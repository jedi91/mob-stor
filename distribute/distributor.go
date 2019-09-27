// Copyright (c) Michael Kovacevich 2019. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package distribute

import (
	"github.com/jedi91/mob-stor/transmit"
)

// Distributor for Object Stores
type Distributor struct {
	Transmitters []transmit.Transmitter
}

// Distribute file data to object stores
func (d Distributor) Distribute(
	data []byte,
	filePath string,
	containerName string,
) []Result {
	if d.inputsInvalid(
		data,
		filePath,
	) {
		return []Result{}
	}

	results := []Result{}
	for _, transmitter := range d.Transmitters {
		//TODO: Move into a go routine and use a channel to gather the results
		err := transmitter.Transmit(
			data,
			filePath,
			containerName,
		)

		result := Result{
			transmitter.GetName(),
			err == nil,
			err,
		}

		results = append(
			results,
			result,
		)
	}

	return results
}

func (d Distributor) inputsInvalid(
	data []byte,
	filePath string,
) bool {
	return data == nil ||
		len(data) == 0 ||
		len(filePath) == 0
}
