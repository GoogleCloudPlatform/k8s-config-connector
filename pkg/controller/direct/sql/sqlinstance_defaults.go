// Copyright 2024 Google LLC
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

package sql

import (
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	api "google.golang.org/api/sqladmin/v1beta4"
)

const DefaultToGCPFieldsAnnotation = "cnrm.cloud.google.com/default-to-gcp-fields"

func ApplySQLInstanceGCPDefaults(in *krm.SQLInstance, out *api.DatabaseInstance, actual *api.DatabaseInstance) {
	fieldsToDefault := make(map[string]bool)
	if actual != nil {
		if fieldsValue, ok := in.GetAnnotations()[DefaultToGCPFieldsAnnotation]; ok && fieldsValue != "" {
			fields := strings.Split(fieldsValue, ",")
			for _, field := range fields {
				fieldsToDefault[strings.TrimSpace(field)] = true
			}
		}
	}

	if in.Spec.DatabaseVersion == nil {
		if fieldsToDefault["databaseVersion"] {
			out.DatabaseVersion = actual.DatabaseVersion
		}
	}
	if in.Spec.InstanceType == nil {
		if fieldsToDefault["instanceType"] {
			out.InstanceType = actual.InstanceType
		}
	}
	if in.Spec.MaintenanceVersion == nil {
		if fieldsToDefault["maintenanceVersion"] {
			out.MaintenanceVersion = actual.MaintenanceVersion
		}
	}
	if in.Spec.MasterInstanceRef == nil {
		if fieldsToDefault["masterInstanceName"] {
			out.MasterInstanceName = actual.MasterInstanceName
		}
	}
	if in.Spec.Region == nil {
		if fieldsToDefault["region"] {
			out.Region = actual.Region
		}
	}

	// Settings
	if out.Settings == nil {
		out.Settings = &api.Settings{}
	}
	if in.Spec.Settings.ActivationPolicy == nil {
		if fieldsToDefault["settings.activationPolicy"] {
			out.Settings.ActivationPolicy = actual.Settings.ActivationPolicy
		}
	}
	if in.Spec.Settings.AuthorizedGaeApplications == nil {
		// For some reason, GCP API uses empty slice instead of nil.
		out.Settings.AuthorizedGaeApplications = make([]string, 0)
	}
	if in.Spec.Settings.AvailabilityType == nil {
		if fieldsToDefault["settings.availabilityType"] {
			out.Settings.AvailabilityType = actual.Settings.AvailabilityType
		}
	}
	if in.Spec.Settings.BackupConfiguration == nil && actual != nil && actual.Settings.BackupConfiguration != nil && !actual.Settings.BackupConfiguration.Enabled {
		out.Settings.BackupConfiguration = actual.Settings.BackupConfiguration
	}
	if in.Spec.Settings.BackupConfiguration != nil {
		if in.Spec.Settings.BackupConfiguration.BackupRetentionSettings != nil && in.Spec.Settings.BackupConfiguration.BackupRetentionSettings.RetentionUnit == nil {
			if fieldsToDefault["settings.backupConfiguration.backupRetentionSettings.retentionUnit"] {
				// default to gcp as user specified
				out.Settings.BackupConfiguration.BackupRetentionSettings.RetentionUnit = actual.Settings.BackupConfiguration.BackupRetentionSettings.RetentionUnit
			}
		}
	}
	if in.Spec.Settings.Collation == nil {
		if fieldsToDefault["settings.collation"] {
			out.Settings.Collation = actual.Settings.Collation
		}
	}
	if in.Spec.Settings.ConnectorEnforcement == nil {
		if fieldsToDefault["settings.connectorEnforcement"] {
			out.Settings.ConnectorEnforcement = actual.Settings.ConnectorEnforcement
		}
	}
	if in.Spec.Settings.CrashSafeReplication == nil {
		if fieldsToDefault["settings.crashSafeReplication"] {
			out.Settings.CrashSafeReplicationEnabled = actual.Settings.CrashSafeReplicationEnabled
		}
	}
	if in.Spec.Settings.DiskAutoresize == nil {
		if fieldsToDefault["settings.diskAutoresize"] {
			out.Settings.StorageAutoResize = actual.Settings.StorageAutoResize
		}
	}
	if in.Spec.Settings.DiskAutoresizeLimit == nil {
		if fieldsToDefault["settings.diskAutoresizeLimit"] {
			out.Settings.StorageAutoResizeLimit = actual.Settings.StorageAutoResizeLimit
		}
	}
	if in.Spec.Settings.DiskSize == nil {
		if fieldsToDefault["settings.diskSize"] {
			out.Settings.DataDiskSizeGb = actual.Settings.DataDiskSizeGb
		} else if actual != nil && out.Settings.StorageAutoResize != nil && *out.Settings.StorageAutoResize {
			out.Settings.DataDiskSizeGb = actual.Settings.DataDiskSizeGb
		}
	}
	if in.Spec.Settings.DiskType == nil {
		if fieldsToDefault["settings.diskType"] {
			out.Settings.DataDiskType = actual.Settings.DataDiskType
		}
	}
	if in.Spec.Settings.DeletionProtectionEnabled == nil {
		if fieldsToDefault["settings.deletionProtectionEnabled"] {
			out.Settings.DeletionProtectionEnabled = actual.Settings.DeletionProtectionEnabled
		}
	}
	if in.Spec.Settings.Edition == nil {
		if fieldsToDefault["settings.edition"] {
			out.Settings.Edition = actual.Settings.Edition
		}
	}
	if in.Spec.Settings.IpConfiguration != nil {
		if in.Spec.Settings.IpConfiguration.Ipv4Enabled == nil {
			if fieldsToDefault["settings.ipConfiguration.ipv4Enabled"] {
				out.Settings.IpConfiguration.Ipv4Enabled = actual.Settings.IpConfiguration.Ipv4Enabled
			}
		}
		if in.Spec.Settings.IpConfiguration.SslMode == nil {
			if fieldsToDefault["settings.ipConfiguration.sslMode"] {
				out.Settings.IpConfiguration.SslMode = actual.Settings.IpConfiguration.SslMode
			}
		}
	}
	if in.Spec.Settings.LocationPreference == nil {
		if fieldsToDefault["settings.locationPreference"] && actual != nil {
			out.Settings.LocationPreference = actual.Settings.LocationPreference
		}
	}
	if in.Spec.Settings.PricingPlan == nil {
		if fieldsToDefault["settings.pricingPlan"] {
			out.Settings.PricingPlan = actual.Settings.PricingPlan
		}
	}
	if in.Spec.Settings.ReplicationType == nil {
		if fieldsToDefault["settings.replicationType"] {
			out.Settings.ReplicationType = actual.Settings.ReplicationType
		}
	}
	if in.Spec.Settings.Tier == "" {
		if fieldsToDefault["settings.tier"] {
			out.Settings.Tier = actual.Settings.Tier
		}
	}
	if in.Spec.Settings.TimeZone == nil {
		if fieldsToDefault["settings.timeZone"] {
			out.Settings.TimeZone = actual.Settings.TimeZone
		}
	}

	if actual != nil {
		out.Settings.SettingsVersion = actual.Settings.SettingsVersion
	}
}
