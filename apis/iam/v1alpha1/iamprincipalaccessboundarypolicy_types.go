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

// API sources for IAMPrincipalAccessBoundaryPolicy, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/iam/v3/principal_access_boundary_policy_resources.proto
// +kcc:source:service-docs=https://cloud.google.com/iam/
// +kcc:source:resource-docs=https://docs.cloud.google.com/iam/docs/reference/rest/v3/organizations.locations.principalAccessBoundaryPolicies

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var IAMPrincipalAccessBoundaryPolicyGVK = GroupVersion.WithKind("IAMPrincipalAccessBoundaryPolicy")

// IAMPrincipalAccessBoundaryPolicySpec defines the desired state of IAMPrincipalAccessBoundaryPolicy
// +kcc:spec:proto=google.iam.v3.PrincipalAccessBoundaryPolicy
// +kcc:required-from-proto
type IAMPrincipalAccessBoundaryPolicySpec struct {
	// The organization that this resource belongs to.
	OrganizationRef *refsv1beta1.OrganizationRef `json:"organizationRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The IAMPrincipalAccessBoundaryPolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. The etag for the principal access boundary.
	//  If this is provided on update, it must match the server's etag.
	// +kcc:proto:field=google.iam.v3.PrincipalAccessBoundaryPolicy.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. The description of the principal access boundary policy. Must be
	//  less than or equal to 63 characters.
	// +kcc:proto:field=google.iam.v3.PrincipalAccessBoundaryPolicy.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User defined annotations. See
	//  https://google.aip.dev/148#annotations for more details such as format and
	//  size limitations
	// +kcc:proto:field=google.iam.v3.PrincipalAccessBoundaryPolicy.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. The details for the principal access boundary policy.
	// +kcc:proto:field=google.iam.v3.PrincipalAccessBoundaryPolicy.details
	Details *PrincipalAccessBoundaryPolicyDetails `json:"details,omitempty"`
}

// IAMPrincipalAccessBoundaryPolicyStatus defines the config connector machine state of IAMPrincipalAccessBoundaryPolicy
type IAMPrincipalAccessBoundaryPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the IAMPrincipalAccessBoundaryPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *IAMPrincipalAccessBoundaryPolicyObservedState `json:"observedState,omitempty"`
}

// IAMPrincipalAccessBoundaryPolicyObservedState is the state of the IAMPrincipalAccessBoundaryPolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.iam.v3.PrincipalAccessBoundaryPolicy
type IAMPrincipalAccessBoundaryPolicyObservedState struct {
	// Output only. The globally unique ID of the principal access boundary
	//  policy.
	// +kcc:proto:field=google.iam.v3.PrincipalAccessBoundaryPolicy.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the principal access boundary policy was
	//  created.
	// +kcc:proto:field=google.iam.v3.PrincipalAccessBoundaryPolicy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the principal access boundary policy was most
	//  recently updated.
	// +kcc:proto:field=google.iam.v3.PrincipalAccessBoundaryPolicy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpiamprincipalaccessboundarypolicy;gcpiamprincipalaccessboundarypolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// IAMPrincipalAccessBoundaryPolicy is the Schema for the IAMPrincipalAccessBoundaryPolicy API
// +k8s:openapi-gen=true
type IAMPrincipalAccessBoundaryPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   IAMPrincipalAccessBoundaryPolicySpec   `json:"spec,omitempty"`
	Status IAMPrincipalAccessBoundaryPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// IAMPrincipalAccessBoundaryPolicyList contains a list of IAMPrincipalAccessBoundaryPolicy
type IAMPrincipalAccessBoundaryPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IAMPrincipalAccessBoundaryPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IAMPrincipalAccessBoundaryPolicy{}, &IAMPrincipalAccessBoundaryPolicyList{})
}
