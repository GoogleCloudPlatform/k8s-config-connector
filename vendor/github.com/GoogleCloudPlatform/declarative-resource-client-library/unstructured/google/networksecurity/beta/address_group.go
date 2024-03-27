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
package networksecurity

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networksecurity/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type AddressGroup struct{}

func AddressGroupToUnstructured(r *dclService.AddressGroup) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "networksecurity",
			Version: "beta",
			Type:    "AddressGroup",
		},
		Object: make(map[string]interface{}),
	}
	if r.Capacity != nil {
		u.Object["capacity"] = *r.Capacity
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	var rItems []interface{}
	for _, rItemsVal := range r.Items {
		rItems = append(rItems, rItemsVal)
	}
	u.Object["items"] = rItems
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Parent != nil {
		u.Object["parent"] = *r.Parent
	}
	if r.Type != nil {
		u.Object["type"] = string(*r.Type)
	}
	return u
}

func UnstructuredToAddressGroup(u *unstructured.Resource) (*dclService.AddressGroup, error) {
	r := &dclService.AddressGroup{}
	if _, ok := u.Object["capacity"]; ok {
		if i, ok := u.Object["capacity"].(int64); ok {
			r.Capacity = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.Capacity: expected int64")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["items"]; ok {
		if s, ok := u.Object["items"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.Items = append(r.Items, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Items: expected []interface{}")
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
	if _, ok := u.Object["parent"]; ok {
		if s, ok := u.Object["parent"].(string); ok {
			r.Parent = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Parent: expected string")
		}
	}
	if _, ok := u.Object["type"]; ok {
		if s, ok := u.Object["type"].(string); ok {
			r.Type = dclService.AddressGroupTypeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Type: expected string")
		}
	}
	return r, nil
}

func GetAddressGroup(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAddressGroup(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetAddressGroup(ctx, r)
	if err != nil {
		return nil, err
	}
	return AddressGroupToUnstructured(r), nil
}

func ListAddressGroup(ctx context.Context, config *dcl.Config, location string, parent string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListAddressGroup(ctx, location, parent)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, AddressGroupToUnstructured(r))
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

func ApplyAddressGroup(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAddressGroup(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToAddressGroup(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyAddressGroup(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return AddressGroupToUnstructured(r), nil
}

func AddressGroupHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAddressGroup(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToAddressGroup(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyAddressGroup(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteAddressGroup(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAddressGroup(u)
	if err != nil {
		return err
	}
	return c.DeleteAddressGroup(ctx, r)
}

func AddressGroupID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToAddressGroup(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *AddressGroup) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"networksecurity",
		"AddressGroup",
		"beta",
	}
}

func (r *AddressGroup) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *AddressGroup) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *AddressGroup) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *AddressGroup) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *AddressGroup) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *AddressGroup) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *AddressGroup) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetAddressGroup(ctx, config, resource)
}

func (r *AddressGroup) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyAddressGroup(ctx, config, resource, opts...)
}

func (r *AddressGroup) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return AddressGroupHasDiff(ctx, config, resource, opts...)
}

func (r *AddressGroup) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteAddressGroup(ctx, config, resource)
}

func (r *AddressGroup) ID(resource *unstructured.Resource) (string, error) {
	return AddressGroupID(resource)
}

func init() {
	unstructured.Register(&AddressGroup{})
}
