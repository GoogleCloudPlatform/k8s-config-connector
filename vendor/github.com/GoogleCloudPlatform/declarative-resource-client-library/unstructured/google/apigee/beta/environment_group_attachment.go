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

type EnvironmentGroupAttachment struct{}

func EnvironmentGroupAttachmentToUnstructured(r *dclService.EnvironmentGroupAttachment) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "apigee",
			Version: "beta",
			Type:    "EnvironmentGroupAttachment",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreatedAt != nil {
		u.Object["createdAt"] = *r.CreatedAt
	}
	if r.Envgroup != nil {
		u.Object["envgroup"] = *r.Envgroup
	}
	if r.Environment != nil {
		u.Object["environment"] = *r.Environment
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	return u
}

func UnstructuredToEnvironmentGroupAttachment(u *unstructured.Resource) (*dclService.EnvironmentGroupAttachment, error) {
	r := &dclService.EnvironmentGroupAttachment{}
	if _, ok := u.Object["createdAt"]; ok {
		if i, ok := u.Object["createdAt"].(int64); ok {
			r.CreatedAt = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.CreatedAt: expected int64")
		}
	}
	if _, ok := u.Object["envgroup"]; ok {
		if s, ok := u.Object["envgroup"].(string); ok {
			r.Envgroup = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Envgroup: expected string")
		}
	}
	if _, ok := u.Object["environment"]; ok {
		if s, ok := u.Object["environment"].(string); ok {
			r.Environment = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Environment: expected string")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	return r, nil
}

func GetEnvironmentGroupAttachment(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEnvironmentGroupAttachment(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetEnvironmentGroupAttachment(ctx, r)
	if err != nil {
		return nil, err
	}
	return EnvironmentGroupAttachmentToUnstructured(r), nil
}

func ListEnvironmentGroupAttachment(ctx context.Context, config *dcl.Config, envgroup string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListEnvironmentGroupAttachment(ctx, envgroup)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, EnvironmentGroupAttachmentToUnstructured(r))
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

func ApplyEnvironmentGroupAttachment(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEnvironmentGroupAttachment(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToEnvironmentGroupAttachment(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyEnvironmentGroupAttachment(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return EnvironmentGroupAttachmentToUnstructured(r), nil
}

func EnvironmentGroupAttachmentHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEnvironmentGroupAttachment(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToEnvironmentGroupAttachment(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyEnvironmentGroupAttachment(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteEnvironmentGroupAttachment(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEnvironmentGroupAttachment(u)
	if err != nil {
		return err
	}
	return c.DeleteEnvironmentGroupAttachment(ctx, r)
}

func EnvironmentGroupAttachmentID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToEnvironmentGroupAttachment(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *EnvironmentGroupAttachment) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"apigee",
		"EnvironmentGroupAttachment",
		"beta",
	}
}

func (r *EnvironmentGroupAttachment) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *EnvironmentGroupAttachment) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *EnvironmentGroupAttachment) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *EnvironmentGroupAttachment) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *EnvironmentGroupAttachment) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *EnvironmentGroupAttachment) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *EnvironmentGroupAttachment) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetEnvironmentGroupAttachment(ctx, config, resource)
}

func (r *EnvironmentGroupAttachment) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyEnvironmentGroupAttachment(ctx, config, resource, opts...)
}

func (r *EnvironmentGroupAttachment) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return EnvironmentGroupAttachmentHasDiff(ctx, config, resource, opts...)
}

func (r *EnvironmentGroupAttachment) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteEnvironmentGroupAttachment(ctx, config, resource)
}

func (r *EnvironmentGroupAttachment) ID(resource *unstructured.Resource) (string, error) {
	return EnvironmentGroupAttachmentID(resource)
}

func init() {
	unstructured.Register(&EnvironmentGroupAttachment{})
}
