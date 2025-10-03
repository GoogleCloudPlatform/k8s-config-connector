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
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/identitytoolkit/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Tenant struct{}

func TenantToUnstructured(r *dclService.Tenant) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "identitytoolkit",
			Version: "beta",
			Type:    "Tenant",
		},
		Object: make(map[string]interface{}),
	}
	if r.AllowPasswordSignup != nil {
		u.Object["allowPasswordSignup"] = *r.AllowPasswordSignup
	}
	if r.DisableAuth != nil {
		u.Object["disableAuth"] = *r.DisableAuth
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.EnableAnonymousUser != nil {
		u.Object["enableAnonymousUser"] = *r.EnableAnonymousUser
	}
	if r.EnableEmailLinkSignin != nil {
		u.Object["enableEmailLinkSignin"] = *r.EnableEmailLinkSignin
	}
	if r.MfaConfig != nil && r.MfaConfig != dclService.EmptyTenantMfaConfig {
		rMfaConfig := make(map[string]interface{})
		var rMfaConfigEnabledProviders []interface{}
		for _, rMfaConfigEnabledProvidersVal := range r.MfaConfig.EnabledProviders {
			rMfaConfigEnabledProviders = append(rMfaConfigEnabledProviders, string(rMfaConfigEnabledProvidersVal))
		}
		rMfaConfig["enabledProviders"] = rMfaConfigEnabledProviders
		if r.MfaConfig.State != nil {
			rMfaConfig["state"] = string(*r.MfaConfig.State)
		}
		u.Object["mfaConfig"] = rMfaConfig
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.TestPhoneNumbers != nil {
		rTestPhoneNumbers := make(map[string]interface{})
		for k, v := range r.TestPhoneNumbers {
			rTestPhoneNumbers[k] = v
		}
		u.Object["testPhoneNumbers"] = rTestPhoneNumbers
	}
	return u
}

func UnstructuredToTenant(u *unstructured.Resource) (*dclService.Tenant, error) {
	r := &dclService.Tenant{}
	if _, ok := u.Object["allowPasswordSignup"]; ok {
		if b, ok := u.Object["allowPasswordSignup"].(bool); ok {
			r.AllowPasswordSignup = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.AllowPasswordSignup: expected bool")
		}
	}
	if _, ok := u.Object["disableAuth"]; ok {
		if b, ok := u.Object["disableAuth"].(bool); ok {
			r.DisableAuth = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.DisableAuth: expected bool")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
		}
	}
	if _, ok := u.Object["enableAnonymousUser"]; ok {
		if b, ok := u.Object["enableAnonymousUser"].(bool); ok {
			r.EnableAnonymousUser = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.EnableAnonymousUser: expected bool")
		}
	}
	if _, ok := u.Object["enableEmailLinkSignin"]; ok {
		if b, ok := u.Object["enableEmailLinkSignin"].(bool); ok {
			r.EnableEmailLinkSignin = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.EnableEmailLinkSignin: expected bool")
		}
	}
	if _, ok := u.Object["mfaConfig"]; ok {
		if rMfaConfig, ok := u.Object["mfaConfig"].(map[string]interface{}); ok {
			r.MfaConfig = &dclService.TenantMfaConfig{}
			if _, ok := rMfaConfig["enabledProviders"]; ok {
				if s, ok := rMfaConfig["enabledProviders"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.MfaConfig.EnabledProviders = append(r.MfaConfig.EnabledProviders, dclService.TenantMfaConfigEnabledProvidersEnum(strval))
						}
					}
				} else {
					return nil, fmt.Errorf("r.MfaConfig.EnabledProviders: expected []interface{}")
				}
			}
			if _, ok := rMfaConfig["state"]; ok {
				if s, ok := rMfaConfig["state"].(string); ok {
					r.MfaConfig.State = dclService.TenantMfaConfigStateEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.MfaConfig.State: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.MfaConfig: expected map[string]interface{}")
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
	if _, ok := u.Object["testPhoneNumbers"]; ok {
		if rTestPhoneNumbers, ok := u.Object["testPhoneNumbers"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rTestPhoneNumbers {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.TestPhoneNumbers = m
		} else {
			return nil, fmt.Errorf("r.TestPhoneNumbers: expected map[string]interface{}")
		}
	}
	return r, nil
}

func GetTenant(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTenant(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetTenant(ctx, r)
	if err != nil {
		return nil, err
	}
	return TenantToUnstructured(r), nil
}

func ListTenant(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListTenant(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, TenantToUnstructured(r))
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

func ApplyTenant(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTenant(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToTenant(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyTenant(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return TenantToUnstructured(r), nil
}

func TenantHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTenant(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToTenant(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyTenant(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteTenant(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTenant(u)
	if err != nil {
		return err
	}
	return c.DeleteTenant(ctx, r)
}

func TenantID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToTenant(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Tenant) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"identitytoolkit",
		"Tenant",
		"beta",
	}
}

func (r *Tenant) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Tenant) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Tenant) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Tenant) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Tenant) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Tenant) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Tenant) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetTenant(ctx, config, resource)
}

func (r *Tenant) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyTenant(ctx, config, resource, opts...)
}

func (r *Tenant) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return TenantHasDiff(ctx, config, resource, opts...)
}

func (r *Tenant) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteTenant(ctx, config, resource)
}

func (r *Tenant) ID(resource *unstructured.Resource) (string, error) {
	return TenantID(resource)
}

func init() {
	unstructured.Register(&Tenant{})
}
