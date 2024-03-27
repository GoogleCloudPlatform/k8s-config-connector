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
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *TlsRoute) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "rules"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	return nil
}
func (r *TlsRouteRules) validate() error {
	if err := dcl.Required(r, "matches"); err != nil {
		return err
	}
	if err := dcl.Required(r, "action"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Action) {
		if err := r.Action.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *TlsRouteRulesMatches) validate() error {
	return nil
}
func (r *TlsRouteRulesAction) validate() error {
	if err := dcl.Required(r, "destinations"); err != nil {
		return err
	}
	return nil
}
func (r *TlsRouteRulesActionDestinations) validate() error {
	if err := dcl.Required(r, "serviceName"); err != nil {
		return err
	}
	return nil
}
func (r *TlsRoute) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://networkservices.googleapis.com/v1beta1/", params)
}

func (r *TlsRoute) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/tlsRoutes/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *TlsRoute) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/tlsRoutes", nr.basePath(), userBasePath, params), nil

}

func (r *TlsRoute) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/tlsRoutes?tlsRouteId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *TlsRoute) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/tlsRoutes/{{name}}", nr.basePath(), userBasePath, params), nil
}

// tlsRouteApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type tlsRouteApiOperation interface {
	do(context.Context, *TlsRoute, *Client) error
}

// newUpdateTlsRouteUpdateTlsRouteRequest creates a request for an
// TlsRoute resource's UpdateTlsRoute update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateTlsRouteUpdateTlsRouteRequest(ctx context.Context, f *TlsRoute, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v, err := dcl.DeriveField("projects/%s/locations/global/tlsRoutes/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v, err := expandTlsRouteRulesSlice(c, f.Rules, res); err != nil {
		return nil, fmt.Errorf("error expanding Rules into rules: %w", err)
	} else if v != nil {
		req["rules"] = v
	}
	if v := f.Meshes; v != nil {
		req["meshes"] = v
	}
	if v := f.Gateways; v != nil {
		req["gateways"] = v
	}
	req["name"] = fmt.Sprintf("projects/%s/locations/%s/tlsRoutes/%s", *f.Project, *f.Location, *f.Name)

	return req, nil
}

// marshalUpdateTlsRouteUpdateTlsRouteRequest converts the update into
// the final JSON request body.
func marshalUpdateTlsRouteUpdateTlsRouteRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateTlsRouteUpdateTlsRouteOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateTlsRouteUpdateTlsRouteOperation) do(ctx context.Context, r *TlsRoute, c *Client) error {
	_, err := c.GetTlsRoute(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateTlsRoute")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateTlsRouteUpdateTlsRouteRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateTlsRouteUpdateTlsRouteRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
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

func (c *Client) listTlsRouteRaw(ctx context.Context, r *TlsRoute, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != TlsRouteMaxPage {
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

type listTlsRouteOperation struct {
	TlsRoutes []map[string]interface{} `json:"tlsRoutes"`
	Token     string                   `json:"nextPageToken"`
}

func (c *Client) listTlsRoute(ctx context.Context, r *TlsRoute, pageToken string, pageSize int32) ([]*TlsRoute, string, error) {
	b, err := c.listTlsRouteRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listTlsRouteOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*TlsRoute
	for _, v := range m.TlsRoutes {
		res, err := unmarshalMapTlsRoute(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllTlsRoute(ctx context.Context, f func(*TlsRoute) bool, resources []*TlsRoute) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteTlsRoute(ctx, res)
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

type deleteTlsRouteOperation struct{}

func (op *deleteTlsRouteOperation) do(ctx context.Context, r *TlsRoute, c *Client) error {
	r, err := c.GetTlsRoute(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "TlsRoute not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetTlsRoute checking for existence. error: %v", err)
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
		_, err := c.GetTlsRoute(ctx, r)
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
type createTlsRouteOperation struct {
	response map[string]interface{}
}

func (op *createTlsRouteOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createTlsRouteOperation) do(ctx context.Context, r *TlsRoute, c *Client) error {
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

	if _, err := c.GetTlsRoute(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getTlsRouteRaw(ctx context.Context, r *TlsRoute) ([]byte, error) {

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

func (c *Client) tlsRouteDiffsForRawDesired(ctx context.Context, rawDesired *TlsRoute, opts ...dcl.ApplyOption) (initial, desired *TlsRoute, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *TlsRoute
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*TlsRoute); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected TlsRoute, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetTlsRoute(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a TlsRoute resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve TlsRoute resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that TlsRoute resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeTlsRouteDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for TlsRoute: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for TlsRoute: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractTlsRouteFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeTlsRouteInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for TlsRoute: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeTlsRouteDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for TlsRoute: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffTlsRoute(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeTlsRouteInitialState(rawInitial, rawDesired *TlsRoute) (*TlsRoute, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeTlsRouteDesiredState(rawDesired, rawInitial *TlsRoute, opts ...dcl.ApplyOption) (*TlsRoute, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &TlsRoute{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	canonicalDesired.Rules = canonicalizeTlsRouteRulesSlice(rawDesired.Rules, rawInitial.Rules, opts...)
	if dcl.StringArrayCanonicalize(rawDesired.Meshes, rawInitial.Meshes) {
		canonicalDesired.Meshes = rawInitial.Meshes
	} else {
		canonicalDesired.Meshes = rawDesired.Meshes
	}
	if dcl.StringArrayCanonicalize(rawDesired.Gateways, rawInitial.Gateways) {
		canonicalDesired.Gateways = rawInitial.Gateways
	} else {
		canonicalDesired.Gateways = rawDesired.Gateways
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}
	return canonicalDesired, nil
}

func canonicalizeTlsRouteNewState(c *Client, rawNew, rawDesired *TlsRoute) (*TlsRoute, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Rules) && dcl.IsEmptyValueIndirect(rawDesired.Rules) {
		rawNew.Rules = rawDesired.Rules
	} else {
		rawNew.Rules = canonicalizeNewTlsRouteRulesSlice(c, rawDesired.Rules, rawNew.Rules)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Meshes) && dcl.IsEmptyValueIndirect(rawDesired.Meshes) {
		rawNew.Meshes = rawDesired.Meshes
	} else {
		if dcl.StringArrayCanonicalize(rawDesired.Meshes, rawNew.Meshes) {
			rawNew.Meshes = rawDesired.Meshes
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Gateways) && dcl.IsEmptyValueIndirect(rawDesired.Gateways) {
		rawNew.Gateways = rawDesired.Gateways
	} else {
		if dcl.StringArrayCanonicalize(rawDesired.Gateways, rawNew.Gateways) {
			rawNew.Gateways = rawDesired.Gateways
		}
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeTlsRouteRules(des, initial *TlsRouteRules, opts ...dcl.ApplyOption) *TlsRouteRules {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TlsRouteRules{}

	cDes.Matches = canonicalizeTlsRouteRulesMatchesSlice(des.Matches, initial.Matches, opts...)
	cDes.Action = canonicalizeTlsRouteRulesAction(des.Action, initial.Action, opts...)

	return cDes
}

func canonicalizeTlsRouteRulesSlice(des, initial []TlsRouteRules, opts ...dcl.ApplyOption) []TlsRouteRules {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TlsRouteRules, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTlsRouteRules(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TlsRouteRules, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTlsRouteRules(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTlsRouteRules(c *Client, des, nw *TlsRouteRules) *TlsRouteRules {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for TlsRouteRules while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Matches = canonicalizeNewTlsRouteRulesMatchesSlice(c, des.Matches, nw.Matches)
	nw.Action = canonicalizeNewTlsRouteRulesAction(c, des.Action, nw.Action)

	return nw
}

func canonicalizeNewTlsRouteRulesSet(c *Client, des, nw []TlsRouteRules) []TlsRouteRules {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []TlsRouteRules
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareTlsRouteRulesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewTlsRouteRules(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewTlsRouteRulesSlice(c *Client, des, nw []TlsRouteRules) []TlsRouteRules {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TlsRouteRules
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTlsRouteRules(c, &d, &n))
	}

	return items
}

func canonicalizeTlsRouteRulesMatches(des, initial *TlsRouteRulesMatches, opts ...dcl.ApplyOption) *TlsRouteRulesMatches {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TlsRouteRulesMatches{}

	if dcl.StringArrayCanonicalize(des.SniHost, initial.SniHost) {
		cDes.SniHost = initial.SniHost
	} else {
		cDes.SniHost = des.SniHost
	}
	if dcl.StringArrayCanonicalize(des.Alpn, initial.Alpn) {
		cDes.Alpn = initial.Alpn
	} else {
		cDes.Alpn = des.Alpn
	}

	return cDes
}

func canonicalizeTlsRouteRulesMatchesSlice(des, initial []TlsRouteRulesMatches, opts ...dcl.ApplyOption) []TlsRouteRulesMatches {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TlsRouteRulesMatches, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTlsRouteRulesMatches(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TlsRouteRulesMatches, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTlsRouteRulesMatches(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTlsRouteRulesMatches(c *Client, des, nw *TlsRouteRulesMatches) *TlsRouteRulesMatches {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for TlsRouteRulesMatches while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.SniHost, nw.SniHost) {
		nw.SniHost = des.SniHost
	}
	if dcl.StringArrayCanonicalize(des.Alpn, nw.Alpn) {
		nw.Alpn = des.Alpn
	}

	return nw
}

func canonicalizeNewTlsRouteRulesMatchesSet(c *Client, des, nw []TlsRouteRulesMatches) []TlsRouteRulesMatches {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []TlsRouteRulesMatches
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareTlsRouteRulesMatchesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewTlsRouteRulesMatches(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewTlsRouteRulesMatchesSlice(c *Client, des, nw []TlsRouteRulesMatches) []TlsRouteRulesMatches {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TlsRouteRulesMatches
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTlsRouteRulesMatches(c, &d, &n))
	}

	return items
}

func canonicalizeTlsRouteRulesAction(des, initial *TlsRouteRulesAction, opts ...dcl.ApplyOption) *TlsRouteRulesAction {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TlsRouteRulesAction{}

	cDes.Destinations = canonicalizeTlsRouteRulesActionDestinationsSlice(des.Destinations, initial.Destinations, opts...)

	return cDes
}

func canonicalizeTlsRouteRulesActionSlice(des, initial []TlsRouteRulesAction, opts ...dcl.ApplyOption) []TlsRouteRulesAction {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TlsRouteRulesAction, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTlsRouteRulesAction(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TlsRouteRulesAction, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTlsRouteRulesAction(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTlsRouteRulesAction(c *Client, des, nw *TlsRouteRulesAction) *TlsRouteRulesAction {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for TlsRouteRulesAction while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Destinations = canonicalizeNewTlsRouteRulesActionDestinationsSlice(c, des.Destinations, nw.Destinations)

	return nw
}

func canonicalizeNewTlsRouteRulesActionSet(c *Client, des, nw []TlsRouteRulesAction) []TlsRouteRulesAction {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []TlsRouteRulesAction
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareTlsRouteRulesActionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewTlsRouteRulesAction(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewTlsRouteRulesActionSlice(c *Client, des, nw []TlsRouteRulesAction) []TlsRouteRulesAction {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TlsRouteRulesAction
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTlsRouteRulesAction(c, &d, &n))
	}

	return items
}

func canonicalizeTlsRouteRulesActionDestinations(des, initial *TlsRouteRulesActionDestinations, opts ...dcl.ApplyOption) *TlsRouteRulesActionDestinations {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TlsRouteRulesActionDestinations{}

	if dcl.IsZeroValue(des.ServiceName) || (dcl.IsEmptyValueIndirect(des.ServiceName) && dcl.IsEmptyValueIndirect(initial.ServiceName)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ServiceName = initial.ServiceName
	} else {
		cDes.ServiceName = des.ServiceName
	}
	if dcl.IsZeroValue(des.Weight) || (dcl.IsEmptyValueIndirect(des.Weight) && dcl.IsEmptyValueIndirect(initial.Weight)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Weight = initial.Weight
	} else {
		cDes.Weight = des.Weight
	}

	return cDes
}

func canonicalizeTlsRouteRulesActionDestinationsSlice(des, initial []TlsRouteRulesActionDestinations, opts ...dcl.ApplyOption) []TlsRouteRulesActionDestinations {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TlsRouteRulesActionDestinations, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTlsRouteRulesActionDestinations(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TlsRouteRulesActionDestinations, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTlsRouteRulesActionDestinations(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTlsRouteRulesActionDestinations(c *Client, des, nw *TlsRouteRulesActionDestinations) *TlsRouteRulesActionDestinations {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for TlsRouteRulesActionDestinations while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewTlsRouteRulesActionDestinationsSet(c *Client, des, nw []TlsRouteRulesActionDestinations) []TlsRouteRulesActionDestinations {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []TlsRouteRulesActionDestinations
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareTlsRouteRulesActionDestinationsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewTlsRouteRulesActionDestinations(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewTlsRouteRulesActionDestinationsSlice(c *Client, des, nw []TlsRouteRulesActionDestinations) []TlsRouteRulesActionDestinations {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TlsRouteRulesActionDestinations
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTlsRouteRulesActionDestinations(c, &d, &n))
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
func diffTlsRoute(c *Client, desired, actual *TlsRoute, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTlsRouteUpdateTlsRouteOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SelfLink")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTlsRouteUpdateTlsRouteOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Rules, actual.Rules, dcl.DiffInfo{ObjectFunction: compareTlsRouteRulesNewStyle, EmptyObject: EmptyTlsRouteRules, OperationSelector: dcl.TriggersOperation("updateTlsRouteUpdateTlsRouteOperation")}, fn.AddNest("Rules")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Meshes, actual.Meshes, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateTlsRouteUpdateTlsRouteOperation")}, fn.AddNest("Meshes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Gateways, actual.Gateways, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateTlsRouteUpdateTlsRouteOperation")}, fn.AddNest("Gateways")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
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
func compareTlsRouteRulesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TlsRouteRules)
	if !ok {
		desiredNotPointer, ok := d.(TlsRouteRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TlsRouteRules or *TlsRouteRules", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TlsRouteRules)
	if !ok {
		actualNotPointer, ok := a.(TlsRouteRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TlsRouteRules", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Matches, actual.Matches, dcl.DiffInfo{ObjectFunction: compareTlsRouteRulesMatchesNewStyle, EmptyObject: EmptyTlsRouteRulesMatches, OperationSelector: dcl.TriggersOperation("updateTlsRouteUpdateTlsRouteOperation")}, fn.AddNest("Matches")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Action, actual.Action, dcl.DiffInfo{ObjectFunction: compareTlsRouteRulesActionNewStyle, EmptyObject: EmptyTlsRouteRulesAction, OperationSelector: dcl.TriggersOperation("updateTlsRouteUpdateTlsRouteOperation")}, fn.AddNest("Action")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTlsRouteRulesMatchesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TlsRouteRulesMatches)
	if !ok {
		desiredNotPointer, ok := d.(TlsRouteRulesMatches)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TlsRouteRulesMatches or *TlsRouteRulesMatches", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TlsRouteRulesMatches)
	if !ok {
		actualNotPointer, ok := a.(TlsRouteRulesMatches)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TlsRouteRulesMatches", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SniHost, actual.SniHost, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTlsRouteUpdateTlsRouteOperation")}, fn.AddNest("SniHost")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Alpn, actual.Alpn, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTlsRouteUpdateTlsRouteOperation")}, fn.AddNest("Alpn")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTlsRouteRulesActionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TlsRouteRulesAction)
	if !ok {
		desiredNotPointer, ok := d.(TlsRouteRulesAction)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TlsRouteRulesAction or *TlsRouteRulesAction", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TlsRouteRulesAction)
	if !ok {
		actualNotPointer, ok := a.(TlsRouteRulesAction)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TlsRouteRulesAction", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Destinations, actual.Destinations, dcl.DiffInfo{ObjectFunction: compareTlsRouteRulesActionDestinationsNewStyle, EmptyObject: EmptyTlsRouteRulesActionDestinations, OperationSelector: dcl.TriggersOperation("updateTlsRouteUpdateTlsRouteOperation")}, fn.AddNest("Destinations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTlsRouteRulesActionDestinationsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TlsRouteRulesActionDestinations)
	if !ok {
		desiredNotPointer, ok := d.(TlsRouteRulesActionDestinations)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TlsRouteRulesActionDestinations or *TlsRouteRulesActionDestinations", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TlsRouteRulesActionDestinations)
	if !ok {
		actualNotPointer, ok := a.(TlsRouteRulesActionDestinations)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TlsRouteRulesActionDestinations", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ServiceName, actual.ServiceName, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateTlsRouteUpdateTlsRouteOperation")}, fn.AddNest("ServiceName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Weight, actual.Weight, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTlsRouteUpdateTlsRouteOperation")}, fn.AddNest("Weight")); len(ds) != 0 || err != nil {
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
func (r *TlsRoute) urlNormalized() *TlsRoute {
	normalized := dcl.Copy(*r).(TlsRoute)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *TlsRoute) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateTlsRoute" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/tlsRoutes/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the TlsRoute resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *TlsRoute) marshal(c *Client) ([]byte, error) {
	m, err := expandTlsRoute(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling TlsRoute: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalTlsRoute decodes JSON responses into the TlsRoute resource schema.
func unmarshalTlsRoute(b []byte, c *Client, res *TlsRoute) (*TlsRoute, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapTlsRoute(m, c, res)
}

func unmarshalMapTlsRoute(m map[string]interface{}, c *Client, res *TlsRoute) (*TlsRoute, error) {

	flattened := flattenTlsRoute(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandTlsRoute expands TlsRoute into a JSON request object.
func expandTlsRoute(c *Client, f *TlsRoute) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("projects/%s/locations/global/tlsRoutes/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v, err := expandTlsRouteRulesSlice(c, f.Rules, res); err != nil {
		return nil, fmt.Errorf("error expanding Rules into rules: %w", err)
	} else if v != nil {
		m["rules"] = v
	}
	if v := f.Meshes; v != nil {
		m["meshes"] = v
	}
	if v := f.Gateways; v != nil {
		m["gateways"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}

	return m, nil
}

// flattenTlsRoute flattens TlsRoute from a JSON request object into the
// TlsRoute type.
func flattenTlsRoute(c *Client, i interface{}, res *TlsRoute) *TlsRoute {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &TlsRoute{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.SelfLink = dcl.FlattenString(m["selfLink"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.Rules = flattenTlsRouteRulesSlice(c, m["rules"], res)
	resultRes.Meshes = dcl.FlattenStringSlice(m["meshes"])
	resultRes.Gateways = dcl.FlattenStringSlice(m["gateways"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// expandTlsRouteRulesMap expands the contents of TlsRouteRules into a JSON
// request object.
func expandTlsRouteRulesMap(c *Client, f map[string]TlsRouteRules, res *TlsRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTlsRouteRules(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTlsRouteRulesSlice expands the contents of TlsRouteRules into a JSON
// request object.
func expandTlsRouteRulesSlice(c *Client, f []TlsRouteRules, res *TlsRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTlsRouteRules(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTlsRouteRulesMap flattens the contents of TlsRouteRules from a JSON
// response object.
func flattenTlsRouteRulesMap(c *Client, i interface{}, res *TlsRoute) map[string]TlsRouteRules {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TlsRouteRules{}
	}

	if len(a) == 0 {
		return map[string]TlsRouteRules{}
	}

	items := make(map[string]TlsRouteRules)
	for k, item := range a {
		items[k] = *flattenTlsRouteRules(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTlsRouteRulesSlice flattens the contents of TlsRouteRules from a JSON
// response object.
func flattenTlsRouteRulesSlice(c *Client, i interface{}, res *TlsRoute) []TlsRouteRules {
	a, ok := i.([]interface{})
	if !ok {
		return []TlsRouteRules{}
	}

	if len(a) == 0 {
		return []TlsRouteRules{}
	}

	items := make([]TlsRouteRules, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTlsRouteRules(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTlsRouteRules expands an instance of TlsRouteRules into a JSON
// request object.
func expandTlsRouteRules(c *Client, f *TlsRouteRules, res *TlsRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandTlsRouteRulesMatchesSlice(c, f.Matches, res); err != nil {
		return nil, fmt.Errorf("error expanding Matches into matches: %w", err)
	} else if v != nil {
		m["matches"] = v
	}
	if v, err := expandTlsRouteRulesAction(c, f.Action, res); err != nil {
		return nil, fmt.Errorf("error expanding Action into action: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["action"] = v
	}

	return m, nil
}

// flattenTlsRouteRules flattens an instance of TlsRouteRules from a JSON
// response object.
func flattenTlsRouteRules(c *Client, i interface{}, res *TlsRoute) *TlsRouteRules {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TlsRouteRules{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTlsRouteRules
	}
	r.Matches = flattenTlsRouteRulesMatchesSlice(c, m["matches"], res)
	r.Action = flattenTlsRouteRulesAction(c, m["action"], res)

	return r
}

// expandTlsRouteRulesMatchesMap expands the contents of TlsRouteRulesMatches into a JSON
// request object.
func expandTlsRouteRulesMatchesMap(c *Client, f map[string]TlsRouteRulesMatches, res *TlsRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTlsRouteRulesMatches(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTlsRouteRulesMatchesSlice expands the contents of TlsRouteRulesMatches into a JSON
// request object.
func expandTlsRouteRulesMatchesSlice(c *Client, f []TlsRouteRulesMatches, res *TlsRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTlsRouteRulesMatches(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTlsRouteRulesMatchesMap flattens the contents of TlsRouteRulesMatches from a JSON
// response object.
func flattenTlsRouteRulesMatchesMap(c *Client, i interface{}, res *TlsRoute) map[string]TlsRouteRulesMatches {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TlsRouteRulesMatches{}
	}

	if len(a) == 0 {
		return map[string]TlsRouteRulesMatches{}
	}

	items := make(map[string]TlsRouteRulesMatches)
	for k, item := range a {
		items[k] = *flattenTlsRouteRulesMatches(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTlsRouteRulesMatchesSlice flattens the contents of TlsRouteRulesMatches from a JSON
// response object.
func flattenTlsRouteRulesMatchesSlice(c *Client, i interface{}, res *TlsRoute) []TlsRouteRulesMatches {
	a, ok := i.([]interface{})
	if !ok {
		return []TlsRouteRulesMatches{}
	}

	if len(a) == 0 {
		return []TlsRouteRulesMatches{}
	}

	items := make([]TlsRouteRulesMatches, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTlsRouteRulesMatches(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTlsRouteRulesMatches expands an instance of TlsRouteRulesMatches into a JSON
// request object.
func expandTlsRouteRulesMatches(c *Client, f *TlsRouteRulesMatches, res *TlsRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SniHost; v != nil {
		m["sniHost"] = v
	}
	if v := f.Alpn; v != nil {
		m["alpn"] = v
	}

	return m, nil
}

// flattenTlsRouteRulesMatches flattens an instance of TlsRouteRulesMatches from a JSON
// response object.
func flattenTlsRouteRulesMatches(c *Client, i interface{}, res *TlsRoute) *TlsRouteRulesMatches {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TlsRouteRulesMatches{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTlsRouteRulesMatches
	}
	r.SniHost = dcl.FlattenStringSlice(m["sniHost"])
	r.Alpn = dcl.FlattenStringSlice(m["alpn"])

	return r
}

// expandTlsRouteRulesActionMap expands the contents of TlsRouteRulesAction into a JSON
// request object.
func expandTlsRouteRulesActionMap(c *Client, f map[string]TlsRouteRulesAction, res *TlsRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTlsRouteRulesAction(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTlsRouteRulesActionSlice expands the contents of TlsRouteRulesAction into a JSON
// request object.
func expandTlsRouteRulesActionSlice(c *Client, f []TlsRouteRulesAction, res *TlsRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTlsRouteRulesAction(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTlsRouteRulesActionMap flattens the contents of TlsRouteRulesAction from a JSON
// response object.
func flattenTlsRouteRulesActionMap(c *Client, i interface{}, res *TlsRoute) map[string]TlsRouteRulesAction {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TlsRouteRulesAction{}
	}

	if len(a) == 0 {
		return map[string]TlsRouteRulesAction{}
	}

	items := make(map[string]TlsRouteRulesAction)
	for k, item := range a {
		items[k] = *flattenTlsRouteRulesAction(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTlsRouteRulesActionSlice flattens the contents of TlsRouteRulesAction from a JSON
// response object.
func flattenTlsRouteRulesActionSlice(c *Client, i interface{}, res *TlsRoute) []TlsRouteRulesAction {
	a, ok := i.([]interface{})
	if !ok {
		return []TlsRouteRulesAction{}
	}

	if len(a) == 0 {
		return []TlsRouteRulesAction{}
	}

	items := make([]TlsRouteRulesAction, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTlsRouteRulesAction(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTlsRouteRulesAction expands an instance of TlsRouteRulesAction into a JSON
// request object.
func expandTlsRouteRulesAction(c *Client, f *TlsRouteRulesAction, res *TlsRoute) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandTlsRouteRulesActionDestinationsSlice(c, f.Destinations, res); err != nil {
		return nil, fmt.Errorf("error expanding Destinations into destinations: %w", err)
	} else if v != nil {
		m["destinations"] = v
	}

	return m, nil
}

// flattenTlsRouteRulesAction flattens an instance of TlsRouteRulesAction from a JSON
// response object.
func flattenTlsRouteRulesAction(c *Client, i interface{}, res *TlsRoute) *TlsRouteRulesAction {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TlsRouteRulesAction{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTlsRouteRulesAction
	}
	r.Destinations = flattenTlsRouteRulesActionDestinationsSlice(c, m["destinations"], res)

	return r
}

// expandTlsRouteRulesActionDestinationsMap expands the contents of TlsRouteRulesActionDestinations into a JSON
// request object.
func expandTlsRouteRulesActionDestinationsMap(c *Client, f map[string]TlsRouteRulesActionDestinations, res *TlsRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTlsRouteRulesActionDestinations(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTlsRouteRulesActionDestinationsSlice expands the contents of TlsRouteRulesActionDestinations into a JSON
// request object.
func expandTlsRouteRulesActionDestinationsSlice(c *Client, f []TlsRouteRulesActionDestinations, res *TlsRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTlsRouteRulesActionDestinations(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTlsRouteRulesActionDestinationsMap flattens the contents of TlsRouteRulesActionDestinations from a JSON
// response object.
func flattenTlsRouteRulesActionDestinationsMap(c *Client, i interface{}, res *TlsRoute) map[string]TlsRouteRulesActionDestinations {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TlsRouteRulesActionDestinations{}
	}

	if len(a) == 0 {
		return map[string]TlsRouteRulesActionDestinations{}
	}

	items := make(map[string]TlsRouteRulesActionDestinations)
	for k, item := range a {
		items[k] = *flattenTlsRouteRulesActionDestinations(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTlsRouteRulesActionDestinationsSlice flattens the contents of TlsRouteRulesActionDestinations from a JSON
// response object.
func flattenTlsRouteRulesActionDestinationsSlice(c *Client, i interface{}, res *TlsRoute) []TlsRouteRulesActionDestinations {
	a, ok := i.([]interface{})
	if !ok {
		return []TlsRouteRulesActionDestinations{}
	}

	if len(a) == 0 {
		return []TlsRouteRulesActionDestinations{}
	}

	items := make([]TlsRouteRulesActionDestinations, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTlsRouteRulesActionDestinations(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTlsRouteRulesActionDestinations expands an instance of TlsRouteRulesActionDestinations into a JSON
// request object.
func expandTlsRouteRulesActionDestinations(c *Client, f *TlsRouteRulesActionDestinations, res *TlsRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ServiceName; !dcl.IsEmptyValueIndirect(v) {
		m["serviceName"] = v
	}
	if v := f.Weight; !dcl.IsEmptyValueIndirect(v) {
		m["weight"] = v
	}

	return m, nil
}

// flattenTlsRouteRulesActionDestinations flattens an instance of TlsRouteRulesActionDestinations from a JSON
// response object.
func flattenTlsRouteRulesActionDestinations(c *Client, i interface{}, res *TlsRoute) *TlsRouteRulesActionDestinations {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TlsRouteRulesActionDestinations{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTlsRouteRulesActionDestinations
	}
	r.ServiceName = dcl.FlattenString(m["serviceName"])
	r.Weight = dcl.FlattenInteger(m["weight"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *TlsRoute) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalTlsRoute(b, c, r)
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
		if nr.Location == nil && ncr.Location == nil {
			c.Config.Logger.Info("Both Location fields null - considering equal.")
		} else if nr.Location == nil || ncr.Location == nil {
			c.Config.Logger.Info("Only one Location field is null - considering unequal.")
			return false
		} else if *nr.Location != *ncr.Location {
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

type tlsRouteDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         tlsRouteApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToTlsRouteDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]tlsRouteDiff, error) {
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
	var diffs []tlsRouteDiff
	// For each operation name, create a tlsRouteDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := tlsRouteDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToTlsRouteApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToTlsRouteApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (tlsRouteApiOperation, error) {
	switch opName {

	case "updateTlsRouteUpdateTlsRouteOperation":
		return &updateTlsRouteUpdateTlsRouteOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractTlsRouteFields(r *TlsRoute) error {
	return nil
}
func extractTlsRouteRulesFields(r *TlsRoute, o *TlsRouteRules) error {
	vAction := o.Action
	if vAction == nil {
		// note: explicitly not the empty object.
		vAction = &TlsRouteRulesAction{}
	}
	if err := extractTlsRouteRulesActionFields(r, vAction); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAction) {
		o.Action = vAction
	}
	return nil
}
func extractTlsRouteRulesMatchesFields(r *TlsRoute, o *TlsRouteRulesMatches) error {
	return nil
}
func extractTlsRouteRulesActionFields(r *TlsRoute, o *TlsRouteRulesAction) error {
	return nil
}
func extractTlsRouteRulesActionDestinationsFields(r *TlsRoute, o *TlsRouteRulesActionDestinations) error {
	return nil
}

func postReadExtractTlsRouteFields(r *TlsRoute) error {
	return nil
}
func postReadExtractTlsRouteRulesFields(r *TlsRoute, o *TlsRouteRules) error {
	vAction := o.Action
	if vAction == nil {
		// note: explicitly not the empty object.
		vAction = &TlsRouteRulesAction{}
	}
	if err := extractTlsRouteRulesActionFields(r, vAction); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAction) {
		o.Action = vAction
	}
	return nil
}
func postReadExtractTlsRouteRulesMatchesFields(r *TlsRoute, o *TlsRouteRulesMatches) error {
	return nil
}
func postReadExtractTlsRouteRulesActionFields(r *TlsRoute, o *TlsRouteRulesAction) error {
	return nil
}
func postReadExtractTlsRouteRulesActionDestinationsFields(r *TlsRoute, o *TlsRouteRulesActionDestinations) error {
	return nil
}
