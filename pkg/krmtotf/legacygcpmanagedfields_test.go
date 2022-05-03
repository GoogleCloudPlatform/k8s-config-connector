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
	"encoding/json"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	. "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var emptyObject = make(map[string]interface{})

func TestResolveGCPManagedFields(t *testing.T) {
	tests := []struct {
		name              string
		kind              string
		lastAppliedConfig map[string]interface{}
		resourceExists    bool
		inputConfig       map[string]interface{}
		expectedConfig    map[string]interface{}
	}{
		{
			name: "ContainerCluster treats node version as GCP-managed when release channel set",
			kind: "ContainerCluster",
			lastAppliedConfig: map[string]interface{}{
				"spec": map[string]interface{}{
					"otherField": "otherValue",
					"releaseChannel": map[string]interface{}{
						"channel": "REGULAR",
					},
				},
			},
			resourceExists: true,
			inputConfig: map[string]interface{}{
				"otherField": "otherValue",
				"releaseChannel": map[string]interface{}{
					"channel": "REGULAR",
				},
				"nodeVersion": "1.14",
			},
			expectedConfig: map[string]interface{}{
				"otherField": "otherValue",
				"releaseChannel": map[string]interface{}{
					"channel": "REGULAR",
				},
			},
		},
		{
			name: "SQLInstance treats disk size as GCP-managed when disk autoresize set",
			kind: "SQLInstance",
			lastAppliedConfig: map[string]interface{}{
				"spec": map[string]interface{}{
					"otherField": "otherValue",
					"settings": map[string]interface{}{
						"diskAutoresize": true,
					},
				},
			},
			resourceExists: true,
			inputConfig: map[string]interface{}{
				"otherField": "otherValue",
				"settings": map[string]interface{}{
					"diskAutoresize": true,
					"diskSize":       50,
				},
			},
			expectedConfig: map[string]interface{}{
				"otherField": "otherValue",
				"settings": map[string]interface{}{
					"diskAutoresize": true,
				},
			},
		},
		{
			name: "ComputeBackendService treats backends as unmanaged when backends not set",
			kind: "ComputeBackendService",
			lastAppliedConfig: map[string]interface{}{
				"spec": map[string]interface{}{
					"otherField": "otherValue",
				},
			},
			resourceExists: true,
			inputConfig: map[string]interface{}{
				"otherField": "otherValue",
				"backend":    []interface{}{},
			},
			expectedConfig: map[string]interface{}{
				"otherField": "otherValue",
			},
		},
		{
			name: "ComputeBackendService manages user-supplied backends",
			kind: "ComputeBackendService",
			lastAppliedConfig: map[string]interface{}{
				"spec": map[string]interface{}{
					"otherField": "otherValue",
					"backend": []interface{}{
						"backend1",
					},
				},
			},
			resourceExists: true,
			inputConfig: map[string]interface{}{
				"otherField": "otherValue",
				"backend": []interface{}{
					"backend1",
				},
			},
			expectedConfig: map[string]interface{}{
				"otherField": "otherValue",
				"backend": []interface{}{
					"backend1",
				},
			},
		},
		{
			name: "user-applied value is explicit override of GCP",
			kind: "ContainerCluster",
			lastAppliedConfig: map[string]interface{}{
				"spec": map[string]interface{}{
					"otherField": "otherValue",
					"releaseChannel": map[string]interface{}{
						"channel": "REGULAR",
					},
					"nodeVersion": "1.14",
				},
			},
			resourceExists: true,
			inputConfig: map[string]interface{}{
				"otherField": "otherValue",
				"releaseChannel": map[string]interface{}{
					"channel": "REGULAR",
				},
				"nodeVersion": "1.14",
			},
			expectedConfig: map[string]interface{}{
				"otherField": "otherValue",
				"releaseChannel": map[string]interface{}{
					"channel": "REGULAR",
				},
				"nodeVersion": "1.14",
			},
		},
		{
			name:           "use explicit config when creating resource",
			kind:           "ContainerCluster",
			resourceExists: false,
			inputConfig: map[string]interface{}{
				"otherField": "otherValue",
				"releaseChannel": map[string]interface{}{
					"channel": "REGULAR",
				},
				"nodeVersion": "1.14",
			},
			expectedConfig: map[string]interface{}{
				"otherField": "otherValue",
				"releaseChannel": map[string]interface{}{
					"channel": "REGULAR",
				},
				"nodeVersion": "1.14",
			},
		},
		{
			name: "ContainerNodePool treats version field as GCP-managed when autoUpgrade set",
			kind: "ContainerNodePool",
			lastAppliedConfig: map[string]interface{}{
				"spec": map[string]interface{}{
					"otherField": "otherValue",
					"management": map[string]interface{}{
						"autoUpgrade": true,
					},
				},
			},
			resourceExists: true,
			inputConfig: map[string]interface{}{
				"otherField": "otherValue",
				"management": map[string]interface{}{
					"autoUpgrade": true,
				},
				"version": "1.18.0-gke.0",
			},
			expectedConfig: map[string]interface{}{
				"otherField": "otherValue",
				"management": map[string]interface{}{
					"autoUpgrade": true,
				},
			},
		},
		{
			name: "ContainerNodePool uses spec value when autoUpgrade turned off",
			kind: "ContainerNodePool",
			lastAppliedConfig: map[string]interface{}{
				"spec": map[string]interface{}{
					"otherField": "otherValue",
					"management": map[string]interface{}{
						"autoUpgrade": false,
					},
				},
			},
			resourceExists: true,
			inputConfig: map[string]interface{}{
				"otherField": "otherValue",
				"management": map[string]interface{}{
					"autoUpgrade": false,
				},
				"version": "1.18.0-gke.0",
			},
			expectedConfig: map[string]interface{}{
				"otherField": "otherValue",
				"management": map[string]interface{}{
					"autoUpgrade": false,
				},
				"version": "1.18.0-gke.0",
			},
		},

		{
			name: "ContainerNodePool uses user-supplied version when explicitly applied",
			kind: "ContainerNodePool",
			lastAppliedConfig: map[string]interface{}{
				"spec": map[string]interface{}{
					"otherField": "otherValue",
					"management": map[string]interface{}{
						"autoUpgrade": true,
					},
					"version": "1.18.0-gke.0",
				},
			},
			resourceExists: true,
			inputConfig: map[string]interface{}{
				"otherField": "otherValue",
				"management": map[string]interface{}{
					"autoUpgrade": true,
				},
				"version": "1.18.0-gke.0",
			},
			expectedConfig: map[string]interface{}{
				"otherField": "otherValue",
				"management": map[string]interface{}{
					"autoUpgrade": true,
				},
				"version": "1.18.0-gke.0",
			},
		},
		{
			name: "ContainerNodePool treats initialNodeCount as GCP-managed",
			kind: "ContainerNodePool",
			lastAppliedConfig: map[string]interface{}{
				"spec": map[string]interface{}{
					"otherField": "otherValue",
				},
			},
			resourceExists: true,
			inputConfig: map[string]interface{}{
				"otherField":       "otherValue",
				"initialNodeCount": 1,
			},
			expectedConfig: map[string]interface{}{
				"otherField": "otherValue",
			},
		},
		{
			name: "ContainerNodePool treats nodeCount as GCP-managed if autoscaling set",
			kind: "ContainerNodePool",
			lastAppliedConfig: map[string]interface{}{
				"spec": map[string]interface{}{
					"otherField": "otherValue",
					"autoscaling": map[string]interface{}{
						"minNodeCount": 1,
						"maxNodeCount": 3,
					},
				},
			},
			resourceExists: true,
			inputConfig: map[string]interface{}{
				"otherField": "otherValue",
				"autoscaling": map[string]interface{}{
					"minNodeCount": 1,
					"maxNodeCount": 3,
				},
				"nodeCount": 1,
			},
			expectedConfig: map[string]interface{}{
				"otherField": "otherValue",
				"autoscaling": map[string]interface{}{
					"minNodeCount": 1,
					"maxNodeCount": 3,
				},
			},
		},
		{
			name: "ContainerNodePool uses explicit nodeCount when autoscaling disabled",
			kind: "ContainerNodePool",
			lastAppliedConfig: map[string]interface{}{
				"spec": map[string]interface{}{
					"otherField": "otherValue",
				},
			},
			resourceExists: true,
			inputConfig: map[string]interface{}{
				"otherField": "otherValue",
				"nodeCount":  1,
			},
			expectedConfig: map[string]interface{}{
				"otherField": "otherValue",
				"nodeCount":  1,
			},
		},
		{
			name: "ContainerNodePool uses explicit nodeCount when autoscaling set explicitly to nil",
			kind: "ContainerNodePool",
			lastAppliedConfig: map[string]interface{}{
				"spec": map[string]interface{}{
					"otherField":  "otherValue",
					"autoscaling": nil,
				},
			},
			resourceExists: true,
			inputConfig: map[string]interface{}{
				"otherField":  "otherValue",
				"autoscaling": nil,
				"nodeCount":   1,
			},
			expectedConfig: map[string]interface{}{
				"otherField":  "otherValue",
				"autoscaling": nil,
				"nodeCount":   1,
			},
		},
		{
			name: "BigtableInstance removes numNodes when not set",
			kind: "BigtableInstance",
			lastAppliedConfig: map[string]interface{}{
				"spec": map[string]interface{}{
					"otherField": "otherValue",
					"cluster": []interface{}{
						map[string]interface{}{
							"clusterId": "test1",
							"zone":      "us-central1-a",
						},
						map[string]interface{}{
							"clusterId": "test2",
							"zone":      "us-west1-a",
						},
					},
				},
			},
			resourceExists: true,
			inputConfig: map[string]interface{}{
				"otherField": "otherValue",
				"cluster": []interface{}{
					map[string]interface{}{
						"clusterId": "test1",
						"zone":      "us-central1-a",
						"numNodes":  float64(1),
					},
					map[string]interface{}{
						"clusterId": "test2",
						"zone":      "us-west1-a",
						"numNodes":  float64(1),
					},
				},
			},
			expectedConfig: map[string]interface{}{
				"otherField": "otherValue",
				"cluster": []interface{}{
					map[string]interface{}{
						"clusterId": "test1",
						"zone":      "us-central1-a",
					},
					map[string]interface{}{
						"clusterId": "test2",
						"zone":      "us-west1-a",
					},
				},
			},
		},
		{
			name: "BigtableInstance respects numNodes when specified",
			kind: "BigtableInstance",
			lastAppliedConfig: map[string]interface{}{
				"spec": map[string]interface{}{
					"otherField": "otherValue",
					"cluster": []interface{}{
						map[string]interface{}{
							"clusterId": "test1",
							"zone":      "us-central1-a",
							"numNodes":  float64(2),
						},
						map[string]interface{}{
							"clusterId": "test2",
							"zone":      "us-west1-a",
							"numNodes":  float64(3),
						},
					},
				},
			},
			resourceExists: true,
			inputConfig: map[string]interface{}{
				"otherField": "otherValue",
				"cluster": []interface{}{
					map[string]interface{}{
						"clusterId": "test1",
						"zone":      "us-central1-a",
						"numNodes":  float64(2),
					},
					map[string]interface{}{
						"clusterId": "test2",
						"zone":      "us-west1-a",
						"numNodes":  float64(3),
					},
				},
			},
			expectedConfig: map[string]interface{}{
				"otherField": "otherValue",
				"cluster": []interface{}{
					map[string]interface{}{
						"clusterId": "test1",
						"zone":      "us-central1-a",
						"numNodes":  float64(2),
					},
					map[string]interface{}{
						"clusterId": "test2",
						"zone":      "us-west1-a",
						"numNodes":  float64(3),
					},
				},
			},
		},
		{
			name: "BigtableInstance handles mixed set and unset numNodes",
			kind: "BigtableInstance",
			lastAppliedConfig: map[string]interface{}{
				"spec": map[string]interface{}{
					"otherField": "otherValue",
					"cluster": []interface{}{
						map[string]interface{}{
							"clusterId": "test1",
							"zone":      "us-central1-a",
							"numNodes":  float64(2),
						},
						map[string]interface{}{
							"clusterId": "test2",
							"zone":      "us-west1-a",
						},
					},
				},
			},
			resourceExists: true,
			inputConfig: map[string]interface{}{
				"otherField": "otherValue",
				"cluster": []interface{}{
					map[string]interface{}{
						"clusterId": "test1",
						"zone":      "us-central1-a",
						"numNodes":  float64(2),
					},
					map[string]interface{}{
						"clusterId": "test2",
						"zone":      "us-west1-a",
						"numNodes":  float64(3),
					},
				},
			},
			expectedConfig: map[string]interface{}{
				"otherField": "otherValue",
				"cluster": []interface{}{
					map[string]interface{}{
						"clusterId": "test1",
						"zone":      "us-central1-a",
						"numNodes":  float64(2),
					},
					map[string]interface{}{
						"clusterId": "test2",
						"zone":      "us-west1-a",
					},
				},
			},
		},
		{
			name: "BigtableInstance removes an existing numNodes, if the value is removed",
			kind: "BigtableInstance",
			lastAppliedConfig: map[string]interface{}{
				"spec": map[string]interface{}{
					"otherField": "otherValue",
					"cluster": []interface{}{
						map[string]interface{}{
							"clusterId": "test1",
							"zone":      "us-central1-a",
							"numNodes":  float64(2),
						},
					},
				},
			},
			resourceExists: true,
			inputConfig: map[string]interface{}{
				"otherField": "otherValue",
				"cluster": []interface{}{
					map[string]interface{}{
						"clusterId": "test1",
						"zone":      "us-central1-a",
					},
				},
			},
			expectedConfig: map[string]interface{}{
				"otherField": "otherValue",
				"cluster": []interface{}{
					map[string]interface{}{
						"clusterId": "test1",
						"zone":      "us-central1-a",
					},
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := resourceSkeleton()
			r.SetGroupVersionKind(schema.GroupVersionKind{Kind: tc.kind})
			lastAppliedConfigJSON, err := json.Marshal(tc.lastAppliedConfig)
			if err != nil {
				t.Fatalf("error marshaling last applied config: %v", err)
			}
			r.SetAnnotations(map[string]string{
				k8s.LastAppliedConfigurationAnnotation: string(lastAppliedConfigJSON),
			})
			var liveState *terraform.InstanceState
			if tc.resourceExists {
				// The content of the instance state does not matter here,
				// as the diff function later handles the merge. We just
				// need to supply *something* to mark this resource
				// as already existing.
				liveState = &terraform.InstanceState{
					ID: "foo",
				}
			}
			config := tc.inputConfig
			if err := ResolveLegacyGCPManagedFields(r, liveState, config); err != nil {
				t.Fatalf("error resolving GCP-managed fields: %v", err)
			}
			if !test.Equals(t, tc.expectedConfig, config) {
				t.Fatalf("expected config: %+v, actual: %+v", tc.expectedConfig, config)
			}
		})
	}
}
