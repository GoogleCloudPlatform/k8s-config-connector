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

type Tenant struct {
	Name                  *string           `json:"name"`
	DisplayName           *string           `json:"displayName"`
	AllowPasswordSignup   *bool             `json:"allowPasswordSignup"`
	EnableEmailLinkSignin *bool             `json:"enableEmailLinkSignin"`
	DisableAuth           *bool             `json:"disableAuth"`
	EnableAnonymousUser   *bool             `json:"enableAnonymousUser"`
	MfaConfig             *TenantMfaConfig  `json:"mfaConfig"`
	TestPhoneNumbers      map[string]string `json:"testPhoneNumbers"`
	Project               *string           `json:"project"`
}

func (r *Tenant) String() string {
	return dcl.SprintResource(r)
}

// The enum TenantMfaConfigStateEnum.
type TenantMfaConfigStateEnum string

// TenantMfaConfigStateEnumRef returns a *TenantMfaConfigStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func TenantMfaConfigStateEnumRef(s string) *TenantMfaConfigStateEnum {
	v := TenantMfaConfigStateEnum(s)
	return &v
}

func (v TenantMfaConfigStateEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"STATE_UNSPECIFIED", "DISABLED", "ENABLED", "MANDATORY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "TenantMfaConfigStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum TenantMfaConfigEnabledProvidersEnum.
type TenantMfaConfigEnabledProvidersEnum string

// TenantMfaConfigEnabledProvidersEnumRef returns a *TenantMfaConfigEnabledProvidersEnum with the value of string s
// If the empty string is provided, nil is returned.
func TenantMfaConfigEnabledProvidersEnumRef(s string) *TenantMfaConfigEnabledProvidersEnum {
	v := TenantMfaConfigEnabledProvidersEnum(s)
	return &v
}

func (v TenantMfaConfigEnabledProvidersEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"PROVIDER_UNSPECIFIED", "PHONE_SMS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "TenantMfaConfigEnabledProvidersEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type TenantMfaConfig struct {
	empty            bool                                  `json:"-"`
	State            *TenantMfaConfigStateEnum             `json:"state"`
	EnabledProviders []TenantMfaConfigEnabledProvidersEnum `json:"enabledProviders"`
}

type jsonTenantMfaConfig TenantMfaConfig

func (r *TenantMfaConfig) UnmarshalJSON(data []byte) error {
	var res jsonTenantMfaConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTenantMfaConfig
	} else {

		r.State = res.State

		r.EnabledProviders = res.EnabledProviders

	}
	return nil
}

// This object is used to assert a desired state where this TenantMfaConfig is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyTenantMfaConfig *TenantMfaConfig = &TenantMfaConfig{empty: true}

func (r *TenantMfaConfig) Empty() bool {
	return r.empty
}

func (r *TenantMfaConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *TenantMfaConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Tenant) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "identity_toolkit",
		Type:    "Tenant",
		Version: "beta",
	}
}

func (r *Tenant) ID() (string, error) {
	if err := extractTenantFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                     dcl.ValueOrEmptyString(nr.Name),
		"display_name":             dcl.ValueOrEmptyString(nr.DisplayName),
		"allow_password_signup":    dcl.ValueOrEmptyString(nr.AllowPasswordSignup),
		"enable_email_link_signin": dcl.ValueOrEmptyString(nr.EnableEmailLinkSignin),
		"disable_auth":             dcl.ValueOrEmptyString(nr.DisableAuth),
		"enable_anonymous_user":    dcl.ValueOrEmptyString(nr.EnableAnonymousUser),
		"mfa_config":               dcl.ValueOrEmptyString(nr.MfaConfig),
		"test_phone_numbers":       dcl.ValueOrEmptyString(nr.TestPhoneNumbers),
		"project":                  dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.Nprintf("projects/{{project}}/tenants/{{name}}", params), nil
}

const TenantMaxPage = -1

type TenantList struct {
	Items []*Tenant

	nextToken string

	pageSize int32

	resource *Tenant
}

func (l *TenantList) HasNext() bool {
	return l.nextToken != ""
}

func (l *TenantList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listTenant(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListTenant(ctx context.Context, project string) (*TenantList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListTenantWithMaxResults(ctx, project, TenantMaxPage)

}

func (c *Client) ListTenantWithMaxResults(ctx context.Context, project string, pageSize int32) (*TenantList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &Tenant{
		Project: &project,
	}
	items, token, err := c.listTenant(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &TenantList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetTenant(ctx context.Context, r *Tenant) (*Tenant, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractTenantFields(r)

	b, err := c.getTenantRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalTenant(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeTenantNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractTenantFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteTenant(ctx context.Context, r *Tenant) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Tenant resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Tenant...")
	deleteOp := deleteTenantOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllTenant deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllTenant(ctx context.Context, project string, filter func(*Tenant) bool) error {
	listObj, err := c.ListTenant(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllTenant(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllTenant(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyTenant(ctx context.Context, rawDesired *Tenant, opts ...dcl.ApplyOption) (*Tenant, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Tenant
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyTenantHelper(c, ctx, rawDesired, opts...)
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

func applyTenantHelper(c *Client, ctx context.Context, rawDesired *Tenant, opts ...dcl.ApplyOption) (*Tenant, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyTenant...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractTenantFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.tenantDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToTenantDiffs(c.Config, fieldDiffs, opts)
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
	var ops []tenantApiOperation
	if create {
		ops = append(ops, &createTenantOperation{})
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
	return applyTenantDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyTenantDiff(c *Client, ctx context.Context, desired *Tenant, rawDesired *Tenant, ops []tenantApiOperation, opts ...dcl.ApplyOption) (*Tenant, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetTenant(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createTenantOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapTenant(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeTenantNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeTenantNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeTenantDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractTenantFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractTenantFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffTenant(c, newDesired, newState)
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
