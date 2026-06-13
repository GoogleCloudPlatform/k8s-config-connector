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

// ComputeMachineImage_v1alpha1_FromProto maps from pb.MachineImage to the root KRM type.
// We hand-code this function because the proto has a `Status` field (of type *string)
// which conflicts with the KRM type's custom status field (of type ComputeMachineImageStatus).
func ComputeMachineImage_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.MachineImage) *krm.ComputeMachineImage {
	if in == nil {
		return nil
	}
	out := &krm.ComputeMachineImage{}
	out.Spec = *ComputeMachineImageSpec_v1alpha1_FromProto(mapCtx, in)
	out.Status = *ComputeMachineImageStatus_v1alpha1_FromProto(mapCtx, in)
	return out
}

// ComputeMachineImage_v1alpha1_ToProto maps from KRM type to pb.MachineImage.
// We hand-code this function because the proto has a `Status` field (of type *string)
// which conflicts with the KRM type's custom status field (of type ComputeMachineImageStatus).
func ComputeMachineImage_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeMachineImage) *pb.MachineImage {
	if in == nil {
		return nil
	}
	out := ComputeMachineImageSpec_v1alpha1_ToProto(mapCtx, &in.Spec)
	if out == nil {
		out = &pb.MachineImage{}
	}
	return out
}

// ComputeMachineImageStatus_v1alpha1_FromProto maps from pb.MachineImage to ComputeMachineImageStatus.
// We hand-code this because it cannot be auto-generated due to status field name collision.
func ComputeMachineImageStatus_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.MachineImage) *krm.ComputeMachineImageStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeMachineImageStatus{}
	out.SelfLink = in.SelfLink
	out.StorageLocations = in.StorageLocations
	return out
}

// ComputeMachineImageStatus_v1alpha1_ToProto maps from ComputeMachineImageStatus to pb.MachineImage.
// We hand-code this because it cannot be auto-generated due to status field name collision.
func ComputeMachineImageStatus_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeMachineImageStatus) *pb.MachineImage {
	if in == nil {
		return nil
	}
	out := &pb.MachineImage{}
	out.SelfLink = in.SelfLink
	out.StorageLocations = in.StorageLocations
	return out
}
