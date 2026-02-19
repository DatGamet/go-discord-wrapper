package events

var EventFactories = map[EventType]func() Event{
	EventMessageCreate:     MessageCreateEvent{}.DesiredEventType,
	EventReady:             ReadyEvent{}.DesiredEventType,
	EventGuildCreate:       GuildCreateEvent{}.DesiredEventType,
	EventInteractionCreate: InteractionCreateEvent{}.DesiredEventType,
	EventGuildDelete:       GuildDeleteEvent{}.DesiredEventType,
	EventInviteCreate:      InviteCreateEvent{}.DesiredEventType,
	EventInviteDelete:      InviteDeleteEvent{}.DesiredEventType,
	EventChannelCreate:     ChannelCreateEvent{}.DesiredEventType,
	EventChannelDelete:     ChannelDeleteEvent{}.DesiredEventType,
}

type Event interface {
	Event() EventType
	DesiredEventType() Event
}

type EventType string

const (
	EventMessageCreate     EventType = "MESSAGE_CREATE"
	EventReady             EventType = "READY"
	EventGuildCreate       EventType = "GUILD_CREATE"
	EventInteractionCreate EventType = "INTERACTION_CREATE"
	EventGuildDelete       EventType = "GUILD_DELETE"
	/*TODO
	MessageDelete
	MessageUpdate

	GuildAuditLogEntryCreate
	*/
	EventChannelCreate EventType = "CHANNEL_CREATE"
	// ChannelUpdate     EventType = "CHANNEL_UPDATE"
	EventChannelDelete EventType = "CHANNEL_DELETE"
	/*
		ChannelPinsUpdate EventType = "CHANNEL_PINS_UPDATE"


			EventChannelDelete
			ChannelPinsUpdate

			RoleCreate
			RoleUpdate
			RoleDelete

			WebhookUpdate

			IntegrationCreate
			IntegrationUpdate
			IntegrationDelete

			AutoModerationRuleCreate
			AutoModerationRuleUpdate
			AutoModerationRuleDelete
			AutoModerationActionExecute

			ThreadCreate
			ThreadUpdate
			ThreadDelete
			ThreadMemberUpdate
			ThreadMembersUpdate

			EntitlementCreate
			EntitlementUpdate
			EntitlementDelete

			GuildBanAdd
			GuildBanRemove
			GuildEmojisUpdate
			GuildStickersUpdate
			GuildIntegrationsUpdate
			GuildMemberAdd
			GuildMemberRemove

			ScheduledEventCreate
			ScheduledEventUpdate
			ScheduledEventDelete
			ScheduledEventUserAdd
			ScheduledEventUserRemove

			SoundboardSoundsCreate
			SoundboardSoundsUpdate
			SoundboardSoundsDelete

	*/
	EventInviteCreate EventType = "INVITE_CREATE"
	EventInviteDelete EventType = "INVITE_DELETE"

	/*

		MessageReactionAdd
		MessageReactionRemove
		MessageReactionRemoveAll
		MessageReactionRemoveEmoji

		PresenceUpdate

		StageInstanceUpdate
		StageInstanceCreate
		StageInstanceDelete

		SubscriptionCreate
		SubscriptionDelete
		SubscriptionUpdate

		TypingStart

		UserUpdate

		VoiceStateUpdate

		MessagePollVoteAdd
		MessagePollVoteRemove
	*/
)
