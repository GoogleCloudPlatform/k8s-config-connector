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

package webhook

import (
	"context"
	"errors"
	"testing"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"github.com/google/go-cmp/cmp"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

// newKCCStateReader returns a fake reader that holds the given ConfigConnector and
// ConfigConnectorContext objects (either may be nil).
func newKCCStateReader(t *testing.T, cc *operatorv1beta1.ConfigConnector, ccc *operatorv1beta1.ConfigConnectorContext) client.Reader {
	t.Helper()
	scheme := runtime.NewScheme()
	if err := operatorv1beta1.AddToScheme(scheme); err != nil {
		t.Fatalf("adding operator types to scheme: %v", err)
	}
	builder := fake.NewClientBuilder().WithScheme(scheme)
	if cc != nil {
		builder = builder.WithObjects(cc)
	}
	if ccc != nil {
		builder = builder.WithObjects(ccc)
	}
	return builder.Build()
}

func newConfigConnector(mode string) *operatorv1beta1.ConfigConnector {
	return &operatorv1beta1.ConfigConnector{
		ObjectMeta: metav1.ObjectMeta{Name: operatorv1beta1.ConfigConnectorAllowedName},
		Spec:       operatorv1beta1.ConfigConnectorSpec{Mode: mode},
	}
}

const testNamespace = "test-namespace"

func newConfigConnectorContext(overrides map[string]k8s.ReconcilerType) *operatorv1beta1.ConfigConnectorContext {
	return &operatorv1beta1.ConfigConnectorContext{
		ObjectMeta: metav1.ObjectMeta{Namespace: testNamespace, Name: operatorv1beta1.ConfigConnectorContextAllowedName},
		Spec: operatorv1beta1.ConfigConnectorContextSpec{
			GoogleServiceAccount: "kcc@test-project.iam.gserviceaccount.com",
			Experiments:          &operatorv1beta1.Experiments{ControllerOverrides: overrides},
		},
	}
}

var (
	configConnectorKind        = operatorv1beta1.ConfigConnectorGroupVersionKind.Kind
	configConnectorContextKind = operatorv1beta1.ConfigConnectorContextGroupVersionKind.Kind
	namespacedConfigConnector  = map[string]any{"spec": map[string]any{"mode": "namespaced"}}
)

// stubKCCStateReader returns canned ConfigConnector and ConfigConnectorContext content, or an
// error, by kind. It produces responses that the fake client can't, such as RBAC errors, missing
// CRDs and malformed objects.
type stubKCCStateReader struct {
	client.Reader
	objects map[string]map[string]any
	errs    map[string]error
}

func (r stubKCCStateReader) Get(_ context.Context, key client.ObjectKey, obj client.Object, _ ...client.GetOption) error {
	kind := obj.GetObjectKind().GroupVersionKind().Kind
	if err := r.errs[kind]; err != nil {
		return err
	}
	content, ok := r.objects[kind]
	if !ok {
		return apierrors.NewNotFound(schema.GroupResource{Group: "core.cnrm.cloud.google.com", Resource: kind}, key.Name)
	}
	obj.(*unstructured.Unstructured).Object = content
	return nil
}

func TestControllerOverrides(t *testing.T) {
	overrides := map[string]k8s.ReconcilerType{"ArtifactRegistryRepository.artifactregistry.cnrm.cloud.google.com": k8s.ReconcilerTypeDirect}
	crdNotInstalled := &meta.NoKindMatchError{GroupKind: schema.GroupKind{Group: "core.cnrm.cloud.google.com", Kind: "ConfigConnectorContext"}}
	tests := []struct {
		name   string
		reader client.Reader
		want   map[string]k8s.ReconcilerType
	}{
		{
			name:   "no reader",
			reader: nil,
			want:   nil,
		},
		{
			name:   "no ConfigConnector",
			reader: newKCCStateReader(t, nil, nil),
			want:   nil,
		},
		{
			name:   "ConfigConnector CRD not installed",
			reader: stubKCCStateReader{errs: map[string]error{configConnectorKind: crdNotInstalled}},
			want:   nil,
		},
		{
			name:   "cluster mode ignores ConfigConnectorContext",
			reader: newKCCStateReader(t, newConfigConnector("cluster"), newConfigConnectorContext(overrides)),
			want:   nil,
		},
		{
			name:   "namespaced mode without ConfigConnectorContext",
			reader: newKCCStateReader(t, newConfigConnector("namespaced"), nil),
			want:   nil,
		},
		{
			name: "ConfigConnectorContext CRD not installed",
			reader: stubKCCStateReader{
				objects: map[string]map[string]any{configConnectorKind: namespacedConfigConnector},
				errs:    map[string]error{configConnectorContextKind: crdNotInstalled},
			},
			want: nil,
		},
		{
			name:   "namespaced mode with overrides",
			reader: newKCCStateReader(t, newConfigConnector("namespaced"), newConfigConnectorContext(overrides)),
			want:   overrides,
		},
		{
			name:   "empty mode means namespaced mode",
			reader: newKCCStateReader(t, newConfigConnector(""), newConfigConnectorContext(overrides)),
			want:   overrides,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := controllerOverrides(t.Context(), tc.reader, testNamespace)

			if err != nil {
				t.Fatalf("controllerOverrides(%q) returned error: %v", testNamespace, err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("controllerOverrides(%q) mismatch (-want +got):\n%s", testNamespace, diff)
			}
		})
	}
}

func TestControllerOverridesErrors(t *testing.T) {
	forbidden := apierrors.NewForbidden(schema.GroupResource{Group: "core.cnrm.cloud.google.com", Resource: "configconnectorcontexts"}, "", errors.New("RBAC: access denied"))
	tests := []struct {
		name    string
		reader  client.Reader
		wantErr error
	}{
		{
			name:    "ConfigConnector read fails",
			reader:  stubKCCStateReader{errs: map[string]error{configConnectorKind: forbidden}},
			wantErr: forbidden,
		},
		{
			name: "ConfigConnectorContext read fails",
			reader: stubKCCStateReader{
				objects: map[string]map[string]any{configConnectorKind: namespacedConfigConnector},
				errs:    map[string]error{configConnectorContextKind: forbidden},
			},
			wantErr: forbidden,
		},
		{
			name: "mode is not a string",
			reader: stubKCCStateReader{objects: map[string]map[string]any{
				configConnectorKind: {"spec": map[string]any{"mode": int64(1)}},
			}},
		},
		{
			name: "controllerOverrides is not a map of strings",
			reader: stubKCCStateReader{objects: map[string]map[string]any{
				configConnectorKind: namespacedConfigConnector,
				configConnectorContextKind: {"spec": map[string]any{"experiments": map[string]any{
					"controllerOverrides": map[string]any{"ArtifactRegistryRepository.artifactregistry.cnrm.cloud.google.com": int64(1)},
				}}},
			}},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := controllerOverrides(t.Context(), tc.reader, testNamespace)

			if err == nil {
				t.Fatalf("controllerOverrides(%q) returned nil error, want an error", testNamespace)
			}
			if tc.wantErr != nil && !errors.Is(err, tc.wantErr) {
				t.Errorf("controllerOverrides(%q) error = %v, want %v", testNamespace, err, tc.wantErr)
			}
		})
	}
}

func newTestObject(apiVersion, kind string, annotations map[string]string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion(apiVersion)
	u.SetKind(kind)
	u.SetNamespace(testNamespace)
	u.SetName("test-object")
	u.SetAnnotations(annotations)
	return u
}

func TestReconcilerTypeForObject(t *testing.T) {
	directAnnotation := map[string]string{k8s.AlphaReconcilerAnnotation: "direct"}
	unreadableKCCState := stubKCCStateReader{errs: map[string]error{
		configConnectorKind: apierrors.NewForbidden(schema.GroupResource{Group: "core.cnrm.cloud.google.com", Resource: "configconnectors"}, "", errors.New("RBAC: access denied")),
	}}
	tests := []struct {
		name   string
		obj    *unstructured.Unstructured
		reader client.Reader
		want   k8s.ReconcilerType
		wantOK bool
	}{
		{
			name:   "Direct by default",
			obj:    newTestObject("documentai.cnrm.cloud.google.com/v1alpha1", "DocumentAIProcessorDefaultVersion", nil),
			reader: newKCCStateReader(t, newConfigConnector("namespaced"), newConfigConnectorContext(nil)),
			want:   k8s.ReconcilerTypeDirect,
			wantOK: true,
		},
		{
			name: "namespace override back to Terraform",
			obj:  newTestObject("documentai.cnrm.cloud.google.com/v1alpha1", "DocumentAIProcessorDefaultVersion", nil),
			reader: newKCCStateReader(t, newConfigConnector("namespaced"), newConfigConnectorContext(map[string]k8s.ReconcilerType{
				"DocumentAIProcessorDefaultVersion.documentai.cnrm.cloud.google.com": k8s.ReconcilerTypeTerraform,
			})),
			want:   k8s.ReconcilerTypeTerraform,
			wantOK: true,
		},
		{
			name:   "annotation opts in to Direct without reading the namespace",
			obj:    newTestObject("artifactregistry.cnrm.cloud.google.com/v1beta1", "ArtifactRegistryRepository", directAnnotation),
			reader: unreadableKCCState,
			want:   k8s.ReconcilerTypeDirect,
			wantOK: true,
		},
		{
			name: "unsupported override falls back to the default controller",
			obj:  newTestObject("artifactregistry.cnrm.cloud.google.com/v1beta1", "ArtifactRegistryRepository", nil),
			reader: newKCCStateReader(t, newConfigConnector("namespaced"), newConfigConnectorContext(map[string]k8s.ReconcilerType{
				"ArtifactRegistryRepository.artifactregistry.cnrm.cloud.google.com": k8s.ReconcilerTypeDCL,
			})),
			want:   k8s.ReconcilerTypeTerraform,
			wantOK: true,
		},
		{
			name:   "Direct annotation on a kind without a Direct controller falls back",
			obj:    newTestObject("compute.cnrm.cloud.google.com/v1beta1", "ComputeBackendService", directAnnotation),
			reader: nil,
			want:   k8s.ReconcilerTypeTerraform,
			wantOK: true,
		},
		{
			name:   "undecided when the namespace can't be read",
			obj:    newTestObject("documentai.cnrm.cloud.google.com/v1alpha1", "DocumentAIProcessorDefaultVersion", nil),
			reader: unreadableKCCState,
			wantOK: false,
		},
		{
			name:   "undecided for an unknown kind",
			obj:    newTestObject("test.cnrm.cloud.google.com/v1alpha1", "UnknownKind", directAnnotation),
			reader: nil,
			wantOK: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := reconcilerTypeForObject(t.Context(), tc.reader, tc.obj)

			if ok != tc.wantOK {
				t.Fatalf("reconcilerTypeForObject(%v) ok = %v, want %v", tc.obj.GetKind(), ok, tc.wantOK)
			}
			if got != tc.want {
				t.Errorf("reconcilerTypeForObject(%v) = %q, want %q", tc.obj.GetKind(), got, tc.want)
			}
		})
	}
}
