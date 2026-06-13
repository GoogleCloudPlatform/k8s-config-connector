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

// +generated:mapper
// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.compute.v1

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeInstanceGroupNamedPortSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.NamedPort) *krm.ComputeInstanceGroupNamedPortSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeInstanceGroupNamedPortSpec{}
	out.Port = int64(in.GetPort())
	out.ResourceID = direct.LazyPtr(in.GetName())
	return out
}

func ComputeInstanceGroupNamedPortSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInstanceGroupNamedPortSpec) *pb.NamedPort {
	if in == nil {
		return nil
	}
	out := &pb.NamedPort{}
	out.Port = direct.LazyPtr(int32(in.Port))
	out.Name = in.ResourceID
	return out
}

func ComputeInstanceGroupNamedPortObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.NamedPort) *krm.ComputeInstanceGroupNamedPortObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ComputeInstanceGroupNamedPortObservedState{}
	return out
}

func ComputeInstanceGroupNamedPortObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInstanceGroupNamedPortObservedState) *pb.NamedPort {
	if in == nil {
		return nil
	}
	out := &pb.NamedPort{}
	return out
}
