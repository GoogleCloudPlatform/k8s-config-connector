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

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// ComputeInstanceGroupSpec_v1beta1_FromProto is handcoded because the KRM's Zone field
// is a non-pointer string while the proto's Zone is a pointer string, and to skip generating it.
func ComputeInstanceGroupSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroup) *krm.ComputeInstanceGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeInstanceGroupSpec{}
	out.Description = in.Description
	out.NamedPorts = direct.Slice_FromProto(mapCtx, in.NamedPorts, InstanceGroupNamedPort_v1beta1_FromProto)
	if in.GetNetwork() != "" {
		out.NetworkRef = &krm.ComputeNetworkRef{External: in.GetNetwork()}
	}
	out.Zone = in.GetZone()
	return out
}

// ComputeInstanceGroupSpec_v1beta1_ToProto is handcoded because the KRM's Zone field
// is a non-pointer string while the proto's Zone is a pointer string, and to skip generating it.
func ComputeInstanceGroupSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInstanceGroupSpec) *pb.InstanceGroup {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroup{}
	out.Description = in.Description
	out.NamedPorts = direct.Slice_ToProto(mapCtx, in.NamedPorts, InstanceGroupNamedPort_v1beta1_ToProto)
	if in.NetworkRef != nil {
		out.Network = &in.NetworkRef.External
	}
	if in.Zone != "" {
		out.Zone = &in.Zone
	}
	return out
}

// InstanceGroupNamedPort_v1beta1_FromProto is handcoded because the fields of pb.NamedPort
// are pointer types in protobuf, but are non-pointer types in KRM.
func InstanceGroupNamedPort_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.NamedPort) *krm.InstanceGroupNamedPort {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupNamedPort{}
	out.Name = in.GetName()
	out.Port = in.GetPort()
	return out
}

// InstanceGroupNamedPort_v1beta1_ToProto is handcoded because the fields of pb.NamedPort
// are pointer types in protobuf, but are non-pointer types in KRM.
func InstanceGroupNamedPort_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupNamedPort) *pb.NamedPort {
	if in == nil {
		return nil
	}
	out := &pb.NamedPort{}
	if in.Name != "" {
		out.Name = &in.Name
	}
	if in.Port != 0 {
		out.Port = &in.Port
	}
	return out
}

func ComputeInstanceGroupStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroup) *krm.ComputeInstanceGroupStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeInstanceGroupStatus{}
	out.SelfLink = in.SelfLink
	if in.Size != nil {
		val := int64(*in.Size)
		out.Size = &val
	}
	return out
}

func ComputeInstanceGroupStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInstanceGroupStatus) *pb.InstanceGroup {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroup{}
	out.SelfLink = in.SelfLink
	if in.Size != nil {
		val := int32(*in.Size)
		out.Size = &val
	}
	return out
}
