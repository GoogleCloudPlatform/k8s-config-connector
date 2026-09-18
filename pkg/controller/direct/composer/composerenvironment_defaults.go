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

package composer

import (
	composerpb "cloud.google.com/go/orchestration/airflow/service/apiv1/servicepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/composer/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
)

// defaultEnvironmentPb returns a proto Environment populated with static server-defaulted fields.
// It serves as the centralized baseline template for universal static defaults defined in
// docs/resource-behaviors/composerenvironment/default-value-analysis.md.
//
// Universal static defaults are deterministic, fixed constants applied across environments when the
// corresponding spec field is unset.
func defaultEnvironmentPb() *composerpb.Environment {
	return &composerpb.Environment{
		Config: &composerpb.EnvironmentConfig{
			// Mutable: WebServerNetworkAccessControl defaults to open access for IPv4 and IPv6 across all environments.
			WebServerNetworkAccessControl: &composerpb.WebServerNetworkAccessControl{
				AllowedIpRanges: []*composerpb.WebServerNetworkAccessControl_AllowedIpRange{
					{
						Value:       "0.0.0.0/0",
						Description: "Allows access from all IPv4 addresses (default value)",
					},
					{
						Value:       "::0/0",
						Description: "Allows access from all IPv6 addresses (default value)",
					},
				},
			},
		},
	}
}

// populateDesiredWithDefaults populates configurable fields in desiredPb that have static
// server-defaulted values using the centralized defaultEnvironmentPb template as a baseline.
//
// These fields represent known, deterministic constants that GCP uses when a user omits them from spec.
func populateDesiredWithDefaults(desired *krm.ComposerEnvironment, desiredPb *composerpb.Environment) {
	if desiredPb == nil {
		return
	}

	cfg := desired.Spec.Config
	if cfg == nil {
		cfg = &krm.EnvironmentConfig{}
	}

	if desiredPb.Config == nil {
		desiredPb.Config = &composerpb.EnvironmentConfig{}
	}

	// 1. Static Server-Defaulted Fields
	defaultConfig := defaultEnvironmentPb().GetConfig()

	// Mutable: webServerNetworkAccessControl static default (open access 0.0.0.0/0, ::0/0)
	if cfg.WebServerNetworkAccessControl == nil {
		desiredPb.Config.WebServerNetworkAccessControl = defaultConfig.GetWebServerNetworkAccessControl()
	}
}

// computedFieldPaths lists the KRM field paths for server-assigned values.
// Composite sub-messages whose fields are 100% server-computed (or immutable) are listed by their
// parent sub-message path; PopulateComputedFields automatically recurses into partially specified sub-messages.
// Sub-messages containing user-clearable fields (like SoftwareConfig) list only their server-computed leaf fields.
var computedFieldPaths = []string{
	// 1. StorageConfig
	"StorageConfig",

	// 2. Config top-level dynamic fields & composite sub-messages (sorted alphabetically)
	"Config.DatabaseConfig",
	"Config.DataRetentionConfig",
	"Config.EnvironmentSize",
	"Config.MaintenanceWindow",
	"Config.MasterAuthorizedNetworksConfig",
	"Config.NodeConfig",
	"Config.NodeCount",
	"Config.PrivateEnvironmentConfig",
	"Config.RecoveryConfig",
	"Config.WebServerConfig",
	"Config.WorkloadsConfig",

	// 3. SoftwareConfig (explicitly list server-computed leaves and optional feature sub-messages so user-clearable maps can be unset)
	"Config.SoftwareConfig.CloudDataLineageIntegration",
	"Config.SoftwareConfig.ImageVersion",
	"Config.SoftwareConfig.PythonVersion",
	"Config.SoftwareConfig.SchedulerCount",
	"Config.SoftwareConfig.WebServerPluginsMode",
}

// populateDesiredWithActualIfComputed populates dynamic/server-generated values in O(N) linear time
// by checking omitted KRM fields against non-nil parents in actualPb and directly assigning their values to desiredPb.
func populateDesiredWithActualIfComputed(desired *krm.ComposerEnvironment, desiredPb, actualPb *composerpb.Environment) {
	if desired == nil {
		common.PopulateComputedFields(nil, desiredPb, actualPb, computedFieldPaths)
		normalizeDisabledFeatureBlocks(desiredPb, actualPb)
		return
	}
	common.PopulateComputedFields(desired.Spec, desiredPb, actualPb, computedFieldPaths)
	normalizeDisabledFeatureBlocks(desiredPb, actualPb)
}

// normalizeDisabledFeatureBlocks aligns desiredPb with actualPb when optional boolean-toggle
// blocks (CloudDataLineageIntegration, ScheduledSnapshotsConfig, MasterAuthorizedNetworksConfig) are
// explicitly set to enabled: false in spec while GCP returns nil (absent message) in actualPb.
// On GCP, an absent sub-message (nil) is semantically identical to {enabled: false}. Normalizing
// desiredPb to nil declaratively before diffing avoids false diffs in both validateUpdatableFields
// and fieldUpdaters without requiring procedural nil-guards inside individual fieldUpdaters.
func normalizeDisabledFeatureBlocks(desiredPb, actualPb *composerpb.Environment) {
	if desiredPb == nil || desiredPb.Config == nil {
		return
	}

	// 1. CloudDataLineageIntegration
	if desiredPb.Config.SoftwareConfig != nil {
		mergedLineage := desiredPb.Config.SoftwareConfig.GetCloudDataLineageIntegration()
		actualLineage := actualPb.GetConfig().GetSoftwareConfig().GetCloudDataLineageIntegration()
		if mergedLineage != nil && !mergedLineage.GetEnabled() && actualLineage == nil {
			desiredPb.Config.SoftwareConfig.CloudDataLineageIntegration = nil
		}
	}

	// 2. ScheduledSnapshotsConfig
	if desiredPb.Config.RecoveryConfig != nil {
		mergedSnapshots := desiredPb.Config.RecoveryConfig.GetScheduledSnapshotsConfig()
		actualSnapshots := actualPb.GetConfig().GetRecoveryConfig().GetScheduledSnapshotsConfig()
		if mergedSnapshots != nil && !mergedSnapshots.GetEnabled() &&
			mergedSnapshots.GetSnapshotLocation() == "" &&
			mergedSnapshots.GetSnapshotCreationSchedule() == "" &&
			mergedSnapshots.GetTimeZone() == "" &&
			actualSnapshots == nil {
			desiredPb.Config.RecoveryConfig.ScheduledSnapshotsConfig = nil
			if actualPb.GetConfig().GetRecoveryConfig() == nil {
				desiredPb.Config.RecoveryConfig = nil
			}
		}
	}

	// 3. MasterAuthorizedNetworksConfig
	mergedAuth := desiredPb.Config.GetMasterAuthorizedNetworksConfig()
	actualAuth := actualPb.GetConfig().GetMasterAuthorizedNetworksConfig()
	if mergedAuth != nil && !mergedAuth.GetEnabled() && len(mergedAuth.GetCidrBlocks()) == 0 && actualAuth == nil {
		desiredPb.Config.MasterAuthorizedNetworksConfig = nil
	}
}
