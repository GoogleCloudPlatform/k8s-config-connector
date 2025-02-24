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
	if err := dcl.Required(r, "type"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.NetworkConfig) {
		if err := r.NetworkConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceNetworkConfig) validate() error {
	return nil
}
func (r *InstanceAvailableVersion) validate() error {
	return nil
}
func (r *Instance) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://datafusion.googleapis.com/v1beta1/", params)
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

	if v := f.EnableStackdriverLogging; !dcl.IsEmptyValueIndirect(v) {
		req["enableStackdriverLogging"] = v
	}
	if v := f.EnableStackdriverMonitoring; !dcl.IsEmptyValueIndirect(v) {
		req["enableStackdriverMonitoring"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		req["version"] = v
	}
	if v := f.DataprocServiceAccount; !dcl.IsEmptyValueIndirect(v) {
		req["dataprocServiceAccount"] = v
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
	mask := op.UpdateMask()
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
		rawDesired.NetworkConfig = canonicalizeInstanceNetworkConfig(rawDesired.NetworkConfig, nil, opts...)

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
	if dcl.IsZeroValue(rawDesired.Type) || (dcl.IsEmptyValueIndirect(rawDesired.Type) && dcl.IsEmptyValueIndirect(rawInitial.Type)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Type = rawInitial.Type
	} else {
		canonicalDesired.Type = rawDesired.Type
	}
	if dcl.BoolCanonicalize(rawDesired.EnableStackdriverLogging, rawInitial.EnableStackdriverLogging) {
		canonicalDesired.EnableStackdriverLogging = rawInitial.EnableStackdriverLogging
	} else {
		canonicalDesired.EnableStackdriverLogging = rawDesired.EnableStackdriverLogging
	}
	if dcl.BoolCanonicalize(rawDesired.EnableStackdriverMonitoring, rawInitial.EnableStackdriverMonitoring) {
		canonicalDesired.EnableStackdriverMonitoring = rawInitial.EnableStackdriverMonitoring
	} else {
		canonicalDesired.EnableStackdriverMonitoring = rawDesired.EnableStackdriverMonitoring
	}
	if dcl.BoolCanonicalize(rawDesired.PrivateInstance, rawInitial.PrivateInstance) {
		canonicalDesired.PrivateInstance = rawInitial.PrivateInstance
	} else {
		canonicalDesired.PrivateInstance = rawDesired.PrivateInstance
	}
	canonicalDesired.NetworkConfig = canonicalizeInstanceNetworkConfig(rawDesired.NetworkConfig, rawInitial.NetworkConfig, opts...)
	if dcl.IsZeroValue(rawDesired.Labels) || (dcl.IsEmptyValueIndirect(rawDesired.Labels) && dcl.IsEmptyValueIndirect(rawInitial.Labels)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	if dcl.IsZeroValue(rawDesired.Options) || (dcl.IsEmptyValueIndirect(rawDesired.Options) && dcl.IsEmptyValueIndirect(rawInitial.Options)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Options = rawInitial.Options
	} else {
		canonicalDesired.Options = rawDesired.Options
	}
	if dcl.StringCanonicalize(rawDesired.Zone, rawInitial.Zone) {
		canonicalDesired.Zone = rawInitial.Zone
	} else {
		canonicalDesired.Zone = rawDesired.Zone
	}
	if dcl.StringCanonicalize(rawDesired.Version, rawInitial.Version) {
		canonicalDesired.Version = rawInitial.Version
	} else {
		canonicalDesired.Version = rawDesired.Version
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		canonicalDesired.DisplayName = rawInitial.DisplayName
	} else {
		canonicalDesired.DisplayName = rawDesired.DisplayName
	}
	if dcl.IsZeroValue(rawDesired.DataprocServiceAccount) || (dcl.IsEmptyValueIndirect(rawDesired.DataprocServiceAccount) && dcl.IsEmptyValueIndirect(rawInitial.DataprocServiceAccount)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.DataprocServiceAccount = rawInitial.DataprocServiceAccount
	} else {
		canonicalDesired.DataprocServiceAccount = rawDesired.DataprocServiceAccount
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

func canonicalizeInstanceNewState(c *Client, rawNew, rawDesired *Instance) (*Instance, error) {

	rawNew.Name = rawDesired.Name

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Type) && dcl.IsEmptyValueIndirect(rawDesired.Type) {
		rawNew.Type = rawDesired.Type
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.EnableStackdriverLogging) && dcl.IsEmptyValueIndirect(rawDesired.EnableStackdriverLogging) {
		rawNew.EnableStackdriverLogging = rawDesired.EnableStackdriverLogging
	} else {
		if dcl.BoolCanonicalize(rawDesired.EnableStackdriverLogging, rawNew.EnableStackdriverLogging) {
			rawNew.EnableStackdriverLogging = rawDesired.EnableStackdriverLogging
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.EnableStackdriverMonitoring) && dcl.IsEmptyValueIndirect(rawDesired.EnableStackdriverMonitoring) {
		rawNew.EnableStackdriverMonitoring = rawDesired.EnableStackdriverMonitoring
	} else {
		if dcl.BoolCanonicalize(rawDesired.EnableStackdriverMonitoring, rawNew.EnableStackdriverMonitoring) {
			rawNew.EnableStackdriverMonitoring = rawDesired.EnableStackdriverMonitoring
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.PrivateInstance) && dcl.IsEmptyValueIndirect(rawDesired.PrivateInstance) {
		rawNew.PrivateInstance = rawDesired.PrivateInstance
	} else {
		if dcl.BoolCanonicalize(rawDesired.PrivateInstance, rawNew.PrivateInstance) {
			rawNew.PrivateInstance = rawDesired.PrivateInstance
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.NetworkConfig) && dcl.IsEmptyValueIndirect(rawDesired.NetworkConfig) {
		rawNew.NetworkConfig = rawDesired.NetworkConfig
	} else {
		rawNew.NetworkConfig = canonicalizeNewInstanceNetworkConfig(c, rawDesired.NetworkConfig, rawNew.NetworkConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Options) && dcl.IsEmptyValueIndirect(rawDesired.Options) {
		rawNew.Options = rawDesired.Options
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

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.StateMessage) && dcl.IsEmptyValueIndirect(rawDesired.StateMessage) {
		rawNew.StateMessage = rawDesired.StateMessage
	} else {
		if dcl.StringCanonicalize(rawDesired.StateMessage, rawNew.StateMessage) {
			rawNew.StateMessage = rawDesired.StateMessage
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ServiceEndpoint) && dcl.IsEmptyValueIndirect(rawDesired.ServiceEndpoint) {
		rawNew.ServiceEndpoint = rawDesired.ServiceEndpoint
	} else {
		if dcl.StringCanonicalize(rawDesired.ServiceEndpoint, rawNew.ServiceEndpoint) {
			rawNew.ServiceEndpoint = rawDesired.ServiceEndpoint
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Zone) && dcl.IsEmptyValueIndirect(rawDesired.Zone) {
		rawNew.Zone = rawDesired.Zone
	} else {
		if dcl.StringCanonicalize(rawDesired.Zone, rawNew.Zone) {
			rawNew.Zone = rawDesired.Zone
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Version) && dcl.IsEmptyValueIndirect(rawDesired.Version) {
		rawNew.Version = rawDesired.Version
	} else {
		if dcl.StringCanonicalize(rawDesired.Version, rawNew.Version) {
			rawNew.Version = rawDesired.Version
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.AvailableVersion) && dcl.IsEmptyValueIndirect(rawDesired.AvailableVersion) {
		rawNew.AvailableVersion = rawDesired.AvailableVersion
	} else {
		rawNew.AvailableVersion = canonicalizeNewInstanceAvailableVersionSlice(c, rawDesired.AvailableVersion, rawNew.AvailableVersion)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ApiEndpoint) && dcl.IsEmptyValueIndirect(rawDesired.ApiEndpoint) {
		rawNew.ApiEndpoint = rawDesired.ApiEndpoint
	} else {
		if dcl.StringCanonicalize(rawDesired.ApiEndpoint, rawNew.ApiEndpoint) {
			rawNew.ApiEndpoint = rawDesired.ApiEndpoint
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.GcsBucket) && dcl.IsEmptyValueIndirect(rawDesired.GcsBucket) {
		rawNew.GcsBucket = rawDesired.GcsBucket
	} else {
		if dcl.StringCanonicalize(rawDesired.GcsBucket, rawNew.GcsBucket) {
			rawNew.GcsBucket = rawDesired.GcsBucket
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.P4ServiceAccount) && dcl.IsEmptyValueIndirect(rawDesired.P4ServiceAccount) {
		rawNew.P4ServiceAccount = rawDesired.P4ServiceAccount
	} else {
		if dcl.StringCanonicalize(rawDesired.P4ServiceAccount, rawNew.P4ServiceAccount) {
			rawNew.P4ServiceAccount = rawDesired.P4ServiceAccount
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.TenantProjectId) && dcl.IsEmptyValueIndirect(rawDesired.TenantProjectId) {
		rawNew.TenantProjectId = rawDesired.TenantProjectId
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DataprocServiceAccount) && dcl.IsEmptyValueIndirect(rawDesired.DataprocServiceAccount) {
		rawNew.DataprocServiceAccount = rawDesired.DataprocServiceAccount
	} else {
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeInstanceNetworkConfig(des, initial *InstanceNetworkConfig, opts ...dcl.ApplyOption) *InstanceNetworkConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceNetworkConfig{}

	if dcl.IsZeroValue(des.Network) || (dcl.IsEmptyValueIndirect(des.Network) && dcl.IsEmptyValueIndirect(initial.Network)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Network = initial.Network
	} else {
		cDes.Network = des.Network
	}
	if dcl.StringCanonicalize(des.IPAllocation, initial.IPAllocation) || dcl.IsZeroValue(des.IPAllocation) {
		cDes.IPAllocation = initial.IPAllocation
	} else {
		cDes.IPAllocation = des.IPAllocation
	}

	return cDes
}

func canonicalizeInstanceNetworkConfigSlice(des, initial []InstanceNetworkConfig, opts ...dcl.ApplyOption) []InstanceNetworkConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InstanceNetworkConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInstanceNetworkConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InstanceNetworkConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInstanceNetworkConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInstanceNetworkConfig(c *Client, des, nw *InstanceNetworkConfig) *InstanceNetworkConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InstanceNetworkConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.IPAllocation, nw.IPAllocation) {
		nw.IPAllocation = des.IPAllocation
	}

	return nw
}

func canonicalizeNewInstanceNetworkConfigSet(c *Client, des, nw []InstanceNetworkConfig) []InstanceNetworkConfig {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InstanceNetworkConfig
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInstanceNetworkConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInstanceNetworkConfig(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInstanceNetworkConfigSlice(c *Client, des, nw []InstanceNetworkConfig) []InstanceNetworkConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceNetworkConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceNetworkConfig(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceAvailableVersion(des, initial *InstanceAvailableVersion, opts ...dcl.ApplyOption) *InstanceAvailableVersion {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InstanceAvailableVersion{}

	return cDes
}

func canonicalizeInstanceAvailableVersionSlice(des, initial []InstanceAvailableVersion, opts ...dcl.ApplyOption) []InstanceAvailableVersion {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InstanceAvailableVersion, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInstanceAvailableVersion(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InstanceAvailableVersion, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInstanceAvailableVersion(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInstanceAvailableVersion(c *Client, des, nw *InstanceAvailableVersion) *InstanceAvailableVersion {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InstanceAvailableVersion while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.VersionNumber, nw.VersionNumber) {
		nw.VersionNumber = des.VersionNumber
	}
	if dcl.BoolCanonicalize(des.DefaultVersion, nw.DefaultVersion) {
		nw.DefaultVersion = des.DefaultVersion
	}
	if dcl.StringArrayCanonicalize(des.AvailableFeatures, nw.AvailableFeatures) {
		nw.AvailableFeatures = des.AvailableFeatures
	}

	return nw
}

func canonicalizeNewInstanceAvailableVersionSet(c *Client, des, nw []InstanceAvailableVersion) []InstanceAvailableVersion {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InstanceAvailableVersion
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInstanceAvailableVersionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInstanceAvailableVersion(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInstanceAvailableVersionSlice(c *Client, des, nw []InstanceAvailableVersion) []InstanceAvailableVersion {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InstanceAvailableVersion
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceAvailableVersion(c, &d, &n))
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

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnableStackdriverLogging, actual.EnableStackdriverLogging, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("EnableStackdriverLogging")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnableStackdriverMonitoring, actual.EnableStackdriverMonitoring, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("EnableStackdriverMonitoring")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PrivateInstance, actual.PrivateInstance, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PrivateInstance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NetworkConfig, actual.NetworkConfig, dcl.DiffInfo{ObjectFunction: compareInstanceNetworkConfigNewStyle, EmptyObject: EmptyInstanceNetworkConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NetworkConfig")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Options, actual.Options, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Options")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StateMessage, actual.StateMessage, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StateMessage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceEndpoint, actual.ServiceEndpoint, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServiceEndpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Zone, actual.Zone, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Zone")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Version, actual.Version, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("Version")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AvailableVersion, actual.AvailableVersion, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareInstanceAvailableVersionNewStyle, EmptyObject: EmptyInstanceAvailableVersion, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AvailableVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ApiEndpoint, actual.ApiEndpoint, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ApiEndpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GcsBucket, actual.GcsBucket, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GcsBucket")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.P4ServiceAccount, actual.P4ServiceAccount, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("P4ServiceAccount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TenantProjectId, actual.TenantProjectId, dcl.DiffInfo{OutputOnly: true, Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TenantProjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DataprocServiceAccount, actual.DataprocServiceAccount, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateInstanceUpdateInstanceOperation")}, fn.AddNest("DataprocServiceAccount")); len(ds) != 0 || err != nil {
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
func compareInstanceNetworkConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceNetworkConfig)
	if !ok {
		desiredNotPointer, ok := d.(InstanceNetworkConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceNetworkConfig or *InstanceNetworkConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceNetworkConfig)
	if !ok {
		actualNotPointer, ok := a.(InstanceNetworkConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceNetworkConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Network, actual.Network, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Network")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IPAllocation, actual.IPAllocation, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IpAllocation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInstanceAvailableVersionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InstanceAvailableVersion)
	if !ok {
		desiredNotPointer, ok := d.(InstanceAvailableVersion)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceAvailableVersion or *InstanceAvailableVersion", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InstanceAvailableVersion)
	if !ok {
		actualNotPointer, ok := a.(InstanceAvailableVersion)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InstanceAvailableVersion", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.VersionNumber, actual.VersionNumber, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("VersionNumber")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultVersion, actual.DefaultVersion, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DefaultVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AvailableFeatures, actual.AvailableFeatures, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AvailableFeatures")); len(ds) != 0 || err != nil {
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
	normalized.StateMessage = dcl.SelfLinkToName(r.StateMessage)
	normalized.ServiceEndpoint = dcl.SelfLinkToName(r.ServiceEndpoint)
	normalized.Zone = dcl.SelfLinkToName(r.Zone)
	normalized.Version = dcl.SelfLinkToName(r.Version)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.ApiEndpoint = dcl.SelfLinkToName(r.ApiEndpoint)
	normalized.GcsBucket = dcl.SelfLinkToName(r.GcsBucket)
	normalized.P4ServiceAccount = dcl.SelfLinkToName(r.P4ServiceAccount)
	normalized.TenantProjectId = dcl.SelfLinkToName(r.TenantProjectId)
	normalized.DataprocServiceAccount = dcl.SelfLinkToName(r.DataprocServiceAccount)
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
	if v := f.Type; dcl.ValueShouldBeSent(v) {
		m["type"] = v
	}
	if v := f.EnableStackdriverLogging; dcl.ValueShouldBeSent(v) {
		m["enableStackdriverLogging"] = v
	}
	if v := f.EnableStackdriverMonitoring; dcl.ValueShouldBeSent(v) {
		m["enableStackdriverMonitoring"] = v
	}
	if v := f.PrivateInstance; dcl.ValueShouldBeSent(v) {
		m["privateInstance"] = v
	}
	if v, err := expandInstanceNetworkConfig(c, f.NetworkConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding NetworkConfig into networkConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["networkConfig"] = v
	}
	if v := f.Labels; dcl.ValueShouldBeSent(v) {
		m["labels"] = v
	}
	if v := f.Options; dcl.ValueShouldBeSent(v) {
		m["options"] = v
	}
	if v := f.Zone; dcl.ValueShouldBeSent(v) {
		m["zone"] = v
	}
	if v := f.Version; dcl.ValueShouldBeSent(v) {
		m["version"] = v
	}
	if v := f.DisplayName; dcl.ValueShouldBeSent(v) {
		m["displayName"] = v
	}
	if v := f.DataprocServiceAccount; dcl.ValueShouldBeSent(v) {
		m["dataprocServiceAccount"] = v
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
	resultRes.Type = flattenInstanceTypeEnum(m["type"])
	resultRes.EnableStackdriverLogging = dcl.FlattenBool(m["enableStackdriverLogging"])
	resultRes.EnableStackdriverMonitoring = dcl.FlattenBool(m["enableStackdriverMonitoring"])
	resultRes.PrivateInstance = dcl.FlattenBool(m["privateInstance"])
	resultRes.NetworkConfig = flattenInstanceNetworkConfig(c, m["networkConfig"], res)
	resultRes.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	resultRes.Options = dcl.FlattenKeyValuePairs(m["options"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.State = flattenInstanceStateEnum(m["state"])
	resultRes.StateMessage = dcl.FlattenString(m["stateMessage"])
	resultRes.ServiceEndpoint = dcl.FlattenString(m["serviceEndpoint"])
	resultRes.Zone = dcl.FlattenString(m["zone"])
	resultRes.Version = dcl.FlattenString(m["version"])
	resultRes.DisplayName = dcl.FlattenString(m["displayName"])
	resultRes.AvailableVersion = flattenInstanceAvailableVersionSlice(c, m["availableVersion"], res)
	resultRes.ApiEndpoint = dcl.FlattenString(m["apiEndpoint"])
	resultRes.GcsBucket = dcl.FlattenString(m["gcsBucket"])
	resultRes.P4ServiceAccount = dcl.FlattenString(m["p4ServiceAccount"])
	resultRes.TenantProjectId = dcl.FlattenString(m["tenantProjectId"])
	resultRes.DataprocServiceAccount = dcl.FlattenString(m["dataprocServiceAccount"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// expandInstanceNetworkConfigMap expands the contents of InstanceNetworkConfig into a JSON
// request object.
func expandInstanceNetworkConfigMap(c *Client, f map[string]InstanceNetworkConfig, res *Instance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceNetworkConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceNetworkConfigSlice expands the contents of InstanceNetworkConfig into a JSON
// request object.
func expandInstanceNetworkConfigSlice(c *Client, f []InstanceNetworkConfig, res *Instance) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceNetworkConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceNetworkConfigMap flattens the contents of InstanceNetworkConfig from a JSON
// response object.
func flattenInstanceNetworkConfigMap(c *Client, i interface{}, res *Instance) map[string]InstanceNetworkConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceNetworkConfig{}
	}

	if len(a) == 0 {
		return map[string]InstanceNetworkConfig{}
	}

	items := make(map[string]InstanceNetworkConfig)
	for k, item := range a {
		items[k] = *flattenInstanceNetworkConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInstanceNetworkConfigSlice flattens the contents of InstanceNetworkConfig from a JSON
// response object.
func flattenInstanceNetworkConfigSlice(c *Client, i interface{}, res *Instance) []InstanceNetworkConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceNetworkConfig{}
	}

	if len(a) == 0 {
		return []InstanceNetworkConfig{}
	}

	items := make([]InstanceNetworkConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceNetworkConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInstanceNetworkConfig expands an instance of InstanceNetworkConfig into a JSON
// request object.
func expandInstanceNetworkConfig(c *Client, f *InstanceNetworkConfig, res *Instance) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Network; !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v := f.IPAllocation; !dcl.IsEmptyValueIndirect(v) {
		m["ipAllocation"] = v
	}

	return m, nil
}

// flattenInstanceNetworkConfig flattens an instance of InstanceNetworkConfig from a JSON
// response object.
func flattenInstanceNetworkConfig(c *Client, i interface{}, res *Instance) *InstanceNetworkConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceNetworkConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceNetworkConfig
	}
	r.Network = dcl.FlattenString(m["network"])
	r.IPAllocation = dcl.FlattenString(m["ipAllocation"])

	return r
}

// expandInstanceAvailableVersionMap expands the contents of InstanceAvailableVersion into a JSON
// request object.
func expandInstanceAvailableVersionMap(c *Client, f map[string]InstanceAvailableVersion, res *Instance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceAvailableVersion(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceAvailableVersionSlice expands the contents of InstanceAvailableVersion into a JSON
// request object.
func expandInstanceAvailableVersionSlice(c *Client, f []InstanceAvailableVersion, res *Instance) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceAvailableVersion(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceAvailableVersionMap flattens the contents of InstanceAvailableVersion from a JSON
// response object.
func flattenInstanceAvailableVersionMap(c *Client, i interface{}, res *Instance) map[string]InstanceAvailableVersion {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceAvailableVersion{}
	}

	if len(a) == 0 {
		return map[string]InstanceAvailableVersion{}
	}

	items := make(map[string]InstanceAvailableVersion)
	for k, item := range a {
		items[k] = *flattenInstanceAvailableVersion(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInstanceAvailableVersionSlice flattens the contents of InstanceAvailableVersion from a JSON
// response object.
func flattenInstanceAvailableVersionSlice(c *Client, i interface{}, res *Instance) []InstanceAvailableVersion {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceAvailableVersion{}
	}

	if len(a) == 0 {
		return []InstanceAvailableVersion{}
	}

	items := make([]InstanceAvailableVersion, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceAvailableVersion(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInstanceAvailableVersion expands an instance of InstanceAvailableVersion into a JSON
// request object.
func expandInstanceAvailableVersion(c *Client, f *InstanceAvailableVersion, res *Instance) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenInstanceAvailableVersion flattens an instance of InstanceAvailableVersion from a JSON
// response object.
func flattenInstanceAvailableVersion(c *Client, i interface{}, res *Instance) *InstanceAvailableVersion {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceAvailableVersion{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInstanceAvailableVersion
	}
	r.VersionNumber = dcl.FlattenString(m["versionNumber"])
	r.DefaultVersion = dcl.FlattenBool(m["defaultVersion"])
	r.AvailableFeatures = dcl.FlattenStringSlice(m["availableFeatures"])

	return r
}

// flattenInstanceTypeEnumMap flattens the contents of InstanceTypeEnum from a JSON
// response object.
func flattenInstanceTypeEnumMap(c *Client, i interface{}, res *Instance) map[string]InstanceTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]InstanceTypeEnum{}
	}

	items := make(map[string]InstanceTypeEnum)
	for k, item := range a {
		items[k] = *flattenInstanceTypeEnum(item.(interface{}))
	}

	return items
}

// flattenInstanceTypeEnumSlice flattens the contents of InstanceTypeEnum from a JSON
// response object.
func flattenInstanceTypeEnumSlice(c *Client, i interface{}, res *Instance) []InstanceTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceTypeEnum{}
	}

	if len(a) == 0 {
		return []InstanceTypeEnum{}
	}

	items := make([]InstanceTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceTypeEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceTypeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceTypeEnum with the same value as that string.
func flattenInstanceTypeEnum(i interface{}) *InstanceTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return InstanceTypeEnumRef(s)
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

	case "updateInstanceUpdateInstanceOperation":
		return &updateInstanceUpdateInstanceOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractInstanceFields(r *Instance) error {
	vNetworkConfig := r.NetworkConfig
	if vNetworkConfig == nil {
		// note: explicitly not the empty object.
		vNetworkConfig = &InstanceNetworkConfig{}
	}
	if err := extractInstanceNetworkConfigFields(r, vNetworkConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vNetworkConfig) {
		r.NetworkConfig = vNetworkConfig
	}
	return nil
}
func extractInstanceNetworkConfigFields(r *Instance, o *InstanceNetworkConfig) error {
	return nil
}
func extractInstanceAvailableVersionFields(r *Instance, o *InstanceAvailableVersion) error {
	return nil
}

func postReadExtractInstanceFields(r *Instance) error {
	vNetworkConfig := r.NetworkConfig
	if vNetworkConfig == nil {
		// note: explicitly not the empty object.
		vNetworkConfig = &InstanceNetworkConfig{}
	}
	if err := postReadExtractInstanceNetworkConfigFields(r, vNetworkConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vNetworkConfig) {
		r.NetworkConfig = vNetworkConfig
	}
	return nil
}
func postReadExtractInstanceNetworkConfigFields(r *Instance, o *InstanceNetworkConfig) error {
	return nil
}
func postReadExtractInstanceAvailableVersionFields(r *Instance, o *InstanceAvailableVersion) error {
	return nil
}
