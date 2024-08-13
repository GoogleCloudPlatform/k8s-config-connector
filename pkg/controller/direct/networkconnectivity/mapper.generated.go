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

package networkconnectivity

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/v1alpha1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AuditConfig_FromProto(mapCtx *direct.MapContext, in *pb.AuditConfig) *krm.AuditConfig {
	if in == nil {
		return nil
	}
	out := &krm.AuditConfig{}
	out.AuditLogConfigs = direct.Slice_FromProto(mapCtx, in.AuditLogConfigs, AuditLogConfig_FromProto)
	out.Service = direct.LazyPtr(in.GetService())
	return out
}
func AuditConfig_ToProto(mapCtx *direct.MapContext, in *krm.AuditConfig) *pb.AuditConfig {
	if in == nil {
		return nil
	}
	out := &pb.AuditConfig{}
	out.AuditLogConfigs = direct.Slice_ToProto(mapCtx, in.AuditLogConfigs, AuditLogConfig_ToProto)
	out.Service = direct.ValueOf(in.Service)
	return out
}
func AuditLogConfig_FromProto(mapCtx *direct.MapContext, in *pb.AuditLogConfig) *krm.AuditLogConfig {
	if in == nil {
		return nil
	}
	out := &krm.AuditLogConfig{}
	out.ExemptedMembers = in.ExemptedMembers
	out.LogType = direct.LazyPtr(in.GetLogType())
	return out
}
func AuditLogConfig_ToProto(mapCtx *direct.MapContext, in *krm.AuditLogConfig) *pb.AuditLogConfig {
	if in == nil {
		return nil
	}
	out := &pb.AuditLogConfig{}
	out.ExemptedMembers = in.ExemptedMembers
	out.LogType = direct.ValueOf(in.LogType)
	return out
}
func AutoAccept_FromProto(mapCtx *direct.MapContext, in *pb.AutoAccept) *krm.AutoAccept {
	if in == nil {
		return nil
	}
	out := &krm.AutoAccept{}
	out.AutoAcceptProjects = in.AutoAcceptProjects
	return out
}
func AutoAccept_ToProto(mapCtx *direct.MapContext, in *krm.AutoAccept) *pb.AutoAccept {
	if in == nil {
		return nil
	}
	out := &pb.AutoAccept{}
	out.AutoAcceptProjects = in.AutoAcceptProjects
	return out
}
func Binding_FromProto(mapCtx *direct.MapContext, in *pb.Binding) *krm.Binding {
	if in == nil {
		return nil
	}
	out := &krm.Binding{}
	out.Condition = Expr_FromProto(mapCtx, in.GetCondition())
	out.Members = in.Members
	out.Role = direct.LazyPtr(in.GetRole())
	return out
}
func Binding_ToProto(mapCtx *direct.MapContext, in *krm.Binding) *pb.Binding {
	if in == nil {
		return nil
	}
	out := &pb.Binding{}
	out.Condition = Expr_ToProto(mapCtx, in.Condition)
	out.Members = in.Members
	out.Role = direct.ValueOf(in.Role)
	return out
}
func ConsumerPscConfig_FromProto(mapCtx *direct.MapContext, in *pb.ConsumerPscConfig) *krm.ConsumerPscConfig {
	if in == nil {
		return nil
	}
	out := &krm.ConsumerPscConfig{}
	out.ConsumerInstanceProject = direct.LazyPtr(in.GetConsumerInstanceProject())
	out.DisableGlobalAccess = direct.LazyPtr(in.GetDisableGlobalAccess())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.ProducerInstanceID = direct.LazyPtr(in.GetProducerInstanceId())
	out.Project = direct.LazyPtr(in.GetProject())
	out.ServiceAttachmentIpAddressMap = in.ServiceAttachmentIpAddressMap
	out.State = direct.LazyPtr(in.GetState())
	return out
}
func ConsumerPscConfig_ToProto(mapCtx *direct.MapContext, in *krm.ConsumerPscConfig) *pb.ConsumerPscConfig {
	if in == nil {
		return nil
	}
	out := &pb.ConsumerPscConfig{}
	out.ConsumerInstanceProject = direct.ValueOf(in.ConsumerInstanceProject)
	out.DisableGlobalAccess = direct.ValueOf(in.DisableGlobalAccess)
	out.Network = direct.ValueOf(in.Network)
	out.ProducerInstanceId = direct.ValueOf(in.ProducerInstanceID)
	out.Project = direct.ValueOf(in.Project)
	out.ServiceAttachmentIpAddressMap = in.ServiceAttachmentIpAddressMap
	out.State = direct.ValueOf(in.State)
	return out
}
func ConsumerPscConnection_FromProto(mapCtx *direct.MapContext, in *pb.ConsumerPscConnection) *krm.ConsumerPscConnection {
	if in == nil {
		return nil
	}
	out := &krm.ConsumerPscConnection{}
	out.Error = GoogleRpcStatus_FromProto(mapCtx, in.GetError())
	out.ErrorInfo = GoogleRpcErrorInfo_FromProto(mapCtx, in.GetErrorInfo())
	out.ErrorType = direct.LazyPtr(in.GetErrorType())
	out.ForwardingRule = direct.LazyPtr(in.GetForwardingRule())
	out.GceOperation = direct.LazyPtr(in.GetGceOperation())
	out.Ip = direct.LazyPtr(in.GetIp())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.ProducerInstanceID = direct.LazyPtr(in.GetProducerInstanceId())
	out.Project = direct.LazyPtr(in.GetProject())
	out.PscConnectionID = direct.LazyPtr(in.GetPscConnectionId())
	out.SelectedSubnetwork = direct.LazyPtr(in.GetSelectedSubnetwork())
	out.ServiceAttachmentUri = direct.LazyPtr(in.GetServiceAttachmentUri())
	out.State = direct.LazyPtr(in.GetState())
	return out
}
func ConsumerPscConnection_ToProto(mapCtx *direct.MapContext, in *krm.ConsumerPscConnection) *pb.ConsumerPscConnection {
	if in == nil {
		return nil
	}
	out := &pb.ConsumerPscConnection{}
	out.Error = GoogleRpcStatus_ToProto(mapCtx, in.Error)
	out.ErrorInfo = GoogleRpcErrorInfo_ToProto(mapCtx, in.ErrorInfo)
	out.ErrorType = direct.ValueOf(in.ErrorType)
	out.ForwardingRule = direct.ValueOf(in.ForwardingRule)
	out.GceOperation = direct.ValueOf(in.GceOperation)
	out.Ip = direct.ValueOf(in.Ip)
	out.Network = direct.ValueOf(in.Network)
	out.ProducerInstanceId = direct.ValueOf(in.ProducerInstanceID)
	out.Project = direct.ValueOf(in.Project)
	out.PscConnectionId = direct.ValueOf(in.PscConnectionID)
	out.SelectedSubnetwork = direct.ValueOf(in.SelectedSubnetwork)
	out.ServiceAttachmentUri = direct.ValueOf(in.ServiceAttachmentUri)
	out.State = direct.ValueOf(in.State)
	return out
}
func Empty_FromProto(mapCtx *direct.MapContext, in *pb.Empty) *krm.Empty {
	if in == nil {
		return nil
	}
	out := &krm.Empty{}
	return out
}
func Empty_ToProto(mapCtx *direct.MapContext, in *krm.Empty) *pb.Empty {
	if in == nil {
		return nil
	}
	out := &pb.Empty{}
	return out
}
func Expr_FromProto(mapCtx *direct.MapContext, in *pb.Expr) *krm.Expr {
	if in == nil {
		return nil
	}
	out := &krm.Expr{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Expression = direct.LazyPtr(in.GetExpression())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.Title = direct.LazyPtr(in.GetTitle())
	return out
}
func Expr_ToProto(mapCtx *direct.MapContext, in *krm.Expr) *pb.Expr {
	if in == nil {
		return nil
	}
	out := &pb.Expr{}
	out.Description = direct.ValueOf(in.Description)
	out.Expression = direct.ValueOf(in.Expression)
	out.Location = direct.ValueOf(in.Location)
	out.Title = direct.ValueOf(in.Title)
	return out
}
func Filter_FromProto(mapCtx *direct.MapContext, in *pb.Filter) *krm.Filter {
	if in == nil {
		return nil
	}
	out := &krm.Filter{}
	out.DestRange = direct.LazyPtr(in.GetDestRange())
	out.IpProtocol = direct.LazyPtr(in.GetIpProtocol())
	out.ProtocolVersion = direct.LazyPtr(in.GetProtocolVersion())
	out.SrcRange = direct.LazyPtr(in.GetSrcRange())
	return out
}
func Filter_ToProto(mapCtx *direct.MapContext, in *krm.Filter) *pb.Filter {
	if in == nil {
		return nil
	}
	out := &pb.Filter{}
	out.DestRange = direct.ValueOf(in.DestRange)
	out.IpProtocol = direct.ValueOf(in.IpProtocol)
	out.ProtocolVersion = direct.ValueOf(in.ProtocolVersion)
	out.SrcRange = direct.ValueOf(in.SrcRange)
	return out
}
func GoogleRpcErrorInfo_FromProto(mapCtx *direct.MapContext, in *pb.GoogleRpcErrorInfo) *krm.GoogleRpcErrorInfo {
	if in == nil {
		return nil
	}
	out := &krm.GoogleRpcErrorInfo{}
	out.Domain = direct.LazyPtr(in.GetDomain())
	out.Metadata = in.Metadata
	out.Reason = direct.LazyPtr(in.GetReason())
	return out
}
func GoogleRpcErrorInfo_ToProto(mapCtx *direct.MapContext, in *krm.GoogleRpcErrorInfo) *pb.GoogleRpcErrorInfo {
	if in == nil {
		return nil
	}
	out := &pb.GoogleRpcErrorInfo{}
	out.Domain = direct.ValueOf(in.Domain)
	out.Metadata = in.Metadata
	out.Reason = direct.ValueOf(in.Reason)
	return out
}
func GoogleRpcStatus_FromProto(mapCtx *direct.MapContext, in *pb.GoogleRpcStatus) *krm.GoogleRpcStatus {
	if in == nil {
		return nil
	}
	out := &krm.GoogleRpcStatus{}
	out.Code = direct.LazyPtr(in.GetCode())
	// MISSING: Details
	out.Message = direct.LazyPtr(in.GetMessage())
	return out
}
func GoogleRpcStatus_ToProto(mapCtx *direct.MapContext, in *krm.GoogleRpcStatus) *pb.GoogleRpcStatus {
	if in == nil {
		return nil
	}
	out := &pb.GoogleRpcStatus{}
	out.Code = direct.ValueOf(in.Code)
	// MISSING: Details
	out.Message = direct.ValueOf(in.Message)
	return out
}
func Group_FromProto(mapCtx *direct.MapContext, in *pb.Group) *krm.Group {
	if in == nil {
		return nil
	}
	out := &krm.Group{}
	out.AutoAccept = AutoAccept_FromProto(mapCtx, in.GetAutoAccept())
	out.CreateTime = Group_CreateTime_FromProto(mapCtx, in.GetCreateTime())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	out.Name = direct.LazyPtr(in.GetName())
	out.RouteTable = direct.LazyPtr(in.GetRouteTable())
	out.State = direct.LazyPtr(in.GetState())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.UpdateTime = Group_UpdateTime_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func Group_ToProto(mapCtx *direct.MapContext, in *krm.Group) *pb.Group {
	if in == nil {
		return nil
	}
	out := &pb.Group{}
	out.AutoAccept = AutoAccept_ToProto(mapCtx, in.AutoAccept)
	out.CreateTime = Group_CreateTime_ToProto(mapCtx, in.CreateTime)
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	out.Name = direct.ValueOf(in.Name)
	out.RouteTable = direct.ValueOf(in.RouteTable)
	out.State = direct.ValueOf(in.State)
	out.Uid = direct.ValueOf(in.Uid)
	out.UpdateTime = Group_UpdateTime_ToProto(mapCtx, in.UpdateTime)
	return out
}
func Hub_FromProto(mapCtx *direct.MapContext, in *pb.Hub) *krm.Hub {
	if in == nil {
		return nil
	}
	out := &krm.Hub{}
	out.CreateTime = Hub_CreateTime_FromProto(mapCtx, in.GetCreateTime())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.ExportPsc = direct.LazyPtr(in.GetExportPsc())
	out.Labels = in.Labels
	out.Name = direct.LazyPtr(in.GetName())
	out.PolicyMode = direct.LazyPtr(in.GetPolicyMode())
	out.PresetTopology = direct.LazyPtr(in.GetPresetTopology())
	out.RouteTables = in.RouteTables
	out.RoutingVpcs = direct.Slice_FromProto(mapCtx, in.RoutingVpcs, RoutingVPC_FromProto)
	out.SpokeSummary = SpokeSummary_FromProto(mapCtx, in.GetSpokeSummary())
	out.State = direct.LazyPtr(in.GetState())
	out.UniqueID = direct.LazyPtr(in.GetUniqueId())
	out.UpdateTime = Hub_UpdateTime_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func Hub_ToProto(mapCtx *direct.MapContext, in *krm.Hub) *pb.Hub {
	if in == nil {
		return nil
	}
	out := &pb.Hub{}
	out.CreateTime = Hub_CreateTime_ToProto(mapCtx, in.CreateTime)
	out.Description = direct.ValueOf(in.Description)
	out.ExportPsc = direct.ValueOf(in.ExportPsc)
	out.Labels = in.Labels
	out.Name = direct.ValueOf(in.Name)
	out.PolicyMode = direct.ValueOf(in.PolicyMode)
	out.PresetTopology = direct.ValueOf(in.PresetTopology)
	out.RouteTables = in.RouteTables
	out.RoutingVpcs = direct.Slice_ToProto(mapCtx, in.RoutingVpcs, RoutingVPC_ToProto)
	out.SpokeSummary = SpokeSummary_ToProto(mapCtx, in.SpokeSummary)
	out.State = direct.ValueOf(in.State)
	out.UniqueId = direct.ValueOf(in.UniqueID)
	out.UpdateTime = Hub_UpdateTime_ToProto(mapCtx, in.UpdateTime)
	return out
}
func InterconnectAttachment_FromProto(mapCtx *direct.MapContext, in *pb.InterconnectAttachment) *krm.InterconnectAttachment {
	if in == nil {
		return nil
	}
	out := &krm.InterconnectAttachment{}
	out.Region = direct.LazyPtr(in.GetRegion())
	return out
}
func InterconnectAttachment_ToProto(mapCtx *direct.MapContext, in *krm.InterconnectAttachment) *pb.InterconnectAttachment {
	if in == nil {
		return nil
	}
	out := &pb.InterconnectAttachment{}
	out.Region = direct.ValueOf(in.Region)
	return out
}
func InternalRange_FromProto(mapCtx *direct.MapContext, in *pb.InternalRange) *krm.InternalRange {
	if in == nil {
		return nil
	}
	out := &krm.InternalRange{}
	out.CreateTime = InternalRange_CreateTime_FromProto(mapCtx, in.GetCreateTime())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.IpCidrRange = direct.LazyPtr(in.GetIpCidrRange())
	out.Labels = in.Labels
	out.Name = direct.LazyPtr(in.GetName())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.Overlaps = in.Overlaps
	out.Peering = direct.LazyPtr(in.GetPeering())
	out.PrefixLength = direct.LazyPtr(in.GetPrefixLength())
	out.TargetCidrRange = in.TargetCidrRange
	out.UpdateTime = InternalRange_UpdateTime_FromProto(mapCtx, in.GetUpdateTime())
	out.Usage = direct.LazyPtr(in.GetUsage())
	out.Users = in.Users
	return out
}
func InternalRange_ToProto(mapCtx *direct.MapContext, in *krm.InternalRange) *pb.InternalRange {
	if in == nil {
		return nil
	}
	out := &pb.InternalRange{}
	out.CreateTime = InternalRange_CreateTime_ToProto(mapCtx, in.CreateTime)
	out.Description = direct.ValueOf(in.Description)
	out.IpCidrRange = direct.ValueOf(in.IpCidrRange)
	out.Labels = in.Labels
	out.Name = direct.ValueOf(in.Name)
	out.Network = direct.ValueOf(in.Network)
	out.Overlaps = in.Overlaps
	out.Peering = direct.ValueOf(in.Peering)
	out.PrefixLength = direct.ValueOf(in.PrefixLength)
	out.TargetCidrRange = in.TargetCidrRange
	out.UpdateTime = InternalRange_UpdateTime_ToProto(mapCtx, in.UpdateTime)
	out.Usage = direct.ValueOf(in.Usage)
	out.Users = in.Users
	return out
}
func LinkedInterconnectAttachments_FromProto(mapCtx *direct.MapContext, in *pb.LinkedInterconnectAttachments) *krm.LinkedInterconnectAttachments {
	if in == nil {
		return nil
	}
	out := &krm.LinkedInterconnectAttachments{}
	out.IncludeImportRanges = in.IncludeImportRanges
	out.SiteToSiteDataTransfer = direct.LazyPtr(in.GetSiteToSiteDataTransfer())
	out.Uris = in.Uris
	out.VpcNetwork = direct.LazyPtr(in.GetVpcNetwork())
	return out
}
func LinkedInterconnectAttachments_ToProto(mapCtx *direct.MapContext, in *krm.LinkedInterconnectAttachments) *pb.LinkedInterconnectAttachments {
	if in == nil {
		return nil
	}
	out := &pb.LinkedInterconnectAttachments{}
	out.IncludeImportRanges = in.IncludeImportRanges
	out.SiteToSiteDataTransfer = direct.ValueOf(in.SiteToSiteDataTransfer)
	out.Uris = in.Uris
	out.VpcNetwork = direct.ValueOf(in.VpcNetwork)
	return out
}
func LinkedRouterApplianceInstances_FromProto(mapCtx *direct.MapContext, in *pb.LinkedRouterApplianceInstances) *krm.LinkedRouterApplianceInstances {
	if in == nil {
		return nil
	}
	out := &krm.LinkedRouterApplianceInstances{}
	out.IncludeImportRanges = in.IncludeImportRanges
	out.Instances = direct.Slice_FromProto(mapCtx, in.Instances, RouterApplianceInstance_FromProto)
	out.SiteToSiteDataTransfer = direct.LazyPtr(in.GetSiteToSiteDataTransfer())
	out.VpcNetwork = direct.LazyPtr(in.GetVpcNetwork())
	return out
}
func LinkedRouterApplianceInstances_ToProto(mapCtx *direct.MapContext, in *krm.LinkedRouterApplianceInstances) *pb.LinkedRouterApplianceInstances {
	if in == nil {
		return nil
	}
	out := &pb.LinkedRouterApplianceInstances{}
	out.IncludeImportRanges = in.IncludeImportRanges
	out.Instances = direct.Slice_ToProto(mapCtx, in.Instances, RouterApplianceInstance_ToProto)
	out.SiteToSiteDataTransfer = direct.ValueOf(in.SiteToSiteDataTransfer)
	out.VpcNetwork = direct.ValueOf(in.VpcNetwork)
	return out
}
func LinkedVpcNetwork_FromProto(mapCtx *direct.MapContext, in *pb.LinkedVpcNetwork) *krm.LinkedVpcNetwork {
	if in == nil {
		return nil
	}
	out := &krm.LinkedVpcNetwork{}
	out.ExcludeExportRanges = in.ExcludeExportRanges
	out.IncludeExportRanges = in.IncludeExportRanges
	out.Uri = direct.LazyPtr(in.GetUri())
	return out
}
func LinkedVpcNetwork_ToProto(mapCtx *direct.MapContext, in *krm.LinkedVpcNetwork) *pb.LinkedVpcNetwork {
	if in == nil {
		return nil
	}
	out := &pb.LinkedVpcNetwork{}
	out.ExcludeExportRanges = in.ExcludeExportRanges
	out.IncludeExportRanges = in.IncludeExportRanges
	out.Uri = direct.ValueOf(in.Uri)
	return out
}
func LinkedVpnTunnels_FromProto(mapCtx *direct.MapContext, in *pb.LinkedVpnTunnels) *krm.LinkedVpnTunnels {
	if in == nil {
		return nil
	}
	out := &krm.LinkedVpnTunnels{}
	out.IncludeImportRanges = in.IncludeImportRanges
	out.SiteToSiteDataTransfer = direct.LazyPtr(in.GetSiteToSiteDataTransfer())
	out.Uris = in.Uris
	out.VpcNetwork = direct.LazyPtr(in.GetVpcNetwork())
	return out
}
func LinkedVpnTunnels_ToProto(mapCtx *direct.MapContext, in *krm.LinkedVpnTunnels) *pb.LinkedVpnTunnels {
	if in == nil {
		return nil
	}
	out := &pb.LinkedVpnTunnels{}
	out.IncludeImportRanges = in.IncludeImportRanges
	out.SiteToSiteDataTransfer = direct.ValueOf(in.SiteToSiteDataTransfer)
	out.Uris = in.Uris
	out.VpcNetwork = direct.ValueOf(in.VpcNetwork)
	return out
}
func Location_FromProto(mapCtx *direct.MapContext, in *pb.Location) *krm.Location {
	if in == nil {
		return nil
	}
	out := &krm.Location{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Labels = in.Labels
	out.LocationID = direct.LazyPtr(in.GetLocationId())
	// MISSING: Metadata
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func Location_ToProto(mapCtx *direct.MapContext, in *krm.Location) *pb.Location {
	if in == nil {
		return nil
	}
	out := &pb.Location{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Labels = in.Labels
	out.LocationId = direct.ValueOf(in.LocationID)
	// MISSING: Metadata
	out.Name = direct.ValueOf(in.Name)
	return out
}
func NetworkConnectivityServiceConnectionPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ServiceConnectionPolicy) *krm.NetworkConnectivityServiceConnectionPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConnectivityServiceConnectionPolicyObservedState{}
	out.CreateTime = ServiceConnectionPolicy_CreateTime_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: Description
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Infrastructure = direct.LazyPtr(in.GetInfrastructure())
	// MISSING: Labels
	// MISSING: Name
	// MISSING: Network
	// MISSING: PscConfig
	out.PscConnections = direct.Slice_FromProto(mapCtx, in.PscConnections, PscConnection_FromProto)
	// MISSING: ServiceClass
	out.UpdateTime = ServiceConnectionPolicy_UpdateTime_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func NetworkConnectivityServiceConnectionPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConnectivityServiceConnectionPolicyObservedState) *pb.ServiceConnectionPolicy {
	if in == nil {
		return nil
	}
	out := &pb.ServiceConnectionPolicy{}
	out.CreateTime = ServiceConnectionPolicy_CreateTime_ToProto(mapCtx, in.CreateTime)
	// MISSING: Description
	out.Etag = direct.ValueOf(in.Etag)
	out.Infrastructure = direct.ValueOf(in.Infrastructure)
	// MISSING: Labels
	// MISSING: Name
	// MISSING: Network
	// MISSING: PscConfig
	out.PscConnections = direct.Slice_ToProto(mapCtx, in.PscConnections, PscConnection_ToProto)
	// MISSING: ServiceClass
	out.UpdateTime = ServiceConnectionPolicy_UpdateTime_ToProto(mapCtx, in.UpdateTime)
	return out
}
func NetworkConnectivityServiceConnectionPolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConnectivityServiceConnectionPolicySpec) *pb.ServiceConnectionPolicy {
	if in == nil {
		return nil
	}
	out := &pb.ServiceConnectionPolicy{}
	// MISSING: CreateTime
	out.Description = direct.ValueOf(in.Description)
	// MISSING: Etag
	// MISSING: Infrastructure
	// MISSING: Labels
	// MISSING: Name
	out.Network = NetworkConnectivityServiceConnectionPolicySpec_Network_ToProto(mapCtx, in.Network)
	out.PscConfig = PscConfig_ToProto(mapCtx, in.PscConfig)
	// MISSING: PscConnections
	out.ServiceClass = direct.ValueOf(in.ServiceClass)
	// MISSING: UpdateTime
	return out
}
func NextHopInterconnectAttachment_FromProto(mapCtx *direct.MapContext, in *pb.NextHopInterconnectAttachment) *krm.NextHopInterconnectAttachment {
	if in == nil {
		return nil
	}
	out := &krm.NextHopInterconnectAttachment{}
	out.SiteToSiteDataTransfer = direct.LazyPtr(in.GetSiteToSiteDataTransfer())
	out.Uri = direct.LazyPtr(in.GetUri())
	out.VpcNetwork = direct.LazyPtr(in.GetVpcNetwork())
	return out
}
func NextHopInterconnectAttachment_ToProto(mapCtx *direct.MapContext, in *krm.NextHopInterconnectAttachment) *pb.NextHopInterconnectAttachment {
	if in == nil {
		return nil
	}
	out := &pb.NextHopInterconnectAttachment{}
	out.SiteToSiteDataTransfer = direct.ValueOf(in.SiteToSiteDataTransfer)
	out.Uri = direct.ValueOf(in.Uri)
	out.VpcNetwork = direct.ValueOf(in.VpcNetwork)
	return out
}
func NextHopRouterApplianceInstance_FromProto(mapCtx *direct.MapContext, in *pb.NextHopRouterApplianceInstance) *krm.NextHopRouterApplianceInstance {
	if in == nil {
		return nil
	}
	out := &krm.NextHopRouterApplianceInstance{}
	out.SiteToSiteDataTransfer = direct.LazyPtr(in.GetSiteToSiteDataTransfer())
	out.Uri = direct.LazyPtr(in.GetUri())
	out.VpcNetwork = direct.LazyPtr(in.GetVpcNetwork())
	return out
}
func NextHopRouterApplianceInstance_ToProto(mapCtx *direct.MapContext, in *krm.NextHopRouterApplianceInstance) *pb.NextHopRouterApplianceInstance {
	if in == nil {
		return nil
	}
	out := &pb.NextHopRouterApplianceInstance{}
	out.SiteToSiteDataTransfer = direct.ValueOf(in.SiteToSiteDataTransfer)
	out.Uri = direct.ValueOf(in.Uri)
	out.VpcNetwork = direct.ValueOf(in.VpcNetwork)
	return out
}
func NextHopVPNTunnel_FromProto(mapCtx *direct.MapContext, in *pb.NextHopVPNTunnel) *krm.NextHopVPNTunnel {
	if in == nil {
		return nil
	}
	out := &krm.NextHopVPNTunnel{}
	out.SiteToSiteDataTransfer = direct.LazyPtr(in.GetSiteToSiteDataTransfer())
	out.Uri = direct.LazyPtr(in.GetUri())
	out.VpcNetwork = direct.LazyPtr(in.GetVpcNetwork())
	return out
}
func NextHopVPNTunnel_ToProto(mapCtx *direct.MapContext, in *krm.NextHopVPNTunnel) *pb.NextHopVPNTunnel {
	if in == nil {
		return nil
	}
	out := &pb.NextHopVPNTunnel{}
	out.SiteToSiteDataTransfer = direct.ValueOf(in.SiteToSiteDataTransfer)
	out.Uri = direct.ValueOf(in.Uri)
	out.VpcNetwork = direct.ValueOf(in.VpcNetwork)
	return out
}
func NextHopVpcNetwork_FromProto(mapCtx *direct.MapContext, in *pb.NextHopVpcNetwork) *krm.NextHopVpcNetwork {
	if in == nil {
		return nil
	}
	out := &krm.NextHopVpcNetwork{}
	out.Uri = direct.LazyPtr(in.GetUri())
	return out
}
func NextHopVpcNetwork_ToProto(mapCtx *direct.MapContext, in *krm.NextHopVpcNetwork) *pb.NextHopVpcNetwork {
	if in == nil {
		return nil
	}
	out := &pb.NextHopVpcNetwork{}
	out.Uri = direct.ValueOf(in.Uri)
	return out
}
func Policy_FromProto(mapCtx *direct.MapContext, in *pb.Policy) *krm.Policy {
	if in == nil {
		return nil
	}
	out := &krm.Policy{}
	out.AuditConfigs = direct.Slice_FromProto(mapCtx, in.AuditConfigs, AuditConfig_FromProto)
	out.Bindings = direct.Slice_FromProto(mapCtx, in.Bindings, Binding_FromProto)
	out.Etag = in.GetEtag()
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}
func Policy_ToProto(mapCtx *direct.MapContext, in *krm.Policy) *pb.Policy {
	if in == nil {
		return nil
	}
	out := &pb.Policy{}
	out.AuditConfigs = direct.Slice_ToProto(mapCtx, in.AuditConfigs, AuditConfig_ToProto)
	out.Bindings = direct.Slice_ToProto(mapCtx, in.Bindings, Binding_ToProto)
	out.Etag = in.Etag
	out.Version = direct.ValueOf(in.Version)
	return out
}
func PolicyBasedRoute_FromProto(mapCtx *direct.MapContext, in *pb.PolicyBasedRoute) *krm.PolicyBasedRoute {
	if in == nil {
		return nil
	}
	out := &krm.PolicyBasedRoute{}
	out.CreateTime = PolicyBasedRoute_CreateTime_FromProto(mapCtx, in.GetCreateTime())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Filter = Filter_FromProto(mapCtx, in.GetFilter())
	out.InterconnectAttachment = InterconnectAttachment_FromProto(mapCtx, in.GetInterconnectAttachment())
	out.Kind = direct.LazyPtr(in.GetKind())
	out.Labels = in.Labels
	out.Name = direct.LazyPtr(in.GetName())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.NextHopIlbIp = direct.LazyPtr(in.GetNextHopIlbIp())
	out.NextHopOtherRoutes = direct.LazyPtr(in.GetNextHopOtherRoutes())
	out.Priority = direct.LazyPtr(in.GetPriority())
	out.SelfLink = direct.LazyPtr(in.GetSelfLink())
	out.UpdateTime = PolicyBasedRoute_UpdateTime_FromProto(mapCtx, in.GetUpdateTime())
	out.VirtualMachine = VirtualMachine_FromProto(mapCtx, in.GetVirtualMachine())
	out.Warnings = direct.Slice_FromProto(mapCtx, in.Warnings, Warnings_FromProto)
	return out
}
func PolicyBasedRoute_ToProto(mapCtx *direct.MapContext, in *krm.PolicyBasedRoute) *pb.PolicyBasedRoute {
	if in == nil {
		return nil
	}
	out := &pb.PolicyBasedRoute{}
	out.CreateTime = PolicyBasedRoute_CreateTime_ToProto(mapCtx, in.CreateTime)
	out.Description = direct.ValueOf(in.Description)
	out.Filter = Filter_ToProto(mapCtx, in.Filter)
	out.InterconnectAttachment = InterconnectAttachment_ToProto(mapCtx, in.InterconnectAttachment)
	out.Kind = direct.ValueOf(in.Kind)
	out.Labels = in.Labels
	out.Name = direct.ValueOf(in.Name)
	out.Network = direct.ValueOf(in.Network)
	out.NextHopIlbIp = direct.ValueOf(in.NextHopIlbIp)
	out.NextHopOtherRoutes = direct.ValueOf(in.NextHopOtherRoutes)
	out.Priority = direct.ValueOf(in.Priority)
	out.SelfLink = direct.ValueOf(in.SelfLink)
	out.UpdateTime = PolicyBasedRoute_UpdateTime_ToProto(mapCtx, in.UpdateTime)
	out.VirtualMachine = VirtualMachine_ToProto(mapCtx, in.VirtualMachine)
	out.Warnings = direct.Slice_ToProto(mapCtx, in.Warnings, Warnings_ToProto)
	return out
}
func ProducerPscConfig_FromProto(mapCtx *direct.MapContext, in *pb.ProducerPscConfig) *krm.ProducerPscConfig {
	if in == nil {
		return nil
	}
	out := &krm.ProducerPscConfig{}
	out.ServiceAttachmentUri = direct.LazyPtr(in.GetServiceAttachmentUri())
	return out
}
func ProducerPscConfig_ToProto(mapCtx *direct.MapContext, in *krm.ProducerPscConfig) *pb.ProducerPscConfig {
	if in == nil {
		return nil
	}
	out := &pb.ProducerPscConfig{}
	out.ServiceAttachmentUri = direct.ValueOf(in.ServiceAttachmentUri)
	return out
}
func PscConfig_FromProto(mapCtx *direct.MapContext, in *pb.PscConfig) *krm.PscConfig {
	if in == nil {
		return nil
	}
	out := &krm.PscConfig{}
	// MISSING: AllowedGoogleProducersResourceHierarchyLevel
	out.Limit = direct.LazyPtr(in.GetLimit())
	out.ProducerInstanceLocation = direct.LazyPtr(in.GetProducerInstanceLocation())
	out.Subnetworks = PscConfig_Subnetworks_FromProto(mapCtx, in.Subnetworks)
	return out
}
func PscConfig_ToProto(mapCtx *direct.MapContext, in *krm.PscConfig) *pb.PscConfig {
	if in == nil {
		return nil
	}
	out := &pb.PscConfig{}
	// MISSING: AllowedGoogleProducersResourceHierarchyLevel
	out.Limit = direct.ValueOf(in.Limit)
	out.ProducerInstanceLocation = direct.ValueOf(in.ProducerInstanceLocation)
	out.Subnetworks = PscConfig_Subnetworks_ToProto(mapCtx, in.Subnetworks)
	return out
}
func PscConnection_FromProto(mapCtx *direct.MapContext, in *pb.PscConnection) *krm.PscConnection {
	if in == nil {
		return nil
	}
	out := &krm.PscConnection{}
	out.ConsumerAddress = direct.LazyPtr(in.GetConsumerAddress())
	out.ConsumerForwardingRule = direct.LazyPtr(in.GetConsumerForwardingRule())
	out.ConsumerTargetProject = direct.LazyPtr(in.GetConsumerTargetProject())
	out.Error = GoogleRpcStatus_FromProto(mapCtx, in.GetError())
	out.ErrorInfo = GoogleRpcErrorInfo_FromProto(mapCtx, in.GetErrorInfo())
	out.ErrorType = direct.LazyPtr(in.GetErrorType())
	out.GceOperation = direct.LazyPtr(in.GetGceOperation())
	out.ProducerInstanceID = direct.LazyPtr(in.GetProducerInstanceId())
	out.PscConnectionID = direct.LazyPtr(in.GetPscConnectionId())
	out.SelectedSubnetwork = direct.LazyPtr(in.GetSelectedSubnetwork())
	out.State = direct.LazyPtr(in.GetState())
	return out
}
func PscConnection_ToProto(mapCtx *direct.MapContext, in *krm.PscConnection) *pb.PscConnection {
	if in == nil {
		return nil
	}
	out := &pb.PscConnection{}
	out.ConsumerAddress = direct.ValueOf(in.ConsumerAddress)
	out.ConsumerForwardingRule = direct.ValueOf(in.ConsumerForwardingRule)
	out.ConsumerTargetProject = direct.ValueOf(in.ConsumerTargetProject)
	out.Error = GoogleRpcStatus_ToProto(mapCtx, in.Error)
	out.ErrorInfo = GoogleRpcErrorInfo_ToProto(mapCtx, in.ErrorInfo)
	out.ErrorType = direct.ValueOf(in.ErrorType)
	out.GceOperation = direct.ValueOf(in.GceOperation)
	out.ProducerInstanceId = direct.ValueOf(in.ProducerInstanceID)
	out.PscConnectionId = direct.ValueOf(in.PscConnectionID)
	out.SelectedSubnetwork = direct.ValueOf(in.SelectedSubnetwork)
	out.State = direct.ValueOf(in.State)
	return out
}
func RegionalEndpoint_FromProto(mapCtx *direct.MapContext, in *pb.RegionalEndpoint) *krm.RegionalEndpoint {
	if in == nil {
		return nil
	}
	out := &krm.RegionalEndpoint{}
	out.AccessType = direct.LazyPtr(in.GetAccessType())
	out.Address = direct.LazyPtr(in.GetAddress())
	out.CreateTime = RegionalEndpoint_CreateTime_FromProto(mapCtx, in.GetCreateTime())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.IpAddress = direct.LazyPtr(in.GetIpAddress())
	out.Labels = in.Labels
	out.Name = direct.LazyPtr(in.GetName())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.PscForwardingRule = direct.LazyPtr(in.GetPscForwardingRule())
	out.Subnetwork = direct.LazyPtr(in.GetSubnetwork())
	out.TargetGoogleApi = direct.LazyPtr(in.GetTargetGoogleApi())
	out.UpdateTime = RegionalEndpoint_UpdateTime_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func RegionalEndpoint_ToProto(mapCtx *direct.MapContext, in *krm.RegionalEndpoint) *pb.RegionalEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.RegionalEndpoint{}
	out.AccessType = direct.ValueOf(in.AccessType)
	out.Address = direct.ValueOf(in.Address)
	out.CreateTime = RegionalEndpoint_CreateTime_ToProto(mapCtx, in.CreateTime)
	out.Description = direct.ValueOf(in.Description)
	out.IpAddress = direct.ValueOf(in.IpAddress)
	out.Labels = in.Labels
	out.Name = direct.ValueOf(in.Name)
	out.Network = direct.ValueOf(in.Network)
	out.PscForwardingRule = direct.ValueOf(in.PscForwardingRule)
	out.Subnetwork = direct.ValueOf(in.Subnetwork)
	out.TargetGoogleApi = direct.ValueOf(in.TargetGoogleApi)
	out.UpdateTime = RegionalEndpoint_UpdateTime_ToProto(mapCtx, in.UpdateTime)
	return out
}
func Route_FromProto(mapCtx *direct.MapContext, in *pb.Route) *krm.Route {
	if in == nil {
		return nil
	}
	out := &krm.Route{}
	out.CreateTime = Route_CreateTime_FromProto(mapCtx, in.GetCreateTime())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.IpCidrRange = direct.LazyPtr(in.GetIpCidrRange())
	out.Labels = in.Labels
	out.Location = direct.LazyPtr(in.GetLocation())
	out.Name = direct.LazyPtr(in.GetName())
	out.NextHopInterconnectAttachment = NextHopInterconnectAttachment_FromProto(mapCtx, in.GetNextHopInterconnectAttachment())
	out.NextHopRouterApplianceInstance = NextHopRouterApplianceInstance_FromProto(mapCtx, in.GetNextHopRouterApplianceInstance())
	out.NextHopVpcNetwork = NextHopVpcNetwork_FromProto(mapCtx, in.GetNextHopVpcNetwork())
	out.NextHopVpnTunnel = NextHopVPNTunnel_FromProto(mapCtx, in.GetNextHopVpnTunnel())
	out.Priority = direct.LazyPtr(in.GetPriority())
	out.Spoke = direct.LazyPtr(in.GetSpoke())
	out.State = direct.LazyPtr(in.GetState())
	out.Type = direct.LazyPtr(in.GetType())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.UpdateTime = Route_UpdateTime_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func Route_ToProto(mapCtx *direct.MapContext, in *krm.Route) *pb.Route {
	if in == nil {
		return nil
	}
	out := &pb.Route{}
	out.CreateTime = Route_CreateTime_ToProto(mapCtx, in.CreateTime)
	out.Description = direct.ValueOf(in.Description)
	out.IpCidrRange = direct.ValueOf(in.IpCidrRange)
	out.Labels = in.Labels
	out.Location = direct.ValueOf(in.Location)
	out.Name = direct.ValueOf(in.Name)
	out.NextHopInterconnectAttachment = NextHopInterconnectAttachment_ToProto(mapCtx, in.NextHopInterconnectAttachment)
	out.NextHopRouterApplianceInstance = NextHopRouterApplianceInstance_ToProto(mapCtx, in.NextHopRouterApplianceInstance)
	out.NextHopVpcNetwork = NextHopVpcNetwork_ToProto(mapCtx, in.NextHopVpcNetwork)
	out.NextHopVpnTunnel = NextHopVPNTunnel_ToProto(mapCtx, in.NextHopVpnTunnel)
	out.Priority = direct.ValueOf(in.Priority)
	out.Spoke = direct.ValueOf(in.Spoke)
	out.State = direct.ValueOf(in.State)
	out.Type = direct.ValueOf(in.Type)
	out.Uid = direct.ValueOf(in.Uid)
	out.UpdateTime = Route_UpdateTime_ToProto(mapCtx, in.UpdateTime)
	return out
}
func RouteTable_FromProto(mapCtx *direct.MapContext, in *pb.RouteTable) *krm.RouteTable {
	if in == nil {
		return nil
	}
	out := &krm.RouteTable{}
	out.CreateTime = RouteTable_CreateTime_FromProto(mapCtx, in.GetCreateTime())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	out.Name = direct.LazyPtr(in.GetName())
	out.State = direct.LazyPtr(in.GetState())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.UpdateTime = RouteTable_UpdateTime_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func RouteTable_ToProto(mapCtx *direct.MapContext, in *krm.RouteTable) *pb.RouteTable {
	if in == nil {
		return nil
	}
	out := &pb.RouteTable{}
	out.CreateTime = RouteTable_CreateTime_ToProto(mapCtx, in.CreateTime)
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	out.Name = direct.ValueOf(in.Name)
	out.State = direct.ValueOf(in.State)
	out.Uid = direct.ValueOf(in.Uid)
	out.UpdateTime = RouteTable_UpdateTime_ToProto(mapCtx, in.UpdateTime)
	return out
}
func RouterApplianceInstance_FromProto(mapCtx *direct.MapContext, in *pb.RouterApplianceInstance) *krm.RouterApplianceInstance {
	if in == nil {
		return nil
	}
	out := &krm.RouterApplianceInstance{}
	out.IpAddress = direct.LazyPtr(in.GetIpAddress())
	out.VirtualMachine = direct.LazyPtr(in.GetVirtualMachine())
	return out
}
func RouterApplianceInstance_ToProto(mapCtx *direct.MapContext, in *krm.RouterApplianceInstance) *pb.RouterApplianceInstance {
	if in == nil {
		return nil
	}
	out := &pb.RouterApplianceInstance{}
	out.IpAddress = direct.ValueOf(in.IpAddress)
	out.VirtualMachine = direct.ValueOf(in.VirtualMachine)
	return out
}
func RoutingVPC_FromProto(mapCtx *direct.MapContext, in *pb.RoutingVPC) *krm.RoutingVPC {
	if in == nil {
		return nil
	}
	out := &krm.RoutingVPC{}
	out.RequiredForNewSiteToSiteDataTransferSpokes = direct.LazyPtr(in.GetRequiredForNewSiteToSiteDataTransferSpokes())
	out.Uri = direct.LazyPtr(in.GetUri())
	return out
}
func RoutingVPC_ToProto(mapCtx *direct.MapContext, in *krm.RoutingVPC) *pb.RoutingVPC {
	if in == nil {
		return nil
	}
	out := &pb.RoutingVPC{}
	out.RequiredForNewSiteToSiteDataTransferSpokes = direct.ValueOf(in.RequiredForNewSiteToSiteDataTransferSpokes)
	out.Uri = direct.ValueOf(in.Uri)
	return out
}
func ServiceClass_FromProto(mapCtx *direct.MapContext, in *pb.ServiceClass) *krm.ServiceClass {
	if in == nil {
		return nil
	}
	out := &krm.ServiceClass{}
	out.CreateTime = ServiceClass_CreateTime_FromProto(mapCtx, in.GetCreateTime())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Labels = in.Labels
	out.Name = direct.LazyPtr(in.GetName())
	out.ServiceClass = direct.LazyPtr(in.GetServiceClass())
	out.UpdateTime = ServiceClass_UpdateTime_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func ServiceClass_ToProto(mapCtx *direct.MapContext, in *krm.ServiceClass) *pb.ServiceClass {
	if in == nil {
		return nil
	}
	out := &pb.ServiceClass{}
	out.CreateTime = ServiceClass_CreateTime_ToProto(mapCtx, in.CreateTime)
	out.Description = direct.ValueOf(in.Description)
	out.Etag = direct.ValueOf(in.Etag)
	out.Labels = in.Labels
	out.Name = direct.ValueOf(in.Name)
	out.ServiceClass = direct.ValueOf(in.ServiceClass)
	out.UpdateTime = ServiceClass_UpdateTime_ToProto(mapCtx, in.UpdateTime)
	return out
}
func ServiceConnectionMap_FromProto(mapCtx *direct.MapContext, in *pb.ServiceConnectionMap) *krm.ServiceConnectionMap {
	if in == nil {
		return nil
	}
	out := &krm.ServiceConnectionMap{}
	out.ConsumerPscConfigs = direct.Slice_FromProto(mapCtx, in.ConsumerPscConfigs, ConsumerPscConfig_FromProto)
	out.ConsumerPscConnections = direct.Slice_FromProto(mapCtx, in.ConsumerPscConnections, ConsumerPscConnection_FromProto)
	out.CreateTime = ServiceConnectionMap_CreateTime_FromProto(mapCtx, in.GetCreateTime())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Infrastructure = direct.LazyPtr(in.GetInfrastructure())
	out.Labels = in.Labels
	out.Name = direct.LazyPtr(in.GetName())
	out.ProducerPscConfigs = direct.Slice_FromProto(mapCtx, in.ProducerPscConfigs, ProducerPscConfig_FromProto)
	out.ServiceClass = direct.LazyPtr(in.GetServiceClass())
	out.ServiceClassUri = direct.LazyPtr(in.GetServiceClassUri())
	out.Token = direct.LazyPtr(in.GetToken())
	out.UpdateTime = ServiceConnectionMap_UpdateTime_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func ServiceConnectionMap_ToProto(mapCtx *direct.MapContext, in *krm.ServiceConnectionMap) *pb.ServiceConnectionMap {
	if in == nil {
		return nil
	}
	out := &pb.ServiceConnectionMap{}
	out.ConsumerPscConfigs = direct.Slice_ToProto(mapCtx, in.ConsumerPscConfigs, ConsumerPscConfig_ToProto)
	out.ConsumerPscConnections = direct.Slice_ToProto(mapCtx, in.ConsumerPscConnections, ConsumerPscConnection_ToProto)
	out.CreateTime = ServiceConnectionMap_CreateTime_ToProto(mapCtx, in.CreateTime)
	out.Description = direct.ValueOf(in.Description)
	out.Etag = direct.ValueOf(in.Etag)
	out.Infrastructure = direct.ValueOf(in.Infrastructure)
	out.Labels = in.Labels
	out.Name = direct.ValueOf(in.Name)
	out.ProducerPscConfigs = direct.Slice_ToProto(mapCtx, in.ProducerPscConfigs, ProducerPscConfig_ToProto)
	out.ServiceClass = direct.ValueOf(in.ServiceClass)
	out.ServiceClassUri = direct.ValueOf(in.ServiceClassUri)
	out.Token = direct.ValueOf(in.Token)
	out.UpdateTime = ServiceConnectionMap_UpdateTime_ToProto(mapCtx, in.UpdateTime)
	return out
}
func ServiceConnectionPolicy_FromProto(mapCtx *direct.MapContext, in *pb.ServiceConnectionPolicy) *krm.ServiceConnectionPolicy {
	if in == nil {
		return nil
	}
	out := &krm.ServiceConnectionPolicy{}
	out.CreateTime = ServiceConnectionPolicy_CreateTime_FromProto(mapCtx, in.GetCreateTime())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Infrastructure = direct.LazyPtr(in.GetInfrastructure())
	out.Labels = in.Labels
	out.Name = direct.LazyPtr(in.GetName())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.PscConfig = PscConfig_FromProto(mapCtx, in.GetPscConfig())
	out.PscConnections = direct.Slice_FromProto(mapCtx, in.PscConnections, PscConnection_FromProto)
	out.ServiceClass = direct.LazyPtr(in.GetServiceClass())
	out.UpdateTime = ServiceConnectionPolicy_UpdateTime_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func ServiceConnectionPolicy_ToProto(mapCtx *direct.MapContext, in *krm.ServiceConnectionPolicy) *pb.ServiceConnectionPolicy {
	if in == nil {
		return nil
	}
	out := &pb.ServiceConnectionPolicy{}
	out.CreateTime = ServiceConnectionPolicy_CreateTime_ToProto(mapCtx, in.CreateTime)
	out.Description = direct.ValueOf(in.Description)
	out.Etag = direct.ValueOf(in.Etag)
	out.Infrastructure = direct.ValueOf(in.Infrastructure)
	out.Labels = in.Labels
	out.Name = direct.ValueOf(in.Name)
	out.Network = direct.ValueOf(in.Network)
	out.PscConfig = PscConfig_ToProto(mapCtx, in.PscConfig)
	out.PscConnections = direct.Slice_ToProto(mapCtx, in.PscConnections, PscConnection_ToProto)
	out.ServiceClass = direct.ValueOf(in.ServiceClass)
	out.UpdateTime = ServiceConnectionPolicy_UpdateTime_ToProto(mapCtx, in.UpdateTime)
	return out
}
func ServiceConnectionToken_FromProto(mapCtx *direct.MapContext, in *pb.ServiceConnectionToken) *krm.ServiceConnectionToken {
	if in == nil {
		return nil
	}
	out := &krm.ServiceConnectionToken{}
	out.CreateTime = ServiceConnectionToken_CreateTime_FromProto(mapCtx, in.GetCreateTime())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.ExpireTime = ServiceConnectionToken_ExpireTime_FromProto(mapCtx, in.GetExpireTime())
	out.Labels = in.Labels
	out.Name = direct.LazyPtr(in.GetName())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.Token = direct.LazyPtr(in.GetToken())
	out.UpdateTime = ServiceConnectionToken_UpdateTime_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func ServiceConnectionToken_ToProto(mapCtx *direct.MapContext, in *krm.ServiceConnectionToken) *pb.ServiceConnectionToken {
	if in == nil {
		return nil
	}
	out := &pb.ServiceConnectionToken{}
	out.CreateTime = ServiceConnectionToken_CreateTime_ToProto(mapCtx, in.CreateTime)
	out.Description = direct.ValueOf(in.Description)
	out.Etag = direct.ValueOf(in.Etag)
	out.ExpireTime = ServiceConnectionToken_ExpireTime_ToProto(mapCtx, in.ExpireTime)
	out.Labels = in.Labels
	out.Name = direct.ValueOf(in.Name)
	out.Network = direct.ValueOf(in.Network)
	out.Token = direct.ValueOf(in.Token)
	out.UpdateTime = ServiceConnectionToken_UpdateTime_ToProto(mapCtx, in.UpdateTime)
	return out
}
func Spoke_FromProto(mapCtx *direct.MapContext, in *pb.Spoke) *krm.Spoke {
	if in == nil {
		return nil
	}
	out := &krm.Spoke{}
	out.CreateTime = Spoke_CreateTime_FromProto(mapCtx, in.GetCreateTime())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Group = direct.LazyPtr(in.GetGroup())
	out.Hub = direct.LazyPtr(in.GetHub())
	out.Labels = in.Labels
	out.LinkedInterconnectAttachments = LinkedInterconnectAttachments_FromProto(mapCtx, in.GetLinkedInterconnectAttachments())
	out.LinkedRouterApplianceInstances = LinkedRouterApplianceInstances_FromProto(mapCtx, in.GetLinkedRouterApplianceInstances())
	out.LinkedVpcNetwork = LinkedVpcNetwork_FromProto(mapCtx, in.GetLinkedVpcNetwork())
	out.LinkedVpnTunnels = LinkedVpnTunnels_FromProto(mapCtx, in.GetLinkedVpnTunnels())
	out.Name = direct.LazyPtr(in.GetName())
	out.Reasons = direct.Slice_FromProto(mapCtx, in.Reasons, StateReason_FromProto)
	out.SpokeType = direct.LazyPtr(in.GetSpokeType())
	out.State = direct.LazyPtr(in.GetState())
	out.UniqueID = direct.LazyPtr(in.GetUniqueId())
	out.UpdateTime = Spoke_UpdateTime_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func Spoke_ToProto(mapCtx *direct.MapContext, in *krm.Spoke) *pb.Spoke {
	if in == nil {
		return nil
	}
	out := &pb.Spoke{}
	out.CreateTime = Spoke_CreateTime_ToProto(mapCtx, in.CreateTime)
	out.Description = direct.ValueOf(in.Description)
	out.Group = direct.ValueOf(in.Group)
	out.Hub = direct.ValueOf(in.Hub)
	out.Labels = in.Labels
	out.LinkedInterconnectAttachments = LinkedInterconnectAttachments_ToProto(mapCtx, in.LinkedInterconnectAttachments)
	out.LinkedRouterApplianceInstances = LinkedRouterApplianceInstances_ToProto(mapCtx, in.LinkedRouterApplianceInstances)
	out.LinkedVpcNetwork = LinkedVpcNetwork_ToProto(mapCtx, in.LinkedVpcNetwork)
	out.LinkedVpnTunnels = LinkedVpnTunnels_ToProto(mapCtx, in.LinkedVpnTunnels)
	out.Name = direct.ValueOf(in.Name)
	out.Reasons = direct.Slice_ToProto(mapCtx, in.Reasons, StateReason_ToProto)
	out.SpokeType = direct.ValueOf(in.SpokeType)
	out.State = direct.ValueOf(in.State)
	out.UniqueId = direct.ValueOf(in.UniqueID)
	out.UpdateTime = Spoke_UpdateTime_ToProto(mapCtx, in.UpdateTime)
	return out
}
func SpokeStateCount_FromProto(mapCtx *direct.MapContext, in *pb.SpokeStateCount) *krm.SpokeStateCount {
	if in == nil {
		return nil
	}
	out := &krm.SpokeStateCount{}
	out.Count = direct.LazyPtr(in.GetCount())
	out.State = direct.LazyPtr(in.GetState())
	return out
}
func SpokeStateCount_ToProto(mapCtx *direct.MapContext, in *krm.SpokeStateCount) *pb.SpokeStateCount {
	if in == nil {
		return nil
	}
	out := &pb.SpokeStateCount{}
	out.Count = direct.ValueOf(in.Count)
	out.State = direct.ValueOf(in.State)
	return out
}
func SpokeStateReasonCount_FromProto(mapCtx *direct.MapContext, in *pb.SpokeStateReasonCount) *krm.SpokeStateReasonCount {
	if in == nil {
		return nil
	}
	out := &krm.SpokeStateReasonCount{}
	out.Count = direct.LazyPtr(in.GetCount())
	out.StateReasonCode = direct.LazyPtr(in.GetStateReasonCode())
	return out
}
func SpokeStateReasonCount_ToProto(mapCtx *direct.MapContext, in *krm.SpokeStateReasonCount) *pb.SpokeStateReasonCount {
	if in == nil {
		return nil
	}
	out := &pb.SpokeStateReasonCount{}
	out.Count = direct.ValueOf(in.Count)
	out.StateReasonCode = direct.ValueOf(in.StateReasonCode)
	return out
}
func SpokeSummary_FromProto(mapCtx *direct.MapContext, in *pb.SpokeSummary) *krm.SpokeSummary {
	if in == nil {
		return nil
	}
	out := &krm.SpokeSummary{}
	out.SpokeStateCounts = direct.Slice_FromProto(mapCtx, in.SpokeStateCounts, SpokeStateCount_FromProto)
	out.SpokeStateReasonCounts = direct.Slice_FromProto(mapCtx, in.SpokeStateReasonCounts, SpokeStateReasonCount_FromProto)
	out.SpokeTypeCounts = direct.Slice_FromProto(mapCtx, in.SpokeTypeCounts, SpokeTypeCount_FromProto)
	return out
}
func SpokeSummary_ToProto(mapCtx *direct.MapContext, in *krm.SpokeSummary) *pb.SpokeSummary {
	if in == nil {
		return nil
	}
	out := &pb.SpokeSummary{}
	out.SpokeStateCounts = direct.Slice_ToProto(mapCtx, in.SpokeStateCounts, SpokeStateCount_ToProto)
	out.SpokeStateReasonCounts = direct.Slice_ToProto(mapCtx, in.SpokeStateReasonCounts, SpokeStateReasonCount_ToProto)
	out.SpokeTypeCounts = direct.Slice_ToProto(mapCtx, in.SpokeTypeCounts, SpokeTypeCount_ToProto)
	return out
}
func SpokeTypeCount_FromProto(mapCtx *direct.MapContext, in *pb.SpokeTypeCount) *krm.SpokeTypeCount {
	if in == nil {
		return nil
	}
	out := &krm.SpokeTypeCount{}
	out.Count = direct.LazyPtr(in.GetCount())
	out.SpokeType = direct.LazyPtr(in.GetSpokeType())
	return out
}
func SpokeTypeCount_ToProto(mapCtx *direct.MapContext, in *krm.SpokeTypeCount) *pb.SpokeTypeCount {
	if in == nil {
		return nil
	}
	out := &pb.SpokeTypeCount{}
	out.Count = direct.ValueOf(in.Count)
	out.SpokeType = direct.ValueOf(in.SpokeType)
	return out
}
func StateReason_FromProto(mapCtx *direct.MapContext, in *pb.StateReason) *krm.StateReason {
	if in == nil {
		return nil
	}
	out := &krm.StateReason{}
	out.Code = direct.LazyPtr(in.GetCode())
	out.Message = direct.LazyPtr(in.GetMessage())
	out.UserDetails = direct.LazyPtr(in.GetUserDetails())
	return out
}
func StateReason_ToProto(mapCtx *direct.MapContext, in *krm.StateReason) *pb.StateReason {
	if in == nil {
		return nil
	}
	out := &pb.StateReason{}
	out.Code = direct.ValueOf(in.Code)
	out.Message = direct.ValueOf(in.Message)
	out.UserDetails = direct.ValueOf(in.UserDetails)
	return out
}
func VirtualMachine_FromProto(mapCtx *direct.MapContext, in *pb.VirtualMachine) *krm.VirtualMachine {
	if in == nil {
		return nil
	}
	out := &krm.VirtualMachine{}
	out.Tags = in.Tags
	return out
}
func VirtualMachine_ToProto(mapCtx *direct.MapContext, in *krm.VirtualMachine) *pb.VirtualMachine {
	if in == nil {
		return nil
	}
	out := &pb.VirtualMachine{}
	out.Tags = in.Tags
	return out
}
func Warnings_FromProto(mapCtx *direct.MapContext, in *pb.Warnings) *krm.Warnings {
	if in == nil {
		return nil
	}
	out := &krm.Warnings{}
	out.Code = direct.LazyPtr(in.GetCode())
	out.Data = in.Data
	out.WarningMessage = direct.LazyPtr(in.GetWarningMessage())
	return out
}
func Warnings_ToProto(mapCtx *direct.MapContext, in *krm.Warnings) *pb.Warnings {
	if in == nil {
		return nil
	}
	out := &pb.Warnings{}
	out.Code = direct.ValueOf(in.Code)
	out.Data = in.Data
	out.WarningMessage = direct.ValueOf(in.WarningMessage)
	return out
}
