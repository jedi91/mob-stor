// Copyright (c) Michael Kovacevich 2019. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package distribute

// Result of a transmission
type Result struct {
	Stor    string
	Success bool
	Error   error
}
