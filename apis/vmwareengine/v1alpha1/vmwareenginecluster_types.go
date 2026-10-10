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

// API sources for VMwareEngineCluster, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/vmwareengine/v1/vmwareengine_resources.proto
// +kcc:source:service-docs=https://cloud.google.com/solutions/vmware-as-a-service
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/solutions/vmware-as-a-service/docs/reference/rest/v1/projects.locations.privateClouds.clusters

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VMwareEngineClusterGVK = GroupVersion.WithKind("VMwareEngineCluster")

// VMwareEngineClusterSpec defines the desired state of VMwareEngineCluster
// +kcc:spec:proto=google.cloud.vmwareengine.v1.Cluster
// +kcc:required-from-proto
type VMwareEngineClusterSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/privateClouds/{private_cloud} this resource belongs to.
	// +kcc:guess
	// PrivateCloudRef *PrivateCloudRef `json:"privateCloudRef,omitempty"`

	// The VMwareEngineCluster name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Configuration of the autoscaling applied to this cluster.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Cluster.autoscaling_settings
	AutoscalingSettings *AutoscalingSettings `json:"autoscalingSettings,omitempty"`

	// Required. The map of cluster node types in this cluster, where the key is
	//  canonical identifier of the node type (corresponds to the `NodeType`).
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Cluster.node_type_configs
	// +required
	NodeTypeConfigs map[string]NodeTypeConfig `json:"nodeTypeConfigs,omitempty"`

	// Optional. Configuration of a stretched cluster. Required for clusters that
	//  belong to a STRETCHED private cloud.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Cluster.stretched_cluster_config
	StretchedClusterConfig *StretchedClusterConfig `json:"stretchedClusterConfig,omitempty"`
}

// VMwareEngineClusterStatus defines the config connector machine state of VMwareEngineCluster
type VMwareEngineClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VMwareEngineCluster resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VMwareEngineClusterObservedState `json:"observedState,omitempty"`
}

// VMwareEngineClusterObservedState is the state of the VMwareEngineCluster resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.vmwareengine.v1.Cluster
type VMwareEngineClusterObservedState struct {
	// Output only. Creation time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Cluster.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Cluster.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. State of the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Cluster.state
	State *string `json:"state,omitempty"`

	// Output only. True if the cluster is a management cluster; false otherwise.
	//  There can only be one management cluster in a private cloud
	//  and it has to be the first one.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Cluster.management
	Management *bool `json:"management,omitempty"`

	// Output only. System-generated unique identifier for the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Cluster.uid
	Uid *string `json:"uid,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvmwareenginecluster;gcpvmwareengineclusters
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VMwareEngineCluster is the Schema for the VMwareEngineCluster API
// +k8s:openapi-gen=true
type VMwareEngineCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VMwareEngineClusterSpec   `json:"spec,omitempty"`
	Status VMwareEngineClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VMwareEngineClusterList contains a list of VMwareEngineCluster
type VMwareEngineClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VMwareEngineCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VMwareEngineCluster{}, &VMwareEngineClusterList{})
}
