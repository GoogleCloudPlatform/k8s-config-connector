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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ParameterManagerParameterGVK = GroupVersion.WithKind("ParameterManagerParameter")

// ParameterManagerParameterSpec defines the desired state of ParameterManagerParameter
// +kcc:spec:proto=google.cloud.parametermanager.v1.Parameter
type ParameterManagerParameterSpec struct {
	// Required. Defines the parent path of the resource.
	*parent.ProjectAndLocationRef `json:",inline"`

	// The ParameterManagerParameter name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Specifies the format of a Parameter.
	// +kcc:proto:field=google.cloud.parametermanager.v1.Parameter.format
	// +kubebuilder:validation:Enum=UNFORMATTED;YAML;JSON
	// +kubebuilder:default=UNFORMATTED
	Format *string `json:"format,omitempty"`

	// Optional. Customer managed encryption key (CMEK) to use for encrypting the
	//  Parameter Versions. If not set, the default Google-managed encryption key
	//  will be used. Cloud KMS CryptoKeys must reside in the same location as the
	//  Parameter. The expected format is
	//  `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
	// +kcc:proto:field=google.cloud.parametermanager.v1.Parameter.kms_key
	KMSKeyRef *refs.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// ParameterManagerParameterStatus defines the config connector machine state of ParameterManagerParameter
type ParameterManagerParameterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ParameterManagerParameter resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ParameterManagerParameterObservedState `json:"observedState,omitempty"`
}

// ParameterManagerParameterObservedState is the state of the ParameterManagerParameter resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.parametermanager.v1.Parameter
type ParameterManagerParameterObservedState struct {
	// Identifier. [Output only] The resource name of the Parameter in the format
	// `projects/*/locations/*/parameters/*`
	// +kcc:proto:field=google.cloud.parametermanager.v1.Parameter.name
	Name *string `json:"name,omitempty"`

	// Output only. [Output only] Create time stamp
	// +kcc:proto:field=google.cloud.parametermanager.v1.Parameter.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. [Output only] Update time stamp
	// +kcc:proto:field=google.cloud.parametermanager.v1.Parameter.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. [Output-only] policy member strings of a Google Cloud
	//  resource.
	// +kcc:proto:field=google.cloud.parametermanager.v1.Parameter.policy_member
	PolicyMember *ResourcePolicyMemberObservedState `json:"policyMember,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpparametermanagerparameter;gcpparametermanagerparameters
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ParameterManagerParameter is the Schema for the ParameterManagerParameter API
// +k8s:openapi-gen=true
type ParameterManagerParameter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ParameterManagerParameterSpec   `json:"spec,omitempty"`
	Status ParameterManagerParameterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ParameterManagerParameterList contains a list of ParameterManagerParameter
type ParameterManagerParameterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ParameterManagerParameter `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ParameterManagerParameter{}, &ParameterManagerParameterList{})
}
