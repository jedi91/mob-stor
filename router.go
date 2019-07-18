package mobStor

// Router for Object Stores
type Router struct {
	providers []provider
}

// Routes file data to the configured object stores
func (r Router) Route(
	data []byte,
	fileName string,
) []routeResult {
	if r.inputsInvalid(
		data,
		fileName,
	) {
		return []routeResult{}
	}

	results := []routeResult{}
	for _, provider := range r.providers {
		success := provider.stor(
			data,
			fileName,
		)

		result := routeResult{
			provider.getName(),
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
