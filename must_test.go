// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tofutestutils_test

import (
	"fmt"
	"testing"

	"github.com/opentofu/tofutestutils"
)

func TestMust(t *testing.T) {
	t.Logf("🔍 Checking if Must() panics with an error...")
	paniced := false
	func() {
		defer func() {
			e := recover()
			paniced = e != nil
		}()

		tofutestutils.Must(fmt.Errorf("this is an error"))
	}()
	if paniced == false {
		t.Fatalf("❌ The Must() function did not panic.")
	}
	t.Logf("✅ The Must() function paniced properly.")

	t.Logf("🔍 Checking if Must() does not panic with nil...")
	paniced = false
	func() {
		defer func() {
			e := recover()
			paniced = e != nil
		}()

		tofutestutils.Must(nil)
	}()
	if paniced != false {
		t.Fatalf("❌ The Must() function paniced.")
	}
	t.Logf("✅ The Must() function did not panic.")
}

func TestMust2(t *testing.T) {
	t.Logf("🔍 Checking if Must() panics with an error...")
	paniced := false

	func() {
		defer func() {
			e := recover()
			paniced = e != nil
		}()
		_ = tofutestutils.Must2(42, fmt.Errorf("this is an error"))
	}()
	if paniced == false {
		t.Fatalf("❌ The Must2() function did not panic.")
	}
	t.Logf("✅ The Must2() function paniced properly.")

	t.Logf("🔍 Checking if Must2() panics does not panic with nil and returns the correct value...")
	paniced = false
	returned := 0
	func() {
		defer func() {
			e := recover()
			paniced = e != nil
		}()

		returned = tofutestutils.Must2(42, nil)
	}()
	if paniced != false {
		t.Fatalf("❌ The Must2() function paniced.")
	}
	if returned != 42 {
		t.Fatalf("❌ The Must2() function did not return the correct value: %d.", returned)
	}

	t.Logf("✅ The Must2() function did not panic and returned the correct value.")
}
