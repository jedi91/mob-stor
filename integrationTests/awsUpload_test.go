// Copyright (c) Michael Kovacevich 2019. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package integrationTests

import (
	"fmt"
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

	containerName := "mob-stor-integrationtests"

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
	id := os.Getenv("AWS_ID")
	fmt.Println("ID: " + id)

	secret := os.Getenv("AWS_SECRET")
	fmt.Println("Secret: " + secret)

	region := os.Getenv("AWS_REGION")
	fmt.Println("Region: " + region)

	return transmit.AWSS3Transmitter{
		Id:     id,
		Secret: secret,
		Token:  "",
		Region: region,
	}
}
