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

func (r *CapacityCommitment) validate() error {

	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.FailureStatus) {
		if err := r.FailureStatus.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CapacityCommitmentFailureStatus) validate() error {
	return nil
}
func (r *CapacityCommitmentFailureStatusDetails) validate() error {
	return nil
}
func (r *CapacityCommitment) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://bigqueryreservation.googleapis.com/v1beta1/", params)
}

func (r *CapacityCommitment) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/capacityCommitments/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *CapacityCommitment) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/capacityCommitments", nr.basePath(), userBasePath, params), nil

}

func (r *CapacityCommitment) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/capacityCommitments?capacityCommitmentId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *CapacityCommitment) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/capacityCommitments/{{name}}", nr.basePath(), userBasePath, params), nil
}

// capacityCommitmentApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type capacityCommitmentApiOperation interface {
	do(context.Context, *CapacityCommitment, *Client) error
}

// newUpdateCapacityCommitmentUpdateCapacityCommitmentRequest creates a request for an
// CapacityCommitment resource's UpdateCapacityCommitment update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateCapacityCommitmentUpdateCapacityCommitmentRequest(ctx context.Context, f *CapacityCommitment, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.Plan; !dcl.IsEmptyValueIndirect(v) {
		req["plan"] = v
	}
	if v := f.RenewalPlan; !dcl.IsEmptyValueIndirect(v) {
		req["renewalPlan"] = v
	}
	return req, nil
}

// marshalUpdateCapacityCommitmentUpdateCapacityCommitmentRequest converts the update into
// the final JSON request body.
func marshalUpdateCapacityCommitmentUpdateCapacityCommitmentRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateCapacityCommitmentUpdateCapacityCommitmentOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateCapacityCommitmentUpdateCapacityCommitmentOperation) do(ctx context.Context, r *CapacityCommitment, c *Client) error {
	_, err := c.GetCapacityCommitment(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateCapacityCommitment")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateCapacityCommitmentUpdateCapacityCommitmentRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateCapacityCommitmentUpdateCapacityCommitmentRequest(c, req)
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

func (c *Client) listCapacityCommitmentRaw(ctx context.Context, r *CapacityCommitment, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != CapacityCommitmentMaxPage {
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

type listCapacityCommitmentOperation struct {
	CapacityCommitments []map[string]interface{} `json:"capacityCommitments"`
	Token               string                   `json:"nextPageToken"`
}

func (c *Client) listCapacityCommitment(ctx context.Context, r *CapacityCommitment, pageToken string, pageSize int32) ([]*CapacityCommitment, string, error) {
	b, err := c.listCapacityCommitmentRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listCapacityCommitmentOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*CapacityCommitment
	for _, v := range m.CapacityCommitments {
		res, err := unmarshalMapCapacityCommitment(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllCapacityCommitment(ctx context.Context, f func(*CapacityCommitment) bool, resources []*CapacityCommitment) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteCapacityCommitment(ctx, res)
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

type deleteCapacityCommitmentOperation struct{}

func (op *deleteCapacityCommitmentOperation) do(ctx context.Context, r *CapacityCommitment, c *Client) error {
	r, err := c.GetCapacityCommitment(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "CapacityCommitment not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetCapacityCommitment checking for existence. error: %v", err)
		return err
	}

	err = r.waitForEndTime(ctx, c)
	if err != nil {
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
		return fmt.Errorf("failed to delete CapacityCommitment: %w", err)
	}

	// We saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// This is the reason we are adding retry to handle that case.
	retriesRemaining := 10
	dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		_, err := c.GetCapacityCommitment(ctx, r)
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
type createCapacityCommitmentOperation struct {
	response map[string]interface{}
}

func (op *createCapacityCommitmentOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createCapacityCommitmentOperation) do(ctx context.Context, r *CapacityCommitment, c *Client) error {
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
		// Allowing creation to continue with Name set could result in a CapacityCommitment with the wrong Name.
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
	name, ok := op.response["name"].(string)
	if !ok {
		return fmt.Errorf("expected name to be a string in %v, was %T", op.response, op.response["name"])
	}
	r.Name = &name

	if _, err := c.GetCapacityCommitment(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getCapacityCommitmentRaw(ctx context.Context, r *CapacityCommitment) ([]byte, error) {

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

func (c *Client) capacityCommitmentDiffsForRawDesired(ctx context.Context, rawDesired *CapacityCommitment, opts ...dcl.ApplyOption) (initial, desired *CapacityCommitment, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *CapacityCommitment
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*CapacityCommitment); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected CapacityCommitment, got %T", sh)
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
		desired, err := canonicalizeCapacityCommitmentDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}
	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetCapacityCommitment(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a CapacityCommitment resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve CapacityCommitment resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that CapacityCommitment resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeCapacityCommitmentDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for CapacityCommitment: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for CapacityCommitment: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractCapacityCommitmentFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeCapacityCommitmentInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for CapacityCommitment: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeCapacityCommitmentDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for CapacityCommitment: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffCapacityCommitment(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeCapacityCommitmentInitialState(rawInitial, rawDesired *CapacityCommitment) (*CapacityCommitment, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeCapacityCommitmentDesiredState(rawDesired, rawInitial *CapacityCommitment, opts ...dcl.ApplyOption) (*CapacityCommitment, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.FailureStatus = canonicalizeCapacityCommitmentFailureStatus(rawDesired.FailureStatus, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &CapacityCommitment{}
	if dcl.IsZeroValue(rawDesired.Name) || (dcl.IsEmptyValueIndirect(rawDesired.Name) && dcl.IsEmptyValueIndirect(rawInitial.Name)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.IsZeroValue(rawDesired.SlotCount) || (dcl.IsEmptyValueIndirect(rawDesired.SlotCount) && dcl.IsEmptyValueIndirect(rawInitial.SlotCount)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.SlotCount = rawInitial.SlotCount
	} else {
		canonicalDesired.SlotCount = rawDesired.SlotCount
	}
	if dcl.IsZeroValue(rawDesired.Plan) || (dcl.IsEmptyValueIndirect(rawDesired.Plan) && dcl.IsEmptyValueIndirect(rawInitial.Plan)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Plan = rawInitial.Plan
	} else {
		canonicalDesired.Plan = rawDesired.Plan
	}
	if dcl.IsZeroValue(rawDesired.RenewalPlan) || (dcl.IsEmptyValueIndirect(rawDesired.RenewalPlan) && dcl.IsEmptyValueIndirect(rawInitial.RenewalPlan)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.RenewalPlan = rawInitial.RenewalPlan
	} else {
		canonicalDesired.RenewalPlan = rawDesired.RenewalPlan
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

	return canonicalDesired, nil
}

func canonicalizeCapacityCommitmentNewState(c *Client, rawNew, rawDesired *CapacityCommitment) (*CapacityCommitment, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SlotCount) && dcl.IsEmptyValueIndirect(rawDesired.SlotCount) {
		rawNew.SlotCount = rawDesired.SlotCount
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Plan) && dcl.IsEmptyValueIndirect(rawDesired.Plan) {
		rawNew.Plan = rawDesired.Plan
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CommitmentStartTime) && dcl.IsEmptyValueIndirect(rawDesired.CommitmentStartTime) {
		rawNew.CommitmentStartTime = rawDesired.CommitmentStartTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CommitmentEndTime) && dcl.IsEmptyValueIndirect(rawDesired.CommitmentEndTime) {
		rawNew.CommitmentEndTime = rawDesired.CommitmentEndTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.FailureStatus) && dcl.IsEmptyValueIndirect(rawDesired.FailureStatus) {
		rawNew.FailureStatus = rawDesired.FailureStatus
	} else {
		rawNew.FailureStatus = canonicalizeNewCapacityCommitmentFailureStatus(c, rawDesired.FailureStatus, rawNew.FailureStatus)
	}

	if dcl.IsEmptyValueIndirect(rawNew.RenewalPlan) && dcl.IsEmptyValueIndirect(rawDesired.RenewalPlan) {
		rawNew.RenewalPlan = rawDesired.RenewalPlan
	} else {
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeCapacityCommitmentFailureStatus(des, initial *CapacityCommitmentFailureStatus, opts ...dcl.ApplyOption) *CapacityCommitmentFailureStatus {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &CapacityCommitmentFailureStatus{}

	if dcl.IsZeroValue(des.Code) || (dcl.IsEmptyValueIndirect(des.Code) && dcl.IsEmptyValueIndirect(initial.Code)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Code = initial.Code
	} else {
		cDes.Code = des.Code
	}
	if dcl.StringCanonicalize(des.Message, initial.Message) || dcl.IsZeroValue(des.Message) {
		cDes.Message = initial.Message
	} else {
		cDes.Message = des.Message
	}
	cDes.Details = canonicalizeCapacityCommitmentFailureStatusDetailsSlice(des.Details, initial.Details, opts...)

	return cDes
}

func canonicalizeCapacityCommitmentFailureStatusSlice(des, initial []CapacityCommitmentFailureStatus, opts ...dcl.ApplyOption) []CapacityCommitmentFailureStatus {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]CapacityCommitmentFailureStatus, 0, len(des))
		for _, d := range des {
			cd := canonicalizeCapacityCommitmentFailureStatus(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]CapacityCommitmentFailureStatus, 0, len(des))
	for i, d := range des {
		cd := canonicalizeCapacityCommitmentFailureStatus(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewCapacityCommitmentFailureStatus(c *Client, des, nw *CapacityCommitmentFailureStatus) *CapacityCommitmentFailureStatus {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for CapacityCommitmentFailureStatus while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Message, nw.Message) {
		nw.Message = des.Message
	}
	nw.Details = canonicalizeNewCapacityCommitmentFailureStatusDetailsSlice(c, des.Details, nw.Details)

	return nw
}

func canonicalizeNewCapacityCommitmentFailureStatusSet(c *Client, des, nw []CapacityCommitmentFailureStatus) []CapacityCommitmentFailureStatus {
	if des == nil {
		return nw
	}
	var reorderedNew []CapacityCommitmentFailureStatus
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCapacityCommitmentFailureStatusNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCapacityCommitmentFailureStatusSlice(c *Client, des, nw []CapacityCommitmentFailureStatus) []CapacityCommitmentFailureStatus {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CapacityCommitmentFailureStatus
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCapacityCommitmentFailureStatus(c, &d, &n))
	}

	return items
}

func canonicalizeCapacityCommitmentFailureStatusDetails(des, initial *CapacityCommitmentFailureStatusDetails, opts ...dcl.ApplyOption) *CapacityCommitmentFailureStatusDetails {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &CapacityCommitmentFailureStatusDetails{}

	if dcl.StringCanonicalize(des.TypeUrl, initial.TypeUrl) || dcl.IsZeroValue(des.TypeUrl) {
		cDes.TypeUrl = initial.TypeUrl
	} else {
		cDes.TypeUrl = des.TypeUrl
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		cDes.Value = initial.Value
	} else {
		cDes.Value = des.Value
	}

	return cDes
}

func canonicalizeCapacityCommitmentFailureStatusDetailsSlice(des, initial []CapacityCommitmentFailureStatusDetails, opts ...dcl.ApplyOption) []CapacityCommitmentFailureStatusDetails {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]CapacityCommitmentFailureStatusDetails, 0, len(des))
		for _, d := range des {
			cd := canonicalizeCapacityCommitmentFailureStatusDetails(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]CapacityCommitmentFailureStatusDetails, 0, len(des))
	for i, d := range des {
		cd := canonicalizeCapacityCommitmentFailureStatusDetails(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewCapacityCommitmentFailureStatusDetails(c *Client, des, nw *CapacityCommitmentFailureStatusDetails) *CapacityCommitmentFailureStatusDetails {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for CapacityCommitmentFailureStatusDetails while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.TypeUrl, nw.TypeUrl) {
		nw.TypeUrl = des.TypeUrl
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}

	return nw
}

func canonicalizeNewCapacityCommitmentFailureStatusDetailsSet(c *Client, des, nw []CapacityCommitmentFailureStatusDetails) []CapacityCommitmentFailureStatusDetails {
	if des == nil {
		return nw
	}
	var reorderedNew []CapacityCommitmentFailureStatusDetails
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCapacityCommitmentFailureStatusDetailsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCapacityCommitmentFailureStatusDetailsSlice(c *Client, des, nw []CapacityCommitmentFailureStatusDetails) []CapacityCommitmentFailureStatusDetails {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CapacityCommitmentFailureStatusDetails
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCapacityCommitmentFailureStatusDetails(c, &d, &n))
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
func diffCapacityCommitment(c *Client, desired, actual *CapacityCommitment, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.SlotCount, actual.SlotCount, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SlotCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Plan, actual.Plan, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateCapacityCommitmentUpdateCapacityCommitmentOperation")}, fn.AddNest("Plan")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.CommitmentStartTime, actual.CommitmentStartTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CommitmentStartTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CommitmentEndTime, actual.CommitmentEndTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CommitmentEndTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FailureStatus, actual.FailureStatus, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareCapacityCommitmentFailureStatusNewStyle, EmptyObject: EmptyCapacityCommitmentFailureStatus, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FailureStatus")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RenewalPlan, actual.RenewalPlan, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateCapacityCommitmentUpdateCapacityCommitmentOperation")}, fn.AddNest("RenewalPlan")); len(ds) != 0 || err != nil {
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

	return newDiffs, nil
}
func compareCapacityCommitmentFailureStatusNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CapacityCommitmentFailureStatus)
	if !ok {
		desiredNotPointer, ok := d.(CapacityCommitmentFailureStatus)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CapacityCommitmentFailureStatus or *CapacityCommitmentFailureStatus", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CapacityCommitmentFailureStatus)
	if !ok {
		actualNotPointer, ok := a.(CapacityCommitmentFailureStatus)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CapacityCommitmentFailureStatus", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Code, actual.Code, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Code")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Message, actual.Message, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Message")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Details, actual.Details, dcl.DiffInfo{ObjectFunction: compareCapacityCommitmentFailureStatusDetailsNewStyle, EmptyObject: EmptyCapacityCommitmentFailureStatusDetails, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Details")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCapacityCommitmentFailureStatusDetailsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CapacityCommitmentFailureStatusDetails)
	if !ok {
		desiredNotPointer, ok := d.(CapacityCommitmentFailureStatusDetails)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CapacityCommitmentFailureStatusDetails or *CapacityCommitmentFailureStatusDetails", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CapacityCommitmentFailureStatusDetails)
	if !ok {
		actualNotPointer, ok := a.(CapacityCommitmentFailureStatusDetails)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CapacityCommitmentFailureStatusDetails", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TypeUrl, actual.TypeUrl, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TypeUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
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
func (r *CapacityCommitment) urlNormalized() *CapacityCommitment {
	normalized := dcl.Copy(*r).(CapacityCommitment)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *CapacityCommitment) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateCapacityCommitment" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/capacityCommitments/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the CapacityCommitment resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *CapacityCommitment) marshal(c *Client) ([]byte, error) {
	m, err := expandCapacityCommitment(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling CapacityCommitment: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalCapacityCommitment decodes JSON responses into the CapacityCommitment resource schema.
func unmarshalCapacityCommitment(b []byte, c *Client, res *CapacityCommitment) (*CapacityCommitment, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapCapacityCommitment(m, c, res)
}

func unmarshalMapCapacityCommitment(m map[string]interface{}, c *Client, res *CapacityCommitment) (*CapacityCommitment, error) {

	flattened := flattenCapacityCommitment(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandCapacityCommitment expands CapacityCommitment into a JSON request object.
func expandCapacityCommitment(c *Client, f *CapacityCommitment) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("projects/%s/locations/%s/capacityCommitments/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Location), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.SlotCount; dcl.ValueShouldBeSent(v) {
		m["slotCount"] = v
	}
	if v := f.Plan; dcl.ValueShouldBeSent(v) {
		m["plan"] = v
	}
	if v := f.RenewalPlan; dcl.ValueShouldBeSent(v) {
		m["renewalPlan"] = v
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

	return m, nil
}

// flattenCapacityCommitment flattens CapacityCommitment from a JSON request object into the
// CapacityCommitment type.
func flattenCapacityCommitment(c *Client, i interface{}, res *CapacityCommitment) *CapacityCommitment {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &CapacityCommitment{}
	resultRes.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	resultRes.SlotCount = dcl.FlattenInteger(m["slotCount"])
	resultRes.Plan = flattenCapacityCommitmentPlanEnum(m["plan"])
	resultRes.State = flattenCapacityCommitmentStateEnum(m["state"])
	resultRes.CommitmentStartTime = dcl.FlattenString(m["commitmentStartTime"])
	resultRes.CommitmentEndTime = dcl.FlattenString(m["commitmentEndTime"])
	resultRes.FailureStatus = flattenCapacityCommitmentFailureStatus(c, m["failureStatus"], res)
	resultRes.RenewalPlan = flattenCapacityCommitmentRenewalPlanEnum(m["renewalPlan"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// expandCapacityCommitmentFailureStatusMap expands the contents of CapacityCommitmentFailureStatus into a JSON
// request object.
func expandCapacityCommitmentFailureStatusMap(c *Client, f map[string]CapacityCommitmentFailureStatus, res *CapacityCommitment) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCapacityCommitmentFailureStatus(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCapacityCommitmentFailureStatusSlice expands the contents of CapacityCommitmentFailureStatus into a JSON
// request object.
func expandCapacityCommitmentFailureStatusSlice(c *Client, f []CapacityCommitmentFailureStatus, res *CapacityCommitment) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCapacityCommitmentFailureStatus(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCapacityCommitmentFailureStatusMap flattens the contents of CapacityCommitmentFailureStatus from a JSON
// response object.
func flattenCapacityCommitmentFailureStatusMap(c *Client, i interface{}, res *CapacityCommitment) map[string]CapacityCommitmentFailureStatus {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CapacityCommitmentFailureStatus{}
	}

	if len(a) == 0 {
		return map[string]CapacityCommitmentFailureStatus{}
	}

	items := make(map[string]CapacityCommitmentFailureStatus)
	for k, item := range a {
		items[k] = *flattenCapacityCommitmentFailureStatus(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenCapacityCommitmentFailureStatusSlice flattens the contents of CapacityCommitmentFailureStatus from a JSON
// response object.
func flattenCapacityCommitmentFailureStatusSlice(c *Client, i interface{}, res *CapacityCommitment) []CapacityCommitmentFailureStatus {
	a, ok := i.([]interface{})
	if !ok {
		return []CapacityCommitmentFailureStatus{}
	}

	if len(a) == 0 {
		return []CapacityCommitmentFailureStatus{}
	}

	items := make([]CapacityCommitmentFailureStatus, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCapacityCommitmentFailureStatus(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandCapacityCommitmentFailureStatus expands an instance of CapacityCommitmentFailureStatus into a JSON
// request object.
func expandCapacityCommitmentFailureStatus(c *Client, f *CapacityCommitmentFailureStatus, res *CapacityCommitment) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Code; !dcl.IsEmptyValueIndirect(v) {
		m["code"] = v
	}
	if v := f.Message; !dcl.IsEmptyValueIndirect(v) {
		m["message"] = v
	}
	if v, err := expandCapacityCommitmentFailureStatusDetailsSlice(c, f.Details, res); err != nil {
		return nil, fmt.Errorf("error expanding Details into details: %w", err)
	} else if v != nil {
		m["details"] = v
	}

	return m, nil
}

// flattenCapacityCommitmentFailureStatus flattens an instance of CapacityCommitmentFailureStatus from a JSON
// response object.
func flattenCapacityCommitmentFailureStatus(c *Client, i interface{}, res *CapacityCommitment) *CapacityCommitmentFailureStatus {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CapacityCommitmentFailureStatus{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCapacityCommitmentFailureStatus
	}
	r.Code = dcl.FlattenInteger(m["code"])
	r.Message = dcl.FlattenString(m["message"])
	r.Details = flattenCapacityCommitmentFailureStatusDetailsSlice(c, m["details"], res)

	return r
}

// expandCapacityCommitmentFailureStatusDetailsMap expands the contents of CapacityCommitmentFailureStatusDetails into a JSON
// request object.
func expandCapacityCommitmentFailureStatusDetailsMap(c *Client, f map[string]CapacityCommitmentFailureStatusDetails, res *CapacityCommitment) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCapacityCommitmentFailureStatusDetails(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCapacityCommitmentFailureStatusDetailsSlice expands the contents of CapacityCommitmentFailureStatusDetails into a JSON
// request object.
func expandCapacityCommitmentFailureStatusDetailsSlice(c *Client, f []CapacityCommitmentFailureStatusDetails, res *CapacityCommitment) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCapacityCommitmentFailureStatusDetails(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCapacityCommitmentFailureStatusDetailsMap flattens the contents of CapacityCommitmentFailureStatusDetails from a JSON
// response object.
func flattenCapacityCommitmentFailureStatusDetailsMap(c *Client, i interface{}, res *CapacityCommitment) map[string]CapacityCommitmentFailureStatusDetails {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CapacityCommitmentFailureStatusDetails{}
	}

	if len(a) == 0 {
		return map[string]CapacityCommitmentFailureStatusDetails{}
	}

	items := make(map[string]CapacityCommitmentFailureStatusDetails)
	for k, item := range a {
		items[k] = *flattenCapacityCommitmentFailureStatusDetails(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenCapacityCommitmentFailureStatusDetailsSlice flattens the contents of CapacityCommitmentFailureStatusDetails from a JSON
// response object.
func flattenCapacityCommitmentFailureStatusDetailsSlice(c *Client, i interface{}, res *CapacityCommitment) []CapacityCommitmentFailureStatusDetails {
	a, ok := i.([]interface{})
	if !ok {
		return []CapacityCommitmentFailureStatusDetails{}
	}

	if len(a) == 0 {
		return []CapacityCommitmentFailureStatusDetails{}
	}

	items := make([]CapacityCommitmentFailureStatusDetails, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCapacityCommitmentFailureStatusDetails(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandCapacityCommitmentFailureStatusDetails expands an instance of CapacityCommitmentFailureStatusDetails into a JSON
// request object.
func expandCapacityCommitmentFailureStatusDetails(c *Client, f *CapacityCommitmentFailureStatusDetails, res *CapacityCommitment) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.TypeUrl; !dcl.IsEmptyValueIndirect(v) {
		m["typeUrl"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenCapacityCommitmentFailureStatusDetails flattens an instance of CapacityCommitmentFailureStatusDetails from a JSON
// response object.
func flattenCapacityCommitmentFailureStatusDetails(c *Client, i interface{}, res *CapacityCommitment) *CapacityCommitmentFailureStatusDetails {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CapacityCommitmentFailureStatusDetails{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCapacityCommitmentFailureStatusDetails
	}
	r.TypeUrl = dcl.FlattenString(m["typeUrl"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// flattenCapacityCommitmentPlanEnumMap flattens the contents of CapacityCommitmentPlanEnum from a JSON
// response object.
func flattenCapacityCommitmentPlanEnumMap(c *Client, i interface{}, res *CapacityCommitment) map[string]CapacityCommitmentPlanEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CapacityCommitmentPlanEnum{}
	}

	if len(a) == 0 {
		return map[string]CapacityCommitmentPlanEnum{}
	}

	items := make(map[string]CapacityCommitmentPlanEnum)
	for k, item := range a {
		items[k] = *flattenCapacityCommitmentPlanEnum(item.(interface{}))
	}

	return items
}

// flattenCapacityCommitmentPlanEnumSlice flattens the contents of CapacityCommitmentPlanEnum from a JSON
// response object.
func flattenCapacityCommitmentPlanEnumSlice(c *Client, i interface{}, res *CapacityCommitment) []CapacityCommitmentPlanEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []CapacityCommitmentPlanEnum{}
	}

	if len(a) == 0 {
		return []CapacityCommitmentPlanEnum{}
	}

	items := make([]CapacityCommitmentPlanEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCapacityCommitmentPlanEnum(item.(interface{})))
	}

	return items
}

// flattenCapacityCommitmentPlanEnum asserts that an interface is a string, and returns a
// pointer to a *CapacityCommitmentPlanEnum with the same value as that string.
func flattenCapacityCommitmentPlanEnum(i interface{}) *CapacityCommitmentPlanEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return CapacityCommitmentPlanEnumRef(s)
}

// flattenCapacityCommitmentStateEnumMap flattens the contents of CapacityCommitmentStateEnum from a JSON
// response object.
func flattenCapacityCommitmentStateEnumMap(c *Client, i interface{}, res *CapacityCommitment) map[string]CapacityCommitmentStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CapacityCommitmentStateEnum{}
	}

	if len(a) == 0 {
		return map[string]CapacityCommitmentStateEnum{}
	}

	items := make(map[string]CapacityCommitmentStateEnum)
	for k, item := range a {
		items[k] = *flattenCapacityCommitmentStateEnum(item.(interface{}))
	}

	return items
}

// flattenCapacityCommitmentStateEnumSlice flattens the contents of CapacityCommitmentStateEnum from a JSON
// response object.
func flattenCapacityCommitmentStateEnumSlice(c *Client, i interface{}, res *CapacityCommitment) []CapacityCommitmentStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []CapacityCommitmentStateEnum{}
	}

	if len(a) == 0 {
		return []CapacityCommitmentStateEnum{}
	}

	items := make([]CapacityCommitmentStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCapacityCommitmentStateEnum(item.(interface{})))
	}

	return items
}

// flattenCapacityCommitmentStateEnum asserts that an interface is a string, and returns a
// pointer to a *CapacityCommitmentStateEnum with the same value as that string.
func flattenCapacityCommitmentStateEnum(i interface{}) *CapacityCommitmentStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return CapacityCommitmentStateEnumRef(s)
}

// flattenCapacityCommitmentRenewalPlanEnumMap flattens the contents of CapacityCommitmentRenewalPlanEnum from a JSON
// response object.
func flattenCapacityCommitmentRenewalPlanEnumMap(c *Client, i interface{}, res *CapacityCommitment) map[string]CapacityCommitmentRenewalPlanEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CapacityCommitmentRenewalPlanEnum{}
	}

	if len(a) == 0 {
		return map[string]CapacityCommitmentRenewalPlanEnum{}
	}

	items := make(map[string]CapacityCommitmentRenewalPlanEnum)
	for k, item := range a {
		items[k] = *flattenCapacityCommitmentRenewalPlanEnum(item.(interface{}))
	}

	return items
}

// flattenCapacityCommitmentRenewalPlanEnumSlice flattens the contents of CapacityCommitmentRenewalPlanEnum from a JSON
// response object.
func flattenCapacityCommitmentRenewalPlanEnumSlice(c *Client, i interface{}, res *CapacityCommitment) []CapacityCommitmentRenewalPlanEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []CapacityCommitmentRenewalPlanEnum{}
	}

	if len(a) == 0 {
		return []CapacityCommitmentRenewalPlanEnum{}
	}

	items := make([]CapacityCommitmentRenewalPlanEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCapacityCommitmentRenewalPlanEnum(item.(interface{})))
	}

	return items
}

// flattenCapacityCommitmentRenewalPlanEnum asserts that an interface is a string, and returns a
// pointer to a *CapacityCommitmentRenewalPlanEnum with the same value as that string.
func flattenCapacityCommitmentRenewalPlanEnum(i interface{}) *CapacityCommitmentRenewalPlanEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return CapacityCommitmentRenewalPlanEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *CapacityCommitment) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalCapacityCommitment(b, c, r)
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

type capacityCommitmentDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         capacityCommitmentApiOperation
}

func convertFieldDiffsToCapacityCommitmentDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]capacityCommitmentDiff, error) {
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
	var diffs []capacityCommitmentDiff
	// For each operation name, create a capacityCommitmentDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := capacityCommitmentDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToCapacityCommitmentApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToCapacityCommitmentApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (capacityCommitmentApiOperation, error) {
	switch opName {

	case "updateCapacityCommitmentUpdateCapacityCommitmentOperation":
		return &updateCapacityCommitmentUpdateCapacityCommitmentOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractCapacityCommitmentFields(r *CapacityCommitment) error {
	vFailureStatus := r.FailureStatus
	if vFailureStatus == nil {
		// note: explicitly not the empty object.
		vFailureStatus = &CapacityCommitmentFailureStatus{}
	}
	if err := extractCapacityCommitmentFailureStatusFields(r, vFailureStatus); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vFailureStatus) {
		r.FailureStatus = vFailureStatus
	}
	return nil
}
func extractCapacityCommitmentFailureStatusFields(r *CapacityCommitment, o *CapacityCommitmentFailureStatus) error {
	return nil
}
func extractCapacityCommitmentFailureStatusDetailsFields(r *CapacityCommitment, o *CapacityCommitmentFailureStatusDetails) error {
	return nil
}

func postReadExtractCapacityCommitmentFields(r *CapacityCommitment) error {
	vFailureStatus := r.FailureStatus
	if vFailureStatus == nil {
		// note: explicitly not the empty object.
		vFailureStatus = &CapacityCommitmentFailureStatus{}
	}
	if err := postReadExtractCapacityCommitmentFailureStatusFields(r, vFailureStatus); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vFailureStatus) {
		r.FailureStatus = vFailureStatus
	}
	return nil
}
func postReadExtractCapacityCommitmentFailureStatusFields(r *CapacityCommitment, o *CapacityCommitmentFailureStatus) error {
	return nil
}
func postReadExtractCapacityCommitmentFailureStatusDetailsFields(r *CapacityCommitment, o *CapacityCommitmentFailureStatusDetails) error {
	return nil
}
