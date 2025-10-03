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
package apigee

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Organization) validate() error {

	if err := dcl.Required(r, "analyticsRegion"); err != nil {
		return err
	}
	if err := dcl.Required(r, "runtimeType"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.AddonsConfig) {
		if err := r.AddonsConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *OrganizationAddonsConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.AdvancedApiOpsConfig) {
		if err := r.AdvancedApiOpsConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MonetizationConfig) {
		if err := r.MonetizationConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *OrganizationAddonsConfigAdvancedApiOpsConfig) validate() error {
	return nil
}
func (r *OrganizationAddonsConfigMonetizationConfig) validate() error {
	return nil
}
func (r *Organization) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://apigee.googleapis.com/v1/", params)
}

func (r *Organization) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name": dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("organizations/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Organization) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{}
	return dcl.URL("organizations", nr.basePath(), userBasePath, params), nil

}

func (r *Organization) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.URL("organizations?parent={{project}}", nr.basePath(), userBasePath, params), nil

}

func (r *Organization) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name": dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("organizations/{{name}}", nr.basePath(), userBasePath, params), nil
}

// organizationApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type organizationApiOperation interface {
	do(context.Context, *Organization, *Client) error
}

// newUpdateOrganizationSetAddonsRequest creates a request for an
// Organization resource's SetAddons update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateOrganizationSetAddonsRequest(ctx context.Context, f *Organization, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v, err := expandOrganizationAddonsConfig(c, f.AddonsConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding AddonsConfig into addonsConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["addonsConfig"] = v
	}
	return req, nil
}

// marshalUpdateOrganizationSetAddonsRequest converts the update into
// the final JSON request body.
func marshalUpdateOrganizationSetAddonsRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateOrganizationSetAddonsOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateOrganizationSetAddonsOperation) do(ctx context.Context, r *Organization, c *Client) error {
	_, err := c.GetOrganization(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "SetAddons")
	if err != nil {
		return err
	}

	req, err := newUpdateOrganizationSetAddonsRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateOrganizationSetAddonsRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET")

	if err != nil {
		return err
	}

	return nil
}

// newUpdateOrganizationUpdateOrganizationRequest creates a request for an
// Organization resource's UpdateOrganization update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateOrganizationUpdateOrganizationRequest(ctx context.Context, f *Organization, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v, err := dcl.ListOfKeyValuesFromMapInStruct(f.Properties, "property", "name", "value"); err != nil {
		return nil, fmt.Errorf("error expanding Properties into properties: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["properties"] = v
	}
	if v := f.AuthorizedNetwork; !dcl.IsEmptyValueIndirect(v) {
		req["authorizedNetwork"] = v
	}
	if v := f.RuntimeDatabaseEncryptionKeyName; !dcl.IsEmptyValueIndirect(v) {
		req["runtimeDatabaseEncryptionKeyName"] = v
	}
	return req, nil
}

// marshalUpdateOrganizationUpdateOrganizationRequest converts the update into
// the final JSON request body.
func marshalUpdateOrganizationUpdateOrganizationRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateOrganizationUpdateOrganizationOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (c *Client) deleteAllOrganization(ctx context.Context, f func(*Organization) bool, resources []*Organization) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteOrganization(ctx, res)
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("%v", strings.Join(errors, "\n"))
	} else {
		return nil
	}
}

type deleteOrganizationOperation struct{}

func (op *deleteOrganizationOperation) do(ctx context.Context, r *Organization, c *Client) error {
	r, err := c.GetOrganization(ctx, r)
	if err != nil {
		if dcl.IsNotFoundOrCode(err, 403) {
			c.Config.Logger.InfoWithContextf(ctx, "Organization not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetOrganization checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		return err
	}

	// We saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// This is the reason we are adding retry to handle that case.
	retriesRemaining := 10
	dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		_, err := c.GetOrganization(ctx, r)
		if dcl.IsNotFoundOrCode(err, 403) {
			return nil, nil
		}
		if retriesRemaining > 0 {
			retriesRemaining--
			return &dcl.RetryDetails{}, dcl.OperationNotDone{}
		}
		return nil, dcl.NotDeletedError{ExistingResource: r}
	}, c.Config.RetryProvider)
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createOrganizationOperation struct {
	response map[string]interface{}
}

func (op *createOrganizationOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createOrganizationOperation) do(ctx context.Context, r *Organization, c *Client) error {
	c.Config.Logger.InfoWithContextf(ctx, "Attempting to create %v", r)
	u, err := r.createURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	if r.Name != nil {
		// Allowing creation to continue with Name set could result in a Organization with the wrong Name.
		return fmt.Errorf("server-generated parameter Name was specified by user as %v, should be unspecified", dcl.ValueOrEmptyString(r.Name))
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}
	// wait for object to be created.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	// Include Name in URL substitution for initial GET request.
	m := op.response
	r.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))

	if _, err := c.GetOrganization(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getOrganizationRaw(ctx context.Context, r *Organization) ([]byte, error) {

	u, err := r.getURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	b, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) organizationDiffsForRawDesired(ctx context.Context, rawDesired *Organization, opts ...dcl.ApplyOption) (initial, desired *Organization, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Organization
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Organization); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Organization, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	if fetchState.Name == nil {
		// We cannot perform a get because of lack of information. We have to assume
		// that this is being created for the first time.
		desired, err := canonicalizeOrganizationDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}
	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetOrganization(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFoundOrCode(err, 403) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Organization resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Organization resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Organization resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeOrganizationDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Organization: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Organization: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractOrganizationFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeOrganizationInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Organization: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeOrganizationDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Organization: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffOrganization(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeOrganizationInitialState(rawInitial, rawDesired *Organization) (*Organization, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeOrganizationDesiredState(rawDesired, rawInitial *Organization, opts ...dcl.ApplyOption) (*Organization, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.AddonsConfig = canonicalizeOrganizationAddonsConfig(rawDesired.AddonsConfig, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Organization{}
	if dcl.IsZeroValue(rawDesired.Name) || (dcl.IsEmptyValueIndirect(rawDesired.Name) && dcl.IsEmptyValueIndirect(rawInitial.Name)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		canonicalDesired.DisplayName = rawInitial.DisplayName
	} else {
		canonicalDesired.DisplayName = rawDesired.DisplayName
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.IsZeroValue(rawDesired.Properties) || (dcl.IsEmptyValueIndirect(rawDesired.Properties) && dcl.IsEmptyValueIndirect(rawInitial.Properties)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Properties = rawInitial.Properties
	} else {
		canonicalDesired.Properties = rawDesired.Properties
	}
	if dcl.StringCanonicalize(rawDesired.AnalyticsRegion, rawInitial.AnalyticsRegion) {
		canonicalDesired.AnalyticsRegion = rawInitial.AnalyticsRegion
	} else {
		canonicalDesired.AnalyticsRegion = rawDesired.AnalyticsRegion
	}
	if dcl.IsZeroValue(rawDesired.AuthorizedNetwork) || (dcl.IsEmptyValueIndirect(rawDesired.AuthorizedNetwork) && dcl.IsEmptyValueIndirect(rawInitial.AuthorizedNetwork)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.AuthorizedNetwork = rawInitial.AuthorizedNetwork
	} else {
		canonicalDesired.AuthorizedNetwork = rawDesired.AuthorizedNetwork
	}
	if dcl.IsZeroValue(rawDesired.RuntimeType) || (dcl.IsEmptyValueIndirect(rawDesired.RuntimeType) && dcl.IsEmptyValueIndirect(rawInitial.RuntimeType)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.RuntimeType = rawInitial.RuntimeType
	} else {
		canonicalDesired.RuntimeType = rawDesired.RuntimeType
	}
	canonicalDesired.AddonsConfig = canonicalizeOrganizationAddonsConfig(rawDesired.AddonsConfig, rawInitial.AddonsConfig, opts...)
	if dcl.IsZeroValue(rawDesired.RuntimeDatabaseEncryptionKeyName) || (dcl.IsEmptyValueIndirect(rawDesired.RuntimeDatabaseEncryptionKeyName) && dcl.IsEmptyValueIndirect(rawInitial.RuntimeDatabaseEncryptionKeyName)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.RuntimeDatabaseEncryptionKeyName = rawInitial.RuntimeDatabaseEncryptionKeyName
	} else {
		canonicalDesired.RuntimeDatabaseEncryptionKeyName = rawDesired.RuntimeDatabaseEncryptionKeyName
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	return canonicalDesired, nil
}

func canonicalizeOrganizationNewState(c *Client, rawNew, rawDesired *Organization) (*Organization, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreatedAt) && dcl.IsEmptyValueIndirect(rawDesired.CreatedAt) {
		rawNew.CreatedAt = rawDesired.CreatedAt
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.LastModifiedAt) && dcl.IsEmptyValueIndirect(rawDesired.LastModifiedAt) {
		rawNew.LastModifiedAt = rawDesired.LastModifiedAt
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ExpiresAt) && dcl.IsEmptyValueIndirect(rawDesired.ExpiresAt) {
		rawNew.ExpiresAt = rawDesired.ExpiresAt
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Environments) && dcl.IsEmptyValueIndirect(rawDesired.Environments) {
		rawNew.Environments = rawDesired.Environments
	} else {
		if dcl.StringArrayCanonicalize(rawDesired.Environments, rawNew.Environments) {
			rawNew.Environments = rawDesired.Environments
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Properties) && dcl.IsEmptyValueIndirect(rawDesired.Properties) {
		rawNew.Properties = rawDesired.Properties
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.AnalyticsRegion) && dcl.IsEmptyValueIndirect(rawDesired.AnalyticsRegion) {
		rawNew.AnalyticsRegion = rawDesired.AnalyticsRegion
	} else {
		if dcl.StringCanonicalize(rawDesired.AnalyticsRegion, rawNew.AnalyticsRegion) {
			rawNew.AnalyticsRegion = rawDesired.AnalyticsRegion
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.AuthorizedNetwork) && dcl.IsEmptyValueIndirect(rawDesired.AuthorizedNetwork) {
		rawNew.AuthorizedNetwork = rawDesired.AuthorizedNetwork
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.RuntimeType) && dcl.IsEmptyValueIndirect(rawDesired.RuntimeType) {
		rawNew.RuntimeType = rawDesired.RuntimeType
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SubscriptionType) && dcl.IsEmptyValueIndirect(rawDesired.SubscriptionType) {
		rawNew.SubscriptionType = rawDesired.SubscriptionType
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.BillingType) && dcl.IsEmptyValueIndirect(rawDesired.BillingType) {
		rawNew.BillingType = rawDesired.BillingType
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.AddonsConfig) && dcl.IsEmptyValueIndirect(rawDesired.AddonsConfig) {
		rawNew.AddonsConfig = rawDesired.AddonsConfig
	} else {
		rawNew.AddonsConfig = canonicalizeNewOrganizationAddonsConfig(c, rawDesired.AddonsConfig, rawNew.AddonsConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.CaCertificate) && dcl.IsEmptyValueIndirect(rawDesired.CaCertificate) {
		rawNew.CaCertificate = rawDesired.CaCertificate
	} else {
		if dcl.StringCanonicalize(rawDesired.CaCertificate, rawNew.CaCertificate) {
			rawNew.CaCertificate = rawDesired.CaCertificate
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.RuntimeDatabaseEncryptionKeyName) && dcl.IsEmptyValueIndirect(rawDesired.RuntimeDatabaseEncryptionKeyName) {
		rawNew.RuntimeDatabaseEncryptionKeyName = rawDesired.RuntimeDatabaseEncryptionKeyName
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ProjectId) && dcl.IsEmptyValueIndirect(rawDesired.ProjectId) {
		rawNew.ProjectId = rawDesired.ProjectId
	} else {
		if dcl.StringCanonicalize(rawDesired.ProjectId, rawNew.ProjectId) {
			rawNew.ProjectId = rawDesired.ProjectId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeOrganizationAddonsConfig(des, initial *OrganizationAddonsConfig, opts ...dcl.ApplyOption) *OrganizationAddonsConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &OrganizationAddonsConfig{}

	cDes.AdvancedApiOpsConfig = canonicalizeOrganizationAddonsConfigAdvancedApiOpsConfig(des.AdvancedApiOpsConfig, initial.AdvancedApiOpsConfig, opts...)
	cDes.MonetizationConfig = canonicalizeOrganizationAddonsConfigMonetizationConfig(des.MonetizationConfig, initial.MonetizationConfig, opts...)

	return cDes
}

func canonicalizeOrganizationAddonsConfigSlice(des, initial []OrganizationAddonsConfig, opts ...dcl.ApplyOption) []OrganizationAddonsConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]OrganizationAddonsConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeOrganizationAddonsConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]OrganizationAddonsConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeOrganizationAddonsConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewOrganizationAddonsConfig(c *Client, des, nw *OrganizationAddonsConfig) *OrganizationAddonsConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for OrganizationAddonsConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.AdvancedApiOpsConfig = canonicalizeNewOrganizationAddonsConfigAdvancedApiOpsConfig(c, des.AdvancedApiOpsConfig, nw.AdvancedApiOpsConfig)
	nw.MonetizationConfig = canonicalizeNewOrganizationAddonsConfigMonetizationConfig(c, des.MonetizationConfig, nw.MonetizationConfig)

	return nw
}

func canonicalizeNewOrganizationAddonsConfigSet(c *Client, des, nw []OrganizationAddonsConfig) []OrganizationAddonsConfig {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []OrganizationAddonsConfig
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareOrganizationAddonsConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewOrganizationAddonsConfig(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewOrganizationAddonsConfigSlice(c *Client, des, nw []OrganizationAddonsConfig) []OrganizationAddonsConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OrganizationAddonsConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOrganizationAddonsConfig(c, &d, &n))
	}

	return items
}

func canonicalizeOrganizationAddonsConfigAdvancedApiOpsConfig(des, initial *OrganizationAddonsConfigAdvancedApiOpsConfig, opts ...dcl.ApplyOption) *OrganizationAddonsConfigAdvancedApiOpsConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &OrganizationAddonsConfigAdvancedApiOpsConfig{}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		cDes.Enabled = initial.Enabled
	} else {
		cDes.Enabled = des.Enabled
	}

	return cDes
}

func canonicalizeOrganizationAddonsConfigAdvancedApiOpsConfigSlice(des, initial []OrganizationAddonsConfigAdvancedApiOpsConfig, opts ...dcl.ApplyOption) []OrganizationAddonsConfigAdvancedApiOpsConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]OrganizationAddonsConfigAdvancedApiOpsConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeOrganizationAddonsConfigAdvancedApiOpsConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]OrganizationAddonsConfigAdvancedApiOpsConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeOrganizationAddonsConfigAdvancedApiOpsConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewOrganizationAddonsConfigAdvancedApiOpsConfig(c *Client, des, nw *OrganizationAddonsConfigAdvancedApiOpsConfig) *OrganizationAddonsConfigAdvancedApiOpsConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for OrganizationAddonsConfigAdvancedApiOpsConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}

	return nw
}

func canonicalizeNewOrganizationAddonsConfigAdvancedApiOpsConfigSet(c *Client, des, nw []OrganizationAddonsConfigAdvancedApiOpsConfig) []OrganizationAddonsConfigAdvancedApiOpsConfig {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []OrganizationAddonsConfigAdvancedApiOpsConfig
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareOrganizationAddonsConfigAdvancedApiOpsConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewOrganizationAddonsConfigAdvancedApiOpsConfig(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewOrganizationAddonsConfigAdvancedApiOpsConfigSlice(c *Client, des, nw []OrganizationAddonsConfigAdvancedApiOpsConfig) []OrganizationAddonsConfigAdvancedApiOpsConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OrganizationAddonsConfigAdvancedApiOpsConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOrganizationAddonsConfigAdvancedApiOpsConfig(c, &d, &n))
	}

	return items
}

func canonicalizeOrganizationAddonsConfigMonetizationConfig(des, initial *OrganizationAddonsConfigMonetizationConfig, opts ...dcl.ApplyOption) *OrganizationAddonsConfigMonetizationConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &OrganizationAddonsConfigMonetizationConfig{}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		cDes.Enabled = initial.Enabled
	} else {
		cDes.Enabled = des.Enabled
	}

	return cDes
}

func canonicalizeOrganizationAddonsConfigMonetizationConfigSlice(des, initial []OrganizationAddonsConfigMonetizationConfig, opts ...dcl.ApplyOption) []OrganizationAddonsConfigMonetizationConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]OrganizationAddonsConfigMonetizationConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeOrganizationAddonsConfigMonetizationConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]OrganizationAddonsConfigMonetizationConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeOrganizationAddonsConfigMonetizationConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewOrganizationAddonsConfigMonetizationConfig(c *Client, des, nw *OrganizationAddonsConfigMonetizationConfig) *OrganizationAddonsConfigMonetizationConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for OrganizationAddonsConfigMonetizationConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}

	return nw
}

func canonicalizeNewOrganizationAddonsConfigMonetizationConfigSet(c *Client, des, nw []OrganizationAddonsConfigMonetizationConfig) []OrganizationAddonsConfigMonetizationConfig {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []OrganizationAddonsConfigMonetizationConfig
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareOrganizationAddonsConfigMonetizationConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewOrganizationAddonsConfigMonetizationConfig(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewOrganizationAddonsConfigMonetizationConfigSlice(c *Client, des, nw []OrganizationAddonsConfigMonetizationConfig) []OrganizationAddonsConfigMonetizationConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []OrganizationAddonsConfigMonetizationConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewOrganizationAddonsConfigMonetizationConfig(c, &d, &n))
	}

	return items
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffOrganization(c *Client, desired, actual *Organization, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateOrganizationUpdateOrganizationOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateOrganizationUpdateOrganizationOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreatedAt, actual.CreatedAt, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreatedAt")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LastModifiedAt, actual.LastModifiedAt, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LastModifiedAt")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExpiresAt, actual.ExpiresAt, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExpiresAt")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Environments, actual.Environments, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Environments")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Properties, actual.Properties, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateOrganizationUpdateOrganizationOperation")}, fn.AddNest("Properties")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AnalyticsRegion, actual.AnalyticsRegion, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AnalyticsRegion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AuthorizedNetwork, actual.AuthorizedNetwork, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateOrganizationUpdateOrganizationOperation")}, fn.AddNest("AuthorizedNetwork")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RuntimeType, actual.RuntimeType, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RuntimeType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SubscriptionType, actual.SubscriptionType, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SubscriptionType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BillingType, actual.BillingType, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BillingType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AddonsConfig, actual.AddonsConfig, dcl.DiffInfo{ObjectFunction: compareOrganizationAddonsConfigNewStyle, EmptyObject: EmptyOrganizationAddonsConfig, OperationSelector: dcl.TriggersOperation("updateOrganizationSetAddonsOperation")}, fn.AddNest("AddonsConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CaCertificate, actual.CaCertificate, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CaCertificate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RuntimeDatabaseEncryptionKeyName, actual.RuntimeDatabaseEncryptionKeyName, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateOrganizationUpdateOrganizationOperation")}, fn.AddNest("RuntimeDatabaseEncryptionKeyName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ProjectId, actual.ProjectId, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ProjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if len(newDiffs) > 0 {
		c.Config.Logger.Infof("Diff function found diffs: %v", newDiffs)
	}
	return newDiffs, nil
}
func compareOrganizationAddonsConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OrganizationAddonsConfig)
	if !ok {
		desiredNotPointer, ok := d.(OrganizationAddonsConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OrganizationAddonsConfig or *OrganizationAddonsConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OrganizationAddonsConfig)
	if !ok {
		actualNotPointer, ok := a.(OrganizationAddonsConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OrganizationAddonsConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AdvancedApiOpsConfig, actual.AdvancedApiOpsConfig, dcl.DiffInfo{ObjectFunction: compareOrganizationAddonsConfigAdvancedApiOpsConfigNewStyle, EmptyObject: EmptyOrganizationAddonsConfigAdvancedApiOpsConfig, OperationSelector: dcl.TriggersOperation("updateOrganizationSetAddonsOperation")}, fn.AddNest("AdvancedApiOpsConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MonetizationConfig, actual.MonetizationConfig, dcl.DiffInfo{ObjectFunction: compareOrganizationAddonsConfigMonetizationConfigNewStyle, EmptyObject: EmptyOrganizationAddonsConfigMonetizationConfig, OperationSelector: dcl.TriggersOperation("updateOrganizationSetAddonsOperation")}, fn.AddNest("MonetizationConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOrganizationAddonsConfigAdvancedApiOpsConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OrganizationAddonsConfigAdvancedApiOpsConfig)
	if !ok {
		desiredNotPointer, ok := d.(OrganizationAddonsConfigAdvancedApiOpsConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OrganizationAddonsConfigAdvancedApiOpsConfig or *OrganizationAddonsConfigAdvancedApiOpsConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OrganizationAddonsConfigAdvancedApiOpsConfig)
	if !ok {
		actualNotPointer, ok := a.(OrganizationAddonsConfigAdvancedApiOpsConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OrganizationAddonsConfigAdvancedApiOpsConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateOrganizationSetAddonsOperation")}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareOrganizationAddonsConfigMonetizationConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*OrganizationAddonsConfigMonetizationConfig)
	if !ok {
		desiredNotPointer, ok := d.(OrganizationAddonsConfigMonetizationConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OrganizationAddonsConfigMonetizationConfig or *OrganizationAddonsConfigMonetizationConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*OrganizationAddonsConfigMonetizationConfig)
	if !ok {
		actualNotPointer, ok := a.(OrganizationAddonsConfigMonetizationConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a OrganizationAddonsConfigMonetizationConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateOrganizationSetAddonsOperation")}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Organization) urlNormalized() *Organization {
	normalized := dcl.Copy(*r).(Organization)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.AnalyticsRegion = dcl.SelfLinkToName(r.AnalyticsRegion)
	normalized.AuthorizedNetwork = dcl.SelfLinkToName(r.AuthorizedNetwork)
	normalized.CaCertificate = dcl.SelfLinkToName(r.CaCertificate)
	normalized.RuntimeDatabaseEncryptionKeyName = dcl.SelfLinkToName(r.RuntimeDatabaseEncryptionKeyName)
	normalized.ProjectId = dcl.SelfLinkToName(r.ProjectId)
	normalized.Project = r.Project
	return &normalized
}

func (r *Organization) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "SetAddons" {
		fields := map[string]interface{}{
			"name": dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("organizations/{{name}}:setAddons", nr.basePath(), userBasePath, fields), nil

	}
	if updateName == "UpdateOrganization" {
		fields := map[string]interface{}{
			"name": dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("organizations/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Organization resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Organization) marshal(c *Client) ([]byte, error) {
	m, err := expandOrganization(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Organization: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalOrganization decodes JSON responses into the Organization resource schema.
func unmarshalOrganization(b []byte, c *Client, res *Organization) (*Organization, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapOrganization(m, c, res)
}

func unmarshalMapOrganization(m map[string]interface{}, c *Client, res *Organization) (*Organization, error) {
	if v, err := dcl.MapFromListOfKeyValues(m, []string{"properties", "property"}, "name", "value"); err != nil {
		return nil, err
	} else {
		dcl.PutMapEntry(
			m,
			[]string{"properties"},
			v,
		)
	}

	flattened := flattenOrganization(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandOrganization expands Organization into a JSON request object.
func expandOrganization(c *Client, f *Organization) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("organizations/%s", f.Name, dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.DisplayName; dcl.ValueShouldBeSent(v) {
		m["displayName"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v, err := dcl.ListOfKeyValuesFromMapInStruct(f.Properties, "property", "name", "value"); err != nil {
		return nil, fmt.Errorf("error expanding Properties into properties: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["properties"] = v
	}
	if v := f.AnalyticsRegion; dcl.ValueShouldBeSent(v) {
		m["analyticsRegion"] = v
	}
	if v := f.AuthorizedNetwork; dcl.ValueShouldBeSent(v) {
		m["authorizedNetwork"] = v
	}
	if v := f.RuntimeType; dcl.ValueShouldBeSent(v) {
		m["runtimeType"] = v
	}
	if v, err := expandOrganizationAddonsConfig(c, f.AddonsConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding AddonsConfig into addonsConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["addonsConfig"] = v
	}
	if v := f.RuntimeDatabaseEncryptionKeyName; dcl.ValueShouldBeSent(v) {
		m["runtimeDatabaseEncryptionKeyName"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}

	return m, nil
}

// flattenOrganization flattens Organization from a JSON request object into the
// Organization type.
func flattenOrganization(c *Client, i interface{}, res *Organization) *Organization {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Organization{}
	resultRes.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	resultRes.DisplayName = dcl.FlattenString(m["displayName"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.CreatedAt = dcl.FlattenInteger(m["createdAt"])
	resultRes.LastModifiedAt = dcl.FlattenInteger(m["lastModifiedAt"])
	resultRes.ExpiresAt = dcl.FlattenInteger(m["expiresAt"])
	resultRes.Environments = dcl.FlattenStringSlice(m["environments"])
	resultRes.Properties = dcl.FlattenKeyValuePairs(m["properties"])
	resultRes.AnalyticsRegion = dcl.FlattenString(m["analyticsRegion"])
	resultRes.AuthorizedNetwork = dcl.FlattenString(m["authorizedNetwork"])
	resultRes.RuntimeType = flattenOrganizationRuntimeTypeEnum(m["runtimeType"])
	resultRes.SubscriptionType = flattenOrganizationSubscriptionTypeEnum(m["subscriptionType"])
	resultRes.BillingType = flattenOrganizationBillingTypeEnum(m["billingType"])
	resultRes.AddonsConfig = flattenOrganizationAddonsConfig(c, m["addonsConfig"], res)
	resultRes.CaCertificate = dcl.FlattenString(m["caCertificate"])
	resultRes.RuntimeDatabaseEncryptionKeyName = dcl.FlattenString(m["runtimeDatabaseEncryptionKeyName"])
	resultRes.ProjectId = dcl.FlattenString(m["projectId"])
	resultRes.State = flattenOrganizationStateEnum(m["state"])
	resultRes.Project = dcl.FlattenString(m["project"])

	return resultRes
}

// expandOrganizationAddonsConfigMap expands the contents of OrganizationAddonsConfig into a JSON
// request object.
func expandOrganizationAddonsConfigMap(c *Client, f map[string]OrganizationAddonsConfig, res *Organization) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOrganizationAddonsConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOrganizationAddonsConfigSlice expands the contents of OrganizationAddonsConfig into a JSON
// request object.
func expandOrganizationAddonsConfigSlice(c *Client, f []OrganizationAddonsConfig, res *Organization) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOrganizationAddonsConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOrganizationAddonsConfigMap flattens the contents of OrganizationAddonsConfig from a JSON
// response object.
func flattenOrganizationAddonsConfigMap(c *Client, i interface{}, res *Organization) map[string]OrganizationAddonsConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OrganizationAddonsConfig{}
	}

	if len(a) == 0 {
		return map[string]OrganizationAddonsConfig{}
	}

	items := make(map[string]OrganizationAddonsConfig)
	for k, item := range a {
		items[k] = *flattenOrganizationAddonsConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenOrganizationAddonsConfigSlice flattens the contents of OrganizationAddonsConfig from a JSON
// response object.
func flattenOrganizationAddonsConfigSlice(c *Client, i interface{}, res *Organization) []OrganizationAddonsConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []OrganizationAddonsConfig{}
	}

	if len(a) == 0 {
		return []OrganizationAddonsConfig{}
	}

	items := make([]OrganizationAddonsConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOrganizationAddonsConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandOrganizationAddonsConfig expands an instance of OrganizationAddonsConfig into a JSON
// request object.
func expandOrganizationAddonsConfig(c *Client, f *OrganizationAddonsConfig, res *Organization) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandOrganizationAddonsConfigAdvancedApiOpsConfig(c, f.AdvancedApiOpsConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding AdvancedApiOpsConfig into advancedApiOpsConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["advancedApiOpsConfig"] = v
	}
	if v, err := expandOrganizationAddonsConfigMonetizationConfig(c, f.MonetizationConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding MonetizationConfig into monetizationConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["monetizationConfig"] = v
	}

	return m, nil
}

// flattenOrganizationAddonsConfig flattens an instance of OrganizationAddonsConfig from a JSON
// response object.
func flattenOrganizationAddonsConfig(c *Client, i interface{}, res *Organization) *OrganizationAddonsConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OrganizationAddonsConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOrganizationAddonsConfig
	}
	r.AdvancedApiOpsConfig = flattenOrganizationAddonsConfigAdvancedApiOpsConfig(c, m["advancedApiOpsConfig"], res)
	r.MonetizationConfig = flattenOrganizationAddonsConfigMonetizationConfig(c, m["monetizationConfig"], res)

	return r
}

// expandOrganizationAddonsConfigAdvancedApiOpsConfigMap expands the contents of OrganizationAddonsConfigAdvancedApiOpsConfig into a JSON
// request object.
func expandOrganizationAddonsConfigAdvancedApiOpsConfigMap(c *Client, f map[string]OrganizationAddonsConfigAdvancedApiOpsConfig, res *Organization) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOrganizationAddonsConfigAdvancedApiOpsConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOrganizationAddonsConfigAdvancedApiOpsConfigSlice expands the contents of OrganizationAddonsConfigAdvancedApiOpsConfig into a JSON
// request object.
func expandOrganizationAddonsConfigAdvancedApiOpsConfigSlice(c *Client, f []OrganizationAddonsConfigAdvancedApiOpsConfig, res *Organization) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOrganizationAddonsConfigAdvancedApiOpsConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOrganizationAddonsConfigAdvancedApiOpsConfigMap flattens the contents of OrganizationAddonsConfigAdvancedApiOpsConfig from a JSON
// response object.
func flattenOrganizationAddonsConfigAdvancedApiOpsConfigMap(c *Client, i interface{}, res *Organization) map[string]OrganizationAddonsConfigAdvancedApiOpsConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OrganizationAddonsConfigAdvancedApiOpsConfig{}
	}

	if len(a) == 0 {
		return map[string]OrganizationAddonsConfigAdvancedApiOpsConfig{}
	}

	items := make(map[string]OrganizationAddonsConfigAdvancedApiOpsConfig)
	for k, item := range a {
		items[k] = *flattenOrganizationAddonsConfigAdvancedApiOpsConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenOrganizationAddonsConfigAdvancedApiOpsConfigSlice flattens the contents of OrganizationAddonsConfigAdvancedApiOpsConfig from a JSON
// response object.
func flattenOrganizationAddonsConfigAdvancedApiOpsConfigSlice(c *Client, i interface{}, res *Organization) []OrganizationAddonsConfigAdvancedApiOpsConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []OrganizationAddonsConfigAdvancedApiOpsConfig{}
	}

	if len(a) == 0 {
		return []OrganizationAddonsConfigAdvancedApiOpsConfig{}
	}

	items := make([]OrganizationAddonsConfigAdvancedApiOpsConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOrganizationAddonsConfigAdvancedApiOpsConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandOrganizationAddonsConfigAdvancedApiOpsConfig expands an instance of OrganizationAddonsConfigAdvancedApiOpsConfig into a JSON
// request object.
func expandOrganizationAddonsConfigAdvancedApiOpsConfig(c *Client, f *OrganizationAddonsConfigAdvancedApiOpsConfig, res *Organization) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}

	return m, nil
}

// flattenOrganizationAddonsConfigAdvancedApiOpsConfig flattens an instance of OrganizationAddonsConfigAdvancedApiOpsConfig from a JSON
// response object.
func flattenOrganizationAddonsConfigAdvancedApiOpsConfig(c *Client, i interface{}, res *Organization) *OrganizationAddonsConfigAdvancedApiOpsConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OrganizationAddonsConfigAdvancedApiOpsConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOrganizationAddonsConfigAdvancedApiOpsConfig
	}
	r.Enabled = dcl.FlattenBool(m["enabled"])

	return r
}

// expandOrganizationAddonsConfigMonetizationConfigMap expands the contents of OrganizationAddonsConfigMonetizationConfig into a JSON
// request object.
func expandOrganizationAddonsConfigMonetizationConfigMap(c *Client, f map[string]OrganizationAddonsConfigMonetizationConfig, res *Organization) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandOrganizationAddonsConfigMonetizationConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandOrganizationAddonsConfigMonetizationConfigSlice expands the contents of OrganizationAddonsConfigMonetizationConfig into a JSON
// request object.
func expandOrganizationAddonsConfigMonetizationConfigSlice(c *Client, f []OrganizationAddonsConfigMonetizationConfig, res *Organization) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandOrganizationAddonsConfigMonetizationConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenOrganizationAddonsConfigMonetizationConfigMap flattens the contents of OrganizationAddonsConfigMonetizationConfig from a JSON
// response object.
func flattenOrganizationAddonsConfigMonetizationConfigMap(c *Client, i interface{}, res *Organization) map[string]OrganizationAddonsConfigMonetizationConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OrganizationAddonsConfigMonetizationConfig{}
	}

	if len(a) == 0 {
		return map[string]OrganizationAddonsConfigMonetizationConfig{}
	}

	items := make(map[string]OrganizationAddonsConfigMonetizationConfig)
	for k, item := range a {
		items[k] = *flattenOrganizationAddonsConfigMonetizationConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenOrganizationAddonsConfigMonetizationConfigSlice flattens the contents of OrganizationAddonsConfigMonetizationConfig from a JSON
// response object.
func flattenOrganizationAddonsConfigMonetizationConfigSlice(c *Client, i interface{}, res *Organization) []OrganizationAddonsConfigMonetizationConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []OrganizationAddonsConfigMonetizationConfig{}
	}

	if len(a) == 0 {
		return []OrganizationAddonsConfigMonetizationConfig{}
	}

	items := make([]OrganizationAddonsConfigMonetizationConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOrganizationAddonsConfigMonetizationConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandOrganizationAddonsConfigMonetizationConfig expands an instance of OrganizationAddonsConfigMonetizationConfig into a JSON
// request object.
func expandOrganizationAddonsConfigMonetizationConfig(c *Client, f *OrganizationAddonsConfigMonetizationConfig, res *Organization) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}

	return m, nil
}

// flattenOrganizationAddonsConfigMonetizationConfig flattens an instance of OrganizationAddonsConfigMonetizationConfig from a JSON
// response object.
func flattenOrganizationAddonsConfigMonetizationConfig(c *Client, i interface{}, res *Organization) *OrganizationAddonsConfigMonetizationConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &OrganizationAddonsConfigMonetizationConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyOrganizationAddonsConfigMonetizationConfig
	}
	r.Enabled = dcl.FlattenBool(m["enabled"])

	return r
}

// flattenOrganizationRuntimeTypeEnumMap flattens the contents of OrganizationRuntimeTypeEnum from a JSON
// response object.
func flattenOrganizationRuntimeTypeEnumMap(c *Client, i interface{}, res *Organization) map[string]OrganizationRuntimeTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OrganizationRuntimeTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]OrganizationRuntimeTypeEnum{}
	}

	items := make(map[string]OrganizationRuntimeTypeEnum)
	for k, item := range a {
		items[k] = *flattenOrganizationRuntimeTypeEnum(item.(interface{}))
	}

	return items
}

// flattenOrganizationRuntimeTypeEnumSlice flattens the contents of OrganizationRuntimeTypeEnum from a JSON
// response object.
func flattenOrganizationRuntimeTypeEnumSlice(c *Client, i interface{}, res *Organization) []OrganizationRuntimeTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []OrganizationRuntimeTypeEnum{}
	}

	if len(a) == 0 {
		return []OrganizationRuntimeTypeEnum{}
	}

	items := make([]OrganizationRuntimeTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOrganizationRuntimeTypeEnum(item.(interface{})))
	}

	return items
}

// flattenOrganizationRuntimeTypeEnum asserts that an interface is a string, and returns a
// pointer to a *OrganizationRuntimeTypeEnum with the same value as that string.
func flattenOrganizationRuntimeTypeEnum(i interface{}) *OrganizationRuntimeTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return OrganizationRuntimeTypeEnumRef(s)
}

// flattenOrganizationSubscriptionTypeEnumMap flattens the contents of OrganizationSubscriptionTypeEnum from a JSON
// response object.
func flattenOrganizationSubscriptionTypeEnumMap(c *Client, i interface{}, res *Organization) map[string]OrganizationSubscriptionTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OrganizationSubscriptionTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]OrganizationSubscriptionTypeEnum{}
	}

	items := make(map[string]OrganizationSubscriptionTypeEnum)
	for k, item := range a {
		items[k] = *flattenOrganizationSubscriptionTypeEnum(item.(interface{}))
	}

	return items
}

// flattenOrganizationSubscriptionTypeEnumSlice flattens the contents of OrganizationSubscriptionTypeEnum from a JSON
// response object.
func flattenOrganizationSubscriptionTypeEnumSlice(c *Client, i interface{}, res *Organization) []OrganizationSubscriptionTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []OrganizationSubscriptionTypeEnum{}
	}

	if len(a) == 0 {
		return []OrganizationSubscriptionTypeEnum{}
	}

	items := make([]OrganizationSubscriptionTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOrganizationSubscriptionTypeEnum(item.(interface{})))
	}

	return items
}

// flattenOrganizationSubscriptionTypeEnum asserts that an interface is a string, and returns a
// pointer to a *OrganizationSubscriptionTypeEnum with the same value as that string.
func flattenOrganizationSubscriptionTypeEnum(i interface{}) *OrganizationSubscriptionTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return OrganizationSubscriptionTypeEnumRef(s)
}

// flattenOrganizationBillingTypeEnumMap flattens the contents of OrganizationBillingTypeEnum from a JSON
// response object.
func flattenOrganizationBillingTypeEnumMap(c *Client, i interface{}, res *Organization) map[string]OrganizationBillingTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OrganizationBillingTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]OrganizationBillingTypeEnum{}
	}

	items := make(map[string]OrganizationBillingTypeEnum)
	for k, item := range a {
		items[k] = *flattenOrganizationBillingTypeEnum(item.(interface{}))
	}

	return items
}

// flattenOrganizationBillingTypeEnumSlice flattens the contents of OrganizationBillingTypeEnum from a JSON
// response object.
func flattenOrganizationBillingTypeEnumSlice(c *Client, i interface{}, res *Organization) []OrganizationBillingTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []OrganizationBillingTypeEnum{}
	}

	if len(a) == 0 {
		return []OrganizationBillingTypeEnum{}
	}

	items := make([]OrganizationBillingTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOrganizationBillingTypeEnum(item.(interface{})))
	}

	return items
}

// flattenOrganizationBillingTypeEnum asserts that an interface is a string, and returns a
// pointer to a *OrganizationBillingTypeEnum with the same value as that string.
func flattenOrganizationBillingTypeEnum(i interface{}) *OrganizationBillingTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return OrganizationBillingTypeEnumRef(s)
}

// flattenOrganizationStateEnumMap flattens the contents of OrganizationStateEnum from a JSON
// response object.
func flattenOrganizationStateEnumMap(c *Client, i interface{}, res *Organization) map[string]OrganizationStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]OrganizationStateEnum{}
	}

	if len(a) == 0 {
		return map[string]OrganizationStateEnum{}
	}

	items := make(map[string]OrganizationStateEnum)
	for k, item := range a {
		items[k] = *flattenOrganizationStateEnum(item.(interface{}))
	}

	return items
}

// flattenOrganizationStateEnumSlice flattens the contents of OrganizationStateEnum from a JSON
// response object.
func flattenOrganizationStateEnumSlice(c *Client, i interface{}, res *Organization) []OrganizationStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []OrganizationStateEnum{}
	}

	if len(a) == 0 {
		return []OrganizationStateEnum{}
	}

	items := make([]OrganizationStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenOrganizationStateEnum(item.(interface{})))
	}

	return items
}

// flattenOrganizationStateEnum asserts that an interface is a string, and returns a
// pointer to a *OrganizationStateEnum with the same value as that string.
func flattenOrganizationStateEnum(i interface{}) *OrganizationStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return OrganizationStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Organization) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalOrganization(b, c, r)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Name == nil && ncr.Name == nil {
			c.Config.Logger.Info("Both Name fields null - considering equal.")
		} else if nr.Name == nil || ncr.Name == nil {
			c.Config.Logger.Info("Only one Name field is null - considering unequal.")
			return false
		} else if *nr.Name != *ncr.Name {
			return false
		}
		return true
	}
}

type organizationDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         organizationApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToOrganizationDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]organizationDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff: %v", ro, fd)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []organizationDiff
	// For each operation name, create a organizationDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := organizationDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToOrganizationApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToOrganizationApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (organizationApiOperation, error) {
	switch opName {

	case "updateOrganizationSetAddonsOperation":
		return &updateOrganizationSetAddonsOperation{FieldDiffs: fieldDiffs}, nil

	case "updateOrganizationUpdateOrganizationOperation":
		return &updateOrganizationUpdateOrganizationOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractOrganizationFields(r *Organization) error {
	vAddonsConfig := r.AddonsConfig
	if vAddonsConfig == nil {
		// note: explicitly not the empty object.
		vAddonsConfig = &OrganizationAddonsConfig{}
	}
	if err := extractOrganizationAddonsConfigFields(r, vAddonsConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAddonsConfig) {
		r.AddonsConfig = vAddonsConfig
	}
	return nil
}
func extractOrganizationAddonsConfigFields(r *Organization, o *OrganizationAddonsConfig) error {
	vAdvancedApiOpsConfig := o.AdvancedApiOpsConfig
	if vAdvancedApiOpsConfig == nil {
		// note: explicitly not the empty object.
		vAdvancedApiOpsConfig = &OrganizationAddonsConfigAdvancedApiOpsConfig{}
	}
	if err := extractOrganizationAddonsConfigAdvancedApiOpsConfigFields(r, vAdvancedApiOpsConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAdvancedApiOpsConfig) {
		o.AdvancedApiOpsConfig = vAdvancedApiOpsConfig
	}
	vMonetizationConfig := o.MonetizationConfig
	if vMonetizationConfig == nil {
		// note: explicitly not the empty object.
		vMonetizationConfig = &OrganizationAddonsConfigMonetizationConfig{}
	}
	if err := extractOrganizationAddonsConfigMonetizationConfigFields(r, vMonetizationConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMonetizationConfig) {
		o.MonetizationConfig = vMonetizationConfig
	}
	return nil
}
func extractOrganizationAddonsConfigAdvancedApiOpsConfigFields(r *Organization, o *OrganizationAddonsConfigAdvancedApiOpsConfig) error {
	return nil
}
func extractOrganizationAddonsConfigMonetizationConfigFields(r *Organization, o *OrganizationAddonsConfigMonetizationConfig) error {
	return nil
}

func postReadExtractOrganizationFields(r *Organization) error {
	vAddonsConfig := r.AddonsConfig
	if vAddonsConfig == nil {
		// note: explicitly not the empty object.
		vAddonsConfig = &OrganizationAddonsConfig{}
	}
	if err := postReadExtractOrganizationAddonsConfigFields(r, vAddonsConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAddonsConfig) {
		r.AddonsConfig = vAddonsConfig
	}
	return nil
}
func postReadExtractOrganizationAddonsConfigFields(r *Organization, o *OrganizationAddonsConfig) error {
	vAdvancedApiOpsConfig := o.AdvancedApiOpsConfig
	if vAdvancedApiOpsConfig == nil {
		// note: explicitly not the empty object.
		vAdvancedApiOpsConfig = &OrganizationAddonsConfigAdvancedApiOpsConfig{}
	}
	if err := extractOrganizationAddonsConfigAdvancedApiOpsConfigFields(r, vAdvancedApiOpsConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAdvancedApiOpsConfig) {
		o.AdvancedApiOpsConfig = vAdvancedApiOpsConfig
	}
	vMonetizationConfig := o.MonetizationConfig
	if vMonetizationConfig == nil {
		// note: explicitly not the empty object.
		vMonetizationConfig = &OrganizationAddonsConfigMonetizationConfig{}
	}
	if err := extractOrganizationAddonsConfigMonetizationConfigFields(r, vMonetizationConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMonetizationConfig) {
		o.MonetizationConfig = vMonetizationConfig
	}
	return nil
}
func postReadExtractOrganizationAddonsConfigAdvancedApiOpsConfigFields(r *Organization, o *OrganizationAddonsConfigAdvancedApiOpsConfig) error {
	return nil
}
func postReadExtractOrganizationAddonsConfigMonetizationConfigFields(r *Organization, o *OrganizationAddonsConfigMonetizationConfig) error {
	return nil
}
