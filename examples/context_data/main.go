package main

//go:generate go run ../../scripts/tools/get_tdjson.go

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

// --- Constants for Context Data Keys ---
const (
	KeyStartTime   = "start_time"
	KeyAdmins      = "admins"
	KeyUser        = "user"
	KeyRateLimited = "is_rate_limited"
)

// AdminCache  cache for chat administrators.
type AdminCache struct {
	sync.RWMutex
	// Map ChatID -> []UserID
	Admins map[int64][]int64
}

// User represents a mock user model from a database.
type User struct {
	ID        int64
	Username  string
	IsPremium bool
}

// RateLimitCache stores last request time per user.
type RateLimitCache struct {
	sync.Mutex
	LastSeen map[int64]time.Time
}

var adminCache = &AdminCache{
	Admins: make(map[int64][]int64),
}

var rateLimitCache = &RateLimitCache{
	LastSeen: make(map[int64]time.Time),
}

func main() {
	apiID := int32(6)
	apiHash := ""
	botToken := ""

	bot := gotdbot.NewClient(apiID, apiHash, botToken, &gotdbot.ClientConfig{LibraryPath: "./libtdjson.so.1.8.60"})
	gotdbot.SetTdlibLogVerbosityLevel(2)

	dispatcher := ext.NewDispatcher(bot)

	// --- Middleware (Group -1) ---
	// Rate Limiter: Runs BEFORE everything else.
	dispatcher.AddHandlerToGroup(handlers.NewMessage(filters.All, func(ctx *ext.Context) error {
		ctx.Data[KeyStartTime] = time.Now()
		msg := ctx.EffectiveMessage
		senderId := msg.FromID()
		if senderId == 0 {
			return nil
		}

		rateLimitCache.Lock()
		last, exists := rateLimitCache.LastSeen[senderId]
		rateLimitCache.LastSeen[senderId] = time.Now()
		rateLimitCache.Unlock()

		// Simple limit: 1 request per second
		if exists && time.Since(last) < 1*time.Second {
			ctx.Data[KeyRateLimited] = true
			log.Printf("[RateLimit] User %d is sending too fast", senderId)
		}

		mockUser := &User{
			ID:        senderId,
			Username:  "MockUser",
			IsPremium: senderId%2 == 0,
		}
		ctx.Data[KeyUser] = mockUser

		// Admin Caching
		chatId := ctx.EffectiveChatId
		if chatId != 0 {
			admins := getAdmins(bot, chatId)
			ctx.Data[KeyAdmins] = admins
		}

		return nil
	}), -1)

	// Helper to check rate limit at start of handlers
	checkRateLimit := func(ctx *ext.Context) bool {
		if val, ok := ctx.Data[KeyRateLimited].(bool); ok && val {
			ctx.Client.SendTextMessage(ctx.EffectiveChatId, "⏳ <b>Slow down!</b> You are being rate limited.", &gotdbot.SendTextMessageOpts{
				ReplyToMessageID: ctx.EffectiveMessage.Id,
				ParseMode:        "HTML",
			})
			return true
		}
		return false
	}

	// Command: /admin
	dispatcher.AddHandlerToGroup(handlers.NewCommand("admin", func(ctx *ext.Context) error {
		if checkRateLimit(ctx) {
			return nil
		}
		msg := ctx.EffectiveMessage
		c := ctx.Client

		senderId := msg.FromID()
		if checkAdmin(ctx, senderId) {
			msg.ReplyText(c, "👮 <b>Access Granted</b>: You are an admin.", &gotdbot.SendTextMessageOpts{ParseMode: "HTML"})
		} else {
			msg.ReplyText(c, "🚫 <b>Access Denied</b>: You are not an admin. (Try /promote_me)", &gotdbot.SendTextMessageOpts{ParseMode: "HTML"})
		}
		return nil
	}), 1)

	// Command: /me (Uses User Data)
	dispatcher.AddHandlerToGroup(handlers.NewCommand("me", func(ctx *ext.Context) error {
		if checkRateLimit(ctx) {
			return nil
		}
		m := ctx.EffectiveMessage
		c := ctx.Client

		val, ok := ctx.Data[KeyUser]
		if !ok {
			m.ReplyText(c, "❌ User data not found in context.", &gotdbot.SendTextMessageOpts{ParseMode: "HTML"})
			return nil
		}

		user := val.(*User)
		msg := fmt.Sprintf("👤 <b>User Profile</b>\n"+
			"ID: <code>%d</code>\n"+
			"Premium: %v", user.ID, user.IsPremium)

		m.ReplyText(c, msg, &gotdbot.SendTextMessageOpts{ParseMode: "HTML"})
		return nil
	}), 1)

	// Command: /promote_me (For testing Admin Cache)
	dispatcher.AddHandlerToGroup(handlers.NewCommand("promote_me", func(ctx *ext.Context) error {
		msg := ctx.EffectiveMessage
		senderId := msg.FromID()
		chatId := ctx.EffectiveChatId

		adminCache.Lock()
		admins := adminCache.Admins[chatId]
		exists := false
		for _, id := range admins {
			if id == senderId {
				exists = true
				break
			}
		}
		if !exists {
			adminCache.Admins[chatId] = append(admins, senderId)
		}
		adminCache.Unlock()

		ctx.Client.SendTextMessage(chatId, "🎉 You are now cached as an admin! Try /admin again.", nil)
		return nil
	}), 1)

	// Command: /info
	dispatcher.AddHandlerToGroup(handlers.NewCommand("info", func(ctx *ext.Context) error {
		startTime, _ := ctx.Data[KeyStartTime].(time.Time)
		admins, _ := ctx.Data[KeyAdmins].([]int64)

		msg := fmt.Sprintf("ℹ️ <b>Request Info</b>\n"+
			"• Processed in: %v\n"+
			"• Cached Admins: %d", time.Since(startTime), len(admins))

		ctx.Client.SendTextMessage(ctx.EffectiveChatId, msg, &gotdbot.SendTextMessageOpts{ParseMode: "HTML"})
		return nil
	}), 1)

	dispatcher.Start()
	log.Println("Starting bot...")
	if err := bot.Start(); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}

	bot.Idle()
}

func getAdmins(bot *gotdbot.Client, chatId int64) []int64 {
	adminCache.RLock()
	admins, found := adminCache.Admins[chatId]
	adminCache.RUnlock()

	if found {
		return admins
	}

	// Cache miss
	// Return empty list initially. Use /promote_me to test.
	var fetchedAdmins []int64

	//administrators, err := bot.GetChatAdministrators(chatId)

	adminCache.Lock()
	adminCache.Admins[chatId] = fetchedAdmins
	adminCache.Unlock()

	return fetchedAdmins
}

func checkAdmin(ctx *ext.Context, userId int64) bool {
	val, ok := ctx.Data[KeyAdmins]
	if !ok {
		return false
	}
	admins, ok := val.([]int64)
	if !ok {
		return false
	}

	for _, id := range admins {
		if id == userId {
			return true
		}
	}
	return false
}
