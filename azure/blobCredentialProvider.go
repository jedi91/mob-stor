// Copyright (c) Michael Kovacevich 2019. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package azure

import (
	"github.com/Azure/azure-storage-blob-go/azblob"
)

// Provides a SharedKeyCredential from the azblob package
type BlobCredentialProvider struct {
	AccountName string
	AccountKey  string
}

// Creates the SharedKeyCredential
func (b BlobCredentialProvider) CreateCredential() (
	*azblob.SharedKeyCredential,
	error,
) {
	return azblob.NewSharedKeyCredential(
		b.AccountName,
		b.AccountKey,
	)
}
