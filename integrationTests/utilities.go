// Copyright (c) Michael Kovacevich 2019. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package integrationTests

import (
	"fmt"
	"github.com/jedi91/mob-stor/distribute"
	"github.com/jedi91/mob-stor/transmit"
	"testing"
)

const boolTemplate = "Actual: %t | Expected: %t"
const dateFormat = "01-02-2006 15:04:05"

func checkForSuccess(
	results []distribute.Result,
) bool {
	success := true
	for _, result := range results {
		if !result.Success {
			fmt.Println(result.Error)
		}

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

func setupDistributor(
	transmitter transmit.Transmitter,
) distribute.Distributor {
	transmitters := []transmit.Transmitter{
		transmitter,
	}

	return distribute.Distributor{
		Transmitters: transmitters,
	}
}
