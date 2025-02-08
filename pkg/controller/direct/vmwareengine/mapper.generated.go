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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func DnsForwarding_FromProto(mapCtx *direct.MapContext, in *pb.DnsForwarding) *krm.DnsForwarding {
	if in == nil {
		return nil
	}
	out := &krm.DnsForwarding{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.ForwardingRules = direct.Slice_FromProto(mapCtx, in.ForwardingRules, DnsForwarding_ForwardingRule_FromProto)
	return out
}
func DnsForwarding_ToProto(mapCtx *direct.MapContext, in *krm.DnsForwarding) *pb.DnsForwarding {
	if in == nil {
		return nil
	}
	out := &pb.DnsForwarding{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.ForwardingRules = direct.Slice_ToProto(mapCtx, in.ForwardingRules, DnsForwarding_ForwardingRule_ToProto)
	return out
}
func DnsForwardingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DnsForwarding) *krm.DnsForwardingObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DnsForwardingObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: ForwardingRules
	return out
}
func DnsForwardingObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DnsForwardingObservedState) *pb.DnsForwarding {
	if in == nil {
		return nil
	}
	out := &pb.DnsForwarding{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: ForwardingRules
	return out
}
func DnsForwarding_ForwardingRule_FromProto(mapCtx *direct.MapContext, in *pb.DnsForwarding_ForwardingRule) *krm.DnsForwarding_ForwardingRule {
	if in == nil {
		return nil
	}
	out := &krm.DnsForwarding_ForwardingRule{}
	out.Domain = direct.LazyPtr(in.GetDomain())
	out.NameServers = in.NameServers
	return out
}
func DnsForwarding_ForwardingRule_ToProto(mapCtx *direct.MapContext, in *krm.DnsForwarding_ForwardingRule) *pb.DnsForwarding_ForwardingRule {
	if in == nil {
		return nil
	}
	out := &pb.DnsForwarding_ForwardingRule{}
	out.Domain = direct.ValueOf(in.Domain)
	out.NameServers = in.NameServers
	return out
}
func VmwareengineDnsForwardingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DnsForwarding) *krm.VmwareengineDnsForwardingObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmwareengineDnsForwardingObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ForwardingRules
	return out
}
func VmwareengineDnsForwardingObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmwareengineDnsForwardingObservedState) *pb.DnsForwarding {
	if in == nil {
		return nil
	}
	out := &pb.DnsForwarding{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ForwardingRules
	return out
}
func VmwareengineDnsForwardingSpec_FromProto(mapCtx *direct.MapContext, in *pb.DnsForwarding) *krm.VmwareengineDnsForwardingSpec {
	if in == nil {
		return nil
	}
	out := &krm.VmwareengineDnsForwardingSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ForwardingRules
	return out
}
func VmwareengineDnsForwardingSpec_ToProto(mapCtx *direct.MapContext, in *krm.VmwareengineDnsForwardingSpec) *pb.DnsForwarding {
	if in == nil {
		return nil
	}
	out := &pb.DnsForwarding{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ForwardingRules
	return out
}
