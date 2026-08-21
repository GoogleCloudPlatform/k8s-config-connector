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
package firebase

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/firebase/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type FirebaseProject struct{}

func FirebaseProjectToUnstructured(r *dclService.FirebaseProject) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "firebase",
			Version: "alpha",
			Type:    "FirebaseProject",
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
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.ProjectId != nil {
		u.Object["projectId"] = *r.ProjectId
	}
	if r.ProjectNumber != nil {
		u.Object["projectNumber"] = *r.ProjectNumber
	}
	if r.Resources != nil && r.Resources != dclService.EmptyFirebaseProjectResources {
		rResources := make(map[string]interface{})
		if r.Resources.HostingSite != nil {
			rResources["hostingSite"] = *r.Resources.HostingSite
		}
		if r.Resources.LocationId != nil {
			rResources["locationId"] = *r.Resources.LocationId
		}
		if r.Resources.RealtimeDatabaseInstance != nil {
			rResources["realtimeDatabaseInstance"] = *r.Resources.RealtimeDatabaseInstance
		}
		if r.Resources.StorageBucket != nil {
			rResources["storageBucket"] = *r.Resources.StorageBucket
		}
		u.Object["resources"] = rResources
	}
	if r.State != nil {
		u.Object["state"] = string(*r.State)
	}
	return u
}

func UnstructuredToFirebaseProject(u *unstructured.Resource) (*dclService.FirebaseProject, error) {
	r := &dclService.FirebaseProject{}
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
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["projectId"]; ok {
		if s, ok := u.Object["projectId"].(string); ok {
			r.ProjectId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ProjectId: expected string")
		}
	}
	if _, ok := u.Object["projectNumber"]; ok {
		if i, ok := u.Object["projectNumber"].(int64); ok {
			r.ProjectNumber = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.ProjectNumber: expected int64")
		}
	}
	if _, ok := u.Object["resources"]; ok {
		if rResources, ok := u.Object["resources"].(map[string]interface{}); ok {
			r.Resources = &dclService.FirebaseProjectResources{}
			if _, ok := rResources["hostingSite"]; ok {
				if s, ok := rResources["hostingSite"].(string); ok {
					r.Resources.HostingSite = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Resources.HostingSite: expected string")
				}
			}
			if _, ok := rResources["locationId"]; ok {
				if s, ok := rResources["locationId"].(string); ok {
					r.Resources.LocationId = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Resources.LocationId: expected string")
				}
			}
			if _, ok := rResources["realtimeDatabaseInstance"]; ok {
				if s, ok := rResources["realtimeDatabaseInstance"].(string); ok {
					r.Resources.RealtimeDatabaseInstance = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Resources.RealtimeDatabaseInstance: expected string")
				}
			}
			if _, ok := rResources["storageBucket"]; ok {
				if s, ok := rResources["storageBucket"].(string); ok {
					r.Resources.StorageBucket = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Resources.StorageBucket: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Resources: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["state"]; ok {
		if s, ok := u.Object["state"].(string); ok {
			r.State = dclService.FirebaseProjectStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.State: expected string")
		}
	}
	return r, nil
}

func GetFirebaseProject(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFirebaseProject(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetFirebaseProject(ctx, r)
	if err != nil {
		return nil, err
	}
	return FirebaseProjectToUnstructured(r), nil
}

func ListFirebaseProject(ctx context.Context, config *dcl.Config) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListFirebaseProject(ctx)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, FirebaseProjectToUnstructured(r))
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

func ApplyFirebaseProject(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFirebaseProject(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToFirebaseProject(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyFirebaseProject(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return FirebaseProjectToUnstructured(r), nil
}

func FirebaseProjectHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFirebaseProject(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToFirebaseProject(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyFirebaseProject(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteFirebaseProject(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func FirebaseProjectID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToFirebaseProject(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *FirebaseProject) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"firebase",
		"FirebaseProject",
		"alpha",
	}
}

func (r *FirebaseProject) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FirebaseProject) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FirebaseProject) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *FirebaseProject) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FirebaseProject) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FirebaseProject) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FirebaseProject) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetFirebaseProject(ctx, config, resource)
}

func (r *FirebaseProject) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyFirebaseProject(ctx, config, resource, opts...)
}

func (r *FirebaseProject) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return FirebaseProjectHasDiff(ctx, config, resource, opts...)
}

func (r *FirebaseProject) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteFirebaseProject(ctx, config, resource)
}

func (r *FirebaseProject) ID(resource *unstructured.Resource) (string, error) {
	return FirebaseProjectID(resource)
}

func init() {
	unstructured.Register(&FirebaseProject{})
}
