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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataplexEntryGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EntryGroup) *krm.DataplexEntryGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataplexEntryGroupObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: TransferStatus
	return out
}
func DataplexEntryGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataplexEntryGroupObservedState) *pb.EntryGroup {
	if in == nil {
		return nil
	}
	out := &pb.EntryGroup{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: TransferStatus
	return out
}
func DataplexEntryGroupSpec_FromProto(mapCtx *direct.MapContext, in *pb.EntryGroup) *krm.DataplexEntryGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataplexEntryGroupSpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: TransferStatus
	return out
}
func DataplexEntryGroupSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataplexEntryGroupSpec) *pb.EntryGroup {
	if in == nil {
		return nil
	}
	out := &pb.EntryGroup{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: TransferStatus
	return out
}
func EntryGroup_FromProto(mapCtx *direct.MapContext, in *pb.EntryGroup) *krm.EntryGroup {
	if in == nil {
		return nil
	}
	out := &krm.EntryGroup{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Labels = in.Labels
	out.Etag = direct.LazyPtr(in.GetEtag())
	// MISSING: TransferStatus
	return out
}
func EntryGroup_ToProto(mapCtx *direct.MapContext, in *krm.EntryGroup) *pb.EntryGroup {
	if in == nil {
		return nil
	}
	out := &pb.EntryGroup{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.ValueOf(in.Description)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Labels = in.Labels
	out.Etag = direct.ValueOf(in.Etag)
	// MISSING: TransferStatus
	return out
}
func EntryGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EntryGroup) *krm.EntryGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EntryGroupObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: Etag
	out.TransferStatus = direct.Enum_FromProto(mapCtx, in.GetTransferStatus())
	return out
}
func EntryGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EntryGroupObservedState) *pb.EntryGroup {
	if in == nil {
		return nil
	}
	out := &pb.EntryGroup{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: Etag
	out.TransferStatus = direct.Enum_ToProto[pb.TransferStatus](mapCtx, in.TransferStatus)
	return out
}
