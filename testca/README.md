# Certificate authority

This folder contains a basic x509 certificate authority implementation for testing purposes. You can use it whenever you need a certificate for servers or clients.

```go
package your_test

import (
	"crypto/tls"
	"io"
	"net"
	"strconv"
	"testing"
	"time"

	"github.com/opentofu/tofutestutils"
	"github.com/opentofu/tofutestutils/testca"
	"github.com/opentofu/tofutestutils/testrandom"
)

func TestMySocket(t *testing.T) {
	// Configure a desired randomness and time source. You can use this to create deterministic behavior.
	currentTimeSource := time.Now
	ca := testca.New(t, testrandom.DeterministicSource(t), currentTimeSource)

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