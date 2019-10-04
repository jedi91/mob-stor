// Copyright (c) Michael Kovacevich 2019. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package transmit

// Interface for transmitters
type Transmitter interface {
	// Transmits an object to the implemented object store
	Transmit(
		data []byte,
		filePath string,
		containerName string,
	) error

	// Gets the name of the implemented object store
	GetName() string
}
