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
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/kccscheme"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var LiveStreamAssetGVK = GroupVersion.WithKind("LiveStreamAsset")

// LiveStreamAssetSpec defines the desired state of LiveStreamAsset
// +kcc:spec:proto=google.cloud.video.livestream.v1.Asset
type LiveStreamAssetSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The LiveStreamAsset name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// User-defined key/value metadata.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Asset.labels
	Labels map[string]string `json:"labels,omitempty"`

	// VideoAsset represents a video.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Asset.video
	Video *AssetVideoAsset `json:"video,omitempty"`

	// ImageAsset represents an image.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Asset.image
	Image *AssetImageAsset `json:"image,omitempty"`

	// Based64-encoded CRC32c checksum of the asset file. For more information,
	//  see the crc32c checksum of the [Cloud Storage Objects
	//  resource](https://cloud.google.com/storage/docs/json_api/v1/objects).
	//  If crc32c is omitted or left empty when the asset is created, this field is
	//  filled by the crc32c checksum of the Cloud Storage object indicated by
	//  [VideoAsset.uri][google.cloud.video.livestream.v1.Asset.VideoAsset.uri] or
	//  [ImageAsset.uri][google.cloud.video.livestream.v1.Asset.ImageAsset.uri]. If
	//  crc32c is set, the asset can't be created if the crc32c value does not
	//  match with the crc32c checksum of the Cloud Storage object indicated by
	//  [VideoAsset.uri][google.cloud.video.livestream.v1.Asset.VideoAsset.uri] or
	//  [ImageAsset.uri][google.cloud.video.livestream.v1.Asset.ImageAsset.uri].
	// +kcc:proto:field=google.cloud.video.livestream.v1.Asset.crc32c
	Crc32c *string `json:"crc32c,omitempty"`
}

// LiveStreamAssetStatus defines the config connector machine state of LiveStreamAsset
type LiveStreamAssetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the LiveStreamAsset resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *LiveStreamAssetObservedState `json:"observedState,omitempty"`
}

// LiveStreamAssetObservedState is the state of the LiveStreamAsset resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.video.livestream.v1.Asset
type LiveStreamAssetObservedState struct {
	// Output only. The creation time.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Asset.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update time.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Asset.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The state of the asset resource.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Asset.state
	State *string `json:"state,omitempty"`

	// Output only. Only present when `state` is `ERROR`. The reason for the error
	//  state of the asset.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Asset.error
	Error *common.Status `json:"error,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcplivestreamasset;gcplivestreamassets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// LiveStreamAsset is the Schema for the LiveStreamAsset API
// +k8s:openapi-gen=true
type LiveStreamAsset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   LiveStreamAssetSpec   `json:"spec,omitempty"`
	Status LiveStreamAssetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// LiveStreamAssetList contains a list of LiveStreamAsset
type LiveStreamAssetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LiveStreamAsset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LiveStreamAsset{}, &LiveStreamAssetList{})
	kccscheme.RegisterType(LiveStreamAssetGVK, &LiveStreamAsset{})
}

// +kcc:proto=google.cloud.video.livestream.v1.Asset.ImageAsset
type AssetImageAsset struct {
	// Cloud Storage URI of the image. The format is `gs://my-bucket/my-object`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Asset.ImageAsset.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Asset.VideoAsset
type AssetVideoAsset struct {
	// Cloud Storage URI of the video. The format is `gs://my-bucket/my-object`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Asset.VideoAsset.uri
	URI *string `json:"uri,omitempty"`
}
