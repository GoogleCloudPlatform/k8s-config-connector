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
func AzureNodeConfig_FromProto(mapCtx *direct.MapContext, in *pb.AzureNodeConfig) *krm.AzureNodeConfig {
	if in == nil {
		return nil
	}
	out := &krm.AzureNodeConfig{}
	out.VmSize = direct.LazyPtr(in.GetVmSize())
	out.RootVolume = AzureDiskTemplate_FromProto(mapCtx, in.GetRootVolume())
	out.Tags = in.Tags
	out.ImageType = direct.LazyPtr(in.GetImageType())
	out.SSHConfig = AzureSshConfig_FromProto(mapCtx, in.GetSshConfig())
	out.ProxyConfig = AzureProxyConfig_FromProto(mapCtx, in.GetProxyConfig())
	out.ConfigEncryption = AzureConfigEncryption_FromProto(mapCtx, in.GetConfigEncryption())
	out.Taints = direct.Slice_FromProto(mapCtx, in.Taints, NodeTaint_FromProto)
	out.Labels = in.Labels
	return out
}
func AzureNodeConfig_ToProto(mapCtx *direct.MapContext, in *krm.AzureNodeConfig) *pb.AzureNodeConfig {
	if in == nil {
		return nil
	}
	out := &pb.AzureNodeConfig{}
	out.VmSize = direct.ValueOf(in.VmSize)
	out.RootVolume = AzureDiskTemplate_ToProto(mapCtx, in.RootVolume)
	out.Tags = in.Tags
	out.ImageType = direct.ValueOf(in.ImageType)
	out.SshConfig = AzureSshConfig_ToProto(mapCtx, in.SSHConfig)
	out.ProxyConfig = AzureProxyConfig_ToProto(mapCtx, in.ProxyConfig)
	out.ConfigEncryption = AzureConfigEncryption_ToProto(mapCtx, in.ConfigEncryption)
	out.Taints = direct.Slice_ToProto(mapCtx, in.Taints, NodeTaint_ToProto)
	out.Labels = in.Labels
	return out
}
func AzureNodeManagement_FromProto(mapCtx *direct.MapContext, in *pb.AzureNodeManagement) *krm.AzureNodeManagement {
	if in == nil {
		return nil
	}
	out := &krm.AzureNodeManagement{}
	out.AutoRepair = direct.LazyPtr(in.GetAutoRepair())
	return out
}
func AzureNodeManagement_ToProto(mapCtx *direct.MapContext, in *krm.AzureNodeManagement) *pb.AzureNodeManagement {
	if in == nil {
		return nil
	}
	out := &pb.AzureNodeManagement{}
	out.AutoRepair = direct.ValueOf(in.AutoRepair)
	return out
}
func AzureNodePool_FromProto(mapCtx *direct.MapContext, in *pb.AzureNodePool) *krm.AzureNodePool {
	if in == nil {
		return nil
	}
	out := &krm.AzureNodePool{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Version = direct.LazyPtr(in.GetVersion())
	out.Config = AzureNodeConfig_FromProto(mapCtx, in.GetConfig())
	out.SubnetID = direct.LazyPtr(in.GetSubnetId())
	out.Autoscaling = AzureNodePoolAutoscaling_FromProto(mapCtx, in.GetAutoscaling())
	// MISSING: State
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Annotations = in.Annotations
	out.MaxPodsConstraint = MaxPodsConstraint_FromProto(mapCtx, in.GetMaxPodsConstraint())
	out.AzureAvailabilityZone = direct.LazyPtr(in.GetAzureAvailabilityZone())
	// MISSING: Errors
	out.Management = AzureNodeManagement_FromProto(mapCtx, in.GetManagement())
	return out
}
func AzureNodePool_ToProto(mapCtx *direct.MapContext, in *krm.AzureNodePool) *pb.AzureNodePool {
	if in == nil {
		return nil
	}
	out := &pb.AzureNodePool{}
	out.Name = direct.ValueOf(in.Name)
	out.Version = direct.ValueOf(in.Version)
	out.Config = AzureNodeConfig_ToProto(mapCtx, in.Config)
	out.SubnetId = direct.ValueOf(in.SubnetID)
	out.Autoscaling = AzureNodePoolAutoscaling_ToProto(mapCtx, in.Autoscaling)
	// MISSING: State
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Etag = direct.ValueOf(in.Etag)
	out.Annotations = in.Annotations
	out.MaxPodsConstraint = MaxPodsConstraint_ToProto(mapCtx, in.MaxPodsConstraint)
	out.AzureAvailabilityZone = direct.ValueOf(in.AzureAvailabilityZone)
	// MISSING: Errors
	out.Management = AzureNodeManagement_ToProto(mapCtx, in.Management)
	return out
}
func AzureNodePoolAutoscaling_FromProto(mapCtx *direct.MapContext, in *pb.AzureNodePoolAutoscaling) *krm.AzureNodePoolAutoscaling {
	if in == nil {
		return nil
	}
	out := &krm.AzureNodePoolAutoscaling{}
	out.MinNodeCount = direct.LazyPtr(in.GetMinNodeCount())
	out.MaxNodeCount = direct.LazyPtr(in.GetMaxNodeCount())
	return out
}
func AzureNodePoolAutoscaling_ToProto(mapCtx *direct.MapContext, in *krm.AzureNodePoolAutoscaling) *pb.AzureNodePoolAutoscaling {
	if in == nil {
		return nil
	}
	out := &pb.AzureNodePoolAutoscaling{}
	out.MinNodeCount = direct.ValueOf(in.MinNodeCount)
	out.MaxNodeCount = direct.ValueOf(in.MaxNodeCount)
	return out
}
func AzureNodePoolError_FromProto(mapCtx *direct.MapContext, in *pb.AzureNodePoolError) *krm.AzureNodePoolError {
	if in == nil {
		return nil
	}
	out := &krm.AzureNodePoolError{}
	out.Message = direct.LazyPtr(in.GetMessage())
	return out
}
func AzureNodePoolError_ToProto(mapCtx *direct.MapContext, in *krm.AzureNodePoolError) *pb.AzureNodePoolError {
	if in == nil {
		return nil
	}
	out := &pb.AzureNodePoolError{}
	out.Message = direct.ValueOf(in.Message)
	return out
}
func AzureNodePoolObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AzureNodePool) *krm.AzureNodePoolObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AzureNodePoolObservedState{}
	// MISSING: Name
	// MISSING: Version
	// MISSING: Config
	// MISSING: SubnetID
	// MISSING: Autoscaling
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: MaxPodsConstraint
	// MISSING: AzureAvailabilityZone
	out.Errors = direct.Slice_FromProto(mapCtx, in.Errors, AzureNodePoolError_FromProto)
	// MISSING: Management
	return out
}
func AzureNodePoolObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AzureNodePoolObservedState) *pb.AzureNodePool {
	if in == nil {
		return nil
	}
	out := &pb.AzureNodePool{}
	// MISSING: Name
	// MISSING: Version
	// MISSING: Config
	// MISSING: SubnetID
	// MISSING: Autoscaling
	out.State = direct.Enum_ToProto[pb.AzureNodePool_State](mapCtx, in.State)
	out.Uid = direct.ValueOf(in.Uid)
	out.Reconciling = direct.ValueOf(in.Reconciling)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: MaxPodsConstraint
	// MISSING: AzureAvailabilityZone
	out.Errors = direct.Slice_ToProto(mapCtx, in.Errors, AzureNodePoolError_ToProto)
	// MISSING: Management
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
	if in.GetProject() != "" {
		out.ProjectRef = &refs.FleetProjectRef{External: in.GetProject()}
	}
	out.Membership = direct.LazyPtr(in.GetMembership())
	return out
}
func Fleet_ToProto(mapCtx *direct.MapContext, in *krm.Fleet) *pb.Fleet {
	if in == nil {
		return nil
	}
	out := &pb.Fleet{}
	if in.ProjectRef != nil {
		out.Project = in.ProjectRef.External
	}
	out.Membership = direct.ValueOf(in.Membership)
	return out
}
func GkemulticloudAzureNodePoolObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AzureNodePool) *krm.GkemulticloudAzureNodePoolObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GkemulticloudAzureNodePoolObservedState{}
	// MISSING: Name
	// MISSING: Version
	// MISSING: Config
	// MISSING: SubnetID
	// MISSING: Autoscaling
	// MISSING: State
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: MaxPodsConstraint
	// MISSING: AzureAvailabilityZone
	// MISSING: Errors
	// MISSING: Management
	return out
}
func GkemulticloudAzureNodePoolObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GkemulticloudAzureNodePoolObservedState) *pb.AzureNodePool {
	if in == nil {
		return nil
	}
	out := &pb.AzureNodePool{}
	// MISSING: Name
	// MISSING: Version
	// MISSING: Config
	// MISSING: SubnetID
	// MISSING: Autoscaling
	// MISSING: State
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: MaxPodsConstraint
	// MISSING: AzureAvailabilityZone
	// MISSING: Errors
	// MISSING: Management
	return out
}
func GkemulticloudAzureNodePoolSpec_FromProto(mapCtx *direct.MapContext, in *pb.AzureNodePool) *krm.GkemulticloudAzureNodePoolSpec {
	if in == nil {
		return nil
	}
	out := &krm.GkemulticloudAzureNodePoolSpec{}
	// MISSING: Name
	// MISSING: Version
	// MISSING: Config
	// MISSING: SubnetID
	// MISSING: Autoscaling
	// MISSING: State
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: MaxPodsConstraint
	// MISSING: AzureAvailabilityZone
	// MISSING: Errors
	// MISSING: Management
	return out
}
func GkemulticloudAzureNodePoolSpec_ToProto(mapCtx *direct.MapContext, in *krm.GkemulticloudAzureNodePoolSpec) *pb.AzureNodePool {
	if in == nil {
		return nil
	}
	out := &pb.AzureNodePool{}
	// MISSING: Name
	// MISSING: Version
	// MISSING: Config
	// MISSING: SubnetID
	// MISSING: Autoscaling
	// MISSING: State
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: MaxPodsConstraint
	// MISSING: AzureAvailabilityZone
	// MISSING: Errors
	// MISSING: Management
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
func MaxPodsConstraint_FromProto(mapCtx *direct.MapContext, in *pb.MaxPodsConstraint) *krm.MaxPodsConstraint {
	if in == nil {
		return nil
	}
	out := &krm.MaxPodsConstraint{}
	out.MaxPodsPerNode = direct.LazyPtr(in.GetMaxPodsPerNode())
	return out
}
func MaxPodsConstraint_ToProto(mapCtx *direct.MapContext, in *krm.MaxPodsConstraint) *pb.MaxPodsConstraint {
	if in == nil {
		return nil
	}
	out := &pb.MaxPodsConstraint{}
	out.MaxPodsPerNode = direct.ValueOf(in.MaxPodsPerNode)
	return out
}
func MonitoringConfig_FromProto(mapCtx *direct.MapContext, in *pb.MonitoringConfig) *krm.MonitoringConfig {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringConfig{}
	out.ManagedPrometheusConfig = ManagedPrometheusConfig_FromProto(mapCtx, in.GetManagedPrometheusConfig())
	// MISSING: CloudMonitoringConfig
	return out
}
func MonitoringConfig_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringConfig) *pb.MonitoringConfig {
	if in == nil {
		return nil
	}
	out := &pb.MonitoringConfig{}
	out.ManagedPrometheusConfig = ManagedPrometheusConfig_ToProto(mapCtx, in.ManagedPrometheusConfig)
	// MISSING: CloudMonitoringConfig
	return out
}
func NodeTaint_FromProto(mapCtx *direct.MapContext, in *pb.NodeTaint) *krm.NodeTaint {
	if in == nil {
		return nil
	}
	out := &krm.NodeTaint{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.Value = direct.LazyPtr(in.GetValue())
	out.Effect = direct.Enum_FromProto(mapCtx, in.GetEffect())
	return out
}
func NodeTaint_ToProto(mapCtx *direct.MapContext, in *krm.NodeTaint) *pb.NodeTaint {
	if in == nil {
		return nil
	}
	out := &pb.NodeTaint{}
	out.Key = direct.ValueOf(in.Key)
	out.Value = direct.ValueOf(in.Value)
	out.Effect = direct.Enum_ToProto[pb.NodeTaint_Effect](mapCtx, in.Effect)
	return out
}
func WorkloadIdentityConfig_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadIdentityConfig) *krm.WorkloadIdentityConfig {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadIdentityConfig{}
	// MISSING: IssuerURI
	// (near miss): "IssuerURI" vs "IssuerUri"
	out.WorkloadPool = direct.LazyPtr(in.GetWorkloadPool())
	out.IdentityProvider = direct.LazyPtr(in.GetIdentityProvider())
	return out
}
func WorkloadIdentityConfig_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadIdentityConfig) *pb.WorkloadIdentityConfig {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadIdentityConfig{}
	// MISSING: IssuerURI
	// (near miss): "IssuerURI" vs "IssuerUri"
	out.WorkloadPool = direct.ValueOf(in.WorkloadPool)
	out.IdentityProvider = direct.ValueOf(in.IdentityProvider)
	return out
}
