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
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VertexAISpecialistPoolGVK = GroupVersion.WithKind("VertexAISpecialistPool")

// VertexAISpecialistPoolSpec defines the desired state of VertexAISpecialistPool
// +kcc:spec:proto=google.cloud.aiplatform.v1.SpecialistPool
type VertexAISpecialistPoolSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The VertexAISpecialistPool name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The user-defined name of the SpecialistPool.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	//  This field should be unique on project-level.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SpecialistPool.display_name
	// +required
	DisplayName *string `json:"displayName"`

	// The email addresses of the managers in the SpecialistPool.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SpecialistPool.specialist_manager_emails
	SpecialistManagerEmails []string `json:"specialistManagerEmails,omitempty"`

	// The email addresses of workers in the SpecialistPool.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SpecialistPool.specialist_worker_emails
	SpecialistWorkerEmails []string `json:"specialistWorkerEmails,omitempty"`
}

// VertexAISpecialistPoolStatus defines the config connector machine state of VertexAISpecialistPool
type VertexAISpecialistPoolStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAISpecialistPool resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAISpecialistPoolObservedState `json:"observedState,omitempty"`
}

// VertexAISpecialistPoolObservedState is the state of the VertexAISpecialistPool resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.SpecialistPool
type VertexAISpecialistPoolObservedState struct {
	// Output only. The number of managers in this SpecialistPool.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SpecialistPool.specialist_managers_count
	SpecialistManagersCount *int32 `json:"specialistManagersCount,omitempty"`

	// Output only. The resource name of the pending data labeling jobs.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SpecialistPool.pending_data_labeling_jobs
	PendingDataLabelingJobs []string `json:"pendingDataLabelingJobs,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaispecialistpool;gcpvertexaispecialistpools
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAISpecialistPool is the Schema for the VertexAISpecialistPool API
// +k8s:openapi-gen=true
type VertexAISpecialistPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAISpecialistPoolSpec   `json:"spec,omitempty"`
	Status VertexAISpecialistPoolStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAISpecialistPoolList contains a list of VertexAISpecialistPool
type VertexAISpecialistPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAISpecialistPool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAISpecialistPool{}, &VertexAISpecialistPoolList{})
}

// Declaring dummy variable to keep the unused import of apiextensionsv1 if types.generated.go is compiled separately.
var _ = apiextensionsv1.JSON{}
