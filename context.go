// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tofutestutils

import (
	"context"
	"github.com/opentofu/tofutestutils/testcontext"
	"testing"
)

// Context creates a context with a deadline that allows for enough time to clean up a test before the testing framework
// unceremoniously kills the process. Starting with Go 1.24, this function is an alias for t.Context(). In Go 1.23 and
// lower this function reserves 25% of the test timeout, minimum 1 minute, maximum 5 minutes for cleanup.
func Context(t *testing.T) context.Context {
	return testcontext.Context(t)
}
