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

type WorkforcePoolProvider struct {
	Name               *string                         `json:"name"`
	DisplayName        *string                         `json:"displayName"`
	Description        *string                         `json:"description"`
	State              *WorkforcePoolProviderStateEnum `json:"state"`
	Disabled           *bool                           `json:"disabled"`
	AttributeMapping   map[string]string               `json:"attributeMapping"`
	AttributeCondition *string                         `json:"attributeCondition"`
	Saml               *WorkforcePoolProviderSaml      `json:"saml"`
	Oidc               *WorkforcePoolProviderOidc      `json:"oidc"`
	Location           *string                         `json:"location"`
	WorkforcePool      *string                         `json:"workforcePool"`
}

func (r *WorkforcePoolProvider) String() string {
	return dcl.SprintResource(r)
}

// The enum WorkforcePoolProviderStateEnum.
type WorkforcePoolProviderStateEnum string

// WorkforcePoolProviderStateEnumRef returns a *WorkforcePoolProviderStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func WorkforcePoolProviderStateEnumRef(s string) *WorkforcePoolProviderStateEnum {
	v := WorkforcePoolProviderStateEnum(s)
	return &v
}

func (v WorkforcePoolProviderStateEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"STATE_UNSPECIFIED", "ACTIVE", "DELETED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "WorkforcePoolProviderStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum WorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum.
type WorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum string

// WorkforcePoolProviderOidcWebSsoConfigResponseTypeEnumRef returns a *WorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func WorkforcePoolProviderOidcWebSsoConfigResponseTypeEnumRef(s string) *WorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum {
	v := WorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum(s)
	return &v
}

func (v WorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"RESPONSE_TYPE_UNSPECIFIED", "CODE", "ID_TOKEN"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "WorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum WorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum.
type WorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum string

// WorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnumRef returns a *WorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum with the value of string s
// If the empty string is provided, nil is returned.
func WorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnumRef(s string) *WorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum {
	v := WorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum(s)
	return &v
}

func (v WorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"ASSERTION_CLAIMS_BEHAVIOR_UNSPECIFIED", "MERGE_USER_INFO_OVER_ID_TOKEN_CLAIMS", "ONLY_ID_TOKEN_CLAIMS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "WorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type WorkforcePoolProviderSaml struct {
	empty          bool    `json:"-"`
	IdpMetadataXml *string `json:"idpMetadataXml"`
}

type jsonWorkforcePoolProviderSaml WorkforcePoolProviderSaml

func (r *WorkforcePoolProviderSaml) UnmarshalJSON(data []byte) error {
	var res jsonWorkforcePoolProviderSaml
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyWorkforcePoolProviderSaml
	} else {

		r.IdpMetadataXml = res.IdpMetadataXml

	}
	return nil
}

// This object is used to assert a desired state where this WorkforcePoolProviderSaml is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyWorkforcePoolProviderSaml *WorkforcePoolProviderSaml = &WorkforcePoolProviderSaml{empty: true}

func (r *WorkforcePoolProviderSaml) Empty() bool {
	return r.empty
}

func (r *WorkforcePoolProviderSaml) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkforcePoolProviderSaml) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkforcePoolProviderOidc struct {
	empty        bool                                   `json:"-"`
	IssuerUri    *string                                `json:"issuerUri"`
	ClientId     *string                                `json:"clientId"`
	JwksJson     *string                                `json:"jwksJson"`
	WebSsoConfig *WorkforcePoolProviderOidcWebSsoConfig `json:"webSsoConfig"`
	ClientSecret *WorkforcePoolProviderOidcClientSecret `json:"clientSecret"`
}

type jsonWorkforcePoolProviderOidc WorkforcePoolProviderOidc

func (r *WorkforcePoolProviderOidc) UnmarshalJSON(data []byte) error {
	var res jsonWorkforcePoolProviderOidc
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyWorkforcePoolProviderOidc
	} else {

		r.IssuerUri = res.IssuerUri

		r.ClientId = res.ClientId

		r.JwksJson = res.JwksJson

		r.WebSsoConfig = res.WebSsoConfig

		r.ClientSecret = res.ClientSecret

	}
	return nil
}

// This object is used to assert a desired state where this WorkforcePoolProviderOidc is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyWorkforcePoolProviderOidc *WorkforcePoolProviderOidc = &WorkforcePoolProviderOidc{empty: true}

func (r *WorkforcePoolProviderOidc) Empty() bool {
	return r.empty
}

func (r *WorkforcePoolProviderOidc) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkforcePoolProviderOidc) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkforcePoolProviderOidcWebSsoConfig struct {
	empty                   bool                                                              `json:"-"`
	ResponseType            *WorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum            `json:"responseType"`
	AssertionClaimsBehavior *WorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum `json:"assertionClaimsBehavior"`
	AdditionalScopes        []string                                                          `json:"additionalScopes"`
}

type jsonWorkforcePoolProviderOidcWebSsoConfig WorkforcePoolProviderOidcWebSsoConfig

func (r *WorkforcePoolProviderOidcWebSsoConfig) UnmarshalJSON(data []byte) error {
	var res jsonWorkforcePoolProviderOidcWebSsoConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyWorkforcePoolProviderOidcWebSsoConfig
	} else {

		r.ResponseType = res.ResponseType

		r.AssertionClaimsBehavior = res.AssertionClaimsBehavior

		r.AdditionalScopes = res.AdditionalScopes

	}
	return nil
}

// This object is used to assert a desired state where this WorkforcePoolProviderOidcWebSsoConfig is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyWorkforcePoolProviderOidcWebSsoConfig *WorkforcePoolProviderOidcWebSsoConfig = &WorkforcePoolProviderOidcWebSsoConfig{empty: true}

func (r *WorkforcePoolProviderOidcWebSsoConfig) Empty() bool {
	return r.empty
}

func (r *WorkforcePoolProviderOidcWebSsoConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkforcePoolProviderOidcWebSsoConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkforcePoolProviderOidcClientSecret struct {
	empty bool                                        `json:"-"`
	Value *WorkforcePoolProviderOidcClientSecretValue `json:"value"`
}

type jsonWorkforcePoolProviderOidcClientSecret WorkforcePoolProviderOidcClientSecret

func (r *WorkforcePoolProviderOidcClientSecret) UnmarshalJSON(data []byte) error {
	var res jsonWorkforcePoolProviderOidcClientSecret
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyWorkforcePoolProviderOidcClientSecret
	} else {

		r.Value = res.Value

	}
	return nil
}

// This object is used to assert a desired state where this WorkforcePoolProviderOidcClientSecret is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyWorkforcePoolProviderOidcClientSecret *WorkforcePoolProviderOidcClientSecret = &WorkforcePoolProviderOidcClientSecret{empty: true}

func (r *WorkforcePoolProviderOidcClientSecret) Empty() bool {
	return r.empty
}

func (r *WorkforcePoolProviderOidcClientSecret) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkforcePoolProviderOidcClientSecret) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkforcePoolProviderOidcClientSecretValue struct {
	empty      bool    `json:"-"`
	PlainText  *string `json:"plainText"`
	Thumbprint *string `json:"thumbprint"`
}

type jsonWorkforcePoolProviderOidcClientSecretValue WorkforcePoolProviderOidcClientSecretValue

func (r *WorkforcePoolProviderOidcClientSecretValue) UnmarshalJSON(data []byte) error {
	var res jsonWorkforcePoolProviderOidcClientSecretValue
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyWorkforcePoolProviderOidcClientSecretValue
	} else {

		r.PlainText = res.PlainText

		r.Thumbprint = res.Thumbprint

	}
	return nil
}

// This object is used to assert a desired state where this WorkforcePoolProviderOidcClientSecretValue is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyWorkforcePoolProviderOidcClientSecretValue *WorkforcePoolProviderOidcClientSecretValue = &WorkforcePoolProviderOidcClientSecretValue{empty: true}

func (r *WorkforcePoolProviderOidcClientSecretValue) Empty() bool {
	return r.empty
}

func (r *WorkforcePoolProviderOidcClientSecretValue) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkforcePoolProviderOidcClientSecretValue) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *WorkforcePoolProvider) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "iam",
		Type:    "WorkforcePoolProvider",
		Version: "beta",
	}
}

func (r *WorkforcePoolProvider) ID() (string, error) {
	if err := extractWorkforcePoolProviderFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                dcl.ValueOrEmptyString(nr.Name),
		"display_name":        dcl.ValueOrEmptyString(nr.DisplayName),
		"description":         dcl.ValueOrEmptyString(nr.Description),
		"state":               dcl.ValueOrEmptyString(nr.State),
		"disabled":            dcl.ValueOrEmptyString(nr.Disabled),
		"attribute_mapping":   dcl.ValueOrEmptyString(nr.AttributeMapping),
		"attribute_condition": dcl.ValueOrEmptyString(nr.AttributeCondition),
		"saml":                dcl.ValueOrEmptyString(nr.Saml),
		"oidc":                dcl.ValueOrEmptyString(nr.Oidc),
		"location":            dcl.ValueOrEmptyString(nr.Location),
		"workforce_pool":      dcl.ValueOrEmptyString(nr.WorkforcePool),
	}
	return dcl.Nprintf("locations/{{location}}/workforcePools/{{workforce_pool}}/providers/{{name}}", params), nil
}

const WorkforcePoolProviderMaxPage = -1

type WorkforcePoolProviderList struct {
	Items []*WorkforcePoolProvider

	nextToken string

	pageSize int32

	resource *WorkforcePoolProvider
}

func (l *WorkforcePoolProviderList) HasNext() bool {
	return l.nextToken != ""
}

func (l *WorkforcePoolProviderList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listWorkforcePoolProvider(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListWorkforcePoolProvider(ctx context.Context, location, workforcePool string) (*WorkforcePoolProviderList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListWorkforcePoolProviderWithMaxResults(ctx, location, workforcePool, WorkforcePoolProviderMaxPage)

}

func (c *Client) ListWorkforcePoolProviderWithMaxResults(ctx context.Context, location, workforcePool string, pageSize int32) (*WorkforcePoolProviderList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &WorkforcePoolProvider{
		Location:      &location,
		WorkforcePool: &workforcePool,
	}
	items, token, err := c.listWorkforcePoolProvider(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &WorkforcePoolProviderList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetWorkforcePoolProvider(ctx context.Context, r *WorkforcePoolProvider) (*WorkforcePoolProvider, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractWorkforcePoolProviderFields(r)

	b, err := c.getWorkforcePoolProviderRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalWorkforcePoolProvider(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Location = r.Location
	result.WorkforcePool = r.WorkforcePool
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeWorkforcePoolProviderNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractWorkforcePoolProviderFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteWorkforcePoolProvider(ctx context.Context, r *WorkforcePoolProvider) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("WorkforcePoolProvider resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting WorkforcePoolProvider...")
	deleteOp := deleteWorkforcePoolProviderOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllWorkforcePoolProvider deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllWorkforcePoolProvider(ctx context.Context, location, workforcePool string, filter func(*WorkforcePoolProvider) bool) error {
	listObj, err := c.ListWorkforcePoolProvider(ctx, location, workforcePool)
	if err != nil {
		return err
	}

	err = c.deleteAllWorkforcePoolProvider(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllWorkforcePoolProvider(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyWorkforcePoolProvider(ctx context.Context, rawDesired *WorkforcePoolProvider, opts ...dcl.ApplyOption) (*WorkforcePoolProvider, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *WorkforcePoolProvider
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyWorkforcePoolProviderHelper(c, ctx, rawDesired, opts...)
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

func applyWorkforcePoolProviderHelper(c *Client, ctx context.Context, rawDesired *WorkforcePoolProvider, opts ...dcl.ApplyOption) (*WorkforcePoolProvider, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyWorkforcePoolProvider...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractWorkforcePoolProviderFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.workforcePoolProviderDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToWorkforcePoolProviderDiffs(c.Config, fieldDiffs, opts)
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
	var ops []workforcePoolProviderApiOperation
	if create {
		ops = append(ops, &createWorkforcePoolProviderOperation{})
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
	return applyWorkforcePoolProviderDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyWorkforcePoolProviderDiff(c *Client, ctx context.Context, desired *WorkforcePoolProvider, rawDesired *WorkforcePoolProvider, ops []workforcePoolProviderApiOperation, opts ...dcl.ApplyOption) (*WorkforcePoolProvider, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetWorkforcePoolProvider(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createWorkforcePoolProviderOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapWorkforcePoolProvider(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeWorkforcePoolProviderNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeWorkforcePoolProviderNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeWorkforcePoolProviderDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractWorkforcePoolProviderFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractWorkforcePoolProviderFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffWorkforcePoolProvider(c, newDesired, newState)
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
