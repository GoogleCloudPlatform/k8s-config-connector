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
	"context"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type ServiceBinding struct {
	Name        *string           `json:"name"`
	Description *string           `json:"description"`
	CreateTime  *string           `json:"createTime"`
	UpdateTime  *string           `json:"updateTime"`
	Service     *string           `json:"service"`
	Labels      map[string]string `json:"labels"`
	Project     *string           `json:"project"`
	Location    *string           `json:"location"`
}

func (r *ServiceBinding) String() string {
	return dcl.SprintResource(r)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *ServiceBinding) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "network_services",
		Type:    "ServiceBinding",
		Version: "alpha",
	}
}

func (r *ServiceBinding) ID() (string, error) {
	if err := extractServiceBindingFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":        dcl.ValueOrEmptyString(nr.Name),
		"description": dcl.ValueOrEmptyString(nr.Description),
		"create_time": dcl.ValueOrEmptyString(nr.CreateTime),
		"update_time": dcl.ValueOrEmptyString(nr.UpdateTime),
		"service":     dcl.ValueOrEmptyString(nr.Service),
		"labels":      dcl.ValueOrEmptyString(nr.Labels),
		"project":     dcl.ValueOrEmptyString(nr.Project),
		"location":    dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/serviceBindings/{{name}}", params), nil
}

const ServiceBindingMaxPage = -1

type ServiceBindingList struct {
	Items []*ServiceBinding

	nextToken string

	pageSize int32

	resource *ServiceBinding
}

func (l *ServiceBindingList) HasNext() bool {
	return l.nextToken != ""
}

func (l *ServiceBindingList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listServiceBinding(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListServiceBinding(ctx context.Context, project, location string) (*ServiceBindingList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListServiceBindingWithMaxResults(ctx, project, location, ServiceBindingMaxPage)

}

func (c *Client) ListServiceBindingWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*ServiceBindingList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &ServiceBinding{
		Project:  &project,
		Location: &location,
	}
	items, token, err := c.listServiceBinding(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &ServiceBindingList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetServiceBinding(ctx context.Context, r *ServiceBinding) (*ServiceBinding, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractServiceBindingFields(r)

	b, err := c.getServiceBindingRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalServiceBinding(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeServiceBindingNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractServiceBindingFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteServiceBinding(ctx context.Context, r *ServiceBinding) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("ServiceBinding resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting ServiceBinding...")
	deleteOp := deleteServiceBindingOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllServiceBinding deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllServiceBinding(ctx context.Context, project, location string, filter func(*ServiceBinding) bool) error {
	listObj, err := c.ListServiceBinding(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllServiceBinding(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllServiceBinding(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyServiceBinding(ctx context.Context, rawDesired *ServiceBinding, opts ...dcl.ApplyOption) (*ServiceBinding, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *ServiceBinding
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyServiceBindingHelper(c, ctx, rawDesired, opts...)
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

func applyServiceBindingHelper(c *Client, ctx context.Context, rawDesired *ServiceBinding, opts ...dcl.ApplyOption) (*ServiceBinding, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyServiceBinding...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractServiceBindingFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.serviceBindingDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToServiceBindingDiffs(c.Config, fieldDiffs, opts)
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
	var ops []serviceBindingApiOperation
	if create {
		ops = append(ops, &createServiceBindingOperation{})
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
	return applyServiceBindingDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyServiceBindingDiff(c *Client, ctx context.Context, desired *ServiceBinding, rawDesired *ServiceBinding, ops []serviceBindingApiOperation, opts ...dcl.ApplyOption) (*ServiceBinding, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetServiceBinding(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createServiceBindingOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapServiceBinding(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeServiceBindingNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeServiceBindingNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeServiceBindingDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractServiceBindingFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractServiceBindingFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffServiceBinding(c, newDesired, newState)
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
