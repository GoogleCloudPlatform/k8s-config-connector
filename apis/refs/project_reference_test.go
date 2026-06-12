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

package refs

import (
	"context"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type mockProjectReader struct {
	client.Reader
}

func (m *mockProjectReader) Get(ctx context.Context, key types.NamespacedName, obj client.Object, opts ...client.GetOption) error {
	u, ok := obj.(*unstructured.Unstructured)
	if ok && u.GroupVersionKind().Kind == "Project" {
		u.SetName("my-resolved-project")
		u.SetNamespace(key.Namespace)
	}
	return nil
}

func TestProjectRef_Normalize_AlreadyExternal(t *testing.T) {
	r := &ProjectRef{
		External: "my-project",
	}

	err := r.Normalize(context.Background(), nil, "default")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if r.External != "projects/my-project" {
		t.Errorf("expected External to be %q, got %q", "projects/my-project", r.External)
	}
	if r.Name != "" {
		t.Errorf("expected Name to be cleared, got %q", r.Name)
	}
	if r.Namespace != "" {
		t.Errorf("expected Namespace to be cleared, got %q", r.Namespace)
	}
}

func TestProjectRef_Normalize_ResolvesFromK8s(t *testing.T) {
	r := &ProjectRef{
		Name:      "my-name",
		Namespace: "my-namespace",
	}

	reader := &mockProjectReader{}
	err := r.Normalize(context.Background(), reader, "default")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if r.External != "projects/my-resolved-project" {
		t.Errorf("expected External to be %q, got %q", "projects/my-resolved-project", r.External)
	}
	if r.Name != "" {
		t.Errorf("expected Name to be cleared, got %q", r.Name)
	}
	if r.Namespace != "" {
		t.Errorf("expected Namespace to be cleared, got %q", r.Namespace)
	}
}
