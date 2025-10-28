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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataCatalogPolicyTagGVK = GroupVersion.WithKind("DataCatalogPolicyTag")

// DataCatalogPolicyTagSpec defines the desired state of DataCatalogPolicyTag
// +kcc:spec:proto=google.cloud.datacatalog.v1beta1.PolicyTag
type DataCatalogPolicyTagSpec struct {
	/* Description of this policy tag. It must: contain only unicode characters, tabs,
	newlines, carriage returns and page breaks; and be at most 2000 bytes long when
	encoded in UTF-8. If not set, defaults to an empty description.
	If not set, defaults to an empty description. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* User defined name of this policy tag. It must: be unique within the parent
	taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces;
	not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8. */
	DisplayName string `json:"displayName"`

	// +optional
	ParentPolicyTagRef *PolicyTagRef `json:"parentPolicyTagRef,omitempty"`

	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	TaxonomyRef TaxonomyRef `json:"taxonomyRef"`
}

// DataCatalogPolicyTagStatus defines the config connector machine state of DataCatalogPolicyTag
type DataCatalogPolicyTagStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataCatalogPolicyTag resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataCatalogPolicyTagObservedState `json:"observedState,omitempty"`

	/* Resource name of this policy tag, whose format is:
	"projects/{project}/locations/{region}/taxonomies/{taxonomy}/policyTags/{policytag}". */
	// +optional
	Name *string `json:"name,omitempty"`

	/* Resource names of child policy tags of this policy tag. */
	// +optional
	ChildPolicyTags []string `json:"childPolicyTags,omitempty"`
}

// DataCatalogPolicyTagObservedState is the state of the DataCatalogPolicyTag resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.datacatalog.v1beta1.PolicyTag
type DataCatalogPolicyTagObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatacatalogpolicytag;gcpdatacatalogpolicytags
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataCatalogPolicyTag is the Schema for the DataCatalogPolicyTag API
// +k8s:openapi-gen=true
type DataCatalogPolicyTag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataCatalogPolicyTagSpec   `json:"spec,omitempty"`
	Status DataCatalogPolicyTagStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataCatalogPolicyTagList contains a list of DataCatalogPolicyTag
type DataCatalogPolicyTagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataCatalogPolicyTag `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataCatalogPolicyTag{}, &DataCatalogPolicyTagList{})
}
