package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/AshokShau/gotdbot"
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

func getArgs(ctx *gotdbot.Context) []string {
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
