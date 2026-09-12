/*
Copyright 2026 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package parent

import (
	"context"
	"testing"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestDetermineControllerType_Overrides(t *testing.T) {
	scheme := runtime.NewScheme()
	if err := corev1beta1.SchemeBuilder.AddToScheme(scheme); err != nil {
		t.Fatalf("failed to add core v1beta1 scheme to scheme: %v", err)
	}

	testGVK := schema.GroupVersionKind{
		Group:   "bigquery.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "BigQueryDataset",
	}
	testGK := "BigQueryDataset.bigquery.cnrm.cloud.google.com"

	tests := []struct {
		name                string
		objects             []runtime.Object
		resourceAnnotations map[string]string
		expectedResult      k8s.ReconcilerType
		expectError         bool
	}{
		{
			name:           "no CC and no CCC: falls back to default static config (tf)",
			objects:        []runtime.Object{},
			expectedResult: k8s.ReconcilerTypeTerraform,
		},
		{
			name: "CC exists with no overrides: falls back to default static config (tf)",
			objects: []runtime.Object{
				&corev1beta1.ConfigConnector{
					ObjectMeta: metav1.ObjectMeta{
						Name: corev1beta1.ConfigConnectorAllowedName,
					},
					Spec: corev1beta1.ConfigConnectorSpec{
						Mode: "namespaced",
					},
				},
				&corev1beta1.ConfigConnectorContext{
					ObjectMeta: metav1.ObjectMeta{
						Name:      corev1beta1.ConfigConnectorContextAllowedName,
						Namespace: "test-namespace",
					},
				},
			},
			expectedResult: k8s.ReconcilerTypeTerraform,
		},
		{
			name: "CC exists with overrides globally (dcl), no CCC: returns error because KCC is in namespaced mode but no CCC exists",
			objects: []runtime.Object{
				&corev1beta1.ConfigConnector{
					ObjectMeta: metav1.ObjectMeta{
						Name: corev1beta1.ConfigConnectorAllowedName,
					},
					Spec: corev1beta1.ConfigConnectorSpec{
						Mode: "namespaced",
						Experiments: &corev1beta1.CCExperiments{
							ControllerOverrides: map[string]k8s.ReconcilerType{
								testGK: k8s.ReconcilerTypeDCL,
							},
						},
					},
				},
			},
			expectError: true,
		},
		{
			name: "CC exists with overrides globally (dcl) and KCC is in cluster mode: matches global CC override",
			objects: []runtime.Object{
				&corev1beta1.ConfigConnector{
					ObjectMeta: metav1.ObjectMeta{
						Name: corev1beta1.ConfigConnectorAllowedName,
					},
					Spec: corev1beta1.ConfigConnectorSpec{
						Mode: "cluster",
						Experiments: &corev1beta1.CCExperiments{
							ControllerOverrides: map[string]k8s.ReconcilerType{
								testGK: k8s.ReconcilerTypeDCL,
							},
						},
					},
				},
			},
			expectedResult: k8s.ReconcilerTypeDCL,
		},
		{
			name: "CCC exists in namespace with overrides: matches namespace CCC override",
			objects: []runtime.Object{
				&corev1beta1.ConfigConnector{
					ObjectMeta: metav1.ObjectMeta{
						Name: corev1beta1.ConfigConnectorAllowedName,
					},
					Spec: corev1beta1.ConfigConnectorSpec{
						Mode: "namespaced",
					},
				},
				&corev1beta1.ConfigConnectorContext{
					ObjectMeta: metav1.ObjectMeta{
						Name:      corev1beta1.ConfigConnectorContextAllowedName,
						Namespace: "test-namespace",
					},
					Spec: corev1beta1.ConfigConnectorContextSpec{
						Experiments: &corev1beta1.Experiments{
							ControllerOverrides: map[string]k8s.ReconcilerType{
								testGK: k8s.ReconcilerTypeDirect,
							},
						},
					},
				},
			},
			expectedResult: k8s.ReconcilerTypeDirect,
		},
		{
			name: "Both CC and CCC exist with overrides: CCC override takes precedence",
			objects: []runtime.Object{
				&corev1beta1.ConfigConnector{
					ObjectMeta: metav1.ObjectMeta{
						Name: corev1beta1.ConfigConnectorAllowedName,
					},
					Spec: corev1beta1.ConfigConnectorSpec{
						Mode: "namespaced",
						Experiments: &corev1beta1.CCExperiments{
							ControllerOverrides: map[string]k8s.ReconcilerType{
								testGK: k8s.ReconcilerTypeDCL,
							},
						},
					},
				},
				&corev1beta1.ConfigConnectorContext{
					ObjectMeta: metav1.ObjectMeta{
						Name:      corev1beta1.ConfigConnectorContextAllowedName,
						Namespace: "test-namespace",
					},
					Spec: corev1beta1.ConfigConnectorContextSpec{
						Experiments: &corev1beta1.Experiments{
							ControllerOverrides: map[string]k8s.ReconcilerType{
								testGK: k8s.ReconcilerTypeDirect,
							},
						},
					},
				},
			},
			expectedResult: k8s.ReconcilerTypeDirect,
		},
		{
			name: "Both CC and CCC exist with overrides, but CCC does not override this resource type: matches global CC override",
			objects: []runtime.Object{
				&corev1beta1.ConfigConnector{
					ObjectMeta: metav1.ObjectMeta{
						Name: corev1beta1.ConfigConnectorAllowedName,
					},
					Spec: corev1beta1.ConfigConnectorSpec{
						Mode: "namespaced",
						Experiments: &corev1beta1.CCExperiments{
							ControllerOverrides: map[string]k8s.ReconcilerType{
								testGK: k8s.ReconcilerTypeDCL,
							},
						},
					},
				},
				&corev1beta1.ConfigConnectorContext{
					ObjectMeta: metav1.ObjectMeta{
						Name:      corev1beta1.ConfigConnectorContextAllowedName,
						Namespace: "test-namespace",
					},
					Spec: corev1beta1.ConfigConnectorContextSpec{
						Experiments: &corev1beta1.Experiments{
							ControllerOverrides: map[string]k8s.ReconcilerType{
								"SomeOtherGK": k8s.ReconcilerTypeDirect,
							},
						},
					},
				},
			},
			expectedResult: k8s.ReconcilerTypeDCL,
		},
		{
			name: "CCC exists but in a different namespace: returns error because CCC is missing for the resource namespace",
			objects: []runtime.Object{
				&corev1beta1.ConfigConnector{
					ObjectMeta: metav1.ObjectMeta{
						Name: corev1beta1.ConfigConnectorAllowedName,
					},
					Spec: corev1beta1.ConfigConnectorSpec{
						Mode: "namespaced",
						Experiments: &corev1beta1.CCExperiments{
							ControllerOverrides: map[string]k8s.ReconcilerType{
								testGK: k8s.ReconcilerTypeDCL,
							},
						},
					},
				},
				&corev1beta1.ConfigConnectorContext{
					ObjectMeta: metav1.ObjectMeta{
						Name:      corev1beta1.ConfigConnectorContextAllowedName,
						Namespace: "other-namespace",
					},
					Spec: corev1beta1.ConfigConnectorContextSpec{
						Experiments: &corev1beta1.Experiments{
							ControllerOverrides: map[string]k8s.ReconcilerType{
								testGK: k8s.ReconcilerTypeDirect,
							},
						},
					},
				},
			},
			expectError: true,
		},
		{
			name: "resource has direct reconciler annotation: matches direct, taking precedence over everything",
			objects: []runtime.Object{
				&corev1beta1.ConfigConnector{
					ObjectMeta: metav1.ObjectMeta{
						Name: corev1beta1.ConfigConnectorAllowedName,
					},
					Spec: corev1beta1.ConfigConnectorSpec{
						Mode: "namespaced",
						Experiments: &corev1beta1.CCExperiments{
							ControllerOverrides: map[string]k8s.ReconcilerType{
								testGK: k8s.ReconcilerTypeDCL,
							},
						},
					},
				},
				&corev1beta1.ConfigConnectorContext{
					ObjectMeta: metav1.ObjectMeta{
						Name:      corev1beta1.ConfigConnectorContextAllowedName,
						Namespace: "test-namespace",
					},
					Spec: corev1beta1.ConfigConnectorContextSpec{
						Experiments: &corev1beta1.Experiments{
							ControllerOverrides: map[string]k8s.ReconcilerType{
								testGK: k8s.ReconcilerTypeTerraform,
							},
						},
					},
				},
			},
			resourceAnnotations: map[string]string{
				k8s.AlphaReconcilerAnnotation: "direct",
			},
			expectedResult: k8s.ReconcilerTypeDirect,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			resource := &unstructured.Unstructured{}
			resource.SetGroupVersionKind(testGVK)
			resource.SetName("test-dataset")
			resource.SetNamespace("test-namespace")
			if tc.resourceAnnotations != nil {
				resource.SetAnnotations(tc.resourceAnnotations)
			}

			fakeClient := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(tc.objects...).Build()
			r := &ParentReconciler{
				Client: fakeClient,
				gvk:    testGVK,
			}

			result, err := r.determineControllerType(context.Background(), resource)
			if tc.expectError {
				if err == nil {
					t.Errorf("expected error, but got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tc.expectedResult {
				t.Errorf("unexpected controller type. Got: %v, Want: %v", result, tc.expectedResult)
			}
		})
	}
}
