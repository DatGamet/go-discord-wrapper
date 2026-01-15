package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

type DiscordEntitlement struct {
	ID            DiscordSnowflake       `json:"id"`
	SkuID         DiscordSnowflake       `json:"sku_id"`
	ApplicationID DiscordSnowflake       `json:"application_id"`
	UserID        *DiscordSnowflake      `json:"user_id"`
	Type          DiscordEntitlementType `json:"type"`
	Deleted       bool                   `json:"deleted"`
	StartsAt      *time.Time             `json:"starts_at,omitempty"`
	EndsAt        *time.Time             `json:"ends_at,omitempty"`
	GuildID       *DiscordSnowflake      `json:"guild_id,omitempty"`
	Consumed      bool                   `json:"consumed,omitempty"`
}

type DiscordEntitlementType int

const (
	DiscordEntitlementTypePurchase                DiscordEntitlementType = 1
	DiscordEntitlementTypePremiumSubscription     DiscordEntitlementType = 2
	DiscordEntitlementTypeDeveloperGift           DiscordEntitlementType = 3
	DiscordEntitlementTypeTestModePurchase        DiscordEntitlementType = 4
	DiscordEntitlementTypeFreePurchase            DiscordEntitlementType = 5
	DiscordEntitlementTypeUserGift                DiscordEntitlementType = 6
	DiscordEntitlementTypePremiumPurchase         DiscordEntitlementType = 7
	DiscordEntitlementTypeApplicationSubscription DiscordEntitlementType = 8
)

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

type DiscordMessageMessageReference struct {
	Type            *DiscordMessageMessageReferenceType `json:"type,omitempty"`
	MessageID       *DiscordSnowflake                   `json:"message_id,omitempty"`
	ChannelID       *DiscordSnowflake                   `json:"channel_id,omitempty"`
	GuildID         *DiscordSnowflake                   `json:"guild_id,omitempty"`
	FailIfNotExists *bool                               `json:"fail_if_not_exists,omitempty"`
}

type DiscordMessageInteractionMetadataApplicationCommand struct {
	ID                           DiscordSnowflake                                             `json:"id"`
	Type                         DiscordInteractionType                                       `json:"type"`
	User                         DiscordUser                                                  `json:"user,omitempty"`
	AuthorizingIntegrationOwners map[DiscordInteractionApplicationIntegrationType]interface{} `json:"authorizing_integration_owners,omitempty"`
	OriginalResponseMessageID    *DiscordSnowflake                                            `json:"original_response_message_id,omitempty"`
	TargetUser                   *DiscordUser                                                 `json:"target_user,omitempty"`
	TargetMessageID              *DiscordSnowflake                                            `json:"target_message_id,omitempty"`
}

type DiscordMessageInteractionMetadataMessageComponent struct {
	ID                           DiscordSnowflake                                             `json:"id"`
	Type                         DiscordInteractionType                                       `json:"type"`
	User                         DiscordUser                                                  `json:"user,omitempty"`
	AuthorizingIntegrationOwners map[DiscordInteractionApplicationIntegrationType]interface{} `json:"authorizing_integration_owners,omitempty"`
	OriginalResponseMessageID    *DiscordSnowflake                                            `json:"original_response_message_id,omitempty"`
	InteractedMessageID          *DiscordSnowflake                                            `json:"interacted_message_id,omitempty"`
}

type AnyDiscordMessageInteractionMetadata interface{}

type DiscordMessageInteractionMetadata struct {
	Value AnyDiscordMessageInteractionMetadata
}

type AnyDiscordMessageInteractionMetadataModalSubmitTriggeringInteractionMetadata interface{}

type DiscordMessageInteractionMetadataModalSubmitTriggering struct {
	AnyDiscordMessageInteractionMetadataModalSubmitTriggeringInteractionMetadata
}

func (d *DiscordMessageInteractionMetadataModalSubmitTriggering) UnmarshalJSON(data []byte) error {
	var a DiscordMessageInteractionMetadataApplicationCommand
	if err := json.Unmarshal(data, &a); err == nil && a.ID != "" {
		d.AnyDiscordMessageInteractionMetadataModalSubmitTriggeringInteractionMetadata = &a
		return nil
	}

	var b DiscordMessageInteractionMetadataMessageComponent
	if err := json.Unmarshal(data, &b); err == nil && b.ID != "" {
		d.AnyDiscordMessageInteractionMetadataModalSubmitTriggeringInteractionMetadata = &b
		return nil
	}

	return nil
}

func (d *DiscordMessageInteractionMetadata) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		d.Value = nil
		return nil
	}

	var a DiscordMessageInteractionMetadataApplicationCommand
	if err := json.Unmarshal(data, &a); err == nil && a.ID != "" {
		d.Value = &a
		return nil
	}

	var b DiscordMessageInteractionMetadataMessageComponent
	if err := json.Unmarshal(data, &b); err == nil && b.ID != "" {
		d.Value = &b
		return nil
	}

	var c DiscordMessageInteractionMetadataModalSubmit
	if err := json.Unmarshal(data, &c); err == nil && c.ID != "" {
		d.Value = &c
		return nil
	}

	return fmt.Errorf("unknown DiscordMessageInteractionMetadata: %s", string(data))
}

type DiscordMessageInteractionMetadataModalSubmit struct {
	ID                            DiscordSnowflake                                             `json:"id"`
	Type                          DiscordInteractionType                                       `json:"type"`
	User                          DiscordUser                                                  `json:"user,omitempty"`
	AuthorizingIntegrationOwners  map[DiscordInteractionApplicationIntegrationType]interface{} `json:"authorizing_integration_owners,omitempty"`
	OriginalResponseMessageID     *DiscordSnowflake                                            `json:"original_response_message_id,omitempty"`
	TriggeringInteractionMetadata *DiscordMessageInteractionMetadataModalSubmitTriggering      `json:"triggering_interaction_metadata,omitempty"`
}

type InteractionContextType int

const (
	InteractionContextTypeGuild          InteractionContextType = 0
	InteractionContextTypeBotDM          InteractionContextType = 1
	InteractionContextTypePrivateChannel InteractionContextType = 2
)

type DiscordInteraction struct {
	ID                           DiscordSnowflake                                             `json:"id"`
	ApplicationID                DiscordSnowflake                                             `json:"application_id"`
	Type                         DiscordInteractionType                                       `json:"type"`
	Data                         DiscordInteractionData                                       `json:"data,omitempty"`
	GuildID                      *DiscordSnowflake                                            `json:"guild_id,omitempty"`
	ChannelID                    *DiscordSnowflake                                            `json:"channel_id,omitempty"`
	Guild                        *Guild                                                       `json:"guild,omitempty"`
	Channel                      *DiscordChannel                                              `json:"channel,omitempty"`
	Member                       *GuildMember                                                 `json:"member,omitempty"`
	User                         *DiscordUser                                                 `json:"user,omitempty"`
	Token                        string                                                       `json:"token"`
	Version                      int                                                          `json:"version"`
	Message                      *DiscordMessage                                              `json:"message,omitempty"`
	AppPermissions               string                                                       `json:"app_permissions,omitempty"`
	Locale                       *DiscordLocale                                               `json:"locale,omitempty"`
	GuildLocale                  string                                                       `json:"guild_locale,omitempty"`
	Entitlements                 []DiscordEntitlement                                         `json:"entitlements,omitempty"`
	AuthorizingIntegrationOwners map[DiscordInteractionApplicationIntegrationType]interface{} `json:"authorizing_integration_owners,omitempty"`
	Context                      InteractionContextType                                       `json:"context,omitempty"`
	AttachmentSizeLimit          int                                                          `json:"attachment_size_limit,omitempty"`
}

func (i *DiscordInteraction) UnmarshalJSON(data []byte) error {
	type Alias DiscordInteraction
	aux := &struct {
		Data json.RawMessage `json:"data"`
		*Alias
	}{
		Alias: (*Alias)(i),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	if aux.Data == nil {
		return nil
	}

	// Peek into data.type
	var typeProbe struct {
		Type DiscordInteractionDataApplicationCommandType `json:"type"`
	}

	if err := json.Unmarshal(aux.Data, &typeProbe); err != nil {
		return err
	}

	switch typeProbe.Type {
	case DiscordInteractionDataApplicationCommandTypeChatInput:
		var cmd DiscordInteractionDataApplicationCommand
		if err := json.Unmarshal(aux.Data, &cmd); err != nil {
			return err
		}
		i.Data = &cmd

	default:
		return fmt.Errorf("unknown interaction data type %d", typeProbe.Type)
	}

	return nil
}

type DiscordInteractionDataType int

const (
	DiscordInteractionDataTypePing                           DiscordInteractionDataType = 1
	DiscordInteractionDataTypeApplicationCommand             DiscordInteractionDataType = 2
	DiscordInteractionDataTypeMessageComponent               DiscordInteractionDataType = 3
	DiscordInteractionDataTypeApplicationCommandAutocomplete DiscordInteractionDataType = 4
	DiscordInteractionDataTypeModalSubmit                    DiscordInteractionDataType = 5
)

type DiscordInteractionData interface {
	GetType() DiscordInteractionDataType
}

type DiscordInteractionDataApplicationCommandType int

const (
	DiscordInteractionDataApplicationCommandTypeChatInput       DiscordInteractionDataApplicationCommandType = 1
	DiscordInteractionDataApplicationCommandTypeUser            DiscordInteractionDataApplicationCommandType = 2
	DiscordInteractionDataApplicationCommandTypeMessage         DiscordInteractionDataApplicationCommandType = 3
	DiscordInteractionDataApplicationCommandTypePrimaryEndpoint DiscordInteractionDataApplicationCommandType = 4
)

type DiscordInteractionDataApplicationCommand struct {
	ID          DiscordSnowflake                             `json:"id"`
	CommandName string                                       `json:"name"`
	Type        DiscordInteractionDataApplicationCommandType `json:"type"`
	GuildID     *DiscordSnowflake                            `json:"guild_id,omitempty"`
	TargetID    *DiscordSnowflake                            `json:"target_id,omitempty"`
	Resolved    *DiscordResolvedData                         `json:"resolved,omitempty"`
	Options     []AnyDiscordMessageInteractionMetadata       `json:"options,omitempty"`
}

func (d *DiscordInteractionDataApplicationCommand) GetType() DiscordInteractionDataType {
	return DiscordInteractionDataTypeApplicationCommand
}

func (d *DiscordInteractionDataApplicationCommand) UnmarshalJSON(data []byte) error {
	type Alias DiscordInteractionDataApplicationCommand
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(d),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMap); err != nil {
		return err
	}

	if optionsData, ok := rawMap["options"]; ok {
		var rawOptions []json.RawMessage
		if err := json.Unmarshal(optionsData, &rawOptions); err != nil {
			return err
		}

		for _, rawOption := range rawOptions {
			var option AnyDiscordMessageInteractionMetadata
			if err := json.Unmarshal(rawOption, &option); err != nil {
				return err
			}
			d.Options = append(d.Options, option)
		}
	}

	return nil
}
