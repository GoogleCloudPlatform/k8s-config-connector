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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DeployDeliveryPipelineGVK = GroupVersion.WithKind("DeployDeliveryPipeline")

// DeployDeliveryPipelineSpec defines the desired state of DeployDeliveryPipeline
// +kcc:proto=google.cloud.deploy.v1.DeliveryPipeline
type DeployDeliveryPipelineSpec struct {
	// The DeployDeliveryPipeline name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Name of the `DeliveryPipeline`. Format is
	//  `projects/{project}/locations/{location}/deliveryPipelines/{deliveryPipeline}`.
	//  The `deliveryPipeline` component must match
	//  `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipeline.name
	Name *string `json:"name,omitempty"`

	// Description of the `DeliveryPipeline`. Max length is 255 characters.
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipeline.description
	Description *string `json:"description,omitempty"`

	// User annotations. These attributes can only be set and used by the
	//  user, and not by Cloud Deploy.
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipeline.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// NOT YET
	// // Labels are attributes that can be set and used by both the
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
	// // +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipeline.labels
	// Labels map[string]string `json:"labels,omitempty"`

	// SerialPipeline defines a sequential set of stages for a
	//  `DeliveryPipeline`.
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipeline.serial_pipeline
	SerialPipeline *SerialPipeline `json:"serialPipeline,omitempty"`

	// This checksum is computed by the server based on the value of other
	//  fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipeline.etag
	Etag *string `json:"etag,omitempty"`

	// When suspended, no new releases or rollouts can be created,
	//  but in-progress ones will complete.
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipeline.suspended
	Suspended *bool `json:"suspended,omitempty"`
}

// DeployDeliveryPipelineStatus defines the config connector machine state of DeployDeliveryPipeline
type DeployDeliveryPipelineStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DeployDeliveryPipeline resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DeployDeliveryPipelineObservedState `json:"observedState,omitempty"`
}

// DeployDeliveryPipelineObservedState is the state of the DeployDeliveryPipeline resource as most recently observed in GCP.
// +kcc:proto=google.cloud.deploy.v1.DeliveryPipeline
type DeployDeliveryPipelineObservedState struct {
	// Output only. Unique identifier of the `DeliveryPipeline`.
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipeline.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Time at which the pipeline was created.
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipeline.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Most recent time at which the pipeline was updated.
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipeline.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Information around the state of the Delivery Pipeline.
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipeline.condition
	Condition *PipelineCondition `json:"condition,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpdeploydeliverypipeline;gcpdeploydeliverypipelines
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DeployDeliveryPipeline is the Schema for the DeployDeliveryPipeline API
// +k8s:openapi-gen=true
type DeployDeliveryPipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DeployDeliveryPipelineSpec   `json:"spec,omitempty"`
	Status DeployDeliveryPipelineStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DeployDeliveryPipelineList contains a list of DeployDeliveryPipeline
type DeployDeliveryPipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeployDeliveryPipeline `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DeployDeliveryPipeline{}, &DeployDeliveryPipelineList{})
}
