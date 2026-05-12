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

func ReservationSpecificReservation_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AllocationSpecificSKUReservation) *krm.ReservationSpecificReservation {
	if in == nil {
		return nil
	}
	out := &krm.ReservationSpecificReservation{}
	out.Count = Int32_FromProto(in.Count)
	out.InUseCount = Int32_FromProto(in.InUseCount)
	out.InstanceProperties = ReservationInstanceProperties_v1beta1_FromProto(mapCtx, in.GetInstanceProperties())
	return out
}

func ReservationSpecificReservation_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ReservationSpecificReservation) *pb.AllocationSpecificSKUReservation {
	if in == nil {
		return nil
	}
	out := &pb.AllocationSpecificSKUReservation{}
	out.Count = Int32_ToProto(in.Count)
	out.InUseCount = Int32_ToProto(in.InUseCount)
	out.InstanceProperties = ReservationInstanceProperties_v1beta1_ToProto(mapCtx, in.InstanceProperties)
	return out
}

func ReservationLocalSsds_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDisk) *krm.ReservationLocalSsds {
	if in == nil {
		return nil
	}
	out := &krm.ReservationLocalSsds{}
	out.DiskSizeGb = Int32_FromProto(in.DiskSizeGb)
	out.Interface = in.Interface
	return out
}

func ReservationLocalSsds_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ReservationLocalSsds) *pb.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDisk {
	if in == nil {
		return nil
	}
	out := &pb.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDisk{}
	out.DiskSizeGb = Int32_ToProto(in.DiskSizeGb)
	out.Interface = in.Interface
	return out
}

/*
func ComputeReservationObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Reservation) *krm.ComputeReservationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ComputeReservationObservedState{}
	out.ID = Uint64_FromProto(in.Id)
	out.ResourceStatus = AllocationResourceStatus_v1beta1_FromProto(mapCtx, in.GetResourceStatus())
	return out
}

func ComputeReservationObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeReservationObservedState) *pb.Reservation {
	if in == nil {
		return nil
	}
	out := &pb.Reservation{}
	out.Id = Uint64_ToProto(in.ID)
	out.ResourceStatus = AllocationResourceStatus_v1beta1_ToProto(mapCtx, in.ResourceStatus)
	return out
}
*/

func Int32_FromProto(in *int64) *int32 {
	if in == nil {
		return nil
	}
	out := int32(*in)
	return &out
}

func Int32_ToProto(in *int32) *int64 {
	if in == nil {
		return nil
	}
	out := int64(*in)
	return &out
}

func Uint64_FromProto(in *uint64) *int64 {
	if in == nil {
		return nil
	}
	out := int64(*in)
	return &out
}

func Uint64_ToProto(in *int64) *uint64 {
	if in == nil {
		return nil
	}
	out := uint64(*in)
	return &out
}

func ComputeReservationStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Reservation) *krm.ComputeReservationStatus {
	if in == nil {
		return nil
	}
	status := &krm.ComputeReservationStatus{}
	if in.Commitment != nil {
		status.Commitment = in.Commitment
	}
	if in.CreationTimestamp != nil {
		status.CreationTimestamp = in.CreationTimestamp
	}
	if in.SelfLink != nil {
		status.SelfLink = in.SelfLink
	}
	if in.Status != nil {
		status.Status = in.Status
	}
	return status
}

func ComputeReservationStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeReservationStatus) *pb.Reservation {
	if in == nil {
		return nil
	}
	out := &pb.Reservation{}
	if in.Commitment != nil {
		out.Commitment = in.Commitment
	}
	if in.CreationTimestamp != nil {
		out.CreationTimestamp = in.CreationTimestamp
	}
	if in.SelfLink != nil {
		out.SelfLink = in.SelfLink
	}
	if in.Status != nil {
		out.Status = in.Status
	}
	return out
}
