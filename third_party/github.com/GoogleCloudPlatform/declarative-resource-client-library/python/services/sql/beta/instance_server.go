// Copyright 2021 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import (
	"context"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/sql/beta/sql_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/sql/beta"
)

// Server implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstanceBackendTypeEnum converts a InstanceBackendTypeEnum enum from its proto representation.
func ProtoToSqlBetaInstanceBackendTypeEnum(e betapb.SqlBetaInstanceBackendTypeEnum) *beta.InstanceBackendTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.SqlBetaInstanceBackendTypeEnum_name[int32(e)]; ok {
		e := beta.InstanceBackendTypeEnum(n[len("SqlBetaInstanceBackendTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceDatabaseVersionEnum converts a InstanceDatabaseVersionEnum enum from its proto representation.
func ProtoToSqlBetaInstanceDatabaseVersionEnum(e betapb.SqlBetaInstanceDatabaseVersionEnum) *beta.InstanceDatabaseVersionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.SqlBetaInstanceDatabaseVersionEnum_name[int32(e)]; ok {
		e := beta.InstanceDatabaseVersionEnum(n[len("SqlBetaInstanceDatabaseVersionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceInstanceTypeEnum converts a InstanceInstanceTypeEnum enum from its proto representation.
func ProtoToSqlBetaInstanceInstanceTypeEnum(e betapb.SqlBetaInstanceInstanceTypeEnum) *beta.InstanceInstanceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.SqlBetaInstanceInstanceTypeEnum_name[int32(e)]; ok {
		e := beta.InstanceInstanceTypeEnum(n[len("SqlBetaInstanceInstanceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceIPAddressesTypeEnum converts a InstanceIPAddressesTypeEnum enum from its proto representation.
func ProtoToSqlBetaInstanceIPAddressesTypeEnum(e betapb.SqlBetaInstanceIPAddressesTypeEnum) *beta.InstanceIPAddressesTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.SqlBetaInstanceIPAddressesTypeEnum_name[int32(e)]; ok {
		e := beta.InstanceIPAddressesTypeEnum(n[len("SqlBetaInstanceIPAddressesTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceSettingsAvailabilityTypeEnum converts a InstanceSettingsAvailabilityTypeEnum enum from its proto representation.
func ProtoToSqlBetaInstanceSettingsAvailabilityTypeEnum(e betapb.SqlBetaInstanceSettingsAvailabilityTypeEnum) *beta.InstanceSettingsAvailabilityTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.SqlBetaInstanceSettingsAvailabilityTypeEnum_name[int32(e)]; ok {
		e := beta.InstanceSettingsAvailabilityTypeEnum(n[len("SqlBetaInstanceSettingsAvailabilityTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceSettingsPricingPlanEnum converts a InstanceSettingsPricingPlanEnum enum from its proto representation.
func ProtoToSqlBetaInstanceSettingsPricingPlanEnum(e betapb.SqlBetaInstanceSettingsPricingPlanEnum) *beta.InstanceSettingsPricingPlanEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.SqlBetaInstanceSettingsPricingPlanEnum_name[int32(e)]; ok {
		e := beta.InstanceSettingsPricingPlanEnum(n[len("SqlBetaInstanceSettingsPricingPlanEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceSettingsReplicationTypeEnum converts a InstanceSettingsReplicationTypeEnum enum from its proto representation.
func ProtoToSqlBetaInstanceSettingsReplicationTypeEnum(e betapb.SqlBetaInstanceSettingsReplicationTypeEnum) *beta.InstanceSettingsReplicationTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.SqlBetaInstanceSettingsReplicationTypeEnum_name[int32(e)]; ok {
		e := beta.InstanceSettingsReplicationTypeEnum(n[len("SqlBetaInstanceSettingsReplicationTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceSettingsActivationPolicyEnum converts a InstanceSettingsActivationPolicyEnum enum from its proto representation.
func ProtoToSqlBetaInstanceSettingsActivationPolicyEnum(e betapb.SqlBetaInstanceSettingsActivationPolicyEnum) *beta.InstanceSettingsActivationPolicyEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.SqlBetaInstanceSettingsActivationPolicyEnum_name[int32(e)]; ok {
		e := beta.InstanceSettingsActivationPolicyEnum(n[len("SqlBetaInstanceSettingsActivationPolicyEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceSettingsDataDiskTypeEnum converts a InstanceSettingsDataDiskTypeEnum enum from its proto representation.
func ProtoToSqlBetaInstanceSettingsDataDiskTypeEnum(e betapb.SqlBetaInstanceSettingsDataDiskTypeEnum) *beta.InstanceSettingsDataDiskTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.SqlBetaInstanceSettingsDataDiskTypeEnum_name[int32(e)]; ok {
		e := beta.InstanceSettingsDataDiskTypeEnum(n[len("SqlBetaInstanceSettingsDataDiskTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceSettingsMaintenanceWindowUpdateTrackEnum converts a InstanceSettingsMaintenanceWindowUpdateTrackEnum enum from its proto representation.
func ProtoToSqlBetaInstanceSettingsMaintenanceWindowUpdateTrackEnum(e betapb.SqlBetaInstanceSettingsMaintenanceWindowUpdateTrackEnum) *beta.InstanceSettingsMaintenanceWindowUpdateTrackEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.SqlBetaInstanceSettingsMaintenanceWindowUpdateTrackEnum_name[int32(e)]; ok {
		e := beta.InstanceSettingsMaintenanceWindowUpdateTrackEnum(n[len("SqlBetaInstanceSettingsMaintenanceWindowUpdateTrackEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum converts a InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum enum from its proto representation.
func ProtoToSqlBetaInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum(e betapb.SqlBetaInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum) *beta.InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.SqlBetaInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum_name[int32(e)]; ok {
		e := beta.InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum(n[len("SqlBetaInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceMaxDiskSize converts a InstanceMaxDiskSize resource from its proto representation.
func ProtoToSqlBetaInstanceMaxDiskSize(p *betapb.SqlBetaInstanceMaxDiskSize) *beta.InstanceMaxDiskSize {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceMaxDiskSize{
		Value: dcl.Int64OrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceCurrentDiskSize converts a InstanceCurrentDiskSize resource from its proto representation.
func ProtoToSqlBetaInstanceCurrentDiskSize(p *betapb.SqlBetaInstanceCurrentDiskSize) *beta.InstanceCurrentDiskSize {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceCurrentDiskSize{
		Value: dcl.Int64OrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceDiskEncryptionConfiguration converts a InstanceDiskEncryptionConfiguration resource from its proto representation.
func ProtoToSqlBetaInstanceDiskEncryptionConfiguration(p *betapb.SqlBetaInstanceDiskEncryptionConfiguration) *beta.InstanceDiskEncryptionConfiguration {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceDiskEncryptionConfiguration{
		KmsKeyName: dcl.StringOrNil(p.KmsKeyName),
		Kind:       dcl.StringOrNil(p.Kind),
	}
	return obj
}

// ProtoToInstanceFailoverReplica converts a InstanceFailoverReplica resource from its proto representation.
func ProtoToSqlBetaInstanceFailoverReplica(p *betapb.SqlBetaInstanceFailoverReplica) *beta.InstanceFailoverReplica {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceFailoverReplica{
		Name:             dcl.StringOrNil(p.Name),
		Available:        dcl.Bool(p.Available),
		FailoverInstance: ProtoToSqlBetaInstanceFailoverReplicaFailoverInstance(p.GetFailoverInstance()),
	}
	return obj
}

// ProtoToInstanceFailoverReplicaFailoverInstance converts a InstanceFailoverReplicaFailoverInstance resource from its proto representation.
func ProtoToSqlBetaInstanceFailoverReplicaFailoverInstance(p *betapb.SqlBetaInstanceFailoverReplicaFailoverInstance) *beta.InstanceFailoverReplicaFailoverInstance {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceFailoverReplicaFailoverInstance{
		Name:   dcl.StringOrNil(p.Name),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceIPAddresses converts a InstanceIPAddresses resource from its proto representation.
func ProtoToSqlBetaInstanceIPAddresses(p *betapb.SqlBetaInstanceIPAddresses) *beta.InstanceIPAddresses {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceIPAddresses{
		Type:         ProtoToSqlBetaInstanceIPAddressesTypeEnum(p.GetType()),
		IPAddress:    dcl.StringOrNil(p.IpAddress),
		TimeToRetire: ProtoToSqlBetaInstanceIPAddressesTimeToRetire(p.GetTimeToRetire()),
	}
	return obj
}

// ProtoToInstanceIPAddressesTimeToRetire converts a InstanceIPAddressesTimeToRetire resource from its proto representation.
func ProtoToSqlBetaInstanceIPAddressesTimeToRetire(p *betapb.SqlBetaInstanceIPAddressesTimeToRetire) *beta.InstanceIPAddressesTimeToRetire {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceIPAddressesTimeToRetire{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToInstanceMasterInstance converts a InstanceMasterInstance resource from its proto representation.
func ProtoToSqlBetaInstanceMasterInstance(p *betapb.SqlBetaInstanceMasterInstance) *beta.InstanceMasterInstance {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceMasterInstance{
		Name:   dcl.StringOrNil(p.Name),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceReplicaConfiguration converts a InstanceReplicaConfiguration resource from its proto representation.
func ProtoToSqlBetaInstanceReplicaConfiguration(p *betapb.SqlBetaInstanceReplicaConfiguration) *beta.InstanceReplicaConfiguration {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceReplicaConfiguration{
		Kind:                      dcl.StringOrNil(p.Kind),
		MysqlReplicaConfiguration: ProtoToSqlBetaInstanceReplicaConfigurationMysqlReplicaConfiguration(p.GetMysqlReplicaConfiguration()),
		FailoverTarget:            dcl.Bool(p.FailoverTarget),
		ReplicaPoolConfiguration:  ProtoToSqlBetaInstanceReplicaConfigurationReplicaPoolConfiguration(p.GetReplicaPoolConfiguration()),
	}
	return obj
}

// ProtoToInstanceReplicaConfigurationMysqlReplicaConfiguration converts a InstanceReplicaConfigurationMysqlReplicaConfiguration resource from its proto representation.
func ProtoToSqlBetaInstanceReplicaConfigurationMysqlReplicaConfiguration(p *betapb.SqlBetaInstanceReplicaConfigurationMysqlReplicaConfiguration) *beta.InstanceReplicaConfigurationMysqlReplicaConfiguration {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceReplicaConfigurationMysqlReplicaConfiguration{
		DumpFilePath:            dcl.StringOrNil(p.DumpFilePath),
		Username:                dcl.StringOrNil(p.Username),
		Password:                dcl.StringOrNil(p.Password),
		ConnectRetryInterval:    dcl.Int64OrNil(p.ConnectRetryInterval),
		MasterHeartbeatPeriod:   ProtoToSqlBetaInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(p.GetMasterHeartbeatPeriod()),
		CaCertificate:           dcl.StringOrNil(p.CaCertificate),
		ClientCertificate:       dcl.StringOrNil(p.ClientCertificate),
		ClientKey:               dcl.StringOrNil(p.ClientKey),
		SslCipher:               dcl.StringOrNil(p.SslCipher),
		VerifyServerCertificate: dcl.Bool(p.VerifyServerCertificate),
		Kind:                    dcl.StringOrNil(p.Kind),
	}
	return obj
}

// ProtoToInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod converts a InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod resource from its proto representation.
func ProtoToSqlBetaInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(p *betapb.SqlBetaInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod) *beta.InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod{
		Value: dcl.Int64OrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceReplicaConfigurationReplicaPoolConfiguration converts a InstanceReplicaConfigurationReplicaPoolConfiguration resource from its proto representation.
func ProtoToSqlBetaInstanceReplicaConfigurationReplicaPoolConfiguration(p *betapb.SqlBetaInstanceReplicaConfigurationReplicaPoolConfiguration) *beta.InstanceReplicaConfigurationReplicaPoolConfiguration {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceReplicaConfigurationReplicaPoolConfiguration{
		Kind:                         dcl.StringOrNil(p.Kind),
		StaticPoolConfiguration:      ProtoToSqlBetaInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(p.GetStaticPoolConfiguration()),
		AutoscalingPoolConfiguration: ProtoToSqlBetaInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(p.GetAutoscalingPoolConfiguration()),
		ReplicaCount:                 dcl.Int64OrNil(p.ReplicaCount),
		ExposeReplicaIP:              dcl.Bool(p.ExposeReplicaIp),
	}
	return obj
}

// ProtoToInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration converts a InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration resource from its proto representation.
func ProtoToSqlBetaInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(p *betapb.SqlBetaInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration) *beta.InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration{
		Kind:            dcl.StringOrNil(p.Kind),
		ReplicaCount:    dcl.Int64OrNil(p.ReplicaCount),
		ExposeReplicaIP: dcl.Bool(p.ExposeReplicaIp),
	}
	return obj
}

// ProtoToInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration converts a InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration resource from its proto representation.
func ProtoToSqlBetaInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(p *betapb.SqlBetaInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration) *beta.InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration{
		Kind:            dcl.StringOrNil(p.Kind),
		MinReplicaCount: dcl.Int64OrNil(p.MinReplicaCount),
		MaxReplicaCount: dcl.Int64OrNil(p.MaxReplicaCount),
		TargetCpuUtil:   dcl.Float64OrNil(p.TargetCpuUtil),
	}
	return obj
}

// ProtoToInstanceScheduledMaintenance converts a InstanceScheduledMaintenance resource from its proto representation.
func ProtoToSqlBetaInstanceScheduledMaintenance(p *betapb.SqlBetaInstanceScheduledMaintenance) *beta.InstanceScheduledMaintenance {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceScheduledMaintenance{
		StartTime:     ProtoToSqlBetaInstanceScheduledMaintenanceStartTime(p.GetStartTime()),
		CanDefer:      dcl.Bool(p.CanDefer),
		CanReschedule: dcl.Bool(p.CanReschedule),
	}
	return obj
}

// ProtoToInstanceScheduledMaintenanceStartTime converts a InstanceScheduledMaintenanceStartTime resource from its proto representation.
func ProtoToSqlBetaInstanceScheduledMaintenanceStartTime(p *betapb.SqlBetaInstanceScheduledMaintenanceStartTime) *beta.InstanceScheduledMaintenanceStartTime {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceScheduledMaintenanceStartTime{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToInstanceSettings converts a InstanceSettings resource from its proto representation.
func ProtoToSqlBetaInstanceSettings(p *betapb.SqlBetaInstanceSettings) *beta.InstanceSettings {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceSettings{
		Tier:                        dcl.StringOrNil(p.Tier),
		Kind:                        dcl.StringOrNil(p.Kind),
		AvailabilityType:            ProtoToSqlBetaInstanceSettingsAvailabilityTypeEnum(p.GetAvailabilityType()),
		PricingPlan:                 ProtoToSqlBetaInstanceSettingsPricingPlanEnum(p.GetPricingPlan()),
		ReplicationType:             ProtoToSqlBetaInstanceSettingsReplicationTypeEnum(p.GetReplicationType()),
		ActivationPolicy:            ProtoToSqlBetaInstanceSettingsActivationPolicyEnum(p.GetActivationPolicy()),
		StorageAutoResize:           dcl.Bool(p.StorageAutoResize),
		DataDiskType:                ProtoToSqlBetaInstanceSettingsDataDiskTypeEnum(p.GetDataDiskType()),
		DatabaseReplicationEnabled:  dcl.Bool(p.DatabaseReplicationEnabled),
		CrashSafeReplicationEnabled: dcl.Bool(p.CrashSafeReplicationEnabled),
		SettingsVersion:             ProtoToSqlBetaInstanceSettingsSettingsVersion(p.GetSettingsVersion()),
		StorageAutoResizeLimit:      ProtoToSqlBetaInstanceSettingsStorageAutoResizeLimit(p.GetStorageAutoResizeLimit()),
		IPConfiguration:             ProtoToSqlBetaInstanceSettingsIPConfiguration(p.GetIpConfiguration()),
		LocationPreference:          ProtoToSqlBetaInstanceSettingsLocationPreference(p.GetLocationPreference()),
		MaintenanceWindow:           ProtoToSqlBetaInstanceSettingsMaintenanceWindow(p.GetMaintenanceWindow()),
		BackupConfiguration:         ProtoToSqlBetaInstanceSettingsBackupConfiguration(p.GetBackupConfiguration()),
		DataDiskSizeGb:              ProtoToSqlBetaInstanceSettingsDataDiskSizeGb(p.GetDataDiskSizeGb()),
		ActiveDirectoryConfig:       ProtoToSqlBetaInstanceSettingsActiveDirectoryConfig(p.GetActiveDirectoryConfig()),
		Collation:                   dcl.StringOrNil(p.Collation),
		InsightsConfig:              ProtoToSqlBetaInstanceSettingsInsightsConfig(p.GetInsightsConfig()),
	}
	for _, r := range p.GetAuthorizedGaeApplications() {
		obj.AuthorizedGaeApplications = append(obj.AuthorizedGaeApplications, r)
	}
	for _, r := range p.GetDatabaseFlags() {
		obj.DatabaseFlags = append(obj.DatabaseFlags, *ProtoToSqlBetaInstanceSettingsDatabaseFlags(r))
	}
	for _, r := range p.GetDenyMaintenancePeriods() {
		obj.DenyMaintenancePeriods = append(obj.DenyMaintenancePeriods, *ProtoToSqlBetaInstanceSettingsDenyMaintenancePeriods(r))
	}
	return obj
}

// ProtoToInstanceSettingsSettingsVersion converts a InstanceSettingsSettingsVersion resource from its proto representation.
func ProtoToSqlBetaInstanceSettingsSettingsVersion(p *betapb.SqlBetaInstanceSettingsSettingsVersion) *beta.InstanceSettingsSettingsVersion {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceSettingsSettingsVersion{
		Value: dcl.Int64OrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceSettingsStorageAutoResizeLimit converts a InstanceSettingsStorageAutoResizeLimit resource from its proto representation.
func ProtoToSqlBetaInstanceSettingsStorageAutoResizeLimit(p *betapb.SqlBetaInstanceSettingsStorageAutoResizeLimit) *beta.InstanceSettingsStorageAutoResizeLimit {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceSettingsStorageAutoResizeLimit{
		Value: dcl.Int64OrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceSettingsIPConfiguration converts a InstanceSettingsIPConfiguration resource from its proto representation.
func ProtoToSqlBetaInstanceSettingsIPConfiguration(p *betapb.SqlBetaInstanceSettingsIPConfiguration) *beta.InstanceSettingsIPConfiguration {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceSettingsIPConfiguration{
		IPv4Enabled:    dcl.Bool(p.Ipv4Enabled),
		PrivateNetwork: dcl.StringOrNil(p.PrivateNetwork),
		RequireSsl:     dcl.Bool(p.RequireSsl),
	}
	for _, r := range p.GetAuthorizedNetworks() {
		obj.AuthorizedNetworks = append(obj.AuthorizedNetworks, *ProtoToSqlBetaInstanceSettingsIPConfigurationAuthorizedNetworks(r))
	}
	return obj
}

// ProtoToInstanceSettingsIPConfigurationAuthorizedNetworks converts a InstanceSettingsIPConfigurationAuthorizedNetworks resource from its proto representation.
func ProtoToSqlBetaInstanceSettingsIPConfigurationAuthorizedNetworks(p *betapb.SqlBetaInstanceSettingsIPConfigurationAuthorizedNetworks) *beta.InstanceSettingsIPConfigurationAuthorizedNetworks {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceSettingsIPConfigurationAuthorizedNetworks{
		Value:          dcl.StringOrNil(p.Value),
		ExpirationTime: dcl.StringOrNil(p.GetExpirationTime()),
		Name:           dcl.StringOrNil(p.Name),
		Kind:           dcl.StringOrNil(p.Kind),
	}
	return obj
}

// ProtoToInstanceSettingsLocationPreference converts a InstanceSettingsLocationPreference resource from its proto representation.
func ProtoToSqlBetaInstanceSettingsLocationPreference(p *betapb.SqlBetaInstanceSettingsLocationPreference) *beta.InstanceSettingsLocationPreference {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceSettingsLocationPreference{
		Zone: dcl.StringOrNil(p.Zone),
		Kind: dcl.StringOrNil(p.Kind),
	}
	return obj
}

// ProtoToInstanceSettingsDatabaseFlags converts a InstanceSettingsDatabaseFlags resource from its proto representation.
func ProtoToSqlBetaInstanceSettingsDatabaseFlags(p *betapb.SqlBetaInstanceSettingsDatabaseFlags) *beta.InstanceSettingsDatabaseFlags {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceSettingsDatabaseFlags{
		Name:  dcl.StringOrNil(p.Name),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceSettingsMaintenanceWindow converts a InstanceSettingsMaintenanceWindow resource from its proto representation.
func ProtoToSqlBetaInstanceSettingsMaintenanceWindow(p *betapb.SqlBetaInstanceSettingsMaintenanceWindow) *beta.InstanceSettingsMaintenanceWindow {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceSettingsMaintenanceWindow{
		Hour:        dcl.Int64OrNil(p.Hour),
		Day:         dcl.Int64OrNil(p.Day),
		UpdateTrack: ProtoToSqlBetaInstanceSettingsMaintenanceWindowUpdateTrackEnum(p.GetUpdateTrack()),
		Kind:        dcl.StringOrNil(p.Kind),
	}
	return obj
}

// ProtoToInstanceSettingsBackupConfiguration converts a InstanceSettingsBackupConfiguration resource from its proto representation.
func ProtoToSqlBetaInstanceSettingsBackupConfiguration(p *betapb.SqlBetaInstanceSettingsBackupConfiguration) *beta.InstanceSettingsBackupConfiguration {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceSettingsBackupConfiguration{
		StartTime:                   dcl.StringOrNil(p.StartTime),
		Enabled:                     dcl.Bool(p.Enabled),
		Kind:                        dcl.StringOrNil(p.Kind),
		BinaryLogEnabled:            dcl.Bool(p.BinaryLogEnabled),
		Location:                    dcl.StringOrNil(p.Location),
		BackupRetentionSettings:     ProtoToSqlBetaInstanceSettingsBackupConfigurationBackupRetentionSettings(p.GetBackupRetentionSettings()),
		TransactionLogRetentionDays: dcl.Int64OrNil(p.TransactionLogRetentionDays),
	}
	return obj
}

// ProtoToInstanceSettingsBackupConfigurationBackupRetentionSettings converts a InstanceSettingsBackupConfigurationBackupRetentionSettings resource from its proto representation.
func ProtoToSqlBetaInstanceSettingsBackupConfigurationBackupRetentionSettings(p *betapb.SqlBetaInstanceSettingsBackupConfigurationBackupRetentionSettings) *beta.InstanceSettingsBackupConfigurationBackupRetentionSettings {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceSettingsBackupConfigurationBackupRetentionSettings{
		RetentionUnit:   ProtoToSqlBetaInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum(p.GetRetentionUnit()),
		RetainedBackups: dcl.Int64OrNil(p.RetainedBackups),
	}
	return obj
}

// ProtoToInstanceSettingsDataDiskSizeGb converts a InstanceSettingsDataDiskSizeGb resource from its proto representation.
func ProtoToSqlBetaInstanceSettingsDataDiskSizeGb(p *betapb.SqlBetaInstanceSettingsDataDiskSizeGb) *beta.InstanceSettingsDataDiskSizeGb {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceSettingsDataDiskSizeGb{
		Value: dcl.Int64OrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceSettingsActiveDirectoryConfig converts a InstanceSettingsActiveDirectoryConfig resource from its proto representation.
func ProtoToSqlBetaInstanceSettingsActiveDirectoryConfig(p *betapb.SqlBetaInstanceSettingsActiveDirectoryConfig) *beta.InstanceSettingsActiveDirectoryConfig {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceSettingsActiveDirectoryConfig{
		Kind:   dcl.StringOrNil(p.Kind),
		Domain: dcl.StringOrNil(p.Domain),
	}
	return obj
}

// ProtoToInstanceSettingsDenyMaintenancePeriods converts a InstanceSettingsDenyMaintenancePeriods resource from its proto representation.
func ProtoToSqlBetaInstanceSettingsDenyMaintenancePeriods(p *betapb.SqlBetaInstanceSettingsDenyMaintenancePeriods) *beta.InstanceSettingsDenyMaintenancePeriods {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceSettingsDenyMaintenancePeriods{
		StartDate: dcl.StringOrNil(p.StartDate),
		EndDate:   dcl.StringOrNil(p.EndDate),
		Time:      dcl.StringOrNil(p.Time),
	}
	return obj
}

// ProtoToInstanceSettingsInsightsConfig converts a InstanceSettingsInsightsConfig resource from its proto representation.
func ProtoToSqlBetaInstanceSettingsInsightsConfig(p *betapb.SqlBetaInstanceSettingsInsightsConfig) *beta.InstanceSettingsInsightsConfig {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceSettingsInsightsConfig{
		QueryInsightsEnabled:  dcl.Bool(p.QueryInsightsEnabled),
		RecordClientAddress:   dcl.Bool(p.RecordClientAddress),
		RecordApplicationTags: dcl.Bool(p.RecordApplicationTags),
		QueryStringLength:     dcl.Int64OrNil(p.QueryStringLength),
	}
	return obj
}

// ProtoToInstanceReplicaInstances converts a InstanceReplicaInstances resource from its proto representation.
func ProtoToSqlBetaInstanceReplicaInstances(p *betapb.SqlBetaInstanceReplicaInstances) *beta.InstanceReplicaInstances {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceReplicaInstances{
		Name:   dcl.StringOrNil(p.Name),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceServerCaCert converts a InstanceServerCaCert resource from its proto representation.
func ProtoToSqlBetaInstanceServerCaCert(p *betapb.SqlBetaInstanceServerCaCert) *beta.InstanceServerCaCert {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceServerCaCert{
		Kind:             dcl.StringOrNil(p.Kind),
		CertSerialNumber: dcl.StringOrNil(p.CertSerialNumber),
		Cert:             dcl.StringOrNil(p.Cert),
		CreateTime:       dcl.StringOrNil(p.GetCreateTime()),
		CommonName:       dcl.StringOrNil(p.CommonName),
		ExpirationTime:   dcl.StringOrNil(p.GetExpirationTime()),
		Sha1Fingerprint:  dcl.StringOrNil(p.Sha1Fingerprint),
		Instance:         dcl.StringOrNil(p.Instance),
	}
	return obj
}

// ProtoToInstanceOnPremisesConfiguration converts a InstanceOnPremisesConfiguration resource from its proto representation.
func ProtoToSqlBetaInstanceOnPremisesConfiguration(p *betapb.SqlBetaInstanceOnPremisesConfiguration) *beta.InstanceOnPremisesConfiguration {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceOnPremisesConfiguration{
		HostPort:          dcl.StringOrNil(p.HostPort),
		Kind:              dcl.StringOrNil(p.Kind),
		Username:          dcl.StringOrNil(p.Username),
		Password:          dcl.StringOrNil(p.Password),
		CaCertificate:     dcl.StringOrNil(p.CaCertificate),
		ClientCertificate: dcl.StringOrNil(p.ClientCertificate),
		ClientKey:         dcl.StringOrNil(p.ClientKey),
		DumpFilePath:      dcl.StringOrNil(p.DumpFilePath),
		Database:          dcl.StringOrNil(p.Database),
	}
	for _, r := range p.GetReplicatedDatabases() {
		obj.ReplicatedDatabases = append(obj.ReplicatedDatabases, r)
	}
	return obj
}

// ProtoToInstanceDiskEncryptionStatus converts a InstanceDiskEncryptionStatus resource from its proto representation.
func ProtoToSqlBetaInstanceDiskEncryptionStatus(p *betapb.SqlBetaInstanceDiskEncryptionStatus) *beta.InstanceDiskEncryptionStatus {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceDiskEncryptionStatus{
		KmsKeyVersionName: dcl.StringOrNil(p.KmsKeyVersionName),
		Kind:              dcl.StringOrNil(p.Kind),
	}
	return obj
}

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *betapb.SqlBetaInstance) *beta.Instance {
	obj := &beta.Instance{
		BackendType:                 ProtoToSqlBetaInstanceBackendTypeEnum(p.GetBackendType()),
		ConnectionName:              dcl.StringOrNil(p.ConnectionName),
		DatabaseVersion:             ProtoToSqlBetaInstanceDatabaseVersionEnum(p.GetDatabaseVersion()),
		Etag:                        dcl.StringOrNil(p.Etag),
		GceZone:                     dcl.StringOrNil(p.GceZone),
		InstanceType:                ProtoToSqlBetaInstanceInstanceTypeEnum(p.GetInstanceType()),
		MasterInstanceName:          dcl.StringOrNil(p.MasterInstanceName),
		MaxDiskSize:                 ProtoToSqlBetaInstanceMaxDiskSize(p.GetMaxDiskSize()),
		Name:                        dcl.StringOrNil(p.Name),
		Project:                     dcl.StringOrNil(p.Project),
		Region:                      dcl.StringOrNil(p.Region),
		RootPassword:                dcl.StringOrNil(p.RootPassword),
		CurrentDiskSize:             ProtoToSqlBetaInstanceCurrentDiskSize(p.GetCurrentDiskSize()),
		DiskEncryptionConfiguration: ProtoToSqlBetaInstanceDiskEncryptionConfiguration(p.GetDiskEncryptionConfiguration()),
		FailoverReplica:             ProtoToSqlBetaInstanceFailoverReplica(p.GetFailoverReplica()),
		MasterInstance:              ProtoToSqlBetaInstanceMasterInstance(p.GetMasterInstance()),
		ReplicaConfiguration:        ProtoToSqlBetaInstanceReplicaConfiguration(p.GetReplicaConfiguration()),
		ScheduledMaintenance:        ProtoToSqlBetaInstanceScheduledMaintenance(p.GetScheduledMaintenance()),
		Settings:                    ProtoToSqlBetaInstanceSettings(p.GetSettings()),
		State:                       dcl.StringOrNil(p.State),
		ServerCaCert:                ProtoToSqlBetaInstanceServerCaCert(p.GetServerCaCert()),
		IPv6Address:                 dcl.StringOrNil(p.Ipv6Address),
		ServiceAccountEmailAddress:  dcl.StringOrNil(p.ServiceAccountEmailAddress),
		OnPremisesConfiguration:     ProtoToSqlBetaInstanceOnPremisesConfiguration(p.GetOnPremisesConfiguration()),
		DiskEncryptionStatus:        ProtoToSqlBetaInstanceDiskEncryptionStatus(p.GetDiskEncryptionStatus()),
		InstanceUid:                 dcl.StringOrNil(p.InstanceUid),
	}
	for _, r := range p.GetIpAddresses() {
		obj.IPAddresses = append(obj.IPAddresses, *ProtoToSqlBetaInstanceIPAddresses(r))
	}
	for _, r := range p.GetReplicaInstances() {
		obj.ReplicaInstances = append(obj.ReplicaInstances, *ProtoToSqlBetaInstanceReplicaInstances(r))
	}
	for _, r := range p.GetSuspensionReason() {
		obj.SuspensionReason = append(obj.SuspensionReason, r)
	}
	return obj
}

// InstanceBackendTypeEnumToProto converts a InstanceBackendTypeEnum enum to its proto representation.
func SqlBetaInstanceBackendTypeEnumToProto(e *beta.InstanceBackendTypeEnum) betapb.SqlBetaInstanceBackendTypeEnum {
	if e == nil {
		return betapb.SqlBetaInstanceBackendTypeEnum(0)
	}
	if v, ok := betapb.SqlBetaInstanceBackendTypeEnum_value["InstanceBackendTypeEnum"+string(*e)]; ok {
		return betapb.SqlBetaInstanceBackendTypeEnum(v)
	}
	return betapb.SqlBetaInstanceBackendTypeEnum(0)
}

// InstanceDatabaseVersionEnumToProto converts a InstanceDatabaseVersionEnum enum to its proto representation.
func SqlBetaInstanceDatabaseVersionEnumToProto(e *beta.InstanceDatabaseVersionEnum) betapb.SqlBetaInstanceDatabaseVersionEnum {
	if e == nil {
		return betapb.SqlBetaInstanceDatabaseVersionEnum(0)
	}
	if v, ok := betapb.SqlBetaInstanceDatabaseVersionEnum_value["InstanceDatabaseVersionEnum"+string(*e)]; ok {
		return betapb.SqlBetaInstanceDatabaseVersionEnum(v)
	}
	return betapb.SqlBetaInstanceDatabaseVersionEnum(0)
}

// InstanceInstanceTypeEnumToProto converts a InstanceInstanceTypeEnum enum to its proto representation.
func SqlBetaInstanceInstanceTypeEnumToProto(e *beta.InstanceInstanceTypeEnum) betapb.SqlBetaInstanceInstanceTypeEnum {
	if e == nil {
		return betapb.SqlBetaInstanceInstanceTypeEnum(0)
	}
	if v, ok := betapb.SqlBetaInstanceInstanceTypeEnum_value["InstanceInstanceTypeEnum"+string(*e)]; ok {
		return betapb.SqlBetaInstanceInstanceTypeEnum(v)
	}
	return betapb.SqlBetaInstanceInstanceTypeEnum(0)
}

// InstanceIPAddressesTypeEnumToProto converts a InstanceIPAddressesTypeEnum enum to its proto representation.
func SqlBetaInstanceIPAddressesTypeEnumToProto(e *beta.InstanceIPAddressesTypeEnum) betapb.SqlBetaInstanceIPAddressesTypeEnum {
	if e == nil {
		return betapb.SqlBetaInstanceIPAddressesTypeEnum(0)
	}
	if v, ok := betapb.SqlBetaInstanceIPAddressesTypeEnum_value["InstanceIPAddressesTypeEnum"+string(*e)]; ok {
		return betapb.SqlBetaInstanceIPAddressesTypeEnum(v)
	}
	return betapb.SqlBetaInstanceIPAddressesTypeEnum(0)
}

// InstanceSettingsAvailabilityTypeEnumToProto converts a InstanceSettingsAvailabilityTypeEnum enum to its proto representation.
func SqlBetaInstanceSettingsAvailabilityTypeEnumToProto(e *beta.InstanceSettingsAvailabilityTypeEnum) betapb.SqlBetaInstanceSettingsAvailabilityTypeEnum {
	if e == nil {
		return betapb.SqlBetaInstanceSettingsAvailabilityTypeEnum(0)
	}
	if v, ok := betapb.SqlBetaInstanceSettingsAvailabilityTypeEnum_value["InstanceSettingsAvailabilityTypeEnum"+string(*e)]; ok {
		return betapb.SqlBetaInstanceSettingsAvailabilityTypeEnum(v)
	}
	return betapb.SqlBetaInstanceSettingsAvailabilityTypeEnum(0)
}

// InstanceSettingsPricingPlanEnumToProto converts a InstanceSettingsPricingPlanEnum enum to its proto representation.
func SqlBetaInstanceSettingsPricingPlanEnumToProto(e *beta.InstanceSettingsPricingPlanEnum) betapb.SqlBetaInstanceSettingsPricingPlanEnum {
	if e == nil {
		return betapb.SqlBetaInstanceSettingsPricingPlanEnum(0)
	}
	if v, ok := betapb.SqlBetaInstanceSettingsPricingPlanEnum_value["InstanceSettingsPricingPlanEnum"+string(*e)]; ok {
		return betapb.SqlBetaInstanceSettingsPricingPlanEnum(v)
	}
	return betapb.SqlBetaInstanceSettingsPricingPlanEnum(0)
}

// InstanceSettingsReplicationTypeEnumToProto converts a InstanceSettingsReplicationTypeEnum enum to its proto representation.
func SqlBetaInstanceSettingsReplicationTypeEnumToProto(e *beta.InstanceSettingsReplicationTypeEnum) betapb.SqlBetaInstanceSettingsReplicationTypeEnum {
	if e == nil {
		return betapb.SqlBetaInstanceSettingsReplicationTypeEnum(0)
	}
	if v, ok := betapb.SqlBetaInstanceSettingsReplicationTypeEnum_value["InstanceSettingsReplicationTypeEnum"+string(*e)]; ok {
		return betapb.SqlBetaInstanceSettingsReplicationTypeEnum(v)
	}
	return betapb.SqlBetaInstanceSettingsReplicationTypeEnum(0)
}

// InstanceSettingsActivationPolicyEnumToProto converts a InstanceSettingsActivationPolicyEnum enum to its proto representation.
func SqlBetaInstanceSettingsActivationPolicyEnumToProto(e *beta.InstanceSettingsActivationPolicyEnum) betapb.SqlBetaInstanceSettingsActivationPolicyEnum {
	if e == nil {
		return betapb.SqlBetaInstanceSettingsActivationPolicyEnum(0)
	}
	if v, ok := betapb.SqlBetaInstanceSettingsActivationPolicyEnum_value["InstanceSettingsActivationPolicyEnum"+string(*e)]; ok {
		return betapb.SqlBetaInstanceSettingsActivationPolicyEnum(v)
	}
	return betapb.SqlBetaInstanceSettingsActivationPolicyEnum(0)
}

// InstanceSettingsDataDiskTypeEnumToProto converts a InstanceSettingsDataDiskTypeEnum enum to its proto representation.
func SqlBetaInstanceSettingsDataDiskTypeEnumToProto(e *beta.InstanceSettingsDataDiskTypeEnum) betapb.SqlBetaInstanceSettingsDataDiskTypeEnum {
	if e == nil {
		return betapb.SqlBetaInstanceSettingsDataDiskTypeEnum(0)
	}
	if v, ok := betapb.SqlBetaInstanceSettingsDataDiskTypeEnum_value["InstanceSettingsDataDiskTypeEnum"+string(*e)]; ok {
		return betapb.SqlBetaInstanceSettingsDataDiskTypeEnum(v)
	}
	return betapb.SqlBetaInstanceSettingsDataDiskTypeEnum(0)
}

// InstanceSettingsMaintenanceWindowUpdateTrackEnumToProto converts a InstanceSettingsMaintenanceWindowUpdateTrackEnum enum to its proto representation.
func SqlBetaInstanceSettingsMaintenanceWindowUpdateTrackEnumToProto(e *beta.InstanceSettingsMaintenanceWindowUpdateTrackEnum) betapb.SqlBetaInstanceSettingsMaintenanceWindowUpdateTrackEnum {
	if e == nil {
		return betapb.SqlBetaInstanceSettingsMaintenanceWindowUpdateTrackEnum(0)
	}
	if v, ok := betapb.SqlBetaInstanceSettingsMaintenanceWindowUpdateTrackEnum_value["InstanceSettingsMaintenanceWindowUpdateTrackEnum"+string(*e)]; ok {
		return betapb.SqlBetaInstanceSettingsMaintenanceWindowUpdateTrackEnum(v)
	}
	return betapb.SqlBetaInstanceSettingsMaintenanceWindowUpdateTrackEnum(0)
}

// InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnumToProto converts a InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum enum to its proto representation.
func SqlBetaInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnumToProto(e *beta.InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum) betapb.SqlBetaInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum {
	if e == nil {
		return betapb.SqlBetaInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum(0)
	}
	if v, ok := betapb.SqlBetaInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum_value["InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum"+string(*e)]; ok {
		return betapb.SqlBetaInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum(v)
	}
	return betapb.SqlBetaInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum(0)
}

// InstanceMaxDiskSizeToProto converts a InstanceMaxDiskSize resource to its proto representation.
func SqlBetaInstanceMaxDiskSizeToProto(o *beta.InstanceMaxDiskSize) *betapb.SqlBetaInstanceMaxDiskSize {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceMaxDiskSize{
		Value: dcl.ValueOrEmptyInt64(o.Value),
	}
	return p
}

// InstanceCurrentDiskSizeToProto converts a InstanceCurrentDiskSize resource to its proto representation.
func SqlBetaInstanceCurrentDiskSizeToProto(o *beta.InstanceCurrentDiskSize) *betapb.SqlBetaInstanceCurrentDiskSize {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceCurrentDiskSize{
		Value: dcl.ValueOrEmptyInt64(o.Value),
	}
	return p
}

// InstanceDiskEncryptionConfigurationToProto converts a InstanceDiskEncryptionConfiguration resource to its proto representation.
func SqlBetaInstanceDiskEncryptionConfigurationToProto(o *beta.InstanceDiskEncryptionConfiguration) *betapb.SqlBetaInstanceDiskEncryptionConfiguration {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceDiskEncryptionConfiguration{
		KmsKeyName: dcl.ValueOrEmptyString(o.KmsKeyName),
		Kind:       dcl.ValueOrEmptyString(o.Kind),
	}
	return p
}

// InstanceFailoverReplicaToProto converts a InstanceFailoverReplica resource to its proto representation.
func SqlBetaInstanceFailoverReplicaToProto(o *beta.InstanceFailoverReplica) *betapb.SqlBetaInstanceFailoverReplica {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceFailoverReplica{
		Name:             dcl.ValueOrEmptyString(o.Name),
		Available:        dcl.ValueOrEmptyBool(o.Available),
		FailoverInstance: SqlBetaInstanceFailoverReplicaFailoverInstanceToProto(o.FailoverInstance),
	}
	return p
}

// InstanceFailoverReplicaFailoverInstanceToProto converts a InstanceFailoverReplicaFailoverInstance resource to its proto representation.
func SqlBetaInstanceFailoverReplicaFailoverInstanceToProto(o *beta.InstanceFailoverReplicaFailoverInstance) *betapb.SqlBetaInstanceFailoverReplicaFailoverInstance {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceFailoverReplicaFailoverInstance{
		Name:   dcl.ValueOrEmptyString(o.Name),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceIPAddressesToProto converts a InstanceIPAddresses resource to its proto representation.
func SqlBetaInstanceIPAddressesToProto(o *beta.InstanceIPAddresses) *betapb.SqlBetaInstanceIPAddresses {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceIPAddresses{
		Type:         SqlBetaInstanceIPAddressesTypeEnumToProto(o.Type),
		IpAddress:    dcl.ValueOrEmptyString(o.IPAddress),
		TimeToRetire: SqlBetaInstanceIPAddressesTimeToRetireToProto(o.TimeToRetire),
	}
	return p
}

// InstanceIPAddressesTimeToRetireToProto converts a InstanceIPAddressesTimeToRetire resource to its proto representation.
func SqlBetaInstanceIPAddressesTimeToRetireToProto(o *beta.InstanceIPAddressesTimeToRetire) *betapb.SqlBetaInstanceIPAddressesTimeToRetire {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceIPAddressesTimeToRetire{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// InstanceMasterInstanceToProto converts a InstanceMasterInstance resource to its proto representation.
func SqlBetaInstanceMasterInstanceToProto(o *beta.InstanceMasterInstance) *betapb.SqlBetaInstanceMasterInstance {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceMasterInstance{
		Name:   dcl.ValueOrEmptyString(o.Name),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceReplicaConfigurationToProto converts a InstanceReplicaConfiguration resource to its proto representation.
func SqlBetaInstanceReplicaConfigurationToProto(o *beta.InstanceReplicaConfiguration) *betapb.SqlBetaInstanceReplicaConfiguration {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceReplicaConfiguration{
		Kind:                      dcl.ValueOrEmptyString(o.Kind),
		MysqlReplicaConfiguration: SqlBetaInstanceReplicaConfigurationMysqlReplicaConfigurationToProto(o.MysqlReplicaConfiguration),
		FailoverTarget:            dcl.ValueOrEmptyBool(o.FailoverTarget),
		ReplicaPoolConfiguration:  SqlBetaInstanceReplicaConfigurationReplicaPoolConfigurationToProto(o.ReplicaPoolConfiguration),
	}
	return p
}

// InstanceReplicaConfigurationMysqlReplicaConfigurationToProto converts a InstanceReplicaConfigurationMysqlReplicaConfiguration resource to its proto representation.
func SqlBetaInstanceReplicaConfigurationMysqlReplicaConfigurationToProto(o *beta.InstanceReplicaConfigurationMysqlReplicaConfiguration) *betapb.SqlBetaInstanceReplicaConfigurationMysqlReplicaConfiguration {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceReplicaConfigurationMysqlReplicaConfiguration{
		DumpFilePath:            dcl.ValueOrEmptyString(o.DumpFilePath),
		Username:                dcl.ValueOrEmptyString(o.Username),
		Password:                dcl.ValueOrEmptyString(o.Password),
		ConnectRetryInterval:    dcl.ValueOrEmptyInt64(o.ConnectRetryInterval),
		MasterHeartbeatPeriod:   SqlBetaInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodToProto(o.MasterHeartbeatPeriod),
		CaCertificate:           dcl.ValueOrEmptyString(o.CaCertificate),
		ClientCertificate:       dcl.ValueOrEmptyString(o.ClientCertificate),
		ClientKey:               dcl.ValueOrEmptyString(o.ClientKey),
		SslCipher:               dcl.ValueOrEmptyString(o.SslCipher),
		VerifyServerCertificate: dcl.ValueOrEmptyBool(o.VerifyServerCertificate),
		Kind:                    dcl.ValueOrEmptyString(o.Kind),
	}
	return p
}

// InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodToProto converts a InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod resource to its proto representation.
func SqlBetaInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodToProto(o *beta.InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod) *betapb.SqlBetaInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod{
		Value: dcl.ValueOrEmptyInt64(o.Value),
	}
	return p
}

// InstanceReplicaConfigurationReplicaPoolConfigurationToProto converts a InstanceReplicaConfigurationReplicaPoolConfiguration resource to its proto representation.
func SqlBetaInstanceReplicaConfigurationReplicaPoolConfigurationToProto(o *beta.InstanceReplicaConfigurationReplicaPoolConfiguration) *betapb.SqlBetaInstanceReplicaConfigurationReplicaPoolConfiguration {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceReplicaConfigurationReplicaPoolConfiguration{
		Kind:                         dcl.ValueOrEmptyString(o.Kind),
		StaticPoolConfiguration:      SqlBetaInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationToProto(o.StaticPoolConfiguration),
		AutoscalingPoolConfiguration: SqlBetaInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationToProto(o.AutoscalingPoolConfiguration),
		ReplicaCount:                 dcl.ValueOrEmptyInt64(o.ReplicaCount),
		ExposeReplicaIp:              dcl.ValueOrEmptyBool(o.ExposeReplicaIP),
	}
	return p
}

// InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationToProto converts a InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration resource to its proto representation.
func SqlBetaInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationToProto(o *beta.InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration) *betapb.SqlBetaInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration{
		Kind:            dcl.ValueOrEmptyString(o.Kind),
		ReplicaCount:    dcl.ValueOrEmptyInt64(o.ReplicaCount),
		ExposeReplicaIp: dcl.ValueOrEmptyBool(o.ExposeReplicaIP),
	}
	return p
}

// InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationToProto converts a InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration resource to its proto representation.
func SqlBetaInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationToProto(o *beta.InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration) *betapb.SqlBetaInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration{
		Kind:            dcl.ValueOrEmptyString(o.Kind),
		MinReplicaCount: dcl.ValueOrEmptyInt64(o.MinReplicaCount),
		MaxReplicaCount: dcl.ValueOrEmptyInt64(o.MaxReplicaCount),
		TargetCpuUtil:   dcl.ValueOrEmptyDouble(o.TargetCpuUtil),
	}
	return p
}

// InstanceScheduledMaintenanceToProto converts a InstanceScheduledMaintenance resource to its proto representation.
func SqlBetaInstanceScheduledMaintenanceToProto(o *beta.InstanceScheduledMaintenance) *betapb.SqlBetaInstanceScheduledMaintenance {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceScheduledMaintenance{
		StartTime:     SqlBetaInstanceScheduledMaintenanceStartTimeToProto(o.StartTime),
		CanDefer:      dcl.ValueOrEmptyBool(o.CanDefer),
		CanReschedule: dcl.ValueOrEmptyBool(o.CanReschedule),
	}
	return p
}

// InstanceScheduledMaintenanceStartTimeToProto converts a InstanceScheduledMaintenanceStartTime resource to its proto representation.
func SqlBetaInstanceScheduledMaintenanceStartTimeToProto(o *beta.InstanceScheduledMaintenanceStartTime) *betapb.SqlBetaInstanceScheduledMaintenanceStartTime {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceScheduledMaintenanceStartTime{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// InstanceSettingsToProto converts a InstanceSettings resource to its proto representation.
func SqlBetaInstanceSettingsToProto(o *beta.InstanceSettings) *betapb.SqlBetaInstanceSettings {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceSettings{
		Tier:                        dcl.ValueOrEmptyString(o.Tier),
		Kind:                        dcl.ValueOrEmptyString(o.Kind),
		AvailabilityType:            SqlBetaInstanceSettingsAvailabilityTypeEnumToProto(o.AvailabilityType),
		PricingPlan:                 SqlBetaInstanceSettingsPricingPlanEnumToProto(o.PricingPlan),
		ReplicationType:             SqlBetaInstanceSettingsReplicationTypeEnumToProto(o.ReplicationType),
		ActivationPolicy:            SqlBetaInstanceSettingsActivationPolicyEnumToProto(o.ActivationPolicy),
		StorageAutoResize:           dcl.ValueOrEmptyBool(o.StorageAutoResize),
		DataDiskType:                SqlBetaInstanceSettingsDataDiskTypeEnumToProto(o.DataDiskType),
		DatabaseReplicationEnabled:  dcl.ValueOrEmptyBool(o.DatabaseReplicationEnabled),
		CrashSafeReplicationEnabled: dcl.ValueOrEmptyBool(o.CrashSafeReplicationEnabled),
		SettingsVersion:             SqlBetaInstanceSettingsSettingsVersionToProto(o.SettingsVersion),
		StorageAutoResizeLimit:      SqlBetaInstanceSettingsStorageAutoResizeLimitToProto(o.StorageAutoResizeLimit),
		IpConfiguration:             SqlBetaInstanceSettingsIPConfigurationToProto(o.IPConfiguration),
		LocationPreference:          SqlBetaInstanceSettingsLocationPreferenceToProto(o.LocationPreference),
		MaintenanceWindow:           SqlBetaInstanceSettingsMaintenanceWindowToProto(o.MaintenanceWindow),
		BackupConfiguration:         SqlBetaInstanceSettingsBackupConfigurationToProto(o.BackupConfiguration),
		DataDiskSizeGb:              SqlBetaInstanceSettingsDataDiskSizeGbToProto(o.DataDiskSizeGb),
		ActiveDirectoryConfig:       SqlBetaInstanceSettingsActiveDirectoryConfigToProto(o.ActiveDirectoryConfig),
		Collation:                   dcl.ValueOrEmptyString(o.Collation),
		InsightsConfig:              SqlBetaInstanceSettingsInsightsConfigToProto(o.InsightsConfig),
	}
	for _, r := range o.AuthorizedGaeApplications {
		p.AuthorizedGaeApplications = append(p.AuthorizedGaeApplications, r)
	}
	p.UserLabels = make(map[string]string)
	for k, r := range o.UserLabels {
		p.UserLabels[k] = r
	}
	for _, r := range o.DatabaseFlags {
		p.DatabaseFlags = append(p.DatabaseFlags, SqlBetaInstanceSettingsDatabaseFlagsToProto(&r))
	}
	for _, r := range o.DenyMaintenancePeriods {
		p.DenyMaintenancePeriods = append(p.DenyMaintenancePeriods, SqlBetaInstanceSettingsDenyMaintenancePeriodsToProto(&r))
	}
	return p
}

// InstanceSettingsSettingsVersionToProto converts a InstanceSettingsSettingsVersion resource to its proto representation.
func SqlBetaInstanceSettingsSettingsVersionToProto(o *beta.InstanceSettingsSettingsVersion) *betapb.SqlBetaInstanceSettingsSettingsVersion {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceSettingsSettingsVersion{
		Value: dcl.ValueOrEmptyInt64(o.Value),
	}
	return p
}

// InstanceSettingsStorageAutoResizeLimitToProto converts a InstanceSettingsStorageAutoResizeLimit resource to its proto representation.
func SqlBetaInstanceSettingsStorageAutoResizeLimitToProto(o *beta.InstanceSettingsStorageAutoResizeLimit) *betapb.SqlBetaInstanceSettingsStorageAutoResizeLimit {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceSettingsStorageAutoResizeLimit{
		Value: dcl.ValueOrEmptyInt64(o.Value),
	}
	return p
}

// InstanceSettingsIPConfigurationToProto converts a InstanceSettingsIPConfiguration resource to its proto representation.
func SqlBetaInstanceSettingsIPConfigurationToProto(o *beta.InstanceSettingsIPConfiguration) *betapb.SqlBetaInstanceSettingsIPConfiguration {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceSettingsIPConfiguration{
		Ipv4Enabled:    dcl.ValueOrEmptyBool(o.IPv4Enabled),
		PrivateNetwork: dcl.ValueOrEmptyString(o.PrivateNetwork),
		RequireSsl:     dcl.ValueOrEmptyBool(o.RequireSsl),
	}
	for _, r := range o.AuthorizedNetworks {
		p.AuthorizedNetworks = append(p.AuthorizedNetworks, SqlBetaInstanceSettingsIPConfigurationAuthorizedNetworksToProto(&r))
	}
	return p
}

// InstanceSettingsIPConfigurationAuthorizedNetworksToProto converts a InstanceSettingsIPConfigurationAuthorizedNetworks resource to its proto representation.
func SqlBetaInstanceSettingsIPConfigurationAuthorizedNetworksToProto(o *beta.InstanceSettingsIPConfigurationAuthorizedNetworks) *betapb.SqlBetaInstanceSettingsIPConfigurationAuthorizedNetworks {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceSettingsIPConfigurationAuthorizedNetworks{
		Value:          dcl.ValueOrEmptyString(o.Value),
		ExpirationTime: dcl.ValueOrEmptyString(o.ExpirationTime),
		Name:           dcl.ValueOrEmptyString(o.Name),
		Kind:           dcl.ValueOrEmptyString(o.Kind),
	}
	return p
}

// InstanceSettingsLocationPreferenceToProto converts a InstanceSettingsLocationPreference resource to its proto representation.
func SqlBetaInstanceSettingsLocationPreferenceToProto(o *beta.InstanceSettingsLocationPreference) *betapb.SqlBetaInstanceSettingsLocationPreference {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceSettingsLocationPreference{
		Zone: dcl.ValueOrEmptyString(o.Zone),
		Kind: dcl.ValueOrEmptyString(o.Kind),
	}
	return p
}

// InstanceSettingsDatabaseFlagsToProto converts a InstanceSettingsDatabaseFlags resource to its proto representation.
func SqlBetaInstanceSettingsDatabaseFlagsToProto(o *beta.InstanceSettingsDatabaseFlags) *betapb.SqlBetaInstanceSettingsDatabaseFlags {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceSettingsDatabaseFlags{
		Name:  dcl.ValueOrEmptyString(o.Name),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceSettingsMaintenanceWindowToProto converts a InstanceSettingsMaintenanceWindow resource to its proto representation.
func SqlBetaInstanceSettingsMaintenanceWindowToProto(o *beta.InstanceSettingsMaintenanceWindow) *betapb.SqlBetaInstanceSettingsMaintenanceWindow {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceSettingsMaintenanceWindow{
		Hour:        dcl.ValueOrEmptyInt64(o.Hour),
		Day:         dcl.ValueOrEmptyInt64(o.Day),
		UpdateTrack: SqlBetaInstanceSettingsMaintenanceWindowUpdateTrackEnumToProto(o.UpdateTrack),
		Kind:        dcl.ValueOrEmptyString(o.Kind),
	}
	return p
}

// InstanceSettingsBackupConfigurationToProto converts a InstanceSettingsBackupConfiguration resource to its proto representation.
func SqlBetaInstanceSettingsBackupConfigurationToProto(o *beta.InstanceSettingsBackupConfiguration) *betapb.SqlBetaInstanceSettingsBackupConfiguration {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceSettingsBackupConfiguration{
		StartTime:                   dcl.ValueOrEmptyString(o.StartTime),
		Enabled:                     dcl.ValueOrEmptyBool(o.Enabled),
		Kind:                        dcl.ValueOrEmptyString(o.Kind),
		BinaryLogEnabled:            dcl.ValueOrEmptyBool(o.BinaryLogEnabled),
		Location:                    dcl.ValueOrEmptyString(o.Location),
		BackupRetentionSettings:     SqlBetaInstanceSettingsBackupConfigurationBackupRetentionSettingsToProto(o.BackupRetentionSettings),
		TransactionLogRetentionDays: dcl.ValueOrEmptyInt64(o.TransactionLogRetentionDays),
	}
	return p
}

// InstanceSettingsBackupConfigurationBackupRetentionSettingsToProto converts a InstanceSettingsBackupConfigurationBackupRetentionSettings resource to its proto representation.
func SqlBetaInstanceSettingsBackupConfigurationBackupRetentionSettingsToProto(o *beta.InstanceSettingsBackupConfigurationBackupRetentionSettings) *betapb.SqlBetaInstanceSettingsBackupConfigurationBackupRetentionSettings {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceSettingsBackupConfigurationBackupRetentionSettings{
		RetentionUnit:   SqlBetaInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnumToProto(o.RetentionUnit),
		RetainedBackups: dcl.ValueOrEmptyInt64(o.RetainedBackups),
	}
	return p
}

// InstanceSettingsDataDiskSizeGbToProto converts a InstanceSettingsDataDiskSizeGb resource to its proto representation.
func SqlBetaInstanceSettingsDataDiskSizeGbToProto(o *beta.InstanceSettingsDataDiskSizeGb) *betapb.SqlBetaInstanceSettingsDataDiskSizeGb {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceSettingsDataDiskSizeGb{
		Value: dcl.ValueOrEmptyInt64(o.Value),
	}
	return p
}

// InstanceSettingsActiveDirectoryConfigToProto converts a InstanceSettingsActiveDirectoryConfig resource to its proto representation.
func SqlBetaInstanceSettingsActiveDirectoryConfigToProto(o *beta.InstanceSettingsActiveDirectoryConfig) *betapb.SqlBetaInstanceSettingsActiveDirectoryConfig {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceSettingsActiveDirectoryConfig{
		Kind:   dcl.ValueOrEmptyString(o.Kind),
		Domain: dcl.ValueOrEmptyString(o.Domain),
	}
	return p
}

// InstanceSettingsDenyMaintenancePeriodsToProto converts a InstanceSettingsDenyMaintenancePeriods resource to its proto representation.
func SqlBetaInstanceSettingsDenyMaintenancePeriodsToProto(o *beta.InstanceSettingsDenyMaintenancePeriods) *betapb.SqlBetaInstanceSettingsDenyMaintenancePeriods {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceSettingsDenyMaintenancePeriods{
		StartDate: dcl.ValueOrEmptyString(o.StartDate),
		EndDate:   dcl.ValueOrEmptyString(o.EndDate),
		Time:      dcl.ValueOrEmptyString(o.Time),
	}
	return p
}

// InstanceSettingsInsightsConfigToProto converts a InstanceSettingsInsightsConfig resource to its proto representation.
func SqlBetaInstanceSettingsInsightsConfigToProto(o *beta.InstanceSettingsInsightsConfig) *betapb.SqlBetaInstanceSettingsInsightsConfig {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceSettingsInsightsConfig{
		QueryInsightsEnabled:  dcl.ValueOrEmptyBool(o.QueryInsightsEnabled),
		RecordClientAddress:   dcl.ValueOrEmptyBool(o.RecordClientAddress),
		RecordApplicationTags: dcl.ValueOrEmptyBool(o.RecordApplicationTags),
		QueryStringLength:     dcl.ValueOrEmptyInt64(o.QueryStringLength),
	}
	return p
}

// InstanceReplicaInstancesToProto converts a InstanceReplicaInstances resource to its proto representation.
func SqlBetaInstanceReplicaInstancesToProto(o *beta.InstanceReplicaInstances) *betapb.SqlBetaInstanceReplicaInstances {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceReplicaInstances{
		Name:   dcl.ValueOrEmptyString(o.Name),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceServerCaCertToProto converts a InstanceServerCaCert resource to its proto representation.
func SqlBetaInstanceServerCaCertToProto(o *beta.InstanceServerCaCert) *betapb.SqlBetaInstanceServerCaCert {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceServerCaCert{
		Kind:             dcl.ValueOrEmptyString(o.Kind),
		CertSerialNumber: dcl.ValueOrEmptyString(o.CertSerialNumber),
		Cert:             dcl.ValueOrEmptyString(o.Cert),
		CreateTime:       dcl.ValueOrEmptyString(o.CreateTime),
		CommonName:       dcl.ValueOrEmptyString(o.CommonName),
		ExpirationTime:   dcl.ValueOrEmptyString(o.ExpirationTime),
		Sha1Fingerprint:  dcl.ValueOrEmptyString(o.Sha1Fingerprint),
		Instance:         dcl.ValueOrEmptyString(o.Instance),
	}
	return p
}

// InstanceOnPremisesConfigurationToProto converts a InstanceOnPremisesConfiguration resource to its proto representation.
func SqlBetaInstanceOnPremisesConfigurationToProto(o *beta.InstanceOnPremisesConfiguration) *betapb.SqlBetaInstanceOnPremisesConfiguration {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceOnPremisesConfiguration{
		HostPort:          dcl.ValueOrEmptyString(o.HostPort),
		Kind:              dcl.ValueOrEmptyString(o.Kind),
		Username:          dcl.ValueOrEmptyString(o.Username),
		Password:          dcl.ValueOrEmptyString(o.Password),
		CaCertificate:     dcl.ValueOrEmptyString(o.CaCertificate),
		ClientCertificate: dcl.ValueOrEmptyString(o.ClientCertificate),
		ClientKey:         dcl.ValueOrEmptyString(o.ClientKey),
		DumpFilePath:      dcl.ValueOrEmptyString(o.DumpFilePath),
		Database:          dcl.ValueOrEmptyString(o.Database),
	}
	for _, r := range o.ReplicatedDatabases {
		p.ReplicatedDatabases = append(p.ReplicatedDatabases, r)
	}
	return p
}

// InstanceDiskEncryptionStatusToProto converts a InstanceDiskEncryptionStatus resource to its proto representation.
func SqlBetaInstanceDiskEncryptionStatusToProto(o *beta.InstanceDiskEncryptionStatus) *betapb.SqlBetaInstanceDiskEncryptionStatus {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaInstanceDiskEncryptionStatus{
		KmsKeyVersionName: dcl.ValueOrEmptyString(o.KmsKeyVersionName),
		Kind:              dcl.ValueOrEmptyString(o.Kind),
	}
	return p
}

// InstanceToProto converts a Instance resource to its proto representation.
func InstanceToProto(resource *beta.Instance) *betapb.SqlBetaInstance {
	p := &betapb.SqlBetaInstance{
		BackendType:                 SqlBetaInstanceBackendTypeEnumToProto(resource.BackendType),
		ConnectionName:              dcl.ValueOrEmptyString(resource.ConnectionName),
		DatabaseVersion:             SqlBetaInstanceDatabaseVersionEnumToProto(resource.DatabaseVersion),
		Etag:                        dcl.ValueOrEmptyString(resource.Etag),
		GceZone:                     dcl.ValueOrEmptyString(resource.GceZone),
		InstanceType:                SqlBetaInstanceInstanceTypeEnumToProto(resource.InstanceType),
		MasterInstanceName:          dcl.ValueOrEmptyString(resource.MasterInstanceName),
		MaxDiskSize:                 SqlBetaInstanceMaxDiskSizeToProto(resource.MaxDiskSize),
		Name:                        dcl.ValueOrEmptyString(resource.Name),
		Project:                     dcl.ValueOrEmptyString(resource.Project),
		Region:                      dcl.ValueOrEmptyString(resource.Region),
		RootPassword:                dcl.ValueOrEmptyString(resource.RootPassword),
		CurrentDiskSize:             SqlBetaInstanceCurrentDiskSizeToProto(resource.CurrentDiskSize),
		DiskEncryptionConfiguration: SqlBetaInstanceDiskEncryptionConfigurationToProto(resource.DiskEncryptionConfiguration),
		FailoverReplica:             SqlBetaInstanceFailoverReplicaToProto(resource.FailoverReplica),
		MasterInstance:              SqlBetaInstanceMasterInstanceToProto(resource.MasterInstance),
		ReplicaConfiguration:        SqlBetaInstanceReplicaConfigurationToProto(resource.ReplicaConfiguration),
		ScheduledMaintenance:        SqlBetaInstanceScheduledMaintenanceToProto(resource.ScheduledMaintenance),
		Settings:                    SqlBetaInstanceSettingsToProto(resource.Settings),
		State:                       dcl.ValueOrEmptyString(resource.State),
		ServerCaCert:                SqlBetaInstanceServerCaCertToProto(resource.ServerCaCert),
		Ipv6Address:                 dcl.ValueOrEmptyString(resource.IPv6Address),
		ServiceAccountEmailAddress:  dcl.ValueOrEmptyString(resource.ServiceAccountEmailAddress),
		OnPremisesConfiguration:     SqlBetaInstanceOnPremisesConfigurationToProto(resource.OnPremisesConfiguration),
		DiskEncryptionStatus:        SqlBetaInstanceDiskEncryptionStatusToProto(resource.DiskEncryptionStatus),
		InstanceUid:                 dcl.ValueOrEmptyString(resource.InstanceUid),
	}
	for _, r := range resource.IPAddresses {
		p.IpAddresses = append(p.IpAddresses, SqlBetaInstanceIPAddressesToProto(&r))
	}
	for _, r := range resource.ReplicaInstances {
		p.ReplicaInstances = append(p.ReplicaInstances, SqlBetaInstanceReplicaInstancesToProto(&r))
	}
	for _, r := range resource.SuspensionReason {
		p.SuspensionReason = append(p.SuspensionReason, r)
	}

	return p
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) applyInstance(ctx context.Context, c *beta.Client, request *betapb.ApplySqlBetaInstanceRequest) (*betapb.SqlBetaInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplySqlBetaInstance(ctx context.Context, request *betapb.ApplySqlBetaInstanceRequest) (*betapb.SqlBetaInstance, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteSqlBetaInstance(ctx context.Context, request *betapb.DeleteSqlBetaInstanceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))

}

// ListSqlBetaInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListSqlBetaInstance(ctx context.Context, request *betapb.ListSqlBetaInstanceRequest) (*betapb.ListSqlBetaInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.SqlBetaInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListSqlBetaInstanceResponse{Items: protos}, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
