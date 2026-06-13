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
	return nil
}

func ComputeInstanceStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInstanceStatus) *pb.Instance {
	return nil
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

func InstanceInitializeParams_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AttachedDiskInitializeParams) *krm.InstanceInitializeParams {
	if in == nil {
		return nil
	}
	out := &krm.InstanceInitializeParams{}
	out.Size = in.DiskSizeGb
	out.Type = in.DiskType
	if in.SourceImage != nil {
		out.SourceDiskRef = &krm.InstanceResourceRef{External: *in.SourceImage}
	}
	return out
}

func InstanceInitializeParams_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceInitializeParams) *pb.AttachedDiskInitializeParams {
	if in == nil {
		return nil
	}
	out := &pb.AttachedDiskInitializeParams{}
	out.DiskSizeGb = in.Size
	out.DiskType = in.Type
	if in.SourceDiskRef != nil {
		out.SourceImage = &in.SourceDiskRef.External
	}
	return out
}

func InstanceIpv6AccessConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AccessConfig) *krm.InstanceIpv6AccessConfig {
	if in == nil {
		return nil
	}
	out := &krm.InstanceIpv6AccessConfig{}
	out.ExternalIpv6 = in.ExternalIpv6
	if in.ExternalIpv6PrefixLength != nil {
		prefixStr := strconv.FormatInt(int64(*in.ExternalIpv6PrefixLength), 10)
		out.ExternalIpv6PrefixLength = &prefixStr
	}
	out.Name = in.Name
	if in.NetworkTier != nil {
		out.NetworkTier = *in.NetworkTier
	}
	out.PublicPtrDomainName = in.PublicPtrDomainName
	return out
}

func InstanceIpv6AccessConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceIpv6AccessConfig) *pb.AccessConfig {
	if in == nil {
		return nil
	}
	out := &pb.AccessConfig{}
	out.ExternalIpv6 = in.ExternalIpv6
	if in.ExternalIpv6PrefixLength != nil {
		if val, err := strconv.ParseInt(*in.ExternalIpv6PrefixLength, 10, 32); err == nil {
			val32 := int32(val)
			out.ExternalIpv6PrefixLength = &val32
		}
	}
	out.Name = in.Name
	if in.NetworkTier != "" {
		out.NetworkTier = &in.NetworkTier
	}
	out.PublicPtrDomainName = in.PublicPtrDomainName
	return out
}

func InstanceLocalSsdRecoveryTimeout_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Duration) *krm.InstanceLocalSsdRecoveryTimeout {
	if in == nil {
		return nil
	}
	out := &krm.InstanceLocalSsdRecoveryTimeout{}
	if in.Nanos != nil {
		val := int64(*in.Nanos)
		out.Nanos = &val
	}
	if in.Seconds != nil {
		out.Seconds = *in.Seconds
	}
	return out
}

func InstanceLocalSsdRecoveryTimeout_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceLocalSsdRecoveryTimeout) *pb.Duration {
	if in == nil {
		return nil
	}
	out := &pb.Duration{}
	if in.Nanos != nil {
		val := int32(*in.Nanos)
		out.Nanos = &val
	}
	out.Seconds = &in.Seconds
	return out
}

func InstanceMaxRunDuration_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Duration) *krm.InstanceMaxRunDuration {
	if in == nil {
		return nil
	}
	out := &krm.InstanceMaxRunDuration{}
	if in.Nanos != nil {
		val := int64(*in.Nanos)
		out.Nanos = &val
	}
	if in.Seconds != nil {
		out.Seconds = *in.Seconds
	}
	return out
}

func InstanceMaxRunDuration_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceMaxRunDuration) *pb.Duration {
	if in == nil {
		return nil
	}
	out := &pb.Duration{}
	if in.Nanos != nil {
		val := int32(*in.Nanos)
		out.Nanos = &val
	}
	out.Seconds = &in.Seconds
	return out
}

func InstanceNetworkInterface_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.NetworkInterface) *krm.InstanceNetworkInterface {
	if in == nil {
		return nil
	}
	out := &krm.InstanceNetworkInterface{}
	out.AccessConfig = direct.Slice_FromProto(mapCtx, in.AccessConfigs, InstanceAccessConfig_v1beta1_FromProto)
	out.AliasIpRange = direct.Slice_FromProto(mapCtx, in.AliasIpRanges, InstanceAliasIpRange_v1beta1_FromProto)
	if in.InternalIpv6PrefixLength != nil {
		val := int64(*in.InternalIpv6PrefixLength)
		out.InternalIpv6PrefixLength = &val
	}
	out.Ipv6AccessConfig = direct.Slice_FromProto(mapCtx, in.Ipv6AccessConfigs, InstanceIpv6AccessConfig_v1beta1_FromProto)
	out.Ipv6AccessType = in.Ipv6AccessType
	out.Ipv6Address = in.Ipv6Address
	out.Name = in.Name
	out.NetworkIp = in.NetworkIP
	if in.GetNetwork() != "" {
		out.NetworkRef = &krm.InstanceResourceRef{External: in.GetNetwork()}
	}
	out.NicType = in.NicType
	if in.QueueCount != nil {
		val := int64(*in.QueueCount)
		out.QueueCount = &val
	}
	out.StackType = in.StackType
	if in.GetSubnetwork() != "" {
		out.SubnetworkRef = &krm.InstanceResourceRef{External: in.GetSubnetwork()}
	}
	return out
}

func InstanceNetworkInterface_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceNetworkInterface) *pb.NetworkInterface {
	if in == nil {
		return nil
	}
	out := &pb.NetworkInterface{}
	out.AccessConfigs = direct.Slice_ToProto(mapCtx, in.AccessConfig, InstanceAccessConfig_v1beta1_ToProto)
	out.AliasIpRanges = direct.Slice_ToProto(mapCtx, in.AliasIpRange, InstanceAliasIpRange_v1beta1_ToProto)
	if in.InternalIpv6PrefixLength != nil {
		val := int32(*in.InternalIpv6PrefixLength)
		out.InternalIpv6PrefixLength = &val
	}
	out.Ipv6AccessConfigs = direct.Slice_ToProto(mapCtx, in.Ipv6AccessConfig, InstanceIpv6AccessConfig_v1beta1_ToProto)
	out.Ipv6AccessType = in.Ipv6AccessType
	out.Ipv6Address = in.Ipv6Address
	out.Name = in.Name
	out.NetworkIP = in.NetworkIp
	if in.NetworkRef != nil {
		out.Network = &in.NetworkRef.External
	}
	out.NicType = in.NicType
	if in.QueueCount != nil {
		val := int32(*in.QueueCount)
		out.QueueCount = &val
	}
	out.StackType = in.StackType
	if in.SubnetworkRef != nil {
		out.Subnetwork = &in.SubnetworkRef.External
	}
	return out
}

func InstanceScratchDisk_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AttachedDisk) *krm.InstanceScratchDisk {
	if in == nil {
		return nil
	}
	out := &krm.InstanceScratchDisk{}
	if in.Interface != nil {
		out.Interface = *in.Interface
	}
	out.Size = in.DiskSizeGb
	return out
}

func InstanceScratchDisk_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceScratchDisk) *pb.AttachedDisk {
	if in == nil {
		return nil
	}
	out := &pb.AttachedDisk{}
	if in.Interface != "" {
		out.Interface = &in.Interface
	}
	out.DiskSizeGb = in.Size
	out.Type = direct.PtrTo("SCRATCH")
	return out
}
