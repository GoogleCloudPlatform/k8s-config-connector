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

func (r *WorkloadIdentityPoolProvider) validate() error {

	if err := dcl.ValidateExactlyOneOfFieldsSet([]string{"Aws", "Oidc"}, r.Aws, r.Oidc); err != nil {
		return err
	}
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.WorkloadIdentityPool, "WorkloadIdentityPool"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Aws) {
		if err := r.Aws.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Oidc) {
		if err := r.Oidc.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *WorkloadIdentityPoolProviderAws) validate() error {
	if err := dcl.Required(r, "accountId"); err != nil {
		return err
	}
	return nil
}
func (r *WorkloadIdentityPoolProviderOidc) validate() error {
	if err := dcl.Required(r, "issuerUri"); err != nil {
		return err
	}
	return nil
}
func (r *WorkloadIdentityPoolProvider) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://iam.googleapis.com/v1/", params)
}

func (r *WorkloadIdentityPoolProvider) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":              dcl.ValueOrEmptyString(nr.Project),
		"location":             dcl.ValueOrEmptyString(nr.Location),
		"workloadIdentityPool": dcl.ValueOrEmptyString(nr.WorkloadIdentityPool),
		"name":                 dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/workloadIdentityPools/{{workloadIdentityPool}}/providers/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *WorkloadIdentityPoolProvider) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":              dcl.ValueOrEmptyString(nr.Project),
		"location":             dcl.ValueOrEmptyString(nr.Location),
		"workloadIdentityPool": dcl.ValueOrEmptyString(nr.WorkloadIdentityPool),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/workloadIdentityPools/{{workloadIdentityPool}}/providers", nr.basePath(), userBasePath, params), nil

}

func (r *WorkloadIdentityPoolProvider) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":              dcl.ValueOrEmptyString(nr.Project),
		"location":             dcl.ValueOrEmptyString(nr.Location),
		"workloadIdentityPool": dcl.ValueOrEmptyString(nr.WorkloadIdentityPool),
		"name":                 dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/workloadIdentityPools/{{workloadIdentityPool}}/providers?workloadIdentityPoolProviderId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *WorkloadIdentityPoolProvider) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":              dcl.ValueOrEmptyString(nr.Project),
		"location":             dcl.ValueOrEmptyString(nr.Location),
		"workloadIdentityPool": dcl.ValueOrEmptyString(nr.WorkloadIdentityPool),
		"name":                 dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/workloadIdentityPools/{{workloadIdentityPool}}/providers/{{name}}", nr.basePath(), userBasePath, params), nil
}

// workloadIdentityPoolProviderApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type workloadIdentityPoolProviderApiOperation interface {
	do(context.Context, *WorkloadIdentityPoolProvider, *Client) error
}

// newUpdateWorkloadIdentityPoolProviderUpdateWorkloadIdentityPoolProviderRequest creates a request for an
// WorkloadIdentityPoolProvider resource's UpdateWorkloadIdentityPoolProvider update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateWorkloadIdentityPoolProviderUpdateWorkloadIdentityPoolProviderRequest(ctx context.Context, f *WorkloadIdentityPoolProvider, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Disabled; !dcl.IsEmptyValueIndirect(v) {
		req["disabled"] = v
	}
	if v := f.AttributeMapping; !dcl.IsEmptyValueIndirect(v) {
		req["attributeMapping"] = v
	}
	if v := f.AttributeCondition; !dcl.IsEmptyValueIndirect(v) {
		req["attributeCondition"] = v
	}
	if v, err := expandWorkloadIdentityPoolProviderAws(c, f.Aws, res); err != nil {
		return nil, fmt.Errorf("error expanding Aws into aws: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["aws"] = v
	}
	if v, err := expandWorkloadIdentityPoolProviderOidc(c, f.Oidc, res); err != nil {
		return nil, fmt.Errorf("error expanding Oidc into oidc: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["oidc"] = v
	}
	return req, nil
}

// marshalUpdateWorkloadIdentityPoolProviderUpdateWorkloadIdentityPoolProviderRequest converts the update into
// the final JSON request body.
func marshalUpdateWorkloadIdentityPoolProviderUpdateWorkloadIdentityPoolProviderRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateWorkloadIdentityPoolProviderUpdateWorkloadIdentityPoolProviderOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateWorkloadIdentityPoolProviderUpdateWorkloadIdentityPoolProviderOperation) do(ctx context.Context, r *WorkloadIdentityPoolProvider, c *Client) error {
	_, err := c.GetWorkloadIdentityPoolProvider(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateWorkloadIdentityPoolProvider")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateWorkloadIdentityPoolProviderUpdateWorkloadIdentityPoolProviderRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateWorkloadIdentityPoolProviderUpdateWorkloadIdentityPoolProviderRequest(c, req)
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

func (c *Client) listWorkloadIdentityPoolProviderRaw(ctx context.Context, r *WorkloadIdentityPoolProvider, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != WorkloadIdentityPoolProviderMaxPage {
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

type listWorkloadIdentityPoolProviderOperation struct {
	WorkloadIdentityPoolProviders []map[string]interface{} `json:"workloadIdentityPoolProviders"`
	Token                         string                   `json:"nextPageToken"`
}

func (c *Client) listWorkloadIdentityPoolProvider(ctx context.Context, r *WorkloadIdentityPoolProvider, pageToken string, pageSize int32) ([]*WorkloadIdentityPoolProvider, string, error) {
	b, err := c.listWorkloadIdentityPoolProviderRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listWorkloadIdentityPoolProviderOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*WorkloadIdentityPoolProvider
	for _, v := range m.WorkloadIdentityPoolProviders {
		res, err := unmarshalMapWorkloadIdentityPoolProvider(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		res.WorkloadIdentityPool = r.WorkloadIdentityPool
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllWorkloadIdentityPoolProvider(ctx context.Context, f func(*WorkloadIdentityPoolProvider) bool, resources []*WorkloadIdentityPoolProvider) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteWorkloadIdentityPoolProvider(ctx, res)
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

type deleteWorkloadIdentityPoolProviderOperation struct{}

func (op *deleteWorkloadIdentityPoolProviderOperation) do(ctx context.Context, r *WorkloadIdentityPoolProvider, c *Client) error {
	r, err := c.GetWorkloadIdentityPoolProvider(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "WorkloadIdentityPoolProvider not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetWorkloadIdentityPoolProvider checking for existence. error: %v", err)
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
		return fmt.Errorf("failed to delete WorkloadIdentityPoolProvider: %w", err)
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createWorkloadIdentityPoolProviderOperation struct {
	response map[string]interface{}
}

func (op *createWorkloadIdentityPoolProviderOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createWorkloadIdentityPoolProviderOperation) do(ctx context.Context, r *WorkloadIdentityPoolProvider, c *Client) error {
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

	if _, err := c.GetWorkloadIdentityPoolProvider(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getWorkloadIdentityPoolProviderRaw(ctx context.Context, r *WorkloadIdentityPoolProvider) ([]byte, error) {

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

func (c *Client) workloadIdentityPoolProviderDiffsForRawDesired(ctx context.Context, rawDesired *WorkloadIdentityPoolProvider, opts ...dcl.ApplyOption) (initial, desired *WorkloadIdentityPoolProvider, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *WorkloadIdentityPoolProvider
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*WorkloadIdentityPoolProvider); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected WorkloadIdentityPoolProvider, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetWorkloadIdentityPoolProvider(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a WorkloadIdentityPoolProvider resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve WorkloadIdentityPoolProvider resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that WorkloadIdentityPoolProvider resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeWorkloadIdentityPoolProviderDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for WorkloadIdentityPoolProvider: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for WorkloadIdentityPoolProvider: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractWorkloadIdentityPoolProviderFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeWorkloadIdentityPoolProviderInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for WorkloadIdentityPoolProvider: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeWorkloadIdentityPoolProviderDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for WorkloadIdentityPoolProvider: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffWorkloadIdentityPoolProvider(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeWorkloadIdentityPoolProviderInitialState(rawInitial, rawDesired *WorkloadIdentityPoolProvider) (*WorkloadIdentityPoolProvider, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.

	if !dcl.IsZeroValue(rawInitial.Aws) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Oidc) {
			rawInitial.Aws = EmptyWorkloadIdentityPoolProviderAws
		}
	}

	if !dcl.IsZeroValue(rawInitial.Oidc) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Aws) {
			rawInitial.Oidc = EmptyWorkloadIdentityPoolProviderOidc
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

func canonicalizeWorkloadIdentityPoolProviderDesiredState(rawDesired, rawInitial *WorkloadIdentityPoolProvider, opts ...dcl.ApplyOption) (*WorkloadIdentityPoolProvider, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Aws = canonicalizeWorkloadIdentityPoolProviderAws(rawDesired.Aws, nil, opts...)
		rawDesired.Oidc = canonicalizeWorkloadIdentityPoolProviderOidc(rawDesired.Oidc, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &WorkloadIdentityPoolProvider{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
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
	if dcl.BoolCanonicalize(rawDesired.Disabled, rawInitial.Disabled) {
		canonicalDesired.Disabled = rawInitial.Disabled
	} else {
		canonicalDesired.Disabled = rawDesired.Disabled
	}
	if dcl.IsZeroValue(rawDesired.AttributeMapping) || (dcl.IsEmptyValueIndirect(rawDesired.AttributeMapping) && dcl.IsEmptyValueIndirect(rawInitial.AttributeMapping)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.AttributeMapping = rawInitial.AttributeMapping
	} else {
		canonicalDesired.AttributeMapping = rawDesired.AttributeMapping
	}
	if dcl.StringCanonicalize(rawDesired.AttributeCondition, rawInitial.AttributeCondition) {
		canonicalDesired.AttributeCondition = rawInitial.AttributeCondition
	} else {
		canonicalDesired.AttributeCondition = rawDesired.AttributeCondition
	}
	canonicalDesired.Aws = canonicalizeWorkloadIdentityPoolProviderAws(rawDesired.Aws, rawInitial.Aws, opts...)
	canonicalDesired.Oidc = canonicalizeWorkloadIdentityPoolProviderOidc(rawDesired.Oidc, rawInitial.Oidc, opts...)
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
	if dcl.NameToSelfLink(rawDesired.WorkloadIdentityPool, rawInitial.WorkloadIdentityPool) {
		canonicalDesired.WorkloadIdentityPool = rawInitial.WorkloadIdentityPool
	} else {
		canonicalDesired.WorkloadIdentityPool = rawDesired.WorkloadIdentityPool
	}

	if canonicalDesired.Aws != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Oidc) {
			canonicalDesired.Aws = EmptyWorkloadIdentityPoolProviderAws
		}
	}

	if canonicalDesired.Oidc != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Aws) {
			canonicalDesired.Oidc = EmptyWorkloadIdentityPoolProviderOidc
		}
	}

	return canonicalDesired, nil
}

func canonicalizeWorkloadIdentityPoolProviderNewState(c *Client, rawNew, rawDesired *WorkloadIdentityPoolProvider) (*WorkloadIdentityPoolProvider, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
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

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Disabled) && dcl.IsEmptyValueIndirect(rawDesired.Disabled) {
		rawNew.Disabled = rawDesired.Disabled
	} else {
		if dcl.BoolCanonicalize(rawDesired.Disabled, rawNew.Disabled) {
			rawNew.Disabled = rawDesired.Disabled
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.AttributeMapping) && dcl.IsEmptyValueIndirect(rawDesired.AttributeMapping) {
		rawNew.AttributeMapping = rawDesired.AttributeMapping
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.AttributeCondition) && dcl.IsEmptyValueIndirect(rawDesired.AttributeCondition) {
		rawNew.AttributeCondition = rawDesired.AttributeCondition
	} else {
		if dcl.StringCanonicalize(rawDesired.AttributeCondition, rawNew.AttributeCondition) {
			rawNew.AttributeCondition = rawDesired.AttributeCondition
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Aws) && dcl.IsEmptyValueIndirect(rawDesired.Aws) {
		rawNew.Aws = rawDesired.Aws
	} else {
		rawNew.Aws = canonicalizeNewWorkloadIdentityPoolProviderAws(c, rawDesired.Aws, rawNew.Aws)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Oidc) && dcl.IsEmptyValueIndirect(rawDesired.Oidc) {
		rawNew.Oidc = rawDesired.Oidc
	} else {
		rawNew.Oidc = canonicalizeNewWorkloadIdentityPoolProviderOidc(c, rawDesired.Oidc, rawNew.Oidc)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	rawNew.WorkloadIdentityPool = rawDesired.WorkloadIdentityPool

	return rawNew, nil
}

func canonicalizeWorkloadIdentityPoolProviderAws(des, initial *WorkloadIdentityPoolProviderAws, opts ...dcl.ApplyOption) *WorkloadIdentityPoolProviderAws {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &WorkloadIdentityPoolProviderAws{}

	if dcl.StringCanonicalize(des.AccountId, initial.AccountId) || dcl.IsZeroValue(des.AccountId) {
		cDes.AccountId = initial.AccountId
	} else {
		cDes.AccountId = des.AccountId
	}
	if dcl.StringArrayCanonicalize(des.StsUri, initial.StsUri) {
		cDes.StsUri = initial.StsUri
	} else {
		cDes.StsUri = des.StsUri
	}

	return cDes
}

func canonicalizeWorkloadIdentityPoolProviderAwsSlice(des, initial []WorkloadIdentityPoolProviderAws, opts ...dcl.ApplyOption) []WorkloadIdentityPoolProviderAws {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]WorkloadIdentityPoolProviderAws, 0, len(des))
		for _, d := range des {
			cd := canonicalizeWorkloadIdentityPoolProviderAws(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]WorkloadIdentityPoolProviderAws, 0, len(des))
	for i, d := range des {
		cd := canonicalizeWorkloadIdentityPoolProviderAws(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewWorkloadIdentityPoolProviderAws(c *Client, des, nw *WorkloadIdentityPoolProviderAws) *WorkloadIdentityPoolProviderAws {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for WorkloadIdentityPoolProviderAws while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.AccountId, nw.AccountId) {
		nw.AccountId = des.AccountId
	}
	nw.StsUri = des.StsUri

	return nw
}

func canonicalizeNewWorkloadIdentityPoolProviderAwsSet(c *Client, des, nw []WorkloadIdentityPoolProviderAws) []WorkloadIdentityPoolProviderAws {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []WorkloadIdentityPoolProviderAws
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareWorkloadIdentityPoolProviderAwsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewWorkloadIdentityPoolProviderAws(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewWorkloadIdentityPoolProviderAwsSlice(c *Client, des, nw []WorkloadIdentityPoolProviderAws) []WorkloadIdentityPoolProviderAws {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []WorkloadIdentityPoolProviderAws
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkloadIdentityPoolProviderAws(c, &d, &n))
	}

	return items
}

func canonicalizeWorkloadIdentityPoolProviderOidc(des, initial *WorkloadIdentityPoolProviderOidc, opts ...dcl.ApplyOption) *WorkloadIdentityPoolProviderOidc {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &WorkloadIdentityPoolProviderOidc{}

	if dcl.StringCanonicalize(des.IssuerUri, initial.IssuerUri) || dcl.IsZeroValue(des.IssuerUri) {
		cDes.IssuerUri = initial.IssuerUri
	} else {
		cDes.IssuerUri = des.IssuerUri
	}
	if dcl.StringArrayCanonicalize(des.AllowedAudiences, initial.AllowedAudiences) {
		cDes.AllowedAudiences = initial.AllowedAudiences
	} else {
		cDes.AllowedAudiences = des.AllowedAudiences
	}

	return cDes
}

func canonicalizeWorkloadIdentityPoolProviderOidcSlice(des, initial []WorkloadIdentityPoolProviderOidc, opts ...dcl.ApplyOption) []WorkloadIdentityPoolProviderOidc {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]WorkloadIdentityPoolProviderOidc, 0, len(des))
		for _, d := range des {
			cd := canonicalizeWorkloadIdentityPoolProviderOidc(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]WorkloadIdentityPoolProviderOidc, 0, len(des))
	for i, d := range des {
		cd := canonicalizeWorkloadIdentityPoolProviderOidc(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewWorkloadIdentityPoolProviderOidc(c *Client, des, nw *WorkloadIdentityPoolProviderOidc) *WorkloadIdentityPoolProviderOidc {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for WorkloadIdentityPoolProviderOidc while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.IssuerUri, nw.IssuerUri) {
		nw.IssuerUri = des.IssuerUri
	}
	if dcl.StringArrayCanonicalize(des.AllowedAudiences, nw.AllowedAudiences) {
		nw.AllowedAudiences = des.AllowedAudiences
	}

	return nw
}

func canonicalizeNewWorkloadIdentityPoolProviderOidcSet(c *Client, des, nw []WorkloadIdentityPoolProviderOidc) []WorkloadIdentityPoolProviderOidc {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []WorkloadIdentityPoolProviderOidc
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareWorkloadIdentityPoolProviderOidcNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewWorkloadIdentityPoolProviderOidc(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewWorkloadIdentityPoolProviderOidcSlice(c *Client, des, nw []WorkloadIdentityPoolProviderOidc) []WorkloadIdentityPoolProviderOidc {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []WorkloadIdentityPoolProviderOidc
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkloadIdentityPoolProviderOidc(c, &d, &n))
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
func diffWorkloadIdentityPoolProvider(c *Client, desired, actual *WorkloadIdentityPoolProvider, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateWorkloadIdentityPoolProviderUpdateWorkloadIdentityPoolProviderOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateWorkloadIdentityPoolProviderUpdateWorkloadIdentityPoolProviderOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Disabled, actual.Disabled, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateWorkloadIdentityPoolProviderUpdateWorkloadIdentityPoolProviderOperation")}, fn.AddNest("Disabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AttributeMapping, actual.AttributeMapping, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateWorkloadIdentityPoolProviderUpdateWorkloadIdentityPoolProviderOperation")}, fn.AddNest("AttributeMapping")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AttributeCondition, actual.AttributeCondition, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateWorkloadIdentityPoolProviderUpdateWorkloadIdentityPoolProviderOperation")}, fn.AddNest("AttributeCondition")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Aws, actual.Aws, dcl.DiffInfo{ObjectFunction: compareWorkloadIdentityPoolProviderAwsNewStyle, EmptyObject: EmptyWorkloadIdentityPoolProviderAws, OperationSelector: dcl.TriggersOperation("updateWorkloadIdentityPoolProviderUpdateWorkloadIdentityPoolProviderOperation")}, fn.AddNest("Aws")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Oidc, actual.Oidc, dcl.DiffInfo{ObjectFunction: compareWorkloadIdentityPoolProviderOidcNewStyle, EmptyObject: EmptyWorkloadIdentityPoolProviderOidc, OperationSelector: dcl.TriggersOperation("updateWorkloadIdentityPoolProviderUpdateWorkloadIdentityPoolProviderOperation")}, fn.AddNest("Oidc")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.WorkloadIdentityPool, actual.WorkloadIdentityPool, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("WorkloadIdentityPool")); len(ds) != 0 || err != nil {
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
func compareWorkloadIdentityPoolProviderAwsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*WorkloadIdentityPoolProviderAws)
	if !ok {
		desiredNotPointer, ok := d.(WorkloadIdentityPoolProviderAws)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a WorkloadIdentityPoolProviderAws or *WorkloadIdentityPoolProviderAws", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*WorkloadIdentityPoolProviderAws)
	if !ok {
		actualNotPointer, ok := a.(WorkloadIdentityPoolProviderAws)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a WorkloadIdentityPoolProviderAws", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AccountId, actual.AccountId, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateWorkloadIdentityPoolProviderUpdateWorkloadIdentityPoolProviderOperation")}, fn.AddNest("AccountId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StsUri, actual.StsUri, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateWorkloadIdentityPoolProviderUpdateWorkloadIdentityPoolProviderOperation")}, fn.AddNest("StsUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareWorkloadIdentityPoolProviderOidcNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*WorkloadIdentityPoolProviderOidc)
	if !ok {
		desiredNotPointer, ok := d.(WorkloadIdentityPoolProviderOidc)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a WorkloadIdentityPoolProviderOidc or *WorkloadIdentityPoolProviderOidc", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*WorkloadIdentityPoolProviderOidc)
	if !ok {
		actualNotPointer, ok := a.(WorkloadIdentityPoolProviderOidc)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a WorkloadIdentityPoolProviderOidc", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.IssuerUri, actual.IssuerUri, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateWorkloadIdentityPoolProviderUpdateWorkloadIdentityPoolProviderOperation")}, fn.AddNest("IssuerUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowedAudiences, actual.AllowedAudiences, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateWorkloadIdentityPoolProviderUpdateWorkloadIdentityPoolProviderOperation")}, fn.AddNest("AllowedAudiences")); len(ds) != 0 || err != nil {
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
func (r *WorkloadIdentityPoolProvider) urlNormalized() *WorkloadIdentityPoolProvider {
	normalized := dcl.Copy(*r).(WorkloadIdentityPoolProvider)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.AttributeCondition = dcl.SelfLinkToName(r.AttributeCondition)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.WorkloadIdentityPool = dcl.SelfLinkToName(r.WorkloadIdentityPool)
	return &normalized
}

func (r *WorkloadIdentityPoolProvider) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateWorkloadIdentityPoolProvider" {
		fields := map[string]interface{}{
			"project":              dcl.ValueOrEmptyString(nr.Project),
			"location":             dcl.ValueOrEmptyString(nr.Location),
			"workloadIdentityPool": dcl.ValueOrEmptyString(nr.WorkloadIdentityPool),
			"name":                 dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/workloadIdentityPools/{{workloadIdentityPool}}/providers/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the WorkloadIdentityPoolProvider resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *WorkloadIdentityPoolProvider) marshal(c *Client) ([]byte, error) {
	m, err := expandWorkloadIdentityPoolProvider(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling WorkloadIdentityPoolProvider: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalWorkloadIdentityPoolProvider decodes JSON responses into the WorkloadIdentityPoolProvider resource schema.
func unmarshalWorkloadIdentityPoolProvider(b []byte, c *Client, res *WorkloadIdentityPoolProvider) (*WorkloadIdentityPoolProvider, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapWorkloadIdentityPoolProvider(m, c, res)
}

func unmarshalMapWorkloadIdentityPoolProvider(m map[string]interface{}, c *Client, res *WorkloadIdentityPoolProvider) (*WorkloadIdentityPoolProvider, error) {

	flattened := flattenWorkloadIdentityPoolProvider(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandWorkloadIdentityPoolProvider expands WorkloadIdentityPoolProvider into a JSON request object.
func expandWorkloadIdentityPoolProvider(c *Client, f *WorkloadIdentityPoolProvider) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("projects/%s/locations/%s/workloadIdentityPools/%s/providers/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Location), dcl.SelfLinkToName(f.WorkloadIdentityPool), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.DisplayName; dcl.ValueShouldBeSent(v) {
		m["displayName"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v := f.Disabled; dcl.ValueShouldBeSent(v) {
		m["disabled"] = v
	}
	if v := f.AttributeMapping; dcl.ValueShouldBeSent(v) {
		m["attributeMapping"] = v
	}
	if v := f.AttributeCondition; dcl.ValueShouldBeSent(v) {
		m["attributeCondition"] = v
	}
	if v, err := expandWorkloadIdentityPoolProviderAws(c, f.Aws, res); err != nil {
		return nil, fmt.Errorf("error expanding Aws into aws: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["aws"] = v
	}
	if v, err := expandWorkloadIdentityPoolProviderOidc(c, f.Oidc, res); err != nil {
		return nil, fmt.Errorf("error expanding Oidc into oidc: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["oidc"] = v
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
		return nil, fmt.Errorf("error expanding WorkloadIdentityPool into workloadIdentityPool: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["workloadIdentityPool"] = v
	}

	return m, nil
}

// flattenWorkloadIdentityPoolProvider flattens WorkloadIdentityPoolProvider from a JSON request object into the
// WorkloadIdentityPoolProvider type.
func flattenWorkloadIdentityPoolProvider(c *Client, i interface{}, res *WorkloadIdentityPoolProvider) *WorkloadIdentityPoolProvider {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &WorkloadIdentityPoolProvider{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.DisplayName = dcl.FlattenString(m["displayName"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.State = flattenWorkloadIdentityPoolProviderStateEnum(m["state"])
	resultRes.Disabled = dcl.FlattenBool(m["disabled"])
	resultRes.AttributeMapping = dcl.FlattenKeyValuePairs(m["attributeMapping"])
	resultRes.AttributeCondition = dcl.FlattenString(m["attributeCondition"])
	resultRes.Aws = flattenWorkloadIdentityPoolProviderAws(c, m["aws"], res)
	resultRes.Oidc = flattenWorkloadIdentityPoolProviderOidc(c, m["oidc"], res)
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])
	resultRes.WorkloadIdentityPool = dcl.FlattenString(m["workloadIdentityPool"])

	return resultRes
}

// expandWorkloadIdentityPoolProviderAwsMap expands the contents of WorkloadIdentityPoolProviderAws into a JSON
// request object.
func expandWorkloadIdentityPoolProviderAwsMap(c *Client, f map[string]WorkloadIdentityPoolProviderAws, res *WorkloadIdentityPoolProvider) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkloadIdentityPoolProviderAws(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkloadIdentityPoolProviderAwsSlice expands the contents of WorkloadIdentityPoolProviderAws into a JSON
// request object.
func expandWorkloadIdentityPoolProviderAwsSlice(c *Client, f []WorkloadIdentityPoolProviderAws, res *WorkloadIdentityPoolProvider) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkloadIdentityPoolProviderAws(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkloadIdentityPoolProviderAwsMap flattens the contents of WorkloadIdentityPoolProviderAws from a JSON
// response object.
func flattenWorkloadIdentityPoolProviderAwsMap(c *Client, i interface{}, res *WorkloadIdentityPoolProvider) map[string]WorkloadIdentityPoolProviderAws {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkloadIdentityPoolProviderAws{}
	}

	if len(a) == 0 {
		return map[string]WorkloadIdentityPoolProviderAws{}
	}

	items := make(map[string]WorkloadIdentityPoolProviderAws)
	for k, item := range a {
		items[k] = *flattenWorkloadIdentityPoolProviderAws(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenWorkloadIdentityPoolProviderAwsSlice flattens the contents of WorkloadIdentityPoolProviderAws from a JSON
// response object.
func flattenWorkloadIdentityPoolProviderAwsSlice(c *Client, i interface{}, res *WorkloadIdentityPoolProvider) []WorkloadIdentityPoolProviderAws {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkloadIdentityPoolProviderAws{}
	}

	if len(a) == 0 {
		return []WorkloadIdentityPoolProviderAws{}
	}

	items := make([]WorkloadIdentityPoolProviderAws, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkloadIdentityPoolProviderAws(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandWorkloadIdentityPoolProviderAws expands an instance of WorkloadIdentityPoolProviderAws into a JSON
// request object.
func expandWorkloadIdentityPoolProviderAws(c *Client, f *WorkloadIdentityPoolProviderAws, res *WorkloadIdentityPoolProvider) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AccountId; !dcl.IsEmptyValueIndirect(v) {
		m["accountId"] = v
	}
	if v := f.StsUri; v != nil {
		m["stsUri"] = v
	}

	return m, nil
}

// flattenWorkloadIdentityPoolProviderAws flattens an instance of WorkloadIdentityPoolProviderAws from a JSON
// response object.
func flattenWorkloadIdentityPoolProviderAws(c *Client, i interface{}, res *WorkloadIdentityPoolProvider) *WorkloadIdentityPoolProviderAws {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkloadIdentityPoolProviderAws{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyWorkloadIdentityPoolProviderAws
	}
	r.AccountId = dcl.FlattenString(m["accountId"])
	r.StsUri = dcl.FlattenStringSlice(m["stsUri"])

	return r
}

// expandWorkloadIdentityPoolProviderOidcMap expands the contents of WorkloadIdentityPoolProviderOidc into a JSON
// request object.
func expandWorkloadIdentityPoolProviderOidcMap(c *Client, f map[string]WorkloadIdentityPoolProviderOidc, res *WorkloadIdentityPoolProvider) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkloadIdentityPoolProviderOidc(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkloadIdentityPoolProviderOidcSlice expands the contents of WorkloadIdentityPoolProviderOidc into a JSON
// request object.
func expandWorkloadIdentityPoolProviderOidcSlice(c *Client, f []WorkloadIdentityPoolProviderOidc, res *WorkloadIdentityPoolProvider) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkloadIdentityPoolProviderOidc(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkloadIdentityPoolProviderOidcMap flattens the contents of WorkloadIdentityPoolProviderOidc from a JSON
// response object.
func flattenWorkloadIdentityPoolProviderOidcMap(c *Client, i interface{}, res *WorkloadIdentityPoolProvider) map[string]WorkloadIdentityPoolProviderOidc {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkloadIdentityPoolProviderOidc{}
	}

	if len(a) == 0 {
		return map[string]WorkloadIdentityPoolProviderOidc{}
	}

	items := make(map[string]WorkloadIdentityPoolProviderOidc)
	for k, item := range a {
		items[k] = *flattenWorkloadIdentityPoolProviderOidc(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenWorkloadIdentityPoolProviderOidcSlice flattens the contents of WorkloadIdentityPoolProviderOidc from a JSON
// response object.
func flattenWorkloadIdentityPoolProviderOidcSlice(c *Client, i interface{}, res *WorkloadIdentityPoolProvider) []WorkloadIdentityPoolProviderOidc {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkloadIdentityPoolProviderOidc{}
	}

	if len(a) == 0 {
		return []WorkloadIdentityPoolProviderOidc{}
	}

	items := make([]WorkloadIdentityPoolProviderOidc, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkloadIdentityPoolProviderOidc(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandWorkloadIdentityPoolProviderOidc expands an instance of WorkloadIdentityPoolProviderOidc into a JSON
// request object.
func expandWorkloadIdentityPoolProviderOidc(c *Client, f *WorkloadIdentityPoolProviderOidc, res *WorkloadIdentityPoolProvider) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.IssuerUri; !dcl.IsEmptyValueIndirect(v) {
		m["issuerUri"] = v
	}
	if v := f.AllowedAudiences; v != nil {
		m["allowedAudiences"] = v
	}

	return m, nil
}

// flattenWorkloadIdentityPoolProviderOidc flattens an instance of WorkloadIdentityPoolProviderOidc from a JSON
// response object.
func flattenWorkloadIdentityPoolProviderOidc(c *Client, i interface{}, res *WorkloadIdentityPoolProvider) *WorkloadIdentityPoolProviderOidc {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkloadIdentityPoolProviderOidc{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyWorkloadIdentityPoolProviderOidc
	}
	r.IssuerUri = dcl.FlattenString(m["issuerUri"])
	r.AllowedAudiences = dcl.FlattenStringSlice(m["allowedAudiences"])

	return r
}

// flattenWorkloadIdentityPoolProviderStateEnumMap flattens the contents of WorkloadIdentityPoolProviderStateEnum from a JSON
// response object.
func flattenWorkloadIdentityPoolProviderStateEnumMap(c *Client, i interface{}, res *WorkloadIdentityPoolProvider) map[string]WorkloadIdentityPoolProviderStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkloadIdentityPoolProviderStateEnum{}
	}

	if len(a) == 0 {
		return map[string]WorkloadIdentityPoolProviderStateEnum{}
	}

	items := make(map[string]WorkloadIdentityPoolProviderStateEnum)
	for k, item := range a {
		items[k] = *flattenWorkloadIdentityPoolProviderStateEnum(item.(interface{}))
	}

	return items
}

// flattenWorkloadIdentityPoolProviderStateEnumSlice flattens the contents of WorkloadIdentityPoolProviderStateEnum from a JSON
// response object.
func flattenWorkloadIdentityPoolProviderStateEnumSlice(c *Client, i interface{}, res *WorkloadIdentityPoolProvider) []WorkloadIdentityPoolProviderStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkloadIdentityPoolProviderStateEnum{}
	}

	if len(a) == 0 {
		return []WorkloadIdentityPoolProviderStateEnum{}
	}

	items := make([]WorkloadIdentityPoolProviderStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkloadIdentityPoolProviderStateEnum(item.(interface{})))
	}

	return items
}

// flattenWorkloadIdentityPoolProviderStateEnum asserts that an interface is a string, and returns a
// pointer to a *WorkloadIdentityPoolProviderStateEnum with the same value as that string.
func flattenWorkloadIdentityPoolProviderStateEnum(i interface{}) *WorkloadIdentityPoolProviderStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return WorkloadIdentityPoolProviderStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *WorkloadIdentityPoolProvider) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalWorkloadIdentityPoolProvider(b, c, r)
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
		if nr.WorkloadIdentityPool == nil && ncr.WorkloadIdentityPool == nil {
			c.Config.Logger.Info("Both WorkloadIdentityPool fields null - considering equal.")
		} else if nr.WorkloadIdentityPool == nil || ncr.WorkloadIdentityPool == nil {
			c.Config.Logger.Info("Only one WorkloadIdentityPool field is null - considering unequal.")
			return false
		} else if *nr.WorkloadIdentityPool != *ncr.WorkloadIdentityPool {
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

type workloadIdentityPoolProviderDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         workloadIdentityPoolProviderApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToWorkloadIdentityPoolProviderDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]workloadIdentityPoolProviderDiff, error) {
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
	var diffs []workloadIdentityPoolProviderDiff
	// For each operation name, create a workloadIdentityPoolProviderDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := workloadIdentityPoolProviderDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToWorkloadIdentityPoolProviderApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToWorkloadIdentityPoolProviderApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (workloadIdentityPoolProviderApiOperation, error) {
	switch opName {

	case "updateWorkloadIdentityPoolProviderUpdateWorkloadIdentityPoolProviderOperation":
		return &updateWorkloadIdentityPoolProviderUpdateWorkloadIdentityPoolProviderOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractWorkloadIdentityPoolProviderFields(r *WorkloadIdentityPoolProvider) error {
	vAws := r.Aws
	if vAws == nil {
		// note: explicitly not the empty object.
		vAws = &WorkloadIdentityPoolProviderAws{}
	}
	if err := extractWorkloadIdentityPoolProviderAwsFields(r, vAws); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAws) {
		r.Aws = vAws
	}
	vOidc := r.Oidc
	if vOidc == nil {
		// note: explicitly not the empty object.
		vOidc = &WorkloadIdentityPoolProviderOidc{}
	}
	if err := extractWorkloadIdentityPoolProviderOidcFields(r, vOidc); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vOidc) {
		r.Oidc = vOidc
	}
	return nil
}
func extractWorkloadIdentityPoolProviderAwsFields(r *WorkloadIdentityPoolProvider, o *WorkloadIdentityPoolProviderAws) error {
	return nil
}
func extractWorkloadIdentityPoolProviderOidcFields(r *WorkloadIdentityPoolProvider, o *WorkloadIdentityPoolProviderOidc) error {
	return nil
}

func postReadExtractWorkloadIdentityPoolProviderFields(r *WorkloadIdentityPoolProvider) error {
	vAws := r.Aws
	if vAws == nil {
		// note: explicitly not the empty object.
		vAws = &WorkloadIdentityPoolProviderAws{}
	}
	if err := postReadExtractWorkloadIdentityPoolProviderAwsFields(r, vAws); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAws) {
		r.Aws = vAws
	}
	vOidc := r.Oidc
	if vOidc == nil {
		// note: explicitly not the empty object.
		vOidc = &WorkloadIdentityPoolProviderOidc{}
	}
	if err := postReadExtractWorkloadIdentityPoolProviderOidcFields(r, vOidc); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vOidc) {
		r.Oidc = vOidc
	}
	return nil
}
func postReadExtractWorkloadIdentityPoolProviderAwsFields(r *WorkloadIdentityPoolProvider, o *WorkloadIdentityPoolProviderAws) error {
	return nil
}
func postReadExtractWorkloadIdentityPoolProviderOidcFields(r *WorkloadIdentityPoolProvider, o *WorkloadIdentityPoolProviderOidc) error {
	return nil
}
