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

type Service struct{}

func ServiceToUnstructured(r *dclService.Service) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "monitoring",
			Version: "alpha",
			Type:    "Service",
		},
		Object: make(map[string]interface{}),
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.Telemetry != nil && r.Telemetry != dclService.EmptyServiceTelemetry {
		rTelemetry := make(map[string]interface{})
		if r.Telemetry.ResourceName != nil {
			rTelemetry["resourceName"] = *r.Telemetry.ResourceName
		}
		u.Object["telemetry"] = rTelemetry
	}
	if r.UserLabels != nil {
		rUserLabels := make(map[string]interface{})
		for k, v := range r.UserLabels {
			rUserLabels[k] = v
		}
		u.Object["userLabels"] = rUserLabels
	}
	return u
}

func UnstructuredToService(u *unstructured.Resource) (*dclService.Service, error) {
	r := &dclService.Service{}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
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
	if _, ok := u.Object["telemetry"]; ok {
		if rTelemetry, ok := u.Object["telemetry"].(map[string]interface{}); ok {
			r.Telemetry = &dclService.ServiceTelemetry{}
			if _, ok := rTelemetry["resourceName"]; ok {
				if s, ok := rTelemetry["resourceName"].(string); ok {
					r.Telemetry.ResourceName = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Telemetry.ResourceName: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Telemetry: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["userLabels"]; ok {
		if rUserLabels, ok := u.Object["userLabels"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rUserLabels {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.UserLabels = m
		} else {
			return nil, fmt.Errorf("r.UserLabels: expected map[string]interface{}")
		}
	}
	return r, nil
}

func GetService(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToService(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetService(ctx, r)
	if err != nil {
		return nil, err
	}
	return ServiceToUnstructured(r), nil
}

func ListService(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListService(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, ServiceToUnstructured(r))
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

func ApplyService(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToService(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToService(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyService(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return ServiceToUnstructured(r), nil
}

func ServiceHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToService(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToService(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyService(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteService(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToService(u)
	if err != nil {
		return err
	}
	return c.DeleteService(ctx, r)
}

func ServiceID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToService(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Service) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"monitoring",
		"Service",
		"alpha",
	}
}

func (r *Service) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Service) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Service) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Service) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Service) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Service) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Service) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetService(ctx, config, resource)
}

func (r *Service) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyService(ctx, config, resource, opts...)
}

func (r *Service) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return ServiceHasDiff(ctx, config, resource, opts...)
}

func (r *Service) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteService(ctx, config, resource)
}

func (r *Service) ID(resource *unstructured.Resource) (string, error) {
	return ServiceID(resource)
}

func init() {
	unstructured.Register(&Service{})
}
