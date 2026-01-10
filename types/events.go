package types

type DiscordEvent interface {
	Unmarshal(data []byte) (DiscordEvent, error)
}

type DiscordEventType string

const (
	DiscordEventMessageCreate DiscordEventType = "MESSAGE_CREATE"
	DiscordEventReady         DiscordEventType = "READY"
)
