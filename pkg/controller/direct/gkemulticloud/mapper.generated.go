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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkemulticloud/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/containerattached/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func AttachedCluster_FromProto(mapCtx *direct.MapContext, in *pb.AttachedCluster) *krm.AttachedCluster {
	if in == nil {
		return nil
	}
	out := &krm.AttachedCluster{}
	out.Name = direct.LazyPtr(in.GetName())
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
	out.Etag = direct.LazyPtr(in.GetEtag())
	// MISSING: KubernetesVersion
	out.Annotations = in.Annotations
	// MISSING: WorkloadIdentityConfig
	out.LoggingConfig = LoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	// MISSING: Errors
	out.Authorization = AttachedClustersAuthorization_FromProto(mapCtx, in.GetAuthorization())
	out.MonitoringConfig = MonitoringConfig_FromProto(mapCtx, in.GetMonitoringConfig())
	out.ProxyConfig = AttachedProxyConfig_FromProto(mapCtx, in.GetProxyConfig())
	out.BinaryAuthorization = BinaryAuthorization_FromProto(mapCtx, in.GetBinaryAuthorization())
	out.SecurityPostureConfig = SecurityPostureConfig_FromProto(mapCtx, in.GetSecurityPostureConfig())
	out.Tags = in.Tags
	return out
}
func AttachedCluster_ToProto(mapCtx *direct.MapContext, in *krm.AttachedCluster) *pb.AttachedCluster {
	if in == nil {
		return nil
	}
	out := &pb.AttachedCluster{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.OidcConfig = AttachedOidcConfig_ToProto(mapCtx, in.OidcConfig)
	out.PlatformVersion = direct.ValueOf(in.PlatformVersion)
	out.Distribution = direct.ValueOf(in.Distribution)
	// MISSING: ClusterRegion
	out.Fleet = Fleet_ToProto(mapCtx, in.Fleet)
	// MISSING: State
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Etag = direct.ValueOf(in.Etag)
	// MISSING: KubernetesVersion
	out.Annotations = in.Annotations
	// MISSING: WorkloadIdentityConfig
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	// MISSING: Errors
	out.Authorization = AttachedClustersAuthorization_ToProto(mapCtx, in.Authorization)
	out.MonitoringConfig = MonitoringConfig_ToProto(mapCtx, in.MonitoringConfig)
	out.ProxyConfig = AttachedProxyConfig_ToProto(mapCtx, in.ProxyConfig)
	out.BinaryAuthorization = BinaryAuthorization_ToProto(mapCtx, in.BinaryAuthorization)
	out.SecurityPostureConfig = SecurityPostureConfig_ToProto(mapCtx, in.SecurityPostureConfig)
	out.Tags = in.Tags
	return out
}
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
func AttachedClusterGroup_FromProto(mapCtx *direct.MapContext, in *pb.AttachedClusterGroup) *krm.AttachedClusterGroup {
	if in == nil {
		return nil
	}
	out := &krm.AttachedClusterGroup{}
	out.Group = direct.LazyPtr(in.GetGroup())
	return out
}
func AttachedClusterGroup_ToProto(mapCtx *direct.MapContext, in *krm.AttachedClusterGroup) *pb.AttachedClusterGroup {
	if in == nil {
		return nil
	}
	out := &pb.AttachedClusterGroup{}
	out.Group = direct.ValueOf(in.Group)
	return out
}
func AttachedClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AttachedCluster) *krm.AttachedClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AttachedClusterObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: OidcConfig
	// MISSING: PlatformVersion
	// MISSING: Distribution
	out.ClusterRegion = direct.LazyPtr(in.GetClusterRegion())
	out.Fleet = FleetObservedState_FromProto(mapCtx, in.GetFleet())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Etag
	out.KubernetesVersion = direct.LazyPtr(in.GetKubernetesVersion())
	// MISSING: Annotations
	out.WorkloadIdentityConfig = WorkloadIdentityConfig_FromProto(mapCtx, in.GetWorkloadIdentityConfig())
	// MISSING: LoggingConfig
	out.Errors = direct.Slice_FromProto(mapCtx, in.Errors, AttachedClusterError_FromProto)
	// MISSING: Authorization
	// MISSING: MonitoringConfig
	// MISSING: ProxyConfig
	// MISSING: BinaryAuthorization
	// MISSING: SecurityPostureConfig
	// MISSING: Tags
	return out
}
func AttachedClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AttachedClusterObservedState) *pb.AttachedCluster {
	if in == nil {
		return nil
	}
	out := &pb.AttachedCluster{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: OidcConfig
	// MISSING: PlatformVersion
	// MISSING: Distribution
	out.ClusterRegion = direct.ValueOf(in.ClusterRegion)
	out.Fleet = FleetObservedState_ToProto(mapCtx, in.Fleet)
	out.State = direct.Enum_ToProto[pb.AttachedCluster_State](mapCtx, in.State)
	out.Uid = direct.ValueOf(in.Uid)
	out.Reconciling = direct.ValueOf(in.Reconciling)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Etag
	out.KubernetesVersion = direct.ValueOf(in.KubernetesVersion)
	// MISSING: Annotations
	out.WorkloadIdentityConfig = WorkloadIdentityConfig_ToProto(mapCtx, in.WorkloadIdentityConfig)
	// MISSING: LoggingConfig
	out.Errors = direct.Slice_ToProto(mapCtx, in.Errors, AttachedClusterError_ToProto)
	// MISSING: Authorization
	// MISSING: MonitoringConfig
	// MISSING: ProxyConfig
	// MISSING: BinaryAuthorization
	// MISSING: SecurityPostureConfig
	// MISSING: Tags
	return out
}
func AttachedClusterUser_FromProto(mapCtx *direct.MapContext, in *pb.AttachedClusterUser) *krm.AttachedClusterUser {
	if in == nil {
		return nil
	}
	out := &krm.AttachedClusterUser{}
	out.Username = direct.LazyPtr(in.GetUsername())
	return out
}
func AttachedClusterUser_ToProto(mapCtx *direct.MapContext, in *krm.AttachedClusterUser) *pb.AttachedClusterUser {
	if in == nil {
		return nil
	}
	out := &pb.AttachedClusterUser{}
	out.Username = direct.ValueOf(in.Username)
	return out
}
func AttachedClustersAuthorization_FromProto(mapCtx *direct.MapContext, in *pb.AttachedClustersAuthorization) *krm.AttachedClustersAuthorization {
	if in == nil {
		return nil
	}
	out := &krm.AttachedClustersAuthorization{}
	out.AdminUsers = direct.Slice_FromProto(mapCtx, in.AdminUsers, AttachedClusterUser_FromProto)
	out.AdminGroups = direct.Slice_FromProto(mapCtx, in.AdminGroups, AttachedClusterGroup_FromProto)
	return out
}
func AttachedClustersAuthorization_ToProto(mapCtx *direct.MapContext, in *krm.AttachedClustersAuthorization) *pb.AttachedClustersAuthorization {
	if in == nil {
		return nil
	}
	out := &pb.AttachedClustersAuthorization{}
	out.AdminUsers = direct.Slice_ToProto(mapCtx, in.AdminUsers, AttachedClusterUser_ToProto)
	out.AdminGroups = direct.Slice_ToProto(mapCtx, in.AdminGroups, AttachedClusterGroup_ToProto)
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
	out.IssuerUrl = direct.ValueOf(in.IssuerURL)
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
func GkemulticloudAttachedClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AttachedCluster) *krm.GkemulticloudAttachedClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GkemulticloudAttachedClusterObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: OidcConfig
	// MISSING: PlatformVersion
	// MISSING: Distribution
	// MISSING: ClusterRegion
	// MISSING: Fleet
	// MISSING: State
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: KubernetesVersion
	// MISSING: Annotations
	// MISSING: WorkloadIdentityConfig
	// MISSING: LoggingConfig
	// MISSING: Errors
	// MISSING: Authorization
	// MISSING: MonitoringConfig
	// MISSING: ProxyConfig
	// MISSING: BinaryAuthorization
	// MISSING: SecurityPostureConfig
	// MISSING: Tags
	return out
}
func GkemulticloudAttachedClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GkemulticloudAttachedClusterObservedState) *pb.AttachedCluster {
	if in == nil {
		return nil
	}
	out := &pb.AttachedCluster{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: OidcConfig
	// MISSING: PlatformVersion
	// MISSING: Distribution
	// MISSING: ClusterRegion
	// MISSING: Fleet
	// MISSING: State
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: KubernetesVersion
	// MISSING: Annotations
	// MISSING: WorkloadIdentityConfig
	// MISSING: LoggingConfig
	// MISSING: Errors
	// MISSING: Authorization
	// MISSING: MonitoringConfig
	// MISSING: ProxyConfig
	// MISSING: BinaryAuthorization
	// MISSING: SecurityPostureConfig
	// MISSING: Tags
	return out
}
func GkemulticloudAttachedClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.AttachedCluster) *krm.GkemulticloudAttachedClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.GkemulticloudAttachedClusterSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: OidcConfig
	// MISSING: PlatformVersion
	// MISSING: Distribution
	// MISSING: ClusterRegion
	// MISSING: Fleet
	// MISSING: State
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: KubernetesVersion
	// MISSING: Annotations
	// MISSING: WorkloadIdentityConfig
	// MISSING: LoggingConfig
	// MISSING: Errors
	// MISSING: Authorization
	// MISSING: MonitoringConfig
	// MISSING: ProxyConfig
	// MISSING: BinaryAuthorization
	// MISSING: SecurityPostureConfig
	// MISSING: Tags
	return out
}
func GkemulticloudAttachedClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.GkemulticloudAttachedClusterSpec) *pb.AttachedCluster {
	if in == nil {
		return nil
	}
	out := &pb.AttachedCluster{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: OidcConfig
	// MISSING: PlatformVersion
	// MISSING: Distribution
	// MISSING: ClusterRegion
	// MISSING: Fleet
	// MISSING: State
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: KubernetesVersion
	// MISSING: Annotations
	// MISSING: WorkloadIdentityConfig
	// MISSING: LoggingConfig
	// MISSING: Errors
	// MISSING: Authorization
	// MISSING: MonitoringConfig
	// MISSING: ProxyConfig
	// MISSING: BinaryAuthorization
	// MISSING: SecurityPostureConfig
	// MISSING: Tags
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
	out.Name = direct.ValueOf(in.Name)
	out.Namespace = direct.ValueOf(in.Namespace)
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
func SecurityPostureConfig_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPostureConfig) *krm.SecurityPostureConfig {
	if in == nil {
		return nil
	}
	out := &krm.SecurityPostureConfig{}
	out.VulnerabilityMode = direct.Enum_FromProto(mapCtx, in.GetVulnerabilityMode())
	return out
}
func SecurityPostureConfig_ToProto(mapCtx *direct.MapContext, in *krm.SecurityPostureConfig) *pb.SecurityPostureConfig {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPostureConfig{}
	out.VulnerabilityMode = direct.Enum_ToProto[pb.SecurityPostureConfig_VulnerabilityMode](mapCtx, in.VulnerabilityMode)
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
