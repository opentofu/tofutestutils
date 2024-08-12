# Randomness sources

This folder contains a number of randomness sources for testing purposes, some deterministic, some not.

To obtain a deterministic randomness source *tied to a test name*, you can use the `DeterministicSource` (implementing an `io.Reader`) as follows:

```go
package your_test

import (
	"testing"
	"your"
	
	"github.com/opentofu/tofutestutils/testrandom"
)

func TestMyApp(t *testing.T) {
	randomness := testrandom.DeterministicSource(t)
	
    your.App(randomness)	
}
```

For a full list of possible functions, please [check the Go docs](https://pkg.go.dev/github.com/opentofu/tofutestutils/testrandom).