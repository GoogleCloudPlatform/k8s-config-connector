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


// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Experiment
type Experiment struct {
	// The name of the experiment.
	//  Format:
	//  projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/environments/<EnvironmentID>/experiments/<ExperimentID>.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.name
	Name *string `json:"name,omitempty"`

	// Required. The human-readable name of the experiment (unique in an
	//  environment). Limit of 64 characters.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The human-readable description of the experiment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.description
	Description *string `json:"description,omitempty"`

	// The current state of the experiment.
	//  Transition triggered by Experiments.StartExperiment: DRAFT->RUNNING.
	//  Transition triggered by Experiments.CancelExperiment: DRAFT->DONE or
	//  RUNNING->DONE.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.state
	State *string `json:"state,omitempty"`

	// The definition of the experiment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.definition
	Definition *Experiment_Definition `json:"definition,omitempty"`

	// The configuration for auto rollout. If set, there should be exactly two
	//  variants in the experiment (control variant being the default version of
	//  the flow), the traffic allocation for the non-control variant will
	//  gradually increase to 100% when conditions are met, and eventually
	//  replace the control variant to become the default version of the flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.rollout_config
	RolloutConfig *RolloutConfig `json:"rolloutConfig,omitempty"`

	// State of the auto rollout process.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.rollout_state
	RolloutState *RolloutState `json:"rolloutState,omitempty"`

	// The reason why rollout has failed. Should only be set when state is
	//  ROLLOUT_FAILED.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.rollout_failure_reason
	RolloutFailureReason *string `json:"rolloutFailureReason,omitempty"`

	// Inference result of the experiment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.result
	Result *Experiment_Result `json:"result,omitempty"`

	// Creation time of this experiment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Start time of this experiment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.start_time
	StartTime *string `json:"startTime,omitempty"`

	// End time of this experiment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Last update time of this experiment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.last_update_time
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// Maximum number of days to run the experiment. If auto-rollout is
	//  not enabled, default value and maximum will be 30 days. If auto-rollout is
	//  enabled, default value and maximum will be 6 days.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.experiment_length
	ExperimentLength *string `json:"experimentLength,omitempty"`

	// The history of updates to the experiment variants.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.variants_history
	VariantsHistory []VariantsHistory `json:"variantsHistory,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Experiment.Definition
type Experiment_Definition struct {
	// The condition defines which subset of sessions are selected for
	//  this experiment. If not specified, all sessions are eligible. E.g.
	//  "query_input.language_code=en" See the [conditions
	//  reference](https://cloud.google.com/dialogflow/cx/docs/reference/condition).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.Definition.condition
	Condition *string `json:"condition,omitempty"`

	// The flow versions as the variants of this experiment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.Definition.version_variants
	VersionVariants *VersionVariants `json:"versionVariants,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Experiment.Result
type Experiment_Result struct {
	// Version variants and metrics.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.Result.version_metrics
	VersionMetrics []Experiment_Result_VersionMetrics `json:"versionMetrics,omitempty"`

	// The last time the experiment's stats data was updated. Will have default
	//  value if stats have never been computed for this experiment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.Result.last_update_time
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Experiment.Result.ConfidenceInterval
type Experiment_Result_ConfidenceInterval struct {
	// The confidence level used to construct the interval, i.e. there is X%
	//  chance that the true value is within this interval.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.Result.ConfidenceInterval.confidence_level
	ConfidenceLevel *float64 `json:"confidenceLevel,omitempty"`

	// The percent change between an experiment metric's value and the value
	//  for its control.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.Result.ConfidenceInterval.ratio
	Ratio *float64 `json:"ratio,omitempty"`

	// Lower bound of the interval.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.Result.ConfidenceInterval.lower_bound
	LowerBound *float64 `json:"lowerBound,omitempty"`

	// Upper bound of the interval.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.Result.ConfidenceInterval.upper_bound
	UpperBound *float64 `json:"upperBound,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Experiment.Result.Metric
type Experiment_Result_Metric struct {
	// Ratio-based metric type. Only one of type or count_type is specified in
	//  each Metric.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.Result.Metric.type
	Type *string `json:"type,omitempty"`

	// Count-based metric type. Only one of type or count_type is specified in
	//  each Metric.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.Result.Metric.count_type
	CountType *string `json:"countType,omitempty"`

	// Ratio value of a metric.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.Result.Metric.ratio
	Ratio *float64 `json:"ratio,omitempty"`

	// Count value of a metric.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.Result.Metric.count
	Count *float64 `json:"count,omitempty"`

	// The probability that the treatment is better than all other treatments
	//  in the experiment
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.Result.Metric.confidence_interval
	ConfidenceInterval *Experiment_Result_ConfidenceInterval `json:"confidenceInterval,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Experiment.Result.VersionMetrics
type Experiment_Result_VersionMetrics struct {
	// The name of the flow
	//  [Version][google.cloud.dialogflow.cx.v3beta1.Version]. Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/versions/<VersionID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.Result.VersionMetrics.version
	Version *string `json:"version,omitempty"`

	// The metrics and corresponding confidence intervals in the inference
	//  result.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.Result.VersionMetrics.metrics
	Metrics []Experiment_Result_Metric `json:"metrics,omitempty"`

	// Number of sessions that were allocated to this version.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Experiment.Result.VersionMetrics.session_count
	SessionCount *int32 `json:"sessionCount,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.RolloutConfig
type RolloutConfig struct {
	// Steps to roll out a flow version. Steps should be sorted by percentage in
	//  ascending order.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.RolloutConfig.rollout_steps
	RolloutSteps []RolloutConfig_RolloutStep `json:"rolloutSteps,omitempty"`

	// The conditions that are used to evaluate the success of a rollout
	//  step. If not specified, all rollout steps will proceed to the next one
	//  unless failure conditions are met. E.g. "containment_rate > 60% AND
	//  callback_rate < 20%". See the [conditions
	//  reference](https://cloud.google.com/dialogflow/cx/docs/reference/condition).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.RolloutConfig.rollout_condition
	RolloutCondition *string `json:"rolloutCondition,omitempty"`

	// The conditions that are used to evaluate the failure of a rollout
	//  step. If not specified, no rollout steps will fail. E.g. "containment_rate
	//  < 10% OR average_turn_count < 3". See the [conditions
	//  reference](https://cloud.google.com/dialogflow/cx/docs/reference/condition).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.RolloutConfig.failure_condition
	FailureCondition *string `json:"failureCondition,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.RolloutConfig.RolloutStep
type RolloutConfig_RolloutStep struct {
	// The name of the rollout step;
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.RolloutConfig.RolloutStep.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The percentage of traffic allocated to the flow version of this rollout
	//  step. (0%, 100%].
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.RolloutConfig.RolloutStep.traffic_percent
	TrafficPercent *int32 `json:"trafficPercent,omitempty"`

	// The minimum time that this step should last. Should be longer than 1
	//  hour. If not set, the default minimum duration for each step will be 1
	//  hour.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.RolloutConfig.RolloutStep.min_duration
	MinDuration *string `json:"minDuration,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.RolloutState
type RolloutState struct {
	// Display name of the current auto rollout step.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.RolloutState.step
	Step *string `json:"step,omitempty"`

	// Index of the current step in the auto rollout steps list.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.RolloutState.step_index
	StepIndex *int32 `json:"stepIndex,omitempty"`

	// Start time of the current step.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.RolloutState.start_time
	StartTime *string `json:"startTime,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.VariantsHistory
type VariantsHistory struct {
	// The flow versions as the variants.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.VariantsHistory.version_variants
	VersionVariants *VersionVariants `json:"versionVariants,omitempty"`

	// Update time of the variants.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.VariantsHistory.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.VersionVariants
type VersionVariants struct {
	// A list of flow version variants.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.VersionVariants.variants
	Variants []VersionVariants_Variant `json:"variants,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.VersionVariants.Variant
type VersionVariants_Variant struct {
	// The name of the flow version.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/versions/<VersionID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.VersionVariants.Variant.version
	Version *string `json:"version,omitempty"`

	// Percentage of the traffic which should be routed to this
	//  version of flow. Traffic allocation for a single flow must sum up to 1.0.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.VersionVariants.Variant.traffic_allocation
	TrafficAllocation *float32 `json:"trafficAllocation,omitempty"`

	// Whether the variant is for the control group.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.VersionVariants.Variant.is_control_group
	IsControlGroup *bool `json:"isControlGroup,omitempty"`
}
