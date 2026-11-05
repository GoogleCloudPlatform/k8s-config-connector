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
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/types"
)

func TestAppEngineServiceIdentity(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		want     *AppEngineServiceIdentity
		hasError bool
	}{
		{
			name:  "valid standard external reference",
			input: "apps/my-app/services/my-service",
			want: &AppEngineServiceIdentity{
				App:     "my-app",
				Service: "my-service",
			},
			hasError: false,
		},
		{
			name:     "invalid format - missing service",
			input:    "apps/my-app/services",
			want:     nil,
			hasError: true,
		},
		{
			name:     "invalid format - wrong keyword",
			input:    "apps/my-app/versions/v1",
			want:     nil,
			hasError: true,
		},
		{
			name:     "invalid format - empty string",
			input:    "",
			want:     nil,
			hasError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id := &AppEngineServiceIdentity{}
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

			if diff := cmp.Diff(tc.want, id); diff != "" {
				t.Errorf("AppEngineServiceIdentity mismatch (-want +got):\n%s", diff)
			}

			if got := id.String(); got != tc.input {
				t.Errorf("expected String() %q, got %q", tc.input, got)
			}

			wantParent := "apps/" + tc.want.App
			if got := id.ParentString(); got != wantParent {
				t.Errorf("expected ParentString() %q, got %q", wantParent, got)
			}

			if got := id.Host(); got != "appengine.googleapis.com" {
				t.Errorf("expected Host() %q, got %q", "appengine.googleapis.com", got)
			}
		})
	}
}

func TestAppEngineServiceRef(t *testing.T) {
	ref := &AppEngineServiceRef{
		Name:      "test-service",
		Namespace: "test-ns",
	}

	if got := ref.GetGVK(); got != AppEngineServiceGVK {
		t.Errorf("expected GetGVK() %v, got %v", AppEngineServiceGVK, got)
	}

	expectedNN := types.NamespacedName{Name: "test-service", Namespace: "test-ns"}
	if got := ref.GetNamespacedName(); got != expectedNN {
		t.Errorf("expected GetNamespacedName() %v, got %v", expectedNN, got)
	}

	external := "apps/my-app/services/my-service"
	ref.SetExternal(external)
	if got := ref.GetExternal(); got != external {
		t.Errorf("expected GetExternal() %q, got %q", external, got)
	}
	if ref.Name != "" || ref.Namespace != "" {
		t.Errorf("expected Name and Namespace to be cleared after SetExternal, got name=%q ns=%q", ref.Name, ref.Namespace)
	}

	if err := ref.ValidateExternal(external); err != nil {
		t.Errorf("unexpected ValidateExternal error: %v", err)
	}

	if err := ref.ValidateExternal("invalid"); err == nil {
		t.Error("expected ValidateExternal to fail on invalid input, got nil")
	}

	id, err := ref.ParseExternalToIdentity()
	if err != nil {
		t.Fatalf("unexpected ParseExternalToIdentity error: %v", err)
	}
	if id.String() != external {
		t.Errorf("expected identity string %q, got %q", external, id.String())
	}
}
