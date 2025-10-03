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
package recaptchaenterprise

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/recaptchaenterprise/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Key struct{}

func KeyToUnstructured(r *dclService.Key) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "recaptchaenterprise",
			Version: "beta",
			Type:    "Key",
		},
		Object: make(map[string]interface{}),
	}
	if r.AndroidSettings != nil && r.AndroidSettings != dclService.EmptyKeyAndroidSettings {
		rAndroidSettings := make(map[string]interface{})
		if r.AndroidSettings.AllowAllPackageNames != nil {
			rAndroidSettings["allowAllPackageNames"] = *r.AndroidSettings.AllowAllPackageNames
		}
		var rAndroidSettingsAllowedPackageNames []interface{}
		for _, rAndroidSettingsAllowedPackageNamesVal := range r.AndroidSettings.AllowedPackageNames {
			rAndroidSettingsAllowedPackageNames = append(rAndroidSettingsAllowedPackageNames, rAndroidSettingsAllowedPackageNamesVal)
		}
		rAndroidSettings["allowedPackageNames"] = rAndroidSettingsAllowedPackageNames
		u.Object["androidSettings"] = rAndroidSettings
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.IosSettings != nil && r.IosSettings != dclService.EmptyKeyIosSettings {
		rIosSettings := make(map[string]interface{})
		if r.IosSettings.AllowAllBundleIds != nil {
			rIosSettings["allowAllBundleIds"] = *r.IosSettings.AllowAllBundleIds
		}
		var rIosSettingsAllowedBundleIds []interface{}
		for _, rIosSettingsAllowedBundleIdsVal := range r.IosSettings.AllowedBundleIds {
			rIosSettingsAllowedBundleIds = append(rIosSettingsAllowedBundleIds, rIosSettingsAllowedBundleIdsVal)
		}
		rIosSettings["allowedBundleIds"] = rIosSettingsAllowedBundleIds
		u.Object["iosSettings"] = rIosSettings
	}
	if r.Labels != nil {
		rLabels := make(map[string]interface{})
		for k, v := range r.Labels {
			rLabels[k] = v
		}
		u.Object["labels"] = rLabels
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.TestingOptions != nil && r.TestingOptions != dclService.EmptyKeyTestingOptions {
		rTestingOptions := make(map[string]interface{})
		if r.TestingOptions.TestingChallenge != nil {
			rTestingOptions["testingChallenge"] = string(*r.TestingOptions.TestingChallenge)
		}
		if r.TestingOptions.TestingScore != nil {
			rTestingOptions["testingScore"] = *r.TestingOptions.TestingScore
		}
		u.Object["testingOptions"] = rTestingOptions
	}
	if r.WafSettings != nil && r.WafSettings != dclService.EmptyKeyWafSettings {
		rWafSettings := make(map[string]interface{})
		if r.WafSettings.WafFeature != nil {
			rWafSettings["wafFeature"] = string(*r.WafSettings.WafFeature)
		}
		if r.WafSettings.WafService != nil {
			rWafSettings["wafService"] = string(*r.WafSettings.WafService)
		}
		u.Object["wafSettings"] = rWafSettings
	}
	if r.WebSettings != nil && r.WebSettings != dclService.EmptyKeyWebSettings {
		rWebSettings := make(map[string]interface{})
		if r.WebSettings.AllowAllDomains != nil {
			rWebSettings["allowAllDomains"] = *r.WebSettings.AllowAllDomains
		}
		if r.WebSettings.AllowAmpTraffic != nil {
			rWebSettings["allowAmpTraffic"] = *r.WebSettings.AllowAmpTraffic
		}
		var rWebSettingsAllowedDomains []interface{}
		for _, rWebSettingsAllowedDomainsVal := range r.WebSettings.AllowedDomains {
			rWebSettingsAllowedDomains = append(rWebSettingsAllowedDomains, rWebSettingsAllowedDomainsVal)
		}
		rWebSettings["allowedDomains"] = rWebSettingsAllowedDomains
		if r.WebSettings.ChallengeSecurityPreference != nil {
			rWebSettings["challengeSecurityPreference"] = string(*r.WebSettings.ChallengeSecurityPreference)
		}
		if r.WebSettings.IntegrationType != nil {
			rWebSettings["integrationType"] = string(*r.WebSettings.IntegrationType)
		}
		u.Object["webSettings"] = rWebSettings
	}
	return u
}

func UnstructuredToKey(u *unstructured.Resource) (*dclService.Key, error) {
	r := &dclService.Key{}
	if _, ok := u.Object["androidSettings"]; ok {
		if rAndroidSettings, ok := u.Object["androidSettings"].(map[string]interface{}); ok {
			r.AndroidSettings = &dclService.KeyAndroidSettings{}
			if _, ok := rAndroidSettings["allowAllPackageNames"]; ok {
				if b, ok := rAndroidSettings["allowAllPackageNames"].(bool); ok {
					r.AndroidSettings.AllowAllPackageNames = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.AndroidSettings.AllowAllPackageNames: expected bool")
				}
			}
			if _, ok := rAndroidSettings["allowedPackageNames"]; ok {
				if s, ok := rAndroidSettings["allowedPackageNames"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.AndroidSettings.AllowedPackageNames = append(r.AndroidSettings.AllowedPackageNames, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.AndroidSettings.AllowedPackageNames: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.AndroidSettings: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
		}
	}
	if _, ok := u.Object["iosSettings"]; ok {
		if rIosSettings, ok := u.Object["iosSettings"].(map[string]interface{}); ok {
			r.IosSettings = &dclService.KeyIosSettings{}
			if _, ok := rIosSettings["allowAllBundleIds"]; ok {
				if b, ok := rIosSettings["allowAllBundleIds"].(bool); ok {
					r.IosSettings.AllowAllBundleIds = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.IosSettings.AllowAllBundleIds: expected bool")
				}
			}
			if _, ok := rIosSettings["allowedBundleIds"]; ok {
				if s, ok := rIosSettings["allowedBundleIds"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.IosSettings.AllowedBundleIds = append(r.IosSettings.AllowedBundleIds, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.IosSettings.AllowedBundleIds: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.IosSettings: expected map[string]interface{}")
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
	if _, ok := u.Object["testingOptions"]; ok {
		if rTestingOptions, ok := u.Object["testingOptions"].(map[string]interface{}); ok {
			r.TestingOptions = &dclService.KeyTestingOptions{}
			if _, ok := rTestingOptions["testingChallenge"]; ok {
				if s, ok := rTestingOptions["testingChallenge"].(string); ok {
					r.TestingOptions.TestingChallenge = dclService.KeyTestingOptionsTestingChallengeEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.TestingOptions.TestingChallenge: expected string")
				}
			}
			if _, ok := rTestingOptions["testingScore"]; ok {
				if f, ok := rTestingOptions["testingScore"].(float64); ok {
					r.TestingOptions.TestingScore = dcl.Float64(f)
				} else {
					return nil, fmt.Errorf("r.TestingOptions.TestingScore: expected float64")
				}
			}
		} else {
			return nil, fmt.Errorf("r.TestingOptions: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["wafSettings"]; ok {
		if rWafSettings, ok := u.Object["wafSettings"].(map[string]interface{}); ok {
			r.WafSettings = &dclService.KeyWafSettings{}
			if _, ok := rWafSettings["wafFeature"]; ok {
				if s, ok := rWafSettings["wafFeature"].(string); ok {
					r.WafSettings.WafFeature = dclService.KeyWafSettingsWafFeatureEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.WafSettings.WafFeature: expected string")
				}
			}
			if _, ok := rWafSettings["wafService"]; ok {
				if s, ok := rWafSettings["wafService"].(string); ok {
					r.WafSettings.WafService = dclService.KeyWafSettingsWafServiceEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.WafSettings.WafService: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.WafSettings: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["webSettings"]; ok {
		if rWebSettings, ok := u.Object["webSettings"].(map[string]interface{}); ok {
			r.WebSettings = &dclService.KeyWebSettings{}
			if _, ok := rWebSettings["allowAllDomains"]; ok {
				if b, ok := rWebSettings["allowAllDomains"].(bool); ok {
					r.WebSettings.AllowAllDomains = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.WebSettings.AllowAllDomains: expected bool")
				}
			}
			if _, ok := rWebSettings["allowAmpTraffic"]; ok {
				if b, ok := rWebSettings["allowAmpTraffic"].(bool); ok {
					r.WebSettings.AllowAmpTraffic = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.WebSettings.AllowAmpTraffic: expected bool")
				}
			}
			if _, ok := rWebSettings["allowedDomains"]; ok {
				if s, ok := rWebSettings["allowedDomains"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.WebSettings.AllowedDomains = append(r.WebSettings.AllowedDomains, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.WebSettings.AllowedDomains: expected []interface{}")
				}
			}
			if _, ok := rWebSettings["challengeSecurityPreference"]; ok {
				if s, ok := rWebSettings["challengeSecurityPreference"].(string); ok {
					r.WebSettings.ChallengeSecurityPreference = dclService.KeyWebSettingsChallengeSecurityPreferenceEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.WebSettings.ChallengeSecurityPreference: expected string")
				}
			}
			if _, ok := rWebSettings["integrationType"]; ok {
				if s, ok := rWebSettings["integrationType"].(string); ok {
					r.WebSettings.IntegrationType = dclService.KeyWebSettingsIntegrationTypeEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.WebSettings.IntegrationType: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.WebSettings: expected map[string]interface{}")
		}
	}
	return r, nil
}

func GetKey(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToKey(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetKey(ctx, r)
	if err != nil {
		return nil, err
	}
	return KeyToUnstructured(r), nil
}

func ListKey(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListKey(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, KeyToUnstructured(r))
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

func ApplyKey(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToKey(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToKey(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyKey(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return KeyToUnstructured(r), nil
}

func KeyHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToKey(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToKey(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyKey(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteKey(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToKey(u)
	if err != nil {
		return err
	}
	return c.DeleteKey(ctx, r)
}

func KeyID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToKey(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Key) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"recaptchaenterprise",
		"Key",
		"beta",
	}
}

func (r *Key) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Key) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Key) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Key) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Key) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Key) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Key) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetKey(ctx, config, resource)
}

func (r *Key) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyKey(ctx, config, resource, opts...)
}

func (r *Key) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return KeyHasDiff(ctx, config, resource, opts...)
}

func (r *Key) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteKey(ctx, config, resource)
}

func (r *Key) ID(resource *unstructured.Resource) (string, error) {
	return KeyID(resource)
}

func init() {
	unstructured.Register(&Key{})
}
