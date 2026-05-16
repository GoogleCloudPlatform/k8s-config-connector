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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KMSCryptoKeySpec defines the desired state of KMSCryptoKey
// +kcc:spec:proto=google.cloud.kms.v1.CryptoKey
type KMSCryptoKeySpec struct {
	/* The KMSKeyRing that this key belongs to. */
	KeyRingRef *KMSKeyRingRef `json:"keyRingRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. The immutable purpose of this CryptoKey. See the
	[purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
	for possible inputs.
	Default value is "ENCRYPT_DECRYPT". */
	// +optional
	// +kubebuilder:default="ENCRYPT_DECRYPT"
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKey.purpose
	Purpose *string `json:"purpose,omitempty"`

	/* A template describing settings for new crypto key versions. */
	// +optional
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKey.version_template
	VersionTemplate *CryptoKeyVersionTemplate `json:"versionTemplate,omitempty"`

	/* Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
	The first rotation will take place after the specified period. The rotation period has
	the format of a decimal number with up to 9 fractional digits, followed by the
	letter 's' (seconds). It must be greater than a day (ie, 86400). */
	// +optional
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKey.rotation_period
	RotationPeriod *string `json:"rotationPeriod,omitempty"`

	/* Immutable. Whether this key may contain imported versions only. */
	// +optional
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKey.import_only
	ImportOnly *bool `json:"importOnly,omitempty"`

	/* Immutable. The period of time that versions of this key spend in the DESTROY_SCHEDULED state before transitioning to DESTROYED.
	If not specified at creation time, the default duration is 24 hours. */
	// +optional
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKey.destroy_scheduled_duration
	DestroyScheduledDuration *string `json:"destroyScheduledDuration,omitempty"`

	/* Immutable. If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
	You must use the 'google_kms_key_ring_import_job' resource to import the CryptoKeyVersion. */
	// +optional
	// +kcc:proto:field=google.cloud.kms.v1.CreateCryptoKeyRequest.skip_initial_version_creation
	SkipInitialVersionCreation *bool `json:"skipInitialVersionCreation,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.CryptoKeyVersionTemplate
type CryptoKeyVersionTemplate struct {
	/* The algorithm to use when creating a version based on this template.
	See the [algorithm reference](https://cloud.google.com/kms/docs/reference/rest/v1/CryptoKeyVersionAlgorithm) for possible inputs. */
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKeyVersionTemplate.algorithm
	Algorithm *string `json:"algorithm,omitempty"`

	/* Immutable. The protection level to use when creating a version based on this template. Possible values include "SOFTWARE", "HSM", "EXTERNAL", "EXTERNAL_VPC". Defaults to "SOFTWARE". */
	// +optional
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKeyVersionTemplate.protection_level
	ProtectionLevel *string `json:"protectionLevel,omitempty"`
}

// KMSCryptoKeyStatus defines the config connector machine state of KMSCryptoKey
type KMSCryptoKeyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1beta1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the KMSCryptoKey resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// The self link of the created CryptoKey in the format projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{name}.
	SelfLink *string `json:"selfLink,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *KMSCryptoKeyObservedState `json:"observedState,omitempty"`
}

// KMSCryptoKeyObservedState is the state of the KMSCryptoKey resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.kms.v1.CryptoKey
type KMSCryptoKeyObservedState struct {
	// Output only. The time at which this CryptoKey was created.
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKey.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. A copy of the "primary" CryptoKeyVersion that will be used by Encrypt when this CryptoKey is given in EncryptRequest.name without a version.
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKey.primary
	Primary *CryptoKeyVersionObservedState `json:"primary,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.kms.v1.CryptoKeyVersion
type CryptoKeyVersionObservedState struct {
	// The resource name for this CryptoKeyVersion.
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKeyVersion.name
	Name *string `json:"name,omitempty"`

	// The current state of the CryptoKeyVersion.
	// +kcc:proto:field=google.cloud.kms.v1.CryptoKeyVersion.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpkmscryptokey;gcpkmscryptokeys
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/default-controller=direct"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// KMSCryptoKey is the Schema for the KMSCryptoKey API
// +kubebuilder:storageversion
// +k8s:openapi-gen=true
type KMSCryptoKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KMSCryptoKeySpec   `json:"spec,omitempty"`
	Status KMSCryptoKeyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// KMSCryptoKeyList contains a list of KMSCryptoKey
type KMSCryptoKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KMSCryptoKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KMSCryptoKey{}, &KMSCryptoKeyList{})
}
