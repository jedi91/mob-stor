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
