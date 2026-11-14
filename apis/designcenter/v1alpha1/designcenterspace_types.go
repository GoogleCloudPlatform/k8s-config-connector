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

// API sources for DesignCenterSpace, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/547f8fdb4b91dbac70d540b3757520f0608ae356/google/cloud/designcenter/v1/catalog_management.proto
// +kcc:source:service-docs=https://cloud.google.com/application-design-center/docs
// +kcc:source:resource-docs=https://cloud.google.com/application-design-center/docs/reference/rest/v1/projects.locations.spaces

package v1alpha1

import (
	parent "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DesignCenterSpaceGVK = GroupVersion.WithKind("DesignCenterSpace")

// DesignCenterSpaceSpec defines the desired state of DesignCenterSpace
// +kcc:spec:proto=google.cloud.designcenter.v1.Space
// +kcc:required-from-proto
type DesignCenterSpaceSpec struct {
	*parent.ProjectAndLocationRef `json:",inline"`

	// The DesignCenterSpace name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Display name for the space.
	// +kcc:proto:field=google.cloud.designcenter.v1.Space.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description for the space.
	// +kcc:proto:field=google.cloud.designcenter.v1.Space.description
	Description *string `json:"description,omitempty"`

	// Optional. Flag to enable Google opinionated shared templates.
	// +kcc:proto:field=google.cloud.designcenter.v1.Space.enable_gcp_shared_templates
	EnableGcpSharedTemplates *bool `json:"enableGcpSharedTemplates,omitempty"`

	// Optional. An existing Google Cloud Storage bucket that you want to use
	//  instead of creating a new bucket during Application Design Center setup. If
	//  not provided, a default bucket is created during setup. The bucket must
	//  exist in the same project as the space. If the bucket name does not exist
	//  in the same project as the space, the request fails with an
	//  INVALID_ARGUMENT error. If you do not have access to the bucket, the
	//  request fails with a PERMISSION_DENIED error. Format: {$bucket_name} For
	//  example, if the Cloud Storage bucket URI is gs:\/\/{$bucket_name}, the
	//  format is {$bucket_name}.
	// +kcc:proto:field=google.cloud.designcenter.v1.Space.gcs_bucket
	GCSBucketRef *storagev1beta1.StorageBucketRef `json:"gcsBucketRef,omitempty"`

	// Optional. Input only. Immutable. Tags are key/values bound to space
	//  resource. Example:
	//  ```
	//    "123/environment": "production"
	//    "123/costCenter": "marketing"
	//  ```
	//
	//  For more information on tag creation and management, see
	//  https://cloud.google.com/resource-manager/docs/tags/tags-overview.
	// +kcc:proto:field=google.cloud.designcenter.v1.Space.tags
	Tags map[string]string `json:"tags,omitempty"`
}

// DesignCenterSpaceStatus defines the config connector machine state of DesignCenterSpace
type DesignCenterSpaceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DesignCenterSpace resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DesignCenterSpaceObservedState `json:"observedState,omitempty"`
}

// DesignCenterSpaceObservedState is the state of the DesignCenterSpace resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.designcenter.v1.Space
type DesignCenterSpaceObservedState struct {
	// Output only. Space creation timestamp.
	// +kcc:proto:field=google.cloud.designcenter.v1.Space.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Space update timestamp
	// +kcc:proto:field=google.cloud.designcenter.v1.Space.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdesigncenterspace;gcpdesigncenterspaces
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DesignCenterSpace is the Schema for the DesignCenterSpace API
// +k8s:openapi-gen=true
type DesignCenterSpace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DesignCenterSpaceSpec   `json:"spec,omitempty"`
	Status DesignCenterSpaceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DesignCenterSpaceList contains a list of DesignCenterSpace
type DesignCenterSpaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DesignCenterSpace `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DesignCenterSpace{}, &DesignCenterSpaceList{})
}
