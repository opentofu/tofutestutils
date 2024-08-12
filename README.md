# OpenTofu test utilities

This library contains the test fixtures and utilities for testing OpenTofu. This library will be changed to suit the needs of OpenTofu only, please do not use it for any other purpose.

## Test context

Some tests need to create external resources, such as cloud resources, which need to be cleaned up reliably when the test ends. Unfortunately, `go test` unceremoniously kills the test run if it encounters a timeout without leaving time for cleanup functions to run. You can call `tofutestutils.Context` to obtain a `context.Context` that will leave enough time to perform a cleanup:

```go
package your_test

import (
	"testing"
	"your"
	
	"github.com/opentofu/tofutestutils"
)

func TestMyApplication(t *testing.T) {
	ctx := tofutestutils.Context(t)
	
	your.App(ctx)
}
```

## Randomness

The random functions provide a number of randomness sources for testing purposes, some deterministic, some not.

To obtain a deterministic randomness source *tied to a test name*, you can use the `DeterministicSource` (implementing an `io.Reader`) as follows:

```go
package your_test

import (
	"testing"
	"your"
	
	"github.com/opentofu/tofutestutils"
)

func TestMyApp(t *testing.T) {
	randomness := tofutestutils.DeterministicRandomSource(t)
	
    your.App(randomness)	
}
```

For a full list of possible functions, please [check the Go docs](https://pkg.go.dev/github.com/opentofu/tofutestutils).

## Handling errors

This package also provides the `Must()` and `Must2()` functions to make test code easier to read. For example:

```go
package your_test

import (
	"fmt"
	"testing"

	"github.com/opentofu/tofutestutils"
)

func erroringFunction() error {
	return fmt.Errorf("this is an error")
}

func erroringFunctionWithReturn() (int, error) {
	return 42, fmt.Errorf("this is an error")
}

func TestMyApp(t *testing.T) {
	// This will cause a panic:
	tofutestutils.Must(erroringFunction())
}

func TestMyApp2(t *testing.T) {
	// This will also cause a panic:
	result := tofutestutils.Must2(erroringFunctionWithReturn())
	t.Logf("The number is: %d", result)
}
```