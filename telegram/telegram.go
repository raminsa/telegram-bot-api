package telegram

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/raminsa/telegram-bot-api/client"
	"github.com/raminsa/telegram-bot-api/config"
	"github.com/raminsa/telegram-bot-api/types"
)

type Api struct {
	Bot types.BotApi
}

var Core Api

// BaseUrl set custom api base url.
func (t *Api) BaseUrl(baseUrl string) {
	t.Bot.BaseUrl = baseUrl
}

// Client make new telegram client.
func Client() *client.Client {
	return &client.Client{}
}

// New make new telegram bot api response.
func New(token string) (Core Api, err error) {
	if token == "" {
		err = errors.New("bot token missed")
		return
	}

	c := Client()
	c.BaseUrl = config.DefaultBaseUrl

	var makeClient *http.Client
	makeClient, err = c.MakeClient()
	if err != nil {
		return
	}

	Core.Bot = types.BotApi{Token: token, BaseUrl: c.BaseUrl, Client: makeClient}

	return
}

// NewWithBaseUrl make new telegram bot api response with custom base url.
func NewWithBaseUrl(token, baseUrl string) (Core Api, err error) {
	if token == "" {
		err = errors.New("bot token missed")
		return
	}
	if baseUrl == "" {
		err = errors.New("base url missed")
		return
	}

	c := Client()
	c.BaseUrl = baseUrl

	var makeClient *http.Client
	makeClient, err = c.MakeClient()
	if err != nil {
		return
	}

	Core.Bot = types.BotApi{Token: token, BaseUrl: c.BaseUrl, Client: makeClient}

	return
}

// NewWithCustomClient make new telegram bot api response with a custom client.
func NewWithCustomClient(token string, Client *client.Client) (Core Api, err error) {
	if token == "" {
		err = errors.New("bot token missed")
		return
	}
	if Client.BaseUrl == "" {
		Client.BaseUrl = config.DefaultBaseUrl
	}

	var makeClient *http.Client
	makeClient, err = Client.MakeClient()
	if err != nil {
		return
	}

	Core.Bot = types.BotApi{Token: token, BaseUrl: Client.BaseUrl, Client: makeClient}

	return
}

// HandleUpdate parses and return update received via webhook
func HandleUpdate(r *http.Request) (types.Update, error) {
	var update types.Update

	if r.Method != http.MethodPost {
		return update, errors.New("wrong HTTP method required POST")
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		return update, err
	}

	return update, nil
}

// HandleUpdateError response writer error to requested server
func HandleUpdateError(w http.ResponseWriter, wErr error) error {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	errMsg, err := json.Marshal(map[string]string{
		"error": wErr.Error(),
	})
	if err != nil {
		return err
	}
	if _, err = w.Write(errMsg); err != nil {
		return err
	}

	return nil
}
