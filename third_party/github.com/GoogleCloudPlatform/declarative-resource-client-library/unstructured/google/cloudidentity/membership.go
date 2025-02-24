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
package cloudidentity

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudidentity"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Membership struct{}

func MembershipToUnstructured(r *dclService.Membership) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "cloudidentity",
			Version: "ga",
			Type:    "Membership",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.DeliverySetting != nil {
		u.Object["deliverySetting"] = string(*r.DeliverySetting)
	}
	if r.DisplayName != nil && r.DisplayName != dclService.EmptyMembershipDisplayName {
		rDisplayName := make(map[string]interface{})
		if r.DisplayName.FamilyName != nil {
			rDisplayName["familyName"] = *r.DisplayName.FamilyName
		}
		if r.DisplayName.FullName != nil {
			rDisplayName["fullName"] = *r.DisplayName.FullName
		}
		if r.DisplayName.GivenName != nil {
			rDisplayName["givenName"] = *r.DisplayName.GivenName
		}
		u.Object["displayName"] = rDisplayName
	}
	if r.Group != nil {
		u.Object["group"] = *r.Group
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.PreferredMemberKey != nil && r.PreferredMemberKey != dclService.EmptyMembershipPreferredMemberKey {
		rPreferredMemberKey := make(map[string]interface{})
		if r.PreferredMemberKey.Id != nil {
			rPreferredMemberKey["id"] = *r.PreferredMemberKey.Id
		}
		if r.PreferredMemberKey.Namespace != nil {
			rPreferredMemberKey["namespace"] = *r.PreferredMemberKey.Namespace
		}
		u.Object["preferredMemberKey"] = rPreferredMemberKey
	}
	var rRoles []interface{}
	for _, rRolesVal := range r.Roles {
		rRolesObject := make(map[string]interface{})
		if rRolesVal.ExpiryDetail != nil && rRolesVal.ExpiryDetail != dclService.EmptyMembershipRolesExpiryDetail {
			rRolesValExpiryDetail := make(map[string]interface{})
			if rRolesVal.ExpiryDetail.ExpireTime != nil {
				rRolesValExpiryDetail["expireTime"] = *rRolesVal.ExpiryDetail.ExpireTime
			}
			rRolesObject["expiryDetail"] = rRolesValExpiryDetail
		}
		if rRolesVal.Name != nil {
			rRolesObject["name"] = *rRolesVal.Name
		}
		if rRolesVal.RestrictionEvaluations != nil && rRolesVal.RestrictionEvaluations != dclService.EmptyMembershipRolesRestrictionEvaluations {
			rRolesValRestrictionEvaluations := make(map[string]interface{})
			if rRolesVal.RestrictionEvaluations.MemberRestrictionEvaluation != nil && rRolesVal.RestrictionEvaluations.MemberRestrictionEvaluation != dclService.EmptyMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation {
				rRolesValRestrictionEvaluationsMemberRestrictionEvaluation := make(map[string]interface{})
				if rRolesVal.RestrictionEvaluations.MemberRestrictionEvaluation.State != nil {
					rRolesValRestrictionEvaluationsMemberRestrictionEvaluation["state"] = string(*rRolesVal.RestrictionEvaluations.MemberRestrictionEvaluation.State)
				}
				rRolesValRestrictionEvaluations["memberRestrictionEvaluation"] = rRolesValRestrictionEvaluationsMemberRestrictionEvaluation
			}
			rRolesObject["restrictionEvaluations"] = rRolesValRestrictionEvaluations
		}
		rRoles = append(rRoles, rRolesObject)
	}
	u.Object["roles"] = rRoles
	if r.Type != nil {
		u.Object["type"] = string(*r.Type)
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToMembership(u *unstructured.Resource) (*dclService.Membership, error) {
	r := &dclService.Membership{}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["deliverySetting"]; ok {
		if s, ok := u.Object["deliverySetting"].(string); ok {
			r.DeliverySetting = dclService.MembershipDeliverySettingEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.DeliverySetting: expected string")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if rDisplayName, ok := u.Object["displayName"].(map[string]interface{}); ok {
			r.DisplayName = &dclService.MembershipDisplayName{}
			if _, ok := rDisplayName["familyName"]; ok {
				if s, ok := rDisplayName["familyName"].(string); ok {
					r.DisplayName.FamilyName = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.DisplayName.FamilyName: expected string")
				}
			}
			if _, ok := rDisplayName["fullName"]; ok {
				if s, ok := rDisplayName["fullName"].(string); ok {
					r.DisplayName.FullName = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.DisplayName.FullName: expected string")
				}
			}
			if _, ok := rDisplayName["givenName"]; ok {
				if s, ok := rDisplayName["givenName"].(string); ok {
					r.DisplayName.GivenName = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.DisplayName.GivenName: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["group"]; ok {
		if s, ok := u.Object["group"].(string); ok {
			r.Group = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Group: expected string")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["preferredMemberKey"]; ok {
		if rPreferredMemberKey, ok := u.Object["preferredMemberKey"].(map[string]interface{}); ok {
			r.PreferredMemberKey = &dclService.MembershipPreferredMemberKey{}
			if _, ok := rPreferredMemberKey["id"]; ok {
				if s, ok := rPreferredMemberKey["id"].(string); ok {
					r.PreferredMemberKey.Id = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.PreferredMemberKey.Id: expected string")
				}
			}
			if _, ok := rPreferredMemberKey["namespace"]; ok {
				if s, ok := rPreferredMemberKey["namespace"].(string); ok {
					r.PreferredMemberKey.Namespace = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.PreferredMemberKey.Namespace: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.PreferredMemberKey: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["roles"]; ok {
		if s, ok := u.Object["roles"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rRoles dclService.MembershipRoles
					if _, ok := objval["expiryDetail"]; ok {
						if rRolesExpiryDetail, ok := objval["expiryDetail"].(map[string]interface{}); ok {
							rRoles.ExpiryDetail = &dclService.MembershipRolesExpiryDetail{}
							if _, ok := rRolesExpiryDetail["expireTime"]; ok {
								if s, ok := rRolesExpiryDetail["expireTime"].(string); ok {
									rRoles.ExpiryDetail.ExpireTime = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rRoles.ExpiryDetail.ExpireTime: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("rRoles.ExpiryDetail: expected map[string]interface{}")
						}
					}
					if _, ok := objval["name"]; ok {
						if s, ok := objval["name"].(string); ok {
							rRoles.Name = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rRoles.Name: expected string")
						}
					}
					if _, ok := objval["restrictionEvaluations"]; ok {
						if rRolesRestrictionEvaluations, ok := objval["restrictionEvaluations"].(map[string]interface{}); ok {
							rRoles.RestrictionEvaluations = &dclService.MembershipRolesRestrictionEvaluations{}
							if _, ok := rRolesRestrictionEvaluations["memberRestrictionEvaluation"]; ok {
								if rRolesRestrictionEvaluationsMemberRestrictionEvaluation, ok := rRolesRestrictionEvaluations["memberRestrictionEvaluation"].(map[string]interface{}); ok {
									rRoles.RestrictionEvaluations.MemberRestrictionEvaluation = &dclService.MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation{}
									if _, ok := rRolesRestrictionEvaluationsMemberRestrictionEvaluation["state"]; ok {
										if s, ok := rRolesRestrictionEvaluationsMemberRestrictionEvaluation["state"].(string); ok {
											rRoles.RestrictionEvaluations.MemberRestrictionEvaluation.State = dclService.MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnumRef(s)
										} else {
											return nil, fmt.Errorf("rRoles.RestrictionEvaluations.MemberRestrictionEvaluation.State: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("rRoles.RestrictionEvaluations.MemberRestrictionEvaluation: expected map[string]interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("rRoles.RestrictionEvaluations: expected map[string]interface{}")
						}
					}
					r.Roles = append(r.Roles, rRoles)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Roles: expected []interface{}")
		}
	}
	if _, ok := u.Object["type"]; ok {
		if s, ok := u.Object["type"].(string); ok {
			r.Type = dclService.MembershipTypeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Type: expected string")
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

func GetMembership(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToMembership(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetMembership(ctx, r)
	if err != nil {
		return nil, err
	}
	return MembershipToUnstructured(r), nil
}

func ListMembership(ctx context.Context, config *dcl.Config, group string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListMembership(ctx, group)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, MembershipToUnstructured(r))
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

func ApplyMembership(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToMembership(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToMembership(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyMembership(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return MembershipToUnstructured(r), nil
}

func MembershipHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToMembership(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToMembership(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyMembership(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteMembership(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToMembership(u)
	if err != nil {
		return err
	}
	return c.DeleteMembership(ctx, r)
}

func MembershipID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToMembership(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Membership) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"cloudidentity",
		"Membership",
		"ga",
	}
}

func (r *Membership) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Membership) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Membership) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Membership) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Membership) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Membership) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Membership) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetMembership(ctx, config, resource)
}

func (r *Membership) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyMembership(ctx, config, resource, opts...)
}

func (r *Membership) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return MembershipHasDiff(ctx, config, resource, opts...)
}

func (r *Membership) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteMembership(ctx, config, resource)
}

func (r *Membership) ID(resource *unstructured.Resource) (string, error) {
	return MembershipID(resource)
}

func init() {
	unstructured.Register(&Membership{})
}
