package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
	"reflect"
	"strconv"
)

func (fb FileBytes) NeedsUpload() bool {
	return true
}

func (fb FileBytes) UploadData() (string, io.Reader, error) {
	return fb.Name, bytes.NewReader(fb.Bytes), nil
}

func (fb FileBytes) SendData() string {
	return "FileBytes must be uploaded"
}

func (fr FileReader) NeedsUpload() bool {
	return true
}

func (fr FileReader) UploadData() (string, io.Reader, error) {
	return fr.Name, fr.Reader, nil
}

func (fr FileReader) SendData() string {
	return "FileReader must be uploaded"
}

func (fp FilePath) NeedsUpload() bool {
	return true
}

func (fp FilePath) UploadData() (string, io.Reader, error) {
	fileHandle, err := os.Open(string(fp))
	if err != nil {
		return "", nil, err
	}

	name := fileHandle.Name()
	return name, fileHandle, err
}

func (fp FilePath) SendData() string {
	return "FilePath must be uploaded"
}

func (fu FileURL) NeedsUpload() bool {
	return false
}

func (fu FileURL) UploadData() (string, io.Reader, error) {
	return "", nil, errors.New("FileURL cannot be uploaded")
}

func (fu FileURL) SendData() string {
	return string(fu)
}

func (fi FileID) NeedsUpload() bool {
	return false
}

func (fi FileID) UploadData() (string, io.Reader, error) {
	return "", nil, errors.New("FileID cannot be uploaded")
}

func (fi FileID) SendData() string {
	return string(fi)
}

func (fa FileAttach) NeedsUpload() bool {
	return false
}

func (fa FileAttach) UploadData() (string, io.Reader, error) {
	return "", nil, errors.New("fileAttach cannot be uploaded")
}

func (fa FileAttach) SendData() string {
	return string(fa)
}

// Error message string.
func (e Error) Error() string {
	return e.Message
}

// IsPrivate returns if the Chat is a private conversation.
func (c Chat) IsPrivate() bool {
	return c.Type == "private"
}

// IsGroup returns if the Chat is a group.
func (c Chat) IsGroup() bool {
	return c.Type == "group"
}

// IsSuperGroup returns if the Chat is a supergroup.
func (c Chat) IsSuperGroup() bool {
	return c.Type == "supergroup"
}

// IsChannel returns if the Chat is a channel.
func (c Chat) IsChannel() bool {
	return c.Type == "channel"
}

// IsCommand returns true if message starts with a "bot_command" entity.
func (m *Message) IsCommand() bool {
	if m.Entities == nil || len(m.Entities) == 0 {
		return false
	}

	entity := m.Entities[0]
	return entity.Offset == 0 && entity.IsCommand()
}

// IsMention returns true if the type of the message entity is "mention" (@username).
func (e MessageEntity) IsMention() bool {
	return e.Type == "mention"
}

// IsHashtag returns true if the type of the message entity is "hashtag".
func (e MessageEntity) IsHashtag() bool {
	return e.Type == "hashtag"
}

// IsCommand returns true if the type of the message entity is "bot_command".
func (e MessageEntity) IsCommand() bool {
	return e.Type == "bot_command"
}

// IsURL returns true if the type of the message entity is "url".
func (e MessageEntity) IsURL() bool {
	return e.Type == "url"
}

// IsEmail returns true if the type of the message entity is "email".
func (e MessageEntity) IsEmail() bool {
	return e.Type == "email"
}

// IsBold returns true if the type of the message entity is "bold" (bold text).
func (e MessageEntity) IsBold() bool {
	return e.Type == "bold"
}

// IsItalic returns true if the type of the message entity is "italic" (italic text).
func (e MessageEntity) IsItalic() bool {
	return e.Type == "italic"
}

// IsCode returns true if the type of the message entity is "code" (monoWidth string).
func (e MessageEntity) IsCode() bool {
	return e.Type == "code"
}

// IsPre returns true if the type of the message entity is "pre" (monoWidth block).
func (e MessageEntity) IsPre() bool {
	return e.Type == "pre"
}

// IsTextLink returns true if the type of the message entity is "text_link" (clickable text URL).
func (e MessageEntity) IsTextLink() bool {
	return e.Type == "text_link"
}

// AddNonEmpty adds a value if it not an empty string.
func (p Params) AddNonEmpty(key, value string) {
	if value != "" {
		p[key] = value
	}
}

// AddNonZero adds a value if it is not zero.
func (p Params) AddNonZero(key string, value int) {
	if value != 0 {
		p[key] = strconv.Itoa(value)
	}
}

// AddNonZero64 is the same as AddNonZero except uses an int64.
func (p Params) AddNonZero64(key string, value int64) {
	if value != 0 {
		p[key] = strconv.FormatInt(value, 10)
	}
}

// AddBool adds a value of a bool if it is true.
func (p Params) AddBool(key string, value bool) {
	if value {
		p[key] = strconv.FormatBool(value)
	}
}

// AddNonZeroFloat adds a floating point value that is not zero.
func (p Params) AddNonZeroFloat(key string, value float64) {
	if value != 0 {
		p[key] = strconv.FormatFloat(value, 'f', 6, 64)
	}
}

// AddInterface adds an interface if it is not nil and can be JSON marshalled.
func (p Params) AddInterface(key string, value interface{}) error {
	if value == nil || (reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil()) {
		return nil
	}

	b, err := json.Marshal(value)
	if err != nil {
		return err
	}

	p[key] = string(b)

	return nil
}

// AddAt adds @ if it not an empty string.
func (p Params) AddAt(value string) {
	if value != "" && value[0:1] != "@" {
		value = "@" + value
	}
}

// AddFirstValid attempts to add the first item that is not a default value. For example, AddFirstValid(0, "", "test") would add "test".
func (p Params) AddFirstValid(key string, args ...interface{}) error {
	for _, arg := range args {
		switch v := arg.(type) {
		case int:
			if v != 0 {
				p[key] = strconv.Itoa(v)
				return nil
			}
		case int64:
			if v != 0 {
				p[key] = strconv.FormatInt(v, 10)
				return nil
			}
		case string:
			if v != "" {
				p[key] = v
				return nil
			}
		case nil:
		default:
			b, err := json.Marshal(arg)
			if err != nil {
				return err
			}

			p[key] = string(b)
			return nil
		}
	}

	return nil
}
