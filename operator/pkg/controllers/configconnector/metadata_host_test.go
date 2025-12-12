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

	customizev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/customize/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/controllers"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/controller"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestApplyMetadataHost(t *testing.T) {
	t.Parallel()

	statefulSetGVK := schema.GroupVersionKind{
		Group:   appsv1.SchemeGroupVersion.Group,
		Version: appsv1.SchemeGroupVersion.Version,
		Kind:    "StatefulSet",
	}

	deploymentGVK := schema.GroupVersionKind{
		Group:   appsv1.SchemeGroupVersion.Group,
		Version: appsv1.SchemeGroupVersion.Version,
		Kind:    "Deployment",
	}

	tests := []struct {
		name           string
		controllerName string
		controllerGVK  schema.GroupVersionKind
		metadataHost   string
		wantEnvVar     bool
	}{
		{
			name:           "metadataHost not set - no changes",
			controllerName: "cnrm-controller-manager",
			controllerGVK:  statefulSetGVK,
			metadataHost:   "",
			wantEnvVar:     false,
		},
		{
			name:           "metadataHost set on StatefulSet - inject env var",
			controllerName: "cnrm-controller-manager",
			controllerGVK:  statefulSetGVK,
			metadataHost:   "metadata.google.internal",
			wantEnvVar:     true,
		},
		{
			name:           "metadataHost set on Deployment - inject env var",
			controllerName: "cnrm-webhook-manager",
			controllerGVK:  deploymentGVK,
			metadataHost:   "metadata.google.internal",
			wantEnvVar:     true,
		},
		{
			name:           "IPv6 address format",
			controllerName: "cnrm-controller-manager",
			controllerGVK:  statefulSetGVK,
			metadataHost:   "[fd20:ce::254]",
			wantEnvVar:     true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := t.Context()

			m := testcontroller.ParseObjects(ctx, t, testcontroller.GetClusterModeWorkloadIdentityManifest())

			err := controllers.ApplyMetadataHost(m, tc.controllerName, tc.controllerGVK, tc.metadataHost)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Check the specified controller for env var
			for _, item := range m.Items {
				if item.GroupVersionKind() != tc.controllerGVK {
					continue
				}
				if item.GetName() != tc.controllerName {
					continue
				}

				hasEnvVar := checkForEnvVar(t, item.UnstructuredObject(), "GCE_METADATA_HOST")
				if hasEnvVar != tc.wantEnvVar {
					t.Errorf("kind=%s name=%s: GCE_METADATA_HOST present=%v, want=%v",
						item.Kind, item.GetName(), hasEnvVar, tc.wantEnvVar)
				}
				if tc.wantEnvVar && tc.metadataHost != "" {
					value := getEnvVarValue(t, item.UnstructuredObject(), "GCE_METADATA_HOST")
					if value != tc.metadataHost {
						t.Errorf("GCE_METADATA_HOST value = %q, want %q", value, tc.metadataHost)
					}
				}
			}
		})
	}
}

func TestApplyControllerResourceCR_MetadataHost(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		controllerName string
		metadataHost   string
		wantEnvVar     bool
	}{
		{
			name:           "cnrm-controller-manager with metadataHost",
			controllerName: "cnrm-controller-manager",
			metadataHost:   "metadata.google.internal",
			wantEnvVar:     true,
		},
		{
			name:           "cnrm-webhook-manager with metadataHost",
			controllerName: "cnrm-webhook-manager",
			metadataHost:   "metadata.google.internal",
			wantEnvVar:     true,
		},
		{
			name:           "cnrm-controller-manager without metadataHost",
			controllerName: "cnrm-controller-manager",
			metadataHost:   "",
			wantEnvVar:     false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := t.Context()

			cr := &customizev1beta1.ControllerResource{}
			cr.Name = tc.controllerName
			cr.Spec.MetadataHost = tc.metadataHost

			m := testcontroller.ParseObjects(ctx, t, testcontroller.GetClusterModeWorkloadIdentityManifest())

			// Determine the GVK based on controller name
			var controllerGVK schema.GroupVersionKind
			switch tc.controllerName {
			case "cnrm-controller-manager", "cnrm-deletiondefender", "cnrm-unmanaged-detector":
				controllerGVK = schema.GroupVersionKind{
					Group:   appsv1.SchemeGroupVersion.Group,
					Version: appsv1.SchemeGroupVersion.Version,
					Kind:    "StatefulSet",
				}
			case "cnrm-webhook-manager", "cnrm-resource-stats-recorder":
				controllerGVK = schema.GroupVersionKind{
					Group:   appsv1.SchemeGroupVersion.Group,
					Version: appsv1.SchemeGroupVersion.Version,
					Kind:    "Deployment",
				}
			}

			err := controllers.ApplyMetadataHost(m, cr.Name, controllerGVK, cr.Spec.MetadataHost)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Check for env var
			for _, item := range m.Items {
				if item.GroupVersionKind() != controllerGVK {
					continue
				}
				if item.GetName() != tc.controllerName {
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

// getEnvVarValue returns the value of a specific env var from a workload object
func getEnvVarValue(t *testing.T, obj *unstructured.Unstructured, envName string) string {
	t.Helper()

	containers, found, _ := unstructured.NestedSlice(obj.Object, "spec", "template", "spec", "containers")
	if !found {
		return ""
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
				value, _, _ := unstructured.NestedString(envMap, "value")
				return value
			}
		}
	}
	return ""
}
