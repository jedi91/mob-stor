package mobStor

import (
	"testing"
)

const errorTemplate = "Actual: %t | Expected: %t"

func TestRouteObject(t *testing.T) {
	r := router{}
	data := []byte("test file")
	fileName := "TestFile"
	result := r.RouteObject(data, fileName)
	if result {
		return
	}

	t.Errorf(
		errorTemplate,
		result,
		true,
	)
}

func TestInvalidInputsNilData(t *testing.T) {
	r := router{}
	fileName := "TestFile"
	result := r.invalidInputs(nil, fileName)
	if result {
		return
	}

	t.Errorf(
		errorTemplate,
		result,
		true,
	)
}

func TestInvalidInputsEmptyData(t *testing.T) {
	r := router{}
	fileName := "TestFile"
	data := []byte("")
	result := r.invalidInputs(data, fileName)
	if result {
		return
	}

	t.Errorf(
		errorTemplate,
		result,
		true,
	)
}

func TestInvalidInputsEmptyFileName(t *testing.T) {
	r := router{}
	fileName := ""
	data := []byte("Some Test Data")
	result := r.invalidInputs(data, fileName)
	if result {
		return
	}

	t.Errorf(
		errorTemplate,
		result,
		true,
	)
}
