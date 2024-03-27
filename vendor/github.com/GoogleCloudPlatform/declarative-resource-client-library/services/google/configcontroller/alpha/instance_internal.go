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

	if err := dcl.Required(r, "managementConfig"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.BundlesConfig) {
		if err := r.BundlesConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ManagementConfig) {
		if err := r.ManagementConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceBundlesConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.ConfigControllerConfig) {
		if err := r.ConfigControllerConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceBundlesConfigConfigControllerConfig) validate() error {
	return nil
}
func (r *InstanceManagementConfig) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"StandardManagementConfig", "FullManagementConfig"}, r.StandardManagementConfig, r.FullManagementConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.StandardManagementConfig) {
		if err := r.StandardManagementConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.FullManagementConfig) {
		if err := r.FullManagementConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceManagementConfigStandardManagementConfig) validate() error {
	if err := dcl.Required(r, "masterIPv4CidrBlock"); err != nil {
		return err
	}
	return nil
}
func (r *InstanceManagementConfigFullManagementConfig) validate() error {
	return nil
}
func (r *Instance) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://krmapihosting.googleapis.com/v1alpha1/", params)
}

func (r *Instance) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/krmApiHosts/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Instance) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/krmApiHosts", nr.basePath(), userBasePath, params), nil

}

func (r *Instance) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/krmApiHosts?krm_api_host_id={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *Instance) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/krmApiHosts/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Instance) SetPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{}
	return dcl.URL("", nr.basePath(), userBasePath, fields)
}

func (r *Instance) SetPolicyVerb() string {
	return ""
}

func (r *Instance) getPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{}
	return dcl.URL("", nr.basePath(), userBasePath, fields)
}

func (r *Instance) IAMPolicyVersion() int {
	return 3
}

// instanceApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type instanceApiOperation interface {
	do(context.Context, *Instance, *Client) error
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
	KrmApiHosts []map[string]interface{} `json:"krmApiHosts"`
	Token       string                   `json:"nextPageToken"`
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
	for _, v := range m.KrmApiHosts {
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

func (op *createInstanceOperation) do(ctx context.Context, r *Instance, c *Client) error {
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

	if _, err := c.GetInstance(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
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
		rawDesired.BundlesConfig = canonicalizeInstanceBundlesConfig(rawDesired.BundlesConfig, nil, opts...)
		rawDesired.ManagementConfig = canonicalizeInstanceManagementConfig(rawDesired.ManagementConfig, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Instance{}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
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
	canonicalDesired.BundlesConfig = canonicalizeInstanceBundlesConfig(rawDesired.BundlesConfig, rawInitial.BundlesConfig, opts...)
	if dcl.BoolCanonicalize(rawDesired.UsePrivateEndpoint, rawInitial.UsePrivateEndpoint) {
		canonicalDesired.UsePrivateEndpoint = rawInitial.UsePrivateEndpoint
	} else {
		canonicalDesired.UsePrivateEndpoint = rawDesired.UsePrivateEndpoint
	}
	canonicalDesired.ManagementConfig = canonicalizeInstanceManagementConfig(rawDesired.ManagementConfig, rawInitial.ManagementConfig, opts...)
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

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.BundlesConfig) && dcl.IsEmptyValueIndirect(rawDesired.BundlesConfig) {
		rawNew.BundlesConfig = rawDesired.BundlesConfig
	} else {
		rawNew.BundlesConfig = canonicalizeNewInstanceBundlesConfig(c, rawDesired.BundlesConfig, rawNew.BundlesConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.UsePrivateEndpoint) && dcl.IsEmptyValueIndirect(rawDesired.UsePrivateEndpoint) {
		rawNew.UsePrivateEndpoint = rawDesired.UsePrivateEndpoint
	} else {
		if dcl.BoolCanonicalize(rawDesired.UsePrivateEndpoint, rawNew.UsePrivateEndpoint) {
			rawNew.UsePrivateEndpoint = rawDesired.UsePrivateEndpoint
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.GkeResourceLink) && dcl.IsEmptyValueIndirect(rawDesired.GkeResourceLink) {
		rawNew.GkeResourceLink = rawDesired.GkeResourceLink
	} else {
		if dcl.StringCanonicalize(rawDesired.GkeResourceLink, rawNew.GkeResourceLink) {
			rawNew.GkeResourceLink = rawDesired.GkeResourceLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ManagementConfig) && dcl.IsEmptyValueIndirect(rawDesired.ManagementConfig) {
		rawNew.ManagementConfig = rawDesired.ManagementConfig
	} else {
		rawNew.ManagementConfig = canonicalizeNewInstanceManagementConfig(c, rawDesired.ManagementConfig, rawNew.ManagementConfig)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeInstanceBundlesConfig(des, initial *InstanceBundlesConfig, opts ...dcl.ApplyOption) *InstanceBundlesConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceBundlesConfig{}

	cDes.ConfigControllerConfig = canonicalizeInstanceBundlesConfigConfigControllerConfig(des.ConfigControllerConfig, initial.ConfigControllerConfig, opts...)

	return cDes
}

func canonicalizeInstanceBundlesConfigSlice(des, initial []InstanceBundlesConfig, opts ...dcl.ApplyOption) []InstanceBundlesConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InstanceBundlesConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInstanceBundlesConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InstanceBundlesConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInstanceBundlesConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInstanceBundlesConfig(c *Client, des, nw *InstanceBundlesConfig) *InstanceBundlesConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InstanceBundlesConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.ConfigControllerConfig = canonicalizeNewInstanceBundlesConfigConfigControllerConfig(c, des.ConfigControllerConfig, nw.ConfigControllerConfig)

	return nw
}

func canonicalizeNewInstanceBundlesConfigSet(c *Client, des, nw []InstanceBundlesConfig) []InstanceBundlesConfig {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InstanceBundlesConfig
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInstanceBundlesConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInstanceBundlesConfig(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInstanceBundlesConfigSlice(c *Client, des, nw []InstanceBundlesConfig) []InstanceBundlesConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceBundlesConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceBundlesConfig(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceBundlesConfigConfigControllerConfig(des, initial *InstanceBundlesConfigConfigControllerConfig, opts ...dcl.ApplyOption) *InstanceBundlesConfigConfigControllerConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceBundlesConfigConfigControllerConfig{}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		cDes.Enabled = initial.Enabled
	} else {
		cDes.Enabled = des.Enabled
	}

	return cDes
}

func canonicalizeInstanceBundlesConfigConfigControllerConfigSlice(des, initial []InstanceBundlesConfigConfigControllerConfig, opts ...dcl.ApplyOption) []InstanceBundlesConfigConfigControllerConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InstanceBundlesConfigConfigControllerConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInstanceBundlesConfigConfigControllerConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InstanceBundlesConfigConfigControllerConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInstanceBundlesConfigConfigControllerConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInstanceBundlesConfigConfigControllerConfig(c *Client, des, nw *InstanceBundlesConfigConfigControllerConfig) *InstanceBundlesConfigConfigControllerConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InstanceBundlesConfigConfigControllerConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}

	return nw
}

func canonicalizeNewInstanceBundlesConfigConfigControllerConfigSet(c *Client, des, nw []InstanceBundlesConfigConfigControllerConfig) []InstanceBundlesConfigConfigControllerConfig {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InstanceBundlesConfigConfigControllerConfig
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInstanceBundlesConfigConfigControllerConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInstanceBundlesConfigConfigControllerConfig(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInstanceBundlesConfigConfigControllerConfigSlice(c *Client, des, nw []InstanceBundlesConfigConfigControllerConfig) []InstanceBundlesConfigConfigControllerConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceBundlesConfigConfigControllerConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceBundlesConfigConfigControllerConfig(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceManagementConfig(des, initial *InstanceManagementConfig, opts ...dcl.ApplyOption) *InstanceManagementConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.StandardManagementConfig != nil || (initial != nil && initial.StandardManagementConfig != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.FullManagementConfig) {
			des.StandardManagementConfig = nil
			if initial != nil {
				initial.StandardManagementConfig = nil
			}
		}
	}

	if des.FullManagementConfig != nil || (initial != nil && initial.FullManagementConfig != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.StandardManagementConfig) {
			des.FullManagementConfig = nil
			if initial != nil {
				initial.FullManagementConfig = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceManagementConfig{}

	cDes.StandardManagementConfig = canonicalizeInstanceManagementConfigStandardManagementConfig(des.StandardManagementConfig, initial.StandardManagementConfig, opts...)
	cDes.FullManagementConfig = canonicalizeInstanceManagementConfigFullManagementConfig(des.FullManagementConfig, initial.FullManagementConfig, opts...)

	return cDes
}

func canonicalizeInstanceManagementConfigSlice(des, initial []InstanceManagementConfig, opts ...dcl.ApplyOption) []InstanceManagementConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InstanceManagementConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInstanceManagementConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InstanceManagementConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInstanceManagementConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInstanceManagementConfig(c *Client, des, nw *InstanceManagementConfig) *InstanceManagementConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InstanceManagementConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.StandardManagementConfig = canonicalizeNewInstanceManagementConfigStandardManagementConfig(c, des.StandardManagementConfig, nw.StandardManagementConfig)
	nw.FullManagementConfig = canonicalizeNewInstanceManagementConfigFullManagementConfig(c, des.FullManagementConfig, nw.FullManagementConfig)

	return nw
}

func canonicalizeNewInstanceManagementConfigSet(c *Client, des, nw []InstanceManagementConfig) []InstanceManagementConfig {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InstanceManagementConfig
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInstanceManagementConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInstanceManagementConfig(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInstanceManagementConfigSlice(c *Client, des, nw []InstanceManagementConfig) []InstanceManagementConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceManagementConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceManagementConfig(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceManagementConfigStandardManagementConfig(des, initial *InstanceManagementConfigStandardManagementConfig, opts ...dcl.ApplyOption) *InstanceManagementConfigStandardManagementConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceManagementConfigStandardManagementConfig{}

	if dcl.IsZeroValue(des.Network) || (dcl.IsEmptyValueIndirect(des.Network) && dcl.IsEmptyValueIndirect(initial.Network)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Network = initial.Network
	} else {
		cDes.Network = des.Network
	}
	if dcl.StringCanonicalize(des.MasterIPv4CidrBlock, initial.MasterIPv4CidrBlock) || dcl.IsZeroValue(des.MasterIPv4CidrBlock) {
		cDes.MasterIPv4CidrBlock = initial.MasterIPv4CidrBlock
	} else {
		cDes.MasterIPv4CidrBlock = des.MasterIPv4CidrBlock
	}
	if dcl.StringCanonicalize(des.ManBlock, initial.ManBlock) || dcl.IsZeroValue(des.ManBlock) {
		cDes.ManBlock = initial.ManBlock
	} else {
		cDes.ManBlock = des.ManBlock
	}
	if dcl.StringCanonicalize(des.ClusterCidrBlock, initial.ClusterCidrBlock) || dcl.IsZeroValue(des.ClusterCidrBlock) {
		cDes.ClusterCidrBlock = initial.ClusterCidrBlock
	} else {
		cDes.ClusterCidrBlock = des.ClusterCidrBlock
	}
	if dcl.StringCanonicalize(des.ServicesCidrBlock, initial.ServicesCidrBlock) || dcl.IsZeroValue(des.ServicesCidrBlock) {
		cDes.ServicesCidrBlock = initial.ServicesCidrBlock
	} else {
		cDes.ServicesCidrBlock = des.ServicesCidrBlock
	}
	if dcl.StringCanonicalize(des.ClusterNamedRange, initial.ClusterNamedRange) || dcl.IsZeroValue(des.ClusterNamedRange) {
		cDes.ClusterNamedRange = initial.ClusterNamedRange
	} else {
		cDes.ClusterNamedRange = des.ClusterNamedRange
	}
	if dcl.StringCanonicalize(des.ServicesNamedRange, initial.ServicesNamedRange) || dcl.IsZeroValue(des.ServicesNamedRange) {
		cDes.ServicesNamedRange = initial.ServicesNamedRange
	} else {
		cDes.ServicesNamedRange = des.ServicesNamedRange
	}

	return cDes
}

func canonicalizeInstanceManagementConfigStandardManagementConfigSlice(des, initial []InstanceManagementConfigStandardManagementConfig, opts ...dcl.ApplyOption) []InstanceManagementConfigStandardManagementConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InstanceManagementConfigStandardManagementConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInstanceManagementConfigStandardManagementConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InstanceManagementConfigStandardManagementConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInstanceManagementConfigStandardManagementConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInstanceManagementConfigStandardManagementConfig(c *Client, des, nw *InstanceManagementConfigStandardManagementConfig) *InstanceManagementConfigStandardManagementConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InstanceManagementConfigStandardManagementConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.MasterIPv4CidrBlock, nw.MasterIPv4CidrBlock) {
		nw.MasterIPv4CidrBlock = des.MasterIPv4CidrBlock
	}
	if dcl.StringCanonicalize(des.ManBlock, nw.ManBlock) {
		nw.ManBlock = des.ManBlock
	}
	if dcl.StringCanonicalize(des.ClusterCidrBlock, nw.ClusterCidrBlock) {
		nw.ClusterCidrBlock = des.ClusterCidrBlock
	}
	if dcl.StringCanonicalize(des.ServicesCidrBlock, nw.ServicesCidrBlock) {
		nw.ServicesCidrBlock = des.ServicesCidrBlock
	}
	if dcl.StringCanonicalize(des.ClusterNamedRange, nw.ClusterNamedRange) {
		nw.ClusterNamedRange = des.ClusterNamedRange
	}
	if dcl.StringCanonicalize(des.ServicesNamedRange, nw.ServicesNamedRange) {
		nw.ServicesNamedRange = des.ServicesNamedRange
	}

	return nw
}

func canonicalizeNewInstanceManagementConfigStandardManagementConfigSet(c *Client, des, nw []InstanceManagementConfigStandardManagementConfig) []InstanceManagementConfigStandardManagementConfig {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InstanceManagementConfigStandardManagementConfig
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInstanceManagementConfigStandardManagementConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInstanceManagementConfigStandardManagementConfig(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInstanceManagementConfigStandardManagementConfigSlice(c *Client, des, nw []InstanceManagementConfigStandardManagementConfig) []InstanceManagementConfigStandardManagementConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceManagementConfigStandardManagementConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceManagementConfigStandardManagementConfig(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceManagementConfigFullManagementConfig(des, initial *InstanceManagementConfigFullManagementConfig, opts ...dcl.ApplyOption) *InstanceManagementConfigFullManagementConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceManagementConfigFullManagementConfig{}

	if dcl.IsZeroValue(des.Network) || (dcl.IsEmptyValueIndirect(des.Network) && dcl.IsEmptyValueIndirect(initial.Network)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Network = initial.Network
	} else {
		cDes.Network = des.Network
	}
	if dcl.StringCanonicalize(des.MasterIPv4CidrBlock, initial.MasterIPv4CidrBlock) || dcl.IsZeroValue(des.MasterIPv4CidrBlock) {
		cDes.MasterIPv4CidrBlock = initial.MasterIPv4CidrBlock
	} else {
		cDes.MasterIPv4CidrBlock = des.MasterIPv4CidrBlock
	}
	if dcl.StringCanonicalize(des.ManBlock, initial.ManBlock) || dcl.IsZeroValue(des.ManBlock) {
		cDes.ManBlock = initial.ManBlock
	} else {
		cDes.ManBlock = des.ManBlock
	}
	if dcl.StringCanonicalize(des.ClusterCidrBlock, initial.ClusterCidrBlock) || dcl.IsZeroValue(des.ClusterCidrBlock) {
		cDes.ClusterCidrBlock = initial.ClusterCidrBlock
	} else {
		cDes.ClusterCidrBlock = des.ClusterCidrBlock
	}
	if dcl.StringCanonicalize(des.ServicesCidrBlock, initial.ServicesCidrBlock) || dcl.IsZeroValue(des.ServicesCidrBlock) {
		cDes.ServicesCidrBlock = initial.ServicesCidrBlock
	} else {
		cDes.ServicesCidrBlock = des.ServicesCidrBlock
	}
	if dcl.StringCanonicalize(des.ClusterNamedRange, initial.ClusterNamedRange) || dcl.IsZeroValue(des.ClusterNamedRange) {
		cDes.ClusterNamedRange = initial.ClusterNamedRange
	} else {
		cDes.ClusterNamedRange = des.ClusterNamedRange
	}
	if dcl.StringCanonicalize(des.ServicesNamedRange, initial.ServicesNamedRange) || dcl.IsZeroValue(des.ServicesNamedRange) {
		cDes.ServicesNamedRange = initial.ServicesNamedRange
	} else {
		cDes.ServicesNamedRange = des.ServicesNamedRange
	}

	return cDes
}

func canonicalizeInstanceManagementConfigFullManagementConfigSlice(des, initial []InstanceManagementConfigFullManagementConfig, opts ...dcl.ApplyOption) []InstanceManagementConfigFullManagementConfig {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InstanceManagementConfigFullManagementConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInstanceManagementConfigFullManagementConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InstanceManagementConfigFullManagementConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInstanceManagementConfigFullManagementConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInstanceManagementConfigFullManagementConfig(c *Client, des, nw *InstanceManagementConfigFullManagementConfig) *InstanceManagementConfigFullManagementConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InstanceManagementConfigFullManagementConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.MasterIPv4CidrBlock, nw.MasterIPv4CidrBlock) {
		nw.MasterIPv4CidrBlock = des.MasterIPv4CidrBlock
	}
	if dcl.StringCanonicalize(des.ManBlock, nw.ManBlock) {
		nw.ManBlock = des.ManBlock
	}
	if dcl.StringCanonicalize(des.ClusterCidrBlock, nw.ClusterCidrBlock) {
		nw.ClusterCidrBlock = des.ClusterCidrBlock
	}
	if dcl.StringCanonicalize(des.ServicesCidrBlock, nw.ServicesCidrBlock) {
		nw.ServicesCidrBlock = des.ServicesCidrBlock
	}
	if dcl.StringCanonicalize(des.ClusterNamedRange, nw.ClusterNamedRange) {
		nw.ClusterNamedRange = des.ClusterNamedRange
	}
	if dcl.StringCanonicalize(des.ServicesNamedRange, nw.ServicesNamedRange) {
		nw.ServicesNamedRange = des.ServicesNamedRange
	}

	return nw
}

func canonicalizeNewInstanceManagementConfigFullManagementConfigSet(c *Client, des, nw []InstanceManagementConfigFullManagementConfig) []InstanceManagementConfigFullManagementConfig {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InstanceManagementConfigFullManagementConfig
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInstanceManagementConfigFullManagementConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInstanceManagementConfigFullManagementConfig(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInstanceManagementConfigFullManagementConfigSlice(c *Client, des, nw []InstanceManagementConfigFullManagementConfig) []InstanceManagementConfigFullManagementConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceManagementConfigFullManagementConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceManagementConfigFullManagementConfig(c, &d, &n))
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

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BundlesConfig, actual.BundlesConfig, dcl.DiffInfo{ObjectFunction: compareInstanceBundlesConfigNewStyle, EmptyObject: EmptyInstanceBundlesConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BundlesConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UsePrivateEndpoint, actual.UsePrivateEndpoint, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UsePrivateEndpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GkeResourceLink, actual.GkeResourceLink, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GkeResourceLink")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.ManagementConfig, actual.ManagementConfig, dcl.DiffInfo{ObjectFunction: compareInstanceManagementConfigNewStyle, EmptyObject: EmptyInstanceManagementConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ManagementConfig")); len(ds) != 0 || err != nil {
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
func compareInstanceBundlesConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceBundlesConfig)
	if !ok {
		desiredNotPointer, ok := d.(InstanceBundlesConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceBundlesConfig or *InstanceBundlesConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceBundlesConfig)
	if !ok {
		actualNotPointer, ok := a.(InstanceBundlesConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceBundlesConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ConfigControllerConfig, actual.ConfigControllerConfig, dcl.DiffInfo{ObjectFunction: compareInstanceBundlesConfigConfigControllerConfigNewStyle, EmptyObject: EmptyInstanceBundlesConfigConfigControllerConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ConfigControllerConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceBundlesConfigConfigControllerConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceBundlesConfigConfigControllerConfig)
	if !ok {
		desiredNotPointer, ok := d.(InstanceBundlesConfigConfigControllerConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceBundlesConfigConfigControllerConfig or *InstanceBundlesConfigConfigControllerConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceBundlesConfigConfigControllerConfig)
	if !ok {
		actualNotPointer, ok := a.(InstanceBundlesConfigConfigControllerConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceBundlesConfigConfigControllerConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceManagementConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceManagementConfig)
	if !ok {
		desiredNotPointer, ok := d.(InstanceManagementConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceManagementConfig or *InstanceManagementConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceManagementConfig)
	if !ok {
		actualNotPointer, ok := a.(InstanceManagementConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceManagementConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.StandardManagementConfig, actual.StandardManagementConfig, dcl.DiffInfo{ObjectFunction: compareInstanceManagementConfigStandardManagementConfigNewStyle, EmptyObject: EmptyInstanceManagementConfigStandardManagementConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StandardManagementConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FullManagementConfig, actual.FullManagementConfig, dcl.DiffInfo{ServerDefault: true, ObjectFunction: compareInstanceManagementConfigFullManagementConfigNewStyle, EmptyObject: EmptyInstanceManagementConfigFullManagementConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FullManagementConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceManagementConfigStandardManagementConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceManagementConfigStandardManagementConfig)
	if !ok {
		desiredNotPointer, ok := d.(InstanceManagementConfigStandardManagementConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceManagementConfigStandardManagementConfig or *InstanceManagementConfigStandardManagementConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceManagementConfigStandardManagementConfig)
	if !ok {
		actualNotPointer, ok := a.(InstanceManagementConfigStandardManagementConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceManagementConfigStandardManagementConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Network, actual.Network, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Network")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MasterIPv4CidrBlock, actual.MasterIPv4CidrBlock, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MasterIpv4CidrBlock")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ManBlock, actual.ManBlock, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ManBlock")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterCidrBlock, actual.ClusterCidrBlock, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ClusterCidrBlock")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServicesCidrBlock, actual.ServicesCidrBlock, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServicesCidrBlock")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterNamedRange, actual.ClusterNamedRange, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ClusterNamedRange")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServicesNamedRange, actual.ServicesNamedRange, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServicesNamedRange")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceManagementConfigFullManagementConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceManagementConfigFullManagementConfig)
	if !ok {
		desiredNotPointer, ok := d.(InstanceManagementConfigFullManagementConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceManagementConfigFullManagementConfig or *InstanceManagementConfigFullManagementConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceManagementConfigFullManagementConfig)
	if !ok {
		actualNotPointer, ok := a.(InstanceManagementConfigFullManagementConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceManagementConfigFullManagementConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Network, actual.Network, dcl.DiffInfo{ServerDefault: true, Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Network")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MasterIPv4CidrBlock, actual.MasterIPv4CidrBlock, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MasterIpv4CidrBlock")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ManBlock, actual.ManBlock, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ManBlock")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterCidrBlock, actual.ClusterCidrBlock, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ClusterCidrBlock")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServicesCidrBlock, actual.ServicesCidrBlock, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServicesCidrBlock")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterNamedRange, actual.ClusterNamedRange, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ClusterNamedRange")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServicesNamedRange, actual.ServicesNamedRange, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServicesNamedRange")); len(ds) != 0 || err != nil {
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
	normalized.GkeResourceLink = dcl.SelfLinkToName(r.GkeResourceLink)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Instance) updateURL(userBasePath, updateName string) (string, error) {
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
	m = EncodeInstanceCreateRequest(m)

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
	if v := f.Name; dcl.ValueShouldBeSent(v) {
		m["name"] = v
	}
	if v := f.Labels; dcl.ValueShouldBeSent(v) {
		m["labels"] = v
	}
	if v, err := expandInstanceBundlesConfig(c, f.BundlesConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding BundlesConfig into bundlesConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["bundlesConfig"] = v
	}
	if v := f.UsePrivateEndpoint; dcl.ValueShouldBeSent(v) {
		m["usePrivateEndpoint"] = v
	}
	if v, err := expandInstanceManagementConfig(c, f.ManagementConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding ManagementConfig into managementConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["managementConfig"] = v
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
	resultRes.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	resultRes.BundlesConfig = flattenInstanceBundlesConfig(c, m["bundlesConfig"], res)
	resultRes.UsePrivateEndpoint = dcl.FlattenBool(m["usePrivateEndpoint"])
	resultRes.GkeResourceLink = dcl.FlattenString(m["gkeResourceLink"])
	resultRes.State = flattenInstanceStateEnum(m["state"])
	resultRes.ManagementConfig = flattenInstanceManagementConfig(c, m["managementConfig"], res)
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// expandInstanceBundlesConfigMap expands the contents of InstanceBundlesConfig into a JSON
// request object.
func expandInstanceBundlesConfigMap(c *Client, f map[string]InstanceBundlesConfig, res *Instance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceBundlesConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceBundlesConfigSlice expands the contents of InstanceBundlesConfig into a JSON
// request object.
func expandInstanceBundlesConfigSlice(c *Client, f []InstanceBundlesConfig, res *Instance) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceBundlesConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceBundlesConfigMap flattens the contents of InstanceBundlesConfig from a JSON
// response object.
func flattenInstanceBundlesConfigMap(c *Client, i interface{}, res *Instance) map[string]InstanceBundlesConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceBundlesConfig{}
	}

	if len(a) == 0 {
		return map[string]InstanceBundlesConfig{}
	}

	items := make(map[string]InstanceBundlesConfig)
	for k, item := range a {
		items[k] = *flattenInstanceBundlesConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInstanceBundlesConfigSlice flattens the contents of InstanceBundlesConfig from a JSON
// response object.
func flattenInstanceBundlesConfigSlice(c *Client, i interface{}, res *Instance) []InstanceBundlesConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceBundlesConfig{}
	}

	if len(a) == 0 {
		return []InstanceBundlesConfig{}
	}

	items := make([]InstanceBundlesConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceBundlesConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInstanceBundlesConfig expands an instance of InstanceBundlesConfig into a JSON
// request object.
func expandInstanceBundlesConfig(c *Client, f *InstanceBundlesConfig, res *Instance) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandInstanceBundlesConfigConfigControllerConfig(c, f.ConfigControllerConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding ConfigControllerConfig into configControllerConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["configControllerConfig"] = v
	}

	return m, nil
}

// flattenInstanceBundlesConfig flattens an instance of InstanceBundlesConfig from a JSON
// response object.
func flattenInstanceBundlesConfig(c *Client, i interface{}, res *Instance) *InstanceBundlesConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceBundlesConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceBundlesConfig
	}
	r.ConfigControllerConfig = flattenInstanceBundlesConfigConfigControllerConfig(c, m["configControllerConfig"], res)

	return r
}

// expandInstanceBundlesConfigConfigControllerConfigMap expands the contents of InstanceBundlesConfigConfigControllerConfig into a JSON
// request object.
func expandInstanceBundlesConfigConfigControllerConfigMap(c *Client, f map[string]InstanceBundlesConfigConfigControllerConfig, res *Instance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceBundlesConfigConfigControllerConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceBundlesConfigConfigControllerConfigSlice expands the contents of InstanceBundlesConfigConfigControllerConfig into a JSON
// request object.
func expandInstanceBundlesConfigConfigControllerConfigSlice(c *Client, f []InstanceBundlesConfigConfigControllerConfig, res *Instance) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceBundlesConfigConfigControllerConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceBundlesConfigConfigControllerConfigMap flattens the contents of InstanceBundlesConfigConfigControllerConfig from a JSON
// response object.
func flattenInstanceBundlesConfigConfigControllerConfigMap(c *Client, i interface{}, res *Instance) map[string]InstanceBundlesConfigConfigControllerConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceBundlesConfigConfigControllerConfig{}
	}

	if len(a) == 0 {
		return map[string]InstanceBundlesConfigConfigControllerConfig{}
	}

	items := make(map[string]InstanceBundlesConfigConfigControllerConfig)
	for k, item := range a {
		items[k] = *flattenInstanceBundlesConfigConfigControllerConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInstanceBundlesConfigConfigControllerConfigSlice flattens the contents of InstanceBundlesConfigConfigControllerConfig from a JSON
// response object.
func flattenInstanceBundlesConfigConfigControllerConfigSlice(c *Client, i interface{}, res *Instance) []InstanceBundlesConfigConfigControllerConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceBundlesConfigConfigControllerConfig{}
	}

	if len(a) == 0 {
		return []InstanceBundlesConfigConfigControllerConfig{}
	}

	items := make([]InstanceBundlesConfigConfigControllerConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceBundlesConfigConfigControllerConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInstanceBundlesConfigConfigControllerConfig expands an instance of InstanceBundlesConfigConfigControllerConfig into a JSON
// request object.
func expandInstanceBundlesConfigConfigControllerConfig(c *Client, f *InstanceBundlesConfigConfigControllerConfig, res *Instance) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}

	return m, nil
}

// flattenInstanceBundlesConfigConfigControllerConfig flattens an instance of InstanceBundlesConfigConfigControllerConfig from a JSON
// response object.
func flattenInstanceBundlesConfigConfigControllerConfig(c *Client, i interface{}, res *Instance) *InstanceBundlesConfigConfigControllerConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceBundlesConfigConfigControllerConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceBundlesConfigConfigControllerConfig
	}
	r.Enabled = dcl.FlattenBool(m["enabled"])

	return r
}

// expandInstanceManagementConfigMap expands the contents of InstanceManagementConfig into a JSON
// request object.
func expandInstanceManagementConfigMap(c *Client, f map[string]InstanceManagementConfig, res *Instance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceManagementConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceManagementConfigSlice expands the contents of InstanceManagementConfig into a JSON
// request object.
func expandInstanceManagementConfigSlice(c *Client, f []InstanceManagementConfig, res *Instance) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceManagementConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceManagementConfigMap flattens the contents of InstanceManagementConfig from a JSON
// response object.
func flattenInstanceManagementConfigMap(c *Client, i interface{}, res *Instance) map[string]InstanceManagementConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceManagementConfig{}
	}

	if len(a) == 0 {
		return map[string]InstanceManagementConfig{}
	}

	items := make(map[string]InstanceManagementConfig)
	for k, item := range a {
		items[k] = *flattenInstanceManagementConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInstanceManagementConfigSlice flattens the contents of InstanceManagementConfig from a JSON
// response object.
func flattenInstanceManagementConfigSlice(c *Client, i interface{}, res *Instance) []InstanceManagementConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceManagementConfig{}
	}

	if len(a) == 0 {
		return []InstanceManagementConfig{}
	}

	items := make([]InstanceManagementConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceManagementConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInstanceManagementConfig expands an instance of InstanceManagementConfig into a JSON
// request object.
func expandInstanceManagementConfig(c *Client, f *InstanceManagementConfig, res *Instance) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandInstanceManagementConfigStandardManagementConfig(c, f.StandardManagementConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding StandardManagementConfig into standardManagementConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["standardManagementConfig"] = v
	}
	if v, err := expandInstanceManagementConfigFullManagementConfig(c, f.FullManagementConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding FullManagementConfig into fullManagementConfig: %w", err)
	} else if v != nil {
		m["fullManagementConfig"] = v
	}

	return m, nil
}

// flattenInstanceManagementConfig flattens an instance of InstanceManagementConfig from a JSON
// response object.
func flattenInstanceManagementConfig(c *Client, i interface{}, res *Instance) *InstanceManagementConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceManagementConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceManagementConfig
	}
	r.StandardManagementConfig = flattenInstanceManagementConfigStandardManagementConfig(c, m["standardManagementConfig"], res)
	r.FullManagementConfig = flattenInstanceManagementConfigFullManagementConfig(c, m["fullManagementConfig"], res)

	return r
}

// expandInstanceManagementConfigStandardManagementConfigMap expands the contents of InstanceManagementConfigStandardManagementConfig into a JSON
// request object.
func expandInstanceManagementConfigStandardManagementConfigMap(c *Client, f map[string]InstanceManagementConfigStandardManagementConfig, res *Instance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceManagementConfigStandardManagementConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceManagementConfigStandardManagementConfigSlice expands the contents of InstanceManagementConfigStandardManagementConfig into a JSON
// request object.
func expandInstanceManagementConfigStandardManagementConfigSlice(c *Client, f []InstanceManagementConfigStandardManagementConfig, res *Instance) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceManagementConfigStandardManagementConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceManagementConfigStandardManagementConfigMap flattens the contents of InstanceManagementConfigStandardManagementConfig from a JSON
// response object.
func flattenInstanceManagementConfigStandardManagementConfigMap(c *Client, i interface{}, res *Instance) map[string]InstanceManagementConfigStandardManagementConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceManagementConfigStandardManagementConfig{}
	}

	if len(a) == 0 {
		return map[string]InstanceManagementConfigStandardManagementConfig{}
	}

	items := make(map[string]InstanceManagementConfigStandardManagementConfig)
	for k, item := range a {
		items[k] = *flattenInstanceManagementConfigStandardManagementConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInstanceManagementConfigStandardManagementConfigSlice flattens the contents of InstanceManagementConfigStandardManagementConfig from a JSON
// response object.
func flattenInstanceManagementConfigStandardManagementConfigSlice(c *Client, i interface{}, res *Instance) []InstanceManagementConfigStandardManagementConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceManagementConfigStandardManagementConfig{}
	}

	if len(a) == 0 {
		return []InstanceManagementConfigStandardManagementConfig{}
	}

	items := make([]InstanceManagementConfigStandardManagementConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceManagementConfigStandardManagementConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInstanceManagementConfigStandardManagementConfig expands an instance of InstanceManagementConfigStandardManagementConfig into a JSON
// request object.
func expandInstanceManagementConfigStandardManagementConfig(c *Client, f *InstanceManagementConfigStandardManagementConfig, res *Instance) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Network; !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v := f.MasterIPv4CidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["masterIpv4CidrBlock"] = v
	}
	if v := f.ManBlock; !dcl.IsEmptyValueIndirect(v) {
		m["manBlock"] = v
	}
	if v := f.ClusterCidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["clusterCidrBlock"] = v
	}
	if v := f.ServicesCidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["servicesCidrBlock"] = v
	}
	if v := f.ClusterNamedRange; !dcl.IsEmptyValueIndirect(v) {
		m["clusterNamedRange"] = v
	}
	if v := f.ServicesNamedRange; !dcl.IsEmptyValueIndirect(v) {
		m["servicesNamedRange"] = v
	}

	return m, nil
}

// flattenInstanceManagementConfigStandardManagementConfig flattens an instance of InstanceManagementConfigStandardManagementConfig from a JSON
// response object.
func flattenInstanceManagementConfigStandardManagementConfig(c *Client, i interface{}, res *Instance) *InstanceManagementConfigStandardManagementConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceManagementConfigStandardManagementConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceManagementConfigStandardManagementConfig
	}
	r.Network = dcl.FlattenString(m["network"])
	r.MasterIPv4CidrBlock = dcl.FlattenString(m["masterIpv4CidrBlock"])
	r.ManBlock = dcl.FlattenString(m["manBlock"])
	r.ClusterCidrBlock = dcl.FlattenString(m["clusterCidrBlock"])
	r.ServicesCidrBlock = dcl.FlattenString(m["servicesCidrBlock"])
	r.ClusterNamedRange = dcl.FlattenString(m["clusterNamedRange"])
	r.ServicesNamedRange = dcl.FlattenString(m["servicesNamedRange"])

	return r
}

// expandInstanceManagementConfigFullManagementConfigMap expands the contents of InstanceManagementConfigFullManagementConfig into a JSON
// request object.
func expandInstanceManagementConfigFullManagementConfigMap(c *Client, f map[string]InstanceManagementConfigFullManagementConfig, res *Instance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceManagementConfigFullManagementConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceManagementConfigFullManagementConfigSlice expands the contents of InstanceManagementConfigFullManagementConfig into a JSON
// request object.
func expandInstanceManagementConfigFullManagementConfigSlice(c *Client, f []InstanceManagementConfigFullManagementConfig, res *Instance) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceManagementConfigFullManagementConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceManagementConfigFullManagementConfigMap flattens the contents of InstanceManagementConfigFullManagementConfig from a JSON
// response object.
func flattenInstanceManagementConfigFullManagementConfigMap(c *Client, i interface{}, res *Instance) map[string]InstanceManagementConfigFullManagementConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceManagementConfigFullManagementConfig{}
	}

	if len(a) == 0 {
		return map[string]InstanceManagementConfigFullManagementConfig{}
	}

	items := make(map[string]InstanceManagementConfigFullManagementConfig)
	for k, item := range a {
		items[k] = *flattenInstanceManagementConfigFullManagementConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInstanceManagementConfigFullManagementConfigSlice flattens the contents of InstanceManagementConfigFullManagementConfig from a JSON
// response object.
func flattenInstanceManagementConfigFullManagementConfigSlice(c *Client, i interface{}, res *Instance) []InstanceManagementConfigFullManagementConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceManagementConfigFullManagementConfig{}
	}

	if len(a) == 0 {
		return []InstanceManagementConfigFullManagementConfig{}
	}

	items := make([]InstanceManagementConfigFullManagementConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceManagementConfigFullManagementConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInstanceManagementConfigFullManagementConfig expands an instance of InstanceManagementConfigFullManagementConfig into a JSON
// request object.
func expandInstanceManagementConfigFullManagementConfig(c *Client, f *InstanceManagementConfigFullManagementConfig, res *Instance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Network; !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v := f.MasterIPv4CidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["masterIpv4CidrBlock"] = v
	}
	if v := f.ManBlock; !dcl.IsEmptyValueIndirect(v) {
		m["manBlock"] = v
	}
	if v := f.ClusterCidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["clusterCidrBlock"] = v
	}
	if v := f.ServicesCidrBlock; !dcl.IsEmptyValueIndirect(v) {
		m["servicesCidrBlock"] = v
	}
	if v := f.ClusterNamedRange; !dcl.IsEmptyValueIndirect(v) {
		m["clusterNamedRange"] = v
	}
	if v := f.ServicesNamedRange; !dcl.IsEmptyValueIndirect(v) {
		m["servicesNamedRange"] = v
	}

	return m, nil
}

// flattenInstanceManagementConfigFullManagementConfig flattens an instance of InstanceManagementConfigFullManagementConfig from a JSON
// response object.
func flattenInstanceManagementConfigFullManagementConfig(c *Client, i interface{}, res *Instance) *InstanceManagementConfigFullManagementConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceManagementConfigFullManagementConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceManagementConfigFullManagementConfig
	}
	r.Network = dcl.FlattenString(m["network"])
	r.MasterIPv4CidrBlock = dcl.FlattenString(m["masterIpv4CidrBlock"])
	r.ManBlock = dcl.FlattenString(m["manBlock"])
	r.ClusterCidrBlock = dcl.FlattenString(m["clusterCidrBlock"])
	r.ServicesCidrBlock = dcl.FlattenString(m["servicesCidrBlock"])
	r.ClusterNamedRange = dcl.FlattenString(m["clusterNamedRange"])
	r.ServicesNamedRange = dcl.FlattenString(m["servicesNamedRange"])

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

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractInstanceFields(r *Instance) error {
	vBundlesConfig := r.BundlesConfig
	if vBundlesConfig == nil {
		// note: explicitly not the empty object.
		vBundlesConfig = &InstanceBundlesConfig{}
	}
	if err := extractInstanceBundlesConfigFields(r, vBundlesConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vBundlesConfig) {
		r.BundlesConfig = vBundlesConfig
	}
	vManagementConfig := r.ManagementConfig
	if vManagementConfig == nil {
		// note: explicitly not the empty object.
		vManagementConfig = &InstanceManagementConfig{}
	}
	if err := extractInstanceManagementConfigFields(r, vManagementConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vManagementConfig) {
		r.ManagementConfig = vManagementConfig
	}
	return nil
}
func extractInstanceBundlesConfigFields(r *Instance, o *InstanceBundlesConfig) error {
	vConfigControllerConfig := o.ConfigControllerConfig
	if vConfigControllerConfig == nil {
		// note: explicitly not the empty object.
		vConfigControllerConfig = &InstanceBundlesConfigConfigControllerConfig{}
	}
	if err := extractInstanceBundlesConfigConfigControllerConfigFields(r, vConfigControllerConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vConfigControllerConfig) {
		o.ConfigControllerConfig = vConfigControllerConfig
	}
	return nil
}
func extractInstanceBundlesConfigConfigControllerConfigFields(r *Instance, o *InstanceBundlesConfigConfigControllerConfig) error {
	return nil
}
func extractInstanceManagementConfigFields(r *Instance, o *InstanceManagementConfig) error {
	vStandardManagementConfig := o.StandardManagementConfig
	if vStandardManagementConfig == nil {
		// note: explicitly not the empty object.
		vStandardManagementConfig = &InstanceManagementConfigStandardManagementConfig{}
	}
	if err := extractInstanceManagementConfigStandardManagementConfigFields(r, vStandardManagementConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vStandardManagementConfig) {
		o.StandardManagementConfig = vStandardManagementConfig
	}
	vFullManagementConfig := o.FullManagementConfig
	if vFullManagementConfig == nil {
		// note: explicitly not the empty object.
		vFullManagementConfig = &InstanceManagementConfigFullManagementConfig{}
	}
	if err := extractInstanceManagementConfigFullManagementConfigFields(r, vFullManagementConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vFullManagementConfig) {
		o.FullManagementConfig = vFullManagementConfig
	}
	return nil
}
func extractInstanceManagementConfigStandardManagementConfigFields(r *Instance, o *InstanceManagementConfigStandardManagementConfig) error {
	return nil
}
func extractInstanceManagementConfigFullManagementConfigFields(r *Instance, o *InstanceManagementConfigFullManagementConfig) error {
	return nil
}

func postReadExtractInstanceFields(r *Instance) error {
	vBundlesConfig := r.BundlesConfig
	if vBundlesConfig == nil {
		// note: explicitly not the empty object.
		vBundlesConfig = &InstanceBundlesConfig{}
	}
	if err := postReadExtractInstanceBundlesConfigFields(r, vBundlesConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vBundlesConfig) {
		r.BundlesConfig = vBundlesConfig
	}
	vManagementConfig := r.ManagementConfig
	if vManagementConfig == nil {
		// note: explicitly not the empty object.
		vManagementConfig = &InstanceManagementConfig{}
	}
	if err := postReadExtractInstanceManagementConfigFields(r, vManagementConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vManagementConfig) {
		r.ManagementConfig = vManagementConfig
	}
	return nil
}
func postReadExtractInstanceBundlesConfigFields(r *Instance, o *InstanceBundlesConfig) error {
	vConfigControllerConfig := o.ConfigControllerConfig
	if vConfigControllerConfig == nil {
		// note: explicitly not the empty object.
		vConfigControllerConfig = &InstanceBundlesConfigConfigControllerConfig{}
	}
	if err := extractInstanceBundlesConfigConfigControllerConfigFields(r, vConfigControllerConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vConfigControllerConfig) {
		o.ConfigControllerConfig = vConfigControllerConfig
	}
	return nil
}
func postReadExtractInstanceBundlesConfigConfigControllerConfigFields(r *Instance, o *InstanceBundlesConfigConfigControllerConfig) error {
	return nil
}
func postReadExtractInstanceManagementConfigFields(r *Instance, o *InstanceManagementConfig) error {
	vStandardManagementConfig := o.StandardManagementConfig
	if vStandardManagementConfig == nil {
		// note: explicitly not the empty object.
		vStandardManagementConfig = &InstanceManagementConfigStandardManagementConfig{}
	}
	if err := extractInstanceManagementConfigStandardManagementConfigFields(r, vStandardManagementConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vStandardManagementConfig) {
		o.StandardManagementConfig = vStandardManagementConfig
	}
	vFullManagementConfig := o.FullManagementConfig
	if vFullManagementConfig == nil {
		// note: explicitly not the empty object.
		vFullManagementConfig = &InstanceManagementConfigFullManagementConfig{}
	}
	if err := extractInstanceManagementConfigFullManagementConfigFields(r, vFullManagementConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vFullManagementConfig) {
		o.FullManagementConfig = vFullManagementConfig
	}
	return nil
}
func postReadExtractInstanceManagementConfigStandardManagementConfigFields(r *Instance, o *InstanceManagementConfigStandardManagementConfig) error {
	return nil
}
func postReadExtractInstanceManagementConfigFullManagementConfigFields(r *Instance, o *InstanceManagementConfigFullManagementConfig) error {
	return nil
}
