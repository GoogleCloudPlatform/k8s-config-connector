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

func (r *EnvironmentGroup) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "hostnames"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.ApigeeOrganization, "ApigeeOrganization"); err != nil {
		return err
	}
	return nil
}
func (r *EnvironmentGroup) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://apigee.googleapis.com/v1/", params)
}

func (r *EnvironmentGroup) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"apigeeOrganization": dcl.ValueOrEmptyString(nr.ApigeeOrganization),
		"name":               dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("organizations/{{apigeeOrganization}}/envgroups/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *EnvironmentGroup) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"apigeeOrganization": dcl.ValueOrEmptyString(nr.ApigeeOrganization),
	}
	return dcl.URL("organizations/{{apigeeOrganization}}/envgroups", nr.basePath(), userBasePath, params), nil

}

func (r *EnvironmentGroup) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"apigeeOrganization": dcl.ValueOrEmptyString(nr.ApigeeOrganization),
	}
	return dcl.URL("organizations/{{apigeeOrganization}}/envgroups", nr.basePath(), userBasePath, params), nil

}

func (r *EnvironmentGroup) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"apigeeOrganization": dcl.ValueOrEmptyString(nr.ApigeeOrganization),
		"name":               dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("organizations/{{apigeeOrganization}}/envgroups/{{name}}", nr.basePath(), userBasePath, params), nil
}

// environmentGroupApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type environmentGroupApiOperation interface {
	do(context.Context, *EnvironmentGroup, *Client) error
}

// newUpdateEnvironmentGroupPatchEnvironmentGroupRequest creates a request for an
// EnvironmentGroup resource's PatchEnvironmentGroup update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateEnvironmentGroupPatchEnvironmentGroupRequest(ctx context.Context, f *EnvironmentGroup, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.Hostnames; v != nil {
		req["hostnames"] = v
	}
	return req, nil
}

// marshalUpdateEnvironmentGroupPatchEnvironmentGroupRequest converts the update into
// the final JSON request body.
func marshalUpdateEnvironmentGroupPatchEnvironmentGroupRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateEnvironmentGroupPatchEnvironmentGroupOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateEnvironmentGroupPatchEnvironmentGroupOperation) do(ctx context.Context, r *EnvironmentGroup, c *Client) error {
	_, err := c.GetEnvironmentGroup(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "PatchEnvironmentGroup")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateEnvironmentGroupPatchEnvironmentGroupRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateEnvironmentGroupPatchEnvironmentGroupRequest(c, req)
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

func (c *Client) listEnvironmentGroupRaw(ctx context.Context, r *EnvironmentGroup, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != EnvironmentGroupMaxPage {
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

type listEnvironmentGroupOperation struct {
	EnvironmentGroups []map[string]interface{} `json:"environmentGroups"`
	Token             string                   `json:"nextPageToken"`
}

func (c *Client) listEnvironmentGroup(ctx context.Context, r *EnvironmentGroup, pageToken string, pageSize int32) ([]*EnvironmentGroup, string, error) {
	b, err := c.listEnvironmentGroupRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listEnvironmentGroupOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*EnvironmentGroup
	for _, v := range m.EnvironmentGroups {
		res, err := unmarshalMapEnvironmentGroup(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.ApigeeOrganization = r.ApigeeOrganization
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllEnvironmentGroup(ctx context.Context, f func(*EnvironmentGroup) bool, resources []*EnvironmentGroup) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteEnvironmentGroup(ctx, res)
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

type deleteEnvironmentGroupOperation struct{}

func (op *deleteEnvironmentGroupOperation) do(ctx context.Context, r *EnvironmentGroup, c *Client) error {
	r, err := c.GetEnvironmentGroup(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "EnvironmentGroup not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetEnvironmentGroup checking for existence. error: %v", err)
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
		_, err := c.GetEnvironmentGroup(ctx, r)
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
type createEnvironmentGroupOperation struct {
	response map[string]interface{}
}

func (op *createEnvironmentGroupOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createEnvironmentGroupOperation) do(ctx context.Context, r *EnvironmentGroup, c *Client) error {
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

	if _, err := c.GetEnvironmentGroup(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getEnvironmentGroupRaw(ctx context.Context, r *EnvironmentGroup) ([]byte, error) {

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

func (c *Client) environmentGroupDiffsForRawDesired(ctx context.Context, rawDesired *EnvironmentGroup, opts ...dcl.ApplyOption) (initial, desired *EnvironmentGroup, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *EnvironmentGroup
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*EnvironmentGroup); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected EnvironmentGroup, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetEnvironmentGroup(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a EnvironmentGroup resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve EnvironmentGroup resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that EnvironmentGroup resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeEnvironmentGroupDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for EnvironmentGroup: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for EnvironmentGroup: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractEnvironmentGroupFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeEnvironmentGroupInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for EnvironmentGroup: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeEnvironmentGroupDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for EnvironmentGroup: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffEnvironmentGroup(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeEnvironmentGroupInitialState(rawInitial, rawDesired *EnvironmentGroup) (*EnvironmentGroup, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeEnvironmentGroupDesiredState(rawDesired, rawInitial *EnvironmentGroup, opts ...dcl.ApplyOption) (*EnvironmentGroup, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &EnvironmentGroup{}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringArrayCanonicalize(rawDesired.Hostnames, rawInitial.Hostnames) {
		canonicalDesired.Hostnames = rawInitial.Hostnames
	} else {
		canonicalDesired.Hostnames = rawDesired.Hostnames
	}
	if dcl.NameToSelfLink(rawDesired.ApigeeOrganization, rawInitial.ApigeeOrganization) {
		canonicalDesired.ApigeeOrganization = rawInitial.ApigeeOrganization
	} else {
		canonicalDesired.ApigeeOrganization = rawDesired.ApigeeOrganization
	}

	return canonicalDesired, nil
}

func canonicalizeEnvironmentGroupNewState(c *Client, rawNew, rawDesired *EnvironmentGroup) (*EnvironmentGroup, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Hostnames) && dcl.IsEmptyValueIndirect(rawDesired.Hostnames) {
		rawNew.Hostnames = rawDesired.Hostnames
	} else {
		if dcl.StringArrayCanonicalize(rawDesired.Hostnames, rawNew.Hostnames) {
			rawNew.Hostnames = rawDesired.Hostnames
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreatedAt) && dcl.IsEmptyValueIndirect(rawDesired.CreatedAt) {
		rawNew.CreatedAt = rawDesired.CreatedAt
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.LastModifiedAt) && dcl.IsEmptyValueIndirect(rawDesired.LastModifiedAt) {
		rawNew.LastModifiedAt = rawDesired.LastModifiedAt
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	rawNew.ApigeeOrganization = rawDesired.ApigeeOrganization

	return rawNew, nil
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffEnvironmentGroup(c *Client, desired, actual *EnvironmentGroup, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Hostnames, actual.Hostnames, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateEnvironmentGroupPatchEnvironmentGroupOperation")}, fn.AddNest("Hostnames")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreatedAt, actual.CreatedAt, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreatedAt")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LastModifiedAt, actual.LastModifiedAt, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LastModifiedAt")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.ApigeeOrganization, actual.ApigeeOrganization, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ApigeeOrganization")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *EnvironmentGroup) urlNormalized() *EnvironmentGroup {
	normalized := dcl.Copy(*r).(EnvironmentGroup)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.ApigeeOrganization = dcl.SelfLinkToName(r.ApigeeOrganization)
	return &normalized
}

func (r *EnvironmentGroup) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "PatchEnvironmentGroup" {
		fields := map[string]interface{}{
			"apigeeOrganization": dcl.ValueOrEmptyString(nr.ApigeeOrganization),
			"name":               dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("organizations/{{apigeeOrganization}}/envgroups/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the EnvironmentGroup resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *EnvironmentGroup) marshal(c *Client) ([]byte, error) {
	m, err := expandEnvironmentGroup(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling EnvironmentGroup: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalEnvironmentGroup decodes JSON responses into the EnvironmentGroup resource schema.
func unmarshalEnvironmentGroup(b []byte, c *Client, res *EnvironmentGroup) (*EnvironmentGroup, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapEnvironmentGroup(m, c, res)
}

func unmarshalMapEnvironmentGroup(m map[string]interface{}, c *Client, res *EnvironmentGroup) (*EnvironmentGroup, error) {

	flattened := flattenEnvironmentGroup(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandEnvironmentGroup expands EnvironmentGroup into a JSON request object.
func expandEnvironmentGroup(c *Client, f *EnvironmentGroup) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v := f.Name; dcl.ValueShouldBeSent(v) {
		m["name"] = v
	}
	if v := f.Hostnames; v != nil {
		m["hostnames"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding ApigeeOrganization into apigeeOrganization: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["apigeeOrganization"] = v
	}

	return m, nil
}

// flattenEnvironmentGroup flattens EnvironmentGroup from a JSON request object into the
// EnvironmentGroup type.
func flattenEnvironmentGroup(c *Client, i interface{}, res *EnvironmentGroup) *EnvironmentGroup {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &EnvironmentGroup{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.Hostnames = dcl.FlattenStringSlice(m["hostnames"])
	resultRes.CreatedAt = dcl.FlattenInteger(m["createdAt"])
	resultRes.LastModifiedAt = dcl.FlattenInteger(m["lastModifiedAt"])
	resultRes.State = flattenEnvironmentGroupStateEnum(m["state"])
	resultRes.ApigeeOrganization = dcl.FlattenString(m["apigeeOrganization"])

	return resultRes
}

// flattenEnvironmentGroupStateEnumMap flattens the contents of EnvironmentGroupStateEnum from a JSON
// response object.
func flattenEnvironmentGroupStateEnumMap(c *Client, i interface{}, res *EnvironmentGroup) map[string]EnvironmentGroupStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EnvironmentGroupStateEnum{}
	}

	if len(a) == 0 {
		return map[string]EnvironmentGroupStateEnum{}
	}

	items := make(map[string]EnvironmentGroupStateEnum)
	for k, item := range a {
		items[k] = *flattenEnvironmentGroupStateEnum(item.(interface{}))
	}

	return items
}

// flattenEnvironmentGroupStateEnumSlice flattens the contents of EnvironmentGroupStateEnum from a JSON
// response object.
func flattenEnvironmentGroupStateEnumSlice(c *Client, i interface{}, res *EnvironmentGroup) []EnvironmentGroupStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []EnvironmentGroupStateEnum{}
	}

	if len(a) == 0 {
		return []EnvironmentGroupStateEnum{}
	}

	items := make([]EnvironmentGroupStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEnvironmentGroupStateEnum(item.(interface{})))
	}

	return items
}

// flattenEnvironmentGroupStateEnum asserts that an interface is a string, and returns a
// pointer to a *EnvironmentGroupStateEnum with the same value as that string.
func flattenEnvironmentGroupStateEnum(i interface{}) *EnvironmentGroupStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return EnvironmentGroupStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *EnvironmentGroup) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalEnvironmentGroup(b, c, r)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.ApigeeOrganization == nil && ncr.ApigeeOrganization == nil {
			c.Config.Logger.Info("Both ApigeeOrganization fields null - considering equal.")
		} else if nr.ApigeeOrganization == nil || ncr.ApigeeOrganization == nil {
			c.Config.Logger.Info("Only one ApigeeOrganization field is null - considering unequal.")
			return false
		} else if *nr.ApigeeOrganization != *ncr.ApigeeOrganization {
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

type environmentGroupDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         environmentGroupApiOperation
}

func convertFieldDiffsToEnvironmentGroupDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]environmentGroupDiff, error) {
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
	var diffs []environmentGroupDiff
	// For each operation name, create a environmentGroupDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := environmentGroupDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToEnvironmentGroupApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToEnvironmentGroupApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (environmentGroupApiOperation, error) {
	switch opName {

	case "updateEnvironmentGroupPatchEnvironmentGroupOperation":
		return &updateEnvironmentGroupPatchEnvironmentGroupOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractEnvironmentGroupFields(r *EnvironmentGroup) error {
	return nil
}

func postReadExtractEnvironmentGroupFields(r *EnvironmentGroup) error {
	return nil
}
