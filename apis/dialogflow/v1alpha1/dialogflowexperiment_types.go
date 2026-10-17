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

// API sources for DialogflowExperiment, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/dialogflow/cx/v3/experiment.proto
// +kcc:source:service-docs=https://cloud.google.com/dialogflow/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/dialogflow/docs/reference/rest/v3/projects.locations.agents.environments.experiments

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowExperimentGVK = GroupVersion.WithKind("DialogflowExperiment")

// DialogflowExperimentSpec defines the desired state of DialogflowExperiment
// +kcc:spec:proto=google.cloud.dialogflow.cx.v3.Experiment
// +kcc:required-from-proto
type DialogflowExperimentSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/agents/{agent}/environments/{environment} this resource belongs to.
	// +kcc:guess
	// EnvironmentRef *DialogflowEnvironmentRef `json:"environmentRef,omitempty"`

	// The DialogflowExperiment name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The human-readable name of the experiment (unique in an
	//  environment). Limit of 64 characters.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Experiment.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// The human-readable description of the experiment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Experiment.description
	Description *string `json:"description,omitempty"`

	// The current state of the experiment.
	//  Transition triggered by Experiments.StartExperiment: DRAFT->RUNNING.
	//  Transition triggered by Experiments.CancelExperiment: DRAFT->DONE or
	//  RUNNING->DONE.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Experiment.state
	State *string `json:"state,omitempty"`

	// The definition of the experiment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Experiment.definition
	Definition *Experiment_Definition `json:"definition,omitempty"`

	// The configuration for auto rollout. If set, there should be exactly two
	//  variants in the experiment (control variant being the default version of
	//  the flow), the traffic allocation for the non-control variant will
	//  gradually increase to 100% when conditions are met, and eventually
	//  replace the control variant to become the default version of the flow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Experiment.rollout_config
	RolloutConfig *RolloutConfig `json:"rolloutConfig,omitempty"`

	// State of the auto rollout process.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Experiment.rollout_state
	RolloutState *RolloutState `json:"rolloutState,omitempty"`

	// The reason why rollout has failed. Should only be set when state is
	//  ROLLOUT_FAILED.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Experiment.rollout_failure_reason
	RolloutFailureReason *string `json:"rolloutFailureReason,omitempty"`

	// Inference result of the experiment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Experiment.result
	Result *Experiment_Result `json:"result,omitempty"`

	// Creation time of this experiment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Experiment.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Start time of this experiment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Experiment.start_time
	StartTime *string `json:"startTime,omitempty"`

	// End time of this experiment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Experiment.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Last update time of this experiment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Experiment.last_update_time
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// Maximum number of days to run the experiment/rollout. If auto-rollout is
	//  not enabled, default value and maximum will be 30 days. If auto-rollout is
	//  enabled, default value and maximum will be 6 days.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Experiment.experiment_length
	ExperimentLength *string `json:"experimentLength,omitempty"`

	// The history of updates to the experiment variants.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Experiment.variants_history
	VariantsHistory []VariantsHistory `json:"variantsHistory,omitempty"`
}

// DialogflowExperimentStatus defines the config connector machine state of DialogflowExperiment
type DialogflowExperimentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowExperiment resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *DialogflowExperimentObservedState `json:"observedState,omitempty"`
}

// DialogflowExperimentObservedState is the state of the DialogflowExperiment resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.cx.v3.Experiment
// type DialogflowExperimentObservedState struct {
// }

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowexperiment;gcpdialogflowexperiments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowExperiment is the Schema for the DialogflowExperiment API
// +k8s:openapi-gen=true
type DialogflowExperiment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowExperimentSpec   `json:"spec,omitempty"`
	Status DialogflowExperimentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowExperimentList contains a list of DialogflowExperiment
type DialogflowExperimentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowExperiment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowExperiment{}, &DialogflowExperimentList{})
}
