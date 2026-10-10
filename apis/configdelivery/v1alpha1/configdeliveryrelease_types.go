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

// API sources for ConfigDeliveryRelease, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/configdelivery/v1/config_delivery.proto
// +kcc:source:service-docs=https://cloud.google.com/kubernetes-engine/enterprise/config-sync/docs/concepts/fleet-packages
// +kcc:source:resource-docs=https://docs.cloud.google.com/kubernetes-engine/enterprise/config-sync/docs/reference/rest/v1/projects.locations.resourceBundles.releases

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ConfigDeliveryReleaseGVK = GroupVersion.WithKind("ConfigDeliveryRelease")

// ConfigDeliveryReleaseSpec defines the desired state of ConfigDeliveryRelease
// +kcc:spec:proto=google.cloud.configdelivery.v1.Release
// +kcc:required-from-proto
type ConfigDeliveryReleaseSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/resourceBundles/{resource_bundle} this resource belongs to.
	// +kcc:guess
	// ResourceBundleRef *ConfigDeliveryResourceBundleRef `json:"resourceBundleRef,omitempty"`

	// The ConfigDeliveryRelease name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Labels as key value pairs.
	// +kcc:proto:field=google.cloud.configdelivery.v1.Release.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. lifecycle of the `Release`.
	// +kcc:proto:field=google.cloud.configdelivery.v1.Release.lifecycle
	Lifecycle *string `json:"lifecycle,omitempty"`

	// Required. version of the `Release`. This must be v<major>.<minor>.<patch>.
	// +kcc:proto:field=google.cloud.configdelivery.v1.Release.version
	// +required
	Version *string `json:"version,omitempty"`

	// Optional. `ResourceBundle` Release extra information e.g., artifact
	//  registry image path.
	// +kcc:proto:field=google.cloud.configdelivery.v1.Release.info
	Info *ReleaseInfo `json:"info,omitempty"`
}

// ConfigDeliveryReleaseStatus defines the config connector machine state of ConfigDeliveryRelease
type ConfigDeliveryReleaseStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ConfigDeliveryRelease resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ConfigDeliveryReleaseObservedState `json:"observedState,omitempty"`
}

// ConfigDeliveryReleaseObservedState is the state of the ConfigDeliveryRelease resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.configdelivery.v1.Release
type ConfigDeliveryReleaseObservedState struct {
	// Output only. Time `Release` was created.
	// +kcc:proto:field=google.cloud.configdelivery.v1.Release.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time `Release` was last updated.
	// +kcc:proto:field=google.cloud.configdelivery.v1.Release.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Time the `Release` was published.
	// +kcc:proto:field=google.cloud.configdelivery.v1.Release.publish_time
	PublishTime *string `json:"publishTime,omitempty"`

	// Optional. `ResourceBundle` Release extra information e.g., artifact
	//  registry image path.
	// +kcc:proto:field=google.cloud.configdelivery.v1.Release.info
	Info *ReleaseInfoObservedState `json:"info,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpconfigdeliveryrelease;gcpconfigdeliveryreleases
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ConfigDeliveryRelease is the Schema for the ConfigDeliveryRelease API
// +k8s:openapi-gen=true
type ConfigDeliveryRelease struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ConfigDeliveryReleaseSpec   `json:"spec,omitempty"`
	Status ConfigDeliveryReleaseStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ConfigDeliveryReleaseList contains a list of ConfigDeliveryRelease
type ConfigDeliveryReleaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigDeliveryRelease `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConfigDeliveryRelease{}, &ConfigDeliveryReleaseList{})
}
