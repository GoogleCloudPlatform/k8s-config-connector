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

var _ = apiextensionsv1.JSON{}

var VisionAIApplicationGVK = GroupVersion.WithKind("VisionAIApplication")

// VisionAIApplicationSpec defines the desired state of VisionAIApplication
// +kcc:spec:proto=google.cloud.visionai.v1.Application
type VisionAIApplicationSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The VisionAIApplication name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. A user friendly display name for the solution.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName,omitempty"`

	// A description for this application.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`

	// Application graph configuration.
	// +kubebuilder:validation:Optional
	ApplicationConfigs *ApplicationConfigs `json:"applicationConfigs,omitempty"`

	// Billing mode of the application.
	// +kubebuilder:validation:Optional
	BillingMode *string `json:"billingMode,omitempty"`
}

// +kubebuilder:validation:XPreserveUnknownFields
// +kcc:proto=google.cloud.visionai.v1.AIEnabledDevicesInputConfig
type AiEnabledDevicesInputConfig struct {
}

// +kubebuilder:validation:XPreserveUnknownFields
// +kcc:proto=google.cloud.visionai.v1.GeneralObjectDetectionConfig
type GeneralObjectDetectionConfig struct {
}

// +kubebuilder:validation:XPreserveUnknownFields
// +kcc:proto=google.cloud.visionai.v1.UniversalInputConfig
type UniversalInputConfig struct {
}

// VisionAIApplicationStatus defines the config connector machine state of VisionAIApplication
type VisionAIApplicationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VisionAIApplication resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VisionAIApplicationObservedState `json:"observedState,omitempty"`
}

// VisionAIApplicationObservedState is the state of the VisionAIApplication resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.visionai.v1.Application
type VisionAIApplicationObservedState struct {
	// Output only. Create timestamp
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update timestamp
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Application graph runtime info. Only exists when application state equals to DEPLOYED.
	RuntimeInfo *Application_ApplicationRuntimeInfo `json:"runtimeInfo,omitempty"`

	// Output only. State of the application.
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvisionaiapplication;gcpvisionaiapplications
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VisionAIApplication is the Schema for the VisionAIApplication API
// +k8s:openapi-gen=true
type VisionAIApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VisionAIApplicationSpec   `json:"spec,omitempty"`
	Status VisionAIApplicationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VisionAIApplicationList contains a list of VisionAIApplication
type VisionAIApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VisionAIApplication `json:"items"`
}

type VisionAICorpusRef struct {
	// A reference to an externally managed VisionAI Corpus resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/corpora/{{corpusID}}".
	External string `json:"external,omitempty"`

	// The name of a VisionAICorpus resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VisionAICorpus resource.
	Namespace string `json:"namespace,omitempty"`
}

// +kubebuilder:validation:XPreserveUnknownFields
// +kcc:proto=google.cloud.visionai.v1.MediaWarehouseConfig
type MediaWarehouseConfig struct {
	// Resource name of the Media Warehouse corpus.
	// +kcc:proto:field=google.cloud.visionai.v1.MediaWarehouseConfig.corpus
	CorpusRef *VisionAICorpusRef `json:"corpusRef,omitempty"`

	// Deprecated.
	// +kcc:proto:field=google.cloud.visionai.v1.MediaWarehouseConfig.region
	Region *string `json:"region,omitempty"`

	// The duration for which all media assets, associated metadata, and search
	//  documents can exist.
	// +kcc:proto:field=google.cloud.visionai.v1.MediaWarehouseConfig.ttl
	TTL *string `json:"ttl,omitempty"`
}

func init() {
	SchemeBuilder.Register(&VisionAIApplication{}, &VisionAIApplicationList{})
}
