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
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *TcpRoute) validate() error {

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
func (r *TcpRouteRules) validate() error {
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
func (r *TcpRouteRulesMatches) validate() error {
	if err := dcl.Required(r, "address"); err != nil {
		return err
	}
	if err := dcl.Required(r, "port"); err != nil {
		return err
	}
	return nil
}
func (r *TcpRouteRulesAction) validate() error {
	return nil
}
func (r *TcpRouteRulesActionDestinations) validate() error {
	if err := dcl.Required(r, "serviceName"); err != nil {
		return err
	}
	return nil
}
func (r *TcpRoute) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://networkservices.googleapis.com/v1alpha1/", params)
}

func (r *TcpRoute) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/tcpRoutes/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *TcpRoute) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/tcpRoutes", nr.basePath(), userBasePath, params), nil

}

func (r *TcpRoute) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/tcpRoutes?tcpRouteId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *TcpRoute) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/tcpRoutes/{{name}}", nr.basePath(), userBasePath, params), nil
}

// tcpRouteApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type tcpRouteApiOperation interface {
	do(context.Context, *TcpRoute, *Client) error
}

// newUpdateTcpRouteUpdateTcpRouteRequest creates a request for an
// TcpRoute resource's UpdateTcpRoute update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateTcpRouteUpdateTcpRouteRequest(ctx context.Context, f *TcpRoute, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v, err := dcl.DeriveField("projects/%s/locations/global/tcpRoutes/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v, err := expandTcpRouteRulesSlice(c, f.Rules, res); err != nil {
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
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	req["name"] = fmt.Sprintf("projects/%s/locations/%s/tcpRoutes/%s", *f.Project, *f.Location, *f.Name)

	return req, nil
}

// marshalUpdateTcpRouteUpdateTcpRouteRequest converts the update into
// the final JSON request body.
func marshalUpdateTcpRouteUpdateTcpRouteRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateTcpRouteUpdateTcpRouteOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateTcpRouteUpdateTcpRouteOperation) do(ctx context.Context, r *TcpRoute, c *Client) error {
	_, err := c.GetTcpRoute(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateTcpRoute")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateTcpRouteUpdateTcpRouteRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateTcpRouteUpdateTcpRouteRequest(c, req)
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

func (c *Client) listTcpRouteRaw(ctx context.Context, r *TcpRoute, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != TcpRouteMaxPage {
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

type listTcpRouteOperation struct {
	TcpRoutes []map[string]interface{} `json:"tcpRoutes"`
	Token     string                   `json:"nextPageToken"`
}

func (c *Client) listTcpRoute(ctx context.Context, r *TcpRoute, pageToken string, pageSize int32) ([]*TcpRoute, string, error) {
	b, err := c.listTcpRouteRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listTcpRouteOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*TcpRoute
	for _, v := range m.TcpRoutes {
		res, err := unmarshalMapTcpRoute(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllTcpRoute(ctx context.Context, f func(*TcpRoute) bool, resources []*TcpRoute) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteTcpRoute(ctx, res)
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

type deleteTcpRouteOperation struct{}

func (op *deleteTcpRouteOperation) do(ctx context.Context, r *TcpRoute, c *Client) error {
	r, err := c.GetTcpRoute(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "TcpRoute not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetTcpRoute checking for existence. error: %v", err)
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
		_, err := c.GetTcpRoute(ctx, r)
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
type createTcpRouteOperation struct {
	response map[string]interface{}
}

func (op *createTcpRouteOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createTcpRouteOperation) do(ctx context.Context, r *TcpRoute, c *Client) error {
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

	if _, err := c.GetTcpRoute(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getTcpRouteRaw(ctx context.Context, r *TcpRoute) ([]byte, error) {

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

func (c *Client) tcpRouteDiffsForRawDesired(ctx context.Context, rawDesired *TcpRoute, opts ...dcl.ApplyOption) (initial, desired *TcpRoute, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *TcpRoute
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*TcpRoute); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected TcpRoute, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetTcpRoute(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a TcpRoute resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve TcpRoute resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that TcpRoute resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeTcpRouteDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for TcpRoute: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for TcpRoute: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractTcpRouteFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeTcpRouteInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for TcpRoute: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeTcpRouteDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for TcpRoute: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffTcpRoute(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeTcpRouteInitialState(rawInitial, rawDesired *TcpRoute) (*TcpRoute, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeTcpRouteDesiredState(rawDesired, rawInitial *TcpRoute, opts ...dcl.ApplyOption) (*TcpRoute, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &TcpRoute{}
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
	canonicalDesired.Rules = canonicalizeTcpRouteRulesSlice(rawDesired.Rules, rawInitial.Rules, opts...)
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
	if dcl.IsZeroValue(rawDesired.Labels) || (dcl.IsEmptyValueIndirect(rawDesired.Labels) && dcl.IsEmptyValueIndirect(rawInitial.Labels)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
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

func canonicalizeTcpRouteNewState(c *Client, rawNew, rawDesired *TcpRoute) (*TcpRoute, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
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
		rawNew.Rules = canonicalizeNewTcpRouteRulesSlice(c, rawDesired.Rules, rawNew.Rules)
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

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	return rawNew, nil
}

func canonicalizeTcpRouteRules(des, initial *TcpRouteRules, opts ...dcl.ApplyOption) *TcpRouteRules {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TcpRouteRules{}

	cDes.Matches = canonicalizeTcpRouteRulesMatchesSlice(des.Matches, initial.Matches, opts...)
	cDes.Action = canonicalizeTcpRouteRulesAction(des.Action, initial.Action, opts...)

	return cDes
}

func canonicalizeTcpRouteRulesSlice(des, initial []TcpRouteRules, opts ...dcl.ApplyOption) []TcpRouteRules {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TcpRouteRules, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTcpRouteRules(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TcpRouteRules, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTcpRouteRules(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTcpRouteRules(c *Client, des, nw *TcpRouteRules) *TcpRouteRules {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for TcpRouteRules while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Matches = canonicalizeNewTcpRouteRulesMatchesSlice(c, des.Matches, nw.Matches)
	nw.Action = canonicalizeNewTcpRouteRulesAction(c, des.Action, nw.Action)

	return nw
}

func canonicalizeNewTcpRouteRulesSet(c *Client, des, nw []TcpRouteRules) []TcpRouteRules {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []TcpRouteRules
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareTcpRouteRulesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewTcpRouteRules(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewTcpRouteRulesSlice(c *Client, des, nw []TcpRouteRules) []TcpRouteRules {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TcpRouteRules
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTcpRouteRules(c, &d, &n))
	}

	return items
}

func canonicalizeTcpRouteRulesMatches(des, initial *TcpRouteRulesMatches, opts ...dcl.ApplyOption) *TcpRouteRulesMatches {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TcpRouteRulesMatches{}

	if dcl.StringCanonicalize(des.Address, initial.Address) || dcl.IsZeroValue(des.Address) {
		cDes.Address = initial.Address
	} else {
		cDes.Address = des.Address
	}
	if dcl.StringCanonicalize(des.Port, initial.Port) || dcl.IsZeroValue(des.Port) {
		cDes.Port = initial.Port
	} else {
		cDes.Port = des.Port
	}

	return cDes
}

func canonicalizeTcpRouteRulesMatchesSlice(des, initial []TcpRouteRulesMatches, opts ...dcl.ApplyOption) []TcpRouteRulesMatches {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TcpRouteRulesMatches, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTcpRouteRulesMatches(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TcpRouteRulesMatches, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTcpRouteRulesMatches(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTcpRouteRulesMatches(c *Client, des, nw *TcpRouteRulesMatches) *TcpRouteRulesMatches {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for TcpRouteRulesMatches while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Address, nw.Address) {
		nw.Address = des.Address
	}
	if dcl.StringCanonicalize(des.Port, nw.Port) {
		nw.Port = des.Port
	}

	return nw
}

func canonicalizeNewTcpRouteRulesMatchesSet(c *Client, des, nw []TcpRouteRulesMatches) []TcpRouteRulesMatches {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []TcpRouteRulesMatches
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareTcpRouteRulesMatchesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewTcpRouteRulesMatches(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewTcpRouteRulesMatchesSlice(c *Client, des, nw []TcpRouteRulesMatches) []TcpRouteRulesMatches {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TcpRouteRulesMatches
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTcpRouteRulesMatches(c, &d, &n))
	}

	return items
}

func canonicalizeTcpRouteRulesAction(des, initial *TcpRouteRulesAction, opts ...dcl.ApplyOption) *TcpRouteRulesAction {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TcpRouteRulesAction{}

	cDes.Destinations = canonicalizeTcpRouteRulesActionDestinationsSlice(des.Destinations, initial.Destinations, opts...)
	if dcl.BoolCanonicalize(des.OriginalDestination, initial.OriginalDestination) || dcl.IsZeroValue(des.OriginalDestination) {
		cDes.OriginalDestination = initial.OriginalDestination
	} else {
		cDes.OriginalDestination = des.OriginalDestination
	}

	return cDes
}

func canonicalizeTcpRouteRulesActionSlice(des, initial []TcpRouteRulesAction, opts ...dcl.ApplyOption) []TcpRouteRulesAction {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TcpRouteRulesAction, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTcpRouteRulesAction(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TcpRouteRulesAction, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTcpRouteRulesAction(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTcpRouteRulesAction(c *Client, des, nw *TcpRouteRulesAction) *TcpRouteRulesAction {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for TcpRouteRulesAction while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Destinations = canonicalizeNewTcpRouteRulesActionDestinationsSlice(c, des.Destinations, nw.Destinations)
	if dcl.BoolCanonicalize(des.OriginalDestination, nw.OriginalDestination) {
		nw.OriginalDestination = des.OriginalDestination
	}

	return nw
}

func canonicalizeNewTcpRouteRulesActionSet(c *Client, des, nw []TcpRouteRulesAction) []TcpRouteRulesAction {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []TcpRouteRulesAction
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareTcpRouteRulesActionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewTcpRouteRulesAction(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewTcpRouteRulesActionSlice(c *Client, des, nw []TcpRouteRulesAction) []TcpRouteRulesAction {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TcpRouteRulesAction
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTcpRouteRulesAction(c, &d, &n))
	}

	return items
}

func canonicalizeTcpRouteRulesActionDestinations(des, initial *TcpRouteRulesActionDestinations, opts ...dcl.ApplyOption) *TcpRouteRulesActionDestinations {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TcpRouteRulesActionDestinations{}

	if dcl.IsZeroValue(des.Weight) || (dcl.IsEmptyValueIndirect(des.Weight) && dcl.IsEmptyValueIndirect(initial.Weight)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Weight = initial.Weight
	} else {
		cDes.Weight = des.Weight
	}
	if dcl.IsZeroValue(des.ServiceName) || (dcl.IsEmptyValueIndirect(des.ServiceName) && dcl.IsEmptyValueIndirect(initial.ServiceName)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ServiceName = initial.ServiceName
	} else {
		cDes.ServiceName = des.ServiceName
	}

	return cDes
}

func canonicalizeTcpRouteRulesActionDestinationsSlice(des, initial []TcpRouteRulesActionDestinations, opts ...dcl.ApplyOption) []TcpRouteRulesActionDestinations {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TcpRouteRulesActionDestinations, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTcpRouteRulesActionDestinations(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TcpRouteRulesActionDestinations, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTcpRouteRulesActionDestinations(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTcpRouteRulesActionDestinations(c *Client, des, nw *TcpRouteRulesActionDestinations) *TcpRouteRulesActionDestinations {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for TcpRouteRulesActionDestinations while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewTcpRouteRulesActionDestinationsSet(c *Client, des, nw []TcpRouteRulesActionDestinations) []TcpRouteRulesActionDestinations {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []TcpRouteRulesActionDestinations
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareTcpRouteRulesActionDestinationsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewTcpRouteRulesActionDestinations(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewTcpRouteRulesActionDestinationsSlice(c *Client, des, nw []TcpRouteRulesActionDestinations) []TcpRouteRulesActionDestinations {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TcpRouteRulesActionDestinations
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTcpRouteRulesActionDestinations(c, &d, &n))
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
func diffTcpRoute(c *Client, desired, actual *TcpRoute, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTcpRouteUpdateTcpRouteOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTcpRouteUpdateTcpRouteOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Rules, actual.Rules, dcl.DiffInfo{ObjectFunction: compareTcpRouteRulesNewStyle, EmptyObject: EmptyTcpRouteRules, OperationSelector: dcl.TriggersOperation("updateTcpRouteUpdateTcpRouteOperation")}, fn.AddNest("Rules")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Meshes, actual.Meshes, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateTcpRouteUpdateTcpRouteOperation")}, fn.AddNest("Meshes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Gateways, actual.Gateways, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateTcpRouteUpdateTcpRouteOperation")}, fn.AddNest("Gateways")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTcpRouteUpdateTcpRouteOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SelfLink")); len(ds) != 0 || err != nil {
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
func compareTcpRouteRulesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TcpRouteRules)
	if !ok {
		desiredNotPointer, ok := d.(TcpRouteRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TcpRouteRules or *TcpRouteRules", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TcpRouteRules)
	if !ok {
		actualNotPointer, ok := a.(TcpRouteRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TcpRouteRules", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Matches, actual.Matches, dcl.DiffInfo{ObjectFunction: compareTcpRouteRulesMatchesNewStyle, EmptyObject: EmptyTcpRouteRulesMatches, OperationSelector: dcl.TriggersOperation("updateTcpRouteUpdateTcpRouteOperation")}, fn.AddNest("Matches")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Action, actual.Action, dcl.DiffInfo{ObjectFunction: compareTcpRouteRulesActionNewStyle, EmptyObject: EmptyTcpRouteRulesAction, OperationSelector: dcl.TriggersOperation("updateTcpRouteUpdateTcpRouteOperation")}, fn.AddNest("Action")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTcpRouteRulesMatchesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TcpRouteRulesMatches)
	if !ok {
		desiredNotPointer, ok := d.(TcpRouteRulesMatches)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TcpRouteRulesMatches or *TcpRouteRulesMatches", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TcpRouteRulesMatches)
	if !ok {
		actualNotPointer, ok := a.(TcpRouteRulesMatches)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TcpRouteRulesMatches", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Address, actual.Address, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTcpRouteUpdateTcpRouteOperation")}, fn.AddNest("Address")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Port, actual.Port, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTcpRouteUpdateTcpRouteOperation")}, fn.AddNest("Port")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTcpRouteRulesActionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TcpRouteRulesAction)
	if !ok {
		desiredNotPointer, ok := d.(TcpRouteRulesAction)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TcpRouteRulesAction or *TcpRouteRulesAction", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TcpRouteRulesAction)
	if !ok {
		actualNotPointer, ok := a.(TcpRouteRulesAction)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TcpRouteRulesAction", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Destinations, actual.Destinations, dcl.DiffInfo{ObjectFunction: compareTcpRouteRulesActionDestinationsNewStyle, EmptyObject: EmptyTcpRouteRulesActionDestinations, OperationSelector: dcl.TriggersOperation("updateTcpRouteUpdateTcpRouteOperation")}, fn.AddNest("Destinations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OriginalDestination, actual.OriginalDestination, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTcpRouteUpdateTcpRouteOperation")}, fn.AddNest("OriginalDestination")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTcpRouteRulesActionDestinationsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TcpRouteRulesActionDestinations)
	if !ok {
		desiredNotPointer, ok := d.(TcpRouteRulesActionDestinations)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TcpRouteRulesActionDestinations or *TcpRouteRulesActionDestinations", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TcpRouteRulesActionDestinations)
	if !ok {
		actualNotPointer, ok := a.(TcpRouteRulesActionDestinations)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TcpRouteRulesActionDestinations", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Weight, actual.Weight, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTcpRouteUpdateTcpRouteOperation")}, fn.AddNest("Weight")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceName, actual.ServiceName, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateTcpRouteUpdateTcpRouteOperation")}, fn.AddNest("ServiceName")); len(ds) != 0 || err != nil {
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
func (r *TcpRoute) urlNormalized() *TcpRoute {
	normalized := dcl.Copy(*r).(TcpRoute)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	return &normalized
}

func (r *TcpRoute) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateTcpRoute" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/tcpRoutes/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the TcpRoute resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *TcpRoute) marshal(c *Client) ([]byte, error) {
	m, err := expandTcpRoute(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling TcpRoute: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalTcpRoute decodes JSON responses into the TcpRoute resource schema.
func unmarshalTcpRoute(b []byte, c *Client, res *TcpRoute) (*TcpRoute, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapTcpRoute(m, c, res)
}

func unmarshalMapTcpRoute(m map[string]interface{}, c *Client, res *TcpRoute) (*TcpRoute, error) {

	flattened := flattenTcpRoute(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandTcpRoute expands TcpRoute into a JSON request object.
func expandTcpRoute(c *Client, f *TcpRoute) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("projects/%s/locations/global/tcpRoutes/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v, err := expandTcpRouteRulesSlice(c, f.Rules, res); err != nil {
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
	if v := f.Labels; dcl.ValueShouldBeSent(v) {
		m["labels"] = v
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

// flattenTcpRoute flattens TcpRoute from a JSON request object into the
// TcpRoute type.
func flattenTcpRoute(c *Client, i interface{}, res *TcpRoute) *TcpRoute {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &TcpRoute{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.Rules = flattenTcpRouteRulesSlice(c, m["rules"], res)
	resultRes.Meshes = dcl.FlattenStringSlice(m["meshes"])
	resultRes.Gateways = dcl.FlattenStringSlice(m["gateways"])
	resultRes.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])
	resultRes.SelfLink = dcl.FlattenString(m["selfLink"])

	return resultRes
}

// expandTcpRouteRulesMap expands the contents of TcpRouteRules into a JSON
// request object.
func expandTcpRouteRulesMap(c *Client, f map[string]TcpRouteRules, res *TcpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTcpRouteRules(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTcpRouteRulesSlice expands the contents of TcpRouteRules into a JSON
// request object.
func expandTcpRouteRulesSlice(c *Client, f []TcpRouteRules, res *TcpRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTcpRouteRules(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTcpRouteRulesMap flattens the contents of TcpRouteRules from a JSON
// response object.
func flattenTcpRouteRulesMap(c *Client, i interface{}, res *TcpRoute) map[string]TcpRouteRules {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TcpRouteRules{}
	}

	if len(a) == 0 {
		return map[string]TcpRouteRules{}
	}

	items := make(map[string]TcpRouteRules)
	for k, item := range a {
		items[k] = *flattenTcpRouteRules(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTcpRouteRulesSlice flattens the contents of TcpRouteRules from a JSON
// response object.
func flattenTcpRouteRulesSlice(c *Client, i interface{}, res *TcpRoute) []TcpRouteRules {
	a, ok := i.([]interface{})
	if !ok {
		return []TcpRouteRules{}
	}

	if len(a) == 0 {
		return []TcpRouteRules{}
	}

	items := make([]TcpRouteRules, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTcpRouteRules(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTcpRouteRules expands an instance of TcpRouteRules into a JSON
// request object.
func expandTcpRouteRules(c *Client, f *TcpRouteRules, res *TcpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandTcpRouteRulesMatchesSlice(c, f.Matches, res); err != nil {
		return nil, fmt.Errorf("error expanding Matches into matches: %w", err)
	} else if v != nil {
		m["matches"] = v
	}
	if v, err := expandTcpRouteRulesAction(c, f.Action, res); err != nil {
		return nil, fmt.Errorf("error expanding Action into action: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["action"] = v
	}

	return m, nil
}

// flattenTcpRouteRules flattens an instance of TcpRouteRules from a JSON
// response object.
func flattenTcpRouteRules(c *Client, i interface{}, res *TcpRoute) *TcpRouteRules {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TcpRouteRules{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTcpRouteRules
	}
	r.Matches = flattenTcpRouteRulesMatchesSlice(c, m["matches"], res)
	r.Action = flattenTcpRouteRulesAction(c, m["action"], res)

	return r
}

// expandTcpRouteRulesMatchesMap expands the contents of TcpRouteRulesMatches into a JSON
// request object.
func expandTcpRouteRulesMatchesMap(c *Client, f map[string]TcpRouteRulesMatches, res *TcpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTcpRouteRulesMatches(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTcpRouteRulesMatchesSlice expands the contents of TcpRouteRulesMatches into a JSON
// request object.
func expandTcpRouteRulesMatchesSlice(c *Client, f []TcpRouteRulesMatches, res *TcpRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTcpRouteRulesMatches(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTcpRouteRulesMatchesMap flattens the contents of TcpRouteRulesMatches from a JSON
// response object.
func flattenTcpRouteRulesMatchesMap(c *Client, i interface{}, res *TcpRoute) map[string]TcpRouteRulesMatches {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TcpRouteRulesMatches{}
	}

	if len(a) == 0 {
		return map[string]TcpRouteRulesMatches{}
	}

	items := make(map[string]TcpRouteRulesMatches)
	for k, item := range a {
		items[k] = *flattenTcpRouteRulesMatches(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTcpRouteRulesMatchesSlice flattens the contents of TcpRouteRulesMatches from a JSON
// response object.
func flattenTcpRouteRulesMatchesSlice(c *Client, i interface{}, res *TcpRoute) []TcpRouteRulesMatches {
	a, ok := i.([]interface{})
	if !ok {
		return []TcpRouteRulesMatches{}
	}

	if len(a) == 0 {
		return []TcpRouteRulesMatches{}
	}

	items := make([]TcpRouteRulesMatches, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTcpRouteRulesMatches(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTcpRouteRulesMatches expands an instance of TcpRouteRulesMatches into a JSON
// request object.
func expandTcpRouteRulesMatches(c *Client, f *TcpRouteRulesMatches, res *TcpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Address; !dcl.IsEmptyValueIndirect(v) {
		m["address"] = v
	}
	if v := f.Port; !dcl.IsEmptyValueIndirect(v) {
		m["port"] = v
	}

	return m, nil
}

// flattenTcpRouteRulesMatches flattens an instance of TcpRouteRulesMatches from a JSON
// response object.
func flattenTcpRouteRulesMatches(c *Client, i interface{}, res *TcpRoute) *TcpRouteRulesMatches {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TcpRouteRulesMatches{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTcpRouteRulesMatches
	}
	r.Address = dcl.FlattenString(m["address"])
	r.Port = dcl.FlattenString(m["port"])

	return r
}

// expandTcpRouteRulesActionMap expands the contents of TcpRouteRulesAction into a JSON
// request object.
func expandTcpRouteRulesActionMap(c *Client, f map[string]TcpRouteRulesAction, res *TcpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTcpRouteRulesAction(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTcpRouteRulesActionSlice expands the contents of TcpRouteRulesAction into a JSON
// request object.
func expandTcpRouteRulesActionSlice(c *Client, f []TcpRouteRulesAction, res *TcpRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTcpRouteRulesAction(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTcpRouteRulesActionMap flattens the contents of TcpRouteRulesAction from a JSON
// response object.
func flattenTcpRouteRulesActionMap(c *Client, i interface{}, res *TcpRoute) map[string]TcpRouteRulesAction {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TcpRouteRulesAction{}
	}

	if len(a) == 0 {
		return map[string]TcpRouteRulesAction{}
	}

	items := make(map[string]TcpRouteRulesAction)
	for k, item := range a {
		items[k] = *flattenTcpRouteRulesAction(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTcpRouteRulesActionSlice flattens the contents of TcpRouteRulesAction from a JSON
// response object.
func flattenTcpRouteRulesActionSlice(c *Client, i interface{}, res *TcpRoute) []TcpRouteRulesAction {
	a, ok := i.([]interface{})
	if !ok {
		return []TcpRouteRulesAction{}
	}

	if len(a) == 0 {
		return []TcpRouteRulesAction{}
	}

	items := make([]TcpRouteRulesAction, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTcpRouteRulesAction(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTcpRouteRulesAction expands an instance of TcpRouteRulesAction into a JSON
// request object.
func expandTcpRouteRulesAction(c *Client, f *TcpRouteRulesAction, res *TcpRoute) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandTcpRouteRulesActionDestinationsSlice(c, f.Destinations, res); err != nil {
		return nil, fmt.Errorf("error expanding Destinations into destinations: %w", err)
	} else if v != nil {
		m["destinations"] = v
	}
	if v := f.OriginalDestination; !dcl.IsEmptyValueIndirect(v) {
		m["originalDestination"] = v
	}

	return m, nil
}

// flattenTcpRouteRulesAction flattens an instance of TcpRouteRulesAction from a JSON
// response object.
func flattenTcpRouteRulesAction(c *Client, i interface{}, res *TcpRoute) *TcpRouteRulesAction {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TcpRouteRulesAction{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTcpRouteRulesAction
	}
	r.Destinations = flattenTcpRouteRulesActionDestinationsSlice(c, m["destinations"], res)
	r.OriginalDestination = dcl.FlattenBool(m["originalDestination"])

	return r
}

// expandTcpRouteRulesActionDestinationsMap expands the contents of TcpRouteRulesActionDestinations into a JSON
// request object.
func expandTcpRouteRulesActionDestinationsMap(c *Client, f map[string]TcpRouteRulesActionDestinations, res *TcpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTcpRouteRulesActionDestinations(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTcpRouteRulesActionDestinationsSlice expands the contents of TcpRouteRulesActionDestinations into a JSON
// request object.
func expandTcpRouteRulesActionDestinationsSlice(c *Client, f []TcpRouteRulesActionDestinations, res *TcpRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTcpRouteRulesActionDestinations(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTcpRouteRulesActionDestinationsMap flattens the contents of TcpRouteRulesActionDestinations from a JSON
// response object.
func flattenTcpRouteRulesActionDestinationsMap(c *Client, i interface{}, res *TcpRoute) map[string]TcpRouteRulesActionDestinations {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TcpRouteRulesActionDestinations{}
	}

	if len(a) == 0 {
		return map[string]TcpRouteRulesActionDestinations{}
	}

	items := make(map[string]TcpRouteRulesActionDestinations)
	for k, item := range a {
		items[k] = *flattenTcpRouteRulesActionDestinations(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTcpRouteRulesActionDestinationsSlice flattens the contents of TcpRouteRulesActionDestinations from a JSON
// response object.
func flattenTcpRouteRulesActionDestinationsSlice(c *Client, i interface{}, res *TcpRoute) []TcpRouteRulesActionDestinations {
	a, ok := i.([]interface{})
	if !ok {
		return []TcpRouteRulesActionDestinations{}
	}

	if len(a) == 0 {
		return []TcpRouteRulesActionDestinations{}
	}

	items := make([]TcpRouteRulesActionDestinations, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTcpRouteRulesActionDestinations(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTcpRouteRulesActionDestinations expands an instance of TcpRouteRulesActionDestinations into a JSON
// request object.
func expandTcpRouteRulesActionDestinations(c *Client, f *TcpRouteRulesActionDestinations, res *TcpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Weight; !dcl.IsEmptyValueIndirect(v) {
		m["weight"] = v
	}
	if v := f.ServiceName; !dcl.IsEmptyValueIndirect(v) {
		m["serviceName"] = v
	}

	return m, nil
}

// flattenTcpRouteRulesActionDestinations flattens an instance of TcpRouteRulesActionDestinations from a JSON
// response object.
func flattenTcpRouteRulesActionDestinations(c *Client, i interface{}, res *TcpRoute) *TcpRouteRulesActionDestinations {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TcpRouteRulesActionDestinations{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTcpRouteRulesActionDestinations
	}
	r.Weight = dcl.FlattenInteger(m["weight"])
	r.ServiceName = dcl.FlattenString(m["serviceName"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *TcpRoute) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalTcpRoute(b, c, r)
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

type tcpRouteDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         tcpRouteApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToTcpRouteDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]tcpRouteDiff, error) {
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
	var diffs []tcpRouteDiff
	// For each operation name, create a tcpRouteDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := tcpRouteDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToTcpRouteApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToTcpRouteApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (tcpRouteApiOperation, error) {
	switch opName {

	case "updateTcpRouteUpdateTcpRouteOperation":
		return &updateTcpRouteUpdateTcpRouteOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractTcpRouteFields(r *TcpRoute) error {
	return nil
}
func extractTcpRouteRulesFields(r *TcpRoute, o *TcpRouteRules) error {
	vAction := o.Action
	if vAction == nil {
		// note: explicitly not the empty object.
		vAction = &TcpRouteRulesAction{}
	}
	if err := extractTcpRouteRulesActionFields(r, vAction); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAction) {
		o.Action = vAction
	}
	return nil
}
func extractTcpRouteRulesMatchesFields(r *TcpRoute, o *TcpRouteRulesMatches) error {
	return nil
}
func extractTcpRouteRulesActionFields(r *TcpRoute, o *TcpRouteRulesAction) error {
	return nil
}
func extractTcpRouteRulesActionDestinationsFields(r *TcpRoute, o *TcpRouteRulesActionDestinations) error {
	return nil
}

func postReadExtractTcpRouteFields(r *TcpRoute) error {
	return nil
}
func postReadExtractTcpRouteRulesFields(r *TcpRoute, o *TcpRouteRules) error {
	vAction := o.Action
	if vAction == nil {
		// note: explicitly not the empty object.
		vAction = &TcpRouteRulesAction{}
	}
	if err := extractTcpRouteRulesActionFields(r, vAction); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAction) {
		o.Action = vAction
	}
	return nil
}
func postReadExtractTcpRouteRulesMatchesFields(r *TcpRoute, o *TcpRouteRulesMatches) error {
	return nil
}
func postReadExtractTcpRouteRulesActionFields(r *TcpRoute, o *TcpRouteRulesAction) error {
	return nil
}
func postReadExtractTcpRouteRulesActionDestinationsFields(r *TcpRoute, o *TcpRouteRulesActionDestinations) error {
	return nil
}
