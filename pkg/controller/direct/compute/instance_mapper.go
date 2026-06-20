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
	"strconv"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeInstanceSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.ComputeInstanceSpec {
	if in == nil {
		return nil
	}
	if in.Tags != nil && len(in.Tags.Items) == 0 {
		in.Tags = nil
	}
	if in.Metadata != nil && len(in.Metadata.Items) == 0 {
		in.Metadata = nil
	}
	out := &krm.ComputeInstanceSpec{}
	out.AdvancedMachineFeatures = InstanceAdvancedMachineFeatures_v1beta1_FromProto(mapCtx, in.GetAdvancedMachineFeatures())
	out.ConfidentialInstanceConfig = InstanceConfidentialInstanceConfig_v1beta1_FromProto(mapCtx, in.GetConfidentialInstanceConfig())
	out.DeletionProtection = in.DeletionProtection
	out.Description = in.Description
	out.Hostname = in.Hostname
	out.MachineType = in.MachineType
	out.Metadata = InstanceMetadata_v1beta1_FromProto(mapCtx, in.GetMetadata())
	out.NetworkPerformanceConfig = InstanceNetworkPerformanceConfig_v1beta1_FromProto(mapCtx, in.GetNetworkPerformanceConfig())
	out.Params = InstanceParams_v1beta1_FromProto(mapCtx, in.GetParams())
	out.ReservationAffinity = InstanceReservationAffinity_v1beta1_FromProto(mapCtx, in.GetReservationAffinity())
	out.ResourcePolicies = ComputeInstanceSpec_ResourcePolicies_FromProto(mapCtx, in.ResourcePolicies)
	out.Scheduling = InstanceScheduling_v1beta1_FromProto(mapCtx, in.GetScheduling())
	out.ShieldedInstanceConfig = InstanceShieldedInstanceConfig_v1beta1_FromProto(mapCtx, in.GetShieldedInstanceConfig())
	if v := in.GetTags(); v != nil {
		out.Tags = v.GetItems()
	}
	out.Zone = in.Zone
	return out
}

func ComputeInstanceSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.AdvancedMachineFeatures = InstanceAdvancedMachineFeatures_v1beta1_ToProto(mapCtx, in.AdvancedMachineFeatures)
	out.ConfidentialInstanceConfig = InstanceConfidentialInstanceConfig_v1beta1_ToProto(mapCtx, in.ConfidentialInstanceConfig)
	out.DeletionProtection = in.DeletionProtection
	out.Description = in.Description
	out.Hostname = in.Hostname
	out.MachineType = in.MachineType
	out.Metadata = InstanceMetadata_v1beta1_ToProto(mapCtx, in.Metadata)
	out.NetworkPerformanceConfig = InstanceNetworkPerformanceConfig_v1beta1_ToProto(mapCtx, in.NetworkPerformanceConfig)
	out.Params = InstanceParams_v1beta1_ToProto(mapCtx, in.Params)
	out.ReservationAffinity = InstanceReservationAffinity_v1beta1_ToProto(mapCtx, in.ReservationAffinity)
	out.ResourcePolicies = ComputeInstanceSpec_ResourcePolicies_ToProto(mapCtx, in.ResourcePolicies)
	out.Scheduling = InstanceScheduling_v1beta1_ToProto(mapCtx, in.Scheduling)
	out.ShieldedInstanceConfig = InstanceShieldedInstanceConfig_v1beta1_ToProto(mapCtx, in.ShieldedInstanceConfig)
	if len(in.Tags) > 0 {
		out.Tags = &pb.Tags{
			Items: in.Tags,
		}
	}
	out.Zone = in.Zone
	return out
}

func ComputeInstanceStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.ComputeInstanceStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeInstanceStatus{}
	out.CpuPlatform = in.CpuPlatform
	out.CurrentStatus = in.Status
	if in.Id != nil {
		idStr := strconv.FormatUint(*in.Id, 10)
		out.InstanceId = &idStr
	}
	out.LabelFingerprint = in.LabelFingerprint
	if in.Metadata != nil {
		out.MetadataFingerprint = in.Metadata.Fingerprint
	}
	out.SelfLink = in.SelfLink
	if in.Tags != nil {
		out.TagsFingerprint = in.Tags.Fingerprint
	}
	return out
}

func ComputeInstanceStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInstanceStatus) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.CpuPlatform = in.CpuPlatform
	out.Status = in.CurrentStatus
	if in.InstanceId != nil {
		idVal, err := strconv.ParseUint(*in.InstanceId, 10, 64)
		if err != nil {
			mapCtx.Errorf("error converting InstanceId string %s to uint64: %v", *in.InstanceId, err)
		} else {
			out.Id = &idVal
		}
	}
	out.LabelFingerprint = in.LabelFingerprint
	if in.MetadataFingerprint != nil {
		if out.Metadata == nil {
			out.Metadata = &pb.Metadata{}
		}
		out.Metadata.Fingerprint = in.MetadataFingerprint
	}
	out.SelfLink = in.SelfLink
	if in.TagsFingerprint != nil {
		if out.Tags == nil {
			out.Tags = &pb.Tags{}
		}
		out.Tags.Fingerprint = in.TagsFingerprint
	}
	return out
}

func ComputeInstanceSpec_ResourcePolicies_FromProto(mapCtx *direct.MapContext, in []string) []krm.InstanceResourceRef {
	if in == nil {
		return nil
	}
	var out []krm.InstanceResourceRef
	for _, i := range in {
		out = append(out, krm.InstanceResourceRef{
			External: i,
		})
	}
	return out
}

func ComputeInstanceSpec_ResourcePolicies_ToProto(mapCtx *direct.MapContext, in []krm.InstanceResourceRef) []string {
	if in == nil {
		return nil
	}
	var out []string
	for _, i := range in {
		if i.External == "" {
			mapCtx.Errorf("reference %s was not pre-resolved", i.Name)
		}
		out = append(out, i.External)
	}
	return out
}

func InstanceMetadata_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Metadata) []krm.InstanceMetadata {
	if in == nil {
		return nil
	}
	var out []krm.InstanceMetadata
	for _, item := range in.Items {
		if item == nil {
			continue
		}
		if item.Key == nil {
			item.Key = direct.PtrTo("")
		}
		if item.Value == nil {
			item.Value = direct.PtrTo("")
		}
		out = append(out, krm.InstanceMetadata{
			Key:   item.GetKey(),
			Value: item.GetValue(),
		})
	}
	return out
}

func InstanceMetadata_v1beta1_ToProto(mapCtx *direct.MapContext, in []krm.InstanceMetadata) *pb.Metadata {
	if in == nil {
		return nil
	}
	out := &pb.Metadata{}
	for _, item := range in {
		out.Items = append(out.Items, &pb.Items{
			Key:   direct.PtrTo(item.Key),
			Value: direct.PtrTo(item.Value),
		})
	}
	return out
}

func InstanceAdvancedMachineFeatures_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AdvancedMachineFeatures) *krm.InstanceAdvancedMachineFeatures {
	return nil
}

func InstanceAdvancedMachineFeatures_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceAdvancedMachineFeatures) *pb.AdvancedMachineFeatures {
	return nil
}

func InstanceConfidentialInstanceConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ConfidentialInstanceConfig) *krm.InstanceConfidentialInstanceConfig {
	return nil
}

func InstanceConfidentialInstanceConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceConfidentialInstanceConfig) *pb.ConfidentialInstanceConfig {
	return nil
}

func InstanceNetworkPerformanceConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPerformanceConfig) *krm.InstanceNetworkPerformanceConfig {
	return nil
}

func InstanceNetworkPerformanceConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceNetworkPerformanceConfig) *pb.NetworkPerformanceConfig {
	return nil
}

func InstanceParams_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceParams) *krm.InstanceParams {
	return nil
}

func InstanceParams_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceParams) *pb.InstanceParams {
	return nil
}

func InstanceReservationAffinity_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ReservationAffinity) *krm.InstanceReservationAffinity {
	return nil
}

func InstanceReservationAffinity_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceReservationAffinity) *pb.ReservationAffinity {
	return nil
}

func InstanceScheduling_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Scheduling) *krm.InstanceScheduling {
	return nil
}

func InstanceScheduling_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceScheduling) *pb.Scheduling {
	return nil
}

func InstanceShieldedInstanceConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ShieldedInstanceConfig) *krm.InstanceShieldedInstanceConfig {
	return nil
}

func InstanceShieldedInstanceConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceShieldedInstanceConfig) *pb.ShieldedInstanceConfig {
	return nil
}
