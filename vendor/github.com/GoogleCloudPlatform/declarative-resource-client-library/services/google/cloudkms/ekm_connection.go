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
package cloudkms

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

type EkmConnection struct {
	Name             *string                         `json:"name"`
	CreateTime       *string                         `json:"createTime"`
	ServiceResolvers []EkmConnectionServiceResolvers `json:"serviceResolvers"`
	Etag             *string                         `json:"etag"`
	Project          *string                         `json:"project"`
	Location         *string                         `json:"location"`
}

func (r *EkmConnection) String() string {
	return dcl.SprintResource(r)
}

type EkmConnectionServiceResolvers struct {
	empty                   bool                                              `json:"-"`
	ServiceDirectoryService *string                                           `json:"serviceDirectoryService"`
	EndpointFilter          *string                                           `json:"endpointFilter"`
	Hostname                *string                                           `json:"hostname"`
	ServerCertificates      []EkmConnectionServiceResolversServerCertificates `json:"serverCertificates"`
}

type jsonEkmConnectionServiceResolvers EkmConnectionServiceResolvers

func (r *EkmConnectionServiceResolvers) UnmarshalJSON(data []byte) error {
	var res jsonEkmConnectionServiceResolvers
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEkmConnectionServiceResolvers
	} else {

		r.ServiceDirectoryService = res.ServiceDirectoryService

		r.EndpointFilter = res.EndpointFilter

		r.Hostname = res.Hostname

		r.ServerCertificates = res.ServerCertificates

	}
	return nil
}

// This object is used to assert a desired state where this EkmConnectionServiceResolvers is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyEkmConnectionServiceResolvers *EkmConnectionServiceResolvers = &EkmConnectionServiceResolvers{empty: true}

func (r *EkmConnectionServiceResolvers) Empty() bool {
	return r.empty
}

func (r *EkmConnectionServiceResolvers) String() string {
	return dcl.SprintResource(r)
}

func (r *EkmConnectionServiceResolvers) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type EkmConnectionServiceResolversServerCertificates struct {
	empty                      bool     `json:"-"`
	RawDer                     *string  `json:"rawDer"`
	Parsed                     *bool    `json:"parsed"`
	Issuer                     *string  `json:"issuer"`
	Subject                    *string  `json:"subject"`
	SubjectAlternativeDnsNames []string `json:"subjectAlternativeDnsNames"`
	NotBeforeTime              *string  `json:"notBeforeTime"`
	NotAfterTime               *string  `json:"notAfterTime"`
	SerialNumber               *string  `json:"serialNumber"`
	Sha256Fingerprint          *string  `json:"sha256Fingerprint"`
}

type jsonEkmConnectionServiceResolversServerCertificates EkmConnectionServiceResolversServerCertificates

func (r *EkmConnectionServiceResolversServerCertificates) UnmarshalJSON(data []byte) error {
	var res jsonEkmConnectionServiceResolversServerCertificates
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEkmConnectionServiceResolversServerCertificates
	} else {

		r.RawDer = res.RawDer

		r.Parsed = res.Parsed

		r.Issuer = res.Issuer

		r.Subject = res.Subject

		r.SubjectAlternativeDnsNames = res.SubjectAlternativeDnsNames

		r.NotBeforeTime = res.NotBeforeTime

		r.NotAfterTime = res.NotAfterTime

		r.SerialNumber = res.SerialNumber

		r.Sha256Fingerprint = res.Sha256Fingerprint

	}
	return nil
}

// This object is used to assert a desired state where this EkmConnectionServiceResolversServerCertificates is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyEkmConnectionServiceResolversServerCertificates *EkmConnectionServiceResolversServerCertificates = &EkmConnectionServiceResolversServerCertificates{empty: true}

func (r *EkmConnectionServiceResolversServerCertificates) Empty() bool {
	return r.empty
}

func (r *EkmConnectionServiceResolversServerCertificates) String() string {
	return dcl.SprintResource(r)
}

func (r *EkmConnectionServiceResolversServerCertificates) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *EkmConnection) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "cloudkms",
		Type:    "EkmConnection",
		Version: "cloudkms",
	}
}

func (r *EkmConnection) ID() (string, error) {
	if err := extractEkmConnectionFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":              dcl.ValueOrEmptyString(nr.Name),
		"create_time":       dcl.ValueOrEmptyString(nr.CreateTime),
		"service_resolvers": dcl.ValueOrEmptyString(nr.ServiceResolvers),
		"etag":              dcl.ValueOrEmptyString(nr.Etag),
		"project":           dcl.ValueOrEmptyString(nr.Project),
		"location":          dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/ekmConnections/{{name}}", params), nil
}

const EkmConnectionMaxPage = -1

type EkmConnectionList struct {
	Items []*EkmConnection

	nextToken string

	pageSize int32

	resource *EkmConnection
}

func (l *EkmConnectionList) HasNext() bool {
	return l.nextToken != ""
}

func (l *EkmConnectionList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listEkmConnection(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListEkmConnection(ctx context.Context, project, location string) (*EkmConnectionList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListEkmConnectionWithMaxResults(ctx, project, location, EkmConnectionMaxPage)

}

func (c *Client) ListEkmConnectionWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*EkmConnectionList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &EkmConnection{
		Project:  &project,
		Location: &location,
	}
	items, token, err := c.listEkmConnection(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &EkmConnectionList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetEkmConnection(ctx context.Context, r *EkmConnection) (*EkmConnection, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractEkmConnectionFields(r)

	b, err := c.getEkmConnectionRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalEkmConnection(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeEkmConnectionNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractEkmConnectionFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) ApplyEkmConnection(ctx context.Context, rawDesired *EkmConnection, opts ...dcl.ApplyOption) (*EkmConnection, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *EkmConnection
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyEkmConnectionHelper(c, ctx, rawDesired, opts...)
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

func applyEkmConnectionHelper(c *Client, ctx context.Context, rawDesired *EkmConnection, opts ...dcl.ApplyOption) (*EkmConnection, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyEkmConnection...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractEkmConnectionFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.ekmConnectionDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToEkmConnectionDiffs(c.Config, fieldDiffs, opts)
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
	var ops []ekmConnectionApiOperation
	if create {
		ops = append(ops, &createEkmConnectionOperation{})
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
	return applyEkmConnectionDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyEkmConnectionDiff(c *Client, ctx context.Context, desired *EkmConnection, rawDesired *EkmConnection, ops []ekmConnectionApiOperation, opts ...dcl.ApplyOption) (*EkmConnection, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetEkmConnection(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createEkmConnectionOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapEkmConnection(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeEkmConnectionNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeEkmConnectionNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeEkmConnectionDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractEkmConnectionFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractEkmConnectionFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffEkmConnection(c, newDesired, newState)
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

func (r *EkmConnection) GetPolicy(basePath string) (string, string, *bytes.Buffer, error) {
	u := r.getPolicyURL(basePath)
	body := &bytes.Buffer{}
	u, err := dcl.AddQueryParams(u, map[string]string{"options.requestedPolicyVersion": fmt.Sprintf("%d", r.IAMPolicyVersion())})
	if err != nil {
		return "", "", nil, err
	}
	return u, "GET", body, nil
}
