package backblaze

import (
	"go_jwt_auth/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func NewBackblazeClient(conf config.StorageConf) (*s3manager.Uploader, error) {
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(conf.KEYID, conf.APPKEY, ""),
		Endpoint:    &conf.ENDPOINT,
		Region:      &conf.REGION,
	})

	if err != nil {
		return nil, err
	}
	s3Client := s3.New(sess)
	Uploader := s3manager.NewUploaderWithClient(s3Client)
	return Uploader, nil
}
