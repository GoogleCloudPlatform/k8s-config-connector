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
package cloudbuildv2

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuildv2"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Repository struct{}

func RepositoryToUnstructured(r *dclService.Repository) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "cloudbuildv2",
			Version: "ga",
			Type:    "Repository",
		},
		Object: make(map[string]interface{}),
	}
	if r.Annotations != nil {
		rAnnotations := make(map[string]interface{})
		for k, v := range r.Annotations {
			rAnnotations[k] = v
		}
		u.Object["annotations"] = rAnnotations
	}
	if r.Connection != nil {
		u.Object["connection"] = *r.Connection
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Etag != nil {
		u.Object["etag"] = *r.Etag
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
	if r.RemoteUri != nil {
		u.Object["remoteUri"] = *r.RemoteUri
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToRepository(u *unstructured.Resource) (*dclService.Repository, error) {
	r := &dclService.Repository{}
	if _, ok := u.Object["annotations"]; ok {
		if rAnnotations, ok := u.Object["annotations"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rAnnotations {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.Annotations = m
		} else {
			return nil, fmt.Errorf("r.Annotations: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["connection"]; ok {
		if s, ok := u.Object["connection"].(string); ok {
			r.Connection = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Connection: expected string")
		}
	}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["etag"]; ok {
		if s, ok := u.Object["etag"].(string); ok {
			r.Etag = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Etag: expected string")
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
	if _, ok := u.Object["remoteUri"]; ok {
		if s, ok := u.Object["remoteUri"].(string); ok {
			r.RemoteUri = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.RemoteUri: expected string")
		}
	}
	if _, ok := u.Object["updateTime"]; ok {
		if s, ok := u.Object["updateTime"].(string); ok {
			r.UpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UpdateTime: expected string")
		}
	}
	return r, nil
}

func GetRepository(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToRepository(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetRepository(ctx, r)
	if err != nil {
		return nil, err
	}
	return RepositoryToUnstructured(r), nil
}

func ListRepository(ctx context.Context, config *dcl.Config, project string, location string, connection string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListRepository(ctx, project, location, connection)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, RepositoryToUnstructured(r))
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

func ApplyRepository(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToRepository(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToRepository(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyRepository(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return RepositoryToUnstructured(r), nil
}

func RepositoryHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToRepository(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToRepository(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyRepository(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteRepository(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToRepository(u)
	if err != nil {
		return err
	}
	return c.DeleteRepository(ctx, r)
}

func RepositoryID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToRepository(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Repository) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"cloudbuildv2",
		"Repository",
		"ga",
	}
}

func (r *Repository) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Repository) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Repository) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Repository) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Repository) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Repository) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Repository) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetRepository(ctx, config, resource)
}

func (r *Repository) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyRepository(ctx, config, resource, opts...)
}

func (r *Repository) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return RepositoryHasDiff(ctx, config, resource, opts...)
}

func (r *Repository) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteRepository(ctx, config, resource)
}

func (r *Repository) ID(resource *unstructured.Resource) (string, error) {
	return RepositoryID(resource)
}

func init() {
	unstructured.Register(&Repository{})
}
