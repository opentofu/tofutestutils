# Logging

This package contains a set of logging tools that write to `*testing.T` as a log output. You can use `testlog.NewHCLogAdapter()`, `testlog.NewGoTestLogger()`, `testlog.NewSlogHandler()` or `testlog.NewTestContainersLogger()` to obtain a logger for your use case.