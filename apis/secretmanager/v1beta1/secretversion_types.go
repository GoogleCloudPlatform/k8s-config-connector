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
	refsv1beta1secret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SecretManagerSecretVersionGVK = GroupVersion.WithKind("SecretManagerSecretVersion")

// SecretManagerSecretVersionSpec defines the desired state of SecretManagerSecretVersion
// +kcc:proto=google.cloud.secretmanager.v1.SecretVersion
type SecretManagerSecretVersionSpec struct {
	// The resource name of the [Secret][google.cloud.secretmanager.v1.Secret] to create a [SecretVersion][google.cloud.secretmanager.v1.SecretVersion] for.
	SecretRef *SecretRef `json:"secretRef,omitempty"`

	// The SecretVersion number. If given, Config Connector acquires the resource from the Secret Manager service.
	// If not given, Config Connector adds a new secret version to the GCP service, and you can find out the version number
	// from `status.observedState.version`
	ResourceID *string `json:"resourceID,omitempty"`

	// Should enable or disable the current SecretVersion.
	// - Enabled version can be accessed and described.
	// - Disabled version cannot be accessed, but the secret's contents still exist
	Enabled *bool `json:"enabled,omitempty"`

	// The actual secret data. Config Connector supports secret data stored in Kubernetes secret or plain data (base64)
	SecretData *SecretData_OneOf `json:"secretData,omitempty"`

	// DEPRECATED. You do not need to set this field in direct reconciler mode. Use delete-policy annotation instead. https://cloud.google.com/config-connector/docs/how-to/managing-deleting-resources#keeping_resources_after_deletion
	// The deletion policy for the secret version. Setting 'ABANDON' allows the resource
	// to be abandoned rather than deleted. Setting 'DISABLE' allows the resource to be
	// disabled rather than deleted. Default is 'DELETE'. Possible values are:
	// * DELETE
	// * DISABLE
	// * ABANDON.
	DeletionPolicy *string `json:"deletionPolicy,omitempty"`

	// DEPRECATED. You do not need to set this field in direct reconciler mode.
	IsSecretDataBase64 *bool `json:"isSecretDataBase64,omitempty"`
}

type SecretData_OneOf struct {
	*refsv1beta1secret.Legacy `json:",inline"`
	// TODO: support getting secret data from other places.
}

// SecretManagerSecretVersionStatus defines the config connector machine state of SecretManagerSecretVersion
type SecretManagerSecretVersionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SecretManagerSecretVersion resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SecretManagerSecretVersionObservedState `json:"observedState,omitempty"`

	// DEPRECATING NOTE: Please use status.observedState.createTime instead.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// DEPRECATING NOTE: Please use status.observedState.destroyTime instead.
	DestroyTime *string `json:"destroyTime,omitempty" tf:"destroy_time,omitempty"`

	// DEPRECATING NOTE: Please use status.observedState.name instead.
	Name *string `json:"name,omitempty"`

	// DEPRECATED.
	Version *string `json:"version,omitempty"`
}

// SecretManagerSecretVersionObserved is the state of the SecretManagerSecretVersion resource as most recently observed in GCP.
// +kcc:proto=google.cloud.secretmanager.v1.SecretVersion
type SecretManagerSecretVersionObservedState struct {
	// Output only. The time at which the
	// [SecretVersion][google.cloud.secretmanager.v1.SecretVersion] was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Output only. The time this
	// [SecretVersion][google.cloud.secretmanager.v1.SecretVersion] was destroyed.
	// Only present if
	// [state][google.cloud.secretmanager.v1.SecretVersion.state] is
	// [DESTROYED][google.cloud.secretmanager.v1.SecretVersion.State.DESTROYED].
	DestroyTime *string `json:"destroyTime,omitempty" tf:"destroy_time,omitempty"`

	// Output only. The resource name of the
	// [SecretVersion][google.cloud.secretmanager.v1.SecretVersion] in the
	// format `projects/*/secrets/*/versions/*`.
	//
	// [SecretVersion][google.cloud.secretmanager.v1.SecretVersion] IDs in a
	// [Secret][google.cloud.secretmanager.v1.Secret] start at 1 and are
	// incremented for each subsequent version of the secret.
	Name *string `json:"name,omitempty"`

	// The replication status of the
	// [SecretVersion][google.cloud.secretmanager.v1.SecretVersion].
	ReplicationStatus *ReplicationStatus `json:"replicationStatus,omitempty"`

	// Optional. Output only. Scheduled destroy time for secret version.
	// This is a part of the Delayed secret version destroy feature. For a
	// Secret with a valid version destroy TTL, when a secert version is
	// destroyed, the version is moved to disabled state and it is scheduled for
	// destruction. The version is destroyed only after the
	// `scheduled_destroy_time`.
	ScheduledDestroyTime *string `json:"scheduledDestroyTime,omitempty"`

	// Output only. True if payload checksum specified in
	//  [SecretPayload][google.cloud.secretmanager.v1.SecretPayload] object has
	//  been received by
	//  [SecretManagerService][google.cloud.secretmanager.v1.SecretManagerService]
	//  on
	//  [SecretManagerService.AddSecretVersion][google.cloud.secretmanager.v1.SecretManagerService.AddSecretVersion].
	ClientSpecifiedPayloadChecksum *bool `json:"clientSpecifiedPayloadChecksum,omitempty"`

	// Output only. The customer-managed encryption status of the
	//  [SecretVersion][google.cloud.secretmanager.v1.SecretVersion]. Only
	//  populated if customer-managed encryption is used and
	//  [Secret][google.cloud.secretmanager.v1.Secret] is a Regionalised Secret.
	CustomerManagedEncryption *CustomerManagedEncryptionStatus `json:"customerManagedEncryption,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsecretmanagersecretversion;gcpsecretmanagersecretversions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SecretManagerSecretVersion is the Schema for the SecretManagerSecretVersion API
// +k8s:openapi-gen=true
type SecretManagerSecretVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SecretManagerSecretVersionSpec   `json:"spec,omitempty"`
	Status SecretManagerSecretVersionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SecretManagerSecretVersionList contains a list of SecretManagerSecretVersion
type SecretManagerSecretVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretManagerSecretVersion `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecretManagerSecretVersion{}, &SecretManagerSecretVersionList{})
}
