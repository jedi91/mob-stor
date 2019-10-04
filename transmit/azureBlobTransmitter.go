// Copyright (c) Michael Kovacevich 2019. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package transmit

import (
	"context"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/jedi91/mob-stor/azure"
)

// AzureBlobTransmitter - Transmits objects to Azure Blob storage
type AzureBlobTransmitter struct {
	ContainerURLProvider azure.ContainerURLProvider
}

// GetName - Gets the name of the transmitter
func (t AzureBlobTransmitter) GetName() string {
	return "AzureBlob"
}

// Transmit - Transmits an object to an Azure Blob container
func (t AzureBlobTransmitter) Transmit(
	data []byte,
	filePath string,
	containerName string,
) error {
	blobURL, blobURLErr := t.createBlobURL(
		containerName,
		filePath,
	)

	if blobURLErr != nil {
		return blobURLErr
	}

	context := context.Background()
	options := azblob.UploadToBlockBlobOptions{
		BlockSize:   4 * 1024 * 1024, //TODO: Need to do this dynamically
		Parallelism: 16,              //TODO: Run tests with different inputs for this
	}

	_, uploadErr := azblob.UploadBufferToBlockBlob(
		context,
		data,
		blobURL,
		options,
	)

	return uploadErr
}

func (t AzureBlobTransmitter) createBlobURL(
	containerName string,
	filePath string,
) (
	azblob.BlockBlobURL,
	error,
) {
	containerURL, err := t.
		ContainerURLProvider.
		CreateContainerURL(containerName)

	if err != nil {
		return azblob.BlockBlobURL{}, err
	}

	blobURL := containerURL.NewBlockBlobURL(
		filePath,
	)

	return blobURL, nil
}
