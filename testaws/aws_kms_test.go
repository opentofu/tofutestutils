// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package testaws_test

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/opentofu/tofutestutils"
	"github.com/opentofu/tofutestutils/testaws"
)

func testKMSService(t *testing.T, kmsService testaws.AWSKMSTestService) {
	ctx := tofutestutils.Context(t)
	kmsClient := kms.NewFromConfig(kmsService.Config())
	t.Logf("🔍 Checking if the KMS key from the AWS test service can be described...")
	if _, err := kmsClient.DescribeKey(ctx, &kms.DescribeKeyInput{
		KeyId: aws.String(kmsService.KMSKeyID()),
	}); err != nil {
		t.Fatalf("❌ %v", err)
	}
	t.Logf("✅ KMS works as intended.")
}
