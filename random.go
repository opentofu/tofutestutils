// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tofutestutils

import (
	"math/rand"
	"testing"

	"github.com/opentofu/tofutestutils/testrandom"
)

// DeterministicRandomID generates a pseudo-random identifier for the given test, using its name as a seed for
// randomness. This function guarantees that when queried in order, the values are always the same as long as the name
// of the test doesn't change.
func DeterministicRandomID(t *testing.T, length uint, characterSpace testrandom.CharacterRange) string {
	return testrandom.DeterministicID(t, length, characterSpace)
}

// RandomID returns a non-deterministic, pseudo-random identifier.
func RandomID(length uint, characterSpace testrandom.CharacterRange) string {
	return testrandom.ID(length, characterSpace)
}

// RandomIDPrefix returns a random identifier with a given prefix. The prefix length does not count towards the
// length.
func RandomIDPrefix(prefix string, length uint, characterSpace testrandom.CharacterRange) string {
	return testrandom.IDPrefix(prefix, length, characterSpace)
}

// RandomIDFromSource generates a random ID with the specified length based on the provided random parameter.
func RandomIDFromSource(random *rand.Rand, length uint, characterSpace testrandom.CharacterRange) string {
	return testrandom.IDFromSource(random, length, characterSpace)
}

// DeterministicRandomInt produces a deterministic random integer based on the test name between the specified min and
// max value (inclusive).
func DeterministicRandomInt(t *testing.T, minValue int, maxValue int) int {
	return testrandom.DeterministicInt(t, minValue, maxValue)
}

// RandomInt produces a random integer between the specified min and max value (inclusive).
func RandomInt(minValue int, maxValue int) int {
	return testrandom.Int(minValue, maxValue)
}

// RandomIntFromSource produces a random integer between the specified min and max value (inclusive).
func RandomIntFromSource(source *rand.Rand, minValue int, maxValue int) int {
	return testrandom.IntFromSource(source, minValue, maxValue)
}

// RandomSource produces a rand.Rand randomness source that is non-deterministic.
func RandomSource() *rand.Rand {
	return testrandom.Source()
}

// DeterministicRandomSource produces a rand.Rand that is deterministic based on the provided test name. It will always
// supply the same values as long as the test name doesn't change.
func DeterministicRandomSource(t *testing.T) *rand.Rand {
	return testrandom.DeterministicSource(t)
}
