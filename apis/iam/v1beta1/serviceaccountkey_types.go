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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var IAMServiceAccountKeyGVK = GroupVersion.WithKind("IAMServiceAccountKey")

// IAMServiceAccountKeySpec defines the desired state of IAMServiceAccountKey
// +kcc:spec:proto=google.iam.admin.v1.ServiceAccountKey
type IAMServiceAccountKeySpec struct {
	// Immutable. The Service Account to create a key for.
	// +required
	ServiceAccountRef *refs.IAMServiceAccountRef `json:"serviceAccountRef"`

	// Immutable. The algorithm used to generate the key, used only on create. KEY_ALG_RSA_2048 is the default algorithm. Valid values are: "KEY_ALG_RSA_1024", "KEY_ALG_RSA_2048".
	// +optional
	KeyAlgorithm *string `json:"keyAlgorithm,omitempty"`

	// Immutable. The output format of the private key. TYPE_GOOGLE_CREDENTIALS_FILE is the default output format.
	// +optional
	PrivateKeyType *string `json:"privateKeyType,omitempty"`

	// Immutable. A field that allows clients to upload their own public key. If set, use this public key data to create a service account key for given service account. Please note, the expected format for this field is a base64 encoded X509_PEM.
	// +optional
	PublicKeyData *string `json:"publicKeyData,omitempty"`

	// Immutable. The output format for the public key. TYPE_NONE is the default for public key output.
	// +optional
	PublicKeyType *string `json:"publicKeyType,omitempty"`
}

// IAMServiceAccountKeyStatus defines the config connector machine state of IAMServiceAccountKey
type IAMServiceAccountKeyStatus struct {
	// Conditions represent the latest available observations of the object's current state.
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// Immutable. The name used for this key pair.
	// +optional
	Name *string `json:"name,omitempty"`

	// The private key in JSON format, base64 encoded. This is what you normally get as a file when creating service account keys through the CLI or web console. This is only populated when creating a new key.
	// +optional
	PrivateKey *string `json:"privateKey,omitempty"`

	// Immutable. The public key, base64 encoded.
	// +optional
	PublicKey *string `json:"publicKey,omitempty"`

	// The key can be used after this timestamp. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	// +optional
	ValidAfter *string `json:"validAfter,omitempty"`

	// The key can be used before this timestamp. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	// +optional
	ValidBefore *string `json:"validBefore,omitempty"`
}

// IAMServiceAccountKeyObservedState is the state of the IAMServiceAccountKey resource as most recently observed in GCP.
// +kcc:proto=google.iam.admin.v1.ServiceAccountKey
type IAMServiceAccountKeyObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpiamserviceaccountkey;gcpiamserviceaccountkeys
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true";"cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// IAMServiceAccountKey is the Schema for the IAMServiceAccountKey API
// +k8s:openapi-gen=true
type IAMServiceAccountKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   IAMServiceAccountKeySpec   `json:"spec,omitempty"`
	Status IAMServiceAccountKeyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// IAMServiceAccountKeyList contains a list of IAMServiceAccountKey
type IAMServiceAccountKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IAMServiceAccountKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IAMServiceAccountKey{}, &IAMServiceAccountKeyList{})
}
