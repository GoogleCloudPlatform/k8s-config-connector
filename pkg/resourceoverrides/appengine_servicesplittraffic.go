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
	"strconv"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides/operations"
)

const appEngineMigrateTrafficField = "migrate_traffic"

func GetAppEngineServiceSplitTrafficResourceOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "AppEngineServiceSplitTraffic",
	}
	ro.Overrides = append(ro.Overrides, useDesiredMigrateTraffic())
	return ro
}

// useDesiredMigrateTraffic stops a change to spec.migrateTraffic on its own from
// triggering an update.
//
// migrateTraffic is a request-only parameter. It is sent as the migrateTraffic query
// parameter of UpdateService, it decides whether a change to spec.split is rolled out
// gradually or at once, and App Engine never stores or returns it. The Terraform
// controller rebuilds the live state from App Engine on every reconcile, so the value can
// never be read back. Copying the desired value into the live state means that:
//   - a change to migrateTraffic alone produces no diff and no API call. Nothing on the
//     service changes, and App Engine rejects migrateTraffic=true without a split change.
//   - when spec.split changes, the update is sent with the desired migrateTraffic value.
//     When spec.migrateTraffic is omitted, false is used, which is the API default, so
//     the update is sent with migrateTraffic=false rather than an empty value.
//
// mutableButUnreadableFields is not used because it keeps the last applied value, so a
// change to migrateTraffic alone would still produce a diff and an empty update.
func useDesiredMigrateTraffic() ResourceOverride {
	o := ResourceOverride{}
	o.PreTerraformApply = func(ctx context.Context, op *operations.PreTerraformApply) error {
		// On create there is no live state, and Terraform sends the desired value as is.
		if op.LiveState.Empty() || op.TerraformConfig == nil {
			return nil
		}

		value := false
		if desired, found := op.TerraformConfig.Get(appEngineMigrateTrafficField); found && desired != nil {
			v, ok := desired.(bool)
			if !ok {
				// Not a concrete bool (e.g. an unknown value); leave the live state as is.
				return nil
			}
			value = v
		}

		if op.LiveState.Attributes == nil {
			op.LiveState.Attributes = map[string]string{}
		}
		op.LiveState.Attributes[appEngineMigrateTrafficField] = strconv.FormatBool(value)
		return nil
	}
	return o
}
