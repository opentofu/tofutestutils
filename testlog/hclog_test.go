// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package testlog

import (
	"strings"
	"testing"

	"github.com/hashicorp/go-hclog"
)

func TestHCLogAdapter(t *testing.T) {
	t2 := &fakeT{}
	logger := newAdapter(t2)
	const testString = "Hello world!"

	interceptLogger := hclog.NewInterceptLogger(nil)
	interceptLogger.RegisterSink(logger)
	interceptLogger.Log(hclog.Error, testString)
	for _, line := range t2.lines {
		if strings.Contains(line, testString) {
			t.Logf("✅ Found the test string in the log output.")
			return
		}
	}
	t.Fatalf("❌ Failed to find test string in the log output.")
}
