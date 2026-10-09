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

package v1beta1

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestSecretRefNormalize(t *testing.T) {
	ctx := context.TODO()

	readySecret := &unstructured.Unstructured{}
	readySecret.SetGroupVersionKind(SecretManagerSecretGVK)
	readySecret.SetName("my-secret")
	readySecret.SetNamespace("test-ns")
	if err := unstructured.SetNestedField(readySecret.Object, "projects/test-project/secrets/my-secret", "status", "externalRef"); err != nil {
		t.Fatalf("failed to set status.externalRef: %v", err)
	}

	reader := fake.NewClientBuilder().WithLists(&unstructured.UnstructuredList{Items: []unstructured.Unstructured{*readySecret}}).Build()

	t.Run("resolves by name", func(t *testing.T) {
		ref := &SecretRef{
			Name:      "my-secret",
			Namespace: "test-ns",
		}
		if err := ref.Normalize(ctx, reader, "test-ns"); err != nil {
			t.Fatalf("Normalize() unexpected error: %v", err)
		}
		want := "projects/test-project/secrets/my-secret"
		if ref.External != want {
			t.Errorf("ref.External = %q, want %q", ref.External, want)
		}

		id, err := ref.ParseExternalToIdentity()
		if err != nil {
			t.Fatalf("ParseExternalToIdentity() unexpected error: %v", err)
		}
		if id.String() != want {
			t.Errorf("id.String() = %q, want %q", id.String(), want)
		}
	})

	t.Run("resolves direct external", func(t *testing.T) {
		ref := &SecretRef{
			External: "projects/test-project/secrets/my-secret",
		}
		if err := ref.Normalize(ctx, reader, "test-ns"); err != nil {
			t.Fatalf("Normalize() unexpected error: %v", err)
		}
		want := "projects/test-project/secrets/my-secret"
		if ref.External != want {
			t.Errorf("ref.External = %q, want %q", ref.External, want)
		}
	})
}

func TestSecretVersionRefNormalize(t *testing.T) {
	ctx := context.TODO()

	readyVersion := &unstructured.Unstructured{}
	readyVersion.SetGroupVersionKind(SecretManagerSecretVersionGVK)
	readyVersion.SetName("my-version")
	readyVersion.SetNamespace("test-ns")
	if err := unstructured.SetNestedField(readyVersion.Object, "projects/test-project/secrets/my-secret/versions/1", "status", "externalRef"); err != nil {
		t.Fatalf("failed to set status.externalRef: %v", err)
	}

	reader := fake.NewClientBuilder().WithLists(&unstructured.UnstructuredList{Items: []unstructured.Unstructured{*readyVersion}}).Build()

	t.Run("resolves by name", func(t *testing.T) {
		ref := &SecretVersionRef{
			Name:      "my-version",
			Namespace: "test-ns",
		}
		if err := ref.Normalize(ctx, reader, "test-ns"); err != nil {
			t.Fatalf("Normalize() unexpected error: %v", err)
		}
		want := "projects/test-project/secrets/my-secret/versions/1"
		if ref.External != want {
			t.Errorf("ref.External = %q, want %q", ref.External, want)
		}

		id, err := ref.ParseExternalToIdentity()
		if err != nil {
			t.Fatalf("ParseExternalToIdentity() unexpected error: %v", err)
		}
		if id.String() != want {
			t.Errorf("id.String() = %q, want %q", id.String(), want)
		}
	})

	t.Run("resolves direct external", func(t *testing.T) {
		ref := &SecretVersionRef{
			External: "projects/test-project/secrets/my-secret/versions/1",
		}
		if err := ref.Normalize(ctx, reader, "test-ns"); err != nil {
			t.Fatalf("Normalize() unexpected error: %v", err)
		}
		want := "projects/test-project/secrets/my-secret/versions/1"
		if ref.External != want {
			t.Errorf("ref.External = %q, want %q", ref.External, want)
		}
	})
}

func TestSecretIdentityParse(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      *SecretIdentity
		wantError bool
	}{
		{
			name:  "valid external",
			input: "projects/p1/secrets/s1",
			want: &SecretIdentity{
				parent: &SecretParent{ProjectID: "p1"},
				id:     "s1",
			},
			wantError: false,
		},
		{
			name:      "invalid external",
			input:     "projects/p1/locations/l1/secrets/s1",
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id := &SecretIdentity{}
			err := id.FromExternal(tc.input)
			if tc.wantError {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.want, id, cmp.AllowUnexported(SecretIdentity{}, SecretParent{})); diff != "" {
				t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestSecretVersionIdentityParse(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      *SecretVersionIdentity
		wantError bool
	}{
		{
			name:  "valid external",
			input: "projects/p1/secrets/s1/versions/1",
			want: &SecretVersionIdentity{
				parent: &SecretVersionParent{ProjectID: "p1", SecretID: "s1"},
				id:     "1",
			},
			wantError: false,
		},
		{
			name:      "invalid external",
			input:     "projects/p1/locations/l1/secretversions/1",
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id := &SecretVersionIdentity{}
			err := id.FromExternal(tc.input)
			if tc.wantError {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.want, id, cmp.AllowUnexported(SecretVersionIdentity{}, SecretVersionParent{})); diff != "" {
				t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
