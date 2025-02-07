// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1


// +kcc:proto=google.cloud.binaryauthorization.v1.AdmissionRule
type AdmissionRule struct {
	// Required. How this admission rule will be evaluated.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.AdmissionRule.evaluation_mode
	EvaluationMode *string `json:"evaluationMode,omitempty"`

	// Optional. The resource names of the attestors that must attest to
	//  a container image, in the format `projects/*/attestors/*`. Each
	//  attestor must exist before a policy can reference it.  To add an attestor
	//  to a policy the principal issuing the policy change request must be able
	//  to read the attestor resource.
	//
	//  Note: this field must be non-empty when the evaluation_mode field specifies
	//  REQUIRE_ATTESTATION, otherwise it must be empty.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.AdmissionRule.require_attestations_by
	RequireAttestationsBy []string `json:"requireAttestationsBy,omitempty"`

	// Required. The action when a pod creation is denied by the admission rule.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.AdmissionRule.enforcement_mode
	EnforcementMode *string `json:"enforcementMode,omitempty"`
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

// +kcc:proto=google.cloud.binaryauthorization.v1.Policy
type Policy struct {

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

	// TODO: unsupported map type with key string and value message


	// TODO: unsupported map type with key string and value message


	// TODO: unsupported map type with key string and value message


	// TODO: unsupported map type with key string and value message


	// Required. Default admission rule for a cluster without a per-cluster, per-
	//  kubernetes-service-account, or per-istio-service-identity admission rule.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.Policy.default_admission_rule
	DefaultAdmissionRule *AdmissionRule `json:"defaultAdmissionRule,omitempty"`
}

// +kcc:proto=google.cloud.binaryauthorization.v1.Policy
type PolicyObservedState struct {
	// Output only. The resource name, in the format `projects/*/policy`. There is
	//  at most one policy per project.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.Policy.name
	Name *string `json:"name,omitempty"`

	// Output only. Time when the policy was last updated.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.Policy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
