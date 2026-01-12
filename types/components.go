package types

type DiscordComponentType int

type AnyComponent interface {
	Type() DiscordComponentType
	UnmarshalJSON(data []byte) error
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

type AnyApplicationCommandInteractionOption interface {
	GetName() string
	GetType() DiscordApplicationCommandInteractionOptionType
}

type DiscordApplicationCommandInteractionOptionString struct {
	Name    string                                         `json:"name"`
	Type    DiscordApplicationCommandInteractionOptionType `json:"type"`
	Value   string                                         `json:"value"`
	Focused *bool                                          `json:"focused,omitempty"`
}

func (o DiscordApplicationCommandInteractionOptionString) GetName() string {
	return o.Name
}

func (o DiscordApplicationCommandInteractionOptionString) GetType() DiscordApplicationCommandInteractionOptionType {
	return o.Type
}
