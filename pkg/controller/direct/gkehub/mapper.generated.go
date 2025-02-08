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

package gkehub

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/gkehub/apiv1/gkehubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Authority_FromProto(mapCtx *direct.MapContext, in *pb.Authority) *krm.Authority {
	if in == nil {
		return nil
	}
	out := &krm.Authority{}
	out.Issuer = direct.LazyPtr(in.GetIssuer())
	// MISSING: WorkloadIdentityPool
	// MISSING: IdentityProvider
	out.OidcJwks = in.GetOidcJwks()
	return out
}
func Authority_ToProto(mapCtx *direct.MapContext, in *krm.Authority) *pb.Authority {
	if in == nil {
		return nil
	}
	out := &pb.Authority{}
	out.Issuer = direct.ValueOf(in.Issuer)
	// MISSING: WorkloadIdentityPool
	// MISSING: IdentityProvider
	out.OidcJwks = in.OidcJwks
	return out
}
func AuthorityObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Authority) *krm.AuthorityObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AuthorityObservedState{}
	// MISSING: Issuer
	out.WorkloadIdentityPool = direct.LazyPtr(in.GetWorkloadIdentityPool())
	out.IdentityProvider = direct.LazyPtr(in.GetIdentityProvider())
	// MISSING: OidcJwks
	return out
}
func AuthorityObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AuthorityObservedState) *pb.Authority {
	if in == nil {
		return nil
	}
	out := &pb.Authority{}
	// MISSING: Issuer
	out.WorkloadIdentityPool = direct.ValueOf(in.WorkloadIdentityPool)
	out.IdentityProvider = direct.ValueOf(in.IdentityProvider)
	// MISSING: OidcJwks
	return out
}
func GkeCluster_FromProto(mapCtx *direct.MapContext, in *pb.GkeCluster) *krm.GkeCluster {
	if in == nil {
		return nil
	}
	out := &krm.GkeCluster{}
	out.ResourceLink = direct.LazyPtr(in.GetResourceLink())
	// MISSING: ClusterMissing
	return out
}
func GkeCluster_ToProto(mapCtx *direct.MapContext, in *krm.GkeCluster) *pb.GkeCluster {
	if in == nil {
		return nil
	}
	out := &pb.GkeCluster{}
	out.ResourceLink = direct.ValueOf(in.ResourceLink)
	// MISSING: ClusterMissing
	return out
}
func GkeClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GkeCluster) *krm.GkeClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GkeClusterObservedState{}
	// MISSING: ResourceLink
	out.ClusterMissing = direct.LazyPtr(in.GetClusterMissing())
	return out
}
func GkeClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GkeClusterObservedState) *pb.GkeCluster {
	if in == nil {
		return nil
	}
	out := &pb.GkeCluster{}
	// MISSING: ResourceLink
	out.ClusterMissing = direct.ValueOf(in.ClusterMissing)
	return out
}
func GkehubMembershipObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Membership) *krm.GkehubMembershipObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GkehubMembershipObservedState{}
	// MISSING: Endpoint
	// MISSING: Name
	// MISSING: Labels
	// MISSING: Description
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExternalID
	// MISSING: LastConnectionTime
	// MISSING: UniqueID
	// MISSING: Authority
	// MISSING: MonitoringConfig
	return out
}
func GkehubMembershipObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GkehubMembershipObservedState) *pb.Membership {
	if in == nil {
		return nil
	}
	out := &pb.Membership{}
	// MISSING: Endpoint
	// MISSING: Name
	// MISSING: Labels
	// MISSING: Description
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExternalID
	// MISSING: LastConnectionTime
	// MISSING: UniqueID
	// MISSING: Authority
	// MISSING: MonitoringConfig
	return out
}
func GkehubMembershipSpec_FromProto(mapCtx *direct.MapContext, in *pb.Membership) *krm.GkehubMembershipSpec {
	if in == nil {
		return nil
	}
	out := &krm.GkehubMembershipSpec{}
	// MISSING: Endpoint
	// MISSING: Name
	// MISSING: Labels
	// MISSING: Description
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExternalID
	// MISSING: LastConnectionTime
	// MISSING: UniqueID
	// MISSING: Authority
	// MISSING: MonitoringConfig
	return out
}
func GkehubMembershipSpec_ToProto(mapCtx *direct.MapContext, in *krm.GkehubMembershipSpec) *pb.Membership {
	if in == nil {
		return nil
	}
	out := &pb.Membership{}
	// MISSING: Endpoint
	// MISSING: Name
	// MISSING: Labels
	// MISSING: Description
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExternalID
	// MISSING: LastConnectionTime
	// MISSING: UniqueID
	// MISSING: Authority
	// MISSING: MonitoringConfig
	return out
}
func KubernetesResource_FromProto(mapCtx *direct.MapContext, in *pb.KubernetesResource) *krm.KubernetesResource {
	if in == nil {
		return nil
	}
	out := &krm.KubernetesResource{}
	out.MembershipCrManifest = direct.LazyPtr(in.GetMembershipCrManifest())
	// MISSING: MembershipResources
	// MISSING: ConnectResources
	out.ResourceOptions = ResourceOptions_FromProto(mapCtx, in.GetResourceOptions())
	return out
}
func KubernetesResource_ToProto(mapCtx *direct.MapContext, in *krm.KubernetesResource) *pb.KubernetesResource {
	if in == nil {
		return nil
	}
	out := &pb.KubernetesResource{}
	out.MembershipCrManifest = direct.ValueOf(in.MembershipCrManifest)
	// MISSING: MembershipResources
	// MISSING: ConnectResources
	out.ResourceOptions = ResourceOptions_ToProto(mapCtx, in.ResourceOptions)
	return out
}
func KubernetesResourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.KubernetesResource) *krm.KubernetesResourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.KubernetesResourceObservedState{}
	// MISSING: MembershipCrManifest
	out.MembershipResources = direct.Slice_FromProto(mapCtx, in.MembershipResources, ResourceManifest_FromProto)
	out.ConnectResources = direct.Slice_FromProto(mapCtx, in.ConnectResources, ResourceManifest_FromProto)
	// MISSING: ResourceOptions
	return out
}
func KubernetesResourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.KubernetesResourceObservedState) *pb.KubernetesResource {
	if in == nil {
		return nil
	}
	out := &pb.KubernetesResource{}
	// MISSING: MembershipCrManifest
	out.MembershipResources = direct.Slice_ToProto(mapCtx, in.MembershipResources, ResourceManifest_ToProto)
	out.ConnectResources = direct.Slice_ToProto(mapCtx, in.ConnectResources, ResourceManifest_ToProto)
	// MISSING: ResourceOptions
	return out
}
func Membership_FromProto(mapCtx *direct.MapContext, in *pb.Membership) *krm.Membership {
	if in == nil {
		return nil
	}
	out := &krm.Membership{}
	out.Endpoint = MembershipEndpoint_FromProto(mapCtx, in.GetEndpoint())
	// MISSING: Name
	out.Labels = in.Labels
	// MISSING: Description
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	out.ExternalID = direct.LazyPtr(in.GetExternalId())
	// MISSING: LastConnectionTime
	// MISSING: UniqueID
	out.Authority = Authority_FromProto(mapCtx, in.GetAuthority())
	out.MonitoringConfig = MonitoringConfig_FromProto(mapCtx, in.GetMonitoringConfig())
	return out
}
func Membership_ToProto(mapCtx *direct.MapContext, in *krm.Membership) *pb.Membership {
	if in == nil {
		return nil
	}
	out := &pb.Membership{}
	if oneof := MembershipEndpoint_ToProto(mapCtx, in.Endpoint); oneof != nil {
		out.Type = &pb.Membership_Endpoint{Endpoint: oneof}
	}
	// MISSING: Name
	out.Labels = in.Labels
	// MISSING: Description
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	out.ExternalId = direct.ValueOf(in.ExternalID)
	// MISSING: LastConnectionTime
	// MISSING: UniqueID
	out.Authority = Authority_ToProto(mapCtx, in.Authority)
	out.MonitoringConfig = MonitoringConfig_ToProto(mapCtx, in.MonitoringConfig)
	return out
}
func MembershipEndpoint_FromProto(mapCtx *direct.MapContext, in *pb.MembershipEndpoint) *krm.MembershipEndpoint {
	if in == nil {
		return nil
	}
	out := &krm.MembershipEndpoint{}
	out.GkeCluster = GkeCluster_FromProto(mapCtx, in.GetGkeCluster())
	// MISSING: KubernetesMetadata
	out.KubernetesResource = KubernetesResource_FromProto(mapCtx, in.GetKubernetesResource())
	// MISSING: GoogleManaged
	return out
}
func MembershipEndpoint_ToProto(mapCtx *direct.MapContext, in *krm.MembershipEndpoint) *pb.MembershipEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.MembershipEndpoint{}
	out.GkeCluster = GkeCluster_ToProto(mapCtx, in.GkeCluster)
	// MISSING: KubernetesMetadata
	out.KubernetesResource = KubernetesResource_ToProto(mapCtx, in.KubernetesResource)
	// MISSING: GoogleManaged
	return out
}
func MembershipEndpointObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MembershipEndpoint) *krm.MembershipEndpointObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MembershipEndpointObservedState{}
	out.GkeCluster = GkeClusterObservedState_FromProto(mapCtx, in.GetGkeCluster())
	out.KubernetesMetadata = KubernetesMetadata_FromProto(mapCtx, in.GetKubernetesMetadata())
	out.KubernetesResource = KubernetesResourceObservedState_FromProto(mapCtx, in.GetKubernetesResource())
	out.GoogleManaged = direct.LazyPtr(in.GetGoogleManaged())
	return out
}
func MembershipEndpointObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MembershipEndpointObservedState) *pb.MembershipEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.MembershipEndpoint{}
	out.GkeCluster = GkeClusterObservedState_ToProto(mapCtx, in.GkeCluster)
	out.KubernetesMetadata = KubernetesMetadata_ToProto(mapCtx, in.KubernetesMetadata)
	out.KubernetesResource = KubernetesResourceObservedState_ToProto(mapCtx, in.KubernetesResource)
	out.GoogleManaged = direct.ValueOf(in.GoogleManaged)
	return out
}
func MembershipObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Membership) *krm.MembershipObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MembershipObservedState{}
	out.Endpoint = MembershipEndpointObservedState_FromProto(mapCtx, in.GetEndpoint())
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.State = MembershipState_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	// MISSING: ExternalID
	out.LastConnectionTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastConnectionTime())
	out.UniqueID = direct.LazyPtr(in.GetUniqueId())
	out.Authority = AuthorityObservedState_FromProto(mapCtx, in.GetAuthority())
	// MISSING: MonitoringConfig
	return out
}
func MembershipObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MembershipObservedState) *pb.Membership {
	if in == nil {
		return nil
	}
	out := &pb.Membership{}
	if oneof := MembershipEndpointObservedState_ToProto(mapCtx, in.Endpoint); oneof != nil {
		out.Type = &pb.Membership_Endpoint{Endpoint: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Labels
	out.Description = direct.ValueOf(in.Description)
	out.State = MembershipState_ToProto(mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	// MISSING: ExternalID
	out.LastConnectionTime = direct.StringTimestamp_ToProto(mapCtx, in.LastConnectionTime)
	out.UniqueId = direct.ValueOf(in.UniqueID)
	out.Authority = AuthorityObservedState_ToProto(mapCtx, in.Authority)
	// MISSING: MonitoringConfig
	return out
}
func MembershipState_FromProto(mapCtx *direct.MapContext, in *pb.MembershipState) *krm.MembershipState {
	if in == nil {
		return nil
	}
	out := &krm.MembershipState{}
	// MISSING: Code
	return out
}
func MembershipState_ToProto(mapCtx *direct.MapContext, in *krm.MembershipState) *pb.MembershipState {
	if in == nil {
		return nil
	}
	out := &pb.MembershipState{}
	// MISSING: Code
	return out
}
func MembershipStateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MembershipState) *krm.MembershipStateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MembershipStateObservedState{}
	out.Code = direct.Enum_FromProto(mapCtx, in.GetCode())
	return out
}
func MembershipStateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MembershipStateObservedState) *pb.MembershipState {
	if in == nil {
		return nil
	}
	out := &pb.MembershipState{}
	out.Code = direct.Enum_ToProto[pb.MembershipState_Code](mapCtx, in.Code)
	return out
}
func MonitoringConfig_FromProto(mapCtx *direct.MapContext, in *pb.MonitoringConfig) *krm.MonitoringConfig {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringConfig{}
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.Cluster = direct.LazyPtr(in.GetCluster())
	out.KubernetesMetricsPrefix = direct.LazyPtr(in.GetKubernetesMetricsPrefix())
	out.ClusterHash = direct.LazyPtr(in.GetClusterHash())
	return out
}
func MonitoringConfig_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringConfig) *pb.MonitoringConfig {
	if in == nil {
		return nil
	}
	out := &pb.MonitoringConfig{}
	out.ProjectId = direct.ValueOf(in.ProjectID)
	out.Location = direct.ValueOf(in.Location)
	out.Cluster = direct.ValueOf(in.Cluster)
	out.KubernetesMetricsPrefix = direct.ValueOf(in.KubernetesMetricsPrefix)
	out.ClusterHash = direct.ValueOf(in.ClusterHash)
	return out
}
func ResourceManifest_FromProto(mapCtx *direct.MapContext, in *pb.ResourceManifest) *krm.ResourceManifest {
	if in == nil {
		return nil
	}
	out := &krm.ResourceManifest{}
	out.Manifest = direct.LazyPtr(in.GetManifest())
	out.ClusterScoped = direct.LazyPtr(in.GetClusterScoped())
	return out
}
func ResourceManifest_ToProto(mapCtx *direct.MapContext, in *krm.ResourceManifest) *pb.ResourceManifest {
	if in == nil {
		return nil
	}
	out := &pb.ResourceManifest{}
	out.Manifest = direct.ValueOf(in.Manifest)
	out.ClusterScoped = direct.ValueOf(in.ClusterScoped)
	return out
}
func ResourceOptions_FromProto(mapCtx *direct.MapContext, in *pb.ResourceOptions) *krm.ResourceOptions {
	if in == nil {
		return nil
	}
	out := &krm.ResourceOptions{}
	out.ConnectVersion = direct.LazyPtr(in.GetConnectVersion())
	out.V1beta1Crd = direct.LazyPtr(in.GetV1beta1Crd())
	out.K8sVersion = direct.LazyPtr(in.GetK8sVersion())
	return out
}
func ResourceOptions_ToProto(mapCtx *direct.MapContext, in *krm.ResourceOptions) *pb.ResourceOptions {
	if in == nil {
		return nil
	}
	out := &pb.ResourceOptions{}
	out.ConnectVersion = direct.ValueOf(in.ConnectVersion)
	out.V1beta1Crd = direct.ValueOf(in.V1beta1Crd)
	out.K8sVersion = direct.ValueOf(in.K8sVersion)
	return out
}
