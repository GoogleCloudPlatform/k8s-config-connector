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
package apikeys

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apikeys"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Key struct{}

func KeyToUnstructured(r *dclService.Key) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "apikeys",
			Version: "ga",
			Type:    "Key",
		},
		Object: make(map[string]interface{}),
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.KeyString != nil {
		u.Object["keyString"] = *r.KeyString
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.Restrictions != nil && r.Restrictions != dclService.EmptyKeyRestrictions {
		rRestrictions := make(map[string]interface{})
		if r.Restrictions.AndroidKeyRestrictions != nil && r.Restrictions.AndroidKeyRestrictions != dclService.EmptyKeyRestrictionsAndroidKeyRestrictions {
			rRestrictionsAndroidKeyRestrictions := make(map[string]interface{})
			var rRestrictionsAndroidKeyRestrictionsAllowedApplications []interface{}
			for _, rRestrictionsAndroidKeyRestrictionsAllowedApplicationsVal := range r.Restrictions.AndroidKeyRestrictions.AllowedApplications {
				rRestrictionsAndroidKeyRestrictionsAllowedApplicationsObject := make(map[string]interface{})
				if rRestrictionsAndroidKeyRestrictionsAllowedApplicationsVal.PackageName != nil {
					rRestrictionsAndroidKeyRestrictionsAllowedApplicationsObject["packageName"] = *rRestrictionsAndroidKeyRestrictionsAllowedApplicationsVal.PackageName
				}
				if rRestrictionsAndroidKeyRestrictionsAllowedApplicationsVal.Sha1Fingerprint != nil {
					rRestrictionsAndroidKeyRestrictionsAllowedApplicationsObject["sha1Fingerprint"] = *rRestrictionsAndroidKeyRestrictionsAllowedApplicationsVal.Sha1Fingerprint
				}
				rRestrictionsAndroidKeyRestrictionsAllowedApplications = append(rRestrictionsAndroidKeyRestrictionsAllowedApplications, rRestrictionsAndroidKeyRestrictionsAllowedApplicationsObject)
			}
			rRestrictionsAndroidKeyRestrictions["allowedApplications"] = rRestrictionsAndroidKeyRestrictionsAllowedApplications
			rRestrictions["androidKeyRestrictions"] = rRestrictionsAndroidKeyRestrictions
		}
		var rRestrictionsApiTargets []interface{}
		for _, rRestrictionsApiTargetsVal := range r.Restrictions.ApiTargets {
			rRestrictionsApiTargetsObject := make(map[string]interface{})
			var rRestrictionsApiTargetsValMethods []interface{}
			for _, rRestrictionsApiTargetsValMethodsVal := range rRestrictionsApiTargetsVal.Methods {
				rRestrictionsApiTargetsValMethods = append(rRestrictionsApiTargetsValMethods, rRestrictionsApiTargetsValMethodsVal)
			}
			rRestrictionsApiTargetsObject["methods"] = rRestrictionsApiTargetsValMethods
			if rRestrictionsApiTargetsVal.Service != nil {
				rRestrictionsApiTargetsObject["service"] = *rRestrictionsApiTargetsVal.Service
			}
			rRestrictionsApiTargets = append(rRestrictionsApiTargets, rRestrictionsApiTargetsObject)
		}
		rRestrictions["apiTargets"] = rRestrictionsApiTargets
		if r.Restrictions.BrowserKeyRestrictions != nil && r.Restrictions.BrowserKeyRestrictions != dclService.EmptyKeyRestrictionsBrowserKeyRestrictions {
			rRestrictionsBrowserKeyRestrictions := make(map[string]interface{})
			var rRestrictionsBrowserKeyRestrictionsAllowedReferrers []interface{}
			for _, rRestrictionsBrowserKeyRestrictionsAllowedReferrersVal := range r.Restrictions.BrowserKeyRestrictions.AllowedReferrers {
				rRestrictionsBrowserKeyRestrictionsAllowedReferrers = append(rRestrictionsBrowserKeyRestrictionsAllowedReferrers, rRestrictionsBrowserKeyRestrictionsAllowedReferrersVal)
			}
			rRestrictionsBrowserKeyRestrictions["allowedReferrers"] = rRestrictionsBrowserKeyRestrictionsAllowedReferrers
			rRestrictions["browserKeyRestrictions"] = rRestrictionsBrowserKeyRestrictions
		}
		if r.Restrictions.IosKeyRestrictions != nil && r.Restrictions.IosKeyRestrictions != dclService.EmptyKeyRestrictionsIosKeyRestrictions {
			rRestrictionsIosKeyRestrictions := make(map[string]interface{})
			var rRestrictionsIosKeyRestrictionsAllowedBundleIds []interface{}
			for _, rRestrictionsIosKeyRestrictionsAllowedBundleIdsVal := range r.Restrictions.IosKeyRestrictions.AllowedBundleIds {
				rRestrictionsIosKeyRestrictionsAllowedBundleIds = append(rRestrictionsIosKeyRestrictionsAllowedBundleIds, rRestrictionsIosKeyRestrictionsAllowedBundleIdsVal)
			}
			rRestrictionsIosKeyRestrictions["allowedBundleIds"] = rRestrictionsIosKeyRestrictionsAllowedBundleIds
			rRestrictions["iosKeyRestrictions"] = rRestrictionsIosKeyRestrictions
		}
		if r.Restrictions.ServerKeyRestrictions != nil && r.Restrictions.ServerKeyRestrictions != dclService.EmptyKeyRestrictionsServerKeyRestrictions {
			rRestrictionsServerKeyRestrictions := make(map[string]interface{})
			var rRestrictionsServerKeyRestrictionsAllowedIps []interface{}
			for _, rRestrictionsServerKeyRestrictionsAllowedIpsVal := range r.Restrictions.ServerKeyRestrictions.AllowedIps {
				rRestrictionsServerKeyRestrictionsAllowedIps = append(rRestrictionsServerKeyRestrictionsAllowedIps, rRestrictionsServerKeyRestrictionsAllowedIpsVal)
			}
			rRestrictionsServerKeyRestrictions["allowedIps"] = rRestrictionsServerKeyRestrictionsAllowedIps
			rRestrictions["serverKeyRestrictions"] = rRestrictionsServerKeyRestrictions
		}
		u.Object["restrictions"] = rRestrictions
	}
	if r.Uid != nil {
		u.Object["uid"] = *r.Uid
	}
	return u
}

func UnstructuredToKey(u *unstructured.Resource) (*dclService.Key, error) {
	r := &dclService.Key{}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
		}
	}
	if _, ok := u.Object["keyString"]; ok {
		if s, ok := u.Object["keyString"].(string); ok {
			r.KeyString = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.KeyString: expected string")
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
	if _, ok := u.Object["restrictions"]; ok {
		if rRestrictions, ok := u.Object["restrictions"].(map[string]interface{}); ok {
			r.Restrictions = &dclService.KeyRestrictions{}
			if _, ok := rRestrictions["androidKeyRestrictions"]; ok {
				if rRestrictionsAndroidKeyRestrictions, ok := rRestrictions["androidKeyRestrictions"].(map[string]interface{}); ok {
					r.Restrictions.AndroidKeyRestrictions = &dclService.KeyRestrictionsAndroidKeyRestrictions{}
					if _, ok := rRestrictionsAndroidKeyRestrictions["allowedApplications"]; ok {
						if s, ok := rRestrictionsAndroidKeyRestrictions["allowedApplications"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rRestrictionsAndroidKeyRestrictionsAllowedApplications dclService.KeyRestrictionsAndroidKeyRestrictionsAllowedApplications
									if _, ok := objval["packageName"]; ok {
										if s, ok := objval["packageName"].(string); ok {
											rRestrictionsAndroidKeyRestrictionsAllowedApplications.PackageName = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rRestrictionsAndroidKeyRestrictionsAllowedApplications.PackageName: expected string")
										}
									}
									if _, ok := objval["sha1Fingerprint"]; ok {
										if s, ok := objval["sha1Fingerprint"].(string); ok {
											rRestrictionsAndroidKeyRestrictionsAllowedApplications.Sha1Fingerprint = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rRestrictionsAndroidKeyRestrictionsAllowedApplications.Sha1Fingerprint: expected string")
										}
									}
									r.Restrictions.AndroidKeyRestrictions.AllowedApplications = append(r.Restrictions.AndroidKeyRestrictions.AllowedApplications, rRestrictionsAndroidKeyRestrictionsAllowedApplications)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Restrictions.AndroidKeyRestrictions.AllowedApplications: expected []interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Restrictions.AndroidKeyRestrictions: expected map[string]interface{}")
				}
			}
			if _, ok := rRestrictions["apiTargets"]; ok {
				if s, ok := rRestrictions["apiTargets"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rRestrictionsApiTargets dclService.KeyRestrictionsApiTargets
							if _, ok := objval["methods"]; ok {
								if s, ok := objval["methods"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rRestrictionsApiTargets.Methods = append(rRestrictionsApiTargets.Methods, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rRestrictionsApiTargets.Methods: expected []interface{}")
								}
							}
							if _, ok := objval["service"]; ok {
								if s, ok := objval["service"].(string); ok {
									rRestrictionsApiTargets.Service = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rRestrictionsApiTargets.Service: expected string")
								}
							}
							r.Restrictions.ApiTargets = append(r.Restrictions.ApiTargets, rRestrictionsApiTargets)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Restrictions.ApiTargets: expected []interface{}")
				}
			}
			if _, ok := rRestrictions["browserKeyRestrictions"]; ok {
				if rRestrictionsBrowserKeyRestrictions, ok := rRestrictions["browserKeyRestrictions"].(map[string]interface{}); ok {
					r.Restrictions.BrowserKeyRestrictions = &dclService.KeyRestrictionsBrowserKeyRestrictions{}
					if _, ok := rRestrictionsBrowserKeyRestrictions["allowedReferrers"]; ok {
						if s, ok := rRestrictionsBrowserKeyRestrictions["allowedReferrers"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.Restrictions.BrowserKeyRestrictions.AllowedReferrers = append(r.Restrictions.BrowserKeyRestrictions.AllowedReferrers, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Restrictions.BrowserKeyRestrictions.AllowedReferrers: expected []interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Restrictions.BrowserKeyRestrictions: expected map[string]interface{}")
				}
			}
			if _, ok := rRestrictions["iosKeyRestrictions"]; ok {
				if rRestrictionsIosKeyRestrictions, ok := rRestrictions["iosKeyRestrictions"].(map[string]interface{}); ok {
					r.Restrictions.IosKeyRestrictions = &dclService.KeyRestrictionsIosKeyRestrictions{}
					if _, ok := rRestrictionsIosKeyRestrictions["allowedBundleIds"]; ok {
						if s, ok := rRestrictionsIosKeyRestrictions["allowedBundleIds"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.Restrictions.IosKeyRestrictions.AllowedBundleIds = append(r.Restrictions.IosKeyRestrictions.AllowedBundleIds, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Restrictions.IosKeyRestrictions.AllowedBundleIds: expected []interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Restrictions.IosKeyRestrictions: expected map[string]interface{}")
				}
			}
			if _, ok := rRestrictions["serverKeyRestrictions"]; ok {
				if rRestrictionsServerKeyRestrictions, ok := rRestrictions["serverKeyRestrictions"].(map[string]interface{}); ok {
					r.Restrictions.ServerKeyRestrictions = &dclService.KeyRestrictionsServerKeyRestrictions{}
					if _, ok := rRestrictionsServerKeyRestrictions["allowedIps"]; ok {
						if s, ok := rRestrictionsServerKeyRestrictions["allowedIps"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.Restrictions.ServerKeyRestrictions.AllowedIps = append(r.Restrictions.ServerKeyRestrictions.AllowedIps, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Restrictions.ServerKeyRestrictions.AllowedIps: expected []interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Restrictions.ServerKeyRestrictions: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Restrictions: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["uid"]; ok {
		if s, ok := u.Object["uid"].(string); ok {
			r.Uid = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Uid: expected string")
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
		"apikeys",
		"Key",
		"ga",
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
