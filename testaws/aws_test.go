// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package testaws_test

import (
	"testing"

	"github.com/opentofu/tofutestutils/testaws"
)

func TestAWS(t *testing.T) {
	t.Parallel()
	awsService := testaws.New(t)
	t.Run("DynamoDB", func(t *testing.T) {
		t.Parallel()
		testDynamoDBService(t, awsService)
	})
	t.Run("IAM", func(t *testing.T) {
		t.Parallel()
		testIAMService(t, awsService)
	})
	t.Run("KMS", func(t *testing.T) {
		t.Parallel()
		testKMSService(t, awsService)
	})
	t.Run("S3", func(t *testing.T) {
		t.Parallel()
		testS3Service(t, awsService)
	})
	t.Run("STS", func(t *testing.T) {
		t.Parallel()
		testSTSService(t, awsService)
	})
}
