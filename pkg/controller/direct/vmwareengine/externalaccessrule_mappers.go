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
)

func VMwareEngineExternalAccessRuleSpec_FromProto(mapCtx *direct.MapContext, in *pb.ExternalAccessRule) *krm.VMwareEngineExternalAccessRuleSpec {
	if in == nil {
		return nil
	}
	out := &krm.VMwareEngineExternalAccessRuleSpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Priority = direct.LazyPtr(in.GetPriority())
	out.Action = direct.Enum_FromProto(mapCtx, in.GetAction())
	out.IPProtocol = direct.LazyPtr(in.GetIpProtocol())
	out.SourceIPRanges = direct.Slice_FromProto(mapCtx, in.SourceIpRanges, ExternalAccessRule_IPRange_FromProto)
	out.SourcePorts = in.SourcePorts
	out.DestinationIPRanges = direct.Slice_FromProto(mapCtx, in.DestinationIpRanges, ExternalAccessRule_IPRange_FromProto)
	out.DestinationPorts = in.DestinationPorts
	return out
}
func VMwareEngineExternalAccessRuleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ExternalAccessRule) *krm.VMwareEngineExternalAccessRuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VMwareEngineExternalAccessRuleObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.UID = direct.LazyPtr(in.GetUid())
	return out
}
func VMwareEngineExternalAccessRuleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VMwareEngineExternalAccessRuleObservedState) *pb.ExternalAccessRule {
	if in == nil {
		return nil
	}
	out := &pb.ExternalAccessRule{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.ExternalAccessRule_State](mapCtx, in.State)
	out.Uid = direct.ValueOf(in.UID)
	return out
}
func ExternalAccessRule_IPRange_FromProto(mapCtx *direct.MapContext, in *pb.ExternalAccessRule_IpRange) *krm.ExternalAccessRule_IPRange {
	if in == nil {
		return nil
	}
	out := &krm.ExternalAccessRule_IPRange{}
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	out.IPAddressRange = direct.LazyPtr(in.GetIpAddressRange())
	if in.GetExternalAddress() != "" {
		out.ExternalAddressRef = &krm.ExternalAddressRef{External: in.GetExternalAddress()}
	}
	return out
}
func ExternalAccessRule_IPRange_ToProto(mapCtx *direct.MapContext, in *krm.ExternalAccessRule_IPRange) *pb.ExternalAccessRule_IpRange {
	if in == nil {
		return nil
	}
	out := &pb.ExternalAccessRule_IpRange{}
	if oneof := ExternalAccessRule_IPRange_IpAddress_ToProto(mapCtx, in.IPAddress); oneof != nil {
		out.IpRange = oneof
	}
	if oneof := ExternalAccessRule_IPRange_IpAddressRange_ToProto(mapCtx, in.IPAddressRange); oneof != nil {
		out.IpRange = oneof
	}
	if oneof := ExternalAccessRule_IPRange_ExternalAddress_ToProto(mapCtx, in.ExternalAddressRef); oneof != nil {
		out.IpRange = oneof
	}
	return out
}
func ExternalAccessRule_IPRange_IpAddress_ToProto(mapCtx *direct.MapContext, IPAddress *string) *pb.ExternalAccessRule_IpRange_IpAddress {
	if IPAddress == nil {
		return nil
	}
	return &pb.ExternalAccessRule_IpRange_IpAddress{
		IpAddress: *IPAddress,
	}
}
func ExternalAccessRule_IPRange_IpAddressRange_ToProto(mapCtx *direct.MapContext, IPAddressRange *string) *pb.ExternalAccessRule_IpRange_IpAddressRange {
	if IPAddressRange == nil {
		return nil
	}
	return &pb.ExternalAccessRule_IpRange_IpAddressRange{
		IpAddressRange: *IPAddressRange,
	}
}
func ExternalAccessRule_IPRange_ExternalAddress_ToProto(mapCtx *direct.MapContext, ExternalAddressRef *krm.ExternalAddressRef) *pb.ExternalAccessRule_IpRange_ExternalAddress {
	if ExternalAddressRef == nil {
		return nil
	}
	return &pb.ExternalAccessRule_IpRange_ExternalAddress{
		ExternalAddress: ExternalAddressRef.External,
	}
}
