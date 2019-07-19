package distribute

import (
	"github.com/jedi91/mob-stor/transmit"
)

// Router for Object Stores
type Distributor struct {
	Transmitters []transmit.Transmitter
}

// Routes file data to the configured object stores
func (d Distributor) Distribute(
	data []byte,
	fileName string,
) []DistributeResult {
	if d.inputsInvalid(
		data,
		fileName,
	) {
		return []DistributeResult{}
	}

	results := []DistributeResult{}
	for _, transmitter := range d.Transmitters {
		success := transmitter.Stor(
			data,
			fileName,
		)

		result := DistributeResult{
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
