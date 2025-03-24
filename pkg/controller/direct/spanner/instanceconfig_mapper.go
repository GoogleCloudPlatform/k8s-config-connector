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
// krm.group: spanner.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.spanner.admin.instance.v1

package spanner

import (
	pb "cloud.google.com/go/spanner/admin/instance/apiv1/instancepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/spanner/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ReplicaInfo_FromProto(mapCtx *direct.MapContext, in *pb.ReplicaInfo) *krm.ReplicaInfo {
	if in == nil {
		return nil
	}
	out := &krm.ReplicaInfo{}
	out.Location = direct.LazyPtr(in.GetLocation())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.DefaultLeaderLocation = direct.LazyPtr(in.GetDefaultLeaderLocation())
	return out
}
func ReplicaInfo_ToProto(mapCtx *direct.MapContext, in *krm.ReplicaInfo) *pb.ReplicaInfo {
	if in == nil {
		return nil
	}
	out := &pb.ReplicaInfo{}
	out.Location = direct.ValueOf(in.Location)
	out.Type = direct.Enum_ToProto[pb.ReplicaInfo_ReplicaType](mapCtx, in.Type)
	out.DefaultLeaderLocation = direct.ValueOf(in.DefaultLeaderLocation)
	return out
}
func SpannerInstanceConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InstanceConfig) *krm.SpannerInstanceConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SpannerInstanceConfigObservedState{}
	// MISSING: Name
	out.ConfigType = direct.Enum_FromProto(mapCtx, in.GetConfigType())
	out.OptionalReplicas = direct.Slice_FromProto(mapCtx, in.OptionalReplicas, ReplicaInfo_FromProto)
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.FreeInstanceAvailability = direct.Enum_FromProto(mapCtx, in.GetFreeInstanceAvailability())
	out.QuorumType = direct.Enum_FromProto(mapCtx, in.GetQuorumType())
	out.StorageLimitPerProcessingUnit = direct.LazyPtr(in.GetStorageLimitPerProcessingUnit())
	return out
}
func SpannerInstanceConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SpannerInstanceConfigObservedState) *pb.InstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceConfig{}
	// MISSING: Name
	out.ConfigType = direct.Enum_ToProto[pb.InstanceConfig_Type](mapCtx, in.ConfigType)
	out.OptionalReplicas = direct.Slice_ToProto(mapCtx, in.OptionalReplicas, ReplicaInfo_ToProto)
	out.Reconciling = direct.ValueOf(in.Reconciling)
	out.State = direct.Enum_ToProto[pb.InstanceConfig_State](mapCtx, in.State)
	out.FreeInstanceAvailability = direct.Enum_ToProto[pb.InstanceConfig_FreeInstanceAvailability](mapCtx, in.FreeInstanceAvailability)
	out.QuorumType = direct.Enum_ToProto[pb.InstanceConfig_QuorumType](mapCtx, in.QuorumType)
	out.StorageLimitPerProcessingUnit = direct.ValueOf(in.StorageLimitPerProcessingUnit)
	return out
}
func SpannerInstanceConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.InstanceConfig) *krm.SpannerInstanceConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.SpannerInstanceConfigSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Replicas = direct.Slice_FromProto(mapCtx, in.Replicas, ReplicaInfo_FromProto)
	if in.GetBaseConfig() != "" {
		out.BaseConfigRef = &krm.InstanceConfigRef{External: in.BaseConfig}
	}
	out.Labels = in.Labels
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.LeaderOptions = in.LeaderOptions
	return out
}
func SpannerInstanceConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.SpannerInstanceConfigSpec) *pb.InstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceConfig{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Replicas = direct.Slice_ToProto(mapCtx, in.Replicas, ReplicaInfo_ToProto)
	if in.BaseConfigRef != nil {
		out.BaseConfig = in.BaseConfigRef.External
	}
	out.Labels = in.Labels
	out.Etag = direct.ValueOf(in.Etag)
	out.LeaderOptions = in.LeaderOptions
	return out
}
