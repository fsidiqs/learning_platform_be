package backblaze

import (
	"fmt"
	"go_jwt_auth/config"
	"io"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type backblaze struct {
	basepath string
	Uploader *s3manager.Uploader
}

func NewBackblaze(basepath string, conf config.StorageConf) (*backblaze, error) {
	c, err := NewBackblazeClient(conf)
	if err != nil {
		return nil, err
	}
	backblaze := &backblaze{
		basepath: basepath,
		Uploader: c,
	}
	return backblaze, nil
}

func (b *backblaze) UploadImage(filename string, file io.Reader) (string, error) {

	var extension = filepath.Ext(filename)
	var name = filename[0 : len(filename)-len(extension)]
	result, err := b.Uploader.Upload(&s3manager.UploadInput{
		ACL:         aws.String("public-read"),
		Bucket:      aws.String(config.NewAppConfig().StorageConf.BUCKETNAME),
		Key:         aws.String(fmt.Sprintf("%v%v-%v%v", b.basepath, name, time.Now().Unix(), extension)),
		Body:        file,
		ContentType: aws.String("image/jpeg"),
	})

	if err != nil {
		fmt.Printf("error %#v \n", err)
		fmt.Println(err)

		return "", err
	}

	return aws.StringValue(&result.Location), nil
}

func (b *backblaze) UploadVideo(filename string, file io.Reader) (string, error) {

	var extension = filepath.Ext(filename)
	var name = filename[0 : len(filename)-len(extension)]
	result, err := b.Uploader.Upload(&s3manager.UploadInput{
		ACL:         aws.String("public-read"),
		Bucket:      aws.String(config.NewAppConfig().StorageConf.BUCKETNAME),
		Key:         aws.String(fmt.Sprintf("%v%v-%v%v", b.basepath, name, time.Now().Unix(), extension)),
		Body:        file,
		ContentType: aws.String("video/mp4"),
	})

	if err != nil {
		fmt.Printf("error %#v \n", err)
		fmt.Println(err)

		return "", err
	}

	return aws.StringValue(&result.Location), nil
}
