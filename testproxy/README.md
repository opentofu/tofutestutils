# HTTP proxy

This package implements a full HTTP proxy for testing purposes. You can use this package to test if a proxy configuration handles settings correctly. The simplest version to create a proxy is as follows:

```go
package your_test

import (
	"testing"

	"github.com/opentofu/tofutestutils/testproxy"
)

func TestMyApp(t *testing.T) {
	proxy := testproxy.HTTP(t)
	
	t.Setenv("http_proxy", proxy.HTTPProxy().String())
	t.Setenv("https_proxy", proxy.HTTPSProxy().String())
	
	// Proxy-using code here
}
```

> [!NOTE]
> You can use the HTTP proxy for HTTPS URLs. This is because the connection to the proxy is not necessarily encrypted, the proxy can perform certificate verification. If you use the `HTTPSProxy()` call, you should configure your HTTP clients to accept the certificate presented in the `proxy.CACert()` function.

You can customize the proxy behavior and force connections to go to a specific endpoint. For details, please check [the documentation for this package](https://pkg.go.dev/github.com/opentofu/tofutestutils/testproxy).