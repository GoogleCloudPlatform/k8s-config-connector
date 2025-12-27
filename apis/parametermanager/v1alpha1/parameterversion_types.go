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

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ParameterManagerParameterVersionGVK = GroupVersion.WithKind("ParameterManagerParameterVersion")

// ParameterManagerParameterVersionSpec defines the desired state of ParameterManagerParameterVersion
// +kcc:spec:proto=google.cloud.parametermanager.v1.ParameterVersion
type ParameterManagerParameterVersionSpec struct {
	// The resource name of the [Parameter][google.cloud.parametermanager.v1.Parameter] to create a [ParameterVersion][google.cloud.parametermanager.v1.ParameterVersion] for.
	ParameterRef *ParameterRef `json:"parameterRef,omitempty"`

	// The ParameterManagerParameterVersion name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Disabled boolean to determine if a ParameterVersion acts as a
	//  metadata only resource (payload is never returned if disabled is true). If
	//  true any calls will always default to BASIC view even if the user
	//  explicitly passes FULL view as part of the request. A render call on a
	//  disabled resource fails with an error. Default value is False.
	// +kcc:proto:field=google.cloud.parametermanager.v1.ParameterVersion.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// Required. Immutable. Payload content of a ParameterVersion resource.  This
	//  is only returned when the request provides the View value of FULL (default
	//  for GET request).
	// +kcc:proto:field=google.cloud.parametermanager.v1.ParameterVersion.payload
	Payload *ParameterVersionPayload `json:"payload,omitempty"`
}

// ParameterManagerParameterVersionStatus defines the config connector machine state of ParameterManagerParameterVersion
type ParameterManagerParameterVersionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ParameterManagerParameterVersion resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ParameterManagerParameterVersionObservedState `json:"observedState,omitempty"`
}

// ParameterManagerParameterVersionObservedState is the state of the ParameterManagerParameterVersion resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.parametermanager.v1.ParameterVersion
type ParameterManagerParameterVersionObservedState struct {
	// Identifier. [Output only] The resource name of the ParameterVersion in the
	//  format `projects/*/locations/*/parameters/*/versions/*`.
	// +kcc:proto:field=google.cloud.parametermanager.v1.ParameterVersion.name
	Name *string `json:"name,omitempty"`

	// Output only. [Output only] Create time stamp
	// +kcc:proto:field=google.cloud.parametermanager.v1.ParameterVersion.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. [Output only] Update time stamp
	// +kcc:proto:field=google.cloud.parametermanager.v1.ParameterVersion.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. Output only. [Output only] The resource name of the KMS key
	//  version used to encrypt the ParameterVersion payload. This field is
	//  populated only if the Parameter resource has customer managed encryption
	//  key (CMEK) configured.
	// +kcc:proto:field=google.cloud.parametermanager.v1.ParameterVersion.kms_key_version
	KMSKeyVersion *refs.KMSCryptoKeyRef `json:"kmsKeyVersion,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpparametermanagerparameterversion;gcpparametermanagerparameterversions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ParameterManagerParameterVersion is the Schema for the ParameterManagerParameterVersion API
// +k8s:openapi-gen=true
type ParameterManagerParameterVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ParameterManagerParameterVersionSpec   `json:"spec,omitempty"`
	Status ParameterManagerParameterVersionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ParameterManagerParameterVersionList contains a list of ParameterManagerParameterVersion
type ParameterManagerParameterVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ParameterManagerParameterVersion `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ParameterManagerParameterVersion{}, &ParameterManagerParameterVersionList{})
}
