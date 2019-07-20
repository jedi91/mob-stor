package azure

import (
	"github.com/Azure/azure-storage-blob-go/azblob"
)

type BlobCredentialProvider struct {
	accountName string
	accountKey  string
}

func (b BlobCredentialProvider) CreateCredential() (
	*azblob.SharedKeyCredential,
	error,
) {
	return azblob.NewSharedKeyCredential(
		b.accountName,
		b.accountKey,
	)
}
