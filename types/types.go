package types

import (
	"encoding/json"
)

// APIResponse is a response from the Telegram API with the result. stored raw.
type APIResponse struct {
	Ok          bool                `json:"ok"`
	Result      json.RawMessage     `json:"result,omitempty"`
	ErrorCode   int                 `json:"error_code,omitempty"`
	Description string              `json:"description,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
}

// Update is an update response, from GetUpdates.
type Update struct {
	UpdateID                int                          `json:"update_id"`                           // The update's unique identifier. Update identifiers start from a certain positive number and increase sequentially. This ID becomes especially handy if you're using webhooks, since it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order. If there are no new updates for at least a week, then the identifier of the next update will be chosen randomly instead of sequentially.
	Message                 *Message                     `json:"message,omitempty"`                   // Optional. New incoming message of any kind - text, photo, sticker, etc.
	EditedMessage           *Message                     `json:"edited_message,omitempty"`            // Optional. New version of a message that is known to the bot and was edited
	ChannelPost             *Message                     `json:"channel_post,omitempty"`              // Optional. New incoming channel post of any kind - text, photo, sticker, etc.
	EditedChannelPost       *Message                     `json:"edited_channel_post,omitempty"`       // Optional. New version of a channel post that is known to the bot and was edited
	BusinessConnection      *BusinessConnection          `json:"business_connection,omitempty"`       // Optional. The bot was connected to or disconnected from a business account, or a user edited an existing connection with the bot
	BusinessMessage         *Message                     `json:"business_message,omitempty"`          // Optional. New non-service message from a connected business account
	EditedBusinessMessage   *Message                     `json:"edited_business_message,omitempty"`   // Optional. New version of a message from a connected business account
	DeletedBusinessMessages *BusinessMessagesDeleted     `json:"deleted_business_messages,omitempty"` // Optional. Messages were deleted from a connected business account
	MessageReaction         *MessageReactionUpdated      `json:"message_reaction,omitempty"`          // Optional. A reaction to a message was changed by a user. The bot must be an administrator in the chat and must explicitly specify "message_reaction" in the list of allowed_updates to receive these updates. The update isn't received for reactions set by bots.
	MessageReactionCount    *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`    // Optional. Reactions to a message with anonymous reactions were changed. The bot must be an administrator in the chat and must explicitly specify "message_reaction_count" in the list of allowed_updates to receive these updates.
	InlineQuery             *InlineQuery                 `json:"inline_query,omitempty"`              // Optional. New incoming inline query
	ChosenInlineResult      *ChosenInlineResult          `json:"chosen_inline_result,omitempty"`      // Optional. The result of an inline query that was chosen by a user and sent to their chat partner. Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
	CallbackQuery           *CallbackQuery               `json:"callback_query,omitempty"`            // Optional. New incoming callback query
	ShippingQuery           *ShippingQuery               `json:"shipping_query,omitempty"`            // Optional. New incoming shipping query. Only for invoices with flexible price
	PreCheckoutQuery        *PreCheckoutQuery            `json:"pre_checkout_query,omitempty"`        // Optional. New incoming pre-checkout query. Contains full information about checkout
	Poll                    *Poll                        `json:"poll,omitempty"`                      // Optional. New poll state. Bots receive only updates about stopped polls and polls, which are sent by the bot
	PollAnswer              *PollAnswer                  `json:"poll_answer,omitempty"`               // Optional. A user changed their answer in a non-anonymous poll. Bots receive new votes only in polls that were sent by the bot itself.
	MyChatMember            *ChatMemberUpdated           `json:"my_chat_member,omitempty"`            // Optional. The bot's chat member status was updated in a chat. For private chats, this update is received only when the bot is blocked or unblocked by the user.
	ChatMember              *ChatMemberUpdated           `json:"chat_member,omitempty"`               // Optional. A chat member's status was updated in a chat. The bot must be an administrator in the chat and must explicitly specify “chat_member” in the list of allowed_updates to receive these updates.
	ChatJoinRequest         *ChatJoinRequest             `json:"chat_join_request,omitempty"`         // Optional. A request to join the chat has been sent. The bot must have the can_invite_users administrator right in the chat to receive these updates.
	ChatBoost               *ChatBoostUpdated            `json:"chat_boost,omitempty"`                // Optional. A chat boost was added or changed. The bot must be an administrator in the chat to receive these updates.
	RemovedChatBoost        *ChatBoostRemoved            `json:"removed_chat_boost,omitempty"`        // Optional. A boost was removed from a chat. The bot must be an administrator in the chat to receive these updates.
}

// WebhookInfo Describes the current status of a webhook.
type WebhookInfo struct {
	URL                          string   `json:"url"`                                       // Webhook URL, may be empty if webhook is not set up
	HasCustomCertificate         bool     `json:"has_custom_certificate"`                    // True, if a custom certificate was provided for webhook certificate checks
	PendingUpdateCount           int      `json:"pending_update_count"`                      // Number of updates awaiting delivery
	IPAddress                    string   `json:"ip_address,omitempty"`                      // Optional. Currently used webhook IP address
	LastErrorDate                int      `json:"last_error_date,omitempty"`                 // Optional. Unix time for the most recent error that happened when trying to deliver an update via webhook
	LastErrorMessage             string   `json:"last_error_message,omitempty"`              // Optional. Error message in human-readable format for the most recent error that happened when trying to deliver an update via webhook
	LastSynchronizationErrorDate int      `json:"last_synchronization_error_date,omitempty"` // Optional. Unix time of the most recent error that happened when trying to synchronize available updates with Telegram datacenters
	MaxConnections               int      `json:"max_connections,omitempty"`                 // Optional. The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery
	AllowedUpdates               []string `json:"allowed_updates,omitempty"`                 // Optional. A list of update types the bot is subscribed to. Defaults to all update types except chat_member
}

// User Represents a Telegram user or bot.
type User struct {
	ID                      int64  `json:"id"`                                    // Unique identifier for this user or bot. This number may have more than 32 significant bits, and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type is safe for storing this identifier.
	IDString                string `json:"-"`                                     // String ID
	IsBot                   bool   `json:"is_bot"`                                // True, if this user is a bot
	FirstName               string `json:"first_name"`                            // User's or bot's first name
	LastName                string `json:"last_name,omitempty"`                   // Optional. User's or bot's last name
	UserName                string `json:"username,omitempty"`                    // Optional. user's or bot's username
	LanguageCode            string `json:"language_code,omitempty"`               // Optional. (IETF language tag)[https://en.wikipedia.org/wiki/IETF_language_tag] of the user's language
	IsPremium               bool   `json:"is_premium,omitempty"`                  // Optional. True, if this user is a Telegram Premium user
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu,omitempty"`    // Optional. True, if this user added the bot to the attachment menu
	CanJoinGroups           bool   `json:"can_join_groups,omitempty"`             // Optional. Is true if the bot can be invited to groups. Returned only in getMe.
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"` // Optional. is true if privacy mode is disabled for the bot. Returned only in getMe.
	SupportsInlineQueries   bool   `json:"supports_inline_queries,omitempty"`     // Optional. Is true if the bot supports inline queries. Returned only in getMe.
	CanConnectToBusiness    bool   `json:"can_connect_to_business,omitempty"`     // Optional. True, if the bot can be connected to a Telegram Business account to receive its messages. Returned only in getMe.
}

// Chat Represents a chat.
type Chat struct {
	ID        int64  `json:"id"`                   // Unique identifier for this chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	IDString  string `json:"-"`                    // String ID
	Type      string `json:"type"`                 // Type of the chat, can be either “private”, “group”, “supergroup” or “channel”
	Title     string `json:"title,omitempty"`      // Optional. Title, for supergroups, channels and group chats
	UserName  string `json:"username,omitempty"`   // Optional. Username, for private chats, supergroups and channels if available
	FirstName string `json:"first_name,omitempty"` // Optional. First name of the other party in a private chat
	LastName  string `json:"last_name,omitempty"`  // Optional. Last name of the other party in a private chat
	IsForum   bool   `json:"is_forum,omitempty"`   // Optional. True, if the supergroup chat is a forum (has topics enabled)
}

// ChatFullInfo Contains full information about a chat.
type ChatFullInfo struct {
	ID                                 int64                 `json:"id"`                                                // Unique identifier for this chat. This number may have more than 32 significant bits, and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type is safe for storing this identifier.
	IDString                           string                `json:"-"`                                                 // String ID
	Type                               string                `json:"type"`                                              // Type of chat can be either “private,” “group,” “supergroup” or “channel”
	Title                              string                `json:"title,omitempty"`                                   // Optional. Title, for supergroups, channels and group chats
	UserName                           string                `json:"username,omitempty"`                                // Optional. Username, for private chats, supergroups and channels if available
	FirstName                          string                `json:"first_name,omitempty"`                              // Optional. First name of the other party in a private chat
	LastName                           string                `json:"last_name,omitempty"`                               // Optional. Last name of the other party in a private chat
	IsForum                            bool                  `json:"is_forum,omitempty"`                                // Optional. True, if the supergroup chat is a forum (has topics enabled)
	Photo                              *ChatPhoto            `json:"photo,omitempty"`                                   // Optional. Chat photo. Returned only in getChat.
	ActiveUsernames                    []string              `json:"active_usernames,omitempty"`                        // Optional. If non-empty, the list of all active chat usernames; for private chats, supergroups and channels. Returned only in getChat.
	Birthdate                          *Birthdate            `json:"birthdate,omitempty"`                               // Optional. For private chats, the date of birth of the user. Returned only in getChat.
	BusinessIntro                      *BusinessIntro        `json:"business_intro,omitempty"`                          // Optional. For private chats with business accounts, the intro of the business. Returned only in getChat.
	BusinessLocation                   *BusinessLocation     `json:"business_location,omitempty"`                       // Optional. For private chats with business accounts, the location of the business. Returned only in getChat.
	BusinessOpeningHours               *BusinessOpeningHours `json:"business_opening_hours,omitempty"`                  // Optional. For private chats with business accounts, the opening hours of the business. Returned only in getChat.
	PersonalChat                       *Chat                 `json:"personal_chat,omitempty"`                           // Optional. For private chats, the personal channel of the user. Returned only in getChat.
	AvailableReactions                 []ReactionType        `json:"available_reactions,omitempty"`                     // Optional. List of available reactions allowed in the chat. If omitted, then all emoji reactions are allowed. Returned only in getChat.
	AccentColorId                      int                   `json:"accent_color_id,omitempty"`                         // Identifier of the accent color for the chat name and backgrounds of the chat photo, reply header, and link preview. See accent colors for more details. Returned only in getChat. Always returned in getChat.
	MaxReactionCount                   int                   `json:"max_reaction_count,omitempty"`                      // The maximum number of reactions that can be set on a message in the chat
	BackgroundCustomEmojiId            string                `json:"background_custom_emoji_id,omitempty"`              // Optional. Custom emoji identifier of emoji chosen by the chat for the reply header and link preview background. Returned only in getChat.
	ProfileAccentColorId               int                   `json:"profile_accent_color_id,omitempty"`                 // Optional. Identifier of the accent color for the chat's profile background. See profile accent colors for more details. Returned only in getChat.
	ProfileBackgroundCustomEmojiId     string                `json:"profile_background_custom_emoji_id,omitempty"`      // Optional. Custom emoji identifier of the emoji chosen by the chat for its profile background. Returned only in getChat.
	EmojiStatusCustomEmojiId           string                `json:"emoji_status_custom_emoji_id,omitempty"`            // Optional. Custom emoji identifier of the emoji status of the chat or the other party in a private chat. Returned only in getChat.
	EmojiStatusExpirationDate          int64                 `json:"emoji_status_expiration_date,omitempty"`            // Optional. Expiration date of the emoji status of the chat or the other party in a private chat, in Unix time, if any. Returned only in getChat.
	Bio                                string                `json:"bio,omitempty"`                                     // Optional. Bio of the other party in a private chat. Returned only in getChat.
	HasPrivateForwards                 bool                  `json:"has_private_forwards,omitempty"`                    // Optional. True, if privacy settings of the other party in the private chat allow to use tg://user?id=<user_id> links only in chats with the user. Returned only in getChat.
	HasRestrictedVoiceAndVideoMessages bool                  `json:"has_restricted_voice_and_video_messages,omitempty"` // Optional. True, if the privacy settings of the other party restrict sending voice and video note messages in the private chat. Returned only in getChat.
	JoinToSendMessages                 bool                  `json:"join_to_send_messages,omitempty"`                   // Optional. True, if users need to join the supergroup before they can send messages. Returned only in getChat.
	JoinByRequest                      bool                  `json:"join_by_request,omitempty"`                         // Optional. True, if all users directly joining the supergroup need to be approved by supergroup administrators. Returned only in getChat.
	Description                        string                `json:"description,omitempty"`                             // Optional. Description, for groups, supergroups and channel chats. Returned only in getChat.
	InviteLink                         string                `json:"invite_link,omitempty"`                             // Optional. Primary invite link, for groups, supergroups and channel chats. Returned only in getChat.
	PinnedMessage                      *Message              `json:"pinned_message,omitempty"`                          // Optional. The most recent pinned message (by sending date). Returned only in getChat.
	Permissions                        *ChatPermissions      `json:"permissions,omitempty"`                             // Optional. Default chat member permissions, for groups and supergroups. Returned only in getChat.
	SlowModeDelay                      int                   `json:"slow_mode_delay,omitempty"`                         // Optional. For supergroups, the minimum allowed delay between consecutive messages sent by each unprivileged user; in seconds. Returned only in getChat.
	UnRestrictBoostCount               int                   `json:"unrestrict_boost_count,omitempty"`                  // Optional. For supergroups, the minimum number of boosts that a non-administrator user needs to add in order to ignore slow mode and chat permissions. Returned only in getChat.
	MessageAutoDeleteTime              int                   `json:"message_auto_delete_time,omitempty"`                // Optional. The time after which all messages sent to the chat will be automatically deleted in seconds. Returned only in getChat.
	HasAggressiveAntiSpamEnabled       bool                  `json:"has_aggressive_anti_spam_enabled,omitempty"`        // Optional. True, if aggressive anti-spam checks are enabled in the supergroup. The field is only available to chat administrators. Returned only in getChat.
	HasHiddenMembers                   bool                  `json:"has_hidden_members,omitempty"`                      // Optional. True, if non-administrators can only get the list of bots and administrators in the chat. Returned only in getChat.
	HasProtectedContent                bool                  `json:"has_protected_content,omitempty"`                   // Optional. True, if messages from the chat can't be forwarded to other chats. Returned only in getChat.
	HasVisibleHistory                  bool                  `json:"has_visible_history,omitempty"`                     // Optional. True, if new chat members will have access to old messages; available only to chat administrators. Returned only in getChat.
	StickerSetName                     string                `json:"sticker_set_name,omitempty"`                        // Optional. For supergroups, the name of a group sticker set. Returned only in getChat.
	CanSetStickerSet                   bool                  `json:"can_set_sticker_set,omitempty"`                     // Optional. True, if the bot can change the group sticker set. Returned only in getChat.
	CustomEmojiStickerSetName          string                `json:"custom_emoji_sticker_set_name,omitempty"`           // Optional. For supergroups, the name of the group's custom emoji sticker set. Custom emoji from this set can be used by all users and bots in the group. Returned only in getChat.
	LinkedChatID                       int64                 `json:"linked_chat_id,omitempty"`                          // Optional. Unique identifier for the linked chat, i.e., the discussion group identifier for a channel and vice versa, for supergroups and channel chats. This identifier may be greater than 32 bits, and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64-bit integer or double-precision float type is safe for storing this identifier. Returned only in getChat.
	Location                           *ChatLocation         `json:"location,omitempty"`                                // Optional. For supergroups, the location to which the supergroup is connected. Returned only in getChat.
}

// Message Represents a message.
type Message struct {
	MessageID                     int                            `json:"message_id"`                                  // Unique message identifier inside this chat
	MessageThreadID               int                            `json:"message_thread_id,omitempty"`                 // Optional. Unique identifier of a message thread to which the message belongs; for supergroups only
	From                          *User                          `json:"from,omitempty"`                              // Optional. Sender of the message; empty for messages sent to channels. For backward compatibility, the field contains a fake sender user in non-channel chats if the message was sent on behalf of a chat.
	SenderChat                    *Chat                          `json:"sender_chat,omitempty"`                       // Optional. Sender of the message, sent on behalf of a chat. For example, the channel itself for channel posts, the supergroup itself for messages from anonymous group administrators, the linked channel for messages automatically forwarded to the discussion group. For backward compatibility, the field contains a fake sender user in non-channel chats if the message was sent on behalf of a chat.
	SenderBoostCount              int                            `json:"sender_boost_count,omitempty"`                // Optional. If the sender of the message boosted the chat, the number of boosts added by the user
	SenderBusinessBot             *User                          `json:"sender_business_bot,omitempty"`               // Optional. The bot that actually sent the message on behalf of the business account. Available only for outgoing messages sent on behalf of the connected business account.
	Date                          int64                          `json:"date"`                                        // date the message was sent in Unix time
	BusinessConnectionId          string                         `json:"business_connection_id,omitempty"`            // Optional. Unique identifier of the business connection from which the message was received. If non-empty, the message belongs to a chat of the corresponding business account that is independent from any potential bot chat which might share the same identifier.
	Chat                          Chat                           `json:"chat"`                                        // The Conversation the message belongs to
	ForwardOrigin                 *MessageOrigin                 `json:"forward_origin,omitempty"`                    // Optional. Information about the original message for forwarded messages
	IsTopicMessage                bool                           `json:"is_topic_message,omitempty"`                  // Optional. True, if the message is sent to a forum topic
	IsAutomaticForward            bool                           `json:"is_automatic_forward,omitempty"`              // Optional. True, if the message is a channel post that was automatically forwarded to the connected discussion group
	ReplyToMessage                *Message                       `json:"reply_to_message,omitempty"`                  // Optional. For replies, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	ExternalReply                 *ExternalReplyInfo             `json:"external_reply,omitempty"`                    // Optional. Information about the message that is being replied to, which may come from another chat or forum topic
	Quote                         *TextQuote                     `json:"quote,omitempty"`                             // Optional. For replies that quote part of the original message, the quoted part of the message
	ReplyToStory                  *Story                         `json:"reply_to_story,omitempty"`                    // Optional. For replies to a story, the original story
	ViaBot                        *User                          `json:"via_bot,omitempty"`                           // Optional. Bot through which the message was sent
	EditDate                      int64                          `json:"edit_date,omitempty"`                         // Optional. Date the message was last edited in Unix time
	HasProtectedContent           bool                           `json:"has_protected_content,omitempty"`             // Optional. True, if the message can't be forwarded
	IsFromOffline                 bool                           `json:"is_from_offline,omitempty"`                   // Optional. True, if the message was sent by an implicit action, for example, as an away or a greeting business message, or as a scheduled message
	MediaGroupID                  string                         `json:"media_group_id,omitempty"`                    // Optional. The unique identifier of a media message group this message belongs to
	AuthorSignature               string                         `json:"author_signature,omitempty"`                  // Optional. Signature of the post-author for messages in channels, or the custom title of an anonymous group administrator
	Text                          string                         `json:"text,omitempty"`                              // Optional. For text messages, the actual UTF-8 text of the message
	Entities                      []*MessageEntity               `json:"entities,omitempty"`                          // Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	LinkPreviewOptions            *LinkPreviewOptions            `json:"link_preview_options,omitempty"`              // Optional. Options used for link preview generation for the message, if it is a text message and link preview options were changed
	EffectId                      string                         `json:"effect_id,omitempty"`                         // Optional. Unique identifier of the message effect added to the message
	Animation                     *Animation                     `json:"animation,omitempty"`                         // Optional. Message is an animation, information about the animation. For backward compatibility, when this field is set, the document field will also be set
	Audio                         *Audio                         `json:"audio,omitempty"`                             // Optional. Message is an audio file, information about the file
	Document                      *Document                      `json:"document,omitempty"`                          // Optional. Message is a general file, information about the file
	Photo                         []*PhotoSize                   `json:"photo,omitempty"`                             // Optional. Message is a photo, available sizes of the photo
	Sticker                       *Sticker                       `json:"sticker,omitempty"`                           // Optional. Message is a sticker, information about the sticker
	Story                         *Story                         `json:"story,omitempty"`                             // Optional. Message is a forwarded story
	Video                         *Video                         `json:"video,omitempty"`                             // Optional. Message is a video, information about the video
	VideoNote                     *VideoNote                     `json:"video_note,omitempty"`                        // Optional. Message is a video note, information about the video message
	Voice                         *Voice                         `json:"voice,omitempty"`                             // Optional. Message is a voice message, information about the file
	Caption                       string                         `json:"caption,omitempty"`                           // Optional. Caption for the animation, audio, document, photo, video or voice
	CaptionEntities               []*MessageEntity               `json:"caption_entities,omitempty"`                  // Optional. For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption
	ShowCaptionAboveMedia         bool                           `json:"show_caption_above_media,omitempty"`          // Optional. True, if the caption must be shown above the message media
	HasMediaSpoiler               bool                           `json:"has_media_spoiler,omitempty"`                 // Optional. True, if the message media is covered by a spoiler animation
	Contact                       *Contact                       `json:"contact,omitempty"`                           // Optional. Message is a shared contact, information about the contact
	Dice                          *Dice                          `json:"dice,omitempty"`                              // Optional. Message is a dice with random value
	Game                          *Game                          `json:"game,omitempty"`                              // Optional. Message is a game, information about the game.
	Poll                          *Poll                          `json:"poll,omitempty"`                              // Optional. Message is a native poll, information about the poll
	Venue                         *Venue                         `json:"venue,omitempty"`                             // Optional. Message is a venue, information about the venue. For backward compatibility, when this field is set, the location field will also be set
	Location                      *Location                      `json:"location,omitempty"`                          // Optional. Message is a shared location, information about the location
	NewChatMembers                []*User                        `json:"new_chat_members,omitempty"`                  // Optional. New members that were added to the group or supergroup and information about them (the bot itself may be one of these members)
	LeftChatMember                *User                          `json:"left_chat_member,omitempty"`                  // Optional. A member was removed from the group, information about them (this member may be the bot itself)
	NewChatTitle                  string                         `json:"new_chat_title,omitempty"`                    // Optional. A chat title was changed to this value
	NewChatPhoto                  []*PhotoSize                   `json:"new_chat_photo,omitempty"`                    // Optional. A chat photo was changed to this value
	DeleteChatPhoto               bool                           `json:"delete_chat_photo,omitempty"`                 // Optional. Service message: the chat photo was deleted
	GroupChatCreated              bool                           `json:"group_chat_created,omitempty"`                // Optional. Service message: the group has been created
	SuperGroupChatCreated         bool                           `json:"supergroup_chat_created,omitempty"`           // Optional. Service message: the supergroup has been created. This field can't be received in a message coming through updates, because bot can't be a member of a supergroup when it is created. It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup.
	ChannelChatCreated            bool                           `json:"channel_chat_created,omitempty"`              // Optional. Service message: the channel has been created. This field can't be received in a message coming through updates, because bot can't be a member of a channel when it is created. It can only be found in reply_to_message if someone replies to a very first message in a channel.
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"` // Optional. Service message: auto-delete timer settings changed in the chat
	MigrateToChatID               int64                          `json:"migrate_to_chat_id,omitempty"`                // Optional. The group has been migrated to a supergroup with the specified identifier. This number may have more than 32 significant bits, and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type is safe for storing this identifier.
	MigrateFromChatID             int64                          `json:"migrate_from_chat_id,omitempty"`              // Optional. The supergroup has been migrated from a group with the specified identifier. This number may have more than 32 significant bits, and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type is safe for storing this identifier.
	PinnedMessage                 *Message                       `json:"pinned_message,omitempty"`                    // Optional. The Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it is itself a reply.
	Invoice                       *Invoice                       `json:"invoice,omitempty"`                           // Optional. Message is an invoice for a payment, information about the invoice.
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment,omitempty"`                // Optional. Message is a service message about a successful payment, information about the payment.
	UsersShared                   *UsersShared                   `json:"users_shared,omitempty"`                      // Optional. Service message: a user was shared with the bot
	ChatShared                    *ChatShared                    `json:"chat_shared,omitempty"`                       // Optional. Service message: a chat was shared with the bot
	ConnectedWebsite              string                         `json:"connected_website,omitempty"`                 // Optional. The domain name of the website on which the user has logged in.
	WriteAccessAllowed            *WriteAccessAllowed            `json:"write_access_allowed,omitempty"`              // Optional. Service message: the user allowed the bot added to the attachment menu to write messages
	PassportData                  *PassportData                  `json:"passport_data,omitempty"`                     // Optional. Telegram Passport data
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`         // Optional. Service message. A user in the chat triggered another user's proximity alert while sharing a Live Location.
	BoostAdded                    *ChatBoostAdded                `json:"boost_added,omitempty"`                       // Optional. Service message: user boosted the chat
	ChatBackgroundSet             *ChatBackground                `json:"chat_background_set,omitempty"`               // Optional. Service message: chat background set
	ForumTopicCreated             *ForumTopicCreated             `json:"forum_topic_created,omitempty"`               // Optional. Service message: forum topic created
	ForumTopicEdited              *ForumTopicEdited              `json:"forum_topic_edited,omitempty"`                // Optional. Service message: forum topic edited
	ForumTopicClosed              *ForumTopicClosed              `json:"forum_topic_closed,omitempty"`                // Optional. Service message: a forum topic closed
	ForumTopicReopened            *ForumTopicReopened            `json:"forum_topic_reopened,omitempty"`              // Optional. Service message: a forum topic reopened
	GeneralForumTopicHidden       *GeneralForumTopicHidden       `json:"general_forum_topic_hidden,omitempty"`        // Optional. Service message: the 'General' forum topic hidden
	GeneralForumTopicUnhidden     *GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden,omitempty"`      // Optional. Service message: the 'General' forum topic unhidden
	GiveawayCreated               *GiveawayCreated               `json:"giveaway_created,omitempty"`                  // Optional. Service message: a scheduled giveaway was created
	Giveaway                      *Giveaway                      `json:"giveaway,omitempty"`                          // Optional. The message is a scheduled giveaway message
	GiveawayWinners               *GiveawayWinners               `json:"giveaway_winners,omitempty"`                  // Optional. A giveaway with public winners was completed
	GiveawayCompleted             *GiveawayCompleted             `json:"giveaway_completed,omitempty"`                // Optional. a giveaway without public winners was completed
	VideoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled,omitempty"`              // Optional. Service message: video chat scheduled
	VideoChatStarted              *VideoChatStarted              `json:"video_chat_started,omitempty"`                // Optional. Service message: video chat started
	VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended,omitempty"`                  // Optional. Service message: video chat ended
	VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited,omitempty"`   // Optional. Service message: new participants invited to a video chat
	WebAppData                    *WebAppData                    `json:"web_app_data,omitempty"`                      // Optional. Service message: data sent by a Web App
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup,omitempty"`                      // Optional. Inline keyboard attached to the message. Login_url buttons are represented as ordinary url buttons.
}

// MessageID Represents a unique message identifier.
type MessageID struct {
	MessageID int `json:"message_id"` // Unique message identifier
}

// InaccessibleMessage Describes a message that was deleted or is otherwise inaccessible to the bot.
type InaccessibleMessage struct {
	Chat      Chat  `json:"chat"`       // Chat the message belonged to
	MessageID int   `json:"message_id"` // Unique message identifier inside the chat
	Date      int64 `json:"date"`       // Always 0. The field can be used to differentiate regular and inaccessible messages.
}

// MaybeInaccessibleMessage Describes a message that can be inaccessible to the bot. It can be one of
type MaybeInaccessibleMessage struct {
	Message
	InaccessibleMessage
}

// MessageEntity Represents one special entity in a text message. For example, hashtags, usernames, URLs, etc
type MessageEntity struct {
	Type          string `json:"type"`                      // Type of the entity. Currently, can be “mention” (@username), “hashtag” (#hashtag), “hashtag” ($USD), “bot_command” (/start@jobs_bot), “url” (https://telegram.org), “email” (do-not-reply@telegram.org), “phone_number” (+1-212-555-0123), “bold” (bold text), “italic” (italic text), “underline” (underlined text), “strikethrough” (strikethrough text), “spoiler” (spoiler message), “code” (mono-width string), “pre” (mono-width block), “text_link” (for clickable text URLs), “text_mention” (for users without usernames)
	Offset        int    `json:"offset"`                    // Offset in UTF-16 code units to the start of the entity
	Length        int    `json:"length"`                    // Length of the entity in UTF-16 code units
	URL           string `json:"url,omitempty"`             // Optional. For “text_link” only, URL that will be opened after user taps on the text
	User          *User  `json:"user,omitempty"`            // Optional. For “text_mention” only, the mentioned user
	Language      string `json:"language,omitempty"`        // Optional. For “pre” only, the programming language of the entity text
	CustomEmojiId string `json:"custom_emoji_id,omitempty"` // Optional. For “custom_emoji” only, unique identifier of the custom emoji. Use getCustomEmojiStickers to get full information about the sticker
}

// TextQuote Contains information about the quoted part of a message that is replied to by the given message.
type TextQuote struct {
	Text     string          `json:"text"`                // Text of the quoted part of a message that is replied to by the given message
	Entities []MessageEntity `json:"entities,omitempty"`  // Optional. Special entities that appear in the quote. Currently, only bold, italic, underline, strikethrough, spoiler, and custom_emoji entities are kept in quotes.
	Position int             `json:"position"`            // Approximate quote position in the original message in UTF-16 code units as specified by the sender
	IsManual bool            `json:"is_manual,omitempty"` // Optional. True, if the quote was chosen manually by the message sender. Otherwise, the quote was added automatically by the server.
}

// ExternalReplyInfo Contains information about a message that is being replied to, which may come from another chat or forum topic.
type ExternalReplyInfo struct {
	Origin             MessageOrigin       `json:"origin"`                         // Origin of the message replied to by the given message
	Chat               *Chat               `json:"chat,omitempty"`                 // Optional. Chat the original message belongs to. Available only if the chat is a supergroup or a channel.
	MessageID          int                 `json:"message_id,omitempty"`           // Optional. Unique message identifier inside the original chat. Available only if the original chat is a supergroup or a channel.
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"` // Optional. Options used for link preview generation for the original message, if it is a text message
	Animation          *Animation          `json:"animation,omitempty"`            // Optional. Message is an animation, information about the animation
	Audio              *Audio              `json:"audio,omitempty"`                // Optional. Message is an audio file, information about the file
	Document           *Document           `json:"document,omitempty"`             // Optional. Message is a general file, information about the file
	Photo              []PhotoSize         `json:"photo,omitempty"`                // Optional. Message is a photo, available sizes of the photo
	Sticker            *Sticker            `json:"sticker,omitempty"`              // Optional. Message is a sticker, information about the sticker
	Story              *Story              `json:"story,omitempty"`                // Optional. Message is a forwarded story
	Video              *Video              `json:"video,omitempty"`                // Optional. Message is a video, information about the video
	VideoNote          *VideoNote          `json:"video_note,omitempty"`           // Optional. Message is a video note, information about the video message
	Voice              *Voice              `json:"voice,omitempty"`                // Optional. Message is a voice message, information about the file
	HasMediaSpoiler    bool                `json:"has_media_spoiler,omitempty"`    // Optional. True, if the message media is covered by a spoiler animation
	Contact            *Contact            `json:"contact,omitempty"`              // Optional. Message is a shared contact, information about the contact
	Dice               *Dice               `json:"dice,omitempty"`                 // Optional. Message is a dice with random value
	Game               *Game               `json:"game,omitempty"`                 // Optional. Message is a game, information about the game. More about games
	Giveaway           *Giveaway           `json:"giveaway,omitempty"`             // Optional. Message is a scheduled giveaway, information about the giveaway
	GiveawayWinners    *GiveawayWinners    `json:"giveaway_winners,omitempty"`     // Optional. A giveaway with public winners was completed
	Invoice            *Invoice            `json:"invoice,omitempty"`              // Optional. Message is an invoice for a payment, information about the invoice.
	Location           *Location           `json:"location,omitempty"`             // Optional. Message is a shared location, information about the location
	Poll               *Poll               `json:"poll,omitempty"`                 // Optional. Message is a native poll, information about the poll
	Venue              *Venue              `json:"venue,omitempty"`                // Optional. Message is a venue, information about the venue
}

// ReplyParameters Describes reply parameters for the message that is being sent.
type ReplyParameters struct {
	MessageID                int             `json:"message_id"`                            // Identifier of the message that will be replied to in the current chat, or in the chat chat_id if it is specified
	ChatID                   string          `json:"chat_id,omitempty"`                     // Optional. If the message to be replied to is from a different chat, unique identifier for the chat or username of the channel (in the format @channelusername)
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified message to be replied to is not found; can be used only for replies in the same chat and forum topic.
	Quote                    string          `json:"quote,omitempty"`                       // Optional. Quoted part of the message to be replied to; 0-1024 characters after entities parsing. The quote must be an exact substring of the message to be replied to, including bold, italic, underline, strikethrough, spoiler, and custom_emoji entities. The message will fail to send if the quote isn't found in the original message.
	QuoteParseMode           string          `json:"quote_parse_mode,omitempty"`            // Optional. Mode for parsing entities in the quote. See formatting options for more details.
	QuoteEntities            []MessageEntity `json:"quote_entities,omitempty"`              // Optional. A JSON-serialized list of special entities that appear in the quote. It can be specified instead of quote_parse_mode.
	QuotePosition            int             `json:"quote_position,omitempty"`              // Optional. Position of the quote in the original message in UTF-16 code units
}

// MessageOrigin Describes the origin of a message. It can be one of
type MessageOrigin struct {
	MessageOriginUser
	MessageOriginHiddenUser
	MessageOriginChat
	MessageOriginChannel
}

// MessageOriginUser Originally sent by a known user.
type MessageOriginUser struct {
	Type       string `json:"type"`        // Type of the message origin, always “user”
	Date       int64  `json:"date"`        // Date the message was sent originally in Unix time
	SenderUser User   `json:"sender_user"` // User that sent the message originally
}

// MessageOriginHiddenUser Originally sent by an unknown user.
type MessageOriginHiddenUser struct {
	Type           string `json:"type"`             // Type of the message origin, always “hidden_user”
	Date           int64  `json:"date"`             // Date the message was sent originally in Unix time
	SenderUserName string `json:"sender_user_name"` // Name of the user that sent the message originally
}

// MessageOriginChat Originally sent on behalf of a chat to a group chat.
type MessageOriginChat struct {
	Type            string `json:"type"`                       // Type of the message origin, always “chat”
	Date            int64  `json:"date"`                       // Date the message was sent originally in Unix time
	SenderChat      Chat   `json:"sender_chat"`                // Chat that sent the message originally
	AuthorSignature string `json:"author_signature,omitempty"` // Optional. For messages originally sent by an anonymous chat administrator, original message author signature
}

// MessageOriginChannel Originally sent to a channel chat.
type MessageOriginChannel struct {
	Type            string `json:"type"`                       // Type of the message origin, always “channel”
	Date            int64  `json:"date"`                       // Date the message was sent originally in Unix time
	SenderChat      Chat   `json:"sender_chat"`                // Channel chat to which the message was originally sent
	MessageID       int    `json:"message_id"`                 // Unique message identifier inside the chat
	AuthorSignature string `json:"author_signature,omitempty"` // Optional. Signature of the original post author
}

// PhotoSize Represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	FileID       string `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string `json:"file_unique_id"`      // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        int    `json:"width"`               // Photo width
	Height       int    `json:"height"`              // Photo height
	FileSize     int64  `json:"file_size,omitempty"` // Optional. File size in bytes
}

// Animation Represents an animation file (GIF or H.264/MPEG-4 AVC video without a sound).
type Animation struct {
	FileID       string     `json:"file_id"`                   // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string     `json:"file_unique_id"`            // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        int        `json:"width"`                     // Video width as defined by sender
	Height       int        `json:"height"`                    // Video height as defined by sender
	Duration     int        `json:"duration"`                  // Duration of the video in seconds as defined by sender
	Thumbnail    *PhotoSize `json:"thumb,thumbnail,omitempty"` // Optional. Animation thumbnail as defined by sender
	FileName     string     `json:"file_name,omitempty"`       // Optional. Original animation filename as defined by sender
	MimeType     string     `json:"mime_type,omitempty"`       // Optional. MIME type of the file as defined by sender
	FileSize     int64      `json:"file_size,omitempty"`       // Optional. File size in bytes. It can be bigger than 2^31, and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type is safe for storing this value.
}

// Audio Represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	FileID       string     `json:"file_id"`                   // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string     `json:"file_unique_id"`            // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Duration     int        `json:"duration"`                  // Duration of the audio in seconds as defined by sender
	Performer    string     `json:"performer,omitempty"`       // Optional. Performer of the audio as defined by sender or by audio tags
	Title        string     `json:"title,omitempty"`           // Optional. Title of the audio as defined by sender or by audio tags
	FileName     string     `json:"file_name,omitempty"`       // Optional. Original filename as defined by sender
	MimeType     string     `json:"mime_type,omitempty"`       // Optional. MIME type of the file as defined by sender
	FileSize     int64      `json:"file_size,omitempty"`       // Optional. File size in bytes. It can be bigger than 2^31, and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type is safe for storing this value.
	Thumbnail    *PhotoSize `json:"thumb,thumbnail,omitempty"` // Optional. Thumbnail of the album cover to which the music file belongs
}

// The Document Represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {
	FileID       string     `json:"file_id"`                   // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string     `json:"file_unique_id"`            // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Thumbnail    *PhotoSize `json:"thumb,thumbnail,omitempty"` // Optional. Document thumbnail as defined by sender
	FileName     string     `json:"file_name,omitempty"`       // Optional. Original filename as defined by sender
	MimeType     string     `json:"mime_type,omitempty"`       // Optional. MIME type of the file as defined by sender
	FileSize     int64      `json:"file_size,omitempty"`       // Optional. File size in bytes. It can be bigger than 2^31, and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type is safe for storing this value.
}

// Story Represents a message about a forwarded story in the chat. Currently, holds no information.
type Story struct {
	Chat Chat  `json:"chat"` // Chat that posted the story
	ID   int64 `json:"id"`   // Unique identifier for the story in the chat
}

// Video Represents a video file.
type Video struct {
	FileID       string     `json:"file_id"`                   // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string     `json:"file_unique_id"`            // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        int        `json:"width"`                     // Video width as defined by sender
	Height       int        `json:"height"`                    // Video height as defined by sender
	Duration     int        `json:"duration"`                  // Duration of the video in seconds as defined by sender
	Thumbnail    *PhotoSize `json:"thumb,thumbnail,omitempty"` // Optional. Video thumbnail
	FileName     string     `json:"file_name,omitempty"`       // Optional. Original filename as defined by sender
	MimeType     string     `json:"mime_type,omitempty"`       // Optional. MIME type of the file as defined by sender
	FileSize     int64      `json:"file_size,omitempty"`       // Optional. File size in bytes. It can be bigger than 2^31, and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type is safe for storing this value.
}

// VideoNote Represents a video message (available in Telegram apps as of v.4.0).
type VideoNote struct {
	FileID       string     `json:"file_id"`                   // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string     `json:"file_unique_id"`            // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Length       int        `json:"length"`                    // Video width and height (diameter of the video message) as defined by sender
	Duration     int        `json:"duration"`                  // Duration of the video in seconds as defined by sender
	Thumbnail    *PhotoSize `json:"thumb,thumbnail,omitempty"` // Optional. Video thumbnail
	FileSize     int64      `json:"file_size,omitempty"`       // Optional. File size in bytes
}

// Voice Represents a voice note.
type Voice struct {
	FileID       string `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string `json:"file_unique_id"`      //	Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Duration     int    `json:"duration"`            //	Duration of the audio in seconds as defined by sender
	MimeType     string `json:"mime_type,omitempty"` //	Optional. MIME type of the file as defined by sender
	FileSize     int64  `json:"file_size,omitempty"` //	Optional. File size in bytes. It can be bigger than 2^31, and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type is safe for storing this value.
}

// Contact Represents a phone contact.
type Contact struct {
	PhoneNumber string `json:"phone_number"`        // Contact's phone number
	FirstName   string `json:"first_name"`          // Contact's first name
	LastName    string `json:"last_name,omitempty"` // Optional. Contact's last name
	UserID      int64  `json:"user_id,omitempty"`   // Optional. Contact's user identifier in Telegram. This number may have more than 32 significant bits, and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type is safe for storing this identifier.
	VCard       string `json:"vcard,omitempty"`     // Optional. Additional data about the contact in the form of a vCard
}

// Dice Represents an animated emoji that displays a random value.
type Dice struct {
	Emoji string `json:"emoji"` // Emoji on which the dice throw animation is based
	Value int    `json:"value"` // Value of the dice, 1-6 for “🎲”, “🎯” and “🎳” base emoji, 1-5 for “🏀” and “⚽” base emoji, 1-64 for “🎰” base emoji
}

// PollOption Contains information about one answer option in a poll.
type PollOption struct {
	Text         string           `json:"text"`          // Option text, 1-100 characters
	TextEntities []*MessageEntity `json:"text_entities"` // Optional. Special entities that appear in the option text. Currently, only custom emoji entities are allowed in poll option texts
	VoterCount   int              `json:"voter_count"`   // Number of users that voted for this option
}

// InputPollOption Contains information about one answer option in a poll to send.
type InputPollOption struct {
	Text          string           `json:"text"`            // Option text, 1-100 characters
	TextParseMode string           `json:"text_parse_mode"` // Optional. Mode for parsing entities in the text. See formatting options for more details. Currently, only custom emoji entities are allowed
	TextEntities  []*MessageEntity `json:"text_entities"`   // Optional. A JSON-serialized list of special entities that appear in the poll option text. It can be specified instead of text_parse_mode
}

// PollAnswer Represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	PollID    string `json:"poll_id"`              // Unique poll identifier
	VoterChat *Chat  `json:"voter_chat,omitempty"` // Optional. The chat that changed the answer to the poll if the voter is anonymous
	User      *User  `json:"user,omitempty"`       // Optional. The user that changed the answer to the poll if the voter isn't anonymous
	OptionIDs []int  `json:"option_ids"`           // 0-based identifiers of answer options, chosen by the user. Maybe empty if the user retracted their vote.
}

// The Poll Contains information about a poll.
type Poll struct {
	ID                    string           `json:"id"`                             // Unique poll identifier
	Question              string           `json:"question"`                       // Poll question, 1-300 characters
	QuestionEntities      []*MessageEntity `json:"question_entities"`              // Optional. Special entities that appear in the question. Currently, only custom emoji entities are allowed in poll questions
	Options               []PollOption     `json:"options"`                        // List of poll options
	TotalVoterCount       int              `json:"total_voter_count"`              // Total number of users that voted in the poll
	IsClosed              bool             `json:"is_closed"`                      // True, if the poll is closed
	IsAnonymous           bool             `json:"is_anonymous"`                   // True, if the poll is anonymous
	Type                  string           `json:"type"`                           // Poll type, currently can be “regular” or “quiz”
	AllowsMultipleAnswers bool             `json:"allows_multiple_answers"`        // True, if the poll allows multiple answers
	CorrectOptionID       int              `json:"correct_option_id,omitempty"`    // Optional. 0-based identifier of the correct answer option. Available only for polls in the quiz mode, which are closed, or was sent (not forwarded) by the bot or to the private chat with the bot.
	Explanation           string           `json:"explanation,omitempty"`          // Optional. Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters
	ExplanationEntities   []*MessageEntity `json:"explanation_entities,omitempty"` // Optional. Special entities like usernames, URLs, bot commands, etc. that appear in the explanation
	OpenPeriod            int              `json:"open_period,omitempty"`          // Optional. Amount of time in seconds the poll will be active after creation
	CloseDate             int64            `json:"close_date,omitempty"`           // Optional. Point in time (Unix timestamp) when the poll will be automatically closed
}

// Location Represents a point on the map.
type Location struct {
	Longitude            float64 `json:"longitude"`                        // Longitude as defined by sender
	Latitude             float64 `json:"latitude"`                         // Latitude as defined by sender
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`    // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod           int     `json:"live_period,omitempty"`            // Optional. Time relative to the message sending date, during which the location can be updated; in seconds. For active live locations only.
	Heading              int     `json:"heading,omitempty"`                // Optional. The direction in which user is moving, in degrees; 1-360. For active live locations only.
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"` // Optional. The maximum distance for proximity alerts about approaching another chat member, in meters. For sent live locations only.
}

// Venue Represents a venue.
type Venue struct {
	Location        Location `json:"location"`                    // Venue location. Can't be a live location
	Title           string   `json:"title"`                       // Name of the venue
	Address         string   `json:"address"`                     // address of the venue
	FoursquareID    string   `json:"foursquare_id,omitempty"`     // Optional. Foursquare identifier of the venue
	FoursquareType  string   `json:"foursquare_type,omitempty"`   // Optional. Foursquare type of the venue. (For example, “arts_entertainment/default,” “arts_entertainment/aquarium” or “food/ice cream.”)
	GooglePlaceID   string   `json:"google_place_id,omitempty"`   // Optional. Google Places identifier of the venue
	GooglePlaceType string   `json:"google_place_type,omitempty"` // Optional. Google Places a type of the venue. (See supported types.)
}

// WebAppData Describes data sent from a Web App to the bot.
type WebAppData struct {
	Data       string `json:"data"`        // The data. Be aware that a bad client can send arbitrary data in this field.
	ButtonText string `json:"button_text"` // Text of the web_app keyboard button from which the Web App was opened. Be aware that a bad client can send arbitrary data in this field.
}

// ProximityAlertTriggered Represents the content of a service message, sent whenever a user in the chat triggers a proximity alert set by another user.
type ProximityAlertTriggered struct {
	Traveler User `json:"traveler"` // User that triggered the alert
	Watcher  User `json:"watcher"`  // User that set the alert
	Distance int  `json:"distance"` // The distance between the users
}

// MessageAutoDeleteTimerChanged Represents a service message about a change in auto-delete timer settings.
type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"` // New auto-delete time for messages in the chat; in seconds
}

// ChatBoostAdded Represents a service message about a user boosting a chat.
type ChatBoostAdded struct {
	BoostCount int `json:"boost_count"` // Number of boosts added by the user
}

// BackgroundFill Describes the way a background is filled based on the selected colors. Currently, it can be one of
type BackgroundFill struct {
	BackgroundFillSolid
	BackgroundFillGradient
	BackgroundFillFreeformGradient
}

// BackgroundFillSolid Filled using the selected color.
type BackgroundFillSolid struct {
	Type  string `json:"type"`  // Type of the background fill, always “solid”
	Color int    `json:"color"` // The color of the background fill in the RGB24 format
}

// BackgroundFillGradient Automatically filled based on the selected colors.
type BackgroundFillGradient struct {
	Type          string `json:"type"`           // Type of the background fill, always “gradient”
	TopColor      int    `json:"top_color"`      // Top color of the gradient in the RGB24 format
	BottomColor   int    `json:"bottom_color"`   // Bottom color of the gradient in the RGB24 format
	RotationAngle int    `json:"rotation_angle"` // Clockwise rotation angle of the background fill in degrees; 0-359
}

// BackgroundFillFreeformGradient Automatically filled based on the selected colors.
type BackgroundFillFreeformGradient struct {
	Type   string `json:"type"`   // Type of the background fill, always “freeform_gradient”
	Colors []int  `json:"colors"` // A list of the 3 or 4 base colors that are used to generate the freeform gradient in the RGB24 format
}

// BackgroundType Describes the type of a background. Currently, it can be one of
type BackgroundType struct {
	BackgroundTypeFill
	BackgroundTypeWallpaper
	BackgroundTypePattern
	BackgroundTypeChatTheme
}

// BackgroundTypeFill Automatically filled based on the selected colors.
type BackgroundTypeFill struct {
	Type             string         `json:"type"`               // Type of the background, always “fill”
	Fill             BackgroundFill `json:"fill"`               // The background fill
	DarkThemeDimming int            `json:"dark_theme_dimming"` // Dimming of the background in dark themes, as a percentage; 0-100
}

// BackgroundTypeWallpaper Wallpaper in the JPEG format.
type BackgroundTypeWallpaper struct {
	Type             string   `json:"type"`               // Type of the background, always “wallpaper”
	Document         Document `json:"document"`           // Document with the wallpaper
	DarkThemeDimming int      `json:"dark_theme_dimming"` // Dimming of the background in dark themes, as a percentage; 0-100
	IsBlurred        bool     `json:"is_blurred"`         // Optional. True, if the wallpaper is downscaled to fit in a 450x450 square and then box-blurred with radius 12
	IsMoving         bool     `json:"is_moving"`          // Optional. True, if the background moves slightly when the device is tilted
}

// BackgroundTypePattern PNG or TGV (gzipped subset of SVG with MIME type “application/x-tgwallpattern”) pattern to be combined with the background fill chosen by the user.
type BackgroundTypePattern struct {
	Type       string         `json:"type"`        // Type of the background, always “pattern”
	Document   Document       `json:"document"`    // Document with the pattern
	Fill       BackgroundFill `json:"fill"`        // The background fill that is combined with the pattern
	Intensity  int            `json:"intensity"`   // Intensity of the pattern when it is shown above the filled background; 0-100
	IsInverted bool           `json:"is_inverted"` // Optional. True, if the background fill must be applied only to the pattern itself. All other pixels are black in this case. For dark themes only
	IsMoving   bool           `json:"is_moving"`   // Optional. True, if the background moves slightly when the device is tilted
}

// BackgroundTypeChatTheme Taken directly from a built-in chat theme.
type BackgroundTypeChatTheme struct {
	Type      string `json:"type"`       // Type of the background, always “chat_theme”
	ThemeName string `json:"theme_name"` // Name of the chat theme, which is usually an emoji
}

// ChatBackground Represents a chat background.
type ChatBackground struct {
	Type BackgroundType `json:"type"` // Number of boosts added by the user
}

// ForumTopicCreated Represents a service message about a new forum topic created in the chat.
type ForumTopicCreated struct {
	Name              string `json:"name"`                           // Name of the topic
	IconColor         int    `json:"icon_color"`                     // Color of the topic icon in RGB format
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"` // Optional. Unique identifier of the custom emoji shown as the topic icon
}

// ForumTopicClosed Represents a service message about a forum topic closed in the chat. Currently, it holds no information.
type ForumTopicClosed struct {
}

// ForumTopicEdited Represents a service message about an edited forum topic.
type ForumTopicEdited struct {
	Name              string `json:"name,omitempty"`                 // Optional. New name of the topic, if it was edited
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"` // Optional. New identifier of the custom emoji shown as the topic icon if it was edited; an empty string if the icon was removed
}

// ForumTopicReopened Represents a service message about a forum topic reopened in the chat. Currently, it holds no information.
type ForumTopicReopened struct {
}

// GeneralForumTopicHidden Represents a service message about a General forum topic hidden in the chat. Currently, it holds no information
type GeneralForumTopicHidden struct {
}

// GeneralForumTopicUnhidden Represents a service message about a General forum topic unhidden in the chat. Currently, it holds no information.
type GeneralForumTopicUnhidden struct {
}

// SharedUser Contains information about a user that was shared with the bot using a KeyboardButtonRequestUser button.
type SharedUser struct {
	UserId    int64        `json:"user_id"`              // Identifier of the shared user. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so 64-bit integers or double-precision float types are safe for storing these identifiers. The bot may not have access to the user and could be unable to use this identifier, unless the user is already known to the bot by some other means.
	FirstName string       `json:"first_name,omitempty"` // Optional. First name of the user, if the name was requested by the bot
	LastName  string       `json:"last_name,omitempty"`  // Optional. Last name of the user, if the name was requested by the bot
	Username  string       `json:"username,omitempty"`   // Optional. Username of the user, if the username was requested by the bot
	Photo     []*PhotoSize `json:"photo,omitempty"`      // Optional. Available sizes of the chat photo, if the photo was requested by the bot
}

// UsersShared Contains information about the users whose identifiers were shared with the bot using a KeyboardButtonRequestUsers button.
type UsersShared struct {
	RequestID int           `json:"request_id"` // Identifier of the request
	Users     []*SharedUser `json:"users"`      // Information about users shared with the bot.
}

// ChatShared Contains information about the chat whose identifier was shared with the bot using a KeyboardButtonRequestChat button.
type ChatShared struct {
	RequestID int          `json:"request_id"`         // Identifier of the request
	ChatID    int64        `json:"chat_id"`            // Identifier of the shared chat. This number may have more than 32 significant bits, and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type is safe for storing this identifier. The bot may not have access to the chat and could be unable to use this identifier, unless the chat is already known to the bot by some other means.
	Title     string       `json:"title,omitempty"`    // Optional. Title of the chat, if the title was requested by the bot.
	Username  string       `json:"username,omitempty"` // Optional. Username of the chat, if the username was requested by the bot and available.
	Photo     []*PhotoSize `json:"photo,omitempty"`    // Optional. Available sizes of the chat photo, if the photo was requested by the bot
}

// WriteAccessAllowed Represents a service message about a user allowing a bot to write messages
// after adding the bot to the attachment menu or launching a Web App from a link.
type WriteAccessAllowed struct {
	FromRequest        bool   `json:"from_request,omitempty"`         // Optional. True, if the access was granted after the user accepted an explicit request from a Web App sent by the method requestWriteAccess
	WebAppName         string `json:"web_app_name,omitempty"`         // Optional. Name of the Web App which was launched from a link
	FromAttachmentMenu bool   `json:"from_attachment_menu,omitempty"` // Optional. True, if the access was granted when the bot was added to the attachment or side menu
}

// VideoChatScheduled Represents a service message about a video chat scheduled in the chat.
type VideoChatScheduled struct {
	StartDate int64 `json:"start_date"` // Point in time (Unix timestamp) when the video chat is supposed to be started by a chat administrator
}

// VideoChatStarted Represents a service message about a video chat started in the chat. Currently, it holds no information.
type VideoChatStarted struct{}

// VideoChatEnded Represents a service message about a video chat ended in the chat.
type VideoChatEnded struct {
	Duration int `json:"duration"` // Video chat duration in seconds
}

// VideoChatParticipantsInvited Represents a service message about new members invited to a video chat.
type VideoChatParticipantsInvited struct {
	Users []User `json:"users"` // New members that were invited to the video chat
}

// GiveawayCreated Represents a service message about the creation of a scheduled giveaway. Currently holds no information.
type GiveawayCreated struct {
}

// Giveaway Represents a message about a scheduled giveaway.
type Giveaway struct {
	Chats                         []Chat   `json:"chats"`                                      // The list of chats which the user must join to participate in the giveaway
	WinnersSelectionDate          int64    `json:"winners_selection_date"`                     // Point in time (Unix timestamp) when winners of the giveaway will be selected
	WinnerCount                   int      `json:"winner_count"`                               // The number of users which are supposed to be selected as winners of the giveaway
	OnlyNewMembers                bool     `json:"only_new_members,omitempty"`                 // Optional. True, if only users who join the chats after the giveaway started should be eligible to win
	HasPublicWinners              bool     `json:"has_public_winners,omitempty"`               // Optional. True, if the list of giveaway winners will be visible to everyone
	PrizeDescription              string   `json:"prize_description,omitempty"`                // Optional. Description of additional giveaway prize
	CountryCodes                  []string `json:"country_codes,omitempty"`                    // Optional. A list of two-letter ISO 3166-1 alpha-2 country codes indicating the countries from which eligible users for the giveaway must come. If empty, then all users can participate in the giveaway. Users with a phone number that was bought on Fragment can always participate in giveaways.
	PremiumSubscriptionMonthCount int      `json:"premium_subscription_month_count,omitempty"` // Optional. The number of months the Telegram Premium subscription won from the giveaway will be active for
}

// GiveawayWinners Represents a message about the completion of a giveaway with public winners.
type GiveawayWinners struct {
	Chat                          Chat   `json:"chats"`                                      // The chat that created the giveaway
	GiveawayMessageId             int    `json:"giveaway_message_id"`                        // Identifier of the messsage with the giveaway in the chat
	WinnersSelectionDate          int64  `json:"winners_selection_date"`                     // Point in time (Unix timestamp) when winners of the giveaway were selected
	WinnerCount                   int    `json:"winner_count"`                               // Total number of winners in the giveaway
	Winners                       []User `json:"winners"`                                    // List of up to 100 winners of the giveaway
	AdditionalChatCount           int    `json:"additional_chat_count,omitempty"`            // Optional. The number of other chats the user had to join in order to be eligible for the giveaway
	PremiumSubscriptionMonthCount int    `json:"premium_subscription_month_count,omitempty"` // Optional. The number of months the Telegram Premium subscription won from the giveaway will be active for
	UnclaimedPrizeCount           int    `json:"unclaimed_prize_count,omitempty"`            // Optional. Number of undistributed prizes
	OnlyNewMembers                bool   `json:"only_new_members,omitempty"`                 // Optional. True, if only users who had joined the chats after the giveaway started were eligible to win
	WasRefunded                   bool   `json:"was_refunded,omitempty"`                     // Optional. True, if the giveaway was canceled because the payment for it was refunded
	PrizeDescription              string `json:"prize_description,omitempty"`                // Optional. Description of additional giveaway prize
}

// GiveawayCompleted Represents a service message about the completion of a giveaway without public winners.
type GiveawayCompleted struct {
	WinnerCount         int      `json:"winner_count"`                    // Number of winners in the giveaway
	UnclaimedPrizeCount int      `json:"unclaimed_prize_count,omitempty"` // Optional. Number of undistributed prizes
	GiveawayMessage     *Message `json:"giveaway_message,omitempty"`      // Optional. Message with the giveaway that was completed, if it wasn't deleted
}

// LinkPreviewOptions Describes the options used for link preview generation.
type LinkPreviewOptions struct {
	IsDisabled       bool   `json:"is_disabled,omitempty"`        // Optional. True, if the link preview is disabled
	Url              string `json:"url,omitempty"`                // Optional. URL to use for the link preview. If empty, then the first URL found in the message text will be used
	PreferSmallMedia bool   `json:"prefer_small_media,omitempty"` // Optional. True, if the media in the link preview is suppposed to be shrunk; ignored if the URL isn't explicitly specified or media size change isn't supported for the preview
	PreferLargeMedia bool   `json:"prefer_large_media,omitempty"` // Optional. True, if the media in the link preview is suppposed to be enlarged; ignored if the URL isn't explicitly specified or media size change isn't supported for the preview
	ShowAboveText    bool   `json:"show_above_text,omitempty"`    // Optional. True, if the link preview must be shown above the message text; otherwise, the link preview will be shown below the message text
}

// UserProfilePhotos represent a user's profile pictures.
type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"` // Total number of profile pictures the target user has
	Photos     [][]PhotoSize `json:"photos"`      // Requested profile pictures (in up to 4 sizes each)
}

// File Represents a file ready to be downloaded.
// The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>. It is guaranteed
// that the link will be valid for at least 1 hour.
// When the link expires, a new one can be requested by calling getFile.
// Note: The maximum file size to download is 20 MB
type File struct {
	FileID       string `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string `json:"file_unique_id"`      // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileSize     int64  `json:"file_size,omitempty"` //	Optional. File size in bytes. It can be bigger than 2^31, and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type is safe for storing this value.
	FilePath     string `json:"file_path,omitempty"` // Optional. File a path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
}

// WebAppInfo Describes a Web App.
type WebAppInfo struct {
	URL string `json:"url"` // An HTTPS URL of a Web App to be opened with additional data as specified in Initializing Web Apps
}

// ReplyKeyboardMarkup Represents a custom keyboard with reply options (see Introduction to bots for details and examples).
type ReplyKeyboardMarkup struct {
	Keyboard              [][]KeyboardButton `json:"keyboard"`                          // Array of button rows, each represented by an Array of KeyboardButton objects
	IsPersistent          bool               `json:"is_persistent,omitempty"`           // Optional. Requests clients to always show the keyboard when the regular keyboard is hidden. Defaults to false, in which case the custom keyboard can be hidden and opened with a keyboard icon.
	ResizeKeyboard        bool               `json:"resize_keyboard,omitempty"`         // Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	OneTimeKeyboard       bool               `json:"one_time_keyboard,omitempty"`       // Optional. Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat - the user can press a special button in the input field to see the custom keyboard again. Defaults to false.
	InputFieldPlaceholder string             `json:"input_field_placeholder,omitempty"` // Optional. The placeholder to be shown in the input field when the keyboard is active; 1-64 characters
	Selective             bool               `json:"selective,omitempty"`               // Optional. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.  Example: A user requests to change the bot's language, bot replies to the request with a keyboard to select the new language. Other users in the group don't see the keyboard.
}

// KeyboardButton Represents one button of the reply keyboard.
// For simple text buttons, String can be used instead of this object to specify the text of the button.
// Optional fields web_app, request_contact, request_location, and request_poll are mutually exclusive.
// Note: request_contact and request_location options will only work in Telegram versions released after 9 April 2016.
// Older clients will display unsupported message.
// Note: request_poll option will only work in Telegram versions released after 23 January 2020.
// Older clients will display unsupported message.
// Note: web_app option will only work in Telegram versions released after 16 April 2022.
// Older clients will display an unsupported message.
type KeyboardButton struct {
	Text            string                      `json:"text"`                       // Text of the button. If none of the optional fields are used, it will be sent as a message when the button is pressed
	RequestUsers    *KeyboardButtonRequestUsers `json:"request_users,omitempty"`    // Optional. If specified, pressing the button will open a list of suitable users. Tapping on any user will send their identifier to the bot in an “users_shared” service message. Available in private chats only.
	RequestChat     *KeyboardButtonRequestChat  `json:"request_chat,omitempty"`     // Optional. If specified, pressing the button will open a list of suitable chats. Tapping on a chat will send its identifier to the bot in a “chat_shared” service message. Available in private chats only.
	RequestContact  bool                        `json:"request_contact,omitempty"`  // Optional. If True, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only.
	RequestLocation bool                        `json:"request_location,omitempty"` // Optional. If True, the user's current location will be sent when the button is pressed. Available in private chats only.
	RequestPoll     *KeyboardButtonPollType     `json:"request_poll,omitempty"`     // Optional. If specified, the user will be asked to create a poll and send it to the bot when the button is pressed. Available in private chats only.
	WebApp          *WebAppInfo                 `json:"web_app,omitempty"`          // Optional. If specified, the described Web App will be launched when the button is pressed. The Web App will be able to send a “web_app_data” service message. Available in private chats only.
}

// KeyboardButtonRequestUsers Defines the criteria used to request suitable users.
// The identifiers of the selected users will be shared with the bot when the corresponding button is pressed.
type KeyboardButtonRequestUsers struct {
	RequestId       int  `json:"request_id"`                 // Signed 32-bit identifier of the request that will be received back in the UsersShared object. Must be unique within the message.
	UserIsBot       bool `json:"user_is_bot,omitempty"`      // Optional. Pass True to request bots, pass False to request regular users. If not specified, no additional restrictions are applied.
	UserIsPremium   bool `json:"user_is_premium,omitempty"`  // Optional. Pass True to request premium users, pass False to request non-premium users. If not specified, no additional restrictions are applied.
	MaxQuantity     int  `json:"max_quantity,omitempty"`     // Optional. The maximum number of users to be selected; 1-10. Defaults to 1.
	RequestName     bool `json:"request_name,omitempty"`     // Optional. Pass True to request the users' first and last name
	RequestUsername bool `json:"request_username,omitempty"` // Optional. Pass True to request the users' username
	RequestPhoto    bool `json:"request_photo,omitempty"`    // Optional. Pass True to request the users' photo
}

// KeyboardButtonRequestChat Defines the criteria used to request a suitable chat.
// The identifier of the selected chat will be shared with the bot when the corresponding button is pressed.
type KeyboardButtonRequestChat struct {
	RequestID               int                      `json:"request_id"`                          // Signed 32-bit identifier of the request, which will be received back in the ChatShared object. Must be unique within the message
	ChatIsChannel           bool                     `json:"chat_is_channel"`                     // Pass True to request a channel chat, pass False to request a group or a supergroup chat.
	ChatIsForum             bool                     `json:"chat_is_forum,omitempty"`             // Optional. Pass True to request a forum supergroup, pass False to request a non-forum chat. If not specified, no additional restrictions are applied.
	ChatHasUsername         bool                     `json:"chat_has_username,omitempty"`         // Optional. Pass True to request a supergroup or a channel with a username, pass False to request a chat without a username. If not specified, no additional restrictions are applied.
	ChatIsCreated           bool                     `json:"chat_is_created,omitempty"`           // Optional. Pass True to request a chat owned by the user. Otherwise, no additional restrictions are applied.
	UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights,omitempty"` // Optional. A JSON-serialized object lists the required administrator rights of the user in the chat. The rights must be a superset of bot_administrator_rights. If not specified, no additional restrictions are applied.
	BotAdministratorRights  *ChatAdministratorRights `json:"bot_administrator_rights,omitempty"`  // Optional. A JSON-serialized object listing the required administrator rights of the bot in the chat. The rights must be a subset of user_administrator_rights. If not specified, no additional restrictions are applied.
	BotIsMember             bool                     `json:"bot_is_member,omitempty"`             // Optional. Pass True to request a chat with the bot as a member. Otherwise, no additional restrictions are applied.
	RequestTitle            bool                     `json:"request_title,omitempty"`             // Optional. Pass True to request a chat with the bot as a member. Otherwise, no additional restrictions are applied.
	RequestUsername         bool                     `json:"request_username,omitempty"`          // Optional. Pass True to request a chat with the bot as a member. Otherwise, no additional restrictions are applied.
	RequestPhoto            bool                     `json:"request_photo,omitempty"`             // Optional. Pass True to request a chat with the bot as a member. Otherwise, no additional restrictions are applied.
}

// KeyboardButtonPollType Represents a type of poll, which is allowed to be created and sent when the corresponding button is pressed.
type KeyboardButtonPollType struct {
	Type string `json:"type"` // Optional. If quiz is passed, the user will be allowed to create only polls in the quiz mode. If regular is passed, only regular polls will be allowed. Otherwise, the user will be allowed to create a poll of any type.
}

// ReplyKeyboardRemove Upon receiving a message with this object,
// Telegram clients will remove the current custom keyboard and display the default letter-keyboard.
// By default, custom keyboards are displayed until a new keyboard is sent by a bot.
// An exception is made for one-time keyboards that are hidden immediately after the user presses a button
// (see ReplyKeyboardMarkup).
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`     // Requests clients to remove the custom keyboard (user will not be able to summon this keyboard; if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in ReplyKeyboardMarkup)
	Selective      bool `json:"selective,omitempty"` //	Optional. Use this parameter if you want to remove the keyboard for specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message. Example: A user votes in a poll, bot returns a confirmation message in reply to the vote and removes the keyboard for that user, while still showing the keyboard with poll options to users who haven't voted yet.
}

// InlineKeyboardMarkup Represents an inline keyboard that appears right next to
// the message it belongs to. Note: This will only work in Telegram versions
// released after 9 April 2016. Older clients will display an unsupported message.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"` // Array of button rows, each represented by an Array of InlineKeyboardButton objects
}

// InlineKeyboardButton Represents one button of an inline keyboard. You must use exactly one of the optional fields.
type InlineKeyboardButton struct {
	Text                         string                       `json:"text"`                                       // Label text on the button
	URL                          string                       `json:"url,omitempty"`                              // Optional. HTTP or tg:// URL to be opened when the button is pressed. Links tg://user?id=<user_id> can be used to mention a user by their ID without using a username, if this is allowed by their privacy settings.
	CallbackData                 string                       `json:"callback_data,omitempty"`                    // Optional. Description of the Web App that will be launched when the user presses the button. The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery. Available only in private chats between a user and the bot.
	WebApp                       *WebAppInfo                  `json:"web_app,omitempty"`                          // Optional. An HTTPS URL used to automatically authorize the user. Can be used as a replacement for the Telegram Login Widget.
	LoginURL                     *LoginURL                    `json:"login_url,omitempty"`                        // Optional. Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
	SwitchInlineQuery            *string                      `json:"switch_inline_query,omitempty"`              // Optional. If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot's username and the specified inline query in the input field. Maybe empty, in which case just the bot's username will be inserted. Note: This offers an easy way for users to start using your bot in inline mode when they are currently in a private chat with it. Especially useful when combined with switch_pm… actions - in this case, the user will be automatically returned to the chat they switched from, skipping the chat selection screen.
	SwitchInlineQueryCurrentChat *string                      `json:"switch_inline_query_current_chat,omitempty"` // Optional. If set, pressing the button will insert the bot's username and the specified inline query in the current chat's input field. Maybe empty, in which case only the bot's username will be inserted. This offers a quick way for the user to open your bot in inline mode in the same chat - good for selecting something from multiple options.
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`  // Optional. If set, pressing the button will prompt the user to select one of their chats of the specified type, open that chat and insert the bot's username and the specified inline query in the input field.
	CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`                    // Optional. Description of the game that will be launched when the user presses the button. NOTE: This type of button must always be the first button in the first row.
	Pay                          bool                         `json:"pay,omitempty"`                              // Optional. Specify True, to send a Pay button. NOTE: This type of button must always be the first button in the first row and can only be used in invoice messages.
}

// LoginURL Represents a parameter of the inline keyboard button used to automatically authorize a user.
// Serves as a great replacement for the Telegram Login Widget when the user is coming from Telegram.
// All the user needs to do is tap/click a button and confirm that they want to log in:
// Telegram apps support these buttons as of version 5.7.
// Sample bot: @discuss-bot
type LoginURL struct {
	URL                string `json:"url"`                            // An HTTP URL to be opened with user authorization data added to the query string when the button is pressed. If the user refuses to provide authorization data, the original URL without information about the user will be opened. The data added is the same as described in Receiving authorization data. NOTE: You must always check the hash of the received data to verify the authentication and the integrity of the data as described in Checking authorization.
	ForwardText        string `json:"forward_text,omitempty"`         // Optional. New text of the button in forwarded messages.
	BotUsername        string `json:"bot_username,omitempty"`         // Optional. Username of a bot, which will be used for user authorization. See Setting up a bot for more details. If not specified, the current bot's username will be assumed. The hurl's domain must be the same as the domain linked with the bot. See Linking your domain to the bot for more details.
	RequestWriteAccess bool   `json:"request_write_access,omitempty"` // Optional. Pass True to request the permission for your bot to send messages to the user.
}

// SwitchInlineQueryChosenChat Represents an inline button that switches the current user to inline mode in a chosen chat,
// with an optional default inline query.
type SwitchInlineQueryChosenChat struct {
	Query             string `json:"query,omitempty"`               // Optional. The default inline query to be inserted in the input field. If left empty, only the bot's username will be inserted
	AllowUserChats    bool   `json:"allow_user_chats,omitempty"`    // Optional. True, if private chats with users can be chosen
	AllowBotChats     bool   `json:"allow_bot_chats,omitempty"`     // Optional. True, if private chats with bots can be chosen
	AllowGroupChats   bool   `json:"allow_group_chats,omitempty"`   // Optional. True, if group and supergroup chats can be chosen
	AllowChannelChats bool   `json:"allow_channel_chats,omitempty"` // Optional. True, if channel chats can be chosen
}

// CallbackQuery Represents an incoming callback query from a callback button in an inline keyboard.
// If the button that originated the query was attached to a message sent by the bot, the field message will be present.
// If the button was attached to a message sent via the bot (in inline mode),
// the field inline_message_id will be present.
// Exactly one of the field data or game_short_name will be present.
// NOTE:
// After the user presses a callback button,
// Telegram clients will display a progress bar until you call answerCallbackQuery.
// It is, therefore, necessary to react by calling answerCallbackQuery even if no notification to the user is needed
// (e.g., without specifying any of the optional parameters).
type CallbackQuery struct {
	ID              string   `json:"id"`                          // Unique identifier for this query
	From            User     `json:"from"`                        // Sender
	Message         *Message `json:"message,omitempty"`           // Optional. Message with the callback button that originated the query. Note that message content and message date will not be available if the message is too old
	InlineMessageID string   `json:"inline_message_id,omitempty"` // Optional. Identifier of the message sent via the bot in inline mode, that originated the query.
	ChatInstance    string   `json:"chat_instance"`               // Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent. Useful for high scores in games.
	Data            string   `json:"data,omitempty"`              // Optional. Data associated with the callback button. Be aware that the message originated the query can contain no callback buttons with this data.
	GameShortName   string   `json:"game_short_name,omitempty"`   // Optional. Short name of a Game to be returned, serves as the unique identifier for the game
}

// ForceReply Upon receiving a message with this object,
// Telegram clients will display a reply interface to the user
// (act as if the user has selected the bot's message and tapped 'Reply').
// This can be extremely useful
// if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.
// Example: A poll bot for groups runs in privacy mode (only receives commands, replies to its messages and mentions).
// There could be two ways to create a new poll:
//
//	Explain the user how to send a command with parameters (e.g. /new poll question answer1 answer2).
//	Maybe appealing for hardcore users but lacks modern day polish.
//	Guide the user through a step-by-step process.
//	'Please send me your question', 'Cool, now let's add the first answer option', 'Great.
//	Keep adding answer options, then send /done when you're ready'.
//	The last option is definitely more attractive.
//	And if you use ForceReply in your bot's questions,
//	it will receive the user's answers even if it only receives replies,
//	commands and mentions - without any extra work for the user.
type ForceReply struct {
	ForceReply            bool   `json:"force_reply"`                       // Shows reply interface to the user, as if they manually selected the bot's message and tapped 'Reply'
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"` // Optional. The placeholder to be shown in the input field when the reply is active; 1-64 characters
	Selective             bool   `json:"selective,omitempty"`               // Optional. Use this parameter if you want to force reply from specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
}

// ChatPhoto Represents a chat photo.
type ChatPhoto struct {
	SmallFileID       string `json:"small_file_id"`        // File identifier of a small (160x160) chat photo. This file_id can be used only for photo download and only for as long as the photo is not changed.
	SmallFileUniqueID string `json:"small_file_unique_id"` // Unique file identifier of a small (160x160) chat photo, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	BigFileID         string `json:"big_file_id"`          // File identifier of a big (640x640) chat photo. This file_id can be used only for photo download and only for as long as the photo is not changed.
	BigFileUniqueID   string `json:"big_file_unique_id"`   // Unique file identifier of a big (640x640) chat photo, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
}

// ChatInviteLink Represents an invitation link for a chat.
type ChatInviteLink struct {
	InviteLink              string `json:"invite_link"`                          // The invite link. If the link was created by another chat administrator, then the second part of the link will be replaced with “…”.
	Creator                 User   `json:"creator"`                              // Creator of the link
	CreatesJoinRequest      bool   `json:"creates_join_request"`                 // True, if users joining the chat via the link need to be approved by chat administrators
	IsPrimary               bool   `json:"is_primary"`                           // True, if the link is primary
	IsRevoked               bool   `json:"is_revoked"`                           // True, if the link is revoked
	Name                    string `json:"name,omitempty"`                       // Optional. Invite link name
	ExpireDate              int64  `json:"expire_date,omitempty"`                // Optional. Point in time (Unix timestamp) when the link will expire or has been expired
	MemberLimit             int    `json:"member_limit,omitempty"`               // Optional. The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	PendingJoinRequestCount int    `json:"pending_join_request_count,omitempty"` // Optional. Number of pending join requests created using this link
}

// ChatAdministratorRights Represents the rights of an administrator in a chat.
type ChatAdministratorRights struct {
	IsAnonymous         bool `json:"is_anonymous"`                 // True, if the user's presence in the chat is hidden
	CanManageChat       bool `json:"can_manage_chat"`              // True, if the administrator can access the chat event log, chat statistics, message statistics in channels, see channel members, see anonymous administrators in supergroups and ignore slow mode. Implied by any other administrator privilege
	CanDeleteMessages   bool `json:"can_delete_messages"`          // True, if the administrator can delete messages of other users
	CanManageVideoChats bool `json:"can_manage_video_chats"`       // True, if the administrator can manage video chats
	CanRestrictMembers  bool `json:"can_restrict_members"`         // True, if the administrator can restrict, ban or unban chat members
	CanPromoteMembers   bool `json:"can_promote_members"`          // True, if the administrator can add new administrators with a subset of their own privileges or demote administrators that he has promoted, directly or indirectly (promoted by administrators that were appointed by the user)
	CanChangeInfo       bool `json:"can_change_info"`              //	True, if the user is allowed to change the chat title, photo and other settings
	CanInviteUsers      bool `json:"can_invite_users"`             // True, if the user is allowed to invite new users to the chat
	CanPostMessages     bool `json:"can_post_messages,omitempty"`  // Optional. True, if the administrator can post in the channel; channels only
	CanEditMessages     bool `json:"can_edit_messages,omitempty"`  // Optional. True, if the administrator can edit messages of other users and can pin messages; channels only
	CanPinMessages      bool `json:"can_pin_messages,omitempty"`   // Optional. True, if the user is allowed to pin messages, groups and supergroups only
	CanPostStories      bool `json:"can_post_stories,omitempty"`   // Optional. True, if the administrator can post stories in the channel; channels only
	CanEditStories      bool `json:"can_edit_stories,omitempty"`   // Optional. True, if the administrator can edit stories posted by other users; channels only
	CanDeleteStories    bool `json:"can_delete_stories,omitempty"` // Optional. True, if the administrator can delete stories posted by other users; channels only
	CanManageTopics     bool `json:"can_manage_topics,omitempty"`  // Optional. True, if the user is allowed to create, rename, close, and reopen forum topics; supergroups only
}

// ChatMemberUpdated Represents changes in the status of a chat member.
type ChatMemberUpdated struct {
	Chat                    Chat            `json:"chat"`                                  // the user belongs to
	From                    User            `json:"from"`                                  // Performer of the action, which resulted in the change
	Date                    int             `json:"date"`                                  // the change was done in Unix time
	OldChatMember           ChatMember      `json:"old_chat_member"`                       // Previous information about the chat member
	NewChatMember           ChatMember      `json:"new_chat_member"`                       // New information about the chat member
	InviteLink              *ChatInviteLink `json:"invite_link,omitempty"`                 // Optional. Chat invite link, which was used by the user to join the chat; for joining by invite link events only.
	ViaJoinRequest          bool            `json:"via_join_request,omitempty"`            // Optional. True, if the user joined the chat after sending a direct join request without using an invite link and being approved by an administrator
	ViaChatFolderInviteLink bool            `json:"via_chat_folder_invite_link,omitempty"` // Optional. True, if the user joined the chat via a chat folder invite link
}

// ChatMember Contains information about one member of a chat. Currently, the following six types of chat members are supported:
type ChatMember struct {
	User                  *User  `json:"user"`                                // User information about the user
	Status                string `json:"status"`                              // The member's status in the chat. Can be “creator”,“administrator,” “member,” “restricted,” “left” or “kicked”
	CustomTitle           string `json:"custom_title,omitempty"`              // Optional. owner and administrators only. Custom title for this user
	IsAnonymous           bool   `json:"is_anonymous,omitempty"`              // Optional. Owner and administrators only. True, if the user's presence in the chat is hidden
	UntilDate             int64  `json:"until_date,omitempty"`                // Optional. Restricted and kicked only. Date when restrictions will be lifted for this user; unix time.
	CanBeEdited           bool   `json:"can_be_edited,omitempty"`             // Optional. Administrators only. True, if the bot is allowed to edit administrator privileges of that user.
	CanManageChat         bool   `json:"can_manage_chat,omitempty"`           // Optional. Administrators only. True, if the administrator can access the chat event log, chat statistics, message statistics in channels, see channel members, see anonymous administrators in supergroups and ignore slow mode. Implied by any other administrator privilege.
	CanPostMessages       bool   `json:"can_post_messages,omitempty"`         // Optional. administrators only. True, if the administrator can post in the channel; channels only.
	CanEditMessages       bool   `json:"can_edit_messages,omitempty"`         // Optional. administrators only. True, if the administrator can edit messages of other users and can pin messages; channels only.
	CanDeleteMessages     bool   `json:"can_delete_messages,omitempty"`       // Optional. administrators only. True, if the administrator can delete messages of other users.
	CanPostStories        bool   `json:"can_post_stories,omitempty"`          // Optional. True, if the administrator can post stories in the channel; channels only
	CanEditStories        bool   `json:"can_edit_stories,omitempty"`          // Optional. True, if the administrator can edit stories posted by other users; channels only
	CanDeleteStories      bool   `json:"can_delete_stories,omitempty"`        // Optional. True, if the administrator can delete stories posted by other users; channels only
	CanManageVideoChats   bool   `json:"can_manage_video_chats,omitempty"`    // Optional. administrators only. True, if the administrator can manage video chats.
	CanRestrictMembers    bool   `json:"can_restrict_members,omitempty"`      // Optional. administrators only. True, if the administrator can restrict, ban or unban chat members.
	CanPromoteMembers     bool   `json:"can_promote_members,omitempty"`       // Optional. administrators only. True, if the administrator can add new administrators with a subset of their own privileges or demote administrators that he has promoted, directly or indirectly (promoted by administrators that were appointed by the user).
	CanChangeInfo         bool   `json:"can_change_info,omitempty"`           // Optional. administrators and restricted only. True, if the user is allowed to change the chat title, photo and other settings.
	CanInviteUsers        bool   `json:"can_invite_users,omitempty"`          // Optional. administrators and restricted only. True, if the user is allowed to invite new users to the chat.
	CanPinMessages        bool   `json:"can_pin_messages,omitempty"`          // Optional. administrators and restricted only. True, if the user is allowed to pin messages, groups and supergroups only
	IsMember              bool   `json:"is_member,omitempty"`                 // Optional. restricted only. True, if the user is a member of the chat at the moment of the request
	CanSendMessages       bool   `json:"can_send_messages,omitempty"`         // Optional. restricted only. True, if the user is allowed to send text messages, contacts, invoices, locations and venues
	CanSendAudios         bool   `json:"can_send_audios,omitempty"`           // Optional. restricted only. True, if the user is allowed to send audios
	CanSendDocuments      bool   `json:"can_send_documents,omitempty"`        // Optional. restricted only. True, if the user is allowed to send documents
	CanSendPhotos         bool   `json:"can_send_photos,omitempty"`           // Optional. restricted only. True, if the user is allowed to send photos
	CanSendVideos         bool   `json:"can_send_videos,omitempty"`           // Optional. restricted only. True, if the user is allowed to send videos
	CanSendVideoNotes     bool   `json:"can_send_video_notes,omitempty"`      // Optional. restricted only. True, if the user is allowed to send video notes
	CanSendVoiceNotes     bool   `json:"can_send_voice_notes,omitempty"`      // Optional. Restricted only. True, if the user is allowed to send voice notes
	CanSendPolls          bool   `json:"can_send_polls,omitempty"`            // Optional. restricted only. True, if the user is allowed to send polls
	CanSendOtherMessages  bool   `json:"can_send_other_messages,omitempty"`   // Optional. restricted only. True, if the user is allowed to send audio, documents, photos, videos, video notes and voice notes.
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews,omitempty"` // Optional. Restricted only. True, if the user is allowed to add web page previews to their messages.
	CanManageTopics       bool   `json:"can_manage_topics,omitempty"`         // Optional. administrators and restricted only. true, if the user is allowed to create, rename, close, and reopen forum topics; supergroups only
}

// ChatJoinRequest Represents a join request sent to a chat.
type ChatJoinRequest struct {
	Chat       Chat            `json:"chat"`                  // to which the request was sent
	From       User            `json:"from"`                  // that sent the join request
	UserChatId int64           `json:"user_chat_id"`          // Identifier of a private chat with the user who sent the join request. This number may have more than 32 significant bits, and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type is safe for storing this identifier. The bot can use this identifier for 24 hours to send messages until the join request is processed, assuming no other administrator contacted the user.
	Date       int64           `json:"date"`                  // the request was sent in Unix time
	Bio        string          `json:"bio,omitempty"`         // Optional. Bio of the user.
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"` // Optional. Chat invite link that was used by the user to send the join request
}

// ChatPermissions Describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages,omitempty"`         // Optional. True, if the user is allowed to send text messages, contacts, locations and venues
	CanSendAudios         bool `json:"can_send_audios,omitempty"`           // Optional. True, if the user is allowed to send audios
	CanSendDocuments      bool `json:"can_send_documents,omitempty"`        // Optional. True, if the user is allowed to send documents
	CanSendPhotos         bool `json:"can_send_photos,omitempty"`           // Optional. True, if the user is allowed to send photos
	CanSendVideos         bool `json:"can_send_videos,omitempty"`           // Optional. True, if the user is allowed to send videos
	CanSendVideoNotes     bool `json:"can_send_video_notes,omitempty"`      // Optional. True, if the user is allowed to send video notes
	CanSendVoiceNotes     bool `json:"can_send_voice_notes,omitempty"`      // Optional. True, if the user is allowed to send voice notes
	CanSendPolls          bool `json:"can_send_polls,omitempty"`            // Optional. True, if the user is allowed to send polls, it implies can_send_messages
	CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`   // Optional. True, if the user is allowed to send animations, games, stickers and use inline bots, it implies can_send_media_messages
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"` // Optional. True, if the user is allowed to add web page previews to their messages, it implies can_send_media_messages
	CanChangeInfo         bool `json:"can_change_info,omitempty"`           // Optional. True, if the user is allowed to change the chat title, photo and other settings. Ignored in public supergroups
	CanInviteUsers        bool `json:"can_invite_users,omitempty"`          // Optional. True, if the user is allowed to invite new users to the chat
	CanPinMessages        bool `json:"can_pin_messages,omitempty"`          // Optional. True, if the user is allowed to pin messages. Ignored in public supergroups
	CanManageTopics       bool `json:"can_manage_topics,omitempty"`         // Optional. True, if the user is allowed to create forum topics. If omitted, defaults to the value of can_pin_messages
}

// Birthdate //
type Birthdate struct {
	Day   int `json:"day"`            // Day of the user's birth; 1-31
	Month int `json:"month"`          // Month of the user's birth; 1-12
	Year  int `json:"year,omitempty"` // Optional. Year of the user's birth
}

// BusinessIntro //
type BusinessIntro struct {
	Title   string   `json:"title,omitempty"`   // Optional. Title text of the business intro
	Message string   `json:"message,omitempty"` // Optional. Message text of the business intro
	Sticker *Sticker `json:"sticker,omitempty"` // Optional. Sticker of the business intro
}

// BusinessLocation //
type BusinessLocation struct {
	Address  string    `json:"address"`            // Address of the business
	Location *Location `json:"location,omitempty"` // Optional. Location of the business
}

// BusinessOpeningHoursInterval //
type BusinessOpeningHoursInterval struct {
	OpeningMinute int `json:"opening_minute"` // The minute's sequence number in a week, starting on Monday, marking the start of the time interval during which the business is open; 0 - 7 * 24 * 60
	ClosingMinute int `json:"closing_minute"` // The minute's sequence number in a week, starting on Monday, marking the end of the time interval during which the business is open; 0 - 8 * 24 * 60
}

// BusinessOpeningHours //
type BusinessOpeningHours struct {
	TimeZoneName string                          `json:"time_zone_name"` // Unique name of the time zone for which the opening hours are defined
	OpeningHours []*BusinessOpeningHoursInterval `json:"opening_hours"`  // List of time intervals describing business opening hours
}

// ChatLocation Represents a location to which a chat is connected.
type ChatLocation struct {
	Location Location `json:"location"` // The location to which the supergroup is connected. Can't be a live location.
	Address  string   `json:"address"`  // Location address; 1-64 characters, as defined by the chat owner
}

// ReactionType Describes the type of a reaction. Currently, it can be one of
type ReactionType struct {
	ReactionTypeEmoji       ReactionTypeEmoji
	ReactionTypeCustomEmoji ReactionTypeCustomEmoji
}

// ReactionTypeEmoji Reaction is based on an emoji.
type ReactionTypeEmoji struct {
	Type  string `json:"type"`  // Type of the reaction, always “emoji”
	Emoji string `json:"emoji"` // Reaction emoji. Currently, it can be one of "👍", "👎", "❤", "🔥", "🥰", "👏", "😁", "🤔", "🤯", "😱", "🤬", "😢", "🎉", "🤩", "🤮", "💩", "🙏", "👌", "🕊", "🤡", "🥱", "🥴", "😍", "🐳", "❤‍🔥", "🌚", "🌭", "💯", "🤣", "⚡", "🍌", "🏆", "💔", "🤨", "😐", "🍓", "🍾", "💋", "🖕", "😈", "😴", "😭", "🤓", "👻", "👨‍💻", "👀", "🎃", "🙈", "😇", "😨", "🤝", "✍", "🤗", "🫡", "🎅", "🎄", "☃", "💅", "🤪", "🗿", "🆒", "💘", "🙉", "🦄", "😘", "💊", "🙊", "😎", "👾", "🤷‍♂", "🤷", "🤷‍♀", "😡"
}

// ReactionTypeCustomEmoji Reaction is based on a custom emoji.
type ReactionTypeCustomEmoji struct {
	Type        string `json:"type"`         // Type of the reaction, always “custom_emoji”
	CustomEmoji string `json:"custom_emoji"` // Custom emoji identifier
}

// ReactionCount Represents a reaction added to a message along with the number of times it was added.
type ReactionCount struct {
	Type       ReactionType `json:"type"`        // Type of the reaction
	TotalCount int          `json:"total_count"` // Number of times the reaction was added
}

// MessageReactionUpdated Represents a change of a reaction on a message performed by a user.
type MessageReactionUpdated struct {
	Chat        Chat           `json:"chat"`                 // The chat containing the message the user reacted to
	MessageId   int            `json:"message_id"`           // Unique identifier of the message inside the chat
	User        *User          `json:"user,omitempty"`       // Optional. The user that changed the reaction, if the user isn't anonymous
	ActorChat   *Chat          `json:"actor_chat,omitempty"` // Optional. The chat on behalf of which the reaction was changed, if the user is anonymous
	Date        int64          `json:"date"`                 // Date of the change in Unix time
	OldReaction []ReactionType `json:"old_reaction"`         // Previous list of reaction types that were set by the user
	NewReaction []ReactionType `json:"new_reaction"`         // New list of reaction types that have been set by the user
}

// MessageReactionCountUpdated Represents reaction changes on a message with anonymous reactions.
type MessageReactionCountUpdated struct {
	Chat      Chat            `json:"chat"`       // The chat containing the message
	MessageId int             `json:"message_id"` // Unique message identifier inside the chat
	Date      int64           `json:"date"`       // Date of the change in Unix time
	Reactions []ReactionCount `json:"reactions"`  // List of reactions that are present on the message
}

// ForumTopic Represents a forum topic.
type ForumTopic struct {
	MessageThreadID   int64  `json:"message_thread_id"`              // Unique identifier of the forum topic
	Name              string `json:"name"`                           // Name of the topic
	IconColor         int    `json:"icon_color"`                     // Color of the topic icon in RGB format
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"` // Optional. Unique identifier of the custom emoji shown as the topic icon
}

// BotCommand Represents a bot command.
type BotCommand struct {
	Command     string `json:"command"`     // Text of the command; 1-32 characters. Can contain only lowercase English letters, digits and underscores.
	Description string `json:"description"` // Description of the command; 1-256 characters.
}

// BotCommandScope Represents the scope to which bot commands are applied. Currently, the following seven scopes are supported.
type BotCommandScope struct {
	Type   string `json:"type"`
	ChatID int64  `json:"chat_id"`
	UserID int64  `json:"user_id"`
}

// BotCommandScopeDefault Represents the default scope of bot commands.
// Default commands are used if no commands with a narrower scope are specified for the user.
type BotCommandScopeDefault struct {
	Type string `json:"type"` // Scope type must be default
}

// BotCommandScopeAllPrivateChats Represents the scope of bot commands, covering all private chats.
type BotCommandScopeAllPrivateChats struct {
	Type string `json:"type"` // Scope type, must be all_private_chats
}

// BotCommandScopeAllGroupChats Represents the scope of bot commands, covering all group and supergroup chats.
type BotCommandScopeAllGroupChats struct {
	Type string `json:"type"` // Scope type, must be all_group_chats
}

// BotCommandScopeAllChatAdministrators Represents the scope of bot commands, covering all group and supergroup chat administrators.
type BotCommandScopeAllChatAdministrators struct {
	Type string `json:"type"` // Scope type, must be all_chat_administrators
}

// BotCommandScopeChat Represents the scope of bot commands, covering a specific chat.
type BotCommandScopeChat struct {
	Type   string `json:"type"`    // Scope type must be chat
	ChatID string `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroup username)
}

// BotCommandScopeChatAdministrators Represents the scope of bot commands, covering all administrators of a specific group or supergroup chat.
type BotCommandScopeChatAdministrators struct {
	Type   string `json:"type"`    // Scope type, must be chat_administrators
	ChatID string `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroup username)
}

// BotCommandScopeChatMember Represents the scope of bot commands, covering a specific member of a group or supergroup chat.
type BotCommandScopeChatMember struct {
	Type   string `json:"type"`    // Scope type, must be chat_member
	ChatID string `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroup username)
	UserID int64  `json:"user_id"` // Unique identifier of the target user
}

// BotName represents the bot's name.
type BotName struct {
	Name string `json:"name"` // The bot's name
}

// BotDescription represents the bot's description.
type BotDescription struct {
	Description string `json:"description"` // The bot's description
}

// BotShortDescription represents the bot's short description.
type BotShortDescription struct {
	ShortDescription string `json:"short_description"` // The bot's short description
}

// MenuButton describes the bot's menu button in a private chat.
// If a menu button other than MenuButtonDefault is set for a private chat, then it is applied in the chat.
// Otherwise, the default menu button is applied.
// By default, the menu button opens the list of bot commands.
type MenuButton struct {
	Type   string      `json:"type"`    // Type of the button must be commands
	Text   string      `json:"text"`    // Text on the button
	WebApp *WebAppInfo `json:"web_app"` // Description of the Web App that will be launched when the user presses the button. The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery.
}

// MenuButtons describes the bot's menu button in a private chat.
// If a menu button other than MenuButtonDefault is set for a private chat, then it is applied in the chat.
// Otherwise, the default menu button is applied.
// By default, the menu button opens the list of bot commands.
type MenuButtons struct {
	MenuButtonCommands MenuButtonCommands
	MenuButtonWebApp   MenuButtonWebApp
	MenuButtonDefault  MenuButtonDefault
}

// MenuButtonCommands Represents a menu button, which opens the bot's list of commands.
type MenuButtonCommands struct {
	Type string `json:"type"` // Type of the button must be commands
}

// MenuButtonWebApp Represents a menu button, which launches a Web App.
type MenuButtonWebApp struct {
	Type   string      `json:"type"`    // The Type of the button, must be web_app
	Text   string      `json:"text"`    // Text on the button
	WebApp *WebAppInfo `json:"web_app"` // Description of the Web App that will be launched when the user presses the button. The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery.
}

// MenuButtonDefault Describes that no specific value for the menu button was set.
type MenuButtonDefault struct {
	Type string `json:"type"` // Type of the button must be default
}

// ChatBoostSource Describes the source of a chat boost. It can be one of
type ChatBoostSource struct {
	ChatBoostSourcePremium
	ChatBoostSourceGiftCode
	ChatBoostSourceGiveaway
}

// ChatBoostSourcePremium Obtained by subscribing to Telegram Premium or by gifting a Telegram Premium subscription to another user.
type ChatBoostSourcePremium struct {
	Source string `json:"source"` // Source of the boost, always “premium”
	User   User   `json:"user"`   // User that boosted the chat
}

// ChatBoostSourceGiftCode Obtained by the creation of Telegram Premium gift codes to boost a chat. Each such code boosts the chat 4 times for the duration of the corresponding Telegram Premium subscription.
type ChatBoostSourceGiftCode struct {
	Source string `json:"source"` // Source of the boost, always “gift_code”
	User   User   `json:"user"`   // User for which the gift code was created
}

// ChatBoostSourceGiveaway Obtained by the creation of a Telegram Premium giveaway. This boosts the chat 4 times for the duration of the corresponding Telegram Premium subscription.
type ChatBoostSourceGiveaway struct {
	Source            string `json:"source"`                 // Source of the boost, always “giveaway”
	GiveawayMessageId int    `json:"giveaway_message_id"`    // Identifier of a message in the chat with the giveaway; the message could have been deleted already. May be 0 if the message isn't sent yet.
	User              *User  `json:"user,omitempty"`         // Optional. User that won the prize in the giveaway if any
	IsUnclaimed       bool   `json:"is_unclaimed,omitempty"` // Optional. True, if the giveaway was completed, but there was no user to win the prize
}

// ChatBoost Represents a boost added to a chat or changed.
type ChatBoost struct {
	BoostID        string          `json:"boost_id"`        // Unique identifier of the boost
	AddDate        int64           `json:"add_date"`        // Point in time (Unix timestamp) when the chat was boosted
	ExpirationDate int64           `json:"expiration_date"` // Point in time (Unix timestamp) when the boost will automatically expire, unless the booster's Telegram Premium subscription is prolonged
	Source         ChatBoostSource `json:"source"`          // Source of the added boost
}

// ChatBoostUpdated Represents a boost added to a chat or changed.
type ChatBoostUpdated struct {
	Chat  Chat      `json:"chat"`  // Chat which was boosted
	Boost ChatBoost `json:"boost"` // Infomation about the chat boost
}

// ChatBoostRemoved Represents a boost removed from a chat.
type ChatBoostRemoved struct {
	Chat       Chat            `json:"chat"`        // Chat which was boosted
	BoostId    string          `json:"boost_id"`    // Unique identifier of the boost
	RemoveDate int64           `json:"remove_date"` // Point in time (Unix timestamp) when the boost was removed
	Source     ChatBoostSource `json:"source"`      // Source of the removed boost
}

// UserChatBoosts Represents a boost removed from a chat.
type UserChatBoosts struct {
	Boosts []ChatBoost `json:"boosts"` // The list of boosts added to the chat by the user
}

// BusinessConnection Describes the connection of the bot with a business account.
type BusinessConnection struct {
	ID         string `json:"id"`           // Unique identifier of the business connection
	User       User   `json:"user"`         // Business account user that created the business connection
	UserChatId int64  `json:"user_chat_id"` // Identifier of a private chat with the user who created the business connection. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	Date       int64  `json:"date"`         // Date the connection was established in Unix time
	CanReply   bool   `json:"can_reply"`    // True, if the bot can act on behalf of the business account in chats that were active in the last 24 hours
	IsEnabled  bool   `json:"is_enabled"`   // True, if the connection is active
}

// BusinessMessagesDeleted Received when messages are deleted from a connected business account.
type BusinessMessagesDeleted struct {
	BusinessConnectionId string `json:"business_connection_id"` // Unique identifier of the business connection
	Chat                 Chat   `json:"chat"`                   // Information about a chat in the business account. The bot may not have access to the chat or the corresponding user.
	MessageIds           []int  `json:"message_ids"`            // A JSON-serialized list of identifiers of deleted messages in the chat of the business account
}

// ResponseParameters are various errors that can be returned in APIResponse.
type ResponseParameters struct {
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"` // Optional. The group has been migrated to a supergroup with the specified identifier. This number may have more than 32 significant bits, and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type is safe for storing this identifier.
	RetryAfter      int   `json:"retry_after,omitempty"`        // Optional. In case of exceeding flood control, the number of seconds left to wait before the request can be repeated
}

// InputMedia Represents the content of a media message to be sent. It should be one of
type InputMedia struct {
	InputMediaAnimation InputMediaAnimation
	InputMediaDocument  InputMediaDocument
	InputMediaAudio     InputMediaAudio
	InputMediaPhoto     InputMediaPhoto
	InputMediaVideo     InputMediaVideo
}

// InputMediaPhoto Represents a photo to be sent.
type InputMediaPhoto struct {
	Type                  string           `json:"type"`                               // Type of the result must be photoed
	Media                 RequestFileData  `json:"media"`                              // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name.
	Caption               string           `json:"caption,omitempty"`                  // Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	ParseMode             string           `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the photo caption. See [formatting options](https://core.telegram.org/bots/api#formatting-options) for more details.
	CaptionEntities       []*MessageEntity `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool             `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
	HasSpoiler            bool             `json:"has_spoiler,omitempty"`              // Optional. Pass True if the photo needs to be covered with a spoiler animation
}

// InputMediaVideo Represents a video to be sent.
type InputMediaVideo struct {
	Type                  string           `json:"type"`                               // Type of the result must be video
	Media                 RequestFileData  `json:"media"`                              // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name.
	Thumbnail             RequestFileData  `json:"thumb,thumbnail,omitempty"`          // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	Caption               string           `json:"caption,omitempty"`                  // Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	ParseMode             string           `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the video caption. See [formatting options](https://core.telegram.org/bots/api#formatting-options) for more details.
	CaptionEntities       []*MessageEntity `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool             `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
	Width                 int              `json:"width,omitempty"`                    // Optional. Video width
	Height                int              `json:"height,omitempty"`                   // Optional. Video height
	Duration              int              `json:"duration,omitempty"`                 // Optional. Video duration in seconds
	SupportsStreaming     bool             `json:"supports_streaming,omitempty"`       // Optional. Pass True, if the uploaded video is suitable for streaming
	HasSpoiler            bool             `json:"has_spoiler,omitempty"`              // Optional. Pass True if the video needs to be covered with a spoiler animation
}

// InputMediaAnimation Represents an animation file (GIF or H.264/MPEG-4 AVC video without a sound) to be sent.
type InputMediaAnimation struct {
	Type                  string           `json:"type"`                               // The Type of the result must be animation
	Media                 RequestFileData  `json:"media"`                              // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name.
	Thumbnail             RequestFileData  `json:"thumb,thumbnail,omitempty"`          // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	Caption               string           `json:"caption,omitempty"`                  // Optional. Caption of the animation to be sent, 0-1024 characters after entities parsing
	ParseMode             string           `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the animation caption. See [formatting options](https://core.telegram.org/bots/api#formatting-options) for more details.
	CaptionEntities       []*MessageEntity `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool             `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
	Width                 int              `json:"width,omitempty"`                    // Optional. Animation width
	Height                int              `json:"height,omitempty"`                   // Optional. Animation height
	Duration              int              `json:"duration,omitempty"`                 // Optional. Animation duration in seconds
	HasSpoiler            bool             `json:"has_spoiler,omitempty"`              // Optional. Pass True if the animation needs to be covered with a spoiler animation
}

// InputMediaAudio Represents an audio file to be treated as music to be sent.
type InputMediaAudio struct {
	Type            string           `json:"type"`                       // The Type of the result must be audio
	Media           RequestFileData  `json:"media"`                      // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name.
	Thumbnail       RequestFileData  `json:"thumb,thumbnail,omitempty"`  // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	Caption         string           `json:"caption,omitempty"`          // Optional. Caption of the audio to be sent, 0-1024 characters after entities parsing
	ParseMode       string           `json:"parse_mode,omitempty"`       // Optional. Mode for parsing entities in the audio caption. See [formatting options](https://core.telegram.org/bots/api#formatting-options) for more details.
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"` // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	Duration        int              `json:"duration,omitempty"`         // Optional. Duration of the audio in seconds
	Performer       string           `json:"performer,omitempty"`        // Optional. Performer of the audio
	Title           string           `json:"title,omitempty"`            // Optional. Title of the audio
}

// InputMediaDocument Represents a general file to be sent.
type InputMediaDocument struct {
	Type                        string           `json:"type"`                                     // The Type of the result must be documented
	Media                       RequestFileData  `json:"media"`                                    // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name.
	Thumbnail                   RequestFileData  `json:"thumb,thumbnail,omitempty"`                // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	Caption                     string           `json:"caption,omitempty"`                        // Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	ParseMode                   string           `json:"parse_mode,omitempty"`                     // Optional. Mode for parsing entities in the document caption. See [formatting options](https://core.telegram.org/bots/api#formatting-options) for more details.
	CaptionEntities             []*MessageEntity `json:"caption_entities,omitempty"`               // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	DisableContentTypeDetection bool             `json:"disable_content_type_detection,omitempty"` // Optional. Disables automatic server-side content type detection for files uploaded using multipart/form-data. Always True, if the document is sent as part of an album.
}

// Sticker Represents a sticker.
type Sticker struct {
	FileID           string        `json:"file_id"`                     // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID     string        `json:"file_unique_id"`              // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Type             string        `json:"type"`                        // Type of the sticker, currently one of “regular,” “masks,” “custom_emoji.” The type of the sticker is independent of its format, which is determined by the fields is_animated and is_video.
	Width            int           `json:"width"`                       // Sticker width
	Height           int           `json:"height"`                      // Sticker height
	IsAnimated       bool          `json:"is_animated"`                 // True, if the sticker is animated
	IsVideo          bool          `json:"is_video"`                    // True, if the sticker is a video sticker
	Thumbnail        *PhotoSize    `json:"thumb,thumbnail,omitempty"`   // Optional. Sticker thumbnail in the .WEBP or .JPG format
	Emoji            string        `json:"emoji,omitempty"`             // Optional. Emoji associated with the sticker
	SetName          string        `json:"set_name,omitempty"`          // Optional. Name of the sticker set to which the sticker belongs
	PremiumAnimation *File         `json:"premium_animation,omitempty"` // Optional. Premium animation for the sticker, if the sticker is premium
	MaskPosition     *MaskPosition `json:"mask_position,omitempty"`     // Optional. For mask stickers, the position where the mask should be placed
	CustomEmojiID    string        `json:"custom_emoji_id,omitempty"`   // Optional. For custom emoji stickers, unique identifier of the custom emoji
	NeedsRepainting  bool          `json:"needs_repainting,omitempty"`  // Optional. True, if the sticker must be repainted to a text color in messages, the color of the Telegram Premium badge in emoji status, white color on chat photos, or another appropriate color in other places
	FileSize         int64         `json:"file_size,omitempty"`         // Optional. File size in bytes
}

// StickerSet Represents a sticker set.
type StickerSet struct {
	Name        string     `json:"name"`                      // Sticker set name
	Title       string     `json:"title"`                     // Sticker set title
	StickerType string     `json:"sticker_type"`              // Type of stickers in the set, currently one of “regular,” “masks,” “custom_emoji”
	Stickers    []Sticker  `json:"stickers"`                  // List of all set stickers
	Thumbnail   *PhotoSize `json:"thumb,thumbnail,omitempty"` // Optional. Sticker set thumbnail in the .WEBP, .TGS, or .WEBM format
}

// MaskPosition Describes the position on faces where a mask should be placed by default. by default.
type MaskPosition struct {
	Point  string  `json:"point"`   // The part of the face relative to which the mask should be placed. One of “foreheads,” “eyes,” “mouth,” or “chin.”
	XShift float64 `json:"x_shift"` // Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. For example, choosing -1.0 will place mask just to the left of the default mask position.
	YShift float64 `json:"y_shift"` // Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. For example, 1.0 will place the mask just below the default mask position.
	Scale  float64 `json:"scale"`   // Mask scaling coefficient. For example, 2.0 means double size.
}

// InputSticker Describes a sticker to be added to a sticker set.
type InputSticker struct {
	Sticker      RequestFileData `json:"sticker"`       // The added sticker. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. Animated and video stickers can't be uploaded via HTTP URL.
	Format       string          `json:"format"`        // Format of the added sticker, must be one of “static” for a .WEBP or .PNG image, “animated” for a .TGS animation, “video” for a WEBM video
	EmojiList    []string        `json:"emoji_list"`    // List of 1-20 emoji associated with the sticker
	MaskPosition *MaskPosition   `json:"mask_position"` // Optional. Position where the mask should be placed on faces. For “mask” stickers only.
	Keywords     []string        `json:"keywords"`      // Optional. List of 0-20 search keywords for the sticker with total length of up to 64 characters. For “regular” and “custom_emoji” stickers only.
}

// InlineQuery Represents an incoming inline query. When the user sends an empty query, your bot could return some default or trending results.
type InlineQuery struct {
	ID       string    `json:"id"`                  // Unique identifier for this query
	From     User      `json:"from"`                // Sender
	Query    string    `json:"query"`               // Text of the query (up to 256 characters)
	Offset   string    `json:"offset"`              // the bot can control the offset of the results to be returned.
	ChatType string    `json:"chat_type,omitempty"` // Optional. Type of the chat from which the inline query was sent. Can be either “sender” for a private chat with the inline query sender, “private,” “group,” “supergroup,” or “channel.” The chat type should always be known for requests sent from official clients and most third-party clients, unless the request was sent from a secret chat
	Location *Location `json:"location,omitempty"`  // Optional. Sender location, only for bots that request user location
}

// InlineQueryResultsButton Represents a button to be shown above inline query results. You must use exactly one of the optional fields.
type InlineQueryResultsButton struct {
	Text           string      `json:"text"`                      // Label text on the button
	WebApp         *WebAppInfo `json:"web_app,omitempty"`         // Optional. Description of the Web App that will be launched when the user presses the button. The Web App will be able to switch back to the inline mode using the method web_app_switch_inline_query inside the Web App.
	StartParameter *string     `json:"start_parameter,omitempty"` // Optional. Deep-linking parameter for the /start message sent to the bot when a user presses the button. 1-64 characters, only A-Z, a-z, 0-9, _ and - are allowed. Example: An inline bot that sends YouTube videos can ask the user to connect the bot to their YouTube account to adapt search results accordingly. To do this, it displays a 'Connect your YouTube account' button above the results, or even before shows any. The user presses the button, switches to a private chat with the bot and, in doing so, passes a start parameter that instructs the bot to return an OAuth link. Once done, the bot can offer a switch_inline button so that the user can easily return to the chat where they wanted to use the bot's inline capabilities.
}

// InlineQueryResult Represents one result of an inline query. Telegram clients currently support results of the following 20 types
// Note: All URLs passed in inline query results will be available to end users and therefore must be assumed to be public.
type InlineQueryResult struct {
	InlineQueryResultArticle        InlineQueryResultArticle
	InlineQueryResultPhoto          InlineQueryResultPhoto
	InlineQueryResultGif            InlineQueryResultGIF
	InlineQueryResultMpeg4Gif       InlineQueryResultMPEG4GIF
	InlineQueryResultVideo          InlineQueryResultVideo
	InlineQueryResultAudio          InlineQueryResultAudio
	InlineQueryResultVoice          InlineQueryResultVoice
	InlineQueryResultDocument       InlineQueryResultDocument
	InlineQueryResultLocation       InlineQueryResultLocation
	InlineQueryResultVenue          InlineQueryResultVenue
	InlineQueryResultContact        InlineQueryResultContact
	InlineQueryResultGame           InlineQueryResultGame
	InlineQueryResultCachedPhoto    InlineQueryResultCachedPhoto
	InlineQueryResultCachedGif      InlineQueryResultCachedGIF
	InlineQueryResultCachedMpeg4Gif InlineQueryResultCachedMPEG4GIF
	InlineQueryResultCachedSticker  InlineQueryResultCachedSticker
	InlineQueryResultCachedDocument InlineQueryResultCachedDocument
	InlineQueryResultCachedVideo    InlineQueryResultCachedVideo
	InlineQueryResultCachedVoice    InlineQueryResultCachedVoice
	InlineQueryResultCachedAudio    InlineQueryResultCachedAudio
}

// InlineQueryResultArticle Represents a link to an article or web page.
type InlineQueryResultArticle struct {
	Type                string                `json:"type"`                                    // the type of the result must be article
	ID                  string                `json:"id"`                                      // Unique identifier for this result, 1-64 Bytes
	Title               string                `json:"title"`                                   // title of the result
	InputMessageContent any                   `json:"input_message_content"`                   // Content of the message to be sent
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`                  // Optional. Inline keyboard attached to the message
	URL                 string                `json:"url,omitempty"`                           // Optional. URL of the result
	HideURL             bool                  `json:"hide_url,omitempty"`                      // Optional. Pass True, if you don't want the URL to be shown in the message
	Description         string                `json:"description,omitempty"`                   // Optional. Short description of the result
	ThumbnailURL        string                `json:"thumb_url,thumbnail_url,omitempty"`       // Optional. Url of the thumbnail for the result
	ThumbnailWidth      int                   `json:"thumb_width,thumbnail_width,omitempty"`   // Optional. Thumbnail width
	ThumbnailHeight     int                   `json:"thumb_height,thumbnail_height,omitempty"` // Optional. Thumbnail height
}

// InlineQueryResultPhoto Represents a link to a photo.
// By default, this photo will be sent by the user with optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultPhoto struct {
	Type                  string                `json:"type"`                               // type of the result must be photoed
	ID                    string                `json:"id"`                                 // Unique identifier for this result, 1-64 bytes
	URL                   string                `json:"photo_url"`                          // A valid URL of the photo. Photo must be in JPEG format. Photo size must not exceed 5MB
	ThumbnailURL          string                `json:"thumb_url,thumbnail_url"`            // url of the thumbnail for the photo
	Width                 int                   `json:"photo_width,omitempty"`              // Optional. Width of the photo
	Height                int                   `json:"photo_height,omitempty"`             // Optional. Height of the photo
	Title                 string                `json:"title"`                              // title of the result
	Description           string                `json:"description,omitempty"`              // Optional. Short description of the result
	Caption               string                `json:"caption,omitempty"`                  // Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	ParseMode             string                `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the photo caption. See [formatting options](https://core.telegram.org/bots/api#formatting-options) for more details.
	CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // Optional. Inline keyboard attached to the message
	InputMessageContent   any                   `json:"input_message_content,omitempty"`    // Optional. Content of the message to be sent instead of the photo
}

// InlineQueryResultGIF Represents a link to an animated GIF file.
// By default, this animated GIF file will be sent by the user with optional caption.
// Alternatively,
// you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultGIF struct {
	Type                  string                `json:"type"`                                          // the type of the result must be gif
	ID                    string                `json:"id"`                                            // Unique identifier for this result, 1-64 bytes
	URL                   string                `json:"gif_url"`                                       // A valid URL for the GIF file. File size must not exceed 1MB
	Width                 int                   `json:"gif_width,omitempty"`                           // Optional. Width of the GIF
	Height                int                   `json:"gif_height,omitempty"`                          // Optional. Height of the GIF
	Duration              int                   `json:"gif_duration,omitempty"`                        // Optional. Duration of the GIF in seconds
	ThumbnailURL          string                `json:"thumb_url,thumbnail_url"`                       // url of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
	ThumbnailMimeType     string                `json:"thumb_mime_type,thumbnail_mime_type,omitempty"` // Optional. MIME type of the thumbnail must be one of “image/jpeg,” “image/gif,” or “video/mp4.” Defaults to “image/jpeg”
	Title                 string                `json:"title,omitempty"`                               // Optional. Title for the result
	Caption               string                `json:"caption,omitempty"`                             // Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
	ParseMode             string                `json:"parse_mode,omitempty"`                          // Optional. Mode for parsing entities in the caption. See [formatting options](https://core.telegram.org/bots/api#formatting-options) for more details.
	CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`                    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`            // Optional. Pass True, if the caption must be shown above the message media
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`                        // Optional. Inline keyboard attached to the message
	InputMessageContent   any                   `json:"input_message_content,omitempty"`               // Optional. Content of the message to be sent instead of the GIF animation
}

// InlineQueryResultMPEG4GIF Represents a link to a video animation (H.264/MPEG-4 AVC video without a sound).
// By default, this animated MPEG-4 file will be sent by the user with optional caption.
// Alternatively,
// you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultMPEG4GIF struct {
	Type                  string                `json:"type"`                                          // type of the result must be mpeg4_gif
	ID                    string                `json:"id"`                                            // Unique identifier for this result, 1-64 bytes
	URL                   string                `json:"mpeg4_url"`                                     // A valid URL for the MPEG4 file. File size must not exceed 1MB
	Width                 int                   `json:"mpeg4_width,omitempty"`                         // Optional. Video width
	Height                int                   `json:"mpeg4_height,omitempty"`                        // Optional. Video height
	Duration              int                   `json:"mpeg4_duration,omitempty"`                      // Optional. Video duration in seconds
	ThumbnailURL          string                `json:"thumb_url,thumbnail_url"`                       // url of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
	ThumbnailMimeType     string                `json:"thumb_mime_type,thumbnail_mime_type,omitempty"` // Optional. MIME type of the thumbnail must be one of “image/jpeg,” “image/gif,” or “video/mp4.” Defaults to “image/jpeg”
	Title                 string                `json:"title,omitempty"`                               // Optional. Title for the result
	Caption               string                `json:"caption,omitempty"`                             // Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
	ParseMode             string                `json:"parse_mode,omitempty"`                          // Optional. Mode for parsing entities in the caption. See [formatting options](https://core.telegram.org/bots/api#formatting-options) for more details.
	CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`                    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`            // Optional. Pass True, if the caption must be shown above the message media
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`                        // Optional. Inline keyboard attached to the message
	InputMessageContent   any                   `json:"input_message_content,omitempty"`               // Optional. Content of the message to be sent instead of the video animation
}

// InlineQueryResultVideo Represents a link to a page containing an embedded video player or a video file.
// By default, this video file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
// If an InlineQueryResultVideo message contains an embedded video (e.g., YouTube),
// you must replace its content using input_message_content.
type InlineQueryResultVideo struct {
	Type                  string                `json:"type"`                               // type of the result must be video
	ID                    string                `json:"id"`                                 // Unique identifier for this result, 1-64 bytes
	URL                   string                `json:"video_url"`                          // A valid URL for the embedded video player or video file
	MimeType              string                `json:"mime_type"`                          // MIME type of the content of the video URL, “text/html” or “video/mp4”
	ThumbnailURL          string                `json:"thumb_url,thumbnail_url"`            // url of the thumbnail (JPEG only) for the video
	Title                 string                `json:"title"`                              // title for the result
	Caption               string                `json:"caption,omitempty"`                  // Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	ParseMode             string                `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the video caption. See [formatting options](https://core.telegram.org/bots/api#formatting-options) for more details.
	CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	Width                 int                   `json:"video_width,omitempty"`              // Optional. Video width
	Height                int                   `json:"video_height,omitempty"`             // Optional. Video height
	Duration              int                   `json:"video_duration,omitempty"`           // Optional. Video duration in seconds
	Description           string                `json:"description,omitempty"`              // Optional. Short description of the result
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // Optional. Inline keyboard attached to the message
	InputMessageContent   any                   `json:"input_message_content,omitempty"`    // Optional. Content of the message to be sent instead of the video. This field is required if InlineQueryResultVideo is used to send an HTML-page as a result (e.g., a YouTube video).
}

// InlineQueryResultAudio Represents a link to an MP3 audio file.
// By default, this audio file will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
// Note: This will only work in Telegram versions released after 9 April 2016.
// Older clients will ignore them.
type InlineQueryResultAudio struct {
	Type                string                `json:"type"`                            // the type of the result must be audio
	ID                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	URL                 string                `json:"audio_url"`                       // A valid URL for the audio file
	Title               string                `json:"title"`                           // title
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption, 0-1024 characters after entities parsing
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the audio caption. See [formatting options](https://core.telegram.org/bots/api#formatting-options) for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	Performer           string                `json:"performer,omitempty"`             // Optional. Performer
	Duration            int                   `json:"audio_duration,omitempty"`        // Optional. Audio duration in seconds
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent any                   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the audio
}

// InlineQueryResultVoice Represents a link to a voice recording in an .OGG container encoded with OPUS.
// By default, this voice recording will be sent by the user.
// Alternatively,
// you can use input_message_content to send a message with the specified content instead of the voice message.
// Note: This will only work in Telegram versions released after 9 April 2016.
// Older clients will ignore them.
type InlineQueryResultVoice struct {
	Type                string                `json:"type"`                            // the type of the result must be voice
	ID                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	URL                 string                `json:"voice_url"`                       // A valid URL for the voice recording
	Title               string                `json:"title"`                           // Recording title
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption, 0-1024 characters after entities parsing
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the voice message caption. See [formatting options](https://core.telegram.org/bots/api#formatting-options) for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	Duration            int                   `json:"voice_duration,omitempty"`        // Optional. Recording duration in seconds
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent any                   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the voice recording
}

// InlineQueryResultDocument Represents a link to a file.
// By default, this file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the file.
// Currently, only .PDF and .ZIP files can be sent using this method.
// Note: This will only work in Telegram versions released after 9 April 2016.
// Older clients will ignore them.
type InlineQueryResultDocument struct {
	Type                string                `json:"type"`                                    // the type of the result must be documented
	ID                  string                `json:"id"`                                      // Unique identifier for this result, 1-64 bytes
	Title               string                `json:"title"`                                   // title for the result
	Caption             string                `json:"caption,omitempty"`                       // Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	ParseMode           string                `json:"parse_mode,omitempty"`                    // Optional. Mode for parsing entities in the document caption. See [formatting options](https://core.telegram.org/bots/api#formatting-options) for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`              // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	URL                 string                `json:"document_url"`                            // A valid URL for the file
	MimeType            string                `json:"mime_type"`                               // MIME type of the content of the file, either “application/pdf” or “application/zip”
	Description         string                `json:"description,omitempty"`                   // Optional. Short description of the result
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`                  // Optional. Inline keyboard attached to the message
	InputMessageContent any                   `json:"input_message_content,omitempty"`         // Optional. Content of the message to be sent instead of the file
	ThumbnailURL        string                `json:"thumb_url,thumbnail_url,omitempty"`       // Optional. URL of the thumbnail (JPEG only) for the file
	ThumbnailWidth      int                   `json:"thumb_width,thumbnail_width,omitempty"`   // Optional. Thumbnail width
	ThumbnailHeight     int                   `json:"thumb_height,thumbnail_height,omitempty"` // Optional. Thumbnail height
}

// InlineQueryResultLocation Represents a location on a map.
// By default, the location will be sent by the user.
// Alternatively,
// you can use input_message_content to send a message with the specified content instead of the location.
// Note: This will only work in Telegram versions released after 9 April 2016.
// Older clients will ignore them.
type InlineQueryResultLocation struct {
	Type                 string                `json:"type"`                                    // type of the result must be location
	ID                   string                `json:"id"`                                      // Unique identifier for this result, 1-64 Bytes
	Latitude             float64               `json:"latitude"`                                // Location latitude in degrees
	Longitude            float64               `json:"longitude"`                               // Location longitude in degrees
	Title                string                `json:"title"`                                   // Location title
	HorizontalAccuracy   float64               `json:"horizontal_accuracy,omitempty"`           // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod           int                   `json:"live_period,omitempty"`                   // Optional. The Period in seconds for which the location can be updated should be between 60 and 86400.
	Heading              int                   `json:"heading,omitempty"`                       // Optional. For live locations, the direction in which the user is moving in degrees. It Must be between 1 and 360 if specified.
	ProximityAlertRadius int                   `json:"proximity_alert_radius,omitempty"`        // Optional. For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. It Must be between 1 and 100000 if specified.
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`                  // Optional. Inline keyboard attached to the message
	InputMessageContent  any                   `json:"input_message_content,omitempty"`         // Optional. Content of the message to be sent instead of the location
	ThumbnailURL         string                `json:"thumb_url,thumbnail_url,omitempty"`       // Optional. Url of the thumbnail for the result
	ThumbnailWidth       int                   `json:"thumb_width,thumbnail_width,omitempty"`   // Optional. Thumbnail width
	ThumbnailHeight      int                   `json:"thumb_height,thumbnail_height,omitempty"` // Optional. Thumbnail height
}

// InlineQueryResultVenue Represents a venue.
// By default, the venue will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the venue.
// Note: This will only work in Telegram versions released after 9 April 2016.
// Older clients will ignore them.
type InlineQueryResultVenue struct {
	Type                string                `json:"type"`                                    // type of the result must be venue
	ID                  string                `json:"id"`                                      // Unique identifier for this result, 1-64 Bytes
	Latitude            float64               `json:"latitude"`                                // latitude of the venue location in degrees
	Longitude           float64               `json:"longitude"`                               // longitude of the venue location in degrees
	Title               string                `json:"title"`                                   // title of the venue
	Address             string                `json:"address"`                                 // address of the venue
	FoursquareID        string                `json:"foursquare_id,omitempty"`                 // Optional. Foursquare identifier of the venue if known
	FoursquareType      string                `json:"foursquare_type,omitempty"`               // Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default,” “arts_entertainment/aquarium” or “food/ice cream.”)
	GooglePlaceID       string                `json:"google_place_id,omitempty"`               // Optional. Google Places identifier of the venue
	GooglePlaceType     string                `json:"google_place_type,omitempty"`             // Optional. Google Places a type of the venue. See [supported types](https://developers.google.com/places/web-service/supported_types).
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`                  // Optional. Inline keyboard attached to the message
	InputMessageContent any                   `json:"input_message_content,omitempty"`         // Optional. Content of the message to be sent instead of the venue
	ThumbnailURL        string                `json:"thumb_url,thumbnail_url,omitempty"`       // Optional. Url of the thumbnail for the result
	ThumbnailWidth      int                   `json:"thumb_width,thumbnail_width,omitempty"`   // Optional. Thumbnail width
	ThumbnailHeight     int                   `json:"thumb_height,thumbnail_height,omitempty"` // Optional. Thumbnail height
}

// InlineQueryResultContact Represents a contact with a phone number.
// By default, this contact will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the contact.
// Note: This will only work in Telegram versions released after 9 April 2016.
// Older clients will ignore them.
type InlineQueryResultContact struct {
	Type                string                `json:"type"`                                    // type of the result must be contact
	ID                  string                `json:"id"`                                      // Unique identifier for this result, 1-64 Bytes
	PhoneNumber         string                `json:"phone_number"`                            // contact's phone number
	FirstName           string                `json:"first_name"`                              // contact's first name
	LastName            string                `json:"last_name,omitempty"`                     // Optional. Contact's last name
	VCard               string                `json:"vcard,omitempty"`                         // Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`                  // Optional. Inline keyboard attached to the message
	InputMessageContent any                   `json:"input_message_content,omitempty"`         // Optional. Content of the message to be sent instead of the contact
	ThumbnailURL        string                `json:"thumb_url,thumbnail_url,omitempty"`       // Optional. Url of the thumbnail for the result
	ThumbnailWidth      int                   `json:"thumb_width,thumbnail_width,omitempty"`   // Optional. Thumbnail width
	ThumbnailHeight     int                   `json:"thumb_height,thumbnail_height,omitempty"` // Optional. Thumbnail height
}

// InlineQueryResultGame Represents a Game.
// Note: This will only work in Telegram versions released after October 1, 2016.
// Older clients will not display any inline results if a game result is among them.
type InlineQueryResultGame struct {
	Type          string                `json:"type"`                   // type of the result must be game
	ID            string                `json:"id"`                     // Unique identifier for this result, 1-64 bytes
	GameShortName string                `json:"game_short_name"`        // Short name of the game
	ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup,omitempty"` // Optional. Inline keyboard attached to the message
}

// InlineQueryResultCachedPhoto Represents a link to a photo stored on the Telegram servers.
// By default, this photo will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultCachedPhoto struct {
	Type                  string                `json:"type"`                               // type of the result must be photoed
	ID                    string                `json:"id"`                                 // Unique identifier for this result, 1-64 bytes
	PhotoID               string                `json:"photo_file_id"`                      // A valid file identifier of the photo
	Title                 string                `json:"title,omitempty"`                    // Optional. Title for the result
	Description           string                `json:"description,omitempty"`              // Optional. Short description of the result
	Caption               string                `json:"caption,omitempty"`                  // Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	ParseMode             string                `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the photo caption. See [formatting options](https://core.telegram.org/bots/api#formatting-options) for more details.
	CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // Optional. Inline keyboard attached to the message
	InputMessageContent   any                   `json:"input_message_content,omitempty"`    // Optional. Content of the message to be sent instead of the photo
}

// InlineQueryResultCachedGIF Represents a link to an animated GIF file stored on the Telegram servers.
// By default, this animated GIF file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with specified content instead of the animation.
type InlineQueryResultCachedGIF struct {
	Type                  string                `json:"type"`                               // the type of the result must be gif
	ID                    string                `json:"id"`                                 // Unique identifier for this result, 1-64 bytes
	GifID                 string                `json:"gif_file_id"`                        // A valid file identifier for the GIF file
	Title                 string                `json:"title,omitempty"`                    // Optional. Title for the result
	Caption               string                `json:"caption,omitempty"`                  // Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
	ParseMode             string                `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the caption. See [formatting options](https://core.telegram.org/bots/api#formatting-options) for more details.
	CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // Optional. Inline keyboard attached to the message
	InputMessageContent   any                   `json:"input_message_content,omitempty"`    // Optional. Content of the message to be sent instead of the GIF animation
}

// InlineQueryResultCachedMPEG4GIF Represents a link to a video animation
// (H.264/MPEG-4 AVC video without a sound) stored on the Telegram servers.
// By default, this animated MPEG-4 file will be sent by the user with an optional caption.
// Alternatively,
// you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultCachedMPEG4GIF struct {
	Type                  string                `json:"type"`                               // type of the result must be mpeg4_gif
	ID                    string                `json:"id"`                                 // Unique identifier for this result, 1-64 bytes
	MPEG4FileID           string                `json:"mpeg4_file_id"`                      // A valid file identifier for the MPEG4 file
	Title                 string                `json:"title,omitempty"`                    // Optional. Title for the result
	Caption               string                `json:"caption,omitempty"`                  // Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
	ParseMode             string                `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the caption. See [formatting options](https://core.telegram.org/bots/api#formatting-options) for more details.
	CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // Optional. Inline keyboard attached to the message
	InputMessageContent   any                   `json:"input_message_content,omitempty"`    // Optional. Content of the message to be sent instead of the video animation
}

// InlineQueryResultCachedSticker Represents a link to a sticker stored on the Telegram servers.
// By default, this sticker will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the sticker.
// Note:
// This will only work in Telegram versions
// released after 9 April 2016 for static stickers and after 06 July 2019 for animated stickers.
// Older clients will ignore them.
type InlineQueryResultCachedSticker struct {
	Type                string                `json:"type"`                            // the type of the result must be stickered
	ID                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	StickerID           string                `json:"sticker_file_id"`                 // A valid file identifier of the sticker
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent any                   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the sticker
}

// InlineQueryResultCachedDocument Represents a link to a file stored on the Telegram servers.
// By default, this file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the file.
// Note: This will only work in Telegram versions released after 9 April 2016.
// Older clients will ignore them.
type InlineQueryResultCachedDocument struct {
	Type                string                `json:"type"`                            // the type of the result must be documented
	ID                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	Title               string                `json:"title"`                           // title for the result
	DocumentID          string                `json:"document_file_id"`                // A valid file identifier for the file
	Description         string                `json:"description,omitempty"`           // Optional. Short description of the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the document caption. See [formatting options](https://core.telegram.org/bots/api#formatting-options) for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent any                   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the file
}

// InlineQueryResultCachedVideo Represents a link to a video file stored on the Telegram servers.
// By default, this video file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
type InlineQueryResultCachedVideo struct {
	Type                  string                `json:"type"`                               // type of the result must be video
	ID                    string                `json:"id"`                                 // Unique identifier for this result, 1-64 bytes
	VideoID               string                `json:"video_file_id"`                      // A valid file identifier for the video file
	Title                 string                `json:"title"`                              // title for the result
	Description           string                `json:"description,omitempty"`              // Optional. Short description of the result
	Caption               string                `json:"caption,omitempty"`                  // Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	ParseMode             string                `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the video caption. See [formatting options](https://core.telegram.org/bots/api#formatting-options) for more details.
	CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // Optional. Inline keyboard attached to the message
	InputMessageContent   any                   `json:"input_message_content,omitempty"`    // Optional. Content of the message to be sent instead of the video
}

// InlineQueryResultCachedVoice Represents a link to a voice message stored on the Telegram servers.
// By default, this voice message will be sent by the user.
// Alternatively,
// you can use input_message_content to send a message with the specified content instead of the voice message.
// Note: This will only work in Telegram versions released after 9 April 2016.
// Older clients will ignore them.
type InlineQueryResultCachedVoice struct {
	Type                string                `json:"type"`                            // the type of the result must be voice
	ID                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	VoiceID             string                `json:"voice_file_id"`                   // A valid file identifier for the voice message
	Title               string                `json:"title"`                           // Voice message title
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption, 0-1024 characters after entities parsing
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the voice message caption. See [formatting options](https://core.telegram.org/bots/api#formatting-options) for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent any                   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the voice message
}

// InlineQueryResultCachedAudio Represents a link to an MP3 audio file stored on the Telegram servers.
// By default, this audio file will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
// Note: This will only work in Telegram versions released after 9 April 2016.
// Older clients will ignore them.
type InlineQueryResultCachedAudio struct {
	Type                string                `json:"type"`                            // The Type of the result must be audio
	ID                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	AudioID             string                `json:"audio_file_id"`                   // A valid file identifier for the audio file
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption, 0-1024 characters after entities parsing
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional. Mode for parsing entities in the audio caption. See [formatting options](https://core.telegram.org/bots/api#formatting-options) for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent any                   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the audio
}

// InputMessageContent Represents the content of a message to be sent as a result of an inline query.
// Telegram clients currently support the following five types:
type InputMessageContent struct {
	InputTextMessageContent     InputTextMessageContent
	InputLocationMessageContent InputLocationMessageContent
	InputVenueMessageContent    InputVenueMessageContent
	InputContactMessageContent  InputContactMessageContent
	InputInvoiceMessageContent  InputInvoiceMessageContent
}

type InputTextMessageContent struct {
	Text               string              `json:"message_text"`                   // text of the message to be sent, 1-4096 characters
	ParseMode          string              `json:"parse_mode,omitempty"`           // Optional. Mode for parsing entities in the message text. See formatting options for more details.
	Entities           []MessageEntity     `json:"entities,omitempty"`             // Optional. List of special entities that appear in message text, which can be specified instead of parse_mode
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"` // Optional. Link preview generation options for the message
}

// InputLocationMessageContent Represents the content of a location message to be sent as the result of an inline query.
type InputLocationMessageContent struct {
	Latitude             float64 `json:"latitude"`                         // latitude of the location in degrees
	Longitude            float64 `json:"longitude"`                        // longitude of the location in degrees
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`    // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod           int     `json:"live_period,omitempty"`            // Optional. The Period in seconds for which the location can be updated should be between 60 and 86400.
	Heading              int     `json:"heading,omitempty"`                // Optional. For live locations, the direction in which the user is moving in degrees. It Must be between 1 and 360 if specified.
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"` // Optional. For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. It Must be between 1 and 100000 if specified.
}

// InputVenueMessageContent Represents the content of a venue message to be sent as the result of an inline query.
type InputVenueMessageContent struct {
	Latitude        float64 `json:"latitude"`                    // Latitude of the venue in degrees
	Longitude       float64 `json:"longitude"`                   // Longitude of the venue in degrees
	Title           string  `json:"title"`                       // Name of the venue
	Address         string  `json:"address"`                     // Address of the venue
	FoursquareID    string  `json:"foursquare_id,omitempty"`     // Optional. Foursquare identifier of the venue, if known
	FoursquareType  string  `json:"foursquare_type,omitempty"`   // Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default,” “arts_entertainment/aquarium” or “food/ice cream.”)
	GooglePlaceID   string  `json:"google_place_id,omitempty"`   // Optional. Google Places identifier of the venue
	GooglePlaceType string  `json:"google_place_type,omitempty"` // Optional. Google Places a type of the venue. See [supported types](https://developers.google.com/places/web-service/supported_types).
}

// InputContactMessageContent Represents the content of a contact message to be sent as the result of an inline query.
type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"`        // Contact's phone number
	FirstName   string `json:"first_name"`          // Contact's first name
	LastName    string `json:"last_name,omitempty"` // Optional. Contact's last name
	VCard       string `json:"vcard,omitempty"`     // Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
}

// InputInvoiceMessageContent Represents the content of an invoice message to be sent as the result of an inline query.
type InputInvoiceMessageContent struct {
	Title                     string         `json:"title"`                                   // Product name, 1-32 characters
	Description               string         `json:"description"`                             // Product description, 1-255 characters
	Payload                   string         `json:"payload"`                                 // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use it for your internal processes.
	ProviderToken             string         `json:"provider_token"`                          // Payment provider token, obtained via @BotFather
	Currency                  string         `json:"currency"`                                // Three-letter ISO 4217 currency code
	Prices                    []LabeledPrice `json:"prices"`                                  // Price breakdown, a JSON-serialized list of components (e.g., product price, tax, discount, delivery cost, delivery tax, bonus, etc.)
	MaxTipAmount              int            `json:"max_tip_amount,omitempty"`                // Optional. The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Default to 0
	SuggestedTipAmounts       []int          `json:"suggested_tip_amounts,omitempty"`         // Optional. A JSON-serialized array of suggested amounts of tip in the smallest units of the currency (integer, not float/double). At most four suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
	ProviderData              string         `json:"provider_data,omitempty"`                 // Optional. A JSON-serialized object for data about the invoice, which will be shared with the payment provider. The payment provider should provide a detailed description of the required fields.
	PhotoURL                  string         `json:"photo_url,omitempty"`                     // Optional. URL of the product photo for the invoice. It Can be a photo of the goods or a marketing image for a service.
	PhotoSize                 int            `json:"photo_size,omitempty"`                    // Optional. Photo size in bytes
	PhotoWidth                int            `json:"photo_width,omitempty"`                   // Optional. Photo width
	PhotoHeight               int            `json:"photo_height,omitempty"`                  // Optional. Photo height
	NeedName                  bool           `json:"need_name,omitempty"`                     // Optional. Pass True, if you require the user's full name to complete the order
	NeedPhoneNumber           bool           `json:"need_phone_number,omitempty"`             // Optional. Pass True, if you require the user's phone number to complete the order
	NeedEmail                 bool           `json:"need_email,omitempty"`                    // Optional. Pass True, if you require the user's email address to complete the order
	NeedShippingAddress       bool           `json:"need_shipping_address,omitempty"`         // Optional. Pass True, if you require the user's shipping address to complete the order
	SendPhoneNumberToProvider bool           `json:"send_phone_number_to_provider,omitempty"` // Optional. Pass True, if the user's phone number should be sent to the provider
	SendEmailToProvider       bool           `json:"send_email_to_provider,omitempty"`        // Optional. Pass True, if the user's email address should be sent to the provider
	IsFlexible                bool           `json:"is_flexible,omitempty"`                   // Optional. Pass True, if the final price depends on the shipping method
}

// ChosenInlineResult Represents a result of an inline query that was chosen by the user and sent to their chat partner.
// Note: It is necessary to enable inline feedback via @BotFather in order to receive these objects in updates.
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`                   // The unique identifier for the result that was chosen
	From            User      `json:"from"`                        // The user that chose the result
	Location        *Location `json:"location,omitempty"`          // Optional. Sender location, only for bots that require user location
	InlineMessageID string    `json:"inline_message_id,omitempty"` // Optional. Identifier of the sent inline message. Available only if there is an inline keyboard attached to the message. It Will be also received in callback queries and can be used to edit the message.
	Query           string    `json:"query"`                       // The query that was used to obtain the result
}

// SentWebAppMessage Describes an inline message sent by a Web App on behalf of a user.
type SentWebAppMessage struct {
	InlineMessageID string `json:"inline_message_id"` // Optional. Identifier of the sent inline message. Available only if there is an inline keyboard attached to the message.
}

// LabeledPrice Represents a portion of the price for goods or services.
type LabeledPrice struct {
	Label  string `json:"label"`  // Portion label
	Amount int    `json:"amount"` // Price of the product in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
}

// Invoice Contains basic information about an invoice
type Invoice struct {
	Title          string `json:"title"`           // Product name
	Description    string `json:"description"`     // Product description
	StartParameter string `json:"start_parameter"` // Unique bot deep-linking parameter that can be used to generate this invoice
	Currency       string `json:"currency"`        // Three-letter ISO 4217 [currency](https://core.telegram.org/bots/payments#supported-currencies) code
	TotalAmount    int    `json:"total_amount"`    // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in [currencies.json](https://core.telegram.org/bots/payments/currencies.json), it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
}

// ShippingAddress Represents a shipping address.
type ShippingAddress struct {
	CountryCode string `json:"country_code"` // Two-letter ISO 3166-1 alpha-2 country code
	State       string `json:"state"`        // State, if applicable
	City        string `json:"city"`         // City
	StreetLine1 string `json:"street_line1"` // First line for the address
	StreetLine2 string `json:"street_line2"` // Second line for the address
	PostCode    string `json:"post_code"`    // Address post code
}

// OrderInfo Represents information about an order.
type OrderInfo struct {
	Name            string           `json:"name,omitempty"`             // Optional. Username
	PhoneNumber     string           `json:"phone_number,omitempty"`     // Optional. User's phone number
	Email           string           `json:"email,omitempty"`            // Optional. User email
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"` // Optional. User shipping address
}

// ShippingOption Represents one shipping option.
type ShippingOption struct {
	ID     string         `json:"id"`     // Shipping option identifier
	Title  string         `json:"title"`  // Option title
	Prices []LabeledPrice `json:"prices"` // List of price portions
}

// SuccessfulPayment Contains basic information about a successful payment.
type SuccessfulPayment struct {
	Currency                string     `json:"currency"`                     // Three-letter ISO 4217 currency code
	TotalAmount             int        `json:"total_amount"`                 // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	InvoicePayload          string     `json:"invoice_payload"`              // Bot specified invoice payload
	ShippingOptionID        string     `json:"shipping_option_id,omitempty"` // Optional. Identifier of the shipping option chosen by the user
	OrderInfo               *OrderInfo `json:"order_info,omitempty"`         // Optional. Order information provided by the user
	TelegramPaymentChargeID string     `json:"telegram_payment_charge_id"`   // Telegram payment identifier
	ProviderPaymentChargeID string     `json:"provider_payment_charge_id"`   // Provider payment identifier
}

// ShippingQuery Contains information about an incoming shipping query.
type ShippingQuery struct {
	ID              string          `json:"id"`               // Unique query identifier
	From            User            `json:"from"`             // User who sent the query
	InvoicePayload  string          `json:"invoice_payload"`  // Bot specified invoice payload
	ShippingAddress ShippingAddress `json:"shipping_address"` // User specified shipping address
}

// PreCheckoutQuery Contains a list of Telegram Star transactions.
type PreCheckoutQuery struct {
	Transactions []StarTransaction `json:"transactions"` // The list of transactions
}

// RevenueWithdrawalState Describes the state of a revenue withdrawal operation. Currently, it can be one of
type RevenueWithdrawalState struct {
	RevenueWithdrawalStatePending
	RevenueWithdrawalStateSucceeded
	RevenueWithdrawalStateFailed
}

// RevenueWithdrawalStatePending The withdrawal is in progress.
type RevenueWithdrawalStatePending struct {
	Type string `json:"type"` // Type of the state, always “pending”
}

// RevenueWithdrawalStateSucceeded The withdrawal succeeded.
type RevenueWithdrawalStateSucceeded struct {
	Type string `json:"type"` // 	Type of the state, always “succeeded”
	Date int64  `json:"date"` // 	Date the withdrawal was completed in Unix time
	Url  string `json:"url"`  // 	An HTTPS URL that can be used to see transaction details
}

// RevenueWithdrawalStateFailed The withdrawal failed and the transaction was refunded.
type RevenueWithdrawalStateFailed struct {
	Type string `json:"type"` // Type of the state, always “failed”
}

// TransactionPartner Describes the source of a transaction, or its recipient for outgoing transactions. Currently, it can be one of.
type TransactionPartner struct {
	TransactionPartnerFragment
	TransactionPartnerUser
	TransactionPartnerOther
}

// TransactionPartnerFragment Describes a withdrawal transaction with Fragment.
type TransactionPartnerFragment struct {
	Type            string                 `json:"type"`                       // Type of the transaction partner, always “fragment”
	WithdrawalState RevenueWithdrawalState `json:"withdrawal_state,omitempty"` // Optional. State of the transaction if the transaction is outgoing
}

// TransactionPartnerUser Describes a transaction with a user.
type TransactionPartnerUser struct {
	Type string `json:"type"` // Type of the transaction partner, always “user”
	User User   `json:"user"` // Information about the user
}

// TransactionPartnerOther Describes a transaction with an unknown source or recipient.
type TransactionPartnerOther struct {
	Type string `json:"type"` // 	Type of the transaction partner, always “other”
}

// StarTransaction Describes a Telegram Star transaction.
type StarTransaction struct {
	ID       string             `json:"id"`       // Unique identifier of the transaction. Coincides with the identifer of the original transaction for refund transactions. Coincides with SuccessfulPayment.telegram_payment_charge_id for successful incoming payments from users.
	Amount   int                `json:"amount"`   // Number of Telegram Stars transferred by the transaction
	Date     int                `json:"date"`     // Date the transaction was created in Unix time
	Source   TransactionPartner `json:"source"`   // Optional. Source of an incoming transaction (e.g., a user purchasing goods or services, Fragment refunding a failed withdrawal). Only for incoming transactions
	Receiver TransactionPartner `json:"receiver"` // Optional. Receiver of an outgoing transaction (e.g., a user for a purchase refund, Fragment for a withdrawal). Only for outgoing transactions
}

// StarTransactions Describes Telegram Passport data shared with the bot by the user.
type StarTransactions struct {
	Data        []EncryptedPassportElement `json:"data"`        // Array with information about documents and other Telegram Passport elements that was shared with the bot
	Credentials EncryptedCredentials       `json:"credentials"` // Encrypted credentials required to decrypt the data
}

// PassportData Describes Telegram Passport data shared with the bot by the user.
type PassportData struct {
	Data        []EncryptedPassportElement `json:"data"`        // Array with information about documents and other Telegram Passport elements that was shared with the bot
	Credentials EncryptedCredentials       `json:"credentials"` // Encrypted credentials required to decrypt the data
}

// PassportFile Represents a file uploaded to Telegram Passport.
// Currently, all Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.
type PassportFile struct {
	FileID       string `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileSize     int    `json:"file_size"`      // File size in bytes
	FileDate     int64  `json:"file_date"`      // Unix time when the file was uploaded
}

// EncryptedPassportElement Describes documents or other Telegram Passport elements shared with the bot by the user.
type EncryptedPassportElement struct {
	Type        string          `json:"type"`                   // Element type. One of “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport”, “address”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”, “phone_number”, “email”.
	Data        string          `json:"data,omitempty"`         // Optional. Base64-encoded encrypted Telegram Passport element data provided by the user, available for “personal_details,” “passport,” “driver_license,” “identity_card,” “internal_passport” and “address” types. Can be decrypted and verified using the accompanying [EncryptedCredentials](https://core.telegram.org/bots/api#encryptedcredentials).
	PhoneNumber string          `json:"phone_number,omitempty"` // Optional. User's verified phone number, available only for “phone_number” type
	Email       string          `json:"email,omitempty"`        // Optional. User's verified email address, available only for “email” type
	Files       []*PassportFile `json:"files,omitempty"`        // Optional. Array of encrypted files with documents provided by the user, available for “utility_bill,” “bank_statement,” “rental_agreement,” “passport_registration” and “temporary_registration” types. Files can be decrypted and verified using the accompanying [EncryptedCredentials](https://core.telegram.org/bots/api#encryptedcredentials).
	FrontSide   *PassportFile   `json:"front_side,omitempty"`   // Optional. Encrypted file with the front side of the document, provided by the user. Available for “passport,” “driver_license,” “identity_card” and “internal_passport.” The file can be decrypted and verified using the accompanying [EncryptedCredentials](https://core.telegram.org/bots/api#encryptedcredentials).
	ReverseSide *PassportFile   `json:"reverse_side,omitempty"` // Optional. Encrypted file with the reverse side of the document, provided by the user. Available for “driver_license” and “identity_card.” The file can be decrypted and verified using the accompanying [EncryptedCredentials](https://core.telegram.org/bots/api#encryptedcredentials).
	Selfie      *PassportFile   `json:"selfie,omitempty"`       // Optional. Encrypted file with the selfie of the user holding a document, provided by the user; available for “passport,” “driver_license,” “identity_card” and “internal_passport.” The file can be decrypted and verified using the accompanying [EncryptedCredentials](https://core.telegram.org/bots/api#encryptedcredentials).
	Translation []*PassportFile `json:"translation,omitempty"`  // Optional. Array of encrypted files with translated versions of documents provided by the user. Available if requested for “passport,” “driver_license,” “identity_card,” “internal_passport,” “utility_bill,” “bank_statement,” “rental_agreement,” “passport_registration” and “temporary_registration” types. Files can be decrypted and verified using the accompanying [EncryptedCredentials](https://core.telegram.org/bots/api#encryptedcredentials).
	Hash        string          `json:"hash"`                   // Base64-encoded element hash for using in [PassportElementErrorUnspecified](https://core.telegram.org/bots/api#passportelementerrorunspecified)
}

// EncryptedCredentials Describes data required for decrypting and authenticating EncryptedPassportElement.
// See the Telegram Passport Documentation for a complete description of the data decryption and authentication processes.
type EncryptedCredentials struct {
	Data   string `json:"data"`   // Base64-encoded encrypted JSON-serialized data with unique user's payload, data hashes and secrets required for [EncryptedPassportElement](https://core.telegram.org/bots/api#encryptedpassportelement) decryption and authentication
	Hash   string `json:"hash"`   // Base64-encoded data hash for data authentication
	Secret string `json:"secret"` // Base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption
}

// PassportElementError Represents an error in the Telegram Passport element which was submitted that should be resolved by the user.
// It should be one of:
type PassportElementError struct {
	PassportElementErrorDataField        PassportElementErrorDataField
	PassportElementErrorFrontSide        PassportElementErrorFrontSide
	PassportElementErrorReverseSide      PassportElementErrorReverseSide
	PassportElementErrorSelfie           PassportElementErrorSelfie
	PassportElementErrorFile             PassportElementErrorFile
	PassportElementErrorFiles            PassportElementErrorFiles
	PassportElementErrorTranslationFile  PassportElementErrorTranslationFile
	PassportElementErrorTranslationFiles PassportElementErrorTranslationFiles
	PassportElementErrorUnspecified      PassportElementErrorUnspecified
}

// PassportElementErrorDataField Represents an issue in one of the data fields that was provided by the user.
// The error is considered resolved when the field's value changes.
type PassportElementErrorDataField struct {
	Source    string `json:"source"`     // Error source, must be data
	Type      string `json:"type"`       // The section of the user's Telegram Passport, which has the error, one of “personal_details,” “passport,” “driver_license,” “identity_card,” “internal_passport,” “address”
	FieldName string `json:"field_name"` // Name of the data field which has the error
	DataHash  string `json:"data_hash"`  // Base64-encoded data hash
	Message   string `json:"message"`    // Error message
}

// PassportElementErrorFrontSide Represents an issue with the front side of a document.
// The error is considered resolved when the file with the front side of the document changes.
type PassportElementErrorFrontSide struct {
	Source   string `json:"source"`    // Error source must be front_side
	Type     string `json:"type"`      // The section of the user's Telegram Passport which has the issue, one of “passports,” “driver_license,” “identity_card,” “internal_passport”
	FileHash string `json:"file_hash"` // Base64-encoded hash of the file with the front side of the document
	Message  string `json:"message"`   // Error message
}

// PassportElementErrorReverseSide Represents an issue with the reverse side of a document.
// The error is considered resolved when the file with the reverse side of the document changes.
type PassportElementErrorReverseSide struct {
	Source   string `json:"source"`    // Error source, must be reverse_side
	Type     string `json:"type"`      // The section of the user's Telegram Passport which has the issue, one of “driver_license,” “identity_card”
	FileHash string `json:"file_hash"` // Base64-encoded hash of the file with the reverse side of the document
	Message  string `json:"message"`   // Error message
}

// PassportElementErrorSelfie Represents an issue with the selfie with a document.
// The error is considered resolved when the file with the selfie changes.
type PassportElementErrorSelfie struct {
	Source   string `json:"source"`    // Error source, must be selfie
	Type     string `json:"type"`      // The section of the user's Telegram Passport which has the issue, one of “passports,” “driver_license,” “identity_card,” “internal_passport”
	FileHash string `json:"file_hash"` // Base64-encoded hash of the file with the selfie
	Message  string `json:"message"`   // Error message
}

// PassportElementErrorFile Represents an issue with a document scan.
// The error is considered resolved when the file with the document scan changes.
type PassportElementErrorFile struct {
	Source   string `json:"source"`    // Error source must be filed
	Type     string `json:"type"`      // The section of the user's Telegram Passport which has the issue, one of “utility_bill,” “bank_statement,” “rental_agreement,” “passport_registration,” “temporary_registration”
	FileHash string `json:"file_hash"` // Base64-encoded file hash
	Message  string `json:"message"`   // Error message
}

// PassportElementErrorFiles Represents an issue with a list of scans.
// The error is considered resolved when the list of files containing the scan changes.
type PassportElementErrorFiles struct {
	Source     string   `json:"source"`      // Error source, must be files
	Type       string   `json:"type"`        // The section of the user's Telegram Passport which has the issue, one of “utility_bill,” “bank_statement,” “rental_agreement,” “passport_registration,” “temporary_registration”
	FileHashes []string `json:"file_hashes"` // List of base64-encoded file hashes
	Message    string   `json:"message"`     // Error message
}

// PassportElementErrorTranslationFile Represents an issue with one of the files
// that constitute the translation of a document.
// The error is considered resolved when the file changes.
type PassportElementErrorTranslationFile struct {
	Source   string   `json:"source"`    // Error source, must be translation_file
	Type     string   `json:"type"`      // Type of element of the user's Telegram Passport which has the issue, one of “passports,” “driver_license,” “identity_card,” “internal_passport,” “utility_bill,” “bank_statement,” “rental_agreement,” “passport_registration,” “temporary_registration”
	FileHash []string `json:"file_hash"` // Base64-encoded file hash
	Message  string   `json:"message"`   // Error message
}

// PassportElementErrorTranslationFiles Represents an issue with the translated version of a document.
// The error is considered resolved when a file with the document translation changes.
type PassportElementErrorTranslationFiles struct {
	Source     string   `json:"source"`      // Error source, must be translation_files
	Type       string   `json:"type"`        // Type of element of the user's Telegram Passport which has the issue, one of “passports,” “driver_license,” “identity_card,” “internal_passport,” “utility_bill,” “bank_statement,” “rental_agreement,” “passport_registration,” “temporary_registration”
	FileHashes []string `json:"file_hashes"` // List of base64-encoded file hashes
	Message    string   `json:"message"`     // Error message
}

// PassportElementErrorUnspecified Represents an issue in an unspecified place.
// The error is considered resolved when new data is added.
type PassportElementErrorUnspecified struct {
	Source   string `json:"source"`    // Error source must be unspecified
	Type     string `json:"type"`      // Type of element of the user's Telegram Passport which has the issue
	FileHash string `json:"file_hash"` // Base64-encoded element hash
	Message  string `json:"message"`   // Error message
}

// Game Represents a game. Use BotFather to create and edit games, their short names will act as unique identifiers.
type Game struct {
	Title        string           `json:"title"`                   // Title of the game
	Description  string           `json:"description"`             // Description of the game
	Photo        []PhotoSize      `json:"photo"`                   // Photo that will be displayed in the game message in chats.
	Text         string           `json:"text,omitempty"`          // Optional. Brief description of the game or high scores included in the game message. It Can be automatically edited to include current high scores for the game when the bot calls setGameScore, or manually edited using editMessageText. 0-4096 characters.
	TextEntities []*MessageEntity `json:"text_entities,omitempty"` // Optional. Special entities that appear in text, such as usernames, URLs, bot commands, etc.
	Animation    *Animation       `json:"animation,omitempty"`     // Optional. Animation that will be displayed in the game message in chats. Upload via BotFather
}

// CallbackGame A placeholder currently holds no information. Use BotFather to set up your game.
type CallbackGame struct{}

// GameHighScore Represents one row of the high scores table for a game
type GameHighScore struct {
	Position int  `json:"position"` // Position in high-score table for the game
	User     User `json:"user"`     // User
	Score    int  `json:"score"`    // Score
}
