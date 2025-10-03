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
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides/operations"
)

func GetBigtableInstanceOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "BigtableInstance",
	}
	ro.Overrides = append(ro.Overrides, noNodesWhereAutoscaling())
	return ro
}

func noNodesWhereAutoscaling() ResourceOverride {
	o := ResourceOverride{}

	// When we are applying a cluster which has autoscaling, we should not specify num_nodes (serve_nodes in GCP)
	o.PreTerraformApply = func(ctx context.Context, op *operations.PreTerraformApply) error {
		removeNumNodesIfAutoscaling := func(clusters []any) {
			for i, clusterAny := range clusters {
				cluster, ok := clusterAny.(map[string]any)
				if !ok {
					continue
				}

				autoscalingConfig := cluster["autoscaling_config"]
				if autoscalingConfig == nil {
					continue
				}

				delete(cluster, "num_nodes")

				// Delete from liveState also, otherwise TF gets confused
				if op.LiveState != nil {
					k := fmt.Sprintf("cluster.%d.num_nodes", i)
					delete(op.LiveState.Attributes, k)
				}
			}
		}

		if clusters, ok := op.TerraformConfig.Config["cluster"].([]any); ok {
			removeNumNodesIfAutoscaling(clusters)
		}

		if clusters, ok := op.TerraformConfig.Raw["cluster"].([]any); ok {
			removeNumNodesIfAutoscaling(clusters)
		}

		return nil
	}

	return o
}
