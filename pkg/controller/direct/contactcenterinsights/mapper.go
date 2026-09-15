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

package contactcenterinsights

import (
	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contactcenterinsights/v1alpha1"
	krmdialogflowv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CCInsightsAnalysisRuleAnnotatorSelector_FromProto(mapCtx *direct.MapContext, in *pb.AnnotatorSelector) *krm.CCInsightsAnalysisRuleAnnotatorSelector {
	if in == nil {
		return nil
	}
	out := &krm.CCInsightsAnalysisRuleAnnotatorSelector{}
	out.RunInterruptionAnnotator = direct.LazyPtr(in.GetRunInterruptionAnnotator())
	out.RunSilenceAnnotator = direct.LazyPtr(in.GetRunSilenceAnnotator())
	out.RunPhraseMatcherAnnotator = direct.LazyPtr(in.GetRunPhraseMatcherAnnotator())

	if v := in.GetPhraseMatchers(); len(v) != 0 {
		for i := range v {
			out.PhraseMatcherRefs = append(out.PhraseMatcherRefs, krm.CCInsightsPhraseMatcherRef{External: v[i]})
		}
	}

	out.RunSentimentAnnotator = direct.LazyPtr(in.GetRunSentimentAnnotator())
	out.RunEntityAnnotator = direct.LazyPtr(in.GetRunEntityAnnotator())
	out.RunIntentAnnotator = direct.LazyPtr(in.GetRunIntentAnnotator())
	out.RunIssueModelAnnotator = direct.LazyPtr(in.GetRunIssueModelAnnotator())

	if v := in.GetIssueModels(); len(v) != 0 {
		for i := range v {
			out.IssueModelRefs = append(out.IssueModelRefs, krm.CCInsightsIssueModelRef{External: v[i]})
		}
	}

	out.RunSummarizationAnnotator = direct.LazyPtr(in.GetRunSummarizationAnnotator())
	out.SummarizationConfig = CCInsightsAnalysisRuleAnnotatorSelector_SummarizationConfig_FromProto(mapCtx, in.GetSummarizationConfig())
	out.RunQaAnnotator = direct.LazyPtr(in.GetRunQaAnnotator())
	out.QaConfig = AnnotatorSelector_QaConfig_FromProto(mapCtx, in.GetQaConfig())
	return out
}

func CCInsightsAnalysisRuleAnnotatorSelector_ToProto(mapCtx *direct.MapContext, in *krm.CCInsightsAnalysisRuleAnnotatorSelector) *pb.AnnotatorSelector {
	if in == nil {
		return nil
	}
	out := &pb.AnnotatorSelector{}
	out.RunInterruptionAnnotator = direct.ValueOf(in.RunInterruptionAnnotator)
	out.RunSilenceAnnotator = direct.ValueOf(in.RunSilenceAnnotator)
	out.RunPhraseMatcherAnnotator = direct.ValueOf(in.RunPhraseMatcherAnnotator)

	if v := in.PhraseMatcherRefs; len(v) != 0 {
		for i := range v {
			out.PhraseMatchers = append(out.PhraseMatchers, v[i].External)
		}
	}

	out.RunSentimentAnnotator = direct.ValueOf(in.RunSentimentAnnotator)
	out.RunEntityAnnotator = direct.ValueOf(in.RunEntityAnnotator)
	out.RunIntentAnnotator = direct.ValueOf(in.RunIntentAnnotator)
	out.RunIssueModelAnnotator = direct.ValueOf(in.RunIssueModelAnnotator)

	if v := in.IssueModelRefs; len(v) != 0 {
		for i := range v {
			out.IssueModels = append(out.IssueModels, v[i].External)
		}
	}

	out.RunSummarizationAnnotator = direct.ValueOf(in.RunSummarizationAnnotator)
	out.SummarizationConfig = CCInsightsAnalysisRuleAnnotatorSelector_SummarizationConfig_ToProto(mapCtx, in.SummarizationConfig)
	out.RunQaAnnotator = direct.ValueOf(in.RunQaAnnotator)
	out.QaConfig = AnnotatorSelector_QaConfig_ToProto(mapCtx, in.QaConfig)
	return out
}

func CCInsightsAnalysisRuleAnnotatorSelector_SummarizationConfig_FromProto(mapCtx *direct.MapContext, in *pb.AnnotatorSelector_SummarizationConfig) *krm.CCInsightsAnalysisRuleAnnotatorSelector_SummarizationConfig {
	if in == nil {
		return nil
	}
	out := &krm.CCInsightsAnalysisRuleAnnotatorSelector_SummarizationConfig{}
	if in.GetConversationProfile() != "" {
		out.ConversationProfileRef = &krmdialogflowv1alpha1.DialogflowConversationProfileRef{External: in.GetConversationProfile()}
	}
	out.SummarizationModel = direct.Enum_FromProto(mapCtx, in.GetSummarizationModel())
	return out
}

func CCInsightsAnalysisRuleAnnotatorSelector_SummarizationConfig_ToProto(mapCtx *direct.MapContext, in *krm.CCInsightsAnalysisRuleAnnotatorSelector_SummarizationConfig) *pb.AnnotatorSelector_SummarizationConfig {
	if in == nil {
		return nil
	}
	out := &pb.AnnotatorSelector_SummarizationConfig{}
	if in.ConversationProfileRef != nil {
		out.ModelSource = &pb.AnnotatorSelector_SummarizationConfig_ConversationProfile{ConversationProfile: in.ConversationProfileRef.External}
	}
	if oneof := AnnotatorSelector_SummarizationConfig_SummarizationModel_ToProto(mapCtx, in.SummarizationModel); oneof != nil {
		out.ModelSource = oneof
	}
	return out
}
