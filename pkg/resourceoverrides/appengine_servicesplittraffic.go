// Copyright 2026 Google LLC
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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides/operations"
)

const (
	appEngineSplitShardByField          = "split.0.shard_by"
	appEngineSplitAllocationsCountField = "split.0.allocations.%"
)

func GetAppEngineServiceSplitTrafficResourceOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "AppEngineServiceSplitTraffic",
	}
	ro.Overrides = append(ro.Overrides, keepShardByDroppedByAppEngine())
	return ro
}

// keepShardByDroppedByAppEngine stops a shardBy that App Engine dropped from looking like a
// change.
//
// When an update that is not a traffic migration leaves only one version in the split, App
// Engine clears shardBy ("the diversion method is irrelevant"), so it is never read back. If
// the live split has one version and no shardBy, the desired shardBy is copied into the live
// state. A shardBy-only difference then produces no diff and no update, while a change to the
// allocations still updates the split, and the update still sends the desired shardBy.
func keepShardByDroppedByAppEngine() ResourceOverride {
	o := ResourceOverride{}
	o.PreTerraformApply = func(ctx context.Context, op *operations.PreTerraformApply) error {
		// On create there is no live state, and Terraform sends the desired value as is.
		if op.LiveState.Empty() || op.TerraformConfig == nil {
			return nil
		}

		live := op.LiveState.Attributes
		if live[appEngineSplitAllocationsCountField] != "1" || live[appEngineSplitShardByField] != "" {
			return nil
		}
		if op.TerraformConfig.IsComputed(appEngineSplitShardByField) {
			return nil
		}
		desired, found := op.TerraformConfig.Get(appEngineSplitShardByField)
		if !found {
			return nil
		}
		shardBy, ok := desired.(string)
		if !ok || shardBy == "" {
			return nil
		}
		live[appEngineSplitShardByField] = shardBy
		return nil
	}
	return o
}
