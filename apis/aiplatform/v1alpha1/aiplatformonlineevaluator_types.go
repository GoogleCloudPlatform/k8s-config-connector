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

var AIPlatformOnlineEvaluatorGVK = GroupVersion.WithKind("AIPlatformOnlineEvaluator")

// AIPlatformOnlineEvaluatorSpec defines the desired state of AIPlatformOnlineEvaluator
// kcc:spec:proto=google.cloud.aiplatform.v1beta1.OnlineEvaluator
type AIPlatformOnlineEvaluatorSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The AIPlatformOnlineEvaluator name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Human-readable name for the `OnlineEvaluator`.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.OnlineEvaluator.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Immutable. The name of the agent that the OnlineEvaluator evaluates periodically.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.OnlineEvaluator.agent_resource
	AgentRef *DialogflowCXAgentRef `json:"agentRef"`

	// Data source for the OnlineEvaluator, based on GCP Observability stack
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.OnlineEvaluator.cloud_observability
	CloudObservability *OnlineEvaluator_CloudObservability `json:"cloudObservability,omitempty"`

	// Required. A list of metric sources to be used for evaluating samples.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.OnlineEvaluator.metric_sources
	MetricSources []MetricSource `json:"metricSources"`

	// Required. Configuration for the OnlineEvaluator.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.OnlineEvaluator.config
	Config *OnlineEvaluator_Config `json:"config"`
}

// AIPlatformOnlineEvaluatorStatus defines the config connector machine state of AIPlatformOnlineEvaluator
type AIPlatformOnlineEvaluatorStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AIPlatformOnlineEvaluator resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AIPlatformOnlineEvaluatorObservedState `json:"observedState,omitempty"`
}

// AIPlatformOnlineEvaluatorObservedState is the state of the AIPlatformOnlineEvaluator resource as most recently observed in GCP.
// kcc:observedstate:proto=google.cloud.aiplatform.v1beta1.OnlineEvaluator
type AIPlatformOnlineEvaluatorObservedState struct {
	// Output only. The state of the OnlineEvaluator.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.OnlineEvaluator.state
	State *string `json:"state,omitempty"`

	// Output only. Contains additional information about the state of the
	//  OnlineEvaluator. This is used to provide more details in the event of a
	//  failure.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.OnlineEvaluator.state_details
	StateDetails []OnlineEvaluator_StateDetailsObservedState `json:"stateDetails,omitempty"`

	// Output only. Timestamp when the OnlineEvaluator was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.OnlineEvaluator.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when the OnlineEvaluator was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.OnlineEvaluator.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaiplatformonlineevaluator;gcpaiplatformonlineevaluators
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AIPlatformOnlineEvaluator is the Schema for the AIPlatformOnlineEvaluator API
// +k8s:openapi-gen=true
type AIPlatformOnlineEvaluator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AIPlatformOnlineEvaluatorSpec   `json:"spec,omitempty"`
	Status AIPlatformOnlineEvaluatorStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AIPlatformOnlineEvaluatorList contains a list of AIPlatformOnlineEvaluator
type AIPlatformOnlineEvaluatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AIPlatformOnlineEvaluator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AIPlatformOnlineEvaluator{}, &AIPlatformOnlineEvaluatorList{})
}
