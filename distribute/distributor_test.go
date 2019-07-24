package distribute

import (
	"errors"
	"github.com/jedi91/mob-stor/transmit"
	"testing"
)

//TODO: Way to do any type in str format?
const boolTemplate = "Actual: %t | Expected: %t"
const intTemplate = "Actual: %d | Expected: %d"

type testTransmitter struct {
	err  error
	name string
}

func (t testTransmitter) Transmit(
	data []byte,
	fileName string,
	path string,
) error {
	return t.err
}

func (t testTransmitter) GetName() string {
	return t.name
}

func TestDistributeSuccess(t *testing.T) {
	d := setupDistributor(
		nil,
		nil,
	)

	data := []byte("test file")
	fileName := "TestFile"
	path := "test/test"

	results := d.Distribute(
		data,
		fileName,
		path,
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

func TestDistributeNilData(t *testing.T) {
	d := setupDistributor(
		nil,
		nil,
	)

	fileName := "TestFile"
	path := "test/test"
	results := d.Distribute(
		nil,
		fileName,
		path,
	)

	checkExpectedInt(
		t,
		len(results),
		0,
	)
}

func TestDistributeEmptyData(t *testing.T) {
	d := setupDistributor(
		nil,
		nil,
	)

	data := []byte("")
	fileName := "TestFile"
	path := "test/test"
	results := d.Distribute(
		data,
		fileName,
		path,
	)

	checkExpectedInt(
		t,
		len(results),
		0,
	)
}

func TestDistributeEmptyFileName(t *testing.T) {
	d := setupDistributor(
		nil,
		nil,
	)

	fileName := ""
	data := []byte("Some Test Data")
	path := "test/test"
	results := d.Distribute(
		data,
		fileName,
		path,
	)

	checkExpectedInt(
		t,
		len(results),
		0,
	)
}

func TestDistributeSingleStorFails(t *testing.T) {
	d := setupDistributor(
		nil,
		errors.New("test"),
	)

	fileName := "TestFile"
	data := []byte("Some Test Data")
	path := "test/test"
	results := d.Distribute(
		data,
		fileName,
		path,
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

func setupDistributor(
	err1 error,
	err2 error,
) Distributor {
	t1 := testTransmitter{
		name: "Test 1",
		err:  err1,
	}

	t2 := testTransmitter{
		name: "Test 2",
		err:  err2,
	}

	return Distributor{
		Transmitters: []transmit.Transmitter{
			t1,
			t2,
		},
	}
}

func checkForSuccess(
	results []Result,
) bool {
	success := true
	for _, result := range results {
		success = success && result.Success
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
