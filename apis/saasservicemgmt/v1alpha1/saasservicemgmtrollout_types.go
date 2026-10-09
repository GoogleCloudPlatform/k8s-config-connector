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

// API sources for SaaSServiceMgmtRollout, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/saasplatform/saasservicemgmt/v1beta1/rollouts_resources.proto
// +kcc:source:service-docs=https://cloud.google.com/saas-runtime/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/saas-runtime/docs/reference/rest/v1beta1/projects.locations.rollouts

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SaaSServiceMgmtRolloutGVK = GroupVersion.WithKind("SaaSServiceMgmtRollout")

// SaaSServiceMgmtRolloutSpec defines the desired state of SaaSServiceMgmtRollout
// +kcc:spec:proto=google.cloud.saasplatform.saasservicemgmt.v1beta1.Rollout
type SaaSServiceMgmtRolloutSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The SaaSServiceMgmtRollout name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Immutable. Name of the Release that gets rolled out to target
	//  Units. Required if no other type of release is specified.
	// +kcc:guess=possible-reference target=SaasServiceMgmtRelease
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Rollout.release
	Release *string `json:"release,omitempty"`

	// Optional. The strategy used for executing this Rollout.
	//  This strategy will override whatever strategy is specified in the
	//  RolloutType. If not specified on creation, the
	//  strategy from RolloutType will be used.
	//
	//  There are two supported values strategies which are used to control
	//  - "Google.Cloud.Simple.AllAtOnce"
	//  - "Google.Cloud.Simple.OneLocationAtATime"
	//
	//  A rollout with one of these simple strategies will rollout across
	//  all locations defined in the targeted UnitKind's Saas Locations.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Rollout.rollout_orchestration_strategy
	RolloutOrchestrationStrategy *string `json:"rolloutOrchestrationStrategy,omitempty"`

	// Optional. CEL(https://github.com/google/cel-spec) formatted filter string
	//  against Unit. The filter will be applied to determine the eligible unit
	//  population. This filter can only reduce, but not expand the scope of the
	//  rollout. If not provided, the unit_filter from the RolloutType will be
	//  used.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Rollout.unit_filter
	UnitFilter *string `json:"unitFilter,omitempty"`

	// Required. Immutable. Name of the RolloutKind this rollout is stemming from
	//  and adhering to.
	// +required
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Rollout.rollout_kind
	RolloutKind *string `json:"rolloutKind"`

	// Optional. Requested change to the execution of this rollout.
	//  Default RolloutControl.action is ROLLOUT_ACTION_RUN meaning
	//  the rollout will be executed to completion while progressing through
	//  all natural Rollout States (such as RUNNING -> SUCCEEDED or RUNNING ->
	//  FAILED). Requests can only be made when the Rollout is in a non-terminal
	//  state.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Rollout.control
	Control *RolloutControl `json:"control,omitempty"`

	// Optional. The labels on the resource, which can be used for categorization.
	//  similar to Kubernetes resource labels.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Rollout.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Annotations is an unstructured key-value map stored with a
	//  resource that may be set by external tools to store and retrieve arbitrary
	//  metadata. They are not queryable and should be preserved when modifying
	//  objects.
	//
	//  More info: https://kubernetes.io/docs/user-guide/annotations
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Rollout.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// SaaSServiceMgmtRolloutStatus defines the config connector machine state of SaaSServiceMgmtRollout
type SaaSServiceMgmtRolloutStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SaaSServiceMgmtRollout resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SaaSServiceMgmtRolloutObservedState `json:"observedState,omitempty"`
}

// SaaSServiceMgmtRolloutObservedState is the state of the SaaSServiceMgmtRollout resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.saasplatform.saasservicemgmt.v1beta1.Rollout
type SaaSServiceMgmtRolloutObservedState struct {
	// Optional. Output only. The time when the rollout started executing. Will be
	//  empty if the rollout hasn't started yet.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Rollout.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Optional. Output only. The time when the rollout finished execution
	//  (regardless of  success, failure, or cancellation). Will be empty if the
	//  rollout hasn't finished yet. Once set, the rollout is in terminal state and
	//  all the results are final.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Rollout.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Current state of the rollout.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Rollout.state
	State *string `json:"state,omitempty"`

	// Output only. Human readable message indicating details about the last state
	//  transition.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Rollout.state_message
	StateMessage *string `json:"stateMessage,omitempty"`

	// Optional. Output only. The time when the rollout transitioned into its
	//  current state.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Rollout.state_transition_time
	StateTransitionTime *string `json:"stateTransitionTime,omitempty"`

	// Optional. Output only. The root rollout that this rollout is stemming from.
	//  The resource name (full URI of the resource) following the standard naming
	//  scheme:
	//
	//    "projects/{project}/locations/{location}/rollouts/{rollout_id}"
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Rollout.root_rollout
	RootRollout *string `json:"rootRollout,omitempty"`

	// Optional. Output only. The direct parent rollout that this rollout is
	//  stemming from. The resource name (full URI of the resource) following the
	//  standard naming scheme:
	//
	//    "projects/{project}/locations/{location}/rollouts/{rollout_id}"
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Rollout.parent_rollout
	ParentRollout *string `json:"parentRollout,omitempty"`

	// Optional. Output only. Details about the progress of the rollout.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Rollout.stats
	Stats *RolloutStatsObservedState `json:"stats,omitempty"`

	// Output only. The unique identifier of the resource. UID is unique in the
	//  time and space for this resource within the scope of the service. It is
	//  typically generated by the server on successful creation of a resource
	//  and must not be changed. UID is used to uniquely identify resources
	//  with resource name reuses. This should be a UUID4.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Rollout.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. An opaque value that uniquely identifies a version or
	//  generation of a resource. It can be used to confirm that the client
	//  and server agree on the ordering of a resource being written.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Rollout.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Rollout.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was last updated. Any
	//  change to the resource made by users must refresh this value.
	//  Changes to a resource made by the service should refresh this value.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Rollout.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsaasservicemgmtrollout;gcpsaasservicemgmtrollouts
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SaaSServiceMgmtRollout is the Schema for the SaaSServiceMgmtRollout API
// +k8s:openapi-gen=true
type SaaSServiceMgmtRollout struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SaaSServiceMgmtRolloutSpec   `json:"spec,omitempty"`
	Status SaaSServiceMgmtRolloutStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SaaSServiceMgmtRolloutList contains a list of SaaSServiceMgmtRollout
type SaaSServiceMgmtRolloutList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SaaSServiceMgmtRollout `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SaaSServiceMgmtRollout{}, &SaaSServiceMgmtRolloutList{})
}
