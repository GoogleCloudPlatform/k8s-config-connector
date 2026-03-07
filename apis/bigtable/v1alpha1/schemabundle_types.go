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
	bigtablev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigtableSchemaBundleGVK = GroupVersion.WithKind("BigtableSchemaBundle")

// BigtableSchemaBundleSpec defines the desired state of BigtableSchemaBundle
// +kcc:spec:proto=google.bigtable.admin.v2.SchemaBundle
type BigtableSchemaBundleSpec struct {
	// The table to create the schema bundle in.
	// +required
	TableRef bigtablev1beta1.TableRef `json:"tableRef"`

	// The BigtableSchemaBundle name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. The protobuf schema for the table.
	// +optional
	ProtoSchema *ProtoSchema `json:"protoSchema,omitempty"`
}

// ProtoSchema contains a protobuf-serialized FileDescriptorSet.
type ProtoSchema struct {
	// Optional. The protobuf schema descriptor.
	// The schema descriptor must be a file descriptor set, which can be generated using `protoc --descriptor_set_out=myschema.fd myschema.proto`.
	// The file descriptor set must be base64 encoded.
	// +optional
	ProtoDescriptors []byte `json:"protoDescriptors,omitempty"`
}

// BigtableSchemaBundleStatus defines the config connector machine state of BigtableSchemaBundle
type BigtableSchemaBundleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []k8sv1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigtableSchemaBundle resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigtableSchemaBundleObservedState `json:"observedState,omitempty"`
}

// BigtableSchemaBundleObservedState is the state of the BigtableSchemaBundle resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.bigtable.admin.v2.SchemaBundle
type BigtableSchemaBundleObservedState struct {
	// Optional. The etag for this schema bundle.
	Etag *string `json:"etag,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigtableschemabundle;gcpbigtableschemabundles
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:annotations="cnrm.cloud.google.com/owner=fkc1e100/Frank Currie/fcurrie"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigtableSchemaBundle is the Schema for the BigtableSchemaBundle API
// +k8s:openapi-gen=true
type BigtableSchemaBundle struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigtableSchemaBundleSpec   `json:"spec,omitempty"`
	Status BigtableSchemaBundleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigtableSchemaBundleList contains a list of BigtableSchemaBundle
type BigtableSchemaBundleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigtableSchemaBundle `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigtableSchemaBundle{}, &BigtableSchemaBundleList{})
}
