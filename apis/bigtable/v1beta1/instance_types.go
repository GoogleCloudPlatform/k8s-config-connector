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

package v1beta1

import (
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	BigtableInstanceGVK = GroupVersion.WithKind("BigtableInstance")
)

// BigtableInstanceSpec defines the desired state of BigtableInstance
// +kcc:proto=google.bigtable.admin.v2.Instance
type BigtableInstanceSpec struct {
	// The Instance name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	//     // The unique name of the instance. Values are of the form
	// //  `projects/{project}/instances/[a-z][a-z0-9\\-]+[a-z0-9]`.
	// Name *string `json:"name,omitempty"`

	// DEPRECATED. This field no longer serves any function and is intended to be dropped in a later version of the resource.
	// +optional
	DeletionProtection *bool `json:"deletionProtection,omitempty"`

	// Required. The descriptive name for this instance as it appears in UIs.
	//  Can be changed at any time, but should be kept globally unique
	//  to avoid confusion.
	DisplayName *string `json:"displayName,omitempty"`

	// DEPRECATED. It is recommended to leave this field unspecified since the distinction between "DEVELOPMENT" and "PRODUCTION" instances is going away, and all instances will become "PRODUCTION" instances. This means that new and existing "DEVELOPMENT" instances will be converted to "PRODUCTION" instances. It is recommended for users to use "PRODUCTION" instances in any case, since a 1-node "PRODUCTION" instance is functionally identical to a "DEVELOPMENT" instance, but without the accompanying restrictions. The instance type to create. One of "DEVELOPMENT" or "PRODUCTION". Defaults to "PRODUCTION".
	// +optional
	InstanceType *string `json:"instanceType,omitempty"`

	/*NOTYET
	// Labels are a flexible and lightweight mechanism for organizing cloud
	//  resources into groups that reflect a customer's organizational needs and
	//  deployment strategies. They can be used to filter resources and aggregate
	//  metrics.
	//
	//  * Label keys must be between 1 and 63 characters long and must conform to
	//    the regular expression: `[\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}`.
	//  * Label values must be between 0 and 63 characters long and must conform to
	//    the regular expression: `[\p{Ll}\p{Lo}\p{N}_-]{0,63}`.
	//  * No more than 64 labels can be associated with a given resource.
	//  * Keys and values must both be under 128 bytes.
	Labels map[string]string `json:"labels,omitempty"`
	*/

	// A block of cluster configuration options. This can be specified at least once.
	Cluster []InstanceCluster `json:"cluster,omitempty"`
}

type InstanceAutoscalingConfig struct {
	/* The target CPU utilization for autoscaling. Value must be between 10 and 80. */
	CpuTarget int64 `json:"cpuTarget"`

	/* The maximum number of nodes for autoscaling. */
	MaxNodes int64 `json:"maxNodes"`

	/* The minimum number of nodes for autoscaling. */
	MinNodes int64 `json:"minNodes"`

	/* The target storage utilization for autoscaling, in GB, for each node in a cluster. This number is limited between 2560 (2.5TiB) and 5120 (5TiB) for a SSD cluster and between 8192 (8TiB) and 16384 (16 TiB) for an HDD cluster. If not set, whatever is already set for the cluster will not change, or if the cluster is just being created, it will use the default value of 2560 for SSD clusters and 8192 for HDD clusters. */
	// +optional
	StorageTarget *int64 `json:"storageTarget,omitempty"`
}

type InstanceCluster struct {
	/* A list of Autoscaling configurations. Only one element is used and allowed. */
	// +optional
	AutoscalingConfig *InstanceAutoscalingConfig `json:"autoscalingConfig,omitempty"`

	/* The ID of the Cloud Bigtable cluster. Must be 6-30 characters and must only contain hyphens, lowercase letters and numbers. */
	ClusterId string `json:"clusterId"`

	// Describes the Cloud KMS encryption key that will be used to protect the destination Bigtable
	// cluster. The requirements for this key are:
	//
	// 1) The Cloud Bigtable service account associated with the project that contains
	// this cluster must be granted the cloudkms.cryptoKeyEncrypterDecrypter role on the CMEK key.
	// 2) Only regional keys can be used and the region of the CMEK key must match the region of the cluster.
	// 3) All clusters within an instance must use the same CMEK key access to this encryption key.
	// +optional
	KmsKeyRef *kmsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`

	/* The number of nodes in the cluster. If no value is set, Cloud Bigtable automatically allocates nodes based on your data footprint and optimized for 50% storage utilization. */
	// +optional
	NumNodes *int64 `json:"numNodes,omitempty"`

	/* The storage type to use. One of "SSD" or "HDD". Defaults to "SSD". */
	// +optional
	StorageType *string `json:"storageType,omitempty"`

	/* The zone to create the Cloud Bigtable cluster in. Each cluster must have a different zone in the same region. Zones that support Bigtable instances are noted on the Cloud Bigtable locations page. */
	Zone string `json:"zone"`
}

// BigtableInstanceStatus defines the config connector machine state of BigtableInstance
type BigtableInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* NOTYET
	// A unique specifier for the BigtableInstance resource in GCP.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`
	*/

	/* NOTYET
	// ObservedState is the state of the resource as most recently observed in GCP.
	// +optional
	ObservedState *BigtableInstanceObservedState `json:"observedState,omitempty"`
	*/
}

/* NOTYET
// BigtableInstanceSpec defines the desired state of BigtableInstance
// +kcc:proto=google.bigtable.admin.v2.Instance
type BigtableInstanceObservedState struct {

	// (`OutputOnly`)
	//  The current state of the instance.
	State *string `json:"state,omitempty"`

	// The type of the instance. Defaults to `PRODUCTION`.
	Type *string `json:"type,omitempty"`

	// Output only. A server-assigned timestamp representing when this Instance
	//  was created. For instances created before this field was added (August
	//  2021), this value is `seconds: 0, nanos: 1`.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Reserved for future use.
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`
}
*/

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigtableinstance;gcpbigtableinstances
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/tf2crd=true"

// BigtableInstance is the Schema for the BigtableInstance API
// +k8s:openapi-gen=true
type BigtableInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigtableInstanceSpec   `json:"spec,omitempty"`
	Status BigtableInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigtableInstanceList contains a list of BigtableInstance
type BigtableInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigtableInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigtableInstance{}, &BigtableInstanceList{})
}
