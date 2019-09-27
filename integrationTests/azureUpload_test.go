// Copyright (c) Michael Kovacevich 2019. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package integrationTests

import (
	"github.com/jedi91/mob-stor/azure"
	"github.com/jedi91/mob-stor/transmit"
	"os"
	"testing"
	"time"
)

func TestAzureFileUpload(t *testing.T) {
	transmitter := setupAzureTransmitter()
	distributor := setupDistributor(
		transmitter,
	)

	data := []byte(
		"Testing uploading a file to azureblob with mob-stor.",
	)

	filePath := "TestAzureFileUpload_" +
		time.Now().Format(dateFormat)

	containerName := "integrationtests"
	results := distributor.Distribute(
		data,
		filePath,
		containerName,
	)

	success := checkForSuccess(
		results,
	)

	checkExpectedBool(
		t,
		success,
		true,
	)
}

func TestAzureFileUploadWithPath(t *testing.T) {
	transmitter := setupAzureTransmitter()
	distributor := setupDistributor(
		transmitter,
	)

	data := []byte(
		"Testing uploading a file to azureblob with mob-stor.",
	)

	filePath := "TestFolder/TestAzureFileUpload_" +
		time.Now().Format(dateFormat)

	containerName := "integrationtests"
	results := distributor.Distribute(
		data,
		filePath,
		containerName,
	)

	success := checkForSuccess(
		results,
	)

	checkExpectedBool(
		t,
		success,
		true,
	)
}

func setupAzureTransmitter() transmit.Transmitter {
	credsProvider := azure.BlobCredentialProvider{
		AccountName: os.Getenv("AZURE_STORAGE_ACCOUNT"),
		AccountKey:  os.Getenv("AZURE_STORAGE_ACCESS_KEY"),
	}

	urlProvider := azure.ContainerUrlProvider{
		CredsProvider: credsProvider,
	}

	return transmit.AzureBlobTransmitter{
		ContainerUrlProvider: urlProvider,
	}
}
