// A minimal bot. Set TELEGRAM_API_ID, TELEGRAM_API_HASH,
// TELEGRAM_BOT_TOKEN, and TDLIB_LIBRARY before running it.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/filters"
	"github.com/AshokShau/gotdbot/filters/message"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	apiID, err := strconv.ParseInt(strings.TrimSpace(os.Getenv("TELEGRAM_API_ID")), 10, 32)
	if err != nil || apiID <= 0 {
		return fmt.Errorf("set TELEGRAM_API_ID to a positive Telegram API ID")
	}
	apiHash := strings.TrimSpace(os.Getenv("TELEGRAM_API_HASH"))
	botToken := strings.TrimSpace(os.Getenv("TELEGRAM_BOT_TOKEN"))
	libraryPath := strings.TrimSpace(os.Getenv("TDLIB_LIBRARY"))
	if apiHash == "" || botToken == "" || libraryPath == "" {
		return fmt.Errorf("set TELEGRAM_API_HASH, TELEGRAM_BOT_TOKEN, and TDLIB_LIBRARY")
	}
	tokenID, tokenBody, ok := strings.Cut(botToken, ":")
	if _, err := strconv.ParseUint(tokenID, 10, 64); !ok || err != nil || tokenBody == "" || strings.Trim(tokenBody, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-") != "" {
		return fmt.Errorf("TELEGRAM_BOT_TOKEN must be a bot token from BotFather")
	}
	dataDir := os.Getenv("GOTDBOT_DATA_DIR")
	if dataDir == "" {
		dataDir = "./data/bot"
	}
	cfg := gotdbot.DefaultClientConfig()
	cfg.LibraryPath = libraryPath
	cfg.DatabaseDirectory = filepath.Join(dataDir, "database")
	cfg.FilesDirectory = filepath.Join(dataDir, "files")
	bot, err := gotdbot.NewClient(int32(apiID), apiHash, botToken, cfg)
	if err != nil {
		return fmt.Errorf("create bot: %w", err)
	}
	defer bot.Close()
	bot.OnCommand("start", func(c *gotdbot.Client, msg *gotdbot.Message) error {
		_, err := msg.ReplyText(c, "Hello! Send me a private text message and I will echo it.", nil)
		return err
	})
	bot.OnMessage(func(c *gotdbot.Client, msg *gotdbot.Message) error {
		_, err := msg.ReplyText(c, msg.Text(), nil)
		return err
	}, filters.AllFilters(
		message.Private,
		message.Incoming,
		filters.NotFilter(message.Command),
		func(msg *gotdbot.Message) bool { return msg.Text() != "" },
	))
	if err := bot.Start(); err != nil {
		return fmt.Errorf("start bot: %w", err)
	}
	bot.Idle()
	return nil
}
