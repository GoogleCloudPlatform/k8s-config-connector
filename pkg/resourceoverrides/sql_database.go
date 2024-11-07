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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides/operations"
)

func GetSQLDatabaseResourceOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "SQLDatabase",
	}

	ro.Overrides = append(ro.Overrides, ignoreChangesToTerraformVirtualFields())

	return ro
}

func ignoreChangesToTerraformVirtualFields() ResourceOverride {
	o := ResourceOverride{}
	o.PreTerraformApply = func(ctx context.Context, op *operations.PreTerraformApply) error {
		// This field is only used in the deletion path, and otherwise causes spurious diffs.
		delete(op.TerraformConfig.Config, "deletion_policy")

		return nil
	}
	return o
}
