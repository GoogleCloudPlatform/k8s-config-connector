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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var KMSCryptoKeyGVK = GroupVersion.WithKind("KMSCryptoKey")

// KMSCryptoKeySpec defines the desired state of KMSCryptoKey
// +kcc:proto=google.cloud.kms.v1.CryptoKey
type KMSCryptoKeySpec struct {
	// /* Immutable. The Project that this resource belongs to. */
	// ProjectRef refs.ProjectRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. Whether this key may contain imported versions only.
	ImportOnly *bool `json:"importOnly,omitempty"`

	// Immutable. The period of time that versions of this key spend in the
	//  [DESTROY_SCHEDULED][google.cloud.kms.v1.CryptoKeyVersion.CryptoKeyVersionState.DESTROY_SCHEDULED]
	//  state before transitioning to
	//  [DESTROYED][google.cloud.kms.v1.CryptoKeyVersion.CryptoKeyVersionState.DESTROYED].
	//  If not specified at creation time, the default duration is 24 hours.
	DestroyScheduledDuration *string `json:"destroyScheduledDuration,omitempty"`

	// Immutable. The immutable purpose of this
	//  [CryptoKey][google.cloud.kms.v1.CryptoKey].
	Purpose *string `json:"purpose,omitempty"`

	// [next_rotation_time][google.cloud.kms.v1.CryptoKey.next_rotation_time]
	//  will be advanced by this period when the service automatically rotates a
	//  key. Must be at least 24 hours and at most 876,000 hours.
	//
	//  If [rotation_period][google.cloud.kms.v1.CryptoKey.rotation_period] is
	//  set,
	//  [next_rotation_time][google.cloud.kms.v1.CryptoKey.next_rotation_time]
	//  must also be set.
	//
	//  Keys with [purpose][google.cloud.kms.v1.CryptoKey.purpose]
	//  [ENCRYPT_DECRYPT][google.cloud.kms.v1.CryptoKey.CryptoKeyPurpose.ENCRYPT_DECRYPT]
	//  support automatic rotation. For other keys, this field must be omitted.
	RotationPeriod *string `json:"rotationPeriod,omitempty"`

	// Immutable. If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
	SkipInitialVersionCreation *bool `json:"skipInitialVersionCreation,omitempty"`

	// A template describing settings for new
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion] instances. The
	//  properties of new [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion]
	//  instances created by either
	//  [CreateCryptoKeyVersion][google.cloud.kms.v1.KeyManagementService.CreateCryptoKeyVersion]
	//  or auto-rotation are controlled by this template.
	VersionTemplate *CryptoKeyVersionTemplate `json:"versionTemplate,omitempty"`

	// The KMSKeyRing that this key belongs to.
	// +required
	KeyRingRef *refs.KMSKeyRingRef `json:"keyRingRef,omitempty"`
}

// KMSCryptoKeyStatus defines the config connector machine state of KMSCryptoKey
// +kcc:proto=google.cloud.kms.v1.CryptoKey
type KMSCryptoKeyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* NOTYET
	// A unique specifier for the KMSCryptoKey resource in GCP.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`
	*/

	/* NOTYET
	// ObservedState is the state of the resource as most recently observed in GCP.
	// +optional
	ObservedState *KMSCryptoKeyObservedState `json:"observedState,omitempty"`
	*/

	// The self link of the created key in the format projects/{project}/locations/{location}/keyRings/{keyRingName}/cryptoKeys/{name}.
	SelfLink *string `json:"selfLink,omitempty"`

	// // Output only. The resource name for this
	// //  [CryptoKey][google.cloud.kms.v1.CryptoKey] in the format
	// //  `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
	// Name *string `json:"name,omitempty"`

	// // Output only. A copy of the "primary"
	// //  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion] that will be used
	// //  by [Encrypt][google.cloud.kms.v1.KeyManagementService.Encrypt] when this
	// //  [CryptoKey][google.cloud.kms.v1.CryptoKey] is given in
	// //  [EncryptRequest.name][google.cloud.kms.v1.EncryptRequest.name].
	// //
	// //  The [CryptoKey][google.cloud.kms.v1.CryptoKey]'s primary version can be
	// //  updated via
	// //  [UpdateCryptoKeyPrimaryVersion][google.cloud.kms.v1.KeyManagementService.UpdateCryptoKeyPrimaryVersion].
	// //
	// //  Keys with [purpose][google.cloud.kms.v1.CryptoKey.purpose]
	// //  [ENCRYPT_DECRYPT][google.cloud.kms.v1.CryptoKey.CryptoKeyPurpose.ENCRYPT_DECRYPT]
	// //  may have a primary. For other keys, this field will be omitted.
	// Primary *CryptoKeyVersion `json:"primary,omitempty"`

	// // Output only. The time at which this
	// //  [CryptoKey][google.cloud.kms.v1.CryptoKey] was created.
	// CreateTime *string `json:"createTime,omitempty"`

	// // At [next_rotation_time][google.cloud.kms.v1.CryptoKey.next_rotation_time],
	// //  the Key Management Service will automatically:
	// //
	// //  1. Create a new version of this [CryptoKey][google.cloud.kms.v1.CryptoKey].
	// //  2. Mark the new version as primary.
	// //
	// //  Key rotations performed manually via
	// //  [CreateCryptoKeyVersion][google.cloud.kms.v1.KeyManagementService.CreateCryptoKeyVersion]
	// //  and
	// //  [UpdateCryptoKeyPrimaryVersion][google.cloud.kms.v1.KeyManagementService.UpdateCryptoKeyPrimaryVersion]
	// //  do not affect
	// //  [next_rotation_time][google.cloud.kms.v1.CryptoKey.next_rotation_time].
	// //
	// //  Keys with [purpose][google.cloud.kms.v1.CryptoKey.purpose]
	// //  [ENCRYPT_DECRYPT][google.cloud.kms.v1.CryptoKey.CryptoKeyPurpose.ENCRYPT_DECRYPT]
	// //  support automatic rotation. For other keys, this field must be omitted.
	// NextRotationTime *string `json:"nextRotationTime,omitempty"`

	// // Labels with user-defined metadata. For more information, see
	// //  [Labeling Keys](https://cloud.google.com/kms/docs/labeling-keys).
	// Labels map[string]string `json:"labels,omitempty"`

	// // Immutable. The resource name of the backend environment where the key
	// //  material for all [CryptoKeyVersions][google.cloud.kms.v1.CryptoKeyVersion]
	// //  associated with this [CryptoKey][google.cloud.kms.v1.CryptoKey] reside and
	// //  where all related cryptographic operations are performed. Only applicable
	// //  if [CryptoKeyVersions][google.cloud.kms.v1.CryptoKeyVersion] have a
	// //  [ProtectionLevel][google.cloud.kms.v1.ProtectionLevel] of
	// //  [EXTERNAL_VPC][CryptoKeyVersion.ProtectionLevel.EXTERNAL_VPC], with the
	// //  resource name in the format `projects/*/locations/*/ekmConnections/*`.
	// //  Note, this list is non-exhaustive and may apply to additional
	// //  [ProtectionLevels][google.cloud.kms.v1.ProtectionLevel] in the future.
	// CryptoKeyBackend *string `json:"cryptoKeyBackend,omitempty"`

	// // Optional. The policy used for Key Access Justifications Policy Enforcement.
	// //  If this field is present and this key is enrolled in Key Access
	// //  Justifications Policy Enforcement, the policy will be evaluated in encrypt,
	// //  decrypt, and sign operations, and the operation will fail if rejected by
	// //  the policy. The policy is defined by specifying zero or more allowed
	// //  justification codes.
	// //  https://cloud.google.com/assured-workloads/key-access-justifications/docs/justification-codes
	// //  By default, this field is absent, and all justification codes are allowed.
	// KeyAccessJustificationsPolicy *KeyAccessJustificationsPolicy `json:"keyAccessJustificationsPolicy,omitempty"`

}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpkmscryptokey;gcpkmscryptokeys
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true";"cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// KMSCryptoKey is the Schema for the KMSCryptoKey API
// +k8s:openapi-gen=true
type KMSCryptoKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
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
