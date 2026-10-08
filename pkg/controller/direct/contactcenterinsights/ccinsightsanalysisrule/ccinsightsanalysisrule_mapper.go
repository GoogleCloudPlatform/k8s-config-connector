// Copyright 2026 Google LLC
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

package ccinsightsanalysisrule

import (
	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contactcenterinsights/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CCInsightsAnalysisRuleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AnalysisRule) *krm.CCInsightsAnalysisRuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CCInsightsAnalysisRuleObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func CCInsightsAnalysisRuleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CCInsightsAnalysisRuleObservedState) *pb.AnalysisRule {
	if in == nil {
		return nil
	}
	out := &pb.AnalysisRule{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}

func CCInsightsAnalysisRuleSpec_FromProto(mapCtx *direct.MapContext, in *pb.AnalysisRule) *krm.CCInsightsAnalysisRuleSpec {
	if in == nil {
		return nil
	}
	out := &krm.CCInsightsAnalysisRuleSpec{}
	// MISSING: Name
	out.DisplayName = in.DisplayName
	out.ConversationFilter = direct.LazyPtr(in.GetConversationFilter())
	out.AnnotatorSelector = AnnotatorSelector_FromProto(mapCtx, in.GetAnnotatorSelector())
	out.AnalysisPercentage = direct.LazyPtr(in.GetAnalysisPercentage())
	out.Active = direct.LazyPtr(in.GetActive())
	return out
}

func CCInsightsAnalysisRuleSpec_ToProto(mapCtx *direct.MapContext, in *krm.CCInsightsAnalysisRuleSpec) *pb.AnalysisRule {
	if in == nil {
		return nil
	}
	out := &pb.AnalysisRule{}
	// MISSING: Name
	out.DisplayName = in.DisplayName
	out.ConversationFilter = direct.ValueOf(in.ConversationFilter)
	out.AnnotatorSelector = AnnotatorSelector_ToProto(mapCtx, in.AnnotatorSelector)
	out.AnalysisPercentage = direct.ValueOf(in.AnalysisPercentage)
	out.Active = direct.ValueOf(in.Active)
	return out
}

func AnnotatorSelector_FromProto(mapCtx *direct.MapContext, in *pb.AnnotatorSelector) *krm.AnnotatorSelector {
	if in == nil {
		return nil
	}
	out := &krm.AnnotatorSelector{}
	out.RunInterruptionAnnotator = direct.LazyPtr(in.GetRunInterruptionAnnotator())
	out.RunSilenceAnnotator = direct.LazyPtr(in.GetRunSilenceAnnotator())
	out.RunPhraseMatcherAnnotator = direct.LazyPtr(in.GetRunPhraseMatcherAnnotator())
	out.PhraseMatchers = in.PhraseMatchers
	out.RunSentimentAnnotator = direct.LazyPtr(in.GetRunSentimentAnnotator())
	out.RunEntityAnnotator = direct.LazyPtr(in.GetRunEntityAnnotator())
	out.RunIntentAnnotator = direct.LazyPtr(in.GetRunIntentAnnotator())
	out.RunIssueModelAnnotator = direct.LazyPtr(in.GetRunIssueModelAnnotator())
	out.IssueModels = in.IssueModels
	out.RunSummarizationAnnotator = direct.LazyPtr(in.GetRunSummarizationAnnotator())
	out.SummarizationConfig = AnnotatorSelector_SummarizationConfig_FromProto(mapCtx, in.GetSummarizationConfig())
	out.RunQaAnnotator = direct.LazyPtr(in.GetRunQaAnnotator())
	out.QaConfig = AnnotatorSelector_QaConfig_FromProto(mapCtx, in.GetQaConfig())
	return out
}

func AnnotatorSelector_ToProto(mapCtx *direct.MapContext, in *krm.AnnotatorSelector) *pb.AnnotatorSelector {
	if in == nil {
		return nil
	}
	out := &pb.AnnotatorSelector{}
	out.RunInterruptionAnnotator = direct.ValueOf(in.RunInterruptionAnnotator)
	out.RunSilenceAnnotator = direct.ValueOf(in.RunSilenceAnnotator)
	out.RunPhraseMatcherAnnotator = direct.ValueOf(in.RunPhraseMatcherAnnotator)
	out.PhraseMatchers = in.PhraseMatchers
	out.RunSentimentAnnotator = direct.ValueOf(in.RunSentimentAnnotator)
	out.RunEntityAnnotator = direct.ValueOf(in.RunEntityAnnotator)
	out.RunIntentAnnotator = direct.ValueOf(in.RunIntentAnnotator)
	out.RunIssueModelAnnotator = direct.ValueOf(in.RunIssueModelAnnotator)
	out.IssueModels = in.IssueModels
	out.RunSummarizationAnnotator = direct.ValueOf(in.RunSummarizationAnnotator)
	out.SummarizationConfig = AnnotatorSelector_SummarizationConfig_ToProto(mapCtx, in.SummarizationConfig)
	out.RunQaAnnotator = direct.ValueOf(in.RunQaAnnotator)
	out.QaConfig = AnnotatorSelector_QaConfig_ToProto(mapCtx, in.QaConfig)
	return out
}

func AnnotatorSelector_QaConfig_FromProto(mapCtx *direct.MapContext, in *pb.AnnotatorSelector_QaConfig) *krm.AnnotatorSelector_QaConfig {
	if in == nil {
		return nil
	}
	out := &krm.AnnotatorSelector_QaConfig{}
	out.ScorecardList = AnnotatorSelector_QaConfig_ScorecardList_FromProto(mapCtx, in.GetScorecardList())
	return out
}

func AnnotatorSelector_QaConfig_ToProto(mapCtx *direct.MapContext, in *krm.AnnotatorSelector_QaConfig) *pb.AnnotatorSelector_QaConfig {
	if in == nil {
		return nil
	}
	out := &pb.AnnotatorSelector_QaConfig{}
	if oneof := AnnotatorSelector_QaConfig_ScorecardList_ToProto(mapCtx, in.ScorecardList); oneof != nil {
		out.ScorecardSource = &pb.AnnotatorSelector_QaConfig_ScorecardList_{ScorecardList: oneof}
	}
	return out
}

func AnnotatorSelector_QaConfig_ScorecardList_FromProto(mapCtx *direct.MapContext, in *pb.AnnotatorSelector_QaConfig_ScorecardList) *krm.AnnotatorSelector_QaConfig_ScorecardList {
	if in == nil {
		return nil
	}
	out := &krm.AnnotatorSelector_QaConfig_ScorecardList{}
	out.QaScorecardRevisions = in.QaScorecardRevisions
	return out
}

func AnnotatorSelector_QaConfig_ScorecardList_ToProto(mapCtx *direct.MapContext, in *krm.AnnotatorSelector_QaConfig_ScorecardList) *pb.AnnotatorSelector_QaConfig_ScorecardList {
	if in == nil {
		return nil
	}
	out := &pb.AnnotatorSelector_QaConfig_ScorecardList{}
	out.QaScorecardRevisions = in.QaScorecardRevisions
	return out
}

func AnnotatorSelector_SummarizationConfig_FromProto(mapCtx *direct.MapContext, in *pb.AnnotatorSelector_SummarizationConfig) *krm.AnnotatorSelector_SummarizationConfig {
	if in == nil {
		return nil
	}
	out := &krm.AnnotatorSelector_SummarizationConfig{}
	switch modelSource := in.GetModelSource().(type) {
	case *pb.AnnotatorSelector_SummarizationConfig_ConversationProfile:
		out.ConversationProfile = direct.LazyPtr(modelSource.ConversationProfile)
	case *pb.AnnotatorSelector_SummarizationConfig_SummarizationModel_:
		out.SummarizationModel = direct.ZeroBasedEnum_FromProto(mapCtx, modelSource.SummarizationModel)
	}
	return out
}

func AnnotatorSelector_SummarizationConfig_ToProto(mapCtx *direct.MapContext, in *krm.AnnotatorSelector_SummarizationConfig) *pb.AnnotatorSelector_SummarizationConfig {
	if in == nil {
		return nil
	}
	out := &pb.AnnotatorSelector_SummarizationConfig{}
	if in.ConversationProfile != nil {
		out.ModelSource = &pb.AnnotatorSelector_SummarizationConfig_ConversationProfile{ConversationProfile: *in.ConversationProfile}
	}
	if in.SummarizationModel != nil {
		out.ModelSource = &pb.AnnotatorSelector_SummarizationConfig_SummarizationModel_{SummarizationModel: direct.Enum_ToProto[pb.AnnotatorSelector_SummarizationConfig_SummarizationModel](mapCtx, in.SummarizationModel)}
	}
	return out
}
