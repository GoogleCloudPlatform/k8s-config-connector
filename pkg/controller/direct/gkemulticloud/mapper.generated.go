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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkemulticloud/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/gkemulticloud/apiv1/gkemulticloudpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/containerattached/v1beta1"
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
func AwsAutoscalingGroupMetricsCollection_FromProto(mapCtx *direct.MapContext, in *pb.AwsAutoscalingGroupMetricsCollection) *krm.AwsAutoscalingGroupMetricsCollection {
	if in == nil {
		return nil
	}
	out := &krm.AwsAutoscalingGroupMetricsCollection{}
	out.Granularity = direct.LazyPtr(in.GetGranularity())
	out.Metrics = in.Metrics
	return out
}
func AwsAutoscalingGroupMetricsCollection_ToProto(mapCtx *direct.MapContext, in *krm.AwsAutoscalingGroupMetricsCollection) *pb.AwsAutoscalingGroupMetricsCollection {
	if in == nil {
		return nil
	}
	out := &pb.AwsAutoscalingGroupMetricsCollection{}
	out.Granularity = direct.ValueOf(in.Granularity)
	out.Metrics = in.Metrics
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
func AwsNodeConfig_FromProto(mapCtx *direct.MapContext, in *pb.AwsNodeConfig) *krm.AwsNodeConfig {
	if in == nil {
		return nil
	}
	out := &krm.AwsNodeConfig{}
	out.InstanceType = direct.LazyPtr(in.GetInstanceType())
	out.RootVolume = AwsVolumeTemplate_FromProto(mapCtx, in.GetRootVolume())
	out.Taints = direct.Slice_FromProto(mapCtx, in.Taints, NodeTaint_FromProto)
	out.Labels = in.Labels
	out.Tags = in.Tags
	out.IamInstanceProfile = direct.LazyPtr(in.GetIamInstanceProfile())
	out.ImageType = direct.LazyPtr(in.GetImageType())
	out.SSHConfig = AwsSshConfig_FromProto(mapCtx, in.GetSshConfig())
	out.SecurityGroupIds = in.SecurityGroupIds
	out.ProxyConfig = AwsProxyConfig_FromProto(mapCtx, in.GetProxyConfig())
	out.ConfigEncryption = AwsConfigEncryption_FromProto(mapCtx, in.GetConfigEncryption())
	out.InstancePlacement = AwsInstancePlacement_FromProto(mapCtx, in.GetInstancePlacement())
	out.AutoscalingMetricsCollection = AwsAutoscalingGroupMetricsCollection_FromProto(mapCtx, in.GetAutoscalingMetricsCollection())
	out.SpotConfig = SpotConfig_FromProto(mapCtx, in.GetSpotConfig())
	return out
}
func AwsNodeConfig_ToProto(mapCtx *direct.MapContext, in *krm.AwsNodeConfig) *pb.AwsNodeConfig {
	if in == nil {
		return nil
	}
	out := &pb.AwsNodeConfig{}
	out.InstanceType = direct.ValueOf(in.InstanceType)
	out.RootVolume = AwsVolumeTemplate_ToProto(mapCtx, in.RootVolume)
	out.Taints = direct.Slice_ToProto(mapCtx, in.Taints, NodeTaint_ToProto)
	out.Labels = in.Labels
	out.Tags = in.Tags
	out.IamInstanceProfile = direct.ValueOf(in.IamInstanceProfile)
	out.ImageType = direct.ValueOf(in.ImageType)
	out.SshConfig = AwsSshConfig_ToProto(mapCtx, in.SSHConfig)
	out.SecurityGroupIds = in.SecurityGroupIds
	out.ProxyConfig = AwsProxyConfig_ToProto(mapCtx, in.ProxyConfig)
	out.ConfigEncryption = AwsConfigEncryption_ToProto(mapCtx, in.ConfigEncryption)
	out.InstancePlacement = AwsInstancePlacement_ToProto(mapCtx, in.InstancePlacement)
	out.AutoscalingMetricsCollection = AwsAutoscalingGroupMetricsCollection_ToProto(mapCtx, in.AutoscalingMetricsCollection)
	out.SpotConfig = SpotConfig_ToProto(mapCtx, in.SpotConfig)
	return out
}
func AwsNodeManagement_FromProto(mapCtx *direct.MapContext, in *pb.AwsNodeManagement) *krm.AwsNodeManagement {
	if in == nil {
		return nil
	}
	out := &krm.AwsNodeManagement{}
	out.AutoRepair = direct.LazyPtr(in.GetAutoRepair())
	return out
}
func AwsNodeManagement_ToProto(mapCtx *direct.MapContext, in *krm.AwsNodeManagement) *pb.AwsNodeManagement {
	if in == nil {
		return nil
	}
	out := &pb.AwsNodeManagement{}
	out.AutoRepair = direct.ValueOf(in.AutoRepair)
	return out
}
func AwsNodePool_FromProto(mapCtx *direct.MapContext, in *pb.AwsNodePool) *krm.AwsNodePool {
	if in == nil {
		return nil
	}
	out := &krm.AwsNodePool{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Version = direct.LazyPtr(in.GetVersion())
	out.Config = AwsNodeConfig_FromProto(mapCtx, in.GetConfig())
	out.Autoscaling = AwsNodePoolAutoscaling_FromProto(mapCtx, in.GetAutoscaling())
	out.SubnetID = direct.LazyPtr(in.GetSubnetId())
	// MISSING: State
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Annotations = in.Annotations
	out.MaxPodsConstraint = MaxPodsConstraint_FromProto(mapCtx, in.GetMaxPodsConstraint())
	// MISSING: Errors
	out.Management = AwsNodeManagement_FromProto(mapCtx, in.GetManagement())
	out.KubeletConfig = NodeKubeletConfig_FromProto(mapCtx, in.GetKubeletConfig())
	out.UpdateSettings = UpdateSettings_FromProto(mapCtx, in.GetUpdateSettings())
	return out
}
func AwsNodePool_ToProto(mapCtx *direct.MapContext, in *krm.AwsNodePool) *pb.AwsNodePool {
	if in == nil {
		return nil
	}
	out := &pb.AwsNodePool{}
	out.Name = direct.ValueOf(in.Name)
	out.Version = direct.ValueOf(in.Version)
	out.Config = AwsNodeConfig_ToProto(mapCtx, in.Config)
	out.Autoscaling = AwsNodePoolAutoscaling_ToProto(mapCtx, in.Autoscaling)
	out.SubnetId = direct.ValueOf(in.SubnetID)
	// MISSING: State
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Etag = direct.ValueOf(in.Etag)
	out.Annotations = in.Annotations
	out.MaxPodsConstraint = MaxPodsConstraint_ToProto(mapCtx, in.MaxPodsConstraint)
	// MISSING: Errors
	out.Management = AwsNodeManagement_ToProto(mapCtx, in.Management)
	out.KubeletConfig = NodeKubeletConfig_ToProto(mapCtx, in.KubeletConfig)
	out.UpdateSettings = UpdateSettings_ToProto(mapCtx, in.UpdateSettings)
	return out
}
func AwsNodePoolAutoscaling_FromProto(mapCtx *direct.MapContext, in *pb.AwsNodePoolAutoscaling) *krm.AwsNodePoolAutoscaling {
	if in == nil {
		return nil
	}
	out := &krm.AwsNodePoolAutoscaling{}
	out.MinNodeCount = direct.LazyPtr(in.GetMinNodeCount())
	out.MaxNodeCount = direct.LazyPtr(in.GetMaxNodeCount())
	return out
}
func AwsNodePoolAutoscaling_ToProto(mapCtx *direct.MapContext, in *krm.AwsNodePoolAutoscaling) *pb.AwsNodePoolAutoscaling {
	if in == nil {
		return nil
	}
	out := &pb.AwsNodePoolAutoscaling{}
	out.MinNodeCount = direct.ValueOf(in.MinNodeCount)
	out.MaxNodeCount = direct.ValueOf(in.MaxNodeCount)
	return out
}
func AwsNodePoolError_FromProto(mapCtx *direct.MapContext, in *pb.AwsNodePoolError) *krm.AwsNodePoolError {
	if in == nil {
		return nil
	}
	out := &krm.AwsNodePoolError{}
	out.Message = direct.LazyPtr(in.GetMessage())
	return out
}
func AwsNodePoolError_ToProto(mapCtx *direct.MapContext, in *krm.AwsNodePoolError) *pb.AwsNodePoolError {
	if in == nil {
		return nil
	}
	out := &pb.AwsNodePoolError{}
	out.Message = direct.ValueOf(in.Message)
	return out
}
func AwsNodePoolObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AwsNodePool) *krm.AwsNodePoolObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AwsNodePoolObservedState{}
	// MISSING: Name
	// MISSING: Version
	// MISSING: Config
	// MISSING: Autoscaling
	// MISSING: SubnetID
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: MaxPodsConstraint
	out.Errors = direct.Slice_FromProto(mapCtx, in.Errors, AwsNodePoolError_FromProto)
	// MISSING: Management
	// MISSING: KubeletConfig
	// MISSING: UpdateSettings
	return out
}
func AwsNodePoolObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AwsNodePoolObservedState) *pb.AwsNodePool {
	if in == nil {
		return nil
	}
	out := &pb.AwsNodePool{}
	// MISSING: Name
	// MISSING: Version
	// MISSING: Config
	// MISSING: Autoscaling
	// MISSING: SubnetID
	out.State = direct.Enum_ToProto[pb.AwsNodePool_State](mapCtx, in.State)
	out.Uid = direct.ValueOf(in.Uid)
	out.Reconciling = direct.ValueOf(in.Reconciling)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: MaxPodsConstraint
	out.Errors = direct.Slice_ToProto(mapCtx, in.Errors, AwsNodePoolError_ToProto)
	// MISSING: Management
	// MISSING: KubeletConfig
	// MISSING: UpdateSettings
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
func GkemulticloudAwsNodePoolObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AwsNodePool) *krm.GkemulticloudAwsNodePoolObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GkemulticloudAwsNodePoolObservedState{}
	// MISSING: Name
	// MISSING: Version
	// MISSING: Config
	// MISSING: Autoscaling
	// MISSING: SubnetID
	// MISSING: State
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: MaxPodsConstraint
	// MISSING: Errors
	// MISSING: Management
	// MISSING: KubeletConfig
	// MISSING: UpdateSettings
	return out
}
func GkemulticloudAwsNodePoolObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GkemulticloudAwsNodePoolObservedState) *pb.AwsNodePool {
	if in == nil {
		return nil
	}
	out := &pb.AwsNodePool{}
	// MISSING: Name
	// MISSING: Version
	// MISSING: Config
	// MISSING: Autoscaling
	// MISSING: SubnetID
	// MISSING: State
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: MaxPodsConstraint
	// MISSING: Errors
	// MISSING: Management
	// MISSING: KubeletConfig
	// MISSING: UpdateSettings
	return out
}
func GkemulticloudAwsNodePoolSpec_FromProto(mapCtx *direct.MapContext, in *pb.AwsNodePool) *krm.GkemulticloudAwsNodePoolSpec {
	if in == nil {
		return nil
	}
	out := &krm.GkemulticloudAwsNodePoolSpec{}
	// MISSING: Name
	// MISSING: Version
	// MISSING: Config
	// MISSING: Autoscaling
	// MISSING: SubnetID
	// MISSING: State
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: MaxPodsConstraint
	// MISSING: Errors
	// MISSING: Management
	// MISSING: KubeletConfig
	// MISSING: UpdateSettings
	return out
}
func GkemulticloudAwsNodePoolSpec_ToProto(mapCtx *direct.MapContext, in *krm.GkemulticloudAwsNodePoolSpec) *pb.AwsNodePool {
	if in == nil {
		return nil
	}
	out := &pb.AwsNodePool{}
	// MISSING: Name
	// MISSING: Version
	// MISSING: Config
	// MISSING: Autoscaling
	// MISSING: SubnetID
	// MISSING: State
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: MaxPodsConstraint
	// MISSING: Errors
	// MISSING: Management
	// MISSING: KubeletConfig
	// MISSING: UpdateSettings
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
func NodeKubeletConfig_FromProto(mapCtx *direct.MapContext, in *pb.NodeKubeletConfig) *krm.NodeKubeletConfig {
	if in == nil {
		return nil
	}
	out := &krm.NodeKubeletConfig{}
	out.InsecureKubeletReadonlyPortEnabled = direct.LazyPtr(in.GetInsecureKubeletReadonlyPortEnabled())
	out.CpuManagerPolicy = in.CpuManagerPolicy
	out.CpuCfsQuota = in.CpuCfsQuota
	out.CpuCfsQuotaPeriod = in.CpuCfsQuotaPeriod
	out.PodPidsLimit = in.PodPidsLimit
	return out
}
func NodeKubeletConfig_ToProto(mapCtx *direct.MapContext, in *krm.NodeKubeletConfig) *pb.NodeKubeletConfig {
	if in == nil {
		return nil
	}
	out := &pb.NodeKubeletConfig{}
	out.InsecureKubeletReadonlyPortEnabled = direct.ValueOf(in.InsecureKubeletReadonlyPortEnabled)
	out.CpuManagerPolicy = in.CpuManagerPolicy
	out.CpuCfsQuota = in.CpuCfsQuota
	out.CpuCfsQuotaPeriod = in.CpuCfsQuotaPeriod
	out.PodPidsLimit = in.PodPidsLimit
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
func SpotConfig_FromProto(mapCtx *direct.MapContext, in *pb.SpotConfig) *krm.SpotConfig {
	if in == nil {
		return nil
	}
	out := &krm.SpotConfig{}
	out.InstanceTypes = in.InstanceTypes
	return out
}
func SpotConfig_ToProto(mapCtx *direct.MapContext, in *krm.SpotConfig) *pb.SpotConfig {
	if in == nil {
		return nil
	}
	out := &pb.SpotConfig{}
	out.InstanceTypes = in.InstanceTypes
	return out
}
func SurgeSettings_FromProto(mapCtx *direct.MapContext, in *pb.SurgeSettings) *krm.SurgeSettings {
	if in == nil {
		return nil
	}
	out := &krm.SurgeSettings{}
	out.MaxSurge = direct.LazyPtr(in.GetMaxSurge())
	out.MaxUnavailable = direct.LazyPtr(in.GetMaxUnavailable())
	return out
}
func SurgeSettings_ToProto(mapCtx *direct.MapContext, in *krm.SurgeSettings) *pb.SurgeSettings {
	if in == nil {
		return nil
	}
	out := &pb.SurgeSettings{}
	out.MaxSurge = direct.ValueOf(in.MaxSurge)
	out.MaxUnavailable = direct.ValueOf(in.MaxUnavailable)
	return out
}
func UpdateSettings_FromProto(mapCtx *direct.MapContext, in *pb.UpdateSettings) *krm.UpdateSettings {
	if in == nil {
		return nil
	}
	out := &krm.UpdateSettings{}
	out.SurgeSettings = SurgeSettings_FromProto(mapCtx, in.GetSurgeSettings())
	return out
}
func UpdateSettings_ToProto(mapCtx *direct.MapContext, in *krm.UpdateSettings) *pb.UpdateSettings {
	if in == nil {
		return nil
	}
	out := &pb.UpdateSettings{}
	out.SurgeSettings = SurgeSettings_ToProto(mapCtx, in.SurgeSettings)
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
