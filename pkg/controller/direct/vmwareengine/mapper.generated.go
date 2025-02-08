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

package vmwareengine

import (
	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func ExternalAccessRule_FromProto(mapCtx *direct.MapContext, in *pb.ExternalAccessRule) *krm.ExternalAccessRule {
	if in == nil {
		return nil
	}
	out := &krm.ExternalAccessRule{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Priority = direct.LazyPtr(in.GetPriority())
	out.Action = direct.Enum_FromProto(mapCtx, in.GetAction())
	out.IPProtocol = direct.LazyPtr(in.GetIpProtocol())
	out.SourceIPRanges = direct.Slice_FromProto(mapCtx, in.SourceIPRanges, ExternalAccessRule_IpRange_FromProto)
	out.SourcePorts = in.SourcePorts
	out.DestinationIPRanges = direct.Slice_FromProto(mapCtx, in.DestinationIPRanges, ExternalAccessRule_IpRange_FromProto)
	out.DestinationPorts = in.DestinationPorts
	// MISSING: State
	// MISSING: Uid
	return out
}
func ExternalAccessRule_ToProto(mapCtx *direct.MapContext, in *krm.ExternalAccessRule) *pb.ExternalAccessRule {
	if in == nil {
		return nil
	}
	out := &pb.ExternalAccessRule{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.ValueOf(in.Description)
	out.Priority = direct.ValueOf(in.Priority)
	out.Action = direct.Enum_ToProto[pb.ExternalAccessRule_Action](mapCtx, in.Action)
	out.IpProtocol = direct.ValueOf(in.IPProtocol)
	out.SourceIpRanges = direct.Slice_ToProto(mapCtx, in.SourceIPRanges, ExternalAccessRule_IpRange_ToProto)
	out.SourcePorts = in.SourcePorts
	out.DestinationIpRanges = direct.Slice_ToProto(mapCtx, in.DestinationIPRanges, ExternalAccessRule_IpRange_ToProto)
	out.DestinationPorts = in.DestinationPorts
	// MISSING: State
	// MISSING: Uid
	return out
}
func ExternalAccessRuleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ExternalAccessRule) *krm.ExternalAccessRuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ExternalAccessRuleObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Description
	// MISSING: Priority
	// MISSING: Action
	// MISSING: IPProtocol
	// MISSING: SourceIPRanges
	// MISSING: SourcePorts
	// MISSING: DestinationIPRanges
	// MISSING: DestinationPorts
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Uid = direct.LazyPtr(in.GetUid())
	return out
}
func ExternalAccessRuleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ExternalAccessRuleObservedState) *pb.ExternalAccessRule {
	if in == nil {
		return nil
	}
	out := &pb.ExternalAccessRule{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Description
	// MISSING: Priority
	// MISSING: Action
	// MISSING: IPProtocol
	// MISSING: SourceIPRanges
	// MISSING: SourcePorts
	// MISSING: DestinationIPRanges
	// MISSING: DestinationPorts
	out.State = direct.Enum_ToProto[pb.ExternalAccessRule_State](mapCtx, in.State)
	out.Uid = direct.ValueOf(in.Uid)
	return out
}
func ExternalAccessRule_IpRange_FromProto(mapCtx *direct.MapContext, in *pb.ExternalAccessRule_IpRange) *krm.ExternalAccessRule_IpRange {
	if in == nil {
		return nil
	}
	out := &krm.ExternalAccessRule_IpRange{}
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	out.IPAddressRange = direct.LazyPtr(in.GetIpAddressRange())
	out.ExternalAddress = direct.LazyPtr(in.GetExternalAddress())
	return out
}
func ExternalAccessRule_IpRange_ToProto(mapCtx *direct.MapContext, in *krm.ExternalAccessRule_IpRange) *pb.ExternalAccessRule_IpRange {
	if in == nil {
		return nil
	}
	out := &pb.ExternalAccessRule_IpRange{}
	if oneof := ExternalAccessRule_IpRange_IpAddress_ToProto(mapCtx, in.IPAddress); oneof != nil {
		out.IpRange = oneof
	}
	if oneof := ExternalAccessRule_IpRange_IpAddressRange_ToProto(mapCtx, in.IPAddressRange); oneof != nil {
		out.IpRange = oneof
	}
	if oneof := ExternalAccessRule_IpRange_ExternalAddress_ToProto(mapCtx, in.ExternalAddress); oneof != nil {
		out.IpRange = oneof
	}
	return out
}
func VmwareengineExternalAccessRuleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ExternalAccessRule) *krm.VmwareengineExternalAccessRuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmwareengineExternalAccessRuleObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: Priority
	// MISSING: Action
	// MISSING: IPProtocol
	// MISSING: SourceIPRanges
	// MISSING: SourcePorts
	// MISSING: DestinationIPRanges
	// MISSING: DestinationPorts
	// MISSING: State
	// MISSING: Uid
	return out
}
func VmwareengineExternalAccessRuleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmwareengineExternalAccessRuleObservedState) *pb.ExternalAccessRule {
	if in == nil {
		return nil
	}
	out := &pb.ExternalAccessRule{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: Priority
	// MISSING: Action
	// MISSING: IPProtocol
	// MISSING: SourceIPRanges
	// MISSING: SourcePorts
	// MISSING: DestinationIPRanges
	// MISSING: DestinationPorts
	// MISSING: State
	// MISSING: Uid
	return out
}
func VmwareengineExternalAccessRuleSpec_FromProto(mapCtx *direct.MapContext, in *pb.ExternalAccessRule) *krm.VmwareengineExternalAccessRuleSpec {
	if in == nil {
		return nil
	}
	out := &krm.VmwareengineExternalAccessRuleSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: Priority
	// MISSING: Action
	// MISSING: IPProtocol
	// MISSING: SourceIPRanges
	// MISSING: SourcePorts
	// MISSING: DestinationIPRanges
	// MISSING: DestinationPorts
	// MISSING: State
	// MISSING: Uid
	return out
}
func VmwareengineExternalAccessRuleSpec_ToProto(mapCtx *direct.MapContext, in *krm.VmwareengineExternalAccessRuleSpec) *pb.ExternalAccessRule {
	if in == nil {
		return nil
	}
	out := &pb.ExternalAccessRule{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: Priority
	// MISSING: Action
	// MISSING: IPProtocol
	// MISSING: SourceIPRanges
	// MISSING: SourcePorts
	// MISSING: DestinationIPRanges
	// MISSING: DestinationPorts
	// MISSING: State
	// MISSING: Uid
	return out
}
