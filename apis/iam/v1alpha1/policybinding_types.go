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
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var IAMPolicyBindingGVK = GroupVersion.WithKind("IAMPolicyBinding")

type Parent struct {
	// +required
	Location string `json:"location"`
	// +optional
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef,omitempty"`
	// +optional
	OrganizationRef *refv1beta1.OrganizationRef `json:"organizationRef,omitempty"`
	// +optional
	FolderRef *refv1beta1.FolderRef `json:"folderRef,omitempty"`
}

// IAMPolicyBindingSpec defines the desired state of IAMPolicyBinding
// +kcc:proto=google.iam.v3.PolicyBinding
type IAMPolicyBindingSpec struct {
	Parent `json:",inline"`
	// The IAMPolicyBinding name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. The etag for the policy binding.
	//  If this is provided on update, it must match the server's etag.
	// +kcc:proto:field=google.iam.v3.PolicyBinding.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. The description of the policy binding. Must be less than or equal
	//  to 63 characters.
	// +kcc:proto:field=google.iam.v3.PolicyBinding.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User defined annotations. See
	//  https://google.aip.dev/148#annotations for more details such as format and
	//  size limitations
	// +kcc:proto:field=google.iam.v3.PolicyBinding.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Required. Immutable. Target is the full resource name of the resource to
	//  which the policy will be bound. Immutable once set.
	// +kcc:proto:field=google.iam.v3.PolicyBinding.target
	//+required
	Target *PolicyBinding_Target `json:"target,omitempty"`

	// Immutable. The kind of the policy to attach in this binding. This field
	//  must be one of the following:
	//
	//  - Left empty (will be automatically set to the policy kind)
	//  - The input policy kind
	// +kcc:proto:field=google.iam.v3.PolicyBinding.policy_kind
	PolicyKind *string `json:"policyKind,omitempty"`

	// Required. Immutable. The resource name of the policy to be bound. The
	//  binding parent and policy must belong to the same Organization (or
	//  Project).
	// +kcc:proto:field=google.iam.v3.PolicyBinding.policy
	//+required
	Policy *string `json:"policy,omitempty"`

	// Optional. Condition can either be a principal condition or a resource
	//  condition. It depends on the type of target, the policy it is attached to,
	//  and/or the expression itself. When set, the `expression` field in the
	//  `Expr` must include from 1 to 10 subexpressions, joined by the "||"(Logical
	//  OR),
	//  "&&"(Logical AND) or "!"(Logical NOT) operators and cannot contain more
	//  than 250 characters.
	//  Allowed operations for principal.subject:
	//
	//  - `principal.subject == <principal subject string>`
	//  - `principal.subject != <principal subject string>`
	//  - `principal.subject in [<list of principal subjects>]`
	//  - `principal.subject.startsWith(<string>)`
	//  - `principal.subject.endsWith(<string>)`
	//
	//  Allowed operations for principal.type:
	//
	//  - `principal.type == <principal type string>`
	//  - `principal.type != <principal type string>`
	//  - `principal.type in [<list of principal types>]`
	//
	//  Supported principal types are Workspace, Workforce Pool, Workload Pool and
	//  Service Account. Allowed string must be one of:
	//
	//  - iam.googleapis.com/WorkspaceIdentity
	//  - iam.googleapis.com/WorkforcePoolIdentity
	//  - iam.googleapis.com/WorkloadPoolIdentity
	//  - iam.googleapis.com/ServiceAccount
	//
	//  When the bound policy is a principal access boundary policy, the only
	//  supported attributes in any subexpression are `principal.type` and
	//  `principal.subject`. An example expression is: "principal.type ==
	//  'iam.googleapis.com/ServiceAccount'" or "principal.subject ==
	//  'bob@example.com'".
	// +kcc:proto:field=google.iam.v3.PolicyBinding.condition
	Condition *Expr `json:"condition,omitempty"`
}

// IAMPolicyBindingStatus defines the config connector machine state of IAMPolicyBinding
type IAMPolicyBindingStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the IAMPolicyBinding resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *IAMPolicyBindingObservedState `json:"observedState,omitempty"`
}

// IAMPolicyBindingObservedState is the state of the IAMPolicyBinding resource as most recently observed in GCP.
// +kcc:proto=google.iam.v3.PolicyBinding
type IAMPolicyBindingObservedState struct {
	// Output only. The globally unique ID of the policy binding. Assigned when
	//  the policy binding is created.
	// +kcc:proto:field=google.iam.v3.PolicyBinding.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The globally unique ID of the policy to be bound.
	// +kcc:proto:field=google.iam.v3.PolicyBinding.policy_uid
	PolicyUid *string `json:"policyUid,omitempty"`

	// Output only. The time when the policy binding was created.
	// +kcc:proto:field=google.iam.v3.PolicyBinding.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the policy binding was most recently updated.
	// +kcc:proto:field=google.iam.v3.PolicyBinding.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpiampolicybinding;gcpiampolicybindings
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// IAMPolicyBinding is the Schema for the IAMPolicyBinding API
// +k8s:openapi-gen=true
type IAMPolicyBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   IAMPolicyBindingSpec   `json:"spec,omitempty"`
	Status IAMPolicyBindingStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// IAMPolicyBindingList contains a list of IAMPolicyBinding
type IAMPolicyBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IAMPolicyBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IAMPolicyBinding{}, &IAMPolicyBindingList{})
}
