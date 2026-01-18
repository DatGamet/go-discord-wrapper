package types

import (
	"encoding/json"
)

type DiscordComponentType int

const (
	DiscordComponentTypeActionRow        DiscordComponentType = 1
	DiscordComponentTypeButton           DiscordComponentType = 2
	DiscordComponentTypeStringSelectMenu DiscordComponentType = 3
	DiscordComponentTypeTextInput        DiscordComponentType = 4
	DiscordComponentTypeUserSelectMenu   DiscordComponentType = 5
	DiscordComponentTypeRoleSelectMenu   DiscordComponentType = 6
	DiscordComponentTypeMentionableMenu  DiscordComponentType = 7
	DiscordComponentTypeChannelSelect    DiscordComponentType = 8
	DiscordComponentTypeSection          DiscordComponentType = 9
	//TODO
)

func (c DiscordComponentType) IsAnySelectMenu() bool {
	return c == DiscordComponentTypeStringSelectMenu ||
		c == DiscordComponentTypeUserSelectMenu ||
		c == DiscordComponentTypeRoleSelectMenu ||
		c == DiscordComponentTypeMentionableMenu ||
		c == DiscordComponentTypeChannelSelect
}

type AnyComponent interface {
	GetType() DiscordComponentType
}

type ActionRow struct {
	Type       DiscordComponentType `json:"type"`
	ID         *int                 `json:"id"`
	Components []AnyComponent       `json:"components"`
}

func (a *ActionRow) GetType() DiscordComponentType {
	return DiscordComponentTypeActionRow
}

func (a *ActionRow) UnmarshalJSON(data []byte) error {
	type Alias ActionRow

	var raw struct {
		Alias
		Components []json.RawMessage `json:"components"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	*a = ActionRow(raw.Alias)

	for _, c := range raw.Components {
		var probe struct {
			Type DiscordComponentType `json:"type"`
		}

		if err := json.Unmarshal(c, &probe); err != nil {
			return err
		}

		switch probe.Type {
		case DiscordComponentTypeButton:
			var b ButtonComponent
			if err := json.Unmarshal(c, &b); err != nil {
				return err
			}
			a.Components = append(a.Components, &b)
		}
	}

	return nil
}

type ButtonStyle int

const (
	ButtonStylePrimary   ButtonStyle = 1
	ButtonStyleSecondary ButtonStyle = 2
	ButtonStyleSuccess   ButtonStyle = 3
	ButtonStyleDanger    ButtonStyle = 4
	ButtonStyleLink      ButtonStyle = 5
	ButtonStylePremium   ButtonStyle = 6
)

type ButtonComponent struct {
	Type     DiscordComponentType `json:"type"`
	ID       *int                 `json:"id,omitempty"`
	Style    ButtonStyle          `json:"style"`
	Label    string               `json:"label,omitempty"`
	Emoji    *DiscordEmoji        `json:"emoji,omitempty"`
	CustomID string               `json:"custom_id,omitempty"`
	SkuID    *DiscordSnowflake    `json:"sku_id,omitempty"`
	URL      string               `json:"url,omitempty"`
	Disabled bool                 `json:"disabled,omitempty"`
}

func (b ButtonComponent) GetType() DiscordComponentType {
	return DiscordComponentTypeButton
}

type DiscordApplicationCommandInteractionOptionType int

const (
	DiscordApplicationCommandInteractionOptionTypeSubCommand      DiscordApplicationCommandInteractionOptionType = 1
	DiscordApplicationCommandInteractionOptionTypeSubCommandGroup DiscordApplicationCommandInteractionOptionType = 2
	DiscordApplicationCommandInteractionOptionTypeString          DiscordApplicationCommandInteractionOptionType = 3
	DiscordApplicationCommandInteractionOptionTypeInteger         DiscordApplicationCommandInteractionOptionType = 4
	DiscordApplicationCommandInteractionOptionTypeBoolean         DiscordApplicationCommandInteractionOptionType = 5
	DiscordApplicationCommandInteractionOptionTypeUser            DiscordApplicationCommandInteractionOptionType = 6
	DiscordApplicationCommandInteractionOptionTypeChannel         DiscordApplicationCommandInteractionOptionType = 7
	DiscordApplicationCommandInteractionOptionTypeRole            DiscordApplicationCommandInteractionOptionType = 8
	DiscordApplicationCommandInteractionOptionTypeMentionable     DiscordApplicationCommandInteractionOptionType = 9
	DiscordApplicationCommandInteractionOptionTypeNumber          DiscordApplicationCommandInteractionOptionType = 10
	DiscordApplicationCommandInteractionOptionTypeAttachment      DiscordApplicationCommandInteractionOptionType = 11
)
