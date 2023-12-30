package telegram

import (
	"io"
	"strings"

	"github.com/raminsa/telegram-bot-api/config"
	"github.com/raminsa/telegram-bot-api/types"
)

// ModeMarkdown return markdown mode
func (t *Api) ModeMarkdown() string {
	return config.ModeMarkdown
}

// ModeMarkdownV2 return markdown2 mode
func (t *Api) ModeMarkdownV2() string {
	return config.ModeMarkdownV2
}

// ModeHTML return html mode
func (t *Api) ModeHTML() string {
	return config.ModeHTML
}

// EscapeText takes an input text and escapes Telegram markup symbols.
// In this way, we can send a text without being afraid of having to escape the characters manually.
// Note that you don't have to include the formatting style in the input text, or it will be escaped too.
// If there is an error, an empty string will be returned.
func (t *Api) EscapeText(parseMode, text string) string {
	var replacer *strings.Replacer

	if parseMode == config.ModeHTML {
		replacer = strings.NewReplacer("<", "&lt;", ">", "&gt;", "&", "&amp;")
	} else if parseMode == config.ModeMarkdown {
		replacer = strings.NewReplacer("_", "\\_", "*", "\\*", "`", "\\`", "[", "\\[")
	} else if parseMode == config.ModeMarkdownV2 {
		replacer = strings.NewReplacer(
			"_", "\\_", "*", "\\*", "[", "\\[", "]", "\\]", "(",
			"\\(", ")", "\\)", "~", "\\~", "`", "\\`", ">", "\\>",
			"#", "\\#", "+", "\\+", "-", "\\-", "=", "\\=", "|",
			"\\|", "{", "\\{", "}", "\\}", ".", "\\.", "!", "\\!",
		)
	} else {
		return ""
	}

	return replacer.Replace(text)
}

// FileBytes return file bytes style.
func (t *Api) FileBytes(name string, byte []byte) *types.FileBytes {
	return &types.FileBytes{
		Name:  name,
		Bytes: byte,
	}
}

// FileReader return file reader style.
func (t *Api) FileReader(name string, reader io.Reader) *types.FileReader {
	return &types.FileReader{
		Name:   name,
		Reader: reader,
	}
}

// FilePath return file path style.
func (t *Api) FilePath(path string) types.FilePath {
	return types.FilePath(path)
}

// FileURL return file url style.
func (t *Api) FileURL(Url string) types.FileURL {
	return types.FileURL(Url)
}

// FileID return file id style already uploaded to Telegram.
func (t *Api) FileID(ID string) types.FileID {
	return types.FileID(ID)
}

// FileAttach return file attaches style used for processed media groups.
func (t *Api) FileAttach(Attach string) types.FileAttach {
	return types.FileAttach(Attach)
}

// GetFileDirectURL returns direct download URL from file
func (t *Api) GetFileDirectURL(fileID string) (string, error) {
	file, err := t.GetFile(&types.GetFile{FileID: fileID})
	if err != nil {
		return "", err
	}

	return file.Link(t.Bot.Token), nil
}

// NewGetUpdates create a new get updates message.
func (t *Api) NewGetUpdates() *types.GetUpdates {
	return &types.GetUpdates{}
}

// NewSetWebhook create a new set webhook message.
func (t *Api) NewSetWebhook() *types.SetWebhook {
	return &types.SetWebhook{}
}

// NewDeleteWebhook create a new delete webhook message.
func (t *Api) NewDeleteWebhook() *types.DeleteWebhook {
	return &types.DeleteWebhook{}
}

// NewReplyKeyboardMarkup creates a new regular keyboard with correct defaults.
func (t *Api) NewReplyKeyboardMarkup(rows ...[]types.KeyboardButton) types.ReplyKeyboardMarkup {
	var keyboard [][]types.KeyboardButton

	keyboard = append(keyboard, rows...)

	return types.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard:       keyboard,
	}
}

// NewKeyboardButton creates a regular keyboard button.
func (t *Api) NewKeyboardButton(text string) types.KeyboardButton {
	return types.KeyboardButton{
		Text: text,
	}
}

// NewKeyboardButtonContact creates a keyboard button that requests. user contact information upon click.
func (t *Api) NewKeyboardButtonContact(text string) types.KeyboardButton {
	return types.KeyboardButton{
		Text:           text,
		RequestContact: true,
	}
}

// NewKeyboardButtonLocation creates a keyboard button that requests. user location information upon click.
func (t *Api) NewKeyboardButtonLocation(text string) types.KeyboardButton {
	return types.KeyboardButton{
		Text:            text,
		RequestLocation: true,
	}
}

// NewKeyboardButtonPool creates a keyboard button that requests.
// For pool, if `quiz` is passed, the user will be allowed to create only polls in the quiz mode.
// If `regular` is passed, only regular polls will be allowed.
func (t *Api) NewKeyboardButtonPool(text, pool string) types.KeyboardButton {
	return types.KeyboardButton{
		Text: text,
		RequestPoll: &types.KeyboardButtonPollType{
			Type: pool,
		},
	}
}

// NewKeyboardButtonWebApp creates a keyboard button that requests. WebApp information upon click.
func (t *Api) NewKeyboardButtonWebApp(text, Url string) types.KeyboardButton {
	return types.KeyboardButton{
		Text: text,
		WebApp: &types.WebAppInfo{
			URL: Url,
		},
	}
}

// NewKeyboardButtonRow creates a row of keyboard buttons.
func (t *Api) NewKeyboardButtonRow(buttons ...types.KeyboardButton) []types.KeyboardButton {
	var row []types.KeyboardButton

	row = append(row, buttons...)

	return row
}

// NewReplyKeyboardRemove hides the keyboard, with the option for being selective or hiding for everyone.
func (t *Api) NewReplyKeyboardRemove(selective bool) types.ReplyKeyboardRemove {
	return types.ReplyKeyboardRemove{
		RemoveKeyboard: true,
		Selective:      selective,
	}
}

// NewForceReply creates a new force reply.
func (t *Api) NewForceReply(inputFieldPlaceholder string, selective bool) types.ForceReply {
	return types.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: inputFieldPlaceholder,
		Selective:             selective,
	}
}

// NewInlineKeyboardMarkup creates a new inline keyboard.
func (t *Api) NewInlineKeyboardMarkup(rows ...[]types.InlineKeyboardButton) types.InlineKeyboardMarkup {
	var keyboard [][]types.InlineKeyboardButton

	keyboard = append(keyboard, rows...)

	return types.InlineKeyboardMarkup{
		InlineKeyboard: keyboard,
	}
}

// NewInlineKeyboardButtonURL creates an inline keyboard url button with text
func (t *Api) NewInlineKeyboardButtonURL(text, url string) types.InlineKeyboardButton {
	return types.InlineKeyboardButton{
		Text: text,
		URL:  url,
	}
}

// NewInlineKeyboardCallbackData creates an inline keyboard callback data button with text
func (t *Api) NewInlineKeyboardCallbackData(text, data string) types.InlineKeyboardButton {
	return types.InlineKeyboardButton{
		Text:         text,
		CallbackData: data,
	}
}

// NewInlineKeyboardWebApp creates an inline keyboard webapp button with text
func (t *Api) NewInlineKeyboardWebApp(text, url string) types.InlineKeyboardButton {
	return types.InlineKeyboardButton{
		Text: text,
		WebApp: &types.WebAppInfo{
			URL: url,
		},
	}
}

// NewInlineKeyboardButtonLoginURL creates an inline keyboard login url button with text
func (t *Api) NewInlineKeyboardButtonLoginURL(text string, loginURL types.LoginURL) types.InlineKeyboardButton {
	return types.InlineKeyboardButton{
		Text:     text,
		LoginURL: &loginURL,
	}
}

// NewInlineKeyboardButtonSwitch creates an inline keyboard switch inline query button with text
func (t *Api) NewInlineKeyboardButtonSwitch(text, switchInline string) types.InlineKeyboardButton {
	return types.InlineKeyboardButton{
		Text:              text,
		SwitchInlineQuery: &switchInline,
	}
}

// NewInlineKeyboardButtonSwitchCurrentChat creates an inline keyboard switch inline query bot chat button with text
func (t *Api) NewInlineKeyboardButtonSwitchCurrentChat(text, switchInline string) types.InlineKeyboardButton {
	return types.InlineKeyboardButton{
		Text:                         text,
		SwitchInlineQueryCurrentChat: &switchInline,
	}
}

// NewInlineKeyboardButtonSwitchChosenChat creates an inline keyboard switch inline query chosen chat-bot chat button with text
func (t *Api) NewInlineKeyboardButtonSwitchChosenChat(text string, switchInline *types.SwitchInlineQueryChosenChat) types.InlineKeyboardButton {
	return types.InlineKeyboardButton{
		Text:                        text,
		SwitchInlineQueryChosenChat: switchInline,
	}
}

// NewInlineKeyboardRow creates an inline keyboard row with buttons.
func (t *Api) NewInlineKeyboardRow(buttons ...types.InlineKeyboardButton) []types.InlineKeyboardButton {
	var row []types.InlineKeyboardButton

	row = append(row, buttons...)

	return row
}

// NewSendMessage create a new sent message.
func (t *Api) NewSendMessage() *types.SendMessage {
	return &types.SendMessage{}
}

// NewForwardMessage create a new forward message.
func (t *Api) NewForwardMessage() *types.ForwardMessage {
	return &types.ForwardMessage{}
}

// NewForwardMessages create a new forward messages.
func (t *Api) NewForwardMessages() *types.ForwardMessages {
	return &types.ForwardMessages{}
}

// NewCopyMessage create a new copy message.
func (t *Api) NewCopyMessage() *types.CopyMessage {
	return &types.CopyMessage{}
}

// NewCopyMessages create a new copy messages.
func (t *Api) NewCopyMessages() *types.CopyMessages {
	return &types.CopyMessages{}
}

// NewSendPhoto create a new photo message.
func (t *Api) NewSendPhoto() *types.SendPhoto {
	return &types.SendPhoto{}
}

// NewSendAudio create a new audio message.
func (t *Api) NewSendAudio() *types.SendAudio {
	return &types.SendAudio{}
}

// NewSendDocument create a new document message.
func (t *Api) NewSendDocument() *types.SendDocument {
	return &types.SendDocument{}
}

// NewSendVideo create a new video message.
func (t *Api) NewSendVideo() *types.SendVideo {
	return &types.SendVideo{}
}

// NewSendAnimation create a new animation message.
func (t *Api) NewSendAnimation() *types.SendAnimation {
	return &types.SendAnimation{}
}

// NewSendVoice create a new voice message.
func (t *Api) NewSendVoice() *types.SendVoice {
	return &types.SendVoice{}
}

// NewSendVideoNote create a new videoNote message.
func (t *Api) NewSendVideoNote() *types.SendVideoNote {
	return &types.SendVideoNote{}
}

// NewSendMediaGroup create a new mediaGroup message.
func (t *Api) NewSendMediaGroup() *types.SendMediaGroup {
	return &types.SendMediaGroup{}
}

// NewInputMediaPhoto creates a new inputMediaPhoto.
func (t *Api) NewInputMediaPhoto() types.InputMediaPhoto {
	return types.InputMediaPhoto{Type: "photo"}
}

// NewInputMediaVideo creates a new inputMediaVideo.
func (t *Api) NewInputMediaVideo() types.InputMediaVideo {
	return types.InputMediaVideo{Type: "video"}
}

// NewInputMediaAnimation creates a new inputMediaAnimation.
func (t *Api) NewInputMediaAnimation() types.InputMediaAnimation {
	return types.InputMediaAnimation{Type: "animation"}
}

// NewInputMediaAudio creates a new inputMediaAudio.
func (t *Api) NewInputMediaAudio() types.InputMediaAudio {
	return types.InputMediaAudio{Type: "audio"}
}

// NewInputMediaDocument creates a new inputMediaDocument.
func (t *Api) NewInputMediaDocument() types.InputMediaDocument {
	return types.InputMediaDocument{Type: "document"}
}

// NewSendLocation creates a new location message.
func (t *Api) NewSendLocation() *types.SendLocation {
	return &types.SendLocation{}
}

// NewEditMessageLiveLocation creates a new edit message live location message.
func (t *Api) NewEditMessageLiveLocation() *types.EditMessageLiveLocation {
	return &types.EditMessageLiveLocation{}
}

// NewStopMessageLiveLocation creates a new stop message live location message.
func (t *Api) NewStopMessageLiveLocation() *types.StopMessageLiveLocation {
	return &types.StopMessageLiveLocation{}
}

// NewSendVenue creates a new venue message.
func (t *Api) NewSendVenue() *types.SendVenue {
	return &types.SendVenue{}
}

// NewSendContact creates a new contact message.
func (t *Api) NewSendContact() *types.SendContact {
	return &types.SendContact{}
}

// NewSendPoll creates a new poll message.
func (t *Api) NewSendPoll() *types.SendPoll {
	return &types.SendPoll{}
}

// NewSendDice creates a new dice message.
func (t *Api) NewSendDice() *types.SendDice {
	return &types.SendDice{}
}

// NewSendChatAction creates a new chat action message.
func (t *Api) NewSendChatAction() *types.SendChatAction {
	return &types.SendChatAction{}
}

// TypingChatAction return typing chat action
func (t *Api) TypingChatAction() string {
	return config.ChatTyping
}

// UploadPhotoChatAction return upload photo action
func (t *Api) UploadPhotoChatAction() string {
	return config.ChatUploadPhoto
}

// RecordVideoChatAction return record video chat action
func (t *Api) RecordVideoChatAction() string {
	return config.ChatRecordVideo
}

// UploadVideoChatAction return upload video chat action
func (t *Api) UploadVideoChatAction() string {
	return config.ChatUploadVideo
}

// RecordVoiceChatAction return record voice chat action
func (t *Api) RecordVoiceChatAction() string {
	return config.ChatRecordVoice
}

// UploadVoiceChatAction return upload voice chat action
func (t *Api) UploadVoiceChatAction() string {
	return config.ChatUploadVoice
}

// UploadDocumentChatAction return upload document chat action
func (t *Api) UploadDocumentChatAction() string {
	return config.ChatUploadDocument
}

// ChooseStickerChatAction return choose sticker chat action
func (t *Api) ChooseStickerChatAction() string {
	return config.ChatChooseSticker
}

// FindLocationChatAction return find location chat action
func (t *Api) FindLocationChatAction() string {
	return config.ChatFindLocation
}

// RecordVideoNoteChatAction return record video note chat action
func (t *Api) RecordVideoNoteChatAction() string {
	return config.ChatRecordVideoNote
}

// UploadVideoNoteChatAction return upload video note chat action
func (t *Api) UploadVideoNoteChatAction() string {
	return config.ChatUploadVideoNote
}

// NewSetMessageReaction creates a new set message reaction.
func (t *Api) NewSetMessageReaction() *types.SetMessageReaction {
	return &types.SetMessageReaction{}
}

// NewGetUserProfilePhotos creates a new get user profile photos message.
func (t *Api) NewGetUserProfilePhotos() *types.GetUserProfilePhotos {
	return &types.GetUserProfilePhotos{}
}

// NewGetFile creates a new get file message.
func (t *Api) NewGetFile() *types.GetFile {
	return &types.GetFile{}
}

// NewBanChatMember creates a new ban chat member message.
func (t *Api) NewBanChatMember() *types.BanChatMember {
	return &types.BanChatMember{}
}

// NewUnbanChatMember creates a new unban chat member message.
func (t *Api) NewUnbanChatMember() *types.UnbanChatMember {
	return &types.UnbanChatMember{}
}

// NewRestrictChatMember creates a new restrict chat member message.
func (t *Api) NewRestrictChatMember() *types.RestrictChatMember {
	return &types.RestrictChatMember{}
}

// NewChatPermissions creates a new chat permissions message.
func (t *Api) NewChatPermissions() types.ChatPermissions {
	return types.ChatPermissions{}
}

// NewPromoteChatMember creates a new promote chat member message.
func (t *Api) NewPromoteChatMember() *types.PromoteChatMember {
	return &types.PromoteChatMember{}
}

// NewSetChatAdministratorCustomTitle creates a new set chat administrator custom title message.
func (t *Api) NewSetChatAdministratorCustomTitle() *types.SetChatAdministratorCustomTitle {
	return &types.SetChatAdministratorCustomTitle{}
}

// NewBanChatSenderChat creates a new ban chat sender chat message.
func (t *Api) NewBanChatSenderChat() *types.BanChatSenderChat {
	return &types.BanChatSenderChat{}
}

// NewUnbanChatSenderChat creates a new unban chat sender chat message.
func (t *Api) NewUnbanChatSenderChat() *types.UnbanChatSenderChat {
	return &types.UnbanChatSenderChat{}
}

// NewSetChatPermissions creates a new set chat permission message.
func (t *Api) NewSetChatPermissions() *types.SetChatPermissions {
	return &types.SetChatPermissions{}
}

// NewExportChatInviteLink creates a new export chat invite link message.
func (t *Api) NewExportChatInviteLink() *types.ExportChatInviteLink {
	return &types.ExportChatInviteLink{}
}

// NewCreateChatInviteLink creates a new create chat invite link message.
func (t *Api) NewCreateChatInviteLink() *types.CreateChatInviteLink {
	return &types.CreateChatInviteLink{}
}

// NewEditChatInviteLink creates a new edit chat invite link message.
func (t *Api) NewEditChatInviteLink() *types.EditChatInviteLink {
	return &types.EditChatInviteLink{}
}

// NewRevokeChatInviteLink creates a new revoke chat invite link message.
func (t *Api) NewRevokeChatInviteLink() *types.RevokeChatInviteLink {
	return &types.RevokeChatInviteLink{}
}

// NewApproveChatJoinRequest creates a new approval chat join request message.
func (t *Api) NewApproveChatJoinRequest() *types.ApproveChatJoinRequest {
	return &types.ApproveChatJoinRequest{}
}

// NewDeclineChatJoinRequest creates a new decline chat join request message.
func (t *Api) NewDeclineChatJoinRequest() *types.DeclineChatJoinRequest {
	return &types.DeclineChatJoinRequest{}
}

// NewSetChatPhoto creates a new set chat photo message.
func (t *Api) NewSetChatPhoto() *types.SetChatPhoto {
	return &types.SetChatPhoto{}
}

// NewDeleteChatPhoto creates a new delete chat photo message.
func (t *Api) NewDeleteChatPhoto() *types.DeleteChatPhoto {
	return &types.DeleteChatPhoto{}
}

// NewSetChatTitle creates a new set chat title message.
func (t *Api) NewSetChatTitle() *types.SetChatTitle {
	return &types.SetChatTitle{}
}

// NewSetChatDescription creates a new set chat description message.
func (t *Api) NewSetChatDescription() *types.SetChatDescription {
	return &types.SetChatDescription{}
}

// NewPinChatMessage creates a new pin chat message.
func (t *Api) NewPinChatMessage() *types.PinChatMessage {
	return &types.PinChatMessage{}
}

// NewUnpinChatMessage creates a new unpin chat message.
func (t *Api) NewUnpinChatMessage() *types.UnpinChatMessage {
	return &types.UnpinChatMessage{}
}

// NewUnpinAllChatMessages creates a new unpin all chat message.
func (t *Api) NewUnpinAllChatMessages() *types.UnpinAllChatMessages {
	return &types.UnpinAllChatMessages{}
}

// NewLeaveChat creates a new leave chat message.
func (t *Api) NewLeaveChat() *types.LeaveChat {
	return &types.LeaveChat{}
}

// NewGetChat creates a new get chat message.
func (t *Api) NewGetChat() *types.GetChat {
	return &types.GetChat{}
}

// NewGetChatAdministrators creates a new get chat administrators message.
func (t *Api) NewGetChatAdministrators() *types.GetChatAdministrators {
	return &types.GetChatAdministrators{}
}

// NewGetChatMemberCount creates a new get chat member count message.
func (t *Api) NewGetChatMemberCount() *types.GetChatMemberCount {
	return &types.GetChatMemberCount{}
}

// NewGetChatMember creates a new get chat member message.
func (t *Api) NewGetChatMember() *types.GetChatMember {
	return &types.GetChatMember{}
}

// NewSetChatStickerSet creates a new set chat sticker set message.
func (t *Api) NewSetChatStickerSet() *types.SetChatStickerSet {
	return &types.SetChatStickerSet{}
}

// NewDeleteChatStickerSet creates a new delete chat sticker set message.
func (t *Api) NewDeleteChatStickerSet() *types.DeleteChatStickerSet {
	return &types.DeleteChatStickerSet{}
}

// NewCreateForumTopic creates a new create forum topic message.
func (t *Api) NewCreateForumTopic() *types.CreateForumTopic {
	return &types.CreateForumTopic{}
}

// NewEditForumTopic creates a new edit forum topic message.
func (t *Api) NewEditForumTopic() *types.EditForumTopic {
	return &types.EditForumTopic{}
}

// NewCloseForumTopic creates a new close forum topic message.
func (t *Api) NewCloseForumTopic() *types.CloseForumTopic {
	return &types.CloseForumTopic{}
}

// NewReopenForumTopic creates a new reopen forum topic message.
func (t *Api) NewReopenForumTopic() *types.ReopenForumTopic {
	return &types.ReopenForumTopic{}
}

// NewDeleteForumTopic creates a new delete forum topic message.
func (t *Api) NewDeleteForumTopic() *types.DeleteForumTopic {
	return &types.DeleteForumTopic{}
}

// NewUnpinAllForumTopicMessages creates a new unpinned all forum topic messages.
func (t *Api) NewUnpinAllForumTopicMessages() *types.UnpinAllForumTopicMessages {
	return &types.UnpinAllForumTopicMessages{}
}

// NewEditGeneralForumTopic creates a new edit general forum topic message.
func (t *Api) NewEditGeneralForumTopic() *types.EditGeneralForumTopic {
	return &types.EditGeneralForumTopic{}
}

// NewCloseGeneralForumTopic creates a new close general forum topic message.
func (t *Api) NewCloseGeneralForumTopic() *types.CloseGeneralForumTopic {
	return &types.CloseGeneralForumTopic{}
}

// NewUnHideGeneralForumTopic creates a new unHide general forum topic.
func (t *Api) NewUnHideGeneralForumTopic() *types.UnHideGeneralForumTopic {
	return &types.UnHideGeneralForumTopic{}
}

// NewUnpinAllGeneralForumTopicMessages creates a new unpin all general forum topic messages.
func (t *Api) NewUnpinAllGeneralForumTopicMessages() *types.UnpinAllGeneralForumTopicMessages {
	return &types.UnpinAllGeneralForumTopicMessages{}
}

// NewHideGeneralForumTopic creates a new hide general forum topic.
func (t *Api) NewHideGeneralForumTopic() *types.HideGeneralForumTopic {
	return &types.HideGeneralForumTopic{}
}

// NewReopenGeneralForumTopic creates a new reopened general forum topic message.
func (t *Api) NewReopenGeneralForumTopic() *types.ReopenGeneralForumTopic {
	return &types.ReopenGeneralForumTopic{}
}

// NewGetForumTopicIconStickers creates a new get forum topic icon stickers.
func (t *Api) NewGetForumTopicIconStickers() *types.GetForumTopicIconStickers {
	return &types.GetForumTopicIconStickers{}
}

// NewAnswerCallbackQuery creates a new answer callback query message.
func (t *Api) NewAnswerCallbackQuery() *types.AnswerCallbackQuery {
	return &types.AnswerCallbackQuery{}
}

// NewGetUserChatBoosts creates a new get user chat boosts.
func (t *Api) NewGetUserChatBoosts() *types.GetUserChatBoosts {
	return &types.GetUserChatBoosts{}
}

// NewSetMyCommands creates a new set my commands message.
func (t *Api) NewSetMyCommands(commands ...types.BotCommand) *types.SetMyCommands {
	return &types.SetMyCommands{Commands: commands}
}

// NewSetMyCommandsWithScope creates a set my commands with a scope message.
func (t *Api) NewSetMyCommandsWithScope(scope *types.BotCommandScope, commands ...types.BotCommand) *types.SetMyCommands {
	return &types.SetMyCommands{Commands: commands, Scope: scope}
}

// NewSetMyCommandsWithScopeAndLanguage creates a set my commands with scope and language message.
func (t *Api) NewSetMyCommandsWithScopeAndLanguage(scope *types.BotCommandScope, languageCode string, commands ...types.BotCommand) *types.SetMyCommands {
	return &types.SetMyCommands{Commands: commands, Scope: scope, LanguageCode: languageCode}
}

// NewGetMyCommandsWithScope creates a get my commands with a scope message.
func (t *Api) NewGetMyCommandsWithScope(scope *types.BotCommandScope) *types.GetMyCommands {
	return &types.GetMyCommands{Scope: scope}
}

// NewGetMyCommandsWithScopeAndLanguage creates a get my commands with scope and language message.
func (t *Api) NewGetMyCommandsWithScopeAndLanguage(scope *types.BotCommandScope, languageCode string) *types.GetMyCommands {
	return &types.GetMyCommands{Scope: scope, LanguageCode: languageCode}
}

// NewDeleteMyCommands creates a new deleted my commands message.
func (t *Api) NewDeleteMyCommands() *types.DeleteMyCommands {
	return &types.DeleteMyCommands{}
}

// NewDeleteMyCommandsWithScope creates a new delete my commands with a scope message.
func (t *Api) NewDeleteMyCommandsWithScope(scope *types.BotCommandScope) *types.DeleteMyCommands {
	return &types.DeleteMyCommands{Scope: scope}
}

// NewDeleteMyCommandsWithScopeAndLanguage creates a new delete my commands with a scope and language message.
func (t *Api) NewDeleteMyCommandsWithScopeAndLanguage(scope *types.BotCommandScope, languageCode string) *types.DeleteMyCommands {
	return &types.DeleteMyCommands{Scope: scope, LanguageCode: languageCode}
}

// NewBotCommand creates a new bot command message.
func (t *Api) NewBotCommand(command, description string) types.BotCommand {
	return types.BotCommand{
		Command:     command,
		Description: description,
	}
}

// NewBotCommandScopeDefault represents the default scope of bot commands.
func (t *Api) NewBotCommandScopeDefault() *types.BotCommandScope {
	return &types.BotCommandScope{Type: "default"}
}

// NewBotCommandScopeAllPrivateChats represents the scope of bot commands,
// covering all private chats.
func (t *Api) NewBotCommandScopeAllPrivateChats() *types.BotCommandScope {
	return &types.BotCommandScope{Type: "all_private_chats"}
}

// NewBotCommandScopeAllGroupChats represents the scope of bot commands,
// covering all group and supergroup chats.
func (t *Api) NewBotCommandScopeAllGroupChats() *types.BotCommandScope {
	return &types.BotCommandScope{Type: "all_group_chats"}
}

// NewBotCommandScopeAllChatAdministrators represents the scope of bot commands, covering all group and supergroup chat administrators.
func (t *Api) NewBotCommandScopeAllChatAdministrators() *types.BotCommandScope {
	return &types.BotCommandScope{Type: "all_chat_administrators"}
}

// NewBotCommandScopeChat represents the scope of bot commands, covering a specific chat.
func (t *Api) NewBotCommandScopeChat(chatID int64) *types.BotCommandScope {
	return &types.BotCommandScope{
		Type:   "chat",
		ChatID: chatID,
	}
}

// NewBotCommandScopeChatAdministrators represents the scope of bot commands,
// covering all administrators of a specific group or supergroup chat.
func (t *Api) NewBotCommandScopeChatAdministrators(chatID int64) *types.BotCommandScope {
	return &types.BotCommandScope{
		Type:   "chat_administrators",
		ChatID: chatID,
	}
}

// NewBotCommandScopeChatMember represents the scope of bot commands, covering a specific member of a group or supergroup chat.
func (t *Api) NewBotCommandScopeChatMember(chatID, userID int64) *types.BotCommandScope {
	return &types.BotCommandScope{
		Type:   "chat_member",
		ChatID: chatID,
		UserID: userID,
	}
}

// NewSetMyName creates a set my name message.
func (t *Api) NewSetMyName() *types.SetMyName {
	return &types.SetMyName{}
}

// NewGetMyName creates a get my name message.
func (t *Api) NewGetMyName() *types.GetMyName {
	return &types.GetMyName{}
}

// NewSetMyDescription creates a set my description message.
func (t *Api) NewSetMyDescription() *types.SetMyDescription {
	return &types.SetMyDescription{}
}

// NewGetMyDescription creates a get my description message.
func (t *Api) NewGetMyDescription() *types.GetMyDescription {
	return &types.GetMyDescription{}
}

// NewSetMyShortDescription creates a set my short description message.
func (t *Api) NewSetMyShortDescription() *types.SetMyShortDescription {
	return &types.SetMyShortDescription{}
}

// NewGetMyShortDescription creates a get my short description message.
func (t *Api) NewGetMyShortDescription() *types.GetMyShortDescription {
	return &types.GetMyShortDescription{}
}

// NewSetChatMenuButton creates a new set chat menu button message.
func (t *Api) NewSetChatMenuButton() *types.SetChatMenuButton {
	return &types.SetChatMenuButton{}
}

// NewMenuButtonCommands represents the menu button of a bot menu.
func (t *Api) NewMenuButtonCommands(command string) *types.MenuButton {
	return &types.MenuButton{
		Type: command,
	}
}

// NewMenuButtonWebApp represents the menu button of a bot menu.
func (t *Api) NewMenuButtonWebApp(text, Url string) *types.MenuButton {
	return &types.MenuButton{
		Type: "web_app",
		Text: text,
		WebApp: &types.WebAppInfo{
			URL: Url,
		},
	}
}

// NewMenuButtonDefault represents the menu button of a bot menu.
func (t *Api) NewMenuButtonDefault() *types.MenuButton {
	return &types.MenuButton{
		Type: "default",
	}
}

// NewGetChatMenuButton creates a new get chat menu button message.
func (t *Api) NewGetChatMenuButton() *types.GetChatMenuButton {
	return &types.GetChatMenuButton{}
}

// NewSetMyDefaultAdministratorRights creates a new a set my default administrator rights message.
func (t *Api) NewSetMyDefaultAdministratorRights() *types.SetMyDefaultAdministratorRights {
	return &types.SetMyDefaultAdministratorRights{}
}

// NewGetMyDefaultAdministratorRights creates a new get my default administrator rights message.
func (t *Api) NewGetMyDefaultAdministratorRights() *types.GetMyDefaultAdministratorRights {
	return &types.GetMyDefaultAdministratorRights{}
}

// NewEditMessageText creates a new edit message text message.
func (t *Api) NewEditMessageText() *types.EditMessageText {
	return &types.EditMessageText{}
}

// NewEditMessageCaption creates a new edit message caption message.
func (t *Api) NewEditMessageCaption() *types.EditMessageCaption {
	return &types.EditMessageCaption{}
}

// NewEditMessageMedia creates a new edit message media message.
func (t *Api) NewEditMessageMedia() *types.EditMessageMedia {
	return &types.EditMessageMedia{}
}

// NewEditMessageReplyMarkup creates a new edit message reply markup message.
func (t *Api) NewEditMessageReplyMarkup() *types.EditMessageReplyMarkup {
	return &types.EditMessageReplyMarkup{}
}

// NewStopPoll creates a new stop poll message.
func (t *Api) NewStopPoll() *types.StopPoll {
	return &types.StopPoll{}
}

// NewDeleteMessage creates a new delete message.
func (t *Api) NewDeleteMessage() *types.DeleteMessage {
	return &types.DeleteMessage{}
}

// NewDeleteMessages creates a new delete messages.
func (t *Api) NewDeleteMessages() *types.DeleteMessages {
	return &types.DeleteMessages{}
}

// NewSendSticker creates a new send sticker message.
func (t *Api) NewSendSticker() *types.SendSticker {
	return &types.SendSticker{}
}

// NewGetStickerSet creates a new get sticker set message.
func (t *Api) NewGetStickerSet() *types.GetStickerSet {
	return &types.GetStickerSet{}
}

// NewGetCustomEmojiStickers creates a new get custom emoji stickers message.
func (t *Api) NewGetCustomEmojiStickers() *types.GetCustomEmojiStickers {
	return &types.GetCustomEmojiStickers{}
}

// NewUploadStickerFile creates a new upload sticker file message.
func (t *Api) NewUploadStickerFile() *types.UploadStickerFile {
	return &types.UploadStickerFile{}
}

// NewCreateNewStickerSet creates a new creation new sticker set message.
func (t *Api) NewCreateNewStickerSet() *types.CreateNewStickerSet {
	return &types.CreateNewStickerSet{}
}

// NewAddStickerToSet creates a new added sticker to set a message.
func (t *Api) NewAddStickerToSet() *types.AddStickerToSet {
	return &types.AddStickerToSet{}
}

// NewInputSticker creates a new input sticker message.
func (t *Api) NewInputSticker() types.InputSticker {
	return types.InputSticker{}
}

// NewMaskPosition creates a new mask position message.
func (t *Api) NewMaskPosition() *types.MaskPosition {
	return &types.MaskPosition{}
}

// NewSetStickerPositionInSet creates a new set sticker position in a set message.
func (t *Api) NewSetStickerPositionInSet() *types.SetStickerPositionInSet {
	return &types.SetStickerPositionInSet{}
}

// NewDeleteStickerFromSet creates a new delete sticker from set message.
func (t *Api) NewDeleteStickerFromSet() *types.DeleteStickerFromSet {
	return &types.DeleteStickerFromSet{}
}

// NewSetStickerEmojiList creates a new set sticker emoji list message.
func (t *Api) NewSetStickerEmojiList() *types.SetStickerEmojiList {
	return &types.SetStickerEmojiList{}
}

// NewSetStickerKeywords creates a new set sticker keywords message.
func (t *Api) NewSetStickerKeywords() *types.SetStickerKeywords {
	return &types.SetStickerKeywords{}
}

// NewSetStickerMaskPosition creates a new set sticker mask position message.
func (t *Api) NewSetStickerMaskPosition() *types.SetStickerMaskPosition {
	return &types.SetStickerMaskPosition{}
}

// NewSetStickerSetTitle creates a new set sticker set title message.
func (t *Api) NewSetStickerSetTitle() *types.SetStickerSetTitle {
	return &types.SetStickerSetTitle{}
}

// NewSetStickerSetThumbnail creates a new set sticker set thumbnail message.
func (t *Api) NewSetStickerSetThumbnail() *types.SetStickerSetThumbnail {
	return &types.SetStickerSetThumbnail{}
}

// NewSetCustomEmojiStickerSetThumbnail creates a new set custom emoji sticker set thumbnail message.
func (t *Api) NewSetCustomEmojiStickerSetThumbnail() *types.SetCustomEmojiStickerSetThumbnail {
	return &types.SetCustomEmojiStickerSetThumbnail{}
}

// NewDeleteStickerSet creates a new delete sticker set message.
func (t *Api) NewDeleteStickerSet() *types.DeleteStickerSet {
	return &types.DeleteStickerSet{}
}

// NewInlineQueryResultsButtonStartParameter creates a new inline query results button startParameter.
func (t *Api) NewInlineQueryResultsButtonStartParameter(title, startParameter string) *types.InlineQueryResultsButton {
	return &types.InlineQueryResultsButton{
		Text:           title,
		StartParameter: &startParameter,
	}
}

// NewInlineQueryResultsButtonWebApp creates a new inline query results button webApp.
func (t *Api) NewInlineQueryResultsButtonWebApp(title string, webApp *types.WebAppInfo) *types.InlineQueryResultsButton {
	return &types.InlineQueryResultsButton{
		Text:   title,
		WebApp: webApp,
	}
}

// NewAnswerInlineQuery creates a new answer inline query message.
func (t *Api) NewAnswerInlineQuery() *types.AnswerInlineQuery {
	return &types.AnswerInlineQuery{}
}

// NewInlineQueryResultArticle creates a new inline query article.
func (t *Api) NewInlineQueryResultArticle(id, title string) *types.InlineQueryResultArticle {
	return &types.InlineQueryResultArticle{
		Type:  "article",
		ID:    id,
		Title: title,
	}
}

// NewInputTextMessageContent creates a new input text message content.
func (t *Api) NewInputTextMessageContent() *types.InputTextMessageContent {
	return &types.InputTextMessageContent{}
}

// NewInputLocationMessageContent creates a new input location message content.
func (t *Api) NewInputLocationMessageContent() *types.InputLocationMessageContent {
	return &types.InputLocationMessageContent{}
}

// NewInputVenueMessageContent creates a new input venue message content.
func (t *Api) NewInputVenueMessageContent() *types.InputVenueMessageContent {
	return &types.InputVenueMessageContent{}
}

// NewInputContactMessageContent creates a new input contact message content.
func (t *Api) NewInputContactMessageContent() *types.InputContactMessageContent {
	return &types.InputContactMessageContent{}
}

// NewInputInvoiceMessageContent creates a new input invoice message content.
func (t *Api) NewInputInvoiceMessageContent() *types.InputInvoiceMessageContent {
	return &types.InputInvoiceMessageContent{}
}

// NewInlineQueryResultGIF creates a new inline query GIF.
func (t *Api) NewInlineQueryResultGIF(id, url string) *types.InlineQueryResultGIF {
	return &types.InlineQueryResultGIF{
		Type: "gif",
		ID:   id,
		URL:  url,
	}
}

// NewInlineQueryResultCachedGIF create a new inline query with a cached photo.
func (t *Api) NewInlineQueryResultCachedGIF(id, gifID string) *types.InlineQueryResultCachedGIF {
	return &types.InlineQueryResultCachedGIF{
		Type:  "gif",
		ID:    id,
		GifID: gifID,
	}
}

// NewInlineQueryResultMPEG4GIF creates a new inline query MPEG4 GIF.
func (t *Api) NewInlineQueryResultMPEG4GIF(id, url string) *types.InlineQueryResultMPEG4GIF {
	return &types.InlineQueryResultMPEG4GIF{
		Type: "mpeg4_gif",
		ID:   id,
		URL:  url,
	}
}

// NewInlineQueryResultCachedMPEG4GIF create a new inline query with cached MPEG4 GIF.
func (t *Api) NewInlineQueryResultCachedMPEG4GIF(id, MPEG4GifID string) *types.InlineQueryResultCachedMPEG4GIF {
	return &types.InlineQueryResultCachedMPEG4GIF{
		Type:        "mpeg4_gif",
		ID:          id,
		MPEG4FileID: MPEG4GifID,
	}
}

// NewInlineQueryResultPhoto creates a new inline query photo.
func (t *Api) NewInlineQueryResultPhoto(id, url string) *types.InlineQueryResultPhoto {
	return &types.InlineQueryResultPhoto{
		Type: "photo",
		ID:   id,
		URL:  url,
	}
}

// NewInlineQueryResultPhotoWithThumbnail creates a new inline query photo.
func (t *Api) NewInlineQueryResultPhotoWithThumbnail(id, url, thumbnail string) *types.InlineQueryResultPhoto {
	return &types.InlineQueryResultPhoto{
		Type:         "photo",
		ID:           id,
		URL:          url,
		ThumbnailURL: thumbnail,
	}
}

// NewInlineQueryResultCachedPhoto create a new inline query with a cached photo.
func (t *Api) NewInlineQueryResultCachedPhoto(id, photoID string) *types.InlineQueryResultCachedPhoto {
	return &types.InlineQueryResultCachedPhoto{
		Type:    "photo",
		ID:      id,
		PhotoID: photoID,
	}
}

// NewInlineQueryResultVideo creates a new inline query video.
func (t *Api) NewInlineQueryResultVideo(id, url string) *types.InlineQueryResultVideo {
	return &types.InlineQueryResultVideo{
		Type: "video",
		ID:   id,
		URL:  url,
	}
}

// NewInlineQueryResultCachedVideo create a new inline query with cached video.
func (t *Api) NewInlineQueryResultCachedVideo(id, videoID, title string) *types.InlineQueryResultCachedVideo {
	return &types.InlineQueryResultCachedVideo{
		Type:    "video",
		ID:      id,
		VideoID: videoID,
		Title:   title,
	}
}

// NewInlineQueryResultCachedSticker create a new inline query with cached sticker.
func (t *Api) NewInlineQueryResultCachedSticker(id, stickerID string) *types.InlineQueryResultCachedSticker {
	return &types.InlineQueryResultCachedSticker{
		Type:      "sticker",
		ID:        id,
		StickerID: stickerID,
	}
}

// NewInlineQueryResultAudio creates a new inline query audio.
func (t *Api) NewInlineQueryResultAudio(id, url, title string) *types.InlineQueryResultAudio {
	return &types.InlineQueryResultAudio{
		Type:  "audio",
		ID:    id,
		URL:   url,
		Title: title,
	}
}

// NewInlineQueryResultCachedAudio create a new inline query with a cached photo.
func (t *Api) NewInlineQueryResultCachedAudio(id, audioID string) *types.InlineQueryResultCachedAudio {
	return &types.InlineQueryResultCachedAudio{
		Type:    "audio",
		ID:      id,
		AudioID: audioID,
	}
}

// NewInlineQueryResultVoice creates a new inline query voice.
func (t *Api) NewInlineQueryResultVoice(id, url, title string) *types.InlineQueryResultVoice {
	return &types.InlineQueryResultVoice{
		Type:  "voice",
		ID:    id,
		URL:   url,
		Title: title,
	}
}

// NewInlineQueryResultCachedVoice create a new inline query with a cached photo.
func (t *Api) NewInlineQueryResultCachedVoice(id, voiceID, title string) *types.InlineQueryResultCachedVoice {
	return &types.InlineQueryResultCachedVoice{
		Type:    "voice",
		ID:      id,
		VoiceID: voiceID,
		Title:   title,
	}
}

// NewInlineQueryResultDocument creates a new inline query document.
func (t *Api) NewInlineQueryResultDocument(id, url, title, mimeType string) *types.InlineQueryResultDocument {
	return &types.InlineQueryResultDocument{
		Type:     "document",
		ID:       id,
		URL:      url,
		Title:    title,
		MimeType: mimeType,
	}
}

// NewInlineQueryResultCachedDocument create a new inline query with a cached photo.
func (t *Api) NewInlineQueryResultCachedDocument(id, documentID, title string) *types.InlineQueryResultCachedDocument {
	return &types.InlineQueryResultCachedDocument{
		Type:       "document",
		ID:         id,
		DocumentID: documentID,
		Title:      title,
	}
}

// NewInlineQueryResultLocation creates a new inline query location.
func (t *Api) NewInlineQueryResultLocation(id, title string, latitude, longitude float64) *types.InlineQueryResultLocation {
	return &types.InlineQueryResultLocation{
		Type:      "location",
		ID:        id,
		Title:     title,
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// NewInlineQueryResultVenue creates a new inline query venue.
func (t *Api) NewInlineQueryResultVenue(id, title, address string, latitude, longitude float64) *types.InlineQueryResultVenue {
	return &types.InlineQueryResultVenue{
		Type:      "venue",
		ID:        id,
		Title:     title,
		Address:   address,
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// NewAnswerWebAppQuery creates a new answer web app query.
func (t *Api) NewAnswerWebAppQuery() *types.AnswerWebAppQuery {
	return &types.AnswerWebAppQuery{}
}

// NewSendInvoice creates a new sent invoice message.
func (t *Api) NewSendInvoice() *types.SendInvoice {
	return &types.SendInvoice{}
}

// NewLabeledPrice creates a new labeled price message.
func (t *Api) NewLabeledPrice(label string, amount int) types.LabeledPrice {
	return types.LabeledPrice{
		Label:  label,
		Amount: amount,
	}
}

// NewLabeledPrices creates a new labeled prices message.
func (t *Api) NewLabeledPrices(labeledPrice ...types.LabeledPrice) []types.LabeledPrice {
	var row []types.LabeledPrice

	row = append(row, labeledPrice...)

	return row
}

// NewCreateInvoiceLink creates a new create invoice link.
func (t *Api) NewCreateInvoiceLink() *types.CreateInvoiceLink {
	return &types.CreateInvoiceLink{}
}

// NewAnswerShippingQuery creates a new answer shipping query.
func (t *Api) NewAnswerShippingQuery() *types.AnswerShippingQuery {
	return &types.AnswerShippingQuery{}
}

// NewAnswerPreCheckoutQuery creates a new answer pre-checkout query.
func (t *Api) NewAnswerPreCheckoutQuery() *types.AnswerPreCheckoutQuery {
	return &types.AnswerPreCheckoutQuery{}
}

// NewSetPassportDataErrors creates a new set passport data errors.
func (t *Api) NewSetPassportDataErrors() *types.SetPassportDataErrors {
	return &types.SetPassportDataErrors{}
}

// PassportElementErrorDataField creates a passport element error data field.
func (t *Api) PassportElementErrorDataField() *types.PassportElementErrorDataField {
	return &types.PassportElementErrorDataField{}
}

// PassportElementErrorFrontSide creates a passport element error front side.
func (t *Api) PassportElementErrorFrontSide() *types.PassportElementErrorFrontSide {
	return &types.PassportElementErrorFrontSide{}
}

// PassportElementErrorReverseSide creates a passport element error reverse side.
func (t *Api) PassportElementErrorReverseSide() *types.PassportElementErrorReverseSide {
	return &types.PassportElementErrorReverseSide{}
}

// PassportElementErrorSelfie creates a passport element error selfie.
func (t *Api) PassportElementErrorSelfie() *types.PassportElementErrorSelfie {
	return &types.PassportElementErrorSelfie{}
}

// PassportElementErrorFile creates a passport element error file.
func (t *Api) PassportElementErrorFile() *types.PassportElementErrorFile {
	return &types.PassportElementErrorFile{}
}

// PassportElementErrorFiles creates a passport element error files.
func (t *Api) PassportElementErrorFiles() *types.PassportElementErrorFiles {
	return &types.PassportElementErrorFiles{}
}

// PassportElementErrorTranslationFile creates a passport element error translation file.
func (t *Api) PassportElementErrorTranslationFile() *types.PassportElementErrorTranslationFile {
	return &types.PassportElementErrorTranslationFile{}
}

// PassportElementErrorTranslationFiles creates a passport element error translation files.
func (t *Api) PassportElementErrorTranslationFiles() *types.PassportElementErrorTranslationFiles {
	return &types.PassportElementErrorTranslationFiles{}
}

// PassportElementErrorUnspecified creates a passport element error unspecified.
func (t *Api) PassportElementErrorUnspecified() *types.PassportElementErrorUnspecified {
	return &types.PassportElementErrorUnspecified{}
}

// NewSendGame creates a new send game message.
func (t *Api) NewSendGame() *types.SendGame {
	return &types.SendGame{}
}

// NewSetGameScore creates a new send game message.
func (t *Api) NewSetGameScore() *types.SetGameScore {
	return &types.SetGameScore{}
}

// NewGetGameHighScores creates a new get game high scores.
func (t *Api) NewGetGameHighScores() *types.GetGameHighScores {
	return &types.GetGameHighScores{}
}
