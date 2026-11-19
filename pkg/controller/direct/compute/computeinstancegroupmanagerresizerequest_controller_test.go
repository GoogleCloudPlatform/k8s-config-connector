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
	"context"
	"testing"

	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestInstanceGroupManagerResizeRequestAdapter_Export(t *testing.T) {
	ctx := context.Background()

	adapter := &InstanceGroupManagerResizeRequestAdapter{
		id: &krm.ComputeInstanceGroupManagerResizeRequestIdentity{
			Project:              "test-project",
			Zone:                 "us-central1-a",
			InstanceGroupManager: "my-igm",
			ResizeRequest:        "my-resize-request",
		},
		actual: &computepb.InstanceGroupManagerResizeRequest{
			Name:        proto.String("my-resize-request"),
			Description: proto.String("A sample resize request"),
			ResizeBy:    proto.Int32(2),
		},
	}

	u, err := adapter.Export(ctx)
	if err != nil {
		t.Fatalf("Export() failed: %v", err)
	}

	// Verify Kubernetes object metadata name is the short name, not the full resource path
	wantName := "my-resize-request"
	if u.GetName() != wantName {
		t.Errorf("u.GetName() = %q, want %q", u.GetName(), wantName)
	}

	// Verify parentRef is populated with ParentString
	parentRef, found, err := unstructured.NestedString(u.Object, "spec", "parentRef", "external")
	if err != nil || !found {
		t.Errorf("spec.parentRef.external not found: %v", err)
	}
	wantParent := "projects/test-project/zones/us-central1-a/instanceGroupManagers/my-igm"
	if parentRef != wantParent {
		t.Errorf("spec.parentRef.external = %q, want %q", parentRef, wantParent)
	}

	// Verify resourceID is populated
	resourceID, found, err := unstructured.NestedString(u.Object, "spec", "resourceID")
	if err != nil || !found {
		t.Errorf("spec.resourceID not found: %v", err)
	}
	if resourceID != wantName {
		t.Errorf("spec.resourceID = %q, want %q", resourceID, wantName)
	}

	// Verify mapped spec fields
	desc, found, err := unstructured.NestedString(u.Object, "spec", "description")
	if err != nil || !found || desc != "A sample resize request" {
		t.Errorf("spec.description = %q (found=%v, err=%v), want %q", desc, found, err, "A sample resize request")
	}

	resizeBy, found, err := unstructured.NestedInt64(u.Object, "spec", "resizeBy")
	if err != nil || !found || resizeBy != 2 {
		t.Errorf("spec.resizeBy = %d (found=%v, err=%v), want 2", resizeBy, found, err)
	}
}
