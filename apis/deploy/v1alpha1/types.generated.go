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


// +kcc:proto=google.cloud.deploy.v1.AdvanceRolloutRule
type AdvanceRolloutRule struct {
	// Required. ID of the rule. This id must be unique in the `Automation`
	//  resource to which this rule belongs. The format is
	//  `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`.
	// +kcc:proto:field=google.cloud.deploy.v1.AdvanceRolloutRule.id
	ID *string `json:"id,omitempty"`

	// Optional. Proceeds only after phase name matched any one in the list.
	//  This value must consist of lower-case letters, numbers, and hyphens,
	//  start with a letter and end with a letter or a number, and have a max
	//  length of 63 characters. In other words, it must match the following
	//  regex: `^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$`.
	// +kcc:proto:field=google.cloud.deploy.v1.AdvanceRolloutRule.source_phases
	SourcePhases []string `json:"sourcePhases,omitempty"`

	// Optional. How long to wait after a rollout is finished.
	// +kcc:proto:field=google.cloud.deploy.v1.AdvanceRolloutRule.wait
	Wait *string `json:"wait,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.Automation
type Automation struct {

	// Optional. Description of the `Automation`. Max length is 255 characters.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.description
	Description *string `json:"description,omitempty"`

	// Optional. User annotations. These attributes can only be set and used by
	//  the user, and not by Cloud Deploy. Annotations must meet the following
	//  constraints:
	//
	//  * Annotations are key/value pairs.
	//  * Valid annotation keys have two segments: an optional prefix and name,
	//  separated by a slash (`/`).
	//  * The name segment is required and must be 63 characters or less,
	//  beginning and ending with an alphanumeric character (`[a-z0-9A-Z]`) with
	//  dashes (`-`), underscores (`_`), dots (`.`), and alphanumerics between.
	//  * The prefix is optional. If specified, the prefix must be a DNS subdomain:
	//  a series of DNS labels separated by dots(`.`), not longer than 253
	//  characters in total, followed by a slash (`/`).
	//
	//  See
	//  https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/#syntax-and-character-set
	//  for more details.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. Labels are attributes that can be set and used by both the
	//  user and by Cloud Deploy. Labels must meet the following constraints:
	//
	//  * Keys and values can contain only lowercase letters, numeric characters,
	//  underscores, and dashes.
	//  * All characters must use UTF-8 encoding, and international characters are
	//  allowed.
	//  * Keys must start with a lowercase letter or international character.
	//  * Each resource is limited to a maximum of 64 labels.
	//
	//  Both keys and values are additionally constrained to be <= 63 characters.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The weak etag of the `Automation` resource.
	//  This checksum is computed by the server based on the value of other
	//  fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. When Suspended, automation is deactivated from execution.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.suspended
	Suspended *bool `json:"suspended,omitempty"`

	// Required. Email address of the user-managed IAM service account that
	//  creates Cloud Deploy release and rollout resources.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Required. Selected resources to which the automation will be applied.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.selector
	Selector *AutomationResourceSelector `json:"selector,omitempty"`

	// Required. List of Automation rules associated with the Automation resource.
	//  Must have at least one rule and limited to 250 rules per Delivery Pipeline.
	//  Note: the order of the rules here is not the same as the order of
	//  execution.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.rules
	Rules []AutomationRule `json:"rules,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.AutomationResourceSelector
type AutomationResourceSelector struct {
	// Contains attributes about a target.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationResourceSelector.targets
	Targets []TargetAttribute `json:"targets,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.AutomationRule
type AutomationRule struct {
	// Optional. `PromoteReleaseRule` will automatically promote a release from
	//  the current target to a specified target.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationRule.promote_release_rule
	PromoteReleaseRule *PromoteReleaseRule `json:"promoteReleaseRule,omitempty"`

	// Optional. The `AdvanceRolloutRule` will automatically advance a
	//  successful Rollout.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationRule.advance_rollout_rule
	AdvanceRolloutRule *AdvanceRolloutRule `json:"advanceRolloutRule,omitempty"`

	// Optional. The `RepairRolloutRule` will automatically repair a failed
	//  rollout.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationRule.repair_rollout_rule
	RepairRolloutRule *RepairRolloutRule `json:"repairRolloutRule,omitempty"`

	// Optional. The `TimedPromoteReleaseRule` will automatically promote a
	//  release from the current target(s) to the specified target(s) on a
	//  configured schedule.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationRule.timed_promote_release_rule
	TimedPromoteReleaseRule *TimedPromoteReleaseRule `json:"timedPromoteReleaseRule,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.AutomationRuleCondition
type AutomationRuleCondition struct {
	// Optional. Details around targets enumerated in the rule.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationRuleCondition.targets_present_condition
	TargetsPresentCondition *TargetsPresentCondition `json:"targetsPresentCondition,omitempty"`

	// Optional. TimedPromoteReleaseCondition contains rule conditions specific
	//  to a an Automation with a timed promote release rule defined.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationRuleCondition.timed_promote_release_condition
	TimedPromoteReleaseCondition *TimedPromoteReleaseCondition `json:"timedPromoteReleaseCondition,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.PromoteReleaseRule
type PromoteReleaseRule struct {
	// Required. ID of the rule. This id must be unique in the `Automation`
	//  resource to which this rule belongs. The format is
	//  `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`.
	// +kcc:proto:field=google.cloud.deploy.v1.PromoteReleaseRule.id
	ID *string `json:"id,omitempty"`

	// Optional. How long the release need to be paused until being promoted to
	//  the next target.
	// +kcc:proto:field=google.cloud.deploy.v1.PromoteReleaseRule.wait
	Wait *string `json:"wait,omitempty"`

	// Optional. The ID of the stage in the pipeline to which this `Release` is
	//  deploying. If unspecified, default it to the next stage in the promotion
	//  flow. The value of this field could be one of the following:
	//
	//  * The last segment of a target name
	//  * "@next", the next target in the promotion sequence
	// +kcc:proto:field=google.cloud.deploy.v1.PromoteReleaseRule.destination_target_id
	DestinationTargetID *string `json:"destinationTargetID,omitempty"`

	// Optional. The starting phase of the rollout created by this operation.
	//  Default to the first phase.
	// +kcc:proto:field=google.cloud.deploy.v1.PromoteReleaseRule.destination_phase
	DestinationPhase *string `json:"destinationPhase,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.RepairPhaseConfig
type RepairPhaseConfig struct {
	// Optional. Retries a failed job.
	// +kcc:proto:field=google.cloud.deploy.v1.RepairPhaseConfig.retry
	Retry *Retry `json:"retry,omitempty"`

	// Optional. Rolls back a `Rollout`.
	// +kcc:proto:field=google.cloud.deploy.v1.RepairPhaseConfig.rollback
	Rollback *Rollback `json:"rollback,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.RepairRolloutRule
type RepairRolloutRule struct {
	// Required. ID of the rule. This id must be unique in the `Automation`
	//  resource to which this rule belongs. The format is
	//  `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`.
	// +kcc:proto:field=google.cloud.deploy.v1.RepairRolloutRule.id
	ID *string `json:"id,omitempty"`

	// Optional. Phases within which jobs are subject to automatic repair actions
	//  on failure. Proceeds only after phase name matched any one in the list, or
	//  for all phases if unspecified. This value must consist of lower-case
	//  letters, numbers, and hyphens, start with a letter and end with a letter or
	//  a number, and have a max length of 63 characters. In other words, it must
	//  match the following regex: `^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$`.
	// +kcc:proto:field=google.cloud.deploy.v1.RepairRolloutRule.phases
	Phases []string `json:"phases,omitempty"`

	// Optional. Jobs to repair. Proceeds only after job name matched any one in
	//  the list, or for all jobs if unspecified or empty. The phase that includes
	//  the job must match the phase ID specified in `source_phase`. This value
	//  must consist of lower-case letters, numbers, and hyphens, start with a
	//  letter and end with a letter or a number, and have a max length of 63
	//  characters. In other words, it must match the following regex:
	//  `^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$`.
	// +kcc:proto:field=google.cloud.deploy.v1.RepairRolloutRule.jobs
	Jobs []string `json:"jobs,omitempty"`

	// Required. Defines the types of automatic repair phases for failed jobs.
	// +kcc:proto:field=google.cloud.deploy.v1.RepairRolloutRule.repair_phases
	RepairPhases []RepairPhaseConfig `json:"repairPhases,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.Retry
type Retry struct {
	// Required. Total number of retries. Retry is skipped if set to 0; The
	//  minimum value is 1, and the maximum value is 10.
	// +kcc:proto:field=google.cloud.deploy.v1.Retry.attempts
	Attempts *int64 `json:"attempts,omitempty"`

	// Optional. How long to wait for the first retry. Default is 0, and the
	//  maximum value is 14d.
	// +kcc:proto:field=google.cloud.deploy.v1.Retry.wait
	Wait *string `json:"wait,omitempty"`

	// Optional. The pattern of how wait time will be increased. Default is
	//  linear. Backoff mode will be ignored if `wait` is 0.
	// +kcc:proto:field=google.cloud.deploy.v1.Retry.backoff_mode
	BackoffMode *string `json:"backoffMode,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.Rollback
type Rollback struct {
	// Optional. The starting phase ID for the `Rollout`. If unspecified, the
	//  `Rollout` will start in the stable phase.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollback.destination_phase
	DestinationPhase *string `json:"destinationPhase,omitempty"`

	// Optional. If pending rollout exists on the target, the rollback operation
	//  will be aborted.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollback.disable_rollback_if_rollout_pending
	DisableRollbackIfRolloutPending *bool `json:"disableRollbackIfRolloutPending,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.TargetAttribute
type TargetAttribute struct {
	// ID of the `Target`. The value of this field could be one of the
	//  following:
	//
	//  * The last segment of a target name
	//  * "*", all targets in a location
	// +kcc:proto:field=google.cloud.deploy.v1.TargetAttribute.id
	ID *string `json:"id,omitempty"`

	// Target labels.
	// +kcc:proto:field=google.cloud.deploy.v1.TargetAttribute.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.TargetsPresentCondition
type TargetsPresentCondition struct {
	// True if there aren't any missing Targets.
	// +kcc:proto:field=google.cloud.deploy.v1.TargetsPresentCondition.status
	Status *bool `json:"status,omitempty"`

	// The list of Target names that do not exist. For example,
	//  `projects/{project_id}/locations/{location_name}/targets/{target_name}`.
	// +kcc:proto:field=google.cloud.deploy.v1.TargetsPresentCondition.missing_targets
	MissingTargets []string `json:"missingTargets,omitempty"`

	// Last time the condition was updated.
	// +kcc:proto:field=google.cloud.deploy.v1.TargetsPresentCondition.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.TimedPromoteReleaseCondition
type TimedPromoteReleaseCondition struct {
}

// +kcc:proto=google.cloud.deploy.v1.TimedPromoteReleaseCondition.Targets
type TimedPromoteReleaseCondition_Targets struct {
	// Optional. The source target ID.
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseCondition.Targets.source_target_id
	SourceTargetID *string `json:"sourceTargetID,omitempty"`

	// Optional. The destination target ID.
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseCondition.Targets.destination_target_id
	DestinationTargetID *string `json:"destinationTargetID,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.TimedPromoteReleaseRule
type TimedPromoteReleaseRule struct {
	// Required. ID of the rule. This ID must be unique in the `Automation`
	//  resource to which this rule belongs. The format is
	//  `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`.
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseRule.id
	ID *string `json:"id,omitempty"`

	// Optional. The ID of the stage in the pipeline to which this `Release` is
	//  deploying. If unspecified, default it to the next stage in the promotion
	//  flow. The value of this field could be one of the following:
	//
	//  * The last segment of a target name
	//  * "@next", the next target in the promotion sequence
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseRule.destination_target_id
	DestinationTargetID *string `json:"destinationTargetID,omitempty"`

	// Required. Schedule in crontab format. e.g. "0 9 * * 1" for every Monday at
	//  9am.
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseRule.schedule
	Schedule *string `json:"schedule,omitempty"`

	// Required. The time zone in IANA format [IANA Time Zone
	//  Database](https://www.iana.org/time-zones) (e.g. America/New_York).
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseRule.time_zone
	TimeZone *string `json:"timeZone,omitempty"`

	// Optional. The starting phase of the rollout created by this rule. Default
	//  to the first phase.
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseRule.destination_phase
	DestinationPhase *string `json:"destinationPhase,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.AdvanceRolloutRule
type AdvanceRolloutRuleObservedState struct {
	// Output only. Information around the state of the Automation rule.
	// +kcc:proto:field=google.cloud.deploy.v1.AdvanceRolloutRule.condition
	Condition *AutomationRuleCondition `json:"condition,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.Automation
type AutomationObservedState struct {
	// Output only. Name of the `Automation`. Format is
	//  `projects/{project}/locations/{location}/deliveryPipelines/{delivery_pipeline}/automations/{automation}`.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.name
	Name *string `json:"name,omitempty"`

	// Output only. Unique identifier of the `Automation`.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Time at which the automation was created.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time at which the automation was updated.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Required. List of Automation rules associated with the Automation resource.
	//  Must have at least one rule and limited to 250 rules per Delivery Pipeline.
	//  Note: the order of the rules here is not the same as the order of
	//  execution.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.rules
	Rules []AutomationRuleObservedState `json:"rules,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.AutomationRule
type AutomationRuleObservedState struct {
	// Optional. `PromoteReleaseRule` will automatically promote a release from
	//  the current target to a specified target.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationRule.promote_release_rule
	PromoteReleaseRule *PromoteReleaseRuleObservedState `json:"promoteReleaseRule,omitempty"`

	// Optional. The `AdvanceRolloutRule` will automatically advance a
	//  successful Rollout.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationRule.advance_rollout_rule
	AdvanceRolloutRule *AdvanceRolloutRuleObservedState `json:"advanceRolloutRule,omitempty"`

	// Optional. The `RepairRolloutRule` will automatically repair a failed
	//  rollout.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationRule.repair_rollout_rule
	RepairRolloutRule *RepairRolloutRuleObservedState `json:"repairRolloutRule,omitempty"`

	// Optional. The `TimedPromoteReleaseRule` will automatically promote a
	//  release from the current target(s) to the specified target(s) on a
	//  configured schedule.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationRule.timed_promote_release_rule
	TimedPromoteReleaseRule *TimedPromoteReleaseRuleObservedState `json:"timedPromoteReleaseRule,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.AutomationRuleCondition
type AutomationRuleConditionObservedState struct {
	// Optional. TimedPromoteReleaseCondition contains rule conditions specific
	//  to a an Automation with a timed promote release rule defined.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationRuleCondition.timed_promote_release_condition
	TimedPromoteReleaseCondition *TimedPromoteReleaseConditionObservedState `json:"timedPromoteReleaseCondition,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.PromoteReleaseRule
type PromoteReleaseRuleObservedState struct {
	// Output only. Information around the state of the Automation rule.
	// +kcc:proto:field=google.cloud.deploy.v1.PromoteReleaseRule.condition
	Condition *AutomationRuleCondition `json:"condition,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.RepairRolloutRule
type RepairRolloutRuleObservedState struct {
	// Output only. Information around the state of the 'Automation' rule.
	// +kcc:proto:field=google.cloud.deploy.v1.RepairRolloutRule.condition
	Condition *AutomationRuleCondition `json:"condition,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.TimedPromoteReleaseCondition
type TimedPromoteReleaseConditionObservedState struct {
	// Output only. When the next scheduled promotion(s) will occur.
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseCondition.next_promotion_time
	NextPromotionTime *string `json:"nextPromotionTime,omitempty"`

	// Output only. A list of targets involved in the upcoming timed promotion(s).
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseCondition.targets_list
	TargetsList []TimedPromoteReleaseCondition_Targets `json:"targetsList,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.TimedPromoteReleaseRule
type TimedPromoteReleaseRuleObservedState struct {
	// Output only. Information around the state of the Automation rule.
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseRule.condition
	Condition *AutomationRuleCondition `json:"condition,omitempty"`
}
