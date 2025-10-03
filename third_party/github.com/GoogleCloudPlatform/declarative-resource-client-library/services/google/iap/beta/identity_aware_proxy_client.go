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
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type IdentityAwareProxyClient struct {
	Name        *string `json:"name"`
	Secret      *string `json:"secret"`
	DisplayName *string `json:"displayName"`
	Project     *string `json:"project"`
	Brand       *string `json:"brand"`
}

func (r *IdentityAwareProxyClient) String() string {
	return dcl.SprintResource(r)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *IdentityAwareProxyClient) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "iap",
		Type:    "IdentityAwareProxyClient",
		Version: "beta",
	}
}

func (r *IdentityAwareProxyClient) ID() (string, error) {
	if err := extractIdentityAwareProxyClientFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":         dcl.ValueOrEmptyString(nr.Name),
		"secret":       dcl.ValueOrEmptyString(nr.Secret),
		"display_name": dcl.ValueOrEmptyString(nr.DisplayName),
		"project":      dcl.ValueOrEmptyString(nr.Project),
		"brand":        dcl.ValueOrEmptyString(nr.Brand),
	}
	return dcl.Nprintf("projects/{{project}}/brands/{{brand}}/identityAwareProxyClients/{{name}}", params), nil
}

const IdentityAwareProxyClientMaxPage = -1

type IdentityAwareProxyClientList struct {
	Items []*IdentityAwareProxyClient

	nextToken string

	pageSize int32

	resource *IdentityAwareProxyClient
}

func (l *IdentityAwareProxyClientList) HasNext() bool {
	return l.nextToken != ""
}

func (l *IdentityAwareProxyClientList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listIdentityAwareProxyClient(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListIdentityAwareProxyClient(ctx context.Context, project, brand string) (*IdentityAwareProxyClientList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListIdentityAwareProxyClientWithMaxResults(ctx, project, brand, IdentityAwareProxyClientMaxPage)

}

func (c *Client) ListIdentityAwareProxyClientWithMaxResults(ctx context.Context, project, brand string, pageSize int32) (*IdentityAwareProxyClientList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &IdentityAwareProxyClient{
		Project: &project,
		Brand:   &brand,
	}
	items, token, err := c.listIdentityAwareProxyClient(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &IdentityAwareProxyClientList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetIdentityAwareProxyClient(ctx context.Context, r *IdentityAwareProxyClient) (*IdentityAwareProxyClient, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractIdentityAwareProxyClientFields(r)

	b, err := c.getIdentityAwareProxyClientRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalIdentityAwareProxyClient(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Brand = r.Brand
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeIdentityAwareProxyClientNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractIdentityAwareProxyClientFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteIdentityAwareProxyClient(ctx context.Context, r *IdentityAwareProxyClient) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("IdentityAwareProxyClient resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting IdentityAwareProxyClient...")
	deleteOp := deleteIdentityAwareProxyClientOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllIdentityAwareProxyClient deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllIdentityAwareProxyClient(ctx context.Context, project, brand string, filter func(*IdentityAwareProxyClient) bool) error {
	listObj, err := c.ListIdentityAwareProxyClient(ctx, project, brand)
	if err != nil {
		return err
	}

	err = c.deleteAllIdentityAwareProxyClient(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllIdentityAwareProxyClient(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyIdentityAwareProxyClient(ctx context.Context, rawDesired *IdentityAwareProxyClient, opts ...dcl.ApplyOption) (*IdentityAwareProxyClient, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *IdentityAwareProxyClient
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyIdentityAwareProxyClientHelper(c, ctx, rawDesired, opts...)
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

func applyIdentityAwareProxyClientHelper(c *Client, ctx context.Context, rawDesired *IdentityAwareProxyClient, opts ...dcl.ApplyOption) (*IdentityAwareProxyClient, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyIdentityAwareProxyClient...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractIdentityAwareProxyClientFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.identityAwareProxyClientDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToIdentityAwareProxyClientDiffs(c.Config, fieldDiffs, opts)
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
	var ops []identityAwareProxyClientApiOperation
	if create {
		ops = append(ops, &createIdentityAwareProxyClientOperation{})
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
	return applyIdentityAwareProxyClientDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyIdentityAwareProxyClientDiff(c *Client, ctx context.Context, desired *IdentityAwareProxyClient, rawDesired *IdentityAwareProxyClient, ops []identityAwareProxyClientApiOperation, opts ...dcl.ApplyOption) (*IdentityAwareProxyClient, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetIdentityAwareProxyClient(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createIdentityAwareProxyClientOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapIdentityAwareProxyClient(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeIdentityAwareProxyClientNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeIdentityAwareProxyClientNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeIdentityAwareProxyClientDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractIdentityAwareProxyClientFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractIdentityAwareProxyClientFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffIdentityAwareProxyClient(c, newDesired, newState)
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
