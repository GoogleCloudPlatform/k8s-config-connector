// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddeploy/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	commonv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/common/v1alpha1"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudDeployAutomationGVK = GroupVersion.WithKind("CloudDeployAutomation")

// CloudDeployAutomationSpec defines the desired state of CloudDeployAutomation
// +kcc:spec:proto=google.cloud.deploy.v1.Automation
type CloudDeployAutomationSpec struct {
	commonv1alpha1.CommonSpec `json:",inline"`

	// Immutable. The location of this resource.
	// +required
	Location *string `json:"location,omitempty"`

	// Immutable. Required. The DeliveryPipeline that this automation belongs to.
	// +required
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.delivery_pipeline
	DeliveryPipelineRef *v1beta1.DeliveryPipelineRef `json:"deliveryPipelineRef,omitempty"`

	// Optional. Description of the `Automation`. Max length is 255 characters.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.description
	Description *string `json:"description,omitempty"`

	// Optional. User annotations. These attributes can only be set and used by
	// the user, and not by Cloud Deploy. See
	// https://google.aip.dev/128#annotations for more details such as format and
	// size limitations.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. When Suspended, automation is deactivated from execution.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.suspended
	Suspended *bool `json:"suspended,omitempty"`

	// Required. Email address of the user-managed IAM service account that
	//  creates Cloud Deploy release and rollout resources.
	// +required
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.service_account
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Required. Selected resources to which the automation will be applied.
	// +required
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.selector
	Selector *AutomationResourceSelector `json:"selector,omitempty"`

	// Required. List of Automation rules associated with the Automation resource.
	// +required
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.rules
	Rules []AutomationRule `json:"rules,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.AutomationResourceSelector
type AutomationResourceSelector struct {
	// Optional. Contains attributes about a target.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationResourceSelector.targets
	Targets []TargetAttribute `json:"targets,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.TargetAttribute
type TargetAttribute struct {
	// Optional. The Target to which the rule applies.
	// +kcc:proto:field=google.cloud.deploy.v1.TargetAttribute.id
	TargetRef *CloudDeployTargetRef `json:"targetRef,omitempty"`

	// Target labels.
	// +kcc:proto:field=google.cloud.deploy.v1.TargetAttribute.labels
	Labels map[string]string `json:"labels,omitempty"`
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

// +kcc:proto=google.cloud.deploy.v1.PromoteReleaseRule
type PromoteReleaseRule struct {
	// Required. ID of the rule.
	// +required
	// +kcc:proto:field=google.cloud.deploy.v1.PromoteReleaseRule.id
	ID *string `json:"id,omitempty"`

	// Optional. How long the release need to be paused until being promoted to
	//  the next target.
	// +kcc:proto:field=google.cloud.deploy.v1.PromoteReleaseRule.wait
	Wait *string `json:"wait,omitempty"`

	// Optional. The ID of the stage in the pipeline to which this `Release` is
	//  deploying.
	// +kcc:proto:field=google.cloud.deploy.v1.PromoteReleaseRule.destination_target_id
	DestinationTargetRef *CloudDeployTargetRef `json:"destinationTargetRef,omitempty"`

	// Optional. The starting phase of the rollout created by this operation.
	//  Default to the first phase.
	// +kcc:proto:field=google.cloud.deploy.v1.PromoteReleaseRule.destination_phase
	DestinationPhase *string `json:"destinationPhase,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.AdvanceRolloutRule
type AdvanceRolloutRule struct {
	// Required. ID of the rule.
	// +required
	// +kcc:proto:field=google.cloud.deploy.v1.AdvanceRolloutRule.id
	ID *string `json:"id,omitempty"`

	// Optional. Proceeds only after phase name matched any one in the list.
	// +kcc:proto:field=google.cloud.deploy.v1.AdvanceRolloutRule.source_phases
	SourcePhases []string `json:"sourcePhases,omitempty"`

	// Optional. How long to wait after a rollout is finished.
	// +kcc:proto:field=google.cloud.deploy.v1.AdvanceRolloutRule.wait
	Wait *string `json:"wait,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.RepairRolloutRule
type RepairRolloutRule struct {
	// Required. ID of the rule.
	// +required
	// +kcc:proto:field=google.cloud.deploy.v1.RepairRolloutRule.id
	ID *string `json:"id,omitempty"`

	// Optional. Phases within which jobs are subject to automatic repair actions
	//  on failure.
	// +kcc:proto:field=google.cloud.deploy.v1.RepairRolloutRule.phases
	Phases []string `json:"phases,omitempty"`

	// Optional. Jobs to repair.
	// +kcc:proto:field=google.cloud.deploy.v1.RepairRolloutRule.jobs
	Jobs []string `json:"jobs,omitempty"`

	// Required. Defines the types of automatic repair phases for failed jobs.
	// +required
	// +kcc:proto:field=google.cloud.deploy.v1.RepairRolloutRule.repair_phases
	RepairPhases []RepairPhaseConfig `json:"repairPhases,omitempty"`
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

// +kcc:proto=google.cloud.deploy.v1.Retry
type Retry struct {
	// Required. Total number of retries.
	// +required
	// +kcc:proto:field=google.cloud.deploy.v1.Retry.attempts
	Attempts *int64 `json:"attempts,omitempty"`

	// Optional. How long to wait for the first retry.
	// +kcc:proto:field=google.cloud.deploy.v1.Retry.wait
	Wait *string `json:"wait,omitempty"`

	// Optional. The pattern of how wait time will be increased.
	// +kcc:proto:field=google.cloud.deploy.v1.Retry.backoff_mode
	BackoffMode *string `json:"backoffMode,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.Rollback
type Rollback struct {
	// Optional. The starting phase ID for the `Rollout`.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollback.destination_phase
	DestinationPhase *string `json:"destinationPhase,omitempty"`

	// Optional. If pending rollout exists on the target, the rollback operation
	//  will be aborted.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollback.disable_rollback_if_rollout_pending
	DisableRollbackIfRolloutPending *bool `json:"disableRollbackIfRolloutPending,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.TimedPromoteReleaseRule
type TimedPromoteReleaseRule struct {
	// Required. ID of the rule.
	// +required
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseRule.id
	ID *string `json:"id,omitempty"`

	// Optional. The ID of the stage in the pipeline to which this `Release` is
	//  deploying.
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseRule.destination_target_id
	DestinationTargetRef *CloudDeployTargetRef `json:"destinationTargetRef,omitempty"`

	// Required. Schedule in crontab format.
	// +required
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseRule.schedule
	Schedule *string `json:"schedule,omitempty"`

	// Required. The time zone in IANA format.
	// +required
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseRule.time_zone
	TimeZone *string `json:"timeZone,omitempty"`

	// Optional. The starting phase of the rollout created by this rule.
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseRule.destination_phase
	DestinationPhase *string `json:"destinationPhase,omitempty"`
}

// CloudDeployAutomationStatus defines the config connector machine state of CloudDeployAutomation
type CloudDeployAutomationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []k8sv1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudDeployAutomation resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudDeployAutomationObservedState `json:"observedState,omitempty"`
}

// CloudDeployAutomationObservedState is the state of the CloudDeployAutomation resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.deploy.v1.Automation
type CloudDeployAutomationObservedState struct {
	// Output only. Unique identifier of the `Automation`.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Time at which the automation was created.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time at which the automation was updated.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. The weak etag of the `Automation` resource.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.etag
	Etag *string `json:"etag,omitempty"`

	// Required. List of Automation rules associated with the Automation resource.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.rules
	Rules []AutomationRuleObservedState `json:"rules,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.deploy.v1.AutomationRule
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

// +kcc:observedstate:proto=google.cloud.deploy.v1.PromoteReleaseRule
type PromoteReleaseRuleObservedState struct {
	// Output only. Information around the state of the Automation rule.
	// +kcc:proto:field=google.cloud.deploy.v1.PromoteReleaseRule.condition
	Condition *AutomationRuleCondition `json:"condition,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.deploy.v1.AdvanceRolloutRule
type AdvanceRolloutRuleObservedState struct {
	// Output only. Information around the state of the Automation rule.
	// +kcc:proto:field=google.cloud.deploy.v1.AdvanceRolloutRule.condition
	Condition *AutomationRuleCondition `json:"condition,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.deploy.v1.RepairRolloutRule
type RepairRolloutRuleObservedState struct {
	// Output only. Information around the state of the 'Automation' rule.
	// +kcc:proto:field=google.cloud.deploy.v1.RepairRolloutRule.condition
	Condition *AutomationRuleCondition `json:"condition,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.deploy.v1.TimedPromoteReleaseRule
type TimedPromoteReleaseRuleObservedState struct {
	// Output only. Information around the state of the Automation rule.
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseRule.condition
	Condition *AutomationRuleCondition `json:"condition,omitempty"`
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

// +kcc:proto=google.cloud.deploy.v1.TargetsPresentCondition
type TargetsPresentCondition struct {
	// True if there aren't any missing Targets.
	// +kcc:proto:field=google.cloud.deploy.v1.TargetsPresentCondition.status
	Status *bool `json:"status,omitempty"`

	// The list of Target names that do not exist.
	// +kcc:proto:field=google.cloud.deploy.v1.TargetsPresentCondition.missing_targets
	MissingTargets []string `json:"missingTargets,omitempty"`

	// Last time the condition was updated.
	// +kcc:proto:field=google.cloud.deploy.v1.TargetsPresentCondition.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.TimedPromoteReleaseCondition
type TimedPromoteReleaseCondition struct {
	// Output only. When the next scheduled promotion(s) will occur.
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseCondition.next_promotion_time
	NextPromotionTime *string `json:"nextPromotionTime,omitempty"`

	// Output only. A list of targets involved in the upcoming timed promotion(s).
	// +kcc:proto:field=google.cloud.deploy.v1.TimedPromoteReleaseCondition.targets_list
	TargetsList []TimedPromoteReleaseCondition_Targets `json:"targetsList,omitempty"`
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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpclouddeployautomation;gcpclouddeployautomations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudDeployAutomation is the Schema for the CloudDeployAutomation API
// +k8s:openapi-gen=true
type CloudDeployAutomation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudDeployAutomationSpec   `json:"spec,omitempty"`
	Status CloudDeployAutomationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudDeployAutomationList contains a list of CloudDeployAutomation
type CloudDeployAutomationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudDeployAutomation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudDeployAutomation{}, &CloudDeployAutomationList{})
}
