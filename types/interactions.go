package types

type DiscordInteractionType int

const (
	DiscordInteractionTypePing                           DiscordInteractionType = 1
	DiscordInteractionTypeApplicationCommand             DiscordInteractionType = 2
	DiscordInteractionTypeMessageComponent               DiscordInteractionType = 3
	DiscordInteractionTypeApplicationCommandAutocomplete DiscordInteractionType = 4
	DiscordInteractionTypeModalSubmit                    DiscordInteractionType = 5
)

type DiscordInteractionApplicationIntegrationType int

const (
	DiscordInteractionApplicationIntegrationTypeGuildInstall DiscordInteractionApplicationIntegrationType = 0
	DiscordInteractionApplicationIntegrationTypeUserInstall  DiscordInteractionApplicationIntegrationType = 1
)
