package disgo

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/slf4go/logger"
)

type Shard struct {
	session *Session

	webSocket *websocket.Conn
	shard     int
	sessionID string
	sequence  int
	heartbeat int

	mutex        sync.Mutex
	closeMessage chan int
	stopListen   chan bool
	stopRead     chan bool
}

func connectShard(session *Session, shard int) (*Shard, error) {
	conn, _, err := websocket.DefaultDialer.Dial(session.wsUrl, http.Header{})
	if err != nil {
		return nil, err
	}

	s := &Shard{
		session:      session,
		webSocket:    conn,
		shard:        shard,
		closeMessage: make(chan int, 1),
		stopListen:   make(chan bool),
		stopRead:     make(chan bool),
	}

	if err = s.handshake(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Shard) handshake() error {
	s.webSocket.SetCloseHandler(s.onClose)

	helloFrame, err := s.readFrame()
	if err != nil {
		return err
	} else if helloFrame.Op != opHello {
		return errors.New("First frame sent from the Discord Gateway is not hello")
	}

	hello := helloPayload{}
	if err = json.Unmarshal(helloFrame.Data, &hello); err != nil {
		return err
	}

	logger.Debugf("Connected to Discord servers: %s", strings.Join(hello.Servers, ", "))
	logger.Debugf("Setting up a heartbeat interval of %d ms", hello.HeartbeatInterval)
	s.heartbeat = hello.HeartbeatInterval

	if err = s.identify(); err != nil {
		return err
	}

	go s.mainLoop()
	return nil
}

func (s *Shard) identify() error {
	if s.sessionID != "" {
		logger.Debugf("Resuming connectiong starting at sequence %d.", s.sequence)
		s.sendFrame(&gatewayFrame{opResume, resumePayload{
			Token:     s.session.token,
			SessionID: s.sessionID,
			Sequence:  s.sequence,
		}})
	} else {
		logger.Debugf("Identifying to websocket")
		s.sendFrame(&gatewayFrame{opIdentify, identifyPayload{
			Token:          s.session.token,
			Compress:       true,
			LargeThreshold: 250,
			Shard:          [2]int{s.shard, cap(s.session.shards)},
			Properties: propertiesPayload{
				OS:      runtime.GOOS,
				Browser: "DisGo",
				Device:  "DisGo",
			},
		}})
	}

	frame, err := s.readFrame()
	if err != nil {
		return err
	}

	if frame.Op == opDispatch {
		if frame.EventName == "READY" {
			ready := ReadyEvent{}
			if err := json.Unmarshal(frame.Data, &ready); err != nil {
				return err
			}

			s.sessionID = ready.SessionID
			s.session.dispatchEvent(frame)
			return nil
		} else if frame.EventName == "RESUMED" {
			s.session.dispatchEvent(frame)
			return nil
		} else {
			return errors.New(fmt.Sprintf("Discord sent event of type '%s', expected 'READY'", frame.EventName))
		}
	} else if frame.Op == opInvalidSession {
		s.sessionID = "" // Invalidate session and retry
		s.identify()
		return nil
	} else {
		return errors.New(fmt.Sprintf("Unexpected opCode received from Discord: %d", frame.Op))
	}
}

func (s *Shard) mainLoop() {
	logger.Debugf("Starting main loop for shard [%d/%d]", s.shard+1, cap(s.session.shards))
	defer logger.Debugf("Exiting main loop for shard [%d/%d]", s.shard+1, cap(s.session.shards))

	heartbeat := time.NewTicker(time.Duration(s.heartbeat) * time.Millisecond)
	defer heartbeat.Stop()
	sentHeartBeat := false

	reader := make(chan *receivedFrame, 1)
	go s.readWebSocket(reader)

	for {
		select {
		case <-heartbeat.C:
			if !sentHeartBeat {
				s.sendFrame(&gatewayFrame{opHeartbeat, s.sequence})
				sentHeartBeat = true
			} else {
				s.disconnect(websocket.CloseAbnormalClosure, "Did not respond to previous ping")
			}
		case <-s.stopListen:
			return
		case frame := <-reader:
			switch opCode := frame.Op; opCode {
			case opHeartbeat:
				s.sendFrame(&gatewayFrame{Op: opHeartbeatAck})
			case opHeartbeatAck:
				sentHeartBeat = false
			case opReconnect:
				s.disconnect(websocket.CloseNormalClosure, "op Reconnect")
			case opDispatch:
				s.session.dispatchEvent(frame)
			default:
				logger.Errorf("Unknown opCode received: %d", opCode)
			}
		}
	}
}

func (s *Shard) sendFrame(frame *gatewayFrame) {
	logger.Debugf("Sending frame with opCode: %d", frame.Op)
	s.webSocket.WriteJSON(frame)
}

func (s *Shard) readWebSocket(reader chan *receivedFrame) {
	logger.Debugf("Starting read loop for shard [%d/%d]", s.shard+1, cap(s.session.shards))
	defer logger.Debugf("Exiting read loop for shard [%d/%d]", s.shard+1, cap(s.session.shards))

	for {
		s.mutex.Lock()
		s.mutex.Unlock()
		select {
		case <-s.stopRead:
			return
		default:
			frame, err := s.readFrame()
			if err != nil {
				if !s.session.shuttingDown {
					go s.disconnect(websocket.CloseAbnormalClosure, err.Error())
				}
				<-s.stopRead // Wait for the connection to be closed
				return
			}

			reader <- frame
		}
	}

}

func (s *Shard) readFrame() (*receivedFrame, error) {
	msgType, msg, err := s.webSocket.ReadMessage()
	if err != nil {
		return nil, err
	}

	var reader io.Reader
	reader = bytes.NewBuffer(msg)

	if msgType == websocket.BinaryMessage {
		zReader, err := zlib.NewReader(reader)
		if err != nil {
			return nil, err
		}

		defer zReader.Close()

		reader = zReader
	}

	frame := receivedFrame{Sequence: -1}
	json.NewDecoder(reader).Decode(&frame)

	logger.Debugf("Received frame with opCode: %d", frame.Op)

	if frame.Sequence != -1 {
		logger.Tracef("Last sequence received set to %d.", frame.Sequence)
		s.sequence = frame.Sequence
	}
	return &frame, nil
}

func (s *Shard) reconnect() {
	if !s.session.shuttingDown {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		logger.Noticef("Reconnecting shard [%d/%d]", s.shard+1, cap(s.session.shards))
		conn, _, err := websocket.DefaultDialer.Dial(s.session.wsUrl, http.Header{})

		if err == nil {
			s.webSocket = conn
			err = s.handshake()
		}

		if err != nil {
			logger.ErrorE(err)
			time.Sleep(1 * time.Second)
			go s.reconnect()
		}
	}
}

func (s *Shard) onClose(code int, text string) error {
	logger.Infof("Received Close Frame from Discord. Code: %d. Text: %s", code, text)
	s.closeMessage <- code
	s.reconnect()
	return nil
}

func (s *Shard) disconnect(code int, text string) {
	s.mutex.Lock()
	s.stopListen <- true

	err := s.webSocket.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(code, text))
	if err != nil {
		logger.ErrorE(err)
	}

	select {
	case code := <-s.closeMessage:
		logger.Infof("Discord connection closed with code %d", code)
	case <-time.After(1 * time.Second):
		logger.Warn("Discord did not reply to the close message, force-closing connection.")
	}

	s.webSocket.Close()
	s.mutex.Unlock()
	s.stopRead <- true

	s.reconnect()
}
