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
// Package iam includes tools for setting and getting policies, bindings, and members of IAM policies in the DCL.
package beta

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"reflect"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// ResourceWithPolicy is any DCL resource which has an IAM policy.
type ResourceWithPolicy interface {
	SetPolicyURL(string) string
	SetPolicyVerb() string
	GetPolicy(string) (string, string, *bytes.Buffer, error)
	IAMPolicyVersion() int
}

// Policy is the core resource of an IAM policy.
type Policy struct {
	Bindings []Binding          `json:"bindings"`
	Etag     *string            `json:"etag"`
	Version  *int               `json:"version"`
	Resource ResourceWithPolicy `json:"resource"`
}

// Binding maps a single role to all of its members.
type Binding struct {
	Role      *string            `json:"role"`
	Members   []string           `json:"members"`
	Condition *Condition         `json:"condition,omitempty"`
	Resource  ResourceWithPolicy `json:"resource"`
}

// Condition represents an IAM condition.
// See https://cloud.google.com/iam/docs/conditions-overview#resources for details.
type Condition struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Expression  *string `json:"expression"`
}

// Member maps a single IAM member to one of its roles.
type Member struct {
	Role     *string            `json:"role"`
	Member   *string            `json:"member"`
	Resource ResourceWithPolicy `json:"resource"`
}

// Encode encodes the bindings, tag, and version of an IAM policy.
func (p *Policy) Encode() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	var bindings []map[string]interface{}
	for _, b := range p.Bindings {
		bb, err := b.Encode()
		if err != nil {
			return nil, err
		}
		bindings = append(bindings, bb)
	}
	m["bindings"] = bindings
	m["etag"] = p.Etag
	m["version"] = p.Version
	return map[string]interface{}{"policy": m}, nil
}

// Encode encodes the members and role of an IAM binding.
func (b *Binding) Encode() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	m["members"] = b.Members
	m["role"] = b.Role
	if b.Condition != nil {
		m["condition"] = b.Condition
	}
	return m, nil
}

// Encode encodes the role and member of a single IAM member.
func (m *Member) Encode() (map[string]interface{}, error) {
	return map[string]interface{}{
		"role":   m.Role,
		"member": m.Member,
	}, nil
}

// GetPolicy returns the policy for the given resource.
func (c *Client) GetPolicy(ctx context.Context, r ResourceWithPolicy) (*Policy, error) {
	ctx = context.WithValue(ctx, dcl.APIRequestIDKey, dcl.CreateAPIRequestID())
	u, v, body, err := r.GetPolicy(c.Config.BasePath)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, v, u, body, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	b, err := io.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}
	p := &Policy{}
	if err := json.Unmarshal(b, p); err != nil {
		return nil, err
	}
	p.Resource = r
	return p, nil
}

// SetPolicy sets the policy for the given resource.
func (c *Client) SetPolicy(ctx context.Context, p *Policy) (*Policy, error) {
	ctx = context.WithValue(ctx, dcl.APIRequestIDKey, dcl.CreateAPIRequestID())
	currentPolicy, err := c.GetPolicy(ctx, p.Resource)
	if err != nil {
		return nil, err
	}
	p.Etag = currentPolicy.Etag
	return c.SetPolicyWithEtag(ctx, p)
}

// SetPolicyWithEtag sets the policy for the given resource using the etag contained in the Policy.
func (c *Client) SetPolicyWithEtag(ctx context.Context, p *Policy) (*Policy, error) {
	ctx = context.WithValue(ctx, dcl.APIRequestIDKey, dcl.CreateAPIRequestID())
	versionPtr := p.Resource.IAMPolicyVersion()
	p.Version = &versionPtr
	verb := p.Resource.SetPolicyVerb()
	m, err := p.Encode()
	if err != nil {
		return nil, err
	}
	if verb == "PUT" {
		// Currently only storage has this verb and requires a different format for the request body.
		policyMap, ok := m["policy"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("no policy found in map: %v", m)
		}
		m = policyMap
	}
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, verb, p.Resource.SetPolicyURL(c.Config.BasePath), bytes.NewBuffer(b), c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	respB, err := io.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}
	newP := &Policy{}
	err = json.Unmarshal(respB, newP)
	return newP, err
}

// SetBinding sets one binding, authoritatively on the role, for the given resource.
func (c *Client) SetBinding(ctx context.Context, b *Binding) (*Policy, error) {
	ctx = context.WithValue(ctx, dcl.APIRequestIDKey, dcl.CreateAPIRequestID())
	p, err := c.GetPolicy(ctx, b.Resource)
	if err != nil {
		return nil, err
	}
	roleExists := false
	for i, eb := range p.Bindings {
		if dcl.StringEquals(eb.Role, b.Role) && reflect.DeepEqual(eb.Condition, b.Condition) {
			p.Bindings[i].Members = b.Members
			roleExists = true
			break
		}
	}
	if !roleExists {
		p.Bindings = append(p.Bindings, *b)
	}
	return c.SetPolicy(ctx, p)
}

// GetBinding returns the binding for the given role, or nil if there is no such binding.
func (c *Client) GetBinding(ctx context.Context, r ResourceWithPolicy, role string) (*Binding, error) {
	ctx = context.WithValue(ctx, dcl.APIRequestIDKey, dcl.CreateAPIRequestID())
	p, err := c.GetPolicy(ctx, r)
	if err != nil {
		return nil, err
	}
	for _, eb := range p.Bindings {
		if dcl.StringEquals(eb.Role, &role) {
			return &eb, nil
		}
	}
	return nil, nil
}

// SetMember adds a member to the binding for its role if not already present.
func (c *Client) SetMember(ctx context.Context, m *Member) (*Policy, error) {
	ctx = context.WithValue(ctx, dcl.APIRequestIDKey, dcl.CreateAPIRequestID())
	b, err := c.GetBinding(ctx, m.Resource, dcl.ValueOrEmptyString(m.Role))
	if err != nil {
		return nil, err
	}
	if b == nil {
		b = &Binding{Role: m.Role}
	}
	containsMember := false
	for _, em := range b.Members {
		if dcl.StringEquals(&em, m.Member) {
			containsMember = true
			break
		}
	}
	if !containsMember {
		b.Members = append(b.Members, dcl.ValueOrEmptyString(m.Member))
	}
	b.Resource = m.Resource
	return c.SetBinding(ctx, b)
}

// GetMember returns a Member struct if the role/member pair exists on the resource's policy,
// or nil if they do not.
func (c *Client) GetMember(ctx context.Context, r ResourceWithPolicy, role, member string) (*Member, error) {
	ctx = context.WithValue(ctx, dcl.APIRequestIDKey, dcl.CreateAPIRequestID())
	b, err := c.GetBinding(ctx, r, role)
	if err != nil {
		return nil, err
	}
	if b == nil {
		return nil, &googleapi.Error{
			Code:    404,
			Message: "this role does not have a binding.",
		}
	}
	for _, em := range b.Members {
		if em == member {
			return &Member{
				Role:     &role,
				Member:   &member,
				Resource: r,
			}, nil
		}
	}
	return nil, nil
}

// convenience methods for Member, Binding, and Policy.

// ApplyBinding is a convenience method to create a binding if it does not exist.  It supports BlockAcquire and BlockCreation but
// ignores other lifecycle parameters as they are not relevant to IAM bindings.
func (c *Client) ApplyBinding(ctx context.Context, binding *Binding, opts ...dcl.ApplyOption) (*Binding, error) {
	ctx = context.WithValue(ctx, dcl.APIRequestIDKey, dcl.CreateAPIRequestID())
	lp := dcl.FetchLifecycleParams(opts)
	exists, err := c.GetBinding(ctx, binding.Resource, dcl.ValueOrEmptyString(binding.Role))
	if exists != nil && dcl.HasLifecycleParam(lp, dcl.BlockAcquire) {
		return exists, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource already exists - apply blocked by lifecycle params: %#v.", exists),
		}
	}
	if err != nil && dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
		return nil, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource does not exist - apply blocked by lifecycle params: %#v.", opts),
		}
	}
	if _, err := c.SetBinding(ctx, binding); err != nil {
		return nil, err
	}
	return binding, nil
}

// DeleteBinding deletes a binding from its specified resource.
func (c *Client) DeleteBinding(ctx context.Context, binding *Binding) error {
	ctx = context.WithValue(ctx, dcl.APIRequestIDKey, dcl.CreateAPIRequestID())
	policy, err := c.GetPolicy(ctx, binding.Resource)
	if err != nil {
		return err
	}
	var bindings []Binding
	for _, b := range policy.Bindings {
		if !dcl.StringEquals(b.Role, binding.Role) {
			bindings = append(bindings, b)
		}
	}
	policy.Bindings = bindings
	_, err = c.SetPolicy(ctx, policy)
	return err
}

// ApplyMember is a convenience method to create a member if it does not exist.  It supports BlockAcquire and BlockCreation but
// ignores other lifecycle parameters as they are not relevant to IAM members.
func (c *Client) ApplyMember(ctx context.Context, member *Member, opts ...dcl.ApplyOption) (*Member, error) {
	ctx = context.WithValue(ctx, dcl.APIRequestIDKey, dcl.CreateAPIRequestID())
	lp := dcl.FetchLifecycleParams(opts)
	role := dcl.ValueOrEmptyString(member.Role)
	memberString := dcl.ValueOrEmptyString(member.Member)
	exists, err := c.GetMember(ctx, member.Resource, role, memberString)
	if exists != nil && dcl.HasLifecycleParam(lp, dcl.BlockAcquire) {
		return exists, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource already exists - apply blocked by lifecycle params: %#v.", exists),
		}
	}
	if err != nil && dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
		return nil, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource does not exist - apply blocked by lifecycle params: %#v.", opts),
		}
	}
	_, err = c.SetMember(ctx, member)
	if err != nil {
		return nil, err
	}
	exists, err = c.GetMember(ctx, member.Resource, dcl.ValueOrEmptyString(member.Role), dcl.ValueOrEmptyString(member.Member))
	if err != nil {
		return nil, err
	}
	if exists == nil {
		return nil, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource does not exist after creation: (%#v, %#v).", dcl.ValueOrEmptyString(member.Role), dcl.ValueOrEmptyString(member.Member)),
		}
	}
	return &Member{
		Resource: member.Resource,
		Role:     exists.Role,
		Member:   exists.Member,
	}, nil
}

// DeleteMember deletes a member from its specified binding.
func (c *Client) DeleteMember(ctx context.Context, member *Member) error {
	ctx = context.WithValue(ctx, dcl.APIRequestIDKey, dcl.CreateAPIRequestID())
	binding, err := c.GetBinding(ctx, member.Resource, dcl.ValueOrEmptyString(member.Role))
	if err != nil {
		return err
	}
	if binding == nil {
		return nil
	}
	var members []string
	for _, m := range binding.Members {
		if !dcl.StringEquals(&m, member.Member) {
			members = append(members, m)
		}
	}
	binding.Members = members
	binding.Resource = member.Resource
	_, err = c.SetBinding(ctx, binding)
	return err
}

func (p *Policy) String() string {
	return dcl.SprintResource(p)
}

func (m *Member) String() string {
	return dcl.SprintResource(m)
}
