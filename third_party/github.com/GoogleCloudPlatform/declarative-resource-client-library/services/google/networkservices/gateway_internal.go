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

func (r *Gateway) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "ports"); err != nil {
		return err
	}
	if err := dcl.Required(r, "scope"); err != nil {
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
func (r *Gateway) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://networkservices.googleapis.com/v1/", params)
}

func (r *Gateway) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/gateways/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Gateway) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/gateways", nr.basePath(), userBasePath, params), nil

}

func (r *Gateway) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/gateways?gatewayId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *Gateway) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/gateways/{{name}}", nr.basePath(), userBasePath, params), nil
}

// gatewayApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type gatewayApiOperation interface {
	do(context.Context, *Gateway, *Client) error
}

// newUpdateGatewayUpdateGatewayRequest creates a request for an
// Gateway resource's UpdateGateway update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateGatewayUpdateGatewayRequest(ctx context.Context, f *Gateway, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v, err := dcl.DeriveField("projects/%s/locations/global/gateways/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Name)); err != nil {
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
	if v := f.Addresses; v != nil {
		req["addresses"] = v
	}
	if v := f.Ports; v != nil {
		req["ports"] = v
	}
	if v := f.ServerTlsPolicy; !dcl.IsEmptyValueIndirect(v) {
		req["serverTlsPolicy"] = v
	}
	req["name"] = fmt.Sprintf("projects/%s/locations/%s/gateways/%s", *f.Project, *f.Location, *f.Name)

	return req, nil
}

// marshalUpdateGatewayUpdateGatewayRequest converts the update into
// the final JSON request body.
func marshalUpdateGatewayUpdateGatewayRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateGatewayUpdateGatewayOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateGatewayUpdateGatewayOperation) do(ctx context.Context, r *Gateway, c *Client) error {
	_, err := c.GetGateway(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateGateway")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateGatewayUpdateGatewayRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateGatewayUpdateGatewayRequest(c, req)
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

func (c *Client) listGatewayRaw(ctx context.Context, r *Gateway, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != GatewayMaxPage {
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

type listGatewayOperation struct {
	Gateways []map[string]interface{} `json:"gateways"`
	Token    string                   `json:"nextPageToken"`
}

func (c *Client) listGateway(ctx context.Context, r *Gateway, pageToken string, pageSize int32) ([]*Gateway, string, error) {
	b, err := c.listGatewayRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listGatewayOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Gateway
	for _, v := range m.Gateways {
		res, err := unmarshalMapGateway(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllGateway(ctx context.Context, f func(*Gateway) bool, resources []*Gateway) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteGateway(ctx, res)
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

type deleteGatewayOperation struct{}

func (op *deleteGatewayOperation) do(ctx context.Context, r *Gateway, c *Client) error {
	r, err := c.GetGateway(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Gateway not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetGateway checking for existence. error: %v", err)
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
		_, err := c.GetGateway(ctx, r)
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
type createGatewayOperation struct {
	response map[string]interface{}
}

func (op *createGatewayOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createGatewayOperation) do(ctx context.Context, r *Gateway, c *Client) error {
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

	if _, err := c.GetGateway(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getGatewayRaw(ctx context.Context, r *Gateway) ([]byte, error) {

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

func (c *Client) gatewayDiffsForRawDesired(ctx context.Context, rawDesired *Gateway, opts ...dcl.ApplyOption) (initial, desired *Gateway, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Gateway
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Gateway); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Gateway, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetGateway(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Gateway resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Gateway resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Gateway resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeGatewayDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Gateway: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Gateway: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractGatewayFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeGatewayInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Gateway: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeGatewayDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Gateway: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffGateway(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeGatewayInitialState(rawInitial, rawDesired *Gateway) (*Gateway, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeGatewayDesiredState(rawDesired, rawInitial *Gateway, opts ...dcl.ApplyOption) (*Gateway, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &Gateway{}
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
	if dcl.IsZeroValue(rawDesired.Type) || (dcl.IsEmptyValueIndirect(rawDesired.Type) && dcl.IsEmptyValueIndirect(rawInitial.Type)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Type = rawInitial.Type
	} else {
		canonicalDesired.Type = rawDesired.Type
	}
	if dcl.StringArrayCanonicalize(rawDesired.Addresses, rawInitial.Addresses) {
		canonicalDesired.Addresses = rawInitial.Addresses
	} else {
		canonicalDesired.Addresses = rawDesired.Addresses
	}
	if dcl.IsZeroValue(rawDesired.Ports) || (dcl.IsEmptyValueIndirect(rawDesired.Ports) && dcl.IsEmptyValueIndirect(rawInitial.Ports)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Ports = rawInitial.Ports
	} else {
		canonicalDesired.Ports = rawDesired.Ports
	}
	if dcl.StringCanonicalize(rawDesired.Scope, rawInitial.Scope) {
		canonicalDesired.Scope = rawInitial.Scope
	} else {
		canonicalDesired.Scope = rawDesired.Scope
	}
	if dcl.IsZeroValue(rawDesired.ServerTlsPolicy) || (dcl.IsEmptyValueIndirect(rawDesired.ServerTlsPolicy) && dcl.IsEmptyValueIndirect(rawInitial.ServerTlsPolicy)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.ServerTlsPolicy = rawInitial.ServerTlsPolicy
	} else {
		canonicalDesired.ServerTlsPolicy = rawDesired.ServerTlsPolicy
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

func canonicalizeGatewayNewState(c *Client, rawNew, rawDesired *Gateway) (*Gateway, error) {

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

	if dcl.IsEmptyValueIndirect(rawNew.Type) && dcl.IsEmptyValueIndirect(rawDesired.Type) {
		rawNew.Type = rawDesired.Type
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Addresses) && dcl.IsEmptyValueIndirect(rawDesired.Addresses) {
		rawNew.Addresses = rawDesired.Addresses
	} else {
		if dcl.StringArrayCanonicalize(rawDesired.Addresses, rawNew.Addresses) {
			rawNew.Addresses = rawDesired.Addresses
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Ports) && dcl.IsEmptyValueIndirect(rawDesired.Ports) {
		rawNew.Ports = rawDesired.Ports
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Scope) && dcl.IsEmptyValueIndirect(rawDesired.Scope) {
		rawNew.Scope = rawDesired.Scope
	} else {
		if dcl.StringCanonicalize(rawDesired.Scope, rawNew.Scope) {
			rawNew.Scope = rawDesired.Scope
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ServerTlsPolicy) && dcl.IsEmptyValueIndirect(rawDesired.ServerTlsPolicy) {
		rawNew.ServerTlsPolicy = rawDesired.ServerTlsPolicy
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

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffGateway(c *Client, desired, actual *Gateway, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateGatewayUpdateGatewayOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateGatewayUpdateGatewayOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateGatewayUpdateGatewayOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Addresses, actual.Addresses, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateGatewayUpdateGatewayOperation")}, fn.AddNest("Addresses")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Ports, actual.Ports, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateGatewayUpdateGatewayOperation")}, fn.AddNest("Ports")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Scope, actual.Scope, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Scope")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServerTlsPolicy, actual.ServerTlsPolicy, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateGatewayUpdateGatewayOperation")}, fn.AddNest("ServerTlsPolicy")); len(ds) != 0 || err != nil {
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

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Gateway) urlNormalized() *Gateway {
	normalized := dcl.Copy(*r).(Gateway)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Scope = dcl.SelfLinkToName(r.Scope)
	normalized.ServerTlsPolicy = r.ServerTlsPolicy
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	return &normalized
}

func (r *Gateway) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateGateway" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/gateways/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Gateway resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Gateway) marshal(c *Client) ([]byte, error) {
	m, err := expandGateway(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Gateway: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalGateway decodes JSON responses into the Gateway resource schema.
func unmarshalGateway(b []byte, c *Client, res *Gateway) (*Gateway, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapGateway(m, c, res)
}

func unmarshalMapGateway(m map[string]interface{}, c *Client, res *Gateway) (*Gateway, error) {

	flattened := flattenGateway(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandGateway expands Gateway into a JSON request object.
func expandGateway(c *Client, f *Gateway) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("projects/%s/locations/global/gateways/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Name)); err != nil {
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
	if v := f.Type; dcl.ValueShouldBeSent(v) {
		m["type"] = v
	}
	if v := f.Addresses; v != nil {
		m["addresses"] = v
	}
	if v := f.Ports; v != nil {
		m["ports"] = v
	}
	if v := f.Scope; dcl.ValueShouldBeSent(v) {
		m["scope"] = v
	}
	if v := f.ServerTlsPolicy; dcl.ValueShouldBeSent(v) {
		m["serverTlsPolicy"] = v
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

// flattenGateway flattens Gateway from a JSON request object into the
// Gateway type.
func flattenGateway(c *Client, i interface{}, res *Gateway) *Gateway {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Gateway{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.Type = flattenGatewayTypeEnum(m["type"])
	resultRes.Addresses = dcl.FlattenStringSlice(m["addresses"])
	resultRes.Ports = dcl.FlattenIntSlice(m["ports"])
	resultRes.Scope = dcl.FlattenString(m["scope"])
	resultRes.ServerTlsPolicy = dcl.FlattenString(m["serverTlsPolicy"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])
	resultRes.SelfLink = dcl.FlattenString(m["selfLink"])

	return resultRes
}

// flattenGatewayTypeEnumMap flattens the contents of GatewayTypeEnum from a JSON
// response object.
func flattenGatewayTypeEnumMap(c *Client, i interface{}, res *Gateway) map[string]GatewayTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GatewayTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]GatewayTypeEnum{}
	}

	items := make(map[string]GatewayTypeEnum)
	for k, item := range a {
		items[k] = *flattenGatewayTypeEnum(item.(interface{}))
	}

	return items
}

// flattenGatewayTypeEnumSlice flattens the contents of GatewayTypeEnum from a JSON
// response object.
func flattenGatewayTypeEnumSlice(c *Client, i interface{}, res *Gateway) []GatewayTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []GatewayTypeEnum{}
	}

	if len(a) == 0 {
		return []GatewayTypeEnum{}
	}

	items := make([]GatewayTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGatewayTypeEnum(item.(interface{})))
	}

	return items
}

// flattenGatewayTypeEnum asserts that an interface is a string, and returns a
// pointer to a *GatewayTypeEnum with the same value as that string.
func flattenGatewayTypeEnum(i interface{}) *GatewayTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return GatewayTypeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Gateway) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalGateway(b, c, r)
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

type gatewayDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         gatewayApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToGatewayDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]gatewayDiff, error) {
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
	var diffs []gatewayDiff
	// For each operation name, create a gatewayDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := gatewayDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToGatewayApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToGatewayApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (gatewayApiOperation, error) {
	switch opName {

	case "updateGatewayUpdateGatewayOperation":
		return &updateGatewayUpdateGatewayOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractGatewayFields(r *Gateway) error {
	return nil
}

func postReadExtractGatewayFields(r *Gateway) error {
	return nil
}
