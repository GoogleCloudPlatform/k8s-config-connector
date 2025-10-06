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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeFirewallPolicyRuleSpec_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyRule) *krm.ComputeFirewallPolicyRuleSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeFirewallPolicyRuleSpec{}
	out.Action = in.Action
	out.Description = in.Description
	out.Direction = in.Direction
	out.Disabled = in.Disabled
	out.EnableLogging = in.EnableLogging
	// MISSING: Kind
	out.Match = FirewallPolicyRuleMatch_FromProto(mapCtx, in.GetMatch())
	out.Priority = in.Priority
	// MISSING: RuleName
	// MISSING: RuleTupleCount
	// MISSING: SecurityProfileGroup
	out.TargetResources = ComputeFirewallPolicyRuleSpec_TargetResources_FromProto(mapCtx, in.TargetResources)
	// MISSING: TargetSecureTags
	out.TargetServiceAccounts = ComputeFirewallPolicyRuleSpec_TargetServiceAccounts_FromProto(mapCtx, in.TargetServiceAccounts)
	// MISSING: TLSInspect
	return out
}
func ComputeFirewallPolicyRuleSpec_ToProto(mapCtx *direct.MapContext, in *krm.ComputeFirewallPolicyRuleSpec) *pb.FirewallPolicyRule {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicyRule{}
	out.Action = in.Action
	out.Description = in.Description
	out.Direction = in.Direction
	out.Disabled = in.Disabled
	out.EnableLogging = in.EnableLogging
	// MISSING: Kind
	out.Match = FirewallPolicyRuleMatch_ToProto(mapCtx, in.Match)
	out.Priority = in.Priority
	// MISSING: RuleName
	// MISSING: RuleTupleCount
	// MISSING: SecurityProfileGroup
	out.TargetResources = ComputeFirewallPolicyRuleSpec_TargetResources_ToProto(mapCtx, in.TargetResources)
	// MISSING: TargetSecureTags
	out.TargetServiceAccounts = ComputeFirewallPolicyRuleSpec_TargetServiceAccounts_ToProto(mapCtx, in.TargetServiceAccounts)
	// MISSING: TLSInspect
	return out
}
func ComputeFirewallPolicyRuleStatus_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyRule) *krm.ComputeFirewallPolicyRuleStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeFirewallPolicyRuleStatus{}
	// MISSING: Action
	// MISSING: Description
	// MISSING: Direction
	// MISSING: Disabled
	// MISSING: EnableLogging
	out.Kind = in.Kind
	// MISSING: Match
	// MISSING: Priority
	// MISSING: RuleName
	out.RuleTupleCount = in.RuleTupleCount
	// MISSING: SecurityProfileGroup
	// MISSING: TargetResources
	// MISSING: TargetSecureTags
	// MISSING: TargetServiceAccounts
	// MISSING: TLSInspect
	return out
}
func ComputeFirewallPolicyRuleStatus_ToProto(mapCtx *direct.MapContext, in *krm.ComputeFirewallPolicyRuleStatus) *pb.FirewallPolicyRule {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicyRule{}
	// MISSING: Action
	// MISSING: Description
	// MISSING: Direction
	// MISSING: Disabled
	// MISSING: EnableLogging
	out.Kind = in.Kind
	// MISSING: Match
	// MISSING: Priority
	// MISSING: RuleName
	out.RuleTupleCount = in.RuleTupleCount
	// MISSING: SecurityProfileGroup
	// MISSING: TargetResources
	// MISSING: TargetSecureTags
	// MISSING: TargetServiceAccounts
	// MISSING: TLSInspect
	return out
}
func ComputeForwardingRuleSpec_FromProto(mapCtx *direct.MapContext, in *pb.ForwardingRule) *krm.ComputeForwardingRuleSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeForwardingRuleSpec{}
	// MISSING: IPAddress
	// (near miss): "IPAddress" vs "IpAddress"
	// MISSING: IPProtocol
	// (near miss): "IPProtocol" vs "IpProtocol"
	out.AllPorts = in.AllPorts
	out.AllowGlobalAccess = in.AllowGlobalAccess
	// MISSING: AllowPSCGlobalAccess
	// (near miss): "AllowPSCGlobalAccess" vs "AllowPscGlobalAccess"
	if in.GetBackendService() != "" {
		out.BackendServiceRef = &krm.ComputeBackendServiceRef{External: in.GetBackendService()}
	}
	// MISSING: BaseForwardingRule
	// MISSING: CreationTimestamp
	out.Description = in.Description
	// MISSING: ExternalManagedBackendBucketMigrationState
	// MISSING: ExternalManagedBackendBucketMigrationTestingPercentage
	// MISSING: Fingerprint
	// MISSING: ID
	// MISSING: IPCollection
	// MISSING: IPVersion
	// (near miss): "IPVersion" vs "IpVersion"
	out.IsMirroringCollector = in.IsMirroringCollector
	// MISSING: Kind
	// MISSING: LabelFingerprint
	// MISSING: Labels
	out.LoadBalancingScheme = in.LoadBalancingScheme
	out.MetadataFilters = direct.Slice_FromProto(mapCtx, in.MetadataFilters, ForwardingruleMetadataFilters_FromProto)
	// MISSING: Name
	if in.GetNetwork() != "" {
		out.NetworkRef = &refsv1beta1.ComputeNetworkRef{External: in.GetNetwork()}
	}
	out.NetworkTier = in.NetworkTier
	// MISSING: NoAutomateDNSZone
	// (near miss): "NoAutomateDNSZone" vs "NoAutomateDnsZone"
	out.PortRange = in.PortRange
	out.Ports = in.Ports
	// MISSING: PSCConnectionID
	// MISSING: PSCConnectionStatus
	// MISSING: Region
	// MISSING: SelfLink
	// MISSING: SelfLinkWithID
	out.ServiceDirectoryRegistrations = direct.Slice_FromProto(mapCtx, in.ServiceDirectoryRegistrations, ForwardingruleServiceDirectoryRegistrations_FromProto)
	out.ServiceLabel = in.ServiceLabel
	// MISSING: ServiceName
	// MISSING: SourceIPRanges
	// (near miss): "SourceIPRanges" vs "SourceIpRanges"
	if in.GetSubnetwork() != "" {
		out.SubnetworkRef = &refsv1beta1.ComputeSubnetworkRef{External: in.GetSubnetwork()}
	}
	out.Target = in.Target
	return out
}
func ComputeForwardingRuleSpec_ToProto(mapCtx *direct.MapContext, in *krm.ComputeForwardingRuleSpec) *pb.ForwardingRule {
	if in == nil {
		return nil
	}
	out := &pb.ForwardingRule{}
	// MISSING: IPAddress
	// (near miss): "IPAddress" vs "IpAddress"
	// MISSING: IPProtocol
	// (near miss): "IPProtocol" vs "IpProtocol"
	out.AllPorts = in.AllPorts
	out.AllowGlobalAccess = in.AllowGlobalAccess
	// MISSING: AllowPSCGlobalAccess
	// (near miss): "AllowPSCGlobalAccess" vs "AllowPscGlobalAccess"
	if in.BackendServiceRef != nil {
		out.BackendService = in.BackendServiceRef.External
	}
	// MISSING: BaseForwardingRule
	// MISSING: CreationTimestamp
	out.Description = in.Description
	// MISSING: ExternalManagedBackendBucketMigrationState
	// MISSING: ExternalManagedBackendBucketMigrationTestingPercentage
	// MISSING: Fingerprint
	// MISSING: ID
	// MISSING: IPCollection
	// MISSING: IPVersion
	// (near miss): "IPVersion" vs "IpVersion"
	out.IsMirroringCollector = in.IsMirroringCollector
	// MISSING: Kind
	// MISSING: LabelFingerprint
	// MISSING: Labels
	out.LoadBalancingScheme = in.LoadBalancingScheme
	out.MetadataFilters = direct.Slice_ToProto(mapCtx, in.MetadataFilters, ForwardingruleMetadataFilters_ToProto)
	// MISSING: Name
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	out.NetworkTier = in.NetworkTier
	// MISSING: NoAutomateDNSZone
	// (near miss): "NoAutomateDNSZone" vs "NoAutomateDnsZone"
	out.PortRange = in.PortRange
	out.Ports = in.Ports
	// MISSING: PSCConnectionID
	// MISSING: PSCConnectionStatus
	// MISSING: Region
	// MISSING: SelfLink
	// MISSING: SelfLinkWithID
	out.ServiceDirectoryRegistrations = direct.Slice_ToProto(mapCtx, in.ServiceDirectoryRegistrations, ForwardingruleServiceDirectoryRegistrations_ToProto)
	out.ServiceLabel = in.ServiceLabel
	// MISSING: ServiceName
	// MISSING: SourceIPRanges
	// (near miss): "SourceIPRanges" vs "SourceIpRanges"
	if in.SubnetworkRef != nil {
		out.Subnetwork = in.SubnetworkRef.External
	}
	out.Target = in.Target
	return out
}
func ComputeForwardingRuleStatus_FromProto(mapCtx *direct.MapContext, in *pb.ForwardingRule) *krm.ComputeForwardingRuleStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeForwardingRuleStatus{}
	// MISSING: IPAddress
	// MISSING: IPProtocol
	// MISSING: AllPorts
	// MISSING: AllowGlobalAccess
	// MISSING: AllowPSCGlobalAccess
	// MISSING: BackendService
	out.BaseForwardingRule = in.BaseForwardingRule
	out.CreationTimestamp = in.CreationTimestamp
	// MISSING: Description
	// MISSING: ExternalManagedBackendBucketMigrationState
	// MISSING: ExternalManagedBackendBucketMigrationTestingPercentage
	// MISSING: Fingerprint
	// MISSING: ID
	// MISSING: IPCollection
	// MISSING: IPVersion
	// MISSING: IsMirroringCollector
	// MISSING: Kind
	out.LabelFingerprint = in.LabelFingerprint
	// MISSING: Labels
	// MISSING: LoadBalancingScheme
	// MISSING: MetadataFilters
	// MISSING: Name
	// MISSING: Network
	// MISSING: NetworkTier
	// MISSING: NoAutomateDNSZone
	// MISSING: PortRange
	// MISSING: Ports
	// MISSING: PSCConnectionID
	// (near miss): "PSCConnectionID" vs "PscConnectionId"
	// MISSING: PSCConnectionStatus
	// (near miss): "PSCConnectionStatus" vs "PscConnectionStatus"
	// MISSING: Region
	out.SelfLink = in.SelfLink
	// MISSING: SelfLinkWithID
	// MISSING: ServiceDirectoryRegistrations
	// MISSING: ServiceLabel
	out.ServiceName = in.ServiceName
	// MISSING: SourceIPRanges
	// MISSING: Subnetwork
	// MISSING: Target
	return out
}
func ComputeForwardingRuleStatus_ToProto(mapCtx *direct.MapContext, in *krm.ComputeForwardingRuleStatus) *pb.ForwardingRule {
	if in == nil {
		return nil
	}
	out := &pb.ForwardingRule{}
	// MISSING: IPAddress
	// MISSING: IPProtocol
	// MISSING: AllPorts
	// MISSING: AllowGlobalAccess
	// MISSING: AllowPSCGlobalAccess
	// MISSING: BackendService
	out.BaseForwardingRule = in.BaseForwardingRule
	out.CreationTimestamp = in.CreationTimestamp
	// MISSING: Description
	// MISSING: ExternalManagedBackendBucketMigrationState
	// MISSING: ExternalManagedBackendBucketMigrationTestingPercentage
	// MISSING: Fingerprint
	// MISSING: ID
	// MISSING: IPCollection
	// MISSING: IPVersion
	// MISSING: IsMirroringCollector
	// MISSING: Kind
	out.LabelFingerprint = in.LabelFingerprint
	// MISSING: Labels
	// MISSING: LoadBalancingScheme
	// MISSING: MetadataFilters
	// MISSING: Name
	// MISSING: Network
	// MISSING: NetworkTier
	// MISSING: NoAutomateDNSZone
	// MISSING: PortRange
	// MISSING: Ports
	// MISSING: PSCConnectionID
	// (near miss): "PSCConnectionID" vs "PscConnectionId"
	// MISSING: PSCConnectionStatus
	// (near miss): "PSCConnectionStatus" vs "PscConnectionStatus"
	// MISSING: Region
	out.SelfLink = in.SelfLink
	// MISSING: SelfLinkWithID
	// MISSING: ServiceDirectoryRegistrations
	// MISSING: ServiceLabel
	out.ServiceName = in.ServiceName
	// MISSING: SourceIPRanges
	// MISSING: Subnetwork
	// MISSING: Target
	return out
}
func FirewallPolicyRuleLayer4Configs_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyRuleMatcherLayer4Config) *krm.FirewallPolicyRuleLayer4Configs {
	if in == nil {
		return nil
	}
	out := &krm.FirewallPolicyRuleLayer4Configs{}
	out.IPProtocol = in.IpProtocol
	out.Ports = in.Ports
	return out
}
func FirewallPolicyRuleLayer4Configs_ToProto(mapCtx *direct.MapContext, in *krm.FirewallPolicyRuleLayer4Configs) *pb.FirewallPolicyRuleMatcherLayer4Config {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicyRuleMatcherLayer4Config{}
	out.IpProtocol = in.IPProtocol
	out.Ports = in.Ports
	return out
}
func FirewallPolicyRuleMatch_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyRuleMatcher) *krm.FirewallPolicyRuleMatch {
	if in == nil {
		return nil
	}
	out := &krm.FirewallPolicyRuleMatch{}
	out.DestAddressGroups = in.DestAddressGroups
	out.DestFqdns = in.DestFqdns
	out.DestIPRanges = in.DestIpRanges
	// MISSING: DestNetworkType
	out.DestRegionCodes = in.DestRegionCodes
	out.DestThreatIntelligences = in.DestThreatIntelligences
	out.Layer4Configs = direct.Slice_FromProto(mapCtx, in.Layer4Configs, FirewallPolicyRuleLayer4Configs_FromProto)
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
func FirewallPolicyRuleMatch_ToProto(mapCtx *direct.MapContext, in *krm.FirewallPolicyRuleMatch) *pb.FirewallPolicyRuleMatcher {
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
	out.Layer4Configs = direct.Slice_ToProto(mapCtx, in.Layer4Configs, FirewallPolicyRuleLayer4Configs_ToProto)
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
func ForwardingruleFilterLabels_FromProto(mapCtx *direct.MapContext, in *pb.MetadataFilterLabelMatch) *krm.ForwardingruleFilterLabels {
	if in == nil {
		return nil
	}
	out := &krm.ForwardingruleFilterLabels{}
	out.Name = in.Name
	out.Value = in.Value
	return out
}
func ForwardingruleFilterLabels_ToProto(mapCtx *direct.MapContext, in *krm.ForwardingruleFilterLabels) *pb.MetadataFilterLabelMatch {
	if in == nil {
		return nil
	}
	out := &pb.MetadataFilterLabelMatch{}
	out.Name = in.Name
	out.Value = in.Value
	return out
}
func ForwardingruleMetadataFilters_FromProto(mapCtx *direct.MapContext, in *pb.MetadataFilter) *krm.ForwardingruleMetadataFilters {
	if in == nil {
		return nil
	}
	out := &krm.ForwardingruleMetadataFilters{}
	out.FilterLabels = direct.Slice_FromProto(mapCtx, in.FilterLabels, ForwardingruleFilterLabels_FromProto)
	out.FilterMatchCriteria = in.FilterMatchCriteria
	return out
}
func ForwardingruleMetadataFilters_ToProto(mapCtx *direct.MapContext, in *krm.ForwardingruleMetadataFilters) *pb.MetadataFilter {
	if in == nil {
		return nil
	}
	out := &pb.MetadataFilter{}
	out.FilterLabels = direct.Slice_ToProto(mapCtx, in.FilterLabels, ForwardingruleFilterLabels_ToProto)
	out.FilterMatchCriteria = in.FilterMatchCriteria
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
