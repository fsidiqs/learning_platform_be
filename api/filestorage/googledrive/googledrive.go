package googledrive

import (
	"fmt"
	"io"
	"log"

	"google.golang.org/api/drive/v3"
)

type GoogleDriveStorage struct {
	Basepath string
	client   *drive.Service
}

func NewGoogleDriveStorage() (*GoogleDriveStorage, error) {

	c, err := getService()
	if err != nil {
		return &GoogleDriveStorage{}, err
	}
	ikStorage := GoogleDriveStorage{client: c}
	return &ikStorage, nil
}

func (s *GoogleDriveStorage) UploadImage(name string, content io.Reader) (string, error) {
	f := &drive.File{
		MimeType: "image/jpg",
		Name:     name,
	}
	file, err := s.client.Files.Create(f).Media(content).Do()

	if err != nil {
		log.Println("Could not create file: " + err.Error())
		return "", err
	}
	fmt.Println("success")
	fmt.Printf("%#v\n", file)
	return "ok", nil
}
