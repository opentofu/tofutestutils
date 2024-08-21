// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package testaws_test

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/opentofu/tofutestutils"
	"github.com/opentofu/tofutestutils/testaws"
)

func testSTSService(t *testing.T, stsService testaws.AWSSTSTestService) {
	ctx := tofutestutils.Context(t)
	stsClient := sts.NewFromConfig(stsService.Config())
	t.Logf("\U0001FAAA Checking if the caller identity can be retrieved...")
	output, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		t.Fatalf("❌ Failed to get caller identity: %v", err)
	}
	t.Logf("✅ Caller identity: %s", *output.UserId)
}
