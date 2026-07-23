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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var HypercomputeClusterClusterGVK = GroupVersion.WithKind("HypercomputeClusterCluster")

// HypercomputeClusterClusterSpec defines the desired state of HypercomputeClusterCluster
// +kcc:spec:proto=google.cloud.hypercomputecluster.v1.Cluster
type HypercomputeClusterClusterSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The HypercomputeClusterCluster name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. User-provided description of the cluster.
	Description *string `json:"description,omitempty"`

	// Optional. Labels applied to the cluster.
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Network resources available to the cluster.
	NetworkResources map[string]NetworkResource `json:"networkResources,omitempty"`

	// Optional. Storage resources available to the cluster.
	StorageResources map[string]StorageResource `json:"storageResources,omitempty"`

	// Optional. Compute resources available to the cluster.
	ComputeResources map[string]ComputeResource `json:"computeResources,omitempty"`

	// Optional. Orchestrator that is responsible for scheduling and running jobs on the cluster.
	Orchestrator *Orchestrator `json:"orchestrator,omitempty"`
}

// HypercomputeClusterClusterStatus defines the config connector machine state of HypercomputeClusterCluster
type HypercomputeClusterClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the HypercomputeClusterCluster resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *HypercomputeClusterClusterObservedState `json:"observedState,omitempty"`
}

// HypercomputeClusterClusterObservedState is the state of the HypercomputeClusterCluster resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.hypercomputecluster.v1.Cluster
type HypercomputeClusterClusterObservedState struct {
	// Output only. Time that the cluster was originally created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time that the cluster was most recently updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Indicates whether changes to the cluster are currently in
	//  flight. If this is `true`, then the current state might not match the
	//  cluster's intended state.
	Reconciling *bool `json:"reconciling,omitempty"`

	// Optional. Orchestrator that is responsible for scheduling and running jobs
	//  on the cluster.
	Orchestrator *OrchestratorObservedState `json:"orchestrator,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcphypercomputeclustercluster;gcphypercomputeclusterclusters
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// HypercomputeClusterCluster is the Schema for the HypercomputeClusterCluster API
// +k8s:openapi-gen=true
type HypercomputeClusterCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   HypercomputeClusterClusterSpec   `json:"spec,omitempty"`
	Status HypercomputeClusterClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// HypercomputeClusterClusterList contains a list of HypercomputeClusterCluster
type HypercomputeClusterClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HypercomputeClusterCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HypercomputeClusterCluster{}, &HypercomputeClusterClusterList{})
}
