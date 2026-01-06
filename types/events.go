package types

import "time"

type HelloPayloadData struct {
	HeartbeatInterval float64 `json:"heartbeat_interval"`
}

type DiscordEventMessageCreate struct {
	Attachments     []any        `json:"attachments"`
	Author          *User        `json:"author"`
	ChannelID       string       `json:"channel_id"`
	ChannelType     int          `json:"channel_type"`
	Components      []any        `json:"components"`
	Content         string       `json:"content"`
	EditedTimestamp *int64       `json:"edited_timestamp"`
	Embeds          any          `json:"embeds"`
	Flags           int64        `json:"flags"`
	ID              string       `json:"id"`
	Member          *GuildMember `json:"member,omitempty"`
	MentionEveryone bool         `json:"mention_everyone"`
	MentionChannels *any         `json:"mention_channels"`
	MentionRoles    []string     `json:"mention_roles"`
	Mentions        *[]User      `json:"mentions"`
	Nonce           interface{}  `json:"nonce,omitempty"`
	Pinned          bool         `json:"pinned"`
	Timestamp       *time.Time   `json:"timestamp,omitempty"`
	TTS             bool         `json:"tts"`
	Type            int          `json:"type"`
}
