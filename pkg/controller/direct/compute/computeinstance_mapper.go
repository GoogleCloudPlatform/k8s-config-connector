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
// See the License for the_specific language governing permissions and
// limitations under the License.

package compute

import (
	"strconv"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
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
	out.CanIpForward = in.CanIpForward
	out.MinCpuPlatform = in.MinCpuPlatform

	// Map disks
	for _, pbDisk := range in.Disks {
		if pbDisk.GetBoot() {
			out.BootDisk = InstanceBootDisk_v1beta1_FromProto(mapCtx, pbDisk)
		} else {
			krmDisk := InstanceAttachedDisk_v1beta1_FromProto(mapCtx, pbDisk)
			if krmDisk != nil {
				out.AttachedDisk = append(out.AttachedDisk, *krmDisk)
			}
		}
	}

	// Map NetworkInterfaces
	for _, pbNi := range in.NetworkInterfaces {
		krmNi := InstanceNetworkInterface_v1beta1_FromProto(mapCtx, pbNi)
		if krmNi != nil {
			out.NetworkInterface = append(out.NetworkInterface, *krmNi)
		}
	}

	// Map ServiceAccounts
	if len(in.ServiceAccounts) > 0 {
		out.ServiceAccount = InstanceServiceAccount_v1beta1_FromProto(mapCtx, in.ServiceAccounts[0])
	}

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
	out.CanIpForward = in.CanIpForward
	out.MinCpuPlatform = in.MinCpuPlatform

	// Map disks
	var pbDisks []*pb.AttachedDisk
	if in.BootDisk != nil {
		pbBootDisk := InstanceBootDisk_v1beta1_ToProto(mapCtx, in.BootDisk)
		if pbBootDisk != nil {
			pbBootDisk.Boot = direct.PtrTo(true)
			pbDisks = append(pbDisks, pbBootDisk)
		}
	}
	for i := range in.AttachedDisk {
		pbDisk := InstanceAttachedDisk_v1beta1_ToProto(mapCtx, &in.AttachedDisk[i])
		if pbDisk != nil {
			pbDisk.Boot = direct.PtrTo(false)
			pbDisks = append(pbDisks, pbDisk)
		}
	}
	out.Disks = pbDisks

	// Map NetworkInterfaces
	for i := range in.NetworkInterface {
		pbNi := InstanceNetworkInterface_v1beta1_ToProto(mapCtx, &in.NetworkInterface[i])
		if pbNi != nil {
			out.NetworkInterfaces = append(out.NetworkInterfaces, pbNi)
		}
	}

	// Map ServiceAccounts
	if in.ServiceAccount != nil {
		pbSa := InstanceServiceAccount_v1beta1_ToProto(mapCtx, in.ServiceAccount)
		if pbSa != nil {
			out.ServiceAccounts = append(out.ServiceAccounts, pbSa)
		}
	}

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
	if in == nil {
		return nil
	}
	out := &krm.InstanceAdvancedMachineFeatures{}
	out.EnableNestedVirtualization = in.EnableNestedVirtualization
	if in.ThreadsPerCore != nil {
		v := int64(*in.ThreadsPerCore)
		out.ThreadsPerCore = &v
	}
	if in.VisibleCoreCount != nil {
		v := int64(*in.VisibleCoreCount)
		out.VisibleCoreCount = &v
	}
	return out
}

func InstanceAdvancedMachineFeatures_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceAdvancedMachineFeatures) *pb.AdvancedMachineFeatures {
	if in == nil {
		return nil
	}
	out := &pb.AdvancedMachineFeatures{}
	out.EnableNestedVirtualization = in.EnableNestedVirtualization
	if in.ThreadsPerCore != nil {
		v := int32(*in.ThreadsPerCore)
		out.ThreadsPerCore = &v
	}
	if in.VisibleCoreCount != nil {
		v := int32(*in.VisibleCoreCount)
		out.VisibleCoreCount = &v
	}
	return out
}

func InstanceConfidentialInstanceConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ConfidentialInstanceConfig) *krm.InstanceConfidentialInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.InstanceConfidentialInstanceConfig{}
	out.EnableConfidentialCompute = direct.ValueOf(in.EnableConfidentialCompute)
	return out
}

func InstanceConfidentialInstanceConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceConfidentialInstanceConfig) *pb.ConfidentialInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.ConfidentialInstanceConfig{}
	out.EnableConfidentialCompute = &in.EnableConfidentialCompute
	return out
}

func InstanceNetworkPerformanceConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPerformanceConfig) *krm.InstanceNetworkPerformanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.InstanceNetworkPerformanceConfig{}
	if in.TotalEgressBandwidthTier != nil {
		out.TotalEgressBandwidthTier = *in.TotalEgressBandwidthTier
	}
	return out
}

func InstanceNetworkPerformanceConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceNetworkPerformanceConfig) *pb.NetworkPerformanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkPerformanceConfig{}
	out.TotalEgressBandwidthTier = &in.TotalEgressBandwidthTier
	return out
}

func InstanceParams_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceParams) *krm.InstanceParams {
	if in == nil {
		return nil
	}
	out := &krm.InstanceParams{}
	return out
}

func InstanceParams_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceParams) *pb.InstanceParams {
	if in == nil {
		return nil
	}
	out := &pb.InstanceParams{}
	return out
}

func InstanceReservationAffinity_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ReservationAffinity) *krm.InstanceReservationAffinity {
	if in == nil {
		return nil
	}
	out := &krm.InstanceReservationAffinity{}
	if in.ConsumeReservationType != nil {
		out.Type = *in.ConsumeReservationType
	}
	if in.Key != nil || len(in.Values) > 0 {
		out.SpecificReservation = &krm.InstanceSpecificReservation{}
		if in.Key != nil {
			out.SpecificReservation.Key = *in.Key
		}
		out.SpecificReservation.Values = in.Values
	}
	return out
}

func InstanceReservationAffinity_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceReservationAffinity) *pb.ReservationAffinity {
	if in == nil {
		return nil
	}
	out := &pb.ReservationAffinity{}
	out.ConsumeReservationType = &in.Type
	if in.SpecificReservation != nil {
		out.Key = &in.SpecificReservation.Key
		out.Values = in.SpecificReservation.Values
	}
	return out
}

func InstanceScheduling_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Scheduling) *krm.InstanceScheduling {
	if in == nil {
		return nil
	}
	out := &krm.InstanceScheduling{}
	out.AutomaticRestart = in.AutomaticRestart
	out.OnHostMaintenance = in.OnHostMaintenance
	out.Preemptible = in.Preemptible
	out.ProvisioningModel = in.ProvisioningModel
	out.InstanceTerminationAction = in.InstanceTerminationAction
	if in.MinNodeCpus != nil {
		v := int64(*in.MinNodeCpus)
		out.MinNodeCpus = &v
	}
	return out
}

func InstanceScheduling_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceScheduling) *pb.Scheduling {
	if in == nil {
		return nil
	}
	out := &pb.Scheduling{}
	out.AutomaticRestart = in.AutomaticRestart
	out.OnHostMaintenance = in.OnHostMaintenance
	out.Preemptible = in.Preemptible
	out.ProvisioningModel = in.ProvisioningModel
	out.InstanceTerminationAction = in.InstanceTerminationAction
	if in.MinNodeCpus != nil {
		v := int32(*in.MinNodeCpus)
		out.MinNodeCpus = &v
	}
	return out
}

func InstanceShieldedInstanceConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ShieldedInstanceConfig) *krm.InstanceShieldedInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.InstanceShieldedInstanceConfig{}
	out.EnableIntegrityMonitoring = in.EnableIntegrityMonitoring
	out.EnableSecureBoot = in.EnableSecureBoot
	out.EnableVtpm = in.EnableVtpm
	return out
}

func InstanceShieldedInstanceConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceShieldedInstanceConfig) *pb.ShieldedInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.ShieldedInstanceConfig{}
	out.EnableIntegrityMonitoring = in.EnableIntegrityMonitoring
	out.EnableSecureBoot = in.EnableSecureBoot
	out.EnableVtpm = in.EnableVtpm
	return out
}

func InstanceDiskEncryptionKey_v1beta1_ToProto(mapCtx *direct.MapContext, kmsKeyRef *krm.InstanceResourceRef, diskEncryptionKeyRaw *krm.InstanceDiskEncryptionKeyRaw) *pb.CustomerEncryptionKey {
	if kmsKeyRef == nil {
		return nil
	}
	out := &pb.CustomerEncryptionKey{}
	if kmsKeyRef.External != "" {
		out.KmsKeyName = direct.PtrTo(kmsKeyRef.External)
	}
	return out
}

func InstanceDiskEncryptionKey_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.CustomerEncryptionKey) (*krm.InstanceResourceRef, *string) {
	if in == nil {
		return nil, nil
	}
	var kmsKeyRef *krm.InstanceResourceRef
	if in.KmsKeyName != nil {
		kmsKeyRef = &krm.InstanceResourceRef{
			External: *in.KmsKeyName,
		}
	}
	return kmsKeyRef, in.Sha256
}

func InstanceBootDisk_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AttachedDisk) *krm.InstanceBootDisk {
	if in == nil {
		return nil
	}
	out := &krm.InstanceBootDisk{}
	out.AutoDelete = in.AutoDelete
	out.DeviceName = in.DeviceName
	out.Mode = in.Mode
	if in.Source != nil {
		out.SourceDiskRef = &krm.InstanceResourceRef{
			External: *in.Source,
		}
	}
	if in.DiskEncryptionKey != nil {
		kmsKeyRef, sha256 := InstanceDiskEncryptionKey_v1beta1_FromProto(mapCtx, in.DiskEncryptionKey)
		out.KmsKeyRef = kmsKeyRef
		out.DiskEncryptionKeySha256 = sha256
	}
	if in.InitializeParams != nil {
		out.InitializeParams = &krm.InstanceInitializeParams{}
		out.InitializeParams.Size = in.InitializeParams.DiskSizeGb
		out.InitializeParams.Type = in.InitializeParams.DiskType
		if in.InitializeParams.SourceImage != nil {
			out.InitializeParams.SourceImageRef = &krm.InstanceResourceRef{
				External: *in.InitializeParams.SourceImage,
			}
		}
	}
	return out
}

func InstanceBootDisk_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceBootDisk) *pb.AttachedDisk {
	if in == nil {
		return nil
	}
	out := &pb.AttachedDisk{}
	out.AutoDelete = in.AutoDelete
	out.DeviceName = in.DeviceName
	out.Mode = in.Mode
	if in.SourceDiskRef != nil {
		out.Source = &in.SourceDiskRef.External
	}
	if in.KmsKeyRef != nil {
		out.DiskEncryptionKey = InstanceDiskEncryptionKey_v1beta1_ToProto(mapCtx, in.KmsKeyRef, in.DiskEncryptionKeyRaw)
	}
	if in.InitializeParams != nil {
		out.InitializeParams = &pb.AttachedDiskInitializeParams{}
		out.InitializeParams.DiskSizeGb = in.InitializeParams.Size
		out.InitializeParams.DiskType = in.InitializeParams.Type
		if in.InitializeParams.SourceImageRef != nil {
			out.InitializeParams.SourceImage = &in.InitializeParams.SourceImageRef.External
		}
	}
	return out
}

func InstanceAttachedDisk_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AttachedDisk) *krm.InstanceAttachedDisk {
	if in == nil {
		return nil
	}
	out := &krm.InstanceAttachedDisk{}
	out.DeviceName = in.DeviceName
	out.Mode = in.Mode
	if in.Source != nil {
		out.SourceDiskRef = krm.InstanceResourceRef{
			External: *in.Source,
		}
	}
	if in.DiskEncryptionKey != nil {
		kmsKeyRef, sha256 := InstanceDiskEncryptionKey_v1beta1_FromProto(mapCtx, in.DiskEncryptionKey)
		out.KmsKeyRef = kmsKeyRef
		out.DiskEncryptionKeySha256 = sha256
	}
	return out
}

func InstanceAttachedDisk_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceAttachedDisk) *pb.AttachedDisk {
	if in == nil {
		return nil
	}
	out := &pb.AttachedDisk{}
	out.DeviceName = in.DeviceName
	out.Mode = in.Mode
	out.Source = &in.SourceDiskRef.External
	if in.KmsKeyRef != nil {
		out.DiskEncryptionKey = InstanceDiskEncryptionKey_v1beta1_ToProto(mapCtx, in.KmsKeyRef, in.DiskEncryptionKeyRaw)
	}
	return out
}

func InstanceAccessConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AccessConfig) *krm.InstanceAccessConfig {
	if in == nil {
		return nil
	}
	out := &krm.InstanceAccessConfig{}
	out.NetworkTier = in.NetworkTier
	out.PublicPtrDomainName = in.PublicPtrDomainName
	if in.NatIP != nil {
		out.NatIpRef = &krm.InstanceResourceRef{
			External: *in.NatIP,
		}
	}
	return out
}

func InstanceAccessConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceAccessConfig) *pb.AccessConfig {
	if in == nil {
		return nil
	}
	out := &pb.AccessConfig{}
	out.NetworkTier = in.NetworkTier
	out.PublicPtrDomainName = in.PublicPtrDomainName
	out.Name = direct.PtrTo("External NAT")
	out.Type = direct.PtrTo("ONE_TO_ONE_NAT")
	if in.NatIpRef != nil {
		out.NatIP = &in.NatIpRef.External
	}
	return out
}

func InstanceAliasIpRange_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AliasIpRange) *krm.InstanceAliasIpRange {
	if in == nil {
		return nil
	}
	out := &krm.InstanceAliasIpRange{}
	if in.IpCidrRange != nil {
		out.IpCidrRange = *in.IpCidrRange
	}
	out.SubnetworkRangeName = in.SubnetworkRangeName
	return out
}

func InstanceAliasIpRange_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceAliasIpRange) *pb.AliasIpRange {
	if in == nil {
		return nil
	}
	out := &pb.AliasIpRange{}
	out.IpCidrRange = &in.IpCidrRange
	out.SubnetworkRangeName = in.SubnetworkRangeName
	return out
}

func InstanceIpv6AccessConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AccessConfig) *krm.InstanceIpv6AccessConfig {
	if in == nil {
		return nil
	}
	out := &krm.InstanceIpv6AccessConfig{}
	out.NetworkTier = direct.ValueOf(in.NetworkTier)
	out.PublicPtrDomainName = in.PublicPtrDomainName
	return out
}

func InstanceIpv6AccessConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceIpv6AccessConfig) *pb.AccessConfig {
	if in == nil {
		return nil
	}
	out := &pb.AccessConfig{}
	out.NetworkTier = &in.NetworkTier
	out.PublicPtrDomainName = in.PublicPtrDomainName
	out.Name = direct.PtrTo("External IPv6")
	out.Type = direct.PtrTo("DIRECT_IPV6")
	return out
}

func InstanceNetworkInterface_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.NetworkInterface) *krm.InstanceNetworkInterface {
	if in == nil {
		return nil
	}
	out := &krm.InstanceNetworkInterface{}
	out.Name = in.Name
	out.Ipv6AccessType = in.Ipv6AccessType
	out.Ipv6Address = in.Ipv6Address
	out.NicType = in.NicType
	out.StackType = in.StackType
	if in.InternalIpv6PrefixLength != nil {
		v := int64(*in.InternalIpv6PrefixLength)
		out.InternalIpv6PrefixLength = &v
	}
	if in.QueueCount != nil {
		v := int64(*in.QueueCount)
		out.QueueCount = &v
	}
	if in.Network != nil {
		out.NetworkRef = &krm.InstanceResourceRef{
			External: *in.Network,
		}
	}
	if in.Subnetwork != nil {
		out.SubnetworkRef = &krm.InstanceResourceRef{
			External: *in.Subnetwork,
		}
	}
	if in.NetworkIP != nil {
		out.NetworkIpRef = &k8sv1alpha1.ResourceRef{
			External: *in.NetworkIP,
		}
	}
	for _, ac := range in.AccessConfigs {
		krmAc := InstanceAccessConfig_v1beta1_FromProto(mapCtx, ac)
		if krmAc != nil {
			out.AccessConfig = append(out.AccessConfig, *krmAc)
		}
	}
	for _, ac := range in.Ipv6AccessConfigs {
		krmAc := InstanceIpv6AccessConfig_v1beta1_FromProto(mapCtx, ac)
		if krmAc != nil {
			out.Ipv6AccessConfig = append(out.Ipv6AccessConfig, *krmAc)
		}
	}
	for _, r := range in.AliasIpRanges {
		krmR := InstanceAliasIpRange_v1beta1_FromProto(mapCtx, r)
		if krmR != nil {
			out.AliasIpRange = append(out.AliasIpRange, *krmR)
		}
	}
	return out
}

func InstanceNetworkInterface_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceNetworkInterface) *pb.NetworkInterface {
	if in == nil {
		return nil
	}
	out := &pb.NetworkInterface{}
	out.Name = in.Name
	out.Ipv6AccessType = in.Ipv6AccessType
	out.Ipv6Address = in.Ipv6Address
	out.NicType = in.NicType
	out.StackType = in.StackType
	if in.InternalIpv6PrefixLength != nil {
		v := int32(*in.InternalIpv6PrefixLength)
		out.InternalIpv6PrefixLength = &v
	}
	if in.QueueCount != nil {
		v := int32(*in.QueueCount)
		out.QueueCount = &v
	}
	if in.NetworkRef != nil {
		out.Network = &in.NetworkRef.External
	}
	if in.SubnetworkRef != nil {
		out.Subnetwork = &in.SubnetworkRef.External
	}
	if in.NetworkIpRef != nil {
		out.NetworkIP = &in.NetworkIpRef.External
	}
	for i := range in.AccessConfig {
		ac := InstanceAccessConfig_v1beta1_ToProto(mapCtx, &in.AccessConfig[i])
		if ac != nil {
			out.AccessConfigs = append(out.AccessConfigs, ac)
		}
	}
	for i := range in.Ipv6AccessConfig {
		ac := InstanceIpv6AccessConfig_v1beta1_ToProto(mapCtx, &in.Ipv6AccessConfig[i])
		if ac != nil {
			out.Ipv6AccessConfigs = append(out.Ipv6AccessConfigs, ac)
		}
	}
	for i := range in.AliasIpRange {
		r := InstanceAliasIpRange_v1beta1_ToProto(mapCtx, &in.AliasIpRange[i])
		if r != nil {
			out.AliasIpRanges = append(out.AliasIpRanges, r)
		}
	}
	return out
}

func InstanceServiceAccount_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ServiceAccount) *krm.InstanceServiceAccount {
	if in == nil {
		return nil
	}
	out := &krm.InstanceServiceAccount{}
	out.Scopes = in.Scopes
	if in.Email != nil {
		out.ServiceAccountRef = &krm.InstanceResourceRef{
			External: *in.Email,
		}
	}
	return out
}

func InstanceServiceAccount_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceServiceAccount) *pb.ServiceAccount {
	if in == nil {
		return nil
	}
	out := &pb.ServiceAccount{}
	out.Scopes = in.Scopes
	if in.ServiceAccountRef != nil {
		out.Email = &in.ServiceAccountRef.External
	}
	return out
}
