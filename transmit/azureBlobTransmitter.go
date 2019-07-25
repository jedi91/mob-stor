package transmit

import (
	"context"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/jedi91/mob-stor/azure"
)

type AzureBlobTransmitter struct {
	ContainerUrlProvider azure.ContainerUrlProvider
}

func (t AzureBlobTransmitter) GetName() string {
	return "AzureBlob"
}

func (t AzureBlobTransmitter) Transmit(
	data []byte,
	fileName string,
	path string,
) error {
	blobUrl, blobUrlErr := t.createBlobUrl(
		path,
		fileName,
	)

	if blobUrlErr != nil {
		return blobUrlErr
	}

	context := context.Background()
	options := azblob.UploadToBlockBlobOptions{
		BlockSize:   4 * 1024 * 1024, //TODO: Need to do this dynamically
		Parallelism: 16,              //TODO: Run tests with different inputs for this
	}

	_, uploadErr := azblob.UploadBufferToBlockBlob(
		context,
		data,
		blobUrl,
		options,
	)

	return uploadErr
}

func (t AzureBlobTransmitter) createBlobUrl(
	path string,
	fileName string,
) (
	azblob.BlockBlobURL,
	error,
) {
	containerUrl, err := t.
		ContainerUrlProvider.
		CreateContainerUrl(path)

	if err != nil {
		return azblob.BlockBlobURL{}, err
	}

	blobUrl := containerUrl.NewBlockBlobURL(
		fileName,
	)

	return blobUrl, nil
}
