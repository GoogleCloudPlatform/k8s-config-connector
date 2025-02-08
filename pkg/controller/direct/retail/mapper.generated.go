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

package retail

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/retail/apiv2/retailpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/retail/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func RetailServingConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ServingConfig) *krm.RetailServingConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RetailServingConfigObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ModelID
	// MISSING: PriceRerankingLevel
	// MISSING: FacetControlIds
	// MISSING: DynamicFacetSpec
	// MISSING: BoostControlIds
	// MISSING: FilterControlIds
	// MISSING: RedirectControlIds
	// MISSING: TwowaySynonymsControlIds
	// MISSING: OnewaySynonymsControlIds
	// MISSING: DoNotAssociateControlIds
	// MISSING: ReplacementControlIds
	// MISSING: IgnoreControlIds
	// MISSING: DiversityLevel
	// MISSING: DiversityType
	// MISSING: EnableCategoryFilterLevel
	// MISSING: IgnoreRecsDenylist
	// MISSING: PersonalizationSpec
	// MISSING: SolutionTypes
	return out
}
func RetailServingConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RetailServingConfigObservedState) *pb.ServingConfig {
	if in == nil {
		return nil
	}
	out := &pb.ServingConfig{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ModelID
	// MISSING: PriceRerankingLevel
	// MISSING: FacetControlIds
	// MISSING: DynamicFacetSpec
	// MISSING: BoostControlIds
	// MISSING: FilterControlIds
	// MISSING: RedirectControlIds
	// MISSING: TwowaySynonymsControlIds
	// MISSING: OnewaySynonymsControlIds
	// MISSING: DoNotAssociateControlIds
	// MISSING: ReplacementControlIds
	// MISSING: IgnoreControlIds
	// MISSING: DiversityLevel
	// MISSING: DiversityType
	// MISSING: EnableCategoryFilterLevel
	// MISSING: IgnoreRecsDenylist
	// MISSING: PersonalizationSpec
	// MISSING: SolutionTypes
	return out
}
func RetailServingConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.ServingConfig) *krm.RetailServingConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.RetailServingConfigSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ModelID
	// MISSING: PriceRerankingLevel
	// MISSING: FacetControlIds
	// MISSING: DynamicFacetSpec
	// MISSING: BoostControlIds
	// MISSING: FilterControlIds
	// MISSING: RedirectControlIds
	// MISSING: TwowaySynonymsControlIds
	// MISSING: OnewaySynonymsControlIds
	// MISSING: DoNotAssociateControlIds
	// MISSING: ReplacementControlIds
	// MISSING: IgnoreControlIds
	// MISSING: DiversityLevel
	// MISSING: DiversityType
	// MISSING: EnableCategoryFilterLevel
	// MISSING: IgnoreRecsDenylist
	// MISSING: PersonalizationSpec
	// MISSING: SolutionTypes
	return out
}
func RetailServingConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.RetailServingConfigSpec) *pb.ServingConfig {
	if in == nil {
		return nil
	}
	out := &pb.ServingConfig{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ModelID
	// MISSING: PriceRerankingLevel
	// MISSING: FacetControlIds
	// MISSING: DynamicFacetSpec
	// MISSING: BoostControlIds
	// MISSING: FilterControlIds
	// MISSING: RedirectControlIds
	// MISSING: TwowaySynonymsControlIds
	// MISSING: OnewaySynonymsControlIds
	// MISSING: DoNotAssociateControlIds
	// MISSING: ReplacementControlIds
	// MISSING: IgnoreControlIds
	// MISSING: DiversityLevel
	// MISSING: DiversityType
	// MISSING: EnableCategoryFilterLevel
	// MISSING: IgnoreRecsDenylist
	// MISSING: PersonalizationSpec
	// MISSING: SolutionTypes
	return out
}
func ServingConfig_FromProto(mapCtx *direct.MapContext, in *pb.ServingConfig) *krm.ServingConfig {
	if in == nil {
		return nil
	}
	out := &krm.ServingConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.ModelID = direct.LazyPtr(in.GetModelId())
	out.PriceRerankingLevel = direct.LazyPtr(in.GetPriceRerankingLevel())
	out.FacetControlIds = in.FacetControlIds
	out.DynamicFacetSpec = SearchRequest_DynamicFacetSpec_FromProto(mapCtx, in.GetDynamicFacetSpec())
	out.BoostControlIds = in.BoostControlIds
	out.FilterControlIds = in.FilterControlIds
	out.RedirectControlIds = in.RedirectControlIds
	out.TwowaySynonymsControlIds = in.TwowaySynonymsControlIds
	out.OnewaySynonymsControlIds = in.OnewaySynonymsControlIds
	out.DoNotAssociateControlIds = in.DoNotAssociateControlIds
	out.ReplacementControlIds = in.ReplacementControlIds
	out.IgnoreControlIds = in.IgnoreControlIds
	out.DiversityLevel = direct.LazyPtr(in.GetDiversityLevel())
	out.DiversityType = direct.Enum_FromProto(mapCtx, in.GetDiversityType())
	out.EnableCategoryFilterLevel = direct.LazyPtr(in.GetEnableCategoryFilterLevel())
	out.IgnoreRecsDenylist = direct.LazyPtr(in.GetIgnoreRecsDenylist())
	out.PersonalizationSpec = SearchRequest_PersonalizationSpec_FromProto(mapCtx, in.GetPersonalizationSpec())
	out.SolutionTypes = direct.EnumSlice_FromProto(mapCtx, in.SolutionTypes)
	return out
}
func ServingConfig_ToProto(mapCtx *direct.MapContext, in *krm.ServingConfig) *pb.ServingConfig {
	if in == nil {
		return nil
	}
	out := &pb.ServingConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.ModelId = direct.ValueOf(in.ModelID)
	out.PriceRerankingLevel = direct.ValueOf(in.PriceRerankingLevel)
	out.FacetControlIds = in.FacetControlIds
	out.DynamicFacetSpec = SearchRequest_DynamicFacetSpec_ToProto(mapCtx, in.DynamicFacetSpec)
	out.BoostControlIds = in.BoostControlIds
	out.FilterControlIds = in.FilterControlIds
	out.RedirectControlIds = in.RedirectControlIds
	out.TwowaySynonymsControlIds = in.TwowaySynonymsControlIds
	out.OnewaySynonymsControlIds = in.OnewaySynonymsControlIds
	out.DoNotAssociateControlIds = in.DoNotAssociateControlIds
	out.ReplacementControlIds = in.ReplacementControlIds
	out.IgnoreControlIds = in.IgnoreControlIds
	out.DiversityLevel = direct.ValueOf(in.DiversityLevel)
	out.DiversityType = direct.Enum_ToProto[pb.ServingConfig_DiversityType](mapCtx, in.DiversityType)
	out.EnableCategoryFilterLevel = direct.ValueOf(in.EnableCategoryFilterLevel)
	out.IgnoreRecsDenylist = direct.ValueOf(in.IgnoreRecsDenylist)
	out.PersonalizationSpec = SearchRequest_PersonalizationSpec_ToProto(mapCtx, in.PersonalizationSpec)
	out.SolutionTypes = direct.EnumSlice_ToProto[pb.SolutionType](mapCtx, in.SolutionTypes)
	return out
}
