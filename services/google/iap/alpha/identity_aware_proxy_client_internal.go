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
)

func (r *IdentityAwareProxyClient) validate() error {

	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Brand, "Brand"); err != nil {
		return err
	}
	return nil
}
func (r *IdentityAwareProxyClient) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://iap.googleapis.com/v1/", params)
}

func (r *IdentityAwareProxyClient) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"brand":   dcl.ValueOrEmptyString(nr.Brand),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/brands/{{brand}}/identityAwareProxyClients/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *IdentityAwareProxyClient) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"brand":   dcl.ValueOrEmptyString(nr.Brand),
	}
	return dcl.URL("projects/{{project}}/brands/{{brand}}/identityAwareProxyClients", nr.basePath(), userBasePath, params), nil

}

func (r *IdentityAwareProxyClient) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"brand":   dcl.ValueOrEmptyString(nr.Brand),
	}
	return dcl.URL("projects/{{project}}/brands/{{brand}}/identityAwareProxyClients", nr.basePath(), userBasePath, params), nil

}

func (r *IdentityAwareProxyClient) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"brand":   dcl.ValueOrEmptyString(nr.Brand),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/brands/{{brand}}/identityAwareProxyClients/{{name}}", nr.basePath(), userBasePath, params), nil
}

// identityAwareProxyClientApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type identityAwareProxyClientApiOperation interface {
	do(context.Context, *IdentityAwareProxyClient, *Client) error
}

func (c *Client) listIdentityAwareProxyClientRaw(ctx context.Context, r *IdentityAwareProxyClient, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != IdentityAwareProxyClientMaxPage {
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

type listIdentityAwareProxyClientOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listIdentityAwareProxyClient(ctx context.Context, r *IdentityAwareProxyClient, pageToken string, pageSize int32) ([]*IdentityAwareProxyClient, string, error) {
	b, err := c.listIdentityAwareProxyClientRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listIdentityAwareProxyClientOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*IdentityAwareProxyClient
	for _, v := range m.Items {
		res, err := unmarshalMapIdentityAwareProxyClient(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Brand = r.Brand
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllIdentityAwareProxyClient(ctx context.Context, f func(*IdentityAwareProxyClient) bool, resources []*IdentityAwareProxyClient) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteIdentityAwareProxyClient(ctx, res)
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

type deleteIdentityAwareProxyClientOperation struct{}

func (op *deleteIdentityAwareProxyClientOperation) do(ctx context.Context, r *IdentityAwareProxyClient, c *Client) error {
	r, err := c.GetIdentityAwareProxyClient(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "IdentityAwareProxyClient not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetIdentityAwareProxyClient checking for existence. error: %v", err)
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
		return fmt.Errorf("failed to delete IdentityAwareProxyClient: %w", err)
	}

	// We saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// This is the reason we are adding retry to handle that case.
	retriesRemaining := 10
	dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		_, err := c.GetIdentityAwareProxyClient(ctx, r)
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
type createIdentityAwareProxyClientOperation struct {
	response map[string]interface{}
}

func (op *createIdentityAwareProxyClientOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createIdentityAwareProxyClientOperation) do(ctx context.Context, r *IdentityAwareProxyClient, c *Client) error {
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
		// Allowing creation to continue with Name set could result in a IdentityAwareProxyClient with the wrong Name.
		return fmt.Errorf("server-generated parameter Name was specified by user as %v, should be unspecified", dcl.ValueOrEmptyString(r.Name))
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

	// Include Name in URL substitution for initial GET request.
	m := op.response
	r.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))

	if _, err := c.GetIdentityAwareProxyClient(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getIdentityAwareProxyClientRaw(ctx context.Context, r *IdentityAwareProxyClient) ([]byte, error) {

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

func (c *Client) identityAwareProxyClientDiffsForRawDesired(ctx context.Context, rawDesired *IdentityAwareProxyClient, opts ...dcl.ApplyOption) (initial, desired *IdentityAwareProxyClient, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *IdentityAwareProxyClient
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*IdentityAwareProxyClient); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected IdentityAwareProxyClient, got %T", sh)
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
		desired, err := canonicalizeIdentityAwareProxyClientDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}
	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetIdentityAwareProxyClient(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a IdentityAwareProxyClient resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve IdentityAwareProxyClient resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that IdentityAwareProxyClient resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeIdentityAwareProxyClientDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for IdentityAwareProxyClient: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for IdentityAwareProxyClient: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractIdentityAwareProxyClientFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeIdentityAwareProxyClientInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for IdentityAwareProxyClient: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeIdentityAwareProxyClientDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for IdentityAwareProxyClient: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffIdentityAwareProxyClient(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeIdentityAwareProxyClientInitialState(rawInitial, rawDesired *IdentityAwareProxyClient) (*IdentityAwareProxyClient, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeIdentityAwareProxyClientDesiredState(rawDesired, rawInitial *IdentityAwareProxyClient, opts ...dcl.ApplyOption) (*IdentityAwareProxyClient, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &IdentityAwareProxyClient{}
	if dcl.IsZeroValue(rawDesired.Name) || (dcl.IsEmptyValueIndirect(rawDesired.Name) && dcl.IsEmptyValueIndirect(rawInitial.Name)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		canonicalDesired.DisplayName = rawInitial.DisplayName
	} else {
		canonicalDesired.DisplayName = rawDesired.DisplayName
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.NameToSelfLink(rawDesired.Brand, rawInitial.Brand) {
		canonicalDesired.Brand = rawInitial.Brand
	} else {
		canonicalDesired.Brand = rawDesired.Brand
	}
	return canonicalDesired, nil
}

func canonicalizeIdentityAwareProxyClientNewState(c *Client, rawNew, rawDesired *IdentityAwareProxyClient) (*IdentityAwareProxyClient, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Secret) && dcl.IsEmptyValueIndirect(rawDesired.Secret) {
		rawNew.Secret = rawDesired.Secret
	} else {
		if dcl.StringCanonicalize(rawDesired.Secret, rawNew.Secret) {
			rawNew.Secret = rawDesired.Secret
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	rawNew.Project = rawDesired.Project

	rawNew.Brand = rawDesired.Brand

	return rawNew, nil
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffIdentityAwareProxyClient(c *Client, desired, actual *IdentityAwareProxyClient, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Secret, actual.Secret, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Secret")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Brand, actual.Brand, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Brand")); len(ds) != 0 || err != nil {
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
func (r *IdentityAwareProxyClient) urlNormalized() *IdentityAwareProxyClient {
	normalized := dcl.Copy(*r).(IdentityAwareProxyClient)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Secret = dcl.SelfLinkToName(r.Secret)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Brand = dcl.SelfLinkToName(r.Brand)
	return &normalized
}

func (r *IdentityAwareProxyClient) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the IdentityAwareProxyClient resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *IdentityAwareProxyClient) marshal(c *Client) ([]byte, error) {
	m, err := expandIdentityAwareProxyClient(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling IdentityAwareProxyClient: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalIdentityAwareProxyClient decodes JSON responses into the IdentityAwareProxyClient resource schema.
func unmarshalIdentityAwareProxyClient(b []byte, c *Client, res *IdentityAwareProxyClient) (*IdentityAwareProxyClient, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapIdentityAwareProxyClient(m, c, res)
}

func unmarshalMapIdentityAwareProxyClient(m map[string]interface{}, c *Client, res *IdentityAwareProxyClient) (*IdentityAwareProxyClient, error) {

	flattened := flattenIdentityAwareProxyClient(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandIdentityAwareProxyClient expands IdentityAwareProxyClient into a JSON request object.
func expandIdentityAwareProxyClient(c *Client, f *IdentityAwareProxyClient) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("projects/%s/brands/%s/identityAwareProxyClients/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Brand), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.DisplayName; dcl.ValueShouldBeSent(v) {
		m["displayName"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Brand into brand: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["brand"] = v
	}

	return m, nil
}

// flattenIdentityAwareProxyClient flattens IdentityAwareProxyClient from a JSON request object into the
// IdentityAwareProxyClient type.
func flattenIdentityAwareProxyClient(c *Client, i interface{}, res *IdentityAwareProxyClient) *IdentityAwareProxyClient {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &IdentityAwareProxyClient{}
	resultRes.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	resultRes.Secret = dcl.FlattenString(m["secret"])
	resultRes.DisplayName = dcl.FlattenString(m["displayName"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Brand = dcl.FlattenString(m["brand"])

	return resultRes
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *IdentityAwareProxyClient) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalIdentityAwareProxyClient(b, c, r)
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
		if nr.Brand == nil && ncr.Brand == nil {
			c.Config.Logger.Info("Both Brand fields null - considering equal.")
		} else if nr.Brand == nil || ncr.Brand == nil {
			c.Config.Logger.Info("Only one Brand field is null - considering unequal.")
			return false
		} else if *nr.Brand != *ncr.Brand {
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

type identityAwareProxyClientDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         identityAwareProxyClientApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToIdentityAwareProxyClientDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]identityAwareProxyClientDiff, error) {
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
	var diffs []identityAwareProxyClientDiff
	// For each operation name, create a identityAwareProxyClientDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := identityAwareProxyClientDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToIdentityAwareProxyClientApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToIdentityAwareProxyClientApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (identityAwareProxyClientApiOperation, error) {
	switch opName {

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractIdentityAwareProxyClientFields(r *IdentityAwareProxyClient) error {
	return nil
}

func postReadExtractIdentityAwareProxyClientFields(r *IdentityAwareProxyClient) error {
	return nil
}
