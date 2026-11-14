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

package resourcemanager

import (
	"context"
	"testing"

	resourcemanagerpb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/resourcemanager/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestTagHoldAdapter_Export(t *testing.T) {
	ctx := context.Background()

	id := &krm.ResourceManagerTagHoldIdentity{
		TagValue: "1234567890",
		TagHold:  "hold-abc",
	}

	actual := &resourcemanagerpb.TagHold{
		Name:     "tagValues/1234567890/tagHolds/hold-abc",
		Holder:   "//compute.googleapis.com/compute/projects/test-project/regions/us-central1/instanceGroupManagers/igm-test",
		Origin:   "test-origin",
		HelpLink: "https://cloud.google.com/resource-manager/docs/tags",
	}

	adapter := &TagHoldAdapter{
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

	if got, want := u.GetName(), id.String(); got != want {
		t.Errorf("GetName() = %q, want %q", got, want)
	}

	gvk := u.GroupVersionKind()
	if gvk != krm.ResourceManagerTagHoldGVK {
		t.Errorf("GroupVersionKind() = %v, want %v", gvk, krm.ResourceManagerTagHoldGVK)
	}

	spec, found, err := unstructured.NestedMap(u.Object, "spec")
	if err != nil {
		t.Fatalf("NestedMap spec failed: %v", err)
	}
	if !found {
		t.Fatal("spec field not found in unstructured object")
	}

	parentRef, found, _ := unstructured.NestedString(spec, "parentRef", "external")
	if !found || parentRef != "tagValues/1234567890" {
		t.Errorf("parentRef.external = %q, want %q", parentRef, "tagValues/1234567890")
	}

	holder, found, _ := unstructured.NestedString(spec, "holder")
	if !found || holder != actual.Holder {
		t.Errorf("holder = %q, want %q", holder, actual.Holder)
	}

	origin, found, _ := unstructured.NestedString(spec, "origin")
	if !found || origin != actual.Origin {
		t.Errorf("origin = %q, want %q", origin, actual.Origin)
	}

	helpLink, found, _ := unstructured.NestedString(spec, "helpLink")
	if !found || helpLink != actual.HelpLink {
		t.Errorf("helpLink = %q, want %q", helpLink, actual.HelpLink)
	}
}
