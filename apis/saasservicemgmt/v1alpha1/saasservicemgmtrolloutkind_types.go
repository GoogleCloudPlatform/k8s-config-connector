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

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SaaSServiceMgmtRolloutKindGVK = GroupVersion.WithKind("SaaSServiceMgmtRolloutKind")

// SaaSServiceMgmtRolloutKindSpec defines the desired state of SaaSServiceMgmtRolloutKind
// +kcc:spec:proto=google.cloud.saasplatform.saasservicemgmt.v1beta1.RolloutKind
type SaaSServiceMgmtRolloutKindSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable. The location of this resource.
	// +required
	Location *string `json:"location"`

	// The SaaSServiceMgmtRolloutKind name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="UnitKind field is immutable"
	// Required. Immutable. UnitKind that this rollout kind corresponds to.
	//  Rollouts stemming from this rollout kind will target the units of this unit
	//  kind. In other words, this defines the population of target units to be
	//  upgraded by rollouts.
	// +required
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.RolloutKind.unit_kind
	UnitKind *string `json:"unitKind"`

	// Optional. The strategy used for executing a Rollout. This is a required
	//  field.
	//
	//  There are two supported values strategies which are used to control
	//  - "Google.Cloud.Simple.AllAtOnce"
	//  - "Google.Cloud.Simple.OneLocationAtATime"
	//
	//  A rollout with one of these simple strategies will rollout across
	//  all locations defined in the associated UnitKind's Saas Locations.
	// +optional
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.RolloutKind.rollout_orchestration_strategy
	RolloutOrchestrationStrategy *string `json:"rolloutOrchestrationStrategy,omitempty"`

	// Optional. CEL(https://github.com/google/cel-spec) formatted filter string
	//  against Unit. The filter will be applied to determine the eligible unit
	//  population. This filter can only reduce, but not expand the scope of the
	//  rollout.
	// +optional
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.RolloutKind.unit_filter
	UnitFilter *string `json:"unitFilter,omitempty"`

	// Optional. The config for updating the unit kind. By default, the unit kind
	//  will be updated on the rollout start.
	// +optional
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.RolloutKind.update_unit_kind_strategy
	// +kubebuilder:validation:Enum=UPDATE_UNIT_KIND_STRATEGY_UNSPECIFIED;UPDATE_UNIT_KIND_STRATEGY_ON_START;UPDATE_UNIT_KIND_STRATEGY_NEVER
	UpdateUnitKindStrategy *string `json:"updateUnitKindStrategy,omitempty"`

	// Optional. The configuration for error budget. If the number of failed units
	//  exceeds max(allowed_count, allowed_ratio * total_units), the rollout will
	//  be paused. If not set, all units will be attempted to be updated regardless
	//  of the number of failures encountered.
	// +optional
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.RolloutKind.error_budget
	ErrorBudget *ErrorBudget `json:"errorBudget,omitempty"`
}

// SaaSServiceMgmtRolloutKindStatus defines the config connector machine state of SaaSServiceMgmtRolloutKind
type SaaSServiceMgmtRolloutKindStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SaaSServiceMgmtRolloutKind resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SaaSServiceMgmtRolloutKindObservedState `json:"observedState,omitempty"`
}

// SaaSServiceMgmtRolloutKindObservedState is the state of the SaaSServiceMgmtRolloutKind resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.saasplatform.saasservicemgmt.v1beta1.RolloutKind
type SaaSServiceMgmtRolloutKindObservedState struct {
	// Output only. The unique identifier of the resource. UID is unique in the
	//  time and space for this resource within the scope of the service. It is
	//  typically generated by the server on successful creation of a resource
	//  and must not be changed. UID is used to uniquely identify resources
	//  with resource name reuses. This should be a UUID4.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.RolloutKind.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. An opaque value that uniquely identifies a version or
	//  generation of a resource. It can be used to confirm that the client
	//  and server agree on the ordering of a resource being written.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.RolloutKind.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.RolloutKind.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was last updated. Any
	//  change to the resource made by users must refresh this value.
	//  Changes to a resource made by the service should refresh this value.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.RolloutKind.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsaasservicemgmtrolloutkind;gcpsaasservicemgmtrolloutkinds
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SaaSServiceMgmtRolloutKind is the Schema for the SaaSServiceMgmtRolloutKind API
// +k8s:openapi-gen=true
type SaaSServiceMgmtRolloutKind struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SaaSServiceMgmtRolloutKindSpec   `json:"spec,omitempty"`
	Status SaaSServiceMgmtRolloutKindStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SaaSServiceMgmtRolloutKindList contains a list of SaaSServiceMgmtRolloutKind
type SaaSServiceMgmtRolloutKindList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SaaSServiceMgmtRolloutKind `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SaaSServiceMgmtRolloutKind{}, &SaaSServiceMgmtRolloutKindList{})
}
