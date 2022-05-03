// Copyright 2022 Google LLC. All Rights Reserved.
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
package apigee

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apigee/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Envgroup struct{}

func EnvgroupToUnstructured(r *dclService.Envgroup) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "apigee",
			Version: "beta",
			Type:    "Envgroup",
		},
		Object: make(map[string]interface{}),
	}
	if r.ApigeeOrganization != nil {
		u.Object["apigeeOrganization"] = *r.ApigeeOrganization
	}
	if r.CreatedAt != nil {
		u.Object["createdAt"] = *r.CreatedAt
	}
	var rHostnames []interface{}
	for _, rHostnamesVal := range r.Hostnames {
		rHostnames = append(rHostnames, rHostnamesVal)
	}
	u.Object["hostnames"] = rHostnames
	if r.LastModifiedAt != nil {
		u.Object["lastModifiedAt"] = *r.LastModifiedAt
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.State != nil {
		u.Object["state"] = string(*r.State)
	}
	return u
}

func UnstructuredToEnvgroup(u *unstructured.Resource) (*dclService.Envgroup, error) {
	r := &dclService.Envgroup{}
	if _, ok := u.Object["apigeeOrganization"]; ok {
		if s, ok := u.Object["apigeeOrganization"].(string); ok {
			r.ApigeeOrganization = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ApigeeOrganization: expected string")
		}
	}
	if _, ok := u.Object["createdAt"]; ok {
		if i, ok := u.Object["createdAt"].(int64); ok {
			r.CreatedAt = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.CreatedAt: expected int64")
		}
	}
	if _, ok := u.Object["hostnames"]; ok {
		if s, ok := u.Object["hostnames"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.Hostnames = append(r.Hostnames, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Hostnames: expected []interface{}")
		}
	}
	if _, ok := u.Object["lastModifiedAt"]; ok {
		if i, ok := u.Object["lastModifiedAt"].(int64); ok {
			r.LastModifiedAt = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.LastModifiedAt: expected int64")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["state"]; ok {
		if s, ok := u.Object["state"].(string); ok {
			r.State = dclService.EnvgroupStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.State: expected string")
		}
	}
	return r, nil
}

func GetEnvgroup(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEnvgroup(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetEnvgroup(ctx, r)
	if err != nil {
		return nil, err
	}
	return EnvgroupToUnstructured(r), nil
}

func ListEnvgroup(ctx context.Context, config *dcl.Config, apigeeorganization string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListEnvgroup(ctx, apigeeorganization)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, EnvgroupToUnstructured(r))
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

func ApplyEnvgroup(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEnvgroup(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToEnvgroup(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyEnvgroup(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return EnvgroupToUnstructured(r), nil
}

func EnvgroupHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEnvgroup(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToEnvgroup(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyEnvgroup(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteEnvgroup(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEnvgroup(u)
	if err != nil {
		return err
	}
	return c.DeleteEnvgroup(ctx, r)
}

func EnvgroupID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToEnvgroup(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Envgroup) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"apigee",
		"Envgroup",
		"beta",
	}
}

func (r *Envgroup) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Envgroup) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Envgroup) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Envgroup) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Envgroup) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Envgroup) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Envgroup) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetEnvgroup(ctx, config, resource)
}

func (r *Envgroup) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyEnvgroup(ctx, config, resource, opts...)
}

func (r *Envgroup) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return EnvgroupHasDiff(ctx, config, resource, opts...)
}

func (r *Envgroup) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteEnvgroup(ctx, config, resource)
}

func (r *Envgroup) ID(resource *unstructured.Resource) (string, error) {
	return EnvgroupID(resource)
}

func init() {
	unstructured.Register(&Envgroup{})
}
