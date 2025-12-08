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
	"testing"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/controller"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/main"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestAddMetadataHostEnvVar(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		container    map[string]interface{}
		metadataHost string
		wantEnvValue string
		wantPreserve bool // true if existing GCE_METADATA_HOST should be preserved
	}{
		{
			name: "add to container with no env vars",
			container: map[string]interface{}{
				"name": "manager",
			},
			metadataHost: "metadata.google.internal",
			wantEnvValue: "metadata.google.internal",
		},
		{
			name: "add to container with existing env vars",
			container: map[string]interface{}{
				"name": "manager",
				"env": []interface{}{
					map[string]interface{}{"name": "EXISTING_VAR", "value": "existing"},
				},
			},
			metadataHost: "metadata.google.internal",
			wantEnvValue: "metadata.google.internal",
		},
		{
			name: "preserve existing GCE_METADATA_HOST",
			container: map[string]interface{}{
				"name": "manager",
				"env": []interface{}{
					map[string]interface{}{"name": "GCE_METADATA_HOST", "value": "custom.endpoint"},
				},
			},
			metadataHost: "metadata.google.internal",
			wantEnvValue: "custom.endpoint", // preserved, not overwritten
			wantPreserve: true,
		},
		{
			name: "IPv6 address format",
			container: map[string]interface{}{
				"name": "manager",
			},
			metadataHost: "[fd20:ce::254]",
			wantEnvValue: "[fd20:ce::254]",
		},
		{
			name: "DNS hostname with port",
			container: map[string]interface{}{
				"name": "manager",
			},
			metadataHost: "metadata.google.internal:8080",
			wantEnvValue: "metadata.google.internal:8080",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Make a copy to avoid mutation issues
			container := copyContainer(tc.container)

			err := addMetadataHostEnvVar(container, tc.metadataHost)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			envs, _, _ := unstructured.NestedSlice(container, "env")
			var found bool
			for _, e := range envs {
				envMap, ok := e.(map[string]interface{})
				if !ok {
					continue
				}
				name, _, _ := unstructured.NestedString(envMap, "name")
				if name == "GCE_METADATA_HOST" {
					found = true
					value, _, _ := unstructured.NestedString(envMap, "value")
					if value != tc.wantEnvValue {
						t.Errorf("env var value = %q, want %q", value, tc.wantEnvValue)
					}
				}
			}
			if !found {
				t.Errorf("env var GCE_METADATA_HOST not found")
			}
		})
	}
}

func TestApplyMetadataHost(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		metadataHost string
		wantEnvVar   bool // whether GCE_METADATA_HOST should be in containers
	}{
		{
			name:         "metadataHost not set - no changes",
			metadataHost: "",
			wantEnvVar:   false,
		},
		{
			name:         "metadataHost set - inject env var",
			metadataHost: "metadata.google.internal",
			wantEnvVar:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := t.Context()

			mgr, stop := testmain.StartTestManagerFromNewTestEnv()
			defer stop()
			c := mgr.GetClient()

			cc := &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "configconnector.core.cnrm.cloud.google.com",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode:         "cluster",
					MetadataHost: tc.metadataHost,
				},
			}

			m := testcontroller.ParseObjects(ctx, t, testcontroller.GetClusterModeWorkloadIdentityManifest())
			r := newConfigConnectorReconciler(c)

			if err := r.applyMetadataHost(ctx, cc, m); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Check StatefulSet and Deployment containers for env var
			for _, item := range m.Items {
				if item.Kind != "StatefulSet" && item.Kind != "Deployment" {
					continue
				}

				hasEnvVar := checkForEnvVar(t, item.UnstructuredObject(), "GCE_METADATA_HOST")
				if hasEnvVar != tc.wantEnvVar {
					t.Errorf("kind=%s name=%s: GCE_METADATA_HOST present=%v, want=%v",
						item.Kind, item.GetName(), hasEnvVar, tc.wantEnvVar)
				}
			}
		})
	}
}

func TestTransformForMetadataHost(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		metadataHost string
		wantErr      bool
	}{
		{
			name:         "valid hostname",
			metadataHost: "metadata.google.internal",
			wantErr:      false,
		},
		{
			name:         "empty metadataHost",
			metadataHost: "",
			wantErr:      false,
		},
		{
			name:         "IPv6 address",
			metadataHost: "[fd20:ce::254]",
			wantErr:      false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := t.Context()

			mgr, stop := testmain.StartTestManagerFromNewTestEnv()
			defer stop()
			c := mgr.GetClient()

			r := newConfigConnectorReconciler(c)
			transform := r.transformForMetadataHost()

			cc := &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{Name: "configconnector.core.cnrm.cloud.google.com"},
				Spec:       corev1beta1.ConfigConnectorSpec{MetadataHost: tc.metadataHost},
			}

			m := testcontroller.ParseObjects(ctx, t, testcontroller.GetClusterModeWorkloadIdentityManifest())

			err := transform(ctx, cc, m)
			if (err != nil) != tc.wantErr {
				t.Errorf("transform error = %v, wantErr = %v", err, tc.wantErr)
			}
		})
	}
}

// copyContainer creates a deep copy of a container map
func copyContainer(src map[string]interface{}) map[string]interface{} {
	dst := make(map[string]interface{})
	for k, v := range src {
		switch val := v.(type) {
		case []interface{}:
			newSlice := make([]interface{}, len(val))
			for i, item := range val {
				if m, ok := item.(map[string]interface{}); ok {
					newSlice[i] = copyContainer(m)
				} else {
					newSlice[i] = item
				}
			}
			dst[k] = newSlice
		case map[string]interface{}:
			dst[k] = copyContainer(val)
		default:
			dst[k] = v
		}
	}
	return dst
}

// checkForEnvVar checks if a workload object has a specific env var in any container
func checkForEnvVar(t *testing.T, obj *unstructured.Unstructured, envName string) bool {
	t.Helper()

	containers, found, _ := unstructured.NestedSlice(obj.Object, "spec", "template", "spec", "containers")
	if !found {
		return false
	}

	for _, c := range containers {
		container, ok := c.(map[string]interface{})
		if !ok {
			continue
		}
		envs, _, _ := unstructured.NestedSlice(container, "env")
		for _, e := range envs {
			envMap, ok := e.(map[string]interface{})
			if !ok {
				continue
			}
			name, _, _ := unstructured.NestedString(envMap, "name")
			if name == envName {
				return true
			}
		}
	}
	return false
}
