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

package iam

import (
	"net/url"
	"testing"

	pb "cloud.google.com/go/iam/apiv2/iampb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestIAMDenyPolicySpec_ToProto(t *testing.T) {
	mapCtx := &direct.MapContext{}
	spec := &krm.IAMDenyPolicySpec{
		DisplayName: direct.LazyPtr("My Deny Policy"),
		Rules: []krm.PolicyRule{
			{
				Description: direct.LazyPtr("Deny all for storage"),
				DenyRule: &krm.DenyRule{
					DeniedPrincipals:  []string{"principalSet://goog/public:all"},
					DeniedPermissions: []string{"storage.googleapis.com/buckets.list"},
				},
			},
		},
	}

	got := IAMDenyPolicySpec_ToProto(mapCtx, spec)
	if mapCtx.Err() != nil {
		t.Fatalf("unexpected error: %v", mapCtx.Err())
	}

	want := &pb.Policy{
		DisplayName: "My Deny Policy",
		Rules: []*pb.PolicyRule{
			{
				Description: "Deny all for storage",
				Kind: &pb.PolicyRule_DenyRule{
					DenyRule: &pb.DenyRule{
						DeniedPrincipals:  []string{"principalSet://goog/public:all"},
						DeniedPermissions: []string{"storage.googleapis.com/buckets.list"},
					},
				},
			},
		},
	}

	if diff := cmp.Diff(got, want, protocmp.Transform()); diff != "" {
		t.Errorf("unexpected diff: %v", diff)
	}
}

func TestFullyQualifiedName(t *testing.T) {
	a := &denyPolicyAdapter{
		attachmentPoint: url.PathEscape("cloudresourcemanager.googleapis.com/projects/my-project"),
		policyID:        "my-policy",
	}

	got := a.fullyQualifiedName()
	want := "policies/cloudresourcemanager.googleapis.com%2Fprojects%2Fmy-project/denypolicies/my-policy"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
