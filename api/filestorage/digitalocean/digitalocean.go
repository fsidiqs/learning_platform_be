package digitalocean

import "io"

type FileStorage struct{}

func (f *FileStorage) UploadFile(filename string, contents io.Reader) error {
	return nil
}
