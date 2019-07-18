package router

import (
	"github.com/jedi91/mob-stor/providers"
)

// Router for Object Stores
type Router struct {
	Providers []providers.Provider
}

// Routes file data to the configured object stores
func (r Router) Route(
	data []byte,
	fileName string,
) []RouteResult {
	if r.inputsInvalid(
		data,
		fileName,
	) {
		return []RouteResult{}
	}

	results := []RouteResult{}
	for _, provider := range r.Providers {
		success := provider.Stor(
			data,
			fileName,
		)

		result := RouteResult{
			provider.GetName(),
			success,
		}

		results = append(
			results,
			result,
		)
	}

	return results
}

func (r Router) inputsInvalid(
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
