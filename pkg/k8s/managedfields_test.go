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
	"encoding/json"
	"reflect"
	"testing"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/deepcopy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testk8s "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/k8s"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/structured-merge-diff/v4/fieldpath"
)

var emptyObject = make(map[string]interface{})

func TestConstructManagedFieldSet(t *testing.T) {
	tests := []struct {
		name                string
		managedFieldEntries []v1.ManagedFieldsEntry
		expectedSet         *fieldpath.Set
	}{
		{
			name: "fields from separate managers are combined",
			managedFieldEntries: []v1.ManagedFieldsEntry{
				mapToManagedFieldEntry(t, "managerA", map[string]interface{}{
					"f:spec": map[string]interface{}{
						".":              emptyObject,
						"f:simpleFieldA": emptyObject,
						"f:nestedObjectA": map[string]interface{}{
							".":              emptyObject,
							"f:nestedFieldA": emptyObject,
						},
					},
				}),
				mapToManagedFieldEntry(t, "managerB", map[string]interface{}{
					"f:spec": map[string]interface{}{
						".":              emptyObject,
						"f:simpleFieldB": emptyObject,
						"f:nestedObjectB": map[string]interface{}{
							".":              emptyObject,
							"f:nestedFieldB": emptyObject,
						},
					},
				}),
			},
			expectedSet: testk8s.MapToFieldPathSet(t, map[string]interface{}{
				"f:simpleFieldA": emptyObject,
				"f:simpleFieldB": emptyObject,
				"f:nestedObjectA": map[string]interface{}{
					".":              emptyObject,
					"f:nestedFieldA": emptyObject,
				},
				"f:nestedObjectB": map[string]interface{}{
					".":              emptyObject,
					"f:nestedFieldB": emptyObject,
				},
			}),
		},
		{
			name: "fields from the cnrm-controller-manager manager are ignored",
			managedFieldEntries: []v1.ManagedFieldsEntry{
				mapToManagedFieldEntry(t, "managerA", map[string]interface{}{
					"f:spec": map[string]interface{}{
						".":              emptyObject,
						"f:simpleFieldA": emptyObject,
						"f:nestedObjectA": map[string]interface{}{
							".":              emptyObject,
							"f:nestedFieldA": emptyObject,
						},
					},
				}),
				mapToManagedFieldEntry(t, k8s.ControllerManagedFieldManager, map[string]interface{}{
					"f:spec": map[string]interface{}{
						".":                emptyObject,
						"f:simpleFieldKCC": emptyObject,
						"f:nestedObjectKCC": map[string]interface{}{
							".":                emptyObject,
							"f:nestedFieldKCC": emptyObject,
						},
					},
				}),
			},
			expectedSet: testk8s.MapToFieldPathSet(t, map[string]interface{}{
				"f:simpleFieldA": emptyObject,
				"f:nestedObjectA": map[string]interface{}{
					".":              emptyObject,
					"f:nestedFieldA": emptyObject,
				},
			}),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			managedFieldSet, err := k8s.ConstructManagedFieldsV1Set(tc.managedFieldEntries)
			if err != nil {
				t.Error("error constructing managed field set:", err)
				return
			}
			if !managedFieldSet.Equals(tc.expectedSet) {
				t.Errorf("actual and expected sets do not match: actual: %v, expected: %v",
					string(fieldPathSetToJSON(t, managedFieldSet)),
					string(fieldPathSetToJSON(t, tc.expectedSet)))
				return
			}
		})
	}
}

var schema = &apiextensions.JSONSchemaProps{
	Properties: map[string]apiextensions.JSONSchemaProps{
		"spec": {
			Properties: map[string]apiextensions.JSONSchemaProps{
				"field":     {Type: "string"},
				"external":  {Type: "string"},
				"unrelated": {Type: "string"},
				"projectRef": {
					Properties: map[string]apiextensions.JSONSchemaProps{
						"external":  {Type: "string"},
						"name":      {Type: "string"},
						"namespace": {Type: "string"},
					},
					Type: "object",
				},
				"obj": {
					Properties: map[string]apiextensions.JSONSchemaProps{
						"field":    {Type: "string"},
						"external": {Type: "string"},
						"nestedObj": {
							Properties: map[string]apiextensions.JSONSchemaProps{
								"field":    {Type: "string"},
								"external": {Type: "string"},
							},
							Type: "object",
						},
						"nestedList": {
							Items: &apiextensions.JSONSchemaPropsOrArray{
								Schema: &apiextensions.JSONSchemaProps{Type: "string"},
							},
							Type: "array",
						},
					},
					Type: "object",
				},
				"objMap": {
					AdditionalProperties: &apiextensions.JSONSchemaPropsOrBool{
						Schema: &apiextensions.JSONSchemaProps{Type: "string"},
					},
					Type: "object",
				},
				"objMapSchemaless": {
					Type: "object",
				},
				"list": {
					Items: &apiextensions.JSONSchemaPropsOrArray{
						Schema: &apiextensions.JSONSchemaProps{Type: "string"},
					},
					Type: "array",
				},
			},
			Type: "object",
		},
	},
}

func TestOverlayManagedFieldsOntoState(t *testing.T) {
	tests := []struct {
		name             string
		spec             map[string]interface{}
		krmState         map[string]interface{}
		managedFields    *fieldpath.Set
		hierarchicalRefs []corekccv1alpha1.HierarchicalReference
		expected         map[string]interface{}
	}{
		{
			name: "use spec values for k8s-managed fields",
			spec: map[string]interface{}{
				"field": "k8s",
			},
			krmState: map[string]interface{}{
				"field": "external",
			},
			managedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
				"f:field": emptyObject,
			}),
			expected: map[string]interface{}{
				"field": "k8s",
			},
		},
		{
			name: "use state values for externally-managed fields",
			spec: map[string]interface{}{
				"field": "k8s",
			},
			krmState: map[string]interface{}{
				"field": "external",
			},
			managedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
				"f:unrelated": emptyObject,
			}),
			expected: map[string]interface{}{
				"field": "external",
			},
		},
		{
			name: "use spec values for k8s-managed nested fields",
			spec: map[string]interface{}{
				"obj": map[string]interface{}{
					"field": "k8s",
				},
			},
			krmState: map[string]interface{}{
				"obj": map[string]interface{}{
					"field": "external",
				},
			},
			managedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
				"f:obj": map[string]interface{}{
					".":       emptyObject,
					"f:field": emptyObject,
				},
			}),
			expected: map[string]interface{}{
				"obj": map[string]interface{}{
					"field": "k8s",
				},
			},
		},
		{
			name: "use state value for externally-managed nested fields",
			spec: map[string]interface{}{
				"obj": map[string]interface{}{
					"field": "k8s",
				},
			},
			krmState: map[string]interface{}{
				"obj": map[string]interface{}{
					"field": "external",
				},
			},
			managedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
				"f:unrelated": emptyObject,
			}),
			expected: map[string]interface{}{
				"obj": map[string]interface{}{
					"field": "external",
				},
			},
		},
		{
			name: "top-level mixed management fields are merged",
			spec: map[string]interface{}{
				"field":    "k8s",
				"external": "k8s",
			},
			krmState: map[string]interface{}{
				"field":    "external",
				"external": "external",
			},
			managedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
				"f:field": emptyObject,
			}),
			expected: map[string]interface{}{
				"field":    "k8s",
				"external": "external",
			},
		},
		{
			name: "nested mixed management fields are merged",
			spec: map[string]interface{}{
				"obj": map[string]interface{}{
					"field":    "k8s",
					"external": "k8s",
				},
			},
			krmState: map[string]interface{}{
				"obj": map[string]interface{}{
					"field":    "external",
					"external": "external",
				},
			},
			managedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
				"f:obj": map[string]interface{}{
					".":       emptyObject,
					"f:field": emptyObject,
				},
			}),
			expected: map[string]interface{}{
				"obj": map[string]interface{}{
					"field":    "k8s",
					"external": "external",
				},
			},
		},
		{
			name: "object map fields are merged",
			spec: map[string]interface{}{
				"objMap": map[string]interface{}{
					"field":    "k8s",
					"external": "k8s",
				},
			},
			krmState: map[string]interface{}{
				"objMap": map[string]interface{}{
					"field":    "external",
					"external": "external",
				},
			},
			managedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
				"f:objMap": map[string]interface{}{
					"f:field": emptyObject,
				},
			}),
			expected: map[string]interface{}{
				"objMap": map[string]interface{}{
					"field":    "k8s",
					"external": "external",
				},
			},
		},
		{
			name: "schemaless object map fields are merged",
			spec: map[string]interface{}{
				"objMapSchemaless": map[string]interface{}{
					"field":    "k8s",
					"external": "k8s",
				},
			},
			krmState: map[string]interface{}{
				"objMapSchemaless": map[string]interface{}{
					"field":    "external",
					"external": "external",
				},
			},
			managedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
				"f:objMapSchemaless": map[string]interface{}{
					"f:field": emptyObject,
				},
			}),
			expected: map[string]interface{}{
				"objMapSchemaless": map[string]interface{}{
					"field":    "k8s",
					"external": "external",
				},
			},
		},

		{
			// TODO(b/160160236): Externally-managed list merging is not yet
			//  supported.
			name: "always use k8s value for lists set in spec",
			spec: map[string]interface{}{
				"list": []interface{}{"k8s-first", "k8s-second"},
			},
			krmState: map[string]interface{}{
				"list": []interface{}{"external-first", "external-second"},
			},
			managedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
				"f:unrelated": emptyObject,
			}),
			expected: map[string]interface{}{
				"list": []interface{}{"k8s-first", "k8s-second"},
			},
		},
		{
			name: "use external value for lists not set in spec",
			spec: map[string]interface{}{
				"field": "k8s",
			},
			krmState: map[string]interface{}{
				"field": "external",
				"list":  []interface{}{"external-first", "external-second"},
			},
			managedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
				"f:field": emptyObject,
			}),
			expected: map[string]interface{}{
				"field": "k8s",
				"list":  []interface{}{"external-first", "external-second"},
			},
		},
		{
			// name-only acquisition case
			name: "empty spec and managed fields are supported",
			spec: nil,
			krmState: map[string]interface{}{
				"field": "external",
			},
			managedFields: nil,
			expected: map[string]interface{}{
				"field": "external",
			},
		},
		{
			name: "empty state is supported",
			spec: map[string]interface{}{
				"field": "k8s",
			},
			krmState: nil,
			managedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
				"f:field": emptyObject,
			}),
			expected: map[string]interface{}{
				"field": "k8s",
			},
		},
		{
			name: "externally-managed fields can be cleared",
			spec: map[string]interface{}{
				"field":     "external",
				"unrelated": "val",
			},
			krmState: map[string]interface{}{
				"unrelated": "val",
			},
			managedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
				"f:unrelated": emptyObject,
			}),
			expected: map[string]interface{}{
				"unrelated": "val",
			},
		},
		{
			name: "hierarchical reference is preserved",
			spec: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"external": "project_id",
				},
				"unrelated": "val",
			},
			krmState: map[string]interface{}{
				"unrelated": "val",
			},
			managedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
				"f:unrelated": emptyObject,
				"f:projectRef": map[string]interface{}{
					".":          emptyObject,
					"f:external": emptyObject,
				},
			}),
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "projectRef",
				},
			},
			expected: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"external": "project_id",
				},
				"unrelated": "val",
			},
		},
		{
			name: "hierarchical reference is preserved even if not in managed fields as long as in spec",
			spec: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"external": "project_id",
				},
				"unrelated": "val",
			},
			krmState: map[string]interface{}{
				"unrelated": "val",
			},
			managedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
				"f:unrelated": emptyObject,
			}),
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "projectRef",
				},
			},
			expected: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"external": "project_id",
				},
				"unrelated": "val",
			},
		},
		{
			name: "no hierarchical reference in output config if none in spec",
			spec: map[string]interface{}{
				"unrelated": "val",
			},
			krmState: map[string]interface{}{
				"unrelated": "val",
			},
			managedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
				"f:unrelated": emptyObject,
			}),
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "projectRef",
				},
			},
			expected: map[string]interface{}{
				"unrelated": "val",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			spec := deepcopy.MapStringInterface(tc.spec)
			res, err := k8s.OverlayManagedFieldsOntoState(spec, tc.krmState, tc.managedFields, schema, tc.hierarchicalRefs)
			if err != nil {
				t.Error("error overlaying externally-managed fields:", err)
				return
			}
			if !test.Equals(t, res, tc.expected) {
				t.Errorf("actual: %+v, expected: %+v", res, tc.expected)
				return
			}
		})
	}
}

func TestConstructTrimmedSpecWithManagedFields(t *testing.T) {
	tests := []struct {
		name             string
		resource         *k8s.Resource
		hierarchicalRefs []corekccv1alpha1.HierarchicalReference
		expected         map[string]interface{}
	}{
		{
			name: "no managed field information present",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"field": "value",
					"list":  []interface{}{"a", "b"},
				},
			},
			expected: map[string]interface{}{
				"field": "value",
				"list":  []interface{}{"a", "b"},
			},
		},
		{
			name: "preserve k8s managed fields on the top level",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"field":    "k8s",
					"external": "external",
				},
				ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
					"f:field": emptyObject,
				}),
			},
			expected: map[string]interface{}{
				"field": "k8s",
			},
		},
		{
			name: "preserve k8s-managed nested fields",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"obj": map[string]interface{}{
						"field": "k8s",
						"nestedObj": map[string]interface{}{
							"field":    "k8s",
							"external": "external",
						},
						"external": "external",
					},
				},
				ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
					"f:obj": map[string]interface{}{
						".":       emptyObject,
						"f:field": emptyObject,
						"f:nestedObj": map[string]interface{}{
							".":       emptyObject,
							"f:field": emptyObject,
						},
					},
				}),
			},
			expected: map[string]interface{}{
				"obj": map[string]interface{}{
					"field": "k8s",
					"nestedObj": map[string]interface{}{
						"field": "k8s",
					},
				},
			},
		},
		{
			// TODO(b/160160236): Externally-managed list merging is not yet
			//  supported.
			name: "always use k8s value for lists set in spec",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"list": []interface{}{"k8s-first", "k8s-second"},
				},
				ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
					"f:unrelated": emptyObject,
				}),
			},
			expected: map[string]interface{}{
				"list": []interface{}{"k8s-first", "k8s-second"},
			},
		},
		{
			name:     "empty spec and managed fields are supported",
			resource: &k8s.Resource{},
			expected: nil,
		},
		{
			name: "hierarchical reference is preserved",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"projectRef": map[string]interface{}{
						"external": "project_id",
					},
					"unrelated": "val",
				},
				ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
					"f:unrelated": emptyObject,
					"f:projectRef": map[string]interface{}{
						".":          emptyObject,
						"f:external": emptyObject,
					},
				}),
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "projectRef",
				},
			},
			expected: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"external": "project_id",
				},
				"unrelated": "val",
			},
		},
		{
			name: "hierarchical reference is preserved even if not in managed fields as long as in spec",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"projectRef": map[string]interface{}{
						"external": "project_id",
					},
					"unrelated": "val",
				},
				ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
					"f:unrelated": emptyObject,
				}),
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "projectRef",
				},
			},
			expected: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"external": "project_id",
				},
				"unrelated": "val",
			},
		},
		{
			name: "no hierarchical reference in output config if none in spec",
			resource: &k8s.Resource{
				Spec: map[string]interface{}{
					"unrelated": "val",
				},
				ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
					"f:unrelated": emptyObject,
				}),
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "projectRef",
				},
			},
			expected: map[string]interface{}{
				"unrelated": "val",
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			trimmedSpec, err := k8s.ConstructTrimmedSpecWithManagedFields(tc.resource, schema, tc.hierarchicalRefs)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got, want := trimmedSpec, tc.expected; !test.Equals(t, got, want) {
				t.Fatalf("got: %v, want: %v", got, want)
			}
		})
	}
}

func fieldPathSetToJSON(t *testing.T, s *fieldpath.Set) []byte {
	b, err := s.ToJSON()
	if err != nil {
		t.Fatal("error converting set to JSON:", err)
	}
	return b
}

func mapToManagedFieldEntry(t *testing.T, manager string, fields map[string]interface{}) v1.ManagedFieldsEntry {
	b, err := json.Marshal(fields)
	if err != nil {
		t.Fatal("error marshaling to JSON:", err)
	}
	return v1.ManagedFieldsEntry{
		Manager:    manager,
		FieldsType: k8s.ManagedFieldsTypeFieldsV1,
		FieldsV1:   &v1.FieldsV1{Raw: b},
	}
}

func TestSanitizeManagedFields(t *testing.T) {
	tests := []struct {
		name                string
		managedFieldEntries []v1.ManagedFieldsEntry
		expectedEntries     []v1.ManagedFieldsEntry
	}{
		{
			name: "remove subresource from spec managed field",
			managedFieldEntries: []v1.ManagedFieldsEntry{
				{
					Manager:     "manager",
					FieldsType:  k8s.ManagedFieldsTypeFieldsV1,
					FieldsV1:    &v1.FieldsV1{Raw: []byte(`{"f:spec":{}}`)},
					Subresource: "status",
				},
			},
			expectedEntries: []v1.ManagedFieldsEntry{
				{
					Manager:     "manager",
					FieldsType:  k8s.ManagedFieldsTypeFieldsV1,
					FieldsV1:    &v1.FieldsV1{Raw: []byte(`{"f:spec":{}}`)},
					Subresource: "",
				},
			},
		},
		{
			name: "keep subresource for non-spec managed field",
			managedFieldEntries: []v1.ManagedFieldsEntry{
				{
					Manager:     "manager",
					FieldsType:  k8s.ManagedFieldsTypeFieldsV1,
					FieldsV1:    &v1.FieldsV1{Raw: []byte(`{"f:status":{}}`)},
					Subresource: "status",
				},
			},
			expectedEntries: []v1.ManagedFieldsEntry{
				{
					Manager:     "manager",
					FieldsType:  k8s.ManagedFieldsTypeFieldsV1,
					FieldsV1:    &v1.FieldsV1{Raw: []byte(`{"f:status":{}}`)},
					Subresource: "status",
				},
			},
		},
		{
			name: "no change if subresource is empty",
			managedFieldEntries: []v1.ManagedFieldsEntry{
				{
					Manager:     "manager",
					FieldsType:  k8s.ManagedFieldsTypeFieldsV1,
					FieldsV1:    &v1.FieldsV1{Raw: []byte(`{"f:spec":{}}`)},
					Subresource: "",
				},
			},
			expectedEntries: []v1.ManagedFieldsEntry{
				{
					Manager:     "manager",
					FieldsType:  k8s.ManagedFieldsTypeFieldsV1,
					FieldsV1:    &v1.FieldsV1{Raw: []byte(`{"f:spec":{}}`)},
					Subresource: "",
				},
			},
		},
		{
			name: "remove subresource from mixed spec and metadata managed field",
			managedFieldEntries: []v1.ManagedFieldsEntry{
				{
					Manager:    "kubectl",
					FieldsType: k8s.ManagedFieldsTypeFieldsV1,
					FieldsV1: &v1.FieldsV1{Raw: []byte(`{
						"f:metadata": {
							"f:annotations": {
								".": {},
								"f:cnrm.cloud.google.com/deletion-policy": {},
								"f:kubectl.kubernetes.io/last-applied-configuration": {}
							}
						},
						"f:spec": {
							".": {},
							"f:attemptDeadline": {},
							"f:description": {}
						}
					}`)},
					Subresource: "status",
				},
			},
			expectedEntries: []v1.ManagedFieldsEntry{
				{
					Manager:    "kubectl",
					FieldsType: k8s.ManagedFieldsTypeFieldsV1,
					FieldsV1: &v1.FieldsV1{Raw: []byte(`{
						"f:metadata": {
							"f:annotations": {
								".": {},
								"f:cnrm.cloud.google.com/deletion-policy": {},
								"f:kubectl.kubernetes.io/last-applied-configuration": {}
							}
						},
						"f:spec": {
							".": {},
							"f:attemptDeadline": {},
							"f:description": {}
						}
					}`)},
					Subresource: "",
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := &k8s.Resource{
				ObjectMeta: v1.ObjectMeta{
					ManagedFields: tc.managedFieldEntries,
				},
			}
			k8s.SanitizeSpecManagedFields(r)
			if !reflect.DeepEqual(r.ObjectMeta.ManagedFields, tc.expectedEntries) {
				t.Errorf("actual: %+v, expected: %+v", r.ObjectMeta.ManagedFields, tc.expectedEntries)
			}
		})
	}
}
