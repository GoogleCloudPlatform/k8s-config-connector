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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ReservationSpecificReservation_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AllocationSpecificSKUReservation) *krm.ReservationSpecificReservation {
	if in == nil {
		return nil
	}
	out := &krm.ReservationSpecificReservation{}
	out.Count = direct.PtrInt64ToPtrInt32(in.Count)
	out.InUseCount = direct.PtrInt64ToPtrInt32(in.InUseCount)
	out.InstanceProperties = ReservationInstanceProperties_v1beta1_FromProto(mapCtx, in.GetInstanceProperties())
	return out
}

func ReservationSpecificReservation_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ReservationSpecificReservation) *pb.AllocationSpecificSKUReservation {
	if in == nil {
		return nil
	}
	out := &pb.AllocationSpecificSKUReservation{}
	out.Count = direct.PtrInt32ToPtrInt64(in.Count)
	out.InUseCount = direct.PtrInt32ToPtrInt64(in.InUseCount)
	out.InstanceProperties = ReservationInstanceProperties_v1beta1_ToProto(mapCtx, in.InstanceProperties)
	return out
}

func ReservationLocalSsds_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDisk) *krm.ReservationLocalSsds {
	if in == nil {
		return nil
	}
	out := &krm.ReservationLocalSsds{}
	out.DiskSizeGb = direct.PtrInt64ToPtrInt32(in.DiskSizeGb)
	out.Interface = in.Interface
	return out
}

func ReservationLocalSsds_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ReservationLocalSsds) *pb.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDisk {
	if in == nil {
		return nil
	}
	out := &pb.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDisk{}
	out.DiskSizeGb = direct.PtrInt32ToPtrInt64(in.DiskSizeGb)
	out.Interface = in.Interface
	return out
}

/*
func ComputeReservationObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Reservation) *krm.ComputeReservationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ComputeReservationObservedState{}
	out.ID = direct.PtrUint64ToPtrInt64(in.Id)
	out.ResourceStatus = AllocationResourceStatus_v1beta1_FromProto(mapCtx, in.GetResourceStatus())
	return out
}

func ComputeReservationObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeReservationObservedState) *pb.Reservation {
	if in == nil {
		return nil
	}
	out := &pb.Reservation{}
	out.Id = direct.PtrInt64ToPtrUint64(in.ID)
	out.ResourceStatus = AllocationResourceStatus_v1beta1_ToProto(mapCtx, in.ResourceStatus)
	return out
}
*/

func ReservationInstanceProperties_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AllocationSpecificSKUAllocationReservedInstanceProperties) *krm.ReservationInstanceProperties {
	if in == nil {
		return nil
	}
	out := &krm.ReservationInstanceProperties{}
	out.GuestAccelerators = direct.Slice_FromProto(mapCtx, in.GuestAccelerators, ReservationGuestAccelerators_v1beta1_FromProto)
	out.LocalSsds = direct.Slice_FromProto(mapCtx, in.LocalSsds, ReservationLocalSsds_v1beta1_FromProto)
	out.MachineType = in.MachineType
	out.MinCpuPlatform = in.MinCpuPlatform
	return out
}

func ReservationInstanceProperties_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ReservationInstanceProperties) *pb.AllocationSpecificSKUAllocationReservedInstanceProperties {
	if in == nil {
		return nil
	}
	out := &pb.AllocationSpecificSKUAllocationReservedInstanceProperties{}
	out.GuestAccelerators = direct.Slice_ToProto(mapCtx, in.GuestAccelerators, ReservationGuestAccelerators_v1beta1_ToProto)
	out.LocalSsds = direct.Slice_ToProto(mapCtx, in.LocalSsds, ReservationLocalSsds_v1beta1_ToProto)
	out.MachineType = in.MachineType
	out.MinCpuPlatform = in.MinCpuPlatform
	return out
}

func ComputeReservationStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Reservation) *krm.ComputeReservationStatus {
	if in == nil {
		return nil
	}
	status := &krm.ComputeReservationStatus{}
	status.Commitment = in.Commitment
	status.CreationTimestamp = in.CreationTimestamp
	status.SelfLink = in.SelfLink
	status.Status = in.Status
	return status
}

func ComputeReservationStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeReservationStatus) *pb.Reservation {
	if in == nil {
		return nil
	}
	out := &pb.Reservation{}
	out.Commitment = in.Commitment
	out.CreationTimestamp = in.CreationTimestamp
	out.SelfLink = in.SelfLink
	out.Status = in.Status
	return out
}

func ShareSettings_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ShareSettings) *krm.ShareSettings {
	if in == nil {
		return nil
	}
	out := &krm.ShareSettings{}
	if in.ProjectMap != nil {
		for k, v := range in.ProjectMap {
			out.ProjectMap = append(out.ProjectMap, krm.ShareSettingsProjectMap{
				KeyRef: &refsv1beta1.ExtendedProjectRef{
					External: k,
				},
				Value: ShareSettingsProjectConfig_v1beta1_FromProto(mapCtx, v),
			})
		}
	}
	out.ShareType = in.ShareType
	return out
}
func ShareSettings_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ShareSettings) *pb.ShareSettings {
	if in == nil {
		return nil
	}
	out := &pb.ShareSettings{}
	if in.ProjectMap != nil {
		out.ProjectMap = make(map[string]*pb.ShareSettingsProjectConfig)
		for _, entry := range in.ProjectMap {
			if entry.KeyRef != nil {
				if entry.KeyRef.External == "" {
					mapCtx.Errorf("reference %s was not pre-resolved", entry.KeyRef.Name)
					continue
				}
			}
			out.ProjectMap[entry.KeyRef.External] = ShareSettingsProjectConfig_v1beta1_ToProto(mapCtx, entry.Value)
		}
	}
	out.ShareType = in.ShareType
	return out
}
func ShareSettingsProjectConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ShareSettingsProjectConfig) *krm.ShareSettingsProjectConfig {
	if in == nil {
		return nil
	}
	out := &krm.ShareSettingsProjectConfig{}
	if in.GetProjectId() != "" {
		out.ProjectIDRef = &refsv1beta1.ProjectRef{External: in.GetProjectId()}
	}
	return out
}
func ShareSettingsProjectConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ShareSettingsProjectConfig) *pb.ShareSettingsProjectConfig {
	if in == nil {
		return nil
	}
	out := &pb.ShareSettingsProjectConfig{}
	if in.ProjectIDRef != nil {
		if in.ProjectIDRef.External == "" {
			mapCtx.Errorf("reference %s was not pre-resolved", in.ProjectIDRef.Name)
		}
		out.ProjectId = &in.ProjectIDRef.External
	}
	return out
}
func ShareSettingsObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ShareSettings) *krm.ShareSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ShareSettingsObservedState{}
	if in.ProjectMap != nil {
		out.ProjectMap = make(map[string]krm.ShareSettingsProjectConfigObservedState)
		for k, v := range in.ProjectMap {
			val := ShareSettingsProjectConfigObservedState_v1beta1_FromProto(mapCtx, v)
			if val != nil {
				out.ProjectMap[k] = *val
			}
		}
	}
	out.ShareType = in.ShareType
	return out
}
func ShareSettingsObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ShareSettingsObservedState) *pb.ShareSettings {
	if in == nil {
		return nil
	}
	out := &pb.ShareSettings{}
	if in.ProjectMap != nil {
		out.ProjectMap = make(map[string]*pb.ShareSettingsProjectConfig)
		for k, v := range in.ProjectMap {
			out.ProjectMap[k] = ShareSettingsProjectConfigObservedState_v1beta1_ToProto(mapCtx, &v)
		}
	}
	out.ShareType = in.ShareType
	return out
}
func ShareSettingsProjectConfigObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ShareSettingsProjectConfig) *krm.ShareSettingsProjectConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ShareSettingsProjectConfigObservedState{}
	out.ProjectID = in.ProjectId
	return out
}
func ShareSettingsProjectConfigObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ShareSettingsProjectConfigObservedState) *pb.ShareSettingsProjectConfig {
	if in == nil {
		return nil
	}
	out := &pb.ShareSettingsProjectConfig{}
	out.ProjectId = in.ProjectID
	return out
}
