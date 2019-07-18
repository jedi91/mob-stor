package mobStor

import (
	"testing"
)

//TODO: Way to do any type in str format?
const boolTemplate = "Actual: %t | Expected: %t"
const intTemplate = "Actual: %d | Expected: %d"

type testProvider struct {
	result bool
	name   string
}

func (p testProvider) stor(
	data []byte,
	fileName string,
) bool {
	return p.result
}

func (p testProvider) getName() string {
	return p.name
}

func TestRouteSuccess(t *testing.T) {
	r := setupRouter(
		true,
		true,
	)

	data := []byte("test file")
	fileName := "TestFile"

	results := r.Route(
		data,
		fileName,
	)

	success := checkForSuccess(
		results,
	)

	checkExpectedBool(
		t,
		success,
		true,
	)
}

func TestRouteNilData(t *testing.T) {
	r := setupRouter(
		true,
		true,
	)

	fileName := "TestFile"
	results := r.Route(
		nil,
		fileName,
	)

	checkExpectedInt(
		t,
		len(results),
		0,
	)
}

func TestRouteEmptyData(t *testing.T) {
	r := setupRouter(
		true,
		true,
	)

	data := []byte("")
	fileName := "TestFile"
	results := r.Route(
		data,
		fileName,
	)

	checkExpectedInt(
		t,
		len(results),
		0,
	)
}

func TestRouteEmptyFileName(t *testing.T) {
	r := setupRouter(
		true,
		true,
	)

	fileName := ""
	data := []byte("Some Test Data")
	results := r.Route(
		data,
		fileName,
	)

	checkExpectedInt(
		t,
		len(results),
		0,
	)
}

func TestRouteSingleStorFails(t *testing.T) {
	r := setupRouter(
		true,
		false,
	)

	fileName := "TestFile"
	data := []byte("Some Test Data")
	results := r.Route(
		data,
		fileName,
	)

	success := checkForSuccess(
		results,
	)

	checkExpectedBool(
		t,
		success,
		false,
	)
}

func setupRouter(
	r1 bool,
	r2 bool,
) Router {
	p1 := testProvider{
		name:   "Test 1",
		result: r1,
	}

	p2 := testProvider{
		name:   "Test 2",
		result: r2,
	}

	return Router{
		[]provider{
			p1,
			p2,
		},
	}
}

func checkForSuccess(
	results []routeResult,
) bool {
	success := true
	for _, result := range results {
		success = success && result.success
	}

	return success
}

func checkExpectedBool(
	t *testing.T,
	actual bool,
	expected bool,
) {
	if actual == expected {
		return
	}

	t.Errorf(
		boolTemplate,
		actual,
		expected,
	)
}

func checkExpectedInt(
	t *testing.T,
	actual int,
	expected int,
) {
	if actual == expected {
		return
	}

	t.Errorf(
		intTemplate,
		actual,
		expected,
	)
}
