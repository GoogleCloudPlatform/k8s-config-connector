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

package compute

import (
	"reflect"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func TestComputeOrganizationSecurityPolicyRuleSpecMappers(t *testing.T) {
	action := "allow"
	desc := "test rule"
	resourceID := "100"
	versionedExpr := "FIREWALL"
	preview := false

	spec := &krm.ComputeOrganizationSecurityPolicyRuleSpec{
		Action:      action,
		Description: &desc,
		Preview:     &preview,
		ResourceID:  &resourceID,
		Match: &krm.OrganizationSecurityPolicyRuleMatch{
			VersionedExpr: &versionedExpr,
			Config: &krm.OrganizationSecurityPolicyRuleMatchConfig{
				SrcIPRanges: []string{"10.0.0.0/8"},
			},
		},
	}

	mapCtx := &direct.MapContext{}
	proto := ComputeOrganizationSecurityPolicyRuleSpec_v1alpha1_ToProto(mapCtx, spec)
	if mapCtx.Err() != nil {
		t.Fatalf("unexpected error converting to proto: %v", mapCtx.Err())
	}

	if proto.GetAction() != action {
		t.Errorf("GetAction() = %q, want %q", proto.GetAction(), action)
	}
	if proto.GetPriority() != 100 {
		t.Errorf("GetPriority() = %d, want 100", proto.GetPriority())
	}

	convertedSpec := ComputeOrganizationSecurityPolicyRuleSpec_v1alpha1_FromProto(mapCtx, proto)
	if mapCtx.Err() != nil {
		t.Fatalf("unexpected error converting from proto: %v", mapCtx.Err())
	}

	if !reflect.DeepEqual(convertedSpec, spec) {
		t.Errorf("roundtrip mismatch: got %+v, want %+v", convertedSpec, spec)
	}

	if ComputeOrganizationSecurityPolicyRuleSpec_v1alpha1_ToProto(mapCtx, nil) != nil {
		t.Errorf("expected nil for nil input")
	}
	if ComputeOrganizationSecurityPolicyRuleSpec_v1alpha1_FromProto(mapCtx, nil) != nil {
		t.Errorf("expected nil for nil input")
	}
}
