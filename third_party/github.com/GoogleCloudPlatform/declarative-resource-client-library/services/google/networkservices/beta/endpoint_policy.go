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

type EndpointPolicy struct {
	Name                *string                            `json:"name"`
	CreateTime          *string                            `json:"createTime"`
	UpdateTime          *string                            `json:"updateTime"`
	Labels              map[string]string                  `json:"labels"`
	Type                *EndpointPolicyTypeEnum            `json:"type"`
	AuthorizationPolicy *string                            `json:"authorizationPolicy"`
	EndpointMatcher     *EndpointPolicyEndpointMatcher     `json:"endpointMatcher"`
	TrafficPortSelector *EndpointPolicyTrafficPortSelector `json:"trafficPortSelector"`
	Description         *string                            `json:"description"`
	ServerTlsPolicy     *string                            `json:"serverTlsPolicy"`
	ClientTlsPolicy     *string                            `json:"clientTlsPolicy"`
	Project             *string                            `json:"project"`
	Location            *string                            `json:"location"`
}

func (r *EndpointPolicy) String() string {
	return dcl.SprintResource(r)
}

// The enum EndpointPolicyTypeEnum.
type EndpointPolicyTypeEnum string

// EndpointPolicyTypeEnumRef returns a *EndpointPolicyTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func EndpointPolicyTypeEnumRef(s string) *EndpointPolicyTypeEnum {
	v := EndpointPolicyTypeEnum(s)
	return &v
}

func (v EndpointPolicyTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"ENDPOINT_CONFIG_SELECTOR_TYPE_UNSPECIFIED", "SIDECAR_PROXY", "GRPC_SERVER"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "EndpointPolicyTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum.
type EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum string

// EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumRef returns a *EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum with the value of string s
// If the empty string is provided, nil is returned.
func EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumRef(s string) *EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum {
	v := EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(s)
	return &v
}

func (v EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"METADATA_LABEL_MATCH_CRITERIA_UNSPECIFIED", "MATCH_ANY", "MATCH_ALL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type EndpointPolicyEndpointMatcher struct {
	empty                bool                                               `json:"-"`
	MetadataLabelMatcher *EndpointPolicyEndpointMatcherMetadataLabelMatcher `json:"metadataLabelMatcher"`
}

type jsonEndpointPolicyEndpointMatcher EndpointPolicyEndpointMatcher

func (r *EndpointPolicyEndpointMatcher) UnmarshalJSON(data []byte) error {
	var res jsonEndpointPolicyEndpointMatcher
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEndpointPolicyEndpointMatcher
	} else {

		r.MetadataLabelMatcher = res.MetadataLabelMatcher

	}
	return nil
}

// This object is used to assert a desired state where this EndpointPolicyEndpointMatcher is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyEndpointPolicyEndpointMatcher *EndpointPolicyEndpointMatcher = &EndpointPolicyEndpointMatcher{empty: true}

func (r *EndpointPolicyEndpointMatcher) Empty() bool {
	return r.empty
}

func (r *EndpointPolicyEndpointMatcher) String() string {
	return dcl.SprintResource(r)
}

func (r *EndpointPolicyEndpointMatcher) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type EndpointPolicyEndpointMatcherMetadataLabelMatcher struct {
	empty                      bool                                                                             `json:"-"`
	MetadataLabelMatchCriteria *EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum `json:"metadataLabelMatchCriteria"`
	MetadataLabels             []EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels                `json:"metadataLabels"`
}

type jsonEndpointPolicyEndpointMatcherMetadataLabelMatcher EndpointPolicyEndpointMatcherMetadataLabelMatcher

func (r *EndpointPolicyEndpointMatcherMetadataLabelMatcher) UnmarshalJSON(data []byte) error {
	var res jsonEndpointPolicyEndpointMatcherMetadataLabelMatcher
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEndpointPolicyEndpointMatcherMetadataLabelMatcher
	} else {

		r.MetadataLabelMatchCriteria = res.MetadataLabelMatchCriteria

		r.MetadataLabels = res.MetadataLabels

	}
	return nil
}

// This object is used to assert a desired state where this EndpointPolicyEndpointMatcherMetadataLabelMatcher is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyEndpointPolicyEndpointMatcherMetadataLabelMatcher *EndpointPolicyEndpointMatcherMetadataLabelMatcher = &EndpointPolicyEndpointMatcherMetadataLabelMatcher{empty: true}

func (r *EndpointPolicyEndpointMatcherMetadataLabelMatcher) Empty() bool {
	return r.empty
}

func (r *EndpointPolicyEndpointMatcherMetadataLabelMatcher) String() string {
	return dcl.SprintResource(r)
}

func (r *EndpointPolicyEndpointMatcherMetadataLabelMatcher) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels struct {
	empty      bool    `json:"-"`
	LabelName  *string `json:"labelName"`
	LabelValue *string `json:"labelValue"`
}

type jsonEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels

func (r *EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels) UnmarshalJSON(data []byte) error {
	var res jsonEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels
	} else {

		r.LabelName = res.LabelName

		r.LabelValue = res.LabelValue

	}
	return nil
}

// This object is used to assert a desired state where this EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels *EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels = &EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels{empty: true}

func (r *EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels) Empty() bool {
	return r.empty
}

func (r *EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels) String() string {
	return dcl.SprintResource(r)
}

func (r *EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type EndpointPolicyTrafficPortSelector struct {
	empty bool     `json:"-"`
	Ports []string `json:"ports"`
}

type jsonEndpointPolicyTrafficPortSelector EndpointPolicyTrafficPortSelector

func (r *EndpointPolicyTrafficPortSelector) UnmarshalJSON(data []byte) error {
	var res jsonEndpointPolicyTrafficPortSelector
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEndpointPolicyTrafficPortSelector
	} else {

		r.Ports = res.Ports

	}
	return nil
}

// This object is used to assert a desired state where this EndpointPolicyTrafficPortSelector is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyEndpointPolicyTrafficPortSelector *EndpointPolicyTrafficPortSelector = &EndpointPolicyTrafficPortSelector{empty: true}

func (r *EndpointPolicyTrafficPortSelector) Empty() bool {
	return r.empty
}

func (r *EndpointPolicyTrafficPortSelector) String() string {
	return dcl.SprintResource(r)
}

func (r *EndpointPolicyTrafficPortSelector) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *EndpointPolicy) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "network_services",
		Type:    "EndpointPolicy",
		Version: "beta",
	}
}

func (r *EndpointPolicy) ID() (string, error) {
	if err := extractEndpointPolicyFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                  dcl.ValueOrEmptyString(nr.Name),
		"create_time":           dcl.ValueOrEmptyString(nr.CreateTime),
		"update_time":           dcl.ValueOrEmptyString(nr.UpdateTime),
		"labels":                dcl.ValueOrEmptyString(nr.Labels),
		"type":                  dcl.ValueOrEmptyString(nr.Type),
		"authorization_policy":  dcl.ValueOrEmptyString(nr.AuthorizationPolicy),
		"endpoint_matcher":      dcl.ValueOrEmptyString(nr.EndpointMatcher),
		"traffic_port_selector": dcl.ValueOrEmptyString(nr.TrafficPortSelector),
		"description":           dcl.ValueOrEmptyString(nr.Description),
		"server_tls_policy":     dcl.ValueOrEmptyString(nr.ServerTlsPolicy),
		"client_tls_policy":     dcl.ValueOrEmptyString(nr.ClientTlsPolicy),
		"project":               dcl.ValueOrEmptyString(nr.Project),
		"location":              dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/endpointPolicies/{{name}}", params), nil
}

const EndpointPolicyMaxPage = -1

type EndpointPolicyList struct {
	Items []*EndpointPolicy

	nextToken string

	pageSize int32

	resource *EndpointPolicy
}

func (l *EndpointPolicyList) HasNext() bool {
	return l.nextToken != ""
}

func (l *EndpointPolicyList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listEndpointPolicy(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListEndpointPolicy(ctx context.Context, project, location string) (*EndpointPolicyList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListEndpointPolicyWithMaxResults(ctx, project, location, EndpointPolicyMaxPage)

}

func (c *Client) ListEndpointPolicyWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*EndpointPolicyList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &EndpointPolicy{
		Project:  &project,
		Location: &location,
	}
	items, token, err := c.listEndpointPolicy(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &EndpointPolicyList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetEndpointPolicy(ctx context.Context, r *EndpointPolicy) (*EndpointPolicy, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractEndpointPolicyFields(r)

	b, err := c.getEndpointPolicyRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalEndpointPolicy(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeEndpointPolicyNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractEndpointPolicyFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteEndpointPolicy(ctx context.Context, r *EndpointPolicy) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("EndpointPolicy resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting EndpointPolicy...")
	deleteOp := deleteEndpointPolicyOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllEndpointPolicy deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllEndpointPolicy(ctx context.Context, project, location string, filter func(*EndpointPolicy) bool) error {
	listObj, err := c.ListEndpointPolicy(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllEndpointPolicy(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllEndpointPolicy(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyEndpointPolicy(ctx context.Context, rawDesired *EndpointPolicy, opts ...dcl.ApplyOption) (*EndpointPolicy, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *EndpointPolicy
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyEndpointPolicyHelper(c, ctx, rawDesired, opts...)
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

func applyEndpointPolicyHelper(c *Client, ctx context.Context, rawDesired *EndpointPolicy, opts ...dcl.ApplyOption) (*EndpointPolicy, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyEndpointPolicy...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractEndpointPolicyFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.endpointPolicyDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToEndpointPolicyDiffs(c.Config, fieldDiffs, opts)
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
	var ops []endpointPolicyApiOperation
	if create {
		ops = append(ops, &createEndpointPolicyOperation{})
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
	return applyEndpointPolicyDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyEndpointPolicyDiff(c *Client, ctx context.Context, desired *EndpointPolicy, rawDesired *EndpointPolicy, ops []endpointPolicyApiOperation, opts ...dcl.ApplyOption) (*EndpointPolicy, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetEndpointPolicy(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createEndpointPolicyOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapEndpointPolicy(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeEndpointPolicyNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeEndpointPolicyNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeEndpointPolicyDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractEndpointPolicyFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractEndpointPolicyFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffEndpointPolicy(c, newDesired, newState)
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
