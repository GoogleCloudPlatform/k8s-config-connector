// Copyright 2023 Google LLC. All Rights Reserved.
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

func (r *VmwareEngineNetwork) validate() error {

	if err := dcl.RequiredParameter(r.Name, "Name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "type"); err != nil {
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
func (r *VmwareEngineNetworkVPCNetworks) validate() error {
	return nil
}
func (r *VmwareEngineNetwork) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://vmwareengine.googleapis.com/v1/", params)
}

func (r *VmwareEngineNetwork) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/vmwareEngineNetworks/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *VmwareEngineNetwork) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/vmwareEngineNetworks", nr.basePath(), userBasePath, params), nil

}

func (r *VmwareEngineNetwork) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/vmwareEngineNetworks?vmwareEngineNetworkId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *VmwareEngineNetwork) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/vmwareEngineNetworks/{{name}}", nr.basePath(), userBasePath, params), nil
}

// vmwareEngineNetworkApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type vmwareEngineNetworkApiOperation interface {
	do(context.Context, *VmwareEngineNetwork, *Client) error
}

// newUpdateVmwareEngineNetworkUpdateVmwareEngineNetworkRequest creates a request for an
// VmwareEngineNetwork resource's UpdateVmwareEngineNetwork update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateVmwareEngineNetworkUpdateVmwareEngineNetworkRequest(ctx context.Context, f *VmwareEngineNetwork, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		req["type"] = v
	}
	b, err := c.getVmwareEngineNetworkRaw(ctx, f)
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
	req["name"] = fmt.Sprintf("projects/%s/locations/%s/vmwareEngineNetworks/%s", *f.Project, *f.Location, *f.Name)

	return req, nil
}

// marshalUpdateVmwareEngineNetworkUpdateVmwareEngineNetworkRequest converts the update into
// the final JSON request body.
func marshalUpdateVmwareEngineNetworkUpdateVmwareEngineNetworkRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateVmwareEngineNetworkUpdateVmwareEngineNetworkOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateVmwareEngineNetworkUpdateVmwareEngineNetworkOperation) do(ctx context.Context, r *VmwareEngineNetwork, c *Client) error {
	_, err := c.GetVmwareEngineNetwork(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateVmwareEngineNetwork")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateVmwareEngineNetworkUpdateVmwareEngineNetworkRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateVmwareEngineNetworkUpdateVmwareEngineNetworkRequest(c, req)
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

func (c *Client) listVmwareEngineNetworkRaw(ctx context.Context, r *VmwareEngineNetwork, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != VmwareEngineNetworkMaxPage {
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

type listVmwareEngineNetworkOperation struct {
	VmwareEngineNetworks []map[string]interface{} `json:"vmwareEngineNetworks"`
	Token                string                   `json:"nextPageToken"`
}

func (c *Client) listVmwareEngineNetwork(ctx context.Context, r *VmwareEngineNetwork, pageToken string, pageSize int32) ([]*VmwareEngineNetwork, string, error) {
	b, err := c.listVmwareEngineNetworkRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listVmwareEngineNetworkOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*VmwareEngineNetwork
	for _, v := range m.VmwareEngineNetworks {
		res, err := unmarshalMapVmwareEngineNetwork(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllVmwareEngineNetwork(ctx context.Context, f func(*VmwareEngineNetwork) bool, resources []*VmwareEngineNetwork) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteVmwareEngineNetwork(ctx, res)
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

type deleteVmwareEngineNetworkOperation struct{}

func (op *deleteVmwareEngineNetworkOperation) do(ctx context.Context, r *VmwareEngineNetwork, c *Client) error {
	r, err := c.GetVmwareEngineNetwork(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "VmwareEngineNetwork not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetVmwareEngineNetwork checking for existence. error: %v", err)
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
		_, err := c.GetVmwareEngineNetwork(ctx, r)
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
type createVmwareEngineNetworkOperation struct {
	response map[string]interface{}
}

func (op *createVmwareEngineNetworkOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createVmwareEngineNetworkOperation) do(ctx context.Context, r *VmwareEngineNetwork, c *Client) error {
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

	if _, err := c.GetVmwareEngineNetwork(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getVmwareEngineNetworkRaw(ctx context.Context, r *VmwareEngineNetwork) ([]byte, error) {

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

func (c *Client) vmwareEngineNetworkDiffsForRawDesired(ctx context.Context, rawDesired *VmwareEngineNetwork, opts ...dcl.ApplyOption) (initial, desired *VmwareEngineNetwork, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *VmwareEngineNetwork
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*VmwareEngineNetwork); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected VmwareEngineNetwork, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetVmwareEngineNetwork(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a VmwareEngineNetwork resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve VmwareEngineNetwork resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that VmwareEngineNetwork resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeVmwareEngineNetworkDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for VmwareEngineNetwork: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for VmwareEngineNetwork: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractVmwareEngineNetworkFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeVmwareEngineNetworkInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for VmwareEngineNetwork: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeVmwareEngineNetworkDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for VmwareEngineNetwork: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffVmwareEngineNetwork(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeVmwareEngineNetworkInitialState(rawInitial, rawDesired *VmwareEngineNetwork) (*VmwareEngineNetwork, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeVmwareEngineNetworkDesiredState(rawDesired, rawInitial *VmwareEngineNetwork, opts ...dcl.ApplyOption) (*VmwareEngineNetwork, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &VmwareEngineNetwork{}
	if dcl.NameToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
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

func canonicalizeVmwareEngineNetworkNewState(c *Client, rawNew, rawDesired *VmwareEngineNetwork) (*VmwareEngineNetwork, error) {

	rawNew.Name = rawDesired.Name

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

	if dcl.IsEmptyValueIndirect(rawNew.VPCNetworks) && dcl.IsEmptyValueIndirect(rawDesired.VPCNetworks) {
		rawNew.VPCNetworks = rawDesired.VPCNetworks
	} else {
		rawNew.VPCNetworks = canonicalizeNewVmwareEngineNetworkVPCNetworksSlice(c, rawDesired.VPCNetworks, rawNew.VPCNetworks)
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Type) && dcl.IsEmptyValueIndirect(rawDesired.Type) {
		rawNew.Type = rawDesired.Type
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Uid) && dcl.IsEmptyValueIndirect(rawDesired.Uid) {
		rawNew.Uid = rawDesired.Uid
	} else {
		if dcl.StringCanonicalize(rawDesired.Uid, rawNew.Uid) {
			rawNew.Uid = rawDesired.Uid
		}
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

func canonicalizeVmwareEngineNetworkVPCNetworks(des, initial *VmwareEngineNetworkVPCNetworks, opts ...dcl.ApplyOption) *VmwareEngineNetworkVPCNetworks {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &VmwareEngineNetworkVPCNetworks{}

	return cDes
}

func canonicalizeVmwareEngineNetworkVPCNetworksSlice(des, initial []VmwareEngineNetworkVPCNetworks, opts ...dcl.ApplyOption) []VmwareEngineNetworkVPCNetworks {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]VmwareEngineNetworkVPCNetworks, 0, len(des))
		for _, d := range des {
			cd := canonicalizeVmwareEngineNetworkVPCNetworks(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]VmwareEngineNetworkVPCNetworks, 0, len(des))
	for i, d := range des {
		cd := canonicalizeVmwareEngineNetworkVPCNetworks(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewVmwareEngineNetworkVPCNetworks(c *Client, des, nw *VmwareEngineNetworkVPCNetworks) *VmwareEngineNetworkVPCNetworks {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for VmwareEngineNetworkVPCNetworks while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewVmwareEngineNetworkVPCNetworksSet(c *Client, des, nw []VmwareEngineNetworkVPCNetworks) []VmwareEngineNetworkVPCNetworks {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []VmwareEngineNetworkVPCNetworks
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareVmwareEngineNetworkVPCNetworksNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewVmwareEngineNetworkVPCNetworks(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewVmwareEngineNetworkVPCNetworksSlice(c *Client, des, nw []VmwareEngineNetworkVPCNetworks) []VmwareEngineNetworkVPCNetworks {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []VmwareEngineNetworkVPCNetworks
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewVmwareEngineNetworkVPCNetworks(c, &d, &n))
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
func diffVmwareEngineNetwork(c *Client, desired, actual *VmwareEngineNetwork, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateVmwareEngineNetworkUpdateVmwareEngineNetworkOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VPCNetworks, actual.VPCNetworks, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareVmwareEngineNetworkVPCNetworksNewStyle, EmptyObject: EmptyVmwareEngineNetworkVPCNetworks, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("VpcNetworks")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateVmwareEngineNetworkUpdateVmwareEngineNetworkOperation")}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Uid, actual.Uid, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Uid")); len(ds) != 0 || err != nil {
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
func compareVmwareEngineNetworkVPCNetworksNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*VmwareEngineNetworkVPCNetworks)
	if !ok {
		desiredNotPointer, ok := d.(VmwareEngineNetworkVPCNetworks)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VmwareEngineNetworkVPCNetworks or *VmwareEngineNetworkVPCNetworks", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*VmwareEngineNetworkVPCNetworks)
	if !ok {
		actualNotPointer, ok := a.(VmwareEngineNetworkVPCNetworks)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a VmwareEngineNetworkVPCNetworks", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Network, actual.Network, dcl.DiffInfo{OutputOnly: true, Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Network")); len(ds) != 0 || err != nil {
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
func (r *VmwareEngineNetwork) urlNormalized() *VmwareEngineNetwork {
	normalized := dcl.Copy(*r).(VmwareEngineNetwork)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Uid = dcl.SelfLinkToName(r.Uid)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *VmwareEngineNetwork) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateVmwareEngineNetwork" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/vmwareEngineNetworks/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the VmwareEngineNetwork resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *VmwareEngineNetwork) marshal(c *Client) ([]byte, error) {
	m, err := expandVmwareEngineNetwork(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling VmwareEngineNetwork: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalVmwareEngineNetwork decodes JSON responses into the VmwareEngineNetwork resource schema.
func unmarshalVmwareEngineNetwork(b []byte, c *Client, res *VmwareEngineNetwork) (*VmwareEngineNetwork, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapVmwareEngineNetwork(m, c, res)
}

func unmarshalMapVmwareEngineNetwork(m map[string]interface{}, c *Client, res *VmwareEngineNetwork) (*VmwareEngineNetwork, error) {

	flattened := flattenVmwareEngineNetwork(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandVmwareEngineNetwork expands VmwareEngineNetwork into a JSON request object.
func expandVmwareEngineNetwork(c *Client, f *VmwareEngineNetwork) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v := f.Type; dcl.ValueShouldBeSent(v) {
		m["type"] = v
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

// flattenVmwareEngineNetwork flattens VmwareEngineNetwork from a JSON request object into the
// VmwareEngineNetwork type.
func flattenVmwareEngineNetwork(c *Client, i interface{}, res *VmwareEngineNetwork) *VmwareEngineNetwork {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &VmwareEngineNetwork{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.VPCNetworks = flattenVmwareEngineNetworkVPCNetworksSlice(c, m["vpcNetworks"], res)
	resultRes.State = flattenVmwareEngineNetworkStateEnum(m["state"])
	resultRes.Type = flattenVmwareEngineNetworkTypeEnum(m["type"])
	resultRes.Uid = dcl.FlattenString(m["uid"])
	resultRes.Etag = dcl.FlattenString(m["etag"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// expandVmwareEngineNetworkVPCNetworksMap expands the contents of VmwareEngineNetworkVPCNetworks into a JSON
// request object.
func expandVmwareEngineNetworkVPCNetworksMap(c *Client, f map[string]VmwareEngineNetworkVPCNetworks, res *VmwareEngineNetwork) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandVmwareEngineNetworkVPCNetworks(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandVmwareEngineNetworkVPCNetworksSlice expands the contents of VmwareEngineNetworkVPCNetworks into a JSON
// request object.
func expandVmwareEngineNetworkVPCNetworksSlice(c *Client, f []VmwareEngineNetworkVPCNetworks, res *VmwareEngineNetwork) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandVmwareEngineNetworkVPCNetworks(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenVmwareEngineNetworkVPCNetworksMap flattens the contents of VmwareEngineNetworkVPCNetworks from a JSON
// response object.
func flattenVmwareEngineNetworkVPCNetworksMap(c *Client, i interface{}, res *VmwareEngineNetwork) map[string]VmwareEngineNetworkVPCNetworks {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VmwareEngineNetworkVPCNetworks{}
	}

	if len(a) == 0 {
		return map[string]VmwareEngineNetworkVPCNetworks{}
	}

	items := make(map[string]VmwareEngineNetworkVPCNetworks)
	for k, item := range a {
		items[k] = *flattenVmwareEngineNetworkVPCNetworks(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenVmwareEngineNetworkVPCNetworksSlice flattens the contents of VmwareEngineNetworkVPCNetworks from a JSON
// response object.
func flattenVmwareEngineNetworkVPCNetworksSlice(c *Client, i interface{}, res *VmwareEngineNetwork) []VmwareEngineNetworkVPCNetworks {
	a, ok := i.([]interface{})
	if !ok {
		return []VmwareEngineNetworkVPCNetworks{}
	}

	if len(a) == 0 {
		return []VmwareEngineNetworkVPCNetworks{}
	}

	items := make([]VmwareEngineNetworkVPCNetworks, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVmwareEngineNetworkVPCNetworks(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandVmwareEngineNetworkVPCNetworks expands an instance of VmwareEngineNetworkVPCNetworks into a JSON
// request object.
func expandVmwareEngineNetworkVPCNetworks(c *Client, f *VmwareEngineNetworkVPCNetworks, res *VmwareEngineNetwork) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenVmwareEngineNetworkVPCNetworks flattens an instance of VmwareEngineNetworkVPCNetworks from a JSON
// response object.
func flattenVmwareEngineNetworkVPCNetworks(c *Client, i interface{}, res *VmwareEngineNetwork) *VmwareEngineNetworkVPCNetworks {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &VmwareEngineNetworkVPCNetworks{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyVmwareEngineNetworkVPCNetworks
	}
	r.Type = flattenVmwareEngineNetworkVPCNetworksTypeEnum(m["type"])
	r.Network = dcl.FlattenString(m["network"])

	return r
}

// flattenVmwareEngineNetworkVPCNetworksTypeEnumMap flattens the contents of VmwareEngineNetworkVPCNetworksTypeEnum from a JSON
// response object.
func flattenVmwareEngineNetworkVPCNetworksTypeEnumMap(c *Client, i interface{}, res *VmwareEngineNetwork) map[string]VmwareEngineNetworkVPCNetworksTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VmwareEngineNetworkVPCNetworksTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]VmwareEngineNetworkVPCNetworksTypeEnum{}
	}

	items := make(map[string]VmwareEngineNetworkVPCNetworksTypeEnum)
	for k, item := range a {
		items[k] = *flattenVmwareEngineNetworkVPCNetworksTypeEnum(item.(interface{}))
	}

	return items
}

// flattenVmwareEngineNetworkVPCNetworksTypeEnumSlice flattens the contents of VmwareEngineNetworkVPCNetworksTypeEnum from a JSON
// response object.
func flattenVmwareEngineNetworkVPCNetworksTypeEnumSlice(c *Client, i interface{}, res *VmwareEngineNetwork) []VmwareEngineNetworkVPCNetworksTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []VmwareEngineNetworkVPCNetworksTypeEnum{}
	}

	if len(a) == 0 {
		return []VmwareEngineNetworkVPCNetworksTypeEnum{}
	}

	items := make([]VmwareEngineNetworkVPCNetworksTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVmwareEngineNetworkVPCNetworksTypeEnum(item.(interface{})))
	}

	return items
}

// flattenVmwareEngineNetworkVPCNetworksTypeEnum asserts that an interface is a string, and returns a
// pointer to a *VmwareEngineNetworkVPCNetworksTypeEnum with the same value as that string.
func flattenVmwareEngineNetworkVPCNetworksTypeEnum(i interface{}) *VmwareEngineNetworkVPCNetworksTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return VmwareEngineNetworkVPCNetworksTypeEnumRef(s)
}

// flattenVmwareEngineNetworkStateEnumMap flattens the contents of VmwareEngineNetworkStateEnum from a JSON
// response object.
func flattenVmwareEngineNetworkStateEnumMap(c *Client, i interface{}, res *VmwareEngineNetwork) map[string]VmwareEngineNetworkStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VmwareEngineNetworkStateEnum{}
	}

	if len(a) == 0 {
		return map[string]VmwareEngineNetworkStateEnum{}
	}

	items := make(map[string]VmwareEngineNetworkStateEnum)
	for k, item := range a {
		items[k] = *flattenVmwareEngineNetworkStateEnum(item.(interface{}))
	}

	return items
}

// flattenVmwareEngineNetworkStateEnumSlice flattens the contents of VmwareEngineNetworkStateEnum from a JSON
// response object.
func flattenVmwareEngineNetworkStateEnumSlice(c *Client, i interface{}, res *VmwareEngineNetwork) []VmwareEngineNetworkStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []VmwareEngineNetworkStateEnum{}
	}

	if len(a) == 0 {
		return []VmwareEngineNetworkStateEnum{}
	}

	items := make([]VmwareEngineNetworkStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVmwareEngineNetworkStateEnum(item.(interface{})))
	}

	return items
}

// flattenVmwareEngineNetworkStateEnum asserts that an interface is a string, and returns a
// pointer to a *VmwareEngineNetworkStateEnum with the same value as that string.
func flattenVmwareEngineNetworkStateEnum(i interface{}) *VmwareEngineNetworkStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return VmwareEngineNetworkStateEnumRef(s)
}

// flattenVmwareEngineNetworkTypeEnumMap flattens the contents of VmwareEngineNetworkTypeEnum from a JSON
// response object.
func flattenVmwareEngineNetworkTypeEnumMap(c *Client, i interface{}, res *VmwareEngineNetwork) map[string]VmwareEngineNetworkTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]VmwareEngineNetworkTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]VmwareEngineNetworkTypeEnum{}
	}

	items := make(map[string]VmwareEngineNetworkTypeEnum)
	for k, item := range a {
		items[k] = *flattenVmwareEngineNetworkTypeEnum(item.(interface{}))
	}

	return items
}

// flattenVmwareEngineNetworkTypeEnumSlice flattens the contents of VmwareEngineNetworkTypeEnum from a JSON
// response object.
func flattenVmwareEngineNetworkTypeEnumSlice(c *Client, i interface{}, res *VmwareEngineNetwork) []VmwareEngineNetworkTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []VmwareEngineNetworkTypeEnum{}
	}

	if len(a) == 0 {
		return []VmwareEngineNetworkTypeEnum{}
	}

	items := make([]VmwareEngineNetworkTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenVmwareEngineNetworkTypeEnum(item.(interface{})))
	}

	return items
}

// flattenVmwareEngineNetworkTypeEnum asserts that an interface is a string, and returns a
// pointer to a *VmwareEngineNetworkTypeEnum with the same value as that string.
func flattenVmwareEngineNetworkTypeEnum(i interface{}) *VmwareEngineNetworkTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return VmwareEngineNetworkTypeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *VmwareEngineNetwork) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalVmwareEngineNetwork(b, c, r)
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

type vmwareEngineNetworkDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         vmwareEngineNetworkApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToVmwareEngineNetworkDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]vmwareEngineNetworkDiff, error) {
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
	var diffs []vmwareEngineNetworkDiff
	// For each operation name, create a vmwareEngineNetworkDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := vmwareEngineNetworkDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToVmwareEngineNetworkApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToVmwareEngineNetworkApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (vmwareEngineNetworkApiOperation, error) {
	switch opName {

	case "updateVmwareEngineNetworkUpdateVmwareEngineNetworkOperation":
		return &updateVmwareEngineNetworkUpdateVmwareEngineNetworkOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractVmwareEngineNetworkFields(r *VmwareEngineNetwork) error {
	return nil
}
func extractVmwareEngineNetworkVPCNetworksFields(r *VmwareEngineNetwork, o *VmwareEngineNetworkVPCNetworks) error {
	return nil
}

func postReadExtractVmwareEngineNetworkFields(r *VmwareEngineNetwork) error {
	return nil
}
func postReadExtractVmwareEngineNetworkVPCNetworksFields(r *VmwareEngineNetwork, o *VmwareEngineNetworkVPCNetworks) error {
	return nil
}
