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

func EntryType_FromProto(mapCtx *direct.MapContext, in *pb.EntryType) *krm.EntryType {
	if in == nil {
		return nil
	}
	out := &krm.EntryType{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Labels = in.Labels
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.TypeAliases = in.TypeAliases
	out.Platform = direct.LazyPtr(in.GetPlatform())
	out.System = direct.LazyPtr(in.GetSystem())
	out.RequiredAspects = direct.Slice_FromProto(mapCtx, in.RequiredAspects, EntryType_AspectInfo_FromProto)
	out.Authorization = EntryType_Authorization_FromProto(mapCtx, in.GetAuthorization())
	return out
}
func EntryType_ToProto(mapCtx *direct.MapContext, in *krm.EntryType) *pb.EntryType {
	if in == nil {
		return nil
	}
	out := &pb.EntryType{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.ValueOf(in.Description)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Labels = in.Labels
	out.Etag = direct.ValueOf(in.Etag)
	out.TypeAliases = in.TypeAliases
	out.Platform = direct.ValueOf(in.Platform)
	out.System = direct.ValueOf(in.System)
	out.RequiredAspects = direct.Slice_ToProto(mapCtx, in.RequiredAspects, EntryType_AspectInfo_ToProto)
	out.Authorization = EntryType_Authorization_ToProto(mapCtx, in.Authorization)
	return out
}
func EntryTypeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EntryType) *krm.EntryTypeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EntryTypeObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: TypeAliases
	// MISSING: Platform
	// MISSING: System
	// MISSING: RequiredAspects
	// MISSING: Authorization
	return out
}
func EntryTypeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EntryTypeObservedState) *pb.EntryType {
	if in == nil {
		return nil
	}
	out := &pb.EntryType{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: TypeAliases
	// MISSING: Platform
	// MISSING: System
	// MISSING: RequiredAspects
	// MISSING: Authorization
	return out
}
