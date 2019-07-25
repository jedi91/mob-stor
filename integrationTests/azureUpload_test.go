package integrationTests

import (
	"github.com/jedi91/mob-stor/azure"
	"github.com/jedi91/mob-stor/distribute"
	"github.com/jedi91/mob-stor/transmit"
	"os"
	"testing"
	"time"
)

const boolTemplate = "Actual: %t | Expected: %t"
const dateFormat = "01-02-2006 15:04:05"

func TestAzureFileUpload(t *testing.T) {
	transmitter := setupTransmitter()
	distributor := setupDistributor(
		transmitter,
	)

	data := []byte(
		"Testing uploading a file to azureblob with mob-stor.",
	)

	fileName := "TestAzureFileUpload_" +
		time.Now().Format(dateFormat)

	path := "integrationtests"
	results := distributor.Distribute(
		data,
		fileName,
		path,
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

func setupTransmitter() transmit.Transmitter {
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

func setupDistributor(
	transmitter transmit.Transmitter,
) distribute.Distributor {
	transmitters := []transmit.Transmitter{
		transmitter,
	}

	return distribute.Distributor{
		Transmitters: transmitters,
	}
}

func checkForSuccess(
	results []distribute.Result,
) bool {
	success := true
	for _, result := range results {
		success = success && result.Success
	}

	return success
}

func checkExpectedBool(
	t *testing.T,
	actual bool,
	expected bool,
) {
	if actual == expected {
		return
	}

	t.Errorf(
		boolTemplate,
		actual,
		expected,
	)
}
