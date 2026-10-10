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

// API sources for SpannerInstancePartition, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/spanner/admin/instance/v1/spanner_instance_admin.proto
// +kcc:source:service-docs=https://cloud.google.com/spanner/
// +kcc:source:resource-docs=https://docs.cloud.google.com/spanner/docs/reference/rest/v1/projects.instances.instancePartitions

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SpannerInstancePartitionGVK = GroupVersion.WithKind("SpannerInstancePartition")

// SpannerInstancePartitionSpec defines the desired state of SpannerInstancePartition
// +kcc:spec:proto=google.spanner.admin.database.v1.InstancePartition
// +kcc:required-from-proto
type SpannerInstancePartitionSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The SpannerInstancePartition name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The name of the instance partition's configuration. Values are of
	//  the form `projects/<project>/instanceConfigs/<configuration>`. See also
	//  [InstanceConfig][google.spanner.admin.instance.v1.InstanceConfig] and
	//  [ListInstanceConfigs][google.spanner.admin.instance.v1.InstanceAdmin.ListInstanceConfigs].
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstancePartition.config
	// +required
	Config *string `json:"config,omitempty"`

	// Required. The descriptive name for this instance partition as it appears in
	//  UIs. Must be unique per project and between 4 and 30 characters in length.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstancePartition.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// The number of nodes allocated to this instance partition.
	//
	//  Users can set the `node_count` field to specify the target number of
	//  nodes allocated to the instance partition.
	//
	//  This may be zero in API responses for instance partitions that are not
	//  yet in state `READY`.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstancePartition.node_count
	NodeCount *int32 `json:"nodeCount,omitempty"`

	// The number of processing units allocated to this instance partition.
	//
	//  Users can set the `processing_units` field to specify the target number
	//  of processing units allocated to the instance partition.
	//
	//  This might be zero in API responses for instance partitions that are not
	//  yet in the `READY` state.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstancePartition.processing_units
	ProcessingUnits *int32 `json:"processingUnits,omitempty"`

	// Used for optimistic concurrency control as a way
	//  to help prevent simultaneous updates of a instance partition from
	//  overwriting each other. It is strongly suggested that systems make use of
	//  the etag in the read-modify-write cycle to perform instance partition
	//  updates in order to avoid race conditions: An etag is returned in the
	//  response which contains instance partitions, and systems are expected to
	//  put that etag in the request to update instance partitions to ensure that
	//  their change will be applied to the same version of the instance partition.
	//  If no etag is provided in the call to update instance partition, then the
	//  existing instance partition is overwritten blindly.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstancePartition.etag
	Etag *string `json:"etag,omitempty"`
}

// SpannerInstancePartitionStatus defines the config connector machine state of SpannerInstancePartition
type SpannerInstancePartitionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SpannerInstancePartition resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SpannerInstancePartitionObservedState `json:"observedState,omitempty"`
}

// SpannerInstancePartitionObservedState is the state of the SpannerInstancePartition resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.spanner.admin.database.v1.InstancePartition
type SpannerInstancePartitionObservedState struct {
	// Output only. The current instance partition state.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstancePartition.state
	State *string `json:"state,omitempty"`

	// Output only. The time at which the instance partition was created.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstancePartition.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the instance partition was most recently
	//  updated.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstancePartition.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The names of the databases that reference this
	//  instance partition. Referencing databases should share the parent instance.
	//  The existence of any referencing database prevents the instance partition
	//  from being deleted.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstancePartition.referencing_databases
	ReferencingDatabases []string `json:"referencingDatabases,omitempty"`

	// Output only. Deprecated: This field is not populated.
	//  Output only. The names of the backups that reference this instance
	//  partition. Referencing backups should share the parent instance. The
	//  existence of any referencing backup prevents the instance partition from
	//  being deleted.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstancePartition.referencing_backups
	ReferencingBackups []string `json:"referencingBackups,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpspannerinstancepartition;gcpspannerinstancepartitions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SpannerInstancePartition is the Schema for the SpannerInstancePartition API
// +k8s:openapi-gen=true
type SpannerInstancePartition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SpannerInstancePartitionSpec   `json:"spec,omitempty"`
	Status SpannerInstancePartitionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SpannerInstancePartitionList contains a list of SpannerInstancePartition
type SpannerInstancePartitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpannerInstancePartition `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SpannerInstancePartition{}, &SpannerInstancePartitionList{})
}
