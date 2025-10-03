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
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Group struct{}

func GroupToUnstructured(r *dclService.Group) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "monitoring",
			Version: "alpha",
			Type:    "Group",
		},
		Object: make(map[string]interface{}),
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.Filter != nil {
		u.Object["filter"] = *r.Filter
	}
	if r.IsCluster != nil {
		u.Object["isCluster"] = *r.IsCluster
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.ParentName != nil {
		u.Object["parentName"] = *r.ParentName
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	return u
}

func UnstructuredToGroup(u *unstructured.Resource) (*dclService.Group, error) {
	r := &dclService.Group{}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
		}
	}
	if _, ok := u.Object["filter"]; ok {
		if s, ok := u.Object["filter"].(string); ok {
			r.Filter = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Filter: expected string")
		}
	}
	if _, ok := u.Object["isCluster"]; ok {
		if b, ok := u.Object["isCluster"].(bool); ok {
			r.IsCluster = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.IsCluster: expected bool")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["parentName"]; ok {
		if s, ok := u.Object["parentName"].(string); ok {
			r.ParentName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ParentName: expected string")
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

func GetGroup(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToGroup(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetGroup(ctx, r)
	if err != nil {
		return nil, err
	}
	return GroupToUnstructured(r), nil
}

func ListGroup(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListGroup(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, GroupToUnstructured(r))
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

func ApplyGroup(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToGroup(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToGroup(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyGroup(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return GroupToUnstructured(r), nil
}

func GroupHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToGroup(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToGroup(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyGroup(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteGroup(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToGroup(u)
	if err != nil {
		return err
	}
	return c.DeleteGroup(ctx, r)
}

func GroupID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToGroup(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Group) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"monitoring",
		"Group",
		"alpha",
	}
}

func (r *Group) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Group) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Group) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Group) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Group) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Group) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Group) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetGroup(ctx, config, resource)
}

func (r *Group) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyGroup(ctx, config, resource, opts...)
}

func (r *Group) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return GroupHasDiff(ctx, config, resource, opts...)
}

func (r *Group) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteGroup(ctx, config, resource)
}

func (r *Group) ID(resource *unstructured.Resource) (string, error) {
	return GroupID(resource)
}

func init() {
	unstructured.Register(&Group{})
}
