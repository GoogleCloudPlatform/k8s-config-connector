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
package networkservices

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

func (r *HttpRoute) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "hostnames"); err != nil {
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
func (r *HttpRouteRules) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Action) {
		if err := r.Action.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *HttpRouteRulesMatches) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"FullPathMatch", "PrefixMatch", "RegexMatch"}, r.FullPathMatch, r.PrefixMatch, r.RegexMatch); err != nil {
		return err
	}
	return nil
}
func (r *HttpRouteRulesMatchesHeaders) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"ExactMatch", "RegexMatch", "PrefixMatch", "PresentMatch", "SuffixMatch", "RangeMatch"}, r.ExactMatch, r.RegexMatch, r.PrefixMatch, r.PresentMatch, r.SuffixMatch, r.RangeMatch); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.RangeMatch) {
		if err := r.RangeMatch.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *HttpRouteRulesMatchesHeadersRangeMatch) validate() error {
	return nil
}
func (r *HttpRouteRulesMatchesQueryParameters) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"ExactMatch", "RegexMatch", "PresentMatch"}, r.ExactMatch, r.RegexMatch, r.PresentMatch); err != nil {
		return err
	}
	return nil
}
func (r *HttpRouteRulesAction) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Redirect) {
		if err := r.Redirect.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.FaultInjectionPolicy) {
		if err := r.FaultInjectionPolicy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.RequestHeaderModifier) {
		if err := r.RequestHeaderModifier.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ResponseHeaderModifier) {
		if err := r.ResponseHeaderModifier.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.UrlRewrite) {
		if err := r.UrlRewrite.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.RetryPolicy) {
		if err := r.RetryPolicy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.RequestMirrorPolicy) {
		if err := r.RequestMirrorPolicy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CorsPolicy) {
		if err := r.CorsPolicy.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *HttpRouteRulesActionDestinations) validate() error {
	return nil
}
func (r *HttpRouteRulesActionRedirect) validate() error {
	return nil
}
func (r *HttpRouteRulesActionFaultInjectionPolicy) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Delay) {
		if err := r.Delay.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Abort) {
		if err := r.Abort.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *HttpRouteRulesActionFaultInjectionPolicyDelay) validate() error {
	return nil
}
func (r *HttpRouteRulesActionFaultInjectionPolicyAbort) validate() error {
	return nil
}
func (r *HttpRouteRulesActionRequestHeaderModifier) validate() error {
	return nil
}
func (r *HttpRouteRulesActionResponseHeaderModifier) validate() error {
	return nil
}
func (r *HttpRouteRulesActionUrlRewrite) validate() error {
	return nil
}
func (r *HttpRouteRulesActionRetryPolicy) validate() error {
	return nil
}
func (r *HttpRouteRulesActionRequestMirrorPolicy) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Destination) {
		if err := r.Destination.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *HttpRouteRulesActionRequestMirrorPolicyDestination) validate() error {
	return nil
}
func (r *HttpRouteRulesActionCorsPolicy) validate() error {
	return nil
}
func (r *HttpRoute) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://networkservices.googleapis.com/v1/", params)
}

func (r *HttpRoute) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/httpRoutes/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *HttpRoute) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/httpRoutes", nr.basePath(), userBasePath, params), nil

}

func (r *HttpRoute) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/httpRoutes?httpRouteId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *HttpRoute) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/httpRoutes/{{name}}", nr.basePath(), userBasePath, params), nil
}

// httpRouteApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type httpRouteApiOperation interface {
	do(context.Context, *HttpRoute, *Client) error
}

// newUpdateHttpRouteUpdateHttpRouteRequest creates a request for an
// HttpRoute resource's UpdateHttpRoute update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateHttpRouteUpdateHttpRouteRequest(ctx context.Context, f *HttpRoute, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v, err := dcl.DeriveField("projects/%s/locations/global/httpRoutes/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Hostnames; v != nil {
		req["hostnames"] = v
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
	if v, err := expandHttpRouteRulesSlice(c, f.Rules, res); err != nil {
		return nil, fmt.Errorf("error expanding Rules into rules: %w", err)
	} else if v != nil {
		req["rules"] = v
	}
	req["name"] = fmt.Sprintf("projects/%s/locations/%s/httpRoutes/%s", *f.Project, *f.Location, *f.Name)

	return req, nil
}

// marshalUpdateHttpRouteUpdateHttpRouteRequest converts the update into
// the final JSON request body.
func marshalUpdateHttpRouteUpdateHttpRouteRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateHttpRouteUpdateHttpRouteOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateHttpRouteUpdateHttpRouteOperation) do(ctx context.Context, r *HttpRoute, c *Client) error {
	_, err := c.GetHttpRoute(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateHttpRoute")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateHttpRouteUpdateHttpRouteRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateHttpRouteUpdateHttpRouteRequest(c, req)
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

func (c *Client) listHttpRouteRaw(ctx context.Context, r *HttpRoute, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != HttpRouteMaxPage {
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

type listHttpRouteOperation struct {
	HttpRoutes []map[string]interface{} `json:"httpRoutes"`
	Token      string                   `json:"nextPageToken"`
}

func (c *Client) listHttpRoute(ctx context.Context, r *HttpRoute, pageToken string, pageSize int32) ([]*HttpRoute, string, error) {
	b, err := c.listHttpRouteRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listHttpRouteOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*HttpRoute
	for _, v := range m.HttpRoutes {
		res, err := unmarshalMapHttpRoute(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllHttpRoute(ctx context.Context, f func(*HttpRoute) bool, resources []*HttpRoute) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteHttpRoute(ctx, res)
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

type deleteHttpRouteOperation struct{}

func (op *deleteHttpRouteOperation) do(ctx context.Context, r *HttpRoute, c *Client) error {
	r, err := c.GetHttpRoute(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "HttpRoute not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetHttpRoute checking for existence. error: %v", err)
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
		_, err := c.GetHttpRoute(ctx, r)
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
type createHttpRouteOperation struct {
	response map[string]interface{}
}

func (op *createHttpRouteOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createHttpRouteOperation) do(ctx context.Context, r *HttpRoute, c *Client) error {
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

	if _, err := c.GetHttpRoute(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getHttpRouteRaw(ctx context.Context, r *HttpRoute) ([]byte, error) {

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

func (c *Client) httpRouteDiffsForRawDesired(ctx context.Context, rawDesired *HttpRoute, opts ...dcl.ApplyOption) (initial, desired *HttpRoute, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *HttpRoute
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*HttpRoute); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected HttpRoute, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetHttpRoute(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a HttpRoute resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve HttpRoute resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that HttpRoute resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeHttpRouteDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for HttpRoute: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for HttpRoute: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractHttpRouteFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeHttpRouteInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for HttpRoute: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeHttpRouteDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for HttpRoute: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffHttpRoute(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeHttpRouteInitialState(rawInitial, rawDesired *HttpRoute) (*HttpRoute, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeHttpRouteDesiredState(rawDesired, rawInitial *HttpRoute, opts ...dcl.ApplyOption) (*HttpRoute, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &HttpRoute{}
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
	if dcl.StringArrayCanonicalize(rawDesired.Hostnames, rawInitial.Hostnames) {
		canonicalDesired.Hostnames = rawInitial.Hostnames
	} else {
		canonicalDesired.Hostnames = rawDesired.Hostnames
	}
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
	canonicalDesired.Rules = canonicalizeHttpRouteRulesSlice(rawDesired.Rules, rawInitial.Rules, opts...)
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

func canonicalizeHttpRouteNewState(c *Client, rawNew, rawDesired *HttpRoute) (*HttpRoute, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
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

	if dcl.IsEmptyValueIndirect(rawNew.Hostnames) && dcl.IsEmptyValueIndirect(rawDesired.Hostnames) {
		rawNew.Hostnames = rawDesired.Hostnames
	} else {
		if dcl.StringArrayCanonicalize(rawDesired.Hostnames, rawNew.Hostnames) {
			rawNew.Hostnames = rawDesired.Hostnames
		}
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

	if dcl.IsEmptyValueIndirect(rawNew.Rules) && dcl.IsEmptyValueIndirect(rawDesired.Rules) {
		rawNew.Rules = rawDesired.Rules
	} else {
		rawNew.Rules = canonicalizeNewHttpRouteRulesSlice(c, rawDesired.Rules, rawNew.Rules)
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

func canonicalizeHttpRouteRules(des, initial *HttpRouteRules, opts ...dcl.ApplyOption) *HttpRouteRules {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &HttpRouteRules{}

	cDes.Matches = canonicalizeHttpRouteRulesMatchesSlice(des.Matches, initial.Matches, opts...)
	cDes.Action = canonicalizeHttpRouteRulesAction(des.Action, initial.Action, opts...)

	return cDes
}

func canonicalizeHttpRouteRulesSlice(des, initial []HttpRouteRules, opts ...dcl.ApplyOption) []HttpRouteRules {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]HttpRouteRules, 0, len(des))
		for _, d := range des {
			cd := canonicalizeHttpRouteRules(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]HttpRouteRules, 0, len(des))
	for i, d := range des {
		cd := canonicalizeHttpRouteRules(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewHttpRouteRules(c *Client, des, nw *HttpRouteRules) *HttpRouteRules {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for HttpRouteRules while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Matches = canonicalizeNewHttpRouteRulesMatchesSlice(c, des.Matches, nw.Matches)
	nw.Action = canonicalizeNewHttpRouteRulesAction(c, des.Action, nw.Action)

	return nw
}

func canonicalizeNewHttpRouteRulesSet(c *Client, des, nw []HttpRouteRules) []HttpRouteRules {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []HttpRouteRules
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareHttpRouteRulesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewHttpRouteRules(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewHttpRouteRulesSlice(c *Client, des, nw []HttpRouteRules) []HttpRouteRules {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []HttpRouteRules
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewHttpRouteRules(c, &d, &n))
	}

	return items
}

func canonicalizeHttpRouteRulesMatches(des, initial *HttpRouteRulesMatches, opts ...dcl.ApplyOption) *HttpRouteRulesMatches {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.FullPathMatch != nil || (initial != nil && initial.FullPathMatch != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.PrefixMatch, des.RegexMatch) {
			des.FullPathMatch = nil
			if initial != nil {
				initial.FullPathMatch = nil
			}
		}
	}

	if des.PrefixMatch != nil || (initial != nil && initial.PrefixMatch != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.FullPathMatch, des.RegexMatch) {
			des.PrefixMatch = nil
			if initial != nil {
				initial.PrefixMatch = nil
			}
		}
	}

	if des.RegexMatch != nil || (initial != nil && initial.RegexMatch != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.FullPathMatch, des.PrefixMatch) {
			des.RegexMatch = nil
			if initial != nil {
				initial.RegexMatch = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &HttpRouteRulesMatches{}

	if dcl.StringCanonicalize(des.FullPathMatch, initial.FullPathMatch) || dcl.IsZeroValue(des.FullPathMatch) {
		cDes.FullPathMatch = initial.FullPathMatch
	} else {
		cDes.FullPathMatch = des.FullPathMatch
	}
	if dcl.StringCanonicalize(des.PrefixMatch, initial.PrefixMatch) || dcl.IsZeroValue(des.PrefixMatch) {
		cDes.PrefixMatch = initial.PrefixMatch
	} else {
		cDes.PrefixMatch = des.PrefixMatch
	}
	if dcl.StringCanonicalize(des.RegexMatch, initial.RegexMatch) || dcl.IsZeroValue(des.RegexMatch) {
		cDes.RegexMatch = initial.RegexMatch
	} else {
		cDes.RegexMatch = des.RegexMatch
	}
	if dcl.BoolCanonicalize(des.IgnoreCase, initial.IgnoreCase) || dcl.IsZeroValue(des.IgnoreCase) {
		cDes.IgnoreCase = initial.IgnoreCase
	} else {
		cDes.IgnoreCase = des.IgnoreCase
	}
	cDes.Headers = canonicalizeHttpRouteRulesMatchesHeadersSlice(des.Headers, initial.Headers, opts...)
	cDes.QueryParameters = canonicalizeHttpRouteRulesMatchesQueryParametersSlice(des.QueryParameters, initial.QueryParameters, opts...)

	return cDes
}

func canonicalizeHttpRouteRulesMatchesSlice(des, initial []HttpRouteRulesMatches, opts ...dcl.ApplyOption) []HttpRouteRulesMatches {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]HttpRouteRulesMatches, 0, len(des))
		for _, d := range des {
			cd := canonicalizeHttpRouteRulesMatches(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]HttpRouteRulesMatches, 0, len(des))
	for i, d := range des {
		cd := canonicalizeHttpRouteRulesMatches(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewHttpRouteRulesMatches(c *Client, des, nw *HttpRouteRulesMatches) *HttpRouteRulesMatches {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for HttpRouteRulesMatches while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.FullPathMatch, nw.FullPathMatch) {
		nw.FullPathMatch = des.FullPathMatch
	}
	if dcl.StringCanonicalize(des.PrefixMatch, nw.PrefixMatch) {
		nw.PrefixMatch = des.PrefixMatch
	}
	if dcl.StringCanonicalize(des.RegexMatch, nw.RegexMatch) {
		nw.RegexMatch = des.RegexMatch
	}
	if dcl.BoolCanonicalize(des.IgnoreCase, nw.IgnoreCase) {
		nw.IgnoreCase = des.IgnoreCase
	}
	nw.Headers = canonicalizeNewHttpRouteRulesMatchesHeadersSlice(c, des.Headers, nw.Headers)
	nw.QueryParameters = canonicalizeNewHttpRouteRulesMatchesQueryParametersSlice(c, des.QueryParameters, nw.QueryParameters)

	return nw
}

func canonicalizeNewHttpRouteRulesMatchesSet(c *Client, des, nw []HttpRouteRulesMatches) []HttpRouteRulesMatches {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []HttpRouteRulesMatches
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareHttpRouteRulesMatchesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewHttpRouteRulesMatches(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewHttpRouteRulesMatchesSlice(c *Client, des, nw []HttpRouteRulesMatches) []HttpRouteRulesMatches {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []HttpRouteRulesMatches
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewHttpRouteRulesMatches(c, &d, &n))
	}

	return items
}

func canonicalizeHttpRouteRulesMatchesHeaders(des, initial *HttpRouteRulesMatchesHeaders, opts ...dcl.ApplyOption) *HttpRouteRulesMatchesHeaders {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.ExactMatch != nil || (initial != nil && initial.ExactMatch != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.RegexMatch, des.PrefixMatch, des.PresentMatch, des.SuffixMatch, des.RangeMatch) {
			des.ExactMatch = nil
			if initial != nil {
				initial.ExactMatch = nil
			}
		}
	}

	if des.RegexMatch != nil || (initial != nil && initial.RegexMatch != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.ExactMatch, des.PrefixMatch, des.PresentMatch, des.SuffixMatch, des.RangeMatch) {
			des.RegexMatch = nil
			if initial != nil {
				initial.RegexMatch = nil
			}
		}
	}

	if des.PrefixMatch != nil || (initial != nil && initial.PrefixMatch != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.ExactMatch, des.RegexMatch, des.PresentMatch, des.SuffixMatch, des.RangeMatch) {
			des.PrefixMatch = nil
			if initial != nil {
				initial.PrefixMatch = nil
			}
		}
	}

	if des.PresentMatch != nil || (initial != nil && initial.PresentMatch != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.ExactMatch, des.RegexMatch, des.PrefixMatch, des.SuffixMatch, des.RangeMatch) {
			des.PresentMatch = nil
			if initial != nil {
				initial.PresentMatch = nil
			}
		}
	}

	if des.SuffixMatch != nil || (initial != nil && initial.SuffixMatch != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.ExactMatch, des.RegexMatch, des.PrefixMatch, des.PresentMatch, des.RangeMatch) {
			des.SuffixMatch = nil
			if initial != nil {
				initial.SuffixMatch = nil
			}
		}
	}

	if des.RangeMatch != nil || (initial != nil && initial.RangeMatch != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.ExactMatch, des.RegexMatch, des.PrefixMatch, des.PresentMatch, des.SuffixMatch) {
			des.RangeMatch = nil
			if initial != nil {
				initial.RangeMatch = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &HttpRouteRulesMatchesHeaders{}

	if dcl.StringCanonicalize(des.Header, initial.Header) || dcl.IsZeroValue(des.Header) {
		cDes.Header = initial.Header
	} else {
		cDes.Header = des.Header
	}
	if dcl.StringCanonicalize(des.ExactMatch, initial.ExactMatch) || dcl.IsZeroValue(des.ExactMatch) {
		cDes.ExactMatch = initial.ExactMatch
	} else {
		cDes.ExactMatch = des.ExactMatch
	}
	if dcl.StringCanonicalize(des.RegexMatch, initial.RegexMatch) || dcl.IsZeroValue(des.RegexMatch) {
		cDes.RegexMatch = initial.RegexMatch
	} else {
		cDes.RegexMatch = des.RegexMatch
	}
	if dcl.StringCanonicalize(des.PrefixMatch, initial.PrefixMatch) || dcl.IsZeroValue(des.PrefixMatch) {
		cDes.PrefixMatch = initial.PrefixMatch
	} else {
		cDes.PrefixMatch = des.PrefixMatch
	}
	if dcl.BoolCanonicalize(des.PresentMatch, initial.PresentMatch) || dcl.IsZeroValue(des.PresentMatch) {
		cDes.PresentMatch = initial.PresentMatch
	} else {
		cDes.PresentMatch = des.PresentMatch
	}
	if dcl.StringCanonicalize(des.SuffixMatch, initial.SuffixMatch) || dcl.IsZeroValue(des.SuffixMatch) {
		cDes.SuffixMatch = initial.SuffixMatch
	} else {
		cDes.SuffixMatch = des.SuffixMatch
	}
	cDes.RangeMatch = canonicalizeHttpRouteRulesMatchesHeadersRangeMatch(des.RangeMatch, initial.RangeMatch, opts...)
	if dcl.BoolCanonicalize(des.InvertMatch, initial.InvertMatch) || dcl.IsZeroValue(des.InvertMatch) {
		cDes.InvertMatch = initial.InvertMatch
	} else {
		cDes.InvertMatch = des.InvertMatch
	}

	return cDes
}

func canonicalizeHttpRouteRulesMatchesHeadersSlice(des, initial []HttpRouteRulesMatchesHeaders, opts ...dcl.ApplyOption) []HttpRouteRulesMatchesHeaders {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]HttpRouteRulesMatchesHeaders, 0, len(des))
		for _, d := range des {
			cd := canonicalizeHttpRouteRulesMatchesHeaders(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]HttpRouteRulesMatchesHeaders, 0, len(des))
	for i, d := range des {
		cd := canonicalizeHttpRouteRulesMatchesHeaders(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewHttpRouteRulesMatchesHeaders(c *Client, des, nw *HttpRouteRulesMatchesHeaders) *HttpRouteRulesMatchesHeaders {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for HttpRouteRulesMatchesHeaders while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Header, nw.Header) {
		nw.Header = des.Header
	}
	if dcl.StringCanonicalize(des.ExactMatch, nw.ExactMatch) {
		nw.ExactMatch = des.ExactMatch
	}
	if dcl.StringCanonicalize(des.RegexMatch, nw.RegexMatch) {
		nw.RegexMatch = des.RegexMatch
	}
	if dcl.StringCanonicalize(des.PrefixMatch, nw.PrefixMatch) {
		nw.PrefixMatch = des.PrefixMatch
	}
	if dcl.BoolCanonicalize(des.PresentMatch, nw.PresentMatch) {
		nw.PresentMatch = des.PresentMatch
	}
	if dcl.StringCanonicalize(des.SuffixMatch, nw.SuffixMatch) {
		nw.SuffixMatch = des.SuffixMatch
	}
	nw.RangeMatch = canonicalizeNewHttpRouteRulesMatchesHeadersRangeMatch(c, des.RangeMatch, nw.RangeMatch)
	if dcl.BoolCanonicalize(des.InvertMatch, nw.InvertMatch) {
		nw.InvertMatch = des.InvertMatch
	}

	return nw
}

func canonicalizeNewHttpRouteRulesMatchesHeadersSet(c *Client, des, nw []HttpRouteRulesMatchesHeaders) []HttpRouteRulesMatchesHeaders {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []HttpRouteRulesMatchesHeaders
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareHttpRouteRulesMatchesHeadersNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewHttpRouteRulesMatchesHeaders(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewHttpRouteRulesMatchesHeadersSlice(c *Client, des, nw []HttpRouteRulesMatchesHeaders) []HttpRouteRulesMatchesHeaders {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []HttpRouteRulesMatchesHeaders
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewHttpRouteRulesMatchesHeaders(c, &d, &n))
	}

	return items
}

func canonicalizeHttpRouteRulesMatchesHeadersRangeMatch(des, initial *HttpRouteRulesMatchesHeadersRangeMatch, opts ...dcl.ApplyOption) *HttpRouteRulesMatchesHeadersRangeMatch {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &HttpRouteRulesMatchesHeadersRangeMatch{}

	if dcl.IsZeroValue(des.Start) || (dcl.IsEmptyValueIndirect(des.Start) && dcl.IsEmptyValueIndirect(initial.Start)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Start = initial.Start
	} else {
		cDes.Start = des.Start
	}
	if dcl.IsZeroValue(des.End) || (dcl.IsEmptyValueIndirect(des.End) && dcl.IsEmptyValueIndirect(initial.End)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.End = initial.End
	} else {
		cDes.End = des.End
	}

	return cDes
}

func canonicalizeHttpRouteRulesMatchesHeadersRangeMatchSlice(des, initial []HttpRouteRulesMatchesHeadersRangeMatch, opts ...dcl.ApplyOption) []HttpRouteRulesMatchesHeadersRangeMatch {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]HttpRouteRulesMatchesHeadersRangeMatch, 0, len(des))
		for _, d := range des {
			cd := canonicalizeHttpRouteRulesMatchesHeadersRangeMatch(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]HttpRouteRulesMatchesHeadersRangeMatch, 0, len(des))
	for i, d := range des {
		cd := canonicalizeHttpRouteRulesMatchesHeadersRangeMatch(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewHttpRouteRulesMatchesHeadersRangeMatch(c *Client, des, nw *HttpRouteRulesMatchesHeadersRangeMatch) *HttpRouteRulesMatchesHeadersRangeMatch {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for HttpRouteRulesMatchesHeadersRangeMatch while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewHttpRouteRulesMatchesHeadersRangeMatchSet(c *Client, des, nw []HttpRouteRulesMatchesHeadersRangeMatch) []HttpRouteRulesMatchesHeadersRangeMatch {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []HttpRouteRulesMatchesHeadersRangeMatch
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareHttpRouteRulesMatchesHeadersRangeMatchNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewHttpRouteRulesMatchesHeadersRangeMatch(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewHttpRouteRulesMatchesHeadersRangeMatchSlice(c *Client, des, nw []HttpRouteRulesMatchesHeadersRangeMatch) []HttpRouteRulesMatchesHeadersRangeMatch {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []HttpRouteRulesMatchesHeadersRangeMatch
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewHttpRouteRulesMatchesHeadersRangeMatch(c, &d, &n))
	}

	return items
}

func canonicalizeHttpRouteRulesMatchesQueryParameters(des, initial *HttpRouteRulesMatchesQueryParameters, opts ...dcl.ApplyOption) *HttpRouteRulesMatchesQueryParameters {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.ExactMatch != nil || (initial != nil && initial.ExactMatch != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.RegexMatch, des.PresentMatch) {
			des.ExactMatch = nil
			if initial != nil {
				initial.ExactMatch = nil
			}
		}
	}

	if des.RegexMatch != nil || (initial != nil && initial.RegexMatch != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.ExactMatch, des.PresentMatch) {
			des.RegexMatch = nil
			if initial != nil {
				initial.RegexMatch = nil
			}
		}
	}

	if des.PresentMatch != nil || (initial != nil && initial.PresentMatch != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.ExactMatch, des.RegexMatch) {
			des.PresentMatch = nil
			if initial != nil {
				initial.PresentMatch = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &HttpRouteRulesMatchesQueryParameters{}

	if dcl.StringCanonicalize(des.QueryParameter, initial.QueryParameter) || dcl.IsZeroValue(des.QueryParameter) {
		cDes.QueryParameter = initial.QueryParameter
	} else {
		cDes.QueryParameter = des.QueryParameter
	}
	if dcl.StringCanonicalize(des.ExactMatch, initial.ExactMatch) || dcl.IsZeroValue(des.ExactMatch) {
		cDes.ExactMatch = initial.ExactMatch
	} else {
		cDes.ExactMatch = des.ExactMatch
	}
	if dcl.StringCanonicalize(des.RegexMatch, initial.RegexMatch) || dcl.IsZeroValue(des.RegexMatch) {
		cDes.RegexMatch = initial.RegexMatch
	} else {
		cDes.RegexMatch = des.RegexMatch
	}
	if dcl.BoolCanonicalize(des.PresentMatch, initial.PresentMatch) || dcl.IsZeroValue(des.PresentMatch) {
		cDes.PresentMatch = initial.PresentMatch
	} else {
		cDes.PresentMatch = des.PresentMatch
	}

	return cDes
}

func canonicalizeHttpRouteRulesMatchesQueryParametersSlice(des, initial []HttpRouteRulesMatchesQueryParameters, opts ...dcl.ApplyOption) []HttpRouteRulesMatchesQueryParameters {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]HttpRouteRulesMatchesQueryParameters, 0, len(des))
		for _, d := range des {
			cd := canonicalizeHttpRouteRulesMatchesQueryParameters(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]HttpRouteRulesMatchesQueryParameters, 0, len(des))
	for i, d := range des {
		cd := canonicalizeHttpRouteRulesMatchesQueryParameters(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewHttpRouteRulesMatchesQueryParameters(c *Client, des, nw *HttpRouteRulesMatchesQueryParameters) *HttpRouteRulesMatchesQueryParameters {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for HttpRouteRulesMatchesQueryParameters while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.QueryParameter, nw.QueryParameter) {
		nw.QueryParameter = des.QueryParameter
	}
	if dcl.StringCanonicalize(des.ExactMatch, nw.ExactMatch) {
		nw.ExactMatch = des.ExactMatch
	}
	if dcl.StringCanonicalize(des.RegexMatch, nw.RegexMatch) {
		nw.RegexMatch = des.RegexMatch
	}
	if dcl.BoolCanonicalize(des.PresentMatch, nw.PresentMatch) {
		nw.PresentMatch = des.PresentMatch
	}

	return nw
}

func canonicalizeNewHttpRouteRulesMatchesQueryParametersSet(c *Client, des, nw []HttpRouteRulesMatchesQueryParameters) []HttpRouteRulesMatchesQueryParameters {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []HttpRouteRulesMatchesQueryParameters
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareHttpRouteRulesMatchesQueryParametersNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewHttpRouteRulesMatchesQueryParameters(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewHttpRouteRulesMatchesQueryParametersSlice(c *Client, des, nw []HttpRouteRulesMatchesQueryParameters) []HttpRouteRulesMatchesQueryParameters {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []HttpRouteRulesMatchesQueryParameters
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewHttpRouteRulesMatchesQueryParameters(c, &d, &n))
	}

	return items
}

func canonicalizeHttpRouteRulesAction(des, initial *HttpRouteRulesAction, opts ...dcl.ApplyOption) *HttpRouteRulesAction {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &HttpRouteRulesAction{}

	cDes.Destinations = canonicalizeHttpRouteRulesActionDestinationsSlice(des.Destinations, initial.Destinations, opts...)
	cDes.Redirect = canonicalizeHttpRouteRulesActionRedirect(des.Redirect, initial.Redirect, opts...)
	cDes.FaultInjectionPolicy = canonicalizeHttpRouteRulesActionFaultInjectionPolicy(des.FaultInjectionPolicy, initial.FaultInjectionPolicy, opts...)
	cDes.RequestHeaderModifier = canonicalizeHttpRouteRulesActionRequestHeaderModifier(des.RequestHeaderModifier, initial.RequestHeaderModifier, opts...)
	cDes.ResponseHeaderModifier = canonicalizeHttpRouteRulesActionResponseHeaderModifier(des.ResponseHeaderModifier, initial.ResponseHeaderModifier, opts...)
	cDes.UrlRewrite = canonicalizeHttpRouteRulesActionUrlRewrite(des.UrlRewrite, initial.UrlRewrite, opts...)
	if dcl.StringCanonicalize(des.Timeout, initial.Timeout) || dcl.IsZeroValue(des.Timeout) {
		cDes.Timeout = initial.Timeout
	} else {
		cDes.Timeout = des.Timeout
	}
	cDes.RetryPolicy = canonicalizeHttpRouteRulesActionRetryPolicy(des.RetryPolicy, initial.RetryPolicy, opts...)
	cDes.RequestMirrorPolicy = canonicalizeHttpRouteRulesActionRequestMirrorPolicy(des.RequestMirrorPolicy, initial.RequestMirrorPolicy, opts...)
	cDes.CorsPolicy = canonicalizeHttpRouteRulesActionCorsPolicy(des.CorsPolicy, initial.CorsPolicy, opts...)

	return cDes
}

func canonicalizeHttpRouteRulesActionSlice(des, initial []HttpRouteRulesAction, opts ...dcl.ApplyOption) []HttpRouteRulesAction {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]HttpRouteRulesAction, 0, len(des))
		for _, d := range des {
			cd := canonicalizeHttpRouteRulesAction(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]HttpRouteRulesAction, 0, len(des))
	for i, d := range des {
		cd := canonicalizeHttpRouteRulesAction(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewHttpRouteRulesAction(c *Client, des, nw *HttpRouteRulesAction) *HttpRouteRulesAction {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for HttpRouteRulesAction while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Destinations = canonicalizeNewHttpRouteRulesActionDestinationsSlice(c, des.Destinations, nw.Destinations)
	nw.Redirect = canonicalizeNewHttpRouteRulesActionRedirect(c, des.Redirect, nw.Redirect)
	nw.FaultInjectionPolicy = canonicalizeNewHttpRouteRulesActionFaultInjectionPolicy(c, des.FaultInjectionPolicy, nw.FaultInjectionPolicy)
	nw.RequestHeaderModifier = canonicalizeNewHttpRouteRulesActionRequestHeaderModifier(c, des.RequestHeaderModifier, nw.RequestHeaderModifier)
	nw.ResponseHeaderModifier = canonicalizeNewHttpRouteRulesActionResponseHeaderModifier(c, des.ResponseHeaderModifier, nw.ResponseHeaderModifier)
	nw.UrlRewrite = canonicalizeNewHttpRouteRulesActionUrlRewrite(c, des.UrlRewrite, nw.UrlRewrite)
	if dcl.StringCanonicalize(des.Timeout, nw.Timeout) {
		nw.Timeout = des.Timeout
	}
	nw.RetryPolicy = canonicalizeNewHttpRouteRulesActionRetryPolicy(c, des.RetryPolicy, nw.RetryPolicy)
	nw.RequestMirrorPolicy = canonicalizeNewHttpRouteRulesActionRequestMirrorPolicy(c, des.RequestMirrorPolicy, nw.RequestMirrorPolicy)
	nw.CorsPolicy = canonicalizeNewHttpRouteRulesActionCorsPolicy(c, des.CorsPolicy, nw.CorsPolicy)

	return nw
}

func canonicalizeNewHttpRouteRulesActionSet(c *Client, des, nw []HttpRouteRulesAction) []HttpRouteRulesAction {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []HttpRouteRulesAction
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareHttpRouteRulesActionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewHttpRouteRulesAction(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewHttpRouteRulesActionSlice(c *Client, des, nw []HttpRouteRulesAction) []HttpRouteRulesAction {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []HttpRouteRulesAction
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewHttpRouteRulesAction(c, &d, &n))
	}

	return items
}

func canonicalizeHttpRouteRulesActionDestinations(des, initial *HttpRouteRulesActionDestinations, opts ...dcl.ApplyOption) *HttpRouteRulesActionDestinations {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &HttpRouteRulesActionDestinations{}

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

func canonicalizeHttpRouteRulesActionDestinationsSlice(des, initial []HttpRouteRulesActionDestinations, opts ...dcl.ApplyOption) []HttpRouteRulesActionDestinations {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]HttpRouteRulesActionDestinations, 0, len(des))
		for _, d := range des {
			cd := canonicalizeHttpRouteRulesActionDestinations(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]HttpRouteRulesActionDestinations, 0, len(des))
	for i, d := range des {
		cd := canonicalizeHttpRouteRulesActionDestinations(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewHttpRouteRulesActionDestinations(c *Client, des, nw *HttpRouteRulesActionDestinations) *HttpRouteRulesActionDestinations {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for HttpRouteRulesActionDestinations while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewHttpRouteRulesActionDestinationsSet(c *Client, des, nw []HttpRouteRulesActionDestinations) []HttpRouteRulesActionDestinations {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []HttpRouteRulesActionDestinations
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareHttpRouteRulesActionDestinationsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewHttpRouteRulesActionDestinations(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewHttpRouteRulesActionDestinationsSlice(c *Client, des, nw []HttpRouteRulesActionDestinations) []HttpRouteRulesActionDestinations {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []HttpRouteRulesActionDestinations
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewHttpRouteRulesActionDestinations(c, &d, &n))
	}

	return items
}

func canonicalizeHttpRouteRulesActionRedirect(des, initial *HttpRouteRulesActionRedirect, opts ...dcl.ApplyOption) *HttpRouteRulesActionRedirect {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &HttpRouteRulesActionRedirect{}

	if dcl.StringCanonicalize(des.HostRedirect, initial.HostRedirect) || dcl.IsZeroValue(des.HostRedirect) {
		cDes.HostRedirect = initial.HostRedirect
	} else {
		cDes.HostRedirect = des.HostRedirect
	}
	if dcl.StringCanonicalize(des.PathRedirect, initial.PathRedirect) || dcl.IsZeroValue(des.PathRedirect) {
		cDes.PathRedirect = initial.PathRedirect
	} else {
		cDes.PathRedirect = des.PathRedirect
	}
	if dcl.StringCanonicalize(des.PrefixRewrite, initial.PrefixRewrite) || dcl.IsZeroValue(des.PrefixRewrite) {
		cDes.PrefixRewrite = initial.PrefixRewrite
	} else {
		cDes.PrefixRewrite = des.PrefixRewrite
	}
	if dcl.IsZeroValue(des.ResponseCode) || (dcl.IsEmptyValueIndirect(des.ResponseCode) && dcl.IsEmptyValueIndirect(initial.ResponseCode)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ResponseCode = initial.ResponseCode
	} else {
		cDes.ResponseCode = des.ResponseCode
	}
	if dcl.BoolCanonicalize(des.HttpsRedirect, initial.HttpsRedirect) || dcl.IsZeroValue(des.HttpsRedirect) {
		cDes.HttpsRedirect = initial.HttpsRedirect
	} else {
		cDes.HttpsRedirect = des.HttpsRedirect
	}
	if dcl.BoolCanonicalize(des.StripQuery, initial.StripQuery) || dcl.IsZeroValue(des.StripQuery) {
		cDes.StripQuery = initial.StripQuery
	} else {
		cDes.StripQuery = des.StripQuery
	}
	if dcl.IsZeroValue(des.PortRedirect) || (dcl.IsEmptyValueIndirect(des.PortRedirect) && dcl.IsEmptyValueIndirect(initial.PortRedirect)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.PortRedirect = initial.PortRedirect
	} else {
		cDes.PortRedirect = des.PortRedirect
	}

	return cDes
}

func canonicalizeHttpRouteRulesActionRedirectSlice(des, initial []HttpRouteRulesActionRedirect, opts ...dcl.ApplyOption) []HttpRouteRulesActionRedirect {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]HttpRouteRulesActionRedirect, 0, len(des))
		for _, d := range des {
			cd := canonicalizeHttpRouteRulesActionRedirect(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]HttpRouteRulesActionRedirect, 0, len(des))
	for i, d := range des {
		cd := canonicalizeHttpRouteRulesActionRedirect(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewHttpRouteRulesActionRedirect(c *Client, des, nw *HttpRouteRulesActionRedirect) *HttpRouteRulesActionRedirect {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for HttpRouteRulesActionRedirect while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.HostRedirect, nw.HostRedirect) {
		nw.HostRedirect = des.HostRedirect
	}
	if dcl.StringCanonicalize(des.PathRedirect, nw.PathRedirect) {
		nw.PathRedirect = des.PathRedirect
	}
	if dcl.StringCanonicalize(des.PrefixRewrite, nw.PrefixRewrite) {
		nw.PrefixRewrite = des.PrefixRewrite
	}
	if dcl.BoolCanonicalize(des.HttpsRedirect, nw.HttpsRedirect) {
		nw.HttpsRedirect = des.HttpsRedirect
	}
	if dcl.BoolCanonicalize(des.StripQuery, nw.StripQuery) {
		nw.StripQuery = des.StripQuery
	}

	return nw
}

func canonicalizeNewHttpRouteRulesActionRedirectSet(c *Client, des, nw []HttpRouteRulesActionRedirect) []HttpRouteRulesActionRedirect {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []HttpRouteRulesActionRedirect
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareHttpRouteRulesActionRedirectNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewHttpRouteRulesActionRedirect(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewHttpRouteRulesActionRedirectSlice(c *Client, des, nw []HttpRouteRulesActionRedirect) []HttpRouteRulesActionRedirect {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []HttpRouteRulesActionRedirect
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewHttpRouteRulesActionRedirect(c, &d, &n))
	}

	return items
}

func canonicalizeHttpRouteRulesActionFaultInjectionPolicy(des, initial *HttpRouteRulesActionFaultInjectionPolicy, opts ...dcl.ApplyOption) *HttpRouteRulesActionFaultInjectionPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &HttpRouteRulesActionFaultInjectionPolicy{}

	cDes.Delay = canonicalizeHttpRouteRulesActionFaultInjectionPolicyDelay(des.Delay, initial.Delay, opts...)
	cDes.Abort = canonicalizeHttpRouteRulesActionFaultInjectionPolicyAbort(des.Abort, initial.Abort, opts...)

	return cDes
}

func canonicalizeHttpRouteRulesActionFaultInjectionPolicySlice(des, initial []HttpRouteRulesActionFaultInjectionPolicy, opts ...dcl.ApplyOption) []HttpRouteRulesActionFaultInjectionPolicy {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]HttpRouteRulesActionFaultInjectionPolicy, 0, len(des))
		for _, d := range des {
			cd := canonicalizeHttpRouteRulesActionFaultInjectionPolicy(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]HttpRouteRulesActionFaultInjectionPolicy, 0, len(des))
	for i, d := range des {
		cd := canonicalizeHttpRouteRulesActionFaultInjectionPolicy(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewHttpRouteRulesActionFaultInjectionPolicy(c *Client, des, nw *HttpRouteRulesActionFaultInjectionPolicy) *HttpRouteRulesActionFaultInjectionPolicy {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for HttpRouteRulesActionFaultInjectionPolicy while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Delay = canonicalizeNewHttpRouteRulesActionFaultInjectionPolicyDelay(c, des.Delay, nw.Delay)
	nw.Abort = canonicalizeNewHttpRouteRulesActionFaultInjectionPolicyAbort(c, des.Abort, nw.Abort)

	return nw
}

func canonicalizeNewHttpRouteRulesActionFaultInjectionPolicySet(c *Client, des, nw []HttpRouteRulesActionFaultInjectionPolicy) []HttpRouteRulesActionFaultInjectionPolicy {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []HttpRouteRulesActionFaultInjectionPolicy
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareHttpRouteRulesActionFaultInjectionPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewHttpRouteRulesActionFaultInjectionPolicy(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewHttpRouteRulesActionFaultInjectionPolicySlice(c *Client, des, nw []HttpRouteRulesActionFaultInjectionPolicy) []HttpRouteRulesActionFaultInjectionPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []HttpRouteRulesActionFaultInjectionPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewHttpRouteRulesActionFaultInjectionPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeHttpRouteRulesActionFaultInjectionPolicyDelay(des, initial *HttpRouteRulesActionFaultInjectionPolicyDelay, opts ...dcl.ApplyOption) *HttpRouteRulesActionFaultInjectionPolicyDelay {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &HttpRouteRulesActionFaultInjectionPolicyDelay{}

	if dcl.StringCanonicalize(des.FixedDelay, initial.FixedDelay) || dcl.IsZeroValue(des.FixedDelay) {
		cDes.FixedDelay = initial.FixedDelay
	} else {
		cDes.FixedDelay = des.FixedDelay
	}
	if dcl.IsZeroValue(des.Percentage) || (dcl.IsEmptyValueIndirect(des.Percentage) && dcl.IsEmptyValueIndirect(initial.Percentage)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Percentage = initial.Percentage
	} else {
		cDes.Percentage = des.Percentage
	}

	return cDes
}

func canonicalizeHttpRouteRulesActionFaultInjectionPolicyDelaySlice(des, initial []HttpRouteRulesActionFaultInjectionPolicyDelay, opts ...dcl.ApplyOption) []HttpRouteRulesActionFaultInjectionPolicyDelay {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]HttpRouteRulesActionFaultInjectionPolicyDelay, 0, len(des))
		for _, d := range des {
			cd := canonicalizeHttpRouteRulesActionFaultInjectionPolicyDelay(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]HttpRouteRulesActionFaultInjectionPolicyDelay, 0, len(des))
	for i, d := range des {
		cd := canonicalizeHttpRouteRulesActionFaultInjectionPolicyDelay(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewHttpRouteRulesActionFaultInjectionPolicyDelay(c *Client, des, nw *HttpRouteRulesActionFaultInjectionPolicyDelay) *HttpRouteRulesActionFaultInjectionPolicyDelay {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for HttpRouteRulesActionFaultInjectionPolicyDelay while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.FixedDelay, nw.FixedDelay) {
		nw.FixedDelay = des.FixedDelay
	}

	return nw
}

func canonicalizeNewHttpRouteRulesActionFaultInjectionPolicyDelaySet(c *Client, des, nw []HttpRouteRulesActionFaultInjectionPolicyDelay) []HttpRouteRulesActionFaultInjectionPolicyDelay {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []HttpRouteRulesActionFaultInjectionPolicyDelay
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareHttpRouteRulesActionFaultInjectionPolicyDelayNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewHttpRouteRulesActionFaultInjectionPolicyDelay(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewHttpRouteRulesActionFaultInjectionPolicyDelaySlice(c *Client, des, nw []HttpRouteRulesActionFaultInjectionPolicyDelay) []HttpRouteRulesActionFaultInjectionPolicyDelay {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []HttpRouteRulesActionFaultInjectionPolicyDelay
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewHttpRouteRulesActionFaultInjectionPolicyDelay(c, &d, &n))
	}

	return items
}

func canonicalizeHttpRouteRulesActionFaultInjectionPolicyAbort(des, initial *HttpRouteRulesActionFaultInjectionPolicyAbort, opts ...dcl.ApplyOption) *HttpRouteRulesActionFaultInjectionPolicyAbort {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &HttpRouteRulesActionFaultInjectionPolicyAbort{}

	if dcl.IsZeroValue(des.HttpStatus) || (dcl.IsEmptyValueIndirect(des.HttpStatus) && dcl.IsEmptyValueIndirect(initial.HttpStatus)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.HttpStatus = initial.HttpStatus
	} else {
		cDes.HttpStatus = des.HttpStatus
	}
	if dcl.IsZeroValue(des.Percentage) || (dcl.IsEmptyValueIndirect(des.Percentage) && dcl.IsEmptyValueIndirect(initial.Percentage)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Percentage = initial.Percentage
	} else {
		cDes.Percentage = des.Percentage
	}

	return cDes
}

func canonicalizeHttpRouteRulesActionFaultInjectionPolicyAbortSlice(des, initial []HttpRouteRulesActionFaultInjectionPolicyAbort, opts ...dcl.ApplyOption) []HttpRouteRulesActionFaultInjectionPolicyAbort {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]HttpRouteRulesActionFaultInjectionPolicyAbort, 0, len(des))
		for _, d := range des {
			cd := canonicalizeHttpRouteRulesActionFaultInjectionPolicyAbort(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]HttpRouteRulesActionFaultInjectionPolicyAbort, 0, len(des))
	for i, d := range des {
		cd := canonicalizeHttpRouteRulesActionFaultInjectionPolicyAbort(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewHttpRouteRulesActionFaultInjectionPolicyAbort(c *Client, des, nw *HttpRouteRulesActionFaultInjectionPolicyAbort) *HttpRouteRulesActionFaultInjectionPolicyAbort {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for HttpRouteRulesActionFaultInjectionPolicyAbort while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewHttpRouteRulesActionFaultInjectionPolicyAbortSet(c *Client, des, nw []HttpRouteRulesActionFaultInjectionPolicyAbort) []HttpRouteRulesActionFaultInjectionPolicyAbort {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []HttpRouteRulesActionFaultInjectionPolicyAbort
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareHttpRouteRulesActionFaultInjectionPolicyAbortNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewHttpRouteRulesActionFaultInjectionPolicyAbort(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewHttpRouteRulesActionFaultInjectionPolicyAbortSlice(c *Client, des, nw []HttpRouteRulesActionFaultInjectionPolicyAbort) []HttpRouteRulesActionFaultInjectionPolicyAbort {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []HttpRouteRulesActionFaultInjectionPolicyAbort
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewHttpRouteRulesActionFaultInjectionPolicyAbort(c, &d, &n))
	}

	return items
}

func canonicalizeHttpRouteRulesActionRequestHeaderModifier(des, initial *HttpRouteRulesActionRequestHeaderModifier, opts ...dcl.ApplyOption) *HttpRouteRulesActionRequestHeaderModifier {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &HttpRouteRulesActionRequestHeaderModifier{}

	if dcl.IsZeroValue(des.Set) || (dcl.IsEmptyValueIndirect(des.Set) && dcl.IsEmptyValueIndirect(initial.Set)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Set = initial.Set
	} else {
		cDes.Set = des.Set
	}
	if dcl.IsZeroValue(des.Add) || (dcl.IsEmptyValueIndirect(des.Add) && dcl.IsEmptyValueIndirect(initial.Add)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Add = initial.Add
	} else {
		cDes.Add = des.Add
	}
	if dcl.StringArrayCanonicalize(des.Remove, initial.Remove) {
		cDes.Remove = initial.Remove
	} else {
		cDes.Remove = des.Remove
	}

	return cDes
}

func canonicalizeHttpRouteRulesActionRequestHeaderModifierSlice(des, initial []HttpRouteRulesActionRequestHeaderModifier, opts ...dcl.ApplyOption) []HttpRouteRulesActionRequestHeaderModifier {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]HttpRouteRulesActionRequestHeaderModifier, 0, len(des))
		for _, d := range des {
			cd := canonicalizeHttpRouteRulesActionRequestHeaderModifier(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]HttpRouteRulesActionRequestHeaderModifier, 0, len(des))
	for i, d := range des {
		cd := canonicalizeHttpRouteRulesActionRequestHeaderModifier(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewHttpRouteRulesActionRequestHeaderModifier(c *Client, des, nw *HttpRouteRulesActionRequestHeaderModifier) *HttpRouteRulesActionRequestHeaderModifier {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for HttpRouteRulesActionRequestHeaderModifier while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.Remove, nw.Remove) {
		nw.Remove = des.Remove
	}

	return nw
}

func canonicalizeNewHttpRouteRulesActionRequestHeaderModifierSet(c *Client, des, nw []HttpRouteRulesActionRequestHeaderModifier) []HttpRouteRulesActionRequestHeaderModifier {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []HttpRouteRulesActionRequestHeaderModifier
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareHttpRouteRulesActionRequestHeaderModifierNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewHttpRouteRulesActionRequestHeaderModifier(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewHttpRouteRulesActionRequestHeaderModifierSlice(c *Client, des, nw []HttpRouteRulesActionRequestHeaderModifier) []HttpRouteRulesActionRequestHeaderModifier {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []HttpRouteRulesActionRequestHeaderModifier
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewHttpRouteRulesActionRequestHeaderModifier(c, &d, &n))
	}

	return items
}

func canonicalizeHttpRouteRulesActionResponseHeaderModifier(des, initial *HttpRouteRulesActionResponseHeaderModifier, opts ...dcl.ApplyOption) *HttpRouteRulesActionResponseHeaderModifier {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &HttpRouteRulesActionResponseHeaderModifier{}

	if dcl.IsZeroValue(des.Set) || (dcl.IsEmptyValueIndirect(des.Set) && dcl.IsEmptyValueIndirect(initial.Set)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Set = initial.Set
	} else {
		cDes.Set = des.Set
	}
	if dcl.IsZeroValue(des.Add) || (dcl.IsEmptyValueIndirect(des.Add) && dcl.IsEmptyValueIndirect(initial.Add)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Add = initial.Add
	} else {
		cDes.Add = des.Add
	}
	if dcl.StringArrayCanonicalize(des.Remove, initial.Remove) {
		cDes.Remove = initial.Remove
	} else {
		cDes.Remove = des.Remove
	}

	return cDes
}

func canonicalizeHttpRouteRulesActionResponseHeaderModifierSlice(des, initial []HttpRouteRulesActionResponseHeaderModifier, opts ...dcl.ApplyOption) []HttpRouteRulesActionResponseHeaderModifier {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]HttpRouteRulesActionResponseHeaderModifier, 0, len(des))
		for _, d := range des {
			cd := canonicalizeHttpRouteRulesActionResponseHeaderModifier(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]HttpRouteRulesActionResponseHeaderModifier, 0, len(des))
	for i, d := range des {
		cd := canonicalizeHttpRouteRulesActionResponseHeaderModifier(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewHttpRouteRulesActionResponseHeaderModifier(c *Client, des, nw *HttpRouteRulesActionResponseHeaderModifier) *HttpRouteRulesActionResponseHeaderModifier {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for HttpRouteRulesActionResponseHeaderModifier while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.Remove, nw.Remove) {
		nw.Remove = des.Remove
	}

	return nw
}

func canonicalizeNewHttpRouteRulesActionResponseHeaderModifierSet(c *Client, des, nw []HttpRouteRulesActionResponseHeaderModifier) []HttpRouteRulesActionResponseHeaderModifier {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []HttpRouteRulesActionResponseHeaderModifier
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareHttpRouteRulesActionResponseHeaderModifierNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewHttpRouteRulesActionResponseHeaderModifier(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewHttpRouteRulesActionResponseHeaderModifierSlice(c *Client, des, nw []HttpRouteRulesActionResponseHeaderModifier) []HttpRouteRulesActionResponseHeaderModifier {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []HttpRouteRulesActionResponseHeaderModifier
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewHttpRouteRulesActionResponseHeaderModifier(c, &d, &n))
	}

	return items
}

func canonicalizeHttpRouteRulesActionUrlRewrite(des, initial *HttpRouteRulesActionUrlRewrite, opts ...dcl.ApplyOption) *HttpRouteRulesActionUrlRewrite {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &HttpRouteRulesActionUrlRewrite{}

	if dcl.StringCanonicalize(des.PathPrefixRewrite, initial.PathPrefixRewrite) || dcl.IsZeroValue(des.PathPrefixRewrite) {
		cDes.PathPrefixRewrite = initial.PathPrefixRewrite
	} else {
		cDes.PathPrefixRewrite = des.PathPrefixRewrite
	}
	if dcl.StringCanonicalize(des.HostRewrite, initial.HostRewrite) || dcl.IsZeroValue(des.HostRewrite) {
		cDes.HostRewrite = initial.HostRewrite
	} else {
		cDes.HostRewrite = des.HostRewrite
	}

	return cDes
}

func canonicalizeHttpRouteRulesActionUrlRewriteSlice(des, initial []HttpRouteRulesActionUrlRewrite, opts ...dcl.ApplyOption) []HttpRouteRulesActionUrlRewrite {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]HttpRouteRulesActionUrlRewrite, 0, len(des))
		for _, d := range des {
			cd := canonicalizeHttpRouteRulesActionUrlRewrite(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]HttpRouteRulesActionUrlRewrite, 0, len(des))
	for i, d := range des {
		cd := canonicalizeHttpRouteRulesActionUrlRewrite(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewHttpRouteRulesActionUrlRewrite(c *Client, des, nw *HttpRouteRulesActionUrlRewrite) *HttpRouteRulesActionUrlRewrite {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for HttpRouteRulesActionUrlRewrite while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.PathPrefixRewrite, nw.PathPrefixRewrite) {
		nw.PathPrefixRewrite = des.PathPrefixRewrite
	}
	if dcl.StringCanonicalize(des.HostRewrite, nw.HostRewrite) {
		nw.HostRewrite = des.HostRewrite
	}

	return nw
}

func canonicalizeNewHttpRouteRulesActionUrlRewriteSet(c *Client, des, nw []HttpRouteRulesActionUrlRewrite) []HttpRouteRulesActionUrlRewrite {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []HttpRouteRulesActionUrlRewrite
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareHttpRouteRulesActionUrlRewriteNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewHttpRouteRulesActionUrlRewrite(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewHttpRouteRulesActionUrlRewriteSlice(c *Client, des, nw []HttpRouteRulesActionUrlRewrite) []HttpRouteRulesActionUrlRewrite {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []HttpRouteRulesActionUrlRewrite
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewHttpRouteRulesActionUrlRewrite(c, &d, &n))
	}

	return items
}

func canonicalizeHttpRouteRulesActionRetryPolicy(des, initial *HttpRouteRulesActionRetryPolicy, opts ...dcl.ApplyOption) *HttpRouteRulesActionRetryPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &HttpRouteRulesActionRetryPolicy{}

	if dcl.StringArrayCanonicalize(des.RetryConditions, initial.RetryConditions) {
		cDes.RetryConditions = initial.RetryConditions
	} else {
		cDes.RetryConditions = des.RetryConditions
	}
	if dcl.IsZeroValue(des.NumRetries) || (dcl.IsEmptyValueIndirect(des.NumRetries) && dcl.IsEmptyValueIndirect(initial.NumRetries)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.NumRetries = initial.NumRetries
	} else {
		cDes.NumRetries = des.NumRetries
	}
	if dcl.StringCanonicalize(des.PerTryTimeout, initial.PerTryTimeout) || dcl.IsZeroValue(des.PerTryTimeout) {
		cDes.PerTryTimeout = initial.PerTryTimeout
	} else {
		cDes.PerTryTimeout = des.PerTryTimeout
	}

	return cDes
}

func canonicalizeHttpRouteRulesActionRetryPolicySlice(des, initial []HttpRouteRulesActionRetryPolicy, opts ...dcl.ApplyOption) []HttpRouteRulesActionRetryPolicy {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]HttpRouteRulesActionRetryPolicy, 0, len(des))
		for _, d := range des {
			cd := canonicalizeHttpRouteRulesActionRetryPolicy(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]HttpRouteRulesActionRetryPolicy, 0, len(des))
	for i, d := range des {
		cd := canonicalizeHttpRouteRulesActionRetryPolicy(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewHttpRouteRulesActionRetryPolicy(c *Client, des, nw *HttpRouteRulesActionRetryPolicy) *HttpRouteRulesActionRetryPolicy {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for HttpRouteRulesActionRetryPolicy while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.RetryConditions, nw.RetryConditions) {
		nw.RetryConditions = des.RetryConditions
	}
	if dcl.StringCanonicalize(des.PerTryTimeout, nw.PerTryTimeout) {
		nw.PerTryTimeout = des.PerTryTimeout
	}

	return nw
}

func canonicalizeNewHttpRouteRulesActionRetryPolicySet(c *Client, des, nw []HttpRouteRulesActionRetryPolicy) []HttpRouteRulesActionRetryPolicy {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []HttpRouteRulesActionRetryPolicy
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareHttpRouteRulesActionRetryPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewHttpRouteRulesActionRetryPolicy(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewHttpRouteRulesActionRetryPolicySlice(c *Client, des, nw []HttpRouteRulesActionRetryPolicy) []HttpRouteRulesActionRetryPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []HttpRouteRulesActionRetryPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewHttpRouteRulesActionRetryPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeHttpRouteRulesActionRequestMirrorPolicy(des, initial *HttpRouteRulesActionRequestMirrorPolicy, opts ...dcl.ApplyOption) *HttpRouteRulesActionRequestMirrorPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &HttpRouteRulesActionRequestMirrorPolicy{}

	cDes.Destination = canonicalizeHttpRouteRulesActionRequestMirrorPolicyDestination(des.Destination, initial.Destination, opts...)

	return cDes
}

func canonicalizeHttpRouteRulesActionRequestMirrorPolicySlice(des, initial []HttpRouteRulesActionRequestMirrorPolicy, opts ...dcl.ApplyOption) []HttpRouteRulesActionRequestMirrorPolicy {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]HttpRouteRulesActionRequestMirrorPolicy, 0, len(des))
		for _, d := range des {
			cd := canonicalizeHttpRouteRulesActionRequestMirrorPolicy(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]HttpRouteRulesActionRequestMirrorPolicy, 0, len(des))
	for i, d := range des {
		cd := canonicalizeHttpRouteRulesActionRequestMirrorPolicy(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewHttpRouteRulesActionRequestMirrorPolicy(c *Client, des, nw *HttpRouteRulesActionRequestMirrorPolicy) *HttpRouteRulesActionRequestMirrorPolicy {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for HttpRouteRulesActionRequestMirrorPolicy while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Destination = canonicalizeNewHttpRouteRulesActionRequestMirrorPolicyDestination(c, des.Destination, nw.Destination)

	return nw
}

func canonicalizeNewHttpRouteRulesActionRequestMirrorPolicySet(c *Client, des, nw []HttpRouteRulesActionRequestMirrorPolicy) []HttpRouteRulesActionRequestMirrorPolicy {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []HttpRouteRulesActionRequestMirrorPolicy
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareHttpRouteRulesActionRequestMirrorPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewHttpRouteRulesActionRequestMirrorPolicy(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewHttpRouteRulesActionRequestMirrorPolicySlice(c *Client, des, nw []HttpRouteRulesActionRequestMirrorPolicy) []HttpRouteRulesActionRequestMirrorPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []HttpRouteRulesActionRequestMirrorPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewHttpRouteRulesActionRequestMirrorPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeHttpRouteRulesActionRequestMirrorPolicyDestination(des, initial *HttpRouteRulesActionRequestMirrorPolicyDestination, opts ...dcl.ApplyOption) *HttpRouteRulesActionRequestMirrorPolicyDestination {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &HttpRouteRulesActionRequestMirrorPolicyDestination{}

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

func canonicalizeHttpRouteRulesActionRequestMirrorPolicyDestinationSlice(des, initial []HttpRouteRulesActionRequestMirrorPolicyDestination, opts ...dcl.ApplyOption) []HttpRouteRulesActionRequestMirrorPolicyDestination {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]HttpRouteRulesActionRequestMirrorPolicyDestination, 0, len(des))
		for _, d := range des {
			cd := canonicalizeHttpRouteRulesActionRequestMirrorPolicyDestination(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]HttpRouteRulesActionRequestMirrorPolicyDestination, 0, len(des))
	for i, d := range des {
		cd := canonicalizeHttpRouteRulesActionRequestMirrorPolicyDestination(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewHttpRouteRulesActionRequestMirrorPolicyDestination(c *Client, des, nw *HttpRouteRulesActionRequestMirrorPolicyDestination) *HttpRouteRulesActionRequestMirrorPolicyDestination {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for HttpRouteRulesActionRequestMirrorPolicyDestination while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewHttpRouteRulesActionRequestMirrorPolicyDestinationSet(c *Client, des, nw []HttpRouteRulesActionRequestMirrorPolicyDestination) []HttpRouteRulesActionRequestMirrorPolicyDestination {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []HttpRouteRulesActionRequestMirrorPolicyDestination
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareHttpRouteRulesActionRequestMirrorPolicyDestinationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewHttpRouteRulesActionRequestMirrorPolicyDestination(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewHttpRouteRulesActionRequestMirrorPolicyDestinationSlice(c *Client, des, nw []HttpRouteRulesActionRequestMirrorPolicyDestination) []HttpRouteRulesActionRequestMirrorPolicyDestination {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []HttpRouteRulesActionRequestMirrorPolicyDestination
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewHttpRouteRulesActionRequestMirrorPolicyDestination(c, &d, &n))
	}

	return items
}

func canonicalizeHttpRouteRulesActionCorsPolicy(des, initial *HttpRouteRulesActionCorsPolicy, opts ...dcl.ApplyOption) *HttpRouteRulesActionCorsPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &HttpRouteRulesActionCorsPolicy{}

	if dcl.StringArrayCanonicalize(des.AllowOrigins, initial.AllowOrigins) {
		cDes.AllowOrigins = initial.AllowOrigins
	} else {
		cDes.AllowOrigins = des.AllowOrigins
	}
	if dcl.StringArrayCanonicalize(des.AllowOriginRegexes, initial.AllowOriginRegexes) {
		cDes.AllowOriginRegexes = initial.AllowOriginRegexes
	} else {
		cDes.AllowOriginRegexes = des.AllowOriginRegexes
	}
	if dcl.StringArrayCanonicalize(des.AllowMethods, initial.AllowMethods) {
		cDes.AllowMethods = initial.AllowMethods
	} else {
		cDes.AllowMethods = des.AllowMethods
	}
	if dcl.StringArrayCanonicalize(des.AllowHeaders, initial.AllowHeaders) {
		cDes.AllowHeaders = initial.AllowHeaders
	} else {
		cDes.AllowHeaders = des.AllowHeaders
	}
	if dcl.StringArrayCanonicalize(des.ExposeHeaders, initial.ExposeHeaders) {
		cDes.ExposeHeaders = initial.ExposeHeaders
	} else {
		cDes.ExposeHeaders = des.ExposeHeaders
	}
	if dcl.StringCanonicalize(des.MaxAge, initial.MaxAge) || dcl.IsZeroValue(des.MaxAge) {
		cDes.MaxAge = initial.MaxAge
	} else {
		cDes.MaxAge = des.MaxAge
	}
	if dcl.BoolCanonicalize(des.AllowCredentials, initial.AllowCredentials) || dcl.IsZeroValue(des.AllowCredentials) {
		cDes.AllowCredentials = initial.AllowCredentials
	} else {
		cDes.AllowCredentials = des.AllowCredentials
	}
	if dcl.BoolCanonicalize(des.Disabled, initial.Disabled) || dcl.IsZeroValue(des.Disabled) {
		cDes.Disabled = initial.Disabled
	} else {
		cDes.Disabled = des.Disabled
	}

	return cDes
}

func canonicalizeHttpRouteRulesActionCorsPolicySlice(des, initial []HttpRouteRulesActionCorsPolicy, opts ...dcl.ApplyOption) []HttpRouteRulesActionCorsPolicy {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]HttpRouteRulesActionCorsPolicy, 0, len(des))
		for _, d := range des {
			cd := canonicalizeHttpRouteRulesActionCorsPolicy(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]HttpRouteRulesActionCorsPolicy, 0, len(des))
	for i, d := range des {
		cd := canonicalizeHttpRouteRulesActionCorsPolicy(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewHttpRouteRulesActionCorsPolicy(c *Client, des, nw *HttpRouteRulesActionCorsPolicy) *HttpRouteRulesActionCorsPolicy {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for HttpRouteRulesActionCorsPolicy while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.AllowOrigins, nw.AllowOrigins) {
		nw.AllowOrigins = des.AllowOrigins
	}
	if dcl.StringArrayCanonicalize(des.AllowOriginRegexes, nw.AllowOriginRegexes) {
		nw.AllowOriginRegexes = des.AllowOriginRegexes
	}
	if dcl.StringArrayCanonicalize(des.AllowMethods, nw.AllowMethods) {
		nw.AllowMethods = des.AllowMethods
	}
	if dcl.StringArrayCanonicalize(des.AllowHeaders, nw.AllowHeaders) {
		nw.AllowHeaders = des.AllowHeaders
	}
	if dcl.StringArrayCanonicalize(des.ExposeHeaders, nw.ExposeHeaders) {
		nw.ExposeHeaders = des.ExposeHeaders
	}
	if dcl.StringCanonicalize(des.MaxAge, nw.MaxAge) {
		nw.MaxAge = des.MaxAge
	}
	if dcl.BoolCanonicalize(des.AllowCredentials, nw.AllowCredentials) {
		nw.AllowCredentials = des.AllowCredentials
	}
	if dcl.BoolCanonicalize(des.Disabled, nw.Disabled) {
		nw.Disabled = des.Disabled
	}

	return nw
}

func canonicalizeNewHttpRouteRulesActionCorsPolicySet(c *Client, des, nw []HttpRouteRulesActionCorsPolicy) []HttpRouteRulesActionCorsPolicy {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []HttpRouteRulesActionCorsPolicy
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareHttpRouteRulesActionCorsPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewHttpRouteRulesActionCorsPolicy(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewHttpRouteRulesActionCorsPolicySlice(c *Client, des, nw []HttpRouteRulesActionCorsPolicy) []HttpRouteRulesActionCorsPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []HttpRouteRulesActionCorsPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewHttpRouteRulesActionCorsPolicy(c, &d, &n))
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
func diffHttpRoute(c *Client, desired, actual *HttpRoute, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Hostnames, actual.Hostnames, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Hostnames")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Meshes, actual.Meshes, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Meshes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Gateways, actual.Gateways, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Gateways")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Rules, actual.Rules, dcl.DiffInfo{ObjectFunction: compareHttpRouteRulesNewStyle, EmptyObject: EmptyHttpRouteRules, OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Rules")); len(ds) != 0 || err != nil {
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
func compareHttpRouteRulesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*HttpRouteRules)
	if !ok {
		desiredNotPointer, ok := d.(HttpRouteRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRules or *HttpRouteRules", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*HttpRouteRules)
	if !ok {
		actualNotPointer, ok := a.(HttpRouteRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRules", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Matches, actual.Matches, dcl.DiffInfo{ObjectFunction: compareHttpRouteRulesMatchesNewStyle, EmptyObject: EmptyHttpRouteRulesMatches, OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Matches")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Action, actual.Action, dcl.DiffInfo{ObjectFunction: compareHttpRouteRulesActionNewStyle, EmptyObject: EmptyHttpRouteRulesAction, OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Action")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareHttpRouteRulesMatchesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*HttpRouteRulesMatches)
	if !ok {
		desiredNotPointer, ok := d.(HttpRouteRulesMatches)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesMatches or *HttpRouteRulesMatches", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*HttpRouteRulesMatches)
	if !ok {
		actualNotPointer, ok := a.(HttpRouteRulesMatches)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesMatches", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.FullPathMatch, actual.FullPathMatch, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("FullPathMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PrefixMatch, actual.PrefixMatch, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("PrefixMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RegexMatch, actual.RegexMatch, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("RegexMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IgnoreCase, actual.IgnoreCase, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("IgnoreCase")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Headers, actual.Headers, dcl.DiffInfo{ObjectFunction: compareHttpRouteRulesMatchesHeadersNewStyle, EmptyObject: EmptyHttpRouteRulesMatchesHeaders, OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Headers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.QueryParameters, actual.QueryParameters, dcl.DiffInfo{ObjectFunction: compareHttpRouteRulesMatchesQueryParametersNewStyle, EmptyObject: EmptyHttpRouteRulesMatchesQueryParameters, OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("QueryParameters")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareHttpRouteRulesMatchesHeadersNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*HttpRouteRulesMatchesHeaders)
	if !ok {
		desiredNotPointer, ok := d.(HttpRouteRulesMatchesHeaders)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesMatchesHeaders or *HttpRouteRulesMatchesHeaders", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*HttpRouteRulesMatchesHeaders)
	if !ok {
		actualNotPointer, ok := a.(HttpRouteRulesMatchesHeaders)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesMatchesHeaders", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Header, actual.Header, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Header")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExactMatch, actual.ExactMatch, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("ExactMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RegexMatch, actual.RegexMatch, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("RegexMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PrefixMatch, actual.PrefixMatch, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("PrefixMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PresentMatch, actual.PresentMatch, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("PresentMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SuffixMatch, actual.SuffixMatch, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("SuffixMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RangeMatch, actual.RangeMatch, dcl.DiffInfo{ObjectFunction: compareHttpRouteRulesMatchesHeadersRangeMatchNewStyle, EmptyObject: EmptyHttpRouteRulesMatchesHeadersRangeMatch, OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("RangeMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InvertMatch, actual.InvertMatch, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("InvertMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareHttpRouteRulesMatchesHeadersRangeMatchNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*HttpRouteRulesMatchesHeadersRangeMatch)
	if !ok {
		desiredNotPointer, ok := d.(HttpRouteRulesMatchesHeadersRangeMatch)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesMatchesHeadersRangeMatch or *HttpRouteRulesMatchesHeadersRangeMatch", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*HttpRouteRulesMatchesHeadersRangeMatch)
	if !ok {
		actualNotPointer, ok := a.(HttpRouteRulesMatchesHeadersRangeMatch)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesMatchesHeadersRangeMatch", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Start, actual.Start, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Start")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.End, actual.End, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("End")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareHttpRouteRulesMatchesQueryParametersNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*HttpRouteRulesMatchesQueryParameters)
	if !ok {
		desiredNotPointer, ok := d.(HttpRouteRulesMatchesQueryParameters)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesMatchesQueryParameters or *HttpRouteRulesMatchesQueryParameters", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*HttpRouteRulesMatchesQueryParameters)
	if !ok {
		actualNotPointer, ok := a.(HttpRouteRulesMatchesQueryParameters)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesMatchesQueryParameters", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.QueryParameter, actual.QueryParameter, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("QueryParameter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExactMatch, actual.ExactMatch, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("ExactMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RegexMatch, actual.RegexMatch, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("RegexMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PresentMatch, actual.PresentMatch, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("PresentMatch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareHttpRouteRulesActionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*HttpRouteRulesAction)
	if !ok {
		desiredNotPointer, ok := d.(HttpRouteRulesAction)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesAction or *HttpRouteRulesAction", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*HttpRouteRulesAction)
	if !ok {
		actualNotPointer, ok := a.(HttpRouteRulesAction)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesAction", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Destinations, actual.Destinations, dcl.DiffInfo{ObjectFunction: compareHttpRouteRulesActionDestinationsNewStyle, EmptyObject: EmptyHttpRouteRulesActionDestinations, OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Destinations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Redirect, actual.Redirect, dcl.DiffInfo{ObjectFunction: compareHttpRouteRulesActionRedirectNewStyle, EmptyObject: EmptyHttpRouteRulesActionRedirect, OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Redirect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FaultInjectionPolicy, actual.FaultInjectionPolicy, dcl.DiffInfo{ObjectFunction: compareHttpRouteRulesActionFaultInjectionPolicyNewStyle, EmptyObject: EmptyHttpRouteRulesActionFaultInjectionPolicy, OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("FaultInjectionPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RequestHeaderModifier, actual.RequestHeaderModifier, dcl.DiffInfo{ObjectFunction: compareHttpRouteRulesActionRequestHeaderModifierNewStyle, EmptyObject: EmptyHttpRouteRulesActionRequestHeaderModifier, OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("RequestHeaderModifier")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResponseHeaderModifier, actual.ResponseHeaderModifier, dcl.DiffInfo{ObjectFunction: compareHttpRouteRulesActionResponseHeaderModifierNewStyle, EmptyObject: EmptyHttpRouteRulesActionResponseHeaderModifier, OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("ResponseHeaderModifier")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UrlRewrite, actual.UrlRewrite, dcl.DiffInfo{ObjectFunction: compareHttpRouteRulesActionUrlRewriteNewStyle, EmptyObject: EmptyHttpRouteRulesActionUrlRewrite, OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("UrlRewrite")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Timeout, actual.Timeout, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Timeout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RetryPolicy, actual.RetryPolicy, dcl.DiffInfo{ObjectFunction: compareHttpRouteRulesActionRetryPolicyNewStyle, EmptyObject: EmptyHttpRouteRulesActionRetryPolicy, OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("RetryPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RequestMirrorPolicy, actual.RequestMirrorPolicy, dcl.DiffInfo{ObjectFunction: compareHttpRouteRulesActionRequestMirrorPolicyNewStyle, EmptyObject: EmptyHttpRouteRulesActionRequestMirrorPolicy, OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("RequestMirrorPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CorsPolicy, actual.CorsPolicy, dcl.DiffInfo{ObjectFunction: compareHttpRouteRulesActionCorsPolicyNewStyle, EmptyObject: EmptyHttpRouteRulesActionCorsPolicy, OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("CorsPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareHttpRouteRulesActionDestinationsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*HttpRouteRulesActionDestinations)
	if !ok {
		desiredNotPointer, ok := d.(HttpRouteRulesActionDestinations)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesActionDestinations or *HttpRouteRulesActionDestinations", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*HttpRouteRulesActionDestinations)
	if !ok {
		actualNotPointer, ok := a.(HttpRouteRulesActionDestinations)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesActionDestinations", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Weight, actual.Weight, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Weight")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceName, actual.ServiceName, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("ServiceName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareHttpRouteRulesActionRedirectNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*HttpRouteRulesActionRedirect)
	if !ok {
		desiredNotPointer, ok := d.(HttpRouteRulesActionRedirect)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesActionRedirect or *HttpRouteRulesActionRedirect", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*HttpRouteRulesActionRedirect)
	if !ok {
		actualNotPointer, ok := a.(HttpRouteRulesActionRedirect)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesActionRedirect", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HostRedirect, actual.HostRedirect, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("HostRedirect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PathRedirect, actual.PathRedirect, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("PathRedirect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PrefixRewrite, actual.PrefixRewrite, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("PrefixRewrite")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResponseCode, actual.ResponseCode, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("ResponseCode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HttpsRedirect, actual.HttpsRedirect, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("HttpsRedirect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StripQuery, actual.StripQuery, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("StripQuery")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PortRedirect, actual.PortRedirect, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("PortRedirect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareHttpRouteRulesActionFaultInjectionPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*HttpRouteRulesActionFaultInjectionPolicy)
	if !ok {
		desiredNotPointer, ok := d.(HttpRouteRulesActionFaultInjectionPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesActionFaultInjectionPolicy or *HttpRouteRulesActionFaultInjectionPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*HttpRouteRulesActionFaultInjectionPolicy)
	if !ok {
		actualNotPointer, ok := a.(HttpRouteRulesActionFaultInjectionPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesActionFaultInjectionPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Delay, actual.Delay, dcl.DiffInfo{ObjectFunction: compareHttpRouteRulesActionFaultInjectionPolicyDelayNewStyle, EmptyObject: EmptyHttpRouteRulesActionFaultInjectionPolicyDelay, OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Delay")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Abort, actual.Abort, dcl.DiffInfo{ObjectFunction: compareHttpRouteRulesActionFaultInjectionPolicyAbortNewStyle, EmptyObject: EmptyHttpRouteRulesActionFaultInjectionPolicyAbort, OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Abort")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareHttpRouteRulesActionFaultInjectionPolicyDelayNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*HttpRouteRulesActionFaultInjectionPolicyDelay)
	if !ok {
		desiredNotPointer, ok := d.(HttpRouteRulesActionFaultInjectionPolicyDelay)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesActionFaultInjectionPolicyDelay or *HttpRouteRulesActionFaultInjectionPolicyDelay", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*HttpRouteRulesActionFaultInjectionPolicyDelay)
	if !ok {
		actualNotPointer, ok := a.(HttpRouteRulesActionFaultInjectionPolicyDelay)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesActionFaultInjectionPolicyDelay", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.FixedDelay, actual.FixedDelay, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("FixedDelay")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percentage, actual.Percentage, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Percentage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareHttpRouteRulesActionFaultInjectionPolicyAbortNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*HttpRouteRulesActionFaultInjectionPolicyAbort)
	if !ok {
		desiredNotPointer, ok := d.(HttpRouteRulesActionFaultInjectionPolicyAbort)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesActionFaultInjectionPolicyAbort or *HttpRouteRulesActionFaultInjectionPolicyAbort", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*HttpRouteRulesActionFaultInjectionPolicyAbort)
	if !ok {
		actualNotPointer, ok := a.(HttpRouteRulesActionFaultInjectionPolicyAbort)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesActionFaultInjectionPolicyAbort", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HttpStatus, actual.HttpStatus, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("HttpStatus")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percentage, actual.Percentage, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Percentage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareHttpRouteRulesActionRequestHeaderModifierNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*HttpRouteRulesActionRequestHeaderModifier)
	if !ok {
		desiredNotPointer, ok := d.(HttpRouteRulesActionRequestHeaderModifier)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesActionRequestHeaderModifier or *HttpRouteRulesActionRequestHeaderModifier", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*HttpRouteRulesActionRequestHeaderModifier)
	if !ok {
		actualNotPointer, ok := a.(HttpRouteRulesActionRequestHeaderModifier)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesActionRequestHeaderModifier", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Set, actual.Set, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Set")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Add, actual.Add, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Add")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Remove, actual.Remove, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Remove")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareHttpRouteRulesActionResponseHeaderModifierNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*HttpRouteRulesActionResponseHeaderModifier)
	if !ok {
		desiredNotPointer, ok := d.(HttpRouteRulesActionResponseHeaderModifier)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesActionResponseHeaderModifier or *HttpRouteRulesActionResponseHeaderModifier", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*HttpRouteRulesActionResponseHeaderModifier)
	if !ok {
		actualNotPointer, ok := a.(HttpRouteRulesActionResponseHeaderModifier)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesActionResponseHeaderModifier", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Set, actual.Set, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Set")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Add, actual.Add, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Add")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Remove, actual.Remove, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Remove")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareHttpRouteRulesActionUrlRewriteNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*HttpRouteRulesActionUrlRewrite)
	if !ok {
		desiredNotPointer, ok := d.(HttpRouteRulesActionUrlRewrite)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesActionUrlRewrite or *HttpRouteRulesActionUrlRewrite", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*HttpRouteRulesActionUrlRewrite)
	if !ok {
		actualNotPointer, ok := a.(HttpRouteRulesActionUrlRewrite)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesActionUrlRewrite", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PathPrefixRewrite, actual.PathPrefixRewrite, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("PathPrefixRewrite")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HostRewrite, actual.HostRewrite, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("HostRewrite")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareHttpRouteRulesActionRetryPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*HttpRouteRulesActionRetryPolicy)
	if !ok {
		desiredNotPointer, ok := d.(HttpRouteRulesActionRetryPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesActionRetryPolicy or *HttpRouteRulesActionRetryPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*HttpRouteRulesActionRetryPolicy)
	if !ok {
		actualNotPointer, ok := a.(HttpRouteRulesActionRetryPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesActionRetryPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RetryConditions, actual.RetryConditions, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("RetryConditions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NumRetries, actual.NumRetries, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("NumRetries")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PerTryTimeout, actual.PerTryTimeout, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("PerTryTimeout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareHttpRouteRulesActionRequestMirrorPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*HttpRouteRulesActionRequestMirrorPolicy)
	if !ok {
		desiredNotPointer, ok := d.(HttpRouteRulesActionRequestMirrorPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesActionRequestMirrorPolicy or *HttpRouteRulesActionRequestMirrorPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*HttpRouteRulesActionRequestMirrorPolicy)
	if !ok {
		actualNotPointer, ok := a.(HttpRouteRulesActionRequestMirrorPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesActionRequestMirrorPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Destination, actual.Destination, dcl.DiffInfo{ObjectFunction: compareHttpRouteRulesActionRequestMirrorPolicyDestinationNewStyle, EmptyObject: EmptyHttpRouteRulesActionRequestMirrorPolicyDestination, OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Destination")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareHttpRouteRulesActionRequestMirrorPolicyDestinationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*HttpRouteRulesActionRequestMirrorPolicyDestination)
	if !ok {
		desiredNotPointer, ok := d.(HttpRouteRulesActionRequestMirrorPolicyDestination)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesActionRequestMirrorPolicyDestination or *HttpRouteRulesActionRequestMirrorPolicyDestination", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*HttpRouteRulesActionRequestMirrorPolicyDestination)
	if !ok {
		actualNotPointer, ok := a.(HttpRouteRulesActionRequestMirrorPolicyDestination)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesActionRequestMirrorPolicyDestination", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Weight, actual.Weight, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Weight")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceName, actual.ServiceName, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("ServiceName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareHttpRouteRulesActionCorsPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*HttpRouteRulesActionCorsPolicy)
	if !ok {
		desiredNotPointer, ok := d.(HttpRouteRulesActionCorsPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesActionCorsPolicy or *HttpRouteRulesActionCorsPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*HttpRouteRulesActionCorsPolicy)
	if !ok {
		actualNotPointer, ok := a.(HttpRouteRulesActionCorsPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a HttpRouteRulesActionCorsPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AllowOrigins, actual.AllowOrigins, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("AllowOrigins")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowOriginRegexes, actual.AllowOriginRegexes, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("AllowOriginRegexes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowMethods, actual.AllowMethods, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("AllowMethods")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowHeaders, actual.AllowHeaders, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("AllowHeaders")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExposeHeaders, actual.ExposeHeaders, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("ExposeHeaders")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxAge, actual.MaxAge, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("MaxAge")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowCredentials, actual.AllowCredentials, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("AllowCredentials")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Disabled, actual.Disabled, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateHttpRouteUpdateHttpRouteOperation")}, fn.AddNest("Disabled")); len(ds) != 0 || err != nil {
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
func (r *HttpRoute) urlNormalized() *HttpRoute {
	normalized := dcl.Copy(*r).(HttpRoute)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	return &normalized
}

func (r *HttpRoute) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateHttpRoute" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/httpRoutes/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the HttpRoute resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *HttpRoute) marshal(c *Client) ([]byte, error) {
	m, err := expandHttpRoute(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling HttpRoute: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalHttpRoute decodes JSON responses into the HttpRoute resource schema.
func unmarshalHttpRoute(b []byte, c *Client, res *HttpRoute) (*HttpRoute, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapHttpRoute(m, c, res)
}

func unmarshalMapHttpRoute(m map[string]interface{}, c *Client, res *HttpRoute) (*HttpRoute, error) {

	flattened := flattenHttpRoute(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandHttpRoute expands HttpRoute into a JSON request object.
func expandHttpRoute(c *Client, f *HttpRoute) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("projects/%s/locations/global/httpRoutes/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v := f.Hostnames; v != nil {
		m["hostnames"] = v
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
	if v, err := expandHttpRouteRulesSlice(c, f.Rules, res); err != nil {
		return nil, fmt.Errorf("error expanding Rules into rules: %w", err)
	} else if v != nil {
		m["rules"] = v
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

// flattenHttpRoute flattens HttpRoute from a JSON request object into the
// HttpRoute type.
func flattenHttpRoute(c *Client, i interface{}, res *HttpRoute) *HttpRoute {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &HttpRoute{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.Hostnames = dcl.FlattenStringSlice(m["hostnames"])
	resultRes.Meshes = dcl.FlattenStringSlice(m["meshes"])
	resultRes.Gateways = dcl.FlattenStringSlice(m["gateways"])
	resultRes.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	resultRes.Rules = flattenHttpRouteRulesSlice(c, m["rules"], res)
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])
	resultRes.SelfLink = dcl.FlattenString(m["selfLink"])

	return resultRes
}

// expandHttpRouteRulesMap expands the contents of HttpRouteRules into a JSON
// request object.
func expandHttpRouteRulesMap(c *Client, f map[string]HttpRouteRules, res *HttpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandHttpRouteRules(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandHttpRouteRulesSlice expands the contents of HttpRouteRules into a JSON
// request object.
func expandHttpRouteRulesSlice(c *Client, f []HttpRouteRules, res *HttpRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandHttpRouteRules(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenHttpRouteRulesMap flattens the contents of HttpRouteRules from a JSON
// response object.
func flattenHttpRouteRulesMap(c *Client, i interface{}, res *HttpRoute) map[string]HttpRouteRules {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]HttpRouteRules{}
	}

	if len(a) == 0 {
		return map[string]HttpRouteRules{}
	}

	items := make(map[string]HttpRouteRules)
	for k, item := range a {
		items[k] = *flattenHttpRouteRules(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenHttpRouteRulesSlice flattens the contents of HttpRouteRules from a JSON
// response object.
func flattenHttpRouteRulesSlice(c *Client, i interface{}, res *HttpRoute) []HttpRouteRules {
	a, ok := i.([]interface{})
	if !ok {
		return []HttpRouteRules{}
	}

	if len(a) == 0 {
		return []HttpRouteRules{}
	}

	items := make([]HttpRouteRules, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHttpRouteRules(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandHttpRouteRules expands an instance of HttpRouteRules into a JSON
// request object.
func expandHttpRouteRules(c *Client, f *HttpRouteRules, res *HttpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandHttpRouteRulesMatchesSlice(c, f.Matches, res); err != nil {
		return nil, fmt.Errorf("error expanding Matches into matches: %w", err)
	} else if v != nil {
		m["matches"] = v
	}
	if v, err := expandHttpRouteRulesAction(c, f.Action, res); err != nil {
		return nil, fmt.Errorf("error expanding Action into action: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["action"] = v
	}

	return m, nil
}

// flattenHttpRouteRules flattens an instance of HttpRouteRules from a JSON
// response object.
func flattenHttpRouteRules(c *Client, i interface{}, res *HttpRoute) *HttpRouteRules {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &HttpRouteRules{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyHttpRouteRules
	}
	r.Matches = flattenHttpRouteRulesMatchesSlice(c, m["matches"], res)
	r.Action = flattenHttpRouteRulesAction(c, m["action"], res)

	return r
}

// expandHttpRouteRulesMatchesMap expands the contents of HttpRouteRulesMatches into a JSON
// request object.
func expandHttpRouteRulesMatchesMap(c *Client, f map[string]HttpRouteRulesMatches, res *HttpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandHttpRouteRulesMatches(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandHttpRouteRulesMatchesSlice expands the contents of HttpRouteRulesMatches into a JSON
// request object.
func expandHttpRouteRulesMatchesSlice(c *Client, f []HttpRouteRulesMatches, res *HttpRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandHttpRouteRulesMatches(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenHttpRouteRulesMatchesMap flattens the contents of HttpRouteRulesMatches from a JSON
// response object.
func flattenHttpRouteRulesMatchesMap(c *Client, i interface{}, res *HttpRoute) map[string]HttpRouteRulesMatches {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]HttpRouteRulesMatches{}
	}

	if len(a) == 0 {
		return map[string]HttpRouteRulesMatches{}
	}

	items := make(map[string]HttpRouteRulesMatches)
	for k, item := range a {
		items[k] = *flattenHttpRouteRulesMatches(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenHttpRouteRulesMatchesSlice flattens the contents of HttpRouteRulesMatches from a JSON
// response object.
func flattenHttpRouteRulesMatchesSlice(c *Client, i interface{}, res *HttpRoute) []HttpRouteRulesMatches {
	a, ok := i.([]interface{})
	if !ok {
		return []HttpRouteRulesMatches{}
	}

	if len(a) == 0 {
		return []HttpRouteRulesMatches{}
	}

	items := make([]HttpRouteRulesMatches, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHttpRouteRulesMatches(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandHttpRouteRulesMatches expands an instance of HttpRouteRulesMatches into a JSON
// request object.
func expandHttpRouteRulesMatches(c *Client, f *HttpRouteRulesMatches, res *HttpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.FullPathMatch; !dcl.IsEmptyValueIndirect(v) {
		m["fullPathMatch"] = v
	}
	if v := f.PrefixMatch; !dcl.IsEmptyValueIndirect(v) {
		m["prefixMatch"] = v
	}
	if v := f.RegexMatch; !dcl.IsEmptyValueIndirect(v) {
		m["regexMatch"] = v
	}
	if v := f.IgnoreCase; !dcl.IsEmptyValueIndirect(v) {
		m["ignoreCase"] = v
	}
	if v, err := expandHttpRouteRulesMatchesHeadersSlice(c, f.Headers, res); err != nil {
		return nil, fmt.Errorf("error expanding Headers into headers: %w", err)
	} else if v != nil {
		m["headers"] = v
	}
	if v, err := expandHttpRouteRulesMatchesQueryParametersSlice(c, f.QueryParameters, res); err != nil {
		return nil, fmt.Errorf("error expanding QueryParameters into queryParameters: %w", err)
	} else if v != nil {
		m["queryParameters"] = v
	}

	return m, nil
}

// flattenHttpRouteRulesMatches flattens an instance of HttpRouteRulesMatches from a JSON
// response object.
func flattenHttpRouteRulesMatches(c *Client, i interface{}, res *HttpRoute) *HttpRouteRulesMatches {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &HttpRouteRulesMatches{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyHttpRouteRulesMatches
	}
	r.FullPathMatch = dcl.FlattenString(m["fullPathMatch"])
	r.PrefixMatch = dcl.FlattenString(m["prefixMatch"])
	r.RegexMatch = dcl.FlattenString(m["regexMatch"])
	r.IgnoreCase = dcl.FlattenBool(m["ignoreCase"])
	r.Headers = flattenHttpRouteRulesMatchesHeadersSlice(c, m["headers"], res)
	r.QueryParameters = flattenHttpRouteRulesMatchesQueryParametersSlice(c, m["queryParameters"], res)

	return r
}

// expandHttpRouteRulesMatchesHeadersMap expands the contents of HttpRouteRulesMatchesHeaders into a JSON
// request object.
func expandHttpRouteRulesMatchesHeadersMap(c *Client, f map[string]HttpRouteRulesMatchesHeaders, res *HttpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandHttpRouteRulesMatchesHeaders(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandHttpRouteRulesMatchesHeadersSlice expands the contents of HttpRouteRulesMatchesHeaders into a JSON
// request object.
func expandHttpRouteRulesMatchesHeadersSlice(c *Client, f []HttpRouteRulesMatchesHeaders, res *HttpRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandHttpRouteRulesMatchesHeaders(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenHttpRouteRulesMatchesHeadersMap flattens the contents of HttpRouteRulesMatchesHeaders from a JSON
// response object.
func flattenHttpRouteRulesMatchesHeadersMap(c *Client, i interface{}, res *HttpRoute) map[string]HttpRouteRulesMatchesHeaders {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]HttpRouteRulesMatchesHeaders{}
	}

	if len(a) == 0 {
		return map[string]HttpRouteRulesMatchesHeaders{}
	}

	items := make(map[string]HttpRouteRulesMatchesHeaders)
	for k, item := range a {
		items[k] = *flattenHttpRouteRulesMatchesHeaders(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenHttpRouteRulesMatchesHeadersSlice flattens the contents of HttpRouteRulesMatchesHeaders from a JSON
// response object.
func flattenHttpRouteRulesMatchesHeadersSlice(c *Client, i interface{}, res *HttpRoute) []HttpRouteRulesMatchesHeaders {
	a, ok := i.([]interface{})
	if !ok {
		return []HttpRouteRulesMatchesHeaders{}
	}

	if len(a) == 0 {
		return []HttpRouteRulesMatchesHeaders{}
	}

	items := make([]HttpRouteRulesMatchesHeaders, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHttpRouteRulesMatchesHeaders(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandHttpRouteRulesMatchesHeaders expands an instance of HttpRouteRulesMatchesHeaders into a JSON
// request object.
func expandHttpRouteRulesMatchesHeaders(c *Client, f *HttpRouteRulesMatchesHeaders, res *HttpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Header; !dcl.IsEmptyValueIndirect(v) {
		m["header"] = v
	}
	if v := f.ExactMatch; !dcl.IsEmptyValueIndirect(v) {
		m["exactMatch"] = v
	}
	if v := f.RegexMatch; !dcl.IsEmptyValueIndirect(v) {
		m["regexMatch"] = v
	}
	if v := f.PrefixMatch; !dcl.IsEmptyValueIndirect(v) {
		m["prefixMatch"] = v
	}
	if v := f.PresentMatch; !dcl.IsEmptyValueIndirect(v) {
		m["presentMatch"] = v
	}
	if v := f.SuffixMatch; !dcl.IsEmptyValueIndirect(v) {
		m["suffixMatch"] = v
	}
	if v, err := expandHttpRouteRulesMatchesHeadersRangeMatch(c, f.RangeMatch, res); err != nil {
		return nil, fmt.Errorf("error expanding RangeMatch into rangeMatch: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["rangeMatch"] = v
	}
	if v := f.InvertMatch; !dcl.IsEmptyValueIndirect(v) {
		m["invertMatch"] = v
	}

	return m, nil
}

// flattenHttpRouteRulesMatchesHeaders flattens an instance of HttpRouteRulesMatchesHeaders from a JSON
// response object.
func flattenHttpRouteRulesMatchesHeaders(c *Client, i interface{}, res *HttpRoute) *HttpRouteRulesMatchesHeaders {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &HttpRouteRulesMatchesHeaders{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyHttpRouteRulesMatchesHeaders
	}
	r.Header = dcl.FlattenString(m["header"])
	r.ExactMatch = dcl.FlattenString(m["exactMatch"])
	r.RegexMatch = dcl.FlattenString(m["regexMatch"])
	r.PrefixMatch = dcl.FlattenString(m["prefixMatch"])
	r.PresentMatch = dcl.FlattenBool(m["presentMatch"])
	r.SuffixMatch = dcl.FlattenString(m["suffixMatch"])
	r.RangeMatch = flattenHttpRouteRulesMatchesHeadersRangeMatch(c, m["rangeMatch"], res)
	r.InvertMatch = dcl.FlattenBool(m["invertMatch"])

	return r
}

// expandHttpRouteRulesMatchesHeadersRangeMatchMap expands the contents of HttpRouteRulesMatchesHeadersRangeMatch into a JSON
// request object.
func expandHttpRouteRulesMatchesHeadersRangeMatchMap(c *Client, f map[string]HttpRouteRulesMatchesHeadersRangeMatch, res *HttpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandHttpRouteRulesMatchesHeadersRangeMatch(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandHttpRouteRulesMatchesHeadersRangeMatchSlice expands the contents of HttpRouteRulesMatchesHeadersRangeMatch into a JSON
// request object.
func expandHttpRouteRulesMatchesHeadersRangeMatchSlice(c *Client, f []HttpRouteRulesMatchesHeadersRangeMatch, res *HttpRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandHttpRouteRulesMatchesHeadersRangeMatch(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenHttpRouteRulesMatchesHeadersRangeMatchMap flattens the contents of HttpRouteRulesMatchesHeadersRangeMatch from a JSON
// response object.
func flattenHttpRouteRulesMatchesHeadersRangeMatchMap(c *Client, i interface{}, res *HttpRoute) map[string]HttpRouteRulesMatchesHeadersRangeMatch {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]HttpRouteRulesMatchesHeadersRangeMatch{}
	}

	if len(a) == 0 {
		return map[string]HttpRouteRulesMatchesHeadersRangeMatch{}
	}

	items := make(map[string]HttpRouteRulesMatchesHeadersRangeMatch)
	for k, item := range a {
		items[k] = *flattenHttpRouteRulesMatchesHeadersRangeMatch(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenHttpRouteRulesMatchesHeadersRangeMatchSlice flattens the contents of HttpRouteRulesMatchesHeadersRangeMatch from a JSON
// response object.
func flattenHttpRouteRulesMatchesHeadersRangeMatchSlice(c *Client, i interface{}, res *HttpRoute) []HttpRouteRulesMatchesHeadersRangeMatch {
	a, ok := i.([]interface{})
	if !ok {
		return []HttpRouteRulesMatchesHeadersRangeMatch{}
	}

	if len(a) == 0 {
		return []HttpRouteRulesMatchesHeadersRangeMatch{}
	}

	items := make([]HttpRouteRulesMatchesHeadersRangeMatch, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHttpRouteRulesMatchesHeadersRangeMatch(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandHttpRouteRulesMatchesHeadersRangeMatch expands an instance of HttpRouteRulesMatchesHeadersRangeMatch into a JSON
// request object.
func expandHttpRouteRulesMatchesHeadersRangeMatch(c *Client, f *HttpRouteRulesMatchesHeadersRangeMatch, res *HttpRoute) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Start; !dcl.IsEmptyValueIndirect(v) {
		m["start"] = v
	}
	if v := f.End; !dcl.IsEmptyValueIndirect(v) {
		m["end"] = v
	}

	return m, nil
}

// flattenHttpRouteRulesMatchesHeadersRangeMatch flattens an instance of HttpRouteRulesMatchesHeadersRangeMatch from a JSON
// response object.
func flattenHttpRouteRulesMatchesHeadersRangeMatch(c *Client, i interface{}, res *HttpRoute) *HttpRouteRulesMatchesHeadersRangeMatch {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &HttpRouteRulesMatchesHeadersRangeMatch{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyHttpRouteRulesMatchesHeadersRangeMatch
	}
	r.Start = dcl.FlattenInteger(m["start"])
	r.End = dcl.FlattenInteger(m["end"])

	return r
}

// expandHttpRouteRulesMatchesQueryParametersMap expands the contents of HttpRouteRulesMatchesQueryParameters into a JSON
// request object.
func expandHttpRouteRulesMatchesQueryParametersMap(c *Client, f map[string]HttpRouteRulesMatchesQueryParameters, res *HttpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandHttpRouteRulesMatchesQueryParameters(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandHttpRouteRulesMatchesQueryParametersSlice expands the contents of HttpRouteRulesMatchesQueryParameters into a JSON
// request object.
func expandHttpRouteRulesMatchesQueryParametersSlice(c *Client, f []HttpRouteRulesMatchesQueryParameters, res *HttpRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandHttpRouteRulesMatchesQueryParameters(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenHttpRouteRulesMatchesQueryParametersMap flattens the contents of HttpRouteRulesMatchesQueryParameters from a JSON
// response object.
func flattenHttpRouteRulesMatchesQueryParametersMap(c *Client, i interface{}, res *HttpRoute) map[string]HttpRouteRulesMatchesQueryParameters {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]HttpRouteRulesMatchesQueryParameters{}
	}

	if len(a) == 0 {
		return map[string]HttpRouteRulesMatchesQueryParameters{}
	}

	items := make(map[string]HttpRouteRulesMatchesQueryParameters)
	for k, item := range a {
		items[k] = *flattenHttpRouteRulesMatchesQueryParameters(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenHttpRouteRulesMatchesQueryParametersSlice flattens the contents of HttpRouteRulesMatchesQueryParameters from a JSON
// response object.
func flattenHttpRouteRulesMatchesQueryParametersSlice(c *Client, i interface{}, res *HttpRoute) []HttpRouteRulesMatchesQueryParameters {
	a, ok := i.([]interface{})
	if !ok {
		return []HttpRouteRulesMatchesQueryParameters{}
	}

	if len(a) == 0 {
		return []HttpRouteRulesMatchesQueryParameters{}
	}

	items := make([]HttpRouteRulesMatchesQueryParameters, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHttpRouteRulesMatchesQueryParameters(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandHttpRouteRulesMatchesQueryParameters expands an instance of HttpRouteRulesMatchesQueryParameters into a JSON
// request object.
func expandHttpRouteRulesMatchesQueryParameters(c *Client, f *HttpRouteRulesMatchesQueryParameters, res *HttpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.QueryParameter; !dcl.IsEmptyValueIndirect(v) {
		m["queryParameter"] = v
	}
	if v := f.ExactMatch; !dcl.IsEmptyValueIndirect(v) {
		m["exactMatch"] = v
	}
	if v := f.RegexMatch; !dcl.IsEmptyValueIndirect(v) {
		m["regexMatch"] = v
	}
	if v := f.PresentMatch; !dcl.IsEmptyValueIndirect(v) {
		m["presentMatch"] = v
	}

	return m, nil
}

// flattenHttpRouteRulesMatchesQueryParameters flattens an instance of HttpRouteRulesMatchesQueryParameters from a JSON
// response object.
func flattenHttpRouteRulesMatchesQueryParameters(c *Client, i interface{}, res *HttpRoute) *HttpRouteRulesMatchesQueryParameters {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &HttpRouteRulesMatchesQueryParameters{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyHttpRouteRulesMatchesQueryParameters
	}
	r.QueryParameter = dcl.FlattenString(m["queryParameter"])
	r.ExactMatch = dcl.FlattenString(m["exactMatch"])
	r.RegexMatch = dcl.FlattenString(m["regexMatch"])
	r.PresentMatch = dcl.FlattenBool(m["presentMatch"])

	return r
}

// expandHttpRouteRulesActionMap expands the contents of HttpRouteRulesAction into a JSON
// request object.
func expandHttpRouteRulesActionMap(c *Client, f map[string]HttpRouteRulesAction, res *HttpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandHttpRouteRulesAction(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandHttpRouteRulesActionSlice expands the contents of HttpRouteRulesAction into a JSON
// request object.
func expandHttpRouteRulesActionSlice(c *Client, f []HttpRouteRulesAction, res *HttpRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandHttpRouteRulesAction(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenHttpRouteRulesActionMap flattens the contents of HttpRouteRulesAction from a JSON
// response object.
func flattenHttpRouteRulesActionMap(c *Client, i interface{}, res *HttpRoute) map[string]HttpRouteRulesAction {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]HttpRouteRulesAction{}
	}

	if len(a) == 0 {
		return map[string]HttpRouteRulesAction{}
	}

	items := make(map[string]HttpRouteRulesAction)
	for k, item := range a {
		items[k] = *flattenHttpRouteRulesAction(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenHttpRouteRulesActionSlice flattens the contents of HttpRouteRulesAction from a JSON
// response object.
func flattenHttpRouteRulesActionSlice(c *Client, i interface{}, res *HttpRoute) []HttpRouteRulesAction {
	a, ok := i.([]interface{})
	if !ok {
		return []HttpRouteRulesAction{}
	}

	if len(a) == 0 {
		return []HttpRouteRulesAction{}
	}

	items := make([]HttpRouteRulesAction, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHttpRouteRulesAction(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandHttpRouteRulesAction expands an instance of HttpRouteRulesAction into a JSON
// request object.
func expandHttpRouteRulesAction(c *Client, f *HttpRouteRulesAction, res *HttpRoute) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandHttpRouteRulesActionDestinationsSlice(c, f.Destinations, res); err != nil {
		return nil, fmt.Errorf("error expanding Destinations into destinations: %w", err)
	} else if v != nil {
		m["destinations"] = v
	}
	if v, err := expandHttpRouteRulesActionRedirect(c, f.Redirect, res); err != nil {
		return nil, fmt.Errorf("error expanding Redirect into redirect: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["redirect"] = v
	}
	if v, err := expandHttpRouteRulesActionFaultInjectionPolicy(c, f.FaultInjectionPolicy, res); err != nil {
		return nil, fmt.Errorf("error expanding FaultInjectionPolicy into faultInjectionPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["faultInjectionPolicy"] = v
	}
	if v, err := expandHttpRouteRulesActionRequestHeaderModifier(c, f.RequestHeaderModifier, res); err != nil {
		return nil, fmt.Errorf("error expanding RequestHeaderModifier into requestHeaderModifier: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["requestHeaderModifier"] = v
	}
	if v, err := expandHttpRouteRulesActionResponseHeaderModifier(c, f.ResponseHeaderModifier, res); err != nil {
		return nil, fmt.Errorf("error expanding ResponseHeaderModifier into responseHeaderModifier: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["responseHeaderModifier"] = v
	}
	if v, err := expandHttpRouteRulesActionUrlRewrite(c, f.UrlRewrite, res); err != nil {
		return nil, fmt.Errorf("error expanding UrlRewrite into urlRewrite: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["urlRewrite"] = v
	}
	if v := f.Timeout; !dcl.IsEmptyValueIndirect(v) {
		m["timeout"] = v
	}
	if v, err := expandHttpRouteRulesActionRetryPolicy(c, f.RetryPolicy, res); err != nil {
		return nil, fmt.Errorf("error expanding RetryPolicy into retryPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["retryPolicy"] = v
	}
	if v, err := expandHttpRouteRulesActionRequestMirrorPolicy(c, f.RequestMirrorPolicy, res); err != nil {
		return nil, fmt.Errorf("error expanding RequestMirrorPolicy into requestMirrorPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["requestMirrorPolicy"] = v
	}
	if v, err := expandHttpRouteRulesActionCorsPolicy(c, f.CorsPolicy, res); err != nil {
		return nil, fmt.Errorf("error expanding CorsPolicy into corsPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["corsPolicy"] = v
	}

	return m, nil
}

// flattenHttpRouteRulesAction flattens an instance of HttpRouteRulesAction from a JSON
// response object.
func flattenHttpRouteRulesAction(c *Client, i interface{}, res *HttpRoute) *HttpRouteRulesAction {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &HttpRouteRulesAction{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyHttpRouteRulesAction
	}
	r.Destinations = flattenHttpRouteRulesActionDestinationsSlice(c, m["destinations"], res)
	r.Redirect = flattenHttpRouteRulesActionRedirect(c, m["redirect"], res)
	r.FaultInjectionPolicy = flattenHttpRouteRulesActionFaultInjectionPolicy(c, m["faultInjectionPolicy"], res)
	r.RequestHeaderModifier = flattenHttpRouteRulesActionRequestHeaderModifier(c, m["requestHeaderModifier"], res)
	r.ResponseHeaderModifier = flattenHttpRouteRulesActionResponseHeaderModifier(c, m["responseHeaderModifier"], res)
	r.UrlRewrite = flattenHttpRouteRulesActionUrlRewrite(c, m["urlRewrite"], res)
	r.Timeout = dcl.FlattenString(m["timeout"])
	r.RetryPolicy = flattenHttpRouteRulesActionRetryPolicy(c, m["retryPolicy"], res)
	r.RequestMirrorPolicy = flattenHttpRouteRulesActionRequestMirrorPolicy(c, m["requestMirrorPolicy"], res)
	r.CorsPolicy = flattenHttpRouteRulesActionCorsPolicy(c, m["corsPolicy"], res)

	return r
}

// expandHttpRouteRulesActionDestinationsMap expands the contents of HttpRouteRulesActionDestinations into a JSON
// request object.
func expandHttpRouteRulesActionDestinationsMap(c *Client, f map[string]HttpRouteRulesActionDestinations, res *HttpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandHttpRouteRulesActionDestinations(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandHttpRouteRulesActionDestinationsSlice expands the contents of HttpRouteRulesActionDestinations into a JSON
// request object.
func expandHttpRouteRulesActionDestinationsSlice(c *Client, f []HttpRouteRulesActionDestinations, res *HttpRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandHttpRouteRulesActionDestinations(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenHttpRouteRulesActionDestinationsMap flattens the contents of HttpRouteRulesActionDestinations from a JSON
// response object.
func flattenHttpRouteRulesActionDestinationsMap(c *Client, i interface{}, res *HttpRoute) map[string]HttpRouteRulesActionDestinations {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]HttpRouteRulesActionDestinations{}
	}

	if len(a) == 0 {
		return map[string]HttpRouteRulesActionDestinations{}
	}

	items := make(map[string]HttpRouteRulesActionDestinations)
	for k, item := range a {
		items[k] = *flattenHttpRouteRulesActionDestinations(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenHttpRouteRulesActionDestinationsSlice flattens the contents of HttpRouteRulesActionDestinations from a JSON
// response object.
func flattenHttpRouteRulesActionDestinationsSlice(c *Client, i interface{}, res *HttpRoute) []HttpRouteRulesActionDestinations {
	a, ok := i.([]interface{})
	if !ok {
		return []HttpRouteRulesActionDestinations{}
	}

	if len(a) == 0 {
		return []HttpRouteRulesActionDestinations{}
	}

	items := make([]HttpRouteRulesActionDestinations, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHttpRouteRulesActionDestinations(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandHttpRouteRulesActionDestinations expands an instance of HttpRouteRulesActionDestinations into a JSON
// request object.
func expandHttpRouteRulesActionDestinations(c *Client, f *HttpRouteRulesActionDestinations, res *HttpRoute) (map[string]interface{}, error) {
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

// flattenHttpRouteRulesActionDestinations flattens an instance of HttpRouteRulesActionDestinations from a JSON
// response object.
func flattenHttpRouteRulesActionDestinations(c *Client, i interface{}, res *HttpRoute) *HttpRouteRulesActionDestinations {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &HttpRouteRulesActionDestinations{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyHttpRouteRulesActionDestinations
	}
	r.Weight = dcl.FlattenInteger(m["weight"])
	r.ServiceName = dcl.FlattenString(m["serviceName"])

	return r
}

// expandHttpRouteRulesActionRedirectMap expands the contents of HttpRouteRulesActionRedirect into a JSON
// request object.
func expandHttpRouteRulesActionRedirectMap(c *Client, f map[string]HttpRouteRulesActionRedirect, res *HttpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandHttpRouteRulesActionRedirect(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandHttpRouteRulesActionRedirectSlice expands the contents of HttpRouteRulesActionRedirect into a JSON
// request object.
func expandHttpRouteRulesActionRedirectSlice(c *Client, f []HttpRouteRulesActionRedirect, res *HttpRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandHttpRouteRulesActionRedirect(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenHttpRouteRulesActionRedirectMap flattens the contents of HttpRouteRulesActionRedirect from a JSON
// response object.
func flattenHttpRouteRulesActionRedirectMap(c *Client, i interface{}, res *HttpRoute) map[string]HttpRouteRulesActionRedirect {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]HttpRouteRulesActionRedirect{}
	}

	if len(a) == 0 {
		return map[string]HttpRouteRulesActionRedirect{}
	}

	items := make(map[string]HttpRouteRulesActionRedirect)
	for k, item := range a {
		items[k] = *flattenHttpRouteRulesActionRedirect(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenHttpRouteRulesActionRedirectSlice flattens the contents of HttpRouteRulesActionRedirect from a JSON
// response object.
func flattenHttpRouteRulesActionRedirectSlice(c *Client, i interface{}, res *HttpRoute) []HttpRouteRulesActionRedirect {
	a, ok := i.([]interface{})
	if !ok {
		return []HttpRouteRulesActionRedirect{}
	}

	if len(a) == 0 {
		return []HttpRouteRulesActionRedirect{}
	}

	items := make([]HttpRouteRulesActionRedirect, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHttpRouteRulesActionRedirect(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandHttpRouteRulesActionRedirect expands an instance of HttpRouteRulesActionRedirect into a JSON
// request object.
func expandHttpRouteRulesActionRedirect(c *Client, f *HttpRouteRulesActionRedirect, res *HttpRoute) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.HostRedirect; !dcl.IsEmptyValueIndirect(v) {
		m["hostRedirect"] = v
	}
	if v := f.PathRedirect; !dcl.IsEmptyValueIndirect(v) {
		m["pathRedirect"] = v
	}
	if v := f.PrefixRewrite; !dcl.IsEmptyValueIndirect(v) {
		m["prefixRewrite"] = v
	}
	if v := f.ResponseCode; !dcl.IsEmptyValueIndirect(v) {
		m["responseCode"] = v
	}
	if v := f.HttpsRedirect; !dcl.IsEmptyValueIndirect(v) {
		m["httpsRedirect"] = v
	}
	if v := f.StripQuery; !dcl.IsEmptyValueIndirect(v) {
		m["stripQuery"] = v
	}
	if v := f.PortRedirect; !dcl.IsEmptyValueIndirect(v) {
		m["portRedirect"] = v
	}

	return m, nil
}

// flattenHttpRouteRulesActionRedirect flattens an instance of HttpRouteRulesActionRedirect from a JSON
// response object.
func flattenHttpRouteRulesActionRedirect(c *Client, i interface{}, res *HttpRoute) *HttpRouteRulesActionRedirect {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &HttpRouteRulesActionRedirect{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyHttpRouteRulesActionRedirect
	}
	r.HostRedirect = dcl.FlattenString(m["hostRedirect"])
	r.PathRedirect = dcl.FlattenString(m["pathRedirect"])
	r.PrefixRewrite = dcl.FlattenString(m["prefixRewrite"])
	r.ResponseCode = flattenHttpRouteRulesActionRedirectResponseCodeEnum(m["responseCode"])
	r.HttpsRedirect = dcl.FlattenBool(m["httpsRedirect"])
	r.StripQuery = dcl.FlattenBool(m["stripQuery"])
	r.PortRedirect = dcl.FlattenInteger(m["portRedirect"])

	return r
}

// expandHttpRouteRulesActionFaultInjectionPolicyMap expands the contents of HttpRouteRulesActionFaultInjectionPolicy into a JSON
// request object.
func expandHttpRouteRulesActionFaultInjectionPolicyMap(c *Client, f map[string]HttpRouteRulesActionFaultInjectionPolicy, res *HttpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandHttpRouteRulesActionFaultInjectionPolicy(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandHttpRouteRulesActionFaultInjectionPolicySlice expands the contents of HttpRouteRulesActionFaultInjectionPolicy into a JSON
// request object.
func expandHttpRouteRulesActionFaultInjectionPolicySlice(c *Client, f []HttpRouteRulesActionFaultInjectionPolicy, res *HttpRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandHttpRouteRulesActionFaultInjectionPolicy(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenHttpRouteRulesActionFaultInjectionPolicyMap flattens the contents of HttpRouteRulesActionFaultInjectionPolicy from a JSON
// response object.
func flattenHttpRouteRulesActionFaultInjectionPolicyMap(c *Client, i interface{}, res *HttpRoute) map[string]HttpRouteRulesActionFaultInjectionPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]HttpRouteRulesActionFaultInjectionPolicy{}
	}

	if len(a) == 0 {
		return map[string]HttpRouteRulesActionFaultInjectionPolicy{}
	}

	items := make(map[string]HttpRouteRulesActionFaultInjectionPolicy)
	for k, item := range a {
		items[k] = *flattenHttpRouteRulesActionFaultInjectionPolicy(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenHttpRouteRulesActionFaultInjectionPolicySlice flattens the contents of HttpRouteRulesActionFaultInjectionPolicy from a JSON
// response object.
func flattenHttpRouteRulesActionFaultInjectionPolicySlice(c *Client, i interface{}, res *HttpRoute) []HttpRouteRulesActionFaultInjectionPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []HttpRouteRulesActionFaultInjectionPolicy{}
	}

	if len(a) == 0 {
		return []HttpRouteRulesActionFaultInjectionPolicy{}
	}

	items := make([]HttpRouteRulesActionFaultInjectionPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHttpRouteRulesActionFaultInjectionPolicy(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandHttpRouteRulesActionFaultInjectionPolicy expands an instance of HttpRouteRulesActionFaultInjectionPolicy into a JSON
// request object.
func expandHttpRouteRulesActionFaultInjectionPolicy(c *Client, f *HttpRouteRulesActionFaultInjectionPolicy, res *HttpRoute) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandHttpRouteRulesActionFaultInjectionPolicyDelay(c, f.Delay, res); err != nil {
		return nil, fmt.Errorf("error expanding Delay into delay: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["delay"] = v
	}
	if v, err := expandHttpRouteRulesActionFaultInjectionPolicyAbort(c, f.Abort, res); err != nil {
		return nil, fmt.Errorf("error expanding Abort into abort: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["abort"] = v
	}

	return m, nil
}

// flattenHttpRouteRulesActionFaultInjectionPolicy flattens an instance of HttpRouteRulesActionFaultInjectionPolicy from a JSON
// response object.
func flattenHttpRouteRulesActionFaultInjectionPolicy(c *Client, i interface{}, res *HttpRoute) *HttpRouteRulesActionFaultInjectionPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &HttpRouteRulesActionFaultInjectionPolicy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyHttpRouteRulesActionFaultInjectionPolicy
	}
	r.Delay = flattenHttpRouteRulesActionFaultInjectionPolicyDelay(c, m["delay"], res)
	r.Abort = flattenHttpRouteRulesActionFaultInjectionPolicyAbort(c, m["abort"], res)

	return r
}

// expandHttpRouteRulesActionFaultInjectionPolicyDelayMap expands the contents of HttpRouteRulesActionFaultInjectionPolicyDelay into a JSON
// request object.
func expandHttpRouteRulesActionFaultInjectionPolicyDelayMap(c *Client, f map[string]HttpRouteRulesActionFaultInjectionPolicyDelay, res *HttpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandHttpRouteRulesActionFaultInjectionPolicyDelay(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandHttpRouteRulesActionFaultInjectionPolicyDelaySlice expands the contents of HttpRouteRulesActionFaultInjectionPolicyDelay into a JSON
// request object.
func expandHttpRouteRulesActionFaultInjectionPolicyDelaySlice(c *Client, f []HttpRouteRulesActionFaultInjectionPolicyDelay, res *HttpRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandHttpRouteRulesActionFaultInjectionPolicyDelay(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenHttpRouteRulesActionFaultInjectionPolicyDelayMap flattens the contents of HttpRouteRulesActionFaultInjectionPolicyDelay from a JSON
// response object.
func flattenHttpRouteRulesActionFaultInjectionPolicyDelayMap(c *Client, i interface{}, res *HttpRoute) map[string]HttpRouteRulesActionFaultInjectionPolicyDelay {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]HttpRouteRulesActionFaultInjectionPolicyDelay{}
	}

	if len(a) == 0 {
		return map[string]HttpRouteRulesActionFaultInjectionPolicyDelay{}
	}

	items := make(map[string]HttpRouteRulesActionFaultInjectionPolicyDelay)
	for k, item := range a {
		items[k] = *flattenHttpRouteRulesActionFaultInjectionPolicyDelay(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenHttpRouteRulesActionFaultInjectionPolicyDelaySlice flattens the contents of HttpRouteRulesActionFaultInjectionPolicyDelay from a JSON
// response object.
func flattenHttpRouteRulesActionFaultInjectionPolicyDelaySlice(c *Client, i interface{}, res *HttpRoute) []HttpRouteRulesActionFaultInjectionPolicyDelay {
	a, ok := i.([]interface{})
	if !ok {
		return []HttpRouteRulesActionFaultInjectionPolicyDelay{}
	}

	if len(a) == 0 {
		return []HttpRouteRulesActionFaultInjectionPolicyDelay{}
	}

	items := make([]HttpRouteRulesActionFaultInjectionPolicyDelay, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHttpRouteRulesActionFaultInjectionPolicyDelay(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandHttpRouteRulesActionFaultInjectionPolicyDelay expands an instance of HttpRouteRulesActionFaultInjectionPolicyDelay into a JSON
// request object.
func expandHttpRouteRulesActionFaultInjectionPolicyDelay(c *Client, f *HttpRouteRulesActionFaultInjectionPolicyDelay, res *HttpRoute) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.FixedDelay; !dcl.IsEmptyValueIndirect(v) {
		m["fixedDelay"] = v
	}
	if v := f.Percentage; !dcl.IsEmptyValueIndirect(v) {
		m["percentage"] = v
	}

	return m, nil
}

// flattenHttpRouteRulesActionFaultInjectionPolicyDelay flattens an instance of HttpRouteRulesActionFaultInjectionPolicyDelay from a JSON
// response object.
func flattenHttpRouteRulesActionFaultInjectionPolicyDelay(c *Client, i interface{}, res *HttpRoute) *HttpRouteRulesActionFaultInjectionPolicyDelay {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &HttpRouteRulesActionFaultInjectionPolicyDelay{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyHttpRouteRulesActionFaultInjectionPolicyDelay
	}
	r.FixedDelay = dcl.FlattenString(m["fixedDelay"])
	r.Percentage = dcl.FlattenInteger(m["percentage"])

	return r
}

// expandHttpRouteRulesActionFaultInjectionPolicyAbortMap expands the contents of HttpRouteRulesActionFaultInjectionPolicyAbort into a JSON
// request object.
func expandHttpRouteRulesActionFaultInjectionPolicyAbortMap(c *Client, f map[string]HttpRouteRulesActionFaultInjectionPolicyAbort, res *HttpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandHttpRouteRulesActionFaultInjectionPolicyAbort(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandHttpRouteRulesActionFaultInjectionPolicyAbortSlice expands the contents of HttpRouteRulesActionFaultInjectionPolicyAbort into a JSON
// request object.
func expandHttpRouteRulesActionFaultInjectionPolicyAbortSlice(c *Client, f []HttpRouteRulesActionFaultInjectionPolicyAbort, res *HttpRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandHttpRouteRulesActionFaultInjectionPolicyAbort(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenHttpRouteRulesActionFaultInjectionPolicyAbortMap flattens the contents of HttpRouteRulesActionFaultInjectionPolicyAbort from a JSON
// response object.
func flattenHttpRouteRulesActionFaultInjectionPolicyAbortMap(c *Client, i interface{}, res *HttpRoute) map[string]HttpRouteRulesActionFaultInjectionPolicyAbort {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]HttpRouteRulesActionFaultInjectionPolicyAbort{}
	}

	if len(a) == 0 {
		return map[string]HttpRouteRulesActionFaultInjectionPolicyAbort{}
	}

	items := make(map[string]HttpRouteRulesActionFaultInjectionPolicyAbort)
	for k, item := range a {
		items[k] = *flattenHttpRouteRulesActionFaultInjectionPolicyAbort(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenHttpRouteRulesActionFaultInjectionPolicyAbortSlice flattens the contents of HttpRouteRulesActionFaultInjectionPolicyAbort from a JSON
// response object.
func flattenHttpRouteRulesActionFaultInjectionPolicyAbortSlice(c *Client, i interface{}, res *HttpRoute) []HttpRouteRulesActionFaultInjectionPolicyAbort {
	a, ok := i.([]interface{})
	if !ok {
		return []HttpRouteRulesActionFaultInjectionPolicyAbort{}
	}

	if len(a) == 0 {
		return []HttpRouteRulesActionFaultInjectionPolicyAbort{}
	}

	items := make([]HttpRouteRulesActionFaultInjectionPolicyAbort, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHttpRouteRulesActionFaultInjectionPolicyAbort(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandHttpRouteRulesActionFaultInjectionPolicyAbort expands an instance of HttpRouteRulesActionFaultInjectionPolicyAbort into a JSON
// request object.
func expandHttpRouteRulesActionFaultInjectionPolicyAbort(c *Client, f *HttpRouteRulesActionFaultInjectionPolicyAbort, res *HttpRoute) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.HttpStatus; !dcl.IsEmptyValueIndirect(v) {
		m["httpStatus"] = v
	}
	if v := f.Percentage; !dcl.IsEmptyValueIndirect(v) {
		m["percentage"] = v
	}

	return m, nil
}

// flattenHttpRouteRulesActionFaultInjectionPolicyAbort flattens an instance of HttpRouteRulesActionFaultInjectionPolicyAbort from a JSON
// response object.
func flattenHttpRouteRulesActionFaultInjectionPolicyAbort(c *Client, i interface{}, res *HttpRoute) *HttpRouteRulesActionFaultInjectionPolicyAbort {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &HttpRouteRulesActionFaultInjectionPolicyAbort{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyHttpRouteRulesActionFaultInjectionPolicyAbort
	}
	r.HttpStatus = dcl.FlattenInteger(m["httpStatus"])
	r.Percentage = dcl.FlattenInteger(m["percentage"])

	return r
}

// expandHttpRouteRulesActionRequestHeaderModifierMap expands the contents of HttpRouteRulesActionRequestHeaderModifier into a JSON
// request object.
func expandHttpRouteRulesActionRequestHeaderModifierMap(c *Client, f map[string]HttpRouteRulesActionRequestHeaderModifier, res *HttpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandHttpRouteRulesActionRequestHeaderModifier(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandHttpRouteRulesActionRequestHeaderModifierSlice expands the contents of HttpRouteRulesActionRequestHeaderModifier into a JSON
// request object.
func expandHttpRouteRulesActionRequestHeaderModifierSlice(c *Client, f []HttpRouteRulesActionRequestHeaderModifier, res *HttpRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandHttpRouteRulesActionRequestHeaderModifier(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenHttpRouteRulesActionRequestHeaderModifierMap flattens the contents of HttpRouteRulesActionRequestHeaderModifier from a JSON
// response object.
func flattenHttpRouteRulesActionRequestHeaderModifierMap(c *Client, i interface{}, res *HttpRoute) map[string]HttpRouteRulesActionRequestHeaderModifier {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]HttpRouteRulesActionRequestHeaderModifier{}
	}

	if len(a) == 0 {
		return map[string]HttpRouteRulesActionRequestHeaderModifier{}
	}

	items := make(map[string]HttpRouteRulesActionRequestHeaderModifier)
	for k, item := range a {
		items[k] = *flattenHttpRouteRulesActionRequestHeaderModifier(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenHttpRouteRulesActionRequestHeaderModifierSlice flattens the contents of HttpRouteRulesActionRequestHeaderModifier from a JSON
// response object.
func flattenHttpRouteRulesActionRequestHeaderModifierSlice(c *Client, i interface{}, res *HttpRoute) []HttpRouteRulesActionRequestHeaderModifier {
	a, ok := i.([]interface{})
	if !ok {
		return []HttpRouteRulesActionRequestHeaderModifier{}
	}

	if len(a) == 0 {
		return []HttpRouteRulesActionRequestHeaderModifier{}
	}

	items := make([]HttpRouteRulesActionRequestHeaderModifier, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHttpRouteRulesActionRequestHeaderModifier(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandHttpRouteRulesActionRequestHeaderModifier expands an instance of HttpRouteRulesActionRequestHeaderModifier into a JSON
// request object.
func expandHttpRouteRulesActionRequestHeaderModifier(c *Client, f *HttpRouteRulesActionRequestHeaderModifier, res *HttpRoute) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Set; !dcl.IsEmptyValueIndirect(v) {
		m["set"] = v
	}
	if v := f.Add; !dcl.IsEmptyValueIndirect(v) {
		m["add"] = v
	}
	if v := f.Remove; v != nil {
		m["remove"] = v
	}

	return m, nil
}

// flattenHttpRouteRulesActionRequestHeaderModifier flattens an instance of HttpRouteRulesActionRequestHeaderModifier from a JSON
// response object.
func flattenHttpRouteRulesActionRequestHeaderModifier(c *Client, i interface{}, res *HttpRoute) *HttpRouteRulesActionRequestHeaderModifier {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &HttpRouteRulesActionRequestHeaderModifier{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyHttpRouteRulesActionRequestHeaderModifier
	}
	r.Set = dcl.FlattenKeyValuePairs(m["set"])
	r.Add = dcl.FlattenKeyValuePairs(m["add"])
	r.Remove = dcl.FlattenStringSlice(m["remove"])

	return r
}

// expandHttpRouteRulesActionResponseHeaderModifierMap expands the contents of HttpRouteRulesActionResponseHeaderModifier into a JSON
// request object.
func expandHttpRouteRulesActionResponseHeaderModifierMap(c *Client, f map[string]HttpRouteRulesActionResponseHeaderModifier, res *HttpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandHttpRouteRulesActionResponseHeaderModifier(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandHttpRouteRulesActionResponseHeaderModifierSlice expands the contents of HttpRouteRulesActionResponseHeaderModifier into a JSON
// request object.
func expandHttpRouteRulesActionResponseHeaderModifierSlice(c *Client, f []HttpRouteRulesActionResponseHeaderModifier, res *HttpRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandHttpRouteRulesActionResponseHeaderModifier(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenHttpRouteRulesActionResponseHeaderModifierMap flattens the contents of HttpRouteRulesActionResponseHeaderModifier from a JSON
// response object.
func flattenHttpRouteRulesActionResponseHeaderModifierMap(c *Client, i interface{}, res *HttpRoute) map[string]HttpRouteRulesActionResponseHeaderModifier {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]HttpRouteRulesActionResponseHeaderModifier{}
	}

	if len(a) == 0 {
		return map[string]HttpRouteRulesActionResponseHeaderModifier{}
	}

	items := make(map[string]HttpRouteRulesActionResponseHeaderModifier)
	for k, item := range a {
		items[k] = *flattenHttpRouteRulesActionResponseHeaderModifier(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenHttpRouteRulesActionResponseHeaderModifierSlice flattens the contents of HttpRouteRulesActionResponseHeaderModifier from a JSON
// response object.
func flattenHttpRouteRulesActionResponseHeaderModifierSlice(c *Client, i interface{}, res *HttpRoute) []HttpRouteRulesActionResponseHeaderModifier {
	a, ok := i.([]interface{})
	if !ok {
		return []HttpRouteRulesActionResponseHeaderModifier{}
	}

	if len(a) == 0 {
		return []HttpRouteRulesActionResponseHeaderModifier{}
	}

	items := make([]HttpRouteRulesActionResponseHeaderModifier, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHttpRouteRulesActionResponseHeaderModifier(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandHttpRouteRulesActionResponseHeaderModifier expands an instance of HttpRouteRulesActionResponseHeaderModifier into a JSON
// request object.
func expandHttpRouteRulesActionResponseHeaderModifier(c *Client, f *HttpRouteRulesActionResponseHeaderModifier, res *HttpRoute) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Set; !dcl.IsEmptyValueIndirect(v) {
		m["set"] = v
	}
	if v := f.Add; !dcl.IsEmptyValueIndirect(v) {
		m["add"] = v
	}
	if v := f.Remove; v != nil {
		m["remove"] = v
	}

	return m, nil
}

// flattenHttpRouteRulesActionResponseHeaderModifier flattens an instance of HttpRouteRulesActionResponseHeaderModifier from a JSON
// response object.
func flattenHttpRouteRulesActionResponseHeaderModifier(c *Client, i interface{}, res *HttpRoute) *HttpRouteRulesActionResponseHeaderModifier {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &HttpRouteRulesActionResponseHeaderModifier{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyHttpRouteRulesActionResponseHeaderModifier
	}
	r.Set = dcl.FlattenKeyValuePairs(m["set"])
	r.Add = dcl.FlattenKeyValuePairs(m["add"])
	r.Remove = dcl.FlattenStringSlice(m["remove"])

	return r
}

// expandHttpRouteRulesActionUrlRewriteMap expands the contents of HttpRouteRulesActionUrlRewrite into a JSON
// request object.
func expandHttpRouteRulesActionUrlRewriteMap(c *Client, f map[string]HttpRouteRulesActionUrlRewrite, res *HttpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandHttpRouteRulesActionUrlRewrite(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandHttpRouteRulesActionUrlRewriteSlice expands the contents of HttpRouteRulesActionUrlRewrite into a JSON
// request object.
func expandHttpRouteRulesActionUrlRewriteSlice(c *Client, f []HttpRouteRulesActionUrlRewrite, res *HttpRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandHttpRouteRulesActionUrlRewrite(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenHttpRouteRulesActionUrlRewriteMap flattens the contents of HttpRouteRulesActionUrlRewrite from a JSON
// response object.
func flattenHttpRouteRulesActionUrlRewriteMap(c *Client, i interface{}, res *HttpRoute) map[string]HttpRouteRulesActionUrlRewrite {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]HttpRouteRulesActionUrlRewrite{}
	}

	if len(a) == 0 {
		return map[string]HttpRouteRulesActionUrlRewrite{}
	}

	items := make(map[string]HttpRouteRulesActionUrlRewrite)
	for k, item := range a {
		items[k] = *flattenHttpRouteRulesActionUrlRewrite(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenHttpRouteRulesActionUrlRewriteSlice flattens the contents of HttpRouteRulesActionUrlRewrite from a JSON
// response object.
func flattenHttpRouteRulesActionUrlRewriteSlice(c *Client, i interface{}, res *HttpRoute) []HttpRouteRulesActionUrlRewrite {
	a, ok := i.([]interface{})
	if !ok {
		return []HttpRouteRulesActionUrlRewrite{}
	}

	if len(a) == 0 {
		return []HttpRouteRulesActionUrlRewrite{}
	}

	items := make([]HttpRouteRulesActionUrlRewrite, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHttpRouteRulesActionUrlRewrite(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandHttpRouteRulesActionUrlRewrite expands an instance of HttpRouteRulesActionUrlRewrite into a JSON
// request object.
func expandHttpRouteRulesActionUrlRewrite(c *Client, f *HttpRouteRulesActionUrlRewrite, res *HttpRoute) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.PathPrefixRewrite; !dcl.IsEmptyValueIndirect(v) {
		m["pathPrefixRewrite"] = v
	}
	if v := f.HostRewrite; !dcl.IsEmptyValueIndirect(v) {
		m["hostRewrite"] = v
	}

	return m, nil
}

// flattenHttpRouteRulesActionUrlRewrite flattens an instance of HttpRouteRulesActionUrlRewrite from a JSON
// response object.
func flattenHttpRouteRulesActionUrlRewrite(c *Client, i interface{}, res *HttpRoute) *HttpRouteRulesActionUrlRewrite {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &HttpRouteRulesActionUrlRewrite{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyHttpRouteRulesActionUrlRewrite
	}
	r.PathPrefixRewrite = dcl.FlattenString(m["pathPrefixRewrite"])
	r.HostRewrite = dcl.FlattenString(m["hostRewrite"])

	return r
}

// expandHttpRouteRulesActionRetryPolicyMap expands the contents of HttpRouteRulesActionRetryPolicy into a JSON
// request object.
func expandHttpRouteRulesActionRetryPolicyMap(c *Client, f map[string]HttpRouteRulesActionRetryPolicy, res *HttpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandHttpRouteRulesActionRetryPolicy(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandHttpRouteRulesActionRetryPolicySlice expands the contents of HttpRouteRulesActionRetryPolicy into a JSON
// request object.
func expandHttpRouteRulesActionRetryPolicySlice(c *Client, f []HttpRouteRulesActionRetryPolicy, res *HttpRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandHttpRouteRulesActionRetryPolicy(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenHttpRouteRulesActionRetryPolicyMap flattens the contents of HttpRouteRulesActionRetryPolicy from a JSON
// response object.
func flattenHttpRouteRulesActionRetryPolicyMap(c *Client, i interface{}, res *HttpRoute) map[string]HttpRouteRulesActionRetryPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]HttpRouteRulesActionRetryPolicy{}
	}

	if len(a) == 0 {
		return map[string]HttpRouteRulesActionRetryPolicy{}
	}

	items := make(map[string]HttpRouteRulesActionRetryPolicy)
	for k, item := range a {
		items[k] = *flattenHttpRouteRulesActionRetryPolicy(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenHttpRouteRulesActionRetryPolicySlice flattens the contents of HttpRouteRulesActionRetryPolicy from a JSON
// response object.
func flattenHttpRouteRulesActionRetryPolicySlice(c *Client, i interface{}, res *HttpRoute) []HttpRouteRulesActionRetryPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []HttpRouteRulesActionRetryPolicy{}
	}

	if len(a) == 0 {
		return []HttpRouteRulesActionRetryPolicy{}
	}

	items := make([]HttpRouteRulesActionRetryPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHttpRouteRulesActionRetryPolicy(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandHttpRouteRulesActionRetryPolicy expands an instance of HttpRouteRulesActionRetryPolicy into a JSON
// request object.
func expandHttpRouteRulesActionRetryPolicy(c *Client, f *HttpRouteRulesActionRetryPolicy, res *HttpRoute) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RetryConditions; v != nil {
		m["retryConditions"] = v
	}
	if v := f.NumRetries; !dcl.IsEmptyValueIndirect(v) {
		m["numRetries"] = v
	}
	if v := f.PerTryTimeout; !dcl.IsEmptyValueIndirect(v) {
		m["perTryTimeout"] = v
	}

	return m, nil
}

// flattenHttpRouteRulesActionRetryPolicy flattens an instance of HttpRouteRulesActionRetryPolicy from a JSON
// response object.
func flattenHttpRouteRulesActionRetryPolicy(c *Client, i interface{}, res *HttpRoute) *HttpRouteRulesActionRetryPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &HttpRouteRulesActionRetryPolicy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyHttpRouteRulesActionRetryPolicy
	}
	r.RetryConditions = dcl.FlattenStringSlice(m["retryConditions"])
	r.NumRetries = dcl.FlattenInteger(m["numRetries"])
	r.PerTryTimeout = dcl.FlattenString(m["perTryTimeout"])

	return r
}

// expandHttpRouteRulesActionRequestMirrorPolicyMap expands the contents of HttpRouteRulesActionRequestMirrorPolicy into a JSON
// request object.
func expandHttpRouteRulesActionRequestMirrorPolicyMap(c *Client, f map[string]HttpRouteRulesActionRequestMirrorPolicy, res *HttpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandHttpRouteRulesActionRequestMirrorPolicy(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandHttpRouteRulesActionRequestMirrorPolicySlice expands the contents of HttpRouteRulesActionRequestMirrorPolicy into a JSON
// request object.
func expandHttpRouteRulesActionRequestMirrorPolicySlice(c *Client, f []HttpRouteRulesActionRequestMirrorPolicy, res *HttpRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandHttpRouteRulesActionRequestMirrorPolicy(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenHttpRouteRulesActionRequestMirrorPolicyMap flattens the contents of HttpRouteRulesActionRequestMirrorPolicy from a JSON
// response object.
func flattenHttpRouteRulesActionRequestMirrorPolicyMap(c *Client, i interface{}, res *HttpRoute) map[string]HttpRouteRulesActionRequestMirrorPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]HttpRouteRulesActionRequestMirrorPolicy{}
	}

	if len(a) == 0 {
		return map[string]HttpRouteRulesActionRequestMirrorPolicy{}
	}

	items := make(map[string]HttpRouteRulesActionRequestMirrorPolicy)
	for k, item := range a {
		items[k] = *flattenHttpRouteRulesActionRequestMirrorPolicy(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenHttpRouteRulesActionRequestMirrorPolicySlice flattens the contents of HttpRouteRulesActionRequestMirrorPolicy from a JSON
// response object.
func flattenHttpRouteRulesActionRequestMirrorPolicySlice(c *Client, i interface{}, res *HttpRoute) []HttpRouteRulesActionRequestMirrorPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []HttpRouteRulesActionRequestMirrorPolicy{}
	}

	if len(a) == 0 {
		return []HttpRouteRulesActionRequestMirrorPolicy{}
	}

	items := make([]HttpRouteRulesActionRequestMirrorPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHttpRouteRulesActionRequestMirrorPolicy(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandHttpRouteRulesActionRequestMirrorPolicy expands an instance of HttpRouteRulesActionRequestMirrorPolicy into a JSON
// request object.
func expandHttpRouteRulesActionRequestMirrorPolicy(c *Client, f *HttpRouteRulesActionRequestMirrorPolicy, res *HttpRoute) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandHttpRouteRulesActionRequestMirrorPolicyDestination(c, f.Destination, res); err != nil {
		return nil, fmt.Errorf("error expanding Destination into destination: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["destination"] = v
	}

	return m, nil
}

// flattenHttpRouteRulesActionRequestMirrorPolicy flattens an instance of HttpRouteRulesActionRequestMirrorPolicy from a JSON
// response object.
func flattenHttpRouteRulesActionRequestMirrorPolicy(c *Client, i interface{}, res *HttpRoute) *HttpRouteRulesActionRequestMirrorPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &HttpRouteRulesActionRequestMirrorPolicy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyHttpRouteRulesActionRequestMirrorPolicy
	}
	r.Destination = flattenHttpRouteRulesActionRequestMirrorPolicyDestination(c, m["destination"], res)

	return r
}

// expandHttpRouteRulesActionRequestMirrorPolicyDestinationMap expands the contents of HttpRouteRulesActionRequestMirrorPolicyDestination into a JSON
// request object.
func expandHttpRouteRulesActionRequestMirrorPolicyDestinationMap(c *Client, f map[string]HttpRouteRulesActionRequestMirrorPolicyDestination, res *HttpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandHttpRouteRulesActionRequestMirrorPolicyDestination(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandHttpRouteRulesActionRequestMirrorPolicyDestinationSlice expands the contents of HttpRouteRulesActionRequestMirrorPolicyDestination into a JSON
// request object.
func expandHttpRouteRulesActionRequestMirrorPolicyDestinationSlice(c *Client, f []HttpRouteRulesActionRequestMirrorPolicyDestination, res *HttpRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandHttpRouteRulesActionRequestMirrorPolicyDestination(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenHttpRouteRulesActionRequestMirrorPolicyDestinationMap flattens the contents of HttpRouteRulesActionRequestMirrorPolicyDestination from a JSON
// response object.
func flattenHttpRouteRulesActionRequestMirrorPolicyDestinationMap(c *Client, i interface{}, res *HttpRoute) map[string]HttpRouteRulesActionRequestMirrorPolicyDestination {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]HttpRouteRulesActionRequestMirrorPolicyDestination{}
	}

	if len(a) == 0 {
		return map[string]HttpRouteRulesActionRequestMirrorPolicyDestination{}
	}

	items := make(map[string]HttpRouteRulesActionRequestMirrorPolicyDestination)
	for k, item := range a {
		items[k] = *flattenHttpRouteRulesActionRequestMirrorPolicyDestination(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenHttpRouteRulesActionRequestMirrorPolicyDestinationSlice flattens the contents of HttpRouteRulesActionRequestMirrorPolicyDestination from a JSON
// response object.
func flattenHttpRouteRulesActionRequestMirrorPolicyDestinationSlice(c *Client, i interface{}, res *HttpRoute) []HttpRouteRulesActionRequestMirrorPolicyDestination {
	a, ok := i.([]interface{})
	if !ok {
		return []HttpRouteRulesActionRequestMirrorPolicyDestination{}
	}

	if len(a) == 0 {
		return []HttpRouteRulesActionRequestMirrorPolicyDestination{}
	}

	items := make([]HttpRouteRulesActionRequestMirrorPolicyDestination, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHttpRouteRulesActionRequestMirrorPolicyDestination(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandHttpRouteRulesActionRequestMirrorPolicyDestination expands an instance of HttpRouteRulesActionRequestMirrorPolicyDestination into a JSON
// request object.
func expandHttpRouteRulesActionRequestMirrorPolicyDestination(c *Client, f *HttpRouteRulesActionRequestMirrorPolicyDestination, res *HttpRoute) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
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

// flattenHttpRouteRulesActionRequestMirrorPolicyDestination flattens an instance of HttpRouteRulesActionRequestMirrorPolicyDestination from a JSON
// response object.
func flattenHttpRouteRulesActionRequestMirrorPolicyDestination(c *Client, i interface{}, res *HttpRoute) *HttpRouteRulesActionRequestMirrorPolicyDestination {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &HttpRouteRulesActionRequestMirrorPolicyDestination{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyHttpRouteRulesActionRequestMirrorPolicyDestination
	}
	r.Weight = dcl.FlattenInteger(m["weight"])
	r.ServiceName = dcl.FlattenString(m["serviceName"])

	return r
}

// expandHttpRouteRulesActionCorsPolicyMap expands the contents of HttpRouteRulesActionCorsPolicy into a JSON
// request object.
func expandHttpRouteRulesActionCorsPolicyMap(c *Client, f map[string]HttpRouteRulesActionCorsPolicy, res *HttpRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandHttpRouteRulesActionCorsPolicy(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandHttpRouteRulesActionCorsPolicySlice expands the contents of HttpRouteRulesActionCorsPolicy into a JSON
// request object.
func expandHttpRouteRulesActionCorsPolicySlice(c *Client, f []HttpRouteRulesActionCorsPolicy, res *HttpRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandHttpRouteRulesActionCorsPolicy(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenHttpRouteRulesActionCorsPolicyMap flattens the contents of HttpRouteRulesActionCorsPolicy from a JSON
// response object.
func flattenHttpRouteRulesActionCorsPolicyMap(c *Client, i interface{}, res *HttpRoute) map[string]HttpRouteRulesActionCorsPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]HttpRouteRulesActionCorsPolicy{}
	}

	if len(a) == 0 {
		return map[string]HttpRouteRulesActionCorsPolicy{}
	}

	items := make(map[string]HttpRouteRulesActionCorsPolicy)
	for k, item := range a {
		items[k] = *flattenHttpRouteRulesActionCorsPolicy(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenHttpRouteRulesActionCorsPolicySlice flattens the contents of HttpRouteRulesActionCorsPolicy from a JSON
// response object.
func flattenHttpRouteRulesActionCorsPolicySlice(c *Client, i interface{}, res *HttpRoute) []HttpRouteRulesActionCorsPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []HttpRouteRulesActionCorsPolicy{}
	}

	if len(a) == 0 {
		return []HttpRouteRulesActionCorsPolicy{}
	}

	items := make([]HttpRouteRulesActionCorsPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHttpRouteRulesActionCorsPolicy(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandHttpRouteRulesActionCorsPolicy expands an instance of HttpRouteRulesActionCorsPolicy into a JSON
// request object.
func expandHttpRouteRulesActionCorsPolicy(c *Client, f *HttpRouteRulesActionCorsPolicy, res *HttpRoute) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AllowOrigins; v != nil {
		m["allowOrigins"] = v
	}
	if v := f.AllowOriginRegexes; v != nil {
		m["allowOriginRegexes"] = v
	}
	if v := f.AllowMethods; v != nil {
		m["allowMethods"] = v
	}
	if v := f.AllowHeaders; v != nil {
		m["allowHeaders"] = v
	}
	if v := f.ExposeHeaders; v != nil {
		m["exposeHeaders"] = v
	}
	if v := f.MaxAge; !dcl.IsEmptyValueIndirect(v) {
		m["maxAge"] = v
	}
	if v := f.AllowCredentials; !dcl.IsEmptyValueIndirect(v) {
		m["allowCredentials"] = v
	}
	if v := f.Disabled; !dcl.IsEmptyValueIndirect(v) {
		m["disabled"] = v
	}

	return m, nil
}

// flattenHttpRouteRulesActionCorsPolicy flattens an instance of HttpRouteRulesActionCorsPolicy from a JSON
// response object.
func flattenHttpRouteRulesActionCorsPolicy(c *Client, i interface{}, res *HttpRoute) *HttpRouteRulesActionCorsPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &HttpRouteRulesActionCorsPolicy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyHttpRouteRulesActionCorsPolicy
	}
	r.AllowOrigins = dcl.FlattenStringSlice(m["allowOrigins"])
	r.AllowOriginRegexes = dcl.FlattenStringSlice(m["allowOriginRegexes"])
	r.AllowMethods = dcl.FlattenStringSlice(m["allowMethods"])
	r.AllowHeaders = dcl.FlattenStringSlice(m["allowHeaders"])
	r.ExposeHeaders = dcl.FlattenStringSlice(m["exposeHeaders"])
	r.MaxAge = dcl.FlattenString(m["maxAge"])
	r.AllowCredentials = dcl.FlattenBool(m["allowCredentials"])
	r.Disabled = dcl.FlattenBool(m["disabled"])

	return r
}

// flattenHttpRouteRulesActionRedirectResponseCodeEnumMap flattens the contents of HttpRouteRulesActionRedirectResponseCodeEnum from a JSON
// response object.
func flattenHttpRouteRulesActionRedirectResponseCodeEnumMap(c *Client, i interface{}, res *HttpRoute) map[string]HttpRouteRulesActionRedirectResponseCodeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]HttpRouteRulesActionRedirectResponseCodeEnum{}
	}

	if len(a) == 0 {
		return map[string]HttpRouteRulesActionRedirectResponseCodeEnum{}
	}

	items := make(map[string]HttpRouteRulesActionRedirectResponseCodeEnum)
	for k, item := range a {
		items[k] = *flattenHttpRouteRulesActionRedirectResponseCodeEnum(item.(interface{}))
	}

	return items
}

// flattenHttpRouteRulesActionRedirectResponseCodeEnumSlice flattens the contents of HttpRouteRulesActionRedirectResponseCodeEnum from a JSON
// response object.
func flattenHttpRouteRulesActionRedirectResponseCodeEnumSlice(c *Client, i interface{}, res *HttpRoute) []HttpRouteRulesActionRedirectResponseCodeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []HttpRouteRulesActionRedirectResponseCodeEnum{}
	}

	if len(a) == 0 {
		return []HttpRouteRulesActionRedirectResponseCodeEnum{}
	}

	items := make([]HttpRouteRulesActionRedirectResponseCodeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenHttpRouteRulesActionRedirectResponseCodeEnum(item.(interface{})))
	}

	return items
}

// flattenHttpRouteRulesActionRedirectResponseCodeEnum asserts that an interface is a string, and returns a
// pointer to a *HttpRouteRulesActionRedirectResponseCodeEnum with the same value as that string.
func flattenHttpRouteRulesActionRedirectResponseCodeEnum(i interface{}) *HttpRouteRulesActionRedirectResponseCodeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return HttpRouteRulesActionRedirectResponseCodeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *HttpRoute) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalHttpRoute(b, c, r)
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

type httpRouteDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         httpRouteApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToHttpRouteDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]httpRouteDiff, error) {
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
	var diffs []httpRouteDiff
	// For each operation name, create a httpRouteDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := httpRouteDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToHttpRouteApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToHttpRouteApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (httpRouteApiOperation, error) {
	switch opName {

	case "updateHttpRouteUpdateHttpRouteOperation":
		return &updateHttpRouteUpdateHttpRouteOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractHttpRouteFields(r *HttpRoute) error {
	return nil
}
func extractHttpRouteRulesFields(r *HttpRoute, o *HttpRouteRules) error {
	vAction := o.Action
	if vAction == nil {
		// note: explicitly not the empty object.
		vAction = &HttpRouteRulesAction{}
	}
	if err := extractHttpRouteRulesActionFields(r, vAction); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAction) {
		o.Action = vAction
	}
	return nil
}
func extractHttpRouteRulesMatchesFields(r *HttpRoute, o *HttpRouteRulesMatches) error {
	return nil
}
func extractHttpRouteRulesMatchesHeadersFields(r *HttpRoute, o *HttpRouteRulesMatchesHeaders) error {
	vRangeMatch := o.RangeMatch
	if vRangeMatch == nil {
		// note: explicitly not the empty object.
		vRangeMatch = &HttpRouteRulesMatchesHeadersRangeMatch{}
	}
	if err := extractHttpRouteRulesMatchesHeadersRangeMatchFields(r, vRangeMatch); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRangeMatch) {
		o.RangeMatch = vRangeMatch
	}
	return nil
}
func extractHttpRouteRulesMatchesHeadersRangeMatchFields(r *HttpRoute, o *HttpRouteRulesMatchesHeadersRangeMatch) error {
	return nil
}
func extractHttpRouteRulesMatchesQueryParametersFields(r *HttpRoute, o *HttpRouteRulesMatchesQueryParameters) error {
	return nil
}
func extractHttpRouteRulesActionFields(r *HttpRoute, o *HttpRouteRulesAction) error {
	vRedirect := o.Redirect
	if vRedirect == nil {
		// note: explicitly not the empty object.
		vRedirect = &HttpRouteRulesActionRedirect{}
	}
	if err := extractHttpRouteRulesActionRedirectFields(r, vRedirect); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRedirect) {
		o.Redirect = vRedirect
	}
	vFaultInjectionPolicy := o.FaultInjectionPolicy
	if vFaultInjectionPolicy == nil {
		// note: explicitly not the empty object.
		vFaultInjectionPolicy = &HttpRouteRulesActionFaultInjectionPolicy{}
	}
	if err := extractHttpRouteRulesActionFaultInjectionPolicyFields(r, vFaultInjectionPolicy); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vFaultInjectionPolicy) {
		o.FaultInjectionPolicy = vFaultInjectionPolicy
	}
	vRequestHeaderModifier := o.RequestHeaderModifier
	if vRequestHeaderModifier == nil {
		// note: explicitly not the empty object.
		vRequestHeaderModifier = &HttpRouteRulesActionRequestHeaderModifier{}
	}
	if err := extractHttpRouteRulesActionRequestHeaderModifierFields(r, vRequestHeaderModifier); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRequestHeaderModifier) {
		o.RequestHeaderModifier = vRequestHeaderModifier
	}
	vResponseHeaderModifier := o.ResponseHeaderModifier
	if vResponseHeaderModifier == nil {
		// note: explicitly not the empty object.
		vResponseHeaderModifier = &HttpRouteRulesActionResponseHeaderModifier{}
	}
	if err := extractHttpRouteRulesActionResponseHeaderModifierFields(r, vResponseHeaderModifier); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vResponseHeaderModifier) {
		o.ResponseHeaderModifier = vResponseHeaderModifier
	}
	vUrlRewrite := o.UrlRewrite
	if vUrlRewrite == nil {
		// note: explicitly not the empty object.
		vUrlRewrite = &HttpRouteRulesActionUrlRewrite{}
	}
	if err := extractHttpRouteRulesActionUrlRewriteFields(r, vUrlRewrite); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vUrlRewrite) {
		o.UrlRewrite = vUrlRewrite
	}
	vRetryPolicy := o.RetryPolicy
	if vRetryPolicy == nil {
		// note: explicitly not the empty object.
		vRetryPolicy = &HttpRouteRulesActionRetryPolicy{}
	}
	if err := extractHttpRouteRulesActionRetryPolicyFields(r, vRetryPolicy); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRetryPolicy) {
		o.RetryPolicy = vRetryPolicy
	}
	vRequestMirrorPolicy := o.RequestMirrorPolicy
	if vRequestMirrorPolicy == nil {
		// note: explicitly not the empty object.
		vRequestMirrorPolicy = &HttpRouteRulesActionRequestMirrorPolicy{}
	}
	if err := extractHttpRouteRulesActionRequestMirrorPolicyFields(r, vRequestMirrorPolicy); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRequestMirrorPolicy) {
		o.RequestMirrorPolicy = vRequestMirrorPolicy
	}
	vCorsPolicy := o.CorsPolicy
	if vCorsPolicy == nil {
		// note: explicitly not the empty object.
		vCorsPolicy = &HttpRouteRulesActionCorsPolicy{}
	}
	if err := extractHttpRouteRulesActionCorsPolicyFields(r, vCorsPolicy); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCorsPolicy) {
		o.CorsPolicy = vCorsPolicy
	}
	return nil
}
func extractHttpRouteRulesActionDestinationsFields(r *HttpRoute, o *HttpRouteRulesActionDestinations) error {
	return nil
}
func extractHttpRouteRulesActionRedirectFields(r *HttpRoute, o *HttpRouteRulesActionRedirect) error {
	return nil
}
func extractHttpRouteRulesActionFaultInjectionPolicyFields(r *HttpRoute, o *HttpRouteRulesActionFaultInjectionPolicy) error {
	vDelay := o.Delay
	if vDelay == nil {
		// note: explicitly not the empty object.
		vDelay = &HttpRouteRulesActionFaultInjectionPolicyDelay{}
	}
	if err := extractHttpRouteRulesActionFaultInjectionPolicyDelayFields(r, vDelay); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDelay) {
		o.Delay = vDelay
	}
	vAbort := o.Abort
	if vAbort == nil {
		// note: explicitly not the empty object.
		vAbort = &HttpRouteRulesActionFaultInjectionPolicyAbort{}
	}
	if err := extractHttpRouteRulesActionFaultInjectionPolicyAbortFields(r, vAbort); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAbort) {
		o.Abort = vAbort
	}
	return nil
}
func extractHttpRouteRulesActionFaultInjectionPolicyDelayFields(r *HttpRoute, o *HttpRouteRulesActionFaultInjectionPolicyDelay) error {
	return nil
}
func extractHttpRouteRulesActionFaultInjectionPolicyAbortFields(r *HttpRoute, o *HttpRouteRulesActionFaultInjectionPolicyAbort) error {
	return nil
}
func extractHttpRouteRulesActionRequestHeaderModifierFields(r *HttpRoute, o *HttpRouteRulesActionRequestHeaderModifier) error {
	return nil
}
func extractHttpRouteRulesActionResponseHeaderModifierFields(r *HttpRoute, o *HttpRouteRulesActionResponseHeaderModifier) error {
	return nil
}
func extractHttpRouteRulesActionUrlRewriteFields(r *HttpRoute, o *HttpRouteRulesActionUrlRewrite) error {
	return nil
}
func extractHttpRouteRulesActionRetryPolicyFields(r *HttpRoute, o *HttpRouteRulesActionRetryPolicy) error {
	return nil
}
func extractHttpRouteRulesActionRequestMirrorPolicyFields(r *HttpRoute, o *HttpRouteRulesActionRequestMirrorPolicy) error {
	vDestination := o.Destination
	if vDestination == nil {
		// note: explicitly not the empty object.
		vDestination = &HttpRouteRulesActionRequestMirrorPolicyDestination{}
	}
	if err := extractHttpRouteRulesActionRequestMirrorPolicyDestinationFields(r, vDestination); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDestination) {
		o.Destination = vDestination
	}
	return nil
}
func extractHttpRouteRulesActionRequestMirrorPolicyDestinationFields(r *HttpRoute, o *HttpRouteRulesActionRequestMirrorPolicyDestination) error {
	return nil
}
func extractHttpRouteRulesActionCorsPolicyFields(r *HttpRoute, o *HttpRouteRulesActionCorsPolicy) error {
	return nil
}

func postReadExtractHttpRouteFields(r *HttpRoute) error {
	return nil
}
func postReadExtractHttpRouteRulesFields(r *HttpRoute, o *HttpRouteRules) error {
	vAction := o.Action
	if vAction == nil {
		// note: explicitly not the empty object.
		vAction = &HttpRouteRulesAction{}
	}
	if err := extractHttpRouteRulesActionFields(r, vAction); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAction) {
		o.Action = vAction
	}
	return nil
}
func postReadExtractHttpRouteRulesMatchesFields(r *HttpRoute, o *HttpRouteRulesMatches) error {
	return nil
}
func postReadExtractHttpRouteRulesMatchesHeadersFields(r *HttpRoute, o *HttpRouteRulesMatchesHeaders) error {
	vRangeMatch := o.RangeMatch
	if vRangeMatch == nil {
		// note: explicitly not the empty object.
		vRangeMatch = &HttpRouteRulesMatchesHeadersRangeMatch{}
	}
	if err := extractHttpRouteRulesMatchesHeadersRangeMatchFields(r, vRangeMatch); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRangeMatch) {
		o.RangeMatch = vRangeMatch
	}
	return nil
}
func postReadExtractHttpRouteRulesMatchesHeadersRangeMatchFields(r *HttpRoute, o *HttpRouteRulesMatchesHeadersRangeMatch) error {
	return nil
}
func postReadExtractHttpRouteRulesMatchesQueryParametersFields(r *HttpRoute, o *HttpRouteRulesMatchesQueryParameters) error {
	return nil
}
func postReadExtractHttpRouteRulesActionFields(r *HttpRoute, o *HttpRouteRulesAction) error {
	vRedirect := o.Redirect
	if vRedirect == nil {
		// note: explicitly not the empty object.
		vRedirect = &HttpRouteRulesActionRedirect{}
	}
	if err := extractHttpRouteRulesActionRedirectFields(r, vRedirect); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRedirect) {
		o.Redirect = vRedirect
	}
	vFaultInjectionPolicy := o.FaultInjectionPolicy
	if vFaultInjectionPolicy == nil {
		// note: explicitly not the empty object.
		vFaultInjectionPolicy = &HttpRouteRulesActionFaultInjectionPolicy{}
	}
	if err := extractHttpRouteRulesActionFaultInjectionPolicyFields(r, vFaultInjectionPolicy); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vFaultInjectionPolicy) {
		o.FaultInjectionPolicy = vFaultInjectionPolicy
	}
	vRequestHeaderModifier := o.RequestHeaderModifier
	if vRequestHeaderModifier == nil {
		// note: explicitly not the empty object.
		vRequestHeaderModifier = &HttpRouteRulesActionRequestHeaderModifier{}
	}
	if err := extractHttpRouteRulesActionRequestHeaderModifierFields(r, vRequestHeaderModifier); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRequestHeaderModifier) {
		o.RequestHeaderModifier = vRequestHeaderModifier
	}
	vResponseHeaderModifier := o.ResponseHeaderModifier
	if vResponseHeaderModifier == nil {
		// note: explicitly not the empty object.
		vResponseHeaderModifier = &HttpRouteRulesActionResponseHeaderModifier{}
	}
	if err := extractHttpRouteRulesActionResponseHeaderModifierFields(r, vResponseHeaderModifier); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vResponseHeaderModifier) {
		o.ResponseHeaderModifier = vResponseHeaderModifier
	}
	vUrlRewrite := o.UrlRewrite
	if vUrlRewrite == nil {
		// note: explicitly not the empty object.
		vUrlRewrite = &HttpRouteRulesActionUrlRewrite{}
	}
	if err := extractHttpRouteRulesActionUrlRewriteFields(r, vUrlRewrite); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vUrlRewrite) {
		o.UrlRewrite = vUrlRewrite
	}
	vRetryPolicy := o.RetryPolicy
	if vRetryPolicy == nil {
		// note: explicitly not the empty object.
		vRetryPolicy = &HttpRouteRulesActionRetryPolicy{}
	}
	if err := extractHttpRouteRulesActionRetryPolicyFields(r, vRetryPolicy); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRetryPolicy) {
		o.RetryPolicy = vRetryPolicy
	}
	vRequestMirrorPolicy := o.RequestMirrorPolicy
	if vRequestMirrorPolicy == nil {
		// note: explicitly not the empty object.
		vRequestMirrorPolicy = &HttpRouteRulesActionRequestMirrorPolicy{}
	}
	if err := extractHttpRouteRulesActionRequestMirrorPolicyFields(r, vRequestMirrorPolicy); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRequestMirrorPolicy) {
		o.RequestMirrorPolicy = vRequestMirrorPolicy
	}
	vCorsPolicy := o.CorsPolicy
	if vCorsPolicy == nil {
		// note: explicitly not the empty object.
		vCorsPolicy = &HttpRouteRulesActionCorsPolicy{}
	}
	if err := extractHttpRouteRulesActionCorsPolicyFields(r, vCorsPolicy); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCorsPolicy) {
		o.CorsPolicy = vCorsPolicy
	}
	return nil
}
func postReadExtractHttpRouteRulesActionDestinationsFields(r *HttpRoute, o *HttpRouteRulesActionDestinations) error {
	return nil
}
func postReadExtractHttpRouteRulesActionRedirectFields(r *HttpRoute, o *HttpRouteRulesActionRedirect) error {
	return nil
}
func postReadExtractHttpRouteRulesActionFaultInjectionPolicyFields(r *HttpRoute, o *HttpRouteRulesActionFaultInjectionPolicy) error {
	vDelay := o.Delay
	if vDelay == nil {
		// note: explicitly not the empty object.
		vDelay = &HttpRouteRulesActionFaultInjectionPolicyDelay{}
	}
	if err := extractHttpRouteRulesActionFaultInjectionPolicyDelayFields(r, vDelay); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDelay) {
		o.Delay = vDelay
	}
	vAbort := o.Abort
	if vAbort == nil {
		// note: explicitly not the empty object.
		vAbort = &HttpRouteRulesActionFaultInjectionPolicyAbort{}
	}
	if err := extractHttpRouteRulesActionFaultInjectionPolicyAbortFields(r, vAbort); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAbort) {
		o.Abort = vAbort
	}
	return nil
}
func postReadExtractHttpRouteRulesActionFaultInjectionPolicyDelayFields(r *HttpRoute, o *HttpRouteRulesActionFaultInjectionPolicyDelay) error {
	return nil
}
func postReadExtractHttpRouteRulesActionFaultInjectionPolicyAbortFields(r *HttpRoute, o *HttpRouteRulesActionFaultInjectionPolicyAbort) error {
	return nil
}
func postReadExtractHttpRouteRulesActionRequestHeaderModifierFields(r *HttpRoute, o *HttpRouteRulesActionRequestHeaderModifier) error {
	return nil
}
func postReadExtractHttpRouteRulesActionResponseHeaderModifierFields(r *HttpRoute, o *HttpRouteRulesActionResponseHeaderModifier) error {
	return nil
}
func postReadExtractHttpRouteRulesActionUrlRewriteFields(r *HttpRoute, o *HttpRouteRulesActionUrlRewrite) error {
	return nil
}
func postReadExtractHttpRouteRulesActionRetryPolicyFields(r *HttpRoute, o *HttpRouteRulesActionRetryPolicy) error {
	return nil
}
func postReadExtractHttpRouteRulesActionRequestMirrorPolicyFields(r *HttpRoute, o *HttpRouteRulesActionRequestMirrorPolicy) error {
	vDestination := o.Destination
	if vDestination == nil {
		// note: explicitly not the empty object.
		vDestination = &HttpRouteRulesActionRequestMirrorPolicyDestination{}
	}
	if err := extractHttpRouteRulesActionRequestMirrorPolicyDestinationFields(r, vDestination); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDestination) {
		o.Destination = vDestination
	}
	return nil
}
func postReadExtractHttpRouteRulesActionRequestMirrorPolicyDestinationFields(r *HttpRoute, o *HttpRouteRulesActionRequestMirrorPolicyDestination) error {
	return nil
}
func postReadExtractHttpRouteRulesActionCorsPolicyFields(r *HttpRoute, o *HttpRouteRulesActionCorsPolicy) error {
	return nil
}
