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

package common

import (
	"context"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

type mockRef struct {
	normalized bool
}

func (m *mockRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{}
}

func (m *mockRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{}
}

func (m *mockRef) GetExternal() string {
	return ""
}

func (m *mockRef) SetExternal(string) {}

func (m *mockRef) ValidateExternal(ref string) error {
	return nil
}

func (m *mockRef) Normalize(ctx context.Context, reader client.Reader, otherNamespace string) error {
	m.normalized = true
	return nil
}

type testObj struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec testSpec `json:"spec,omitempty"`
}

type testSpec struct {
	RefA *mockRef
	RefB *mockRef
}

func (t *testObj) DeepCopyObject() runtime.Object {
	return t
}

func TestNormalizeReferencesWithSkipPath(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	kube := fake.NewClientBuilder().WithScheme(scheme).Build()

	t.Run("default behavior resolves both references", func(t *testing.T) {
		obj := &testObj{
			Spec: testSpec{
				RefA: &mockRef{},
				RefB: &mockRef{},
			},
		}

		err := NormalizeReferences(ctx, kube, obj, nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !obj.Spec.RefA.normalized {
			t.Errorf("expected RefA to be normalized")
		}
		if !obj.Spec.RefB.normalized {
			t.Errorf("expected RefB to be normalized")
		}
	})

	t.Run("SkipPath with leading dot", func(t *testing.T) {
		obj := &testObj{
			Spec: testSpec{
				RefA: &mockRef{},
				RefB: &mockRef{},
			},
		}

		err := NormalizeReferences(ctx, kube, obj, nil, SkipPath(".Spec.RefA"))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if obj.Spec.RefA.normalized {
			t.Errorf("expected RefA to be skipped, but it was normalized")
		}
		if !obj.Spec.RefB.normalized {
			t.Errorf("expected RefB to be normalized")
		}
	})

	t.Run("SkipPath without leading dot", func(t *testing.T) {
		obj := &testObj{
			Spec: testSpec{
				RefA: &mockRef{},
				RefB: &mockRef{},
			},
		}

		err := NormalizeReferences(ctx, kube, obj, nil, SkipPath("Spec.RefB"))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !obj.Spec.RefA.normalized {
			t.Errorf("expected RefA to be normalized")
		}
		if obj.Spec.RefB.normalized {
			t.Errorf("expected RefB to be skipped, but it was normalized")
		}
	})

	t.Run("multiple SkipPath options", func(t *testing.T) {
		obj := &testObj{
			Spec: testSpec{
				RefA: &mockRef{},
				RefB: &mockRef{},
			},
		}

		err := NormalizeReferences(ctx, kube, obj, nil, SkipPath(".Spec.RefA"), SkipPath(".Spec.RefB"))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if obj.Spec.RefA.normalized {
			t.Errorf("expected RefA to be skipped, but it was normalized")
		}
		if obj.Spec.RefB.normalized {
			t.Errorf("expected RefB to be skipped, but it was normalized")
		}
	})
}
