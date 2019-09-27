// Copyright (c) Michael Kovacevich 2019. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package azure

import (
	"github.com/Azure/azure-storage-blob-go/azblob"
)

type BlobCredentialProvider struct {
	AccountName string
	AccountKey  string
}

func (b BlobCredentialProvider) CreateCredential() (
	*azblob.SharedKeyCredential,
	error,
) {
	return azblob.NewSharedKeyCredential(
		b.AccountName,
		b.AccountKey,
	)
}
