// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	pubsubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudSupportSupportEventSubscriptionGVK = GroupVersion.WithKind("CloudSupportSupportEventSubscription")

// CloudSupportSupportEventSubscriptionSpec defines the desired state of CloudSupportSupportEventSubscription
// +kcc:spec:proto=google.cloud.support.v2.SupportEventSubscription
type CloudSupportSupportEventSubscriptionSpec struct {
	// The organization that this resource belongs to.
	// +kubebuilder:validation:Required
	OrganizationRef *refsv1beta1.OrganizationRef `json:"organizationRef"`

	// The Pub/Sub topic to publish notifications to.
	// +kubebuilder:validation:Required
	PubSubTopicRef *pubsubv1beta1.PubSubTopicRef `json:"pubSubTopicRef"`

	// The CloudSupportSupportEventSubscription name. If not given, the metadata.name will be used.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// CloudSupportSupportEventSubscriptionStatus defines the config connector machine state of CloudSupportSupportEventSubscription
type CloudSupportSupportEventSubscriptionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudSupportSupportEventSubscription resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudSupportSupportEventSubscriptionObservedState `json:"observedState,omitempty"`
}

// CloudSupportSupportEventSubscriptionObservedState is the state of the CloudSupportSupportEventSubscription resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.support.v2.SupportEventSubscription
type CloudSupportSupportEventSubscriptionObservedState struct {
	// Output only. The state of the subscription.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty"`

	// Output only. Reason why subscription is failing. State of subscription
	//  must be FAILING in order for this to have a value.
	// +kubebuilder:validation:Optional
	FailureReason *string `json:"failureReason,omitempty"`

	// Output only. The time at which the subscription was created.
	// +kubebuilder:validation:Optional
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the subscription was last updated.
	// +kubebuilder:validation:Optional
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The time at which the subscription was deleted.
	// +kubebuilder:validation:Optional
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. The time at which the subscription will be purged.
	// +kubebuilder:validation:Optional
	PurgeTime *string `json:"purgeTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudsupportsupporteventsubscription;gcpcloudsupportsupporteventsubscriptions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudSupportSupportEventSubscription is the Schema for the CloudSupportSupportEventSubscription API
// +k8s:openapi-gen=true
type CloudSupportSupportEventSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudSupportSupportEventSubscriptionSpec   `json:"spec,omitempty"`
	Status CloudSupportSupportEventSubscriptionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudSupportSupportEventSubscriptionList contains a list of CloudSupportSupportEventSubscription
type CloudSupportSupportEventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudSupportSupportEventSubscription `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudSupportSupportEventSubscription{}, &CloudSupportSupportEventSubscriptionList{})
}
