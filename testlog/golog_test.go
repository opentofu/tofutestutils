// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package testlog

import (
	"testing"
)

func TestGoTestLogger(t *testing.T) {
	t2 := &fakeT{}
	const testString = "Hello world!"

	logger := newGoTestLogger(t2)
	logger.Print(testString)

	if len(t2.lines) != 1 {
		t.Fatalf("❌ Expected 1 line, got %d", len(t2.lines))
	}
	t2.RunCleanup()
	if len(t2.lines) != 1 {
		t.Fatalf("❌ Expected 1 line, got %d", len(t2.lines))
	}
	if t2.lines[0] != testString {
		t.Fatalf("❌ Expected 'Hello world!', got '%s'", t2.lines[0])
	}
	t.Logf("✅ Correctly logged text.")
}

func TestGoTestLoggerMultiline(t *testing.T) {
	t2 := &fakeT{}
	const testString1 = "Hello"
	const testString2 = "world!"
	const testString = testString1 + "\n" + testString2
	logger := newGoTestLogger(t2)
	logger.Print(testString)

	if len(t2.lines) != 2 {
		t.Fatalf("❌ Expected 2 lines, got %d", len(t2.lines))
	}
	t2.RunCleanup()
	if len(t2.lines) != 2 {
		t.Fatalf("❌ Expected 2 lines, got %d", len(t2.lines))
	}
	if t2.lines[0] != testString1 {
		t.Fatalf("❌ Expected '%s', got '%s'", testString1, t2.lines[0])
	}
	if t2.lines[1] != testString2 {
		t.Fatalf("❌ Expected '%s', got '%s'", testString2, t2.lines[0])
	}
	t.Logf("✅ Correctly logged multiline text.")
}
