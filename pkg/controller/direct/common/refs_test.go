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

	computepb "cloud.google.com/go/compute/apiv1/computepb"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/google/go-cmp/cmp"
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

type mockComputeRef struct {
	External  string `json:"external,omitempty"`
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

var _ refs.Ref = &mockComputeRef{}

func (r *mockComputeRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeAddress",
	}
}

func (r *mockComputeRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
}

func (r *mockComputeRef) GetExternal() string {
	return r.External
}

func (r *mockComputeRef) SetExternal(ref string) {
	r.External = ref
}

func (r *mockComputeRef) ValidateExternal(ref string) error {
	return nil
}

func (r *mockComputeRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return nil
}

type testComputeSpec struct {
	NatIps []mockComputeRef `json:"natIps,omitempty"`
}

func TestNormalizeManagedComputeURIs(t *testing.T) {
	pb := &computepb.RouterNat{
		NatIps: []string{
			"https://www.googleapis.com/compute/v1/projects/my-project/regions/us-central1/addresses/addr-1",
			"projects/my-project/regions/us-central1/addresses/addr-2",
		},
	}

	specFromProto := func(mapCtx *direct.MapContext, in *computepb.RouterNat) *testComputeSpec {
		if in == nil {
			return nil
		}
		spec := &testComputeSpec{}
		for _, ip := range in.NatIps {
			spec.NatIps = append(spec.NatIps, mockComputeRef{External: ip})
		}
		return spec
	}

	specToProto := func(mapCtx *direct.MapContext, in *testComputeSpec) *computepb.RouterNat {
		if in == nil {
			return nil
		}
		out := &computepb.RouterNat{}
		for _, ref := range in.NatIps {
			if ref.External != "" {
				out.NatIps = append(out.NatIps, ref.External)
			}
		}
		return out
	}

	mapCtx := &direct.MapContext{}
	err := NormalizeManagedComputeURIs(mapCtx, pb, specFromProto, specToProto)
	if err != nil {
		t.Fatalf("unexpected error normalizing URIs: %v", err)
	}

	want := []string{
		"projects/my-project/regions/us-central1/addresses/addr-1",
		"projects/my-project/regions/us-central1/addresses/addr-2",
	}

	if !cmp.Equal(pb.NatIps, want) {
		t.Errorf("unexpected normalized URIs (-want +got):\n%s", cmp.Diff(want, pb.NatIps))
	}
}
