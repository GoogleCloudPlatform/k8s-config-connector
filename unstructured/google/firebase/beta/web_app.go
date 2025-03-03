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
package firebase

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/firebase/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type WebApp struct{}

func WebAppToUnstructured(r *dclService.WebApp) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "firebase",
			Version: "beta",
			Type:    "WebApp",
		},
		Object: make(map[string]interface{}),
	}
	if r.ApiKeyId != nil {
		u.Object["apiKeyId"] = *r.ApiKeyId
	}
	if r.AppId != nil {
		u.Object["appId"] = *r.AppId
	}
	var rAppUrls []interface{}
	for _, rAppUrlsVal := range r.AppUrls {
		rAppUrls = append(rAppUrls, rAppUrlsVal)
	}
	u.Object["appUrls"] = rAppUrls
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
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
	if r.WebId != nil {
		u.Object["webId"] = *r.WebId
	}
	return u
}

func UnstructuredToWebApp(u *unstructured.Resource) (*dclService.WebApp, error) {
	r := &dclService.WebApp{}
	if _, ok := u.Object["apiKeyId"]; ok {
		if s, ok := u.Object["apiKeyId"].(string); ok {
			r.ApiKeyId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ApiKeyId: expected string")
		}
	}
	if _, ok := u.Object["appId"]; ok {
		if s, ok := u.Object["appId"].(string); ok {
			r.AppId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.AppId: expected string")
		}
	}
	if _, ok := u.Object["appUrls"]; ok {
		if s, ok := u.Object["appUrls"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.AppUrls = append(r.AppUrls, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.AppUrls: expected []interface{}")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
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
	if _, ok := u.Object["webId"]; ok {
		if s, ok := u.Object["webId"].(string); ok {
			r.WebId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.WebId: expected string")
		}
	}
	return r, nil
}

func GetWebApp(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWebApp(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetWebApp(ctx, r)
	if err != nil {
		return nil, err
	}
	return WebAppToUnstructured(r), nil
}

func ListWebApp(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListWebApp(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, WebAppToUnstructured(r))
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

func ApplyWebApp(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWebApp(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToWebApp(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyWebApp(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return WebAppToUnstructured(r), nil
}

func WebAppHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWebApp(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToWebApp(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyWebApp(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteWebApp(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWebApp(u)
	if err != nil {
		return err
	}
	return c.DeleteWebApp(ctx, r)
}

func WebAppID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToWebApp(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *WebApp) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"firebase",
		"WebApp",
		"beta",
	}
}

func (r *WebApp) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *WebApp) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *WebApp) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *WebApp) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *WebApp) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *WebApp) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *WebApp) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetWebApp(ctx, config, resource)
}

func (r *WebApp) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyWebApp(ctx, config, resource, opts...)
}

func (r *WebApp) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return WebAppHasDiff(ctx, config, resource, opts...)
}

func (r *WebApp) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteWebApp(ctx, config, resource)
}

func (r *WebApp) ID(resource *unstructured.Resource) (string, error) {
	return WebAppID(resource)
}

func init() {
	unstructured.Register(&WebApp{})
}
