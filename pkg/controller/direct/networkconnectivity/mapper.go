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

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/v1alpha1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
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
func NetworkConnectivityServiceConnectionPolicySpec_Network_ToProto(mapCtx *direct.MapContext, in *computerefs.ComputeNetworkRef) string {
	if in == nil {
		return ""
	}
	return in.External
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

func Services_FromProto(mapCtx *direct.MapContext, in map[string]*pb.StateTimeline) map[string]krm.StateTimeline {
	if in == nil {
		return nil
	}
	out := make(map[string]krm.StateTimeline)
	for k := range in {
		out[k] = krm.StateTimeline{}
	}
	return out
}

func Services_ToProto(mapCtx *direct.MapContext, in map[string]krm.StateTimeline) map[string]*pb.StateTimeline {
	if in == nil {
		return nil
	}
	out := make(map[string]*pb.StateTimeline)
	for k := range in {
		out[k] = &pb.StateTimeline{}
	}
	return out
}

func ServicesObservedState_FromProto(mapCtx *direct.MapContext, in map[string]*pb.StateTimeline) map[string]krm.StateTimelineObservedState {
	if in == nil {
		return nil
	}
	out := make(map[string]krm.StateTimelineObservedState)
	for k, v := range in {
		st := StateTimelineObservedState_FromProto(mapCtx, v)
		if st != nil {
			out[k] = *st
		}
	}
	return out
}

func ServicesObservedState_ToProto(mapCtx *direct.MapContext, in map[string]krm.StateTimelineObservedState) map[string]*pb.StateTimeline {
	if in == nil {
		return nil
	}
	out := make(map[string]*pb.StateTimeline)
	for k, v := range in {
		out[k] = StateTimelineObservedState_ToProto(mapCtx, &v)
	}
	return out
}

func StateMetadataObservedState_FromProto(mapCtx *direct.MapContext, in *pb.StateMetadata) *krm.StateMetadataObservedState {
	if in == nil {
		return nil
	}
	out := &krm.StateMetadataObservedState{}
	out.EffectiveTime = Timestamp_FromProto(mapCtx, in.GetEffectiveTime())
	out.State = direct.LazyPtr(in.GetState())
	return out
}

func StateMetadataObservedState_ToProto(mapCtx *direct.MapContext, in *krm.StateMetadataObservedState) *pb.StateMetadata {
	if in == nil {
		return nil
	}
	out := &pb.StateMetadata{}
	out.EffectiveTime = Timestamp_ToProto(mapCtx, in.EffectiveTime)
	out.State = direct.ValueOf(in.State)
	return out
}

func NetworkConnectivityMulticloudDataTransferConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MulticloudDataTransferConfig) *krm.NetworkConnectivityMulticloudDataTransferConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConnectivityMulticloudDataTransferConfigObservedState{}
	out.CreateTime = Timestamp_FromProto(mapCtx, in.GetCreateTime())
	out.DestinationsActiveCount = direct.LazyPtr(in.GetDestinationsActiveCount())
	out.DestinationsCount = direct.LazyPtr(in.GetDestinationsCount())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Services = ServicesObservedState_FromProto(mapCtx, in.Services)
	out.Uid = direct.LazyPtr(in.GetUid())
	out.UpdateTime = Timestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func NetworkConnectivityMulticloudDataTransferConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConnectivityMulticloudDataTransferConfigObservedState) *pb.MulticloudDataTransferConfig {
	if in == nil {
		return nil
	}
	out := &pb.MulticloudDataTransferConfig{}
	out.CreateTime = Timestamp_ToProto(mapCtx, in.CreateTime)
	out.DestinationsActiveCount = direct.ValueOf(in.DestinationsActiveCount)
	out.DestinationsCount = direct.ValueOf(in.DestinationsCount)
	out.Etag = direct.ValueOf(in.Etag)
	out.Services = ServicesObservedState_ToProto(mapCtx, in.Services)
	out.Uid = direct.ValueOf(in.Uid)
	out.UpdateTime = Timestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
