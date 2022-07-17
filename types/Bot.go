package types

import (
	"bytes"
	"net/http"
	"time"
)

// BotApi api config data
type BotApi struct {
	Token             string
	BaseUrl           string
	Debug             bool
	Log               bytes.Buffer
	RequestTimeout    time.Duration
	Client            *http.Client
	SecretToken       string
	GetUpdatesChannel chan interface{}
}
