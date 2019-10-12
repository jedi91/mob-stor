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
		result := d.transmit(
			data,
			filePath,
			containerName,
			transmitter,
		)

		results = append(
			results,
			result,
		)
	}

	return results
}

// DistributeConcurrently distributes file data to object sotres concurrently
func (d Distributor) DistributeConcurrently(
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

	resultsChan := make(chan Result)
	for _, transmitter := range d.Transmitters {
		go d.transmitConcurrently(
			data,
			filePath,
			containerName,
			transmitter,
			resultsChan,
		)
	}

	results := []Result{}
	for range d.Transmitters {
		results = append(
			results,
			<-resultsChan,
		)
	}

	return results
}

func (d Distributor) transmitConcurrently(
	data []byte,
	filePath string,
	containerName string,
	transmitter transmit.Transmitter,
	resultsChan chan Result,
) {
	resultsChan <- d.transmit(
		data,
		filePath,
		containerName,
		transmitter,
	)
}

func (d Distributor) transmit(
	data []byte,
	filePath string,
	containerName string,
	transmitter transmit.Transmitter,
) Result {
	err := transmitter.Transmit(
		data,
		filePath,
		containerName,
	)

	return Result{
		transmitter.GetName(),
		err == nil,
		err,
	}
}

func (d Distributor) inputsInvalid(
	data []byte,
	filePath string,
) bool {
	return data == nil ||
		len(data) == 0 ||
		len(filePath) == 0
}
