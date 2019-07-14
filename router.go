package mobStor

type router struct {
	//TODO: Add array of providers
}

func (r router) RouteObject(
	data []byte,
	fileName string,
) bool {
	if r.invalidInputs(data, fileName) {
		return false
	}

	return true
}

func (r router) invalidInputs(
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
