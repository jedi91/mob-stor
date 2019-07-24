package integrationTests

import (
	"github.com/jedi91/mob-stor/azure"
	"github.com/jedi91/mob-stor/distribute"
	"github.com/jedi91/mob-stor/transmit"
	"testing"
)

func TestAzureFileUpload(t *testing.T) {
	transmitter := setupTransmitter()
	distributor := setupDistributor(
		transmitter,
	)

	data := []byte(
		"Testing uploading a file to azureblob with mob-stor.",
	)

	fileName := "TestAzureFileUpload"
	path := "mob-stor/integration-tests"
	results := distributor.Distribute(
		data,
		fileName,
		path,
	)
}

func setupTransmitter() transmit.Transmitter {
	credsProvider := azure.BlobCredentialProvider{
		"testAccountName",
		"testAccountKey",
	}

	urlProvider := azure.ContainerUrlProvider{
		credsProvider,
	}

	return transmit.AzureBlobTransmitter{
		urlProvider,
	}
}

func setupDistributor(
	transmitter transmit.Transmitter,
) distribute.Distributor {
	transmitters := []transmit.Transmitter{
		transmitter,
	}

	return distribute.Distributor{
		transmitters,
	}
}
