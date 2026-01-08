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
// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.compute.v1

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krmcomputev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AcceleratorConfig_FromProto(mapCtx *direct.MapContext, in *pb.AcceleratorConfig) *krm.AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &krm.AcceleratorConfig{}
	out.AcceleratorCount = in.AcceleratorCount
	out.AcceleratorType = in.AcceleratorType
	return out
}
func AcceleratorConfig_ToProto(mapCtx *direct.MapContext, in *krm.AcceleratorConfig) *pb.AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &pb.AcceleratorConfig{}
	out.AcceleratorCount = in.AcceleratorCount
	out.AcceleratorType = in.AcceleratorType
	return out
}
func AccessConfig_FromProto(mapCtx *direct.MapContext, in *pb.AccessConfig) *krm.AccessConfig {
	if in == nil {
		return nil
	}
	out := &krm.AccessConfig{}
	out.ExternalIPV6 = in.ExternalIpv6
	out.ExternalIPV6PrefixLength = in.ExternalIpv6PrefixLength
	out.Kind = in.Kind
	out.Name = in.Name
	out.NATIP = in.NatIP
	out.NetworkTier = in.NetworkTier
	out.PublicPtrDomainName = in.PublicPtrDomainName
	out.SecurityPolicy = in.SecurityPolicy
	out.SetPublicPtr = in.SetPublicPtr
	out.Type = in.Type
	return out
}
func AccessConfig_ToProto(mapCtx *direct.MapContext, in *krm.AccessConfig) *pb.AccessConfig {
	if in == nil {
		return nil
	}
	out := &pb.AccessConfig{}
	out.ExternalIpv6 = in.ExternalIPV6
	out.ExternalIpv6PrefixLength = in.ExternalIPV6PrefixLength
	out.Kind = in.Kind
	out.Name = in.Name
	out.NatIP = in.NATIP
	out.NetworkTier = in.NetworkTier
	out.PublicPtrDomainName = in.PublicPtrDomainName
	out.SecurityPolicy = in.SecurityPolicy
	out.SetPublicPtr = in.SetPublicPtr
	out.Type = in.Type
	return out
}
func AdvancedMachineFeatures_FromProto(mapCtx *direct.MapContext, in *pb.AdvancedMachineFeatures) *krm.AdvancedMachineFeatures {
	if in == nil {
		return nil
	}
	out := &krm.AdvancedMachineFeatures{}
	out.EnableNestedVirtualization = in.EnableNestedVirtualization
	out.EnableUefiNetworking = in.EnableUefiNetworking
	out.PerformanceMonitoringUnit = in.PerformanceMonitoringUnit
	out.ThreadsPerCore = in.ThreadsPerCore
	out.TurboMode = in.TurboMode
	out.VisibleCoreCount = in.VisibleCoreCount
	return out
}
func AdvancedMachineFeatures_ToProto(mapCtx *direct.MapContext, in *krm.AdvancedMachineFeatures) *pb.AdvancedMachineFeatures {
	if in == nil {
		return nil
	}
	out := &pb.AdvancedMachineFeatures{}
	out.EnableNestedVirtualization = in.EnableNestedVirtualization
	out.EnableUefiNetworking = in.EnableUefiNetworking
	out.PerformanceMonitoringUnit = in.PerformanceMonitoringUnit
	out.ThreadsPerCore = in.ThreadsPerCore
	out.TurboMode = in.TurboMode
	out.VisibleCoreCount = in.VisibleCoreCount
	return out
}
func AliasIPRange_FromProto(mapCtx *direct.MapContext, in *pb.AliasIpRange) *krm.AliasIPRange {
	if in == nil {
		return nil
	}
	out := &krm.AliasIPRange{}
	out.IPCIDRRange = in.IpCidrRange
	out.SubnetworkRangeName = in.SubnetworkRangeName
	return out
}
func AliasIPRange_ToProto(mapCtx *direct.MapContext, in *krm.AliasIPRange) *pb.AliasIpRange {
	if in == nil {
		return nil
	}
	out := &pb.AliasIpRange{}
	out.IpCidrRange = in.IPCIDRRange
	out.SubnetworkRangeName = in.SubnetworkRangeName
	return out
}
func AttachedDisk_FromProto(mapCtx *direct.MapContext, in *pb.AttachedDisk) *krm.AttachedDisk {
	if in == nil {
		return nil
	}
	out := &krm.AttachedDisk{}
	out.Architecture = in.Architecture
	out.AutoDelete = in.AutoDelete
	out.Boot = in.Boot
	out.DeviceName = in.DeviceName
	out.DiskEncryptionKey = CustomerEncryptionKey_FromProto(mapCtx, in.GetDiskEncryptionKey())
	out.DiskSizeGB = in.DiskSizeGb
	out.ForceAttach = in.ForceAttach
	out.GuestOSFeatures = direct.Slice_FromProto(mapCtx, in.GuestOsFeatures, GuestOSFeature_FromProto)
	out.Index = in.Index
	out.InitializeParams = AttachedDiskInitializeParams_FromProto(mapCtx, in.GetInitializeParams())
	out.Interface = in.Interface
	out.Kind = in.Kind
	out.Licenses = in.Licenses
	out.Mode = in.Mode
	out.SavedState = in.SavedState
	out.ShieldedInstanceInitialState = InitialStateConfig_FromProto(mapCtx, in.GetShieldedInstanceInitialState())
	out.Source = in.Source
	out.Type = in.Type
	return out
}
func AttachedDisk_ToProto(mapCtx *direct.MapContext, in *krm.AttachedDisk) *pb.AttachedDisk {
	if in == nil {
		return nil
	}
	out := &pb.AttachedDisk{}
	out.Architecture = in.Architecture
	out.AutoDelete = in.AutoDelete
	out.Boot = in.Boot
	out.DeviceName = in.DeviceName
	out.DiskEncryptionKey = CustomerEncryptionKey_ToProto(mapCtx, in.DiskEncryptionKey)
	out.DiskSizeGb = in.DiskSizeGB
	out.ForceAttach = in.ForceAttach
	out.GuestOsFeatures = direct.Slice_ToProto(mapCtx, in.GuestOSFeatures, GuestOSFeature_ToProto)
	out.Index = in.Index
	out.InitializeParams = AttachedDiskInitializeParams_ToProto(mapCtx, in.InitializeParams)
	out.Interface = in.Interface
	out.Kind = in.Kind
	out.Licenses = in.Licenses
	out.Mode = in.Mode
	out.SavedState = in.SavedState
	out.ShieldedInstanceInitialState = InitialStateConfig_ToProto(mapCtx, in.ShieldedInstanceInitialState)
	out.Source = in.Source
	out.Type = in.Type
	return out
}
func AttachedDiskInitializeParams_FromProto(mapCtx *direct.MapContext, in *pb.AttachedDiskInitializeParams) *krm.AttachedDiskInitializeParams {
	if in == nil {
		return nil
	}
	out := &krm.AttachedDiskInitializeParams{}
	out.Architecture = in.Architecture
	out.Description = in.Description
	out.DiskName = in.DiskName
	out.DiskSizeGB = in.DiskSizeGb
	out.DiskType = in.DiskType
	out.EnableConfidentialCompute = in.EnableConfidentialCompute
	out.Labels = in.Labels
	out.Licenses = in.Licenses
	out.OnUpdateAction = in.OnUpdateAction
	out.ProvisionedIops = in.ProvisionedIops
	out.ProvisionedThroughput = in.ProvisionedThroughput
	out.ReplicaZones = in.ReplicaZones
	out.ResourceManagerTags = in.ResourceManagerTags
	out.ResourcePolicies = in.ResourcePolicies
	out.SourceImage = in.SourceImage
	out.SourceImageEncryptionKey = CustomerEncryptionKey_FromProto(mapCtx, in.GetSourceImageEncryptionKey())
	out.SourceSnapshot = in.SourceSnapshot
	out.SourceSnapshotEncryptionKey = CustomerEncryptionKey_FromProto(mapCtx, in.GetSourceSnapshotEncryptionKey())
	out.StoragePool = in.StoragePool
	return out
}
func AttachedDiskInitializeParams_ToProto(mapCtx *direct.MapContext, in *krm.AttachedDiskInitializeParams) *pb.AttachedDiskInitializeParams {
	if in == nil {
		return nil
	}
	out := &pb.AttachedDiskInitializeParams{}
	out.Architecture = in.Architecture
	out.Description = in.Description
	out.DiskName = in.DiskName
	out.DiskSizeGb = in.DiskSizeGB
	out.DiskType = in.DiskType
	out.EnableConfidentialCompute = in.EnableConfidentialCompute
	out.Labels = in.Labels
	out.Licenses = in.Licenses
	out.OnUpdateAction = in.OnUpdateAction
	out.ProvisionedIops = in.ProvisionedIops
	out.ProvisionedThroughput = in.ProvisionedThroughput
	out.ReplicaZones = in.ReplicaZones
	out.ResourceManagerTags = in.ResourceManagerTags
	out.ResourcePolicies = in.ResourcePolicies
	out.SourceImage = in.SourceImage
	out.SourceImageEncryptionKey = CustomerEncryptionKey_ToProto(mapCtx, in.SourceImageEncryptionKey)
	out.SourceSnapshot = in.SourceSnapshot
	out.SourceSnapshotEncryptionKey = CustomerEncryptionKey_ToProto(mapCtx, in.SourceSnapshotEncryptionKey)
	out.StoragePool = in.StoragePool
	return out
}
func ComputeInstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.ComputeInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ComputeInstanceObservedState{}
	// MISSING: CPUPlatform
	// MISSING: CreationTimestamp
	// MISSING: Disks
	// MISSING: DisplayDevice
	// MISSING: Fingerprint
	// MISSING: GuestAccelerators
	// MISSING: ID
	// MISSING: InstanceEncryptionKey
	// MISSING: KeyRevocationActionType
	// MISSING: Kind
	// MISSING: LabelFingerprint
	// MISSING: Labels
	// MISSING: LastStartTimestamp
	// MISSING: LastStopTimestamp
	// MISSING: LastSuspendedTimestamp
	// MISSING: Name
	// MISSING: NetworkInterfaces
	// MISSING: PrivateIPV6GoogleAccess
	// MISSING: ResourceStatus
	// MISSING: SatisfiesPzi
	// MISSING: SatisfiesPzs
	// MISSING: SelfLink
	// MISSING: ServiceAccounts
	// MISSING: ShieldedInstanceIntegrityPolicy
	// MISSING: SourceMachineImage
	// MISSING: SourceMachineImageEncryptionKey
	// MISSING: StartRestricted
	// MISSING: Status
	// MISSING: StatusMessage
	return out
}
func ComputeInstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: CPUPlatform
	// MISSING: CreationTimestamp
	// MISSING: Disks
	// MISSING: DisplayDevice
	// MISSING: Fingerprint
	// MISSING: GuestAccelerators
	// MISSING: ID
	// MISSING: InstanceEncryptionKey
	// MISSING: KeyRevocationActionType
	// MISSING: Kind
	// MISSING: LabelFingerprint
	// MISSING: Labels
	// MISSING: LastStartTimestamp
	// MISSING: LastStopTimestamp
	// MISSING: LastSuspendedTimestamp
	// MISSING: Name
	// MISSING: NetworkInterfaces
	// MISSING: PrivateIPV6GoogleAccess
	// MISSING: ResourceStatus
	// MISSING: SatisfiesPzi
	// MISSING: SatisfiesPzs
	// MISSING: SelfLink
	// MISSING: ServiceAccounts
	// MISSING: ShieldedInstanceIntegrityPolicy
	// MISSING: SourceMachineImage
	// MISSING: SourceMachineImageEncryptionKey
	// MISSING: StartRestricted
	// MISSING: Status
	// MISSING: StatusMessage
	return out
}
func ComputeInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.ComputeInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeInstanceSpec{}
	out.AdvancedMachineFeatures = ComputeInstanceAdvancedMachineFeatures_FromProto(mapCtx, in.GetAdvancedMachineFeatures())
	out.CanIPForward = in.CanIpForward
	out.ConfidentialInstanceConfig = ComputeInstanceConfidentialInstanceConfig_FromProto(mapCtx, in.GetConfidentialInstanceConfig())
	// MISSING: CPUPlatform
	// MISSING: CreationTimestamp
	out.DeletionProtection = in.DeletionProtection
	out.Description = in.Description
	// MISSING: Disks
	// MISSING: DisplayDevice
	// MISSING: Fingerprint
	// MISSING: GuestAccelerators
	out.Hostname = in.Hostname
	// MISSING: ID
	// MISSING: InstanceEncryptionKey
	// MISSING: KeyRevocationActionType
	// MISSING: Kind
	// MISSING: LabelFingerprint
	// MISSING: Labels
	// MISSING: LastStartTimestamp
	// MISSING: LastStopTimestamp
	// MISSING: LastSuspendedTimestamp
	out.MachineType = in.MachineType
	if v := in.GetMetadata(); v != nil {
		out.Metadata = []krm.ComputeInstanceMetadata{ComputeInstanceMetadata_FromProto(mapCtx, v)}
	}
	out.MinCPUPlatform = in.MinCpuPlatform
	// MISSING: Name
	// MISSING: NetworkInterfaces
	out.NetworkPerformanceConfig = NetworkPerformanceConfig_FromProto(mapCtx, in.GetNetworkPerformanceConfig())
	out.Params = InstanceParams_FromProto(mapCtx, in.GetParams())
	// MISSING: PrivateIPV6GoogleAccess
	out.ReservationAffinity = ReservationAffinity_FromProto(mapCtx, in.GetReservationAffinity())
	out.ResourcePolicies = ComputeInstanceSpec_ResourcePolicies_FromProto(mapCtx, in.ResourcePolicies)
	// MISSING: ResourceStatus
	// MISSING: SatisfiesPzi
	// MISSING: SatisfiesPzs
	out.Scheduling = Scheduling_FromProto(mapCtx, in.GetScheduling())
	// MISSING: SelfLink
	// MISSING: ServiceAccounts
	out.ShieldedInstanceConfig = ComputeInstanceShieldedInstanceConfig_FromProto(mapCtx, in.GetShieldedInstanceConfig())
	// MISSING: ShieldedInstanceIntegrityPolicy
	// MISSING: SourceMachineImage
	// MISSING: SourceMachineImageEncryptionKey
	// MISSING: StartRestricted
	// MISSING: Status
	// MISSING: StatusMessage
	if v := in.GetTags(); v != nil {
		out.Tags = []krm.string{string_FromProto(mapCtx, v)}
	}
	out.Zone = in.Zone
	return out
}
func ComputeInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.AdvancedMachineFeatures = ComputeInstanceAdvancedMachineFeatures_ToProto(mapCtx, in.AdvancedMachineFeatures)
	out.CanIpForward = in.CanIPForward
	out.ConfidentialInstanceConfig = ComputeInstanceConfidentialInstanceConfig_ToProto(mapCtx, in.ConfidentialInstanceConfig)
	// MISSING: CPUPlatform
	// MISSING: CreationTimestamp
	out.DeletionProtection = in.DeletionProtection
	out.Description = in.Description
	// MISSING: Disks
	// MISSING: DisplayDevice
	// MISSING: Fingerprint
	// MISSING: GuestAccelerators
	out.Hostname = in.Hostname
	// MISSING: ID
	// MISSING: InstanceEncryptionKey
	// MISSING: KeyRevocationActionType
	// MISSING: Kind
	// MISSING: LabelFingerprint
	// MISSING: Labels
	// MISSING: LastStartTimestamp
	// MISSING: LastStopTimestamp
	// MISSING: LastSuspendedTimestamp
	out.MachineType = in.MachineType
	if len(in.Metadata) > 0 && in.Metadata[0] != nil {
		out.Metadata = ComputeInstanceMetadata_ToProto(mapCtx, in.Metadata[0])
	}
	out.MinCpuPlatform = in.MinCPUPlatform
	// MISSING: Name
	// MISSING: NetworkInterfaces
	out.NetworkPerformanceConfig = NetworkPerformanceConfig_ToProto(mapCtx, in.NetworkPerformanceConfig)
	out.Params = InstanceParams_ToProto(mapCtx, in.Params)
	// MISSING: PrivateIPV6GoogleAccess
	out.ReservationAffinity = ReservationAffinity_ToProto(mapCtx, in.ReservationAffinity)
	out.ResourcePolicies = ComputeInstanceSpec_ResourcePolicies_ToProto(mapCtx, in.ResourcePolicies)
	// MISSING: ResourceStatus
	// MISSING: SatisfiesPzi
	// MISSING: SatisfiesPzs
	out.Scheduling = Scheduling_ToProto(mapCtx, in.Scheduling)
	// MISSING: SelfLink
	// MISSING: ServiceAccounts
	out.ShieldedInstanceConfig = ComputeInstanceShieldedInstanceConfig_ToProto(mapCtx, in.ShieldedInstanceConfig)
	// MISSING: ShieldedInstanceIntegrityPolicy
	// MISSING: SourceMachineImage
	// MISSING: SourceMachineImageEncryptionKey
	// MISSING: StartRestricted
	// MISSING: Status
	// MISSING: StatusMessage
	if len(in.Tags) > 0 && in.Tags[0] != nil {
		out.Tags = string_ToProto(mapCtx, in.Tags[0])
	}
	out.Zone = in.Zone
	return out
}
func ComputeSubnetworkSpec_FromProto(mapCtx *direct.MapContext, in *pb.Subnetwork) *krm.ComputeSubnetworkSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeSubnetworkSpec{}
	// MISSING: CreationTimestamp
	out.Description = in.Description
	// MISSING: EnableFlowLogs
	// MISSING: ExternalIPV6Prefix
	// MISSING: Fingerprint
	// MISSING: GatewayAddress
	// MISSING: ID
	// MISSING: InternalIPV6Prefix
	out.IPCIDRRange = in.IpCidrRange
	// MISSING: IPCollection
	out.IPV6AccessType = in.Ipv6AccessType
	// MISSING: IPV6CIDRRange
	// MISSING: IPV6GCEEndpoint
	// MISSING: Kind
	out.LogConfig = SubnetworkLogConfig_FromProto(mapCtx, in.GetLogConfig())
	// MISSING: Name
	if in.GetNetwork() != "" {
		out.NetworkRef = &krm.ComputeNetworkRef{External: in.GetNetwork()}
	}
	// MISSING: Params
	out.PrivateIPGoogleAccess = in.PrivateIpGoogleAccess
	out.PrivateIPV6GoogleAccess = in.PrivateIpv6GoogleAccess
	out.Purpose = in.Purpose
	out.Region = in.Region
	// MISSING: ReservedInternalRange
	out.Role = in.Role
	out.SecondaryIPRanges = direct.Slice_FromProto(mapCtx, in.SecondaryIpRanges, SubnetworkSecondaryRange_FromProto)
	// MISSING: SelfLink
	out.StackType = in.StackType
	// MISSING: State
	// MISSING: SystemReservedExternalIPV6Ranges
	// MISSING: SystemReservedInternalIPV6Ranges
	return out
}
func ComputeSubnetworkSpec_ToProto(mapCtx *direct.MapContext, in *krm.ComputeSubnetworkSpec) *pb.Subnetwork {
	if in == nil {
		return nil
	}
	out := &pb.Subnetwork{}
	// MISSING: CreationTimestamp
	out.Description = in.Description
	// MISSING: EnableFlowLogs
	// MISSING: ExternalIPV6Prefix
	// MISSING: Fingerprint
	// MISSING: GatewayAddress
	// MISSING: ID
	// MISSING: InternalIPV6Prefix
	out.IpCidrRange = in.IPCIDRRange
	// MISSING: IPCollection
	out.Ipv6AccessType = in.IPV6AccessType
	// MISSING: IPV6CIDRRange
	// MISSING: IPV6GCEEndpoint
	// MISSING: Kind
	out.LogConfig = SubnetworkLogConfig_ToProto(mapCtx, in.LogConfig)
	// MISSING: Name
	if in.NetworkRef != nil {
		out.Network = &in.NetworkRef.External
	}
	// MISSING: Params
	out.PrivateIpGoogleAccess = in.PrivateIPGoogleAccess
	out.PrivateIpv6GoogleAccess = in.PrivateIPV6GoogleAccess
	out.Purpose = in.Purpose
	out.Region = in.Region
	// MISSING: ReservedInternalRange
	out.Role = in.Role
	out.SecondaryIpRanges = direct.Slice_ToProto(mapCtx, in.SecondaryIPRanges, SubnetworkSecondaryRange_ToProto)
	// MISSING: SelfLink
	out.StackType = in.StackType
	// MISSING: State
	// MISSING: SystemReservedExternalIPV6Ranges
	// MISSING: SystemReservedInternalIPV6Ranges
	return out
}
func ComputeSubnetworkStatus_FromProto(mapCtx *direct.MapContext, in *pb.Subnetwork) *krm.ComputeSubnetworkStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeSubnetworkStatus{}
	out.CreationTimestamp = in.CreationTimestamp
	// MISSING: Description
	// MISSING: EnableFlowLogs
	out.ExternalIPV6Prefix = in.ExternalIpv6Prefix
	out.Fingerprint = in.Fingerprint
	out.GatewayAddress = in.GatewayAddress
	// MISSING: ID
	out.InternalIPV6Prefix = in.InternalIpv6Prefix
	// MISSING: IPCIDRRange
	// MISSING: IPCollection
	// MISSING: IPV6AccessType
	out.IPV6CIDRRange = in.Ipv6CidrRange
	// MISSING: IPV6GCEEndpoint
	// MISSING: Kind
	// MISSING: LogConfig
	// MISSING: Name
	// MISSING: Network
	// MISSING: Params
	// MISSING: PrivateIPGoogleAccess
	// MISSING: PrivateIPV6GoogleAccess
	// MISSING: Purpose
	// MISSING: Region
	// MISSING: ReservedInternalRange
	// MISSING: Role
	// MISSING: SecondaryIPRanges
	out.SelfLink = in.SelfLink
	// MISSING: StackType
	// MISSING: State
	// MISSING: SystemReservedExternalIPV6Ranges
	// MISSING: SystemReservedInternalIPV6Ranges
	return out
}
func ComputeSubnetworkStatus_ToProto(mapCtx *direct.MapContext, in *krm.ComputeSubnetworkStatus) *pb.Subnetwork {
	if in == nil {
		return nil
	}
	out := &pb.Subnetwork{}
	out.CreationTimestamp = in.CreationTimestamp
	// MISSING: Description
	// MISSING: EnableFlowLogs
	out.ExternalIpv6Prefix = in.ExternalIPV6Prefix
	out.Fingerprint = in.Fingerprint
	out.GatewayAddress = in.GatewayAddress
	// MISSING: ID
	out.InternalIpv6Prefix = in.InternalIPV6Prefix
	// MISSING: IPCIDRRange
	// MISSING: IPCollection
	// MISSING: IPV6AccessType
	out.Ipv6CidrRange = in.IPV6CIDRRange
	// MISSING: IPV6GCEEndpoint
	// MISSING: Kind
	// MISSING: LogConfig
	// MISSING: Name
	// MISSING: Network
	// MISSING: Params
	// MISSING: PrivateIPGoogleAccess
	// MISSING: PrivateIPV6GoogleAccess
	// MISSING: Purpose
	// MISSING: Region
	// MISSING: ReservedInternalRange
	// MISSING: Role
	// MISSING: SecondaryIPRanges
	out.SelfLink = in.SelfLink
	// MISSING: StackType
	// MISSING: State
	// MISSING: SystemReservedExternalIPV6Ranges
	// MISSING: SystemReservedInternalIPV6Ranges
	return out
}
func ConfidentialInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.ConfidentialInstanceConfig) *krm.ConfidentialInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.ConfidentialInstanceConfig{}
	out.ConfidentialInstanceType = in.ConfidentialInstanceType
	out.EnableConfidentialCompute = in.EnableConfidentialCompute
	return out
}
func ConfidentialInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.ConfidentialInstanceConfig) *pb.ConfidentialInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.ConfidentialInstanceConfig{}
	out.ConfidentialInstanceType = in.ConfidentialInstanceType
	out.EnableConfidentialCompute = in.EnableConfidentialCompute
	return out
}
func CustomerEncryptionKey_FromProto(mapCtx *direct.MapContext, in *pb.CustomerEncryptionKey) *krm.CustomerEncryptionKey {
	if in == nil {
		return nil
	}
	out := &krm.CustomerEncryptionKey{}
	out.KMSKeyName = in.KmsKeyName
	out.KMSKeyServiceAccount = in.KmsKeyServiceAccount
	out.RawKey = in.RawKey
	out.RsaEncryptedKey = in.RsaEncryptedKey
	out.Sha256 = in.Sha256
	return out
}
func CustomerEncryptionKey_ToProto(mapCtx *direct.MapContext, in *krm.CustomerEncryptionKey) *pb.CustomerEncryptionKey {
	if in == nil {
		return nil
	}
	out := &pb.CustomerEncryptionKey{}
	out.KmsKeyName = in.KMSKeyName
	out.KmsKeyServiceAccount = in.KMSKeyServiceAccount
	out.RawKey = in.RawKey
	out.RsaEncryptedKey = in.RsaEncryptedKey
	out.Sha256 = in.Sha256
	return out
}
func DisplayDevice_FromProto(mapCtx *direct.MapContext, in *pb.DisplayDevice) *krm.DisplayDevice {
	if in == nil {
		return nil
	}
	out := &krm.DisplayDevice{}
	out.EnableDisplay = in.EnableDisplay
	return out
}
func DisplayDevice_ToProto(mapCtx *direct.MapContext, in *krm.DisplayDevice) *pb.DisplayDevice {
	if in == nil {
		return nil
	}
	out := &pb.DisplayDevice{}
	out.EnableDisplay = in.EnableDisplay
	return out
}
func Duration_FromProto(mapCtx *direct.MapContext, in *pb.Duration) *krm.Duration {
	if in == nil {
		return nil
	}
	out := &krm.Duration{}
	out.Nanos = in.Nanos
	out.Seconds = in.Seconds
	return out
}
func Duration_ToProto(mapCtx *direct.MapContext, in *krm.Duration) *pb.Duration {
	if in == nil {
		return nil
	}
	out := &pb.Duration{}
	out.Nanos = in.Nanos
	out.Seconds = in.Seconds
	return out
}
func FileContentBuffer_FromProto(mapCtx *direct.MapContext, in *pb.FileContentBuffer) *krm.FileContentBuffer {
	if in == nil {
		return nil
	}
	out := &krm.FileContentBuffer{}
	out.Content = in.Content
	out.FileType = in.FileType
	return out
}
func FileContentBuffer_ToProto(mapCtx *direct.MapContext, in *krm.FileContentBuffer) *pb.FileContentBuffer {
	if in == nil {
		return nil
	}
	out := &pb.FileContentBuffer{}
	out.Content = in.Content
	out.FileType = in.FileType
	return out
}
func FirewallPolicyRuleMatcher_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyRuleMatcher) *krm.FirewallPolicyRuleMatcher {
	if in == nil {
		return nil
	}
	out := &krm.FirewallPolicyRuleMatcher{}
	out.DestAddressGroups = in.DestAddressGroups
	out.DestFqdns = in.DestFqdns
	out.DestIPRanges = in.DestIpRanges
	// MISSING: DestNetworkType
	out.DestRegionCodes = in.DestRegionCodes
	out.DestThreatIntelligences = in.DestThreatIntelligences
	out.Layer4Configs = direct.Slice_FromProto(mapCtx, in.Layer4Configs, FirewallPolicyRuleMatcherLayer4Config_FromProto)
	out.SrcAddressGroups = in.SrcAddressGroups
	out.SrcFqdns = in.SrcFqdns
	out.SrcIPRanges = in.SrcIpRanges
	// MISSING: SrcNetworkType
	// MISSING: SrcNetworks
	out.SrcRegionCodes = in.SrcRegionCodes
	// MISSING: SrcSecureTags
	out.SrcThreatIntelligences = in.SrcThreatIntelligences
	return out
}
func FirewallPolicyRuleMatcher_ToProto(mapCtx *direct.MapContext, in *krm.FirewallPolicyRuleMatcher) *pb.FirewallPolicyRuleMatcher {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicyRuleMatcher{}
	out.DestAddressGroups = in.DestAddressGroups
	out.DestFqdns = in.DestFqdns
	out.DestIpRanges = in.DestIPRanges
	// MISSING: DestNetworkType
	out.DestRegionCodes = in.DestRegionCodes
	out.DestThreatIntelligences = in.DestThreatIntelligences
	out.Layer4Configs = direct.Slice_ToProto(mapCtx, in.Layer4Configs, FirewallPolicyRuleMatcherLayer4Config_ToProto)
	out.SrcAddressGroups = in.SrcAddressGroups
	out.SrcFqdns = in.SrcFqdns
	out.SrcIpRanges = in.SrcIPRanges
	// MISSING: SrcNetworkType
	// MISSING: SrcNetworks
	out.SrcRegionCodes = in.SrcRegionCodes
	// MISSING: SrcSecureTags
	out.SrcThreatIntelligences = in.SrcThreatIntelligences
	return out
}
func FirewallPolicyRuleSecureTag_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyRuleSecureTag) *krm.FirewallPolicyRuleSecureTag {
	if in == nil {
		return nil
	}
	out := &krm.FirewallPolicyRuleSecureTag{}
	out.Name = in.Name
	out.State = in.State
	return out
}
func FirewallPolicyRuleSecureTag_ToProto(mapCtx *direct.MapContext, in *krm.FirewallPolicyRuleSecureTag) *pb.FirewallPolicyRuleSecureTag {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicyRuleSecureTag{}
	out.Name = in.Name
	out.State = in.State
	return out
}
func ForwardingruleServiceDirectoryRegistrations_FromProto(mapCtx *direct.MapContext, in *pb.ForwardingRuleServiceDirectoryRegistration) *krm.ForwardingruleServiceDirectoryRegistrations {
	if in == nil {
		return nil
	}
	out := &krm.ForwardingruleServiceDirectoryRegistrations{}
	out.Namespace = in.Namespace
	out.Service = in.Service
	// MISSING: ServiceDirectoryRegion
	return out
}
func ForwardingruleServiceDirectoryRegistrations_ToProto(mapCtx *direct.MapContext, in *krm.ForwardingruleServiceDirectoryRegistrations) *pb.ForwardingRuleServiceDirectoryRegistration {
	if in == nil {
		return nil
	}
	out := &pb.ForwardingRuleServiceDirectoryRegistration{}
	out.Namespace = in.Namespace
	out.Service = in.Service
	// MISSING: ServiceDirectoryRegion
	return out
}
func GuestOSFeature_FromProto(mapCtx *direct.MapContext, in *pb.GuestOsFeature) *krm.GuestOSFeature {
	if in == nil {
		return nil
	}
	out := &krm.GuestOSFeature{}
	out.Type = in.Type
	return out
}
func GuestOSFeature_ToProto(mapCtx *direct.MapContext, in *krm.GuestOSFeature) *pb.GuestOsFeature {
	if in == nil {
		return nil
	}
	out := &pb.GuestOsFeature{}
	out.Type = in.Type
	return out
}
func InitialStateConfig_FromProto(mapCtx *direct.MapContext, in *pb.InitialStateConfig) *krm.InitialStateConfig {
	if in == nil {
		return nil
	}
	out := &krm.InitialStateConfig{}
	out.Dbs = direct.Slice_FromProto(mapCtx, in.Dbs, FileContentBuffer_FromProto)
	out.Dbxs = direct.Slice_FromProto(mapCtx, in.Dbxs, FileContentBuffer_FromProto)
	out.Keks = direct.Slice_FromProto(mapCtx, in.Keks, FileContentBuffer_FromProto)
	out.Pk = FileContentBuffer_FromProto(mapCtx, in.GetPk())
	return out
}
func InitialStateConfig_ToProto(mapCtx *direct.MapContext, in *krm.InitialStateConfig) *pb.InitialStateConfig {
	if in == nil {
		return nil
	}
	out := &pb.InitialStateConfig{}
	out.Dbs = direct.Slice_ToProto(mapCtx, in.Dbs, FileContentBuffer_ToProto)
	out.Dbxs = direct.Slice_ToProto(mapCtx, in.Dbxs, FileContentBuffer_ToProto)
	out.Keks = direct.Slice_ToProto(mapCtx, in.Keks, FileContentBuffer_ToProto)
	out.Pk = FileContentBuffer_ToProto(mapCtx, in.Pk)
	return out
}
func InstanceParams_FromProto(mapCtx *direct.MapContext, in *pb.InstanceParams) *krm.InstanceParams {
	if in == nil {
		return nil
	}
	out := &krm.InstanceParams{}
	out.ResourceManagerTags = in.ResourceManagerTags
	return out
}
func InstanceParams_ToProto(mapCtx *direct.MapContext, in *krm.InstanceParams) *pb.InstanceParams {
	if in == nil {
		return nil
	}
	out := &pb.InstanceParams{}
	out.ResourceManagerTags = in.ResourceManagerTags
	return out
}
func Interconnect_FromProto(mapCtx *direct.MapContext, in *pb.Interconnect) *krmcomputev1alpha1.Interconnect {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.Interconnect{}
	// MISSING: AaiEnabled
	out.AdminEnabled = in.AdminEnabled
	// MISSING: ApplicationAwareInterconnect
	out.AvailableFeatures = in.AvailableFeatures
	out.CircuitInfos = direct.Slice_FromProto(mapCtx, in.CircuitInfos, InterconnectCircuitInfo_FromProto)
	out.CreationTimestamp = in.CreationTimestamp
	out.CustomerName = in.CustomerName
	out.Description = in.Description
	out.ExpectedOutages = direct.Slice_FromProto(mapCtx, in.ExpectedOutages, InterconnectOutageNotification_FromProto)
	out.GoogleIPAddress = in.GoogleIpAddress
	out.GoogleReferenceID = in.GoogleReferenceId
	out.ID = in.Id
	out.InterconnectAttachments = in.InterconnectAttachments
	// MISSING: InterconnectGroups
	out.InterconnectType = in.InterconnectType
	out.Kind = in.Kind
	out.LabelFingerprint = in.LabelFingerprint
	out.Labels = in.Labels
	out.LinkType = in.LinkType
	out.Location = in.Location
	out.Macsec = InterconnectMacsec_FromProto(mapCtx, in.GetMacsec())
	out.MacsecEnabled = in.MacsecEnabled
	out.Name = in.Name
	out.NocContactEmail = in.NocContactEmail
	out.OperationalStatus = in.OperationalStatus
	out.PeerIPAddress = in.PeerIpAddress
	out.ProvisionedLinkCount = in.ProvisionedLinkCount
	out.RemoteLocation = in.RemoteLocation
	out.RequestedFeatures = in.RequestedFeatures
	out.RequestedLinkCount = in.RequestedLinkCount
	out.SatisfiesPzs = in.SatisfiesPzs
	out.SelfLink = in.SelfLink
	out.State = in.State
	return out
}
func Interconnect_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.Interconnect) *pb.Interconnect {
	if in == nil {
		return nil
	}
	out := &pb.Interconnect{}
	// MISSING: AaiEnabled
	out.AdminEnabled = in.AdminEnabled
	// MISSING: ApplicationAwareInterconnect
	out.AvailableFeatures = in.AvailableFeatures
	out.CircuitInfos = direct.Slice_ToProto(mapCtx, in.CircuitInfos, InterconnectCircuitInfo_ToProto)
	out.CreationTimestamp = in.CreationTimestamp
	out.CustomerName = in.CustomerName
	out.Description = in.Description
	out.ExpectedOutages = direct.Slice_ToProto(mapCtx, in.ExpectedOutages, InterconnectOutageNotification_ToProto)
	out.GoogleIpAddress = in.GoogleIPAddress
	out.GoogleReferenceId = in.GoogleReferenceID
	out.Id = in.ID
	out.InterconnectAttachments = in.InterconnectAttachments
	// MISSING: InterconnectGroups
	out.InterconnectType = in.InterconnectType
	out.Kind = in.Kind
	out.LabelFingerprint = in.LabelFingerprint
	out.Labels = in.Labels
	out.LinkType = in.LinkType
	out.Location = in.Location
	out.Macsec = InterconnectMacsec_ToProto(mapCtx, in.Macsec)
	out.MacsecEnabled = in.MacsecEnabled
	out.Name = in.Name
	out.NocContactEmail = in.NocContactEmail
	out.OperationalStatus = in.OperationalStatus
	out.PeerIpAddress = in.PeerIPAddress
	out.ProvisionedLinkCount = in.ProvisionedLinkCount
	out.RemoteLocation = in.RemoteLocation
	out.RequestedFeatures = in.RequestedFeatures
	out.RequestedLinkCount = in.RequestedLinkCount
	out.SatisfiesPzs = in.SatisfiesPzs
	out.SelfLink = in.SelfLink
	out.State = in.State
	return out
}
func Items_FromProto(mapCtx *direct.MapContext, in *pb.Items) *krm.Items {
	if in == nil {
		return nil
	}
	out := &krm.Items{}
	out.Key = in.Key
	out.Value = in.Value
	return out
}
func Items_ToProto(mapCtx *direct.MapContext, in *krm.Items) *pb.Items {
	if in == nil {
		return nil
	}
	out := &pb.Items{}
	out.Key = in.Key
	out.Value = in.Value
	return out
}
func MetadataFilter_FromProto(mapCtx *direct.MapContext, in *pb.MetadataFilter) *krm.MetadataFilter {
	if in == nil {
		return nil
	}
	out := &krm.MetadataFilter{}
	out.FilterLabels = direct.Slice_FromProto(mapCtx, in.FilterLabels, MetadataFilterLabelMatch_FromProto)
	out.FilterMatchCriteria = in.FilterMatchCriteria
	return out
}
func MetadataFilter_ToProto(mapCtx *direct.MapContext, in *krm.MetadataFilter) *pb.MetadataFilter {
	if in == nil {
		return nil
	}
	out := &pb.MetadataFilter{}
	out.FilterLabels = direct.Slice_ToProto(mapCtx, in.FilterLabels, MetadataFilterLabelMatch_ToProto)
	out.FilterMatchCriteria = in.FilterMatchCriteria
	return out
}
func MetadataFilterLabelMatch_FromProto(mapCtx *direct.MapContext, in *pb.MetadataFilterLabelMatch) *krm.MetadataFilterLabelMatch {
	if in == nil {
		return nil
	}
	out := &krm.MetadataFilterLabelMatch{}
	out.Name = in.Name
	out.Value = in.Value
	return out
}
func MetadataFilterLabelMatch_ToProto(mapCtx *direct.MapContext, in *krm.MetadataFilterLabelMatch) *pb.MetadataFilterLabelMatch {
	if in == nil {
		return nil
	}
	out := &pb.MetadataFilterLabelMatch{}
	out.Name = in.Name
	out.Value = in.Value
	return out
}
func NetworkAttachment_FromProto(mapCtx *direct.MapContext, in *pb.NetworkAttachment) *krmcomputev1alpha1.NetworkAttachment {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.NetworkAttachment{}
	out.ConnectionEndpoints = direct.Slice_FromProto(mapCtx, in.ConnectionEndpoints, NetworkAttachmentConnectedEndpoint_FromProto)
	out.ConnectionPreference = in.ConnectionPreference
	out.CreationTimestamp = in.CreationTimestamp
	out.Description = in.Description
	out.Fingerprint = in.Fingerprint
	out.ID = in.Id
	out.Kind = in.Kind
	out.Name = in.Name
	out.Network = in.Network
	out.ProducerAcceptLists = in.ProducerAcceptLists
	out.ProducerRejectLists = in.ProducerRejectLists
	out.Region = in.Region
	out.SelfLink = in.SelfLink
	out.SelfLinkWithID = in.SelfLinkWithId
	out.Subnetworks = in.Subnetworks
	return out
}
func NetworkAttachment_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.NetworkAttachment) *pb.NetworkAttachment {
	if in == nil {
		return nil
	}
	out := &pb.NetworkAttachment{}
	out.ConnectionEndpoints = direct.Slice_ToProto(mapCtx, in.ConnectionEndpoints, NetworkAttachmentConnectedEndpoint_ToProto)
	out.ConnectionPreference = in.ConnectionPreference
	out.CreationTimestamp = in.CreationTimestamp
	out.Description = in.Description
	out.Fingerprint = in.Fingerprint
	out.Id = in.ID
	out.Kind = in.Kind
	out.Name = in.Name
	out.Network = in.Network
	out.ProducerAcceptLists = in.ProducerAcceptLists
	out.ProducerRejectLists = in.ProducerRejectLists
	out.Region = in.Region
	out.SelfLink = in.SelfLink
	out.SelfLinkWithId = in.SelfLinkWithID
	out.Subnetworks = in.Subnetworks
	return out
}
func NetworkEdgeSecurityService_FromProto(mapCtx *direct.MapContext, in *pb.NetworkEdgeSecurityService) *krmcomputev1alpha1.NetworkEdgeSecurityService {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.NetworkEdgeSecurityService{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Description = in.Description
	out.Fingerprint = in.Fingerprint
	out.ID = in.Id
	out.Kind = in.Kind
	out.Name = in.Name
	out.Region = in.Region
	out.SecurityPolicy = in.SecurityPolicy
	out.SelfLink = in.SelfLink
	out.SelfLinkWithID = in.SelfLinkWithId
	return out
}
func NetworkEdgeSecurityService_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.NetworkEdgeSecurityService) *pb.NetworkEdgeSecurityService {
	if in == nil {
		return nil
	}
	out := &pb.NetworkEdgeSecurityService{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Description = in.Description
	out.Fingerprint = in.Fingerprint
	out.Id = in.ID
	out.Kind = in.Kind
	out.Name = in.Name
	out.Region = in.Region
	out.SecurityPolicy = in.SecurityPolicy
	out.SelfLink = in.SelfLink
	out.SelfLinkWithId = in.SelfLinkWithID
	return out
}
func NetworkInterface_FromProto(mapCtx *direct.MapContext, in *pb.NetworkInterface) *krm.NetworkInterface {
	if in == nil {
		return nil
	}
	out := &krm.NetworkInterface{}
	out.AccessConfigs = direct.Slice_FromProto(mapCtx, in.AccessConfigs, AccessConfig_FromProto)
	out.AliasIPRanges = direct.Slice_FromProto(mapCtx, in.AliasIpRanges, AliasIPRange_FromProto)
	out.Fingerprint = in.Fingerprint
	out.InternalIPV6PrefixLength = in.InternalIpv6PrefixLength
	out.IPV6AccessConfigs = direct.Slice_FromProto(mapCtx, in.Ipv6AccessConfigs, AccessConfig_FromProto)
	out.IPV6AccessType = in.Ipv6AccessType
	out.IPV6Address = in.Ipv6Address
	out.Kind = in.Kind
	out.Name = in.Name
	out.Network = in.Network
	out.NetworkAttachment = in.NetworkAttachment
	out.NetworkIP = in.NetworkIP
	out.NicType = in.NicType
	out.QueueCount = in.QueueCount
	out.StackType = in.StackType
	out.Subnetwork = in.Subnetwork
	return out
}
func NetworkInterface_ToProto(mapCtx *direct.MapContext, in *krm.NetworkInterface) *pb.NetworkInterface {
	if in == nil {
		return nil
	}
	out := &pb.NetworkInterface{}
	out.AccessConfigs = direct.Slice_ToProto(mapCtx, in.AccessConfigs, AccessConfig_ToProto)
	out.AliasIpRanges = direct.Slice_ToProto(mapCtx, in.AliasIPRanges, AliasIPRange_ToProto)
	out.Fingerprint = in.Fingerprint
	out.InternalIpv6PrefixLength = in.InternalIPV6PrefixLength
	out.Ipv6AccessConfigs = direct.Slice_ToProto(mapCtx, in.IPV6AccessConfigs, AccessConfig_ToProto)
	out.Ipv6AccessType = in.IPV6AccessType
	out.Ipv6Address = in.IPV6Address
	out.Kind = in.Kind
	out.Name = in.Name
	out.Network = in.Network
	out.NetworkAttachment = in.NetworkAttachment
	out.NetworkIP = in.NetworkIP
	out.NicType = in.NicType
	out.QueueCount = in.QueueCount
	out.StackType = in.StackType
	out.Subnetwork = in.Subnetwork
	return out
}
func NetworkPerformanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPerformanceConfig) *krm.NetworkPerformanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.NetworkPerformanceConfig{}
	out.TotalEgressBandwidthTier = in.TotalEgressBandwidthTier
	return out
}
func NetworkPerformanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.NetworkPerformanceConfig) *pb.NetworkPerformanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkPerformanceConfig{}
	out.TotalEgressBandwidthTier = in.TotalEgressBandwidthTier
	return out
}
func ReservationAffinity_FromProto(mapCtx *direct.MapContext, in *pb.ReservationAffinity) *krm.ReservationAffinity {
	if in == nil {
		return nil
	}
	out := &krm.ReservationAffinity{}
	out.ConsumeReservationType = in.ConsumeReservationType
	out.Key = in.Key
	out.Values = in.Values
	return out
}
func ReservationAffinity_ToProto(mapCtx *direct.MapContext, in *krm.ReservationAffinity) *pb.ReservationAffinity {
	if in == nil {
		return nil
	}
	out := &pb.ReservationAffinity{}
	out.ConsumeReservationType = in.ConsumeReservationType
	out.Key = in.Key
	out.Values = in.Values
	return out
}
func ResourceStatus_FromProto(mapCtx *direct.MapContext, in *pb.ResourceStatus) *krm.ResourceStatus {
	if in == nil {
		return nil
	}
	out := &krm.ResourceStatus{}
	out.PhysicalHost = in.PhysicalHost
	out.PhysicalHostTopology = ResourceStatusPhysicalHostTopology_FromProto(mapCtx, in.GetPhysicalHostTopology())
	out.Scheduling = ResourceStatusScheduling_FromProto(mapCtx, in.GetScheduling())
	out.UpcomingMaintenance = UpcomingMaintenance_FromProto(mapCtx, in.GetUpcomingMaintenance())
	return out
}
func ResourceStatus_ToProto(mapCtx *direct.MapContext, in *krm.ResourceStatus) *pb.ResourceStatus {
	if in == nil {
		return nil
	}
	out := &pb.ResourceStatus{}
	out.PhysicalHost = in.PhysicalHost
	out.PhysicalHostTopology = ResourceStatusPhysicalHostTopology_ToProto(mapCtx, in.PhysicalHostTopology)
	out.Scheduling = ResourceStatusScheduling_ToProto(mapCtx, in.Scheduling)
	out.UpcomingMaintenance = UpcomingMaintenance_ToProto(mapCtx, in.UpcomingMaintenance)
	return out
}
func ResourceStatusPhysicalHostTopology_FromProto(mapCtx *direct.MapContext, in *pb.ResourceStatusPhysicalHostTopology) *krm.ResourceStatusPhysicalHostTopology {
	if in == nil {
		return nil
	}
	out := &krm.ResourceStatusPhysicalHostTopology{}
	out.Block = in.Block
	out.Cluster = in.Cluster
	out.Host = in.Host
	out.Subblock = in.Subblock
	return out
}
func ResourceStatusPhysicalHostTopology_ToProto(mapCtx *direct.MapContext, in *krm.ResourceStatusPhysicalHostTopology) *pb.ResourceStatusPhysicalHostTopology {
	if in == nil {
		return nil
	}
	out := &pb.ResourceStatusPhysicalHostTopology{}
	out.Block = in.Block
	out.Cluster = in.Cluster
	out.Host = in.Host
	out.Subblock = in.Subblock
	return out
}
func ResourceStatusScheduling_FromProto(mapCtx *direct.MapContext, in *pb.ResourceStatusScheduling) *krm.ResourceStatusScheduling {
	if in == nil {
		return nil
	}
	out := &krm.ResourceStatusScheduling{}
	out.AvailabilityDomain = in.AvailabilityDomain
	return out
}
func ResourceStatusScheduling_ToProto(mapCtx *direct.MapContext, in *krm.ResourceStatusScheduling) *pb.ResourceStatusScheduling {
	if in == nil {
		return nil
	}
	out := &pb.ResourceStatusScheduling{}
	out.AvailabilityDomain = in.AvailabilityDomain
	return out
}
func Scheduling_FromProto(mapCtx *direct.MapContext, in *pb.Scheduling) *krm.Scheduling {
	if in == nil {
		return nil
	}
	out := &krm.Scheduling{}
	out.AutomaticRestart = in.AutomaticRestart
	out.AvailabilityDomain = in.AvailabilityDomain
	out.HostErrorTimeoutSeconds = in.HostErrorTimeoutSeconds
	out.InstanceTerminationAction = in.InstanceTerminationAction
	out.LocalSsdRecoveryTimeout = Duration_FromProto(mapCtx, in.GetLocalSsdRecoveryTimeout())
	out.LocationHint = in.LocationHint
	out.MaxRunDuration = Duration_FromProto(mapCtx, in.GetMaxRunDuration())
	out.MinNodeCpus = in.MinNodeCpus
	out.NodeAffinities = direct.Slice_FromProto(mapCtx, in.NodeAffinities, SchedulingNodeAffinity_FromProto)
	out.OnHostMaintenance = in.OnHostMaintenance
	out.OnInstanceStopAction = SchedulingOnInstanceStopAction_FromProto(mapCtx, in.GetOnInstanceStopAction())
	out.Preemptible = in.Preemptible
	out.ProvisioningModel = in.ProvisioningModel
	out.TerminationTime = in.TerminationTime
	return out
}
func Scheduling_ToProto(mapCtx *direct.MapContext, in *krm.Scheduling) *pb.Scheduling {
	if in == nil {
		return nil
	}
	out := &pb.Scheduling{}
	out.AutomaticRestart = in.AutomaticRestart
	out.AvailabilityDomain = in.AvailabilityDomain
	out.HostErrorTimeoutSeconds = in.HostErrorTimeoutSeconds
	out.InstanceTerminationAction = in.InstanceTerminationAction
	out.LocalSsdRecoveryTimeout = Duration_ToProto(mapCtx, in.LocalSsdRecoveryTimeout)
	out.LocationHint = in.LocationHint
	out.MaxRunDuration = Duration_ToProto(mapCtx, in.MaxRunDuration)
	out.MinNodeCpus = in.MinNodeCpus
	out.NodeAffinities = direct.Slice_ToProto(mapCtx, in.NodeAffinities, SchedulingNodeAffinity_ToProto)
	out.OnHostMaintenance = in.OnHostMaintenance
	out.OnInstanceStopAction = SchedulingOnInstanceStopAction_ToProto(mapCtx, in.OnInstanceStopAction)
	out.Preemptible = in.Preemptible
	out.ProvisioningModel = in.ProvisioningModel
	out.TerminationTime = in.TerminationTime
	return out
}
func SchedulingNodeAffinity_FromProto(mapCtx *direct.MapContext, in *pb.SchedulingNodeAffinity) *krm.SchedulingNodeAffinity {
	if in == nil {
		return nil
	}
	out := &krm.SchedulingNodeAffinity{}
	out.Key = in.Key
	out.Operator = in.Operator
	out.Values = in.Values
	return out
}
func SchedulingNodeAffinity_ToProto(mapCtx *direct.MapContext, in *krm.SchedulingNodeAffinity) *pb.SchedulingNodeAffinity {
	if in == nil {
		return nil
	}
	out := &pb.SchedulingNodeAffinity{}
	out.Key = in.Key
	out.Operator = in.Operator
	out.Values = in.Values
	return out
}
func SchedulingOnInstanceStopAction_FromProto(mapCtx *direct.MapContext, in *pb.SchedulingOnInstanceStopAction) *krm.SchedulingOnInstanceStopAction {
	if in == nil {
		return nil
	}
	out := &krm.SchedulingOnInstanceStopAction{}
	out.DiscardLocalSsd = in.DiscardLocalSsd
	return out
}
func SchedulingOnInstanceStopAction_ToProto(mapCtx *direct.MapContext, in *krm.SchedulingOnInstanceStopAction) *pb.SchedulingOnInstanceStopAction {
	if in == nil {
		return nil
	}
	out := &pb.SchedulingOnInstanceStopAction{}
	out.DiscardLocalSsd = in.DiscardLocalSsd
	return out
}
func ServiceAccount_FromProto(mapCtx *direct.MapContext, in *pb.ServiceAccount) *krm.ServiceAccount {
	if in == nil {
		return nil
	}
	out := &krm.ServiceAccount{}
	out.Email = in.Email
	out.Scopes = in.Scopes
	return out
}
func ServiceAccount_ToProto(mapCtx *direct.MapContext, in *krm.ServiceAccount) *pb.ServiceAccount {
	if in == nil {
		return nil
	}
	out := &pb.ServiceAccount{}
	out.Email = in.Email
	out.Scopes = in.Scopes
	return out
}
func ShieldedInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.ShieldedInstanceConfig) *krm.ShieldedInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.ShieldedInstanceConfig{}
	out.EnableIntegrityMonitoring = in.EnableIntegrityMonitoring
	out.EnableSecureBoot = in.EnableSecureBoot
	out.EnableVTPM = in.EnableVtpm
	return out
}
func ShieldedInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.ShieldedInstanceConfig) *pb.ShieldedInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.ShieldedInstanceConfig{}
	out.EnableIntegrityMonitoring = in.EnableIntegrityMonitoring
	out.EnableSecureBoot = in.EnableSecureBoot
	out.EnableVtpm = in.EnableVTPM
	return out
}
func ShieldedInstanceIntegrityPolicy_FromProto(mapCtx *direct.MapContext, in *pb.ShieldedInstanceIntegrityPolicy) *krm.ShieldedInstanceIntegrityPolicy {
	if in == nil {
		return nil
	}
	out := &krm.ShieldedInstanceIntegrityPolicy{}
	out.UpdateAutoLearnPolicy = in.UpdateAutoLearnPolicy
	return out
}
func ShieldedInstanceIntegrityPolicy_ToProto(mapCtx *direct.MapContext, in *krm.ShieldedInstanceIntegrityPolicy) *pb.ShieldedInstanceIntegrityPolicy {
	if in == nil {
		return nil
	}
	out := &pb.ShieldedInstanceIntegrityPolicy{}
	out.UpdateAutoLearnPolicy = in.UpdateAutoLearnPolicy
	return out
}
func SubnetworkLogConfig_FromProto(mapCtx *direct.MapContext, in *pb.SubnetworkLogConfig) *krm.SubnetworkLogConfig {
	if in == nil {
		return nil
	}
	out := &krm.SubnetworkLogConfig{}
	out.AggregationInterval = in.AggregationInterval
	// MISSING: Enable
	out.FilterExpr = in.FilterExpr
	out.FlowSampling = in.FlowSampling
	out.Metadata = in.Metadata
	out.MetadataFields = in.MetadataFields
	return out
}
func SubnetworkLogConfig_ToProto(mapCtx *direct.MapContext, in *krm.SubnetworkLogConfig) *pb.SubnetworkLogConfig {
	if in == nil {
		return nil
	}
	out := &pb.SubnetworkLogConfig{}
	out.AggregationInterval = in.AggregationInterval
	// MISSING: Enable
	out.FilterExpr = in.FilterExpr
	out.FlowSampling = in.FlowSampling
	out.Metadata = in.Metadata
	out.MetadataFields = in.MetadataFields
	return out
}
func SubnetworkParams_FromProto(mapCtx *direct.MapContext, in *pb.SubnetworkParams) *krm.SubnetworkParams {
	if in == nil {
		return nil
	}
	out := &krm.SubnetworkParams{}
	out.ResourceManagerTags = in.ResourceManagerTags
	return out
}
func SubnetworkParams_ToProto(mapCtx *direct.MapContext, in *krm.SubnetworkParams) *pb.SubnetworkParams {
	if in == nil {
		return nil
	}
	out := &pb.SubnetworkParams{}
	out.ResourceManagerTags = in.ResourceManagerTags
	return out
}
func SubnetworkSecondaryRange_FromProto(mapCtx *direct.MapContext, in *pb.SubnetworkSecondaryRange) *krm.SubnetworkSecondaryRange {
	if in == nil {
		return nil
	}
	out := &krm.SubnetworkSecondaryRange{}
	out.IPCIDRRange = in.IpCidrRange
	out.RangeName = in.RangeName
	// MISSING: ReservedInternalRange
	return out
}
func SubnetworkSecondaryRange_ToProto(mapCtx *direct.MapContext, in *krm.SubnetworkSecondaryRange) *pb.SubnetworkSecondaryRange {
	if in == nil {
		return nil
	}
	out := &pb.SubnetworkSecondaryRange{}
	out.IpCidrRange = in.IPCIDRRange
	out.RangeName = in.RangeName
	// MISSING: ReservedInternalRange
	return out
}
func Tags_FromProto(mapCtx *direct.MapContext, in *pb.Tags) *krm.Tags {
	if in == nil {
		return nil
	}
	out := &krm.Tags{}
	out.Fingerprint = in.Fingerprint
	out.Items = in.Items
	return out
}
func Tags_ToProto(mapCtx *direct.MapContext, in *krm.Tags) *pb.Tags {
	if in == nil {
		return nil
	}
	out := &pb.Tags{}
	out.Fingerprint = in.Fingerprint
	out.Items = in.Items
	return out
}
func UpcomingMaintenance_FromProto(mapCtx *direct.MapContext, in *pb.UpcomingMaintenance) *krm.UpcomingMaintenance {
	if in == nil {
		return nil
	}
	out := &krm.UpcomingMaintenance{}
	out.CanReschedule = in.CanReschedule
	out.LatestWindowStartTime = in.LatestWindowStartTime
	out.MaintenanceOnShutdown = in.MaintenanceOnShutdown
	out.MaintenanceReasons = in.MaintenanceReasons
	out.MaintenanceStatus = in.MaintenanceStatus
	out.Type = in.Type
	out.WindowEndTime = in.WindowEndTime
	out.WindowStartTime = in.WindowStartTime
	return out
}
func UpcomingMaintenance_ToProto(mapCtx *direct.MapContext, in *krm.UpcomingMaintenance) *pb.UpcomingMaintenance {
	if in == nil {
		return nil
	}
	out := &pb.UpcomingMaintenance{}
	out.CanReschedule = in.CanReschedule
	out.LatestWindowStartTime = in.LatestWindowStartTime
	out.MaintenanceOnShutdown = in.MaintenanceOnShutdown
	out.MaintenanceReasons = in.MaintenanceReasons
	out.MaintenanceStatus = in.MaintenanceStatus
	out.Type = in.Type
	out.WindowEndTime = in.WindowEndTime
	out.WindowStartTime = in.WindowStartTime
	return out
}
