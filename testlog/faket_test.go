// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package testlog

import (
	"fmt"
)

type fakeT struct {
	lines        []string
	cleanupFuncs []func()
}

func (f *fakeT) Helper() {
}

func (f *fakeT) Logf(format string, args ...interface{}) {
	f.lines = append(f.lines, fmt.Sprintf(format, args...))
}

func (f *fakeT) Cleanup(cleanupFunc func()) {
	f.cleanupFuncs = append(f.cleanupFuncs, cleanupFunc)
}

func (f *fakeT) RunCleanup() {
	for _, cleanupFunc := range f.cleanupFuncs {
		cleanupFunc()
	}
}
