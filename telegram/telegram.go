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

var Core Api

type Api struct {
	Bot *types.BotApi
}

// BaseUrl set custom api base url.
func (t *Api) BaseUrl(baseUrl string) {
	t.Bot.BaseUrl = baseUrl
}

// Client make new telegram client.
func Client() *client.Config {
	return &client.Config{}
}

// New make new telegram bot api response.
func New(token string) (*Api, error) {
	if token == "" {
		return nil, errors.New("bot token missed")
	}

	c := Client()
	c.BaseUrl = config.DefaultBaseUrl

	err := c.Setup()
	if err != nil {
		return nil, err
	}

	Core.Bot = &types.BotApi{Token: token, BaseUrl: c.BaseUrl, Client: c.HttpC}
	return &Core, nil
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

	err := c.Setup()
	if err != nil {
		return nil, err
	}

	Core.Bot = &types.BotApi{Token: token, BaseUrl: c.BaseUrl, Client: c.HttpC}
	return &Core, nil
}

// NewWithCustomClient make new telegram bot api response with a custom client.
func NewWithCustomClient(token string, c *client.Config) (*Api, error) {
	if token == "" {
		return nil, errors.New("bot token missed")
	}
	if c.BaseUrl == "" {
		c.BaseUrl = config.DefaultBaseUrl
	}

	err := c.Setup()
	if err != nil {
		return nil, err
	}

	Core.Bot = &types.BotApi{Token: token, BaseUrl: c.BaseUrl, Client: c.HttpC}
	return &Core, nil
}

// HandleUpdate parses and return update received via webhook
func HandleUpdate(r *http.Request) (*types.Update, error) {
	var update types.Update

	if r.Method != http.MethodPost {
		return nil, errors.New("wrong HTTP method required POST")
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		return nil, err
	}

	return &update, nil
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
