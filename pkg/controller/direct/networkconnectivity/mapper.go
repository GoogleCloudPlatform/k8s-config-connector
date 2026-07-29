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
	"time"

	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"

	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"

	pb "cloud.google.com/go/networkconnectivity/apiv1/networkconnectivitypb"
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/v1alpha1"
	pbgen "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NetworkConnectivityServiceConnectionPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.ServiceConnectionPolicy) *krm.NetworkConnectivityServiceConnectionPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConnectivityServiceConnectionPolicySpec{}
	// MISSING: CreateTime
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: Etag
	// MISSING: Infrastructure
	// MISSING: Labels
	// MISSING: Name
	if in.Network != "" {
		out.Network = &computerefs.ComputeNetworkRef{External: in.Network}
	}
	out.PSCConfig = PSCConfig_FromProto(mapCtx, in.GetPscConfig())
	// MISSING: PscConnections
	out.ServiceClass = direct.LazyPtr(in.GetServiceClass())
	// MISSING: UpdateTime
	return out
}
func NetworkConnectivityServiceConnectionPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ServiceConnectionPolicy) *krm.NetworkConnectivityServiceConnectionPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConnectivityServiceConnectionPolicyObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Infrastructure = direct.Enum_FromProto(mapCtx, in.GetInfrastructure())
	out.PSCConnections = direct.Slice_FromProto(mapCtx, in.PscConnections, PSCConnection_FromProto)
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func NetworkConnectivityServiceConnectionPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConnectivityServiceConnectionPolicyObservedState) *pb.ServiceConnectionPolicy {
	if in == nil {
		return nil
	}
	out := &pb.ServiceConnectionPolicy{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.Etag = in.Etag
	out.Infrastructure = direct.Enum_ToProto[pb.Infrastructure](mapCtx, in.Infrastructure)
	out.PscConnections = direct.Slice_ToProto(mapCtx, in.PSCConnections, PSCConnection_ToProto)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}

func NetworkConnectivityServiceConnectionPolicySpec_Network_ToProto(mapCtx *direct.MapContext, in *computerefs.ComputeNetworkRef) string {
	if in == nil {
		return ""
	}
	return in.External
}
func PSCConfig_FromProto(mapCtx *direct.MapContext, in *pb.ServiceConnectionPolicy_PscConfig) *krm.PSCConfig {
	if in == nil {
		return nil
	}
	out := &krm.PSCConfig{}
	out.Subnetworks = PSCConfig_Subnetworks_FromProto(mapCtx, in.Subnetworks)
	out.Limit = in.Limit
	out.ProducerInstanceLocation = direct.Enum_FromProto(mapCtx, in.GetProducerInstanceLocation())
	return out
}

func PSCConfig_ToProto(mapCtx *direct.MapContext, in *krm.PSCConfig) *pb.ServiceConnectionPolicy_PscConfig {
	if in == nil {
		return nil
	}
	out := &pb.ServiceConnectionPolicy_PscConfig{}
	out.Subnetworks = PSCConfig_Subnetworks_ToProto(mapCtx, in.Subnetworks)
	out.Limit = in.Limit
	out.ProducerInstanceLocation = direct.Enum_ToProto[pb.ServiceConnectionPolicy_PscConfig_ProducerInstanceLocation](mapCtx, in.ProducerInstanceLocation)
	return out
}

func PSCConnection_FromProto(mapCtx *direct.MapContext, in *pb.ServiceConnectionPolicy_PscConnection) *krm.PSCConnection {
	if in == nil {
		return nil
	}
	out := &krm.PSCConnection{}
	out.ConsumerAddress = direct.LazyPtr(in.GetConsumerAddress())
	out.ConsumerForwardingRule = direct.LazyPtr(in.GetConsumerForwardingRule())
	out.ConsumerTargetProject = direct.LazyPtr(in.GetConsumerTargetProject())
	out.Error = direct.Status_FromProto(mapCtx, in.GetError())
	out.ErrorType = direct.Enum_FromProto(mapCtx, in.GetErrorType())
	out.GCEOperation = direct.LazyPtr(in.GetGceOperation())
	out.ProducerInstanceID = direct.LazyPtr(in.GetProducerInstanceId())
	out.PSCConnectionID = direct.LazyPtr(in.GetPscConnectionId())
	out.SelectedSubnetwork = direct.LazyPtr(in.GetSelectedSubnetwork())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}

func PSCConnection_ToProto(mapCtx *direct.MapContext, in *krm.PSCConnection) *pb.ServiceConnectionPolicy_PscConnection {
	if in == nil {
		return nil
	}
	out := &pb.ServiceConnectionPolicy_PscConnection{}
	out.ConsumerAddress = direct.ValueOf(in.ConsumerAddress)
	out.ConsumerForwardingRule = direct.ValueOf(in.ConsumerForwardingRule)
	out.ConsumerTargetProject = direct.ValueOf(in.ConsumerTargetProject)
	out.Error = direct.Status_ToProto(mapCtx, in.Error)
	out.ErrorType = direct.Enum_ToProto[pb.ConnectionErrorType](mapCtx, in.ErrorType)
	out.GceOperation = direct.ValueOf(in.GCEOperation)
	out.ProducerInstanceId = direct.ValueOf(in.ProducerInstanceID)
	out.PscConnectionId = direct.ValueOf(in.PSCConnectionID)
	out.SelectedSubnetwork = direct.ValueOf(in.SelectedSubnetwork)
	out.State = direct.Enum_ToProto[pb.ServiceConnectionPolicy_State](mapCtx, in.State)
	return out
}

func AllocationOptions_FromProto(mapCtx *direct.MapContext, in *pb.InternalRange_AllocationOptions) *krm.AllocationOptions {
	if in == nil {
		return nil
	}
	out := &krm.AllocationOptions{}
	out.AllocationStrategy = direct.Enum_FromProto(mapCtx, in.GetAllocationStrategy())
	out.FirstAvailableRangesLookupSize = direct.LazyPtr(in.GetFirstAvailableRangesLookupSize())
	return out
}

func AllocationOptions_ToProto(mapCtx *direct.MapContext, in *krm.AllocationOptions) *pb.InternalRange_AllocationOptions {
	if in == nil {
		return nil
	}
	out := &pb.InternalRange_AllocationOptions{}
	out.AllocationStrategy = direct.Enum_ToProto[pb.InternalRange_AllocationStrategy](mapCtx, in.AllocationStrategy)
	out.FirstAvailableRangesLookupSize = direct.ValueOf(in.FirstAvailableRangesLookupSize)
	return out
}

func Migration_FromProto(mapCtx *direct.MapContext, in *pb.InternalRange_Migration) *krm.Migration {
	if in == nil {
		return nil
	}
	out := &krm.Migration{}
	out.Source = direct.LazyPtr(in.GetSource())
	out.Target = direct.LazyPtr(in.GetTarget())
	return out
}

func Migration_ToProto(mapCtx *direct.MapContext, in *krm.Migration) *pb.InternalRange_Migration {
	if in == nil {
		return nil
	}
	out := &pb.InternalRange_Migration{}
	out.Source = direct.ValueOf(in.Source)
	out.Target = direct.ValueOf(in.Target)
	return out
}

func NetworkConnectivityInternalRangeSpec_FromProto(mapCtx *direct.MapContext, in *pb.InternalRange) *krm.NetworkConnectivityInternalRangeSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConnectivityInternalRangeSpec{}
	out.AllocationOptions = AllocationOptions_FromProto(mapCtx, in.GetAllocationOptions())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.IPCIDRRange = direct.LazyPtr(in.GetIpCidrRange())
	out.Labels = in.Labels
	out.Migration = Migration_FromProto(mapCtx, in.GetMigration())
	if in.GetNetwork() != "" {
		out.NetworkRef = &computerefs.ComputeNetworkRef{External: in.GetNetwork()}
	}
	out.Overlaps = direct.EnumSlice_FromProto(mapCtx, in.Overlaps)
	out.Peering = direct.Enum_FromProto(mapCtx, in.GetPeering())
	out.PrefixLength = direct.LazyPtr(in.GetPrefixLength())
	out.TargetCIDRRange = in.TargetCidrRange
	out.Usage = direct.Enum_FromProto(mapCtx, in.GetUsage())
	return out
}

func NetworkConnectivityInternalRangeSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConnectivityInternalRangeSpec) *pb.InternalRange {
	if in == nil {
		return nil
	}
	out := &pb.InternalRange{}
	out.AllocationOptions = AllocationOptions_ToProto(mapCtx, in.AllocationOptions)
	out.Description = direct.ValueOf(in.Description)
	out.IpCidrRange = direct.ValueOf(in.IPCIDRRange)
	out.Labels = in.Labels
	out.Migration = Migration_ToProto(mapCtx, in.Migration)
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	out.Overlaps = direct.EnumSlice_ToProto[pb.InternalRange_Overlap](mapCtx, in.Overlaps)
	out.Peering = direct.Enum_ToProto[pb.InternalRange_Peering](mapCtx, in.Peering)
	out.PrefixLength = direct.ValueOf(in.PrefixLength)
	out.TargetCidrRange = in.TargetCIDRRange
	out.Usage = direct.Enum_ToProto[pb.InternalRange_Usage](mapCtx, in.Usage)
	return out
}

func PSCConfig_Subnetworks_FromProto(mapCtx *direct.MapContext, in []string) []computev1beta1.ComputeSubnetworkRef {
	if in == nil {
		return nil
	}
	var out []computev1beta1.ComputeSubnetworkRef
	for _, s := range in {
		out = append(out, computev1beta1.ComputeSubnetworkRef{External: s})
	}
	return out
}
func PSCConfig_Subnetworks_ToProto(mapCtx *direct.MapContext, in []computev1beta1.ComputeSubnetworkRef) []string {
	if in == nil {
		return nil
	}
	var out []string
	for _, ref := range in {
		out = append(out, ref.External)
	}
	return out
}

func Group_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func Group_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func Group_UpdateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func Group_UpdateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func InternalRange_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func InternalRange_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func InternalRange_UpdateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func InternalRange_UpdateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func Hub_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func Hub_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func Hub_UpdateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func Hub_UpdateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func ServiceConnectionPolicy_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	return Timestamp_FromProto(mapCtx, in)
}

func ServiceConnectionPolicy_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	return Timestamp_ToProto(mapCtx, in)
}

func ServiceConnectionPolicy_UpdateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	return Timestamp_FromProto(mapCtx, in)
}
func ServiceConnectionPolicy_UpdateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	return Timestamp_ToProto(mapCtx, in)
}

func OperationMetadata_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func OperationMetadata_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func OperationMetadata_EndTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func OperationMetadata_EndTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func RegionalEndpoint_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func RegionalEndpoint_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func RegionalEndpoint_UpdateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func RegionalEndpoint_UpdateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func PolicyBasedRoute_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func PolicyBasedRoute_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func PolicyBasedRoute_UpdateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func PolicyBasedRoute_UpdateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func Route_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func Route_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func Route_UpdateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func Route_UpdateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func RouteTable_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func RouteTable_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func RouteTable_UpdateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func RouteTable_UpdateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func ServiceClass_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func ServiceClass_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func ServiceClass_UpdateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func ServiceClass_UpdateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func ServiceConnectionMap_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func ServiceConnectionMap_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func ServiceConnectionMap_UpdateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func ServiceConnectionMap_UpdateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func ServiceConnectionToken_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func ServiceConnectionToken_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func ServiceConnectionToken_UpdateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func ServiceConnectionToken_UpdateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func Spoke_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func Spoke_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func Spoke_UpdateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func Spoke_UpdateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func ServiceConnectionToken_ExpireTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func ServiceConnectionToken_ExpireTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func Timestamp_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	if in == nil {
		return nil
	}
	t := in.AsTime()
	s := t.Format(time.RFC3339Nano)
	return &s
}

func Timestamp_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	if in == nil {
		return nil
	}
	t, err := time.Parse(time.RFC3339Nano, *in)
	if err != nil {
		mapCtx.Errorf("parsing timestamp %q", *in)
		return nil
	}
	return timestamppb.New(t)
}

func GoogleRpcStatus_FromProto(mapCtx *direct.MapContext, in *pbgen.GoogleRpcStatus) *common.Status {
	if in == nil {
		return nil
	}
	return &common.Status{
		Code:    direct.LazyPtr(in.GetCode()),
		Message: direct.LazyPtr(in.GetMessage()),
	}
}

func GoogleRpcStatus_ToProto(mapCtx *direct.MapContext, in *common.Status) *pbgen.GoogleRpcStatus {
	if in == nil {
		return nil
	}
	return &pbgen.GoogleRpcStatus{
		Code:    direct.ValueOf(in.Code),
		Message: direct.ValueOf(in.Message),
	}
}

func StateMetadata_FromProto(mapCtx *direct.MapContext, in *pb.StateTimeline_StateMetadata) *krm.StateMetadata {
	if in == nil {
		return nil
	}
	out := &krm.StateMetadata{}
	out.EffectiveTime = Timestamp_FromProto(mapCtx, in.EffectiveTime)
	out.State = direct.Enum_FromProto(mapCtx, in.State)
	return out
}

func StateMetadata_ToProto(mapCtx *direct.MapContext, in *krm.StateMetadata) *pb.StateTimeline_StateMetadata {
	if in == nil {
		return nil
	}
	out := &pb.StateTimeline_StateMetadata{}
	out.EffectiveTime = Timestamp_ToProto(mapCtx, in.EffectiveTime)
	out.State = direct.Enum_ToProto[pb.StateTimeline_StateMetadata_State](mapCtx, in.State)
	return out
}

func Services_FromProto(mapCtx *direct.MapContext, in map[string]*pb.StateTimeline) map[string]krm.StateTimeline {
	if in == nil {
		return nil
	}
	out := make(map[string]krm.StateTimeline)
	for k, v := range in {
		val := StateTimeline_FromProto(mapCtx, v)
		if val != nil {
			out[k] = *val
		}
	}
	return out
}

func Services_ToProto(mapCtx *direct.MapContext, in map[string]krm.StateTimeline) map[string]*pb.StateTimeline {
	if in == nil {
		return nil
	}
	out := make(map[string]*pb.StateTimeline)
	for k, v := range in {
		out[k] = StateTimeline_ToProto(mapCtx, &v)
	}
	return out
}
