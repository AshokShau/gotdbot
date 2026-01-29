package filters

import "github.com/AshokShau/gotdbot"

type (
	Message            func(msg *gotdbot.Message) bool
	CallbackQuery      func(cq *gotdbot.UpdateNewCallbackQuery) bool
	InlineQuery        func(iq *gotdbot.UpdateNewInlineQuery) bool
	ChosenInlineResult func(cir *gotdbot.UpdateNewChosenInlineResult) bool
	ShippingQuery      func(sq *gotdbot.UpdateNewShippingQuery) bool
	PreCheckoutQuery   func(pcq *gotdbot.UpdateNewPreCheckoutQuery) bool
	Poll               func(p *gotdbot.UpdatePoll) bool
	PollAnswer         func(pa *gotdbot.UpdatePollAnswer) bool
	ChatMember         func(cm *gotdbot.UpdateChatMember) bool
	ChatJoinRequest    func(cjr *gotdbot.UpdateNewChatJoinRequest) bool
	User               func(u *gotdbot.UpdateUser) bool
	UserStatus         func(us *gotdbot.UpdateUserStatus) bool
	AuthorizationState func(as *gotdbot.UpdateAuthorizationState) bool
)
