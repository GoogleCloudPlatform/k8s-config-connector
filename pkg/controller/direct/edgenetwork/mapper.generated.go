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

package edgenetwork

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/edgenetwork/apiv1/edgenetworkpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/edgenetwork/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func EdgenetworkInterconnectAttachmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InterconnectAttachment) *krm.EdgenetworkInterconnectAttachmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EdgenetworkInterconnectAttachmentObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Interconnect
	// MISSING: Network
	// MISSING: VlanID
	// MISSING: Mtu
	// MISSING: State
	return out
}
func EdgenetworkInterconnectAttachmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EdgenetworkInterconnectAttachmentObservedState) *pb.InterconnectAttachment {
	if in == nil {
		return nil
	}
	out := &pb.InterconnectAttachment{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Interconnect
	// MISSING: Network
	// MISSING: VlanID
	// MISSING: Mtu
	// MISSING: State
	return out
}
func EdgenetworkInterconnectAttachmentSpec_FromProto(mapCtx *direct.MapContext, in *pb.InterconnectAttachment) *krm.EdgenetworkInterconnectAttachmentSpec {
	if in == nil {
		return nil
	}
	out := &krm.EdgenetworkInterconnectAttachmentSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Interconnect
	// MISSING: Network
	// MISSING: VlanID
	// MISSING: Mtu
	// MISSING: State
	return out
}
func EdgenetworkInterconnectAttachmentSpec_ToProto(mapCtx *direct.MapContext, in *krm.EdgenetworkInterconnectAttachmentSpec) *pb.InterconnectAttachment {
	if in == nil {
		return nil
	}
	out := &pb.InterconnectAttachment{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Interconnect
	// MISSING: Network
	// MISSING: VlanID
	// MISSING: Mtu
	// MISSING: State
	return out
}
func InterconnectAttachment_FromProto(mapCtx *direct.MapContext, in *pb.InterconnectAttachment) *krm.InterconnectAttachment {
	if in == nil {
		return nil
	}
	out := &krm.InterconnectAttachment{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Interconnect = direct.LazyPtr(in.GetInterconnect())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.VlanID = direct.LazyPtr(in.GetVlanId())
	out.Mtu = direct.LazyPtr(in.GetMtu())
	// MISSING: State
	return out
}
func InterconnectAttachment_ToProto(mapCtx *direct.MapContext, in *krm.InterconnectAttachment) *pb.InterconnectAttachment {
	if in == nil {
		return nil
	}
	out := &pb.InterconnectAttachment{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	out.Interconnect = direct.ValueOf(in.Interconnect)
	out.Network = direct.ValueOf(in.Network)
	out.VlanId = direct.ValueOf(in.VlanID)
	out.Mtu = direct.ValueOf(in.Mtu)
	// MISSING: State
	return out
}
func InterconnectAttachmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InterconnectAttachment) *krm.InterconnectAttachmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InterconnectAttachmentObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Interconnect
	// MISSING: Network
	// MISSING: VlanID
	// MISSING: Mtu
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func InterconnectAttachmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InterconnectAttachmentObservedState) *pb.InterconnectAttachment {
	if in == nil {
		return nil
	}
	out := &pb.InterconnectAttachment{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Interconnect
	// MISSING: Network
	// MISSING: VlanID
	// MISSING: Mtu
	out.State = direct.Enum_ToProto[pb.ResourceState](mapCtx, in.State)
	return out
}
