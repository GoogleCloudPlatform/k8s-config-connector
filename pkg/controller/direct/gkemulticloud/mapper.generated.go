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

package gkemulticloud

import (
	pb "cloud.google.com/go/gkemulticloud/apiv1/gkemulticloudpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/containerattached/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkemulticloud/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func AttachedClusterError_FromProto(mapCtx *direct.MapContext, in *pb.AttachedClusterError) *krm.AttachedClusterError {
	if in == nil {
		return nil
	}
	out := &krm.AttachedClusterError{}
	out.Message = direct.LazyPtr(in.GetMessage())
	return out
}
func AttachedClusterError_ToProto(mapCtx *direct.MapContext, in *krm.AttachedClusterError) *pb.AttachedClusterError {
	if in == nil {
		return nil
	}
	out := &pb.AttachedClusterError{}
	out.Message = direct.ValueOf(in.Message)
	return out
}
func AttachedClustersAuthorization_FromProto(mapCtx *direct.MapContext, in *pb.AttachedClustersAuthorization) *krm.AttachedClustersAuthorization {
	if in == nil {
		return nil
	}
	out := &krm.AttachedClustersAuthorization{}
	out.AdminUsers = direct.Slice_FromProto(mapCtx, in.AdminUsers, string_FromProto)
	// MISSING: AdminGroups
	return out
}
func AttachedClustersAuthorization_ToProto(mapCtx *direct.MapContext, in *krm.AttachedClustersAuthorization) *pb.AttachedClustersAuthorization {
	if in == nil {
		return nil
	}
	out := &pb.AttachedClustersAuthorization{}
	out.AdminUsers = direct.Slice_ToProto(mapCtx, in.AdminUsers, string_ToProto)
	// MISSING: AdminGroups
	return out
}
func AttachedOidcConfig_FromProto(mapCtx *direct.MapContext, in *pb.AttachedOidcConfig) *krm.AttachedOidcConfig {
	if in == nil {
		return nil
	}
	out := &krm.AttachedOidcConfig{}
	out.IssuerURL = direct.LazyPtr(in.GetIssuerUrl())
	out.Jwks = in.GetJwks()
	return out
}
func AttachedOidcConfig_ToProto(mapCtx *direct.MapContext, in *krm.AttachedOidcConfig) *pb.AttachedOidcConfig {
	if in == nil {
		return nil
	}
	out := &pb.AttachedOidcConfig{}
	out.IssuerURL = AttachedOidcConfig_IssuerUrl_ToProto(mapCtx, in.IssuerURL)
	out.Jwks = in.Jwks
	return out
}
func AttachedProxyConfig_FromProto(mapCtx *direct.MapContext, in *pb.AttachedProxyConfig) *krm.AttachedProxyConfig {
	if in == nil {
		return nil
	}
	out := &krm.AttachedProxyConfig{}
	out.KubernetesSecret = KubernetesSecret_FromProto(mapCtx, in.GetKubernetesSecret())
	return out
}
func AttachedProxyConfig_ToProto(mapCtx *direct.MapContext, in *krm.AttachedProxyConfig) *pb.AttachedProxyConfig {
	if in == nil {
		return nil
	}
	out := &pb.AttachedProxyConfig{}
	out.KubernetesSecret = KubernetesSecret_ToProto(mapCtx, in.KubernetesSecret)
	return out
}
func AwsAuthorization_FromProto(mapCtx *direct.MapContext, in *pb.AwsAuthorization) *krm.AwsAuthorization {
	if in == nil {
		return nil
	}
	out := &krm.AwsAuthorization{}
	out.AdminUsers = direct.Slice_FromProto(mapCtx, in.AdminUsers, AwsClusterUser_FromProto)
	out.AdminGroups = direct.Slice_FromProto(mapCtx, in.AdminGroups, AwsClusterGroup_FromProto)
	return out
}
func AwsAuthorization_ToProto(mapCtx *direct.MapContext, in *krm.AwsAuthorization) *pb.AwsAuthorization {
	if in == nil {
		return nil
	}
	out := &pb.AwsAuthorization{}
	out.AdminUsers = direct.Slice_ToProto(mapCtx, in.AdminUsers, AwsClusterUser_ToProto)
	out.AdminGroups = direct.Slice_ToProto(mapCtx, in.AdminGroups, AwsClusterGroup_ToProto)
	return out
}
func AwsCluster_FromProto(mapCtx *direct.MapContext, in *pb.AwsCluster) *krm.AwsCluster {
	if in == nil {
		return nil
	}
	out := &krm.AwsCluster{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Networking = AwsClusterNetworking_FromProto(mapCtx, in.GetNetworking())
	out.AwsRegion = direct.LazyPtr(in.GetAwsRegion())
	out.ControlPlane = AwsControlPlane_FromProto(mapCtx, in.GetControlPlane())
	out.Authorization = AwsAuthorization_FromProto(mapCtx, in.GetAuthorization())
	// MISSING: State
	// MISSING: Endpoint
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Annotations = in.Annotations
	// MISSING: WorkloadIdentityConfig
	// MISSING: ClusterCaCertificate
	out.Fleet = Fleet_FromProto(mapCtx, in.GetFleet())
	out.LoggingConfig = LoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	// MISSING: Errors
	out.MonitoringConfig = MonitoringConfig_FromProto(mapCtx, in.GetMonitoringConfig())
	out.BinaryAuthorization = BinaryAuthorization_FromProto(mapCtx, in.GetBinaryAuthorization())
	return out
}
func AwsCluster_ToProto(mapCtx *direct.MapContext, in *krm.AwsCluster) *pb.AwsCluster {
	if in == nil {
		return nil
	}
	out := &pb.AwsCluster{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.Networking = AwsClusterNetworking_ToProto(mapCtx, in.Networking)
	out.AwsRegion = direct.ValueOf(in.AwsRegion)
	out.ControlPlane = AwsControlPlane_ToProto(mapCtx, in.ControlPlane)
	out.Authorization = AwsAuthorization_ToProto(mapCtx, in.Authorization)
	// MISSING: State
	// MISSING: Endpoint
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Etag = direct.ValueOf(in.Etag)
	out.Annotations = in.Annotations
	// MISSING: WorkloadIdentityConfig
	// MISSING: ClusterCaCertificate
	out.Fleet = Fleet_ToProto(mapCtx, in.Fleet)
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	// MISSING: Errors
	out.MonitoringConfig = MonitoringConfig_ToProto(mapCtx, in.MonitoringConfig)
	out.BinaryAuthorization = BinaryAuthorization_ToProto(mapCtx, in.BinaryAuthorization)
	return out
}
func AwsClusterError_FromProto(mapCtx *direct.MapContext, in *pb.AwsClusterError) *krm.AwsClusterError {
	if in == nil {
		return nil
	}
	out := &krm.AwsClusterError{}
	out.Message = direct.LazyPtr(in.GetMessage())
	return out
}
func AwsClusterError_ToProto(mapCtx *direct.MapContext, in *krm.AwsClusterError) *pb.AwsClusterError {
	if in == nil {
		return nil
	}
	out := &pb.AwsClusterError{}
	out.Message = direct.ValueOf(in.Message)
	return out
}
func AwsClusterGroup_FromProto(mapCtx *direct.MapContext, in *pb.AwsClusterGroup) *krm.AwsClusterGroup {
	if in == nil {
		return nil
	}
	out := &krm.AwsClusterGroup{}
	out.Group = direct.LazyPtr(in.GetGroup())
	return out
}
func AwsClusterGroup_ToProto(mapCtx *direct.MapContext, in *krm.AwsClusterGroup) *pb.AwsClusterGroup {
	if in == nil {
		return nil
	}
	out := &pb.AwsClusterGroup{}
	out.Group = direct.ValueOf(in.Group)
	return out
}
func AwsClusterNetworking_FromProto(mapCtx *direct.MapContext, in *pb.AwsClusterNetworking) *krm.AwsClusterNetworking {
	if in == nil {
		return nil
	}
	out := &krm.AwsClusterNetworking{}
	out.VpcID = direct.LazyPtr(in.GetVpcId())
	out.PodAddressCidrBlocks = in.PodAddressCidrBlocks
	out.ServiceAddressCidrBlocks = in.ServiceAddressCidrBlocks
	out.PerNodePoolSgRulesDisabled = direct.LazyPtr(in.GetPerNodePoolSgRulesDisabled())
	return out
}
func AwsClusterNetworking_ToProto(mapCtx *direct.MapContext, in *krm.AwsClusterNetworking) *pb.AwsClusterNetworking {
	if in == nil {
		return nil
	}
	out := &pb.AwsClusterNetworking{}
	out.VpcId = direct.ValueOf(in.VpcID)
	out.PodAddressCidrBlocks = in.PodAddressCidrBlocks
	out.ServiceAddressCidrBlocks = in.ServiceAddressCidrBlocks
	out.PerNodePoolSgRulesDisabled = direct.ValueOf(in.PerNodePoolSgRulesDisabled)
	return out
}
func AwsClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AwsCluster) *krm.AwsClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AwsClusterObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Networking
	// MISSING: AwsRegion
	// MISSING: ControlPlane
	// MISSING: Authorization
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Endpoint = direct.LazyPtr(in.GetEndpoint())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Etag
	// MISSING: Annotations
	out.WorkloadIdentityConfig = WorkloadIdentityConfig_FromProto(mapCtx, in.GetWorkloadIdentityConfig())
	out.ClusterCaCertificate = direct.LazyPtr(in.GetClusterCaCertificate())
	out.Fleet = FleetObservedState_FromProto(mapCtx, in.GetFleet())
	// MISSING: LoggingConfig
	out.Errors = direct.Slice_FromProto(mapCtx, in.Errors, AwsClusterError_FromProto)
	// MISSING: MonitoringConfig
	// MISSING: BinaryAuthorization
	return out
}
func AwsClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AwsClusterObservedState) *pb.AwsCluster {
	if in == nil {
		return nil
	}
	out := &pb.AwsCluster{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Networking
	// MISSING: AwsRegion
	// MISSING: ControlPlane
	// MISSING: Authorization
	out.State = direct.Enum_ToProto[pb.AwsCluster_State](mapCtx, in.State)
	out.Endpoint = direct.ValueOf(in.Endpoint)
	out.Uid = direct.ValueOf(in.Uid)
	out.Reconciling = direct.ValueOf(in.Reconciling)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Etag
	// MISSING: Annotations
	out.WorkloadIdentityConfig = WorkloadIdentityConfig_ToProto(mapCtx, in.WorkloadIdentityConfig)
	out.ClusterCaCertificate = direct.ValueOf(in.ClusterCaCertificate)
	out.Fleet = FleetObservedState_ToProto(mapCtx, in.Fleet)
	// MISSING: LoggingConfig
	out.Errors = direct.Slice_ToProto(mapCtx, in.Errors, AwsClusterError_ToProto)
	// MISSING: MonitoringConfig
	// MISSING: BinaryAuthorization
	return out
}
func AwsClusterUser_FromProto(mapCtx *direct.MapContext, in *pb.AwsClusterUser) *krm.AwsClusterUser {
	if in == nil {
		return nil
	}
	out := &krm.AwsClusterUser{}
	out.Username = direct.LazyPtr(in.GetUsername())
	return out
}
func AwsClusterUser_ToProto(mapCtx *direct.MapContext, in *krm.AwsClusterUser) *pb.AwsClusterUser {
	if in == nil {
		return nil
	}
	out := &pb.AwsClusterUser{}
	out.Username = direct.ValueOf(in.Username)
	return out
}
func AwsConfigEncryption_FromProto(mapCtx *direct.MapContext, in *pb.AwsConfigEncryption) *krm.AwsConfigEncryption {
	if in == nil {
		return nil
	}
	out := &krm.AwsConfigEncryption{}
	out.KMSKeyArn = direct.LazyPtr(in.GetKmsKeyArn())
	return out
}
func AwsConfigEncryption_ToProto(mapCtx *direct.MapContext, in *krm.AwsConfigEncryption) *pb.AwsConfigEncryption {
	if in == nil {
		return nil
	}
	out := &pb.AwsConfigEncryption{}
	out.KmsKeyArn = direct.ValueOf(in.KMSKeyArn)
	return out
}
func AwsControlPlane_FromProto(mapCtx *direct.MapContext, in *pb.AwsControlPlane) *krm.AwsControlPlane {
	if in == nil {
		return nil
	}
	out := &krm.AwsControlPlane{}
	out.Version = direct.LazyPtr(in.GetVersion())
	out.InstanceType = direct.LazyPtr(in.GetInstanceType())
	out.SSHConfig = AwsSshConfig_FromProto(mapCtx, in.GetSshConfig())
	out.SubnetIds = in.SubnetIds
	out.SecurityGroupIds = in.SecurityGroupIds
	out.IamInstanceProfile = direct.LazyPtr(in.GetIamInstanceProfile())
	out.RootVolume = AwsVolumeTemplate_FromProto(mapCtx, in.GetRootVolume())
	out.MainVolume = AwsVolumeTemplate_FromProto(mapCtx, in.GetMainVolume())
	out.DatabaseEncryption = AwsDatabaseEncryption_FromProto(mapCtx, in.GetDatabaseEncryption())
	out.Tags = in.Tags
	out.AwsServicesAuthentication = AwsServicesAuthentication_FromProto(mapCtx, in.GetAwsServicesAuthentication())
	out.ProxyConfig = AwsProxyConfig_FromProto(mapCtx, in.GetProxyConfig())
	out.ConfigEncryption = AwsConfigEncryption_FromProto(mapCtx, in.GetConfigEncryption())
	out.InstancePlacement = AwsInstancePlacement_FromProto(mapCtx, in.GetInstancePlacement())
	return out
}
func AwsControlPlane_ToProto(mapCtx *direct.MapContext, in *krm.AwsControlPlane) *pb.AwsControlPlane {
	if in == nil {
		return nil
	}
	out := &pb.AwsControlPlane{}
	out.Version = direct.ValueOf(in.Version)
	out.InstanceType = direct.ValueOf(in.InstanceType)
	out.SshConfig = AwsSshConfig_ToProto(mapCtx, in.SSHConfig)
	out.SubnetIds = in.SubnetIds
	out.SecurityGroupIds = in.SecurityGroupIds
	out.IamInstanceProfile = direct.ValueOf(in.IamInstanceProfile)
	out.RootVolume = AwsVolumeTemplate_ToProto(mapCtx, in.RootVolume)
	out.MainVolume = AwsVolumeTemplate_ToProto(mapCtx, in.MainVolume)
	out.DatabaseEncryption = AwsDatabaseEncryption_ToProto(mapCtx, in.DatabaseEncryption)
	out.Tags = in.Tags
	out.AwsServicesAuthentication = AwsServicesAuthentication_ToProto(mapCtx, in.AwsServicesAuthentication)
	out.ProxyConfig = AwsProxyConfig_ToProto(mapCtx, in.ProxyConfig)
	out.ConfigEncryption = AwsConfigEncryption_ToProto(mapCtx, in.ConfigEncryption)
	out.InstancePlacement = AwsInstancePlacement_ToProto(mapCtx, in.InstancePlacement)
	return out
}
func AwsDatabaseEncryption_FromProto(mapCtx *direct.MapContext, in *pb.AwsDatabaseEncryption) *krm.AwsDatabaseEncryption {
	if in == nil {
		return nil
	}
	out := &krm.AwsDatabaseEncryption{}
	out.KMSKeyArn = direct.LazyPtr(in.GetKmsKeyArn())
	return out
}
func AwsDatabaseEncryption_ToProto(mapCtx *direct.MapContext, in *krm.AwsDatabaseEncryption) *pb.AwsDatabaseEncryption {
	if in == nil {
		return nil
	}
	out := &pb.AwsDatabaseEncryption{}
	out.KmsKeyArn = direct.ValueOf(in.KMSKeyArn)
	return out
}
func AwsInstancePlacement_FromProto(mapCtx *direct.MapContext, in *pb.AwsInstancePlacement) *krm.AwsInstancePlacement {
	if in == nil {
		return nil
	}
	out := &krm.AwsInstancePlacement{}
	out.Tenancy = direct.Enum_FromProto(mapCtx, in.GetTenancy())
	return out
}
func AwsInstancePlacement_ToProto(mapCtx *direct.MapContext, in *krm.AwsInstancePlacement) *pb.AwsInstancePlacement {
	if in == nil {
		return nil
	}
	out := &pb.AwsInstancePlacement{}
	out.Tenancy = direct.Enum_ToProto[pb.AwsInstancePlacement_Tenancy](mapCtx, in.Tenancy)
	return out
}
func AwsProxyConfig_FromProto(mapCtx *direct.MapContext, in *pb.AwsProxyConfig) *krm.AwsProxyConfig {
	if in == nil {
		return nil
	}
	out := &krm.AwsProxyConfig{}
	out.SecretArn = direct.LazyPtr(in.GetSecretArn())
	out.SecretVersion = direct.LazyPtr(in.GetSecretVersion())
	return out
}
func AwsProxyConfig_ToProto(mapCtx *direct.MapContext, in *krm.AwsProxyConfig) *pb.AwsProxyConfig {
	if in == nil {
		return nil
	}
	out := &pb.AwsProxyConfig{}
	out.SecretArn = direct.ValueOf(in.SecretArn)
	out.SecretVersion = direct.ValueOf(in.SecretVersion)
	return out
}
func AwsServicesAuthentication_FromProto(mapCtx *direct.MapContext, in *pb.AwsServicesAuthentication) *krm.AwsServicesAuthentication {
	if in == nil {
		return nil
	}
	out := &krm.AwsServicesAuthentication{}
	out.RoleArn = direct.LazyPtr(in.GetRoleArn())
	out.RoleSessionName = direct.LazyPtr(in.GetRoleSessionName())
	return out
}
func AwsServicesAuthentication_ToProto(mapCtx *direct.MapContext, in *krm.AwsServicesAuthentication) *pb.AwsServicesAuthentication {
	if in == nil {
		return nil
	}
	out := &pb.AwsServicesAuthentication{}
	out.RoleArn = direct.ValueOf(in.RoleArn)
	out.RoleSessionName = direct.ValueOf(in.RoleSessionName)
	return out
}
func AwsSshConfig_FromProto(mapCtx *direct.MapContext, in *pb.AwsSshConfig) *krm.AwsSshConfig {
	if in == nil {
		return nil
	}
	out := &krm.AwsSshConfig{}
	out.Ec2KeyPair = direct.LazyPtr(in.GetEc2KeyPair())
	return out
}
func AwsSshConfig_ToProto(mapCtx *direct.MapContext, in *krm.AwsSshConfig) *pb.AwsSshConfig {
	if in == nil {
		return nil
	}
	out := &pb.AwsSshConfig{}
	out.Ec2KeyPair = direct.ValueOf(in.Ec2KeyPair)
	return out
}
func AwsVolumeTemplate_FromProto(mapCtx *direct.MapContext, in *pb.AwsVolumeTemplate) *krm.AwsVolumeTemplate {
	if in == nil {
		return nil
	}
	out := &krm.AwsVolumeTemplate{}
	out.SizeGib = direct.LazyPtr(in.GetSizeGib())
	out.VolumeType = direct.Enum_FromProto(mapCtx, in.GetVolumeType())
	out.Iops = direct.LazyPtr(in.GetIops())
	out.Throughput = direct.LazyPtr(in.GetThroughput())
	out.KMSKeyArn = direct.LazyPtr(in.GetKmsKeyArn())
	return out
}
func AwsVolumeTemplate_ToProto(mapCtx *direct.MapContext, in *krm.AwsVolumeTemplate) *pb.AwsVolumeTemplate {
	if in == nil {
		return nil
	}
	out := &pb.AwsVolumeTemplate{}
	out.SizeGib = direct.ValueOf(in.SizeGib)
	out.VolumeType = direct.Enum_ToProto[pb.AwsVolumeTemplate_VolumeType](mapCtx, in.VolumeType)
	out.Iops = direct.ValueOf(in.Iops)
	out.Throughput = direct.ValueOf(in.Throughput)
	out.KmsKeyArn = direct.ValueOf(in.KMSKeyArn)
	return out
}
func BinaryAuthorization_FromProto(mapCtx *direct.MapContext, in *pb.BinaryAuthorization) *krm.BinaryAuthorization {
	if in == nil {
		return nil
	}
	out := &krm.BinaryAuthorization{}
	out.EvaluationMode = direct.Enum_FromProto(mapCtx, in.GetEvaluationMode())
	return out
}
func BinaryAuthorization_ToProto(mapCtx *direct.MapContext, in *krm.BinaryAuthorization) *pb.BinaryAuthorization {
	if in == nil {
		return nil
	}
	out := &pb.BinaryAuthorization{}
	out.EvaluationMode = direct.Enum_ToProto[pb.BinaryAuthorization_EvaluationMode](mapCtx, in.EvaluationMode)
	return out
}
func CloudMonitoringConfig_FromProto(mapCtx *direct.MapContext, in *pb.CloudMonitoringConfig) *krm.CloudMonitoringConfig {
	if in == nil {
		return nil
	}
	out := &krm.CloudMonitoringConfig{}
	out.Enabled = in.Enabled
	return out
}
func CloudMonitoringConfig_ToProto(mapCtx *direct.MapContext, in *krm.CloudMonitoringConfig) *pb.CloudMonitoringConfig {
	if in == nil {
		return nil
	}
	out := &pb.CloudMonitoringConfig{}
	out.Enabled = in.Enabled
	return out
}
func ContainerAttachedClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.AttachedCluster) *krm.ContainerAttachedClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.ContainerAttachedClusterSpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	out.OidcConfig = AttachedOidcConfig_FromProto(mapCtx, in.GetOidcConfig())
	out.PlatformVersion = direct.LazyPtr(in.GetPlatformVersion())
	out.Distribution = direct.LazyPtr(in.GetDistribution())
	// MISSING: ClusterRegion
	out.Fleet = Fleet_FromProto(mapCtx, in.GetFleet())
	// MISSING: State
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: KubernetesVersion
	out.Annotations = in.Annotations
	// MISSING: WorkloadIdentityConfig
	out.LoggingConfig = LoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	// MISSING: Errors
	out.Authorization = AttachedClustersAuthorization_FromProto(mapCtx, in.GetAuthorization())
	out.MonitoringConfig = MonitoringConfig_FromProto(mapCtx, in.GetMonitoringConfig())
	// MISSING: ProxyConfig
	out.BinaryAuthorization = BinaryAuthorization_FromProto(mapCtx, in.GetBinaryAuthorization())
	// MISSING: SecurityPostureConfig
	// MISSING: Tags
	return out
}
func ContainerAttachedClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.ContainerAttachedClusterSpec) *pb.AttachedCluster {
	if in == nil {
		return nil
	}
	out := &pb.AttachedCluster{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	out.OidcConfig = AttachedOidcConfig_ToProto(mapCtx, in.OidcConfig)
	out.PlatformVersion = ContainerAttachedClusterSpec_PlatformVersion_ToProto(mapCtx, in.PlatformVersion)
	out.Distribution = ContainerAttachedClusterSpec_Distribution_ToProto(mapCtx, in.Distribution)
	// MISSING: ClusterRegion
	out.Fleet = Fleet_ToProto(mapCtx, in.Fleet)
	// MISSING: State
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: KubernetesVersion
	out.Annotations = in.Annotations
	// MISSING: WorkloadIdentityConfig
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	// MISSING: Errors
	out.Authorization = AttachedClustersAuthorization_ToProto(mapCtx, in.Authorization)
	out.MonitoringConfig = MonitoringConfig_ToProto(mapCtx, in.MonitoringConfig)
	// MISSING: ProxyConfig
	out.BinaryAuthorization = BinaryAuthorization_ToProto(mapCtx, in.BinaryAuthorization)
	// MISSING: SecurityPostureConfig
	// MISSING: Tags
	return out
}
func Fleet_FromProto(mapCtx *direct.MapContext, in *pb.Fleet) *krm.Fleet {
	if in == nil {
		return nil
	}
	out := &krm.Fleet{}
	out.Project = direct.LazyPtr(in.GetProject())
	// MISSING: Membership
	return out
}
func Fleet_ToProto(mapCtx *direct.MapContext, in *krm.Fleet) *pb.Fleet {
	if in == nil {
		return nil
	}
	out := &pb.Fleet{}
	out.Project = direct.ValueOf(in.Project)
	// MISSING: Membership
	return out
}
func FleetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Fleet) *krm.FleetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FleetObservedState{}
	// MISSING: Project
	out.Membership = direct.LazyPtr(in.GetMembership())
	return out
}
func FleetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FleetObservedState) *pb.Fleet {
	if in == nil {
		return nil
	}
	out := &pb.Fleet{}
	// MISSING: Project
	out.Membership = direct.ValueOf(in.Membership)
	return out
}
func GkemulticloudAwsClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AwsCluster) *krm.GkemulticloudAwsClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GkemulticloudAwsClusterObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Networking
	// MISSING: AwsRegion
	// MISSING: ControlPlane
	// MISSING: Authorization
	// MISSING: State
	// MISSING: Endpoint
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: WorkloadIdentityConfig
	// MISSING: ClusterCaCertificate
	// MISSING: Fleet
	// MISSING: LoggingConfig
	// MISSING: Errors
	// MISSING: MonitoringConfig
	// MISSING: BinaryAuthorization
	return out
}
func GkemulticloudAwsClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GkemulticloudAwsClusterObservedState) *pb.AwsCluster {
	if in == nil {
		return nil
	}
	out := &pb.AwsCluster{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Networking
	// MISSING: AwsRegion
	// MISSING: ControlPlane
	// MISSING: Authorization
	// MISSING: State
	// MISSING: Endpoint
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: WorkloadIdentityConfig
	// MISSING: ClusterCaCertificate
	// MISSING: Fleet
	// MISSING: LoggingConfig
	// MISSING: Errors
	// MISSING: MonitoringConfig
	// MISSING: BinaryAuthorization
	return out
}
func GkemulticloudAwsClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.AwsCluster) *krm.GkemulticloudAwsClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.GkemulticloudAwsClusterSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Networking
	// MISSING: AwsRegion
	// MISSING: ControlPlane
	// MISSING: Authorization
	// MISSING: State
	// MISSING: Endpoint
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: WorkloadIdentityConfig
	// MISSING: ClusterCaCertificate
	// MISSING: Fleet
	// MISSING: LoggingConfig
	// MISSING: Errors
	// MISSING: MonitoringConfig
	// MISSING: BinaryAuthorization
	return out
}
func GkemulticloudAwsClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.GkemulticloudAwsClusterSpec) *pb.AwsCluster {
	if in == nil {
		return nil
	}
	out := &pb.AwsCluster{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Networking
	// MISSING: AwsRegion
	// MISSING: ControlPlane
	// MISSING: Authorization
	// MISSING: State
	// MISSING: Endpoint
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: WorkloadIdentityConfig
	// MISSING: ClusterCaCertificate
	// MISSING: Fleet
	// MISSING: LoggingConfig
	// MISSING: Errors
	// MISSING: MonitoringConfig
	// MISSING: BinaryAuthorization
	return out
}
func KubernetesSecret_FromProto(mapCtx *direct.MapContext, in *pb.KubernetesSecret) *krm.KubernetesSecret {
	if in == nil {
		return nil
	}
	out := &krm.KubernetesSecret{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Namespace = direct.LazyPtr(in.GetNamespace())
	return out
}
func KubernetesSecret_ToProto(mapCtx *direct.MapContext, in *krm.KubernetesSecret) *pb.KubernetesSecret {
	if in == nil {
		return nil
	}
	out := &pb.KubernetesSecret{}
	out.Name = KubernetesSecret_Name_ToProto(mapCtx, in.Name)
	out.Namespace = KubernetesSecret_Namespace_ToProto(mapCtx, in.Namespace)
	return out
}
func LoggingComponentConfig_FromProto(mapCtx *direct.MapContext, in *pb.LoggingComponentConfig) *krm.LoggingComponentConfig {
	if in == nil {
		return nil
	}
	out := &krm.LoggingComponentConfig{}
	out.EnableComponents = direct.EnumSlice_FromProto(mapCtx, in.EnableComponents)
	return out
}
func LoggingComponentConfig_ToProto(mapCtx *direct.MapContext, in *krm.LoggingComponentConfig) *pb.LoggingComponentConfig {
	if in == nil {
		return nil
	}
	out := &pb.LoggingComponentConfig{}
	out.EnableComponents = direct.EnumSlice_ToProto[pb.LoggingComponentConfig_Component](mapCtx, in.EnableComponents)
	return out
}
func LoggingConfig_FromProto(mapCtx *direct.MapContext, in *pb.LoggingConfig) *krm.LoggingConfig {
	if in == nil {
		return nil
	}
	out := &krm.LoggingConfig{}
	out.ComponentConfig = LoggingComponentConfig_FromProto(mapCtx, in.GetComponentConfig())
	return out
}
func LoggingConfig_ToProto(mapCtx *direct.MapContext, in *krm.LoggingConfig) *pb.LoggingConfig {
	if in == nil {
		return nil
	}
	out := &pb.LoggingConfig{}
	out.ComponentConfig = LoggingComponentConfig_ToProto(mapCtx, in.ComponentConfig)
	return out
}
func ManagedPrometheusConfig_FromProto(mapCtx *direct.MapContext, in *pb.ManagedPrometheusConfig) *krm.ManagedPrometheusConfig {
	if in == nil {
		return nil
	}
	out := &krm.ManagedPrometheusConfig{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	return out
}
func ManagedPrometheusConfig_ToProto(mapCtx *direct.MapContext, in *krm.ManagedPrometheusConfig) *pb.ManagedPrometheusConfig {
	if in == nil {
		return nil
	}
	out := &pb.ManagedPrometheusConfig{}
	out.Enabled = direct.ValueOf(in.Enabled)
	return out
}
func MonitoringConfig_FromProto(mapCtx *direct.MapContext, in *pb.MonitoringConfig) *krm.MonitoringConfig {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringConfig{}
	out.ManagedPrometheusConfig = ManagedPrometheusConfig_FromProto(mapCtx, in.GetManagedPrometheusConfig())
	out.CloudMonitoringConfig = CloudMonitoringConfig_FromProto(mapCtx, in.GetCloudMonitoringConfig())
	return out
}
func MonitoringConfig_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringConfig) *pb.MonitoringConfig {
	if in == nil {
		return nil
	}
	out := &pb.MonitoringConfig{}
	out.ManagedPrometheusConfig = ManagedPrometheusConfig_ToProto(mapCtx, in.ManagedPrometheusConfig)
	out.CloudMonitoringConfig = CloudMonitoringConfig_ToProto(mapCtx, in.CloudMonitoringConfig)
	return out
}
func WorkloadIdentityConfig_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadIdentityConfig) *krm.WorkloadIdentityConfig {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadIdentityConfig{}
	out.IssuerURI = direct.LazyPtr(in.GetIssuerUri())
	out.WorkloadPool = direct.LazyPtr(in.GetWorkloadPool())
	out.IdentityProvider = direct.LazyPtr(in.GetIdentityProvider())
	return out
}
func WorkloadIdentityConfig_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadIdentityConfig) *pb.WorkloadIdentityConfig {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadIdentityConfig{}
	out.IssuerUri = direct.ValueOf(in.IssuerURI)
	out.WorkloadPool = direct.ValueOf(in.WorkloadPool)
	out.IdentityProvider = direct.ValueOf(in.IdentityProvider)
	return out
}
