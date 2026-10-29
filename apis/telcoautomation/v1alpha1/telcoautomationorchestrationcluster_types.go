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

// API sources for TelcoAutomationOrchestrationCluster, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/telcoautomation/v1/telcoautomation.proto
// +kcc:source:service-docs=https://cloud.google.com/telecom-network-automation
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/telecom-network-automation/docs/reference/rest/v1/projects.locations.orchestrationClusters

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var TelcoAutomationOrchestrationClusterGVK = GroupVersion.WithKind("TelcoAutomationOrchestrationCluster")

// TelcoAutomationOrchestrationClusterSpec defines the desired state of TelcoAutomationOrchestrationCluster
// +kcc:spec:proto=google.cloud.telcoautomation.v1.OrchestrationCluster
// +kcc:required-from-proto
type TelcoAutomationOrchestrationClusterSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The TelcoAutomationOrchestrationCluster name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Management configuration of the underlying GKE cluster.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.OrchestrationCluster.management_config
	ManagementConfig *ManagementConfig `json:"managementConfig,omitempty"`

	// Labels as key value pairs.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.OrchestrationCluster.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// TelcoAutomationOrchestrationClusterStatus defines the config connector machine state of TelcoAutomationOrchestrationCluster
type TelcoAutomationOrchestrationClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the TelcoAutomationOrchestrationCluster resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *TelcoAutomationOrchestrationClusterObservedState `json:"observedState,omitempty"`
}

// TelcoAutomationOrchestrationClusterObservedState is the state of the TelcoAutomationOrchestrationCluster resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.telcoautomation.v1.OrchestrationCluster
type TelcoAutomationOrchestrationClusterObservedState struct {
	// Output only. [Output only] Create time stamp.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.OrchestrationCluster.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. [Output only] Update time stamp.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.OrchestrationCluster.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Provides the TNA version installed on the cluster.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.OrchestrationCluster.tna_version
	TnaVersion *string `json:"tnaVersion,omitempty"`

	// Output only. State of the Orchestration Cluster.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.OrchestrationCluster.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcptelcoautomationorchestrationcluster;gcptelcoautomationorchestrationclusters
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// TelcoAutomationOrchestrationCluster is the Schema for the TelcoAutomationOrchestrationCluster API
// +k8s:openapi-gen=true
type TelcoAutomationOrchestrationCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   TelcoAutomationOrchestrationClusterSpec   `json:"spec,omitempty"`
	Status TelcoAutomationOrchestrationClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TelcoAutomationOrchestrationClusterList contains a list of TelcoAutomationOrchestrationCluster
type TelcoAutomationOrchestrationClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TelcoAutomationOrchestrationCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TelcoAutomationOrchestrationCluster{}, &TelcoAutomationOrchestrationClusterList{})
}
