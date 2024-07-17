# OpenTofu test utilities

This library contains the test fixtures and utilities for testing OpenTofu. This library will be changed to suit the needs of OpenTofu only, please do not use it for any other purpose.

## Test context

Some tests need to create external resources, such as cloud resources, which need to be cleaned up reliably when the test ends. Unfortunately, `go test` unceremoniously kills the test run if it encounters a timeout without leaving time for cleanup functions to run. You can call `tofutestutils.Context` to obtain a `context.Context` that will leave enough time to perform a cleanup:

```go
package your_test

import (
	"your"
	
	"github.com/opentofu/tofutestutils"
)

func TestMyApplication(t *testing.T) {
	ctx := tofutestutils.Context(t)
	
	your.App(ctx)
}
```