// Copyright (c) Michael Kovacevich 2019. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package transmit

type Transmitter interface {
	Transmit(
		data []byte,
		filePath string,
		containerName string,
	) error

	GetName() string
}
