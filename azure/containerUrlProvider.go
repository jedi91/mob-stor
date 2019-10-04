// Copyright (c) Michael Kovacevich 2019. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package azure

import (
	"fmt"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"net/url"
)

const urlTemplate = "https://%s.blob.core.windows.net/%s"

// ContainerURLProvider - Provides a ContainerURL from the azblob package
type ContainerURLProvider struct {
	CredsProvider BlobCredentialProvider
}

// CreateContainerURL - Creates a ContainerURL
func (c ContainerURLProvider) CreateContainerURL(
	containerName string,
) (
	azblob.ContainerURL,
	error,
) {
	pipeline, pipelineErr := c.createPipeline()
	if pipelineErr != nil {
		return azblob.ContainerURL{}, pipelineErr
	}

	URL, urlErr := c.createURL(containerName)
	if urlErr != nil {
		return azblob.ContainerURL{}, urlErr
	}

	containerURL := azblob.NewContainerURL(
		*URL,
		pipeline,
	)

	return containerURL, nil
}

func (c ContainerURLProvider) createPipeline() (
	pipeline.Pipeline,
	error,
) {
	creds, err := c.CredsProvider.CreateCredential()
	if err != nil {
		return nil, err
	}

	pipeline := azblob.NewPipeline(
		creds,
		azblob.PipelineOptions{},
	)

	return pipeline, nil
}

func (c ContainerURLProvider) createURL(
	containerName string,
) (
	*url.URL,
	error,
) {
	return url.Parse(
		fmt.Sprintf(
			urlTemplate,
			c.CredsProvider.AccountName,
			containerName,
		),
	)
}
