// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !go1.24

package testcontext

import (
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	const checkTime = 20 * time.Second
	ctx := contextWithParameters(
		t,
		4,
		30*time.Second,
		60*time.Second,
	)
	if ctx == nil {
		t.Fatalf("No context returned from testutils.Context")
	}
	tDeadline, tOk := t.Deadline()
	ctxDeadline, ctxOk := ctx.Deadline()
	if tOk != ctxOk {
		t.Fatalf("The testutils.Context function does not correctly set up the deadline ('ok' value mismatch)")
	}
	if tOk {
		if !ctxDeadline.Before(tDeadline.Add(checkTime)) {
			t.Fatalf(
				"The testutils.Context function does not correctly set up the deadline (not enough time left for cleanup)",
			)
		}
	}
}
