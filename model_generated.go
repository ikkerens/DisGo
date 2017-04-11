package disgo

// Warning: This file has been automatically generated by generate/apimodel/main.go
// Do NOT make changes here, instead adapt model.go and run go generate

import (
	"encoding/json"
	"time"
)

type Attachment struct {
	session  *Session
	internal *internalAttachment
}

func (s *Attachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.internal)
}

func (s *Attachment) UnmarshalJSON(b []byte) error {
	s.internal = &internalAttachment{}
	return json.Unmarshal(b, &s.internal)
}

func (s *Attachment) ID() Snowflake {
	return s.internal.ID
}

func (s *Attachment) Filename() string {
	return s.internal.Filename
}

func (s *Attachment) Size() int {
	return s.internal.Size
}

func (s *Attachment) URL() string {
	return s.internal.URL
}

func (s *Attachment) ProxyURL() string {
	return s.internal.ProxyURL
}

func (s *Attachment) Height() int {
	return s.internal.Height
}

func (s *Attachment) Width() int {
	return s.internal.Width
}

type Channel struct {
	session  *Session
	internal *internalChannel
}

func (s *Channel) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.internal)
}

func (s *Channel) UnmarshalJSON(b []byte) error {
	s.internal = &internalChannel{}
	return json.Unmarshal(b, &s.internal)
}

func (s *Channel) ID() Snowflake {
	return s.internal.ID
}

func (s *Channel) GuildID() Snowflake {
	return s.internal.GuildID
}

func (s *Channel) Name() string {
	return s.internal.Name
}

func (s *Channel) Type() string {
	return s.internal.Type
}

func (s *Channel) Position() int {
	return s.internal.Position
}

func (s *Channel) IsPrivate() bool {
	return s.internal.IsPrivate
}

func (s *Channel) PermissionOverwrites() []Overwrite {
	return s.internal.PermissionOverwrites
}

func (s *Channel) Topic() string {
	return s.internal.Topic
}

func (s *Channel) LastMessageID() Snowflake {
	return s.internal.LastMessageID
}

func (s *Channel) Bitrate() int {
	return s.internal.Bitrate
}

func (s *Channel) UserLimit() int {
	return s.internal.UserLimit
}

type DMChannel struct {
	session  *Session
	internal *internalDMChannel
}

func (s *DMChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.internal)
}

func (s *DMChannel) UnmarshalJSON(b []byte) error {
	s.internal = &internalDMChannel{}
	return json.Unmarshal(b, &s.internal)
}

func (s *DMChannel) ID() Snowflake {
	return s.internal.ID
}

func (s *DMChannel) IsPrivate() bool {
	return s.internal.IsPrivate
}

func (s *DMChannel) Recipient() *User {
	return s.internal.Recipient
}

func (s *DMChannel) LastMessageID() Snowflake {
	return s.internal.LastMessageID
}

type Emoji struct {
	session  *Session
	internal *internalEmoji
}

func (s *Emoji) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.internal)
}

func (s *Emoji) UnmarshalJSON(b []byte) error {
	s.internal = &internalEmoji{}
	return json.Unmarshal(b, &s.internal)
}

func (s *Emoji) ID() Snowflake {
	return s.internal.ID
}

func (s *Emoji) Name() string {
	return s.internal.Name
}

func (s *Emoji) Roles() []json.RawMessage {
	return s.internal.Roles
}

func (s *Emoji) RequireColons() bool {
	return s.internal.RequireColons
}

func (s *Emoji) Managed() bool {
	return s.internal.Managed
}

type Game struct {
	session  *Session
	internal *internalGame
}

func (s *Game) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.internal)
}

func (s *Game) UnmarshalJSON(b []byte) error {
	s.internal = &internalGame{}
	return json.Unmarshal(b, &s.internal)
}

func (s *Game) Name() string {
	return s.internal.Name
}

func (s *Game) Type() int {
	return s.internal.Type
}

func (s *Game) URL() string {
	return s.internal.URL
}

type Guild struct {
	session  *Session
	internal *internalGuild
}

func (s *Guild) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.internal)
}

func (s *Guild) UnmarshalJSON(b []byte) error {
	s.internal = &internalGuild{}
	return json.Unmarshal(b, &s.internal)
}

func (s *Guild) ID() Snowflake {
	return s.internal.ID
}

func (s *Guild) Name() string {
	return s.internal.Name
}

func (s *Guild) IconHash() string {
	return s.internal.IconHash
}

func (s *Guild) SplashHash() string {
	return s.internal.SplashHash
}

func (s *Guild) OwnerID() Snowflake {
	return s.internal.OwnerID
}

func (s *Guild) Region() string {
	return s.internal.Region
}

func (s *Guild) AFKChannelID() Snowflake {
	return s.internal.AFKChannelID
}

func (s *Guild) AFKTimeout() int {
	return s.internal.AFKTimeout
}

func (s *Guild) EmbedEnabled() bool {
	return s.internal.EmbedEnabled
}

func (s *Guild) EmbedChannelID() Snowflake {
	return s.internal.EmbedChannelID
}

func (s *Guild) VerificationLevel() int {
	return s.internal.VerificationLevel
}

func (s *Guild) DefaultMessageNotifications() int {
	return s.internal.DefaultMessageNotifications
}

func (s *Guild) Roles() []*Role {
	return s.internal.Roles
}

func (s *Guild) Emojis() []*Emoji {
	return s.internal.Emojis
}

func (s *Guild) Features() []json.RawMessage {
	return s.internal.Features
}

func (s *Guild) MFALevel() int {
	return s.internal.MFALevel
}

func (s *Guild) JoinedAt() time.Time {
	return s.internal.JoinedAt
}

func (s *Guild) Large() bool {
	return s.internal.Large
}

func (s *Guild) Unavailable() bool {
	return s.internal.Unavailable
}

func (s *Guild) MemberCount() int {
	return s.internal.MemberCount
}

func (s *Guild) VoiceStates() []json.RawMessage {
	return s.internal.VoiceStates
}

func (s *Guild) Members() []GuildMember {
	return s.internal.Members
}

func (s *Guild) Channels() []*Channel {
	return s.internal.Channels
}

func (s *Guild) Presences() []Presence {
	return s.internal.Presences
}

type GuildMember struct {
	session  *Session
	internal *internalGuildMember
}

func (s *GuildMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.internal)
}

func (s *GuildMember) UnmarshalJSON(b []byte) error {
	s.internal = &internalGuildMember{}
	return json.Unmarshal(b, &s.internal)
}

func (s *GuildMember) User() *User {
	return s.internal.User
}

func (s *GuildMember) Nick() string {
	return s.internal.Nick
}

func (s *GuildMember) Roles() []Snowflake {
	return s.internal.Roles
}

func (s *GuildMember) JoinedAt() time.Time {
	return s.internal.JoinedAt
}

func (s *GuildMember) Deaf() bool {
	return s.internal.Deaf
}

func (s *GuildMember) Mute() bool {
	return s.internal.Mute
}

type Message struct {
	session  *Session
	internal *internalMessage
}

func (s *Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.internal)
}

func (s *Message) UnmarshalJSON(b []byte) error {
	s.internal = &internalMessage{}
	return json.Unmarshal(b, &s.internal)
}

func (s *Message) ID() Snowflake {
	return s.internal.ID
}

func (s *Message) ChannelID() Snowflake {
	return s.internal.ChannelID
}

func (s *Message) Author() *User {
	return s.internal.Author
}

func (s *Message) Content() string {
	return s.internal.Content
}

func (s *Message) Timestamp() time.Time {
	return s.internal.Timestamp
}

func (s *Message) EditedTimestamp() time.Time {
	return s.internal.EditedTimestamp
}

func (s *Message) TTS() bool {
	return s.internal.TTS
}

func (s *Message) MentionEveryone() bool {
	return s.internal.MentionEveryone
}

func (s *Message) MentionRoles() []*Role {
	return s.internal.MentionRoles
}

func (s *Message) Attachments() []Attachment {
	return s.internal.Attachments
}

func (s *Message) Embeds() []Embed {
	return s.internal.Embeds
}

func (s *Message) Reactions() []Reaction {
	return s.internal.Reactions
}

func (s *Message) NOnce() Snowflake {
	return s.internal.NOnce
}

func (s *Message) Pinned() bool {
	return s.internal.Pinned
}

func (s *Message) WebhookID() string {
	return s.internal.WebhookID
}

type Overwrite struct {
	session  *Session
	internal *internalOverwrite
}

func (s *Overwrite) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.internal)
}

func (s *Overwrite) UnmarshalJSON(b []byte) error {
	s.internal = &internalOverwrite{}
	return json.Unmarshal(b, &s.internal)
}

func (s *Overwrite) ID() Snowflake {
	return s.internal.ID
}

func (s *Overwrite) Type() string {
	return s.internal.Type
}

func (s *Overwrite) Allow() int {
	return s.internal.Allow
}

func (s *Overwrite) Deny() int {
	return s.internal.Deny
}

type Presence struct {
	session  *Session
	internal *internalPresence
}

func (s *Presence) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.internal)
}

func (s *Presence) UnmarshalJSON(b []byte) error {
	s.internal = &internalPresence{}
	return json.Unmarshal(b, &s.internal)
}

func (s *Presence) User() *User {
	return s.internal.User
}

func (s *Presence) Roles() []Snowflake {
	return s.internal.Roles
}

func (s *Presence) Game() Game {
	return s.internal.Game
}

func (s *Presence) GuildID() Snowflake {
	return s.internal.GuildID
}

func (s *Presence) Status() string {
	return s.internal.Status
}

type Reaction struct {
	session  *Session
	internal *internalReaction
}

func (s *Reaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.internal)
}

func (s *Reaction) UnmarshalJSON(b []byte) error {
	s.internal = &internalReaction{}
	return json.Unmarshal(b, &s.internal)
}

func (s *Reaction) Count() int {
	return s.internal.Count
}

func (s *Reaction) Me() bool {
	return s.internal.Me
}

func (s *Reaction) Emoji() *Emoji {
	return s.internal.Emoji
}

type Role struct {
	session  *Session
	internal *internalRole
}

func (s *Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.internal)
}

func (s *Role) UnmarshalJSON(b []byte) error {
	s.internal = &internalRole{}
	return json.Unmarshal(b, &s.internal)
}

func (s *Role) ID() Snowflake {
	return s.internal.ID
}

func (s *Role) Name() string {
	return s.internal.Name
}

func (s *Role) Color() int {
	return s.internal.Color
}

func (s *Role) Hoist() bool {
	return s.internal.Hoist
}

func (s *Role) Position() int {
	return s.internal.Position
}

func (s *Role) Permissions() int {
	return s.internal.Permissions
}

func (s *Role) Managed() bool {
	return s.internal.Managed
}

func (s *Role) Mentionable() bool {
	return s.internal.Mentionable
}

type User struct {
	session  *Session
	internal *internalUser
}

func (s *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.internal)
}

func (s *User) UnmarshalJSON(b []byte) error {
	s.internal = &internalUser{}
	return json.Unmarshal(b, &s.internal)
}

func (s *User) ID() Snowflake {
	return s.internal.ID
}

func (s *User) Username() string {
	return s.internal.Username
}

func (s *User) Discriminator() string {
	return s.internal.Discriminator
}

func (s *User) AvatarHash() string {
	return s.internal.AvatarHash
}

func (s *User) Bot() bool {
	return s.internal.Bot
}

func (s *User) MFAEnabled() bool {
	return s.internal.MFAEnabled
}

func (s *User) Verified() bool {
	return s.internal.Verified
}

func (s *User) EMail() string {
	return s.internal.EMail
}
