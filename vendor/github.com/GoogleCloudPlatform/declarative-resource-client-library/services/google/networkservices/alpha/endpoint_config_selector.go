// Copyright 2021 Google LLC. All Rights Reserved.
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

type EndpointConfigSelector struct {
	Name                *string                                    `json:"name"`
	CreateTime          *string                                    `json:"createTime"`
	UpdateTime          *string                                    `json:"updateTime"`
	Labels              map[string]string                          `json:"labels"`
	Type                *EndpointConfigSelectorTypeEnum            `json:"type"`
	AuthorizationPolicy *string                                    `json:"authorizationPolicy"`
	HttpFilters         *EndpointConfigSelectorHttpFilters         `json:"httpFilters"`
	EndpointMatcher     *EndpointConfigSelectorEndpointMatcher     `json:"endpointMatcher"`
	TrafficPortSelector *EndpointConfigSelectorTrafficPortSelector `json:"trafficPortSelector"`
	Description         *string                                    `json:"description"`
	ServerTlsPolicy     *string                                    `json:"serverTlsPolicy"`
	ClientTlsPolicy     *string                                    `json:"clientTlsPolicy"`
	Project             *string                                    `json:"project"`
	Location            *string                                    `json:"location"`
}

func (r *EndpointConfigSelector) String() string {
	return dcl.SprintResource(r)
}

// The enum EndpointConfigSelectorTypeEnum.
type EndpointConfigSelectorTypeEnum string

// EndpointConfigSelectorTypeEnumRef returns a *EndpointConfigSelectorTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func EndpointConfigSelectorTypeEnumRef(s string) *EndpointConfigSelectorTypeEnum {
	if s == "" {
		return nil
	}

	v := EndpointConfigSelectorTypeEnum(s)
	return &v
}

func (v EndpointConfigSelectorTypeEnum) Validate() error {
	for _, s := range []string{"ENDPOINT_CONFIG_SELECTOR_TYPE_UNSPECIFIED", "SIDECAR_PROXY", "GRPC_SERVER"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "EndpointConfigSelectorTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum.
type EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum string

// EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumRef returns a *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum with the value of string s
// If the empty string is provided, nil is returned.
func EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumRef(s string) *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum {
	if s == "" {
		return nil
	}

	v := EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(s)
	return &v
}

func (v EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum) Validate() error {
	for _, s := range []string{"METADATA_LABEL_MATCH_CRITERIA_UNSPECIFIED", "MATCH_ANY", "MATCH_ALL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type EndpointConfigSelectorHttpFilters struct {
	empty       bool     `json:"-"`
	HttpFilters []string `json:"httpFilters"`
}

type jsonEndpointConfigSelectorHttpFilters EndpointConfigSelectorHttpFilters

func (r *EndpointConfigSelectorHttpFilters) UnmarshalJSON(data []byte) error {
	var res jsonEndpointConfigSelectorHttpFilters
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEndpointConfigSelectorHttpFilters
	} else {

		r.HttpFilters = res.HttpFilters

	}
	return nil
}

// This object is used to assert a desired state where this EndpointConfigSelectorHttpFilters is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyEndpointConfigSelectorHttpFilters *EndpointConfigSelectorHttpFilters = &EndpointConfigSelectorHttpFilters{empty: true}

func (r *EndpointConfigSelectorHttpFilters) Empty() bool {
	return r.empty
}

func (r *EndpointConfigSelectorHttpFilters) String() string {
	return dcl.SprintResource(r)
}

func (r *EndpointConfigSelectorHttpFilters) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type EndpointConfigSelectorEndpointMatcher struct {
	empty                bool                                                       `json:"-"`
	MetadataLabelMatcher *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher `json:"metadataLabelMatcher"`
}

type jsonEndpointConfigSelectorEndpointMatcher EndpointConfigSelectorEndpointMatcher

func (r *EndpointConfigSelectorEndpointMatcher) UnmarshalJSON(data []byte) error {
	var res jsonEndpointConfigSelectorEndpointMatcher
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEndpointConfigSelectorEndpointMatcher
	} else {

		r.MetadataLabelMatcher = res.MetadataLabelMatcher

	}
	return nil
}

// This object is used to assert a desired state where this EndpointConfigSelectorEndpointMatcher is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyEndpointConfigSelectorEndpointMatcher *EndpointConfigSelectorEndpointMatcher = &EndpointConfigSelectorEndpointMatcher{empty: true}

func (r *EndpointConfigSelectorEndpointMatcher) Empty() bool {
	return r.empty
}

func (r *EndpointConfigSelectorEndpointMatcher) String() string {
	return dcl.SprintResource(r)
}

func (r *EndpointConfigSelectorEndpointMatcher) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher struct {
	empty                      bool                                                                                     `json:"-"`
	MetadataLabelMatchCriteria *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum `json:"metadataLabelMatchCriteria"`
	MetadataLabels             []EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels                `json:"metadataLabels"`
}

type jsonEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher

func (r *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher) UnmarshalJSON(data []byte) error {
	var res jsonEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher
	} else {

		r.MetadataLabelMatchCriteria = res.MetadataLabelMatchCriteria

		r.MetadataLabels = res.MetadataLabels

	}
	return nil
}

// This object is used to assert a desired state where this EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher = &EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher{empty: true}

func (r *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher) Empty() bool {
	return r.empty
}

func (r *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher) String() string {
	return dcl.SprintResource(r)
}

func (r *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels struct {
	empty      bool    `json:"-"`
	LabelName  *string `json:"labelName"`
	LabelValue *string `json:"labelValue"`
}

type jsonEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels

func (r *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels) UnmarshalJSON(data []byte) error {
	var res jsonEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels
	} else {

		r.LabelName = res.LabelName

		r.LabelValue = res.LabelValue

	}
	return nil
}

// This object is used to assert a desired state where this EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels = &EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels{empty: true}

func (r *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels) Empty() bool {
	return r.empty
}

func (r *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels) String() string {
	return dcl.SprintResource(r)
}

func (r *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type EndpointConfigSelectorTrafficPortSelector struct {
	empty bool     `json:"-"`
	Ports []string `json:"ports"`
}

type jsonEndpointConfigSelectorTrafficPortSelector EndpointConfigSelectorTrafficPortSelector

func (r *EndpointConfigSelectorTrafficPortSelector) UnmarshalJSON(data []byte) error {
	var res jsonEndpointConfigSelectorTrafficPortSelector
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyEndpointConfigSelectorTrafficPortSelector
	} else {

		r.Ports = res.Ports

	}
	return nil
}

// This object is used to assert a desired state where this EndpointConfigSelectorTrafficPortSelector is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyEndpointConfigSelectorTrafficPortSelector *EndpointConfigSelectorTrafficPortSelector = &EndpointConfigSelectorTrafficPortSelector{empty: true}

func (r *EndpointConfigSelectorTrafficPortSelector) Empty() bool {
	return r.empty
}

func (r *EndpointConfigSelectorTrafficPortSelector) String() string {
	return dcl.SprintResource(r)
}

func (r *EndpointConfigSelectorTrafficPortSelector) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *EndpointConfigSelector) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "network_services",
		Type:    "EndpointConfigSelector",
		Version: "alpha",
	}
}

func (r *EndpointConfigSelector) ID() (string, error) {
	if err := extractEndpointConfigSelectorFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                dcl.ValueOrEmptyString(nr.Name),
		"createTime":          dcl.ValueOrEmptyString(nr.CreateTime),
		"updateTime":          dcl.ValueOrEmptyString(nr.UpdateTime),
		"labels":              dcl.ValueOrEmptyString(nr.Labels),
		"type":                dcl.ValueOrEmptyString(nr.Type),
		"authorizationPolicy": dcl.ValueOrEmptyString(nr.AuthorizationPolicy),
		"httpFilters":         dcl.ValueOrEmptyString(nr.HttpFilters),
		"endpointMatcher":     dcl.ValueOrEmptyString(nr.EndpointMatcher),
		"trafficPortSelector": dcl.ValueOrEmptyString(nr.TrafficPortSelector),
		"description":         dcl.ValueOrEmptyString(nr.Description),
		"serverTlsPolicy":     dcl.ValueOrEmptyString(nr.ServerTlsPolicy),
		"clientTlsPolicy":     dcl.ValueOrEmptyString(nr.ClientTlsPolicy),
		"project":             dcl.ValueOrEmptyString(nr.Project),
		"location":            dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/endpointConfigSelectors/{{name}}", params), nil
}

const EndpointConfigSelectorMaxPage = -1

type EndpointConfigSelectorList struct {
	Items []*EndpointConfigSelector

	nextToken string

	pageSize int32

	resource *EndpointConfigSelector
}

func (l *EndpointConfigSelectorList) HasNext() bool {
	return l.nextToken != ""
}

func (l *EndpointConfigSelectorList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listEndpointConfigSelector(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListEndpointConfigSelector(ctx context.Context, r *EndpointConfigSelector) (*EndpointConfigSelectorList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListEndpointConfigSelectorWithMaxResults(ctx, r, EndpointConfigSelectorMaxPage)

}

func (c *Client) ListEndpointConfigSelectorWithMaxResults(ctx context.Context, r *EndpointConfigSelector, pageSize int32) (*EndpointConfigSelectorList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listEndpointConfigSelector(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &EndpointConfigSelectorList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetEndpointConfigSelector(ctx context.Context, r *EndpointConfigSelector) (*EndpointConfigSelector, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getEndpointConfigSelectorRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalEndpointConfigSelector(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeEndpointConfigSelectorNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteEndpointConfigSelector(ctx context.Context, r *EndpointConfigSelector) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("EndpointConfigSelector resource is nil")
	}
	c.Config.Logger.Info("Deleting EndpointConfigSelector...")
	deleteOp := deleteEndpointConfigSelectorOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllEndpointConfigSelector deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllEndpointConfigSelector(ctx context.Context, project, location string, filter func(*EndpointConfigSelector) bool) error {
	r := &EndpointConfigSelector{

		Project: &project,

		Location: &location,
	}
	listObj, err := c.ListEndpointConfigSelector(ctx, r)
	if err != nil {
		return err
	}

	err = c.deleteAllEndpointConfigSelector(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllEndpointConfigSelector(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyEndpointConfigSelector(ctx context.Context, rawDesired *EndpointConfigSelector, opts ...dcl.ApplyOption) (*EndpointConfigSelector, error) {
	var resultNewState *EndpointConfigSelector
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyEndpointConfigSelectorHelper(c, ctx, rawDesired, opts...)
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

func applyEndpointConfigSelectorHelper(c *Client, ctx context.Context, rawDesired *EndpointConfigSelector, opts ...dcl.ApplyOption) (*EndpointConfigSelector, error) {
	c.Config.Logger.Info("Beginning ApplyEndpointConfigSelector...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractEndpointConfigSelectorFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.endpointConfigSelectorDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToEndpointConfigSelectorDiffs(c.Config, fieldDiffs, opts)
	if err != nil {
		return nil, err
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
	var recreate bool
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
				if dcl.HasLifecycleParam(lp, dcl.BlockDestruction) || dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
					return nil, dcl.ApplyInfeasibleError{
						Message: fmt.Sprintf("Infeasible update: (%v) would require recreation.", d),
					}
				}
				recreate = true
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []endpointConfigSelectorApiOperation
	if create {
		ops = append(ops, &createEndpointConfigSelectorOperation{})
	} else if recreate {
		ops = append(ops, &deleteEndpointConfigSelectorOperation{})
		ops = append(ops, &createEndpointConfigSelectorOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeEndpointConfigSelectorDesiredState(rawDesired, nil)
		if err != nil {
			return nil, err
		}
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.Infof("Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.Infof("Performing operation %T %+v", op, op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.Infof("Failed operation %T %+v: %v", op, op, err)
			return nil, err
		}
		c.Config.Logger.Infof("Finished operation %T %+v", op, op)
	}

	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.Info("Retrieving raw new state...")
	rawNew, err := c.GetEndpointConfigSelector(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createEndpointConfigSelectorOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapEndpointConfigSelector(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeEndpointConfigSelectorNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeEndpointConfigSelectorNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeEndpointConfigSelectorDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffEndpointConfigSelector(c, newDesired, newState)
	if err != nil {
		return nil, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.Info("No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.Infof("Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.Info("Done Apply.")
	return newState, nil
}
