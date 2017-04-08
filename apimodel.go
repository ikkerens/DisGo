package disgo

// Warning: This file has been automatically generated by generate/apimodel/main.go
// Do NOT make changes here, instead adapt model.go and run go generate

import (
	"encoding/json"
	"time"
)

type DMChannel struct {
	discordObject *internalDMChannel
}

func (s *DMChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.discordObject)
}

func (s *DMChannel) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &s.discordObject)
}

func (s *DMChannel) ID() Snowflake {
	return s.discordObject.ID
}

func (s *DMChannel) IsPrivate() bool {
	return s.discordObject.IsPrivate
}

func (s *DMChannel) Recipient() *User {
	return s.discordObject.Recipient
}

func (s *DMChannel) LastMessageID() Snowflake {
	return s.discordObject.LastMessageID
}

type User struct {
	discordObject *internalUser
}

func (s *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.discordObject)
}

func (s *User) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &s.discordObject)
}

func (s *User) ID() Snowflake {
	return s.discordObject.ID
}

func (s *User) Username() string {
	return s.discordObject.Username
}

func (s *User) Discriminator() string {
	return s.discordObject.Discriminator
}

func (s *User) AvatarHash() string {
	return s.discordObject.AvatarHash
}

func (s *User) Bot() bool {
	return s.discordObject.Bot
}

func (s *User) MFAEnabled() bool {
	return s.discordObject.MFAEnabled
}

func (s *User) Verified() bool {
	return s.discordObject.Verified
}

func (s *User) EMail() string {
	return s.discordObject.EMail
}

type Emoji struct {
	discordObject *internalEmoji
}

func (s *Emoji) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.discordObject)
}

func (s *Emoji) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &s.discordObject)
}

func (s *Emoji) ID() Snowflake {
	return s.discordObject.ID
}

func (s *Emoji) Name() string {
	return s.discordObject.Name
}

type Channel struct {
	discordObject *internalChannel
}

func (s *Channel) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.discordObject)
}

func (s *Channel) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &s.discordObject)
}

func (s *Channel) ID() Snowflake {
	return s.discordObject.ID
}

func (s *Channel) GuildID() Snowflake {
	return s.discordObject.GuildID
}

func (s *Channel) Name() string {
	return s.discordObject.Name
}

func (s *Channel) Type() string {
	return s.discordObject.Type
}

func (s *Channel) Position() int {
	return s.discordObject.Position
}

func (s *Channel) IsPrivate() bool {
	return s.discordObject.IsPrivate
}

func (s *Channel) PermissionOverwrites() []Overwrite {
	return s.discordObject.PermissionOverwrites
}

func (s *Channel) Topic() string {
	return s.discordObject.Topic
}

func (s *Channel) LastMessageID() Snowflake {
	return s.discordObject.LastMessageID
}

func (s *Channel) Bitrate() int {
	return s.discordObject.Bitrate
}

func (s *Channel) UserLimit() int {
	return s.discordObject.UserLimit
}

type Reaction struct {
	discordObject *internalReaction
}

func (s *Reaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.discordObject)
}

func (s *Reaction) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &s.discordObject)
}

func (s *Reaction) Count() int {
	return s.discordObject.Count
}

func (s *Reaction) Me() bool {
	return s.discordObject.Me
}

func (s *Reaction) Emoji() Emoji {
	return s.discordObject.Emoji
}

type Overwrite struct {
	discordObject *internalOverwrite
}

func (s *Overwrite) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.discordObject)
}

func (s *Overwrite) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &s.discordObject)
}

func (s *Overwrite) ID() Snowflake {
	return s.discordObject.ID
}

func (s *Overwrite) Type() string {
	return s.discordObject.Type
}

func (s *Overwrite) Allow() int {
	return s.discordObject.Allow
}

func (s *Overwrite) Deny() int {
	return s.discordObject.Deny
}

type Attachment struct {
	discordObject *internalAttachment
}

func (s *Attachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.discordObject)
}

func (s *Attachment) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &s.discordObject)
}

func (s *Attachment) ID() Snowflake {
	return s.discordObject.ID
}

func (s *Attachment) Filename() string {
	return s.discordObject.Filename
}

func (s *Attachment) Size() int {
	return s.discordObject.Size
}

func (s *Attachment) URL() string {
	return s.discordObject.URL
}

func (s *Attachment) ProxyURL() string {
	return s.discordObject.ProxyURL
}

func (s *Attachment) Height() int {
	return s.discordObject.Height
}

func (s *Attachment) Width() int {
	return s.discordObject.Width
}

type Message struct {
	discordObject *internalMessage
}

func (s *Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.discordObject)
}

func (s *Message) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &s.discordObject)
}

func (s *Message) ID() Snowflake {
	return s.discordObject.ID
}

func (s *Message) ChannelID() Snowflake {
	return s.discordObject.ChannelID
}

func (s *Message) Author() *User {
	return s.discordObject.Author
}

func (s *Message) Content() string {
	return s.discordObject.Content
}

func (s *Message) Timestamp() time.Time {
	return s.discordObject.Timestamp
}

func (s *Message) EditedTimestamp() time.Time {
	return s.discordObject.EditedTimestamp
}

func (s *Message) TTS() bool {
	return s.discordObject.TTS
}

func (s *Message) MentionEveryone() bool {
	return s.discordObject.MentionEveryone
}

func (s *Message) MentionRoles() []json.RawMessage {
	return s.discordObject.MentionRoles
}

func (s *Message) Attachments() []Attachment {
	return s.discordObject.Attachments
}

func (s *Message) Embeds() []Embed {
	return s.discordObject.Embeds
}

func (s *Message) Reactions() []Reaction {
	return s.discordObject.Reactions
}

func (s *Message) NOnce() Snowflake {
	return s.discordObject.NOnce
}

func (s *Message) Pinned() bool {
	return s.discordObject.Pinned
}

func (s *Message) WebhookID() string {
	return s.discordObject.WebhookID
}
