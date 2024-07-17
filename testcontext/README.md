# Test context

Some tests need to create external resources, such as cloud resources, which need to be cleaned up reliably when the test ends. Unfortunately, `go test` unceremoniously kills the test run if it encounters a timeout without leaving time for cleanup functions to run.

This package solves this problem by creating a `context.Context` that has a timeout earlier than the hard timeout of `tofu test`. You can pass this context to any parts of the application that need to abort before the test aborts.

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