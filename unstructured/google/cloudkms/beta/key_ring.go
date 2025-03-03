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
package cloudkms

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudkms/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type KeyRing struct{}

func KeyRingToUnstructured(r *dclService.KeyRing) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "cloudkms",
			Version: "beta",
			Type:    "KeyRing",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	return u
}

func UnstructuredToKeyRing(u *unstructured.Resource) (*dclService.KeyRing, error) {
	r := &dclService.KeyRing{}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
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
	return r, nil
}

func GetKeyRing(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToKeyRing(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetKeyRing(ctx, r)
	if err != nil {
		return nil, err
	}
	return KeyRingToUnstructured(r), nil
}

func ListKeyRing(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListKeyRing(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, KeyRingToUnstructured(r))
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

func ApplyKeyRing(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToKeyRing(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToKeyRing(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyKeyRing(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return KeyRingToUnstructured(r), nil
}

func KeyRingHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToKeyRing(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToKeyRing(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyKeyRing(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteKeyRing(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func KeyRingID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToKeyRing(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *KeyRing) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"cloudkms",
		"KeyRing",
		"beta",
	}
}

func (r *KeyRing) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *KeyRing) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *KeyRing) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *KeyRing) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *KeyRing) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *KeyRing) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *KeyRing) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetKeyRing(ctx, config, resource)
}

func (r *KeyRing) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyKeyRing(ctx, config, resource, opts...)
}

func (r *KeyRing) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return KeyRingHasDiff(ctx, config, resource, opts...)
}

func (r *KeyRing) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteKeyRing(ctx, config, resource)
}

func (r *KeyRing) ID(resource *unstructured.Resource) (string, error) {
	return KeyRingID(resource)
}

func init() {
	unstructured.Register(&KeyRing{})
}
