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

func (r *WorkforcePool) validate() error {

	if err := dcl.RequiredParameter(r.Name, "Name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "parent"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	return nil
}
func (r *WorkforcePool) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://iam.googleapis.com/v1/", params)
}

func (r *WorkforcePool) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("locations/{{location}}/workforcePools/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *WorkforcePool) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(nr.Location),
		"parent":   dcl.ValueOrEmptyString(nr.Parent),
	}
	return dcl.URL("locations/{{location}}/workforcePools?parent={{parent}}", nr.basePath(), userBasePath, params), nil

}

func (r *WorkforcePool) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("locations/{{location}}/workforcePools?workforcePoolId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *WorkforcePool) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("locations/{{location}}/workforcePools/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *WorkforcePool) SetPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"location": *nr.Location,
		"name":     *nr.Name,
	}
	return dcl.URL("locations/{{location}}/workforcePools/{{name}}:setIamPolicy", nr.basePath(), userBasePath, fields)
}

func (r *WorkforcePool) SetPolicyVerb() string {
	return "POST"
}

func (r *WorkforcePool) getPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"location": *nr.Location,
		"name":     *nr.Name,
	}
	return dcl.URL("locations/{{location}}/workforcePools/{{name}}:getIamPolicy", nr.basePath(), userBasePath, fields)
}

func (r *WorkforcePool) IAMPolicyVersion() int {
	return 3
}

// workforcePoolApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type workforcePoolApiOperation interface {
	do(context.Context, *WorkforcePool, *Client) error
}

// newUpdateWorkforcePoolUpdateWorkforcePoolRequest creates a request for an
// WorkforcePool resource's UpdateWorkforcePool update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateWorkforcePoolUpdateWorkforcePoolRequest(ctx context.Context, f *WorkforcePool, c *Client) (map[string]interface{}, error) {
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
	if v := f.SessionDuration; !dcl.IsEmptyValueIndirect(v) {
		req["sessionDuration"] = v
	}
	return req, nil
}

// marshalUpdateWorkforcePoolUpdateWorkforcePoolRequest converts the update into
// the final JSON request body.
func marshalUpdateWorkforcePoolUpdateWorkforcePoolRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateWorkforcePoolUpdateWorkforcePoolOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateWorkforcePoolUpdateWorkforcePoolOperation) do(ctx context.Context, r *WorkforcePool, c *Client) error {
	_, err := c.GetWorkforcePool(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateWorkforcePool")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateWorkforcePoolUpdateWorkforcePoolRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateWorkforcePoolUpdateWorkforcePoolRequest(c, req)
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

func (c *Client) listWorkforcePoolRaw(ctx context.Context, r *WorkforcePool, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != WorkforcePoolMaxPage {
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

type listWorkforcePoolOperation struct {
	WorkforcePools []map[string]interface{} `json:"workforcePools"`
	Token          string                   `json:"nextPageToken"`
}

func (c *Client) listWorkforcePool(ctx context.Context, r *WorkforcePool, pageToken string, pageSize int32) ([]*WorkforcePool, string, error) {
	b, err := c.listWorkforcePoolRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listWorkforcePoolOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*WorkforcePool
	for _, v := range m.WorkforcePools {
		res, err := unmarshalMapWorkforcePool(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Location = r.Location
		res.Parent = r.Parent
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllWorkforcePool(ctx context.Context, f func(*WorkforcePool) bool, resources []*WorkforcePool) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteWorkforcePool(ctx, res)
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

type deleteWorkforcePoolOperation struct{}

func (op *deleteWorkforcePoolOperation) do(ctx context.Context, r *WorkforcePool, c *Client) error {
	r, err := c.GetWorkforcePool(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "WorkforcePool not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetWorkforcePool checking for existence. error: %v", err)
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
		return fmt.Errorf("failed to delete WorkforcePool: %w", err)
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createWorkforcePoolOperation struct {
	response map[string]interface{}
}

func (op *createWorkforcePoolOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createWorkforcePoolOperation) do(ctx context.Context, r *WorkforcePool, c *Client) error {
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

	if _, err := c.GetWorkforcePool(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getWorkforcePoolRaw(ctx context.Context, r *WorkforcePool) ([]byte, error) {

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

func (c *Client) workforcePoolDiffsForRawDesired(ctx context.Context, rawDesired *WorkforcePool, opts ...dcl.ApplyOption) (initial, desired *WorkforcePool, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *WorkforcePool
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*WorkforcePool); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected WorkforcePool, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetWorkforcePool(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a WorkforcePool resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve WorkforcePool resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that WorkforcePool resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeWorkforcePoolDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for WorkforcePool: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for WorkforcePool: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractWorkforcePoolFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeWorkforcePoolInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for WorkforcePool: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeWorkforcePoolDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for WorkforcePool: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffWorkforcePool(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeWorkforcePoolInitialState(rawInitial, rawDesired *WorkforcePool) (*WorkforcePool, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeWorkforcePoolDesiredState(rawDesired, rawInitial *WorkforcePool, opts ...dcl.ApplyOption) (*WorkforcePool, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &WorkforcePool{}
	if dcl.NameToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.Parent, rawInitial.Parent) {
		canonicalDesired.Parent = rawInitial.Parent
	} else {
		canonicalDesired.Parent = rawDesired.Parent
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
	if dcl.StringCanonicalize(rawDesired.SessionDuration, rawInitial.SessionDuration) {
		canonicalDesired.SessionDuration = rawInitial.SessionDuration
	} else {
		canonicalDesired.SessionDuration = rawDesired.SessionDuration
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}
	return canonicalDesired, nil
}

func canonicalizeWorkforcePoolNewState(c *Client, rawNew, rawDesired *WorkforcePool) (*WorkforcePool, error) {

	rawNew.Name = rawDesired.Name

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Parent) && dcl.IsEmptyValueIndirect(rawDesired.Parent) {
		rawNew.Parent = rawDesired.Parent
	} else {
		if dcl.StringCanonicalize(rawDesired.Parent, rawNew.Parent) {
			rawNew.Parent = rawDesired.Parent
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

	if dcl.IsEmptyValueIndirect(rawNew.SessionDuration) && dcl.IsEmptyValueIndirect(rawDesired.SessionDuration) {
		rawNew.SessionDuration = rawDesired.SessionDuration
	} else {
		if dcl.StringCanonicalize(rawDesired.SessionDuration, rawNew.SessionDuration) {
			rawNew.SessionDuration = rawDesired.SessionDuration
		}
	}

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
func diffWorkforcePool(c *Client, desired, actual *WorkforcePool, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateWorkforcePoolUpdateWorkforcePoolOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateWorkforcePoolUpdateWorkforcePoolOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Disabled, actual.Disabled, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateWorkforcePoolUpdateWorkforcePoolOperation")}, fn.AddNest("Disabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SessionDuration, actual.SessionDuration, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.TriggersOperation("updateWorkforcePoolUpdateWorkforcePoolOperation")}, fn.AddNest("SessionDuration")); len(ds) != 0 || err != nil {
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
func (r *WorkforcePool) urlNormalized() *WorkforcePool {
	normalized := dcl.Copy(*r).(WorkforcePool)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Parent = r.Parent
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.SessionDuration = dcl.SelfLinkToName(r.SessionDuration)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *WorkforcePool) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateWorkforcePool" {
		fields := map[string]interface{}{
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("locations/{{location}}/workforcePools/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the WorkforcePool resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *WorkforcePool) marshal(c *Client) ([]byte, error) {
	m, err := expandWorkforcePool(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling WorkforcePool: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalWorkforcePool decodes JSON responses into the WorkforcePool resource schema.
func unmarshalWorkforcePool(b []byte, c *Client, res *WorkforcePool) (*WorkforcePool, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapWorkforcePool(m, c, res)
}

func unmarshalMapWorkforcePool(m map[string]interface{}, c *Client, res *WorkforcePool) (*WorkforcePool, error) {

	flattened := flattenWorkforcePool(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandWorkforcePool expands WorkforcePool into a JSON request object.
func expandWorkforcePool(c *Client, f *WorkforcePool) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Parent; dcl.ValueShouldBeSent(v) {
		m["parent"] = v
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
	if v := f.SessionDuration; dcl.ValueShouldBeSent(v) {
		m["sessionDuration"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}

	return m, nil
}

// flattenWorkforcePool flattens WorkforcePool from a JSON request object into the
// WorkforcePool type.
func flattenWorkforcePool(c *Client, i interface{}, res *WorkforcePool) *WorkforcePool {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &WorkforcePool{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.SelfLink = dcl.FlattenString(m["name"])
	resultRes.Parent = dcl.FlattenString(m["parent"])
	resultRes.DisplayName = dcl.FlattenString(m["displayName"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.State = flattenWorkforcePoolStateEnum(m["state"])
	resultRes.Disabled = dcl.FlattenBool(m["disabled"])
	resultRes.SessionDuration = dcl.FlattenString(m["sessionDuration"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// flattenWorkforcePoolStateEnumMap flattens the contents of WorkforcePoolStateEnum from a JSON
// response object.
func flattenWorkforcePoolStateEnumMap(c *Client, i interface{}, res *WorkforcePool) map[string]WorkforcePoolStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkforcePoolStateEnum{}
	}

	if len(a) == 0 {
		return map[string]WorkforcePoolStateEnum{}
	}

	items := make(map[string]WorkforcePoolStateEnum)
	for k, item := range a {
		items[k] = *flattenWorkforcePoolStateEnum(item.(interface{}))
	}

	return items
}

// flattenWorkforcePoolStateEnumSlice flattens the contents of WorkforcePoolStateEnum from a JSON
// response object.
func flattenWorkforcePoolStateEnumSlice(c *Client, i interface{}, res *WorkforcePool) []WorkforcePoolStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkforcePoolStateEnum{}
	}

	if len(a) == 0 {
		return []WorkforcePoolStateEnum{}
	}

	items := make([]WorkforcePoolStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkforcePoolStateEnum(item.(interface{})))
	}

	return items
}

// flattenWorkforcePoolStateEnum asserts that an interface is a string, and returns a
// pointer to a *WorkforcePoolStateEnum with the same value as that string.
func flattenWorkforcePoolStateEnum(i interface{}) *WorkforcePoolStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return WorkforcePoolStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *WorkforcePool) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalWorkforcePool(b, c, r)
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

type workforcePoolDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         workforcePoolApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToWorkforcePoolDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]workforcePoolDiff, error) {
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
	var diffs []workforcePoolDiff
	// For each operation name, create a workforcePoolDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := workforcePoolDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToWorkforcePoolApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToWorkforcePoolApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (workforcePoolApiOperation, error) {
	switch opName {

	case "updateWorkforcePoolUpdateWorkforcePoolOperation":
		return &updateWorkforcePoolUpdateWorkforcePoolOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractWorkforcePoolFields(r *WorkforcePool) error {
	return nil
}

func postReadExtractWorkforcePoolFields(r *WorkforcePool) error {
	return nil
}
