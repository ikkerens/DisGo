package disgo

// Warning: This file has been automatically generated by generate/apimodel/main.go
// Do NOT make changes here, instead adapt model.go and run go generate

import (
	"encoding/json"
	"sync"
)

// Attachment is based on the Discord object with the same name.
// Any fields can be obtained by calling the respective getters.
type Attachment struct {
	session  *Session
	internal *internalAttachment
}

// MarshalJSON is used to convert this object into its json representation for Discord
func (s *Attachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.internal)
}

// UnmarshalJSON is used to convert json discord objects back into their respective structs
func (s *Attachment) UnmarshalJSON(b []byte) error {
	s.internal = &internalAttachment{}
	return json.Unmarshal(b, &s.internal)
}

// ID is used to export the ID from this struct.
func (s *Attachment) ID() Snowflake {
	return s.internal.ID
}

// Filename is used to export the Filename from this struct.
func (s *Attachment) Filename() string {
	return s.internal.Filename
}

// Size is used to export the Size from this struct.
func (s *Attachment) Size() int {
	return s.internal.Size
}

// URL is used to export the URL from this struct.
func (s *Attachment) URL() string {
	return s.internal.URL
}

// ProxyURL is used to export the ProxyURL from this struct.
func (s *Attachment) ProxyURL() string {
	return s.internal.ProxyURL
}

// Height is used to export the Height from this struct.
func (s *Attachment) Height() int {
	return s.internal.Height
}

// Width is used to export the Width from this struct.
func (s *Attachment) Width() int {
	return s.internal.Width
}

// Channel is based on the Discord object with the same name.
// Any fields can be obtained by calling the respective getters.
type Channel struct {
	session  *Session
	internal *internalChannel

	lock *sync.RWMutex
}

// MarshalJSON is used to convert this object into its json representation for Discord
func (s *Channel) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.internal)
}

// UnmarshalJSON is used to convert json discord objects back into their respective structs
func (s *Channel) UnmarshalJSON(b []byte) error {
	id := IDObject{}
	if err := json.Unmarshal(b, &id); err != nil {
		return err
	}

	registered := objects.registerChannel(&id)
	registered.lock.Lock()
	defer registered.lock.Unlock()

	s.lock = registered.lock
	s.internal = registered.internal
	return json.Unmarshal(b, &s.internal)
}

func (s *Channel) setSession(session *Session) {
	s.session = session
}

// ID is used to export the ID from this struct.
func (s *Channel) ID() Snowflake {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.ID
}

// GuildID is used to export the GuildID from this struct.
func (s *Channel) GuildID() Snowflake {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.GuildID
}

// Name is used to export the Name from this struct.
func (s *Channel) Name() string {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Name
}

// Type is used to export the Type from this struct.
func (s *Channel) Type() string {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Type
}

// Position is used to export the Position from this struct.
func (s *Channel) Position() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Position
}

// IsPrivate is used to export the IsPrivate from this struct.
func (s *Channel) IsPrivate() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.IsPrivate
}

// PermissionOverwrites is used to export the PermissionOverwrites from this struct.
func (s *Channel) PermissionOverwrites() []Overwrite {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.PermissionOverwrites
}

// Topic is used to export the Topic from this struct.
func (s *Channel) Topic() string {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Topic
}

// LastMessageID is used to export the LastMessageID from this struct.
func (s *Channel) LastMessageID() Snowflake {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.LastMessageID
}

// Bitrate is used to export the Bitrate from this struct.
func (s *Channel) Bitrate() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Bitrate
}

// UserLimit is used to export the UserLimit from this struct.
func (s *Channel) UserLimit() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.UserLimit
}

// Recipient is used to export the Recipient from this struct.
func (s *Channel) Recipient() *User {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Recipient
}

// Emoji is based on the Discord object with the same name.
// Any fields can be obtained by calling the respective getters.
type Emoji struct {
	session  *Session
	internal *internalEmoji
}

// MarshalJSON is used to convert this object into its json representation for Discord
func (s *Emoji) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.internal)
}

// UnmarshalJSON is used to convert json discord objects back into their respective structs
func (s *Emoji) UnmarshalJSON(b []byte) error {
	s.internal = &internalEmoji{}
	return json.Unmarshal(b, &s.internal)
}

// ID is used to export the ID from this struct.
func (s *Emoji) ID() Snowflake {
	return s.internal.ID
}

// Name is used to export the Name from this struct.
func (s *Emoji) Name() string {
	return s.internal.Name
}

// Roles is used to export the Roles from this struct.
func (s *Emoji) Roles() []*Role {
	return s.internal.Roles
}

// RequireColons is used to export the RequireColons from this struct.
func (s *Emoji) RequireColons() bool {
	return s.internal.RequireColons
}

// Managed is used to export the Managed from this struct.
func (s *Emoji) Managed() bool {
	return s.internal.Managed
}

// Game is based on the Discord object with the same name.
// Any fields can be obtained by calling the respective getters.
type Game struct {
	session  *Session
	internal *internalGame
}

// MarshalJSON is used to convert this object into its json representation for Discord
func (s *Game) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.internal)
}

// UnmarshalJSON is used to convert json discord objects back into their respective structs
func (s *Game) UnmarshalJSON(b []byte) error {
	s.internal = &internalGame{}
	return json.Unmarshal(b, &s.internal)
}

// Name is used to export the Name from this struct.
func (s *Game) Name() string {
	return s.internal.Name
}

// Type is used to export the Type from this struct.
func (s *Game) Type() int {
	return s.internal.Type
}

// URL is used to export the URL from this struct.
func (s *Game) URL() string {
	return s.internal.URL
}

// Guild is based on the Discord object with the same name.
// Any fields can be obtained by calling the respective getters.
type Guild struct {
	session  *Session
	internal *internalGuild

	lock *sync.RWMutex
}

// MarshalJSON is used to convert this object into its json representation for Discord
func (s *Guild) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.internal)
}

// UnmarshalJSON is used to convert json discord objects back into their respective structs
func (s *Guild) UnmarshalJSON(b []byte) error {
	id := IDObject{}
	if err := json.Unmarshal(b, &id); err != nil {
		return err
	}

	registered := objects.registerGuild(&id)
	registered.lock.Lock()
	defer registered.lock.Unlock()

	s.lock = registered.lock
	s.internal = registered.internal
	return json.Unmarshal(b, &s.internal)
}

func (s *Guild) setSession(session *Session) {
	s.session = session
}

// ID is used to export the ID from this struct.
func (s *Guild) ID() Snowflake {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.ID
}

// Name is used to export the Name from this struct.
func (s *Guild) Name() string {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Name
}

// IconHash is used to export the IconHash from this struct.
func (s *Guild) IconHash() string {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.IconHash
}

// SplashHash is used to export the SplashHash from this struct.
func (s *Guild) SplashHash() string {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.SplashHash
}

// OwnerID is used to export the OwnerID from this struct.
func (s *Guild) OwnerID() Snowflake {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.OwnerID
}

// Region is used to export the Region from this struct.
func (s *Guild) Region() string {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Region
}

// AFKChannelID is used to export the AFKChannelID from this struct.
func (s *Guild) AFKChannelID() Snowflake {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.AFKChannelID
}

// AFKTimeout is used to export the AFKTimeout from this struct.
func (s *Guild) AFKTimeout() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.AFKTimeout
}

// EmbedEnabled is used to export the EmbedEnabled from this struct.
func (s *Guild) EmbedEnabled() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.EmbedEnabled
}

// EmbedChannelID is used to export the EmbedChannelID from this struct.
func (s *Guild) EmbedChannelID() Snowflake {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.EmbedChannelID
}

// VerificationLevel is used to export the VerificationLevel from this struct.
func (s *Guild) VerificationLevel() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.VerificationLevel
}

// DefaultMessageNotifications is used to export the DefaultMessageNotifications from this struct.
func (s *Guild) DefaultMessageNotifications() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.DefaultMessageNotifications
}

// Roles is used to export the Roles from this struct.
func (s *Guild) Roles() []*Role {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Roles
}

// Emojis is used to export the Emojis from this struct.
func (s *Guild) Emojis() []Emoji {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Emojis
}

// Features is used to export the Features from this struct.
func (s *Guild) Features() []json.RawMessage {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Features
}

// MFALevel is used to export the MFALevel from this struct.
func (s *Guild) MFALevel() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.MFALevel
}

// JoinedAt is used to export the JoinedAt from this struct.
func (s *Guild) JoinedAt() DiscordTime {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.JoinedAt
}

// Large is used to export the Large from this struct.
func (s *Guild) Large() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Large
}

// Unavailable is used to export the Unavailable from this struct.
func (s *Guild) Unavailable() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Unavailable
}

// MemberCount is used to export the MemberCount from this struct.
func (s *Guild) MemberCount() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.MemberCount
}

// VoiceStates is used to export the VoiceStates from this struct.
func (s *Guild) VoiceStates() []json.RawMessage {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.VoiceStates
}

// Members is used to export the Members from this struct.
func (s *Guild) Members() []*GuildMember {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Members
}

// Channels is used to export the Channels from this struct.
func (s *Guild) Channels() []*Channel {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Channels
}

// Presences is used to export the Presences from this struct.
func (s *Guild) Presences() []Presence {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Presences
}

// GuildMember is based on the Discord object with the same name.
// Any fields can be obtained by calling the respective getters.
type GuildMember struct {
	session  *Session
	internal *internalGuildMember
}

// MarshalJSON is used to convert this object into its json representation for Discord
func (s *GuildMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.internal)
}

// UnmarshalJSON is used to convert json discord objects back into their respective structs
func (s *GuildMember) UnmarshalJSON(b []byte) error {
	s.internal = &internalGuildMember{}
	return json.Unmarshal(b, &s.internal)
}

// User is used to export the User from this struct.
func (s *GuildMember) User() *User {
	return s.internal.User
}

// Nick is used to export the Nick from this struct.
func (s *GuildMember) Nick() string {
	return s.internal.Nick
}

// RolesIDs is used to export the RolesIDs from this struct.
func (s *GuildMember) RolesIDs() []Snowflake {
	return s.internal.RolesIDs
}

// JoinedAt is used to export the JoinedAt from this struct.
func (s *GuildMember) JoinedAt() DiscordTime {
	return s.internal.JoinedAt
}

// Deaf is used to export the Deaf from this struct.
func (s *GuildMember) Deaf() bool {
	return s.internal.Deaf
}

// Mute is used to export the Mute from this struct.
func (s *GuildMember) Mute() bool {
	return s.internal.Mute
}

// Message is based on the Discord object with the same name.
// Any fields can be obtained by calling the respective getters.
type Message struct {
	session  *Session
	internal *internalMessage

	lock *sync.RWMutex
}

// MarshalJSON is used to convert this object into its json representation for Discord
func (s *Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.internal)
}

// UnmarshalJSON is used to convert json discord objects back into their respective structs
func (s *Message) UnmarshalJSON(b []byte) error {
	id := IDObject{}
	if err := json.Unmarshal(b, &id); err != nil {
		return err
	}

	registered := objects.registerMessage(&id)
	registered.lock.Lock()
	defer registered.lock.Unlock()

	s.lock = registered.lock
	s.internal = registered.internal
	return json.Unmarshal(b, &s.internal)
}

func (s *Message) setSession(session *Session) {
	s.session = session
}

// ID is used to export the ID from this struct.
func (s *Message) ID() Snowflake {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.ID
}

// ChannelID is used to export the ChannelID from this struct.
func (s *Message) ChannelID() Snowflake {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.ChannelID
}

// Author is used to export the Author from this struct.
func (s *Message) Author() *User {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Author
}

// Content is used to export the Content from this struct.
func (s *Message) Content() string {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Content
}

// Timestamp is used to export the Timestamp from this struct.
func (s *Message) Timestamp() DiscordTime {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Timestamp
}

// EditedTimestamp is used to export the EditedTimestamp from this struct.
func (s *Message) EditedTimestamp() DiscordTime {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.EditedTimestamp
}

// TTS is used to export the TTS from this struct.
func (s *Message) TTS() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.TTS
}

// MentionEveryone is used to export the MentionEveryone from this struct.
func (s *Message) MentionEveryone() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.MentionEveryone
}

// Mentions is used to export the Mentions from this struct.
func (s *Message) Mentions() []*User {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Mentions
}

// MentionRoles is used to export the MentionRoles from this struct.
func (s *Message) MentionRoles() []*Role {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.MentionRoles
}

// Attachments is used to export the Attachments from this struct.
func (s *Message) Attachments() []Attachment {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Attachments
}

// Embeds is used to export the Embeds from this struct.
func (s *Message) Embeds() []Embed {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Embeds
}

// Reactions is used to export the Reactions from this struct.
func (s *Message) Reactions() []Reaction {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Reactions
}

// NOnce is used to export the NOnce from this struct.
func (s *Message) NOnce() Snowflake {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.NOnce
}

// Pinned is used to export the Pinned from this struct.
func (s *Message) Pinned() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Pinned
}

// WebhookID is used to export the WebhookID from this struct.
func (s *Message) WebhookID() string {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.WebhookID
}

// Presence is based on the Discord object with the same name.
// Any fields can be obtained by calling the respective getters.
type Presence struct {
	session  *Session
	internal *internalPresence
}

// MarshalJSON is used to convert this object into its json representation for Discord
func (s *Presence) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.internal)
}

// UnmarshalJSON is used to convert json discord objects back into their respective structs
func (s *Presence) UnmarshalJSON(b []byte) error {
	s.internal = &internalPresence{}
	return json.Unmarshal(b, &s.internal)
}

// User is used to export the User from this struct.
func (s *Presence) User() *User {
	return s.internal.User
}

// Roles is used to export the Roles from this struct.
func (s *Presence) Roles() []Snowflake {
	return s.internal.Roles
}

// Game is used to export the Game from this struct.
func (s *Presence) Game() Game {
	return s.internal.Game
}

// GuildID is used to export the GuildID from this struct.
func (s *Presence) GuildID() Snowflake {
	return s.internal.GuildID
}

// Status is used to export the Status from this struct.
func (s *Presence) Status() string {
	return s.internal.Status
}

// Reaction is based on the Discord object with the same name.
// Any fields can be obtained by calling the respective getters.
type Reaction struct {
	session  *Session
	internal *internalReaction
}

// MarshalJSON is used to convert this object into its json representation for Discord
func (s *Reaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.internal)
}

// UnmarshalJSON is used to convert json discord objects back into their respective structs
func (s *Reaction) UnmarshalJSON(b []byte) error {
	s.internal = &internalReaction{}
	return json.Unmarshal(b, &s.internal)
}

// Count is used to export the Count from this struct.
func (s *Reaction) Count() int {
	return s.internal.Count
}

// Me is used to export the Me from this struct.
func (s *Reaction) Me() bool {
	return s.internal.Me
}

// Emoji is used to export the Emoji from this struct.
func (s *Reaction) Emoji() *Emoji {
	return s.internal.Emoji
}

// Role is based on the Discord object with the same name.
// Any fields can be obtained by calling the respective getters.
type Role struct {
	session  *Session
	internal *internalRole

	lock *sync.RWMutex
}

// MarshalJSON is used to convert this object into its json representation for Discord
func (s *Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.internal)
}

// UnmarshalJSON is used to convert json discord objects back into their respective structs
func (s *Role) UnmarshalJSON(b []byte) error {
	id := IDObject{}
	if err := json.Unmarshal(b, &id); err != nil {
		return err
	}

	registered := objects.registerRole(&id)
	registered.lock.Lock()
	defer registered.lock.Unlock()

	s.lock = registered.lock
	s.internal = registered.internal
	return json.Unmarshal(b, &s.internal)
}

func (s *Role) setSession(session *Session) {
	s.session = session
}

// ID is used to export the ID from this struct.
func (s *Role) ID() Snowflake {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.ID
}

// Name is used to export the Name from this struct.
func (s *Role) Name() string {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Name
}

// Color is used to export the Color from this struct.
func (s *Role) Color() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Color
}

// Hoist is used to export the Hoist from this struct.
func (s *Role) Hoist() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Hoist
}

// Position is used to export the Position from this struct.
func (s *Role) Position() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Position
}

// Permissions is used to export the Permissions from this struct.
func (s *Role) Permissions() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Permissions
}

// Managed is used to export the Managed from this struct.
func (s *Role) Managed() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Managed
}

// Mentionable is used to export the Mentionable from this struct.
func (s *Role) Mentionable() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Mentionable
}

// User is based on the Discord object with the same name.
// Any fields can be obtained by calling the respective getters.
type User struct {
	session  *Session
	internal *internalUser

	lock *sync.RWMutex
}

// MarshalJSON is used to convert this object into its json representation for Discord
func (s *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.internal)
}

// UnmarshalJSON is used to convert json discord objects back into their respective structs
func (s *User) UnmarshalJSON(b []byte) error {
	id := IDObject{}
	if err := json.Unmarshal(b, &id); err != nil {
		return err
	}

	registered := objects.registerUser(&id)
	registered.lock.Lock()
	defer registered.lock.Unlock()

	s.lock = registered.lock
	s.internal = registered.internal
	return json.Unmarshal(b, &s.internal)
}

func (s *User) setSession(session *Session) {
	s.session = session
}

// ID is used to export the ID from this struct.
func (s *User) ID() Snowflake {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.ID
}

// Username is used to export the Username from this struct.
func (s *User) Username() string {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Username
}

// Discriminator is used to export the Discriminator from this struct.
func (s *User) Discriminator() string {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Discriminator
}

// AvatarHash is used to export the AvatarHash from this struct.
func (s *User) AvatarHash() string {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.AvatarHash
}

// Bot is used to export the Bot from this struct.
func (s *User) Bot() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Bot
}

// MFAEnabled is used to export the MFAEnabled from this struct.
func (s *User) MFAEnabled() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.MFAEnabled
}

// Verified is used to export the Verified from this struct.
func (s *User) Verified() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.Verified
}

// EMail is used to export the EMail from this struct.
func (s *User) EMail() string {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.internal.EMail
}
