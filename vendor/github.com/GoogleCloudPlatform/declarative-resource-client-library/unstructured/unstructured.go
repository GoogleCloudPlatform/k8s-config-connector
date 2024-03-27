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
// Package unstructured provides a generic unstructured client to invoke DCL.
package unstructured

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
)

var (
	registrations     []RegisteredResource
	registrationMutex sync.RWMutex

	// ErrNoSuchMethod is the error returned when calling Get, List, Apply, or Delete
	// on an API that doesn't support the requested method.
	ErrNoSuchMethod = errors.New("non-existent API method")
)

// ServiceTypeVersion describes a single DCL resource.
type ServiceTypeVersion struct {
	// Service to which this resource belongs, e.g., "compute".
	Service string
	// Type of the resource, e.g., "ComputeInstance"
	Type string
	// Version of the resource, e.g., "ga". There may be multiple versions of the
	// same Type and Service in a single DCL build.
	Version string
}

// Resource is the untyped representation of a typed DCL resource.
type Resource struct {
	// ResourceWithPolicy is included so that the Resource struct can interact with the
	// IAMClient (IAMClient methods expect a ResourceWithPolicy)
	iam.ResourceWithPolicy

	// Object is a JSON compatible map with string, float, int,
	// bool, []interface{}, or map[string]interface{} children.
	Object map[string]interface{}

	// STV indicates the type of this resource
	STV ServiceTypeVersion
}

// RegisteredResource is used by generated unstructured library code to
// make type-specific operations available in a type-agnostic manner.
type RegisteredResource interface {
	// STV indicates the type of this resource.
	STV() ServiceTypeVersion

	// Get provides an indirection for the type-specific Get call.
	Get(ctx context.Context, config *dcl.Config, r *Resource) (*Resource, error)

	// Apply provides an indirection for the type-specific Apply call.
	Apply(ctx context.Context, config *dcl.Config, r *Resource, opts ...dcl.ApplyOption) (*Resource, error)

	// HasDiff provides an indirection for the type-specific HasDiff call.
	HasDiff(ctx context.Context, config *dcl.Config, r *Resource, opts ...dcl.ApplyOption) (bool, error)

	// Delete provides an indirection for the type-specific Delete call.
	Delete(ctx context.Context, config *dcl.Config, r *Resource) error

	// GetPolicy provides an indirection for the type-specific GetPolicy call.
	GetPolicy(ctx context.Context, config *dcl.Config, r *Resource) (*Resource, error)

	// SetPolicy provides an indirection for the type-specific SetPolicy call.
	SetPolicy(ctx context.Context, config *dcl.Config, r *Resource, p *Resource) (*Resource, error)

	// SetPolicyWithEtag provides an indirection for the type-specific SetPolicy call.
	SetPolicyWithEtag(ctx context.Context, config *dcl.Config, r *Resource, p *Resource) (*Resource, error)

	// GetPolicyMember provides an indirection for the type-specific GetPolicyMember call.
	GetPolicyMember(ctx context.Context, config *dcl.Config, r *Resource, role, member string) (*Resource, error)

	// SetPolicyMember provides an indirection for the type-specific SetPolicyMember call.
	SetPolicyMember(ctx context.Context, config *dcl.Config, r *Resource, m *Resource) (*Resource, error)

	// DeletePolicyMember provides an indirection for the type-specific DeletePolicyMember call.
	DeletePolicyMember(ctx context.Context, config *dcl.Config, r *Resource, m *Resource) error

	// ID returns a string uniquely identifying this resource.
	ID(r *Resource) (string, error)
}

// Equals compares two ServiceTypeVersion structures.
func (stv ServiceTypeVersion) Equals(o ServiceTypeVersion) bool {
	return stv.Service == o.Service && stv.Type == o.Type && stv.Version == o.Version
}

// String returns a loggable description of this ServiceTypeVersion.
func (stv ServiceTypeVersion) String() string {
	return fmt.Sprintf(`"%s.%s.%s"`, stv.Service, stv.Type, stv.Version)
}

// StateHint is a dcl.ApplyOption that acts as the unstructured analog to dcl.stateHint.
type stateHint struct {
	state *Resource
}

// Apply is a no-op to conform to the dcl.ApplyOption interface.
func (s stateHint) Apply(o *dcl.ApplyOpts) {}

// WithStateHint performs the same function as dcl.WithStateHint, but
// takes an unstructured resource.
func WithStateHint(r *Resource) dcl.ApplyOption {
	return stateHint{state: r}
}

// FetchStateHint returns either nil or a Resource representing the pre-apply state.
func FetchStateHint(c []dcl.ApplyOption) *Resource {
	for _, p := range c {
		if sh, ok := p.(stateHint); ok {
			return sh.state
		}
	}
	return nil
}

// Register adds the provided resource to the list of resources available
// via the generic Get/List/Apply/Delete functions.
func Register(rr RegisteredResource) {
	registrationMutex.Lock()
	defer registrationMutex.Unlock()
	registrations = append(registrations, rr)
}

func registration(r *Resource) RegisteredResource {
	registrationMutex.RLock()
	defer registrationMutex.RUnlock()
	for _, rr := range registrations {
		if rr.STV().Equals(r.STV) {
			return rr
		}
	}
	return nil
}

// Get returns the current version of a given resource (usually from the
// result of a previous Apply()).
func Get(ctx context.Context, config *dcl.Config, r *Resource) (*Resource, error) {
	rr := registration(r)
	if rr == nil {
		return nil, fmt.Errorf("unknown resource type %s", r.STV.String())
	}
	return rr.Get(ctx, config, r)
}

// Apply creates or updates the provided resource.
func Apply(ctx context.Context, config *dcl.Config, r *Resource, opts ...dcl.ApplyOption) (*Resource, error) {
	rr := registration(r)
	if rr == nil {
		return nil, fmt.Errorf("unknown resource type %s", r.STV.String())
	}
	return rr.Apply(ctx, config, r, opts...)
}

// HasDiff returns whether the provided resource config matches the live resource, i.e.,
// a return value of true indicates that calling Apply() will cause a creation or update of
// the live resource. The `opts` parameter can optionally include a state hint.
func HasDiff(ctx context.Context, config *dcl.Config, r *Resource, opts ...dcl.ApplyOption) (bool, error) {
	rr := registration(r)
	if rr == nil {
		return false, fmt.Errorf("unknown resource type %s", r.STV.String())
	}
	return rr.HasDiff(ctx, config, r, opts...)
}

// Delete deletes the provided resource.
func Delete(ctx context.Context, config *dcl.Config, r *Resource) error {
	rr := registration(r)
	if rr == nil {
		return fmt.Errorf("unknown resource type %s", r.STV.String())
	}
	return rr.Delete(ctx, config, r)
}

// GetPolicy gets the IAMPolicy for the provided resource.
func GetPolicy(ctx context.Context, config *dcl.Config, r *Resource) (*Resource, error) {
	rr := registration(r)
	if rr == nil {
		return nil, fmt.Errorf("unknown resource type %s", r.STV.String())
	}
	return rr.GetPolicy(ctx, config, r)
}

// SetPolicy sets the IAMPolicy for the provided resource.
func SetPolicy(ctx context.Context, config *dcl.Config, r *Resource, p *Resource) (*Resource, error) {
	rr := registration(r)
	if rr == nil {
		return nil, fmt.Errorf("unknown resource type %s", r.STV.String())
	}
	return rr.SetPolicy(ctx, config, r, p)
}

// SetPolicyWithEtag sets the IAMPolicy using the etag container for the provided resource.
func SetPolicyWithEtag(ctx context.Context, config *dcl.Config, r *Resource, p *Resource) (*Resource, error) {
	rr := registration(r)
	if rr == nil {
		return nil, fmt.Errorf("unknown resource type %s", r.STV.String())
	}
	return rr.SetPolicyWithEtag(ctx, config, r, p)
}

// GetPolicyMember gets the IAMPolicyMember for the provided resource.
func GetPolicyMember(ctx context.Context, config *dcl.Config, r *Resource, role, member string) (*Resource, error) {
	rr := registration(r)
	if rr == nil {
		return nil, fmt.Errorf("unknown resource type %s", r.STV.String())
	}
	return rr.GetPolicyMember(ctx, config, r, role, member)
}

// SetPolicyMember sets the IAMPolicyMember for the provided resource.
func SetPolicyMember(ctx context.Context, config *dcl.Config, r *Resource, m *Resource) (*Resource, error) {
	rr := registration(r)
	if rr == nil {
		return nil, fmt.Errorf("unknown resource type %s", r.STV.String())
	}
	return rr.SetPolicyMember(ctx, config, r, m)
}

// DeletePolicyMember deletes the IAMPolicyMember for the provided resource.
func DeletePolicyMember(ctx context.Context, config *dcl.Config, r *Resource, m *Resource) error {
	rr := registration(r)
	if rr == nil {
		return fmt.Errorf("unknown resource type %s", r.STV.String())
	}
	return rr.DeletePolicyMember(ctx, config, r, m)
}

// ID returns a unique ID for the provided resource.
func ID(r *Resource) (string, error) {
	rr := registration(r)
	if rr == nil {
		return "", fmt.Errorf("unknown resource type %s", r.STV.String())
	}
	return rr.ID(r)
}
