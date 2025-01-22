// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !go1.24

package testcontext

import (
	"context"
	"testing"
	"time"
)

const defaultCleanupTimeFraction = 4
const defaultMinimumCleanupTime = time.Minute
const defaultMaximumCleanupTime = 5 * time.Minute

// Context creates a context with a deadline that allows for enough time to clean up a test before the testing framework
// unceremoniously kills the process. Starting with Go 1.24, this function is an alias for t.Context(). In Go 1.23 and
// lower this function reserves 25% of the test timeout, minimum 1 minute, maximum 5 minutes for cleanup.
func Context(t *testing.T) context.Context {
	return contextWithParameters(t, defaultCleanupTimeFraction, defaultMinimumCleanupTime, defaultMaximumCleanupTime)
}

func contextWithParameters(
	t *testing.T,
	cleanupTimeFraction uint,
	minimumCleanupTime time.Duration,
	maximumCleanupTime time.Duration,
) context.Context {
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
