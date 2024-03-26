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
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type TcpRoute struct {
	Name        *string           `json:"name"`
	CreateTime  *string           `json:"createTime"`
	UpdateTime  *string           `json:"updateTime"`
	Description *string           `json:"description"`
	Rules       []TcpRouteRules   `json:"rules"`
	Meshes      []string          `json:"meshes"`
	Gateways    []string          `json:"gateways"`
	Labels      map[string]string `json:"labels"`
	Project     *string           `json:"project"`
	Location    *string           `json:"location"`
	SelfLink    *string           `json:"selfLink"`
}

func (r *TcpRoute) String() string {
	return dcl.SprintResource(r)
}

type TcpRouteRules struct {
	empty   bool                   `json:"-"`
	Matches []TcpRouteRulesMatches `json:"matches"`
	Action  *TcpRouteRulesAction   `json:"action"`
}

type jsonTcpRouteRules TcpRouteRules

func (r *TcpRouteRules) UnmarshalJSON(data []byte) error {
	var res jsonTcpRouteRules
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTcpRouteRules
	} else {

		r.Matches = res.Matches

		r.Action = res.Action

	}
	return nil
}

// This object is used to assert a desired state where this TcpRouteRules is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTcpRouteRules *TcpRouteRules = &TcpRouteRules{empty: true}

func (r *TcpRouteRules) Empty() bool {
	return r.empty
}

func (r *TcpRouteRules) String() string {
	return dcl.SprintResource(r)
}

func (r *TcpRouteRules) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TcpRouteRulesMatches struct {
	empty   bool    `json:"-"`
	Address *string `json:"address"`
	Port    *string `json:"port"`
}

type jsonTcpRouteRulesMatches TcpRouteRulesMatches

func (r *TcpRouteRulesMatches) UnmarshalJSON(data []byte) error {
	var res jsonTcpRouteRulesMatches
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTcpRouteRulesMatches
	} else {

		r.Address = res.Address

		r.Port = res.Port

	}
	return nil
}

// This object is used to assert a desired state where this TcpRouteRulesMatches is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTcpRouteRulesMatches *TcpRouteRulesMatches = &TcpRouteRulesMatches{empty: true}

func (r *TcpRouteRulesMatches) Empty() bool {
	return r.empty
}

func (r *TcpRouteRulesMatches) String() string {
	return dcl.SprintResource(r)
}

func (r *TcpRouteRulesMatches) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TcpRouteRulesAction struct {
	empty               bool                              `json:"-"`
	Destinations        []TcpRouteRulesActionDestinations `json:"destinations"`
	OriginalDestination *bool                             `json:"originalDestination"`
}

type jsonTcpRouteRulesAction TcpRouteRulesAction

func (r *TcpRouteRulesAction) UnmarshalJSON(data []byte) error {
	var res jsonTcpRouteRulesAction
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTcpRouteRulesAction
	} else {

		r.Destinations = res.Destinations

		r.OriginalDestination = res.OriginalDestination

	}
	return nil
}

// This object is used to assert a desired state where this TcpRouteRulesAction is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTcpRouteRulesAction *TcpRouteRulesAction = &TcpRouteRulesAction{empty: true}

func (r *TcpRouteRulesAction) Empty() bool {
	return r.empty
}

func (r *TcpRouteRulesAction) String() string {
	return dcl.SprintResource(r)
}

func (r *TcpRouteRulesAction) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TcpRouteRulesActionDestinations struct {
	empty       bool    `json:"-"`
	Weight      *int64  `json:"weight"`
	ServiceName *string `json:"serviceName"`
}

type jsonTcpRouteRulesActionDestinations TcpRouteRulesActionDestinations

func (r *TcpRouteRulesActionDestinations) UnmarshalJSON(data []byte) error {
	var res jsonTcpRouteRulesActionDestinations
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTcpRouteRulesActionDestinations
	} else {

		r.Weight = res.Weight

		r.ServiceName = res.ServiceName

	}
	return nil
}

// This object is used to assert a desired state where this TcpRouteRulesActionDestinations is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTcpRouteRulesActionDestinations *TcpRouteRulesActionDestinations = &TcpRouteRulesActionDestinations{empty: true}

func (r *TcpRouteRulesActionDestinations) Empty() bool {
	return r.empty
}

func (r *TcpRouteRulesActionDestinations) String() string {
	return dcl.SprintResource(r)
}

func (r *TcpRouteRulesActionDestinations) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *TcpRoute) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "network_services",
		Type:    "TcpRoute",
		Version: "alpha",
	}
}

func (r *TcpRoute) ID() (string, error) {
	if err := extractTcpRouteFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":        dcl.ValueOrEmptyString(nr.Name),
		"create_time": dcl.ValueOrEmptyString(nr.CreateTime),
		"update_time": dcl.ValueOrEmptyString(nr.UpdateTime),
		"description": dcl.ValueOrEmptyString(nr.Description),
		"rules":       dcl.ValueOrEmptyString(nr.Rules),
		"meshes":      dcl.ValueOrEmptyString(nr.Meshes),
		"gateways":    dcl.ValueOrEmptyString(nr.Gateways),
		"labels":      dcl.ValueOrEmptyString(nr.Labels),
		"project":     dcl.ValueOrEmptyString(nr.Project),
		"location":    dcl.ValueOrEmptyString(nr.Location),
		"self_link":   dcl.ValueOrEmptyString(nr.SelfLink),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/tcpRoutes/{{name}}", params), nil
}

const TcpRouteMaxPage = -1

type TcpRouteList struct {
	Items []*TcpRoute

	nextToken string

	pageSize int32

	resource *TcpRoute
}

func (l *TcpRouteList) HasNext() bool {
	return l.nextToken != ""
}

func (l *TcpRouteList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listTcpRoute(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListTcpRoute(ctx context.Context, project, location string) (*TcpRouteList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListTcpRouteWithMaxResults(ctx, project, location, TcpRouteMaxPage)

}

func (c *Client) ListTcpRouteWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*TcpRouteList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &TcpRoute{
		Project:  &project,
		Location: &location,
	}
	items, token, err := c.listTcpRoute(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &TcpRouteList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetTcpRoute(ctx context.Context, r *TcpRoute) (*TcpRoute, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractTcpRouteFields(r)

	b, err := c.getTcpRouteRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalTcpRoute(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeTcpRouteNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractTcpRouteFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteTcpRoute(ctx context.Context, r *TcpRoute) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("TcpRoute resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting TcpRoute...")
	deleteOp := deleteTcpRouteOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllTcpRoute deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllTcpRoute(ctx context.Context, project, location string, filter func(*TcpRoute) bool) error {
	listObj, err := c.ListTcpRoute(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllTcpRoute(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllTcpRoute(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyTcpRoute(ctx context.Context, rawDesired *TcpRoute, opts ...dcl.ApplyOption) (*TcpRoute, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *TcpRoute
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyTcpRouteHelper(c, ctx, rawDesired, opts...)
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

func applyTcpRouteHelper(c *Client, ctx context.Context, rawDesired *TcpRoute, opts ...dcl.ApplyOption) (*TcpRoute, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyTcpRoute...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractTcpRouteFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.tcpRouteDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToTcpRouteDiffs(c.Config, fieldDiffs, opts)
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
	var ops []tcpRouteApiOperation
	if create {
		ops = append(ops, &createTcpRouteOperation{})
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
	return applyTcpRouteDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyTcpRouteDiff(c *Client, ctx context.Context, desired *TcpRoute, rawDesired *TcpRoute, ops []tcpRouteApiOperation, opts ...dcl.ApplyOption) (*TcpRoute, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetTcpRoute(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createTcpRouteOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapTcpRoute(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeTcpRouteNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeTcpRouteNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeTcpRouteDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractTcpRouteFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractTcpRouteFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffTcpRoute(c, newDesired, newState)
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
