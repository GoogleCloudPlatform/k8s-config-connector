// Copyright 2024 Google LLC
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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var LoggingLinkGVK = GroupVersion.WithKind("LoggingLink")

// LoggingLinkSpec defines the desired state of LoggingLink
// +kcc:proto=google.logging.v2.Link
type LoggingLinkSpec struct {
	// Immutable.
	// The LoggingLink name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Describes this link.
	//  The maximum length of the description is 8000 characters.
	Description *string `json:"description,omitempty"`

	// The LoggingLogBucket that this Link is associated with
	LoggingLogBucketRef *refs.LoggingLogBucketRef `json:"loggingLogBucketRef,omitempty"`
}

// LoggingLinkStatus defines the config connector machine state of LoggingLink
type LoggingLinkStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the LoggingLink resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *LoggingLinkObservedState `json:"observedState,omitempty"`
}

// LoggingLinkObservedState is the state of the LoggingLink resource as most recently observed in GCP.
type LoggingLinkObservedState struct {

	// +optional
	// Output only. The creation timestamp of the link.
	CreateTime *string `json:"createTime,omitempty"`

	// the lifecycle state might be something more complicated
	// this is an ENUM, should be safe to use a string

	// Output only. The resource lifecycle state.
	LifecycleState *string `json:"lifecycleState,omitempty"`

	// this field is just a string, but its an object(string)
	// https://github.com/googleapis/googleapis/blob/master/google/logging/v2/logging_config.proto#L1063

	// The information of a BigQuery Dataset. When a link is created, a BigQuery dataset is created
	// along with it, in the same project as the LogBucket it's linked to. This dataset will also have
	// BigQuery Views corresponding to the LogViews in the bucket.
	BigQueryDataset *BigQueryDatasetObservedState `json:"bigQueryDataset,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcplogginglink;gcplogginglinks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// LoggingLink is the Schema for the LoggingLink API
// +k8s:openapi-gen=true
type LoggingLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   LoggingLinkSpec   `json:"spec,omitempty"`
	Status LoggingLinkStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// LoggingLinkList contains a list of LoggingLink
type LoggingLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoggingLink `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LoggingLink{}, &LoggingLinkList{})
}
