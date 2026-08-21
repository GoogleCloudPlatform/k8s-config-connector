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

type AppleApp struct{}

func AppleAppToUnstructured(r *dclService.AppleApp) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "firebase",
			Version: "beta",
			Type:    "AppleApp",
		},
		Object: make(map[string]interface{}),
	}
	if r.ApiKeyId != nil {
		u.Object["apiKeyId"] = *r.ApiKeyId
	}
	if r.AppId != nil {
		u.Object["appId"] = *r.AppId
	}
	if r.AppStoreId != nil {
		u.Object["appStoreId"] = *r.AppStoreId
	}
	if r.BundleId != nil {
		u.Object["bundleId"] = *r.BundleId
	}
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
	if r.TeamId != nil {
		u.Object["teamId"] = *r.TeamId
	}
	return u
}

func UnstructuredToAppleApp(u *unstructured.Resource) (*dclService.AppleApp, error) {
	r := &dclService.AppleApp{}
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
	if _, ok := u.Object["appStoreId"]; ok {
		if s, ok := u.Object["appStoreId"].(string); ok {
			r.AppStoreId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.AppStoreId: expected string")
		}
	}
	if _, ok := u.Object["bundleId"]; ok {
		if s, ok := u.Object["bundleId"].(string); ok {
			r.BundleId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.BundleId: expected string")
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
	if _, ok := u.Object["teamId"]; ok {
		if s, ok := u.Object["teamId"].(string); ok {
			r.TeamId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.TeamId: expected string")
		}
	}
	return r, nil
}

func GetAppleApp(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAppleApp(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetAppleApp(ctx, r)
	if err != nil {
		return nil, err
	}
	return AppleAppToUnstructured(r), nil
}

func ListAppleApp(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListAppleApp(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, AppleAppToUnstructured(r))
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

func ApplyAppleApp(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAppleApp(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToAppleApp(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyAppleApp(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return AppleAppToUnstructured(r), nil
}

func AppleAppHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAppleApp(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToAppleApp(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyAppleApp(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteAppleApp(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAppleApp(u)
	if err != nil {
		return err
	}
	return c.DeleteAppleApp(ctx, r)
}

func AppleAppID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToAppleApp(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *AppleApp) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"firebase",
		"AppleApp",
		"beta",
	}
}

func (r *AppleApp) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *AppleApp) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *AppleApp) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *AppleApp) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *AppleApp) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *AppleApp) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *AppleApp) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetAppleApp(ctx, config, resource)
}

func (r *AppleApp) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyAppleApp(ctx, config, resource, opts...)
}

func (r *AppleApp) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return AppleAppHasDiff(ctx, config, resource, opts...)
}

func (r *AppleApp) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteAppleApp(ctx, config, resource)
}

func (r *AppleApp) ID(resource *unstructured.Resource) (string, error) {
	return AppleAppID(resource)
}

func init() {
	unstructured.Register(&AppleApp{})
}
