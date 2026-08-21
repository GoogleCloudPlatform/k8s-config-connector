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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *TenantOAuthIdpConfig) validate() error {

	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Tenant, "Tenant"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.ResponseType) {
		if err := r.ResponseType.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *TenantOAuthIdpConfigResponseType) validate() error {
	return nil
}
func (r *TenantOAuthIdpConfig) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://identitytoolkit.googleapis.com/admin/v2", params)
}

func (r *TenantOAuthIdpConfig) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"tenant":  dcl.ValueOrEmptyString(nr.Tenant),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/tenants/{{tenant}}/oauthIdpConfigs/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *TenantOAuthIdpConfig) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"tenant":  dcl.ValueOrEmptyString(nr.Tenant),
	}
	return dcl.URL("projects/{{project}}/tenants/{{tenant}}/oauthIdpConfigs", nr.basePath(), userBasePath, params), nil

}

func (r *TenantOAuthIdpConfig) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"tenant":  dcl.ValueOrEmptyString(nr.Tenant),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/tenants/{{tenant}}/oauthIdpConfigs?oauthIdpConfigId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *TenantOAuthIdpConfig) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"tenant":  dcl.ValueOrEmptyString(nr.Tenant),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/tenants/{{tenant}}/oauthIdpConfigs/{{name}}", nr.basePath(), userBasePath, params), nil
}

// tenantOAuthIdpConfigApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type tenantOAuthIdpConfigApiOperation interface {
	do(context.Context, *TenantOAuthIdpConfig, *Client) error
}

// newUpdateTenantOAuthIdpConfigUpdateConfigRequest creates a request for an
// TenantOAuthIdpConfig resource's UpdateConfig update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateTenantOAuthIdpConfigUpdateConfigRequest(ctx context.Context, f *TenantOAuthIdpConfig, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v := f.ClientId; !dcl.IsEmptyValueIndirect(v) {
		req["clientId"] = v
	}
	if v := f.Issuer; !dcl.IsEmptyValueIndirect(v) {
		req["issuer"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		req["enabled"] = v
	}
	if v := f.ClientSecret; !dcl.IsEmptyValueIndirect(v) {
		req["clientSecret"] = v
	}
	if v, err := expandTenantOAuthIdpConfigResponseType(c, f.ResponseType, res); err != nil {
		return nil, fmt.Errorf("error expanding ResponseType into responseType: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["responseType"] = v
	}
	return req, nil
}

// marshalUpdateTenantOAuthIdpConfigUpdateConfigRequest converts the update into
// the final JSON request body.
func marshalUpdateTenantOAuthIdpConfigUpdateConfigRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	dcl.MoveMapEntry(
		m,
		[]string{"name"},
		[]string{},
	)
	return json.Marshal(m)
}

type updateTenantOAuthIdpConfigUpdateConfigOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateTenantOAuthIdpConfigUpdateConfigOperation) do(ctx context.Context, r *TenantOAuthIdpConfig, c *Client) error {
	_, err := c.GetTenantOAuthIdpConfig(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateConfig")
	if err != nil {
		return err
	}

	req, err := newUpdateTenantOAuthIdpConfigUpdateConfigRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateTenantOAuthIdpConfigUpdateConfigRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listTenantOAuthIdpConfigRaw(ctx context.Context, r *TenantOAuthIdpConfig, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != TenantOAuthIdpConfigMaxPage {
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

type listTenantOAuthIdpConfigOperation struct {
	OauthIdpConfigs []map[string]interface{} `json:"oauthIdpConfigs"`
	Token           string                   `json:"nextPageToken"`
}

func (c *Client) listTenantOAuthIdpConfig(ctx context.Context, r *TenantOAuthIdpConfig, pageToken string, pageSize int32) ([]*TenantOAuthIdpConfig, string, error) {
	b, err := c.listTenantOAuthIdpConfigRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listTenantOAuthIdpConfigOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*TenantOAuthIdpConfig
	for _, v := range m.OauthIdpConfigs {
		res, err := unmarshalMapTenantOAuthIdpConfig(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Tenant = r.Tenant
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllTenantOAuthIdpConfig(ctx context.Context, f func(*TenantOAuthIdpConfig) bool, resources []*TenantOAuthIdpConfig) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteTenantOAuthIdpConfig(ctx, res)
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

type deleteTenantOAuthIdpConfigOperation struct{}

func (op *deleteTenantOAuthIdpConfigOperation) do(ctx context.Context, r *TenantOAuthIdpConfig, c *Client) error {
	r, err := c.GetTenantOAuthIdpConfig(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "TenantOAuthIdpConfig not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetTenantOAuthIdpConfig checking for existence. error: %v", err)
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
		return fmt.Errorf("failed to delete TenantOAuthIdpConfig: %w", err)
	}

	// We saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// This is the reason we are adding retry to handle that case.
	retriesRemaining := 10
	dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		_, err := c.GetTenantOAuthIdpConfig(ctx, r)
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
type createTenantOAuthIdpConfigOperation struct {
	response map[string]interface{}
}

func (op *createTenantOAuthIdpConfigOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createTenantOAuthIdpConfigOperation) do(ctx context.Context, r *TenantOAuthIdpConfig, c *Client) error {
	c.Config.Logger.InfoWithContextf(ctx, "Attempting to create %v", r)
	u, err := r.createURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
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

	if _, err := c.GetTenantOAuthIdpConfig(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getTenantOAuthIdpConfigRaw(ctx context.Context, r *TenantOAuthIdpConfig) ([]byte, error) {

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

func (c *Client) tenantOAuthIdpConfigDiffsForRawDesired(ctx context.Context, rawDesired *TenantOAuthIdpConfig, opts ...dcl.ApplyOption) (initial, desired *TenantOAuthIdpConfig, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *TenantOAuthIdpConfig
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*TenantOAuthIdpConfig); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected TenantOAuthIdpConfig, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetTenantOAuthIdpConfig(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a TenantOAuthIdpConfig resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve TenantOAuthIdpConfig resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that TenantOAuthIdpConfig resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeTenantOAuthIdpConfigDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for TenantOAuthIdpConfig: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for TenantOAuthIdpConfig: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractTenantOAuthIdpConfigFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeTenantOAuthIdpConfigInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for TenantOAuthIdpConfig: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeTenantOAuthIdpConfigDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for TenantOAuthIdpConfig: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffTenantOAuthIdpConfig(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeTenantOAuthIdpConfigInitialState(rawInitial, rawDesired *TenantOAuthIdpConfig) (*TenantOAuthIdpConfig, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeTenantOAuthIdpConfigDesiredState(rawDesired, rawInitial *TenantOAuthIdpConfig, opts ...dcl.ApplyOption) (*TenantOAuthIdpConfig, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.ResponseType = canonicalizeTenantOAuthIdpConfigResponseType(rawDesired.ResponseType, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &TenantOAuthIdpConfig{}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.ClientId, rawInitial.ClientId) {
		canonicalDesired.ClientId = rawInitial.ClientId
	} else {
		canonicalDesired.ClientId = rawDesired.ClientId
	}
	if dcl.StringCanonicalize(rawDesired.Issuer, rawInitial.Issuer) {
		canonicalDesired.Issuer = rawInitial.Issuer
	} else {
		canonicalDesired.Issuer = rawDesired.Issuer
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		canonicalDesired.DisplayName = rawInitial.DisplayName
	} else {
		canonicalDesired.DisplayName = rawDesired.DisplayName
	}
	if dcl.BoolCanonicalize(rawDesired.Enabled, rawInitial.Enabled) {
		canonicalDesired.Enabled = rawInitial.Enabled
	} else {
		canonicalDesired.Enabled = rawDesired.Enabled
	}
	if dcl.StringCanonicalize(rawDesired.ClientSecret, rawInitial.ClientSecret) {
		canonicalDesired.ClientSecret = rawInitial.ClientSecret
	} else {
		canonicalDesired.ClientSecret = rawDesired.ClientSecret
	}
	canonicalDesired.ResponseType = canonicalizeTenantOAuthIdpConfigResponseType(rawDesired.ResponseType, rawInitial.ResponseType, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.NameToSelfLink(rawDesired.Tenant, rawInitial.Tenant) {
		canonicalDesired.Tenant = rawInitial.Tenant
	} else {
		canonicalDesired.Tenant = rawDesired.Tenant
	}
	return canonicalDesired, nil
}

func canonicalizeTenantOAuthIdpConfigNewState(c *Client, rawNew, rawDesired *TenantOAuthIdpConfig) (*TenantOAuthIdpConfig, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ClientId) && dcl.IsEmptyValueIndirect(rawDesired.ClientId) {
		rawNew.ClientId = rawDesired.ClientId
	} else {
		if dcl.StringCanonicalize(rawDesired.ClientId, rawNew.ClientId) {
			rawNew.ClientId = rawDesired.ClientId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Issuer) && dcl.IsEmptyValueIndirect(rawDesired.Issuer) {
		rawNew.Issuer = rawDesired.Issuer
	} else {
		if dcl.StringCanonicalize(rawDesired.Issuer, rawNew.Issuer) {
			rawNew.Issuer = rawDesired.Issuer
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Enabled) && dcl.IsEmptyValueIndirect(rawDesired.Enabled) {
		rawNew.Enabled = rawDesired.Enabled
	} else {
		if dcl.BoolCanonicalize(rawDesired.Enabled, rawNew.Enabled) {
			rawNew.Enabled = rawDesired.Enabled
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ClientSecret) && dcl.IsEmptyValueIndirect(rawDesired.ClientSecret) {
		rawNew.ClientSecret = rawDesired.ClientSecret
	} else {
		if dcl.StringCanonicalize(rawDesired.ClientSecret, rawNew.ClientSecret) {
			rawNew.ClientSecret = rawDesired.ClientSecret
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ResponseType) && dcl.IsEmptyValueIndirect(rawDesired.ResponseType) {
		rawNew.ResponseType = rawDesired.ResponseType
	} else {
		rawNew.ResponseType = canonicalizeNewTenantOAuthIdpConfigResponseType(c, rawDesired.ResponseType, rawNew.ResponseType)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Tenant = rawDesired.Tenant

	return rawNew, nil
}

func canonicalizeTenantOAuthIdpConfigResponseType(des, initial *TenantOAuthIdpConfigResponseType, opts ...dcl.ApplyOption) *TenantOAuthIdpConfigResponseType {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TenantOAuthIdpConfigResponseType{}

	if dcl.BoolCanonicalize(des.IdToken, initial.IdToken) || dcl.IsZeroValue(des.IdToken) {
		cDes.IdToken = initial.IdToken
	} else {
		cDes.IdToken = des.IdToken
	}
	if dcl.BoolCanonicalize(des.Code, initial.Code) || dcl.IsZeroValue(des.Code) {
		cDes.Code = initial.Code
	} else {
		cDes.Code = des.Code
	}
	if dcl.BoolCanonicalize(des.Token, initial.Token) || dcl.IsZeroValue(des.Token) {
		cDes.Token = initial.Token
	} else {
		cDes.Token = des.Token
	}

	return cDes
}

func canonicalizeTenantOAuthIdpConfigResponseTypeSlice(des, initial []TenantOAuthIdpConfigResponseType, opts ...dcl.ApplyOption) []TenantOAuthIdpConfigResponseType {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TenantOAuthIdpConfigResponseType, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTenantOAuthIdpConfigResponseType(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TenantOAuthIdpConfigResponseType, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTenantOAuthIdpConfigResponseType(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTenantOAuthIdpConfigResponseType(c *Client, des, nw *TenantOAuthIdpConfigResponseType) *TenantOAuthIdpConfigResponseType {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for TenantOAuthIdpConfigResponseType while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.IdToken, nw.IdToken) {
		nw.IdToken = des.IdToken
	}
	if dcl.BoolCanonicalize(des.Code, nw.Code) {
		nw.Code = des.Code
	}
	if dcl.BoolCanonicalize(des.Token, nw.Token) {
		nw.Token = des.Token
	}

	return nw
}

func canonicalizeNewTenantOAuthIdpConfigResponseTypeSet(c *Client, des, nw []TenantOAuthIdpConfigResponseType) []TenantOAuthIdpConfigResponseType {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []TenantOAuthIdpConfigResponseType
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareTenantOAuthIdpConfigResponseTypeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewTenantOAuthIdpConfigResponseType(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewTenantOAuthIdpConfigResponseTypeSlice(c *Client, des, nw []TenantOAuthIdpConfigResponseType) []TenantOAuthIdpConfigResponseType {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TenantOAuthIdpConfigResponseType
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTenantOAuthIdpConfigResponseType(c, &d, &n))
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
func diffTenantOAuthIdpConfig(c *Client, desired, actual *TenantOAuthIdpConfig, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTenantOAuthIdpConfigUpdateConfigOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClientId, actual.ClientId, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTenantOAuthIdpConfigUpdateConfigOperation")}, fn.AddNest("ClientId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Issuer, actual.Issuer, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTenantOAuthIdpConfigUpdateConfigOperation")}, fn.AddNest("Issuer")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTenantOAuthIdpConfigUpdateConfigOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTenantOAuthIdpConfigUpdateConfigOperation")}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClientSecret, actual.ClientSecret, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTenantOAuthIdpConfigUpdateConfigOperation")}, fn.AddNest("ClientSecret")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResponseType, actual.ResponseType, dcl.DiffInfo{ServerDefault: true, ObjectFunction: compareTenantOAuthIdpConfigResponseTypeNewStyle, EmptyObject: EmptyTenantOAuthIdpConfigResponseType, OperationSelector: dcl.TriggersOperation("updateTenantOAuthIdpConfigUpdateConfigOperation")}, fn.AddNest("ResponseType")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Tenant, actual.Tenant, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Tenant")); len(ds) != 0 || err != nil {
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
func compareTenantOAuthIdpConfigResponseTypeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TenantOAuthIdpConfigResponseType)
	if !ok {
		desiredNotPointer, ok := d.(TenantOAuthIdpConfigResponseType)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TenantOAuthIdpConfigResponseType or *TenantOAuthIdpConfigResponseType", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TenantOAuthIdpConfigResponseType)
	if !ok {
		actualNotPointer, ok := a.(TenantOAuthIdpConfigResponseType)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TenantOAuthIdpConfigResponseType", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.IdToken, actual.IdToken, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IdToken")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Code, actual.Code, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Code")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Token, actual.Token, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Token")); len(ds) != 0 || err != nil {
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
func (r *TenantOAuthIdpConfig) urlNormalized() *TenantOAuthIdpConfig {
	normalized := dcl.Copy(*r).(TenantOAuthIdpConfig)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.ClientId = dcl.SelfLinkToName(r.ClientId)
	normalized.Issuer = dcl.SelfLinkToName(r.Issuer)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.ClientSecret = dcl.SelfLinkToName(r.ClientSecret)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Tenant = dcl.SelfLinkToName(r.Tenant)
	return &normalized
}

func (r *TenantOAuthIdpConfig) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateConfig" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(nr.Project),
			"tenant":  dcl.ValueOrEmptyString(nr.Tenant),
			"name":    dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/tenants/{{tenant}}/oauthIdpConfigs/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the TenantOAuthIdpConfig resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *TenantOAuthIdpConfig) marshal(c *Client) ([]byte, error) {
	m, err := expandTenantOAuthIdpConfig(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling TenantOAuthIdpConfig: %w", err)
	}
	dcl.MoveMapEntry(
		m,
		[]string{"name"},
		[]string{},
	)

	return json.Marshal(m)
}

// unmarshalTenantOAuthIdpConfig decodes JSON responses into the TenantOAuthIdpConfig resource schema.
func unmarshalTenantOAuthIdpConfig(b []byte, c *Client, res *TenantOAuthIdpConfig) (*TenantOAuthIdpConfig, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapTenantOAuthIdpConfig(m, c, res)
}

func unmarshalMapTenantOAuthIdpConfig(m map[string]interface{}, c *Client, res *TenantOAuthIdpConfig) (*TenantOAuthIdpConfig, error) {

	flattened := flattenTenantOAuthIdpConfig(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandTenantOAuthIdpConfig expands TenantOAuthIdpConfig into a JSON request object.
func expandTenantOAuthIdpConfig(c *Client, f *TenantOAuthIdpConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v := f.Name; dcl.ValueShouldBeSent(v) {
		m["name"] = v
	}
	if v := f.ClientId; dcl.ValueShouldBeSent(v) {
		m["clientId"] = v
	}
	if v := f.Issuer; dcl.ValueShouldBeSent(v) {
		m["issuer"] = v
	}
	if v := f.DisplayName; dcl.ValueShouldBeSent(v) {
		m["displayName"] = v
	}
	if v := f.Enabled; dcl.ValueShouldBeSent(v) {
		m["enabled"] = v
	}
	if v := f.ClientSecret; dcl.ValueShouldBeSent(v) {
		m["clientSecret"] = v
	}
	if v, err := expandTenantOAuthIdpConfigResponseType(c, f.ResponseType, res); err != nil {
		return nil, fmt.Errorf("error expanding ResponseType into responseType: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["responseType"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Tenant into tenant: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["tenant"] = v
	}

	return m, nil
}

// flattenTenantOAuthIdpConfig flattens TenantOAuthIdpConfig from a JSON request object into the
// TenantOAuthIdpConfig type.
func flattenTenantOAuthIdpConfig(c *Client, i interface{}, res *TenantOAuthIdpConfig) *TenantOAuthIdpConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &TenantOAuthIdpConfig{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.ClientId = dcl.FlattenString(m["clientId"])
	resultRes.Issuer = dcl.FlattenString(m["issuer"])
	resultRes.DisplayName = dcl.FlattenString(m["displayName"])
	resultRes.Enabled = dcl.FlattenBool(m["enabled"])
	resultRes.ClientSecret = dcl.FlattenString(m["clientSecret"])
	resultRes.ResponseType = flattenTenantOAuthIdpConfigResponseType(c, m["responseType"], res)
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Tenant = dcl.FlattenString(m["tenant"])

	return resultRes
}

// expandTenantOAuthIdpConfigResponseTypeMap expands the contents of TenantOAuthIdpConfigResponseType into a JSON
// request object.
func expandTenantOAuthIdpConfigResponseTypeMap(c *Client, f map[string]TenantOAuthIdpConfigResponseType, res *TenantOAuthIdpConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTenantOAuthIdpConfigResponseType(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTenantOAuthIdpConfigResponseTypeSlice expands the contents of TenantOAuthIdpConfigResponseType into a JSON
// request object.
func expandTenantOAuthIdpConfigResponseTypeSlice(c *Client, f []TenantOAuthIdpConfigResponseType, res *TenantOAuthIdpConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTenantOAuthIdpConfigResponseType(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTenantOAuthIdpConfigResponseTypeMap flattens the contents of TenantOAuthIdpConfigResponseType from a JSON
// response object.
func flattenTenantOAuthIdpConfigResponseTypeMap(c *Client, i interface{}, res *TenantOAuthIdpConfig) map[string]TenantOAuthIdpConfigResponseType {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TenantOAuthIdpConfigResponseType{}
	}

	if len(a) == 0 {
		return map[string]TenantOAuthIdpConfigResponseType{}
	}

	items := make(map[string]TenantOAuthIdpConfigResponseType)
	for k, item := range a {
		items[k] = *flattenTenantOAuthIdpConfigResponseType(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTenantOAuthIdpConfigResponseTypeSlice flattens the contents of TenantOAuthIdpConfigResponseType from a JSON
// response object.
func flattenTenantOAuthIdpConfigResponseTypeSlice(c *Client, i interface{}, res *TenantOAuthIdpConfig) []TenantOAuthIdpConfigResponseType {
	a, ok := i.([]interface{})
	if !ok {
		return []TenantOAuthIdpConfigResponseType{}
	}

	if len(a) == 0 {
		return []TenantOAuthIdpConfigResponseType{}
	}

	items := make([]TenantOAuthIdpConfigResponseType, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTenantOAuthIdpConfigResponseType(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTenantOAuthIdpConfigResponseType expands an instance of TenantOAuthIdpConfigResponseType into a JSON
// request object.
func expandTenantOAuthIdpConfigResponseType(c *Client, f *TenantOAuthIdpConfigResponseType, res *TenantOAuthIdpConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.IdToken; !dcl.IsEmptyValueIndirect(v) {
		m["idToken"] = v
	}
	if v := f.Code; !dcl.IsEmptyValueIndirect(v) {
		m["code"] = v
	}
	if v := f.Token; !dcl.IsEmptyValueIndirect(v) {
		m["token"] = v
	}

	return m, nil
}

// flattenTenantOAuthIdpConfigResponseType flattens an instance of TenantOAuthIdpConfigResponseType from a JSON
// response object.
func flattenTenantOAuthIdpConfigResponseType(c *Client, i interface{}, res *TenantOAuthIdpConfig) *TenantOAuthIdpConfigResponseType {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TenantOAuthIdpConfigResponseType{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTenantOAuthIdpConfigResponseType
	}
	r.IdToken = dcl.FlattenBool(m["idToken"])
	r.Code = dcl.FlattenBool(m["code"])
	r.Token = dcl.FlattenBool(m["token"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *TenantOAuthIdpConfig) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalTenantOAuthIdpConfig(b, c, r)
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
		if nr.Tenant == nil && ncr.Tenant == nil {
			c.Config.Logger.Info("Both Tenant fields null - considering equal.")
		} else if nr.Tenant == nil || ncr.Tenant == nil {
			c.Config.Logger.Info("Only one Tenant field is null - considering unequal.")
			return false
		} else if *nr.Tenant != *ncr.Tenant {
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

type tenantOAuthIdpConfigDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         tenantOAuthIdpConfigApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToTenantOAuthIdpConfigDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]tenantOAuthIdpConfigDiff, error) {
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
	var diffs []tenantOAuthIdpConfigDiff
	// For each operation name, create a tenantOAuthIdpConfigDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := tenantOAuthIdpConfigDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToTenantOAuthIdpConfigApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToTenantOAuthIdpConfigApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (tenantOAuthIdpConfigApiOperation, error) {
	switch opName {

	case "updateTenantOAuthIdpConfigUpdateConfigOperation":
		return &updateTenantOAuthIdpConfigUpdateConfigOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractTenantOAuthIdpConfigFields(r *TenantOAuthIdpConfig) error {
	vResponseType := r.ResponseType
	if vResponseType == nil {
		// note: explicitly not the empty object.
		vResponseType = &TenantOAuthIdpConfigResponseType{}
	}
	if err := extractTenantOAuthIdpConfigResponseTypeFields(r, vResponseType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vResponseType) {
		r.ResponseType = vResponseType
	}
	return nil
}
func extractTenantOAuthIdpConfigResponseTypeFields(r *TenantOAuthIdpConfig, o *TenantOAuthIdpConfigResponseType) error {
	return nil
}

func postReadExtractTenantOAuthIdpConfigFields(r *TenantOAuthIdpConfig) error {
	vResponseType := r.ResponseType
	if vResponseType == nil {
		// note: explicitly not the empty object.
		vResponseType = &TenantOAuthIdpConfigResponseType{}
	}
	if err := postReadExtractTenantOAuthIdpConfigResponseTypeFields(r, vResponseType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vResponseType) {
		r.ResponseType = vResponseType
	}
	return nil
}
func postReadExtractTenantOAuthIdpConfigResponseTypeFields(r *TenantOAuthIdpConfig, o *TenantOAuthIdpConfigResponseType) error {
	return nil
}
