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
package firebaserules

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/firebaserules/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Ruleset struct{}

func RulesetToUnstructured(r *dclService.Ruleset) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "firebaserules",
			Version: "beta",
			Type:    "Ruleset",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Metadata != nil && r.Metadata != dclService.EmptyRulesetMetadata {
		rMetadata := make(map[string]interface{})
		var rMetadataServices []interface{}
		for _, rMetadataServicesVal := range r.Metadata.Services {
			rMetadataServices = append(rMetadataServices, rMetadataServicesVal)
		}
		rMetadata["services"] = rMetadataServices
		u.Object["metadata"] = rMetadata
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.Source != nil && r.Source != dclService.EmptyRulesetSource {
		rSource := make(map[string]interface{})
		var rSourceFiles []interface{}
		for _, rSourceFilesVal := range r.Source.Files {
			rSourceFilesObject := make(map[string]interface{})
			if rSourceFilesVal.Content != nil {
				rSourceFilesObject["content"] = *rSourceFilesVal.Content
			}
			if rSourceFilesVal.Fingerprint != nil {
				rSourceFilesObject["fingerprint"] = *rSourceFilesVal.Fingerprint
			}
			if rSourceFilesVal.Name != nil {
				rSourceFilesObject["name"] = *rSourceFilesVal.Name
			}
			rSourceFiles = append(rSourceFiles, rSourceFilesObject)
		}
		rSource["files"] = rSourceFiles
		if r.Source.Language != nil {
			rSource["language"] = string(*r.Source.Language)
		}
		u.Object["source"] = rSource
	}
	return u
}

func UnstructuredToRuleset(u *unstructured.Resource) (*dclService.Ruleset, error) {
	r := &dclService.Ruleset{}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["metadata"]; ok {
		if rMetadata, ok := u.Object["metadata"].(map[string]interface{}); ok {
			r.Metadata = &dclService.RulesetMetadata{}
			if _, ok := rMetadata["services"]; ok {
				if s, ok := rMetadata["services"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Metadata.Services = append(r.Metadata.Services, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Metadata.Services: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Metadata: expected map[string]interface{}")
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
	if _, ok := u.Object["source"]; ok {
		if rSource, ok := u.Object["source"].(map[string]interface{}); ok {
			r.Source = &dclService.RulesetSource{}
			if _, ok := rSource["files"]; ok {
				if s, ok := rSource["files"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rSourceFiles dclService.RulesetSourceFiles
							if _, ok := objval["content"]; ok {
								if s, ok := objval["content"].(string); ok {
									rSourceFiles.Content = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rSourceFiles.Content: expected string")
								}
							}
							if _, ok := objval["fingerprint"]; ok {
								if s, ok := objval["fingerprint"].(string); ok {
									rSourceFiles.Fingerprint = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rSourceFiles.Fingerprint: expected string")
								}
							}
							if _, ok := objval["name"]; ok {
								if s, ok := objval["name"].(string); ok {
									rSourceFiles.Name = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rSourceFiles.Name: expected string")
								}
							}
							r.Source.Files = append(r.Source.Files, rSourceFiles)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Source.Files: expected []interface{}")
				}
			}
			if _, ok := rSource["language"]; ok {
				if s, ok := rSource["language"].(string); ok {
					r.Source.Language = dclService.RulesetSourceLanguageEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.Source.Language: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Source: expected map[string]interface{}")
		}
	}
	return r, nil
}

func GetRuleset(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToRuleset(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetRuleset(ctx, r)
	if err != nil {
		return nil, err
	}
	return RulesetToUnstructured(r), nil
}

func ListRuleset(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListRuleset(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, RulesetToUnstructured(r))
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

func ApplyRuleset(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToRuleset(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToRuleset(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyRuleset(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return RulesetToUnstructured(r), nil
}

func RulesetHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToRuleset(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToRuleset(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyRuleset(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteRuleset(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToRuleset(u)
	if err != nil {
		return err
	}
	return c.DeleteRuleset(ctx, r)
}

func RulesetID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToRuleset(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Ruleset) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"firebaserules",
		"Ruleset",
		"beta",
	}
}

func (r *Ruleset) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Ruleset) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Ruleset) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Ruleset) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Ruleset) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Ruleset) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Ruleset) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetRuleset(ctx, config, resource)
}

func (r *Ruleset) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyRuleset(ctx, config, resource, opts...)
}

func (r *Ruleset) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return RulesetHasDiff(ctx, config, resource, opts...)
}

func (r *Ruleset) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteRuleset(ctx, config, resource)
}

func (r *Ruleset) ID(resource *unstructured.Resource) (string, error) {
	return RulesetID(resource)
}

func init() {
	unstructured.Register(&Ruleset{})
}
