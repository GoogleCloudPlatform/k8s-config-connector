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
	"fmt"
	"testing"
	"time"

	computepb "cloud.google.com/go/compute/apiv1/computepb"
	artifactregistryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/artifactregistry/v1beta1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/google/go-cmp/cmp"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
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

type configurableRef struct {
	name        string
	namespace   string
	external    string
	resolvedExt string
	normErr     error
}

var _ refs.Ref = &configurableRef{}

func (r *configurableRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "kms.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "KMSCryptoKey",
	}
}

func (r *configurableRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{Name: r.name, Namespace: r.namespace}
}

func (r *configurableRef) GetExternal() string {
	return r.external
}

func (r *configurableRef) SetExternal(ext string) {
	r.external = ext
}

func (r *configurableRef) ValidateExternal(_ string) error {
	return nil
}

func (r *configurableRef) Normalize(_ context.Context, _ client.Reader, _ string) error {
	if r.normErr != nil {
		return r.normErr
	}
	if r.external == "" && r.resolvedExt != "" {
		r.external = r.resolvedExt
	}
	return nil
}

type deleteRefTestObj struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec deleteRefTestSpec `json:"spec,omitempty"`
}

type deleteRefTestSpec struct {
	PrimaryRef *configurableRef
	HealthyRef *configurableRef
	ProjectRef *refs.ProjectRef
}

func (o *deleteRefTestObj) DeepCopyObject() runtime.Object {
	return o
}

type failingGetReader struct {
	client.Reader
	getErr error
}

func (f *failingGetReader) Get(_ context.Context, _ client.ObjectKey, _ client.Object, _ ...client.GetOption) error {
	return f.getErr
}

func TestNormalizeReferences_DeletingSkipsUnresolvableRefs(t *testing.T) {
	refGVK := schema.GroupVersionKind{
		Group:   "kms.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "KMSCryptoKey",
	}
	refNN := types.NamespacedName{Namespace: "test-ns", Name: "my-key"}
	wantHealthyExternal := "projects/my-proj/locations/us-central1/keyRings/kr/cryptoKeys/ready-key"

	tests := []struct {
		name       string
		primaryRef *configurableRef
		projectRef *refs.ProjectRef
	}{
		{
			name: "missing_krm_ref",
			primaryRef: &configurableRef{
				name:    "my-key",
				normErr: k8s.NewReferenceNotFoundError(refGVK, refNN),
			},
		},
		{
			name: "unready_krm_ref",
			primaryRef: &configurableRef{
				name:    "my-key",
				normErr: k8s.NewReferenceNotReadyError(refGVK, refNN),
			},
		},
		{
			name: "missing_transitive_dependency",
			primaryRef: &configurableRef{
				name:    "my-key",
				normErr: k8s.NewTransitiveDependencyNotFoundError(refGVK, refNN),
			},
		},
		{
			name: "unready_transitive_dependency",
			primaryRef: &configurableRef{
				name:    "my-key",
				normErr: k8s.NewTransitiveDependencyNotReadyError(refGVK, refNN),
			},
		},
		{
			name: "missing_project_ref",
			projectRef: &refs.ProjectRef{
				Name: "deleted-project",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			kube := fake.NewClientBuilder().WithScheme(runtime.NewScheme()).Build()
			now := metav1.Now()
			obj := &deleteRefTestObj{
				ObjectMeta: metav1.ObjectMeta{
					Name:              "test-obj",
					Namespace:         "test-ns",
					DeletionTimestamp: &now,
				},
				Spec: deleteRefTestSpec{
					PrimaryRef: tc.primaryRef,
					HealthyRef: &configurableRef{
						name:        "ready-key",
						resolvedExt: wantHealthyExternal,
					},
					ProjectRef: tc.projectRef,
				},
			}

			err := NormalizeReferences(t.Context(), kube, obj, nil)

			if err != nil {
				t.Fatalf("NormalizeReferences(%s) error = %v, want nil", tc.name, err)
			}
			if got := obj.Spec.HealthyRef.GetExternal(); got != wantHealthyExternal {
				t.Errorf("NormalizeReferences(%s) HealthyRef.GetExternal() = %q, want %q", tc.name, got, wantHealthyExternal)
			}
			if tc.primaryRef != nil && tc.primaryRef.GetExternal() != "" {
				t.Errorf("NormalizeReferences(%s) PrimaryRef.GetExternal() = %q, want empty string", tc.name, tc.primaryRef.GetExternal())
			}
		})
	}
}

func TestNormalizeReferences_DeletingRejectsInvalidOrTransientErrors(t *testing.T) {
	scheme := runtime.NewScheme()
	baseKube := fake.NewClientBuilder().WithScheme(scheme).Build()

	tests := []struct {
		name       string
		reader     client.Reader
		primaryRef *configurableRef
		projectRef *refs.ProjectRef
	}{
		{
			name:   "invalid_external_format",
			reader: baseKube,
			primaryRef: &configurableRef{
				external: "invalid://format",
				normErr:  fmt.Errorf("invalid external format"),
			},
		},
		{
			name:   "empty_ref_without_name_or_external",
			reader: baseKube,
			primaryRef: &configurableRef{
				normErr: fmt.Errorf("must specify name or external"),
			},
		},
		{
			name:   "invalid_project_external_format",
			reader: baseKube,
			projectRef: &refs.ProjectRef{
				External: "projects/extra/segment",
			},
		},
		{
			name:   "transient_project_reader_error",
			reader: &failingGetReader{Reader: baseKube, getErr: fmt.Errorf("apiserver unavailable")},
			projectRef: &refs.ProjectRef{
				Name: "my-proj",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			now := metav1.Now()
			obj := &deleteRefTestObj{
				ObjectMeta: metav1.ObjectMeta{
					Name:              "test-obj",
					Namespace:         "test-ns",
					DeletionTimestamp: &now,
				},
				Spec: deleteRefTestSpec{
					PrimaryRef: tc.primaryRef,
					ProjectRef: tc.projectRef,
				},
			}

			err := NormalizeReferences(t.Context(), tc.reader, obj, nil)

			if err == nil {
				t.Errorf("NormalizeReferences(%s) error = nil, want non-nil error", tc.name)
			}
		})
	}
}

func TestNormalizeReferences_ActiveReconcileFailsOnUnresolvableRefs(t *testing.T) {
	refGVK := schema.GroupVersionKind{
		Group:   "kms.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "KMSCryptoKey",
	}
	refNN := types.NamespacedName{Namespace: "test-ns", Name: "my-key"}

	tests := []struct {
		name       string
		primaryRef *configurableRef
		projectRef *refs.ProjectRef
	}{
		{
			name: "missing_krm_ref",
			primaryRef: &configurableRef{
				name:    "my-key",
				normErr: k8s.NewReferenceNotFoundError(refGVK, refNN),
			},
		},
		{
			name: "unready_krm_ref",
			primaryRef: &configurableRef{
				name:    "my-key",
				normErr: k8s.NewReferenceNotReadyError(refGVK, refNN),
			},
		},
		{
			name: "missing_project_ref",
			projectRef: &refs.ProjectRef{
				Name: "deleted-project",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			kube := fake.NewClientBuilder().WithScheme(runtime.NewScheme()).Build()
			obj := &deleteRefTestObj{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-obj",
					Namespace: "test-ns",
				},
				Spec: deleteRefTestSpec{
					PrimaryRef: tc.primaryRef,
					ProjectRef: tc.projectRef,
				},
			}

			err := NormalizeReferences(t.Context(), kube, obj, nil)

			if err == nil {
				t.Errorf("NormalizeReferences(%s) error = nil, want non-nil error", tc.name)
			}
		})
	}
}

func TestNormalizeReferences_ArtifactRegistryRepositoryDeleting(t *testing.T) {
	unreadyKey := &unstructured.Unstructured{}
	unreadyKey.SetGroupVersionKind(kmsv1beta1.KMSCryptoKeyGVK)
	unreadyKey.SetNamespace("test-ns")
	unreadyKey.SetName("unready-key")

	tests := []struct {
		name      string
		kmsKeyRef *kmsv1beta1.KMSCryptoKeyRef
		upstreams []artifactregistryv1beta1.ArtifactRegistryRepositoryUpstreamPolicy
		extraObjs []runtime.Object
	}{
		{
			name:      "deleted_kms_key_ref",
			kmsKeyRef: &kmsv1beta1.KMSCryptoKeyRef{Name: "deleted-key"},
		},
		{
			name:      "unready_kms_key_ref",
			kmsKeyRef: &kmsv1beta1.KMSCryptoKeyRef{Name: "unready-key"},
			extraObjs: []runtime.Object{unreadyKey},
		},
		{
			name: "deleted_upstream_repository_ref",
			upstreams: []artifactregistryv1beta1.ArtifactRegistryRepositoryUpstreamPolicy{
				{
					ID:            direct.LazyPtr("up-1"),
					RepositoryRef: &artifactregistryv1beta1.ArtifactRegistryRepositoryRef{Name: "deleted-upstream"},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			kube := fake.NewClientBuilder().WithScheme(runtime.NewScheme()).WithRuntimeObjects(tc.extraObjs...).Build()
			now := metav1.Now()
			repo := &artifactregistryv1beta1.ArtifactRegistryRepository{
				ObjectMeta: metav1.ObjectMeta{
					Name:              "cmek-repo",
					Namespace:         "test-ns",
					DeletionTimestamp: &now,
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "test-proj",
					},
				},
				Spec: artifactregistryv1beta1.ArtifactRegistryRepositorySpec{
					Location:      "us-central1",
					Format:        direct.LazyPtr("DOCKER"),
					KMSKeyNameRef: tc.kmsKeyRef,
				},
			}
			if len(tc.upstreams) > 0 {
				repo.Spec.VirtualRepositoryConfig = &artifactregistryv1beta1.ArtifactRegistryRepositoryVirtualRepositoryConfig{
					UpstreamPolicies: tc.upstreams,
				}
			}

			normErr := NormalizeReferences(t.Context(), kube, repo, nil)
			id, idErr := repo.GetIdentity(t.Context(), kube)

			if normErr != nil {
				t.Fatalf("NormalizeReferences(%s) error = %v, want nil", tc.name, normErr)
			}
			if idErr != nil {
				t.Fatalf("GetIdentity(%s) error = %v, want nil", tc.name, idErr)
			}
			wantID := "projects/test-proj/locations/us-central1/repositories/cmek-repo"
			if got := id.String(); got != wantID {
				t.Errorf("GetIdentity(%s) = %q, want %q", tc.name, got, wantID)
			}
		})
	}
}

type stubDirectManager struct {
	manager.Manager
	kubeClient client.Client
	scheme     *runtime.Scheme
	recorder   record.EventRecorder
}

func (m *stubDirectManager) GetClient() client.Client                          { return m.kubeClient }
func (m *stubDirectManager) GetScheme() *runtime.Scheme                        { return m.scheme }
func (m *stubDirectManager) GetConfig() *rest.Config                           { return &rest.Config{} }
func (m *stubDirectManager) GetEventRecorderFor(_ string) record.EventRecorder { return m.recorder }

type stubJitter struct{}

func (s *stubJitter) WatchJitteredTimeout() time.Duration { return time.Second }
func (s *stubJitter) JitteredReenqueue(_ schema.GroupVersionKind, _ metav1.Object) (time.Duration, error) {
	return time.Second, nil
}

type realNormalizeARModel struct {
	findCalls int
	deletedID string
}

func (m *realNormalizeARModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	obj := &artifactregistryv1beta1.ArtifactRegistryRepository{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(op.GetUnstructured().Object, &obj); err != nil {
		return nil, err
	}
	if err := NormalizeReferences(ctx, op.Reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}
	id, err := obj.GetIdentity(ctx, op.Reader)
	if err != nil {
		return nil, err
	}
	return &realNormalizeARAdapter{id: id.String(), model: m}, nil
}

func (m *realNormalizeARModel) AdapterForURL(_ context.Context, _ string) (directbase.Adapter, error) {
	return nil, nil
}

type realNormalizeARAdapter struct {
	id    string
	model *realNormalizeARModel
}

func (a *realNormalizeARAdapter) Find(_ context.Context) (bool, error) {
	a.model.findCalls++
	return true, nil
}

func (a *realNormalizeARAdapter) Delete(_ context.Context, _ *directbase.DeleteOperation) (bool, error) {
	a.model.deletedID = a.id
	return true, nil
}

func (a *realNormalizeARAdapter) Create(_ context.Context, _ *directbase.CreateOperation) error {
	return nil
}

func (a *realNormalizeARAdapter) Update(_ context.Context, _ *directbase.UpdateOperation) error {
	return nil
}

func (a *realNormalizeARAdapter) Export(_ context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func TestNormalizeReferences_DirectReconcilerDeleteEndToEnd(t *testing.T) {
	arGVK := artifactregistryv1beta1.ArtifactRegistryRepositoryGVK
	unreadyKey := &unstructured.Unstructured{}
	unreadyKey.SetGroupVersionKind(kmsv1beta1.KMSCryptoKeyGVK)
	unreadyKey.SetNamespace("test-ns")
	unreadyKey.SetName("unready-key")

	tests := []struct {
		name      string
		spec      map[string]any
		extraObjs []runtime.Object
	}{
		{
			name: "missing_kms_crypto_key_ref",
			spec: map[string]any{
				"location": "us-central1",
				"format":   "DOCKER",
				"kmsKeyRef": map[string]any{
					"name": "deleted-key",
				},
			},
		},
		{
			name: "unready_kms_crypto_key_ref",
			spec: map[string]any{
				"location": "us-central1",
				"format":   "DOCKER",
				"kmsKeyRef": map[string]any{
					"name": "unready-key",
				},
			},
			extraObjs: []runtime.Object{unreadyKey},
		},
		{
			name: "missing_upstream_repository_ref",
			spec: map[string]any{
				"location": "us-central1",
				"format":   "DOCKER",
				"mode":     "VIRTUAL_REPOSITORY",
				"virtualRepositoryConfig": map[string]any{
					"upstreamPolicies": []any{
						map[string]any{
							"id":       "up-1",
							"priority": int64(1),
							"repositoryRef": map[string]any{
								"name": "deleted-upstream",
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			scheme := runtime.NewScheme()
			if err := corev1.AddToScheme(scheme); err != nil {
				t.Fatalf("corev1.AddToScheme() failed: %v", err)
			}
			if err := operatorv1beta1.AddToScheme(scheme); err != nil {
				t.Fatalf("operatorv1beta1.AddToScheme() failed: %v", err)
			}

			now := metav1.Now()
			u := &unstructured.Unstructured{
				Object: map[string]any{
					"spec": tc.spec,
				},
			}
			u.SetGroupVersionKind(arGVK)
			u.SetName("cmek-repo")
			u.SetNamespace("test-ns")
			u.SetDeletionTimestamp(&now)
			u.SetFinalizers([]string{k8s.ControllerFinalizerName})
			u.SetAnnotations(map[string]string{
				"cnrm.cloud.google.com/project-id": "test-proj",
			})

			objs := append([]runtime.Object{u}, tc.extraObjs...)
			kubeClient := fake.NewClientBuilder().
				WithScheme(scheme).
				WithStatusSubresource(u).
				WithRuntimeObjects(objs...).
				Build()
			mgr := &stubDirectManager{
				kubeClient: kubeClient,
				scheme:     scheme,
				recorder:   record.NewFakeRecorder(100),
			}
			model := &realNormalizeARModel{}
			r, err := directbase.NewReconciler(mgr, nil, nil, arGVK, model, directbase.Deps{
				JitterGenerator: &stubJitter{},
			})
			if err != nil {
				t.Fatalf("NewReconciler() error = %v", err)
			}
			req := reconcile.Request{
				NamespacedName: types.NamespacedName{Namespace: "test-ns", Name: "cmek-repo"},
			}

			_, err = r.Reconcile(t.Context(), req)

			wantID := "projects/test-proj/locations/us-central1/repositories/cmek-repo"
			if err != nil {
				t.Fatalf("Reconcile(%s) error = %v, want nil", tc.name, err)
			}
			if model.findCalls != 1 {
				t.Errorf("Reconcile(%s) Find calls = %d, want 1", tc.name, model.findCalls)
			}
			if model.deletedID != wantID {
				t.Errorf("Reconcile(%s) deletedID = %q, want %q", tc.name, model.deletedID, wantID)
			}
			gotObj := &unstructured.Unstructured{}
			gotObj.SetGroupVersionKind(arGVK)
			if getErr := kubeClient.Get(t.Context(), req.NamespacedName, gotObj); !apierrors.IsNotFound(getErr) {
				t.Errorf("Reconcile(%s) kubeClient.Get() error = %v, want NotFound", tc.name, getErr)
			}
		})
	}
}
