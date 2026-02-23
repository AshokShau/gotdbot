package main

//go:generate go run ../../scripts/tools/get_tdjson.go

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers"
)

var (
	// activeTasks maps FileId (int64) to *Progress
	activeTasks sync.Map

	// activeTasksByName maps basename string -> *Progress for uploads
	activeTasksByName sync.Map
)

func main() {
	apiID := int32(6)
	apiHash := "API_HASH"
	botToken := "BOT_TOKEN"

	// Initialize bot
	bot, err := gotdbot.NewClient(apiID, apiHash, botToken, &gotdbot.ClientConfig{LibraryPath: "./libtdjson.so.1.8.61"})
	if err != nil {
		panic(err)
	}

	gotdbot.SetTdlibLogVerbosityLevel(2)
	// gotdbot.SetTdlibLogStreamFile("tdlib.log", 10*1024*1024, false)

	dispatcher := ext.NewDispatcher(bot)

	dispatcher.AddHandler(handlers.NewCommand("saved", func(ctx *ext.Context) error {
		message := ctx.EffectiveMessage
		client := ctx.Client
		userId := message.SenderID()
		if message.ReplyToMessageID() != 0 {
			replyMsg, err := message.GetRepliedMessage(client)
			if err != nil {
				_, _ = message.ReplyText(client, fmt.Sprintf("Failed to get replied message: %v", err), nil)
				return err
			}
			userId = replyMsg.SenderID()
		}

		// Tg allow Sends 2-10 messages grouped together into an album
		const maxAudiosToSend int32 = 10

		profileAudios, err := client.GetUserProfileAudios(0, maxAudiosToSend, userId)
		if err != nil {
			_, _ = message.ReplyText(client, fmt.Sprintf("Failed to get saved audios: %v", err), nil)
			return err
		}

		if len(profileAudios.Audios) == 0 {
			_, _ = message.ReplyText(client, "No saved audios found.", &gotdbot.SendTextMessageOpts{ParseMode: "HTML"})
			return nil
		}

		inputMessages := make([]gotdbot.InputMessageContent, 0, len(profileAudios.Audios))
		for _, audioItem := range profileAudios.Audios {
			if audioItem.Audio == nil ||
				audioItem.Audio.Remote == nil ||
				audioItem.Audio.Remote.Id == "" {
				continue
			}

			inputMessages = append(inputMessages, &gotdbot.InputMessageAudio{
				Audio: gotdbot.InputFileRemote{Id: audioItem.Audio.Remote.Id},
			})
		}

		_, err = client.SendMessageAlbum(ctx.EffectiveChatId, inputMessages, nil)
		return err
	}))

	// Register /download command
	dispatcher.AddHandler(handlers.NewCommand("download", downloadCmd))

	// Register /upload command
	dispatcher.AddHandler(handlers.NewCommand("upload", uploadCmd))

	// Register /fileid command
	dispatcher.AddHandler(handlers.NewCommand("fileid", getFileID))

	//  Register /file command
	dispatcher.AddHandler(handlers.NewCommand("file", sendWithFileID))

	// Register UpdateFile handler for progress tracking
	dispatcher.AddHandler(handlers.NewUpdateFile(nil, progressHandler))

	log.Println("Starting bot...")
	dispatcher.Start()

	// Start bot
	if err := bot.Start(); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}

	bot.Idle()
}

func downloadCmd(ctx *ext.Context) error {
	msg := ctx.EffectiveMessage
	c := ctx.Client
	if msg.ReplyToMessageID() == 0 {
		_, err := msg.ReplyText(c, "Please reply to a message with a file to download.", &gotdbot.SendTextMessageOpts{ParseMode: "HTML"})
		return err
	}

	replyMsg, err := msg.GetRepliedMessage(c)
	if err != nil {
		return fmt.Errorf("failed to get reply message: %w", err)
	}

	fileId, size, name := getFileInfo(replyMsg.Content)
	if fileId == 0 {
		_, err = msg.ReplyText(c, "No supported media found in the replied message.", &gotdbot.SendTextMessageOpts{ParseMode: "HTML"})
		return err
	}

	sentMsg, err := msg.ReplyText(c, fmt.Sprintf("📥 Starting download for %s...", gotdbot.EscapeHTML(name)), &gotdbot.SendTextMessageOpts{ParseMode: "HTML"})
	if err != nil {
		return err
	}

	activeTasks.Store(int64(fileId), &Progress{
		ChatId:     sentMsg.ChatId,
		MessageId:  sentMsg.Id,
		StartTime:  time.Now(),
		LastUpdate: time.Now(),
		LastSize:   0,
		FileName:   name,
		IsUpload:   false,
		TotalSize:  size,
	})

	dFile, err := c.DownloadFile(fileId, 0, 0, 1, &gotdbot.DownloadFileOpts{Synchronous: true})
	if err != nil {
		activeTasks.Delete(int64(fileId))
		_, _ = sentMsg.EditText(c, fmt.Sprintf("Failed to start download: %v", err), nil)
		return err
	}

	if dFile.Local.IsDownloadingCompleted {
		activeTasks.Delete(int64(fileId))
		text := fmt.Sprintf("✅ Download already completed for <code>%s</code>", gotdbot.EscapeHTML(name))
		_, _ = sentMsg.EditText(c, text, nil)
	}

	return nil
}

func uploadCmd(ctx *ext.Context) error {
	args := getArgs(ctx)
	msg := ctx.EffectiveMessage
	c := ctx.Client

	if len(args) == 0 {
		_, err := msg.ReplyText(c, "Please provide a file path to upload. Usage: /upload <path>", &gotdbot.SendTextMessageOpts{ParseMode: "HTML"})
		return err
	}

	filePath := strings.Join(args, " ")
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		_, err := msg.ReplyText(c, "file not found", &gotdbot.SendTextMessageOpts{ParseMode: "HTML"})
		return err
	}

	fileName := filepath.Base(filePath)
	fileSize := info.Size()

	text := fmt.Sprintf("📤 Starting upload for %s...", gotdbot.EscapeHTML(fileName))
	progressMsg, err := msg.ReplyText(c, text, &gotdbot.SendTextMessageOpts{ParseMode: "HTML"})
	if err != nil {
		return err
	}

	base := filepath.Base(filePath)
	activeTasksByName.Store(base, &Progress{
		ChatId:     progressMsg.ChatId,
		MessageId:  progressMsg.Id,
		StartTime:  time.Now(),
		LastUpdate: time.Now(),
		LastSize:   0,
		FileName:   fileName,
		IsUpload:   true,
		TotalSize:  fileSize,
		LocalBase:  base,
	})

	_, err = c.SendDocument(ctx.EffectiveChatId, gotdbot.InputFileLocal{Path: filePath}, &gotdbot.SendDocumentOpts{
		Caption: "Uploaded",
	})

	if err != nil {
		activeTasksByName.Delete(base)
		_, _ = progressMsg.EditText(c, fmt.Sprintf("Failed to start upload: %v", err), nil)
		return err
	}
	return nil
}

func getFileID(ctx *ext.Context) error {
	msg := ctx.EffectiveMessage
	c := ctx.Client
	if msg.ReplyToMessageID() == 0 {
		_, err := msg.ReplyText(c, "Please reply to a message with a file to get its File ID.", &gotdbot.SendTextMessageOpts{ParseMode: "HTML"})
		return err
	}

	replyMsg, err := msg.GetRepliedMessage(c)
	if err != nil {
		_, _ = msg.ReplyText(c, fmt.Sprintf("Failed to get reply message: %v", err), &gotdbot.SendTextMessageOpts{ParseMode: "HTML"})
		return fmt.Errorf("failed to get reply message: %w", err)
	}

	fileID := replyMsg.RemoteFileID()
	if fileID == "" {
		_, err = msg.ReplyText(c, "No supported media found in the replied message.", &gotdbot.SendTextMessageOpts{ParseMode: "HTML"})
		return err
	}

	_, err = msg.ReplyText(c, fmt.Sprintf("📁 File ID: <code>%s</code>", gotdbot.EscapeHTML(fileID)), &gotdbot.SendTextMessageOpts{ParseMode: "HTML"})
	return err
}

func sendWithFileID(ctx *ext.Context) error {
	args := getArgs(ctx)
	msg := ctx.EffectiveMessage
	c := ctx.Client

	if len(args) == 0 {
		_, err := msg.ReplyText(c, "Please provide a File ID to send. Usage: /file <file_id>", &gotdbot.SendTextMessageOpts{ParseMode: "HTML"})
		return err
	}

	_, err := c.SendDocument(ctx.EffectiveChatId, gotdbot.InputFileRemote{Id: args[0]}, &gotdbot.SendDocumentOpts{
		Caption: "Here is your file",
	})

	if err != nil {
		_, _ = msg.ReplyText(c, fmt.Sprintf("Failed to send file: %v", err), &gotdbot.SendTextMessageOpts{ParseMode: "HTML"})
		return err
	}
	return nil
}

func progressHandler(ctx *ext.Context) error {
	update := ctx.Update.UpdateFile
	//jsonData, _ := json.MarshalIndent(update, "", "  ")
	//log.Println(string(jsonData))

	file := update.File
	c := ctx.Client

	val, ok := activeTasks.Load(int64(file.Id))
	var prog *Progress
	if ok {
		prog = val.(*Progress)
	} else {
		if file.Local != nil && file.Local.Path != "" {
			base := filepath.Base(file.Local.Path)
			if v, ok2 := activeTasksByName.Load(base); ok2 {
				prog = v.(*Progress)
				activeTasks.Store(int64(file.Id), prog)
			}
		}
	}

	if prog == nil {
		log.Printf("No active task for FileID: %d", file.Id)
		return nil
	}

	var currentSize int64
	var isCompleted bool
	var totalSize = prog.TotalSize

	if prog.IsUpload {
		if file.Remote.IsUploadingActive {
			currentSize = file.Remote.UploadedSize
		} else if file.Remote.IsUploadingCompleted {
			currentSize = file.Size
			isCompleted = true
		} else {
			currentSize = file.Remote.UploadedSize
		}
		if file.Size > 0 {
			totalSize = file.Size
		}
	} else {
		// Download
		currentSize = file.Local.DownloadedSize
		isCompleted = file.Local.IsDownloadingCompleted
		if file.Size > 0 {
			totalSize = file.Size
		} else if file.ExpectedSize > 0 {
			totalSize = file.ExpectedSize
		}
	}

	now := time.Now()
	if isCompleted {
		activeTasks.Delete(int64(file.Id))
		if prog.LocalBase != "" {
			activeTasksByName.Delete(prog.LocalBase)
		}
		duration := now.Sub(prog.StartTime).Seconds()
		avgSpeed := float64(totalSize) / math.Max(duration, 0.000001)

		text := fmt.Sprintf(
			"✅ <b>%s Complete:</b> <code>%s</code>\n"+
				"💾 <b>Size:</b> %s\n"+
				"⏱ <b>Time Taken:</b> %s\n"+
				"⚡ <b>Average Speed:</b> %s/s",
			getActionName(prog.IsUpload),
			gotdbot.EscapeHTML(prog.FileName),
			formatBytes(totalSize),
			formatTime(duration),
			formatBytes(int64(avgSpeed)),
		)

		_, err := c.EditTextMessage(prog.ChatId, prog.MessageId, text, &gotdbot.EditTextMessageOpts{ParseMode: "HTML"})
		return err
	}

	elapsed := now.Sub(prog.LastUpdate).Seconds()
	delta := currentSize - prog.LastSize
	speed := float64(delta) / elapsed
	if elapsed <= 0 {
		speed = 0
	}

	interval := calculateUpdateInterval(totalSize, speed)

	if elapsed < interval {
		return nil
	}

	prog.LastUpdate = now
	prog.LastSize = currentSize
	percentage := 0
	if totalSize > 0 {
		percentage = int(float64(currentSize) / float64(totalSize) * 100)
	}
	if percentage > 100 {
		percentage = 100
	}

	eta := int64(-1)
	if speed > 0 && totalSize > currentSize {
		eta = int64(float64(totalSize-currentSize) / speed)
	}

	action := "Downloading"
	if prog.IsUpload {
		action = "Uploading"
	}
	icon := "📥"
	if prog.IsUpload {
		icon = "📤"
	}

	text := fmt.Sprintf(
		"%s <b>%s:</b> <code>%s</code>\n"+
			"💾 <b>Size:</b> %s / %s\n"+
			"📊 <b>Progress:</b> %d%% %s\n"+
			"🚀 <b>Speed:</b> %s/s\n"+
			"⏳ <b>ETA:</b> %s",
		icon, action, gotdbot.EscapeHTML(prog.FileName),
		formatBytes(currentSize), formatBytes(totalSize),
		percentage, progressBar(percentage),
		formatBytes(int64(speed)),
		formatTime(float64(eta)),
	)

	_, err := c.EditTextMessage(prog.ChatId, prog.MessageId, text, &gotdbot.EditTextMessageOpts{ParseMode: "HTML"})
	if err != nil {
		log.Printf("Failed to edit message (ChatID: %d, MessageID: %d): %v", prog.ChatId, prog.MessageId, err)
	}
	return err
}
