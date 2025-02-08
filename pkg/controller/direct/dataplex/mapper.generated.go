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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
)
func Aspect_FromProto(mapCtx *direct.MapContext, in *pb.Aspect) *krm.Aspect {
	if in == nil {
		return nil
	}
	out := &krm.Aspect{}
	// MISSING: AspectType
	// MISSING: Path
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Data = Data_FromProto(mapCtx, in.GetData())
	out.AspectSource = AspectSource_FromProto(mapCtx, in.GetAspectSource())
	return out
}
func Aspect_ToProto(mapCtx *direct.MapContext, in *krm.Aspect) *pb.Aspect {
	if in == nil {
		return nil
	}
	out := &pb.Aspect{}
	// MISSING: AspectType
	// MISSING: Path
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Data = Data_ToProto(mapCtx, in.Data)
	out.AspectSource = AspectSource_ToProto(mapCtx, in.AspectSource)
	return out
}
func AspectSource_FromProto(mapCtx *direct.MapContext, in *pb.AspectSource) *krm.AspectSource {
	if in == nil {
		return nil
	}
	out := &krm.AspectSource{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DataVersion = direct.LazyPtr(in.GetDataVersion())
	return out
}
func AspectSource_ToProto(mapCtx *direct.MapContext, in *krm.AspectSource) *pb.AspectSource {
	if in == nil {
		return nil
	}
	out := &pb.AspectSource{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DataVersion = direct.ValueOf(in.DataVersion)
	return out
}
func DataplexEntryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Entry) *krm.DataplexEntryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataplexEntryObservedState{}
	// MISSING: Name
	// MISSING: EntryType
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Aspects
	// MISSING: ParentEntry
	// MISSING: FullyQualifiedName
	// MISSING: EntrySource
	return out
}
func DataplexEntryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataplexEntryObservedState) *pb.Entry {
	if in == nil {
		return nil
	}
	out := &pb.Entry{}
	// MISSING: Name
	// MISSING: EntryType
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Aspects
	// MISSING: ParentEntry
	// MISSING: FullyQualifiedName
	// MISSING: EntrySource
	return out
}
func DataplexEntrySpec_FromProto(mapCtx *direct.MapContext, in *pb.Entry) *krm.DataplexEntrySpec {
	if in == nil {
		return nil
	}
	out := &krm.DataplexEntrySpec{}
	// MISSING: Name
	// MISSING: EntryType
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Aspects
	// MISSING: ParentEntry
	// MISSING: FullyQualifiedName
	// MISSING: EntrySource
	return out
}
func DataplexEntrySpec_ToProto(mapCtx *direct.MapContext, in *krm.DataplexEntrySpec) *pb.Entry {
	if in == nil {
		return nil
	}
	out := &pb.Entry{}
	// MISSING: Name
	// MISSING: EntryType
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Aspects
	// MISSING: ParentEntry
	// MISSING: FullyQualifiedName
	// MISSING: EntrySource
	return out
}
func Entry_FromProto(mapCtx *direct.MapContext, in *pb.Entry) *krm.Entry {
	if in == nil {
		return nil
	}
	out := &krm.Entry{}
	out.Name = direct.LazyPtr(in.GetName())
	out.EntryType = direct.LazyPtr(in.GetEntryType())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Aspects
	out.ParentEntry = direct.LazyPtr(in.GetParentEntry())
	out.FullyQualifiedName = direct.LazyPtr(in.GetFullyQualifiedName())
	out.EntrySource = EntrySource_FromProto(mapCtx, in.GetEntrySource())
	return out
}
func Entry_ToProto(mapCtx *direct.MapContext, in *krm.Entry) *pb.Entry {
	if in == nil {
		return nil
	}
	out := &pb.Entry{}
	out.Name = direct.ValueOf(in.Name)
	out.EntryType = direct.ValueOf(in.EntryType)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Aspects
	out.ParentEntry = direct.ValueOf(in.ParentEntry)
	out.FullyQualifiedName = direct.ValueOf(in.FullyQualifiedName)
	out.EntrySource = EntrySource_ToProto(mapCtx, in.EntrySource)
	return out
}
func EntryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Entry) *krm.EntryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EntryObservedState{}
	// MISSING: Name
	// MISSING: EntryType
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Aspects
	// MISSING: ParentEntry
	// MISSING: FullyQualifiedName
	out.EntrySource = EntrySourceObservedState_FromProto(mapCtx, in.GetEntrySource())
	return out
}
func EntryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EntryObservedState) *pb.Entry {
	if in == nil {
		return nil
	}
	out := &pb.Entry{}
	// MISSING: Name
	// MISSING: EntryType
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Aspects
	// MISSING: ParentEntry
	// MISSING: FullyQualifiedName
	out.EntrySource = EntrySourceObservedState_ToProto(mapCtx, in.EntrySource)
	return out
}
func EntrySource_FromProto(mapCtx *direct.MapContext, in *pb.EntrySource) *krm.EntrySource {
	if in == nil {
		return nil
	}
	out := &krm.EntrySource{}
	out.Resource = direct.LazyPtr(in.GetResource())
	out.System = direct.LazyPtr(in.GetSystem())
	out.Platform = direct.LazyPtr(in.GetPlatform())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	out.Ancestors = direct.Slice_FromProto(mapCtx, in.Ancestors, EntrySource_Ancestor_FromProto)
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Location
	return out
}
func EntrySource_ToProto(mapCtx *direct.MapContext, in *krm.EntrySource) *pb.EntrySource {
	if in == nil {
		return nil
	}
	out := &pb.EntrySource{}
	out.Resource = direct.ValueOf(in.Resource)
	out.System = direct.ValueOf(in.System)
	out.Platform = direct.ValueOf(in.Platform)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	out.Ancestors = direct.Slice_ToProto(mapCtx, in.Ancestors, EntrySource_Ancestor_ToProto)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Location
	return out
}
func EntrySourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EntrySource) *krm.EntrySourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EntrySourceObservedState{}
	// MISSING: Resource
	// MISSING: System
	// MISSING: Platform
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Labels
	// MISSING: Ancestors
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Location = direct.LazyPtr(in.GetLocation())
	return out
}
func EntrySourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EntrySourceObservedState) *pb.EntrySource {
	if in == nil {
		return nil
	}
	out := &pb.EntrySource{}
	// MISSING: Resource
	// MISSING: System
	// MISSING: Platform
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Labels
	// MISSING: Ancestors
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Location = direct.ValueOf(in.Location)
	return out
}
func EntrySource_Ancestor_FromProto(mapCtx *direct.MapContext, in *pb.EntrySource_Ancestor) *krm.EntrySource_Ancestor {
	if in == nil {
		return nil
	}
	out := &krm.EntrySource_Ancestor{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Type = direct.LazyPtr(in.GetType())
	return out
}
func EntrySource_Ancestor_ToProto(mapCtx *direct.MapContext, in *krm.EntrySource_Ancestor) *pb.EntrySource_Ancestor {
	if in == nil {
		return nil
	}
	out := &pb.EntrySource_Ancestor{}
	out.Name = direct.ValueOf(in.Name)
	out.Type = direct.ValueOf(in.Type)
	return out
}
