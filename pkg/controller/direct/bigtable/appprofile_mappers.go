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

package bigtable

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	pb "google.golang.org/genproto/googleapis/bigtable/admin/v2"
)

func BigtableAppProfileSpec_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile) *krm.BigtableAppProfileSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigtableAppProfileSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.MultiClusterRoutingUseAny = AppProfile_MultiClusterRoutingUseAny_FromProto(mapCtx, in.GetMultiClusterRoutingUseAny())
	out.SingleClusterRouting = AppProfile_SingleClusterRouting_FromProto(mapCtx, in.GetSingleClusterRouting())
	out.Priority = direct.Enum_FromProto(mapCtx, in.GetPriority())
	out.StandardIsolation = AppProfile_StandardIsolation_FromProto(mapCtx, in.GetStandardIsolation())
	out.DataBoostIsolationReadOnly = AppProfile_DataBoostIsolationReadOnly_FromProto(mapCtx, in.GetDataBoostIsolationReadOnly())
	return out
}
func BigtableAppProfileSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigtableAppProfileSpec) *pb.AppProfile {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile{}
	out.Description = direct.ValueOf(in.Description)
	if oneof := AppProfile_MultiClusterRoutingUseAny_ToProto(mapCtx, in.MultiClusterRoutingUseAny); oneof != nil {
		out.RoutingPolicy = &pb.AppProfile_MultiClusterRoutingUseAny_{MultiClusterRoutingUseAny: oneof}
	}
	if oneof := AppProfile_SingleClusterRouting_ToProto(mapCtx, in.SingleClusterRouting); oneof != nil {
		out.RoutingPolicy = &pb.AppProfile_SingleClusterRouting_{SingleClusterRouting: oneof}
	}
	if oneof := AppProfile_Priority_ToProto(mapCtx, in.Priority); oneof != nil {
		out.Isolation = oneof
	}
	if oneof := AppProfile_StandardIsolation_ToProto(mapCtx, in.StandardIsolation); oneof != nil {
		out.Isolation = &pb.AppProfile_StandardIsolation_{StandardIsolation: oneof}
	}
	if oneof := AppProfile_DataBoostIsolationReadOnly_ToProto(mapCtx, in.DataBoostIsolationReadOnly); oneof != nil {
		out.Isolation = &pb.AppProfile_DataBoostIsolationReadOnly_{DataBoostIsolationReadOnly: oneof}
	}
	return out
}
func BigtableAppProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile) *krm.BigtableAppProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigtableAppProfileObservedState{}
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func BigtableAppProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigtableAppProfileObservedState) *pb.AppProfile {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile{}
	out.Etag = direct.ValueOf(in.Etag)
	return out
}

func AppProfile_MultiClusterRoutingUseAny_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_MultiClusterRoutingUseAny) *krm.AppProfile_MultiClusterRoutingUseAny {
	if in == nil {
		return nil
	}
	out := &krm.AppProfile_MultiClusterRoutingUseAny{}
	out.ClusterIds = in.ClusterIds
	return out
}
func AppProfile_MultiClusterRoutingUseAny_ToProto(mapCtx *direct.MapContext, in *krm.AppProfile_MultiClusterRoutingUseAny) *pb.AppProfile_MultiClusterRoutingUseAny {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_MultiClusterRoutingUseAny{}
	out.ClusterIds = in.ClusterIds
	return out
}
func AppProfile_SingleClusterRouting_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_SingleClusterRouting) *krm.AppProfile_SingleClusterRouting {
	if in == nil {
		return nil
	}
	out := &krm.AppProfile_SingleClusterRouting{}
	out.ClusterID = direct.LazyPtr(in.GetClusterId())
	out.AllowTransactionalWrites = direct.LazyPtr(in.GetAllowTransactionalWrites())
	return out
}
func AppProfile_SingleClusterRouting_ToProto(mapCtx *direct.MapContext, in *krm.AppProfile_SingleClusterRouting) *pb.AppProfile_SingleClusterRouting {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_SingleClusterRouting{}
	out.ClusterId = direct.ValueOf(in.ClusterID)
	out.AllowTransactionalWrites = direct.ValueOf(in.AllowTransactionalWrites)
	return out
}
func AppProfile_StandardIsolation_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_StandardIsolation) *krm.AppProfile_StandardIsolation {
	if in == nil {
		return nil
	}
	out := &krm.AppProfile_StandardIsolation{}
	out.Priority = direct.Enum_FromProto(mapCtx, in.GetPriority())
	return out
}
func AppProfile_StandardIsolation_ToProto(mapCtx *direct.MapContext, in *krm.AppProfile_StandardIsolation) *pb.AppProfile_StandardIsolation {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_StandardIsolation{}
	out.Priority = direct.Enum_ToProto[pb.AppProfile_Priority](mapCtx, in.Priority)
	return out
}
func AppProfile_DataBoostIsolationReadOnly_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_DataBoostIsolationReadOnly) *krm.AppProfile_DataBoostIsolationReadOnly {
	if in == nil {
		return nil
	}
	out := &krm.AppProfile_DataBoostIsolationReadOnly{}
	out.ComputeBillingOwner = direct.Enum_FromProto(mapCtx, in.GetComputeBillingOwner())
	return out
}
func AppProfile_DataBoostIsolationReadOnly_ToProto(mapCtx *direct.MapContext, in *krm.AppProfile_DataBoostIsolationReadOnly) *pb.AppProfile_DataBoostIsolationReadOnly {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_DataBoostIsolationReadOnly{}
	if oneof := AppProfile_DataBoostIsolationReadOnly_ComputeBillingOwner_ToProto(mapCtx, in.ComputeBillingOwner); oneof != nil {
		out.ComputeBillingOwner = oneof
	}
	return out
}
