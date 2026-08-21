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

type WorkloadIdentityPoolProvider struct {
	Name                 *string                                `json:"name"`
	DisplayName          *string                                `json:"displayName"`
	Description          *string                                `json:"description"`
	State                *WorkloadIdentityPoolProviderStateEnum `json:"state"`
	Disabled             *bool                                  `json:"disabled"`
	AttributeMapping     map[string]string                      `json:"attributeMapping"`
	AttributeCondition   *string                                `json:"attributeCondition"`
	Aws                  *WorkloadIdentityPoolProviderAws       `json:"aws"`
	Oidc                 *WorkloadIdentityPoolProviderOidc      `json:"oidc"`
	Project              *string                                `json:"project"`
	Location             *string                                `json:"location"`
	WorkloadIdentityPool *string                                `json:"workloadIdentityPool"`
}

func (r *WorkloadIdentityPoolProvider) String() string {
	return dcl.SprintResource(r)
}

// The enum WorkloadIdentityPoolProviderStateEnum.
type WorkloadIdentityPoolProviderStateEnum string

// WorkloadIdentityPoolProviderStateEnumRef returns a *WorkloadIdentityPoolProviderStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func WorkloadIdentityPoolProviderStateEnumRef(s string) *WorkloadIdentityPoolProviderStateEnum {
	v := WorkloadIdentityPoolProviderStateEnum(s)
	return &v
}

func (v WorkloadIdentityPoolProviderStateEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"STATE_UNSPECIFIED", "ACTIVE", "DELETED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "WorkloadIdentityPoolProviderStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type WorkloadIdentityPoolProviderAws struct {
	empty     bool     `json:"-"`
	AccountId *string  `json:"accountId"`
	StsUri    []string `json:"stsUri"`
}

type jsonWorkloadIdentityPoolProviderAws WorkloadIdentityPoolProviderAws

func (r *WorkloadIdentityPoolProviderAws) UnmarshalJSON(data []byte) error {
	var res jsonWorkloadIdentityPoolProviderAws
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyWorkloadIdentityPoolProviderAws
	} else {

		r.AccountId = res.AccountId

		r.StsUri = res.StsUri

	}
	return nil
}

// This object is used to assert a desired state where this WorkloadIdentityPoolProviderAws is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyWorkloadIdentityPoolProviderAws *WorkloadIdentityPoolProviderAws = &WorkloadIdentityPoolProviderAws{empty: true}

func (r *WorkloadIdentityPoolProviderAws) Empty() bool {
	return r.empty
}

func (r *WorkloadIdentityPoolProviderAws) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkloadIdentityPoolProviderAws) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkloadIdentityPoolProviderOidc struct {
	empty            bool     `json:"-"`
	IssuerUri        *string  `json:"issuerUri"`
	AllowedAudiences []string `json:"allowedAudiences"`
}

type jsonWorkloadIdentityPoolProviderOidc WorkloadIdentityPoolProviderOidc

func (r *WorkloadIdentityPoolProviderOidc) UnmarshalJSON(data []byte) error {
	var res jsonWorkloadIdentityPoolProviderOidc
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyWorkloadIdentityPoolProviderOidc
	} else {

		r.IssuerUri = res.IssuerUri

		r.AllowedAudiences = res.AllowedAudiences

	}
	return nil
}

// This object is used to assert a desired state where this WorkloadIdentityPoolProviderOidc is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyWorkloadIdentityPoolProviderOidc *WorkloadIdentityPoolProviderOidc = &WorkloadIdentityPoolProviderOidc{empty: true}

func (r *WorkloadIdentityPoolProviderOidc) Empty() bool {
	return r.empty
}

func (r *WorkloadIdentityPoolProviderOidc) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkloadIdentityPoolProviderOidc) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *WorkloadIdentityPoolProvider) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "iam",
		Type:    "WorkloadIdentityPoolProvider",
		Version: "beta",
	}
}

func (r *WorkloadIdentityPoolProvider) ID() (string, error) {
	if err := extractWorkloadIdentityPoolProviderFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                   dcl.ValueOrEmptyString(nr.Name),
		"display_name":           dcl.ValueOrEmptyString(nr.DisplayName),
		"description":            dcl.ValueOrEmptyString(nr.Description),
		"state":                  dcl.ValueOrEmptyString(nr.State),
		"disabled":               dcl.ValueOrEmptyString(nr.Disabled),
		"attribute_mapping":      dcl.ValueOrEmptyString(nr.AttributeMapping),
		"attribute_condition":    dcl.ValueOrEmptyString(nr.AttributeCondition),
		"aws":                    dcl.ValueOrEmptyString(nr.Aws),
		"oidc":                   dcl.ValueOrEmptyString(nr.Oidc),
		"project":                dcl.ValueOrEmptyString(nr.Project),
		"location":               dcl.ValueOrEmptyString(nr.Location),
		"workload_identity_pool": dcl.ValueOrEmptyString(nr.WorkloadIdentityPool),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/workloadIdentityPools/{{workload_identity_pool}}/providers/{{name}}", params), nil
}

const WorkloadIdentityPoolProviderMaxPage = -1

type WorkloadIdentityPoolProviderList struct {
	Items []*WorkloadIdentityPoolProvider

	nextToken string

	pageSize int32

	resource *WorkloadIdentityPoolProvider
}

func (l *WorkloadIdentityPoolProviderList) HasNext() bool {
	return l.nextToken != ""
}

func (l *WorkloadIdentityPoolProviderList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listWorkloadIdentityPoolProvider(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListWorkloadIdentityPoolProvider(ctx context.Context, project, location, workloadIdentityPool string) (*WorkloadIdentityPoolProviderList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListWorkloadIdentityPoolProviderWithMaxResults(ctx, project, location, workloadIdentityPool, WorkloadIdentityPoolProviderMaxPage)

}

func (c *Client) ListWorkloadIdentityPoolProviderWithMaxResults(ctx context.Context, project, location, workloadIdentityPool string, pageSize int32) (*WorkloadIdentityPoolProviderList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &WorkloadIdentityPoolProvider{
		Project:              &project,
		Location:             &location,
		WorkloadIdentityPool: &workloadIdentityPool,
	}
	items, token, err := c.listWorkloadIdentityPoolProvider(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &WorkloadIdentityPoolProviderList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetWorkloadIdentityPoolProvider(ctx context.Context, r *WorkloadIdentityPoolProvider) (*WorkloadIdentityPoolProvider, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractWorkloadIdentityPoolProviderFields(r)

	b, err := c.getWorkloadIdentityPoolProviderRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalWorkloadIdentityPoolProvider(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.WorkloadIdentityPool = r.WorkloadIdentityPool
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeWorkloadIdentityPoolProviderNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractWorkloadIdentityPoolProviderFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteWorkloadIdentityPoolProvider(ctx context.Context, r *WorkloadIdentityPoolProvider) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("WorkloadIdentityPoolProvider resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting WorkloadIdentityPoolProvider...")
	deleteOp := deleteWorkloadIdentityPoolProviderOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllWorkloadIdentityPoolProvider deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllWorkloadIdentityPoolProvider(ctx context.Context, project, location, workloadIdentityPool string, filter func(*WorkloadIdentityPoolProvider) bool) error {
	listObj, err := c.ListWorkloadIdentityPoolProvider(ctx, project, location, workloadIdentityPool)
	if err != nil {
		return err
	}

	err = c.deleteAllWorkloadIdentityPoolProvider(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllWorkloadIdentityPoolProvider(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyWorkloadIdentityPoolProvider(ctx context.Context, rawDesired *WorkloadIdentityPoolProvider, opts ...dcl.ApplyOption) (*WorkloadIdentityPoolProvider, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *WorkloadIdentityPoolProvider
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyWorkloadIdentityPoolProviderHelper(c, ctx, rawDesired, opts...)
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

func applyWorkloadIdentityPoolProviderHelper(c *Client, ctx context.Context, rawDesired *WorkloadIdentityPoolProvider, opts ...dcl.ApplyOption) (*WorkloadIdentityPoolProvider, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyWorkloadIdentityPoolProvider...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractWorkloadIdentityPoolProviderFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.workloadIdentityPoolProviderDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToWorkloadIdentityPoolProviderDiffs(c.Config, fieldDiffs, opts)
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
	var ops []workloadIdentityPoolProviderApiOperation
	if create {
		ops = append(ops, &createWorkloadIdentityPoolProviderOperation{})
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
	return applyWorkloadIdentityPoolProviderDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyWorkloadIdentityPoolProviderDiff(c *Client, ctx context.Context, desired *WorkloadIdentityPoolProvider, rawDesired *WorkloadIdentityPoolProvider, ops []workloadIdentityPoolProviderApiOperation, opts ...dcl.ApplyOption) (*WorkloadIdentityPoolProvider, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetWorkloadIdentityPoolProvider(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createWorkloadIdentityPoolProviderOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapWorkloadIdentityPoolProvider(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeWorkloadIdentityPoolProviderNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeWorkloadIdentityPoolProviderNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeWorkloadIdentityPoolProviderDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractWorkloadIdentityPoolProviderFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractWorkloadIdentityPoolProviderFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffWorkloadIdentityPoolProvider(c, newDesired, newState)
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
