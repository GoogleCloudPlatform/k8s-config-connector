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

package resourceoverrides

import (
	"reflect"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	testk8s "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/k8s"

	"github.com/google/go-cmp/cmp"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	fooKind     = "Foo"
	emptyObject = make(map[string]interface{})
)

func TestPreserveUserSpecifiedLegacyField(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		original   *k8s.Resource
		reconciled *k8s.Resource
		fieldPath  []string
		expected   *k8s.Resource
	}{
		{
			name: "no legacy field is set",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"field2": "defaultValue",
				},
			},
			fieldPath: []string{"legacyField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"field2": "defaultValue",
				},
			},
		},
		{
			name: "user-specified legacy field is preserved",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":      "value",
					"legacyField": "value",
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"field2": "defaultValue",
				},
			},
			fieldPath: []string{"legacyField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":      "value",
					"legacyField": "value",
					"field2":      "defaultValue",
				},
			},
		},
		{
			name: "user-specified nested legacy field is preserved",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"legacyField": "value",
					},
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"newField": "value",
					},
				},
			},
			fieldPath: []string{"topField", "legacyField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"legacyField": "value",
						"newField":    "value",
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if err := PreserveUserSpecifiedLegacyField(tc.original, tc.reconciled, tc.fieldPath...); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(tc.reconciled, tc.expected) {
				t.Fatalf("unexpected diff (-want +got): \n%v", cmp.Diff(tc.expected, tc.reconciled))
			}
		})
	}
}

func TestPreserveUserSpecifiedLegacyFieldUnderSlice(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		original      *k8s.Resource
		reconciled    *k8s.Resource
		pathUpToSlice []string
		fieldPath     []string
		expected      *k8s.Resource
	}{
		{
			name: "no legacy field is set",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"field2": "defaultValue",
				},
			},
			pathUpToSlice: []string{"pathUpToSlice"},
			fieldPath:     []string{"legacyField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"field2": "defaultValue",
				},
			},
		},
		{
			name: "user-specified legacy field is preserved for the correct slice element",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"pathUpToSlice": []interface{}{
						map[string]interface{}{
							"subfield1": "value1",
						},
						map[string]interface{}{
							"legacyField": "value",
						},
					},
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"pathUpToSlice": []interface{}{
						map[string]interface{}{
							"subfield1": "value1",
						},
						map[string]interface{}{},
					},
					"field2": "defaultValue",
				},
			},
			pathUpToSlice: []string{"pathUpToSlice"},
			fieldPath:     []string{"legacyField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"pathUpToSlice": []interface{}{
						map[string]interface{}{
							"subfield1": "value1",
						},
						map[string]interface{}{
							"legacyField": "value",
						},
					},
					"field2": "defaultValue",
				},
			},
		},
		{
			name: "user-specified nested legacy field is preserved",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField1": map[string]interface{}{
						"pathUpToSlice": []interface{}{
							map[string]interface{}{
								"topField2": map[string]interface{}{
									"legacyField": "value",
								},
							},
						},
					},
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField1": map[string]interface{}{
						"pathUpToSlice": []interface{}{
							map[string]interface{}{
								"topField2": map[string]interface{}{
									"newField": "value1",
								},
							},
						},
					},
				},
			},
			pathUpToSlice: []string{"topField1", "pathUpToSlice"},
			fieldPath:     []string{"topField2", "legacyField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField1": map[string]interface{}{
						"pathUpToSlice": []interface{}{
							map[string]interface{}{
								"topField2": map[string]interface{}{
									"newField":    "value1",
									"legacyField": "value",
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if err := PreserveUserSpecifiedLegacyFieldUnderSlice(tc.original, tc.reconciled, tc.pathUpToSlice, tc.fieldPath); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(tc.reconciled, tc.expected) {
				t.Fatalf("unexpected diff (-want +got): \n%v", cmp.Diff(tc.expected, tc.reconciled))
			}
		})
	}
}

func TestPreserveUserSpecifiedLegacyArrayField(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		original   *k8s.Resource
		reconciled *k8s.Resource
		fieldPath  []string
		expected   *k8s.Resource
		hasError   bool
	}{
		{
			name: "no legacy array field is set",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"field2": "defaultValue",
				},
			},
			fieldPath: []string{"legacyArray"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"field2": "defaultValue",
				},
			},
		},
		{
			name: "user-specified legacy array field is preserved",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"legacyArray": []interface{}{
						"testValue1",
						"testValue2",
					},
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"field2": "defaultValue",
				},
			},
			fieldPath: []string{"legacyArray"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"field2": "defaultValue",
					"legacyArray": []interface{}{
						"testValue1",
						"testValue2",
					},
				},
			},
		},
		{
			name: "user-specified nested legacy array field is preserved",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"legacyArray": []interface{}{
							"testValue1",
							"testValue2",
						},
					},
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"newField": "value",
					},
				},
			},
			fieldPath: []string{"topField", "legacyArray"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"legacyArray": []interface{}{
							"testValue1",
							"testValue2",
						},
						"newField": "value",
					},
				},
			},
		},
		{
			name: "user-specified legacy field is not preserved because it's " +
				"not an array",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":      "value",
					"legacyArray": "testValue",
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"field2": "defaultValue",
				},
			},
			fieldPath: []string{"legacyArray"},
			hasError:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := PreserveUserSpecifiedLegacyArrayField(tc.original, tc.reconciled, tc.fieldPath...)
			if tc.hasError {
				if err == nil {
					t.Fatalf("got nil, but expect to have error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got, want := tc.reconciled, tc.expected; !reflect.DeepEqual(got, want) {
				t.Fatalf("unexpected diff (-want +got): \n%v", cmp.Diff(want, got))
			}
		})
	}
}

func TestPruneDefaultedAuthoritativeFieldIfOnlyLegacyFieldSpecified(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		original        *k8s.Resource
		reconciled      *k8s.Resource
		legacyFieldPath []string
		fieldPath       []string
		expected        *k8s.Resource
	}{
		{
			name: "neither the legacy field nor the authoritative field is set; the authoritative field is not defaulted",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"field2": "defaultValue",
				},
			},
			legacyFieldPath: []string{"legacyField"},
			fieldPath:       []string{"authoritativeField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"field2": "defaultValue",
				},
			},
		},
		{
			name: "neither the legacy field nor the authoritative field is set; the authoritative field is defaulted",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":             "value",
					"authoritativeField": "defaultValue",
				},
			},
			legacyFieldPath: []string{"legacyField"},
			fieldPath:       []string{"authoritativeField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":             "value",
					"authoritativeField": "defaultValue",
				},
			},
		},
		{
			name: "prune the authoritative field if users only specify the legacy field",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":      "value",
					"legacyField": "value",
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":             "value",
					"authoritativeField": "value",
				},
			},
			legacyFieldPath: []string{"legacyField"},
			fieldPath:       []string{"authoritativeField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
				},
			},
		},
		{
			name: "prune the nested authoritative field if users only specify the legacy field",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"field2":      "value",
						"legacyField": "value",
					},
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"field2":             "value",
						"authoritativeField": "value",
					},
				},
			},
			legacyFieldPath: []string{"topField", "legacyField"},
			fieldPath:       []string{"topField", "authoritativeField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"field2": "value",
					},
				},
			},
		},
		{
			name: "users only specify the authoritative field",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":             "value",
					"authoritativeField": "value",
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":             "value",
					"authoritativeField": "value",
				},
			},
			legacyFieldPath: []string{"legacyField"},
			fieldPath:       []string{"authoritativeField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":             "value",
					"authoritativeField": "value",
				},
			},
		},
		{
			name: "users specify both the legacy and the authoritative fields",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":             "value",
					"legacyField":        "value",
					"authoritativeField": "value",
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":             "value",
					"legacyField":        "value",
					"authoritativeField": "value",
				},
			},
			legacyFieldPath: []string{"legacyField"},
			fieldPath:       []string{"authoritativeField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":             "value",
					"legacyField":        "value",
					"authoritativeField": "value",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if err := PruneDefaultedAuthoritativeFieldIfOnlyLegacyFieldSpecified(tc.original, tc.reconciled, tc.legacyFieldPath, tc.fieldPath); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(tc.reconciled, tc.expected) {
				t.Fatalf("unexpected diff (-want +got): \n%v", cmp.Diff(tc.expected, tc.reconciled))
			}
		})
	}
}

func TestPruneDefaultedAuthoritativeFieldIfOnlyLegacyFieldSpecifiedUnderSlice(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                   string
		original               *k8s.Resource
		reconciled             *k8s.Resource
		pathUpToSlice          []string
		legacyFieldPath        []string
		authoritativeFieldPath []string
		expected               *k8s.Resource
	}{
		{
			name: "neither the legacy field nor the authoritative field is set; the authoritative field is not defaulted",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"field2": "defaultValue",
				},
			},
			pathUpToSlice:          []string{"pathUpToSlice"},
			legacyFieldPath:        []string{"legacyField"},
			authoritativeFieldPath: []string{"authoritativeField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"field2": "defaultValue",
				},
			},
		},
		{
			name: "neither the legacy field nor the authoritative field is set; the authoritative field is defaulted",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"pathUpToSlice": []interface{}{
						map[string]interface{}{
							"authoritativeField": "defaultValue",
						},
					},
				},
			},
			pathUpToSlice:          []string{"pathUpToSlice"},
			legacyFieldPath:        []string{"legacyField"},
			authoritativeFieldPath: []string{"authoritativeField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"pathUpToSlice": []interface{}{
						map[string]interface{}{
							"authoritativeField": "defaultValue",
						},
					},
				},
			},
		},
		{
			name: "prune the authoritative field if users only specify the legacy field",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"pathUpToSlice": []interface{}{
						map[string]interface{}{
							"legacyField": "value",
						},
					},
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"pathUpToSlice": []interface{}{
						map[string]interface{}{
							"authoritativeField": "value",
						},
					},
				},
			},
			pathUpToSlice:          []string{"pathUpToSlice"},
			legacyFieldPath:        []string{"legacyField"},
			authoritativeFieldPath: []string{"authoritativeField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"pathUpToSlice": []interface{}{
						map[string]interface{}{},
					},
				},
			},
		},
		{
			name: "prune the nested authoritative field if users only specify the legacy field",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField1": map[string]interface{}{
						"pathUpToSlice": []interface{}{
							map[string]interface{}{
								"topField2": map[string]interface{}{
									"legacyField": "value",
									"field2":      "value",
								},
							},
						},
					},
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField1": map[string]interface{}{
						"pathUpToSlice": []interface{}{
							map[string]interface{}{
								"topField2": map[string]interface{}{
									"authoritativeField": "value",
									"field2":             "value",
								},
							},
						},
					},
				},
			},
			pathUpToSlice:          []string{"topField1", "pathUpToSlice"},
			legacyFieldPath:        []string{"topField2", "legacyField"},
			authoritativeFieldPath: []string{"topField2", "authoritativeField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField1": map[string]interface{}{
						"pathUpToSlice": []interface{}{
							map[string]interface{}{
								"topField2": map[string]interface{}{
									"field2": "value",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "users only specify the authoritative field",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"pathUpToSlice": []interface{}{
						map[string]interface{}{
							"authoritativeField": "value",
						},
					},
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"pathUpToSlice": []interface{}{
						map[string]interface{}{
							"authoritativeField": "value",
						},
					},
				},
			},
			pathUpToSlice:          []string{"pathUpToSlice"},
			legacyFieldPath:        []string{"legacyField"},
			authoritativeFieldPath: []string{"authoritativeField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"pathUpToSlice": []interface{}{
						map[string]interface{}{
							"authoritativeField": "value",
						},
					},
				},
			},
		},
		{
			name: "handle multiple elements with different use-cases",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"pathUpToSlice": []interface{}{
						map[string]interface{}{
							"legacyField": "value",
						},
						map[string]interface{}{
							"authoritativeField": "value1",
						},
					},
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"pathUpToSlice": []interface{}{
						map[string]interface{}{
							"authoritativeField": "value",
						},
						map[string]interface{}{
							"authoritativeField": "value1",
						},
						map[string]interface{}{
							"authoritativeField": "defaultValue",
						},
					},
				},
			},
			pathUpToSlice:          []string{"pathUpToSlice"},
			legacyFieldPath:        []string{"legacyField"},
			authoritativeFieldPath: []string{"authoritativeField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"pathUpToSlice": []interface{}{
						map[string]interface{}{},
						map[string]interface{}{
							"authoritativeField": "value1",
						},
						map[string]interface{}{
							"authoritativeField": "defaultValue",
						},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if err := PruneDefaultedAuthoritativeFieldIfOnlyLegacyFieldSpecifiedUnderSlice(tc.original, tc.reconciled, tc.pathUpToSlice, tc.legacyFieldPath, tc.authoritativeFieldPath); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(tc.reconciled, tc.expected) {
				t.Fatalf("unexpected diff (-want +got): \n%v", cmp.Diff(tc.expected, tc.reconciled))
			}
		})
	}
}

func TestPruneDefaultedAuthoritativeArrayFieldIfOnlyLegacyArrayFieldSpecified(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		original        *k8s.Resource
		reconciled      *k8s.Resource
		legacyFieldPath []string
		fieldPath       []string
		expected        *k8s.Resource
		hasError        bool
	}{
		{
			name: "neither the legacy array field nor the authoritative array " +
				"field is set; the authoritative array field is not defaulted",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"field2": "defaultValue",
				},
			},
			legacyFieldPath: []string{"legacyArray"},
			fieldPath:       []string{"authoritativeArray"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"field2": "defaultValue",
				},
			},
		},
		{
			name: "neither the legacy array field nor the authoritative array " +
				"field is set; the authoritative array field is defaulted",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"authoritativeArray": []interface{}{
						"defaultedValue1",
						"defaultedValue2",
					},
				},
			},
			legacyFieldPath: []string{"legacyArray"},
			fieldPath:       []string{"authoritativeArray"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"authoritativeArray": []interface{}{
						"defaultedValue1",
						"defaultedValue2",
					},
				},
			},
		},
		{
			name: "prune the authoritative array field if users only specify " +
				"the legacy array field",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"legacyArray": []interface{}{
						"testValue1",
						"testValue2",
					},
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"authoritativeArray": []interface{}{
						"testValue1",
						"testValue2",
					},
				},
			},
			legacyFieldPath: []string{"legacyArray"},
			fieldPath:       []string{"authoritativeArray"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
				},
			},
		},
		{
			name: "prune the authoritative array field of objects if users " +
				"only specify the legacy array field of strings",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"legacyArray": []interface{}{
						"testValue1",
						"testValue2",
					},
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"authoritativeArray": []interface{}{
						map[string]interface{}{
							"external": "testValue1",
						},
						map[string]interface{}{
							"external": "testValue2",
						},
					},
				},
			},
			legacyFieldPath: []string{"legacyArray"},
			fieldPath:       []string{"authoritativeArray"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
				},
			},
		},
		{
			name: "prune the nested authoritative array field if users only " +
				"specify the legacy array field",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"field2": "value",
						"legacyArray": []interface{}{
							"testValue1",
							"testValue2",
						},
					},
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"field2": "value",
						"authoritativeArray": []interface{}{
							"testValue1",
							"testValue2",
						},
					},
				},
			},
			legacyFieldPath: []string{"topField", "legacyArray"},
			fieldPath:       []string{"topField", "authoritativeArray"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"field2": "value",
					},
				},
			},
		},
		{
			name: "users only specify the authoritative array field",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"authoritativeArray": []interface{}{
						"testValue1",
						"testValue2",
					},
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"authoritativeArray": []interface{}{
						"testValue1",
						"testValue2",
					},
				},
			},
			legacyFieldPath: []string{"legacyArray"},
			fieldPath:       []string{"authoritativeArray"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"authoritativeArray": []interface{}{
						"testValue1",
						"testValue2",
					},
				},
			},
		},
		{
			name: "users specify both the legacy and the authoritative fields",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"legacyArray": []interface{}{
						"testValue1",
						"testValue2",
					},
					"authoritativeArray": []interface{}{
						"testValue1",
						"testValue2",
					},
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"legacyArray": []interface{}{
						"testValue1",
						"testValue2",
					},
					"authoritativeArray": []interface{}{
						"testValue1",
						"testValue2",
					},
				},
			},
			legacyFieldPath: []string{"legacyArray"},
			fieldPath:       []string{"authoritativeArray"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"legacyArray": []interface{}{
						"testValue1",
						"testValue2",
					},
					"authoritativeArray": []interface{}{
						"testValue1",
						"testValue2",
					},
				},
			},
		},
		{
			name: "error pruning the authoritative array field if the legacy " +
				"field is not an array",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":      "value",
					"legacyField": "testValue1",
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"authoritativeArray": []interface{}{
						"testValue1",
						"testValue2",
					},
				},
			},
			legacyFieldPath: []string{"legacyField"},
			fieldPath:       []string{"authoritativeArray"},
			hasError:        true,
		},
		{
			name: "error when users only specify the non-array authoritative " +
				"field",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":             "value",
					"authoritativeField": "testValue1",
				},
			},
			reconciled: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":             "value",
					"authoritativeField": "testValue1",
				},
			},
			legacyFieldPath: []string{"legacyArray"},
			fieldPath:       []string{"authoritativeField"},
			hasError:        true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := PruneDefaultedAuthoritativeArrayFieldIfOnlyLegacyArrayFieldSpecified(tc.original, tc.reconciled, tc.legacyFieldPath, tc.fieldPath)
			if tc.hasError {
				if err == nil {
					t.Fatalf("got nil, but expect to have error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got, want := tc.reconciled, tc.expected; !reflect.DeepEqual(got, want) {
				t.Fatalf("unexpected diff (-want +got): \n%v", cmp.Diff(want, got))
			}
		})
	}
}

func TestFavorAuthoritativeFieldOverLegacyField(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		original        *k8s.Resource
		legacyFieldPath []string
		fieldPath       []string
		expected        *k8s.Resource
		hasError        bool
	}{
		{
			name: "neither the legacy field nor the authoritative field is set",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
				},
			},
			legacyFieldPath: []string{"legacyField"},
			fieldPath:       []string{"authoritativeField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
				},
			},
		},
		{
			name: "only the legacy field is set",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":      "value",
					"legacyField": "value",
				},
				ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
					"f:field1":      emptyObject,
					"f:legacyField": emptyObject,
				}),
			},
			legacyFieldPath: []string{"legacyField"},
			fieldPath:       []string{"authoritativeField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":             "value",
					"authoritativeField": "value",
				},
				ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
					"f:field1":             emptyObject,
					"f:legacyField":        emptyObject,
					"f:authoritativeField": emptyObject,
				}),
			},
		},
		{
			name: "only the nested legacy field is set",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"field2":      "value",
						"legacyField": "value",
					},
				},
				ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
					"f:field1": emptyObject,
					"f:topField": map[string]interface{}{
						".":             emptyObject,
						"f:legacyField": emptyObject,
					},
				}),
			},
			legacyFieldPath: []string{"topField", "legacyField"},
			fieldPath:       []string{"topField", "authoritativeField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"field2":             "value",
						"authoritativeField": "value",
					},
				},
				ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
					"f:field1": emptyObject,
					"f:topField": map[string]interface{}{
						".":                    emptyObject,
						"f:legacyField":        emptyObject,
						"f:authoritativeField": emptyObject,
					},
				}),
			},
		},
		{
			name: "only the nested legacy field is set and the authoritative field is a reference field",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"field2":      "value",
						"legacyField": "value",
					},
				},
				ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
					"f:field1": emptyObject,
					"f:topField": map[string]interface{}{
						".":             emptyObject,
						"f:legacyField": emptyObject,
					},
				}),
			},
			legacyFieldPath: []string{"topField", "legacyField"},
			fieldPath:       []string{"topField", "authoritativeFieldRef"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"field2": "value",
						"authoritativeFieldRef": map[string]interface{}{
							"external": "value",
						},
					},
				},
				ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
					"f:field1": emptyObject,
					"f:topField": map[string]interface{}{
						".":             emptyObject,
						"f:legacyField": emptyObject,
						"f:authoritativeFieldRef": map[string]interface{}{
							".":          emptyObject,
							"f:external": emptyObject,
						},
					},
				}),
			},
		},
		{
			name: "only the authoritative field is set",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":             "value",
					"authoritativeField": "value",
				},
				ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
					"f:field1":             emptyObject,
					"f:authoritativeField": emptyObject,
				}),
			},
			legacyFieldPath: []string{"legacyField"},
			fieldPath:       []string{"authoritativeField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":             "value",
					"authoritativeField": "value",
				},
				ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
					"f:field1":             emptyObject,
					"f:authoritativeField": emptyObject,
				}),
			},
		},
		{
			name: "only the nested authoritative field is set",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"field2":             "value",
						"authoritativeField": "value",
					},
				},
				ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
					"f:field1": emptyObject,
					"f:topField": map[string]interface{}{
						".":                    emptyObject,
						"f:authoritativeField": emptyObject,
					},
				}),
			},
			legacyFieldPath: []string{"topField", "legacyField"},
			fieldPath:       []string{"topField", "authoritativeField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"field2":             "value",
						"authoritativeField": "value",
					},
				},
				ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
					"f:field1": emptyObject,
					"f:topField": map[string]interface{}{
						".":                    emptyObject,
						"f:authoritativeField": emptyObject,
					},
				}),
			},
		},
		{
			name: "both the legacy field and the authoritative field are set with the same value",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":             "value",
					"legacyField":        "value1",
					"authoritativeField": "value1",
				},
				ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
					"f:field1":             emptyObject,
					"f:legacyField":        emptyObject,
					"f:authoritativeField": emptyObject,
				}),
			},
			legacyFieldPath: []string{"legacyField"},
			fieldPath:       []string{"authoritativeField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":             "value",
					"authoritativeField": "value1",
				},
				ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
					"f:field1":             emptyObject,
					"f:legacyField":        emptyObject,
					"f:authoritativeField": emptyObject,
				}),
			},
		},
		{
			name: "both the legacy field and the authoritative field are set with the different values",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":             "value",
					"legacyField":        "value1",
					"authoritativeField": "value2",
				},
				ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
					"f:field1":             emptyObject,
					"f:legacyField":        emptyObject,
					"f:authoritativeField": emptyObject,
				}),
			},
			legacyFieldPath: []string{"legacyField"},
			fieldPath:       []string{"authoritativeField"},
			hasError:        true,
		},
		{
			name: "both the nested legacy field and the nested authoritative field are set with the same value",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"legacyField":        "value1",
						"authoritativeField": "value1",
					},
				},
				ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
					"f:field1": emptyObject,
					"f:topField": map[string]interface{}{
						".":                    emptyObject,
						"f:authoritativeField": emptyObject,
						"f:legacyField":        emptyObject,
					},
				}),
			},
			legacyFieldPath: []string{"topField", "legacyField"},
			fieldPath:       []string{"topField", "authoritativeField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"authoritativeField": "value1",
					},
				},
				ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
					"f:field1": emptyObject,
					"f:topField": map[string]interface{}{
						".":                    emptyObject,
						"f:authoritativeField": emptyObject,
						"f:legacyField":        emptyObject,
					},
				}),
			},
		},
		{
			name: "both the nested legacy field and the nested authoritative field are set with different values",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"legacyField":        "value1",
						"authoritativeField": "value2",
					},
				},
				ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
					"f:field1": emptyObject,
					"f:topField": map[string]interface{}{
						".":                    emptyObject,
						"f:authoritativeField": emptyObject,
						"f:legacyField":        emptyObject,
					},
				}),
			},
			legacyFieldPath: []string{"topField", "legacyField"},
			fieldPath:       []string{"topField", "authoritativeField"},
			hasError:        true,
		},
		{
			name: "the nested legacy field and the nested authoritative field are structurally different",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"legacyField": "value1",
						"authoritativeField": map[string]interface{}{
							"subfield": "value1",
						},
					},
				},
			},
			legacyFieldPath: []string{"topField", "legacyField"},
			fieldPath:       []string{"topField", "authoritativeField"},
			hasError:        true,
		},
		{
			name: "the nested authoritative field has a slice somewhere in the path",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":      "value",
					"legacyField": "value1",
					"topField": []interface{}{
						map[string]interface{}{
							"authoritativeField": "value2",
						},
					},
				},
			},
			legacyFieldPath: []string{"legacyField"},
			fieldPath:       []string{"topField", "authoritativeField"},
			hasError:        true,
		},
		{
			name: "the nested legacy field has a slice somewhere in the path",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":             "value",
					"authoritativeField": "value1",
					"topField": []interface{}{
						map[string]interface{}{
							"legacyField": "value2",
						},
					},
				},
			},
			legacyFieldPath: []string{"topField", "legacyField"},
			fieldPath:       []string{"authoritativeField"},
			hasError:        true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := FavorAuthoritativeFieldOverLegacyField(tc.original, tc.legacyFieldPath, tc.fieldPath)
			if tc.hasError {
				if err == nil {
					t.Fatalf("got nil, but expect to have error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Compare ManagedFields separately.
			if tc.original.ManagedFields != nil && !tc.original.ManagedFields.Equals(tc.expected.ManagedFields) {
				t.Fatalf("got %v, want %v", tc.original.ManagedFields, tc.expected.ManagedFields)
			}
			tc.original.ManagedFields = nil
			tc.expected.ManagedFields = nil
			if !reflect.DeepEqual(tc.original, tc.expected) {
				t.Fatalf("unexpected diff (-want +got): \n%v", cmp.Diff(tc.expected, tc.original))
			}
		})
	}
}

func TestFavorReferenceFieldOverNonReferenceFieldUnderSlice(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                  string
		original              *k8s.Resource
		pathUpToSlice         []string
		nonReferenceFieldPath []string
		referenceFieldPath    []string
		expected              *k8s.Resource
		hasError              bool
	}{
		{
			name: "neither the legacy field nor the authoritative field is set without path up to slice",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
				},
			},
			pathUpToSlice:         []string{"pathUpToSlice"},
			nonReferenceFieldPath: []string{"legacyField"},
			referenceFieldPath:    []string{"authoritativeField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
				},
			},
		},
		{
			name: "neither the legacy field nor the authoritative field is set with path up to slice",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":        "value",
					"pathUpToSlice": []interface{}{},
				},
			},
			pathUpToSlice:         []string{"pathUpToSlice"},
			nonReferenceFieldPath: []string{"legacyField"},
			referenceFieldPath:    []string{"authoritativeField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":        "value",
					"pathUpToSlice": []interface{}{},
				},
			},
		},
		{
			name: "only the legacy field is set",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"pathUpToSlice": []interface{}{
						map[string]interface{}{
							"field1":      "value",
							"legacyField": "value",
						},
					},
				},
			},
			pathUpToSlice:         []string{"pathUpToSlice"},
			nonReferenceFieldPath: []string{"legacyField"},
			referenceFieldPath:    []string{"authoritativeField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"pathUpToSlice": []interface{}{
						map[string]interface{}{
							"field1": "value",
							"authoritativeField": map[string]interface{}{
								"external": "value",
							},
						},
					},
				},
			},
		},
		{
			name: "only the authoritative field is set",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"pathUpToSlice": []interface{}{
						map[string]interface{}{
							"authoritativeField": "value",
						},
					},
				},
			},
			pathUpToSlice:         []string{"pathUpToSlice"},
			nonReferenceFieldPath: []string{"legacyField"},
			referenceFieldPath:    []string{"authoritativeField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"pathUpToSlice": []interface{}{
						map[string]interface{}{
							"authoritativeField": "value",
						},
					},
				},
			},
		},
		{
			name: "only the legacy field within the nested slice field is set",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"field2": "value",
						"pathUpToSlice": []interface{}{
							map[string]interface{}{
								"legacyField": "value",
							},
						},
					},
				},
			},
			pathUpToSlice:         []string{"topField", "pathUpToSlice"},
			nonReferenceFieldPath: []string{"legacyField"},
			referenceFieldPath:    []string{"authoritativeField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"field2": "value",
						"pathUpToSlice": []interface{}{
							map[string]interface{}{
								"authoritativeField": map[string]interface{}{
									"external": "value",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "only the nested legacy field within the slice field is set",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"pathUpToSlice": []interface{}{
						map[string]interface{}{
							"topField": map[string]interface{}{
								"legacyField": "value",
							},
						},
					},
				},
			},
			pathUpToSlice:         []string{"pathUpToSlice"},
			nonReferenceFieldPath: []string{"topField", "legacyField"},
			referenceFieldPath:    []string{"topField", "authoritativeField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"pathUpToSlice": []interface{}{
						map[string]interface{}{
							"topField": map[string]interface{}{
								"authoritativeField": map[string]interface{}{
									"external": "value",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "slice has one authoritative field and one legacy field set on different elements",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"pathUpToSlice": []interface{}{
						map[string]interface{}{
							"authoritativeField": map[string]interface{}{
								"name": "resourcename",
								"kind": "resourcekind",
							},
						},
						map[string]interface{}{
							"legacyField": "value1",
						},
					},
				},
			},
			pathUpToSlice:         []string{"pathUpToSlice"},
			nonReferenceFieldPath: []string{"legacyField"},
			referenceFieldPath:    []string{"authoritativeField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"pathUpToSlice": []interface{}{
						map[string]interface{}{
							"authoritativeField": map[string]interface{}{
								"name": "resourcename",
								"kind": "resourcekind",
							},
						},
						map[string]interface{}{
							"authoritativeField": map[string]interface{}{
								"external": "value1",
							},
						},
					},
				},
			},
		},
		{
			name: "both legacy field and authoritative field are set",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"pathUpToSlice": []interface{}{
						map[string]interface{}{
							"legacyField": "value1",
							"authoritativeField": map[string]interface{}{
								"external": "value2",
							},
						},
					},
				},
			},
			pathUpToSlice:         []string{"pathUpToSlice"},
			nonReferenceFieldPath: []string{"legacyField"},
			referenceFieldPath:    []string{"authoritativeField"},
			hasError:              true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := FavorReferenceFieldOverNonReferenceFieldUnderSlice(tc.original, tc.pathUpToSlice, tc.nonReferenceFieldPath, tc.referenceFieldPath)
			if tc.hasError {
				if err == nil {
					t.Fatalf("got nil, but expect to have error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(tc.original, tc.expected) {
				t.Fatalf("unexpected diff (-want +got): \n%v", cmp.Diff(tc.expected, tc.original))
			}
		})
	}
}

func TestFavorReferenceArrayFieldOverNonReferenceArrayField(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                  string
		original              *k8s.Resource
		nonReferenceFieldPath []string
		referenceFieldPath    []string
		expected              *k8s.Resource
		hasError              bool
	}{
		{
			name: "neither the non-reference array field nor the reference " +
				"array field is set",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
				},
			},
			nonReferenceFieldPath: []string{"nonReferenceField"},
			referenceFieldPath:    []string{"referenceField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
				},
			},
		},
		{
			name: "only the non-reference array field is set",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"nonReferenceField": []interface{}{
						"testValue1",
						"testValue2",
					},
				},
			},
			nonReferenceFieldPath: []string{"nonReferenceField"},
			referenceFieldPath:    []string{"referenceField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"referenceField": []interface{}{
						map[string]interface{}{
							"external": "testValue1",
						},
						map[string]interface{}{
							"external": "testValue2",
						},
					},
				},
			},
		},
		{
			name: "only the nested non-reference array field is set",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"field2": "value",
						"nonReferenceField": []interface{}{
							"testValue1",
							"testValue2",
						},
					},
				},
			},
			nonReferenceFieldPath: []string{"topField", "nonReferenceField"},
			referenceFieldPath:    []string{"topField", "referenceField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"field2": "value",
						"referenceField": []interface{}{
							map[string]interface{}{
								"external": "testValue1",
							},
							map[string]interface{}{
								"external": "testValue2",
							},
						},
					},
				},
			},
		},
		{
			name: "only the reference array field is set",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"referenceField": []interface{}{
						map[string]interface{}{
							"name": "reference1",
						},
						map[string]interface{}{
							"external": "reference2",
						},
					},
				},
			},
			nonReferenceFieldPath: []string{"nonReferenceField"},
			referenceFieldPath:    []string{"referenceField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"referenceField": []interface{}{
						map[string]interface{}{
							"name": "reference1",
						},
						map[string]interface{}{
							"external": "reference2",
						},
					},
				},
			},
		},
		{
			name: "only the nested reference array field is set",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"field2": "value",
						"referenceField": []interface{}{
							map[string]interface{}{
								"name": "reference1",
							},
							map[string]interface{}{
								"external": "reference2",
							},
						},
					},
				},
			},
			nonReferenceFieldPath: []string{"topField", "nonReferenceField"},
			referenceFieldPath:    []string{"topField", "referenceField"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"field2": "value",
						"referenceField": []interface{}{
							map[string]interface{}{
								"name": "reference1",
							},
							map[string]interface{}{
								"external": "reference2",
							},
						},
					},
				},
			},
		},
		{
			name: "both the non-reference array field and the reference array " +
				"field are set",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"nonReferenceField": []interface{}{
						"testValue1",
						"testValue2",
					},
					"referenceField": []interface{}{
						map[string]interface{}{
							"name": "reference1",
						},
						map[string]interface{}{
							"external": "reference2",
						},
					},
				},
			},
			nonReferenceFieldPath: []string{"nonReferenceField"},
			referenceFieldPath:    []string{"referenceField"},
			hasError:              true,
		},
		{
			name: "both the nested non-reference array field and the nested " +
				"reference array field are set",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"nonReferenceField": []interface{}{
							"testValue1",
							"testValue2",
						},
						"referenceField": []interface{}{
							map[string]interface{}{
								"name": "reference1",
							},
							map[string]interface{}{
								"external": "reference2",
							},
						},
					},
				},
			},
			nonReferenceFieldPath: []string{"topField", "nonReferenceField"},
			referenceFieldPath:    []string{"topField", "referenceField"},
			hasError:              true,
		},
		{
			name: "the non-reference field is not an array",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":            "value",
					"nonReferenceField": "testValue",
				},
			},
			nonReferenceFieldPath: []string{"nonReferenceField"},
			referenceFieldPath:    []string{"referenceField"},
			hasError:              true,
		},
		{
			name: "the reference field is not an array",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1":         "value",
					"referenceField": "testValue",
				},
			},
			nonReferenceFieldPath: []string{"nonReferenceField"},
			referenceFieldPath:    []string{"referenceField"},
			hasError:              true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := FavorReferenceArrayFieldOverNonReferenceArrayField(tc.original, tc.nonReferenceFieldPath, tc.referenceFieldPath)
			if tc.hasError {
				if err == nil {
					t.Fatalf("got nil, but expect to have error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got, want := tc.original, tc.expected; !reflect.DeepEqual(got, want) {
				t.Fatalf("unexpected diff (-want +got): \n%v", cmp.Diff(want, got))
			}
		})
	}
}

func TestPruneNoOpsField(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		original  *k8s.Resource
		fieldPath []string
		expected  *k8s.Resource
	}{
		{
			name: "no-ops fields are not set",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
				},
			},
			fieldPath: []string{"noops"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
				},
			},
		},
		{
			name: "prune no-ops fields",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"noops":  "value",
				},
			},
			fieldPath: []string{"noops"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
				},
			},
		},
		{
			name: "prune nested no-ops fields",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"field2": "value",
						"noops":  "value",
					},
				},
			},
			fieldPath: []string{"topField", "noops"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value",
					"topField": map[string]interface{}{
						"field2": "value",
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if err := PruneNoOpsField(tc.original, tc.fieldPath...); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(tc.original, tc.expected) {
				t.Fatalf("unexpected diff (-want +got): \n%v", cmp.Diff(tc.expected, tc.original))
			}
		})
	}
}

func TestPreserveMutuallyExclusiveNonRefField(t *testing.T) {
	t.Parallel()
	crdWithOptionalReferenceField := &apiextensions.CustomResourceDefinition{
		Spec: apiextensions.CustomResourceDefinitionSpec{
			Versions: []apiextensions.CustomResourceDefinitionVersion{
				{
					Name: "v1beta1",
					Schema: &apiextensions.CustomResourceValidation{
						OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
							Properties: map[string]apiextensions.JSONSchemaProps{
								"spec": {
									Properties: map[string]apiextensions.JSONSchemaProps{
										"testRef": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"name": {
													Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													Type:        "string",
												},
												"namespace": {
													Description: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
													Type:        "string",
												},
												"kind": {
													Description: "Kind of the referent. Allowed values: ReferenceKind",
													Type:        "string",
												},
												"external": {
													Description: "Test description",
													Type:        "string",
												},
											},
											OneOf: []apiextensions.JSONSchemaProps{
												{
													Required: []string{"name", "kind"},
													Not: &apiextensions.JSONSchemaProps{
														Required: []string{"external"},
													},
												},
												{
													Required: []string{"external"},
													Not: &apiextensions.JSONSchemaProps{
														AnyOf: []apiextensions.JSONSchemaProps{
															{Required: []string{"name"}},
															{Required: []string{"namespace"}},
															{Required: []string{"kind"}},
														},
													},
												},
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
							Type: "object",
						},
					},
				},
			},
		},
	}

	tests := []struct {
		name                  string
		originalCRD           *apiextensions.CustomResourceDefinition
		parentPath            []string
		referenceFieldName    string
		nonReferenceFieldName string
		expectedCRD           *apiextensions.CustomResourceDefinition
		hasError              bool
	}{
		{
			name: "required top-level non-reference field is added when there" +
				"is another required field",
			originalCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"topLevelRef": {
													Properties: map[string]apiextensions.JSONSchemaProps{
														"kind": {},
													},
													Description: "fake reference schema for testing",
													Type:        "object",
												},
												"otherRequired": {Type: "string"},
											},
											Required: []string{
												"topLevelRef",
												"otherRequired",
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
			referenceFieldName:    "topLevelRef",
			nonReferenceFieldName: "topLevel",
			expectedCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"topLevelRef": {
													Properties: map[string]apiextensions.JSONSchemaProps{
														"kind": {},
													},
													Description: "fake reference schema for testing",
													Type:        "object",
												},
												"topLevel": {
													Description: "DEPRECATED. Although this field is still available, there is limited support. " +
														"We recommend that you use `spec.topLevelRef` instead.",
													Type: "string",
												},
												"otherRequired": {Type: "string"},
											},
											OneOf: []apiextensions.JSONSchemaProps{
												{Required: []string{"topLevel"}},
												{Required: []string{"topLevelRef"}},
											},
											Required: []string{
												"otherRequired",
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "optional top-level non-reference field is added",
			originalCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"topLevelRef": {
													Properties: map[string]apiextensions.JSONSchemaProps{
														"kind": {},
													},
													Description: "fake reference schema for testing",
													Type:        "object",
												},
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
			referenceFieldName:    "topLevelRef",
			nonReferenceFieldName: "topLevel",
			expectedCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"topLevelRef": {
													Properties: map[string]apiextensions.JSONSchemaProps{
														"kind": {},
													},
													Description: "fake reference schema for testing",
													Type:        "object",
												},
												"topLevel": {
													Description: "DEPRECATED. Although this field is still available, there is limited support. " +
														"We recommend that you use `spec.topLevelRef` instead.",
													Type: "string",
												},
											},
											Not: &apiextensions.JSONSchemaProps{
												Required: []string{"topLevel", "topLevelRef"},
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "required nested non-reference array field is added",
			originalCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"topLevelObject": {
													Properties: map[string]apiextensions.JSONSchemaProps{
														"nestedArrayRefs": {
															Items: &apiextensions.JSONSchemaPropsOrArray{
																Schema: &apiextensions.JSONSchemaProps{
																	Properties: map[string]apiextensions.JSONSchemaProps{
																		"kind": {},
																	},
																	Description: "fake reference schema for testing",
																	Type:        "object",
																},
															},
															Type: "array",
														},
													},
													Required: []string{
														"nestedArrayRefs",
													},
													Type: "object",
												},
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
			parentPath:            []string{"topLevelObject"},
			referenceFieldName:    "nestedArrayRefs",
			nonReferenceFieldName: "nestedArray",
			expectedCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"topLevelObject": {
													Properties: map[string]apiextensions.JSONSchemaProps{
														"nestedArrayRefs": {
															Items: &apiextensions.JSONSchemaPropsOrArray{
																Schema: &apiextensions.JSONSchemaProps{
																	Properties: map[string]apiextensions.JSONSchemaProps{
																		"kind": {},
																	},
																	Description: "fake reference schema for testing",
																	Type:        "object",
																},
															},
															Type: "array",
														},
														"nestedArray": {
															Description: "DEPRECATED. Although this field is still available, there is limited support. " +
																"We recommend that you use `spec.topLevelObject.nestedArrayRefs` instead.",
															Items: &apiextensions.JSONSchemaPropsOrArray{
																Schema: &apiextensions.JSONSchemaProps{
																	Type: "string",
																},
															},
															Type: "array",
														},
													},
													OneOf: []apiextensions.JSONSchemaProps{
														{Required: []string{"nestedArray"}},
														{Required: []string{"nestedArrayRefs"}},
													},
													Type: "object",
												},
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "optional nested non-reference field is added when the " +
				"parent's parent is an array field",
			originalCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"topLevelArray": {
													Items: &apiextensions.JSONSchemaPropsOrArray{
														Schema: &apiextensions.JSONSchemaProps{
															Properties: map[string]apiextensions.JSONSchemaProps{
																"secondLevelObject": {
																	Properties: map[string]apiextensions.JSONSchemaProps{
																		"nestedRef": {
																			Properties: map[string]apiextensions.JSONSchemaProps{
																				"kind": {},
																			},
																			Description: "fake reference schema for testing",
																			Type:        "object",
																		},
																		"otherField": {Type: "string"},
																	},
																	Required: []string{"otherField"},
																	Type:     "object",
																},
																"secondLevelString": {Type: "string"},
															},
															Type: "object",
														},
													},
													Type: "array",
												},
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
			parentPath:            []string{"topLevelArray", "secondLevelObject"},
			referenceFieldName:    "nestedRef",
			nonReferenceFieldName: "nested",
			expectedCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"topLevelArray": {
													Items: &apiextensions.JSONSchemaPropsOrArray{
														Schema: &apiextensions.JSONSchemaProps{
															Properties: map[string]apiextensions.JSONSchemaProps{
																"secondLevelObject": {
																	Properties: map[string]apiextensions.JSONSchemaProps{
																		"nestedRef": {
																			Properties: map[string]apiextensions.JSONSchemaProps{
																				"kind": {},
																			},
																			Description: "fake reference schema for testing",
																			Type:        "object",
																		},
																		"nested": {
																			Description: "DEPRECATED. Although this field is still available, there is limited support. " +
																				"We recommend that you use `spec.topLevelArray.secondLevelObject.nestedRef` instead.",
																			Type: "string",
																		},
																		"otherField": {Type: "string"},
																	},
																	Not: &apiextensions.JSONSchemaProps{
																		Required: []string{"nested", "nestedRef"},
																	},
																	Required: []string{"otherField"},
																	Type:     "object",
																},
																"secondLevelString": {Type: "string"},
															},
															Type: "object",
														},
													},
													Type: "array",
												},
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "optional nested non-reference array field is added when " +
				"the parent is also an array field",
			originalCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"topLevelArray": {
													Items: &apiextensions.JSONSchemaPropsOrArray{
														Schema: &apiextensions.JSONSchemaProps{
															Properties: map[string]apiextensions.JSONSchemaProps{
																"nestedArrayRefs": {
																	Items: &apiextensions.JSONSchemaPropsOrArray{
																		Schema: &apiextensions.JSONSchemaProps{
																			Properties: map[string]apiextensions.JSONSchemaProps{
																				"kind": {},
																			},
																			Description: "fake reference schema for testing",
																			Type:        "object",
																		},
																	},
																	Type: "array",
																},
															},
															Type: "object",
														},
													},
													Type: "array",
												},
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
			parentPath:            []string{"topLevelArray"},
			referenceFieldName:    "nestedArrayRefs",
			nonReferenceFieldName: "nestedArray",
			expectedCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"topLevelArray": {
													Items: &apiextensions.JSONSchemaPropsOrArray{
														Schema: &apiextensions.JSONSchemaProps{
															Properties: map[string]apiextensions.JSONSchemaProps{
																"nestedArrayRefs": {
																	Items: &apiextensions.JSONSchemaPropsOrArray{
																		Schema: &apiextensions.JSONSchemaProps{
																			Properties: map[string]apiextensions.JSONSchemaProps{
																				"kind": {},
																			},
																			Description: "fake reference schema for testing",
																			Type:        "object",
																		},
																	},
																	Type: "array",
																},
																"nestedArray": {
																	Description: "DEPRECATED. Although this field is still available, there is limited support. " +
																		"We recommend that you use `spec.topLevelArray.nestedArrayRefs` instead.",
																	Items: &apiextensions.JSONSchemaPropsOrArray{
																		Schema: &apiextensions.JSONSchemaProps{
																			Type: "string",
																		},
																	},
																	Type: "array",
																},
															},
															Not: &apiextensions.JSONSchemaProps{
																Required: []string{"nestedArray", "nestedArrayRefs"},
															},
															Type: "object",
														},
													},
													Type: "array",
												},
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "reference with no 'kind' field",
			originalCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"testRef": {
													Properties: map[string]apiextensions.JSONSchemaProps{
														"external": {Description: "Test description"},
													},
													Description: "reference schema with no 'kind' field",
													Type:        "object",
												},
											},
											Required: []string{
												"testRef",
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
			referenceFieldName:    "testRef",
			nonReferenceFieldName: "test",
			expectedCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"testRef": {
													Properties: map[string]apiextensions.JSONSchemaProps{
														"external": {Description: "Test description"},
													},
													Description: "reference schema with no 'kind' field",
													Type:        "object",
												},
												"test": {
													Description: "DEPRECATED. Although this field is still available, there is limited support. " +
														"We recommend that you use `spec.testRef` instead.",
													Type: "string",
												},
											},
											OneOf: []apiextensions.JSONSchemaProps{
												{Required: []string{"test"}},
												{Required: []string{"testRef"}},
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "required non-reference field can't be added due to " +
				"existing oneOf rule",
			originalCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"testRef": {
													Description: "fake reference schema for testing",
													Type:        "object",
												},
											},
											OneOf: []apiextensions.JSONSchemaProps{
												{Required: []string{"randomField"}},
												{Required: []string{"randomField2"}},
											},
											Required: []string{
												"testRef",
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
			referenceFieldName:    "testRef",
			nonReferenceFieldName: "test",
			hasError:              true,
		},
		{
			name: "optional non-reference field can't be added due to " +
				"existing not rule",
			originalCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"testRef": {
													Description: "fake reference schema for testing",
													Type:        "object",
												},
											},
											Not: &apiextensions.JSONSchemaProps{
												Required: []string{"randomField"},
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
			parentPath:            []string{},
			referenceFieldName:    "testRef",
			nonReferenceFieldName: "test",
			hasError:              true,
		},
		{
			name: "non-reference field can't be added due to empty reference " +
				"field name",
			originalCRD:           crdWithOptionalReferenceField,
			parentPath:            []string{},
			referenceFieldName:    "",
			nonReferenceFieldName: "test",
			hasError:              true,
		},
		{
			name: "non-reference field can't be added due to empty " +
				"non-reference field name",
			originalCRD:           crdWithOptionalReferenceField,
			parentPath:            []string{},
			referenceFieldName:    "testRef",
			nonReferenceFieldName: "",
			hasError:              true,
		},
		{
			name: "non-reference field can't be added due to incorrect " +
				"reference field name",
			originalCRD:           crdWithOptionalReferenceField,
			parentPath:            []string{},
			referenceFieldName:    "wrongRef",
			nonReferenceFieldName: "wrong",
			hasError:              true,
		},
		{
			name: "non-reference field can't be added due to incorrect " +
				"reference field type",
			originalCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"testRef": {
													Description: "reference schema of wrong type",
													Type:        "string",
												},
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
			parentPath:            []string{},
			referenceFieldName:    "testRef",
			nonReferenceFieldName: "test",
			hasError:              true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := PreserveMutuallyExclusiveNonReferenceField(tc.originalCRD, tc.parentPath, tc.referenceFieldName, tc.nonReferenceFieldName)
			if err != nil {
				if !tc.hasError {
					t.Fatalf("error preserving the mutually exclusive non-reference field: got an error, but want no error: %v", err)
				}
				return
			}
			if !reflect.DeepEqual(tc.expectedCRD, tc.originalCRD) {
				t.Fatalf("unexpected diff in CRD after supporting the mutually exclusive non-reference field (-want +got): \n%v", cmp.Diff(tc.expectedCRD, tc.originalCRD))
			}
		})
	}
}

func TestEnsureReferenceFieldIsMultiKind(t *testing.T) {
	t.Parallel()
	crdWithOptionalReferenceField := &apiextensions.CustomResourceDefinition{
		Spec: apiextensions.CustomResourceDefinitionSpec{
			Versions: []apiextensions.CustomResourceDefinitionVersion{
				{
					Name: "v1beta1",
					Schema: &apiextensions.CustomResourceValidation{
						OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
							Properties: map[string]apiextensions.JSONSchemaProps{
								"spec": {
									Properties: map[string]apiextensions.JSONSchemaProps{
										"testRef": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"name": {
													Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													Type:        "string",
												},
												"namespace": {
													Description: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
													Type:        "string",
												},
												"kind": {
													Description: "Kind of the referent. Allowed values: ReferenceKind",
													Type:        "string",
												},
												"external": {
													Description: "Test description",
													Type:        "string",
												},
											},
											OneOf: []apiextensions.JSONSchemaProps{
												{
													Required: []string{"name", "kind"},
													Not: &apiextensions.JSONSchemaProps{
														Required: []string{"external"},
													},
												},
												{
													Required: []string{"external"},
													Not: &apiextensions.JSONSchemaProps{
														AnyOf: []apiextensions.JSONSchemaProps{
															{Required: []string{"name"}},
															{Required: []string{"namespace"}},
															{Required: []string{"kind"}},
														},
													},
												},
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
							Type: "object",
						},
					},
				},
			},
		},
	}

	tests := []struct {
		name               string
		originalCRD        *apiextensions.CustomResourceDefinition
		parentPath         []string
		referenceFieldName string
		supportedKinds     []MultiKindRef
		expectedCRD        *apiextensions.CustomResourceDefinition
		hasError           bool
	}{
		{
			name: "top-level reference field has 'kind' field",
			originalCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"topLevelRef": {
													Properties: map[string]apiextensions.JSONSchemaProps{
														"kind": {},
													},
													Description: "fake reference schema for testing",
													Type:        "object",
												},
												"topLevel": {
													Description: "DEPRECATED. Although this field is still available, there is limited support. " +
														"We recommend that you use `spec.topLevelRef` instead.",
													Type: "string",
												},
												"otherRequired": {Type: "string"},
											},
											OneOf: []apiextensions.JSONSchemaProps{
												{Required: []string{"topLevel"}},
												{Required: []string{"topLevelRef"}},
											},
											Required: []string{
												"otherRequired",
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
			referenceFieldName: "topLevelRef",
			expectedCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"topLevelRef": {
													Properties: map[string]apiextensions.JSONSchemaProps{
														"kind": {},
													},
													Description: "fake reference schema for testing",
													Type:        "object",
												},
												"topLevel": {
													Description: "DEPRECATED. Although this field is still available, there is limited support. " +
														"We recommend that you use `spec.topLevelRef` instead.",
													Type: "string",
												},
												"otherRequired": {Type: "string"},
											},
											OneOf: []apiextensions.JSONSchemaProps{
												{Required: []string{"topLevel"}},
												{Required: []string{"topLevelRef"}},
											},
											Required: []string{
												"otherRequired",
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "top-level reference field doesn't have 'kind' field",
			originalCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"topLevelRef": {
													Properties: map[string]apiextensions.JSONSchemaProps{
														"external": {Description: "Test description"},
													},
													Description: "reference schema with no 'kind' field",
													Type:        "object",
												},
												"topLevel": {
													Description: "DEPRECATED. Although this field is still available, there is limited support. " +
														"We recommend that you use `spec.topLevelRef` instead.",
													Type: "string",
												},
											},
											Not: &apiextensions.JSONSchemaProps{
												Required: []string{"topLevel", "topLevelRef"},
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
			referenceFieldName: "topLevelRef",
			supportedKinds:     []MultiKindRef{{Kind: "ReferenceKind"}},
			expectedCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"topLevelRef": {
													Properties: map[string]apiextensions.JSONSchemaProps{
														"name": {
															Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
															Type:        "string",
														},
														"namespace": {
															Description: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
															Type:        "string",
														},
														"kind": {
															Description: "Kind of the referent. Allowed values: ReferenceKind",
															Type:        "string",
														},
														"external": {
															Description: "Test description",
															Type:        "string",
														},
													},
													OneOf: []apiextensions.JSONSchemaProps{
														{
															Required: []string{"name", "kind"},
															Not: &apiextensions.JSONSchemaProps{
																Required: []string{"external"},
															},
														},
														{
															Required: []string{"external"},
															Not: &apiextensions.JSONSchemaProps{
																AnyOf: []apiextensions.JSONSchemaProps{
																	{Required: []string{"name"}},
																	{Required: []string{"namespace"}},
																	{Required: []string{"kind"}},
																},
															},
														},
													},
													Type: "object",
												},
												"topLevel": {
													Description: "DEPRECATED. Although this field is still available, there is limited support. " +
														"We recommend that you use `spec.topLevelRef` instead.",
													Type: "string",
												},
											},
											Not: &apiextensions.JSONSchemaProps{
												Required: []string{"topLevel", "topLevelRef"},
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "nested reference array field has 'kind' field",
			originalCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"topLevelObject": {
													Properties: map[string]apiextensions.JSONSchemaProps{
														"nestedArrayRefs": {
															Items: &apiextensions.JSONSchemaPropsOrArray{
																Schema: &apiextensions.JSONSchemaProps{
																	Properties: map[string]apiextensions.JSONSchemaProps{
																		"kind": {},
																	},
																	Description: "fake reference schema for testing",
																	Type:        "object",
																},
															},
															Type: "array",
														},
													},
													Required: []string{
														"nestedArrayRefs",
													},
													Type: "object",
												},
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
			parentPath:         []string{"topLevelObject"},
			referenceFieldName: "nestedArrayRefs",
			expectedCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"topLevelObject": {
													Properties: map[string]apiextensions.JSONSchemaProps{
														"nestedArrayRefs": {
															Items: &apiextensions.JSONSchemaPropsOrArray{
																Schema: &apiextensions.JSONSchemaProps{
																	Properties: map[string]apiextensions.JSONSchemaProps{
																		"kind": {},
																	},
																	Description: "fake reference schema for testing",
																	Type:        "object",
																},
															},
															Type: "array",
														},
													},
													Required: []string{
														"nestedArrayRefs",
													},
													Type: "object",
												},
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "nested reference field doesn't have 'kind' field when the " +
				"parent's parent is an array field",
			originalCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"topLevelArray": {
													Items: &apiextensions.JSONSchemaPropsOrArray{
														Schema: &apiextensions.JSONSchemaProps{
															Properties: map[string]apiextensions.JSONSchemaProps{
																"secondLevelObject": {
																	Properties: map[string]apiextensions.JSONSchemaProps{
																		"nestedRef": {
																			Properties: map[string]apiextensions.JSONSchemaProps{
																				"external": {Description: "Test description"},
																			},
																			Description: "reference schema with no 'kind' field",
																			Type:        "object",
																		},
																		"otherField": {Type: "string"},
																	},
																	Required: []string{"otherField"},
																	Type:     "object",
																},
																"secondLevelString": {Type: "string"},
															},
															Type: "object",
														},
													},
													Type: "array",
												},
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
			parentPath:         []string{"topLevelArray", "secondLevelObject"},
			referenceFieldName: "nestedRef",
			supportedKinds:     []MultiKindRef{{Kind: "ReferenceKind"}},
			expectedCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"topLevelArray": {
													Items: &apiextensions.JSONSchemaPropsOrArray{
														Schema: &apiextensions.JSONSchemaProps{
															Properties: map[string]apiextensions.JSONSchemaProps{
																"secondLevelObject": {
																	Properties: map[string]apiextensions.JSONSchemaProps{
																		"nestedRef": {
																			Properties: map[string]apiextensions.JSONSchemaProps{
																				"name": {
																					Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																					Type:        "string",
																				},
																				"namespace": {
																					Description: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
																					Type:        "string",
																				},
																				"kind": {
																					Description: "Kind of the referent. Allowed values: ReferenceKind",
																					Type:        "string",
																				},
																				"external": {
																					Description: "Test description",
																					Type:        "string",
																				},
																			},
																			OneOf: []apiextensions.JSONSchemaProps{
																				{
																					Required: []string{"name", "kind"},
																					Not: &apiextensions.JSONSchemaProps{
																						Required: []string{"external"},
																					},
																				},
																				{
																					Required: []string{"external"},
																					Not: &apiextensions.JSONSchemaProps{
																						AnyOf: []apiextensions.JSONSchemaProps{
																							{Required: []string{"name"}},
																							{Required: []string{"namespace"}},
																							{Required: []string{"kind"}},
																						},
																					},
																				},
																			},
																			Type: "object",
																		},
																		"otherField": {Type: "string"},
																	},
																	Required: []string{"otherField"},
																	Type:     "object",
																},
																"secondLevelString": {Type: "string"},
															},
															Type: "object",
														},
													},
													Type: "array",
												},
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "nested reference array field doesn't have 'kind' field " +
				"when the parent is also an array field",
			originalCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"topLevelArray": {
													Items: &apiextensions.JSONSchemaPropsOrArray{
														Schema: &apiextensions.JSONSchemaProps{
															Properties: map[string]apiextensions.JSONSchemaProps{
																"nestedArrayRefs": {
																	Items: &apiextensions.JSONSchemaPropsOrArray{
																		Schema: &apiextensions.JSONSchemaProps{
																			Properties: map[string]apiextensions.JSONSchemaProps{
																				"external": {Description: "Test description"},
																			},
																			Description: "reference schema with no 'kind' field",
																			Type:        "object",
																		},
																	},
																	Type: "array",
																},
															},
															Type: "object",
														},
													},
													Type: "array",
												},
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
			parentPath:         []string{"topLevelArray"},
			referenceFieldName: "nestedArrayRefs",
			supportedKinds:     []MultiKindRef{{Kind: "ReferenceKind"}},
			expectedCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"topLevelArray": {
													Items: &apiextensions.JSONSchemaPropsOrArray{
														Schema: &apiextensions.JSONSchemaProps{
															Properties: map[string]apiextensions.JSONSchemaProps{
																"nestedArrayRefs": {
																	Items: &apiextensions.JSONSchemaPropsOrArray{
																		Schema: &apiextensions.JSONSchemaProps{
																			Properties: map[string]apiextensions.JSONSchemaProps{
																				"name": {
																					Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																					Type:        "string",
																				},
																				"namespace": {
																					Description: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
																					Type:        "string",
																				},
																				"kind": {
																					Description: "Kind of the referent. Allowed values: ReferenceKind",
																					Type:        "string",
																				},
																				"external": {
																					Description: "Test description",
																					Type:        "string",
																				},
																			},
																			OneOf: []apiextensions.JSONSchemaProps{
																				{
																					Required: []string{"name", "kind"},
																					Not: &apiextensions.JSONSchemaProps{
																						Required: []string{"external"},
																					},
																				},
																				{
																					Required: []string{"external"},
																					Not: &apiextensions.JSONSchemaProps{
																						AnyOf: []apiextensions.JSONSchemaProps{
																							{Required: []string{"name"}},
																							{Required: []string{"namespace"}},
																							{Required: []string{"kind"}},
																						},
																					},
																				},
																			},
																			Type: "object",
																		},
																	},
																	Type: "array",
																},
															},
															Type: "object",
														},
													},
													Type: "array",
												},
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
		},
		{
			name:               "error to empty reference field name",
			originalCRD:        crdWithOptionalReferenceField,
			parentPath:         []string{},
			referenceFieldName: "",
			hasError:           true,
		},
		{
			name:               "error due to incorrect reference field name",
			originalCRD:        crdWithOptionalReferenceField,
			parentPath:         []string{},
			referenceFieldName: "wrongRef",
			hasError:           true,
		},
		{
			name: "error due to incorrect reference field type",
			originalCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"testRef": {
													Description: "reference schema of wrong type",
													Type:        "string",
												},
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
			parentPath:         []string{},
			referenceFieldName: "testRef",
			hasError:           true,
		},
		{
			name: "non-reference field can't be added because the 'external' " +
				"field is missing when the 'kind' field doesn't exist",
			originalCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"testRef": {
													Description: "reference schema with no 'kind' and 'external' fields",
													Type:        "object",
												},
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
			parentPath:         []string{},
			referenceFieldName: "testRef",
			hasError:           true,
		},
		{
			name: "non-reference field can't be added because supportedKinds " +
				"param is unset when the 'kind' field doesn't exist",
			originalCRD: &apiextensions.CustomResourceDefinition{
				Spec: apiextensions.CustomResourceDefinitionSpec{
					Versions: []apiextensions.CustomResourceDefinitionVersion{
						{
							Name: "v1beta1",
							Schema: &apiextensions.CustomResourceValidation{
								OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
									Properties: map[string]apiextensions.JSONSchemaProps{
										"spec": {
											Properties: map[string]apiextensions.JSONSchemaProps{
												"testRef": {
													Properties: map[string]apiextensions.JSONSchemaProps{
														"external": {Description: "Test description"},
													},
													Description: "reference schema with no 'kind' field",
													Type:        "object",
												},
											},
											Type: "object",
										},
									},
									Type: "object",
								},
							},
						},
					},
				},
			},
			parentPath:         []string{},
			referenceFieldName: "testRef",
			hasError:           true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := EnsureReferenceFieldIsMultiKind(tc.originalCRD, tc.parentPath, tc.referenceFieldName, tc.supportedKinds)
			if err != nil {
				if !tc.hasError {
					t.Fatalf("error ensuring the reference field supports multi-kind: got an error, but want no error: %v", err)
				}
				return
			}
			if !reflect.DeepEqual(tc.expectedCRD, tc.originalCRD) {
				t.Fatalf("unexpected diff in CRD after ensuring the reference field supports multi-kind (-want +got): \n%v", cmp.Diff(tc.expectedCRD, tc.originalCRD))
			}
		})
	}
}

func TestRemovePrefixFromStringFieldInSpec(t *testing.T) {
	t.Parallel()
	prefixToRemove := "prefix/"
	tests := []struct {
		name      string
		original  *k8s.Resource
		fieldPath []string
		expected  *k8s.Resource
		expectErr bool
	}{
		{
			name: "field exists, is string, has matching prefix",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": prefixToRemove + "value1",
					"field2": prefixToRemove + "value2",
				},
			},
			fieldPath: []string{"field1"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value1",
					"field2": prefixToRemove + "value2",
				},
			},
		},
		{
			name: "nested field exists, is string, has matching prefix",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field": prefixToRemove + "value",
					"topField": map[string]interface{}{
						"field": prefixToRemove + "value",
					},
				},
			},
			fieldPath: []string{"topField", "field"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field": prefixToRemove + "value",
					"topField": map[string]interface{}{
						"field": "value",
					},
				},
			},
		},
		{
			name: "field does not exist",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value1",
				},
			},
			fieldPath: []string{"field2"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "value1",
				},
			},
		},
		{
			name: "field exits, is not a string",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": true,
				},
			},
			fieldPath: []string{"field1"},
			expectErr: true,
		},
		{
			name: "field exists, prefix does not match",
			original: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "prefix1/value1",
				},
			},
			fieldPath: []string{"field1"},
			expected: &k8s.Resource{
				TypeMeta: v1.TypeMeta{
					Kind: fooKind,
				},
				Spec: map[string]interface{}{
					"field1": "prefix1/value1",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := removePrefixFromStringFieldInSpec(tc.original, prefixToRemove, tc.fieldPath...)
			if tc.expectErr {
				if err == nil {
					t.Fatalf("got nil, but expect to have error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(tc.original, tc.expected) {
				t.Fatalf("unexpected diff (-want +got): \n%v", cmp.Diff(tc.expected, tc.original))
			}
		})
	}
}
