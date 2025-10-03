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
package apigee

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Environment struct {
	Name               *string               `json:"name"`
	Description        *string               `json:"description"`
	CreatedAt          *int64                `json:"createdAt"`
	LastModifiedAt     *int64                `json:"lastModifiedAt"`
	Properties         map[string]string     `json:"properties"`
	DisplayName        *string               `json:"displayName"`
	State              *EnvironmentStateEnum `json:"state"`
	ApigeeOrganization *string               `json:"apigeeOrganization"`
}

func (r *Environment) String() string {
	return dcl.SprintResource(r)
}

// The enum EnvironmentStateEnum.
type EnvironmentStateEnum string

// EnvironmentStateEnumRef returns a *EnvironmentStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func EnvironmentStateEnumRef(s string) *EnvironmentStateEnum {
	v := EnvironmentStateEnum(s)
	return &v
}

func (v EnvironmentStateEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"STATE_UNSPECIFIED", "CREATING", "ACTIVE", "DELETING"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "EnvironmentStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Environment) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "apigee",
		Type:    "Environment",
		Version: "apigee",
	}
}

func (r *Environment) ID() (string, error) {
	if err := extractEnvironmentFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                dcl.ValueOrEmptyString(nr.Name),
		"description":         dcl.ValueOrEmptyString(nr.Description),
		"created_at":          dcl.ValueOrEmptyString(nr.CreatedAt),
		"last_modified_at":    dcl.ValueOrEmptyString(nr.LastModifiedAt),
		"properties":          dcl.ValueOrEmptyString(nr.Properties),
		"display_name":        dcl.ValueOrEmptyString(nr.DisplayName),
		"state":               dcl.ValueOrEmptyString(nr.State),
		"apigee_organization": dcl.ValueOrEmptyString(nr.ApigeeOrganization),
	}
	return dcl.Nprintf("organizations/{{apigee_organization}}/environments/{{name}}", params), nil
}

const EnvironmentMaxPage = -1

type EnvironmentList struct {
	Items []*Environment

	nextToken string

	pageSize int32

	resource *Environment
}

func (c *Client) GetEnvironment(ctx context.Context, r *Environment) (*Environment, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractEnvironmentFields(r)

	b, err := c.getEnvironmentRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalEnvironment(b, c, r)
	if err != nil {
		return nil, err
	}
	result.ApigeeOrganization = r.ApigeeOrganization
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeEnvironmentNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractEnvironmentFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteEnvironment(ctx context.Context, r *Environment) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Environment resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Environment...")
	deleteOp := deleteEnvironmentOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllEnvironment deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllEnvironment(ctx context.Context, apigeeOrganization string, filter func(*Environment) bool) error {
	listObj, err := c.ListEnvironment(ctx, apigeeOrganization)
	if err != nil {
		return err
	}

	err = c.deleteAllEnvironment(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllEnvironment(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyEnvironment(ctx context.Context, rawDesired *Environment, opts ...dcl.ApplyOption) (*Environment, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Environment
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyEnvironmentHelper(c, ctx, rawDesired, opts...)
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

func applyEnvironmentHelper(c *Client, ctx context.Context, rawDesired *Environment, opts ...dcl.ApplyOption) (*Environment, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyEnvironment...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractEnvironmentFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.environmentDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToEnvironmentDiffs(c.Config, fieldDiffs, opts)
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
	var ops []environmentApiOperation
	if create {
		ops = append(ops, &createEnvironmentOperation{})
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
	return applyEnvironmentDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyEnvironmentDiff(c *Client, ctx context.Context, desired *Environment, rawDesired *Environment, ops []environmentApiOperation, opts ...dcl.ApplyOption) (*Environment, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetEnvironment(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createEnvironmentOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapEnvironment(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeEnvironmentNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeEnvironmentNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeEnvironmentDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractEnvironmentFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractEnvironmentFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffEnvironment(c, newDesired, newState)
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

func (r *Environment) GetPolicy(basePath string) (string, string, *bytes.Buffer, error) {
	u := r.getPolicyURL(basePath)
	body := &bytes.Buffer{}
	u, err := dcl.AddQueryParams(u, map[string]string{"options.requestedPolicyVersion": fmt.Sprintf("%d", r.IAMPolicyVersion())})
	if err != nil {
		return "", "", nil, err
	}
	return u, "GET", body, nil
}
