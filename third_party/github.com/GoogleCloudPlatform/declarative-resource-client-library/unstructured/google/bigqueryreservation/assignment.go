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
package bigqueryreservation

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryreservation"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Assignment struct{}

func AssignmentToUnstructured(r *dclService.Assignment) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "bigqueryreservation",
			Version: "ga",
			Type:    "Assignment",
		},
		Object: make(map[string]interface{}),
	}
	if r.Assignee != nil {
		u.Object["assignee"] = *r.Assignee
	}
	if r.JobType != nil {
		u.Object["jobType"] = string(*r.JobType)
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
	if r.Reservation != nil {
		u.Object["reservation"] = *r.Reservation
	}
	if r.State != nil {
		u.Object["state"] = string(*r.State)
	}
	return u
}

func UnstructuredToAssignment(u *unstructured.Resource) (*dclService.Assignment, error) {
	r := &dclService.Assignment{}
	if _, ok := u.Object["assignee"]; ok {
		if s, ok := u.Object["assignee"].(string); ok {
			r.Assignee = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Assignee: expected string")
		}
	}
	if _, ok := u.Object["jobType"]; ok {
		if s, ok := u.Object["jobType"].(string); ok {
			r.JobType = dclService.AssignmentJobTypeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.JobType: expected string")
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
	if _, ok := u.Object["reservation"]; ok {
		if s, ok := u.Object["reservation"].(string); ok {
			r.Reservation = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Reservation: expected string")
		}
	}
	if _, ok := u.Object["state"]; ok {
		if s, ok := u.Object["state"].(string); ok {
			r.State = dclService.AssignmentStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.State: expected string")
		}
	}
	return r, nil
}

func GetAssignment(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAssignment(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetAssignment(ctx, r)
	if err != nil {
		return nil, err
	}
	return AssignmentToUnstructured(r), nil
}

func ListAssignment(ctx context.Context, config *dcl.Config, project string, location string, reservation string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListAssignment(ctx, project, location, reservation)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, AssignmentToUnstructured(r))
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

func ApplyAssignment(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAssignment(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToAssignment(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyAssignment(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return AssignmentToUnstructured(r), nil
}

func AssignmentHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAssignment(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToAssignment(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyAssignment(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteAssignment(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAssignment(u)
	if err != nil {
		return err
	}
	return c.DeleteAssignment(ctx, r)
}

func AssignmentID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToAssignment(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Assignment) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"bigqueryreservation",
		"Assignment",
		"ga",
	}
}

func (r *Assignment) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Assignment) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Assignment) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Assignment) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Assignment) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Assignment) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Assignment) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetAssignment(ctx, config, resource)
}

func (r *Assignment) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyAssignment(ctx, config, resource, opts...)
}

func (r *Assignment) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return AssignmentHasDiff(ctx, config, resource, opts...)
}

func (r *Assignment) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteAssignment(ctx, config, resource)
}

func (r *Assignment) ID(resource *unstructured.Resource) (string, error) {
	return AssignmentID(resource)
}

func init() {
	unstructured.Register(&Assignment{})
}
