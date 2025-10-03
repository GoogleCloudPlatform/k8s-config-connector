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
package iam

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *Role) validate() error {

	if !dcl.IsEmptyValueIndirect(r.LocalizedValues) {
		if err := r.LocalizedValues.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *RoleLocalizedValues) validate() error {
	return nil
}
func (r *Role) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://iam.googleapis.com/v1/", params)
}

func (r *Role) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"parent": dcl.ValueOrEmptyString(nr.Parent),
		"name":   dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("{{parent}}/roles/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Role) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"parent": dcl.ValueOrEmptyString(nr.Parent),
	}
	return dcl.URL("{{parent}}/roles", nr.basePath(), userBasePath, params), nil

}

func (r *Role) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"parent": dcl.ValueOrEmptyString(nr.Parent),
	}
	return dcl.URL("{{parent}}/roles", nr.basePath(), userBasePath, params), nil

}

func (r *Role) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"parent": dcl.ValueOrEmptyString(nr.Parent),
		"name":   dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("{{parent}}/roles/{{name}}", nr.basePath(), userBasePath, params), nil
}

// roleApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type roleApiOperation interface {
	do(context.Context, *Role, *Client) error
}

// newUpdateRoleUpdateRoleRequest creates a request for an
// Role resource's UpdateRole update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateRoleUpdateRoleRequest(ctx context.Context, f *Role, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	b, err := c.getRoleRaw(ctx, f)
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

// marshalUpdateRoleUpdateRoleRequest converts the update into
// the final JSON request body.
func marshalUpdateRoleUpdateRoleRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateRoleUpdateRoleOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateRoleUpdateRoleOperation) do(ctx context.Context, r *Role, c *Client) error {
	_, err := c.GetRole(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateRole")
	if err != nil {
		return err
	}

	req, err := newUpdateRoleUpdateRoleRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateRoleUpdateRoleRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listRoleRaw(ctx context.Context, r *Role, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != RoleMaxPage {
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

type listRoleOperation struct {
	Roles []map[string]interface{} `json:"roles"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listRole(ctx context.Context, r *Role, pageToken string, pageSize int32) ([]*Role, string, error) {
	b, err := c.listRoleRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listRoleOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Role
	for _, v := range m.Roles {
		res, err := unmarshalMapRole(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Parent = r.Parent
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllRole(ctx context.Context, f func(*Role) bool, resources []*Role) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteRole(ctx, res)
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

type deleteRoleOperation struct{}

func (op *deleteRoleOperation) do(ctx context.Context, r *Role, c *Client) error {
	r, err := c.GetRole(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Role not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetRole checking for existence. error: %v", err)
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
		return fmt.Errorf("failed to delete Role: %w", err)
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createRoleOperation struct {
	response map[string]interface{}
}

func (op *createRoleOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createRoleOperation) do(ctx context.Context, r *Role, c *Client) error {
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

	if _, err := c.GetRole(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getRoleRaw(ctx context.Context, r *Role) ([]byte, error) {

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

func (c *Client) roleDiffsForRawDesired(ctx context.Context, rawDesired *Role, opts ...dcl.ApplyOption) (initial, desired *Role, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Role
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Role); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Role, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetRole(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Role resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Role resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Role resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeRoleDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Role: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Role: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractRoleFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeRoleInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Role: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeRoleDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Role: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffRole(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeRoleInitialState(rawInitial, rawDesired *Role) (*Role, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeRoleDesiredState(rawDesired, rawInitial *Role, opts ...dcl.ApplyOption) (*Role, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.LocalizedValues = canonicalizeRoleLocalizedValues(rawDesired.LocalizedValues, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Role{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.Title, rawInitial.Title) {
		canonicalDesired.Title = rawInitial.Title
	} else {
		canonicalDesired.Title = rawDesired.Title
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	canonicalDesired.LocalizedValues = canonicalizeRoleLocalizedValues(rawDesired.LocalizedValues, rawInitial.LocalizedValues, opts...)
	if dcl.StringCanonicalize(rawDesired.LifecyclePhase, rawInitial.LifecyclePhase) {
		canonicalDesired.LifecyclePhase = rawInitial.LifecyclePhase
	} else {
		canonicalDesired.LifecyclePhase = rawDesired.LifecyclePhase
	}
	if dcl.StringCanonicalize(rawDesired.GroupName, rawInitial.GroupName) {
		canonicalDesired.GroupName = rawInitial.GroupName
	} else {
		canonicalDesired.GroupName = rawDesired.GroupName
	}
	if dcl.StringCanonicalize(rawDesired.GroupTitle, rawInitial.GroupTitle) {
		canonicalDesired.GroupTitle = rawInitial.GroupTitle
	} else {
		canonicalDesired.GroupTitle = rawDesired.GroupTitle
	}
	if dcl.StringArrayCanonicalize(rawDesired.IncludedPermissions, rawInitial.IncludedPermissions) {
		canonicalDesired.IncludedPermissions = rawInitial.IncludedPermissions
	} else {
		canonicalDesired.IncludedPermissions = rawDesired.IncludedPermissions
	}
	if dcl.IsZeroValue(rawDesired.Stage) || (dcl.IsEmptyValueIndirect(rawDesired.Stage) && dcl.IsEmptyValueIndirect(rawInitial.Stage)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Stage = rawInitial.Stage
	} else {
		canonicalDesired.Stage = rawDesired.Stage
	}
	if dcl.StringCanonicalize(rawDesired.Etag, rawInitial.Etag) {
		canonicalDesired.Etag = rawInitial.Etag
	} else {
		canonicalDesired.Etag = rawDesired.Etag
	}
	if dcl.BoolCanonicalize(rawDesired.Deleted, rawInitial.Deleted) {
		canonicalDesired.Deleted = rawInitial.Deleted
	} else {
		canonicalDesired.Deleted = rawDesired.Deleted
	}
	if dcl.StringArrayCanonicalize(rawDesired.IncludedRoles, rawInitial.IncludedRoles) {
		canonicalDesired.IncludedRoles = rawInitial.IncludedRoles
	} else {
		canonicalDesired.IncludedRoles = rawDesired.IncludedRoles
	}
	if dcl.NameToSelfLink(rawDesired.Parent, rawInitial.Parent) {
		canonicalDesired.Parent = rawInitial.Parent
	} else {
		canonicalDesired.Parent = rawDesired.Parent
	}
	return canonicalDesired, nil
}

func canonicalizeRoleNewState(c *Client, rawNew, rawDesired *Role) (*Role, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Title) && dcl.IsEmptyValueIndirect(rawDesired.Title) {
		rawNew.Title = rawDesired.Title
	} else {
		if dcl.StringCanonicalize(rawDesired.Title, rawNew.Title) {
			rawNew.Title = rawDesired.Title
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.LocalizedValues) && dcl.IsEmptyValueIndirect(rawDesired.LocalizedValues) {
		rawNew.LocalizedValues = rawDesired.LocalizedValues
	} else {
		rawNew.LocalizedValues = canonicalizeNewRoleLocalizedValues(c, rawDesired.LocalizedValues, rawNew.LocalizedValues)
	}

	if dcl.IsEmptyValueIndirect(rawNew.LifecyclePhase) && dcl.IsEmptyValueIndirect(rawDesired.LifecyclePhase) {
		rawNew.LifecyclePhase = rawDesired.LifecyclePhase
	} else {
		if dcl.StringCanonicalize(rawDesired.LifecyclePhase, rawNew.LifecyclePhase) {
			rawNew.LifecyclePhase = rawDesired.LifecyclePhase
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.GroupName) && dcl.IsEmptyValueIndirect(rawDesired.GroupName) {
		rawNew.GroupName = rawDesired.GroupName
	} else {
		if dcl.StringCanonicalize(rawDesired.GroupName, rawNew.GroupName) {
			rawNew.GroupName = rawDesired.GroupName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.GroupTitle) && dcl.IsEmptyValueIndirect(rawDesired.GroupTitle) {
		rawNew.GroupTitle = rawDesired.GroupTitle
	} else {
		if dcl.StringCanonicalize(rawDesired.GroupTitle, rawNew.GroupTitle) {
			rawNew.GroupTitle = rawDesired.GroupTitle
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.IncludedPermissions) && dcl.IsEmptyValueIndirect(rawDesired.IncludedPermissions) {
		rawNew.IncludedPermissions = rawDesired.IncludedPermissions
	} else {
		if dcl.StringArrayCanonicalize(rawDesired.IncludedPermissions, rawNew.IncludedPermissions) {
			rawNew.IncludedPermissions = rawDesired.IncludedPermissions
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Stage) && dcl.IsEmptyValueIndirect(rawDesired.Stage) {
		rawNew.Stage = rawDesired.Stage
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Etag) && dcl.IsEmptyValueIndirect(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
		if dcl.StringCanonicalize(rawDesired.Etag, rawNew.Etag) {
			rawNew.Etag = rawDesired.Etag
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Deleted) && dcl.IsEmptyValueIndirect(rawDesired.Deleted) {
		rawNew.Deleted = rawDesired.Deleted
	} else {
		if dcl.BoolCanonicalize(rawDesired.Deleted, rawNew.Deleted) {
			rawNew.Deleted = rawDesired.Deleted
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.IncludedRoles) && dcl.IsEmptyValueIndirect(rawDesired.IncludedRoles) {
		rawNew.IncludedRoles = rawDesired.IncludedRoles
	} else {
		if dcl.StringArrayCanonicalize(rawDesired.IncludedRoles, rawNew.IncludedRoles) {
			rawNew.IncludedRoles = rawDesired.IncludedRoles
		}
	}

	rawNew.Parent = rawDesired.Parent

	return rawNew, nil
}

func canonicalizeRoleLocalizedValues(des, initial *RoleLocalizedValues, opts ...dcl.ApplyOption) *RoleLocalizedValues {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &RoleLocalizedValues{}

	if dcl.StringCanonicalize(des.LocalizedTitle, initial.LocalizedTitle) || dcl.IsZeroValue(des.LocalizedTitle) {
		cDes.LocalizedTitle = initial.LocalizedTitle
	} else {
		cDes.LocalizedTitle = des.LocalizedTitle
	}
	if dcl.StringCanonicalize(des.LocalizedDescription, initial.LocalizedDescription) || dcl.IsZeroValue(des.LocalizedDescription) {
		cDes.LocalizedDescription = initial.LocalizedDescription
	} else {
		cDes.LocalizedDescription = des.LocalizedDescription
	}

	return cDes
}

func canonicalizeRoleLocalizedValuesSlice(des, initial []RoleLocalizedValues, opts ...dcl.ApplyOption) []RoleLocalizedValues {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]RoleLocalizedValues, 0, len(des))
		for _, d := range des {
			cd := canonicalizeRoleLocalizedValues(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]RoleLocalizedValues, 0, len(des))
	for i, d := range des {
		cd := canonicalizeRoleLocalizedValues(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewRoleLocalizedValues(c *Client, des, nw *RoleLocalizedValues) *RoleLocalizedValues {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for RoleLocalizedValues while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.LocalizedTitle, nw.LocalizedTitle) {
		nw.LocalizedTitle = des.LocalizedTitle
	}
	if dcl.StringCanonicalize(des.LocalizedDescription, nw.LocalizedDescription) {
		nw.LocalizedDescription = des.LocalizedDescription
	}

	return nw
}

func canonicalizeNewRoleLocalizedValuesSet(c *Client, des, nw []RoleLocalizedValues) []RoleLocalizedValues {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []RoleLocalizedValues
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareRoleLocalizedValuesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewRoleLocalizedValues(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewRoleLocalizedValuesSlice(c *Client, des, nw []RoleLocalizedValues) []RoleLocalizedValues {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []RoleLocalizedValues
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewRoleLocalizedValues(c, &d, &n))
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
func diffRole(c *Client, desired, actual *Role, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Title, actual.Title, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Title")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LocalizedValues, actual.LocalizedValues, dcl.DiffInfo{ObjectFunction: compareRoleLocalizedValuesNewStyle, EmptyObject: EmptyRoleLocalizedValues, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LocalizedValues")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LifecyclePhase, actual.LifecyclePhase, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LifecyclePhase")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupName, actual.GroupName, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GroupName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupTitle, actual.GroupTitle, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GroupTitle")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IncludedPermissions, actual.IncludedPermissions, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IncludedPermissions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Stage, actual.Stage, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Stage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Etag, actual.Etag, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Etag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Deleted, actual.Deleted, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Deleted")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IncludedRoles, actual.IncludedRoles, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IncludedRoles")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Parent, actual.Parent, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Parent")); len(ds) != 0 || err != nil {
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
func compareRoleLocalizedValuesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*RoleLocalizedValues)
	if !ok {
		desiredNotPointer, ok := d.(RoleLocalizedValues)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a RoleLocalizedValues or *RoleLocalizedValues", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*RoleLocalizedValues)
	if !ok {
		actualNotPointer, ok := a.(RoleLocalizedValues)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a RoleLocalizedValues", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.LocalizedTitle, actual.LocalizedTitle, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LocalizedTitle")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LocalizedDescription, actual.LocalizedDescription, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LocalizedDescription")); len(ds) != 0 || err != nil {
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
func (r *Role) urlNormalized() *Role {
	normalized := dcl.Copy(*r).(Role)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Title = dcl.SelfLinkToName(r.Title)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.LifecyclePhase = dcl.SelfLinkToName(r.LifecyclePhase)
	normalized.GroupName = dcl.SelfLinkToName(r.GroupName)
	normalized.GroupTitle = dcl.SelfLinkToName(r.GroupTitle)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.Parent = r.Parent
	return &normalized
}

func (r *Role) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateRole" {
		fields := map[string]interface{}{
			"parent": dcl.ValueOrEmptyString(nr.Parent),
			"name":   dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("{{parent}}/roles/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Role resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Role) marshal(c *Client) ([]byte, error) {
	m, err := expandRole(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Role: %w", err)
	}
	m = EncodeRoleCreateRequest(m)

	return json.Marshal(m)
}

// unmarshalRole decodes JSON responses into the Role resource schema.
func unmarshalRole(b []byte, c *Client, res *Role) (*Role, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapRole(m, c, res)
}

func unmarshalMapRole(m map[string]interface{}, c *Client, res *Role) (*Role, error) {

	flattened := flattenRole(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandRole expands Role into a JSON request object.
func expandRole(c *Client, f *Role) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("%s/roles/%s", f.Name, f.Parent, dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Title; dcl.ValueShouldBeSent(v) {
		m["title"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v, err := expandRoleLocalizedValues(c, f.LocalizedValues, res); err != nil {
		return nil, fmt.Errorf("error expanding LocalizedValues into localizedValues: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["localizedValues"] = v
	}
	if v := f.LifecyclePhase; dcl.ValueShouldBeSent(v) {
		m["lifecyclePhase"] = v
	}
	if v := f.GroupName; dcl.ValueShouldBeSent(v) {
		m["groupName"] = v
	}
	if v := f.GroupTitle; dcl.ValueShouldBeSent(v) {
		m["groupTitle"] = v
	}
	if v := f.IncludedPermissions; v != nil {
		m["includedPermissions"] = v
	}
	if v := f.Stage; dcl.ValueShouldBeSent(v) {
		m["stage"] = v
	}
	if v := f.Etag; dcl.ValueShouldBeSent(v) {
		m["etag"] = v
	}
	if v := f.Deleted; dcl.ValueShouldBeSent(v) {
		m["deleted"] = v
	}
	if v := f.IncludedRoles; v != nil {
		m["includedRoles"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Parent into parent: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["parent"] = v
	}

	return m, nil
}

// flattenRole flattens Role from a JSON request object into the
// Role type.
func flattenRole(c *Client, i interface{}, res *Role) *Role {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Role{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.Title = dcl.FlattenString(m["title"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.LocalizedValues = flattenRoleLocalizedValues(c, m["localizedValues"], res)
	resultRes.LifecyclePhase = dcl.FlattenString(m["lifecyclePhase"])
	resultRes.GroupName = dcl.FlattenString(m["groupName"])
	resultRes.GroupTitle = dcl.FlattenString(m["groupTitle"])
	resultRes.IncludedPermissions = dcl.FlattenStringSlice(m["includedPermissions"])
	resultRes.Stage = flattenRoleStageEnum(m["stage"])
	resultRes.Etag = dcl.FlattenString(m["etag"])
	resultRes.Deleted = dcl.FlattenBool(m["deleted"])
	resultRes.IncludedRoles = dcl.FlattenStringSlice(m["includedRoles"])
	resultRes.Parent = dcl.FlattenString(m["parent"])

	return resultRes
}

// expandRoleLocalizedValuesMap expands the contents of RoleLocalizedValues into a JSON
// request object.
func expandRoleLocalizedValuesMap(c *Client, f map[string]RoleLocalizedValues, res *Role) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandRoleLocalizedValues(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandRoleLocalizedValuesSlice expands the contents of RoleLocalizedValues into a JSON
// request object.
func expandRoleLocalizedValuesSlice(c *Client, f []RoleLocalizedValues, res *Role) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandRoleLocalizedValues(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenRoleLocalizedValuesMap flattens the contents of RoleLocalizedValues from a JSON
// response object.
func flattenRoleLocalizedValuesMap(c *Client, i interface{}, res *Role) map[string]RoleLocalizedValues {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]RoleLocalizedValues{}
	}

	if len(a) == 0 {
		return map[string]RoleLocalizedValues{}
	}

	items := make(map[string]RoleLocalizedValues)
	for k, item := range a {
		items[k] = *flattenRoleLocalizedValues(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenRoleLocalizedValuesSlice flattens the contents of RoleLocalizedValues from a JSON
// response object.
func flattenRoleLocalizedValuesSlice(c *Client, i interface{}, res *Role) []RoleLocalizedValues {
	a, ok := i.([]interface{})
	if !ok {
		return []RoleLocalizedValues{}
	}

	if len(a) == 0 {
		return []RoleLocalizedValues{}
	}

	items := make([]RoleLocalizedValues, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRoleLocalizedValues(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandRoleLocalizedValues expands an instance of RoleLocalizedValues into a JSON
// request object.
func expandRoleLocalizedValues(c *Client, f *RoleLocalizedValues, res *Role) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.LocalizedTitle; !dcl.IsEmptyValueIndirect(v) {
		m["localizedTitle"] = v
	}
	if v := f.LocalizedDescription; !dcl.IsEmptyValueIndirect(v) {
		m["localizedDescription"] = v
	}

	return m, nil
}

// flattenRoleLocalizedValues flattens an instance of RoleLocalizedValues from a JSON
// response object.
func flattenRoleLocalizedValues(c *Client, i interface{}, res *Role) *RoleLocalizedValues {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &RoleLocalizedValues{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyRoleLocalizedValues
	}
	r.LocalizedTitle = dcl.FlattenString(m["localizedTitle"])
	r.LocalizedDescription = dcl.FlattenString(m["localizedDescription"])

	return r
}

// flattenRoleStageEnumMap flattens the contents of RoleStageEnum from a JSON
// response object.
func flattenRoleStageEnumMap(c *Client, i interface{}, res *Role) map[string]RoleStageEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]RoleStageEnum{}
	}

	if len(a) == 0 {
		return map[string]RoleStageEnum{}
	}

	items := make(map[string]RoleStageEnum)
	for k, item := range a {
		items[k] = *flattenRoleStageEnum(item.(interface{}))
	}

	return items
}

// flattenRoleStageEnumSlice flattens the contents of RoleStageEnum from a JSON
// response object.
func flattenRoleStageEnumSlice(c *Client, i interface{}, res *Role) []RoleStageEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []RoleStageEnum{}
	}

	if len(a) == 0 {
		return []RoleStageEnum{}
	}

	items := make([]RoleStageEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenRoleStageEnum(item.(interface{})))
	}

	return items
}

// flattenRoleStageEnum asserts that an interface is a string, and returns a
// pointer to a *RoleStageEnum with the same value as that string.
func flattenRoleStageEnum(i interface{}) *RoleStageEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return RoleStageEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Role) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalRole(b, c, r)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Parent == nil && ncr.Parent == nil {
			c.Config.Logger.Info("Both Parent fields null - considering equal.")
		} else if nr.Parent == nil || ncr.Parent == nil {
			c.Config.Logger.Info("Only one Parent field is null - considering unequal.")
			return false
		} else if *nr.Parent != *ncr.Parent {
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

type roleDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         roleApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToRoleDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]roleDiff, error) {
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
	var diffs []roleDiff
	// For each operation name, create a roleDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := roleDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToRoleApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToRoleApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (roleApiOperation, error) {
	switch opName {

	case "updateRoleUpdateRoleOperation":
		return &updateRoleUpdateRoleOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractRoleFields(r *Role) error {
	vLocalizedValues := r.LocalizedValues
	if vLocalizedValues == nil {
		// note: explicitly not the empty object.
		vLocalizedValues = &RoleLocalizedValues{}
	}
	if err := extractRoleLocalizedValuesFields(r, vLocalizedValues); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLocalizedValues) {
		r.LocalizedValues = vLocalizedValues
	}
	return nil
}
func extractRoleLocalizedValuesFields(r *Role, o *RoleLocalizedValues) error {
	return nil
}

func postReadExtractRoleFields(r *Role) error {
	vLocalizedValues := r.LocalizedValues
	if vLocalizedValues == nil {
		// note: explicitly not the empty object.
		vLocalizedValues = &RoleLocalizedValues{}
	}
	if err := postReadExtractRoleLocalizedValuesFields(r, vLocalizedValues); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLocalizedValues) {
		r.LocalizedValues = vLocalizedValues
	}
	return nil
}
func postReadExtractRoleLocalizedValuesFields(r *Role, o *RoleLocalizedValues) error {
	return nil
}
