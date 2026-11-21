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

// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.compute.v1

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeInstanceGroupManagerResizeRequestSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupManagerResizeRequest) *krm.ComputeInstanceGroupManagerResizeRequestSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeInstanceGroupManagerResizeRequestSpec{}
	out.Description = in.Description
	out.RequestedRunDuration = Duration_v1alpha1_FromProto(mapCtx, in.RequestedRunDuration)
	out.ResizeBy = in.ResizeBy
	return out
}

func ComputeInstanceGroupManagerResizeRequestSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInstanceGroupManagerResizeRequestSpec) *pb.InstanceGroupManagerResizeRequest {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupManagerResizeRequest{}
	out.Description = in.Description
	out.RequestedRunDuration = Duration_v1alpha1_ToProto(mapCtx, in.RequestedRunDuration)
	out.ResizeBy = in.ResizeBy
	return out
}

func ComputeInstanceGroupManagerResizeRequestObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupManagerResizeRequest) *krm.ComputeInstanceGroupManagerResizeRequestObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ComputeInstanceGroupManagerResizeRequestObservedState{}
	out.CreationTimestamp = in.CreationTimestamp
	out.ID = in.Id
	out.Kind = in.Kind
	out.SelfLink = in.SelfLink
	out.SelfLinkWithID = in.SelfLinkWithId
	out.State = in.State
	out.Status = InstanceGroupManagerResizeRequestStatus_v1alpha1_FromProto(mapCtx, in.Status)
	out.Zone = in.Zone
	return out
}

func ComputeInstanceGroupManagerResizeRequestObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInstanceGroupManagerResizeRequestObservedState) *pb.InstanceGroupManagerResizeRequest {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupManagerResizeRequest{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Id = in.ID
	out.Kind = in.Kind
	out.SelfLink = in.SelfLink
	out.SelfLinkWithId = in.SelfLinkWithID
	out.State = in.State
	out.Status = InstanceGroupManagerResizeRequestStatus_v1alpha1_ToProto(mapCtx, in.Status)
	out.Zone = in.Zone
	return out
}
