// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package testlog

import (
	"log/slog"
	"testing"
)

// NewSlogHandler produces an slog.Handler that writes to t.
func NewSlogHandler(t *testing.T) slog.Handler {
	return newSlogHandler(t)
}

func newSlogHandler(t testingT) slog.Handler {
	return newAdapter(t)
}
