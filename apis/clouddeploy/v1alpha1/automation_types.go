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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	commonv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/common/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudDeployAutomationGVK = GroupVersion.WithKind("CloudDeployAutomation")

// CloudDeployAutomationSpec defines the desired state of CloudDeployAutomation
// +kcc:spec:proto=google.cloud.deploy.v1.Automation
type CloudDeployAutomationSpec struct {
	commonv1alpha1.CommonSpec `json:",inline"`

	// The delivery pipeline to which this automation belongs
	// +required
	DeliveryPipelineRef DeliveryPipelineRef `json:"deliveryPipelineRef"`

	// Optional. Description of the `Automation`. Max length is 255 characters.
	Description *string `json:"description,omitempty"`

	// Optional. User-defined key/value metadata.
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. Labels are attributes that can be set and used by both the
	// user and by Cloud Deploy. It is used to select objects for filtering and
	// finding.
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Selectors contain the selection criteria for choosing the resources to
	// act on.
	Selector *AutomationResourceSelector `json:"selector,omitempty"`

	// Optional. When suspended, automation is deactivated from execution.
	Suspended *bool `json:"suspended,omitempty"`

	// Required. Email address of the user-managed IAM service account that creates the Cloud
	// Deploy release.
	ServiceAccountRef *refs.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Required. A sequence of rules that determines connections between stages in a
	// delivery pipeline.
	Rules []*AutomationRule `json:"rules,omitempty"`
}

// CloudDeployAutomationStatus defines the config connector machine state of CloudDeployAutomation
type CloudDeployAutomationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudDeployAutomation resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudDeployAutomationObservedState `json:"observedState,omitempty"`
}

// CloudDeployAutomationObservedState is the state of the CloudDeployAutomation resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.deploy.v1.Automation
type CloudDeployAutomationObservedState struct {
	// Name of the 'Automation'. Format is
	// 'projects/{project}/locations/{location}/deliveryPipelines/{delivery_pipeline}/automations/{automation}'.
	Name *string `json:"name,omitempty"`

	// Unique identifier of the 'Automation'.
	Uid *string `json:"uid,omitempty"`

	// Time at which the automation was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Time at which the automation was updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// The weak etag of the 'Automation' resource.
	// This checksum is computed by the server based on the value of other
	// fields, and may be sent on update and delete requests to ensure the
	// client has an up-to-date value before proceeding.
	Etag *string `json:"etag,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpclouddeployautomation;gcpclouddeployautomations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudDeployAutomation is the Schema for the CloudDeployAutomation API
// +k8s:openapi-gen=true
type CloudDeployAutomation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudDeployAutomationSpec   `json:"spec,omitempty"`
	Status CloudDeployAutomationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudDeployAutomationList contains a list of CloudDeployAutomation
type CloudDeployAutomationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudDeployAutomation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudDeployAutomation{}, &CloudDeployAutomationList{})
}
