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

// Progress tracks the state of a download or upload
type Progress struct {
	ChatId     int64
	MessageId  int64
	StartTime  time.Time
	LastUpdate time.Time
	LastSize   int64
	FileName   string
	IsUpload   bool
	TotalSize  int64
	LocalBase  string
}

var (
	// activeTasks maps FileId (int64) to *Progress
	activeTasks sync.Map

	// activeTasksByName maps basename string -> *Progress for uploads
	activeTasksByName sync.Map
)

func main() {
	apiID := int32(0)
	apiHash := ""
	botToken := ""

	// Initialize bot
	bot := gotdbot.NewClient(apiID, apiHash, botToken, &gotdbot.ClientConfig{LibraryPath: "./libtdjson.so.1.8.60"})
	gotdbot.SetTdlibLogVerbosityLevel(2)

	dispatcher := ext.NewDispatcher(bot)

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

func getArgs(ctx *ext.Context) []string {
	if ctx.EffectiveMessage == nil {
		return nil
	}
	txt, ok := ctx.EffectiveMessage.Content.(*gotdbot.MessageText)
	if !ok || txt.Text == nil {
		return nil
	}
	parts := strings.Fields(txt.Text.Text)
	if len(parts) > 1 {
		return parts[1:]
	}
	return nil
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

	dFile, err := c.DownloadFile(fileId, 1, 0, 0, false)
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

	_, err = c.SendDocument(ctx.EffectiveChatId, filePath, &gotdbot.SendDocumentOpts{
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

	_, err := c.SendDocument(ctx.EffectiveChatId, args[0], &gotdbot.SendDocumentOpts{
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

func getActionName(isUpload bool) string {
	if isUpload {
		return "Upload"
	}
	return "Download"
}

func calculateUpdateInterval(fileSize int64, speed float64) float64 {
	base := 1.0
	const threshold = 5 * 1024 * 1024
	if fileSize >= threshold {
		ratio := float64(fileSize) / float64(threshold)
		scale := math.Min(math.Log10(ratio), 2)
		base = 1.0 + scale*2.0
	}

	speedMod := 1.0
	if speed > 1024*1024 {
		speedMod = math.Max(0.5, 2.0-(speed/float64(threshold)))
	}

	res := base * speedMod
	if res < 1.0 {
		res = 1.0
	}
	if res > 5.0 {
		res = 5.0
	}
	return res
}

func getFileInfo(content gotdbot.MessageContent) (int32, int64, string) {
	switch c := content.(type) {
	case *gotdbot.MessageDocument:
		return c.Document.Document.Id, c.Document.Document.Size, c.Document.FileName
	case *gotdbot.MessagePhoto:
		if len(c.Photo.Sizes) > 0 {
			last := c.Photo.Sizes[len(c.Photo.Sizes)-1]
			return last.Photo.Id, last.Photo.Size, "photo.jpg"
		}
	case *gotdbot.MessageVideo:
		return c.Video.Video.Id, c.Video.Video.Size, c.Video.FileName
	case *gotdbot.MessageAudio:
		return c.Audio.Audio.Id, c.Audio.Audio.Size, c.Audio.FileName
	case *gotdbot.MessageVoiceNote:
		return c.VoiceNote.Voice.Id, c.VoiceNote.Voice.Size, "voice.ogg"
	case *gotdbot.MessageVideoNote:
		return c.VideoNote.Video.Id, c.VideoNote.Video.Size, "video_note.mp4"
	case *gotdbot.MessageAnimation:
		return c.Animation.Animation.Id, c.Animation.Animation.Size, c.Animation.FileName
	case *gotdbot.MessageSticker:
		return c.Sticker.Sticker.Id, c.Sticker.Sticker.Size, "sticker.webp"
	}
	return 0, 0, ""
}

func formatBytes(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}

func formatTime(seconds float64) string {
	if seconds < 0 {
		return "Calculating..."
	}
	if seconds < 60 {
		return fmt.Sprintf("%ds", int(seconds))
	}
	minutes := int(seconds) / 60
	secondsRem := int(seconds) % 60
	if minutes < 60 {
		return fmt.Sprintf("%dm %ds", minutes, secondsRem)
	}
	hours := minutes / 60
	minutesRem := minutes % 60
	return fmt.Sprintf("%dh %dm", hours, minutesRem)
}

func progressBar(percentage int) string {
	const length = 10
	filled := int(math.Round(float64(length) * float64(percentage) / 100))
	if filled > length {
		filled = length
	}
	if filled < 0 {
		filled = 0
	}
	return strings.Repeat("⬢", filled) + strings.Repeat("⬡", length-filled)
}
