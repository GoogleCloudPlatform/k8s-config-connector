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
package pubsub

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/pubsub"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Topic struct{}

func TopicToUnstructured(r *dclService.Topic) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "pubsub",
			Version: "ga",
			Type:    "Topic",
		},
		Object: make(map[string]interface{}),
	}
	if r.KmsKeyName != nil {
		u.Object["kmsKeyName"] = *r.KmsKeyName
	}
	if r.Labels != nil {
		rLabels := make(map[string]interface{})
		for k, v := range r.Labels {
			rLabels[k] = v
		}
		u.Object["labels"] = rLabels
	}
	if r.MessageStoragePolicy != nil && r.MessageStoragePolicy != dclService.EmptyTopicMessageStoragePolicy {
		rMessageStoragePolicy := make(map[string]interface{})
		var rMessageStoragePolicyAllowedPersistenceRegions []interface{}
		for _, rMessageStoragePolicyAllowedPersistenceRegionsVal := range r.MessageStoragePolicy.AllowedPersistenceRegions {
			rMessageStoragePolicyAllowedPersistenceRegions = append(rMessageStoragePolicyAllowedPersistenceRegions, rMessageStoragePolicyAllowedPersistenceRegionsVal)
		}
		rMessageStoragePolicy["allowedPersistenceRegions"] = rMessageStoragePolicyAllowedPersistenceRegions
		u.Object["messageStoragePolicy"] = rMessageStoragePolicy
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	return u
}

func UnstructuredToTopic(u *unstructured.Resource) (*dclService.Topic, error) {
	r := &dclService.Topic{}
	if _, ok := u.Object["kmsKeyName"]; ok {
		if s, ok := u.Object["kmsKeyName"].(string); ok {
			r.KmsKeyName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.KmsKeyName: expected string")
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
	if _, ok := u.Object["messageStoragePolicy"]; ok {
		if rMessageStoragePolicy, ok := u.Object["messageStoragePolicy"].(map[string]interface{}); ok {
			r.MessageStoragePolicy = &dclService.TopicMessageStoragePolicy{}
			if _, ok := rMessageStoragePolicy["allowedPersistenceRegions"]; ok {
				if s, ok := rMessageStoragePolicy["allowedPersistenceRegions"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.MessageStoragePolicy.AllowedPersistenceRegions = append(r.MessageStoragePolicy.AllowedPersistenceRegions, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.MessageStoragePolicy.AllowedPersistenceRegions: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.MessageStoragePolicy: expected map[string]interface{}")
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
	return r, nil
}

func GetTopic(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTopic(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetTopic(ctx, r)
	if err != nil {
		return nil, err
	}
	return TopicToUnstructured(r), nil
}

func ListTopic(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListTopic(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, TopicToUnstructured(r))
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

func ApplyTopic(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTopic(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToTopic(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyTopic(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return TopicToUnstructured(r), nil
}

func TopicHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTopic(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToTopic(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyTopic(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteTopic(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTopic(u)
	if err != nil {
		return err
	}
	return c.DeleteTopic(ctx, r)
}

func TopicID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToTopic(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Topic) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"pubsub",
		"Topic",
		"ga",
	}
}

func (r *Topic) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Topic) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Topic) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Topic) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Topic) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Topic) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Topic) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetTopic(ctx, config, resource)
}

func (r *Topic) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyTopic(ctx, config, resource, opts...)
}

func (r *Topic) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return TopicHasDiff(ctx, config, resource, opts...)
}

func (r *Topic) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteTopic(ctx, config, resource)
}

func (r *Topic) ID(resource *unstructured.Resource) (string, error) {
	return TopicID(resource)
}

func init() {
	unstructured.Register(&Topic{})
}
