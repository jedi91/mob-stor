// Copyright (c) Michael Kovacevich 2019. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

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
	filePath string,
	containerName string,
) error {
	blobUrl, blobUrlErr := t.createBlobUrl(
		containerName,
		filePath,
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
	containerName string,
	filePath string,
) (
	azblob.BlockBlobURL,
	error,
) {
	containerUrl, err := t.
		ContainerUrlProvider.
		CreateContainerUrl(containerName)

	if err != nil {
		return azblob.BlockBlobURL{}, err
	}

	blobUrl := containerUrl.NewBlockBlobURL(
		filePath,
	)

	return blobUrl, nil
}
