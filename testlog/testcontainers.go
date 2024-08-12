// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package testlog

import (
	"testing"

	"github.com/testcontainers/testcontainers-go"
)

// NewTestContainersLogger produces a logger for testcontainers.
func NewTestContainersLogger(t *testing.T) testcontainers.Logging {
	return newTestContainersLogger(t)
}

func newTestContainersLogger(t testingT) testcontainers.Logging {
	return newAdapter(t)
}
