package azure

import (
	"fmt"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"github.com/Azure/azure-storage-blob-go/azblob"
)

type ContainerUrlProvider struct {
	credsProvider BlobCredentialProvider
}

func (c ContainerUrlProvider) CreateContainerUrl(
	containerName string,
) (
	azblob.ContainerURL,
	error,
) {
	pipeline, err := c.createPipeline()
	if err != nil {
		return azblob.ContainerURL{}, err
	}

	URL, _ := url.Parse(
		fmt.Sprintf(
			"https://%s.blob.core.windows.net/%s",
			c.credsProvider.accountName,
			containerName,
		),
	)

	containerUrl = azblob.NewContainerURL(
		*URL,
		pipeline,
	)
}

func (c ContainerUrlProvider) createPipeline() (
	pipeline.Pipeline,
	error,
) {
	creds, err := c.credsProvider.CreateCredential()
	if err != nil {
		return nil, err
	}

	pipeline := azblob.NewPipeline(
		creds,
		azblob.PipelineOptions{},
	)

	return pipeline, nil
}
