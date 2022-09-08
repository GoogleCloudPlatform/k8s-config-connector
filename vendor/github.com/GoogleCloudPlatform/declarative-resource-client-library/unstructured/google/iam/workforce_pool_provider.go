// Copyright 2022 Google LLC. All Rights Reserved.
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
package iam

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type WorkforcePoolProvider struct{}

func WorkforcePoolProviderToUnstructured(r *dclService.WorkforcePoolProvider) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "iam",
			Version: "ga",
			Type:    "WorkforcePoolProvider",
		},
		Object: make(map[string]interface{}),
	}
	if r.AttributeCondition != nil {
		u.Object["attributeCondition"] = *r.AttributeCondition
	}
	if r.AttributeMapping != nil {
		rAttributeMapping := make(map[string]interface{})
		for k, v := range r.AttributeMapping {
			rAttributeMapping[k] = v
		}
		u.Object["attributeMapping"] = rAttributeMapping
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.Disabled != nil {
		u.Object["disabled"] = *r.Disabled
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Oidc != nil && r.Oidc != dclService.EmptyWorkforcePoolProviderOidc {
		rOidc := make(map[string]interface{})
		if r.Oidc.ClientId != nil {
			rOidc["clientId"] = *r.Oidc.ClientId
		}
		if r.Oidc.IssuerUri != nil {
			rOidc["issuerUri"] = *r.Oidc.IssuerUri
		}
		u.Object["oidc"] = rOidc
	}
	if r.Saml != nil && r.Saml != dclService.EmptyWorkforcePoolProviderSaml {
		rSaml := make(map[string]interface{})
		if r.Saml.IdpMetadataXml != nil {
			rSaml["idpMetadataXml"] = *r.Saml.IdpMetadataXml
		}
		u.Object["saml"] = rSaml
	}
	if r.State != nil {
		u.Object["state"] = string(*r.State)
	}
	if r.WorkforcePool != nil {
		u.Object["workforcePool"] = *r.WorkforcePool
	}
	return u
}

func UnstructuredToWorkforcePoolProvider(u *unstructured.Resource) (*dclService.WorkforcePoolProvider, error) {
	r := &dclService.WorkforcePoolProvider{}
	if _, ok := u.Object["attributeCondition"]; ok {
		if s, ok := u.Object["attributeCondition"].(string); ok {
			r.AttributeCondition = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.AttributeCondition: expected string")
		}
	}
	if _, ok := u.Object["attributeMapping"]; ok {
		if rAttributeMapping, ok := u.Object["attributeMapping"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rAttributeMapping {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.AttributeMapping = m
		} else {
			return nil, fmt.Errorf("r.AttributeMapping: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["disabled"]; ok {
		if b, ok := u.Object["disabled"].(bool); ok {
			r.Disabled = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.Disabled: expected bool")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
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
	if _, ok := u.Object["oidc"]; ok {
		if rOidc, ok := u.Object["oidc"].(map[string]interface{}); ok {
			r.Oidc = &dclService.WorkforcePoolProviderOidc{}
			if _, ok := rOidc["clientId"]; ok {
				if s, ok := rOidc["clientId"].(string); ok {
					r.Oidc.ClientId = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Oidc.ClientId: expected string")
				}
			}
			if _, ok := rOidc["issuerUri"]; ok {
				if s, ok := rOidc["issuerUri"].(string); ok {
					r.Oidc.IssuerUri = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Oidc.IssuerUri: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Oidc: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["saml"]; ok {
		if rSaml, ok := u.Object["saml"].(map[string]interface{}); ok {
			r.Saml = &dclService.WorkforcePoolProviderSaml{}
			if _, ok := rSaml["idpMetadataXml"]; ok {
				if s, ok := rSaml["idpMetadataXml"].(string); ok {
					r.Saml.IdpMetadataXml = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Saml.IdpMetadataXml: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Saml: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["state"]; ok {
		if s, ok := u.Object["state"].(string); ok {
			r.State = dclService.WorkforcePoolProviderStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.State: expected string")
		}
	}
	if _, ok := u.Object["workforcePool"]; ok {
		if s, ok := u.Object["workforcePool"].(string); ok {
			r.WorkforcePool = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.WorkforcePool: expected string")
		}
	}
	return r, nil
}

func GetWorkforcePoolProvider(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWorkforcePoolProvider(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetWorkforcePoolProvider(ctx, r)
	if err != nil {
		return nil, err
	}
	return WorkforcePoolProviderToUnstructured(r), nil
}

func ListWorkforcePoolProvider(ctx context.Context, config *dcl.Config, location string, workforcePool string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListWorkforcePoolProvider(ctx, location, workforcePool)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, WorkforcePoolProviderToUnstructured(r))
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

func ApplyWorkforcePoolProvider(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWorkforcePoolProvider(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToWorkforcePoolProvider(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyWorkforcePoolProvider(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return WorkforcePoolProviderToUnstructured(r), nil
}

func WorkforcePoolProviderHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWorkforcePoolProvider(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToWorkforcePoolProvider(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyWorkforcePoolProvider(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteWorkforcePoolProvider(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWorkforcePoolProvider(u)
	if err != nil {
		return err
	}
	return c.DeleteWorkforcePoolProvider(ctx, r)
}

func WorkforcePoolProviderID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToWorkforcePoolProvider(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *WorkforcePoolProvider) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"iam",
		"WorkforcePoolProvider",
		"ga",
	}
}

func (r *WorkforcePoolProvider) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *WorkforcePoolProvider) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *WorkforcePoolProvider) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *WorkforcePoolProvider) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *WorkforcePoolProvider) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *WorkforcePoolProvider) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *WorkforcePoolProvider) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetWorkforcePoolProvider(ctx, config, resource)
}

func (r *WorkforcePoolProvider) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyWorkforcePoolProvider(ctx, config, resource, opts...)
}

func (r *WorkforcePoolProvider) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return WorkforcePoolProviderHasDiff(ctx, config, resource, opts...)
}

func (r *WorkforcePoolProvider) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteWorkforcePoolProvider(ctx, config, resource)
}

func (r *WorkforcePoolProvider) ID(resource *unstructured.Resource) (string, error) {
	return WorkforcePoolProviderID(resource)
}

func init() {
	unstructured.Register(&WorkforcePoolProvider{})
}
