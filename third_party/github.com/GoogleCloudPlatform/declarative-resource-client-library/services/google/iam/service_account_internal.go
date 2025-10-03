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
package iam

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *ServiceAccount) validate() error {

	if !dcl.IsEmptyValueIndirect(r.ActasResources) {
		if err := r.ActasResources.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceAccountActasResources) validate() error {
	return nil
}
func (r *ServiceAccountActasResourcesResources) validate() error {
	return nil
}
func (r *ServiceAccount) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://iam.googleapis.com/v1/", params)
}

func (r *ServiceAccount) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.URL("projects/{{project}}/serviceAccounts", nr.basePath(), userBasePath, params), nil

}

// serviceAccountApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type serviceAccountApiOperation interface {
	do(context.Context, *ServiceAccount, *Client) error
}

// newUpdateServiceAccountPatchServiceAccountRequest creates a request for an
// ServiceAccount resource's PatchServiceAccount update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateServiceAccountPatchServiceAccountRequest(ctx context.Context, f *ServiceAccount, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	return req, nil
}

// marshalUpdateServiceAccountPatchServiceAccountRequest converts the update into
// the final JSON request body.
func marshalUpdateServiceAccountPatchServiceAccountRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	dcl.MoveMapEntry(
		m,
		[]string{"projectId"},
		[]string{},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"email"},
		[]string{},
	)
	return json.Marshal(m)
}

type updateServiceAccountPatchServiceAccountOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateServiceAccountPatchServiceAccountOperation) do(ctx context.Context, r *ServiceAccount, c *Client) error {
	_, err := c.GetServiceAccount(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "PatchServiceAccount")
	if err != nil {
		return err
	}

	req, err := newUpdateServiceAccountPatchServiceAccountRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateServiceAccountPatchServiceAccountRequest(c, req)
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

func (c *Client) listServiceAccountRaw(ctx context.Context, r *ServiceAccount, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ServiceAccountMaxPage {
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

type listServiceAccountOperation struct {
	Accounts []map[string]interface{} `json:"accounts"`
	Token    string                   `json:"nextPageToken"`
}

func (c *Client) listServiceAccount(ctx context.Context, r *ServiceAccount, pageToken string, pageSize int32) ([]*ServiceAccount, string, error) {
	b, err := c.listServiceAccountRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listServiceAccountOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*ServiceAccount
	for _, v := range m.Accounts {
		res, err := unmarshalMapServiceAccount(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllServiceAccount(ctx context.Context, f func(*ServiceAccount) bool, resources []*ServiceAccount) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteServiceAccount(ctx, res)
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

type deleteServiceAccountOperation struct{}

func (op *deleteServiceAccountOperation) do(ctx context.Context, r *ServiceAccount, c *Client) error {
	r, err := c.GetServiceAccount(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "ServiceAccount not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetServiceAccount checking for existence. error: %v", err)
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
		return fmt.Errorf("failed to delete ServiceAccount: %w", err)
	}

	// We saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// This is the reason we are adding retry to handle that case.
	retriesRemaining := 10
	dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		_, err := c.GetServiceAccount(ctx, r)
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
type createServiceAccountOperation struct {
	response map[string]interface{}
}

func (op *createServiceAccountOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createServiceAccountOperation) do(ctx context.Context, r *ServiceAccount, c *Client) error {
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

	// Poll for the ServiceAccount resource to be created. ServiceAccount resources are eventually consistent but do not support operations
	// so we must repeatedly poll to check for their creation.
	requiredSuccesses := 1
	start := time.Now()
	err = dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		u, err := r.getURL(c.Config.BasePath)
		if err != nil {
			return nil, err
		}
		getResp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, nil)
		if err != nil {
			// If the error is a transient server error (e.g., 500) or not found (i.e., the resource has not yet been created),
			// continue retrying until the transient error is resolved, the resource is created, or we time out.
			if dcl.IsRetryableRequestError(c.Config, err, true, start) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		getResp.Response.Body.Close()
		requiredSuccesses--
		if requiredSuccesses > 0 {
			return &dcl.RetryDetails{}, dcl.OperationNotDone{}
		}
		return getResp, nil
	}, c.Config.RetryProvider)

	if _, err := c.GetServiceAccount(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getServiceAccountRaw(ctx context.Context, r *ServiceAccount) ([]byte, error) {

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

func (c *Client) serviceAccountDiffsForRawDesired(ctx context.Context, rawDesired *ServiceAccount, opts ...dcl.ApplyOption) (initial, desired *ServiceAccount, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *ServiceAccount
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*ServiceAccount); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected ServiceAccount, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetServiceAccount(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a ServiceAccount resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve ServiceAccount resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that ServiceAccount resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeServiceAccountDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for ServiceAccount: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for ServiceAccount: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractServiceAccountFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeServiceAccountInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for ServiceAccount: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeServiceAccountDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for ServiceAccount: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffServiceAccount(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeServiceAccountInitialState(rawInitial, rawDesired *ServiceAccount) (*ServiceAccount, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeServiceAccountDesiredState(rawDesired, rawInitial *ServiceAccount, opts ...dcl.ApplyOption) (*ServiceAccount, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.ActasResources = canonicalizeServiceAccountActasResources(rawDesired.ActasResources, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &ServiceAccount{}
	if canonicalizeServiceAccountName(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.IsZeroValue(rawDesired.Project) || (dcl.IsEmptyValueIndirect(rawDesired.Project) && dcl.IsEmptyValueIndirect(rawInitial.Project)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
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
	canonicalDesired.ActasResources = canonicalizeServiceAccountActasResources(rawDesired.ActasResources, rawInitial.ActasResources, opts...)
	return canonicalDesired, nil
}

func canonicalizeServiceAccountNewState(c *Client, rawNew, rawDesired *ServiceAccount) (*ServiceAccount, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if canonicalizeServiceAccountName(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Project) && dcl.IsEmptyValueIndirect(rawDesired.Project) {
		rawNew.Project = rawDesired.Project
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UniqueId) && dcl.IsEmptyValueIndirect(rawDesired.UniqueId) {
		rawNew.UniqueId = rawDesired.UniqueId
	} else {
		if dcl.StringCanonicalize(rawDesired.UniqueId, rawNew.UniqueId) {
			rawNew.UniqueId = rawDesired.UniqueId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Email) && dcl.IsEmptyValueIndirect(rawDesired.Email) {
		rawNew.Email = rawDesired.Email
	} else {
		if dcl.StringCanonicalize(rawDesired.Email, rawNew.Email) {
			rawNew.Email = rawDesired.Email
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

	if dcl.IsEmptyValueIndirect(rawNew.OAuth2ClientId) && dcl.IsEmptyValueIndirect(rawDesired.OAuth2ClientId) {
		rawNew.OAuth2ClientId = rawDesired.OAuth2ClientId
	} else {
		if dcl.StringCanonicalize(rawDesired.OAuth2ClientId, rawNew.OAuth2ClientId) {
			rawNew.OAuth2ClientId = rawDesired.OAuth2ClientId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ActasResources) && dcl.IsEmptyValueIndirect(rawDesired.ActasResources) {
		rawNew.ActasResources = rawDesired.ActasResources
	} else {
		rawNew.ActasResources = canonicalizeNewServiceAccountActasResources(c, rawDesired.ActasResources, rawNew.ActasResources)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Disabled) && dcl.IsEmptyValueIndirect(rawDesired.Disabled) {
		rawNew.Disabled = rawDesired.Disabled
	} else {
		if dcl.BoolCanonicalize(rawDesired.Disabled, rawNew.Disabled) {
			rawNew.Disabled = rawDesired.Disabled
		}
	}

	return rawNew, nil
}

func canonicalizeServiceAccountActasResources(des, initial *ServiceAccountActasResources, opts ...dcl.ApplyOption) *ServiceAccountActasResources {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceAccountActasResources{}

	cDes.Resources = canonicalizeServiceAccountActasResourcesResourcesSlice(des.Resources, initial.Resources, opts...)

	return cDes
}

func canonicalizeServiceAccountActasResourcesSlice(des, initial []ServiceAccountActasResources, opts ...dcl.ApplyOption) []ServiceAccountActasResources {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceAccountActasResources, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceAccountActasResources(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceAccountActasResources, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceAccountActasResources(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceAccountActasResources(c *Client, des, nw *ServiceAccountActasResources) *ServiceAccountActasResources {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceAccountActasResources while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Resources = canonicalizeNewServiceAccountActasResourcesResourcesSlice(c, des.Resources, nw.Resources)

	return nw
}

func canonicalizeNewServiceAccountActasResourcesSet(c *Client, des, nw []ServiceAccountActasResources) []ServiceAccountActasResources {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ServiceAccountActasResources
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareServiceAccountActasResourcesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewServiceAccountActasResources(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewServiceAccountActasResourcesSlice(c *Client, des, nw []ServiceAccountActasResources) []ServiceAccountActasResources {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceAccountActasResources
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceAccountActasResources(c, &d, &n))
	}

	return items
}

func canonicalizeServiceAccountActasResourcesResources(des, initial *ServiceAccountActasResourcesResources, opts ...dcl.ApplyOption) *ServiceAccountActasResourcesResources {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceAccountActasResourcesResources{}

	if dcl.StringCanonicalize(des.FullResourceName, initial.FullResourceName) || dcl.IsZeroValue(des.FullResourceName) {
		cDes.FullResourceName = initial.FullResourceName
	} else {
		cDes.FullResourceName = des.FullResourceName
	}

	return cDes
}

func canonicalizeServiceAccountActasResourcesResourcesSlice(des, initial []ServiceAccountActasResourcesResources, opts ...dcl.ApplyOption) []ServiceAccountActasResourcesResources {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceAccountActasResourcesResources, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceAccountActasResourcesResources(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceAccountActasResourcesResources, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceAccountActasResourcesResources(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceAccountActasResourcesResources(c *Client, des, nw *ServiceAccountActasResourcesResources) *ServiceAccountActasResourcesResources {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceAccountActasResourcesResources while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.FullResourceName, nw.FullResourceName) {
		nw.FullResourceName = des.FullResourceName
	}

	return nw
}

func canonicalizeNewServiceAccountActasResourcesResourcesSet(c *Client, des, nw []ServiceAccountActasResourcesResources) []ServiceAccountActasResourcesResources {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ServiceAccountActasResourcesResources
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareServiceAccountActasResourcesResourcesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewServiceAccountActasResourcesResources(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewServiceAccountActasResourcesResourcesSlice(c *Client, des, nw []ServiceAccountActasResourcesResources) []ServiceAccountActasResourcesResources {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceAccountActasResourcesResources
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceAccountActasResourcesResources(c, &d, &n))
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
func diffServiceAccount(c *Client, desired, actual *ServiceAccount, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{CustomDiff: canonicalizeServiceAccountName, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ProjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UniqueId, actual.UniqueId, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UniqueId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Email, actual.Email, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Email")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceAccountPatchServiceAccountOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceAccountPatchServiceAccountOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OAuth2ClientId, actual.OAuth2ClientId, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Oauth2ClientId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ActasResources, actual.ActasResources, dcl.DiffInfo{ObjectFunction: compareServiceAccountActasResourcesNewStyle, EmptyObject: EmptyServiceAccountActasResources, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ActasResources")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Disabled, actual.Disabled, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Disabled")); len(ds) != 0 || err != nil {
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
func compareServiceAccountActasResourcesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceAccountActasResources)
	if !ok {
		desiredNotPointer, ok := d.(ServiceAccountActasResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceAccountActasResources or *ServiceAccountActasResources", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceAccountActasResources)
	if !ok {
		actualNotPointer, ok := a.(ServiceAccountActasResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceAccountActasResources", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Resources, actual.Resources, dcl.DiffInfo{ObjectFunction: compareServiceAccountActasResourcesResourcesNewStyle, EmptyObject: EmptyServiceAccountActasResourcesResources, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Resources")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceAccountActasResourcesResourcesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceAccountActasResourcesResources)
	if !ok {
		desiredNotPointer, ok := d.(ServiceAccountActasResourcesResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceAccountActasResourcesResources or *ServiceAccountActasResourcesResources", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceAccountActasResourcesResources)
	if !ok {
		actualNotPointer, ok := a.(ServiceAccountActasResourcesResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceAccountActasResourcesResources", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.FullResourceName, actual.FullResourceName, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FullResourceName")); len(ds) != 0 || err != nil {
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
func (r *ServiceAccount) urlNormalized() *ServiceAccount {
	normalized := dcl.Copy(*r).(ServiceAccount)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.UniqueId = dcl.SelfLinkToName(r.UniqueId)
	normalized.Email = dcl.SelfLinkToName(r.Email)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.OAuth2ClientId = dcl.SelfLinkToName(r.OAuth2ClientId)
	return &normalized
}

func (r *ServiceAccount) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "PatchServiceAccount" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(nr.Project),
			"name":    dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/serviceAccounts/{{name}}@{{project}}.iam.gserviceaccount.com", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the ServiceAccount resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *ServiceAccount) marshal(c *Client) ([]byte, error) {
	m, err := expandServiceAccount(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling ServiceAccount: %w", err)
	}
	dcl.MoveMapEntry(
		m,
		[]string{"projectId"},
		[]string{},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"email"},
		[]string{},
	)
	m = EncodeServiceAccountCreateRequest(m)

	return json.Marshal(m)
}

// unmarshalServiceAccount decodes JSON responses into the ServiceAccount resource schema.
func unmarshalServiceAccount(b []byte, c *Client, res *ServiceAccount) (*ServiceAccount, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapServiceAccount(m, c, res)
}

func unmarshalMapServiceAccount(m map[string]interface{}, c *Client, res *ServiceAccount) (*ServiceAccount, error) {

	flattened := flattenServiceAccount(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandServiceAccount expands ServiceAccount into a JSON request object.
func expandServiceAccount(c *Client, f *ServiceAccount) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("projects/%s/serviceAccounts/%s@%s.iam.gserviceaccounts.com", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Project; dcl.ValueShouldBeSent(v) {
		m["projectId"] = v
	}
	if v := f.DisplayName; dcl.ValueShouldBeSent(v) {
		m["displayName"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v, err := expandServiceAccountActasResources(c, f.ActasResources, res); err != nil {
		return nil, fmt.Errorf("error expanding ActasResources into actasResources: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["actasResources"] = v
	}

	return m, nil
}

// flattenServiceAccount flattens ServiceAccount from a JSON request object into the
// ServiceAccount type.
func flattenServiceAccount(c *Client, i interface{}, res *ServiceAccount) *ServiceAccount {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &ServiceAccount{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.Project = dcl.FlattenString(m["projectId"])
	resultRes.UniqueId = dcl.FlattenString(m["uniqueId"])
	resultRes.Email = dcl.FlattenString(m["email"])
	resultRes.DisplayName = dcl.FlattenString(m["displayName"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.OAuth2ClientId = dcl.FlattenString(m["oauth2ClientId"])
	resultRes.ActasResources = flattenServiceAccountActasResources(c, m["actasResources"], res)
	resultRes.Disabled = dcl.FlattenBool(m["disabled"])

	return resultRes
}

// expandServiceAccountActasResourcesMap expands the contents of ServiceAccountActasResources into a JSON
// request object.
func expandServiceAccountActasResourcesMap(c *Client, f map[string]ServiceAccountActasResources, res *ServiceAccount) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceAccountActasResources(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceAccountActasResourcesSlice expands the contents of ServiceAccountActasResources into a JSON
// request object.
func expandServiceAccountActasResourcesSlice(c *Client, f []ServiceAccountActasResources, res *ServiceAccount) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceAccountActasResources(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceAccountActasResourcesMap flattens the contents of ServiceAccountActasResources from a JSON
// response object.
func flattenServiceAccountActasResourcesMap(c *Client, i interface{}, res *ServiceAccount) map[string]ServiceAccountActasResources {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceAccountActasResources{}
	}

	if len(a) == 0 {
		return map[string]ServiceAccountActasResources{}
	}

	items := make(map[string]ServiceAccountActasResources)
	for k, item := range a {
		items[k] = *flattenServiceAccountActasResources(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenServiceAccountActasResourcesSlice flattens the contents of ServiceAccountActasResources from a JSON
// response object.
func flattenServiceAccountActasResourcesSlice(c *Client, i interface{}, res *ServiceAccount) []ServiceAccountActasResources {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceAccountActasResources{}
	}

	if len(a) == 0 {
		return []ServiceAccountActasResources{}
	}

	items := make([]ServiceAccountActasResources, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceAccountActasResources(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandServiceAccountActasResources expands an instance of ServiceAccountActasResources into a JSON
// request object.
func expandServiceAccountActasResources(c *Client, f *ServiceAccountActasResources, res *ServiceAccount) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandServiceAccountActasResourcesResourcesSlice(c, f.Resources, res); err != nil {
		return nil, fmt.Errorf("error expanding Resources into resources: %w", err)
	} else if v != nil {
		m["resources"] = v
	}

	return m, nil
}

// flattenServiceAccountActasResources flattens an instance of ServiceAccountActasResources from a JSON
// response object.
func flattenServiceAccountActasResources(c *Client, i interface{}, res *ServiceAccount) *ServiceAccountActasResources {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceAccountActasResources{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceAccountActasResources
	}
	r.Resources = flattenServiceAccountActasResourcesResourcesSlice(c, m["resources"], res)

	return r
}

// expandServiceAccountActasResourcesResourcesMap expands the contents of ServiceAccountActasResourcesResources into a JSON
// request object.
func expandServiceAccountActasResourcesResourcesMap(c *Client, f map[string]ServiceAccountActasResourcesResources, res *ServiceAccount) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceAccountActasResourcesResources(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceAccountActasResourcesResourcesSlice expands the contents of ServiceAccountActasResourcesResources into a JSON
// request object.
func expandServiceAccountActasResourcesResourcesSlice(c *Client, f []ServiceAccountActasResourcesResources, res *ServiceAccount) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceAccountActasResourcesResources(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceAccountActasResourcesResourcesMap flattens the contents of ServiceAccountActasResourcesResources from a JSON
// response object.
func flattenServiceAccountActasResourcesResourcesMap(c *Client, i interface{}, res *ServiceAccount) map[string]ServiceAccountActasResourcesResources {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceAccountActasResourcesResources{}
	}

	if len(a) == 0 {
		return map[string]ServiceAccountActasResourcesResources{}
	}

	items := make(map[string]ServiceAccountActasResourcesResources)
	for k, item := range a {
		items[k] = *flattenServiceAccountActasResourcesResources(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenServiceAccountActasResourcesResourcesSlice flattens the contents of ServiceAccountActasResourcesResources from a JSON
// response object.
func flattenServiceAccountActasResourcesResourcesSlice(c *Client, i interface{}, res *ServiceAccount) []ServiceAccountActasResourcesResources {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceAccountActasResourcesResources{}
	}

	if len(a) == 0 {
		return []ServiceAccountActasResourcesResources{}
	}

	items := make([]ServiceAccountActasResourcesResources, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceAccountActasResourcesResources(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandServiceAccountActasResourcesResources expands an instance of ServiceAccountActasResourcesResources into a JSON
// request object.
func expandServiceAccountActasResourcesResources(c *Client, f *ServiceAccountActasResourcesResources, res *ServiceAccount) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.FullResourceName; !dcl.IsEmptyValueIndirect(v) {
		m["fullResourceName"] = v
	}

	return m, nil
}

// flattenServiceAccountActasResourcesResources flattens an instance of ServiceAccountActasResourcesResources from a JSON
// response object.
func flattenServiceAccountActasResourcesResources(c *Client, i interface{}, res *ServiceAccount) *ServiceAccountActasResourcesResources {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceAccountActasResourcesResources{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceAccountActasResourcesResources
	}
	r.FullResourceName = dcl.FlattenString(m["fullResourceName"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *ServiceAccount) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalServiceAccount(b, c, r)
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

type serviceAccountDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         serviceAccountApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToServiceAccountDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]serviceAccountDiff, error) {
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
	var diffs []serviceAccountDiff
	// For each operation name, create a serviceAccountDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := serviceAccountDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToServiceAccountApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToServiceAccountApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (serviceAccountApiOperation, error) {
	switch opName {

	case "updateServiceAccountPatchServiceAccountOperation":
		return &updateServiceAccountPatchServiceAccountOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractServiceAccountFields(r *ServiceAccount) error {
	vActasResources := r.ActasResources
	if vActasResources == nil {
		// note: explicitly not the empty object.
		vActasResources = &ServiceAccountActasResources{}
	}
	if err := extractServiceAccountActasResourcesFields(r, vActasResources); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vActasResources) {
		r.ActasResources = vActasResources
	}
	return nil
}
func extractServiceAccountActasResourcesFields(r *ServiceAccount, o *ServiceAccountActasResources) error {
	return nil
}
func extractServiceAccountActasResourcesResourcesFields(r *ServiceAccount, o *ServiceAccountActasResourcesResources) error {
	return nil
}

func postReadExtractServiceAccountFields(r *ServiceAccount) error {
	vActasResources := r.ActasResources
	if vActasResources == nil {
		// note: explicitly not the empty object.
		vActasResources = &ServiceAccountActasResources{}
	}
	if err := postReadExtractServiceAccountActasResourcesFields(r, vActasResources); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vActasResources) {
		r.ActasResources = vActasResources
	}
	return nil
}
func postReadExtractServiceAccountActasResourcesFields(r *ServiceAccount, o *ServiceAccountActasResources) error {
	return nil
}
func postReadExtractServiceAccountActasResourcesResourcesFields(r *ServiceAccount, o *ServiceAccountActasResourcesResources) error {
	return nil
}
