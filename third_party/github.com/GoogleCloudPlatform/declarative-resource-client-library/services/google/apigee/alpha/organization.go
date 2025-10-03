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

type Organization struct {
	Name                             *string                           `json:"name"`
	DisplayName                      *string                           `json:"displayName"`
	Description                      *string                           `json:"description"`
	CreatedAt                        *int64                            `json:"createdAt"`
	LastModifiedAt                   *int64                            `json:"lastModifiedAt"`
	ExpiresAt                        *int64                            `json:"expiresAt"`
	Environments                     []string                          `json:"environments"`
	Properties                       map[string]string                 `json:"properties"`
	AnalyticsRegion                  *string                           `json:"analyticsRegion"`
	AuthorizedNetwork                *string                           `json:"authorizedNetwork"`
	RuntimeType                      *OrganizationRuntimeTypeEnum      `json:"runtimeType"`
	SubscriptionType                 *OrganizationSubscriptionTypeEnum `json:"subscriptionType"`
	BillingType                      *OrganizationBillingTypeEnum      `json:"billingType"`
	AddonsConfig                     *OrganizationAddonsConfig         `json:"addonsConfig"`
	CaCertificate                    *string                           `json:"caCertificate"`
	RuntimeDatabaseEncryptionKeyName *string                           `json:"runtimeDatabaseEncryptionKeyName"`
	ProjectId                        *string                           `json:"projectId"`
	State                            *OrganizationStateEnum            `json:"state"`
	Project                          *string                           `json:"project"`
}

func (r *Organization) String() string {
	return dcl.SprintResource(r)
}

// The enum OrganizationRuntimeTypeEnum.
type OrganizationRuntimeTypeEnum string

// OrganizationRuntimeTypeEnumRef returns a *OrganizationRuntimeTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func OrganizationRuntimeTypeEnumRef(s string) *OrganizationRuntimeTypeEnum {
	v := OrganizationRuntimeTypeEnum(s)
	return &v
}

func (v OrganizationRuntimeTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"RUNTIME_TYPE_UNSPECIFIED", "CLOUD", "HYBRID"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "OrganizationRuntimeTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum OrganizationSubscriptionTypeEnum.
type OrganizationSubscriptionTypeEnum string

// OrganizationSubscriptionTypeEnumRef returns a *OrganizationSubscriptionTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func OrganizationSubscriptionTypeEnumRef(s string) *OrganizationSubscriptionTypeEnum {
	v := OrganizationSubscriptionTypeEnum(s)
	return &v
}

func (v OrganizationSubscriptionTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"SUBSCRIPTION_TYPE_UNSPECIFIED", "PAID", "TRIAL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "OrganizationSubscriptionTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum OrganizationBillingTypeEnum.
type OrganizationBillingTypeEnum string

// OrganizationBillingTypeEnumRef returns a *OrganizationBillingTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func OrganizationBillingTypeEnumRef(s string) *OrganizationBillingTypeEnum {
	v := OrganizationBillingTypeEnum(s)
	return &v
}

func (v OrganizationBillingTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"BILLING_TYPE_UNSPECIFIED", "SUBSCRIPTION", "EVALUATION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "OrganizationBillingTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum OrganizationStateEnum.
type OrganizationStateEnum string

// OrganizationStateEnumRef returns a *OrganizationStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func OrganizationStateEnumRef(s string) *OrganizationStateEnum {
	v := OrganizationStateEnum(s)
	return &v
}

func (v OrganizationStateEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"SNAPSHOT_STATE_UNSPECIFIED", "MISSING", "OK_DOCSTORE", "OK_SUBMITTED", "OK_EXTERNAL", "DELETED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "OrganizationStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type OrganizationAddonsConfig struct {
	empty                bool                                          `json:"-"`
	AdvancedApiOpsConfig *OrganizationAddonsConfigAdvancedApiOpsConfig `json:"advancedApiOpsConfig"`
	MonetizationConfig   *OrganizationAddonsConfigMonetizationConfig   `json:"monetizationConfig"`
}

type jsonOrganizationAddonsConfig OrganizationAddonsConfig

func (r *OrganizationAddonsConfig) UnmarshalJSON(data []byte) error {
	var res jsonOrganizationAddonsConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOrganizationAddonsConfig
	} else {

		r.AdvancedApiOpsConfig = res.AdvancedApiOpsConfig

		r.MonetizationConfig = res.MonetizationConfig

	}
	return nil
}

// This object is used to assert a desired state where this OrganizationAddonsConfig is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyOrganizationAddonsConfig *OrganizationAddonsConfig = &OrganizationAddonsConfig{empty: true}

func (r *OrganizationAddonsConfig) Empty() bool {
	return r.empty
}

func (r *OrganizationAddonsConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *OrganizationAddonsConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OrganizationAddonsConfigAdvancedApiOpsConfig struct {
	empty   bool  `json:"-"`
	Enabled *bool `json:"enabled"`
}

type jsonOrganizationAddonsConfigAdvancedApiOpsConfig OrganizationAddonsConfigAdvancedApiOpsConfig

func (r *OrganizationAddonsConfigAdvancedApiOpsConfig) UnmarshalJSON(data []byte) error {
	var res jsonOrganizationAddonsConfigAdvancedApiOpsConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOrganizationAddonsConfigAdvancedApiOpsConfig
	} else {

		r.Enabled = res.Enabled

	}
	return nil
}

// This object is used to assert a desired state where this OrganizationAddonsConfigAdvancedApiOpsConfig is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyOrganizationAddonsConfigAdvancedApiOpsConfig *OrganizationAddonsConfigAdvancedApiOpsConfig = &OrganizationAddonsConfigAdvancedApiOpsConfig{empty: true}

func (r *OrganizationAddonsConfigAdvancedApiOpsConfig) Empty() bool {
	return r.empty
}

func (r *OrganizationAddonsConfigAdvancedApiOpsConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *OrganizationAddonsConfigAdvancedApiOpsConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type OrganizationAddonsConfigMonetizationConfig struct {
	empty   bool  `json:"-"`
	Enabled *bool `json:"enabled"`
}

type jsonOrganizationAddonsConfigMonetizationConfig OrganizationAddonsConfigMonetizationConfig

func (r *OrganizationAddonsConfigMonetizationConfig) UnmarshalJSON(data []byte) error {
	var res jsonOrganizationAddonsConfigMonetizationConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyOrganizationAddonsConfigMonetizationConfig
	} else {

		r.Enabled = res.Enabled

	}
	return nil
}

// This object is used to assert a desired state where this OrganizationAddonsConfigMonetizationConfig is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyOrganizationAddonsConfigMonetizationConfig *OrganizationAddonsConfigMonetizationConfig = &OrganizationAddonsConfigMonetizationConfig{empty: true}

func (r *OrganizationAddonsConfigMonetizationConfig) Empty() bool {
	return r.empty
}

func (r *OrganizationAddonsConfigMonetizationConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *OrganizationAddonsConfigMonetizationConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Organization) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "apigee",
		Type:    "Organization",
		Version: "alpha",
	}
}

func (r *Organization) ID() (string, error) {
	if err := extractOrganizationFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                                 dcl.ValueOrEmptyString(nr.Name),
		"display_name":                         dcl.ValueOrEmptyString(nr.DisplayName),
		"description":                          dcl.ValueOrEmptyString(nr.Description),
		"created_at":                           dcl.ValueOrEmptyString(nr.CreatedAt),
		"last_modified_at":                     dcl.ValueOrEmptyString(nr.LastModifiedAt),
		"expires_at":                           dcl.ValueOrEmptyString(nr.ExpiresAt),
		"environments":                         dcl.ValueOrEmptyString(nr.Environments),
		"properties":                           dcl.ValueOrEmptyString(nr.Properties),
		"analytics_region":                     dcl.ValueOrEmptyString(nr.AnalyticsRegion),
		"authorized_network":                   dcl.ValueOrEmptyString(nr.AuthorizedNetwork),
		"runtime_type":                         dcl.ValueOrEmptyString(nr.RuntimeType),
		"subscription_type":                    dcl.ValueOrEmptyString(nr.SubscriptionType),
		"billing_type":                         dcl.ValueOrEmptyString(nr.BillingType),
		"addons_config":                        dcl.ValueOrEmptyString(nr.AddonsConfig),
		"ca_certificate":                       dcl.ValueOrEmptyString(nr.CaCertificate),
		"runtime_database_encryption_key_name": dcl.ValueOrEmptyString(nr.RuntimeDatabaseEncryptionKeyName),
		"project_id":                           dcl.ValueOrEmptyString(nr.ProjectId),
		"state":                                dcl.ValueOrEmptyString(nr.State),
		"project":                              dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.Nprintf("organizations/{{name}}", params), nil
}

const OrganizationMaxPage = -1

type OrganizationList struct {
	Items []*Organization

	nextToken string

	pageSize int32

	resource *Organization
}

func (c *Client) GetOrganization(ctx context.Context, r *Organization) (*Organization, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractOrganizationFields(r)

	b, err := c.getOrganizationRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFoundOrCode(err, 403) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalOrganization(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeOrganizationNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractOrganizationFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteOrganization(ctx context.Context, r *Organization) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(4800*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Organization resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Organization...")
	deleteOp := deleteOrganizationOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllOrganization deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllOrganization(ctx context.Context, filter func(*Organization) bool) error {
	listObj, err := c.ListOrganization(ctx)
	if err != nil {
		return err
	}

	err = c.deleteAllOrganization(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllOrganization(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyOrganization(ctx context.Context, rawDesired *Organization, opts ...dcl.ApplyOption) (*Organization, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(4800*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Organization
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyOrganizationHelper(c, ctx, rawDesired, opts...)
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

func applyOrganizationHelper(c *Client, ctx context.Context, rawDesired *Organization, opts ...dcl.ApplyOption) (*Organization, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyOrganization...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractOrganizationFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.organizationDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToOrganizationDiffs(c.Config, fieldDiffs, opts)
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
	var ops []organizationApiOperation
	if create {
		ops = append(ops, &createOrganizationOperation{})
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	ops, err = createProperties(ops)
	if err != nil {
		return nil, err
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
	return applyOrganizationDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyOrganizationDiff(c *Client, ctx context.Context, desired *Organization, rawDesired *Organization, ops []organizationApiOperation, opts ...dcl.ApplyOption) (*Organization, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetOrganization(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createOrganizationOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapOrganization(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeOrganizationNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeOrganizationNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeOrganizationDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractOrganizationFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractOrganizationFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffOrganization(c, newDesired, newState)
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
