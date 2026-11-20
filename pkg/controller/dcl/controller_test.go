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

package dclcontroller

import (
	"context"
	"reflect"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/lifecyclehandler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"github.com/google/go-cmp/cmp"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestApplyChangesForBackwardsCompatibility_NetworkConnectivityHub(t *testing.T) {
	tests := []struct {
		name                 string
		objectAnnotations    map[string]string
		namespaceAnnotations map[string]string
		spec                 map[string]interface{}
		expectedProjectRef   map[string]interface{}
	}{
		{
			name: "respects object-level project-id annotation over namespace name",
			objectAnnotations: map[string]string{
				k8s.ProjectIDAnnotation: "my-custom-project",
			},
			namespaceAnnotations: nil,
			spec: map[string]interface{}{
				"description": "test hub",
			},
			expectedProjectRef: map[string]interface{}{
				"external": "my-custom-project",
			},
		},
		{
			name: "respects object-level project-id annotation over namespace annotation",
			objectAnnotations: map[string]string{
				k8s.ProjectIDAnnotation: "my-custom-project",
			},
			namespaceAnnotations: map[string]string{
				k8s.ProjectIDAnnotation: "ns-project",
			},
			spec: map[string]interface{}{
				"description": "test hub",
			},
			expectedProjectRef: map[string]interface{}{
				"external": "my-custom-project",
			},
		},
		{
			name:              "falls back to namespace annotation when object annotation is absent",
			objectAnnotations: nil,
			namespaceAnnotations: map[string]string{
				k8s.ProjectIDAnnotation: "ns-project",
			},
			spec: map[string]interface{}{
				"description": "test hub",
			},
			expectedProjectRef: map[string]interface{}{
				"external": "ns-project",
			},
		},
		{
			name:                 "falls back to namespace name when both object and namespace annotations are absent",
			objectAnnotations:    nil,
			namespaceAnnotations: nil,
			spec: map[string]interface{}{
				"description": "test hub",
			},
			expectedProjectRef: map[string]interface{}{
				"external": "test-ns",
			},
		},
		{
			name: "preserves explicit spec.projectRef even when object annotation is present",
			objectAnnotations: map[string]string{
				k8s.ProjectIDAnnotation: "ignored-annotation-project",
			},
			namespaceAnnotations: map[string]string{
				k8s.ProjectIDAnnotation: "ignored-ns-project",
			},
			spec: map[string]interface{}{
				"description": "test hub",
				"projectRef": map[string]interface{}{
					"external": "explicit-project",
				},
			},
			expectedProjectRef: map[string]interface{}{
				"external": "explicit-project",
			},
		},
	}

	smLoader := dclmetadata.New()
	schemaLoader, err := dclschemaloader.New()
	if err != nil {
		t.Fatalf("unexpected error creating schema loader: %v", err)
	}

	hubGVK := schema.GroupVersionKind{
		Group:   "networkconnectivity.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "NetworkConnectivityHub",
	}

	dclSchema, err := dclschemaloader.GetDCLSchemaForGVK(hubGVK, smLoader, schemaLoader)
	if err != nil {
		t.Fatalf("unexpected error getting DCL schema for NetworkConnectivityHub: %v", err)
	}

	converter := conversion.New(schemaLoader, smLoader)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ns := &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:        "test-ns",
					Annotations: tc.namespaceAnnotations,
				},
			}

			s := runtime.NewScheme()
			if err := corev1.AddToScheme(s); err != nil {
				t.Fatalf("failed to add corev1 to scheme: %v", err)
			}
			client := fake.NewClientBuilder().WithScheme(s).WithObjects(ns).Build()

			r := &Reconciler{
				LifecycleHandler: lifecyclehandler.NewLifecycleHandler(client, nil),
				converter:        converter,
				schema:           dclSchema,
			}

			u := &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "networkconnectivity.cnrm.cloud.google.com/v1beta1",
					"kind":       "NetworkConnectivityHub",
					"metadata": map[string]interface{}{
						"name":      "test-hub",
						"namespace": "test-ns",
					},
					"spec": tc.spec,
				},
			}
			if tc.objectAnnotations != nil {
				u.SetAnnotations(tc.objectAnnotations)
			}

			res, err := dcl.NewResource(u, r.schema)
			if err != nil {
				t.Fatalf("unexpected error creating resource: %v", err)
			}

			ctx := context.Background()
			if err := r.applyChangesForBackwardsCompatibility(ctx, res); err != nil {
				t.Fatalf("applyChangesForBackwardsCompatibility failed: %v", err)
			}

			actualProjectRef, found, err := unstructured.NestedMap(res.Spec, "projectRef")
			if err != nil {
				t.Fatalf("error getting projectRef: %v", err)
			}
			if !found {
				t.Fatalf("projectRef not found in resource spec: %v", res.Spec)
			}

			if !reflect.DeepEqual(tc.expectedProjectRef, actualProjectRef) {
				t.Fatalf("got unexpected projectRef (-want +got):\n%v", cmp.Diff(tc.expectedProjectRef, actualProjectRef))
			}
		})
	}
}
