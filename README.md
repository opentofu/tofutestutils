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

## Certificate authority

When you need an x509 certificate for a server or a client, you can use the `tofutestutils.CA` function to obtain a `testca.CertificateAuthority` implementation using a pseudo-random number generator. You can use this to create a certificate for a socket server:

```go
package your_test

import (
	"crypto/tls"
	"io"
	"net"
	"strconv"
	"testing"

	"github.com/opentofu/tofutestutils"
)

func TestMySocket(t *testing.T) {
	ca := tofutestutils.CA(t)

	// Server side:
	tlsListener := tofutestutils.Must2(tls.Listen("tcp", "127.0.0.1:0", ca.CreateLocalhostServerCert().GetServerTLSConfig()))
	go func() {
		conn, serverErr := tlsListener.Accept()
		if serverErr != nil {
			return
		}
		defer func() {
			_ = conn.Close()
		}()
		_, _ = conn.Write([]byte("Hello world!"))
	}()

	// Client side:
	port := tlsListener.Addr().(*net.TCPAddr).Port
	client := tofutestutils.Must2(tls.Dial("tcp", net.JoinHostPort("127.0.0.1", strconv.Itoa(port)), ca.GetClientTLSConfig()))
	defer func() {
		_ = client.Close()
	}()

	t.Logf("%s", tofutestutils.Must2(io.ReadAll(client)))
}
```
