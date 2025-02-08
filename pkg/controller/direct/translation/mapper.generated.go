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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/translate/apiv3/translatepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/translation/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func GlossaryEntry_FromProto(mapCtx *direct.MapContext, in *pb.GlossaryEntry) *krm.GlossaryEntry {
	if in == nil {
		return nil
	}
	out := &krm.GlossaryEntry{}
	out.Name = direct.LazyPtr(in.GetName())
	out.TermsPair = GlossaryEntry_GlossaryTermsPair_FromProto(mapCtx, in.GetTermsPair())
	out.TermsSet = GlossaryEntry_GlossaryTermsSet_FromProto(mapCtx, in.GetTermsSet())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func GlossaryEntry_ToProto(mapCtx *direct.MapContext, in *krm.GlossaryEntry) *pb.GlossaryEntry {
	if in == nil {
		return nil
	}
	out := &pb.GlossaryEntry{}
	out.Name = direct.ValueOf(in.Name)
	if oneof := GlossaryEntry_GlossaryTermsPair_ToProto(mapCtx, in.TermsPair); oneof != nil {
		out.Data = &pb.GlossaryEntry_TermsPair{TermsPair: oneof}
	}
	if oneof := GlossaryEntry_GlossaryTermsSet_ToProto(mapCtx, in.TermsSet); oneof != nil {
		out.Data = &pb.GlossaryEntry_TermsSet{TermsSet: oneof}
	}
	out.Description = direct.ValueOf(in.Description)
	return out
}
func GlossaryEntry_GlossaryTermsPair_FromProto(mapCtx *direct.MapContext, in *pb.GlossaryEntry_GlossaryTermsPair) *krm.GlossaryEntry_GlossaryTermsPair {
	if in == nil {
		return nil
	}
	out := &krm.GlossaryEntry_GlossaryTermsPair{}
	out.SourceTerm = GlossaryTerm_FromProto(mapCtx, in.GetSourceTerm())
	out.TargetTerm = GlossaryTerm_FromProto(mapCtx, in.GetTargetTerm())
	return out
}
func GlossaryEntry_GlossaryTermsPair_ToProto(mapCtx *direct.MapContext, in *krm.GlossaryEntry_GlossaryTermsPair) *pb.GlossaryEntry_GlossaryTermsPair {
	if in == nil {
		return nil
	}
	out := &pb.GlossaryEntry_GlossaryTermsPair{}
	out.SourceTerm = GlossaryTerm_ToProto(mapCtx, in.SourceTerm)
	out.TargetTerm = GlossaryTerm_ToProto(mapCtx, in.TargetTerm)
	return out
}
func GlossaryEntry_GlossaryTermsSet_FromProto(mapCtx *direct.MapContext, in *pb.GlossaryEntry_GlossaryTermsSet) *krm.GlossaryEntry_GlossaryTermsSet {
	if in == nil {
		return nil
	}
	out := &krm.GlossaryEntry_GlossaryTermsSet{}
	out.Terms = direct.Slice_FromProto(mapCtx, in.Terms, GlossaryTerm_FromProto)
	return out
}
func GlossaryEntry_GlossaryTermsSet_ToProto(mapCtx *direct.MapContext, in *krm.GlossaryEntry_GlossaryTermsSet) *pb.GlossaryEntry_GlossaryTermsSet {
	if in == nil {
		return nil
	}
	out := &pb.GlossaryEntry_GlossaryTermsSet{}
	out.Terms = direct.Slice_ToProto(mapCtx, in.Terms, GlossaryTerm_ToProto)
	return out
}
func GlossaryTerm_FromProto(mapCtx *direct.MapContext, in *pb.GlossaryTerm) *krm.GlossaryTerm {
	if in == nil {
		return nil
	}
	out := &krm.GlossaryTerm{}
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	out.Text = direct.LazyPtr(in.GetText())
	return out
}
func GlossaryTerm_ToProto(mapCtx *direct.MapContext, in *krm.GlossaryTerm) *pb.GlossaryTerm {
	if in == nil {
		return nil
	}
	out := &pb.GlossaryTerm{}
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	out.Text = direct.ValueOf(in.Text)
	return out
}
func TranslationGlossaryEntryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GlossaryEntry) *krm.TranslationGlossaryEntryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TranslationGlossaryEntryObservedState{}
	// MISSING: Name
	// MISSING: TermsPair
	// MISSING: TermsSet
	// MISSING: Description
	return out
}
func TranslationGlossaryEntryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TranslationGlossaryEntryObservedState) *pb.GlossaryEntry {
	if in == nil {
		return nil
	}
	out := &pb.GlossaryEntry{}
	// MISSING: Name
	// MISSING: TermsPair
	// MISSING: TermsSet
	// MISSING: Description
	return out
}
func TranslationGlossaryEntrySpec_FromProto(mapCtx *direct.MapContext, in *pb.GlossaryEntry) *krm.TranslationGlossaryEntrySpec {
	if in == nil {
		return nil
	}
	out := &krm.TranslationGlossaryEntrySpec{}
	// MISSING: Name
	// MISSING: TermsPair
	// MISSING: TermsSet
	// MISSING: Description
	return out
}
func TranslationGlossaryEntrySpec_ToProto(mapCtx *direct.MapContext, in *krm.TranslationGlossaryEntrySpec) *pb.GlossaryEntry {
	if in == nil {
		return nil
	}
	out := &pb.GlossaryEntry{}
	// MISSING: Name
	// MISSING: TermsPair
	// MISSING: TermsSet
	// MISSING: Description
	return out
}
