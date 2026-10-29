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

// API sources for SecurityposturePostureDeployment, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/securityposture/v1/securityposture.proto
// +kcc:source:service-docs=https://cloud.google.com/security-command-center
// +kcc:source:resource-docs=https://docs.cloud.google.com/security-command-center/docs/reference/securityposture/rest/v1/organizations.locations.postureDeployments

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SecurityposturePostureDeploymentGVK = GroupVersion.WithKind("SecurityposturePostureDeployment")

// SecurityposturePostureDeploymentSpec defines the desired state of SecurityposturePostureDeployment
// +kcc:spec:proto=google.cloud.securityposture.v1.PostureDeployment
// +kcc:required-from-proto
type SecurityposturePostureDeploymentSpec struct {
	// The organization that this resource belongs to.
	OrganizationRef *refsv1beta1.OrganizationRef `json:"organizationRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The SecurityposturePostureDeployment name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Target resource where the Posture will be deployed. Currently
	//  supported resources are of types: projects/projectNumber,
	//  folders/folderNumber, organizations/organizationNumber.
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.target_resource
	// +required
	TargetResource *string `json:"targetResource,omitempty"`

	// Required. Posture that needs to be deployed.
	//  Format:
	//  organizations/{org_id}/locations/{location_id}/postures/<posture>
	//  Example:
	//  organizations/99/locations/global/postures/les-miserables.
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.posture_id
	// +required
	PostureID *string `json:"postureID,omitempty"`

	// Required. Revision_id of the Posture that is to be deployed.
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.posture_revision_id
	// +required
	PostureRevisionID *string `json:"postureRevisionID,omitempty"`

	// Optional. User provided description of the PostureDeployment.
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.description
	Description *string `json:"description,omitempty"`

	// Optional. An opaque tag indicating the current version of the
	//  PostureDeployment, used for concurrency control. When the
	//  `PostureDeployment` is returned from either a `GetPostureDeployment` or a
	//  `ListPostureDeployments` request, this `etag` indicates the version of the
	//  current `PostureDeployment` to use when executing a read-modify-write loop.
	//
	//  When the `PostureDeployment` is used in a `UpdatePostureDeployment` method,
	//  use the `etag` value that was returned from a `GetPostureDeployment`
	//  request as part of a read-modify-write loop for concurrency control. Not
	//  setting the `etag` in a `UpdatePostureDeployment` request will result in an
	//  unconditional write of the `PostureDeployment`.
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. User annotations. These attributes can only be set and used by
	//  the user, and not by Google Security Postures.
	//  .
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// SecurityposturePostureDeploymentStatus defines the config connector machine state of SecurityposturePostureDeployment
type SecurityposturePostureDeploymentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SecurityposturePostureDeployment resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SecurityposturePostureDeploymentObservedState `json:"observedState,omitempty"`
}

// SecurityposturePostureDeploymentObservedState is the state of the SecurityposturePostureDeployment resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.securityposture.v1.PostureDeployment
type SecurityposturePostureDeploymentObservedState struct {
	// Output only. State of PostureDeployment resource.
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.state
	State *string `json:"state,omitempty"`

	// Output only. The timestamp that the PostureDeployment was created.
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp that the PostureDeployment was updated.
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Whether or not this Posture is in the process of being
	//  updated.
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. This is a output only optional field which will be filled in
	//  case where PostureDeployment state is UPDATE_FAILED or CREATE_FAILED or
	//  DELETE_FAILED. It denotes the desired Posture.
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.desired_posture_id
	DesiredPostureID *string `json:"desiredPostureID,omitempty"`

	// Output only. Output only optional field which provides revision_id of the
	//  desired_posture_id.
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.desired_posture_revision_id
	DesiredPostureRevisionID *string `json:"desiredPostureRevisionID,omitempty"`

	// Output only. This is a output only optional field which will be filled in
	//  case where PostureDeployment enters a failure state like UPDATE_FAILED or
	//  CREATE_FAILED or DELETE_FAILED.
	// +kcc:proto:field=google.cloud.securityposture.v1.PostureDeployment.failure_message
	FailureMessage *string `json:"failureMessage,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsecuritypostureposturedeployment;gcpsecuritypostureposturedeployments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SecurityposturePostureDeployment is the Schema for the SecurityposturePostureDeployment API
// +k8s:openapi-gen=true
type SecurityposturePostureDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SecurityposturePostureDeploymentSpec   `json:"spec,omitempty"`
	Status SecurityposturePostureDeploymentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SecurityposturePostureDeploymentList contains a list of SecurityposturePostureDeployment
type SecurityposturePostureDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityposturePostureDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecurityposturePostureDeployment{}, &SecurityposturePostureDeploymentList{})
}
