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
	"google.golang.org/protobuf/reflect/protoreflect"
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
	"Config.WorkloadsConfig.DagProcessor",
	"Config.WorkloadsConfig.Triggerer",
	"Config.WorkloadsConfig.WebServer",
	"Config.WorkloadsConfig.Worker",

	// 8. WebServerConfig
	"Config.WebServerConfig.MachineType",
}

// parentPair tracks matching actual and desired proto messages for a given parent path.
type parentPair struct {
	actual  protoreflect.Message
	desired protoreflect.Message
}

// buildParentMap collects all non-nil parent sub-messages from actualPb and initializes
// corresponding message containers on desiredPb.
func buildParentMap(desiredPb, actualPb *composerpb.Environment) map[string]parentPair {
	parents := make(map[string]parentPair)
	if actualPb == nil || desiredPb == nil {
		return parents
	}

	parents[""] = parentPair{actual: actualPb.ProtoReflect(), desired: desiredPb.ProtoReflect()}

	if actualPb.StorageConfig != nil {
		if desiredPb.StorageConfig == nil {
			desiredPb.StorageConfig = &composerpb.StorageConfig{}
		}
		parents["StorageConfig"] = parentPair{
			actual:  actualPb.StorageConfig.ProtoReflect(),
			desired: desiredPb.StorageConfig.ProtoReflect(),
		}
	}

	if actualPb.Config != nil {
		if desiredPb.Config == nil {
			desiredPb.Config = &composerpb.EnvironmentConfig{}
		}
		parents["Config"] = parentPair{
			actual:  actualPb.Config.ProtoReflect(),
			desired: desiredPb.Config.ProtoReflect(),
		}

		cfg := actualPb.Config
		dCfg := desiredPb.Config

		if cfg.NodeConfig != nil {
			if dCfg.NodeConfig == nil {
				dCfg.NodeConfig = &composerpb.NodeConfig{}
			}
			parents["Config.NodeConfig"] = parentPair{
				actual:  cfg.NodeConfig.ProtoReflect(),
				desired: dCfg.NodeConfig.ProtoReflect(),
			}
		}
		if cfg.SoftwareConfig != nil {
			if dCfg.SoftwareConfig == nil {
				dCfg.SoftwareConfig = &composerpb.SoftwareConfig{}
			}
			parents["Config.SoftwareConfig"] = parentPair{
				actual:  cfg.SoftwareConfig.ProtoReflect(),
				desired: dCfg.SoftwareConfig.ProtoReflect(),
			}
		}
		if cfg.DatabaseConfig != nil {
			if dCfg.DatabaseConfig == nil {
				dCfg.DatabaseConfig = &composerpb.DatabaseConfig{}
			}
			parents["Config.DatabaseConfig"] = parentPair{
				actual:  cfg.DatabaseConfig.ProtoReflect(),
				desired: dCfg.DatabaseConfig.ProtoReflect(),
			}
		}
		if cfg.WebServerConfig != nil {
			if dCfg.WebServerConfig == nil {
				dCfg.WebServerConfig = &composerpb.WebServerConfig{}
			}
			parents["Config.WebServerConfig"] = parentPair{
				actual:  cfg.WebServerConfig.ProtoReflect(),
				desired: dCfg.WebServerConfig.ProtoReflect(),
			}
		}
		if cfg.PrivateEnvironmentConfig != nil {
			if dCfg.PrivateEnvironmentConfig == nil {
				dCfg.PrivateEnvironmentConfig = &composerpb.PrivateEnvironmentConfig{}
			}
			parents["Config.PrivateEnvironmentConfig"] = parentPair{
				actual:  cfg.PrivateEnvironmentConfig.ProtoReflect(),
				desired: dCfg.PrivateEnvironmentConfig.ProtoReflect(),
			}
		}
		if cfg.WorkloadsConfig != nil {
			if dCfg.WorkloadsConfig == nil {
				dCfg.WorkloadsConfig = &composerpb.WorkloadsConfig{}
			}
			parents["Config.WorkloadsConfig"] = parentPair{
				actual:  cfg.WorkloadsConfig.ProtoReflect(),
				desired: dCfg.WorkloadsConfig.ProtoReflect(),
			}
		}
	}

	return parents
}

// findProtoField matches a KRM leaf field name to its corresponding proto field descriptor
// by comparing the field's JSONName against the KRM field name (ignoring case).
func findProtoField(desc protoreflect.MessageDescriptor, krmLeaf string) protoreflect.FieldDescriptor {
	krmName := strings.TrimSuffix(krmLeaf, "Ref")
	for i := 0; i < desc.Fields().Len(); i++ {
		fd := desc.Fields().Get(i)
		if strings.EqualFold(fd.JSONName(), krmName) {
			return fd
		}
	}
	return nil
}

// populateDesiredWithActualIfComputed populates dynamic/computed server-generated values in O(N) linear time
// by checking omitted KRM fields against non-nil parents in actualPb and directly assigning their values to desiredPb.
func populateDesiredWithActualIfComputed(desired *krm.ComposerEnvironment, desiredPb, actualPb *composerpb.Environment) {
	if desiredPb == nil || actualPb == nil {
		return
	}

	// 1. Build map of non-nil parents from actualPb and initialize them on desiredPb
	parentMap := buildParentMap(desiredPb, actualPb)

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

		fd := findProtoField(pair.actual.Descriptor(), leafName)
		if fd == nil {
			klog.V(0).Infof("internal error: field %q not found on proto message %s", leafName, pair.actual.Descriptor().FullName())
			continue
		}
		if pair.actual.Has(fd) {
			pair.desired.Set(fd, pair.actual.Get(fd))
		}
	}
}
