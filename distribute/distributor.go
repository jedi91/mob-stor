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
	fileName string,
) []Result {
	if d.inputsInvalid(
		data,
		fileName,
	) {
		return []Result{}
	}

	results := []Result{}
	for _, transmitter := range d.Transmitters {
		success := transmitter.Stor(
			data,
			fileName,
		)

		result := Result{
			transmitter.GetName(),
			success,
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
	fileName string,
) bool {
	dataIsNil := data == nil
	dataIsEmpty := len(data) == 0
	fileNameIsEmpty := len(fileName) == 0
	return dataIsNil ||
		dataIsEmpty ||
		fileNameIsEmpty
}
