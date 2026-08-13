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
)

// defaultEnvironmentPb returns a proto Environment populated with static server-defaulted fields.
// It serves as the centralized baseline template for universal static defaults defined in
// docs/resource-defaults/composerenvironment.md.
//
// Universal static defaults are deterministic, fixed constants applied across environments when the
// corresponding spec field is unset.
func defaultEnvironmentPb() *composerpb.Environment {
	return &composerpb.Environment{
		Config: &composerpb.EnvironmentConfig{
			// Mutable: Environment size defaults to SMALL (1).
			EnvironmentSize: composerpb.EnvironmentConfig_ENVIRONMENT_SIZE_SMALL,

			// Mutable: WorkloadsConfig scheduler baseline resources for small environment size.
			WorkloadsConfig: &composerpb.WorkloadsConfig{
				Scheduler: &composerpb.WorkloadsConfig_SchedulerResource{
					Count:     1,
					Cpu:       0.5,
					MemoryGb:  2,
					StorageGb: 1,
				},
			},

			// Mutable: WebServerNetworkAccessControl defaults to open access for IPv4 and IPv6.
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

			// Mutable: MaintenanceWindow defaults to a weekly weekend schedule if unspecified.
			MaintenanceWindow: &composerpb.MaintenanceWindow{
				Recurrence: "FREQ=WEEKLY;BYDAY=FR,SA,SU",
			},

			// Immutable: Private environment networkingConfig default.
			PrivateEnvironmentConfig: &composerpb.PrivateEnvironmentConfig{
				NetworkingConfig: &composerpb.NetworkingConfig{},
			},
		},
	}
}

// defaultEnvironmentPb returns the default proto Environment template.
func (a *EnvironmentAdapter) defaultEnvironmentPb() *composerpb.Environment {
	return defaultEnvironmentPb()
}

// populateDesiredWithDefaults populates configurable fields in desiredPb that have static
// server-defaulted values using the centralized defaultEnvironmentPb template as a baseline.
//
// These fields represent known, deterministic constants that GCP uses when a user omits them from spec.
func (a *EnvironmentAdapter) populateDesiredWithDefaults(desired *krm.ComposerEnvironment, desiredPb *composerpb.Environment) {
	if desiredPb == nil {
		return
	}

	cfg := desired.Spec.Config
	if cfg == nil {
		cfg = &krm.EnvironmentConfig{}
	}
	peConfig := cfg.PrivateEnvironmentConfig

	if desiredPb.Config == nil {
		desiredPb.Config = &composerpb.EnvironmentConfig{}
	}

	// 1. Static Server-Defaulted Fields
	defaultConfig := a.defaultEnvironmentPb().GetConfig()

	// Mutable: environmentSize defaults to ENVIRONMENT_SIZE_SMALL (1)
	if cfg.EnvironmentSize == nil {
		desiredPb.Config.EnvironmentSize = defaultConfig.GetEnvironmentSize()
	}

	// Mutable: workloadsConfig.scheduler static sizing default
	if cfg.WorkloadsConfig == nil || cfg.WorkloadsConfig.Scheduler == nil {
		if desiredPb.Config.WorkloadsConfig == nil {
			desiredPb.Config.WorkloadsConfig = &composerpb.WorkloadsConfig{}
		}
		desiredPb.Config.WorkloadsConfig.Scheduler = defaultConfig.GetWorkloadsConfig().GetScheduler()
	}

	// Mutable: webServerNetworkAccessControl static default (open access 0.0.0.0/0, ::0/0)
	if cfg.WebServerNetworkAccessControl == nil {
		desiredPb.Config.WebServerNetworkAccessControl = defaultConfig.GetWebServerNetworkAccessControl()
	}

	// Mutable: maintenanceWindow static weekly schedule default
	if cfg.MaintenanceWindow == nil {
		desiredPb.Config.MaintenanceWindow = defaultConfig.GetMaintenanceWindow()
	}

	// Immutable: privateEnvironmentConfig.networkingConfig
	if peConfig == nil || peConfig.NetworkingConfig == nil {
		if desiredPb.Config.PrivateEnvironmentConfig == nil {
			desiredPb.Config.PrivateEnvironmentConfig = &composerpb.PrivateEnvironmentConfig{}
		}
		desiredPb.Config.PrivateEnvironmentConfig.NetworkingConfig = defaultConfig.GetPrivateEnvironmentConfig().GetNetworkingConfig()
	}
}

// populateDesiredWithActualIfComputed populates dynamic/computed server-generated values by copying
// them from actualPb into desiredPb when the corresponding field is omitted in the desired spec.
//
// These fields represent runtime server-assigned values that are generated or auto-allocated by GCP
// during environment creation or drifted by service upgrades (e.g. auto-created GCS bucket name,
// auto-allocated CIDR blocks, dynamically upgraded image versions, generation-dependent node VM/database configs).
func (a *EnvironmentAdapter) populateDesiredWithActualIfComputed(desired *krm.ComposerEnvironment, desiredPb, actualPb *composerpb.Environment) {
	if desiredPb == nil || actualPb == nil {
		return
	}

	cfg := desired.Spec.Config
	if cfg == nil {
		cfg = &krm.EnvironmentConfig{}
	}
	softwareConfig := cfg.SoftwareConfig
	dbConfig := cfg.DatabaseConfig
	nodeConfig := cfg.NodeConfig
	workloadsConfig := cfg.WorkloadsConfig

	// Immutable: storageConfig.bucket (auto-created GCS bucket)
	if actualPb.GetStorageConfig() != nil {
		if desiredPb.StorageConfig == nil {
			desiredPb.StorageConfig = &composerpb.StorageConfig{}
		}
		if desired.Spec.StorageConfig == nil || desired.Spec.StorageConfig.BucketRef == nil {
			desiredPb.StorageConfig.Bucket = actualPb.GetStorageConfig().GetBucket()
		}
	}

	actualConfig := actualPb.GetConfig()
	if actualConfig == nil {
		return
	}

	if desiredPb.Config == nil {
		desiredPb.Config = &composerpb.EnvironmentConfig{}
	}

	// Dynamic NodeConfig fields (auto-allocated CIDR, subnetwork, network attachment, or server-assigned network/machineType/diskSize)
	if actualNode := actualConfig.GetNodeConfig(); actualNode != nil {
		if desiredPb.Config.NodeConfig == nil {
			desiredPb.Config.NodeConfig = &composerpb.NodeConfig{}
		}
		if nodeConfig == nil || nodeConfig.ComposerInternalIPv4CIDRBlock == nil {
			desiredPb.Config.NodeConfig.ComposerInternalIpv4CidrBlock = actualNode.GetComposerInternalIpv4CidrBlock()
		}
		if nodeConfig == nil || nodeConfig.ComposerNetworkAttachmentRef == nil {
			desiredPb.Config.NodeConfig.ComposerNetworkAttachment = actualNode.GetComposerNetworkAttachment()
		}
		if nodeConfig == nil || nodeConfig.SubnetworkRef == nil {
			desiredPb.Config.NodeConfig.Subnetwork = actualNode.GetSubnetwork()
		}
		if nodeConfig == nil || nodeConfig.IPAllocationPolicy == nil {
			desiredPb.Config.NodeConfig.IpAllocationPolicy = actualNode.GetIpAllocationPolicy()
		}
		if nodeConfig == nil || nodeConfig.NetworkRef == nil {
			desiredPb.Config.NodeConfig.Network = actualNode.GetNetwork()
		}
		if nodeConfig == nil || nodeConfig.MachineType == nil {
			desiredPb.Config.NodeConfig.MachineType = actualNode.GetMachineType()
		}
		if nodeConfig == nil || nodeConfig.DiskSizeGB == nil {
			desiredPb.Config.NodeConfig.DiskSizeGb = actualNode.GetDiskSizeGb()
		}
	}

	// Dynamic SoftwareConfig fields (dynamic imageVersion, pythonVersion, schedulerCount, webServerPluginsMode)
	if actualSoft := actualConfig.GetSoftwareConfig(); actualSoft != nil {
		if desiredPb.Config.SoftwareConfig == nil {
			desiredPb.Config.SoftwareConfig = &composerpb.SoftwareConfig{}
		}
		if softwareConfig == nil || softwareConfig.ImageVersion == nil {
			desiredPb.Config.SoftwareConfig.ImageVersion = actualSoft.GetImageVersion()
		}
		if softwareConfig == nil || softwareConfig.PythonVersion == nil {
			desiredPb.Config.SoftwareConfig.PythonVersion = actualSoft.GetPythonVersion()
		}
		if softwareConfig == nil || softwareConfig.SchedulerCount == nil {
			desiredPb.Config.SoftwareConfig.SchedulerCount = actualSoft.GetSchedulerCount()
		}
		if softwareConfig == nil || softwareConfig.WebServerPluginsMode == nil {
			desiredPb.Config.SoftwareConfig.WebServerPluginsMode = actualSoft.GetWebServerPluginsMode()
		}
	}

	// Dynamic DatabaseConfig fields (dynamic machineType, zone)
	if actualDb := actualConfig.GetDatabaseConfig(); actualDb != nil {
		if desiredPb.Config.DatabaseConfig == nil {
			desiredPb.Config.DatabaseConfig = &composerpb.DatabaseConfig{}
		}
		if dbConfig == nil || dbConfig.MachineType == nil {
			desiredPb.Config.DatabaseConfig.MachineType = actualDb.GetMachineType()
		}
		if dbConfig == nil || dbConfig.Zone == nil {
			desiredPb.Config.DatabaseConfig.Zone = actualDb.GetZone()
		}
	}

	// Mutable: dataRetentionConfig (dynamic default in Composer 3)
	if cfg.DataRetentionConfig == nil {
		desiredPb.Config.DataRetentionConfig = actualConfig.GetDataRetentionConfig()
	}

	// Mutable: workloadsConfig (dagProcessor, triggerer, webServer, worker dynamic sizing)
	if actualWorkloads := actualConfig.GetWorkloadsConfig(); actualWorkloads != nil {
		if desiredPb.Config.WorkloadsConfig == nil {
			desiredPb.Config.WorkloadsConfig = &composerpb.WorkloadsConfig{}
		}
		if workloadsConfig == nil || workloadsConfig.DagProcessor == nil {
			desiredPb.Config.WorkloadsConfig.DagProcessor = actualWorkloads.GetDagProcessor()
		}
		if workloadsConfig == nil || workloadsConfig.Triggerer == nil {
			desiredPb.Config.WorkloadsConfig.Triggerer = actualWorkloads.GetTriggerer()
		}
		if workloadsConfig == nil || workloadsConfig.WebServer == nil {
			desiredPb.Config.WorkloadsConfig.WebServer = actualWorkloads.GetWebServer()
		}
		if workloadsConfig == nil || workloadsConfig.Worker == nil {
			desiredPb.Config.WorkloadsConfig.Worker = actualWorkloads.GetWorker()
		}
	}
}
