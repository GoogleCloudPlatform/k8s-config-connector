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
package cloudbuild

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuild/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type WorkerPool struct{}

func WorkerPoolToUnstructured(r *dclService.WorkerPool) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "cloudbuild",
			Version: "alpha",
			Type:    "WorkerPool",
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
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.DeleteTime != nil {
		u.Object["deleteTime"] = *r.DeleteTime
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
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
	if r.NetworkConfig != nil && r.NetworkConfig != dclService.EmptyWorkerPoolNetworkConfig {
		rNetworkConfig := make(map[string]interface{})
		if r.NetworkConfig.PeeredNetwork != nil {
			rNetworkConfig["peeredNetwork"] = *r.NetworkConfig.PeeredNetwork
		}
		if r.NetworkConfig.PeeredNetworkIPRange != nil {
			rNetworkConfig["peeredNetworkIPRange"] = *r.NetworkConfig.PeeredNetworkIPRange
		}
		u.Object["networkConfig"] = rNetworkConfig
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.State != nil {
		u.Object["state"] = string(*r.State)
	}
	if r.Uid != nil {
		u.Object["uid"] = *r.Uid
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	if r.WorkerConfig != nil && r.WorkerConfig != dclService.EmptyWorkerPoolWorkerConfig {
		rWorkerConfig := make(map[string]interface{})
		if r.WorkerConfig.DiskSizeGb != nil {
			rWorkerConfig["diskSizeGb"] = *r.WorkerConfig.DiskSizeGb
		}
		if r.WorkerConfig.MachineType != nil {
			rWorkerConfig["machineType"] = *r.WorkerConfig.MachineType
		}
		if r.WorkerConfig.NoExternalIP != nil {
			rWorkerConfig["noExternalIP"] = *r.WorkerConfig.NoExternalIP
		}
		u.Object["workerConfig"] = rWorkerConfig
	}
	return u
}

func UnstructuredToWorkerPool(u *unstructured.Resource) (*dclService.WorkerPool, error) {
	r := &dclService.WorkerPool{}
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
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["deleteTime"]; ok {
		if s, ok := u.Object["deleteTime"].(string); ok {
			r.DeleteTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DeleteTime: expected string")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
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
	if _, ok := u.Object["networkConfig"]; ok {
		if rNetworkConfig, ok := u.Object["networkConfig"].(map[string]interface{}); ok {
			r.NetworkConfig = &dclService.WorkerPoolNetworkConfig{}
			if _, ok := rNetworkConfig["peeredNetwork"]; ok {
				if s, ok := rNetworkConfig["peeredNetwork"].(string); ok {
					r.NetworkConfig.PeeredNetwork = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.NetworkConfig.PeeredNetwork: expected string")
				}
			}
			if _, ok := rNetworkConfig["peeredNetworkIPRange"]; ok {
				if s, ok := rNetworkConfig["peeredNetworkIPRange"].(string); ok {
					r.NetworkConfig.PeeredNetworkIPRange = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.NetworkConfig.PeeredNetworkIPRange: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.NetworkConfig: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["state"]; ok {
		if s, ok := u.Object["state"].(string); ok {
			r.State = dclService.WorkerPoolStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.State: expected string")
		}
	}
	if _, ok := u.Object["uid"]; ok {
		if s, ok := u.Object["uid"].(string); ok {
			r.Uid = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Uid: expected string")
		}
	}
	if _, ok := u.Object["updateTime"]; ok {
		if s, ok := u.Object["updateTime"].(string); ok {
			r.UpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UpdateTime: expected string")
		}
	}
	if _, ok := u.Object["workerConfig"]; ok {
		if rWorkerConfig, ok := u.Object["workerConfig"].(map[string]interface{}); ok {
			r.WorkerConfig = &dclService.WorkerPoolWorkerConfig{}
			if _, ok := rWorkerConfig["diskSizeGb"]; ok {
				if i, ok := rWorkerConfig["diskSizeGb"].(int64); ok {
					r.WorkerConfig.DiskSizeGb = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.WorkerConfig.DiskSizeGb: expected int64")
				}
			}
			if _, ok := rWorkerConfig["machineType"]; ok {
				if s, ok := rWorkerConfig["machineType"].(string); ok {
					r.WorkerConfig.MachineType = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.WorkerConfig.MachineType: expected string")
				}
			}
			if _, ok := rWorkerConfig["noExternalIP"]; ok {
				if b, ok := rWorkerConfig["noExternalIP"].(bool); ok {
					r.WorkerConfig.NoExternalIP = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.WorkerConfig.NoExternalIP: expected bool")
				}
			}
		} else {
			return nil, fmt.Errorf("r.WorkerConfig: expected map[string]interface{}")
		}
	}
	return r, nil
}

func GetWorkerPool(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWorkerPool(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetWorkerPool(ctx, r)
	if err != nil {
		return nil, err
	}
	return WorkerPoolToUnstructured(r), nil
}

func ListWorkerPool(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListWorkerPool(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, WorkerPoolToUnstructured(r))
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

func ApplyWorkerPool(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWorkerPool(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToWorkerPool(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyWorkerPool(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return WorkerPoolToUnstructured(r), nil
}

func WorkerPoolHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWorkerPool(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToWorkerPool(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyWorkerPool(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteWorkerPool(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWorkerPool(u)
	if err != nil {
		return err
	}
	return c.DeleteWorkerPool(ctx, r)
}

func WorkerPoolID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToWorkerPool(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *WorkerPool) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"cloudbuild",
		"WorkerPool",
		"alpha",
	}
}

func (r *WorkerPool) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *WorkerPool) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *WorkerPool) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *WorkerPool) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *WorkerPool) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *WorkerPool) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *WorkerPool) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetWorkerPool(ctx, config, resource)
}

func (r *WorkerPool) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyWorkerPool(ctx, config, resource, opts...)
}

func (r *WorkerPool) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return WorkerPoolHasDiff(ctx, config, resource, opts...)
}

func (r *WorkerPool) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteWorkerPool(ctx, config, resource)
}

func (r *WorkerPool) ID(resource *unstructured.Resource) (string, error) {
	return WorkerPoolID(resource)
}

func init() {
	unstructured.Register(&WorkerPool{})
}
