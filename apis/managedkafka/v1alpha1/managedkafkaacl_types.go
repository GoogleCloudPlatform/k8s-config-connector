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

// API sources for ManagedKafkaACL, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/managedkafka/v1/resources.proto
// +kcc:source:service-docs=https://cloud.google.com/managed-service-for-apache-kafka/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/managed-service-for-apache-kafka/docs/reference/rest/v1/projects.locations.clusters.acls

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ManagedKafkaACLGVK = GroupVersion.WithKind("ManagedKafkaACL")

// ManagedKafkaACLSpec defines the desired state of ManagedKafkaACL
// +kcc:spec:proto=google.cloud.managedkafka.v1.Acl
// +kcc:required-from-proto
type ManagedKafkaACLSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/clusters/{cluster} this resource belongs to.
	// +kcc:guess
	// ClusterRef *ClusterRef `json:"clusterRef,omitempty"`

	// The ManagedKafkaACL name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The ACL entries that apply to the resource pattern. The maximum
	//  number of allowed entries 100.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Acl.acl_entries
	// +required
	AclEntries []AclEntry `json:"aclEntries,omitempty"`

	// Optional. `etag` is used for concurrency control. An `etag` is returned in
	//  the response to `GetAcl` and `CreateAcl`. Callers are required to put that
	//  etag in the request to `UpdateAcl` to ensure that their change will be
	//  applied to the same version of the acl that exists in the Kafka Cluster.
	//
	//  A terminal 'T' character in the etag indicates that the AclEntries were
	//  truncated; more entries for the Acl exist on the Kafka Cluster, but can't
	//  be returned in the Acl due to repeated field limits.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Acl.etag
	Etag *string `json:"etag,omitempty"`
}

// ManagedKafkaACLStatus defines the config connector machine state of ManagedKafkaACL
type ManagedKafkaACLStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ManagedKafkaACL resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ManagedKafkaACLObservedState `json:"observedState,omitempty"`
}

// ManagedKafkaACLObservedState is the state of the ManagedKafkaACL resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.managedkafka.v1.Acl
type ManagedKafkaACLObservedState struct {
	// Output only. The ACL resource type derived from the name. One of: CLUSTER,
	//  TOPIC, GROUP, TRANSACTIONAL_ID.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Acl.resource_type
	ResourceType *string `json:"resourceType,omitempty"`

	// Output only. The ACL resource name derived from the name. For cluster
	//  resource_type, this is always "kafka-cluster". Can be the wildcard literal
	//  "*".
	// +kcc:proto:field=google.cloud.managedkafka.v1.Acl.resource_name
	ResourceName *string `json:"resourceName,omitempty"`

	// Output only. The ACL pattern type derived from the name. One of: LITERAL,
	//  PREFIXED.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Acl.pattern_type
	PatternType *string `json:"patternType,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmanagedkafkaacl;gcpmanagedkafkaacls
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ManagedKafkaACL is the Schema for the ManagedKafkaACL API
// +k8s:openapi-gen=true
type ManagedKafkaACL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ManagedKafkaACLSpec   `json:"spec,omitempty"`
	Status ManagedKafkaACLStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ManagedKafkaACLList contains a list of ManagedKafkaACL
type ManagedKafkaACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedKafkaACL `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ManagedKafkaACL{}, &ManagedKafkaACLList{})
}
