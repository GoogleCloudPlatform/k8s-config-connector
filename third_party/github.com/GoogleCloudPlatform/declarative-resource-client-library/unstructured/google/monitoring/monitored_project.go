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
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type MonitoredProject struct{}

func MonitoredProjectToUnstructured(r *dclService.MonitoredProject) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "monitoring",
			Version: "ga",
			Type:    "MonitoredProject",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.MetricsScope != nil {
		u.Object["metricsScope"] = *r.MetricsScope
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	return u
}

func UnstructuredToMonitoredProject(u *unstructured.Resource) (*dclService.MonitoredProject, error) {
	r := &dclService.MonitoredProject{}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["metricsScope"]; ok {
		if s, ok := u.Object["metricsScope"].(string); ok {
			r.MetricsScope = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.MetricsScope: expected string")
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

func GetMonitoredProject(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToMonitoredProject(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetMonitoredProject(ctx, r)
	if err != nil {
		return nil, err
	}
	return MonitoredProjectToUnstructured(r), nil
}

func ListMonitoredProject(ctx context.Context, config *dcl.Config, metricsScope string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListMonitoredProject(ctx, metricsScope)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, MonitoredProjectToUnstructured(r))
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

func ApplyMonitoredProject(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToMonitoredProject(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToMonitoredProject(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyMonitoredProject(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return MonitoredProjectToUnstructured(r), nil
}

func MonitoredProjectHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToMonitoredProject(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToMonitoredProject(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyMonitoredProject(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteMonitoredProject(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToMonitoredProject(u)
	if err != nil {
		return err
	}
	return c.DeleteMonitoredProject(ctx, r)
}

func MonitoredProjectID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToMonitoredProject(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *MonitoredProject) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"monitoring",
		"MonitoredProject",
		"ga",
	}
}

func (r *MonitoredProject) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *MonitoredProject) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *MonitoredProject) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *MonitoredProject) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *MonitoredProject) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *MonitoredProject) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *MonitoredProject) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetMonitoredProject(ctx, config, resource)
}

func (r *MonitoredProject) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyMonitoredProject(ctx, config, resource, opts...)
}

func (r *MonitoredProject) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return MonitoredProjectHasDiff(ctx, config, resource, opts...)
}

func (r *MonitoredProject) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteMonitoredProject(ctx, config, resource)
}

func (r *MonitoredProject) ID(resource *unstructured.Resource) (string, error) {
	return MonitoredProjectID(resource)
}

func init() {
	unstructured.Register(&MonitoredProject{})
}
