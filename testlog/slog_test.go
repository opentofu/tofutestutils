// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package testlog

import (
	"log/slog"
	"strings"
	"testing"
)

func TestSlog(t *testing.T) {
	const testString = "Hello world!"

	tLogger := &fakeT{}

	logger := slog.New(newSlogHandler(tLogger))

	logger.Debug(testString)

	tLogger.RunCleanup()

	if len(tLogger.lines) != 1 {
		t.Fatalf("❌ Incorrect number of lines in log: %d", len(tLogger.lines))
	}

	if !strings.Contains(tLogger.lines[0], testString) {
		t.Fatalf("❌ The log output doesn't contain the test string: %s", tLogger.lines[0])
	}
	t.Logf("✅ Correctly logged text.")
}
