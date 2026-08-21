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
package networksecurity

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networksecurity/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	iamUnstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/iam"
)

type AuthorizationPolicy struct{}

func AuthorizationPolicyToUnstructured(r *dclService.AuthorizationPolicy) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "networksecurity",
			Version: "beta",
			Type:    "AuthorizationPolicy",
		},
		Object: make(map[string]interface{}),
	}
	if r.Action != nil {
		u.Object["action"] = string(*r.Action)
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.Labels != nil {
		rLabels := make(map[string]interface{})
		for k, v := range r.Labels {
			rLabels[k] = v
		}
		u.Object["labels"] = rLabels
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	var rRules []interface{}
	for _, rRulesVal := range r.Rules {
		rRulesObject := make(map[string]interface{})
		var rRulesValDestinations []interface{}
		for _, rRulesValDestinationsVal := range rRulesVal.Destinations {
			rRulesValDestinationsObject := make(map[string]interface{})
			var rRulesValDestinationsValHosts []interface{}
			for _, rRulesValDestinationsValHostsVal := range rRulesValDestinationsVal.Hosts {
				rRulesValDestinationsValHosts = append(rRulesValDestinationsValHosts, rRulesValDestinationsValHostsVal)
			}
			rRulesValDestinationsObject["hosts"] = rRulesValDestinationsValHosts
			if rRulesValDestinationsVal.HttpHeaderMatch != nil && rRulesValDestinationsVal.HttpHeaderMatch != dclService.EmptyAuthorizationPolicyRulesDestinationsHttpHeaderMatch {
				rRulesValDestinationsValHttpHeaderMatch := make(map[string]interface{})
				if rRulesValDestinationsVal.HttpHeaderMatch.HeaderName != nil {
					rRulesValDestinationsValHttpHeaderMatch["headerName"] = *rRulesValDestinationsVal.HttpHeaderMatch.HeaderName
				}
				if rRulesValDestinationsVal.HttpHeaderMatch.RegexMatch != nil {
					rRulesValDestinationsValHttpHeaderMatch["regexMatch"] = *rRulesValDestinationsVal.HttpHeaderMatch.RegexMatch
				}
				rRulesValDestinationsObject["httpHeaderMatch"] = rRulesValDestinationsValHttpHeaderMatch
			}
			var rRulesValDestinationsValMethods []interface{}
			for _, rRulesValDestinationsValMethodsVal := range rRulesValDestinationsVal.Methods {
				rRulesValDestinationsValMethods = append(rRulesValDestinationsValMethods, rRulesValDestinationsValMethodsVal)
			}
			rRulesValDestinationsObject["methods"] = rRulesValDestinationsValMethods
			var rRulesValDestinationsValPorts []interface{}
			for _, rRulesValDestinationsValPortsVal := range rRulesValDestinationsVal.Ports {
				rRulesValDestinationsValPorts = append(rRulesValDestinationsValPorts, rRulesValDestinationsValPortsVal)
			}
			rRulesValDestinationsObject["ports"] = rRulesValDestinationsValPorts
			rRulesValDestinations = append(rRulesValDestinations, rRulesValDestinationsObject)
		}
		rRulesObject["destinations"] = rRulesValDestinations
		var rRulesValSources []interface{}
		for _, rRulesValSourcesVal := range rRulesVal.Sources {
			rRulesValSourcesObject := make(map[string]interface{})
			var rRulesValSourcesValIPBlocks []interface{}
			for _, rRulesValSourcesValIPBlocksVal := range rRulesValSourcesVal.IPBlocks {
				rRulesValSourcesValIPBlocks = append(rRulesValSourcesValIPBlocks, rRulesValSourcesValIPBlocksVal)
			}
			rRulesValSourcesObject["ipBlocks"] = rRulesValSourcesValIPBlocks
			var rRulesValSourcesValPrincipals []interface{}
			for _, rRulesValSourcesValPrincipalsVal := range rRulesValSourcesVal.Principals {
				rRulesValSourcesValPrincipals = append(rRulesValSourcesValPrincipals, rRulesValSourcesValPrincipalsVal)
			}
			rRulesValSourcesObject["principals"] = rRulesValSourcesValPrincipals
			rRulesValSources = append(rRulesValSources, rRulesValSourcesObject)
		}
		rRulesObject["sources"] = rRulesValSources
		rRules = append(rRules, rRulesObject)
	}
	u.Object["rules"] = rRules
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToAuthorizationPolicy(u *unstructured.Resource) (*dclService.AuthorizationPolicy, error) {
	r := &dclService.AuthorizationPolicy{}
	if _, ok := u.Object["action"]; ok {
		if s, ok := u.Object["action"].(string); ok {
			r.Action = dclService.AuthorizationPolicyActionEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Action: expected string")
		}
	}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["labels"]; ok {
		if rLabels, ok := u.Object["labels"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rLabels {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.Labels = m
		} else {
			return nil, fmt.Errorf("r.Labels: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["rules"]; ok {
		if s, ok := u.Object["rules"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rRules dclService.AuthorizationPolicyRules
					if _, ok := objval["destinations"]; ok {
						if s, ok := objval["destinations"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rRulesDestinations dclService.AuthorizationPolicyRulesDestinations
									if _, ok := objval["hosts"]; ok {
										if s, ok := objval["hosts"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rRulesDestinations.Hosts = append(rRulesDestinations.Hosts, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rRulesDestinations.Hosts: expected []interface{}")
										}
									}
									if _, ok := objval["httpHeaderMatch"]; ok {
										if rRulesDestinationsHttpHeaderMatch, ok := objval["httpHeaderMatch"].(map[string]interface{}); ok {
											rRulesDestinations.HttpHeaderMatch = &dclService.AuthorizationPolicyRulesDestinationsHttpHeaderMatch{}
											if _, ok := rRulesDestinationsHttpHeaderMatch["headerName"]; ok {
												if s, ok := rRulesDestinationsHttpHeaderMatch["headerName"].(string); ok {
													rRulesDestinations.HttpHeaderMatch.HeaderName = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRulesDestinations.HttpHeaderMatch.HeaderName: expected string")
												}
											}
											if _, ok := rRulesDestinationsHttpHeaderMatch["regexMatch"]; ok {
												if s, ok := rRulesDestinationsHttpHeaderMatch["regexMatch"].(string); ok {
													rRulesDestinations.HttpHeaderMatch.RegexMatch = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRulesDestinations.HttpHeaderMatch.RegexMatch: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rRulesDestinations.HttpHeaderMatch: expected map[string]interface{}")
										}
									}
									if _, ok := objval["methods"]; ok {
										if s, ok := objval["methods"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rRulesDestinations.Methods = append(rRulesDestinations.Methods, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rRulesDestinations.Methods: expected []interface{}")
										}
									}
									if _, ok := objval["ports"]; ok {
										if s, ok := objval["ports"].([]interface{}); ok {
											for _, ss := range s {
												if intval, ok := ss.(int64); ok {
													rRulesDestinations.Ports = append(rRulesDestinations.Ports, intval)
												}
											}
										} else {
											return nil, fmt.Errorf("rRulesDestinations.Ports: expected []interface{}")
										}
									}
									rRules.Destinations = append(rRules.Destinations, rRulesDestinations)
								}
							}
						} else {
							return nil, fmt.Errorf("rRules.Destinations: expected []interface{}")
						}
					}
					if _, ok := objval["sources"]; ok {
						if s, ok := objval["sources"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rRulesSources dclService.AuthorizationPolicyRulesSources
									if _, ok := objval["ipBlocks"]; ok {
										if s, ok := objval["ipBlocks"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rRulesSources.IPBlocks = append(rRulesSources.IPBlocks, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rRulesSources.IPBlocks: expected []interface{}")
										}
									}
									if _, ok := objval["principals"]; ok {
										if s, ok := objval["principals"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rRulesSources.Principals = append(rRulesSources.Principals, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rRulesSources.Principals: expected []interface{}")
										}
									}
									rRules.Sources = append(rRules.Sources, rRulesSources)
								}
							}
						} else {
							return nil, fmt.Errorf("rRules.Sources: expected []interface{}")
						}
					}
					r.Rules = append(r.Rules, rRules)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Rules: expected []interface{}")
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

func GetAuthorizationPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAuthorizationPolicy(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetAuthorizationPolicy(ctx, r)
	if err != nil {
		return nil, err
	}
	return AuthorizationPolicyToUnstructured(r), nil
}

func ListAuthorizationPolicy(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListAuthorizationPolicy(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, AuthorizationPolicyToUnstructured(r))
		}
		if !l.HasNext() {
			break
		}
		if err := l.Next(ctx, c); err != nil {
			return nil, err
		}
	}
	return resources, nil
}

func ApplyAuthorizationPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAuthorizationPolicy(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToAuthorizationPolicy(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyAuthorizationPolicy(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return AuthorizationPolicyToUnstructured(r), nil
}

func AuthorizationPolicyHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAuthorizationPolicy(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToAuthorizationPolicy(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyAuthorizationPolicy(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteAuthorizationPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAuthorizationPolicy(u)
	if err != nil {
		return err
	}
	return c.DeleteAuthorizationPolicy(ctx, r)
}

func AuthorizationPolicyID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToAuthorizationPolicy(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *AuthorizationPolicy) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"networksecurity",
		"AuthorizationPolicy",
		"beta",
	}
}

func SetPolicyAuthorizationPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToAuthorizationPolicy(u)
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

func SetPolicyWithEtagAuthorizationPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToAuthorizationPolicy(u)
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

func GetPolicyAuthorizationPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToAuthorizationPolicy(u)
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

func SetPolicyMemberAuthorizationPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToAuthorizationPolicy(u)
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

func GetPolicyMemberAuthorizationPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	r, err := UnstructuredToAuthorizationPolicy(u)
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

func DeletePolicyMemberAuthorizationPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) error {
	r, err := UnstructuredToAuthorizationPolicy(u)
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

func (r *AuthorizationPolicy) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyMemberAuthorizationPolicy(ctx, config, resource, member)
}

func (r *AuthorizationPolicy) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return GetPolicyMemberAuthorizationPolicy(ctx, config, resource, role, member)
}

func (r *AuthorizationPolicy) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return DeletePolicyMemberAuthorizationPolicy(ctx, config, resource, member)
}

func (r *AuthorizationPolicy) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyAuthorizationPolicy(ctx, config, resource, policy)
}

func (r *AuthorizationPolicy) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyWithEtagAuthorizationPolicy(ctx, config, resource, policy)
}

func (r *AuthorizationPolicy) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetPolicyAuthorizationPolicy(ctx, config, resource)
}

func (r *AuthorizationPolicy) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetAuthorizationPolicy(ctx, config, resource)
}

func (r *AuthorizationPolicy) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyAuthorizationPolicy(ctx, config, resource, opts...)
}

func (r *AuthorizationPolicy) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return AuthorizationPolicyHasDiff(ctx, config, resource, opts...)
}

func (r *AuthorizationPolicy) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteAuthorizationPolicy(ctx, config, resource)
}

func (r *AuthorizationPolicy) ID(resource *unstructured.Resource) (string, error) {
	return AuthorizationPolicyID(resource)
}

func init() {
	unstructured.Register(&AuthorizationPolicy{})
}
