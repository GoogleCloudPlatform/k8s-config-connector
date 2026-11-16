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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestParseNotebookRuntimeExternal(t *testing.T) {
	tests := []struct {
		name           string
		external       string
		wantProjectID  string
		wantLocation   string
		wantResourceID string
		wantErr        bool
	}{
		{
			name:           "valid relative externalRef",
			external:       "projects/my-project/locations/us-central1/notebookRuntimes/my-runtime",
			wantProjectID:  "my-project",
			wantLocation:   "us-central1",
			wantResourceID: "my-runtime",
		},
		{
			name:           "valid full URL externalRef",
			external:       "//aiplatform.googleapis.com/projects/my-project/locations/europe-west4/notebookRuntimes/my-runtime",
			wantProjectID:  "my-project",
			wantLocation:   "europe-west4",
			wantResourceID: "my-runtime",
		},
		{
			name:     "rejects SSRF location with slash",
			external: "projects/my-project/locations/127.0.0.1:34567/?/notebookRuntimes/my-runtime",
			wantErr:  true,
		},
		{
			name:     "rejects SSRF location without slash (query delimiter)",
			external: "projects/my-project/locations/127.0.0.1:34567?/notebookRuntimes/my-runtime",
			wantErr:  true,
		},
		{
			name:     "rejects SSRF location without slash (fragment delimiter)",
			external: "projects/my-project/locations/attacker.com#/notebookRuntimes/my-runtime",
			wantErr:  true,
		},
		{
			name:     "rejects path traversal in resourceID",
			external: "projects/my-project/locations/us-central1/notebookRuntimes/..",
			wantErr:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			parent, resourceID, err := ParseNotebookRuntimeExternal(tc.external)
			if (err != nil) != tc.wantErr {
				t.Fatalf("ParseNotebookRuntimeExternal(%q) error = %v, wantErr %v", tc.external, err, tc.wantErr)
			}
			if !tc.wantErr {
				if parent.ProjectID != tc.wantProjectID || parent.Location != tc.wantLocation || resourceID != tc.wantResourceID {
					t.Errorf("ParseNotebookRuntimeExternal(%q) = (%+v, %q), want (%s, %s, %s)",
						tc.external, parent, resourceID, tc.wantProjectID, tc.wantLocation, tc.wantResourceID)
				}
			}
		})
	}
}

func TestNewNotebookRuntimeIdentity_ValidatesLocation(t *testing.T) {
	ctx := context.Background()

	maliciousLocations := []string{
		"127.0.0.1:34567/?",
		"127.0.0.1:34567?",
		"attacker.com/?",
		"attacker.com?",
		"attacker.com#",
		"user@attacker.com?",
		"..",
		"us-central1/..",
		"",
	}

	for _, loc := range maliciousLocations {
		t.Run(loc, func(t *testing.T) {
			obj := &ColabRuntime{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-runtime",
					Namespace: "default",
				},
				Spec: ColabRuntimeSpec{
					Runtime_Parent: Runtime_Parent{
						ProjectRef: &refsv1beta1.ProjectRef{
							External: "test-project",
						},
						Location: loc,
					},
				},
			}
			if _, err := NewNotebookRuntimeIdentity(ctx, nil, obj); err == nil {
				t.Errorf("NewNotebookRuntimeIdentity with location=%q succeeded unexpectedly, want error", loc)
			}

			tmplObj := &ColabRuntimeTemplate{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-template",
					Namespace: "default",
				},
				Spec: ColabRuntimeTemplateSpec{
					RuntimeTemplate_Parent: RuntimeTemplate_Parent{
						ProjectRef: &refsv1beta1.ProjectRef{
							External: "test-project",
						},
						Location: loc,
					},
				},
			}
			if _, err := NewNotebookRuntimeTemplateIdentity(ctx, nil, tmplObj); err == nil {
				t.Errorf("NewNotebookRuntimeTemplateIdentity with location=%q succeeded unexpectedly, want error", loc)
			}
		})
	}
}
