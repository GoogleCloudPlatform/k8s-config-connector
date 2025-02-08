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

package dialogflow

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func DialogflowEntityTypeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EntityType) *krm.DialogflowEntityTypeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowEntityTypeObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Kind
	// MISSING: AutoExpansionMode
	// MISSING: Entities
	// MISSING: ExcludedPhrases
	// MISSING: EnableFuzzyExtraction
	// MISSING: Redact
	return out
}
func DialogflowEntityTypeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowEntityTypeObservedState) *pb.EntityType {
	if in == nil {
		return nil
	}
	out := &pb.EntityType{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Kind
	// MISSING: AutoExpansionMode
	// MISSING: Entities
	// MISSING: ExcludedPhrases
	// MISSING: EnableFuzzyExtraction
	// MISSING: Redact
	return out
}
func DialogflowEntityTypeSpec_FromProto(mapCtx *direct.MapContext, in *pb.EntityType) *krm.DialogflowEntityTypeSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowEntityTypeSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Kind
	// MISSING: AutoExpansionMode
	// MISSING: Entities
	// MISSING: ExcludedPhrases
	// MISSING: EnableFuzzyExtraction
	// MISSING: Redact
	return out
}
func DialogflowEntityTypeSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowEntityTypeSpec) *pb.EntityType {
	if in == nil {
		return nil
	}
	out := &pb.EntityType{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Kind
	// MISSING: AutoExpansionMode
	// MISSING: Entities
	// MISSING: ExcludedPhrases
	// MISSING: EnableFuzzyExtraction
	// MISSING: Redact
	return out
}
func EntityType_FromProto(mapCtx *direct.MapContext, in *pb.EntityType) *krm.EntityType {
	if in == nil {
		return nil
	}
	out := &krm.EntityType{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Kind = direct.Enum_FromProto(mapCtx, in.GetKind())
	out.AutoExpansionMode = direct.Enum_FromProto(mapCtx, in.GetAutoExpansionMode())
	out.Entities = direct.Slice_FromProto(mapCtx, in.Entities, EntityType_Entity_FromProto)
	out.ExcludedPhrases = direct.Slice_FromProto(mapCtx, in.ExcludedPhrases, EntityType_ExcludedPhrase_FromProto)
	out.EnableFuzzyExtraction = direct.LazyPtr(in.GetEnableFuzzyExtraction())
	out.Redact = direct.LazyPtr(in.GetRedact())
	return out
}
func EntityType_ToProto(mapCtx *direct.MapContext, in *krm.EntityType) *pb.EntityType {
	if in == nil {
		return nil
	}
	out := &pb.EntityType{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Kind = direct.Enum_ToProto[pb.EntityType_Kind](mapCtx, in.Kind)
	out.AutoExpansionMode = direct.Enum_ToProto[pb.EntityType_AutoExpansionMode](mapCtx, in.AutoExpansionMode)
	out.Entities = direct.Slice_ToProto(mapCtx, in.Entities, EntityType_Entity_ToProto)
	out.ExcludedPhrases = direct.Slice_ToProto(mapCtx, in.ExcludedPhrases, EntityType_ExcludedPhrase_ToProto)
	out.EnableFuzzyExtraction = direct.ValueOf(in.EnableFuzzyExtraction)
	out.Redact = direct.ValueOf(in.Redact)
	return out
}
func EntityType_Entity_FromProto(mapCtx *direct.MapContext, in *pb.EntityType_Entity) *krm.EntityType_Entity {
	if in == nil {
		return nil
	}
	out := &krm.EntityType_Entity{}
	out.Value = direct.LazyPtr(in.GetValue())
	out.Synonyms = in.Synonyms
	return out
}
func EntityType_Entity_ToProto(mapCtx *direct.MapContext, in *krm.EntityType_Entity) *pb.EntityType_Entity {
	if in == nil {
		return nil
	}
	out := &pb.EntityType_Entity{}
	out.Value = direct.ValueOf(in.Value)
	out.Synonyms = in.Synonyms
	return out
}
func EntityType_ExcludedPhrase_FromProto(mapCtx *direct.MapContext, in *pb.EntityType_ExcludedPhrase) *krm.EntityType_ExcludedPhrase {
	if in == nil {
		return nil
	}
	out := &krm.EntityType_ExcludedPhrase{}
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func EntityType_ExcludedPhrase_ToProto(mapCtx *direct.MapContext, in *krm.EntityType_ExcludedPhrase) *pb.EntityType_ExcludedPhrase {
	if in == nil {
		return nil
	}
	out := &pb.EntityType_ExcludedPhrase{}
	out.Value = direct.ValueOf(in.Value)
	return out
}
