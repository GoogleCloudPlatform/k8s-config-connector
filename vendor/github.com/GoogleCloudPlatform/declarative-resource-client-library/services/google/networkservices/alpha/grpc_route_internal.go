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

func (r *GrpcRoute) validate() error {

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
func (r *GrpcRouteRules) validate() error {
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
func (r *GrpcRouteRulesMatches) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"Method"}, r.Method); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Method) {
		if err := r.Method.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *GrpcRouteRulesMatchesMethod) validate() error {
	if err := dcl.Required(r, "grpcService"); err != nil {
		return err
	}
	if err := dcl.Required(r, "grpcMethod"); err != nil {
		return err
	}
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"CaseSensitive"}, r.CaseSensitive); err != nil {
		return err
	}
	return nil
}
func (r *GrpcRouteRulesMatchesHeaders) validate() error {
	if err := dcl.Required(r, "key"); err != nil {
		return err
	}
	if err := dcl.Required(r, "value"); err != nil {
		return err
	}
	return nil
}
func (r *GrpcRouteRulesAction) validate() error {
	if !dcl.IsEmptyValueIndirect(r.FaultInjectionPolicy) {
		if err := r.FaultInjectionPolicy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.RetryPolicy) {
		if err := r.RetryPolicy.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *GrpcRouteRulesActionDestinations) validate() error {
	if err := dcl.Required(r, "serviceName"); err != nil {
		return err
	}
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"Weight"}, r.Weight); err != nil {
		return err
	}
	return nil
}
func (r *GrpcRouteRulesActionFaultInjectionPolicy) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"Delay"}, r.Delay); err != nil {
		return err
	}
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"Abort"}, r.Abort); err != nil {
		return err
	}
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
func (r *GrpcRouteRulesActionFaultInjectionPolicyDelay) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"FixedDelay"}, r.FixedDelay); err != nil {
		return err
	}
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"Percentage"}, r.Percentage); err != nil {
		return err
	}
	return nil
}
func (r *GrpcRouteRulesActionFaultInjectionPolicyAbort) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"HttpStatus"}, r.HttpStatus); err != nil {
		return err
	}
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"Percentage"}, r.Percentage); err != nil {
		return err
	}
	return nil
}
func (r *GrpcRouteRulesActionRetryPolicy) validate() error {
	return nil
}
func (r *GrpcRoute) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://networkservices.googleapis.com/v1alpha1/", params)
}

func (r *GrpcRoute) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/grpcRoutes/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *GrpcRoute) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/grpcRoutes", nr.basePath(), userBasePath, params), nil

}

func (r *GrpcRoute) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/grpcRoutes?grpcRouteId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *GrpcRoute) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/grpcRoutes/{{name}}", nr.basePath(), userBasePath, params), nil
}

// grpcRouteApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type grpcRouteApiOperation interface {
	do(context.Context, *GrpcRoute, *Client) error
}

// newUpdateGrpcRouteUpdateGrpcRouteRequest creates a request for an
// GrpcRoute resource's UpdateGrpcRoute update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateGrpcRouteUpdateGrpcRouteRequest(ctx context.Context, f *GrpcRoute, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v, err := dcl.DeriveField("projects/%s/locations/global/grpcRoutes/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
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
	if v, err := expandGrpcRouteRulesSlice(c, f.Rules, res); err != nil {
		return nil, fmt.Errorf("error expanding Rules into rules: %w", err)
	} else if v != nil {
		req["rules"] = v
	}
	req["name"] = fmt.Sprintf("projects/%s/locations/%s/grpcRoutes/%s", *f.Project, *f.Location, *f.Name)

	return req, nil
}

// marshalUpdateGrpcRouteUpdateGrpcRouteRequest converts the update into
// the final JSON request body.
func marshalUpdateGrpcRouteUpdateGrpcRouteRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateGrpcRouteUpdateGrpcRouteOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateGrpcRouteUpdateGrpcRouteOperation) do(ctx context.Context, r *GrpcRoute, c *Client) error {
	_, err := c.GetGrpcRoute(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateGrpcRoute")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateGrpcRouteUpdateGrpcRouteRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateGrpcRouteUpdateGrpcRouteRequest(c, req)
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

func (c *Client) listGrpcRouteRaw(ctx context.Context, r *GrpcRoute, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != GrpcRouteMaxPage {
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

type listGrpcRouteOperation struct {
	GrpcRoutes []map[string]interface{} `json:"grpcRoutes"`
	Token      string                   `json:"nextPageToken"`
}

func (c *Client) listGrpcRoute(ctx context.Context, r *GrpcRoute, pageToken string, pageSize int32) ([]*GrpcRoute, string, error) {
	b, err := c.listGrpcRouteRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listGrpcRouteOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*GrpcRoute
	for _, v := range m.GrpcRoutes {
		res, err := unmarshalMapGrpcRoute(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllGrpcRoute(ctx context.Context, f func(*GrpcRoute) bool, resources []*GrpcRoute) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteGrpcRoute(ctx, res)
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

type deleteGrpcRouteOperation struct{}

func (op *deleteGrpcRouteOperation) do(ctx context.Context, r *GrpcRoute, c *Client) error {
	r, err := c.GetGrpcRoute(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "GrpcRoute not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetGrpcRoute checking for existence. error: %v", err)
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
		_, err := c.GetGrpcRoute(ctx, r)
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
type createGrpcRouteOperation struct {
	response map[string]interface{}
}

func (op *createGrpcRouteOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createGrpcRouteOperation) do(ctx context.Context, r *GrpcRoute, c *Client) error {
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

	if _, err := c.GetGrpcRoute(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getGrpcRouteRaw(ctx context.Context, r *GrpcRoute) ([]byte, error) {

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

func (c *Client) grpcRouteDiffsForRawDesired(ctx context.Context, rawDesired *GrpcRoute, opts ...dcl.ApplyOption) (initial, desired *GrpcRoute, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *GrpcRoute
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*GrpcRoute); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected GrpcRoute, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetGrpcRoute(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a GrpcRoute resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve GrpcRoute resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that GrpcRoute resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeGrpcRouteDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for GrpcRoute: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for GrpcRoute: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractGrpcRouteFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeGrpcRouteInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for GrpcRoute: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeGrpcRouteDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for GrpcRoute: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffGrpcRoute(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeGrpcRouteInitialState(rawInitial, rawDesired *GrpcRoute) (*GrpcRoute, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeGrpcRouteDesiredState(rawDesired, rawInitial *GrpcRoute, opts ...dcl.ApplyOption) (*GrpcRoute, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &GrpcRoute{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.IsZeroValue(rawDesired.Labels) || (dcl.IsEmptyValueIndirect(rawDesired.Labels) && dcl.IsEmptyValueIndirect(rawInitial.Labels)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
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
	canonicalDesired.Rules = canonicalizeGrpcRouteRulesSlice(rawDesired.Rules, rawInitial.Rules, opts...)
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

func canonicalizeGrpcRouteNewState(c *Client, rawNew, rawDesired *GrpcRoute) (*GrpcRoute, error) {

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

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
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

	if dcl.IsEmptyValueIndirect(rawNew.Rules) && dcl.IsEmptyValueIndirect(rawDesired.Rules) {
		rawNew.Rules = rawDesired.Rules
	} else {
		rawNew.Rules = canonicalizeNewGrpcRouteRulesSlice(c, rawDesired.Rules, rawNew.Rules)
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

func canonicalizeGrpcRouteRules(des, initial *GrpcRouteRules, opts ...dcl.ApplyOption) *GrpcRouteRules {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &GrpcRouteRules{}

	cDes.Matches = canonicalizeGrpcRouteRulesMatchesSlice(des.Matches, initial.Matches, opts...)
	cDes.Action = canonicalizeGrpcRouteRulesAction(des.Action, initial.Action, opts...)

	return cDes
}

func canonicalizeGrpcRouteRulesSlice(des, initial []GrpcRouteRules, opts ...dcl.ApplyOption) []GrpcRouteRules {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]GrpcRouteRules, 0, len(des))
		for _, d := range des {
			cd := canonicalizeGrpcRouteRules(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]GrpcRouteRules, 0, len(des))
	for i, d := range des {
		cd := canonicalizeGrpcRouteRules(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewGrpcRouteRules(c *Client, des, nw *GrpcRouteRules) *GrpcRouteRules {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for GrpcRouteRules while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Matches = canonicalizeNewGrpcRouteRulesMatchesSlice(c, des.Matches, nw.Matches)
	nw.Action = canonicalizeNewGrpcRouteRulesAction(c, des.Action, nw.Action)

	return nw
}

func canonicalizeNewGrpcRouteRulesSet(c *Client, des, nw []GrpcRouteRules) []GrpcRouteRules {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []GrpcRouteRules
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareGrpcRouteRulesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewGrpcRouteRules(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewGrpcRouteRulesSlice(c *Client, des, nw []GrpcRouteRules) []GrpcRouteRules {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []GrpcRouteRules
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGrpcRouteRules(c, &d, &n))
	}

	return items
}

func canonicalizeGrpcRouteRulesMatches(des, initial *GrpcRouteRulesMatches, opts ...dcl.ApplyOption) *GrpcRouteRulesMatches {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.Method != nil || (initial != nil && initial.Method != nil) {
		// Check if anything else is set.
		if dcl.AnySet() {
			des.Method = nil
			if initial != nil {
				initial.Method = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &GrpcRouteRulesMatches{}

	cDes.Method = canonicalizeGrpcRouteRulesMatchesMethod(des.Method, initial.Method, opts...)
	cDes.Headers = canonicalizeGrpcRouteRulesMatchesHeadersSlice(des.Headers, initial.Headers, opts...)

	return cDes
}

func canonicalizeGrpcRouteRulesMatchesSlice(des, initial []GrpcRouteRulesMatches, opts ...dcl.ApplyOption) []GrpcRouteRulesMatches {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]GrpcRouteRulesMatches, 0, len(des))
		for _, d := range des {
			cd := canonicalizeGrpcRouteRulesMatches(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]GrpcRouteRulesMatches, 0, len(des))
	for i, d := range des {
		cd := canonicalizeGrpcRouteRulesMatches(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewGrpcRouteRulesMatches(c *Client, des, nw *GrpcRouteRulesMatches) *GrpcRouteRulesMatches {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for GrpcRouteRulesMatches while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Method = canonicalizeNewGrpcRouteRulesMatchesMethod(c, des.Method, nw.Method)
	nw.Headers = canonicalizeNewGrpcRouteRulesMatchesHeadersSlice(c, des.Headers, nw.Headers)

	return nw
}

func canonicalizeNewGrpcRouteRulesMatchesSet(c *Client, des, nw []GrpcRouteRulesMatches) []GrpcRouteRulesMatches {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []GrpcRouteRulesMatches
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareGrpcRouteRulesMatchesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewGrpcRouteRulesMatches(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewGrpcRouteRulesMatchesSlice(c *Client, des, nw []GrpcRouteRulesMatches) []GrpcRouteRulesMatches {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []GrpcRouteRulesMatches
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGrpcRouteRulesMatches(c, &d, &n))
	}

	return items
}

func canonicalizeGrpcRouteRulesMatchesMethod(des, initial *GrpcRouteRulesMatchesMethod, opts ...dcl.ApplyOption) *GrpcRouteRulesMatchesMethod {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.CaseSensitive != nil || (initial != nil && initial.CaseSensitive != nil) {
		// Check if anything else is set.
		if dcl.AnySet() {
			des.CaseSensitive = nil
			if initial != nil {
				initial.CaseSensitive = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &GrpcRouteRulesMatchesMethod{}

	if dcl.IsZeroValue(des.Type) || (dcl.IsEmptyValueIndirect(des.Type) && dcl.IsEmptyValueIndirect(initial.Type)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Type = initial.Type
	} else {
		cDes.Type = des.Type
	}
	if dcl.StringCanonicalize(des.GrpcService, initial.GrpcService) || dcl.IsZeroValue(des.GrpcService) {
		cDes.GrpcService = initial.GrpcService
	} else {
		cDes.GrpcService = des.GrpcService
	}
	if dcl.StringCanonicalize(des.GrpcMethod, initial.GrpcMethod) || dcl.IsZeroValue(des.GrpcMethod) {
		cDes.GrpcMethod = initial.GrpcMethod
	} else {
		cDes.GrpcMethod = des.GrpcMethod
	}
	if dcl.BoolCanonicalize(des.CaseSensitive, initial.CaseSensitive) || dcl.IsZeroValue(des.CaseSensitive) {
		cDes.CaseSensitive = initial.CaseSensitive
	} else {
		cDes.CaseSensitive = des.CaseSensitive
	}

	return cDes
}

func canonicalizeGrpcRouteRulesMatchesMethodSlice(des, initial []GrpcRouteRulesMatchesMethod, opts ...dcl.ApplyOption) []GrpcRouteRulesMatchesMethod {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]GrpcRouteRulesMatchesMethod, 0, len(des))
		for _, d := range des {
			cd := canonicalizeGrpcRouteRulesMatchesMethod(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]GrpcRouteRulesMatchesMethod, 0, len(des))
	for i, d := range des {
		cd := canonicalizeGrpcRouteRulesMatchesMethod(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewGrpcRouteRulesMatchesMethod(c *Client, des, nw *GrpcRouteRulesMatchesMethod) *GrpcRouteRulesMatchesMethod {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for GrpcRouteRulesMatchesMethod while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.GrpcService, nw.GrpcService) {
		nw.GrpcService = des.GrpcService
	}
	if dcl.StringCanonicalize(des.GrpcMethod, nw.GrpcMethod) {
		nw.GrpcMethod = des.GrpcMethod
	}
	if dcl.BoolCanonicalize(des.CaseSensitive, nw.CaseSensitive) {
		nw.CaseSensitive = des.CaseSensitive
	}

	return nw
}

func canonicalizeNewGrpcRouteRulesMatchesMethodSet(c *Client, des, nw []GrpcRouteRulesMatchesMethod) []GrpcRouteRulesMatchesMethod {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []GrpcRouteRulesMatchesMethod
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareGrpcRouteRulesMatchesMethodNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewGrpcRouteRulesMatchesMethod(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewGrpcRouteRulesMatchesMethodSlice(c *Client, des, nw []GrpcRouteRulesMatchesMethod) []GrpcRouteRulesMatchesMethod {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []GrpcRouteRulesMatchesMethod
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGrpcRouteRulesMatchesMethod(c, &d, &n))
	}

	return items
}

func canonicalizeGrpcRouteRulesMatchesHeaders(des, initial *GrpcRouteRulesMatchesHeaders, opts ...dcl.ApplyOption) *GrpcRouteRulesMatchesHeaders {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &GrpcRouteRulesMatchesHeaders{}

	if dcl.IsZeroValue(des.Type) || (dcl.IsEmptyValueIndirect(des.Type) && dcl.IsEmptyValueIndirect(initial.Type)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Type = initial.Type
	} else {
		cDes.Type = des.Type
	}
	if dcl.StringCanonicalize(des.Key, initial.Key) || dcl.IsZeroValue(des.Key) {
		cDes.Key = initial.Key
	} else {
		cDes.Key = des.Key
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		cDes.Value = initial.Value
	} else {
		cDes.Value = des.Value
	}

	return cDes
}

func canonicalizeGrpcRouteRulesMatchesHeadersSlice(des, initial []GrpcRouteRulesMatchesHeaders, opts ...dcl.ApplyOption) []GrpcRouteRulesMatchesHeaders {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]GrpcRouteRulesMatchesHeaders, 0, len(des))
		for _, d := range des {
			cd := canonicalizeGrpcRouteRulesMatchesHeaders(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]GrpcRouteRulesMatchesHeaders, 0, len(des))
	for i, d := range des {
		cd := canonicalizeGrpcRouteRulesMatchesHeaders(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewGrpcRouteRulesMatchesHeaders(c *Client, des, nw *GrpcRouteRulesMatchesHeaders) *GrpcRouteRulesMatchesHeaders {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for GrpcRouteRulesMatchesHeaders while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Key, nw.Key) {
		nw.Key = des.Key
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}

	return nw
}

func canonicalizeNewGrpcRouteRulesMatchesHeadersSet(c *Client, des, nw []GrpcRouteRulesMatchesHeaders) []GrpcRouteRulesMatchesHeaders {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []GrpcRouteRulesMatchesHeaders
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareGrpcRouteRulesMatchesHeadersNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewGrpcRouteRulesMatchesHeaders(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewGrpcRouteRulesMatchesHeadersSlice(c *Client, des, nw []GrpcRouteRulesMatchesHeaders) []GrpcRouteRulesMatchesHeaders {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []GrpcRouteRulesMatchesHeaders
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGrpcRouteRulesMatchesHeaders(c, &d, &n))
	}

	return items
}

func canonicalizeGrpcRouteRulesAction(des, initial *GrpcRouteRulesAction, opts ...dcl.ApplyOption) *GrpcRouteRulesAction {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &GrpcRouteRulesAction{}

	cDes.Destinations = canonicalizeGrpcRouteRulesActionDestinationsSlice(des.Destinations, initial.Destinations, opts...)
	cDes.FaultInjectionPolicy = canonicalizeGrpcRouteRulesActionFaultInjectionPolicy(des.FaultInjectionPolicy, initial.FaultInjectionPolicy, opts...)
	if dcl.StringCanonicalize(des.Timeout, initial.Timeout) || dcl.IsZeroValue(des.Timeout) {
		cDes.Timeout = initial.Timeout
	} else {
		cDes.Timeout = des.Timeout
	}
	cDes.RetryPolicy = canonicalizeGrpcRouteRulesActionRetryPolicy(des.RetryPolicy, initial.RetryPolicy, opts...)

	return cDes
}

func canonicalizeGrpcRouteRulesActionSlice(des, initial []GrpcRouteRulesAction, opts ...dcl.ApplyOption) []GrpcRouteRulesAction {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]GrpcRouteRulesAction, 0, len(des))
		for _, d := range des {
			cd := canonicalizeGrpcRouteRulesAction(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]GrpcRouteRulesAction, 0, len(des))
	for i, d := range des {
		cd := canonicalizeGrpcRouteRulesAction(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewGrpcRouteRulesAction(c *Client, des, nw *GrpcRouteRulesAction) *GrpcRouteRulesAction {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for GrpcRouteRulesAction while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Destinations = canonicalizeNewGrpcRouteRulesActionDestinationsSlice(c, des.Destinations, nw.Destinations)
	nw.FaultInjectionPolicy = canonicalizeNewGrpcRouteRulesActionFaultInjectionPolicy(c, des.FaultInjectionPolicy, nw.FaultInjectionPolicy)
	if dcl.StringCanonicalize(des.Timeout, nw.Timeout) {
		nw.Timeout = des.Timeout
	}
	nw.RetryPolicy = canonicalizeNewGrpcRouteRulesActionRetryPolicy(c, des.RetryPolicy, nw.RetryPolicy)

	return nw
}

func canonicalizeNewGrpcRouteRulesActionSet(c *Client, des, nw []GrpcRouteRulesAction) []GrpcRouteRulesAction {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []GrpcRouteRulesAction
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareGrpcRouteRulesActionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewGrpcRouteRulesAction(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewGrpcRouteRulesActionSlice(c *Client, des, nw []GrpcRouteRulesAction) []GrpcRouteRulesAction {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []GrpcRouteRulesAction
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGrpcRouteRulesAction(c, &d, &n))
	}

	return items
}

func canonicalizeGrpcRouteRulesActionDestinations(des, initial *GrpcRouteRulesActionDestinations, opts ...dcl.ApplyOption) *GrpcRouteRulesActionDestinations {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.Weight != nil || (initial != nil && initial.Weight != nil) {
		// Check if anything else is set.
		if dcl.AnySet() {
			des.Weight = nil
			if initial != nil {
				initial.Weight = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &GrpcRouteRulesActionDestinations{}

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

func canonicalizeGrpcRouteRulesActionDestinationsSlice(des, initial []GrpcRouteRulesActionDestinations, opts ...dcl.ApplyOption) []GrpcRouteRulesActionDestinations {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]GrpcRouteRulesActionDestinations, 0, len(des))
		for _, d := range des {
			cd := canonicalizeGrpcRouteRulesActionDestinations(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]GrpcRouteRulesActionDestinations, 0, len(des))
	for i, d := range des {
		cd := canonicalizeGrpcRouteRulesActionDestinations(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewGrpcRouteRulesActionDestinations(c *Client, des, nw *GrpcRouteRulesActionDestinations) *GrpcRouteRulesActionDestinations {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for GrpcRouteRulesActionDestinations while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewGrpcRouteRulesActionDestinationsSet(c *Client, des, nw []GrpcRouteRulesActionDestinations) []GrpcRouteRulesActionDestinations {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []GrpcRouteRulesActionDestinations
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareGrpcRouteRulesActionDestinationsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewGrpcRouteRulesActionDestinations(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewGrpcRouteRulesActionDestinationsSlice(c *Client, des, nw []GrpcRouteRulesActionDestinations) []GrpcRouteRulesActionDestinations {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []GrpcRouteRulesActionDestinations
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGrpcRouteRulesActionDestinations(c, &d, &n))
	}

	return items
}

func canonicalizeGrpcRouteRulesActionFaultInjectionPolicy(des, initial *GrpcRouteRulesActionFaultInjectionPolicy, opts ...dcl.ApplyOption) *GrpcRouteRulesActionFaultInjectionPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.Delay != nil || (initial != nil && initial.Delay != nil) {
		// Check if anything else is set.
		if dcl.AnySet() {
			des.Delay = nil
			if initial != nil {
				initial.Delay = nil
			}
		}
	}

	if des.Abort != nil || (initial != nil && initial.Abort != nil) {
		// Check if anything else is set.
		if dcl.AnySet() {
			des.Abort = nil
			if initial != nil {
				initial.Abort = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &GrpcRouteRulesActionFaultInjectionPolicy{}

	cDes.Delay = canonicalizeGrpcRouteRulesActionFaultInjectionPolicyDelay(des.Delay, initial.Delay, opts...)
	cDes.Abort = canonicalizeGrpcRouteRulesActionFaultInjectionPolicyAbort(des.Abort, initial.Abort, opts...)

	return cDes
}

func canonicalizeGrpcRouteRulesActionFaultInjectionPolicySlice(des, initial []GrpcRouteRulesActionFaultInjectionPolicy, opts ...dcl.ApplyOption) []GrpcRouteRulesActionFaultInjectionPolicy {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]GrpcRouteRulesActionFaultInjectionPolicy, 0, len(des))
		for _, d := range des {
			cd := canonicalizeGrpcRouteRulesActionFaultInjectionPolicy(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]GrpcRouteRulesActionFaultInjectionPolicy, 0, len(des))
	for i, d := range des {
		cd := canonicalizeGrpcRouteRulesActionFaultInjectionPolicy(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewGrpcRouteRulesActionFaultInjectionPolicy(c *Client, des, nw *GrpcRouteRulesActionFaultInjectionPolicy) *GrpcRouteRulesActionFaultInjectionPolicy {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for GrpcRouteRulesActionFaultInjectionPolicy while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Delay = canonicalizeNewGrpcRouteRulesActionFaultInjectionPolicyDelay(c, des.Delay, nw.Delay)
	nw.Abort = canonicalizeNewGrpcRouteRulesActionFaultInjectionPolicyAbort(c, des.Abort, nw.Abort)

	return nw
}

func canonicalizeNewGrpcRouteRulesActionFaultInjectionPolicySet(c *Client, des, nw []GrpcRouteRulesActionFaultInjectionPolicy) []GrpcRouteRulesActionFaultInjectionPolicy {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []GrpcRouteRulesActionFaultInjectionPolicy
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareGrpcRouteRulesActionFaultInjectionPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewGrpcRouteRulesActionFaultInjectionPolicy(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewGrpcRouteRulesActionFaultInjectionPolicySlice(c *Client, des, nw []GrpcRouteRulesActionFaultInjectionPolicy) []GrpcRouteRulesActionFaultInjectionPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []GrpcRouteRulesActionFaultInjectionPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGrpcRouteRulesActionFaultInjectionPolicy(c, &d, &n))
	}

	return items
}

func canonicalizeGrpcRouteRulesActionFaultInjectionPolicyDelay(des, initial *GrpcRouteRulesActionFaultInjectionPolicyDelay, opts ...dcl.ApplyOption) *GrpcRouteRulesActionFaultInjectionPolicyDelay {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.FixedDelay != nil || (initial != nil && initial.FixedDelay != nil) {
		// Check if anything else is set.
		if dcl.AnySet() {
			des.FixedDelay = nil
			if initial != nil {
				initial.FixedDelay = nil
			}
		}
	}

	if des.Percentage != nil || (initial != nil && initial.Percentage != nil) {
		// Check if anything else is set.
		if dcl.AnySet() {
			des.Percentage = nil
			if initial != nil {
				initial.Percentage = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &GrpcRouteRulesActionFaultInjectionPolicyDelay{}

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

func canonicalizeGrpcRouteRulesActionFaultInjectionPolicyDelaySlice(des, initial []GrpcRouteRulesActionFaultInjectionPolicyDelay, opts ...dcl.ApplyOption) []GrpcRouteRulesActionFaultInjectionPolicyDelay {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]GrpcRouteRulesActionFaultInjectionPolicyDelay, 0, len(des))
		for _, d := range des {
			cd := canonicalizeGrpcRouteRulesActionFaultInjectionPolicyDelay(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]GrpcRouteRulesActionFaultInjectionPolicyDelay, 0, len(des))
	for i, d := range des {
		cd := canonicalizeGrpcRouteRulesActionFaultInjectionPolicyDelay(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewGrpcRouteRulesActionFaultInjectionPolicyDelay(c *Client, des, nw *GrpcRouteRulesActionFaultInjectionPolicyDelay) *GrpcRouteRulesActionFaultInjectionPolicyDelay {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for GrpcRouteRulesActionFaultInjectionPolicyDelay while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.FixedDelay, nw.FixedDelay) {
		nw.FixedDelay = des.FixedDelay
	}

	return nw
}

func canonicalizeNewGrpcRouteRulesActionFaultInjectionPolicyDelaySet(c *Client, des, nw []GrpcRouteRulesActionFaultInjectionPolicyDelay) []GrpcRouteRulesActionFaultInjectionPolicyDelay {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []GrpcRouteRulesActionFaultInjectionPolicyDelay
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareGrpcRouteRulesActionFaultInjectionPolicyDelayNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewGrpcRouteRulesActionFaultInjectionPolicyDelay(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewGrpcRouteRulesActionFaultInjectionPolicyDelaySlice(c *Client, des, nw []GrpcRouteRulesActionFaultInjectionPolicyDelay) []GrpcRouteRulesActionFaultInjectionPolicyDelay {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []GrpcRouteRulesActionFaultInjectionPolicyDelay
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGrpcRouteRulesActionFaultInjectionPolicyDelay(c, &d, &n))
	}

	return items
}

func canonicalizeGrpcRouteRulesActionFaultInjectionPolicyAbort(des, initial *GrpcRouteRulesActionFaultInjectionPolicyAbort, opts ...dcl.ApplyOption) *GrpcRouteRulesActionFaultInjectionPolicyAbort {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.HttpStatus != nil || (initial != nil && initial.HttpStatus != nil) {
		// Check if anything else is set.
		if dcl.AnySet() {
			des.HttpStatus = nil
			if initial != nil {
				initial.HttpStatus = nil
			}
		}
	}

	if des.Percentage != nil || (initial != nil && initial.Percentage != nil) {
		// Check if anything else is set.
		if dcl.AnySet() {
			des.Percentage = nil
			if initial != nil {
				initial.Percentage = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &GrpcRouteRulesActionFaultInjectionPolicyAbort{}

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

func canonicalizeGrpcRouteRulesActionFaultInjectionPolicyAbortSlice(des, initial []GrpcRouteRulesActionFaultInjectionPolicyAbort, opts ...dcl.ApplyOption) []GrpcRouteRulesActionFaultInjectionPolicyAbort {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]GrpcRouteRulesActionFaultInjectionPolicyAbort, 0, len(des))
		for _, d := range des {
			cd := canonicalizeGrpcRouteRulesActionFaultInjectionPolicyAbort(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]GrpcRouteRulesActionFaultInjectionPolicyAbort, 0, len(des))
	for i, d := range des {
		cd := canonicalizeGrpcRouteRulesActionFaultInjectionPolicyAbort(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewGrpcRouteRulesActionFaultInjectionPolicyAbort(c *Client, des, nw *GrpcRouteRulesActionFaultInjectionPolicyAbort) *GrpcRouteRulesActionFaultInjectionPolicyAbort {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for GrpcRouteRulesActionFaultInjectionPolicyAbort while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewGrpcRouteRulesActionFaultInjectionPolicyAbortSet(c *Client, des, nw []GrpcRouteRulesActionFaultInjectionPolicyAbort) []GrpcRouteRulesActionFaultInjectionPolicyAbort {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []GrpcRouteRulesActionFaultInjectionPolicyAbort
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareGrpcRouteRulesActionFaultInjectionPolicyAbortNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewGrpcRouteRulesActionFaultInjectionPolicyAbort(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewGrpcRouteRulesActionFaultInjectionPolicyAbortSlice(c *Client, des, nw []GrpcRouteRulesActionFaultInjectionPolicyAbort) []GrpcRouteRulesActionFaultInjectionPolicyAbort {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []GrpcRouteRulesActionFaultInjectionPolicyAbort
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGrpcRouteRulesActionFaultInjectionPolicyAbort(c, &d, &n))
	}

	return items
}

func canonicalizeGrpcRouteRulesActionRetryPolicy(des, initial *GrpcRouteRulesActionRetryPolicy, opts ...dcl.ApplyOption) *GrpcRouteRulesActionRetryPolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &GrpcRouteRulesActionRetryPolicy{}

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

	return cDes
}

func canonicalizeGrpcRouteRulesActionRetryPolicySlice(des, initial []GrpcRouteRulesActionRetryPolicy, opts ...dcl.ApplyOption) []GrpcRouteRulesActionRetryPolicy {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]GrpcRouteRulesActionRetryPolicy, 0, len(des))
		for _, d := range des {
			cd := canonicalizeGrpcRouteRulesActionRetryPolicy(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]GrpcRouteRulesActionRetryPolicy, 0, len(des))
	for i, d := range des {
		cd := canonicalizeGrpcRouteRulesActionRetryPolicy(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewGrpcRouteRulesActionRetryPolicy(c *Client, des, nw *GrpcRouteRulesActionRetryPolicy) *GrpcRouteRulesActionRetryPolicy {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for GrpcRouteRulesActionRetryPolicy while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.RetryConditions, nw.RetryConditions) {
		nw.RetryConditions = des.RetryConditions
	}

	return nw
}

func canonicalizeNewGrpcRouteRulesActionRetryPolicySet(c *Client, des, nw []GrpcRouteRulesActionRetryPolicy) []GrpcRouteRulesActionRetryPolicy {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []GrpcRouteRulesActionRetryPolicy
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareGrpcRouteRulesActionRetryPolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewGrpcRouteRulesActionRetryPolicy(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewGrpcRouteRulesActionRetryPolicySlice(c *Client, des, nw []GrpcRouteRulesActionRetryPolicy) []GrpcRouteRulesActionRetryPolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []GrpcRouteRulesActionRetryPolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGrpcRouteRulesActionRetryPolicy(c, &d, &n))
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
func diffGrpcRoute(c *Client, desired, actual *GrpcRoute, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Hostnames, actual.Hostnames, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("Hostnames")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Meshes, actual.Meshes, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("Meshes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Gateways, actual.Gateways, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("Gateways")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Rules, actual.Rules, dcl.DiffInfo{ObjectFunction: compareGrpcRouteRulesNewStyle, EmptyObject: EmptyGrpcRouteRules, OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("Rules")); len(ds) != 0 || err != nil {
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
func compareGrpcRouteRulesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GrpcRouteRules)
	if !ok {
		desiredNotPointer, ok := d.(GrpcRouteRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GrpcRouteRules or *GrpcRouteRules", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GrpcRouteRules)
	if !ok {
		actualNotPointer, ok := a.(GrpcRouteRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GrpcRouteRules", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Matches, actual.Matches, dcl.DiffInfo{ObjectFunction: compareGrpcRouteRulesMatchesNewStyle, EmptyObject: EmptyGrpcRouteRulesMatches, OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("Matches")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Action, actual.Action, dcl.DiffInfo{ObjectFunction: compareGrpcRouteRulesActionNewStyle, EmptyObject: EmptyGrpcRouteRulesAction, OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("Action")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGrpcRouteRulesMatchesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GrpcRouteRulesMatches)
	if !ok {
		desiredNotPointer, ok := d.(GrpcRouteRulesMatches)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GrpcRouteRulesMatches or *GrpcRouteRulesMatches", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GrpcRouteRulesMatches)
	if !ok {
		actualNotPointer, ok := a.(GrpcRouteRulesMatches)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GrpcRouteRulesMatches", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Method, actual.Method, dcl.DiffInfo{ObjectFunction: compareGrpcRouteRulesMatchesMethodNewStyle, EmptyObject: EmptyGrpcRouteRulesMatchesMethod, OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("Method")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Headers, actual.Headers, dcl.DiffInfo{ObjectFunction: compareGrpcRouteRulesMatchesHeadersNewStyle, EmptyObject: EmptyGrpcRouteRulesMatchesHeaders, OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("Headers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGrpcRouteRulesMatchesMethodNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GrpcRouteRulesMatchesMethod)
	if !ok {
		desiredNotPointer, ok := d.(GrpcRouteRulesMatchesMethod)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GrpcRouteRulesMatchesMethod or *GrpcRouteRulesMatchesMethod", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GrpcRouteRulesMatchesMethod)
	if !ok {
		actualNotPointer, ok := a.(GrpcRouteRulesMatchesMethod)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GrpcRouteRulesMatchesMethod", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GrpcService, actual.GrpcService, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("GrpcService")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GrpcMethod, actual.GrpcMethod, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("GrpcMethod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CaseSensitive, actual.CaseSensitive, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("CaseSensitive")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGrpcRouteRulesMatchesHeadersNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GrpcRouteRulesMatchesHeaders)
	if !ok {
		desiredNotPointer, ok := d.(GrpcRouteRulesMatchesHeaders)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GrpcRouteRulesMatchesHeaders or *GrpcRouteRulesMatchesHeaders", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GrpcRouteRulesMatchesHeaders)
	if !ok {
		actualNotPointer, ok := a.(GrpcRouteRulesMatchesHeaders)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GrpcRouteRulesMatchesHeaders", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Key, actual.Key, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("Key")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGrpcRouteRulesActionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GrpcRouteRulesAction)
	if !ok {
		desiredNotPointer, ok := d.(GrpcRouteRulesAction)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GrpcRouteRulesAction or *GrpcRouteRulesAction", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GrpcRouteRulesAction)
	if !ok {
		actualNotPointer, ok := a.(GrpcRouteRulesAction)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GrpcRouteRulesAction", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Destinations, actual.Destinations, dcl.DiffInfo{ObjectFunction: compareGrpcRouteRulesActionDestinationsNewStyle, EmptyObject: EmptyGrpcRouteRulesActionDestinations, OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("Destinations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FaultInjectionPolicy, actual.FaultInjectionPolicy, dcl.DiffInfo{ObjectFunction: compareGrpcRouteRulesActionFaultInjectionPolicyNewStyle, EmptyObject: EmptyGrpcRouteRulesActionFaultInjectionPolicy, OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("FaultInjectionPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Timeout, actual.Timeout, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("Timeout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RetryPolicy, actual.RetryPolicy, dcl.DiffInfo{ObjectFunction: compareGrpcRouteRulesActionRetryPolicyNewStyle, EmptyObject: EmptyGrpcRouteRulesActionRetryPolicy, OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("RetryPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGrpcRouteRulesActionDestinationsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GrpcRouteRulesActionDestinations)
	if !ok {
		desiredNotPointer, ok := d.(GrpcRouteRulesActionDestinations)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GrpcRouteRulesActionDestinations or *GrpcRouteRulesActionDestinations", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GrpcRouteRulesActionDestinations)
	if !ok {
		actualNotPointer, ok := a.(GrpcRouteRulesActionDestinations)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GrpcRouteRulesActionDestinations", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Weight, actual.Weight, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("Weight")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceName, actual.ServiceName, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("ServiceName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGrpcRouteRulesActionFaultInjectionPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GrpcRouteRulesActionFaultInjectionPolicy)
	if !ok {
		desiredNotPointer, ok := d.(GrpcRouteRulesActionFaultInjectionPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GrpcRouteRulesActionFaultInjectionPolicy or *GrpcRouteRulesActionFaultInjectionPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GrpcRouteRulesActionFaultInjectionPolicy)
	if !ok {
		actualNotPointer, ok := a.(GrpcRouteRulesActionFaultInjectionPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GrpcRouteRulesActionFaultInjectionPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Delay, actual.Delay, dcl.DiffInfo{ObjectFunction: compareGrpcRouteRulesActionFaultInjectionPolicyDelayNewStyle, EmptyObject: EmptyGrpcRouteRulesActionFaultInjectionPolicyDelay, OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("Delay")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Abort, actual.Abort, dcl.DiffInfo{ObjectFunction: compareGrpcRouteRulesActionFaultInjectionPolicyAbortNewStyle, EmptyObject: EmptyGrpcRouteRulesActionFaultInjectionPolicyAbort, OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("Abort")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGrpcRouteRulesActionFaultInjectionPolicyDelayNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GrpcRouteRulesActionFaultInjectionPolicyDelay)
	if !ok {
		desiredNotPointer, ok := d.(GrpcRouteRulesActionFaultInjectionPolicyDelay)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GrpcRouteRulesActionFaultInjectionPolicyDelay or *GrpcRouteRulesActionFaultInjectionPolicyDelay", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GrpcRouteRulesActionFaultInjectionPolicyDelay)
	if !ok {
		actualNotPointer, ok := a.(GrpcRouteRulesActionFaultInjectionPolicyDelay)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GrpcRouteRulesActionFaultInjectionPolicyDelay", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.FixedDelay, actual.FixedDelay, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("FixedDelay")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percentage, actual.Percentage, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("Percentage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGrpcRouteRulesActionFaultInjectionPolicyAbortNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GrpcRouteRulesActionFaultInjectionPolicyAbort)
	if !ok {
		desiredNotPointer, ok := d.(GrpcRouteRulesActionFaultInjectionPolicyAbort)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GrpcRouteRulesActionFaultInjectionPolicyAbort or *GrpcRouteRulesActionFaultInjectionPolicyAbort", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GrpcRouteRulesActionFaultInjectionPolicyAbort)
	if !ok {
		actualNotPointer, ok := a.(GrpcRouteRulesActionFaultInjectionPolicyAbort)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GrpcRouteRulesActionFaultInjectionPolicyAbort", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HttpStatus, actual.HttpStatus, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("HttpStatus")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percentage, actual.Percentage, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("Percentage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGrpcRouteRulesActionRetryPolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GrpcRouteRulesActionRetryPolicy)
	if !ok {
		desiredNotPointer, ok := d.(GrpcRouteRulesActionRetryPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GrpcRouteRulesActionRetryPolicy or *GrpcRouteRulesActionRetryPolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GrpcRouteRulesActionRetryPolicy)
	if !ok {
		actualNotPointer, ok := a.(GrpcRouteRulesActionRetryPolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GrpcRouteRulesActionRetryPolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RetryConditions, actual.RetryConditions, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("RetryConditions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NumRetries, actual.NumRetries, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateGrpcRouteUpdateGrpcRouteOperation")}, fn.AddNest("NumRetries")); len(ds) != 0 || err != nil {
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
func (r *GrpcRoute) urlNormalized() *GrpcRoute {
	normalized := dcl.Copy(*r).(GrpcRoute)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	return &normalized
}

func (r *GrpcRoute) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateGrpcRoute" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/grpcRoutes/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the GrpcRoute resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *GrpcRoute) marshal(c *Client) ([]byte, error) {
	m, err := expandGrpcRoute(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling GrpcRoute: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalGrpcRoute decodes JSON responses into the GrpcRoute resource schema.
func unmarshalGrpcRoute(b []byte, c *Client, res *GrpcRoute) (*GrpcRoute, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapGrpcRoute(m, c, res)
}

func unmarshalMapGrpcRoute(m map[string]interface{}, c *Client, res *GrpcRoute) (*GrpcRoute, error) {

	flattened := flattenGrpcRoute(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandGrpcRoute expands GrpcRoute into a JSON request object.
func expandGrpcRoute(c *Client, f *GrpcRoute) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("projects/%s/locations/global/grpcRoutes/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Labels; dcl.ValueShouldBeSent(v) {
		m["labels"] = v
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
	if v, err := expandGrpcRouteRulesSlice(c, f.Rules, res); err != nil {
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

// flattenGrpcRoute flattens GrpcRoute from a JSON request object into the
// GrpcRoute type.
func flattenGrpcRoute(c *Client, i interface{}, res *GrpcRoute) *GrpcRoute {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &GrpcRoute{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.Hostnames = dcl.FlattenStringSlice(m["hostnames"])
	resultRes.Meshes = dcl.FlattenStringSlice(m["meshes"])
	resultRes.Gateways = dcl.FlattenStringSlice(m["gateways"])
	resultRes.Rules = flattenGrpcRouteRulesSlice(c, m["rules"], res)
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])
	resultRes.SelfLink = dcl.FlattenString(m["selfLink"])

	return resultRes
}

// expandGrpcRouteRulesMap expands the contents of GrpcRouteRules into a JSON
// request object.
func expandGrpcRouteRulesMap(c *Client, f map[string]GrpcRouteRules, res *GrpcRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGrpcRouteRules(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGrpcRouteRulesSlice expands the contents of GrpcRouteRules into a JSON
// request object.
func expandGrpcRouteRulesSlice(c *Client, f []GrpcRouteRules, res *GrpcRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGrpcRouteRules(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGrpcRouteRulesMap flattens the contents of GrpcRouteRules from a JSON
// response object.
func flattenGrpcRouteRulesMap(c *Client, i interface{}, res *GrpcRoute) map[string]GrpcRouteRules {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GrpcRouteRules{}
	}

	if len(a) == 0 {
		return map[string]GrpcRouteRules{}
	}

	items := make(map[string]GrpcRouteRules)
	for k, item := range a {
		items[k] = *flattenGrpcRouteRules(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenGrpcRouteRulesSlice flattens the contents of GrpcRouteRules from a JSON
// response object.
func flattenGrpcRouteRulesSlice(c *Client, i interface{}, res *GrpcRoute) []GrpcRouteRules {
	a, ok := i.([]interface{})
	if !ok {
		return []GrpcRouteRules{}
	}

	if len(a) == 0 {
		return []GrpcRouteRules{}
	}

	items := make([]GrpcRouteRules, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGrpcRouteRules(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandGrpcRouteRules expands an instance of GrpcRouteRules into a JSON
// request object.
func expandGrpcRouteRules(c *Client, f *GrpcRouteRules, res *GrpcRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandGrpcRouteRulesMatchesSlice(c, f.Matches, res); err != nil {
		return nil, fmt.Errorf("error expanding Matches into matches: %w", err)
	} else if v != nil {
		m["matches"] = v
	}
	if v, err := expandGrpcRouteRulesAction(c, f.Action, res); err != nil {
		return nil, fmt.Errorf("error expanding Action into action: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["action"] = v
	}

	return m, nil
}

// flattenGrpcRouteRules flattens an instance of GrpcRouteRules from a JSON
// response object.
func flattenGrpcRouteRules(c *Client, i interface{}, res *GrpcRoute) *GrpcRouteRules {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GrpcRouteRules{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGrpcRouteRules
	}
	r.Matches = flattenGrpcRouteRulesMatchesSlice(c, m["matches"], res)
	r.Action = flattenGrpcRouteRulesAction(c, m["action"], res)

	return r
}

// expandGrpcRouteRulesMatchesMap expands the contents of GrpcRouteRulesMatches into a JSON
// request object.
func expandGrpcRouteRulesMatchesMap(c *Client, f map[string]GrpcRouteRulesMatches, res *GrpcRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGrpcRouteRulesMatches(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGrpcRouteRulesMatchesSlice expands the contents of GrpcRouteRulesMatches into a JSON
// request object.
func expandGrpcRouteRulesMatchesSlice(c *Client, f []GrpcRouteRulesMatches, res *GrpcRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGrpcRouteRulesMatches(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGrpcRouteRulesMatchesMap flattens the contents of GrpcRouteRulesMatches from a JSON
// response object.
func flattenGrpcRouteRulesMatchesMap(c *Client, i interface{}, res *GrpcRoute) map[string]GrpcRouteRulesMatches {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GrpcRouteRulesMatches{}
	}

	if len(a) == 0 {
		return map[string]GrpcRouteRulesMatches{}
	}

	items := make(map[string]GrpcRouteRulesMatches)
	for k, item := range a {
		items[k] = *flattenGrpcRouteRulesMatches(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenGrpcRouteRulesMatchesSlice flattens the contents of GrpcRouteRulesMatches from a JSON
// response object.
func flattenGrpcRouteRulesMatchesSlice(c *Client, i interface{}, res *GrpcRoute) []GrpcRouteRulesMatches {
	a, ok := i.([]interface{})
	if !ok {
		return []GrpcRouteRulesMatches{}
	}

	if len(a) == 0 {
		return []GrpcRouteRulesMatches{}
	}

	items := make([]GrpcRouteRulesMatches, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGrpcRouteRulesMatches(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandGrpcRouteRulesMatches expands an instance of GrpcRouteRulesMatches into a JSON
// request object.
func expandGrpcRouteRulesMatches(c *Client, f *GrpcRouteRulesMatches, res *GrpcRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandGrpcRouteRulesMatchesMethod(c, f.Method, res); err != nil {
		return nil, fmt.Errorf("error expanding Method into method: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["method"] = v
	}
	if v, err := expandGrpcRouteRulesMatchesHeadersSlice(c, f.Headers, res); err != nil {
		return nil, fmt.Errorf("error expanding Headers into headers: %w", err)
	} else if v != nil {
		m["headers"] = v
	}

	return m, nil
}

// flattenGrpcRouteRulesMatches flattens an instance of GrpcRouteRulesMatches from a JSON
// response object.
func flattenGrpcRouteRulesMatches(c *Client, i interface{}, res *GrpcRoute) *GrpcRouteRulesMatches {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GrpcRouteRulesMatches{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGrpcRouteRulesMatches
	}
	r.Method = flattenGrpcRouteRulesMatchesMethod(c, m["method"], res)
	r.Headers = flattenGrpcRouteRulesMatchesHeadersSlice(c, m["headers"], res)

	return r
}

// expandGrpcRouteRulesMatchesMethodMap expands the contents of GrpcRouteRulesMatchesMethod into a JSON
// request object.
func expandGrpcRouteRulesMatchesMethodMap(c *Client, f map[string]GrpcRouteRulesMatchesMethod, res *GrpcRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGrpcRouteRulesMatchesMethod(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGrpcRouteRulesMatchesMethodSlice expands the contents of GrpcRouteRulesMatchesMethod into a JSON
// request object.
func expandGrpcRouteRulesMatchesMethodSlice(c *Client, f []GrpcRouteRulesMatchesMethod, res *GrpcRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGrpcRouteRulesMatchesMethod(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGrpcRouteRulesMatchesMethodMap flattens the contents of GrpcRouteRulesMatchesMethod from a JSON
// response object.
func flattenGrpcRouteRulesMatchesMethodMap(c *Client, i interface{}, res *GrpcRoute) map[string]GrpcRouteRulesMatchesMethod {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GrpcRouteRulesMatchesMethod{}
	}

	if len(a) == 0 {
		return map[string]GrpcRouteRulesMatchesMethod{}
	}

	items := make(map[string]GrpcRouteRulesMatchesMethod)
	for k, item := range a {
		items[k] = *flattenGrpcRouteRulesMatchesMethod(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenGrpcRouteRulesMatchesMethodSlice flattens the contents of GrpcRouteRulesMatchesMethod from a JSON
// response object.
func flattenGrpcRouteRulesMatchesMethodSlice(c *Client, i interface{}, res *GrpcRoute) []GrpcRouteRulesMatchesMethod {
	a, ok := i.([]interface{})
	if !ok {
		return []GrpcRouteRulesMatchesMethod{}
	}

	if len(a) == 0 {
		return []GrpcRouteRulesMatchesMethod{}
	}

	items := make([]GrpcRouteRulesMatchesMethod, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGrpcRouteRulesMatchesMethod(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandGrpcRouteRulesMatchesMethod expands an instance of GrpcRouteRulesMatchesMethod into a JSON
// request object.
func expandGrpcRouteRulesMatchesMethod(c *Client, f *GrpcRouteRulesMatchesMethod, res *GrpcRoute) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.GrpcService; !dcl.IsEmptyValueIndirect(v) {
		m["grpcService"] = v
	}
	if v := f.GrpcMethod; !dcl.IsEmptyValueIndirect(v) {
		m["grpcMethod"] = v
	}
	if v := f.CaseSensitive; !dcl.IsEmptyValueIndirect(v) {
		m["caseSensitive"] = v
	}

	return m, nil
}

// flattenGrpcRouteRulesMatchesMethod flattens an instance of GrpcRouteRulesMatchesMethod from a JSON
// response object.
func flattenGrpcRouteRulesMatchesMethod(c *Client, i interface{}, res *GrpcRoute) *GrpcRouteRulesMatchesMethod {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GrpcRouteRulesMatchesMethod{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGrpcRouteRulesMatchesMethod
	}
	r.Type = flattenGrpcRouteRulesMatchesMethodTypeEnum(m["type"])
	r.GrpcService = dcl.FlattenString(m["grpcService"])
	r.GrpcMethod = dcl.FlattenString(m["grpcMethod"])
	r.CaseSensitive = dcl.FlattenBool(m["caseSensitive"])

	return r
}

// expandGrpcRouteRulesMatchesHeadersMap expands the contents of GrpcRouteRulesMatchesHeaders into a JSON
// request object.
func expandGrpcRouteRulesMatchesHeadersMap(c *Client, f map[string]GrpcRouteRulesMatchesHeaders, res *GrpcRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGrpcRouteRulesMatchesHeaders(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGrpcRouteRulesMatchesHeadersSlice expands the contents of GrpcRouteRulesMatchesHeaders into a JSON
// request object.
func expandGrpcRouteRulesMatchesHeadersSlice(c *Client, f []GrpcRouteRulesMatchesHeaders, res *GrpcRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGrpcRouteRulesMatchesHeaders(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGrpcRouteRulesMatchesHeadersMap flattens the contents of GrpcRouteRulesMatchesHeaders from a JSON
// response object.
func flattenGrpcRouteRulesMatchesHeadersMap(c *Client, i interface{}, res *GrpcRoute) map[string]GrpcRouteRulesMatchesHeaders {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GrpcRouteRulesMatchesHeaders{}
	}

	if len(a) == 0 {
		return map[string]GrpcRouteRulesMatchesHeaders{}
	}

	items := make(map[string]GrpcRouteRulesMatchesHeaders)
	for k, item := range a {
		items[k] = *flattenGrpcRouteRulesMatchesHeaders(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenGrpcRouteRulesMatchesHeadersSlice flattens the contents of GrpcRouteRulesMatchesHeaders from a JSON
// response object.
func flattenGrpcRouteRulesMatchesHeadersSlice(c *Client, i interface{}, res *GrpcRoute) []GrpcRouteRulesMatchesHeaders {
	a, ok := i.([]interface{})
	if !ok {
		return []GrpcRouteRulesMatchesHeaders{}
	}

	if len(a) == 0 {
		return []GrpcRouteRulesMatchesHeaders{}
	}

	items := make([]GrpcRouteRulesMatchesHeaders, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGrpcRouteRulesMatchesHeaders(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandGrpcRouteRulesMatchesHeaders expands an instance of GrpcRouteRulesMatchesHeaders into a JSON
// request object.
func expandGrpcRouteRulesMatchesHeaders(c *Client, f *GrpcRouteRulesMatchesHeaders, res *GrpcRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.Key; !dcl.IsEmptyValueIndirect(v) {
		m["key"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenGrpcRouteRulesMatchesHeaders flattens an instance of GrpcRouteRulesMatchesHeaders from a JSON
// response object.
func flattenGrpcRouteRulesMatchesHeaders(c *Client, i interface{}, res *GrpcRoute) *GrpcRouteRulesMatchesHeaders {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GrpcRouteRulesMatchesHeaders{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGrpcRouteRulesMatchesHeaders
	}
	r.Type = flattenGrpcRouteRulesMatchesHeadersTypeEnum(m["type"])
	r.Key = dcl.FlattenString(m["key"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// expandGrpcRouteRulesActionMap expands the contents of GrpcRouteRulesAction into a JSON
// request object.
func expandGrpcRouteRulesActionMap(c *Client, f map[string]GrpcRouteRulesAction, res *GrpcRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGrpcRouteRulesAction(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGrpcRouteRulesActionSlice expands the contents of GrpcRouteRulesAction into a JSON
// request object.
func expandGrpcRouteRulesActionSlice(c *Client, f []GrpcRouteRulesAction, res *GrpcRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGrpcRouteRulesAction(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGrpcRouteRulesActionMap flattens the contents of GrpcRouteRulesAction from a JSON
// response object.
func flattenGrpcRouteRulesActionMap(c *Client, i interface{}, res *GrpcRoute) map[string]GrpcRouteRulesAction {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GrpcRouteRulesAction{}
	}

	if len(a) == 0 {
		return map[string]GrpcRouteRulesAction{}
	}

	items := make(map[string]GrpcRouteRulesAction)
	for k, item := range a {
		items[k] = *flattenGrpcRouteRulesAction(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenGrpcRouteRulesActionSlice flattens the contents of GrpcRouteRulesAction from a JSON
// response object.
func flattenGrpcRouteRulesActionSlice(c *Client, i interface{}, res *GrpcRoute) []GrpcRouteRulesAction {
	a, ok := i.([]interface{})
	if !ok {
		return []GrpcRouteRulesAction{}
	}

	if len(a) == 0 {
		return []GrpcRouteRulesAction{}
	}

	items := make([]GrpcRouteRulesAction, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGrpcRouteRulesAction(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandGrpcRouteRulesAction expands an instance of GrpcRouteRulesAction into a JSON
// request object.
func expandGrpcRouteRulesAction(c *Client, f *GrpcRouteRulesAction, res *GrpcRoute) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandGrpcRouteRulesActionDestinationsSlice(c, f.Destinations, res); err != nil {
		return nil, fmt.Errorf("error expanding Destinations into destinations: %w", err)
	} else if v != nil {
		m["destinations"] = v
	}
	if v, err := expandGrpcRouteRulesActionFaultInjectionPolicy(c, f.FaultInjectionPolicy, res); err != nil {
		return nil, fmt.Errorf("error expanding FaultInjectionPolicy into faultInjectionPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["faultInjectionPolicy"] = v
	}
	if v := f.Timeout; !dcl.IsEmptyValueIndirect(v) {
		m["timeout"] = v
	}
	if v, err := expandGrpcRouteRulesActionRetryPolicy(c, f.RetryPolicy, res); err != nil {
		return nil, fmt.Errorf("error expanding RetryPolicy into retryPolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["retryPolicy"] = v
	}

	return m, nil
}

// flattenGrpcRouteRulesAction flattens an instance of GrpcRouteRulesAction from a JSON
// response object.
func flattenGrpcRouteRulesAction(c *Client, i interface{}, res *GrpcRoute) *GrpcRouteRulesAction {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GrpcRouteRulesAction{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGrpcRouteRulesAction
	}
	r.Destinations = flattenGrpcRouteRulesActionDestinationsSlice(c, m["destinations"], res)
	r.FaultInjectionPolicy = flattenGrpcRouteRulesActionFaultInjectionPolicy(c, m["faultInjectionPolicy"], res)
	r.Timeout = dcl.FlattenString(m["timeout"])
	r.RetryPolicy = flattenGrpcRouteRulesActionRetryPolicy(c, m["retryPolicy"], res)

	return r
}

// expandGrpcRouteRulesActionDestinationsMap expands the contents of GrpcRouteRulesActionDestinations into a JSON
// request object.
func expandGrpcRouteRulesActionDestinationsMap(c *Client, f map[string]GrpcRouteRulesActionDestinations, res *GrpcRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGrpcRouteRulesActionDestinations(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGrpcRouteRulesActionDestinationsSlice expands the contents of GrpcRouteRulesActionDestinations into a JSON
// request object.
func expandGrpcRouteRulesActionDestinationsSlice(c *Client, f []GrpcRouteRulesActionDestinations, res *GrpcRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGrpcRouteRulesActionDestinations(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGrpcRouteRulesActionDestinationsMap flattens the contents of GrpcRouteRulesActionDestinations from a JSON
// response object.
func flattenGrpcRouteRulesActionDestinationsMap(c *Client, i interface{}, res *GrpcRoute) map[string]GrpcRouteRulesActionDestinations {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GrpcRouteRulesActionDestinations{}
	}

	if len(a) == 0 {
		return map[string]GrpcRouteRulesActionDestinations{}
	}

	items := make(map[string]GrpcRouteRulesActionDestinations)
	for k, item := range a {
		items[k] = *flattenGrpcRouteRulesActionDestinations(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenGrpcRouteRulesActionDestinationsSlice flattens the contents of GrpcRouteRulesActionDestinations from a JSON
// response object.
func flattenGrpcRouteRulesActionDestinationsSlice(c *Client, i interface{}, res *GrpcRoute) []GrpcRouteRulesActionDestinations {
	a, ok := i.([]interface{})
	if !ok {
		return []GrpcRouteRulesActionDestinations{}
	}

	if len(a) == 0 {
		return []GrpcRouteRulesActionDestinations{}
	}

	items := make([]GrpcRouteRulesActionDestinations, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGrpcRouteRulesActionDestinations(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandGrpcRouteRulesActionDestinations expands an instance of GrpcRouteRulesActionDestinations into a JSON
// request object.
func expandGrpcRouteRulesActionDestinations(c *Client, f *GrpcRouteRulesActionDestinations, res *GrpcRoute) (map[string]interface{}, error) {
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

// flattenGrpcRouteRulesActionDestinations flattens an instance of GrpcRouteRulesActionDestinations from a JSON
// response object.
func flattenGrpcRouteRulesActionDestinations(c *Client, i interface{}, res *GrpcRoute) *GrpcRouteRulesActionDestinations {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GrpcRouteRulesActionDestinations{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGrpcRouteRulesActionDestinations
	}
	r.Weight = dcl.FlattenInteger(m["weight"])
	r.ServiceName = dcl.FlattenString(m["serviceName"])

	return r
}

// expandGrpcRouteRulesActionFaultInjectionPolicyMap expands the contents of GrpcRouteRulesActionFaultInjectionPolicy into a JSON
// request object.
func expandGrpcRouteRulesActionFaultInjectionPolicyMap(c *Client, f map[string]GrpcRouteRulesActionFaultInjectionPolicy, res *GrpcRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGrpcRouteRulesActionFaultInjectionPolicy(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGrpcRouteRulesActionFaultInjectionPolicySlice expands the contents of GrpcRouteRulesActionFaultInjectionPolicy into a JSON
// request object.
func expandGrpcRouteRulesActionFaultInjectionPolicySlice(c *Client, f []GrpcRouteRulesActionFaultInjectionPolicy, res *GrpcRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGrpcRouteRulesActionFaultInjectionPolicy(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGrpcRouteRulesActionFaultInjectionPolicyMap flattens the contents of GrpcRouteRulesActionFaultInjectionPolicy from a JSON
// response object.
func flattenGrpcRouteRulesActionFaultInjectionPolicyMap(c *Client, i interface{}, res *GrpcRoute) map[string]GrpcRouteRulesActionFaultInjectionPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GrpcRouteRulesActionFaultInjectionPolicy{}
	}

	if len(a) == 0 {
		return map[string]GrpcRouteRulesActionFaultInjectionPolicy{}
	}

	items := make(map[string]GrpcRouteRulesActionFaultInjectionPolicy)
	for k, item := range a {
		items[k] = *flattenGrpcRouteRulesActionFaultInjectionPolicy(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenGrpcRouteRulesActionFaultInjectionPolicySlice flattens the contents of GrpcRouteRulesActionFaultInjectionPolicy from a JSON
// response object.
func flattenGrpcRouteRulesActionFaultInjectionPolicySlice(c *Client, i interface{}, res *GrpcRoute) []GrpcRouteRulesActionFaultInjectionPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []GrpcRouteRulesActionFaultInjectionPolicy{}
	}

	if len(a) == 0 {
		return []GrpcRouteRulesActionFaultInjectionPolicy{}
	}

	items := make([]GrpcRouteRulesActionFaultInjectionPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGrpcRouteRulesActionFaultInjectionPolicy(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandGrpcRouteRulesActionFaultInjectionPolicy expands an instance of GrpcRouteRulesActionFaultInjectionPolicy into a JSON
// request object.
func expandGrpcRouteRulesActionFaultInjectionPolicy(c *Client, f *GrpcRouteRulesActionFaultInjectionPolicy, res *GrpcRoute) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandGrpcRouteRulesActionFaultInjectionPolicyDelay(c, f.Delay, res); err != nil {
		return nil, fmt.Errorf("error expanding Delay into delay: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["delay"] = v
	}
	if v, err := expandGrpcRouteRulesActionFaultInjectionPolicyAbort(c, f.Abort, res); err != nil {
		return nil, fmt.Errorf("error expanding Abort into abort: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["abort"] = v
	}

	return m, nil
}

// flattenGrpcRouteRulesActionFaultInjectionPolicy flattens an instance of GrpcRouteRulesActionFaultInjectionPolicy from a JSON
// response object.
func flattenGrpcRouteRulesActionFaultInjectionPolicy(c *Client, i interface{}, res *GrpcRoute) *GrpcRouteRulesActionFaultInjectionPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GrpcRouteRulesActionFaultInjectionPolicy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGrpcRouteRulesActionFaultInjectionPolicy
	}
	r.Delay = flattenGrpcRouteRulesActionFaultInjectionPolicyDelay(c, m["delay"], res)
	r.Abort = flattenGrpcRouteRulesActionFaultInjectionPolicyAbort(c, m["abort"], res)

	return r
}

// expandGrpcRouteRulesActionFaultInjectionPolicyDelayMap expands the contents of GrpcRouteRulesActionFaultInjectionPolicyDelay into a JSON
// request object.
func expandGrpcRouteRulesActionFaultInjectionPolicyDelayMap(c *Client, f map[string]GrpcRouteRulesActionFaultInjectionPolicyDelay, res *GrpcRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGrpcRouteRulesActionFaultInjectionPolicyDelay(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGrpcRouteRulesActionFaultInjectionPolicyDelaySlice expands the contents of GrpcRouteRulesActionFaultInjectionPolicyDelay into a JSON
// request object.
func expandGrpcRouteRulesActionFaultInjectionPolicyDelaySlice(c *Client, f []GrpcRouteRulesActionFaultInjectionPolicyDelay, res *GrpcRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGrpcRouteRulesActionFaultInjectionPolicyDelay(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGrpcRouteRulesActionFaultInjectionPolicyDelayMap flattens the contents of GrpcRouteRulesActionFaultInjectionPolicyDelay from a JSON
// response object.
func flattenGrpcRouteRulesActionFaultInjectionPolicyDelayMap(c *Client, i interface{}, res *GrpcRoute) map[string]GrpcRouteRulesActionFaultInjectionPolicyDelay {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GrpcRouteRulesActionFaultInjectionPolicyDelay{}
	}

	if len(a) == 0 {
		return map[string]GrpcRouteRulesActionFaultInjectionPolicyDelay{}
	}

	items := make(map[string]GrpcRouteRulesActionFaultInjectionPolicyDelay)
	for k, item := range a {
		items[k] = *flattenGrpcRouteRulesActionFaultInjectionPolicyDelay(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenGrpcRouteRulesActionFaultInjectionPolicyDelaySlice flattens the contents of GrpcRouteRulesActionFaultInjectionPolicyDelay from a JSON
// response object.
func flattenGrpcRouteRulesActionFaultInjectionPolicyDelaySlice(c *Client, i interface{}, res *GrpcRoute) []GrpcRouteRulesActionFaultInjectionPolicyDelay {
	a, ok := i.([]interface{})
	if !ok {
		return []GrpcRouteRulesActionFaultInjectionPolicyDelay{}
	}

	if len(a) == 0 {
		return []GrpcRouteRulesActionFaultInjectionPolicyDelay{}
	}

	items := make([]GrpcRouteRulesActionFaultInjectionPolicyDelay, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGrpcRouteRulesActionFaultInjectionPolicyDelay(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandGrpcRouteRulesActionFaultInjectionPolicyDelay expands an instance of GrpcRouteRulesActionFaultInjectionPolicyDelay into a JSON
// request object.
func expandGrpcRouteRulesActionFaultInjectionPolicyDelay(c *Client, f *GrpcRouteRulesActionFaultInjectionPolicyDelay, res *GrpcRoute) (map[string]interface{}, error) {
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

// flattenGrpcRouteRulesActionFaultInjectionPolicyDelay flattens an instance of GrpcRouteRulesActionFaultInjectionPolicyDelay from a JSON
// response object.
func flattenGrpcRouteRulesActionFaultInjectionPolicyDelay(c *Client, i interface{}, res *GrpcRoute) *GrpcRouteRulesActionFaultInjectionPolicyDelay {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GrpcRouteRulesActionFaultInjectionPolicyDelay{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGrpcRouteRulesActionFaultInjectionPolicyDelay
	}
	r.FixedDelay = dcl.FlattenString(m["fixedDelay"])
	r.Percentage = dcl.FlattenInteger(m["percentage"])

	return r
}

// expandGrpcRouteRulesActionFaultInjectionPolicyAbortMap expands the contents of GrpcRouteRulesActionFaultInjectionPolicyAbort into a JSON
// request object.
func expandGrpcRouteRulesActionFaultInjectionPolicyAbortMap(c *Client, f map[string]GrpcRouteRulesActionFaultInjectionPolicyAbort, res *GrpcRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGrpcRouteRulesActionFaultInjectionPolicyAbort(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGrpcRouteRulesActionFaultInjectionPolicyAbortSlice expands the contents of GrpcRouteRulesActionFaultInjectionPolicyAbort into a JSON
// request object.
func expandGrpcRouteRulesActionFaultInjectionPolicyAbortSlice(c *Client, f []GrpcRouteRulesActionFaultInjectionPolicyAbort, res *GrpcRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGrpcRouteRulesActionFaultInjectionPolicyAbort(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGrpcRouteRulesActionFaultInjectionPolicyAbortMap flattens the contents of GrpcRouteRulesActionFaultInjectionPolicyAbort from a JSON
// response object.
func flattenGrpcRouteRulesActionFaultInjectionPolicyAbortMap(c *Client, i interface{}, res *GrpcRoute) map[string]GrpcRouteRulesActionFaultInjectionPolicyAbort {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GrpcRouteRulesActionFaultInjectionPolicyAbort{}
	}

	if len(a) == 0 {
		return map[string]GrpcRouteRulesActionFaultInjectionPolicyAbort{}
	}

	items := make(map[string]GrpcRouteRulesActionFaultInjectionPolicyAbort)
	for k, item := range a {
		items[k] = *flattenGrpcRouteRulesActionFaultInjectionPolicyAbort(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenGrpcRouteRulesActionFaultInjectionPolicyAbortSlice flattens the contents of GrpcRouteRulesActionFaultInjectionPolicyAbort from a JSON
// response object.
func flattenGrpcRouteRulesActionFaultInjectionPolicyAbortSlice(c *Client, i interface{}, res *GrpcRoute) []GrpcRouteRulesActionFaultInjectionPolicyAbort {
	a, ok := i.([]interface{})
	if !ok {
		return []GrpcRouteRulesActionFaultInjectionPolicyAbort{}
	}

	if len(a) == 0 {
		return []GrpcRouteRulesActionFaultInjectionPolicyAbort{}
	}

	items := make([]GrpcRouteRulesActionFaultInjectionPolicyAbort, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGrpcRouteRulesActionFaultInjectionPolicyAbort(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandGrpcRouteRulesActionFaultInjectionPolicyAbort expands an instance of GrpcRouteRulesActionFaultInjectionPolicyAbort into a JSON
// request object.
func expandGrpcRouteRulesActionFaultInjectionPolicyAbort(c *Client, f *GrpcRouteRulesActionFaultInjectionPolicyAbort, res *GrpcRoute) (map[string]interface{}, error) {
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

// flattenGrpcRouteRulesActionFaultInjectionPolicyAbort flattens an instance of GrpcRouteRulesActionFaultInjectionPolicyAbort from a JSON
// response object.
func flattenGrpcRouteRulesActionFaultInjectionPolicyAbort(c *Client, i interface{}, res *GrpcRoute) *GrpcRouteRulesActionFaultInjectionPolicyAbort {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GrpcRouteRulesActionFaultInjectionPolicyAbort{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGrpcRouteRulesActionFaultInjectionPolicyAbort
	}
	r.HttpStatus = dcl.FlattenInteger(m["httpStatus"])
	r.Percentage = dcl.FlattenInteger(m["percentage"])

	return r
}

// expandGrpcRouteRulesActionRetryPolicyMap expands the contents of GrpcRouteRulesActionRetryPolicy into a JSON
// request object.
func expandGrpcRouteRulesActionRetryPolicyMap(c *Client, f map[string]GrpcRouteRulesActionRetryPolicy, res *GrpcRoute) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGrpcRouteRulesActionRetryPolicy(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGrpcRouteRulesActionRetryPolicySlice expands the contents of GrpcRouteRulesActionRetryPolicy into a JSON
// request object.
func expandGrpcRouteRulesActionRetryPolicySlice(c *Client, f []GrpcRouteRulesActionRetryPolicy, res *GrpcRoute) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGrpcRouteRulesActionRetryPolicy(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGrpcRouteRulesActionRetryPolicyMap flattens the contents of GrpcRouteRulesActionRetryPolicy from a JSON
// response object.
func flattenGrpcRouteRulesActionRetryPolicyMap(c *Client, i interface{}, res *GrpcRoute) map[string]GrpcRouteRulesActionRetryPolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GrpcRouteRulesActionRetryPolicy{}
	}

	if len(a) == 0 {
		return map[string]GrpcRouteRulesActionRetryPolicy{}
	}

	items := make(map[string]GrpcRouteRulesActionRetryPolicy)
	for k, item := range a {
		items[k] = *flattenGrpcRouteRulesActionRetryPolicy(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenGrpcRouteRulesActionRetryPolicySlice flattens the contents of GrpcRouteRulesActionRetryPolicy from a JSON
// response object.
func flattenGrpcRouteRulesActionRetryPolicySlice(c *Client, i interface{}, res *GrpcRoute) []GrpcRouteRulesActionRetryPolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []GrpcRouteRulesActionRetryPolicy{}
	}

	if len(a) == 0 {
		return []GrpcRouteRulesActionRetryPolicy{}
	}

	items := make([]GrpcRouteRulesActionRetryPolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGrpcRouteRulesActionRetryPolicy(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandGrpcRouteRulesActionRetryPolicy expands an instance of GrpcRouteRulesActionRetryPolicy into a JSON
// request object.
func expandGrpcRouteRulesActionRetryPolicy(c *Client, f *GrpcRouteRulesActionRetryPolicy, res *GrpcRoute) (map[string]interface{}, error) {
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

	return m, nil
}

// flattenGrpcRouteRulesActionRetryPolicy flattens an instance of GrpcRouteRulesActionRetryPolicy from a JSON
// response object.
func flattenGrpcRouteRulesActionRetryPolicy(c *Client, i interface{}, res *GrpcRoute) *GrpcRouteRulesActionRetryPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GrpcRouteRulesActionRetryPolicy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGrpcRouteRulesActionRetryPolicy
	}
	r.RetryConditions = dcl.FlattenStringSlice(m["retryConditions"])
	r.NumRetries = dcl.FlattenInteger(m["numRetries"])

	return r
}

// flattenGrpcRouteRulesMatchesMethodTypeEnumMap flattens the contents of GrpcRouteRulesMatchesMethodTypeEnum from a JSON
// response object.
func flattenGrpcRouteRulesMatchesMethodTypeEnumMap(c *Client, i interface{}, res *GrpcRoute) map[string]GrpcRouteRulesMatchesMethodTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GrpcRouteRulesMatchesMethodTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]GrpcRouteRulesMatchesMethodTypeEnum{}
	}

	items := make(map[string]GrpcRouteRulesMatchesMethodTypeEnum)
	for k, item := range a {
		items[k] = *flattenGrpcRouteRulesMatchesMethodTypeEnum(item.(interface{}))
	}

	return items
}

// flattenGrpcRouteRulesMatchesMethodTypeEnumSlice flattens the contents of GrpcRouteRulesMatchesMethodTypeEnum from a JSON
// response object.
func flattenGrpcRouteRulesMatchesMethodTypeEnumSlice(c *Client, i interface{}, res *GrpcRoute) []GrpcRouteRulesMatchesMethodTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []GrpcRouteRulesMatchesMethodTypeEnum{}
	}

	if len(a) == 0 {
		return []GrpcRouteRulesMatchesMethodTypeEnum{}
	}

	items := make([]GrpcRouteRulesMatchesMethodTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGrpcRouteRulesMatchesMethodTypeEnum(item.(interface{})))
	}

	return items
}

// flattenGrpcRouteRulesMatchesMethodTypeEnum asserts that an interface is a string, and returns a
// pointer to a *GrpcRouteRulesMatchesMethodTypeEnum with the same value as that string.
func flattenGrpcRouteRulesMatchesMethodTypeEnum(i interface{}) *GrpcRouteRulesMatchesMethodTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return GrpcRouteRulesMatchesMethodTypeEnumRef(s)
}

// flattenGrpcRouteRulesMatchesHeadersTypeEnumMap flattens the contents of GrpcRouteRulesMatchesHeadersTypeEnum from a JSON
// response object.
func flattenGrpcRouteRulesMatchesHeadersTypeEnumMap(c *Client, i interface{}, res *GrpcRoute) map[string]GrpcRouteRulesMatchesHeadersTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GrpcRouteRulesMatchesHeadersTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]GrpcRouteRulesMatchesHeadersTypeEnum{}
	}

	items := make(map[string]GrpcRouteRulesMatchesHeadersTypeEnum)
	for k, item := range a {
		items[k] = *flattenGrpcRouteRulesMatchesHeadersTypeEnum(item.(interface{}))
	}

	return items
}

// flattenGrpcRouteRulesMatchesHeadersTypeEnumSlice flattens the contents of GrpcRouteRulesMatchesHeadersTypeEnum from a JSON
// response object.
func flattenGrpcRouteRulesMatchesHeadersTypeEnumSlice(c *Client, i interface{}, res *GrpcRoute) []GrpcRouteRulesMatchesHeadersTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []GrpcRouteRulesMatchesHeadersTypeEnum{}
	}

	if len(a) == 0 {
		return []GrpcRouteRulesMatchesHeadersTypeEnum{}
	}

	items := make([]GrpcRouteRulesMatchesHeadersTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGrpcRouteRulesMatchesHeadersTypeEnum(item.(interface{})))
	}

	return items
}

// flattenGrpcRouteRulesMatchesHeadersTypeEnum asserts that an interface is a string, and returns a
// pointer to a *GrpcRouteRulesMatchesHeadersTypeEnum with the same value as that string.
func flattenGrpcRouteRulesMatchesHeadersTypeEnum(i interface{}) *GrpcRouteRulesMatchesHeadersTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return GrpcRouteRulesMatchesHeadersTypeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *GrpcRoute) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalGrpcRoute(b, c, r)
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

type grpcRouteDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         grpcRouteApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToGrpcRouteDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]grpcRouteDiff, error) {
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
	var diffs []grpcRouteDiff
	// For each operation name, create a grpcRouteDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := grpcRouteDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToGrpcRouteApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToGrpcRouteApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (grpcRouteApiOperation, error) {
	switch opName {

	case "updateGrpcRouteUpdateGrpcRouteOperation":
		return &updateGrpcRouteUpdateGrpcRouteOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractGrpcRouteFields(r *GrpcRoute) error {
	return nil
}
func extractGrpcRouteRulesFields(r *GrpcRoute, o *GrpcRouteRules) error {
	vAction := o.Action
	if vAction == nil {
		// note: explicitly not the empty object.
		vAction = &GrpcRouteRulesAction{}
	}
	if err := extractGrpcRouteRulesActionFields(r, vAction); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAction) {
		o.Action = vAction
	}
	return nil
}
func extractGrpcRouteRulesMatchesFields(r *GrpcRoute, o *GrpcRouteRulesMatches) error {
	vMethod := o.Method
	if vMethod == nil {
		// note: explicitly not the empty object.
		vMethod = &GrpcRouteRulesMatchesMethod{}
	}
	if err := extractGrpcRouteRulesMatchesMethodFields(r, vMethod); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMethod) {
		o.Method = vMethod
	}
	return nil
}
func extractGrpcRouteRulesMatchesMethodFields(r *GrpcRoute, o *GrpcRouteRulesMatchesMethod) error {
	return nil
}
func extractGrpcRouteRulesMatchesHeadersFields(r *GrpcRoute, o *GrpcRouteRulesMatchesHeaders) error {
	return nil
}
func extractGrpcRouteRulesActionFields(r *GrpcRoute, o *GrpcRouteRulesAction) error {
	vFaultInjectionPolicy := o.FaultInjectionPolicy
	if vFaultInjectionPolicy == nil {
		// note: explicitly not the empty object.
		vFaultInjectionPolicy = &GrpcRouteRulesActionFaultInjectionPolicy{}
	}
	if err := extractGrpcRouteRulesActionFaultInjectionPolicyFields(r, vFaultInjectionPolicy); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vFaultInjectionPolicy) {
		o.FaultInjectionPolicy = vFaultInjectionPolicy
	}
	vRetryPolicy := o.RetryPolicy
	if vRetryPolicy == nil {
		// note: explicitly not the empty object.
		vRetryPolicy = &GrpcRouteRulesActionRetryPolicy{}
	}
	if err := extractGrpcRouteRulesActionRetryPolicyFields(r, vRetryPolicy); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRetryPolicy) {
		o.RetryPolicy = vRetryPolicy
	}
	return nil
}
func extractGrpcRouteRulesActionDestinationsFields(r *GrpcRoute, o *GrpcRouteRulesActionDestinations) error {
	return nil
}
func extractGrpcRouteRulesActionFaultInjectionPolicyFields(r *GrpcRoute, o *GrpcRouteRulesActionFaultInjectionPolicy) error {
	vDelay := o.Delay
	if vDelay == nil {
		// note: explicitly not the empty object.
		vDelay = &GrpcRouteRulesActionFaultInjectionPolicyDelay{}
	}
	if err := extractGrpcRouteRulesActionFaultInjectionPolicyDelayFields(r, vDelay); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDelay) {
		o.Delay = vDelay
	}
	vAbort := o.Abort
	if vAbort == nil {
		// note: explicitly not the empty object.
		vAbort = &GrpcRouteRulesActionFaultInjectionPolicyAbort{}
	}
	if err := extractGrpcRouteRulesActionFaultInjectionPolicyAbortFields(r, vAbort); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAbort) {
		o.Abort = vAbort
	}
	return nil
}
func extractGrpcRouteRulesActionFaultInjectionPolicyDelayFields(r *GrpcRoute, o *GrpcRouteRulesActionFaultInjectionPolicyDelay) error {
	return nil
}
func extractGrpcRouteRulesActionFaultInjectionPolicyAbortFields(r *GrpcRoute, o *GrpcRouteRulesActionFaultInjectionPolicyAbort) error {
	return nil
}
func extractGrpcRouteRulesActionRetryPolicyFields(r *GrpcRoute, o *GrpcRouteRulesActionRetryPolicy) error {
	return nil
}

func postReadExtractGrpcRouteFields(r *GrpcRoute) error {
	return nil
}
func postReadExtractGrpcRouteRulesFields(r *GrpcRoute, o *GrpcRouteRules) error {
	vAction := o.Action
	if vAction == nil {
		// note: explicitly not the empty object.
		vAction = &GrpcRouteRulesAction{}
	}
	if err := extractGrpcRouteRulesActionFields(r, vAction); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAction) {
		o.Action = vAction
	}
	return nil
}
func postReadExtractGrpcRouteRulesMatchesFields(r *GrpcRoute, o *GrpcRouteRulesMatches) error {
	vMethod := o.Method
	if vMethod == nil {
		// note: explicitly not the empty object.
		vMethod = &GrpcRouteRulesMatchesMethod{}
	}
	if err := extractGrpcRouteRulesMatchesMethodFields(r, vMethod); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMethod) {
		o.Method = vMethod
	}
	return nil
}
func postReadExtractGrpcRouteRulesMatchesMethodFields(r *GrpcRoute, o *GrpcRouteRulesMatchesMethod) error {
	return nil
}
func postReadExtractGrpcRouteRulesMatchesHeadersFields(r *GrpcRoute, o *GrpcRouteRulesMatchesHeaders) error {
	return nil
}
func postReadExtractGrpcRouteRulesActionFields(r *GrpcRoute, o *GrpcRouteRulesAction) error {
	vFaultInjectionPolicy := o.FaultInjectionPolicy
	if vFaultInjectionPolicy == nil {
		// note: explicitly not the empty object.
		vFaultInjectionPolicy = &GrpcRouteRulesActionFaultInjectionPolicy{}
	}
	if err := extractGrpcRouteRulesActionFaultInjectionPolicyFields(r, vFaultInjectionPolicy); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vFaultInjectionPolicy) {
		o.FaultInjectionPolicy = vFaultInjectionPolicy
	}
	vRetryPolicy := o.RetryPolicy
	if vRetryPolicy == nil {
		// note: explicitly not the empty object.
		vRetryPolicy = &GrpcRouteRulesActionRetryPolicy{}
	}
	if err := extractGrpcRouteRulesActionRetryPolicyFields(r, vRetryPolicy); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRetryPolicy) {
		o.RetryPolicy = vRetryPolicy
	}
	return nil
}
func postReadExtractGrpcRouteRulesActionDestinationsFields(r *GrpcRoute, o *GrpcRouteRulesActionDestinations) error {
	return nil
}
func postReadExtractGrpcRouteRulesActionFaultInjectionPolicyFields(r *GrpcRoute, o *GrpcRouteRulesActionFaultInjectionPolicy) error {
	vDelay := o.Delay
	if vDelay == nil {
		// note: explicitly not the empty object.
		vDelay = &GrpcRouteRulesActionFaultInjectionPolicyDelay{}
	}
	if err := extractGrpcRouteRulesActionFaultInjectionPolicyDelayFields(r, vDelay); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDelay) {
		o.Delay = vDelay
	}
	vAbort := o.Abort
	if vAbort == nil {
		// note: explicitly not the empty object.
		vAbort = &GrpcRouteRulesActionFaultInjectionPolicyAbort{}
	}
	if err := extractGrpcRouteRulesActionFaultInjectionPolicyAbortFields(r, vAbort); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAbort) {
		o.Abort = vAbort
	}
	return nil
}
func postReadExtractGrpcRouteRulesActionFaultInjectionPolicyDelayFields(r *GrpcRoute, o *GrpcRouteRulesActionFaultInjectionPolicyDelay) error {
	return nil
}
func postReadExtractGrpcRouteRulesActionFaultInjectionPolicyAbortFields(r *GrpcRoute, o *GrpcRouteRulesActionFaultInjectionPolicyAbort) error {
	return nil
}
func postReadExtractGrpcRouteRulesActionRetryPolicyFields(r *GrpcRoute, o *GrpcRouteRulesActionRetryPolicy) error {
	return nil
}
