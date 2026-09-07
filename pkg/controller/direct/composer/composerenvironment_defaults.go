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
	"strings"

	composerpb "cloud.google.com/go/orchestration/airflow/service/apiv1/servicepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/composer/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
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
var computedFieldPaths = []string{
	// 1. StorageConfig
	"StorageConfig.BucketRef",

	// 2. Config top-level dynamic fields
	"Config.EnvironmentSize",
	"Config.NodeCount",
	"Config.MaintenanceWindow",
	"Config.DataRetentionConfig",

	// 3. NodeConfig
	"Config.NodeConfig.ComposerInternalIPv4CIDRBlock",
	"Config.NodeConfig.ComposerNetworkAttachmentRef",
	"Config.NodeConfig.SubnetworkRef",
	"Config.NodeConfig.IPAllocationPolicy",
	"Config.NodeConfig.NetworkRef",
	"Config.NodeConfig.MachineType",
	"Config.NodeConfig.DiskSizeGB",

	// 4. SoftwareConfig
	"Config.SoftwareConfig.ImageVersion",
	"Config.SoftwareConfig.PythonVersion",
	"Config.SoftwareConfig.SchedulerCount",
	"Config.SoftwareConfig.WebServerPluginsMode",

	// 5. DatabaseConfig
	"Config.DatabaseConfig.MachineType",
	"Config.DatabaseConfig.Zone",

	// 6. PrivateEnvironmentConfig
	"Config.PrivateEnvironmentConfig.CloudComposerNetworkIPv4CIDRBlock",
	"Config.PrivateEnvironmentConfig.CloudSQLIPv4CIDRBlock",
	"Config.PrivateEnvironmentConfig.PrivateClusterConfig",
	"Config.PrivateEnvironmentConfig.WebServerIPv4CIDRBlock",
	"Config.PrivateEnvironmentConfig.CloudComposerConnectionSubnetworkRef",
	"Config.PrivateEnvironmentConfig.NetworkingConfig",

	// 7. WorkloadsConfig
	"Config.WorkloadsConfig.Scheduler",
	"Config.WorkloadsConfig.Scheduler.CPU",
	"Config.WorkloadsConfig.Scheduler.MemoryGB",
	"Config.WorkloadsConfig.Scheduler.StorageGB",
	"Config.WorkloadsConfig.Scheduler.Count",
	"Config.WorkloadsConfig.DagProcessor",
	"Config.WorkloadsConfig.DagProcessor.CPU",
	"Config.WorkloadsConfig.DagProcessor.MemoryGB",
	"Config.WorkloadsConfig.DagProcessor.StorageGB",
	"Config.WorkloadsConfig.DagProcessor.Count",
	"Config.WorkloadsConfig.Triggerer",
	"Config.WorkloadsConfig.Triggerer.CPU",
	"Config.WorkloadsConfig.Triggerer.MemoryGB",
	"Config.WorkloadsConfig.Triggerer.Count",
	"Config.WorkloadsConfig.WebServer",
	"Config.WorkloadsConfig.WebServer.CPU",
	"Config.WorkloadsConfig.WebServer.MemoryGB",
	"Config.WorkloadsConfig.WebServer.StorageGB",
	"Config.WorkloadsConfig.Worker",
	"Config.WorkloadsConfig.Worker.CPU",
	"Config.WorkloadsConfig.Worker.MemoryGB",
	"Config.WorkloadsConfig.Worker.StorageGB",
	"Config.WorkloadsConfig.Worker.MinCount",
	"Config.WorkloadsConfig.Worker.MaxCount",

	// 8. WebServerConfig
	"Config.WebServerConfig.MachineType",
}

// populateDesiredWithActualIfComputed populates dynamic/computed server-generated values in O(N) linear time
// by checking omitted KRM fields against non-nil parents in actualPb and directly assigning their values to desiredPb.
func populateDesiredWithActualIfComputed(desired *krm.ComposerEnvironment, desiredPb, actualPb *composerpb.Environment) {
	if desiredPb == nil || actualPb == nil {
		return
	}

	// 1. Build map of non-nil parents from actualPb and initialize them on desiredPb
	parentMap := common.BuildParentMap(desiredPb, actualPb, computedFieldPaths)

	// 2. Collect all paths explicitly set in desired.Spec in a single O(N) pass
	var presentFields sets.Set[string]
	if desired != nil {
		presentFields = common.CollectPresentFields(desired.Spec)
	} else {
		presentFields = sets.New[string]()
	}

	// 3. For any computed field omitted in desired.Spec, copy from actualPb if its parent exists
	for _, path := range computedFieldPaths {
		if presentFields.Has(path) {
			continue
		}

		lastDot := strings.LastIndex(path, ".")
		var parentPath, leafName string
		if lastDot == -1 {
			parentPath = ""
			leafName = path
		} else {
			parentPath = path[:lastDot]
			leafName = path[lastDot+1:]
		}

		pair, ok := parentMap[parentPath]
		if !ok {
			continue
		}

		fd := common.FindProtoField(pair.Actual.Descriptor(), leafName)
		if fd == nil {
			klog.V(0).Infof("internal error: field %q not found on proto message %s", leafName, pair.Actual.Descriptor().FullName())
			continue
		}
		if pair.Actual.Has(fd) {
			pair.Desired.Set(fd, pair.Actual.Get(fd))
		}
	}
}
