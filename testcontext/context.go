// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package testcontext

import (
	"context"
	"testing"
	"time"
)

// Context creates a context with a deadline that allows for enough time to clean up a test before the testing framework
// unceremoniously kills the process. The desired time is expressed as a fraction of the time remaining until the hard
// timeout, such as 4 for 25%. The minimumCleanupTime and maximumCleanupTime clamp the remaining time.
//
// For a simpler to use version of this function call tofutestutils.Context.
func Context(t *testing.T, cleanupTimeFraction int64, minimumCleanupTime time.Duration, maximumCleanupTime time.Duration) context.Context {
	ctx := context.Background()
	if deadline, ok := t.Deadline(); ok {
		var cancel func()
		timeoutDuration := time.Until(deadline)
		cleanupSafety := min(max(
			timeoutDuration/time.Duration(cleanupTimeFraction), minimumCleanupTime), maximumCleanupTime,
		)
		ctx, cancel = context.WithDeadline(ctx, deadline.Add(-1*cleanupSafety))
		t.Cleanup(cancel)
	}
	return ctx
}
