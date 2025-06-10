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

package v1alpha1

import (
	cloudbuildv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CustomTargetTypeGVK = GroupVersion.WithKind("DeployCustomTargetType")

// CustomTargetTypeSpec defines the desired state of DeployCustomTargetType
// +kcc:spec:proto=google.cloud.deploy.v1.CustomTargetType
type CustomTargetTypeSpec struct {
	// The DeployCustomTargetType name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The project that this resource belongs to.
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`

	// +required
	Location string `json:"location"`

	// Optional. Description of the `CustomTargetType`. Max length is 255
	//  characters.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetType.description
	Description *string `json:"description,omitempty"`

	// NOT YET
	// // Optional. User annotations. These attributes can only be set and used by
	// //  the user, and not by Cloud Deploy. See
	// //  https://google.aip.dev/128#annotations for more details such as format and
	// //  size limitations.
	// // +kcc:proto:field=google.cloud.deploy.v1.CustomTargetType.annotations
	// Annotations map[string]string `json:"annotations,omitempty"`

	// NOT YET
	// // Optional. Labels are attributes that can be set and used by both the
	// //  user and by Cloud Deploy. Labels must meet the following constraints:
	// //
	// //  * Keys and values can contain only lowercase letters, numeric characters,
	// //  underscores, and dashes.
	// //  * All characters must use UTF-8 encoding, and international characters are
	// //  allowed.
	// //  * Keys must start with a lowercase letter or international character.
	// //  * Each resource is limited to a maximum of 64 labels.
	// //
	// //  Both keys and values are additionally constrained to be <= 128 bytes.
	// // +kcc:proto:field=google.cloud.deploy.v1.CustomTargetType.labels
	// Labels map[string]string `json:"labels,omitempty"`

	// // Optional. This checksum is computed by the server based on the value of
	// //  other fields, and may be sent on update and delete requests to ensure the
	// //  client has an up-to-date value before proceeding.
	// // +kcc:proto:field=google.cloud.deploy.v1.CustomTargetType.etag
	// Etag *string `json:"etag,omitempty"`

	// Configures render and deploy for the `CustomTargetType` using Skaffold
	//  custom actions.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetType.custom_actions
	CustomActions *CustomTargetSkaffoldActions `json:"customActions,omitempty"`
}

// CustomTargetTypeStatus defines the config connector machine state of DeployCustomTargetType
type CustomTargetTypeStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DeployCustomTargetType resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CustomTargetTypeObservedState `json:"observedState,omitempty"`
}

// CustomTargetTypeObservedState is the state of the DeployCustomTargetType resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.deploy.v1.CustomTargetType
type CustomTargetTypeObservedState struct {
	// Output only. Resource id of the `CustomTargetType`.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetType.custom_target_type_id
	CustomTargetTypeID *string `json:"customTargetTypeID,omitempty"`

	// Output only. Unique identifier of the `CustomTargetType`.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetType.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Time at which the `CustomTargetType` was created.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetType.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Most recent time at which the `CustomTargetType` was updated.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetType.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdeploycustomtargettype;gcpdeploycustomtargettypes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DeployCustomTargetType is the Schema for the DeployCustomTargetType API
// +k8s:openapi-gen=true
type DeployCustomTargetType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CustomTargetTypeSpec   `json:"spec,omitempty"`
	Status CustomTargetTypeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DeployCustomTargetTypeList contains a list of DeployCustomTargetType
type DeployCustomTargetTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeployCustomTargetType `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DeployCustomTargetType{}, &DeployCustomTargetTypeList{})
}

// +kcc:proto=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGCBRepoSource
type SkaffoldModules_SkaffoldGcbRepoSource struct {
	// Required. Name of the Cloud Build V2 RepositoryRef.
	//  Format is
	//  projects/{project}/locations/{location}/connections/{connection}/repositories/{repository}.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGCBRepoSource.repository
	RepositoryRef *cloudbuildv1alpha1.RepositoryRef `json:"repositoryRef,omitempty"`

	// Optional. Relative path from the repository root to the Skaffold Config
	//  file.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGCBRepoSource.path
	Path *string `json:"path,omitempty"`

	// Optional. Branch or tag to use when cloning the repository.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGCBRepoSource.ref
	Ref *string `json:"ref,omitempty"`
}
