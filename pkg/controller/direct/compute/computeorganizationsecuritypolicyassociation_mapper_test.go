// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package compute

import (
	"reflect"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func TestComputeOrganizationSecurityPolicyAssociationSpecMappers(t *testing.T) {
	attachmentID := "organizations/123456789"
	policyID := "987654321"
	resourceID := "test-association"

	spec := &krm.ComputeOrganizationSecurityPolicyAssociationSpec{
		AttachmentID: attachmentID,
		PolicyID:     &policyID,
		ResourceID:   &resourceID,
	}

	mapCtx := &direct.MapContext{}
	proto := ComputeOrganizationSecurityPolicyAssociationSpec_v1alpha1_ToProto(mapCtx, spec)
	if mapCtx.Err() != nil {
		t.Fatalf("unexpected error converting to proto: %v", mapCtx.Err())
	}

	if proto.GetAttachmentId() != attachmentID {
		t.Errorf("GetAttachmentId() = %q, want %q", proto.GetAttachmentId(), attachmentID)
	}
	if proto.GetSecurityPolicyId() != policyID {
		t.Errorf("GetSecurityPolicyId() = %q, want %q", proto.GetSecurityPolicyId(), policyID)
	}
	if proto.GetName() != resourceID {
		t.Errorf("GetName() = %q, want %q", proto.GetName(), resourceID)
	}

	convertedSpec := ComputeOrganizationSecurityPolicyAssociationSpec_v1alpha1_FromProto(mapCtx, proto)
	if mapCtx.Err() != nil {
		t.Fatalf("unexpected error converting from proto: %v", mapCtx.Err())
	}

	if !reflect.DeepEqual(convertedSpec, spec) {
		t.Errorf("roundtrip mismatch: got %+v, want %+v", convertedSpec, spec)
	}

	if ComputeOrganizationSecurityPolicyAssociationSpec_v1alpha1_ToProto(mapCtx, nil) != nil {
		t.Errorf("expected nil for nil input")
	}
	if ComputeOrganizationSecurityPolicyAssociationSpec_v1alpha1_FromProto(mapCtx, nil) != nil {
		t.Errorf("expected nil for nil input")
	}
}

func TestComputeOrganizationSecurityPolicyAssociationStatusMappers(t *testing.T) {
	displayName := "my-security-policy"

	status := &krm.ComputeOrganizationSecurityPolicyAssociationStatus{
		DisplayName: &displayName,
	}

	mapCtx := &direct.MapContext{}
	proto := ComputeOrganizationSecurityPolicyAssociationStatus_v1alpha1_ToProto(mapCtx, status)
	if mapCtx.Err() != nil {
		t.Fatalf("unexpected error converting to proto: %v", mapCtx.Err())
	}

	if proto.GetDisplayName() != displayName {
		t.Errorf("GetDisplayName() = %q, want %q", proto.GetDisplayName(), displayName)
	}

	convertedStatus := ComputeOrganizationSecurityPolicyAssociationStatus_v1alpha1_FromProto(mapCtx, proto)
	if mapCtx.Err() != nil {
		t.Fatalf("unexpected error converting from proto: %v", mapCtx.Err())
	}

	if !reflect.DeepEqual(convertedStatus, status) {
		t.Errorf("roundtrip mismatch: got %+v, want %+v", convertedStatus, status)
	}

	if ComputeOrganizationSecurityPolicyAssociationStatus_v1alpha1_ToProto(mapCtx, nil) != nil {
		t.Errorf("expected nil for nil input")
	}
	if ComputeOrganizationSecurityPolicyAssociationStatus_v1alpha1_FromProto(mapCtx, nil) != nil {
		t.Errorf("expected nil for nil input")
	}
}
