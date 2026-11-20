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

// API sources for LustreMirror, recorded by generate-types:
// +kcc:source:proto=https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/mockgcp/apis/google/cloud/lustre/v1/mirror.proto
// +kcc:source:service-docs=https://cloud.google.com/managed-lustre/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/managed-lustre/docs/reference/rest/v1/projects.locations.instances.mirrors

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var LustreMirrorGVK = GroupVersion.WithKind("LustreMirror")

// LustreMirrorSpec defines the desired state of LustreMirror
// +kcc:spec:proto=google.cloud.lustre.v1.Mirror
// +kcc:required-from-proto
type LustreMirrorSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/instances/{instance} this resource belongs to.
	// +kcc:guess
	// InstanceRef *LustreInstanceRef `json:"instanceRef,omitempty"`

	// The LustreMirror name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Immutable. The URI to a Cloud Storage bucket, or a path within a
	//  bucket, using the format `gs://{BUCKET_NAME}/{OPTIONAL_PATH}/`. If a path
	//  inside the bucket is specified, it must end with a forward slash (`/`).
	// +kcc:proto:field=google.cloud.lustre.v1.Mirror.gcs_path
	// +required
	GCSPath *GCSPath `json:"gcsPath,omitempty"`

	// Required. Immutable. The Managed Lustre directory to mirror to. Must be an
	//  absolute path starting with `/`, for example `/data` or `/data/subdir`.
	//  Defaults to the root directory, `/`. If the specified directory doesn't
	//  exist, it is created.
	// +kcc:proto:field=google.cloud.lustre.v1.Mirror.lustre_path
	// +required
	LustrePath *LustrePath `json:"lustrePath,omitempty"`

	// Required. Immutable. Represents the direction of the mirror.
	// +kcc:proto:field=google.cloud.lustre.v1.Mirror.direction
	// +required
	Direction *string `json:"direction,omitempty"`

	// Optional. If `true`, files are not deleted from Managed Lustre when the
	//  source files are deleted from Cloud Storage. Default is `false`.
	// +kcc:proto:field=google.cloud.lustre.v1.Mirror.deleted_files_retained
	DeletedFilesRetained *bool `json:"deletedFilesRetained,omitempty"`

	// Optional. Description of the mirror.
	// +kcc:proto:field=google.cloud.lustre.v1.Mirror.description
	Description *string `json:"description,omitempty"`

	// Optional. Labels to apply to the mirror.
	// +kcc:proto:field=google.cloud.lustre.v1.Mirror.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// LustreMirrorStatus defines the config connector machine state of LustreMirror
type LustreMirrorStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the LustreMirror resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *LustreMirrorObservedState `json:"observedState,omitempty"`
}

// LustreMirrorObservedState is the state of the LustreMirror resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.lustre.v1.Mirror
type LustreMirrorObservedState struct {
	// Output only. [Output only] Create time stamp.
	// +kcc:proto:field=google.cloud.lustre.v1.Mirror.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. [Output only] Update time stamp.
	// +kcc:proto:field=google.cloud.lustre.v1.Mirror.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Unique ID of the resource.
	// +kcc:proto:field=google.cloud.lustre.v1.Mirror.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. [Output only] The current state of the mirror.
	// +kcc:proto:field=google.cloud.lustre.v1.Mirror.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcplustremirror;gcplustremirrors
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// LustreMirror is the Schema for the LustreMirror API
// +k8s:openapi-gen=true
type LustreMirror struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   LustreMirrorSpec   `json:"spec,omitempty"`
	Status LustreMirrorStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// LustreMirrorList contains a list of LustreMirror
type LustreMirrorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LustreMirror `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LustreMirror{}, &LustreMirrorList{})
}
