package tofutestutils

import (
	"log"
	"log/slog"
	"testing"

	"github.com/hashicorp/go-hclog"
	"github.com/opentofu/tofutestutils/testlog"
	"github.com/testcontainers/testcontainers-go"
)

// NewGoTestLogger returns a log.Logger implementation that writes to testing.T.
func NewGoLogger(t *testing.T) *log.Logger {
	return testlog.NewGoTestLogger(t)
}

// NewSlogHandler produces an slog.Handler that writes to t.
func NewSlogHandler(t *testing.T) slog.Handler {
	return testlog.NewSlogHandler(t)
}

// NewHCLogAdapter returns a hclog.SinkAdapter-compatible logger that logs into a test facility.
func NewHCLogAdapter(t *testing.T) hclog.SinkAdapter {
	return testlog.NewHCLogAdapter(t)
}

// NewTestContainersLogger produces a logger for testcontainers.
func NewTestContainersLogger(t *testing.T) testcontainers.Logging {
	return testlog.NewTestContainersLogger(t)
}
