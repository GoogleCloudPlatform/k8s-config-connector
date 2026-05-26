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
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataLineageProcessGVK = GroupVersion.WithKind("DataLineageProcess")

// +kcc:proto=google.cloud.datacatalog.lineage.v1.Origin
type DataLineageProcess_Origin struct {
	// Type of the source.
	//
	//  Use of a sourceType other than CUSTOM for process creation
	//  or updating is highly discouraged. It might be restricted in the future
	//  without notice. There will be increase in cost if you use any of the
	//  source types other than CUSTOM.
	// +kcc:proto:field=google.cloud.datacatalog.lineage.v1.Origin.source_type
	SourceType *string `json:"sourceType,omitempty"`

	// If the source_type isn't CUSTOM, the value of this field should be a GCP
	//  resource name of the system, which reports lineage. The project and
	//  location parts of the resource name must match the project and location of
	//  the lineage resource being created. Examples:
	//
	//  - `{source_type: COMPOSER, name:
	//    "projects/foo/locations/us/environments/bar"}`
	//  - `{source_type: BIGQUERY, name: "projects/foo/locations/eu"}`
	//  - `{source_type: CUSTOM,   name: "myCustomIntegration"}`
	// +kcc:proto:field=google.cloud.datacatalog.lineage.v1.Origin.name
	Name *string `json:"name,omitempty"`
}

// DataLineageProcessSpec defines the desired state of DataLineageProcess
// +kcc:spec:proto=google.cloud.datacatalog.lineage.v1.Process
type DataLineageProcessSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location,omitempty"`

	// The DataLineageProcess name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. A human-readable name you can set to display in a user interface.
	//  Must be not longer than 200 characters and only contain UTF-8 letters
	//  or numbers, spaces or characters like `_-:&.`
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The attributes of the process. Should only be used for the purpose of
	//  non-semantic management (classifying, describing or labeling the process).
	//
	//  Up to 100 attributes are allowed.
	// +optional
	Attributes map[string]apiextensionsv1.JSON `json:"attributes,omitempty"`

	// Optional. The origin of this process and its runs and lineage events.
	// +optional
	Origin *DataLineageProcess_Origin `json:"origin,omitempty"`
}

// DataLineageProcessStatus defines the config connector machine state of DataLineageProcess
type DataLineageProcessStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataLineageProcess resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataLineageProcessObservedState `json:"observedState,omitempty"`
}

// DataLineageProcessObservedState is the state of the DataLineageProcess resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.datacatalog.lineage.v1.Process
type DataLineageProcessObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatalineageprocess;gcpdatalineageprocesss
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataLineageProcess is the Schema for the DataLineageProcess API
// +k8s:openapi-gen=true
type DataLineageProcess struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataLineageProcessSpec   `json:"spec,omitempty"`
	Status DataLineageProcessStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataLineageProcessList contains a list of DataLineageProcess
type DataLineageProcessList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataLineageProcess `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataLineageProcess{}, &DataLineageProcessList{})
}
