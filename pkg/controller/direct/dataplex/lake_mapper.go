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

package dataplex

import (
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	dataprocv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AssetStatus_FromProto(mapCtx *direct.MapContext, in *pb.AssetStatus) *krm.AssetStatus {
	if in == nil {
		return nil
	}
	out := &krm.AssetStatus{}
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.ActiveAssets = direct.LazyPtr(in.GetActiveAssets())
	out.SecurityPolicyApplyingAssets = direct.LazyPtr(in.GetSecurityPolicyApplyingAssets())
	return out
}
func AssetStatus_ToProto(mapCtx *direct.MapContext, in *krm.AssetStatus) *pb.AssetStatus {
	if in == nil {
		return nil
	}
	out := &pb.AssetStatus{}
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.ActiveAssets = direct.ValueOf(in.ActiveAssets)
	out.SecurityPolicyApplyingAssets = direct.ValueOf(in.SecurityPolicyApplyingAssets)
	return out
}
func DataplexLakeSpec_FromProto(mapCtx *direct.MapContext, in *pb.Lake) *krm.DataplexLakeSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataplexLakeSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: State
	// MISSING: ServiceAccount
	out.Metastore = Lake_Metastore_FromProto(mapCtx, in.GetMetastore())
	// MISSING: AssetStatus
	// MISSING: MetastoreStatus
	return out
}
func DataplexLakeSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataplexLakeSpec) *pb.Lake {
	if in == nil {
		return nil
	}
	out := &pb.Lake{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	// MISSING: State
	// MISSING: ServiceAccount
	out.Metastore = Lake_Metastore_ToProto(mapCtx, in.Metastore)
	// MISSING: AssetStatus
	// MISSING: MetastoreStatus
	return out
}
func DataplexLakeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Lake) *krm.DataplexLakeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataplexLakeObservedState{}
	// MISSING: DisplayName
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	// MISSING: Metastore
	out.AssetStatus = AssetStatus_FromProto(mapCtx, in.GetAssetStatus())
	out.MetastoreStatus = Lake_MetastoreStatus_FromProto(mapCtx, in.GetMetastoreStatus())
	return out
}
func DataplexLakeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataplexLakeObservedState) *pb.Lake {
	if in == nil {
		return nil
	}
	out := &pb.Lake{}
	// MISSING: DisplayName
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	out.State = direct.Enum_ToProto[pb.State](mapCtx, in.State)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	// MISSING: Metastore
	out.AssetStatus = AssetStatus_ToProto(mapCtx, in.AssetStatus)
	out.MetastoreStatus = Lake_MetastoreStatus_ToProto(mapCtx, in.MetastoreStatus)
	return out
}
func Lake_Metastore_FromProto(mapCtx *direct.MapContext, in *pb.Lake_Metastore) *krm.Lake_Metastore {
	if in == nil {
		return nil
	}
	out := &krm.Lake_Metastore{}
	if in.GetService() != "" {
		out.ServiceRef = &dataprocv1alpha1.ServiceRef{External: in.GetService()}
	}
	return out
}
func Lake_Metastore_ToProto(mapCtx *direct.MapContext, in *krm.Lake_Metastore) *pb.Lake_Metastore {
	if in == nil {
		return nil
	}
	out := &pb.Lake_Metastore{}
	if in.ServiceRef != nil {
		out.Service = in.ServiceRef.External
	}
	return out
}
func Lake_MetastoreStatus_FromProto(mapCtx *direct.MapContext, in *pb.Lake_MetastoreStatus) *krm.Lake_MetastoreStatus {
	if in == nil {
		return nil
	}
	out := &krm.Lake_MetastoreStatus{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Message = direct.LazyPtr(in.GetMessage())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Endpoint = direct.LazyPtr(in.GetEndpoint())
	return out
}
func Lake_MetastoreStatus_ToProto(mapCtx *direct.MapContext, in *krm.Lake_MetastoreStatus) *pb.Lake_MetastoreStatus {
	if in == nil {
		return nil
	}
	out := &pb.Lake_MetastoreStatus{}
	out.State = direct.Enum_ToProto[pb.Lake_MetastoreStatus_State](mapCtx, in.State)
	out.Message = direct.ValueOf(in.Message)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Endpoint = direct.ValueOf(in.Endpoint)
	return out
}
