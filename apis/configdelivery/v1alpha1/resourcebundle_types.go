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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ConfigDeliveryResourceBundleGVK = GroupVersion.WithKind("ConfigDeliveryResourceBundle")

// ConfigDeliveryResourceBundleSpec defines the desired state of ConfigDeliveryResourceBundle
// +kcc:spec:proto=google.cloud.configdelivery.v1.ResourceBundle
type ConfigDeliveryResourceBundleSpec struct {
	// Required. Defines the parent path of the resource.
	*parent.ProjectAndLocationRef `json:",inline"`

	// The ConfigDeliveryResourceBundle name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Human readable description of the `ResourceBundle`.
	// +kcc:proto:field=google.cloud.configdelivery.v1.ResourceBundle.description
	Description *string `json:"description,omitempty"`

	// Optional. Labels as key value pairs.
	// +kcc:proto:field=google.cloud.configdelivery.v1.ResourceBundle.labels
	// Labels map[string]string `json:"labels,omitempty"`
}

// ConfigDeliveryResourceBundleStatus defines the config connector machine state of ConfigDeliveryResourceBundle
type ConfigDeliveryResourceBundleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ConfigDeliveryResourceBundle resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ConfigDeliveryResourceBundleObservedState `json:"observedState,omitempty"`
}

// ConfigDeliveryResourceBundleObservedState is the state of the ConfigDeliveryResourceBundle resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.configdelivery.v1.ResourceBundle
type ConfigDeliveryResourceBundleObservedState struct {
	// Output only. Time `ResourceBundle` was created.
	// +kcc:proto:field=google.cloud.configdelivery.v1.ResourceBundle.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time `ResourceBundle` was last updated.
	// +kcc:proto:field=google.cloud.configdelivery.v1.ResourceBundle.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpconfigdeliveryresourcebundle;gcpconfigdeliveryresourcebundles
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ConfigDeliveryResourceBundle is the Schema for the ConfigDeliveryResourceBundle API
// +k8s:openapi-gen=true
type ConfigDeliveryResourceBundle struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ConfigDeliveryResourceBundleSpec   `json:"spec,omitempty"`
	Status ConfigDeliveryResourceBundleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ConfigDeliveryResourceBundleList contains a list of ConfigDeliveryResourceBundle
type ConfigDeliveryResourceBundleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigDeliveryResourceBundle `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConfigDeliveryResourceBundle{}, &ConfigDeliveryResourceBundleList{})
}
