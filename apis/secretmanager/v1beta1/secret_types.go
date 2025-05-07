// Copyright 2024 Google LLC
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
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SecretManagerSecretGVK = GroupVersion.WithKind("SecretManagerSecret")

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SecretManagerSecretSpec defines the desired state of SecretManagerSecret
// +kcc:proto=google.cloud.secretmanager.v1.Secret
type SecretManagerSecretSpec struct {
	// The SecretManagerSecret name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Immutable. The replication policy of the secret data attached to
	//  the [Secret][google.cloud.secretmanager.v1.Secret].
	//
	//  The replication policy cannot be changed after the Secret has been created.
	Replication *Replication `json:"replication,omitempty"`

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
	//  Version-Alias pairs will be viewable via GetSecret and modifiable via
	//  UpdateSecret. Access by alias is only be supported on
	//  GetSecretVersion and AccessSecretVersion.
	VersionAliases map[string]string `json:"versionAliases,omitempty"`

	// Optional. Custom metadata about the secret.
	//
	//  Annotations are distinct from various forms of labels.
	//  Annotations exist to allow client tools to store their own state
	//  information without requiring a database.
	//
	//  Annotation keys must be between 1 and 63 characters long, have a UTF-8
	//  encoding of maximum 128 bytes, begin and end with an alphanumeric character
	//  ([a-z0-9A-Z]), and may have dashes (-), underscores (_), dots (.), and
	//  alphanumerics in between these symbols.
	//
	//  The total size of annotation keys and values must be less than 16KiB.
	Annotations map[string]string `json:"annotations,omitempty"`

	// The labels assigned to this Secret.
	//
	// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding
	// of maximum 128 bytes, and must conform to the following PCRE regular
	// expression: `[\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}`
	//
	// Label values must be between 0 and 63 characters long, have a UTF-8
	// encoding of maximum 128 bytes, and must conform to the following PCRE
	// regular expression: `[\p{Ll}\p{Lo}\p{N}_-]{0,63}`
	//
	// No more than 64 labels can be assigned to a given resource.
	Labels map[string]string `json:"labels,omitempty"`

	/*NOTYET
	// Optional. Secret Version TTL after destruction request
	//
	//  This is a part of the Delayed secret version destroy feature.
	//  For secret with TTL>0, version destruction doesn't happen immediately
	//  on calling destroy instead the version goes to a disabled state and
	//  destruction happens after the TTL expires.
	VersionDestroyTtl *string `json:"versionDestroyTtl,omitempty"`
	*/

	/*NOTYET
	// Optional. The customer-managed encryption configuration of the Regionalised
	//  Secrets. If no configuration is provided, Google-managed default encryption
	//  is used.
	//
	//  Updates to the [Secret][google.cloud.secretmanager.v1.Secret] encryption
	//  configuration only apply to
	//  [SecretVersions][google.cloud.secretmanager.v1.SecretVersion] added
	//  afterwards. They do not apply retroactively to existing
	//  [SecretVersions][google.cloud.secretmanager.v1.SecretVersion].
	CustomerManagedEncryption *CustomerManagedEncryption `json:"customerManagedEncryption,omitempty"`
	*/
}

// +kcc:proto=google.cloud.secretmanager.v1.Replication
type Replication struct {
	// The Secret will automatically be replicated without any restrictions.
	LegacyAuto *bool `json:"automatic,omitempty"`

	// The [Secret][google.cloud.secretmanager.v1.Secret] will automatically be
	//  replicated without any restrictions.
	LegacyAutomatic *Replication_Automatic `json:"auto,omitempty"`

	// The [Secret][google.cloud.secretmanager.v1.Secret] will only be
	//  replicated into the locations specified.
	UserManaged *Replication_UserManaged `json:"userManaged,omitempty"`
}

// +kcc:proto=google.cloud.secretmanager.v1.Replication.UserManaged
type Replication_UserManaged struct {
	// +required
	// Required. The list of Replicas for this
	//  [Secret][google.cloud.secretmanager.v1.Secret].
	//
	//  Cannot be empty.
	Replicas []Replication_UserManaged_Replica `json:"replicas,omitempty"`
}

// +kcc:proto=google.cloud.secretmanager.v1.Replication.UserManaged.Replica
type Replication_UserManaged_Replica struct {
	// +required
	// The canonical IDs of the location to replicate data.
	//  For example: `"us-east1"`.
	Location *string `json:"location,omitempty"`

	// Optional. The customer-managed encryption configuration of the
	//  [User-Managed Replica][Replication.UserManaged.Replica]. If no
	//  configuration is provided, Google-managed default encryption is used.
	//
	//  Updates to the [Secret][google.cloud.secretmanager.v1.Secret]
	//  encryption configuration only apply to
	//  [SecretVersions][google.cloud.secretmanager.v1.SecretVersion] added
	//  afterwards. They do not apply retroactively to existing
	//  [SecretVersions][google.cloud.secretmanager.v1.SecretVersion].
	CustomerManagedEncryption *CustomerManagedEncryption `json:"customerManagedEncryption,omitempty"`
}

type TopicRef struct {
	// +required
	PubSubTopicRef *refv1beta1.PubSubTopicRef `json:"topicRef,omitempty"`
}

// +kcc:proto=google.cloud.secretmanager.v1.CustomerManagedEncryption
type CustomerManagedEncryption struct {
	// +required
	// Required. The resource name of the Cloud KMS CryptoKey used to encrypt
	//  secret payloads.
	//
	//  For secrets using the
	//  [UserManaged][google.cloud.secretmanager.v1.Replication.UserManaged]
	//  replication policy type, Cloud KMS CryptoKeys must reside in the same
	//  location as the [replica location][Secret.UserManaged.Replica.location].
	//
	//  For secrets using the
	//  [Automatic][google.cloud.secretmanager.v1.Replication.Automatic]
	//  replication policy type, Cloud KMS CryptoKeys must reside in `global`.
	//
	//  The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
	KmsKeyRef *refv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// SecretManagerSecretStatus defines the config connector machine state of SecretManagerSecret
type SecretManagerSecretStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SecretManagerSecret resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SecretManagerSecretObservedState `json:"observedState,omitempty"`

	// [DEPRECATED] Please read from `.status.externalRef` instead. Config Connector will remove the `.status.name` in v1 Version.
	Name string `json:"name,omitempty"`
}

// SecretManagerSecretSpec defines the desired state of SecretManagerSecret
// +kcc:proto=google.cloud.secretmanager.v1.Secret
type SecretManagerSecretObservedState struct {
	VersionAliases map[string]string `json:"versionAliases,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsecretmanagersecret;gcpsecretmanagersecrets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SecretManagerSecret is the Schema for the SecretManagerSecret API
// +k8s:openapi-gen=true
type SecretManagerSecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecretManagerSecretSpec   `json:"spec,omitempty"`
	Status SecretManagerSecretStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SecretManagerSecretList contains a list of SecretManagerSecret
type SecretManagerSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretManagerSecret `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecretManagerSecret{}, &SecretManagerSecretList{})
}
