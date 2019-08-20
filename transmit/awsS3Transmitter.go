package transmit

import (
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AWSS3Transmitter struct {
}

func (t AWSS3Transmitter) GetName() string {
	return "AWSS3"
}

func (t AWSS3Transmitter) Transmit(
	data []byte,
	filepath string,
	containerName string,
) {

}
