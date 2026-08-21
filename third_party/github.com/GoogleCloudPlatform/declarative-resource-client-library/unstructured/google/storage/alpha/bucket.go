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
package storage

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/storage/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	iamUnstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/iam"
)

type Bucket struct{}

func BucketToUnstructured(r *dclService.Bucket) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "storage",
			Version: "alpha",
			Type:    "Bucket",
		},
		Object: make(map[string]interface{}),
	}
	var rCors []interface{}
	for _, rCorsVal := range r.Cors {
		rCorsObject := make(map[string]interface{})
		if rCorsVal.MaxAgeSeconds != nil {
			rCorsObject["maxAgeSeconds"] = *rCorsVal.MaxAgeSeconds
		}
		var rCorsValMethod []interface{}
		for _, rCorsValMethodVal := range rCorsVal.Method {
			rCorsValMethod = append(rCorsValMethod, rCorsValMethodVal)
		}
		rCorsObject["method"] = rCorsValMethod
		var rCorsValOrigin []interface{}
		for _, rCorsValOriginVal := range rCorsVal.Origin {
			rCorsValOrigin = append(rCorsValOrigin, rCorsValOriginVal)
		}
		rCorsObject["origin"] = rCorsValOrigin
		var rCorsValResponseHeader []interface{}
		for _, rCorsValResponseHeaderVal := range rCorsVal.ResponseHeader {
			rCorsValResponseHeader = append(rCorsValResponseHeader, rCorsValResponseHeaderVal)
		}
		rCorsObject["responseHeader"] = rCorsValResponseHeader
		rCors = append(rCors, rCorsObject)
	}
	u.Object["cors"] = rCors
	if r.Lifecycle != nil && r.Lifecycle != dclService.EmptyBucketLifecycle {
		rLifecycle := make(map[string]interface{})
		var rLifecycleRule []interface{}
		for _, rLifecycleRuleVal := range r.Lifecycle.Rule {
			rLifecycleRuleObject := make(map[string]interface{})
			if rLifecycleRuleVal.Action != nil && rLifecycleRuleVal.Action != dclService.EmptyBucketLifecycleRuleAction {
				rLifecycleRuleValAction := make(map[string]interface{})
				if rLifecycleRuleVal.Action.StorageClass != nil {
					rLifecycleRuleValAction["storageClass"] = *rLifecycleRuleVal.Action.StorageClass
				}
				if rLifecycleRuleVal.Action.Type != nil {
					rLifecycleRuleValAction["type"] = string(*rLifecycleRuleVal.Action.Type)
				}
				rLifecycleRuleObject["action"] = rLifecycleRuleValAction
			}
			if rLifecycleRuleVal.Condition != nil && rLifecycleRuleVal.Condition != dclService.EmptyBucketLifecycleRuleCondition {
				rLifecycleRuleValCondition := make(map[string]interface{})
				if rLifecycleRuleVal.Condition.Age != nil {
					rLifecycleRuleValCondition["age"] = *rLifecycleRuleVal.Condition.Age
				}
				if rLifecycleRuleVal.Condition.CreatedBefore != nil {
					rLifecycleRuleValCondition["createdBefore"] = *rLifecycleRuleVal.Condition.CreatedBefore
				}
				var rLifecycleRuleValConditionMatchesStorageClass []interface{}
				for _, rLifecycleRuleValConditionMatchesStorageClassVal := range rLifecycleRuleVal.Condition.MatchesStorageClass {
					rLifecycleRuleValConditionMatchesStorageClass = append(rLifecycleRuleValConditionMatchesStorageClass, rLifecycleRuleValConditionMatchesStorageClassVal)
				}
				rLifecycleRuleValCondition["matchesStorageClass"] = rLifecycleRuleValConditionMatchesStorageClass
				if rLifecycleRuleVal.Condition.NumNewerVersions != nil {
					rLifecycleRuleValCondition["numNewerVersions"] = *rLifecycleRuleVal.Condition.NumNewerVersions
				}
				if rLifecycleRuleVal.Condition.WithState != nil {
					rLifecycleRuleValCondition["withState"] = string(*rLifecycleRuleVal.Condition.WithState)
				}
				rLifecycleRuleObject["condition"] = rLifecycleRuleValCondition
			}
			rLifecycleRule = append(rLifecycleRule, rLifecycleRuleObject)
		}
		rLifecycle["rule"] = rLifecycleRule
		u.Object["lifecycle"] = rLifecycle
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Logging != nil && r.Logging != dclService.EmptyBucketLogging {
		rLogging := make(map[string]interface{})
		if r.Logging.LogBucket != nil {
			rLogging["logBucket"] = *r.Logging.LogBucket
		}
		if r.Logging.LogObjectPrefix != nil {
			rLogging["logObjectPrefix"] = *r.Logging.LogObjectPrefix
		}
		u.Object["logging"] = rLogging
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.StorageClass != nil {
		u.Object["storageClass"] = string(*r.StorageClass)
	}
	if r.Versioning != nil && r.Versioning != dclService.EmptyBucketVersioning {
		rVersioning := make(map[string]interface{})
		if r.Versioning.Enabled != nil {
			rVersioning["enabled"] = *r.Versioning.Enabled
		}
		u.Object["versioning"] = rVersioning
	}
	if r.Website != nil && r.Website != dclService.EmptyBucketWebsite {
		rWebsite := make(map[string]interface{})
		if r.Website.MainPageSuffix != nil {
			rWebsite["mainPageSuffix"] = *r.Website.MainPageSuffix
		}
		if r.Website.NotFoundPage != nil {
			rWebsite["notFoundPage"] = *r.Website.NotFoundPage
		}
		u.Object["website"] = rWebsite
	}
	return u
}

func UnstructuredToBucket(u *unstructured.Resource) (*dclService.Bucket, error) {
	r := &dclService.Bucket{}
	if _, ok := u.Object["cors"]; ok {
		if s, ok := u.Object["cors"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rCors dclService.BucketCors
					if _, ok := objval["maxAgeSeconds"]; ok {
						if i, ok := objval["maxAgeSeconds"].(int64); ok {
							rCors.MaxAgeSeconds = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("rCors.MaxAgeSeconds: expected int64")
						}
					}
					if _, ok := objval["method"]; ok {
						if s, ok := objval["method"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									rCors.Method = append(rCors.Method, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("rCors.Method: expected []interface{}")
						}
					}
					if _, ok := objval["origin"]; ok {
						if s, ok := objval["origin"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									rCors.Origin = append(rCors.Origin, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("rCors.Origin: expected []interface{}")
						}
					}
					if _, ok := objval["responseHeader"]; ok {
						if s, ok := objval["responseHeader"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									rCors.ResponseHeader = append(rCors.ResponseHeader, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("rCors.ResponseHeader: expected []interface{}")
						}
					}
					r.Cors = append(r.Cors, rCors)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Cors: expected []interface{}")
		}
	}
	if _, ok := u.Object["lifecycle"]; ok {
		if rLifecycle, ok := u.Object["lifecycle"].(map[string]interface{}); ok {
			r.Lifecycle = &dclService.BucketLifecycle{}
			if _, ok := rLifecycle["rule"]; ok {
				if s, ok := rLifecycle["rule"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rLifecycleRule dclService.BucketLifecycleRule
							if _, ok := objval["action"]; ok {
								if rLifecycleRuleAction, ok := objval["action"].(map[string]interface{}); ok {
									rLifecycleRule.Action = &dclService.BucketLifecycleRuleAction{}
									if _, ok := rLifecycleRuleAction["storageClass"]; ok {
										if s, ok := rLifecycleRuleAction["storageClass"].(string); ok {
											rLifecycleRule.Action.StorageClass = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rLifecycleRule.Action.StorageClass: expected string")
										}
									}
									if _, ok := rLifecycleRuleAction["type"]; ok {
										if s, ok := rLifecycleRuleAction["type"].(string); ok {
											rLifecycleRule.Action.Type = dclService.BucketLifecycleRuleActionTypeEnumRef(s)
										} else {
											return nil, fmt.Errorf("rLifecycleRule.Action.Type: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("rLifecycleRule.Action: expected map[string]interface{}")
								}
							}
							if _, ok := objval["condition"]; ok {
								if rLifecycleRuleCondition, ok := objval["condition"].(map[string]interface{}); ok {
									rLifecycleRule.Condition = &dclService.BucketLifecycleRuleCondition{}
									if _, ok := rLifecycleRuleCondition["age"]; ok {
										if i, ok := rLifecycleRuleCondition["age"].(int64); ok {
											rLifecycleRule.Condition.Age = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("rLifecycleRule.Condition.Age: expected int64")
										}
									}
									if _, ok := rLifecycleRuleCondition["createdBefore"]; ok {
										if s, ok := rLifecycleRuleCondition["createdBefore"].(string); ok {
											rLifecycleRule.Condition.CreatedBefore = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rLifecycleRule.Condition.CreatedBefore: expected string")
										}
									}
									if _, ok := rLifecycleRuleCondition["matchesStorageClass"]; ok {
										if s, ok := rLifecycleRuleCondition["matchesStorageClass"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rLifecycleRule.Condition.MatchesStorageClass = append(rLifecycleRule.Condition.MatchesStorageClass, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rLifecycleRule.Condition.MatchesStorageClass: expected []interface{}")
										}
									}
									if _, ok := rLifecycleRuleCondition["numNewerVersions"]; ok {
										if i, ok := rLifecycleRuleCondition["numNewerVersions"].(int64); ok {
											rLifecycleRule.Condition.NumNewerVersions = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("rLifecycleRule.Condition.NumNewerVersions: expected int64")
										}
									}
									if _, ok := rLifecycleRuleCondition["withState"]; ok {
										if s, ok := rLifecycleRuleCondition["withState"].(string); ok {
											rLifecycleRule.Condition.WithState = dclService.BucketLifecycleRuleConditionWithStateEnumRef(s)
										} else {
											return nil, fmt.Errorf("rLifecycleRule.Condition.WithState: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("rLifecycleRule.Condition: expected map[string]interface{}")
								}
							}
							r.Lifecycle.Rule = append(r.Lifecycle.Rule, rLifecycleRule)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Lifecycle.Rule: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Lifecycle: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["logging"]; ok {
		if rLogging, ok := u.Object["logging"].(map[string]interface{}); ok {
			r.Logging = &dclService.BucketLogging{}
			if _, ok := rLogging["logBucket"]; ok {
				if s, ok := rLogging["logBucket"].(string); ok {
					r.Logging.LogBucket = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Logging.LogBucket: expected string")
				}
			}
			if _, ok := rLogging["logObjectPrefix"]; ok {
				if s, ok := rLogging["logObjectPrefix"].(string); ok {
					r.Logging.LogObjectPrefix = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Logging.LogObjectPrefix: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Logging: expected map[string]interface{}")
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
	if _, ok := u.Object["storageClass"]; ok {
		if s, ok := u.Object["storageClass"].(string); ok {
			r.StorageClass = dclService.BucketStorageClassEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.StorageClass: expected string")
		}
	}
	if _, ok := u.Object["versioning"]; ok {
		if rVersioning, ok := u.Object["versioning"].(map[string]interface{}); ok {
			r.Versioning = &dclService.BucketVersioning{}
			if _, ok := rVersioning["enabled"]; ok {
				if b, ok := rVersioning["enabled"].(bool); ok {
					r.Versioning.Enabled = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.Versioning.Enabled: expected bool")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Versioning: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["website"]; ok {
		if rWebsite, ok := u.Object["website"].(map[string]interface{}); ok {
			r.Website = &dclService.BucketWebsite{}
			if _, ok := rWebsite["mainPageSuffix"]; ok {
				if s, ok := rWebsite["mainPageSuffix"].(string); ok {
					r.Website.MainPageSuffix = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Website.MainPageSuffix: expected string")
				}
			}
			if _, ok := rWebsite["notFoundPage"]; ok {
				if s, ok := rWebsite["notFoundPage"].(string); ok {
					r.Website.NotFoundPage = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Website.NotFoundPage: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Website: expected map[string]interface{}")
		}
	}
	return r, nil
}

func GetBucket(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToBucket(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetBucket(ctx, r)
	if err != nil {
		return nil, err
	}
	return BucketToUnstructured(r), nil
}

func ListBucket(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListBucket(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, BucketToUnstructured(r))
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

func ApplyBucket(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToBucket(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToBucket(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyBucket(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return BucketToUnstructured(r), nil
}

func BucketHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToBucket(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToBucket(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyBucket(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteBucket(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToBucket(u)
	if err != nil {
		return err
	}
	return c.DeleteBucket(ctx, r)
}

func BucketID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToBucket(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Bucket) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"storage",
		"Bucket",
		"alpha",
	}
}

func SetPolicyBucket(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToBucket(u)
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

func SetPolicyWithEtagBucket(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToBucket(u)
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

func GetPolicyBucket(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToBucket(u)
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

func SetPolicyMemberBucket(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToBucket(u)
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

func GetPolicyMemberBucket(ctx context.Context, config *dcl.Config, u *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	r, err := UnstructuredToBucket(u)
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

func DeletePolicyMemberBucket(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) error {
	r, err := UnstructuredToBucket(u)
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

func (r *Bucket) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyMemberBucket(ctx, config, resource, member)
}

func (r *Bucket) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return GetPolicyMemberBucket(ctx, config, resource, role, member)
}

func (r *Bucket) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return DeletePolicyMemberBucket(ctx, config, resource, member)
}

func (r *Bucket) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyBucket(ctx, config, resource, policy)
}

func (r *Bucket) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyWithEtagBucket(ctx, config, resource, policy)
}

func (r *Bucket) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetPolicyBucket(ctx, config, resource)
}

func (r *Bucket) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetBucket(ctx, config, resource)
}

func (r *Bucket) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyBucket(ctx, config, resource, opts...)
}

func (r *Bucket) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return BucketHasDiff(ctx, config, resource, opts...)
}

func (r *Bucket) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteBucket(ctx, config, resource)
}

func (r *Bucket) ID(resource *unstructured.Resource) (string, error) {
	return BucketID(resource)
}

func init() {
	unstructured.Register(&Bucket{})
}
