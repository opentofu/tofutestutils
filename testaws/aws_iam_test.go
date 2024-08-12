// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package testaws_test

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/opentofu/tofutestutils"
	"github.com/opentofu/tofutestutils/testaws"
)

func testIAMService(t *testing.T, iamService testaws.AWSIAMTestService) {
	ctx := tofutestutils.Context(t)
	iamClient := iam.NewFromConfig(iamService.Config())
	t.Logf("\U0001FAAA Checking if the caller identity can be retrieved...")
	roles, err := iamClient.ListRoles(ctx, &iam.ListRolesInput{})
	if err != nil {
		t.Fatalf("❌ Failed to get caller identity: %v", err)
	}
	t.Logf("✅ %d roles returned.", len(roles.Roles))
}
