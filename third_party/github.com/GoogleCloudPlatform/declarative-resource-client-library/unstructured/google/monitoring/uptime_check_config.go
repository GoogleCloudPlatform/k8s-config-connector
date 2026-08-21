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
package monitoring

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type UptimeCheckConfig struct{}

func UptimeCheckConfigToUnstructured(r *dclService.UptimeCheckConfig) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "monitoring",
			Version: "ga",
			Type:    "UptimeCheckConfig",
		},
		Object: make(map[string]interface{}),
	}
	var rContentMatchers []interface{}
	for _, rContentMatchersVal := range r.ContentMatchers {
		rContentMatchersObject := make(map[string]interface{})
		if rContentMatchersVal.Content != nil {
			rContentMatchersObject["content"] = *rContentMatchersVal.Content
		}
		if rContentMatchersVal.Matcher != nil {
			rContentMatchersObject["matcher"] = string(*rContentMatchersVal.Matcher)
		}
		rContentMatchers = append(rContentMatchers, rContentMatchersObject)
	}
	u.Object["contentMatchers"] = rContentMatchers
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.HttpCheck != nil && r.HttpCheck != dclService.EmptyUptimeCheckConfigHttpCheck {
		rHttpCheck := make(map[string]interface{})
		if r.HttpCheck.AuthInfo != nil && r.HttpCheck.AuthInfo != dclService.EmptyUptimeCheckConfigHttpCheckAuthInfo {
			rHttpCheckAuthInfo := make(map[string]interface{})
			if r.HttpCheck.AuthInfo.Password != nil {
				rHttpCheckAuthInfo["password"] = *r.HttpCheck.AuthInfo.Password
			}
			if r.HttpCheck.AuthInfo.Username != nil {
				rHttpCheckAuthInfo["username"] = *r.HttpCheck.AuthInfo.Username
			}
			rHttpCheck["authInfo"] = rHttpCheckAuthInfo
		}
		if r.HttpCheck.Body != nil {
			rHttpCheck["body"] = *r.HttpCheck.Body
		}
		if r.HttpCheck.ContentType != nil {
			rHttpCheck["contentType"] = string(*r.HttpCheck.ContentType)
		}
		if r.HttpCheck.Headers != nil {
			rHttpCheckHeaders := make(map[string]interface{})
			for k, v := range r.HttpCheck.Headers {
				rHttpCheckHeaders[k] = v
			}
			rHttpCheck["headers"] = rHttpCheckHeaders
		}
		if r.HttpCheck.MaskHeaders != nil {
			rHttpCheck["maskHeaders"] = *r.HttpCheck.MaskHeaders
		}
		if r.HttpCheck.Path != nil {
			rHttpCheck["path"] = *r.HttpCheck.Path
		}
		if r.HttpCheck.Port != nil {
			rHttpCheck["port"] = *r.HttpCheck.Port
		}
		if r.HttpCheck.RequestMethod != nil {
			rHttpCheck["requestMethod"] = string(*r.HttpCheck.RequestMethod)
		}
		if r.HttpCheck.UseSsl != nil {
			rHttpCheck["useSsl"] = *r.HttpCheck.UseSsl
		}
		if r.HttpCheck.ValidateSsl != nil {
			rHttpCheck["validateSsl"] = *r.HttpCheck.ValidateSsl
		}
		u.Object["httpCheck"] = rHttpCheck
	}
	if r.MonitoredResource != nil && r.MonitoredResource != dclService.EmptyUptimeCheckConfigMonitoredResource {
		rMonitoredResource := make(map[string]interface{})
		if r.MonitoredResource.FilterLabels != nil {
			rMonitoredResourceFilterLabels := make(map[string]interface{})
			for k, v := range r.MonitoredResource.FilterLabels {
				rMonitoredResourceFilterLabels[k] = v
			}
			rMonitoredResource["filterLabels"] = rMonitoredResourceFilterLabels
		}
		if r.MonitoredResource.Type != nil {
			rMonitoredResource["type"] = *r.MonitoredResource.Type
		}
		u.Object["monitoredResource"] = rMonitoredResource
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Period != nil {
		u.Object["period"] = *r.Period
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.ResourceGroup != nil && r.ResourceGroup != dclService.EmptyUptimeCheckConfigResourceGroup {
		rResourceGroup := make(map[string]interface{})
		if r.ResourceGroup.GroupId != nil {
			rResourceGroup["groupId"] = *r.ResourceGroup.GroupId
		}
		if r.ResourceGroup.ResourceType != nil {
			rResourceGroup["resourceType"] = string(*r.ResourceGroup.ResourceType)
		}
		u.Object["resourceGroup"] = rResourceGroup
	}
	var rSelectedRegions []interface{}
	for _, rSelectedRegionsVal := range r.SelectedRegions {
		rSelectedRegions = append(rSelectedRegions, rSelectedRegionsVal)
	}
	u.Object["selectedRegions"] = rSelectedRegions
	if r.TcpCheck != nil && r.TcpCheck != dclService.EmptyUptimeCheckConfigTcpCheck {
		rTcpCheck := make(map[string]interface{})
		if r.TcpCheck.Port != nil {
			rTcpCheck["port"] = *r.TcpCheck.Port
		}
		u.Object["tcpCheck"] = rTcpCheck
	}
	if r.Timeout != nil {
		u.Object["timeout"] = *r.Timeout
	}
	return u
}

func UnstructuredToUptimeCheckConfig(u *unstructured.Resource) (*dclService.UptimeCheckConfig, error) {
	r := &dclService.UptimeCheckConfig{}
	if _, ok := u.Object["contentMatchers"]; ok {
		if s, ok := u.Object["contentMatchers"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rContentMatchers dclService.UptimeCheckConfigContentMatchers
					if _, ok := objval["content"]; ok {
						if s, ok := objval["content"].(string); ok {
							rContentMatchers.Content = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rContentMatchers.Content: expected string")
						}
					}
					if _, ok := objval["matcher"]; ok {
						if s, ok := objval["matcher"].(string); ok {
							rContentMatchers.Matcher = dclService.UptimeCheckConfigContentMatchersMatcherEnumRef(s)
						} else {
							return nil, fmt.Errorf("rContentMatchers.Matcher: expected string")
						}
					}
					r.ContentMatchers = append(r.ContentMatchers, rContentMatchers)
				}
			}
		} else {
			return nil, fmt.Errorf("r.ContentMatchers: expected []interface{}")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
		}
	}
	if _, ok := u.Object["httpCheck"]; ok {
		if rHttpCheck, ok := u.Object["httpCheck"].(map[string]interface{}); ok {
			r.HttpCheck = &dclService.UptimeCheckConfigHttpCheck{}
			if _, ok := rHttpCheck["authInfo"]; ok {
				if rHttpCheckAuthInfo, ok := rHttpCheck["authInfo"].(map[string]interface{}); ok {
					r.HttpCheck.AuthInfo = &dclService.UptimeCheckConfigHttpCheckAuthInfo{}
					if _, ok := rHttpCheckAuthInfo["password"]; ok {
						if s, ok := rHttpCheckAuthInfo["password"].(string); ok {
							r.HttpCheck.AuthInfo.Password = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.HttpCheck.AuthInfo.Password: expected string")
						}
					}
					if _, ok := rHttpCheckAuthInfo["username"]; ok {
						if s, ok := rHttpCheckAuthInfo["username"].(string); ok {
							r.HttpCheck.AuthInfo.Username = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.HttpCheck.AuthInfo.Username: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.HttpCheck.AuthInfo: expected map[string]interface{}")
				}
			}
			if _, ok := rHttpCheck["body"]; ok {
				if s, ok := rHttpCheck["body"].(string); ok {
					r.HttpCheck.Body = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.HttpCheck.Body: expected string")
				}
			}
			if _, ok := rHttpCheck["contentType"]; ok {
				if s, ok := rHttpCheck["contentType"].(string); ok {
					r.HttpCheck.ContentType = dclService.UptimeCheckConfigHttpCheckContentTypeEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.HttpCheck.ContentType: expected string")
				}
			}
			if _, ok := rHttpCheck["headers"]; ok {
				if rHttpCheckHeaders, ok := rHttpCheck["headers"].(map[string]interface{}); ok {
					m := make(map[string]string)
					for k, v := range rHttpCheckHeaders {
						if s, ok := v.(string); ok {
							m[k] = s
						}
					}
					r.HttpCheck.Headers = m
				} else {
					return nil, fmt.Errorf("r.HttpCheck.Headers: expected map[string]interface{}")
				}
			}
			if _, ok := rHttpCheck["maskHeaders"]; ok {
				if b, ok := rHttpCheck["maskHeaders"].(bool); ok {
					r.HttpCheck.MaskHeaders = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.HttpCheck.MaskHeaders: expected bool")
				}
			}
			if _, ok := rHttpCheck["path"]; ok {
				if s, ok := rHttpCheck["path"].(string); ok {
					r.HttpCheck.Path = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.HttpCheck.Path: expected string")
				}
			}
			if _, ok := rHttpCheck["port"]; ok {
				if i, ok := rHttpCheck["port"].(int64); ok {
					r.HttpCheck.Port = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.HttpCheck.Port: expected int64")
				}
			}
			if _, ok := rHttpCheck["requestMethod"]; ok {
				if s, ok := rHttpCheck["requestMethod"].(string); ok {
					r.HttpCheck.RequestMethod = dclService.UptimeCheckConfigHttpCheckRequestMethodEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.HttpCheck.RequestMethod: expected string")
				}
			}
			if _, ok := rHttpCheck["useSsl"]; ok {
				if b, ok := rHttpCheck["useSsl"].(bool); ok {
					r.HttpCheck.UseSsl = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.HttpCheck.UseSsl: expected bool")
				}
			}
			if _, ok := rHttpCheck["validateSsl"]; ok {
				if b, ok := rHttpCheck["validateSsl"].(bool); ok {
					r.HttpCheck.ValidateSsl = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.HttpCheck.ValidateSsl: expected bool")
				}
			}
		} else {
			return nil, fmt.Errorf("r.HttpCheck: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["monitoredResource"]; ok {
		if rMonitoredResource, ok := u.Object["monitoredResource"].(map[string]interface{}); ok {
			r.MonitoredResource = &dclService.UptimeCheckConfigMonitoredResource{}
			if _, ok := rMonitoredResource["filterLabels"]; ok {
				if rMonitoredResourceFilterLabels, ok := rMonitoredResource["filterLabels"].(map[string]interface{}); ok {
					m := make(map[string]string)
					for k, v := range rMonitoredResourceFilterLabels {
						if s, ok := v.(string); ok {
							m[k] = s
						}
					}
					r.MonitoredResource.FilterLabels = m
				} else {
					return nil, fmt.Errorf("r.MonitoredResource.FilterLabels: expected map[string]interface{}")
				}
			}
			if _, ok := rMonitoredResource["type"]; ok {
				if s, ok := rMonitoredResource["type"].(string); ok {
					r.MonitoredResource.Type = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.MonitoredResource.Type: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.MonitoredResource: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["period"]; ok {
		if s, ok := u.Object["period"].(string); ok {
			r.Period = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Period: expected string")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["resourceGroup"]; ok {
		if rResourceGroup, ok := u.Object["resourceGroup"].(map[string]interface{}); ok {
			r.ResourceGroup = &dclService.UptimeCheckConfigResourceGroup{}
			if _, ok := rResourceGroup["groupId"]; ok {
				if s, ok := rResourceGroup["groupId"].(string); ok {
					r.ResourceGroup.GroupId = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.ResourceGroup.GroupId: expected string")
				}
			}
			if _, ok := rResourceGroup["resourceType"]; ok {
				if s, ok := rResourceGroup["resourceType"].(string); ok {
					r.ResourceGroup.ResourceType = dclService.UptimeCheckConfigResourceGroupResourceTypeEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.ResourceGroup.ResourceType: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.ResourceGroup: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["selectedRegions"]; ok {
		if s, ok := u.Object["selectedRegions"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.SelectedRegions = append(r.SelectedRegions, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.SelectedRegions: expected []interface{}")
		}
	}
	if _, ok := u.Object["tcpCheck"]; ok {
		if rTcpCheck, ok := u.Object["tcpCheck"].(map[string]interface{}); ok {
			r.TcpCheck = &dclService.UptimeCheckConfigTcpCheck{}
			if _, ok := rTcpCheck["port"]; ok {
				if i, ok := rTcpCheck["port"].(int64); ok {
					r.TcpCheck.Port = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.TcpCheck.Port: expected int64")
				}
			}
		} else {
			return nil, fmt.Errorf("r.TcpCheck: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["timeout"]; ok {
		if s, ok := u.Object["timeout"].(string); ok {
			r.Timeout = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Timeout: expected string")
		}
	}
	return r, nil
}

func GetUptimeCheckConfig(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToUptimeCheckConfig(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetUptimeCheckConfig(ctx, r)
	if err != nil {
		return nil, err
	}
	return UptimeCheckConfigToUnstructured(r), nil
}

func ListUptimeCheckConfig(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListUptimeCheckConfig(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, UptimeCheckConfigToUnstructured(r))
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

func ApplyUptimeCheckConfig(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToUptimeCheckConfig(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToUptimeCheckConfig(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyUptimeCheckConfig(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return UptimeCheckConfigToUnstructured(r), nil
}

func UptimeCheckConfigHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToUptimeCheckConfig(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToUptimeCheckConfig(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyUptimeCheckConfig(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteUptimeCheckConfig(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToUptimeCheckConfig(u)
	if err != nil {
		return err
	}
	return c.DeleteUptimeCheckConfig(ctx, r)
}

func UptimeCheckConfigID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToUptimeCheckConfig(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *UptimeCheckConfig) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"monitoring",
		"UptimeCheckConfig",
		"ga",
	}
}

func (r *UptimeCheckConfig) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *UptimeCheckConfig) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *UptimeCheckConfig) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *UptimeCheckConfig) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *UptimeCheckConfig) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *UptimeCheckConfig) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *UptimeCheckConfig) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetUptimeCheckConfig(ctx, config, resource)
}

func (r *UptimeCheckConfig) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyUptimeCheckConfig(ctx, config, resource, opts...)
}

func (r *UptimeCheckConfig) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return UptimeCheckConfigHasDiff(ctx, config, resource, opts...)
}

func (r *UptimeCheckConfig) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteUptimeCheckConfig(ctx, config, resource)
}

func (r *UptimeCheckConfig) ID(resource *unstructured.Resource) (string, error) {
	return UptimeCheckConfigID(resource)
}

func init() {
	unstructured.Register(&UptimeCheckConfig{})
}
