package distribute

import (
	"github.com/jedi91/mob-stor/transmit"
	"testing"
)

//TODO: Way to do any type in str format?
const boolTemplate = "Actual: %t | Expected: %t"
const intTemplate = "Actual: %d | Expected: %d"

type testTransmitter struct {
	result bool
	name   string
}

func (t testTransmitter) Stor(
	data []byte,
	fileName string,
) bool {
	return t.result
}

func (t testTransmitter) GetName() string {
	return t.name
}

func TestDistributeSuccess(t *testing.T) {
	d := setupDistributor(
		true,
		true,
	)

	data := []byte("test file")
	fileName := "TestFile"

	results := d.Distribute(
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

func TestDistributeNilData(t *testing.T) {
	d := setupDistributor(
		true,
		true,
	)

	fileName := "TestFile"
	results := d.Distribute(
		nil,
		fileName,
	)

	checkExpectedInt(
		t,
		len(results),
		0,
	)
}

func TestDistributeEmptyData(t *testing.T) {
	d := setupDistributor(
		true,
		true,
	)

	data := []byte("")
	fileName := "TestFile"
	results := d.Distribute(
		data,
		fileName,
	)

	checkExpectedInt(
		t,
		len(results),
		0,
	)
}

func TestDistributeEmptyFileName(t *testing.T) {
	d := setupDistributor(
		true,
		true,
	)

	fileName := ""
	data := []byte("Some Test Data")
	results := d.Distribute(
		data,
		fileName,
	)

	checkExpectedInt(
		t,
		len(results),
		0,
	)
}

func TestDistributeSingleStorFails(t *testing.T) {
	d := setupDistributor(
		true,
		false,
	)

	fileName := "TestFile"
	data := []byte("Some Test Data")
	results := d.Distribute(
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

func setupDistributor(
	r1 bool,
	r2 bool,
) Distributor {
	t1 := testTransmitter{
		name:   "Test 1",
		result: r1,
	}

	t2 := testTransmitter{
		name:   "Test 2",
		result: r2,
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
