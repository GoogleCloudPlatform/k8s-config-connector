// Copyright 2023 Google LLC. All Rights Reserved.
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

type PrivateCloud struct {
	Name              *string                        `json:"name"`
	CreateTime        *string                        `json:"createTime"`
	UpdateTime        *string                        `json:"updateTime"`
	DeleteTime        *string                        `json:"deleteTime"`
	ExpireTime        *string                        `json:"expireTime"`
	State             *PrivateCloudStateEnum         `json:"state"`
	NetworkConfig     *PrivateCloudNetworkConfig     `json:"networkConfig"`
	ManagementCluster *PrivateCloudManagementCluster `json:"managementCluster"`
	Description       *string                        `json:"description"`
	Hcx               *PrivateCloudHcx               `json:"hcx"`
	Nsx               *PrivateCloudNsx               `json:"nsx"`
	Vcenter           *PrivateCloudVcenter           `json:"vcenter"`
	Uid               *string                        `json:"uid"`
	Project           *string                        `json:"project"`
	Location          *string                        `json:"location"`
}

func (r *PrivateCloud) String() string {
	return dcl.SprintResource(r)
}

// The enum PrivateCloudStateEnum.
type PrivateCloudStateEnum string

// PrivateCloudStateEnumRef returns a *PrivateCloudStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func PrivateCloudStateEnumRef(s string) *PrivateCloudStateEnum {
	v := PrivateCloudStateEnum(s)
	return &v
}

func (v PrivateCloudStateEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"STATE_UNSPECIFIED", "ACTIVE", "CREATING", "UPDATING", "FAILED", "DELETED", "PURGING"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "PrivateCloudStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum PrivateCloudHcxStateEnum.
type PrivateCloudHcxStateEnum string

// PrivateCloudHcxStateEnumRef returns a *PrivateCloudHcxStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func PrivateCloudHcxStateEnumRef(s string) *PrivateCloudHcxStateEnum {
	v := PrivateCloudHcxStateEnum(s)
	return &v
}

func (v PrivateCloudHcxStateEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"STATE_UNSPECIFIED", "ACTIVE", "CREATING"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "PrivateCloudHcxStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum PrivateCloudNsxStateEnum.
type PrivateCloudNsxStateEnum string

// PrivateCloudNsxStateEnumRef returns a *PrivateCloudNsxStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func PrivateCloudNsxStateEnumRef(s string) *PrivateCloudNsxStateEnum {
	v := PrivateCloudNsxStateEnum(s)
	return &v
}

func (v PrivateCloudNsxStateEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"STATE_UNSPECIFIED", "ACTIVE", "CREATING"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "PrivateCloudNsxStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum PrivateCloudVcenterStateEnum.
type PrivateCloudVcenterStateEnum string

// PrivateCloudVcenterStateEnumRef returns a *PrivateCloudVcenterStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func PrivateCloudVcenterStateEnumRef(s string) *PrivateCloudVcenterStateEnum {
	v := PrivateCloudVcenterStateEnum(s)
	return &v
}

func (v PrivateCloudVcenterStateEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"STATE_UNSPECIFIED", "ACTIVE", "CREATING"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "PrivateCloudVcenterStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type PrivateCloudNetworkConfig struct {
	empty                            bool    `json:"-"`
	ManagementCidr                   *string `json:"managementCidr"`
	VmwareEngineNetwork              *string `json:"vmwareEngineNetwork"`
	VmwareEngineNetworkCanonical     *string `json:"vmwareEngineNetworkCanonical"`
	ManagementIPAddressLayoutVersion *int64  `json:"managementIPAddressLayoutVersion"`
}

type jsonPrivateCloudNetworkConfig PrivateCloudNetworkConfig

func (r *PrivateCloudNetworkConfig) UnmarshalJSON(data []byte) error {
	var res jsonPrivateCloudNetworkConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyPrivateCloudNetworkConfig
	} else {

		r.ManagementCidr = res.ManagementCidr

		r.VmwareEngineNetwork = res.VmwareEngineNetwork

		r.VmwareEngineNetworkCanonical = res.VmwareEngineNetworkCanonical

		r.ManagementIPAddressLayoutVersion = res.ManagementIPAddressLayoutVersion

	}
	return nil
}

// This object is used to assert a desired state where this PrivateCloudNetworkConfig is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyPrivateCloudNetworkConfig *PrivateCloudNetworkConfig = &PrivateCloudNetworkConfig{empty: true}

func (r *PrivateCloudNetworkConfig) Empty() bool {
	return r.empty
}

func (r *PrivateCloudNetworkConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *PrivateCloudNetworkConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type PrivateCloudManagementCluster struct {
	empty     bool    `json:"-"`
	ClusterId *string `json:"clusterId"`
}

type jsonPrivateCloudManagementCluster PrivateCloudManagementCluster

func (r *PrivateCloudManagementCluster) UnmarshalJSON(data []byte) error {
	var res jsonPrivateCloudManagementCluster
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyPrivateCloudManagementCluster
	} else {

		r.ClusterId = res.ClusterId

	}
	return nil
}

// This object is used to assert a desired state where this PrivateCloudManagementCluster is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyPrivateCloudManagementCluster *PrivateCloudManagementCluster = &PrivateCloudManagementCluster{empty: true}

func (r *PrivateCloudManagementCluster) Empty() bool {
	return r.empty
}

func (r *PrivateCloudManagementCluster) String() string {
	return dcl.SprintResource(r)
}

func (r *PrivateCloudManagementCluster) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type PrivateCloudHcx struct {
	empty      bool                      `json:"-"`
	InternalIP *string                   `json:"internalIP"`
	Version    *string                   `json:"version"`
	State      *PrivateCloudHcxStateEnum `json:"state"`
	Fqdn       *string                   `json:"fqdn"`
}

type jsonPrivateCloudHcx PrivateCloudHcx

func (r *PrivateCloudHcx) UnmarshalJSON(data []byte) error {
	var res jsonPrivateCloudHcx
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyPrivateCloudHcx
	} else {

		r.InternalIP = res.InternalIP

		r.Version = res.Version

		r.State = res.State

		r.Fqdn = res.Fqdn

	}
	return nil
}

// This object is used to assert a desired state where this PrivateCloudHcx is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyPrivateCloudHcx *PrivateCloudHcx = &PrivateCloudHcx{empty: true}

func (r *PrivateCloudHcx) Empty() bool {
	return r.empty
}

func (r *PrivateCloudHcx) String() string {
	return dcl.SprintResource(r)
}

func (r *PrivateCloudHcx) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type PrivateCloudNsx struct {
	empty      bool                      `json:"-"`
	InternalIP *string                   `json:"internalIP"`
	Version    *string                   `json:"version"`
	State      *PrivateCloudNsxStateEnum `json:"state"`
	Fqdn       *string                   `json:"fqdn"`
}

type jsonPrivateCloudNsx PrivateCloudNsx

func (r *PrivateCloudNsx) UnmarshalJSON(data []byte) error {
	var res jsonPrivateCloudNsx
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyPrivateCloudNsx
	} else {

		r.InternalIP = res.InternalIP

		r.Version = res.Version

		r.State = res.State

		r.Fqdn = res.Fqdn

	}
	return nil
}

// This object is used to assert a desired state where this PrivateCloudNsx is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyPrivateCloudNsx *PrivateCloudNsx = &PrivateCloudNsx{empty: true}

func (r *PrivateCloudNsx) Empty() bool {
	return r.empty
}

func (r *PrivateCloudNsx) String() string {
	return dcl.SprintResource(r)
}

func (r *PrivateCloudNsx) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type PrivateCloudVcenter struct {
	empty      bool                          `json:"-"`
	InternalIP *string                       `json:"internalIP"`
	Version    *string                       `json:"version"`
	State      *PrivateCloudVcenterStateEnum `json:"state"`
	Fqdn       *string                       `json:"fqdn"`
}

type jsonPrivateCloudVcenter PrivateCloudVcenter

func (r *PrivateCloudVcenter) UnmarshalJSON(data []byte) error {
	var res jsonPrivateCloudVcenter
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyPrivateCloudVcenter
	} else {

		r.InternalIP = res.InternalIP

		r.Version = res.Version

		r.State = res.State

		r.Fqdn = res.Fqdn

	}
	return nil
}

// This object is used to assert a desired state where this PrivateCloudVcenter is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyPrivateCloudVcenter *PrivateCloudVcenter = &PrivateCloudVcenter{empty: true}

func (r *PrivateCloudVcenter) Empty() bool {
	return r.empty
}

func (r *PrivateCloudVcenter) String() string {
	return dcl.SprintResource(r)
}

func (r *PrivateCloudVcenter) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *PrivateCloud) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "vmware",
		Type:    "PrivateCloud",
		Version: "alpha",
	}
}

func (r *PrivateCloud) ID() (string, error) {
	if err := extractPrivateCloudFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":               dcl.ValueOrEmptyString(nr.Name),
		"create_time":        dcl.ValueOrEmptyString(nr.CreateTime),
		"update_time":        dcl.ValueOrEmptyString(nr.UpdateTime),
		"delete_time":        dcl.ValueOrEmptyString(nr.DeleteTime),
		"expire_time":        dcl.ValueOrEmptyString(nr.ExpireTime),
		"state":              dcl.ValueOrEmptyString(nr.State),
		"network_config":     dcl.ValueOrEmptyString(nr.NetworkConfig),
		"management_cluster": dcl.ValueOrEmptyString(nr.ManagementCluster),
		"description":        dcl.ValueOrEmptyString(nr.Description),
		"hcx":                dcl.ValueOrEmptyString(nr.Hcx),
		"nsx":                dcl.ValueOrEmptyString(nr.Nsx),
		"vcenter":            dcl.ValueOrEmptyString(nr.Vcenter),
		"uid":                dcl.ValueOrEmptyString(nr.Uid),
		"project":            dcl.ValueOrEmptyString(nr.Project),
		"location":           dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/privateClouds/{{name}}", params), nil
}

const PrivateCloudMaxPage = -1

type PrivateCloudList struct {
	Items []*PrivateCloud

	nextToken string

	pageSize int32

	resource *PrivateCloud
}

func (l *PrivateCloudList) HasNext() bool {
	return l.nextToken != ""
}

func (l *PrivateCloudList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listPrivateCloud(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListPrivateCloud(ctx context.Context, project, location string) (*PrivateCloudList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListPrivateCloudWithMaxResults(ctx, project, location, PrivateCloudMaxPage)

}

func (c *Client) ListPrivateCloudWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*PrivateCloudList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &PrivateCloud{
		Project:  &project,
		Location: &location,
	}
	items, token, err := c.listPrivateCloud(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &PrivateCloudList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetPrivateCloud(ctx context.Context, r *PrivateCloud) (*PrivateCloud, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractPrivateCloudFields(r)

	b, err := c.getPrivateCloudRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalPrivateCloud(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizePrivateCloudNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractPrivateCloudFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeletePrivateCloud(ctx context.Context, r *PrivateCloud) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(7200*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("PrivateCloud resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting PrivateCloud...")
	deleteOp := deletePrivateCloudOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllPrivateCloud deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllPrivateCloud(ctx context.Context, project, location string, filter func(*PrivateCloud) bool) error {
	listObj, err := c.ListPrivateCloud(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllPrivateCloud(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllPrivateCloud(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyPrivateCloud(ctx context.Context, rawDesired *PrivateCloud, opts ...dcl.ApplyOption) (*PrivateCloud, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(9600*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *PrivateCloud
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyPrivateCloudHelper(c, ctx, rawDesired, opts...)
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

func applyPrivateCloudHelper(c *Client, ctx context.Context, rawDesired *PrivateCloud, opts ...dcl.ApplyOption) (*PrivateCloud, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyPrivateCloud...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractPrivateCloudFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.privateCloudDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToPrivateCloudDiffs(c.Config, fieldDiffs, opts)
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
	var ops []privateCloudApiOperation
	if create {
		ops = append(ops, &createPrivateCloudOperation{})
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
	return applyPrivateCloudDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyPrivateCloudDiff(c *Client, ctx context.Context, desired *PrivateCloud, rawDesired *PrivateCloud, ops []privateCloudApiOperation, opts ...dcl.ApplyOption) (*PrivateCloud, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetPrivateCloud(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createPrivateCloudOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapPrivateCloud(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizePrivateCloudNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizePrivateCloudNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizePrivateCloudDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractPrivateCloudFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractPrivateCloudFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffPrivateCloud(c, newDesired, newState)
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

func (r *PrivateCloud) GetPolicy(basePath string) (string, string, *bytes.Buffer, error) {
	u := r.getPolicyURL(basePath)
	body := &bytes.Buffer{}
	u, err := dcl.AddQueryParams(u, map[string]string{"options.requestedPolicyVersion": fmt.Sprintf("%d", r.IAMPolicyVersion())})
	if err != nil {
		return "", "", nil, err
	}
	return u, "GET", body, nil
}
