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
package beta

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Role struct {
	Name                *string              `json:"name"`
	Title               *string              `json:"title"`
	Description         *string              `json:"description"`
	LocalizedValues     *RoleLocalizedValues `json:"localizedValues"`
	LifecyclePhase      *string              `json:"lifecyclePhase"`
	GroupName           *string              `json:"groupName"`
	GroupTitle          *string              `json:"groupTitle"`
	IncludedPermissions []string             `json:"includedPermissions"`
	Stage               *RoleStageEnum       `json:"stage"`
	Etag                *string              `json:"etag"`
	Deleted             *bool                `json:"deleted"`
	IncludedRoles       []string             `json:"includedRoles"`
	Parent              *string              `json:"parent"`
}

func (r *Role) String() string {
	return dcl.SprintResource(r)
}

// The enum RoleStageEnum.
type RoleStageEnum string

// RoleStageEnumRef returns a *RoleStageEnum with the value of string s
// If the empty string is provided, nil is returned.
func RoleStageEnumRef(s string) *RoleStageEnum {
	v := RoleStageEnum(s)
	return &v
}

func (v RoleStageEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"ALPHA", "BETA", "GA", "DEPRECATED", "DISABLED", "EAP"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "RoleStageEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type RoleLocalizedValues struct {
	empty                bool    `json:"-"`
	LocalizedTitle       *string `json:"localizedTitle"`
	LocalizedDescription *string `json:"localizedDescription"`
}

type jsonRoleLocalizedValues RoleLocalizedValues

func (r *RoleLocalizedValues) UnmarshalJSON(data []byte) error {
	var res jsonRoleLocalizedValues
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyRoleLocalizedValues
	} else {

		r.LocalizedTitle = res.LocalizedTitle

		r.LocalizedDescription = res.LocalizedDescription

	}
	return nil
}

// This object is used to assert a desired state where this RoleLocalizedValues is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyRoleLocalizedValues *RoleLocalizedValues = &RoleLocalizedValues{empty: true}

func (r *RoleLocalizedValues) Empty() bool {
	return r.empty
}

func (r *RoleLocalizedValues) String() string {
	return dcl.SprintResource(r)
}

func (r *RoleLocalizedValues) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Role) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "iam",
		Type:    "Role",
		Version: "beta",
	}
}

func (r *Role) ID() (string, error) {
	if err := extractRoleFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                 dcl.ValueOrEmptyString(nr.Name),
		"title":                dcl.ValueOrEmptyString(nr.Title),
		"description":          dcl.ValueOrEmptyString(nr.Description),
		"localized_values":     dcl.ValueOrEmptyString(nr.LocalizedValues),
		"lifecycle_phase":      dcl.ValueOrEmptyString(nr.LifecyclePhase),
		"group_name":           dcl.ValueOrEmptyString(nr.GroupName),
		"group_title":          dcl.ValueOrEmptyString(nr.GroupTitle),
		"included_permissions": dcl.ValueOrEmptyString(nr.IncludedPermissions),
		"stage":                dcl.ValueOrEmptyString(nr.Stage),
		"etag":                 dcl.ValueOrEmptyString(nr.Etag),
		"deleted":              dcl.ValueOrEmptyString(nr.Deleted),
		"included_roles":       dcl.ValueOrEmptyString(nr.IncludedRoles),
		"parent":               dcl.ValueOrEmptyString(nr.Parent),
	}
	return dcl.Nprintf("{{parent}}/roles/{{name}}", params), nil
}

const RoleMaxPage = -1

type RoleList struct {
	Items []*Role

	nextToken string

	pageSize int32

	resource *Role
}

func (l *RoleList) HasNext() bool {
	return l.nextToken != ""
}

func (l *RoleList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listRole(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListRole(ctx context.Context, parent string) (*RoleList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListRoleWithMaxResults(ctx, parent, RoleMaxPage)

}

func (c *Client) ListRoleWithMaxResults(ctx context.Context, parent string, pageSize int32) (*RoleList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &Role{
		Parent: &parent,
	}
	items, token, err := c.listRole(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &RoleList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetRole(ctx context.Context, r *Role) (*Role, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractRoleFields(r)

	b, err := c.getRoleRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalRole(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Parent = r.Parent
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeRoleNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractRoleFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteRole(ctx context.Context, r *Role) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Role resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Role...")
	deleteOp := deleteRoleOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllRole deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllRole(ctx context.Context, parent string, filter func(*Role) bool) error {
	listObj, err := c.ListRole(ctx, parent)
	if err != nil {
		return err
	}

	err = c.deleteAllRole(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllRole(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyRole(ctx context.Context, rawDesired *Role, opts ...dcl.ApplyOption) (*Role, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Role
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyRoleHelper(c, ctx, rawDesired, opts...)
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

func applyRoleHelper(c *Client, ctx context.Context, rawDesired *Role, opts ...dcl.ApplyOption) (*Role, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyRole...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractRoleFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.roleDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToRoleDiffs(c.Config, fieldDiffs, opts)
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
	var ops []roleApiOperation
	if create {
		ops = append(ops, &createRoleOperation{})
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
	return applyRoleDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyRoleDiff(c *Client, ctx context.Context, desired *Role, rawDesired *Role, ops []roleApiOperation, opts ...dcl.ApplyOption) (*Role, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetRole(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createRoleOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapRole(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeRoleNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeRoleNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeRoleDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractRoleFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractRoleFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffRole(c, newDesired, newState)
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
