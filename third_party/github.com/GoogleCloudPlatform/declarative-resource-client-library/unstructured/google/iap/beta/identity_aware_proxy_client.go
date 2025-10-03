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
package iap

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iap/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type IdentityAwareProxyClient struct{}

func IdentityAwareProxyClientToUnstructured(r *dclService.IdentityAwareProxyClient) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "iap",
			Version: "beta",
			Type:    "IdentityAwareProxyClient",
		},
		Object: make(map[string]interface{}),
	}
	if r.Brand != nil {
		u.Object["brand"] = *r.Brand
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
	if r.Secret != nil {
		u.Object["secret"] = *r.Secret
	}
	return u
}

func UnstructuredToIdentityAwareProxyClient(u *unstructured.Resource) (*dclService.IdentityAwareProxyClient, error) {
	r := &dclService.IdentityAwareProxyClient{}
	if _, ok := u.Object["brand"]; ok {
		if s, ok := u.Object["brand"].(string); ok {
			r.Brand = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Brand: expected string")
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
	if _, ok := u.Object["secret"]; ok {
		if s, ok := u.Object["secret"].(string); ok {
			r.Secret = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Secret: expected string")
		}
	}
	return r, nil
}

func GetIdentityAwareProxyClient(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToIdentityAwareProxyClient(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetIdentityAwareProxyClient(ctx, r)
	if err != nil {
		return nil, err
	}
	return IdentityAwareProxyClientToUnstructured(r), nil
}

func ListIdentityAwareProxyClient(ctx context.Context, config *dcl.Config, project string, brand string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListIdentityAwareProxyClient(ctx, project, brand)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, IdentityAwareProxyClientToUnstructured(r))
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

func ApplyIdentityAwareProxyClient(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToIdentityAwareProxyClient(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToIdentityAwareProxyClient(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyIdentityAwareProxyClient(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return IdentityAwareProxyClientToUnstructured(r), nil
}

func IdentityAwareProxyClientHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToIdentityAwareProxyClient(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToIdentityAwareProxyClient(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyIdentityAwareProxyClient(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteIdentityAwareProxyClient(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToIdentityAwareProxyClient(u)
	if err != nil {
		return err
	}
	return c.DeleteIdentityAwareProxyClient(ctx, r)
}

func IdentityAwareProxyClientID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToIdentityAwareProxyClient(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *IdentityAwareProxyClient) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"iap",
		"IdentityAwareProxyClient",
		"beta",
	}
}

func (r *IdentityAwareProxyClient) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *IdentityAwareProxyClient) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *IdentityAwareProxyClient) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *IdentityAwareProxyClient) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *IdentityAwareProxyClient) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *IdentityAwareProxyClient) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *IdentityAwareProxyClient) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetIdentityAwareProxyClient(ctx, config, resource)
}

func (r *IdentityAwareProxyClient) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyIdentityAwareProxyClient(ctx, config, resource, opts...)
}

func (r *IdentityAwareProxyClient) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return IdentityAwareProxyClientHasDiff(ctx, config, resource, opts...)
}

func (r *IdentityAwareProxyClient) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteIdentityAwareProxyClient(ctx, config, resource)
}

func (r *IdentityAwareProxyClient) ID(resource *unstructured.Resource) (string, error) {
	return IdentityAwareProxyClientID(resource)
}

func init() {
	unstructured.Register(&IdentityAwareProxyClient{})
}
