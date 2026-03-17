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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BinaryAuthorizationPolicyGVK = GroupVersion.WithKind("BinaryAuthorizationPolicy")

// BinaryAuthorizationPolicySpec defines the desired state of BinaryAuthorizationPolicy
// +kcc:spec:proto=google.cloud.binaryauthorization.v1.Policy
type BinaryAuthorizationPolicySpec struct {
	// Immutable. The Project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Optional. A descriptive comment.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.Policy.description
	Description *string `json:"description,omitempty"`

	// Optional. Controls the evaluation of a Google-maintained global admission
	//  policy for common system-level images. Images not covered by the global
	//  policy will be subject to the project admission policy. This setting
	//  has no effect when specified inside a global admission policy.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.Policy.global_policy_evaluation_mode
	GlobalPolicyEvaluationMode *string `json:"globalPolicyEvaluationMode,omitempty"`

	// Optional. Admission policy allowlisting. A matching admission request will
	//  always be permitted. This feature is typically used to exclude Google or
	//  third-party infrastructure images from Binary Authorization policies.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.Policy.admission_whitelist_patterns
	AdmissionWhitelistPatterns []AdmissionWhitelistPattern `json:"admissionWhitelistPatterns,omitempty"`

	// Optional. Per-cluster admission rules. Cluster spec
	//  format: location.clusterId. There can be at most one admission rule
	//  per cluster spec. A location is either a compute zone (e.g. us-central1-a)
	//  or a region (e.g. us-central1). For clusterId syntax restrictions
	//  see https://cloud.google.com/container-engine/reference/rest/v1/projects.zones.clusters.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.Policy.cluster_admission_rules
	ClusterAdmissionRules map[string]AdmissionRule `json:"clusterAdmissionRules,omitempty"`

	// Optional. Per-kubernetes-namespace admission rules.
	//  K8s namespace spec format: [a-z.-]+, e.g. 'some-namespace'
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.Policy.kubernetes_namespace_admission_rules
	KubernetesNamespaceAdmissionRules map[string]AdmissionRule `json:"kubernetesNamespaceAdmissionRules,omitempty"`

	// Optional. Per-kubernetes-service-account admission rules.
	//  Service account spec format: namespace:serviceaccount. e.g. 'test-ns:default'
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.Policy.kubernetes_service_account_admission_rules
	KubernetesServiceAccountAdmissionRules map[string]AdmissionRule `json:"kubernetesServiceAccountAdmissionRules,omitempty"`

	// Optional. Per-istio-service-identity admission rules.
	//  Istio service identity spec format: spiffe:///ns//sa/ or /ns//sa/
	//  e.g. spiffe://example.com/ns/test-ns/sa/default
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.Policy.istio_service_identity_admission_rules
	IstioServiceIdentityAdmissionRules map[string]AdmissionRule `json:"istioServiceIdentityAdmissionRules,omitempty"`

	// Required. Default admission rule for a cluster without a per-cluster, per-
	//  kubernetes-service-account, or per-istio-service-identity admission rule.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.Policy.default_admission_rule
	// +required
	DefaultAdmissionRule *AdmissionRule `json:"defaultAdmissionRule"`
}

// +kcc:proto=google.cloud.binaryauthorization.v1.AdmissionRule
type AdmissionRule struct {
	// Required. How this admission rule will be evaluated.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.AdmissionRule.evaluation_mode
	// +required
	EvaluationMode *string `json:"evaluationMode"`

	// Optional. The resource names of the attestors that must attest to
	//  a container image, in the format `projects/*/attestors/*`. Each
	//  attestor must exist before a policy can reference it.  To add an attestor
	//  to a policy the principal issuing the policy change request must be able
	//  to read the attestor resource.
	//
	//  Note: this field must be non-empty when the evaluation_mode field specifies
	//  REQUIRE_ATTESTATION, otherwise it must be empty.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.AdmissionRule.require_attestations_by
	RequireAttestationsBy []refsv1beta1.BinaryAuthorizationAttestorRef `json:"requireAttestationsBy,omitempty"`

	// Required. The action when a pod creation is denied by the admission rule.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.AdmissionRule.enforcement_mode
	// +required
	EnforcementMode *string `json:"enforcementMode"`
}

// +kcc:proto=google.cloud.binaryauthorization.v1.AdmissionWhitelistPattern
type AdmissionWhitelistPattern struct {
	// An image name pattern to allowlist, in the form `registry/path/to/image`.
	//  This supports a trailing `*` wildcard, but this is allowed only in
	//  text after the `registry/` part. This also supports a trailing `**`
	//  wildcard which matches subdirectories of a given entry.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.AdmissionWhitelistPattern.name_pattern
	NamePattern *string `json:"namePattern,omitempty"`
}

// BinaryAuthorizationPolicyStatus defines the config connector machine state of BinaryAuthorizationPolicy
type BinaryAuthorizationPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// Output only. The resource name, in the format `projects/*/policy`. There is
	//  at most one policy per project.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.Policy.name
	SelfLink *string `json:"selfLink,omitempty"`

	// Output only. Time when the policy was last updated.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.Policy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbinaryauthorizationpolicy;gcpbinaryauthorizationpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BinaryAuthorizationPolicy is the Schema for the BinaryAuthorizationPolicy API
// +k8s:openapi-gen=true
type BinaryAuthorizationPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BinaryAuthorizationPolicySpec   `json:"spec,omitempty"`
	Status BinaryAuthorizationPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BinaryAuthorizationPolicyList contains a list of BinaryAuthorizationPolicy
type BinaryAuthorizationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BinaryAuthorizationPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BinaryAuthorizationPolicy{}, &BinaryAuthorizationPolicyList{})
}
