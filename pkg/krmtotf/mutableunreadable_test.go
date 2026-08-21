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

package krmtotf_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"
)

func TestMutableButUnreadableFieldsAnnotationFor(t *testing.T) {
	tests := []struct {
		name                                    string
		resource                                *krmtotf.Resource
		expectedMutableButUnreadableFieldsState map[string]interface{}
	}{
		{
			name: "top-level fields",
			resource: &krmtotf.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"fieldA": "val1",
						"fieldB": "val2",
						"fieldC": map[string]interface{}{
							"field1": "val1",
							"field2": "val2",
						},
						"fieldD": []interface{}{
							"val1",
							"val2",
						},
					},
				},
				ResourceConfig: v1alpha1.ResourceConfig{
					MutableButUnreadableFields: []string{
						"field_a",
						"field_c",
						"field_d",
					},
				},
			},
			expectedMutableButUnreadableFieldsState: map[string]interface{}{
				"spec": map[string]interface{}{
					"fieldA": "val1",
					"fieldC": map[string]interface{}{
						"field1": "val1",
						"field2": "val2",
					},
					"fieldD": []interface{}{
						"val1",
						"val2",
					},
				},
			},
		},
		{
			name: "nested fields",
			resource: &krmtotf.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"parentField": map[string]interface{}{
							"fieldA": "val1",
							"fieldB": "val2",
							"fieldC": map[string]interface{}{
								"field1": "val1",
								"field2": "val2",
							},
							"fieldD": []interface{}{
								"val1",
								"val2",
							},
						},
					},
				},
				ResourceConfig: v1alpha1.ResourceConfig{
					MutableButUnreadableFields: []string{
						"parent_field.field_a",
						"parent_field.field_c",
						"parent_field.field_d",
					},
				},
			},
			expectedMutableButUnreadableFieldsState: map[string]interface{}{
				"spec": map[string]interface{}{
					"parentField": map[string]interface{}{
						"fieldA": "val1",
						"fieldC": map[string]interface{}{
							"field1": "val1",
							"field2": "val2",
						},
						"fieldD": []interface{}{
							"val1",
							"val2",
						},
					},
				},
			},
		},
		{
			name: "no mutable-but-unreadable fields set in spec",
			resource: &krmtotf.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"fieldB": "val2",
						"parentField": map[string]interface{}{
							"fieldB": "val2",
						},
					},
				},
				ResourceConfig: v1alpha1.ResourceConfig{
					MutableButUnreadableFields: []string{
						"field_a",
						"field_c",
						"field_d",
						"parent_field.field_a",
						"parent_field.field_c",
						"parent_field.field_d",
					},
				},
			},
			expectedMutableButUnreadableFieldsState: map[string]interface{}{},
		},
		{
			name: "no fields marked mutable-but-unreadable",
			resource: &krmtotf.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"fieldA": "val1",
					},
				},
				ResourceConfig: v1alpha1.ResourceConfig{},
			},
			expectedMutableButUnreadableFieldsState: map[string]interface{}{},
		},
		{
			name: "no spec",
			resource: &krmtotf.Resource{
				ResourceConfig: v1alpha1.ResourceConfig{
					MutableButUnreadableFields: []string{
						"field_a",
					},
				},
			},
			expectedMutableButUnreadableFieldsState: map[string]interface{}{},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			annotationInString, err := krmtotf.MutableButUnreadableFieldsAnnotationFor(tc.resource)
			if err != nil {
				t.Fatal(err)
			}

			expectedStateInString, err := util.MarshalToJSONString(tc.expectedMutableButUnreadableFieldsState)
			if err != nil {
				t.Fatalf("error marshaling the expected state to string: %v", err)
			}
			if got, want := annotationInString, expectedStateInString; got != want {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}
