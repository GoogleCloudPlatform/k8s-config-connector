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

package v1alpha1


// +kcc:proto=google.cloud.dialogflow.v2.ConversationModelEvaluation
type ConversationModelEvaluation struct {
	// The resource name of the evaluation. Format:
	//  `projects/<Project ID>/conversationModels/<Conversation Model
	//  ID>/evaluations/<Evaluation ID>`
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationModelEvaluation.name
	Name *string `json:"name,omitempty"`

	// Optional. The display name of the model evaluation. At most 64 bytes long.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationModelEvaluation.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The configuration of the evaluation task.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationModelEvaluation.evaluation_config
	EvaluationConfig *EvaluationConfig `json:"evaluationConfig,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.EvaluationConfig
type EvaluationConfig struct {
	// Required. Datasets used for evaluation.
	// +kcc:proto:field=google.cloud.dialogflow.v2.EvaluationConfig.datasets
	Datasets []InputDataset `json:"datasets,omitempty"`

	// Configuration for smart reply model evalution.
	// +kcc:proto:field=google.cloud.dialogflow.v2.EvaluationConfig.smart_reply_config
	SmartReplyConfig *EvaluationConfig_SmartReplyConfig `json:"smartReplyConfig,omitempty"`

	// Configuration for smart compose model evalution.
	// +kcc:proto:field=google.cloud.dialogflow.v2.EvaluationConfig.smart_compose_config
	SmartComposeConfig *EvaluationConfig_SmartComposeConfig `json:"smartComposeConfig,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.EvaluationConfig.SmartComposeConfig
type EvaluationConfig_SmartComposeConfig struct {
	// The allowlist document resource name.
	//  Format: `projects/<Project ID>/knowledgeBases/<Knowledge Base
	//  ID>/documents/<Document ID>`. Only used for smart compose model.
	// +kcc:proto:field=google.cloud.dialogflow.v2.EvaluationConfig.SmartComposeConfig.allowlist_document
	AllowlistDocument *string `json:"allowlistDocument,omitempty"`

	// Required. The model to be evaluated can return multiple results with
	//  confidence score on each query. These results will be sorted by the
	//  descending order of the scores and we only keep the first
	//  max_result_count results as the final results to evaluate.
	// +kcc:proto:field=google.cloud.dialogflow.v2.EvaluationConfig.SmartComposeConfig.max_result_count
	MaxResultCount *int32 `json:"maxResultCount,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.EvaluationConfig.SmartReplyConfig
type EvaluationConfig_SmartReplyConfig struct {
	// The allowlist document resource name.
	//  Format: `projects/<Project ID>/knowledgeBases/<Knowledge Base
	//  ID>/documents/<Document ID>`. Only used for smart reply model.
	// +kcc:proto:field=google.cloud.dialogflow.v2.EvaluationConfig.SmartReplyConfig.allowlist_document
	AllowlistDocument *string `json:"allowlistDocument,omitempty"`

	// Required. The model to be evaluated can return multiple results with
	//  confidence score on each query. These results will be sorted by the
	//  descending order of the scores and we only keep the first
	//  max_result_count results as the final results to evaluate.
	// +kcc:proto:field=google.cloud.dialogflow.v2.EvaluationConfig.SmartReplyConfig.max_result_count
	MaxResultCount *int32 `json:"maxResultCount,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.InputDataset
type InputDataset struct {
	// Required. ConversationDataset resource name. Format:
	//  `projects/<Project ID>/locations/<Location
	//  ID>/conversationDatasets/<Conversation Dataset ID>`
	// +kcc:proto:field=google.cloud.dialogflow.v2.InputDataset.dataset
	Dataset *string `json:"dataset,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.SmartReplyMetrics
type SmartReplyMetrics struct {
	// Percentage of target participant messages in the evaluation dataset for
	//  which similar messages have appeared at least once in the allowlist. Should
	//  be [0, 1].
	// +kcc:proto:field=google.cloud.dialogflow.v2.SmartReplyMetrics.allowlist_coverage
	AllowlistCoverage *float32 `json:"allowlistCoverage,omitempty"`

	// Metrics of top n smart replies, sorted by [TopNMetric.n][].
	// +kcc:proto:field=google.cloud.dialogflow.v2.SmartReplyMetrics.top_n_metrics
	TopNMetrics []SmartReplyMetrics_TopNMetrics `json:"topNMetrics,omitempty"`

	// Total number of conversations used to generate this metric.
	// +kcc:proto:field=google.cloud.dialogflow.v2.SmartReplyMetrics.conversation_count
	ConversationCount *int64 `json:"conversationCount,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.SmartReplyMetrics.TopNMetrics
type SmartReplyMetrics_TopNMetrics struct {
	// Number of retrieved smart replies. For example, when `n` is 3, this
	//  evaluation contains metrics for when Dialogflow retrieves 3 smart replies
	//  with the model.
	// +kcc:proto:field=google.cloud.dialogflow.v2.SmartReplyMetrics.TopNMetrics.n
	N *int32 `json:"n,omitempty"`

	// Defined as `number of queries whose top n smart replies have at least one
	//  similar (token match similarity above the defined threshold) reply as the
	//  real reply` divided by `number of queries with at least one smart reply`.
	//  Value ranges from 0.0 to 1.0 inclusive.
	// +kcc:proto:field=google.cloud.dialogflow.v2.SmartReplyMetrics.TopNMetrics.recall
	Recall *float32 `json:"recall,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.ConversationModelEvaluation
type ConversationModelEvaluationObservedState struct {
	// Output only. Creation time of this model.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationModelEvaluation.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Only available when model is for smart reply.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationModelEvaluation.smart_reply_metrics
	SmartReplyMetrics *SmartReplyMetrics `json:"smartReplyMetrics,omitempty"`

	// Output only. Human eval template in csv format.
	//  It tooks real-world conversations provided through input dataset, generates
	//  example suggestions for customer to verify quality of the model.
	//  For Smart Reply, the generated csv file contains columns of
	//  Context, (Suggestions,Q1,Q2)*3, Actual reply.
	//  Context contains at most 10 latest messages in the conversation prior to
	//  the current suggestion.
	//  Q1: "Would you send it as the next message of agent?"
	//  Evaluated based on whether the suggest is appropriate to be sent by
	//  agent in current context.
	//  Q2: "Does the suggestion move the conversation closer to resolution?"
	//  Evaluated based on whether the suggestion provide solutions, or answers
	//  customer's question or collect information from customer to resolve the
	//  customer's issue.
	//  Actual reply column contains the actual agent reply sent in the context.
	// +kcc:proto:field=google.cloud.dialogflow.v2.ConversationModelEvaluation.raw_human_eval_template_csv
	RawHumanEvalTemplateCsv *string `json:"rawHumanEvalTemplateCsv,omitempty"`
}
