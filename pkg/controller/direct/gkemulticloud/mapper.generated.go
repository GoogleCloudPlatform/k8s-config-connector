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
)

func GkeMultiCloudAttachedClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.AttachedCluster) *krm.AttachedCluster {
	if in == nil {
		return nil
	}
	out := &krm.AttachedCluster{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.OIDCConfig = AttachedOIDCConfig_FromProto(mapCtx, in.GetOidcConfig())
	out.PlatformVersion = direct.LazyPtr(in.GetPlatformVersion())
	out.Distribution = direct.LazyPtr(in.GetDistribution())
	out.Fleet = Fleet_FromProto(mapCtx, in.GetFleet())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Annotations = in.Annotations
	out.LoggingConfig = LoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	out.Authorization = AttachedClustersAuthorization_FromProto(mapCtx, in.GetAuthorization())
	out.MonitoringConfig = MonitoringConfig_FromProto(mapCtx, in.GetMonitoringConfig())
	out.ProxyConfig = AttachedProxyConfig_FromProto(mapCtx, in.GetProxyConfig())
	out.BinaryAuthorization = BinaryAuthorization_FromProto(mapCtx, in.GetBinaryAuthorization())
	out.SecurityPostureConfig = SecurityPostureConfig_FromProto(mapCtx, in.GetSecurityPostureConfig())
	out.Tags = in.Tags
	return out
}
func GkeMultiCloudAttachedClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.AttachedCluster) *pb.AttachedCluster {
	if in == nil {
		return nil
	}
	out := &pb.AttachedCluster{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.OidcConfig = AttachedOIDCConfig_ToProto(mapCtx, in.OIDCConfig)
	out.PlatformVersion = direct.ValueOf(in.PlatformVersion)
	out.Distribution = direct.ValueOf(in.Distribution)
	out.Fleet = Fleet_ToProto(mapCtx, in.Fleet)
	out.Etag = direct.ValueOf(in.Etag)
	out.Annotations = in.Annotations
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	out.Authorization = AttachedClustersAuthorization_ToProto(mapCtx, in.Authorization)
	out.MonitoringConfig = MonitoringConfig_ToProto(mapCtx, in.MonitoringConfig)
	out.ProxyConfig = AttachedProxyConfig_ToProto(mapCtx, in.ProxyConfig)
	out.BinaryAuthorization = BinaryAuthorization_ToProto(mapCtx, in.BinaryAuthorization)
	out.SecurityPostureConfig = SecurityPostureConfig_ToProto(mapCtx, in.SecurityPostureConfig)
	out.Tags = in.Tags
	return out
}
func GkeMultiCloudAttachedClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AttachedCluster) *krm.AttachedClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AttachedClusterObservedState{}
	out.ClusterRegion = direct.LazyPtr(in.GetClusterRegion())
	out.Fleet = FleetObservedState_FromProto(mapCtx, in.GetFleet())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.KubernetesVersion = direct.LazyPtr(in.GetKubernetesVersion())
	out.WorkloadIdentityConfig = WorkloadIdentityConfig_FromProto(mapCtx, in.GetWorkloadIdentityConfig())
	out.Errors = direct.Slice_FromProto(mapCtx, in.Errors, AttachedClusterError_FromProto)
	return out
}
func GkeMultiCloudAttachedClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AttachedClusterObservedState) *pb.AttachedCluster {
	if in == nil {
		return nil
	}
	out := &pb.AttachedCluster{}
	out.ClusterRegion = direct.ValueOf(in.ClusterRegion)
	out.Fleet = FleetObservedState_ToProto(mapCtx, in.Fleet)
	out.State = direct.Enum_ToProto[pb.AttachedCluster_State](mapCtx, in.State)
	out.Uid = direct.ValueOf(in.Uid)
	out.Reconciling = direct.ValueOf(in.Reconciling)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.KubernetesVersion = direct.ValueOf(in.KubernetesVersion)
	out.WorkloadIdentityConfig = WorkloadIdentityConfig_ToProto(mapCtx, in.WorkloadIdentityConfig)
	out.Errors = direct.Slice_ToProto(mapCtx, in.Errors, AttachedClusterError_ToProto)
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
func AttachedOIDCConfig_FromProto(mapCtx *direct.MapContext, in *pb.AttachedOidcConfig) *krm.AttachedOIDCConfig {
	if in == nil {
		return nil
	}
	out := &krm.AttachedOIDCConfig{}
	out.IssuerURL = direct.LazyPtr(in.GetIssuerUrl())
	out.Jwks = in.GetJwks()
	return out
}
func AttachedOIDCConfig_ToProto(mapCtx *direct.MapContext, in *krm.AttachedOIDCConfig) *pb.AttachedOidcConfig {
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
