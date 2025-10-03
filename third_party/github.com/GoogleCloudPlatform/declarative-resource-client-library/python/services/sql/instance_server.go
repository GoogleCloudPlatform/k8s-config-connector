// Copyright 2020 Google LLC. All Rights Reserved.
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
	sqlpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/sql/sql_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/sql"
)

// Server implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstanceBackendTypeEnum converts a InstanceBackendTypeEnum enum from its proto representation.
func ProtoToSqlInstanceBackendTypeEnum(e sqlpb.SqlInstanceBackendTypeEnum) *sql.InstanceBackendTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := sqlpb.SqlInstanceBackendTypeEnum_name[int32(e)]; ok {
		e := sql.InstanceBackendTypeEnum(n[len("InstanceBackendTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceDatabaseVersionEnum converts a InstanceDatabaseVersionEnum enum from its proto representation.
func ProtoToSqlInstanceDatabaseVersionEnum(e sqlpb.SqlInstanceDatabaseVersionEnum) *sql.InstanceDatabaseVersionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := sqlpb.SqlInstanceDatabaseVersionEnum_name[int32(e)]; ok {
		e := sql.InstanceDatabaseVersionEnum(n[len("InstanceDatabaseVersionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceInstanceTypeEnum converts a InstanceInstanceTypeEnum enum from its proto representation.
func ProtoToSqlInstanceInstanceTypeEnum(e sqlpb.SqlInstanceInstanceTypeEnum) *sql.InstanceInstanceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := sqlpb.SqlInstanceInstanceTypeEnum_name[int32(e)]; ok {
		e := sql.InstanceInstanceTypeEnum(n[len("InstanceInstanceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceIPAddressesTypeEnum converts a InstanceIPAddressesTypeEnum enum from its proto representation.
func ProtoToSqlInstanceIPAddressesTypeEnum(e sqlpb.SqlInstanceIPAddressesTypeEnum) *sql.InstanceIPAddressesTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := sqlpb.SqlInstanceIPAddressesTypeEnum_name[int32(e)]; ok {
		e := sql.InstanceIPAddressesTypeEnum(n[len("InstanceIPAddressesTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceSettingsAvailabilityTypeEnum converts a InstanceSettingsAvailabilityTypeEnum enum from its proto representation.
func ProtoToSqlInstanceSettingsAvailabilityTypeEnum(e sqlpb.SqlInstanceSettingsAvailabilityTypeEnum) *sql.InstanceSettingsAvailabilityTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := sqlpb.SqlInstanceSettingsAvailabilityTypeEnum_name[int32(e)]; ok {
		e := sql.InstanceSettingsAvailabilityTypeEnum(n[len("InstanceSettingsAvailabilityTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceSettingsPricingPlanEnum converts a InstanceSettingsPricingPlanEnum enum from its proto representation.
func ProtoToSqlInstanceSettingsPricingPlanEnum(e sqlpb.SqlInstanceSettingsPricingPlanEnum) *sql.InstanceSettingsPricingPlanEnum {
	if e == 0 {
		return nil
	}
	if n, ok := sqlpb.SqlInstanceSettingsPricingPlanEnum_name[int32(e)]; ok {
		e := sql.InstanceSettingsPricingPlanEnum(n[len("InstanceSettingsPricingPlanEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceSettingsReplicationTypeEnum converts a InstanceSettingsReplicationTypeEnum enum from its proto representation.
func ProtoToSqlInstanceSettingsReplicationTypeEnum(e sqlpb.SqlInstanceSettingsReplicationTypeEnum) *sql.InstanceSettingsReplicationTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := sqlpb.SqlInstanceSettingsReplicationTypeEnum_name[int32(e)]; ok {
		e := sql.InstanceSettingsReplicationTypeEnum(n[len("InstanceSettingsReplicationTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceSettingsActivationPolicyEnum converts a InstanceSettingsActivationPolicyEnum enum from its proto representation.
func ProtoToSqlInstanceSettingsActivationPolicyEnum(e sqlpb.SqlInstanceSettingsActivationPolicyEnum) *sql.InstanceSettingsActivationPolicyEnum {
	if e == 0 {
		return nil
	}
	if n, ok := sqlpb.SqlInstanceSettingsActivationPolicyEnum_name[int32(e)]; ok {
		e := sql.InstanceSettingsActivationPolicyEnum(n[len("InstanceSettingsActivationPolicyEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceSettingsDataDiskTypeEnum converts a InstanceSettingsDataDiskTypeEnum enum from its proto representation.
func ProtoToSqlInstanceSettingsDataDiskTypeEnum(e sqlpb.SqlInstanceSettingsDataDiskTypeEnum) *sql.InstanceSettingsDataDiskTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := sqlpb.SqlInstanceSettingsDataDiskTypeEnum_name[int32(e)]; ok {
		e := sql.InstanceSettingsDataDiskTypeEnum(n[len("InstanceSettingsDataDiskTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceMaxDiskSize converts a InstanceMaxDiskSize resource from its proto representation.
func ProtoToSqlInstanceMaxDiskSize(p *sqlpb.SqlInstanceMaxDiskSize) *sql.InstanceMaxDiskSize {
	if p == nil {
		return nil
	}
	obj := &sql.InstanceMaxDiskSize{
		Value: dcl.Int64OrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceCurrentDiskSize converts a InstanceCurrentDiskSize resource from its proto representation.
func ProtoToSqlInstanceCurrentDiskSize(p *sqlpb.SqlInstanceCurrentDiskSize) *sql.InstanceCurrentDiskSize {
	if p == nil {
		return nil
	}
	obj := &sql.InstanceCurrentDiskSize{
		Value: dcl.Int64OrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceDiskEncryptionConfiguration converts a InstanceDiskEncryptionConfiguration resource from its proto representation.
func ProtoToSqlInstanceDiskEncryptionConfiguration(p *sqlpb.SqlInstanceDiskEncryptionConfiguration) *sql.InstanceDiskEncryptionConfiguration {
	if p == nil {
		return nil
	}
	obj := &sql.InstanceDiskEncryptionConfiguration{
		KmsKeyName: dcl.StringOrNil(p.KmsKeyName),
		Kind:       dcl.StringOrNil(p.Kind),
	}
	return obj
}

// ProtoToInstanceFailoverReplica converts a InstanceFailoverReplica resource from its proto representation.
func ProtoToSqlInstanceFailoverReplica(p *sqlpb.SqlInstanceFailoverReplica) *sql.InstanceFailoverReplica {
	if p == nil {
		return nil
	}
	obj := &sql.InstanceFailoverReplica{
		Name:      dcl.StringOrNil(p.Name),
		Available: dcl.Bool(p.Available),
	}
	return obj
}

// ProtoToInstanceIPAddresses converts a InstanceIPAddresses resource from its proto representation.
func ProtoToSqlInstanceIPAddresses(p *sqlpb.SqlInstanceIPAddresses) *sql.InstanceIPAddresses {
	if p == nil {
		return nil
	}
	obj := &sql.InstanceIPAddresses{
		Type:         ProtoToSqlInstanceIPAddressesTypeEnum(p.GetType()),
		IPAddress:    dcl.StringOrNil(p.IpAddress),
		TimeToRetire: ProtoToSqlInstanceIPAddressesTimeToRetire(p.GetTimeToRetire()),
	}
	return obj
}

// ProtoToInstanceIPAddressesTimeToRetire converts a InstanceIPAddressesTimeToRetire resource from its proto representation.
func ProtoToSqlInstanceIPAddressesTimeToRetire(p *sqlpb.SqlInstanceIPAddressesTimeToRetire) *sql.InstanceIPAddressesTimeToRetire {
	if p == nil {
		return nil
	}
	obj := &sql.InstanceIPAddressesTimeToRetire{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToInstanceMasterInstance converts a InstanceMasterInstance resource from its proto representation.
func ProtoToSqlInstanceMasterInstance(p *sqlpb.SqlInstanceMasterInstance) *sql.InstanceMasterInstance {
	if p == nil {
		return nil
	}
	obj := &sql.InstanceMasterInstance{
		Name:   dcl.StringOrNil(p.Name),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceReplicaConfiguration converts a InstanceReplicaConfiguration resource from its proto representation.
func ProtoToSqlInstanceReplicaConfiguration(p *sqlpb.SqlInstanceReplicaConfiguration) *sql.InstanceReplicaConfiguration {
	if p == nil {
		return nil
	}
	obj := &sql.InstanceReplicaConfiguration{
		Kind:                      dcl.StringOrNil(p.Kind),
		MysqlReplicaConfiguration: ProtoToSqlInstanceReplicaConfigurationMysqlReplicaConfiguration(p.GetMysqlReplicaConfiguration()),
		FailoverTarget:            dcl.Bool(p.FailoverTarget),
		ReplicaPoolConfiguration:  ProtoToSqlInstanceReplicaConfigurationReplicaPoolConfiguration(p.GetReplicaPoolConfiguration()),
	}
	return obj
}

// ProtoToInstanceReplicaConfigurationMysqlReplicaConfiguration converts a InstanceReplicaConfigurationMysqlReplicaConfiguration resource from its proto representation.
func ProtoToSqlInstanceReplicaConfigurationMysqlReplicaConfiguration(p *sqlpb.SqlInstanceReplicaConfigurationMysqlReplicaConfiguration) *sql.InstanceReplicaConfigurationMysqlReplicaConfiguration {
	if p == nil {
		return nil
	}
	obj := &sql.InstanceReplicaConfigurationMysqlReplicaConfiguration{
		DumpFilePath:            dcl.StringOrNil(p.DumpFilePath),
		Username:                dcl.StringOrNil(p.Username),
		Password:                dcl.StringOrNil(p.Password),
		ConnectRetryInterval:    dcl.Int64OrNil(p.ConnectRetryInterval),
		MasterHeartbeatPeriod:   ProtoToSqlInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(p.GetMasterHeartbeatPeriod()),
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
func ProtoToSqlInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(p *sqlpb.SqlInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod) *sql.InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod {
	if p == nil {
		return nil
	}
	obj := &sql.InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod{
		Value: dcl.Int64OrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceReplicaConfigurationReplicaPoolConfiguration converts a InstanceReplicaConfigurationReplicaPoolConfiguration resource from its proto representation.
func ProtoToSqlInstanceReplicaConfigurationReplicaPoolConfiguration(p *sqlpb.SqlInstanceReplicaConfigurationReplicaPoolConfiguration) *sql.InstanceReplicaConfigurationReplicaPoolConfiguration {
	if p == nil {
		return nil
	}
	obj := &sql.InstanceReplicaConfigurationReplicaPoolConfiguration{
		Kind:                         dcl.StringOrNil(p.Kind),
		StaticPoolConfiguration:      ProtoToSqlInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(p.GetStaticPoolConfiguration()),
		AutoscalingPoolConfiguration: ProtoToSqlInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(p.GetAutoscalingPoolConfiguration()),
	}
	return obj
}

// ProtoToInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration converts a InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration resource from its proto representation.
func ProtoToSqlInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(p *sqlpb.SqlInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration) *sql.InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration {
	if p == nil {
		return nil
	}
	obj := &sql.InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration{
		Kind:            dcl.StringOrNil(p.Kind),
		ReplicaCount:    dcl.Int64OrNil(p.ReplicaCount),
		ExposeReplicaIP: dcl.Bool(p.ExposeReplicaIp),
	}
	return obj
}

// ProtoToInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration converts a InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration resource from its proto representation.
func ProtoToSqlInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(p *sqlpb.SqlInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration) *sql.InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration {
	if p == nil {
		return nil
	}
	obj := &sql.InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration{
		Kind:            dcl.StringOrNil(p.Kind),
		MinReplicaCount: dcl.Int64OrNil(p.MinReplicaCount),
		MaxReplicaCount: dcl.Int64OrNil(p.MaxReplicaCount),
		TargetCpuUtil:   dcl.Float64OrNil(p.TargetCpuUtil),
	}
	return obj
}

// ProtoToInstanceScheduledMaintenance converts a InstanceScheduledMaintenance resource from its proto representation.
func ProtoToSqlInstanceScheduledMaintenance(p *sqlpb.SqlInstanceScheduledMaintenance) *sql.InstanceScheduledMaintenance {
	if p == nil {
		return nil
	}
	obj := &sql.InstanceScheduledMaintenance{
		StartTime:     ProtoToSqlInstanceScheduledMaintenanceStartTime(p.GetStartTime()),
		CanDefer:      dcl.Bool(p.CanDefer),
		CanReschedule: dcl.Bool(p.CanReschedule),
	}
	return obj
}

// ProtoToInstanceScheduledMaintenanceStartTime converts a InstanceScheduledMaintenanceStartTime resource from its proto representation.
func ProtoToSqlInstanceScheduledMaintenanceStartTime(p *sqlpb.SqlInstanceScheduledMaintenanceStartTime) *sql.InstanceScheduledMaintenanceStartTime {
	if p == nil {
		return nil
	}
	obj := &sql.InstanceScheduledMaintenanceStartTime{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToInstanceSettings converts a InstanceSettings resource from its proto representation.
func ProtoToSqlInstanceSettings(p *sqlpb.SqlInstanceSettings) *sql.InstanceSettings {
	if p == nil {
		return nil
	}
	obj := &sql.InstanceSettings{
		Tier:                        dcl.StringOrNil(p.Tier),
		Kind:                        dcl.StringOrNil(p.Kind),
		AvailabilityType:            ProtoToSqlInstanceSettingsAvailabilityTypeEnum(p.GetAvailabilityType()),
		PricingPlan:                 ProtoToSqlInstanceSettingsPricingPlanEnum(p.GetPricingPlan()),
		ReplicationType:             ProtoToSqlInstanceSettingsReplicationTypeEnum(p.GetReplicationType()),
		ActivationPolicy:            ProtoToSqlInstanceSettingsActivationPolicyEnum(p.GetActivationPolicy()),
		StorageAutoResize:           dcl.Bool(p.StorageAutoResize),
		DataDiskType:                ProtoToSqlInstanceSettingsDataDiskTypeEnum(p.GetDataDiskType()),
		DatabaseReplicationEnabled:  dcl.Bool(p.DatabaseReplicationEnabled),
		CrashSafeReplicationEnabled: dcl.Bool(p.CrashSafeReplicationEnabled),
	}
	for _, r := range p.GetAuthorizedGaeApplications() {
		obj.AuthorizedGaeApplications = append(obj.AuthorizedGaeApplications, r)
	}
	return obj
}

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *sqlpb.SqlInstance) *sql.Instance {
	obj := &sql.Instance{
		BackendType:                 ProtoToSqlInstanceBackendTypeEnum(p.GetBackendType()),
		ConnectionName:              dcl.StringOrNil(p.ConnectionName),
		DatabaseVersion:             ProtoToSqlInstanceDatabaseVersionEnum(p.GetDatabaseVersion()),
		Etag:                        dcl.StringOrNil(p.Etag),
		GceZone:                     dcl.StringOrNil(p.GceZone),
		InstanceType:                ProtoToSqlInstanceInstanceTypeEnum(p.GetInstanceType()),
		MasterInstanceName:          dcl.StringOrNil(p.MasterInstanceName),
		MaxDiskSize:                 ProtoToSqlInstanceMaxDiskSize(p.GetMaxDiskSize()),
		Name:                        dcl.StringOrNil(p.Name),
		Project:                     dcl.StringOrNil(p.Project),
		Region:                      dcl.StringOrNil(p.Region),
		RootPassword:                dcl.StringOrNil(p.RootPassword),
		CurrentDiskSize:             ProtoToSqlInstanceCurrentDiskSize(p.GetCurrentDiskSize()),
		DiskEncryptionConfiguration: ProtoToSqlInstanceDiskEncryptionConfiguration(p.GetDiskEncryptionConfiguration()),
		FailoverReplica:             ProtoToSqlInstanceFailoverReplica(p.GetFailoverReplica()),
		MasterInstance:              ProtoToSqlInstanceMasterInstance(p.GetMasterInstance()),
		ReplicaConfiguration:        ProtoToSqlInstanceReplicaConfiguration(p.GetReplicaConfiguration()),
		ScheduledMaintenance:        ProtoToSqlInstanceScheduledMaintenance(p.GetScheduledMaintenance()),
		Settings:                    ProtoToSqlInstanceSettings(p.GetSettings()),
	}
	for _, r := range p.GetIpAddresses() {
		obj.IPAddresses = append(obj.IPAddresses, *ProtoToSqlInstanceIPAddresses(r))
	}
	return obj
}

// InstanceBackendTypeEnumToProto converts a InstanceBackendTypeEnum enum to its proto representation.
func SqlInstanceBackendTypeEnumToProto(e *sql.InstanceBackendTypeEnum) sqlpb.SqlInstanceBackendTypeEnum {
	if e == nil {
		return sqlpb.SqlInstanceBackendTypeEnum(0)
	}
	if v, ok := sqlpb.SqlInstanceBackendTypeEnum_value["InstanceBackendTypeEnum"+string(*e)]; ok {
		return sqlpb.SqlInstanceBackendTypeEnum(v)
	}
	return sqlpb.SqlInstanceBackendTypeEnum(0)
}

// InstanceDatabaseVersionEnumToProto converts a InstanceDatabaseVersionEnum enum to its proto representation.
func SqlInstanceDatabaseVersionEnumToProto(e *sql.InstanceDatabaseVersionEnum) sqlpb.SqlInstanceDatabaseVersionEnum {
	if e == nil {
		return sqlpb.SqlInstanceDatabaseVersionEnum(0)
	}
	if v, ok := sqlpb.SqlInstanceDatabaseVersionEnum_value["InstanceDatabaseVersionEnum"+string(*e)]; ok {
		return sqlpb.SqlInstanceDatabaseVersionEnum(v)
	}
	return sqlpb.SqlInstanceDatabaseVersionEnum(0)
}

// InstanceInstanceTypeEnumToProto converts a InstanceInstanceTypeEnum enum to its proto representation.
func SqlInstanceInstanceTypeEnumToProto(e *sql.InstanceInstanceTypeEnum) sqlpb.SqlInstanceInstanceTypeEnum {
	if e == nil {
		return sqlpb.SqlInstanceInstanceTypeEnum(0)
	}
	if v, ok := sqlpb.SqlInstanceInstanceTypeEnum_value["InstanceInstanceTypeEnum"+string(*e)]; ok {
		return sqlpb.SqlInstanceInstanceTypeEnum(v)
	}
	return sqlpb.SqlInstanceInstanceTypeEnum(0)
}

// InstanceIPAddressesTypeEnumToProto converts a InstanceIPAddressesTypeEnum enum to its proto representation.
func SqlInstanceIPAddressesTypeEnumToProto(e *sql.InstanceIPAddressesTypeEnum) sqlpb.SqlInstanceIPAddressesTypeEnum {
	if e == nil {
		return sqlpb.SqlInstanceIPAddressesTypeEnum(0)
	}
	if v, ok := sqlpb.SqlInstanceIPAddressesTypeEnum_value["InstanceIPAddressesTypeEnum"+string(*e)]; ok {
		return sqlpb.SqlInstanceIPAddressesTypeEnum(v)
	}
	return sqlpb.SqlInstanceIPAddressesTypeEnum(0)
}

// InstanceSettingsAvailabilityTypeEnumToProto converts a InstanceSettingsAvailabilityTypeEnum enum to its proto representation.
func SqlInstanceSettingsAvailabilityTypeEnumToProto(e *sql.InstanceSettingsAvailabilityTypeEnum) sqlpb.SqlInstanceSettingsAvailabilityTypeEnum {
	if e == nil {
		return sqlpb.SqlInstanceSettingsAvailabilityTypeEnum(0)
	}
	if v, ok := sqlpb.SqlInstanceSettingsAvailabilityTypeEnum_value["InstanceSettingsAvailabilityTypeEnum"+string(*e)]; ok {
		return sqlpb.SqlInstanceSettingsAvailabilityTypeEnum(v)
	}
	return sqlpb.SqlInstanceSettingsAvailabilityTypeEnum(0)
}

// InstanceSettingsPricingPlanEnumToProto converts a InstanceSettingsPricingPlanEnum enum to its proto representation.
func SqlInstanceSettingsPricingPlanEnumToProto(e *sql.InstanceSettingsPricingPlanEnum) sqlpb.SqlInstanceSettingsPricingPlanEnum {
	if e == nil {
		return sqlpb.SqlInstanceSettingsPricingPlanEnum(0)
	}
	if v, ok := sqlpb.SqlInstanceSettingsPricingPlanEnum_value["InstanceSettingsPricingPlanEnum"+string(*e)]; ok {
		return sqlpb.SqlInstanceSettingsPricingPlanEnum(v)
	}
	return sqlpb.SqlInstanceSettingsPricingPlanEnum(0)
}

// InstanceSettingsReplicationTypeEnumToProto converts a InstanceSettingsReplicationTypeEnum enum to its proto representation.
func SqlInstanceSettingsReplicationTypeEnumToProto(e *sql.InstanceSettingsReplicationTypeEnum) sqlpb.SqlInstanceSettingsReplicationTypeEnum {
	if e == nil {
		return sqlpb.SqlInstanceSettingsReplicationTypeEnum(0)
	}
	if v, ok := sqlpb.SqlInstanceSettingsReplicationTypeEnum_value["InstanceSettingsReplicationTypeEnum"+string(*e)]; ok {
		return sqlpb.SqlInstanceSettingsReplicationTypeEnum(v)
	}
	return sqlpb.SqlInstanceSettingsReplicationTypeEnum(0)
}

// InstanceSettingsActivationPolicyEnumToProto converts a InstanceSettingsActivationPolicyEnum enum to its proto representation.
func SqlInstanceSettingsActivationPolicyEnumToProto(e *sql.InstanceSettingsActivationPolicyEnum) sqlpb.SqlInstanceSettingsActivationPolicyEnum {
	if e == nil {
		return sqlpb.SqlInstanceSettingsActivationPolicyEnum(0)
	}
	if v, ok := sqlpb.SqlInstanceSettingsActivationPolicyEnum_value["InstanceSettingsActivationPolicyEnum"+string(*e)]; ok {
		return sqlpb.SqlInstanceSettingsActivationPolicyEnum(v)
	}
	return sqlpb.SqlInstanceSettingsActivationPolicyEnum(0)
}

// InstanceSettingsDataDiskTypeEnumToProto converts a InstanceSettingsDataDiskTypeEnum enum to its proto representation.
func SqlInstanceSettingsDataDiskTypeEnumToProto(e *sql.InstanceSettingsDataDiskTypeEnum) sqlpb.SqlInstanceSettingsDataDiskTypeEnum {
	if e == nil {
		return sqlpb.SqlInstanceSettingsDataDiskTypeEnum(0)
	}
	if v, ok := sqlpb.SqlInstanceSettingsDataDiskTypeEnum_value["InstanceSettingsDataDiskTypeEnum"+string(*e)]; ok {
		return sqlpb.SqlInstanceSettingsDataDiskTypeEnum(v)
	}
	return sqlpb.SqlInstanceSettingsDataDiskTypeEnum(0)
}

// InstanceMaxDiskSizeToProto converts a InstanceMaxDiskSize resource to its proto representation.
func SqlInstanceMaxDiskSizeToProto(o *sql.InstanceMaxDiskSize) *sqlpb.SqlInstanceMaxDiskSize {
	if o == nil {
		return nil
	}
	p := &sqlpb.SqlInstanceMaxDiskSize{
		Value: dcl.ValueOrEmptyInt64(o.Value),
	}
	return p
}

// InstanceCurrentDiskSizeToProto converts a InstanceCurrentDiskSize resource to its proto representation.
func SqlInstanceCurrentDiskSizeToProto(o *sql.InstanceCurrentDiskSize) *sqlpb.SqlInstanceCurrentDiskSize {
	if o == nil {
		return nil
	}
	p := &sqlpb.SqlInstanceCurrentDiskSize{
		Value: dcl.ValueOrEmptyInt64(o.Value),
	}
	return p
}

// InstanceDiskEncryptionConfigurationToProto converts a InstanceDiskEncryptionConfiguration resource to its proto representation.
func SqlInstanceDiskEncryptionConfigurationToProto(o *sql.InstanceDiskEncryptionConfiguration) *sqlpb.SqlInstanceDiskEncryptionConfiguration {
	if o == nil {
		return nil
	}
	p := &sqlpb.SqlInstanceDiskEncryptionConfiguration{
		KmsKeyName: dcl.ValueOrEmptyString(o.KmsKeyName),
		Kind:       dcl.ValueOrEmptyString(o.Kind),
	}
	return p
}

// InstanceFailoverReplicaToProto converts a InstanceFailoverReplica resource to its proto representation.
func SqlInstanceFailoverReplicaToProto(o *sql.InstanceFailoverReplica) *sqlpb.SqlInstanceFailoverReplica {
	if o == nil {
		return nil
	}
	p := &sqlpb.SqlInstanceFailoverReplica{
		Name:      dcl.ValueOrEmptyString(o.Name),
		Available: dcl.ValueOrEmptyBool(o.Available),
	}
	return p
}

// InstanceIPAddressesToProto converts a InstanceIPAddresses resource to its proto representation.
func SqlInstanceIPAddressesToProto(o *sql.InstanceIPAddresses) *sqlpb.SqlInstanceIPAddresses {
	if o == nil {
		return nil
	}
	p := &sqlpb.SqlInstanceIPAddresses{
		Type:         SqlInstanceIPAddressesTypeEnumToProto(o.Type),
		IpAddress:    dcl.ValueOrEmptyString(o.IPAddress),
		TimeToRetire: SqlInstanceIPAddressesTimeToRetireToProto(o.TimeToRetire),
	}
	return p
}

// InstanceIPAddressesTimeToRetireToProto converts a InstanceIPAddressesTimeToRetire resource to its proto representation.
func SqlInstanceIPAddressesTimeToRetireToProto(o *sql.InstanceIPAddressesTimeToRetire) *sqlpb.SqlInstanceIPAddressesTimeToRetire {
	if o == nil {
		return nil
	}
	p := &sqlpb.SqlInstanceIPAddressesTimeToRetire{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// InstanceMasterInstanceToProto converts a InstanceMasterInstance resource to its proto representation.
func SqlInstanceMasterInstanceToProto(o *sql.InstanceMasterInstance) *sqlpb.SqlInstanceMasterInstance {
	if o == nil {
		return nil
	}
	p := &sqlpb.SqlInstanceMasterInstance{
		Name:   dcl.ValueOrEmptyString(o.Name),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceReplicaConfigurationToProto converts a InstanceReplicaConfiguration resource to its proto representation.
func SqlInstanceReplicaConfigurationToProto(o *sql.InstanceReplicaConfiguration) *sqlpb.SqlInstanceReplicaConfiguration {
	if o == nil {
		return nil
	}
	p := &sqlpb.SqlInstanceReplicaConfiguration{
		Kind:                      dcl.ValueOrEmptyString(o.Kind),
		MysqlReplicaConfiguration: SqlInstanceReplicaConfigurationMysqlReplicaConfigurationToProto(o.MysqlReplicaConfiguration),
		FailoverTarget:            dcl.ValueOrEmptyBool(o.FailoverTarget),
		ReplicaPoolConfiguration:  SqlInstanceReplicaConfigurationReplicaPoolConfigurationToProto(o.ReplicaPoolConfiguration),
	}
	return p
}

// InstanceReplicaConfigurationMysqlReplicaConfigurationToProto converts a InstanceReplicaConfigurationMysqlReplicaConfiguration resource to its proto representation.
func SqlInstanceReplicaConfigurationMysqlReplicaConfigurationToProto(o *sql.InstanceReplicaConfigurationMysqlReplicaConfiguration) *sqlpb.SqlInstanceReplicaConfigurationMysqlReplicaConfiguration {
	if o == nil {
		return nil
	}
	p := &sqlpb.SqlInstanceReplicaConfigurationMysqlReplicaConfiguration{
		DumpFilePath:            dcl.ValueOrEmptyString(o.DumpFilePath),
		Username:                dcl.ValueOrEmptyString(o.Username),
		Password:                dcl.ValueOrEmptyString(o.Password),
		ConnectRetryInterval:    dcl.ValueOrEmptyInt64(o.ConnectRetryInterval),
		MasterHeartbeatPeriod:   SqlInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodToProto(o.MasterHeartbeatPeriod),
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
func SqlInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodToProto(o *sql.InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod) *sqlpb.SqlInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod {
	if o == nil {
		return nil
	}
	p := &sqlpb.SqlInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod{
		Value: dcl.ValueOrEmptyInt64(o.Value),
	}
	return p
}

// InstanceReplicaConfigurationReplicaPoolConfigurationToProto converts a InstanceReplicaConfigurationReplicaPoolConfiguration resource to its proto representation.
func SqlInstanceReplicaConfigurationReplicaPoolConfigurationToProto(o *sql.InstanceReplicaConfigurationReplicaPoolConfiguration) *sqlpb.SqlInstanceReplicaConfigurationReplicaPoolConfiguration {
	if o == nil {
		return nil
	}
	p := &sqlpb.SqlInstanceReplicaConfigurationReplicaPoolConfiguration{
		Kind:                         dcl.ValueOrEmptyString(o.Kind),
		StaticPoolConfiguration:      SqlInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationToProto(o.StaticPoolConfiguration),
		AutoscalingPoolConfiguration: SqlInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationToProto(o.AutoscalingPoolConfiguration),
	}
	return p
}

// InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationToProto converts a InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration resource to its proto representation.
func SqlInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationToProto(o *sql.InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration) *sqlpb.SqlInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration {
	if o == nil {
		return nil
	}
	p := &sqlpb.SqlInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration{
		Kind:            dcl.ValueOrEmptyString(o.Kind),
		ReplicaCount:    dcl.ValueOrEmptyInt64(o.ReplicaCount),
		ExposeReplicaIp: dcl.ValueOrEmptyBool(o.ExposeReplicaIP),
	}
	return p
}

// InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationToProto converts a InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration resource to its proto representation.
func SqlInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationToProto(o *sql.InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration) *sqlpb.SqlInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration {
	if o == nil {
		return nil
	}
	p := &sqlpb.SqlInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration{
		Kind:            dcl.ValueOrEmptyString(o.Kind),
		MinReplicaCount: dcl.ValueOrEmptyInt64(o.MinReplicaCount),
		MaxReplicaCount: dcl.ValueOrEmptyInt64(o.MaxReplicaCount),
		TargetCpuUtil:   dcl.ValueOrEmptyDouble(o.TargetCpuUtil),
	}
	return p
}

// InstanceScheduledMaintenanceToProto converts a InstanceScheduledMaintenance resource to its proto representation.
func SqlInstanceScheduledMaintenanceToProto(o *sql.InstanceScheduledMaintenance) *sqlpb.SqlInstanceScheduledMaintenance {
	if o == nil {
		return nil
	}
	p := &sqlpb.SqlInstanceScheduledMaintenance{
		StartTime:     SqlInstanceScheduledMaintenanceStartTimeToProto(o.StartTime),
		CanDefer:      dcl.ValueOrEmptyBool(o.CanDefer),
		CanReschedule: dcl.ValueOrEmptyBool(o.CanReschedule),
	}
	return p
}

// InstanceScheduledMaintenanceStartTimeToProto converts a InstanceScheduledMaintenanceStartTime resource to its proto representation.
func SqlInstanceScheduledMaintenanceStartTimeToProto(o *sql.InstanceScheduledMaintenanceStartTime) *sqlpb.SqlInstanceScheduledMaintenanceStartTime {
	if o == nil {
		return nil
	}
	p := &sqlpb.SqlInstanceScheduledMaintenanceStartTime{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// InstanceSettingsToProto converts a InstanceSettings resource to its proto representation.
func SqlInstanceSettingsToProto(o *sql.InstanceSettings) *sqlpb.SqlInstanceSettings {
	if o == nil {
		return nil
	}
	p := &sqlpb.SqlInstanceSettings{
		Tier:                        dcl.ValueOrEmptyString(o.Tier),
		Kind:                        dcl.ValueOrEmptyString(o.Kind),
		AvailabilityType:            SqlInstanceSettingsAvailabilityTypeEnumToProto(o.AvailabilityType),
		PricingPlan:                 SqlInstanceSettingsPricingPlanEnumToProto(o.PricingPlan),
		ReplicationType:             SqlInstanceSettingsReplicationTypeEnumToProto(o.ReplicationType),
		ActivationPolicy:            SqlInstanceSettingsActivationPolicyEnumToProto(o.ActivationPolicy),
		StorageAutoResize:           dcl.ValueOrEmptyBool(o.StorageAutoResize),
		DataDiskType:                SqlInstanceSettingsDataDiskTypeEnumToProto(o.DataDiskType),
		DatabaseReplicationEnabled:  dcl.ValueOrEmptyBool(o.DatabaseReplicationEnabled),
		CrashSafeReplicationEnabled: dcl.ValueOrEmptyBool(o.CrashSafeReplicationEnabled),
	}
	for _, r := range o.AuthorizedGaeApplications {
		p.AuthorizedGaeApplications = append(p.AuthorizedGaeApplications, r)
	}
	return p
}

// InstanceToProto converts a Instance resource to its proto representation.
func InstanceToProto(resource *sql.Instance) *sqlpb.SqlInstance {
	p := &sqlpb.SqlInstance{
		BackendType:                 SqlInstanceBackendTypeEnumToProto(resource.BackendType),
		ConnectionName:              dcl.ValueOrEmptyString(resource.ConnectionName),
		DatabaseVersion:             SqlInstanceDatabaseVersionEnumToProto(resource.DatabaseVersion),
		Etag:                        dcl.ValueOrEmptyString(resource.Etag),
		GceZone:                     dcl.ValueOrEmptyString(resource.GceZone),
		InstanceType:                SqlInstanceInstanceTypeEnumToProto(resource.InstanceType),
		MasterInstanceName:          dcl.ValueOrEmptyString(resource.MasterInstanceName),
		MaxDiskSize:                 SqlInstanceMaxDiskSizeToProto(resource.MaxDiskSize),
		Name:                        dcl.ValueOrEmptyString(resource.Name),
		Project:                     dcl.ValueOrEmptyString(resource.Project),
		Region:                      dcl.ValueOrEmptyString(resource.Region),
		RootPassword:                dcl.ValueOrEmptyString(resource.RootPassword),
		CurrentDiskSize:             SqlInstanceCurrentDiskSizeToProto(resource.CurrentDiskSize),
		DiskEncryptionConfiguration: SqlInstanceDiskEncryptionConfigurationToProto(resource.DiskEncryptionConfiguration),
		FailoverReplica:             SqlInstanceFailoverReplicaToProto(resource.FailoverReplica),
		MasterInstance:              SqlInstanceMasterInstanceToProto(resource.MasterInstance),
		ReplicaConfiguration:        SqlInstanceReplicaConfigurationToProto(resource.ReplicaConfiguration),
		ScheduledMaintenance:        SqlInstanceScheduledMaintenanceToProto(resource.ScheduledMaintenance),
		Settings:                    SqlInstanceSettingsToProto(resource.Settings),
	}
	for _, r := range resource.IPAddresses {
		p.IpAddresses = append(p.IpAddresses, SqlInstanceIPAddressesToProto(&r))
	}

	return p
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) applyInstance(ctx context.Context, c *sql.Client, request *sqlpb.ApplySqlInstanceRequest) (*sqlpb.SqlInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplySqlInstance(ctx context.Context, request *sqlpb.ApplySqlInstanceRequest) (*sqlpb.SqlInstance, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteSqlInstance(ctx context.Context, request *sqlpb.DeleteSqlInstanceRequest) (*emptypb.Empty, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))
}

// ListInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListSqlInstance(ctx context.Context, request *sqlpb.ListSqlInstanceRequest) (*sqlpb.ListSqlInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*sqlpb.SqlInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	return &sqlpb.ListSqlInstanceResponse{Items: protos}, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*sql.Client, error) {

	client, err := dcl.FromCredentialsFile(ctx, service_account_file)
	if err != nil {
		return nil, err
	}

	conf := dcl.NewConfig(client, dcl.WithUserAgent("dcl-test"))
	if err != nil {
		return nil, err
	}
	return sql.NewClient(conf), nil
}
