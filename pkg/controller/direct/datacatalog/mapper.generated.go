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
// krm.group: datacatalog.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.datacatalog.v1

package datacatalog

import (
	pb "cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datacatalog/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataCatalogEntryGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EntryGroup) *krm.DataCatalogEntryGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataCatalogEntryGroupObservedState{}
	// MISSING: Name
	out.DataCatalogTimestamps = SystemTimestamps_FromProto(mapCtx, in.GetDataCatalogTimestamps())
	return out
}
func DataCatalogEntryGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataCatalogEntryGroupObservedState) *pb.EntryGroup {
	if in == nil {
		return nil
	}
	out := &pb.EntryGroup{}
	// MISSING: Name
	out.DataCatalogTimestamps = SystemTimestamps_ToProto(mapCtx, in.DataCatalogTimestamps)
	return out
}
func DataCatalogEntryGroupSpec_FromProto(mapCtx *direct.MapContext, in *pb.EntryGroup) *krm.DataCatalogEntryGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataCatalogEntryGroupSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.TransferredToDataplex = direct.LazyPtr(in.GetTransferredToDataplex())
	return out
}
func DataCatalogEntryGroupSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataCatalogEntryGroupSpec) *pb.EntryGroup {
	if in == nil {
		return nil
	}
	out := &pb.EntryGroup{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.TransferredToDataplex = direct.ValueOf(in.TransferredToDataplex)
	return out
}
func SystemTimestamps_FromProto(mapCtx *direct.MapContext, in *pb.SystemTimestamps) *krm.SystemTimestamps {
	if in == nil {
		return nil
	}
	out := &krm.SystemTimestamps{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: ExpireTime
	return out
}
func SystemTimestamps_ToProto(mapCtx *direct.MapContext, in *krm.SystemTimestamps) *pb.SystemTimestamps {
	if in == nil {
		return nil
	}
	out := &pb.SystemTimestamps{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: ExpireTime
	return out
}
func SystemTimestampsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SystemTimestamps) *krm.SystemTimestampsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SystemTimestampsObservedState{}
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	return out
}
func SystemTimestampsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SystemTimestampsObservedState) *pb.SystemTimestamps {
	if in == nil {
		return nil
	}
	out := &pb.SystemTimestamps{}
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	return out
}
