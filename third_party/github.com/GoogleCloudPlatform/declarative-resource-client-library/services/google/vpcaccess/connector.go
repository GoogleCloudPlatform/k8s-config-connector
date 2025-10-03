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
package vpcaccess

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Connector struct {
	Name              *string             `json:"name"`
	Network           *string             `json:"network"`
	IPCidrRange       *string             `json:"ipCidrRange"`
	State             *ConnectorStateEnum `json:"state"`
	MinThroughput     *int64              `json:"minThroughput"`
	MaxThroughput     *int64              `json:"maxThroughput"`
	ConnectedProjects []string            `json:"connectedProjects"`
	Subnet            *ConnectorSubnet    `json:"subnet"`
	MachineType       *string             `json:"machineType"`
	MinInstances      *int64              `json:"minInstances"`
	MaxInstances      *int64              `json:"maxInstances"`
	Project           *string             `json:"project"`
	Location          *string             `json:"location"`
}

func (r *Connector) String() string {
	return dcl.SprintResource(r)
}

// The enum ConnectorStateEnum.
type ConnectorStateEnum string

// ConnectorStateEnumRef returns a *ConnectorStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func ConnectorStateEnumRef(s string) *ConnectorStateEnum {
	v := ConnectorStateEnum(s)
	return &v
}

func (v ConnectorStateEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"STATE_UNSPECIFIED", "READY", "CREATING", "DELETING", "ERROR", "UPDATING"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ConnectorStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type ConnectorSubnet struct {
	empty     bool    `json:"-"`
	Name      *string `json:"name"`
	ProjectId *string `json:"projectId"`
}

type jsonConnectorSubnet ConnectorSubnet

func (r *ConnectorSubnet) UnmarshalJSON(data []byte) error {
	var res jsonConnectorSubnet
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConnectorSubnet
	} else {

		r.Name = res.Name

		r.ProjectId = res.ProjectId

	}
	return nil
}

// This object is used to assert a desired state where this ConnectorSubnet is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConnectorSubnet *ConnectorSubnet = &ConnectorSubnet{empty: true}

func (r *ConnectorSubnet) Empty() bool {
	return r.empty
}

func (r *ConnectorSubnet) String() string {
	return dcl.SprintResource(r)
}

func (r *ConnectorSubnet) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Connector) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "vpc_access",
		Type:    "Connector",
		Version: "vpcaccess",
	}
}

func (r *Connector) ID() (string, error) {
	if err := extractConnectorFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":               dcl.ValueOrEmptyString(nr.Name),
		"network":            dcl.ValueOrEmptyString(nr.Network),
		"ip_cidr_range":      dcl.ValueOrEmptyString(nr.IPCidrRange),
		"state":              dcl.ValueOrEmptyString(nr.State),
		"min_throughput":     dcl.ValueOrEmptyString(nr.MinThroughput),
		"max_throughput":     dcl.ValueOrEmptyString(nr.MaxThroughput),
		"connected_projects": dcl.ValueOrEmptyString(nr.ConnectedProjects),
		"subnet":             dcl.ValueOrEmptyString(nr.Subnet),
		"machine_type":       dcl.ValueOrEmptyString(nr.MachineType),
		"min_instances":      dcl.ValueOrEmptyString(nr.MinInstances),
		"max_instances":      dcl.ValueOrEmptyString(nr.MaxInstances),
		"project":            dcl.ValueOrEmptyString(nr.Project),
		"location":           dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/connectors/{{name}}", params), nil
}

const ConnectorMaxPage = -1

type ConnectorList struct {
	Items []*Connector

	nextToken string

	pageSize int32

	resource *Connector
}

func (l *ConnectorList) HasNext() bool {
	return l.nextToken != ""
}

func (l *ConnectorList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listConnector(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListConnector(ctx context.Context, project, location string) (*ConnectorList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListConnectorWithMaxResults(ctx, project, location, ConnectorMaxPage)

}

func (c *Client) ListConnectorWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*ConnectorList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &Connector{
		Project:  &project,
		Location: &location,
	}
	items, token, err := c.listConnector(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &ConnectorList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetConnector(ctx context.Context, r *Connector) (*Connector, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractConnectorFields(r)

	b, err := c.getConnectorRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalConnector(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeConnectorNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractConnectorFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteConnector(ctx context.Context, r *Connector) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Connector resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Connector...")
	deleteOp := deleteConnectorOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllConnector deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllConnector(ctx context.Context, project, location string, filter func(*Connector) bool) error {
	listObj, err := c.ListConnector(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllConnector(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllConnector(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyConnector(ctx context.Context, rawDesired *Connector, opts ...dcl.ApplyOption) (*Connector, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Connector
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyConnectorHelper(c, ctx, rawDesired, opts...)
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

func applyConnectorHelper(c *Client, ctx context.Context, rawDesired *Connector, opts ...dcl.ApplyOption) (*Connector, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyConnector...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractConnectorFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.connectorDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToConnectorDiffs(c.Config, fieldDiffs, opts)
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
	var ops []connectorApiOperation
	if create {
		ops = append(ops, &createConnectorOperation{})
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
	return applyConnectorDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyConnectorDiff(c *Client, ctx context.Context, desired *Connector, rawDesired *Connector, ops []connectorApiOperation, opts ...dcl.ApplyOption) (*Connector, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetConnector(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createConnectorOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapConnector(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeConnectorNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeConnectorNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeConnectorDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractConnectorFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractConnectorFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffConnector(c, newDesired, newState)
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
