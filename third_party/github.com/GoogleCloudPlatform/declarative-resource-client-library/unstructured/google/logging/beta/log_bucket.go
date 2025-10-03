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
package logging

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/logging/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type LogBucket struct{}

func LogBucketToUnstructured(r *dclService.LogBucket) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "logging",
			Version: "beta",
			Type:    "LogBucket",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.EnableAnalytics != nil {
		u.Object["enableAnalytics"] = *r.EnableAnalytics
	}
	if r.LifecycleState != nil {
		u.Object["lifecycleState"] = string(*r.LifecycleState)
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Locked != nil {
		u.Object["locked"] = *r.Locked
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Parent != nil {
		u.Object["parent"] = *r.Parent
	}
	if r.RetentionDays != nil {
		u.Object["retentionDays"] = *r.RetentionDays
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToLogBucket(u *unstructured.Resource) (*dclService.LogBucket, error) {
	r := &dclService.LogBucket{}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["enableAnalytics"]; ok {
		if b, ok := u.Object["enableAnalytics"].(bool); ok {
			r.EnableAnalytics = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.EnableAnalytics: expected bool")
		}
	}
	if _, ok := u.Object["lifecycleState"]; ok {
		if s, ok := u.Object["lifecycleState"].(string); ok {
			r.LifecycleState = dclService.LogBucketLifecycleStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.LifecycleState: expected string")
		}
	}
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["locked"]; ok {
		if b, ok := u.Object["locked"].(bool); ok {
			r.Locked = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.Locked: expected bool")
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
	if _, ok := u.Object["retentionDays"]; ok {
		if i, ok := u.Object["retentionDays"].(int64); ok {
			r.RetentionDays = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.RetentionDays: expected int64")
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

func GetLogBucket(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToLogBucket(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetLogBucket(ctx, r)
	if err != nil {
		return nil, err
	}
	return LogBucketToUnstructured(r), nil
}

func ListLogBucket(ctx context.Context, config *dcl.Config, location string, parent string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListLogBucket(ctx, location, parent)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, LogBucketToUnstructured(r))
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

func ApplyLogBucket(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToLogBucket(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToLogBucket(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyLogBucket(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return LogBucketToUnstructured(r), nil
}

func LogBucketHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToLogBucket(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToLogBucket(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyLogBucket(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteLogBucket(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToLogBucket(u)
	if err != nil {
		return err
	}
	return c.DeleteLogBucket(ctx, r)
}

func LogBucketID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToLogBucket(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *LogBucket) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"logging",
		"LogBucket",
		"beta",
	}
}

func (r *LogBucket) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *LogBucket) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *LogBucket) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *LogBucket) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *LogBucket) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *LogBucket) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *LogBucket) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetLogBucket(ctx, config, resource)
}

func (r *LogBucket) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyLogBucket(ctx, config, resource, opts...)
}

func (r *LogBucket) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return LogBucketHasDiff(ctx, config, resource, opts...)
}

func (r *LogBucket) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteLogBucket(ctx, config, resource)
}

func (r *LogBucket) ID(resource *unstructured.Resource) (string, error) {
	return LogBucketID(resource)
}

func init() {
	unstructured.Register(&LogBucket{})
}
