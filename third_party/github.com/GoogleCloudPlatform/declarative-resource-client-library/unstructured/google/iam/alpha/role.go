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
package iam

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Role struct{}

func RoleToUnstructured(r *dclService.Role) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "iam",
			Version: "alpha",
			Type:    "Role",
		},
		Object: make(map[string]interface{}),
	}
	if r.Deleted != nil {
		u.Object["deleted"] = *r.Deleted
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.Etag != nil {
		u.Object["etag"] = *r.Etag
	}
	if r.GroupName != nil {
		u.Object["groupName"] = *r.GroupName
	}
	if r.GroupTitle != nil {
		u.Object["groupTitle"] = *r.GroupTitle
	}
	var rIncludedPermissions []interface{}
	for _, rIncludedPermissionsVal := range r.IncludedPermissions {
		rIncludedPermissions = append(rIncludedPermissions, rIncludedPermissionsVal)
	}
	u.Object["includedPermissions"] = rIncludedPermissions
	var rIncludedRoles []interface{}
	for _, rIncludedRolesVal := range r.IncludedRoles {
		rIncludedRoles = append(rIncludedRoles, rIncludedRolesVal)
	}
	u.Object["includedRoles"] = rIncludedRoles
	if r.LifecyclePhase != nil {
		u.Object["lifecyclePhase"] = *r.LifecyclePhase
	}
	if r.LocalizedValues != nil && r.LocalizedValues != dclService.EmptyRoleLocalizedValues {
		rLocalizedValues := make(map[string]interface{})
		if r.LocalizedValues.LocalizedDescription != nil {
			rLocalizedValues["localizedDescription"] = *r.LocalizedValues.LocalizedDescription
		}
		if r.LocalizedValues.LocalizedTitle != nil {
			rLocalizedValues["localizedTitle"] = *r.LocalizedValues.LocalizedTitle
		}
		u.Object["localizedValues"] = rLocalizedValues
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Parent != nil {
		u.Object["parent"] = *r.Parent
	}
	if r.Stage != nil {
		u.Object["stage"] = string(*r.Stage)
	}
	if r.Title != nil {
		u.Object["title"] = *r.Title
	}
	return u
}

func UnstructuredToRole(u *unstructured.Resource) (*dclService.Role, error) {
	r := &dclService.Role{}
	if _, ok := u.Object["deleted"]; ok {
		if b, ok := u.Object["deleted"].(bool); ok {
			r.Deleted = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.Deleted: expected bool")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["etag"]; ok {
		if s, ok := u.Object["etag"].(string); ok {
			r.Etag = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Etag: expected string")
		}
	}
	if _, ok := u.Object["groupName"]; ok {
		if s, ok := u.Object["groupName"].(string); ok {
			r.GroupName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.GroupName: expected string")
		}
	}
	if _, ok := u.Object["groupTitle"]; ok {
		if s, ok := u.Object["groupTitle"].(string); ok {
			r.GroupTitle = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.GroupTitle: expected string")
		}
	}
	if _, ok := u.Object["includedPermissions"]; ok {
		if s, ok := u.Object["includedPermissions"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.IncludedPermissions = append(r.IncludedPermissions, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.IncludedPermissions: expected []interface{}")
		}
	}
	if _, ok := u.Object["includedRoles"]; ok {
		if s, ok := u.Object["includedRoles"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.IncludedRoles = append(r.IncludedRoles, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.IncludedRoles: expected []interface{}")
		}
	}
	if _, ok := u.Object["lifecyclePhase"]; ok {
		if s, ok := u.Object["lifecyclePhase"].(string); ok {
			r.LifecyclePhase = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.LifecyclePhase: expected string")
		}
	}
	if _, ok := u.Object["localizedValues"]; ok {
		if rLocalizedValues, ok := u.Object["localizedValues"].(map[string]interface{}); ok {
			r.LocalizedValues = &dclService.RoleLocalizedValues{}
			if _, ok := rLocalizedValues["localizedDescription"]; ok {
				if s, ok := rLocalizedValues["localizedDescription"].(string); ok {
					r.LocalizedValues.LocalizedDescription = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.LocalizedValues.LocalizedDescription: expected string")
				}
			}
			if _, ok := rLocalizedValues["localizedTitle"]; ok {
				if s, ok := rLocalizedValues["localizedTitle"].(string); ok {
					r.LocalizedValues.LocalizedTitle = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.LocalizedValues.LocalizedTitle: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.LocalizedValues: expected map[string]interface{}")
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
	if _, ok := u.Object["stage"]; ok {
		if s, ok := u.Object["stage"].(string); ok {
			r.Stage = dclService.RoleStageEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Stage: expected string")
		}
	}
	if _, ok := u.Object["title"]; ok {
		if s, ok := u.Object["title"].(string); ok {
			r.Title = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Title: expected string")
		}
	}
	return r, nil
}

func GetRole(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToRole(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetRole(ctx, r)
	if err != nil {
		return nil, err
	}
	return RoleToUnstructured(r), nil
}

func ListRole(ctx context.Context, config *dcl.Config, parent string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListRole(ctx, parent)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, RoleToUnstructured(r))
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

func ApplyRole(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToRole(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToRole(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyRole(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return RoleToUnstructured(r), nil
}

func RoleHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToRole(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToRole(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyRole(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteRole(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToRole(u)
	if err != nil {
		return err
	}
	return c.DeleteRole(ctx, r)
}

func RoleID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToRole(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Role) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"iam",
		"Role",
		"alpha",
	}
}

func (r *Role) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Role) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Role) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Role) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Role) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Role) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Role) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetRole(ctx, config, resource)
}

func (r *Role) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyRole(ctx, config, resource, opts...)
}

func (r *Role) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return RoleHasDiff(ctx, config, resource, opts...)
}

func (r *Role) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteRole(ctx, config, resource)
}

func (r *Role) ID(resource *unstructured.Resource) (string, error) {
	return RoleID(resource)
}

func init() {
	unstructured.Register(&Role{})
}
