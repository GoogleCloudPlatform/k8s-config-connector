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
package cloudfunctions

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

type Function struct {
	Name                       *string                                 `json:"name"`
	Description                *string                                 `json:"description"`
	SourceArchiveUrl           *string                                 `json:"sourceArchiveUrl"`
	SourceRepository           *FunctionSourceRepository               `json:"sourceRepository"`
	HttpsTrigger               *FunctionHttpsTrigger                   `json:"httpsTrigger"`
	EventTrigger               *FunctionEventTrigger                   `json:"eventTrigger"`
	Status                     *FunctionStatusEnum                     `json:"status"`
	EntryPoint                 *string                                 `json:"entryPoint"`
	Runtime                    *string                                 `json:"runtime"`
	Timeout                    *string                                 `json:"timeout"`
	AvailableMemoryMb          *int64                                  `json:"availableMemoryMb"`
	ServiceAccountEmail        *string                                 `json:"serviceAccountEmail"`
	UpdateTime                 *string                                 `json:"updateTime"`
	VersionId                  *int64                                  `json:"versionId"`
	Labels                     map[string]string                       `json:"labels"`
	EnvironmentVariables       map[string]string                       `json:"environmentVariables"`
	MaxInstances               *int64                                  `json:"maxInstances"`
	VPCConnector               *string                                 `json:"vpcConnector"`
	VPCConnectorEgressSettings *FunctionVPCConnectorEgressSettingsEnum `json:"vpcConnectorEgressSettings"`
	IngressSettings            *FunctionIngressSettingsEnum            `json:"ingressSettings"`
	Region                     *string                                 `json:"region"`
	Project                    *string                                 `json:"project"`
}

func (r *Function) String() string {
	return dcl.SprintResource(r)
}

// The enum FunctionHttpsTriggerSecurityLevelEnum.
type FunctionHttpsTriggerSecurityLevelEnum string

// FunctionHttpsTriggerSecurityLevelEnumRef returns a *FunctionHttpsTriggerSecurityLevelEnum with the value of string s
// If the empty string is provided, nil is returned.
func FunctionHttpsTriggerSecurityLevelEnumRef(s string) *FunctionHttpsTriggerSecurityLevelEnum {
	v := FunctionHttpsTriggerSecurityLevelEnum(s)
	return &v
}

func (v FunctionHttpsTriggerSecurityLevelEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"SECURITY_LEVEL_UNSPECIFIED", "SECURE_ALWAYS", "SECURE_OPTIONAL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "FunctionHttpsTriggerSecurityLevelEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum FunctionStatusEnum.
type FunctionStatusEnum string

// FunctionStatusEnumRef returns a *FunctionStatusEnum with the value of string s
// If the empty string is provided, nil is returned.
func FunctionStatusEnumRef(s string) *FunctionStatusEnum {
	v := FunctionStatusEnum(s)
	return &v
}

func (v FunctionStatusEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"CLOUD_FUNCTION_STATUS_UNSPECIFIED", "ACTIVE", "OFFLINE", "DEPLOY_IN_PROGRESS", "DELETE_IN_PROGRESS", "UNKNOWN"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "FunctionStatusEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum FunctionVPCConnectorEgressSettingsEnum.
type FunctionVPCConnectorEgressSettingsEnum string

// FunctionVPCConnectorEgressSettingsEnumRef returns a *FunctionVPCConnectorEgressSettingsEnum with the value of string s
// If the empty string is provided, nil is returned.
func FunctionVPCConnectorEgressSettingsEnumRef(s string) *FunctionVPCConnectorEgressSettingsEnum {
	v := FunctionVPCConnectorEgressSettingsEnum(s)
	return &v
}

func (v FunctionVPCConnectorEgressSettingsEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED", "PRIVATE_RANGES_ONLY", "ALL_TRAFFIC"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "FunctionVPCConnectorEgressSettingsEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum FunctionIngressSettingsEnum.
type FunctionIngressSettingsEnum string

// FunctionIngressSettingsEnumRef returns a *FunctionIngressSettingsEnum with the value of string s
// If the empty string is provided, nil is returned.
func FunctionIngressSettingsEnumRef(s string) *FunctionIngressSettingsEnum {
	v := FunctionIngressSettingsEnum(s)
	return &v
}

func (v FunctionIngressSettingsEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"INGRESS_SETTINGS_UNSPECIFIED", "ALLOW_ALL", "ALLOW_INTERNAL_ONLY", "ALLOW_INTERNAL_AND_GCLB"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "FunctionIngressSettingsEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type FunctionSourceRepository struct {
	empty       bool    `json:"-"`
	Url         *string `json:"url"`
	DeployedUrl *string `json:"deployedUrl"`
}

type jsonFunctionSourceRepository FunctionSourceRepository

func (r *FunctionSourceRepository) UnmarshalJSON(data []byte) error {
	var res jsonFunctionSourceRepository
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFunctionSourceRepository
	} else {

		r.Url = res.Url

		r.DeployedUrl = res.DeployedUrl

	}
	return nil
}

// This object is used to assert a desired state where this FunctionSourceRepository is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyFunctionSourceRepository *FunctionSourceRepository = &FunctionSourceRepository{empty: true}

func (r *FunctionSourceRepository) Empty() bool {
	return r.empty
}

func (r *FunctionSourceRepository) String() string {
	return dcl.SprintResource(r)
}

func (r *FunctionSourceRepository) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FunctionHttpsTrigger struct {
	empty         bool                                   `json:"-"`
	Url           *string                                `json:"url"`
	SecurityLevel *FunctionHttpsTriggerSecurityLevelEnum `json:"securityLevel"`
}

type jsonFunctionHttpsTrigger FunctionHttpsTrigger

func (r *FunctionHttpsTrigger) UnmarshalJSON(data []byte) error {
	var res jsonFunctionHttpsTrigger
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFunctionHttpsTrigger
	} else {

		r.Url = res.Url

		r.SecurityLevel = res.SecurityLevel

	}
	return nil
}

// This object is used to assert a desired state where this FunctionHttpsTrigger is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyFunctionHttpsTrigger *FunctionHttpsTrigger = &FunctionHttpsTrigger{empty: true}

func (r *FunctionHttpsTrigger) Empty() bool {
	return r.empty
}

func (r *FunctionHttpsTrigger) String() string {
	return dcl.SprintResource(r)
}

func (r *FunctionHttpsTrigger) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FunctionEventTrigger struct {
	empty         bool    `json:"-"`
	EventType     *string `json:"eventType"`
	Resource      *string `json:"resource"`
	Service       *string `json:"service"`
	FailurePolicy *bool   `json:"failurePolicy"`
}

type jsonFunctionEventTrigger FunctionEventTrigger

func (r *FunctionEventTrigger) UnmarshalJSON(data []byte) error {
	var res jsonFunctionEventTrigger
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFunctionEventTrigger
	} else {

		r.EventType = res.EventType

		r.Resource = res.Resource

		r.Service = res.Service

		r.FailurePolicy = res.FailurePolicy

	}
	return nil
}

// This object is used to assert a desired state where this FunctionEventTrigger is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyFunctionEventTrigger *FunctionEventTrigger = &FunctionEventTrigger{empty: true}

func (r *FunctionEventTrigger) Empty() bool {
	return r.empty
}

func (r *FunctionEventTrigger) String() string {
	return dcl.SprintResource(r)
}

func (r *FunctionEventTrigger) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Function) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "cloud_functions",
		Type:    "Function",
		Version: "cloudfunctions",
	}
}

func (r *Function) ID() (string, error) {
	if err := extractFunctionFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                          dcl.ValueOrEmptyString(nr.Name),
		"description":                   dcl.ValueOrEmptyString(nr.Description),
		"source_archive_url":            dcl.ValueOrEmptyString(nr.SourceArchiveUrl),
		"source_repository":             dcl.ValueOrEmptyString(nr.SourceRepository),
		"https_trigger":                 dcl.ValueOrEmptyString(nr.HttpsTrigger),
		"event_trigger":                 dcl.ValueOrEmptyString(nr.EventTrigger),
		"status":                        dcl.ValueOrEmptyString(nr.Status),
		"entry_point":                   dcl.ValueOrEmptyString(nr.EntryPoint),
		"runtime":                       dcl.ValueOrEmptyString(nr.Runtime),
		"timeout":                       dcl.ValueOrEmptyString(nr.Timeout),
		"available_memory_mb":           dcl.ValueOrEmptyString(nr.AvailableMemoryMb),
		"service_account_email":         dcl.ValueOrEmptyString(nr.ServiceAccountEmail),
		"update_time":                   dcl.ValueOrEmptyString(nr.UpdateTime),
		"version_id":                    dcl.ValueOrEmptyString(nr.VersionId),
		"labels":                        dcl.ValueOrEmptyString(nr.Labels),
		"environment_variables":         dcl.ValueOrEmptyString(nr.EnvironmentVariables),
		"max_instances":                 dcl.ValueOrEmptyString(nr.MaxInstances),
		"vpc_connector":                 dcl.ValueOrEmptyString(nr.VPCConnector),
		"vpc_connector_egress_settings": dcl.ValueOrEmptyString(nr.VPCConnectorEgressSettings),
		"ingress_settings":              dcl.ValueOrEmptyString(nr.IngressSettings),
		"region":                        dcl.ValueOrEmptyString(nr.Region),
		"project":                       dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{region}}/functions/{{name}}", params), nil
}

const FunctionMaxPage = -1

type FunctionList struct {
	Items []*Function

	nextToken string

	pageSize int32

	resource *Function
}

func (l *FunctionList) HasNext() bool {
	return l.nextToken != ""
}

func (l *FunctionList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listFunction(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListFunction(ctx context.Context, project, region string) (*FunctionList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListFunctionWithMaxResults(ctx, project, region, FunctionMaxPage)

}

func (c *Client) ListFunctionWithMaxResults(ctx context.Context, project, region string, pageSize int32) (*FunctionList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &Function{
		Project: &project,
		Region:  &region,
	}
	items, token, err := c.listFunction(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &FunctionList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetFunction(ctx context.Context, r *Function) (*Function, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractFunctionFields(r)

	b, err := c.getFunctionRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalFunction(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Region = r.Region
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeFunctionNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractFunctionFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteFunction(ctx context.Context, r *Function) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Function resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Function...")
	deleteOp := deleteFunctionOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllFunction deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllFunction(ctx context.Context, project, region string, filter func(*Function) bool) error {
	listObj, err := c.ListFunction(ctx, project, region)
	if err != nil {
		return err
	}

	err = c.deleteAllFunction(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllFunction(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyFunction(ctx context.Context, rawDesired *Function, opts ...dcl.ApplyOption) (*Function, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Function
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyFunctionHelper(c, ctx, rawDesired, opts...)
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

func applyFunctionHelper(c *Client, ctx context.Context, rawDesired *Function, opts ...dcl.ApplyOption) (*Function, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyFunction...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractFunctionFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.functionDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToFunctionDiffs(c.Config, fieldDiffs, opts)
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
	var ops []functionApiOperation
	if create {
		ops = append(ops, &createFunctionOperation{})
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
	return applyFunctionDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyFunctionDiff(c *Client, ctx context.Context, desired *Function, rawDesired *Function, ops []functionApiOperation, opts ...dcl.ApplyOption) (*Function, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetFunction(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createFunctionOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapFunction(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeFunctionNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeFunctionNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeFunctionDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractFunctionFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractFunctionFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffFunction(c, newDesired, newState)
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

func (r *Function) GetPolicy(basePath string) (string, string, *bytes.Buffer, error) {
	u := r.getPolicyURL(basePath)
	body := &bytes.Buffer{}
	u, err := dcl.AddQueryParams(u, map[string]string{"options.requestedPolicyVersion": fmt.Sprintf("%d", r.IAMPolicyVersion())})
	if err != nil {
		return "", "", nil, err
	}
	return u, "GET", body, nil
}
