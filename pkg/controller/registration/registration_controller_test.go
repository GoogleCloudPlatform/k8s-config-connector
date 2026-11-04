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

package registration

import (
	"context"
	"net/http"
	"strings"
	"sync"
	"testing"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/jitter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/parent"
	dclconversion "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/metrics/server"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/structured-merge-diff/v4/fieldpath"

	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/compute"
)

func TestMaxConcurrentReconcilesOverrides(t *testing.T) {
	testGVK := schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeForwardingRule",
	}
	testGK := testGVK.GroupKind().String()
	otherGK := "BigQueryDataset.bigquery.cnrm.cloud.google.com"

	tests := []struct {
		name         string
		ccOverrides  map[string]int
		cccOverrides map[string]int
		want         int
	}{
		{
			name: "no overrides: returns default ControllerMaxConcurrentReconciles (20)",
			want: k8s.ControllerMaxConcurrentReconciles,
		},
		{
			name: "CC override only (cluster mode or global default): returns CC value",
			ccOverrides: map[string]int{
				testGK: 5,
			},
			want: 5,
		},
		{
			name: "CCC override only: returns CCC value",
			cccOverrides: map[string]int{
				testGK: 8,
			},
			want: 8,
		},
		{
			name: "CCC overrides another GK: falls back to global CC value for target GK",
			ccOverrides: map[string]int{
				testGK: 4,
			},
			cccOverrides: map[string]int{
				otherGK: 12,
			},
			want: 4,
		},
		{
			name: "both CCC and CC override same GK: CCC takes precedence",
			ccOverrides: map[string]int{
				testGK: 3,
			},
			cccOverrides: map[string]int{
				testGK: 15,
			},
			want: 15,
		},
		{
			name: "non-positive value in CCC falls back to CC override",
			ccOverrides: map[string]int{
				testGK: 7,
			},
			cccOverrides: map[string]int{
				testGK: 0,
			},
			want: 7,
		},
		{
			name: "non-positive values in both fall back to default (20)",
			ccOverrides: map[string]int{
				testGK: -1,
			},
			cccOverrides: map[string]int{
				testGK: 0,
			},
			want: k8s.ControllerMaxConcurrentReconciles,
		},
		{
			name: "unrecognized GroupKind is ignored",
			ccOverrides: map[string]int{
				"UnknownKind.unknown.cnrm.cloud.google.com": 50,
			},
			want: k8s.ControllerMaxConcurrentReconciles,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			opts := &defaultControllerOptions{
				maxConcurrentReconcilesOverrides: mergeMaxConcurrentReconcilesOverrides(ctx, tc.ccOverrides, tc.cccOverrides),
			}
			if got := opts.maxConcurrentReconcilesForGVK(testGVK); got != tc.want {
				t.Errorf("maxConcurrentReconcilesForGVK(%v) = %d, want %d", testGVK, got, tc.want)
			}
			if _, exists := opts.maxConcurrentReconcilesOverrides[schema.ParseGroupKind("UnknownKind.unknown.cnrm.cloud.google.com")]; exists {
				t.Errorf("expected unrecognized GroupKind to be ignored")
			}
		})
	}
}

type recordingSchemaUpdater struct {
	delegate    k8s.SchemaReferenceUpdater
	updateCount int
	lastCRD     *apiextensions.CustomResourceDefinition
}

func (u *recordingSchemaUpdater) UpdateSchema(crd *apiextensions.CustomResourceDefinition) error {
	u.updateCount++
	u.lastCRD = crd
	return u.delegate.UpdateSchema(crd)
}

var initRegistryOnce sync.Once

func newTestRegistrationReconciler(t *testing.T, initObjs ...client.Object) (*ReconcileRegistration, client.Client) {
	t.Helper()
	ctx := context.Background()
	initRegistryOnce.Do(func() {
		if err := registry.Init(ctx, &config.ControllerConfig{}); err != nil {
			t.Fatalf("error initializing direct registry: %v", err)
		}
	})

	scheme := runtime.NewScheme()
	if err := apiextensions.AddToScheme(scheme); err != nil {
		t.Fatalf("error adding apiextensions to scheme: %v", err)
	}
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).WithObjects(initObjs...).Build()

	mgr, err := manager.New(&rest.Config{Host: "http://localhost:8080"}, manager.Options{
		Scheme: scheme,
		Metrics: server.Options{
			BindAddress: "0",
		},
		HealthProbeBindAddress: "0",
		MapperProvider: func(c *rest.Config, httpClient *http.Client) (meta.RESTMapper, error) {
			return meta.NewDefaultRESTMapper(nil), nil
		},
		NewClient: func(config *rest.Config, options client.Options) (client.Client, error) {
			return fakeClient, nil
		},
	})
	if err != nil {
		t.Fatalf("error creating test manager: %v", err)
	}

	dclSchemaLoader, err := dclschemaloader.New()
	if err != nil {
		t.Fatalf("error creating DCL schema loader: %v", err)
	}
	dclConverter := dclconversion.New(dclSchemaLoader, dclmetadata.New())

	regOpts := &defaultControllerOptions{}
	r := &ReconcileRegistration{
		Client:             fakeClient,
		mgr:                mgr,
		dclConfig:          dcl.NewConfig(),
		dclConverter:       dclConverter,
		jitterGenerator:    &jitter.SimpleJitterGenerator{},
		controllers:        make(map[string]map[string]controllerContext),
		registrationFunc:   registerDefaultControllers(ctx, regOpts),
		reconcilers:        make(map[schema.GroupVersionKind]*parent.Reconcilers),
		SkipNameValidation: true,
	}
	return r, fakeClient
}

func buildTestCRD(name, group, version, kind string, specProperties map[string]apiextensions.JSONSchemaProps) *apiextensions.CustomResourceDefinition {
	return &apiextensions.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: apiextensions.CustomResourceDefinitionSpec{
			Group: group,
			Names: apiextensions.CustomResourceDefinitionNames{
				Kind: kind,
			},
			Versions: []apiextensions.CustomResourceDefinitionVersion{
				{
					Name:    version,
					Storage: true,
					Served:  true,
					Schema: &apiextensions.CustomResourceValidation{
						OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
							Type: "object",
							Properties: map[string]apiextensions.JSONSchemaProps{
								"spec": {
									Type:       "object",
									Properties: specProperties,
								},
							},
						},
					},
				},
			},
		},
	}
}

func TestReconcileRegistration_DynamicSchemaUpdate_ComputeRouterNAT(t *testing.T) {
	ctx := context.Background()
	crdName := "computerouternats.compute.cnrm.cloud.google.com"
	gvk := schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeRouterNAT",
	}

	// Initial CRD schema without the newly introduced 'type' and 'endpointTypes' fields.
	initialCRD := buildTestCRD(crdName, gvk.Group, gvk.Version, gvk.Kind, map[string]apiextensions.JSONSchemaProps{
		"region": {Type: "string"},
		"natIpAllocateOption": {
			Type: "string",
		},
	})

	r, fakeClient := newTestRegistrationReconciler(t, initialCRD)

	// 1. Initial registration reconcile.
	req := reconcile.Request{NamespacedName: types.NamespacedName{Name: crdName}}
	if _, err := r.Reconcile(ctx, req); err != nil {
		t.Fatalf("unexpected error on initial Reconcile: %v", err)
	}

	ctrlCtx := r.controllers[gvk.Group][gvk.Kind]
	if !ctrlCtx.registered {
		t.Fatalf("expected controller for %v to be registered", gvk)
	}
	if ctrlCtx.schemaUpdater == nil {
		t.Fatalf("expected non-nil schemaUpdater for %v (Terraform+Direct controller), got nil", gvk)
	}
	if ctrlCtx.schemaUpdater != r.reconcilers[gvk].TF {
		t.Fatalf("expected schemaUpdater to be wired to reconcilers.TF")
	}

	recorder := &recordingSchemaUpdater{delegate: ctrlCtx.schemaUpdater}
	ctrlCtx.schemaUpdater = recorder
	r.controllers[gvk.Group][gvk.Kind] = ctrlCtx

	// Verify that with the old schema, overlaying managed fields onto state containing '.type' and '.endpointTypes'
	// fails with "field not declared in schema" (reproducing b/535980652).
	k8sManagedFields := fieldpath.NewSet(
		fieldpath.MakePathOrDie("f:spec", "f:region"),
	)
	stateAsKRM := map[string]interface{}{
		"region":              "asia-southeast1",
		"natIpAllocateOption": "MANUAL_ONLY",
		"type":                "PUBLIC",
		"endpointTypes":       []interface{}{"ENDPOINT_TYPE_VM"},
	}
	spec := map[string]interface{}{
		"region": "asia-southeast1",
	}

	oldSchema := k8s.GetOpenAPIV3SchemaFromCRD(initialCRD)
	if _, err := k8s.OverlayManagedFieldsOntoState(spec, stateAsKRM, k8sManagedFields, oldSchema, nil); err == nil || !strings.Contains(err.Error(), "field not declared in schema") {
		t.Fatalf("expected 'field not declared in schema' error with old CRD schema, got: %v", err)
	}

	// 2. Update the CRD in the API server to add 'type' and 'endpointTypes' to openAPIV3Schema.
	updatedCRD := buildTestCRD(crdName, gvk.Group, gvk.Version, gvk.Kind, map[string]apiextensions.JSONSchemaProps{
		"region":              {Type: "string"},
		"natIpAllocateOption": {Type: "string"},
		"type":                {Type: "string"},
		"endpointTypes": {
			Type: "array",
			Items: &apiextensions.JSONSchemaPropsOrArray{
				Schema: &apiextensions.JSONSchemaProps{Type: "string"},
			},
		},
	})
	updatedCRD.ResourceVersion = initialCRD.ResourceVersion
	if err := fakeClient.Update(ctx, updatedCRD); err != nil {
		t.Fatalf("error updating CRD in fake client: %v", err)
	}

	// 3. Reconcile the CRD update event through ReconcileRegistration.
	if _, err := r.Reconcile(ctx, req); err != nil {
		t.Fatalf("unexpected error on CRD update Reconcile: %v", err)
	}

	if recorder.updateCount != 1 {
		t.Fatalf("expected schemaUpdater.UpdateSchema to be called once on CRD update, got %d", recorder.updateCount)
	}

	// Verify that the updated schema received by the TF reconciler allows OverlayManagedFieldsOntoState to succeed.
	newSchema := k8s.GetOpenAPIV3SchemaFromCRD(recorder.lastCRD)
	mergedSpecs, err := k8s.OverlayManagedFieldsOntoState(spec, stateAsKRM, k8sManagedFields, newSchema, nil)
	if err != nil {
		t.Fatalf("expected OverlayManagedFieldsOntoState to succeed after CRD schema update, got: %v", err)
	}
	if got := mergedSpecs["type"]; got != "PUBLIC" {
		t.Errorf("expected mergedSpecs[\"type\"] = \"PUBLIC\", got %v", got)
	}
}

func TestReconcileRegistration_SchemaUpdaterByControllerType(t *testing.T) {
	ctx := context.Background()

	t.Run("DCL controller sets schemaUpdater to reconcilers.DCL", func(t *testing.T) {
		crdName := "computefirewallpolicies.compute.cnrm.cloud.google.com"
		gvk := schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ComputeFirewallPolicy",
		}
		crd := buildTestCRD(crdName, gvk.Group, gvk.Version, gvk.Kind, map[string]apiextensions.JSONSchemaProps{
			"shortName": {Type: "string"},
		})

		r, fakeClient := newTestRegistrationReconciler(t, crd)
		req := reconcile.Request{NamespacedName: types.NamespacedName{Name: crdName}}
		if _, err := r.Reconcile(ctx, req); err != nil {
			t.Fatalf("unexpected error on initial Reconcile: %v", err)
		}

		ctrlCtx := r.controllers[gvk.Group][gvk.Kind]
		if !ctrlCtx.registered {
			t.Fatalf("expected controller for %v to be registered", gvk)
		}
		if ctrlCtx.schemaUpdater == nil {
			t.Fatalf("expected non-nil schemaUpdater for DCL controller %v, got nil", gvk)
		}
		if ctrlCtx.schemaUpdater != r.reconcilers[gvk].DCL {
			t.Fatalf("expected schemaUpdater to be wired to reconcilers.DCL")
		}

		recorder := &recordingSchemaUpdater{delegate: ctrlCtx.schemaUpdater}
		ctrlCtx.schemaUpdater = recorder
		r.controllers[gvk.Group][gvk.Kind] = ctrlCtx

		// Update the CRD and verify second Reconcile updates schema without error.
		updatedCRD := buildTestCRD(crdName, gvk.Group, gvk.Version, gvk.Kind, map[string]apiextensions.JSONSchemaProps{
			"shortName":   {Type: "string"},
			"description": {Type: "string"},
		})
		updatedCRD.ResourceVersion = crd.ResourceVersion
		if err := fakeClient.Update(ctx, updatedCRD); err != nil {
			t.Fatalf("error updating CRD: %v", err)
		}
		if _, err := r.Reconcile(ctx, req); err != nil {
			t.Fatalf("unexpected error on CRD update Reconcile: %v", err)
		}
		if recorder.updateCount != 1 {
			t.Fatalf("expected schemaUpdater.UpdateSchema to be called once on CRD update, got %d", recorder.updateCount)
		}
	})

	t.Run("Direct-only controller leaves schemaUpdater nil", func(t *testing.T) {
		crdName := "computefirewallpolicyrules.compute.cnrm.cloud.google.com"
		gvk := schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ComputeFirewallPolicyRule",
		}
		crd := buildTestCRD(crdName, gvk.Group, gvk.Version, gvk.Kind, map[string]apiextensions.JSONSchemaProps{
			"action": {Type: "string"},
		})

		r, _ := newTestRegistrationReconciler(t, crd)
		req := reconcile.Request{NamespacedName: types.NamespacedName{Name: crdName}}
		if _, err := r.Reconcile(ctx, req); err != nil {
			t.Fatalf("unexpected error on initial Reconcile: %v", err)
		}

		ctrlCtx := r.controllers[gvk.Group][gvk.Kind]
		if !ctrlCtx.registered {
			t.Fatalf("expected controller for %v to be registered", gvk)
		}
		if ctrlCtx.schemaUpdater != nil {
			t.Fatalf("expected nil schemaUpdater for Direct-only controller %v, got %v", gvk, ctrlCtx.schemaUpdater)
		}
	})
}
