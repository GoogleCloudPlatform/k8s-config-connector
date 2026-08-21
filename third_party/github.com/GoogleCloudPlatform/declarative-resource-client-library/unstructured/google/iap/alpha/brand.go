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
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iap/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Brand struct{}

func BrandToUnstructured(r *dclService.Brand) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "iap",
			Version: "alpha",
			Type:    "Brand",
		},
		Object: make(map[string]interface{}),
	}
	if r.ApplicationTitle != nil {
		u.Object["applicationTitle"] = *r.ApplicationTitle
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.OrgInternalOnly != nil {
		u.Object["orgInternalOnly"] = *r.OrgInternalOnly
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.SupportEmail != nil {
		u.Object["supportEmail"] = *r.SupportEmail
	}
	return u
}

func UnstructuredToBrand(u *unstructured.Resource) (*dclService.Brand, error) {
	r := &dclService.Brand{}
	if _, ok := u.Object["applicationTitle"]; ok {
		if s, ok := u.Object["applicationTitle"].(string); ok {
			r.ApplicationTitle = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ApplicationTitle: expected string")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["orgInternalOnly"]; ok {
		if b, ok := u.Object["orgInternalOnly"].(bool); ok {
			r.OrgInternalOnly = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.OrgInternalOnly: expected bool")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["supportEmail"]; ok {
		if s, ok := u.Object["supportEmail"].(string); ok {
			r.SupportEmail = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.SupportEmail: expected string")
		}
	}
	return r, nil
}

func GetBrand(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToBrand(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetBrand(ctx, r)
	if err != nil {
		return nil, err
	}
	return BrandToUnstructured(r), nil
}

func ListBrand(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListBrand(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, BrandToUnstructured(r))
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

func ApplyBrand(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToBrand(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToBrand(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyBrand(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return BrandToUnstructured(r), nil
}

func BrandHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToBrand(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToBrand(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyBrand(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteBrand(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func BrandID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToBrand(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Brand) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"iap",
		"Brand",
		"alpha",
	}
}

func (r *Brand) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Brand) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Brand) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Brand) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Brand) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Brand) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Brand) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetBrand(ctx, config, resource)
}

func (r *Brand) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyBrand(ctx, config, resource, opts...)
}

func (r *Brand) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return BrandHasDiff(ctx, config, resource, opts...)
}

func (r *Brand) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteBrand(ctx, config, resource)
}

func (r *Brand) ID(resource *unstructured.Resource) (string, error) {
	return BrandID(resource)
}

func init() {
	unstructured.Register(&Brand{})
}
