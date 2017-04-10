package disgo

// Warning: This file has been automatically generated by generate/apimodel/main.go
// Do NOT make changes here, instead adapt model.go and run go generate

import (
	"encoding/json"
	"time"
)

type Attachment struct {
	session       *Session `json:"-"`
	discordObject *internalAttachment
}

func (s *Attachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.discordObject)
}

func (s *Attachment) UnmarshalJSON(b []byte) error {
	s.discordObject = &internalAttachment{}
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

type Channel struct {
	session       *Session `json:"-"`
	discordObject *internalChannel
}

func (s *Channel) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.discordObject)
}

func (s *Channel) UnmarshalJSON(b []byte) error {
	s.discordObject = &internalChannel{}
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

type DMChannel struct {
	session       *Session `json:"-"`
	discordObject *internalDMChannel
}

func (s *DMChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.discordObject)
}

func (s *DMChannel) UnmarshalJSON(b []byte) error {
	s.discordObject = &internalDMChannel{}
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

type Emoji struct {
	session       *Session `json:"-"`
	discordObject *internalEmoji
}

func (s *Emoji) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.discordObject)
}

func (s *Emoji) UnmarshalJSON(b []byte) error {
	s.discordObject = &internalEmoji{}
	return json.Unmarshal(b, &s.discordObject)
}

func (s *Emoji) ID() Snowflake {
	return s.discordObject.ID
}

func (s *Emoji) Name() string {
	return s.discordObject.Name
}

func (s *Emoji) Roles() []json.RawMessage {
	return s.discordObject.Roles
}

func (s *Emoji) RequireColons() bool {
	return s.discordObject.RequireColons
}

func (s *Emoji) Managed() bool {
	return s.discordObject.Managed
}

type Game struct {
	session       *Session `json:"-"`
	discordObject *internalGame
}

func (s *Game) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.discordObject)
}

func (s *Game) UnmarshalJSON(b []byte) error {
	s.discordObject = &internalGame{}
	return json.Unmarshal(b, &s.discordObject)
}

func (s *Game) Name() string {
	return s.discordObject.Name
}

func (s *Game) Type() int {
	return s.discordObject.Type
}

func (s *Game) URL() string {
	return s.discordObject.URL
}

type Guild struct {
	session       *Session `json:"-"`
	discordObject *internalGuild
}

func (s *Guild) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.discordObject)
}

func (s *Guild) UnmarshalJSON(b []byte) error {
	s.discordObject = &internalGuild{}
	return json.Unmarshal(b, &s.discordObject)
}

func (s *Guild) ID() Snowflake {
	return s.discordObject.ID
}

func (s *Guild) Name() string {
	return s.discordObject.Name
}

func (s *Guild) IconHash() string {
	return s.discordObject.IconHash
}

func (s *Guild) SplashHash() string {
	return s.discordObject.SplashHash
}

func (s *Guild) OwnerID() Snowflake {
	return s.discordObject.OwnerID
}

func (s *Guild) Region() string {
	return s.discordObject.Region
}

func (s *Guild) AFKChannelID() Snowflake {
	return s.discordObject.AFKChannelID
}

func (s *Guild) AFKTimeout() int {
	return s.discordObject.AFKTimeout
}

func (s *Guild) EmbedEnabled() bool {
	return s.discordObject.EmbedEnabled
}

func (s *Guild) EmbedChannelID() Snowflake {
	return s.discordObject.EmbedChannelID
}

func (s *Guild) VerificationLevel() int {
	return s.discordObject.VerificationLevel
}

func (s *Guild) DefaultMessageNotifications() int {
	return s.discordObject.DefaultMessageNotifications
}

func (s *Guild) Roles() []*Role {
	return s.discordObject.Roles
}

func (s *Guild) Emojis() []*Emoji {
	return s.discordObject.Emojis
}

func (s *Guild) Features() []json.RawMessage {
	return s.discordObject.Features
}

func (s *Guild) MFALevel() int {
	return s.discordObject.MFALevel
}

func (s *Guild) JoinedAt() time.Time {
	return s.discordObject.JoinedAt
}

func (s *Guild) Large() bool {
	return s.discordObject.Large
}

func (s *Guild) Unavailable() bool {
	return s.discordObject.Unavailable
}

func (s *Guild) MemberCount() int {
	return s.discordObject.MemberCount
}

func (s *Guild) VoiceStates() []json.RawMessage {
	return s.discordObject.VoiceStates
}

func (s *Guild) Members() []GuildMember {
	return s.discordObject.Members
}

func (s *Guild) Channels() []*Channel {
	return s.discordObject.Channels
}

func (s *Guild) Presences() []Presence {
	return s.discordObject.Presences
}

type GuildMember struct {
	session       *Session `json:"-"`
	discordObject *internalGuildMember
}

func (s *GuildMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.discordObject)
}

func (s *GuildMember) UnmarshalJSON(b []byte) error {
	s.discordObject = &internalGuildMember{}
	return json.Unmarshal(b, &s.discordObject)
}

func (s *GuildMember) User() *User {
	return s.discordObject.User
}

func (s *GuildMember) Nick() string {
	return s.discordObject.Nick
}

func (s *GuildMember) Roles() []Snowflake {
	return s.discordObject.Roles
}

func (s *GuildMember) JoinedAt() time.Time {
	return s.discordObject.JoinedAt
}

func (s *GuildMember) Deaf() bool {
	return s.discordObject.Deaf
}

func (s *GuildMember) Mute() bool {
	return s.discordObject.Mute
}

type Message struct {
	session       *Session `json:"-"`
	discordObject *internalMessage
}

func (s *Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.discordObject)
}

func (s *Message) UnmarshalJSON(b []byte) error {
	s.discordObject = &internalMessage{}
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

func (s *Message) MentionRoles() []*Role {
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

type Overwrite struct {
	session       *Session `json:"-"`
	discordObject *internalOverwrite
}

func (s *Overwrite) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.discordObject)
}

func (s *Overwrite) UnmarshalJSON(b []byte) error {
	s.discordObject = &internalOverwrite{}
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

type Presence struct {
	session       *Session `json:"-"`
	discordObject *internalPresence
}

func (s *Presence) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.discordObject)
}

func (s *Presence) UnmarshalJSON(b []byte) error {
	s.discordObject = &internalPresence{}
	return json.Unmarshal(b, &s.discordObject)
}

func (s *Presence) User() *User {
	return s.discordObject.User
}

func (s *Presence) Roles() []Snowflake {
	return s.discordObject.Roles
}

func (s *Presence) Game() Game {
	return s.discordObject.Game
}

func (s *Presence) GuildID() Snowflake {
	return s.discordObject.GuildID
}

func (s *Presence) Status() string {
	return s.discordObject.Status
}

type Reaction struct {
	session       *Session `json:"-"`
	discordObject *internalReaction
}

func (s *Reaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.discordObject)
}

func (s *Reaction) UnmarshalJSON(b []byte) error {
	s.discordObject = &internalReaction{}
	return json.Unmarshal(b, &s.discordObject)
}

func (s *Reaction) Count() int {
	return s.discordObject.Count
}

func (s *Reaction) Me() bool {
	return s.discordObject.Me
}

func (s *Reaction) Emoji() *Emoji {
	return s.discordObject.Emoji
}

type Role struct {
	session       *Session `json:"-"`
	discordObject *internalRole
}

func (s *Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.discordObject)
}

func (s *Role) UnmarshalJSON(b []byte) error {
	s.discordObject = &internalRole{}
	return json.Unmarshal(b, &s.discordObject)
}

func (s *Role) ID() Snowflake {
	return s.discordObject.ID
}

func (s *Role) Name() string {
	return s.discordObject.Name
}

func (s *Role) Color() int {
	return s.discordObject.Color
}

func (s *Role) Hoist() bool {
	return s.discordObject.Hoist
}

func (s *Role) Position() int {
	return s.discordObject.Position
}

func (s *Role) Permissions() int {
	return s.discordObject.Permissions
}

func (s *Role) Managed() bool {
	return s.discordObject.Managed
}

func (s *Role) Mentionable() bool {
	return s.discordObject.Mentionable
}

type User struct {
	session       *Session `json:"-"`
	discordObject *internalUser
}

func (s *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.discordObject)
}

func (s *User) UnmarshalJSON(b []byte) error {
	s.discordObject = &internalUser{}
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
