package connection

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-discord-wrapper/functions"
	"go-discord-wrapper/types"
	"go-discord-wrapper/util"
	"net/http"
	"net/url"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"
)

type DiscordClient struct {
	Token *string

	APIVersion *types.DiscordAPIVersion

	Logger *zerolog.Logger

	Intents *types.DiscordIntent

	Websocket *websocket.Conn

	Events map[types.DiscordEventType]func(session *DiscordClient, event types.DiscordEvent)

	mu sync.RWMutex

	LastEventNum *int

	ReconnectURL *string

	SessionID *string
}

func NewDiscordClient(token string, intents types.DiscordIntent) *DiscordClient {
	return &DiscordClient{
		Token:      &token,
		APIVersion: functions.PointerTo(types.DiscordAPIVersion10),
		Logger:     util.NewLogger(),
		Intents:    &intents,
	}
}

func (d *DiscordClient) initializeGatewayConnection() (*types.DiscordBotRegisterResponse, error) {
	do, err := http.DefaultClient.Do(&http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme: "https",
			Host:   "discord.com",
			Path:   types.DiscordAPIBaseString(*d.APIVersion) + types.DiscordAPIGatewayRequest,
		},
		Header: http.Header{
			"Authorization": []string{"Bot " + *d.Token},
		},
	})
	if err != nil {
		return nil, err
	}

	if do.StatusCode != http.StatusOK {
		return nil, errors.New("failed to register bot gateway connection, status code: " + do.Status)
	}

	defer func() {
		_ = do.Body.Close()
	}()

	var resp types.DiscordBotRegisterResponse
	if err := json.NewDecoder(do.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (d *DiscordClient) Login() error {
	gatewayResp, err := d.initializeGatewayConnection()
	if err != nil {
		return err
	}

	d.Logger.Info().Msgf("Connecting to gateway websocket at %s with %d shards", gatewayResp.Url, gatewayResp.Shards)

	if err := d.connectWebsocket(gatewayResp.Url); err != nil {
		return err
	}

	go func() {
		if err := d.listenWebsocket(); err != nil {
			d.Logger.Err(err).Msg("Error listening to websocket")
		}
	}()

	return nil
}

func OnEvent[T types.DiscordEvent](d *DiscordClient, event types.DiscordEventType, handler func(*DiscordClient, T)) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.Events == nil {
		d.Events = make(map[types.DiscordEventType]func(session *DiscordClient, event types.DiscordEvent))
	}

	d.Events[event] = func(
		session *DiscordClient,
		ev types.DiscordEvent,
	) {
		typed, ok := ev.(T)
		if !ok {
			session.Logger.Warn().
				Str("expected", fmt.Sprintf("%T", *new(T))).
				Str("got", fmt.Sprintf("%T", ev)).
				Msg("event type mismatch")
			return
		}
		handler(session, typed)
	}
}

func (d *DiscordClient) dispatch(event types.DiscordEventType, payload types.Payload) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if d.Events == nil {
		return
	}

	d.LastEventNum = payload.S

	rawEvent := d.Events[event]
	if rawEvent != nil {
		discordEvent, err := d.convertToEvent(event, payload.D)
		if err != nil {
			d.Logger.Err(err).Msgf("Failed to convert event %s", event)
			return
		}

		go rawEvent(d, discordEvent)
	}
}

func (d *DiscordClient) convertToEvent(event types.DiscordEventType, data json.RawMessage) (types.DiscordEvent, error) {
	switch event {
	case "MESSAGE_CREATE":
		var msgCreateEvent types.DiscordMessageCreateEvent
		if err := json.Unmarshal(data, &msgCreateEvent); err != nil {
			return nil, err
		}
		return &msgCreateEvent, nil
	case "READY":
		var readyEvent types.DiscordReadyEvent
		if err := json.Unmarshal(data, &readyEvent); err != nil {
			return nil, err
		}
		return &readyEvent, nil
	default:
		return nil, errors.New("unsupported event type: " + string(event))
	}
}

func (d *DiscordClient) Shutdown() error {
	if d.Websocket != nil {
		return d.Websocket.Close()
	}
	return nil
}
