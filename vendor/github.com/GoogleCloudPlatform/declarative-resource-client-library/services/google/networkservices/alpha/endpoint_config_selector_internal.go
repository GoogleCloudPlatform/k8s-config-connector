// Copyright 2021 Google LLC. All Rights Reserved.
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
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *EndpointConfigSelector) validate() error {

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
	if !dcl.IsEmptyValueIndirect(r.HttpFilters) {
		if err := r.HttpFilters.validate(); err != nil {
			return err
		}
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
func (r *EndpointConfigSelectorHttpFilters) validate() error {
	if err := dcl.Required(r, "httpFilters"); err != nil {
		return err
	}
	return nil
}
func (r *EndpointConfigSelectorEndpointMatcher) validate() error {
	if !dcl.IsEmptyValueIndirect(r.MetadataLabelMatcher) {
		if err := r.MetadataLabelMatcher.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher) validate() error {
	return nil
}
func (r *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels) validate() error {
	if err := dcl.Required(r, "labelName"); err != nil {
		return err
	}
	if err := dcl.Required(r, "labelValue"); err != nil {
		return err
	}
	return nil
}
func (r *EndpointConfigSelectorTrafficPortSelector) validate() error {
	return nil
}
func (r *EndpointConfigSelector) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://networkservices.googleapis.com/v1alpha1/", params)
}

func (r *EndpointConfigSelector) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/endpointConfigSelectors/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *EndpointConfigSelector) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/endpointConfigSelectors", nr.basePath(), userBasePath, params), nil

}

func (r *EndpointConfigSelector) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/endpointConfigSelectors?endpointConfigSelectorId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *EndpointConfigSelector) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/endpointConfigSelectors/{{name}}", nr.basePath(), userBasePath, params), nil
}

// endpointConfigSelectorApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type endpointConfigSelectorApiOperation interface {
	do(context.Context, *EndpointConfigSelector, *Client) error
}

// newUpdateEndpointConfigSelectorUpdateEndpointConfigSelectorRequest creates a request for an
// EndpointConfigSelector resource's UpdateEndpointConfigSelector update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateEndpointConfigSelectorUpdateEndpointConfigSelectorRequest(ctx context.Context, f *EndpointConfigSelector, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		req["type"] = v
	}
	if v := f.AuthorizationPolicy; !dcl.IsEmptyValueIndirect(v) {
		req["authorizationPolicy"] = v
	}
	if v, err := expandEndpointConfigSelectorHttpFilters(c, f.HttpFilters); err != nil {
		return nil, fmt.Errorf("error expanding HttpFilters into httpFilters: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["httpFilters"] = v
	}
	if v, err := expandEndpointConfigSelectorEndpointMatcher(c, f.EndpointMatcher); err != nil {
		return nil, fmt.Errorf("error expanding EndpointMatcher into endpointMatcher: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["endpointMatcher"] = v
	}
	if v, err := expandEndpointConfigSelectorTrafficPortSelector(c, f.TrafficPortSelector); err != nil {
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

// marshalUpdateEndpointConfigSelectorUpdateEndpointConfigSelectorRequest converts the update into
// the final JSON request body.
func marshalUpdateEndpointConfigSelectorUpdateEndpointConfigSelectorRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateEndpointConfigSelectorUpdateEndpointConfigSelectorOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateEndpointConfigSelectorUpdateEndpointConfigSelectorOperation) do(ctx context.Context, r *EndpointConfigSelector, c *Client) error {
	_, err := c.GetEndpointConfigSelector(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateEndpointConfigSelector")
	if err != nil {
		return err
	}

	req, err := newUpdateEndpointConfigSelectorUpdateEndpointConfigSelectorRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateEndpointConfigSelectorUpdateEndpointConfigSelectorRequest(c, req)
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

func (c *Client) listEndpointConfigSelectorRaw(ctx context.Context, r *EndpointConfigSelector, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != EndpointConfigSelectorMaxPage {
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

type listEndpointConfigSelectorOperation struct {
	EndpointConfigSelectors []map[string]interface{} `json:"endpointConfigSelectors"`
	Token                   string                   `json:"nextPageToken"`
}

func (c *Client) listEndpointConfigSelector(ctx context.Context, r *EndpointConfigSelector, pageToken string, pageSize int32) ([]*EndpointConfigSelector, string, error) {
	b, err := c.listEndpointConfigSelectorRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listEndpointConfigSelectorOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*EndpointConfigSelector
	for _, v := range m.EndpointConfigSelectors {
		res, err := unmarshalMapEndpointConfigSelector(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllEndpointConfigSelector(ctx context.Context, f func(*EndpointConfigSelector) bool, resources []*EndpointConfigSelector) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteEndpointConfigSelector(ctx, res)
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

type deleteEndpointConfigSelectorOperation struct{}

func (op *deleteEndpointConfigSelectorOperation) do(ctx context.Context, r *EndpointConfigSelector, c *Client) error {
	r, err := c.GetEndpointConfigSelector(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("EndpointConfigSelector not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetEndpointConfigSelector checking for existence. error: %v", err)
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

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetEndpointConfigSelector(ctx, r)
		if !dcl.IsNotFound(err) {
			if i == maxRetry {
				return dcl.NotDeletedError{ExistingResource: r}
			}
			time.Sleep(1000 * time.Millisecond)
		} else {
			break
		}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createEndpointConfigSelectorOperation struct {
	response map[string]interface{}
}

func (op *createEndpointConfigSelectorOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createEndpointConfigSelectorOperation) do(ctx context.Context, r *EndpointConfigSelector, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)
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
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetEndpointConfigSelector(ctx, r); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getEndpointConfigSelectorRaw(ctx context.Context, r *EndpointConfigSelector) ([]byte, error) {

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

func (c *Client) endpointConfigSelectorDiffsForRawDesired(ctx context.Context, rawDesired *EndpointConfigSelector, opts ...dcl.ApplyOption) (initial, desired *EndpointConfigSelector, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *EndpointConfigSelector
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*EndpointConfigSelector); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected EndpointConfigSelector, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetEndpointConfigSelector(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a EndpointConfigSelector resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve EndpointConfigSelector resource: %v", err)
		}
		c.Config.Logger.Info("Found that EndpointConfigSelector resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeEndpointConfigSelectorDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for EndpointConfigSelector: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for EndpointConfigSelector: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeEndpointConfigSelectorInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for EndpointConfigSelector: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeEndpointConfigSelectorDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for EndpointConfigSelector: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffEndpointConfigSelector(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeEndpointConfigSelectorInitialState(rawInitial, rawDesired *EndpointConfigSelector) (*EndpointConfigSelector, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeEndpointConfigSelectorDesiredState(rawDesired, rawInitial *EndpointConfigSelector, opts ...dcl.ApplyOption) (*EndpointConfigSelector, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.HttpFilters = canonicalizeEndpointConfigSelectorHttpFilters(rawDesired.HttpFilters, nil, opts...)
		rawDesired.EndpointMatcher = canonicalizeEndpointConfigSelectorEndpointMatcher(rawDesired.EndpointMatcher, nil, opts...)
		rawDesired.TrafficPortSelector = canonicalizeEndpointConfigSelectorTrafficPortSelector(rawDesired.TrafficPortSelector, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &EndpointConfigSelector{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	if dcl.IsZeroValue(rawDesired.Type) {
		canonicalDesired.Type = rawInitial.Type
	} else {
		canonicalDesired.Type = rawDesired.Type
	}
	if dcl.NameToSelfLink(rawDesired.AuthorizationPolicy, rawInitial.AuthorizationPolicy) {
		canonicalDesired.AuthorizationPolicy = rawInitial.AuthorizationPolicy
	} else {
		canonicalDesired.AuthorizationPolicy = rawDesired.AuthorizationPolicy
	}
	canonicalDesired.HttpFilters = canonicalizeEndpointConfigSelectorHttpFilters(rawDesired.HttpFilters, rawInitial.HttpFilters, opts...)
	canonicalDesired.EndpointMatcher = canonicalizeEndpointConfigSelectorEndpointMatcher(rawDesired.EndpointMatcher, rawInitial.EndpointMatcher, opts...)
	canonicalDesired.TrafficPortSelector = canonicalizeEndpointConfigSelectorTrafficPortSelector(rawDesired.TrafficPortSelector, rawInitial.TrafficPortSelector, opts...)
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.NameToSelfLink(rawDesired.ServerTlsPolicy, rawInitial.ServerTlsPolicy) {
		canonicalDesired.ServerTlsPolicy = rawInitial.ServerTlsPolicy
	} else {
		canonicalDesired.ServerTlsPolicy = rawDesired.ServerTlsPolicy
	}
	if dcl.NameToSelfLink(rawDesired.ClientTlsPolicy, rawInitial.ClientTlsPolicy) {
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

func canonicalizeEndpointConfigSelectorNewState(c *Client, rawNew, rawDesired *EndpointConfigSelector) (*EndpointConfigSelector, error) {

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
		if dcl.NameToSelfLink(rawDesired.AuthorizationPolicy, rawNew.AuthorizationPolicy) {
			rawNew.AuthorizationPolicy = rawDesired.AuthorizationPolicy
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.HttpFilters) && dcl.IsEmptyValueIndirect(rawDesired.HttpFilters) {
		rawNew.HttpFilters = rawDesired.HttpFilters
	} else {
		rawNew.HttpFilters = canonicalizeNewEndpointConfigSelectorHttpFilters(c, rawDesired.HttpFilters, rawNew.HttpFilters)
	}

	if dcl.IsEmptyValueIndirect(rawNew.EndpointMatcher) && dcl.IsEmptyValueIndirect(rawDesired.EndpointMatcher) {
		rawNew.EndpointMatcher = rawDesired.EndpointMatcher
	} else {
		rawNew.EndpointMatcher = canonicalizeNewEndpointConfigSelectorEndpointMatcher(c, rawDesired.EndpointMatcher, rawNew.EndpointMatcher)
	}

	if dcl.IsEmptyValueIndirect(rawNew.TrafficPortSelector) && dcl.IsEmptyValueIndirect(rawDesired.TrafficPortSelector) {
		rawNew.TrafficPortSelector = rawDesired.TrafficPortSelector
	} else {
		rawNew.TrafficPortSelector = canonicalizeNewEndpointConfigSelectorTrafficPortSelector(c, rawDesired.TrafficPortSelector, rawNew.TrafficPortSelector)
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
		if dcl.NameToSelfLink(rawDesired.ServerTlsPolicy, rawNew.ServerTlsPolicy) {
			rawNew.ServerTlsPolicy = rawDesired.ServerTlsPolicy
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ClientTlsPolicy) && dcl.IsEmptyValueIndirect(rawDesired.ClientTlsPolicy) {
		rawNew.ClientTlsPolicy = rawDesired.ClientTlsPolicy
	} else {
		if dcl.NameToSelfLink(rawDesired.ClientTlsPolicy, rawNew.ClientTlsPolicy) {
			rawNew.ClientTlsPolicy = rawDesired.ClientTlsPolicy
		}
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeEndpointConfigSelectorHttpFilters(des, initial *EndpointConfigSelectorHttpFilters, opts ...dcl.ApplyOption) *EndpointConfigSelectorHttpFilters {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EndpointConfigSelectorHttpFilters{}

	if dcl.IsZeroValue(des.HttpFilters) {
		des.HttpFilters = initial.HttpFilters
	} else {
		cDes.HttpFilters = des.HttpFilters
	}

	return cDes
}

func canonicalizeNewEndpointConfigSelectorHttpFilters(c *Client, des, nw *EndpointConfigSelectorHttpFilters) *EndpointConfigSelectorHttpFilters {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewEndpointConfigSelectorHttpFiltersSet(c *Client, des, nw []EndpointConfigSelectorHttpFilters) []EndpointConfigSelectorHttpFilters {
	if des == nil {
		return nw
	}
	var reorderedNew []EndpointConfigSelectorHttpFilters
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareEndpointConfigSelectorHttpFiltersNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewEndpointConfigSelectorHttpFiltersSlice(c *Client, des, nw []EndpointConfigSelectorHttpFilters) []EndpointConfigSelectorHttpFilters {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EndpointConfigSelectorHttpFilters
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEndpointConfigSelectorHttpFilters(c, &d, &n))
	}

	return items
}

func canonicalizeEndpointConfigSelectorEndpointMatcher(des, initial *EndpointConfigSelectorEndpointMatcher, opts ...dcl.ApplyOption) *EndpointConfigSelectorEndpointMatcher {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EndpointConfigSelectorEndpointMatcher{}

	cDes.MetadataLabelMatcher = canonicalizeEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher(des.MetadataLabelMatcher, initial.MetadataLabelMatcher, opts...)

	return cDes
}

func canonicalizeNewEndpointConfigSelectorEndpointMatcher(c *Client, des, nw *EndpointConfigSelectorEndpointMatcher) *EndpointConfigSelectorEndpointMatcher {
	if des == nil || nw == nil {
		return nw
	}

	nw.MetadataLabelMatcher = canonicalizeNewEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher(c, des.MetadataLabelMatcher, nw.MetadataLabelMatcher)

	return nw
}

func canonicalizeNewEndpointConfigSelectorEndpointMatcherSet(c *Client, des, nw []EndpointConfigSelectorEndpointMatcher) []EndpointConfigSelectorEndpointMatcher {
	if des == nil {
		return nw
	}
	var reorderedNew []EndpointConfigSelectorEndpointMatcher
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareEndpointConfigSelectorEndpointMatcherNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewEndpointConfigSelectorEndpointMatcherSlice(c *Client, des, nw []EndpointConfigSelectorEndpointMatcher) []EndpointConfigSelectorEndpointMatcher {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EndpointConfigSelectorEndpointMatcher
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEndpointConfigSelectorEndpointMatcher(c, &d, &n))
	}

	return items
}

func canonicalizeEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher(des, initial *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher, opts ...dcl.ApplyOption) *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher{}

	if dcl.IsZeroValue(des.MetadataLabelMatchCriteria) {
		des.MetadataLabelMatchCriteria = initial.MetadataLabelMatchCriteria
	} else {
		cDes.MetadataLabelMatchCriteria = des.MetadataLabelMatchCriteria
	}
	if dcl.IsZeroValue(des.MetadataLabels) {
		des.MetadataLabels = initial.MetadataLabels
	} else {
		cDes.MetadataLabels = des.MetadataLabels
	}

	return cDes
}

func canonicalizeNewEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher(c *Client, des, nw *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher) *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher {
	if des == nil || nw == nil {
		return nw
	}

	nw.MetadataLabels = canonicalizeNewEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelsSlice(c, des.MetadataLabels, nw.MetadataLabels)

	return nw
}

func canonicalizeNewEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherSet(c *Client, des, nw []EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher) []EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher {
	if des == nil {
		return nw
	}
	var reorderedNew []EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherSlice(c *Client, des, nw []EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher) []EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher(c, &d, &n))
	}

	return items
}

func canonicalizeEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels(des, initial *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels, opts ...dcl.ApplyOption) *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels{}

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

func canonicalizeNewEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels(c *Client, des, nw *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels) *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.LabelName, nw.LabelName) {
		nw.LabelName = des.LabelName
	}
	if dcl.StringCanonicalize(des.LabelValue, nw.LabelValue) {
		nw.LabelValue = des.LabelValue
	}

	return nw
}

func canonicalizeNewEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelsSet(c *Client, des, nw []EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels) []EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels {
	if des == nil {
		return nw
	}
	var reorderedNew []EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelsSlice(c *Client, des, nw []EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels) []EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels(c, &d, &n))
	}

	return items
}

func canonicalizeEndpointConfigSelectorTrafficPortSelector(des, initial *EndpointConfigSelectorTrafficPortSelector, opts ...dcl.ApplyOption) *EndpointConfigSelectorTrafficPortSelector {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EndpointConfigSelectorTrafficPortSelector{}

	if dcl.IsZeroValue(des.Ports) {
		des.Ports = initial.Ports
	} else {
		cDes.Ports = des.Ports
	}

	return cDes
}

func canonicalizeNewEndpointConfigSelectorTrafficPortSelector(c *Client, des, nw *EndpointConfigSelectorTrafficPortSelector) *EndpointConfigSelectorTrafficPortSelector {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewEndpointConfigSelectorTrafficPortSelectorSet(c *Client, des, nw []EndpointConfigSelectorTrafficPortSelector) []EndpointConfigSelectorTrafficPortSelector {
	if des == nil {
		return nw
	}
	var reorderedNew []EndpointConfigSelectorTrafficPortSelector
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareEndpointConfigSelectorTrafficPortSelectorNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewEndpointConfigSelectorTrafficPortSelectorSlice(c *Client, des, nw []EndpointConfigSelectorTrafficPortSelector) []EndpointConfigSelectorTrafficPortSelector {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EndpointConfigSelectorTrafficPortSelector
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEndpointConfigSelectorTrafficPortSelector(c, &d, &n))
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
func diffEndpointConfigSelector(c *Client, desired, actual *EndpointConfigSelector, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateEndpointConfigSelectorUpdateEndpointConfigSelectorOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateEndpointConfigSelectorUpdateEndpointConfigSelectorOperation")}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AuthorizationPolicy, actual.AuthorizationPolicy, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateEndpointConfigSelectorUpdateEndpointConfigSelectorOperation")}, fn.AddNest("AuthorizationPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HttpFilters, actual.HttpFilters, dcl.Info{ObjectFunction: compareEndpointConfigSelectorHttpFiltersNewStyle, EmptyObject: EmptyEndpointConfigSelectorHttpFilters, OperationSelector: dcl.TriggersOperation("updateEndpointConfigSelectorUpdateEndpointConfigSelectorOperation")}, fn.AddNest("HttpFilters")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EndpointMatcher, actual.EndpointMatcher, dcl.Info{ObjectFunction: compareEndpointConfigSelectorEndpointMatcherNewStyle, EmptyObject: EmptyEndpointConfigSelectorEndpointMatcher, OperationSelector: dcl.TriggersOperation("updateEndpointConfigSelectorUpdateEndpointConfigSelectorOperation")}, fn.AddNest("EndpointMatcher")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TrafficPortSelector, actual.TrafficPortSelector, dcl.Info{ObjectFunction: compareEndpointConfigSelectorTrafficPortSelectorNewStyle, EmptyObject: EmptyEndpointConfigSelectorTrafficPortSelector, OperationSelector: dcl.TriggersOperation("updateEndpointConfigSelectorUpdateEndpointConfigSelectorOperation")}, fn.AddNest("TrafficPortSelector")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateEndpointConfigSelectorUpdateEndpointConfigSelectorOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServerTlsPolicy, actual.ServerTlsPolicy, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateEndpointConfigSelectorUpdateEndpointConfigSelectorOperation")}, fn.AddNest("ServerTlsPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClientTlsPolicy, actual.ClientTlsPolicy, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateEndpointConfigSelectorUpdateEndpointConfigSelectorOperation")}, fn.AddNest("ClientTlsPolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareEndpointConfigSelectorHttpFiltersNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EndpointConfigSelectorHttpFilters)
	if !ok {
		desiredNotPointer, ok := d.(EndpointConfigSelectorHttpFilters)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointConfigSelectorHttpFilters or *EndpointConfigSelectorHttpFilters", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EndpointConfigSelectorHttpFilters)
	if !ok {
		actualNotPointer, ok := a.(EndpointConfigSelectorHttpFilters)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointConfigSelectorHttpFilters", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HttpFilters, actual.HttpFilters, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HttpFilters")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareEndpointConfigSelectorEndpointMatcherNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EndpointConfigSelectorEndpointMatcher)
	if !ok {
		desiredNotPointer, ok := d.(EndpointConfigSelectorEndpointMatcher)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointConfigSelectorEndpointMatcher or *EndpointConfigSelectorEndpointMatcher", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EndpointConfigSelectorEndpointMatcher)
	if !ok {
		actualNotPointer, ok := a.(EndpointConfigSelectorEndpointMatcher)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointConfigSelectorEndpointMatcher", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MetadataLabelMatcher, actual.MetadataLabelMatcher, dcl.Info{ObjectFunction: compareEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherNewStyle, EmptyObject: EmptyEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MetadataLabelMatcher")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher)
	if !ok {
		desiredNotPointer, ok := d.(EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher or *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher)
	if !ok {
		actualNotPointer, ok := a.(EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MetadataLabelMatchCriteria, actual.MetadataLabelMatchCriteria, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MetadataLabelMatchCriteria")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MetadataLabels, actual.MetadataLabels, dcl.Info{ObjectFunction: compareEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelsNewStyle, EmptyObject: EmptyEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MetadataLabels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels)
	if !ok {
		desiredNotPointer, ok := d.(EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels or *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels)
	if !ok {
		actualNotPointer, ok := a.(EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.LabelName, actual.LabelName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LabelName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LabelValue, actual.LabelValue, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LabelValue")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareEndpointConfigSelectorTrafficPortSelectorNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EndpointConfigSelectorTrafficPortSelector)
	if !ok {
		desiredNotPointer, ok := d.(EndpointConfigSelectorTrafficPortSelector)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointConfigSelectorTrafficPortSelector or *EndpointConfigSelectorTrafficPortSelector", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EndpointConfigSelectorTrafficPortSelector)
	if !ok {
		actualNotPointer, ok := a.(EndpointConfigSelectorTrafficPortSelector)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EndpointConfigSelectorTrafficPortSelector", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Ports, actual.Ports, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Ports")); len(ds) != 0 || err != nil {
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
func (r *EndpointConfigSelector) urlNormalized() *EndpointConfigSelector {
	normalized := dcl.Copy(*r).(EndpointConfigSelector)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.AuthorizationPolicy = dcl.SelfLinkToName(r.AuthorizationPolicy)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.ServerTlsPolicy = dcl.SelfLinkToName(r.ServerTlsPolicy)
	normalized.ClientTlsPolicy = dcl.SelfLinkToName(r.ClientTlsPolicy)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *EndpointConfigSelector) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateEndpointConfigSelector" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/endpointConfigSelectors/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the EndpointConfigSelector resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *EndpointConfigSelector) marshal(c *Client) ([]byte, error) {
	m, err := expandEndpointConfigSelector(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling EndpointConfigSelector: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalEndpointConfigSelector decodes JSON responses into the EndpointConfigSelector resource schema.
func unmarshalEndpointConfigSelector(b []byte, c *Client) (*EndpointConfigSelector, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapEndpointConfigSelector(m, c)
}

func unmarshalMapEndpointConfigSelector(m map[string]interface{}, c *Client) (*EndpointConfigSelector, error) {

	return flattenEndpointConfigSelector(c, m), nil
}

// expandEndpointConfigSelector expands EndpointConfigSelector into a JSON request object.
func expandEndpointConfigSelector(c *Client, f *EndpointConfigSelector) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/*/locations/global/endpointConfigSelectors/%s", f.Name, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.AuthorizationPolicy; !dcl.IsEmptyValueIndirect(v) {
		m["authorizationPolicy"] = v
	}
	if v, err := expandEndpointConfigSelectorHttpFilters(c, f.HttpFilters); err != nil {
		return nil, fmt.Errorf("error expanding HttpFilters into httpFilters: %w", err)
	} else if v != nil {
		m["httpFilters"] = v
	}
	if v, err := expandEndpointConfigSelectorEndpointMatcher(c, f.EndpointMatcher); err != nil {
		return nil, fmt.Errorf("error expanding EndpointMatcher into endpointMatcher: %w", err)
	} else if v != nil {
		m["endpointMatcher"] = v
	}
	if v, err := expandEndpointConfigSelectorTrafficPortSelector(c, f.TrafficPortSelector); err != nil {
		return nil, fmt.Errorf("error expanding TrafficPortSelector into trafficPortSelector: %w", err)
	} else if v != nil {
		m["trafficPortSelector"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.ServerTlsPolicy; !dcl.IsEmptyValueIndirect(v) {
		m["serverTlsPolicy"] = v
	}
	if v := f.ClientTlsPolicy; !dcl.IsEmptyValueIndirect(v) {
		m["clientTlsPolicy"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if v != nil {
		m["location"] = v
	}

	return m, nil
}

// flattenEndpointConfigSelector flattens EndpointConfigSelector from a JSON request object into the
// EndpointConfigSelector type.
func flattenEndpointConfigSelector(c *Client, i interface{}) *EndpointConfigSelector {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &EndpointConfigSelector{}
	res.Name = dcl.FlattenString(m["name"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.Type = flattenEndpointConfigSelectorTypeEnum(m["type"])
	res.AuthorizationPolicy = dcl.FlattenString(m["authorizationPolicy"])
	res.HttpFilters = flattenEndpointConfigSelectorHttpFilters(c, m["httpFilters"])
	res.EndpointMatcher = flattenEndpointConfigSelectorEndpointMatcher(c, m["endpointMatcher"])
	res.TrafficPortSelector = flattenEndpointConfigSelectorTrafficPortSelector(c, m["trafficPortSelector"])
	res.Description = dcl.FlattenString(m["description"])
	res.ServerTlsPolicy = dcl.FlattenString(m["serverTlsPolicy"])
	res.ClientTlsPolicy = dcl.FlattenString(m["clientTlsPolicy"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
}

// expandEndpointConfigSelectorHttpFiltersMap expands the contents of EndpointConfigSelectorHttpFilters into a JSON
// request object.
func expandEndpointConfigSelectorHttpFiltersMap(c *Client, f map[string]EndpointConfigSelectorHttpFilters) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEndpointConfigSelectorHttpFilters(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEndpointConfigSelectorHttpFiltersSlice expands the contents of EndpointConfigSelectorHttpFilters into a JSON
// request object.
func expandEndpointConfigSelectorHttpFiltersSlice(c *Client, f []EndpointConfigSelectorHttpFilters) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEndpointConfigSelectorHttpFilters(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEndpointConfigSelectorHttpFiltersMap flattens the contents of EndpointConfigSelectorHttpFilters from a JSON
// response object.
func flattenEndpointConfigSelectorHttpFiltersMap(c *Client, i interface{}) map[string]EndpointConfigSelectorHttpFilters {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EndpointConfigSelectorHttpFilters{}
	}

	if len(a) == 0 {
		return map[string]EndpointConfigSelectorHttpFilters{}
	}

	items := make(map[string]EndpointConfigSelectorHttpFilters)
	for k, item := range a {
		items[k] = *flattenEndpointConfigSelectorHttpFilters(c, item.(map[string]interface{}))
	}

	return items
}

// flattenEndpointConfigSelectorHttpFiltersSlice flattens the contents of EndpointConfigSelectorHttpFilters from a JSON
// response object.
func flattenEndpointConfigSelectorHttpFiltersSlice(c *Client, i interface{}) []EndpointConfigSelectorHttpFilters {
	a, ok := i.([]interface{})
	if !ok {
		return []EndpointConfigSelectorHttpFilters{}
	}

	if len(a) == 0 {
		return []EndpointConfigSelectorHttpFilters{}
	}

	items := make([]EndpointConfigSelectorHttpFilters, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEndpointConfigSelectorHttpFilters(c, item.(map[string]interface{})))
	}

	return items
}

// expandEndpointConfigSelectorHttpFilters expands an instance of EndpointConfigSelectorHttpFilters into a JSON
// request object.
func expandEndpointConfigSelectorHttpFilters(c *Client, f *EndpointConfigSelectorHttpFilters) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.HttpFilters; v != nil {
		m["httpFilters"] = v
	}

	return m, nil
}

// flattenEndpointConfigSelectorHttpFilters flattens an instance of EndpointConfigSelectorHttpFilters from a JSON
// response object.
func flattenEndpointConfigSelectorHttpFilters(c *Client, i interface{}) *EndpointConfigSelectorHttpFilters {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EndpointConfigSelectorHttpFilters{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEndpointConfigSelectorHttpFilters
	}
	r.HttpFilters = dcl.FlattenStringSlice(m["httpFilters"])

	return r
}

// expandEndpointConfigSelectorEndpointMatcherMap expands the contents of EndpointConfigSelectorEndpointMatcher into a JSON
// request object.
func expandEndpointConfigSelectorEndpointMatcherMap(c *Client, f map[string]EndpointConfigSelectorEndpointMatcher) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEndpointConfigSelectorEndpointMatcher(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEndpointConfigSelectorEndpointMatcherSlice expands the contents of EndpointConfigSelectorEndpointMatcher into a JSON
// request object.
func expandEndpointConfigSelectorEndpointMatcherSlice(c *Client, f []EndpointConfigSelectorEndpointMatcher) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEndpointConfigSelectorEndpointMatcher(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEndpointConfigSelectorEndpointMatcherMap flattens the contents of EndpointConfigSelectorEndpointMatcher from a JSON
// response object.
func flattenEndpointConfigSelectorEndpointMatcherMap(c *Client, i interface{}) map[string]EndpointConfigSelectorEndpointMatcher {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EndpointConfigSelectorEndpointMatcher{}
	}

	if len(a) == 0 {
		return map[string]EndpointConfigSelectorEndpointMatcher{}
	}

	items := make(map[string]EndpointConfigSelectorEndpointMatcher)
	for k, item := range a {
		items[k] = *flattenEndpointConfigSelectorEndpointMatcher(c, item.(map[string]interface{}))
	}

	return items
}

// flattenEndpointConfigSelectorEndpointMatcherSlice flattens the contents of EndpointConfigSelectorEndpointMatcher from a JSON
// response object.
func flattenEndpointConfigSelectorEndpointMatcherSlice(c *Client, i interface{}) []EndpointConfigSelectorEndpointMatcher {
	a, ok := i.([]interface{})
	if !ok {
		return []EndpointConfigSelectorEndpointMatcher{}
	}

	if len(a) == 0 {
		return []EndpointConfigSelectorEndpointMatcher{}
	}

	items := make([]EndpointConfigSelectorEndpointMatcher, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEndpointConfigSelectorEndpointMatcher(c, item.(map[string]interface{})))
	}

	return items
}

// expandEndpointConfigSelectorEndpointMatcher expands an instance of EndpointConfigSelectorEndpointMatcher into a JSON
// request object.
func expandEndpointConfigSelectorEndpointMatcher(c *Client, f *EndpointConfigSelectorEndpointMatcher) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher(c, f.MetadataLabelMatcher); err != nil {
		return nil, fmt.Errorf("error expanding MetadataLabelMatcher into metadataLabelMatcher: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["metadataLabelMatcher"] = v
	}

	return m, nil
}

// flattenEndpointConfigSelectorEndpointMatcher flattens an instance of EndpointConfigSelectorEndpointMatcher from a JSON
// response object.
func flattenEndpointConfigSelectorEndpointMatcher(c *Client, i interface{}) *EndpointConfigSelectorEndpointMatcher {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EndpointConfigSelectorEndpointMatcher{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEndpointConfigSelectorEndpointMatcher
	}
	r.MetadataLabelMatcher = flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher(c, m["metadataLabelMatcher"])

	return r
}

// expandEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMap expands the contents of EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher into a JSON
// request object.
func expandEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMap(c *Client, f map[string]EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherSlice expands the contents of EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher into a JSON
// request object.
func expandEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherSlice(c *Client, f []EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMap flattens the contents of EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher from a JSON
// response object.
func flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMap(c *Client, i interface{}) map[string]EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher{}
	}

	if len(a) == 0 {
		return map[string]EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher{}
	}

	items := make(map[string]EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher)
	for k, item := range a {
		items[k] = *flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher(c, item.(map[string]interface{}))
	}

	return items
}

// flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherSlice flattens the contents of EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher from a JSON
// response object.
func flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherSlice(c *Client, i interface{}) []EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher {
	a, ok := i.([]interface{})
	if !ok {
		return []EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher{}
	}

	if len(a) == 0 {
		return []EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher{}
	}

	items := make([]EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher(c, item.(map[string]interface{})))
	}

	return items
}

// expandEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher expands an instance of EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher into a JSON
// request object.
func expandEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher(c *Client, f *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MetadataLabelMatchCriteria; !dcl.IsEmptyValueIndirect(v) {
		m["metadataLabelMatchCriteria"] = v
	}
	if v, err := expandEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelsSlice(c, f.MetadataLabels); err != nil {
		return nil, fmt.Errorf("error expanding MetadataLabels into metadataLabels: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["metadataLabels"] = v
	}

	return m, nil
}

// flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher flattens an instance of EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher from a JSON
// response object.
func flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher(c *Client, i interface{}) *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher
	}
	r.MetadataLabelMatchCriteria = flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(m["metadataLabelMatchCriteria"])
	r.MetadataLabels = flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelsSlice(c, m["metadataLabels"])

	return r
}

// expandEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelsMap expands the contents of EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels into a JSON
// request object.
func expandEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelsMap(c *Client, f map[string]EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelsSlice expands the contents of EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels into a JSON
// request object.
func expandEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelsSlice(c *Client, f []EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelsMap flattens the contents of EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels from a JSON
// response object.
func flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelsMap(c *Client, i interface{}) map[string]EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels{}
	}

	if len(a) == 0 {
		return map[string]EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels{}
	}

	items := make(map[string]EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels)
	for k, item := range a {
		items[k] = *flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels(c, item.(map[string]interface{}))
	}

	return items
}

// flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelsSlice flattens the contents of EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels from a JSON
// response object.
func flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelsSlice(c *Client, i interface{}) []EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels {
	a, ok := i.([]interface{})
	if !ok {
		return []EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels{}
	}

	if len(a) == 0 {
		return []EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels{}
	}

	items := make([]EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels(c, item.(map[string]interface{})))
	}

	return items
}

// expandEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels expands an instance of EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels into a JSON
// request object.
func expandEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels(c *Client, f *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
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

// flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels flattens an instance of EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels from a JSON
// response object.
func flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels(c *Client, i interface{}) *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels
	}
	r.LabelName = dcl.FlattenString(m["labelName"])
	r.LabelValue = dcl.FlattenString(m["labelValue"])

	return r
}

// expandEndpointConfigSelectorTrafficPortSelectorMap expands the contents of EndpointConfigSelectorTrafficPortSelector into a JSON
// request object.
func expandEndpointConfigSelectorTrafficPortSelectorMap(c *Client, f map[string]EndpointConfigSelectorTrafficPortSelector) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEndpointConfigSelectorTrafficPortSelector(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEndpointConfigSelectorTrafficPortSelectorSlice expands the contents of EndpointConfigSelectorTrafficPortSelector into a JSON
// request object.
func expandEndpointConfigSelectorTrafficPortSelectorSlice(c *Client, f []EndpointConfigSelectorTrafficPortSelector) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEndpointConfigSelectorTrafficPortSelector(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEndpointConfigSelectorTrafficPortSelectorMap flattens the contents of EndpointConfigSelectorTrafficPortSelector from a JSON
// response object.
func flattenEndpointConfigSelectorTrafficPortSelectorMap(c *Client, i interface{}) map[string]EndpointConfigSelectorTrafficPortSelector {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EndpointConfigSelectorTrafficPortSelector{}
	}

	if len(a) == 0 {
		return map[string]EndpointConfigSelectorTrafficPortSelector{}
	}

	items := make(map[string]EndpointConfigSelectorTrafficPortSelector)
	for k, item := range a {
		items[k] = *flattenEndpointConfigSelectorTrafficPortSelector(c, item.(map[string]interface{}))
	}

	return items
}

// flattenEndpointConfigSelectorTrafficPortSelectorSlice flattens the contents of EndpointConfigSelectorTrafficPortSelector from a JSON
// response object.
func flattenEndpointConfigSelectorTrafficPortSelectorSlice(c *Client, i interface{}) []EndpointConfigSelectorTrafficPortSelector {
	a, ok := i.([]interface{})
	if !ok {
		return []EndpointConfigSelectorTrafficPortSelector{}
	}

	if len(a) == 0 {
		return []EndpointConfigSelectorTrafficPortSelector{}
	}

	items := make([]EndpointConfigSelectorTrafficPortSelector, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEndpointConfigSelectorTrafficPortSelector(c, item.(map[string]interface{})))
	}

	return items
}

// expandEndpointConfigSelectorTrafficPortSelector expands an instance of EndpointConfigSelectorTrafficPortSelector into a JSON
// request object.
func expandEndpointConfigSelectorTrafficPortSelector(c *Client, f *EndpointConfigSelectorTrafficPortSelector) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Ports; v != nil {
		m["ports"] = v
	}

	return m, nil
}

// flattenEndpointConfigSelectorTrafficPortSelector flattens an instance of EndpointConfigSelectorTrafficPortSelector from a JSON
// response object.
func flattenEndpointConfigSelectorTrafficPortSelector(c *Client, i interface{}) *EndpointConfigSelectorTrafficPortSelector {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EndpointConfigSelectorTrafficPortSelector{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEndpointConfigSelectorTrafficPortSelector
	}
	r.Ports = dcl.FlattenStringSlice(m["ports"])

	return r
}

// flattenEndpointConfigSelectorTypeEnumMap flattens the contents of EndpointConfigSelectorTypeEnum from a JSON
// response object.
func flattenEndpointConfigSelectorTypeEnumMap(c *Client, i interface{}) map[string]EndpointConfigSelectorTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EndpointConfigSelectorTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]EndpointConfigSelectorTypeEnum{}
	}

	items := make(map[string]EndpointConfigSelectorTypeEnum)
	for k, item := range a {
		items[k] = *flattenEndpointConfigSelectorTypeEnum(item.(interface{}))
	}

	return items
}

// flattenEndpointConfigSelectorTypeEnumSlice flattens the contents of EndpointConfigSelectorTypeEnum from a JSON
// response object.
func flattenEndpointConfigSelectorTypeEnumSlice(c *Client, i interface{}) []EndpointConfigSelectorTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []EndpointConfigSelectorTypeEnum{}
	}

	if len(a) == 0 {
		return []EndpointConfigSelectorTypeEnum{}
	}

	items := make([]EndpointConfigSelectorTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEndpointConfigSelectorTypeEnum(item.(interface{})))
	}

	return items
}

// flattenEndpointConfigSelectorTypeEnum asserts that an interface is a string, and returns a
// pointer to a *EndpointConfigSelectorTypeEnum with the same value as that string.
func flattenEndpointConfigSelectorTypeEnum(i interface{}) *EndpointConfigSelectorTypeEnum {
	s, ok := i.(string)
	if !ok {
		return EndpointConfigSelectorTypeEnumRef("")
	}

	return EndpointConfigSelectorTypeEnumRef(s)
}

// flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumMap flattens the contents of EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum from a JSON
// response object.
func flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumMap(c *Client, i interface{}) map[string]EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum{}
	}

	if len(a) == 0 {
		return map[string]EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum{}
	}

	items := make(map[string]EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum)
	for k, item := range a {
		items[k] = *flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(item.(interface{}))
	}

	return items
}

// flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumSlice flattens the contents of EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum from a JSON
// response object.
func flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumSlice(c *Client, i interface{}) []EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum{}
	}

	if len(a) == 0 {
		return []EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum{}
	}

	items := make([]EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(item.(interface{})))
	}

	return items
}

// flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum asserts that an interface is a string, and returns a
// pointer to a *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum with the same value as that string.
func flattenEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(i interface{}) *EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum {
	s, ok := i.(string)
	if !ok {
		return EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumRef("")
	}

	return EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *EndpointConfigSelector) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalEndpointConfigSelector(b, c)
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

type endpointConfigSelectorDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         endpointConfigSelectorApiOperation
}

func convertFieldDiffsToEndpointConfigSelectorDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]endpointConfigSelectorDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff in %q", ro, fd.FieldName)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []endpointConfigSelectorDiff
	// For each operation name, create a endpointConfigSelectorDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := endpointConfigSelectorDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToEndpointConfigSelectorApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToEndpointConfigSelectorApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (endpointConfigSelectorApiOperation, error) {
	switch opName {

	case "updateEndpointConfigSelectorUpdateEndpointConfigSelectorOperation":
		return &updateEndpointConfigSelectorUpdateEndpointConfigSelectorOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractEndpointConfigSelectorFields(r *EndpointConfigSelector) error {
	return nil
}
