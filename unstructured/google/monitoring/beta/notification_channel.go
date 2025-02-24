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
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type NotificationChannel struct{}

func NotificationChannelToUnstructured(r *dclService.NotificationChannel) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "monitoring",
			Version: "beta",
			Type:    "NotificationChannel",
		},
		Object: make(map[string]interface{}),
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.Enabled != nil {
		u.Object["enabled"] = *r.Enabled
	}
	if r.Labels != nil {
		rLabels := make(map[string]interface{})
		for k, v := range r.Labels {
			rLabels[k] = v
		}
		u.Object["labels"] = rLabels
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.Type != nil {
		u.Object["type"] = *r.Type
	}
	if r.UserLabels != nil {
		rUserLabels := make(map[string]interface{})
		for k, v := range r.UserLabels {
			rUserLabels[k] = v
		}
		u.Object["userLabels"] = rUserLabels
	}
	if r.VerificationStatus != nil {
		u.Object["verificationStatus"] = string(*r.VerificationStatus)
	}
	return u
}

func UnstructuredToNotificationChannel(u *unstructured.Resource) (*dclService.NotificationChannel, error) {
	r := &dclService.NotificationChannel{}
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
	if _, ok := u.Object["enabled"]; ok {
		if b, ok := u.Object["enabled"].(bool); ok {
			r.Enabled = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.Enabled: expected bool")
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
	if _, ok := u.Object["type"]; ok {
		if s, ok := u.Object["type"].(string); ok {
			r.Type = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Type: expected string")
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
	if _, ok := u.Object["verificationStatus"]; ok {
		if s, ok := u.Object["verificationStatus"].(string); ok {
			r.VerificationStatus = dclService.NotificationChannelVerificationStatusEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.VerificationStatus: expected string")
		}
	}
	return r, nil
}

func GetNotificationChannel(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNotificationChannel(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetNotificationChannel(ctx, r)
	if err != nil {
		return nil, err
	}
	return NotificationChannelToUnstructured(r), nil
}

func ListNotificationChannel(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListNotificationChannel(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, NotificationChannelToUnstructured(r))
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

func ApplyNotificationChannel(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNotificationChannel(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToNotificationChannel(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyNotificationChannel(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return NotificationChannelToUnstructured(r), nil
}

func NotificationChannelHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNotificationChannel(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToNotificationChannel(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyNotificationChannel(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteNotificationChannel(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNotificationChannel(u)
	if err != nil {
		return err
	}
	return c.DeleteNotificationChannel(ctx, r)
}

func NotificationChannelID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToNotificationChannel(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *NotificationChannel) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"monitoring",
		"NotificationChannel",
		"beta",
	}
}

func (r *NotificationChannel) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *NotificationChannel) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *NotificationChannel) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *NotificationChannel) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *NotificationChannel) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *NotificationChannel) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *NotificationChannel) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetNotificationChannel(ctx, config, resource)
}

func (r *NotificationChannel) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyNotificationChannel(ctx, config, resource, opts...)
}

func (r *NotificationChannel) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return NotificationChannelHasDiff(ctx, config, resource, opts...)
}

func (r *NotificationChannel) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteNotificationChannel(ctx, config, resource)
}

func (r *NotificationChannel) ID(resource *unstructured.Resource) (string, error) {
	return NotificationChannelID(resource)
}

func init() {
	unstructured.Register(&NotificationChannel{})
}
