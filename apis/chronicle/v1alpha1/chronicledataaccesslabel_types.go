// Copyright 2025 Google LLC
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

// API sources for ChronicleDataAccessLabel, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/chronicle/v1/data_access_control.proto
// +kcc:source:service-docs=https://cloud.google.com/chronicle/docs/secops/secops-overview
// +kcc:source:resource-docs=https://docs.cloud.google.com/chronicle/docs/reference/rest/v1/projects.locations.instances.dataAccessLabels

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ChronicleDataAccessLabelGVK = GroupVersion.WithKind("ChronicleDataAccessLabel")

// ChronicleDataAccessLabelSpec defines the desired state of ChronicleDataAccessLabel
// +kcc:spec:proto=google.cloud.chronicle.v1.DataAccessLabel
// +kcc:required-from-proto
type ChronicleDataAccessLabelSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The ChronicleDataAccessLabel name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// A UDM query over event data.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataAccessLabel.udm_query
	UdmQuery *string `json:"udmQuery,omitempty"`

	// Optional. A description of the data access label for a human reader.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataAccessLabel.description
	Description *string `json:"description,omitempty"`
}

// ChronicleDataAccessLabelStatus defines the config connector machine state of ChronicleDataAccessLabel
type ChronicleDataAccessLabelStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ChronicleDataAccessLabel resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ChronicleDataAccessLabelObservedState `json:"observedState,omitempty"`
}

// ChronicleDataAccessLabelObservedState is the state of the ChronicleDataAccessLabel resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.chronicle.v1.DataAccessLabel
type ChronicleDataAccessLabelObservedState struct {
	// Output only. The short name displayed for the label as it appears on event
	//  data.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataAccessLabel.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. The time at which the data access label was created.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataAccessLabel.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the data access label was last updated.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataAccessLabel.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The user who created the data access label.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataAccessLabel.author
	Author *string `json:"author,omitempty"`

	// Output only. The user who last updated the data access label.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataAccessLabel.last_editor
	LastEditor *string `json:"lastEditor,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpchronicledataaccesslabel;gcpchronicledataaccesslabels
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ChronicleDataAccessLabel is the Schema for the ChronicleDataAccessLabel API
// +k8s:openapi-gen=true
type ChronicleDataAccessLabel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ChronicleDataAccessLabelSpec   `json:"spec,omitempty"`
	Status ChronicleDataAccessLabelStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ChronicleDataAccessLabelList contains a list of ChronicleDataAccessLabel
type ChronicleDataAccessLabelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChronicleDataAccessLabel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ChronicleDataAccessLabel{}, &ChronicleDataAccessLabelList{})
}
