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
	filePath string,
	containerName string,
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
	filePath := "TestFile"
	containerName := "test/test"

	results := d.Distribute(
		data,
		filePath,
		containerName,
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

	filePath := "test/TestFile"
	containerName := "test"
	results := d.Distribute(
		nil,
		filePath,
		containerName,
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
	filePath := "test/TestFile"
	containerName := "test"
	results := d.Distribute(
		data,
		filePath,
		containerName,
	)

	checkExpectedInt(
		t,
		len(results),
		0,
	)
}

func TestDistributeEmptyFilePath(t *testing.T) {
	d := setupDistributor(
		nil,
		nil,
	)

	filePath := ""
	data := []byte("Some Test Data")
	containerName := "test"
	results := d.Distribute(
		data,
		filePath,
		containerName,
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

	filePath := "TestFile"
	data := []byte("Some Test Data")
	containerName := "test/test"
	results := d.Distribute(
		data,
		filePath,
		containerName,
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
