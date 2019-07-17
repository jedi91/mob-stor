package mobStor

import (
	"testing"
)

const errorTemplate = "Actual: %t | Expected: %t"

type testProvider struct {
}

func (p testProvider) stor(
	data []byte,
	fileName string,
) bool {
	return true
}

func (p testProvider) name() string {
	return "test"
}

func TestRoute(t *testing.T) {
	p := testProvider{}
	r := Router{[]provider{p}}
	data := []byte("test file")
	fileName := "TestFile"
	results := r.Route(data, fileName)
	success := true
	for _, result := range results {
		success = success && result.success
	}

	if success {
		return
	}

	t.Errorf(
		errorTemplate,
		success,
		true,
	)
}

func TestRouteNilData(t *testing.T) {
	p := testProvider{}
	r := Router{[]provider{p}}
	fileName := "TestFile"
	result := r.Route(nil, fileName)
	if result {
		t.Errorf(
			errorTemplate,
			result,
			false,
		)
	}
}

func TestRouteEmptyData(t *testing.T) {
	p := testProvider{}
	r := Router{[]provider{p}}
	fileName := "TestFile"
	data := []byte("")
	result := r.Route(data, fileName)
	if result {
		t.Errorf(
			errorTemplate,
			result,
			true,
		)
	}
}

func TestRouteEmptyFileName(t *testing.T) {
	p := testProvider{}
	r := Router{[]provider{p}}
	fileName := ""
	data := []byte("Some Test Data")
	result := r.Route(data, fileName)
	if result {
		t.Errorf(
			errorTemplate,
			result,
			true,
		)
	}
}
