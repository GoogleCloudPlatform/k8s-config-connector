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
func AzureK8sVersionInfo_FromProto(mapCtx *direct.MapContext, in *pb.AzureK8sVersionInfo) *krm.AzureK8sVersionInfo {
	if in == nil {
		return nil
	}
	out := &krm.AzureK8sVersionInfo{}
	out.Version = direct.LazyPtr(in.GetVersion())
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	out.EndOfLife = direct.LazyPtr(in.GetEndOfLife())
	out.EndOfLifeDate = Date_FromProto(mapCtx, in.GetEndOfLifeDate())
	out.ReleaseDate = Date_FromProto(mapCtx, in.GetReleaseDate())
	return out
}
func AzureK8sVersionInfo_ToProto(mapCtx *direct.MapContext, in *krm.AzureK8sVersionInfo) *pb.AzureK8sVersionInfo {
	if in == nil {
		return nil
	}
	out := &pb.AzureK8sVersionInfo{}
	out.Version = direct.ValueOf(in.Version)
	out.Enabled = direct.ValueOf(in.Enabled)
	out.EndOfLife = direct.ValueOf(in.EndOfLife)
	out.EndOfLifeDate = Date_ToProto(mapCtx, in.EndOfLifeDate)
	out.ReleaseDate = Date_ToProto(mapCtx, in.ReleaseDate)
	return out
}
func AzureServerConfig_FromProto(mapCtx *direct.MapContext, in *pb.AzureServerConfig) *krm.AzureServerConfig {
	if in == nil {
		return nil
	}
	out := &krm.AzureServerConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ValidVersions = direct.Slice_FromProto(mapCtx, in.ValidVersions, AzureK8sVersionInfo_FromProto)
	out.SupportedAzureRegions = in.SupportedAzureRegions
	return out
}
func AzureServerConfig_ToProto(mapCtx *direct.MapContext, in *krm.AzureServerConfig) *pb.AzureServerConfig {
	if in == nil {
		return nil
	}
	out := &pb.AzureServerConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.ValidVersions = direct.Slice_ToProto(mapCtx, in.ValidVersions, AzureK8sVersionInfo_ToProto)
	out.SupportedAzureRegions = in.SupportedAzureRegions
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
func GkemulticloudAzureServerConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AzureServerConfig) *krm.GkemulticloudAzureServerConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GkemulticloudAzureServerConfigObservedState{}
	// MISSING: Name
	// MISSING: ValidVersions
	// MISSING: SupportedAzureRegions
	return out
}
func GkemulticloudAzureServerConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GkemulticloudAzureServerConfigObservedState) *pb.AzureServerConfig {
	if in == nil {
		return nil
	}
	out := &pb.AzureServerConfig{}
	// MISSING: Name
	// MISSING: ValidVersions
	// MISSING: SupportedAzureRegions
	return out
}
func GkemulticloudAzureServerConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.AzureServerConfig) *krm.GkemulticloudAzureServerConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.GkemulticloudAzureServerConfigSpec{}
	// MISSING: Name
	// MISSING: ValidVersions
	// MISSING: SupportedAzureRegions
	return out
}
func GkemulticloudAzureServerConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.GkemulticloudAzureServerConfigSpec) *pb.AzureServerConfig {
	if in == nil {
		return nil
	}
	out := &pb.AzureServerConfig{}
	// MISSING: Name
	// MISSING: ValidVersions
	// MISSING: SupportedAzureRegions
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
