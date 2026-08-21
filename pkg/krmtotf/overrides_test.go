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
	"reflect"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	. "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/provider"
)

func Test_FlattenComputeInstanceMetadata(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name: "flattens structured list to simple map",
			input: map[string]interface{}{
				"key": "value",
				"metadata": []interface{}{
					map[string]interface{}{
						"key":   "foo",
						"value": "bar",
					},
					map[string]interface{}{
						"key":   "bar",
						"value": "baz",
					},
				},
			},
			expected: map[string]interface{}{
				"key": "value",
				"metadata": map[string]interface{}{
					"foo": "bar",
					"bar": "baz",
				},
			},
		},
		{
			name: "no-op if no metadata given",
			input: map[string]interface{}{
				"key": "value",
			},
			expected: map[string]interface{}{
				"key": "value",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			output, err := FlattenComputeInstanceMetadata(tc.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !test.Equals(t, tc.expected, output) {
				t.Fatalf("expected: %+v, actual: %+v", tc.expected, output)
			}
		})
	}
}

func Test_ExpandComputeInstanceMetadata(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		prev     *Resource
		expected map[string]interface{}
	}{
		{
			name: "expands simple map to structured list",
			input: map[string]interface{}{
				"key": "value",
				"metadata": map[string]interface{}{
					"foo": "bar",
					"bar": "baz",
				},
			},
			prev: &Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{},
				},
			},
			expected: map[string]interface{}{
				"key": "value",
				"metadata": []interface{}{
					map[string]interface{}{
						"key":   "bar",
						"value": "baz",
					},
					map[string]interface{}{
						"key":   "foo",
						"value": "bar",
					},
				},
			},
		},
		{
			name: "previous value is source of truth",
			input: map[string]interface{}{
				"key": "value",
				"metadata": map[string]interface{}{
					"foo": "bar",
					"bar": "baz",
				},
			},
			prev: &Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"metadata": []interface{}{
							map[string]interface{}{
								"key":   "foo",
								"value": "bar",
							},
							map[string]interface{}{
								"key":   "bar",
								"value": "baz",
							},
						},
					},
				},
			},
			expected: map[string]interface{}{
				"key": "value",
				"metadata": []interface{}{
					map[string]interface{}{
						"key":   "foo",
						"value": "bar",
					},
					map[string]interface{}{
						"key":   "bar",
						"value": "baz",
					},
				},
			},
		},
		{
			name: "previous value is source of truth 2",
			input: map[string]interface{}{
				"key": "value",
				"metadata": map[string]interface{}{
					"foo": "bar",
					"bar": "baz",
					"baz": "abc",
				},
			},
			prev: &Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"metadata": []interface{}{
							map[string]interface{}{
								"key":   "foo",
								"value": "bar",
							},
							map[string]interface{}{
								"key":   "bar",
								"value": "baz",
							},
						},
					},
				},
			},
			expected: map[string]interface{}{
				"key": "value",
				"metadata": []interface{}{
					map[string]interface{}{
						"key":   "foo",
						"value": "bar",
					},
					map[string]interface{}{
						"key":   "bar",
						"value": "baz",
					},
				},
			},
		},
		{
			name: "no-op if no metadata given",
			input: map[string]interface{}{
				"key": "value",
			},
			prev: &Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{},
				},
			},
			expected: map[string]interface{}{
				"key": "value",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			output := ExpandComputeInstanceMetadata(tc.input, tc.prev)
			if !test.Equals(t, tc.expected, output) {
				t.Fatalf("unexpected diff in output (-want +got): \n%v", cmp.Diff(tc.expected, output))
			}
		})
	}
}

func Test_MergeClusterConfigsFromLiveState(t *testing.T) {
	t.Parallel()
	resourceMap := provider.ResourceMap()
	tests := []struct {
		name      string
		config    map[string]interface{}
		liveState map[string]interface{}
		expected  map[string]interface{}
	}{
		{
			name: "no changes on the existing clusters",
			config: map[string]interface{}{
				"cluster": []interface{}{
					map[string]interface{}{
						"cluster_id":   "test-c1",
						"num_nodes":    float64(2),
						"storage_type": "SSD",
						"zone":         "us-west1-b",
					},
					map[string]interface{}{
						"cluster_id": "test-c2",
						"num_nodes":  float64(3),
						"zone":       "us-east1-c",
					},
					map[string]interface{}{
						"cluster_id":   "test-c3",
						"num_nodes":    float64(1),
						"storage_type": "SSD",
						"zone":         "us-east1-b",
					},
				},
			},
			liveState: map[string]interface{}{
				"cluster": []interface{}{
					map[string]interface{}{
						"cluster_id":   "test-c3",
						"num_nodes":    float64(1),
						"storage_type": "SSD",
						"zone":         "us-east1-b",
					},
					map[string]interface{}{
						"cluster_id":   "test-c1",
						"num_nodes":    float64(2),
						"storage_type": "SSD",
						"zone":         "us-west1-b",
					},
					map[string]interface{}{
						"cluster_id":   "test-c2",
						"num_nodes":    float64(3),
						"zone":         "us-east1-c",
						"storage_type": "SSD",
					},
				},
			},
			expected: map[string]interface{}{
				"cluster": []interface{}{
					map[string]interface{}{
						"cluster_id":   "test-c1",
						"num_nodes":    float64(2),
						"storage_type": "SSD",
						"zone":         "us-west1-b",
					},
					map[string]interface{}{
						"cluster_id":   "test-c2",
						"num_nodes":    float64(3),
						"storage_type": "SSD",
						"zone":         "us-east1-c",
					},
					map[string]interface{}{
						"cluster_id":   "test-c3",
						"num_nodes":    float64(1),
						"storage_type": "SSD",
						"zone":         "us-east1-b",
					},
				},
			},
		},
		{
			name: "changes on the existing clusters are preserved",
			config: map[string]interface{}{
				"cluster": []interface{}{
					map[string]interface{}{
						"cluster_id":   "test-c1",
						"num_nodes":    float64(3),
						"storage_type": "SSD",
						"zone":         "us-west1-b",
					},
					map[string]interface{}{
						"cluster_id": "test-c2",
						"num_nodes":  float64(3),
						"zone":       "us-east1-c",
					},
					map[string]interface{}{
						"cluster_id":   "test-c3",
						"num_nodes":    float64(2),
						"storage_type": "SSD",
						"zone":         "us-east1-b",
					},
				},
			},
			liveState: map[string]interface{}{
				"cluster": []interface{}{
					map[string]interface{}{
						"cluster_id":   "test-c3",
						"num_nodes":    float64(1),
						"storage_type": "SSD",
						"zone":         "us-east1-b",
					},
					map[string]interface{}{
						"cluster_id":   "test-c1",
						"num_nodes":    float64(2),
						"storage_type": "SSD",
						"zone":         "us-west1-b",
					},
					map[string]interface{}{
						"cluster_id":   "test-c2",
						"num_nodes":    float64(3),
						"zone":         "us-east1-c",
						"storage_type": "SSD",
					},
				},
			},
			expected: map[string]interface{}{
				"cluster": []interface{}{
					map[string]interface{}{
						"cluster_id":   "test-c1",
						"num_nodes":    float64(3),
						"storage_type": "SSD",
						"zone":         "us-west1-b",
					},
					map[string]interface{}{
						"cluster_id":   "test-c2",
						"num_nodes":    float64(3),
						"storage_type": "SSD",
						"zone":         "us-east1-c",
					},
					map[string]interface{}{
						"cluster_id":   "test-c3",
						"num_nodes":    float64(2),
						"storage_type": "SSD",
						"zone":         "us-east1-b",
					},
				},
			},
		},
		{
			name: "omitted optional fields for the existing clusters are populated from live state",
			config: map[string]interface{}{
				"cluster": []interface{}{
					map[string]interface{}{
						"cluster_id": "test-c1",
						"zone":       "us-west1-b",
					},
					map[string]interface{}{
						"cluster_id": "test-c2",
						"zone":       "us-east1-c",
					},
					map[string]interface{}{
						"cluster_id": "test-c3",
						"zone":       "us-east1-b",
					},
				},
			},
			liveState: map[string]interface{}{
				"cluster": []interface{}{
					map[string]interface{}{
						"cluster_id":   "test-c3",
						"num_nodes":    float64(1),
						"storage_type": "SSD",
						"zone":         "us-east1-b",
					},
					map[string]interface{}{
						"cluster_id":   "test-c1",
						"num_nodes":    float64(2),
						"storage_type": "SSD",
						"zone":         "us-west1-b",
					},
					map[string]interface{}{
						"cluster_id":   "test-c2",
						"num_nodes":    float64(3),
						"zone":         "us-east1-c",
						"storage_type": "SSD",
					},
				},
			},
			expected: map[string]interface{}{
				"cluster": []interface{}{
					map[string]interface{}{
						"cluster_id":   "test-c1",
						"num_nodes":    float64(2),
						"storage_type": "SSD",
						"zone":         "us-west1-b",
					},
					map[string]interface{}{
						"cluster_id":   "test-c2",
						"num_nodes":    float64(3),
						"storage_type": "SSD",
						"zone":         "us-east1-c",
					},
					map[string]interface{}{
						"cluster_id":   "test-c3",
						"num_nodes":    float64(1),
						"storage_type": "SSD",
						"zone":         "us-east1-b",
					},
				},
			},
		},
		{
			name: "remove an old cluster, add a new cluster and change an existing cluster",
			config: map[string]interface{}{
				"cluster": []interface{}{
					map[string]interface{}{
						"cluster_id": "test-c2",
						"zone":       "us-east1-c",
					},
					map[string]interface{}{
						"cluster_id": "test-c4",
						"zone":       "us-east1-b",
						"num_nodes":  float64(4),
					},
					map[string]interface{}{
						"cluster_id": "test-c3",
						"zone":       "us-east1-b",
						"num_nodes":  float64(2),
					},
				},
			},
			liveState: map[string]interface{}{
				"cluster": []interface{}{
					map[string]interface{}{
						"cluster_id":   "test-c3",
						"num_nodes":    float64(1),
						"storage_type": "SSD",
						"zone":         "us-east1-b",
					},
					map[string]interface{}{
						"cluster_id":   "test-c1",
						"num_nodes":    float64(2),
						"storage_type": "SSD",
						"zone":         "us-west1-b",
					},
					map[string]interface{}{
						"cluster_id":   "test-c2",
						"num_nodes":    float64(3),
						"zone":         "us-east1-c",
						"storage_type": "SSD",
					},
				},
			},
			expected: map[string]interface{}{
				"cluster": []interface{}{
					map[string]interface{}{
						"cluster_id":   "test-c2",
						"num_nodes":    float64(3),
						"storage_type": "SSD",
						"zone":         "us-east1-c",
					},
					map[string]interface{}{
						"cluster_id": "test-c4",
						"zone":       "us-east1-b",
						"num_nodes":  float64(4),
					},
					map[string]interface{}{
						"cluster_id":   "test-c3",
						"num_nodes":    float64(2),
						"storage_type": "SSD",
						"zone":         "us-east1-b",
					},
				},
			},
		},
		{
			name: "remove all old clusters",
			config: map[string]interface{}{
				"cluster": []interface{}{},
			},
			liveState: map[string]interface{}{
				"cluster": []interface{}{
					map[string]interface{}{
						"cluster_id":   "test-c3",
						"num_nodes":    float64(1),
						"storage_type": "SSD",
						"zone":         "us-east1-b",
					},
					map[string]interface{}{
						"cluster_id":   "test-c1",
						"num_nodes":    float64(2),
						"storage_type": "SSD",
						"zone":         "us-west1-b",
					},
					map[string]interface{}{
						"cluster_id":   "test-c2",
						"num_nodes":    float64(3),
						"zone":         "us-east1-c",
						"storage_type": "SSD",
					},
				},
			},
			expected: map[string]interface{}{
				"cluster": []interface{}{},
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			res, err := MergeClusterConfigsFromLiveStateForBigtableInstance(tc.config, tc.liveState, resourceMap["google_bigtable_instance"])
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(res, tc.expected) {
				diff := cmp.Diff(tc.expected, res)
				t.Fatalf("The merged config has diff (-want, +got):\n%v", diff)
			}
		})
	}
}
