package telegram

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/raminsa/telegram-bot-api/config"
	"github.com/raminsa/telegram-bot-api/types"
)

// GetUpdates Use this method to receive incoming updates using long polling (wiki). An Array of Update objects is returned.
// Notes:
// 1. This method will not work if an outgoing webhook is set up.
// 2. In order to avoid getting duplicate updates, recalculate offset after each server response.
func (t *Api) GetUpdates(c *types.GetUpdates) ([]types.Update, error) {
	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var updates []types.Update
	err = json.Unmarshal(resp.Result, &updates)

	return updates, err
}

// GetUpdatesChan starts and returns a channel for getting updates.
func (t *Api) GetUpdatesChan(c *types.GetUpdates) types.UpdatesChannel {
	t.Bot.GetUpdateChannel = make(chan any)

	if c.Limit < 1 || c.Limit > 100 {
		c.Limit = 100
	}
	ch := make(chan types.Update, c.Limit)

	go func() {
		for {
			select {
			case <-t.Bot.GetUpdateChannel:
				close(ch)
				return
			default:
			}

			updates, err := t.GetUpdates(c)
			if err != nil {
				if t.Bot.Debug {
					t.WriteDebugLog(fmt.Sprintf("Failed to get updates, retrying in 3 seconds... Error:%s", err.Error()))
				}
				log.Println(err)
				log.Println("Failed to get updates, retrying in 3 seconds...")
				time.Sleep(time.Second * 3)

				continue
			}

			for _, update := range updates {
				if update.UpdateID >= c.Offset {
					c.Offset = update.UpdateID + 1
					ch <- update
				}
			}
		}
	}()

	return ch
}

// StopReceivingUpdates stops the go routine which receives updates
func (t *Api) StopReceivingUpdates() {
	close(t.Bot.GetUpdateChannel)
}

// SetWebhook Use this method to specify a URL and receive incoming updates via an outgoing webhook.
// Whenever there is an update for the bot, we will send an HTTPS POST request to the specified URL,
// containing a JSON-serialized Update.
// In case of an unsuccessful request, we will give up after a reasonable number of attempts.
// Returns True to success.
// If you'd like to make sure that you set the webhook,
// you can specify secret data in the parameter secret_token.
// If specified, the request will contain a header “X-Telegram-Bot-Api-Secret-Token” with the secret token as content.
func (t *Api) SetWebhook(c *types.SetWebhook) (*json.RawMessage, error) {
	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	return &resp.Result, err
}

// DeleteWebhook Use this method to remove webhook integration if you decide to switch back to getUpdates.
// Returns True to success.
func (t *Api) DeleteWebhook(c *types.DeleteWebhook) (*json.RawMessage, error) {
	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	return &resp.Result, err
}

// GetWebhook Use this method to get current webhook status.
// Requires no parameters.
// On success, returns a WebhookInfo object.
// If the bot is using getUpdates, will return an object with the url field empty.
func (t *Api) GetWebhook() (*types.WebhookInfo, error) {
	resp, err := t.MakeRequest(config.EndpointGetWebhook, nil)
	if err != nil {
		return nil, err
	}

	var webhookInfo types.WebhookInfo
	err = json.Unmarshal(resp.Result, &webhookInfo)

	return &webhookInfo, err
}

// GetMe A simple method for testing your bot authentication token.
// Requires no parameters.
// Returns basic information about the bot in the form of a User object.
func (t *Api) GetMe() (*types.User, error) {
	resp, err := t.MakeRequest(config.EndpointGetMe, nil)
	if err != nil {
		return nil, err
	}

	var user types.User
	err = json.Unmarshal(resp.Result, &user)

	if err == nil {
		user.IDString = fmt.Sprintf("%d", user.ID)
	}

	return &user, err
}

// LogOut Use this method to log out from the cloud Bot API server before launching the bot locally.
// You must log out the bot before running it locally,
// otherwise there is no guarantee that the bot will receive updates.
// After a successful call, you can immediately log in on a local server,
// but will not be able to log in back to the cloud Bot API server for 10 minutes.
// Returns True to success.
// Requires no parameters.
func (t *Api) LogOut() (bool, error) {
	resp, err := t.MakeRequest(config.EndpointLogOut, nil)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// Close Use this method to close the bot instance before moving it from one local server to another.
// You need to delete the webhook before calling this method to ensure that the
// bot isn't launched again after the server restarts.
// The method will return
// error 429 in the first 10 minutes after the bot is launched.
// Returns True to
// success.
// Requires no parameters.
func (t *Api) Close() (bool, error) {
	resp, err := t.MakeRequest(config.EndpointClose, nil)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// SendMessage Use this method to send text messages. On success, the sent Message is returned.
func (t *Api) SendMessage(c *types.SendMessage) (*types.Message, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}
	if c.Text == "" {
		return nil, errors.New("text Required")
	}

	return t.Send(c)
}

// ForwardMessage Use this method to forward messages of any kind.
// Service messages can't be forwarded.
// On success, the sent Message is returned.
func (t *Api) ForwardMessage(c *types.ForwardMessage) (*types.Message, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}
	if c.FromChatID == 0 && c.FromChatIDStr == "" && c.FromUsername == "" {
		return nil, errors.New("FromChatID or FromUsername Required")
	}
	if c.MessageID == 0 {
		return nil, errors.New("MessageID Required")
	}

	return t.Send(c)
}

// ForwardMessages Use this method to forward multiple messages of any kind.
// If some of the specified messages can't be found or forwarded, they are skipped.
// Service messages and messages with protected content can't be forwarded.
// Album grouping is kept for forwarded messages.
// On success, an array of MessageId of the sent messages is returned.
func (t *Api) ForwardMessages(c *types.ForwardMessages) ([]types.MessageID, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}
	if c.FromChatID == 0 && c.FromChatIDStr == "" && c.FromUsername == "" {
		return nil, errors.New("FromChatID or FromUsername Required")
	}
	if c.MessageIds == nil {
		return nil, errors.New("MessageIds Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var messages []types.MessageID
	err = json.Unmarshal(resp.Result, &messages)

	return messages, err
}

// CopyMessage Use this method to copy messages of any kind.
// Service messages and invoice messages can't be copied.
// The method is analogous to the method forwardMessage,
// but the copied message doesn't have a link to the original message.
// Returns the MessageId of the sent message on success.
func (t *Api) CopyMessage(c *types.CopyMessage) (*types.Message, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}
	if c.FromChatID == 0 && c.FromChatIDStr == "" && c.FromUsername == "" {
		return nil, errors.New("FromChatID or FromUsername Required")
	}
	if c.MessageID == 0 {
		return nil, errors.New("MessageID Required")
	}

	return t.Send(c)
}

// CopyMessages Use this method to copy messages of any kind.
// If some of the specified messages can't be found or copied, they are skipped.
// Service messages, giveaway messages, giveaway winners messages, and invoice messages can't be copied.
// A quiz poll can be copied only if the value of the field correct_option_id is known to the bot.
// The method is analogous to the method forwardMessages, but the copied messages don't have a link to the original message.
// Album grouping is kept for copied messages.
// On success, an array of MessageId of the sent messages is returned.
func (t *Api) CopyMessages(c *types.CopyMessages) ([]types.MessageID, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}
	if c.FromChatID == 0 && c.FromChatIDStr == "" && c.FromUsername == "" {
		return nil, errors.New("FromChatID or FromUsername Required")
	}
	if c.MessageIds == nil {
		return nil, errors.New("MessageIds Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var messages []types.MessageID
	err = json.Unmarshal(resp.Result, &messages)

	return messages, err
}

// SendPhoto Use this method to send photos. On success, the sent Message is returned.
func (t *Api) SendPhoto(c *types.SendPhoto) (*types.Message, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}
	if c.Photo == nil {
		return nil, errors.New("photo Required")
	}

	return t.Send(c)
}

// SendAudio Use this method to send audio files if you want Telegram clients to display them in the music player.
// Your audio must be in the .MP3 or .M4A format.
// On success, the sent Message is returned.
// Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
// For sending voice messages, use the sendVoice method instead.
func (t *Api) SendAudio(c *types.SendAudio) (*types.Message, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}
	if c.Audio == nil {
		return nil, errors.New("audio Required")
	}

	return t.Send(c)
}

// SendDocument Use this method to send general files.
// On success, the sent Message is returned.
// Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
func (t *Api) SendDocument(c *types.SendDocument) (*types.Message, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}
	if c.Document == nil {
		return nil, errors.New("document Required")
	}

	return t.Send(c)
}

// SendVideo Use this method to send video files,
// Telegram clients support MPEG4 videos (other formats may be sent as Document).
// On success, the sent Message is returned.
// Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
func (t *Api) SendVideo(c *types.SendVideo) (*types.Message, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}
	if c.Video == nil {
		return nil, errors.New("video Required")
	}

	return t.Send(c)
}

// SendAnimation Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without a sound).
// On success, the sent Message is returned.
// Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
func (t *Api) SendAnimation(c *types.SendAnimation) (*types.Message, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}
	if c.Animation == nil {
		return nil, errors.New("animation Required")
	}

	return t.Send(c)
}

// SendVoice Use this method
// to send audio files if you want Telegram clients to display the file as a playable voice message.
// For this to work, your audio must be in an .OGG file encoded with OPUS
// (other formats may be sent as Audio or Document).
// On success, the sent Message is returned.
// Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
func (t *Api) SendVoice(c *types.SendVoice) (*types.Message, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}
	if c.Voice == nil {
		return nil, errors.New("voice Required")
	}

	return t.Send(c)
}

// SendVideoNote As of v.4.0, Telegram clients support rounded square MPEG4 videos of up to 1 minute long.
// Use this method to send video messages.
// On success, the sent Message is returned.
func (t *Api) SendVideoNote(c *types.SendVideoNote) (*types.Message, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}
	if c.VideoNote == nil {
		return nil, errors.New("VideoNote Required")
	}

	return t.Send(c)
}

// SendMediaGroup Use this method to send a group of photos, videos, documents or audios as an album.
// Documents and audio files can be only grouped on an album with messages of the same type.
// On success, an array of Messages that were sent is returned.
func (t *Api) SendMediaGroup(c *types.SendMediaGroup) ([]types.Message, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}
	if c.Media == nil {
		return nil, errors.New("media Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var messages []types.Message
	err = json.Unmarshal(resp.Result, &messages)

	return messages, err
}

// SendLocation Use this method to send point on the map. On success, the sent Message is returned.
func (t *Api) SendLocation(c *types.SendLocation) (*types.Message, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}
	if c.Latitude == 0 {
		return nil, errors.New("latitude Required")
	}
	if c.Longitude == 0 {
		return nil, errors.New("longitude Required")
	}

	return t.Send(c)
}

// EditMessageLiveLocation Use this method to edit live location messages.
// A location can be edited
// until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation.
// On success, if the edited message is not an inline message, the edited Message is returned;
// otherwise True is returned.
func (t *Api) EditMessageLiveLocation(c *types.EditMessageLiveLocation) (*types.Message, error) {
	if c.InlineMessageID == "" {
		if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
			return nil, errors.New("ChatID or Username Required")
		}
		if c.MessageID == 0 {
			return nil, errors.New("MessageID Required")
		}
	}
	if c.Latitude == 0 {
		return nil, errors.New("latitude Required")
	}
	if c.Longitude == 0 {
		return nil, errors.New("longitude Required")
	}

	return t.Send(c)
}

// StopMessageLiveLocation Use this method to stop updating a live location message before live_period expires.
// On success, if the message is not an inline message, the edited Message is returned, otherwise True is returned.
func (t *Api) StopMessageLiveLocation(c *types.StopMessageLiveLocation) (*types.Message, error) {
	if c.InlineMessageID == "" {
		if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
			return nil, errors.New("ChatID or Username Required")
		}
		if c.MessageID == 0 {
			return nil, errors.New("MessageID Required")
		}
	}

	return t.Send(c)
}

// SendVenue Use this method to send information about a venue. On success, the sent Message is returned.
func (t *Api) SendVenue(c *types.SendVenue) (*types.Message, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}
	if c.Latitude == 0 {
		return nil, errors.New("latitude Required")
	}
	if c.Longitude == 0 {
		return nil, errors.New("longitude Required")
	}
	if c.Title == "" {
		return nil, errors.New("title Required")
	}
	if c.Address == "" {
		return nil, errors.New("address Required")
	}

	return t.Send(c)
}

// SendContact Use this method to send phone contacts. On success, the sent Message is returned.
func (t *Api) SendContact(c *types.SendContact) (*types.Message, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}
	if c.PhoneNumber == "" {
		return nil, errors.New("PhoneNumber Required")
	}
	if c.FirstName == "" {
		return nil, errors.New("FirstName Required")
	}

	return t.Send(c)
}

// SendPoll Use this method to send a native poll. On success, the sent Message is returned.
func (t *Api) SendPoll(c *types.SendPoll) (*types.Message, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}
	if c.Question == "" {
		return nil, errors.New("question Required")
	}
	if len(c.Options) == 0 {
		return nil, errors.New("options Required")
	}

	return t.Send(c)
}

// SendDice Use this method to send an animated emoji that will display a random value. On success, the sent Message is returned.
func (t *Api) SendDice(c *types.SendDice) (*types.Message, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}

	return t.Send(c)
}

// SendChatAction Use this method when you need to tell the user that something is happening on the bot side.
// The status is set for 5 seconds or less
// (when a message arrives from your bot, Telegram clients clear its typing status).
// Returns True on success.
// Example: The ImageBot needs some time to process a request and upload the image.
// Instead of sending a text message along the lines of “Retrieving image,
// please wait…”, the bot may use sendChatAction with action = upload_photo.
// The user will see a “sending photo” status for the bot.
// We only recommend using this method when a response from the bot will take a noticeable amount of time to arrive.
func (t *Api) SendChatAction(c *types.SendChatAction) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}
	if c.Action == "" {
		return false, errors.New("action Required")
	}

	_, err := t.Request(c)
	if err != nil {
		return false, err
	}

	return true, nil
}

// SetMessageReaction Use this method to change the chosen reactions on a message. Service messages can't be reacted to.
// Automatically forwarded messages from a channel to its discussion group have the same available reactions as messages in the channel.
// In albums, bots must react to the first message.
// Returns True on success.
func (t *Api) SetMessageReaction(c *types.SetMessageReaction) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}
	if c.MessageID == 0 {
		return false, errors.New("MessageID Required")
	}

	_, err := t.Request(c)
	if err != nil {
		return false, err
	}

	return true, nil
}

// GetUserProfilePhotos Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos object.
func (t *Api) GetUserProfilePhotos(c *types.GetUserProfilePhotos) (*types.UserProfilePhotos, error) {
	if c.UserID == 0 {
		return nil, errors.New("UserID Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var profilePhotos types.UserProfilePhotos
	err = json.Unmarshal(resp.Result, &profilePhotos)

	return &profilePhotos, err
}

// GetFile Use this method to get basic information about a file and prepare it for downloading.
// For the moment, bots can download files of up to 20MB in size.
// On success, a File object is returned.
// The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>,
// where <file_path> is taken from the response.
// It is guaranteed that the link will be valid for at least 1 hour.
// When the link expires, a new one can be requested by calling getFile again.
// Note: This function may not preserve the original file name and MIME type.
// You should save the file's MIME type and name (if available) when the File object is received.
func (t *Api) GetFile(c *types.GetFile) (*types.File, error) {
	if c.FileID == "" {
		return nil, errors.New("FileID Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var file types.File
	err = json.Unmarshal(resp.Result, &file)

	return &file, err
}

// BanChatMember Use this method to ban a user in a group, a supergroup or a channel.
// In the case of supergroups and channels,
// the user will not be able to return to the chat on their own using invite links, etc.,
// unless unbanned first.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns True to success.
func (t *Api) BanChatMember(c *types.BanChatMember) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}
	if c.UserID == 0 {
		return false, errors.New("UserID Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// UnbanChatMember Use this method to unban a previously banned user in a supergroup or channel.
// The user will not return to the group or channel automatically,
// but will be able to join via a link, etc. The bot must be an administrator for this to work.
// By default, this method guarantees that after the call the user is not a member of the chat,
// but will be able to join it.
// So if the user is a member of the chat, they will also be removed from the chat.
// If you don't want this, use the parameter only_if_banned.
// Returns True to success.
func (t *Api) UnbanChatMember(c *types.UnbanChatMember) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}
	if c.UserID == 0 {
		return false, errors.New("user_id Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// RestrictChatMember Use this method to restrict a user in a supergroup.
// The bot must be an administrator in the supergroup for this to work
// and must have the appropriate administrator rights.
// Pass True for all permissions to lift restrictions from a user.
// Returns True to success.
func (t *Api) RestrictChatMember(c *types.RestrictChatMember) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}
	if c.UserID == 0 {
		return false, errors.New("user_id Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// PromoteChatMember Use this method to promote or demote a user in a supergroup or a channel.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Pass False for all boolean parameters to demote a user.
// Returns True to success.
func (t *Api) PromoteChatMember(c *types.PromoteChatMember) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}
	if c.UserID == 0 {
		return false, errors.New("user_id Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// SetChatAdministratorCustomTitle Use this method
// to set a custom title for an administrator in a supergroup promoted by the bot.
// Returns True to success.
func (t *Api) SetChatAdministratorCustomTitle(c *types.SetChatAdministratorCustomTitle) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}
	if c.UserID == 0 {
		return false, errors.New("user_id Required")
	}
	if c.CustomTitle == "" {
		return false, errors.New("custom_title Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// BanChatSenderChat Use this method to ban a channel chat in a supergroup or a channel.
// Until the chat is unbanned,
// the owner of the banned chat won't be able to send messages on behalf of any of their channels.
// The bot must be an administrator in the supergroup or channel for this to work
// and must have the appropriate administrator rights.
// Returns True to success.
func (t *Api) BanChatSenderChat(c *types.BanChatSenderChat) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}
	if c.SenderChatID == 0 {
		return false, errors.New("sender_chatID Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// UnbanChatSenderChat Use this method to unban a previously banned channel chat in a supergroup or channel.
// The bot must be an administrator for this to work and must have the appropriate administrator rights.
// Returns True to success.
func (t *Api) UnbanChatSenderChat(c *types.UnbanChatSenderChat) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}
	if c.SenderChatID == 0 {
		return false, errors.New("sender_chatID Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// SetChatPermissions Use this method to set default chat permissions for all members.
// The bot must be an administrator in the group or a supergroup
// for this to work and must have the can_restrict_members administrator rights.
// Returns True to success.
func (t *Api) SetChatPermissions(c *types.SetChatPermissions) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// ExportChatInviteLink Use this method to generate a new primary invite link for a chat;
// any previously generated primary link is revoked.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns the new invite link as String on success.
func (t *Api) ExportChatInviteLink(c *types.ExportChatInviteLink) (string, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return "", errors.New("ChatID or Username Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return "", err
	}

	var inviteLink string
	err = json.Unmarshal(resp.Result, &inviteLink)

	return inviteLink, err
}

// CreateChatInviteLink Use this method to create an additional invite link for a chat.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// The link can be revoked using the method revokeChatInviteLink.
// Returns the new invite link as ChatInviteLink object.
func (t *Api) CreateChatInviteLink(c *types.CreateChatInviteLink) (*types.ChatInviteLink, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var chatInviteLink types.ChatInviteLink
	err = json.Unmarshal(resp.Result, &chatInviteLink)

	return &chatInviteLink, err
}

// EditChatInviteLink Use this method to edit a non-primary invite link created by the bot.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns the edited invite link as a ChatInviteLink object.
func (t *Api) EditChatInviteLink(c *types.EditChatInviteLink) (*types.ChatInviteLink, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}
	if c.InviteLink == "" {
		return nil, errors.New("invite_link Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var chatInviteLink types.ChatInviteLink
	err = json.Unmarshal(resp.Result, &chatInviteLink)

	return &chatInviteLink, err
}

// RevokeChatInviteLink Use this method to revoke an invitation link created by the bot.
// If the primary link is revoked, a new link is automatically generated.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns the revoked invite link as ChatInviteLink object.
func (t *Api) RevokeChatInviteLink(c *types.RevokeChatInviteLink) (*types.ChatInviteLink, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}
	if c.InviteLink == "" {
		return nil, errors.New("invite_link Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var chatInviteLink types.ChatInviteLink
	err = json.Unmarshal(resp.Result, &chatInviteLink)

	return &chatInviteLink, err
}

// ApproveChatJoinRequest Use this method to approve a chat join request.
// The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right.
// Returns True to success.
func (t *Api) ApproveChatJoinRequest(c *types.ApproveChatJoinRequest) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}
	if c.UserID == 0 {
		return false, errors.New("user_id Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// DeclineChatJoinRequest Use this method to decline a chat join request.
// The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right.
// Returns True to success.
func (t *Api) DeclineChatJoinRequest(c *types.DeclineChatJoinRequest) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}
	if c.UserID == 0 {
		return false, errors.New("user_id Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// SetChatPhoto Use this method to set a new profile photo for the chat.
// Photos can't be changed for private chats.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns True to success.
func (t *Api) SetChatPhoto(c *types.SetChatPhoto) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}
	if c.Photo == nil {
		return false, errors.New("photo Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// DeleteChatPhoto Use this method to delete a chat photo.
// Photos can't be changed for private chats.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns True to success.
func (t *Api) DeleteChatPhoto(c *types.DeleteChatPhoto) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// SetChatTitle Use this method to change the title of a chat.
// Titles can't be changed for private chats.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns True to success.
func (t *Api) SetChatTitle(c *types.SetChatTitle) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}
	if c.Title == "" {
		return false, errors.New("title Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// SetChatDescription Use this method to change the description of a group, a supergroup or a channel.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns True to success.
func (t *Api) SetChatDescription(c *types.SetChatDescription) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// PinChatMessage Use this method to add a message to the list of pinned messages in a chat.
// If the chat is not a private chat,
// the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages'
// administrator right in a supergroup or 'can_edit_messages' administrator right in a channel.
// Returns True to success.
func (t *Api) PinChatMessage(c *types.PinChatMessage) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}
	if c.MessageID == 0 {
		return false, errors.New("MessageID Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// UnpinChatMessage Use this method to remove a message from the list of pinned messages in a chat.
// If the chat is not a private chat,
// the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages'
// administrator right in a supergroup or 'can_edit_messages' administrator right in a channel.
// Returns True to success.
func (t *Api) UnpinChatMessage(c *types.UnpinChatMessage) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// UnpinAllChatMessages Use this method to clear the list of pinned messages in a chat.
// If the chat is not a private chat,
// the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages'
// administrator right in a supergroup or 'can_edit_messages' administrator right in a channel.
// Returns True to success.
func (t *Api) UnpinAllChatMessages(c *types.UnpinAllChatMessages) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// LeaveChat Use this method for your bot to leave a group, supergroup or channel. Returns True to success.
func (t *Api) LeaveChat(c *types.LeaveChat) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// GetChat Use this method to get up-to-date information about the chat
// (current name of the user for one-on-one conversations, current username of a user, group or channel, etc.).
// Returns a Chat object on success.
func (t *Api) GetChat(c *types.GetChat) (*types.ChatFullInfo, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var chat types.ChatFullInfo
	err = json.Unmarshal(resp.Result, &chat)

	if err == nil {
		chat.IDString = fmt.Sprintf("%d", chat.ID)
	}

	return &chat, err
}

// GetChatAdministrators Use this method to get a list of administrators in a chat.
// On success,
// returns an Array of ChatMember objects that contains information about all chat administrators except other bots.
// If the chat is a group or a supergroup and no administrators were appointed, only the creator will be returned.
func (t *Api) GetChatAdministrators(c *types.GetChatAdministrators) ([]types.ChatMember, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var members []types.ChatMember
	err = json.Unmarshal(resp.Result, &members)

	return members, err
}

// GetChatMemberCount Use this method to get the number of members in a chat. Returns Int to success.
func (t *Api) GetChatMemberCount(c *types.GetChatMemberCount) (int64, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return 0, errors.New("ChatID or Username Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return 0, err
	}

	var count int64
	err = json.Unmarshal(resp.Result, &count)

	return count, err
}

// GetChatMember Use this method to get information about a member of a chat. Returns a ChatMember object on success.
func (t *Api) GetChatMember(c *types.GetChatMember) (*types.ChatMember, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}
	if c.UserID == 0 {
		return nil, errors.New("UserID Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var member types.ChatMember
	err = json.Unmarshal(resp.Result, &member)

	return &member, err
}

// SetChatStickerSet Use this method to set a new group sticker set for a supergroup.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method.
// Returns True to success.
func (t *Api) SetChatStickerSet(c *types.SetChatStickerSet) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}
	if c.StickerSetName == "" {
		return false, errors.New("StickerSetName Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// DeleteChatStickerSet Use this method to delete a group sticker set from a supergroup.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method.
// Returns True to success.
func (t *Api) DeleteChatStickerSet(c *types.DeleteChatStickerSet) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// GetForumTopicIconStickers Use this method to get custom emoji stickers,
// which can be used as a forum topic icon by any user.
// Requires no parameters.
// Returns an Array of Sticker objects.
func (t *Api) GetForumTopicIconStickers(c *types.GetForumTopicIconStickers) ([]types.Sticker, error) {
	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var result []types.Sticker
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return nil, err
	}

	return result, err
}

// CreateForumTopic Use this method to create a topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work
// and must have the can_manage_topics administrator rights.
// Returns information about the created topic as a ForumTopic object.
func (t *Api) CreateForumTopic(c *types.CreateForumTopic) (*types.ForumTopic, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}

	if c.Name == "" {
		return nil, errors.New("name is Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var result types.ForumTopic
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return nil, err
	}

	return &result, err
}

// EditForumTopic Use this method to edit the name and icon of a topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have can_manage_topics administrator rights,
// unless it is the creator of the topic.
// Returns True to success.
func (t *Api) EditForumTopic(c *types.EditForumTopic) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}

	if c.MessageThreadID == 0 {
		return false, errors.New("message_thread_id is Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// CloseForumTopic Use this method to close an open topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work
// and must have the can_manage_topics administrator rights,
// unless it is the creator of the topic.
// Returns True to success.
func (t *Api) CloseForumTopic(c *types.CloseForumTopic) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}

	if c.MessageThreadID == 0 {
		return false, errors.New("message_thread_id is Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// ReopenForumTopic Use this method to reopen a closed topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work
// and must have the can_manage_topics administrator rights,
// unless it is the creator of the topic.
// Returns True to success.
func (t *Api) ReopenForumTopic(c *types.ReopenForumTopic) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}

	if c.MessageThreadID == 0 {
		return false, errors.New("message_thread_id is Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// DeleteForumTopic Use this method to delete a forum topic along with all its messages in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work
// and must have the can_delete_messages administrator rights.
// Returns True to success.
func (t *Api) DeleteForumTopic(c *types.DeleteForumTopic) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}

	if c.MessageThreadID == 0 {
		return false, errors.New("message_thread_id is Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// UnpinAllForumTopicMessages Use this method to clear the list of pinned messages in a forum topic.
// The bot must be an administrator in the chat for this to work and must have
// the can_pin_messages administrator right in the supergroup.
// Returns True to
// success.
func (t *Api) UnpinAllForumTopicMessages(c *types.UnpinAllForumTopicMessages) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}

	if c.MessageThreadID == 0 {
		return false, errors.New("message_thread_id is Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// EditGeneralForumTopic Use this method to edit the name of the 'General' topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have can_manage_topics administrator rights.
// Returns True to success.
func (t *Api) EditGeneralForumTopic(c *types.EditGeneralForumTopic) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}

	if c.Name == "" {
		return false, errors.New("name is Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// CloseGeneralForumTopic Use this method to close an open 'General' topic in a
// forum supergroup chat.
// The bot must be an administrator in the chat for this to work
// and must have the can_manage_topics administrator rights.
// Returns True to success.
func (t *Api) CloseGeneralForumTopic(c *types.CloseGeneralForumTopic) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// ReopenGeneralForumTopic Use this method to reopen a closed 'General' topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work
// and must have the can_manage_topics administrator rights.
// The topic will be automatically unhidden if it was hidden.
// Returns True to success.
func (t *Api) ReopenGeneralForumTopic(c *types.ReopenGeneralForumTopic) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// HideGeneralForumTopic Use this method to hide the 'General' topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work
// and must have the can_manage_topics administrator rights.
// The topic will be automatically closed if it is open.
// Returns True to success.
func (t *Api) HideGeneralForumTopic(c *types.HideGeneralForumTopic) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// UnHideGeneralForumTopic Use this method to unhide the 'General' topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work
// and must have the can_manage_topics administrator rights.
// Returns True to success.
func (t *Api) UnHideGeneralForumTopic(c *types.UnHideGeneralForumTopic) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// UnpinAllGeneralForumTopicMessages Use this method to clear the list of pinned messages in a General forum topic.
// The bot must be an administrator in the chat for this to work
// and must have the can_pin_messages administrator right in the supergroup.
// Returns True to success.
func (t *Api) UnpinAllGeneralForumTopicMessages(c *types.UnpinAllGeneralForumTopicMessages) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// AnswerCallbackQuery Use this method to send answers to callback queries sent from inline keyboards.
// The answer will be displayed to the user as a notification at the top of the chat screen or as an alert.
// On success, True is returned.
// Alternatively, the user can be redirected to the specified Game URL.
// For this option to work, you must first create a game for your bot via @BotFather and accept the terms.
// Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
func (t *Api) AnswerCallbackQuery(c *types.AnswerCallbackQuery) (bool, error) {
	if c.CallbackQueryID == "" {
		return false, errors.New("CallbackQueryID Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// GetUserChatBoosts Use this method to get the list of boosts added to a chat by a user. Requires administrator rights in the chat. Returns a UserChatBoosts object.
func (t *Api) GetUserChatBoosts(c *types.GetUserChatBoosts) (*types.UserChatBoosts, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}
	if c.UserID == 0 {
		return nil, errors.New("UserID Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var chatBoosts types.UserChatBoosts
	err = json.Unmarshal(resp.Result, &chatBoosts)

	return &chatBoosts, err
}

// GetBusinessConnection Use this method to get information about the connection of the bot with a business account.
// Returns a BusinessConnection object on success.
func (t *Api) GetBusinessConnection(c *types.GetBusinessConnection) (*types.BusinessConnection, error) {
	if c.BusinessConnectionId == "" {
		return nil, errors.New("BusinessConnectionId Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var businessConnection types.BusinessConnection
	err = json.Unmarshal(resp.Result, &businessConnection)

	return &businessConnection, err
}

// SetMyCommands Use this method to change the list of the bot commands.
// See https://core.telegram.org/bots#commands for more details about bot commands.
// Returns True to success.
func (t *Api) SetMyCommands(c *types.SetMyCommands) (bool, error) {
	if c.Commands == nil {
		return false, errors.New("commands Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// DeleteMyCommands Use this method to change the list of the bot commands.
// See https://core.telegram.org/bots#commands for more details about bot commands.
// Returns True to success.
func (t *Api) DeleteMyCommands(c *types.DeleteMyCommands) (bool, error) {
	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// GetMyCommands Use this method to get the current list of the bot commands for the given scope and user language.
// Returns Array of BotCommand on success.
// If commands aren't set, an empty list is returned.
func (t *Api) GetMyCommands(c *types.GetMyCommands) ([]types.BotCommand, error) {
	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var commands []types.BotCommand
	err = json.Unmarshal(resp.Result, &commands)

	return commands, err
}

// SetMyName Use this method to change the bot name.
// Returns True to success.
func (t *Api) SetMyName(c *types.SetMyName) (bool, error) {
	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)

	return result, err
}

// GetMyName Use this method to get the current bot name for the given user language. Returns BotName on success.
func (t *Api) GetMyName(c *types.GetMyName) (*types.BotName, error) {
	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var result types.BotName
	err = json.Unmarshal(resp.Result, &result)

	return &result, err
}

// SetMyDescription Use this method to change the bot description,
// which is shown in the chat with the bot if the chat is empty.
// Returns True to success.
func (t *Api) SetMyDescription(c *types.SetMyDescription) (bool, error) {
	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// GetMyDescription Use this method to get the current bot description for the given user language.
// Returns BotDescription on success.
func (t *Api) GetMyDescription(c *types.GetMyDescription) (*types.BotDescription, error) {
	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var result types.BotDescription
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return nil, err
	}

	return &result, err
}

// SetMyShortDescription Use this method to change the bot short description,
// which is shown on the bot profile page and is sent together with the link when users share the bot.
// Returns True to success.
func (t *Api) SetMyShortDescription(c *types.SetMyShortDescription) (bool, error) {
	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// GetMyShortDescription Use this method to get the current bot short description for the given user language.
// Returns BotShortDescription on success.
func (t *Api) GetMyShortDescription(c *types.GetMyShortDescription) (*types.BotShortDescription, error) {
	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var result types.BotShortDescription
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return nil, err
	}

	return &result, err
}

// SetChatMenuButton Use this method to change the bot menu button in a private chat, or the default menu button.
// Returns True to success.
func (t *Api) SetChatMenuButton(c *types.SetChatMenuButton) (bool, error) {
	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// GetChatMenuButton Use this method to get the current value of the bot menu button in a private chat,
// or the default menu button.
// Returns MenuButton on success.
func (t *Api) GetChatMenuButton(c *types.GetChatMenuButton) (*types.MenuButtons, error) {
	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var menuButton types.MenuButtons
	err = json.Unmarshal(resp.Result, &menuButton)

	return &menuButton, err
}

// SetMyDefaultAdministratorRights Use this method to change the default administrator rights requested by the bot
// when it's added as an administrator to groups or channels.
// These rights will be suggested to users, but they are free to modify the list before adding the bot.
// Returns True to success.
func (t *Api) SetMyDefaultAdministratorRights(c *types.SetMyDefaultAdministratorRights) (bool, error) {
	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// GetMyDefaultAdministratorRights Use this method to get the current default administrator rights of the bot.
// Returns ChatAdministratorRights on success.
func (t *Api) GetMyDefaultAdministratorRights(c *types.GetMyDefaultAdministratorRights) (bool, error) {
	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// EditMessageText Use this method to edit text and game messages.
// On success, if the edited message is not an inline message, the edited Message is returned;
// otherwise True is returned.
func (t *Api) EditMessageText(c *types.EditMessageText) (*types.Message, error) {
	if c.InlineMessageID == "" {
		if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
			return nil, errors.New("ChatID or Username Required")
		}
		if c.MessageID == 0 {
			return nil, errors.New("MessageID Required")
		}
	}

	if c.Text == "" {
		return nil, errors.New("text Required")
	}

	return t.Send(c)
}

// EditInlineMessageText Use this method to edit text and game messages.
// On success, True is returned.
func (t *Api) EditInlineMessageText(c *types.EditMessageText) (bool, error) {
	if c.InlineMessageID == "" {
		if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
			return false, errors.New("ChatID or Username Required")
		}
		if c.MessageID == 0 {
			return false, errors.New("MessageID Required")
		}
	}

	if c.Text == "" {
		return false, errors.New("text Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// EditMessageCaption Use this method to edit captions of messages.
// On success, the edited Message is returned.
func (t *Api) EditMessageCaption(c *types.EditMessageCaption) (*types.Message, error) {
	if c.InlineMessageID == "" {
		if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
			return nil, errors.New("ChatID or Username Required")
		}
		if c.MessageID == 0 {
			return nil, errors.New("MessageID Required")
		}
	}

	return t.Send(c)
}

// EditInlineMessageCaption Use this method to edit captions of messages.
// On success, the edited Message is returned.
func (t *Api) EditInlineMessageCaption(c *types.EditMessageCaption) (bool, error) {
	if c.InlineMessageID == "" {
		if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
			return false, errors.New("ChatID or Username Required")
		}
		if c.MessageID == 0 {
			return false, errors.New("MessageID Required")
		}
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// EditMessageMedia Use this method to edit animation, audio, document, photo, or video messages.
// If a message is part of a message album, then it can be edited only to an audio for audio albums,
// only to a document for document albums and to a photo or a video otherwise.
// When an inline message is edited, a new file can't be uploaded;
// use a previously uploaded file via its file_id or specify a URL.
// On success, the edited Message is returned.
func (t *Api) EditMessageMedia(c *types.EditMessageMedia) (*types.Message, error) {
	if c.InlineMessageID == "" {
		if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
			return nil, errors.New("ChatID or Username Required")
		}
		if c.MessageID == 0 {
			return nil, errors.New("MessageID Required")
		}
	}
	if c.Media == nil {
		return nil, errors.New("media Required")
	}

	return t.Send(c)
}

// EditInlineMessageMedia Use this method to edit animation, audio, document, photo, or video messages.
// If a message is part of a message album, then it can be edited only to an audio for audio albums,
// only to a document for document albums and to a photo or a video otherwise.
// When an inline message is edited, a new file can't be uploaded;
// use a previously uploaded file via its file_id or specify a URL.
// On success, True is returned.
func (t *Api) EditInlineMessageMedia(c *types.EditMessageMedia) (bool, error) {
	if c.InlineMessageID == "" {
		if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
			return false, errors.New("ChatID or Username Required")
		}
		if c.MessageID == 0 {
			return false, errors.New("MessageID Required")
		}
	}
	if c.Media == nil {
		return false, errors.New("media Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// EditMessageReplyMarkup Use this method to edit only the reply markup of messages.
// On success, the edited Message is returned.
func (t *Api) EditMessageReplyMarkup(c *types.EditMessageReplyMarkup) (*types.Message, error) {
	if c.InlineMessageID == "" {
		if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
			return nil, errors.New("ChatID or Username Required")
		}
		if c.MessageID == 0 {
			return nil, errors.New("MessageID Required")
		}
	}

	return t.Send(c)
}

// EditInlineMessageReplyMarkup Use this method to edit only the reply markup of messages.
// On success, True is returned.
func (t *Api) EditInlineMessageReplyMarkup(c *types.EditMessageReplyMarkup) (bool, error) {
	if c.InlineMessageID == "" {
		if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
			return false, errors.New("ChatID or Username Required")
		}
		if c.MessageID == 0 {
			return false, errors.New("MessageID Required")
		}
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// StopPoll Use this method to stop a poll which was sent by the bot.
// On success, the stopped Poll is returned.
func (t *Api) StopPoll(c *types.StopPoll) (*types.Poll, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}
	if c.MessageID == 0 {
		return nil, errors.New("MessageID Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var poll types.Poll
	err = json.Unmarshal(resp.Result, &poll)

	return &poll, err
}

// DeleteMessage Use this method to delete a message, including service messages, with the following limitations:
// - A message can only be deleted if it was sent less than 48 hours ago.
// - A die message in a private chat can only be deleted if it was sent more than 24 hours ago.
// - Bots can delete outgoing messages in private chats, groups, and supergroups.
// - Bots can delete incoming messages in private chats.
// - Bots granted can_post_messages permissions can delete outgoing messages in channels.
// - If the bot is an administrator of a group, it can delete any message there.
// - If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.
// Returns True to success.
func (t *Api) DeleteMessage(c *types.DeleteMessage) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}
	if c.MessageID == 0 {
		return false, errors.New("MessageID Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// DeleteMessages Use this method to delete multiple messages simultaneously.
// If some of the specified messages can't be found, they are skipped.
// Returns True on success.
func (t *Api) DeleteMessages(c *types.DeleteMessages) (bool, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return false, errors.New("ChatID or Username Required")
	}
	if c.MessageIds == nil {
		return false, errors.New("MessageIds Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// SendSticker Use this method to send static .WEBP, animated .TGS, or video .WEBM stickers.
// On success, the sent Message is returned.
func (t *Api) SendSticker(c *types.SendSticker) (*types.Message, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}
	if c.Sticker == nil {
		return nil, errors.New("MessageID Required")
	}

	return t.Send(c)
}

// GetStickerSet Use this method to get a sticker set. On success, a StickerSet object is returned.
func (t *Api) GetStickerSet(c *types.GetStickerSet) (*types.StickerSet, error) {
	if c.Name == "" {
		return nil, errors.New("name Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var stickerSet types.StickerSet
	err = json.Unmarshal(resp.Result, &stickerSet)

	return &stickerSet, err
}

// GetCustomEmojiStickers Use this method to get information about custom emoji stickers by their identifiers.
// Returns an Array of Sticker objects.
func (t *Api) GetCustomEmojiStickers(c *types.GetCustomEmojiStickers) ([]types.Sticker, error) {
	if len(c.CustomEmojiIds) < 1 {
		return nil, errors.New("customEmojiIds Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var stickerSet []types.Sticker
	err = json.Unmarshal(resp.Result, &stickerSet)

	return stickerSet, err
}

// UploadStickerFile Use this method to upload a file with a sticker for later use in the createNewStickerSet
// and addStickerToSet methods (the file can be used multiple times).
// Returns the uploaded File on success.
func (t *Api) UploadStickerFile(c *types.UploadStickerFile) (*types.File, error) {
	if c.UserID == 0 {
		return nil, errors.New("UserID Required")
	}
	if c.Sticker == nil {
		return nil, errors.New("sticker Required")
	}
	if c.StickerFormat == "" {
		return nil, errors.New("stickerFormat Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var file types.File
	err = json.Unmarshal(resp.Result, &file)

	return &file, err
}

// CreateNewStickerSet Use this method to create a new sticker set owned by a user.
// The bot will be able to edit the sticker set thus created.
// Returns True to success.
func (t *Api) CreateNewStickerSet(c *types.CreateNewStickerSet) (bool, error) {
	if c.UserID == 0 {
		return false, errors.New("UserID Required")
	}
	if c.Name == "" {
		return false, errors.New("name Required")
	}
	if c.Title == "" {
		return false, errors.New("title Required")
	}
	if len(c.Stickers) < 1 {
		return false, errors.New("stickers Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// AddStickerToSet Use this method to add a new sticker to a set created by the bot.
// The format of the added sticker must match the format of the other stickers in the set.
// Emoji sticker sets can have up to 200 stickers.
// Animated and video sticker sets can have up to 50 stickers.
// Static sticker sets can have up to 120 stickers.
// Returns True to success.
func (t *Api) AddStickerToSet(c *types.AddStickerToSet) (bool, error) {
	if c.UserID == 0 {
		return false, errors.New("UserID Required")
	}
	if c.Name == "" {
		return false, errors.New("name Required")
	}
	if c.Sticker.Sticker == nil {
		return false, errors.New("sticker Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// SetStickerPositionInSet Use this method to move a sticker in a set created by the bot to a specific position.
// Returns True to success.
func (t *Api) SetStickerPositionInSet(c *types.SetStickerPositionInSet) (bool, error) {
	if c.Sticker == "" {
		return false, errors.New("sticker Required")
	}
	if c.Position == 0 {
		return false, errors.New("position Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// DeleteStickerFromSet Use this method to delete a sticker from a set created by the bot.
// Returns True to success.
func (t *Api) DeleteStickerFromSet(c *types.DeleteStickerFromSet) (bool, error) {
	if c.Sticker == "" {
		return false, errors.New("sticker Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// ReplaceStickerInSet Use this method to replace an existing sticker in a sticker set with a new one. The method is equivalent to calling deleteStickerFromSet, then addStickerToSet, then setStickerPositionInSet.
// Returns True on success
func (t *Api) ReplaceStickerInSet(c *types.ReplaceStickerInSet) (bool, error) {
	if c.UserID == 0 {
		return false, errors.New("UserID Required")
	}
	if c.Name == "" {
		return false, errors.New("name Required")
	}
	if c.OldSticker == "" {
		return false, errors.New("OldSticker Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// SetStickerEmojiList Use this method to change the list of emoji assigned to a regular or custom emoji sticker.
// The sticker must belong to a sticker set created by the bot.
// Returns True to success.
func (t *Api) SetStickerEmojiList(c *types.SetStickerEmojiList) (bool, error) {
	if c.Sticker == "" {
		return false, errors.New("sticker Required")
	}
	if len(c.EmojiList) < 1 {
		return false, errors.New("emojiList Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// SetStickerKeywords Use this method to change search keywords assigned to a regular or custom emoji sticker.
// The sticker must belong to a sticker set created by the bot.
// Returns True to success.
func (t *Api) SetStickerKeywords(c *types.SetStickerKeywords) (bool, error) {
	if c.Sticker == "" {
		return false, errors.New("sticker Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// SetStickerMaskPosition Use this method to change the mask position of a mask sticker.
// The sticker must belong to a sticker set that was created by the bot.
// Returns True to success.
func (t *Api) SetStickerMaskPosition(c *types.SetStickerMaskPosition) (bool, error) {
	if c.Sticker == "" {
		return false, errors.New("sticker Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// SetStickerSetTitle Use this method to set the title of a created sticker set.
// Returns True to success.
func (t *Api) SetStickerSetTitle(c *types.SetStickerSetTitle) (bool, error) {
	if c.Name == "" {
		return false, errors.New("name Required")
	}
	if c.Title == "" {
		return false, errors.New("title Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// SetStickerSetThumbnail Use this method to set the thumbnail of a regular or mask sticker set.
// The format of the thumbnail file must match the format of the stickers in the set.
// Returns True to success.
func (t *Api) SetStickerSetThumbnail(c *types.SetStickerSetThumbnail) (bool, error) {
	if c.Name == "" {
		return false, errors.New("name Required")
	}
	if c.UserID == 0 {
		return false, errors.New("userID Required")
	}
	if c.Format == "" {
		return false, errors.New("format Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// SetCustomEmojiStickerSetThumbnail Use this method to set the thumbnail of a custom emoji sticker set.
// Returns True to success.
func (t *Api) SetCustomEmojiStickerSetThumbnail(c *types.SetCustomEmojiStickerSetThumbnail) (bool, error) {
	if c.Name == "" {
		return false, errors.New("name Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// DeleteStickerSet Use this method to delete a sticker set that was created by the bot. Returns True to success.
func (t *Api) DeleteStickerSet(c *types.DeleteStickerSet) (bool, error) {
	if c.Name == "" {
		return false, errors.New("name Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// AnswerInlineQuery Use this method to send answers to an inline query.
// On success, True is returned.
// No more than 50 results per query are allowed.
func (t *Api) AnswerInlineQuery(c *types.AnswerInlineQuery) (bool, error) {
	if c.InlineQueryID == "" {
		return false, errors.New("InlineQueryID Required")
	}
	if c.Results == nil {
		return false, errors.New("results Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// AnswerWebAppQuery Use this method to set the result of an interaction with a Web App
// and send a corresponding message on behalf of the user to the chat from which the query originated.
// On success, a SentWebAppMessage object is returned.
func (t *Api) AnswerWebAppQuery(c *types.AnswerWebAppQuery) (bool, error) {
	if c.WebAppQueryID == "" {
		return false, errors.New("WebAppQueryID Required")
	}
	if c.Result == nil {
		return false, errors.New("result Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// SendInvoice Use this method to send invoices. On success, the sent Message is returned.
func (t *Api) SendInvoice(c *types.SendInvoice) (*types.Message, error) {
	if c.ChatID == 0 && c.ChatIDStr == "" && c.Username == "" {
		return nil, errors.New("ChatID or Username Required")
	}
	if c.Title == "" {
		return nil, errors.New("title Required")
	}
	if c.Description == "" {
		return nil, errors.New("description Required")
	}
	if c.Payload == "" {
		return nil, errors.New("payload Required")
	}
	if c.Currency == "" {
		return nil, errors.New("currency Required")
	}
	if c.Prices == nil {
		return nil, errors.New("prices Required")
	}

	return t.Send(c)
}

// CreateInvoiceLink Use this method to create a link for an invoice. Returns the created invoice link as String on success.
func (t *Api) CreateInvoiceLink(c *types.CreateInvoiceLink) (string, error) {
	if c.Title == "" {
		return "", errors.New("title Required")
	}
	if c.Description == "" {
		return "", errors.New("description Required")
	}
	if c.Payload == "" {
		return "", errors.New("payload Required")
	}
	if c.Currency == "" {
		return "", errors.New("currency Required")
	}
	if c.Prices == nil {
		return "", errors.New("prices Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return "", err
	}

	var createInvoiceLink string

	err = json.Unmarshal(resp.Result, &createInvoiceLink)
	if err != nil {
		return "", err
	}

	return createInvoiceLink, nil
}

// AnswerShippingQuery If you sent an invoice requesting a shipping address and the parameter is_flexible was specified,
// the Bot API will send an Update with a shipping_query field to the bot.
// Use this method to reply to shipping queries.
// On success, True is returned.
func (t *Api) AnswerShippingQuery(c *types.AnswerShippingQuery) (string, error) {
	if c.ShippingQueryID == "" {
		return "", errors.New("ShippingQueryID Required")
	}
	if c.OK {
		if c.ShippingOptions == nil {
			return "", errors.New("ShippingOptions Required")
		}
	} else {
		if c.ErrorMessage == "" {
			return "", errors.New("ErrorMessage Required")
		}
	}

	resp, err := t.Request(c)
	if err != nil {
		return "", err
	}

	var createInvoiceLink string

	err = json.Unmarshal(resp.Result, &createInvoiceLink)
	if err != nil {
		return "", err
	}

	return createInvoiceLink, nil
}

// AnswerPreCheckoutQuery If you sent an invoice requesting a shipping address,
// and the parameter is_flexible was specified,
// the Bot API will send an Update with a shipping_query field to the bot.
// Use this method to reply to shipping queries.
// On success, True is returned.
func (t *Api) AnswerPreCheckoutQuery(c *types.AnswerPreCheckoutQuery) (bool, error) {
	if c.PreCheckoutQueryID == "" {
		return false, errors.New("ShippingQueryID Required")
	}
	if !c.OK {
		if c.ErrorMessage == "" {
			return false, errors.New("ErrorMessage Required")
		}
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// GetStarTransactions Returns the bot's Telegram Star transactions in chronological order.
// On success, returns a StarTransactions object.
func (t *Api) GetStarTransactions(c *types.GetStarTransactions) (*types.StarTransactions, error) {
	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var starTransactions types.StarTransactions
	err = json.Unmarshal(resp.Result, &starTransactions)

	return &starTransactions, err
}

// RefundStarPayment Refunds a successful payment in Telegram Stars.
// Returns True on success.
func (t *Api) RefundStarPayment(c *types.RefundStarPayment) (bool, error) {
	if c.UserId == "" {
		return false, errors.New("UserId Required")
	}
	if c.TelegramPaymentChargeId == "" {
		return false, errors.New("TelegramPaymentChargeId Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// SetPassportDataErrors Informs a user that some of the Telegram Passport elements they provided contains errors.
// The user will not be able to re-submit their Passport to you until the errors are fixed
// (the contents of the field for which you returned the error must change).
// Returns True to success.
// Use this if the data submitted by the user doesn't satisfy the standards your service requires for any reason.
// For example, if a birthday date seems invalid, a submitted document is blurry;
// a scan shows evidence of tampering,
// etc. Supply some details in the error message to make sure the user knows how to correct the issues.
func (t *Api) SetPassportDataErrors(c *types.SetPassportDataErrors) (bool, error) {
	if c.UserID == 0 {
		return false, errors.New("UserID Required")
	}
	if c.Errors == nil {
		return false, errors.New("errors Required")
	}

	resp, err := t.Request(c)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return false, err
	}

	return result, err
}

// SendGame Use this method to send a game. On success, the sent Message is returned.
func (t *Api) SendGame(c *types.SendGame) (*types.Message, error) {
	if c.ChatID == 0 {
		return nil, errors.New("ChatID Required")
	}
	if c.GameShortName == "" {
		return nil, errors.New("GameShortName Required")
	}

	return t.Send(c)
}

// SetGameScore Use this method to set the score of the specified user in a game message.
// On success, if the message is not an inline message, the Message is returned, otherwise True is returned.
// Returns an error, if the new score is not greater than the user's current score in the chat and force is False.
func (t *Api) SetGameScore(c *types.SetGameScore) (*types.Message, error) {
	if c.UserID == 0 {
		return nil, errors.New("UserID Required")
	}
	if c.Score == 0 {
		return nil, errors.New("score Required")
	}
	if c.InlineMessageID == "" {
		if c.ChatID == 0 {
			return nil, errors.New("ChatID Required")
		}
		if c.MessageID == 0 {
			return nil, errors.New("MessageID Required")
		}
	}

	return t.Send(c)
}

// GetGameHighScores Use this method to get data for high-score tables.
// Will return the score of the specified user and several of their neighbors in a game.
// On success, returns an Array of GameHighScore objects.
// This method will currently return scores for the target user, plus two of their closest neighbors on each side.
// Will also return the top three users if the user and their neighbors are not among them.
// Please note that this behavior is subject to change.
func (t *Api) GetGameHighScores(c *types.GetGameHighScores) ([]types.GameHighScore, error) {
	if c.UserID == 0 {
		return nil, errors.New("UserID Required")
	}
	if c.InlineMessageID == "" {
		if c.ChatID == 0 {
			return nil, errors.New("ChatID Required")
		}
		if c.MessageID == 0 {
			return nil, errors.New("MessageID Required")
		}
	}

	resp, err := t.Request(c)
	if err != nil {
		return nil, err
	}

	var gameHighScore []types.GameHighScore

	err = json.Unmarshal(resp.Result, &gameHighScore)
	if err != nil {
		return nil, err
	}

	return gameHighScore, nil
}
