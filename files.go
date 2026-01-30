package gotdbot

// Download downloads the file.
func (r *RemoteFile) Download(c *Client, priority int32, offset int64, limit int64, synchronous bool) (*File, error) {
	_, err := c.GetRemoteFile(r.Id, nil)
	if err != nil {
		return nil, err
	}

	return r.Download(c, priority, offset, limit, synchronous)
}

// Delete deletes the file from the TDLib file cache.
func (r *RemoteFile) Delete(c *Client) error {
	fileInfo, err := c.GetRemoteFile(r.Id, nil)
	if err != nil {
		return err
	}
	_, err = c.DeleteFile(fileInfo.Id)
	return err
}
