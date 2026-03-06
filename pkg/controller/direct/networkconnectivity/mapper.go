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

	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NetworkConnectivityServiceConnectionPolicySpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.ServiceConnectionPolicy) *krm.NetworkConnectivityServiceConnectionPolicySpec {
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
		out.Network = &computev1beta1.ComputeNetworkRef{External: in.Network}
	}
	out.PscConfig = PscConfig_v1alpha1_FromProto(mapCtx, in.GetPscConfig())
	// MISSING: PscConnections
	out.ServiceClass = direct.LazyPtr(in.GetServiceClass())
	// MISSING: UpdateTime
	return out
}
func NetworkConnectivityServiceConnectionPolicySpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConnectivityServiceConnectionPolicySpec) *pb.ServiceConnectionPolicy {
	if in == nil {
		return nil
	}
	out := &pb.ServiceConnectionPolicy{}
	out.Description = direct.ValueOf(in.Description)
	if in.Network != nil {
		out.Network = in.Network.External
	}
	out.PscConfig = PscConfig_v1alpha1_ToProto(mapCtx, in.PscConfig)
	out.ServiceClass = direct.ValueOf(in.ServiceClass)
	return out
}

func PscConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.PscConfig) *krm.PscConfig {
	if in == nil {
		return nil
	}
	out := &krm.PscConfig{}
	out.Limit = direct.LazyPtr(in.GetLimit())
	out.ProducerInstanceLocation = direct.LazyPtr(in.GetProducerInstanceLocation())
	out.Subnetworks = PscConfig_Subnetworks_v1alpha1_FromProto(mapCtx, in.GetSubnetworks())
	return out
}

func PscConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.PscConfig) *pb.PscConfig {
	if in == nil {
		return nil
	}
	out := &pb.PscConfig{}
	out.Limit = direct.ValueOf(in.Limit)
	out.ProducerInstanceLocation = direct.ValueOf(in.ProducerInstanceLocation)
	out.Subnetworks = PscConfig_Subnetworks_v1alpha1_ToProto(mapCtx, in.Subnetworks)
	return out
}

func PscConfig_Subnetworks_v1alpha1_FromProto(mapCtx *direct.MapContext, in []string) []refs.ComputeSubnetworkRef {
	if in == nil {
		return nil
	}
	var out []refs.ComputeSubnetworkRef
	for _, s := range in {
		out = append(out, refs.ComputeSubnetworkRef{External: s})
	}
	return out
}
func PscConfig_Subnetworks_v1alpha1_ToProto(mapCtx *direct.MapContext, in []refs.ComputeSubnetworkRef) []string {
	if in == nil {
		return nil
	}
	var out []string
	for _, ref := range in {
		out = append(out, ref.External)
	}
	return out
}

func Group_v1alpha1_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func Group_v1alpha1_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func Group_v1alpha1_UpdateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func Group_v1alpha1_UpdateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func InternalRange_v1alpha1_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func InternalRange_v1alpha1_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func InternalRange_v1alpha1_UpdateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func InternalRange_v1alpha1_UpdateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func Hub_v1alpha1_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func Hub_v1alpha1_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func Hub_v1alpha1_UpdateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func Hub_v1alpha1_UpdateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func ServiceConnectionPolicy_v1alpha1_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	return Timestamp_FromProto(mapCtx, in)
}

func ServiceConnectionPolicy_v1alpha1_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	return Timestamp_ToProto(mapCtx, in)
}

func ServiceConnectionPolicy_v1alpha1_UpdateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	return Timestamp_FromProto(mapCtx, in)
}
func ServiceConnectionPolicy_v1alpha1_UpdateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	return Timestamp_ToProto(mapCtx, in)
}

func Policy_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Policy) *krm.Policy {
	if in == nil {
		return nil
	}
	out := &krm.Policy{}
	out.AuditConfigs = direct.Slice_FromProto(mapCtx, in.AuditConfigs, AuditConfig_v1alpha1_FromProto)
	out.Bindings = direct.Slice_FromProto(mapCtx, in.Bindings, Binding_v1alpha1_FromProto)
	out.Etag = in.GetEtag()
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}

func Policy_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.Policy) *pb.Policy {
	if in == nil {
		return nil
	}
	out := &pb.Policy{}
	out.AuditConfigs = direct.Slice_ToProto(mapCtx, in.AuditConfigs, AuditConfig_v1alpha1_ToProto)
	out.Bindings = direct.Slice_ToProto(mapCtx, in.Bindings, Binding_v1alpha1_ToProto)
	out.Etag = in.Etag
	out.Version = direct.ValueOf(in.Version)
	return out
}

func OperationMetadata_v1alpha1_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func OperationMetadata_v1alpha1_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func OperationMetadata_v1alpha1_EndTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func OperationMetadata_v1alpha1_EndTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func RegionalEndpoint_v1alpha1_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func RegionalEndpoint_v1alpha1_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func RegionalEndpoint_v1alpha1_UpdateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func RegionalEndpoint_v1alpha1_UpdateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func PolicyBasedRoute_v1alpha1_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func PolicyBasedRoute_v1alpha1_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func PolicyBasedRoute_v1alpha1_UpdateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func PolicyBasedRoute_v1alpha1_UpdateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func Route_v1alpha1_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func Route_v1alpha1_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func Route_v1alpha1_UpdateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func Route_v1alpha1_UpdateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func RouteTable_v1alpha1_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func RouteTable_v1alpha1_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func RouteTable_v1alpha1_UpdateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func RouteTable_v1alpha1_UpdateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func ServiceClass_v1alpha1_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func ServiceClass_v1alpha1_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func ServiceClass_v1alpha1_UpdateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func ServiceClass_v1alpha1_UpdateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func ServiceConnectionMap_v1alpha1_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func ServiceConnectionMap_v1alpha1_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func ServiceConnectionMap_v1alpha1_UpdateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func ServiceConnectionMap_v1alpha1_UpdateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func ServiceConnectionToken_v1alpha1_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func ServiceConnectionToken_v1alpha1_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func ServiceConnectionToken_v1alpha1_UpdateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func ServiceConnectionToken_v1alpha1_UpdateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func Spoke_v1alpha1_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func Spoke_v1alpha1_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func Spoke_v1alpha1_UpdateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func Spoke_v1alpha1_UpdateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func ServiceConnectionToken_v1alpha1_ExpireTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func ServiceConnectionToken_v1alpha1_ExpireTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
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
