package telegram

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Raminsa/Telegram_API/client"
	"github.com/Raminsa/Telegram_API/config"
	"github.com/Raminsa/Telegram_API/types"
)

type Api struct {
	Bot types.BotApi
}

var telegram Api

// BaseUrl set custom api base url.
func (t *Api) BaseUrl(baseUrl string) {
	t.Bot.BaseUrl = baseUrl
}

// Client make new telegram client.
func Client() client.Client {
	return client.Client{}
}

// New make new telegram bot api response.
func New(token string) (*Api, error) {
	if token == "" {
		return nil, errors.New("bot token missed")
	}
	c := Client()
	c.BaseUrl = config.DefaultBaseUrl

	makeClient, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	telegram.Bot = types.BotApi{Token: token, BaseUrl: c.BaseUrl, Client: makeClient}

	return &telegram, nil
}

// NewWithBaseUrl make new telegram bot api response with custom base url.
func NewWithBaseUrl(token, baseUrl string) (*Api, error) {
	if token == "" {
		return nil, errors.New("bot token missed")
	}
	if baseUrl == "" {
		return nil, errors.New("base url missed")
	}

	c := Client()
	c.BaseUrl = baseUrl

	makeClient, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	telegram.Bot = types.BotApi{Token: token, BaseUrl: c.BaseUrl, Client: makeClient}

	return &telegram, nil
}

// NewWithCustomClient make new telegram bot api response with custom client.
func NewWithCustomClient(token string, Client *client.Client) (*Api, error) {
	if token == "" {
		return nil, errors.New("bot token missed")
	}
	if Client.BaseUrl == "" {
		Client.BaseUrl = config.DefaultBaseUrl
	}

	makeClient, err := Client.MakeClient()
	if err != nil {
		return nil, err
	}

	telegram.Bot = types.BotApi{Token: token, BaseUrl: Client.BaseUrl, Client: makeClient}

	return &telegram, nil
}

// HandleUpdate parses and returns update received via webhook
func HandleUpdate(r *http.Request) (*types.Update, error) {
	if r.Method != http.MethodPost {
		return nil, errors.New("wrong HTTP method required POST")
	}

	defer r.Body.Close()

	var update types.Update
	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		return nil, err
	}

	return &update, nil
}

// HandleUpdateError response writer error to requested server
func HandleUpdateError(w http.ResponseWriter, err error) {
	errMsg, _ := json.Marshal(map[string]string{"error": err.Error()})
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(errMsg)
}
