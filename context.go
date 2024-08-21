// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tofutestutils

import (
	"context"
	"testing"

	"github.com/opentofu/tofutestutils/testcontext"
)

// Context returns a context configured for the test deadline. This function configures a context with 25% safety for
// any cleanup tasks to finish. For a more flexible function see testcontext.Context.
func Context(t *testing.T) context.Context {
	return testcontext.DefaultContext(t)
}
