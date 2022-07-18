// Copyright 2022 Google LLC. All Rights Reserved.
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
package iam

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

func (r *WorkforcePoolProvider) validate() error {

	if err := dcl.ValidateExactlyOneOfFieldsSet([]string{"Saml", "Oidc"}, r.Saml, r.Oidc); err != nil {
		return err
	}
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "attributeMapping"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.WorkforcePool, "WorkforcePool"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Saml) {
		if err := r.Saml.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Oidc) {
		if err := r.Oidc.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *WorkforcePoolProviderSaml) validate() error {
	if err := dcl.Required(r, "idpMetadataXml"); err != nil {
		return err
	}
	return nil
}
func (r *WorkforcePoolProviderOidc) validate() error {
	if err := dcl.Required(r, "issuerUri"); err != nil {
		return err
	}
	if err := dcl.Required(r, "clientId"); err != nil {
		return err
	}
	return nil
}
func (r *WorkforcePoolProvider) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://iam.googleapis.com/v1/", params)
}

func (r *WorkforcePoolProvider) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location":      dcl.ValueOrEmptyString(nr.Location),
		"workforcePool": dcl.ValueOrEmptyString(nr.WorkforcePool),
		"name":          dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("locations/{{location}}/workforcePools/{{workforcePool}}/providers/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *WorkforcePoolProvider) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location":      dcl.ValueOrEmptyString(nr.Location),
		"workforcePool": dcl.ValueOrEmptyString(nr.WorkforcePool),
	}
	return dcl.URL("locations/{{location}}/workforcePools/{{workforcePool}}/providers", nr.basePath(), userBasePath, params), nil

}

func (r *WorkforcePoolProvider) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location":      dcl.ValueOrEmptyString(nr.Location),
		"workforcePool": dcl.ValueOrEmptyString(nr.WorkforcePool),
		"name":          dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("locations/{{location}}/workforcePools/{{workforcePool}}/providers?workforcePoolProviderId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *WorkforcePoolProvider) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location":      dcl.ValueOrEmptyString(nr.Location),
		"workforcePool": dcl.ValueOrEmptyString(nr.WorkforcePool),
		"name":          dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("locations/{{location}}/workforcePools/{{workforcePool}}/providers/{{name}}", nr.basePath(), userBasePath, params), nil
}

// workforcePoolProviderApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type workforcePoolProviderApiOperation interface {
	do(context.Context, *WorkforcePoolProvider, *Client) error
}

// newUpdateWorkforcePoolProviderUpdateWorkforcePoolProviderRequest creates a request for an
// WorkforcePoolProvider resource's UpdateWorkforcePoolProvider update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateWorkforcePoolProviderUpdateWorkforcePoolProviderRequest(ctx context.Context, f *WorkforcePoolProvider, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Disabled; !dcl.IsEmptyValueIndirect(v) {
		req["disabled"] = v
	}
	if v := f.AttributeMapping; !dcl.IsEmptyValueIndirect(v) {
		req["attributeMapping"] = v
	}
	if v := f.AttributeCondition; !dcl.IsEmptyValueIndirect(v) {
		req["attributeCondition"] = v
	}
	if v, err := expandWorkforcePoolProviderSaml(c, f.Saml, res); err != nil {
		return nil, fmt.Errorf("error expanding Saml into saml: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["saml"] = v
	}
	if v, err := expandWorkforcePoolProviderOidc(c, f.Oidc, res); err != nil {
		return nil, fmt.Errorf("error expanding Oidc into oidc: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["oidc"] = v
	}
	return req, nil
}

// marshalUpdateWorkforcePoolProviderUpdateWorkforcePoolProviderRequest converts the update into
// the final JSON request body.
func marshalUpdateWorkforcePoolProviderUpdateWorkforcePoolProviderRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateWorkforcePoolProviderUpdateWorkforcePoolProviderOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateWorkforcePoolProviderUpdateWorkforcePoolProviderOperation) do(ctx context.Context, r *WorkforcePoolProvider, c *Client) error {
	_, err := c.GetWorkforcePoolProvider(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateWorkforcePoolProvider")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateWorkforcePoolProviderUpdateWorkforcePoolProviderRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateWorkforcePoolProviderUpdateWorkforcePoolProviderRequest(c, req)
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

func (c *Client) listWorkforcePoolProviderRaw(ctx context.Context, r *WorkforcePoolProvider, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != WorkforcePoolProviderMaxPage {
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

type listWorkforcePoolProviderOperation struct {
	WorkforcePoolProviders []map[string]interface{} `json:"workforcePoolProviders"`
	Token                  string                   `json:"nextPageToken"`
}

func (c *Client) listWorkforcePoolProvider(ctx context.Context, r *WorkforcePoolProvider, pageToken string, pageSize int32) ([]*WorkforcePoolProvider, string, error) {
	b, err := c.listWorkforcePoolProviderRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listWorkforcePoolProviderOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*WorkforcePoolProvider
	for _, v := range m.WorkforcePoolProviders {
		res, err := unmarshalMapWorkforcePoolProvider(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Location = r.Location
		res.WorkforcePool = r.WorkforcePool
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllWorkforcePoolProvider(ctx context.Context, f func(*WorkforcePoolProvider) bool, resources []*WorkforcePoolProvider) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteWorkforcePoolProvider(ctx, res)
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

type deleteWorkforcePoolProviderOperation struct{}

func (op *deleteWorkforcePoolProviderOperation) do(ctx context.Context, r *WorkforcePoolProvider, c *Client) error {
	r, err := c.GetWorkforcePoolProvider(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "WorkforcePoolProvider not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetWorkforcePoolProvider checking for existence. error: %v", err)
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
		return fmt.Errorf("failed to delete WorkforcePoolProvider: %w", err)
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createWorkforcePoolProviderOperation struct {
	response map[string]interface{}
}

func (op *createWorkforcePoolProviderOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createWorkforcePoolProviderOperation) do(ctx context.Context, r *WorkforcePoolProvider, c *Client) error {
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

	if _, err := c.GetWorkforcePoolProvider(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getWorkforcePoolProviderRaw(ctx context.Context, r *WorkforcePoolProvider) ([]byte, error) {

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

func (c *Client) workforcePoolProviderDiffsForRawDesired(ctx context.Context, rawDesired *WorkforcePoolProvider, opts ...dcl.ApplyOption) (initial, desired *WorkforcePoolProvider, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *WorkforcePoolProvider
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*WorkforcePoolProvider); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected WorkforcePoolProvider, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetWorkforcePoolProvider(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a WorkforcePoolProvider resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve WorkforcePoolProvider resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that WorkforcePoolProvider resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeWorkforcePoolProviderDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for WorkforcePoolProvider: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for WorkforcePoolProvider: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractWorkforcePoolProviderFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeWorkforcePoolProviderInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for WorkforcePoolProvider: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeWorkforcePoolProviderDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for WorkforcePoolProvider: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffWorkforcePoolProvider(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeWorkforcePoolProviderInitialState(rawInitial, rawDesired *WorkforcePoolProvider) (*WorkforcePoolProvider, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.

	if !dcl.IsZeroValue(rawInitial.Saml) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Oidc) {
			rawInitial.Saml = EmptyWorkforcePoolProviderSaml
		}
	}

	if !dcl.IsZeroValue(rawInitial.Oidc) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Saml) {
			rawInitial.Oidc = EmptyWorkforcePoolProviderOidc
		}
	}

	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeWorkforcePoolProviderDesiredState(rawDesired, rawInitial *WorkforcePoolProvider, opts ...dcl.ApplyOption) (*WorkforcePoolProvider, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Saml = canonicalizeWorkforcePoolProviderSaml(rawDesired.Saml, nil, opts...)
		rawDesired.Oidc = canonicalizeWorkforcePoolProviderOidc(rawDesired.Oidc, nil, opts...)

		return rawDesired, nil
	}

	if rawDesired.Saml != nil || rawInitial.Saml != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Oidc) {
			rawDesired.Saml = nil
			rawInitial.Saml = nil
		}
	}

	if rawDesired.Oidc != nil || rawInitial.Oidc != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Saml) {
			rawDesired.Oidc = nil
			rawInitial.Oidc = nil
		}
	}

	canonicalDesired := &WorkforcePoolProvider{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
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
	if dcl.BoolCanonicalize(rawDesired.Disabled, rawInitial.Disabled) {
		canonicalDesired.Disabled = rawInitial.Disabled
	} else {
		canonicalDesired.Disabled = rawDesired.Disabled
	}
	if dcl.IsZeroValue(rawDesired.AttributeMapping) || (dcl.IsEmptyValueIndirect(rawDesired.AttributeMapping) && dcl.IsEmptyValueIndirect(rawInitial.AttributeMapping)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.AttributeMapping = rawInitial.AttributeMapping
	} else {
		canonicalDesired.AttributeMapping = rawDesired.AttributeMapping
	}
	if dcl.StringCanonicalize(rawDesired.AttributeCondition, rawInitial.AttributeCondition) {
		canonicalDesired.AttributeCondition = rawInitial.AttributeCondition
	} else {
		canonicalDesired.AttributeCondition = rawDesired.AttributeCondition
	}
	canonicalDesired.Saml = canonicalizeWorkforcePoolProviderSaml(rawDesired.Saml, rawInitial.Saml, opts...)
	canonicalDesired.Oidc = canonicalizeWorkforcePoolProviderOidc(rawDesired.Oidc, rawInitial.Oidc, opts...)
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}
	if dcl.NameToSelfLink(rawDesired.WorkforcePool, rawInitial.WorkforcePool) {
		canonicalDesired.WorkforcePool = rawInitial.WorkforcePool
	} else {
		canonicalDesired.WorkforcePool = rawDesired.WorkforcePool
	}

	return canonicalDesired, nil
}

func canonicalizeWorkforcePoolProviderNewState(c *Client, rawNew, rawDesired *WorkforcePoolProvider) (*WorkforcePoolProvider, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
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

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Disabled) && dcl.IsEmptyValueIndirect(rawDesired.Disabled) {
		rawNew.Disabled = rawDesired.Disabled
	} else {
		if dcl.BoolCanonicalize(rawDesired.Disabled, rawNew.Disabled) {
			rawNew.Disabled = rawDesired.Disabled
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.AttributeMapping) && dcl.IsEmptyValueIndirect(rawDesired.AttributeMapping) {
		rawNew.AttributeMapping = rawDesired.AttributeMapping
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.AttributeCondition) && dcl.IsEmptyValueIndirect(rawDesired.AttributeCondition) {
		rawNew.AttributeCondition = rawDesired.AttributeCondition
	} else {
		if dcl.StringCanonicalize(rawDesired.AttributeCondition, rawNew.AttributeCondition) {
			rawNew.AttributeCondition = rawDesired.AttributeCondition
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Saml) && dcl.IsEmptyValueIndirect(rawDesired.Saml) {
		rawNew.Saml = rawDesired.Saml
	} else {
		rawNew.Saml = canonicalizeNewWorkforcePoolProviderSaml(c, rawDesired.Saml, rawNew.Saml)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Oidc) && dcl.IsEmptyValueIndirect(rawDesired.Oidc) {
		rawNew.Oidc = rawDesired.Oidc
	} else {
		rawNew.Oidc = canonicalizeNewWorkforcePoolProviderOidc(c, rawDesired.Oidc, rawNew.Oidc)
	}

	rawNew.Location = rawDesired.Location

	rawNew.WorkforcePool = rawDesired.WorkforcePool

	return rawNew, nil
}

func canonicalizeWorkforcePoolProviderSaml(des, initial *WorkforcePoolProviderSaml, opts ...dcl.ApplyOption) *WorkforcePoolProviderSaml {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &WorkforcePoolProviderSaml{}

	if dcl.StringCanonicalize(des.IdpMetadataXml, initial.IdpMetadataXml) || dcl.IsZeroValue(des.IdpMetadataXml) {
		cDes.IdpMetadataXml = initial.IdpMetadataXml
	} else {
		cDes.IdpMetadataXml = des.IdpMetadataXml
	}

	return cDes
}

func canonicalizeWorkforcePoolProviderSamlSlice(des, initial []WorkforcePoolProviderSaml, opts ...dcl.ApplyOption) []WorkforcePoolProviderSaml {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]WorkforcePoolProviderSaml, 0, len(des))
		for _, d := range des {
			cd := canonicalizeWorkforcePoolProviderSaml(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]WorkforcePoolProviderSaml, 0, len(des))
	for i, d := range des {
		cd := canonicalizeWorkforcePoolProviderSaml(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewWorkforcePoolProviderSaml(c *Client, des, nw *WorkforcePoolProviderSaml) *WorkforcePoolProviderSaml {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for WorkforcePoolProviderSaml while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.IdpMetadataXml, nw.IdpMetadataXml) {
		nw.IdpMetadataXml = des.IdpMetadataXml
	}

	return nw
}

func canonicalizeNewWorkforcePoolProviderSamlSet(c *Client, des, nw []WorkforcePoolProviderSaml) []WorkforcePoolProviderSaml {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkforcePoolProviderSaml
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareWorkforcePoolProviderSamlNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewWorkforcePoolProviderSamlSlice(c *Client, des, nw []WorkforcePoolProviderSaml) []WorkforcePoolProviderSaml {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []WorkforcePoolProviderSaml
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkforcePoolProviderSaml(c, &d, &n))
	}

	return items
}

func canonicalizeWorkforcePoolProviderOidc(des, initial *WorkforcePoolProviderOidc, opts ...dcl.ApplyOption) *WorkforcePoolProviderOidc {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &WorkforcePoolProviderOidc{}

	if dcl.StringCanonicalize(des.IssuerUri, initial.IssuerUri) || dcl.IsZeroValue(des.IssuerUri) {
		cDes.IssuerUri = initial.IssuerUri
	} else {
		cDes.IssuerUri = des.IssuerUri
	}
	if dcl.StringCanonicalize(des.ClientId, initial.ClientId) || dcl.IsZeroValue(des.ClientId) {
		cDes.ClientId = initial.ClientId
	} else {
		cDes.ClientId = des.ClientId
	}

	return cDes
}

func canonicalizeWorkforcePoolProviderOidcSlice(des, initial []WorkforcePoolProviderOidc, opts ...dcl.ApplyOption) []WorkforcePoolProviderOidc {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]WorkforcePoolProviderOidc, 0, len(des))
		for _, d := range des {
			cd := canonicalizeWorkforcePoolProviderOidc(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]WorkforcePoolProviderOidc, 0, len(des))
	for i, d := range des {
		cd := canonicalizeWorkforcePoolProviderOidc(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewWorkforcePoolProviderOidc(c *Client, des, nw *WorkforcePoolProviderOidc) *WorkforcePoolProviderOidc {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for WorkforcePoolProviderOidc while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.IssuerUri, nw.IssuerUri) {
		nw.IssuerUri = des.IssuerUri
	}
	if dcl.StringCanonicalize(des.ClientId, nw.ClientId) {
		nw.ClientId = des.ClientId
	}

	return nw
}

func canonicalizeNewWorkforcePoolProviderOidcSet(c *Client, des, nw []WorkforcePoolProviderOidc) []WorkforcePoolProviderOidc {
	if des == nil {
		return nw
	}
	var reorderedNew []WorkforcePoolProviderOidc
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareWorkforcePoolProviderOidcNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewWorkforcePoolProviderOidcSlice(c *Client, des, nw []WorkforcePoolProviderOidc) []WorkforcePoolProviderOidc {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []WorkforcePoolProviderOidc
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkforcePoolProviderOidc(c, &d, &n))
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
func diffWorkforcePoolProvider(c *Client, desired, actual *WorkforcePoolProvider, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateWorkforcePoolProviderUpdateWorkforcePoolProviderOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateWorkforcePoolProviderUpdateWorkforcePoolProviderOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Disabled, actual.Disabled, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateWorkforcePoolProviderUpdateWorkforcePoolProviderOperation")}, fn.AddNest("Disabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AttributeMapping, actual.AttributeMapping, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateWorkforcePoolProviderUpdateWorkforcePoolProviderOperation")}, fn.AddNest("AttributeMapping")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AttributeCondition, actual.AttributeCondition, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateWorkforcePoolProviderUpdateWorkforcePoolProviderOperation")}, fn.AddNest("AttributeCondition")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Saml, actual.Saml, dcl.DiffInfo{ObjectFunction: compareWorkforcePoolProviderSamlNewStyle, EmptyObject: EmptyWorkforcePoolProviderSaml, OperationSelector: dcl.TriggersOperation("updateWorkforcePoolProviderUpdateWorkforcePoolProviderOperation")}, fn.AddNest("Saml")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Oidc, actual.Oidc, dcl.DiffInfo{ObjectFunction: compareWorkforcePoolProviderOidcNewStyle, EmptyObject: EmptyWorkforcePoolProviderOidc, OperationSelector: dcl.TriggersOperation("updateWorkforcePoolProviderUpdateWorkforcePoolProviderOperation")}, fn.AddNest("Oidc")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.WorkforcePool, actual.WorkforcePool, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("WorkforcePool")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareWorkforcePoolProviderSamlNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*WorkforcePoolProviderSaml)
	if !ok {
		desiredNotPointer, ok := d.(WorkforcePoolProviderSaml)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a WorkforcePoolProviderSaml or *WorkforcePoolProviderSaml", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*WorkforcePoolProviderSaml)
	if !ok {
		actualNotPointer, ok := a.(WorkforcePoolProviderSaml)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a WorkforcePoolProviderSaml", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.IdpMetadataXml, actual.IdpMetadataXml, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateWorkforcePoolProviderUpdateWorkforcePoolProviderOperation")}, fn.AddNest("IdpMetadataXml")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareWorkforcePoolProviderOidcNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*WorkforcePoolProviderOidc)
	if !ok {
		desiredNotPointer, ok := d.(WorkforcePoolProviderOidc)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a WorkforcePoolProviderOidc or *WorkforcePoolProviderOidc", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*WorkforcePoolProviderOidc)
	if !ok {
		actualNotPointer, ok := a.(WorkforcePoolProviderOidc)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a WorkforcePoolProviderOidc", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.IssuerUri, actual.IssuerUri, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateWorkforcePoolProviderUpdateWorkforcePoolProviderOperation")}, fn.AddNest("IssuerUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClientId, actual.ClientId, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateWorkforcePoolProviderUpdateWorkforcePoolProviderOperation")}, fn.AddNest("ClientId")); len(ds) != 0 || err != nil {
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
func (r *WorkforcePoolProvider) urlNormalized() *WorkforcePoolProvider {
	normalized := dcl.Copy(*r).(WorkforcePoolProvider)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.AttributeCondition = dcl.SelfLinkToName(r.AttributeCondition)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.WorkforcePool = dcl.SelfLinkToName(r.WorkforcePool)
	return &normalized
}

func (r *WorkforcePoolProvider) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateWorkforcePoolProvider" {
		fields := map[string]interface{}{
			"location":      dcl.ValueOrEmptyString(nr.Location),
			"workforcePool": dcl.ValueOrEmptyString(nr.WorkforcePool),
			"name":          dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("locations/{{location}}/workforcePools/{{workforcePool}}/providers/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the WorkforcePoolProvider resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *WorkforcePoolProvider) marshal(c *Client) ([]byte, error) {
	m, err := expandWorkforcePoolProvider(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling WorkforcePoolProvider: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalWorkforcePoolProvider decodes JSON responses into the WorkforcePoolProvider resource schema.
func unmarshalWorkforcePoolProvider(b []byte, c *Client, res *WorkforcePoolProvider) (*WorkforcePoolProvider, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapWorkforcePoolProvider(m, c, res)
}

func unmarshalMapWorkforcePoolProvider(m map[string]interface{}, c *Client, res *WorkforcePoolProvider) (*WorkforcePoolProvider, error) {

	flattened := flattenWorkforcePoolProvider(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandWorkforcePoolProvider expands WorkforcePoolProvider into a JSON request object.
func expandWorkforcePoolProvider(c *Client, f *WorkforcePoolProvider) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("locations/%s/workforcePools/%s/providers/%s", f.Name, dcl.SelfLinkToName(f.Location), dcl.SelfLinkToName(f.WorkforcePool), dcl.SelfLinkToName(f.Name)); err != nil {
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
	if v := f.Disabled; dcl.ValueShouldBeSent(v) {
		m["disabled"] = v
	}
	if v := f.AttributeMapping; dcl.ValueShouldBeSent(v) {
		m["attributeMapping"] = v
	}
	if v := f.AttributeCondition; dcl.ValueShouldBeSent(v) {
		m["attributeCondition"] = v
	}
	if v, err := expandWorkforcePoolProviderSaml(c, f.Saml, res); err != nil {
		return nil, fmt.Errorf("error expanding Saml into saml: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["saml"] = v
	}
	if v, err := expandWorkforcePoolProviderOidc(c, f.Oidc, res); err != nil {
		return nil, fmt.Errorf("error expanding Oidc into oidc: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["oidc"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding WorkforcePool into workforcePool: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["workforcePool"] = v
	}

	return m, nil
}

// flattenWorkforcePoolProvider flattens WorkforcePoolProvider from a JSON request object into the
// WorkforcePoolProvider type.
func flattenWorkforcePoolProvider(c *Client, i interface{}, res *WorkforcePoolProvider) *WorkforcePoolProvider {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &WorkforcePoolProvider{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.DisplayName = dcl.FlattenString(m["displayName"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.State = flattenWorkforcePoolProviderStateEnum(m["state"])
	resultRes.Disabled = dcl.FlattenBool(m["disabled"])
	resultRes.AttributeMapping = dcl.FlattenKeyValuePairs(m["attributeMapping"])
	resultRes.AttributeCondition = dcl.FlattenString(m["attributeCondition"])
	resultRes.Saml = flattenWorkforcePoolProviderSaml(c, m["saml"], res)
	resultRes.Oidc = flattenWorkforcePoolProviderOidc(c, m["oidc"], res)
	resultRes.Location = dcl.FlattenString(m["location"])
	resultRes.WorkforcePool = dcl.FlattenString(m["workforcePool"])

	return resultRes
}

// expandWorkforcePoolProviderSamlMap expands the contents of WorkforcePoolProviderSaml into a JSON
// request object.
func expandWorkforcePoolProviderSamlMap(c *Client, f map[string]WorkforcePoolProviderSaml, res *WorkforcePoolProvider) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkforcePoolProviderSaml(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkforcePoolProviderSamlSlice expands the contents of WorkforcePoolProviderSaml into a JSON
// request object.
func expandWorkforcePoolProviderSamlSlice(c *Client, f []WorkforcePoolProviderSaml, res *WorkforcePoolProvider) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkforcePoolProviderSaml(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkforcePoolProviderSamlMap flattens the contents of WorkforcePoolProviderSaml from a JSON
// response object.
func flattenWorkforcePoolProviderSamlMap(c *Client, i interface{}, res *WorkforcePoolProvider) map[string]WorkforcePoolProviderSaml {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkforcePoolProviderSaml{}
	}

	if len(a) == 0 {
		return map[string]WorkforcePoolProviderSaml{}
	}

	items := make(map[string]WorkforcePoolProviderSaml)
	for k, item := range a {
		items[k] = *flattenWorkforcePoolProviderSaml(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenWorkforcePoolProviderSamlSlice flattens the contents of WorkforcePoolProviderSaml from a JSON
// response object.
func flattenWorkforcePoolProviderSamlSlice(c *Client, i interface{}, res *WorkforcePoolProvider) []WorkforcePoolProviderSaml {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkforcePoolProviderSaml{}
	}

	if len(a) == 0 {
		return []WorkforcePoolProviderSaml{}
	}

	items := make([]WorkforcePoolProviderSaml, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkforcePoolProviderSaml(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandWorkforcePoolProviderSaml expands an instance of WorkforcePoolProviderSaml into a JSON
// request object.
func expandWorkforcePoolProviderSaml(c *Client, f *WorkforcePoolProviderSaml, res *WorkforcePoolProvider) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.IdpMetadataXml; !dcl.IsEmptyValueIndirect(v) {
		m["idpMetadataXml"] = v
	}

	return m, nil
}

// flattenWorkforcePoolProviderSaml flattens an instance of WorkforcePoolProviderSaml from a JSON
// response object.
func flattenWorkforcePoolProviderSaml(c *Client, i interface{}, res *WorkforcePoolProvider) *WorkforcePoolProviderSaml {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkforcePoolProviderSaml{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyWorkforcePoolProviderSaml
	}
	r.IdpMetadataXml = dcl.FlattenString(m["idpMetadataXml"])

	return r
}

// expandWorkforcePoolProviderOidcMap expands the contents of WorkforcePoolProviderOidc into a JSON
// request object.
func expandWorkforcePoolProviderOidcMap(c *Client, f map[string]WorkforcePoolProviderOidc, res *WorkforcePoolProvider) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkforcePoolProviderOidc(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkforcePoolProviderOidcSlice expands the contents of WorkforcePoolProviderOidc into a JSON
// request object.
func expandWorkforcePoolProviderOidcSlice(c *Client, f []WorkforcePoolProviderOidc, res *WorkforcePoolProvider) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkforcePoolProviderOidc(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkforcePoolProviderOidcMap flattens the contents of WorkforcePoolProviderOidc from a JSON
// response object.
func flattenWorkforcePoolProviderOidcMap(c *Client, i interface{}, res *WorkforcePoolProvider) map[string]WorkforcePoolProviderOidc {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkforcePoolProviderOidc{}
	}

	if len(a) == 0 {
		return map[string]WorkforcePoolProviderOidc{}
	}

	items := make(map[string]WorkforcePoolProviderOidc)
	for k, item := range a {
		items[k] = *flattenWorkforcePoolProviderOidc(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenWorkforcePoolProviderOidcSlice flattens the contents of WorkforcePoolProviderOidc from a JSON
// response object.
func flattenWorkforcePoolProviderOidcSlice(c *Client, i interface{}, res *WorkforcePoolProvider) []WorkforcePoolProviderOidc {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkforcePoolProviderOidc{}
	}

	if len(a) == 0 {
		return []WorkforcePoolProviderOidc{}
	}

	items := make([]WorkforcePoolProviderOidc, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkforcePoolProviderOidc(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandWorkforcePoolProviderOidc expands an instance of WorkforcePoolProviderOidc into a JSON
// request object.
func expandWorkforcePoolProviderOidc(c *Client, f *WorkforcePoolProviderOidc, res *WorkforcePoolProvider) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.IssuerUri; !dcl.IsEmptyValueIndirect(v) {
		m["issuerUri"] = v
	}
	if v := f.ClientId; !dcl.IsEmptyValueIndirect(v) {
		m["clientId"] = v
	}

	return m, nil
}

// flattenWorkforcePoolProviderOidc flattens an instance of WorkforcePoolProviderOidc from a JSON
// response object.
func flattenWorkforcePoolProviderOidc(c *Client, i interface{}, res *WorkforcePoolProvider) *WorkforcePoolProviderOidc {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkforcePoolProviderOidc{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyWorkforcePoolProviderOidc
	}
	r.IssuerUri = dcl.FlattenString(m["issuerUri"])
	r.ClientId = dcl.FlattenString(m["clientId"])

	return r
}

// flattenWorkforcePoolProviderStateEnumMap flattens the contents of WorkforcePoolProviderStateEnum from a JSON
// response object.
func flattenWorkforcePoolProviderStateEnumMap(c *Client, i interface{}, res *WorkforcePoolProvider) map[string]WorkforcePoolProviderStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkforcePoolProviderStateEnum{}
	}

	if len(a) == 0 {
		return map[string]WorkforcePoolProviderStateEnum{}
	}

	items := make(map[string]WorkforcePoolProviderStateEnum)
	for k, item := range a {
		items[k] = *flattenWorkforcePoolProviderStateEnum(item.(interface{}))
	}

	return items
}

// flattenWorkforcePoolProviderStateEnumSlice flattens the contents of WorkforcePoolProviderStateEnum from a JSON
// response object.
func flattenWorkforcePoolProviderStateEnumSlice(c *Client, i interface{}, res *WorkforcePoolProvider) []WorkforcePoolProviderStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkforcePoolProviderStateEnum{}
	}

	if len(a) == 0 {
		return []WorkforcePoolProviderStateEnum{}
	}

	items := make([]WorkforcePoolProviderStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkforcePoolProviderStateEnum(item.(interface{})))
	}

	return items
}

// flattenWorkforcePoolProviderStateEnum asserts that an interface is a string, and returns a
// pointer to a *WorkforcePoolProviderStateEnum with the same value as that string.
func flattenWorkforcePoolProviderStateEnum(i interface{}) *WorkforcePoolProviderStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return WorkforcePoolProviderStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *WorkforcePoolProvider) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalWorkforcePoolProvider(b, c, r)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Location == nil && ncr.Location == nil {
			c.Config.Logger.Info("Both Location fields null - considering equal.")
		} else if nr.Location == nil || ncr.Location == nil {
			c.Config.Logger.Info("Only one Location field is null - considering unequal.")
			return false
		} else if *nr.Location != *ncr.Location {
			return false
		}
		if nr.WorkforcePool == nil && ncr.WorkforcePool == nil {
			c.Config.Logger.Info("Both WorkforcePool fields null - considering equal.")
		} else if nr.WorkforcePool == nil || ncr.WorkforcePool == nil {
			c.Config.Logger.Info("Only one WorkforcePool field is null - considering unequal.")
			return false
		} else if *nr.WorkforcePool != *ncr.WorkforcePool {
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

type workforcePoolProviderDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         workforcePoolProviderApiOperation
}

func convertFieldDiffsToWorkforcePoolProviderDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]workforcePoolProviderDiff, error) {
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
	var diffs []workforcePoolProviderDiff
	// For each operation name, create a workforcePoolProviderDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := workforcePoolProviderDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToWorkforcePoolProviderApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToWorkforcePoolProviderApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (workforcePoolProviderApiOperation, error) {
	switch opName {

	case "updateWorkforcePoolProviderUpdateWorkforcePoolProviderOperation":
		return &updateWorkforcePoolProviderUpdateWorkforcePoolProviderOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractWorkforcePoolProviderFields(r *WorkforcePoolProvider) error {
	vSaml := r.Saml
	if vSaml == nil {
		// note: explicitly not the empty object.
		vSaml = &WorkforcePoolProviderSaml{}
	}
	if err := extractWorkforcePoolProviderSamlFields(r, vSaml); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSaml) {
		r.Saml = vSaml
	}
	vOidc := r.Oidc
	if vOidc == nil {
		// note: explicitly not the empty object.
		vOidc = &WorkforcePoolProviderOidc{}
	}
	if err := extractWorkforcePoolProviderOidcFields(r, vOidc); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vOidc) {
		r.Oidc = vOidc
	}
	return nil
}
func extractWorkforcePoolProviderSamlFields(r *WorkforcePoolProvider, o *WorkforcePoolProviderSaml) error {
	return nil
}
func extractWorkforcePoolProviderOidcFields(r *WorkforcePoolProvider, o *WorkforcePoolProviderOidc) error {
	return nil
}

func postReadExtractWorkforcePoolProviderFields(r *WorkforcePoolProvider) error {
	vSaml := r.Saml
	if vSaml == nil {
		// note: explicitly not the empty object.
		vSaml = &WorkforcePoolProviderSaml{}
	}
	if err := postReadExtractWorkforcePoolProviderSamlFields(r, vSaml); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSaml) {
		r.Saml = vSaml
	}
	vOidc := r.Oidc
	if vOidc == nil {
		// note: explicitly not the empty object.
		vOidc = &WorkforcePoolProviderOidc{}
	}
	if err := postReadExtractWorkforcePoolProviderOidcFields(r, vOidc); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vOidc) {
		r.Oidc = vOidc
	}
	return nil
}
func postReadExtractWorkforcePoolProviderSamlFields(r *WorkforcePoolProvider, o *WorkforcePoolProviderSaml) error {
	return nil
}
func postReadExtractWorkforcePoolProviderOidcFields(r *WorkforcePoolProvider, o *WorkforcePoolProviderOidc) error {
	return nil
}
