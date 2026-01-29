package client

import "github.com/AshokShau/gotdbot/types"

// File is a wrapper around types.File with bound methods.
type File struct {
	*types.File
	client *Client
}

// NewFile creates a new bound File.
func NewFile(client *Client, f *types.File) *File {
	return &File{
		File:   f,
		client: client,
	}
}

// Download downloads the file.
func (f *File) Download(priority int32, offset int64, limit int64, synchronous bool) (*File, error) {
	file, err := f.client.DownloadFile(f.Id, priority, offset, limit, synchronous)
	if err != nil {
		return nil, err
	}
	return NewFile(f.client, file), nil
}

// Delete deletes the file from the TDLib file cache.
func (f *File) Delete() error {
	_, err := f.client.DeleteFile(f.Id)
	return err
}

// RemoteFile is a wrapper around types.RemoteFile with bound methods.
type RemoteFile struct {
	*types.RemoteFile
	client *Client
}

// NewRemoteFile creates a new bound RemoteFile.
func NewRemoteFile(client *Client, f *types.RemoteFile) *RemoteFile {
	return &RemoteFile{
		RemoteFile: f,
		client:     client,
	}
}

// Download downloads the file.
func (f *RemoteFile) Download(priority int32, offset int64, limit int64, synchronous bool) (*File, error) {
	fileInfo, err := f.client.GetRemoteFile(f.Id, nil)
	if err != nil {
		return nil, err
	}

	return NewFile(f.client, fileInfo).Download(priority, offset, limit, synchronous)
}

// Delete deletes the file.
func (f *RemoteFile) Delete() error {
	fileInfo, err := f.client.GetRemoteFile(f.Id, nil)
	if err != nil {
		return err
	}
	_, err = f.client.DeleteFile(fileInfo.Id)
	return err
}
