package common

type GatewayErrorCode int

const (
	// General errors
	GatewayErrorCodeGeneral GatewayErrorCode = 0

	//1xxxx
	GatewayErrorCodeUnknownAccount                       GatewayErrorCode = 10001
	GatewayErrorCodeUnknownApplication                   GatewayErrorCode = 10002
	GatewayErrorCodeUnknownChannel                       GatewayErrorCode = 10003
	GatewayErrorCodeUnknownGuild                         GatewayErrorCode = 10004
	GatewayErrorCodeUnknownIntegration                   GatewayErrorCode = 10005
	GatewayErrorCodeUnknownInvite                        GatewayErrorCode = 10006
	GatewayErrorCodeUnknownMember                        GatewayErrorCode = 10007
	GatewayErrorCodeUnknownMessage                       GatewayErrorCode = 10008
	GatewayErrorCodeUnknownPermissionOverwrite           GatewayErrorCode = 10009
	GatewayErrorCodeUnknownRole                          GatewayErrorCode = 10010
	GatewayErrorCodeUnknownToken                         GatewayErrorCode = 10011
	GatewayErrorCodeUnknownUser                          GatewayErrorCode = 10012
	GatewayErrorCodeUnknownEmoji                         GatewayErrorCode = 10014
	GatewayErrorCodeUnknownWebhook                       GatewayErrorCode = 10015
	GatewayErrorCodeUnknownWebhookService                GatewayErrorCode = 10016
	GatewayErrorCodeUnknownSession                       GatewayErrorCode = 10020
	GatewayErrorCodeUnknownAsset                         GatewayErrorCode = 10021
	GatewayErrorCodeUnknownBan                           GatewayErrorCode = 10026
	GatewayErrorCodeUnknownSku                           GatewayErrorCode = 10027
	GatewayErrorCodeUnknownStoreListing                  GatewayErrorCode = 10028
	GatewayErrorCodeUnknownEntitlement                   GatewayErrorCode = 10029
	GatewayErrorCodeUnknownBuild                         GatewayErrorCode = 10030
	GatewayErrorCodeUnknownLobby                         GatewayErrorCode = 10031
	GatewayErrorCodeUnknownBranch                        GatewayErrorCode = 10032
	GatewayErrorCodeUnknownRedistributable               GatewayErrorCode = 10036
	GatewayErrorCodeUnknownGiftCode                      GatewayErrorCode = 10038
	GatewayErrorCodeUnknownStream                        GatewayErrorCode = 10049
	GatewayErrorCodeUnknownPremiumSubscriptionOption     GatewayErrorCode = 10050
	GatewayErrorCodeUnknownGuildTemplate                 GatewayErrorCode = 10057
	GatewayErrorCodeUnknownDiscoverableServer            GatewayErrorCode = 10059
	GatewayErrorCodeUnknownStickerPack                   GatewayErrorCode = 10060
	GatewayErrorCodeUnknownSticker                       GatewayErrorCode = 10061
	GatewayErrorCodeUnknownInteraction                   GatewayErrorCode = 10062
	GatewayErrorCodeUnknownApplicationCommand            GatewayErrorCode = 10063
	GatewayErrorCodeUnknownVoiceState                    GatewayErrorCode = 10065
	GatewayErrorCodeUnknownApplicationCommandPermissions GatewayErrorCode = 10066
	GatewayErrorCodeUnknownStageInstance                 GatewayErrorCode = 10067
	GatewayErrorCodeUnknownGuildMemberVerificationForm   GatewayErrorCode = 10068
	GatewayErrorCodeUnknownGuildWelcomeScreen            GatewayErrorCode = 10069
	GatewayErrorCodeUnknownScheduledEvent                GatewayErrorCode = 10070
	GatewayErrorCodeUnknownScheduledEventUser            GatewayErrorCode = 10071
	GatewayErrorCodeUnknownTag                           GatewayErrorCode = 10087
	GatewayErrorCodeUnknownSound                         GatewayErrorCode = 10097
	GatewayErrorCodeUnknownInviteTargetUserJob           GatewayErrorCode = 10124
	GatewayErrorCodeUnknownInviteTargetUsers             GatewayErrorCode = 10129

	//2xxxx
	GatewayErrorCodeBotsCannotUseEndpoint                     GatewayErrorCode = 20001
	GatewayErrorCodeOnlyBotsCanUseEndpoint                    GatewayErrorCode = 20002
	GatewayErrorCodeExplicitContentCannotBeSentToTheRecipient GatewayErrorCode = 20009
	GatewayErrorCodeNotAuthorizedToMessageUser                GatewayErrorCode = 20012
	GatewayErrorCodeActionCannotBePerformedSlowmodeRateLimit  GatewayErrorCode = 20016
	GatewayErrorCodeOnlyOwnerCanPerformAction                 GatewayErrorCode = 20018
	GatewayErrorCodeMessageEditFailedAnnouncementRateLimit    GatewayErrorCode = 20022
	GatewayErrorCodeUnderMinimumAge                           GatewayErrorCode = 20024
	GatewayErrorCodeChannelHasWriteRateLimit                  GatewayErrorCode = 20028
	GatewayErrorCodeServerHasWriteRateLimit                   GatewayErrorCode = 20029
	GatewayErrorCodeContainsDisallowedWords                   GatewayErrorCode = 20031
	GatewayErrorCodePremiumLevelTooLow                        GatewayErrorCode = 20035

	//3xxxx
)
