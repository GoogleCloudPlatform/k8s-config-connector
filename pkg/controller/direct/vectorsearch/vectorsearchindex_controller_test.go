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

package vectorsearch

import (
	"context"
	"testing"

	pb "cloud.google.com/go/vectorsearch/apiv1/vectorsearchpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vectorsearch/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestVectorsearchIndex_Export(t *testing.T) {
	ctx := context.Background()

	id := &krm.VectorsearchIndexIdentity{
		Project:    "test-project",
		Location:   "us-central1",
		Collection: "my-collection-123",
		Index:      "my-index-456",
	}

	actual := &pb.Index{
		Name:        "projects/test-project/locations/us-central1/collections/my-collection-123/indexes/my-index-456",
		DisplayName: "My Vector Index",
		Description: "A test index",
	}

	adapter := &IndexAdapter{
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
	if got, want := u.GetName(), "my-index-456"; got != want {
		t.Errorf("GetName() = %q, want %q", got, want)
	}

	// Verify GVK
	gvk := u.GroupVersionKind()
	if gvk.Group != "vectorsearch.cnrm.cloud.google.com" || gvk.Kind != "VectorsearchIndex" {
		t.Errorf("GroupVersionKind() = %v, want vectorsearch.cnrm.cloud.google.com/v1alpha1, Kind=VectorsearchIndex", gvk)
	}

	// Verify Spec fields (ParentRef, ResourceID)
	spec, found, err := unstructured.NestedMap(u.Object, "spec")
	if err != nil {
		t.Fatalf("NestedMap spec failed: %v", err)
	}
	if !found {
		t.Fatal("spec field not found in unstructured object")
	}

	parentRef, found, _ := unstructured.NestedString(spec, "parentRef", "external")
	if !found || parentRef != id.ParentString() {
		t.Errorf("spec.parentRef.external = %q, want %q", parentRef, id.ParentString())
	}

	resourceID, found, _ := unstructured.NestedString(spec, "resourceID")
	if !found || resourceID != "my-index-456" {
		t.Errorf("spec.resourceID = %q, want %q", resourceID, "my-index-456")
	}

	// Verify DisplayName and Description from proto
	displayName, found, _ := unstructured.NestedString(spec, "displayName")
	if !found || displayName != "My Vector Index" {
		t.Errorf("spec.displayName = %q, want %q", displayName, "My Vector Index")
	}

	description, found, _ := unstructured.NestedString(spec, "description")
	if !found || description != "A test index" {
		t.Errorf("spec.description = %q, want %q", description, "A test index")
	}
}
