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
package identitytoolkit

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/identitytoolkit/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type TenantOAuthIdpConfig struct{}

func TenantOAuthIdpConfigToUnstructured(r *dclService.TenantOAuthIdpConfig) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "identitytoolkit",
			Version: "alpha",
			Type:    "TenantOAuthIdpConfig",
		},
		Object: make(map[string]interface{}),
	}
	if r.ClientId != nil {
		u.Object["clientId"] = *r.ClientId
	}
	if r.ClientSecret != nil {
		u.Object["clientSecret"] = *r.ClientSecret
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.Enabled != nil {
		u.Object["enabled"] = *r.Enabled
	}
	if r.Issuer != nil {
		u.Object["issuer"] = *r.Issuer
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.ResponseType != nil && r.ResponseType != dclService.EmptyTenantOAuthIdpConfigResponseType {
		rResponseType := make(map[string]interface{})
		if r.ResponseType.Code != nil {
			rResponseType["code"] = *r.ResponseType.Code
		}
		if r.ResponseType.IdToken != nil {
			rResponseType["idToken"] = *r.ResponseType.IdToken
		}
		if r.ResponseType.Token != nil {
			rResponseType["token"] = *r.ResponseType.Token
		}
		u.Object["responseType"] = rResponseType
	}
	if r.Tenant != nil {
		u.Object["tenant"] = *r.Tenant
	}
	return u
}

func UnstructuredToTenantOAuthIdpConfig(u *unstructured.Resource) (*dclService.TenantOAuthIdpConfig, error) {
	r := &dclService.TenantOAuthIdpConfig{}
	if _, ok := u.Object["clientId"]; ok {
		if s, ok := u.Object["clientId"].(string); ok {
			r.ClientId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ClientId: expected string")
		}
	}
	if _, ok := u.Object["clientSecret"]; ok {
		if s, ok := u.Object["clientSecret"].(string); ok {
			r.ClientSecret = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ClientSecret: expected string")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
		}
	}
	if _, ok := u.Object["enabled"]; ok {
		if b, ok := u.Object["enabled"].(bool); ok {
			r.Enabled = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.Enabled: expected bool")
		}
	}
	if _, ok := u.Object["issuer"]; ok {
		if s, ok := u.Object["issuer"].(string); ok {
			r.Issuer = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Issuer: expected string")
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
	if _, ok := u.Object["responseType"]; ok {
		if rResponseType, ok := u.Object["responseType"].(map[string]interface{}); ok {
			r.ResponseType = &dclService.TenantOAuthIdpConfigResponseType{}
			if _, ok := rResponseType["code"]; ok {
				if b, ok := rResponseType["code"].(bool); ok {
					r.ResponseType.Code = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.ResponseType.Code: expected bool")
				}
			}
			if _, ok := rResponseType["idToken"]; ok {
				if b, ok := rResponseType["idToken"].(bool); ok {
					r.ResponseType.IdToken = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.ResponseType.IdToken: expected bool")
				}
			}
			if _, ok := rResponseType["token"]; ok {
				if b, ok := rResponseType["token"].(bool); ok {
					r.ResponseType.Token = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.ResponseType.Token: expected bool")
				}
			}
		} else {
			return nil, fmt.Errorf("r.ResponseType: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["tenant"]; ok {
		if s, ok := u.Object["tenant"].(string); ok {
			r.Tenant = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Tenant: expected string")
		}
	}
	return r, nil
}

func GetTenantOAuthIdpConfig(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTenantOAuthIdpConfig(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetTenantOAuthIdpConfig(ctx, r)
	if err != nil {
		return nil, err
	}
	return TenantOAuthIdpConfigToUnstructured(r), nil
}

func ListTenantOAuthIdpConfig(ctx context.Context, config *dcl.Config, project string, tenant string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListTenantOAuthIdpConfig(ctx, project, tenant)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, TenantOAuthIdpConfigToUnstructured(r))
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

func ApplyTenantOAuthIdpConfig(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTenantOAuthIdpConfig(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToTenantOAuthIdpConfig(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyTenantOAuthIdpConfig(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return TenantOAuthIdpConfigToUnstructured(r), nil
}

func TenantOAuthIdpConfigHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTenantOAuthIdpConfig(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToTenantOAuthIdpConfig(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyTenantOAuthIdpConfig(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteTenantOAuthIdpConfig(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTenantOAuthIdpConfig(u)
	if err != nil {
		return err
	}
	return c.DeleteTenantOAuthIdpConfig(ctx, r)
}

func TenantOAuthIdpConfigID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToTenantOAuthIdpConfig(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *TenantOAuthIdpConfig) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"identitytoolkit",
		"TenantOAuthIdpConfig",
		"alpha",
	}
}

func (r *TenantOAuthIdpConfig) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *TenantOAuthIdpConfig) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *TenantOAuthIdpConfig) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *TenantOAuthIdpConfig) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *TenantOAuthIdpConfig) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *TenantOAuthIdpConfig) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *TenantOAuthIdpConfig) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetTenantOAuthIdpConfig(ctx, config, resource)
}

func (r *TenantOAuthIdpConfig) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyTenantOAuthIdpConfig(ctx, config, resource, opts...)
}

func (r *TenantOAuthIdpConfig) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return TenantOAuthIdpConfigHasDiff(ctx, config, resource, opts...)
}

func (r *TenantOAuthIdpConfig) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteTenantOAuthIdpConfig(ctx, config, resource)
}

func (r *TenantOAuthIdpConfig) ID(resource *unstructured.Resource) (string, error) {
	return TenantOAuthIdpConfigID(resource)
}

func init() {
	unstructured.Register(&TenantOAuthIdpConfig{})
}
