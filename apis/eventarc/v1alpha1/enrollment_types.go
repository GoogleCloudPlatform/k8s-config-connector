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

var EventarcEnrollmentGVK = GroupVersion.WithKind("EventarcEnrollment")

// EventarcEnrollmentSpec defines the desired state of EventarcEnrollment
// +kcc:spec:proto=google.cloud.eventarc.v1.Enrollment
type EventarcEnrollmentSpec struct {
	// The project that this resource belongs to.
	// +kubebuilder:validation:Required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +kubebuilder:validation:Required
	Location *string `json:"location,omitempty"`

	// The EventarcEnrollment name. If not given, the metadata.name will be used.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Resource labels.
	// +kcc:proto:field=google.cloud.eventarc.v1.Enrollment.labels
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Resource annotations.
	// +kcc:proto:field=google.cloud.eventarc.v1.Enrollment.annotations
	// +kubebuilder:validation:Optional
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. Resource display name.
	// +kcc:proto:field=google.cloud.eventarc.v1.Enrollment.display_name
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty"`

	// Required. A CEL expression identifying which messages this enrollment applies to.
	// +kcc:proto:field=google.cloud.eventarc.v1.Enrollment.cel_match
	// +kubebuilder:validation:Required
	CELMatch *string `json:"celMatch,omitempty"`

	// Required. Resource name of the message bus identifying the source of the messages.
	// +kubebuilder:validation:Required
	MessageBusRef *EventarcMessageBusRef `json:"messageBusRef,omitempty"`

	// Required. Destination is the Pipeline that the Enrollment is delivering to.
	// +kubebuilder:validation:Required
	DestinationRef *EventarcPipelineRef `json:"destinationRef,omitempty"`
}

// EventarcEnrollmentStatus defines the config connector machine state of EventarcEnrollment
type EventarcEnrollmentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the EventarcEnrollment resource in Google Cloud.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in Google Cloud.
	ObservedState *EventarcEnrollmentObservedState `json:"observedState,omitempty"`
}

// EventarcEnrollmentObservedState is the state of the EventarcEnrollment resource as most recently observed in Google Cloud.
// +kcc:observedstate:proto=google.cloud.eventarc.v1.Enrollment
type EventarcEnrollmentObservedState struct {
	// Output only. Server assigned unique identifier for the enrollment. The value
	// is a UUID4 string and guaranteed to remain unchanged until the resource is
	// deleted.
	// +kcc:proto:field=google.cloud.eventarc.v1.Enrollment.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. This checksum is computed by the server based on the value of
	// other fields, and might be sent only on update and delete requests to
	// ensure that the client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.eventarc.v1.Enrollment.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. The creation time.
	// +kcc:proto:field=google.cloud.eventarc.v1.Enrollment.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last-modified time.
	// +kcc:proto:field=google.cloud.eventarc.v1.Enrollment.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpeventarcenrollment;gcpeventarcenrollments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// EventarcEnrollment is the Schema for the EventarcEnrollment API
// +k8s:openapi-gen=true
type EventarcEnrollment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   EventarcEnrollmentSpec   `json:"spec,omitempty"`
	Status EventarcEnrollmentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// EventarcEnrollmentList contains a list of EventarcEnrollment
type EventarcEnrollmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventarcEnrollment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EventarcEnrollment{}, &EventarcEnrollmentList{})
}
