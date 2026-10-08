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

func TestVectorSearchCollection_Export(t *testing.T) {
	ctx := context.Background()

	id := &krm.VectorSearchCollectionIdentity{
		Project:    "test-project",
		Location:   "us-central1",
		Collection: "my-collection-123",
	}

	actual := &pb.Collection{
		Name:        "projects/test-project/locations/us-central1/collections/my-collection-123",
		DisplayName: "My Vector Collection",
		Description: "A test collection",
	}

	adapter := &vectorSearchCollectionAdapter{
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
	if got, want := u.GetName(), "my-collection-123"; got != want {
		t.Errorf("GetName() = %q, want %q", got, want)
	}

	// Verify GVK
	gvk := u.GroupVersionKind()
	if gvk.Group != "vectorsearch.cnrm.cloud.google.com" || gvk.Kind != "VectorSearchCollection" {
		t.Errorf("GroupVersionKind() = %v, want vectorsearch.cnrm.cloud.google.com/v1alpha1, Kind=VectorSearchCollection", gvk)
	}

	// Verify Spec fields (ProjectRef, Location, ResourceID)
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
	if !found || location != "us-central1" {
		t.Errorf("spec.location = %q, want %q", location, "us-central1")
	}

	resourceID, found, _ := unstructured.NestedString(spec, "resourceID")
	if !found || resourceID != "my-collection-123" {
		t.Errorf("spec.resourceID = %q, want %q", resourceID, "my-collection-123")
	}

	// Verify DisplayName and Description from proto
	displayName, found, _ := unstructured.NestedString(spec, "displayName")
	if !found || displayName != "My Vector Collection" {
		t.Errorf("spec.displayName = %q, want %q", displayName, "My Vector Collection")
	}

	description, found, _ := unstructured.NestedString(spec, "description")
	if !found || description != "A test collection" {
		t.Errorf("spec.description = %q, want %q", description, "A test collection")
	}
}
