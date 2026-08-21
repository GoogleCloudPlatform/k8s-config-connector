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
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type ForwardingRule struct{}

func ForwardingRuleToUnstructured(r *dclService.ForwardingRule) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "compute",
			Version: "ga",
			Type:    "ForwardingRule",
		},
		Object: make(map[string]interface{}),
	}
	if r.AllPorts != nil {
		u.Object["allPorts"] = *r.AllPorts
	}
	if r.AllowGlobalAccess != nil {
		u.Object["allowGlobalAccess"] = *r.AllowGlobalAccess
	}
	if r.BackendService != nil {
		u.Object["backendService"] = *r.BackendService
	}
	if r.BaseForwardingRule != nil {
		u.Object["baseForwardingRule"] = *r.BaseForwardingRule
	}
	if r.CreationTimestamp != nil {
		u.Object["creationTimestamp"] = *r.CreationTimestamp
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.IPAddress != nil {
		u.Object["ipAddress"] = *r.IPAddress
	}
	if r.IPProtocol != nil {
		u.Object["ipProtocol"] = string(*r.IPProtocol)
	}
	if r.IPVersion != nil {
		u.Object["ipVersion"] = string(*r.IPVersion)
	}
	if r.IsMirroringCollector != nil {
		u.Object["isMirroringCollector"] = *r.IsMirroringCollector
	}
	if r.LabelFingerprint != nil {
		u.Object["labelFingerprint"] = *r.LabelFingerprint
	}
	if r.Labels != nil {
		rLabels := make(map[string]interface{})
		for k, v := range r.Labels {
			rLabels[k] = v
		}
		u.Object["labels"] = rLabels
	}
	if r.LoadBalancingScheme != nil {
		u.Object["loadBalancingScheme"] = string(*r.LoadBalancingScheme)
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	var rMetadataFilter []interface{}
	for _, rMetadataFilterVal := range r.MetadataFilter {
		rMetadataFilterObject := make(map[string]interface{})
		var rMetadataFilterValFilterLabel []interface{}
		for _, rMetadataFilterValFilterLabelVal := range rMetadataFilterVal.FilterLabel {
			rMetadataFilterValFilterLabelObject := make(map[string]interface{})
			if rMetadataFilterValFilterLabelVal.Name != nil {
				rMetadataFilterValFilterLabelObject["name"] = *rMetadataFilterValFilterLabelVal.Name
			}
			if rMetadataFilterValFilterLabelVal.Value != nil {
				rMetadataFilterValFilterLabelObject["value"] = *rMetadataFilterValFilterLabelVal.Value
			}
			rMetadataFilterValFilterLabel = append(rMetadataFilterValFilterLabel, rMetadataFilterValFilterLabelObject)
		}
		rMetadataFilterObject["filterLabel"] = rMetadataFilterValFilterLabel
		if rMetadataFilterVal.FilterMatchCriteria != nil {
			rMetadataFilterObject["filterMatchCriteria"] = string(*rMetadataFilterVal.FilterMatchCriteria)
		}
		rMetadataFilter = append(rMetadataFilter, rMetadataFilterObject)
	}
	u.Object["metadataFilter"] = rMetadataFilter
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Network != nil {
		u.Object["network"] = *r.Network
	}
	if r.NetworkTier != nil {
		u.Object["networkTier"] = string(*r.NetworkTier)
	}
	if r.PortRange != nil {
		u.Object["portRange"] = *r.PortRange
	}
	var rPorts []interface{}
	for _, rPortsVal := range r.Ports {
		rPorts = append(rPorts, rPortsVal)
	}
	u.Object["ports"] = rPorts
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.PscConnectionId != nil {
		u.Object["pscConnectionId"] = *r.PscConnectionId
	}
	if r.PscConnectionStatus != nil {
		u.Object["pscConnectionStatus"] = string(*r.PscConnectionStatus)
	}
	if r.Region != nil {
		u.Object["region"] = *r.Region
	}
	if r.SelfLink != nil {
		u.Object["selfLink"] = *r.SelfLink
	}
	var rServiceDirectoryRegistrations []interface{}
	for _, rServiceDirectoryRegistrationsVal := range r.ServiceDirectoryRegistrations {
		rServiceDirectoryRegistrationsObject := make(map[string]interface{})
		if rServiceDirectoryRegistrationsVal.Namespace != nil {
			rServiceDirectoryRegistrationsObject["namespace"] = *rServiceDirectoryRegistrationsVal.Namespace
		}
		if rServiceDirectoryRegistrationsVal.Service != nil {
			rServiceDirectoryRegistrationsObject["service"] = *rServiceDirectoryRegistrationsVal.Service
		}
		rServiceDirectoryRegistrations = append(rServiceDirectoryRegistrations, rServiceDirectoryRegistrationsObject)
	}
	u.Object["serviceDirectoryRegistrations"] = rServiceDirectoryRegistrations
	if r.ServiceLabel != nil {
		u.Object["serviceLabel"] = *r.ServiceLabel
	}
	if r.ServiceName != nil {
		u.Object["serviceName"] = *r.ServiceName
	}
	var rSourceIPRanges []interface{}
	for _, rSourceIPRangesVal := range r.SourceIPRanges {
		rSourceIPRanges = append(rSourceIPRanges, rSourceIPRangesVal)
	}
	u.Object["sourceIPRanges"] = rSourceIPRanges
	if r.Subnetwork != nil {
		u.Object["subnetwork"] = *r.Subnetwork
	}
	if r.Target != nil {
		u.Object["target"] = *r.Target
	}
	return u
}

func UnstructuredToForwardingRule(u *unstructured.Resource) (*dclService.ForwardingRule, error) {
	r := &dclService.ForwardingRule{}
	if _, ok := u.Object["allPorts"]; ok {
		if b, ok := u.Object["allPorts"].(bool); ok {
			r.AllPorts = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.AllPorts: expected bool")
		}
	}
	if _, ok := u.Object["allowGlobalAccess"]; ok {
		if b, ok := u.Object["allowGlobalAccess"].(bool); ok {
			r.AllowGlobalAccess = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.AllowGlobalAccess: expected bool")
		}
	}
	if _, ok := u.Object["backendService"]; ok {
		if s, ok := u.Object["backendService"].(string); ok {
			r.BackendService = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.BackendService: expected string")
		}
	}
	if _, ok := u.Object["baseForwardingRule"]; ok {
		if s, ok := u.Object["baseForwardingRule"].(string); ok {
			r.BaseForwardingRule = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.BaseForwardingRule: expected string")
		}
	}
	if _, ok := u.Object["creationTimestamp"]; ok {
		if s, ok := u.Object["creationTimestamp"].(string); ok {
			r.CreationTimestamp = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreationTimestamp: expected string")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["ipAddress"]; ok {
		if s, ok := u.Object["ipAddress"].(string); ok {
			r.IPAddress = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.IPAddress: expected string")
		}
	}
	if _, ok := u.Object["ipProtocol"]; ok {
		if s, ok := u.Object["ipProtocol"].(string); ok {
			r.IPProtocol = dclService.ForwardingRuleIPProtocolEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.IPProtocol: expected string")
		}
	}
	if _, ok := u.Object["ipVersion"]; ok {
		if s, ok := u.Object["ipVersion"].(string); ok {
			r.IPVersion = dclService.ForwardingRuleIPVersionEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.IPVersion: expected string")
		}
	}
	if _, ok := u.Object["isMirroringCollector"]; ok {
		if b, ok := u.Object["isMirroringCollector"].(bool); ok {
			r.IsMirroringCollector = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.IsMirroringCollector: expected bool")
		}
	}
	if _, ok := u.Object["labelFingerprint"]; ok {
		if s, ok := u.Object["labelFingerprint"].(string); ok {
			r.LabelFingerprint = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.LabelFingerprint: expected string")
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
	if _, ok := u.Object["loadBalancingScheme"]; ok {
		if s, ok := u.Object["loadBalancingScheme"].(string); ok {
			r.LoadBalancingScheme = dclService.ForwardingRuleLoadBalancingSchemeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.LoadBalancingScheme: expected string")
		}
	}
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["metadataFilter"]; ok {
		if s, ok := u.Object["metadataFilter"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rMetadataFilter dclService.ForwardingRuleMetadataFilter
					if _, ok := objval["filterLabel"]; ok {
						if s, ok := objval["filterLabel"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rMetadataFilterFilterLabel dclService.ForwardingRuleMetadataFilterFilterLabel
									if _, ok := objval["name"]; ok {
										if s, ok := objval["name"].(string); ok {
											rMetadataFilterFilterLabel.Name = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rMetadataFilterFilterLabel.Name: expected string")
										}
									}
									if _, ok := objval["value"]; ok {
										if s, ok := objval["value"].(string); ok {
											rMetadataFilterFilterLabel.Value = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rMetadataFilterFilterLabel.Value: expected string")
										}
									}
									rMetadataFilter.FilterLabel = append(rMetadataFilter.FilterLabel, rMetadataFilterFilterLabel)
								}
							}
						} else {
							return nil, fmt.Errorf("rMetadataFilter.FilterLabel: expected []interface{}")
						}
					}
					if _, ok := objval["filterMatchCriteria"]; ok {
						if s, ok := objval["filterMatchCriteria"].(string); ok {
							rMetadataFilter.FilterMatchCriteria = dclService.ForwardingRuleMetadataFilterFilterMatchCriteriaEnumRef(s)
						} else {
							return nil, fmt.Errorf("rMetadataFilter.FilterMatchCriteria: expected string")
						}
					}
					r.MetadataFilter = append(r.MetadataFilter, rMetadataFilter)
				}
			}
		} else {
			return nil, fmt.Errorf("r.MetadataFilter: expected []interface{}")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["network"]; ok {
		if s, ok := u.Object["network"].(string); ok {
			r.Network = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Network: expected string")
		}
	}
	if _, ok := u.Object["networkTier"]; ok {
		if s, ok := u.Object["networkTier"].(string); ok {
			r.NetworkTier = dclService.ForwardingRuleNetworkTierEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.NetworkTier: expected string")
		}
	}
	if _, ok := u.Object["portRange"]; ok {
		if s, ok := u.Object["portRange"].(string); ok {
			r.PortRange = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.PortRange: expected string")
		}
	}
	if _, ok := u.Object["ports"]; ok {
		if s, ok := u.Object["ports"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.Ports = append(r.Ports, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Ports: expected []interface{}")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["pscConnectionId"]; ok {
		if s, ok := u.Object["pscConnectionId"].(string); ok {
			r.PscConnectionId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.PscConnectionId: expected string")
		}
	}
	if _, ok := u.Object["pscConnectionStatus"]; ok {
		if s, ok := u.Object["pscConnectionStatus"].(string); ok {
			r.PscConnectionStatus = dclService.ForwardingRulePscConnectionStatusEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.PscConnectionStatus: expected string")
		}
	}
	if _, ok := u.Object["region"]; ok {
		if s, ok := u.Object["region"].(string); ok {
			r.Region = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Region: expected string")
		}
	}
	if _, ok := u.Object["selfLink"]; ok {
		if s, ok := u.Object["selfLink"].(string); ok {
			r.SelfLink = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.SelfLink: expected string")
		}
	}
	if _, ok := u.Object["serviceDirectoryRegistrations"]; ok {
		if s, ok := u.Object["serviceDirectoryRegistrations"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rServiceDirectoryRegistrations dclService.ForwardingRuleServiceDirectoryRegistrations
					if _, ok := objval["namespace"]; ok {
						if s, ok := objval["namespace"].(string); ok {
							rServiceDirectoryRegistrations.Namespace = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rServiceDirectoryRegistrations.Namespace: expected string")
						}
					}
					if _, ok := objval["service"]; ok {
						if s, ok := objval["service"].(string); ok {
							rServiceDirectoryRegistrations.Service = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rServiceDirectoryRegistrations.Service: expected string")
						}
					}
					r.ServiceDirectoryRegistrations = append(r.ServiceDirectoryRegistrations, rServiceDirectoryRegistrations)
				}
			}
		} else {
			return nil, fmt.Errorf("r.ServiceDirectoryRegistrations: expected []interface{}")
		}
	}
	if _, ok := u.Object["serviceLabel"]; ok {
		if s, ok := u.Object["serviceLabel"].(string); ok {
			r.ServiceLabel = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ServiceLabel: expected string")
		}
	}
	if _, ok := u.Object["serviceName"]; ok {
		if s, ok := u.Object["serviceName"].(string); ok {
			r.ServiceName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ServiceName: expected string")
		}
	}
	if _, ok := u.Object["sourceIPRanges"]; ok {
		if s, ok := u.Object["sourceIPRanges"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.SourceIPRanges = append(r.SourceIPRanges, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.SourceIPRanges: expected []interface{}")
		}
	}
	if _, ok := u.Object["subnetwork"]; ok {
		if s, ok := u.Object["subnetwork"].(string); ok {
			r.Subnetwork = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Subnetwork: expected string")
		}
	}
	if _, ok := u.Object["target"]; ok {
		if s, ok := u.Object["target"].(string); ok {
			r.Target = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Target: expected string")
		}
	}
	return r, nil
}

func GetForwardingRule(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToForwardingRule(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetForwardingRule(ctx, r)
	if err != nil {
		return nil, err
	}
	return ForwardingRuleToUnstructured(r), nil
}

func ListForwardingRule(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListForwardingRule(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, ForwardingRuleToUnstructured(r))
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

func ApplyForwardingRule(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToForwardingRule(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToForwardingRule(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyForwardingRule(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return ForwardingRuleToUnstructured(r), nil
}

func ForwardingRuleHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToForwardingRule(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToForwardingRule(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyForwardingRule(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteForwardingRule(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToForwardingRule(u)
	if err != nil {
		return err
	}
	return c.DeleteForwardingRule(ctx, r)
}

func ForwardingRuleID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToForwardingRule(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *ForwardingRule) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"compute",
		"ForwardingRule",
		"ga",
	}
}

func (r *ForwardingRule) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *ForwardingRule) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *ForwardingRule) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *ForwardingRule) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *ForwardingRule) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *ForwardingRule) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *ForwardingRule) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetForwardingRule(ctx, config, resource)
}

func (r *ForwardingRule) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyForwardingRule(ctx, config, resource, opts...)
}

func (r *ForwardingRule) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return ForwardingRuleHasDiff(ctx, config, resource, opts...)
}

func (r *ForwardingRule) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteForwardingRule(ctx, config, resource)
}

func (r *ForwardingRule) ID(resource *unstructured.Resource) (string, error) {
	return ForwardingRuleID(resource)
}

func init() {
	unstructured.Register(&ForwardingRule{})
}
