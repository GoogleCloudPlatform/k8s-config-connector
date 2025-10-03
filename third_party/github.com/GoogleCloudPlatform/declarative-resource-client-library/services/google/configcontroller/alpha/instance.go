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
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Instance struct {
	Name               *string                   `json:"name"`
	Labels             map[string]string         `json:"labels"`
	BundlesConfig      *InstanceBundlesConfig    `json:"bundlesConfig"`
	UsePrivateEndpoint *bool                     `json:"usePrivateEndpoint"`
	GkeResourceLink    *string                   `json:"gkeResourceLink"`
	State              *InstanceStateEnum        `json:"state"`
	ManagementConfig   *InstanceManagementConfig `json:"managementConfig"`
	Project            *string                   `json:"project"`
	Location           *string                   `json:"location"`
}

func (r *Instance) String() string {
	return dcl.SprintResource(r)
}

// The enum InstanceStateEnum.
type InstanceStateEnum string

// InstanceStateEnumRef returns a *InstanceStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceStateEnumRef(s string) *InstanceStateEnum {
	v := InstanceStateEnum(s)
	return &v
}

func (v InstanceStateEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"STATE_UNSPECIFIED", "CREATING", "RUNNING", "DELETING", "SUSPENDED", "READ_ONLY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type InstanceBundlesConfig struct {
	empty                  bool                                         `json:"-"`
	ConfigControllerConfig *InstanceBundlesConfigConfigControllerConfig `json:"configControllerConfig"`
}

type jsonInstanceBundlesConfig InstanceBundlesConfig

func (r *InstanceBundlesConfig) UnmarshalJSON(data []byte) error {
	var res jsonInstanceBundlesConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInstanceBundlesConfig
	} else {

		r.ConfigControllerConfig = res.ConfigControllerConfig

	}
	return nil
}

// This object is used to assert a desired state where this InstanceBundlesConfig is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInstanceBundlesConfig *InstanceBundlesConfig = &InstanceBundlesConfig{empty: true}

func (r *InstanceBundlesConfig) Empty() bool {
	return r.empty
}

func (r *InstanceBundlesConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceBundlesConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceBundlesConfigConfigControllerConfig struct {
	empty   bool  `json:"-"`
	Enabled *bool `json:"enabled"`
}

type jsonInstanceBundlesConfigConfigControllerConfig InstanceBundlesConfigConfigControllerConfig

func (r *InstanceBundlesConfigConfigControllerConfig) UnmarshalJSON(data []byte) error {
	var res jsonInstanceBundlesConfigConfigControllerConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInstanceBundlesConfigConfigControllerConfig
	} else {

		r.Enabled = res.Enabled

	}
	return nil
}

// This object is used to assert a desired state where this InstanceBundlesConfigConfigControllerConfig is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInstanceBundlesConfigConfigControllerConfig *InstanceBundlesConfigConfigControllerConfig = &InstanceBundlesConfigConfigControllerConfig{empty: true}

func (r *InstanceBundlesConfigConfigControllerConfig) Empty() bool {
	return r.empty
}

func (r *InstanceBundlesConfigConfigControllerConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceBundlesConfigConfigControllerConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceManagementConfig struct {
	empty                    bool                                              `json:"-"`
	StandardManagementConfig *InstanceManagementConfigStandardManagementConfig `json:"standardManagementConfig"`
	FullManagementConfig     *InstanceManagementConfigFullManagementConfig     `json:"fullManagementConfig"`
}

type jsonInstanceManagementConfig InstanceManagementConfig

func (r *InstanceManagementConfig) UnmarshalJSON(data []byte) error {
	var res jsonInstanceManagementConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInstanceManagementConfig
	} else {

		r.StandardManagementConfig = res.StandardManagementConfig

		r.FullManagementConfig = res.FullManagementConfig

	}
	return nil
}

// This object is used to assert a desired state where this InstanceManagementConfig is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInstanceManagementConfig *InstanceManagementConfig = &InstanceManagementConfig{empty: true}

func (r *InstanceManagementConfig) Empty() bool {
	return r.empty
}

func (r *InstanceManagementConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceManagementConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceManagementConfigStandardManagementConfig struct {
	empty               bool    `json:"-"`
	Network             *string `json:"network"`
	MasterIPv4CidrBlock *string `json:"masterIPv4CidrBlock"`
	ManBlock            *string `json:"manBlock"`
	ClusterCidrBlock    *string `json:"clusterCidrBlock"`
	ServicesCidrBlock   *string `json:"servicesCidrBlock"`
	ClusterNamedRange   *string `json:"clusterNamedRange"`
	ServicesNamedRange  *string `json:"servicesNamedRange"`
}

type jsonInstanceManagementConfigStandardManagementConfig InstanceManagementConfigStandardManagementConfig

func (r *InstanceManagementConfigStandardManagementConfig) UnmarshalJSON(data []byte) error {
	var res jsonInstanceManagementConfigStandardManagementConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInstanceManagementConfigStandardManagementConfig
	} else {

		r.Network = res.Network

		r.MasterIPv4CidrBlock = res.MasterIPv4CidrBlock

		r.ManBlock = res.ManBlock

		r.ClusterCidrBlock = res.ClusterCidrBlock

		r.ServicesCidrBlock = res.ServicesCidrBlock

		r.ClusterNamedRange = res.ClusterNamedRange

		r.ServicesNamedRange = res.ServicesNamedRange

	}
	return nil
}

// This object is used to assert a desired state where this InstanceManagementConfigStandardManagementConfig is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInstanceManagementConfigStandardManagementConfig *InstanceManagementConfigStandardManagementConfig = &InstanceManagementConfigStandardManagementConfig{empty: true}

func (r *InstanceManagementConfigStandardManagementConfig) Empty() bool {
	return r.empty
}

func (r *InstanceManagementConfigStandardManagementConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceManagementConfigStandardManagementConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceManagementConfigFullManagementConfig struct {
	empty               bool    `json:"-"`
	Network             *string `json:"network"`
	MasterIPv4CidrBlock *string `json:"masterIPv4CidrBlock"`
	ManBlock            *string `json:"manBlock"`
	ClusterCidrBlock    *string `json:"clusterCidrBlock"`
	ServicesCidrBlock   *string `json:"servicesCidrBlock"`
	ClusterNamedRange   *string `json:"clusterNamedRange"`
	ServicesNamedRange  *string `json:"servicesNamedRange"`
}

type jsonInstanceManagementConfigFullManagementConfig InstanceManagementConfigFullManagementConfig

func (r *InstanceManagementConfigFullManagementConfig) UnmarshalJSON(data []byte) error {
	var res jsonInstanceManagementConfigFullManagementConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInstanceManagementConfigFullManagementConfig
	} else {

		r.Network = res.Network

		r.MasterIPv4CidrBlock = res.MasterIPv4CidrBlock

		r.ManBlock = res.ManBlock

		r.ClusterCidrBlock = res.ClusterCidrBlock

		r.ServicesCidrBlock = res.ServicesCidrBlock

		r.ClusterNamedRange = res.ClusterNamedRange

		r.ServicesNamedRange = res.ServicesNamedRange

	}
	return nil
}

// This object is used to assert a desired state where this InstanceManagementConfigFullManagementConfig is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInstanceManagementConfigFullManagementConfig *InstanceManagementConfigFullManagementConfig = &InstanceManagementConfigFullManagementConfig{empty: true}

func (r *InstanceManagementConfigFullManagementConfig) Empty() bool {
	return r.empty
}

func (r *InstanceManagementConfigFullManagementConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceManagementConfigFullManagementConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Instance) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "config_controller",
		Type:    "Instance",
		Version: "alpha",
	}
}

func (r *Instance) ID() (string, error) {
	if err := extractInstanceFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                 dcl.ValueOrEmptyString(nr.Name),
		"labels":               dcl.ValueOrEmptyString(nr.Labels),
		"bundles_config":       dcl.ValueOrEmptyString(nr.BundlesConfig),
		"use_private_endpoint": dcl.ValueOrEmptyString(nr.UsePrivateEndpoint),
		"gke_resource_link":    dcl.ValueOrEmptyString(nr.GkeResourceLink),
		"state":                dcl.ValueOrEmptyString(nr.State),
		"management_config":    dcl.ValueOrEmptyString(nr.ManagementConfig),
		"project":              dcl.ValueOrEmptyString(nr.Project),
		"location":             dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/krmApiHosts/{{name}}", params), nil
}

const InstanceMaxPage = -1

type InstanceList struct {
	Items []*Instance

	nextToken string

	pageSize int32

	resource *Instance
}

func (l *InstanceList) HasNext() bool {
	return l.nextToken != ""
}

func (l *InstanceList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listInstance(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListInstance(ctx context.Context, project, location string) (*InstanceList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListInstanceWithMaxResults(ctx, project, location, InstanceMaxPage)

}

func (c *Client) ListInstanceWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*InstanceList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &Instance{
		Project:  &project,
		Location: &location,
	}
	items, token, err := c.listInstance(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &InstanceList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetInstance(ctx context.Context, r *Instance) (*Instance, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractInstanceFields(r)

	b, err := c.getInstanceRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalInstance(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeInstanceNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractInstanceFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteInstance(ctx context.Context, r *Instance) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Instance resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Instance...")
	deleteOp := deleteInstanceOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllInstance deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllInstance(ctx context.Context, project, location string, filter func(*Instance) bool) error {
	listObj, err := c.ListInstance(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllInstance(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllInstance(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyInstance(ctx context.Context, rawDesired *Instance, opts ...dcl.ApplyOption) (*Instance, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Instance
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyInstanceHelper(c, ctx, rawDesired, opts...)
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

func applyInstanceHelper(c *Client, ctx context.Context, rawDesired *Instance, opts ...dcl.ApplyOption) (*Instance, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyInstance...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractInstanceFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.instanceDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToInstanceDiffs(c.Config, fieldDiffs, opts)
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
	var ops []instanceApiOperation
	if create {
		ops = append(ops, &createInstanceOperation{})
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
	return applyInstanceDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyInstanceDiff(c *Client, ctx context.Context, desired *Instance, rawDesired *Instance, ops []instanceApiOperation, opts ...dcl.ApplyOption) (*Instance, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetInstance(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createInstanceOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapInstance(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeInstanceNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeInstanceNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeInstanceDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractInstanceFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractInstanceFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffInstance(c, newDesired, newState)
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

func (r *Instance) GetPolicy(basePath string) (string, string, *bytes.Buffer, error) {
	u := r.getPolicyURL(basePath)
	body := &bytes.Buffer{}
	u, err := dcl.AddQueryParams(u, map[string]string{"optionsRequestedPolicyVersion": fmt.Sprintf("%d", r.IAMPolicyVersion())})
	if err != nil {
		return "", "", nil, err
	}
	return u, "", body, nil
}
