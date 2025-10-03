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

func (r *Function) validate() error {

	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"SourceArchiveUrl", "SourceRepository"}, r.SourceArchiveUrl, r.SourceRepository); err != nil {
		return err
	}
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"EventTrigger", "HttpsTrigger"}, r.EventTrigger, r.HttpsTrigger); err != nil {
		return err
	}
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "runtime"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Region, "Region"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.SourceRepository) {
		if err := r.SourceRepository.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.HttpsTrigger) {
		if err := r.HttpsTrigger.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.EventTrigger) {
		if err := r.EventTrigger.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FunctionSourceRepository) validate() error {
	if err := dcl.Required(r, "url"); err != nil {
		return err
	}
	return nil
}
func (r *FunctionHttpsTrigger) validate() error {
	return nil
}
func (r *FunctionEventTrigger) validate() error {
	if err := dcl.Required(r, "eventType"); err != nil {
		return err
	}
	if err := dcl.Required(r, "resource"); err != nil {
		return err
	}
	return nil
}
func (r *Function) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://cloudfunctions.googleapis.com/v1/", params)
}

func (r *Function) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"region":  dcl.ValueOrEmptyString(nr.Region),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{region}}/functions/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Function) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"region":  dcl.ValueOrEmptyString(nr.Region),
	}
	return dcl.URL("projects/{{project}}/locations/{{region}}/functions", nr.basePath(), userBasePath, params), nil

}

func (r *Function) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"region":  dcl.ValueOrEmptyString(nr.Region),
	}
	return dcl.URL("projects/{{project}}/locations/{{region}}/functions", nr.basePath(), userBasePath, params), nil

}

func (r *Function) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"region":  dcl.ValueOrEmptyString(nr.Region),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{region}}/functions/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Function) SetPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"project": *nr.Project,
		"region":  *nr.Region,
		"name":    *nr.Name,
	}
	return dcl.URL("projects/{{project}}/locations/{{region}}/functions/{{name}}:setIamPolicy", nr.basePath(), userBasePath, fields)
}

func (r *Function) SetPolicyVerb() string {
	return "POST"
}

func (r *Function) getPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"project": *nr.Project,
		"region":  *nr.Region,
		"name":    *nr.Name,
	}
	return dcl.URL("projects/{{project}}/locations/{{region}}/functions/{{name}}:getIamPolicy", nr.basePath(), userBasePath, fields)
}

func (r *Function) IAMPolicyVersion() int {
	return 3
}

// functionApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type functionApiOperation interface {
	do(context.Context, *Function, *Client) error
}

// newUpdateFunctionUpdateRequest creates a request for an
// Function resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateFunctionUpdateRequest(ctx context.Context, f *Function, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Runtime; !dcl.IsEmptyValueIndirect(v) {
		req["runtime"] = v
	}
	if v := f.Timeout; !dcl.IsEmptyValueIndirect(v) {
		req["timeout"] = v
	}
	if v := f.AvailableMemoryMb; !dcl.IsEmptyValueIndirect(v) {
		req["availableMemoryMb"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v := f.EnvironmentVariables; !dcl.IsEmptyValueIndirect(v) {
		req["environmentVariables"] = v
	}
	if v := f.MaxInstances; !dcl.IsEmptyValueIndirect(v) {
		req["maxInstances"] = v
	}
	if v := f.VPCConnector; !dcl.IsEmptyValueIndirect(v) {
		req["vpcConnector"] = v
	}
	if v := f.VPCConnectorEgressSettings; !dcl.IsEmptyValueIndirect(v) {
		req["vpcConnectorEgressSettings"] = v
	}
	if v := f.IngressSettings; !dcl.IsEmptyValueIndirect(v) {
		req["ingressSettings"] = v
	}
	return req, nil
}

// marshalUpdateFunctionUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateFunctionUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateFunctionUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateFunctionUpdateOperation) do(ctx context.Context, r *Function, c *Client) error {
	_, err := c.GetFunction(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateFunctionUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateFunctionUpdateRequest(c, req)
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

func (c *Client) listFunctionRaw(ctx context.Context, r *Function, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != FunctionMaxPage {
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

type listFunctionOperation struct {
	Functions []map[string]interface{} `json:"functions"`
	Token     string                   `json:"nextPageToken"`
}

func (c *Client) listFunction(ctx context.Context, r *Function, pageToken string, pageSize int32) ([]*Function, string, error) {
	b, err := c.listFunctionRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listFunctionOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Function
	for _, v := range m.Functions {
		res, err := unmarshalMapFunction(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Region = r.Region
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllFunction(ctx context.Context, f func(*Function) bool, resources []*Function) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteFunction(ctx, res)
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

type deleteFunctionOperation struct{}

func (op *deleteFunctionOperation) do(ctx context.Context, r *Function, c *Client) error {
	r, err := c.GetFunction(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Function not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetFunction checking for existence. error: %v", err)
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
		_, err := c.GetFunction(ctx, r)
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
type createFunctionOperation struct {
	response map[string]interface{}
}

func (op *createFunctionOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createFunctionOperation) do(ctx context.Context, r *Function, c *Client) error {
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

	if _, err := c.GetFunction(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getFunctionRaw(ctx context.Context, r *Function) ([]byte, error) {

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

func (c *Client) functionDiffsForRawDesired(ctx context.Context, rawDesired *Function, opts ...dcl.ApplyOption) (initial, desired *Function, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Function
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Function); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Function, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetFunction(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Function resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Function resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Function resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeFunctionDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Function: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Function: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractFunctionFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeFunctionInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Function: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeFunctionDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Function: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffFunction(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeFunctionInitialState(rawInitial, rawDesired *Function) (*Function, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.

	if !dcl.IsZeroValue(rawInitial.SourceArchiveUrl) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.SourceRepository) {
			rawInitial.SourceArchiveUrl = dcl.String("")
		}
	}

	if !dcl.IsZeroValue(rawInitial.SourceRepository) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.SourceArchiveUrl) {
			rawInitial.SourceRepository = EmptyFunctionSourceRepository
		}
	}

	if !dcl.IsZeroValue(rawInitial.EventTrigger) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.HttpsTrigger) {
			rawInitial.EventTrigger = EmptyFunctionEventTrigger
		}
	}

	if !dcl.IsZeroValue(rawInitial.HttpsTrigger) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.EventTrigger) {
			rawInitial.HttpsTrigger = EmptyFunctionHttpsTrigger
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

func canonicalizeFunctionDesiredState(rawDesired, rawInitial *Function, opts ...dcl.ApplyOption) (*Function, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.SourceRepository = canonicalizeFunctionSourceRepository(rawDesired.SourceRepository, nil, opts...)
		rawDesired.HttpsTrigger = canonicalizeFunctionHttpsTrigger(rawDesired.HttpsTrigger, nil, opts...)
		rawDesired.EventTrigger = canonicalizeFunctionEventTrigger(rawDesired.EventTrigger, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Function{}
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
	if dcl.StringCanonicalize(rawDesired.SourceArchiveUrl, rawInitial.SourceArchiveUrl) {
		canonicalDesired.SourceArchiveUrl = rawInitial.SourceArchiveUrl
	} else {
		canonicalDesired.SourceArchiveUrl = rawDesired.SourceArchiveUrl
	}
	canonicalDesired.SourceRepository = canonicalizeFunctionSourceRepository(rawDesired.SourceRepository, rawInitial.SourceRepository, opts...)
	canonicalDesired.HttpsTrigger = canonicalizeFunctionHttpsTrigger(rawDesired.HttpsTrigger, rawInitial.HttpsTrigger, opts...)
	canonicalDesired.EventTrigger = canonicalizeFunctionEventTrigger(rawDesired.EventTrigger, rawInitial.EventTrigger, opts...)
	if dcl.StringCanonicalize(rawDesired.EntryPoint, rawInitial.EntryPoint) {
		canonicalDesired.EntryPoint = rawInitial.EntryPoint
	} else {
		canonicalDesired.EntryPoint = rawDesired.EntryPoint
	}
	if dcl.StringCanonicalize(rawDesired.Runtime, rawInitial.Runtime) {
		canonicalDesired.Runtime = rawInitial.Runtime
	} else {
		canonicalDesired.Runtime = rawDesired.Runtime
	}
	if dcl.StringCanonicalize(rawDesired.Timeout, rawInitial.Timeout) {
		canonicalDesired.Timeout = rawInitial.Timeout
	} else {
		canonicalDesired.Timeout = rawDesired.Timeout
	}
	if dcl.IsZeroValue(rawDesired.AvailableMemoryMb) || (dcl.IsEmptyValueIndirect(rawDesired.AvailableMemoryMb) && dcl.IsEmptyValueIndirect(rawInitial.AvailableMemoryMb)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.AvailableMemoryMb = rawInitial.AvailableMemoryMb
	} else {
		canonicalDesired.AvailableMemoryMb = rawDesired.AvailableMemoryMb
	}
	if dcl.IsZeroValue(rawDesired.ServiceAccountEmail) || (dcl.IsEmptyValueIndirect(rawDesired.ServiceAccountEmail) && dcl.IsEmptyValueIndirect(rawInitial.ServiceAccountEmail)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.ServiceAccountEmail = rawInitial.ServiceAccountEmail
	} else {
		canonicalDesired.ServiceAccountEmail = rawDesired.ServiceAccountEmail
	}
	if dcl.IsZeroValue(rawDesired.Labels) || (dcl.IsEmptyValueIndirect(rawDesired.Labels) && dcl.IsEmptyValueIndirect(rawInitial.Labels)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	if dcl.IsZeroValue(rawDesired.EnvironmentVariables) || (dcl.IsEmptyValueIndirect(rawDesired.EnvironmentVariables) && dcl.IsEmptyValueIndirect(rawInitial.EnvironmentVariables)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.EnvironmentVariables = rawInitial.EnvironmentVariables
	} else {
		canonicalDesired.EnvironmentVariables = rawDesired.EnvironmentVariables
	}
	if dcl.IsZeroValue(rawDesired.MaxInstances) || (dcl.IsEmptyValueIndirect(rawDesired.MaxInstances) && dcl.IsEmptyValueIndirect(rawInitial.MaxInstances)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.MaxInstances = rawInitial.MaxInstances
	} else {
		canonicalDesired.MaxInstances = rawDesired.MaxInstances
	}
	if dcl.IsZeroValue(rawDesired.VPCConnector) || (dcl.IsEmptyValueIndirect(rawDesired.VPCConnector) && dcl.IsEmptyValueIndirect(rawInitial.VPCConnector)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.VPCConnector = rawInitial.VPCConnector
	} else {
		canonicalDesired.VPCConnector = rawDesired.VPCConnector
	}
	if dcl.IsZeroValue(rawDesired.VPCConnectorEgressSettings) || (dcl.IsEmptyValueIndirect(rawDesired.VPCConnectorEgressSettings) && dcl.IsEmptyValueIndirect(rawInitial.VPCConnectorEgressSettings)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.VPCConnectorEgressSettings = rawInitial.VPCConnectorEgressSettings
	} else {
		canonicalDesired.VPCConnectorEgressSettings = rawDesired.VPCConnectorEgressSettings
	}
	if dcl.IsZeroValue(rawDesired.IngressSettings) || (dcl.IsEmptyValueIndirect(rawDesired.IngressSettings) && dcl.IsEmptyValueIndirect(rawInitial.IngressSettings)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.IngressSettings = rawInitial.IngressSettings
	} else {
		canonicalDesired.IngressSettings = rawDesired.IngressSettings
	}
	if dcl.NameToSelfLink(rawDesired.Region, rawInitial.Region) {
		canonicalDesired.Region = rawInitial.Region
	} else {
		canonicalDesired.Region = rawDesired.Region
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}

	if canonicalDesired.SourceArchiveUrl != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.SourceRepository) {
			canonicalDesired.SourceArchiveUrl = dcl.String("")
		}
	}

	if canonicalDesired.SourceRepository != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.SourceArchiveUrl) {
			canonicalDesired.SourceRepository = EmptyFunctionSourceRepository
		}
	}

	if canonicalDesired.EventTrigger != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.HttpsTrigger) {
			canonicalDesired.EventTrigger = EmptyFunctionEventTrigger
		}
	}

	if canonicalDesired.HttpsTrigger != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.EventTrigger) {
			canonicalDesired.HttpsTrigger = EmptyFunctionHttpsTrigger
		}
	}

	return canonicalDesired, nil
}

func canonicalizeFunctionNewState(c *Client, rawNew, rawDesired *Function) (*Function, error) {

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

	if dcl.IsEmptyValueIndirect(rawNew.SourceArchiveUrl) && dcl.IsEmptyValueIndirect(rawDesired.SourceArchiveUrl) {
		rawNew.SourceArchiveUrl = rawDesired.SourceArchiveUrl
	} else {
		if dcl.StringCanonicalize(rawDesired.SourceArchiveUrl, rawNew.SourceArchiveUrl) {
			rawNew.SourceArchiveUrl = rawDesired.SourceArchiveUrl
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceRepository) && dcl.IsEmptyValueIndirect(rawDesired.SourceRepository) {
		rawNew.SourceRepository = rawDesired.SourceRepository
	} else {
		rawNew.SourceRepository = canonicalizeNewFunctionSourceRepository(c, rawDesired.SourceRepository, rawNew.SourceRepository)
	}

	if dcl.IsEmptyValueIndirect(rawNew.HttpsTrigger) && dcl.IsEmptyValueIndirect(rawDesired.HttpsTrigger) {
		rawNew.HttpsTrigger = rawDesired.HttpsTrigger
	} else {
		rawNew.HttpsTrigger = canonicalizeNewFunctionHttpsTrigger(c, rawDesired.HttpsTrigger, rawNew.HttpsTrigger)
	}

	if dcl.IsEmptyValueIndirect(rawNew.EventTrigger) && dcl.IsEmptyValueIndirect(rawDesired.EventTrigger) {
		rawNew.EventTrigger = rawDesired.EventTrigger
	} else {
		rawNew.EventTrigger = canonicalizeNewFunctionEventTrigger(c, rawDesired.EventTrigger, rawNew.EventTrigger)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Status) && dcl.IsEmptyValueIndirect(rawDesired.Status) {
		rawNew.Status = rawDesired.Status
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.EntryPoint) && dcl.IsEmptyValueIndirect(rawDesired.EntryPoint) {
		rawNew.EntryPoint = rawDesired.EntryPoint
	} else {
		if dcl.StringCanonicalize(rawDesired.EntryPoint, rawNew.EntryPoint) {
			rawNew.EntryPoint = rawDesired.EntryPoint
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Runtime) && dcl.IsEmptyValueIndirect(rawDesired.Runtime) {
		rawNew.Runtime = rawDesired.Runtime
	} else {
		if dcl.StringCanonicalize(rawDesired.Runtime, rawNew.Runtime) {
			rawNew.Runtime = rawDesired.Runtime
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Timeout) && dcl.IsEmptyValueIndirect(rawDesired.Timeout) {
		rawNew.Timeout = rawDesired.Timeout
	} else {
		if dcl.StringCanonicalize(rawDesired.Timeout, rawNew.Timeout) {
			rawNew.Timeout = rawDesired.Timeout
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.AvailableMemoryMb) && dcl.IsEmptyValueIndirect(rawDesired.AvailableMemoryMb) {
		rawNew.AvailableMemoryMb = rawDesired.AvailableMemoryMb
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ServiceAccountEmail) && dcl.IsEmptyValueIndirect(rawDesired.ServiceAccountEmail) {
		rawNew.ServiceAccountEmail = rawDesired.ServiceAccountEmail
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
		if dcl.StringCanonicalize(rawDesired.UpdateTime, rawNew.UpdateTime) {
			rawNew.UpdateTime = rawDesired.UpdateTime
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.VersionId) && dcl.IsEmptyValueIndirect(rawDesired.VersionId) {
		rawNew.VersionId = rawDesired.VersionId
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.EnvironmentVariables) && dcl.IsEmptyValueIndirect(rawDesired.EnvironmentVariables) {
		rawNew.EnvironmentVariables = rawDesired.EnvironmentVariables
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.MaxInstances) && dcl.IsEmptyValueIndirect(rawDesired.MaxInstances) {
		rawNew.MaxInstances = rawDesired.MaxInstances
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.VPCConnector) && dcl.IsEmptyValueIndirect(rawDesired.VPCConnector) {
		rawNew.VPCConnector = rawDesired.VPCConnector
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.VPCConnectorEgressSettings) && dcl.IsEmptyValueIndirect(rawDesired.VPCConnectorEgressSettings) {
		rawNew.VPCConnectorEgressSettings = rawDesired.VPCConnectorEgressSettings
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.IngressSettings) && dcl.IsEmptyValueIndirect(rawDesired.IngressSettings) {
		rawNew.IngressSettings = rawDesired.IngressSettings
	} else {
	}

	rawNew.Region = rawDesired.Region

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeFunctionSourceRepository(des, initial *FunctionSourceRepository, opts ...dcl.ApplyOption) *FunctionSourceRepository {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FunctionSourceRepository{}

	if CanonicalizeFunctionSourceRepoURL(des.Url, initial.Url) || dcl.IsZeroValue(des.Url) {
		cDes.Url = initial.Url
	} else {
		cDes.Url = des.Url
	}

	return cDes
}

func canonicalizeFunctionSourceRepositorySlice(des, initial []FunctionSourceRepository, opts ...dcl.ApplyOption) []FunctionSourceRepository {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FunctionSourceRepository, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFunctionSourceRepository(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FunctionSourceRepository, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFunctionSourceRepository(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFunctionSourceRepository(c *Client, des, nw *FunctionSourceRepository) *FunctionSourceRepository {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FunctionSourceRepository while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if CanonicalizeFunctionSourceRepoURL(des.Url, nw.Url) {
		nw.Url = des.Url
	}
	if dcl.StringCanonicalize(des.DeployedUrl, nw.DeployedUrl) {
		nw.DeployedUrl = des.DeployedUrl
	}

	return nw
}

func canonicalizeNewFunctionSourceRepositorySet(c *Client, des, nw []FunctionSourceRepository) []FunctionSourceRepository {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FunctionSourceRepository
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFunctionSourceRepositoryNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFunctionSourceRepository(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFunctionSourceRepositorySlice(c *Client, des, nw []FunctionSourceRepository) []FunctionSourceRepository {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FunctionSourceRepository
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFunctionSourceRepository(c, &d, &n))
	}

	return items
}

func canonicalizeFunctionHttpsTrigger(des, initial *FunctionHttpsTrigger, opts ...dcl.ApplyOption) *FunctionHttpsTrigger {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FunctionHttpsTrigger{}

	if dcl.IsZeroValue(des.SecurityLevel) || (dcl.IsEmptyValueIndirect(des.SecurityLevel) && dcl.IsEmptyValueIndirect(initial.SecurityLevel)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.SecurityLevel = initial.SecurityLevel
	} else {
		cDes.SecurityLevel = des.SecurityLevel
	}

	return cDes
}

func canonicalizeFunctionHttpsTriggerSlice(des, initial []FunctionHttpsTrigger, opts ...dcl.ApplyOption) []FunctionHttpsTrigger {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FunctionHttpsTrigger, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFunctionHttpsTrigger(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FunctionHttpsTrigger, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFunctionHttpsTrigger(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFunctionHttpsTrigger(c *Client, des, nw *FunctionHttpsTrigger) *FunctionHttpsTrigger {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FunctionHttpsTrigger while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Url, nw.Url) {
		nw.Url = des.Url
	}

	return nw
}

func canonicalizeNewFunctionHttpsTriggerSet(c *Client, des, nw []FunctionHttpsTrigger) []FunctionHttpsTrigger {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FunctionHttpsTrigger
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFunctionHttpsTriggerNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFunctionHttpsTrigger(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFunctionHttpsTriggerSlice(c *Client, des, nw []FunctionHttpsTrigger) []FunctionHttpsTrigger {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FunctionHttpsTrigger
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFunctionHttpsTrigger(c, &d, &n))
	}

	return items
}

func canonicalizeFunctionEventTrigger(des, initial *FunctionEventTrigger, opts ...dcl.ApplyOption) *FunctionEventTrigger {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FunctionEventTrigger{}

	if dcl.StringCanonicalize(des.EventType, initial.EventType) || dcl.IsZeroValue(des.EventType) {
		cDes.EventType = initial.EventType
	} else {
		cDes.EventType = des.EventType
	}
	if dcl.IsZeroValue(des.Resource) || (dcl.IsEmptyValueIndirect(des.Resource) && dcl.IsEmptyValueIndirect(initial.Resource)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Resource = initial.Resource
	} else {
		cDes.Resource = des.Resource
	}
	if dcl.StringCanonicalize(des.Service, initial.Service) || dcl.IsZeroValue(des.Service) {
		cDes.Service = initial.Service
	} else {
		cDes.Service = des.Service
	}
	if dcl.BoolCanonicalize(des.FailurePolicy, initial.FailurePolicy) || dcl.IsZeroValue(des.FailurePolicy) {
		cDes.FailurePolicy = initial.FailurePolicy
	} else {
		cDes.FailurePolicy = des.FailurePolicy
	}

	return cDes
}

func canonicalizeFunctionEventTriggerSlice(des, initial []FunctionEventTrigger, opts ...dcl.ApplyOption) []FunctionEventTrigger {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FunctionEventTrigger, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFunctionEventTrigger(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FunctionEventTrigger, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFunctionEventTrigger(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFunctionEventTrigger(c *Client, des, nw *FunctionEventTrigger) *FunctionEventTrigger {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FunctionEventTrigger while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.EventType, nw.EventType) {
		nw.EventType = des.EventType
	}
	if dcl.StringCanonicalize(des.Service, nw.Service) {
		nw.Service = des.Service
	}
	if dcl.BoolCanonicalize(des.FailurePolicy, nw.FailurePolicy) {
		nw.FailurePolicy = des.FailurePolicy
	}

	return nw
}

func canonicalizeNewFunctionEventTriggerSet(c *Client, des, nw []FunctionEventTrigger) []FunctionEventTrigger {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FunctionEventTrigger
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFunctionEventTriggerNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFunctionEventTrigger(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFunctionEventTriggerSlice(c *Client, des, nw []FunctionEventTrigger) []FunctionEventTrigger {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FunctionEventTrigger
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFunctionEventTrigger(c, &d, &n))
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
func diffFunction(c *Client, desired, actual *Function, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateFunctionUpdateOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceArchiveUrl, actual.SourceArchiveUrl, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceArchiveUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceRepository, actual.SourceRepository, dcl.DiffInfo{ObjectFunction: compareFunctionSourceRepositoryNewStyle, EmptyObject: EmptyFunctionSourceRepository, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceRepository")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HttpsTrigger, actual.HttpsTrigger, dcl.DiffInfo{ObjectFunction: compareFunctionHttpsTriggerNewStyle, EmptyObject: EmptyFunctionHttpsTrigger, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HttpsTrigger")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EventTrigger, actual.EventTrigger, dcl.DiffInfo{ObjectFunction: compareFunctionEventTriggerNewStyle, EmptyObject: EmptyFunctionEventTrigger, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EventTrigger")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Status, actual.Status, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Status")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EntryPoint, actual.EntryPoint, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EntryPoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Runtime, actual.Runtime, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateFunctionUpdateOperation")}, fn.AddNest("Runtime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Timeout, actual.Timeout, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.TriggersOperation("updateFunctionUpdateOperation")}, fn.AddNest("Timeout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AvailableMemoryMb, actual.AvailableMemoryMb, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.TriggersOperation("updateFunctionUpdateOperation")}, fn.AddNest("AvailableMemoryMb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceAccountEmail, actual.ServiceAccountEmail, dcl.DiffInfo{ServerDefault: true, Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServiceAccountEmail")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.VersionId, actual.VersionId, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("VersionId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateFunctionUpdateOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnvironmentVariables, actual.EnvironmentVariables, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateFunctionUpdateOperation")}, fn.AddNest("EnvironmentVariables")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxInstances, actual.MaxInstances, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateFunctionUpdateOperation")}, fn.AddNest("MaxInstances")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VPCConnector, actual.VPCConnector, dcl.DiffInfo{ServerDefault: true, Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateFunctionUpdateOperation")}, fn.AddNest("VpcConnector")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VPCConnectorEgressSettings, actual.VPCConnectorEgressSettings, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateFunctionUpdateOperation")}, fn.AddNest("VpcConnectorEgressSettings")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IngressSettings, actual.IngressSettings, dcl.DiffInfo{ServerDefault: true, Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateFunctionUpdateOperation")}, fn.AddNest("IngressSettings")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Region, actual.Region, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Region")); len(ds) != 0 || err != nil {
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
func compareFunctionSourceRepositoryNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FunctionSourceRepository)
	if !ok {
		desiredNotPointer, ok := d.(FunctionSourceRepository)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FunctionSourceRepository or *FunctionSourceRepository", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FunctionSourceRepository)
	if !ok {
		actualNotPointer, ok := a.(FunctionSourceRepository)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FunctionSourceRepository", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Url, actual.Url, dcl.DiffInfo{CustomDiff: CanonicalizeFunctionSourceRepoURL, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Url")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DeployedUrl, actual.DeployedUrl, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DeployedUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFunctionHttpsTriggerNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FunctionHttpsTrigger)
	if !ok {
		desiredNotPointer, ok := d.(FunctionHttpsTrigger)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FunctionHttpsTrigger or *FunctionHttpsTrigger", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FunctionHttpsTrigger)
	if !ok {
		actualNotPointer, ok := a.(FunctionHttpsTrigger)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FunctionHttpsTrigger", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Url, actual.Url, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Url")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SecurityLevel, actual.SecurityLevel, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SecurityLevel")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFunctionEventTriggerNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FunctionEventTrigger)
	if !ok {
		desiredNotPointer, ok := d.(FunctionEventTrigger)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FunctionEventTrigger or *FunctionEventTrigger", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FunctionEventTrigger)
	if !ok {
		actualNotPointer, ok := a.(FunctionEventTrigger)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FunctionEventTrigger", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.EventType, actual.EventType, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EventType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Resource, actual.Resource, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Resource")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Service, actual.Service, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Service")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FailurePolicy, actual.FailurePolicy, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FailurePolicy")); len(ds) != 0 || err != nil {
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
func (r *Function) urlNormalized() *Function {
	normalized := dcl.Copy(*r).(Function)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.SourceArchiveUrl = dcl.SelfLinkToName(r.SourceArchiveUrl)
	normalized.EntryPoint = dcl.SelfLinkToName(r.EntryPoint)
	normalized.Runtime = dcl.SelfLinkToName(r.Runtime)
	normalized.Timeout = dcl.SelfLinkToName(r.Timeout)
	normalized.ServiceAccountEmail = dcl.SelfLinkToName(r.ServiceAccountEmail)
	normalized.UpdateTime = dcl.SelfLinkToName(r.UpdateTime)
	normalized.VPCConnector = dcl.SelfLinkToName(r.VPCConnector)
	normalized.Region = dcl.SelfLinkToName(r.Region)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *Function) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(nr.Project),
			"region":  dcl.ValueOrEmptyString(nr.Region),
			"name":    dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{region}}/functions/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Function resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Function) marshal(c *Client) ([]byte, error) {
	m, err := expandFunction(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Function: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalFunction decodes JSON responses into the Function resource schema.
func unmarshalFunction(b []byte, c *Client, res *Function) (*Function, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapFunction(m, c, res)
}

func unmarshalMapFunction(m map[string]interface{}, c *Client, res *Function) (*Function, error) {

	flattened := flattenFunction(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandFunction expands Function into a JSON request object.
func expandFunction(c *Client, f *Function) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("projects/%s/locations/%s/functions/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Region), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v := f.SourceArchiveUrl; dcl.ValueShouldBeSent(v) {
		m["sourceArchiveUrl"] = v
	}
	if v, err := expandFunctionSourceRepository(c, f.SourceRepository, res); err != nil {
		return nil, fmt.Errorf("error expanding SourceRepository into sourceRepository: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["sourceRepository"] = v
	}
	if v, err := expandFunctionHttpsTrigger(c, f.HttpsTrigger, res); err != nil {
		return nil, fmt.Errorf("error expanding HttpsTrigger into httpsTrigger: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["httpsTrigger"] = v
	}
	if v, err := expandFunctionEventTrigger(c, f.EventTrigger, res); err != nil {
		return nil, fmt.Errorf("error expanding EventTrigger into eventTrigger: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["eventTrigger"] = v
	}
	if v := f.EntryPoint; dcl.ValueShouldBeSent(v) {
		m["entryPoint"] = v
	}
	if v := f.Runtime; dcl.ValueShouldBeSent(v) {
		m["runtime"] = v
	}
	if v := f.Timeout; dcl.ValueShouldBeSent(v) {
		m["timeout"] = v
	}
	if v := f.AvailableMemoryMb; dcl.ValueShouldBeSent(v) {
		m["availableMemoryMb"] = v
	}
	if v := f.ServiceAccountEmail; dcl.ValueShouldBeSent(v) {
		m["serviceAccountEmail"] = v
	}
	if v := f.Labels; dcl.ValueShouldBeSent(v) {
		m["labels"] = v
	}
	if v := f.EnvironmentVariables; dcl.ValueShouldBeSent(v) {
		m["environmentVariables"] = v
	}
	if v := f.MaxInstances; dcl.ValueShouldBeSent(v) {
		m["maxInstances"] = v
	}
	if v := f.VPCConnector; dcl.ValueShouldBeSent(v) {
		m["vpcConnector"] = v
	}
	if v := f.VPCConnectorEgressSettings; dcl.ValueShouldBeSent(v) {
		m["vpcConnectorEgressSettings"] = v
	}
	if v := f.IngressSettings; dcl.ValueShouldBeSent(v) {
		m["ingressSettings"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Region into region: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["region"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}

	return m, nil
}

// flattenFunction flattens Function from a JSON request object into the
// Function type.
func flattenFunction(c *Client, i interface{}, res *Function) *Function {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Function{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.SourceArchiveUrl = dcl.FlattenString(m["sourceArchiveUrl"])
	resultRes.SourceRepository = flattenFunctionSourceRepository(c, m["sourceRepository"], res)
	resultRes.HttpsTrigger = flattenFunctionHttpsTrigger(c, m["httpsTrigger"], res)
	resultRes.EventTrigger = flattenFunctionEventTrigger(c, m["eventTrigger"], res)
	resultRes.Status = flattenFunctionStatusEnum(m["status"])
	resultRes.EntryPoint = dcl.FlattenString(m["entryPoint"])
	resultRes.Runtime = dcl.FlattenString(m["runtime"])
	resultRes.Timeout = dcl.FlattenString(m["timeout"])
	resultRes.AvailableMemoryMb = dcl.FlattenInteger(m["availableMemoryMb"])
	resultRes.ServiceAccountEmail = dcl.FlattenString(m["serviceAccountEmail"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.VersionId = dcl.FlattenInteger(m["versionId"])
	resultRes.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	resultRes.EnvironmentVariables = dcl.FlattenKeyValuePairs(m["environmentVariables"])
	resultRes.MaxInstances = dcl.FlattenInteger(m["maxInstances"])
	resultRes.VPCConnector = dcl.FlattenString(m["vpcConnector"])
	resultRes.VPCConnectorEgressSettings = flattenFunctionVPCConnectorEgressSettingsEnum(m["vpcConnectorEgressSettings"])
	resultRes.IngressSettings = flattenFunctionIngressSettingsEnum(m["ingressSettings"])
	resultRes.Region = dcl.FlattenString(m["region"])
	resultRes.Project = dcl.FlattenString(m["project"])

	return resultRes
}

// expandFunctionSourceRepositoryMap expands the contents of FunctionSourceRepository into a JSON
// request object.
func expandFunctionSourceRepositoryMap(c *Client, f map[string]FunctionSourceRepository, res *Function) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFunctionSourceRepository(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFunctionSourceRepositorySlice expands the contents of FunctionSourceRepository into a JSON
// request object.
func expandFunctionSourceRepositorySlice(c *Client, f []FunctionSourceRepository, res *Function) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFunctionSourceRepository(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFunctionSourceRepositoryMap flattens the contents of FunctionSourceRepository from a JSON
// response object.
func flattenFunctionSourceRepositoryMap(c *Client, i interface{}, res *Function) map[string]FunctionSourceRepository {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FunctionSourceRepository{}
	}

	if len(a) == 0 {
		return map[string]FunctionSourceRepository{}
	}

	items := make(map[string]FunctionSourceRepository)
	for k, item := range a {
		items[k] = *flattenFunctionSourceRepository(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFunctionSourceRepositorySlice flattens the contents of FunctionSourceRepository from a JSON
// response object.
func flattenFunctionSourceRepositorySlice(c *Client, i interface{}, res *Function) []FunctionSourceRepository {
	a, ok := i.([]interface{})
	if !ok {
		return []FunctionSourceRepository{}
	}

	if len(a) == 0 {
		return []FunctionSourceRepository{}
	}

	items := make([]FunctionSourceRepository, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFunctionSourceRepository(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFunctionSourceRepository expands an instance of FunctionSourceRepository into a JSON
// request object.
func expandFunctionSourceRepository(c *Client, f *FunctionSourceRepository, res *Function) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Url; !dcl.IsEmptyValueIndirect(v) {
		m["url"] = v
	}

	return m, nil
}

// flattenFunctionSourceRepository flattens an instance of FunctionSourceRepository from a JSON
// response object.
func flattenFunctionSourceRepository(c *Client, i interface{}, res *Function) *FunctionSourceRepository {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FunctionSourceRepository{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFunctionSourceRepository
	}
	r.Url = dcl.FlattenString(m["url"])
	r.DeployedUrl = dcl.FlattenString(m["deployedUrl"])

	return r
}

// expandFunctionHttpsTriggerMap expands the contents of FunctionHttpsTrigger into a JSON
// request object.
func expandFunctionHttpsTriggerMap(c *Client, f map[string]FunctionHttpsTrigger, res *Function) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFunctionHttpsTrigger(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFunctionHttpsTriggerSlice expands the contents of FunctionHttpsTrigger into a JSON
// request object.
func expandFunctionHttpsTriggerSlice(c *Client, f []FunctionHttpsTrigger, res *Function) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFunctionHttpsTrigger(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFunctionHttpsTriggerMap flattens the contents of FunctionHttpsTrigger from a JSON
// response object.
func flattenFunctionHttpsTriggerMap(c *Client, i interface{}, res *Function) map[string]FunctionHttpsTrigger {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FunctionHttpsTrigger{}
	}

	if len(a) == 0 {
		return map[string]FunctionHttpsTrigger{}
	}

	items := make(map[string]FunctionHttpsTrigger)
	for k, item := range a {
		items[k] = *flattenFunctionHttpsTrigger(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFunctionHttpsTriggerSlice flattens the contents of FunctionHttpsTrigger from a JSON
// response object.
func flattenFunctionHttpsTriggerSlice(c *Client, i interface{}, res *Function) []FunctionHttpsTrigger {
	a, ok := i.([]interface{})
	if !ok {
		return []FunctionHttpsTrigger{}
	}

	if len(a) == 0 {
		return []FunctionHttpsTrigger{}
	}

	items := make([]FunctionHttpsTrigger, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFunctionHttpsTrigger(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFunctionHttpsTrigger expands an instance of FunctionHttpsTrigger into a JSON
// request object.
func expandFunctionHttpsTrigger(c *Client, f *FunctionHttpsTrigger, res *Function) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SecurityLevel; !dcl.IsEmptyValueIndirect(v) {
		m["securityLevel"] = v
	}

	return m, nil
}

// flattenFunctionHttpsTrigger flattens an instance of FunctionHttpsTrigger from a JSON
// response object.
func flattenFunctionHttpsTrigger(c *Client, i interface{}, res *Function) *FunctionHttpsTrigger {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FunctionHttpsTrigger{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFunctionHttpsTrigger
	}
	r.Url = dcl.FlattenString(m["url"])
	r.SecurityLevel = flattenFunctionHttpsTriggerSecurityLevelEnum(m["securityLevel"])

	return r
}

// expandFunctionEventTriggerMap expands the contents of FunctionEventTrigger into a JSON
// request object.
func expandFunctionEventTriggerMap(c *Client, f map[string]FunctionEventTrigger, res *Function) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFunctionEventTrigger(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFunctionEventTriggerSlice expands the contents of FunctionEventTrigger into a JSON
// request object.
func expandFunctionEventTriggerSlice(c *Client, f []FunctionEventTrigger, res *Function) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFunctionEventTrigger(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFunctionEventTriggerMap flattens the contents of FunctionEventTrigger from a JSON
// response object.
func flattenFunctionEventTriggerMap(c *Client, i interface{}, res *Function) map[string]FunctionEventTrigger {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FunctionEventTrigger{}
	}

	if len(a) == 0 {
		return map[string]FunctionEventTrigger{}
	}

	items := make(map[string]FunctionEventTrigger)
	for k, item := range a {
		items[k] = *flattenFunctionEventTrigger(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFunctionEventTriggerSlice flattens the contents of FunctionEventTrigger from a JSON
// response object.
func flattenFunctionEventTriggerSlice(c *Client, i interface{}, res *Function) []FunctionEventTrigger {
	a, ok := i.([]interface{})
	if !ok {
		return []FunctionEventTrigger{}
	}

	if len(a) == 0 {
		return []FunctionEventTrigger{}
	}

	items := make([]FunctionEventTrigger, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFunctionEventTrigger(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFunctionEventTrigger expands an instance of FunctionEventTrigger into a JSON
// request object.
func expandFunctionEventTrigger(c *Client, f *FunctionEventTrigger, res *Function) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.EventType; !dcl.IsEmptyValueIndirect(v) {
		m["eventType"] = v
	}
	if v, err := ExpandFunctionEventTriggerResource(c, f.Resource, res); err != nil {
		return nil, fmt.Errorf("error expanding Resource into resource: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["resource"] = v
	}
	if v := f.Service; !dcl.IsEmptyValueIndirect(v) {
		m["service"] = v
	}
	if v, err := ExpandFunctionEventTriggerFailurePolicy(c, f.FailurePolicy, res); err != nil {
		return nil, fmt.Errorf("error expanding FailurePolicy into failurePolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["failurePolicy"] = v
	}

	return m, nil
}

// flattenFunctionEventTrigger flattens an instance of FunctionEventTrigger from a JSON
// response object.
func flattenFunctionEventTrigger(c *Client, i interface{}, res *Function) *FunctionEventTrigger {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FunctionEventTrigger{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFunctionEventTrigger
	}
	r.EventType = dcl.FlattenString(m["eventType"])
	r.Resource = dcl.FlattenString(m["resource"])
	r.Service = dcl.FlattenString(m["service"])
	r.FailurePolicy = flattenFunctionEventTriggerFailurePolicy(c, m["failurePolicy"], res)

	return r
}

// flattenFunctionHttpsTriggerSecurityLevelEnumMap flattens the contents of FunctionHttpsTriggerSecurityLevelEnum from a JSON
// response object.
func flattenFunctionHttpsTriggerSecurityLevelEnumMap(c *Client, i interface{}, res *Function) map[string]FunctionHttpsTriggerSecurityLevelEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FunctionHttpsTriggerSecurityLevelEnum{}
	}

	if len(a) == 0 {
		return map[string]FunctionHttpsTriggerSecurityLevelEnum{}
	}

	items := make(map[string]FunctionHttpsTriggerSecurityLevelEnum)
	for k, item := range a {
		items[k] = *flattenFunctionHttpsTriggerSecurityLevelEnum(item.(interface{}))
	}

	return items
}

// flattenFunctionHttpsTriggerSecurityLevelEnumSlice flattens the contents of FunctionHttpsTriggerSecurityLevelEnum from a JSON
// response object.
func flattenFunctionHttpsTriggerSecurityLevelEnumSlice(c *Client, i interface{}, res *Function) []FunctionHttpsTriggerSecurityLevelEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FunctionHttpsTriggerSecurityLevelEnum{}
	}

	if len(a) == 0 {
		return []FunctionHttpsTriggerSecurityLevelEnum{}
	}

	items := make([]FunctionHttpsTriggerSecurityLevelEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFunctionHttpsTriggerSecurityLevelEnum(item.(interface{})))
	}

	return items
}

// flattenFunctionHttpsTriggerSecurityLevelEnum asserts that an interface is a string, and returns a
// pointer to a *FunctionHttpsTriggerSecurityLevelEnum with the same value as that string.
func flattenFunctionHttpsTriggerSecurityLevelEnum(i interface{}) *FunctionHttpsTriggerSecurityLevelEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return FunctionHttpsTriggerSecurityLevelEnumRef(s)
}

// flattenFunctionStatusEnumMap flattens the contents of FunctionStatusEnum from a JSON
// response object.
func flattenFunctionStatusEnumMap(c *Client, i interface{}, res *Function) map[string]FunctionStatusEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FunctionStatusEnum{}
	}

	if len(a) == 0 {
		return map[string]FunctionStatusEnum{}
	}

	items := make(map[string]FunctionStatusEnum)
	for k, item := range a {
		items[k] = *flattenFunctionStatusEnum(item.(interface{}))
	}

	return items
}

// flattenFunctionStatusEnumSlice flattens the contents of FunctionStatusEnum from a JSON
// response object.
func flattenFunctionStatusEnumSlice(c *Client, i interface{}, res *Function) []FunctionStatusEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FunctionStatusEnum{}
	}

	if len(a) == 0 {
		return []FunctionStatusEnum{}
	}

	items := make([]FunctionStatusEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFunctionStatusEnum(item.(interface{})))
	}

	return items
}

// flattenFunctionStatusEnum asserts that an interface is a string, and returns a
// pointer to a *FunctionStatusEnum with the same value as that string.
func flattenFunctionStatusEnum(i interface{}) *FunctionStatusEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return FunctionStatusEnumRef(s)
}

// flattenFunctionVPCConnectorEgressSettingsEnumMap flattens the contents of FunctionVPCConnectorEgressSettingsEnum from a JSON
// response object.
func flattenFunctionVPCConnectorEgressSettingsEnumMap(c *Client, i interface{}, res *Function) map[string]FunctionVPCConnectorEgressSettingsEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FunctionVPCConnectorEgressSettingsEnum{}
	}

	if len(a) == 0 {
		return map[string]FunctionVPCConnectorEgressSettingsEnum{}
	}

	items := make(map[string]FunctionVPCConnectorEgressSettingsEnum)
	for k, item := range a {
		items[k] = *flattenFunctionVPCConnectorEgressSettingsEnum(item.(interface{}))
	}

	return items
}

// flattenFunctionVPCConnectorEgressSettingsEnumSlice flattens the contents of FunctionVPCConnectorEgressSettingsEnum from a JSON
// response object.
func flattenFunctionVPCConnectorEgressSettingsEnumSlice(c *Client, i interface{}, res *Function) []FunctionVPCConnectorEgressSettingsEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FunctionVPCConnectorEgressSettingsEnum{}
	}

	if len(a) == 0 {
		return []FunctionVPCConnectorEgressSettingsEnum{}
	}

	items := make([]FunctionVPCConnectorEgressSettingsEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFunctionVPCConnectorEgressSettingsEnum(item.(interface{})))
	}

	return items
}

// flattenFunctionVPCConnectorEgressSettingsEnum asserts that an interface is a string, and returns a
// pointer to a *FunctionVPCConnectorEgressSettingsEnum with the same value as that string.
func flattenFunctionVPCConnectorEgressSettingsEnum(i interface{}) *FunctionVPCConnectorEgressSettingsEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return FunctionVPCConnectorEgressSettingsEnumRef(s)
}

// flattenFunctionIngressSettingsEnumMap flattens the contents of FunctionIngressSettingsEnum from a JSON
// response object.
func flattenFunctionIngressSettingsEnumMap(c *Client, i interface{}, res *Function) map[string]FunctionIngressSettingsEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FunctionIngressSettingsEnum{}
	}

	if len(a) == 0 {
		return map[string]FunctionIngressSettingsEnum{}
	}

	items := make(map[string]FunctionIngressSettingsEnum)
	for k, item := range a {
		items[k] = *flattenFunctionIngressSettingsEnum(item.(interface{}))
	}

	return items
}

// flattenFunctionIngressSettingsEnumSlice flattens the contents of FunctionIngressSettingsEnum from a JSON
// response object.
func flattenFunctionIngressSettingsEnumSlice(c *Client, i interface{}, res *Function) []FunctionIngressSettingsEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FunctionIngressSettingsEnum{}
	}

	if len(a) == 0 {
		return []FunctionIngressSettingsEnum{}
	}

	items := make([]FunctionIngressSettingsEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFunctionIngressSettingsEnum(item.(interface{})))
	}

	return items
}

// flattenFunctionIngressSettingsEnum asserts that an interface is a string, and returns a
// pointer to a *FunctionIngressSettingsEnum with the same value as that string.
func flattenFunctionIngressSettingsEnum(i interface{}) *FunctionIngressSettingsEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return FunctionIngressSettingsEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Function) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalFunction(b, c, r)
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
		if nr.Region == nil && ncr.Region == nil {
			c.Config.Logger.Info("Both Region fields null - considering equal.")
		} else if nr.Region == nil || ncr.Region == nil {
			c.Config.Logger.Info("Only one Region field is null - considering unequal.")
			return false
		} else if *nr.Region != *ncr.Region {
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

type functionDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         functionApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToFunctionDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]functionDiff, error) {
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
	var diffs []functionDiff
	// For each operation name, create a functionDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := functionDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToFunctionApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToFunctionApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (functionApiOperation, error) {
	switch opName {

	case "updateFunctionUpdateOperation":
		return &updateFunctionUpdateOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractFunctionFields(r *Function) error {
	vSourceRepository := r.SourceRepository
	if vSourceRepository == nil {
		// note: explicitly not the empty object.
		vSourceRepository = &FunctionSourceRepository{}
	}
	if err := extractFunctionSourceRepositoryFields(r, vSourceRepository); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSourceRepository) {
		r.SourceRepository = vSourceRepository
	}
	vHttpsTrigger := r.HttpsTrigger
	if vHttpsTrigger == nil {
		// note: explicitly not the empty object.
		vHttpsTrigger = &FunctionHttpsTrigger{}
	}
	if err := extractFunctionHttpsTriggerFields(r, vHttpsTrigger); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vHttpsTrigger) {
		r.HttpsTrigger = vHttpsTrigger
	}
	vEventTrigger := r.EventTrigger
	if vEventTrigger == nil {
		// note: explicitly not the empty object.
		vEventTrigger = &FunctionEventTrigger{}
	}
	if err := extractFunctionEventTriggerFields(r, vEventTrigger); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vEventTrigger) {
		r.EventTrigger = vEventTrigger
	}
	return nil
}
func extractFunctionSourceRepositoryFields(r *Function, o *FunctionSourceRepository) error {
	return nil
}
func extractFunctionHttpsTriggerFields(r *Function, o *FunctionHttpsTrigger) error {
	return nil
}
func extractFunctionEventTriggerFields(r *Function, o *FunctionEventTrigger) error {
	return nil
}

func postReadExtractFunctionFields(r *Function) error {
	vSourceRepository := r.SourceRepository
	if vSourceRepository == nil {
		// note: explicitly not the empty object.
		vSourceRepository = &FunctionSourceRepository{}
	}
	if err := postReadExtractFunctionSourceRepositoryFields(r, vSourceRepository); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSourceRepository) {
		r.SourceRepository = vSourceRepository
	}
	vHttpsTrigger := r.HttpsTrigger
	if vHttpsTrigger == nil {
		// note: explicitly not the empty object.
		vHttpsTrigger = &FunctionHttpsTrigger{}
	}
	if err := postReadExtractFunctionHttpsTriggerFields(r, vHttpsTrigger); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vHttpsTrigger) {
		r.HttpsTrigger = vHttpsTrigger
	}
	vEventTrigger := r.EventTrigger
	if vEventTrigger == nil {
		// note: explicitly not the empty object.
		vEventTrigger = &FunctionEventTrigger{}
	}
	if err := postReadExtractFunctionEventTriggerFields(r, vEventTrigger); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vEventTrigger) {
		r.EventTrigger = vEventTrigger
	}
	return nil
}
func postReadExtractFunctionSourceRepositoryFields(r *Function, o *FunctionSourceRepository) error {
	return nil
}
func postReadExtractFunctionHttpsTriggerFields(r *Function, o *FunctionHttpsTrigger) error {
	return nil
}
func postReadExtractFunctionEventTriggerFields(r *Function, o *FunctionEventTrigger) error {
	return nil
}
