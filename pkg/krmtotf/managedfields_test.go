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

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestRemoveFieldsFromStateThatConflictWithSpec(t *testing.T) {
	tests := []struct {
		name          string
		rc            corekccv1alpha1.ResourceConfig
		schemaMap     map[string]*tfschema.Schema
		state         map[string]interface{}
		spec          map[string]interface{}
		expectedState map[string]interface{}
	}{
		{
			name: "empty spec",
			rc:   corekccv1alpha1.ResourceConfig{},
			schemaMap: map[string]*tfschema.Schema{
				"foo": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
			},
			state: map[string]interface{}{
				"foo": "fooVal",
			},
			spec: map[string]interface{}{},
			expectedState: map[string]interface{}{
				"foo": "fooVal",
			},
		},
		{
			name: "empty state",
			rc:   corekccv1alpha1.ResourceConfig{},
			schemaMap: map[string]*tfschema.Schema{
				"foo": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
			},
			state: map[string]interface{}{},
			spec: map[string]interface{}{
				"foo": "fooVal",
			},
			expectedState: map[string]interface{}{},
		},
		{
			name: "no conflicting fields",
			rc:   corekccv1alpha1.ResourceConfig{},
			schemaMap: map[string]*tfschema.Schema{
				"foo": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
				"bar": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
			},
			state: map[string]interface{}{
				"foo": "fooVal",
			},
			spec: map[string]interface{}{
				"bar": "barVal",
			},
			expectedState: map[string]interface{}{
				"foo": "fooVal",
			},
		},

		// ConflictsWith tests
		{
			name: "state has top-level field that conflicts with top-level field in spec via ConflictsWith",
			rc:   corekccv1alpha1.ResourceConfig{},
			schemaMap: map[string]*tfschema.Schema{
				"foo": {
					Type:          tfschema.TypeString,
					Optional:      true,
					ConflictsWith: []string{"bar"},
				},
				"bar": {
					Type:          tfschema.TypeString,
					Optional:      true,
					ConflictsWith: []string{"foo"},
				},
				"unrelated": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
			},
			state: map[string]interface{}{
				"foo":       "fooVal",
				"unrelated": "unrelatedVal",
			},
			spec: map[string]interface{}{
				"bar": "barVal",
			},
			expectedState: map[string]interface{}{
				"unrelated": "unrelatedVal",
			},
		},
		{
			name: "state has multiple top-level fields that conflicts with top-level field in spec via ConflictsWith",
			rc:   corekccv1alpha1.ResourceConfig{},
			schemaMap: map[string]*tfschema.Schema{
				"foo": {
					Type:          tfschema.TypeString,
					Optional:      true,
					ConflictsWith: []string{"baz"},
				},
				"bar": {
					Type:          tfschema.TypeString,
					Optional:      true,
					ConflictsWith: []string{"baz"},
				},
				"baz": {
					Type:          tfschema.TypeString,
					Optional:      true,
					ConflictsWith: []string{"foo", "bar"},
				},
				"unrelated": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
			},
			state: map[string]interface{}{
				"foo":       "fooVal",
				"bar":       "barVal",
				"unrelated": "unrelatedVal",
			},
			spec: map[string]interface{}{
				"baz": "bazVal",
			},
			expectedState: map[string]interface{}{
				"unrelated": "unrelatedVal",
			},
		},
		{
			name: "state has multiple top-level fields that conflict with different top-level fields in spec via ConflictsWith",
			rc:   corekccv1alpha1.ResourceConfig{},
			schemaMap: map[string]*tfschema.Schema{
				"foo1": {
					Type:          tfschema.TypeString,
					Optional:      true,
					ConflictsWith: []string{"bar1"},
				},
				"foo2": {
					Type:          tfschema.TypeString,
					Optional:      true,
					ConflictsWith: []string{"bar2"},
				},
				"bar1": {
					Type:          tfschema.TypeString,
					Optional:      true,
					ConflictsWith: []string{"foo1"},
				},
				"bar2": {
					Type:          tfschema.TypeString,
					Optional:      true,
					ConflictsWith: []string{"foo2"},
				},
				"unrelated": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
			},
			state: map[string]interface{}{
				"foo1":      "foo1Val",
				"foo2":      "foo2Val",
				"unrelated": "unrelatedVal",
			},
			spec: map[string]interface{}{
				"bar1": "bar1Val",
				"bar2": "bar2Val",
			},
			expectedState: map[string]interface{}{
				"unrelated": "unrelatedVal",
			},
		},
		{
			name: "state has top-level field that conflicts with nested field in spec via ConflictsWith",
			rc:   corekccv1alpha1.ResourceConfig{},
			schemaMap: map[string]*tfschema.Schema{
				"foo": {
					Type:          tfschema.TypeString,
					Optional:      true,
					ConflictsWith: []string{"bar.0.barbar"},
				},
				"bar": {
					Type:     tfschema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem: &tfschema.Resource{
						Schema: map[string]*tfschema.Schema{
							"barbar": {
								Type:          tfschema.TypeString,
								Optional:      true,
								ConflictsWith: []string{"foo"},
							},
						},
					},
				},
				"unrelated": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
			},
			state: map[string]interface{}{
				"foo":       "fooVal",
				"unrelated": "unrelatedVal",
			},
			spec: map[string]interface{}{
				"bar": map[string]interface{}{
					"barbar": "barbarValInSpec",
				},
			},
			expectedState: map[string]interface{}{
				"unrelated": "unrelatedVal",
			},
		},
		{
			name: "state has nested field that conflicts with top-level field in spec via ConflictsWith",
			rc:   corekccv1alpha1.ResourceConfig{},
			schemaMap: map[string]*tfschema.Schema{
				"foo": {
					Type:          tfschema.TypeString,
					Optional:      true,
					ConflictsWith: []string{"bar.0.barbar"},
				},
				"bar": {
					Type:     tfschema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem: &tfschema.Resource{
						Schema: map[string]*tfschema.Schema{
							"barbar": {
								Type:          tfschema.TypeString,
								Optional:      true,
								ConflictsWith: []string{"foo"},
							},
						},
					},
				},
				"unrelated": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
			},
			state: map[string]interface{}{
				"bar": map[string]interface{}{
					"barbar": "barbarVal",
				},
				"unrelated": "unrelatedVal",
			},
			spec: map[string]interface{}{
				"foo": "fooValInSpec",
			},
			expectedState: map[string]interface{}{
				"unrelated": "unrelatedVal",
			},
		},
		{
			name: "state has nested field that conflicts with top-level field in spec via ConflictsWith (2)",
			rc:   corekccv1alpha1.ResourceConfig{},
			schemaMap: map[string]*tfschema.Schema{
				"foo": {
					Type:          tfschema.TypeString,
					Optional:      true,
					ConflictsWith: []string{"bar.0.barbar"},
				},
				"bar": {
					Type:     tfschema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem: &tfschema.Resource{
						Schema: map[string]*tfschema.Schema{
							"barfoo": {
								Type:     tfschema.TypeString,
								Optional: true,
							},
							"barbar": {
								Type:          tfschema.TypeString,
								Optional:      true,
								ConflictsWith: []string{"foo"},
							},
						},
					},
				},
				"unrelated": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
			},
			state: map[string]interface{}{
				"bar": map[string]interface{}{
					"barfoo": "barfooVal",
					"barbar": "barbarVal",
				},
				"unrelated": "unrelatedVal",
			},
			spec: map[string]interface{}{
				"foo": "fooValInSpec",
			},
			expectedState: map[string]interface{}{
				"bar": map[string]interface{}{
					"barfoo": "barfooVal",
				},
				"unrelated": "unrelatedVal",
			},
		},
		{
			name: "state has nested-nested field that conflicts with top-level field in spec via ConflictsWith",
			rc:   corekccv1alpha1.ResourceConfig{},
			schemaMap: map[string]*tfschema.Schema{
				"foo": {
					Type:          tfschema.TypeString,
					Optional:      true,
					ConflictsWith: []string{"bar.0.barbar.0.barbarbar"},
				},
				"bar": {
					Type:     tfschema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem: &tfschema.Resource{
						Schema: map[string]*tfschema.Schema{
							"barbar": {
								Type:     tfschema.TypeList,
								Optional: true,
								MaxItems: 1,
								Elem: &tfschema.Resource{
									Schema: map[string]*tfschema.Schema{
										"barbarbar": {
											Type:          tfschema.TypeString,
											Optional:      true,
											ConflictsWith: []string{"foo"},
										},
									},
								},
							},
						},
					},
				},
				"unrelated": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
			},
			state: map[string]interface{}{
				"bar": map[string]interface{}{
					"barbar": map[string]interface{}{
						"barbarbar": "barbarbarVal",
					},
				},
				"unrelated": "unrelatedVal",
			},
			spec: map[string]interface{}{
				"foo": "fooValInSpec",
			},
			expectedState: map[string]interface{}{
				"unrelated": "unrelatedVal",
			},
		},
		{
			name: "state has nested field that conflicts with nested field in spec via ConflictsWith",
			rc:   corekccv1alpha1.ResourceConfig{},
			schemaMap: map[string]*tfschema.Schema{
				"foo": {
					Type:     tfschema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem: &tfschema.Resource{
						Schema: map[string]*tfschema.Schema{
							"foofoo": {
								Type:          tfschema.TypeString,
								Optional:      true,
								ConflictsWith: []string{"bar.0.barbar"},
							},
						},
					},
				},
				"bar": {
					Type:     tfschema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem: &tfschema.Resource{
						Schema: map[string]*tfschema.Schema{
							"barbar": {
								Type:          tfschema.TypeString,
								Optional:      true,
								ConflictsWith: []string{"foo.0.foofoo"},
							},
						},
					},
				},
				"unrelated": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
			},
			state: map[string]interface{}{
				"foo": map[string]interface{}{
					"foofoo": "foofooVal",
				},
				"unrelated": "unrelatedVal",
			},
			spec: map[string]interface{}{
				"bar": map[string]interface{}{
					"barbar": "barbarValInSpec",
				},
			},
			expectedState: map[string]interface{}{
				"unrelated": "unrelatedVal",
			},
		},
		{
			name: "state has top-level field that conflicts with top-level reference field in spec via ConflictsWith",
			rc: corekccv1alpha1.ResourceConfig{
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					{
						TFField: "foo",
						TypeConfig: corekccv1alpha1.TypeConfig{
							Key: "fooRef",
						},
					},
				},
			},
			schemaMap: map[string]*tfschema.Schema{
				"foo": {
					Type:          tfschema.TypeString,
					Optional:      true,
					ConflictsWith: []string{"bar"},
				},
				"bar": {
					Type:          tfschema.TypeString,
					Optional:      true,
					ConflictsWith: []string{"foo"},
				},
				"unrelated": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
			},
			state: map[string]interface{}{
				"bar":       "barVal",
				"unrelated": "unrelatedVal",
			},
			spec: map[string]interface{}{
				"fooRef": map[string]interface{}{
					"name": "fooName",
				},
			},
			expectedState: map[string]interface{}{
				"unrelated": "unrelatedVal",
			},
		},
		{
			name: "state has top-level reference field that conflicts with top-level field in spec via ConflictsWith",
			rc: corekccv1alpha1.ResourceConfig{
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					{
						TFField: "foo",
						TypeConfig: corekccv1alpha1.TypeConfig{
							Key: "fooRef",
						},
					},
				},
			},
			schemaMap: map[string]*tfschema.Schema{
				"foo": {
					Type:          tfschema.TypeString,
					Optional:      true,
					ConflictsWith: []string{"bar"},
				},
				"bar": {
					Type:          tfschema.TypeString,
					Optional:      true,
					ConflictsWith: []string{"foo"},
				},
				"unrelated": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
			},
			state: map[string]interface{}{
				"fooRef": map[string]interface{}{
					"name": "fooName",
				},
				"unrelated": "unrelatedVal",
			},
			spec: map[string]interface{}{
				"bar": "barValInSpec",
			},
			expectedState: map[string]interface{}{
				"unrelated": "unrelatedVal",
			},
		},
		{
			name: "state has top-level field that conflicts with nested reference field in spec via ConflictsWith",
			rc: corekccv1alpha1.ResourceConfig{
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					{
						TFField: "foo.foofoo",
						TypeConfig: corekccv1alpha1.TypeConfig{
							Key: "foofooRef",
						},
					},
				},
			},
			schemaMap: map[string]*tfschema.Schema{
				"foo": {
					Type:     tfschema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem: &tfschema.Resource{
						Schema: map[string]*tfschema.Schema{
							"foofoo": {
								Type:          tfschema.TypeString,
								Optional:      true,
								ConflictsWith: []string{"bar"},
							},
						},
					},
				},
				"bar": {
					Type:          tfschema.TypeString,
					Optional:      true,
					ConflictsWith: []string{"foo.0.foofoo"},
				},
				"unrelated": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
			},
			state: map[string]interface{}{
				"bar":       "barVal",
				"unrelated": "unrelatedVal",
			},
			spec: map[string]interface{}{
				"foo": map[string]interface{}{
					"foofooRef": map[string]interface{}{
						"name": "foofooName",
					},
				},
			},
			expectedState: map[string]interface{}{
				"unrelated": "unrelatedVal",
			},
		},
		{
			name: "state has nested field that conflicts with top-level reference field in spec via ConflictsWith",
			rc: corekccv1alpha1.ResourceConfig{
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					{
						TFField: "foo",
						TypeConfig: corekccv1alpha1.TypeConfig{
							Key: "fooRef",
						},
					},
				},
			},
			schemaMap: map[string]*tfschema.Schema{
				"foo": {
					Type:          tfschema.TypeString,
					Optional:      true,
					ConflictsWith: []string{"bar.0.barbar"},
				},
				"bar": {
					Type:     tfschema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem: &tfschema.Resource{
						Schema: map[string]*tfschema.Schema{
							"barbar": {
								Type:          tfschema.TypeString,
								Optional:      true,
								ConflictsWith: []string{"foo"},
							},
						},
					},
				},
				"unrelated": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
			},
			state: map[string]interface{}{
				"bar": map[string]interface{}{
					"barbar": "barbar",
				},
				"unrelated": "unrelatedVal",
			},
			spec: map[string]interface{}{
				"fooRef": map[string]interface{}{
					"name": "fooName",
				},
			},
			expectedState: map[string]interface{}{
				"unrelated": "unrelatedVal",
			},
		},
		{
			name: "state has nested reference field that conflicts with nested reference field in spec via ConflictsWith",
			rc: corekccv1alpha1.ResourceConfig{
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					{
						TFField: "foo.foofoo",
						TypeConfig: corekccv1alpha1.TypeConfig{
							Key: "foofooRef",
						},
					},
					{
						TFField: "bar.barbar",
						TypeConfig: corekccv1alpha1.TypeConfig{
							Key: "barbarRef",
						},
					},
				},
			},
			schemaMap: map[string]*tfschema.Schema{
				"foo": {
					Type:     tfschema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem: &tfschema.Resource{
						Schema: map[string]*tfschema.Schema{
							"foofoo": {
								Type:          tfschema.TypeString,
								Optional:      true,
								ConflictsWith: []string{"bar.0.barbar"},
							},
						},
					},
				},
				"bar": {
					Type:     tfschema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem: &tfschema.Resource{
						Schema: map[string]*tfschema.Schema{
							"barbar": {
								Type:          tfschema.TypeString,
								Optional:      true,
								ConflictsWith: []string{"foo.0.foofoo"},
							},
						},
					},
				},
				"unrelated": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
			},
			state: map[string]interface{}{
				"bar": map[string]interface{}{
					"barbarRef": map[string]interface{}{
						"name": "barbarName",
					},
				},
				"unrelated": "unrelatedVal",
			},
			spec: map[string]interface{}{
				"foo": map[string]interface{}{
					"foofooRef": map[string]interface{}{
						"name": "foofooName",
					},
				},
			},
			expectedState: map[string]interface{}{
				"unrelated": "unrelatedVal",
			},
		},
		{
			name: "state has top-level list field that conflicts with top-level field in spec via ConflictsWith",
			rc:   corekccv1alpha1.ResourceConfig{},
			schemaMap: map[string]*tfschema.Schema{
				"foo": {
					Type:          tfschema.TypeString,
					Optional:      true,
					ConflictsWith: []string{"bar"},
				},
				"bar": {
					Type:     tfschema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					ConflictsWith: []string{"foo"},
				},
				"unrelated": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
			},
			state: map[string]interface{}{
				"bar": []interface{}{
					"barVal1",
					"barVal2",
				},
				"unrelated": "unrelatedVal",
			},
			spec: map[string]interface{}{
				"foo": "fooVal",
			},
			expectedState: map[string]interface{}{
				"unrelated": "unrelatedVal",
			},
		},
		{
			name: "state has top-level list of references field that conflicts with top-level field in spec via ConflictsWith",
			rc: corekccv1alpha1.ResourceConfig{
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					{
						TFField: "bar",
						TypeConfig: corekccv1alpha1.TypeConfig{
							Key: "bars",
						},
					},
				},
			},
			schemaMap: map[string]*tfschema.Schema{
				"foo": {
					Type:          tfschema.TypeString,
					Optional:      true,
					ConflictsWith: []string{"bar"},
				},
				"bar": {
					Type:     tfschema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					ConflictsWith: []string{"foo"},
				},
				"unrelated": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
			},
			state: map[string]interface{}{
				"bars": []interface{}{
					map[string]interface{}{
						"name": "barName1",
					},
					map[string]interface{}{
						"name": "barName2",
					},
				},
				"unrelated": "unrelatedVal",
			},
			spec: map[string]interface{}{
				"foo": "fooVal",
			},
			expectedState: map[string]interface{}{
				"unrelated": "unrelatedVal",
			},
		},

		// ExactlyOneOf tests (select subset of ConflictsWith tests)
		{
			name: "state has top-level field that conflicts with top-level field in spec via ExactlyOneOf",
			rc:   corekccv1alpha1.ResourceConfig{},
			schemaMap: map[string]*tfschema.Schema{
				"foo": {
					Type:         tfschema.TypeString,
					Optional:     true,
					ExactlyOneOf: []string{"foo", "bar"},
				},
				"bar": {
					Type:         tfschema.TypeString,
					Optional:     true,
					ExactlyOneOf: []string{"foo", "bar"},
				},
				"unrelated": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
			},
			state: map[string]interface{}{
				"foo":       "fooVal",
				"unrelated": "unrelatedVal",
			},
			spec: map[string]interface{}{
				"bar": "barVal",
			},
			expectedState: map[string]interface{}{
				"unrelated": "unrelatedVal",
			},
		},
		{
			name: "state has multiple top-level fields that conflicts with top-level field in spec via ExactlyOneOf",
			rc:   corekccv1alpha1.ResourceConfig{},
			schemaMap: map[string]*tfschema.Schema{
				"foo": {
					Type:         tfschema.TypeString,
					Optional:     true,
					ExactlyOneOf: []string{"foo", "baz"},
				},
				"bar": {
					Type:         tfschema.TypeString,
					Optional:     true,
					ExactlyOneOf: []string{"bar", "baz"},
				},
				"baz": {
					Type:         tfschema.TypeString,
					Optional:     true,
					ExactlyOneOf: []string{"foo", "bar", "baz"},
				},
				"unrelated": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
			},
			state: map[string]interface{}{
				"foo":       "fooVal",
				"bar":       "barVal",
				"unrelated": "unrelatedVal",
			},
			spec: map[string]interface{}{
				"baz": "bazVal",
			},
			expectedState: map[string]interface{}{
				"unrelated": "unrelatedVal",
			},
		},
		{
			name: "state has nested field that conflicts with top-level field in spec via ExactlyOneOf",
			rc:   corekccv1alpha1.ResourceConfig{},
			schemaMap: map[string]*tfschema.Schema{
				"foo": {
					Type:         tfschema.TypeString,
					Optional:     true,
					ExactlyOneOf: []string{"foo", "bar.0.barbar"},
				},
				"bar": {
					Type:     tfschema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem: &tfschema.Resource{
						Schema: map[string]*tfschema.Schema{
							"barbar": {
								Type:         tfschema.TypeString,
								Optional:     true,
								ExactlyOneOf: []string{"foo", "bar.0.barbar"},
							},
						},
					},
				},
				"unrelated": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
			},
			state: map[string]interface{}{
				"bar": map[string]interface{}{
					"barbar": "barbarVal",
				},
				"unrelated": "unrelatedVal",
			},
			spec: map[string]interface{}{
				"foo": "fooVal",
			},
			expectedState: map[string]interface{}{
				"unrelated": "unrelatedVal",
			},
		},
		{
			name: "state has top-level reference field that conflicts with top-level field in spec via ExactlyOneOf",
			rc: corekccv1alpha1.ResourceConfig{
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					{
						TFField: "foo",
						TypeConfig: corekccv1alpha1.TypeConfig{
							Key: "fooRef",
						},
					},
				},
			},
			schemaMap: map[string]*tfschema.Schema{
				"foo": {
					Type:         tfschema.TypeString,
					Optional:     true,
					ExactlyOneOf: []string{"foo", "bar"},
				},
				"bar": {
					Type:         tfschema.TypeString,
					Optional:     true,
					ExactlyOneOf: []string{"foo", "bar"},
				},
				"unrelated": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
			},
			state: map[string]interface{}{
				"fooRef": map[string]interface{}{
					"name": "fooName",
				},
				"unrelated": "unrelatedVal",
			},
			spec: map[string]interface{}{
				"bar": "barVal",
			},
			expectedState: map[string]interface{}{
				"unrelated": "unrelatedVal",
			},
		},
		{
			name: "state has nested reference field that conflicts with nested reference field in spec via ExactlyOneOf",
			rc: corekccv1alpha1.ResourceConfig{
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					{
						TFField: "foo.foofoo",
						TypeConfig: corekccv1alpha1.TypeConfig{
							Key: "foofooRef",
						},
					},
					{
						TFField: "bar.barbar",
						TypeConfig: corekccv1alpha1.TypeConfig{
							Key: "barbarRef",
						},
					},
				},
			},
			schemaMap: map[string]*tfschema.Schema{
				"foo": {
					Type:     tfschema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem: &tfschema.Resource{
						Schema: map[string]*tfschema.Schema{
							"foofoo": {
								Type:         tfschema.TypeString,
								Optional:     true,
								ExactlyOneOf: []string{"foo.0.foofoo", "bar.0.barbar"},
							},
						},
					},
				},
				"bar": {
					Type:     tfschema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem: &tfschema.Resource{
						Schema: map[string]*tfschema.Schema{
							"barbar": {
								Type:         tfschema.TypeString,
								Optional:     true,
								ExactlyOneOf: []string{"foo.0.foofoo", "bar.0.barbar"},
							},
						},
					},
				},
				"unrelated": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
			},
			state: map[string]interface{}{
				"bar": map[string]interface{}{
					"barbarRef": map[string]interface{}{
						"name": "barbarName",
					},
				},
				"unrelated": "unrelatedVal",
			},
			spec: map[string]interface{}{
				"foo": map[string]interface{}{
					"foofooRef": map[string]interface{}{
						"name": "foofooName",
					},
				},
			},
			expectedState: map[string]interface{}{
				"unrelated": "unrelatedVal",
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if err := krmtotf.RemoveFieldsFromStateThatConflictWithSpec(tc.state, tc.spec, tc.rc, []string{}, tc.schemaMap); err != nil {
				t.Fatalf("got error, wanted none: %v", err)
			}
			if got, want := tc.state, tc.expectedState; !test.Equals(t, got, want) {
				t.Fatalf("unexpected diff in resulting state (-want +got): \n%v", cmp.Diff(want, got))
			}
		})
	}
}
