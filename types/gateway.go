package types

import (
	"encoding/json"
	"fmt"
)

type DiscordAPIVersion uint8

var (
	DiscordAPIBaseString = func(v DiscordAPIVersion) string {
		return fmt.Sprintf("/api/v%d/", v)
	}

	DiscordAPIVersion10 DiscordAPIVersion = 10
	DiscordAPIVersion9  DiscordAPIVersion = 9

	DiscordAPIGatewayRequest = "gateway/bot"
)

type Payload struct {
	Op int             `json:"op"`
	D  json.RawMessage `json:"d"`
	T  string          `json:"t,omitempty"`
	S  *int            `json:"s,omitempty"`
}
