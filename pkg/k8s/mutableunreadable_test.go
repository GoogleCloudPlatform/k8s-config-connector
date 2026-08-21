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

package k8s_test

import (
	"reflect"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"

	"github.com/google/go-cmp/cmp"
)

var mutableButUnreadableFieldAnnotationTests = []struct {
	name                                    string
	resource                                *krmtotf.Resource
	mutableButUnreadablePaths               [][]string
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
		},
		mutableButUnreadablePaths: [][]string{
			{"fieldA"},
			{"fieldC"},
			{"fieldD"},
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
		},
		mutableButUnreadablePaths: [][]string{
			{"parentField", "fieldA"},
			{"parentField", "fieldC"},
			{"parentField", "fieldD"},
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
		},
		mutableButUnreadablePaths: [][]string{
			{"fieldA"},
			{"fieldC"},
			{"fieldD"},
			{"parentField", "fieldA"},
			{"parentField", "fieldC"},
			{"parentField", "fieldD"},
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
		},
		mutableButUnreadablePaths:               [][]string{},
		expectedMutableButUnreadableFieldsState: map[string]interface{}{},
	},
	{
		name:     "no spec",
		resource: &krmtotf.Resource{},
		mutableButUnreadablePaths: [][]string{
			{"fieldA"},
		},
		expectedMutableButUnreadableFieldsState: map[string]interface{}{},
	},
}

func TestGenerateMutableButUnreadableFieldsAnnotation(t *testing.T) {
	for _, tc := range mutableButUnreadableFieldAnnotationTests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			annotationInString, err := k8s.GenerateMutableButUnreadableFieldsAnnotation(&tc.resource.Resource, tc.mutableButUnreadablePaths)
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

func TestGetMutableButUnreadableFieldsFromAnnotations(t *testing.T) {
	for _, tc := range mutableButUnreadableFieldAnnotationTests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			mutableButUnreadableFields, err := k8s.GetMutableButUnreadableFieldsFromAnnotations(&tc.resource.Resource, tc.mutableButUnreadablePaths)
			if err != nil {
				t.Fatal(err)
			}
			if got, want := mutableButUnreadableFields, tc.expectedMutableButUnreadableFieldsState; !reflect.DeepEqual(got, want) {
				t.Fatalf("unexpected mutable-but-unreadable fields diff (-want +got): \n%v", cmp.Diff(want, got))
			}
		})
	}
}
