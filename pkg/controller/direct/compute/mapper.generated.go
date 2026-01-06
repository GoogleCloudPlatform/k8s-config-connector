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

func ComputeSubnetworkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Subnetwork) *krm.ComputeSubnetworkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ComputeSubnetworkObservedState{}
	// MISSING: CreationTimestamp
	// MISSING: EnableFlowLogs
	// MISSING: ExternalIPV6Prefix
	// MISSING: Fingerprint
	// MISSING: GatewayAddress
	// MISSING: ID
	// MISSING: InternalIPV6Prefix
	// MISSING: IPCollection
	// MISSING: IPV6CIDRRange
	// MISSING: IPV6GCEEndpoint
	// MISSING: Kind
	// MISSING: Name
	// MISSING: Params
	// MISSING: ReservedInternalRange
	// MISSING: SelfLink
	// MISSING: State
	// MISSING: SystemReservedExternalIPV6Ranges
	// MISSING: SystemReservedInternalIPV6Ranges
	return out
}
func ComputeSubnetworkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ComputeSubnetworkObservedState) *pb.Subnetwork {
	if in == nil {
		return nil
	}
	out := &pb.Subnetwork{}
	// MISSING: CreationTimestamp
	// MISSING: EnableFlowLogs
	// MISSING: ExternalIPV6Prefix
	// MISSING: Fingerprint
	// MISSING: GatewayAddress
	// MISSING: ID
	// MISSING: InternalIPV6Prefix
	// MISSING: IPCollection
	// MISSING: IPV6CIDRRange
	// MISSING: IPV6GCEEndpoint
	// MISSING: Kind
	// MISSING: Name
	// MISSING: Params
	// MISSING: ReservedInternalRange
	// MISSING: SelfLink
	// MISSING: State
	// MISSING: SystemReservedExternalIPV6Ranges
	// MISSING: SystemReservedInternalIPV6Ranges
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
		out.Network = in.NetworkRef.External
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
