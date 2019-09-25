package transmit

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AWSS3Transmitter struct {
	Id     string
	Secret string
	Token  string
	Region string
}

func (t AWSS3Transmitter) GetName() string {
	return "AWSS3" //TODO: look for a way to get the name of a type
}

func (t AWSS3Transmitter) Transmit(
	data []byte,
	filepath string,
	containerName string,
) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(t.Region),
		Credentials: credentials.NewStaticCredentials(
			t.Id,
			t.Secret,
			t.Token,
		),
	})

	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(containerName),
		Key:    aws.String(filepath),
		Body:   bytes.NewReader(data),
	})

	if err != nil {
		return err
	}

	return nil
}
