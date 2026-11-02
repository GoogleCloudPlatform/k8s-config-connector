// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mockcompute

import (
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestOrganizationSecurityPolicyAssociationsValidation(t *testing.T) {
	ctx := context.Background()
	env := &common.MockEnvironment{}
	memStorage := storage.NewInMemoryStorage()
	service := &OrganizationSecurityPoliciesV1{
		MockService: &MockService{
			MockEnvironment:   env,
			storage:           memStorage,
			computeOperations: newComputeOperationsService(memStorage),
		},
	}

	// Test ListAssociations without targetResource
	_, err := service.ListAssociations(ctx, &pb.ListAssociationsOrganizationSecurityPolicyRequest{})
	if err == nil {
		t.Fatalf("expected error for empty targetResource, got nil")
	}
	if status.Code(err) != codes.InvalidArgument {
		t.Fatalf("expected InvalidArgument code, got %v", status.Code(err))
	}

	// Create a dummy policy so RemoveAssociation can look up the policy
	policyFQN := "locations/global/securityPolicies/123456789"
	dummyPolicy := &pb.SecurityPolicy{
		Id:       PtrTo(uint64(123456789)),
		SelfLink: PtrTo("https://compute.googleapis.com/compute/v1/" + policyFQN),
	}
	if err := memStorage.Create(ctx, policyFQN, dummyPolicy); err != nil {
		t.Fatalf("failed to create dummy policy: %v", err)
	}

	// Test RemoveAssociation for non-existent association
	_, err = service.RemoveAssociation(ctx, &pb.RemoveAssociationOrganizationSecurityPolicyRequest{
		SecurityPolicy: "123456789",
		Name:           PtrTo("non-existent-assoc"),
	})
	if err == nil {
		t.Fatalf("expected error for non-existent association, got nil")
	}
	if status.Code(err) != codes.InvalidArgument {
		t.Fatalf("expected InvalidArgument code, got %v", status.Code(err))
	}
}

func TestOrganizationSecurityPolicyRules(t *testing.T) {
	ctx := context.Background()
	env := &common.MockEnvironment{}
	memStorage := storage.NewInMemoryStorage()
	service := &OrganizationSecurityPoliciesV1{
		MockService: &MockService{
			MockEnvironment:   env,
			storage:           memStorage,
			computeOperations: newComputeOperationsService(memStorage),
		},
	}

	policyFQN := "locations/global/securityPolicies/123456789"
	dummyPolicy := &pb.SecurityPolicy{
		Id:       PtrTo(uint64(123456789)),
		SelfLink: PtrTo("https://compute.googleapis.com/compute/v1/" + policyFQN),
	}
	if err := memStorage.Create(ctx, policyFQN, dummyPolicy); err != nil {
		t.Fatalf("failed to create dummy policy: %v", err)
	}

	// 1. AddRule
	addOp, err := service.AddRule(ctx, &pb.AddRuleOrganizationSecurityPolicyRequest{
		SecurityPolicy: "123456789",
		SecurityPolicyRuleResource: &pb.SecurityPolicyRule{
			Action:      PtrTo("allow"),
			Description: PtrTo("initial rule"),
			Priority:    PtrTo(int32(100)),
			Match: &pb.SecurityPolicyRuleMatcher{
				Config: &pb.SecurityPolicyRuleMatcherConfig{
					SrcIpRanges: []string{"192.168.0.0/16"},
				},
				VersionedExpr: PtrTo("SRC_IPS_V1"),
			},
		},
	})
	if err != nil {
		t.Fatalf("AddRule failed: %v", err)
	}
	if addOp.GetOperationType() != "addOrganizationSecurityPolicyRule" {
		t.Errorf("unexpected operation type: %s", addOp.GetOperationType())
	}

	// 2. GetRule
	rule, err := service.GetRule(ctx, &pb.GetRuleOrganizationSecurityPolicyRequest{
		SecurityPolicy: "123456789",
		Priority:       PtrTo(int32(100)),
	})
	if err != nil {
		t.Fatalf("GetRule failed: %v", err)
	}
	if rule.GetDescription() != "initial rule" {
		t.Errorf("unexpected description: %s", rule.GetDescription())
	}

	// 3. PatchRule
	patchOp, err := service.PatchRule(ctx, &pb.PatchRuleOrganizationSecurityPolicyRequest{
		SecurityPolicy: "123456789",
		Priority:       PtrTo(int32(100)),
		SecurityPolicyRuleResource: &pb.SecurityPolicyRule{
			Description: PtrTo("updated rule"),
		},
	})
	if err != nil {
		t.Fatalf("PatchRule failed: %v", err)
	}
	if patchOp.GetOperationType() != "patchOrganizationSecurityPolicyRule" {
		t.Errorf("unexpected operation type: %s", patchOp.GetOperationType())
	}

	// Verify patch
	rule, err = service.GetRule(ctx, &pb.GetRuleOrganizationSecurityPolicyRequest{
		SecurityPolicy: "123456789",
		Priority:       PtrTo(int32(100)),
	})
	if err != nil {
		t.Fatalf("GetRule failed: %v", err)
	}
	if rule.GetDescription() != "updated rule" {
		t.Errorf("unexpected description after patch: %s", rule.GetDescription())
	}

	// 4. RemoveRule
	removeOp, err := service.RemoveRule(ctx, &pb.RemoveRuleOrganizationSecurityPolicyRequest{
		SecurityPolicy: "123456789",
		Priority:       PtrTo(int32(100)),
	})
	if err != nil {
		t.Fatalf("RemoveRule failed: %v", err)
	}
	if removeOp.GetOperationType() != "removeOrganizationSecurityPolicyRule" {
		t.Errorf("unexpected operation type: %s", removeOp.GetOperationType())
	}

	// 5. GetRule non-existent
	_, err = service.GetRule(ctx, &pb.GetRuleOrganizationSecurityPolicyRequest{
		SecurityPolicy: "123456789",
		Priority:       PtrTo(int32(100)),
	})
	if err == nil {
		t.Fatalf("expected error for removed rule, got nil")
	}
	if status.Code(err) != codes.InvalidArgument {
		t.Fatalf("expected InvalidArgument, got %v", status.Code(err))
	}
}
