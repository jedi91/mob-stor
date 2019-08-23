package integrationTests

import (
	"github.com/jedi91/mob-stor/distribute"
	"github.com/jedi91/mob-stor/transmit"
	"os"
	"testing"
	"time"
)

func TestAWSFileUpload(t *testing.T) {
	transmitter := setupAwsTransmitter()

	transmitters := []transmit.Transmitter{
		transmitter,
	}

	distributor := distribute.Distributor{
		Transmitters: transmitters,
	}

	data := []byte(
		"Testing uploading a file to aws s3 with mob-stor",
	)

	filepath := "TestAWSFileUpload_" +
		time.Now().Format(dateFormat)

	containerName := "integrationtests"

	results := distributor.Distribute(
		data,
		filepath,
		containerName,
	)

	success := checkForSuccess(results)

	checkExpectedBool(
		t,
		success,
		true,
	)
}

func setupAwsTransmitter() transmit.Transmitter {
	return transmit.AWSS3Transmitter{
		Id:     os.Getenv("AWS_ID"),
		Secret: os.Getenv("AWS_SECRET"),
		Token:  os.Getenv("AWS_TOKEN"),
		Region: os.Getenv("AWS_REGION"),
	}
}
