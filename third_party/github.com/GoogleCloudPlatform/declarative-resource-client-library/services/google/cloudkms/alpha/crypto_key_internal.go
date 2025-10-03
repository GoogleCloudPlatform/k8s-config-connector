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

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *CryptoKey) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "purpose"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.KeyRing, "KeyRing"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Primary) {
		if err := r.Primary.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.VersionTemplate) {
		if err := r.VersionTemplate.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CryptoKeyPrimary) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Attestation) {
		if err := r.Attestation.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ExternalProtectionLevelOptions) {
		if err := r.ExternalProtectionLevelOptions.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CryptoKeyPrimaryAttestation) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string(nil)); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.CertChains) {
		if err := r.CertChains.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CryptoKeyPrimaryAttestationCertChains) validate() error {
	return nil
}
func (r *CryptoKeyPrimaryExternalProtectionLevelOptions) validate() error {
	return nil
}
func (r *CryptoKeyVersionTemplate) validate() error {
	if err := dcl.Required(r, "algorithm"); err != nil {
		return err
	}
	return nil
}
func (r *CryptoKey) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://cloudkms.googleapis.com/v1/", params)
}

func (r *CryptoKey) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"keyRing":  dcl.ValueOrEmptyString(nr.KeyRing),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *CryptoKey) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"keyRing":  dcl.ValueOrEmptyString(nr.KeyRing),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys", nr.basePath(), userBasePath, params), nil

}

func (r *CryptoKey) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"keyRing":  dcl.ValueOrEmptyString(nr.KeyRing),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys?cryptoKeyId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *CryptoKey) SetPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"project":  *nr.Project,
		"location": *nr.Location,
		"keyRing":  *nr.KeyRing,
		"name":     *nr.Name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{name}}:setIamPolicy", nr.basePath(), userBasePath, fields)
}

func (r *CryptoKey) SetPolicyVerb() string {
	return "POST"
}

func (r *CryptoKey) getPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"project":  *nr.Project,
		"location": *nr.Location,
		"keyRing":  *nr.KeyRing,
		"name":     *nr.Name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{name}}:getIamPolicy", nr.basePath(), userBasePath, fields)
}

func (r *CryptoKey) IAMPolicyVersion() int {
	return 3
}

// cryptoKeyApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type cryptoKeyApiOperation interface {
	do(context.Context, *CryptoKey, *Client) error
}

// newUpdateCryptoKeyUpdateCryptoKeyRequest creates a request for an
// CryptoKey resource's UpdateCryptoKey update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateCryptoKeyUpdateCryptoKeyRequest(ctx context.Context, f *CryptoKey, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v, err := expandCryptoKeyPrimary(c, f.Primary, res); err != nil {
		return nil, fmt.Errorf("error expanding Primary into primary: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["primary"] = v
	}
	if v := f.NextRotationTime; !dcl.IsEmptyValueIndirect(v) {
		req["nextRotationTime"] = v
	}
	if v := f.RotationPeriod; !dcl.IsEmptyValueIndirect(v) {
		req["rotationPeriod"] = v
	}
	if v, err := expandCryptoKeyVersionTemplate(c, f.VersionTemplate, res); err != nil {
		return nil, fmt.Errorf("error expanding VersionTemplate into versionTemplate: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["versionTemplate"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	return req, nil
}

// marshalUpdateCryptoKeyUpdateCryptoKeyRequest converts the update into
// the final JSON request body.
func marshalUpdateCryptoKeyUpdateCryptoKeyRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	dcl.MoveMapEntry(
		m,
		[]string{"name"},
		[]string{},
	)
	return json.Marshal(m)
}

type updateCryptoKeyUpdateCryptoKeyOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateCryptoKeyUpdateCryptoKeyOperation) do(ctx context.Context, r *CryptoKey, c *Client) error {
	_, err := c.GetCryptoKey(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateCryptoKey")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateCryptoKeyUpdateCryptoKeyRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateCryptoKeyUpdateCryptoKeyRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listCryptoKeyRaw(ctx context.Context, r *CryptoKey, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != CryptoKeyMaxPage {
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

type listCryptoKeyOperation struct {
	CryptoKeys []map[string]interface{} `json:"cryptoKeys"`
	Token      string                   `json:"nextPageToken"`
}

func (c *Client) listCryptoKey(ctx context.Context, r *CryptoKey, pageToken string, pageSize int32) ([]*CryptoKey, string, error) {
	b, err := c.listCryptoKeyRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listCryptoKeyOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*CryptoKey
	for _, v := range m.CryptoKeys {
		res, err := unmarshalMapCryptoKey(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		res.KeyRing = r.KeyRing
		l = append(l, res)
	}

	return l, m.Token, nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createCryptoKeyOperation struct {
	response map[string]interface{}
}

func (op *createCryptoKeyOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createCryptoKeyOperation) do(ctx context.Context, r *CryptoKey, c *Client) error {
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

	if _, err := c.GetCryptoKey(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getCryptoKeyRaw(ctx context.Context, r *CryptoKey) ([]byte, error) {

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

func (c *Client) cryptoKeyDiffsForRawDesired(ctx context.Context, rawDesired *CryptoKey, opts ...dcl.ApplyOption) (initial, desired *CryptoKey, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *CryptoKey
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*CryptoKey); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected CryptoKey, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetCryptoKey(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a CryptoKey resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve CryptoKey resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that CryptoKey resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeCryptoKeyDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for CryptoKey: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for CryptoKey: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractCryptoKeyFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeCryptoKeyInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for CryptoKey: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeCryptoKeyDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for CryptoKey: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffCryptoKey(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeCryptoKeyInitialState(rawInitial, rawDesired *CryptoKey) (*CryptoKey, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeCryptoKeyDesiredState(rawDesired, rawInitial *CryptoKey, opts ...dcl.ApplyOption) (*CryptoKey, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Primary = canonicalizeCryptoKeyPrimary(rawDesired.Primary, nil, opts...)
		rawDesired.VersionTemplate = canonicalizeCryptoKeyVersionTemplate(rawDesired.VersionTemplate, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &CryptoKey{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.IsZeroValue(rawDesired.Purpose) || (dcl.IsEmptyValueIndirect(rawDesired.Purpose) && dcl.IsEmptyValueIndirect(rawInitial.Purpose)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Purpose = rawInitial.Purpose
	} else {
		canonicalDesired.Purpose = rawDesired.Purpose
	}
	if dcl.IsZeroValue(rawDesired.NextRotationTime) || (dcl.IsEmptyValueIndirect(rawDesired.NextRotationTime) && dcl.IsEmptyValueIndirect(rawInitial.NextRotationTime)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.NextRotationTime = rawInitial.NextRotationTime
	} else {
		canonicalDesired.NextRotationTime = rawDesired.NextRotationTime
	}
	if dcl.StringCanonicalize(rawDesired.RotationPeriod, rawInitial.RotationPeriod) {
		canonicalDesired.RotationPeriod = rawInitial.RotationPeriod
	} else {
		canonicalDesired.RotationPeriod = rawDesired.RotationPeriod
	}
	canonicalDesired.VersionTemplate = canonicalizeCryptoKeyVersionTemplate(rawDesired.VersionTemplate, rawInitial.VersionTemplate, opts...)
	if dcl.IsZeroValue(rawDesired.Labels) || (dcl.IsEmptyValueIndirect(rawDesired.Labels) && dcl.IsEmptyValueIndirect(rawInitial.Labels)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	if dcl.BoolCanonicalize(rawDesired.ImportOnly, rawInitial.ImportOnly) {
		canonicalDesired.ImportOnly = rawInitial.ImportOnly
	} else {
		canonicalDesired.ImportOnly = rawDesired.ImportOnly
	}
	if dcl.StringCanonicalize(rawDesired.DestroyScheduledDuration, rawInitial.DestroyScheduledDuration) {
		canonicalDesired.DestroyScheduledDuration = rawInitial.DestroyScheduledDuration
	} else {
		canonicalDesired.DestroyScheduledDuration = rawDesired.DestroyScheduledDuration
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
	if dcl.NameToSelfLink(rawDesired.KeyRing, rawInitial.KeyRing) {
		canonicalDesired.KeyRing = rawInitial.KeyRing
	} else {
		canonicalDesired.KeyRing = rawDesired.KeyRing
	}
	return canonicalDesired, nil
}

func canonicalizeCryptoKeyNewState(c *Client, rawNew, rawDesired *CryptoKey) (*CryptoKey, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Primary) && dcl.IsEmptyValueIndirect(rawDesired.Primary) {
		rawNew.Primary = rawDesired.Primary
	} else {
		rawNew.Primary = canonicalizeNewCryptoKeyPrimary(c, rawDesired.Primary, rawNew.Primary)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Purpose) && dcl.IsEmptyValueIndirect(rawDesired.Purpose) {
		rawNew.Purpose = rawDesired.Purpose
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.NextRotationTime) && dcl.IsEmptyValueIndirect(rawDesired.NextRotationTime) {
		rawNew.NextRotationTime = rawDesired.NextRotationTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.RotationPeriod) && dcl.IsEmptyValueIndirect(rawDesired.RotationPeriod) {
		rawNew.RotationPeriod = rawDesired.RotationPeriod
	} else {
		if dcl.StringCanonicalize(rawDesired.RotationPeriod, rawNew.RotationPeriod) {
			rawNew.RotationPeriod = rawDesired.RotationPeriod
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.VersionTemplate) && dcl.IsEmptyValueIndirect(rawDesired.VersionTemplate) {
		rawNew.VersionTemplate = rawDesired.VersionTemplate
	} else {
		rawNew.VersionTemplate = canonicalizeNewCryptoKeyVersionTemplate(c, rawDesired.VersionTemplate, rawNew.VersionTemplate)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ImportOnly) && dcl.IsEmptyValueIndirect(rawDesired.ImportOnly) {
		rawNew.ImportOnly = rawDesired.ImportOnly
	} else {
		if dcl.BoolCanonicalize(rawDesired.ImportOnly, rawNew.ImportOnly) {
			rawNew.ImportOnly = rawDesired.ImportOnly
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DestroyScheduledDuration) && dcl.IsEmptyValueIndirect(rawDesired.DestroyScheduledDuration) {
		rawNew.DestroyScheduledDuration = rawDesired.DestroyScheduledDuration
	} else {
		if dcl.StringCanonicalize(rawDesired.DestroyScheduledDuration, rawNew.DestroyScheduledDuration) {
			rawNew.DestroyScheduledDuration = rawDesired.DestroyScheduledDuration
		}
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	rawNew.KeyRing = rawDesired.KeyRing

	return rawNew, nil
}

func canonicalizeCryptoKeyPrimary(des, initial *CryptoKeyPrimary, opts ...dcl.ApplyOption) *CryptoKeyPrimary {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &CryptoKeyPrimary{}

	if dcl.IsZeroValue(des.Name) || (dcl.IsEmptyValueIndirect(des.Name) && dcl.IsEmptyValueIndirect(initial.Name)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.IsZeroValue(des.State) || (dcl.IsEmptyValueIndirect(des.State) && dcl.IsEmptyValueIndirect(initial.State)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.State = initial.State
	} else {
		cDes.State = des.State
	}
	cDes.ExternalProtectionLevelOptions = canonicalizeCryptoKeyPrimaryExternalProtectionLevelOptions(des.ExternalProtectionLevelOptions, initial.ExternalProtectionLevelOptions, opts...)

	return cDes
}

func canonicalizeCryptoKeyPrimarySlice(des, initial []CryptoKeyPrimary, opts ...dcl.ApplyOption) []CryptoKeyPrimary {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]CryptoKeyPrimary, 0, len(des))
		for _, d := range des {
			cd := canonicalizeCryptoKeyPrimary(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]CryptoKeyPrimary, 0, len(des))
	for i, d := range des {
		cd := canonicalizeCryptoKeyPrimary(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewCryptoKeyPrimary(c *Client, des, nw *CryptoKeyPrimary) *CryptoKeyPrimary {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for CryptoKeyPrimary while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Attestation = canonicalizeNewCryptoKeyPrimaryAttestation(c, des.Attestation, nw.Attestation)
	if dcl.StringCanonicalize(des.ImportJob, nw.ImportJob) {
		nw.ImportJob = des.ImportJob
	}
	if dcl.StringCanonicalize(des.ImportFailureReason, nw.ImportFailureReason) {
		nw.ImportFailureReason = des.ImportFailureReason
	}
	nw.ExternalProtectionLevelOptions = canonicalizeNewCryptoKeyPrimaryExternalProtectionLevelOptions(c, des.ExternalProtectionLevelOptions, nw.ExternalProtectionLevelOptions)
	if dcl.BoolCanonicalize(des.ReimportEligible, nw.ReimportEligible) {
		nw.ReimportEligible = des.ReimportEligible
	}

	return nw
}

func canonicalizeNewCryptoKeyPrimarySet(c *Client, des, nw []CryptoKeyPrimary) []CryptoKeyPrimary {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []CryptoKeyPrimary
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareCryptoKeyPrimaryNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewCryptoKeyPrimary(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewCryptoKeyPrimarySlice(c *Client, des, nw []CryptoKeyPrimary) []CryptoKeyPrimary {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CryptoKeyPrimary
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCryptoKeyPrimary(c, &d, &n))
	}

	return items
}

func canonicalizeCryptoKeyPrimaryAttestation(des, initial *CryptoKeyPrimaryAttestation, opts ...dcl.ApplyOption) *CryptoKeyPrimaryAttestation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &CryptoKeyPrimaryAttestation{}

	return cDes
}

func canonicalizeCryptoKeyPrimaryAttestationSlice(des, initial []CryptoKeyPrimaryAttestation, opts ...dcl.ApplyOption) []CryptoKeyPrimaryAttestation {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]CryptoKeyPrimaryAttestation, 0, len(des))
		for _, d := range des {
			cd := canonicalizeCryptoKeyPrimaryAttestation(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]CryptoKeyPrimaryAttestation, 0, len(des))
	for i, d := range des {
		cd := canonicalizeCryptoKeyPrimaryAttestation(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewCryptoKeyPrimaryAttestation(c *Client, des, nw *CryptoKeyPrimaryAttestation) *CryptoKeyPrimaryAttestation {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for CryptoKeyPrimaryAttestation while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Content, nw.Content) {
		nw.Content = des.Content
	}
	nw.CertChains = canonicalizeNewCryptoKeyPrimaryAttestationCertChains(c, des.CertChains, nw.CertChains)

	return nw
}

func canonicalizeNewCryptoKeyPrimaryAttestationSet(c *Client, des, nw []CryptoKeyPrimaryAttestation) []CryptoKeyPrimaryAttestation {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []CryptoKeyPrimaryAttestation
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareCryptoKeyPrimaryAttestationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewCryptoKeyPrimaryAttestation(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewCryptoKeyPrimaryAttestationSlice(c *Client, des, nw []CryptoKeyPrimaryAttestation) []CryptoKeyPrimaryAttestation {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CryptoKeyPrimaryAttestation
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCryptoKeyPrimaryAttestation(c, &d, &n))
	}

	return items
}

func canonicalizeCryptoKeyPrimaryAttestationCertChains(des, initial *CryptoKeyPrimaryAttestationCertChains, opts ...dcl.ApplyOption) *CryptoKeyPrimaryAttestationCertChains {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &CryptoKeyPrimaryAttestationCertChains{}

	if dcl.StringArrayCanonicalize(des.CaviumCerts, initial.CaviumCerts) {
		cDes.CaviumCerts = initial.CaviumCerts
	} else {
		cDes.CaviumCerts = des.CaviumCerts
	}
	if dcl.StringArrayCanonicalize(des.GoogleCardCerts, initial.GoogleCardCerts) {
		cDes.GoogleCardCerts = initial.GoogleCardCerts
	} else {
		cDes.GoogleCardCerts = des.GoogleCardCerts
	}
	if dcl.StringArrayCanonicalize(des.GooglePartitionCerts, initial.GooglePartitionCerts) {
		cDes.GooglePartitionCerts = initial.GooglePartitionCerts
	} else {
		cDes.GooglePartitionCerts = des.GooglePartitionCerts
	}

	return cDes
}

func canonicalizeCryptoKeyPrimaryAttestationCertChainsSlice(des, initial []CryptoKeyPrimaryAttestationCertChains, opts ...dcl.ApplyOption) []CryptoKeyPrimaryAttestationCertChains {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]CryptoKeyPrimaryAttestationCertChains, 0, len(des))
		for _, d := range des {
			cd := canonicalizeCryptoKeyPrimaryAttestationCertChains(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]CryptoKeyPrimaryAttestationCertChains, 0, len(des))
	for i, d := range des {
		cd := canonicalizeCryptoKeyPrimaryAttestationCertChains(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewCryptoKeyPrimaryAttestationCertChains(c *Client, des, nw *CryptoKeyPrimaryAttestationCertChains) *CryptoKeyPrimaryAttestationCertChains {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for CryptoKeyPrimaryAttestationCertChains while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.CaviumCerts, nw.CaviumCerts) {
		nw.CaviumCerts = des.CaviumCerts
	}
	if dcl.StringArrayCanonicalize(des.GoogleCardCerts, nw.GoogleCardCerts) {
		nw.GoogleCardCerts = des.GoogleCardCerts
	}
	if dcl.StringArrayCanonicalize(des.GooglePartitionCerts, nw.GooglePartitionCerts) {
		nw.GooglePartitionCerts = des.GooglePartitionCerts
	}

	return nw
}

func canonicalizeNewCryptoKeyPrimaryAttestationCertChainsSet(c *Client, des, nw []CryptoKeyPrimaryAttestationCertChains) []CryptoKeyPrimaryAttestationCertChains {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []CryptoKeyPrimaryAttestationCertChains
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareCryptoKeyPrimaryAttestationCertChainsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewCryptoKeyPrimaryAttestationCertChains(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewCryptoKeyPrimaryAttestationCertChainsSlice(c *Client, des, nw []CryptoKeyPrimaryAttestationCertChains) []CryptoKeyPrimaryAttestationCertChains {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CryptoKeyPrimaryAttestationCertChains
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCryptoKeyPrimaryAttestationCertChains(c, &d, &n))
	}

	return items
}

func canonicalizeCryptoKeyPrimaryExternalProtectionLevelOptions(des, initial *CryptoKeyPrimaryExternalProtectionLevelOptions, opts ...dcl.ApplyOption) *CryptoKeyPrimaryExternalProtectionLevelOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &CryptoKeyPrimaryExternalProtectionLevelOptions{}

	if dcl.StringCanonicalize(des.ExternalKeyUri, initial.ExternalKeyUri) || dcl.IsZeroValue(des.ExternalKeyUri) {
		cDes.ExternalKeyUri = initial.ExternalKeyUri
	} else {
		cDes.ExternalKeyUri = des.ExternalKeyUri
	}

	return cDes
}

func canonicalizeCryptoKeyPrimaryExternalProtectionLevelOptionsSlice(des, initial []CryptoKeyPrimaryExternalProtectionLevelOptions, opts ...dcl.ApplyOption) []CryptoKeyPrimaryExternalProtectionLevelOptions {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]CryptoKeyPrimaryExternalProtectionLevelOptions, 0, len(des))
		for _, d := range des {
			cd := canonicalizeCryptoKeyPrimaryExternalProtectionLevelOptions(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]CryptoKeyPrimaryExternalProtectionLevelOptions, 0, len(des))
	for i, d := range des {
		cd := canonicalizeCryptoKeyPrimaryExternalProtectionLevelOptions(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewCryptoKeyPrimaryExternalProtectionLevelOptions(c *Client, des, nw *CryptoKeyPrimaryExternalProtectionLevelOptions) *CryptoKeyPrimaryExternalProtectionLevelOptions {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for CryptoKeyPrimaryExternalProtectionLevelOptions while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.ExternalKeyUri, nw.ExternalKeyUri) {
		nw.ExternalKeyUri = des.ExternalKeyUri
	}

	return nw
}

func canonicalizeNewCryptoKeyPrimaryExternalProtectionLevelOptionsSet(c *Client, des, nw []CryptoKeyPrimaryExternalProtectionLevelOptions) []CryptoKeyPrimaryExternalProtectionLevelOptions {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []CryptoKeyPrimaryExternalProtectionLevelOptions
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareCryptoKeyPrimaryExternalProtectionLevelOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewCryptoKeyPrimaryExternalProtectionLevelOptions(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewCryptoKeyPrimaryExternalProtectionLevelOptionsSlice(c *Client, des, nw []CryptoKeyPrimaryExternalProtectionLevelOptions) []CryptoKeyPrimaryExternalProtectionLevelOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CryptoKeyPrimaryExternalProtectionLevelOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCryptoKeyPrimaryExternalProtectionLevelOptions(c, &d, &n))
	}

	return items
}

func canonicalizeCryptoKeyVersionTemplate(des, initial *CryptoKeyVersionTemplate, opts ...dcl.ApplyOption) *CryptoKeyVersionTemplate {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &CryptoKeyVersionTemplate{}

	if dcl.IsZeroValue(des.ProtectionLevel) || (dcl.IsEmptyValueIndirect(des.ProtectionLevel) && dcl.IsEmptyValueIndirect(initial.ProtectionLevel)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ProtectionLevel = initial.ProtectionLevel
	} else {
		cDes.ProtectionLevel = des.ProtectionLevel
	}
	if dcl.IsZeroValue(des.Algorithm) || (dcl.IsEmptyValueIndirect(des.Algorithm) && dcl.IsEmptyValueIndirect(initial.Algorithm)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Algorithm = initial.Algorithm
	} else {
		cDes.Algorithm = des.Algorithm
	}

	return cDes
}

func canonicalizeCryptoKeyVersionTemplateSlice(des, initial []CryptoKeyVersionTemplate, opts ...dcl.ApplyOption) []CryptoKeyVersionTemplate {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]CryptoKeyVersionTemplate, 0, len(des))
		for _, d := range des {
			cd := canonicalizeCryptoKeyVersionTemplate(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]CryptoKeyVersionTemplate, 0, len(des))
	for i, d := range des {
		cd := canonicalizeCryptoKeyVersionTemplate(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewCryptoKeyVersionTemplate(c *Client, des, nw *CryptoKeyVersionTemplate) *CryptoKeyVersionTemplate {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for CryptoKeyVersionTemplate while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewCryptoKeyVersionTemplateSet(c *Client, des, nw []CryptoKeyVersionTemplate) []CryptoKeyVersionTemplate {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []CryptoKeyVersionTemplate
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareCryptoKeyVersionTemplateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewCryptoKeyVersionTemplate(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewCryptoKeyVersionTemplateSlice(c *Client, des, nw []CryptoKeyVersionTemplate) []CryptoKeyVersionTemplate {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CryptoKeyVersionTemplate
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCryptoKeyVersionTemplate(c, &d, &n))
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
func diffCryptoKey(c *Client, desired, actual *CryptoKey, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Primary, actual.Primary, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareCryptoKeyPrimaryNewStyle, EmptyObject: EmptyCryptoKeyPrimary, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Primary")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Purpose, actual.Purpose, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Purpose")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.NextRotationTime, actual.NextRotationTime, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateCryptoKeyUpdateCryptoKeyOperation")}, fn.AddNest("NextRotationTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RotationPeriod, actual.RotationPeriod, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateCryptoKeyUpdateCryptoKeyOperation")}, fn.AddNest("RotationPeriod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VersionTemplate, actual.VersionTemplate, dcl.DiffInfo{ServerDefault: true, ObjectFunction: compareCryptoKeyVersionTemplateNewStyle, EmptyObject: EmptyCryptoKeyVersionTemplate, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("VersionTemplate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateCryptoKeyUpdateCryptoKeyOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ImportOnly, actual.ImportOnly, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ImportOnly")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DestroyScheduledDuration, actual.DestroyScheduledDuration, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DestroyScheduledDuration")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.KeyRing, actual.KeyRing, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KeyRing")); len(ds) != 0 || err != nil {
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
func compareCryptoKeyPrimaryNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CryptoKeyPrimary)
	if !ok {
		desiredNotPointer, ok := d.(CryptoKeyPrimary)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CryptoKeyPrimary or *CryptoKeyPrimary", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CryptoKeyPrimary)
	if !ok {
		actualNotPointer, ok := a.(CryptoKeyPrimary)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CryptoKeyPrimary", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateCryptoKeyUpdateCryptoKeyOperation")}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ProtectionLevel, actual.ProtectionLevel, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ProtectionLevel")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Algorithm, actual.Algorithm, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Algorithm")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Attestation, actual.Attestation, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareCryptoKeyPrimaryAttestationNewStyle, EmptyObject: EmptyCryptoKeyPrimaryAttestation, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Attestation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GenerateTime, actual.GenerateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GenerateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DestroyTime, actual.DestroyTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DestroyTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DestroyEventTime, actual.DestroyEventTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DestroyEventTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ImportJob, actual.ImportJob, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ImportJob")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ImportTime, actual.ImportTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ImportTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ImportFailureReason, actual.ImportFailureReason, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ImportFailureReason")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExternalProtectionLevelOptions, actual.ExternalProtectionLevelOptions, dcl.DiffInfo{ObjectFunction: compareCryptoKeyPrimaryExternalProtectionLevelOptionsNewStyle, EmptyObject: EmptyCryptoKeyPrimaryExternalProtectionLevelOptions, OperationSelector: dcl.TriggersOperation("updateCryptoKeyUpdateCryptoKeyOperation")}, fn.AddNest("ExternalProtectionLevelOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReimportEligible, actual.ReimportEligible, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ReimportEligible")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCryptoKeyPrimaryAttestationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CryptoKeyPrimaryAttestation)
	if !ok {
		desiredNotPointer, ok := d.(CryptoKeyPrimaryAttestation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CryptoKeyPrimaryAttestation or *CryptoKeyPrimaryAttestation", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CryptoKeyPrimaryAttestation)
	if !ok {
		actualNotPointer, ok := a.(CryptoKeyPrimaryAttestation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CryptoKeyPrimaryAttestation", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Format, actual.Format, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Format")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Content, actual.Content, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Content")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CertChains, actual.CertChains, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareCryptoKeyPrimaryAttestationCertChainsNewStyle, EmptyObject: EmptyCryptoKeyPrimaryAttestationCertChains, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CertChains")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCryptoKeyPrimaryAttestationCertChainsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CryptoKeyPrimaryAttestationCertChains)
	if !ok {
		desiredNotPointer, ok := d.(CryptoKeyPrimaryAttestationCertChains)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CryptoKeyPrimaryAttestationCertChains or *CryptoKeyPrimaryAttestationCertChains", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CryptoKeyPrimaryAttestationCertChains)
	if !ok {
		actualNotPointer, ok := a.(CryptoKeyPrimaryAttestationCertChains)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CryptoKeyPrimaryAttestationCertChains", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.CaviumCerts, actual.CaviumCerts, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateCryptoKeyUpdateCryptoKeyOperation")}, fn.AddNest("CaviumCerts")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GoogleCardCerts, actual.GoogleCardCerts, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateCryptoKeyUpdateCryptoKeyOperation")}, fn.AddNest("GoogleCardCerts")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GooglePartitionCerts, actual.GooglePartitionCerts, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateCryptoKeyUpdateCryptoKeyOperation")}, fn.AddNest("GooglePartitionCerts")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCryptoKeyPrimaryExternalProtectionLevelOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CryptoKeyPrimaryExternalProtectionLevelOptions)
	if !ok {
		desiredNotPointer, ok := d.(CryptoKeyPrimaryExternalProtectionLevelOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CryptoKeyPrimaryExternalProtectionLevelOptions or *CryptoKeyPrimaryExternalProtectionLevelOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CryptoKeyPrimaryExternalProtectionLevelOptions)
	if !ok {
		actualNotPointer, ok := a.(CryptoKeyPrimaryExternalProtectionLevelOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CryptoKeyPrimaryExternalProtectionLevelOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ExternalKeyUri, actual.ExternalKeyUri, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateCryptoKeyUpdateCryptoKeyOperation")}, fn.AddNest("ExternalKeyUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCryptoKeyVersionTemplateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CryptoKeyVersionTemplate)
	if !ok {
		desiredNotPointer, ok := d.(CryptoKeyVersionTemplate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CryptoKeyVersionTemplate or *CryptoKeyVersionTemplate", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CryptoKeyVersionTemplate)
	if !ok {
		actualNotPointer, ok := a.(CryptoKeyVersionTemplate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CryptoKeyVersionTemplate", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ProtectionLevel, actual.ProtectionLevel, dcl.DiffInfo{ServerDefault: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ProtectionLevel")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Algorithm, actual.Algorithm, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateCryptoKeyUpdateCryptoKeyOperation")}, fn.AddNest("Algorithm")); len(ds) != 0 || err != nil {
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
func (r *CryptoKey) urlNormalized() *CryptoKey {
	normalized := dcl.Copy(*r).(CryptoKey)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.RotationPeriod = dcl.SelfLinkToName(r.RotationPeriod)
	normalized.DestroyScheduledDuration = dcl.SelfLinkToName(r.DestroyScheduledDuration)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.KeyRing = dcl.SelfLinkToName(r.KeyRing)
	return &normalized
}

func (r *CryptoKey) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateCryptoKey" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"keyRing":  dcl.ValueOrEmptyString(nr.KeyRing),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the CryptoKey resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *CryptoKey) marshal(c *Client) ([]byte, error) {
	m, err := expandCryptoKey(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling CryptoKey: %w", err)
	}
	dcl.MoveMapEntry(
		m,
		[]string{"name"},
		[]string{},
	)

	return json.Marshal(m)
}

// unmarshalCryptoKey decodes JSON responses into the CryptoKey resource schema.
func unmarshalCryptoKey(b []byte, c *Client, res *CryptoKey) (*CryptoKey, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapCryptoKey(m, c, res)
}

func unmarshalMapCryptoKey(m map[string]interface{}, c *Client, res *CryptoKey) (*CryptoKey, error) {

	flattened := flattenCryptoKey(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandCryptoKey expands CryptoKey into a JSON request object.
func expandCryptoKey(c *Client, f *CryptoKey) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Location), dcl.SelfLinkToName(f.KeyRing), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Purpose; dcl.ValueShouldBeSent(v) {
		m["purpose"] = v
	}
	if v := f.NextRotationTime; dcl.ValueShouldBeSent(v) {
		m["nextRotationTime"] = v
	}
	if v := f.RotationPeriod; dcl.ValueShouldBeSent(v) {
		m["rotationPeriod"] = v
	}
	if v, err := expandCryptoKeyVersionTemplate(c, f.VersionTemplate, res); err != nil {
		return nil, fmt.Errorf("error expanding VersionTemplate into versionTemplate: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["versionTemplate"] = v
	}
	if v := f.Labels; dcl.ValueShouldBeSent(v) {
		m["labels"] = v
	}
	if v := f.ImportOnly; dcl.ValueShouldBeSent(v) {
		m["importOnly"] = v
	}
	if v := f.DestroyScheduledDuration; dcl.ValueShouldBeSent(v) {
		m["destroyScheduledDuration"] = v
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
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding KeyRing into keyRing: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["keyRing"] = v
	}

	return m, nil
}

// flattenCryptoKey flattens CryptoKey from a JSON request object into the
// CryptoKey type.
func flattenCryptoKey(c *Client, i interface{}, res *CryptoKey) *CryptoKey {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &CryptoKey{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.Primary = flattenCryptoKeyPrimary(c, m["primary"], res)
	resultRes.Purpose = flattenCryptoKeyPurposeEnum(m["purpose"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.NextRotationTime = dcl.FlattenString(m["nextRotationTime"])
	resultRes.RotationPeriod = dcl.FlattenString(m["rotationPeriod"])
	resultRes.VersionTemplate = flattenCryptoKeyVersionTemplate(c, m["versionTemplate"], res)
	resultRes.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	resultRes.ImportOnly = dcl.FlattenBool(m["importOnly"])
	resultRes.DestroyScheduledDuration = dcl.FlattenString(m["destroyScheduledDuration"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])
	resultRes.KeyRing = dcl.FlattenString(m["keyRing"])

	return resultRes
}

// expandCryptoKeyPrimaryMap expands the contents of CryptoKeyPrimary into a JSON
// request object.
func expandCryptoKeyPrimaryMap(c *Client, f map[string]CryptoKeyPrimary, res *CryptoKey) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCryptoKeyPrimary(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCryptoKeyPrimarySlice expands the contents of CryptoKeyPrimary into a JSON
// request object.
func expandCryptoKeyPrimarySlice(c *Client, f []CryptoKeyPrimary, res *CryptoKey) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCryptoKeyPrimary(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCryptoKeyPrimaryMap flattens the contents of CryptoKeyPrimary from a JSON
// response object.
func flattenCryptoKeyPrimaryMap(c *Client, i interface{}, res *CryptoKey) map[string]CryptoKeyPrimary {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CryptoKeyPrimary{}
	}

	if len(a) == 0 {
		return map[string]CryptoKeyPrimary{}
	}

	items := make(map[string]CryptoKeyPrimary)
	for k, item := range a {
		items[k] = *flattenCryptoKeyPrimary(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenCryptoKeyPrimarySlice flattens the contents of CryptoKeyPrimary from a JSON
// response object.
func flattenCryptoKeyPrimarySlice(c *Client, i interface{}, res *CryptoKey) []CryptoKeyPrimary {
	a, ok := i.([]interface{})
	if !ok {
		return []CryptoKeyPrimary{}
	}

	if len(a) == 0 {
		return []CryptoKeyPrimary{}
	}

	items := make([]CryptoKeyPrimary, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCryptoKeyPrimary(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandCryptoKeyPrimary expands an instance of CryptoKeyPrimary into a JSON
// request object.
func expandCryptoKeyPrimary(c *Client, f *CryptoKeyPrimary, res *CryptoKey) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v, err := expandCryptoKeyPrimaryExternalProtectionLevelOptions(c, f.ExternalProtectionLevelOptions, res); err != nil {
		return nil, fmt.Errorf("error expanding ExternalProtectionLevelOptions into externalProtectionLevelOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["externalProtectionLevelOptions"] = v
	}

	return m, nil
}

// flattenCryptoKeyPrimary flattens an instance of CryptoKeyPrimary from a JSON
// response object.
func flattenCryptoKeyPrimary(c *Client, i interface{}, res *CryptoKey) *CryptoKeyPrimary {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CryptoKeyPrimary{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCryptoKeyPrimary
	}
	r.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	r.State = flattenCryptoKeyPrimaryStateEnum(m["state"])
	r.ProtectionLevel = flattenCryptoKeyPrimaryProtectionLevelEnum(m["protectionLevel"])
	r.Algorithm = flattenCryptoKeyPrimaryAlgorithmEnum(m["algorithm"])
	r.Attestation = flattenCryptoKeyPrimaryAttestation(c, m["attestation"], res)
	r.CreateTime = dcl.FlattenString(m["createTime"])
	r.GenerateTime = dcl.FlattenString(m["generateTime"])
	r.DestroyTime = dcl.FlattenString(m["destroyTime"])
	r.DestroyEventTime = dcl.FlattenString(m["destroyEventTime"])
	r.ImportJob = dcl.FlattenString(m["importJob"])
	r.ImportTime = dcl.FlattenString(m["importTime"])
	r.ImportFailureReason = dcl.FlattenString(m["importFailureReason"])
	r.ExternalProtectionLevelOptions = flattenCryptoKeyPrimaryExternalProtectionLevelOptions(c, m["externalProtectionLevelOptions"], res)
	r.ReimportEligible = dcl.FlattenBool(m["reimportEligible"])

	return r
}

// expandCryptoKeyPrimaryAttestationMap expands the contents of CryptoKeyPrimaryAttestation into a JSON
// request object.
func expandCryptoKeyPrimaryAttestationMap(c *Client, f map[string]CryptoKeyPrimaryAttestation, res *CryptoKey) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCryptoKeyPrimaryAttestation(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCryptoKeyPrimaryAttestationSlice expands the contents of CryptoKeyPrimaryAttestation into a JSON
// request object.
func expandCryptoKeyPrimaryAttestationSlice(c *Client, f []CryptoKeyPrimaryAttestation, res *CryptoKey) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCryptoKeyPrimaryAttestation(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCryptoKeyPrimaryAttestationMap flattens the contents of CryptoKeyPrimaryAttestation from a JSON
// response object.
func flattenCryptoKeyPrimaryAttestationMap(c *Client, i interface{}, res *CryptoKey) map[string]CryptoKeyPrimaryAttestation {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CryptoKeyPrimaryAttestation{}
	}

	if len(a) == 0 {
		return map[string]CryptoKeyPrimaryAttestation{}
	}

	items := make(map[string]CryptoKeyPrimaryAttestation)
	for k, item := range a {
		items[k] = *flattenCryptoKeyPrimaryAttestation(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenCryptoKeyPrimaryAttestationSlice flattens the contents of CryptoKeyPrimaryAttestation from a JSON
// response object.
func flattenCryptoKeyPrimaryAttestationSlice(c *Client, i interface{}, res *CryptoKey) []CryptoKeyPrimaryAttestation {
	a, ok := i.([]interface{})
	if !ok {
		return []CryptoKeyPrimaryAttestation{}
	}

	if len(a) == 0 {
		return []CryptoKeyPrimaryAttestation{}
	}

	items := make([]CryptoKeyPrimaryAttestation, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCryptoKeyPrimaryAttestation(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandCryptoKeyPrimaryAttestation expands an instance of CryptoKeyPrimaryAttestation into a JSON
// request object.
func expandCryptoKeyPrimaryAttestation(c *Client, f *CryptoKeyPrimaryAttestation, res *CryptoKey) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenCryptoKeyPrimaryAttestation flattens an instance of CryptoKeyPrimaryAttestation from a JSON
// response object.
func flattenCryptoKeyPrimaryAttestation(c *Client, i interface{}, res *CryptoKey) *CryptoKeyPrimaryAttestation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CryptoKeyPrimaryAttestation{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCryptoKeyPrimaryAttestation
	}
	r.Format = flattenCryptoKeyPrimaryAttestationFormatEnum(m["format"])
	r.Content = dcl.FlattenString(m["content"])
	r.CertChains = flattenCryptoKeyPrimaryAttestationCertChains(c, m["certChains"], res)

	return r
}

// expandCryptoKeyPrimaryAttestationCertChainsMap expands the contents of CryptoKeyPrimaryAttestationCertChains into a JSON
// request object.
func expandCryptoKeyPrimaryAttestationCertChainsMap(c *Client, f map[string]CryptoKeyPrimaryAttestationCertChains, res *CryptoKey) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCryptoKeyPrimaryAttestationCertChains(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCryptoKeyPrimaryAttestationCertChainsSlice expands the contents of CryptoKeyPrimaryAttestationCertChains into a JSON
// request object.
func expandCryptoKeyPrimaryAttestationCertChainsSlice(c *Client, f []CryptoKeyPrimaryAttestationCertChains, res *CryptoKey) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCryptoKeyPrimaryAttestationCertChains(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCryptoKeyPrimaryAttestationCertChainsMap flattens the contents of CryptoKeyPrimaryAttestationCertChains from a JSON
// response object.
func flattenCryptoKeyPrimaryAttestationCertChainsMap(c *Client, i interface{}, res *CryptoKey) map[string]CryptoKeyPrimaryAttestationCertChains {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CryptoKeyPrimaryAttestationCertChains{}
	}

	if len(a) == 0 {
		return map[string]CryptoKeyPrimaryAttestationCertChains{}
	}

	items := make(map[string]CryptoKeyPrimaryAttestationCertChains)
	for k, item := range a {
		items[k] = *flattenCryptoKeyPrimaryAttestationCertChains(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenCryptoKeyPrimaryAttestationCertChainsSlice flattens the contents of CryptoKeyPrimaryAttestationCertChains from a JSON
// response object.
func flattenCryptoKeyPrimaryAttestationCertChainsSlice(c *Client, i interface{}, res *CryptoKey) []CryptoKeyPrimaryAttestationCertChains {
	a, ok := i.([]interface{})
	if !ok {
		return []CryptoKeyPrimaryAttestationCertChains{}
	}

	if len(a) == 0 {
		return []CryptoKeyPrimaryAttestationCertChains{}
	}

	items := make([]CryptoKeyPrimaryAttestationCertChains, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCryptoKeyPrimaryAttestationCertChains(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandCryptoKeyPrimaryAttestationCertChains expands an instance of CryptoKeyPrimaryAttestationCertChains into a JSON
// request object.
func expandCryptoKeyPrimaryAttestationCertChains(c *Client, f *CryptoKeyPrimaryAttestationCertChains, res *CryptoKey) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.CaviumCerts; v != nil {
		m["caviumCerts"] = v
	}
	if v := f.GoogleCardCerts; v != nil {
		m["googleCardCerts"] = v
	}
	if v := f.GooglePartitionCerts; v != nil {
		m["googlePartitionCerts"] = v
	}

	return m, nil
}

// flattenCryptoKeyPrimaryAttestationCertChains flattens an instance of CryptoKeyPrimaryAttestationCertChains from a JSON
// response object.
func flattenCryptoKeyPrimaryAttestationCertChains(c *Client, i interface{}, res *CryptoKey) *CryptoKeyPrimaryAttestationCertChains {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CryptoKeyPrimaryAttestationCertChains{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCryptoKeyPrimaryAttestationCertChains
	}
	r.CaviumCerts = dcl.FlattenStringSlice(m["caviumCerts"])
	r.GoogleCardCerts = dcl.FlattenStringSlice(m["googleCardCerts"])
	r.GooglePartitionCerts = dcl.FlattenStringSlice(m["googlePartitionCerts"])

	return r
}

// expandCryptoKeyPrimaryExternalProtectionLevelOptionsMap expands the contents of CryptoKeyPrimaryExternalProtectionLevelOptions into a JSON
// request object.
func expandCryptoKeyPrimaryExternalProtectionLevelOptionsMap(c *Client, f map[string]CryptoKeyPrimaryExternalProtectionLevelOptions, res *CryptoKey) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCryptoKeyPrimaryExternalProtectionLevelOptions(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCryptoKeyPrimaryExternalProtectionLevelOptionsSlice expands the contents of CryptoKeyPrimaryExternalProtectionLevelOptions into a JSON
// request object.
func expandCryptoKeyPrimaryExternalProtectionLevelOptionsSlice(c *Client, f []CryptoKeyPrimaryExternalProtectionLevelOptions, res *CryptoKey) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCryptoKeyPrimaryExternalProtectionLevelOptions(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCryptoKeyPrimaryExternalProtectionLevelOptionsMap flattens the contents of CryptoKeyPrimaryExternalProtectionLevelOptions from a JSON
// response object.
func flattenCryptoKeyPrimaryExternalProtectionLevelOptionsMap(c *Client, i interface{}, res *CryptoKey) map[string]CryptoKeyPrimaryExternalProtectionLevelOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CryptoKeyPrimaryExternalProtectionLevelOptions{}
	}

	if len(a) == 0 {
		return map[string]CryptoKeyPrimaryExternalProtectionLevelOptions{}
	}

	items := make(map[string]CryptoKeyPrimaryExternalProtectionLevelOptions)
	for k, item := range a {
		items[k] = *flattenCryptoKeyPrimaryExternalProtectionLevelOptions(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenCryptoKeyPrimaryExternalProtectionLevelOptionsSlice flattens the contents of CryptoKeyPrimaryExternalProtectionLevelOptions from a JSON
// response object.
func flattenCryptoKeyPrimaryExternalProtectionLevelOptionsSlice(c *Client, i interface{}, res *CryptoKey) []CryptoKeyPrimaryExternalProtectionLevelOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []CryptoKeyPrimaryExternalProtectionLevelOptions{}
	}

	if len(a) == 0 {
		return []CryptoKeyPrimaryExternalProtectionLevelOptions{}
	}

	items := make([]CryptoKeyPrimaryExternalProtectionLevelOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCryptoKeyPrimaryExternalProtectionLevelOptions(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandCryptoKeyPrimaryExternalProtectionLevelOptions expands an instance of CryptoKeyPrimaryExternalProtectionLevelOptions into a JSON
// request object.
func expandCryptoKeyPrimaryExternalProtectionLevelOptions(c *Client, f *CryptoKeyPrimaryExternalProtectionLevelOptions, res *CryptoKey) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ExternalKeyUri; !dcl.IsEmptyValueIndirect(v) {
		m["externalKeyUri"] = v
	}

	return m, nil
}

// flattenCryptoKeyPrimaryExternalProtectionLevelOptions flattens an instance of CryptoKeyPrimaryExternalProtectionLevelOptions from a JSON
// response object.
func flattenCryptoKeyPrimaryExternalProtectionLevelOptions(c *Client, i interface{}, res *CryptoKey) *CryptoKeyPrimaryExternalProtectionLevelOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CryptoKeyPrimaryExternalProtectionLevelOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCryptoKeyPrimaryExternalProtectionLevelOptions
	}
	r.ExternalKeyUri = dcl.FlattenString(m["externalKeyUri"])

	return r
}

// expandCryptoKeyVersionTemplateMap expands the contents of CryptoKeyVersionTemplate into a JSON
// request object.
func expandCryptoKeyVersionTemplateMap(c *Client, f map[string]CryptoKeyVersionTemplate, res *CryptoKey) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCryptoKeyVersionTemplate(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCryptoKeyVersionTemplateSlice expands the contents of CryptoKeyVersionTemplate into a JSON
// request object.
func expandCryptoKeyVersionTemplateSlice(c *Client, f []CryptoKeyVersionTemplate, res *CryptoKey) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCryptoKeyVersionTemplate(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCryptoKeyVersionTemplateMap flattens the contents of CryptoKeyVersionTemplate from a JSON
// response object.
func flattenCryptoKeyVersionTemplateMap(c *Client, i interface{}, res *CryptoKey) map[string]CryptoKeyVersionTemplate {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CryptoKeyVersionTemplate{}
	}

	if len(a) == 0 {
		return map[string]CryptoKeyVersionTemplate{}
	}

	items := make(map[string]CryptoKeyVersionTemplate)
	for k, item := range a {
		items[k] = *flattenCryptoKeyVersionTemplate(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenCryptoKeyVersionTemplateSlice flattens the contents of CryptoKeyVersionTemplate from a JSON
// response object.
func flattenCryptoKeyVersionTemplateSlice(c *Client, i interface{}, res *CryptoKey) []CryptoKeyVersionTemplate {
	a, ok := i.([]interface{})
	if !ok {
		return []CryptoKeyVersionTemplate{}
	}

	if len(a) == 0 {
		return []CryptoKeyVersionTemplate{}
	}

	items := make([]CryptoKeyVersionTemplate, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCryptoKeyVersionTemplate(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandCryptoKeyVersionTemplate expands an instance of CryptoKeyVersionTemplate into a JSON
// request object.
func expandCryptoKeyVersionTemplate(c *Client, f *CryptoKeyVersionTemplate, res *CryptoKey) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ProtectionLevel; !dcl.IsEmptyValueIndirect(v) {
		m["protectionLevel"] = v
	}
	if v := f.Algorithm; !dcl.IsEmptyValueIndirect(v) {
		m["algorithm"] = v
	}

	return m, nil
}

// flattenCryptoKeyVersionTemplate flattens an instance of CryptoKeyVersionTemplate from a JSON
// response object.
func flattenCryptoKeyVersionTemplate(c *Client, i interface{}, res *CryptoKey) *CryptoKeyVersionTemplate {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CryptoKeyVersionTemplate{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCryptoKeyVersionTemplate
	}
	r.ProtectionLevel = flattenCryptoKeyVersionTemplateProtectionLevelEnum(m["protectionLevel"])
	r.Algorithm = flattenCryptoKeyVersionTemplateAlgorithmEnum(m["algorithm"])

	return r
}

// flattenCryptoKeyPrimaryStateEnumMap flattens the contents of CryptoKeyPrimaryStateEnum from a JSON
// response object.
func flattenCryptoKeyPrimaryStateEnumMap(c *Client, i interface{}, res *CryptoKey) map[string]CryptoKeyPrimaryStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CryptoKeyPrimaryStateEnum{}
	}

	if len(a) == 0 {
		return map[string]CryptoKeyPrimaryStateEnum{}
	}

	items := make(map[string]CryptoKeyPrimaryStateEnum)
	for k, item := range a {
		items[k] = *flattenCryptoKeyPrimaryStateEnum(item.(interface{}))
	}

	return items
}

// flattenCryptoKeyPrimaryStateEnumSlice flattens the contents of CryptoKeyPrimaryStateEnum from a JSON
// response object.
func flattenCryptoKeyPrimaryStateEnumSlice(c *Client, i interface{}, res *CryptoKey) []CryptoKeyPrimaryStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []CryptoKeyPrimaryStateEnum{}
	}

	if len(a) == 0 {
		return []CryptoKeyPrimaryStateEnum{}
	}

	items := make([]CryptoKeyPrimaryStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCryptoKeyPrimaryStateEnum(item.(interface{})))
	}

	return items
}

// flattenCryptoKeyPrimaryStateEnum asserts that an interface is a string, and returns a
// pointer to a *CryptoKeyPrimaryStateEnum with the same value as that string.
func flattenCryptoKeyPrimaryStateEnum(i interface{}) *CryptoKeyPrimaryStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return CryptoKeyPrimaryStateEnumRef(s)
}

// flattenCryptoKeyPrimaryProtectionLevelEnumMap flattens the contents of CryptoKeyPrimaryProtectionLevelEnum from a JSON
// response object.
func flattenCryptoKeyPrimaryProtectionLevelEnumMap(c *Client, i interface{}, res *CryptoKey) map[string]CryptoKeyPrimaryProtectionLevelEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CryptoKeyPrimaryProtectionLevelEnum{}
	}

	if len(a) == 0 {
		return map[string]CryptoKeyPrimaryProtectionLevelEnum{}
	}

	items := make(map[string]CryptoKeyPrimaryProtectionLevelEnum)
	for k, item := range a {
		items[k] = *flattenCryptoKeyPrimaryProtectionLevelEnum(item.(interface{}))
	}

	return items
}

// flattenCryptoKeyPrimaryProtectionLevelEnumSlice flattens the contents of CryptoKeyPrimaryProtectionLevelEnum from a JSON
// response object.
func flattenCryptoKeyPrimaryProtectionLevelEnumSlice(c *Client, i interface{}, res *CryptoKey) []CryptoKeyPrimaryProtectionLevelEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []CryptoKeyPrimaryProtectionLevelEnum{}
	}

	if len(a) == 0 {
		return []CryptoKeyPrimaryProtectionLevelEnum{}
	}

	items := make([]CryptoKeyPrimaryProtectionLevelEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCryptoKeyPrimaryProtectionLevelEnum(item.(interface{})))
	}

	return items
}

// flattenCryptoKeyPrimaryProtectionLevelEnum asserts that an interface is a string, and returns a
// pointer to a *CryptoKeyPrimaryProtectionLevelEnum with the same value as that string.
func flattenCryptoKeyPrimaryProtectionLevelEnum(i interface{}) *CryptoKeyPrimaryProtectionLevelEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return CryptoKeyPrimaryProtectionLevelEnumRef(s)
}

// flattenCryptoKeyPrimaryAlgorithmEnumMap flattens the contents of CryptoKeyPrimaryAlgorithmEnum from a JSON
// response object.
func flattenCryptoKeyPrimaryAlgorithmEnumMap(c *Client, i interface{}, res *CryptoKey) map[string]CryptoKeyPrimaryAlgorithmEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CryptoKeyPrimaryAlgorithmEnum{}
	}

	if len(a) == 0 {
		return map[string]CryptoKeyPrimaryAlgorithmEnum{}
	}

	items := make(map[string]CryptoKeyPrimaryAlgorithmEnum)
	for k, item := range a {
		items[k] = *flattenCryptoKeyPrimaryAlgorithmEnum(item.(interface{}))
	}

	return items
}

// flattenCryptoKeyPrimaryAlgorithmEnumSlice flattens the contents of CryptoKeyPrimaryAlgorithmEnum from a JSON
// response object.
func flattenCryptoKeyPrimaryAlgorithmEnumSlice(c *Client, i interface{}, res *CryptoKey) []CryptoKeyPrimaryAlgorithmEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []CryptoKeyPrimaryAlgorithmEnum{}
	}

	if len(a) == 0 {
		return []CryptoKeyPrimaryAlgorithmEnum{}
	}

	items := make([]CryptoKeyPrimaryAlgorithmEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCryptoKeyPrimaryAlgorithmEnum(item.(interface{})))
	}

	return items
}

// flattenCryptoKeyPrimaryAlgorithmEnum asserts that an interface is a string, and returns a
// pointer to a *CryptoKeyPrimaryAlgorithmEnum with the same value as that string.
func flattenCryptoKeyPrimaryAlgorithmEnum(i interface{}) *CryptoKeyPrimaryAlgorithmEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return CryptoKeyPrimaryAlgorithmEnumRef(s)
}

// flattenCryptoKeyPrimaryAttestationFormatEnumMap flattens the contents of CryptoKeyPrimaryAttestationFormatEnum from a JSON
// response object.
func flattenCryptoKeyPrimaryAttestationFormatEnumMap(c *Client, i interface{}, res *CryptoKey) map[string]CryptoKeyPrimaryAttestationFormatEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CryptoKeyPrimaryAttestationFormatEnum{}
	}

	if len(a) == 0 {
		return map[string]CryptoKeyPrimaryAttestationFormatEnum{}
	}

	items := make(map[string]CryptoKeyPrimaryAttestationFormatEnum)
	for k, item := range a {
		items[k] = *flattenCryptoKeyPrimaryAttestationFormatEnum(item.(interface{}))
	}

	return items
}

// flattenCryptoKeyPrimaryAttestationFormatEnumSlice flattens the contents of CryptoKeyPrimaryAttestationFormatEnum from a JSON
// response object.
func flattenCryptoKeyPrimaryAttestationFormatEnumSlice(c *Client, i interface{}, res *CryptoKey) []CryptoKeyPrimaryAttestationFormatEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []CryptoKeyPrimaryAttestationFormatEnum{}
	}

	if len(a) == 0 {
		return []CryptoKeyPrimaryAttestationFormatEnum{}
	}

	items := make([]CryptoKeyPrimaryAttestationFormatEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCryptoKeyPrimaryAttestationFormatEnum(item.(interface{})))
	}

	return items
}

// flattenCryptoKeyPrimaryAttestationFormatEnum asserts that an interface is a string, and returns a
// pointer to a *CryptoKeyPrimaryAttestationFormatEnum with the same value as that string.
func flattenCryptoKeyPrimaryAttestationFormatEnum(i interface{}) *CryptoKeyPrimaryAttestationFormatEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return CryptoKeyPrimaryAttestationFormatEnumRef(s)
}

// flattenCryptoKeyPurposeEnumMap flattens the contents of CryptoKeyPurposeEnum from a JSON
// response object.
func flattenCryptoKeyPurposeEnumMap(c *Client, i interface{}, res *CryptoKey) map[string]CryptoKeyPurposeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CryptoKeyPurposeEnum{}
	}

	if len(a) == 0 {
		return map[string]CryptoKeyPurposeEnum{}
	}

	items := make(map[string]CryptoKeyPurposeEnum)
	for k, item := range a {
		items[k] = *flattenCryptoKeyPurposeEnum(item.(interface{}))
	}

	return items
}

// flattenCryptoKeyPurposeEnumSlice flattens the contents of CryptoKeyPurposeEnum from a JSON
// response object.
func flattenCryptoKeyPurposeEnumSlice(c *Client, i interface{}, res *CryptoKey) []CryptoKeyPurposeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []CryptoKeyPurposeEnum{}
	}

	if len(a) == 0 {
		return []CryptoKeyPurposeEnum{}
	}

	items := make([]CryptoKeyPurposeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCryptoKeyPurposeEnum(item.(interface{})))
	}

	return items
}

// flattenCryptoKeyPurposeEnum asserts that an interface is a string, and returns a
// pointer to a *CryptoKeyPurposeEnum with the same value as that string.
func flattenCryptoKeyPurposeEnum(i interface{}) *CryptoKeyPurposeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return CryptoKeyPurposeEnumRef(s)
}

// flattenCryptoKeyVersionTemplateProtectionLevelEnumMap flattens the contents of CryptoKeyVersionTemplateProtectionLevelEnum from a JSON
// response object.
func flattenCryptoKeyVersionTemplateProtectionLevelEnumMap(c *Client, i interface{}, res *CryptoKey) map[string]CryptoKeyVersionTemplateProtectionLevelEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CryptoKeyVersionTemplateProtectionLevelEnum{}
	}

	if len(a) == 0 {
		return map[string]CryptoKeyVersionTemplateProtectionLevelEnum{}
	}

	items := make(map[string]CryptoKeyVersionTemplateProtectionLevelEnum)
	for k, item := range a {
		items[k] = *flattenCryptoKeyVersionTemplateProtectionLevelEnum(item.(interface{}))
	}

	return items
}

// flattenCryptoKeyVersionTemplateProtectionLevelEnumSlice flattens the contents of CryptoKeyVersionTemplateProtectionLevelEnum from a JSON
// response object.
func flattenCryptoKeyVersionTemplateProtectionLevelEnumSlice(c *Client, i interface{}, res *CryptoKey) []CryptoKeyVersionTemplateProtectionLevelEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []CryptoKeyVersionTemplateProtectionLevelEnum{}
	}

	if len(a) == 0 {
		return []CryptoKeyVersionTemplateProtectionLevelEnum{}
	}

	items := make([]CryptoKeyVersionTemplateProtectionLevelEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCryptoKeyVersionTemplateProtectionLevelEnum(item.(interface{})))
	}

	return items
}

// flattenCryptoKeyVersionTemplateProtectionLevelEnum asserts that an interface is a string, and returns a
// pointer to a *CryptoKeyVersionTemplateProtectionLevelEnum with the same value as that string.
func flattenCryptoKeyVersionTemplateProtectionLevelEnum(i interface{}) *CryptoKeyVersionTemplateProtectionLevelEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return CryptoKeyVersionTemplateProtectionLevelEnumRef(s)
}

// flattenCryptoKeyVersionTemplateAlgorithmEnumMap flattens the contents of CryptoKeyVersionTemplateAlgorithmEnum from a JSON
// response object.
func flattenCryptoKeyVersionTemplateAlgorithmEnumMap(c *Client, i interface{}, res *CryptoKey) map[string]CryptoKeyVersionTemplateAlgorithmEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CryptoKeyVersionTemplateAlgorithmEnum{}
	}

	if len(a) == 0 {
		return map[string]CryptoKeyVersionTemplateAlgorithmEnum{}
	}

	items := make(map[string]CryptoKeyVersionTemplateAlgorithmEnum)
	for k, item := range a {
		items[k] = *flattenCryptoKeyVersionTemplateAlgorithmEnum(item.(interface{}))
	}

	return items
}

// flattenCryptoKeyVersionTemplateAlgorithmEnumSlice flattens the contents of CryptoKeyVersionTemplateAlgorithmEnum from a JSON
// response object.
func flattenCryptoKeyVersionTemplateAlgorithmEnumSlice(c *Client, i interface{}, res *CryptoKey) []CryptoKeyVersionTemplateAlgorithmEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []CryptoKeyVersionTemplateAlgorithmEnum{}
	}

	if len(a) == 0 {
		return []CryptoKeyVersionTemplateAlgorithmEnum{}
	}

	items := make([]CryptoKeyVersionTemplateAlgorithmEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCryptoKeyVersionTemplateAlgorithmEnum(item.(interface{})))
	}

	return items
}

// flattenCryptoKeyVersionTemplateAlgorithmEnum asserts that an interface is a string, and returns a
// pointer to a *CryptoKeyVersionTemplateAlgorithmEnum with the same value as that string.
func flattenCryptoKeyVersionTemplateAlgorithmEnum(i interface{}) *CryptoKeyVersionTemplateAlgorithmEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return CryptoKeyVersionTemplateAlgorithmEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *CryptoKey) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalCryptoKey(b, c, r)
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
		if nr.KeyRing == nil && ncr.KeyRing == nil {
			c.Config.Logger.Info("Both KeyRing fields null - considering equal.")
		} else if nr.KeyRing == nil || ncr.KeyRing == nil {
			c.Config.Logger.Info("Only one KeyRing field is null - considering unequal.")
			return false
		} else if *nr.KeyRing != *ncr.KeyRing {
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

type cryptoKeyDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         cryptoKeyApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToCryptoKeyDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]cryptoKeyDiff, error) {
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
	var diffs []cryptoKeyDiff
	// For each operation name, create a cryptoKeyDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := cryptoKeyDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToCryptoKeyApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToCryptoKeyApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (cryptoKeyApiOperation, error) {
	switch opName {

	case "updateCryptoKeyUpdateCryptoKeyOperation":
		return &updateCryptoKeyUpdateCryptoKeyOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractCryptoKeyFields(r *CryptoKey) error {
	vPrimary := r.Primary
	if vPrimary == nil {
		// note: explicitly not the empty object.
		vPrimary = &CryptoKeyPrimary{}
	}
	if err := extractCryptoKeyPrimaryFields(r, vPrimary); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPrimary) {
		r.Primary = vPrimary
	}
	vVersionTemplate := r.VersionTemplate
	if vVersionTemplate == nil {
		// note: explicitly not the empty object.
		vVersionTemplate = &CryptoKeyVersionTemplate{}
	}
	if err := extractCryptoKeyVersionTemplateFields(r, vVersionTemplate); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vVersionTemplate) {
		r.VersionTemplate = vVersionTemplate
	}
	return nil
}
func extractCryptoKeyPrimaryFields(r *CryptoKey, o *CryptoKeyPrimary) error {
	vAttestation := o.Attestation
	if vAttestation == nil {
		// note: explicitly not the empty object.
		vAttestation = &CryptoKeyPrimaryAttestation{}
	}
	if err := extractCryptoKeyPrimaryAttestationFields(r, vAttestation); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAttestation) {
		o.Attestation = vAttestation
	}
	vExternalProtectionLevelOptions := o.ExternalProtectionLevelOptions
	if vExternalProtectionLevelOptions == nil {
		// note: explicitly not the empty object.
		vExternalProtectionLevelOptions = &CryptoKeyPrimaryExternalProtectionLevelOptions{}
	}
	if err := extractCryptoKeyPrimaryExternalProtectionLevelOptionsFields(r, vExternalProtectionLevelOptions); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vExternalProtectionLevelOptions) {
		o.ExternalProtectionLevelOptions = vExternalProtectionLevelOptions
	}
	return nil
}
func extractCryptoKeyPrimaryAttestationFields(r *CryptoKey, o *CryptoKeyPrimaryAttestation) error {
	vCertChains := o.CertChains
	if vCertChains == nil {
		// note: explicitly not the empty object.
		vCertChains = &CryptoKeyPrimaryAttestationCertChains{}
	}
	if err := extractCryptoKeyPrimaryAttestationCertChainsFields(r, vCertChains); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCertChains) {
		o.CertChains = vCertChains
	}
	return nil
}
func extractCryptoKeyPrimaryAttestationCertChainsFields(r *CryptoKey, o *CryptoKeyPrimaryAttestationCertChains) error {
	return nil
}
func extractCryptoKeyPrimaryExternalProtectionLevelOptionsFields(r *CryptoKey, o *CryptoKeyPrimaryExternalProtectionLevelOptions) error {
	return nil
}
func extractCryptoKeyVersionTemplateFields(r *CryptoKey, o *CryptoKeyVersionTemplate) error {
	return nil
}

func postReadExtractCryptoKeyFields(r *CryptoKey) error {
	vPrimary := r.Primary
	if vPrimary == nil {
		// note: explicitly not the empty object.
		vPrimary = &CryptoKeyPrimary{}
	}
	if err := postReadExtractCryptoKeyPrimaryFields(r, vPrimary); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPrimary) {
		r.Primary = vPrimary
	}
	vVersionTemplate := r.VersionTemplate
	if vVersionTemplate == nil {
		// note: explicitly not the empty object.
		vVersionTemplate = &CryptoKeyVersionTemplate{}
	}
	if err := postReadExtractCryptoKeyVersionTemplateFields(r, vVersionTemplate); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vVersionTemplate) {
		r.VersionTemplate = vVersionTemplate
	}
	return nil
}
func postReadExtractCryptoKeyPrimaryFields(r *CryptoKey, o *CryptoKeyPrimary) error {
	vAttestation := o.Attestation
	if vAttestation == nil {
		// note: explicitly not the empty object.
		vAttestation = &CryptoKeyPrimaryAttestation{}
	}
	if err := extractCryptoKeyPrimaryAttestationFields(r, vAttestation); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAttestation) {
		o.Attestation = vAttestation
	}
	vExternalProtectionLevelOptions := o.ExternalProtectionLevelOptions
	if vExternalProtectionLevelOptions == nil {
		// note: explicitly not the empty object.
		vExternalProtectionLevelOptions = &CryptoKeyPrimaryExternalProtectionLevelOptions{}
	}
	if err := extractCryptoKeyPrimaryExternalProtectionLevelOptionsFields(r, vExternalProtectionLevelOptions); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vExternalProtectionLevelOptions) {
		o.ExternalProtectionLevelOptions = vExternalProtectionLevelOptions
	}
	return nil
}
func postReadExtractCryptoKeyPrimaryAttestationFields(r *CryptoKey, o *CryptoKeyPrimaryAttestation) error {
	vCertChains := o.CertChains
	if vCertChains == nil {
		// note: explicitly not the empty object.
		vCertChains = &CryptoKeyPrimaryAttestationCertChains{}
	}
	if err := extractCryptoKeyPrimaryAttestationCertChainsFields(r, vCertChains); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCertChains) {
		o.CertChains = vCertChains
	}
	return nil
}
func postReadExtractCryptoKeyPrimaryAttestationCertChainsFields(r *CryptoKey, o *CryptoKeyPrimaryAttestationCertChains) error {
	return nil
}
func postReadExtractCryptoKeyPrimaryExternalProtectionLevelOptionsFields(r *CryptoKey, o *CryptoKeyPrimaryExternalProtectionLevelOptions) error {
	return nil
}
func postReadExtractCryptoKeyVersionTemplateFields(r *CryptoKey, o *CryptoKeyVersionTemplate) error {
	return nil
}
