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
package dataplex

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataplex/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Lake struct{}

func LakeToUnstructured(r *dclService.Lake) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "dataplex",
			Version: "beta",
			Type:    "Lake",
		},
		Object: make(map[string]interface{}),
	}
	if r.AssetStatus != nil && r.AssetStatus != dclService.EmptyLakeAssetStatus {
		rAssetStatus := make(map[string]interface{})
		if r.AssetStatus.ActiveAssets != nil {
			rAssetStatus["activeAssets"] = *r.AssetStatus.ActiveAssets
		}
		if r.AssetStatus.SecurityPolicyApplyingAssets != nil {
			rAssetStatus["securityPolicyApplyingAssets"] = *r.AssetStatus.SecurityPolicyApplyingAssets
		}
		if r.AssetStatus.UpdateTime != nil {
			rAssetStatus["updateTime"] = *r.AssetStatus.UpdateTime
		}
		u.Object["assetStatus"] = rAssetStatus
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
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
	if r.Metastore != nil && r.Metastore != dclService.EmptyLakeMetastore {
		rMetastore := make(map[string]interface{})
		if r.Metastore.Service != nil {
			rMetastore["service"] = *r.Metastore.Service
		}
		u.Object["metastore"] = rMetastore
	}
	if r.MetastoreStatus != nil && r.MetastoreStatus != dclService.EmptyLakeMetastoreStatus {
		rMetastoreStatus := make(map[string]interface{})
		if r.MetastoreStatus.Endpoint != nil {
			rMetastoreStatus["endpoint"] = *r.MetastoreStatus.Endpoint
		}
		if r.MetastoreStatus.Message != nil {
			rMetastoreStatus["message"] = *r.MetastoreStatus.Message
		}
		if r.MetastoreStatus.State != nil {
			rMetastoreStatus["state"] = string(*r.MetastoreStatus.State)
		}
		if r.MetastoreStatus.UpdateTime != nil {
			rMetastoreStatus["updateTime"] = *r.MetastoreStatus.UpdateTime
		}
		u.Object["metastoreStatus"] = rMetastoreStatus
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.ServiceAccount != nil {
		u.Object["serviceAccount"] = *r.ServiceAccount
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
	return u
}

func UnstructuredToLake(u *unstructured.Resource) (*dclService.Lake, error) {
	r := &dclService.Lake{}
	if _, ok := u.Object["assetStatus"]; ok {
		if rAssetStatus, ok := u.Object["assetStatus"].(map[string]interface{}); ok {
			r.AssetStatus = &dclService.LakeAssetStatus{}
			if _, ok := rAssetStatus["activeAssets"]; ok {
				if i, ok := rAssetStatus["activeAssets"].(int64); ok {
					r.AssetStatus.ActiveAssets = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.AssetStatus.ActiveAssets: expected int64")
				}
			}
			if _, ok := rAssetStatus["securityPolicyApplyingAssets"]; ok {
				if i, ok := rAssetStatus["securityPolicyApplyingAssets"].(int64); ok {
					r.AssetStatus.SecurityPolicyApplyingAssets = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.AssetStatus.SecurityPolicyApplyingAssets: expected int64")
				}
			}
			if _, ok := rAssetStatus["updateTime"]; ok {
				if s, ok := rAssetStatus["updateTime"].(string); ok {
					r.AssetStatus.UpdateTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.AssetStatus.UpdateTime: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.AssetStatus: expected map[string]interface{}")
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
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
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
	if _, ok := u.Object["metastore"]; ok {
		if rMetastore, ok := u.Object["metastore"].(map[string]interface{}); ok {
			r.Metastore = &dclService.LakeMetastore{}
			if _, ok := rMetastore["service"]; ok {
				if s, ok := rMetastore["service"].(string); ok {
					r.Metastore.Service = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Metastore.Service: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Metastore: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["metastoreStatus"]; ok {
		if rMetastoreStatus, ok := u.Object["metastoreStatus"].(map[string]interface{}); ok {
			r.MetastoreStatus = &dclService.LakeMetastoreStatus{}
			if _, ok := rMetastoreStatus["endpoint"]; ok {
				if s, ok := rMetastoreStatus["endpoint"].(string); ok {
					r.MetastoreStatus.Endpoint = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.MetastoreStatus.Endpoint: expected string")
				}
			}
			if _, ok := rMetastoreStatus["message"]; ok {
				if s, ok := rMetastoreStatus["message"].(string); ok {
					r.MetastoreStatus.Message = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.MetastoreStatus.Message: expected string")
				}
			}
			if _, ok := rMetastoreStatus["state"]; ok {
				if s, ok := rMetastoreStatus["state"].(string); ok {
					r.MetastoreStatus.State = dclService.LakeMetastoreStatusStateEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.MetastoreStatus.State: expected string")
				}
			}
			if _, ok := rMetastoreStatus["updateTime"]; ok {
				if s, ok := rMetastoreStatus["updateTime"].(string); ok {
					r.MetastoreStatus.UpdateTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.MetastoreStatus.UpdateTime: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.MetastoreStatus: expected map[string]interface{}")
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
	if _, ok := u.Object["serviceAccount"]; ok {
		if s, ok := u.Object["serviceAccount"].(string); ok {
			r.ServiceAccount = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ServiceAccount: expected string")
		}
	}
	if _, ok := u.Object["state"]; ok {
		if s, ok := u.Object["state"].(string); ok {
			r.State = dclService.LakeStateEnumRef(s)
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
	return r, nil
}

func GetLake(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToLake(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetLake(ctx, r)
	if err != nil {
		return nil, err
	}
	return LakeToUnstructured(r), nil
}

func ListLake(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListLake(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, LakeToUnstructured(r))
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

func ApplyLake(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToLake(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToLake(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyLake(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return LakeToUnstructured(r), nil
}

func LakeHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToLake(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToLake(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyLake(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteLake(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToLake(u)
	if err != nil {
		return err
	}
	return c.DeleteLake(ctx, r)
}

func LakeID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToLake(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Lake) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"dataplex",
		"Lake",
		"beta",
	}
}

func (r *Lake) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Lake) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Lake) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Lake) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Lake) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Lake) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Lake) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetLake(ctx, config, resource)
}

func (r *Lake) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyLake(ctx, config, resource, opts...)
}

func (r *Lake) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return LakeHasDiff(ctx, config, resource, opts...)
}

func (r *Lake) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteLake(ctx, config, resource)
}

func (r *Lake) ID(resource *unstructured.Resource) (string, error) {
	return LakeID(resource)
}

func init() {
	unstructured.Register(&Lake{})
}
