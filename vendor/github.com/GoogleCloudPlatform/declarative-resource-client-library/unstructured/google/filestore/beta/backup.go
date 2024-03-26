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
package filestore

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/filestore/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Backup struct{}

func BackupToUnstructured(r *dclService.Backup) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "filestore",
			Version: "beta",
			Type:    "Backup",
		},
		Object: make(map[string]interface{}),
	}
	if r.CapacityGb != nil {
		u.Object["capacityGb"] = *r.CapacityGb
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.DownloadBytes != nil {
		u.Object["downloadBytes"] = *r.DownloadBytes
	}
	if r.Labels != nil {
		rLabels := make(map[string]interface{})
		for k, v := range r.Labels {
			rLabels[k] = v
		}
		u.Object["labels"] = rLabels
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
	if r.SourceFileShare != nil {
		u.Object["sourceFileShare"] = *r.SourceFileShare
	}
	if r.SourceInstance != nil {
		u.Object["sourceInstance"] = *r.SourceInstance
	}
	if r.SourceInstanceTier != nil {
		u.Object["sourceInstanceTier"] = string(*r.SourceInstanceTier)
	}
	if r.State != nil {
		u.Object["state"] = string(*r.State)
	}
	if r.StorageBytes != nil {
		u.Object["storageBytes"] = *r.StorageBytes
	}
	return u
}

func UnstructuredToBackup(u *unstructured.Resource) (*dclService.Backup, error) {
	r := &dclService.Backup{}
	if _, ok := u.Object["capacityGb"]; ok {
		if i, ok := u.Object["capacityGb"].(int64); ok {
			r.CapacityGb = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.CapacityGb: expected int64")
		}
	}
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
	if _, ok := u.Object["downloadBytes"]; ok {
		if i, ok := u.Object["downloadBytes"].(int64); ok {
			r.DownloadBytes = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.DownloadBytes: expected int64")
		}
	}
	if _, ok := u.Object["labels"]; ok {
		if rLabels, ok := u.Object["labels"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rLabels {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.Labels = m
		} else {
			return nil, fmt.Errorf("r.Labels: expected map[string]interface{}")
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
	if _, ok := u.Object["sourceFileShare"]; ok {
		if s, ok := u.Object["sourceFileShare"].(string); ok {
			r.SourceFileShare = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.SourceFileShare: expected string")
		}
	}
	if _, ok := u.Object["sourceInstance"]; ok {
		if s, ok := u.Object["sourceInstance"].(string); ok {
			r.SourceInstance = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.SourceInstance: expected string")
		}
	}
	if _, ok := u.Object["sourceInstanceTier"]; ok {
		if s, ok := u.Object["sourceInstanceTier"].(string); ok {
			r.SourceInstanceTier = dclService.BackupSourceInstanceTierEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.SourceInstanceTier: expected string")
		}
	}
	if _, ok := u.Object["state"]; ok {
		if s, ok := u.Object["state"].(string); ok {
			r.State = dclService.BackupStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.State: expected string")
		}
	}
	if _, ok := u.Object["storageBytes"]; ok {
		if i, ok := u.Object["storageBytes"].(int64); ok {
			r.StorageBytes = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.StorageBytes: expected int64")
		}
	}
	return r, nil
}

func GetBackup(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToBackup(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetBackup(ctx, r)
	if err != nil {
		return nil, err
	}
	return BackupToUnstructured(r), nil
}

func ListBackup(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListBackup(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, BackupToUnstructured(r))
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

func ApplyBackup(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToBackup(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToBackup(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyBackup(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return BackupToUnstructured(r), nil
}

func BackupHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToBackup(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToBackup(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyBackup(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteBackup(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToBackup(u)
	if err != nil {
		return err
	}
	return c.DeleteBackup(ctx, r)
}

func BackupID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToBackup(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Backup) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"filestore",
		"Backup",
		"beta",
	}
}

func (r *Backup) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Backup) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Backup) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Backup) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Backup) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Backup) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Backup) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetBackup(ctx, config, resource)
}

func (r *Backup) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyBackup(ctx, config, resource, opts...)
}

func (r *Backup) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return BackupHasDiff(ctx, config, resource, opts...)
}

func (r *Backup) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteBackup(ctx, config, resource)
}

func (r *Backup) ID(resource *unstructured.Resource) (string, error) {
	return BackupID(resource)
}

func init() {
	unstructured.Register(&Backup{})
}
