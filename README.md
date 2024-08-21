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

## AWS backend

# AWS test backend

This library implements an AWS test backend using [LocalStack](https://www.localstack.cloud/) and Docker. To use it, you will need a local Docker daemon running. You can use the backend as follows:

```go
package your_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/opentofu/tofutestutils"
)

func TestMyApp(t *testing.T) {
	awsBackend := tofutestutils.AWS(t)

	s3Connection := s3.NewFromConfig(awsBackend.Config(), func(options *s3.Options) {
		options.UsePathStyle = awsBackend.S3UsePathStyle()
	})

	if _, err := s3Connection.PutObject(
		context.TODO(),
		&s3.PutObjectInput{
			Key:    aws.String("test.txt"),
			Body:   bytes.NewReader([]byte("Hello world!")),
			Bucket: aws.String(awsBackend.S3Bucket()),
		},
	); err != nil {
		t.Fatalf("❌ Failed to put object (%v)", err)
	}
}
```

> [!WARNING]
> Always use the [test context](#test-context) for timeouts as described above so the backend has time to clean up the test container.

## HTTP proxy

This library implements a full HTTP proxy for testing purposes. You can use this package to test if a proxy configuration handles settings correctly. The simplest version to create a proxy is as follows:

```go
package your_test

import (
	"testing"

	"github.com/opentofu/tofutestutils"
)

func TestMyApp(t *testing.T) {
	proxy := tofutestutils.HTTPProxy(t)
	
	t.Setenv("http_proxy", proxy.HTTPProxy().String())
	t.Setenv("https_proxy", proxy.HTTPSProxy().String())
	
	// Proxy-using code here
}
```

> [!NOTE]
> You can use the HTTP proxy for HTTPS URLs. This is because the connection to the proxy is not necessarily encrypted, the proxy can perform certificate verification. If you use the `HTTPSProxy()` call, you should configure your HTTP clients to accept the certificate presented in the `proxy.CACert()` function.

You can customize the proxy behavior and force connections to go to a specific endpoint. For details, please check [the documentation for the testproxy package](https://pkg.go.dev/github.com/opentofu/tofutestutils/testproxy).