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

// API sources for DatalineageRun, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/datacatalog/lineage/v1/lineage.proto
// +kcc:source:service-docs=https://cloud.google.com/data-catalog
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.processes.runs

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DatalineageRunGVK = GroupVersion.WithKind("DatalineageRun")

// DatalineageRunSpec defines the desired state of DatalineageRun
// +kcc:spec:proto=google.cloud.datacatalog.lineage.v1.Run
// +kcc:required-from-proto
type DatalineageRunSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/processes/{process} this resource belongs to.
	// +kcc:guess
	// ProcessRef *DataLineageProcessRef `json:"processRef,omitempty"`

	// The DatalineageRun name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. A human-readable name you can set to display in a user interface.
	//  Must be not longer than 1024 characters and only contain UTF-8 letters
	//  or numbers, spaces or characters like `_-:&.`
	// +kcc:proto:field=google.cloud.datacatalog.lineage.v1.Run.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The attributes of the run. Should only be used for the purpose of
	//  non-semantic management (classifying, describing or labeling the run).
	//
	//  Up to 100 attributes are allowed.
	// +kcc:proto:field=google.cloud.datacatalog.lineage.v1.Run.attributes
	Attributes map[string]Value `json:"attributes,omitempty"`

	// Required. The timestamp of the start of the run.
	// +kcc:proto:field=google.cloud.datacatalog.lineage.v1.Run.start_time
	// +required
	StartTime *string `json:"startTime,omitempty"`

	// Optional. The timestamp of the end of the run.
	// +kcc:proto:field=google.cloud.datacatalog.lineage.v1.Run.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Required. The state of the run.
	// +kcc:proto:field=google.cloud.datacatalog.lineage.v1.Run.state
	// +required
	State *string `json:"state,omitempty"`
}

// DatalineageRunStatus defines the config connector machine state of DatalineageRun
type DatalineageRunStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DatalineageRun resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *DatalineageRunObservedState `json:"observedState,omitempty"`
}

// DatalineageRunObservedState is the state of the DatalineageRun resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.datacatalog.lineage.v1.Run
// type DatalineageRunObservedState struct {
// }

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatalineagerun;gcpdatalineageruns
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DatalineageRun is the Schema for the DatalineageRun API
// +k8s:openapi-gen=true
type DatalineageRun struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DatalineageRunSpec   `json:"spec,omitempty"`
	Status DatalineageRunStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DatalineageRunList contains a list of DatalineageRun
type DatalineageRunList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatalineageRun `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatalineageRun{}, &DatalineageRunList{})
}
