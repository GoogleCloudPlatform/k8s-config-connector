// Copyright 2024 Google LLC
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
)

func ComputeForwardingRuleSpec_FromProto(mapCtx *direct.MapContext, in *pb.ForwardingRule) *krm.ComputeForwardingRuleSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeForwardingRuleSpec{}
	out.IpAddress = ComputeForwardingRuleSpec_IpAddress_FromProto(mapCtx, in.GetIPAddress())
	out.IpProtocol = in.IPProtocol
	out.AllPorts = in.AllPorts
	out.AllowGlobalAccess = in.AllowGlobalAccess
	out.AllowPscGlobalAccess = in.AllowPscGlobalAccess
	out.BackendServiceRef = ComputeForwardingRuleSpec_BackendSeriviceRef_FromProto(mapCtx, in.GetBackendService())
	// MISSING: BaseForwardingRule
	// MISSING: CreationTimestamp
	out.Description = in.Description
	// MISSING: Fingerprint
	// MISSING: Id
	// MISSING: IpCollection
	out.IpVersion = in.IpVersion
	out.IsMirroringCollector = in.IsMirroringCollector
	// MISSING: Kind
	// MISSING: LabelFingerprint
	// MISSING: Labels
	out.LoadBalancingScheme = in.LoadBalancingScheme
	out.MetadataFilters = direct.Slice_FromProto(mapCtx, in.MetadataFilters, ForwardingruleMetadataFilters_FromProto)
	// MISSING: Name
	out.NetworkRef = ComputeForwardingRuleSpec_NetworkRef_FromProto(mapCtx, in.GetNetwork())
	out.NetworkTier = in.NetworkTier
	out.NoAutomateDnsZone = in.NoAutomateDnsZone
	out.PortRange = in.PortRange
	out.Ports = in.Ports
	// MISSING: PscConnectionId
	// MISSING: PscConnectionStatus
	// MISSING: Region
	// MISSING: SelfLink
	out.ServiceDirectoryRegistrations = direct.Slice_FromProto(mapCtx, in.ServiceDirectoryRegistrations, ForwardingruleServiceDirectoryRegistrations_FromProto)
	out.ServiceLabel = in.ServiceLabel
	// MISSING: ServiceName
	out.SourceIpRanges = in.SourceIpRanges
	out.SubnetworkRef = ComputeForwardingRuleSpec_SubnetworkRef_FromProto(mapCtx, in.GetSubnetwork())
	out.Target = ComputeForwardingRuleSpec_Target_FromProto(mapCtx, in.GetTarget())
	return out
}
func ComputeForwardingRuleSpec_ToProto(mapCtx *direct.MapContext, in *krm.ComputeForwardingRuleSpec) *pb.ForwardingRule {
	if in == nil {
		return nil
	}
	out := &pb.ForwardingRule{}
	out.IPAddress = ComputeForwardingRuleSpec_IpAddress_ToProto(mapCtx, in.IpAddress)
	out.IPProtocol = in.IpProtocol
	out.AllPorts = in.AllPorts
	out.AllowGlobalAccess = in.AllowGlobalAccess
	out.AllowPscGlobalAccess = in.AllowPscGlobalAccess
	out.BackendService = ComputeForwardingRuleSpec_BackendSeriviceRef_ToProto(mapCtx, in.BackendServiceRef)
	// MISSING: BaseForwardingRule
	// MISSING: CreationTimestamp
	out.Description = in.Description
	// MISSING: Fingerprint
	// MISSING: Id
	// MISSING: IpCollection
	out.IpVersion = in.IpVersion
	out.IsMirroringCollector = in.IsMirroringCollector
	// MISSING: Kind
	// MISSING: LabelFingerprint
	// MISSING: Labels
	out.LoadBalancingScheme = in.LoadBalancingScheme
	out.MetadataFilters = direct.Slice_ToProto(mapCtx, in.MetadataFilters, ForwardingruleMetadataFilters_ToProto)
	// MISSING: Name
	out.Network = ComputeForwardingRuleSpec_NetworkRef_ToProto(mapCtx, in.NetworkRef)
	out.NetworkTier = in.NetworkTier
	out.NoAutomateDnsZone = in.NoAutomateDnsZone
	out.PortRange = in.PortRange
	out.Ports = in.Ports
	// MISSING: PscConnectionId
	// MISSING: PscConnectionStatus
	// MISSING: Region
	// MISSING: SelfLink
	out.ServiceDirectoryRegistrations = direct.Slice_ToProto(mapCtx, in.ServiceDirectoryRegistrations, ForwardingruleServiceDirectoryRegistrations_ToProto)
	out.ServiceLabel = in.ServiceLabel
	// MISSING: ServiceName
	out.SourceIpRanges = in.SourceIpRanges
	out.Subnetwork = ComputeForwardingRuleSpec_SubnetworkRef_ToProto(mapCtx, in.SubnetworkRef)
	out.Target = ComputeForwardingRuleSpec_Target_ToProto(mapCtx, in.Target)
	return out
}
func ComputeForwardingRuleStatus_FromProto(mapCtx *direct.MapContext, in *pb.ForwardingRule) *krm.ComputeForwardingRuleStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeForwardingRuleStatus{}
	out.BaseForwardingRule = in.BaseForwardingRule
	out.CreationTimestamp = in.CreationTimestamp
	out.LabelFingerprint = in.LabelFingerprint
	out.PscConnectionId = ComputeForwardingRuleStatus_PscConnectionId_FromProto(mapCtx, in.GetPscConnectionId())
	out.PscConnectionStatus = in.PscConnectionStatus
	out.SelfLink = in.SelfLink
	out.ServiceName = in.ServiceName
	return out
}
func ComputeForwardingRuleStatus_ToProto(mapCtx *direct.MapContext, in *krm.ComputeForwardingRuleStatus) *pb.ForwardingRule {
	if in == nil {
		return nil
	}
	out := &pb.ForwardingRule{}
	out.BaseForwardingRule = in.BaseForwardingRule
	out.CreationTimestamp = in.CreationTimestamp
	out.LabelFingerprint = in.LabelFingerprint
	out.PscConnectionId = ComputeForwardingRuleStatus_PscConnectionId_ToProto(mapCtx, in.PscConnectionId)
	out.PscConnectionStatus = in.PscConnectionStatus
	out.SelfLink = in.SelfLink
	out.ServiceName = in.ServiceName
	return out
}
func ForwardingruleFilterLabels_FromProto(mapCtx *direct.MapContext, in *pb.MetadataFilterLabelMatch) *krm.ForwardingruleFilterLabels {
	if in == nil {
		return nil
	}
	out := &krm.ForwardingruleFilterLabels{}
	out.Name = in.GetName()
	out.Value = in.GetValue()
	return out
}
func ForwardingruleFilterLabels_ToProto(mapCtx *direct.MapContext, in *krm.ForwardingruleFilterLabels) *pb.MetadataFilterLabelMatch {
	if in == nil {
		return nil
	}
	out := &pb.MetadataFilterLabelMatch{}
	out.Name = direct.LazyPtr(in.Name)
	out.Value = direct.LazyPtr(in.Value)
	return out
}
func ForwardingruleMetadataFilters_FromProto(mapCtx *direct.MapContext, in *pb.MetadataFilter) *krm.ForwardingruleMetadataFilters {
	if in == nil {
		return nil
	}
	out := &krm.ForwardingruleMetadataFilters{}
	out.FilterLabels = direct.Slice_FromProto(mapCtx, in.FilterLabels, ForwardingruleFilterLabels_FromProto)
	out.FilterMatchCriteria = in.GetFilterMatchCriteria()
	return out
}
func ForwardingruleMetadataFilters_ToProto(mapCtx *direct.MapContext, in *krm.ForwardingruleMetadataFilters) *pb.MetadataFilter {
	if in == nil {
		return nil
	}
	out := &pb.MetadataFilter{}
	out.FilterLabels = direct.Slice_ToProto(mapCtx, in.FilterLabels, ForwardingruleFilterLabels_ToProto)
	out.FilterMatchCriteria = direct.LazyPtr(in.FilterMatchCriteria)
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
