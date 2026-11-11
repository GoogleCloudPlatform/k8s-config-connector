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
	"sync"
	"testing"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/jitter"
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
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/metrics/server"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

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

func TestReconcileRegistration_DynamicSchemaUpdate(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name                   string
		crdName                string
		gvk                    schema.GroupVersionKind
		initialSpecProps       map[string]apiextensions.JSONSchemaProps
		updatedSpecProps       map[string]apiextensions.JSONSchemaProps
		expectHasSchemaUpdater bool
	}{
		{
			name:    "Terraform controller registers schemaUpdater and handles CRD schema update",
			crdName: "computerouternats.compute.cnrm.cloud.google.com",
			gvk: schema.GroupVersionKind{
				Group:   "compute.cnrm.cloud.google.com",
				Version: "v1beta1",
				Kind:    "ComputeRouterNAT",
			},
			initialSpecProps: map[string]apiextensions.JSONSchemaProps{
				"region":              {Type: "string"},
				"natIpAllocateOption": {Type: "string"},
			},
			updatedSpecProps: map[string]apiextensions.JSONSchemaProps{
				"region":              {Type: "string"},
				"natIpAllocateOption": {Type: "string"},
				"type":                {Type: "string"},
			},
			expectHasSchemaUpdater: true,
		},
		{
			name:    "DCL controller registers schemaUpdater and handles CRD schema update",
			crdName: "computefirewallpolicies.compute.cnrm.cloud.google.com",
			gvk: schema.GroupVersionKind{
				Group:   "compute.cnrm.cloud.google.com",
				Version: "v1beta1",
				Kind:    "ComputeFirewallPolicy",
			},
			initialSpecProps: map[string]apiextensions.JSONSchemaProps{
				"shortName": {Type: "string"},
			},
			updatedSpecProps: map[string]apiextensions.JSONSchemaProps{
				"shortName":   {Type: "string"},
				"description": {Type: "string"},
			},
			expectHasSchemaUpdater: true,
		},
		{
			name:    "Direct-only controller leaves schemaUpdater nil",
			crdName: "computefirewallpolicyrules.compute.cnrm.cloud.google.com",
			gvk: schema.GroupVersionKind{
				Group:   "compute.cnrm.cloud.google.com",
				Version: "v1beta1",
				Kind:    "ComputeFirewallPolicyRule",
			},
			initialSpecProps: map[string]apiextensions.JSONSchemaProps{
				"action": {Type: "string"},
			},
			updatedSpecProps: map[string]apiextensions.JSONSchemaProps{
				"action":      {Type: "string"},
				"description": {Type: "string"},
			},
			expectHasSchemaUpdater: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			initialCRD := buildTestCRD(tc.crdName, tc.gvk.Group, tc.gvk.Version, tc.gvk.Kind, tc.initialSpecProps)
			r, fakeClient := newTestRegistrationReconciler(t, initialCRD)
			req := reconcile.Request{NamespacedName: types.NamespacedName{Name: tc.crdName}}

			// Act: initial registration
			res, err := r.Reconcile(ctx, req)

			// Assert: controller is registered with expected schema updater presence
			if err != nil {
				t.Fatalf("unexpected error on initial Reconcile: %v", err)
			}
			if res.Requeue {
				t.Fatalf("expected Requeue=false on initial Reconcile, got true")
			}
			registered, hasSchemaUpdater := r.controllerContextFor(tc.gvk)
			if !registered {
				t.Fatalf("expected controller for %v to be registered", tc.gvk)
			}
			if hasSchemaUpdater != tc.expectHasSchemaUpdater {
				t.Fatalf("hasSchemaUpdater for %v = %v, want %v", tc.gvk, hasSchemaUpdater, tc.expectHasSchemaUpdater)
			}

			// Arrange: update CRD schema in the API server
			var currentCRD apiextensions.CustomResourceDefinition
			if err := fakeClient.Get(ctx, req.NamespacedName, &currentCRD); err != nil {
				t.Fatalf("error fetching current CRD: %v", err)
			}
			updatedCRD := buildTestCRD(tc.crdName, tc.gvk.Group, tc.gvk.Version, tc.gvk.Kind, tc.updatedSpecProps)
			updatedCRD.ResourceVersion = currentCRD.ResourceVersion
			if err := fakeClient.Update(ctx, updatedCRD); err != nil {
				t.Fatalf("error updating CRD in fake client: %v", err)
			}

			// Act: reconcile CRD schema update event
			updateRes, updateErr := r.Reconcile(ctx, req)

			// Assert: update event reconciled successfully
			if updateErr != nil {
				t.Fatalf("unexpected error reconciling CRD schema update: %v", updateErr)
			}
			if updateRes.Requeue {
				t.Fatalf("expected Requeue=false on CRD update Reconcile, got true")
			}
		})
	}
}

var (
	initRegistryOnce sync.Once
	initRegistryErr  error
)

func newTestRegistrationReconciler(t *testing.T, initObjs ...client.Object) (*ReconcileRegistration, client.Client) {
	t.Helper()
	ctx := context.Background()
	initRegistryOnce.Do(func() {
		initRegistryErr = registry.Init(ctx, &config.ControllerConfig{})
	})
	if initRegistryErr != nil {
		t.Fatalf("error initializing direct registry: %v", initRegistryErr)
	}

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
	deps := &controller.Deps{
		DCLConfig:          dcl.NewConfig(),
		DCLConverter:       dclConverter,
		JitterGen:          &jitter.SimpleJitterGenerator{},
		SkipNameValidation: true,
	}
	return newReconcileRegistration(mgr, deps, registerDefaultControllers(ctx, regOpts)), fakeClient
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

func TestUsePriorityQueueOptionPlumbing(t *testing.T) {
	tests := []struct {
		name                 string
		ccUsePriorityQueue   *bool
		cccUsePriorityQueue  *bool
		wantUsePriorityQueue bool
	}{
		{
			name:                 "both unset -> false",
			wantUsePriorityQueue: false,
		},
		{
			name:                 "cc true -> true",
			ccUsePriorityQueue:   ptr.To(true),
			wantUsePriorityQueue: true,
		},
		{
			name:                 "cc false -> false",
			ccUsePriorityQueue:   ptr.To(false),
			wantUsePriorityQueue: false,
		},
		{
			name:                 "cc true but ccc false -> false (ccc override)",
			ccUsePriorityQueue:   ptr.To(true),
			cccUsePriorityQueue:  ptr.To(false),
			wantUsePriorityQueue: false,
		},
		{
			name:                 "cc false but ccc true -> true (ccc override)",
			ccUsePriorityQueue:   ptr.To(false),
			cccUsePriorityQueue:  ptr.To(true),
			wantUsePriorityQueue: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			usePriorityQueue := false
			if tc.ccUsePriorityQueue != nil && *tc.ccUsePriorityQueue {
				usePriorityQueue = true
			}
			if tc.cccUsePriorityQueue != nil {
				usePriorityQueue = *tc.cccUsePriorityQueue
			}
			if usePriorityQueue != tc.wantUsePriorityQueue {
				t.Errorf("got %v, want %v", usePriorityQueue, tc.wantUsePriorityQueue)
			}
		})
	}
}
