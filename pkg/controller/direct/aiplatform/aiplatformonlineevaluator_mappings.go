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

package aiplatform

import (
	aiplatformv1beta1pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/aiplatform/pb"
)

func OnlineEvaluator_GenerationConfig_FromProto(mapCtx *direct.MapContext, in *aiplatformv1beta1pb.GenerationConfig) *krm.GenerationConfig {
	if in == nil {
		return nil
	}
	out := &krm.GenerationConfig{}
	out.Temperature = in.Temperature
	out.TopP = in.TopP
	out.TopK = in.TopK
	out.CandidateCount = in.CandidateCount
	out.MaxOutputTokens = in.MaxOutputTokens
	out.StopSequences = in.StopSequences
	out.ResponseLogprobs = in.ResponseLogprobs
	out.Logprobs = in.Logprobs
	out.FrequencyPenalty = in.FrequencyPenalty
	out.PresencePenalty = in.PresencePenalty
	out.Seed = in.Seed
	out.ResponseMimeType = direct.LazyPtr(in.GetResponseMimeType())
	return out
}

func OnlineEvaluator_GenerationConfig_ToProto(mapCtx *direct.MapContext, in *krm.GenerationConfig) *aiplatformv1beta1pb.GenerationConfig {
	if in == nil {
		return nil
	}
	out := &aiplatformv1beta1pb.GenerationConfig{}
	out.Temperature = in.Temperature
	out.TopP = in.TopP
	out.TopK = in.TopK
	out.CandidateCount = in.CandidateCount
	out.MaxOutputTokens = in.MaxOutputTokens
	out.StopSequences = in.StopSequences
	out.ResponseLogprobs = in.ResponseLogprobs
	out.Logprobs = in.Logprobs
	out.FrequencyPenalty = in.FrequencyPenalty
	out.PresencePenalty = in.PresencePenalty
	out.Seed = in.Seed
	out.ResponseMimeType = direct.ValueOf(in.ResponseMimeType)
	return out
}

func OnlineEvaluator_PredefinedMetricSpec_FromProto(mapCtx *direct.MapContext, in *pb.PredefinedMetricSpec) *krm.PredefinedMetricSpec {
	if in == nil {
		return nil
	}
	out := &krm.PredefinedMetricSpec{}
	out.MetricSpecName = direct.LazyPtr(in.GetMetricSpecName())
	if v := direct.Struct_FromProto(mapCtx, in.GetMetricSpecParameters()); v != nil {
		out.MetricSpecParameters = *v
	}
	return out
}

func OnlineEvaluator_PredefinedMetricSpec_ToProto(mapCtx *direct.MapContext, in *krm.PredefinedMetricSpec) *pb.PredefinedMetricSpec {
	if in == nil {
		return nil
	}
	out := &pb.PredefinedMetricSpec{}
	out.MetricSpecName = direct.ValueOf(in.MetricSpecName)
	out.MetricSpecParameters = direct.Struct_ToProto(mapCtx, &in.MetricSpecParameters)
	return out
}

func OnlineEvaluator_AutoraterConfig_FromProto(mapCtx *direct.MapContext, in *pb.AutoraterConfig) *krm.AutoraterConfig {
	if in == nil {
		return nil
	}
	out := &krm.AutoraterConfig{}
	out.SamplingCount = in.SamplingCount
	out.FlipEnabled = in.FlipEnabled
	out.AutoraterModel = direct.LazyPtr(in.GetAutoraterModel())
	out.GenerationConfig = OnlineEvaluator_GenerationConfig_FromProto(mapCtx, in.GetGenerationConfig())
	return out
}

func OnlineEvaluator_AutoraterConfig_ToProto(mapCtx *direct.MapContext, in *krm.AutoraterConfig) *pb.AutoraterConfig {
	if in == nil {
		return nil
	}
	out := &pb.AutoraterConfig{}
	out.SamplingCount = in.SamplingCount
	out.FlipEnabled = in.FlipEnabled
	out.AutoraterModel = direct.ValueOf(in.AutoraterModel)
	out.GenerationConfig = OnlineEvaluator_GenerationConfig_ToProto(mapCtx, in.GenerationConfig)
	return out
}

func MetricSource_FromProto(mapCtx *direct.MapContext, in *pb.MetricSource) *krm.MetricSource {
	if in == nil {
		return nil
	}
	out := &krm.MetricSource{}
	if in.GetMetric() != nil {
		out.Metric = Metric_FromProto(mapCtx, in.GetMetric())
	}
	if in.GetMetricResourceName() != "" {
		out.MetricResourceName = direct.LazyPtr(in.GetMetricResourceName())
	}
	return out
}

func MetricSource_ToProto(mapCtx *direct.MapContext, in *krm.MetricSource) *pb.MetricSource {
	if in == nil {
		return nil
	}
	out := &pb.MetricSource{}
	if in.MetricResourceName != nil {
		out.MetricSource = &pb.MetricSource_MetricResourceName{
			MetricResourceName: direct.ValueOf(in.MetricResourceName),
		}
	} else if in.Metric != nil {
		if m := Metric_ToProto(mapCtx, in.Metric); m != nil {
			out.MetricSource = &pb.MetricSource_Metric{
				Metric: m,
			}
		}
	}
	return out
}

func Metric_FromProto(mapCtx *direct.MapContext, in *pb.Metric) *krm.Metric {
	if in == nil {
		return nil
	}
	out := &krm.Metric{}
	out.AggregationMetrics = direct.EnumSlice_FromProto(mapCtx, in.AggregationMetrics)
	if spec := in.GetPredefinedMetricSpec(); spec != nil {
		out.PredefinedMetricSpec = OnlineEvaluator_PredefinedMetricSpec_FromProto(mapCtx, spec)
	}
	if spec := in.GetComputationBasedMetricSpec(); spec != nil {
		out.ComputationBasedMetricSpec = &krm.ComputationBasedMetricSpec{}
		if spec.Type != nil {
			out.ComputationBasedMetricSpec.Type = direct.Enum_FromProto(mapCtx, spec.GetType())
		}
		if v := direct.Struct_FromProto(mapCtx, spec.GetParameters()); v != nil {
			out.ComputationBasedMetricSpec.Parameters = *v
		}
	}
	if spec := in.GetLlmBasedMetricSpec(); spec != nil {
		out.LlmBasedMetricSpec = &krm.LlmBasedMetricSpec{}
		if spec.GetRubricGroupKey() != "" {
			out.LlmBasedMetricSpec.RubricGroupKey = direct.LazyPtr(spec.GetRubricGroupKey())
		}
		if spec.GetPredefinedRubricGenerationSpec() != nil {
			out.LlmBasedMetricSpec.PredefinedRubricGenerationSpec = OnlineEvaluator_PredefinedMetricSpec_FromProto(mapCtx, spec.GetPredefinedRubricGenerationSpec())
		}
		if spec.MetricPromptTemplate != nil {
			out.LlmBasedMetricSpec.MetricPromptTemplate = spec.MetricPromptTemplate
		}
		if spec.SystemInstruction != nil {
			out.LlmBasedMetricSpec.SystemInstruction = spec.SystemInstruction
		}
		if spec.GetJudgeAutoraterConfig() != nil {
			out.LlmBasedMetricSpec.JudgeAutoraterConfig = OnlineEvaluator_AutoraterConfig_FromProto(mapCtx, spec.GetJudgeAutoraterConfig())
		}
		if v := direct.Struct_FromProto(mapCtx, spec.GetAdditionalConfig()); v != nil {
			out.LlmBasedMetricSpec.AdditionalConfig = *v
		}
	}
	if spec := in.GetPointwiseMetricSpec(); spec != nil {
		out.PointwiseMetricSpec = &krm.PointwiseMetricSpec{}
		if spec.GetMetricPromptTemplate() != "" {
			out.PointwiseMetricSpec.MetricPromptTemplate = direct.LazyPtr(spec.GetMetricPromptTemplate())
		}
		if spec.GetSystemInstruction() != "" {
			out.PointwiseMetricSpec.SystemInstruction = direct.LazyPtr(spec.GetSystemInstruction())
		}
	}
	if spec := in.GetPairwiseMetricSpec(); spec != nil {
		out.PairwiseMetricSpec = &krm.PairwiseMetricSpec{}
		if spec.GetMetricPromptTemplate() != "" {
			out.PairwiseMetricSpec.MetricPromptTemplate = direct.LazyPtr(spec.GetMetricPromptTemplate())
		}
		if spec.GetCandidateResponseFieldName() != "" {
			out.PairwiseMetricSpec.CandidateResponseFieldName = direct.LazyPtr(spec.GetCandidateResponseFieldName())
		}
		if spec.GetBaselineResponseFieldName() != "" {
			out.PairwiseMetricSpec.BaselineResponseFieldName = direct.LazyPtr(spec.GetBaselineResponseFieldName())
		}
		if spec.GetSystemInstruction() != "" {
			out.PairwiseMetricSpec.SystemInstruction = direct.LazyPtr(spec.GetSystemInstruction())
		}
	}
	if spec := in.GetExactMatchSpec(); spec != nil {
		out.ExactMatchSpec = &krm.ExactMatchSpec{}
	}
	if spec := in.GetBleuSpec(); spec != nil {
		out.BleuSpec = &krm.BleuSpec{
			UseEffectiveOrder: direct.LazyPtr(spec.GetUseEffectiveOrder()),
		}
	}
	if spec := in.GetRougeSpec(); spec != nil {
		out.RougeSpec = &krm.RougeSpec{}
		if spec.GetRougeType() != "" {
			out.RougeSpec.RougeType = direct.LazyPtr(spec.GetRougeType())
		}
		out.RougeSpec.UseStemmer = direct.LazyPtr(spec.GetUseStemmer())
		out.RougeSpec.SplitSummaries = direct.LazyPtr(spec.GetSplitSummaries())
	}
	return out
}

func Metric_ToProto(mapCtx *direct.MapContext, in *krm.Metric) *pb.Metric {
	if in == nil {
		return nil
	}
	out := &pb.Metric{}
	out.AggregationMetrics = direct.EnumSlice_ToProto[pb.Metric_AggregationMetric](mapCtx, in.AggregationMetrics)
	if in.PredefinedMetricSpec != nil {
		out.MetricSpec = &pb.Metric_PredefinedMetricSpec{
			PredefinedMetricSpec: OnlineEvaluator_PredefinedMetricSpec_ToProto(mapCtx, in.PredefinedMetricSpec),
		}
	} else if in.ComputationBasedMetricSpec != nil {
		spec := &pb.ComputationBasedMetricSpec{
			Parameters: direct.Struct_ToProto(mapCtx, &in.ComputationBasedMetricSpec.Parameters),
		}
		if in.ComputationBasedMetricSpec.Type != nil {
			t := direct.Enum_ToProto[pb.ComputationBasedMetricSpec_ComputationBasedMetricType](mapCtx, in.ComputationBasedMetricSpec.Type)
			spec.Type = &t
		}
		out.MetricSpec = &pb.Metric_ComputationBasedMetricSpec{
			ComputationBasedMetricSpec: spec,
		}
	} else if in.LlmBasedMetricSpec != nil {
		spec := &pb.LLMBasedMetricSpec{
			MetricPromptTemplate: in.LlmBasedMetricSpec.MetricPromptTemplate,
			SystemInstruction:    in.LlmBasedMetricSpec.SystemInstruction,
			AdditionalConfig:     direct.Struct_ToProto(mapCtx, &in.LlmBasedMetricSpec.AdditionalConfig),
		}
		if in.LlmBasedMetricSpec.JudgeAutoraterConfig != nil {
			spec.JudgeAutoraterConfig = OnlineEvaluator_AutoraterConfig_ToProto(mapCtx, in.LlmBasedMetricSpec.JudgeAutoraterConfig)
		}
		if in.LlmBasedMetricSpec.RubricGroupKey != nil && *in.LlmBasedMetricSpec.RubricGroupKey != "" {
			spec.RubricsSource = &pb.LLMBasedMetricSpec_RubricGroupKey{
				RubricGroupKey: *in.LlmBasedMetricSpec.RubricGroupKey,
			}
		} else if in.LlmBasedMetricSpec.PredefinedRubricGenerationSpec != nil {
			spec.RubricsSource = &pb.LLMBasedMetricSpec_PredefinedRubricGenerationSpec{
				PredefinedRubricGenerationSpec: OnlineEvaluator_PredefinedMetricSpec_ToProto(mapCtx, in.LlmBasedMetricSpec.PredefinedRubricGenerationSpec),
			}
		}
		out.MetricSpec = &pb.Metric_LlmBasedMetricSpec{
			LlmBasedMetricSpec: spec,
		}
	} else if in.PointwiseMetricSpec != nil {
		out.MetricSpec = &pb.Metric_PointwiseMetricSpec{
			PointwiseMetricSpec: &pb.PointwiseMetricSpec{
				MetricPromptTemplate: in.PointwiseMetricSpec.MetricPromptTemplate,
				SystemInstruction:    in.PointwiseMetricSpec.SystemInstruction,
			},
		}
	} else if in.PairwiseMetricSpec != nil {
		out.MetricSpec = &pb.Metric_PairwiseMetricSpec{
			PairwiseMetricSpec: &pb.PairwiseMetricSpec{
				MetricPromptTemplate:       in.PairwiseMetricSpec.MetricPromptTemplate,
				CandidateResponseFieldName: direct.ValueOf(in.PairwiseMetricSpec.CandidateResponseFieldName),
				BaselineResponseFieldName:  direct.ValueOf(in.PairwiseMetricSpec.BaselineResponseFieldName),
				SystemInstruction:          in.PairwiseMetricSpec.SystemInstruction,
			},
		}
	} else if in.ExactMatchSpec != nil {
		out.MetricSpec = &pb.Metric_ExactMatchSpec{
			ExactMatchSpec: &pb.ExactMatchSpec{},
		}
	} else if in.BleuSpec != nil {
		out.MetricSpec = &pb.Metric_BleuSpec{
			BleuSpec: &pb.BleuSpec{
				UseEffectiveOrder: direct.ValueOf(in.BleuSpec.UseEffectiveOrder),
			},
		}
	} else if in.RougeSpec != nil {
		out.MetricSpec = &pb.Metric_RougeSpec{
			RougeSpec: &pb.RougeSpec{
				RougeType:      direct.ValueOf(in.RougeSpec.RougeType),
				UseStemmer:     direct.ValueOf(in.RougeSpec.UseStemmer),
				SplitSummaries: direct.ValueOf(in.RougeSpec.SplitSummaries),
			},
		}
	}
	return out
}
