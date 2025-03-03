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
package networkservices

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type EndpointPolicy struct{}

func EndpointPolicyToUnstructured(r *dclService.EndpointPolicy) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "networkservices",
			Version: "alpha",
			Type:    "EndpointPolicy",
		},
		Object: make(map[string]interface{}),
	}
	if r.AuthorizationPolicy != nil {
		u.Object["authorizationPolicy"] = *r.AuthorizationPolicy
	}
	if r.ClientTlsPolicy != nil {
		u.Object["clientTlsPolicy"] = *r.ClientTlsPolicy
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.EndpointMatcher != nil && r.EndpointMatcher != dclService.EmptyEndpointPolicyEndpointMatcher {
		rEndpointMatcher := make(map[string]interface{})
		if r.EndpointMatcher.MetadataLabelMatcher != nil && r.EndpointMatcher.MetadataLabelMatcher != dclService.EmptyEndpointPolicyEndpointMatcherMetadataLabelMatcher {
			rEndpointMatcherMetadataLabelMatcher := make(map[string]interface{})
			if r.EndpointMatcher.MetadataLabelMatcher.MetadataLabelMatchCriteria != nil {
				rEndpointMatcherMetadataLabelMatcher["metadataLabelMatchCriteria"] = string(*r.EndpointMatcher.MetadataLabelMatcher.MetadataLabelMatchCriteria)
			}
			var rEndpointMatcherMetadataLabelMatcherMetadataLabels []interface{}
			for _, rEndpointMatcherMetadataLabelMatcherMetadataLabelsVal := range r.EndpointMatcher.MetadataLabelMatcher.MetadataLabels {
				rEndpointMatcherMetadataLabelMatcherMetadataLabelsObject := make(map[string]interface{})
				if rEndpointMatcherMetadataLabelMatcherMetadataLabelsVal.LabelName != nil {
					rEndpointMatcherMetadataLabelMatcherMetadataLabelsObject["labelName"] = *rEndpointMatcherMetadataLabelMatcherMetadataLabelsVal.LabelName
				}
				if rEndpointMatcherMetadataLabelMatcherMetadataLabelsVal.LabelValue != nil {
					rEndpointMatcherMetadataLabelMatcherMetadataLabelsObject["labelValue"] = *rEndpointMatcherMetadataLabelMatcherMetadataLabelsVal.LabelValue
				}
				rEndpointMatcherMetadataLabelMatcherMetadataLabels = append(rEndpointMatcherMetadataLabelMatcherMetadataLabels, rEndpointMatcherMetadataLabelMatcherMetadataLabelsObject)
			}
			rEndpointMatcherMetadataLabelMatcher["metadataLabels"] = rEndpointMatcherMetadataLabelMatcherMetadataLabels
			rEndpointMatcher["metadataLabelMatcher"] = rEndpointMatcherMetadataLabelMatcher
		}
		u.Object["endpointMatcher"] = rEndpointMatcher
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
	if r.ServerTlsPolicy != nil {
		u.Object["serverTlsPolicy"] = *r.ServerTlsPolicy
	}
	if r.TrafficPortSelector != nil && r.TrafficPortSelector != dclService.EmptyEndpointPolicyTrafficPortSelector {
		rTrafficPortSelector := make(map[string]interface{})
		var rTrafficPortSelectorPorts []interface{}
		for _, rTrafficPortSelectorPortsVal := range r.TrafficPortSelector.Ports {
			rTrafficPortSelectorPorts = append(rTrafficPortSelectorPorts, rTrafficPortSelectorPortsVal)
		}
		rTrafficPortSelector["ports"] = rTrafficPortSelectorPorts
		u.Object["trafficPortSelector"] = rTrafficPortSelector
	}
	if r.Type != nil {
		u.Object["type"] = string(*r.Type)
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToEndpointPolicy(u *unstructured.Resource) (*dclService.EndpointPolicy, error) {
	r := &dclService.EndpointPolicy{}
	if _, ok := u.Object["authorizationPolicy"]; ok {
		if s, ok := u.Object["authorizationPolicy"].(string); ok {
			r.AuthorizationPolicy = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.AuthorizationPolicy: expected string")
		}
	}
	if _, ok := u.Object["clientTlsPolicy"]; ok {
		if s, ok := u.Object["clientTlsPolicy"].(string); ok {
			r.ClientTlsPolicy = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ClientTlsPolicy: expected string")
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
	if _, ok := u.Object["endpointMatcher"]; ok {
		if rEndpointMatcher, ok := u.Object["endpointMatcher"].(map[string]interface{}); ok {
			r.EndpointMatcher = &dclService.EndpointPolicyEndpointMatcher{}
			if _, ok := rEndpointMatcher["metadataLabelMatcher"]; ok {
				if rEndpointMatcherMetadataLabelMatcher, ok := rEndpointMatcher["metadataLabelMatcher"].(map[string]interface{}); ok {
					r.EndpointMatcher.MetadataLabelMatcher = &dclService.EndpointPolicyEndpointMatcherMetadataLabelMatcher{}
					if _, ok := rEndpointMatcherMetadataLabelMatcher["metadataLabelMatchCriteria"]; ok {
						if s, ok := rEndpointMatcherMetadataLabelMatcher["metadataLabelMatchCriteria"].(string); ok {
							r.EndpointMatcher.MetadataLabelMatcher.MetadataLabelMatchCriteria = dclService.EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.EndpointMatcher.MetadataLabelMatcher.MetadataLabelMatchCriteria: expected string")
						}
					}
					if _, ok := rEndpointMatcherMetadataLabelMatcher["metadataLabels"]; ok {
						if s, ok := rEndpointMatcherMetadataLabelMatcher["metadataLabels"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rEndpointMatcherMetadataLabelMatcherMetadataLabels dclService.EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels
									if _, ok := objval["labelName"]; ok {
										if s, ok := objval["labelName"].(string); ok {
											rEndpointMatcherMetadataLabelMatcherMetadataLabels.LabelName = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rEndpointMatcherMetadataLabelMatcherMetadataLabels.LabelName: expected string")
										}
									}
									if _, ok := objval["labelValue"]; ok {
										if s, ok := objval["labelValue"].(string); ok {
											rEndpointMatcherMetadataLabelMatcherMetadataLabels.LabelValue = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rEndpointMatcherMetadataLabelMatcherMetadataLabels.LabelValue: expected string")
										}
									}
									r.EndpointMatcher.MetadataLabelMatcher.MetadataLabels = append(r.EndpointMatcher.MetadataLabelMatcher.MetadataLabels, rEndpointMatcherMetadataLabelMatcherMetadataLabels)
								}
							}
						} else {
							return nil, fmt.Errorf("r.EndpointMatcher.MetadataLabelMatcher.MetadataLabels: expected []interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.EndpointMatcher.MetadataLabelMatcher: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.EndpointMatcher: expected map[string]interface{}")
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
	if _, ok := u.Object["serverTlsPolicy"]; ok {
		if s, ok := u.Object["serverTlsPolicy"].(string); ok {
			r.ServerTlsPolicy = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ServerTlsPolicy: expected string")
		}
	}
	if _, ok := u.Object["trafficPortSelector"]; ok {
		if rTrafficPortSelector, ok := u.Object["trafficPortSelector"].(map[string]interface{}); ok {
			r.TrafficPortSelector = &dclService.EndpointPolicyTrafficPortSelector{}
			if _, ok := rTrafficPortSelector["ports"]; ok {
				if s, ok := rTrafficPortSelector["ports"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.TrafficPortSelector.Ports = append(r.TrafficPortSelector.Ports, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.TrafficPortSelector.Ports: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.TrafficPortSelector: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["type"]; ok {
		if s, ok := u.Object["type"].(string); ok {
			r.Type = dclService.EndpointPolicyTypeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Type: expected string")
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

func GetEndpointPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEndpointPolicy(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetEndpointPolicy(ctx, r)
	if err != nil {
		return nil, err
	}
	return EndpointPolicyToUnstructured(r), nil
}

func ListEndpointPolicy(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListEndpointPolicy(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, EndpointPolicyToUnstructured(r))
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

func ApplyEndpointPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEndpointPolicy(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToEndpointPolicy(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyEndpointPolicy(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return EndpointPolicyToUnstructured(r), nil
}

func EndpointPolicyHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEndpointPolicy(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToEndpointPolicy(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyEndpointPolicy(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteEndpointPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEndpointPolicy(u)
	if err != nil {
		return err
	}
	return c.DeleteEndpointPolicy(ctx, r)
}

func EndpointPolicyID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToEndpointPolicy(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *EndpointPolicy) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"networkservices",
		"EndpointPolicy",
		"alpha",
	}
}

func (r *EndpointPolicy) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *EndpointPolicy) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *EndpointPolicy) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *EndpointPolicy) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *EndpointPolicy) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *EndpointPolicy) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *EndpointPolicy) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetEndpointPolicy(ctx, config, resource)
}

func (r *EndpointPolicy) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyEndpointPolicy(ctx, config, resource, opts...)
}

func (r *EndpointPolicy) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return EndpointPolicyHasDiff(ctx, config, resource, opts...)
}

func (r *EndpointPolicy) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteEndpointPolicy(ctx, config, resource)
}

func (r *EndpointPolicy) ID(resource *unstructured.Resource) (string, error) {
	return EndpointPolicyID(resource)
}

func init() {
	unstructured.Register(&EndpointPolicy{})
}
