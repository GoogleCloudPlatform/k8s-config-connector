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

func (r *EnvironmentGroupAttachment) validate() error {

	if err := dcl.Required(r, "environment"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Envgroup, "Envgroup"); err != nil {
		return err
	}
	return nil
}
func (r *EnvironmentGroupAttachment) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://apigee.googleapis.com/v1/", params)
}

func (r *EnvironmentGroupAttachment) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"envgroup": dcl.ValueOrEmptyString(nr.Envgroup),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("{{envgroup}}/attachments/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *EnvironmentGroupAttachment) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"envgroup": dcl.ValueOrEmptyString(nr.Envgroup),
	}
	return dcl.URL("{{envgroup}}/attachments", nr.basePath(), userBasePath, params), nil

}

func (r *EnvironmentGroupAttachment) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"envgroup": dcl.ValueOrEmptyString(nr.Envgroup),
	}
	return dcl.URL("{{envgroup}}/attachments", nr.basePath(), userBasePath, params), nil

}

func (r *EnvironmentGroupAttachment) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"envgroup": dcl.ValueOrEmptyString(nr.Envgroup),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("{{envgroup}}/attachments/{{name}}", nr.basePath(), userBasePath, params), nil
}

// environmentGroupAttachmentApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type environmentGroupAttachmentApiOperation interface {
	do(context.Context, *EnvironmentGroupAttachment, *Client) error
}

func (c *Client) listEnvironmentGroupAttachmentRaw(ctx context.Context, r *EnvironmentGroupAttachment, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != EnvironmentGroupAttachmentMaxPage {
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

type listEnvironmentGroupAttachmentOperation struct {
	EnvironmentGroupAttachments []map[string]interface{} `json:"environmentGroupAttachments"`
	Token                       string                   `json:"nextPageToken"`
}

func (c *Client) listEnvironmentGroupAttachment(ctx context.Context, r *EnvironmentGroupAttachment, pageToken string, pageSize int32) ([]*EnvironmentGroupAttachment, string, error) {
	b, err := c.listEnvironmentGroupAttachmentRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listEnvironmentGroupAttachmentOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*EnvironmentGroupAttachment
	for _, v := range m.EnvironmentGroupAttachments {
		res, err := unmarshalMapEnvironmentGroupAttachment(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Envgroup = r.Envgroup
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllEnvironmentGroupAttachment(ctx context.Context, f func(*EnvironmentGroupAttachment) bool, resources []*EnvironmentGroupAttachment) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteEnvironmentGroupAttachment(ctx, res)
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

type deleteEnvironmentGroupAttachmentOperation struct{}

func (op *deleteEnvironmentGroupAttachmentOperation) do(ctx context.Context, r *EnvironmentGroupAttachment, c *Client) error {
	r, err := c.GetEnvironmentGroupAttachment(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "EnvironmentGroupAttachment not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetEnvironmentGroupAttachment checking for existence. error: %v", err)
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
		_, err := c.GetEnvironmentGroupAttachment(ctx, r)
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
type createEnvironmentGroupAttachmentOperation struct {
	response map[string]interface{}
}

func (op *createEnvironmentGroupAttachmentOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createEnvironmentGroupAttachmentOperation) do(ctx context.Context, r *EnvironmentGroupAttachment, c *Client) error {
	c.Config.Logger.InfoWithContextf(ctx, "Attempting to create %v", r)
	u, err := r.createURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	if r.Name != nil {
		// Allowing creation to continue with Name set could result in a EnvironmentGroupAttachment with the wrong Name.
		return fmt.Errorf("server-generated parameter Name was specified by user as %v, should be unspecified", dcl.ValueOrEmptyString(r.Name))
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

	// Include Name in URL substitution for initial GET request.
	m := op.response
	r.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))

	if _, err := c.GetEnvironmentGroupAttachment(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getEnvironmentGroupAttachmentRaw(ctx context.Context, r *EnvironmentGroupAttachment) ([]byte, error) {

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

func (c *Client) environmentGroupAttachmentDiffsForRawDesired(ctx context.Context, rawDesired *EnvironmentGroupAttachment, opts ...dcl.ApplyOption) (initial, desired *EnvironmentGroupAttachment, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *EnvironmentGroupAttachment
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*EnvironmentGroupAttachment); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected EnvironmentGroupAttachment, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	if fetchState.Name == nil {
		// We cannot perform a get because of lack of information. We have to assume
		// that this is being created for the first time.
		desired, err := canonicalizeEnvironmentGroupAttachmentDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}
	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetEnvironmentGroupAttachment(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a EnvironmentGroupAttachment resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve EnvironmentGroupAttachment resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that EnvironmentGroupAttachment resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeEnvironmentGroupAttachmentDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for EnvironmentGroupAttachment: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for EnvironmentGroupAttachment: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractEnvironmentGroupAttachmentFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeEnvironmentGroupAttachmentInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for EnvironmentGroupAttachment: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeEnvironmentGroupAttachmentDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for EnvironmentGroupAttachment: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffEnvironmentGroupAttachment(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeEnvironmentGroupAttachmentInitialState(rawInitial, rawDesired *EnvironmentGroupAttachment) (*EnvironmentGroupAttachment, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeEnvironmentGroupAttachmentDesiredState(rawDesired, rawInitial *EnvironmentGroupAttachment, opts ...dcl.ApplyOption) (*EnvironmentGroupAttachment, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &EnvironmentGroupAttachment{}
	if dcl.IsZeroValue(rawDesired.Name) || (dcl.IsEmptyValueIndirect(rawDesired.Name) && dcl.IsEmptyValueIndirect(rawInitial.Name)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.IsZeroValue(rawDesired.Environment) || (dcl.IsEmptyValueIndirect(rawDesired.Environment) && dcl.IsEmptyValueIndirect(rawInitial.Environment)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Environment = rawInitial.Environment
	} else {
		canonicalDesired.Environment = rawDesired.Environment
	}
	if dcl.NameToSelfLink(rawDesired.Envgroup, rawInitial.Envgroup) {
		canonicalDesired.Envgroup = rawInitial.Envgroup
	} else {
		canonicalDesired.Envgroup = rawDesired.Envgroup
	}

	return canonicalDesired, nil
}

func canonicalizeEnvironmentGroupAttachmentNewState(c *Client, rawNew, rawDesired *EnvironmentGroupAttachment) (*EnvironmentGroupAttachment, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Environment) && dcl.IsEmptyValueIndirect(rawDesired.Environment) {
		rawNew.Environment = rawDesired.Environment
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreatedAt) && dcl.IsEmptyValueIndirect(rawDesired.CreatedAt) {
		rawNew.CreatedAt = rawDesired.CreatedAt
	} else {
	}

	rawNew.Envgroup = rawDesired.Envgroup

	return rawNew, nil
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffEnvironmentGroupAttachment(c *Client, desired, actual *EnvironmentGroupAttachment, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Environment, actual.Environment, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Environment")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Envgroup, actual.Envgroup, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Envgroup")); len(ds) != 0 || err != nil {
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
func (r *EnvironmentGroupAttachment) urlNormalized() *EnvironmentGroupAttachment {
	normalized := dcl.Copy(*r).(EnvironmentGroupAttachment)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Environment = dcl.SelfLinkToName(r.Environment)
	normalized.Envgroup = r.Envgroup
	return &normalized
}

func (r *EnvironmentGroupAttachment) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the EnvironmentGroupAttachment resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *EnvironmentGroupAttachment) marshal(c *Client) ([]byte, error) {
	m, err := expandEnvironmentGroupAttachment(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling EnvironmentGroupAttachment: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalEnvironmentGroupAttachment decodes JSON responses into the EnvironmentGroupAttachment resource schema.
func unmarshalEnvironmentGroupAttachment(b []byte, c *Client, res *EnvironmentGroupAttachment) (*EnvironmentGroupAttachment, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapEnvironmentGroupAttachment(m, c, res)
}

func unmarshalMapEnvironmentGroupAttachment(m map[string]interface{}, c *Client, res *EnvironmentGroupAttachment) (*EnvironmentGroupAttachment, error) {

	flattened := flattenEnvironmentGroupAttachment(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandEnvironmentGroupAttachment expands EnvironmentGroupAttachment into a JSON request object.
func expandEnvironmentGroupAttachment(c *Client, f *EnvironmentGroupAttachment) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v := f.Name; dcl.ValueShouldBeSent(v) {
		m["name"] = v
	}
	if v, err := dcl.SelfLinkToNameExpander(f.Environment); err != nil {
		return nil, fmt.Errorf("error expanding Environment into environment: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["environment"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Envgroup into envgroup: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["envgroup"] = v
	}

	return m, nil
}

// flattenEnvironmentGroupAttachment flattens EnvironmentGroupAttachment from a JSON request object into the
// EnvironmentGroupAttachment type.
func flattenEnvironmentGroupAttachment(c *Client, i interface{}, res *EnvironmentGroupAttachment) *EnvironmentGroupAttachment {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &EnvironmentGroupAttachment{}
	resultRes.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	resultRes.Environment = dcl.FlattenString(m["environment"])
	resultRes.CreatedAt = dcl.FlattenInteger(m["createdAt"])
	resultRes.Envgroup = dcl.FlattenString(m["envgroup"])

	return resultRes
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *EnvironmentGroupAttachment) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalEnvironmentGroupAttachment(b, c, r)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Envgroup == nil && ncr.Envgroup == nil {
			c.Config.Logger.Info("Both Envgroup fields null - considering equal.")
		} else if nr.Envgroup == nil || ncr.Envgroup == nil {
			c.Config.Logger.Info("Only one Envgroup field is null - considering unequal.")
			return false
		} else if *nr.Envgroup != *ncr.Envgroup {
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

type environmentGroupAttachmentDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         environmentGroupAttachmentApiOperation
}

func convertFieldDiffsToEnvironmentGroupAttachmentDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]environmentGroupAttachmentDiff, error) {
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
	var diffs []environmentGroupAttachmentDiff
	// For each operation name, create a environmentGroupAttachmentDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := environmentGroupAttachmentDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToEnvironmentGroupAttachmentApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToEnvironmentGroupAttachmentApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (environmentGroupAttachmentApiOperation, error) {
	switch opName {

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractEnvironmentGroupAttachmentFields(r *EnvironmentGroupAttachment) error {
	return nil
}

func postReadExtractEnvironmentGroupAttachmentFields(r *EnvironmentGroupAttachment) error {
	return nil
}
