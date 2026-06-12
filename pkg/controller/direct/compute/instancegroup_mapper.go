// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package compute

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	pb "google.golang.org/genproto/googleapis/cloud/compute/v1"
)

// ComputeInstanceGroupSpec_v1beta1_FromProto maps a pb.InstanceGroup to a krm.ComputeInstanceGroupSpec.
// Note: Instances field is handled manually in the controller (via separate API calls like addInstances / removeInstances),
// as it is not part of the base pb.InstanceGroup message.
func ComputeInstanceGroupSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroup) *krm.ComputeInstanceGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeInstanceGroupSpec{}
	out.Description = in.Description
	if in.Name != nil {
		out.ResourceID = in.Name
	}
	if in.Network != nil {
		out.NetworkRef = &krm.ComputeNetworkRef{External: *in.Network}
	}
	if in.Zone != nil {
		out.Zone = *in.Zone
	}
	if in.NamedPorts != nil {
		out.NamedPort = make([]krm.InstancegroupNamedPort, len(in.NamedPorts))
		for i, p := range in.NamedPorts {
			if p.Name != nil {
				out.NamedPort[i].Name = *p.Name
			}
			if p.Port != nil {
				out.NamedPort[i].Port = int(*p.Port)
			}
		}
	}
	return out
}

// ComputeInstanceGroupSpec_v1beta1_ToProto maps a krm.ComputeInstanceGroupSpec to a pb.InstanceGroup.
// Note: Instances field is handled manually in the controller (via separate API calls like addInstances / removeInstances),
// as it is not part of the base pb.InstanceGroup message.
func ComputeInstanceGroupSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInstanceGroupSpec) *pb.InstanceGroup {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroup{}
	out.Description = in.Description
	if in.ResourceID != nil {
		out.Name = in.ResourceID
	}
	if in.NetworkRef != nil {
		out.Network = &in.NetworkRef.External
	}
	if in.Zone != "" {
		out.Zone = &in.Zone
	}
	if in.NamedPort != nil {
		out.NamedPorts = make([]*pb.NamedPort, len(in.NamedPort))
		for i, p := range in.NamedPort {
			name := p.Name
			port := int32(p.Port)
			out.NamedPorts[i] = &pb.NamedPort{
				Name: &name,
				Port: &port,
			}
		}
	}
	return out
}

// ComputeInstanceGroupStatus_v1beta1_FromProto maps a pb.InstanceGroup to a krm.ComputeInstanceGroupStatus.
// Note: ObservedGeneration is managed by the Kubernetes controller and is not part of the GCP API proto.
func ComputeInstanceGroupStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroup) *krm.ComputeInstanceGroupStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeInstanceGroupStatus{}
	out.SelfLink = in.SelfLink
	if in.Size != nil {
		sizeVal := int(*in.Size)
		out.Size = &sizeVal
	}
	return out
}

// ComputeInstanceGroupStatus_v1beta1_ToProto maps a krm.ComputeInstanceGroupStatus to a pb.InstanceGroup.
// Note: ObservedGeneration is managed by the Kubernetes controller and is not part of the GCP API proto.
func ComputeInstanceGroupStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInstanceGroupStatus) *pb.InstanceGroup {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroup{}
	out.SelfLink = in.SelfLink
	if in.Size != nil {
		sizeVal := int32(*in.Size)
		out.Size = &sizeVal
	}
	return out
}

// InstancegroupNamedPort_v1beta1_FromProto maps a pb.NamedPort to a krm.InstancegroupNamedPort.
func InstancegroupNamedPort_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.NamedPort) *krm.InstancegroupNamedPort {
	if in == nil {
		return nil
	}
	out := &krm.InstancegroupNamedPort{}
	if in.Name != nil {
		out.Name = *in.Name
	}
	if in.Port != nil {
		out.Port = int(*in.Port)
	}
	return out
}

// InstancegroupNamedPort_v1beta1_ToProto maps a krm.InstancegroupNamedPort to a pb.NamedPort.
func InstancegroupNamedPort_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstancegroupNamedPort) *pb.NamedPort {
	if in == nil {
		return nil
	}
	out := &pb.NamedPort{}
	name := in.Name
	port := int32(in.Port)
	out.Name = &name
	out.Port = &port
	return out
}
