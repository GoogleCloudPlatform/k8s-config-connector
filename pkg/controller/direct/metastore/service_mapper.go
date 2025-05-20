// Copyright 2025 Google LLC
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

// +generated:mapper
// krm.group: metastore.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.metastore.v1

package metastore

import (
	"fmt"
	"strconv"

	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"

	pb "cloud.google.com/go/metastore/apiv1/metastorepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/metastore/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	secretmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	dayofweek "google.golang.org/genproto/googleapis/type/dayofweek"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func AuxiliaryVersionConfig_FromProto(mapCtx *direct.MapContext, in *pb.AuxiliaryVersionConfig) *krm.AuxiliaryVersionConfig {
	if in == nil {
		return nil
	}
	out := &krm.AuxiliaryVersionConfig{}
	out.Version = direct.LazyPtr(in.GetVersion())
	out.ConfigOverrides = in.ConfigOverrides
	// MISSING: NetworkConfig
	return out
}
func AuxiliaryVersionConfig_ToProto(mapCtx *direct.MapContext, in *krm.AuxiliaryVersionConfig) *pb.AuxiliaryVersionConfig {
	if in == nil {
		return nil
	}
	out := &pb.AuxiliaryVersionConfig{}
	out.Version = direct.ValueOf(in.Version)
	out.ConfigOverrides = in.ConfigOverrides
	// MISSING: NetworkConfig
	return out
}
func BackendMetastore_FromProto(mapCtx *direct.MapContext, in *pb.BackendMetastore) *krm.BackendMetastore {
	if in == nil {
		return nil
	}
	out := &krm.BackendMetastore{}
	out.ServiceRef = &krm.ServiceRef{
		External: in.GetName(),
	}
	out.MetastoreType = direct.Enum_FromProto(mapCtx, in.GetMetastoreType())
	return out
}
func BackendMetastore_ToProto(mapCtx *direct.MapContext, in *krm.BackendMetastore) *pb.BackendMetastore {
	if in == nil {
		return nil
	}
	out := &pb.BackendMetastore{}
	out.Name = in.ServiceRef.External
	out.MetastoreType = direct.Enum_ToProto[pb.BackendMetastore_MetastoreType](mapCtx, in.MetastoreType)
	return out
}
func EncryptionConfig_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionConfig) *krm.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionConfig{}
	out.KMSKeyRef = &kmsv1beta1.KMSKeyRef_OneOf{External: in.GetKmsKey()}
	return out
}
func EncryptionConfig_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionConfig) *pb.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionConfig{}
	out.KmsKey = in.KMSKeyRef.External
	return out
}
func HiveMetastoreConfig_FromProto(mapCtx *direct.MapContext, in *pb.HiveMetastoreConfig) *krm.HiveMetastoreConfig {
	if in == nil {
		return nil
	}
	out := &krm.HiveMetastoreConfig{}
	out.Version = direct.LazyPtr(in.GetVersion())
	out.ConfigOverrides = in.ConfigOverrides
	out.KerberosConfig = KerberosConfig_FromProto(mapCtx, in.GetKerberosConfig())
	out.EndpointProtocol = direct.Enum_FromProto(mapCtx, in.GetEndpointProtocol())
	// MISSING: AuxiliaryVersions
	return out
}
func HiveMetastoreConfig_ToProto(mapCtx *direct.MapContext, in *krm.HiveMetastoreConfig) *pb.HiveMetastoreConfig {
	if in == nil {
		return nil
	}
	out := &pb.HiveMetastoreConfig{}
	out.Version = direct.ValueOf(in.Version)
	out.ConfigOverrides = in.ConfigOverrides
	out.KerberosConfig = KerberosConfig_ToProto(mapCtx, in.KerberosConfig)
	out.EndpointProtocol = direct.Enum_ToProto[pb.HiveMetastoreConfig_EndpointProtocol](mapCtx, in.EndpointProtocol)
	// MISSING: AuxiliaryVersions
	return out
}
func KerberosConfig_FromProto(mapCtx *direct.MapContext, in *pb.KerberosConfig) *krm.KerberosConfig {
	if in == nil {
		return nil
	}
	out := &krm.KerberosConfig{}
	out.Keytab = Secret_FromProto(mapCtx, in.GetKeytab())
	out.Principal = direct.LazyPtr(in.GetPrincipal())
	out.Krb5ConfigGCSURI = direct.LazyPtr(in.GetKrb5ConfigGcsUri())
	return out
}
func KerberosConfig_ToProto(mapCtx *direct.MapContext, in *krm.KerberosConfig) *pb.KerberosConfig {
	if in == nil {
		return nil
	}
	out := &pb.KerberosConfig{}
	out.Keytab = Secret_ToProto(mapCtx, in.Keytab)
	out.Principal = direct.ValueOf(in.Principal)
	out.Krb5ConfigGcsUri = direct.ValueOf(in.Krb5ConfigGCSURI)
	return out
}
func MaintenanceWindow_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceWindow) *krm.MaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &krm.MaintenanceWindow{}
	if in.GetHourOfDay() != nil {
		value := in.GetHourOfDay().GetValue()
		out.HourOfDay = &krm.Int32Value{
			Value: &value,
		}
	}
	out.DayOfWeek = direct.LazyPtr(in.GetDayOfWeek().String())
	return out
}
func MaintenanceWindow_ToProto(mapCtx *direct.MapContext, in *krm.MaintenanceWindow) *pb.MaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceWindow{}
	if in.HourOfDay != nil && in.HourOfDay.Value != nil {
		out.HourOfDay = &wrapperspb.Int32Value{Value: *in.HourOfDay.Value}
	}
	if in.DayOfWeek != nil {
		for enumValue, enumName := range dayofweek.DayOfWeek_name {
			if enumName == *in.DayOfWeek {
				dow := dayofweek.DayOfWeek(enumValue)
				out.DayOfWeek = dow
				break
			}
		}
	}
	return out
}
func MetadataExport_FromProto(mapCtx *direct.MapContext, in *pb.MetadataExport) *krm.MetadataExport {
	if in == nil {
		return nil
	}
	out := &krm.MetadataExport{}
	// MISSING: DestinationGCSURI
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	// MISSING: DatabaseDumpType
	return out
}
func MetadataExport_ToProto(mapCtx *direct.MapContext, in *krm.MetadataExport) *pb.MetadataExport {
	if in == nil {
		return nil
	}
	out := &pb.MetadataExport{}
	// MISSING: DestinationGCSURI
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	// MISSING: DatabaseDumpType
	return out
}
func MetadataExportObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MetadataExport) *krm.MetadataExportObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MetadataExportObservedState{}
	out.DestinationGCSURI = direct.LazyPtr(in.GetDestinationGcsUri())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.DatabaseDumpType = direct.Enum_FromProto(mapCtx, in.GetDatabaseDumpType())
	return out
}
func MetadataExportObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MetadataExportObservedState) *pb.MetadataExport {
	if in == nil {
		return nil
	}
	out := &pb.MetadataExport{}
	if in.DestinationGCSURI != nil {
		out.Destination = &pb.MetadataExport_DestinationGcsUri{
			DestinationGcsUri: *in.DestinationGCSURI,
		}
	}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.State = direct.Enum_ToProto[pb.MetadataExport_State](mapCtx, in.State)
	out.DatabaseDumpType = direct.Enum_ToProto[pb.DatabaseDumpSpec_Type](mapCtx, in.DatabaseDumpType)
	return out
}
func MetadataManagementActivity_FromProto(mapCtx *direct.MapContext, in *pb.MetadataManagementActivity) *krm.MetadataManagementActivity {
	if in == nil {
		return nil
	}
	out := &krm.MetadataManagementActivity{}
	// MISSING: MetadataExports
	// MISSING: Restores
	return out
}
func MetadataManagementActivity_ToProto(mapCtx *direct.MapContext, in *krm.MetadataManagementActivity) *pb.MetadataManagementActivity {
	if in == nil {
		return nil
	}
	out := &pb.MetadataManagementActivity{}
	// MISSING: MetadataExports
	// MISSING: Restores
	return out
}
func MetadataManagementActivityObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MetadataManagementActivity) *krm.MetadataManagementActivityObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MetadataManagementActivityObservedState{}
	out.MetadataExports = direct.Slice_FromProto(mapCtx, in.MetadataExports, MetadataExportObservedState_FromProto)
	out.Restores = direct.Slice_FromProto(mapCtx, in.Restores, RestoreObservedState_FromProto)
	return out
}
func MetadataManagementActivityObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MetadataManagementActivityObservedState) *pb.MetadataManagementActivity {
	if in == nil {
		return nil
	}
	out := &pb.MetadataManagementActivity{}
	out.MetadataExports = direct.Slice_ToProto(mapCtx, in.MetadataExports, MetadataExportObservedState_ToProto)
	out.Restores = direct.Slice_ToProto(mapCtx, in.Restores, RestoreObservedState_ToProto)
	return out
}

func MetastoreFederationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Federation) *krm.MetastoreFederationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MetastoreFederationObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.EndpointURI = direct.LazyPtr(in.GetEndpointUri())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateMessage = direct.LazyPtr(in.GetStateMessage())
	out.Uid = direct.LazyPtr(in.GetUid())
	return out
}
func MetastoreFederationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MetastoreFederationObservedState) *pb.Federation {
	if in == nil {
		return nil
	}
	out := &pb.Federation{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.EndpointUri = direct.ValueOf(in.EndpointURI)
	out.State = direct.Enum_ToProto[pb.Federation_State](mapCtx, in.State)
	out.StateMessage = direct.ValueOf(in.StateMessage)
	out.Uid = direct.ValueOf(in.Uid)
	return out
}
func MetastoreFederationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Federation) *krm.MetastoreFederationSpec {
	if in == nil {
		return nil
	}
	out := &krm.MetastoreFederationSpec{}
	// MISSING: Name
	out.Labels = in.Labels
	out.Version = direct.LazyPtr(in.GetVersion())
	out.BackendMetastores = make(map[string]krm.BackendMetastore)
	for k, v := range in.GetBackendMetastores() {
		out.BackendMetastores[fmt.Sprintf("%d", k)] = *BackendMetastore_FromProto(mapCtx, v)
	}
	return out
}
func MetastoreFederationSpec_ToProto(mapCtx *direct.MapContext, in *krm.MetastoreFederationSpec) *pb.Federation {
	if in == nil {
		return nil
	}
	out := &pb.Federation{}
	// MISSING: Name
	out.Labels = in.Labels
	out.Version = direct.ValueOf(in.Version)
	// TODO: map type int32 message for field BackendMetastores
	out.BackendMetastores = make(map[int32]*pb.BackendMetastore)
	for k, v := range in.BackendMetastores {
		ik, err := strconv.ParseInt(k, 10, 32)
		if err != nil {
			// we prevalidate in create and update.
			// so we can safely ignore this error
			continue
		}
		out.BackendMetastores[int32(ik)] = BackendMetastore_ToProto(mapCtx, &v)
	}

	return out
}

func MetastoreServiceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Service) *krm.MetastoreServiceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MetastoreServiceObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.EndpointURI = direct.LazyPtr(in.GetEndpointUri())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateMessage = direct.LazyPtr(in.GetStateMessage())
	out.ArtifactGCSURI = direct.LazyPtr(in.GetArtifactGcsUri())
	out.UID = direct.LazyPtr(in.GetUid())
	out.MetadataManagementActivity = MetadataManagementActivityObservedState_FromProto(mapCtx, in.GetMetadataManagementActivity())
	out.NetworkConfig = NetworkConfigObservedState_FromProto(mapCtx, in.GetNetworkConfig())
	return out
}
func MetastoreServiceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MetastoreServiceObservedState) *pb.Service {
	if in == nil {
		return nil
	}
	out := &pb.Service{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.EndpointUri = direct.ValueOf(in.EndpointURI)
	out.State = direct.Enum_ToProto[pb.Service_State](mapCtx, in.State)
	out.StateMessage = direct.ValueOf(in.StateMessage)
	out.ArtifactGcsUri = direct.ValueOf(in.ArtifactGCSURI)
	out.Uid = direct.ValueOf(in.UID)
	out.MetadataManagementActivity = MetadataManagementActivityObservedState_ToProto(mapCtx, in.MetadataManagementActivity)
	out.NetworkConfig = NetworkConfigObservedState_ToProto(mapCtx, in.NetworkConfig)
	return out
}
func MetastoreServiceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Service) *krm.MetastoreServiceSpec {
	if in == nil {
		return nil
	}
	out := &krm.MetastoreServiceSpec{}
	out.HiveMetastoreConfig = HiveMetastoreConfig_FromProto(mapCtx, in.GetHiveMetastoreConfig())
	// MISSING: Name
	out.Labels = in.Labels
	if in.GetNetwork() != "" {
		out.NetworkRef = &refsv1beta1.ComputeNetworkRef{External: in.GetNetwork()}
	}
	out.Port = direct.LazyPtr(in.GetPort())
	out.Tier = direct.Enum_FromProto(mapCtx, in.GetTier())
	out.MaintenanceWindow = MaintenanceWindow_FromProto(mapCtx, in.GetMaintenanceWindow())
	out.ReleaseChannel = direct.Enum_FromProto(mapCtx, in.GetReleaseChannel())
	out.EncryptionConfig = EncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	out.NetworkConfig = NetworkConfig_FromProto(mapCtx, in.GetNetworkConfig())
	out.DatabaseType = direct.Enum_FromProto(mapCtx, in.GetDatabaseType())
	out.TelemetryConfig = TelemetryConfig_FromProto(mapCtx, in.GetTelemetryConfig())
	out.ScalingConfig = ScalingConfig_FromProto(mapCtx, in.GetScalingConfig())
	return out
}
func MetastoreServiceSpec_ToProto(mapCtx *direct.MapContext, in *krm.MetastoreServiceSpec) *pb.Service {
	if in == nil {
		return nil
	}
	out := &pb.Service{}
	if oneof := HiveMetastoreConfig_ToProto(mapCtx, in.HiveMetastoreConfig); oneof != nil {
		out.MetastoreConfig = &pb.Service_HiveMetastoreConfig{HiveMetastoreConfig: oneof}
	}
	// MISSING: Name
	out.Labels = in.Labels
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	out.Port = direct.ValueOf(in.Port)
	out.Tier = direct.Enum_ToProto[pb.Service_Tier](mapCtx, in.Tier)
	out.MaintenanceWindow = MaintenanceWindow_ToProto(mapCtx, in.MaintenanceWindow)
	out.ReleaseChannel = direct.Enum_ToProto[pb.Service_ReleaseChannel](mapCtx, in.ReleaseChannel)
	out.EncryptionConfig = EncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	out.NetworkConfig = NetworkConfig_ToProto(mapCtx, in.NetworkConfig)
	out.DatabaseType = direct.Enum_ToProto[pb.Service_DatabaseType](mapCtx, in.DatabaseType)
	out.TelemetryConfig = TelemetryConfig_ToProto(mapCtx, in.TelemetryConfig)
	out.ScalingConfig = ScalingConfig_ToProto(mapCtx, in.ScalingConfig)
	return out
}
func NetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig) *krm.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConfig{}
	out.Consumers = direct.Slice_FromProto(mapCtx, in.Consumers, NetworkConfig_Consumer_FromProto)
	return out
}
func NetworkConfig_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConfig) *pb.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig{}
	out.Consumers = direct.Slice_ToProto(mapCtx, in.Consumers, NetworkConfig_Consumer_ToProto)
	return out
}
func NetworkConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig) *krm.NetworkConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConfigObservedState{}
	out.Consumers = direct.Slice_FromProto(mapCtx, in.Consumers, NetworkConfig_ConsumerObservedState_FromProto)
	return out
}
func NetworkConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConfigObservedState) *pb.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig{}
	out.Consumers = direct.Slice_ToProto(mapCtx, in.Consumers, NetworkConfig_ConsumerObservedState_ToProto)
	return out
}
func NetworkConfig_Consumer_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig_Consumer) *krm.NetworkConfig_Consumer {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConfig_Consumer{}
	out.SubnetworkRef = &refsv1beta1.ComputeSubnetworkRef{External: in.GetSubnetwork()}
	// MISSING: EndpointURI
	// MISSING: EndpointLocation
	return out
}
func NetworkConfig_Consumer_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConfig_Consumer) *pb.NetworkConfig_Consumer {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig_Consumer{}
	if in.SubnetworkRef != nil {
		out.VpcResource = &pb.NetworkConfig_Consumer_Subnetwork{
			Subnetwork: in.SubnetworkRef.External,
		}
	}
	// MISSING: EndpointURI
	// MISSING: EndpointLocation
	return out
}
func NetworkConfig_ConsumerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig_Consumer) *krm.NetworkConfig_ConsumerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConfig_ConsumerObservedState{}
	// MISSING: Subnetwork
	out.EndpointURI = direct.LazyPtr(in.GetEndpointUri())
	out.EndpointLocation = direct.LazyPtr(in.GetEndpointLocation())
	return out
}
func NetworkConfig_ConsumerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConfig_ConsumerObservedState) *pb.NetworkConfig_Consumer {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig_Consumer{}
	// MISSING: Subnetwork
	out.EndpointUri = direct.ValueOf(in.EndpointURI)
	out.EndpointLocation = direct.ValueOf(in.EndpointLocation)
	return out
}
func Restore_FromProto(mapCtx *direct.MapContext, in *pb.Restore) *krm.Restore {
	if in == nil {
		return nil
	}
	out := &krm.Restore{}
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	// MISSING: Backup
	// MISSING: Type
	// MISSING: Details
	return out
}
func Restore_ToProto(mapCtx *direct.MapContext, in *krm.Restore) *pb.Restore {
	if in == nil {
		return nil
	}
	out := &pb.Restore{}
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	// MISSING: Backup
	// MISSING: Type
	// MISSING: Details
	return out
}
func RestoreObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Restore) *krm.RestoreObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RestoreObservedState{}
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Backup = direct.LazyPtr(in.GetBackup())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Details = direct.LazyPtr(in.GetDetails())
	return out
}
func RestoreObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RestoreObservedState) *pb.Restore {
	if in == nil {
		return nil
	}
	out := &pb.Restore{}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.State = direct.Enum_ToProto[pb.Restore_State](mapCtx, in.State)
	out.Backup = direct.ValueOf(in.Backup)
	out.Type = direct.Enum_ToProto[pb.Restore_RestoreType](mapCtx, in.Type)
	out.Details = direct.ValueOf(in.Details)
	return out
}
func ScalingConfig_FromProto(mapCtx *direct.MapContext, in *pb.ScalingConfig) *krm.ScalingConfig {
	if in == nil {
		return nil
	}
	out := &krm.ScalingConfig{}
	out.InstanceSize = direct.Enum_FromProto(mapCtx, in.GetInstanceSize())
	out.ScalingFactor = direct.LazyPtr(in.GetScalingFactor())
	return out
}
func ScalingConfig_ToProto(mapCtx *direct.MapContext, in *krm.ScalingConfig) *pb.ScalingConfig {
	if in == nil {
		return nil
	}
	out := &pb.ScalingConfig{}
	if in.InstanceSize != nil {
		instanceSize := direct.Enum_ToProto[pb.ScalingConfig_InstanceSize](mapCtx, in.InstanceSize)
		out.ScalingModel = &pb.ScalingConfig_InstanceSize_{
			InstanceSize: instanceSize,
		}
	}
	if in.ScalingFactor != nil {
		out.ScalingModel = &pb.ScalingConfig_ScalingFactor{
			ScalingFactor: *in.ScalingFactor,
		}
	}
	return out
}
func Secret_FromProto(mapCtx *direct.MapContext, in *pb.Secret) *krm.Secret {
	if in == nil {
		return nil
	}
	out := &krm.Secret{}
	out.SecretRef = &secretmanagerv1beta1.SecretRef{External: in.GetCloudSecret()}
	return out
}
func Secret_ToProto(mapCtx *direct.MapContext, in *krm.Secret) *pb.Secret {
	if in == nil {
		return nil
	}
	out := &pb.Secret{}
	if in.SecretRef != nil && in.SecretRef.External != "" {
		out.Value = &pb.Secret_CloudSecret{CloudSecret: in.SecretRef.External}
	}
	return out
}
func TelemetryConfig_FromProto(mapCtx *direct.MapContext, in *pb.TelemetryConfig) *krm.TelemetryConfig {
	if in == nil {
		return nil
	}
	out := &krm.TelemetryConfig{}
	out.LogFormat = direct.Enum_FromProto(mapCtx, in.GetLogFormat())
	return out
}
func TelemetryConfig_ToProto(mapCtx *direct.MapContext, in *krm.TelemetryConfig) *pb.TelemetryConfig {
	if in == nil {
		return nil
	}
	out := &pb.TelemetryConfig{}
	out.LogFormat = direct.Enum_ToProto[pb.TelemetryConfig_LogFormat](mapCtx, in.LogFormat)
	return out
}
