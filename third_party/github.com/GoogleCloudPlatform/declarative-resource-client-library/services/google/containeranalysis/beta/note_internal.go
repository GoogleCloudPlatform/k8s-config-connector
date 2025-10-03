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
)

func (r *Note) validate() error {

	if err := dcl.ValidateExactlyOneOfFieldsSet([]string{"Vulnerability", "Build", "Image", "Package", "Deployment", "Discovery", "Attestation"}, r.Vulnerability, r.Build, r.Image, r.Package, r.Deployment, r.Discovery, r.Attestation); err != nil {
		return err
	}
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Vulnerability) {
		if err := r.Vulnerability.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Build) {
		if err := r.Build.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Image) {
		if err := r.Image.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Package) {
		if err := r.Package.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Discovery) {
		if err := r.Discovery.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Deployment) {
		if err := r.Deployment.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Attestation) {
		if err := r.Attestation.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *NoteRelatedUrl) validate() error {
	return nil
}
func (r *NoteVulnerability) validate() error {
	if !dcl.IsEmptyValueIndirect(r.CvssV3) {
		if err := r.CvssV3.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *NoteVulnerabilityDetails) validate() error {
	if err := dcl.Required(r, "affectedCpeUri"); err != nil {
		return err
	}
	if err := dcl.Required(r, "affectedPackage"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.AffectedVersionStart) {
		if err := r.AffectedVersionStart.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.AffectedVersionEnd) {
		if err := r.AffectedVersionEnd.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.FixedVersion) {
		if err := r.FixedVersion.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *NoteVulnerabilityDetailsAffectedVersionStart) validate() error {
	if err := dcl.Required(r, "kind"); err != nil {
		return err
	}
	return nil
}
func (r *NoteVulnerabilityDetailsAffectedVersionEnd) validate() error {
	if err := dcl.Required(r, "kind"); err != nil {
		return err
	}
	return nil
}
func (r *NoteVulnerabilityDetailsFixedVersion) validate() error {
	if err := dcl.Required(r, "kind"); err != nil {
		return err
	}
	return nil
}
func (r *NoteVulnerabilityCvssV3) validate() error {
	return nil
}
func (r *NoteVulnerabilityWindowsDetails) validate() error {
	if err := dcl.Required(r, "cpeUri"); err != nil {
		return err
	}
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "fixingKbs"); err != nil {
		return err
	}
	return nil
}
func (r *NoteVulnerabilityWindowsDetailsFixingKbs) validate() error {
	return nil
}
func (r *NoteBuild) validate() error {
	if err := dcl.Required(r, "builderVersion"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Signature) {
		if err := r.Signature.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *NoteBuildSignature) validate() error {
	if err := dcl.Required(r, "signature"); err != nil {
		return err
	}
	return nil
}
func (r *NoteImage) validate() error {
	if err := dcl.Required(r, "resourceUrl"); err != nil {
		return err
	}
	if err := dcl.Required(r, "fingerprint"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Fingerprint) {
		if err := r.Fingerprint.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *NoteImageFingerprint) validate() error {
	if err := dcl.Required(r, "v1Name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "v2Blob"); err != nil {
		return err
	}
	return nil
}
func (r *NotePackage) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	return nil
}
func (r *NotePackageDistribution) validate() error {
	if err := dcl.Required(r, "cpeUri"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.LatestVersion) {
		if err := r.LatestVersion.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *NotePackageDistributionLatestVersion) validate() error {
	if err := dcl.Required(r, "kind"); err != nil {
		return err
	}
	return nil
}
func (r *NoteDiscovery) validate() error {
	if err := dcl.Required(r, "analysisKind"); err != nil {
		return err
	}
	return nil
}
func (r *NoteDeployment) validate() error {
	if err := dcl.Required(r, "resourceUri"); err != nil {
		return err
	}
	return nil
}
func (r *NoteAttestation) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Hint) {
		if err := r.Hint.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *NoteAttestationHint) validate() error {
	if err := dcl.Required(r, "humanReadableName"); err != nil {
		return err
	}
	return nil
}
func (r *Note) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://containeranalysis.googleapis.com/v1beta1/", params)
}

func (r *Note) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/notes/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Note) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.URL("projects/{{project}}/notes", nr.basePath(), userBasePath, params), nil

}

func (r *Note) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/notes?noteId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *Note) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/notes/{{name}}", nr.basePath(), userBasePath, params), nil
}

// noteApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type noteApiOperation interface {
	do(context.Context, *Note, *Client) error
}

// newUpdateNoteUpdateNoteRequest creates a request for an
// Note resource's UpdateNote update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateNoteUpdateNoteRequest(ctx context.Context, f *Note, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.ShortDescription; !dcl.IsEmptyValueIndirect(v) {
		req["shortDescription"] = v
	}
	if v := f.LongDescription; !dcl.IsEmptyValueIndirect(v) {
		req["longDescription"] = v
	}
	if v, err := expandNoteRelatedUrlSlice(c, f.RelatedUrl, res); err != nil {
		return nil, fmt.Errorf("error expanding RelatedUrl into relatedUrl: %w", err)
	} else if v != nil {
		req["relatedUrl"] = v
	}
	if v := f.ExpirationTime; !dcl.IsEmptyValueIndirect(v) {
		req["expirationTime"] = v
	}
	if v := f.RelatedNoteNames; v != nil {
		req["relatedNoteNames"] = v
	}
	if v, err := expandNoteVulnerability(c, f.Vulnerability, res); err != nil {
		return nil, fmt.Errorf("error expanding Vulnerability into vulnerability: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["vulnerability"] = v
	}
	if v, err := expandNoteBuild(c, f.Build, res); err != nil {
		return nil, fmt.Errorf("error expanding Build into build: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["build"] = v
	}
	if v, err := expandNoteImage(c, f.Image, res); err != nil {
		return nil, fmt.Errorf("error expanding Image into baseImage: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["baseImage"] = v
	}
	if v, err := expandNotePackage(c, f.Package, res); err != nil {
		return nil, fmt.Errorf("error expanding Package into package: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["package"] = v
	}
	if v, err := expandNoteDiscovery(c, f.Discovery, res); err != nil {
		return nil, fmt.Errorf("error expanding Discovery into discovery: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["discovery"] = v
	}
	if v, err := expandNoteDeployment(c, f.Deployment, res); err != nil {
		return nil, fmt.Errorf("error expanding Deployment into deployable: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["deployable"] = v
	}
	if v, err := expandNoteAttestation(c, f.Attestation, res); err != nil {
		return nil, fmt.Errorf("error expanding Attestation into attestationAuthority: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["attestationAuthority"] = v
	}
	return req, nil
}

// marshalUpdateNoteUpdateNoteRequest converts the update into
// the final JSON request body.
func marshalUpdateNoteUpdateNoteRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateNoteUpdateNoteOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateNoteUpdateNoteOperation) do(ctx context.Context, r *Note, c *Client) error {
	_, err := c.GetNote(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateNote")
	if err != nil {
		return err
	}

	req, err := newUpdateNoteUpdateNoteRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateNoteUpdateNoteRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listNoteRaw(ctx context.Context, r *Note, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != NoteMaxPage {
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

type listNoteOperation struct {
	Notes []map[string]interface{} `json:"notes"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listNote(ctx context.Context, r *Note, pageToken string, pageSize int32) ([]*Note, string, error) {
	b, err := c.listNoteRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listNoteOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Note
	for _, v := range m.Notes {
		res, err := unmarshalMapNote(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllNote(ctx context.Context, f func(*Note) bool, resources []*Note) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteNote(ctx, res)
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

type deleteNoteOperation struct{}

func (op *deleteNoteOperation) do(ctx context.Context, r *Note, c *Client) error {
	r, err := c.GetNote(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Note not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetNote checking for existence. error: %v", err)
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
		return fmt.Errorf("failed to delete Note: %w", err)
	}

	// We saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// This is the reason we are adding retry to handle that case.
	retriesRemaining := 10
	dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		_, err := c.GetNote(ctx, r)
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
type createNoteOperation struct {
	response map[string]interface{}
}

func (op *createNoteOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createNoteOperation) do(ctx context.Context, r *Note, c *Client) error {
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

	if _, err := c.GetNote(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getNoteRaw(ctx context.Context, r *Note) ([]byte, error) {

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

func (c *Client) noteDiffsForRawDesired(ctx context.Context, rawDesired *Note, opts ...dcl.ApplyOption) (initial, desired *Note, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Note
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Note); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Note, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetNote(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Note resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Note resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Note resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeNoteDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Note: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Note: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractNoteFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeNoteInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Note: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeNoteDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Note: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffNote(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeNoteInitialState(rawInitial, rawDesired *Note) (*Note, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.

	if !dcl.IsZeroValue(rawInitial.Vulnerability) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Build, rawInitial.Image, rawInitial.Package, rawInitial.Deployment, rawInitial.Discovery, rawInitial.Attestation) {
			rawInitial.Vulnerability = EmptyNoteVulnerability
		}
	}

	if !dcl.IsZeroValue(rawInitial.Build) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Vulnerability, rawInitial.Image, rawInitial.Package, rawInitial.Deployment, rawInitial.Discovery, rawInitial.Attestation) {
			rawInitial.Build = EmptyNoteBuild
		}
	}

	if !dcl.IsZeroValue(rawInitial.Image) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Vulnerability, rawInitial.Build, rawInitial.Package, rawInitial.Deployment, rawInitial.Discovery, rawInitial.Attestation) {
			rawInitial.Image = EmptyNoteImage
		}
	}

	if !dcl.IsZeroValue(rawInitial.Package) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Vulnerability, rawInitial.Build, rawInitial.Image, rawInitial.Deployment, rawInitial.Discovery, rawInitial.Attestation) {
			rawInitial.Package = EmptyNotePackage
		}
	}

	if !dcl.IsZeroValue(rawInitial.Deployment) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Vulnerability, rawInitial.Build, rawInitial.Image, rawInitial.Package, rawInitial.Discovery, rawInitial.Attestation) {
			rawInitial.Deployment = EmptyNoteDeployment
		}
	}

	if !dcl.IsZeroValue(rawInitial.Discovery) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Vulnerability, rawInitial.Build, rawInitial.Image, rawInitial.Package, rawInitial.Deployment, rawInitial.Attestation) {
			rawInitial.Discovery = EmptyNoteDiscovery
		}
	}

	if !dcl.IsZeroValue(rawInitial.Attestation) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Vulnerability, rawInitial.Build, rawInitial.Image, rawInitial.Package, rawInitial.Deployment, rawInitial.Discovery) {
			rawInitial.Attestation = EmptyNoteAttestation
		}
	}

	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeNoteDesiredState(rawDesired, rawInitial *Note, opts ...dcl.ApplyOption) (*Note, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Vulnerability = canonicalizeNoteVulnerability(rawDesired.Vulnerability, nil, opts...)
		rawDesired.Build = canonicalizeNoteBuild(rawDesired.Build, nil, opts...)
		rawDesired.Image = canonicalizeNoteImage(rawDesired.Image, nil, opts...)
		rawDesired.Package = canonicalizeNotePackage(rawDesired.Package, nil, opts...)
		rawDesired.Discovery = canonicalizeNoteDiscovery(rawDesired.Discovery, nil, opts...)
		rawDesired.Deployment = canonicalizeNoteDeployment(rawDesired.Deployment, nil, opts...)
		rawDesired.Attestation = canonicalizeNoteAttestation(rawDesired.Attestation, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Note{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.ShortDescription, rawInitial.ShortDescription) {
		canonicalDesired.ShortDescription = rawInitial.ShortDescription
	} else {
		canonicalDesired.ShortDescription = rawDesired.ShortDescription
	}
	if dcl.StringCanonicalize(rawDesired.LongDescription, rawInitial.LongDescription) {
		canonicalDesired.LongDescription = rawInitial.LongDescription
	} else {
		canonicalDesired.LongDescription = rawDesired.LongDescription
	}
	canonicalDesired.RelatedUrl = canonicalizeNoteRelatedUrlSlice(rawDesired.RelatedUrl, rawInitial.RelatedUrl, opts...)
	if dcl.IsZeroValue(rawDesired.ExpirationTime) || (dcl.IsEmptyValueIndirect(rawDesired.ExpirationTime) && dcl.IsEmptyValueIndirect(rawInitial.ExpirationTime)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.ExpirationTime = rawInitial.ExpirationTime
	} else {
		canonicalDesired.ExpirationTime = rawDesired.ExpirationTime
	}
	if dcl.StringArrayCanonicalize(rawDesired.RelatedNoteNames, rawInitial.RelatedNoteNames) {
		canonicalDesired.RelatedNoteNames = rawInitial.RelatedNoteNames
	} else {
		canonicalDesired.RelatedNoteNames = rawDesired.RelatedNoteNames
	}
	canonicalDesired.Vulnerability = canonicalizeNoteVulnerability(rawDesired.Vulnerability, rawInitial.Vulnerability, opts...)
	canonicalDesired.Build = canonicalizeNoteBuild(rawDesired.Build, rawInitial.Build, opts...)
	canonicalDesired.Image = canonicalizeNoteImage(rawDesired.Image, rawInitial.Image, opts...)
	canonicalDesired.Package = canonicalizeNotePackage(rawDesired.Package, rawInitial.Package, opts...)
	canonicalDesired.Discovery = canonicalizeNoteDiscovery(rawDesired.Discovery, rawInitial.Discovery, opts...)
	canonicalDesired.Deployment = canonicalizeNoteDeployment(rawDesired.Deployment, rawInitial.Deployment, opts...)
	canonicalDesired.Attestation = canonicalizeNoteAttestation(rawDesired.Attestation, rawInitial.Attestation, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}

	if canonicalDesired.Vulnerability != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Build, rawDesired.Image, rawDesired.Package, rawDesired.Deployment, rawDesired.Discovery, rawDesired.Attestation) {
			canonicalDesired.Vulnerability = EmptyNoteVulnerability
		}
	}

	if canonicalDesired.Build != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Vulnerability, rawDesired.Image, rawDesired.Package, rawDesired.Deployment, rawDesired.Discovery, rawDesired.Attestation) {
			canonicalDesired.Build = EmptyNoteBuild
		}
	}

	if canonicalDesired.Image != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Vulnerability, rawDesired.Build, rawDesired.Package, rawDesired.Deployment, rawDesired.Discovery, rawDesired.Attestation) {
			canonicalDesired.Image = EmptyNoteImage
		}
	}

	if canonicalDesired.Package != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Vulnerability, rawDesired.Build, rawDesired.Image, rawDesired.Deployment, rawDesired.Discovery, rawDesired.Attestation) {
			canonicalDesired.Package = EmptyNotePackage
		}
	}

	if canonicalDesired.Deployment != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Vulnerability, rawDesired.Build, rawDesired.Image, rawDesired.Package, rawDesired.Discovery, rawDesired.Attestation) {
			canonicalDesired.Deployment = EmptyNoteDeployment
		}
	}

	if canonicalDesired.Discovery != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Vulnerability, rawDesired.Build, rawDesired.Image, rawDesired.Package, rawDesired.Deployment, rawDesired.Attestation) {
			canonicalDesired.Discovery = EmptyNoteDiscovery
		}
	}

	if canonicalDesired.Attestation != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Vulnerability, rawDesired.Build, rawDesired.Image, rawDesired.Package, rawDesired.Deployment, rawDesired.Discovery) {
			canonicalDesired.Attestation = EmptyNoteAttestation
		}
	}

	return canonicalDesired, nil
}

func canonicalizeNoteNewState(c *Client, rawNew, rawDesired *Note) (*Note, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ShortDescription) && dcl.IsEmptyValueIndirect(rawDesired.ShortDescription) {
		rawNew.ShortDescription = rawDesired.ShortDescription
	} else {
		if dcl.StringCanonicalize(rawDesired.ShortDescription, rawNew.ShortDescription) {
			rawNew.ShortDescription = rawDesired.ShortDescription
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.LongDescription) && dcl.IsEmptyValueIndirect(rawDesired.LongDescription) {
		rawNew.LongDescription = rawDesired.LongDescription
	} else {
		if dcl.StringCanonicalize(rawDesired.LongDescription, rawNew.LongDescription) {
			rawNew.LongDescription = rawDesired.LongDescription
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.RelatedUrl) && dcl.IsEmptyValueIndirect(rawDesired.RelatedUrl) {
		rawNew.RelatedUrl = rawDesired.RelatedUrl
	} else {
		rawNew.RelatedUrl = canonicalizeNewNoteRelatedUrlSlice(c, rawDesired.RelatedUrl, rawNew.RelatedUrl)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ExpirationTime) && dcl.IsEmptyValueIndirect(rawDesired.ExpirationTime) {
		rawNew.ExpirationTime = rawDesired.ExpirationTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.RelatedNoteNames) && dcl.IsEmptyValueIndirect(rawDesired.RelatedNoteNames) {
		rawNew.RelatedNoteNames = rawDesired.RelatedNoteNames
	} else {
		if dcl.StringArrayCanonicalize(rawDesired.RelatedNoteNames, rawNew.RelatedNoteNames) {
			rawNew.RelatedNoteNames = rawDesired.RelatedNoteNames
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Vulnerability) && dcl.IsEmptyValueIndirect(rawDesired.Vulnerability) {
		rawNew.Vulnerability = rawDesired.Vulnerability
	} else {
		rawNew.Vulnerability = canonicalizeNewNoteVulnerability(c, rawDesired.Vulnerability, rawNew.Vulnerability)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Build) && dcl.IsEmptyValueIndirect(rawDesired.Build) {
		rawNew.Build = rawDesired.Build
	} else {
		rawNew.Build = canonicalizeNewNoteBuild(c, rawDesired.Build, rawNew.Build)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Image) && dcl.IsEmptyValueIndirect(rawDesired.Image) {
		rawNew.Image = rawDesired.Image
	} else {
		rawNew.Image = canonicalizeNewNoteImage(c, rawDesired.Image, rawNew.Image)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Package) && dcl.IsEmptyValueIndirect(rawDesired.Package) {
		rawNew.Package = rawDesired.Package
	} else {
		rawNew.Package = canonicalizeNewNotePackage(c, rawDesired.Package, rawNew.Package)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Discovery) && dcl.IsEmptyValueIndirect(rawDesired.Discovery) {
		rawNew.Discovery = rawDesired.Discovery
	} else {
		rawNew.Discovery = canonicalizeNewNoteDiscovery(c, rawDesired.Discovery, rawNew.Discovery)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Deployment) && dcl.IsEmptyValueIndirect(rawDesired.Deployment) {
		rawNew.Deployment = rawDesired.Deployment
	} else {
		rawNew.Deployment = canonicalizeNewNoteDeployment(c, rawDesired.Deployment, rawNew.Deployment)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Attestation) && dcl.IsEmptyValueIndirect(rawDesired.Attestation) {
		rawNew.Attestation = rawDesired.Attestation
	} else {
		rawNew.Attestation = canonicalizeNewNoteAttestation(c, rawDesired.Attestation, rawNew.Attestation)
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeNoteRelatedUrl(des, initial *NoteRelatedUrl, opts ...dcl.ApplyOption) *NoteRelatedUrl {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NoteRelatedUrl{}

	if dcl.StringCanonicalize(des.Url, initial.Url) || dcl.IsZeroValue(des.Url) {
		cDes.Url = initial.Url
	} else {
		cDes.Url = des.Url
	}
	if dcl.StringCanonicalize(des.Label, initial.Label) || dcl.IsZeroValue(des.Label) {
		cDes.Label = initial.Label
	} else {
		cDes.Label = des.Label
	}

	return cDes
}

func canonicalizeNoteRelatedUrlSlice(des, initial []NoteRelatedUrl, opts ...dcl.ApplyOption) []NoteRelatedUrl {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NoteRelatedUrl, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNoteRelatedUrl(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NoteRelatedUrl, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNoteRelatedUrl(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNoteRelatedUrl(c *Client, des, nw *NoteRelatedUrl) *NoteRelatedUrl {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for NoteRelatedUrl while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Url, nw.Url) {
		nw.Url = des.Url
	}
	if dcl.StringCanonicalize(des.Label, nw.Label) {
		nw.Label = des.Label
	}

	return nw
}

func canonicalizeNewNoteRelatedUrlSet(c *Client, des, nw []NoteRelatedUrl) []NoteRelatedUrl {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []NoteRelatedUrl
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareNoteRelatedUrlNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewNoteRelatedUrl(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewNoteRelatedUrlSlice(c *Client, des, nw []NoteRelatedUrl) []NoteRelatedUrl {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NoteRelatedUrl
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteRelatedUrl(c, &d, &n))
	}

	return items
}

func canonicalizeNoteVulnerability(des, initial *NoteVulnerability, opts ...dcl.ApplyOption) *NoteVulnerability {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NoteVulnerability{}

	if dcl.IsZeroValue(des.CvssScore) || (dcl.IsEmptyValueIndirect(des.CvssScore) && dcl.IsEmptyValueIndirect(initial.CvssScore)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.CvssScore = initial.CvssScore
	} else {
		cDes.CvssScore = des.CvssScore
	}
	if dcl.IsZeroValue(des.Severity) || (dcl.IsEmptyValueIndirect(des.Severity) && dcl.IsEmptyValueIndirect(initial.Severity)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Severity = initial.Severity
	} else {
		cDes.Severity = des.Severity
	}
	cDes.Details = canonicalizeNoteVulnerabilityDetailsSlice(des.Details, initial.Details, opts...)
	cDes.CvssV3 = canonicalizeNoteVulnerabilityCvssV3(des.CvssV3, initial.CvssV3, opts...)
	cDes.WindowsDetails = canonicalizeNoteVulnerabilityWindowsDetailsSlice(des.WindowsDetails, initial.WindowsDetails, opts...)
	if dcl.IsZeroValue(des.SourceUpdateTime) || (dcl.IsEmptyValueIndirect(des.SourceUpdateTime) && dcl.IsEmptyValueIndirect(initial.SourceUpdateTime)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.SourceUpdateTime = initial.SourceUpdateTime
	} else {
		cDes.SourceUpdateTime = des.SourceUpdateTime
	}

	return cDes
}

func canonicalizeNoteVulnerabilitySlice(des, initial []NoteVulnerability, opts ...dcl.ApplyOption) []NoteVulnerability {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NoteVulnerability, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNoteVulnerability(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NoteVulnerability, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNoteVulnerability(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNoteVulnerability(c *Client, des, nw *NoteVulnerability) *NoteVulnerability {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for NoteVulnerability while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Details = canonicalizeNewNoteVulnerabilityDetailsSlice(c, des.Details, nw.Details)
	nw.CvssV3 = canonicalizeNewNoteVulnerabilityCvssV3(c, des.CvssV3, nw.CvssV3)
	nw.WindowsDetails = canonicalizeNewNoteVulnerabilityWindowsDetailsSlice(c, des.WindowsDetails, nw.WindowsDetails)

	return nw
}

func canonicalizeNewNoteVulnerabilitySet(c *Client, des, nw []NoteVulnerability) []NoteVulnerability {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []NoteVulnerability
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareNoteVulnerabilityNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewNoteVulnerability(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewNoteVulnerabilitySlice(c *Client, des, nw []NoteVulnerability) []NoteVulnerability {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NoteVulnerability
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteVulnerability(c, &d, &n))
	}

	return items
}

func canonicalizeNoteVulnerabilityDetails(des, initial *NoteVulnerabilityDetails, opts ...dcl.ApplyOption) *NoteVulnerabilityDetails {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NoteVulnerabilityDetails{}

	if dcl.StringCanonicalize(des.SeverityName, initial.SeverityName) || dcl.IsZeroValue(des.SeverityName) {
		cDes.SeverityName = initial.SeverityName
	} else {
		cDes.SeverityName = des.SeverityName
	}
	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		cDes.Description = initial.Description
	} else {
		cDes.Description = des.Description
	}
	if dcl.StringCanonicalize(des.PackageType, initial.PackageType) || dcl.IsZeroValue(des.PackageType) {
		cDes.PackageType = initial.PackageType
	} else {
		cDes.PackageType = des.PackageType
	}
	if dcl.StringCanonicalize(des.AffectedCpeUri, initial.AffectedCpeUri) || dcl.IsZeroValue(des.AffectedCpeUri) {
		cDes.AffectedCpeUri = initial.AffectedCpeUri
	} else {
		cDes.AffectedCpeUri = des.AffectedCpeUri
	}
	if dcl.StringCanonicalize(des.AffectedPackage, initial.AffectedPackage) || dcl.IsZeroValue(des.AffectedPackage) {
		cDes.AffectedPackage = initial.AffectedPackage
	} else {
		cDes.AffectedPackage = des.AffectedPackage
	}
	cDes.AffectedVersionStart = canonicalizeNoteVulnerabilityDetailsAffectedVersionStart(des.AffectedVersionStart, initial.AffectedVersionStart, opts...)
	cDes.AffectedVersionEnd = canonicalizeNoteVulnerabilityDetailsAffectedVersionEnd(des.AffectedVersionEnd, initial.AffectedVersionEnd, opts...)
	if dcl.StringCanonicalize(des.FixedCpeUri, initial.FixedCpeUri) || dcl.IsZeroValue(des.FixedCpeUri) {
		cDes.FixedCpeUri = initial.FixedCpeUri
	} else {
		cDes.FixedCpeUri = des.FixedCpeUri
	}
	if dcl.StringCanonicalize(des.FixedPackage, initial.FixedPackage) || dcl.IsZeroValue(des.FixedPackage) {
		cDes.FixedPackage = initial.FixedPackage
	} else {
		cDes.FixedPackage = des.FixedPackage
	}
	cDes.FixedVersion = canonicalizeNoteVulnerabilityDetailsFixedVersion(des.FixedVersion, initial.FixedVersion, opts...)
	if dcl.BoolCanonicalize(des.IsObsolete, initial.IsObsolete) || dcl.IsZeroValue(des.IsObsolete) {
		cDes.IsObsolete = initial.IsObsolete
	} else {
		cDes.IsObsolete = des.IsObsolete
	}
	if dcl.IsZeroValue(des.SourceUpdateTime) || (dcl.IsEmptyValueIndirect(des.SourceUpdateTime) && dcl.IsEmptyValueIndirect(initial.SourceUpdateTime)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.SourceUpdateTime = initial.SourceUpdateTime
	} else {
		cDes.SourceUpdateTime = des.SourceUpdateTime
	}

	return cDes
}

func canonicalizeNoteVulnerabilityDetailsSlice(des, initial []NoteVulnerabilityDetails, opts ...dcl.ApplyOption) []NoteVulnerabilityDetails {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NoteVulnerabilityDetails, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNoteVulnerabilityDetails(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NoteVulnerabilityDetails, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNoteVulnerabilityDetails(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNoteVulnerabilityDetails(c *Client, des, nw *NoteVulnerabilityDetails) *NoteVulnerabilityDetails {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for NoteVulnerabilityDetails while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.SeverityName, nw.SeverityName) {
		nw.SeverityName = des.SeverityName
	}
	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}
	if dcl.StringCanonicalize(des.PackageType, nw.PackageType) {
		nw.PackageType = des.PackageType
	}
	if dcl.StringCanonicalize(des.AffectedCpeUri, nw.AffectedCpeUri) {
		nw.AffectedCpeUri = des.AffectedCpeUri
	}
	if dcl.StringCanonicalize(des.AffectedPackage, nw.AffectedPackage) {
		nw.AffectedPackage = des.AffectedPackage
	}
	nw.AffectedVersionStart = canonicalizeNewNoteVulnerabilityDetailsAffectedVersionStart(c, des.AffectedVersionStart, nw.AffectedVersionStart)
	nw.AffectedVersionEnd = canonicalizeNewNoteVulnerabilityDetailsAffectedVersionEnd(c, des.AffectedVersionEnd, nw.AffectedVersionEnd)
	if dcl.StringCanonicalize(des.FixedCpeUri, nw.FixedCpeUri) {
		nw.FixedCpeUri = des.FixedCpeUri
	}
	if dcl.StringCanonicalize(des.FixedPackage, nw.FixedPackage) {
		nw.FixedPackage = des.FixedPackage
	}
	nw.FixedVersion = canonicalizeNewNoteVulnerabilityDetailsFixedVersion(c, des.FixedVersion, nw.FixedVersion)
	if dcl.BoolCanonicalize(des.IsObsolete, nw.IsObsolete) {
		nw.IsObsolete = des.IsObsolete
	}

	return nw
}

func canonicalizeNewNoteVulnerabilityDetailsSet(c *Client, des, nw []NoteVulnerabilityDetails) []NoteVulnerabilityDetails {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []NoteVulnerabilityDetails
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareNoteVulnerabilityDetailsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewNoteVulnerabilityDetails(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewNoteVulnerabilityDetailsSlice(c *Client, des, nw []NoteVulnerabilityDetails) []NoteVulnerabilityDetails {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NoteVulnerabilityDetails
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteVulnerabilityDetails(c, &d, &n))
	}

	return items
}

func canonicalizeNoteVulnerabilityDetailsAffectedVersionStart(des, initial *NoteVulnerabilityDetailsAffectedVersionStart, opts ...dcl.ApplyOption) *NoteVulnerabilityDetailsAffectedVersionStart {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NoteVulnerabilityDetailsAffectedVersionStart{}

	if dcl.IsZeroValue(des.Epoch) || (dcl.IsEmptyValueIndirect(des.Epoch) && dcl.IsEmptyValueIndirect(initial.Epoch)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Epoch = initial.Epoch
	} else {
		cDes.Epoch = des.Epoch
	}
	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Revision, initial.Revision) || dcl.IsZeroValue(des.Revision) {
		cDes.Revision = initial.Revision
	} else {
		cDes.Revision = des.Revision
	}
	if dcl.IsZeroValue(des.Kind) || (dcl.IsEmptyValueIndirect(des.Kind) && dcl.IsEmptyValueIndirect(initial.Kind)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Kind = initial.Kind
	} else {
		cDes.Kind = des.Kind
	}
	if dcl.StringCanonicalize(des.FullName, initial.FullName) || dcl.IsZeroValue(des.FullName) {
		cDes.FullName = initial.FullName
	} else {
		cDes.FullName = des.FullName
	}

	return cDes
}

func canonicalizeNoteVulnerabilityDetailsAffectedVersionStartSlice(des, initial []NoteVulnerabilityDetailsAffectedVersionStart, opts ...dcl.ApplyOption) []NoteVulnerabilityDetailsAffectedVersionStart {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NoteVulnerabilityDetailsAffectedVersionStart, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNoteVulnerabilityDetailsAffectedVersionStart(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NoteVulnerabilityDetailsAffectedVersionStart, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNoteVulnerabilityDetailsAffectedVersionStart(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNoteVulnerabilityDetailsAffectedVersionStart(c *Client, des, nw *NoteVulnerabilityDetailsAffectedVersionStart) *NoteVulnerabilityDetailsAffectedVersionStart {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for NoteVulnerabilityDetailsAffectedVersionStart while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Revision, nw.Revision) {
		nw.Revision = des.Revision
	}
	if dcl.StringCanonicalize(des.FullName, nw.FullName) {
		nw.FullName = des.FullName
	}

	return nw
}

func canonicalizeNewNoteVulnerabilityDetailsAffectedVersionStartSet(c *Client, des, nw []NoteVulnerabilityDetailsAffectedVersionStart) []NoteVulnerabilityDetailsAffectedVersionStart {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []NoteVulnerabilityDetailsAffectedVersionStart
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareNoteVulnerabilityDetailsAffectedVersionStartNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewNoteVulnerabilityDetailsAffectedVersionStart(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewNoteVulnerabilityDetailsAffectedVersionStartSlice(c *Client, des, nw []NoteVulnerabilityDetailsAffectedVersionStart) []NoteVulnerabilityDetailsAffectedVersionStart {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NoteVulnerabilityDetailsAffectedVersionStart
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteVulnerabilityDetailsAffectedVersionStart(c, &d, &n))
	}

	return items
}

func canonicalizeNoteVulnerabilityDetailsAffectedVersionEnd(des, initial *NoteVulnerabilityDetailsAffectedVersionEnd, opts ...dcl.ApplyOption) *NoteVulnerabilityDetailsAffectedVersionEnd {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NoteVulnerabilityDetailsAffectedVersionEnd{}

	if dcl.IsZeroValue(des.Epoch) || (dcl.IsEmptyValueIndirect(des.Epoch) && dcl.IsEmptyValueIndirect(initial.Epoch)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Epoch = initial.Epoch
	} else {
		cDes.Epoch = des.Epoch
	}
	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Revision, initial.Revision) || dcl.IsZeroValue(des.Revision) {
		cDes.Revision = initial.Revision
	} else {
		cDes.Revision = des.Revision
	}
	if dcl.IsZeroValue(des.Kind) || (dcl.IsEmptyValueIndirect(des.Kind) && dcl.IsEmptyValueIndirect(initial.Kind)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Kind = initial.Kind
	} else {
		cDes.Kind = des.Kind
	}
	if dcl.StringCanonicalize(des.FullName, initial.FullName) || dcl.IsZeroValue(des.FullName) {
		cDes.FullName = initial.FullName
	} else {
		cDes.FullName = des.FullName
	}

	return cDes
}

func canonicalizeNoteVulnerabilityDetailsAffectedVersionEndSlice(des, initial []NoteVulnerabilityDetailsAffectedVersionEnd, opts ...dcl.ApplyOption) []NoteVulnerabilityDetailsAffectedVersionEnd {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NoteVulnerabilityDetailsAffectedVersionEnd, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNoteVulnerabilityDetailsAffectedVersionEnd(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NoteVulnerabilityDetailsAffectedVersionEnd, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNoteVulnerabilityDetailsAffectedVersionEnd(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNoteVulnerabilityDetailsAffectedVersionEnd(c *Client, des, nw *NoteVulnerabilityDetailsAffectedVersionEnd) *NoteVulnerabilityDetailsAffectedVersionEnd {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for NoteVulnerabilityDetailsAffectedVersionEnd while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Revision, nw.Revision) {
		nw.Revision = des.Revision
	}
	if dcl.StringCanonicalize(des.FullName, nw.FullName) {
		nw.FullName = des.FullName
	}

	return nw
}

func canonicalizeNewNoteVulnerabilityDetailsAffectedVersionEndSet(c *Client, des, nw []NoteVulnerabilityDetailsAffectedVersionEnd) []NoteVulnerabilityDetailsAffectedVersionEnd {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []NoteVulnerabilityDetailsAffectedVersionEnd
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareNoteVulnerabilityDetailsAffectedVersionEndNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewNoteVulnerabilityDetailsAffectedVersionEnd(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewNoteVulnerabilityDetailsAffectedVersionEndSlice(c *Client, des, nw []NoteVulnerabilityDetailsAffectedVersionEnd) []NoteVulnerabilityDetailsAffectedVersionEnd {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NoteVulnerabilityDetailsAffectedVersionEnd
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteVulnerabilityDetailsAffectedVersionEnd(c, &d, &n))
	}

	return items
}

func canonicalizeNoteVulnerabilityDetailsFixedVersion(des, initial *NoteVulnerabilityDetailsFixedVersion, opts ...dcl.ApplyOption) *NoteVulnerabilityDetailsFixedVersion {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NoteVulnerabilityDetailsFixedVersion{}

	if dcl.IsZeroValue(des.Epoch) || (dcl.IsEmptyValueIndirect(des.Epoch) && dcl.IsEmptyValueIndirect(initial.Epoch)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Epoch = initial.Epoch
	} else {
		cDes.Epoch = des.Epoch
	}
	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Revision, initial.Revision) || dcl.IsZeroValue(des.Revision) {
		cDes.Revision = initial.Revision
	} else {
		cDes.Revision = des.Revision
	}
	if dcl.IsZeroValue(des.Kind) || (dcl.IsEmptyValueIndirect(des.Kind) && dcl.IsEmptyValueIndirect(initial.Kind)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Kind = initial.Kind
	} else {
		cDes.Kind = des.Kind
	}
	if dcl.StringCanonicalize(des.FullName, initial.FullName) || dcl.IsZeroValue(des.FullName) {
		cDes.FullName = initial.FullName
	} else {
		cDes.FullName = des.FullName
	}

	return cDes
}

func canonicalizeNoteVulnerabilityDetailsFixedVersionSlice(des, initial []NoteVulnerabilityDetailsFixedVersion, opts ...dcl.ApplyOption) []NoteVulnerabilityDetailsFixedVersion {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NoteVulnerabilityDetailsFixedVersion, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNoteVulnerabilityDetailsFixedVersion(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NoteVulnerabilityDetailsFixedVersion, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNoteVulnerabilityDetailsFixedVersion(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNoteVulnerabilityDetailsFixedVersion(c *Client, des, nw *NoteVulnerabilityDetailsFixedVersion) *NoteVulnerabilityDetailsFixedVersion {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for NoteVulnerabilityDetailsFixedVersion while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Revision, nw.Revision) {
		nw.Revision = des.Revision
	}
	if dcl.StringCanonicalize(des.FullName, nw.FullName) {
		nw.FullName = des.FullName
	}

	return nw
}

func canonicalizeNewNoteVulnerabilityDetailsFixedVersionSet(c *Client, des, nw []NoteVulnerabilityDetailsFixedVersion) []NoteVulnerabilityDetailsFixedVersion {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []NoteVulnerabilityDetailsFixedVersion
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareNoteVulnerabilityDetailsFixedVersionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewNoteVulnerabilityDetailsFixedVersion(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewNoteVulnerabilityDetailsFixedVersionSlice(c *Client, des, nw []NoteVulnerabilityDetailsFixedVersion) []NoteVulnerabilityDetailsFixedVersion {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NoteVulnerabilityDetailsFixedVersion
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteVulnerabilityDetailsFixedVersion(c, &d, &n))
	}

	return items
}

func canonicalizeNoteVulnerabilityCvssV3(des, initial *NoteVulnerabilityCvssV3, opts ...dcl.ApplyOption) *NoteVulnerabilityCvssV3 {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NoteVulnerabilityCvssV3{}

	if dcl.IsZeroValue(des.BaseScore) || (dcl.IsEmptyValueIndirect(des.BaseScore) && dcl.IsEmptyValueIndirect(initial.BaseScore)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.BaseScore = initial.BaseScore
	} else {
		cDes.BaseScore = des.BaseScore
	}
	if dcl.IsZeroValue(des.ExploitabilityScore) || (dcl.IsEmptyValueIndirect(des.ExploitabilityScore) && dcl.IsEmptyValueIndirect(initial.ExploitabilityScore)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ExploitabilityScore = initial.ExploitabilityScore
	} else {
		cDes.ExploitabilityScore = des.ExploitabilityScore
	}
	if dcl.IsZeroValue(des.ImpactScore) || (dcl.IsEmptyValueIndirect(des.ImpactScore) && dcl.IsEmptyValueIndirect(initial.ImpactScore)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ImpactScore = initial.ImpactScore
	} else {
		cDes.ImpactScore = des.ImpactScore
	}
	if dcl.IsZeroValue(des.AttackVector) || (dcl.IsEmptyValueIndirect(des.AttackVector) && dcl.IsEmptyValueIndirect(initial.AttackVector)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.AttackVector = initial.AttackVector
	} else {
		cDes.AttackVector = des.AttackVector
	}
	if dcl.IsZeroValue(des.AttackComplexity) || (dcl.IsEmptyValueIndirect(des.AttackComplexity) && dcl.IsEmptyValueIndirect(initial.AttackComplexity)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.AttackComplexity = initial.AttackComplexity
	} else {
		cDes.AttackComplexity = des.AttackComplexity
	}
	if dcl.IsZeroValue(des.PrivilegesRequired) || (dcl.IsEmptyValueIndirect(des.PrivilegesRequired) && dcl.IsEmptyValueIndirect(initial.PrivilegesRequired)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.PrivilegesRequired = initial.PrivilegesRequired
	} else {
		cDes.PrivilegesRequired = des.PrivilegesRequired
	}
	if dcl.IsZeroValue(des.UserInteraction) || (dcl.IsEmptyValueIndirect(des.UserInteraction) && dcl.IsEmptyValueIndirect(initial.UserInteraction)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.UserInteraction = initial.UserInteraction
	} else {
		cDes.UserInteraction = des.UserInteraction
	}
	if dcl.IsZeroValue(des.Scope) || (dcl.IsEmptyValueIndirect(des.Scope) && dcl.IsEmptyValueIndirect(initial.Scope)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Scope = initial.Scope
	} else {
		cDes.Scope = des.Scope
	}
	if dcl.IsZeroValue(des.ConfidentialityImpact) || (dcl.IsEmptyValueIndirect(des.ConfidentialityImpact) && dcl.IsEmptyValueIndirect(initial.ConfidentialityImpact)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ConfidentialityImpact = initial.ConfidentialityImpact
	} else {
		cDes.ConfidentialityImpact = des.ConfidentialityImpact
	}
	if dcl.IsZeroValue(des.IntegrityImpact) || (dcl.IsEmptyValueIndirect(des.IntegrityImpact) && dcl.IsEmptyValueIndirect(initial.IntegrityImpact)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.IntegrityImpact = initial.IntegrityImpact
	} else {
		cDes.IntegrityImpact = des.IntegrityImpact
	}
	if dcl.IsZeroValue(des.AvailabilityImpact) || (dcl.IsEmptyValueIndirect(des.AvailabilityImpact) && dcl.IsEmptyValueIndirect(initial.AvailabilityImpact)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.AvailabilityImpact = initial.AvailabilityImpact
	} else {
		cDes.AvailabilityImpact = des.AvailabilityImpact
	}

	return cDes
}

func canonicalizeNoteVulnerabilityCvssV3Slice(des, initial []NoteVulnerabilityCvssV3, opts ...dcl.ApplyOption) []NoteVulnerabilityCvssV3 {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NoteVulnerabilityCvssV3, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNoteVulnerabilityCvssV3(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NoteVulnerabilityCvssV3, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNoteVulnerabilityCvssV3(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNoteVulnerabilityCvssV3(c *Client, des, nw *NoteVulnerabilityCvssV3) *NoteVulnerabilityCvssV3 {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for NoteVulnerabilityCvssV3 while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewNoteVulnerabilityCvssV3Set(c *Client, des, nw []NoteVulnerabilityCvssV3) []NoteVulnerabilityCvssV3 {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []NoteVulnerabilityCvssV3
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareNoteVulnerabilityCvssV3NewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewNoteVulnerabilityCvssV3(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewNoteVulnerabilityCvssV3Slice(c *Client, des, nw []NoteVulnerabilityCvssV3) []NoteVulnerabilityCvssV3 {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NoteVulnerabilityCvssV3
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteVulnerabilityCvssV3(c, &d, &n))
	}

	return items
}

func canonicalizeNoteVulnerabilityWindowsDetails(des, initial *NoteVulnerabilityWindowsDetails, opts ...dcl.ApplyOption) *NoteVulnerabilityWindowsDetails {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NoteVulnerabilityWindowsDetails{}

	if dcl.StringCanonicalize(des.CpeUri, initial.CpeUri) || dcl.IsZeroValue(des.CpeUri) {
		cDes.CpeUri = initial.CpeUri
	} else {
		cDes.CpeUri = des.CpeUri
	}
	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		cDes.Description = initial.Description
	} else {
		cDes.Description = des.Description
	}
	cDes.FixingKbs = canonicalizeNoteVulnerabilityWindowsDetailsFixingKbsSlice(des.FixingKbs, initial.FixingKbs, opts...)

	return cDes
}

func canonicalizeNoteVulnerabilityWindowsDetailsSlice(des, initial []NoteVulnerabilityWindowsDetails, opts ...dcl.ApplyOption) []NoteVulnerabilityWindowsDetails {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NoteVulnerabilityWindowsDetails, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNoteVulnerabilityWindowsDetails(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NoteVulnerabilityWindowsDetails, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNoteVulnerabilityWindowsDetails(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNoteVulnerabilityWindowsDetails(c *Client, des, nw *NoteVulnerabilityWindowsDetails) *NoteVulnerabilityWindowsDetails {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for NoteVulnerabilityWindowsDetails while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.CpeUri, nw.CpeUri) {
		nw.CpeUri = des.CpeUri
	}
	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}
	nw.FixingKbs = canonicalizeNewNoteVulnerabilityWindowsDetailsFixingKbsSlice(c, des.FixingKbs, nw.FixingKbs)

	return nw
}

func canonicalizeNewNoteVulnerabilityWindowsDetailsSet(c *Client, des, nw []NoteVulnerabilityWindowsDetails) []NoteVulnerabilityWindowsDetails {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []NoteVulnerabilityWindowsDetails
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareNoteVulnerabilityWindowsDetailsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewNoteVulnerabilityWindowsDetails(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewNoteVulnerabilityWindowsDetailsSlice(c *Client, des, nw []NoteVulnerabilityWindowsDetails) []NoteVulnerabilityWindowsDetails {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NoteVulnerabilityWindowsDetails
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteVulnerabilityWindowsDetails(c, &d, &n))
	}

	return items
}

func canonicalizeNoteVulnerabilityWindowsDetailsFixingKbs(des, initial *NoteVulnerabilityWindowsDetailsFixingKbs, opts ...dcl.ApplyOption) *NoteVulnerabilityWindowsDetailsFixingKbs {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NoteVulnerabilityWindowsDetailsFixingKbs{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Url, initial.Url) || dcl.IsZeroValue(des.Url) {
		cDes.Url = initial.Url
	} else {
		cDes.Url = des.Url
	}

	return cDes
}

func canonicalizeNoteVulnerabilityWindowsDetailsFixingKbsSlice(des, initial []NoteVulnerabilityWindowsDetailsFixingKbs, opts ...dcl.ApplyOption) []NoteVulnerabilityWindowsDetailsFixingKbs {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NoteVulnerabilityWindowsDetailsFixingKbs, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNoteVulnerabilityWindowsDetailsFixingKbs(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NoteVulnerabilityWindowsDetailsFixingKbs, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNoteVulnerabilityWindowsDetailsFixingKbs(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNoteVulnerabilityWindowsDetailsFixingKbs(c *Client, des, nw *NoteVulnerabilityWindowsDetailsFixingKbs) *NoteVulnerabilityWindowsDetailsFixingKbs {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for NoteVulnerabilityWindowsDetailsFixingKbs while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Url, nw.Url) {
		nw.Url = des.Url
	}

	return nw
}

func canonicalizeNewNoteVulnerabilityWindowsDetailsFixingKbsSet(c *Client, des, nw []NoteVulnerabilityWindowsDetailsFixingKbs) []NoteVulnerabilityWindowsDetailsFixingKbs {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []NoteVulnerabilityWindowsDetailsFixingKbs
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareNoteVulnerabilityWindowsDetailsFixingKbsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewNoteVulnerabilityWindowsDetailsFixingKbs(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewNoteVulnerabilityWindowsDetailsFixingKbsSlice(c *Client, des, nw []NoteVulnerabilityWindowsDetailsFixingKbs) []NoteVulnerabilityWindowsDetailsFixingKbs {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NoteVulnerabilityWindowsDetailsFixingKbs
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteVulnerabilityWindowsDetailsFixingKbs(c, &d, &n))
	}

	return items
}

func canonicalizeNoteBuild(des, initial *NoteBuild, opts ...dcl.ApplyOption) *NoteBuild {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NoteBuild{}

	if dcl.StringCanonicalize(des.BuilderVersion, initial.BuilderVersion) || dcl.IsZeroValue(des.BuilderVersion) {
		cDes.BuilderVersion = initial.BuilderVersion
	} else {
		cDes.BuilderVersion = des.BuilderVersion
	}
	cDes.Signature = canonicalizeNoteBuildSignature(des.Signature, initial.Signature, opts...)

	return cDes
}

func canonicalizeNoteBuildSlice(des, initial []NoteBuild, opts ...dcl.ApplyOption) []NoteBuild {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NoteBuild, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNoteBuild(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NoteBuild, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNoteBuild(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNoteBuild(c *Client, des, nw *NoteBuild) *NoteBuild {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for NoteBuild while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.BuilderVersion, nw.BuilderVersion) {
		nw.BuilderVersion = des.BuilderVersion
	}
	nw.Signature = canonicalizeNewNoteBuildSignature(c, des.Signature, nw.Signature)

	return nw
}

func canonicalizeNewNoteBuildSet(c *Client, des, nw []NoteBuild) []NoteBuild {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []NoteBuild
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareNoteBuildNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewNoteBuild(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewNoteBuildSlice(c *Client, des, nw []NoteBuild) []NoteBuild {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NoteBuild
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteBuild(c, &d, &n))
	}

	return items
}

func canonicalizeNoteBuildSignature(des, initial *NoteBuildSignature, opts ...dcl.ApplyOption) *NoteBuildSignature {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NoteBuildSignature{}

	if dcl.StringCanonicalize(des.PublicKey, initial.PublicKey) || dcl.IsZeroValue(des.PublicKey) {
		cDes.PublicKey = initial.PublicKey
	} else {
		cDes.PublicKey = des.PublicKey
	}
	if dcl.StringCanonicalize(des.Signature, initial.Signature) || dcl.IsZeroValue(des.Signature) {
		cDes.Signature = initial.Signature
	} else {
		cDes.Signature = des.Signature
	}
	if dcl.StringCanonicalize(des.KeyId, initial.KeyId) || dcl.IsZeroValue(des.KeyId) {
		cDes.KeyId = initial.KeyId
	} else {
		cDes.KeyId = des.KeyId
	}
	if dcl.IsZeroValue(des.KeyType) || (dcl.IsEmptyValueIndirect(des.KeyType) && dcl.IsEmptyValueIndirect(initial.KeyType)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.KeyType = initial.KeyType
	} else {
		cDes.KeyType = des.KeyType
	}

	return cDes
}

func canonicalizeNoteBuildSignatureSlice(des, initial []NoteBuildSignature, opts ...dcl.ApplyOption) []NoteBuildSignature {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NoteBuildSignature, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNoteBuildSignature(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NoteBuildSignature, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNoteBuildSignature(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNoteBuildSignature(c *Client, des, nw *NoteBuildSignature) *NoteBuildSignature {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for NoteBuildSignature while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.PublicKey, nw.PublicKey) {
		nw.PublicKey = des.PublicKey
	}
	if dcl.StringCanonicalize(des.Signature, nw.Signature) {
		nw.Signature = des.Signature
	}
	if dcl.StringCanonicalize(des.KeyId, nw.KeyId) {
		nw.KeyId = des.KeyId
	}

	return nw
}

func canonicalizeNewNoteBuildSignatureSet(c *Client, des, nw []NoteBuildSignature) []NoteBuildSignature {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []NoteBuildSignature
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareNoteBuildSignatureNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewNoteBuildSignature(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewNoteBuildSignatureSlice(c *Client, des, nw []NoteBuildSignature) []NoteBuildSignature {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NoteBuildSignature
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteBuildSignature(c, &d, &n))
	}

	return items
}

func canonicalizeNoteImage(des, initial *NoteImage, opts ...dcl.ApplyOption) *NoteImage {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NoteImage{}

	if dcl.StringCanonicalize(des.ResourceUrl, initial.ResourceUrl) || dcl.IsZeroValue(des.ResourceUrl) {
		cDes.ResourceUrl = initial.ResourceUrl
	} else {
		cDes.ResourceUrl = des.ResourceUrl
	}
	cDes.Fingerprint = canonicalizeNoteImageFingerprint(des.Fingerprint, initial.Fingerprint, opts...)

	return cDes
}

func canonicalizeNoteImageSlice(des, initial []NoteImage, opts ...dcl.ApplyOption) []NoteImage {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NoteImage, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNoteImage(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NoteImage, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNoteImage(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNoteImage(c *Client, des, nw *NoteImage) *NoteImage {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for NoteImage while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.ResourceUrl, nw.ResourceUrl) {
		nw.ResourceUrl = des.ResourceUrl
	}
	nw.Fingerprint = canonicalizeNewNoteImageFingerprint(c, des.Fingerprint, nw.Fingerprint)

	return nw
}

func canonicalizeNewNoteImageSet(c *Client, des, nw []NoteImage) []NoteImage {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []NoteImage
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareNoteImageNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewNoteImage(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewNoteImageSlice(c *Client, des, nw []NoteImage) []NoteImage {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NoteImage
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteImage(c, &d, &n))
	}

	return items
}

func canonicalizeNoteImageFingerprint(des, initial *NoteImageFingerprint, opts ...dcl.ApplyOption) *NoteImageFingerprint {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NoteImageFingerprint{}

	if dcl.StringCanonicalize(des.V1Name, initial.V1Name) || dcl.IsZeroValue(des.V1Name) {
		cDes.V1Name = initial.V1Name
	} else {
		cDes.V1Name = des.V1Name
	}
	if dcl.StringArrayCanonicalize(des.V2Blob, initial.V2Blob) {
		cDes.V2Blob = initial.V2Blob
	} else {
		cDes.V2Blob = des.V2Blob
	}

	return cDes
}

func canonicalizeNoteImageFingerprintSlice(des, initial []NoteImageFingerprint, opts ...dcl.ApplyOption) []NoteImageFingerprint {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NoteImageFingerprint, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNoteImageFingerprint(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NoteImageFingerprint, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNoteImageFingerprint(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNoteImageFingerprint(c *Client, des, nw *NoteImageFingerprint) *NoteImageFingerprint {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for NoteImageFingerprint while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.V1Name, nw.V1Name) {
		nw.V1Name = des.V1Name
	}
	if dcl.StringArrayCanonicalize(des.V2Blob, nw.V2Blob) {
		nw.V2Blob = des.V2Blob
	}
	if dcl.StringCanonicalize(des.V2Name, nw.V2Name) {
		nw.V2Name = des.V2Name
	}

	return nw
}

func canonicalizeNewNoteImageFingerprintSet(c *Client, des, nw []NoteImageFingerprint) []NoteImageFingerprint {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []NoteImageFingerprint
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareNoteImageFingerprintNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewNoteImageFingerprint(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewNoteImageFingerprintSlice(c *Client, des, nw []NoteImageFingerprint) []NoteImageFingerprint {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NoteImageFingerprint
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteImageFingerprint(c, &d, &n))
	}

	return items
}

func canonicalizeNotePackage(des, initial *NotePackage, opts ...dcl.ApplyOption) *NotePackage {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NotePackage{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	cDes.Distribution = canonicalizeNotePackageDistributionSlice(des.Distribution, initial.Distribution, opts...)

	return cDes
}

func canonicalizeNotePackageSlice(des, initial []NotePackage, opts ...dcl.ApplyOption) []NotePackage {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NotePackage, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNotePackage(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NotePackage, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNotePackage(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNotePackage(c *Client, des, nw *NotePackage) *NotePackage {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for NotePackage while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	nw.Distribution = canonicalizeNewNotePackageDistributionSlice(c, des.Distribution, nw.Distribution)

	return nw
}

func canonicalizeNewNotePackageSet(c *Client, des, nw []NotePackage) []NotePackage {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []NotePackage
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareNotePackageNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewNotePackage(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewNotePackageSlice(c *Client, des, nw []NotePackage) []NotePackage {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NotePackage
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNotePackage(c, &d, &n))
	}

	return items
}

func canonicalizeNotePackageDistribution(des, initial *NotePackageDistribution, opts ...dcl.ApplyOption) *NotePackageDistribution {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NotePackageDistribution{}

	if dcl.StringCanonicalize(des.CpeUri, initial.CpeUri) || dcl.IsZeroValue(des.CpeUri) {
		cDes.CpeUri = initial.CpeUri
	} else {
		cDes.CpeUri = des.CpeUri
	}
	if dcl.IsZeroValue(des.Architecture) || (dcl.IsEmptyValueIndirect(des.Architecture) && dcl.IsEmptyValueIndirect(initial.Architecture)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Architecture = initial.Architecture
	} else {
		cDes.Architecture = des.Architecture
	}
	cDes.LatestVersion = canonicalizeNotePackageDistributionLatestVersion(des.LatestVersion, initial.LatestVersion, opts...)
	if dcl.StringCanonicalize(des.Maintainer, initial.Maintainer) || dcl.IsZeroValue(des.Maintainer) {
		cDes.Maintainer = initial.Maintainer
	} else {
		cDes.Maintainer = des.Maintainer
	}
	if dcl.StringCanonicalize(des.Url, initial.Url) || dcl.IsZeroValue(des.Url) {
		cDes.Url = initial.Url
	} else {
		cDes.Url = des.Url
	}
	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		cDes.Description = initial.Description
	} else {
		cDes.Description = des.Description
	}

	return cDes
}

func canonicalizeNotePackageDistributionSlice(des, initial []NotePackageDistribution, opts ...dcl.ApplyOption) []NotePackageDistribution {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NotePackageDistribution, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNotePackageDistribution(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NotePackageDistribution, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNotePackageDistribution(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNotePackageDistribution(c *Client, des, nw *NotePackageDistribution) *NotePackageDistribution {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for NotePackageDistribution while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.CpeUri, nw.CpeUri) {
		nw.CpeUri = des.CpeUri
	}
	nw.LatestVersion = canonicalizeNewNotePackageDistributionLatestVersion(c, des.LatestVersion, nw.LatestVersion)
	if dcl.StringCanonicalize(des.Maintainer, nw.Maintainer) {
		nw.Maintainer = des.Maintainer
	}
	if dcl.StringCanonicalize(des.Url, nw.Url) {
		nw.Url = des.Url
	}
	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}

	return nw
}

func canonicalizeNewNotePackageDistributionSet(c *Client, des, nw []NotePackageDistribution) []NotePackageDistribution {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []NotePackageDistribution
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareNotePackageDistributionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewNotePackageDistribution(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewNotePackageDistributionSlice(c *Client, des, nw []NotePackageDistribution) []NotePackageDistribution {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NotePackageDistribution
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNotePackageDistribution(c, &d, &n))
	}

	return items
}

func canonicalizeNotePackageDistributionLatestVersion(des, initial *NotePackageDistributionLatestVersion, opts ...dcl.ApplyOption) *NotePackageDistributionLatestVersion {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NotePackageDistributionLatestVersion{}

	if dcl.IsZeroValue(des.Epoch) || (dcl.IsEmptyValueIndirect(des.Epoch) && dcl.IsEmptyValueIndirect(initial.Epoch)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Epoch = initial.Epoch
	} else {
		cDes.Epoch = des.Epoch
	}
	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Revision, initial.Revision) || dcl.IsZeroValue(des.Revision) {
		cDes.Revision = initial.Revision
	} else {
		cDes.Revision = des.Revision
	}
	if dcl.IsZeroValue(des.Kind) || (dcl.IsEmptyValueIndirect(des.Kind) && dcl.IsEmptyValueIndirect(initial.Kind)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Kind = initial.Kind
	} else {
		cDes.Kind = des.Kind
	}
	if dcl.StringCanonicalize(des.FullName, initial.FullName) || dcl.IsZeroValue(des.FullName) {
		cDes.FullName = initial.FullName
	} else {
		cDes.FullName = des.FullName
	}

	return cDes
}

func canonicalizeNotePackageDistributionLatestVersionSlice(des, initial []NotePackageDistributionLatestVersion, opts ...dcl.ApplyOption) []NotePackageDistributionLatestVersion {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NotePackageDistributionLatestVersion, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNotePackageDistributionLatestVersion(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NotePackageDistributionLatestVersion, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNotePackageDistributionLatestVersion(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNotePackageDistributionLatestVersion(c *Client, des, nw *NotePackageDistributionLatestVersion) *NotePackageDistributionLatestVersion {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for NotePackageDistributionLatestVersion while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Revision, nw.Revision) {
		nw.Revision = des.Revision
	}
	if dcl.StringCanonicalize(des.FullName, nw.FullName) {
		nw.FullName = des.FullName
	}

	return nw
}

func canonicalizeNewNotePackageDistributionLatestVersionSet(c *Client, des, nw []NotePackageDistributionLatestVersion) []NotePackageDistributionLatestVersion {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []NotePackageDistributionLatestVersion
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareNotePackageDistributionLatestVersionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewNotePackageDistributionLatestVersion(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewNotePackageDistributionLatestVersionSlice(c *Client, des, nw []NotePackageDistributionLatestVersion) []NotePackageDistributionLatestVersion {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NotePackageDistributionLatestVersion
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNotePackageDistributionLatestVersion(c, &d, &n))
	}

	return items
}

func canonicalizeNoteDiscovery(des, initial *NoteDiscovery, opts ...dcl.ApplyOption) *NoteDiscovery {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NoteDiscovery{}

	if dcl.IsZeroValue(des.AnalysisKind) || (dcl.IsEmptyValueIndirect(des.AnalysisKind) && dcl.IsEmptyValueIndirect(initial.AnalysisKind)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.AnalysisKind = initial.AnalysisKind
	} else {
		cDes.AnalysisKind = des.AnalysisKind
	}

	return cDes
}

func canonicalizeNoteDiscoverySlice(des, initial []NoteDiscovery, opts ...dcl.ApplyOption) []NoteDiscovery {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NoteDiscovery, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNoteDiscovery(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NoteDiscovery, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNoteDiscovery(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNoteDiscovery(c *Client, des, nw *NoteDiscovery) *NoteDiscovery {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for NoteDiscovery while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewNoteDiscoverySet(c *Client, des, nw []NoteDiscovery) []NoteDiscovery {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []NoteDiscovery
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareNoteDiscoveryNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewNoteDiscovery(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewNoteDiscoverySlice(c *Client, des, nw []NoteDiscovery) []NoteDiscovery {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NoteDiscovery
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteDiscovery(c, &d, &n))
	}

	return items
}

func canonicalizeNoteDeployment(des, initial *NoteDeployment, opts ...dcl.ApplyOption) *NoteDeployment {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NoteDeployment{}

	if dcl.StringArrayCanonicalize(des.ResourceUri, initial.ResourceUri) {
		cDes.ResourceUri = initial.ResourceUri
	} else {
		cDes.ResourceUri = des.ResourceUri
	}

	return cDes
}

func canonicalizeNoteDeploymentSlice(des, initial []NoteDeployment, opts ...dcl.ApplyOption) []NoteDeployment {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NoteDeployment, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNoteDeployment(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NoteDeployment, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNoteDeployment(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNoteDeployment(c *Client, des, nw *NoteDeployment) *NoteDeployment {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for NoteDeployment while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.ResourceUri, nw.ResourceUri) {
		nw.ResourceUri = des.ResourceUri
	}

	return nw
}

func canonicalizeNewNoteDeploymentSet(c *Client, des, nw []NoteDeployment) []NoteDeployment {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []NoteDeployment
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareNoteDeploymentNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewNoteDeployment(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewNoteDeploymentSlice(c *Client, des, nw []NoteDeployment) []NoteDeployment {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NoteDeployment
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteDeployment(c, &d, &n))
	}

	return items
}

func canonicalizeNoteAttestation(des, initial *NoteAttestation, opts ...dcl.ApplyOption) *NoteAttestation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NoteAttestation{}

	cDes.Hint = canonicalizeNoteAttestationHint(des.Hint, initial.Hint, opts...)

	return cDes
}

func canonicalizeNoteAttestationSlice(des, initial []NoteAttestation, opts ...dcl.ApplyOption) []NoteAttestation {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NoteAttestation, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNoteAttestation(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NoteAttestation, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNoteAttestation(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNoteAttestation(c *Client, des, nw *NoteAttestation) *NoteAttestation {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for NoteAttestation while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Hint = canonicalizeNewNoteAttestationHint(c, des.Hint, nw.Hint)

	return nw
}

func canonicalizeNewNoteAttestationSet(c *Client, des, nw []NoteAttestation) []NoteAttestation {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []NoteAttestation
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareNoteAttestationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewNoteAttestation(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewNoteAttestationSlice(c *Client, des, nw []NoteAttestation) []NoteAttestation {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NoteAttestation
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteAttestation(c, &d, &n))
	}

	return items
}

func canonicalizeNoteAttestationHint(des, initial *NoteAttestationHint, opts ...dcl.ApplyOption) *NoteAttestationHint {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &NoteAttestationHint{}

	if dcl.StringCanonicalize(des.HumanReadableName, initial.HumanReadableName) || dcl.IsZeroValue(des.HumanReadableName) {
		cDes.HumanReadableName = initial.HumanReadableName
	} else {
		cDes.HumanReadableName = des.HumanReadableName
	}

	return cDes
}

func canonicalizeNoteAttestationHintSlice(des, initial []NoteAttestationHint, opts ...dcl.ApplyOption) []NoteAttestationHint {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]NoteAttestationHint, 0, len(des))
		for _, d := range des {
			cd := canonicalizeNoteAttestationHint(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]NoteAttestationHint, 0, len(des))
	for i, d := range des {
		cd := canonicalizeNoteAttestationHint(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewNoteAttestationHint(c *Client, des, nw *NoteAttestationHint) *NoteAttestationHint {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for NoteAttestationHint while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.HumanReadableName, nw.HumanReadableName) {
		nw.HumanReadableName = des.HumanReadableName
	}

	return nw
}

func canonicalizeNewNoteAttestationHintSet(c *Client, des, nw []NoteAttestationHint) []NoteAttestationHint {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []NoteAttestationHint
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareNoteAttestationHintNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewNoteAttestationHint(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewNoteAttestationHintSlice(c *Client, des, nw []NoteAttestationHint) []NoteAttestationHint {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []NoteAttestationHint
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewNoteAttestationHint(c, &d, &n))
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
func diffNote(c *Client, desired, actual *Note, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.ShortDescription, actual.ShortDescription, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("ShortDescription")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LongDescription, actual.LongDescription, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("LongDescription")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RelatedUrl, actual.RelatedUrl, dcl.DiffInfo{ObjectFunction: compareNoteRelatedUrlNewStyle, EmptyObject: EmptyNoteRelatedUrl, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("RelatedUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExpirationTime, actual.ExpirationTime, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("ExpirationTime")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.RelatedNoteNames, actual.RelatedNoteNames, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("RelatedNoteNames")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Vulnerability, actual.Vulnerability, dcl.DiffInfo{ObjectFunction: compareNoteVulnerabilityNewStyle, EmptyObject: EmptyNoteVulnerability, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Vulnerability")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Build, actual.Build, dcl.DiffInfo{ObjectFunction: compareNoteBuildNewStyle, EmptyObject: EmptyNoteBuild, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Build")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Image, actual.Image, dcl.DiffInfo{ObjectFunction: compareNoteImageNewStyle, EmptyObject: EmptyNoteImage, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("BaseImage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Package, actual.Package, dcl.DiffInfo{ObjectFunction: compareNotePackageNewStyle, EmptyObject: EmptyNotePackage, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Package")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Discovery, actual.Discovery, dcl.DiffInfo{ObjectFunction: compareNoteDiscoveryNewStyle, EmptyObject: EmptyNoteDiscovery, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Discovery")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Deployment, actual.Deployment, dcl.DiffInfo{ObjectFunction: compareNoteDeploymentNewStyle, EmptyObject: EmptyNoteDeployment, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Deployable")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Attestation, actual.Attestation, dcl.DiffInfo{ObjectFunction: compareNoteAttestationNewStyle, EmptyObject: EmptyNoteAttestation, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("AttestationAuthority")); len(ds) != 0 || err != nil {
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

	newDiffs = analyzeNoteDiff(desired, actual, newDiffs)
	if len(newDiffs) > 0 {
		c.Config.Logger.Infof("Diff function found diffs: %v", newDiffs)
	}
	return newDiffs, nil
}
func compareNoteRelatedUrlNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NoteRelatedUrl)
	if !ok {
		desiredNotPointer, ok := d.(NoteRelatedUrl)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteRelatedUrl or *NoteRelatedUrl", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NoteRelatedUrl)
	if !ok {
		actualNotPointer, ok := a.(NoteRelatedUrl)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteRelatedUrl", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Url, actual.Url, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Url")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Label, actual.Label, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Label")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNoteVulnerabilityNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NoteVulnerability)
	if !ok {
		desiredNotPointer, ok := d.(NoteVulnerability)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteVulnerability or *NoteVulnerability", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NoteVulnerability)
	if !ok {
		actualNotPointer, ok := a.(NoteVulnerability)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteVulnerability", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.CvssScore, actual.CvssScore, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("CvssScore")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Severity, actual.Severity, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Severity")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Details, actual.Details, dcl.DiffInfo{ObjectFunction: compareNoteVulnerabilityDetailsNewStyle, EmptyObject: EmptyNoteVulnerabilityDetails, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Details")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CvssV3, actual.CvssV3, dcl.DiffInfo{ObjectFunction: compareNoteVulnerabilityCvssV3NewStyle, EmptyObject: EmptyNoteVulnerabilityCvssV3, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("CvssV3")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WindowsDetails, actual.WindowsDetails, dcl.DiffInfo{ObjectFunction: compareNoteVulnerabilityWindowsDetailsNewStyle, EmptyObject: EmptyNoteVulnerabilityWindowsDetails, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("WindowsDetails")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceUpdateTime, actual.SourceUpdateTime, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("SourceUpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNoteVulnerabilityDetailsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NoteVulnerabilityDetails)
	if !ok {
		desiredNotPointer, ok := d.(NoteVulnerabilityDetails)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteVulnerabilityDetails or *NoteVulnerabilityDetails", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NoteVulnerabilityDetails)
	if !ok {
		actualNotPointer, ok := a.(NoteVulnerabilityDetails)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteVulnerabilityDetails", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SeverityName, actual.SeverityName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("SeverityName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PackageType, actual.PackageType, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("PackageType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AffectedCpeUri, actual.AffectedCpeUri, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("AffectedCpeUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AffectedPackage, actual.AffectedPackage, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("AffectedPackage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AffectedVersionStart, actual.AffectedVersionStart, dcl.DiffInfo{ObjectFunction: compareNoteVulnerabilityDetailsAffectedVersionStartNewStyle, EmptyObject: EmptyNoteVulnerabilityDetailsAffectedVersionStart, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("AffectedVersionStart")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AffectedVersionEnd, actual.AffectedVersionEnd, dcl.DiffInfo{ObjectFunction: compareNoteVulnerabilityDetailsAffectedVersionEndNewStyle, EmptyObject: EmptyNoteVulnerabilityDetailsAffectedVersionEnd, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("AffectedVersionEnd")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FixedCpeUri, actual.FixedCpeUri, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("FixedCpeUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FixedPackage, actual.FixedPackage, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("FixedPackage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FixedVersion, actual.FixedVersion, dcl.DiffInfo{ObjectFunction: compareNoteVulnerabilityDetailsFixedVersionNewStyle, EmptyObject: EmptyNoteVulnerabilityDetailsFixedVersion, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("FixedVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IsObsolete, actual.IsObsolete, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("IsObsolete")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceUpdateTime, actual.SourceUpdateTime, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("SourceUpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNoteVulnerabilityDetailsAffectedVersionStartNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NoteVulnerabilityDetailsAffectedVersionStart)
	if !ok {
		desiredNotPointer, ok := d.(NoteVulnerabilityDetailsAffectedVersionStart)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteVulnerabilityDetailsAffectedVersionStart or *NoteVulnerabilityDetailsAffectedVersionStart", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NoteVulnerabilityDetailsAffectedVersionStart)
	if !ok {
		actualNotPointer, ok := a.(NoteVulnerabilityDetailsAffectedVersionStart)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteVulnerabilityDetailsAffectedVersionStart", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Epoch, actual.Epoch, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Epoch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Revision, actual.Revision, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Revision")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Kind, actual.Kind, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Kind")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FullName, actual.FullName, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("FullName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNoteVulnerabilityDetailsAffectedVersionEndNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NoteVulnerabilityDetailsAffectedVersionEnd)
	if !ok {
		desiredNotPointer, ok := d.(NoteVulnerabilityDetailsAffectedVersionEnd)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteVulnerabilityDetailsAffectedVersionEnd or *NoteVulnerabilityDetailsAffectedVersionEnd", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NoteVulnerabilityDetailsAffectedVersionEnd)
	if !ok {
		actualNotPointer, ok := a.(NoteVulnerabilityDetailsAffectedVersionEnd)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteVulnerabilityDetailsAffectedVersionEnd", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Epoch, actual.Epoch, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Epoch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Revision, actual.Revision, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Revision")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Kind, actual.Kind, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Kind")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FullName, actual.FullName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("FullName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNoteVulnerabilityDetailsFixedVersionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NoteVulnerabilityDetailsFixedVersion)
	if !ok {
		desiredNotPointer, ok := d.(NoteVulnerabilityDetailsFixedVersion)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteVulnerabilityDetailsFixedVersion or *NoteVulnerabilityDetailsFixedVersion", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NoteVulnerabilityDetailsFixedVersion)
	if !ok {
		actualNotPointer, ok := a.(NoteVulnerabilityDetailsFixedVersion)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteVulnerabilityDetailsFixedVersion", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Epoch, actual.Epoch, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Epoch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Revision, actual.Revision, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Revision")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Kind, actual.Kind, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Kind")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FullName, actual.FullName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("FullName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNoteVulnerabilityCvssV3NewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NoteVulnerabilityCvssV3)
	if !ok {
		desiredNotPointer, ok := d.(NoteVulnerabilityCvssV3)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteVulnerabilityCvssV3 or *NoteVulnerabilityCvssV3", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NoteVulnerabilityCvssV3)
	if !ok {
		actualNotPointer, ok := a.(NoteVulnerabilityCvssV3)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteVulnerabilityCvssV3", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.BaseScore, actual.BaseScore, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("BaseScore")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExploitabilityScore, actual.ExploitabilityScore, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("ExploitabilityScore")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ImpactScore, actual.ImpactScore, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("ImpactScore")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AttackVector, actual.AttackVector, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("AttackVector")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AttackComplexity, actual.AttackComplexity, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("AttackComplexity")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PrivilegesRequired, actual.PrivilegesRequired, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("PrivilegesRequired")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UserInteraction, actual.UserInteraction, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("UserInteraction")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Scope, actual.Scope, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Scope")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ConfidentialityImpact, actual.ConfidentialityImpact, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("ConfidentialityImpact")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IntegrityImpact, actual.IntegrityImpact, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("IntegrityImpact")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AvailabilityImpact, actual.AvailabilityImpact, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("AvailabilityImpact")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNoteVulnerabilityWindowsDetailsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NoteVulnerabilityWindowsDetails)
	if !ok {
		desiredNotPointer, ok := d.(NoteVulnerabilityWindowsDetails)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteVulnerabilityWindowsDetails or *NoteVulnerabilityWindowsDetails", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NoteVulnerabilityWindowsDetails)
	if !ok {
		actualNotPointer, ok := a.(NoteVulnerabilityWindowsDetails)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteVulnerabilityWindowsDetails", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.CpeUri, actual.CpeUri, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("CpeUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FixingKbs, actual.FixingKbs, dcl.DiffInfo{ObjectFunction: compareNoteVulnerabilityWindowsDetailsFixingKbsNewStyle, EmptyObject: EmptyNoteVulnerabilityWindowsDetailsFixingKbs, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("FixingKbs")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNoteVulnerabilityWindowsDetailsFixingKbsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NoteVulnerabilityWindowsDetailsFixingKbs)
	if !ok {
		desiredNotPointer, ok := d.(NoteVulnerabilityWindowsDetailsFixingKbs)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteVulnerabilityWindowsDetailsFixingKbs or *NoteVulnerabilityWindowsDetailsFixingKbs", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NoteVulnerabilityWindowsDetailsFixingKbs)
	if !ok {
		actualNotPointer, ok := a.(NoteVulnerabilityWindowsDetailsFixingKbs)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteVulnerabilityWindowsDetailsFixingKbs", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Url, actual.Url, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Url")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNoteBuildNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NoteBuild)
	if !ok {
		desiredNotPointer, ok := d.(NoteBuild)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteBuild or *NoteBuild", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NoteBuild)
	if !ok {
		actualNotPointer, ok := a.(NoteBuild)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteBuild", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.BuilderVersion, actual.BuilderVersion, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("BuilderVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Signature, actual.Signature, dcl.DiffInfo{ObjectFunction: compareNoteBuildSignatureNewStyle, EmptyObject: EmptyNoteBuildSignature, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Signature")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNoteBuildSignatureNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NoteBuildSignature)
	if !ok {
		desiredNotPointer, ok := d.(NoteBuildSignature)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteBuildSignature or *NoteBuildSignature", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NoteBuildSignature)
	if !ok {
		actualNotPointer, ok := a.(NoteBuildSignature)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteBuildSignature", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PublicKey, actual.PublicKey, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("PublicKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Signature, actual.Signature, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Signature")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KeyId, actual.KeyId, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("KeyId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KeyType, actual.KeyType, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("KeyType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNoteImageNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NoteImage)
	if !ok {
		desiredNotPointer, ok := d.(NoteImage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteImage or *NoteImage", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NoteImage)
	if !ok {
		actualNotPointer, ok := a.(NoteImage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteImage", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ResourceUrl, actual.ResourceUrl, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("ResourceUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Fingerprint, actual.Fingerprint, dcl.DiffInfo{ObjectFunction: compareNoteImageFingerprintNewStyle, EmptyObject: EmptyNoteImageFingerprint, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Fingerprint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNoteImageFingerprintNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NoteImageFingerprint)
	if !ok {
		desiredNotPointer, ok := d.(NoteImageFingerprint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteImageFingerprint or *NoteImageFingerprint", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NoteImageFingerprint)
	if !ok {
		actualNotPointer, ok := a.(NoteImageFingerprint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteImageFingerprint", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.V1Name, actual.V1Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("V1Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.V2Blob, actual.V2Blob, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("V2Blob")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.V2Name, actual.V2Name, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("V2Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNotePackageNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NotePackage)
	if !ok {
		desiredNotPointer, ok := d.(NotePackage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NotePackage or *NotePackage", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NotePackage)
	if !ok {
		actualNotPointer, ok := a.(NotePackage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NotePackage", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Distribution, actual.Distribution, dcl.DiffInfo{ObjectFunction: compareNotePackageDistributionNewStyle, EmptyObject: EmptyNotePackageDistribution, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Distribution")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNotePackageDistributionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NotePackageDistribution)
	if !ok {
		desiredNotPointer, ok := d.(NotePackageDistribution)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NotePackageDistribution or *NotePackageDistribution", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NotePackageDistribution)
	if !ok {
		actualNotPointer, ok := a.(NotePackageDistribution)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NotePackageDistribution", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.CpeUri, actual.CpeUri, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("CpeUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Architecture, actual.Architecture, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Architecture")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LatestVersion, actual.LatestVersion, dcl.DiffInfo{ObjectFunction: compareNotePackageDistributionLatestVersionNewStyle, EmptyObject: EmptyNotePackageDistributionLatestVersion, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("LatestVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Maintainer, actual.Maintainer, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Maintainer")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Url, actual.Url, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Url")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNotePackageDistributionLatestVersionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NotePackageDistributionLatestVersion)
	if !ok {
		desiredNotPointer, ok := d.(NotePackageDistributionLatestVersion)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NotePackageDistributionLatestVersion or *NotePackageDistributionLatestVersion", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NotePackageDistributionLatestVersion)
	if !ok {
		actualNotPointer, ok := a.(NotePackageDistributionLatestVersion)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NotePackageDistributionLatestVersion", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Epoch, actual.Epoch, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Epoch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Revision, actual.Revision, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Revision")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Kind, actual.Kind, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Kind")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FullName, actual.FullName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("FullName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNoteDiscoveryNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NoteDiscovery)
	if !ok {
		desiredNotPointer, ok := d.(NoteDiscovery)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteDiscovery or *NoteDiscovery", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NoteDiscovery)
	if !ok {
		actualNotPointer, ok := a.(NoteDiscovery)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteDiscovery", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AnalysisKind, actual.AnalysisKind, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("AnalysisKind")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNoteDeploymentNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NoteDeployment)
	if !ok {
		desiredNotPointer, ok := d.(NoteDeployment)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteDeployment or *NoteDeployment", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NoteDeployment)
	if !ok {
		actualNotPointer, ok := a.(NoteDeployment)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteDeployment", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ResourceUri, actual.ResourceUri, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("ResourceUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNoteAttestationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NoteAttestation)
	if !ok {
		desiredNotPointer, ok := d.(NoteAttestation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteAttestation or *NoteAttestation", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NoteAttestation)
	if !ok {
		actualNotPointer, ok := a.(NoteAttestation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteAttestation", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Hint, actual.Hint, dcl.DiffInfo{ObjectFunction: compareNoteAttestationHintNewStyle, EmptyObject: EmptyNoteAttestationHint, OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("Hint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareNoteAttestationHintNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*NoteAttestationHint)
	if !ok {
		desiredNotPointer, ok := d.(NoteAttestationHint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteAttestationHint or *NoteAttestationHint", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*NoteAttestationHint)
	if !ok {
		actualNotPointer, ok := a.(NoteAttestationHint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a NoteAttestationHint", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HumanReadableName, actual.HumanReadableName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateNoteUpdateNoteOperation")}, fn.AddNest("HumanReadableName")); len(ds) != 0 || err != nil {
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
func (r *Note) urlNormalized() *Note {
	normalized := dcl.Copy(*r).(Note)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.ShortDescription = dcl.SelfLinkToName(r.ShortDescription)
	normalized.LongDescription = dcl.SelfLinkToName(r.LongDescription)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *Note) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateNote" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(nr.Project),
			"name":    dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/notes/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Note resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Note) marshal(c *Client) ([]byte, error) {
	m, err := expandNote(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Note: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalNote decodes JSON responses into the Note resource schema.
func unmarshalNote(b []byte, c *Client, res *Note) (*Note, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapNote(m, c, res)
}

func unmarshalMapNote(m map[string]interface{}, c *Client, res *Note) (*Note, error) {

	flattened := flattenNote(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandNote expands Note into a JSON request object.
func expandNote(c *Client, f *Note) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("projects/%s/notes/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.ShortDescription; dcl.ValueShouldBeSent(v) {
		m["shortDescription"] = v
	}
	if v := f.LongDescription; dcl.ValueShouldBeSent(v) {
		m["longDescription"] = v
	}
	if v, err := expandNoteRelatedUrlSlice(c, f.RelatedUrl, res); err != nil {
		return nil, fmt.Errorf("error expanding RelatedUrl into relatedUrl: %w", err)
	} else if v != nil {
		m["relatedUrl"] = v
	}
	if v := f.ExpirationTime; dcl.ValueShouldBeSent(v) {
		m["expirationTime"] = v
	}
	if v := f.RelatedNoteNames; v != nil {
		m["relatedNoteNames"] = v
	}
	if v, err := expandNoteVulnerability(c, f.Vulnerability, res); err != nil {
		return nil, fmt.Errorf("error expanding Vulnerability into vulnerability: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["vulnerability"] = v
	}
	if v, err := expandNoteBuild(c, f.Build, res); err != nil {
		return nil, fmt.Errorf("error expanding Build into build: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["build"] = v
	}
	if v, err := expandNoteImage(c, f.Image, res); err != nil {
		return nil, fmt.Errorf("error expanding Image into baseImage: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["baseImage"] = v
	}
	if v, err := expandNotePackage(c, f.Package, res); err != nil {
		return nil, fmt.Errorf("error expanding Package into package: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["package"] = v
	}
	if v, err := expandNoteDiscovery(c, f.Discovery, res); err != nil {
		return nil, fmt.Errorf("error expanding Discovery into discovery: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["discovery"] = v
	}
	if v, err := expandNoteDeployment(c, f.Deployment, res); err != nil {
		return nil, fmt.Errorf("error expanding Deployment into deployable: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["deployable"] = v
	}
	if v, err := expandNoteAttestation(c, f.Attestation, res); err != nil {
		return nil, fmt.Errorf("error expanding Attestation into attestationAuthority: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["attestationAuthority"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}

	return m, nil
}

// flattenNote flattens Note from a JSON request object into the
// Note type.
func flattenNote(c *Client, i interface{}, res *Note) *Note {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Note{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.ShortDescription = dcl.FlattenString(m["shortDescription"])
	resultRes.LongDescription = dcl.FlattenString(m["longDescription"])
	resultRes.RelatedUrl = flattenNoteRelatedUrlSlice(c, m["relatedUrl"], res)
	resultRes.ExpirationTime = dcl.FlattenString(m["expirationTime"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.RelatedNoteNames = dcl.FlattenStringSlice(m["relatedNoteNames"])
	resultRes.Vulnerability = flattenNoteVulnerability(c, m["vulnerability"], res)
	resultRes.Build = flattenNoteBuild(c, m["build"], res)
	resultRes.Image = flattenNoteImage(c, m["baseImage"], res)
	resultRes.Package = flattenNotePackage(c, m["package"], res)
	resultRes.Discovery = flattenNoteDiscovery(c, m["discovery"], res)
	resultRes.Deployment = flattenNoteDeployment(c, m["deployable"], res)
	resultRes.Attestation = flattenNoteAttestation(c, m["attestationAuthority"], res)
	resultRes.Project = dcl.FlattenString(m["project"])

	return resultRes
}

// expandNoteRelatedUrlMap expands the contents of NoteRelatedUrl into a JSON
// request object.
func expandNoteRelatedUrlMap(c *Client, f map[string]NoteRelatedUrl, res *Note) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteRelatedUrl(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteRelatedUrlSlice expands the contents of NoteRelatedUrl into a JSON
// request object.
func expandNoteRelatedUrlSlice(c *Client, f []NoteRelatedUrl, res *Note) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteRelatedUrl(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteRelatedUrlMap flattens the contents of NoteRelatedUrl from a JSON
// response object.
func flattenNoteRelatedUrlMap(c *Client, i interface{}, res *Note) map[string]NoteRelatedUrl {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteRelatedUrl{}
	}

	if len(a) == 0 {
		return map[string]NoteRelatedUrl{}
	}

	items := make(map[string]NoteRelatedUrl)
	for k, item := range a {
		items[k] = *flattenNoteRelatedUrl(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenNoteRelatedUrlSlice flattens the contents of NoteRelatedUrl from a JSON
// response object.
func flattenNoteRelatedUrlSlice(c *Client, i interface{}, res *Note) []NoteRelatedUrl {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteRelatedUrl{}
	}

	if len(a) == 0 {
		return []NoteRelatedUrl{}
	}

	items := make([]NoteRelatedUrl, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteRelatedUrl(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandNoteRelatedUrl expands an instance of NoteRelatedUrl into a JSON
// request object.
func expandNoteRelatedUrl(c *Client, f *NoteRelatedUrl, res *Note) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Url; !dcl.IsEmptyValueIndirect(v) {
		m["url"] = v
	}
	if v := f.Label; !dcl.IsEmptyValueIndirect(v) {
		m["label"] = v
	}

	return m, nil
}

// flattenNoteRelatedUrl flattens an instance of NoteRelatedUrl from a JSON
// response object.
func flattenNoteRelatedUrl(c *Client, i interface{}, res *Note) *NoteRelatedUrl {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteRelatedUrl{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNoteRelatedUrl
	}
	r.Url = dcl.FlattenString(m["url"])
	r.Label = dcl.FlattenString(m["label"])

	return r
}

// expandNoteVulnerabilityMap expands the contents of NoteVulnerability into a JSON
// request object.
func expandNoteVulnerabilityMap(c *Client, f map[string]NoteVulnerability, res *Note) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteVulnerability(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteVulnerabilitySlice expands the contents of NoteVulnerability into a JSON
// request object.
func expandNoteVulnerabilitySlice(c *Client, f []NoteVulnerability, res *Note) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteVulnerability(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteVulnerabilityMap flattens the contents of NoteVulnerability from a JSON
// response object.
func flattenNoteVulnerabilityMap(c *Client, i interface{}, res *Note) map[string]NoteVulnerability {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerability{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerability{}
	}

	items := make(map[string]NoteVulnerability)
	for k, item := range a {
		items[k] = *flattenNoteVulnerability(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenNoteVulnerabilitySlice flattens the contents of NoteVulnerability from a JSON
// response object.
func flattenNoteVulnerabilitySlice(c *Client, i interface{}, res *Note) []NoteVulnerability {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerability{}
	}

	if len(a) == 0 {
		return []NoteVulnerability{}
	}

	items := make([]NoteVulnerability, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerability(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandNoteVulnerability expands an instance of NoteVulnerability into a JSON
// request object.
func expandNoteVulnerability(c *Client, f *NoteVulnerability, res *Note) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.CvssScore; !dcl.IsEmptyValueIndirect(v) {
		m["cvssScore"] = v
	}
	if v := f.Severity; !dcl.IsEmptyValueIndirect(v) {
		m["severity"] = v
	}
	if v, err := expandNoteVulnerabilityDetailsSlice(c, f.Details, res); err != nil {
		return nil, fmt.Errorf("error expanding Details into details: %w", err)
	} else if v != nil {
		m["details"] = v
	}
	if v, err := expandNoteVulnerabilityCvssV3(c, f.CvssV3, res); err != nil {
		return nil, fmt.Errorf("error expanding CvssV3 into cvssV3: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cvssV3"] = v
	}
	if v, err := expandNoteVulnerabilityWindowsDetailsSlice(c, f.WindowsDetails, res); err != nil {
		return nil, fmt.Errorf("error expanding WindowsDetails into windowsDetails: %w", err)
	} else if v != nil {
		m["windowsDetails"] = v
	}
	if v := f.SourceUpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["sourceUpdateTime"] = v
	}

	return m, nil
}

// flattenNoteVulnerability flattens an instance of NoteVulnerability from a JSON
// response object.
func flattenNoteVulnerability(c *Client, i interface{}, res *Note) *NoteVulnerability {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteVulnerability{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNoteVulnerability
	}
	r.CvssScore = dcl.FlattenDouble(m["cvssScore"])
	r.Severity = flattenNoteVulnerabilitySeverityEnum(m["severity"])
	r.Details = flattenNoteVulnerabilityDetailsSlice(c, m["details"], res)
	r.CvssV3 = flattenNoteVulnerabilityCvssV3(c, m["cvssV3"], res)
	r.WindowsDetails = flattenNoteVulnerabilityWindowsDetailsSlice(c, m["windowsDetails"], res)
	r.SourceUpdateTime = dcl.FlattenString(m["sourceUpdateTime"])

	return r
}

// expandNoteVulnerabilityDetailsMap expands the contents of NoteVulnerabilityDetails into a JSON
// request object.
func expandNoteVulnerabilityDetailsMap(c *Client, f map[string]NoteVulnerabilityDetails, res *Note) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteVulnerabilityDetails(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteVulnerabilityDetailsSlice expands the contents of NoteVulnerabilityDetails into a JSON
// request object.
func expandNoteVulnerabilityDetailsSlice(c *Client, f []NoteVulnerabilityDetails, res *Note) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteVulnerabilityDetails(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteVulnerabilityDetailsMap flattens the contents of NoteVulnerabilityDetails from a JSON
// response object.
func flattenNoteVulnerabilityDetailsMap(c *Client, i interface{}, res *Note) map[string]NoteVulnerabilityDetails {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerabilityDetails{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerabilityDetails{}
	}

	items := make(map[string]NoteVulnerabilityDetails)
	for k, item := range a {
		items[k] = *flattenNoteVulnerabilityDetails(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenNoteVulnerabilityDetailsSlice flattens the contents of NoteVulnerabilityDetails from a JSON
// response object.
func flattenNoteVulnerabilityDetailsSlice(c *Client, i interface{}, res *Note) []NoteVulnerabilityDetails {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityDetails{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityDetails{}
	}

	items := make([]NoteVulnerabilityDetails, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityDetails(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandNoteVulnerabilityDetails expands an instance of NoteVulnerabilityDetails into a JSON
// request object.
func expandNoteVulnerabilityDetails(c *Client, f *NoteVulnerabilityDetails, res *Note) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SeverityName; !dcl.IsEmptyValueIndirect(v) {
		m["severityName"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.PackageType; !dcl.IsEmptyValueIndirect(v) {
		m["packageType"] = v
	}
	if v := f.AffectedCpeUri; !dcl.IsEmptyValueIndirect(v) {
		m["affectedCpeUri"] = v
	}
	if v := f.AffectedPackage; !dcl.IsEmptyValueIndirect(v) {
		m["affectedPackage"] = v
	}
	if v, err := expandNoteVulnerabilityDetailsAffectedVersionStart(c, f.AffectedVersionStart, res); err != nil {
		return nil, fmt.Errorf("error expanding AffectedVersionStart into affectedVersionStart: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["affectedVersionStart"] = v
	}
	if v, err := expandNoteVulnerabilityDetailsAffectedVersionEnd(c, f.AffectedVersionEnd, res); err != nil {
		return nil, fmt.Errorf("error expanding AffectedVersionEnd into affectedVersionEnd: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["affectedVersionEnd"] = v
	}
	if v := f.FixedCpeUri; !dcl.IsEmptyValueIndirect(v) {
		m["fixedCpeUri"] = v
	}
	if v := f.FixedPackage; !dcl.IsEmptyValueIndirect(v) {
		m["fixedPackage"] = v
	}
	if v, err := expandNoteVulnerabilityDetailsFixedVersion(c, f.FixedVersion, res); err != nil {
		return nil, fmt.Errorf("error expanding FixedVersion into fixedVersion: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["fixedVersion"] = v
	}
	if v := f.IsObsolete; !dcl.IsEmptyValueIndirect(v) {
		m["isObsolete"] = v
	}
	if v := f.SourceUpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["sourceUpdateTime"] = v
	}

	return m, nil
}

// flattenNoteVulnerabilityDetails flattens an instance of NoteVulnerabilityDetails from a JSON
// response object.
func flattenNoteVulnerabilityDetails(c *Client, i interface{}, res *Note) *NoteVulnerabilityDetails {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteVulnerabilityDetails{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNoteVulnerabilityDetails
	}
	r.SeverityName = dcl.FlattenString(m["severityName"])
	r.Description = dcl.FlattenString(m["description"])
	r.PackageType = dcl.FlattenString(m["packageType"])
	r.AffectedCpeUri = dcl.FlattenString(m["affectedCpeUri"])
	r.AffectedPackage = dcl.FlattenString(m["affectedPackage"])
	r.AffectedVersionStart = flattenNoteVulnerabilityDetailsAffectedVersionStart(c, m["affectedVersionStart"], res)
	r.AffectedVersionEnd = flattenNoteVulnerabilityDetailsAffectedVersionEnd(c, m["affectedVersionEnd"], res)
	r.FixedCpeUri = dcl.FlattenString(m["fixedCpeUri"])
	r.FixedPackage = dcl.FlattenString(m["fixedPackage"])
	r.FixedVersion = flattenNoteVulnerabilityDetailsFixedVersion(c, m["fixedVersion"], res)
	r.IsObsolete = dcl.FlattenBool(m["isObsolete"])
	r.SourceUpdateTime = dcl.FlattenString(m["sourceUpdateTime"])

	return r
}

// expandNoteVulnerabilityDetailsAffectedVersionStartMap expands the contents of NoteVulnerabilityDetailsAffectedVersionStart into a JSON
// request object.
func expandNoteVulnerabilityDetailsAffectedVersionStartMap(c *Client, f map[string]NoteVulnerabilityDetailsAffectedVersionStart, res *Note) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteVulnerabilityDetailsAffectedVersionStart(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteVulnerabilityDetailsAffectedVersionStartSlice expands the contents of NoteVulnerabilityDetailsAffectedVersionStart into a JSON
// request object.
func expandNoteVulnerabilityDetailsAffectedVersionStartSlice(c *Client, f []NoteVulnerabilityDetailsAffectedVersionStart, res *Note) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteVulnerabilityDetailsAffectedVersionStart(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteVulnerabilityDetailsAffectedVersionStartMap flattens the contents of NoteVulnerabilityDetailsAffectedVersionStart from a JSON
// response object.
func flattenNoteVulnerabilityDetailsAffectedVersionStartMap(c *Client, i interface{}, res *Note) map[string]NoteVulnerabilityDetailsAffectedVersionStart {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerabilityDetailsAffectedVersionStart{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerabilityDetailsAffectedVersionStart{}
	}

	items := make(map[string]NoteVulnerabilityDetailsAffectedVersionStart)
	for k, item := range a {
		items[k] = *flattenNoteVulnerabilityDetailsAffectedVersionStart(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenNoteVulnerabilityDetailsAffectedVersionStartSlice flattens the contents of NoteVulnerabilityDetailsAffectedVersionStart from a JSON
// response object.
func flattenNoteVulnerabilityDetailsAffectedVersionStartSlice(c *Client, i interface{}, res *Note) []NoteVulnerabilityDetailsAffectedVersionStart {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityDetailsAffectedVersionStart{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityDetailsAffectedVersionStart{}
	}

	items := make([]NoteVulnerabilityDetailsAffectedVersionStart, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityDetailsAffectedVersionStart(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandNoteVulnerabilityDetailsAffectedVersionStart expands an instance of NoteVulnerabilityDetailsAffectedVersionStart into a JSON
// request object.
func expandNoteVulnerabilityDetailsAffectedVersionStart(c *Client, f *NoteVulnerabilityDetailsAffectedVersionStart, res *Note) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Epoch; !dcl.IsEmptyValueIndirect(v) {
		m["epoch"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Revision; !dcl.IsEmptyValueIndirect(v) {
		m["revision"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.FullName; !dcl.IsEmptyValueIndirect(v) {
		m["fullName"] = v
	}

	return m, nil
}

// flattenNoteVulnerabilityDetailsAffectedVersionStart flattens an instance of NoteVulnerabilityDetailsAffectedVersionStart from a JSON
// response object.
func flattenNoteVulnerabilityDetailsAffectedVersionStart(c *Client, i interface{}, res *Note) *NoteVulnerabilityDetailsAffectedVersionStart {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteVulnerabilityDetailsAffectedVersionStart{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNoteVulnerabilityDetailsAffectedVersionStart
	}
	r.Epoch = dcl.FlattenInteger(m["epoch"])
	r.Name = dcl.FlattenString(m["name"])
	r.Revision = dcl.FlattenString(m["revision"])
	r.Kind = flattenNoteVulnerabilityDetailsAffectedVersionStartKindEnum(m["kind"])
	r.FullName = dcl.FlattenString(m["fullName"])

	return r
}

// expandNoteVulnerabilityDetailsAffectedVersionEndMap expands the contents of NoteVulnerabilityDetailsAffectedVersionEnd into a JSON
// request object.
func expandNoteVulnerabilityDetailsAffectedVersionEndMap(c *Client, f map[string]NoteVulnerabilityDetailsAffectedVersionEnd, res *Note) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteVulnerabilityDetailsAffectedVersionEnd(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteVulnerabilityDetailsAffectedVersionEndSlice expands the contents of NoteVulnerabilityDetailsAffectedVersionEnd into a JSON
// request object.
func expandNoteVulnerabilityDetailsAffectedVersionEndSlice(c *Client, f []NoteVulnerabilityDetailsAffectedVersionEnd, res *Note) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteVulnerabilityDetailsAffectedVersionEnd(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteVulnerabilityDetailsAffectedVersionEndMap flattens the contents of NoteVulnerabilityDetailsAffectedVersionEnd from a JSON
// response object.
func flattenNoteVulnerabilityDetailsAffectedVersionEndMap(c *Client, i interface{}, res *Note) map[string]NoteVulnerabilityDetailsAffectedVersionEnd {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerabilityDetailsAffectedVersionEnd{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerabilityDetailsAffectedVersionEnd{}
	}

	items := make(map[string]NoteVulnerabilityDetailsAffectedVersionEnd)
	for k, item := range a {
		items[k] = *flattenNoteVulnerabilityDetailsAffectedVersionEnd(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenNoteVulnerabilityDetailsAffectedVersionEndSlice flattens the contents of NoteVulnerabilityDetailsAffectedVersionEnd from a JSON
// response object.
func flattenNoteVulnerabilityDetailsAffectedVersionEndSlice(c *Client, i interface{}, res *Note) []NoteVulnerabilityDetailsAffectedVersionEnd {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityDetailsAffectedVersionEnd{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityDetailsAffectedVersionEnd{}
	}

	items := make([]NoteVulnerabilityDetailsAffectedVersionEnd, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityDetailsAffectedVersionEnd(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandNoteVulnerabilityDetailsAffectedVersionEnd expands an instance of NoteVulnerabilityDetailsAffectedVersionEnd into a JSON
// request object.
func expandNoteVulnerabilityDetailsAffectedVersionEnd(c *Client, f *NoteVulnerabilityDetailsAffectedVersionEnd, res *Note) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Epoch; !dcl.IsEmptyValueIndirect(v) {
		m["epoch"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Revision; !dcl.IsEmptyValueIndirect(v) {
		m["revision"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.FullName; !dcl.IsEmptyValueIndirect(v) {
		m["fullName"] = v
	}

	return m, nil
}

// flattenNoteVulnerabilityDetailsAffectedVersionEnd flattens an instance of NoteVulnerabilityDetailsAffectedVersionEnd from a JSON
// response object.
func flattenNoteVulnerabilityDetailsAffectedVersionEnd(c *Client, i interface{}, res *Note) *NoteVulnerabilityDetailsAffectedVersionEnd {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteVulnerabilityDetailsAffectedVersionEnd{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNoteVulnerabilityDetailsAffectedVersionEnd
	}
	r.Epoch = dcl.FlattenInteger(m["epoch"])
	r.Name = dcl.FlattenString(m["name"])
	r.Revision = dcl.FlattenString(m["revision"])
	r.Kind = flattenNoteVulnerabilityDetailsAffectedVersionEndKindEnum(m["kind"])
	r.FullName = dcl.FlattenString(m["fullName"])

	return r
}

// expandNoteVulnerabilityDetailsFixedVersionMap expands the contents of NoteVulnerabilityDetailsFixedVersion into a JSON
// request object.
func expandNoteVulnerabilityDetailsFixedVersionMap(c *Client, f map[string]NoteVulnerabilityDetailsFixedVersion, res *Note) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteVulnerabilityDetailsFixedVersion(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteVulnerabilityDetailsFixedVersionSlice expands the contents of NoteVulnerabilityDetailsFixedVersion into a JSON
// request object.
func expandNoteVulnerabilityDetailsFixedVersionSlice(c *Client, f []NoteVulnerabilityDetailsFixedVersion, res *Note) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteVulnerabilityDetailsFixedVersion(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteVulnerabilityDetailsFixedVersionMap flattens the contents of NoteVulnerabilityDetailsFixedVersion from a JSON
// response object.
func flattenNoteVulnerabilityDetailsFixedVersionMap(c *Client, i interface{}, res *Note) map[string]NoteVulnerabilityDetailsFixedVersion {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerabilityDetailsFixedVersion{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerabilityDetailsFixedVersion{}
	}

	items := make(map[string]NoteVulnerabilityDetailsFixedVersion)
	for k, item := range a {
		items[k] = *flattenNoteVulnerabilityDetailsFixedVersion(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenNoteVulnerabilityDetailsFixedVersionSlice flattens the contents of NoteVulnerabilityDetailsFixedVersion from a JSON
// response object.
func flattenNoteVulnerabilityDetailsFixedVersionSlice(c *Client, i interface{}, res *Note) []NoteVulnerabilityDetailsFixedVersion {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityDetailsFixedVersion{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityDetailsFixedVersion{}
	}

	items := make([]NoteVulnerabilityDetailsFixedVersion, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityDetailsFixedVersion(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandNoteVulnerabilityDetailsFixedVersion expands an instance of NoteVulnerabilityDetailsFixedVersion into a JSON
// request object.
func expandNoteVulnerabilityDetailsFixedVersion(c *Client, f *NoteVulnerabilityDetailsFixedVersion, res *Note) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Epoch; !dcl.IsEmptyValueIndirect(v) {
		m["epoch"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Revision; !dcl.IsEmptyValueIndirect(v) {
		m["revision"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.FullName; !dcl.IsEmptyValueIndirect(v) {
		m["fullName"] = v
	}

	return m, nil
}

// flattenNoteVulnerabilityDetailsFixedVersion flattens an instance of NoteVulnerabilityDetailsFixedVersion from a JSON
// response object.
func flattenNoteVulnerabilityDetailsFixedVersion(c *Client, i interface{}, res *Note) *NoteVulnerabilityDetailsFixedVersion {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteVulnerabilityDetailsFixedVersion{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNoteVulnerabilityDetailsFixedVersion
	}
	r.Epoch = dcl.FlattenInteger(m["epoch"])
	r.Name = dcl.FlattenString(m["name"])
	r.Revision = dcl.FlattenString(m["revision"])
	r.Kind = flattenNoteVulnerabilityDetailsFixedVersionKindEnum(m["kind"])
	r.FullName = dcl.FlattenString(m["fullName"])

	return r
}

// expandNoteVulnerabilityCvssV3Map expands the contents of NoteVulnerabilityCvssV3 into a JSON
// request object.
func expandNoteVulnerabilityCvssV3Map(c *Client, f map[string]NoteVulnerabilityCvssV3, res *Note) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteVulnerabilityCvssV3(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteVulnerabilityCvssV3Slice expands the contents of NoteVulnerabilityCvssV3 into a JSON
// request object.
func expandNoteVulnerabilityCvssV3Slice(c *Client, f []NoteVulnerabilityCvssV3, res *Note) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteVulnerabilityCvssV3(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteVulnerabilityCvssV3Map flattens the contents of NoteVulnerabilityCvssV3 from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3Map(c *Client, i interface{}, res *Note) map[string]NoteVulnerabilityCvssV3 {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerabilityCvssV3{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerabilityCvssV3{}
	}

	items := make(map[string]NoteVulnerabilityCvssV3)
	for k, item := range a {
		items[k] = *flattenNoteVulnerabilityCvssV3(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenNoteVulnerabilityCvssV3Slice flattens the contents of NoteVulnerabilityCvssV3 from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3Slice(c *Client, i interface{}, res *Note) []NoteVulnerabilityCvssV3 {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityCvssV3{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityCvssV3{}
	}

	items := make([]NoteVulnerabilityCvssV3, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityCvssV3(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandNoteVulnerabilityCvssV3 expands an instance of NoteVulnerabilityCvssV3 into a JSON
// request object.
func expandNoteVulnerabilityCvssV3(c *Client, f *NoteVulnerabilityCvssV3, res *Note) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.BaseScore; !dcl.IsEmptyValueIndirect(v) {
		m["baseScore"] = v
	}
	if v := f.ExploitabilityScore; !dcl.IsEmptyValueIndirect(v) {
		m["exploitabilityScore"] = v
	}
	if v := f.ImpactScore; !dcl.IsEmptyValueIndirect(v) {
		m["impactScore"] = v
	}
	if v := f.AttackVector; !dcl.IsEmptyValueIndirect(v) {
		m["attackVector"] = v
	}
	if v := f.AttackComplexity; !dcl.IsEmptyValueIndirect(v) {
		m["attackComplexity"] = v
	}
	if v := f.PrivilegesRequired; !dcl.IsEmptyValueIndirect(v) {
		m["privilegesRequired"] = v
	}
	if v := f.UserInteraction; !dcl.IsEmptyValueIndirect(v) {
		m["userInteraction"] = v
	}
	if v := f.Scope; !dcl.IsEmptyValueIndirect(v) {
		m["scope"] = v
	}
	if v := f.ConfidentialityImpact; !dcl.IsEmptyValueIndirect(v) {
		m["confidentialityImpact"] = v
	}
	if v := f.IntegrityImpact; !dcl.IsEmptyValueIndirect(v) {
		m["integrityImpact"] = v
	}
	if v := f.AvailabilityImpact; !dcl.IsEmptyValueIndirect(v) {
		m["availabilityImpact"] = v
	}

	return m, nil
}

// flattenNoteVulnerabilityCvssV3 flattens an instance of NoteVulnerabilityCvssV3 from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3(c *Client, i interface{}, res *Note) *NoteVulnerabilityCvssV3 {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteVulnerabilityCvssV3{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNoteVulnerabilityCvssV3
	}
	r.BaseScore = dcl.FlattenDouble(m["baseScore"])
	r.ExploitabilityScore = dcl.FlattenDouble(m["exploitabilityScore"])
	r.ImpactScore = dcl.FlattenDouble(m["impactScore"])
	r.AttackVector = flattenNoteVulnerabilityCvssV3AttackVectorEnum(m["attackVector"])
	r.AttackComplexity = flattenNoteVulnerabilityCvssV3AttackComplexityEnum(m["attackComplexity"])
	r.PrivilegesRequired = flattenNoteVulnerabilityCvssV3PrivilegesRequiredEnum(m["privilegesRequired"])
	r.UserInteraction = flattenNoteVulnerabilityCvssV3UserInteractionEnum(m["userInteraction"])
	r.Scope = flattenNoteVulnerabilityCvssV3ScopeEnum(m["scope"])
	r.ConfidentialityImpact = flattenNoteVulnerabilityCvssV3ConfidentialityImpactEnum(m["confidentialityImpact"])
	r.IntegrityImpact = flattenNoteVulnerabilityCvssV3IntegrityImpactEnum(m["integrityImpact"])
	r.AvailabilityImpact = flattenNoteVulnerabilityCvssV3AvailabilityImpactEnum(m["availabilityImpact"])

	return r
}

// expandNoteVulnerabilityWindowsDetailsMap expands the contents of NoteVulnerabilityWindowsDetails into a JSON
// request object.
func expandNoteVulnerabilityWindowsDetailsMap(c *Client, f map[string]NoteVulnerabilityWindowsDetails, res *Note) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteVulnerabilityWindowsDetails(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteVulnerabilityWindowsDetailsSlice expands the contents of NoteVulnerabilityWindowsDetails into a JSON
// request object.
func expandNoteVulnerabilityWindowsDetailsSlice(c *Client, f []NoteVulnerabilityWindowsDetails, res *Note) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteVulnerabilityWindowsDetails(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteVulnerabilityWindowsDetailsMap flattens the contents of NoteVulnerabilityWindowsDetails from a JSON
// response object.
func flattenNoteVulnerabilityWindowsDetailsMap(c *Client, i interface{}, res *Note) map[string]NoteVulnerabilityWindowsDetails {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerabilityWindowsDetails{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerabilityWindowsDetails{}
	}

	items := make(map[string]NoteVulnerabilityWindowsDetails)
	for k, item := range a {
		items[k] = *flattenNoteVulnerabilityWindowsDetails(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenNoteVulnerabilityWindowsDetailsSlice flattens the contents of NoteVulnerabilityWindowsDetails from a JSON
// response object.
func flattenNoteVulnerabilityWindowsDetailsSlice(c *Client, i interface{}, res *Note) []NoteVulnerabilityWindowsDetails {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityWindowsDetails{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityWindowsDetails{}
	}

	items := make([]NoteVulnerabilityWindowsDetails, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityWindowsDetails(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandNoteVulnerabilityWindowsDetails expands an instance of NoteVulnerabilityWindowsDetails into a JSON
// request object.
func expandNoteVulnerabilityWindowsDetails(c *Client, f *NoteVulnerabilityWindowsDetails, res *Note) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.CpeUri; !dcl.IsEmptyValueIndirect(v) {
		m["cpeUri"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v, err := expandNoteVulnerabilityWindowsDetailsFixingKbsSlice(c, f.FixingKbs, res); err != nil {
		return nil, fmt.Errorf("error expanding FixingKbs into fixingKbs: %w", err)
	} else if v != nil {
		m["fixingKbs"] = v
	}

	return m, nil
}

// flattenNoteVulnerabilityWindowsDetails flattens an instance of NoteVulnerabilityWindowsDetails from a JSON
// response object.
func flattenNoteVulnerabilityWindowsDetails(c *Client, i interface{}, res *Note) *NoteVulnerabilityWindowsDetails {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteVulnerabilityWindowsDetails{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNoteVulnerabilityWindowsDetails
	}
	r.CpeUri = dcl.FlattenString(m["cpeUri"])
	r.Name = dcl.FlattenString(m["name"])
	r.Description = dcl.FlattenString(m["description"])
	r.FixingKbs = flattenNoteVulnerabilityWindowsDetailsFixingKbsSlice(c, m["fixingKbs"], res)

	return r
}

// expandNoteVulnerabilityWindowsDetailsFixingKbsMap expands the contents of NoteVulnerabilityWindowsDetailsFixingKbs into a JSON
// request object.
func expandNoteVulnerabilityWindowsDetailsFixingKbsMap(c *Client, f map[string]NoteVulnerabilityWindowsDetailsFixingKbs, res *Note) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteVulnerabilityWindowsDetailsFixingKbs(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteVulnerabilityWindowsDetailsFixingKbsSlice expands the contents of NoteVulnerabilityWindowsDetailsFixingKbs into a JSON
// request object.
func expandNoteVulnerabilityWindowsDetailsFixingKbsSlice(c *Client, f []NoteVulnerabilityWindowsDetailsFixingKbs, res *Note) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteVulnerabilityWindowsDetailsFixingKbs(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteVulnerabilityWindowsDetailsFixingKbsMap flattens the contents of NoteVulnerabilityWindowsDetailsFixingKbs from a JSON
// response object.
func flattenNoteVulnerabilityWindowsDetailsFixingKbsMap(c *Client, i interface{}, res *Note) map[string]NoteVulnerabilityWindowsDetailsFixingKbs {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerabilityWindowsDetailsFixingKbs{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerabilityWindowsDetailsFixingKbs{}
	}

	items := make(map[string]NoteVulnerabilityWindowsDetailsFixingKbs)
	for k, item := range a {
		items[k] = *flattenNoteVulnerabilityWindowsDetailsFixingKbs(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenNoteVulnerabilityWindowsDetailsFixingKbsSlice flattens the contents of NoteVulnerabilityWindowsDetailsFixingKbs from a JSON
// response object.
func flattenNoteVulnerabilityWindowsDetailsFixingKbsSlice(c *Client, i interface{}, res *Note) []NoteVulnerabilityWindowsDetailsFixingKbs {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityWindowsDetailsFixingKbs{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityWindowsDetailsFixingKbs{}
	}

	items := make([]NoteVulnerabilityWindowsDetailsFixingKbs, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityWindowsDetailsFixingKbs(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandNoteVulnerabilityWindowsDetailsFixingKbs expands an instance of NoteVulnerabilityWindowsDetailsFixingKbs into a JSON
// request object.
func expandNoteVulnerabilityWindowsDetailsFixingKbs(c *Client, f *NoteVulnerabilityWindowsDetailsFixingKbs, res *Note) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Url; !dcl.IsEmptyValueIndirect(v) {
		m["url"] = v
	}

	return m, nil
}

// flattenNoteVulnerabilityWindowsDetailsFixingKbs flattens an instance of NoteVulnerabilityWindowsDetailsFixingKbs from a JSON
// response object.
func flattenNoteVulnerabilityWindowsDetailsFixingKbs(c *Client, i interface{}, res *Note) *NoteVulnerabilityWindowsDetailsFixingKbs {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteVulnerabilityWindowsDetailsFixingKbs{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNoteVulnerabilityWindowsDetailsFixingKbs
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Url = dcl.FlattenString(m["url"])

	return r
}

// expandNoteBuildMap expands the contents of NoteBuild into a JSON
// request object.
func expandNoteBuildMap(c *Client, f map[string]NoteBuild, res *Note) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteBuild(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteBuildSlice expands the contents of NoteBuild into a JSON
// request object.
func expandNoteBuildSlice(c *Client, f []NoteBuild, res *Note) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteBuild(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteBuildMap flattens the contents of NoteBuild from a JSON
// response object.
func flattenNoteBuildMap(c *Client, i interface{}, res *Note) map[string]NoteBuild {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteBuild{}
	}

	if len(a) == 0 {
		return map[string]NoteBuild{}
	}

	items := make(map[string]NoteBuild)
	for k, item := range a {
		items[k] = *flattenNoteBuild(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenNoteBuildSlice flattens the contents of NoteBuild from a JSON
// response object.
func flattenNoteBuildSlice(c *Client, i interface{}, res *Note) []NoteBuild {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteBuild{}
	}

	if len(a) == 0 {
		return []NoteBuild{}
	}

	items := make([]NoteBuild, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteBuild(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandNoteBuild expands an instance of NoteBuild into a JSON
// request object.
func expandNoteBuild(c *Client, f *NoteBuild, res *Note) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.BuilderVersion; !dcl.IsEmptyValueIndirect(v) {
		m["builderVersion"] = v
	}
	if v, err := expandNoteBuildSignature(c, f.Signature, res); err != nil {
		return nil, fmt.Errorf("error expanding Signature into signature: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["signature"] = v
	}

	return m, nil
}

// flattenNoteBuild flattens an instance of NoteBuild from a JSON
// response object.
func flattenNoteBuild(c *Client, i interface{}, res *Note) *NoteBuild {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteBuild{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNoteBuild
	}
	r.BuilderVersion = dcl.FlattenString(m["builderVersion"])
	r.Signature = flattenNoteBuildSignature(c, m["signature"], res)

	return r
}

// expandNoteBuildSignatureMap expands the contents of NoteBuildSignature into a JSON
// request object.
func expandNoteBuildSignatureMap(c *Client, f map[string]NoteBuildSignature, res *Note) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteBuildSignature(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteBuildSignatureSlice expands the contents of NoteBuildSignature into a JSON
// request object.
func expandNoteBuildSignatureSlice(c *Client, f []NoteBuildSignature, res *Note) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteBuildSignature(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteBuildSignatureMap flattens the contents of NoteBuildSignature from a JSON
// response object.
func flattenNoteBuildSignatureMap(c *Client, i interface{}, res *Note) map[string]NoteBuildSignature {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteBuildSignature{}
	}

	if len(a) == 0 {
		return map[string]NoteBuildSignature{}
	}

	items := make(map[string]NoteBuildSignature)
	for k, item := range a {
		items[k] = *flattenNoteBuildSignature(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenNoteBuildSignatureSlice flattens the contents of NoteBuildSignature from a JSON
// response object.
func flattenNoteBuildSignatureSlice(c *Client, i interface{}, res *Note) []NoteBuildSignature {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteBuildSignature{}
	}

	if len(a) == 0 {
		return []NoteBuildSignature{}
	}

	items := make([]NoteBuildSignature, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteBuildSignature(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandNoteBuildSignature expands an instance of NoteBuildSignature into a JSON
// request object.
func expandNoteBuildSignature(c *Client, f *NoteBuildSignature, res *Note) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.PublicKey; !dcl.IsEmptyValueIndirect(v) {
		m["publicKey"] = v
	}
	if v := f.Signature; !dcl.IsEmptyValueIndirect(v) {
		m["signature"] = v
	}
	if v := f.KeyId; !dcl.IsEmptyValueIndirect(v) {
		m["keyId"] = v
	}
	if v := f.KeyType; !dcl.IsEmptyValueIndirect(v) {
		m["keyType"] = v
	}

	return m, nil
}

// flattenNoteBuildSignature flattens an instance of NoteBuildSignature from a JSON
// response object.
func flattenNoteBuildSignature(c *Client, i interface{}, res *Note) *NoteBuildSignature {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteBuildSignature{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNoteBuildSignature
	}
	r.PublicKey = dcl.FlattenString(m["publicKey"])
	r.Signature = dcl.FlattenString(m["signature"])
	r.KeyId = dcl.FlattenString(m["keyId"])
	r.KeyType = flattenNoteBuildSignatureKeyTypeEnum(m["keyType"])

	return r
}

// expandNoteImageMap expands the contents of NoteImage into a JSON
// request object.
func expandNoteImageMap(c *Client, f map[string]NoteImage, res *Note) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteImage(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteImageSlice expands the contents of NoteImage into a JSON
// request object.
func expandNoteImageSlice(c *Client, f []NoteImage, res *Note) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteImage(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteImageMap flattens the contents of NoteImage from a JSON
// response object.
func flattenNoteImageMap(c *Client, i interface{}, res *Note) map[string]NoteImage {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteImage{}
	}

	if len(a) == 0 {
		return map[string]NoteImage{}
	}

	items := make(map[string]NoteImage)
	for k, item := range a {
		items[k] = *flattenNoteImage(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenNoteImageSlice flattens the contents of NoteImage from a JSON
// response object.
func flattenNoteImageSlice(c *Client, i interface{}, res *Note) []NoteImage {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteImage{}
	}

	if len(a) == 0 {
		return []NoteImage{}
	}

	items := make([]NoteImage, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteImage(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandNoteImage expands an instance of NoteImage into a JSON
// request object.
func expandNoteImage(c *Client, f *NoteImage, res *Note) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ResourceUrl; !dcl.IsEmptyValueIndirect(v) {
		m["resourceUrl"] = v
	}
	if v, err := expandNoteImageFingerprint(c, f.Fingerprint, res); err != nil {
		return nil, fmt.Errorf("error expanding Fingerprint into fingerprint: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["fingerprint"] = v
	}

	return m, nil
}

// flattenNoteImage flattens an instance of NoteImage from a JSON
// response object.
func flattenNoteImage(c *Client, i interface{}, res *Note) *NoteImage {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteImage{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNoteImage
	}
	r.ResourceUrl = dcl.FlattenString(m["resourceUrl"])
	r.Fingerprint = flattenNoteImageFingerprint(c, m["fingerprint"], res)

	return r
}

// expandNoteImageFingerprintMap expands the contents of NoteImageFingerprint into a JSON
// request object.
func expandNoteImageFingerprintMap(c *Client, f map[string]NoteImageFingerprint, res *Note) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteImageFingerprint(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteImageFingerprintSlice expands the contents of NoteImageFingerprint into a JSON
// request object.
func expandNoteImageFingerprintSlice(c *Client, f []NoteImageFingerprint, res *Note) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteImageFingerprint(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteImageFingerprintMap flattens the contents of NoteImageFingerprint from a JSON
// response object.
func flattenNoteImageFingerprintMap(c *Client, i interface{}, res *Note) map[string]NoteImageFingerprint {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteImageFingerprint{}
	}

	if len(a) == 0 {
		return map[string]NoteImageFingerprint{}
	}

	items := make(map[string]NoteImageFingerprint)
	for k, item := range a {
		items[k] = *flattenNoteImageFingerprint(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenNoteImageFingerprintSlice flattens the contents of NoteImageFingerprint from a JSON
// response object.
func flattenNoteImageFingerprintSlice(c *Client, i interface{}, res *Note) []NoteImageFingerprint {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteImageFingerprint{}
	}

	if len(a) == 0 {
		return []NoteImageFingerprint{}
	}

	items := make([]NoteImageFingerprint, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteImageFingerprint(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandNoteImageFingerprint expands an instance of NoteImageFingerprint into a JSON
// request object.
func expandNoteImageFingerprint(c *Client, f *NoteImageFingerprint, res *Note) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.V1Name; !dcl.IsEmptyValueIndirect(v) {
		m["v1Name"] = v
	}
	if v := f.V2Blob; v != nil {
		m["v2Blob"] = v
	}

	return m, nil
}

// flattenNoteImageFingerprint flattens an instance of NoteImageFingerprint from a JSON
// response object.
func flattenNoteImageFingerprint(c *Client, i interface{}, res *Note) *NoteImageFingerprint {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteImageFingerprint{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNoteImageFingerprint
	}
	r.V1Name = dcl.FlattenString(m["v1Name"])
	r.V2Blob = dcl.FlattenStringSlice(m["v2Blob"])
	r.V2Name = dcl.FlattenString(m["v2Name"])

	return r
}

// expandNotePackageMap expands the contents of NotePackage into a JSON
// request object.
func expandNotePackageMap(c *Client, f map[string]NotePackage, res *Note) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNotePackage(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNotePackageSlice expands the contents of NotePackage into a JSON
// request object.
func expandNotePackageSlice(c *Client, f []NotePackage, res *Note) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNotePackage(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNotePackageMap flattens the contents of NotePackage from a JSON
// response object.
func flattenNotePackageMap(c *Client, i interface{}, res *Note) map[string]NotePackage {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NotePackage{}
	}

	if len(a) == 0 {
		return map[string]NotePackage{}
	}

	items := make(map[string]NotePackage)
	for k, item := range a {
		items[k] = *flattenNotePackage(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenNotePackageSlice flattens the contents of NotePackage from a JSON
// response object.
func flattenNotePackageSlice(c *Client, i interface{}, res *Note) []NotePackage {
	a, ok := i.([]interface{})
	if !ok {
		return []NotePackage{}
	}

	if len(a) == 0 {
		return []NotePackage{}
	}

	items := make([]NotePackage, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNotePackage(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandNotePackage expands an instance of NotePackage into a JSON
// request object.
func expandNotePackage(c *Client, f *NotePackage, res *Note) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandNotePackageDistributionSlice(c, f.Distribution, res); err != nil {
		return nil, fmt.Errorf("error expanding Distribution into distribution: %w", err)
	} else if v != nil {
		m["distribution"] = v
	}

	return m, nil
}

// flattenNotePackage flattens an instance of NotePackage from a JSON
// response object.
func flattenNotePackage(c *Client, i interface{}, res *Note) *NotePackage {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NotePackage{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNotePackage
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Distribution = flattenNotePackageDistributionSlice(c, m["distribution"], res)

	return r
}

// expandNotePackageDistributionMap expands the contents of NotePackageDistribution into a JSON
// request object.
func expandNotePackageDistributionMap(c *Client, f map[string]NotePackageDistribution, res *Note) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNotePackageDistribution(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNotePackageDistributionSlice expands the contents of NotePackageDistribution into a JSON
// request object.
func expandNotePackageDistributionSlice(c *Client, f []NotePackageDistribution, res *Note) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNotePackageDistribution(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNotePackageDistributionMap flattens the contents of NotePackageDistribution from a JSON
// response object.
func flattenNotePackageDistributionMap(c *Client, i interface{}, res *Note) map[string]NotePackageDistribution {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NotePackageDistribution{}
	}

	if len(a) == 0 {
		return map[string]NotePackageDistribution{}
	}

	items := make(map[string]NotePackageDistribution)
	for k, item := range a {
		items[k] = *flattenNotePackageDistribution(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenNotePackageDistributionSlice flattens the contents of NotePackageDistribution from a JSON
// response object.
func flattenNotePackageDistributionSlice(c *Client, i interface{}, res *Note) []NotePackageDistribution {
	a, ok := i.([]interface{})
	if !ok {
		return []NotePackageDistribution{}
	}

	if len(a) == 0 {
		return []NotePackageDistribution{}
	}

	items := make([]NotePackageDistribution, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNotePackageDistribution(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandNotePackageDistribution expands an instance of NotePackageDistribution into a JSON
// request object.
func expandNotePackageDistribution(c *Client, f *NotePackageDistribution, res *Note) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.CpeUri; !dcl.IsEmptyValueIndirect(v) {
		m["cpeUri"] = v
	}
	if v := f.Architecture; !dcl.IsEmptyValueIndirect(v) {
		m["architecture"] = v
	}
	if v, err := expandNotePackageDistributionLatestVersion(c, f.LatestVersion, res); err != nil {
		return nil, fmt.Errorf("error expanding LatestVersion into latestVersion: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["latestVersion"] = v
	}
	if v := f.Maintainer; !dcl.IsEmptyValueIndirect(v) {
		m["maintainer"] = v
	}
	if v := f.Url; !dcl.IsEmptyValueIndirect(v) {
		m["url"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}

	return m, nil
}

// flattenNotePackageDistribution flattens an instance of NotePackageDistribution from a JSON
// response object.
func flattenNotePackageDistribution(c *Client, i interface{}, res *Note) *NotePackageDistribution {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NotePackageDistribution{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNotePackageDistribution
	}
	r.CpeUri = dcl.FlattenString(m["cpeUri"])
	r.Architecture = flattenNotePackageDistributionArchitectureEnum(m["architecture"])
	r.LatestVersion = flattenNotePackageDistributionLatestVersion(c, m["latestVersion"], res)
	r.Maintainer = dcl.FlattenString(m["maintainer"])
	r.Url = dcl.FlattenString(m["url"])
	r.Description = dcl.FlattenString(m["description"])

	return r
}

// expandNotePackageDistributionLatestVersionMap expands the contents of NotePackageDistributionLatestVersion into a JSON
// request object.
func expandNotePackageDistributionLatestVersionMap(c *Client, f map[string]NotePackageDistributionLatestVersion, res *Note) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNotePackageDistributionLatestVersion(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNotePackageDistributionLatestVersionSlice expands the contents of NotePackageDistributionLatestVersion into a JSON
// request object.
func expandNotePackageDistributionLatestVersionSlice(c *Client, f []NotePackageDistributionLatestVersion, res *Note) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNotePackageDistributionLatestVersion(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNotePackageDistributionLatestVersionMap flattens the contents of NotePackageDistributionLatestVersion from a JSON
// response object.
func flattenNotePackageDistributionLatestVersionMap(c *Client, i interface{}, res *Note) map[string]NotePackageDistributionLatestVersion {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NotePackageDistributionLatestVersion{}
	}

	if len(a) == 0 {
		return map[string]NotePackageDistributionLatestVersion{}
	}

	items := make(map[string]NotePackageDistributionLatestVersion)
	for k, item := range a {
		items[k] = *flattenNotePackageDistributionLatestVersion(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenNotePackageDistributionLatestVersionSlice flattens the contents of NotePackageDistributionLatestVersion from a JSON
// response object.
func flattenNotePackageDistributionLatestVersionSlice(c *Client, i interface{}, res *Note) []NotePackageDistributionLatestVersion {
	a, ok := i.([]interface{})
	if !ok {
		return []NotePackageDistributionLatestVersion{}
	}

	if len(a) == 0 {
		return []NotePackageDistributionLatestVersion{}
	}

	items := make([]NotePackageDistributionLatestVersion, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNotePackageDistributionLatestVersion(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandNotePackageDistributionLatestVersion expands an instance of NotePackageDistributionLatestVersion into a JSON
// request object.
func expandNotePackageDistributionLatestVersion(c *Client, f *NotePackageDistributionLatestVersion, res *Note) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Epoch; !dcl.IsEmptyValueIndirect(v) {
		m["epoch"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Revision; !dcl.IsEmptyValueIndirect(v) {
		m["revision"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.FullName; !dcl.IsEmptyValueIndirect(v) {
		m["fullName"] = v
	}

	return m, nil
}

// flattenNotePackageDistributionLatestVersion flattens an instance of NotePackageDistributionLatestVersion from a JSON
// response object.
func flattenNotePackageDistributionLatestVersion(c *Client, i interface{}, res *Note) *NotePackageDistributionLatestVersion {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NotePackageDistributionLatestVersion{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNotePackageDistributionLatestVersion
	}
	r.Epoch = dcl.FlattenInteger(m["epoch"])
	r.Name = dcl.FlattenString(m["name"])
	r.Revision = dcl.FlattenString(m["revision"])
	r.Kind = flattenNotePackageDistributionLatestVersionKindEnum(m["kind"])
	r.FullName = dcl.FlattenString(m["fullName"])

	return r
}

// expandNoteDiscoveryMap expands the contents of NoteDiscovery into a JSON
// request object.
func expandNoteDiscoveryMap(c *Client, f map[string]NoteDiscovery, res *Note) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteDiscovery(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteDiscoverySlice expands the contents of NoteDiscovery into a JSON
// request object.
func expandNoteDiscoverySlice(c *Client, f []NoteDiscovery, res *Note) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteDiscovery(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteDiscoveryMap flattens the contents of NoteDiscovery from a JSON
// response object.
func flattenNoteDiscoveryMap(c *Client, i interface{}, res *Note) map[string]NoteDiscovery {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteDiscovery{}
	}

	if len(a) == 0 {
		return map[string]NoteDiscovery{}
	}

	items := make(map[string]NoteDiscovery)
	for k, item := range a {
		items[k] = *flattenNoteDiscovery(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenNoteDiscoverySlice flattens the contents of NoteDiscovery from a JSON
// response object.
func flattenNoteDiscoverySlice(c *Client, i interface{}, res *Note) []NoteDiscovery {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteDiscovery{}
	}

	if len(a) == 0 {
		return []NoteDiscovery{}
	}

	items := make([]NoteDiscovery, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteDiscovery(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandNoteDiscovery expands an instance of NoteDiscovery into a JSON
// request object.
func expandNoteDiscovery(c *Client, f *NoteDiscovery, res *Note) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AnalysisKind; !dcl.IsEmptyValueIndirect(v) {
		m["analysisKind"] = v
	}

	return m, nil
}

// flattenNoteDiscovery flattens an instance of NoteDiscovery from a JSON
// response object.
func flattenNoteDiscovery(c *Client, i interface{}, res *Note) *NoteDiscovery {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteDiscovery{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNoteDiscovery
	}
	r.AnalysisKind = flattenNoteDiscoveryAnalysisKindEnum(m["analysisKind"])

	return r
}

// expandNoteDeploymentMap expands the contents of NoteDeployment into a JSON
// request object.
func expandNoteDeploymentMap(c *Client, f map[string]NoteDeployment, res *Note) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteDeployment(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteDeploymentSlice expands the contents of NoteDeployment into a JSON
// request object.
func expandNoteDeploymentSlice(c *Client, f []NoteDeployment, res *Note) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteDeployment(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteDeploymentMap flattens the contents of NoteDeployment from a JSON
// response object.
func flattenNoteDeploymentMap(c *Client, i interface{}, res *Note) map[string]NoteDeployment {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteDeployment{}
	}

	if len(a) == 0 {
		return map[string]NoteDeployment{}
	}

	items := make(map[string]NoteDeployment)
	for k, item := range a {
		items[k] = *flattenNoteDeployment(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenNoteDeploymentSlice flattens the contents of NoteDeployment from a JSON
// response object.
func flattenNoteDeploymentSlice(c *Client, i interface{}, res *Note) []NoteDeployment {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteDeployment{}
	}

	if len(a) == 0 {
		return []NoteDeployment{}
	}

	items := make([]NoteDeployment, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteDeployment(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandNoteDeployment expands an instance of NoteDeployment into a JSON
// request object.
func expandNoteDeployment(c *Client, f *NoteDeployment, res *Note) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ResourceUri; v != nil {
		m["resourceUri"] = v
	}

	return m, nil
}

// flattenNoteDeployment flattens an instance of NoteDeployment from a JSON
// response object.
func flattenNoteDeployment(c *Client, i interface{}, res *Note) *NoteDeployment {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteDeployment{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNoteDeployment
	}
	r.ResourceUri = dcl.FlattenStringSlice(m["resourceUri"])

	return r
}

// expandNoteAttestationMap expands the contents of NoteAttestation into a JSON
// request object.
func expandNoteAttestationMap(c *Client, f map[string]NoteAttestation, res *Note) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteAttestation(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteAttestationSlice expands the contents of NoteAttestation into a JSON
// request object.
func expandNoteAttestationSlice(c *Client, f []NoteAttestation, res *Note) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteAttestation(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteAttestationMap flattens the contents of NoteAttestation from a JSON
// response object.
func flattenNoteAttestationMap(c *Client, i interface{}, res *Note) map[string]NoteAttestation {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteAttestation{}
	}

	if len(a) == 0 {
		return map[string]NoteAttestation{}
	}

	items := make(map[string]NoteAttestation)
	for k, item := range a {
		items[k] = *flattenNoteAttestation(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenNoteAttestationSlice flattens the contents of NoteAttestation from a JSON
// response object.
func flattenNoteAttestationSlice(c *Client, i interface{}, res *Note) []NoteAttestation {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteAttestation{}
	}

	if len(a) == 0 {
		return []NoteAttestation{}
	}

	items := make([]NoteAttestation, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteAttestation(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandNoteAttestation expands an instance of NoteAttestation into a JSON
// request object.
func expandNoteAttestation(c *Client, f *NoteAttestation, res *Note) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandNoteAttestationHint(c, f.Hint, res); err != nil {
		return nil, fmt.Errorf("error expanding Hint into hint: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["hint"] = v
	}

	return m, nil
}

// flattenNoteAttestation flattens an instance of NoteAttestation from a JSON
// response object.
func flattenNoteAttestation(c *Client, i interface{}, res *Note) *NoteAttestation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteAttestation{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNoteAttestation
	}
	r.Hint = flattenNoteAttestationHint(c, m["hint"], res)

	return r
}

// expandNoteAttestationHintMap expands the contents of NoteAttestationHint into a JSON
// request object.
func expandNoteAttestationHintMap(c *Client, f map[string]NoteAttestationHint, res *Note) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandNoteAttestationHint(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandNoteAttestationHintSlice expands the contents of NoteAttestationHint into a JSON
// request object.
func expandNoteAttestationHintSlice(c *Client, f []NoteAttestationHint, res *Note) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandNoteAttestationHint(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenNoteAttestationHintMap flattens the contents of NoteAttestationHint from a JSON
// response object.
func flattenNoteAttestationHintMap(c *Client, i interface{}, res *Note) map[string]NoteAttestationHint {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteAttestationHint{}
	}

	if len(a) == 0 {
		return map[string]NoteAttestationHint{}
	}

	items := make(map[string]NoteAttestationHint)
	for k, item := range a {
		items[k] = *flattenNoteAttestationHint(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenNoteAttestationHintSlice flattens the contents of NoteAttestationHint from a JSON
// response object.
func flattenNoteAttestationHintSlice(c *Client, i interface{}, res *Note) []NoteAttestationHint {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteAttestationHint{}
	}

	if len(a) == 0 {
		return []NoteAttestationHint{}
	}

	items := make([]NoteAttestationHint, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteAttestationHint(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandNoteAttestationHint expands an instance of NoteAttestationHint into a JSON
// request object.
func expandNoteAttestationHint(c *Client, f *NoteAttestationHint, res *Note) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.HumanReadableName; !dcl.IsEmptyValueIndirect(v) {
		m["humanReadableName"] = v
	}

	return m, nil
}

// flattenNoteAttestationHint flattens an instance of NoteAttestationHint from a JSON
// response object.
func flattenNoteAttestationHint(c *Client, i interface{}, res *Note) *NoteAttestationHint {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &NoteAttestationHint{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyNoteAttestationHint
	}
	r.HumanReadableName = dcl.FlattenString(m["humanReadableName"])

	return r
}

// flattenNoteVulnerabilitySeverityEnumMap flattens the contents of NoteVulnerabilitySeverityEnum from a JSON
// response object.
func flattenNoteVulnerabilitySeverityEnumMap(c *Client, i interface{}, res *Note) map[string]NoteVulnerabilitySeverityEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerabilitySeverityEnum{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerabilitySeverityEnum{}
	}

	items := make(map[string]NoteVulnerabilitySeverityEnum)
	for k, item := range a {
		items[k] = *flattenNoteVulnerabilitySeverityEnum(item.(interface{}))
	}

	return items
}

// flattenNoteVulnerabilitySeverityEnumSlice flattens the contents of NoteVulnerabilitySeverityEnum from a JSON
// response object.
func flattenNoteVulnerabilitySeverityEnumSlice(c *Client, i interface{}, res *Note) []NoteVulnerabilitySeverityEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilitySeverityEnum{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilitySeverityEnum{}
	}

	items := make([]NoteVulnerabilitySeverityEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilitySeverityEnum(item.(interface{})))
	}

	return items
}

// flattenNoteVulnerabilitySeverityEnum asserts that an interface is a string, and returns a
// pointer to a *NoteVulnerabilitySeverityEnum with the same value as that string.
func flattenNoteVulnerabilitySeverityEnum(i interface{}) *NoteVulnerabilitySeverityEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return NoteVulnerabilitySeverityEnumRef(s)
}

// flattenNoteVulnerabilityDetailsAffectedVersionStartKindEnumMap flattens the contents of NoteVulnerabilityDetailsAffectedVersionStartKindEnum from a JSON
// response object.
func flattenNoteVulnerabilityDetailsAffectedVersionStartKindEnumMap(c *Client, i interface{}, res *Note) map[string]NoteVulnerabilityDetailsAffectedVersionStartKindEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerabilityDetailsAffectedVersionStartKindEnum{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerabilityDetailsAffectedVersionStartKindEnum{}
	}

	items := make(map[string]NoteVulnerabilityDetailsAffectedVersionStartKindEnum)
	for k, item := range a {
		items[k] = *flattenNoteVulnerabilityDetailsAffectedVersionStartKindEnum(item.(interface{}))
	}

	return items
}

// flattenNoteVulnerabilityDetailsAffectedVersionStartKindEnumSlice flattens the contents of NoteVulnerabilityDetailsAffectedVersionStartKindEnum from a JSON
// response object.
func flattenNoteVulnerabilityDetailsAffectedVersionStartKindEnumSlice(c *Client, i interface{}, res *Note) []NoteVulnerabilityDetailsAffectedVersionStartKindEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityDetailsAffectedVersionStartKindEnum{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityDetailsAffectedVersionStartKindEnum{}
	}

	items := make([]NoteVulnerabilityDetailsAffectedVersionStartKindEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityDetailsAffectedVersionStartKindEnum(item.(interface{})))
	}

	return items
}

// flattenNoteVulnerabilityDetailsAffectedVersionStartKindEnum asserts that an interface is a string, and returns a
// pointer to a *NoteVulnerabilityDetailsAffectedVersionStartKindEnum with the same value as that string.
func flattenNoteVulnerabilityDetailsAffectedVersionStartKindEnum(i interface{}) *NoteVulnerabilityDetailsAffectedVersionStartKindEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return NoteVulnerabilityDetailsAffectedVersionStartKindEnumRef(s)
}

// flattenNoteVulnerabilityDetailsAffectedVersionEndKindEnumMap flattens the contents of NoteVulnerabilityDetailsAffectedVersionEndKindEnum from a JSON
// response object.
func flattenNoteVulnerabilityDetailsAffectedVersionEndKindEnumMap(c *Client, i interface{}, res *Note) map[string]NoteVulnerabilityDetailsAffectedVersionEndKindEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerabilityDetailsAffectedVersionEndKindEnum{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerabilityDetailsAffectedVersionEndKindEnum{}
	}

	items := make(map[string]NoteVulnerabilityDetailsAffectedVersionEndKindEnum)
	for k, item := range a {
		items[k] = *flattenNoteVulnerabilityDetailsAffectedVersionEndKindEnum(item.(interface{}))
	}

	return items
}

// flattenNoteVulnerabilityDetailsAffectedVersionEndKindEnumSlice flattens the contents of NoteVulnerabilityDetailsAffectedVersionEndKindEnum from a JSON
// response object.
func flattenNoteVulnerabilityDetailsAffectedVersionEndKindEnumSlice(c *Client, i interface{}, res *Note) []NoteVulnerabilityDetailsAffectedVersionEndKindEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityDetailsAffectedVersionEndKindEnum{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityDetailsAffectedVersionEndKindEnum{}
	}

	items := make([]NoteVulnerabilityDetailsAffectedVersionEndKindEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityDetailsAffectedVersionEndKindEnum(item.(interface{})))
	}

	return items
}

// flattenNoteVulnerabilityDetailsAffectedVersionEndKindEnum asserts that an interface is a string, and returns a
// pointer to a *NoteVulnerabilityDetailsAffectedVersionEndKindEnum with the same value as that string.
func flattenNoteVulnerabilityDetailsAffectedVersionEndKindEnum(i interface{}) *NoteVulnerabilityDetailsAffectedVersionEndKindEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return NoteVulnerabilityDetailsAffectedVersionEndKindEnumRef(s)
}

// flattenNoteVulnerabilityDetailsFixedVersionKindEnumMap flattens the contents of NoteVulnerabilityDetailsFixedVersionKindEnum from a JSON
// response object.
func flattenNoteVulnerabilityDetailsFixedVersionKindEnumMap(c *Client, i interface{}, res *Note) map[string]NoteVulnerabilityDetailsFixedVersionKindEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerabilityDetailsFixedVersionKindEnum{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerabilityDetailsFixedVersionKindEnum{}
	}

	items := make(map[string]NoteVulnerabilityDetailsFixedVersionKindEnum)
	for k, item := range a {
		items[k] = *flattenNoteVulnerabilityDetailsFixedVersionKindEnum(item.(interface{}))
	}

	return items
}

// flattenNoteVulnerabilityDetailsFixedVersionKindEnumSlice flattens the contents of NoteVulnerabilityDetailsFixedVersionKindEnum from a JSON
// response object.
func flattenNoteVulnerabilityDetailsFixedVersionKindEnumSlice(c *Client, i interface{}, res *Note) []NoteVulnerabilityDetailsFixedVersionKindEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityDetailsFixedVersionKindEnum{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityDetailsFixedVersionKindEnum{}
	}

	items := make([]NoteVulnerabilityDetailsFixedVersionKindEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityDetailsFixedVersionKindEnum(item.(interface{})))
	}

	return items
}

// flattenNoteVulnerabilityDetailsFixedVersionKindEnum asserts that an interface is a string, and returns a
// pointer to a *NoteVulnerabilityDetailsFixedVersionKindEnum with the same value as that string.
func flattenNoteVulnerabilityDetailsFixedVersionKindEnum(i interface{}) *NoteVulnerabilityDetailsFixedVersionKindEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return NoteVulnerabilityDetailsFixedVersionKindEnumRef(s)
}

// flattenNoteVulnerabilityCvssV3AttackVectorEnumMap flattens the contents of NoteVulnerabilityCvssV3AttackVectorEnum from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3AttackVectorEnumMap(c *Client, i interface{}, res *Note) map[string]NoteVulnerabilityCvssV3AttackVectorEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerabilityCvssV3AttackVectorEnum{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerabilityCvssV3AttackVectorEnum{}
	}

	items := make(map[string]NoteVulnerabilityCvssV3AttackVectorEnum)
	for k, item := range a {
		items[k] = *flattenNoteVulnerabilityCvssV3AttackVectorEnum(item.(interface{}))
	}

	return items
}

// flattenNoteVulnerabilityCvssV3AttackVectorEnumSlice flattens the contents of NoteVulnerabilityCvssV3AttackVectorEnum from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3AttackVectorEnumSlice(c *Client, i interface{}, res *Note) []NoteVulnerabilityCvssV3AttackVectorEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityCvssV3AttackVectorEnum{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityCvssV3AttackVectorEnum{}
	}

	items := make([]NoteVulnerabilityCvssV3AttackVectorEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityCvssV3AttackVectorEnum(item.(interface{})))
	}

	return items
}

// flattenNoteVulnerabilityCvssV3AttackVectorEnum asserts that an interface is a string, and returns a
// pointer to a *NoteVulnerabilityCvssV3AttackVectorEnum with the same value as that string.
func flattenNoteVulnerabilityCvssV3AttackVectorEnum(i interface{}) *NoteVulnerabilityCvssV3AttackVectorEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return NoteVulnerabilityCvssV3AttackVectorEnumRef(s)
}

// flattenNoteVulnerabilityCvssV3AttackComplexityEnumMap flattens the contents of NoteVulnerabilityCvssV3AttackComplexityEnum from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3AttackComplexityEnumMap(c *Client, i interface{}, res *Note) map[string]NoteVulnerabilityCvssV3AttackComplexityEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerabilityCvssV3AttackComplexityEnum{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerabilityCvssV3AttackComplexityEnum{}
	}

	items := make(map[string]NoteVulnerabilityCvssV3AttackComplexityEnum)
	for k, item := range a {
		items[k] = *flattenNoteVulnerabilityCvssV3AttackComplexityEnum(item.(interface{}))
	}

	return items
}

// flattenNoteVulnerabilityCvssV3AttackComplexityEnumSlice flattens the contents of NoteVulnerabilityCvssV3AttackComplexityEnum from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3AttackComplexityEnumSlice(c *Client, i interface{}, res *Note) []NoteVulnerabilityCvssV3AttackComplexityEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityCvssV3AttackComplexityEnum{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityCvssV3AttackComplexityEnum{}
	}

	items := make([]NoteVulnerabilityCvssV3AttackComplexityEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityCvssV3AttackComplexityEnum(item.(interface{})))
	}

	return items
}

// flattenNoteVulnerabilityCvssV3AttackComplexityEnum asserts that an interface is a string, and returns a
// pointer to a *NoteVulnerabilityCvssV3AttackComplexityEnum with the same value as that string.
func flattenNoteVulnerabilityCvssV3AttackComplexityEnum(i interface{}) *NoteVulnerabilityCvssV3AttackComplexityEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return NoteVulnerabilityCvssV3AttackComplexityEnumRef(s)
}

// flattenNoteVulnerabilityCvssV3PrivilegesRequiredEnumMap flattens the contents of NoteVulnerabilityCvssV3PrivilegesRequiredEnum from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3PrivilegesRequiredEnumMap(c *Client, i interface{}, res *Note) map[string]NoteVulnerabilityCvssV3PrivilegesRequiredEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerabilityCvssV3PrivilegesRequiredEnum{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerabilityCvssV3PrivilegesRequiredEnum{}
	}

	items := make(map[string]NoteVulnerabilityCvssV3PrivilegesRequiredEnum)
	for k, item := range a {
		items[k] = *flattenNoteVulnerabilityCvssV3PrivilegesRequiredEnum(item.(interface{}))
	}

	return items
}

// flattenNoteVulnerabilityCvssV3PrivilegesRequiredEnumSlice flattens the contents of NoteVulnerabilityCvssV3PrivilegesRequiredEnum from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3PrivilegesRequiredEnumSlice(c *Client, i interface{}, res *Note) []NoteVulnerabilityCvssV3PrivilegesRequiredEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityCvssV3PrivilegesRequiredEnum{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityCvssV3PrivilegesRequiredEnum{}
	}

	items := make([]NoteVulnerabilityCvssV3PrivilegesRequiredEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityCvssV3PrivilegesRequiredEnum(item.(interface{})))
	}

	return items
}

// flattenNoteVulnerabilityCvssV3PrivilegesRequiredEnum asserts that an interface is a string, and returns a
// pointer to a *NoteVulnerabilityCvssV3PrivilegesRequiredEnum with the same value as that string.
func flattenNoteVulnerabilityCvssV3PrivilegesRequiredEnum(i interface{}) *NoteVulnerabilityCvssV3PrivilegesRequiredEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return NoteVulnerabilityCvssV3PrivilegesRequiredEnumRef(s)
}

// flattenNoteVulnerabilityCvssV3UserInteractionEnumMap flattens the contents of NoteVulnerabilityCvssV3UserInteractionEnum from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3UserInteractionEnumMap(c *Client, i interface{}, res *Note) map[string]NoteVulnerabilityCvssV3UserInteractionEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerabilityCvssV3UserInteractionEnum{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerabilityCvssV3UserInteractionEnum{}
	}

	items := make(map[string]NoteVulnerabilityCvssV3UserInteractionEnum)
	for k, item := range a {
		items[k] = *flattenNoteVulnerabilityCvssV3UserInteractionEnum(item.(interface{}))
	}

	return items
}

// flattenNoteVulnerabilityCvssV3UserInteractionEnumSlice flattens the contents of NoteVulnerabilityCvssV3UserInteractionEnum from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3UserInteractionEnumSlice(c *Client, i interface{}, res *Note) []NoteVulnerabilityCvssV3UserInteractionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityCvssV3UserInteractionEnum{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityCvssV3UserInteractionEnum{}
	}

	items := make([]NoteVulnerabilityCvssV3UserInteractionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityCvssV3UserInteractionEnum(item.(interface{})))
	}

	return items
}

// flattenNoteVulnerabilityCvssV3UserInteractionEnum asserts that an interface is a string, and returns a
// pointer to a *NoteVulnerabilityCvssV3UserInteractionEnum with the same value as that string.
func flattenNoteVulnerabilityCvssV3UserInteractionEnum(i interface{}) *NoteVulnerabilityCvssV3UserInteractionEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return NoteVulnerabilityCvssV3UserInteractionEnumRef(s)
}

// flattenNoteVulnerabilityCvssV3ScopeEnumMap flattens the contents of NoteVulnerabilityCvssV3ScopeEnum from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3ScopeEnumMap(c *Client, i interface{}, res *Note) map[string]NoteVulnerabilityCvssV3ScopeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerabilityCvssV3ScopeEnum{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerabilityCvssV3ScopeEnum{}
	}

	items := make(map[string]NoteVulnerabilityCvssV3ScopeEnum)
	for k, item := range a {
		items[k] = *flattenNoteVulnerabilityCvssV3ScopeEnum(item.(interface{}))
	}

	return items
}

// flattenNoteVulnerabilityCvssV3ScopeEnumSlice flattens the contents of NoteVulnerabilityCvssV3ScopeEnum from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3ScopeEnumSlice(c *Client, i interface{}, res *Note) []NoteVulnerabilityCvssV3ScopeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityCvssV3ScopeEnum{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityCvssV3ScopeEnum{}
	}

	items := make([]NoteVulnerabilityCvssV3ScopeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityCvssV3ScopeEnum(item.(interface{})))
	}

	return items
}

// flattenNoteVulnerabilityCvssV3ScopeEnum asserts that an interface is a string, and returns a
// pointer to a *NoteVulnerabilityCvssV3ScopeEnum with the same value as that string.
func flattenNoteVulnerabilityCvssV3ScopeEnum(i interface{}) *NoteVulnerabilityCvssV3ScopeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return NoteVulnerabilityCvssV3ScopeEnumRef(s)
}

// flattenNoteVulnerabilityCvssV3ConfidentialityImpactEnumMap flattens the contents of NoteVulnerabilityCvssV3ConfidentialityImpactEnum from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3ConfidentialityImpactEnumMap(c *Client, i interface{}, res *Note) map[string]NoteVulnerabilityCvssV3ConfidentialityImpactEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerabilityCvssV3ConfidentialityImpactEnum{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerabilityCvssV3ConfidentialityImpactEnum{}
	}

	items := make(map[string]NoteVulnerabilityCvssV3ConfidentialityImpactEnum)
	for k, item := range a {
		items[k] = *flattenNoteVulnerabilityCvssV3ConfidentialityImpactEnum(item.(interface{}))
	}

	return items
}

// flattenNoteVulnerabilityCvssV3ConfidentialityImpactEnumSlice flattens the contents of NoteVulnerabilityCvssV3ConfidentialityImpactEnum from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3ConfidentialityImpactEnumSlice(c *Client, i interface{}, res *Note) []NoteVulnerabilityCvssV3ConfidentialityImpactEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityCvssV3ConfidentialityImpactEnum{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityCvssV3ConfidentialityImpactEnum{}
	}

	items := make([]NoteVulnerabilityCvssV3ConfidentialityImpactEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityCvssV3ConfidentialityImpactEnum(item.(interface{})))
	}

	return items
}

// flattenNoteVulnerabilityCvssV3ConfidentialityImpactEnum asserts that an interface is a string, and returns a
// pointer to a *NoteVulnerabilityCvssV3ConfidentialityImpactEnum with the same value as that string.
func flattenNoteVulnerabilityCvssV3ConfidentialityImpactEnum(i interface{}) *NoteVulnerabilityCvssV3ConfidentialityImpactEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return NoteVulnerabilityCvssV3ConfidentialityImpactEnumRef(s)
}

// flattenNoteVulnerabilityCvssV3IntegrityImpactEnumMap flattens the contents of NoteVulnerabilityCvssV3IntegrityImpactEnum from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3IntegrityImpactEnumMap(c *Client, i interface{}, res *Note) map[string]NoteVulnerabilityCvssV3IntegrityImpactEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerabilityCvssV3IntegrityImpactEnum{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerabilityCvssV3IntegrityImpactEnum{}
	}

	items := make(map[string]NoteVulnerabilityCvssV3IntegrityImpactEnum)
	for k, item := range a {
		items[k] = *flattenNoteVulnerabilityCvssV3IntegrityImpactEnum(item.(interface{}))
	}

	return items
}

// flattenNoteVulnerabilityCvssV3IntegrityImpactEnumSlice flattens the contents of NoteVulnerabilityCvssV3IntegrityImpactEnum from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3IntegrityImpactEnumSlice(c *Client, i interface{}, res *Note) []NoteVulnerabilityCvssV3IntegrityImpactEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityCvssV3IntegrityImpactEnum{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityCvssV3IntegrityImpactEnum{}
	}

	items := make([]NoteVulnerabilityCvssV3IntegrityImpactEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityCvssV3IntegrityImpactEnum(item.(interface{})))
	}

	return items
}

// flattenNoteVulnerabilityCvssV3IntegrityImpactEnum asserts that an interface is a string, and returns a
// pointer to a *NoteVulnerabilityCvssV3IntegrityImpactEnum with the same value as that string.
func flattenNoteVulnerabilityCvssV3IntegrityImpactEnum(i interface{}) *NoteVulnerabilityCvssV3IntegrityImpactEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return NoteVulnerabilityCvssV3IntegrityImpactEnumRef(s)
}

// flattenNoteVulnerabilityCvssV3AvailabilityImpactEnumMap flattens the contents of NoteVulnerabilityCvssV3AvailabilityImpactEnum from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3AvailabilityImpactEnumMap(c *Client, i interface{}, res *Note) map[string]NoteVulnerabilityCvssV3AvailabilityImpactEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteVulnerabilityCvssV3AvailabilityImpactEnum{}
	}

	if len(a) == 0 {
		return map[string]NoteVulnerabilityCvssV3AvailabilityImpactEnum{}
	}

	items := make(map[string]NoteVulnerabilityCvssV3AvailabilityImpactEnum)
	for k, item := range a {
		items[k] = *flattenNoteVulnerabilityCvssV3AvailabilityImpactEnum(item.(interface{}))
	}

	return items
}

// flattenNoteVulnerabilityCvssV3AvailabilityImpactEnumSlice flattens the contents of NoteVulnerabilityCvssV3AvailabilityImpactEnum from a JSON
// response object.
func flattenNoteVulnerabilityCvssV3AvailabilityImpactEnumSlice(c *Client, i interface{}, res *Note) []NoteVulnerabilityCvssV3AvailabilityImpactEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteVulnerabilityCvssV3AvailabilityImpactEnum{}
	}

	if len(a) == 0 {
		return []NoteVulnerabilityCvssV3AvailabilityImpactEnum{}
	}

	items := make([]NoteVulnerabilityCvssV3AvailabilityImpactEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteVulnerabilityCvssV3AvailabilityImpactEnum(item.(interface{})))
	}

	return items
}

// flattenNoteVulnerabilityCvssV3AvailabilityImpactEnum asserts that an interface is a string, and returns a
// pointer to a *NoteVulnerabilityCvssV3AvailabilityImpactEnum with the same value as that string.
func flattenNoteVulnerabilityCvssV3AvailabilityImpactEnum(i interface{}) *NoteVulnerabilityCvssV3AvailabilityImpactEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return NoteVulnerabilityCvssV3AvailabilityImpactEnumRef(s)
}

// flattenNoteBuildSignatureKeyTypeEnumMap flattens the contents of NoteBuildSignatureKeyTypeEnum from a JSON
// response object.
func flattenNoteBuildSignatureKeyTypeEnumMap(c *Client, i interface{}, res *Note) map[string]NoteBuildSignatureKeyTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteBuildSignatureKeyTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]NoteBuildSignatureKeyTypeEnum{}
	}

	items := make(map[string]NoteBuildSignatureKeyTypeEnum)
	for k, item := range a {
		items[k] = *flattenNoteBuildSignatureKeyTypeEnum(item.(interface{}))
	}

	return items
}

// flattenNoteBuildSignatureKeyTypeEnumSlice flattens the contents of NoteBuildSignatureKeyTypeEnum from a JSON
// response object.
func flattenNoteBuildSignatureKeyTypeEnumSlice(c *Client, i interface{}, res *Note) []NoteBuildSignatureKeyTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteBuildSignatureKeyTypeEnum{}
	}

	if len(a) == 0 {
		return []NoteBuildSignatureKeyTypeEnum{}
	}

	items := make([]NoteBuildSignatureKeyTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteBuildSignatureKeyTypeEnum(item.(interface{})))
	}

	return items
}

// flattenNoteBuildSignatureKeyTypeEnum asserts that an interface is a string, and returns a
// pointer to a *NoteBuildSignatureKeyTypeEnum with the same value as that string.
func flattenNoteBuildSignatureKeyTypeEnum(i interface{}) *NoteBuildSignatureKeyTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return NoteBuildSignatureKeyTypeEnumRef(s)
}

// flattenNotePackageDistributionArchitectureEnumMap flattens the contents of NotePackageDistributionArchitectureEnum from a JSON
// response object.
func flattenNotePackageDistributionArchitectureEnumMap(c *Client, i interface{}, res *Note) map[string]NotePackageDistributionArchitectureEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NotePackageDistributionArchitectureEnum{}
	}

	if len(a) == 0 {
		return map[string]NotePackageDistributionArchitectureEnum{}
	}

	items := make(map[string]NotePackageDistributionArchitectureEnum)
	for k, item := range a {
		items[k] = *flattenNotePackageDistributionArchitectureEnum(item.(interface{}))
	}

	return items
}

// flattenNotePackageDistributionArchitectureEnumSlice flattens the contents of NotePackageDistributionArchitectureEnum from a JSON
// response object.
func flattenNotePackageDistributionArchitectureEnumSlice(c *Client, i interface{}, res *Note) []NotePackageDistributionArchitectureEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NotePackageDistributionArchitectureEnum{}
	}

	if len(a) == 0 {
		return []NotePackageDistributionArchitectureEnum{}
	}

	items := make([]NotePackageDistributionArchitectureEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNotePackageDistributionArchitectureEnum(item.(interface{})))
	}

	return items
}

// flattenNotePackageDistributionArchitectureEnum asserts that an interface is a string, and returns a
// pointer to a *NotePackageDistributionArchitectureEnum with the same value as that string.
func flattenNotePackageDistributionArchitectureEnum(i interface{}) *NotePackageDistributionArchitectureEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return NotePackageDistributionArchitectureEnumRef(s)
}

// flattenNotePackageDistributionLatestVersionKindEnumMap flattens the contents of NotePackageDistributionLatestVersionKindEnum from a JSON
// response object.
func flattenNotePackageDistributionLatestVersionKindEnumMap(c *Client, i interface{}, res *Note) map[string]NotePackageDistributionLatestVersionKindEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NotePackageDistributionLatestVersionKindEnum{}
	}

	if len(a) == 0 {
		return map[string]NotePackageDistributionLatestVersionKindEnum{}
	}

	items := make(map[string]NotePackageDistributionLatestVersionKindEnum)
	for k, item := range a {
		items[k] = *flattenNotePackageDistributionLatestVersionKindEnum(item.(interface{}))
	}

	return items
}

// flattenNotePackageDistributionLatestVersionKindEnumSlice flattens the contents of NotePackageDistributionLatestVersionKindEnum from a JSON
// response object.
func flattenNotePackageDistributionLatestVersionKindEnumSlice(c *Client, i interface{}, res *Note) []NotePackageDistributionLatestVersionKindEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NotePackageDistributionLatestVersionKindEnum{}
	}

	if len(a) == 0 {
		return []NotePackageDistributionLatestVersionKindEnum{}
	}

	items := make([]NotePackageDistributionLatestVersionKindEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNotePackageDistributionLatestVersionKindEnum(item.(interface{})))
	}

	return items
}

// flattenNotePackageDistributionLatestVersionKindEnum asserts that an interface is a string, and returns a
// pointer to a *NotePackageDistributionLatestVersionKindEnum with the same value as that string.
func flattenNotePackageDistributionLatestVersionKindEnum(i interface{}) *NotePackageDistributionLatestVersionKindEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return NotePackageDistributionLatestVersionKindEnumRef(s)
}

// flattenNoteDiscoveryAnalysisKindEnumMap flattens the contents of NoteDiscoveryAnalysisKindEnum from a JSON
// response object.
func flattenNoteDiscoveryAnalysisKindEnumMap(c *Client, i interface{}, res *Note) map[string]NoteDiscoveryAnalysisKindEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]NoteDiscoveryAnalysisKindEnum{}
	}

	if len(a) == 0 {
		return map[string]NoteDiscoveryAnalysisKindEnum{}
	}

	items := make(map[string]NoteDiscoveryAnalysisKindEnum)
	for k, item := range a {
		items[k] = *flattenNoteDiscoveryAnalysisKindEnum(item.(interface{}))
	}

	return items
}

// flattenNoteDiscoveryAnalysisKindEnumSlice flattens the contents of NoteDiscoveryAnalysisKindEnum from a JSON
// response object.
func flattenNoteDiscoveryAnalysisKindEnumSlice(c *Client, i interface{}, res *Note) []NoteDiscoveryAnalysisKindEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []NoteDiscoveryAnalysisKindEnum{}
	}

	if len(a) == 0 {
		return []NoteDiscoveryAnalysisKindEnum{}
	}

	items := make([]NoteDiscoveryAnalysisKindEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenNoteDiscoveryAnalysisKindEnum(item.(interface{})))
	}

	return items
}

// flattenNoteDiscoveryAnalysisKindEnum asserts that an interface is a string, and returns a
// pointer to a *NoteDiscoveryAnalysisKindEnum with the same value as that string.
func flattenNoteDiscoveryAnalysisKindEnum(i interface{}) *NoteDiscoveryAnalysisKindEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return NoteDiscoveryAnalysisKindEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Note) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalNote(b, c, r)
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

type noteDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         noteApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToNoteDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]noteDiff, error) {
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
	var diffs []noteDiff
	// For each operation name, create a noteDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := noteDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToNoteApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToNoteApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (noteApiOperation, error) {
	switch opName {

	case "updateNoteUpdateNoteOperation":
		return &updateNoteUpdateNoteOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractNoteFields(r *Note) error {
	vVulnerability := r.Vulnerability
	if vVulnerability == nil {
		// note: explicitly not the empty object.
		vVulnerability = &NoteVulnerability{}
	}
	if err := extractNoteVulnerabilityFields(r, vVulnerability); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vVulnerability) {
		r.Vulnerability = vVulnerability
	}
	vBuild := r.Build
	if vBuild == nil {
		// note: explicitly not the empty object.
		vBuild = &NoteBuild{}
	}
	if err := extractNoteBuildFields(r, vBuild); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vBuild) {
		r.Build = vBuild
	}
	vImage := r.Image
	if vImage == nil {
		// note: explicitly not the empty object.
		vImage = &NoteImage{}
	}
	if err := extractNoteImageFields(r, vImage); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vImage) {
		r.Image = vImage
	}
	vPackage := r.Package
	if vPackage == nil {
		// note: explicitly not the empty object.
		vPackage = &NotePackage{}
	}
	if err := extractNotePackageFields(r, vPackage); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPackage) {
		r.Package = vPackage
	}
	vDiscovery := r.Discovery
	if vDiscovery == nil {
		// note: explicitly not the empty object.
		vDiscovery = &NoteDiscovery{}
	}
	if err := extractNoteDiscoveryFields(r, vDiscovery); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDiscovery) {
		r.Discovery = vDiscovery
	}
	vDeployment := r.Deployment
	if vDeployment == nil {
		// note: explicitly not the empty object.
		vDeployment = &NoteDeployment{}
	}
	if err := extractNoteDeploymentFields(r, vDeployment); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDeployment) {
		r.Deployment = vDeployment
	}
	vAttestation := r.Attestation
	if vAttestation == nil {
		// note: explicitly not the empty object.
		vAttestation = &NoteAttestation{}
	}
	if err := extractNoteAttestationFields(r, vAttestation); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAttestation) {
		r.Attestation = vAttestation
	}
	return nil
}
func extractNoteRelatedUrlFields(r *Note, o *NoteRelatedUrl) error {
	return nil
}
func extractNoteVulnerabilityFields(r *Note, o *NoteVulnerability) error {
	vCvssV3 := o.CvssV3
	if vCvssV3 == nil {
		// note: explicitly not the empty object.
		vCvssV3 = &NoteVulnerabilityCvssV3{}
	}
	if err := extractNoteVulnerabilityCvssV3Fields(r, vCvssV3); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCvssV3) {
		o.CvssV3 = vCvssV3
	}
	return nil
}
func extractNoteVulnerabilityDetailsFields(r *Note, o *NoteVulnerabilityDetails) error {
	vAffectedVersionStart := o.AffectedVersionStart
	if vAffectedVersionStart == nil {
		// note: explicitly not the empty object.
		vAffectedVersionStart = &NoteVulnerabilityDetailsAffectedVersionStart{}
	}
	if err := extractNoteVulnerabilityDetailsAffectedVersionStartFields(r, vAffectedVersionStart); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAffectedVersionStart) {
		o.AffectedVersionStart = vAffectedVersionStart
	}
	vAffectedVersionEnd := o.AffectedVersionEnd
	if vAffectedVersionEnd == nil {
		// note: explicitly not the empty object.
		vAffectedVersionEnd = &NoteVulnerabilityDetailsAffectedVersionEnd{}
	}
	if err := extractNoteVulnerabilityDetailsAffectedVersionEndFields(r, vAffectedVersionEnd); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAffectedVersionEnd) {
		o.AffectedVersionEnd = vAffectedVersionEnd
	}
	vFixedVersion := o.FixedVersion
	if vFixedVersion == nil {
		// note: explicitly not the empty object.
		vFixedVersion = &NoteVulnerabilityDetailsFixedVersion{}
	}
	if err := extractNoteVulnerabilityDetailsFixedVersionFields(r, vFixedVersion); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vFixedVersion) {
		o.FixedVersion = vFixedVersion
	}
	return nil
}
func extractNoteVulnerabilityDetailsAffectedVersionStartFields(r *Note, o *NoteVulnerabilityDetailsAffectedVersionStart) error {
	return nil
}
func extractNoteVulnerabilityDetailsAffectedVersionEndFields(r *Note, o *NoteVulnerabilityDetailsAffectedVersionEnd) error {
	return nil
}
func extractNoteVulnerabilityDetailsFixedVersionFields(r *Note, o *NoteVulnerabilityDetailsFixedVersion) error {
	return nil
}
func extractNoteVulnerabilityCvssV3Fields(r *Note, o *NoteVulnerabilityCvssV3) error {
	return nil
}
func extractNoteVulnerabilityWindowsDetailsFields(r *Note, o *NoteVulnerabilityWindowsDetails) error {
	return nil
}
func extractNoteVulnerabilityWindowsDetailsFixingKbsFields(r *Note, o *NoteVulnerabilityWindowsDetailsFixingKbs) error {
	return nil
}
func extractNoteBuildFields(r *Note, o *NoteBuild) error {
	vSignature := o.Signature
	if vSignature == nil {
		// note: explicitly not the empty object.
		vSignature = &NoteBuildSignature{}
	}
	if err := extractNoteBuildSignatureFields(r, vSignature); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSignature) {
		o.Signature = vSignature
	}
	return nil
}
func extractNoteBuildSignatureFields(r *Note, o *NoteBuildSignature) error {
	return nil
}
func extractNoteImageFields(r *Note, o *NoteImage) error {
	vFingerprint := o.Fingerprint
	if vFingerprint == nil {
		// note: explicitly not the empty object.
		vFingerprint = &NoteImageFingerprint{}
	}
	if err := extractNoteImageFingerprintFields(r, vFingerprint); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vFingerprint) {
		o.Fingerprint = vFingerprint
	}
	return nil
}
func extractNoteImageFingerprintFields(r *Note, o *NoteImageFingerprint) error {
	return nil
}
func extractNotePackageFields(r *Note, o *NotePackage) error {
	return nil
}
func extractNotePackageDistributionFields(r *Note, o *NotePackageDistribution) error {
	vLatestVersion := o.LatestVersion
	if vLatestVersion == nil {
		// note: explicitly not the empty object.
		vLatestVersion = &NotePackageDistributionLatestVersion{}
	}
	if err := extractNotePackageDistributionLatestVersionFields(r, vLatestVersion); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLatestVersion) {
		o.LatestVersion = vLatestVersion
	}
	return nil
}
func extractNotePackageDistributionLatestVersionFields(r *Note, o *NotePackageDistributionLatestVersion) error {
	return nil
}
func extractNoteDiscoveryFields(r *Note, o *NoteDiscovery) error {
	return nil
}
func extractNoteDeploymentFields(r *Note, o *NoteDeployment) error {
	return nil
}
func extractNoteAttestationFields(r *Note, o *NoteAttestation) error {
	vHint := o.Hint
	if vHint == nil {
		// note: explicitly not the empty object.
		vHint = &NoteAttestationHint{}
	}
	if err := extractNoteAttestationHintFields(r, vHint); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vHint) {
		o.Hint = vHint
	}
	return nil
}
func extractNoteAttestationHintFields(r *Note, o *NoteAttestationHint) error {
	return nil
}

func postReadExtractNoteFields(r *Note) error {
	vVulnerability := r.Vulnerability
	if vVulnerability == nil {
		// note: explicitly not the empty object.
		vVulnerability = &NoteVulnerability{}
	}
	if err := postReadExtractNoteVulnerabilityFields(r, vVulnerability); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vVulnerability) {
		r.Vulnerability = vVulnerability
	}
	vBuild := r.Build
	if vBuild == nil {
		// note: explicitly not the empty object.
		vBuild = &NoteBuild{}
	}
	if err := postReadExtractNoteBuildFields(r, vBuild); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vBuild) {
		r.Build = vBuild
	}
	vImage := r.Image
	if vImage == nil {
		// note: explicitly not the empty object.
		vImage = &NoteImage{}
	}
	if err := postReadExtractNoteImageFields(r, vImage); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vImage) {
		r.Image = vImage
	}
	vPackage := r.Package
	if vPackage == nil {
		// note: explicitly not the empty object.
		vPackage = &NotePackage{}
	}
	if err := postReadExtractNotePackageFields(r, vPackage); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPackage) {
		r.Package = vPackage
	}
	vDiscovery := r.Discovery
	if vDiscovery == nil {
		// note: explicitly not the empty object.
		vDiscovery = &NoteDiscovery{}
	}
	if err := postReadExtractNoteDiscoveryFields(r, vDiscovery); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDiscovery) {
		r.Discovery = vDiscovery
	}
	vDeployment := r.Deployment
	if vDeployment == nil {
		// note: explicitly not the empty object.
		vDeployment = &NoteDeployment{}
	}
	if err := postReadExtractNoteDeploymentFields(r, vDeployment); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDeployment) {
		r.Deployment = vDeployment
	}
	vAttestation := r.Attestation
	if vAttestation == nil {
		// note: explicitly not the empty object.
		vAttestation = &NoteAttestation{}
	}
	if err := postReadExtractNoteAttestationFields(r, vAttestation); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAttestation) {
		r.Attestation = vAttestation
	}
	return nil
}
func postReadExtractNoteRelatedUrlFields(r *Note, o *NoteRelatedUrl) error {
	return nil
}
func postReadExtractNoteVulnerabilityFields(r *Note, o *NoteVulnerability) error {
	vCvssV3 := o.CvssV3
	if vCvssV3 == nil {
		// note: explicitly not the empty object.
		vCvssV3 = &NoteVulnerabilityCvssV3{}
	}
	if err := extractNoteVulnerabilityCvssV3Fields(r, vCvssV3); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCvssV3) {
		o.CvssV3 = vCvssV3
	}
	return nil
}
func postReadExtractNoteVulnerabilityDetailsFields(r *Note, o *NoteVulnerabilityDetails) error {
	vAffectedVersionStart := o.AffectedVersionStart
	if vAffectedVersionStart == nil {
		// note: explicitly not the empty object.
		vAffectedVersionStart = &NoteVulnerabilityDetailsAffectedVersionStart{}
	}
	if err := extractNoteVulnerabilityDetailsAffectedVersionStartFields(r, vAffectedVersionStart); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAffectedVersionStart) {
		o.AffectedVersionStart = vAffectedVersionStart
	}
	vAffectedVersionEnd := o.AffectedVersionEnd
	if vAffectedVersionEnd == nil {
		// note: explicitly not the empty object.
		vAffectedVersionEnd = &NoteVulnerabilityDetailsAffectedVersionEnd{}
	}
	if err := extractNoteVulnerabilityDetailsAffectedVersionEndFields(r, vAffectedVersionEnd); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAffectedVersionEnd) {
		o.AffectedVersionEnd = vAffectedVersionEnd
	}
	vFixedVersion := o.FixedVersion
	if vFixedVersion == nil {
		// note: explicitly not the empty object.
		vFixedVersion = &NoteVulnerabilityDetailsFixedVersion{}
	}
	if err := extractNoteVulnerabilityDetailsFixedVersionFields(r, vFixedVersion); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vFixedVersion) {
		o.FixedVersion = vFixedVersion
	}
	return nil
}
func postReadExtractNoteVulnerabilityDetailsAffectedVersionStartFields(r *Note, o *NoteVulnerabilityDetailsAffectedVersionStart) error {
	return nil
}
func postReadExtractNoteVulnerabilityDetailsAffectedVersionEndFields(r *Note, o *NoteVulnerabilityDetailsAffectedVersionEnd) error {
	return nil
}
func postReadExtractNoteVulnerabilityDetailsFixedVersionFields(r *Note, o *NoteVulnerabilityDetailsFixedVersion) error {
	return nil
}
func postReadExtractNoteVulnerabilityCvssV3Fields(r *Note, o *NoteVulnerabilityCvssV3) error {
	return nil
}
func postReadExtractNoteVulnerabilityWindowsDetailsFields(r *Note, o *NoteVulnerabilityWindowsDetails) error {
	return nil
}
func postReadExtractNoteVulnerabilityWindowsDetailsFixingKbsFields(r *Note, o *NoteVulnerabilityWindowsDetailsFixingKbs) error {
	return nil
}
func postReadExtractNoteBuildFields(r *Note, o *NoteBuild) error {
	vSignature := o.Signature
	if vSignature == nil {
		// note: explicitly not the empty object.
		vSignature = &NoteBuildSignature{}
	}
	if err := extractNoteBuildSignatureFields(r, vSignature); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSignature) {
		o.Signature = vSignature
	}
	return nil
}
func postReadExtractNoteBuildSignatureFields(r *Note, o *NoteBuildSignature) error {
	return nil
}
func postReadExtractNoteImageFields(r *Note, o *NoteImage) error {
	vFingerprint := o.Fingerprint
	if vFingerprint == nil {
		// note: explicitly not the empty object.
		vFingerprint = &NoteImageFingerprint{}
	}
	if err := extractNoteImageFingerprintFields(r, vFingerprint); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vFingerprint) {
		o.Fingerprint = vFingerprint
	}
	return nil
}
func postReadExtractNoteImageFingerprintFields(r *Note, o *NoteImageFingerprint) error {
	return nil
}
func postReadExtractNotePackageFields(r *Note, o *NotePackage) error {
	return nil
}
func postReadExtractNotePackageDistributionFields(r *Note, o *NotePackageDistribution) error {
	vLatestVersion := o.LatestVersion
	if vLatestVersion == nil {
		// note: explicitly not the empty object.
		vLatestVersion = &NotePackageDistributionLatestVersion{}
	}
	if err := extractNotePackageDistributionLatestVersionFields(r, vLatestVersion); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLatestVersion) {
		o.LatestVersion = vLatestVersion
	}
	return nil
}
func postReadExtractNotePackageDistributionLatestVersionFields(r *Note, o *NotePackageDistributionLatestVersion) error {
	return nil
}
func postReadExtractNoteDiscoveryFields(r *Note, o *NoteDiscovery) error {
	return nil
}
func postReadExtractNoteDeploymentFields(r *Note, o *NoteDeployment) error {
	return nil
}
func postReadExtractNoteAttestationFields(r *Note, o *NoteAttestation) error {
	vHint := o.Hint
	if vHint == nil {
		// note: explicitly not the empty object.
		vHint = &NoteAttestationHint{}
	}
	if err := extractNoteAttestationHintFields(r, vHint); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vHint) {
		o.Hint = vHint
	}
	return nil
}
func postReadExtractNoteAttestationHintFields(r *Note, o *NoteAttestationHint) error {
	return nil
}
