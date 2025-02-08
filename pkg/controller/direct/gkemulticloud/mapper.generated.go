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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/gkemulticloud/apiv1/gkemulticloudpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/containerattached/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkemulticloud/v1alpha1"
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
func AzureAuthorization_FromProto(mapCtx *direct.MapContext, in *pb.AzureAuthorization) *krm.AzureAuthorization {
	if in == nil {
		return nil
	}
	out := &krm.AzureAuthorization{}
	out.AdminUsers = direct.Slice_FromProto(mapCtx, in.AdminUsers, AzureClusterUser_FromProto)
	out.AdminGroups = direct.Slice_FromProto(mapCtx, in.AdminGroups, AzureClusterGroup_FromProto)
	return out
}
func AzureAuthorization_ToProto(mapCtx *direct.MapContext, in *krm.AzureAuthorization) *pb.AzureAuthorization {
	if in == nil {
		return nil
	}
	out := &pb.AzureAuthorization{}
	out.AdminUsers = direct.Slice_ToProto(mapCtx, in.AdminUsers, AzureClusterUser_ToProto)
	out.AdminGroups = direct.Slice_ToProto(mapCtx, in.AdminGroups, AzureClusterGroup_ToProto)
	return out
}
func AzureCluster_FromProto(mapCtx *direct.MapContext, in *pb.AzureCluster) *krm.AzureCluster {
	if in == nil {
		return nil
	}
	out := &krm.AzureCluster{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.AzureRegion = direct.LazyPtr(in.GetAzureRegion())
	out.ResourceGroupID = direct.LazyPtr(in.GetResourceGroupId())
	out.AzureClient = direct.LazyPtr(in.GetAzureClient())
	out.Networking = AzureClusterNetworking_FromProto(mapCtx, in.GetNetworking())
	out.ControlPlane = AzureControlPlane_FromProto(mapCtx, in.GetControlPlane())
	out.Authorization = AzureAuthorization_FromProto(mapCtx, in.GetAuthorization())
	out.AzureServicesAuthentication = AzureServicesAuthentication_FromProto(mapCtx, in.GetAzureServicesAuthentication())
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
	// MISSING: ManagedResources
	out.LoggingConfig = LoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	// MISSING: Errors
	out.MonitoringConfig = MonitoringConfig_FromProto(mapCtx, in.GetMonitoringConfig())
	return out
}
func AzureCluster_ToProto(mapCtx *direct.MapContext, in *krm.AzureCluster) *pb.AzureCluster {
	if in == nil {
		return nil
	}
	out := &pb.AzureCluster{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.AzureRegion = direct.ValueOf(in.AzureRegion)
	out.ResourceGroupId = direct.ValueOf(in.ResourceGroupID)
	out.AzureClient = direct.ValueOf(in.AzureClient)
	out.Networking = AzureClusterNetworking_ToProto(mapCtx, in.Networking)
	out.ControlPlane = AzureControlPlane_ToProto(mapCtx, in.ControlPlane)
	out.Authorization = AzureAuthorization_ToProto(mapCtx, in.Authorization)
	out.AzureServicesAuthentication = AzureServicesAuthentication_ToProto(mapCtx, in.AzureServicesAuthentication)
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
	// MISSING: ManagedResources
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	// MISSING: Errors
	out.MonitoringConfig = MonitoringConfig_ToProto(mapCtx, in.MonitoringConfig)
	return out
}
func AzureClusterError_FromProto(mapCtx *direct.MapContext, in *pb.AzureClusterError) *krm.AzureClusterError {
	if in == nil {
		return nil
	}
	out := &krm.AzureClusterError{}
	out.Message = direct.LazyPtr(in.GetMessage())
	return out
}
func AzureClusterError_ToProto(mapCtx *direct.MapContext, in *krm.AzureClusterError) *pb.AzureClusterError {
	if in == nil {
		return nil
	}
	out := &pb.AzureClusterError{}
	out.Message = direct.ValueOf(in.Message)
	return out
}
func AzureClusterGroup_FromProto(mapCtx *direct.MapContext, in *pb.AzureClusterGroup) *krm.AzureClusterGroup {
	if in == nil {
		return nil
	}
	out := &krm.AzureClusterGroup{}
	out.Group = direct.LazyPtr(in.GetGroup())
	return out
}
func AzureClusterGroup_ToProto(mapCtx *direct.MapContext, in *krm.AzureClusterGroup) *pb.AzureClusterGroup {
	if in == nil {
		return nil
	}
	out := &pb.AzureClusterGroup{}
	out.Group = direct.ValueOf(in.Group)
	return out
}
func AzureClusterNetworking_FromProto(mapCtx *direct.MapContext, in *pb.AzureClusterNetworking) *krm.AzureClusterNetworking {
	if in == nil {
		return nil
	}
	out := &krm.AzureClusterNetworking{}
	out.VirtualNetworkID = direct.LazyPtr(in.GetVirtualNetworkId())
	out.PodAddressCidrBlocks = in.PodAddressCidrBlocks
	out.ServiceAddressCidrBlocks = in.ServiceAddressCidrBlocks
	out.ServiceLoadBalancerSubnetID = direct.LazyPtr(in.GetServiceLoadBalancerSubnetId())
	return out
}
func AzureClusterNetworking_ToProto(mapCtx *direct.MapContext, in *krm.AzureClusterNetworking) *pb.AzureClusterNetworking {
	if in == nil {
		return nil
	}
	out := &pb.AzureClusterNetworking{}
	out.VirtualNetworkId = direct.ValueOf(in.VirtualNetworkID)
	out.PodAddressCidrBlocks = in.PodAddressCidrBlocks
	out.ServiceAddressCidrBlocks = in.ServiceAddressCidrBlocks
	out.ServiceLoadBalancerSubnetId = direct.ValueOf(in.ServiceLoadBalancerSubnetID)
	return out
}
func AzureClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AzureCluster) *krm.AzureClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AzureClusterObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: AzureRegion
	// MISSING: ResourceGroupID
	// MISSING: AzureClient
	// MISSING: Networking
	// MISSING: ControlPlane
	// MISSING: Authorization
	// MISSING: AzureServicesAuthentication
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
	out.ManagedResources = AzureClusterResources_FromProto(mapCtx, in.GetManagedResources())
	// MISSING: LoggingConfig
	out.Errors = direct.Slice_FromProto(mapCtx, in.Errors, AzureClusterError_FromProto)
	// MISSING: MonitoringConfig
	return out
}
func AzureClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AzureClusterObservedState) *pb.AzureCluster {
	if in == nil {
		return nil
	}
	out := &pb.AzureCluster{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: AzureRegion
	// MISSING: ResourceGroupID
	// MISSING: AzureClient
	// MISSING: Networking
	// MISSING: ControlPlane
	// MISSING: Authorization
	// MISSING: AzureServicesAuthentication
	out.State = direct.Enum_ToProto[pb.AzureCluster_State](mapCtx, in.State)
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
	out.ManagedResources = AzureClusterResources_ToProto(mapCtx, in.ManagedResources)
	// MISSING: LoggingConfig
	out.Errors = direct.Slice_ToProto(mapCtx, in.Errors, AzureClusterError_ToProto)
	// MISSING: MonitoringConfig
	return out
}
func AzureClusterResources_FromProto(mapCtx *direct.MapContext, in *pb.AzureClusterResources) *krm.AzureClusterResources {
	if in == nil {
		return nil
	}
	out := &krm.AzureClusterResources{}
	// MISSING: NetworkSecurityGroupID
	// MISSING: ControlPlaneApplicationSecurityGroupID
	return out
}
func AzureClusterResources_ToProto(mapCtx *direct.MapContext, in *krm.AzureClusterResources) *pb.AzureClusterResources {
	if in == nil {
		return nil
	}
	out := &pb.AzureClusterResources{}
	// MISSING: NetworkSecurityGroupID
	// MISSING: ControlPlaneApplicationSecurityGroupID
	return out
}
func AzureClusterResourcesObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AzureClusterResources) *krm.AzureClusterResourcesObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AzureClusterResourcesObservedState{}
	out.NetworkSecurityGroupID = direct.LazyPtr(in.GetNetworkSecurityGroupId())
	out.ControlPlaneApplicationSecurityGroupID = direct.LazyPtr(in.GetControlPlaneApplicationSecurityGroupId())
	return out
}
func AzureClusterResourcesObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AzureClusterResourcesObservedState) *pb.AzureClusterResources {
	if in == nil {
		return nil
	}
	out := &pb.AzureClusterResources{}
	out.NetworkSecurityGroupId = direct.ValueOf(in.NetworkSecurityGroupID)
	out.ControlPlaneApplicationSecurityGroupId = direct.ValueOf(in.ControlPlaneApplicationSecurityGroupID)
	return out
}
func AzureClusterUser_FromProto(mapCtx *direct.MapContext, in *pb.AzureClusterUser) *krm.AzureClusterUser {
	if in == nil {
		return nil
	}
	out := &krm.AzureClusterUser{}
	out.Username = direct.LazyPtr(in.GetUsername())
	return out
}
func AzureClusterUser_ToProto(mapCtx *direct.MapContext, in *krm.AzureClusterUser) *pb.AzureClusterUser {
	if in == nil {
		return nil
	}
	out := &pb.AzureClusterUser{}
	out.Username = direct.ValueOf(in.Username)
	return out
}
func AzureConfigEncryption_FromProto(mapCtx *direct.MapContext, in *pb.AzureConfigEncryption) *krm.AzureConfigEncryption {
	if in == nil {
		return nil
	}
	out := &krm.AzureConfigEncryption{}
	out.KeyID = direct.LazyPtr(in.GetKeyId())
	out.PublicKey = direct.LazyPtr(in.GetPublicKey())
	return out
}
func AzureConfigEncryption_ToProto(mapCtx *direct.MapContext, in *krm.AzureConfigEncryption) *pb.AzureConfigEncryption {
	if in == nil {
		return nil
	}
	out := &pb.AzureConfigEncryption{}
	out.KeyId = direct.ValueOf(in.KeyID)
	out.PublicKey = direct.ValueOf(in.PublicKey)
	return out
}
func AzureControlPlane_FromProto(mapCtx *direct.MapContext, in *pb.AzureControlPlane) *krm.AzureControlPlane {
	if in == nil {
		return nil
	}
	out := &krm.AzureControlPlane{}
	out.Version = direct.LazyPtr(in.GetVersion())
	out.SubnetID = direct.LazyPtr(in.GetSubnetId())
	out.VmSize = direct.LazyPtr(in.GetVmSize())
	out.SSHConfig = AzureSshConfig_FromProto(mapCtx, in.GetSshConfig())
	out.RootVolume = AzureDiskTemplate_FromProto(mapCtx, in.GetRootVolume())
	out.MainVolume = AzureDiskTemplate_FromProto(mapCtx, in.GetMainVolume())
	out.DatabaseEncryption = AzureDatabaseEncryption_FromProto(mapCtx, in.GetDatabaseEncryption())
	out.ProxyConfig = AzureProxyConfig_FromProto(mapCtx, in.GetProxyConfig())
	out.ConfigEncryption = AzureConfigEncryption_FromProto(mapCtx, in.GetConfigEncryption())
	out.Tags = in.Tags
	out.ReplicaPlacements = direct.Slice_FromProto(mapCtx, in.ReplicaPlacements, ReplicaPlacement_FromProto)
	out.EndpointSubnetID = direct.LazyPtr(in.GetEndpointSubnetId())
	return out
}
func AzureControlPlane_ToProto(mapCtx *direct.MapContext, in *krm.AzureControlPlane) *pb.AzureControlPlane {
	if in == nil {
		return nil
	}
	out := &pb.AzureControlPlane{}
	out.Version = direct.ValueOf(in.Version)
	out.SubnetId = direct.ValueOf(in.SubnetID)
	out.VmSize = direct.ValueOf(in.VmSize)
	out.SshConfig = AzureSshConfig_ToProto(mapCtx, in.SSHConfig)
	out.RootVolume = AzureDiskTemplate_ToProto(mapCtx, in.RootVolume)
	out.MainVolume = AzureDiskTemplate_ToProto(mapCtx, in.MainVolume)
	out.DatabaseEncryption = AzureDatabaseEncryption_ToProto(mapCtx, in.DatabaseEncryption)
	out.ProxyConfig = AzureProxyConfig_ToProto(mapCtx, in.ProxyConfig)
	out.ConfigEncryption = AzureConfigEncryption_ToProto(mapCtx, in.ConfigEncryption)
	out.Tags = in.Tags
	out.ReplicaPlacements = direct.Slice_ToProto(mapCtx, in.ReplicaPlacements, ReplicaPlacement_ToProto)
	out.EndpointSubnetId = direct.ValueOf(in.EndpointSubnetID)
	return out
}
func AzureDatabaseEncryption_FromProto(mapCtx *direct.MapContext, in *pb.AzureDatabaseEncryption) *krm.AzureDatabaseEncryption {
	if in == nil {
		return nil
	}
	out := &krm.AzureDatabaseEncryption{}
	out.KeyID = direct.LazyPtr(in.GetKeyId())
	return out
}
func AzureDatabaseEncryption_ToProto(mapCtx *direct.MapContext, in *krm.AzureDatabaseEncryption) *pb.AzureDatabaseEncryption {
	if in == nil {
		return nil
	}
	out := &pb.AzureDatabaseEncryption{}
	out.KeyId = direct.ValueOf(in.KeyID)
	return out
}
func AzureDiskTemplate_FromProto(mapCtx *direct.MapContext, in *pb.AzureDiskTemplate) *krm.AzureDiskTemplate {
	if in == nil {
		return nil
	}
	out := &krm.AzureDiskTemplate{}
	out.SizeGib = direct.LazyPtr(in.GetSizeGib())
	return out
}
func AzureDiskTemplate_ToProto(mapCtx *direct.MapContext, in *krm.AzureDiskTemplate) *pb.AzureDiskTemplate {
	if in == nil {
		return nil
	}
	out := &pb.AzureDiskTemplate{}
	out.SizeGib = direct.ValueOf(in.SizeGib)
	return out
}
func AzureProxyConfig_FromProto(mapCtx *direct.MapContext, in *pb.AzureProxyConfig) *krm.AzureProxyConfig {
	if in == nil {
		return nil
	}
	out := &krm.AzureProxyConfig{}
	out.ResourceGroupID = direct.LazyPtr(in.GetResourceGroupId())
	out.SecretID = direct.LazyPtr(in.GetSecretId())
	return out
}
func AzureProxyConfig_ToProto(mapCtx *direct.MapContext, in *krm.AzureProxyConfig) *pb.AzureProxyConfig {
	if in == nil {
		return nil
	}
	out := &pb.AzureProxyConfig{}
	out.ResourceGroupId = direct.ValueOf(in.ResourceGroupID)
	out.SecretId = direct.ValueOf(in.SecretID)
	return out
}
func AzureServicesAuthentication_FromProto(mapCtx *direct.MapContext, in *pb.AzureServicesAuthentication) *krm.AzureServicesAuthentication {
	if in == nil {
		return nil
	}
	out := &krm.AzureServicesAuthentication{}
	out.TenantID = direct.LazyPtr(in.GetTenantId())
	out.ApplicationID = direct.LazyPtr(in.GetApplicationId())
	return out
}
func AzureServicesAuthentication_ToProto(mapCtx *direct.MapContext, in *krm.AzureServicesAuthentication) *pb.AzureServicesAuthentication {
	if in == nil {
		return nil
	}
	out := &pb.AzureServicesAuthentication{}
	out.TenantId = direct.ValueOf(in.TenantID)
	out.ApplicationId = direct.ValueOf(in.ApplicationID)
	return out
}
func AzureSshConfig_FromProto(mapCtx *direct.MapContext, in *pb.AzureSshConfig) *krm.AzureSshConfig {
	if in == nil {
		return nil
	}
	out := &krm.AzureSshConfig{}
	out.AuthorizedKey = direct.LazyPtr(in.GetAuthorizedKey())
	return out
}
func AzureSshConfig_ToProto(mapCtx *direct.MapContext, in *krm.AzureSshConfig) *pb.AzureSshConfig {
	if in == nil {
		return nil
	}
	out := &pb.AzureSshConfig{}
	out.AuthorizedKey = direct.ValueOf(in.AuthorizedKey)
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
func GkemulticloudAzureClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AzureCluster) *krm.GkemulticloudAzureClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GkemulticloudAzureClusterObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: AzureRegion
	// MISSING: ResourceGroupID
	// MISSING: AzureClient
	// MISSING: Networking
	// MISSING: ControlPlane
	// MISSING: Authorization
	// MISSING: AzureServicesAuthentication
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
	// MISSING: ManagedResources
	// MISSING: LoggingConfig
	// MISSING: Errors
	// MISSING: MonitoringConfig
	return out
}
func GkemulticloudAzureClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GkemulticloudAzureClusterObservedState) *pb.AzureCluster {
	if in == nil {
		return nil
	}
	out := &pb.AzureCluster{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: AzureRegion
	// MISSING: ResourceGroupID
	// MISSING: AzureClient
	// MISSING: Networking
	// MISSING: ControlPlane
	// MISSING: Authorization
	// MISSING: AzureServicesAuthentication
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
	// MISSING: ManagedResources
	// MISSING: LoggingConfig
	// MISSING: Errors
	// MISSING: MonitoringConfig
	return out
}
func GkemulticloudAzureClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.AzureCluster) *krm.GkemulticloudAzureClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.GkemulticloudAzureClusterSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: AzureRegion
	// MISSING: ResourceGroupID
	// MISSING: AzureClient
	// MISSING: Networking
	// MISSING: ControlPlane
	// MISSING: Authorization
	// MISSING: AzureServicesAuthentication
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
	// MISSING: ManagedResources
	// MISSING: LoggingConfig
	// MISSING: Errors
	// MISSING: MonitoringConfig
	return out
}
func GkemulticloudAzureClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.GkemulticloudAzureClusterSpec) *pb.AzureCluster {
	if in == nil {
		return nil
	}
	out := &pb.AzureCluster{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: AzureRegion
	// MISSING: ResourceGroupID
	// MISSING: AzureClient
	// MISSING: Networking
	// MISSING: ControlPlane
	// MISSING: Authorization
	// MISSING: AzureServicesAuthentication
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
	// MISSING: ManagedResources
	// MISSING: LoggingConfig
	// MISSING: Errors
	// MISSING: MonitoringConfig
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
func ReplicaPlacement_FromProto(mapCtx *direct.MapContext, in *pb.ReplicaPlacement) *krm.ReplicaPlacement {
	if in == nil {
		return nil
	}
	out := &krm.ReplicaPlacement{}
	out.SubnetID = direct.LazyPtr(in.GetSubnetId())
	out.AzureAvailabilityZone = direct.LazyPtr(in.GetAzureAvailabilityZone())
	return out
}
func ReplicaPlacement_ToProto(mapCtx *direct.MapContext, in *krm.ReplicaPlacement) *pb.ReplicaPlacement {
	if in == nil {
		return nil
	}
	out := &pb.ReplicaPlacement{}
	out.SubnetId = direct.ValueOf(in.SubnetID)
	out.AzureAvailabilityZone = direct.ValueOf(in.AzureAvailabilityZone)
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
