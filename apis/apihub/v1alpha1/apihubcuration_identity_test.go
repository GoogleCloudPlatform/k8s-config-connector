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

package v1alpha1

import (
	"context"
	"testing"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestAPIHubCurationIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *APIHubCurationIdentity
		hasError bool
	}{
		{
			name:  "Full resource name",
			input: "projects/my-project/locations/us-central1/curations/my-curation",
			expected: &APIHubCurationIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Curation: "my-curation",
			},
			hasError: false,
		},
		{
			name:  "Full resource name with host",
			input: "apihub.googleapis.com/projects/my-project/locations/us-central1/curations/my-curation",
			expected: &APIHubCurationIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Curation: "my-curation",
			},
			hasError: false,
		},
		{
			name:     "Invalid format",
			input:    "projects/my-project/locations/us-central1/invalid/my-curation",
			expected: nil,
			hasError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id := &APIHubCurationIdentity{}
			err := id.FromExternal(tc.input)
			if tc.hasError {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.expected, id); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestAPIHubCurationIdentity_StringAndParentString(t *testing.T) {
	id := &APIHubCurationIdentity{
		Project:  "my-project",
		Location: "us-central1",
		Curation: "my-curation",
	}

	expectedStr := "projects/my-project/locations/us-central1/curations/my-curation"
	if got := id.String(); got != expectedStr {
		t.Errorf("String() = %q, want %q", got, expectedStr)
	}

	expectedParentStr := "projects/my-project/locations/us-central1"
	if got := id.ParentString(); got != expectedParentStr {
		t.Errorf("ParentString() = %q, want %q", got, expectedParentStr)
	}

	// Verify FromExternal round-trip
	parsed := &APIHubCurationIdentity{}
	if err := parsed.FromExternal(expectedStr); err != nil {
		t.Fatalf("unexpected error parsing %q: %v", expectedStr, err)
	}
	if parsed.String() != expectedStr {
		t.Errorf("round-trip failed: got %q, want %q", parsed.String(), expectedStr)
	}
}

func TestAPIHubCuration_GetIdentity(t *testing.T) {
	ctx := context.Background()

	// Test case 1: ResourceID is set explicitly
	customID := "my-custom-curation"
	location := "us-central1"
	curationObjWithCustomID := &APIHubCuration{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-curation-k8s-name",
			Namespace: "default",
		},
		Spec: APIHubCurationSpec{
			ProjectRef: &refsv1beta1.ProjectRef{
				External: "my-project",
			},
			Location:   &location,
			ResourceID: &customID,
		},
	}

	id, err := curationObjWithCustomID.GetIdentity(ctx, nil)
	if err != nil {
		t.Fatalf("unexpected error getting identity: %v", err)
	}
	expectedID := "projects/my-project/locations/us-central1/curations/my-custom-curation"
	if id.String() != expectedID {
		t.Errorf("GetIdentity() = %q, want %q", id.String(), expectedID)
	}

	// Test case 2: ResourceID is not set (should fail because it's required)
	curationObjWithDefaultID := &APIHubCuration{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-curation-k8s-name",
			Namespace: "default",
		},
		Spec: APIHubCurationSpec{
			ProjectRef: &refsv1beta1.ProjectRef{
				External: "my-project",
			},
			Location: &location,
		},
	}

	_, err = curationObjWithDefaultID.GetIdentity(ctx, nil)
	if err == nil {
		t.Fatal("expected error getting identity when ResourceID is not set, but got nil")
	}
}
