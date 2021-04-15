package filestorage

import (
	"io"
)

type IFileStorage interface {
	UploadImage(filename string, contents io.Reader) (string, error)
	UploadVideo(filename string, contents io.Reader) (string, error)
}

type fileStorageService struct {
	storage IFileStorage
}

func NewStorageService(s IFileStorage) fileStorageService {
	return fileStorageService{storage: s}
}

// func (f *fileStorageService) UploadImage(filename string, file io.Reader) (filestorage.UploadResponse, error) {
// 	_ = f.storage.UploadImage(filename, file)
// 	return filestorage.UploadResponse{}, nil

// }
