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
package apigee

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apigee/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Organization struct{}

func OrganizationToUnstructured(r *dclService.Organization) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "apigee",
			Version: "beta",
			Type:    "Organization",
		},
		Object: make(map[string]interface{}),
	}
	if r.AddonsConfig != nil && r.AddonsConfig != dclService.EmptyOrganizationAddonsConfig {
		rAddonsConfig := make(map[string]interface{})
		if r.AddonsConfig.AdvancedApiOpsConfig != nil && r.AddonsConfig.AdvancedApiOpsConfig != dclService.EmptyOrganizationAddonsConfigAdvancedApiOpsConfig {
			rAddonsConfigAdvancedApiOpsConfig := make(map[string]interface{})
			if r.AddonsConfig.AdvancedApiOpsConfig.Enabled != nil {
				rAddonsConfigAdvancedApiOpsConfig["enabled"] = *r.AddonsConfig.AdvancedApiOpsConfig.Enabled
			}
			rAddonsConfig["advancedApiOpsConfig"] = rAddonsConfigAdvancedApiOpsConfig
		}
		if r.AddonsConfig.MonetizationConfig != nil && r.AddonsConfig.MonetizationConfig != dclService.EmptyOrganizationAddonsConfigMonetizationConfig {
			rAddonsConfigMonetizationConfig := make(map[string]interface{})
			if r.AddonsConfig.MonetizationConfig.Enabled != nil {
				rAddonsConfigMonetizationConfig["enabled"] = *r.AddonsConfig.MonetizationConfig.Enabled
			}
			rAddonsConfig["monetizationConfig"] = rAddonsConfigMonetizationConfig
		}
		u.Object["addonsConfig"] = rAddonsConfig
	}
	if r.AnalyticsRegion != nil {
		u.Object["analyticsRegion"] = *r.AnalyticsRegion
	}
	if r.AuthorizedNetwork != nil {
		u.Object["authorizedNetwork"] = *r.AuthorizedNetwork
	}
	if r.BillingType != nil {
		u.Object["billingType"] = string(*r.BillingType)
	}
	if r.CaCertificate != nil {
		u.Object["caCertificate"] = *r.CaCertificate
	}
	if r.CreatedAt != nil {
		u.Object["createdAt"] = *r.CreatedAt
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	var rEnvironments []interface{}
	for _, rEnvironmentsVal := range r.Environments {
		rEnvironments = append(rEnvironments, rEnvironmentsVal)
	}
	u.Object["environments"] = rEnvironments
	if r.ExpiresAt != nil {
		u.Object["expiresAt"] = *r.ExpiresAt
	}
	if r.LastModifiedAt != nil {
		u.Object["lastModifiedAt"] = *r.LastModifiedAt
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.ProjectId != nil {
		u.Object["projectId"] = *r.ProjectId
	}
	if r.Properties != nil {
		rProperties := make(map[string]interface{})
		for k, v := range r.Properties {
			rProperties[k] = v
		}
		u.Object["properties"] = rProperties
	}
	if r.RuntimeDatabaseEncryptionKeyName != nil {
		u.Object["runtimeDatabaseEncryptionKeyName"] = *r.RuntimeDatabaseEncryptionKeyName
	}
	if r.RuntimeType != nil {
		u.Object["runtimeType"] = string(*r.RuntimeType)
	}
	if r.State != nil {
		u.Object["state"] = string(*r.State)
	}
	if r.SubscriptionType != nil {
		u.Object["subscriptionType"] = string(*r.SubscriptionType)
	}
	return u
}

func UnstructuredToOrganization(u *unstructured.Resource) (*dclService.Organization, error) {
	r := &dclService.Organization{}
	if _, ok := u.Object["addonsConfig"]; ok {
		if rAddonsConfig, ok := u.Object["addonsConfig"].(map[string]interface{}); ok {
			r.AddonsConfig = &dclService.OrganizationAddonsConfig{}
			if _, ok := rAddonsConfig["advancedApiOpsConfig"]; ok {
				if rAddonsConfigAdvancedApiOpsConfig, ok := rAddonsConfig["advancedApiOpsConfig"].(map[string]interface{}); ok {
					r.AddonsConfig.AdvancedApiOpsConfig = &dclService.OrganizationAddonsConfigAdvancedApiOpsConfig{}
					if _, ok := rAddonsConfigAdvancedApiOpsConfig["enabled"]; ok {
						if b, ok := rAddonsConfigAdvancedApiOpsConfig["enabled"].(bool); ok {
							r.AddonsConfig.AdvancedApiOpsConfig.Enabled = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.AddonsConfig.AdvancedApiOpsConfig.Enabled: expected bool")
						}
					}
				} else {
					return nil, fmt.Errorf("r.AddonsConfig.AdvancedApiOpsConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rAddonsConfig["monetizationConfig"]; ok {
				if rAddonsConfigMonetizationConfig, ok := rAddonsConfig["monetizationConfig"].(map[string]interface{}); ok {
					r.AddonsConfig.MonetizationConfig = &dclService.OrganizationAddonsConfigMonetizationConfig{}
					if _, ok := rAddonsConfigMonetizationConfig["enabled"]; ok {
						if b, ok := rAddonsConfigMonetizationConfig["enabled"].(bool); ok {
							r.AddonsConfig.MonetizationConfig.Enabled = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.AddonsConfig.MonetizationConfig.Enabled: expected bool")
						}
					}
				} else {
					return nil, fmt.Errorf("r.AddonsConfig.MonetizationConfig: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.AddonsConfig: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["analyticsRegion"]; ok {
		if s, ok := u.Object["analyticsRegion"].(string); ok {
			r.AnalyticsRegion = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.AnalyticsRegion: expected string")
		}
	}
	if _, ok := u.Object["authorizedNetwork"]; ok {
		if s, ok := u.Object["authorizedNetwork"].(string); ok {
			r.AuthorizedNetwork = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.AuthorizedNetwork: expected string")
		}
	}
	if _, ok := u.Object["billingType"]; ok {
		if s, ok := u.Object["billingType"].(string); ok {
			r.BillingType = dclService.OrganizationBillingTypeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.BillingType: expected string")
		}
	}
	if _, ok := u.Object["caCertificate"]; ok {
		if s, ok := u.Object["caCertificate"].(string); ok {
			r.CaCertificate = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CaCertificate: expected string")
		}
	}
	if _, ok := u.Object["createdAt"]; ok {
		if i, ok := u.Object["createdAt"].(int64); ok {
			r.CreatedAt = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.CreatedAt: expected int64")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
		}
	}
	if _, ok := u.Object["environments"]; ok {
		if s, ok := u.Object["environments"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.Environments = append(r.Environments, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Environments: expected []interface{}")
		}
	}
	if _, ok := u.Object["expiresAt"]; ok {
		if i, ok := u.Object["expiresAt"].(int64); ok {
			r.ExpiresAt = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.ExpiresAt: expected int64")
		}
	}
	if _, ok := u.Object["lastModifiedAt"]; ok {
		if i, ok := u.Object["lastModifiedAt"].(int64); ok {
			r.LastModifiedAt = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.LastModifiedAt: expected int64")
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
	if _, ok := u.Object["projectId"]; ok {
		if s, ok := u.Object["projectId"].(string); ok {
			r.ProjectId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ProjectId: expected string")
		}
	}
	if _, ok := u.Object["properties"]; ok {
		if rProperties, ok := u.Object["properties"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rProperties {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.Properties = m
		} else {
			return nil, fmt.Errorf("r.Properties: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["runtimeDatabaseEncryptionKeyName"]; ok {
		if s, ok := u.Object["runtimeDatabaseEncryptionKeyName"].(string); ok {
			r.RuntimeDatabaseEncryptionKeyName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.RuntimeDatabaseEncryptionKeyName: expected string")
		}
	}
	if _, ok := u.Object["runtimeType"]; ok {
		if s, ok := u.Object["runtimeType"].(string); ok {
			r.RuntimeType = dclService.OrganizationRuntimeTypeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.RuntimeType: expected string")
		}
	}
	if _, ok := u.Object["state"]; ok {
		if s, ok := u.Object["state"].(string); ok {
			r.State = dclService.OrganizationStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.State: expected string")
		}
	}
	if _, ok := u.Object["subscriptionType"]; ok {
		if s, ok := u.Object["subscriptionType"].(string); ok {
			r.SubscriptionType = dclService.OrganizationSubscriptionTypeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.SubscriptionType: expected string")
		}
	}
	return r, nil
}

func GetOrganization(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToOrganization(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetOrganization(ctx, r)
	if err != nil {
		return nil, err
	}
	return OrganizationToUnstructured(r), nil
}

func ListOrganization(ctx context.Context, config *dcl.Config) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListOrganization(ctx)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, OrganizationToUnstructured(r))
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

func ApplyOrganization(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToOrganization(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToOrganization(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyOrganization(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return OrganizationToUnstructured(r), nil
}

func OrganizationHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToOrganization(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToOrganization(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyOrganization(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteOrganization(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToOrganization(u)
	if err != nil {
		return err
	}
	return c.DeleteOrganization(ctx, r)
}

func OrganizationID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToOrganization(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Organization) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"apigee",
		"Organization",
		"beta",
	}
}

func (r *Organization) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Organization) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Organization) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Organization) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Organization) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Organization) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Organization) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetOrganization(ctx, config, resource)
}

func (r *Organization) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyOrganization(ctx, config, resource, opts...)
}

func (r *Organization) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return OrganizationHasDiff(ctx, config, resource, opts...)
}

func (r *Organization) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteOrganization(ctx, config, resource)
}

func (r *Organization) ID(resource *unstructured.Resource) (string, error) {
	return OrganizationID(resource)
}

func init() {
	unstructured.Register(&Organization{})
}
