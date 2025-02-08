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
	pb "cloud.google.com/go/retail/apiv2beta/retailpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/retail/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Condition_FromProto(mapCtx *direct.MapContext, in *pb.Condition) *krm.Condition {
	if in == nil {
		return nil
	}
	out := &krm.Condition{}
	out.QueryTerms = direct.Slice_FromProto(mapCtx, in.QueryTerms, Condition_QueryTerm_FromProto)
	out.ActiveTimeRange = direct.Slice_FromProto(mapCtx, in.ActiveTimeRange, Condition_TimeRange_FromProto)
	out.PageCategories = in.PageCategories
	return out
}
func Condition_ToProto(mapCtx *direct.MapContext, in *krm.Condition) *pb.Condition {
	if in == nil {
		return nil
	}
	out := &pb.Condition{}
	out.QueryTerms = direct.Slice_ToProto(mapCtx, in.QueryTerms, Condition_QueryTerm_ToProto)
	out.ActiveTimeRange = direct.Slice_ToProto(mapCtx, in.ActiveTimeRange, Condition_TimeRange_ToProto)
	out.PageCategories = in.PageCategories
	return out
}
func Condition_QueryTerm_FromProto(mapCtx *direct.MapContext, in *pb.Condition_QueryTerm) *krm.Condition_QueryTerm {
	if in == nil {
		return nil
	}
	out := &krm.Condition_QueryTerm{}
	out.Value = direct.LazyPtr(in.GetValue())
	out.FullMatch = direct.LazyPtr(in.GetFullMatch())
	return out
}
func Condition_QueryTerm_ToProto(mapCtx *direct.MapContext, in *krm.Condition_QueryTerm) *pb.Condition_QueryTerm {
	if in == nil {
		return nil
	}
	out := &pb.Condition_QueryTerm{}
	out.Value = direct.ValueOf(in.Value)
	out.FullMatch = direct.ValueOf(in.FullMatch)
	return out
}
func Condition_TimeRange_FromProto(mapCtx *direct.MapContext, in *pb.Condition_TimeRange) *krm.Condition_TimeRange {
	if in == nil {
		return nil
	}
	out := &krm.Condition_TimeRange{}
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}
func Condition_TimeRange_ToProto(mapCtx *direct.MapContext, in *krm.Condition_TimeRange) *pb.Condition_TimeRange {
	if in == nil {
		return nil
	}
	out := &pb.Condition_TimeRange{}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	return out
}
func Control_FromProto(mapCtx *direct.MapContext, in *pb.Control) *krm.Control {
	if in == nil {
		return nil
	}
	out := &krm.Control{}
	out.FacetSpec = SearchRequest_FacetSpec_FromProto(mapCtx, in.GetFacetSpec())
	out.Rule = Rule_FromProto(mapCtx, in.GetRule())
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: AssociatedServingConfigIds
	out.SolutionTypes = direct.EnumSlice_FromProto(mapCtx, in.SolutionTypes)
	out.SearchSolutionUseCase = direct.EnumSlice_FromProto(mapCtx, in.SearchSolutionUseCase)
	return out
}
func Control_ToProto(mapCtx *direct.MapContext, in *krm.Control) *pb.Control {
	if in == nil {
		return nil
	}
	out := &pb.Control{}
	if oneof := SearchRequest_FacetSpec_ToProto(mapCtx, in.FacetSpec); oneof != nil {
		out.Control = &pb.Control_FacetSpec{FacetSpec: oneof}
	}
	if oneof := Rule_ToProto(mapCtx, in.Rule); oneof != nil {
		out.Control = &pb.Control_Rule{Rule: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: AssociatedServingConfigIds
	out.SolutionTypes = direct.EnumSlice_ToProto[pb.SolutionType](mapCtx, in.SolutionTypes)
	out.SearchSolutionUseCase = direct.EnumSlice_ToProto[pb.SearchSolutionUseCase](mapCtx, in.SearchSolutionUseCase)
	return out
}
func ControlObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Control) *krm.ControlObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ControlObservedState{}
	// MISSING: FacetSpec
	// MISSING: Rule
	// MISSING: Name
	// MISSING: DisplayName
	out.AssociatedServingConfigIds = in.AssociatedServingConfigIds
	// MISSING: SolutionTypes
	// MISSING: SearchSolutionUseCase
	return out
}
func ControlObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ControlObservedState) *pb.Control {
	if in == nil {
		return nil
	}
	out := &pb.Control{}
	// MISSING: FacetSpec
	// MISSING: Rule
	// MISSING: Name
	// MISSING: DisplayName
	out.AssociatedServingConfigIds = in.AssociatedServingConfigIds
	// MISSING: SolutionTypes
	// MISSING: SearchSolutionUseCase
	return out
}
func Interval_FromProto(mapCtx *direct.MapContext, in *pb.Interval) *krm.Interval {
	if in == nil {
		return nil
	}
	out := &krm.Interval{}
	out.Minimum = direct.LazyPtr(in.GetMinimum())
	out.ExclusiveMinimum = direct.LazyPtr(in.GetExclusiveMinimum())
	out.Maximum = direct.LazyPtr(in.GetMaximum())
	out.ExclusiveMaximum = direct.LazyPtr(in.GetExclusiveMaximum())
	return out
}
func Interval_ToProto(mapCtx *direct.MapContext, in *krm.Interval) *pb.Interval {
	if in == nil {
		return nil
	}
	out := &pb.Interval{}
	if oneof := Interval_Minimum_ToProto(mapCtx, in.Minimum); oneof != nil {
		out.Min = oneof
	}
	if oneof := Interval_ExclusiveMinimum_ToProto(mapCtx, in.ExclusiveMinimum); oneof != nil {
		out.Min = oneof
	}
	if oneof := Interval_Maximum_ToProto(mapCtx, in.Maximum); oneof != nil {
		out.Max = oneof
	}
	if oneof := Interval_ExclusiveMaximum_ToProto(mapCtx, in.ExclusiveMaximum); oneof != nil {
		out.Max = oneof
	}
	return out
}
func Rule_FromProto(mapCtx *direct.MapContext, in *pb.Rule) *krm.Rule {
	if in == nil {
		return nil
	}
	out := &krm.Rule{}
	out.BoostAction = Rule_BoostAction_FromProto(mapCtx, in.GetBoostAction())
	out.RedirectAction = Rule_RedirectAction_FromProto(mapCtx, in.GetRedirectAction())
	out.OnewaySynonymsAction = Rule_OnewaySynonymsAction_FromProto(mapCtx, in.GetOnewaySynonymsAction())
	out.DoNotAssociateAction = Rule_DoNotAssociateAction_FromProto(mapCtx, in.GetDoNotAssociateAction())
	out.ReplacementAction = Rule_ReplacementAction_FromProto(mapCtx, in.GetReplacementAction())
	out.IgnoreAction = Rule_IgnoreAction_FromProto(mapCtx, in.GetIgnoreAction())
	out.FilterAction = Rule_FilterAction_FromProto(mapCtx, in.GetFilterAction())
	out.TwowaySynonymsAction = Rule_TwowaySynonymsAction_FromProto(mapCtx, in.GetTwowaySynonymsAction())
	out.ForceReturnFacetAction = Rule_ForceReturnFacetAction_FromProto(mapCtx, in.GetForceReturnFacetAction())
	out.RemoveFacetAction = Rule_RemoveFacetAction_FromProto(mapCtx, in.GetRemoveFacetAction())
	out.Condition = Condition_FromProto(mapCtx, in.GetCondition())
	return out
}
func Rule_ToProto(mapCtx *direct.MapContext, in *krm.Rule) *pb.Rule {
	if in == nil {
		return nil
	}
	out := &pb.Rule{}
	if oneof := Rule_BoostAction_ToProto(mapCtx, in.BoostAction); oneof != nil {
		out.Action = &pb.Rule_BoostAction_{BoostAction: oneof}
	}
	if oneof := Rule_RedirectAction_ToProto(mapCtx, in.RedirectAction); oneof != nil {
		out.Action = &pb.Rule_RedirectAction_{RedirectAction: oneof}
	}
	if oneof := Rule_OnewaySynonymsAction_ToProto(mapCtx, in.OnewaySynonymsAction); oneof != nil {
		out.Action = &pb.Rule_OnewaySynonymsAction_{OnewaySynonymsAction: oneof}
	}
	if oneof := Rule_DoNotAssociateAction_ToProto(mapCtx, in.DoNotAssociateAction); oneof != nil {
		out.Action = &pb.Rule_DoNotAssociateAction_{DoNotAssociateAction: oneof}
	}
	if oneof := Rule_ReplacementAction_ToProto(mapCtx, in.ReplacementAction); oneof != nil {
		out.Action = &pb.Rule_ReplacementAction_{ReplacementAction: oneof}
	}
	if oneof := Rule_IgnoreAction_ToProto(mapCtx, in.IgnoreAction); oneof != nil {
		out.Action = &pb.Rule_IgnoreAction_{IgnoreAction: oneof}
	}
	if oneof := Rule_FilterAction_ToProto(mapCtx, in.FilterAction); oneof != nil {
		out.Action = &pb.Rule_FilterAction_{FilterAction: oneof}
	}
	if oneof := Rule_TwowaySynonymsAction_ToProto(mapCtx, in.TwowaySynonymsAction); oneof != nil {
		out.Action = &pb.Rule_TwowaySynonymsAction_{TwowaySynonymsAction: oneof}
	}
	if oneof := Rule_ForceReturnFacetAction_ToProto(mapCtx, in.ForceReturnFacetAction); oneof != nil {
		out.Action = &pb.Rule_ForceReturnFacetAction_{ForceReturnFacetAction: oneof}
	}
	if oneof := Rule_RemoveFacetAction_ToProto(mapCtx, in.RemoveFacetAction); oneof != nil {
		out.Action = &pb.Rule_RemoveFacetAction_{RemoveFacetAction: oneof}
	}
	out.Condition = Condition_ToProto(mapCtx, in.Condition)
	return out
}
func Rule_BoostAction_FromProto(mapCtx *direct.MapContext, in *pb.Rule_BoostAction) *krm.Rule_BoostAction {
	if in == nil {
		return nil
	}
	out := &krm.Rule_BoostAction{}
	out.Boost = direct.LazyPtr(in.GetBoost())
	out.ProductsFilter = direct.LazyPtr(in.GetProductsFilter())
	return out
}
func Rule_BoostAction_ToProto(mapCtx *direct.MapContext, in *krm.Rule_BoostAction) *pb.Rule_BoostAction {
	if in == nil {
		return nil
	}
	out := &pb.Rule_BoostAction{}
	out.Boost = direct.ValueOf(in.Boost)
	out.ProductsFilter = direct.ValueOf(in.ProductsFilter)
	return out
}
func Rule_DoNotAssociateAction_FromProto(mapCtx *direct.MapContext, in *pb.Rule_DoNotAssociateAction) *krm.Rule_DoNotAssociateAction {
	if in == nil {
		return nil
	}
	out := &krm.Rule_DoNotAssociateAction{}
	out.QueryTerms = in.QueryTerms
	out.DoNotAssociateTerms = in.DoNotAssociateTerms
	out.Terms = in.Terms
	return out
}
func Rule_DoNotAssociateAction_ToProto(mapCtx *direct.MapContext, in *krm.Rule_DoNotAssociateAction) *pb.Rule_DoNotAssociateAction {
	if in == nil {
		return nil
	}
	out := &pb.Rule_DoNotAssociateAction{}
	out.QueryTerms = in.QueryTerms
	out.DoNotAssociateTerms = in.DoNotAssociateTerms
	out.Terms = in.Terms
	return out
}
func Rule_FilterAction_FromProto(mapCtx *direct.MapContext, in *pb.Rule_FilterAction) *krm.Rule_FilterAction {
	if in == nil {
		return nil
	}
	out := &krm.Rule_FilterAction{}
	out.Filter = direct.LazyPtr(in.GetFilter())
	return out
}
func Rule_FilterAction_ToProto(mapCtx *direct.MapContext, in *krm.Rule_FilterAction) *pb.Rule_FilterAction {
	if in == nil {
		return nil
	}
	out := &pb.Rule_FilterAction{}
	out.Filter = direct.ValueOf(in.Filter)
	return out
}
func Rule_ForceReturnFacetAction_FromProto(mapCtx *direct.MapContext, in *pb.Rule_ForceReturnFacetAction) *krm.Rule_ForceReturnFacetAction {
	if in == nil {
		return nil
	}
	out := &krm.Rule_ForceReturnFacetAction{}
	out.FacetPositionAdjustments = direct.Slice_FromProto(mapCtx, in.FacetPositionAdjustments, Rule_ForceReturnFacetAction_FacetPositionAdjustment_FromProto)
	return out
}
func Rule_ForceReturnFacetAction_ToProto(mapCtx *direct.MapContext, in *krm.Rule_ForceReturnFacetAction) *pb.Rule_ForceReturnFacetAction {
	if in == nil {
		return nil
	}
	out := &pb.Rule_ForceReturnFacetAction{}
	out.FacetPositionAdjustments = direct.Slice_ToProto(mapCtx, in.FacetPositionAdjustments, Rule_ForceReturnFacetAction_FacetPositionAdjustment_ToProto)
	return out
}
func Rule_ForceReturnFacetAction_FacetPositionAdjustment_FromProto(mapCtx *direct.MapContext, in *pb.Rule_ForceReturnFacetAction_FacetPositionAdjustment) *krm.Rule_ForceReturnFacetAction_FacetPositionAdjustment {
	if in == nil {
		return nil
	}
	out := &krm.Rule_ForceReturnFacetAction_FacetPositionAdjustment{}
	out.AttributeName = direct.LazyPtr(in.GetAttributeName())
	out.Position = direct.LazyPtr(in.GetPosition())
	return out
}
func Rule_ForceReturnFacetAction_FacetPositionAdjustment_ToProto(mapCtx *direct.MapContext, in *krm.Rule_ForceReturnFacetAction_FacetPositionAdjustment) *pb.Rule_ForceReturnFacetAction_FacetPositionAdjustment {
	if in == nil {
		return nil
	}
	out := &pb.Rule_ForceReturnFacetAction_FacetPositionAdjustment{}
	out.AttributeName = direct.ValueOf(in.AttributeName)
	out.Position = direct.ValueOf(in.Position)
	return out
}
func Rule_IgnoreAction_FromProto(mapCtx *direct.MapContext, in *pb.Rule_IgnoreAction) *krm.Rule_IgnoreAction {
	if in == nil {
		return nil
	}
	out := &krm.Rule_IgnoreAction{}
	out.IgnoreTerms = in.IgnoreTerms
	return out
}
func Rule_IgnoreAction_ToProto(mapCtx *direct.MapContext, in *krm.Rule_IgnoreAction) *pb.Rule_IgnoreAction {
	if in == nil {
		return nil
	}
	out := &pb.Rule_IgnoreAction{}
	out.IgnoreTerms = in.IgnoreTerms
	return out
}
func Rule_OnewaySynonymsAction_FromProto(mapCtx *direct.MapContext, in *pb.Rule_OnewaySynonymsAction) *krm.Rule_OnewaySynonymsAction {
	if in == nil {
		return nil
	}
	out := &krm.Rule_OnewaySynonymsAction{}
	out.QueryTerms = in.QueryTerms
	out.Synonyms = in.Synonyms
	out.OnewayTerms = in.OnewayTerms
	return out
}
func Rule_OnewaySynonymsAction_ToProto(mapCtx *direct.MapContext, in *krm.Rule_OnewaySynonymsAction) *pb.Rule_OnewaySynonymsAction {
	if in == nil {
		return nil
	}
	out := &pb.Rule_OnewaySynonymsAction{}
	out.QueryTerms = in.QueryTerms
	out.Synonyms = in.Synonyms
	out.OnewayTerms = in.OnewayTerms
	return out
}
func Rule_RedirectAction_FromProto(mapCtx *direct.MapContext, in *pb.Rule_RedirectAction) *krm.Rule_RedirectAction {
	if in == nil {
		return nil
	}
	out := &krm.Rule_RedirectAction{}
	out.RedirectURI = direct.LazyPtr(in.GetRedirectUri())
	return out
}
func Rule_RedirectAction_ToProto(mapCtx *direct.MapContext, in *krm.Rule_RedirectAction) *pb.Rule_RedirectAction {
	if in == nil {
		return nil
	}
	out := &pb.Rule_RedirectAction{}
	out.RedirectUri = direct.ValueOf(in.RedirectURI)
	return out
}
func Rule_RemoveFacetAction_FromProto(mapCtx *direct.MapContext, in *pb.Rule_RemoveFacetAction) *krm.Rule_RemoveFacetAction {
	if in == nil {
		return nil
	}
	out := &krm.Rule_RemoveFacetAction{}
	out.AttributeNames = in.AttributeNames
	return out
}
func Rule_RemoveFacetAction_ToProto(mapCtx *direct.MapContext, in *krm.Rule_RemoveFacetAction) *pb.Rule_RemoveFacetAction {
	if in == nil {
		return nil
	}
	out := &pb.Rule_RemoveFacetAction{}
	out.AttributeNames = in.AttributeNames
	return out
}
func Rule_ReplacementAction_FromProto(mapCtx *direct.MapContext, in *pb.Rule_ReplacementAction) *krm.Rule_ReplacementAction {
	if in == nil {
		return nil
	}
	out := &krm.Rule_ReplacementAction{}
	out.QueryTerms = in.QueryTerms
	out.ReplacementTerm = direct.LazyPtr(in.GetReplacementTerm())
	out.Term = direct.LazyPtr(in.GetTerm())
	return out
}
func Rule_ReplacementAction_ToProto(mapCtx *direct.MapContext, in *krm.Rule_ReplacementAction) *pb.Rule_ReplacementAction {
	if in == nil {
		return nil
	}
	out := &pb.Rule_ReplacementAction{}
	out.QueryTerms = in.QueryTerms
	out.ReplacementTerm = direct.ValueOf(in.ReplacementTerm)
	out.Term = direct.ValueOf(in.Term)
	return out
}
func Rule_TwowaySynonymsAction_FromProto(mapCtx *direct.MapContext, in *pb.Rule_TwowaySynonymsAction) *krm.Rule_TwowaySynonymsAction {
	if in == nil {
		return nil
	}
	out := &krm.Rule_TwowaySynonymsAction{}
	out.Synonyms = in.Synonyms
	return out
}
func Rule_TwowaySynonymsAction_ToProto(mapCtx *direct.MapContext, in *krm.Rule_TwowaySynonymsAction) *pb.Rule_TwowaySynonymsAction {
	if in == nil {
		return nil
	}
	out := &pb.Rule_TwowaySynonymsAction{}
	out.Synonyms = in.Synonyms
	return out
}
