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

var VertexAITrialGVK = GroupVersion.WithKind("VertexAITrial")

// VertexAITrialSpec defines the desired state of VertexAITrial
type VertexAITrialSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The parent Study of this resource.
	// +required
	StudyRef *VertexAIStudyRef `json:"studyRef"`

	// The VertexAITrial name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// VertexAITrialStatus defines the config connector machine state of VertexAITrial
type VertexAITrialStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAITrial resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAITrialObservedState `json:"observedState,omitempty"`
}

// VertexAITrialObservedState is the state of the VertexAITrial resource as most recently observed in GCP.
type VertexAITrialObservedState struct {
	// Output only. Resource name of the Trial assigned by the service.
	Name *string `json:"name,omitempty"`

	// Output only. The identifier of the Trial assigned by the service.
	ID *string `json:"id,omitempty"`

	// Output only. The detailed state of the Trial.
	State *string `json:"state,omitempty"`

	// Output only. The parameters of the Trial.
	Parameters []Trial_ParameterObservedState `json:"parameters,omitempty"`

	// Output only. The final measurement containing the objective value.
	FinalMeasurement *MeasurementObservedState `json:"finalMeasurement,omitempty"`

	// Output only. A list of measurements that are strictly lexicographically
	//  ordered by their induced tuples (steps, elapsed_duration).
	//  These are used for early stopping computations.
	Measurements []MeasurementObservedState `json:"measurements,omitempty"`

	// Output only. Time when the Trial was started.
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Time when the Trial's status changed to `SUCCEEDED` or
	//  `INFEASIBLE`.
	EndTime *string `json:"endTime,omitempty"`

	// Output only. The identifier of the client that originally requested this
	//  Trial. Each client is identified by a unique client_id. When a client asks
	//  for a suggestion, Vertex AI Vizier will assign it a Trial. The client
	//  should evaluate the Trial, complete it, and report back to Vertex AI
	//  Vizier. If suggestion is asked again by same client_id before the Trial is
	//  completed, the same Trial will be returned. Multiple clients with
	//  different client_ids can ask for suggestions simultaneously, each of them
	//  will get their own Trial.
	ClientID *string `json:"clientID,omitempty"`

	// Output only. A human readable string describing why the Trial is
	//  infeasible. This is set only if Trial state is `INFEASIBLE`.
	InfeasibleReason *string `json:"infeasibleReason,omitempty"`

	// Output only. The CustomJob name linked to the Trial.
	//  It's set for a HyperparameterTuningJob's Trial.
	CustomJob *string `json:"customJob,omitempty"`

	// Output only. URIs for accessing [interactive
	//  shells](https://cloud.google.com/vertex-ai/docs/training/monitor-debug-interactive-shell)
	//  (one URI for each training node). Only available if this trial is part of
	//  a
	//  [HyperparameterTuningJob][google.cloud.aiplatform.v1beta1.HyperparameterTuningJob]
	//  and the job's
	//  [trial_job_spec.enable_web_access][google.cloud.aiplatform.v1beta1.CustomJobSpec.enable_web_access]
	//  field is `true`.
	//
	//  The keys are names of each node used for the trial; for example,
	//  `workerpool0-0` for the primary node, `workerpool1-0` for the first node in
	//  the second worker pool, and `workerpool1-1` for the second node in the
	//  second worker pool.
	//
	//  The values are the URIs for each node's interactive shell.
	WebAccessURIs map[string]string `json:"webAccessURIs,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaitrial;gcpvertexaitrials
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAITrial is the Schema for the VertexAITrial API
// +k8s:openapi-gen=true
type VertexAITrial struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAITrialSpec   `json:"spec,omitempty"`
	Status VertexAITrialStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAITrialList contains a list of VertexAITrial
type VertexAITrialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAITrial `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAITrial{}, &VertexAITrialList{})
}
