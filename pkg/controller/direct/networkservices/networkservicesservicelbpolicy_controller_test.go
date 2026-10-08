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

package networkservices

import (
	"context"
	"testing"

	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestNetworkServicesServiceLBPolicy_Export(t *testing.T) {
	ctx := context.Background()

	id := &krm.NetworkServicesServiceLBPolicyIdentity{
		Project:         "test-project",
		Location:        "global",
		ServiceLbPolicy: "my-policy-123",
	}

	actual := &pb.ServiceLbPolicy{
		Name:        "projects/test-project/locations/global/serviceLbPolicies/my-policy-123",
		Description: "A test policy",
	}

	adapter := &NetworkServicesServiceLBPolicyAdapter{
		id:     id,
		actual: actual,
	}

	u, err := adapter.Export(ctx)
	if err != nil {
		t.Fatalf("Export failed: %v", err)
	}

	if u == nil {
		t.Fatal("Export returned nil unstructured object")
	}

	// Verify Kubernetes object name
	if got, want := u.GetName(), "my-policy-123"; got != want {
		t.Errorf("GetName() = %q, want %q", got, want)
	}

	// Verify GVK
	if got, want := u.GroupVersionKind(), krm.NetworkServicesServiceLBPolicyGVK; got != want {
		t.Errorf("GroupVersionKind() = %v, want %v", got, want)
	}

	// Verify Spec fields (ProjectRef, Location, Description)
	spec, found, err := unstructured.NestedMap(u.Object, "spec")
	if err != nil {
		t.Fatalf("NestedMap spec failed: %v", err)
	}
	if !found {
		t.Fatal("spec field not found in unstructured object")
	}

	projectRef, found, _ := unstructured.NestedString(spec, "projectRef", "external")
	if !found || projectRef != "test-project" {
		t.Errorf("spec.projectRef.external = %q, want %q", projectRef, "test-project")
	}

	location, found, _ := unstructured.NestedString(spec, "location")
	if !found || location != "global" {
		t.Errorf("spec.location = %q, want %q", location, "global")
	}

	description, found, _ := unstructured.NestedString(spec, "description")
	if !found || description != "A test policy" {
		t.Errorf("spec.description = %q, want %q", description, "A test policy")
	}
}
