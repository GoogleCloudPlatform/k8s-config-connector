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

// +tool:krm-type-terraform
// proto.message: google.cloud.pubsublite.v1.Subscription
// crd.kind: PubSubLiteSubscription
// crd.version: v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var PubSubLiteSubscriptionGVK = GroupVersion.WithKind("PubSubLiteSubscription")

// PubSubLiteSubscriptionSpec defines the desired state of PubSubLiteSubscription
// +kcc:spec:proto=google.cloud.pubsublite.v1.Subscription
type PubSubLiteSubscriptionSpec struct {
	// Required. Defines the parent path of the resource.
	// +required
	ProjectRef *parent.ProjectRef `json:"projectRef,omitempty"`

	// The region of the pubsub lite topic.
	Region *string `json:"region,omitempty"`

	// Immutable. A reference to a Topic resource.
	// +required
	Topic *string `json:"topic,omitempty"`

	// The zone of the pubsub lite topic.
	// +required
	Zone *string `json:"zone,omitempty"`

	// The PubSubLiteSubscription name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// // The name of the subscription.
	// //  Structured like:
	// //  projects/{project_number}/locations/{location}/subscriptions/{subscription_id}
	// // +kcc:proto:field=google.cloud.pubsublite.v1.Subscription.name
	// Name *string `json:"name,omitempty"`

	// // The name of the topic this subscription is attached to.
	// //  Structured like:
	// //  projects/{project_number}/locations/{location}/topics/{topic_id}
	// // +kcc:proto:field=google.cloud.pubsublite.v1.Subscription.topic
	// Topic *string `json:"topic,omitempty"`

	// The settings for this subscription's message delivery.
	// +kcc:proto:field=google.cloud.pubsublite.v1.Subscription.delivery_config
	DeliveryConfig *Subscription_DeliveryConfig `json:"deliveryConfig,omitempty"`

	// // If present, messages are automatically written from the Pub/Sub Lite topic
	// //  associated with this subscription to a destination.
	// // +kcc:proto:field=google.cloud.pubsublite.v1.Subscription.export_config
	// ExportConfig *ExportConfig `json:"exportConfig,omitempty"`
}

// PubSubLiteSubscriptionStatus defines the config connector machine state of PubSubLiteSubscription
type PubSubLiteSubscriptionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the PubSubLiteSubscription resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// NOTYET(terraform)
	// // ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *PubSubLiteSubscriptionObservedState `json:"observedState,omitempty"`
}

// PubSubLiteSubscriptionObservedState is the state of the PubSubLiteSubscription resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.pubsublite.v1.Subscription
type PubSubLiteSubscriptionObservedState struct {
	// // If present, messages are automatically written from the Pub/Sub Lite topic
	//    //  associated with this subscription to a destination.
	//    // +kcc:proto:field=google.cloud.pubsublite.v1.Subscription.export_config
	//    ExportConfig *ExportConfigObservedState `json:"exportConfig,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcppubsublitesubscription;gcppubsublitesubscriptions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// PubSubLiteSubscription is the Schema for the PubSubLiteSubscription API
// +k8s:openapi-gen=true
type PubSubLiteSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   PubSubLiteSubscriptionSpec   `json:"spec,omitempty"`
	Status PubSubLiteSubscriptionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// PubSubLiteSubscriptionList contains a list of PubSubLiteSubscription
type PubSubLiteSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PubSubLiteSubscription `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PubSubLiteSubscription{}, &PubSubLiteSubscriptionList{})
}
