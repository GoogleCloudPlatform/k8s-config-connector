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
package containerazure

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containerazure/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Client struct{}

func ClientToUnstructured(r *dclService.AzureClient) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "containerazure",
			Version: "beta",
			Type:    "Client",
		},
		Object: make(map[string]interface{}),
	}
	if r.ApplicationId != nil {
		u.Object["applicationId"] = *r.ApplicationId
	}
	if r.Certificate != nil {
		u.Object["certificate"] = *r.Certificate
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
	if r.TenantId != nil {
		u.Object["tenantId"] = *r.TenantId
	}
	if r.Uid != nil {
		u.Object["uid"] = *r.Uid
	}
	return u
}

func UnstructuredToClient(u *unstructured.Resource) (*dclService.AzureClient, error) {
	r := &dclService.AzureClient{}
	if _, ok := u.Object["applicationId"]; ok {
		if s, ok := u.Object["applicationId"].(string); ok {
			r.ApplicationId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ApplicationId: expected string")
		}
	}
	if _, ok := u.Object["certificate"]; ok {
		if s, ok := u.Object["certificate"].(string); ok {
			r.Certificate = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Certificate: expected string")
		}
	}
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
	if _, ok := u.Object["tenantId"]; ok {
		if s, ok := u.Object["tenantId"].(string); ok {
			r.TenantId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.TenantId: expected string")
		}
	}
	if _, ok := u.Object["uid"]; ok {
		if s, ok := u.Object["uid"].(string); ok {
			r.Uid = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Uid: expected string")
		}
	}
	return r, nil
}

func GetClient(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToClient(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetClient(ctx, r)
	if err != nil {
		return nil, err
	}
	return ClientToUnstructured(r), nil
}

func ListClient(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListClient(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, ClientToUnstructured(r))
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

func ApplyClient(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToClient(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToClient(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyClient(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return ClientToUnstructured(r), nil
}

func ClientHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToClient(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToClient(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyClient(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteClient(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToClient(u)
	if err != nil {
		return err
	}
	return c.DeleteClient(ctx, r)
}

func ClientID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToClient(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Client) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"containerazure",
		"Client",
		"beta",
	}
}

func (r *Client) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Client) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Client) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Client) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Client) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Client) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Client) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetClient(ctx, config, resource)
}

func (r *Client) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyClient(ctx, config, resource, opts...)
}

func (r *Client) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return ClientHasDiff(ctx, config, resource, opts...)
}

func (r *Client) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteClient(ctx, config, resource)
}

func (r *Client) ID(resource *unstructured.Resource) (string, error) {
	return ClientID(resource)
}

func init() {
	unstructured.Register(&Client{})
}
