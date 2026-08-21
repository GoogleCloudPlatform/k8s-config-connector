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

package networksecurity

import (
	pb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func FirewallEndpointEndpointSettings_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.FirewallEndpoint_EndpointSettings) *krm.FirewallEndpointEndpointSettings {
	if in == nil {
		return nil
	}
	out := &krm.FirewallEndpointEndpointSettings{}
	out.JumboFramesEnabled = direct.LazyPtr(in.GetJumboFramesEnabled())
	return out
}

func FirewallEndpointEndpointSettings_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.FirewallEndpointEndpointSettings) *pb.FirewallEndpoint_EndpointSettings {
	if in == nil {
		return nil
	}
	out := &pb.FirewallEndpoint_EndpointSettings{}
	out.JumboFramesEnabled = direct.ValueOf(in.JumboFramesEnabled)
	return out
}

func FirewallEndpointAssociationReference_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.FirewallEndpoint_AssociationReference) *krm.FirewallEndpointAssociationReference {
	if in == nil {
		return nil
	}
	out := &krm.FirewallEndpointAssociationReference{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Network = direct.LazyPtr(in.GetNetwork())
	return out
}

func FirewallEndpointAssociationReference_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.FirewallEndpointAssociationReference) *pb.FirewallEndpoint_AssociationReference {
	if in == nil {
		return nil
	}
	out := &pb.FirewallEndpoint_AssociationReference{}
	out.Name = direct.ValueOf(in.Name)
	out.Network = direct.ValueOf(in.Network)
	return out
}
