// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tofutestutils

import (
	"testing"
	"time"

	"github.com/opentofu/tofutestutils/testca"
)

// CA returns a certificate authority configured for the provided test. This implementation will configure the CA to use
// a pseudorandom source. You can call testca.New() for more configuration options.
func CA(t *testing.T) testca.CertificateAuthority {
	return testca.New(t, RandomSource(), time.Now)
}
