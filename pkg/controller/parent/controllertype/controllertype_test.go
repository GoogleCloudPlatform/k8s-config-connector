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

package controllertype

import (
	"context"
	"errors"
	"testing"

	kccpredicate "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/predicate"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	// Terraform by default, Direct supported.
	artifactRegistryRepositoryGVK = schema.GroupVersionKind{Group: "artifactregistry.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ArtifactRegistryRepository"}
	// Direct by default, Terraform supported.
	documentAIProcessorDefaultVersionGVK = schema.GroupVersionKind{Group: "documentai.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DocumentAIProcessorDefaultVersion"}
	// Terraform by default; the special cases route some objects to Direct based on their spec.
	bigQueryTableGVK         = schema.GroupVersionKind{Group: "bigquery.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BigQueryTable"}
	computeForwardingRuleGVK = schema.GroupVersionKind{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeForwardingRule"}
	secretManagerSecretGVK   = schema.GroupVersionKind{Group: "secretmanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SecretManagerSecret"}
	spannerInstanceGVK       = schema.GroupVersionKind{Group: "spanner.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SpannerInstance"}
)

func newObject(gvk schema.GroupVersionKind, annotations map[string]string, spec map[string]any) *unstructured.Unstructured {
	u := &unstructured.Unstructured{Object: map[string]any{"spec": spec}}
	u.SetGroupVersionKind(gvk)
	u.SetNamespace("test-namespace")
	u.SetName("test-object")
	u.SetAnnotations(annotations)
	return u
}

func staticOverrides(overrides map[string]k8s.ReconcilerType) OverridesFunc {
	return func(ctx context.Context) (map[string]k8s.ReconcilerType, error) {
		return overrides, nil
	}
}

func TestForObject(t *testing.T) {
	labels := map[string]any{"team": "a"}
	tests := []struct {
		name      string
		gvk       schema.GroupVersionKind
		annots    map[string]string
		spec      map[string]any
		overrides OverridesFunc
		want      k8s.ReconcilerType
	}{
		{
			name: "default controller",
			gvk:  artifactRegistryRepositoryGVK,
			want: k8s.ReconcilerTypeTerraform,
		},
		{
			name: "default controller of a Direct-by-default kind",
			gvk:  documentAIProcessorDefaultVersionGVK,
			want: k8s.ReconcilerTypeDirect,
		},
		{
			name:   "reconciler annotation selects Direct",
			gvk:    artifactRegistryRepositoryGVK,
			annots: map[string]string{k8s.AlphaReconcilerAnnotation: "direct"},
			want:   k8s.ReconcilerTypeDirect,
		},
		{
			name:      "namespace override selects Direct",
			gvk:       artifactRegistryRepositoryGVK,
			overrides: staticOverrides(map[string]k8s.ReconcilerType{"ArtifactRegistryRepository.artifactregistry.cnrm.cloud.google.com": k8s.ReconcilerTypeDirect}),
			want:      k8s.ReconcilerTypeDirect,
		},
		{
			name:      "namespace override selects Terraform for a Direct-by-default kind",
			gvk:       documentAIProcessorDefaultVersionGVK,
			overrides: staticOverrides(map[string]k8s.ReconcilerType{"DocumentAIProcessorDefaultVersion.documentai.cnrm.cloud.google.com": k8s.ReconcilerTypeTerraform}),
			want:      k8s.ReconcilerTypeTerraform,
		},
		{
			name:      "override for another kind is ignored",
			gvk:       artifactRegistryRepositoryGVK,
			overrides: staticOverrides(map[string]k8s.ReconcilerType{"PubSubTopic.pubsub.cnrm.cloud.google.com": k8s.ReconcilerTypeDirect}),
			want:      k8s.ReconcilerTypeTerraform,
		},
		{
			name:   "BigQueryTable with the unmanaged annotation",
			gvk:    bigQueryTableGVK,
			annots: map[string]string{kccpredicate.AnnotationUnmanaged: "true"},
			want:   k8s.ReconcilerTypeDirect,
		},
		{
			name: "BigQueryTable with labels",
			gvk:  bigQueryTableGVK,
			spec: map[string]any{"labels": labels},
			want: k8s.ReconcilerTypeDirect,
		},
		{
			name: "ComputeForwardingRule that targets a Google APIs bundle",
			gvk:  computeForwardingRuleGVK,
			spec: map[string]any{"target": map[string]any{"googleAPIsBundle": "all-apis"}},
			want: k8s.ReconcilerTypeDirect,
		},
		{
			name: "SecretManagerSecret with labels",
			gvk:  secretManagerSecretGVK,
			spec: map[string]any{"labels": labels},
			want: k8s.ReconcilerTypeDirect,
		},
		{
			name: "SpannerInstance with an edition",
			gvk:  spannerInstanceGVK,
			spec: map[string]any{"edition": "ENTERPRISE"},
			want: k8s.ReconcilerTypeDirect,
		},
		{
			name:      "special-case kind without its trigger field uses the namespace override",
			gvk:       secretManagerSecretGVK,
			overrides: staticOverrides(map[string]k8s.ReconcilerType{"SecretManagerSecret.secretmanager.cnrm.cloud.google.com": k8s.ReconcilerTypeDirect}),
			want:      k8s.ReconcilerTypeDirect,
		},
		{
			name: "special-case kind without its trigger field uses the default controller",
			gvk:  secretManagerSecretGVK,
			want: k8s.ReconcilerTypeTerraform,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			u := newObject(tc.gvk, tc.annots, tc.spec)

			got, err := ForObject(t.Context(), tc.gvk, u, tc.overrides)

			if err != nil {
				t.Fatalf("ForObject(%v) returned error: %v", tc.gvk.Kind, err)
			}
			if got != tc.want {
				t.Errorf("ForObject(%v) = %q, want %q", tc.gvk.Kind, got, tc.want)
			}
		})
	}
}

func TestForObjectSkipsOverridesWhenAnEarlierRuleDecides(t *testing.T) {
	tests := []struct {
		name   string
		gvk    schema.GroupVersionKind
		annots map[string]string
		spec   map[string]any
	}{
		{
			name:   "reconciler annotation",
			gvk:    artifactRegistryRepositoryGVK,
			annots: map[string]string{k8s.AlphaReconcilerAnnotation: "direct"},
		},
		{
			name: "special case",
			gvk:  secretManagerSecretGVK,
			spec: map[string]any{"labels": map[string]any{"team": "a"}},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			called := false
			overrides := func(ctx context.Context) (map[string]k8s.ReconcilerType, error) {
				called = true
				return nil, errors.New("overrides must not be read")
			}
			u := newObject(tc.gvk, tc.annots, tc.spec)

			got, err := ForObject(t.Context(), tc.gvk, u, overrides)

			if err != nil {
				t.Fatalf("ForObject(%v) returned error: %v", tc.gvk.Kind, err)
			}
			if got != k8s.ReconcilerTypeDirect {
				t.Errorf("ForObject(%v) = %q, want %q", tc.gvk.Kind, got, k8s.ReconcilerTypeDirect)
			}
			if called {
				t.Errorf("ForObject(%v) read the namespace overrides, want them skipped", tc.gvk.Kind)
			}
		})
	}
}

func TestForObjectReturnsOverridesError(t *testing.T) {
	wantErr := errors.New("cannot read ConfigConnectorContext")
	overrides := func(ctx context.Context) (map[string]k8s.ReconcilerType, error) {
		return nil, wantErr
	}
	u := newObject(artifactRegistryRepositoryGVK, nil, nil)

	_, err := ForObject(t.Context(), artifactRegistryRepositoryGVK, u, overrides)

	if !errors.Is(err, wantErr) {
		t.Errorf("ForObject(%v) error = %v, want %v", artifactRegistryRepositoryGVK.Kind, err, wantErr)
	}
}

func TestForObjectReturnsErrorForUnknownKind(t *testing.T) {
	gvk := schema.GroupVersionKind{Group: "test.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "UnknownKind"}
	u := newObject(gvk, nil, nil)

	_, err := ForObject(t.Context(), gvk, u, nil)

	if err == nil {
		t.Errorf("ForObject(%v) returned nil error, want an error", gvk.Kind)
	}
}
