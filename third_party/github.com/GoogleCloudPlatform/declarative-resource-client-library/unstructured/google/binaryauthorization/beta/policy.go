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
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	iamUnstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/iam"
)

type Policy struct{}

func PolicyToUnstructured(r *dclService.Policy) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "binaryauthorization",
			Version: "beta",
			Type:    "Policy",
		},
		Object: make(map[string]interface{}),
	}
	var rAdmissionWhitelistPatterns []interface{}
	for _, rAdmissionWhitelistPatternsVal := range r.AdmissionWhitelistPatterns {
		rAdmissionWhitelistPatternsObject := make(map[string]interface{})
		if rAdmissionWhitelistPatternsVal.NamePattern != nil {
			rAdmissionWhitelistPatternsObject["namePattern"] = *rAdmissionWhitelistPatternsVal.NamePattern
		}
		rAdmissionWhitelistPatterns = append(rAdmissionWhitelistPatterns, rAdmissionWhitelistPatternsObject)
	}
	u.Object["admissionWhitelistPatterns"] = rAdmissionWhitelistPatterns
	if r.ClusterAdmissionRules != nil {
		rClusterAdmissionRules := make(map[string]interface{})
		for k, v := range r.ClusterAdmissionRules {
			rClusterAdmissionRulesMap := make(map[string]interface{})
			if v.EnforcementMode != nil {
				rClusterAdmissionRulesMap["enforcementMode"] = string(*v.EnforcementMode)
			}
			if v.EvaluationMode != nil {
				rClusterAdmissionRulesMap["evaluationMode"] = string(*v.EvaluationMode)
			}
			var vRequireAttestationsBy []interface{}
			for _, vRequireAttestationsByVal := range v.RequireAttestationsBy {
				vRequireAttestationsBy = append(vRequireAttestationsBy, vRequireAttestationsByVal)
			}
			rClusterAdmissionRulesMap["requireAttestationsBy"] = vRequireAttestationsBy
			rClusterAdmissionRules[k] = rClusterAdmissionRulesMap
		}
		u.Object["clusterAdmissionRules"] = rClusterAdmissionRules
	}
	if r.DefaultAdmissionRule != nil && r.DefaultAdmissionRule != dclService.EmptyPolicyDefaultAdmissionRule {
		rDefaultAdmissionRule := make(map[string]interface{})
		if r.DefaultAdmissionRule.EnforcementMode != nil {
			rDefaultAdmissionRule["enforcementMode"] = string(*r.DefaultAdmissionRule.EnforcementMode)
		}
		if r.DefaultAdmissionRule.EvaluationMode != nil {
			rDefaultAdmissionRule["evaluationMode"] = string(*r.DefaultAdmissionRule.EvaluationMode)
		}
		var rDefaultAdmissionRuleRequireAttestationsBy []interface{}
		for _, rDefaultAdmissionRuleRequireAttestationsByVal := range r.DefaultAdmissionRule.RequireAttestationsBy {
			rDefaultAdmissionRuleRequireAttestationsBy = append(rDefaultAdmissionRuleRequireAttestationsBy, rDefaultAdmissionRuleRequireAttestationsByVal)
		}
		rDefaultAdmissionRule["requireAttestationsBy"] = rDefaultAdmissionRuleRequireAttestationsBy
		u.Object["defaultAdmissionRule"] = rDefaultAdmissionRule
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.GlobalPolicyEvaluationMode != nil {
		u.Object["globalPolicyEvaluationMode"] = string(*r.GlobalPolicyEvaluationMode)
	}
	if r.IstioServiceIdentityAdmissionRules != nil {
		rIstioServiceIdentityAdmissionRules := make(map[string]interface{})
		for k, v := range r.IstioServiceIdentityAdmissionRules {
			rIstioServiceIdentityAdmissionRulesMap := make(map[string]interface{})
			if v.EnforcementMode != nil {
				rIstioServiceIdentityAdmissionRulesMap["enforcementMode"] = string(*v.EnforcementMode)
			}
			if v.EvaluationMode != nil {
				rIstioServiceIdentityAdmissionRulesMap["evaluationMode"] = string(*v.EvaluationMode)
			}
			var vRequireAttestationsBy []interface{}
			for _, vRequireAttestationsByVal := range v.RequireAttestationsBy {
				vRequireAttestationsBy = append(vRequireAttestationsBy, vRequireAttestationsByVal)
			}
			rIstioServiceIdentityAdmissionRulesMap["requireAttestationsBy"] = vRequireAttestationsBy
			rIstioServiceIdentityAdmissionRules[k] = rIstioServiceIdentityAdmissionRulesMap
		}
		u.Object["istioServiceIdentityAdmissionRules"] = rIstioServiceIdentityAdmissionRules
	}
	if r.KubernetesNamespaceAdmissionRules != nil {
		rKubernetesNamespaceAdmissionRules := make(map[string]interface{})
		for k, v := range r.KubernetesNamespaceAdmissionRules {
			rKubernetesNamespaceAdmissionRulesMap := make(map[string]interface{})
			if v.EnforcementMode != nil {
				rKubernetesNamespaceAdmissionRulesMap["enforcementMode"] = string(*v.EnforcementMode)
			}
			if v.EvaluationMode != nil {
				rKubernetesNamespaceAdmissionRulesMap["evaluationMode"] = string(*v.EvaluationMode)
			}
			var vRequireAttestationsBy []interface{}
			for _, vRequireAttestationsByVal := range v.RequireAttestationsBy {
				vRequireAttestationsBy = append(vRequireAttestationsBy, vRequireAttestationsByVal)
			}
			rKubernetesNamespaceAdmissionRulesMap["requireAttestationsBy"] = vRequireAttestationsBy
			rKubernetesNamespaceAdmissionRules[k] = rKubernetesNamespaceAdmissionRulesMap
		}
		u.Object["kubernetesNamespaceAdmissionRules"] = rKubernetesNamespaceAdmissionRules
	}
	if r.KubernetesServiceAccountAdmissionRules != nil {
		rKubernetesServiceAccountAdmissionRules := make(map[string]interface{})
		for k, v := range r.KubernetesServiceAccountAdmissionRules {
			rKubernetesServiceAccountAdmissionRulesMap := make(map[string]interface{})
			if v.EnforcementMode != nil {
				rKubernetesServiceAccountAdmissionRulesMap["enforcementMode"] = string(*v.EnforcementMode)
			}
			if v.EvaluationMode != nil {
				rKubernetesServiceAccountAdmissionRulesMap["evaluationMode"] = string(*v.EvaluationMode)
			}
			var vRequireAttestationsBy []interface{}
			for _, vRequireAttestationsByVal := range v.RequireAttestationsBy {
				vRequireAttestationsBy = append(vRequireAttestationsBy, vRequireAttestationsByVal)
			}
			rKubernetesServiceAccountAdmissionRulesMap["requireAttestationsBy"] = vRequireAttestationsBy
			rKubernetesServiceAccountAdmissionRules[k] = rKubernetesServiceAccountAdmissionRulesMap
		}
		u.Object["kubernetesServiceAccountAdmissionRules"] = rKubernetesServiceAccountAdmissionRules
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.SelfLink != nil {
		u.Object["selfLink"] = *r.SelfLink
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToPolicy(u *unstructured.Resource) (*dclService.Policy, error) {
	r := &dclService.Policy{}
	if _, ok := u.Object["admissionWhitelistPatterns"]; ok {
		if s, ok := u.Object["admissionWhitelistPatterns"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rAdmissionWhitelistPatterns dclService.PolicyAdmissionWhitelistPatterns
					if _, ok := objval["namePattern"]; ok {
						if s, ok := objval["namePattern"].(string); ok {
							rAdmissionWhitelistPatterns.NamePattern = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rAdmissionWhitelistPatterns.NamePattern: expected string")
						}
					}
					r.AdmissionWhitelistPatterns = append(r.AdmissionWhitelistPatterns, rAdmissionWhitelistPatterns)
				}
			}
		} else {
			return nil, fmt.Errorf("r.AdmissionWhitelistPatterns: expected []interface{}")
		}
	}
	if _, ok := u.Object["clusterAdmissionRules"]; ok {
		if rClusterAdmissionRules, ok := u.Object["clusterAdmissionRules"].(map[string]interface{}); ok {
			m := make(map[string]dclService.PolicyClusterAdmissionRules)
			for k, v := range rClusterAdmissionRules {
				if objval, ok := v.(map[string]interface{}); ok {
					var rClusterAdmissionRulesObj dclService.PolicyClusterAdmissionRules
					if _, ok := objval["enforcementMode"]; ok {
						if s, ok := objval["enforcementMode"].(string); ok {
							rClusterAdmissionRulesObj.EnforcementMode = dclService.PolicyClusterAdmissionRulesEnforcementModeEnumRef(s)
						} else {
							return nil, fmt.Errorf("rClusterAdmissionRulesObj.EnforcementMode: expected string")
						}
					}
					if _, ok := objval["evaluationMode"]; ok {
						if s, ok := objval["evaluationMode"].(string); ok {
							rClusterAdmissionRulesObj.EvaluationMode = dclService.PolicyClusterAdmissionRulesEvaluationModeEnumRef(s)
						} else {
							return nil, fmt.Errorf("rClusterAdmissionRulesObj.EvaluationMode: expected string")
						}
					}
					if _, ok := objval["requireAttestationsBy"]; ok {
						if s, ok := objval["requireAttestationsBy"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									rClusterAdmissionRulesObj.RequireAttestationsBy = append(rClusterAdmissionRulesObj.RequireAttestationsBy, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("rClusterAdmissionRulesObj.RequireAttestationsBy: expected []interface{}")
						}
					}
					m[k] = rClusterAdmissionRulesObj
				} else {
					return nil, fmt.Errorf("r.ClusterAdmissionRules: expected map[string]interface{}")
				}
			}
			r.ClusterAdmissionRules = m
		} else {
			return nil, fmt.Errorf("r.ClusterAdmissionRules: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["defaultAdmissionRule"]; ok {
		if rDefaultAdmissionRule, ok := u.Object["defaultAdmissionRule"].(map[string]interface{}); ok {
			r.DefaultAdmissionRule = &dclService.PolicyDefaultAdmissionRule{}
			if _, ok := rDefaultAdmissionRule["enforcementMode"]; ok {
				if s, ok := rDefaultAdmissionRule["enforcementMode"].(string); ok {
					r.DefaultAdmissionRule.EnforcementMode = dclService.PolicyDefaultAdmissionRuleEnforcementModeEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.DefaultAdmissionRule.EnforcementMode: expected string")
				}
			}
			if _, ok := rDefaultAdmissionRule["evaluationMode"]; ok {
				if s, ok := rDefaultAdmissionRule["evaluationMode"].(string); ok {
					r.DefaultAdmissionRule.EvaluationMode = dclService.PolicyDefaultAdmissionRuleEvaluationModeEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.DefaultAdmissionRule.EvaluationMode: expected string")
				}
			}
			if _, ok := rDefaultAdmissionRule["requireAttestationsBy"]; ok {
				if s, ok := rDefaultAdmissionRule["requireAttestationsBy"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.DefaultAdmissionRule.RequireAttestationsBy = append(r.DefaultAdmissionRule.RequireAttestationsBy, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.DefaultAdmissionRule.RequireAttestationsBy: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.DefaultAdmissionRule: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["globalPolicyEvaluationMode"]; ok {
		if s, ok := u.Object["globalPolicyEvaluationMode"].(string); ok {
			r.GlobalPolicyEvaluationMode = dclService.PolicyGlobalPolicyEvaluationModeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.GlobalPolicyEvaluationMode: expected string")
		}
	}
	if _, ok := u.Object["istioServiceIdentityAdmissionRules"]; ok {
		if rIstioServiceIdentityAdmissionRules, ok := u.Object["istioServiceIdentityAdmissionRules"].(map[string]interface{}); ok {
			m := make(map[string]dclService.PolicyIstioServiceIdentityAdmissionRules)
			for k, v := range rIstioServiceIdentityAdmissionRules {
				if objval, ok := v.(map[string]interface{}); ok {
					var rIstioServiceIdentityAdmissionRulesObj dclService.PolicyIstioServiceIdentityAdmissionRules
					if _, ok := objval["enforcementMode"]; ok {
						if s, ok := objval["enforcementMode"].(string); ok {
							rIstioServiceIdentityAdmissionRulesObj.EnforcementMode = dclService.PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnumRef(s)
						} else {
							return nil, fmt.Errorf("rIstioServiceIdentityAdmissionRulesObj.EnforcementMode: expected string")
						}
					}
					if _, ok := objval["evaluationMode"]; ok {
						if s, ok := objval["evaluationMode"].(string); ok {
							rIstioServiceIdentityAdmissionRulesObj.EvaluationMode = dclService.PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnumRef(s)
						} else {
							return nil, fmt.Errorf("rIstioServiceIdentityAdmissionRulesObj.EvaluationMode: expected string")
						}
					}
					if _, ok := objval["requireAttestationsBy"]; ok {
						if s, ok := objval["requireAttestationsBy"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									rIstioServiceIdentityAdmissionRulesObj.RequireAttestationsBy = append(rIstioServiceIdentityAdmissionRulesObj.RequireAttestationsBy, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("rIstioServiceIdentityAdmissionRulesObj.RequireAttestationsBy: expected []interface{}")
						}
					}
					m[k] = rIstioServiceIdentityAdmissionRulesObj
				} else {
					return nil, fmt.Errorf("r.IstioServiceIdentityAdmissionRules: expected map[string]interface{}")
				}
			}
			r.IstioServiceIdentityAdmissionRules = m
		} else {
			return nil, fmt.Errorf("r.IstioServiceIdentityAdmissionRules: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["kubernetesNamespaceAdmissionRules"]; ok {
		if rKubernetesNamespaceAdmissionRules, ok := u.Object["kubernetesNamespaceAdmissionRules"].(map[string]interface{}); ok {
			m := make(map[string]dclService.PolicyKubernetesNamespaceAdmissionRules)
			for k, v := range rKubernetesNamespaceAdmissionRules {
				if objval, ok := v.(map[string]interface{}); ok {
					var rKubernetesNamespaceAdmissionRulesObj dclService.PolicyKubernetesNamespaceAdmissionRules
					if _, ok := objval["enforcementMode"]; ok {
						if s, ok := objval["enforcementMode"].(string); ok {
							rKubernetesNamespaceAdmissionRulesObj.EnforcementMode = dclService.PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnumRef(s)
						} else {
							return nil, fmt.Errorf("rKubernetesNamespaceAdmissionRulesObj.EnforcementMode: expected string")
						}
					}
					if _, ok := objval["evaluationMode"]; ok {
						if s, ok := objval["evaluationMode"].(string); ok {
							rKubernetesNamespaceAdmissionRulesObj.EvaluationMode = dclService.PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnumRef(s)
						} else {
							return nil, fmt.Errorf("rKubernetesNamespaceAdmissionRulesObj.EvaluationMode: expected string")
						}
					}
					if _, ok := objval["requireAttestationsBy"]; ok {
						if s, ok := objval["requireAttestationsBy"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									rKubernetesNamespaceAdmissionRulesObj.RequireAttestationsBy = append(rKubernetesNamespaceAdmissionRulesObj.RequireAttestationsBy, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("rKubernetesNamespaceAdmissionRulesObj.RequireAttestationsBy: expected []interface{}")
						}
					}
					m[k] = rKubernetesNamespaceAdmissionRulesObj
				} else {
					return nil, fmt.Errorf("r.KubernetesNamespaceAdmissionRules: expected map[string]interface{}")
				}
			}
			r.KubernetesNamespaceAdmissionRules = m
		} else {
			return nil, fmt.Errorf("r.KubernetesNamespaceAdmissionRules: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["kubernetesServiceAccountAdmissionRules"]; ok {
		if rKubernetesServiceAccountAdmissionRules, ok := u.Object["kubernetesServiceAccountAdmissionRules"].(map[string]interface{}); ok {
			m := make(map[string]dclService.PolicyKubernetesServiceAccountAdmissionRules)
			for k, v := range rKubernetesServiceAccountAdmissionRules {
				if objval, ok := v.(map[string]interface{}); ok {
					var rKubernetesServiceAccountAdmissionRulesObj dclService.PolicyKubernetesServiceAccountAdmissionRules
					if _, ok := objval["enforcementMode"]; ok {
						if s, ok := objval["enforcementMode"].(string); ok {
							rKubernetesServiceAccountAdmissionRulesObj.EnforcementMode = dclService.PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnumRef(s)
						} else {
							return nil, fmt.Errorf("rKubernetesServiceAccountAdmissionRulesObj.EnforcementMode: expected string")
						}
					}
					if _, ok := objval["evaluationMode"]; ok {
						if s, ok := objval["evaluationMode"].(string); ok {
							rKubernetesServiceAccountAdmissionRulesObj.EvaluationMode = dclService.PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnumRef(s)
						} else {
							return nil, fmt.Errorf("rKubernetesServiceAccountAdmissionRulesObj.EvaluationMode: expected string")
						}
					}
					if _, ok := objval["requireAttestationsBy"]; ok {
						if s, ok := objval["requireAttestationsBy"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									rKubernetesServiceAccountAdmissionRulesObj.RequireAttestationsBy = append(rKubernetesServiceAccountAdmissionRulesObj.RequireAttestationsBy, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("rKubernetesServiceAccountAdmissionRulesObj.RequireAttestationsBy: expected []interface{}")
						}
					}
					m[k] = rKubernetesServiceAccountAdmissionRulesObj
				} else {
					return nil, fmt.Errorf("r.KubernetesServiceAccountAdmissionRules: expected map[string]interface{}")
				}
			}
			r.KubernetesServiceAccountAdmissionRules = m
		} else {
			return nil, fmt.Errorf("r.KubernetesServiceAccountAdmissionRules: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["selfLink"]; ok {
		if s, ok := u.Object["selfLink"].(string); ok {
			r.SelfLink = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.SelfLink: expected string")
		}
	}
	if _, ok := u.Object["updateTime"]; ok {
		if s, ok := u.Object["updateTime"].(string); ok {
			r.UpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UpdateTime: expected string")
		}
	}
	return r, nil
}

func GetPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToPolicy(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetPolicy(ctx, r)
	if err != nil {
		return nil, err
	}
	return PolicyToUnstructured(r), nil
}

func ApplyPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToPolicy(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToPolicy(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyPolicy(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return PolicyToUnstructured(r), nil
}

func PolicyHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToPolicy(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToPolicy(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyPolicy(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeletePolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func PolicyID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToPolicy(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Policy) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"binaryauthorization",
		"Policy",
		"beta",
	}
}

func SetPolicyPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToPolicy(u)
	if err != nil {
		return nil, err
	}
	policy, err := iamUnstruct.UnstructuredToPolicy(p)
	if err != nil {
		return nil, err
	}
	policy.Resource = r
	iamClient := iam.NewClient(config)
	newPolicy, err := iamClient.SetPolicy(ctx, policy)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(newPolicy), nil
}

func SetPolicyWithEtagPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToPolicy(u)
	if err != nil {
		return nil, err
	}
	policy, err := iamUnstruct.UnstructuredToPolicy(p)
	if err != nil {
		return nil, err
	}
	policy.Resource = r
	iamClient := iam.NewClient(config)
	newPolicy, err := iamClient.SetPolicyWithEtag(ctx, policy)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(newPolicy), nil
}

func GetPolicyPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToPolicy(u)
	if err != nil {
		return nil, err
	}
	iamClient := iam.NewClient(config)
	policy, err := iamClient.GetPolicy(ctx, r)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(policy), nil
}

func SetPolicyMemberPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToPolicy(u)
	if err != nil {
		return nil, err
	}
	member, err := iamUnstruct.UnstructuredToMember(m)
	if err != nil {
		return nil, err
	}
	member.Resource = r
	iamClient := iam.NewClient(config)
	policy, err := iamClient.SetMember(ctx, member)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(policy), nil
}

func GetPolicyMemberPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	r, err := UnstructuredToPolicy(u)
	if err != nil {
		return nil, err
	}
	iamClient := iam.NewClient(config)
	policyMember, err := iamClient.GetMember(ctx, r, role, member)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.MemberToUnstructured(policyMember), nil
}

func DeletePolicyMemberPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) error {
	r, err := UnstructuredToPolicy(u)
	if err != nil {
		return err
	}
	member, err := iamUnstruct.UnstructuredToMember(m)
	if err != nil {
		return err
	}
	member.Resource = r
	iamClient := iam.NewClient(config)
	if err := iamClient.DeleteMember(ctx, member); err != nil {
		return err
	}
	return nil
}

func (r *Policy) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyMemberPolicy(ctx, config, resource, member)
}

func (r *Policy) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return GetPolicyMemberPolicy(ctx, config, resource, role, member)
}

func (r *Policy) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return DeletePolicyMemberPolicy(ctx, config, resource, member)
}

func (r *Policy) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyPolicy(ctx, config, resource, policy)
}

func (r *Policy) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyWithEtagPolicy(ctx, config, resource, policy)
}

func (r *Policy) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetPolicyPolicy(ctx, config, resource)
}

func (r *Policy) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetPolicy(ctx, config, resource)
}

func (r *Policy) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyPolicy(ctx, config, resource, opts...)
}

func (r *Policy) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return PolicyHasDiff(ctx, config, resource, opts...)
}

func (r *Policy) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeletePolicy(ctx, config, resource)
}

func (r *Policy) ID(resource *unstructured.Resource) (string, error) {
	return PolicyID(resource)
}

func init() {
	unstructured.Register(&Policy{})
}
