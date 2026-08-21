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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SaasServiceMgmtReleaseGVK = GroupVersion.WithKind("SaasServiceMgmtRelease")

// SaasServiceMgmtReleaseSpec defines the desired state of SaasServiceMgmtRelease
// +kcc:spec:proto=google.cloud.saasplatform.saasservicemgmt.v1beta1.Release
type SaasServiceMgmtReleaseSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable. The location of this resource.
	// +required
	Location *string `json:"location"`

	// The SaasServiceMgmtRelease name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="UnitKind field is immutable"
	// Required. Immutable. Reference to the UnitKind this Release corresponds to
	//  (required and immutable once created).
	// +required
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Release.unit_kind
	UnitKind *string `json:"unitKind"`

	// Optional. Blueprints are OCI Images that contain all of the artifacts
	//  needed to provision a unit.
	// +optional
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Release.blueprint
	Blueprint *Blueprint `json:"blueprint,omitempty"`

	// Optional. Set of requirements to be fulfilled on the Unit when using this
	//  Release.
	// +optional
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Release.release_requirements
	ReleaseRequirements *Release_ReleaseRequirements `json:"releaseRequirements,omitempty"`

	// Optional. Mapping of input variables to default values. Maximum 100
	// +optional
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Release.input_variable_defaults
	InputVariableDefaults []UnitVariable `json:"inputVariableDefaults,omitempty"`
}

// SaasServiceMgmtReleaseStatus defines the config connector machine state of SaasServiceMgmtRelease
type SaasServiceMgmtReleaseStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SaasServiceMgmtRelease resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SaasServiceMgmtReleaseObservedState `json:"observedState,omitempty"`
}

// SaasServiceMgmtReleaseObservedState is the state of the SaasServiceMgmtRelease resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.saasplatform.saasservicemgmt.v1beta1.Release
type SaasServiceMgmtReleaseObservedState struct {
	// Optional. Blueprints are OCI Images that contain all of the artifacts
	//  needed to provision a unit.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Release.blueprint
	Blueprint *BlueprintObservedState `json:"blueprint,omitempty"`

	// Optional. Output only. List of input variables declared on the blueprint
	//  and can be present with their values on the unit spec
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Release.input_variables
	InputVariables []UnitVariable `json:"inputVariables,omitempty"`

	// Optional. Output only. List of output variables declared on the blueprint
	//  and can be present with their values on the unit status
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Release.output_variables
	OutputVariables []UnitVariable `json:"outputVariables,omitempty"`

	// Output only. The unique identifier of the resource. UID is unique in the
	//  time and space for this resource within the scope of the service. It is
	//  typically generated by the server on successful creation of a resource
	//  and must not be changed. UID is used to uniquely identify resources
	//  with resource name reuses. This should be a UUID4.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Release.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. An opaque value that uniquely identifies a version or
	//  generation of a resource. It can be used to confirm that the client
	//  and server agree on the ordering of a resource being written.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Release.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Release.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was last updated. Any
	//  change to the resource made by users must refresh this value.
	//  Changes to a resource made by the service should refresh this value.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Release.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsaasservicemgmtrelease;gcpsaasservicemgmtreleases
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SaasServiceMgmtRelease is the Schema for the SaasServiceMgmtRelease API
// +k8s:openapi-gen=true
type SaasServiceMgmtRelease struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SaasServiceMgmtReleaseSpec   `json:"spec,omitempty"`
	Status SaasServiceMgmtReleaseStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SaasServiceMgmtReleaseList contains a list of SaasServiceMgmtRelease
type SaasServiceMgmtReleaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SaasServiceMgmtRelease `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SaasServiceMgmtRelease{}, &SaasServiceMgmtReleaseList{})
}
