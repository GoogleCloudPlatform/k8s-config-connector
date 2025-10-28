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

package v1beta1

import (
	commonv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/common/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudDeployDeliveryPipelineGVK = GroupVersion.WithKind("CloudDeployDeliveryPipeline")

// DeliveryPipelineSpec defines the desired state of DeployDeliveryPipeline
// +kcc:spec:proto=google.cloud.deploy.v1.DeliveryPipeline
type DeliveryPipelineSpec struct {
	commonv1alpha1.CommonSpec `json:",inline"`

	// Immutable. The location where the DeliveryPipeline should reside.
	// +required
	Location *string `json:"location,omitempty"`

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

	// NOTYET: not supported in Config Connector reconciliation
	// Optional. The weak etag of the `DeliveryPipeline` resource.
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipeline.etag
	// Etag *string `json:"etag,omitempty"`

	// When suspended, no new releases or rollouts can be created,
	//  but in-progress ones will complete.
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipeline.suspended
	Suspended *bool `json:"suspended,omitempty"`
}

// DeliveryPipelineStatus defines the config connector machine state of DeployDeliveryPipeline
type DeliveryPipelineStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1beta1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DeployDeliveryPipeline resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DeliveryPipelineObservedState `json:"observedState,omitempty"`
}

// DeliveryPipelineObservedState is the state of the DeployDeliveryPipeline resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.deploy.v1.DeliveryPipeline
type DeliveryPipelineObservedState struct {
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
// +kubebuilder:resource:categories=gcp,shortName=gcpdeploydeliverypipeline;gcpdeploydeliverypipelines
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="internal.cloud.google.com/additional-versions=v1alpha1"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"
// +kubebuilder:storageversion

// CloudDeployDeliveryPipeline is the Schema for the CloudDeployDeliveryPipeline API
// +k8s:openapi-gen=true
type CloudDeployDeliveryPipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DeliveryPipelineSpec   `json:"spec,omitempty"`
	Status DeliveryPipelineStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudDeployDeliveryPipelineList contains a list of DeployDeliveryPipeline
type CloudDeployDeliveryPipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudDeployDeliveryPipeline `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudDeployDeliveryPipeline{}, &CloudDeployDeliveryPipelineList{})
}

// +kcc:proto=google.cloud.deploy.v1.KubernetesConfig.GatewayServiceMesh.RouteDestinations
type KubernetesConfig_GatewayServiceMesh_RouteDestinations struct {
	// Required. The clusters where the Gateway API HTTPRoute resource will be
	//  deployed to. Valid entries include the associated entities IDs
	//  configured in the Target resource and "@self" to include the Target
	//  cluster.
	// +kcc:proto:field=google.cloud.deploy.v1.KubernetesConfig.GatewayServiceMesh.RouteDestinations.destination_ids
	DestinationIDs []string `json:"destinationIDs,omitempty"`

	// Optional. Whether to propagate the Kubernetes Service to the route
	//  destination clusters. The Service will always be deployed to the Target
	//  cluster even if the HTTPRoute is not. This option may be used to
	//  facilitate successful DNS lookup in the route destination clusters. Can
	//  only be set to true if destinations are specified.
	// +kcc:proto:field=google.cloud.deploy.v1.KubernetesConfig.GatewayServiceMesh.RouteDestinations.propagate_service
	PropagateService *bool `json:"propagateService,omitempty"`
}
