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

package discoveryengine

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Condition_FromProto(mapCtx *direct.MapContext, in *pb.Condition) *krm.Condition {
	if in == nil {
		return nil
	}
	out := &krm.Condition{}
	out.QueryTerms = direct.Slice_FromProto(mapCtx, in.QueryTerms, Condition_QueryTerm_FromProto)
	out.ActiveTimeRange = direct.Slice_FromProto(mapCtx, in.ActiveTimeRange, Condition_TimeRange_FromProto)
	out.QueryRegex = direct.LazyPtr(in.GetQueryRegex())
	return out
}
func Condition_ToProto(mapCtx *direct.MapContext, in *krm.Condition) *pb.Condition {
	if in == nil {
		return nil
	}
	out := &pb.Condition{}
	out.QueryTerms = direct.Slice_ToProto(mapCtx, in.QueryTerms, Condition_QueryTerm_ToProto)
	out.ActiveTimeRange = direct.Slice_ToProto(mapCtx, in.ActiveTimeRange, Condition_TimeRange_ToProto)
	out.QueryRegex = direct.ValueOf(in.QueryRegex)
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
	out.BoostAction = Control_BoostAction_FromProto(mapCtx, in.GetBoostAction())
	out.FilterAction = Control_FilterAction_FromProto(mapCtx, in.GetFilterAction())
	out.RedirectAction = Control_RedirectAction_FromProto(mapCtx, in.GetRedirectAction())
	out.SynonymsAction = Control_SynonymsAction_FromProto(mapCtx, in.GetSynonymsAction())
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: AssociatedServingConfigIds
	out.SolutionType = direct.Enum_FromProto(mapCtx, in.GetSolutionType())
	out.UseCases = direct.EnumSlice_FromProto(mapCtx, in.UseCases)
	out.Conditions = direct.Slice_FromProto(mapCtx, in.Conditions, Condition_FromProto)
	return out
}
func Control_ToProto(mapCtx *direct.MapContext, in *krm.Control) *pb.Control {
	if in == nil {
		return nil
	}
	out := &pb.Control{}
	if oneof := Control_BoostAction_ToProto(mapCtx, in.BoostAction); oneof != nil {
		out.Action = &pb.Control_BoostAction_{BoostAction: oneof}
	}
	if oneof := Control_FilterAction_ToProto(mapCtx, in.FilterAction); oneof != nil {
		out.Action = &pb.Control_FilterAction_{FilterAction: oneof}
	}
	if oneof := Control_RedirectAction_ToProto(mapCtx, in.RedirectAction); oneof != nil {
		out.Action = &pb.Control_RedirectAction_{RedirectAction: oneof}
	}
	if oneof := Control_SynonymsAction_ToProto(mapCtx, in.SynonymsAction); oneof != nil {
		out.Action = &pb.Control_SynonymsAction_{SynonymsAction: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: AssociatedServingConfigIds
	out.SolutionType = direct.Enum_ToProto[pb.SolutionType](mapCtx, in.SolutionType)
	out.UseCases = direct.EnumSlice_ToProto[pb.SearchUseCase](mapCtx, in.UseCases)
	out.Conditions = direct.Slice_ToProto(mapCtx, in.Conditions, Condition_ToProto)
	return out
}
func ControlObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Control) *krm.ControlObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ControlObservedState{}
	// MISSING: BoostAction
	// MISSING: FilterAction
	// MISSING: RedirectAction
	// MISSING: SynonymsAction
	// MISSING: Name
	// MISSING: DisplayName
	out.AssociatedServingConfigIds = in.AssociatedServingConfigIds
	// MISSING: SolutionType
	// MISSING: UseCases
	// MISSING: Conditions
	return out
}
func ControlObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ControlObservedState) *pb.Control {
	if in == nil {
		return nil
	}
	out := &pb.Control{}
	// MISSING: BoostAction
	// MISSING: FilterAction
	// MISSING: RedirectAction
	// MISSING: SynonymsAction
	// MISSING: Name
	// MISSING: DisplayName
	out.AssociatedServingConfigIds = in.AssociatedServingConfigIds
	// MISSING: SolutionType
	// MISSING: UseCases
	// MISSING: Conditions
	return out
}
func Control_BoostAction_FromProto(mapCtx *direct.MapContext, in *pb.Control_BoostAction) *krm.Control_BoostAction {
	if in == nil {
		return nil
	}
	out := &krm.Control_BoostAction{}
	out.Boost = direct.LazyPtr(in.GetBoost())
	out.Filter = direct.LazyPtr(in.GetFilter())
	out.DataStore = direct.LazyPtr(in.GetDataStore())
	return out
}
func Control_BoostAction_ToProto(mapCtx *direct.MapContext, in *krm.Control_BoostAction) *pb.Control_BoostAction {
	if in == nil {
		return nil
	}
	out := &pb.Control_BoostAction{}
	out.Boost = direct.ValueOf(in.Boost)
	out.Filter = direct.ValueOf(in.Filter)
	out.DataStore = direct.ValueOf(in.DataStore)
	return out
}
func Control_FilterAction_FromProto(mapCtx *direct.MapContext, in *pb.Control_FilterAction) *krm.Control_FilterAction {
	if in == nil {
		return nil
	}
	out := &krm.Control_FilterAction{}
	out.Filter = direct.LazyPtr(in.GetFilter())
	out.DataStore = direct.LazyPtr(in.GetDataStore())
	return out
}
func Control_FilterAction_ToProto(mapCtx *direct.MapContext, in *krm.Control_FilterAction) *pb.Control_FilterAction {
	if in == nil {
		return nil
	}
	out := &pb.Control_FilterAction{}
	out.Filter = direct.ValueOf(in.Filter)
	out.DataStore = direct.ValueOf(in.DataStore)
	return out
}
func Control_RedirectAction_FromProto(mapCtx *direct.MapContext, in *pb.Control_RedirectAction) *krm.Control_RedirectAction {
	if in == nil {
		return nil
	}
	out := &krm.Control_RedirectAction{}
	out.RedirectURI = direct.LazyPtr(in.GetRedirectUri())
	return out
}
func Control_RedirectAction_ToProto(mapCtx *direct.MapContext, in *krm.Control_RedirectAction) *pb.Control_RedirectAction {
	if in == nil {
		return nil
	}
	out := &pb.Control_RedirectAction{}
	out.RedirectUri = direct.ValueOf(in.RedirectURI)
	return out
}
func Control_SynonymsAction_FromProto(mapCtx *direct.MapContext, in *pb.Control_SynonymsAction) *krm.Control_SynonymsAction {
	if in == nil {
		return nil
	}
	out := &krm.Control_SynonymsAction{}
	out.Synonyms = in.Synonyms
	return out
}
func Control_SynonymsAction_ToProto(mapCtx *direct.MapContext, in *krm.Control_SynonymsAction) *pb.Control_SynonymsAction {
	if in == nil {
		return nil
	}
	out := &pb.Control_SynonymsAction{}
	out.Synonyms = in.Synonyms
	return out
}
