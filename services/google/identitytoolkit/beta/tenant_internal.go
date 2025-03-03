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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *Tenant) validate() error {

	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.MfaConfig) {
		if err := r.MfaConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *TenantMfaConfig) validate() error {
	return nil
}
func (r *Tenant) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://identitytoolkit.googleapis.com/v2", params)
}

func (r *Tenant) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/tenants/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Tenant) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.URL("projects/{{project}}/tenants", nr.basePath(), userBasePath, params), nil

}

func (r *Tenant) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.URL("projects/{{project}}/tenants", nr.basePath(), userBasePath, params), nil

}

func (r *Tenant) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/tenants/{{name}}", nr.basePath(), userBasePath, params), nil
}

// tenantApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type tenantApiOperation interface {
	do(context.Context, *Tenant, *Client) error
}

// newUpdateTenantUpdateTenantRequest creates a request for an
// Tenant resource's UpdateTenant update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateTenantUpdateTenantRequest(ctx context.Context, f *Tenant, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v := f.AllowPasswordSignup; !dcl.IsEmptyValueIndirect(v) {
		req["allowPasswordSignup"] = v
	}
	if v := f.EnableEmailLinkSignin; !dcl.IsEmptyValueIndirect(v) {
		req["enableEmailLinkSignin"] = v
	}
	if v := f.DisableAuth; !dcl.IsEmptyValueIndirect(v) {
		req["disableAuth"] = v
	}
	if v := f.EnableAnonymousUser; !dcl.IsEmptyValueIndirect(v) {
		req["enableAnonymousUser"] = v
	}
	if v, err := expandTenantMfaConfig(c, f.MfaConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding MfaConfig into mfaConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["mfaConfig"] = v
	}
	if v := f.TestPhoneNumbers; !dcl.IsEmptyValueIndirect(v) {
		req["testPhoneNumbers"] = v
	}
	return req, nil
}

// marshalUpdateTenantUpdateTenantRequest converts the update into
// the final JSON request body.
func marshalUpdateTenantUpdateTenantRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateTenantUpdateTenantOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateTenantUpdateTenantOperation) do(ctx context.Context, r *Tenant, c *Client) error {
	_, err := c.GetTenant(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateTenant")
	if err != nil {
		return err
	}

	req, err := newUpdateTenantUpdateTenantRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateTenantUpdateTenantRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listTenantRaw(ctx context.Context, r *Tenant, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != TenantMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listTenantOperation struct {
	Tenants []map[string]interface{} `json:"tenants"`
	Token   string                   `json:"nextPageToken"`
}

func (c *Client) listTenant(ctx context.Context, r *Tenant, pageToken string, pageSize int32) ([]*Tenant, string, error) {
	b, err := c.listTenantRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listTenantOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Tenant
	for _, v := range m.Tenants {
		res, err := unmarshalMapTenant(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllTenant(ctx context.Context, f func(*Tenant) bool, resources []*Tenant) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteTenant(ctx, res)
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

type deleteTenantOperation struct{}

func (op *deleteTenantOperation) do(ctx context.Context, r *Tenant, c *Client) error {
	r, err := c.GetTenant(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Tenant not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetTenant checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Tenant: %w", err)
	}

	// We saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// This is the reason we are adding retry to handle that case.
	retriesRemaining := 10
	dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		_, err := c.GetTenant(ctx, r)
		if dcl.IsNotFound(err) {
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
type createTenantOperation struct {
	response map[string]interface{}
}

func (op *createTenantOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createTenantOperation) do(ctx context.Context, r *Tenant, c *Client) error {
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
		// Allowing creation to continue with Name set could result in a Tenant with the wrong Name.
		return fmt.Errorf("server-generated parameter Name was specified by user as %v, should be unspecified", dcl.ValueOrEmptyString(r.Name))
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	o, err := dcl.ResponseBodyAsJSON(resp)
	if err != nil {
		return fmt.Errorf("error decoding response body into JSON: %w", err)
	}
	op.response = o

	// Include Name in URL substitution for initial GET request.
	m := op.response
	r.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))

	if _, err := c.GetTenant(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getTenantRaw(ctx context.Context, r *Tenant) ([]byte, error) {

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

func (c *Client) tenantDiffsForRawDesired(ctx context.Context, rawDesired *Tenant, opts ...dcl.ApplyOption) (initial, desired *Tenant, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Tenant
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Tenant); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Tenant, got %T", sh)
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
		desired, err := canonicalizeTenantDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}
	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetTenant(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Tenant resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Tenant resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Tenant resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeTenantDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Tenant: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Tenant: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractTenantFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeTenantInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Tenant: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeTenantDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Tenant: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffTenant(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeTenantInitialState(rawInitial, rawDesired *Tenant) (*Tenant, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeTenantDesiredState(rawDesired, rawInitial *Tenant, opts ...dcl.ApplyOption) (*Tenant, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.MfaConfig = canonicalizeTenantMfaConfig(rawDesired.MfaConfig, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Tenant{}
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
	if dcl.BoolCanonicalize(rawDesired.AllowPasswordSignup, rawInitial.AllowPasswordSignup) {
		canonicalDesired.AllowPasswordSignup = rawInitial.AllowPasswordSignup
	} else {
		canonicalDesired.AllowPasswordSignup = rawDesired.AllowPasswordSignup
	}
	if dcl.BoolCanonicalize(rawDesired.EnableEmailLinkSignin, rawInitial.EnableEmailLinkSignin) {
		canonicalDesired.EnableEmailLinkSignin = rawInitial.EnableEmailLinkSignin
	} else {
		canonicalDesired.EnableEmailLinkSignin = rawDesired.EnableEmailLinkSignin
	}
	if dcl.BoolCanonicalize(rawDesired.DisableAuth, rawInitial.DisableAuth) {
		canonicalDesired.DisableAuth = rawInitial.DisableAuth
	} else {
		canonicalDesired.DisableAuth = rawDesired.DisableAuth
	}
	if dcl.BoolCanonicalize(rawDesired.EnableAnonymousUser, rawInitial.EnableAnonymousUser) {
		canonicalDesired.EnableAnonymousUser = rawInitial.EnableAnonymousUser
	} else {
		canonicalDesired.EnableAnonymousUser = rawDesired.EnableAnonymousUser
	}
	canonicalDesired.MfaConfig = canonicalizeTenantMfaConfig(rawDesired.MfaConfig, rawInitial.MfaConfig, opts...)
	if dcl.IsZeroValue(rawDesired.TestPhoneNumbers) || (dcl.IsEmptyValueIndirect(rawDesired.TestPhoneNumbers) && dcl.IsEmptyValueIndirect(rawInitial.TestPhoneNumbers)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.TestPhoneNumbers = rawInitial.TestPhoneNumbers
	} else {
		canonicalDesired.TestPhoneNumbers = rawDesired.TestPhoneNumbers
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	return canonicalDesired, nil
}

func canonicalizeTenantNewState(c *Client, rawNew, rawDesired *Tenant) (*Tenant, error) {

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

	if dcl.IsEmptyValueIndirect(rawNew.AllowPasswordSignup) && dcl.IsEmptyValueIndirect(rawDesired.AllowPasswordSignup) {
		rawNew.AllowPasswordSignup = rawDesired.AllowPasswordSignup
	} else {
		if dcl.BoolCanonicalize(rawDesired.AllowPasswordSignup, rawNew.AllowPasswordSignup) {
			rawNew.AllowPasswordSignup = rawDesired.AllowPasswordSignup
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.EnableEmailLinkSignin) && dcl.IsEmptyValueIndirect(rawDesired.EnableEmailLinkSignin) {
		rawNew.EnableEmailLinkSignin = rawDesired.EnableEmailLinkSignin
	} else {
		if dcl.BoolCanonicalize(rawDesired.EnableEmailLinkSignin, rawNew.EnableEmailLinkSignin) {
			rawNew.EnableEmailLinkSignin = rawDesired.EnableEmailLinkSignin
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisableAuth) && dcl.IsEmptyValueIndirect(rawDesired.DisableAuth) {
		rawNew.DisableAuth = rawDesired.DisableAuth
	} else {
		if dcl.BoolCanonicalize(rawDesired.DisableAuth, rawNew.DisableAuth) {
			rawNew.DisableAuth = rawDesired.DisableAuth
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.EnableAnonymousUser) && dcl.IsEmptyValueIndirect(rawDesired.EnableAnonymousUser) {
		rawNew.EnableAnonymousUser = rawDesired.EnableAnonymousUser
	} else {
		if dcl.BoolCanonicalize(rawDesired.EnableAnonymousUser, rawNew.EnableAnonymousUser) {
			rawNew.EnableAnonymousUser = rawDesired.EnableAnonymousUser
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.MfaConfig) && dcl.IsEmptyValueIndirect(rawDesired.MfaConfig) {
		rawNew.MfaConfig = rawDesired.MfaConfig
	} else {
		rawNew.MfaConfig = canonicalizeNewTenantMfaConfig(c, rawDesired.MfaConfig, rawNew.MfaConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.TestPhoneNumbers) && dcl.IsEmptyValueIndirect(rawDesired.TestPhoneNumbers) {
		rawNew.TestPhoneNumbers = rawDesired.TestPhoneNumbers
	} else {
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeTenantMfaConfig(des, initial *TenantMfaConfig, opts ...dcl.ApplyOption) *TenantMfaConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TenantMfaConfig{}

	if dcl.IsZeroValue(des.State) || (dcl.IsEmptyValueIndirect(des.State) && dcl.IsEmptyValueIndirect(initial.State)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.State = initial.State
	} else {
		cDes.State = des.State
	}
	if dcl.IsZeroValue(des.EnabledProviders) || (dcl.IsEmptyValueIndirect(des.EnabledProviders) && dcl.IsEmptyValueIndirect(initial.EnabledProviders)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.EnabledProviders = initial.EnabledProviders
	} else {
		cDes.EnabledProviders = des.EnabledProviders
	}

	return cDes
}

func canonicalizeTenantMfaConfigSlice(des, initial []TenantMfaConfig, opts ...dcl.ApplyOption) []TenantMfaConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TenantMfaConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTenantMfaConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TenantMfaConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTenantMfaConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTenantMfaConfig(c *Client, des, nw *TenantMfaConfig) *TenantMfaConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for TenantMfaConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewTenantMfaConfigSet(c *Client, des, nw []TenantMfaConfig) []TenantMfaConfig {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []TenantMfaConfig
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareTenantMfaConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewTenantMfaConfig(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewTenantMfaConfigSlice(c *Client, des, nw []TenantMfaConfig) []TenantMfaConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TenantMfaConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTenantMfaConfig(c, &d, &n))
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
func diffTenant(c *Client, desired, actual *Tenant, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTenantUpdateTenantOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowPasswordSignup, actual.AllowPasswordSignup, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTenantUpdateTenantOperation")}, fn.AddNest("AllowPasswordSignup")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnableEmailLinkSignin, actual.EnableEmailLinkSignin, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTenantUpdateTenantOperation")}, fn.AddNest("EnableEmailLinkSignin")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisableAuth, actual.DisableAuth, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTenantUpdateTenantOperation")}, fn.AddNest("DisableAuth")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnableAnonymousUser, actual.EnableAnonymousUser, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTenantUpdateTenantOperation")}, fn.AddNest("EnableAnonymousUser")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MfaConfig, actual.MfaConfig, dcl.DiffInfo{ObjectFunction: compareTenantMfaConfigNewStyle, EmptyObject: EmptyTenantMfaConfig, OperationSelector: dcl.TriggersOperation("updateTenantUpdateTenantOperation")}, fn.AddNest("MfaConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TestPhoneNumbers, actual.TestPhoneNumbers, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTenantUpdateTenantOperation")}, fn.AddNest("TestPhoneNumbers")); len(ds) != 0 || err != nil {
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
func compareTenantMfaConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TenantMfaConfig)
	if !ok {
		desiredNotPointer, ok := d.(TenantMfaConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TenantMfaConfig or *TenantMfaConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TenantMfaConfig)
	if !ok {
		actualNotPointer, ok := a.(TenantMfaConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TenantMfaConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnabledProviders, actual.EnabledProviders, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EnabledProviders")); len(ds) != 0 || err != nil {
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
func (r *Tenant) urlNormalized() *Tenant {
	normalized := dcl.Copy(*r).(Tenant)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *Tenant) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateTenant" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(nr.Project),
			"name":    dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/tenants/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Tenant resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Tenant) marshal(c *Client) ([]byte, error) {
	m, err := expandTenant(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Tenant: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalTenant decodes JSON responses into the Tenant resource schema.
func unmarshalTenant(b []byte, c *Client, res *Tenant) (*Tenant, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapTenant(m, c, res)
}

func unmarshalMapTenant(m map[string]interface{}, c *Client, res *Tenant) (*Tenant, error) {

	flattened := flattenTenant(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandTenant expands Tenant into a JSON request object.
func expandTenant(c *Client, f *Tenant) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("projects/%s/tenants/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.DisplayName; dcl.ValueShouldBeSent(v) {
		m["displayName"] = v
	}
	if v := f.AllowPasswordSignup; dcl.ValueShouldBeSent(v) {
		m["allowPasswordSignup"] = v
	}
	if v := f.EnableEmailLinkSignin; dcl.ValueShouldBeSent(v) {
		m["enableEmailLinkSignin"] = v
	}
	if v := f.DisableAuth; dcl.ValueShouldBeSent(v) {
		m["disableAuth"] = v
	}
	if v := f.EnableAnonymousUser; dcl.ValueShouldBeSent(v) {
		m["enableAnonymousUser"] = v
	}
	if v, err := expandTenantMfaConfig(c, f.MfaConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding MfaConfig into mfaConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["mfaConfig"] = v
	}
	if v := f.TestPhoneNumbers; dcl.ValueShouldBeSent(v) {
		m["testPhoneNumbers"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}

	return m, nil
}

// flattenTenant flattens Tenant from a JSON request object into the
// Tenant type.
func flattenTenant(c *Client, i interface{}, res *Tenant) *Tenant {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Tenant{}
	resultRes.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	resultRes.DisplayName = dcl.FlattenString(m["displayName"])
	resultRes.AllowPasswordSignup = dcl.FlattenBool(m["allowPasswordSignup"])
	resultRes.EnableEmailLinkSignin = dcl.FlattenBool(m["enableEmailLinkSignin"])
	resultRes.DisableAuth = dcl.FlattenBool(m["disableAuth"])
	resultRes.EnableAnonymousUser = dcl.FlattenBool(m["enableAnonymousUser"])
	resultRes.MfaConfig = flattenTenantMfaConfig(c, m["mfaConfig"], res)
	resultRes.TestPhoneNumbers = dcl.FlattenKeyValuePairs(m["testPhoneNumbers"])
	resultRes.Project = dcl.FlattenString(m["project"])

	return resultRes
}

// expandTenantMfaConfigMap expands the contents of TenantMfaConfig into a JSON
// request object.
func expandTenantMfaConfigMap(c *Client, f map[string]TenantMfaConfig, res *Tenant) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTenantMfaConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTenantMfaConfigSlice expands the contents of TenantMfaConfig into a JSON
// request object.
func expandTenantMfaConfigSlice(c *Client, f []TenantMfaConfig, res *Tenant) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTenantMfaConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTenantMfaConfigMap flattens the contents of TenantMfaConfig from a JSON
// response object.
func flattenTenantMfaConfigMap(c *Client, i interface{}, res *Tenant) map[string]TenantMfaConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TenantMfaConfig{}
	}

	if len(a) == 0 {
		return map[string]TenantMfaConfig{}
	}

	items := make(map[string]TenantMfaConfig)
	for k, item := range a {
		items[k] = *flattenTenantMfaConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTenantMfaConfigSlice flattens the contents of TenantMfaConfig from a JSON
// response object.
func flattenTenantMfaConfigSlice(c *Client, i interface{}, res *Tenant) []TenantMfaConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []TenantMfaConfig{}
	}

	if len(a) == 0 {
		return []TenantMfaConfig{}
	}

	items := make([]TenantMfaConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTenantMfaConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTenantMfaConfig expands an instance of TenantMfaConfig into a JSON
// request object.
func expandTenantMfaConfig(c *Client, f *TenantMfaConfig, res *Tenant) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v := f.EnabledProviders; v != nil {
		m["enabledProviders"] = v
	}

	return m, nil
}

// flattenTenantMfaConfig flattens an instance of TenantMfaConfig from a JSON
// response object.
func flattenTenantMfaConfig(c *Client, i interface{}, res *Tenant) *TenantMfaConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TenantMfaConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTenantMfaConfig
	}
	r.State = flattenTenantMfaConfigStateEnum(m["state"])
	r.EnabledProviders = flattenTenantMfaConfigEnabledProvidersEnumSlice(c, m["enabledProviders"], res)

	return r
}

// flattenTenantMfaConfigStateEnumMap flattens the contents of TenantMfaConfigStateEnum from a JSON
// response object.
func flattenTenantMfaConfigStateEnumMap(c *Client, i interface{}, res *Tenant) map[string]TenantMfaConfigStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TenantMfaConfigStateEnum{}
	}

	if len(a) == 0 {
		return map[string]TenantMfaConfigStateEnum{}
	}

	items := make(map[string]TenantMfaConfigStateEnum)
	for k, item := range a {
		items[k] = *flattenTenantMfaConfigStateEnum(item.(interface{}))
	}

	return items
}

// flattenTenantMfaConfigStateEnumSlice flattens the contents of TenantMfaConfigStateEnum from a JSON
// response object.
func flattenTenantMfaConfigStateEnumSlice(c *Client, i interface{}, res *Tenant) []TenantMfaConfigStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []TenantMfaConfigStateEnum{}
	}

	if len(a) == 0 {
		return []TenantMfaConfigStateEnum{}
	}

	items := make([]TenantMfaConfigStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTenantMfaConfigStateEnum(item.(interface{})))
	}

	return items
}

// flattenTenantMfaConfigStateEnum asserts that an interface is a string, and returns a
// pointer to a *TenantMfaConfigStateEnum with the same value as that string.
func flattenTenantMfaConfigStateEnum(i interface{}) *TenantMfaConfigStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return TenantMfaConfigStateEnumRef(s)
}

// flattenTenantMfaConfigEnabledProvidersEnumMap flattens the contents of TenantMfaConfigEnabledProvidersEnum from a JSON
// response object.
func flattenTenantMfaConfigEnabledProvidersEnumMap(c *Client, i interface{}, res *Tenant) map[string]TenantMfaConfigEnabledProvidersEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TenantMfaConfigEnabledProvidersEnum{}
	}

	if len(a) == 0 {
		return map[string]TenantMfaConfigEnabledProvidersEnum{}
	}

	items := make(map[string]TenantMfaConfigEnabledProvidersEnum)
	for k, item := range a {
		items[k] = *flattenTenantMfaConfigEnabledProvidersEnum(item.(interface{}))
	}

	return items
}

// flattenTenantMfaConfigEnabledProvidersEnumSlice flattens the contents of TenantMfaConfigEnabledProvidersEnum from a JSON
// response object.
func flattenTenantMfaConfigEnabledProvidersEnumSlice(c *Client, i interface{}, res *Tenant) []TenantMfaConfigEnabledProvidersEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []TenantMfaConfigEnabledProvidersEnum{}
	}

	if len(a) == 0 {
		return []TenantMfaConfigEnabledProvidersEnum{}
	}

	items := make([]TenantMfaConfigEnabledProvidersEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTenantMfaConfigEnabledProvidersEnum(item.(interface{})))
	}

	return items
}

// flattenTenantMfaConfigEnabledProvidersEnum asserts that an interface is a string, and returns a
// pointer to a *TenantMfaConfigEnabledProvidersEnum with the same value as that string.
func flattenTenantMfaConfigEnabledProvidersEnum(i interface{}) *TenantMfaConfigEnabledProvidersEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return TenantMfaConfigEnabledProvidersEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Tenant) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalTenant(b, c, r)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Project == nil && ncr.Project == nil {
			c.Config.Logger.Info("Both Project fields null - considering equal.")
		} else if nr.Project == nil || ncr.Project == nil {
			c.Config.Logger.Info("Only one Project field is null - considering unequal.")
			return false
		} else if *nr.Project != *ncr.Project {
			return false
		}
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

type tenantDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         tenantApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToTenantDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]tenantDiff, error) {
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
	var diffs []tenantDiff
	// For each operation name, create a tenantDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := tenantDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToTenantApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToTenantApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (tenantApiOperation, error) {
	switch opName {

	case "updateTenantUpdateTenantOperation":
		return &updateTenantUpdateTenantOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractTenantFields(r *Tenant) error {
	vMfaConfig := r.MfaConfig
	if vMfaConfig == nil {
		// note: explicitly not the empty object.
		vMfaConfig = &TenantMfaConfig{}
	}
	if err := extractTenantMfaConfigFields(r, vMfaConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMfaConfig) {
		r.MfaConfig = vMfaConfig
	}
	return nil
}
func extractTenantMfaConfigFields(r *Tenant, o *TenantMfaConfig) error {
	return nil
}

func postReadExtractTenantFields(r *Tenant) error {
	vMfaConfig := r.MfaConfig
	if vMfaConfig == nil {
		// note: explicitly not the empty object.
		vMfaConfig = &TenantMfaConfig{}
	}
	if err := postReadExtractTenantMfaConfigFields(r, vMfaConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMfaConfig) {
		r.MfaConfig = vMfaConfig
	}
	return nil
}
func postReadExtractTenantMfaConfigFields(r *Tenant, o *TenantMfaConfig) error {
	return nil
}
