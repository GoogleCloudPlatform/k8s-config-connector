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
func Node_FromProto(mapCtx *direct.MapContext, in *pb.Node) *krm.Node {
	if in == nil {
		return nil
	}
	out := &krm.Node{}
	// MISSING: Name
	// MISSING: Fqdn
	// MISSING: InternalIP
	// MISSING: NodeTypeID
	// MISSING: Version
	// MISSING: CustomCoreCount
	// MISSING: State
	return out
}
func Node_ToProto(mapCtx *direct.MapContext, in *krm.Node) *pb.Node {
	if in == nil {
		return nil
	}
	out := &pb.Node{}
	// MISSING: Name
	// MISSING: Fqdn
	// MISSING: InternalIP
	// MISSING: NodeTypeID
	// MISSING: Version
	// MISSING: CustomCoreCount
	// MISSING: State
	return out
}
func NodeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Node) *krm.NodeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NodeObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Fqdn = direct.LazyPtr(in.GetFqdn())
	out.InternalIP = direct.LazyPtr(in.GetInternalIp())
	out.NodeTypeID = direct.LazyPtr(in.GetNodeTypeId())
	out.Version = direct.LazyPtr(in.GetVersion())
	out.CustomCoreCount = direct.LazyPtr(in.GetCustomCoreCount())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func NodeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NodeObservedState) *pb.Node {
	if in == nil {
		return nil
	}
	out := &pb.Node{}
	out.Name = direct.ValueOf(in.Name)
	out.Fqdn = direct.ValueOf(in.Fqdn)
	out.InternalIp = direct.ValueOf(in.InternalIP)
	out.NodeTypeId = direct.ValueOf(in.NodeTypeID)
	out.Version = direct.ValueOf(in.Version)
	out.CustomCoreCount = direct.ValueOf(in.CustomCoreCount)
	out.State = direct.Enum_ToProto[pb.Node_State](mapCtx, in.State)
	return out
}
func VmwareengineNodeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Node) *krm.VmwareengineNodeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmwareengineNodeObservedState{}
	// MISSING: Name
	// MISSING: Fqdn
	// MISSING: InternalIP
	// MISSING: NodeTypeID
	// MISSING: Version
	// MISSING: CustomCoreCount
	// MISSING: State
	return out
}
func VmwareengineNodeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmwareengineNodeObservedState) *pb.Node {
	if in == nil {
		return nil
	}
	out := &pb.Node{}
	// MISSING: Name
	// MISSING: Fqdn
	// MISSING: InternalIP
	// MISSING: NodeTypeID
	// MISSING: Version
	// MISSING: CustomCoreCount
	// MISSING: State
	return out
}
func VmwareengineNodeSpec_FromProto(mapCtx *direct.MapContext, in *pb.Node) *krm.VmwareengineNodeSpec {
	if in == nil {
		return nil
	}
	out := &krm.VmwareengineNodeSpec{}
	// MISSING: Name
	// MISSING: Fqdn
	// MISSING: InternalIP
	// MISSING: NodeTypeID
	// MISSING: Version
	// MISSING: CustomCoreCount
	// MISSING: State
	return out
}
func VmwareengineNodeSpec_ToProto(mapCtx *direct.MapContext, in *krm.VmwareengineNodeSpec) *pb.Node {
	if in == nil {
		return nil
	}
	out := &pb.Node{}
	// MISSING: Name
	// MISSING: Fqdn
	// MISSING: InternalIP
	// MISSING: NodeTypeID
	// MISSING: Version
	// MISSING: CustomCoreCount
	// MISSING: State
	return out
}
