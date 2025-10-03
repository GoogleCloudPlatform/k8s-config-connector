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
package alpha

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type AuthorizationPolicy struct {
	Name        *string                        `json:"name"`
	Description *string                        `json:"description"`
	CreateTime  *string                        `json:"createTime"`
	UpdateTime  *string                        `json:"updateTime"`
	Labels      map[string]string              `json:"labels"`
	Action      *AuthorizationPolicyActionEnum `json:"action"`
	Rules       []AuthorizationPolicyRules     `json:"rules"`
	Project     *string                        `json:"project"`
	Location    *string                        `json:"location"`
}

func (r *AuthorizationPolicy) String() string {
	return dcl.SprintResource(r)
}

// The enum AuthorizationPolicyActionEnum.
type AuthorizationPolicyActionEnum string

// AuthorizationPolicyActionEnumRef returns a *AuthorizationPolicyActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func AuthorizationPolicyActionEnumRef(s string) *AuthorizationPolicyActionEnum {
	v := AuthorizationPolicyActionEnum(s)
	return &v
}

func (v AuthorizationPolicyActionEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"ACTION_UNSPECIFIED", "ALLOW", "DENY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AuthorizationPolicyActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type AuthorizationPolicyRules struct {
	empty        bool                                   `json:"-"`
	Sources      []AuthorizationPolicyRulesSources      `json:"sources"`
	Destinations []AuthorizationPolicyRulesDestinations `json:"destinations"`
}

type jsonAuthorizationPolicyRules AuthorizationPolicyRules

func (r *AuthorizationPolicyRules) UnmarshalJSON(data []byte) error {
	var res jsonAuthorizationPolicyRules
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAuthorizationPolicyRules
	} else {

		r.Sources = res.Sources

		r.Destinations = res.Destinations

	}
	return nil
}

// This object is used to assert a desired state where this AuthorizationPolicyRules is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyAuthorizationPolicyRules *AuthorizationPolicyRules = &AuthorizationPolicyRules{empty: true}

func (r *AuthorizationPolicyRules) Empty() bool {
	return r.empty
}

func (r *AuthorizationPolicyRules) String() string {
	return dcl.SprintResource(r)
}

func (r *AuthorizationPolicyRules) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AuthorizationPolicyRulesSources struct {
	empty      bool     `json:"-"`
	Principals []string `json:"principals"`
	IPBlocks   []string `json:"ipBlocks"`
}

type jsonAuthorizationPolicyRulesSources AuthorizationPolicyRulesSources

func (r *AuthorizationPolicyRulesSources) UnmarshalJSON(data []byte) error {
	var res jsonAuthorizationPolicyRulesSources
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAuthorizationPolicyRulesSources
	} else {

		r.Principals = res.Principals

		r.IPBlocks = res.IPBlocks

	}
	return nil
}

// This object is used to assert a desired state where this AuthorizationPolicyRulesSources is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyAuthorizationPolicyRulesSources *AuthorizationPolicyRulesSources = &AuthorizationPolicyRulesSources{empty: true}

func (r *AuthorizationPolicyRulesSources) Empty() bool {
	return r.empty
}

func (r *AuthorizationPolicyRulesSources) String() string {
	return dcl.SprintResource(r)
}

func (r *AuthorizationPolicyRulesSources) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AuthorizationPolicyRulesDestinations struct {
	empty           bool                                                 `json:"-"`
	Hosts           []string                                             `json:"hosts"`
	Ports           []int64                                              `json:"ports"`
	Methods         []string                                             `json:"methods"`
	HttpHeaderMatch *AuthorizationPolicyRulesDestinationsHttpHeaderMatch `json:"httpHeaderMatch"`
}

type jsonAuthorizationPolicyRulesDestinations AuthorizationPolicyRulesDestinations

func (r *AuthorizationPolicyRulesDestinations) UnmarshalJSON(data []byte) error {
	var res jsonAuthorizationPolicyRulesDestinations
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAuthorizationPolicyRulesDestinations
	} else {

		r.Hosts = res.Hosts

		r.Ports = res.Ports

		r.Methods = res.Methods

		r.HttpHeaderMatch = res.HttpHeaderMatch

	}
	return nil
}

// This object is used to assert a desired state where this AuthorizationPolicyRulesDestinations is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyAuthorizationPolicyRulesDestinations *AuthorizationPolicyRulesDestinations = &AuthorizationPolicyRulesDestinations{empty: true}

func (r *AuthorizationPolicyRulesDestinations) Empty() bool {
	return r.empty
}

func (r *AuthorizationPolicyRulesDestinations) String() string {
	return dcl.SprintResource(r)
}

func (r *AuthorizationPolicyRulesDestinations) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AuthorizationPolicyRulesDestinationsHttpHeaderMatch struct {
	empty      bool    `json:"-"`
	HeaderName *string `json:"headerName"`
	RegexMatch *string `json:"regexMatch"`
}

type jsonAuthorizationPolicyRulesDestinationsHttpHeaderMatch AuthorizationPolicyRulesDestinationsHttpHeaderMatch

func (r *AuthorizationPolicyRulesDestinationsHttpHeaderMatch) UnmarshalJSON(data []byte) error {
	var res jsonAuthorizationPolicyRulesDestinationsHttpHeaderMatch
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAuthorizationPolicyRulesDestinationsHttpHeaderMatch
	} else {

		r.HeaderName = res.HeaderName

		r.RegexMatch = res.RegexMatch

	}
	return nil
}

// This object is used to assert a desired state where this AuthorizationPolicyRulesDestinationsHttpHeaderMatch is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyAuthorizationPolicyRulesDestinationsHttpHeaderMatch *AuthorizationPolicyRulesDestinationsHttpHeaderMatch = &AuthorizationPolicyRulesDestinationsHttpHeaderMatch{empty: true}

func (r *AuthorizationPolicyRulesDestinationsHttpHeaderMatch) Empty() bool {
	return r.empty
}

func (r *AuthorizationPolicyRulesDestinationsHttpHeaderMatch) String() string {
	return dcl.SprintResource(r)
}

func (r *AuthorizationPolicyRulesDestinationsHttpHeaderMatch) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *AuthorizationPolicy) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "network_security",
		Type:    "AuthorizationPolicy",
		Version: "alpha",
	}
}

func (r *AuthorizationPolicy) ID() (string, error) {
	if err := extractAuthorizationPolicyFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":        dcl.ValueOrEmptyString(nr.Name),
		"description": dcl.ValueOrEmptyString(nr.Description),
		"create_time": dcl.ValueOrEmptyString(nr.CreateTime),
		"update_time": dcl.ValueOrEmptyString(nr.UpdateTime),
		"labels":      dcl.ValueOrEmptyString(nr.Labels),
		"action":      dcl.ValueOrEmptyString(nr.Action),
		"rules":       dcl.ValueOrEmptyString(nr.Rules),
		"project":     dcl.ValueOrEmptyString(nr.Project),
		"location":    dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/authorizationPolicies/{{name}}", params), nil
}

const AuthorizationPolicyMaxPage = -1

type AuthorizationPolicyList struct {
	Items []*AuthorizationPolicy

	nextToken string

	pageSize int32

	resource *AuthorizationPolicy
}

func (l *AuthorizationPolicyList) HasNext() bool {
	return l.nextToken != ""
}

func (l *AuthorizationPolicyList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listAuthorizationPolicy(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListAuthorizationPolicy(ctx context.Context, project, location string) (*AuthorizationPolicyList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListAuthorizationPolicyWithMaxResults(ctx, project, location, AuthorizationPolicyMaxPage)

}

func (c *Client) ListAuthorizationPolicyWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*AuthorizationPolicyList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &AuthorizationPolicy{
		Project:  &project,
		Location: &location,
	}
	items, token, err := c.listAuthorizationPolicy(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &AuthorizationPolicyList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetAuthorizationPolicy(ctx context.Context, r *AuthorizationPolicy) (*AuthorizationPolicy, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractAuthorizationPolicyFields(r)

	b, err := c.getAuthorizationPolicyRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalAuthorizationPolicy(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeAuthorizationPolicyNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractAuthorizationPolicyFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteAuthorizationPolicy(ctx context.Context, r *AuthorizationPolicy) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("AuthorizationPolicy resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting AuthorizationPolicy...")
	deleteOp := deleteAuthorizationPolicyOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllAuthorizationPolicy deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllAuthorizationPolicy(ctx context.Context, project, location string, filter func(*AuthorizationPolicy) bool) error {
	listObj, err := c.ListAuthorizationPolicy(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllAuthorizationPolicy(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllAuthorizationPolicy(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyAuthorizationPolicy(ctx context.Context, rawDesired *AuthorizationPolicy, opts ...dcl.ApplyOption) (*AuthorizationPolicy, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *AuthorizationPolicy
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyAuthorizationPolicyHelper(c, ctx, rawDesired, opts...)
		resultNewState = newState
		if err != nil {
			// If the error is 409, there is conflict in resource update.
			// Here we want to apply changes based on latest state.
			if dcl.IsConflictError(err) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		return nil, nil
	}, c.Config.RetryProvider)
	return resultNewState, err
}

func applyAuthorizationPolicyHelper(c *Client, ctx context.Context, rawDesired *AuthorizationPolicy, opts ...dcl.ApplyOption) (*AuthorizationPolicy, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyAuthorizationPolicy...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractAuthorizationPolicyFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.authorizationPolicyDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToAuthorizationPolicyDiffs(c.Config, fieldDiffs, opts)
	if err != nil {
		return nil, err
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		if dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
			return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Creation blocked by lifecycle params: %#v.", desired)}
		}
		create = true
	} else if dcl.HasLifecycleParam(lp, dcl.BlockAcquire) {
		return nil, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource already exists - apply blocked by lifecycle params: %#v.", initial),
		}
	} else {
		for _, d := range diffs {
			if d.RequiresRecreate {
				return nil, dcl.ApplyInfeasibleError{
					Message: fmt.Sprintf("infeasible update: (%v) would require recreation", d),
				}
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []authorizationPolicyApiOperation
	if create {
		ops = append(ops, &createAuthorizationPolicyOperation{})
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.InfoWithContextf(ctx, "Performing operation %T %+v", op, op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.InfoWithContextf(ctx, "Failed operation %T %+v: %v", op, op, err)
			return nil, err
		}
		c.Config.Logger.InfoWithContextf(ctx, "Finished operation %T %+v", op, op)
	}
	return applyAuthorizationPolicyDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyAuthorizationPolicyDiff(c *Client, ctx context.Context, desired *AuthorizationPolicy, rawDesired *AuthorizationPolicy, ops []authorizationPolicyApiOperation, opts ...dcl.ApplyOption) (*AuthorizationPolicy, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetAuthorizationPolicy(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createAuthorizationPolicyOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapAuthorizationPolicy(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeAuthorizationPolicyNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeAuthorizationPolicyNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeAuthorizationPolicyDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractAuthorizationPolicyFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractAuthorizationPolicyFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffAuthorizationPolicy(c, newDesired, newState)
	if err != nil {
		return newState, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.InfoWithContext(ctx, "No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.InfoWithContextf(ctx, "Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.InfoWithContext(ctx, "Done Apply.")
	return newState, nil
}

func (r *AuthorizationPolicy) GetPolicy(basePath string) (string, string, *bytes.Buffer, error) {
	u := r.getPolicyURL(basePath)
	body := &bytes.Buffer{}
	u, err := dcl.AddQueryParams(u, map[string]string{"options.requestedPolicyVersion": fmt.Sprintf("%d", r.IAMPolicyVersion())})
	if err != nil {
		return "", "", nil, err
	}
	return u, "GET", body, nil
}
