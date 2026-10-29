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

// API sources for VisionAICluster, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/visionai/v1/common.proto
// +kcc:source:service-docs=https://cloud.google.com/vision-ai/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/vision-ai/docs/reference/rest/v1/projects.locations.clusters

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VisionAIClusterGVK = GroupVersion.WithKind("VisionAICluster")

// VisionAIClusterSpec defines the desired state of VisionAICluster
// +kcc:spec:proto=google.cloud.visionai.v1.Cluster
// +kcc:required-from-proto
type VisionAIClusterSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The VisionAICluster name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Labels as key value pairs
	// +kcc:proto:field=google.cloud.visionai.v1.Cluster.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Annotations to allow clients to store small amounts of arbitrary data.
	// +kcc:proto:field=google.cloud.visionai.v1.Cluster.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// VisionAIClusterStatus defines the config connector machine state of VisionAICluster
type VisionAIClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VisionAICluster resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VisionAIClusterObservedState `json:"observedState,omitempty"`
}

// VisionAIClusterObservedState is the state of the VisionAICluster resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.visionai.v1.Cluster
type VisionAIClusterObservedState struct {
	// Output only. The create timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.Cluster.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.Cluster.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The DNS name of the data plane service
	// +kcc:proto:field=google.cloud.visionai.v1.Cluster.dataplane_service_endpoint
	DataplaneServiceEndpoint *string `json:"dataplaneServiceEndpoint,omitempty"`

	// Output only. The current state of the cluster.
	// +kcc:proto:field=google.cloud.visionai.v1.Cluster.state
	State *string `json:"state,omitempty"`

	// Output only. The private service connection service target name.
	// +kcc:proto:field=google.cloud.visionai.v1.Cluster.psc_target
	PSCTarget *string `json:"pscTarget,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvisionaicluster;gcpvisionaiclusters
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VisionAICluster is the Schema for the VisionAICluster API
// +k8s:openapi-gen=true
type VisionAICluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VisionAIClusterSpec   `json:"spec,omitempty"`
	Status VisionAIClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VisionAIClusterList contains a list of VisionAICluster
type VisionAIClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VisionAICluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VisionAICluster{}, &VisionAIClusterList{})
}
