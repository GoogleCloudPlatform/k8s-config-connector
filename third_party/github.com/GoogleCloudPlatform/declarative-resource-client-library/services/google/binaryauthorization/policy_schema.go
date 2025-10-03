// Copyright 2024 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package binaryauthorization

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLPolicySchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "BinaryAuthorization/Policy",
			Description: "The BinaryAuthorization Policy resource",
			StructName:  "Policy",
			HasIAM:      true,
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Policy",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "policy",
						Required:    true,
						Description: "A full instance of a Policy",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Policy",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "policy",
						Required:    true,
						Description: "A full instance of a Policy",
					},
				},
			},
		},
		Components: &dcl.Components{
			Schemas: map[string]*dcl.Component{
				"Policy": &dcl.Component{
					Title:           "Policy",
					ID:              "projects/{{project}}/policy",
					ParentContainer: "project",
					HasIAM:          true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"defaultAdmissionRule",
						},
						Properties: map[string]*dcl.Property{
							"admissionWhitelistPatterns": &dcl.Property{
								Type:        "array",
								GoName:      "AdmissionWhitelistPatterns",
								Description: "Optional. Admission policy allowlisting. A matching admission request will always be permitted. This feature is typically used to exclude Google or third-party infrastructure images from Binary Authorization policies.",
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "PolicyAdmissionWhitelistPatterns",
									Properties: map[string]*dcl.Property{
										"namePattern": &dcl.Property{
											Type:        "string",
											GoName:      "NamePattern",
											Description: "An image name pattern to allowlist, in the form `registry/path/to/image`. This supports a trailing `*` as a wildcard, but this is allowed only in text after the `registry/` part.",
										},
									},
								},
							},
							"clusterAdmissionRules": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type:   "object",
									GoType: "PolicyClusterAdmissionRules",
									Required: []string{
										"evaluationMode",
										"enforcementMode",
									},
									Properties: map[string]*dcl.Property{
										"enforcementMode": &dcl.Property{
											Type:        "string",
											GoName:      "EnforcementMode",
											GoType:      "PolicyClusterAdmissionRulesEnforcementModeEnum",
											Description: "Required. The action when a pod creation is denied by the admission rule. Possible values: ENFORCEMENT_MODE_UNSPECIFIED, ENFORCED_BLOCK_AND_AUDIT_LOG, DRYRUN_AUDIT_LOG_ONLY",
											Enum: []string{
												"ENFORCEMENT_MODE_UNSPECIFIED",
												"ENFORCED_BLOCK_AND_AUDIT_LOG",
												"DRYRUN_AUDIT_LOG_ONLY",
											},
										},
										"evaluationMode": &dcl.Property{
											Type:        "string",
											GoName:      "EvaluationMode",
											GoType:      "PolicyClusterAdmissionRulesEvaluationModeEnum",
											Description: "Required. How this admission rule will be evaluated. Possible values: ALWAYS_ALLOW, ALWAYS_DENY, REQUIRE_ATTESTATION",
											Enum: []string{
												"ALWAYS_ALLOW",
												"ALWAYS_DENY",
												"REQUIRE_ATTESTATION",
											},
										},
										"requireAttestationsBy": &dcl.Property{
											Type:        "array",
											GoName:      "RequireAttestationsBy",
											Description: "Optional. The resource names of the attestors that must attest to a container image, in the format `projects/*/attestors/*`. Each attestor must exist before a policy can reference it. To add an attestor to a policy the principal issuing the policy change request must be able to read the attestor resource. Note: this field must be non-empty when the evaluation_mode field specifies REQUIRE_ATTESTATION, otherwise it must be empty.",
											SendEmpty:   true,
											ListType:    "list",
											Items: &dcl.Property{
												Type:   "string",
												GoType: "string",
												ResourceReferences: []*dcl.PropertyResourceReference{
													&dcl.PropertyResourceReference{
														Resource: "Binaryauthorization/Attestor",
														Field:    "name",
													},
												},
											},
										},
									},
								},
								GoName:      "ClusterAdmissionRules",
								Description: "Optional. Per-cluster admission rules. Cluster spec format: location.clusterId. There can be at most one admission rule per cluster spec. A location is either a compute zone (e.g. us-central1-a) or a region (e.g. us-central1). For clusterId syntax restrictions see https://cloud.google.com/container-engine/reference/rest/v1/projects.zones.clusters.",
								Conflicts: []string{
									"kubernetesNamespaceAdmissionRules",
									"kubernetesServiceAccountAdmissionRules",
									"istioServiceIdentityAdmissionRules",
								},
							},
							"defaultAdmissionRule": &dcl.Property{
								Type:        "object",
								GoName:      "DefaultAdmissionRule",
								GoType:      "PolicyDefaultAdmissionRule",
								Description: "Required. Default admission rule for a cluster without a per-cluster, per-kubernetes-service-account, or per-istio-service-identity admission rule.",
								Required: []string{
									"evaluationMode",
									"enforcementMode",
								},
								Properties: map[string]*dcl.Property{
									"enforcementMode": &dcl.Property{
										Type:        "string",
										GoName:      "EnforcementMode",
										GoType:      "PolicyDefaultAdmissionRuleEnforcementModeEnum",
										Description: "Required. The action when a pod creation is denied by the admission rule. Possible values: ENFORCEMENT_MODE_UNSPECIFIED, ENFORCED_BLOCK_AND_AUDIT_LOG, DRYRUN_AUDIT_LOG_ONLY",
										Enum: []string{
											"ENFORCEMENT_MODE_UNSPECIFIED",
											"ENFORCED_BLOCK_AND_AUDIT_LOG",
											"DRYRUN_AUDIT_LOG_ONLY",
										},
									},
									"evaluationMode": &dcl.Property{
										Type:        "string",
										GoName:      "EvaluationMode",
										GoType:      "PolicyDefaultAdmissionRuleEvaluationModeEnum",
										Description: "Required. How this admission rule will be evaluated. Possible values: ALWAYS_ALLOW, ALWAYS_DENY, REQUIRE_ATTESTATION",
										Enum: []string{
											"ALWAYS_ALLOW",
											"ALWAYS_DENY",
											"REQUIRE_ATTESTATION",
										},
									},
									"requireAttestationsBy": &dcl.Property{
										Type:        "array",
										GoName:      "RequireAttestationsBy",
										Description: "Optional. The resource names of the attestors that must attest to a container image, in the format `projects/*/attestors/*`. Each attestor must exist before a policy can reference it. To add an attestor to a policy the principal issuing the policy change request must be able to read the attestor resource. Note: this field must be non-empty when the evaluation_mode field specifies REQUIRE_ATTESTATION, otherwise it must be empty.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "string",
											GoType: "string",
											ResourceReferences: []*dcl.PropertyResourceReference{
												&dcl.PropertyResourceReference{
													Resource: "Binaryauthorization/Attestor",
													Field:    "name",
												},
											},
										},
									},
								},
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "Optional. A descriptive comment.",
							},
							"globalPolicyEvaluationMode": &dcl.Property{
								Type:        "string",
								GoName:      "GlobalPolicyEvaluationMode",
								GoType:      "PolicyGlobalPolicyEvaluationModeEnum",
								Description: "Optional. Controls the evaluation of a Google-maintained global admission policy for common system-level images. Images not covered by the global policy will be subject to the project admission policy. This setting has no effect when specified inside a global admission policy. Possible values: GLOBAL_POLICY_EVALUATION_MODE_UNSPECIFIED, ENABLE, DISABLE",
								Enum: []string{
									"GLOBAL_POLICY_EVALUATION_MODE_UNSPECIFIED",
									"ENABLE",
									"DISABLE",
								},
							},
							"istioServiceIdentityAdmissionRules": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type:   "object",
									GoType: "PolicyIstioServiceIdentityAdmissionRules",
									Required: []string{
										"evaluationMode",
										"enforcementMode",
									},
									Properties: map[string]*dcl.Property{
										"enforcementMode": &dcl.Property{
											Type:        "string",
											GoName:      "EnforcementMode",
											GoType:      "PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum",
											Description: "Required. The action when a pod creation is denied by the admission rule. Possible values: ENFORCEMENT_MODE_UNSPECIFIED, ENFORCED_BLOCK_AND_AUDIT_LOG, DRYRUN_AUDIT_LOG_ONLY",
											Enum: []string{
												"ENFORCEMENT_MODE_UNSPECIFIED",
												"ENFORCED_BLOCK_AND_AUDIT_LOG",
												"DRYRUN_AUDIT_LOG_ONLY",
											},
										},
										"evaluationMode": &dcl.Property{
											Type:        "string",
											GoName:      "EvaluationMode",
											GoType:      "PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum",
											Description: "Required. How this admission rule will be evaluated. Possible values: ALWAYS_ALLOW, ALWAYS_DENY, REQUIRE_ATTESTATION",
											Enum: []string{
												"ALWAYS_ALLOW",
												"ALWAYS_DENY",
												"REQUIRE_ATTESTATION",
											},
										},
										"requireAttestationsBy": &dcl.Property{
											Type:        "array",
											GoName:      "RequireAttestationsBy",
											Description: "Optional. The resource names of the attestors that must attest to a container image, in the format `projects/*/attestors/*`. Each attestor must exist before a policy can reference it. To add an attestor to a policy the principal issuing the policy change request must be able to read the attestor resource. Note: this field must be non-empty when the evaluation_mode field specifies REQUIRE_ATTESTATION, otherwise it must be empty.",
											SendEmpty:   true,
											ListType:    "list",
											Items: &dcl.Property{
												Type:   "string",
												GoType: "string",
												ResourceReferences: []*dcl.PropertyResourceReference{
													&dcl.PropertyResourceReference{
														Resource: "Binaryauthorization/Attestor",
														Field:    "name",
													},
												},
											},
										},
									},
								},
								GoName:      "IstioServiceIdentityAdmissionRules",
								Description: "Optional. Per-istio-service-identity admission rules. Istio service identity spec format: spiffe:///ns//sa/ or /ns//sa/ e.g. spiffe://example.com/ns/test-ns/sa/default",
								Conflicts: []string{
									"kubernetesNamespaceAdmissionRules",
									"kubernetesServiceAccountAdmissionRules",
									"clusterAdmissionRules",
								},
							},
							"kubernetesNamespaceAdmissionRules": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type:   "object",
									GoType: "PolicyKubernetesNamespaceAdmissionRules",
									Required: []string{
										"evaluationMode",
										"enforcementMode",
									},
									Properties: map[string]*dcl.Property{
										"enforcementMode": &dcl.Property{
											Type:        "string",
											GoName:      "EnforcementMode",
											GoType:      "PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum",
											Description: "Required. The action when a pod creation is denied by the admission rule. Possible values: ENFORCEMENT_MODE_UNSPECIFIED, ENFORCED_BLOCK_AND_AUDIT_LOG, DRYRUN_AUDIT_LOG_ONLY",
											Enum: []string{
												"ENFORCEMENT_MODE_UNSPECIFIED",
												"ENFORCED_BLOCK_AND_AUDIT_LOG",
												"DRYRUN_AUDIT_LOG_ONLY",
											},
										},
										"evaluationMode": &dcl.Property{
											Type:        "string",
											GoName:      "EvaluationMode",
											GoType:      "PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum",
											Description: "Required. How this admission rule will be evaluated. Possible values: ALWAYS_ALLOW, ALWAYS_DENY, REQUIRE_ATTESTATION",
											Enum: []string{
												"ALWAYS_ALLOW",
												"ALWAYS_DENY",
												"REQUIRE_ATTESTATION",
											},
										},
										"requireAttestationsBy": &dcl.Property{
											Type:        "array",
											GoName:      "RequireAttestationsBy",
											Description: "Optional. The resource names of the attestors that must attest to a container image, in the format `projects/*/attestors/*`. Each attestor must exist before a policy can reference it. To add an attestor to a policy the principal issuing the policy change request must be able to read the attestor resource. Note: this field must be non-empty when the evaluation_mode field specifies REQUIRE_ATTESTATION, otherwise it must be empty.",
											SendEmpty:   true,
											ListType:    "list",
											Items: &dcl.Property{
												Type:   "string",
												GoType: "string",
												ResourceReferences: []*dcl.PropertyResourceReference{
													&dcl.PropertyResourceReference{
														Resource: "Binaryauthorization/Attestor",
														Field:    "name",
													},
												},
											},
										},
									},
								},
								GoName:      "KubernetesNamespaceAdmissionRules",
								Description: "Optional. Per-kubernetes-namespace admission rules. K8s namespace spec format: [a-z.-]+, e.g. 'some-namespace'",
								Conflicts: []string{
									"kubernetesServiceAccountAdmissionRules",
									"istioServiceIdentityAdmissionRules",
									"clusterAdmissionRules",
								},
							},
							"kubernetesServiceAccountAdmissionRules": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type:   "object",
									GoType: "PolicyKubernetesServiceAccountAdmissionRules",
									Required: []string{
										"evaluationMode",
										"enforcementMode",
									},
									Properties: map[string]*dcl.Property{
										"enforcementMode": &dcl.Property{
											Type:        "string",
											GoName:      "EnforcementMode",
											GoType:      "PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum",
											Description: "Required. The action when a pod creation is denied by the admission rule. Possible values: ENFORCEMENT_MODE_UNSPECIFIED, ENFORCED_BLOCK_AND_AUDIT_LOG, DRYRUN_AUDIT_LOG_ONLY",
											Enum: []string{
												"ENFORCEMENT_MODE_UNSPECIFIED",
												"ENFORCED_BLOCK_AND_AUDIT_LOG",
												"DRYRUN_AUDIT_LOG_ONLY",
											},
										},
										"evaluationMode": &dcl.Property{
											Type:        "string",
											GoName:      "EvaluationMode",
											GoType:      "PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum",
											Description: "Required. How this admission rule will be evaluated. Possible values: ALWAYS_ALLOW, ALWAYS_DENY, REQUIRE_ATTESTATION",
											Enum: []string{
												"ALWAYS_ALLOW",
												"ALWAYS_DENY",
												"REQUIRE_ATTESTATION",
											},
										},
										"requireAttestationsBy": &dcl.Property{
											Type:        "array",
											GoName:      "RequireAttestationsBy",
											Description: "Optional. The resource names of the attestors that must attest to a container image, in the format `projects/*/attestors/*`. Each attestor must exist before a policy can reference it. To add an attestor to a policy the principal issuing the policy change request must be able to read the attestor resource. Note: this field must be non-empty when the evaluation_mode field specifies REQUIRE_ATTESTATION, otherwise it must be empty.",
											SendEmpty:   true,
											ListType:    "list",
											Items: &dcl.Property{
												Type:   "string",
												GoType: "string",
												ResourceReferences: []*dcl.PropertyResourceReference{
													&dcl.PropertyResourceReference{
														Resource: "Binaryauthorization/Attestor",
														Field:    "name",
													},
												},
											},
										},
									},
								},
								GoName:      "KubernetesServiceAccountAdmissionRules",
								Description: "Optional. Per-kubernetes-service-account admission rules. Service account spec format: namespace:serviceaccount. e.g. 'test-ns:default'",
								Conflicts: []string{
									"kubernetesNamespaceAdmissionRules",
									"istioServiceIdentityAdmissionRules",
									"clusterAdmissionRules",
								},
							},
							"project": &dcl.Property{
								Type:        "string",
								GoName:      "Project",
								Description: "The project of the resource.",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Project",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"selfLink": &dcl.Property{
								Type:        "string",
								GoName:      "SelfLink",
								ReadOnly:    true,
								Description: "Output only. The resource name, in the format `projects/*/policy`. There is at most one policy per project.",
								Immutable:   true,
							},
							"updateTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "UpdateTime",
								ReadOnly:    true,
								Description: "Output only. Time when the policy was last updated.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
