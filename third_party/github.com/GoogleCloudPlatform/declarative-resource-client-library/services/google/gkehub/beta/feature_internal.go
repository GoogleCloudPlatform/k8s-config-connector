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

func (r *Feature) validate() error {

	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.ResourceState) {
		if err := r.ResourceState.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Spec) {
		if err := r.Spec.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.State) {
		if err := r.State.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FeatureResourceState) validate() error {
	return nil
}
func (r *FeatureSpec) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Multiclusteringress) {
		if err := r.Multiclusteringress.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Fleetobservability) {
		if err := r.Fleetobservability.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FeatureSpecMulticlusteringress) validate() error {
	if err := dcl.Required(r, "configMembership"); err != nil {
		return err
	}
	return nil
}
func (r *FeatureSpecFleetobservability) validate() error {
	if !dcl.IsEmptyValueIndirect(r.LoggingConfig) {
		if err := r.LoggingConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FeatureSpecFleetobservabilityLoggingConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.DefaultConfig) {
		if err := r.DefaultConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.FleetScopeLogsConfig) {
		if err := r.FleetScopeLogsConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FeatureSpecFleetobservabilityLoggingConfigDefaultConfig) validate() error {
	return nil
}
func (r *FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig) validate() error {
	return nil
}
func (r *FeatureState) validate() error {
	if !dcl.IsEmptyValueIndirect(r.State) {
		if err := r.State.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FeatureStateState) validate() error {
	return nil
}
func (r *Feature) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://gkehub.googleapis.com/v1beta1/", params)
}

// featureApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type featureApiOperation interface {
	do(context.Context, *Feature, *Client) error
}

// newUpdateFeatureUpdateFeatureRequest creates a request for an
// Feature resource's UpdateFeature update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateFeatureUpdateFeatureRequest(ctx context.Context, f *Feature, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v, err := expandFeatureSpec(c, f.Spec, res); err != nil {
		return nil, fmt.Errorf("error expanding Spec into spec: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["spec"] = v
	}
	return req, nil
}

// marshalUpdateFeatureUpdateFeatureRequest converts the update into
// the final JSON request body.
func marshalUpdateFeatureUpdateFeatureRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateFeatureUpdateFeatureOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (c *Client) listFeatureRaw(ctx context.Context, r *Feature, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != FeatureMaxPage {
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

type listFeatureOperation struct {
	Resources []map[string]interface{} `json:"resources"`
	Token     string                   `json:"nextPageToken"`
}

func (c *Client) listFeature(ctx context.Context, r *Feature, pageToken string, pageSize int32) ([]*Feature, string, error) {
	b, err := c.listFeatureRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listFeatureOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Feature
	for _, v := range m.Resources {
		res, err := unmarshalMapFeature(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllFeature(ctx context.Context, f func(*Feature) bool, resources []*Feature) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteFeature(ctx, res)
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

type deleteFeatureOperation struct{}

func (op *deleteFeatureOperation) do(ctx context.Context, r *Feature, c *Client) error {
	r, err := c.GetFeature(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Feature not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetFeature checking for existence. error: %v", err)
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
		_, err := c.GetFeature(ctx, r)
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
type createFeatureOperation struct {
	response map[string]interface{}
}

func (op *createFeatureOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createFeatureOperation) do(ctx context.Context, r *Feature, c *Client) error {
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

	if _, err := c.GetFeature(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getFeatureRaw(ctx context.Context, r *Feature) ([]byte, error) {

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

func (c *Client) featureDiffsForRawDesired(ctx context.Context, rawDesired *Feature, opts ...dcl.ApplyOption) (initial, desired *Feature, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Feature
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Feature); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Feature, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetFeature(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Feature resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Feature resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Feature resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeFeatureDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Feature: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Feature: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractFeatureFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeFeatureInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Feature: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeFeatureDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Feature: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffFeature(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeFeatureInitialState(rawInitial, rawDesired *Feature) (*Feature, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeFeatureDesiredState(rawDesired, rawInitial *Feature, opts ...dcl.ApplyOption) (*Feature, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.ResourceState = canonicalizeFeatureResourceState(rawDesired.ResourceState, nil, opts...)
		rawDesired.Spec = canonicalizeFeatureSpec(rawDesired.Spec, nil, opts...)
		rawDesired.State = canonicalizeFeatureState(rawDesired.State, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Feature{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.IsZeroValue(rawDesired.Labels) || (dcl.IsEmptyValueIndirect(rawDesired.Labels) && dcl.IsEmptyValueIndirect(rawInitial.Labels)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	canonicalDesired.Spec = canonicalizeFeatureSpec(rawDesired.Spec, rawInitial.Spec, opts...)
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

func canonicalizeFeatureNewState(c *Client, rawNew, rawDesired *Feature) (*Feature, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ResourceState) && dcl.IsEmptyValueIndirect(rawDesired.ResourceState) {
		rawNew.ResourceState = rawDesired.ResourceState
	} else {
		rawNew.ResourceState = canonicalizeNewFeatureResourceState(c, rawDesired.ResourceState, rawNew.ResourceState)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Spec) && dcl.IsEmptyValueIndirect(rawDesired.Spec) {
		rawNew.Spec = rawDesired.Spec
	} else {
		rawNew.Spec = canonicalizeNewFeatureSpec(c, rawDesired.Spec, rawNew.Spec)
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
		rawNew.State = canonicalizeNewFeatureState(c, rawDesired.State, rawNew.State)
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DeleteTime) && dcl.IsEmptyValueIndirect(rawDesired.DeleteTime) {
		rawNew.DeleteTime = rawDesired.DeleteTime
	} else {
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeFeatureResourceState(des, initial *FeatureResourceState, opts ...dcl.ApplyOption) *FeatureResourceState {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureResourceState{}

	return cDes
}

func canonicalizeFeatureResourceStateSlice(des, initial []FeatureResourceState, opts ...dcl.ApplyOption) []FeatureResourceState {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureResourceState, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureResourceState(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureResourceState, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureResourceState(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureResourceState(c *Client, des, nw *FeatureResourceState) *FeatureResourceState {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureResourceState while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.HasResources, nw.HasResources) {
		nw.HasResources = des.HasResources
	}

	return nw
}

func canonicalizeNewFeatureResourceStateSet(c *Client, des, nw []FeatureResourceState) []FeatureResourceState {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FeatureResourceState
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFeatureResourceStateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFeatureResourceState(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFeatureResourceStateSlice(c *Client, des, nw []FeatureResourceState) []FeatureResourceState {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureResourceState
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureResourceState(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureSpec(des, initial *FeatureSpec, opts ...dcl.ApplyOption) *FeatureSpec {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureSpec{}

	cDes.Multiclusteringress = canonicalizeFeatureSpecMulticlusteringress(des.Multiclusteringress, initial.Multiclusteringress, opts...)
	cDes.Fleetobservability = canonicalizeFeatureSpecFleetobservability(des.Fleetobservability, initial.Fleetobservability, opts...)

	return cDes
}

func canonicalizeFeatureSpecSlice(des, initial []FeatureSpec, opts ...dcl.ApplyOption) []FeatureSpec {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureSpec, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureSpec(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureSpec, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureSpec(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureSpec(c *Client, des, nw *FeatureSpec) *FeatureSpec {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureSpec while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Multiclusteringress = canonicalizeNewFeatureSpecMulticlusteringress(c, des.Multiclusteringress, nw.Multiclusteringress)
	nw.Fleetobservability = canonicalizeNewFeatureSpecFleetobservability(c, des.Fleetobservability, nw.Fleetobservability)

	return nw
}

func canonicalizeNewFeatureSpecSet(c *Client, des, nw []FeatureSpec) []FeatureSpec {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FeatureSpec
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFeatureSpecNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFeatureSpec(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFeatureSpecSlice(c *Client, des, nw []FeatureSpec) []FeatureSpec {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureSpec
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureSpec(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureSpecMulticlusteringress(des, initial *FeatureSpecMulticlusteringress, opts ...dcl.ApplyOption) *FeatureSpecMulticlusteringress {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureSpecMulticlusteringress{}

	if dcl.IsZeroValue(des.ConfigMembership) || (dcl.IsEmptyValueIndirect(des.ConfigMembership) && dcl.IsEmptyValueIndirect(initial.ConfigMembership)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ConfigMembership = initial.ConfigMembership
	} else {
		cDes.ConfigMembership = des.ConfigMembership
	}

	return cDes
}

func canonicalizeFeatureSpecMulticlusteringressSlice(des, initial []FeatureSpecMulticlusteringress, opts ...dcl.ApplyOption) []FeatureSpecMulticlusteringress {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureSpecMulticlusteringress, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureSpecMulticlusteringress(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureSpecMulticlusteringress, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureSpecMulticlusteringress(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureSpecMulticlusteringress(c *Client, des, nw *FeatureSpecMulticlusteringress) *FeatureSpecMulticlusteringress {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureSpecMulticlusteringress while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewFeatureSpecMulticlusteringressSet(c *Client, des, nw []FeatureSpecMulticlusteringress) []FeatureSpecMulticlusteringress {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FeatureSpecMulticlusteringress
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFeatureSpecMulticlusteringressNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFeatureSpecMulticlusteringress(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFeatureSpecMulticlusteringressSlice(c *Client, des, nw []FeatureSpecMulticlusteringress) []FeatureSpecMulticlusteringress {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureSpecMulticlusteringress
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureSpecMulticlusteringress(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureSpecFleetobservability(des, initial *FeatureSpecFleetobservability, opts ...dcl.ApplyOption) *FeatureSpecFleetobservability {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureSpecFleetobservability{}

	cDes.LoggingConfig = canonicalizeFeatureSpecFleetobservabilityLoggingConfig(des.LoggingConfig, initial.LoggingConfig, opts...)

	return cDes
}

func canonicalizeFeatureSpecFleetobservabilitySlice(des, initial []FeatureSpecFleetobservability, opts ...dcl.ApplyOption) []FeatureSpecFleetobservability {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureSpecFleetobservability, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureSpecFleetobservability(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureSpecFleetobservability, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureSpecFleetobservability(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureSpecFleetobservability(c *Client, des, nw *FeatureSpecFleetobservability) *FeatureSpecFleetobservability {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureSpecFleetobservability while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.LoggingConfig = canonicalizeNewFeatureSpecFleetobservabilityLoggingConfig(c, des.LoggingConfig, nw.LoggingConfig)

	return nw
}

func canonicalizeNewFeatureSpecFleetobservabilitySet(c *Client, des, nw []FeatureSpecFleetobservability) []FeatureSpecFleetobservability {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FeatureSpecFleetobservability
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFeatureSpecFleetobservabilityNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFeatureSpecFleetobservability(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFeatureSpecFleetobservabilitySlice(c *Client, des, nw []FeatureSpecFleetobservability) []FeatureSpecFleetobservability {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureSpecFleetobservability
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureSpecFleetobservability(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureSpecFleetobservabilityLoggingConfig(des, initial *FeatureSpecFleetobservabilityLoggingConfig, opts ...dcl.ApplyOption) *FeatureSpecFleetobservabilityLoggingConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureSpecFleetobservabilityLoggingConfig{}

	cDes.DefaultConfig = canonicalizeFeatureSpecFleetobservabilityLoggingConfigDefaultConfig(des.DefaultConfig, initial.DefaultConfig, opts...)
	cDes.FleetScopeLogsConfig = canonicalizeFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(des.FleetScopeLogsConfig, initial.FleetScopeLogsConfig, opts...)

	return cDes
}

func canonicalizeFeatureSpecFleetobservabilityLoggingConfigSlice(des, initial []FeatureSpecFleetobservabilityLoggingConfig, opts ...dcl.ApplyOption) []FeatureSpecFleetobservabilityLoggingConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureSpecFleetobservabilityLoggingConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureSpecFleetobservabilityLoggingConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureSpecFleetobservabilityLoggingConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureSpecFleetobservabilityLoggingConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureSpecFleetobservabilityLoggingConfig(c *Client, des, nw *FeatureSpecFleetobservabilityLoggingConfig) *FeatureSpecFleetobservabilityLoggingConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureSpecFleetobservabilityLoggingConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.DefaultConfig = canonicalizeNewFeatureSpecFleetobservabilityLoggingConfigDefaultConfig(c, des.DefaultConfig, nw.DefaultConfig)
	nw.FleetScopeLogsConfig = canonicalizeNewFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(c, des.FleetScopeLogsConfig, nw.FleetScopeLogsConfig)

	return nw
}

func canonicalizeNewFeatureSpecFleetobservabilityLoggingConfigSet(c *Client, des, nw []FeatureSpecFleetobservabilityLoggingConfig) []FeatureSpecFleetobservabilityLoggingConfig {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FeatureSpecFleetobservabilityLoggingConfig
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFeatureSpecFleetobservabilityLoggingConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFeatureSpecFleetobservabilityLoggingConfig(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFeatureSpecFleetobservabilityLoggingConfigSlice(c *Client, des, nw []FeatureSpecFleetobservabilityLoggingConfig) []FeatureSpecFleetobservabilityLoggingConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureSpecFleetobservabilityLoggingConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureSpecFleetobservabilityLoggingConfig(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureSpecFleetobservabilityLoggingConfigDefaultConfig(des, initial *FeatureSpecFleetobservabilityLoggingConfigDefaultConfig, opts ...dcl.ApplyOption) *FeatureSpecFleetobservabilityLoggingConfigDefaultConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureSpecFleetobservabilityLoggingConfigDefaultConfig{}

	if dcl.IsZeroValue(des.Mode) || (dcl.IsEmptyValueIndirect(des.Mode) && dcl.IsEmptyValueIndirect(initial.Mode)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Mode = initial.Mode
	} else {
		cDes.Mode = des.Mode
	}

	return cDes
}

func canonicalizeFeatureSpecFleetobservabilityLoggingConfigDefaultConfigSlice(des, initial []FeatureSpecFleetobservabilityLoggingConfigDefaultConfig, opts ...dcl.ApplyOption) []FeatureSpecFleetobservabilityLoggingConfigDefaultConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureSpecFleetobservabilityLoggingConfigDefaultConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureSpecFleetobservabilityLoggingConfigDefaultConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureSpecFleetobservabilityLoggingConfigDefaultConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureSpecFleetobservabilityLoggingConfigDefaultConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureSpecFleetobservabilityLoggingConfigDefaultConfig(c *Client, des, nw *FeatureSpecFleetobservabilityLoggingConfigDefaultConfig) *FeatureSpecFleetobservabilityLoggingConfigDefaultConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureSpecFleetobservabilityLoggingConfigDefaultConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewFeatureSpecFleetobservabilityLoggingConfigDefaultConfigSet(c *Client, des, nw []FeatureSpecFleetobservabilityLoggingConfigDefaultConfig) []FeatureSpecFleetobservabilityLoggingConfigDefaultConfig {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FeatureSpecFleetobservabilityLoggingConfigDefaultConfig
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFeatureSpecFleetobservabilityLoggingConfigDefaultConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFeatureSpecFleetobservabilityLoggingConfigDefaultConfig(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFeatureSpecFleetobservabilityLoggingConfigDefaultConfigSlice(c *Client, des, nw []FeatureSpecFleetobservabilityLoggingConfigDefaultConfig) []FeatureSpecFleetobservabilityLoggingConfigDefaultConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureSpecFleetobservabilityLoggingConfigDefaultConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureSpecFleetobservabilityLoggingConfigDefaultConfig(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(des, initial *FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig, opts ...dcl.ApplyOption) *FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig{}

	if dcl.IsZeroValue(des.Mode) || (dcl.IsEmptyValueIndirect(des.Mode) && dcl.IsEmptyValueIndirect(initial.Mode)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Mode = initial.Mode
	} else {
		cDes.Mode = des.Mode
	}

	return cDes
}

func canonicalizeFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigSlice(des, initial []FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig, opts ...dcl.ApplyOption) []FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(c *Client, des, nw *FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig) *FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigSet(c *Client, des, nw []FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig) []FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigSlice(c *Client, des, nw []FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig) []FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureState(des, initial *FeatureState, opts ...dcl.ApplyOption) *FeatureState {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureState{}

	return cDes
}

func canonicalizeFeatureStateSlice(des, initial []FeatureState, opts ...dcl.ApplyOption) []FeatureState {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureState, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureState(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureState, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureState(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureState(c *Client, des, nw *FeatureState) *FeatureState {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureState while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.State = canonicalizeNewFeatureStateState(c, des.State, nw.State)

	return nw
}

func canonicalizeNewFeatureStateSet(c *Client, des, nw []FeatureState) []FeatureState {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FeatureState
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFeatureStateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFeatureState(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFeatureStateSlice(c *Client, des, nw []FeatureState) []FeatureState {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureState
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureState(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureStateState(des, initial *FeatureStateState, opts ...dcl.ApplyOption) *FeatureStateState {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureStateState{}

	return cDes
}

func canonicalizeFeatureStateStateSlice(des, initial []FeatureStateState, opts ...dcl.ApplyOption) []FeatureStateState {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureStateState, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureStateState(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureStateState, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureStateState(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureStateState(c *Client, des, nw *FeatureStateState) *FeatureStateState {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureStateState while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}
	if dcl.StringCanonicalize(des.UpdateTime, nw.UpdateTime) {
		nw.UpdateTime = des.UpdateTime
	}

	return nw
}

func canonicalizeNewFeatureStateStateSet(c *Client, des, nw []FeatureStateState) []FeatureStateState {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FeatureStateState
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFeatureStateStateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFeatureStateState(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFeatureStateStateSlice(c *Client, des, nw []FeatureStateState) []FeatureStateState {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureStateState
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureStateState(c, &d, &n))
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
func diffFeature(c *Client, desired, actual *Feature, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateFeatureUpdateFeatureOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResourceState, actual.ResourceState, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareFeatureResourceStateNewStyle, EmptyObject: EmptyFeatureResourceState, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourceState")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Spec, actual.Spec, dcl.DiffInfo{ObjectFunction: compareFeatureSpecNewStyle, EmptyObject: EmptyFeatureSpec, OperationSelector: dcl.TriggersOperation("updateFeatureUpdateFeatureOperation")}, fn.AddNest("Spec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareFeatureStateNewStyle, EmptyObject: EmptyFeatureState, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.DeleteTime, actual.DeleteTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DeleteTime")); len(ds) != 0 || err != nil {
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
func compareFeatureResourceStateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureResourceState)
	if !ok {
		desiredNotPointer, ok := d.(FeatureResourceState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureResourceState or *FeatureResourceState", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureResourceState)
	if !ok {
		actualNotPointer, ok := a.(FeatureResourceState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureResourceState", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HasResources, actual.HasResources, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HasResources")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureSpecNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureSpec)
	if !ok {
		desiredNotPointer, ok := d.(FeatureSpec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureSpec or *FeatureSpec", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureSpec)
	if !ok {
		actualNotPointer, ok := a.(FeatureSpec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureSpec", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Multiclusteringress, actual.Multiclusteringress, dcl.DiffInfo{ObjectFunction: compareFeatureSpecMulticlusteringressNewStyle, EmptyObject: EmptyFeatureSpecMulticlusteringress, OperationSelector: dcl.TriggersOperation("updateFeatureUpdateFeatureOperation")}, fn.AddNest("Multiclusteringress")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Fleetobservability, actual.Fleetobservability, dcl.DiffInfo{ObjectFunction: compareFeatureSpecFleetobservabilityNewStyle, EmptyObject: EmptyFeatureSpecFleetobservability, OperationSelector: dcl.TriggersOperation("updateFeatureUpdateFeatureOperation")}, fn.AddNest("Fleetobservability")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureSpecMulticlusteringressNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureSpecMulticlusteringress)
	if !ok {
		desiredNotPointer, ok := d.(FeatureSpecMulticlusteringress)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureSpecMulticlusteringress or *FeatureSpecMulticlusteringress", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureSpecMulticlusteringress)
	if !ok {
		actualNotPointer, ok := a.(FeatureSpecMulticlusteringress)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureSpecMulticlusteringress", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ConfigMembership, actual.ConfigMembership, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateFeatureUpdateFeatureOperation")}, fn.AddNest("ConfigMembership")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureSpecFleetobservabilityNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureSpecFleetobservability)
	if !ok {
		desiredNotPointer, ok := d.(FeatureSpecFleetobservability)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureSpecFleetobservability or *FeatureSpecFleetobservability", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureSpecFleetobservability)
	if !ok {
		actualNotPointer, ok := a.(FeatureSpecFleetobservability)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureSpecFleetobservability", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.LoggingConfig, actual.LoggingConfig, dcl.DiffInfo{ObjectFunction: compareFeatureSpecFleetobservabilityLoggingConfigNewStyle, EmptyObject: EmptyFeatureSpecFleetobservabilityLoggingConfig, OperationSelector: dcl.TriggersOperation("updateFeatureUpdateFeatureOperation")}, fn.AddNest("LoggingConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureSpecFleetobservabilityLoggingConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureSpecFleetobservabilityLoggingConfig)
	if !ok {
		desiredNotPointer, ok := d.(FeatureSpecFleetobservabilityLoggingConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureSpecFleetobservabilityLoggingConfig or *FeatureSpecFleetobservabilityLoggingConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureSpecFleetobservabilityLoggingConfig)
	if !ok {
		actualNotPointer, ok := a.(FeatureSpecFleetobservabilityLoggingConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureSpecFleetobservabilityLoggingConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DefaultConfig, actual.DefaultConfig, dcl.DiffInfo{ObjectFunction: compareFeatureSpecFleetobservabilityLoggingConfigDefaultConfigNewStyle, EmptyObject: EmptyFeatureSpecFleetobservabilityLoggingConfigDefaultConfig, OperationSelector: dcl.TriggersOperation("updateFeatureUpdateFeatureOperation")}, fn.AddNest("DefaultConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FleetScopeLogsConfig, actual.FleetScopeLogsConfig, dcl.DiffInfo{ObjectFunction: compareFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigNewStyle, EmptyObject: EmptyFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig, OperationSelector: dcl.TriggersOperation("updateFeatureUpdateFeatureOperation")}, fn.AddNest("FleetScopeLogsConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureSpecFleetobservabilityLoggingConfigDefaultConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureSpecFleetobservabilityLoggingConfigDefaultConfig)
	if !ok {
		desiredNotPointer, ok := d.(FeatureSpecFleetobservabilityLoggingConfigDefaultConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureSpecFleetobservabilityLoggingConfigDefaultConfig or *FeatureSpecFleetobservabilityLoggingConfigDefaultConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureSpecFleetobservabilityLoggingConfigDefaultConfig)
	if !ok {
		actualNotPointer, ok := a.(FeatureSpecFleetobservabilityLoggingConfigDefaultConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureSpecFleetobservabilityLoggingConfigDefaultConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Mode, actual.Mode, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateFeatureUpdateFeatureOperation")}, fn.AddNest("Mode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig)
	if !ok {
		desiredNotPointer, ok := d.(FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig or *FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig)
	if !ok {
		actualNotPointer, ok := a.(FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Mode, actual.Mode, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateFeatureUpdateFeatureOperation")}, fn.AddNest("Mode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureStateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureState)
	if !ok {
		desiredNotPointer, ok := d.(FeatureState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureState or *FeatureState", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureState)
	if !ok {
		actualNotPointer, ok := a.(FeatureState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureState", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareFeatureStateStateNewStyle, EmptyObject: EmptyFeatureStateState, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureStateStateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureStateState)
	if !ok {
		desiredNotPointer, ok := d.(FeatureStateState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureStateState or *FeatureStateState", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureStateState)
	if !ok {
		actualNotPointer, ok := a.(FeatureStateState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureStateState", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Code, actual.Code, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Code")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
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
func (r *Feature) urlNormalized() *Feature {
	normalized := dcl.Copy(*r).(Feature)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Feature) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateFeature" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/features/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Feature resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Feature) marshal(c *Client) ([]byte, error) {
	m, err := expandFeature(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Feature: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalFeature decodes JSON responses into the Feature resource schema.
func unmarshalFeature(b []byte, c *Client, res *Feature) (*Feature, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapFeature(m, c, res)
}

func unmarshalMapFeature(m map[string]interface{}, c *Client, res *Feature) (*Feature, error) {

	flattened := flattenFeature(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandFeature expands Feature into a JSON request object.
func expandFeature(c *Client, f *Feature) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("projects/%s/locations/%s/features/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Location), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Labels; dcl.ValueShouldBeSent(v) {
		m["labels"] = v
	}
	if v, err := expandFeatureSpec(c, f.Spec, res); err != nil {
		return nil, fmt.Errorf("error expanding Spec into spec: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["spec"] = v
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

// flattenFeature flattens Feature from a JSON request object into the
// Feature type.
func flattenFeature(c *Client, i interface{}, res *Feature) *Feature {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Feature{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	resultRes.ResourceState = flattenFeatureResourceState(c, m["resourceState"], res)
	resultRes.Spec = flattenFeatureSpec(c, m["spec"], res)
	resultRes.State = flattenFeatureState(c, m["state"], res)
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.DeleteTime = dcl.FlattenString(m["deleteTime"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// expandFeatureResourceStateMap expands the contents of FeatureResourceState into a JSON
// request object.
func expandFeatureResourceStateMap(c *Client, f map[string]FeatureResourceState, res *Feature) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureResourceState(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureResourceStateSlice expands the contents of FeatureResourceState into a JSON
// request object.
func expandFeatureResourceStateSlice(c *Client, f []FeatureResourceState, res *Feature) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureResourceState(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureResourceStateMap flattens the contents of FeatureResourceState from a JSON
// response object.
func flattenFeatureResourceStateMap(c *Client, i interface{}, res *Feature) map[string]FeatureResourceState {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureResourceState{}
	}

	if len(a) == 0 {
		return map[string]FeatureResourceState{}
	}

	items := make(map[string]FeatureResourceState)
	for k, item := range a {
		items[k] = *flattenFeatureResourceState(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureResourceStateSlice flattens the contents of FeatureResourceState from a JSON
// response object.
func flattenFeatureResourceStateSlice(c *Client, i interface{}, res *Feature) []FeatureResourceState {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureResourceState{}
	}

	if len(a) == 0 {
		return []FeatureResourceState{}
	}

	items := make([]FeatureResourceState, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureResourceState(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureResourceState expands an instance of FeatureResourceState into a JSON
// request object.
func expandFeatureResourceState(c *Client, f *FeatureResourceState, res *Feature) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenFeatureResourceState flattens an instance of FeatureResourceState from a JSON
// response object.
func flattenFeatureResourceState(c *Client, i interface{}, res *Feature) *FeatureResourceState {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureResourceState{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureResourceState
	}
	r.State = flattenFeatureResourceStateStateEnum(m["state"])
	r.HasResources = dcl.FlattenBool(m["hasResources"])

	return r
}

// expandFeatureSpecMap expands the contents of FeatureSpec into a JSON
// request object.
func expandFeatureSpecMap(c *Client, f map[string]FeatureSpec, res *Feature) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureSpec(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureSpecSlice expands the contents of FeatureSpec into a JSON
// request object.
func expandFeatureSpecSlice(c *Client, f []FeatureSpec, res *Feature) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureSpec(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureSpecMap flattens the contents of FeatureSpec from a JSON
// response object.
func flattenFeatureSpecMap(c *Client, i interface{}, res *Feature) map[string]FeatureSpec {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureSpec{}
	}

	if len(a) == 0 {
		return map[string]FeatureSpec{}
	}

	items := make(map[string]FeatureSpec)
	for k, item := range a {
		items[k] = *flattenFeatureSpec(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureSpecSlice flattens the contents of FeatureSpec from a JSON
// response object.
func flattenFeatureSpecSlice(c *Client, i interface{}, res *Feature) []FeatureSpec {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureSpec{}
	}

	if len(a) == 0 {
		return []FeatureSpec{}
	}

	items := make([]FeatureSpec, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureSpec(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureSpec expands an instance of FeatureSpec into a JSON
// request object.
func expandFeatureSpec(c *Client, f *FeatureSpec, res *Feature) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandFeatureSpecMulticlusteringress(c, f.Multiclusteringress, res); err != nil {
		return nil, fmt.Errorf("error expanding Multiclusteringress into multiclusteringress: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["multiclusteringress"] = v
	}
	if v, err := expandFeatureSpecFleetobservability(c, f.Fleetobservability, res); err != nil {
		return nil, fmt.Errorf("error expanding Fleetobservability into fleetobservability: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["fleetobservability"] = v
	}

	return m, nil
}

// flattenFeatureSpec flattens an instance of FeatureSpec from a JSON
// response object.
func flattenFeatureSpec(c *Client, i interface{}, res *Feature) *FeatureSpec {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureSpec{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureSpec
	}
	r.Multiclusteringress = flattenFeatureSpecMulticlusteringress(c, m["multiclusteringress"], res)
	r.Fleetobservability = flattenFeatureSpecFleetobservability(c, m["fleetobservability"], res)

	return r
}

// expandFeatureSpecMulticlusteringressMap expands the contents of FeatureSpecMulticlusteringress into a JSON
// request object.
func expandFeatureSpecMulticlusteringressMap(c *Client, f map[string]FeatureSpecMulticlusteringress, res *Feature) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureSpecMulticlusteringress(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureSpecMulticlusteringressSlice expands the contents of FeatureSpecMulticlusteringress into a JSON
// request object.
func expandFeatureSpecMulticlusteringressSlice(c *Client, f []FeatureSpecMulticlusteringress, res *Feature) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureSpecMulticlusteringress(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureSpecMulticlusteringressMap flattens the contents of FeatureSpecMulticlusteringress from a JSON
// response object.
func flattenFeatureSpecMulticlusteringressMap(c *Client, i interface{}, res *Feature) map[string]FeatureSpecMulticlusteringress {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureSpecMulticlusteringress{}
	}

	if len(a) == 0 {
		return map[string]FeatureSpecMulticlusteringress{}
	}

	items := make(map[string]FeatureSpecMulticlusteringress)
	for k, item := range a {
		items[k] = *flattenFeatureSpecMulticlusteringress(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureSpecMulticlusteringressSlice flattens the contents of FeatureSpecMulticlusteringress from a JSON
// response object.
func flattenFeatureSpecMulticlusteringressSlice(c *Client, i interface{}, res *Feature) []FeatureSpecMulticlusteringress {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureSpecMulticlusteringress{}
	}

	if len(a) == 0 {
		return []FeatureSpecMulticlusteringress{}
	}

	items := make([]FeatureSpecMulticlusteringress, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureSpecMulticlusteringress(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureSpecMulticlusteringress expands an instance of FeatureSpecMulticlusteringress into a JSON
// request object.
func expandFeatureSpecMulticlusteringress(c *Client, f *FeatureSpecMulticlusteringress, res *Feature) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ConfigMembership; !dcl.IsEmptyValueIndirect(v) {
		m["configMembership"] = v
	}

	return m, nil
}

// flattenFeatureSpecMulticlusteringress flattens an instance of FeatureSpecMulticlusteringress from a JSON
// response object.
func flattenFeatureSpecMulticlusteringress(c *Client, i interface{}, res *Feature) *FeatureSpecMulticlusteringress {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureSpecMulticlusteringress{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureSpecMulticlusteringress
	}
	r.ConfigMembership = dcl.FlattenString(m["configMembership"])

	return r
}

// expandFeatureSpecFleetobservabilityMap expands the contents of FeatureSpecFleetobservability into a JSON
// request object.
func expandFeatureSpecFleetobservabilityMap(c *Client, f map[string]FeatureSpecFleetobservability, res *Feature) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureSpecFleetobservability(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureSpecFleetobservabilitySlice expands the contents of FeatureSpecFleetobservability into a JSON
// request object.
func expandFeatureSpecFleetobservabilitySlice(c *Client, f []FeatureSpecFleetobservability, res *Feature) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureSpecFleetobservability(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureSpecFleetobservabilityMap flattens the contents of FeatureSpecFleetobservability from a JSON
// response object.
func flattenFeatureSpecFleetobservabilityMap(c *Client, i interface{}, res *Feature) map[string]FeatureSpecFleetobservability {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureSpecFleetobservability{}
	}

	if len(a) == 0 {
		return map[string]FeatureSpecFleetobservability{}
	}

	items := make(map[string]FeatureSpecFleetobservability)
	for k, item := range a {
		items[k] = *flattenFeatureSpecFleetobservability(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureSpecFleetobservabilitySlice flattens the contents of FeatureSpecFleetobservability from a JSON
// response object.
func flattenFeatureSpecFleetobservabilitySlice(c *Client, i interface{}, res *Feature) []FeatureSpecFleetobservability {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureSpecFleetobservability{}
	}

	if len(a) == 0 {
		return []FeatureSpecFleetobservability{}
	}

	items := make([]FeatureSpecFleetobservability, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureSpecFleetobservability(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureSpecFleetobservability expands an instance of FeatureSpecFleetobservability into a JSON
// request object.
func expandFeatureSpecFleetobservability(c *Client, f *FeatureSpecFleetobservability, res *Feature) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandFeatureSpecFleetobservabilityLoggingConfig(c, f.LoggingConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding LoggingConfig into loggingConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["loggingConfig"] = v
	}

	return m, nil
}

// flattenFeatureSpecFleetobservability flattens an instance of FeatureSpecFleetobservability from a JSON
// response object.
func flattenFeatureSpecFleetobservability(c *Client, i interface{}, res *Feature) *FeatureSpecFleetobservability {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureSpecFleetobservability{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureSpecFleetobservability
	}
	r.LoggingConfig = flattenFeatureSpecFleetobservabilityLoggingConfig(c, m["loggingConfig"], res)

	return r
}

// expandFeatureSpecFleetobservabilityLoggingConfigMap expands the contents of FeatureSpecFleetobservabilityLoggingConfig into a JSON
// request object.
func expandFeatureSpecFleetobservabilityLoggingConfigMap(c *Client, f map[string]FeatureSpecFleetobservabilityLoggingConfig, res *Feature) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureSpecFleetobservabilityLoggingConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureSpecFleetobservabilityLoggingConfigSlice expands the contents of FeatureSpecFleetobservabilityLoggingConfig into a JSON
// request object.
func expandFeatureSpecFleetobservabilityLoggingConfigSlice(c *Client, f []FeatureSpecFleetobservabilityLoggingConfig, res *Feature) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureSpecFleetobservabilityLoggingConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureSpecFleetobservabilityLoggingConfigMap flattens the contents of FeatureSpecFleetobservabilityLoggingConfig from a JSON
// response object.
func flattenFeatureSpecFleetobservabilityLoggingConfigMap(c *Client, i interface{}, res *Feature) map[string]FeatureSpecFleetobservabilityLoggingConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureSpecFleetobservabilityLoggingConfig{}
	}

	if len(a) == 0 {
		return map[string]FeatureSpecFleetobservabilityLoggingConfig{}
	}

	items := make(map[string]FeatureSpecFleetobservabilityLoggingConfig)
	for k, item := range a {
		items[k] = *flattenFeatureSpecFleetobservabilityLoggingConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureSpecFleetobservabilityLoggingConfigSlice flattens the contents of FeatureSpecFleetobservabilityLoggingConfig from a JSON
// response object.
func flattenFeatureSpecFleetobservabilityLoggingConfigSlice(c *Client, i interface{}, res *Feature) []FeatureSpecFleetobservabilityLoggingConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureSpecFleetobservabilityLoggingConfig{}
	}

	if len(a) == 0 {
		return []FeatureSpecFleetobservabilityLoggingConfig{}
	}

	items := make([]FeatureSpecFleetobservabilityLoggingConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureSpecFleetobservabilityLoggingConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureSpecFleetobservabilityLoggingConfig expands an instance of FeatureSpecFleetobservabilityLoggingConfig into a JSON
// request object.
func expandFeatureSpecFleetobservabilityLoggingConfig(c *Client, f *FeatureSpecFleetobservabilityLoggingConfig, res *Feature) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandFeatureSpecFleetobservabilityLoggingConfigDefaultConfig(c, f.DefaultConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding DefaultConfig into defaultConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["defaultConfig"] = v
	}
	if v, err := expandFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(c, f.FleetScopeLogsConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding FleetScopeLogsConfig into fleetScopeLogsConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["fleetScopeLogsConfig"] = v
	}

	return m, nil
}

// flattenFeatureSpecFleetobservabilityLoggingConfig flattens an instance of FeatureSpecFleetobservabilityLoggingConfig from a JSON
// response object.
func flattenFeatureSpecFleetobservabilityLoggingConfig(c *Client, i interface{}, res *Feature) *FeatureSpecFleetobservabilityLoggingConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureSpecFleetobservabilityLoggingConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureSpecFleetobservabilityLoggingConfig
	}
	r.DefaultConfig = flattenFeatureSpecFleetobservabilityLoggingConfigDefaultConfig(c, m["defaultConfig"], res)
	r.FleetScopeLogsConfig = flattenFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(c, m["fleetScopeLogsConfig"], res)

	return r
}

// expandFeatureSpecFleetobservabilityLoggingConfigDefaultConfigMap expands the contents of FeatureSpecFleetobservabilityLoggingConfigDefaultConfig into a JSON
// request object.
func expandFeatureSpecFleetobservabilityLoggingConfigDefaultConfigMap(c *Client, f map[string]FeatureSpecFleetobservabilityLoggingConfigDefaultConfig, res *Feature) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureSpecFleetobservabilityLoggingConfigDefaultConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureSpecFleetobservabilityLoggingConfigDefaultConfigSlice expands the contents of FeatureSpecFleetobservabilityLoggingConfigDefaultConfig into a JSON
// request object.
func expandFeatureSpecFleetobservabilityLoggingConfigDefaultConfigSlice(c *Client, f []FeatureSpecFleetobservabilityLoggingConfigDefaultConfig, res *Feature) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureSpecFleetobservabilityLoggingConfigDefaultConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureSpecFleetobservabilityLoggingConfigDefaultConfigMap flattens the contents of FeatureSpecFleetobservabilityLoggingConfigDefaultConfig from a JSON
// response object.
func flattenFeatureSpecFleetobservabilityLoggingConfigDefaultConfigMap(c *Client, i interface{}, res *Feature) map[string]FeatureSpecFleetobservabilityLoggingConfigDefaultConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureSpecFleetobservabilityLoggingConfigDefaultConfig{}
	}

	if len(a) == 0 {
		return map[string]FeatureSpecFleetobservabilityLoggingConfigDefaultConfig{}
	}

	items := make(map[string]FeatureSpecFleetobservabilityLoggingConfigDefaultConfig)
	for k, item := range a {
		items[k] = *flattenFeatureSpecFleetobservabilityLoggingConfigDefaultConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureSpecFleetobservabilityLoggingConfigDefaultConfigSlice flattens the contents of FeatureSpecFleetobservabilityLoggingConfigDefaultConfig from a JSON
// response object.
func flattenFeatureSpecFleetobservabilityLoggingConfigDefaultConfigSlice(c *Client, i interface{}, res *Feature) []FeatureSpecFleetobservabilityLoggingConfigDefaultConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureSpecFleetobservabilityLoggingConfigDefaultConfig{}
	}

	if len(a) == 0 {
		return []FeatureSpecFleetobservabilityLoggingConfigDefaultConfig{}
	}

	items := make([]FeatureSpecFleetobservabilityLoggingConfigDefaultConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureSpecFleetobservabilityLoggingConfigDefaultConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureSpecFleetobservabilityLoggingConfigDefaultConfig expands an instance of FeatureSpecFleetobservabilityLoggingConfigDefaultConfig into a JSON
// request object.
func expandFeatureSpecFleetobservabilityLoggingConfigDefaultConfig(c *Client, f *FeatureSpecFleetobservabilityLoggingConfigDefaultConfig, res *Feature) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Mode; !dcl.IsEmptyValueIndirect(v) {
		m["mode"] = v
	}

	return m, nil
}

// flattenFeatureSpecFleetobservabilityLoggingConfigDefaultConfig flattens an instance of FeatureSpecFleetobservabilityLoggingConfigDefaultConfig from a JSON
// response object.
func flattenFeatureSpecFleetobservabilityLoggingConfigDefaultConfig(c *Client, i interface{}, res *Feature) *FeatureSpecFleetobservabilityLoggingConfigDefaultConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureSpecFleetobservabilityLoggingConfigDefaultConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureSpecFleetobservabilityLoggingConfigDefaultConfig
	}
	r.Mode = flattenFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum(m["mode"])

	return r
}

// expandFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigMap expands the contents of FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig into a JSON
// request object.
func expandFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigMap(c *Client, f map[string]FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig, res *Feature) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigSlice expands the contents of FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig into a JSON
// request object.
func expandFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigSlice(c *Client, f []FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig, res *Feature) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigMap flattens the contents of FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig from a JSON
// response object.
func flattenFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigMap(c *Client, i interface{}, res *Feature) map[string]FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig{}
	}

	if len(a) == 0 {
		return map[string]FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig{}
	}

	items := make(map[string]FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig)
	for k, item := range a {
		items[k] = *flattenFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigSlice flattens the contents of FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig from a JSON
// response object.
func flattenFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigSlice(c *Client, i interface{}, res *Feature) []FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig{}
	}

	if len(a) == 0 {
		return []FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig{}
	}

	items := make([]FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig expands an instance of FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig into a JSON
// request object.
func expandFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(c *Client, f *FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig, res *Feature) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Mode; !dcl.IsEmptyValueIndirect(v) {
		m["mode"] = v
	}

	return m, nil
}

// flattenFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig flattens an instance of FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig from a JSON
// response object.
func flattenFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(c *Client, i interface{}, res *Feature) *FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig
	}
	r.Mode = flattenFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum(m["mode"])

	return r
}

// expandFeatureStateMap expands the contents of FeatureState into a JSON
// request object.
func expandFeatureStateMap(c *Client, f map[string]FeatureState, res *Feature) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureState(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureStateSlice expands the contents of FeatureState into a JSON
// request object.
func expandFeatureStateSlice(c *Client, f []FeatureState, res *Feature) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureState(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureStateMap flattens the contents of FeatureState from a JSON
// response object.
func flattenFeatureStateMap(c *Client, i interface{}, res *Feature) map[string]FeatureState {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureState{}
	}

	if len(a) == 0 {
		return map[string]FeatureState{}
	}

	items := make(map[string]FeatureState)
	for k, item := range a {
		items[k] = *flattenFeatureState(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureStateSlice flattens the contents of FeatureState from a JSON
// response object.
func flattenFeatureStateSlice(c *Client, i interface{}, res *Feature) []FeatureState {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureState{}
	}

	if len(a) == 0 {
		return []FeatureState{}
	}

	items := make([]FeatureState, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureState(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureState expands an instance of FeatureState into a JSON
// request object.
func expandFeatureState(c *Client, f *FeatureState, res *Feature) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenFeatureState flattens an instance of FeatureState from a JSON
// response object.
func flattenFeatureState(c *Client, i interface{}, res *Feature) *FeatureState {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureState{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureState
	}
	r.State = flattenFeatureStateState(c, m["state"], res)

	return r
}

// expandFeatureStateStateMap expands the contents of FeatureStateState into a JSON
// request object.
func expandFeatureStateStateMap(c *Client, f map[string]FeatureStateState, res *Feature) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureStateState(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureStateStateSlice expands the contents of FeatureStateState into a JSON
// request object.
func expandFeatureStateStateSlice(c *Client, f []FeatureStateState, res *Feature) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureStateState(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureStateStateMap flattens the contents of FeatureStateState from a JSON
// response object.
func flattenFeatureStateStateMap(c *Client, i interface{}, res *Feature) map[string]FeatureStateState {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureStateState{}
	}

	if len(a) == 0 {
		return map[string]FeatureStateState{}
	}

	items := make(map[string]FeatureStateState)
	for k, item := range a {
		items[k] = *flattenFeatureStateState(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureStateStateSlice flattens the contents of FeatureStateState from a JSON
// response object.
func flattenFeatureStateStateSlice(c *Client, i interface{}, res *Feature) []FeatureStateState {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureStateState{}
	}

	if len(a) == 0 {
		return []FeatureStateState{}
	}

	items := make([]FeatureStateState, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureStateState(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureStateState expands an instance of FeatureStateState into a JSON
// request object.
func expandFeatureStateState(c *Client, f *FeatureStateState, res *Feature) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenFeatureStateState flattens an instance of FeatureStateState from a JSON
// response object.
func flattenFeatureStateState(c *Client, i interface{}, res *Feature) *FeatureStateState {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureStateState{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureStateState
	}
	r.Code = flattenFeatureStateStateCodeEnum(m["code"])
	r.Description = dcl.FlattenString(m["description"])
	r.UpdateTime = dcl.FlattenString(m["updateTime"])

	return r
}

// flattenFeatureResourceStateStateEnumMap flattens the contents of FeatureResourceStateStateEnum from a JSON
// response object.
func flattenFeatureResourceStateStateEnumMap(c *Client, i interface{}, res *Feature) map[string]FeatureResourceStateStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureResourceStateStateEnum{}
	}

	if len(a) == 0 {
		return map[string]FeatureResourceStateStateEnum{}
	}

	items := make(map[string]FeatureResourceStateStateEnum)
	for k, item := range a {
		items[k] = *flattenFeatureResourceStateStateEnum(item.(interface{}))
	}

	return items
}

// flattenFeatureResourceStateStateEnumSlice flattens the contents of FeatureResourceStateStateEnum from a JSON
// response object.
func flattenFeatureResourceStateStateEnumSlice(c *Client, i interface{}, res *Feature) []FeatureResourceStateStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureResourceStateStateEnum{}
	}

	if len(a) == 0 {
		return []FeatureResourceStateStateEnum{}
	}

	items := make([]FeatureResourceStateStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureResourceStateStateEnum(item.(interface{})))
	}

	return items
}

// flattenFeatureResourceStateStateEnum asserts that an interface is a string, and returns a
// pointer to a *FeatureResourceStateStateEnum with the same value as that string.
func flattenFeatureResourceStateStateEnum(i interface{}) *FeatureResourceStateStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return FeatureResourceStateStateEnumRef(s)
}

// flattenFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnumMap flattens the contents of FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum from a JSON
// response object.
func flattenFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnumMap(c *Client, i interface{}, res *Feature) map[string]FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum{}
	}

	if len(a) == 0 {
		return map[string]FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum{}
	}

	items := make(map[string]FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum)
	for k, item := range a {
		items[k] = *flattenFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum(item.(interface{}))
	}

	return items
}

// flattenFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnumSlice flattens the contents of FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum from a JSON
// response object.
func flattenFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnumSlice(c *Client, i interface{}, res *Feature) []FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum{}
	}

	if len(a) == 0 {
		return []FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum{}
	}

	items := make([]FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum(item.(interface{})))
	}

	return items
}

// flattenFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum asserts that an interface is a string, and returns a
// pointer to a *FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum with the same value as that string.
func flattenFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum(i interface{}) *FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnumRef(s)
}

// flattenFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnumMap flattens the contents of FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum from a JSON
// response object.
func flattenFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnumMap(c *Client, i interface{}, res *Feature) map[string]FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum{}
	}

	if len(a) == 0 {
		return map[string]FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum{}
	}

	items := make(map[string]FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum)
	for k, item := range a {
		items[k] = *flattenFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum(item.(interface{}))
	}

	return items
}

// flattenFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnumSlice flattens the contents of FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum from a JSON
// response object.
func flattenFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnumSlice(c *Client, i interface{}, res *Feature) []FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum{}
	}

	if len(a) == 0 {
		return []FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum{}
	}

	items := make([]FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum(item.(interface{})))
	}

	return items
}

// flattenFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum asserts that an interface is a string, and returns a
// pointer to a *FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum with the same value as that string.
func flattenFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum(i interface{}) *FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnumRef(s)
}

// flattenFeatureStateStateCodeEnumMap flattens the contents of FeatureStateStateCodeEnum from a JSON
// response object.
func flattenFeatureStateStateCodeEnumMap(c *Client, i interface{}, res *Feature) map[string]FeatureStateStateCodeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureStateStateCodeEnum{}
	}

	if len(a) == 0 {
		return map[string]FeatureStateStateCodeEnum{}
	}

	items := make(map[string]FeatureStateStateCodeEnum)
	for k, item := range a {
		items[k] = *flattenFeatureStateStateCodeEnum(item.(interface{}))
	}

	return items
}

// flattenFeatureStateStateCodeEnumSlice flattens the contents of FeatureStateStateCodeEnum from a JSON
// response object.
func flattenFeatureStateStateCodeEnumSlice(c *Client, i interface{}, res *Feature) []FeatureStateStateCodeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureStateStateCodeEnum{}
	}

	if len(a) == 0 {
		return []FeatureStateStateCodeEnum{}
	}

	items := make([]FeatureStateStateCodeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureStateStateCodeEnum(item.(interface{})))
	}

	return items
}

// flattenFeatureStateStateCodeEnum asserts that an interface is a string, and returns a
// pointer to a *FeatureStateStateCodeEnum with the same value as that string.
func flattenFeatureStateStateCodeEnum(i interface{}) *FeatureStateStateCodeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return FeatureStateStateCodeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Feature) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalFeature(b, c, r)
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

type featureDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         featureApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToFeatureDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]featureDiff, error) {
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
	var diffs []featureDiff
	// For each operation name, create a featureDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := featureDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToFeatureApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToFeatureApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (featureApiOperation, error) {
	switch opName {

	case "updateFeatureUpdateFeatureOperation":
		return &updateFeatureUpdateFeatureOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractFeatureFields(r *Feature) error {
	vResourceState := r.ResourceState
	if vResourceState == nil {
		// note: explicitly not the empty object.
		vResourceState = &FeatureResourceState{}
	}
	if err := extractFeatureResourceStateFields(r, vResourceState); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vResourceState) {
		r.ResourceState = vResourceState
	}
	vSpec := r.Spec
	if vSpec == nil {
		// note: explicitly not the empty object.
		vSpec = &FeatureSpec{}
	}
	if err := extractFeatureSpecFields(r, vSpec); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSpec) {
		r.Spec = vSpec
	}
	vState := r.State
	if vState == nil {
		// note: explicitly not the empty object.
		vState = &FeatureState{}
	}
	if err := extractFeatureStateFields(r, vState); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vState) {
		r.State = vState
	}
	return nil
}
func extractFeatureResourceStateFields(r *Feature, o *FeatureResourceState) error {
	return nil
}
func extractFeatureSpecFields(r *Feature, o *FeatureSpec) error {
	vMulticlusteringress := o.Multiclusteringress
	if vMulticlusteringress == nil {
		// note: explicitly not the empty object.
		vMulticlusteringress = &FeatureSpecMulticlusteringress{}
	}
	if err := extractFeatureSpecMulticlusteringressFields(r, vMulticlusteringress); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMulticlusteringress) {
		o.Multiclusteringress = vMulticlusteringress
	}
	vFleetobservability := o.Fleetobservability
	if vFleetobservability == nil {
		// note: explicitly not the empty object.
		vFleetobservability = &FeatureSpecFleetobservability{}
	}
	if err := extractFeatureSpecFleetobservabilityFields(r, vFleetobservability); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vFleetobservability) {
		o.Fleetobservability = vFleetobservability
	}
	return nil
}
func extractFeatureSpecMulticlusteringressFields(r *Feature, o *FeatureSpecMulticlusteringress) error {
	return nil
}
func extractFeatureSpecFleetobservabilityFields(r *Feature, o *FeatureSpecFleetobservability) error {
	vLoggingConfig := o.LoggingConfig
	if vLoggingConfig == nil {
		// note: explicitly not the empty object.
		vLoggingConfig = &FeatureSpecFleetobservabilityLoggingConfig{}
	}
	if err := extractFeatureSpecFleetobservabilityLoggingConfigFields(r, vLoggingConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLoggingConfig) {
		o.LoggingConfig = vLoggingConfig
	}
	return nil
}
func extractFeatureSpecFleetobservabilityLoggingConfigFields(r *Feature, o *FeatureSpecFleetobservabilityLoggingConfig) error {
	vDefaultConfig := o.DefaultConfig
	if vDefaultConfig == nil {
		// note: explicitly not the empty object.
		vDefaultConfig = &FeatureSpecFleetobservabilityLoggingConfigDefaultConfig{}
	}
	if err := extractFeatureSpecFleetobservabilityLoggingConfigDefaultConfigFields(r, vDefaultConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDefaultConfig) {
		o.DefaultConfig = vDefaultConfig
	}
	vFleetScopeLogsConfig := o.FleetScopeLogsConfig
	if vFleetScopeLogsConfig == nil {
		// note: explicitly not the empty object.
		vFleetScopeLogsConfig = &FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig{}
	}
	if err := extractFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigFields(r, vFleetScopeLogsConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vFleetScopeLogsConfig) {
		o.FleetScopeLogsConfig = vFleetScopeLogsConfig
	}
	return nil
}
func extractFeatureSpecFleetobservabilityLoggingConfigDefaultConfigFields(r *Feature, o *FeatureSpecFleetobservabilityLoggingConfigDefaultConfig) error {
	return nil
}
func extractFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigFields(r *Feature, o *FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig) error {
	return nil
}
func extractFeatureStateFields(r *Feature, o *FeatureState) error {
	vState := o.State
	if vState == nil {
		// note: explicitly not the empty object.
		vState = &FeatureStateState{}
	}
	if err := extractFeatureStateStateFields(r, vState); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vState) {
		o.State = vState
	}
	return nil
}
func extractFeatureStateStateFields(r *Feature, o *FeatureStateState) error {
	return nil
}

func postReadExtractFeatureFields(r *Feature) error {
	vResourceState := r.ResourceState
	if vResourceState == nil {
		// note: explicitly not the empty object.
		vResourceState = &FeatureResourceState{}
	}
	if err := postReadExtractFeatureResourceStateFields(r, vResourceState); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vResourceState) {
		r.ResourceState = vResourceState
	}
	vSpec := r.Spec
	if vSpec == nil {
		// note: explicitly not the empty object.
		vSpec = &FeatureSpec{}
	}
	if err := postReadExtractFeatureSpecFields(r, vSpec); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSpec) {
		r.Spec = vSpec
	}
	vState := r.State
	if vState == nil {
		// note: explicitly not the empty object.
		vState = &FeatureState{}
	}
	if err := postReadExtractFeatureStateFields(r, vState); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vState) {
		r.State = vState
	}
	return nil
}
func postReadExtractFeatureResourceStateFields(r *Feature, o *FeatureResourceState) error {
	return nil
}
func postReadExtractFeatureSpecFields(r *Feature, o *FeatureSpec) error {
	vMulticlusteringress := o.Multiclusteringress
	if vMulticlusteringress == nil {
		// note: explicitly not the empty object.
		vMulticlusteringress = &FeatureSpecMulticlusteringress{}
	}
	if err := extractFeatureSpecMulticlusteringressFields(r, vMulticlusteringress); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMulticlusteringress) {
		o.Multiclusteringress = vMulticlusteringress
	}
	vFleetobservability := o.Fleetobservability
	if vFleetobservability == nil {
		// note: explicitly not the empty object.
		vFleetobservability = &FeatureSpecFleetobservability{}
	}
	if err := extractFeatureSpecFleetobservabilityFields(r, vFleetobservability); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vFleetobservability) {
		o.Fleetobservability = vFleetobservability
	}
	return nil
}
func postReadExtractFeatureSpecMulticlusteringressFields(r *Feature, o *FeatureSpecMulticlusteringress) error {
	return nil
}
func postReadExtractFeatureSpecFleetobservabilityFields(r *Feature, o *FeatureSpecFleetobservability) error {
	vLoggingConfig := o.LoggingConfig
	if vLoggingConfig == nil {
		// note: explicitly not the empty object.
		vLoggingConfig = &FeatureSpecFleetobservabilityLoggingConfig{}
	}
	if err := extractFeatureSpecFleetobservabilityLoggingConfigFields(r, vLoggingConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLoggingConfig) {
		o.LoggingConfig = vLoggingConfig
	}
	return nil
}
func postReadExtractFeatureSpecFleetobservabilityLoggingConfigFields(r *Feature, o *FeatureSpecFleetobservabilityLoggingConfig) error {
	vDefaultConfig := o.DefaultConfig
	if vDefaultConfig == nil {
		// note: explicitly not the empty object.
		vDefaultConfig = &FeatureSpecFleetobservabilityLoggingConfigDefaultConfig{}
	}
	if err := extractFeatureSpecFleetobservabilityLoggingConfigDefaultConfigFields(r, vDefaultConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDefaultConfig) {
		o.DefaultConfig = vDefaultConfig
	}
	vFleetScopeLogsConfig := o.FleetScopeLogsConfig
	if vFleetScopeLogsConfig == nil {
		// note: explicitly not the empty object.
		vFleetScopeLogsConfig = &FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig{}
	}
	if err := extractFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigFields(r, vFleetScopeLogsConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vFleetScopeLogsConfig) {
		o.FleetScopeLogsConfig = vFleetScopeLogsConfig
	}
	return nil
}
func postReadExtractFeatureSpecFleetobservabilityLoggingConfigDefaultConfigFields(r *Feature, o *FeatureSpecFleetobservabilityLoggingConfigDefaultConfig) error {
	return nil
}
func postReadExtractFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigFields(r *Feature, o *FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig) error {
	return nil
}
func postReadExtractFeatureStateFields(r *Feature, o *FeatureState) error {
	vState := o.State
	if vState == nil {
		// note: explicitly not the empty object.
		vState = &FeatureStateState{}
	}
	if err := extractFeatureStateStateFields(r, vState); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vState) {
		o.State = vState
	}
	return nil
}
func postReadExtractFeatureStateStateFields(r *Feature, o *FeatureStateState) error {
	return nil
}
