// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tofutestutils

import (
	"testing"

	"github.com/opentofu/tofutestutils/testproxy"
)

// HTTPProxy starts an HTTP proxy and returns the connection details as a result.
func HTTPProxy(t *testing.T, options ...testproxy.HTTPOption) testproxy.HTTPService {
	return testproxy.HTTP(t, options...)
}
