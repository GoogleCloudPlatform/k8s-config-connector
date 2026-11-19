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

// API sources for BigLakeIcebergNamespace, recorded by generate-types:
// +kcc:source:proto=https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/mockgcp/apis/google/cloud/biglake/v1/iceberg_rest_catalog.proto
// +kcc:source:service-docs=https://cloud.google.com/products/lakehouse
// +kcc:source:resource-docs=https://cloud.google.com/bigquery/docs/iceberg-rest-catalog-about

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigLakeIcebergNamespaceGVK = GroupVersion.WithKind("BigLakeIcebergNamespace")

// BigLakeIcebergNamespaceSpec defines the desired state of BigLakeIcebergNamespace
// +kcc:spec:proto=google.cloud.biglake.v1.IcebergNamespace
// +kcc:required-from-proto
type BigLakeIcebergNamespaceSpec struct {
	// The BigLakeIcebergCatalog this BigLakeIcebergNamespace is created under.
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="parentRef is immutable"
	ParentRef *BigLakeIcebergCatalogRef `json:"parentRef"`

	// The BigLakeIcebergNamespace name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The name of the namespace.
	// +kcc:proto:field=google.cloud.biglake.v1.IcebergNamespace.namespace
	// +required
	Namespace []string `json:"namespace,omitempty"`

	// Optional. The optional properties of the namespace.
	// +kcc:proto:field=google.cloud.biglake.v1.IcebergNamespace.properties
	Properties map[string]string `json:"properties,omitempty"`
}

// BigLakeIcebergNamespaceStatus defines the config connector machine state of BigLakeIcebergNamespace
type BigLakeIcebergNamespaceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigLakeIcebergNamespace resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *BigLakeIcebergNamespaceObservedState `json:"observedState,omitempty"`
}

// BigLakeIcebergNamespaceObservedState is the state of the BigLakeIcebergNamespace resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.biglake.v1.IcebergNamespace
// type BigLakeIcebergNamespaceObservedState struct {
// }

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbiglakeicebergnamespace;gcpbiglakeicebergnamespaces
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigLakeIcebergNamespace is the Schema for the BigLakeIcebergNamespace API
// +k8s:openapi-gen=true
type BigLakeIcebergNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigLakeIcebergNamespaceSpec   `json:"spec,omitempty"`
	Status BigLakeIcebergNamespaceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigLakeIcebergNamespaceList contains a list of BigLakeIcebergNamespace
type BigLakeIcebergNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigLakeIcebergNamespace `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigLakeIcebergNamespace{}, &BigLakeIcebergNamespaceList{})
}
