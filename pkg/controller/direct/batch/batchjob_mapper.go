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

// +generated:mapper
// krm.group: batch.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.batch.v1

package batch

import (
	pb "cloud.google.com/go/batch/apiv1/batchpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/batch/v1alpha1"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AllocationPolicy_FromProto(mapCtx *direct.MapContext, in *pb.AllocationPolicy) *krm.AllocationPolicy {
	if in == nil {
		return nil
	}
	out := &krm.AllocationPolicy{}
	out.Location = AllocationPolicy_LocationPolicy_FromProto(mapCtx, in.GetLocation())
	out.Instances = direct.Slice_FromProto(mapCtx, in.Instances, AllocationPolicy_InstancePolicyOrTemplate_FromProto)

	if in.GetServiceAccount() != nil {
		out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetServiceAccount().Email}
	}
	out.Labels = in.Labels
	out.Network = AllocationPolicy_NetworkPolicy_FromProto(mapCtx, in.GetNetwork())
	out.Placement = AllocationPolicy_PlacementPolicy_FromProto(mapCtx, in.GetPlacement())
	out.Tags = in.Tags
	return out
}
func AllocationPolicy_ToProto(mapCtx *direct.MapContext, in *krm.AllocationPolicy) *pb.AllocationPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AllocationPolicy{}
	out.Location = AllocationPolicy_LocationPolicy_ToProto(mapCtx, in.Location)
	out.Instances = direct.Slice_ToProto(mapCtx, in.Instances, AllocationPolicy_InstancePolicyOrTemplate_ToProto)
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = &pb.ServiceAccount{
			Email: in.ServiceAccountRef.External,
		}
	}
	out.Labels = in.Labels
	out.Network = AllocationPolicy_NetworkPolicy_ToProto(mapCtx, in.Network)
	out.Placement = AllocationPolicy_PlacementPolicy_ToProto(mapCtx, in.Placement)
	out.Tags = in.Tags
	return out
}
func AllocationPolicy_Disk_FromProto(mapCtx *direct.MapContext, in *pb.AllocationPolicy_Disk) *krm.AllocationPolicy_Disk {
	if in == nil {
		return nil
	}
	out := &krm.AllocationPolicy_Disk{}
	if in.GetImage() != "" {
		out.ImageRef = &computev1beta1.ComputeImageRef{External: in.GetImage()}
	}
	out.Snapshot = direct.LazyPtr(in.GetSnapshot())
	out.Type = direct.LazyPtr(in.GetType())
	out.SizeGB = direct.LazyPtr(in.GetSizeGb())
	out.DiskInterface = direct.LazyPtr(in.GetDiskInterface())
	return out
}

func AllocationPolicy_Disk_ToProto(mapCtx *direct.MapContext, in *krm.AllocationPolicy_Disk) *pb.AllocationPolicy_Disk {
	if in == nil {
		return nil
	}
	out := &pb.AllocationPolicy_Disk{}
	if in.ImageRef != nil {
		out.DataSource = &pb.AllocationPolicy_Disk_Image{Image: in.ImageRef.External}
	}
	if in.Snapshot != nil {
		out.DataSource = &pb.AllocationPolicy_Disk_Snapshot{Snapshot: *in.Snapshot}
	}
	out.Type = direct.ValueOf(in.Type)
	out.SizeGb = direct.ValueOf(in.SizeGB)
	out.DiskInterface = direct.ValueOf(in.DiskInterface)
	return out
}
