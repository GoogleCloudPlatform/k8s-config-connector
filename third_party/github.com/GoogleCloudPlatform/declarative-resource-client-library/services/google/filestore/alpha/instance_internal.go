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

func (r *Instance) validate() error {

	if err := dcl.RequiredParameter(r.Name, "Name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	return nil
}
func (r *InstanceFileShares) validate() error {
	return nil
}
func (r *InstanceFileSharesNfsExportOptions) validate() error {
	return nil
}
func (r *InstanceNetworks) validate() error {
	return nil
}
func (r *Instance) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://file.googleapis.com/v1beta1/", params)
}

func (r *Instance) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/instances/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Instance) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/instances", nr.basePath(), userBasePath, params), nil

}

func (r *Instance) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/instances?instanceId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *Instance) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/instances/{{name}}", nr.basePath(), userBasePath, params), nil
}

// instanceApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type instanceApiOperation interface {
	do(context.Context, *Instance, *Client) error
}

// newUpdateInstanceUpdateInstanceRequest creates a request for an
// Instance resource's UpdateInstance update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateInstanceUpdateInstanceRequest(ctx context.Context, f *Instance, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v, err := expandInstanceFileSharesSlice(c, f.FileShares, res); err != nil {
		return nil, fmt.Errorf("error expanding FileShares into fileShares: %w", err)
	} else if v != nil {
		req["fileShares"] = v
	}
	b, err := c.getInstanceRaw(ctx, f)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	rawEtag, err := dcl.GetMapEntry(
		m,
		[]string{"etag"},
	)
	if err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "Failed to fetch from JSON Path: %v", err)
	} else {
		req["etag"] = rawEtag.(string)
	}
	return req, nil
}

// marshalUpdateInstanceUpdateInstanceRequest converts the update into
// the final JSON request body.
func marshalUpdateInstanceUpdateInstanceRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateInstanceUpdateInstanceOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateInstanceUpdateInstanceOperation) do(ctx context.Context, r *Instance, c *Client) error {
	_, err := c.GetInstance(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateInstance")
	if err != nil {
		return err
	}
	mask := dcl.TopLevelUpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateInstanceUpdateInstanceRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateInstanceUpdateInstanceRequest(c, req)
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

func (c *Client) listInstanceRaw(ctx context.Context, r *Instance, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != InstanceMaxPage {
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

type listInstanceOperation struct {
	Instances []map[string]interface{} `json:"instances"`
	Token     string                   `json:"nextPageToken"`
}

func (c *Client) listInstance(ctx context.Context, r *Instance, pageToken string, pageSize int32) ([]*Instance, string, error) {
	b, err := c.listInstanceRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listInstanceOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Instance
	for _, v := range m.Instances {
		res, err := unmarshalMapInstance(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllInstance(ctx context.Context, f func(*Instance) bool, resources []*Instance) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteInstance(ctx, res)
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

type deleteInstanceOperation struct{}

func (op *deleteInstanceOperation) do(ctx context.Context, r *Instance, c *Client) error {
	r, err := c.GetInstance(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Instance not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetInstance checking for existence. error: %v", err)
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
		_, err := c.GetInstance(ctx, r)
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
type createInstanceOperation struct {
	response map[string]interface{}
}

func (op *createInstanceOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (c *Client) getInstanceRaw(ctx context.Context, r *Instance) ([]byte, error) {

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

func (c *Client) instanceDiffsForRawDesired(ctx context.Context, rawDesired *Instance, opts ...dcl.ApplyOption) (initial, desired *Instance, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Instance
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Instance); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Instance, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetInstance(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Instance resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Instance resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Instance resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeInstanceDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Instance: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Instance: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractInstanceFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeInstanceInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Instance: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeInstanceDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Instance: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffInstance(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeInstanceInitialState(rawInitial, rawDesired *Instance) (*Instance, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeInstanceDesiredState(rawDesired, rawInitial *Instance, opts ...dcl.ApplyOption) (*Instance, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &Instance{}
	if dcl.NameToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.IsZeroValue(rawDesired.Tier) || (dcl.IsEmptyValueIndirect(rawDesired.Tier) && dcl.IsEmptyValueIndirect(rawInitial.Tier)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Tier = rawInitial.Tier
	} else {
		canonicalDesired.Tier = rawDesired.Tier
	}
	if dcl.IsZeroValue(rawDesired.Labels) || (dcl.IsEmptyValueIndirect(rawDesired.Labels) && dcl.IsEmptyValueIndirect(rawInitial.Labels)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	canonicalDesired.FileShares = canonicalizeInstanceFileSharesSlice(rawDesired.FileShares, rawInitial.FileShares, opts...)
	canonicalDesired.Networks = canonicalizeInstanceNetworksSlice(rawDesired.Networks, rawInitial.Networks, opts...)
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

func canonicalizeInstanceNewState(c *Client, rawNew, rawDesired *Instance) (*Instance, error) {

	rawNew.Name = rawDesired.Name

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

	if dcl.IsEmptyValueIndirect(rawNew.StatusMessage) && dcl.IsEmptyValueIndirect(rawDesired.StatusMessage) {
		rawNew.StatusMessage = rawDesired.StatusMessage
	} else {
		if dcl.StringCanonicalize(rawDesired.StatusMessage, rawNew.StatusMessage) {
			rawNew.StatusMessage = rawDesired.StatusMessage
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Tier) && dcl.IsEmptyValueIndirect(rawDesired.Tier) {
		rawNew.Tier = rawDesired.Tier
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.FileShares) && dcl.IsEmptyValueIndirect(rawDesired.FileShares) {
		rawNew.FileShares = rawDesired.FileShares
	} else {
		rawNew.FileShares = canonicalizeNewInstanceFileSharesSlice(c, rawDesired.FileShares, rawNew.FileShares)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Networks) && dcl.IsEmptyValueIndirect(rawDesired.Networks) {
		rawNew.Networks = rawDesired.Networks
	} else {
		rawNew.Networks = canonicalizeNewInstanceNetworksSlice(c, rawDesired.Networks, rawNew.Networks)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Etag) && dcl.IsEmptyValueIndirect(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
		if dcl.StringCanonicalize(rawDesired.Etag, rawNew.Etag) {
			rawNew.Etag = rawDesired.Etag
		}
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeInstanceFileShares(des, initial *InstanceFileShares, opts ...dcl.ApplyOption) *InstanceFileShares {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceFileShares{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.IsZeroValue(des.CapacityGb) || (dcl.IsEmptyValueIndirect(des.CapacityGb) && dcl.IsEmptyValueIndirect(initial.CapacityGb)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.CapacityGb = initial.CapacityGb
	} else {
		cDes.CapacityGb = des.CapacityGb
	}
	if dcl.IsZeroValue(des.SourceBackup) || (dcl.IsEmptyValueIndirect(des.SourceBackup) && dcl.IsEmptyValueIndirect(initial.SourceBackup)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.SourceBackup = initial.SourceBackup
	} else {
		cDes.SourceBackup = des.SourceBackup
	}
	cDes.NfsExportOptions = canonicalizeInstanceFileSharesNfsExportOptionsSlice(des.NfsExportOptions, initial.NfsExportOptions, opts...)

	return cDes
}

func canonicalizeInstanceFileSharesSlice(des, initial []InstanceFileShares, opts ...dcl.ApplyOption) []InstanceFileShares {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InstanceFileShares, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInstanceFileShares(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InstanceFileShares, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInstanceFileShares(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInstanceFileShares(c *Client, des, nw *InstanceFileShares) *InstanceFileShares {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InstanceFileShares while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	nw.NfsExportOptions = canonicalizeNewInstanceFileSharesNfsExportOptionsSlice(c, des.NfsExportOptions, nw.NfsExportOptions)

	return nw
}

func canonicalizeNewInstanceFileSharesSet(c *Client, des, nw []InstanceFileShares) []InstanceFileShares {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InstanceFileShares
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInstanceFileSharesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInstanceFileShares(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInstanceFileSharesSlice(c *Client, des, nw []InstanceFileShares) []InstanceFileShares {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceFileShares
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceFileShares(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceFileSharesNfsExportOptions(des, initial *InstanceFileSharesNfsExportOptions, opts ...dcl.ApplyOption) *InstanceFileSharesNfsExportOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceFileSharesNfsExportOptions{}

	if dcl.StringArrayCanonicalize(des.IPRanges, initial.IPRanges) {
		cDes.IPRanges = initial.IPRanges
	} else {
		cDes.IPRanges = des.IPRanges
	}
	if dcl.IsZeroValue(des.AccessMode) || (dcl.IsEmptyValueIndirect(des.AccessMode) && dcl.IsEmptyValueIndirect(initial.AccessMode)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.AccessMode = initial.AccessMode
	} else {
		cDes.AccessMode = des.AccessMode
	}
	if dcl.IsZeroValue(des.SquashMode) || (dcl.IsEmptyValueIndirect(des.SquashMode) && dcl.IsEmptyValueIndirect(initial.SquashMode)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.SquashMode = initial.SquashMode
	} else {
		cDes.SquashMode = des.SquashMode
	}
	if dcl.IsZeroValue(des.AnonUid) || (dcl.IsEmptyValueIndirect(des.AnonUid) && dcl.IsEmptyValueIndirect(initial.AnonUid)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.AnonUid = initial.AnonUid
	} else {
		cDes.AnonUid = des.AnonUid
	}
	if dcl.IsZeroValue(des.AnonGid) || (dcl.IsEmptyValueIndirect(des.AnonGid) && dcl.IsEmptyValueIndirect(initial.AnonGid)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.AnonGid = initial.AnonGid
	} else {
		cDes.AnonGid = des.AnonGid
	}

	return cDes
}

func canonicalizeInstanceFileSharesNfsExportOptionsSlice(des, initial []InstanceFileSharesNfsExportOptions, opts ...dcl.ApplyOption) []InstanceFileSharesNfsExportOptions {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InstanceFileSharesNfsExportOptions, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInstanceFileSharesNfsExportOptions(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InstanceFileSharesNfsExportOptions, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInstanceFileSharesNfsExportOptions(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInstanceFileSharesNfsExportOptions(c *Client, des, nw *InstanceFileSharesNfsExportOptions) *InstanceFileSharesNfsExportOptions {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InstanceFileSharesNfsExportOptions while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.IPRanges, nw.IPRanges) {
		nw.IPRanges = des.IPRanges
	}

	return nw
}

func canonicalizeNewInstanceFileSharesNfsExportOptionsSet(c *Client, des, nw []InstanceFileSharesNfsExportOptions) []InstanceFileSharesNfsExportOptions {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InstanceFileSharesNfsExportOptions
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInstanceFileSharesNfsExportOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInstanceFileSharesNfsExportOptions(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInstanceFileSharesNfsExportOptionsSlice(c *Client, des, nw []InstanceFileSharesNfsExportOptions) []InstanceFileSharesNfsExportOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceFileSharesNfsExportOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceFileSharesNfsExportOptions(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceNetworks(des, initial *InstanceNetworks, opts ...dcl.ApplyOption) *InstanceNetworks {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceNetworks{}

	if dcl.IsZeroValue(des.Network) || (dcl.IsEmptyValueIndirect(des.Network) && dcl.IsEmptyValueIndirect(initial.Network)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Network = initial.Network
	} else {
		cDes.Network = des.Network
	}
	if dcl.IsZeroValue(des.Modes) || (dcl.IsEmptyValueIndirect(des.Modes) && dcl.IsEmptyValueIndirect(initial.Modes)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Modes = initial.Modes
	} else {
		cDes.Modes = des.Modes
	}
	if dcl.StringCanonicalize(des.ReservedIPRange, initial.ReservedIPRange) || dcl.IsZeroValue(des.ReservedIPRange) {
		cDes.ReservedIPRange = initial.ReservedIPRange
	} else {
		cDes.ReservedIPRange = des.ReservedIPRange
	}

	return cDes
}

func canonicalizeInstanceNetworksSlice(des, initial []InstanceNetworks, opts ...dcl.ApplyOption) []InstanceNetworks {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InstanceNetworks, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInstanceNetworks(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InstanceNetworks, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInstanceNetworks(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInstanceNetworks(c *Client, des, nw *InstanceNetworks) *InstanceNetworks {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InstanceNetworks while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.ReservedIPRange, nw.ReservedIPRange) {
		nw.ReservedIPRange = des.ReservedIPRange
	}
	if dcl.StringArrayCanonicalize(des.IPAddresses, nw.IPAddresses) {
		nw.IPAddresses = des.IPAddresses
	}

	return nw
}

func canonicalizeNewInstanceNetworksSet(c *Client, des, nw []InstanceNetworks) []InstanceNetworks {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InstanceNetworks
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInstanceNetworksNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInstanceNetworks(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInstanceNetworksSlice(c *Client, des, nw []InstanceNetworks) []InstanceNetworks {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceNetworks
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceNetworks(c, &d, &n))
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
func diffInstance(c *Client, desired, actual *Instance, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.StatusMessage, actual.StatusMessage, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StatusMessage")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Tier, actual.Tier, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Tier")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FileShares, actual.FileShares, dcl.DiffInfo{ObjectFunction: compareInstanceFileSharesNewStyle, EmptyObject: EmptyInstanceFileShares, OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("FileShares")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Networks, actual.Networks, dcl.DiffInfo{ObjectFunction: compareInstanceNetworksNewStyle, EmptyObject: EmptyInstanceNetworks, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Networks")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Etag, actual.Etag, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Etag")); len(ds) != 0 || err != nil {
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

	if len(newDiffs) > 0 {
		c.Config.Logger.Infof("Diff function found diffs: %v", newDiffs)
	}
	return newDiffs, nil
}
func compareInstanceFileSharesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceFileShares)
	if !ok {
		desiredNotPointer, ok := d.(InstanceFileShares)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceFileShares or *InstanceFileShares", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceFileShares)
	if !ok {
		actualNotPointer, ok := a.(InstanceFileShares)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceFileShares", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CapacityGb, actual.CapacityGb, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("CapacityGb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceBackup, actual.SourceBackup, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("SourceBackup")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NfsExportOptions, actual.NfsExportOptions, dcl.DiffInfo{ObjectFunction: compareInstanceFileSharesNfsExportOptionsNewStyle, EmptyObject: EmptyInstanceFileSharesNfsExportOptions, OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("NfsExportOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceFileSharesNfsExportOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceFileSharesNfsExportOptions)
	if !ok {
		desiredNotPointer, ok := d.(InstanceFileSharesNfsExportOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceFileSharesNfsExportOptions or *InstanceFileSharesNfsExportOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceFileSharesNfsExportOptions)
	if !ok {
		actualNotPointer, ok := a.(InstanceFileSharesNfsExportOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceFileSharesNfsExportOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.IPRanges, actual.IPRanges, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("IpRanges")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AccessMode, actual.AccessMode, dcl.DiffInfo{ServerDefault: true, Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("AccessMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SquashMode, actual.SquashMode, dcl.DiffInfo{ServerDefault: true, Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("SquashMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AnonUid, actual.AnonUid, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("AnonUid")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AnonGid, actual.AnonGid, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("AnonGid")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceNetworksNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceNetworks)
	if !ok {
		desiredNotPointer, ok := d.(InstanceNetworks)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceNetworks or *InstanceNetworks", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceNetworks)
	if !ok {
		actualNotPointer, ok := a.(InstanceNetworks)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceNetworks", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Network, actual.Network, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Network")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Modes, actual.Modes, dcl.DiffInfo{ServerDefault: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Modes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReservedIPRange, actual.ReservedIPRange, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ReservedIpRange")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IPAddresses, actual.IPAddresses, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IpAddresses")); len(ds) != 0 || err != nil {
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
func (r *Instance) urlNormalized() *Instance {
	normalized := dcl.Copy(*r).(Instance)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.StatusMessage = dcl.SelfLinkToName(r.StatusMessage)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Instance) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateInstance" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/instances/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Instance resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Instance) marshal(c *Client) ([]byte, error) {
	m, err := expandInstance(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Instance: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalInstance decodes JSON responses into the Instance resource schema.
func unmarshalInstance(b []byte, c *Client, res *Instance) (*Instance, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapInstance(m, c, res)
}

func unmarshalMapInstance(m map[string]interface{}, c *Client, res *Instance) (*Instance, error) {

	flattened := flattenInstance(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandInstance expands Instance into a JSON request object.
func expandInstance(c *Client, f *Instance) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v := f.Tier; dcl.ValueShouldBeSent(v) {
		m["tier"] = v
	}
	if v := f.Labels; dcl.ValueShouldBeSent(v) {
		m["labels"] = v
	}
	if v, err := expandInstanceFileSharesSlice(c, f.FileShares, res); err != nil {
		return nil, fmt.Errorf("error expanding FileShares into fileShares: %w", err)
	} else if v != nil {
		m["fileShares"] = v
	}
	if v, err := expandInstanceNetworksSlice(c, f.Networks, res); err != nil {
		return nil, fmt.Errorf("error expanding Networks into networks: %w", err)
	} else if v != nil {
		m["networks"] = v
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

// flattenInstance flattens Instance from a JSON request object into the
// Instance type.
func flattenInstance(c *Client, i interface{}, res *Instance) *Instance {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Instance{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.State = flattenInstanceStateEnum(m["state"])
	resultRes.StatusMessage = dcl.FlattenString(m["statusMessage"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.Tier = flattenInstanceTierEnum(m["tier"])
	resultRes.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	resultRes.FileShares = flattenInstanceFileSharesSlice(c, m["fileShares"], res)
	resultRes.Networks = flattenInstanceNetworksSlice(c, m["networks"], res)
	resultRes.Etag = dcl.FlattenString(m["etag"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// expandInstanceFileSharesMap expands the contents of InstanceFileShares into a JSON
// request object.
func expandInstanceFileSharesMap(c *Client, f map[string]InstanceFileShares, res *Instance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceFileShares(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceFileSharesSlice expands the contents of InstanceFileShares into a JSON
// request object.
func expandInstanceFileSharesSlice(c *Client, f []InstanceFileShares, res *Instance) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceFileShares(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceFileSharesMap flattens the contents of InstanceFileShares from a JSON
// response object.
func flattenInstanceFileSharesMap(c *Client, i interface{}, res *Instance) map[string]InstanceFileShares {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceFileShares{}
	}

	if len(a) == 0 {
		return map[string]InstanceFileShares{}
	}

	items := make(map[string]InstanceFileShares)
	for k, item := range a {
		items[k] = *flattenInstanceFileShares(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInstanceFileSharesSlice flattens the contents of InstanceFileShares from a JSON
// response object.
func flattenInstanceFileSharesSlice(c *Client, i interface{}, res *Instance) []InstanceFileShares {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceFileShares{}
	}

	if len(a) == 0 {
		return []InstanceFileShares{}
	}

	items := make([]InstanceFileShares, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceFileShares(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInstanceFileShares expands an instance of InstanceFileShares into a JSON
// request object.
func expandInstanceFileShares(c *Client, f *InstanceFileShares, res *Instance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.CapacityGb; !dcl.IsEmptyValueIndirect(v) {
		m["capacityGb"] = v
	}
	if v := f.SourceBackup; !dcl.IsEmptyValueIndirect(v) {
		m["sourceBackup"] = v
	}
	if v, err := expandInstanceFileSharesNfsExportOptionsSlice(c, f.NfsExportOptions, res); err != nil {
		return nil, fmt.Errorf("error expanding NfsExportOptions into nfsExportOptions: %w", err)
	} else if v != nil {
		m["nfsExportOptions"] = v
	}

	return m, nil
}

// flattenInstanceFileShares flattens an instance of InstanceFileShares from a JSON
// response object.
func flattenInstanceFileShares(c *Client, i interface{}, res *Instance) *InstanceFileShares {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceFileShares{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceFileShares
	}
	r.Name = dcl.FlattenString(m["name"])
	r.CapacityGb = dcl.FlattenInteger(m["capacityGb"])
	r.SourceBackup = dcl.FlattenString(m["sourceBackup"])
	r.NfsExportOptions = flattenInstanceFileSharesNfsExportOptionsSlice(c, m["nfsExportOptions"], res)

	return r
}

// expandInstanceFileSharesNfsExportOptionsMap expands the contents of InstanceFileSharesNfsExportOptions into a JSON
// request object.
func expandInstanceFileSharesNfsExportOptionsMap(c *Client, f map[string]InstanceFileSharesNfsExportOptions, res *Instance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceFileSharesNfsExportOptions(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceFileSharesNfsExportOptionsSlice expands the contents of InstanceFileSharesNfsExportOptions into a JSON
// request object.
func expandInstanceFileSharesNfsExportOptionsSlice(c *Client, f []InstanceFileSharesNfsExportOptions, res *Instance) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceFileSharesNfsExportOptions(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceFileSharesNfsExportOptionsMap flattens the contents of InstanceFileSharesNfsExportOptions from a JSON
// response object.
func flattenInstanceFileSharesNfsExportOptionsMap(c *Client, i interface{}, res *Instance) map[string]InstanceFileSharesNfsExportOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceFileSharesNfsExportOptions{}
	}

	if len(a) == 0 {
		return map[string]InstanceFileSharesNfsExportOptions{}
	}

	items := make(map[string]InstanceFileSharesNfsExportOptions)
	for k, item := range a {
		items[k] = *flattenInstanceFileSharesNfsExportOptions(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInstanceFileSharesNfsExportOptionsSlice flattens the contents of InstanceFileSharesNfsExportOptions from a JSON
// response object.
func flattenInstanceFileSharesNfsExportOptionsSlice(c *Client, i interface{}, res *Instance) []InstanceFileSharesNfsExportOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceFileSharesNfsExportOptions{}
	}

	if len(a) == 0 {
		return []InstanceFileSharesNfsExportOptions{}
	}

	items := make([]InstanceFileSharesNfsExportOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceFileSharesNfsExportOptions(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInstanceFileSharesNfsExportOptions expands an instance of InstanceFileSharesNfsExportOptions into a JSON
// request object.
func expandInstanceFileSharesNfsExportOptions(c *Client, f *InstanceFileSharesNfsExportOptions, res *Instance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.IPRanges; v != nil {
		m["ipRanges"] = v
	}
	if v := f.AccessMode; !dcl.IsEmptyValueIndirect(v) {
		m["accessMode"] = v
	}
	if v := f.SquashMode; !dcl.IsEmptyValueIndirect(v) {
		m["squashMode"] = v
	}
	if v := f.AnonUid; !dcl.IsEmptyValueIndirect(v) {
		m["anonUid"] = v
	}
	if v := f.AnonGid; !dcl.IsEmptyValueIndirect(v) {
		m["anonGid"] = v
	}

	return m, nil
}

// flattenInstanceFileSharesNfsExportOptions flattens an instance of InstanceFileSharesNfsExportOptions from a JSON
// response object.
func flattenInstanceFileSharesNfsExportOptions(c *Client, i interface{}, res *Instance) *InstanceFileSharesNfsExportOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceFileSharesNfsExportOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceFileSharesNfsExportOptions
	}
	r.IPRanges = dcl.FlattenStringSlice(m["ipRanges"])
	r.AccessMode = flattenInstanceFileSharesNfsExportOptionsAccessModeEnum(m["accessMode"])
	r.SquashMode = flattenInstanceFileSharesNfsExportOptionsSquashModeEnum(m["squashMode"])
	r.AnonUid = dcl.FlattenInteger(m["anonUid"])
	r.AnonGid = dcl.FlattenInteger(m["anonGid"])

	return r
}

// expandInstanceNetworksMap expands the contents of InstanceNetworks into a JSON
// request object.
func expandInstanceNetworksMap(c *Client, f map[string]InstanceNetworks, res *Instance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceNetworks(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceNetworksSlice expands the contents of InstanceNetworks into a JSON
// request object.
func expandInstanceNetworksSlice(c *Client, f []InstanceNetworks, res *Instance) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceNetworks(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceNetworksMap flattens the contents of InstanceNetworks from a JSON
// response object.
func flattenInstanceNetworksMap(c *Client, i interface{}, res *Instance) map[string]InstanceNetworks {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceNetworks{}
	}

	if len(a) == 0 {
		return map[string]InstanceNetworks{}
	}

	items := make(map[string]InstanceNetworks)
	for k, item := range a {
		items[k] = *flattenInstanceNetworks(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInstanceNetworksSlice flattens the contents of InstanceNetworks from a JSON
// response object.
func flattenInstanceNetworksSlice(c *Client, i interface{}, res *Instance) []InstanceNetworks {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceNetworks{}
	}

	if len(a) == 0 {
		return []InstanceNetworks{}
	}

	items := make([]InstanceNetworks, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceNetworks(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInstanceNetworks expands an instance of InstanceNetworks into a JSON
// request object.
func expandInstanceNetworks(c *Client, f *InstanceNetworks, res *Instance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Network; !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v := f.Modes; v != nil {
		m["modes"] = v
	}
	if v := f.ReservedIPRange; !dcl.IsEmptyValueIndirect(v) {
		m["reservedIpRange"] = v
	}

	return m, nil
}

// flattenInstanceNetworks flattens an instance of InstanceNetworks from a JSON
// response object.
func flattenInstanceNetworks(c *Client, i interface{}, res *Instance) *InstanceNetworks {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceNetworks{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceNetworks
	}
	r.Network = dcl.FlattenString(m["network"])
	r.Modes = flattenInstanceNetworksModesEnumSlice(c, m["modes"], res)
	r.ReservedIPRange = dcl.FlattenString(m["reservedIpRange"])
	r.IPAddresses = dcl.FlattenStringSlice(m["ipAddresses"])

	return r
}

// flattenInstanceStateEnumMap flattens the contents of InstanceStateEnum from a JSON
// response object.
func flattenInstanceStateEnumMap(c *Client, i interface{}, res *Instance) map[string]InstanceStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceStateEnum{}
	}

	if len(a) == 0 {
		return map[string]InstanceStateEnum{}
	}

	items := make(map[string]InstanceStateEnum)
	for k, item := range a {
		items[k] = *flattenInstanceStateEnum(item.(interface{}))
	}

	return items
}

// flattenInstanceStateEnumSlice flattens the contents of InstanceStateEnum from a JSON
// response object.
func flattenInstanceStateEnumSlice(c *Client, i interface{}, res *Instance) []InstanceStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceStateEnum{}
	}

	if len(a) == 0 {
		return []InstanceStateEnum{}
	}

	items := make([]InstanceStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceStateEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceStateEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceStateEnum with the same value as that string.
func flattenInstanceStateEnum(i interface{}) *InstanceStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return InstanceStateEnumRef(s)
}

// flattenInstanceTierEnumMap flattens the contents of InstanceTierEnum from a JSON
// response object.
func flattenInstanceTierEnumMap(c *Client, i interface{}, res *Instance) map[string]InstanceTierEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceTierEnum{}
	}

	if len(a) == 0 {
		return map[string]InstanceTierEnum{}
	}

	items := make(map[string]InstanceTierEnum)
	for k, item := range a {
		items[k] = *flattenInstanceTierEnum(item.(interface{}))
	}

	return items
}

// flattenInstanceTierEnumSlice flattens the contents of InstanceTierEnum from a JSON
// response object.
func flattenInstanceTierEnumSlice(c *Client, i interface{}, res *Instance) []InstanceTierEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTierEnum{}
	}

	if len(a) == 0 {
		return []InstanceTierEnum{}
	}

	items := make([]InstanceTierEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTierEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceTierEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceTierEnum with the same value as that string.
func flattenInstanceTierEnum(i interface{}) *InstanceTierEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return InstanceTierEnumRef(s)
}

// flattenInstanceFileSharesNfsExportOptionsAccessModeEnumMap flattens the contents of InstanceFileSharesNfsExportOptionsAccessModeEnum from a JSON
// response object.
func flattenInstanceFileSharesNfsExportOptionsAccessModeEnumMap(c *Client, i interface{}, res *Instance) map[string]InstanceFileSharesNfsExportOptionsAccessModeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceFileSharesNfsExportOptionsAccessModeEnum{}
	}

	if len(a) == 0 {
		return map[string]InstanceFileSharesNfsExportOptionsAccessModeEnum{}
	}

	items := make(map[string]InstanceFileSharesNfsExportOptionsAccessModeEnum)
	for k, item := range a {
		items[k] = *flattenInstanceFileSharesNfsExportOptionsAccessModeEnum(item.(interface{}))
	}

	return items
}

// flattenInstanceFileSharesNfsExportOptionsAccessModeEnumSlice flattens the contents of InstanceFileSharesNfsExportOptionsAccessModeEnum from a JSON
// response object.
func flattenInstanceFileSharesNfsExportOptionsAccessModeEnumSlice(c *Client, i interface{}, res *Instance) []InstanceFileSharesNfsExportOptionsAccessModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceFileSharesNfsExportOptionsAccessModeEnum{}
	}

	if len(a) == 0 {
		return []InstanceFileSharesNfsExportOptionsAccessModeEnum{}
	}

	items := make([]InstanceFileSharesNfsExportOptionsAccessModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceFileSharesNfsExportOptionsAccessModeEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceFileSharesNfsExportOptionsAccessModeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceFileSharesNfsExportOptionsAccessModeEnum with the same value as that string.
func flattenInstanceFileSharesNfsExportOptionsAccessModeEnum(i interface{}) *InstanceFileSharesNfsExportOptionsAccessModeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return InstanceFileSharesNfsExportOptionsAccessModeEnumRef(s)
}

// flattenInstanceFileSharesNfsExportOptionsSquashModeEnumMap flattens the contents of InstanceFileSharesNfsExportOptionsSquashModeEnum from a JSON
// response object.
func flattenInstanceFileSharesNfsExportOptionsSquashModeEnumMap(c *Client, i interface{}, res *Instance) map[string]InstanceFileSharesNfsExportOptionsSquashModeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceFileSharesNfsExportOptionsSquashModeEnum{}
	}

	if len(a) == 0 {
		return map[string]InstanceFileSharesNfsExportOptionsSquashModeEnum{}
	}

	items := make(map[string]InstanceFileSharesNfsExportOptionsSquashModeEnum)
	for k, item := range a {
		items[k] = *flattenInstanceFileSharesNfsExportOptionsSquashModeEnum(item.(interface{}))
	}

	return items
}

// flattenInstanceFileSharesNfsExportOptionsSquashModeEnumSlice flattens the contents of InstanceFileSharesNfsExportOptionsSquashModeEnum from a JSON
// response object.
func flattenInstanceFileSharesNfsExportOptionsSquashModeEnumSlice(c *Client, i interface{}, res *Instance) []InstanceFileSharesNfsExportOptionsSquashModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceFileSharesNfsExportOptionsSquashModeEnum{}
	}

	if len(a) == 0 {
		return []InstanceFileSharesNfsExportOptionsSquashModeEnum{}
	}

	items := make([]InstanceFileSharesNfsExportOptionsSquashModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceFileSharesNfsExportOptionsSquashModeEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceFileSharesNfsExportOptionsSquashModeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceFileSharesNfsExportOptionsSquashModeEnum with the same value as that string.
func flattenInstanceFileSharesNfsExportOptionsSquashModeEnum(i interface{}) *InstanceFileSharesNfsExportOptionsSquashModeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return InstanceFileSharesNfsExportOptionsSquashModeEnumRef(s)
}

// flattenInstanceNetworksModesEnumMap flattens the contents of InstanceNetworksModesEnum from a JSON
// response object.
func flattenInstanceNetworksModesEnumMap(c *Client, i interface{}, res *Instance) map[string]InstanceNetworksModesEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceNetworksModesEnum{}
	}

	if len(a) == 0 {
		return map[string]InstanceNetworksModesEnum{}
	}

	items := make(map[string]InstanceNetworksModesEnum)
	for k, item := range a {
		items[k] = *flattenInstanceNetworksModesEnum(item.(interface{}))
	}

	return items
}

// flattenInstanceNetworksModesEnumSlice flattens the contents of InstanceNetworksModesEnum from a JSON
// response object.
func flattenInstanceNetworksModesEnumSlice(c *Client, i interface{}, res *Instance) []InstanceNetworksModesEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceNetworksModesEnum{}
	}

	if len(a) == 0 {
		return []InstanceNetworksModesEnum{}
	}

	items := make([]InstanceNetworksModesEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceNetworksModesEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceNetworksModesEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceNetworksModesEnum with the same value as that string.
func flattenInstanceNetworksModesEnum(i interface{}) *InstanceNetworksModesEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return InstanceNetworksModesEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Instance) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalInstance(b, c, r)
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

type instanceDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         instanceApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToInstanceDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]instanceDiff, error) {
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
	var diffs []instanceDiff
	// For each operation name, create a instanceDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := instanceDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToInstanceApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToInstanceApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (instanceApiOperation, error) {
	switch opName {

	case "updateInstanceUpdateInstanceOperation":
		return &updateInstanceUpdateInstanceOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractInstanceFields(r *Instance) error {
	return nil
}
func extractInstanceFileSharesFields(r *Instance, o *InstanceFileShares) error {
	return nil
}
func extractInstanceFileSharesNfsExportOptionsFields(r *Instance, o *InstanceFileSharesNfsExportOptions) error {
	return nil
}
func extractInstanceNetworksFields(r *Instance, o *InstanceNetworks) error {
	return nil
}

func postReadExtractInstanceFields(r *Instance) error {
	return nil
}
func postReadExtractInstanceFileSharesFields(r *Instance, o *InstanceFileShares) error {
	return nil
}
func postReadExtractInstanceFileSharesNfsExportOptionsFields(r *Instance, o *InstanceFileSharesNfsExportOptions) error {
	return nil
}
func postReadExtractInstanceNetworksFields(r *Instance, o *InstanceNetworks) error {
	return nil
}
