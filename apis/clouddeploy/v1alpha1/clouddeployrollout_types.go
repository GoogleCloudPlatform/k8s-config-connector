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

// API sources for CloudDeployRollout, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/deploy/v1/cloud_deploy.proto
// +kcc:source:service-docs=https://cloud.google.com/deploy/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/deploy/docs/reference/rest/v1/projects.locations.deliveryPipelines.releases.rollouts

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudDeployRolloutGVK = GroupVersion.WithKind("CloudDeployRollout")

// CloudDeployRolloutSpec defines the desired state of CloudDeployRollout
// +kcc:spec:proto=google.cloud.deploy.v1.Rollout
// +kcc:required-from-proto
type CloudDeployRolloutSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/deliveryPipelines/{delivery_pipeline}/releases/{release} this resource belongs to.
	// +kcc:guess
	// ReleaseRef *CloudDeployReleaseRef `json:"releaseRef,omitempty"`

	// The CloudDeployRollout name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Description of the `Rollout` for user purposes. Max length is 255
	//  characters.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.description
	Description *string `json:"description,omitempty"`

	// Optional. User annotations. These attributes can only be set and used by
	//  the user, and not by Cloud Deploy. See
	//  https://google.aip.dev/128#annotations for more details such as format and
	//  size limitations.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Labels are attributes that can be set and used by both the
	//  user and by Cloud Deploy. Labels must meet the following constraints:
	//
	//  * Keys and values can contain only lowercase letters, numeric characters,
	//  underscores, and dashes.
	//  * All characters must use UTF-8 encoding, and international characters are
	//  allowed.
	//  * Keys must start with a lowercase letter or international character.
	//  * Each resource is limited to a maximum of 64 labels.
	//
	//  Both keys and values are additionally constrained to be <= 128 bytes.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. The ID of Target to which this `Rollout` is deploying.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.target_id
	// +required
	TargetID *string `json:"targetID,omitempty"`

	// This checksum is computed by the server based on the value of other
	//  fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.etag
	Etag *string `json:"etag,omitempty"`
}

// CloudDeployRolloutStatus defines the config connector machine state of CloudDeployRollout
type CloudDeployRolloutStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudDeployRollout resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudDeployRolloutObservedState `json:"observedState,omitempty"`
}

// CloudDeployRolloutObservedState is the state of the CloudDeployRollout resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.deploy.v1.Rollout
type CloudDeployRolloutObservedState struct {
	// Output only. Unique identifier of the `Rollout`.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Time at which the `Rollout` was created.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time at which the `Rollout` was approved.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.approve_time
	ApproveTime *string `json:"approveTime,omitempty"`

	// Output only. Time at which the `Rollout` was enqueued.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.enqueue_time
	EnqueueTime *string `json:"enqueueTime,omitempty"`

	// Output only. Time at which the `Rollout` started deploying.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.deploy_start_time
	DeployStartTime *string `json:"deployStartTime,omitempty"`

	// Output only. Time at which the `Rollout` finished deploying.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.deploy_end_time
	DeployEndTime *string `json:"deployEndTime,omitempty"`

	// Output only. Approval state of the `Rollout`.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.approval_state
	ApprovalState *string `json:"approvalState,omitempty"`

	// Output only. Current state of the `Rollout`.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.state
	State *string `json:"state,omitempty"`

	// Output only. Additional information about the rollout failure, if
	//  available.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.failure_reason
	FailureReason *string `json:"failureReason,omitempty"`

	// Output only. The resource name of the Cloud Build `Build` object that is
	//  used to deploy the Rollout. Format is
	//  `projects/{project}/locations/{location}/builds/{build}`.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.deploying_build
	DeployingBuild *string `json:"deployingBuild,omitempty"`

	// Output only. The reason this rollout failed. This will always be
	//  unspecified while the rollout is in progress.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.deploy_failure_cause
	DeployFailureCause *string `json:"deployFailureCause,omitempty"`

	// Output only. The phases that represent the workflows of this `Rollout`.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.phases
	Phases []PhaseObservedState `json:"phases,omitempty"`

	// Output only. Metadata contains information about the rollout.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.metadata
	Metadata *MetadataObservedState `json:"metadata,omitempty"`

	// Output only. Name of the `ControllerRollout`. Format is
	//  `projects/{project}/locations/{location}/deliveryPipelines/{deliveryPipeline}/releases/{release}/rollouts/{rollout}`.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.controller_rollout
	ControllerRollout *string `json:"controllerRollout,omitempty"`

	// Output only. Name of the `Rollout` that is rolled back by this `Rollout`.
	//  Empty if this `Rollout` wasn't created as a rollback.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.rollback_of_rollout
	RollbackOfRollout *string `json:"rollbackOfRollout,omitempty"`

	// Output only. Names of `Rollouts` that rolled back this `Rollout`.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.rolled_back_by_rollouts
	RolledBackByRollouts []string `json:"rolledBackByRollouts,omitempty"`

	// Output only. The AutomationRun actively repairing the rollout.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.active_repair_automation_run
	ActiveRepairAutomationRun *string `json:"activeRepairAutomationRun,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpclouddeployrollout;gcpclouddeployrollouts
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudDeployRollout is the Schema for the CloudDeployRollout API
// +k8s:openapi-gen=true
type CloudDeployRollout struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudDeployRolloutSpec   `json:"spec,omitempty"`
	Status CloudDeployRolloutStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudDeployRolloutList contains a list of CloudDeployRollout
type CloudDeployRolloutList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudDeployRollout `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudDeployRollout{}, &CloudDeployRolloutList{})
}
