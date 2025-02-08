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

package datacatalog

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datacatalog/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
)
func DatacatalogEntryGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EntryGroup) *krm.DatacatalogEntryGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatacatalogEntryGroupObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DataCatalogTimestamps
	// MISSING: TransferredToDataplex
	return out
}
func DatacatalogEntryGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatacatalogEntryGroupObservedState) *pb.EntryGroup {
	if in == nil {
		return nil
	}
	out := &pb.EntryGroup{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DataCatalogTimestamps
	// MISSING: TransferredToDataplex
	return out
}
func DatacatalogEntryGroupSpec_FromProto(mapCtx *direct.MapContext, in *pb.EntryGroup) *krm.DatacatalogEntryGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatacatalogEntryGroupSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DataCatalogTimestamps
	// MISSING: TransferredToDataplex
	return out
}
func DatacatalogEntryGroupSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatacatalogEntryGroupSpec) *pb.EntryGroup {
	if in == nil {
		return nil
	}
	out := &pb.EntryGroup{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DataCatalogTimestamps
	// MISSING: TransferredToDataplex
	return out
}
func EntryGroup_FromProto(mapCtx *direct.MapContext, in *pb.EntryGroup) *krm.EntryGroup {
	if in == nil {
		return nil
	}
	out := &krm.EntryGroup{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: DataCatalogTimestamps
	out.TransferredToDataplex = direct.LazyPtr(in.GetTransferredToDataplex())
	return out
}
func EntryGroup_ToProto(mapCtx *direct.MapContext, in *krm.EntryGroup) *pb.EntryGroup {
	if in == nil {
		return nil
	}
	out := &pb.EntryGroup{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: DataCatalogTimestamps
	out.TransferredToDataplex = direct.ValueOf(in.TransferredToDataplex)
	return out
}
func EntryGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EntryGroup) *krm.EntryGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EntryGroupObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	out.DataCatalogTimestamps = SystemTimestamps_FromProto(mapCtx, in.GetDataCatalogTimestamps())
	// MISSING: TransferredToDataplex
	return out
}
func EntryGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EntryGroupObservedState) *pb.EntryGroup {
	if in == nil {
		return nil
	}
	out := &pb.EntryGroup{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	out.DataCatalogTimestamps = SystemTimestamps_ToProto(mapCtx, in.DataCatalogTimestamps)
	// MISSING: TransferredToDataplex
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
