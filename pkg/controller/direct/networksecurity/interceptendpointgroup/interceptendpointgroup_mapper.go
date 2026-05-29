// Copyright 2026 Google LLC
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

package interceptendpointgroup

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/cloud/networksecurity/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NetworkSecurityInterceptEndpointGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InterceptEndpointGroup) *krm.NetworkSecurityInterceptEndpointGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecurityInterceptEndpointGroupObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.LazyPtr(in.GetState().String())
	out.Reconciling = direct.LazyPtr(in.GetReconciling())

	if in.ConnectedDeploymentGroup != nil {
		out.ConnectedDeploymentGroup = &krm.InterceptEndpointGroup_ConnectedDeploymentGroupObservedState{}
		out.ConnectedDeploymentGroup.Name = direct.LazyPtr(in.ConnectedDeploymentGroup.GetName())
		if len(in.ConnectedDeploymentGroup.GetLocations()) > 0 {
			out.ConnectedDeploymentGroup.Locations = make([]krm.InterceptLocationObservedState, len(in.ConnectedDeploymentGroup.GetLocations()))
			for i, loc := range in.ConnectedDeploymentGroup.GetLocations() {
				out.ConnectedDeploymentGroup.Locations[i] = krm.InterceptLocationObservedState{
					Location: direct.LazyPtr(loc.GetLocation()),
					State:    direct.LazyPtr(loc.GetState().String()),
				}
			}
		}
	}

	if len(in.GetAssociations()) > 0 {
		out.Associations = make([]krm.InterceptEndpointGroup_AssociationDetailsObservedState, len(in.GetAssociations()))
		for i, assoc := range in.GetAssociations() {
			out.Associations[i] = krm.InterceptEndpointGroup_AssociationDetailsObservedState{
				Name:    direct.LazyPtr(assoc.GetName()),
				Network: direct.LazyPtr(assoc.GetNetwork()),
				State:   direct.LazyPtr(assoc.GetState().String()),
			}
		}
	}

	return out
}

func NetworkSecurityInterceptEndpointGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkSecurityInterceptEndpointGroupObservedState) *pb.InterceptEndpointGroup {
	if in == nil {
		return nil
	}
	out := &pb.InterceptEndpointGroup{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	if in.State != nil {
		if val, ok := pb.InterceptEndpointGroup_State_value[*in.State]; ok {
			out.State = pb.InterceptEndpointGroup_State(val)
		}
	}
	out.Reconciling = direct.ValueOf(in.Reconciling)

	if in.ConnectedDeploymentGroup != nil {
		out.ConnectedDeploymentGroup = &pb.InterceptEndpointGroup_ConnectedDeploymentGroup{}
		out.ConnectedDeploymentGroup.Name = direct.ValueOf(in.ConnectedDeploymentGroup.Name)
		if len(in.ConnectedDeploymentGroup.Locations) > 0 {
			out.ConnectedDeploymentGroup.Locations = make([]*pb.InterceptLocation, len(in.ConnectedDeploymentGroup.Locations))
			for i, loc := range in.ConnectedDeploymentGroup.Locations {
				out.ConnectedDeploymentGroup.Locations[i] = &pb.InterceptLocation{
					Location: direct.ValueOf(loc.Location),
				}
				if loc.State != nil {
					if val, ok := pb.InterceptLocation_State_value[*loc.State]; ok {
						out.ConnectedDeploymentGroup.Locations[i].State = pb.InterceptLocation_State(val)
					}
				}
			}
		}
	}

	if len(in.Associations) > 0 {
		out.Associations = make([]*pb.InterceptEndpointGroup_AssociationDetails, len(in.Associations))
		for i, assoc := range in.Associations {
			out.Associations[i] = &pb.InterceptEndpointGroup_AssociationDetails{}
			out.Associations[i].Name = direct.ValueOf(assoc.Name)
			out.Associations[i].Network = direct.ValueOf(assoc.Network)
			if assoc.State != nil {
				if val, ok := pb.InterceptEndpointGroupAssociation_State_value[*assoc.State]; ok {
					out.Associations[i].State = pb.InterceptEndpointGroupAssociation_State(val)
				}
			}
		}
	}

	return out
}

func NetworkSecurityInterceptEndpointGroupSpec_FromProto(mapCtx *direct.MapContext, in *pb.InterceptEndpointGroup) *krm.NetworkSecurityInterceptEndpointGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecurityInterceptEndpointGroupSpec{}
	out.Labels = in.GetLabels()
	if in.GetInterceptDeploymentGroup() != "" {
		out.InterceptDeploymentGroupRef = &refsv1beta1.NetworkSecurityInterceptDeploymentGroupRef{
			External: in.GetInterceptDeploymentGroup(),
		}
	}
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}

func NetworkSecurityInterceptEndpointGroupSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkSecurityInterceptEndpointGroupSpec) *pb.InterceptEndpointGroup {
	if in == nil {
		return nil
	}
	out := &pb.InterceptEndpointGroup{}
	out.Labels = in.Labels
	if in.InterceptDeploymentGroupRef != nil {
		out.InterceptDeploymentGroup = in.InterceptDeploymentGroupRef.External
	}
	out.Description = direct.ValueOf(in.Description)
	return out
}
