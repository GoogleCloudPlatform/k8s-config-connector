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
package compute

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type NetworkFirewallPolicyRule struct{}

func NetworkFirewallPolicyRuleToUnstructured(r *dclService.NetworkFirewallPolicyRule) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "compute",
			Version: "alpha",
			Type:    "NetworkFirewallPolicyRule",
		},
		Object: make(map[string]interface{}),
	}
	if r.Action != nil {
		u.Object["action"] = *r.Action
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.Direction != nil {
		u.Object["direction"] = string(*r.Direction)
	}
	if r.Disabled != nil {
		u.Object["disabled"] = *r.Disabled
	}
	if r.EnableLogging != nil {
		u.Object["enableLogging"] = *r.EnableLogging
	}
	if r.FirewallPolicy != nil {
		u.Object["firewallPolicy"] = *r.FirewallPolicy
	}
	if r.Kind != nil {
		u.Object["kind"] = *r.Kind
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Match != nil && r.Match != dclService.EmptyNetworkFirewallPolicyRuleMatch {
		rMatch := make(map[string]interface{})
		var rMatchDestAddressGroups []interface{}
		for _, rMatchDestAddressGroupsVal := range r.Match.DestAddressGroups {
			rMatchDestAddressGroups = append(rMatchDestAddressGroups, rMatchDestAddressGroupsVal)
		}
		rMatch["destAddressGroups"] = rMatchDestAddressGroups
		var rMatchDestFqdns []interface{}
		for _, rMatchDestFqdnsVal := range r.Match.DestFqdns {
			rMatchDestFqdns = append(rMatchDestFqdns, rMatchDestFqdnsVal)
		}
		rMatch["destFqdns"] = rMatchDestFqdns
		var rMatchDestIPRanges []interface{}
		for _, rMatchDestIPRangesVal := range r.Match.DestIPRanges {
			rMatchDestIPRanges = append(rMatchDestIPRanges, rMatchDestIPRangesVal)
		}
		rMatch["destIPRanges"] = rMatchDestIPRanges
		var rMatchDestRegionCodes []interface{}
		for _, rMatchDestRegionCodesVal := range r.Match.DestRegionCodes {
			rMatchDestRegionCodes = append(rMatchDestRegionCodes, rMatchDestRegionCodesVal)
		}
		rMatch["destRegionCodes"] = rMatchDestRegionCodes
		var rMatchDestThreatIntelligences []interface{}
		for _, rMatchDestThreatIntelligencesVal := range r.Match.DestThreatIntelligences {
			rMatchDestThreatIntelligences = append(rMatchDestThreatIntelligences, rMatchDestThreatIntelligencesVal)
		}
		rMatch["destThreatIntelligences"] = rMatchDestThreatIntelligences
		var rMatchLayer4Configs []interface{}
		for _, rMatchLayer4ConfigsVal := range r.Match.Layer4Configs {
			rMatchLayer4ConfigsObject := make(map[string]interface{})
			if rMatchLayer4ConfigsVal.IPProtocol != nil {
				rMatchLayer4ConfigsObject["ipProtocol"] = *rMatchLayer4ConfigsVal.IPProtocol
			}
			var rMatchLayer4ConfigsValPorts []interface{}
			for _, rMatchLayer4ConfigsValPortsVal := range rMatchLayer4ConfigsVal.Ports {
				rMatchLayer4ConfigsValPorts = append(rMatchLayer4ConfigsValPorts, rMatchLayer4ConfigsValPortsVal)
			}
			rMatchLayer4ConfigsObject["ports"] = rMatchLayer4ConfigsValPorts
			rMatchLayer4Configs = append(rMatchLayer4Configs, rMatchLayer4ConfigsObject)
		}
		rMatch["layer4Configs"] = rMatchLayer4Configs
		var rMatchSrcAddressGroups []interface{}
		for _, rMatchSrcAddressGroupsVal := range r.Match.SrcAddressGroups {
			rMatchSrcAddressGroups = append(rMatchSrcAddressGroups, rMatchSrcAddressGroupsVal)
		}
		rMatch["srcAddressGroups"] = rMatchSrcAddressGroups
		var rMatchSrcFqdns []interface{}
		for _, rMatchSrcFqdnsVal := range r.Match.SrcFqdns {
			rMatchSrcFqdns = append(rMatchSrcFqdns, rMatchSrcFqdnsVal)
		}
		rMatch["srcFqdns"] = rMatchSrcFqdns
		var rMatchSrcIPRanges []interface{}
		for _, rMatchSrcIPRangesVal := range r.Match.SrcIPRanges {
			rMatchSrcIPRanges = append(rMatchSrcIPRanges, rMatchSrcIPRangesVal)
		}
		rMatch["srcIPRanges"] = rMatchSrcIPRanges
		var rMatchSrcRegionCodes []interface{}
		for _, rMatchSrcRegionCodesVal := range r.Match.SrcRegionCodes {
			rMatchSrcRegionCodes = append(rMatchSrcRegionCodes, rMatchSrcRegionCodesVal)
		}
		rMatch["srcRegionCodes"] = rMatchSrcRegionCodes
		var rMatchSrcSecureTags []interface{}
		for _, rMatchSrcSecureTagsVal := range r.Match.SrcSecureTags {
			rMatchSrcSecureTagsObject := make(map[string]interface{})
			if rMatchSrcSecureTagsVal.Name != nil {
				rMatchSrcSecureTagsObject["name"] = *rMatchSrcSecureTagsVal.Name
			}
			if rMatchSrcSecureTagsVal.State != nil {
				rMatchSrcSecureTagsObject["state"] = string(*rMatchSrcSecureTagsVal.State)
			}
			rMatchSrcSecureTags = append(rMatchSrcSecureTags, rMatchSrcSecureTagsObject)
		}
		rMatch["srcSecureTags"] = rMatchSrcSecureTags
		var rMatchSrcThreatIntelligences []interface{}
		for _, rMatchSrcThreatIntelligencesVal := range r.Match.SrcThreatIntelligences {
			rMatchSrcThreatIntelligences = append(rMatchSrcThreatIntelligences, rMatchSrcThreatIntelligencesVal)
		}
		rMatch["srcThreatIntelligences"] = rMatchSrcThreatIntelligences
		u.Object["match"] = rMatch
	}
	if r.Priority != nil {
		u.Object["priority"] = *r.Priority
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.RuleName != nil {
		u.Object["ruleName"] = *r.RuleName
	}
	if r.RuleTupleCount != nil {
		u.Object["ruleTupleCount"] = *r.RuleTupleCount
	}
	var rTargetSecureTags []interface{}
	for _, rTargetSecureTagsVal := range r.TargetSecureTags {
		rTargetSecureTagsObject := make(map[string]interface{})
		if rTargetSecureTagsVal.Name != nil {
			rTargetSecureTagsObject["name"] = *rTargetSecureTagsVal.Name
		}
		if rTargetSecureTagsVal.State != nil {
			rTargetSecureTagsObject["state"] = string(*rTargetSecureTagsVal.State)
		}
		rTargetSecureTags = append(rTargetSecureTags, rTargetSecureTagsObject)
	}
	u.Object["targetSecureTags"] = rTargetSecureTags
	var rTargetServiceAccounts []interface{}
	for _, rTargetServiceAccountsVal := range r.TargetServiceAccounts {
		rTargetServiceAccounts = append(rTargetServiceAccounts, rTargetServiceAccountsVal)
	}
	u.Object["targetServiceAccounts"] = rTargetServiceAccounts
	return u
}

func UnstructuredToNetworkFirewallPolicyRule(u *unstructured.Resource) (*dclService.NetworkFirewallPolicyRule, error) {
	r := &dclService.NetworkFirewallPolicyRule{}
	if _, ok := u.Object["action"]; ok {
		if s, ok := u.Object["action"].(string); ok {
			r.Action = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Action: expected string")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["direction"]; ok {
		if s, ok := u.Object["direction"].(string); ok {
			r.Direction = dclService.NetworkFirewallPolicyRuleDirectionEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Direction: expected string")
		}
	}
	if _, ok := u.Object["disabled"]; ok {
		if b, ok := u.Object["disabled"].(bool); ok {
			r.Disabled = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.Disabled: expected bool")
		}
	}
	if _, ok := u.Object["enableLogging"]; ok {
		if b, ok := u.Object["enableLogging"].(bool); ok {
			r.EnableLogging = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.EnableLogging: expected bool")
		}
	}
	if _, ok := u.Object["firewallPolicy"]; ok {
		if s, ok := u.Object["firewallPolicy"].(string); ok {
			r.FirewallPolicy = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.FirewallPolicy: expected string")
		}
	}
	if _, ok := u.Object["kind"]; ok {
		if s, ok := u.Object["kind"].(string); ok {
			r.Kind = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Kind: expected string")
		}
	}
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["match"]; ok {
		if rMatch, ok := u.Object["match"].(map[string]interface{}); ok {
			r.Match = &dclService.NetworkFirewallPolicyRuleMatch{}
			if _, ok := rMatch["destAddressGroups"]; ok {
				if s, ok := rMatch["destAddressGroups"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Match.DestAddressGroups = append(r.Match.DestAddressGroups, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Match.DestAddressGroups: expected []interface{}")
				}
			}
			if _, ok := rMatch["destFqdns"]; ok {
				if s, ok := rMatch["destFqdns"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Match.DestFqdns = append(r.Match.DestFqdns, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Match.DestFqdns: expected []interface{}")
				}
			}
			if _, ok := rMatch["destIPRanges"]; ok {
				if s, ok := rMatch["destIPRanges"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Match.DestIPRanges = append(r.Match.DestIPRanges, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Match.DestIPRanges: expected []interface{}")
				}
			}
			if _, ok := rMatch["destRegionCodes"]; ok {
				if s, ok := rMatch["destRegionCodes"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Match.DestRegionCodes = append(r.Match.DestRegionCodes, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Match.DestRegionCodes: expected []interface{}")
				}
			}
			if _, ok := rMatch["destThreatIntelligences"]; ok {
				if s, ok := rMatch["destThreatIntelligences"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Match.DestThreatIntelligences = append(r.Match.DestThreatIntelligences, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Match.DestThreatIntelligences: expected []interface{}")
				}
			}
			if _, ok := rMatch["layer4Configs"]; ok {
				if s, ok := rMatch["layer4Configs"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rMatchLayer4Configs dclService.NetworkFirewallPolicyRuleMatchLayer4Configs
							if _, ok := objval["ipProtocol"]; ok {
								if s, ok := objval["ipProtocol"].(string); ok {
									rMatchLayer4Configs.IPProtocol = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rMatchLayer4Configs.IPProtocol: expected string")
								}
							}
							if _, ok := objval["ports"]; ok {
								if s, ok := objval["ports"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rMatchLayer4Configs.Ports = append(rMatchLayer4Configs.Ports, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rMatchLayer4Configs.Ports: expected []interface{}")
								}
							}
							r.Match.Layer4Configs = append(r.Match.Layer4Configs, rMatchLayer4Configs)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Match.Layer4Configs: expected []interface{}")
				}
			}
			if _, ok := rMatch["srcAddressGroups"]; ok {
				if s, ok := rMatch["srcAddressGroups"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Match.SrcAddressGroups = append(r.Match.SrcAddressGroups, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Match.SrcAddressGroups: expected []interface{}")
				}
			}
			if _, ok := rMatch["srcFqdns"]; ok {
				if s, ok := rMatch["srcFqdns"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Match.SrcFqdns = append(r.Match.SrcFqdns, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Match.SrcFqdns: expected []interface{}")
				}
			}
			if _, ok := rMatch["srcIPRanges"]; ok {
				if s, ok := rMatch["srcIPRanges"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Match.SrcIPRanges = append(r.Match.SrcIPRanges, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Match.SrcIPRanges: expected []interface{}")
				}
			}
			if _, ok := rMatch["srcRegionCodes"]; ok {
				if s, ok := rMatch["srcRegionCodes"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Match.SrcRegionCodes = append(r.Match.SrcRegionCodes, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Match.SrcRegionCodes: expected []interface{}")
				}
			}
			if _, ok := rMatch["srcSecureTags"]; ok {
				if s, ok := rMatch["srcSecureTags"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rMatchSrcSecureTags dclService.NetworkFirewallPolicyRuleMatchSrcSecureTags
							if _, ok := objval["name"]; ok {
								if s, ok := objval["name"].(string); ok {
									rMatchSrcSecureTags.Name = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rMatchSrcSecureTags.Name: expected string")
								}
							}
							if _, ok := objval["state"]; ok {
								if s, ok := objval["state"].(string); ok {
									rMatchSrcSecureTags.State = dclService.NetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnumRef(s)
								} else {
									return nil, fmt.Errorf("rMatchSrcSecureTags.State: expected string")
								}
							}
							r.Match.SrcSecureTags = append(r.Match.SrcSecureTags, rMatchSrcSecureTags)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Match.SrcSecureTags: expected []interface{}")
				}
			}
			if _, ok := rMatch["srcThreatIntelligences"]; ok {
				if s, ok := rMatch["srcThreatIntelligences"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Match.SrcThreatIntelligences = append(r.Match.SrcThreatIntelligences, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Match.SrcThreatIntelligences: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Match: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["priority"]; ok {
		if i, ok := u.Object["priority"].(int64); ok {
			r.Priority = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.Priority: expected int64")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["ruleName"]; ok {
		if s, ok := u.Object["ruleName"].(string); ok {
			r.RuleName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.RuleName: expected string")
		}
	}
	if _, ok := u.Object["ruleTupleCount"]; ok {
		if i, ok := u.Object["ruleTupleCount"].(int64); ok {
			r.RuleTupleCount = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.RuleTupleCount: expected int64")
		}
	}
	if _, ok := u.Object["targetSecureTags"]; ok {
		if s, ok := u.Object["targetSecureTags"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rTargetSecureTags dclService.NetworkFirewallPolicyRuleTargetSecureTags
					if _, ok := objval["name"]; ok {
						if s, ok := objval["name"].(string); ok {
							rTargetSecureTags.Name = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rTargetSecureTags.Name: expected string")
						}
					}
					if _, ok := objval["state"]; ok {
						if s, ok := objval["state"].(string); ok {
							rTargetSecureTags.State = dclService.NetworkFirewallPolicyRuleTargetSecureTagsStateEnumRef(s)
						} else {
							return nil, fmt.Errorf("rTargetSecureTags.State: expected string")
						}
					}
					r.TargetSecureTags = append(r.TargetSecureTags, rTargetSecureTags)
				}
			}
		} else {
			return nil, fmt.Errorf("r.TargetSecureTags: expected []interface{}")
		}
	}
	if _, ok := u.Object["targetServiceAccounts"]; ok {
		if s, ok := u.Object["targetServiceAccounts"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.TargetServiceAccounts = append(r.TargetServiceAccounts, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.TargetServiceAccounts: expected []interface{}")
		}
	}
	return r, nil
}

func GetNetworkFirewallPolicyRule(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNetworkFirewallPolicyRule(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetNetworkFirewallPolicyRule(ctx, r)
	if err != nil {
		return nil, err
	}
	return NetworkFirewallPolicyRuleToUnstructured(r), nil
}

func ListNetworkFirewallPolicyRule(ctx context.Context, config *dcl.Config, project string, location string, firewallPolicy string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListNetworkFirewallPolicyRule(ctx, project, location, firewallPolicy)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, NetworkFirewallPolicyRuleToUnstructured(r))
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

func ApplyNetworkFirewallPolicyRule(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNetworkFirewallPolicyRule(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToNetworkFirewallPolicyRule(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyNetworkFirewallPolicyRule(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return NetworkFirewallPolicyRuleToUnstructured(r), nil
}

func NetworkFirewallPolicyRuleHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNetworkFirewallPolicyRule(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToNetworkFirewallPolicyRule(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyNetworkFirewallPolicyRule(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteNetworkFirewallPolicyRule(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNetworkFirewallPolicyRule(u)
	if err != nil {
		return err
	}
	return c.DeleteNetworkFirewallPolicyRule(ctx, r)
}

func NetworkFirewallPolicyRuleID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToNetworkFirewallPolicyRule(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *NetworkFirewallPolicyRule) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"compute",
		"NetworkFirewallPolicyRule",
		"alpha",
	}
}

func (r *NetworkFirewallPolicyRule) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *NetworkFirewallPolicyRule) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *NetworkFirewallPolicyRule) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *NetworkFirewallPolicyRule) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *NetworkFirewallPolicyRule) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *NetworkFirewallPolicyRule) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *NetworkFirewallPolicyRule) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetNetworkFirewallPolicyRule(ctx, config, resource)
}

func (r *NetworkFirewallPolicyRule) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyNetworkFirewallPolicyRule(ctx, config, resource, opts...)
}

func (r *NetworkFirewallPolicyRule) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return NetworkFirewallPolicyRuleHasDiff(ctx, config, resource, opts...)
}

func (r *NetworkFirewallPolicyRule) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteNetworkFirewallPolicyRule(ctx, config, resource)
}

func (r *NetworkFirewallPolicyRule) ID(resource *unstructured.Resource) (string, error) {
	return NetworkFirewallPolicyRuleID(resource)
}

func init() {
	unstructured.Register(&NetworkFirewallPolicyRule{})
}
