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
package binaryauthorization

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *Attestor) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.UserOwnedDrydockNote) {
		if err := r.UserOwnedDrydockNote.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AttestorUserOwnedDrydockNote) validate() error {
	if err := dcl.Required(r, "noteReference"); err != nil {
		return err
	}
	return nil
}
func (r *AttestorUserOwnedDrydockNotePublicKeys) validate() error {
	if !dcl.IsEmptyValueIndirect(r.PkixPublicKey) {
		if err := r.PkixPublicKey.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) validate() error {
	return nil
}
func (r *Attestor) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://binaryauthorization.googleapis.com/v1", params)
}

func (r *Attestor) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/attestors/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Attestor) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.URL("projects/{{project}}/attestors", nr.basePath(), userBasePath, params), nil

}

func (r *Attestor) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/attestors?attestorId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *Attestor) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/attestors/{{name}}", nr.basePath(), userBasePath, params), nil
}

// attestorApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type attestorApiOperation interface {
	do(context.Context, *Attestor, *Client) error
}

// newUpdateAttestorUpdateAttestorRequest creates a request for an
// Attestor resource's UpdateAttestor update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateAttestorUpdateAttestorRequest(ctx context.Context, f *Attestor, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v, err := expandAttestorUserOwnedDrydockNote(c, f.UserOwnedDrydockNote, res); err != nil {
		return nil, fmt.Errorf("error expanding UserOwnedDrydockNote into userOwnedGrafeasNote: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["userOwnedGrafeasNote"] = v
	}
	return req, nil
}

// marshalUpdateAttestorUpdateAttestorRequest converts the update into
// the final JSON request body.
func marshalUpdateAttestorUpdateAttestorRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateAttestorUpdateAttestorOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateAttestorUpdateAttestorOperation) do(ctx context.Context, r *Attestor, c *Client) error {
	_, err := c.GetAttestor(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateAttestor")
	if err != nil {
		return err
	}

	req, err := newUpdateAttestorUpdateAttestorRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateAttestorUpdateAttestorRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listAttestorRaw(ctx context.Context, r *Attestor, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != AttestorMaxPage {
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

type listAttestorOperation struct {
	Attestors []map[string]interface{} `json:"attestors"`
	Token     string                   `json:"nextPageToken"`
}

func (c *Client) listAttestor(ctx context.Context, r *Attestor, pageToken string, pageSize int32) ([]*Attestor, string, error) {
	b, err := c.listAttestorRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listAttestorOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Attestor
	for _, v := range m.Attestors {
		res, err := unmarshalMapAttestor(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllAttestor(ctx context.Context, f func(*Attestor) bool, resources []*Attestor) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteAttestor(ctx, res)
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

type deleteAttestorOperation struct{}

func (op *deleteAttestorOperation) do(ctx context.Context, r *Attestor, c *Client) error {
	r, err := c.GetAttestor(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Attestor not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetAttestor checking for existence. error: %v", err)
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
		return fmt.Errorf("failed to delete Attestor: %w", err)
	}

	// We saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// This is the reason we are adding retry to handle that case.
	retriesRemaining := 10
	dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		_, err := c.GetAttestor(ctx, r)
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
type createAttestorOperation struct {
	response map[string]interface{}
}

func (op *createAttestorOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createAttestorOperation) do(ctx context.Context, r *Attestor, c *Client) error {
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

	if _, err := c.GetAttestor(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getAttestorRaw(ctx context.Context, r *Attestor) ([]byte, error) {

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

func (c *Client) attestorDiffsForRawDesired(ctx context.Context, rawDesired *Attestor, opts ...dcl.ApplyOption) (initial, desired *Attestor, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Attestor
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Attestor); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Attestor, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetAttestor(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Attestor resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Attestor resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Attestor resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeAttestorDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Attestor: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Attestor: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractAttestorFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeAttestorInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Attestor: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeAttestorDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Attestor: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffAttestor(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeAttestorInitialState(rawInitial, rawDesired *Attestor) (*Attestor, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeAttestorDesiredState(rawDesired, rawInitial *Attestor, opts ...dcl.ApplyOption) (*Attestor, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.UserOwnedDrydockNote = canonicalizeAttestorUserOwnedDrydockNote(rawDesired.UserOwnedDrydockNote, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Attestor{}
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
	canonicalDesired.UserOwnedDrydockNote = canonicalizeAttestorUserOwnedDrydockNote(rawDesired.UserOwnedDrydockNote, rawInitial.UserOwnedDrydockNote, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	return canonicalDesired, nil
}

func canonicalizeAttestorNewState(c *Client, rawNew, rawDesired *Attestor) (*Attestor, error) {

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

	if dcl.IsEmptyValueIndirect(rawNew.UserOwnedDrydockNote) && dcl.IsEmptyValueIndirect(rawDesired.UserOwnedDrydockNote) {
		rawNew.UserOwnedDrydockNote = rawDesired.UserOwnedDrydockNote
	} else {
		rawNew.UserOwnedDrydockNote = canonicalizeNewAttestorUserOwnedDrydockNote(c, rawDesired.UserOwnedDrydockNote, rawNew.UserOwnedDrydockNote)
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeAttestorUserOwnedDrydockNote(des, initial *AttestorUserOwnedDrydockNote, opts ...dcl.ApplyOption) *AttestorUserOwnedDrydockNote {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AttestorUserOwnedDrydockNote{}

	if dcl.IsZeroValue(des.NoteReference) || (dcl.IsEmptyValueIndirect(des.NoteReference) && dcl.IsEmptyValueIndirect(initial.NoteReference)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.NoteReference = initial.NoteReference
	} else {
		cDes.NoteReference = des.NoteReference
	}
	cDes.PublicKeys = canonicalizeAttestorUserOwnedDrydockNotePublicKeysSlice(des.PublicKeys, initial.PublicKeys, opts...)

	return cDes
}

func canonicalizeAttestorUserOwnedDrydockNoteSlice(des, initial []AttestorUserOwnedDrydockNote, opts ...dcl.ApplyOption) []AttestorUserOwnedDrydockNote {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]AttestorUserOwnedDrydockNote, 0, len(des))
		for _, d := range des {
			cd := canonicalizeAttestorUserOwnedDrydockNote(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]AttestorUserOwnedDrydockNote, 0, len(des))
	for i, d := range des {
		cd := canonicalizeAttestorUserOwnedDrydockNote(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewAttestorUserOwnedDrydockNote(c *Client, des, nw *AttestorUserOwnedDrydockNote) *AttestorUserOwnedDrydockNote {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for AttestorUserOwnedDrydockNote while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.PublicKeys = canonicalizeNewAttestorUserOwnedDrydockNotePublicKeysSlice(c, des.PublicKeys, nw.PublicKeys)
	if dcl.StringCanonicalize(des.DelegationServiceAccountEmail, nw.DelegationServiceAccountEmail) {
		nw.DelegationServiceAccountEmail = des.DelegationServiceAccountEmail
	}

	return nw
}

func canonicalizeNewAttestorUserOwnedDrydockNoteSet(c *Client, des, nw []AttestorUserOwnedDrydockNote) []AttestorUserOwnedDrydockNote {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []AttestorUserOwnedDrydockNote
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareAttestorUserOwnedDrydockNoteNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewAttestorUserOwnedDrydockNote(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewAttestorUserOwnedDrydockNoteSlice(c *Client, des, nw []AttestorUserOwnedDrydockNote) []AttestorUserOwnedDrydockNote {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AttestorUserOwnedDrydockNote
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAttestorUserOwnedDrydockNote(c, &d, &n))
	}

	return items
}

func canonicalizeAttestorUserOwnedDrydockNotePublicKeys(des, initial *AttestorUserOwnedDrydockNotePublicKeys, opts ...dcl.ApplyOption) *AttestorUserOwnedDrydockNotePublicKeys {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AttestorUserOwnedDrydockNotePublicKeys{}

	if dcl.StringCanonicalize(des.Comment, initial.Comment) || dcl.IsZeroValue(des.Comment) {
		cDes.Comment = initial.Comment
	} else {
		cDes.Comment = des.Comment
	}
	if dcl.StringCanonicalize(des.Id, initial.Id) || dcl.IsZeroValue(des.Id) {
		cDes.Id = initial.Id
	} else {
		cDes.Id = des.Id
	}
	if dcl.StringCanonicalize(des.AsciiArmoredPgpPublicKey, initial.AsciiArmoredPgpPublicKey) || dcl.IsZeroValue(des.AsciiArmoredPgpPublicKey) {
		cDes.AsciiArmoredPgpPublicKey = initial.AsciiArmoredPgpPublicKey
	} else {
		cDes.AsciiArmoredPgpPublicKey = des.AsciiArmoredPgpPublicKey
	}
	cDes.PkixPublicKey = canonicalizeAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(des.PkixPublicKey, initial.PkixPublicKey, opts...)

	return cDes
}

func canonicalizeAttestorUserOwnedDrydockNotePublicKeysSlice(des, initial []AttestorUserOwnedDrydockNotePublicKeys, opts ...dcl.ApplyOption) []AttestorUserOwnedDrydockNotePublicKeys {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]AttestorUserOwnedDrydockNotePublicKeys, 0, len(des))
		for _, d := range des {
			cd := canonicalizeAttestorUserOwnedDrydockNotePublicKeys(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]AttestorUserOwnedDrydockNotePublicKeys, 0, len(des))
	for i, d := range des {
		cd := canonicalizeAttestorUserOwnedDrydockNotePublicKeys(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewAttestorUserOwnedDrydockNotePublicKeys(c *Client, des, nw *AttestorUserOwnedDrydockNotePublicKeys) *AttestorUserOwnedDrydockNotePublicKeys {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for AttestorUserOwnedDrydockNotePublicKeys while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Comment, nw.Comment) {
		nw.Comment = des.Comment
	}
	if dcl.StringCanonicalize(des.Id, nw.Id) {
		nw.Id = des.Id
	}
	if dcl.StringCanonicalize(des.AsciiArmoredPgpPublicKey, nw.AsciiArmoredPgpPublicKey) {
		nw.AsciiArmoredPgpPublicKey = des.AsciiArmoredPgpPublicKey
	}
	nw.PkixPublicKey = canonicalizeNewAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(c, des.PkixPublicKey, nw.PkixPublicKey)

	return nw
}

func canonicalizeNewAttestorUserOwnedDrydockNotePublicKeysSet(c *Client, des, nw []AttestorUserOwnedDrydockNotePublicKeys) []AttestorUserOwnedDrydockNotePublicKeys {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []AttestorUserOwnedDrydockNotePublicKeys
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareAttestorUserOwnedDrydockNotePublicKeysNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewAttestorUserOwnedDrydockNotePublicKeys(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewAttestorUserOwnedDrydockNotePublicKeysSlice(c *Client, des, nw []AttestorUserOwnedDrydockNotePublicKeys) []AttestorUserOwnedDrydockNotePublicKeys {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AttestorUserOwnedDrydockNotePublicKeys
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAttestorUserOwnedDrydockNotePublicKeys(c, &d, &n))
	}

	return items
}

func canonicalizeAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(des, initial *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey, opts ...dcl.ApplyOption) *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey{}

	if dcl.StringCanonicalize(des.PublicKeyPem, initial.PublicKeyPem) || dcl.IsZeroValue(des.PublicKeyPem) {
		cDes.PublicKeyPem = initial.PublicKeyPem
	} else {
		cDes.PublicKeyPem = des.PublicKeyPem
	}
	if dcl.IsZeroValue(des.SignatureAlgorithm) || (dcl.IsEmptyValueIndirect(des.SignatureAlgorithm) && dcl.IsEmptyValueIndirect(initial.SignatureAlgorithm)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.SignatureAlgorithm = initial.SignatureAlgorithm
	} else {
		cDes.SignatureAlgorithm = des.SignatureAlgorithm
	}

	return cDes
}

func canonicalizeAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySlice(des, initial []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey, opts ...dcl.ApplyOption) []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey, 0, len(des))
		for _, d := range des {
			cd := canonicalizeAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey, 0, len(des))
	for i, d := range des {
		cd := canonicalizeAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(c *Client, des, nw *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.PublicKeyPem, nw.PublicKeyPem) {
		nw.PublicKeyPem = des.PublicKeyPem
	}

	return nw
}

func canonicalizeNewAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySet(c *Client, des, nw []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySlice(c *Client, des, nw []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(c, &d, &n))
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
func diffAttestor(c *Client, desired, actual *Attestor, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateAttestorUpdateAttestorOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UserOwnedDrydockNote, actual.UserOwnedDrydockNote, dcl.DiffInfo{ObjectFunction: compareAttestorUserOwnedDrydockNoteNewStyle, EmptyObject: EmptyAttestorUserOwnedDrydockNote, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UserOwnedGrafeasNote")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
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
func compareAttestorUserOwnedDrydockNoteNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AttestorUserOwnedDrydockNote)
	if !ok {
		desiredNotPointer, ok := d.(AttestorUserOwnedDrydockNote)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AttestorUserOwnedDrydockNote or *AttestorUserOwnedDrydockNote", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AttestorUserOwnedDrydockNote)
	if !ok {
		actualNotPointer, ok := a.(AttestorUserOwnedDrydockNote)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AttestorUserOwnedDrydockNote", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.NoteReference, actual.NoteReference, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NoteReference")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PublicKeys, actual.PublicKeys, dcl.DiffInfo{ObjectFunction: compareAttestorUserOwnedDrydockNotePublicKeysNewStyle, EmptyObject: EmptyAttestorUserOwnedDrydockNotePublicKeys, OperationSelector: dcl.TriggersOperation("updateAttestorUpdateAttestorOperation")}, fn.AddNest("PublicKeys")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DelegationServiceAccountEmail, actual.DelegationServiceAccountEmail, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DelegationServiceAccountEmail")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAttestorUserOwnedDrydockNotePublicKeysNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AttestorUserOwnedDrydockNotePublicKeys)
	if !ok {
		desiredNotPointer, ok := d.(AttestorUserOwnedDrydockNotePublicKeys)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AttestorUserOwnedDrydockNotePublicKeys or *AttestorUserOwnedDrydockNotePublicKeys", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AttestorUserOwnedDrydockNotePublicKeys)
	if !ok {
		actualNotPointer, ok := a.(AttestorUserOwnedDrydockNotePublicKeys)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AttestorUserOwnedDrydockNotePublicKeys", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Comment, actual.Comment, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateAttestorUpdateAttestorOperation")}, fn.AddNest("Comment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.TriggersOperation("updateAttestorUpdateAttestorOperation")}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AsciiArmoredPgpPublicKey, actual.AsciiArmoredPgpPublicKey, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateAttestorUpdateAttestorOperation")}, fn.AddNest("AsciiArmoredPgpPublicKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PkixPublicKey, actual.PkixPublicKey, dcl.DiffInfo{ObjectFunction: compareAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyNewStyle, EmptyObject: EmptyAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey, OperationSelector: dcl.TriggersOperation("updateAttestorUpdateAttestorOperation")}, fn.AddNest("PkixPublicKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey)
	if !ok {
		desiredNotPointer, ok := d.(AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey or *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey)
	if !ok {
		actualNotPointer, ok := a.(AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PublicKeyPem, actual.PublicKeyPem, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateAttestorUpdateAttestorOperation")}, fn.AddNest("PublicKeyPem")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SignatureAlgorithm, actual.SignatureAlgorithm, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateAttestorUpdateAttestorOperation")}, fn.AddNest("SignatureAlgorithm")); len(ds) != 0 || err != nil {
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
func (r *Attestor) urlNormalized() *Attestor {
	normalized := dcl.Copy(*r).(Attestor)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *Attestor) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateAttestor" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(nr.Project),
			"name":    dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/attestors/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Attestor resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Attestor) marshal(c *Client) ([]byte, error) {
	m, err := expandAttestor(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Attestor: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalAttestor decodes JSON responses into the Attestor resource schema.
func unmarshalAttestor(b []byte, c *Client, res *Attestor) (*Attestor, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapAttestor(m, c, res)
}

func unmarshalMapAttestor(m map[string]interface{}, c *Client, res *Attestor) (*Attestor, error) {

	flattened := flattenAttestor(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandAttestor expands Attestor into a JSON request object.
func expandAttestor(c *Client, f *Attestor) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("projects/%s/attestors/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v, err := expandAttestorUserOwnedDrydockNote(c, f.UserOwnedDrydockNote, res); err != nil {
		return nil, fmt.Errorf("error expanding UserOwnedDrydockNote into userOwnedGrafeasNote: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["userOwnedGrafeasNote"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}

	return m, nil
}

// flattenAttestor flattens Attestor from a JSON request object into the
// Attestor type.
func flattenAttestor(c *Client, i interface{}, res *Attestor) *Attestor {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Attestor{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.UserOwnedDrydockNote = flattenAttestorUserOwnedDrydockNote(c, m["userOwnedGrafeasNote"], res)
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.Project = dcl.FlattenString(m["project"])

	return resultRes
}

// expandAttestorUserOwnedDrydockNoteMap expands the contents of AttestorUserOwnedDrydockNote into a JSON
// request object.
func expandAttestorUserOwnedDrydockNoteMap(c *Client, f map[string]AttestorUserOwnedDrydockNote, res *Attestor) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAttestorUserOwnedDrydockNote(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAttestorUserOwnedDrydockNoteSlice expands the contents of AttestorUserOwnedDrydockNote into a JSON
// request object.
func expandAttestorUserOwnedDrydockNoteSlice(c *Client, f []AttestorUserOwnedDrydockNote, res *Attestor) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAttestorUserOwnedDrydockNote(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAttestorUserOwnedDrydockNoteMap flattens the contents of AttestorUserOwnedDrydockNote from a JSON
// response object.
func flattenAttestorUserOwnedDrydockNoteMap(c *Client, i interface{}, res *Attestor) map[string]AttestorUserOwnedDrydockNote {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AttestorUserOwnedDrydockNote{}
	}

	if len(a) == 0 {
		return map[string]AttestorUserOwnedDrydockNote{}
	}

	items := make(map[string]AttestorUserOwnedDrydockNote)
	for k, item := range a {
		items[k] = *flattenAttestorUserOwnedDrydockNote(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenAttestorUserOwnedDrydockNoteSlice flattens the contents of AttestorUserOwnedDrydockNote from a JSON
// response object.
func flattenAttestorUserOwnedDrydockNoteSlice(c *Client, i interface{}, res *Attestor) []AttestorUserOwnedDrydockNote {
	a, ok := i.([]interface{})
	if !ok {
		return []AttestorUserOwnedDrydockNote{}
	}

	if len(a) == 0 {
		return []AttestorUserOwnedDrydockNote{}
	}

	items := make([]AttestorUserOwnedDrydockNote, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAttestorUserOwnedDrydockNote(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandAttestorUserOwnedDrydockNote expands an instance of AttestorUserOwnedDrydockNote into a JSON
// request object.
func expandAttestorUserOwnedDrydockNote(c *Client, f *AttestorUserOwnedDrydockNote, res *Attestor) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.NoteReference; !dcl.IsEmptyValueIndirect(v) {
		m["noteReference"] = v
	}
	if v, err := expandAttestorUserOwnedDrydockNotePublicKeysSlice(c, f.PublicKeys, res); err != nil {
		return nil, fmt.Errorf("error expanding PublicKeys into publicKeys: %w", err)
	} else if v != nil {
		m["publicKeys"] = v
	}

	return m, nil
}

// flattenAttestorUserOwnedDrydockNote flattens an instance of AttestorUserOwnedDrydockNote from a JSON
// response object.
func flattenAttestorUserOwnedDrydockNote(c *Client, i interface{}, res *Attestor) *AttestorUserOwnedDrydockNote {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AttestorUserOwnedDrydockNote{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAttestorUserOwnedDrydockNote
	}
	r.NoteReference = dcl.FlattenString(m["noteReference"])
	r.PublicKeys = flattenAttestorUserOwnedDrydockNotePublicKeysSlice(c, m["publicKeys"], res)
	r.DelegationServiceAccountEmail = dcl.FlattenString(m["delegationServiceAccountEmail"])

	return r
}

// expandAttestorUserOwnedDrydockNotePublicKeysMap expands the contents of AttestorUserOwnedDrydockNotePublicKeys into a JSON
// request object.
func expandAttestorUserOwnedDrydockNotePublicKeysMap(c *Client, f map[string]AttestorUserOwnedDrydockNotePublicKeys, res *Attestor) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAttestorUserOwnedDrydockNotePublicKeys(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAttestorUserOwnedDrydockNotePublicKeysSlice expands the contents of AttestorUserOwnedDrydockNotePublicKeys into a JSON
// request object.
func expandAttestorUserOwnedDrydockNotePublicKeysSlice(c *Client, f []AttestorUserOwnedDrydockNotePublicKeys, res *Attestor) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAttestorUserOwnedDrydockNotePublicKeys(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAttestorUserOwnedDrydockNotePublicKeysMap flattens the contents of AttestorUserOwnedDrydockNotePublicKeys from a JSON
// response object.
func flattenAttestorUserOwnedDrydockNotePublicKeysMap(c *Client, i interface{}, res *Attestor) map[string]AttestorUserOwnedDrydockNotePublicKeys {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AttestorUserOwnedDrydockNotePublicKeys{}
	}

	if len(a) == 0 {
		return map[string]AttestorUserOwnedDrydockNotePublicKeys{}
	}

	items := make(map[string]AttestorUserOwnedDrydockNotePublicKeys)
	for k, item := range a {
		items[k] = *flattenAttestorUserOwnedDrydockNotePublicKeys(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenAttestorUserOwnedDrydockNotePublicKeysSlice flattens the contents of AttestorUserOwnedDrydockNotePublicKeys from a JSON
// response object.
func flattenAttestorUserOwnedDrydockNotePublicKeysSlice(c *Client, i interface{}, res *Attestor) []AttestorUserOwnedDrydockNotePublicKeys {
	a, ok := i.([]interface{})
	if !ok {
		return []AttestorUserOwnedDrydockNotePublicKeys{}
	}

	if len(a) == 0 {
		return []AttestorUserOwnedDrydockNotePublicKeys{}
	}

	items := make([]AttestorUserOwnedDrydockNotePublicKeys, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAttestorUserOwnedDrydockNotePublicKeys(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandAttestorUserOwnedDrydockNotePublicKeys expands an instance of AttestorUserOwnedDrydockNotePublicKeys into a JSON
// request object.
func expandAttestorUserOwnedDrydockNotePublicKeys(c *Client, f *AttestorUserOwnedDrydockNotePublicKeys, res *Attestor) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Comment; !dcl.IsEmptyValueIndirect(v) {
		m["comment"] = v
	}
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.AsciiArmoredPgpPublicKey; !dcl.IsEmptyValueIndirect(v) {
		m["asciiArmoredPgpPublicKey"] = v
	}
	if v, err := expandAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(c, f.PkixPublicKey, res); err != nil {
		return nil, fmt.Errorf("error expanding PkixPublicKey into pkixPublicKey: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["pkixPublicKey"] = v
	}

	return m, nil
}

// flattenAttestorUserOwnedDrydockNotePublicKeys flattens an instance of AttestorUserOwnedDrydockNotePublicKeys from a JSON
// response object.
func flattenAttestorUserOwnedDrydockNotePublicKeys(c *Client, i interface{}, res *Attestor) *AttestorUserOwnedDrydockNotePublicKeys {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AttestorUserOwnedDrydockNotePublicKeys{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAttestorUserOwnedDrydockNotePublicKeys
	}
	r.Comment = dcl.FlattenString(m["comment"])
	r.Id = dcl.FlattenString(m["id"])
	r.AsciiArmoredPgpPublicKey = dcl.FlattenString(m["asciiArmoredPgpPublicKey"])
	r.PkixPublicKey = flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(c, m["pkixPublicKey"], res)

	return r
}

// expandAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyMap expands the contents of AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey into a JSON
// request object.
func expandAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyMap(c *Client, f map[string]AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey, res *Attestor) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySlice expands the contents of AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey into a JSON
// request object.
func expandAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySlice(c *Client, f []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey, res *Attestor) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyMap flattens the contents of AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey from a JSON
// response object.
func flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyMap(c *Client, i interface{}, res *Attestor) map[string]AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey{}
	}

	if len(a) == 0 {
		return map[string]AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey{}
	}

	items := make(map[string]AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey)
	for k, item := range a {
		items[k] = *flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySlice flattens the contents of AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey from a JSON
// response object.
func flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySlice(c *Client, i interface{}, res *Attestor) []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey {
	a, ok := i.([]interface{})
	if !ok {
		return []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey{}
	}

	if len(a) == 0 {
		return []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey{}
	}

	items := make([]AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey expands an instance of AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey into a JSON
// request object.
func expandAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(c *Client, f *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey, res *Attestor) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.PublicKeyPem; !dcl.IsEmptyValueIndirect(v) {
		m["publicKeyPem"] = v
	}
	if v := f.SignatureAlgorithm; !dcl.IsEmptyValueIndirect(v) {
		m["signatureAlgorithm"] = v
	}

	return m, nil
}

// flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey flattens an instance of AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey from a JSON
// response object.
func flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(c *Client, i interface{}, res *Attestor) *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey
	}
	r.PublicKeyPem = dcl.FlattenString(m["publicKeyPem"])
	r.SignatureAlgorithm = flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(m["signatureAlgorithm"])

	return r
}

// flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumMap flattens the contents of AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum from a JSON
// response object.
func flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumMap(c *Client, i interface{}, res *Attestor) map[string]AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum{}
	}

	if len(a) == 0 {
		return map[string]AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum{}
	}

	items := make(map[string]AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum)
	for k, item := range a {
		items[k] = *flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(item.(interface{}))
	}

	return items
}

// flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumSlice flattens the contents of AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum from a JSON
// response object.
func flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumSlice(c *Client, i interface{}, res *Attestor) []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum{}
	}

	if len(a) == 0 {
		return []AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum{}
	}

	items := make([]AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(item.(interface{})))
	}

	return items
}

// flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum asserts that an interface is a string, and returns a
// pointer to a *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum with the same value as that string.
func flattenAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(i interface{}) *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Attestor) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalAttestor(b, c, r)
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

type attestorDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         attestorApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToAttestorDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]attestorDiff, error) {
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
	var diffs []attestorDiff
	// For each operation name, create a attestorDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := attestorDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToAttestorApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToAttestorApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (attestorApiOperation, error) {
	switch opName {

	case "updateAttestorUpdateAttestorOperation":
		return &updateAttestorUpdateAttestorOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractAttestorFields(r *Attestor) error {
	vUserOwnedDrydockNote := r.UserOwnedDrydockNote
	if vUserOwnedDrydockNote == nil {
		// note: explicitly not the empty object.
		vUserOwnedDrydockNote = &AttestorUserOwnedDrydockNote{}
	}
	if err := extractAttestorUserOwnedDrydockNoteFields(r, vUserOwnedDrydockNote); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vUserOwnedDrydockNote) {
		r.UserOwnedDrydockNote = vUserOwnedDrydockNote
	}
	return nil
}
func extractAttestorUserOwnedDrydockNoteFields(r *Attestor, o *AttestorUserOwnedDrydockNote) error {
	return nil
}
func extractAttestorUserOwnedDrydockNotePublicKeysFields(r *Attestor, o *AttestorUserOwnedDrydockNotePublicKeys) error {
	vPkixPublicKey := o.PkixPublicKey
	if vPkixPublicKey == nil {
		// note: explicitly not the empty object.
		vPkixPublicKey = &AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey{}
	}
	if err := extractAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyFields(r, vPkixPublicKey); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPkixPublicKey) {
		o.PkixPublicKey = vPkixPublicKey
	}
	return nil
}
func extractAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyFields(r *Attestor, o *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) error {
	return nil
}

func postReadExtractAttestorFields(r *Attestor) error {
	vUserOwnedDrydockNote := r.UserOwnedDrydockNote
	if vUserOwnedDrydockNote == nil {
		// note: explicitly not the empty object.
		vUserOwnedDrydockNote = &AttestorUserOwnedDrydockNote{}
	}
	if err := postReadExtractAttestorUserOwnedDrydockNoteFields(r, vUserOwnedDrydockNote); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vUserOwnedDrydockNote) {
		r.UserOwnedDrydockNote = vUserOwnedDrydockNote
	}
	return nil
}
func postReadExtractAttestorUserOwnedDrydockNoteFields(r *Attestor, o *AttestorUserOwnedDrydockNote) error {
	return nil
}
func postReadExtractAttestorUserOwnedDrydockNotePublicKeysFields(r *Attestor, o *AttestorUserOwnedDrydockNotePublicKeys) error {
	vPkixPublicKey := o.PkixPublicKey
	if vPkixPublicKey == nil {
		// note: explicitly not the empty object.
		vPkixPublicKey = &AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey{}
	}
	if err := extractAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyFields(r, vPkixPublicKey); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPkixPublicKey) {
		o.PkixPublicKey = vPkixPublicKey
	}
	return nil
}
func postReadExtractAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyFields(r *Attestor, o *AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) error {
	return nil
}
