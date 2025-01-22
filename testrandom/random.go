// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package testrandom

import (
	"hash/crc64"
	"math/rand"
	"strings"
	"sync"
	"testing"
	"time"
)

// The functions below contain an assortment of random ID generation functions, partially ported and improved from the
// github.com/opentofu/opentofu/internal/legacy/helper/acctest package.

var randomSources = map[string]*rand.Rand{} //nolint:gochecknoglobals //This variable stores the randomness sources for DeterministicID and needs to be global.
var randomLock = &sync.Mutex{}              //nolint:gochecknoglobals //This variable is required to lock the randomSources above.

// CharacterRange defines which characters to use for generating a random ID.
type CharacterRange string

const (
	CharacterRangeAlphaNumeric      CharacterRange = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	CharacterRangeAlphaNumericLower CharacterRange = "abcdefghijklmnopqrstuvwxyz0123456789"
	CharacterRangeAlphaNumericUpper CharacterRange = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	CharacterRangeAlpha             CharacterRange = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharacterRangeAlphaLower        CharacterRange = "abcdefghijklmnopqrstuvwxyz"
	CharacterRangeAlphaUpper        CharacterRange = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// DeterministicID generates a pseudo-random identifier for the given test, using its name as a seed for
// randomness. This function guarantees that when queried in order, the values are always the same as long as the name
// of the test doesn't change.
func DeterministicID(t *testing.T, length uint, characterSpace CharacterRange) string {
	return IDFromSource(DeterministicSource(t), length, characterSpace)
}

// ID returns a non-deterministic, pseudo-random identifier.
func ID(length uint, characterSpace CharacterRange) string {
	return IDFromSource(Source(), length, characterSpace)
}

// IDPrefix returns a random identifier with a given prefix. The prefix length does not count towards the
// length.
func IDPrefix(prefix string, length uint, characterSpace CharacterRange) string {
	return prefix + ID(length, characterSpace)
}

// IDFromSource generates a random ID with the specified length based on the provided random parameter.
func IDFromSource(random *rand.Rand, length uint, characterSpace CharacterRange) string {
	runes := []rune(characterSpace)
	var builder strings.Builder
	for i := uint(0); i < length; i++ {
		builder.WriteRune(runes[random.Intn(len(runes))])
	}
	return builder.String()
}

// DeterministicInt produces a deterministic random integer based on the test name between the specified min and
// max value (inclusive).
func DeterministicInt(t *testing.T, minValue int, maxValue int) int {
	return IntFromSource(DeterministicSource(t), minValue, maxValue)
}

// Int produces a random integer between the specified min and max value (inclusive).
func Int(minValue int, maxValue int) int {
	return IntFromSource(Source(), minValue, maxValue)
}

// IntFromSource produces a random integer between the specified min and max value (inclusive).
func IntFromSource(source *rand.Rand, minValue int, maxValue int) int {
	// The logic for this function was moved from mock_value_composer.go
	return source.Intn(maxValue+1-minValue) + minValue
}

// Source produces a rand.Rand randomness source that is non-deterministic.
func Source() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano())) //nolint:gosec // Disabling gosec linting because this ID is for testing only.
}

// DeterministicSource produces a rand.Rand that is deterministic based on the provided test name. It will always
// supply the same values as long as the test name doesn't change.
func DeterministicSource(t *testing.T) *rand.Rand {
	var random *rand.Rand
	name := t.Name()
	var ok bool
	randomLock.Lock()
	random, ok = randomSources[name]
	if !ok {
		seed := crc64.Checksum([]byte(name), crc64.MakeTable(crc64.ECMA))
		random = rand.New(rand.NewSource(int64(seed))) //nolint:gosec //This random number generator is intentionally deterministic.
		randomSources[name] = random
		t.Cleanup(func() {
			randomLock.Lock()
			defer randomLock.Unlock()
			delete(randomSources, name)
		})
	}
	randomLock.Unlock()
	return random
}
