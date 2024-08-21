// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package testlog

import (
	"log"
	"testing"
)

// NewGoTestLogger returns a log.Logger implementation that writes to testing.T.
func NewGoTestLogger(t *testing.T) *log.Logger {
	return newGoTestLogger(t)
}

func newGoTestLogger(t testingT) *log.Logger {
	return log.New(newAdapter(t), "", 0)
}
