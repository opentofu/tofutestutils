// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tofutestutils

import (
	"testing"

	"github.com/opentofu/tofutestutils/testaws"
)

// AWS creates an AWS service for testing purposes.
func AWS(t *testing.T) testaws.AWSTestService {
	return testaws.New(t)
}
