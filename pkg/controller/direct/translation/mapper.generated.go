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

package translation

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/translate/apiv3/translatepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/translation/v1alpha1"
)
func GcsSource_FromProto(mapCtx *direct.MapContext, in *pb.GcsSource) *krm.GcsSource {
	if in == nil {
		return nil
	}
	out := &krm.GcsSource{}
	out.InputURI = direct.LazyPtr(in.GetInputUri())
	return out
}
func GcsSource_ToProto(mapCtx *direct.MapContext, in *krm.GcsSource) *pb.GcsSource {
	if in == nil {
		return nil
	}
	out := &pb.GcsSource{}
	out.InputUri = direct.ValueOf(in.InputURI)
	return out
}
func Glossary_FromProto(mapCtx *direct.MapContext, in *pb.Glossary) *krm.Glossary {
	if in == nil {
		return nil
	}
	out := &krm.Glossary{}
	out.Name = direct.LazyPtr(in.GetName())
	out.LanguagePair = Glossary_LanguageCodePair_FromProto(mapCtx, in.GetLanguagePair())
	out.LanguageCodesSet = Glossary_LanguageCodesSet_FromProto(mapCtx, in.GetLanguageCodesSet())
	out.InputConfig = GlossaryInputConfig_FromProto(mapCtx, in.GetInputConfig())
	// MISSING: EntryCount
	// MISSING: SubmitTime
	// MISSING: EndTime
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}
func Glossary_ToProto(mapCtx *direct.MapContext, in *krm.Glossary) *pb.Glossary {
	if in == nil {
		return nil
	}
	out := &pb.Glossary{}
	out.Name = direct.ValueOf(in.Name)
	if oneof := Glossary_LanguageCodePair_ToProto(mapCtx, in.LanguagePair); oneof != nil {
		out.Languages = &pb.Glossary_LanguagePair{LanguagePair: oneof}
	}
	if oneof := Glossary_LanguageCodesSet_ToProto(mapCtx, in.LanguageCodesSet); oneof != nil {
		out.Languages = &pb.Glossary_LanguageCodesSet_{LanguageCodesSet: oneof}
	}
	out.InputConfig = GlossaryInputConfig_ToProto(mapCtx, in.InputConfig)
	// MISSING: EntryCount
	// MISSING: SubmitTime
	// MISSING: EndTime
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}
func GlossaryInputConfig_FromProto(mapCtx *direct.MapContext, in *pb.GlossaryInputConfig) *krm.GlossaryInputConfig {
	if in == nil {
		return nil
	}
	out := &krm.GlossaryInputConfig{}
	out.GcsSource = GcsSource_FromProto(mapCtx, in.GetGcsSource())
	return out
}
func GlossaryInputConfig_ToProto(mapCtx *direct.MapContext, in *krm.GlossaryInputConfig) *pb.GlossaryInputConfig {
	if in == nil {
		return nil
	}
	out := &pb.GlossaryInputConfig{}
	if oneof := GcsSource_ToProto(mapCtx, in.GcsSource); oneof != nil {
		out.Source = &pb.GlossaryInputConfig_GcsSource{GcsSource: oneof}
	}
	return out
}
func GlossaryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Glossary) *krm.GlossaryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GlossaryObservedState{}
	// MISSING: Name
	// MISSING: LanguagePair
	// MISSING: LanguageCodesSet
	// MISSING: InputConfig
	out.EntryCount = direct.LazyPtr(in.GetEntryCount())
	out.SubmitTime = direct.StringTimestamp_FromProto(mapCtx, in.GetSubmitTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	// MISSING: DisplayName
	return out
}
func GlossaryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GlossaryObservedState) *pb.Glossary {
	if in == nil {
		return nil
	}
	out := &pb.Glossary{}
	// MISSING: Name
	// MISSING: LanguagePair
	// MISSING: LanguageCodesSet
	// MISSING: InputConfig
	out.EntryCount = direct.ValueOf(in.EntryCount)
	out.SubmitTime = direct.StringTimestamp_ToProto(mapCtx, in.SubmitTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	// MISSING: DisplayName
	return out
}
func Glossary_LanguageCodePair_FromProto(mapCtx *direct.MapContext, in *pb.Glossary_LanguageCodePair) *krm.Glossary_LanguageCodePair {
	if in == nil {
		return nil
	}
	out := &krm.Glossary_LanguageCodePair{}
	out.SourceLanguageCode = direct.LazyPtr(in.GetSourceLanguageCode())
	out.TargetLanguageCode = direct.LazyPtr(in.GetTargetLanguageCode())
	return out
}
func Glossary_LanguageCodePair_ToProto(mapCtx *direct.MapContext, in *krm.Glossary_LanguageCodePair) *pb.Glossary_LanguageCodePair {
	if in == nil {
		return nil
	}
	out := &pb.Glossary_LanguageCodePair{}
	out.SourceLanguageCode = direct.ValueOf(in.SourceLanguageCode)
	out.TargetLanguageCode = direct.ValueOf(in.TargetLanguageCode)
	return out
}
func Glossary_LanguageCodesSet_FromProto(mapCtx *direct.MapContext, in *pb.Glossary_LanguageCodesSet) *krm.Glossary_LanguageCodesSet {
	if in == nil {
		return nil
	}
	out := &krm.Glossary_LanguageCodesSet{}
	out.LanguageCodes = in.LanguageCodes
	return out
}
func Glossary_LanguageCodesSet_ToProto(mapCtx *direct.MapContext, in *krm.Glossary_LanguageCodesSet) *pb.Glossary_LanguageCodesSet {
	if in == nil {
		return nil
	}
	out := &pb.Glossary_LanguageCodesSet{}
	out.LanguageCodes = in.LanguageCodes
	return out
}
func TranslationGlossaryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Glossary) *krm.TranslationGlossaryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TranslationGlossaryObservedState{}
	// MISSING: Name
	// MISSING: LanguagePair
	// MISSING: LanguageCodesSet
	// MISSING: InputConfig
	// MISSING: EntryCount
	// MISSING: SubmitTime
	// MISSING: EndTime
	// MISSING: DisplayName
	return out
}
func TranslationGlossaryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TranslationGlossaryObservedState) *pb.Glossary {
	if in == nil {
		return nil
	}
	out := &pb.Glossary{}
	// MISSING: Name
	// MISSING: LanguagePair
	// MISSING: LanguageCodesSet
	// MISSING: InputConfig
	// MISSING: EntryCount
	// MISSING: SubmitTime
	// MISSING: EndTime
	// MISSING: DisplayName
	return out
}
func TranslationGlossarySpec_FromProto(mapCtx *direct.MapContext, in *pb.Glossary) *krm.TranslationGlossarySpec {
	if in == nil {
		return nil
	}
	out := &krm.TranslationGlossarySpec{}
	// MISSING: Name
	// MISSING: LanguagePair
	// MISSING: LanguageCodesSet
	// MISSING: InputConfig
	// MISSING: EntryCount
	// MISSING: SubmitTime
	// MISSING: EndTime
	// MISSING: DisplayName
	return out
}
func TranslationGlossarySpec_ToProto(mapCtx *direct.MapContext, in *krm.TranslationGlossarySpec) *pb.Glossary {
	if in == nil {
		return nil
	}
	out := &pb.Glossary{}
	// MISSING: Name
	// MISSING: LanguagePair
	// MISSING: LanguageCodesSet
	// MISSING: InputConfig
	// MISSING: EntryCount
	// MISSING: SubmitTime
	// MISSING: EndTime
	// MISSING: DisplayName
	return out
}
