package types

import "encoding/json"

type DiscordMessageCreateEvent struct {
	DiscordMessage
	GuildID  *string        `json:"guild_id"`
	Member   *GuildMember   `json:"member,omitempty"`
	Mentions *[]DiscordUser `json:"mentions"`
}

func (e DiscordMessageCreateEvent) Unmarshal(data []byte) (DiscordEvent, error) {
	var event DiscordMessageCreateEvent
	err := json.Unmarshal(data, &event)
	return event, err
}
