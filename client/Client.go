package client

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
)

type Config struct {
	HttpC             *http.Client
	Proxy             string
	ForceV4           bool
	DisableSSLVerify  bool
	ForceAttemptHTTP2 bool
	BaseUrl           string
}

// Setup setup a new client to use telegram api
func (c *Config) Setup() error {
	var proxy func(*http.Request) (*url.URL, error)
	var myDC func(ctx context.Context, network, addr string) (net.Conn, error)
	var TLSClientConfig *tls.Config
	var forceAttemptHTTP2 bool

	if c.Proxy != "" {
		proxyUrl, err := url.Parse(c.Proxy)
		if err != nil {
			return err
		}
		if proxyUrl != nil {
			proxy = http.ProxyURL(proxyUrl)
		} else {
			proxy = nil
		}
	} else {
		proxy = nil
	}

	if c.ForceV4 {
		myDC = func(ctx context.Context, network string, addr string) (net.Conn, error) {
			return (&net.Dialer{}).DialContext(ctx, "tcp4", addr)
		}
	} else {
		myDC = nil
	}

	if c.DisableSSLVerify {
		TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	} else {
		TLSClientConfig = &tls.Config{InsecureSkipVerify: false}
	}

	if c.ForceAttemptHTTP2 {
		forceAttemptHTTP2 = true
	}

	c.HttpC = &http.Client{
		Transport: &http.Transport{
			Proxy:             proxy,
			ForceAttemptHTTP2: forceAttemptHTTP2,
			DialContext:       myDC,
			TLSClientConfig:   TLSClientConfig,
		},
	}

	return nil
}
