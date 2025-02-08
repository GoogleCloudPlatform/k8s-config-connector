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

package visionai

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/visionai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/visionai/apiv1/visionaipb"
)
func Corpus_FromProto(mapCtx *direct.MapContext, in *pb.Corpus) *krm.Corpus {
	if in == nil {
		return nil
	}
	out := &krm.Corpus{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DefaultTtl = direct.StringDuration_FromProto(mapCtx, in.GetDefaultTtl())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.SearchCapabilitySetting = SearchCapabilitySetting_FromProto(mapCtx, in.GetSearchCapabilitySetting())
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func Corpus_ToProto(mapCtx *direct.MapContext, in *krm.Corpus) *pb.Corpus {
	if in == nil {
		return nil
	}
	out := &pb.Corpus{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.DefaultTtl = direct.StringDuration_ToProto(mapCtx, in.DefaultTtl)
	out.Type = direct.Enum_ToProto[pb.Corpus_Type](mapCtx, in.Type)
	out.SearchCapabilitySetting = SearchCapabilitySetting_ToProto(mapCtx, in.SearchCapabilitySetting)
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func CorpusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Corpus) *krm.CorpusObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CorpusObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DefaultTtl
	// MISSING: Type
	// MISSING: SearchCapabilitySetting
	out.SatisfiesPzs = in.SatisfiesPzs
	out.SatisfiesPzi = in.SatisfiesPzi
	return out
}
func CorpusObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CorpusObservedState) *pb.Corpus {
	if in == nil {
		return nil
	}
	out := &pb.Corpus{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DefaultTtl
	// MISSING: Type
	// MISSING: SearchCapabilitySetting
	out.SatisfiesPzs = in.SatisfiesPzs
	out.SatisfiesPzi = in.SatisfiesPzi
	return out
}
func SearchCapability_FromProto(mapCtx *direct.MapContext, in *pb.SearchCapability) *krm.SearchCapability {
	if in == nil {
		return nil
	}
	out := &krm.SearchCapability{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func SearchCapability_ToProto(mapCtx *direct.MapContext, in *krm.SearchCapability) *pb.SearchCapability {
	if in == nil {
		return nil
	}
	out := &pb.SearchCapability{}
	out.Type = direct.Enum_ToProto[pb.SearchCapability_Type](mapCtx, in.Type)
	return out
}
func SearchCapabilitySetting_FromProto(mapCtx *direct.MapContext, in *pb.SearchCapabilitySetting) *krm.SearchCapabilitySetting {
	if in == nil {
		return nil
	}
	out := &krm.SearchCapabilitySetting{}
	out.SearchCapabilities = direct.Slice_FromProto(mapCtx, in.SearchCapabilities, SearchCapability_FromProto)
	return out
}
func SearchCapabilitySetting_ToProto(mapCtx *direct.MapContext, in *krm.SearchCapabilitySetting) *pb.SearchCapabilitySetting {
	if in == nil {
		return nil
	}
	out := &pb.SearchCapabilitySetting{}
	out.SearchCapabilities = direct.Slice_ToProto(mapCtx, in.SearchCapabilities, SearchCapability_ToProto)
	return out
}
func VisionaiCorpusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Corpus) *krm.VisionaiCorpusObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiCorpusObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DefaultTtl
	// MISSING: Type
	// MISSING: SearchCapabilitySetting
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func VisionaiCorpusObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiCorpusObservedState) *pb.Corpus {
	if in == nil {
		return nil
	}
	out := &pb.Corpus{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DefaultTtl
	// MISSING: Type
	// MISSING: SearchCapabilitySetting
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func VisionaiCorpusSpec_FromProto(mapCtx *direct.MapContext, in *pb.Corpus) *krm.VisionaiCorpusSpec {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiCorpusSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DefaultTtl
	// MISSING: Type
	// MISSING: SearchCapabilitySetting
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func VisionaiCorpusSpec_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiCorpusSpec) *pb.Corpus {
	if in == nil {
		return nil
	}
	out := &pb.Corpus{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DefaultTtl
	// MISSING: Type
	// MISSING: SearchCapabilitySetting
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
