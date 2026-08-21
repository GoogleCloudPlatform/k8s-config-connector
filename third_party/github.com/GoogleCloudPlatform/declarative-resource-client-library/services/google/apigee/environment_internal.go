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
package apigee

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

func (r *Environment) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.ApigeeOrganization, "ApigeeOrganization"); err != nil {
		return err
	}
	return nil
}
func (r *Environment) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://apigee.googleapis.com/v1/", params)
}

func (r *Environment) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"apigeeOrganization": dcl.ValueOrEmptyString(nr.ApigeeOrganization),
		"name":               dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("organizations/{{apigeeOrganization}}/environments/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Environment) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"apigeeOrganization": dcl.ValueOrEmptyString(nr.ApigeeOrganization),
	}
	return dcl.URL("organizations/{{apigeeOrganization}}/environments", nr.basePath(), userBasePath, params), nil

}

func (r *Environment) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"apigeeOrganization": dcl.ValueOrEmptyString(nr.ApigeeOrganization),
	}
	return dcl.URL("organizations/{{apigeeOrganization}}/environments", nr.basePath(), userBasePath, params), nil

}

func (r *Environment) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"apigeeOrganization": dcl.ValueOrEmptyString(nr.ApigeeOrganization),
		"name":               dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("organizations/{{apigeeOrganization}}/environments/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Environment) SetPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"apigeeOrganization": *nr.ApigeeOrganization,
		"name":               *nr.Name,
	}
	return dcl.URL("organizations/{{apigeeOrganization}}/environments/{{name}}:setIamPolicy", nr.basePath(), userBasePath, fields)
}

func (r *Environment) SetPolicyVerb() string {
	return "POST"
}

func (r *Environment) getPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"apigeeOrganization": *nr.ApigeeOrganization,
		"name":               *nr.Name,
	}
	return dcl.URL("organizations/{{apigeeOrganization}}/environments/{{name}}:getIamPolicy", nr.basePath(), userBasePath, fields)
}

func (r *Environment) IAMPolicyVersion() int {
	return 3
}

// environmentApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type environmentApiOperation interface {
	do(context.Context, *Environment, *Client) error
}

// newUpdateEnvironmentUpdateEnvironmentRequest creates a request for an
// Environment resource's UpdateEnvironment update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateEnvironmentUpdateEnvironmentRequest(ctx context.Context, f *Environment, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v, err := dcl.ListOfKeyValuesFromMapInStruct(f.Properties, "property", "name", "value"); err != nil {
		return nil, fmt.Errorf("error expanding Properties into properties: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["properties"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	return req, nil
}

// marshalUpdateEnvironmentUpdateEnvironmentRequest converts the update into
// the final JSON request body.
func marshalUpdateEnvironmentUpdateEnvironmentRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateEnvironmentUpdateEnvironmentOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateEnvironmentUpdateEnvironmentOperation) do(ctx context.Context, r *Environment, c *Client) error {
	_, err := c.GetEnvironment(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateEnvironment")
	if err != nil {
		return err
	}

	req, err := newUpdateEnvironmentUpdateEnvironmentRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateEnvironmentUpdateEnvironmentRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) deleteAllEnvironment(ctx context.Context, f func(*Environment) bool, resources []*Environment) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteEnvironment(ctx, res)
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

type deleteEnvironmentOperation struct{}

func (op *deleteEnvironmentOperation) do(ctx context.Context, r *Environment, c *Client) error {
	r, err := c.GetEnvironment(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Environment not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetEnvironment checking for existence. error: %v", err)
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
		_, err := c.GetEnvironment(ctx, r)
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
type createEnvironmentOperation struct {
	response map[string]interface{}
}

func (op *createEnvironmentOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createEnvironmentOperation) do(ctx context.Context, r *Environment, c *Client) error {
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

	if _, err := c.GetEnvironment(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getEnvironmentRaw(ctx context.Context, r *Environment) ([]byte, error) {

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

func (c *Client) environmentDiffsForRawDesired(ctx context.Context, rawDesired *Environment, opts ...dcl.ApplyOption) (initial, desired *Environment, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Environment
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Environment); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Environment, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetEnvironment(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Environment resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Environment resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Environment resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeEnvironmentDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Environment: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Environment: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractEnvironmentFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeEnvironmentInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Environment: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeEnvironmentDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Environment: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffEnvironment(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeEnvironmentInitialState(rawInitial, rawDesired *Environment) (*Environment, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeEnvironmentDesiredState(rawDesired, rawInitial *Environment, opts ...dcl.ApplyOption) (*Environment, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &Environment{}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.IsZeroValue(rawDesired.Properties) || (dcl.IsEmptyValueIndirect(rawDesired.Properties) && dcl.IsEmptyValueIndirect(rawInitial.Properties)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Properties = rawInitial.Properties
	} else {
		canonicalDesired.Properties = rawDesired.Properties
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		canonicalDesired.DisplayName = rawInitial.DisplayName
	} else {
		canonicalDesired.DisplayName = rawDesired.DisplayName
	}
	if dcl.NameToSelfLink(rawDesired.ApigeeOrganization, rawInitial.ApigeeOrganization) {
		canonicalDesired.ApigeeOrganization = rawInitial.ApigeeOrganization
	} else {
		canonicalDesired.ApigeeOrganization = rawDesired.ApigeeOrganization
	}
	return canonicalDesired, nil
}

func canonicalizeEnvironmentNewState(c *Client, rawNew, rawDesired *Environment) (*Environment, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
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

	if dcl.IsEmptyValueIndirect(rawNew.CreatedAt) && dcl.IsEmptyValueIndirect(rawDesired.CreatedAt) {
		rawNew.CreatedAt = rawDesired.CreatedAt
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.LastModifiedAt) && dcl.IsEmptyValueIndirect(rawDesired.LastModifiedAt) {
		rawNew.LastModifiedAt = rawDesired.LastModifiedAt
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Properties) && dcl.IsEmptyValueIndirect(rawDesired.Properties) {
		rawNew.Properties = rawDesired.Properties
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
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
func diffEnvironment(c *Client, desired, actual *Environment, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateEnvironmentUpdateEnvironmentOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Properties, actual.Properties, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateEnvironmentUpdateEnvironmentOperation")}, fn.AddNest("Properties")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateEnvironmentUpdateEnvironmentOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
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

	if len(newDiffs) > 0 {
		c.Config.Logger.Infof("Diff function found diffs: %v", newDiffs)
	}
	return newDiffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Environment) urlNormalized() *Environment {
	normalized := dcl.Copy(*r).(Environment)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.ApigeeOrganization = dcl.SelfLinkToName(r.ApigeeOrganization)
	return &normalized
}

func (r *Environment) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateEnvironment" {
		fields := map[string]interface{}{
			"apigeeOrganization": dcl.ValueOrEmptyString(nr.ApigeeOrganization),
			"name":               dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("organizations/{{apigeeOrganization}}/environments/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Environment resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Environment) marshal(c *Client) ([]byte, error) {
	m, err := expandEnvironment(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Environment: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalEnvironment decodes JSON responses into the Environment resource schema.
func unmarshalEnvironment(b []byte, c *Client, res *Environment) (*Environment, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapEnvironment(m, c, res)
}

func unmarshalMapEnvironment(m map[string]interface{}, c *Client, res *Environment) (*Environment, error) {
	if v, err := dcl.MapFromListOfKeyValues(m, []string{"properties", "property"}, "name", "value"); err != nil {
		return nil, err
	} else {
		dcl.PutMapEntry(
			m,
			[]string{"properties"},
			v,
		)
	}

	flattened := flattenEnvironment(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandEnvironment expands Environment into a JSON request object.
func expandEnvironment(c *Client, f *Environment) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v := f.Name; dcl.ValueShouldBeSent(v) {
		m["name"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v, err := dcl.ListOfKeyValuesFromMapInStruct(f.Properties, "property", "name", "value"); err != nil {
		return nil, fmt.Errorf("error expanding Properties into properties: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["properties"] = v
	}
	if v := f.DisplayName; dcl.ValueShouldBeSent(v) {
		m["displayName"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding ApigeeOrganization into apigeeOrganization: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["apigeeOrganization"] = v
	}

	return m, nil
}

// flattenEnvironment flattens Environment from a JSON request object into the
// Environment type.
func flattenEnvironment(c *Client, i interface{}, res *Environment) *Environment {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Environment{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.CreatedAt = dcl.FlattenInteger(m["createdAt"])
	resultRes.LastModifiedAt = dcl.FlattenInteger(m["lastModifiedAt"])
	resultRes.Properties = dcl.FlattenKeyValuePairs(m["properties"])
	resultRes.DisplayName = dcl.FlattenString(m["displayName"])
	resultRes.State = flattenEnvironmentStateEnum(m["state"])
	resultRes.ApigeeOrganization = dcl.FlattenString(m["apigeeOrganization"])

	return resultRes
}

// flattenEnvironmentStateEnumMap flattens the contents of EnvironmentStateEnum from a JSON
// response object.
func flattenEnvironmentStateEnumMap(c *Client, i interface{}, res *Environment) map[string]EnvironmentStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]EnvironmentStateEnum{}
	}

	if len(a) == 0 {
		return map[string]EnvironmentStateEnum{}
	}

	items := make(map[string]EnvironmentStateEnum)
	for k, item := range a {
		items[k] = *flattenEnvironmentStateEnum(item.(interface{}))
	}

	return items
}

// flattenEnvironmentStateEnumSlice flattens the contents of EnvironmentStateEnum from a JSON
// response object.
func flattenEnvironmentStateEnumSlice(c *Client, i interface{}, res *Environment) []EnvironmentStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []EnvironmentStateEnum{}
	}

	if len(a) == 0 {
		return []EnvironmentStateEnum{}
	}

	items := make([]EnvironmentStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenEnvironmentStateEnum(item.(interface{})))
	}

	return items
}

// flattenEnvironmentStateEnum asserts that an interface is a string, and returns a
// pointer to a *EnvironmentStateEnum with the same value as that string.
func flattenEnvironmentStateEnum(i interface{}) *EnvironmentStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return EnvironmentStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Environment) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalEnvironment(b, c, r)
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

type environmentDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         environmentApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToEnvironmentDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]environmentDiff, error) {
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
	var diffs []environmentDiff
	// For each operation name, create a environmentDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := environmentDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToEnvironmentApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToEnvironmentApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (environmentApiOperation, error) {
	switch opName {

	case "updateEnvironmentUpdateEnvironmentOperation":
		return &updateEnvironmentUpdateEnvironmentOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractEnvironmentFields(r *Environment) error {
	return nil
}

func postReadExtractEnvironmentFields(r *Environment) error {
	return nil
}
