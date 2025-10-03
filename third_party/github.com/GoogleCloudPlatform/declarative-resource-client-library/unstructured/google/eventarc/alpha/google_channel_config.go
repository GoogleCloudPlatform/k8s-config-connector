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
package eventarc

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/eventarc/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type GoogleChannelConfig struct{}

func GoogleChannelConfigToUnstructured(r *dclService.GoogleChannelConfig) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "eventarc",
			Version: "alpha",
			Type:    "GoogleChannelConfig",
		},
		Object: make(map[string]interface{}),
	}
	if r.CryptoKeyName != nil {
		u.Object["cryptoKeyName"] = *r.CryptoKeyName
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
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToGoogleChannelConfig(u *unstructured.Resource) (*dclService.GoogleChannelConfig, error) {
	r := &dclService.GoogleChannelConfig{}
	if _, ok := u.Object["cryptoKeyName"]; ok {
		if s, ok := u.Object["cryptoKeyName"].(string); ok {
			r.CryptoKeyName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CryptoKeyName: expected string")
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
	if _, ok := u.Object["updateTime"]; ok {
		if s, ok := u.Object["updateTime"].(string); ok {
			r.UpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UpdateTime: expected string")
		}
	}
	return r, nil
}

func GetGoogleChannelConfig(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToGoogleChannelConfig(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetGoogleChannelConfig(ctx, r)
	if err != nil {
		return nil, err
	}
	return GoogleChannelConfigToUnstructured(r), nil
}

func ApplyGoogleChannelConfig(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToGoogleChannelConfig(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToGoogleChannelConfig(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyGoogleChannelConfig(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return GoogleChannelConfigToUnstructured(r), nil
}

func GoogleChannelConfigHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToGoogleChannelConfig(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToGoogleChannelConfig(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyGoogleChannelConfig(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteGoogleChannelConfig(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToGoogleChannelConfig(u)
	if err != nil {
		return err
	}
	return c.DeleteGoogleChannelConfig(ctx, r)
}

func GoogleChannelConfigID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToGoogleChannelConfig(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *GoogleChannelConfig) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"eventarc",
		"GoogleChannelConfig",
		"alpha",
	}
}

func (r *GoogleChannelConfig) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *GoogleChannelConfig) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *GoogleChannelConfig) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *GoogleChannelConfig) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *GoogleChannelConfig) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *GoogleChannelConfig) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *GoogleChannelConfig) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetGoogleChannelConfig(ctx, config, resource)
}

func (r *GoogleChannelConfig) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyGoogleChannelConfig(ctx, config, resource, opts...)
}

func (r *GoogleChannelConfig) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return GoogleChannelConfigHasDiff(ctx, config, resource, opts...)
}

func (r *GoogleChannelConfig) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteGoogleChannelConfig(ctx, config, resource)
}

func (r *GoogleChannelConfig) ID(resource *unstructured.Resource) (string, error) {
	return GoogleChannelConfigID(resource)
}

func init() {
	unstructured.Register(&GoogleChannelConfig{})
}
