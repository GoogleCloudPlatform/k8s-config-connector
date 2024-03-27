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

type MetricsScope struct{}

func MetricsScopeToUnstructured(r *dclService.MetricsScope) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "monitoring",
			Version: "ga",
			Type:    "MetricsScope",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	var rMonitoredProjects []interface{}
	for _, rMonitoredProjectsVal := range r.MonitoredProjects {
		rMonitoredProjectsObject := make(map[string]interface{})
		if rMonitoredProjectsVal.CreateTime != nil {
			rMonitoredProjectsObject["createTime"] = *rMonitoredProjectsVal.CreateTime
		}
		if rMonitoredProjectsVal.Name != nil {
			rMonitoredProjectsObject["name"] = *rMonitoredProjectsVal.Name
		}
		rMonitoredProjects = append(rMonitoredProjects, rMonitoredProjectsObject)
	}
	u.Object["monitoredProjects"] = rMonitoredProjects
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToMetricsScope(u *unstructured.Resource) (*dclService.MetricsScope, error) {
	r := &dclService.MetricsScope{}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["monitoredProjects"]; ok {
		if s, ok := u.Object["monitoredProjects"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rMonitoredProjects dclService.MetricsScopeMonitoredProjects
					if _, ok := objval["createTime"]; ok {
						if s, ok := objval["createTime"].(string); ok {
							rMonitoredProjects.CreateTime = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rMonitoredProjects.CreateTime: expected string")
						}
					}
					if _, ok := objval["name"]; ok {
						if s, ok := objval["name"].(string); ok {
							rMonitoredProjects.Name = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rMonitoredProjects.Name: expected string")
						}
					}
					r.MonitoredProjects = append(r.MonitoredProjects, rMonitoredProjects)
				}
			}
		} else {
			return nil, fmt.Errorf("r.MonitoredProjects: expected []interface{}")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
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

func GetMetricsScope(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToMetricsScope(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetMetricsScope(ctx, r)
	if err != nil {
		return nil, err
	}
	return MetricsScopeToUnstructured(r), nil
}

func ApplyMetricsScope(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToMetricsScope(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToMetricsScope(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyMetricsScope(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return MetricsScopeToUnstructured(r), nil
}

func MetricsScopeHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToMetricsScope(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToMetricsScope(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyMetricsScope(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteMetricsScope(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func MetricsScopeID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToMetricsScope(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *MetricsScope) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"monitoring",
		"MetricsScope",
		"ga",
	}
}

func (r *MetricsScope) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *MetricsScope) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *MetricsScope) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *MetricsScope) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *MetricsScope) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *MetricsScope) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *MetricsScope) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetMetricsScope(ctx, config, resource)
}

func (r *MetricsScope) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyMetricsScope(ctx, config, resource, opts...)
}

func (r *MetricsScope) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return MetricsScopeHasDiff(ctx, config, resource, opts...)
}

func (r *MetricsScope) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteMetricsScope(ctx, config, resource)
}

func (r *MetricsScope) ID(resource *unstructured.Resource) (string, error) {
	return MetricsScopeID(resource)
}

func init() {
	unstructured.Register(&MetricsScope{})
}
