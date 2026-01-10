package types

type HelloPayloadData struct {
	HeartbeatInterval float64 `json:"heartbeat_interval"`
}

type DiscordEvent interface {
	Unmarshal(data []byte) (DiscordEvent, error)
}

type DiscordEventType string

const (
	DiscordEventMessageCreate DiscordEventType = "MESSAGE_CREATE"
)
