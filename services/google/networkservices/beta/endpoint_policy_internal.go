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

func (r *EndpointPolicy) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "type"); err != nil {
		return err
	}
	if err := dcl.Required(r, "endpointMatcher"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.EndpointMatcher) {
		if err := r.EndpointMatcher.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.TrafficPortSelector) {
		if err := r.TrafficPortSelector.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *EndpointPolicyEndpointMatcher) validate() error {
	if !dcl.IsEmptyValueIndirect(r.MetadataLabelMatcher) {
		if err := r.MetadataLabelMatcher.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *EndpointPolicyEndpointMatcherMetadataLabelMatcher) validate() error {
	return nil
}
func (r *EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels) validate() error {
	if err := dcl.Required(r, "labelName"); err != nil {
		return err
	}
	if err := dcl.Required(r, "labelValue"); err != nil {
		return err
	}
	return nil
}
func (r *EndpointPolicyTrafficPortSelector) validate() error {
	return nil
}
func (r *EndpointPolicy) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://networkservices.googleapis.com/v1beta1/", params)
}

func (r *EndpointPolicy) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/endpointPolicies/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *EndpointPolicy) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/endpointPolicies", nr.basePath(), userBasePath, params), nil

}

func (r *EndpointPolicy) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/endpointPolicies?endpointPolicyId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *EndpointPolicy) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/endpointPolicies/{{name}}", nr.basePath(), userBasePath, params), nil
}

// endpointPolicyApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type endpointPolicyApiOperation interface {
	do(context.Context, *EndpointPolicy, *Client) error
}

// newUpdateEndpointPolicyUpdateEndpointPolicyRequest creates a request for an
// EndpointPolicy resource's UpdateEndpointPolicy update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateEndpointPolicyUpdateEndpointPolicyRequest(ctx context.Context, f *EndpointPolicy, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		req["type"] = v
	}
	if v := f.AuthorizationPolicy; !dcl.IsEmptyValueIndirect(v) {
		req["authorizationPolicy"] = v
	}
	if v, err := expandEndpointPolicyEndpointMatcher(c, f.EndpointMatcher, res); err != nil {
		return nil, fmt.Errorf("error expanding EndpointMatcher into endpointMatcher: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["endpointMatcher"] = v
	}
	if v, err := expandEndpointPolicyTrafficPortSelector(c, f.TrafficPortSelector, res); err != nil {
		return nil, fmt.Errorf("error expanding TrafficPortSelector into trafficPortSelector: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["trafficPortSelector"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.ServerTlsPolicy; !dcl.IsEmptyValueIndirect(v) {
		req["serverTlsPolicy"] = v
	}
	if v := f.ClientTlsPolicy; !dcl.IsEmptyValueIndirect(v) {
		req["clientTlsPolicy"] = v
	}
	return req, nil
}

// marshalUpdateEndpointPolicyUpdateEndpointPolicyRequest converts the update into
// the final JSON request body.
func marshalUpdateEndpointPolicyUpdateEndpointPolicyRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateEndpointPolicyUpdateEndpointPolicyOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateEndpointPolicyUpdateEndpointPolicyOperation) do(ctx context.Context, r *EndpointPolicy, c *Client) error {
	_, err := c.GetEndpointPolicy(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateEndpointPolicy")
	if err != nil {
		return err
	}

	req, err := newUpdateEndpointPolicyUpdateEndpointPolicyRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateEndpointPolicyUpdateEndpointPolicyRequest(c, req)
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

func (c *Client) listEndpointPolicyRaw(ctx context.Context, r *EndpointPolicy, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != EndpointPolicyMaxPage {
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

type listEndpointPolicyOperation struct {
	EndpointPolicies []map[string]interface{} `json:"endpointPolicies"`
	Token            string                   `json:"nextPageToken"`
}

func (c *Client) listEndpointPolicy(ctx context.Context, r *EndpointPolicy, pageToken string, pageSize int32) ([]*EndpointPolicy, string, error) {
	b, err := c.listEndpointPolicyRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listEndpointPolicyOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*EndpointPolicy
	for _, v := range m.EndpointPolicies {
		res, err := unmarshalMapEndpointPolicy(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllEndpointPolicy(ctx context.Context, f func(*EndpointPolicy) bool, resources []*EndpointPolicy) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteEndpointPolicy(ctx, res)
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

type deleteEndpointPolicyOperation struct{}

func (op *deleteEndpointPolicyOperation) do(ctx context.Context, r *EndpointPolicy, c *Client) error {
	r, err := c.GetEndpointPolicy(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "EndpointPolicy not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetEndpointPolicy checking for existence. error: %v", err)
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
		_, err := c.GetEndpointPolicy(ctx, r)
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
type createEndpointPolicyOperation struct {
	response map[string]interface{}
}

func (op *createEndpointPolicyOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createEndpointPolicyOperation) do(ctx context.Context, r *EndpointPolicy, c *Client) error {
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

	if _, err := c.GetEndpointPolicy(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getEndpointPolicyRaw(ctx context.Context, r *EndpointPolicy) ([]byte, error) {

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

func (c *Client) endpointPolicyDiffsForRawDesired(ctx context.Context, rawDesired *EndpointPolicy, opts ...dcl.ApplyOption) (initial, desired *EndpointPolicy, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *EndpointPolicy
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*EndpointPolicy); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected EndpointPolicy, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetEndpointPolicy(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a EndpointPolicy resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve EndpointPolicy resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that EndpointPolicy resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeEndpointPolicyDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for EndpointPolicy: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for EndpointPolicy: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractEndpointPolicyFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeEndpointPolicyInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for EndpointPolicy: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeEndpointPolicyDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for EndpointPolicy: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffEndpointPolicy(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeEndpointPolicyInitialState(rawInitial, rawDesired *EndpointPolicy) (*EndpointPolicy, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeEndpointPolicyDesiredState(rawDesired, rawInitial *EndpointPolicy, opts ...dcl.ApplyOption) (*EndpointPolicy, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.EndpointMatcher = canonicalizeEndpointPolicyEndpointMatcher(rawDesired.EndpointMatcher, nil, opts...)
		rawDesired.TrafficPortSelector = canonicalizeEndpointPolicyTrafficPortSelector(rawDesired.TrafficPortSelector, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &EndpointPolicy{}
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
	if dcl.IsZeroValue(rawDesired.Type) || (dcl.IsEmptyValueIndirect(rawDesired.Type) && dcl.IsEmptyValueIndirect(rawInitial.Type)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Type = rawInitial.Type
	} else {
		canonicalDesired.Type = rawDesired.Type
	}
	if dcl.IsZeroValue(rawDesired.AuthorizationPolicy) || (dcl.IsEmptyValueIndirect(rawDesired.AuthorizationPolicy) && dcl.IsEmptyValueIndirect(rawInitial.AuthorizationPolicy)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.AuthorizationPolicy = rawInitial.AuthorizationPolicy
	} else {
		canonicalDesired.AuthorizationPolicy = rawDesired.AuthorizationPolicy
	}
	canonicalDesired.EndpointMatcher = canonicalizeEndpointPolicyEndpointMatcher(rawDesired.EndpointMatcher, rawInitial.EndpointMatcher, opts...)
	canonicalDesired.TrafficPortSelector = canonicalizeEndpointPolicyTrafficPortSelector(rawDesired.TrafficPortSelector, rawInitial.TrafficPortSelector, opts...)
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.IsZeroValue(rawDesired.ServerTlsPolicy) || (dcl.IsEmptyValueIndirect(rawDesired.ServerTlsPolicy) && dcl.IsEmptyValueIndirect(rawInitial.ServerTlsPolicy)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.ServerTlsPolicy = rawInitial.ServerTlsPolicy
	} else {
		canonicalDesired.ServerTlsPolicy = rawDesired.ServerTlsPolicy
	}
	if dcl.IsZeroValue(rawDesired.ClientTlsPolicy) || (dcl.IsEmptyValueIndirect(rawDesired.ClientTlsPolicy) && dcl.IsEmptyValueIndirect(rawInitial.ClientTlsPolicy)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.ClientTlsPolicy = rawInitial.ClientTlsPolicy
	} else {
		canonicalDesired.ClientTlsPolicy = rawDesired.ClientTlsPolicy
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

func canonicalizeEndpointPolicyNewState(c *Client, rawNew, rawDesired *EndpointPolicy) (*EndpointPolicy, error) {

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

	if dcl.IsEmptyValueIndirect(rawNew.Type) && dcl.IsEmptyValueIndirect(rawDesired.Type) {
		rawNew.Type = rawDesired.Type
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.AuthorizationPolicy) && dcl.IsEmptyValueIndirect(rawDesired.AuthorizationPolicy) {
		rawNew.AuthorizationPolicy = rawDesired.AuthorizationPolicy
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.EndpointMatcher) && dcl.IsEmptyValueIndirect(rawDesired.EndpointMatcher) {
		rawNew.EndpointMatcher = rawDesired.EndpointMatcher
	} else {
		rawNew.EndpointMatcher = canonicalizeNewEndpointPolicyEndpointMatcher(c, rawDesired.EndpointMatcher, rawNew.EndpointMatcher)
	}

	if dcl.IsEmptyValueIndirect(rawNew.TrafficPortSelector) && dcl.IsEmptyValueIndirect(rawDesired.TrafficPortSelector) {
		rawNew.TrafficPortSelector = rawDesired.TrafficPortSelector
	} else {
		rawNew.TrafficPortSelector = canonicalizeNewEndpointPolicyTrafficPortSelector(c, rawDesired.TrafficPortSelector, rawNew.TrafficPortSelector)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ServerTlsPolicy) && dcl.IsEmptyValueIndirect(rawDesired.ServerTlsPolicy) {
		rawNew.ServerTlsPolicy = rawDesired.ServerTlsPolicy
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ClientTlsPolicy) && dcl.IsEmptyValueIndirect(rawDesired.ClientTlsPolicy) {
		rawNew.ClientTlsPolicy = rawDesired.ClientTlsPolicy
	} else {
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeEndpointPolicyEndpointMatcher(des, initial *EndpointPolicyEndpointMatcher, opts ...dcl.ApplyOption) *EndpointPolicyEndpointMatcher {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EndpointPolicyEndpointMatcher{}

	cDes.MetadataLabelMatcher = canonicalizeEndpointPolicyEndpointMatcherMetadataLabelMatcher(des.MetadataLabelMatcher, initial.MetadataLabelMatcher, opts...)

	return cDes
}

func canonicalizeEndpointPolicyEndpointMatcherSlice(des, initial []EndpointPolicyEndpointMatcher, opts ...dcl.ApplyOption) []EndpointPolicyEndpointMatcher {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]EndpointPolicyEndpointMatcher, 0, len(des))
		for _, d := range des {
			cd := canonicalizeEndpointPolicyEndpointMatcher(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]EndpointPolicyEndpointMatcher, 0, len(des))
	for i, d := range des {
		cd := canonicalizeEndpointPolicyEndpointMatcher(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewEndpointPolicyEndpointMatcher(c *Client, des, nw *EndpointPolicyEndpointMatcher) *EndpointPolicyEndpointMatcher {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for EndpointPolicyEndpointMatcher while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.MetadataLabelMatcher = canonicalizeNewEndpointPolicyEndpointMatcherMetadataLabelMatcher(c, des.MetadataLabelMatcher, nw.MetadataLabelMatcher)

	return nw
}

func canonicalizeNewEndpointPolicyEndpointMatcherSet(c *Client, des, nw []EndpointPolicyEndpointMatcher) []EndpointPolicyEndpointMatcher {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []EndpointPolicyEndpointMatcher
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareEndpointPolicyEndpointMatcherNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewEndpointPolicyEndpointMatcher(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewEndpointPolicyEndpointMatcherSlice(c *Client, des, nw []EndpointPolicyEndpointMatcher) []EndpointPolicyEndpointMatcher {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EndpointPolicyEndpointMatcher
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEndpointPolicyEndpointMatcher(c, &d, &n))
	}

	return items
}

func canonicalizeEndpointPolicyEndpointMatcherMetadataLabelMatcher(des, initial *EndpointPolicyEndpointMatcherMetadataLabelMatcher, opts ...dcl.ApplyOption) *EndpointPolicyEndpointMatcherMetadataLabelMatcher {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EndpointPolicyEndpointMatcherMetadataLabelMatcher{}

	if dcl.IsZeroValue(des.MetadataLabelMatchCriteria) || (dcl.IsEmptyValueIndirect(des.MetadataLabelMatchCriteria) && dcl.IsEmptyValueIndirect(initial.MetadataLabelMatchCriteria)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MetadataLabelMatchCriteria = initial.MetadataLabelMatchCriteria
	} else {
		cDes.MetadataLabelMatchCriteria = des.MetadataLabelMatchCriteria
	}
	cDes.MetadataLabels = canonicalizeEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsSlice(des.MetadataLabels, initial.MetadataLabels, opts...)

	return cDes
}

func canonicalizeEndpointPolicyEndpointMatcherMetadataLabelMatcherSlice(des, initial []EndpointPolicyEndpointMatcherMetadataLabelMatcher, opts ...dcl.ApplyOption) []EndpointPolicyEndpointMatcherMetadataLabelMatcher {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]EndpointPolicyEndpointMatcherMetadataLabelMatcher, 0, len(des))
		for _, d := range des {
			cd := canonicalizeEndpointPolicyEndpointMatcherMetadataLabelMatcher(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]EndpointPolicyEndpointMatcherMetadataLabelMatcher, 0, len(des))
	for i, d := range des {
		cd := canonicalizeEndpointPolicyEndpointMatcherMetadataLabelMatcher(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewEndpointPolicyEndpointMatcherMetadataLabelMatcher(c *Client, des, nw *EndpointPolicyEndpointMatcherMetadataLabelMatcher) *EndpointPolicyEndpointMatcherMetadataLabelMatcher {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for EndpointPolicyEndpointMatcherMetadataLabelMatcher while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.MetadataLabels = canonicalizeNewEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsSlice(c, des.MetadataLabels, nw.MetadataLabels)

	return nw
}

func canonicalizeNewEndpointPolicyEndpointMatcherMetadataLabelMatcherSet(c *Client, des, nw []EndpointPolicyEndpointMatcherMetadataLabelMatcher) []EndpointPolicyEndpointMatcherMetadataLabelMatcher {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []EndpointPolicyEndpointMatcherMetadataLabelMatcher
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareEndpointPolicyEndpointMatcherMetadataLabelMatcherNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewEndpointPolicyEndpointMatcherMetadataLabelMatcher(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewEndpointPolicyEndpointMatcherMetadataLabelMatcherSlice(c *Client, des, nw []EndpointPolicyEndpointMatcherMetadataLabelMatcher) []EndpointPolicyEndpointMatcherMetadataLabelMatcher {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EndpointPolicyEndpointMatcherMetadataLabelMatcher
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEndpointPolicyEndpointMatcherMetadataLabelMatcher(c, &d, &n))
	}

	return items
}

func canonicalizeEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels(des, initial *EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels, opts ...dcl.ApplyOption) *EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels{}

	if dcl.StringCanonicalize(des.LabelName, initial.LabelName) || dcl.IsZeroValue(des.LabelName) {
		cDes.LabelName = initial.LabelName
	} else {
		cDes.LabelName = des.LabelName
	}
	if dcl.StringCanonicalize(des.LabelValue, initial.LabelValue) || dcl.IsZeroValue(des.LabelValue) {
		cDes.LabelValue = initial.LabelValue
	} else {
		cDes.LabelValue = des.LabelValue
	}

	return cDes
}

func canonicalizeEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsSlice(des, initial []EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels, opts ...dcl.ApplyOption) []EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels, 0, len(des))
		for _, d := range des {
			cd := canonicalizeEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels, 0, len(des))
	for i, d := range des {
		cd := canonicalizeEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels(c *Client, des, nw *EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels) *EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.LabelName, nw.LabelName) {
		nw.LabelName = des.LabelName
	}
	if dcl.StringCanonicalize(des.LabelValue, nw.LabelValue) {
		nw.LabelValue = des.LabelValue
	}

	return nw
}

func canonicalizeNewEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsSet(c *Client, des, nw []EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels) []EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsSlice(c *Client, des, nw []EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels) []EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels(c, &d, &n))
	}

	return items
}

func canonicalizeEndpointPolicyTrafficPortSelector(des, initial *EndpointPolicyTrafficPortSelector, opts ...dcl.ApplyOption) *EndpointPolicyTrafficPortSelector {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EndpointPolicyTrafficPortSelector{}

	if dcl.StringArrayCanonicalize(des.Ports, initial.Ports) {
		cDes.Ports = initial.Ports
	} else {
		cDes.Ports = des.Ports
	}

	return cDes
}

func canonicalizeEndpointPolicyTrafficPortSelectorSlice(des, initial []EndpointPolicyTrafficPortSelector, opts ...dcl.ApplyOption) []EndpointPolicyTrafficPortSelector {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]EndpointPolicyTrafficPortSelector, 0, len(des))
		for _, d := range des {
			cd := canonicalizeEndpointPolicyTrafficPortSelector(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]EndpointPolicyTrafficPortSelector, 0, len(des))
	for i, d := range des {
		cd := canonicalizeEndpointPolicyTrafficPortSelector(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewEndpointPolicyTrafficPortSelector(c *Client, des, nw *EndpointPolicyTrafficPortSelector) *EndpointPolicyTrafficPortSelector {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for EndpointPolicyTrafficPortSelector while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.Ports, nw.Ports) {
		nw.Ports = des.Ports
	}

	return nw
}

func canonicalizeNewEndpointPolicyTrafficPortSelectorSet(c *Client, des, nw []EndpointPolicyTrafficPortSelector) []EndpointPolicyTrafficPortSelector {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []EndpointPolicyTrafficPortSelector
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareEndpointPolicyTrafficPortSelectorNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewEndpointPolicyTrafficPortSelector(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewEndpointPolicyTrafficPortSelectorSlice(c *Client, des, nw []EndpointPolicyTrafficPortSelector) []EndpointPolicyTrafficPortSelector {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EndpointPolicyTrafficPortSelector
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEndpointPolicyTrafficPortSelector(c, &d, &n))
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
func diffEndpointPolicy(c *Client, desired, actual *EndpointPolicy, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateEndpointPolicyUpdateEndpointPolicyOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateEndpointPolicyUpdateEndpointPolicyOperation")}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AuthorizationPolicy, actual.AuthorizationPolicy, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateEndpointPolicyUpdateEndpointPolicyOperation")}, fn.AddNest("AuthorizationPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EndpointMatcher, actual.EndpointMatcher, dcl.DiffInfo{ObjectFunction: compareEndpointPolicyEndpointMatcherNewStyle, EmptyObject: EmptyEndpointPolicyEndpointMatcher, OperationSelector: dcl.TriggersOperation("updateEndpointPolicyUpdateEndpointPolicyOperation")}, fn.AddNest("EndpointMatcher")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TrafficPortSelector, actual.TrafficPortSelector, dcl.DiffInfo{ObjectFunction: compareEndpointPolicyTrafficPortSelectorNewStyle, EmptyObject: EmptyEndpointPolicyTrafficPortSelector, OperationSelector: dcl.TriggersOperation("updateEndpointPolicyUpdateEndpointPolicyOperation")}, fn.AddNest("TrafficPortSelector")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateEndpointPolicyUpdateEndpointPolicyOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServerTlsPolicy, actual.ServerTlsPolicy, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateEndpointPolicyUpdateEndpointPolicyOperation")}, fn.AddNest("ServerTlsPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClientTlsPolicy, actual.ClientTlsPolicy, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateEndpointPolicyUpdateEndpointPolicyOperation")}, fn.AddNest("ClientTlsPolicy")); len(ds) != 0 || err != nil {
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
func compareEndpointPolicyEndpointMatcherNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EndpointPolicyEndpointMatcher)
	if !ok {
		desiredNotPointer, ok := d.(EndpointPolicyEndpointMatcher)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointPolicyEndpointMatcher or *EndpointPolicyEndpointMatcher", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EndpointPolicyEndpointMatcher)
	if !ok {
		actualNotPointer, ok := a.(EndpointPolicyEndpointMatcher)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointPolicyEndpointMatcher", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MetadataLabelMatcher, actual.MetadataLabelMatcher, dcl.DiffInfo{ObjectFunction: compareEndpointPolicyEndpointMatcherMetadataLabelMatcherNewStyle, EmptyObject: EmptyEndpointPolicyEndpointMatcherMetadataLabelMatcher, OperationSelector: dcl.TriggersOperation("updateEndpointPolicyUpdateEndpointPolicyOperation")}, fn.AddNest("MetadataLabelMatcher")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareEndpointPolicyEndpointMatcherMetadataLabelMatcherNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EndpointPolicyEndpointMatcherMetadataLabelMatcher)
	if !ok {
		desiredNotPointer, ok := d.(EndpointPolicyEndpointMatcherMetadataLabelMatcher)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointPolicyEndpointMatcherMetadataLabelMatcher or *EndpointPolicyEndpointMatcherMetadataLabelMatcher", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EndpointPolicyEndpointMatcherMetadataLabelMatcher)
	if !ok {
		actualNotPointer, ok := a.(EndpointPolicyEndpointMatcherMetadataLabelMatcher)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointPolicyEndpointMatcherMetadataLabelMatcher", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MetadataLabelMatchCriteria, actual.MetadataLabelMatchCriteria, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateEndpointPolicyUpdateEndpointPolicyOperation")}, fn.AddNest("MetadataLabelMatchCriteria")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MetadataLabels, actual.MetadataLabels, dcl.DiffInfo{ObjectFunction: compareEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsNewStyle, EmptyObject: EmptyEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels, OperationSelector: dcl.TriggersOperation("updateEndpointPolicyUpdateEndpointPolicyOperation")}, fn.AddNest("MetadataLabels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels)
	if !ok {
		desiredNotPointer, ok := d.(EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels or *EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels)
	if !ok {
		actualNotPointer, ok := a.(EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.LabelName, actual.LabelName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateEndpointPolicyUpdateEndpointPolicyOperation")}, fn.AddNest("LabelName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LabelValue, actual.LabelValue, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateEndpointPolicyUpdateEndpointPolicyOperation")}, fn.AddNest("LabelValue")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareEndpointPolicyTrafficPortSelectorNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EndpointPolicyTrafficPortSelector)
	if !ok {
		desiredNotPointer, ok := d.(EndpointPolicyTrafficPortSelector)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointPolicyTrafficPortSelector or *EndpointPolicyTrafficPortSelector", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EndpointPolicyTrafficPortSelector)
	if !ok {
		actualNotPointer, ok := a.(EndpointPolicyTrafficPortSelector)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointPolicyTrafficPortSelector", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Ports, actual.Ports, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateEndpointPolicyUpdateEndpointPolicyOperation")}, fn.AddNest("Ports")); len(ds) != 0 || err != nil {
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
func (r *EndpointPolicy) urlNormalized() *EndpointPolicy {
	normalized := dcl.Copy(*r).(EndpointPolicy)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.AuthorizationPolicy = dcl.SelfLinkToName(r.AuthorizationPolicy)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.ServerTlsPolicy = dcl.SelfLinkToName(r.ServerTlsPolicy)
	normalized.ClientTlsPolicy = dcl.SelfLinkToName(r.ClientTlsPolicy)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *EndpointPolicy) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateEndpointPolicy" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/endpointPolicies/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the EndpointPolicy resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *EndpointPolicy) marshal(c *Client) ([]byte, error) {
	m, err := expandEndpointPolicy(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling EndpointPolicy: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalEndpointPolicy decodes JSON responses into the EndpointPolicy resource schema.
func unmarshalEndpointPolicy(b []byte, c *Client, res *EndpointPolicy) (*EndpointPolicy, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapEndpointPolicy(m, c, res)
}

func unmarshalMapEndpointPolicy(m map[string]interface{}, c *Client, res *EndpointPolicy) (*EndpointPolicy, error) {

	flattened := flattenEndpointPolicy(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandEndpointPolicy expands EndpointPolicy into a JSON request object.
func expandEndpointPolicy(c *Client, f *EndpointPolicy) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("projects/*/locations/global/endpointPolicies/%s", f.Name, dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Labels; dcl.ValueShouldBeSent(v) {
		m["labels"] = v
	}
	if v := f.Type; dcl.ValueShouldBeSent(v) {
		m["type"] = v
	}
	if v := f.AuthorizationPolicy; dcl.ValueShouldBeSent(v) {
		m["authorizationPolicy"] = v
	}
	if v, err := expandEndpointPolicyEndpointMatcher(c, f.EndpointMatcher, res); err != nil {
		return nil, fmt.Errorf("error expanding EndpointMatcher into endpointMatcher: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["endpointMatcher"] = v
	}
	if v, err := expandEndpointPolicyTrafficPortSelector(c, f.TrafficPortSelector, res); err != nil {
		return nil, fmt.Errorf("error expanding TrafficPortSelector into trafficPortSelector: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["trafficPortSelector"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v := f.ServerTlsPolicy; dcl.ValueShouldBeSent(v) {
		m["serverTlsPolicy"] = v
	}
	if v := f.ClientTlsPolicy; dcl.ValueShouldBeSent(v) {
		m["clientTlsPolicy"] = v
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

// flattenEndpointPolicy flattens EndpointPolicy from a JSON request object into the
// EndpointPolicy type.
func flattenEndpointPolicy(c *Client, i interface{}, res *EndpointPolicy) *EndpointPolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &EndpointPolicy{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	resultRes.Type = flattenEndpointPolicyTypeEnum(m["type"])
	resultRes.AuthorizationPolicy = dcl.FlattenString(m["authorizationPolicy"])
	resultRes.EndpointMatcher = flattenEndpointPolicyEndpointMatcher(c, m["endpointMatcher"], res)
	resultRes.TrafficPortSelector = flattenEndpointPolicyTrafficPortSelector(c, m["trafficPortSelector"], res)
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.ServerTlsPolicy = dcl.FlattenString(m["serverTlsPolicy"])
	resultRes.ClientTlsPolicy = dcl.FlattenString(m["clientTlsPolicy"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// expandEndpointPolicyEndpointMatcherMap expands the contents of EndpointPolicyEndpointMatcher into a JSON
// request object.
func expandEndpointPolicyEndpointMatcherMap(c *Client, f map[string]EndpointPolicyEndpointMatcher, res *EndpointPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEndpointPolicyEndpointMatcher(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEndpointPolicyEndpointMatcherSlice expands the contents of EndpointPolicyEndpointMatcher into a JSON
// request object.
func expandEndpointPolicyEndpointMatcherSlice(c *Client, f []EndpointPolicyEndpointMatcher, res *EndpointPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEndpointPolicyEndpointMatcher(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEndpointPolicyEndpointMatcherMap flattens the contents of EndpointPolicyEndpointMatcher from a JSON
// response object.
func flattenEndpointPolicyEndpointMatcherMap(c *Client, i interface{}, res *EndpointPolicy) map[string]EndpointPolicyEndpointMatcher {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EndpointPolicyEndpointMatcher{}
	}

	if len(a) == 0 {
		return map[string]EndpointPolicyEndpointMatcher{}
	}

	items := make(map[string]EndpointPolicyEndpointMatcher)
	for k, item := range a {
		items[k] = *flattenEndpointPolicyEndpointMatcher(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenEndpointPolicyEndpointMatcherSlice flattens the contents of EndpointPolicyEndpointMatcher from a JSON
// response object.
func flattenEndpointPolicyEndpointMatcherSlice(c *Client, i interface{}, res *EndpointPolicy) []EndpointPolicyEndpointMatcher {
	a, ok := i.([]interface{})
	if !ok {
		return []EndpointPolicyEndpointMatcher{}
	}

	if len(a) == 0 {
		return []EndpointPolicyEndpointMatcher{}
	}

	items := make([]EndpointPolicyEndpointMatcher, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEndpointPolicyEndpointMatcher(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandEndpointPolicyEndpointMatcher expands an instance of EndpointPolicyEndpointMatcher into a JSON
// request object.
func expandEndpointPolicyEndpointMatcher(c *Client, f *EndpointPolicyEndpointMatcher, res *EndpointPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandEndpointPolicyEndpointMatcherMetadataLabelMatcher(c, f.MetadataLabelMatcher, res); err != nil {
		return nil, fmt.Errorf("error expanding MetadataLabelMatcher into metadataLabelMatcher: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["metadataLabelMatcher"] = v
	}

	return m, nil
}

// flattenEndpointPolicyEndpointMatcher flattens an instance of EndpointPolicyEndpointMatcher from a JSON
// response object.
func flattenEndpointPolicyEndpointMatcher(c *Client, i interface{}, res *EndpointPolicy) *EndpointPolicyEndpointMatcher {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EndpointPolicyEndpointMatcher{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEndpointPolicyEndpointMatcher
	}
	r.MetadataLabelMatcher = flattenEndpointPolicyEndpointMatcherMetadataLabelMatcher(c, m["metadataLabelMatcher"], res)

	return r
}

// expandEndpointPolicyEndpointMatcherMetadataLabelMatcherMap expands the contents of EndpointPolicyEndpointMatcherMetadataLabelMatcher into a JSON
// request object.
func expandEndpointPolicyEndpointMatcherMetadataLabelMatcherMap(c *Client, f map[string]EndpointPolicyEndpointMatcherMetadataLabelMatcher, res *EndpointPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEndpointPolicyEndpointMatcherMetadataLabelMatcher(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEndpointPolicyEndpointMatcherMetadataLabelMatcherSlice expands the contents of EndpointPolicyEndpointMatcherMetadataLabelMatcher into a JSON
// request object.
func expandEndpointPolicyEndpointMatcherMetadataLabelMatcherSlice(c *Client, f []EndpointPolicyEndpointMatcherMetadataLabelMatcher, res *EndpointPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEndpointPolicyEndpointMatcherMetadataLabelMatcher(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEndpointPolicyEndpointMatcherMetadataLabelMatcherMap flattens the contents of EndpointPolicyEndpointMatcherMetadataLabelMatcher from a JSON
// response object.
func flattenEndpointPolicyEndpointMatcherMetadataLabelMatcherMap(c *Client, i interface{}, res *EndpointPolicy) map[string]EndpointPolicyEndpointMatcherMetadataLabelMatcher {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EndpointPolicyEndpointMatcherMetadataLabelMatcher{}
	}

	if len(a) == 0 {
		return map[string]EndpointPolicyEndpointMatcherMetadataLabelMatcher{}
	}

	items := make(map[string]EndpointPolicyEndpointMatcherMetadataLabelMatcher)
	for k, item := range a {
		items[k] = *flattenEndpointPolicyEndpointMatcherMetadataLabelMatcher(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenEndpointPolicyEndpointMatcherMetadataLabelMatcherSlice flattens the contents of EndpointPolicyEndpointMatcherMetadataLabelMatcher from a JSON
// response object.
func flattenEndpointPolicyEndpointMatcherMetadataLabelMatcherSlice(c *Client, i interface{}, res *EndpointPolicy) []EndpointPolicyEndpointMatcherMetadataLabelMatcher {
	a, ok := i.([]interface{})
	if !ok {
		return []EndpointPolicyEndpointMatcherMetadataLabelMatcher{}
	}

	if len(a) == 0 {
		return []EndpointPolicyEndpointMatcherMetadataLabelMatcher{}
	}

	items := make([]EndpointPolicyEndpointMatcherMetadataLabelMatcher, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEndpointPolicyEndpointMatcherMetadataLabelMatcher(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandEndpointPolicyEndpointMatcherMetadataLabelMatcher expands an instance of EndpointPolicyEndpointMatcherMetadataLabelMatcher into a JSON
// request object.
func expandEndpointPolicyEndpointMatcherMetadataLabelMatcher(c *Client, f *EndpointPolicyEndpointMatcherMetadataLabelMatcher, res *EndpointPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MetadataLabelMatchCriteria; !dcl.IsEmptyValueIndirect(v) {
		m["metadataLabelMatchCriteria"] = v
	}
	if v, err := expandEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsSlice(c, f.MetadataLabels, res); err != nil {
		return nil, fmt.Errorf("error expanding MetadataLabels into metadataLabels: %w", err)
	} else if v != nil {
		m["metadataLabels"] = v
	}

	return m, nil
}

// flattenEndpointPolicyEndpointMatcherMetadataLabelMatcher flattens an instance of EndpointPolicyEndpointMatcherMetadataLabelMatcher from a JSON
// response object.
func flattenEndpointPolicyEndpointMatcherMetadataLabelMatcher(c *Client, i interface{}, res *EndpointPolicy) *EndpointPolicyEndpointMatcherMetadataLabelMatcher {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EndpointPolicyEndpointMatcherMetadataLabelMatcher{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEndpointPolicyEndpointMatcherMetadataLabelMatcher
	}
	r.MetadataLabelMatchCriteria = flattenEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(m["metadataLabelMatchCriteria"])
	r.MetadataLabels = flattenEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsSlice(c, m["metadataLabels"], res)

	return r
}

// expandEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsMap expands the contents of EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels into a JSON
// request object.
func expandEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsMap(c *Client, f map[string]EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels, res *EndpointPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsSlice expands the contents of EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels into a JSON
// request object.
func expandEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsSlice(c *Client, f []EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels, res *EndpointPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsMap flattens the contents of EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels from a JSON
// response object.
func flattenEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsMap(c *Client, i interface{}, res *EndpointPolicy) map[string]EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels{}
	}

	if len(a) == 0 {
		return map[string]EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels{}
	}

	items := make(map[string]EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels)
	for k, item := range a {
		items[k] = *flattenEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsSlice flattens the contents of EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels from a JSON
// response object.
func flattenEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsSlice(c *Client, i interface{}, res *EndpointPolicy) []EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels {
	a, ok := i.([]interface{})
	if !ok {
		return []EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels{}
	}

	if len(a) == 0 {
		return []EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels{}
	}

	items := make([]EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels expands an instance of EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels into a JSON
// request object.
func expandEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels(c *Client, f *EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels, res *EndpointPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.LabelName; !dcl.IsEmptyValueIndirect(v) {
		m["labelName"] = v
	}
	if v := f.LabelValue; !dcl.IsEmptyValueIndirect(v) {
		m["labelValue"] = v
	}

	return m, nil
}

// flattenEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels flattens an instance of EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels from a JSON
// response object.
func flattenEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels(c *Client, i interface{}, res *EndpointPolicy) *EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels
	}
	r.LabelName = dcl.FlattenString(m["labelName"])
	r.LabelValue = dcl.FlattenString(m["labelValue"])

	return r
}

// expandEndpointPolicyTrafficPortSelectorMap expands the contents of EndpointPolicyTrafficPortSelector into a JSON
// request object.
func expandEndpointPolicyTrafficPortSelectorMap(c *Client, f map[string]EndpointPolicyTrafficPortSelector, res *EndpointPolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEndpointPolicyTrafficPortSelector(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEndpointPolicyTrafficPortSelectorSlice expands the contents of EndpointPolicyTrafficPortSelector into a JSON
// request object.
func expandEndpointPolicyTrafficPortSelectorSlice(c *Client, f []EndpointPolicyTrafficPortSelector, res *EndpointPolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEndpointPolicyTrafficPortSelector(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEndpointPolicyTrafficPortSelectorMap flattens the contents of EndpointPolicyTrafficPortSelector from a JSON
// response object.
func flattenEndpointPolicyTrafficPortSelectorMap(c *Client, i interface{}, res *EndpointPolicy) map[string]EndpointPolicyTrafficPortSelector {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EndpointPolicyTrafficPortSelector{}
	}

	if len(a) == 0 {
		return map[string]EndpointPolicyTrafficPortSelector{}
	}

	items := make(map[string]EndpointPolicyTrafficPortSelector)
	for k, item := range a {
		items[k] = *flattenEndpointPolicyTrafficPortSelector(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenEndpointPolicyTrafficPortSelectorSlice flattens the contents of EndpointPolicyTrafficPortSelector from a JSON
// response object.
func flattenEndpointPolicyTrafficPortSelectorSlice(c *Client, i interface{}, res *EndpointPolicy) []EndpointPolicyTrafficPortSelector {
	a, ok := i.([]interface{})
	if !ok {
		return []EndpointPolicyTrafficPortSelector{}
	}

	if len(a) == 0 {
		return []EndpointPolicyTrafficPortSelector{}
	}

	items := make([]EndpointPolicyTrafficPortSelector, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEndpointPolicyTrafficPortSelector(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandEndpointPolicyTrafficPortSelector expands an instance of EndpointPolicyTrafficPortSelector into a JSON
// request object.
func expandEndpointPolicyTrafficPortSelector(c *Client, f *EndpointPolicyTrafficPortSelector, res *EndpointPolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Ports; v != nil {
		m["ports"] = v
	}

	return m, nil
}

// flattenEndpointPolicyTrafficPortSelector flattens an instance of EndpointPolicyTrafficPortSelector from a JSON
// response object.
func flattenEndpointPolicyTrafficPortSelector(c *Client, i interface{}, res *EndpointPolicy) *EndpointPolicyTrafficPortSelector {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EndpointPolicyTrafficPortSelector{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEndpointPolicyTrafficPortSelector
	}
	r.Ports = dcl.FlattenStringSlice(m["ports"])

	return r
}

// flattenEndpointPolicyTypeEnumMap flattens the contents of EndpointPolicyTypeEnum from a JSON
// response object.
func flattenEndpointPolicyTypeEnumMap(c *Client, i interface{}, res *EndpointPolicy) map[string]EndpointPolicyTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EndpointPolicyTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]EndpointPolicyTypeEnum{}
	}

	items := make(map[string]EndpointPolicyTypeEnum)
	for k, item := range a {
		items[k] = *flattenEndpointPolicyTypeEnum(item.(interface{}))
	}

	return items
}

// flattenEndpointPolicyTypeEnumSlice flattens the contents of EndpointPolicyTypeEnum from a JSON
// response object.
func flattenEndpointPolicyTypeEnumSlice(c *Client, i interface{}, res *EndpointPolicy) []EndpointPolicyTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []EndpointPolicyTypeEnum{}
	}

	if len(a) == 0 {
		return []EndpointPolicyTypeEnum{}
	}

	items := make([]EndpointPolicyTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEndpointPolicyTypeEnum(item.(interface{})))
	}

	return items
}

// flattenEndpointPolicyTypeEnum asserts that an interface is a string, and returns a
// pointer to a *EndpointPolicyTypeEnum with the same value as that string.
func flattenEndpointPolicyTypeEnum(i interface{}) *EndpointPolicyTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return EndpointPolicyTypeEnumRef(s)
}

// flattenEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumMap flattens the contents of EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum from a JSON
// response object.
func flattenEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumMap(c *Client, i interface{}, res *EndpointPolicy) map[string]EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum{}
	}

	if len(a) == 0 {
		return map[string]EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum{}
	}

	items := make(map[string]EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum)
	for k, item := range a {
		items[k] = *flattenEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(item.(interface{}))
	}

	return items
}

// flattenEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumSlice flattens the contents of EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum from a JSON
// response object.
func flattenEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumSlice(c *Client, i interface{}, res *EndpointPolicy) []EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum{}
	}

	if len(a) == 0 {
		return []EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum{}
	}

	items := make([]EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(item.(interface{})))
	}

	return items
}

// flattenEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum asserts that an interface is a string, and returns a
// pointer to a *EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum with the same value as that string.
func flattenEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(i interface{}) *EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *EndpointPolicy) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalEndpointPolicy(b, c, r)
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

type endpointPolicyDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         endpointPolicyApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToEndpointPolicyDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]endpointPolicyDiff, error) {
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
	var diffs []endpointPolicyDiff
	// For each operation name, create a endpointPolicyDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := endpointPolicyDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToEndpointPolicyApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToEndpointPolicyApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (endpointPolicyApiOperation, error) {
	switch opName {

	case "updateEndpointPolicyUpdateEndpointPolicyOperation":
		return &updateEndpointPolicyUpdateEndpointPolicyOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractEndpointPolicyFields(r *EndpointPolicy) error {
	vEndpointMatcher := r.EndpointMatcher
	if vEndpointMatcher == nil {
		// note: explicitly not the empty object.
		vEndpointMatcher = &EndpointPolicyEndpointMatcher{}
	}
	if err := extractEndpointPolicyEndpointMatcherFields(r, vEndpointMatcher); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vEndpointMatcher) {
		r.EndpointMatcher = vEndpointMatcher
	}
	vTrafficPortSelector := r.TrafficPortSelector
	if vTrafficPortSelector == nil {
		// note: explicitly not the empty object.
		vTrafficPortSelector = &EndpointPolicyTrafficPortSelector{}
	}
	if err := extractEndpointPolicyTrafficPortSelectorFields(r, vTrafficPortSelector); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTrafficPortSelector) {
		r.TrafficPortSelector = vTrafficPortSelector
	}
	return nil
}
func extractEndpointPolicyEndpointMatcherFields(r *EndpointPolicy, o *EndpointPolicyEndpointMatcher) error {
	vMetadataLabelMatcher := o.MetadataLabelMatcher
	if vMetadataLabelMatcher == nil {
		// note: explicitly not the empty object.
		vMetadataLabelMatcher = &EndpointPolicyEndpointMatcherMetadataLabelMatcher{}
	}
	if err := extractEndpointPolicyEndpointMatcherMetadataLabelMatcherFields(r, vMetadataLabelMatcher); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMetadataLabelMatcher) {
		o.MetadataLabelMatcher = vMetadataLabelMatcher
	}
	return nil
}
func extractEndpointPolicyEndpointMatcherMetadataLabelMatcherFields(r *EndpointPolicy, o *EndpointPolicyEndpointMatcherMetadataLabelMatcher) error {
	return nil
}
func extractEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsFields(r *EndpointPolicy, o *EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels) error {
	return nil
}
func extractEndpointPolicyTrafficPortSelectorFields(r *EndpointPolicy, o *EndpointPolicyTrafficPortSelector) error {
	return nil
}

func postReadExtractEndpointPolicyFields(r *EndpointPolicy) error {
	vEndpointMatcher := r.EndpointMatcher
	if vEndpointMatcher == nil {
		// note: explicitly not the empty object.
		vEndpointMatcher = &EndpointPolicyEndpointMatcher{}
	}
	if err := postReadExtractEndpointPolicyEndpointMatcherFields(r, vEndpointMatcher); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vEndpointMatcher) {
		r.EndpointMatcher = vEndpointMatcher
	}
	vTrafficPortSelector := r.TrafficPortSelector
	if vTrafficPortSelector == nil {
		// note: explicitly not the empty object.
		vTrafficPortSelector = &EndpointPolicyTrafficPortSelector{}
	}
	if err := postReadExtractEndpointPolicyTrafficPortSelectorFields(r, vTrafficPortSelector); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTrafficPortSelector) {
		r.TrafficPortSelector = vTrafficPortSelector
	}
	return nil
}
func postReadExtractEndpointPolicyEndpointMatcherFields(r *EndpointPolicy, o *EndpointPolicyEndpointMatcher) error {
	vMetadataLabelMatcher := o.MetadataLabelMatcher
	if vMetadataLabelMatcher == nil {
		// note: explicitly not the empty object.
		vMetadataLabelMatcher = &EndpointPolicyEndpointMatcherMetadataLabelMatcher{}
	}
	if err := extractEndpointPolicyEndpointMatcherMetadataLabelMatcherFields(r, vMetadataLabelMatcher); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMetadataLabelMatcher) {
		o.MetadataLabelMatcher = vMetadataLabelMatcher
	}
	return nil
}
func postReadExtractEndpointPolicyEndpointMatcherMetadataLabelMatcherFields(r *EndpointPolicy, o *EndpointPolicyEndpointMatcherMetadataLabelMatcher) error {
	return nil
}
func postReadExtractEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsFields(r *EndpointPolicy, o *EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels) error {
	return nil
}
func postReadExtractEndpointPolicyTrafficPortSelectorFields(r *EndpointPolicy, o *EndpointPolicyTrafficPortSelector) error {
	return nil
}
