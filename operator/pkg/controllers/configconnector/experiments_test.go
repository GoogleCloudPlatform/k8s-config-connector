// Copyright 2025 Google LLC
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

package configconnector

import (
	"context"
	"testing"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/controller"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

func TestApplyMultiClusterLeaderElection(t *testing.T) {
	ctx := context.Background()

	// Helper to get manifest objects
	getManifests := func(t *testing.T) *manifest.Objects {
		return testcontroller.ParseObjects(ctx, t, testcontroller.ClusterModeComponents)
	}

	tests := []struct {
		name          string
		cc            *corev1beta1.ConfigConnector
		expectEnabled bool
	}{
		{
			name: "experiments nil",
			cc: &corev1beta1.ConfigConnector{
				Spec: corev1beta1.ConfigConnectorSpec{
					Experiments: nil,
				},
			},
			expectEnabled: false,
		},
		{
			name: "multicluster lease nil",
			cc: &corev1beta1.ConfigConnector{
				Spec: corev1beta1.ConfigConnectorSpec{
					Experiments: &corev1beta1.CCExperiments{
						MultiClusterLease: nil,
					},
				},
			},
			expectEnabled: false,
		},
		{
			name: "multicluster lease set",
			cc: &corev1beta1.ConfigConnector{
				Spec: corev1beta1.ConfigConnectorSpec{
					Experiments: &corev1beta1.CCExperiments{
						MultiClusterLease: &corev1beta1.MultiClusterLeaseSpec{
							LeaseName:                "foo",
							Namespace:                "bar",
							ClusterCandidateIdentity: "baz",
						},
					},
				},
			},
			expectEnabled: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := getManifests(t)
			r := &Reconciler{} // minimal reconciler, methods used don't access fields

			if err := r.applyExperiments(ctx, tc.cc, m); err != nil {
				t.Fatalf("applyExperiments failed: %v", err)
			}

			foundStatefulSet := false
			for _, item := range m.Items {
				if IsControllerManagerStatefulSet(item) {
					foundStatefulSet = true

					// Check args
					foundArg := false
					err := item.MutateContainers(func(container map[string]interface{}) error {
						name, _, _ := unstructured.NestedString(container, "name")
						if name != "manager" {
							return nil
						}
						args, _, _ := unstructured.NestedStringSlice(container, "args")
						for _, arg := range args {
							if arg == "--leader-election-type=multicluster" {
								foundArg = true
							}
						}
						return nil
					})
					if err != nil {
						t.Fatalf("MutateContainers failed: %v", err)
					}

					if tc.expectEnabled && !foundArg {
						t.Errorf("expected --leader-election-type=multicluster arg, but not found")
					}
					if !tc.expectEnabled && foundArg {
						t.Errorf("did not expect --leader-election-type=multicluster arg, but found")
					}

					// Check annotation
					annotations := item.UnstructuredObject().GetAnnotations()
					// Check template annotations
					templateAnnotations, _, _ := unstructured.NestedStringMap(item.UnstructuredObject().Object, "spec", "template", "metadata", "annotations")

					if tc.expectEnabled {
						if val, ok := templateAnnotations["cnrm.cloud.google.com/lease-name"]; !ok || val != "foo" {
							t.Errorf("expected pod template annotation cnrm.cloud.google.com/lease-name to be 'foo', got '%s'", val)
						}
						if val, ok := templateAnnotations["cnrm.cloud.google.com/lease-namespace"]; !ok || val != "bar" {
							t.Errorf("expected pod template annotation cnrm.cloud.google.com/lease-namespace to be 'bar', got '%s'", val)
						}
						if val, ok := templateAnnotations["cnrm.cloud.google.com/lease-identity"]; !ok || val != "baz" {
							t.Errorf("expected pod template annotation cnrm.cloud.google.com/lease-identity to be 'baz', got '%s'", val)
						}
					} else {
						if _, ok := templateAnnotations["cnrm.cloud.google.com/lease-name"]; ok {
							t.Errorf("did not expect pod template annotation cnrm.cloud.google.com/lease-name")
						}
					}

					// Log annotations for debugging
					t.Logf("StatefulSet Annotations: %v", annotations)
					t.Logf("Pod Template Annotations: %v", templateAnnotations)
				}
			}
			if !foundStatefulSet {
				t.Fatalf("StatefulSet not found in manifest")
			}
		})
	}
}
