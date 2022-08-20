package telegram

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/raminsa/telegram-bot-api/config"
	"github.com/raminsa/telegram-bot-api/types"
)

// SetSecretToken parse secret token for very webhook request
func (t *Api) SetSecretToken(secretToken string) {
	telegram.Bot.SecretToken = secretToken
}

func (t *Api) WriteDebugLog(msg string) {
	logger := log.New(&t.Bot.Log, "Debug: ", log.LstdFlags|log.Llongfile)
	Info := func(info string) {
		logger.Output(2, info)
	}
	Info(msg)
}

// GetLoggerFile get debug log
func (t *Api) GetLoggerFile() string {
	return t.Bot.Log.String()
}

// WriteLoggerFile write debug to file
func (t *Api) WriteLoggerFile(fileName string) error {
	return ioutil.WriteFile(fileName, t.Bot.Log.Bytes(), 0644)
}

// MakeRequest makes a request to a specific endpoint with our token.
func (t *Api) MakeRequest(endpoint string, params types.Params) (*types.APIResponse, error) {
	if t.Bot.Debug {
		t.WriteDebugLog(fmt.Sprintf("Endpoint: %s, params: %v\n", endpoint, params))
	}

	URL := fmt.Sprintf(t.Bot.BaseUrl+config.APIEndpoint, t.Bot.Token, endpoint)

	values := buildParams(params)

	var timeout time.Duration
	if t.Bot.RequestTimeout == 0 {
		timeout = 2 * time.Minute
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, URL, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if t.Bot.SecretToken != "" {
		req.Header.Set("X-Telegram-Bot-Api-Secret-Token", t.Bot.SecretToken)
	}

	resp, err := t.Bot.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp types.APIResponse
	bytes, err := t.decodeAPIResponse(resp.Body, &apiResp)
	if err != nil {
		return &apiResp, err
	}

	if t.Bot.Debug {
		t.WriteDebugLog(fmt.Sprintf("Endpoint: %s, response: %s\n", endpoint, string(bytes)))
	}

	if !apiResp.Ok {
		var parameters types.ResponseParameters
		if apiResp.Parameters != nil {
			parameters = *apiResp.Parameters
		}
		return &apiResp, &types.Error{
			Code:               apiResp.ErrorCode,
			Message:            apiResp.Description,
			ResponseParameters: parameters,
		}
	}

	return &apiResp, nil
}

// UploadFiles makes a request to the API with files.
func (t *Api) UploadFiles(endpoint string, params types.Params, files []types.RequestFile) (*types.APIResponse, error) {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)

	go func() {
		defer w.Close()
		defer m.Close()

		for field, value := range params {
			if err := m.WriteField(field, value); err != nil {
				w.CloseWithError(err)
				return
			}
		}

		for _, file := range files {
			if file.Data.NeedsUpload() {
				name, reader, err := file.Data.UploadData()
				if err != nil {
					w.CloseWithError(err)
					return
				}
				if file.FileName != "" {
					name = file.FileName
				}

				part, err := m.CreateFormFile(file.Name, name)
				if err != nil {
					w.CloseWithError(err)
					return
				}

				if _, err := io.Copy(part, reader); err != nil {
					w.CloseWithError(err)
					return
				}

				if closer, ok := reader.(io.ReadCloser); ok {
					if err = closer.Close(); err != nil {
						w.CloseWithError(err)
						return
					}
				}
			} else {
				value := file.Data.SendData()

				if err := m.WriteField(file.Name, value); err != nil {
					w.CloseWithError(err)
					return
				}
			}
		}
	}()

	if t.Bot.Debug {
		t.WriteDebugLog(fmt.Sprintf("Endpoint: %s, params: %v, with %d files\n", endpoint, params, len(files)))
	}

	URL := fmt.Sprintf(t.Bot.BaseUrl+config.APIEndpoint, t.Bot.Token, endpoint)

	var timeout time.Duration
	if t.Bot.RequestTimeout == 0 {
		timeout = 2 * time.Minute
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, URL, r)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", m.FormDataContentType())
	if t.Bot.SecretToken != "" {
		req.Header.Set("X-Telegram-Bot-Api-Secret-Token", t.Bot.SecretToken)
	}

	resp, err := t.Bot.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp types.APIResponse
	bytes, err := t.decodeAPIResponse(resp.Body, &apiResp)
	if err != nil {
		return &apiResp, err
	}

	if t.Bot.Debug {
		t.WriteDebugLog(fmt.Sprintf("Endpoint: %s, response: %s\n", endpoint, string(bytes)))
	}

	if !apiResp.Ok {
		var parameters types.ResponseParameters
		if apiResp.Parameters != nil {
			parameters = *apiResp.Parameters
		}
		return &apiResp, &types.Error{
			Code:               apiResp.ErrorCode,
			Message:            apiResp.Description,
			ResponseParameters: parameters,
		}
	}

	return &apiResp, nil
}

// Request sends a Chattable to Telegram, and returns the APIResponse.
func (t *Api) Request(c types.Chattable) (*types.APIResponse, error) {
	params, err := c.Params()
	if err != nil {
		return nil, err
	}

	if f, ok := c.(types.Fileable); ok {
		files := f.Files()

		// If we have files that need to be uploaded, we should delegate the
		// request to UploadFile.
		if hasFilesNeedingUpload(files) {
			return t.UploadFiles(f.EndPoint(), params, files)
		}

		// However, if there are no files to be uploaded, there's likely things
		// that need to be turned into params instead.
		for _, file := range files {
			params[file.Name] = file.Data.SendData()
		}
	}

	return t.MakeRequest(c.EndPoint(), params)
}

// Send will send a Chattable item to Telegram and provides the returned Message.
func (t *Api) Send(c types.Chattable) (*types.Message, error) {
	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var message types.Message
	err = json.Unmarshal(resp.Result, &message)

	return &message, err
}

// decodeAPIResponse decode response and return slice of bytes if debug enabled.
func (t *Api) decodeAPIResponse(responseBody io.Reader, resp *types.APIResponse) ([]byte, error) {
	if !t.Bot.Debug {
		dec := json.NewDecoder(responseBody)
		err := dec.Decode(resp)
		return nil, err
	}

	data, err := ioutil.ReadAll(responseBody)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func buildParams(in types.Params) url.Values {
	if in == nil {
		return url.Values{}
	}

	out := url.Values{}

	for key, value := range in {
		out.Set(key, value)
	}

	return out
}

func hasFilesNeedingUpload(files []types.RequestFile) bool {
	for _, file := range files {
		if file.Data.NeedsUpload() {
			return true
		}
	}

	return false
}
