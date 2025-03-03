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

func (r *AddressGroup) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "type"); err != nil {
		return err
	}
	if err := dcl.Required(r, "capacity"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Parent, "Parent"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	return nil
}
func (r *AddressGroup) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://networksecurity.googleapis.com/v1alpha1/", params)
}

func (r *AddressGroup) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(nr.Location),
		"parent":   dcl.ValueOrEmptyString(nr.Parent),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("{{parent}}/locations/{{location}}/addressGroups/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *AddressGroup) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(nr.Location),
		"parent":   dcl.ValueOrEmptyString(nr.Parent),
	}
	return dcl.URL("{{parent}}/locations/{{location}}/addressGroups", nr.basePath(), userBasePath, params), nil

}

func (r *AddressGroup) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(nr.Location),
		"parent":   dcl.ValueOrEmptyString(nr.Parent),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("{{parent}}/locations/{{location}}/addressGroups?addressGroupId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *AddressGroup) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(nr.Location),
		"parent":   dcl.ValueOrEmptyString(nr.Parent),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("{{parent}}/locations/{{location}}/addressGroups/{{name}}", nr.basePath(), userBasePath, params), nil
}

// addressGroupApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type addressGroupApiOperation interface {
	do(context.Context, *AddressGroup, *Client) error
}

// newUpdateAddressGroupUpdateAddressGroupRequest creates a request for an
// AddressGroup resource's UpdateAddressGroup update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateAddressGroupUpdateAddressGroupRequest(ctx context.Context, f *AddressGroup, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Items; v != nil {
		req["items"] = v
	}
	return req, nil
}

// marshalUpdateAddressGroupUpdateAddressGroupRequest converts the update into
// the final JSON request body.
func marshalUpdateAddressGroupUpdateAddressGroupRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateAddressGroupUpdateAddressGroupOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateAddressGroupUpdateAddressGroupOperation) do(ctx context.Context, r *AddressGroup, c *Client) error {
	_, err := c.GetAddressGroup(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateAddressGroup")
	if err != nil {
		return err
	}

	req, err := newUpdateAddressGroupUpdateAddressGroupRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateAddressGroupUpdateAddressGroupRequest(c, req)
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

func (c *Client) listAddressGroupRaw(ctx context.Context, r *AddressGroup, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != AddressGroupMaxPage {
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

type listAddressGroupOperation struct {
	AddressGroups []map[string]interface{} `json:"addressGroups"`
	Token         string                   `json:"nextPageToken"`
}

func (c *Client) listAddressGroup(ctx context.Context, r *AddressGroup, pageToken string, pageSize int32) ([]*AddressGroup, string, error) {
	b, err := c.listAddressGroupRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listAddressGroupOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*AddressGroup
	for _, v := range m.AddressGroups {
		res, err := unmarshalMapAddressGroup(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Location = r.Location
		res.Parent = r.Parent
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllAddressGroup(ctx context.Context, f func(*AddressGroup) bool, resources []*AddressGroup) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteAddressGroup(ctx, res)
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

type deleteAddressGroupOperation struct{}

func (op *deleteAddressGroupOperation) do(ctx context.Context, r *AddressGroup, c *Client) error {
	r, err := c.GetAddressGroup(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "AddressGroup not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetAddressGroup checking for existence. error: %v", err)
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
		_, err := c.GetAddressGroup(ctx, r)
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
type createAddressGroupOperation struct {
	response map[string]interface{}
}

func (op *createAddressGroupOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createAddressGroupOperation) do(ctx context.Context, r *AddressGroup, c *Client) error {
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

	if _, err := c.GetAddressGroup(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getAddressGroupRaw(ctx context.Context, r *AddressGroup) ([]byte, error) {

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

func (c *Client) addressGroupDiffsForRawDesired(ctx context.Context, rawDesired *AddressGroup, opts ...dcl.ApplyOption) (initial, desired *AddressGroup, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *AddressGroup
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*AddressGroup); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected AddressGroup, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetAddressGroup(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a AddressGroup resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve AddressGroup resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that AddressGroup resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeAddressGroupDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for AddressGroup: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for AddressGroup: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractAddressGroupFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeAddressGroupInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for AddressGroup: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeAddressGroupDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for AddressGroup: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffAddressGroup(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeAddressGroupInitialState(rawInitial, rawDesired *AddressGroup) (*AddressGroup, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeAddressGroupDesiredState(rawDesired, rawInitial *AddressGroup, opts ...dcl.ApplyOption) (*AddressGroup, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &AddressGroup{}
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
	if dcl.IsZeroValue(rawDesired.Type) || (dcl.IsEmptyValueIndirect(rawDesired.Type) && dcl.IsEmptyValueIndirect(rawInitial.Type)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Type = rawInitial.Type
	} else {
		canonicalDesired.Type = rawDesired.Type
	}
	if dcl.StringArrayCanonicalize(rawDesired.Items, rawInitial.Items) {
		canonicalDesired.Items = rawInitial.Items
	} else {
		canonicalDesired.Items = rawDesired.Items
	}
	if dcl.IsZeroValue(rawDesired.Capacity) || (dcl.IsEmptyValueIndirect(rawDesired.Capacity) && dcl.IsEmptyValueIndirect(rawInitial.Capacity)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Capacity = rawInitial.Capacity
	} else {
		canonicalDesired.Capacity = rawDesired.Capacity
	}
	if dcl.NameToSelfLink(rawDesired.Parent, rawInitial.Parent) {
		canonicalDesired.Parent = rawInitial.Parent
	} else {
		canonicalDesired.Parent = rawDesired.Parent
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}
	return canonicalDesired, nil
}

func canonicalizeAddressGroupNewState(c *Client, rawNew, rawDesired *AddressGroup) (*AddressGroup, error) {

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

	if dcl.IsEmptyValueIndirect(rawNew.Type) && dcl.IsEmptyValueIndirect(rawDesired.Type) {
		rawNew.Type = rawDesired.Type
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Items) && dcl.IsEmptyValueIndirect(rawDesired.Items) {
		rawNew.Items = rawDesired.Items
	} else {
		if dcl.StringArrayCanonicalize(rawDesired.Items, rawNew.Items) {
			rawNew.Items = rawDesired.Items
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Capacity) && dcl.IsEmptyValueIndirect(rawDesired.Capacity) {
		rawNew.Capacity = rawDesired.Capacity
	} else {
	}

	rawNew.Parent = rawDesired.Parent

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffAddressGroup(c *Client, desired, actual *AddressGroup, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateAddressGroupUpdateAddressGroupOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Items, actual.Items, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateAddressGroupUpdateAddressGroupOperation")}, fn.AddNest("Items")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Capacity, actual.Capacity, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Capacity")); len(ds) != 0 || err != nil {
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

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *AddressGroup) urlNormalized() *AddressGroup {
	normalized := dcl.Copy(*r).(AddressGroup)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Parent = r.Parent
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *AddressGroup) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateAddressGroup" {
		fields := map[string]interface{}{
			"location": dcl.ValueOrEmptyString(nr.Location),
			"parent":   dcl.ValueOrEmptyString(nr.Parent),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("{{parent}}/locations/{{location}}/addressGroups/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the AddressGroup resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *AddressGroup) marshal(c *Client) ([]byte, error) {
	m, err := expandAddressGroup(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling AddressGroup: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalAddressGroup decodes JSON responses into the AddressGroup resource schema.
func unmarshalAddressGroup(b []byte, c *Client, res *AddressGroup) (*AddressGroup, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapAddressGroup(m, c, res)
}

func unmarshalMapAddressGroup(m map[string]interface{}, c *Client, res *AddressGroup) (*AddressGroup, error) {

	flattened := flattenAddressGroup(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandAddressGroup expands AddressGroup into a JSON request object.
func expandAddressGroup(c *Client, f *AddressGroup) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("%s/locations/%s/addressGroups/%s", f.Name, f.Parent, dcl.SelfLinkToName(f.Location), dcl.SelfLinkToName(f.Name)); err != nil {
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
	if v := f.Items; v != nil {
		m["items"] = v
	}
	if v := f.Capacity; dcl.ValueShouldBeSent(v) {
		m["capacity"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Parent into parent: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["parent"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}

	return m, nil
}

// flattenAddressGroup flattens AddressGroup from a JSON request object into the
// AddressGroup type.
func flattenAddressGroup(c *Client, i interface{}, res *AddressGroup) *AddressGroup {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &AddressGroup{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.Type = flattenAddressGroupTypeEnum(m["type"])
	resultRes.Items = dcl.FlattenStringSlice(m["items"])
	resultRes.Capacity = dcl.FlattenInteger(m["capacity"])
	resultRes.Parent = dcl.FlattenString(m["parent"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// flattenAddressGroupTypeEnumMap flattens the contents of AddressGroupTypeEnum from a JSON
// response object.
func flattenAddressGroupTypeEnumMap(c *Client, i interface{}, res *AddressGroup) map[string]AddressGroupTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AddressGroupTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]AddressGroupTypeEnum{}
	}

	items := make(map[string]AddressGroupTypeEnum)
	for k, item := range a {
		items[k] = *flattenAddressGroupTypeEnum(item.(interface{}))
	}

	return items
}

// flattenAddressGroupTypeEnumSlice flattens the contents of AddressGroupTypeEnum from a JSON
// response object.
func flattenAddressGroupTypeEnumSlice(c *Client, i interface{}, res *AddressGroup) []AddressGroupTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AddressGroupTypeEnum{}
	}

	if len(a) == 0 {
		return []AddressGroupTypeEnum{}
	}

	items := make([]AddressGroupTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAddressGroupTypeEnum(item.(interface{})))
	}

	return items
}

// flattenAddressGroupTypeEnum asserts that an interface is a string, and returns a
// pointer to a *AddressGroupTypeEnum with the same value as that string.
func flattenAddressGroupTypeEnum(i interface{}) *AddressGroupTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return AddressGroupTypeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *AddressGroup) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalAddressGroup(b, c, r)
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

type addressGroupDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         addressGroupApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToAddressGroupDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]addressGroupDiff, error) {
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
	var diffs []addressGroupDiff
	// For each operation name, create a addressGroupDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := addressGroupDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToAddressGroupApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToAddressGroupApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (addressGroupApiOperation, error) {
	switch opName {

	case "updateAddressGroupUpdateAddressGroupOperation":
		return &updateAddressGroupUpdateAddressGroupOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractAddressGroupFields(r *AddressGroup) error {
	return nil
}

func postReadExtractAddressGroupFields(r *AddressGroup) error {
	return nil
}
