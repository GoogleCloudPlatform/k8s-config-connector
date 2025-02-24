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

type ServiceAttachment struct{}

func ServiceAttachmentToUnstructured(r *dclService.ServiceAttachment) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "compute",
			Version: "ga",
			Type:    "ServiceAttachment",
		},
		Object: make(map[string]interface{}),
	}
	var rConnectedEndpoints []interface{}
	for _, rConnectedEndpointsVal := range r.ConnectedEndpoints {
		rConnectedEndpointsObject := make(map[string]interface{})
		if rConnectedEndpointsVal.Endpoint != nil {
			rConnectedEndpointsObject["endpoint"] = *rConnectedEndpointsVal.Endpoint
		}
		if rConnectedEndpointsVal.PscConnectionId != nil {
			rConnectedEndpointsObject["pscConnectionId"] = *rConnectedEndpointsVal.PscConnectionId
		}
		if rConnectedEndpointsVal.Status != nil {
			rConnectedEndpointsObject["status"] = string(*rConnectedEndpointsVal.Status)
		}
		rConnectedEndpoints = append(rConnectedEndpoints, rConnectedEndpointsObject)
	}
	u.Object["connectedEndpoints"] = rConnectedEndpoints
	if r.ConnectionPreference != nil {
		u.Object["connectionPreference"] = string(*r.ConnectionPreference)
	}
	var rConsumerAcceptLists []interface{}
	for _, rConsumerAcceptListsVal := range r.ConsumerAcceptLists {
		rConsumerAcceptListsObject := make(map[string]interface{})
		if rConsumerAcceptListsVal.ConnectionLimit != nil {
			rConsumerAcceptListsObject["connectionLimit"] = *rConsumerAcceptListsVal.ConnectionLimit
		}
		if rConsumerAcceptListsVal.ProjectIdOrNum != nil {
			rConsumerAcceptListsObject["projectIdOrNum"] = *rConsumerAcceptListsVal.ProjectIdOrNum
		}
		rConsumerAcceptLists = append(rConsumerAcceptLists, rConsumerAcceptListsObject)
	}
	u.Object["consumerAcceptLists"] = rConsumerAcceptLists
	var rConsumerRejectLists []interface{}
	for _, rConsumerRejectListsVal := range r.ConsumerRejectLists {
		rConsumerRejectLists = append(rConsumerRejectLists, rConsumerRejectListsVal)
	}
	u.Object["consumerRejectLists"] = rConsumerRejectLists
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.EnableProxyProtocol != nil {
		u.Object["enableProxyProtocol"] = *r.EnableProxyProtocol
	}
	if r.Fingerprint != nil {
		u.Object["fingerprint"] = *r.Fingerprint
	}
	if r.Id != nil {
		u.Object["id"] = *r.Id
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	var rNatSubnets []interface{}
	for _, rNatSubnetsVal := range r.NatSubnets {
		rNatSubnets = append(rNatSubnets, rNatSubnetsVal)
	}
	u.Object["natSubnets"] = rNatSubnets
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.PscServiceAttachmentId != nil && r.PscServiceAttachmentId != dclService.EmptyServiceAttachmentPscServiceAttachmentId {
		rPscServiceAttachmentId := make(map[string]interface{})
		if r.PscServiceAttachmentId.High != nil {
			rPscServiceAttachmentId["high"] = *r.PscServiceAttachmentId.High
		}
		if r.PscServiceAttachmentId.Low != nil {
			rPscServiceAttachmentId["low"] = *r.PscServiceAttachmentId.Low
		}
		u.Object["pscServiceAttachmentId"] = rPscServiceAttachmentId
	}
	if r.Region != nil {
		u.Object["region"] = *r.Region
	}
	if r.SelfLink != nil {
		u.Object["selfLink"] = *r.SelfLink
	}
	if r.TargetService != nil {
		u.Object["targetService"] = *r.TargetService
	}
	return u
}

func UnstructuredToServiceAttachment(u *unstructured.Resource) (*dclService.ServiceAttachment, error) {
	r := &dclService.ServiceAttachment{}
	if _, ok := u.Object["connectedEndpoints"]; ok {
		if s, ok := u.Object["connectedEndpoints"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rConnectedEndpoints dclService.ServiceAttachmentConnectedEndpoints
					if _, ok := objval["endpoint"]; ok {
						if s, ok := objval["endpoint"].(string); ok {
							rConnectedEndpoints.Endpoint = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rConnectedEndpoints.Endpoint: expected string")
						}
					}
					if _, ok := objval["pscConnectionId"]; ok {
						if i, ok := objval["pscConnectionId"].(int64); ok {
							rConnectedEndpoints.PscConnectionId = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("rConnectedEndpoints.PscConnectionId: expected int64")
						}
					}
					if _, ok := objval["status"]; ok {
						if s, ok := objval["status"].(string); ok {
							rConnectedEndpoints.Status = dclService.ServiceAttachmentConnectedEndpointsStatusEnumRef(s)
						} else {
							return nil, fmt.Errorf("rConnectedEndpoints.Status: expected string")
						}
					}
					r.ConnectedEndpoints = append(r.ConnectedEndpoints, rConnectedEndpoints)
				}
			}
		} else {
			return nil, fmt.Errorf("r.ConnectedEndpoints: expected []interface{}")
		}
	}
	if _, ok := u.Object["connectionPreference"]; ok {
		if s, ok := u.Object["connectionPreference"].(string); ok {
			r.ConnectionPreference = dclService.ServiceAttachmentConnectionPreferenceEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.ConnectionPreference: expected string")
		}
	}
	if _, ok := u.Object["consumerAcceptLists"]; ok {
		if s, ok := u.Object["consumerAcceptLists"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rConsumerAcceptLists dclService.ServiceAttachmentConsumerAcceptLists
					if _, ok := objval["connectionLimit"]; ok {
						if i, ok := objval["connectionLimit"].(int64); ok {
							rConsumerAcceptLists.ConnectionLimit = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("rConsumerAcceptLists.ConnectionLimit: expected int64")
						}
					}
					if _, ok := objval["projectIdOrNum"]; ok {
						if s, ok := objval["projectIdOrNum"].(string); ok {
							rConsumerAcceptLists.ProjectIdOrNum = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rConsumerAcceptLists.ProjectIdOrNum: expected string")
						}
					}
					r.ConsumerAcceptLists = append(r.ConsumerAcceptLists, rConsumerAcceptLists)
				}
			}
		} else {
			return nil, fmt.Errorf("r.ConsumerAcceptLists: expected []interface{}")
		}
	}
	if _, ok := u.Object["consumerRejectLists"]; ok {
		if s, ok := u.Object["consumerRejectLists"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.ConsumerRejectLists = append(r.ConsumerRejectLists, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.ConsumerRejectLists: expected []interface{}")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["enableProxyProtocol"]; ok {
		if b, ok := u.Object["enableProxyProtocol"].(bool); ok {
			r.EnableProxyProtocol = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.EnableProxyProtocol: expected bool")
		}
	}
	if _, ok := u.Object["fingerprint"]; ok {
		if s, ok := u.Object["fingerprint"].(string); ok {
			r.Fingerprint = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Fingerprint: expected string")
		}
	}
	if _, ok := u.Object["id"]; ok {
		if i, ok := u.Object["id"].(int64); ok {
			r.Id = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.Id: expected int64")
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
	if _, ok := u.Object["natSubnets"]; ok {
		if s, ok := u.Object["natSubnets"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.NatSubnets = append(r.NatSubnets, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.NatSubnets: expected []interface{}")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["pscServiceAttachmentId"]; ok {
		if rPscServiceAttachmentId, ok := u.Object["pscServiceAttachmentId"].(map[string]interface{}); ok {
			r.PscServiceAttachmentId = &dclService.ServiceAttachmentPscServiceAttachmentId{}
			if _, ok := rPscServiceAttachmentId["high"]; ok {
				if i, ok := rPscServiceAttachmentId["high"].(int64); ok {
					r.PscServiceAttachmentId.High = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.PscServiceAttachmentId.High: expected int64")
				}
			}
			if _, ok := rPscServiceAttachmentId["low"]; ok {
				if i, ok := rPscServiceAttachmentId["low"].(int64); ok {
					r.PscServiceAttachmentId.Low = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.PscServiceAttachmentId.Low: expected int64")
				}
			}
		} else {
			return nil, fmt.Errorf("r.PscServiceAttachmentId: expected map[string]interface{}")
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
	if _, ok := u.Object["targetService"]; ok {
		if s, ok := u.Object["targetService"].(string); ok {
			r.TargetService = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.TargetService: expected string")
		}
	}
	return r, nil
}

func GetServiceAttachment(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToServiceAttachment(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetServiceAttachment(ctx, r)
	if err != nil {
		return nil, err
	}
	return ServiceAttachmentToUnstructured(r), nil
}

func ListServiceAttachment(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListServiceAttachment(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, ServiceAttachmentToUnstructured(r))
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

func ApplyServiceAttachment(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToServiceAttachment(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToServiceAttachment(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyServiceAttachment(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return ServiceAttachmentToUnstructured(r), nil
}

func ServiceAttachmentHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToServiceAttachment(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToServiceAttachment(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyServiceAttachment(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteServiceAttachment(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToServiceAttachment(u)
	if err != nil {
		return err
	}
	return c.DeleteServiceAttachment(ctx, r)
}

func ServiceAttachmentID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToServiceAttachment(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *ServiceAttachment) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"compute",
		"ServiceAttachment",
		"ga",
	}
}

func (r *ServiceAttachment) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *ServiceAttachment) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *ServiceAttachment) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *ServiceAttachment) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *ServiceAttachment) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *ServiceAttachment) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *ServiceAttachment) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetServiceAttachment(ctx, config, resource)
}

func (r *ServiceAttachment) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyServiceAttachment(ctx, config, resource, opts...)
}

func (r *ServiceAttachment) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return ServiceAttachmentHasDiff(ctx, config, resource, opts...)
}

func (r *ServiceAttachment) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteServiceAttachment(ctx, config, resource)
}

func (r *ServiceAttachment) ID(resource *unstructured.Resource) (string, error) {
	return ServiceAttachmentID(resource)
}

func init() {
	unstructured.Register(&ServiceAttachment{})
}
