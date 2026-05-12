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
	pubsubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SecretManagerRegionalSecretGVK = GroupVersion.WithKind("SecretManagerRegionalSecret")

// SecretManagerRegionalSecretSpec defines the desired state of SecretManagerRegionalSecret
// +kcc:spec:proto=google.cloud.secretmanager.v1.Secret
type SecretManagerRegionalSecretSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The SecretManagerRegionalSecret name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. A list of up to 10 Pub/Sub topics to which messages are published
	//  when control plane operations are called on the secret or its versions.
	TopicRefs []*TopicRef `json:"topics,omitempty"`

	// Optional. Timestamp in UTC when the
	//  [Secret][google.cloud.secretmanager.v1.Secret] is scheduled to expire.
	//  This is always provided on output, regardless of what was sent on input.
	ExpireTime *string `json:"expireTime,omitempty"`

	// Input only. A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	TTL *string `json:"ttl,omitempty"`

	// Optional. Rotation policy attached to the
	//  [Secret][google.cloud.secretmanager.v1.Secret]. May be excluded if there is
	//  no rotation policy.
	Rotation *Rotation `json:"rotation,omitempty"`

	// Optional. Mapping from version alias to version name.
	//
	//  A version alias is a string with a maximum length of 63 characters and can
	//  contain uppercase and lowercase letters, numerals, and the hyphen (`-`)
	//  and underscore ('_') characters. An alias string must start with a
	//  letter and cannot be the string 'latest' or 'NEW'.
	//  No more than 50 aliases can be assigned to a given secret.
	//
	//  Version-Alias pairs will be viewable via GetSecret, but not via ListSecrets.
	//
	//  Note: The underlying GCP API restricts the aliases that can be updated during a Create or Update Secret request.
	//  Config Connector removes the aliases that are not updatable during these requests.
	//  See https://cloud.google.com/secret-manager/docs/reference/rest/v1/projects.secrets#Secret for more details.
	VersionAliases map[string]string `json:"versionAliases,omitempty"`

	// Optional. Custom metadata about the secret.
	//
	//  Annotations are dictionary keys with values both of which are strings.
	//
	//  Note: The underlying GCP API restricts the annotations.
	//  Config Connector removes the annotations that are not valid.
	//  See https://cloud.google.com/secret-manager/docs/reference/rest/v1/projects.secrets#Secret for more details.
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. The customer-managed encryption configuration of the regional secret.
	CustomerManagedEncryption *CustomerManagedEncryption `json:"customerManagedEncryption,omitempty"`
}

// TopicRef is a reference to a Pub/Sub topic
type TopicRef struct {
	// A reference to a Pub/Sub topic
	PubSubTopicRef *pubsubv1beta1.PubSubTopicRef `json:"pubsubTopicRef,omitempty"`
}

// +kcc:proto=google.cloud.secretmanager.v1.CustomerManagedEncryption
type CustomerManagedEncryption struct {
	// +required
	// Required. The resource name of the Cloud KMS CryptoKey used to encrypt
	//  secret payloads.
	KmsKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// SecretManagerRegionalSecretStatus defines the config connector machine state of SecretManagerRegionalSecret
type SecretManagerRegionalSecretStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SecretManagerRegionalSecret resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// The name of the SecretManagerRegionalSecret resource.
	Name string `json:"name,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SecretManagerRegionalSecretObservedState `json:"observedState,omitempty"`
}

// SecretManagerRegionalSecretObservedState is the state of the SecretManagerRegionalSecret resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.secretmanager.v1.Secret
type SecretManagerRegionalSecretObservedState struct {
	// Output only. The time at which the
	//  [Secret][google.cloud.secretmanager.v1.Secret] was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Mapping from version alias to version name.
	VersionAliases map[string]string `json:"versionAliases,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsecretmanagerregionalsecret;gcpsecretmanagerregionalsecrets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SecretManagerRegionalSecret is the Schema for the SecretManagerRegionalSecret API
// +k8s:openapi-gen=true
type SecretManagerRegionalSecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SecretManagerRegionalSecretSpec   `json:"spec,omitempty"`
	Status SecretManagerRegionalSecretStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SecretManagerRegionalSecretList contains a list of SecretManagerRegionalSecret
type SecretManagerRegionalSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretManagerRegionalSecret `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecretManagerRegionalSecret{}, &SecretManagerRegionalSecretList{})
}
