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
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryreservation/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Reservation struct{}

func ReservationToUnstructured(r *dclService.Reservation) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "bigqueryreservation",
			Version: "beta",
			Type:    "Reservation",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreationTime != nil {
		u.Object["creationTime"] = *r.CreationTime
	}
	if r.IgnoreIdleSlots != nil {
		u.Object["ignoreIdleSlots"] = *r.IgnoreIdleSlots
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.MaxConcurrency != nil {
		u.Object["maxConcurrency"] = *r.MaxConcurrency
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.SlotCapacity != nil {
		u.Object["slotCapacity"] = *r.SlotCapacity
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToReservation(u *unstructured.Resource) (*dclService.Reservation, error) {
	r := &dclService.Reservation{}
	if _, ok := u.Object["creationTime"]; ok {
		if s, ok := u.Object["creationTime"].(string); ok {
			r.CreationTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreationTime: expected string")
		}
	}
	if _, ok := u.Object["ignoreIdleSlots"]; ok {
		if b, ok := u.Object["ignoreIdleSlots"].(bool); ok {
			r.IgnoreIdleSlots = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.IgnoreIdleSlots: expected bool")
		}
	}
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["maxConcurrency"]; ok {
		if i, ok := u.Object["maxConcurrency"].(int64); ok {
			r.MaxConcurrency = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.MaxConcurrency: expected int64")
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
	if _, ok := u.Object["slotCapacity"]; ok {
		if i, ok := u.Object["slotCapacity"].(int64); ok {
			r.SlotCapacity = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.SlotCapacity: expected int64")
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

func GetReservation(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToReservation(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetReservation(ctx, r)
	if err != nil {
		return nil, err
	}
	return ReservationToUnstructured(r), nil
}

func ListReservation(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListReservation(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, ReservationToUnstructured(r))
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

func ApplyReservation(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToReservation(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToReservation(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyReservation(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return ReservationToUnstructured(r), nil
}

func ReservationHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToReservation(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToReservation(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyReservation(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteReservation(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToReservation(u)
	if err != nil {
		return err
	}
	return c.DeleteReservation(ctx, r)
}

func ReservationID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToReservation(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Reservation) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"bigqueryreservation",
		"Reservation",
		"beta",
	}
}

func (r *Reservation) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Reservation) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Reservation) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Reservation) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Reservation) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Reservation) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Reservation) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetReservation(ctx, config, resource)
}

func (r *Reservation) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyReservation(ctx, config, resource, opts...)
}

func (r *Reservation) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return ReservationHasDiff(ctx, config, resource, opts...)
}

func (r *Reservation) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteReservation(ctx, config, resource)
}

func (r *Reservation) ID(resource *unstructured.Resource) (string, error) {
	return ReservationID(resource)
}

func init() {
	unstructured.Register(&Reservation{})
}
