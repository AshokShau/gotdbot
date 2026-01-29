package gotdbot

// Download downloads the file.
func (t *File) Download(c *Client, priority int32, offset int64, limit int64, synchronous bool) (*File, error) {
	file, err := c.DownloadFile(t.Id, priority, offset, limit, synchronous)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// Delete deletes the file from the TDLib file cache.
func (t *File) Delete(c *Client) error {
	_, err := c.DeleteFile(t.Id)
	return err
}

// Download downloads the file.
func (t *RemoteFile) Download(c *Client, priority int32, offset int64, limit int64, synchronous bool) (*File, error) {
	_, err := c.GetRemoteFile(t.Id, nil)
	if err != nil {
		return nil, err
	}

	return t.Download(c, priority, offset, limit, synchronous)
}

// Delete deletes the file from the TDLib file cache.
func (t *RemoteFile) Delete(c *Client) error {
	fileInfo, err := c.GetRemoteFile(t.Id, nil)
	if err != nil {
		return err
	}
	_, err = c.DeleteFile(fileInfo.Id)
	return err
}
