// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package testlog

// testingT is a simplified interface to *testing.T. This interface is mainly used for internal testing purposes.
type testingT interface {
	Logf(format string, args ...interface{})
	Cleanup(func())
	Helper()
}
