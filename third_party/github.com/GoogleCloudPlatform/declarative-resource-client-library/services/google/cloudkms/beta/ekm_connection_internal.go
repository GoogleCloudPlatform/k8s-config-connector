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

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *EkmConnection) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "serviceResolvers"); err != nil {
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
func (r *EkmConnectionServiceResolvers) validate() error {
	if err := dcl.Required(r, "serviceDirectoryService"); err != nil {
		return err
	}
	if err := dcl.Required(r, "hostname"); err != nil {
		return err
	}
	if err := dcl.Required(r, "serverCertificates"); err != nil {
		return err
	}
	return nil
}
func (r *EkmConnectionServiceResolversServerCertificates) validate() error {
	if err := dcl.Required(r, "rawDer"); err != nil {
		return err
	}
	return nil
}
func (r *EkmConnection) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://cloudkms.googleapis.com/v1/", params)
}

func (r *EkmConnection) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/ekmConnections/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *EkmConnection) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/ekmConnections", nr.basePath(), userBasePath, params), nil

}

func (r *EkmConnection) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/ekmConnections?ekmConnectionId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *EkmConnection) SetPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"project":  *nr.Project,
		"location": *nr.Location,
		"name":     *nr.Name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/ekmConnections/{{name}}:setIamPolicy", nr.basePath(), userBasePath, fields)
}

func (r *EkmConnection) SetPolicyVerb() string {
	return "POST"
}

func (r *EkmConnection) getPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"project":  *nr.Project,
		"location": *nr.Location,
		"name":     *nr.Name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/ekmConnections/{{name}}:getIamPolicy", nr.basePath(), userBasePath, fields)
}

func (r *EkmConnection) IAMPolicyVersion() int {
	return 3
}

// ekmConnectionApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type ekmConnectionApiOperation interface {
	do(context.Context, *EkmConnection, *Client) error
}

// newUpdateEkmConnectionUpdateEkmConnectionRequest creates a request for an
// EkmConnection resource's UpdateEkmConnection update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateEkmConnectionUpdateEkmConnectionRequest(ctx context.Context, f *EkmConnection, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v, err := dcl.DeriveField("projects/%s/locations/%s/ekmConnections/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Location), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v, err := expandEkmConnectionServiceResolversSlice(c, f.ServiceResolvers, res); err != nil {
		return nil, fmt.Errorf("error expanding ServiceResolvers into serviceResolvers: %w", err)
	} else if v != nil {
		req["serviceResolvers"] = v
	}
	b, err := c.getEkmConnectionRaw(ctx, f)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	rawEtag, err := dcl.GetMapEntry(
		m,
		[]string{"etag"},
	)
	if err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "Failed to fetch from JSON Path: %v", err)
	} else {
		req["etag"] = rawEtag.(string)
	}
	return req, nil
}

// marshalUpdateEkmConnectionUpdateEkmConnectionRequest converts the update into
// the final JSON request body.
func marshalUpdateEkmConnectionUpdateEkmConnectionRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateEkmConnectionUpdateEkmConnectionOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateEkmConnectionUpdateEkmConnectionOperation) do(ctx context.Context, r *EkmConnection, c *Client) error {
	_, err := c.GetEkmConnection(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateEkmConnection")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateEkmConnectionUpdateEkmConnectionRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateEkmConnectionUpdateEkmConnectionRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listEkmConnectionRaw(ctx context.Context, r *EkmConnection, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != EkmConnectionMaxPage {
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

type listEkmConnectionOperation struct {
	EkmConnections []map[string]interface{} `json:"ekmConnections"`
	Token          string                   `json:"nextPageToken"`
}

func (c *Client) listEkmConnection(ctx context.Context, r *EkmConnection, pageToken string, pageSize int32) ([]*EkmConnection, string, error) {
	b, err := c.listEkmConnectionRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listEkmConnectionOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*EkmConnection
	for _, v := range m.EkmConnections {
		res, err := unmarshalMapEkmConnection(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createEkmConnectionOperation struct {
	response map[string]interface{}
}

func (op *createEkmConnectionOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createEkmConnectionOperation) do(ctx context.Context, r *EkmConnection, c *Client) error {
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

	if _, err := c.GetEkmConnection(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getEkmConnectionRaw(ctx context.Context, r *EkmConnection) ([]byte, error) {

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

func (c *Client) ekmConnectionDiffsForRawDesired(ctx context.Context, rawDesired *EkmConnection, opts ...dcl.ApplyOption) (initial, desired *EkmConnection, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *EkmConnection
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*EkmConnection); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected EkmConnection, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetEkmConnection(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a EkmConnection resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve EkmConnection resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that EkmConnection resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeEkmConnectionDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for EkmConnection: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for EkmConnection: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractEkmConnectionFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeEkmConnectionInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for EkmConnection: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeEkmConnectionDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for EkmConnection: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffEkmConnection(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeEkmConnectionInitialState(rawInitial, rawDesired *EkmConnection) (*EkmConnection, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeEkmConnectionDesiredState(rawDesired, rawInitial *EkmConnection, opts ...dcl.ApplyOption) (*EkmConnection, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &EkmConnection{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	canonicalDesired.ServiceResolvers = canonicalizeEkmConnectionServiceResolversSlice(rawDesired.ServiceResolvers, rawInitial.ServiceResolvers, opts...)
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

func canonicalizeEkmConnectionNewState(c *Client, rawNew, rawDesired *EkmConnection) (*EkmConnection, error) {

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

	if dcl.IsEmptyValueIndirect(rawNew.ServiceResolvers) && dcl.IsEmptyValueIndirect(rawDesired.ServiceResolvers) {
		rawNew.ServiceResolvers = rawDesired.ServiceResolvers
	} else {
		rawNew.ServiceResolvers = canonicalizeNewEkmConnectionServiceResolversSlice(c, rawDesired.ServiceResolvers, rawNew.ServiceResolvers)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Etag) && dcl.IsEmptyValueIndirect(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
		if dcl.StringCanonicalize(rawDesired.Etag, rawNew.Etag) {
			rawNew.Etag = rawDesired.Etag
		}
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeEkmConnectionServiceResolvers(des, initial *EkmConnectionServiceResolvers, opts ...dcl.ApplyOption) *EkmConnectionServiceResolvers {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EkmConnectionServiceResolvers{}

	if dcl.IsZeroValue(des.ServiceDirectoryService) || (dcl.IsEmptyValueIndirect(des.ServiceDirectoryService) && dcl.IsEmptyValueIndirect(initial.ServiceDirectoryService)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ServiceDirectoryService = initial.ServiceDirectoryService
	} else {
		cDes.ServiceDirectoryService = des.ServiceDirectoryService
	}
	if dcl.StringCanonicalize(des.EndpointFilter, initial.EndpointFilter) || dcl.IsZeroValue(des.EndpointFilter) {
		cDes.EndpointFilter = initial.EndpointFilter
	} else {
		cDes.EndpointFilter = des.EndpointFilter
	}
	if dcl.StringCanonicalize(des.Hostname, initial.Hostname) || dcl.IsZeroValue(des.Hostname) {
		cDes.Hostname = initial.Hostname
	} else {
		cDes.Hostname = des.Hostname
	}
	cDes.ServerCertificates = canonicalizeEkmConnectionServiceResolversServerCertificatesSlice(des.ServerCertificates, initial.ServerCertificates, opts...)

	return cDes
}

func canonicalizeEkmConnectionServiceResolversSlice(des, initial []EkmConnectionServiceResolvers, opts ...dcl.ApplyOption) []EkmConnectionServiceResolvers {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]EkmConnectionServiceResolvers, 0, len(des))
		for _, d := range des {
			cd := canonicalizeEkmConnectionServiceResolvers(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]EkmConnectionServiceResolvers, 0, len(des))
	for i, d := range des {
		cd := canonicalizeEkmConnectionServiceResolvers(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewEkmConnectionServiceResolvers(c *Client, des, nw *EkmConnectionServiceResolvers) *EkmConnectionServiceResolvers {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for EkmConnectionServiceResolvers while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.EndpointFilter, nw.EndpointFilter) {
		nw.EndpointFilter = des.EndpointFilter
	}
	if dcl.StringCanonicalize(des.Hostname, nw.Hostname) {
		nw.Hostname = des.Hostname
	}
	nw.ServerCertificates = canonicalizeNewEkmConnectionServiceResolversServerCertificatesSlice(c, des.ServerCertificates, nw.ServerCertificates)

	return nw
}

func canonicalizeNewEkmConnectionServiceResolversSet(c *Client, des, nw []EkmConnectionServiceResolvers) []EkmConnectionServiceResolvers {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []EkmConnectionServiceResolvers
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareEkmConnectionServiceResolversNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewEkmConnectionServiceResolvers(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewEkmConnectionServiceResolversSlice(c *Client, des, nw []EkmConnectionServiceResolvers) []EkmConnectionServiceResolvers {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EkmConnectionServiceResolvers
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEkmConnectionServiceResolvers(c, &d, &n))
	}

	return items
}

func canonicalizeEkmConnectionServiceResolversServerCertificates(des, initial *EkmConnectionServiceResolversServerCertificates, opts ...dcl.ApplyOption) *EkmConnectionServiceResolversServerCertificates {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &EkmConnectionServiceResolversServerCertificates{}

	if dcl.StringCanonicalize(des.RawDer, initial.RawDer) || dcl.IsZeroValue(des.RawDer) {
		cDes.RawDer = initial.RawDer
	} else {
		cDes.RawDer = des.RawDer
	}

	return cDes
}

func canonicalizeEkmConnectionServiceResolversServerCertificatesSlice(des, initial []EkmConnectionServiceResolversServerCertificates, opts ...dcl.ApplyOption) []EkmConnectionServiceResolversServerCertificates {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]EkmConnectionServiceResolversServerCertificates, 0, len(des))
		for _, d := range des {
			cd := canonicalizeEkmConnectionServiceResolversServerCertificates(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]EkmConnectionServiceResolversServerCertificates, 0, len(des))
	for i, d := range des {
		cd := canonicalizeEkmConnectionServiceResolversServerCertificates(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewEkmConnectionServiceResolversServerCertificates(c *Client, des, nw *EkmConnectionServiceResolversServerCertificates) *EkmConnectionServiceResolversServerCertificates {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for EkmConnectionServiceResolversServerCertificates while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.RawDer, nw.RawDer) {
		nw.RawDer = des.RawDer
	}
	if dcl.BoolCanonicalize(des.Parsed, nw.Parsed) {
		nw.Parsed = des.Parsed
	}
	if dcl.StringCanonicalize(des.Issuer, nw.Issuer) {
		nw.Issuer = des.Issuer
	}
	if dcl.StringCanonicalize(des.Subject, nw.Subject) {
		nw.Subject = des.Subject
	}
	if dcl.StringArrayCanonicalize(des.SubjectAlternativeDnsNames, nw.SubjectAlternativeDnsNames) {
		nw.SubjectAlternativeDnsNames = des.SubjectAlternativeDnsNames
	}
	if dcl.StringCanonicalize(des.SerialNumber, nw.SerialNumber) {
		nw.SerialNumber = des.SerialNumber
	}
	if dcl.StringCanonicalize(des.Sha256Fingerprint, nw.Sha256Fingerprint) {
		nw.Sha256Fingerprint = des.Sha256Fingerprint
	}

	return nw
}

func canonicalizeNewEkmConnectionServiceResolversServerCertificatesSet(c *Client, des, nw []EkmConnectionServiceResolversServerCertificates) []EkmConnectionServiceResolversServerCertificates {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []EkmConnectionServiceResolversServerCertificates
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareEkmConnectionServiceResolversServerCertificatesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewEkmConnectionServiceResolversServerCertificates(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewEkmConnectionServiceResolversServerCertificatesSlice(c *Client, des, nw []EkmConnectionServiceResolversServerCertificates) []EkmConnectionServiceResolversServerCertificates {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []EkmConnectionServiceResolversServerCertificates
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewEkmConnectionServiceResolversServerCertificates(c, &d, &n))
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
func diffEkmConnection(c *Client, desired, actual *EkmConnection, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateEkmConnectionUpdateEkmConnectionOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.ServiceResolvers, actual.ServiceResolvers, dcl.DiffInfo{ObjectFunction: compareEkmConnectionServiceResolversNewStyle, EmptyObject: EmptyEkmConnectionServiceResolvers, OperationSelector: dcl.TriggersOperation("updateEkmConnectionUpdateEkmConnectionOperation")}, fn.AddNest("ServiceResolvers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Etag, actual.Etag, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Etag")); len(ds) != 0 || err != nil {
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
func compareEkmConnectionServiceResolversNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EkmConnectionServiceResolvers)
	if !ok {
		desiredNotPointer, ok := d.(EkmConnectionServiceResolvers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EkmConnectionServiceResolvers or *EkmConnectionServiceResolvers", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EkmConnectionServiceResolvers)
	if !ok {
		actualNotPointer, ok := a.(EkmConnectionServiceResolvers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EkmConnectionServiceResolvers", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ServiceDirectoryService, actual.ServiceDirectoryService, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateEkmConnectionUpdateEkmConnectionOperation")}, fn.AddNest("ServiceDirectoryService")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EndpointFilter, actual.EndpointFilter, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateEkmConnectionUpdateEkmConnectionOperation")}, fn.AddNest("EndpointFilter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Hostname, actual.Hostname, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateEkmConnectionUpdateEkmConnectionOperation")}, fn.AddNest("Hostname")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServerCertificates, actual.ServerCertificates, dcl.DiffInfo{ObjectFunction: compareEkmConnectionServiceResolversServerCertificatesNewStyle, EmptyObject: EmptyEkmConnectionServiceResolversServerCertificates, OperationSelector: dcl.TriggersOperation("updateEkmConnectionUpdateEkmConnectionOperation")}, fn.AddNest("ServerCertificates")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareEkmConnectionServiceResolversServerCertificatesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*EkmConnectionServiceResolversServerCertificates)
	if !ok {
		desiredNotPointer, ok := d.(EkmConnectionServiceResolversServerCertificates)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EkmConnectionServiceResolversServerCertificates or *EkmConnectionServiceResolversServerCertificates", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*EkmConnectionServiceResolversServerCertificates)
	if !ok {
		actualNotPointer, ok := a.(EkmConnectionServiceResolversServerCertificates)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a EkmConnectionServiceResolversServerCertificates", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RawDer, actual.RawDer, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateEkmConnectionUpdateEkmConnectionOperation")}, fn.AddNest("RawDer")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Parsed, actual.Parsed, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Parsed")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Issuer, actual.Issuer, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Issuer")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Subject, actual.Subject, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Subject")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SubjectAlternativeDnsNames, actual.SubjectAlternativeDnsNames, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SubjectAlternativeDnsNames")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NotBeforeTime, actual.NotBeforeTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NotBeforeTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NotAfterTime, actual.NotAfterTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NotAfterTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SerialNumber, actual.SerialNumber, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SerialNumber")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Sha256Fingerprint, actual.Sha256Fingerprint, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Sha256Fingerprint")); len(ds) != 0 || err != nil {
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
func (r *EkmConnection) urlNormalized() *EkmConnection {
	normalized := dcl.Copy(*r).(EkmConnection)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *EkmConnection) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateEkmConnection" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/ekmConnections/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the EkmConnection resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *EkmConnection) marshal(c *Client) ([]byte, error) {
	m, err := expandEkmConnection(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling EkmConnection: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalEkmConnection decodes JSON responses into the EkmConnection resource schema.
func unmarshalEkmConnection(b []byte, c *Client, res *EkmConnection) (*EkmConnection, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapEkmConnection(m, c, res)
}

func unmarshalMapEkmConnection(m map[string]interface{}, c *Client, res *EkmConnection) (*EkmConnection, error) {

	flattened := flattenEkmConnection(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandEkmConnection expands EkmConnection into a JSON request object.
func expandEkmConnection(c *Client, f *EkmConnection) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("projects/%s/locations/%s/ekmConnections/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Location), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandEkmConnectionServiceResolversSlice(c, f.ServiceResolvers, res); err != nil {
		return nil, fmt.Errorf("error expanding ServiceResolvers into serviceResolvers: %w", err)
	} else if v != nil {
		m["serviceResolvers"] = v
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

// flattenEkmConnection flattens EkmConnection from a JSON request object into the
// EkmConnection type.
func flattenEkmConnection(c *Client, i interface{}, res *EkmConnection) *EkmConnection {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &EkmConnection{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.ServiceResolvers = flattenEkmConnectionServiceResolversSlice(c, m["serviceResolvers"], res)
	resultRes.Etag = dcl.FlattenString(m["etag"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// expandEkmConnectionServiceResolversMap expands the contents of EkmConnectionServiceResolvers into a JSON
// request object.
func expandEkmConnectionServiceResolversMap(c *Client, f map[string]EkmConnectionServiceResolvers, res *EkmConnection) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEkmConnectionServiceResolvers(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEkmConnectionServiceResolversSlice expands the contents of EkmConnectionServiceResolvers into a JSON
// request object.
func expandEkmConnectionServiceResolversSlice(c *Client, f []EkmConnectionServiceResolvers, res *EkmConnection) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEkmConnectionServiceResolvers(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEkmConnectionServiceResolversMap flattens the contents of EkmConnectionServiceResolvers from a JSON
// response object.
func flattenEkmConnectionServiceResolversMap(c *Client, i interface{}, res *EkmConnection) map[string]EkmConnectionServiceResolvers {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EkmConnectionServiceResolvers{}
	}

	if len(a) == 0 {
		return map[string]EkmConnectionServiceResolvers{}
	}

	items := make(map[string]EkmConnectionServiceResolvers)
	for k, item := range a {
		items[k] = *flattenEkmConnectionServiceResolvers(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenEkmConnectionServiceResolversSlice flattens the contents of EkmConnectionServiceResolvers from a JSON
// response object.
func flattenEkmConnectionServiceResolversSlice(c *Client, i interface{}, res *EkmConnection) []EkmConnectionServiceResolvers {
	a, ok := i.([]interface{})
	if !ok {
		return []EkmConnectionServiceResolvers{}
	}

	if len(a) == 0 {
		return []EkmConnectionServiceResolvers{}
	}

	items := make([]EkmConnectionServiceResolvers, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEkmConnectionServiceResolvers(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandEkmConnectionServiceResolvers expands an instance of EkmConnectionServiceResolvers into a JSON
// request object.
func expandEkmConnectionServiceResolvers(c *Client, f *EkmConnectionServiceResolvers, res *EkmConnection) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ServiceDirectoryService; !dcl.IsEmptyValueIndirect(v) {
		m["serviceDirectoryService"] = v
	}
	if v := f.EndpointFilter; !dcl.IsEmptyValueIndirect(v) {
		m["endpointFilter"] = v
	}
	if v := f.Hostname; !dcl.IsEmptyValueIndirect(v) {
		m["hostname"] = v
	}
	if v, err := expandEkmConnectionServiceResolversServerCertificatesSlice(c, f.ServerCertificates, res); err != nil {
		return nil, fmt.Errorf("error expanding ServerCertificates into serverCertificates: %w", err)
	} else if v != nil {
		m["serverCertificates"] = v
	}

	return m, nil
}

// flattenEkmConnectionServiceResolvers flattens an instance of EkmConnectionServiceResolvers from a JSON
// response object.
func flattenEkmConnectionServiceResolvers(c *Client, i interface{}, res *EkmConnection) *EkmConnectionServiceResolvers {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EkmConnectionServiceResolvers{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEkmConnectionServiceResolvers
	}
	r.ServiceDirectoryService = dcl.FlattenString(m["serviceDirectoryService"])
	r.EndpointFilter = dcl.FlattenString(m["endpointFilter"])
	r.Hostname = dcl.FlattenString(m["hostname"])
	r.ServerCertificates = flattenEkmConnectionServiceResolversServerCertificatesSlice(c, m["serverCertificates"], res)

	return r
}

// expandEkmConnectionServiceResolversServerCertificatesMap expands the contents of EkmConnectionServiceResolversServerCertificates into a JSON
// request object.
func expandEkmConnectionServiceResolversServerCertificatesMap(c *Client, f map[string]EkmConnectionServiceResolversServerCertificates, res *EkmConnection) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandEkmConnectionServiceResolversServerCertificates(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandEkmConnectionServiceResolversServerCertificatesSlice expands the contents of EkmConnectionServiceResolversServerCertificates into a JSON
// request object.
func expandEkmConnectionServiceResolversServerCertificatesSlice(c *Client, f []EkmConnectionServiceResolversServerCertificates, res *EkmConnection) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandEkmConnectionServiceResolversServerCertificates(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenEkmConnectionServiceResolversServerCertificatesMap flattens the contents of EkmConnectionServiceResolversServerCertificates from a JSON
// response object.
func flattenEkmConnectionServiceResolversServerCertificatesMap(c *Client, i interface{}, res *EkmConnection) map[string]EkmConnectionServiceResolversServerCertificates {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EkmConnectionServiceResolversServerCertificates{}
	}

	if len(a) == 0 {
		return map[string]EkmConnectionServiceResolversServerCertificates{}
	}

	items := make(map[string]EkmConnectionServiceResolversServerCertificates)
	for k, item := range a {
		items[k] = *flattenEkmConnectionServiceResolversServerCertificates(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenEkmConnectionServiceResolversServerCertificatesSlice flattens the contents of EkmConnectionServiceResolversServerCertificates from a JSON
// response object.
func flattenEkmConnectionServiceResolversServerCertificatesSlice(c *Client, i interface{}, res *EkmConnection) []EkmConnectionServiceResolversServerCertificates {
	a, ok := i.([]interface{})
	if !ok {
		return []EkmConnectionServiceResolversServerCertificates{}
	}

	if len(a) == 0 {
		return []EkmConnectionServiceResolversServerCertificates{}
	}

	items := make([]EkmConnectionServiceResolversServerCertificates, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEkmConnectionServiceResolversServerCertificates(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandEkmConnectionServiceResolversServerCertificates expands an instance of EkmConnectionServiceResolversServerCertificates into a JSON
// request object.
func expandEkmConnectionServiceResolversServerCertificates(c *Client, f *EkmConnectionServiceResolversServerCertificates, res *EkmConnection) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RawDer; !dcl.IsEmptyValueIndirect(v) {
		m["rawDer"] = v
	}

	return m, nil
}

// flattenEkmConnectionServiceResolversServerCertificates flattens an instance of EkmConnectionServiceResolversServerCertificates from a JSON
// response object.
func flattenEkmConnectionServiceResolversServerCertificates(c *Client, i interface{}, res *EkmConnection) *EkmConnectionServiceResolversServerCertificates {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &EkmConnectionServiceResolversServerCertificates{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyEkmConnectionServiceResolversServerCertificates
	}
	r.RawDer = dcl.FlattenString(m["rawDer"])
	r.Parsed = dcl.FlattenBool(m["parsed"])
	r.Issuer = dcl.FlattenString(m["issuer"])
	r.Subject = dcl.FlattenString(m["subject"])
	r.SubjectAlternativeDnsNames = dcl.FlattenStringSlice(m["subjectAlternativeDnsNames"])
	r.NotBeforeTime = dcl.FlattenString(m["notBeforeTime"])
	r.NotAfterTime = dcl.FlattenString(m["notAfterTime"])
	r.SerialNumber = dcl.FlattenString(m["serialNumber"])
	r.Sha256Fingerprint = dcl.FlattenString(m["sha256Fingerprint"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *EkmConnection) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalEkmConnection(b, c, r)
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

type ekmConnectionDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         ekmConnectionApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToEkmConnectionDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]ekmConnectionDiff, error) {
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
	var diffs []ekmConnectionDiff
	// For each operation name, create a ekmConnectionDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := ekmConnectionDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToEkmConnectionApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToEkmConnectionApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (ekmConnectionApiOperation, error) {
	switch opName {

	case "updateEkmConnectionUpdateEkmConnectionOperation":
		return &updateEkmConnectionUpdateEkmConnectionOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractEkmConnectionFields(r *EkmConnection) error {
	return nil
}
func extractEkmConnectionServiceResolversFields(r *EkmConnection, o *EkmConnectionServiceResolvers) error {
	return nil
}
func extractEkmConnectionServiceResolversServerCertificatesFields(r *EkmConnection, o *EkmConnectionServiceResolversServerCertificates) error {
	return nil
}

func postReadExtractEkmConnectionFields(r *EkmConnection) error {
	return nil
}
func postReadExtractEkmConnectionServiceResolversFields(r *EkmConnection, o *EkmConnectionServiceResolvers) error {
	return nil
}
func postReadExtractEkmConnectionServiceResolversServerCertificatesFields(r *EkmConnection, o *EkmConnectionServiceResolversServerCertificates) error {
	return nil
}
