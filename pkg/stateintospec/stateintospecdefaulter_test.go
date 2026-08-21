// Copyright 2024 Google LLC
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

package stateintospec

import (
	"context"
	"log"
	"testing"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/kubernetes/scheme"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/main"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
)

func init() {
	s := scheme.Scheme
	if err := corev1beta1.SchemeBuilder.AddToScheme(s); err != nil {
		log.Fatalf("error registering core kcc operator scheme: %v", err)
	}
}

func TestStateIntoSpecDefaulter_ApplyDefaults(t *testing.T) {
	t.Parallel()
	absentValue := corev1beta1.StateIntoSpecAbsent
	mergeValue := corev1beta1.StateIntoSpecMerge
	tests := []struct {
		name          string
		resource      *unstructured.Unstructured
		cc            *corev1beta1.ConfigConnector
		ccc           *corev1beta1.ConfigConnectorContext
		expectChanged bool
		expectValue   string
		expectError   bool
	}{
		{
			name: "use v1beta1 default value for resources supporting 'merge'",
			resource: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "pubsub.cnrm.cloud.google.com/v1beta1",
					"kind":       "PubSubTopic",
					"metadata": map[string]interface{}{
						"name":      "test-name",
						"namespace": "test-ns",
					},
				},
			},
			expectChanged: true,
			expectValue:   "absent",
		},
		{
			name: "use 'absent' if resource doesn't support 'merge'",
			resource: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test.cnrm.cloud.google.com/v1beta1",
					"kind":       "NewKind",
					"metadata": map[string]interface{}{
						"name":      "test-name",
						"namespace": "test-ns",
					},
				},
			},
			expectChanged: true,
			expectValue:   "absent",
		},
		{
			name: "use cc default value",
			resource: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "pubsub.cnrm.cloud.google.com/v1beta1",
					"kind":       "PubSubTopic",
					"metadata": map[string]interface{}{
						"name":      "test-name",
						"namespace": "test-ns",
					},
				},
			},
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: corev1beta1.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode:          "cluster",
					StateIntoSpec: &absentValue,
				},
			},
			expectChanged: true,
			expectValue:   "absent",
		},
		{
			name: "use ccc default value when mode is unset in cc",
			resource: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "pubsub.cnrm.cloud.google.com/v1beta1",
					"kind":       "PubSubTopic",
					"metadata": map[string]interface{}{
						"name":      "test-name",
						"namespace": "test-ns",
					},
				},
			},
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: corev1beta1.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					StateIntoSpec: &mergeValue,
				},
			},
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      corev1beta1.ConfigConnectorContextAllowedName,
					Namespace: "test-ns",
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					StateIntoSpec: &absentValue,
				},
			},
			expectChanged: true,
			expectValue:   "absent",
		},
		{
			name: "use ccc default value",
			resource: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "pubsub.cnrm.cloud.google.com/v1beta1",
					"kind":       "PubSubTopic",
					"metadata": map[string]interface{}{
						"name":      "test-name",
						"namespace": "test-ns",
					},
				},
			},
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: corev1beta1.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      corev1beta1.ConfigConnectorContextAllowedName,
					Namespace: "test-ns",
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					StateIntoSpec: &absentValue,
				},
			},
			expectChanged: true,
			expectValue:   "absent",
		},
		{
			name: "error due to ccc not found when mode is unset in cc",
			resource: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "pubsub.cnrm.cloud.google.com/v1beta1",
					"kind":       "PubSubTopic",
					"metadata": map[string]interface{}{
						"name":      "test-name",
						"namespace": "test-ns",
					},
				},
			},
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: corev1beta1.ConfigConnectorAllowedName,
				},
			},
			expectError: true,
		},
		{
			name: "error due to ccc not found",
			resource: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "pubsub.cnrm.cloud.google.com/v1beta1",
					"kind":       "PubSubTopic",
					"metadata": map[string]interface{}{
						"name":      "test-name",
						"namespace": "test-ns",
					},
				},
			},
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: corev1beta1.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			expectError: true,
		},
		{
			name: "'merge' is valid for allowlisted kind",
			resource: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "pubsub.cnrm.cloud.google.com/v1beta1",
					"kind":       "PubSubTopic",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.StateIntoSpecAnnotation: "merge",
						},
						"name":      "test-name",
						"namespace": "test-ns",
					},
				},
			},
			expectChanged: false,
			expectValue:   "merge",
		},
		{
			name: "'merge' is invalid for non-allowlisted kind",
			resource: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "vertexai.cnrm.cloud.google.com/v1beta1",
					"kind":       "VertexAIDataset",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.StateIntoSpecAnnotation: "merge",
						},
						"name":      "test-name",
						"namespace": "test-ns",
					},
				},
			},
			expectError: true,
		},
		{
			name: "invalid value",
			resource: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "pubsub.cnrm.cloud.google.com/v1beta1",
					"kind":       "PubSubTopic",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.StateIntoSpecAnnotation: "invalid_value",
						},
						"name":      "test-name",
						"namespace": "test-ns",
					},
				},
			},
			expectError: true,
		},
		{
			name: "empty value",
			resource: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "pubsub.cnrm.cloud.google.com/v1beta1",
					"kind":       "PubSubTopic",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.StateIntoSpecAnnotation: "",
						},
						"name":      "test-name",
						"namespace": "test-ns",
					},
				},
			},
			expectError: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mgr, stop := testmain.StartTestManagerFromNewTestEnv()
			defer stop()
			client := mgr.GetClient()
			ns := &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: tc.resource.GetNamespace(),
				},
			}
			if err := client.Create(ctx, ns); err != nil {
				t.Fatalf("error creating %+v: %v", ns.GroupVersionKind(), err)
			}
			if tc.cc != nil {
				if err := client.Create(ctx, tc.cc); err != nil {
					t.Fatalf("error creating %+v: %v", tc.cc.GroupVersionKind(), err)
				}
			}
			if tc.ccc != nil {
				ns := &corev1.Namespace{
					ObjectMeta: metav1.ObjectMeta{
						Name: tc.ccc.GetNamespace(),
					},
				}
				if err := client.Create(ctx, ns); !errors.IsAlreadyExists(err) {
					t.Fatalf("error creating %+v: %v", ns.GroupVersionKind(), err)
				}
				if err := client.Create(ctx, tc.ccc); err != nil {
					t.Fatalf("error creating %+v: %v", tc.ccc.GroupVersionKind(), err)
				}
			}
			defaulter := NewStateIntoSpecDefaulter(client)
			changed, err := defaulter.ApplyDefaults(ctx, k8s.ReconcilerTypeTerraform, tc.resource)
			if tc.expectError {
				if err == nil {
					t.Fatalf("got nil, but expect an error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if changed != tc.expectChanged {
				t.Errorf("'changed': got %v, want %v", changed, tc.expectChanged)
			}
			value, _ := k8s.GetAnnotation(k8s.StateIntoSpecAnnotation, tc.resource)
			if value != tc.expectValue {
				t.Errorf("state-into-spec value: got %v, want %v", value, tc.expectValue)
			}
		})
	}
}
