// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package testlog

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/hashicorp/go-hclog"
)

func newAdapter(t testingT) *testLogAdapter {
	adapter := &testLogAdapter{t: t}
	t.Cleanup(func() {
		_ = adapter.Close()
	})
	return adapter
}

type testLogAdapter struct {
	t   testingT
	buf []byte
}

func (t *testLogAdapter) Enabled(_ context.Context, _ slog.Level) bool {
	t.t.Helper()
	return true
}

func (t *testLogAdapter) Handle(_ context.Context, record slog.Record) error {
	t.t.Helper()
	t.t.Logf("%s\t%s", record.Level, record.Message)
	return nil
}

func (t *testLogAdapter) WithAttrs(_ []slog.Attr) slog.Handler {
	t.t.Helper()
	return t
}

func (t *testLogAdapter) WithGroup(_ string) slog.Handler {
	t.t.Helper()
	return t
}

// Accept implements a hclog SinkAdapter.
func (t *testLogAdapter) Accept(name string, level hclog.Level, msg string, args ...interface{}) {
	t.t.Helper()
	msg = fmt.Sprintf(msg, args...)
	t.t.Logf("%s\t%s\t%s", name, level.String(), msg)
}

// Printf implements a standardized way to write logs, e.g. for the testcontainers package.
func (t *testLogAdapter) Printf(format string, v ...interface{}) {
	t.t.Helper()
	t.t.Logf(format, v...)
}

// Write provides a Go log-compatible writer.
func (t *testLogAdapter) Write(p []byte) (int, error) {
	t.t.Helper()
	t.buf = append(t.buf, p...)
	i := 0
	for i < len(t.buf) {
		if t.buf[i] == '\n' {
			t.t.Logf("%s", strings.TrimRight(string(t.buf[:i]), "\r"))
			t.buf = t.buf[i+1:]
			i = 0
		} else {
			i++
		}
	}
	return len(p), nil
}

func (t *testLogAdapter) Close() error {
	t.t.Helper()
	if len(t.buf) > 0 {
		t.t.Logf("%s", t.buf)
	}
	return nil
}
