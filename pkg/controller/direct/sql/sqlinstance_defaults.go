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
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/sqladmin/v1beta4"
)

const (
	// DefaultQueryPlansPerMinute is the default value for QueryPlansPerMinute.
	// https://docs.cloud.google.com/sql/docs/mysql/admin-api/rest/v1/instances#insightsconfig
	DefaultQueryPlansPerMinute = 5
	defaultRetainedBackups     = 7
)

func ApplySQLInstanceGCPDefaults(in *krm.SQLInstance, out *api.DatabaseInstance, actual *api.DatabaseInstance, fieldMetadata map[string]*FieldMetadata) error {
	// Stage 1: Apply all client-side defaults as if no fields are unmanaged.
	if in.Spec.InstanceType == nil {
		// GCP default InstanceType is CLOUD_SQL_INSTANCE.
		out.InstanceType = "CLOUD_SQL_INSTANCE"
	}
	if in.Spec.MaintenanceVersion == nil && actual != nil {
		// If desired maintenanceVersion is not specified, assume user wants the actual.
		out.MaintenanceVersion = actual.MaintenanceVersion
	}
	if in.Spec.Settings.ActivationPolicy == nil {
		// GCP default ActivationPolicy is ALWAYS.
		out.Settings.ActivationPolicy = "ALWAYS"
	}
	if in.Spec.Settings.AuthorizedGaeApplications == nil {
		// For some reason, GCP API uses empty slice instead of nil.
		out.Settings.AuthorizedGaeApplications = make([]string, 0)
	}
	if in.Spec.Settings.AvailabilityType == nil {
		// GCP default AvailailbilityType is ZONAL.
		out.Settings.AvailabilityType = "ZONAL"
	}
	if in.Spec.Settings.BackupConfiguration == nil && actual != nil && !actual.Settings.BackupConfiguration.Enabled {
		// If desired backupConfiguration is not specified and actual is disabled, use the actual.
		out.Settings.BackupConfiguration = actual.Settings.BackupConfiguration
	}
	if in.Spec.Settings.BackupConfiguration != nil {
		if in.Spec.Settings.BackupConfiguration.BackupRetentionSettings != nil {
			if in.Spec.Settings.BackupConfiguration.BackupRetentionSettings.RetentionUnit == nil {
				// GCP default retentionUnit is COUNT.
				out.Settings.BackupConfiguration.BackupRetentionSettings.RetentionUnit = "COUNT"
			}
			if in.Spec.Settings.BackupConfiguration.BackupRetentionSettings.RetainedBackups == 0 {
				if out.Settings.Edition == "ENTERPRISE_PLUS" {
					out.Settings.BackupConfiguration.BackupRetentionSettings.RetainedBackups = 15
				} else {
					out.Settings.BackupConfiguration.BackupRetentionSettings.RetainedBackups = 7
				}
			}
		}
	}
	if in.Spec.Settings.BackupConfiguration != nil &&
		(in.Spec.Settings.BackupConfiguration.Enabled == nil || *in.Spec.Settings.BackupConfiguration.Enabled) &&
		in.Spec.Settings.RetainedBackups == nil {
		// TODO(b/293437281): Change this to a server-side default and remove this code
		// Why a client-side default? The server-side default is 7, but the API will return "0" if we send "nil".
		// We want to avoid causing a diff for users that don't have this field specified.
		retainedBackups := int64(defaultRetainedBackups)
		in.Spec.Settings.RetainedBackups = &retainedBackups
	}
	if in.Spec.Settings.ConnectorEnforcement == nil {
		// GCP default ConnectorEnforcement is NOT_REQUIRED.
		out.Settings.ConnectorEnforcement = "NOT_REQUIRED"
	}
	if in.Spec.Settings.DiskType == nil {
		// GCP default DiskType is PD_SSD.
		out.Settings.DataDiskType = "PD_SSD"
	}
	if in.Spec.Settings.Edition == nil {
		// Apply client side GCP default Edition is ENTERPRISE.
		out.Settings.Edition = "ENTERPRISE"
	} else {
		// Validate that the edition is one of the supported values.
		if out.Settings.Edition != "ENTERPRISE" && out.Settings.Edition != "ENTERPRISE_PLUS" {
			return fmt.Errorf("invalid edition %q: must be one of \"ENTERPRISE\", \"ENTERPRISE_PLUS\"", out.Settings.Edition)
		}
	}
	if in.Spec.Settings.InsightsConfig != nil {
		if in.Spec.Settings.InsightsConfig.QueryPlansPerMinute == nil {
			out.Settings.InsightsConfig.QueryPlansPerMinute = DefaultQueryPlansPerMinute
		}
	}

	if in.Spec.Settings.IpConfiguration == nil {
		// GCP default IpConfiguration.
		out.Settings.IpConfiguration = &api.IpConfiguration{
			Ipv4Enabled: true,
			SslMode:     "ALLOW_UNENCRYPTED_AND_ENCRYPTED",
		}
	}
	if in.Spec.Settings.IpConfiguration != nil {
		if in.Spec.Settings.IpConfiguration.Ipv4Enabled == nil {
			// GCP default IpConfiguration.Ipv4Enabled is true.
			out.Settings.IpConfiguration.Ipv4Enabled = true
		}
		if in.Spec.Settings.IpConfiguration.SslMode == nil {
			if out.Settings.IpConfiguration.RequireSsl {
				if strings.HasPrefix(out.DatabaseVersion, "MYSQL") || strings.HasPrefix(out.DatabaseVersion, "POSTGRES") {
					// If RequireSsl is true, and db version is MySQL or Postgres,
					// GCP default SslMode is TRUSTED_CLIENT_CERTIFICATE_REQUIRED.
					out.Settings.IpConfiguration.SslMode = "TRUSTED_CLIENT_CERTIFICATE_REQUIRED"
				} else {
					// Otherwise, if RequireSsl is true and db version is SQLSERVER,
					// GCP default SslMode is ENCRYPTED_ONLY.
					out.Settings.IpConfiguration.SslMode = "ENCRYPTED_ONLY"
				}
			} else {
				// If RequireSsl is false, GCP default IpConfiguration.SslMode is ALLOW_UNENCRYPTED_AND_ENCRYPTED.
				out.Settings.IpConfiguration.SslMode = "ALLOW_UNENCRYPTED_AND_ENCRYPTED"
			}
		}
	}
	if in.Spec.Settings.LocationPreference == nil && actual != nil {
		// Use GCP specified locationPreference.
		out.Settings.LocationPreference = actual.Settings.LocationPreference
	}
	if in.Spec.Settings.PricingPlan == nil {
		// GCP default PricingPlan is PER_USE.
		out.Settings.PricingPlan = "PER_USE"
	}
	if in.Spec.Settings.ReplicationType == nil {
		// GCP default ReplicationType is SYNCHRONOUS.
		out.Settings.ReplicationType = "SYNCHRONOUS"
	}
	if in.Spec.Settings.DiskAutoresize == nil {
		// GCP default StorageAutoResize is true.
		out.Settings.StorageAutoResize = direct.PtrTo(true)
	}
	if in.Spec.Settings.DiskSize == nil && actual != nil && *out.Settings.StorageAutoResize {
		// If desired DiskSize is not specified and StorageAutoResize is enabled, use the actual disk size.
		// Note: This must be set AFTER setting the default value for StorageAutoResize.
		out.Settings.DataDiskSizeGb = actual.Settings.DataDiskSizeGb
	}
	if actual != nil {
		// GCP API requires we set the current settings version, otherwise update will fail.
		out.Settings.SettingsVersion = actual.Settings.SettingsVersion
	}

	// Stage 2: Preserve any fields that are unmanaged by the user.
	// This generic loop will overwrite any defaults set in Stage 1.
	if actual == nil {
		return nil
	}
	for _, field := range fieldMetadata {
		if field.isUnmanaged {
			field.preserveActualValue(out, actual)
		}
	}

	return nil
}
