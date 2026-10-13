package gotdbot

// GetFile returns the primary media file attached to the message, or nil.
// For a photo, it returns the file from the largest available photo size.
func (m *Message) GetFile() *File {
	if m == nil {
		return nil
	}
	switch content := m.Content.(type) {
	case *MessagePhoto:
		if content != nil && content.Photo != nil {
			var largest *File
			var largestArea int64
			for _, size := range content.Photo.Sizes {
				area := int64(size.Width) * int64(size.Height)
				if size.Photo != nil && (largest == nil || area >= largestArea) {
					largest, largestArea = size.Photo, area
				}
			}
			return largest
		}
	case *MessageVideo:
		if content != nil && content.Video != nil {
			return content.Video.Video
		}
	case *MessageSticker:
		if content != nil && content.Sticker != nil {
			return content.Sticker.Sticker
		}
	case *MessageAnimation:
		if content != nil && content.Animation != nil {
			return content.Animation.Animation
		}
	case *MessageAudio:
		if content != nil && content.Audio != nil {
			return content.Audio.Audio
		}
	case *MessageDocument:
		if content != nil && content.Document != nil {
			return content.Document.Document
		}
	case *MessageVoiceNote:
		if content != nil && content.VoiceNote != nil {
			return content.VoiceNote.Voice
		}
	case *MessageVideoNote:
		if content != nil && content.VideoNote != nil {
			return content.VideoNote.Video
		}
	}
	return nil
}

// RemoteFileID returns the remote file ID.
func (m *Message) RemoteFileID() string {
	if file := m.GetFile(); file != nil && file.Remote != nil {
		return file.Remote.Id
	}
	return ""
}

// RemoteDuration returns the remote media duration.
func (m *Message) RemoteDuration() int32 {
	if m == nil {
		return 0
	}
	switch content := m.Content.(type) {
	case *MessageVideo:
		if content != nil && content.Video != nil {
			return content.Video.Duration
		}
	case *MessageAnimation:
		if content != nil && content.Animation != nil {
			return content.Animation.Duration
		}
	case *MessageAudio:
		if content != nil && content.Audio != nil {
			return content.Audio.Duration
		}
	case *MessageVoiceNote:
		if content != nil && content.VoiceNote != nil {
			return content.VoiceNote.Duration
		}
	case *MessageVideoNote:
		if content != nil && content.VideoNote != nil {
			return content.VideoNote.Duration
		}
	}
	return 0
}

// RemoteFileSize returns the size, or the expected size if not known yet.
func (m *Message) RemoteFileSize() int64 {
	if file := m.GetFile(); file != nil {
		if file.Size > 0 {
			return file.Size
		}
		return file.ExpectedSize
	}
	return 0
}

// RemoteUniqueFileID returns the remote unique file ID.
func (m *Message) RemoteUniqueFileID() string {
	if file := m.GetFile(); file != nil && file.Remote != nil {
		return file.Remote.UniqueId
	}
	return ""
}

// Download downloads the media file attached to the message.
// It uses the attached File.Id and returns ErrNoMedia if no file is attached.
// Use the client that received the message, since file IDs belong to that client.
func (m *Message) Download(c *Client, priority int32, opts *DownloadFileOpts) (*File, error) {
	fileInfo := m.GetFile()
	if fileInfo == nil {
		return nil, ErrNoMedia
	}
	return c.DownloadFile(fileInfo.Id, priority, opts)
}
