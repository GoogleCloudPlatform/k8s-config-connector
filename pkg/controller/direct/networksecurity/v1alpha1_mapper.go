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

// krm.group: networksecurity.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.networksecurity.v1

package networksecurity

import (
	krmcomputev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpclients/generated/google/cloud/networksecurity/v1"
)

func InterceptEndpointGroup_AssociationDetailsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InterceptEndpointGroup_AssociationDetails) *krm.InterceptEndpointGroup_AssociationDetailsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InterceptEndpointGroup_AssociationDetailsObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func InterceptEndpointGroup_AssociationDetailsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InterceptEndpointGroup_AssociationDetailsObservedState) *pb.InterceptEndpointGroup_AssociationDetails {
	if in == nil {
		return nil
	}
	out := &pb.InterceptEndpointGroup_AssociationDetails{}
	out.Name = direct.ValueOf(in.Name)
	out.Network = direct.ValueOf(in.Network)
	out.State = direct.Enum_ToProto[pb.InterceptEndpointGroupAssociation_State](mapCtx, in.State)
	return out
}
func InterceptEndpointGroup_ConnectedDeploymentGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InterceptEndpointGroup_ConnectedDeploymentGroup) *krm.InterceptEndpointGroup_ConnectedDeploymentGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InterceptEndpointGroup_ConnectedDeploymentGroupObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Locations = direct.Slice_FromProto(mapCtx, in.Locations, InterceptLocationObservedState_FromProto)
	return out
}
func InterceptEndpointGroup_ConnectedDeploymentGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InterceptEndpointGroup_ConnectedDeploymentGroupObservedState) *pb.InterceptEndpointGroup_ConnectedDeploymentGroup {
	if in == nil {
		return nil
	}
	out := &pb.InterceptEndpointGroup_ConnectedDeploymentGroup{}
	out.Name = direct.ValueOf(in.Name)
	out.Locations = direct.Slice_ToProto(mapCtx, in.Locations, InterceptLocationObservedState_ToProto)
	return out
}
func InterceptLocationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InterceptLocation) *krm.InterceptLocationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InterceptLocationObservedState{}
	out.Location = direct.LazyPtr(in.GetLocation())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func InterceptLocationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InterceptLocationObservedState) *pb.InterceptLocation {
	if in == nil {
		return nil
	}
	out := &pb.InterceptLocation{}
	out.Location = direct.ValueOf(in.Location)
	out.State = direct.Enum_ToProto[pb.InterceptLocation_State](mapCtx, in.State)
	return out
}
func MirroringLocationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MirroringLocation) *krm.MirroringLocationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MirroringLocationObservedState{}
	out.Location = direct.LazyPtr(in.GetLocation())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func MirroringLocationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MirroringLocationObservedState) *pb.MirroringLocation {
	if in == nil {
		return nil
	}
	out := &pb.MirroringLocation{}
	out.Location = direct.ValueOf(in.Location)
	out.State = direct.Enum_ToProto[pb.MirroringLocation_State](mapCtx, in.State)
	return out
}
func NetworkSecurityBackendAuthenticationConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackendAuthenticationConfig) *krm.NetworkSecurityBackendAuthenticationConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecurityBackendAuthenticationConfigObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func NetworkSecurityBackendAuthenticationConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkSecurityBackendAuthenticationConfigObservedState) *pb.BackendAuthenticationConfig {
	if in == nil {
		return nil
	}
	out := &pb.BackendAuthenticationConfig{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func NetworkSecurityBackendAuthenticationConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.BackendAuthenticationConfig) *krm.NetworkSecurityBackendAuthenticationConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecurityBackendAuthenticationConfigSpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	if in.GetClientCertificate() != "" {
		out.ClientCertificateRef = &refsv1beta1.CertificateManagerCertificateRef{External: in.GetClientCertificate()}
	}
	if in.GetTrustConfig() != "" {
		out.TrustConfigRef = &refsv1beta1.CertificateManagerTrustConfigRef{External: in.GetTrustConfig()}
	}
	out.WellKnownRoots = direct.Enum_FromProto(mapCtx, in.GetWellKnownRoots())
	return out
}
func NetworkSecurityBackendAuthenticationConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkSecurityBackendAuthenticationConfigSpec) *pb.BackendAuthenticationConfig {
	if in == nil {
		return nil
	}
	out := &pb.BackendAuthenticationConfig{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	if in.ClientCertificateRef != nil {
		out.ClientCertificate = in.ClientCertificateRef.External
	}
	if in.TrustConfigRef != nil {
		out.TrustConfig = in.TrustConfigRef.External
	}
	out.WellKnownRoots = direct.Enum_ToProto[pb.BackendAuthenticationConfig_WellKnownRoots](mapCtx, in.WellKnownRoots)
	return out
}
func NetworkSecurityInterceptDeploymentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InterceptDeployment) *krm.NetworkSecurityInterceptDeploymentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecurityInterceptDeploymentObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	return out
}
func NetworkSecurityInterceptDeploymentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkSecurityInterceptDeploymentObservedState) *pb.InterceptDeployment {
	if in == nil {
		return nil
	}
	out := &pb.InterceptDeployment{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.InterceptDeployment_State](mapCtx, in.State)
	out.Reconciling = direct.ValueOf(in.Reconciling)
	return out
}
func NetworkSecurityInterceptDeploymentSpec_FromProto(mapCtx *direct.MapContext, in *pb.InterceptDeployment) *krm.NetworkSecurityInterceptDeploymentSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecurityInterceptDeploymentSpec{}
	// MISSING: Name
	out.Labels = in.Labels
	if in.GetForwardingRule() != "" {
		out.ForwardingRuleRef = &refsv1beta1.ComputeForwardingRuleRef{External: in.GetForwardingRule()}
	}
	if in.GetInterceptDeploymentGroup() != "" {
		out.InterceptDeploymentGroupRef = &refsv1beta1.NetworkSecurityInterceptDeploymentGroupRef{External: in.GetInterceptDeploymentGroup()}
	}
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func NetworkSecurityInterceptDeploymentSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkSecurityInterceptDeploymentSpec) *pb.InterceptDeployment {
	if in == nil {
		return nil
	}
	out := &pb.InterceptDeployment{}
	// MISSING: Name
	out.Labels = in.Labels
	if in.ForwardingRuleRef != nil {
		out.ForwardingRule = in.ForwardingRuleRef.External
	}
	if in.InterceptDeploymentGroupRef != nil {
		out.InterceptDeploymentGroup = in.InterceptDeploymentGroupRef.External
	}
	out.Description = direct.ValueOf(in.Description)
	return out
}
func NetworkSecurityInterceptEndpointGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InterceptEndpointGroup) *krm.NetworkSecurityInterceptEndpointGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecurityInterceptEndpointGroupObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.ConnectedDeploymentGroup = InterceptEndpointGroup_ConnectedDeploymentGroupObservedState_FromProto(mapCtx, in.GetConnectedDeploymentGroup())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	out.Associations = direct.Slice_FromProto(mapCtx, in.Associations, InterceptEndpointGroup_AssociationDetailsObservedState_FromProto)
	return out
}
func NetworkSecurityInterceptEndpointGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkSecurityInterceptEndpointGroupObservedState) *pb.InterceptEndpointGroup {
	if in == nil {
		return nil
	}
	out := &pb.InterceptEndpointGroup{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.ConnectedDeploymentGroup = InterceptEndpointGroup_ConnectedDeploymentGroupObservedState_ToProto(mapCtx, in.ConnectedDeploymentGroup)
	out.State = direct.Enum_ToProto[pb.InterceptEndpointGroup_State](mapCtx, in.State)
	out.Reconciling = direct.ValueOf(in.Reconciling)
	out.Associations = direct.Slice_ToProto(mapCtx, in.Associations, InterceptEndpointGroup_AssociationDetailsObservedState_ToProto)
	return out
}
func NetworkSecurityInterceptEndpointGroupSpec_FromProto(mapCtx *direct.MapContext, in *pb.InterceptEndpointGroup) *krm.NetworkSecurityInterceptEndpointGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecurityInterceptEndpointGroupSpec{}
	// MISSING: Name
	out.Labels = in.Labels
	if in.GetInterceptDeploymentGroup() != "" {
		out.InterceptDeploymentGroupRef = &refsv1beta1.NetworkSecurityInterceptDeploymentGroupRef{External: in.GetInterceptDeploymentGroup()}
	}
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func NetworkSecurityInterceptEndpointGroupSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkSecurityInterceptEndpointGroupSpec) *pb.InterceptEndpointGroup {
	if in == nil {
		return nil
	}
	out := &pb.InterceptEndpointGroup{}
	// MISSING: Name
	out.Labels = in.Labels
	if in.InterceptDeploymentGroupRef != nil {
		out.InterceptDeploymentGroup = in.InterceptDeploymentGroupRef.External
	}
	out.Description = direct.ValueOf(in.Description)
	return out
}
func NetworkSecurityMirroringEndpointGroupAssociationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MirroringEndpointGroupAssociation) *krm.NetworkSecurityMirroringEndpointGroupAssociationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecurityMirroringEndpointGroupAssociationObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.LocationsDetails = direct.Slice_FromProto(mapCtx, in.LocationsDetails, MirroringEndpointGroupAssociationLocationDetailsObservedState_FromProto)
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	out.Locations = direct.Slice_FromProto(mapCtx, in.Locations, MirroringLocationObservedState_FromProto)
	return out
}
func NetworkSecurityMirroringEndpointGroupAssociationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkSecurityMirroringEndpointGroupAssociationObservedState) *pb.MirroringEndpointGroupAssociation {
	if in == nil {
		return nil
	}
	out := &pb.MirroringEndpointGroupAssociation{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.LocationsDetails = direct.Slice_ToProto(mapCtx, in.LocationsDetails, MirroringEndpointGroupAssociationLocationDetailsObservedState_ToProto)
	out.State = direct.Enum_ToProto[pb.MirroringEndpointGroupAssociation_State](mapCtx, in.State)
	out.Reconciling = direct.ValueOf(in.Reconciling)
	out.Locations = direct.Slice_ToProto(mapCtx, in.Locations, MirroringLocationObservedState_ToProto)
	return out
}
func NetworkSecurityMirroringEndpointGroupAssociationSpec_FromProto(mapCtx *direct.MapContext, in *pb.MirroringEndpointGroupAssociation) *krm.NetworkSecurityMirroringEndpointGroupAssociationSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecurityMirroringEndpointGroupAssociationSpec{}
	// MISSING: Name
	out.Labels = in.Labels
	if in.GetMirroringEndpointGroup() != "" {
		out.MirroringEndpointGroupRef = &refsv1beta1.NetworkSecurityMirroringEndpointGroupRef{External: in.GetMirroringEndpointGroup()}
	}
	if in.GetNetwork() != "" {
		out.NetworkRef = &krmcomputev1beta1.ComputeNetworkRef{External: in.GetNetwork()}
	}
	return out
}
func NetworkSecurityMirroringEndpointGroupAssociationSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkSecurityMirroringEndpointGroupAssociationSpec) *pb.MirroringEndpointGroupAssociation {
	if in == nil {
		return nil
	}
	out := &pb.MirroringEndpointGroupAssociation{}
	// MISSING: Name
	out.Labels = in.Labels
	if in.MirroringEndpointGroupRef != nil {
		out.MirroringEndpointGroup = in.MirroringEndpointGroupRef.External
	}
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	return out
}

func MirroringEndpointGroupAssociationLocationDetailsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MirroringEndpointGroupAssociation_LocationDetails) *krm.MirroringEndpointGroupAssociationLocationDetailsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MirroringEndpointGroupAssociationLocationDetailsObservedState{}
	out.Location = direct.LazyPtr(in.GetLocation())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func MirroringEndpointGroupAssociationLocationDetailsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MirroringEndpointGroupAssociationLocationDetailsObservedState) *pb.MirroringEndpointGroupAssociation_LocationDetails {
	if in == nil {
		return nil
	}
	out := &pb.MirroringEndpointGroupAssociation_LocationDetails{}
	out.Location = direct.ValueOf(in.Location)
	out.State = direct.Enum_ToProto[pb.MirroringEndpointGroupAssociation_LocationDetails_State](mapCtx, in.State)
	return out
}
