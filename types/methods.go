package types

import (
	"fmt"
	"net/url"

	"github.com/raminsa/telegram-bot-api/config"
)

// SetWebhook Specify a URL and receive incoming updates via an outgoing webhook. Whenever there is an update for the bot, we will send an HTTPS POST request to the specified URL, containing a JSON-serialized Update. In case of an unsuccessful request, we will give up after a reasonable amount of attempts. Returns True on success. If you'd like to make sure that the webhook was set by you, you can specify secret data in the parameter secret_token. If specified, the request will contain a header “X-Telegram-Bot-Api-Secret-Token” with the secret token as content.
// Notes:
//1. You will not be able to receive updates using getUpdates for as long as an outgoing webhook is set up.
//2. To use a self-signed certificate, you need to upload your public key certificate using certificate parameter. Please upload as InputFile, sending a String will not work.
//3. Ports currently supported for webhooks: 443, 80, 88, 8443.
//If you're having any trouble setting up webhooks, please check out this amazing guide to webhooks.

// GetUpdates Receive incoming updates using long polling (wiki). An Array of Update objects is returned.
type GetUpdates struct {
	Offset         int
	Limit          int
	Timeout        int
	AllowedUpdates []string
}

func (s *GetUpdates) Params() (Params, error) {
	params := make(Params, 4)

	params.AddNonZero("offset", s.Offset)
	params.AddNonZero("limit", s.Limit)
	params.AddNonZero("timeout", s.Timeout)
	err := params.AddInterface("allowed_updates", s.AllowedUpdates)

	return params, err
}
func (*GetUpdates) EndPoint() string {
	return config.EndpointGetUpdates
}

// UpdatesChannel is the channel for getting updates.
type UpdatesChannel <-chan Update

// Clear discards all unprocessed incoming updates.
func (ch UpdatesChannel) Clear() {
	for len(ch) != 0 {
		<-ch
	}
}

type SetWebhook struct {
	URL                *url.URL        // HTTPS URL to send updates to. Use an empty string to remove webhook integration
	Certificate        RequestFileData // Optional. Upload your public key certificate so that the root certificate in use can be checked. See our self-signed guide for details.
	IPAddress          string          // Optional. The fixed IP address which will be used to send webhook requests instead of the IP address resolved through DNS
	MaxConnections     int             // Optional. The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100. Defaults to 40. Use lower values to limit the load on your bot's server, and higher values to increase your bot's throughput.
	AllowedUpdates     []string        // Optional. A JSON-serialized list of the update types you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all update types except chat_member (default). If not specified, the previous setting will be used. Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted updates may be received for a short period of time.
	DropPendingUpdates bool            // Optional. Pass True to drop all pending updates
}

func (s *SetWebhook) Params() (Params, error) {
	params := make(Params, 5)

	if s.URL != nil {
		params["url"] = s.URL.String()
	}
	params.AddNonEmpty("ip_address", s.IPAddress)
	params.AddNonZero("max_connections", s.MaxConnections)
	err := params.AddInterface("allowed_updates", s.AllowedUpdates)
	if err != nil {
		return params, err
	}
	params.AddBool("drop_pending_updates", s.DropPendingUpdates)

	return params, nil
}
func (s *SetWebhook) Files() []RequestFile {
	if s.Certificate != nil {
		return []RequestFile{{
			Name: "certificate",
			Data: s.Certificate,
		}}
	}

	return nil
}
func (s *SetWebhook) EndPoint() string {
	return config.EndpointSetWebhook
}

// DeleteWebhook Remove webhook integration if you decide to switch back to getUpdates. Returns True on success.
type DeleteWebhook struct {
	DropPendingUpdates bool // Optional. Pass True to drop all pending updates
}

func (s *DeleteWebhook) Params() (Params, error) {
	params := make(Params, 1)

	params.AddBool("drop_pending_updates", s.DropPendingUpdates)

	return params, nil
}
func (s *DeleteWebhook) EndPoint() string {
	return config.EndpointDeleteWebhook
}

// SendMessage Send text messages. On success, the sent Message is returned.
type SendMessage struct {
	ChatID                   int64  // required. use for user|channel as int
	ChatIDStr                string // required. use for user|channel as string
	Username                 string // required. use for channel
	Text                     string // required
	ParseMode                string
	Entities                 []MessageEntity
	DisableWebPagePreview    bool
	DisableNotification      bool
	ProtectContent           bool
	ReplyToMessageID         int
	AllowSendingWithoutReply bool
	ReplyMarkup              interface{}
}

func (s *SendMessage) Params() (Params, error) {
	params := make(Params, 10)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddNonEmpty("text", s.Text)
	params.AddNonEmpty("parse_mode", s.ParseMode)
	err = params.AddInterface("entities", s.Entities)
	if err != nil {
		return params, err
	}
	params.AddBool("disable_web_page_preview", s.DisableWebPagePreview)
	params.AddBool("disable_notification", s.DisableNotification)
	params.AddBool("protect_content", s.ProtectContent)
	params.AddNonZero("reply_to_message_id", s.ReplyToMessageID)
	params.AddBool("allow_sending_without_reply", s.AllowSendingWithoutReply)
	err = params.AddInterface("reply_markup", s.ReplyMarkup)

	return params, err
}
func (s *SendMessage) EndPoint() string {
	return config.EndpointSendMessage
}

// ForwardMessage Forward messages of any kind. Service messages can't be forwarded. On success, the sent Message is returned.
type ForwardMessage struct {
	ChatID              int64  // required. use for user|channel as int
	ChatIDStr           string // required. use for user|channel as string
	Username            string // required. use for channel
	FromChatID          int64  // required. use for user|channel as int
	FromChatIDStr       string // required. use for user|channel as string
	FromUsername        string // required. use for channel
	DisableNotification bool
	ProtectContent      bool
	MessageID           int // required
}

func (s *ForwardMessage) Params() (Params, error) {
	params := make(Params, 5)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	err = params.AddFirstValid("from_chat_id", s.FromChatID, s.FromChatIDStr, s.FromUsername)
	if err != nil {
		return params, err
	}
	params.AddBool("disable_notification", s.DisableNotification)
	params.AddBool("protect_content", s.ProtectContent)
	params.AddNonZero("message_id", s.MessageID)

	return params, nil
}
func (s *ForwardMessage) EndPoint() string {
	return config.EndpointForwardMessage
}

// CopyMessage Copy messages of any kind. Service messages and invoice messages can't be copied. The method is analogous to the method forwardMessage, but the copied message doesn't have a link to the original message. Returns the MessageId of the sent message on success.
type CopyMessage struct {
	ChatID                   int64  // required. use for user|channel as int
	ChatIDStr                string // required. use for user|channel as string
	Username                 string // required. use for channel
	FromChatID               int64  // required. use for user|channel as int
	FromChatIDStr            string // required. use for user|channel as string
	FromUsername             string // required. use for channel
	MessageID                int
	Caption                  string
	ParseMode                string
	CaptionEntities          []MessageEntity
	DisableNotification      bool
	ProtectContent           bool
	ReplyToMessageID         int
	AllowSendingWithoutReply bool
	ReplyMarkup              interface{}
}

func (s *CopyMessage) Params() (Params, error) {
	params := make(Params, 11)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	err = params.AddFirstValid("from_chat_id", s.FromChatID, s.FromChatIDStr, s.FromUsername)
	if err != nil {
		return params, err
	}
	params.AddNonZero("message_id", s.MessageID)
	params.AddNonEmpty("caption", s.Caption)
	params.AddNonEmpty("parse_mode", s.ParseMode)
	err = params.AddInterface("caption_entities", s.CaptionEntities)
	if err != nil {
		return params, err
	}
	params.AddBool("disable_notification", s.DisableNotification)
	params.AddBool("protect_content", s.ProtectContent)
	params.AddNonZero("reply_to_message_id", s.ReplyToMessageID)
	params.AddBool("allow_sending_without_reply", s.AllowSendingWithoutReply)
	err = params.AddInterface("reply_markup", s.ReplyMarkup)

	return params, err
}
func (s *CopyMessage) EndPoint() string {
	return config.EndpointCopyMessage
}

// SendPhoto Send photos. On success, the sent Message is returned.
type SendPhoto struct {
	ChatID                   int64           // required. use for user|channel as int
	ChatIDStr                string          // required. use for user|channel as string
	Username                 string          // required. use for channel
	Photo                    RequestFileData // required
	CustomFileName           string
	Caption                  string
	ParseMode                string
	CaptionEntities          []MessageEntity
	DisableNotification      bool
	ProtectContent           bool
	ReplyToMessageID         int
	AllowSendingWithoutReply bool
	ReplyMarkup              interface{}
}

func (s *SendPhoto) Params() (Params, error) {
	params := make(Params, 9)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddNonEmpty("caption", s.Caption)
	params.AddNonEmpty("parse_mode", s.ParseMode)
	err = params.AddInterface("caption_entities", s.CaptionEntities)
	if err != nil {
		return params, err
	}
	params.AddBool("disable_notification", s.DisableNotification)
	params.AddBool("protect_content", s.ProtectContent)
	params.AddNonZero("reply_to_message_id", s.ReplyToMessageID)
	params.AddBool("allow_sending_without_reply", s.AllowSendingWithoutReply)
	err = params.AddInterface("reply_markup", s.ReplyMarkup)

	return params, err
}
func (s *SendPhoto) Files() []RequestFile {
	files := []RequestFile{{
		Name:     "photo",
		Data:     s.Photo,
		FileName: s.CustomFileName,
	}}

	return files
}
func (s *SendPhoto) EndPoint() string {
	return config.EndpointSendPhoto
}

// SendAudio Send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future. For sending voice messages, use the sendVoice method instead.
type SendAudio struct {
	ChatID                   int64           // required. use for user|channel as int
	ChatIDStr                string          // required. use for user|channel as string
	Username                 string          // required. use for channel
	Audio                    RequestFileData // required
	CustomFileName           string
	Caption                  string
	ParseMode                string
	CaptionEntities          []MessageEntity
	Duration                 int
	Performer                string
	Title                    string
	Thumb                    RequestFileData
	DisableNotification      bool
	ProtectContent           bool
	ReplyToMessageID         int
	AllowSendingWithoutReply bool
	ReplyMarkup              interface{}
}

func (s *SendAudio) Params() (Params, error) {
	params := make(Params, 12)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddNonEmpty("caption", s.Caption)
	params.AddNonEmpty("parse_mode", s.ParseMode)
	err = params.AddInterface("caption_entities", s.CaptionEntities)
	if err != nil {
		return params, err
	}
	params.AddNonZero("duration", s.Duration)
	params.AddNonEmpty("performer", s.Performer)
	params.AddNonEmpty("title", s.Title)
	params.AddBool("disable_notification", s.DisableNotification)
	params.AddBool("protect_content", s.ProtectContent)
	params.AddNonZero("reply_to_message_id", s.ReplyToMessageID)
	params.AddBool("allow_sending_without_reply", s.AllowSendingWithoutReply)
	err = params.AddInterface("reply_markup", s.ReplyMarkup)

	return params, err
}
func (s *SendAudio) Files() []RequestFile {
	files := []RequestFile{{
		Name:     "audio",
		Data:     s.Audio,
		FileName: s.CustomFileName,
	}}

	if s.Thumb != nil {
		files = append(files, RequestFile{
			Name: "thumb",
			Data: s.Thumb,
		})
	}

	return files
}
func (s *SendAudio) EndPoint() string {
	return config.EndpointSendAudio
}

// SendDocument Send general files. On success, the sent Message is returned. Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
type SendDocument struct {
	ChatID                      int64           // required. use for user|channel as int
	ChatIDStr                   string          // required. use for user|channel as string
	Username                    string          // required. use for channel
	Document                    RequestFileData // required
	Thumb                       RequestFileData
	CustomFileName              string
	Caption                     string
	ParseMode                   string
	CaptionEntities             []MessageEntity
	DisableContentTypeDetection bool
	DisableNotification         bool
	ProtectContent              bool
	ReplyToMessageID            int
	AllowSendingWithoutReply    bool
	ReplyMarkup                 interface{}
}

func (s *SendDocument) Params() (Params, error) {
	params := make(Params, 10)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddNonEmpty("caption", s.Caption)
	params.AddNonEmpty("parse_mode", s.ParseMode)
	err = params.AddInterface("caption_entities", s.CaptionEntities)
	if err != nil {
		return params, err
	}
	params.AddBool("disable_content_type_detection", s.DisableContentTypeDetection)
	params.AddBool("disable_notification", s.DisableNotification)
	params.AddBool("protect_content", s.ProtectContent)
	params.AddNonZero("reply_to_message_id", s.ReplyToMessageID)
	params.AddBool("allow_sending_without_reply", s.AllowSendingWithoutReply)
	err = params.AddInterface("reply_markup", s.ReplyMarkup)

	return params, err
}
func (s *SendDocument) Files() []RequestFile {
	files := []RequestFile{{
		Name:     "document",
		Data:     s.Document,
		FileName: s.CustomFileName,
	}}
	if s.Thumb != nil {
		files = append(files, RequestFile{
			Name: "thumb",
			Data: s.Thumb,
		})
	}

	return files
}
func (s *SendDocument) EndPoint() string {
	return config.EndpointSendDocument
}

// SendVideo Send video files, Telegram clients support MPEG4 videos (other formats may be sent as Document). On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
type SendVideo struct {
	ChatID                   int64           // required. use for user|channel as int
	ChatIDStr                string          // required. use for user|channel as string
	Username                 string          // required. use for channel
	Video                    RequestFileData // required
	CustomFileName           string
	Duration                 int
	Weight                   int
	Height                   int
	Thumb                    RequestFileData
	Caption                  string
	ParseMode                string
	CaptionEntities          []MessageEntity
	SupportsStreaming        bool
	DisableNotification      bool
	ProtectContent           bool
	ReplyToMessageID         int
	AllowSendingWithoutReply bool
	ReplyMarkup              interface{}
}

func (s *SendVideo) Params() (Params, error) {
	params := make(Params, 13)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddNonZero("duration", s.Duration)
	params.AddNonZero("weight", s.Weight)
	params.AddNonZero("height", s.Height)
	params.AddNonEmpty("caption", s.Caption)
	params.AddNonEmpty("parse_mode", s.ParseMode)
	err = params.AddInterface("caption_entities", s.CaptionEntities)
	if err != nil {
		return params, err
	}
	params.AddBool("supports_streaming", s.SupportsStreaming)
	params.AddBool("disable_notification", s.DisableNotification)
	params.AddBool("protect_content", s.ProtectContent)
	params.AddNonZero("reply_to_message_id", s.ReplyToMessageID)
	params.AddBool("allow_sending_without_reply", s.AllowSendingWithoutReply)
	err = params.AddInterface("reply_markup", s.ReplyMarkup)

	return params, err
}
func (s *SendVideo) Files() []RequestFile {
	files := []RequestFile{{
		Name:     "video",
		Data:     s.Video,
		FileName: s.CustomFileName,
	}}
	if s.Thumb != nil {
		files = append(files, RequestFile{
			Name: "thumb",
			Data: s.Thumb,
		})
	}

	return files
}
func (s *SendVideo) EndPoint() string {
	return config.EndpointSendVideo
}

// SendAnimation Send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success, the sent Message is returned. Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
type SendAnimation struct {
	ChatID                   int64           // required. use for user|channel as int
	ChatIDStr                string          // required. use for user|channel as string
	Username                 string          // required. use for channel
	Animation                RequestFileData // required
	Duration                 int
	Weight                   int
	Height                   int
	Thumb                    RequestFileData
	Caption                  string
	ParseMode                string
	CaptionEntities          []MessageEntity
	DisableNotification      bool
	ProtectContent           bool
	ReplyToMessageID         int
	AllowSendingWithoutReply bool
	ReplyMarkup              interface{}
}

func (s *SendAnimation) Params() (Params, error) {
	params := make(Params, 12)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddNonZero("duration", s.Duration)
	params.AddNonZero("weight", s.Weight)
	params.AddNonZero("height", s.Height)
	params.AddNonEmpty("caption", s.Caption)
	params.AddNonEmpty("parse_mode", s.ParseMode)
	err = params.AddInterface("caption_entities", s.CaptionEntities)
	if err != nil {
		return params, err
	}
	params.AddBool("disable_notification", s.DisableNotification)
	params.AddBool("protect_content", s.ProtectContent)
	params.AddNonZero("reply_to_message_id", s.ReplyToMessageID)
	params.AddBool("allow_sending_without_reply", s.AllowSendingWithoutReply)
	err = params.AddInterface("reply_markup", s.ReplyMarkup)

	return params, err
}
func (s *SendAnimation) Files() []RequestFile {
	files := []RequestFile{{
		Name: "animation",
		Data: s.Animation,
	}}
	if s.Thumb != nil {
		files = append(files, RequestFile{
			Name: "thumb",
			Data: s.Thumb,
		})
	}

	return files
}
func (s *SendAnimation) EndPoint() string {
	return config.EndpointSendAnimation
}

// SendVoice Send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .OGG file encoded with OPUS (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
type SendVoice struct {
	ChatID                   int64           // required. use for user|channel as int
	ChatIDStr                string          // required. use for user|channel as string
	Username                 string          // required. use for channel
	Voice                    RequestFileData // required
	CustomFileName           string
	Caption                  string
	ParseMode                string
	CaptionEntities          []MessageEntity
	Duration                 int
	DisableNotification      bool
	ProtectContent           bool
	ReplyToMessageID         int
	AllowSendingWithoutReply bool
	ReplyMarkup              interface{}
}

func (s *SendVoice) Params() (Params, error) {
	params := make(Params, 10)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddNonEmpty("caption", s.Caption)
	params.AddNonEmpty("parse_mode", s.ParseMode)
	err = params.AddInterface("caption_entities", s.CaptionEntities)
	if err != nil {
		return params, err
	}
	params.AddNonZero("duration", s.Duration)
	params.AddBool("disable_notification", s.DisableNotification)
	params.AddBool("protect_content", s.ProtectContent)
	params.AddNonZero("reply_to_message_id", s.ReplyToMessageID)
	params.AddBool("allow_sending_without_reply", s.AllowSendingWithoutReply)
	err = params.AddInterface("reply_markup", s.ReplyMarkup)

	return params, err
}
func (s *SendVoice) Files() []RequestFile {
	files := []RequestFile{{
		Name:     "voice",
		Data:     s.Voice,
		FileName: s.CustomFileName,
	}}

	return files
}
func (s *SendVoice) EndPoint() string {
	return config.EndpointSendVoice
}

// SendVideoNote As of v.4.0, Telegram clients support rounded square MPEG4 videos of up to 1 minute long. Use this method to send video messages. On success, the sent Message is returned.
type SendVideoNote struct {
	ChatID                   int64           // required. use for user|channel as int
	ChatIDStr                string          // required. use for user|channel as string
	Username                 string          // required. use for channel
	VideoNote                RequestFileData // required.
	CustomFileName           string
	Duration                 int
	Length                   int
	Thumb                    RequestFileData
	DisableNotification      bool
	ProtectContent           bool
	ReplyToMessageID         int
	AllowSendingWithoutReply bool
	ReplyMarkup              interface{}
}

func (s *SendVideoNote) Params() (Params, error) {
	params := make(Params, 8)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddNonZero("duration", s.Duration)
	params.AddNonZero("length", s.Length)
	params.AddBool("disable_notification", s.DisableNotification)
	params.AddBool("protect_content", s.ProtectContent)
	params.AddNonZero("reply_to_message_id", s.ReplyToMessageID)
	params.AddBool("allow_sending_without_reply", s.AllowSendingWithoutReply)
	err = params.AddInterface("reply_markup", s.ReplyMarkup)

	return params, err
}
func (s *SendVideoNote) Files() []RequestFile {
	files := []RequestFile{{
		Name:     "video_note",
		Data:     s.VideoNote,
		FileName: s.CustomFileName,
	}}
	if s.Thumb != nil {
		files = append(files, RequestFile{
			Name: "thumb",
			Data: s.Thumb,
		})
	}

	return files
}
func (s *SendVideoNote) EndPoint() string {
	return config.EndpointSendVideoNote
}

// SendMediaGroup Use this method to send a group of photos, videos, documents or audios as an album. Documents and audio files can be only grouped in an album with messages of the same type. On success, an array of Messages that were sent is returned.
type SendMediaGroup struct {
	ChatID                   int64         // required. use for user|channel as int
	ChatIDStr                string        // required. use for user|channel as string
	Username                 string        // required. use for channel
	Media                    []interface{} // required
	DisableNotification      bool
	ProtectContent           bool
	ReplyToMessageID         int
	AllowSendingWithoutReply bool
}

func (s *SendMediaGroup) Params() (Params, error) {
	params := make(Params, 6)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddBool("disable_notification", s.DisableNotification)
	params.AddBool("protect_content", s.ProtectContent)
	params.AddNonZero("reply_to_message_id", s.ReplyToMessageID)
	params.AddBool("allow_sending_without_reply", s.AllowSendingWithoutReply)
	err = params.AddInterface("media", prepareInputMediaForParams(s.Media))

	return params, err
}
func (s *SendMediaGroup) Files() []RequestFile {
	return prepareInputMediaForFiles(s.Media)
}
func (s *SendMediaGroup) EndPoint() string {
	return config.EndpointSendMediaGroup
}

// prepareInputMediaParam evaluates a single InputMedia and determines if it needs to be modified for a successful upload. If it returns nil, then the value does not need to be included in the params. Otherwise, it will return the same type as was originally provided. The idx is used to calculate the file field name. If you only have a single file, 0 may be used. It is formatted into "attach://file-%d" for the primary media and "attach://file-%d-thumb" for thumbnails. It is expected to be used in conjunction with prepareInputMediaFile.
func prepareInputMediaParam(inputMedia interface{}, idx int) interface{} {
	switch m := inputMedia.(type) {
	case InputMediaPhoto:
		if m.Media.NeedsUpload() {
			m.Media = FileAttach(fmt.Sprintf("attach://file-%d", idx))
		}

		return m
	case InputMediaVideo:
		if m.Media.NeedsUpload() {
			m.Media = FileAttach(fmt.Sprintf("attach://file-%d", idx))
		}

		if m.Thumb != nil && m.Thumb.NeedsUpload() {
			m.Thumb = FileAttach(fmt.Sprintf("attach://file-%d-thumb", idx))
		}

		return m
	case InputMediaAudio:
		if m.Media.NeedsUpload() {
			m.Media = FileAttach(fmt.Sprintf("attach://file-%d", idx))
		}

		if m.Thumb != nil && m.Thumb.NeedsUpload() {
			m.Thumb = FileAttach(fmt.Sprintf("attach://file-%d-thumb", idx))
		}

		return m
	case InputMediaDocument:
		if m.Media.NeedsUpload() {
			m.Media = FileAttach(fmt.Sprintf("attach://file-%d", idx))
		}

		if m.Thumb != nil && m.Thumb.NeedsUpload() {
			m.Thumb = FileAttach(fmt.Sprintf("attach://file-%d-thumb", idx))
		}

		return m
	}

	return nil
}

// prepareInputMediaFile generates an array of RequestFile to provide for Fileable's files method. It returns an array as a single InputMedia may have multiple files, for the primary media and a thumbnail. The idx parameter is used to generate file field names. It uses the names "file-%d" for the main file and "file-%d-thumb" for the thumbnail. It is expected to be used in conjunction with prepareInputMediaParam.
func prepareInputMediaFile(inputMedia interface{}, idx int) []RequestFile {
	var files []RequestFile

	switch m := inputMedia.(type) {
	case InputMediaPhoto:
		if m.Media.NeedsUpload() {
			files = append(files, RequestFile{
				Name: fmt.Sprintf("file-%d", idx),
				Data: m.Media,
			})
		}
	case InputMediaVideo:
		if m.Media.NeedsUpload() {
			files = append(files, RequestFile{
				Name: fmt.Sprintf("file-%d", idx),
				Data: m.Media,
			})
		}

		if m.Thumb != nil && m.Thumb.NeedsUpload() {
			files = append(files, RequestFile{
				Name: fmt.Sprintf("file-%d", idx),
				Data: m.Thumb,
			})
		}
	case InputMediaDocument:
		if m.Media.NeedsUpload() {
			files = append(files, RequestFile{
				Name: fmt.Sprintf("file-%d", idx),
				Data: m.Media,
			})
		}

		if m.Thumb != nil && m.Thumb.NeedsUpload() {
			files = append(files, RequestFile{
				Name: fmt.Sprintf("file-%d", idx),
				Data: m.Thumb,
			})
		}
	case InputMediaAudio:
		if m.Media.NeedsUpload() {
			files = append(files, RequestFile{
				Name: fmt.Sprintf("file-%d", idx),
				Data: m.Media,
			})
		}

		if m.Thumb != nil && m.Thumb.NeedsUpload() {
			files = append(files, RequestFile{
				Name: fmt.Sprintf("file-%d", idx),
				Data: m.Thumb,
			})
		}
	}

	return files
}

// prepareInputMediaForParams calls prepareInputMediaParam for each item provided and returns a new array with the correct params for a request. It is expected that files will get data from the associated function, prepareInputMediaForFiles.
func prepareInputMediaForParams(inputMedia []interface{}) []interface{} {
	newMedia := make([]interface{}, len(inputMedia))
	copy(newMedia, inputMedia)

	for idx, media := range inputMedia {
		if param := prepareInputMediaParam(media, idx); param != nil {
			newMedia[idx] = param
		}
	}

	return newMedia
}

// prepareInputMediaForFiles calls prepareInputMediaFile for each item provided and returns a new array with the correct files for a request. It is expected that params will get data from the associated function, prepareInputMediaForParams.
func prepareInputMediaForFiles(inputMedia []interface{}) []RequestFile {
	var files []RequestFile

	for idx, media := range inputMedia {
		if file := prepareInputMediaFile(media, idx); file != nil {
			files = append(files, file...)
		}
	}

	return files
}

// SendLocation Send point on the map. On success, the sent Message is returned.
type SendLocation struct {
	ChatID                   int64   // required. use for user|channel as int
	ChatIDStr                string  // required. use for user|channel as string
	Username                 string  // required. use for channel
	Latitude                 float64 // required
	Longitude                float64 // required
	HorizontalAccuracy       float64
	LivePeriod               int
	Heading                  int
	ProximityAlertRadius     int
	DisableNotification      bool
	ProtectContent           bool
	ReplyToMessageID         int
	AllowSendingWithoutReply bool
	ReplyMarkup              interface{}
}

func (s *SendLocation) Params() (Params, error) {
	params := make(Params, 12)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddNonZeroFloat("latitude", s.Latitude)
	params.AddNonZeroFloat("longitude", s.Longitude)
	params.AddNonZeroFloat("horizontal_accuracy", s.HorizontalAccuracy)
	params.AddNonZero("live_period", s.LivePeriod)
	params.AddNonZero("heading", s.Heading)
	params.AddNonZero("proximity_alert_radius", s.ProximityAlertRadius)
	params.AddBool("disable_notification", s.DisableNotification)
	params.AddBool("protect_content", s.ProtectContent)
	params.AddNonZero("reply_to_message_id", s.ReplyToMessageID)
	params.AddBool("allow_sending_without_reply", s.AllowSendingWithoutReply)
	err = params.AddInterface("reply_markup", s.ReplyMarkup)

	return params, err
}
func (s *SendLocation) EndPoint() string {
	return config.EndpointSendLocation
}

// EditMessageLiveLocation Edit live location messages. A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
type EditMessageLiveLocation struct {
	ChatID               int64   // required if InlineMessageID is not specified. use for user|channel as int
	ChatIDStr            string  // required if InlineMessageID is not specified. use for user|channel as string
	Username             string  // required if InlineMessageID is not specified. use for channel
	MessageID            int     // required if InlineMessageID is not specified
	InlineMessageID      string  // required if ChatID & Username & MessageID are not specified
	Latitude             float64 // required
	Longitude            float64 // required
	HorizontalAccuracy   float64
	Heading              int
	ProximityAlertRadius int
	ReplyMarkup          *InlineKeyboardMarkup
}

func (s *EditMessageLiveLocation) Params() (Params, error) {
	params := make(Params, 8)

	if s.InlineMessageID != "" {
		params["inline_message_id"] = s.InlineMessageID
	} else {
		params.AddAt(s.Username)
		err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
		if err != nil {
			return params, err
		}
		params.AddNonZero("message_id", s.MessageID)
	}

	params.AddNonZeroFloat("latitude", s.Latitude)
	params.AddNonZeroFloat("longitude", s.Longitude)
	params.AddNonZeroFloat("horizontal_accuracy", s.HorizontalAccuracy)
	params.AddNonZero("heading", s.Heading)
	params.AddNonZero("proximity_alert_radius", s.ProximityAlertRadius)

	err := params.AddInterface("reply_markup", s.ReplyMarkup)

	return params, err
}
func (s *EditMessageLiveLocation) EndPoint() string {
	return config.EndpointEditMessageLiveLocation
}

// StopMessageLiveLocation Stop updating a live location message before live_period expires. On success, if the message is not an inline message, the edited Message is returned, otherwise True is returned.
type StopMessageLiveLocation struct {
	ChatID          int64  // required if InlineMessageID is not specified. use for user|channel as int
	ChatIDStr       string // required if InlineMessageID is not specified. use for user|channel as string
	Username        string // required if InlineMessageID is not specified. use for channel
	MessageID       int    // required if InlineMessageID is not specified
	InlineMessageID string // required if ChatID & Username & MessageID are not specified
	ReplyMarkup     *InlineKeyboardMarkup
}

func (s *StopMessageLiveLocation) Params() (Params, error) {
	params := make(Params, 3)

	if s.InlineMessageID != "" {
		params["inline_message_id"] = s.InlineMessageID
	} else {
		params.AddAt(s.Username)
		err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
		if err != nil {
			return params, err
		}
		params.AddNonZero("message_id", s.MessageID)
	}

	err := params.AddInterface("reply_markup", s.ReplyMarkup)

	return params, err
}
func (s *StopMessageLiveLocation) EndPoint() string {
	return config.EndpointStopMessageLiveLocation
}

// SendVenue Send information about a venue. On success, the sent Message is returned.
type SendVenue struct {
	ChatID                   int64   // required. use for user|channel as int
	ChatIDStr                string  // required. use for user|channel as string
	Username                 string  // required. use for channel
	Latitude                 float64 // required
	Longitude                float64 // required
	Title                    string  // required
	Address                  string  // required
	FoursquareID             string
	FoursquareType           string
	GooglePlaceID            string
	GooglePlaceType          string
	DisableNotification      bool
	ProtectContent           bool
	ReplyToMessageID         int
	AllowSendingWithoutReply bool
	ReplyMarkup              interface{}
}

func (s *SendVenue) Params() (Params, error) {
	params := make(Params, 14)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddNonZeroFloat("latitude", s.Latitude)
	params.AddNonZeroFloat("longitude", s.Longitude)
	params["title"] = s.Title
	params["address"] = s.Address
	params.AddNonEmpty("foursquare_id", s.FoursquareID)
	params.AddNonEmpty("foursquare_type", s.FoursquareType)
	params.AddNonEmpty("google_place_id", s.GooglePlaceID)
	params.AddNonEmpty("google_place_type", s.GooglePlaceType)
	params.AddBool("disable_notification", s.DisableNotification)
	params.AddBool("protect_content", s.ProtectContent)
	params.AddNonZero("reply_to_message_id", s.ReplyToMessageID)
	params.AddBool("allow_sending_without_reply", s.AllowSendingWithoutReply)
	err = params.AddInterface("reply_markup", s.ReplyMarkup)

	return params, err
}
func (s *SendVenue) EndPoint() string {
	return config.EndpointSendVenue
}

// SendContact Send phone contacts. On success, the sent Message is returned.
type SendContact struct {
	ChatID                   int64  // required. use for user|channel as int
	ChatIDStr                string // required. use for user|channel as string
	Username                 string // required. use for channel
	PhoneNumber              string // required
	FirstName                string // required
	LastName                 string
	VCard                    string
	DisableNotification      bool
	ProtectContent           bool
	ReplyToMessageID         int
	AllowSendingWithoutReply bool
	ReplyMarkup              interface{}
}

func (s *SendContact) Params() (Params, error) {
	params := make(Params, 10)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params["phone_number"] = s.PhoneNumber
	params["first_name"] = s.FirstName
	params.AddNonEmpty("last_name", s.LastName)
	params.AddNonEmpty("vcard", s.VCard)
	params.AddBool("disable_notification", s.DisableNotification)
	params.AddBool("protect_content", s.ProtectContent)
	params.AddNonZero("reply_to_message_id", s.ReplyToMessageID)
	params.AddBool("allow_sending_without_reply", s.AllowSendingWithoutReply)
	err = params.AddInterface("reply_markup", s.ReplyMarkup)

	return params, err
}
func (s *SendContact) EndPoint() string {
	return config.EndpointSendContact
}

// SendPoll Send a native poll. On success, the sent Message is returned.
type SendPoll struct {
	ChatID                   int64    // required. use for user|channel as int
	ChatIDStr                string   // required. use for user|channel as string
	Username                 string   // required. use for channel
	Question                 string   // required
	Options                  []string // required
	IsAnonymous              bool
	Type                     string
	AllowsMultipleAnswers    bool
	CorrectOptionID          int64
	Explanation              string
	ExplanationParseMode     string
	ExplanationEntities      []MessageEntity
	OpenPeriod               int
	CloseDate                int
	IsClosed                 bool
	DisableNotification      bool
	ProtectContent           bool
	ReplyToMessageID         int
	AllowSendingWithoutReply bool
	ReplyMarkup              interface{}
}

func (s *SendPoll) Params() (Params, error) {
	params := make(Params, 18)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params["question"] = s.Question
	err = params.AddInterface("options", s.Options)
	if err != nil {
		return params, err
	}
	params.AddBool("is_anonymous", s.IsAnonymous)
	params.AddNonEmpty("type", s.Type)
	params.AddBool("allows_multiple_answers", s.AllowsMultipleAnswers)
	params.AddNonZero64("correct_option_id", s.CorrectOptionID)
	params.AddBool("is_closed", s.IsClosed)
	params.AddNonEmpty("explanation", s.Explanation)
	params.AddNonEmpty("explanation_parse_mode", s.ExplanationParseMode)
	params.AddNonZero("open_period", s.OpenPeriod)
	params.AddNonZero("close_date", s.CloseDate)
	err = params.AddInterface("explanation_entities", s.ExplanationEntities)
	if err != nil {
		return params, err
	}
	params.AddBool("disable_notification", s.DisableNotification)
	params.AddBool("protect_content", s.ProtectContent)
	params.AddNonZero("reply_to_message_id", s.ReplyToMessageID)
	params.AddBool("allow_sending_without_reply", s.AllowSendingWithoutReply)
	err = params.AddInterface("reply_markup", s.ReplyMarkup)

	return params, err
}
func (s *SendPoll) EndPoint() string {
	return config.EndpointSendPoll
}

// SendDice Send an animated emoji that will display a random value. On success, the sent Message is returned.
type SendDice struct {
	ChatID                   int64  // required. use for user|channel as int
	ChatIDStr                string // required. use for user|channel as string
	Username                 string // required. use for channel
	Emoji                    string // Emoji on which the dice throw animation is based. Currently, must be one of “🎲”, “🎯”, “🏀”, “⚽”, “🎳”, or “🎰”. Dice can have values 1-6 for “🎲”, “🎯” and “🎳”, values 1-5 for “🏀” and “⚽”, and values 1-64 for “🎰”. Defaults to “🎲
	DisableNotification      bool
	ProtectContent           bool
	ReplyToMessageID         int
	AllowSendingWithoutReply bool
	ReplyMarkup              interface{}
}

func (s *SendDice) Params() (Params, error) {
	params := make(Params, 7)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddNonEmpty("emoji", s.Emoji)
	params.AddBool("disable_notification", s.DisableNotification)
	params.AddBool("protect_content", s.ProtectContent)
	params.AddNonZero("reply_to_message_id", s.ReplyToMessageID)
	params.AddBool("allow_sending_without_reply", s.AllowSendingWithoutReply)
	err = params.AddInterface("reply_markup", s.ReplyMarkup)

	return params, err
}
func (s *SendDice) EndPoint() string {
	return config.EndpointSendDice
}

// SendChatAction Send an animated emoji that will display a random value. On success, the sent Message is returned.
type SendChatAction struct {
	ChatID    int64  // required. use for user|channel as int
	ChatIDStr string // required. use for user|channel as string
	Username  string // required. use for channel
	Action    string // required. `typing` for text messages, `upload_photo` for photos, `record_video` or `upload_video` for videos, `record_voice` or `upload_voice` for voice notes, `upload_document` for general files, `choose_sticker` for stickers, `find_location` for location data, `record_video_note` or `upload_video_note` for video notes.
}

func (s *SendChatAction) Params() (Params, error) {
	params := make(Params, 2)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params["action"] = s.Action

	return params, nil
}
func (s *SendChatAction) EndPoint() string {
	return config.EndpointSendChatAction
}

// GetUserProfilePhotos Get a list of profile pictures for a user. Returns a UserProfilePhotos object.
type GetUserProfilePhotos struct {
	UserID int64 // required
	Offset int
	Limit  int
}

func (s *GetUserProfilePhotos) Params() (Params, error) {
	params := make(Params, 3)

	params.AddNonZero64("user_id", s.UserID)
	params.AddNonZero("offset", s.Offset)
	params.AddNonZero("limit", s.Limit)

	return params, nil
}
func (s *GetUserProfilePhotos) EndPoint() string {
	return config.EndpointGetUserProfilePhotos
}

// Link returns a full path to the download URL for a File.
//
// It requires the Bot token to create the link.
func (f *File) Link(token string) string {
	return fmt.Sprintf(config.APIFileEndpoint, token, f.FilePath)
}

// GetFile Get basic information about a file and prepare it for downloading. For the moment, bots can download files of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.
// Note: This function may not preserve the original file name and MIME type. You should save the file's MIME type and name (if available) when the File object is received.
type GetFile struct {
	FileID string // required
}

func (s *GetFile) Params() (Params, error) {
	params := make(Params, 1)

	params["file_id"] = s.FileID

	return params, nil
}
func (s *GetFile) EndPoint() string {
	return config.EndpointGetFile
}

// BanChatMember Ban a user in a group, a supergroup or a channel. In the case of supergroups and channels, the user will not be able to return to the chat on their own using invite links, etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
type BanChatMember struct {
	ChatID         int64  // required. use for group|supergroup|channel as int
	ChatIDStr      string // required. use for group|supergroup|channel as string
	Username       string // required. use for group|supergroup|channel
	UserID         int64  // required
	UntilDate      int64
	RevokeMessages bool
}

func (s *BanChatMember) Params() (Params, error) {
	params := make(Params, 4)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddNonZero64("user_id", s.UserID)
	params.AddNonZero64("until_date", s.UntilDate)
	params.AddBool("revoke_messages", s.RevokeMessages)

	return params, nil
}
func (s *BanChatMember) EndPoint() string {
	return config.EndpointBanChatMember
}

// UnbanChatMember Unban a previously banned user in a supergroup or channel. The user will not return to the group or channel automatically, but will be able to join via link, etc. The bot must be an administrator for this to work. By default, this method guarantees that after the call the user is not a member of the chat, but will be able to join it. So if the user is a member of the chat they will also be removed from the chat. If you don't want this, use the parameter only_if_banned. Returns True on success.
type UnbanChatMember struct {
	ChatID       int64  // required. use for group|supergroup|channel as int
	ChatIDStr    string // required. use for group|supergroup|channel as string
	Username     string // required. use for group|supergroup|channel
	UserID       int64  // required
	OnlyIfBanned bool
}

func (s *UnbanChatMember) Params() (Params, error) {
	params := make(Params, 3)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddNonZero64("user_id", s.UserID)
	params.AddBool("only_if_banned", s.OnlyIfBanned)

	return params, nil
}
func (s *UnbanChatMember) EndPoint() string {
	return config.EndpointUnbanChatMember
}

// RestrictChatMember Restrict a user in a supergroup. The bot must be an administrator in the supergroup for this to work and must have the appropriate administrator rights. Pass True for all permissions to lift restrictions from a user. Returns True on success.
type RestrictChatMember struct {
	ChatID      int64            // required. use for supergroup as int
	ChatIDStr   string           // required. use for supergroup as string
	Username    string           // required. use for supergroup
	UserID      int64            // required
	Permissions *ChatPermissions // required
	UntilDate   int64
}

func (s *RestrictChatMember) Params() (Params, error) {
	params := make(Params, 4)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddNonZero64("user_id", s.UserID)
	err = params.AddInterface("permissions", s.Permissions)
	if err != nil {
		return params, err
	}
	params.AddNonZero64("until_date", s.UntilDate)

	return params, nil
}
func (s *RestrictChatMember) EndPoint() string {
	return config.EndpointRestrictChatMember
}

// PromoteChatMember Promote or demote a user in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Pass False for all boolean parameters to demote a user. Returns True on success.
type PromoteChatMember struct {
	ChatID              int64  // required. use for supergroup|channel as int
	ChatIDStr           string // required. use for supergroup|channel as string
	Username            string // required. use for supergroup|channel
	UserID              int64  // required
	IsAnonymous         bool
	CanManageChat       bool
	CanChangeInfo       bool
	CanPostMessages     bool
	CanEditMessages     bool
	CanDeleteMessages   bool
	CanManageVideoChats bool
	CanInviteUsers      bool
	CanRestrictMembers  bool
	CanPinMessages      bool
	CanPromoteMembers   bool
}

func (s *PromoteChatMember) Params() (Params, error) {
	params := make(Params, 13)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddNonZero64("user_id", s.UserID)
	params.AddBool("is_anonymous", s.IsAnonymous)
	params.AddBool("can_manage_chat", s.CanManageChat)
	params.AddBool("can_change_info", s.CanChangeInfo)
	params.AddBool("can_post_messages", s.CanPostMessages)
	params.AddBool("can_edit_messages", s.CanEditMessages)
	params.AddBool("can_delete_messages", s.CanDeleteMessages)
	params.AddBool("can_manage_video_chats", s.CanManageVideoChats)
	params.AddBool("can_invite_users", s.CanInviteUsers)
	params.AddBool("can_restrict_members", s.CanRestrictMembers)
	params.AddBool("can_pin_messages", s.CanPinMessages)
	params.AddBool("can_promote_members", s.CanPromoteMembers)

	return params, nil
}
func (s *PromoteChatMember) EndPoint() string {
	return config.EndpointPromoteChatMember
}

// SetChatAdministratorCustomTitle Set a custom title for an administrator in a supergroup promoted by the bot. Returns True on success.
type SetChatAdministratorCustomTitle struct {
	ChatID      int64  // required. use for supergroup as int
	ChatIDStr   string // required. use for supergroup as string
	Username    string // required. use for supergroup
	UserID      int64  // required
	CustomTitle string
}

func (s *SetChatAdministratorCustomTitle) Params() (Params, error) {
	params := make(Params, 3)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddNonZero64("user_id", s.UserID)
	params.AddNonEmpty("custom_title", s.CustomTitle)

	return params, nil
}
func (s *SetChatAdministratorCustomTitle) EndPoint() string {
	return config.EndpointSetChatAdministratorCustomTitle
}

// BanChatSenderChat Ban a channel chat in a supergroup or a channel. Until the chat is unbanned, the owner of the banned chat won't be able to send messages on behalf of any of their channels. The bot must be an administrator in the supergroup or channel for this to work and must have the appropriate administrator rights. Returns True on success.
type BanChatSenderChat struct {
	ChatID       int64  // required. use for supergroup|channel as int
	ChatIDStr    string // required. use for supergroup|channel as string
	Username     string // required. use for supergroup|channel
	SenderChatID int64  // required
}

func (s *BanChatSenderChat) Params() (Params, error) {
	params := make(Params, 2)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddNonZero64("sender_chat_id", s.SenderChatID)

	return params, nil
}
func (s *BanChatSenderChat) EndPoint() string {
	return config.EndpointBanChatSenderChat
}

// UnbanChatSenderChat Unban a previously banned channel chat in a supergroup or channel. The bot must be an administrator for this to work and must have the appropriate administrator rights. Returns True on success.
type UnbanChatSenderChat struct {
	ChatID       int64  // required. use for supergroup|channel as int
	ChatIDStr    string // required. use for supergroup|channel as string
	Username     string // required. use for supergroup|channel
	SenderChatID int64  // required
}

func (s *UnbanChatSenderChat) Params() (Params, error) {
	params := make(Params, 2)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddNonZero64("sender_chat_id", s.SenderChatID)

	return params, nil
}
func (s *UnbanChatSenderChat) EndPoint() string {
	return config.EndpointUnbanChatSenderChat
}

// SetChatPermissions Set default chat permissions for all members. The bot must be an administrator in the group or a supergroup for this to work and must have the can_restrict_members administrator rights. Returns True on success.
type SetChatPermissions struct {
	ChatID      int64            // required. use for group|supergroup as int
	ChatIDStr   string           // required. use for group|supergroup as string
	Username    string           // required. use for group|supergroup
	Permissions *ChatPermissions // required
}

func (s *SetChatPermissions) Params() (Params, error) {
	params := make(Params, 2)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	err = params.AddInterface("permissions", s.Permissions)

	return params, err
}
func (s *SetChatPermissions) EndPoint() string {
	return config.EndpointSetChatPermissions
}

// ExportChatInviteLink Generate a new primary invite link for a chat; any previously generated primary link is revoked. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the new invite link as String on success.
// Note: Each administrator in a chat generates their own invite links. Bots can't use invite links generated by other administrators. If you want your bot to work with invite links, it will need to generate its own link using exportChatInviteLink or by calling the getChat method. If your bot needs to generate a new primary invite link replacing its previous one, use exportChatInviteLink again.
type ExportChatInviteLink struct {
	ChatID    int64  // required. use for group|supergroup|channel as int
	ChatIDStr string // required. use for group|supergroup|channel as string
	Username  string // required. use for group|supergroup|channel
}

func (s *ExportChatInviteLink) Params() (Params, error) {
	params := make(Params, 1)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)

	return params, err
}
func (s *ExportChatInviteLink) EndPoint() string {
	return config.EndpointExportChatInviteLink
}

// CreateChatInviteLink Create an additional invite link for a chat. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. The link can be revoked using the method revokeChatInviteLink. Returns the new invite link as ChatInviteLink object.
type CreateChatInviteLink struct {
	ChatID             int64  // required. use for group|supergroup|channel as int
	ChatIDStr          string // required. use for group|supergroup|channel as string
	Username           string // required. use for group|supergroup|channel
	Name               string
	ExpireDate         int
	MemberLimit        int
	CreatesJoinRequest bool
}

func (s *CreateChatInviteLink) Params() (Params, error) {
	params := make(Params, 5)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddNonEmpty("name", s.Name)
	params.AddNonZero("expire_date", s.ExpireDate)
	params.AddNonZero("member_limit", s.MemberLimit)
	params.AddBool("creates_join_request", s.CreatesJoinRequest)

	return params, nil
}
func (s *CreateChatInviteLink) EndPoint() string {
	return config.EndpointCreateChatInviteLink
}

// EditChatInviteLink Edit a non-primary invite link created by the bot. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the edited invite link as a ChatInviteLink object.
type EditChatInviteLink struct {
	ChatID             int64  // required. use for group|supergroup|channel as int
	ChatIDStr          string // required. use for group|supergroup|channel as string
	Username           string // required. use for group|supergroup|channel
	InviteLink         string // required
	Name               string
	ExpireDate         int
	MemberLimit        int
	CreatesJoinRequest bool
}

func (s *EditChatInviteLink) Params() (Params, error) {
	params := make(Params, 6)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params["invite_link"] = s.InviteLink
	params.AddNonEmpty("name", s.Name)
	params.AddNonZero("expire_date", s.ExpireDate)
	params.AddNonZero("member_limit", s.MemberLimit)
	params.AddBool("creates_join_request", s.CreatesJoinRequest)

	return params, nil
}
func (s *EditChatInviteLink) EndPoint() string {
	return config.EndpointEditChatInviteLink
}

// RevokeChatInviteLink Revoke an invite link created by the bot. If the primary link is revoked, a new link is automatically generated. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the revoked invite link as ChatInviteLink object.
type RevokeChatInviteLink struct {
	ChatID     int64  // required. use for group|supergroup|channel as int
	ChatIDStr  string // required. use for group|supergroup|channel as string
	Username   string // required. use for group|supergroup|channel
	InviteLink string // required
}

func (s *RevokeChatInviteLink) Params() (Params, error) {
	params := make(Params, 2)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params["invite_link"] = s.InviteLink

	return params, nil
}
func (s *RevokeChatInviteLink) EndPoint() string {
	return config.EndpointRevokeChatInviteLink
}

// ApproveChatJoinRequest Approve a chat join request. The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
type ApproveChatJoinRequest struct {
	ChatID    int64  // required. use for group|supergroup|channel as int
	ChatIDStr string // required. use for group|supergroup|channel as string
	Username  string // required. use for group|supergroup|channel
	UserID    int64  // required
}

func (s *ApproveChatJoinRequest) Params() (Params, error) {
	params := make(Params, 2)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddNonZero("user_id", int(s.UserID))

	return params, nil
}
func (s *ApproveChatJoinRequest) EndPoint() string {
	return config.EndpointApproveChatJoinRequest
}

// DeclineChatJoinRequest Decline a chat join request. The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
type DeclineChatJoinRequest struct {
	ChatID    int64  // required. use for group|supergroup|channel as int
	ChatIDStr string // required. use for group|supergroup|channel as string
	Username  string // required. use for group|supergroup|channel
	UserID    int64  // required
}

func (s *DeclineChatJoinRequest) Params() (Params, error) {
	params := make(Params, 2)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddNonZero("user_id", int(s.UserID))

	return params, nil
}
func (s *DeclineChatJoinRequest) EndPoint() string {
	return config.EndpointDeclineChatJoinRequest
}

// SetChatPhoto Set a new profile photo for the chat. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
type SetChatPhoto struct {
	ChatID    int64           // required. use for group|supergroup|channel as int
	ChatIDStr string          // required. use for group|supergroup|channel as string
	Username  string          // required. use for group|supergroup|channel
	Photo     RequestFileData // required must be uploaded or string path
}

func (s *SetChatPhoto) Params() (Params, error) {
	params := make(Params, 1)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)

	return params, err
}
func (s *SetChatPhoto) Files() []RequestFile {
	return []RequestFile{{
		Name: "photo",
		Data: s.Photo,
	}}
}
func (s *SetChatPhoto) EndPoint() string {
	return config.EndpointSetChatPhoto
}

// DeleteChatPhoto Delete a chat photo. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
type DeleteChatPhoto struct {
	ChatID    int64  // required. use for group|supergroup|channel as int
	ChatIDStr string // required. use for group|supergroup|channel as string
	Username  string // required. use for group|supergroup|channel
}

func (s *DeleteChatPhoto) Params() (Params, error) {
	params := make(Params, 1)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)

	return params, err
}
func (s *DeleteChatPhoto) EndPoint() string {
	return config.EndpointDeleteChatPhoto
}

// SetChatTitle Change the title of a chat. Titles can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
type SetChatTitle struct {
	ChatID    int64  // required. use for group|supergroup|channel as int
	ChatIDStr string // required. use for group|supergroup|channel as string
	Username  string // required. use for group|supergroup|channel
	Title     string // required
}

func (s *SetChatTitle) Params() (Params, error) {
	params := make(Params, 2)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params["title"] = s.Title

	return params, nil
}
func (s *SetChatTitle) EndPoint() string {
	return config.EndpointSetChatTitle
}

// SetChatDescription Change the description of a group, a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
type SetChatDescription struct {
	ChatID      int64  // required. use for group|supergroup|channel as int
	ChatIDStr   string // required. use for group|supergroup|channel as string
	Username    string // required. use for group|supergroup|channel
	Description string // required
}

func (s *SetChatDescription) Params() (Params, error) {
	params := make(Params, 2)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params["description"] = s.Description

	return params, nil
}
func (s *SetChatDescription) EndPoint() string {
	return config.EndpointSetChatDescription
}

// PinChatMessage Add a message to the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel. Returns True on success.
type PinChatMessage struct {
	ChatID              int64  // required. use for group|supergroup|channel as int
	ChatIDStr           string // required. use for group|supergroup|channel as string
	Username            string // required. use for group|supergroup|channel
	MessageID           int    // required
	DisableNotification bool
}

func (s *PinChatMessage) Params() (Params, error) {
	params := make(Params, 3)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddNonZero("message_id", s.MessageID)
	params.AddBool("disable_notification", s.DisableNotification)

	return params, nil
}
func (s *PinChatMessage) EndPoint() string {
	return config.EndpointPinChatMessage
}

// UnpinChatMessage Remove a message from the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel. Returns True on success.
type UnpinChatMessage struct {
	ChatID    int64  // required. use for group|supergroup|channel as int
	ChatIDStr string // required. use for group|supergroup|channel as string
	Username  string // required. use for group|supergroup|channel
	MessageID int    // required
}

func (s *UnpinChatMessage) Params() (Params, error) {
	params := make(Params, 2)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddNonZero("message_id", s.MessageID)

	return params, nil
}
func (s *UnpinChatMessage) EndPoint() string {
	return config.EndpointUnpinChatMessage
}

// UnpinAllChatMessages Clear the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel. Returns True on success.
type UnpinAllChatMessages struct {
	ChatID    int64  // required. use for group|supergroup|channel as int
	ChatIDStr string // required. use for group|supergroup|channel as string
	Username  string // required. use for group|supergroup|channel
}

func (s *UnpinAllChatMessages) Params() (Params, error) {
	params := make(Params, 1)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)

	return params, err
}
func (s *UnpinAllChatMessages) EndPoint() string {
	return config.EndpointUnpinAllChatMessages
}

// LeaveChat Your bot to leave a group, supergroup or channel. Returns True on success.
type LeaveChat struct {
	ChatID    int64  // required. use for group|supergroup|channel as int
	ChatIDStr string // required. use for group|supergroup|channel as string
	Username  string // required. use for group|supergroup|channel
}

func (s *LeaveChat) Params() (Params, error) {
	params := make(Params, 1)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)

	return params, err
}
func (s *LeaveChat) EndPoint() string {
	return config.EndpointLeaveChat
}

// GetChat Use this method to get up to date information about the chat (current name of the user for one-on-one conversations, current username of a user, group or channel, etc.). Returns a Chat object on success.
type GetChat struct {
	ChatID    int64  // required. use for group|supergroup|channel as int
	ChatIDStr string // required. use for group|supergroup|channel as string
	Username  string // required. use for group|supergroup|channel
}

func (s *GetChat) Params() (Params, error) {
	params := make(Params, 1)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)

	return params, err
}
func (s *GetChat) EndPoint() string {
	return config.EndpointGetChat
}

// GetChatAdministrators Get a list of administrators in a chat. On success, returns an Array of ChatMember objects that contains information about all chat administrators except other bots. If the chat is a group or a supergroup and no administrators were appointed, only the creator will be returned.
type GetChatAdministrators struct {
	ChatID    int64  // required. use for group|supergroup|channel as int
	ChatIDStr string // required. use for group|supergroup|channel as string
	Username  string // required. use for group|supergroup|channel
}

func (s *GetChatAdministrators) Params() (Params, error) {
	params := make(Params, 1)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)

	return params, err
}
func (s *GetChatAdministrators) EndPoint() string {
	return config.EndpointGetChatAdministrators
}

// GetChatMemberCount Get the number of members in a chat. Returns Int on success.
type GetChatMemberCount struct {
	ChatID    int64  // required. use for group|supergroup|channel as int
	ChatIDStr string // required. use for group|supergroup|channel as string
	Username  string // required. use for group|supergroup|channel
}

func (s *GetChatMemberCount) Params() (Params, error) {
	params := make(Params, 1)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)

	return params, err
}
func (s *GetChatMemberCount) EndPoint() string {
	return config.EndpointGetChatMemberCount
}

// GetChatMember Use this method to get information about a member of a chat. Returns a ChatMember object on success.
type GetChatMember struct {
	ChatID    int64  // required. use for group|supergroup|channel as int
	ChatIDStr string // required. use for group|supergroup|channel as string
	Username  string // required. use for group|supergroup|channel
	UserID    int64  // required
}

func (s *GetChatMember) Params() (Params, error) {
	params := make(Params, 2)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddNonZero64("user_id", s.UserID)

	return params, nil
}
func (s *GetChatMember) EndPoint() string {
	return config.EndpointGetChatMember
}

// SetChatStickerSet Set a new group sticker set for a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
type SetChatStickerSet struct {
	ChatID         int64  // required. use for supergroup as int
	ChatIDStr      string // required. use for supergroup as string
	Username       string // required. use for supergroup
	StickerSetName string // required
}

func (s *SetChatStickerSet) Params() (Params, error) {
	params := make(Params, 2)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params["sticker_set_name"] = s.StickerSetName

	return params, nil
}
func (s *SetChatStickerSet) EndPoint() string {
	return config.EndpointSetChatStickerSet
}

// DeleteChatStickerSet Delete a group sticker set from a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
type DeleteChatStickerSet struct {
	ChatID    int64  // required. use for supergroup as int
	ChatIDStr string // required. use for supergroup as string
	Username  string // required. use for supergroup
}

func (s *DeleteChatStickerSet) Params() (Params, error) {
	params := make(Params, 1)

	params.AddAt(s.Username)
	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)

	return params, err
}
func (s *DeleteChatStickerSet) EndPoint() string {
	return config.EndpointDeleteChatStickerSet
}

// AnswerCallbackQuery Send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. On success, True is returned. Alternatively, the user can be redirected to the specified Game URL. For this option to work, you must first create a game for your bot via @BotFather and accept the terms. Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
type AnswerCallbackQuery struct {
	CallbackQueryID string // required
	Text            string
	ShowAlert       bool
	URL             string
	CacheTime       int
}

func (s *AnswerCallbackQuery) Params() (Params, error) {
	params := make(Params, 5)

	params["callback_query_id"] = s.CallbackQueryID
	params.AddNonEmpty("text", s.Text)
	params.AddBool("show_alert", s.ShowAlert)
	params.AddNonEmpty("url", s.URL)
	params.AddNonZero("cache_time", s.CacheTime)

	return params, nil
}
func (s *AnswerCallbackQuery) EndPoint() string {
	return config.EndpointAnswerCallbackQuery
}

// SetMyCommands Change the list of the bot's commands. See https://core.telegram.org/bots#commands for more details about bot commands. Returns True on success.
type SetMyCommands struct {
	Commands     []*BotCommand // required
	Scope        *BotCommandScope
	LanguageCode string
}

func (s *SetMyCommands) Params() (Params, error) {
	params := make(Params, 3)

	err := params.AddInterface("commands", s.Commands)
	if err != nil {
		return params, err
	}
	err = params.AddInterface("scope", s.Scope)
	if err != nil {
		return params, err
	}
	params.AddNonEmpty("language_code", s.LanguageCode)

	return params, nil
}
func (s *SetMyCommands) EndPoint() string {
	return config.EndpointSetMyCommands
}

// DeleteMyCommands Delete the list of the bot's commands for the given scope and user language. After deletion, higher level commands will be shown to affected users. Returns True on success.
type DeleteMyCommands struct {
	Scope        *BotCommandScope
	LanguageCode string
}

func (s *DeleteMyCommands) Params() (Params, error) {
	params := make(Params, 2)

	err := params.AddInterface("scope", s.Scope)
	if err != nil {
		return params, err
	}
	params.AddNonEmpty("language_code", s.LanguageCode)

	return params, nil
}
func (s *DeleteMyCommands) EndPoint() string {
	return config.EndpointDeleteMyCommands
}

// GetMyCommands Get the current list of the bot's commands for the given scope and user language. Returns Array of BotCommand on success. If commands aren't set, an empty list is returned.
type GetMyCommands struct {
	Scope        *BotCommandScope
	LanguageCode string
}

func (s *GetMyCommands) Params() (Params, error) {
	params := make(Params, 2)

	err := params.AddInterface("scope", s.Scope)
	if err != nil {
		return params, err
	}
	params.AddNonEmpty("language_code", s.LanguageCode)

	return params, nil
}
func (s *GetMyCommands) EndPoint() string {
	return config.EndpointGetMyCommands
}

// SetChatMenuButton Change the bot's menu button in a private chat, or the default menu button. Returns True on success.
type SetChatMenuButton struct {
	ChatID    int64  // required. use for chat|channel as int
	ChatIDStr string // required. use for chat|channel as string
	Username  string // required. use for chat|channel

	MenuButton *MenuButton
}

func (s *SetChatMenuButton) Params() (Params, error) {
	params := make(Params, 2)

	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	err = params.AddInterface("menu_button", s.MenuButton)

	return params, err
}
func (s *SetChatMenuButton) EndPoint() string {
	return config.EndpointSetChatMenuButton
}

// GetChatMenuButton Get the current value of the bot's menu button in a private chat, or the default menu button. Returns MenuButton on success.
type GetChatMenuButton struct {
	ChatID    int64  // required. use for chat|channel as int
	ChatIDStr string // required. use for chat|channel as string
	Username  string // required. use for chat|channel
}

func (s *GetChatMenuButton) Params() (Params, error) {
	params := make(Params, 1)

	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)

	return params, err
}
func (s *GetChatMenuButton) EndPoint() string {
	return config.EndpointGetChatMenuButton
}

// SetMyDefaultAdministratorRights Change the default administrator rights requested by the bot when it's added as an administrator to groups or channels. These rights will be suggested to users, but they are are free to modify the list before adding the bot. Returns True on success.
type SetMyDefaultAdministratorRights struct {
	Rights      ChatAdministratorRights
	ForChannels bool
}

func (s *SetMyDefaultAdministratorRights) Params() (Params, error) {
	params := make(Params, 2)

	err := params.AddInterface("rights", s.Rights)
	if err != nil {
		return params, err
	}
	params.AddBool("for_channels", s.ForChannels)

	return params, nil
}
func (s *SetMyDefaultAdministratorRights) EndPoint() string {
	return config.EndpointSetMyDefaultAdministratorRights
}

// GetMyDefaultAdministratorRights Get the current default administrator rights of the bot. Returns ChatAdministratorRights on success.
type GetMyDefaultAdministratorRights struct {
	ForChannels bool
}

func (s *GetMyDefaultAdministratorRights) Params() (Params, error) {
	params := make(Params, 1)

	params.AddBool("for_channels", s.ForChannels)

	return params, nil
}
func (s *GetMyDefaultAdministratorRights) EndPoint() string {
	return config.EndpointGetMyDefaultAdministratorRights
}

// EditMessageText Edit text and game messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
type EditMessageText struct {
	ChatID                int64  // required if InlineMessageID is not specified. use for chat|channel as int
	ChatIDStr             string // required if InlineMessageID is not specified. use for chat|channel as string
	Username              string // required if InlineMessageID is not specified. use for chat|channel
	MessageID             int    // required if InlineMessageID is not specified
	InlineMessageID       string // required if ChatID|Username & MessageID are not specified
	Text                  string // required
	ParseMode             string
	Entities              []MessageEntity
	DisableWebPagePreview bool
	ReplyMarkup           interface{} // only InlineKeyboardMarkup TODO UPDATE 	ReplyMarkup     *InlineKeyboardMarkup
}

func (s *EditMessageText) Params() (Params, error) {
	params := make(Params, 8)

	if s.InlineMessageID == "" {
		err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
		if err != nil {
			return params, err
		}
		params.AddNonZero("message_id", s.MessageID)
	}
	params.AddNonEmpty("inline_message_id", s.InlineMessageID)
	params["text"] = s.Text
	params.AddNonEmpty("parse_mode", s.ParseMode)
	err := params.AddInterface("entities", s.Entities)
	if err != nil {
		return params, err
	}
	params.AddBool("disable_web_page_preview", s.DisableWebPagePreview)
	err = params.AddInterface("reply_markup", s.ReplyMarkup)

	return params, err
}
func (s *EditMessageText) EndPoint() string {
	return config.EndpointEditMessageText
}

// EditMessageCaption Edit captions of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
type EditMessageCaption struct {
	ChatID          int64  // required if InlineMessageID is not specified. use for chat|channel as int
	ChatIDStr       string // required if InlineMessageID is not specified. use for chat|channel as string
	Username        string // required if InlineMessageID is not specified. use for chat|channel
	MessageID       int    // required if InlineMessageID is not specified
	InlineMessageID string // required if ChatID|Username & MessageID are not specified
	Caption         string
	ParseMode       string
	CaptionEntities []MessageEntity
	ReplyMarkup     interface{} // only InlineKeyboardMarkup TODO UPDATE 	ReplyMarkup     *InlineKeyboardMarkup
}

func (s *EditMessageCaption) Params() (Params, error) {
	params := make(Params, 7)

	if s.InlineMessageID == "" {
		err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
		if err != nil {
			return params, err
		}
		params.AddNonZero("message_id", s.MessageID)
	}
	params.AddNonEmpty("inline_message_id", s.InlineMessageID)
	params.AddNonEmpty("caption", s.Caption)
	params.AddNonEmpty("parse_mode", s.ParseMode)
	err := params.AddInterface("caption_entities", s.CaptionEntities)
	if err != nil {
		return params, err
	}
	err = params.AddInterface("reply_markup", s.ReplyMarkup)

	return params, err
}
func (s *EditMessageCaption) EndPoint() string {
	return config.EndpointEditMessageCaption
}

// EditMessageMedia Edit animation, audio, document, photo, or video messages. If a message is part of a message album, then it can be edited only to an audio for audio albums, only to a document for document albums and to a photo or a video otherwise. When an inline message is edited, a new file can't be uploaded; use a previously uploaded file via its file_id or specify a URL. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
type EditMessageMedia struct {
	ChatID          int64       // required if InlineMessageID is not specified. use for chat|channel as int
	ChatIDStr       string      // required if InlineMessageID is not specified. use for chat|channel as string
	Username        string      // required if InlineMessageID is not specified. use for chat|channel
	MessageID       int         // required if InlineMessageID is not specified
	InlineMessageID string      // required if ChatID|Username & MessageID are not specified
	Media           interface{} // required
	ReplyMarkup     interface{} // only InlineKeyboardMarkup TODO UPDATE 	ReplyMarkup     *InlineKeyboardMarkup
}

func (s *EditMessageMedia) Params() (Params, error) {
	params := make(Params, 5)

	if s.InlineMessageID == "" {
		err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
		if err != nil {
			return params, err
		}
		params.AddNonZero("message_id", s.MessageID)
	}
	params.AddNonEmpty("inline_message_id", s.InlineMessageID)
	err := params.AddInterface("media", prepareInputMediaParam(s.Media, 0))
	if err != nil {
		return params, err
	}
	err = params.AddInterface("reply_markup", s.ReplyMarkup)

	return params, err
}
func (s *EditMessageMedia) Files() []RequestFile {
	return prepareInputMediaFile(s.Media, 0)
}
func (s *EditMessageMedia) EndPoint() string {
	return config.EndpointEditMessageMedia
}

// EditMessageReplyMarkup Edit only the reply markup of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
type EditMessageReplyMarkup struct {
	ChatID          int64       // required if InlineMessageID is not specified. use for chat|channel as int
	ChatIDStr       string      // required if InlineMessageID is not specified. use for chat|channel as string
	Username        string      // required if InlineMessageID is not specified. use for chat|channel
	MessageID       int         // required if InlineMessageID is not specified
	InlineMessageID string      // required if ChatID|Username & MessageID are not specified
	ReplyMarkup     interface{} // only InlineKeyboardMarkup TODO UPDATE 	ReplyMarkup     *InlineKeyboardMarkup
}

func (s *EditMessageReplyMarkup) Params() (Params, error) {
	params := make(Params, 4)

	if s.InlineMessageID == "" {
		err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
		if err != nil {
			return params, err
		}
		params.AddNonZero("message_id", s.MessageID)
	}
	params.AddNonEmpty("inline_message_id", s.InlineMessageID)
	err := params.AddInterface("reply_markup", s.ReplyMarkup)

	return params, err
}
func (s *EditMessageReplyMarkup) EndPoint() string {
	return config.EndpointEditMessageReplyMarkup
}

// StopPoll Stop a poll which was sent by the bot. On success, the stopped Poll is returned
type StopPoll struct {
	ChatID      int64       // required. use for chat|channel as int
	ChatIDStr   string      // required. use for chat|channel as string
	Username    string      // required. use for chat|channel
	MessageID   int         // required
	ReplyMarkup interface{} // only InlineKeyboardMarkup TODO UPDATE 	ReplyMarkup     *InlineKeyboardMarkup
}

func (s *StopPoll) Params() (Params, error) {
	params := make(Params, 3)

	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddNonZero("message_id", s.MessageID)
	err = params.AddInterface("reply_markup", s.ReplyMarkup)

	return params, err
}
func (s *StopPoll) EndPoint() string {
	return config.EndpointStopPoll
}

// DeleteMessage Delete a message, including service messages, with the following limitations. Returns True on success.
type DeleteMessage struct {
	ChatID    int64  // required. use for chat|channel as int
	ChatIDStr string // required. use for chat|channel as string
	Username  string // required. use for chat|channel
	MessageID int    // required
}

func (s *DeleteMessage) Params() (Params, error) {
	params := make(Params, 2)

	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddNonZero("message_id", s.MessageID)

	return params, nil
}
func (s *DeleteMessage) EndPoint() string {
	return config.EndpointDeleteMessage
}

// SendSticker Send static .WEBP, animated .TGS, or video .WEBM stickers. On success, the sent Message is returned.
type SendSticker struct {
	ChatID                   int64           // required. use for user|channel as int
	ChatIDStr                string          // required. use for user|channel as string
	Username                 string          // required. use for channel
	Sticker                  RequestFileData //required.
	DisableNotification      bool
	ProtectContent           bool
	ReplyToMessageID         int
	AllowSendingWithoutReply bool
	ReplyMarkup              interface{}
}

func (s *SendSticker) Params() (Params, error) {
	params := make(Params, 6)

	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params.AddBool("disable_notification", s.DisableNotification)
	params.AddBool("protect_content", s.ProtectContent)
	params.AddNonZero("reply_to_message_id", s.ReplyToMessageID)
	params.AddBool("allow_sending_without_reply", s.AllowSendingWithoutReply)
	err = params.AddInterface("reply_markup", s.ReplyMarkup)

	return params, nil
}
func (s *SendSticker) Files() []RequestFile {
	files := []RequestFile{{
		Name: "sticker",
		Data: s.Sticker,
	}}

	return files
}
func (s *SendSticker) EndPoint() string {
	return config.EndpointSendSticker
}

// GetStickerSet Get a sticker set. On success, a StickerSet object is returned.
type GetStickerSet struct {
	Name string // required
}

func (s *GetStickerSet) Params() (Params, error) {
	params := make(Params, 1)

	params["name"] = s.Name

	return params, nil
}
func (s *GetStickerSet) EndPoint() string {
	return config.EndpointGetStickerSet
}

// UploadStickerFile Upload a .PNG file with a sticker for later use in createNewStickerSet and addStickerToSet methods (can be used multiple times). Returns the uploaded File on success.
type UploadStickerFile struct {
	UserID     int64           // required
	PNGSticker RequestFileData // required
}

func (s *UploadStickerFile) Params() (Params, error) {
	params := make(Params, 1)

	params.AddNonZero64("user_id", s.UserID)

	return params, nil
}
func (s *UploadStickerFile) Files() []RequestFile {
	return []RequestFile{{
		Name: "png_sticker",
		Data: s.PNGSticker,
	}}
}
func (s *UploadStickerFile) EndPoint() string {
	return config.EndpointUploadStickerFile
}

// CreateNewStickerSet Create a new sticker set owned by a user. The bot will be able to edit the sticker set thus created. You must use exactly one of the fields png_sticker, tgs_sticker, or webm_sticker. Returns True on success.
type CreateNewStickerSet struct {
	UserID        int64  // required
	Name          string // required
	Title         string // required
	PNGSticker    RequestFileData
	TGSSticker    RequestFileData
	WEBMSticker   RequestFileData
	Emojis        string // required
	ContainsMasks bool
	MaskPosition  MaskPosition
}

func (s *CreateNewStickerSet) Params() (Params, error) {
	params := make(Params, 6)

	params.AddNonZero64("user_id", s.UserID)
	params["name"] = s.Name
	params["title"] = s.Title
	params["emojis"] = s.Emojis
	params.AddBool("contains_masks", s.ContainsMasks)
	err := params.AddInterface("mask_position", s.MaskPosition)

	return params, err
}
func (s *CreateNewStickerSet) Files() []RequestFile {
	if s.PNGSticker != nil {
		return []RequestFile{{
			Name: "png_sticker",
			Data: s.PNGSticker,
		}}
	}
	if s.TGSSticker != nil {
		return []RequestFile{{
			Name: "tgs_sticker",
			Data: s.TGSSticker,
		}}
	}
	if s.WEBMSticker != nil {
		return []RequestFile{{
			Name: "webm_sticker",
			Data: s.WEBMSticker,
		}}
	}

	return nil
}
func (s *CreateNewStickerSet) EndPoint() string {
	return config.EndpointCreateNewStickerSet
}

// AddStickerToSet Add a new sticker to a set created by the bot. You must use exactly one of the fields png_sticker, tgs_sticker, or webm_sticker. Animated stickers can be added to animated sticker sets and only to them. Animated sticker sets can have up to 50 stickers. Static sticker sets can have up to 120 stickers. Returns True on success.
type AddStickerToSet struct {
	UserID       int64  // required
	Name         string // required
	PNGSticker   RequestFileData
	TGSSticker   RequestFileData
	WEBMSticker  RequestFileData
	Emojis       string // required
	MaskPosition MaskPosition
}

func (s *AddStickerToSet) Params() (Params, error) {
	params := make(Params, 4)

	params.AddNonZero64("user_id", s.UserID)
	params["name"] = s.Name
	params["emojis"] = s.Emojis
	err := params.AddInterface("mask_position", s.MaskPosition)

	return params, err
}
func (s *AddStickerToSet) Files() []RequestFile {
	if s.PNGSticker != nil {
		return []RequestFile{{
			Name: "png_sticker",
			Data: s.PNGSticker,
		}}
	}
	if s.TGSSticker != nil {
		return []RequestFile{{
			Name: "tgs_sticker",
			Data: s.TGSSticker,
		}}
	}
	if s.WEBMSticker != nil {
		return []RequestFile{{
			Name: "webm_sticker",
			Data: s.WEBMSticker,
		}}
	}

	return nil
}
func (s *AddStickerToSet) EndPoint() string {
	return config.EndpointAddStickerToSet
}

// SetStickerPositionInSet Move a sticker in a set created by the bot to a specific position. Returns True on success.
type SetStickerPositionInSet struct {
	Sticker  string // required
	Position int    // required
}

func (s *SetStickerPositionInSet) Params() (Params, error) {
	params := make(Params, 2)

	params["sticker"] = s.Sticker
	params.AddNonZero("position", s.Position)

	return params, nil
}
func (s *SetStickerPositionInSet) EndPoint() string {
	return config.EndpointSetStickerPositionInSet
}

// SetStickerSetThumb Set the thumbnail of a sticker set. Animated thumbnails can be set for animated sticker sets only. Video thumbnails can be set only for video sticker sets only. Returns True on success.
type SetStickerSetThumb struct {
	Name   string
	UserID int64
	Thumb  RequestFileData
}

func (s *SetStickerSetThumb) Params() (Params, error) {
	params := make(Params, 2)

	params["name"] = s.Name
	params.AddNonZero64("user_id", s.UserID)

	return params, nil
}
func (s *SetStickerSetThumb) Files() []RequestFile {
	if s.Thumb != nil {
		return []RequestFile{{
			Name: "thumb",
			Data: s.Thumb,
		}}
	}

	return nil
}
func (s *SetStickerSetThumb) EndPoint() string {
	return config.EndpointSetStickerSetThumb
}

// DeleteStickerFromSet Delete a sticker from a set created by the bot. Returns True on success.
type DeleteStickerFromSet struct {
	Sticker string // required
}

func (s *DeleteStickerFromSet) Params() (Params, error) {
	params := make(Params, 1)

	params["sticker"] = s.Sticker

	return params, nil
}
func (s *DeleteStickerFromSet) EndPoint() string {
	return config.EndpointDeleteStickerFromSet
}

// AnswerInlineQuery Send answers to an inline query. On success, True is returned. No more than 50 results per query are allowed.
type AnswerInlineQuery struct {
	InlineQueryID     string        // required
	Results           []interface{} // required
	CacheTime         int
	IsPersonal        bool
	NextOffset        string
	SwitchPMText      string
	SwitchPMParameter string
}

func (s *AnswerInlineQuery) Params() (Params, error) {
	params := make(Params, 7)

	params["inline_query_id"] = s.InlineQueryID
	params.AddNonZero("cache_time", s.CacheTime)
	params.AddBool("is_personal", s.IsPersonal)
	params.AddNonEmpty("next_offset", s.NextOffset)
	params.AddNonEmpty("switch_pm_text", s.SwitchPMText)
	params.AddNonEmpty("switch_pm_parameter", s.SwitchPMParameter)
	err := params.AddInterface("results", s.Results)

	return params, err
}
func (s *AnswerInlineQuery) EndPoint() string {
	return config.EndpointAnswerInlineQuery
}

// AnswerWebAppQuery Set the result of an interaction with a Web App and send a corresponding message on behalf of the user to the chat from which the query originated. On success, a SentWebAppMessage object is returned.
type AnswerWebAppQuery struct {
	WebAppQueryID string      // required
	Result        interface{} // required
}

func (s *AnswerWebAppQuery) Params() (Params, error) {
	params := make(Params, 2)

	params["web_app_query_id"] = s.WebAppQueryID
	err := params.AddInterface("result", s.Result)

	return params, err
}
func (s *AnswerWebAppQuery) EndPoint() string {
	return config.EndpointAnswerWebAppQuery
}

// SendInvoice Send invoices. On success, the sent Message is returned.
type SendInvoice struct {
	ChatID                    int64          // required. use for user|channel as int
	ChatIDStr                 string         // required. use for user|channel as string
	Username                  string         // required. use for channel
	Title                     string         // required
	Description               string         // required
	Payload                   string         // required
	ProviderToken             string         // required
	Currency                  string         // required
	Prices                    []LabeledPrice // required
	MaxTipAmount              int
	SuggestedTipAmounts       []int
	StartParameter            string
	ProviderData              string
	PhotoURL                  string
	PhotoSize                 int
	PhotoWidth                int
	PhotoHeight               int
	NeedName                  bool
	NeedPhoneNumber           bool
	NeedEmail                 bool
	NeedShippingAddress       bool
	SendPhoneNumberToProvider bool
	SendEmailToProvider       bool
	IsFlexible                bool
	DisableNotification       bool
	ProtectContent            bool
	ReplyToMessageID          int
	AllowSendingWithoutReply  bool
	ReplyMarkup               interface{}
}

func (s *SendInvoice) Params() (Params, error) {
	params := make(Params, 27)

	err := params.AddFirstValid("chat_id", s.ChatID, s.ChatIDStr, s.Username)
	if err != nil {
		return params, err
	}
	params["title"] = s.Title
	params["description"] = s.Description
	params["payload"] = s.Payload
	params["provider_token"] = s.ProviderToken
	params["currency"] = s.Currency
	err = params.AddInterface("prices", s.Prices)
	if err != nil {
		return params, err
	}
	params.AddNonZero("max_tip_amount", s.MaxTipAmount)
	err = params.AddInterface("suggested_tip_amounts", s.SuggestedTipAmounts)
	if err != nil {
		return params, err
	}
	params.AddNonEmpty("start_parameter", s.StartParameter)
	params.AddNonEmpty("provider_data", s.ProviderData)
	params.AddNonEmpty("photo_url", s.PhotoURL)
	params.AddNonZero("photo_size", s.PhotoSize)
	params.AddNonZero("photo_width", s.PhotoWidth)
	params.AddNonZero("photo_height", s.PhotoHeight)
	params.AddBool("need_name", s.NeedName)
	params.AddBool("need_phone_number", s.NeedPhoneNumber)
	params.AddBool("need_email", s.NeedEmail)
	params.AddBool("need_shipping_address", s.NeedShippingAddress)
	params.AddBool("send_phone_number_to_provider", s.SendPhoneNumberToProvider)
	params.AddBool("send_email_to_provider", s.SendEmailToProvider)
	params.AddBool("is_flexible", s.IsFlexible)
	params.AddBool("disable_notification", s.DisableNotification)
	params.AddBool("protect_content", s.ProtectContent)
	params.AddNonZero("reply_to_message_id", s.ReplyToMessageID)
	params.AddBool("allow_sending_without_reply", s.AllowSendingWithoutReply)
	err = params.AddInterface("reply_markup", s.ReplyMarkup)
	return params, nil
}
func (s *SendInvoice) EndPoint() string {
	return config.EndpointSendInvoice
}

// CreateInvoiceLink Create a link for an invoice. Returns the created invoice link as String on success.
type CreateInvoiceLink struct {
	Title                     string          // required
	Description               string          // required
	Payload                   string          // required
	ProviderToken             string          // required
	Currency                  string          // required
	Prices                    []*LabeledPrice // required
	MaxTipAmount              int
	SuggestedTipAmounts       []int
	ProviderData              string
	PhotoURL                  string
	PhotoSize                 int
	PhotoWidth                int
	PhotoHeight               int
	NeedName                  bool
	NeedPhoneNumber           bool
	NeedEmail                 bool
	NeedShippingAddress       bool
	SendPhoneNumberToProvider bool
	SendEmailToProvider       bool
	IsFlexible                bool
}

func (s *CreateInvoiceLink) Params() (Params, error) {
	params := make(Params, 20)

	params["title"] = s.Title
	params["description"] = s.Description
	params["payload"] = s.Payload
	params["provider_token"] = s.ProviderToken
	params["currency"] = s.Currency
	err := params.AddInterface("prices", s.Prices)
	if err != nil {
		return params, err
	}
	params.AddNonZero("max_tip_amount", s.MaxTipAmount)
	err = params.AddInterface("suggested_tip_amounts", s.SuggestedTipAmounts)
	if err != nil {
		return params, err
	}
	params.AddNonEmpty("provider_data", s.ProviderData)
	params.AddNonEmpty("photo_url", s.PhotoURL)
	params.AddNonZero("photo_size", s.PhotoSize)
	params.AddNonZero("photo_width", s.PhotoWidth)
	params.AddNonZero("photo_height", s.PhotoHeight)
	params.AddBool("need_name", s.NeedName)
	params.AddBool("need_phone_number", s.NeedPhoneNumber)
	params.AddBool("need_email", s.NeedEmail)
	params.AddBool("need_shipping_address", s.NeedShippingAddress)
	params.AddBool("is_flexible", s.IsFlexible)
	params.AddBool("send_phone_number_to_provider", s.SendPhoneNumberToProvider)
	params.AddBool("send_email_to_provider", s.SendEmailToProvider)

	return params, nil
}
func (s *CreateInvoiceLink) EndPoint() string {
	return config.EndpointCreateInvoiceLink
}

// AnswerShippingQuery Reply to shipping queries. On success, True is returned.
type AnswerShippingQuery struct {
	ShippingQueryID string           // required
	OK              bool             // required
	ShippingOptions []ShippingOption // required if ok is True
	ErrorMessage    string           // required if ok is False
}

func (s *AnswerShippingQuery) Params() (Params, error) {
	params := make(Params, 4)

	params["shipping_query_id"] = s.ShippingQueryID
	params.AddBool("ok", s.OK)
	err := params.AddInterface("shipping_options", s.ShippingOptions)
	if err != nil {
		return params, err
	}
	params.AddNonEmpty("error_message", s.ErrorMessage)

	return params, nil
}
func (s *AnswerShippingQuery) EndPoint() string {
	return config.EndpointAnswerShippingQuery
}

// AnswerPreCheckoutQuery Respond to such pre-checkout queries. On success, True is returned. Note: The Bot API must receive an answer within 10 seconds after the pre-checkout query was sent.
type AnswerPreCheckoutQuery struct {
	PreCheckoutQueryID string // required
	OK                 bool   // required
	ErrorMessage       string // required if ok is False
}

func (s *AnswerPreCheckoutQuery) Params() (Params, error) {
	params := make(Params, 3)

	params["pre_checkout_query_id"] = s.PreCheckoutQueryID
	params.AddBool("ok", s.OK)
	params.AddNonEmpty("error_message", s.ErrorMessage)

	return params, nil
}
func (s *AnswerPreCheckoutQuery) EndPoint() string {
	return config.EndpointAnswerPreCheckoutQuery
}

// SetPassportDataErrors Informs a user that some of the Telegram Passport elements they provided contains errors. The user will not be able to re-submit their Passport to you until the errors are fixed (the contents of the field for which you returned the error must change). Returns True on success.
type SetPassportDataErrors struct {
	UserID int64         // required
	Errors []interface{} // required
}

func (s *SetPassportDataErrors) Params() (Params, error) {
	params := make(Params, 2)

	params.AddNonZero64("user_id", s.UserID)
	err := params.AddInterface("errors", s.Errors)

	return params, err
}
func (s *SetPassportDataErrors) EndPoint() string {
	return config.EndpointSetPassportDataErrors
}

// SendGame Send a game. On success, the sent Message is returned.
type SendGame struct {
	ChatID                   int64  // required
	GameShortName            string // required
	DisableNotification      bool
	ProtectContent           bool
	ReplyToMessageID         int
	AllowSendingWithoutReply bool
	ReplyMarkup              *InlineKeyboardMarkup
}

func (s *SendGame) Params() (Params, error) {
	params := make(Params, 7)

	params.AddNonZero64("chat_id", s.ChatID)
	params["game_short_name"] = s.GameShortName
	params.AddBool("disable_notification", s.DisableNotification)
	params.AddBool("protect_content", s.ProtectContent)
	params.AddNonZero("reply_to_message_id", s.ReplyToMessageID)
	params.AddBool("allow_sending_without_reply", s.AllowSendingWithoutReply)
	err := params.AddInterface("reply_markup", s.ReplyMarkup)
	return params, err
}
func (s *SendGame) EndPoint() string {
	return config.EndpointSendGame
}

// SetGameScore Set the score of the specified user in a game message. On success, if the message is not an inline message, the Message is returned, otherwise True is returned. Returns an error, if the new score is not greater than the user's current score in the chat and force is False.
type SetGameScore struct {
	UserID             int64 // required
	Score              int   // required
	Force              bool
	DisableEditMessage bool
	ChatID             int64  // required if inline_message_id is not specified
	MessageID          int    // required if inline_message_id is not specified
	InlineMessageID    string // required if chat_id and message_id are not specified
}

func (s *SetGameScore) Params() (Params, error) {
	params := make(Params, 5)

	params.AddNonZero64("user_id", s.UserID)
	params.AddNonZero("scrore", s.Score)
	params.AddBool("disable_edit_message", s.DisableEditMessage)
	if s.InlineMessageID != "" {
		params["inline_message_id"] = s.InlineMessageID
	} else {
		params.AddNonZero64("chat_id", s.ChatID)
		params.AddNonZero("message_id", s.MessageID)
	}

	return params, nil
}
func (s *SetGameScore) EndPoint() string {
	return config.EndpointSetGameScore
}

// GetGameHighScores Use this method to get data for high score tables. Will return the score of the specified user and several of their neighbors in a game. On success, returns an Array of GameHighScore objects. This method will currently return scores for the target user, plus two of their closest neighbors on each side. Will also return the top three users if the user and their neighbors are not among them. Please note that this behavior is subject to change.
type GetGameHighScores struct {
	UserID          int64  // required
	ChatID          int64  // required if inline_message_id is not specified
	MessageID       int    // required if inline_message_id is not specified
	InlineMessageID string // required if chat_id and message_id are not specified
}

func (s *GetGameHighScores) Params() (Params, error) {
	params := make(Params, 4)

	params.AddNonZero64("user_id", s.UserID)
	params.AddNonZero64("chat_id", s.ChatID)
	params.AddNonZero("message_id", s.MessageID)
	params.AddNonEmpty("inline_message_id", s.InlineMessageID)

	return params, nil
}
func (s *GetGameHighScores) EndPoint() string {
	return config.EndpointGetGameHighScores
}
