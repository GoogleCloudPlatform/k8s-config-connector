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

package tpu

import (
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/tpu/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpclients/generated/google/cloud/tpu/v2"
	"google.golang.org/genproto/googleapis/type/interval"
)

// ServiceAccount uses email for serviceAccount

func ServiceAccount_FromProto(mapCtx *direct.MapContext, in *pb.ServiceAccount) *krm.ServiceAccount {
	if in == nil {
		return nil
	}
	out := &krm.ServiceAccount{}
	if in.GetEmail() != "" {
		out.ServiceAccountRef = &refs.IAMServiceAccountRef{
			External: in.GetEmail(),
		}
	}
	out.Scope = in.Scope
	return out
}

func ServiceAccount_ToProto(mapCtx *direct.MapContext, in *krm.ServiceAccount) *pb.ServiceAccount {
	if in == nil {
		return nil
	}
	out := &pb.ServiceAccount{}
	if in.ServiceAccountRef != nil {
		out.Email = in.ServiceAccountRef.External
	}
	out.Scope = in.Scope
	return out
}

func Interval_FromProto(mapCtx *direct.MapContext, in *interval.Interval) *krm.Interval {
	if in == nil {
		return nil
	}
	out := &krm.Interval{}
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}

func Interval_ToProto(mapCtx *direct.MapContext, in *krm.Interval) *interval.Interval {
	if in == nil {
		return nil
	}
	out := &interval.Interval{}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	return out
}

func TPUQueuedResourceSpec_FromProto(mapCtx *direct.MapContext, in *pb.QueuedResource) *krm.TPUQueuedResourceSpec {
	if in == nil {
		return nil
	}
	out := &krm.TPUQueuedResourceSpec{}
	out.Tpu = QueuedResource_Tpu_FromProto(mapCtx, in.GetTpu())
	out.Spot = QueuedResource_Spot_FromProto(mapCtx, in.GetSpot())
	out.Guaranteed = QueuedResource_Guaranteed_FromProto(mapCtx, in.GetGuaranteed())
	out.QueueingPolicy = QueuedResource_QueueingPolicy_FromProto(mapCtx, in.GetQueueingPolicy())
	if in.GetReservationName() != "" {
		out.ReservationRef = &computev1beta1.ComputeReservationRef{
			External: in.GetReservationName(),
		}
	}
	return out
}

func TPUQueuedResourceSpec_ToProto(mapCtx *direct.MapContext, in *krm.TPUQueuedResourceSpec) *pb.QueuedResource {
	if in == nil {
		return nil
	}
	out := &pb.QueuedResource{}
	if oneof := QueuedResource_Tpu_ToProto(mapCtx, in.Tpu); oneof != nil {
		out.Resource = &pb.QueuedResource_Tpu_{Tpu: oneof}
	}
	if oneof := QueuedResource_Spot_ToProto(mapCtx, in.Spot); oneof != nil {
		out.Tier = &pb.QueuedResource_Spot_{Spot: oneof}
	}
	if oneof := QueuedResource_Guaranteed_ToProto(mapCtx, in.Guaranteed); oneof != nil {
		out.Tier = &pb.QueuedResource_Guaranteed_{Guaranteed: oneof}
	}
	out.QueueingPolicy = QueuedResource_QueueingPolicy_ToProto(mapCtx, in.QueueingPolicy)
	if in.ReservationRef != nil {
		out.ReservationName = in.ReservationRef.External
	}
	return out
}

func TPUQueuedResourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QueuedResource) *krm.TPUQueuedResourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TPUQueuedResourceObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.Tpu = QueuedResource_TpuObservedState_FromProto(mapCtx, in.GetTpu())
	out.State = QueuedResourceStateObservedState_FromProto(mapCtx, in.GetState())
	return out
}

func TPUQueuedResourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TPUQueuedResourceObservedState) *pb.QueuedResource {
	if in == nil {
		return nil
	}
	out := &pb.QueuedResource{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	if oneof := QueuedResource_TpuObservedState_ToProto(mapCtx, in.Tpu); oneof != nil {
		out.Resource = &pb.QueuedResource_Tpu_{Tpu: oneof}
	}
	out.State = QueuedResourceStateObservedState_ToProto(mapCtx, in.State)
	return out
}
