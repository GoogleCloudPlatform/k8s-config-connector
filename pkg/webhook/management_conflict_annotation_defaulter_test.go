// Copyright 2022 Google LLC
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
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/managementconflict"
	testutil "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testdclschemaloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/dclschemaloader"
	testservicemetadataloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemetadataloader"

	"github.com/nasa9084/go-openapi"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestDefaultManagementConflictAnnotationForDCLBasedResources(t *testing.T) {
	tests := []struct {
		name   string
		obj    *unstructured.Unstructured
		newObj *unstructured.Unstructured
		ns     *corev1.Namespace
		schema *openapi.Schema
		denied bool
	}{
		{
			name: "default annotation to 'none' if not set on either namespace or resource",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							managementconflict.FullyQualifiedAnnotation: "none",
						},
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-ns",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"labels": &openapi.Schema{
						Type: "string",
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-labels": "labels",
				},
			},
		},
		{
			name: "default annotation from namespace",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							managementconflict.FullyQualifiedAnnotation: "none",
						},
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-ns",
					Annotations: map[string]string{
						managementconflict.FullyQualifiedAnnotation: "none",
					},
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"labels": &openapi.Schema{
						Type: "string",
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-labels": "labels",
				},
			},
		},
		{
			name: "no defaulting needed since resource specifies annotation",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							managementconflict.FullyQualifiedAnnotation: "resource",
						},
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							managementconflict.FullyQualifiedAnnotation: "resource",
						},
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-ns",
					Annotations: map[string]string{
						managementconflict.FullyQualifiedAnnotation: "none",
					},
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"labels": &openapi.Schema{
						Type: "string",
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-labels": "labels",
				},
			},
		},
		{
			name: "default annotation to 'none' if resource doesn't support management conflict prevention",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							managementconflict.FullyQualifiedAnnotation: "none",
						},
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-ns",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
			},
		},
		{
			name: "default annotation to 'none' if resource doesn't support management conflict prevention even if namespace has annotation",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							managementconflict.FullyQualifiedAnnotation: "none",
						},
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-ns",
					Annotations: map[string]string{
						managementconflict.FullyQualifiedAnnotation: "resource",
					},
				},
			},
			schema: &openapi.Schema{
				Type: "object",
			},
		},
	}
	smLoader := testservicemetadataloader.NewForUnitTest()
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			dclSchemaMap := make(map[string]*openapi.Schema)
			dclSchemaMap["test1_beta_bar"] = tc.schema
			dclSchemaLoader := testdclschemaloader.New(dclSchemaMap)
			response := defaultManagementConflictAnnotationForDCLBasedResources(tc.obj, tc.ns, dclSchemaLoader, smLoader)
			expectedResponse := constructPatchResponse(tc.obj, tc.newObj)
			if !testutil.Equals(t, response, expectedResponse) {
				t.Fatalf("expect to get response %v, but got %v", expectedResponse, response)
			}
		})
	}
}
