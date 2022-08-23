// Copyright 2022 Google LLC. All Rights Reserved.
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

type EnvironmentGroupAttachment struct {
	Name        *string `json:"name"`
	Environment *string `json:"environment"`
	CreatedAt   *int64  `json:"createdAt"`
	Envgroup    *string `json:"envgroup"`
}

func (r *EnvironmentGroupAttachment) String() string {
	return dcl.SprintResource(r)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *EnvironmentGroupAttachment) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "apigee",
		Type:    "EnvironmentGroupAttachment",
		Version: "beta",
	}
}

func (r *EnvironmentGroupAttachment) ID() (string, error) {
	if err := extractEnvironmentGroupAttachmentFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":        dcl.ValueOrEmptyString(nr.Name),
		"environment": dcl.ValueOrEmptyString(nr.Environment),
		"created_at":  dcl.ValueOrEmptyString(nr.CreatedAt),
		"envgroup":    dcl.ValueOrEmptyString(nr.Envgroup),
	}
	return dcl.Nprintf("{{envgroup}}/attachments/{{name}}", params), nil
}

const EnvironmentGroupAttachmentMaxPage = -1

type EnvironmentGroupAttachmentList struct {
	Items []*EnvironmentGroupAttachment

	nextToken string

	pageSize int32

	resource *EnvironmentGroupAttachment
}

func (l *EnvironmentGroupAttachmentList) HasNext() bool {
	return l.nextToken != ""
}

func (l *EnvironmentGroupAttachmentList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listEnvironmentGroupAttachment(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListEnvironmentGroupAttachment(ctx context.Context, envgroup string) (*EnvironmentGroupAttachmentList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListEnvironmentGroupAttachmentWithMaxResults(ctx, envgroup, EnvironmentGroupAttachmentMaxPage)

}

func (c *Client) ListEnvironmentGroupAttachmentWithMaxResults(ctx context.Context, envgroup string, pageSize int32) (*EnvironmentGroupAttachmentList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &EnvironmentGroupAttachment{
		Envgroup: &envgroup,
	}
	items, token, err := c.listEnvironmentGroupAttachment(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &EnvironmentGroupAttachmentList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetEnvironmentGroupAttachment(ctx context.Context, r *EnvironmentGroupAttachment) (*EnvironmentGroupAttachment, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractEnvironmentGroupAttachmentFields(r)

	b, err := c.getEnvironmentGroupAttachmentRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalEnvironmentGroupAttachment(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Envgroup = r.Envgroup
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeEnvironmentGroupAttachmentNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractEnvironmentGroupAttachmentFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteEnvironmentGroupAttachment(ctx context.Context, r *EnvironmentGroupAttachment) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("EnvironmentGroupAttachment resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting EnvironmentGroupAttachment...")
	deleteOp := deleteEnvironmentGroupAttachmentOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllEnvironmentGroupAttachment deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllEnvironmentGroupAttachment(ctx context.Context, envgroup string, filter func(*EnvironmentGroupAttachment) bool) error {
	listObj, err := c.ListEnvironmentGroupAttachment(ctx, envgroup)
	if err != nil {
		return err
	}

	err = c.deleteAllEnvironmentGroupAttachment(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllEnvironmentGroupAttachment(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyEnvironmentGroupAttachment(ctx context.Context, rawDesired *EnvironmentGroupAttachment, opts ...dcl.ApplyOption) (*EnvironmentGroupAttachment, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *EnvironmentGroupAttachment
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyEnvironmentGroupAttachmentHelper(c, ctx, rawDesired, opts...)
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

func applyEnvironmentGroupAttachmentHelper(c *Client, ctx context.Context, rawDesired *EnvironmentGroupAttachment, opts ...dcl.ApplyOption) (*EnvironmentGroupAttachment, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyEnvironmentGroupAttachment...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractEnvironmentGroupAttachmentFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.environmentGroupAttachmentDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToEnvironmentGroupAttachmentDiffs(c.Config, fieldDiffs, opts)
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
	var ops []environmentGroupAttachmentApiOperation
	if create {
		ops = append(ops, &createEnvironmentGroupAttachmentOperation{})
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
	return applyEnvironmentGroupAttachmentDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyEnvironmentGroupAttachmentDiff(c *Client, ctx context.Context, desired *EnvironmentGroupAttachment, rawDesired *EnvironmentGroupAttachment, ops []environmentGroupAttachmentApiOperation, opts ...dcl.ApplyOption) (*EnvironmentGroupAttachment, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetEnvironmentGroupAttachment(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createEnvironmentGroupAttachmentOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapEnvironmentGroupAttachment(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeEnvironmentGroupAttachmentNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeEnvironmentGroupAttachmentNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeEnvironmentGroupAttachmentDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractEnvironmentGroupAttachmentFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractEnvironmentGroupAttachmentFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffEnvironmentGroupAttachment(c, newDesired, newState)
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
