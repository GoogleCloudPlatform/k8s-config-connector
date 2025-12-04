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

func FirewallPolicyRuleMatcher_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyRuleMatcher) *krm.FirewallPolicyRuleMatcher {
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
	out.Layer4Configs = direct.Slice_FromProto(mapCtx, in.Layer4Configs, FirewallPolicyRuleMatcherLayer4Config_v1beta1_FromProto)
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
func FirewallPolicyRuleMatcher_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.FirewallPolicyRuleMatcher) *pb.FirewallPolicyRuleMatcher {
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
	out.Layer4Configs = direct.Slice_ToProto(mapCtx, in.Layer4Configs, FirewallPolicyRuleMatcherLayer4Config_v1beta1_ToProto)
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
func FirewallPolicyRuleSecureTag_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyRuleSecureTag) *krm.FirewallPolicyRuleSecureTag {
	if in == nil {
		return nil
	}
	out := &krm.FirewallPolicyRuleSecureTag{}
	out.Name = in.Name
	out.State = in.State
	return out
}
func FirewallPolicyRuleSecureTag_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.FirewallPolicyRuleSecureTag) *pb.FirewallPolicyRuleSecureTag {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicyRuleSecureTag{}
	out.Name = in.Name
	out.State = in.State
	return out
}
func ForwardingruleServiceDirectoryRegistrations_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ForwardingRuleServiceDirectoryRegistration) *krm.ForwardingruleServiceDirectoryRegistrations {
	if in == nil {
		return nil
	}
	out := &krm.ForwardingruleServiceDirectoryRegistrations{}
	out.Namespace = in.Namespace
	out.Service = in.Service
	// MISSING: ServiceDirectoryRegion
	return out
}
func ForwardingruleServiceDirectoryRegistrations_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ForwardingruleServiceDirectoryRegistrations) *pb.ForwardingRuleServiceDirectoryRegistration {
	if in == nil {
		return nil
	}
	out := &pb.ForwardingRuleServiceDirectoryRegistration{}
	out.Namespace = in.Namespace
	out.Service = in.Service
	// MISSING: ServiceDirectoryRegion
	return out
}
func InterconnectCircuitInfo_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.InterconnectCircuitInfo) *krmcomputev1alpha1.InterconnectCircuitInfo {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.InterconnectCircuitInfo{}
	out.CustomerDemarcID = in.CustomerDemarcId
	out.GoogleCircuitID = in.GoogleCircuitId
	out.GoogleDemarcID = in.GoogleDemarcId
	return out
}
func InterconnectCircuitInfo_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.InterconnectCircuitInfo) *pb.InterconnectCircuitInfo {
	if in == nil {
		return nil
	}
	out := &pb.InterconnectCircuitInfo{}
	out.CustomerDemarcId = in.CustomerDemarcID
	out.GoogleCircuitId = in.GoogleCircuitID
	out.GoogleDemarcId = in.GoogleDemarcID
	return out
}
func InterconnectMacsec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.InterconnectMacsec) *krmcomputev1alpha1.InterconnectMacsec {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.InterconnectMacsec{}
	out.FailOpen = in.FailOpen
	out.PreSharedKeys = direct.Slice_FromProto(mapCtx, in.PreSharedKeys, InterconnectMacsecPreSharedKey_v1alpha1_FromProto)
	return out
}
func InterconnectMacsec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.InterconnectMacsec) *pb.InterconnectMacsec {
	if in == nil {
		return nil
	}
	out := &pb.InterconnectMacsec{}
	out.FailOpen = in.FailOpen
	out.PreSharedKeys = direct.Slice_ToProto(mapCtx, in.PreSharedKeys, InterconnectMacsecPreSharedKey_v1alpha1_ToProto)
	return out
}
func InterconnectMacsecPreSharedKey_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.InterconnectMacsecPreSharedKey) *krmcomputev1alpha1.InterconnectMacsecPreSharedKey {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.InterconnectMacsecPreSharedKey{}
	out.Name = in.Name
	out.StartTime = in.StartTime
	return out
}
func InterconnectMacsecPreSharedKey_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.InterconnectMacsecPreSharedKey) *pb.InterconnectMacsecPreSharedKey {
	if in == nil {
		return nil
	}
	out := &pb.InterconnectMacsecPreSharedKey{}
	out.Name = in.Name
	out.StartTime = in.StartTime
	return out
}
func InterconnectOutageNotification_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.InterconnectOutageNotification) *krmcomputev1alpha1.InterconnectOutageNotification {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.InterconnectOutageNotification{}
	out.AffectedCircuits = in.AffectedCircuits
	out.Description = in.Description
	out.EndTime = in.EndTime
	out.IssueType = in.IssueType
	out.Name = in.Name
	out.Source = in.Source
	out.StartTime = in.StartTime
	out.State = in.State
	return out
}
func InterconnectOutageNotification_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.InterconnectOutageNotification) *pb.InterconnectOutageNotification {
	if in == nil {
		return nil
	}
	out := &pb.InterconnectOutageNotification{}
	out.AffectedCircuits = in.AffectedCircuits
	out.Description = in.Description
	out.EndTime = in.EndTime
	out.IssueType = in.IssueType
	out.Name = in.Name
	out.Source = in.Source
	out.StartTime = in.StartTime
	out.State = in.State
	return out
}
func MetadataFilter_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.MetadataFilter) *krm.MetadataFilter {
	if in == nil {
		return nil
	}
	out := &krm.MetadataFilter{}
	out.FilterLabels = direct.Slice_FromProto(mapCtx, in.FilterLabels, MetadataFilterLabelMatch_v1beta1_FromProto)
	out.FilterMatchCriteria = in.FilterMatchCriteria
	return out
}
func MetadataFilter_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.MetadataFilter) *pb.MetadataFilter {
	if in == nil {
		return nil
	}
	out := &pb.MetadataFilter{}
	out.FilterLabels = direct.Slice_ToProto(mapCtx, in.FilterLabels, MetadataFilterLabelMatch_v1beta1_ToProto)
	out.FilterMatchCriteria = in.FilterMatchCriteria
	return out
}
func MetadataFilterLabelMatch_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.MetadataFilterLabelMatch) *krm.MetadataFilterLabelMatch {
	if in == nil {
		return nil
	}
	out := &krm.MetadataFilterLabelMatch{}
	out.Name = in.Name
	out.Value = in.Value
	return out
}
func MetadataFilterLabelMatch_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.MetadataFilterLabelMatch) *pb.MetadataFilterLabelMatch {
	if in == nil {
		return nil
	}
	out := &pb.MetadataFilterLabelMatch{}
	out.Name = in.Name
	out.Value = in.Value
	return out
}
func NetworkAttachmentConnectedEndpoint_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.NetworkAttachmentConnectedEndpoint) *krmcomputev1alpha1.NetworkAttachmentConnectedEndpoint {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.NetworkAttachmentConnectedEndpoint{}
	out.IPAddress = in.IpAddress
	out.IPV6Address = in.Ipv6Address
	out.ProjectIDOrNum = in.ProjectIdOrNum
	out.SecondaryIPCIDRRanges = in.SecondaryIpCidrRanges
	out.Status = in.Status
	out.Subnetwork = in.Subnetwork
	out.SubnetworkCIDRRange = in.SubnetworkCidrRange
	return out
}
func NetworkAttachmentConnectedEndpoint_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.NetworkAttachmentConnectedEndpoint) *pb.NetworkAttachmentConnectedEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.NetworkAttachmentConnectedEndpoint{}
	out.IpAddress = in.IPAddress
	out.Ipv6Address = in.IPV6Address
	out.ProjectIdOrNum = in.ProjectIDOrNum
	out.SecondaryIpCidrRanges = in.SecondaryIPCIDRRanges
	out.Status = in.Status
	out.Subnetwork = in.Subnetwork
	out.SubnetworkCidrRange = in.SubnetworkCIDRRange
	return out
}
