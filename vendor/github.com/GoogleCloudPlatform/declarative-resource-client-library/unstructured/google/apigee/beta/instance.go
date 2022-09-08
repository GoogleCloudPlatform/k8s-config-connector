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

type Instance struct{}

func InstanceToUnstructured(r *dclService.Instance) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "apigee",
			Version: "beta",
			Type:    "Instance",
		},
		Object: make(map[string]interface{}),
	}
	if r.ApigeeOrganization != nil {
		u.Object["apigeeOrganization"] = *r.ApigeeOrganization
	}
	if r.CreatedAt != nil {
		u.Object["createdAt"] = *r.CreatedAt
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.DiskEncryptionKeyName != nil {
		u.Object["diskEncryptionKeyName"] = *r.DiskEncryptionKeyName
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.Host != nil {
		u.Object["host"] = *r.Host
	}
	if r.LastModifiedAt != nil {
		u.Object["lastModifiedAt"] = *r.LastModifiedAt
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.PeeringCidrRange != nil {
		u.Object["peeringCidrRange"] = string(*r.PeeringCidrRange)
	}
	if r.Port != nil {
		u.Object["port"] = *r.Port
	}
	if r.State != nil {
		u.Object["state"] = string(*r.State)
	}
	return u
}

func UnstructuredToInstance(u *unstructured.Resource) (*dclService.Instance, error) {
	r := &dclService.Instance{}
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
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["diskEncryptionKeyName"]; ok {
		if s, ok := u.Object["diskEncryptionKeyName"].(string); ok {
			r.DiskEncryptionKeyName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DiskEncryptionKeyName: expected string")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
		}
	}
	if _, ok := u.Object["host"]; ok {
		if s, ok := u.Object["host"].(string); ok {
			r.Host = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Host: expected string")
		}
	}
	if _, ok := u.Object["lastModifiedAt"]; ok {
		if i, ok := u.Object["lastModifiedAt"].(int64); ok {
			r.LastModifiedAt = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.LastModifiedAt: expected int64")
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
	if _, ok := u.Object["peeringCidrRange"]; ok {
		if s, ok := u.Object["peeringCidrRange"].(string); ok {
			r.PeeringCidrRange = dclService.InstancePeeringCidrRangeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.PeeringCidrRange: expected string")
		}
	}
	if _, ok := u.Object["port"]; ok {
		if s, ok := u.Object["port"].(string); ok {
			r.Port = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Port: expected string")
		}
	}
	if _, ok := u.Object["state"]; ok {
		if s, ok := u.Object["state"].(string); ok {
			r.State = dclService.InstanceStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.State: expected string")
		}
	}
	return r, nil
}

func GetInstance(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToInstance(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetInstance(ctx, r)
	if err != nil {
		return nil, err
	}
	return InstanceToUnstructured(r), nil
}

func ListInstance(ctx context.Context, config *dcl.Config, apigeeOrganization string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListInstance(ctx, apigeeOrganization)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, InstanceToUnstructured(r))
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

func ApplyInstance(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToInstance(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToInstance(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyInstance(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return InstanceToUnstructured(r), nil
}

func InstanceHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToInstance(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToInstance(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyInstance(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteInstance(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToInstance(u)
	if err != nil {
		return err
	}
	return c.DeleteInstance(ctx, r)
}

func InstanceID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToInstance(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Instance) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"apigee",
		"Instance",
		"beta",
	}
}

func (r *Instance) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Instance) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Instance) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Instance) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Instance) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Instance) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Instance) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetInstance(ctx, config, resource)
}

func (r *Instance) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyInstance(ctx, config, resource, opts...)
}

func (r *Instance) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return InstanceHasDiff(ctx, config, resource, opts...)
}

func (r *Instance) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteInstance(ctx, config, resource)
}

func (r *Instance) ID(resource *unstructured.Resource) (string, error) {
	return InstanceID(resource)
}

func init() {
	unstructured.Register(&Instance{})
}
