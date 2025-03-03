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
package cloudresourcemanager

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudresourcemanager/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	iamUnstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/iam"
)

type TagKey struct{}

func TagKeyToUnstructured(r *dclService.TagKey) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "cloudresourcemanager",
			Version: "alpha",
			Type:    "TagKey",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.Etag != nil {
		u.Object["etag"] = *r.Etag
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.NamespacedName != nil {
		u.Object["namespacedName"] = *r.NamespacedName
	}
	if r.Parent != nil {
		u.Object["parent"] = *r.Parent
	}
	if r.Purpose != nil {
		u.Object["purpose"] = string(*r.Purpose)
	}
	if r.PurposeData != nil {
		rPurposeData := make(map[string]interface{})
		for k, v := range r.PurposeData {
			rPurposeData[k] = v
		}
		u.Object["purposeData"] = rPurposeData
	}
	if r.ShortName != nil {
		u.Object["shortName"] = *r.ShortName
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToTagKey(u *unstructured.Resource) (*dclService.TagKey, error) {
	r := &dclService.TagKey{}
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
	if _, ok := u.Object["etag"]; ok {
		if s, ok := u.Object["etag"].(string); ok {
			r.Etag = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Etag: expected string")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["namespacedName"]; ok {
		if s, ok := u.Object["namespacedName"].(string); ok {
			r.NamespacedName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.NamespacedName: expected string")
		}
	}
	if _, ok := u.Object["parent"]; ok {
		if s, ok := u.Object["parent"].(string); ok {
			r.Parent = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Parent: expected string")
		}
	}
	if _, ok := u.Object["purpose"]; ok {
		if s, ok := u.Object["purpose"].(string); ok {
			r.Purpose = dclService.TagKeyPurposeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Purpose: expected string")
		}
	}
	if _, ok := u.Object["purposeData"]; ok {
		if rPurposeData, ok := u.Object["purposeData"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rPurposeData {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.PurposeData = m
		} else {
			return nil, fmt.Errorf("r.PurposeData: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["shortName"]; ok {
		if s, ok := u.Object["shortName"].(string); ok {
			r.ShortName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ShortName: expected string")
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

func GetTagKey(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTagKey(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetTagKey(ctx, r)
	if err != nil {
		return nil, err
	}
	return TagKeyToUnstructured(r), nil
}

func ApplyTagKey(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTagKey(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToTagKey(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyTagKey(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return TagKeyToUnstructured(r), nil
}

func TagKeyHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTagKey(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToTagKey(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyTagKey(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteTagKey(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTagKey(u)
	if err != nil {
		return err
	}
	return c.DeleteTagKey(ctx, r)
}

func TagKeyID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToTagKey(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *TagKey) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"cloudresourcemanager",
		"TagKey",
		"alpha",
	}
}

func SetPolicyTagKey(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToTagKey(u)
	if err != nil {
		return nil, err
	}
	policy, err := iamUnstruct.UnstructuredToPolicy(p)
	if err != nil {
		return nil, err
	}
	policy.Resource = r
	iamClient := iam.NewClient(config)
	newPolicy, err := iamClient.SetPolicy(ctx, policy)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(newPolicy), nil
}

func SetPolicyWithEtagTagKey(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToTagKey(u)
	if err != nil {
		return nil, err
	}
	policy, err := iamUnstruct.UnstructuredToPolicy(p)
	if err != nil {
		return nil, err
	}
	policy.Resource = r
	iamClient := iam.NewClient(config)
	newPolicy, err := iamClient.SetPolicyWithEtag(ctx, policy)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(newPolicy), nil
}

func GetPolicyTagKey(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToTagKey(u)
	if err != nil {
		return nil, err
	}
	iamClient := iam.NewClient(config)
	policy, err := iamClient.GetPolicy(ctx, r)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(policy), nil
}

func SetPolicyMemberTagKey(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToTagKey(u)
	if err != nil {
		return nil, err
	}
	member, err := iamUnstruct.UnstructuredToMember(m)
	if err != nil {
		return nil, err
	}
	member.Resource = r
	iamClient := iam.NewClient(config)
	policy, err := iamClient.SetMember(ctx, member)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(policy), nil
}

func GetPolicyMemberTagKey(ctx context.Context, config *dcl.Config, u *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	r, err := UnstructuredToTagKey(u)
	if err != nil {
		return nil, err
	}
	iamClient := iam.NewClient(config)
	policyMember, err := iamClient.GetMember(ctx, r, role, member)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.MemberToUnstructured(policyMember), nil
}

func DeletePolicyMemberTagKey(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) error {
	r, err := UnstructuredToTagKey(u)
	if err != nil {
		return err
	}
	member, err := iamUnstruct.UnstructuredToMember(m)
	if err != nil {
		return err
	}
	member.Resource = r
	iamClient := iam.NewClient(config)
	if err := iamClient.DeleteMember(ctx, member); err != nil {
		return err
	}
	return nil
}

func (r *TagKey) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyMemberTagKey(ctx, config, resource, member)
}

func (r *TagKey) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return GetPolicyMemberTagKey(ctx, config, resource, role, member)
}

func (r *TagKey) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return DeletePolicyMemberTagKey(ctx, config, resource, member)
}

func (r *TagKey) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyTagKey(ctx, config, resource, policy)
}

func (r *TagKey) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyWithEtagTagKey(ctx, config, resource, policy)
}

func (r *TagKey) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetPolicyTagKey(ctx, config, resource)
}

func (r *TagKey) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetTagKey(ctx, config, resource)
}

func (r *TagKey) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyTagKey(ctx, config, resource, opts...)
}

func (r *TagKey) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return TagKeyHasDiff(ctx, config, resource, opts...)
}

func (r *TagKey) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteTagKey(ctx, config, resource)
}

func (r *TagKey) ID(resource *unstructured.Resource) (string, error) {
	return TagKeyID(resource)
}

func init() {
	unstructured.Register(&TagKey{})
}
