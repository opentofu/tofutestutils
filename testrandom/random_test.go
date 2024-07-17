// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package testrandom_test

import (
	"strings"
	"testing"

	"github.com/opentofu/tofutestutils/testrandom"
)

func TestDeterministicID(t *testing.T) {
	const idLength = 12
	if t.Name() != "TestDeterministicID" {
		t.Fatalf(
			"The test name has changed, please update the test as it is used for seeding the random number " +
				"generator.",
		)
	}
	if id := testrandom.DeterministicID(
		t,
		idLength,
		testrandom.CharacterRangeAlphaNumeric,
	); id != "2Uw74WyQkh6P" {
		t.Fatalf(
			"Incorrect first pseudo-random ID returned: %s (the returned ID depends on the test name, make "+
				"sure to verify and update if you changed the test name)",
			id,
		)
	}
	if id := testrandom.DeterministicID(
		t,
		idLength,
		testrandom.CharacterRangeAlphaNumeric,
	); id != "F56iE3wkX1wR" {
		t.Fatalf(
			"Incorrect second pseudo-random ID returned: %s (the returned ID depends on the test name, make "+
				"sure to verify and update if you changed the test name)",
			id,
		)
	}
}

func TestIDPrefix(t *testing.T) {
	const testPrefix = "test-"
	const idLength = 12
	id := testrandom.IDPrefix(testPrefix, idLength, testrandom.CharacterRangeAlphaNumeric)
	if len(id) != idLength+len(testPrefix) {
		t.Fatalf("Incorrect random ID length: %s", id)
	}
	if !strings.HasPrefix(id, testPrefix) {
		t.Fatalf("Missing prefix: %s", id)
	}
}

func TestID(t *testing.T) {
	const idLength = 12
	id := testrandom.ID(idLength, testrandom.CharacterRangeAlphaNumeric)
	if len(id) != idLength {
		t.Fatalf("Incorrect random ID length: %s", id)
	}
}

func TestDeterministicInt(t *testing.T) {
	if t.Name() != "TestDeterministicInt" {
		t.Fatalf(
			"The test name has changed, please update the test as it is used for seeding the random number " +
				"generator.",
		)
	}
	if i := testrandom.DeterministicInt(
		t,
		1,
		42,
	); i != 39 {
		t.Fatalf(
			"Incorrect first pseudo-random int returned: %d (the returned int depends on the test name, make "+
				"sure to verify and update if you changed the test name)",
			i,
		)
	}
	if i := testrandom.DeterministicInt(
		t,
		1,
		42,
	); i != 17 {
		t.Fatalf(
			"Incorrect second pseudo-random int returned: %d (the returned int depends on the test name, make "+
				"sure to verify and update if you changed the test name)",
			i,
		)
	}
}

func TestInt(t *testing.T) {
	i := testrandom.Int(1, 42)
	if i < 1 || i > 42 {
		t.Fatalf("Invalid random integer returned %d (out of range)", i)
	}
}
