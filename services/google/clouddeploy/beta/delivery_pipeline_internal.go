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

func (r *DeliveryPipeline) validate() error {

	if err := dcl.RequiredParameter(r.Name, "Name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.SerialPipeline) {
		if err := r.SerialPipeline.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Condition) {
		if err := r.Condition.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DeliveryPipelineSerialPipeline) validate() error {
	return nil
}
func (r *DeliveryPipelineSerialPipelineStages) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Strategy) {
		if err := r.Strategy.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DeliveryPipelineSerialPipelineStagesStrategy) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Standard) {
		if err := r.Standard.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Canary) {
		if err := r.Canary.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DeliveryPipelineSerialPipelineStagesStrategyStandard) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Predeploy) {
		if err := r.Predeploy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Postdeploy) {
		if err := r.Postdeploy.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy) validate() error {
	return nil
}
func (r *DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy) validate() error {
	return nil
}
func (r *DeliveryPipelineSerialPipelineStagesStrategyCanary) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"CanaryDeployment", "CustomCanaryDeployment"}, r.CanaryDeployment, r.CustomCanaryDeployment); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.RuntimeConfig) {
		if err := r.RuntimeConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CanaryDeployment) {
		if err := r.CanaryDeployment.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CustomCanaryDeployment) {
		if err := r.CustomCanaryDeployment.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"Kubernetes", "CloudRun"}, r.Kubernetes, r.CloudRun); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Kubernetes) {
		if err := r.Kubernetes.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CloudRun) {
		if err := r.CloudRun.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"GatewayServiceMesh", "ServiceNetworking"}, r.GatewayServiceMesh, r.ServiceNetworking); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.GatewayServiceMesh) {
		if err := r.GatewayServiceMesh.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ServiceNetworking) {
		if err := r.ServiceNetworking.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh) validate() error {
	if err := dcl.Required(r, "httpRoute"); err != nil {
		return err
	}
	if err := dcl.Required(r, "service"); err != nil {
		return err
	}
	if err := dcl.Required(r, "deployment"); err != nil {
		return err
	}
	return nil
}
func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking) validate() error {
	if err := dcl.Required(r, "service"); err != nil {
		return err
	}
	if err := dcl.Required(r, "deployment"); err != nil {
		return err
	}
	return nil
}
func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun) validate() error {
	return nil
}
func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment) validate() error {
	if err := dcl.Required(r, "percentages"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Predeploy) {
		if err := r.Predeploy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Postdeploy) {
		if err := r.Postdeploy.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy) validate() error {
	return nil
}
func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy) validate() error {
	return nil
}
func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment) validate() error {
	if err := dcl.Required(r, "phaseConfigs"); err != nil {
		return err
	}
	return nil
}
func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs) validate() error {
	if err := dcl.Required(r, "phaseId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "percentage"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Predeploy) {
		if err := r.Predeploy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Postdeploy) {
		if err := r.Postdeploy.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy) validate() error {
	return nil
}
func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy) validate() error {
	return nil
}
func (r *DeliveryPipelineSerialPipelineStagesDeployParameters) validate() error {
	if err := dcl.Required(r, "values"); err != nil {
		return err
	}
	return nil
}
func (r *DeliveryPipelineCondition) validate() error {
	if !dcl.IsEmptyValueIndirect(r.PipelineReadyCondition) {
		if err := r.PipelineReadyCondition.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.TargetsPresentCondition) {
		if err := r.TargetsPresentCondition.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.TargetsTypeCondition) {
		if err := r.TargetsTypeCondition.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DeliveryPipelineConditionPipelineReadyCondition) validate() error {
	return nil
}
func (r *DeliveryPipelineConditionTargetsPresentCondition) validate() error {
	return nil
}
func (r *DeliveryPipelineConditionTargetsTypeCondition) validate() error {
	return nil
}
func (r *DeliveryPipeline) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://clouddeploy.googleapis.com/v1/", params)
}

func (r *DeliveryPipeline) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/deliveryPipelines/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *DeliveryPipeline) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/deliveryPipelines", nr.basePath(), userBasePath, params), nil

}

func (r *DeliveryPipeline) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/deliveryPipelines?deliveryPipelineId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *DeliveryPipeline) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/deliveryPipelines/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *DeliveryPipeline) SetPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{}
	return dcl.URL("", nr.basePath(), userBasePath, fields)
}

func (r *DeliveryPipeline) SetPolicyVerb() string {
	return ""
}

func (r *DeliveryPipeline) getPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{}
	return dcl.URL("", nr.basePath(), userBasePath, fields)
}

func (r *DeliveryPipeline) IAMPolicyVersion() int {
	return 3
}

// deliveryPipelineApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type deliveryPipelineApiOperation interface {
	do(context.Context, *DeliveryPipeline, *Client) error
}

// newUpdateDeliveryPipelineUpdateDeliveryPipelineRequest creates a request for an
// DeliveryPipeline resource's UpdateDeliveryPipeline update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateDeliveryPipelineUpdateDeliveryPipelineRequest(ctx context.Context, f *DeliveryPipeline, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Annotations; !dcl.IsEmptyValueIndirect(v) {
		req["annotations"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v, err := expandDeliveryPipelineSerialPipeline(c, f.SerialPipeline, res); err != nil {
		return nil, fmt.Errorf("error expanding SerialPipeline into serialPipeline: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["serialPipeline"] = v
	}
	if v, err := expandDeliveryPipelineCondition(c, f.Condition, res); err != nil {
		return nil, fmt.Errorf("error expanding Condition into condition: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["condition"] = v
	}
	if v := f.Suspended; !dcl.IsEmptyValueIndirect(v) {
		req["suspended"] = v
	}
	b, err := c.getDeliveryPipelineRaw(ctx, f)
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
	req["name"] = fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", *f.Project, *f.Location, *f.Name)

	return req, nil
}

// marshalUpdateDeliveryPipelineUpdateDeliveryPipelineRequest converts the update into
// the final JSON request body.
func marshalUpdateDeliveryPipelineUpdateDeliveryPipelineRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateDeliveryPipelineUpdateDeliveryPipelineOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateDeliveryPipelineUpdateDeliveryPipelineOperation) do(ctx context.Context, r *DeliveryPipeline, c *Client) error {
	_, err := c.GetDeliveryPipeline(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateDeliveryPipeline")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateDeliveryPipelineUpdateDeliveryPipelineRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateDeliveryPipelineUpdateDeliveryPipelineRequest(c, req)
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

func (c *Client) listDeliveryPipelineRaw(ctx context.Context, r *DeliveryPipeline, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != DeliveryPipelineMaxPage {
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

type listDeliveryPipelineOperation struct {
	DeliveryPipelines []map[string]interface{} `json:"deliveryPipelines"`
	Token             string                   `json:"nextPageToken"`
}

func (c *Client) listDeliveryPipeline(ctx context.Context, r *DeliveryPipeline, pageToken string, pageSize int32) ([]*DeliveryPipeline, string, error) {
	b, err := c.listDeliveryPipelineRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listDeliveryPipelineOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*DeliveryPipeline
	for _, v := range m.DeliveryPipelines {
		res, err := unmarshalMapDeliveryPipeline(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllDeliveryPipeline(ctx context.Context, f func(*DeliveryPipeline) bool, resources []*DeliveryPipeline) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteDeliveryPipeline(ctx, res)
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

type deleteDeliveryPipelineOperation struct{}

func (op *deleteDeliveryPipelineOperation) do(ctx context.Context, r *DeliveryPipeline, c *Client) error {
	r, err := c.GetDeliveryPipeline(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "DeliveryPipeline not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetDeliveryPipeline checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	u, err = dcl.AddQueryParams(u, map[string]string{"force": "true"})
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
		_, err := c.GetDeliveryPipeline(ctx, r)
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
type createDeliveryPipelineOperation struct {
	response map[string]interface{}
}

func (op *createDeliveryPipelineOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createDeliveryPipelineOperation) do(ctx context.Context, r *DeliveryPipeline, c *Client) error {
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

	if _, err := c.GetDeliveryPipeline(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getDeliveryPipelineRaw(ctx context.Context, r *DeliveryPipeline) ([]byte, error) {

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

func (c *Client) deliveryPipelineDiffsForRawDesired(ctx context.Context, rawDesired *DeliveryPipeline, opts ...dcl.ApplyOption) (initial, desired *DeliveryPipeline, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *DeliveryPipeline
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*DeliveryPipeline); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected DeliveryPipeline, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetDeliveryPipeline(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a DeliveryPipeline resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve DeliveryPipeline resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that DeliveryPipeline resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeDeliveryPipelineDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for DeliveryPipeline: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for DeliveryPipeline: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractDeliveryPipelineFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeDeliveryPipelineInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for DeliveryPipeline: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeDeliveryPipelineDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for DeliveryPipeline: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffDeliveryPipeline(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeDeliveryPipelineInitialState(rawInitial, rawDesired *DeliveryPipeline) (*DeliveryPipeline, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeDeliveryPipelineDesiredState(rawDesired, rawInitial *DeliveryPipeline, opts ...dcl.ApplyOption) (*DeliveryPipeline, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.SerialPipeline = canonicalizeDeliveryPipelineSerialPipeline(rawDesired.SerialPipeline, nil, opts...)
		rawDesired.Condition = canonicalizeDeliveryPipelineCondition(rawDesired.Condition, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &DeliveryPipeline{}
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
	if dcl.IsZeroValue(rawDesired.Annotations) || (dcl.IsEmptyValueIndirect(rawDesired.Annotations) && dcl.IsEmptyValueIndirect(rawInitial.Annotations)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Annotations = rawInitial.Annotations
	} else {
		canonicalDesired.Annotations = rawDesired.Annotations
	}
	if dcl.IsZeroValue(rawDesired.Labels) || (dcl.IsEmptyValueIndirect(rawDesired.Labels) && dcl.IsEmptyValueIndirect(rawInitial.Labels)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	canonicalDesired.SerialPipeline = canonicalizeDeliveryPipelineSerialPipeline(rawDesired.SerialPipeline, rawInitial.SerialPipeline, opts...)
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
	if dcl.BoolCanonicalize(rawDesired.Suspended, rawInitial.Suspended) {
		canonicalDesired.Suspended = rawInitial.Suspended
	} else {
		canonicalDesired.Suspended = rawDesired.Suspended
	}
	return canonicalDesired, nil
}

func canonicalizeDeliveryPipelineNewState(c *Client, rawNew, rawDesired *DeliveryPipeline) (*DeliveryPipeline, error) {

	rawNew.Name = rawDesired.Name

	if dcl.IsEmptyValueIndirect(rawNew.Uid) && dcl.IsEmptyValueIndirect(rawDesired.Uid) {
		rawNew.Uid = rawDesired.Uid
	} else {
		if dcl.StringCanonicalize(rawDesired.Uid, rawNew.Uid) {
			rawNew.Uid = rawDesired.Uid
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Annotations) && dcl.IsEmptyValueIndirect(rawDesired.Annotations) {
		rawNew.Annotations = rawDesired.Annotations
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
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

	if dcl.IsEmptyValueIndirect(rawNew.SerialPipeline) && dcl.IsEmptyValueIndirect(rawDesired.SerialPipeline) {
		rawNew.SerialPipeline = rawDesired.SerialPipeline
	} else {
		rawNew.SerialPipeline = canonicalizeNewDeliveryPipelineSerialPipeline(c, rawDesired.SerialPipeline, rawNew.SerialPipeline)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Condition) && dcl.IsEmptyValueIndirect(rawDesired.Condition) {
		rawNew.Condition = rawDesired.Condition
	} else {
		rawNew.Condition = canonicalizeNewDeliveryPipelineCondition(c, rawDesired.Condition, rawNew.Condition)
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

	if dcl.IsEmptyValueIndirect(rawNew.Suspended) && dcl.IsEmptyValueIndirect(rawDesired.Suspended) {
		rawNew.Suspended = rawDesired.Suspended
	} else {
		if dcl.BoolCanonicalize(rawDesired.Suspended, rawNew.Suspended) {
			rawNew.Suspended = rawDesired.Suspended
		}
	}

	return rawNew, nil
}

func canonicalizeDeliveryPipelineSerialPipeline(des, initial *DeliveryPipelineSerialPipeline, opts ...dcl.ApplyOption) *DeliveryPipelineSerialPipeline {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DeliveryPipelineSerialPipeline{}

	cDes.Stages = canonicalizeDeliveryPipelineSerialPipelineStagesSlice(des.Stages, initial.Stages, opts...)

	return cDes
}

func canonicalizeDeliveryPipelineSerialPipelineSlice(des, initial []DeliveryPipelineSerialPipeline, opts ...dcl.ApplyOption) []DeliveryPipelineSerialPipeline {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DeliveryPipelineSerialPipeline, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDeliveryPipelineSerialPipeline(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DeliveryPipelineSerialPipeline, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDeliveryPipelineSerialPipeline(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDeliveryPipelineSerialPipeline(c *Client, des, nw *DeliveryPipelineSerialPipeline) *DeliveryPipelineSerialPipeline {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DeliveryPipelineSerialPipeline while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Stages = canonicalizeNewDeliveryPipelineSerialPipelineStagesSlice(c, des.Stages, nw.Stages)

	return nw
}

func canonicalizeNewDeliveryPipelineSerialPipelineSet(c *Client, des, nw []DeliveryPipelineSerialPipeline) []DeliveryPipelineSerialPipeline {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DeliveryPipelineSerialPipeline
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDeliveryPipelineSerialPipelineNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDeliveryPipelineSerialPipeline(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDeliveryPipelineSerialPipelineSlice(c *Client, des, nw []DeliveryPipelineSerialPipeline) []DeliveryPipelineSerialPipeline {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DeliveryPipelineSerialPipeline
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDeliveryPipelineSerialPipeline(c, &d, &n))
	}

	return items
}

func canonicalizeDeliveryPipelineSerialPipelineStages(des, initial *DeliveryPipelineSerialPipelineStages, opts ...dcl.ApplyOption) *DeliveryPipelineSerialPipelineStages {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DeliveryPipelineSerialPipelineStages{}

	if dcl.StringCanonicalize(des.TargetId, initial.TargetId) || dcl.IsZeroValue(des.TargetId) {
		cDes.TargetId = initial.TargetId
	} else {
		cDes.TargetId = des.TargetId
	}
	if dcl.StringArrayCanonicalize(des.Profiles, initial.Profiles) {
		cDes.Profiles = initial.Profiles
	} else {
		cDes.Profiles = des.Profiles
	}
	cDes.Strategy = canonicalizeDeliveryPipelineSerialPipelineStagesStrategy(des.Strategy, initial.Strategy, opts...)
	cDes.DeployParameters = canonicalizeDeliveryPipelineSerialPipelineStagesDeployParametersSlice(des.DeployParameters, initial.DeployParameters, opts...)

	return cDes
}

func canonicalizeDeliveryPipelineSerialPipelineStagesSlice(des, initial []DeliveryPipelineSerialPipelineStages, opts ...dcl.ApplyOption) []DeliveryPipelineSerialPipelineStages {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DeliveryPipelineSerialPipelineStages, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDeliveryPipelineSerialPipelineStages(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DeliveryPipelineSerialPipelineStages, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDeliveryPipelineSerialPipelineStages(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDeliveryPipelineSerialPipelineStages(c *Client, des, nw *DeliveryPipelineSerialPipelineStages) *DeliveryPipelineSerialPipelineStages {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DeliveryPipelineSerialPipelineStages while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.TargetId, nw.TargetId) {
		nw.TargetId = des.TargetId
	}
	if dcl.StringArrayCanonicalize(des.Profiles, nw.Profiles) {
		nw.Profiles = des.Profiles
	}
	nw.Strategy = canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategy(c, des.Strategy, nw.Strategy)
	nw.DeployParameters = canonicalizeNewDeliveryPipelineSerialPipelineStagesDeployParametersSlice(c, des.DeployParameters, nw.DeployParameters)

	return nw
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesSet(c *Client, des, nw []DeliveryPipelineSerialPipelineStages) []DeliveryPipelineSerialPipelineStages {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DeliveryPipelineSerialPipelineStages
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDeliveryPipelineSerialPipelineStagesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStages(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesSlice(c *Client, des, nw []DeliveryPipelineSerialPipelineStages) []DeliveryPipelineSerialPipelineStages {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DeliveryPipelineSerialPipelineStages
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStages(c, &d, &n))
	}

	return items
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategy(des, initial *DeliveryPipelineSerialPipelineStagesStrategy, opts ...dcl.ApplyOption) *DeliveryPipelineSerialPipelineStagesStrategy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DeliveryPipelineSerialPipelineStagesStrategy{}

	cDes.Standard = canonicalizeDeliveryPipelineSerialPipelineStagesStrategyStandard(des.Standard, initial.Standard, opts...)
	cDes.Canary = canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanary(des.Canary, initial.Canary, opts...)

	return cDes
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategySlice(des, initial []DeliveryPipelineSerialPipelineStagesStrategy, opts ...dcl.ApplyOption) []DeliveryPipelineSerialPipelineStagesStrategy {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DeliveryPipelineSerialPipelineStagesStrategy, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategy(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategy, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategy(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategy(c *Client, des, nw *DeliveryPipelineSerialPipelineStagesStrategy) *DeliveryPipelineSerialPipelineStagesStrategy {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DeliveryPipelineSerialPipelineStagesStrategy while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Standard = canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyStandard(c, des.Standard, nw.Standard)
	nw.Canary = canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanary(c, des.Canary, nw.Canary)

	return nw
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategySet(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategy) []DeliveryPipelineSerialPipelineStagesStrategy {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DeliveryPipelineSerialPipelineStagesStrategy
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDeliveryPipelineSerialPipelineStagesStrategyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategy(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategySlice(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategy) []DeliveryPipelineSerialPipelineStagesStrategy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DeliveryPipelineSerialPipelineStagesStrategy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategy(c, &d, &n))
	}

	return items
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyStandard(des, initial *DeliveryPipelineSerialPipelineStagesStrategyStandard, opts ...dcl.ApplyOption) *DeliveryPipelineSerialPipelineStagesStrategyStandard {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DeliveryPipelineSerialPipelineStagesStrategyStandard{}

	if dcl.BoolCanonicalize(des.Verify, initial.Verify) || dcl.IsZeroValue(des.Verify) {
		cDes.Verify = initial.Verify
	} else {
		cDes.Verify = des.Verify
	}
	cDes.Predeploy = canonicalizeDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy(des.Predeploy, initial.Predeploy, opts...)
	cDes.Postdeploy = canonicalizeDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy(des.Postdeploy, initial.Postdeploy, opts...)

	return cDes
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyStandardSlice(des, initial []DeliveryPipelineSerialPipelineStagesStrategyStandard, opts ...dcl.ApplyOption) []DeliveryPipelineSerialPipelineStagesStrategyStandard {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DeliveryPipelineSerialPipelineStagesStrategyStandard, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyStandard(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyStandard, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyStandard(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyStandard(c *Client, des, nw *DeliveryPipelineSerialPipelineStagesStrategyStandard) *DeliveryPipelineSerialPipelineStagesStrategyStandard {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DeliveryPipelineSerialPipelineStagesStrategyStandard while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.Verify, nw.Verify) {
		nw.Verify = des.Verify
	}
	nw.Predeploy = canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy(c, des.Predeploy, nw.Predeploy)
	nw.Postdeploy = canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy(c, des.Postdeploy, nw.Postdeploy)

	return nw
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyStandardSet(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyStandard) []DeliveryPipelineSerialPipelineStagesStrategyStandard {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DeliveryPipelineSerialPipelineStagesStrategyStandard
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDeliveryPipelineSerialPipelineStagesStrategyStandardNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyStandard(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyStandardSlice(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyStandard) []DeliveryPipelineSerialPipelineStagesStrategyStandard {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DeliveryPipelineSerialPipelineStagesStrategyStandard
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyStandard(c, &d, &n))
	}

	return items
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy(des, initial *DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy, opts ...dcl.ApplyOption) *DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy{}

	if dcl.StringArrayCanonicalize(des.Actions, initial.Actions) {
		cDes.Actions = initial.Actions
	} else {
		cDes.Actions = des.Actions
	}

	return cDes
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploySlice(des, initial []DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy, opts ...dcl.ApplyOption) []DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy(c *Client, des, nw *DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy) *DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.Actions, nw.Actions) {
		nw.Actions = des.Actions
	}

	return nw
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploySet(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy) []DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDeliveryPipelineSerialPipelineStagesStrategyStandardPredeployNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploySlice(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy) []DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy(c, &d, &n))
	}

	return items
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy(des, initial *DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy, opts ...dcl.ApplyOption) *DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy{}

	if dcl.StringArrayCanonicalize(des.Actions, initial.Actions) {
		cDes.Actions = initial.Actions
	} else {
		cDes.Actions = des.Actions
	}

	return cDes
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploySlice(des, initial []DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy, opts ...dcl.ApplyOption) []DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy(c *Client, des, nw *DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy) *DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.Actions, nw.Actions) {
		nw.Actions = des.Actions
	}

	return nw
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploySet(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy) []DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeployNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploySlice(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy) []DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy(c, &d, &n))
	}

	return items
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanary(des, initial *DeliveryPipelineSerialPipelineStagesStrategyCanary, opts ...dcl.ApplyOption) *DeliveryPipelineSerialPipelineStagesStrategyCanary {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.CanaryDeployment != nil || (initial != nil && initial.CanaryDeployment != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.CustomCanaryDeployment) {
			des.CanaryDeployment = nil
			if initial != nil {
				initial.CanaryDeployment = nil
			}
		}
	}

	if des.CustomCanaryDeployment != nil || (initial != nil && initial.CustomCanaryDeployment != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.CanaryDeployment) {
			des.CustomCanaryDeployment = nil
			if initial != nil {
				initial.CustomCanaryDeployment = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &DeliveryPipelineSerialPipelineStagesStrategyCanary{}

	cDes.RuntimeConfig = canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig(des.RuntimeConfig, initial.RuntimeConfig, opts...)
	cDes.CanaryDeployment = canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment(des.CanaryDeployment, initial.CanaryDeployment, opts...)
	cDes.CustomCanaryDeployment = canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment(des.CustomCanaryDeployment, initial.CustomCanaryDeployment, opts...)

	return cDes
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanarySlice(des, initial []DeliveryPipelineSerialPipelineStagesStrategyCanary, opts ...dcl.ApplyOption) []DeliveryPipelineSerialPipelineStagesStrategyCanary {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanary, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanary(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanary, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanary(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanary(c *Client, des, nw *DeliveryPipelineSerialPipelineStagesStrategyCanary) *DeliveryPipelineSerialPipelineStagesStrategyCanary {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DeliveryPipelineSerialPipelineStagesStrategyCanary while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.RuntimeConfig = canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig(c, des.RuntimeConfig, nw.RuntimeConfig)
	nw.CanaryDeployment = canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment(c, des.CanaryDeployment, nw.CanaryDeployment)
	nw.CustomCanaryDeployment = canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment(c, des.CustomCanaryDeployment, nw.CustomCanaryDeployment)

	return nw
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanarySet(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyCanary) []DeliveryPipelineSerialPipelineStagesStrategyCanary {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DeliveryPipelineSerialPipelineStagesStrategyCanary
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDeliveryPipelineSerialPipelineStagesStrategyCanaryNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanary(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanarySlice(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyCanary) []DeliveryPipelineSerialPipelineStagesStrategyCanary {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DeliveryPipelineSerialPipelineStagesStrategyCanary
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanary(c, &d, &n))
	}

	return items
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig(des, initial *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig, opts ...dcl.ApplyOption) *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.Kubernetes != nil || (initial != nil && initial.Kubernetes != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.CloudRun) {
			des.Kubernetes = nil
			if initial != nil {
				initial.Kubernetes = nil
			}
		}
	}

	if des.CloudRun != nil || (initial != nil && initial.CloudRun != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Kubernetes) {
			des.CloudRun = nil
			if initial != nil {
				initial.CloudRun = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig{}

	cDes.Kubernetes = canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes(des.Kubernetes, initial.Kubernetes, opts...)
	cDes.CloudRun = canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun(des.CloudRun, initial.CloudRun, opts...)

	return cDes
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigSlice(des, initial []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig, opts ...dcl.ApplyOption) []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig(c *Client, des, nw *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig) *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Kubernetes = canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes(c, des.Kubernetes, nw.Kubernetes)
	nw.CloudRun = canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun(c, des.CloudRun, nw.CloudRun)

	return nw
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigSet(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig) []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigSlice(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig) []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig(c, &d, &n))
	}

	return items
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes(des, initial *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes, opts ...dcl.ApplyOption) *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.GatewayServiceMesh != nil || (initial != nil && initial.GatewayServiceMesh != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.ServiceNetworking) {
			des.GatewayServiceMesh = nil
			if initial != nil {
				initial.GatewayServiceMesh = nil
			}
		}
	}

	if des.ServiceNetworking != nil || (initial != nil && initial.ServiceNetworking != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.GatewayServiceMesh) {
			des.ServiceNetworking = nil
			if initial != nil {
				initial.ServiceNetworking = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes{}

	cDes.GatewayServiceMesh = canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh(des.GatewayServiceMesh, initial.GatewayServiceMesh, opts...)
	cDes.ServiceNetworking = canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking(des.ServiceNetworking, initial.ServiceNetworking, opts...)

	return cDes
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesSlice(des, initial []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes, opts ...dcl.ApplyOption) []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes(c *Client, des, nw *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes) *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.GatewayServiceMesh = canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh(c, des.GatewayServiceMesh, nw.GatewayServiceMesh)
	nw.ServiceNetworking = canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking(c, des.ServiceNetworking, nw.ServiceNetworking)

	return nw
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesSet(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes) []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesSlice(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes) []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes(c, &d, &n))
	}

	return items
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh(des, initial *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh, opts ...dcl.ApplyOption) *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh{}

	if dcl.StringCanonicalize(des.HttpRoute, initial.HttpRoute) || dcl.IsZeroValue(des.HttpRoute) {
		cDes.HttpRoute = initial.HttpRoute
	} else {
		cDes.HttpRoute = des.HttpRoute
	}
	if dcl.StringCanonicalize(des.Service, initial.Service) || dcl.IsZeroValue(des.Service) {
		cDes.Service = initial.Service
	} else {
		cDes.Service = des.Service
	}
	if dcl.StringCanonicalize(des.Deployment, initial.Deployment) || dcl.IsZeroValue(des.Deployment) {
		cDes.Deployment = initial.Deployment
	} else {
		cDes.Deployment = des.Deployment
	}
	if dcl.StringCanonicalize(des.RouteUpdateWaitTime, initial.RouteUpdateWaitTime) || dcl.IsZeroValue(des.RouteUpdateWaitTime) {
		cDes.RouteUpdateWaitTime = initial.RouteUpdateWaitTime
	} else {
		cDes.RouteUpdateWaitTime = des.RouteUpdateWaitTime
	}
	if dcl.StringCanonicalize(des.StableCutbackDuration, initial.StableCutbackDuration) || dcl.IsZeroValue(des.StableCutbackDuration) {
		cDes.StableCutbackDuration = initial.StableCutbackDuration
	} else {
		cDes.StableCutbackDuration = des.StableCutbackDuration
	}

	return cDes
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshSlice(des, initial []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh, opts ...dcl.ApplyOption) []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh(c *Client, des, nw *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh) *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.HttpRoute, nw.HttpRoute) {
		nw.HttpRoute = des.HttpRoute
	}
	if dcl.StringCanonicalize(des.Service, nw.Service) {
		nw.Service = des.Service
	}
	if dcl.StringCanonicalize(des.Deployment, nw.Deployment) {
		nw.Deployment = des.Deployment
	}
	if dcl.StringCanonicalize(des.RouteUpdateWaitTime, nw.RouteUpdateWaitTime) {
		nw.RouteUpdateWaitTime = des.RouteUpdateWaitTime
	}
	if dcl.StringCanonicalize(des.StableCutbackDuration, nw.StableCutbackDuration) {
		nw.StableCutbackDuration = des.StableCutbackDuration
	}

	return nw
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshSet(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh) []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshSlice(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh) []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh(c, &d, &n))
	}

	return items
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking(des, initial *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking, opts ...dcl.ApplyOption) *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking{}

	if dcl.StringCanonicalize(des.Service, initial.Service) || dcl.IsZeroValue(des.Service) {
		cDes.Service = initial.Service
	} else {
		cDes.Service = des.Service
	}
	if dcl.StringCanonicalize(des.Deployment, initial.Deployment) || dcl.IsZeroValue(des.Deployment) {
		cDes.Deployment = initial.Deployment
	} else {
		cDes.Deployment = des.Deployment
	}
	if dcl.BoolCanonicalize(des.DisablePodOverprovisioning, initial.DisablePodOverprovisioning) || dcl.IsZeroValue(des.DisablePodOverprovisioning) {
		cDes.DisablePodOverprovisioning = initial.DisablePodOverprovisioning
	} else {
		cDes.DisablePodOverprovisioning = des.DisablePodOverprovisioning
	}

	return cDes
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingSlice(des, initial []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking, opts ...dcl.ApplyOption) []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking(c *Client, des, nw *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking) *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Service, nw.Service) {
		nw.Service = des.Service
	}
	if dcl.StringCanonicalize(des.Deployment, nw.Deployment) {
		nw.Deployment = des.Deployment
	}
	if dcl.BoolCanonicalize(des.DisablePodOverprovisioning, nw.DisablePodOverprovisioning) {
		nw.DisablePodOverprovisioning = des.DisablePodOverprovisioning
	}

	return nw
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingSet(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking) []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingSlice(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking) []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking(c, &d, &n))
	}

	return items
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun(des, initial *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun, opts ...dcl.ApplyOption) *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun{}

	if dcl.BoolCanonicalize(des.AutomaticTrafficControl, initial.AutomaticTrafficControl) || dcl.IsZeroValue(des.AutomaticTrafficControl) {
		cDes.AutomaticTrafficControl = initial.AutomaticTrafficControl
	} else {
		cDes.AutomaticTrafficControl = des.AutomaticTrafficControl
	}
	if dcl.StringArrayCanonicalize(des.CanaryRevisionTags, initial.CanaryRevisionTags) {
		cDes.CanaryRevisionTags = initial.CanaryRevisionTags
	} else {
		cDes.CanaryRevisionTags = des.CanaryRevisionTags
	}
	if dcl.StringArrayCanonicalize(des.PriorRevisionTags, initial.PriorRevisionTags) {
		cDes.PriorRevisionTags = initial.PriorRevisionTags
	} else {
		cDes.PriorRevisionTags = des.PriorRevisionTags
	}
	if dcl.StringArrayCanonicalize(des.StableRevisionTags, initial.StableRevisionTags) {
		cDes.StableRevisionTags = initial.StableRevisionTags
	} else {
		cDes.StableRevisionTags = des.StableRevisionTags
	}

	return cDes
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunSlice(des, initial []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun, opts ...dcl.ApplyOption) []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun(c *Client, des, nw *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun) *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.AutomaticTrafficControl, nw.AutomaticTrafficControl) {
		nw.AutomaticTrafficControl = des.AutomaticTrafficControl
	}
	if dcl.StringArrayCanonicalize(des.CanaryRevisionTags, nw.CanaryRevisionTags) {
		nw.CanaryRevisionTags = des.CanaryRevisionTags
	}
	if dcl.StringArrayCanonicalize(des.PriorRevisionTags, nw.PriorRevisionTags) {
		nw.PriorRevisionTags = des.PriorRevisionTags
	}
	if dcl.StringArrayCanonicalize(des.StableRevisionTags, nw.StableRevisionTags) {
		nw.StableRevisionTags = des.StableRevisionTags
	}

	return nw
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunSet(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun) []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunSlice(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun) []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun(c, &d, &n))
	}

	return items
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment(des, initial *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment, opts ...dcl.ApplyOption) *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment{}

	if dcl.IsZeroValue(des.Percentages) || (dcl.IsEmptyValueIndirect(des.Percentages) && dcl.IsEmptyValueIndirect(initial.Percentages)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Percentages = initial.Percentages
	} else {
		cDes.Percentages = des.Percentages
	}
	if dcl.BoolCanonicalize(des.Verify, initial.Verify) || dcl.IsZeroValue(des.Verify) {
		cDes.Verify = initial.Verify
	} else {
		cDes.Verify = des.Verify
	}
	cDes.Predeploy = canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy(des.Predeploy, initial.Predeploy, opts...)
	cDes.Postdeploy = canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy(des.Postdeploy, initial.Postdeploy, opts...)

	return cDes
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentSlice(des, initial []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment, opts ...dcl.ApplyOption) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment(c *Client, des, nw *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment) *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.Verify, nw.Verify) {
		nw.Verify = des.Verify
	}
	nw.Predeploy = canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy(c, des.Predeploy, nw.Predeploy)
	nw.Postdeploy = canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy(c, des.Postdeploy, nw.Postdeploy)

	return nw
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentSet(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentSlice(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment(c, &d, &n))
	}

	return items
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy(des, initial *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy, opts ...dcl.ApplyOption) *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy{}

	if dcl.StringArrayCanonicalize(des.Actions, initial.Actions) {
		cDes.Actions = initial.Actions
	} else {
		cDes.Actions = des.Actions
	}

	return cDes
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploySlice(des, initial []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy, opts ...dcl.ApplyOption) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy(c *Client, des, nw *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy) *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.Actions, nw.Actions) {
		nw.Actions = des.Actions
	}

	return nw
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploySet(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeployNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploySlice(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy(c, &d, &n))
	}

	return items
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy(des, initial *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy, opts ...dcl.ApplyOption) *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy{}

	if dcl.StringArrayCanonicalize(des.Actions, initial.Actions) {
		cDes.Actions = initial.Actions
	} else {
		cDes.Actions = des.Actions
	}

	return cDes
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploySlice(des, initial []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy, opts ...dcl.ApplyOption) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy(c *Client, des, nw *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy) *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.Actions, nw.Actions) {
		nw.Actions = des.Actions
	}

	return nw
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploySet(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeployNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploySlice(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy(c, &d, &n))
	}

	return items
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment(des, initial *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment, opts ...dcl.ApplyOption) *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment{}

	cDes.PhaseConfigs = canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsSlice(des.PhaseConfigs, initial.PhaseConfigs, opts...)

	return cDes
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentSlice(des, initial []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment, opts ...dcl.ApplyOption) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment(c *Client, des, nw *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment) *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.PhaseConfigs = canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsSlice(c, des.PhaseConfigs, nw.PhaseConfigs)

	return nw
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentSet(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentSlice(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment(c, &d, &n))
	}

	return items
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs(des, initial *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs, opts ...dcl.ApplyOption) *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs{}

	if dcl.StringCanonicalize(des.PhaseId, initial.PhaseId) || dcl.IsZeroValue(des.PhaseId) {
		cDes.PhaseId = initial.PhaseId
	} else {
		cDes.PhaseId = des.PhaseId
	}
	if dcl.IsZeroValue(des.Percentage) || (dcl.IsEmptyValueIndirect(des.Percentage) && dcl.IsEmptyValueIndirect(initial.Percentage)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Percentage = initial.Percentage
	} else {
		cDes.Percentage = des.Percentage
	}
	if dcl.StringArrayCanonicalize(des.Profiles, initial.Profiles) {
		cDes.Profiles = initial.Profiles
	} else {
		cDes.Profiles = des.Profiles
	}
	if dcl.BoolCanonicalize(des.Verify, initial.Verify) || dcl.IsZeroValue(des.Verify) {
		cDes.Verify = initial.Verify
	} else {
		cDes.Verify = des.Verify
	}
	cDes.Predeploy = canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy(des.Predeploy, initial.Predeploy, opts...)
	cDes.Postdeploy = canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy(des.Postdeploy, initial.Postdeploy, opts...)

	return cDes
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsSlice(des, initial []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs, opts ...dcl.ApplyOption) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs(c *Client, des, nw *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs) *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.PhaseId, nw.PhaseId) {
		nw.PhaseId = des.PhaseId
	}
	if dcl.StringArrayCanonicalize(des.Profiles, nw.Profiles) {
		nw.Profiles = des.Profiles
	}
	if dcl.BoolCanonicalize(des.Verify, nw.Verify) {
		nw.Verify = des.Verify
	}
	nw.Predeploy = canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy(c, des.Predeploy, nw.Predeploy)
	nw.Postdeploy = canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy(c, des.Postdeploy, nw.Postdeploy)

	return nw
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsSet(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsSlice(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs(c, &d, &n))
	}

	return items
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy(des, initial *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy, opts ...dcl.ApplyOption) *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy{}

	if dcl.StringArrayCanonicalize(des.Actions, initial.Actions) {
		cDes.Actions = initial.Actions
	} else {
		cDes.Actions = des.Actions
	}

	return cDes
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploySlice(des, initial []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy, opts ...dcl.ApplyOption) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy(c *Client, des, nw *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy) *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.Actions, nw.Actions) {
		nw.Actions = des.Actions
	}

	return nw
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploySet(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeployNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploySlice(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy(c, &d, &n))
	}

	return items
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy(des, initial *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy, opts ...dcl.ApplyOption) *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy{}

	if dcl.StringArrayCanonicalize(des.Actions, initial.Actions) {
		cDes.Actions = initial.Actions
	} else {
		cDes.Actions = des.Actions
	}

	return cDes
}

func canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploySlice(des, initial []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy, opts ...dcl.ApplyOption) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy(c *Client, des, nw *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy) *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.Actions, nw.Actions) {
		nw.Actions = des.Actions
	}

	return nw
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploySet(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeployNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploySlice(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy(c, &d, &n))
	}

	return items
}

func canonicalizeDeliveryPipelineSerialPipelineStagesDeployParameters(des, initial *DeliveryPipelineSerialPipelineStagesDeployParameters, opts ...dcl.ApplyOption) *DeliveryPipelineSerialPipelineStagesDeployParameters {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DeliveryPipelineSerialPipelineStagesDeployParameters{}

	if dcl.IsZeroValue(des.Values) || (dcl.IsEmptyValueIndirect(des.Values) && dcl.IsEmptyValueIndirect(initial.Values)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Values = initial.Values
	} else {
		cDes.Values = des.Values
	}
	if dcl.IsZeroValue(des.MatchTargetLabels) || (dcl.IsEmptyValueIndirect(des.MatchTargetLabels) && dcl.IsEmptyValueIndirect(initial.MatchTargetLabels)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MatchTargetLabels = initial.MatchTargetLabels
	} else {
		cDes.MatchTargetLabels = des.MatchTargetLabels
	}

	return cDes
}

func canonicalizeDeliveryPipelineSerialPipelineStagesDeployParametersSlice(des, initial []DeliveryPipelineSerialPipelineStagesDeployParameters, opts ...dcl.ApplyOption) []DeliveryPipelineSerialPipelineStagesDeployParameters {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DeliveryPipelineSerialPipelineStagesDeployParameters, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDeliveryPipelineSerialPipelineStagesDeployParameters(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DeliveryPipelineSerialPipelineStagesDeployParameters, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDeliveryPipelineSerialPipelineStagesDeployParameters(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesDeployParameters(c *Client, des, nw *DeliveryPipelineSerialPipelineStagesDeployParameters) *DeliveryPipelineSerialPipelineStagesDeployParameters {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DeliveryPipelineSerialPipelineStagesDeployParameters while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesDeployParametersSet(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesDeployParameters) []DeliveryPipelineSerialPipelineStagesDeployParameters {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DeliveryPipelineSerialPipelineStagesDeployParameters
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDeliveryPipelineSerialPipelineStagesDeployParametersNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesDeployParameters(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDeliveryPipelineSerialPipelineStagesDeployParametersSlice(c *Client, des, nw []DeliveryPipelineSerialPipelineStagesDeployParameters) []DeliveryPipelineSerialPipelineStagesDeployParameters {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DeliveryPipelineSerialPipelineStagesDeployParameters
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDeliveryPipelineSerialPipelineStagesDeployParameters(c, &d, &n))
	}

	return items
}

func canonicalizeDeliveryPipelineCondition(des, initial *DeliveryPipelineCondition, opts ...dcl.ApplyOption) *DeliveryPipelineCondition {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DeliveryPipelineCondition{}

	cDes.PipelineReadyCondition = canonicalizeDeliveryPipelineConditionPipelineReadyCondition(des.PipelineReadyCondition, initial.PipelineReadyCondition, opts...)
	cDes.TargetsPresentCondition = canonicalizeDeliveryPipelineConditionTargetsPresentCondition(des.TargetsPresentCondition, initial.TargetsPresentCondition, opts...)
	cDes.TargetsTypeCondition = canonicalizeDeliveryPipelineConditionTargetsTypeCondition(des.TargetsTypeCondition, initial.TargetsTypeCondition, opts...)

	return cDes
}

func canonicalizeDeliveryPipelineConditionSlice(des, initial []DeliveryPipelineCondition, opts ...dcl.ApplyOption) []DeliveryPipelineCondition {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DeliveryPipelineCondition, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDeliveryPipelineCondition(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DeliveryPipelineCondition, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDeliveryPipelineCondition(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDeliveryPipelineCondition(c *Client, des, nw *DeliveryPipelineCondition) *DeliveryPipelineCondition {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DeliveryPipelineCondition while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.PipelineReadyCondition = canonicalizeNewDeliveryPipelineConditionPipelineReadyCondition(c, des.PipelineReadyCondition, nw.PipelineReadyCondition)
	nw.TargetsPresentCondition = canonicalizeNewDeliveryPipelineConditionTargetsPresentCondition(c, des.TargetsPresentCondition, nw.TargetsPresentCondition)
	nw.TargetsTypeCondition = canonicalizeNewDeliveryPipelineConditionTargetsTypeCondition(c, des.TargetsTypeCondition, nw.TargetsTypeCondition)

	return nw
}

func canonicalizeNewDeliveryPipelineConditionSet(c *Client, des, nw []DeliveryPipelineCondition) []DeliveryPipelineCondition {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DeliveryPipelineCondition
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDeliveryPipelineConditionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDeliveryPipelineCondition(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDeliveryPipelineConditionSlice(c *Client, des, nw []DeliveryPipelineCondition) []DeliveryPipelineCondition {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DeliveryPipelineCondition
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDeliveryPipelineCondition(c, &d, &n))
	}

	return items
}

func canonicalizeDeliveryPipelineConditionPipelineReadyCondition(des, initial *DeliveryPipelineConditionPipelineReadyCondition, opts ...dcl.ApplyOption) *DeliveryPipelineConditionPipelineReadyCondition {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DeliveryPipelineConditionPipelineReadyCondition{}

	if dcl.BoolCanonicalize(des.Status, initial.Status) || dcl.IsZeroValue(des.Status) {
		cDes.Status = initial.Status
	} else {
		cDes.Status = des.Status
	}
	if dcl.IsZeroValue(des.UpdateTime) || (dcl.IsEmptyValueIndirect(des.UpdateTime) && dcl.IsEmptyValueIndirect(initial.UpdateTime)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.UpdateTime = initial.UpdateTime
	} else {
		cDes.UpdateTime = des.UpdateTime
	}

	return cDes
}

func canonicalizeDeliveryPipelineConditionPipelineReadyConditionSlice(des, initial []DeliveryPipelineConditionPipelineReadyCondition, opts ...dcl.ApplyOption) []DeliveryPipelineConditionPipelineReadyCondition {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DeliveryPipelineConditionPipelineReadyCondition, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDeliveryPipelineConditionPipelineReadyCondition(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DeliveryPipelineConditionPipelineReadyCondition, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDeliveryPipelineConditionPipelineReadyCondition(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDeliveryPipelineConditionPipelineReadyCondition(c *Client, des, nw *DeliveryPipelineConditionPipelineReadyCondition) *DeliveryPipelineConditionPipelineReadyCondition {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DeliveryPipelineConditionPipelineReadyCondition while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.Status, nw.Status) {
		nw.Status = des.Status
	}

	return nw
}

func canonicalizeNewDeliveryPipelineConditionPipelineReadyConditionSet(c *Client, des, nw []DeliveryPipelineConditionPipelineReadyCondition) []DeliveryPipelineConditionPipelineReadyCondition {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DeliveryPipelineConditionPipelineReadyCondition
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDeliveryPipelineConditionPipelineReadyConditionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDeliveryPipelineConditionPipelineReadyCondition(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDeliveryPipelineConditionPipelineReadyConditionSlice(c *Client, des, nw []DeliveryPipelineConditionPipelineReadyCondition) []DeliveryPipelineConditionPipelineReadyCondition {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DeliveryPipelineConditionPipelineReadyCondition
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDeliveryPipelineConditionPipelineReadyCondition(c, &d, &n))
	}

	return items
}

func canonicalizeDeliveryPipelineConditionTargetsPresentCondition(des, initial *DeliveryPipelineConditionTargetsPresentCondition, opts ...dcl.ApplyOption) *DeliveryPipelineConditionTargetsPresentCondition {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DeliveryPipelineConditionTargetsPresentCondition{}

	if dcl.BoolCanonicalize(des.Status, initial.Status) || dcl.IsZeroValue(des.Status) {
		cDes.Status = initial.Status
	} else {
		cDes.Status = des.Status
	}
	if dcl.StringArrayCanonicalize(des.MissingTargets, initial.MissingTargets) {
		cDes.MissingTargets = initial.MissingTargets
	} else {
		cDes.MissingTargets = des.MissingTargets
	}
	if dcl.IsZeroValue(des.UpdateTime) || (dcl.IsEmptyValueIndirect(des.UpdateTime) && dcl.IsEmptyValueIndirect(initial.UpdateTime)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.UpdateTime = initial.UpdateTime
	} else {
		cDes.UpdateTime = des.UpdateTime
	}

	return cDes
}

func canonicalizeDeliveryPipelineConditionTargetsPresentConditionSlice(des, initial []DeliveryPipelineConditionTargetsPresentCondition, opts ...dcl.ApplyOption) []DeliveryPipelineConditionTargetsPresentCondition {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DeliveryPipelineConditionTargetsPresentCondition, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDeliveryPipelineConditionTargetsPresentCondition(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DeliveryPipelineConditionTargetsPresentCondition, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDeliveryPipelineConditionTargetsPresentCondition(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDeliveryPipelineConditionTargetsPresentCondition(c *Client, des, nw *DeliveryPipelineConditionTargetsPresentCondition) *DeliveryPipelineConditionTargetsPresentCondition {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DeliveryPipelineConditionTargetsPresentCondition while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.Status, nw.Status) {
		nw.Status = des.Status
	}
	if dcl.StringArrayCanonicalize(des.MissingTargets, nw.MissingTargets) {
		nw.MissingTargets = des.MissingTargets
	}

	return nw
}

func canonicalizeNewDeliveryPipelineConditionTargetsPresentConditionSet(c *Client, des, nw []DeliveryPipelineConditionTargetsPresentCondition) []DeliveryPipelineConditionTargetsPresentCondition {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DeliveryPipelineConditionTargetsPresentCondition
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDeliveryPipelineConditionTargetsPresentConditionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDeliveryPipelineConditionTargetsPresentCondition(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDeliveryPipelineConditionTargetsPresentConditionSlice(c *Client, des, nw []DeliveryPipelineConditionTargetsPresentCondition) []DeliveryPipelineConditionTargetsPresentCondition {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DeliveryPipelineConditionTargetsPresentCondition
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDeliveryPipelineConditionTargetsPresentCondition(c, &d, &n))
	}

	return items
}

func canonicalizeDeliveryPipelineConditionTargetsTypeCondition(des, initial *DeliveryPipelineConditionTargetsTypeCondition, opts ...dcl.ApplyOption) *DeliveryPipelineConditionTargetsTypeCondition {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DeliveryPipelineConditionTargetsTypeCondition{}

	if dcl.BoolCanonicalize(des.Status, initial.Status) || dcl.IsZeroValue(des.Status) {
		cDes.Status = initial.Status
	} else {
		cDes.Status = des.Status
	}
	if dcl.StringCanonicalize(des.ErrorDetails, initial.ErrorDetails) || dcl.IsZeroValue(des.ErrorDetails) {
		cDes.ErrorDetails = initial.ErrorDetails
	} else {
		cDes.ErrorDetails = des.ErrorDetails
	}

	return cDes
}

func canonicalizeDeliveryPipelineConditionTargetsTypeConditionSlice(des, initial []DeliveryPipelineConditionTargetsTypeCondition, opts ...dcl.ApplyOption) []DeliveryPipelineConditionTargetsTypeCondition {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DeliveryPipelineConditionTargetsTypeCondition, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDeliveryPipelineConditionTargetsTypeCondition(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DeliveryPipelineConditionTargetsTypeCondition, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDeliveryPipelineConditionTargetsTypeCondition(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDeliveryPipelineConditionTargetsTypeCondition(c *Client, des, nw *DeliveryPipelineConditionTargetsTypeCondition) *DeliveryPipelineConditionTargetsTypeCondition {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DeliveryPipelineConditionTargetsTypeCondition while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.Status, nw.Status) {
		nw.Status = des.Status
	}
	if dcl.StringCanonicalize(des.ErrorDetails, nw.ErrorDetails) {
		nw.ErrorDetails = des.ErrorDetails
	}

	return nw
}

func canonicalizeNewDeliveryPipelineConditionTargetsTypeConditionSet(c *Client, des, nw []DeliveryPipelineConditionTargetsTypeCondition) []DeliveryPipelineConditionTargetsTypeCondition {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DeliveryPipelineConditionTargetsTypeCondition
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDeliveryPipelineConditionTargetsTypeConditionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDeliveryPipelineConditionTargetsTypeCondition(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDeliveryPipelineConditionTargetsTypeConditionSlice(c *Client, des, nw []DeliveryPipelineConditionTargetsTypeCondition) []DeliveryPipelineConditionTargetsTypeCondition {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DeliveryPipelineConditionTargetsTypeCondition
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDeliveryPipelineConditionTargetsTypeCondition(c, &d, &n))
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
func diffDeliveryPipeline(c *Client, desired, actual *DeliveryPipeline, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Uid, actual.Uid, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Uid")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Annotations, actual.Annotations, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Annotations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.SerialPipeline, actual.SerialPipeline, dcl.DiffInfo{ObjectFunction: compareDeliveryPipelineSerialPipelineNewStyle, EmptyObject: EmptyDeliveryPipelineSerialPipeline, OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("SerialPipeline")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Condition, actual.Condition, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareDeliveryPipelineConditionNewStyle, EmptyObject: EmptyDeliveryPipelineCondition, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Condition")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Suspended, actual.Suspended, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Suspended")); len(ds) != 0 || err != nil {
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
func compareDeliveryPipelineSerialPipelineNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DeliveryPipelineSerialPipeline)
	if !ok {
		desiredNotPointer, ok := d.(DeliveryPipelineSerialPipeline)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipeline or *DeliveryPipelineSerialPipeline", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DeliveryPipelineSerialPipeline)
	if !ok {
		actualNotPointer, ok := a.(DeliveryPipelineSerialPipeline)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipeline", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Stages, actual.Stages, dcl.DiffInfo{ObjectFunction: compareDeliveryPipelineSerialPipelineStagesNewStyle, EmptyObject: EmptyDeliveryPipelineSerialPipelineStages, OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Stages")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDeliveryPipelineSerialPipelineStagesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DeliveryPipelineSerialPipelineStages)
	if !ok {
		desiredNotPointer, ok := d.(DeliveryPipelineSerialPipelineStages)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStages or *DeliveryPipelineSerialPipelineStages", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DeliveryPipelineSerialPipelineStages)
	if !ok {
		actualNotPointer, ok := a.(DeliveryPipelineSerialPipelineStages)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStages", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TargetId, actual.TargetId, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("TargetId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Profiles, actual.Profiles, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Profiles")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Strategy, actual.Strategy, dcl.DiffInfo{ObjectFunction: compareDeliveryPipelineSerialPipelineStagesStrategyNewStyle, EmptyObject: EmptyDeliveryPipelineSerialPipelineStagesStrategy, OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Strategy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DeployParameters, actual.DeployParameters, dcl.DiffInfo{ObjectFunction: compareDeliveryPipelineSerialPipelineStagesDeployParametersNewStyle, EmptyObject: EmptyDeliveryPipelineSerialPipelineStagesDeployParameters, OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("DeployParameters")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDeliveryPipelineSerialPipelineStagesStrategyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DeliveryPipelineSerialPipelineStagesStrategy)
	if !ok {
		desiredNotPointer, ok := d.(DeliveryPipelineSerialPipelineStagesStrategy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategy or *DeliveryPipelineSerialPipelineStagesStrategy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DeliveryPipelineSerialPipelineStagesStrategy)
	if !ok {
		actualNotPointer, ok := a.(DeliveryPipelineSerialPipelineStagesStrategy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Standard, actual.Standard, dcl.DiffInfo{ObjectFunction: compareDeliveryPipelineSerialPipelineStagesStrategyStandardNewStyle, EmptyObject: EmptyDeliveryPipelineSerialPipelineStagesStrategyStandard, OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Standard")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Canary, actual.Canary, dcl.DiffInfo{ObjectFunction: compareDeliveryPipelineSerialPipelineStagesStrategyCanaryNewStyle, EmptyObject: EmptyDeliveryPipelineSerialPipelineStagesStrategyCanary, OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Canary")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDeliveryPipelineSerialPipelineStagesStrategyStandardNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DeliveryPipelineSerialPipelineStagesStrategyStandard)
	if !ok {
		desiredNotPointer, ok := d.(DeliveryPipelineSerialPipelineStagesStrategyStandard)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyStandard or *DeliveryPipelineSerialPipelineStagesStrategyStandard", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DeliveryPipelineSerialPipelineStagesStrategyStandard)
	if !ok {
		actualNotPointer, ok := a.(DeliveryPipelineSerialPipelineStagesStrategyStandard)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyStandard", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Verify, actual.Verify, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Verify")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Predeploy, actual.Predeploy, dcl.DiffInfo{ObjectFunction: compareDeliveryPipelineSerialPipelineStagesStrategyStandardPredeployNewStyle, EmptyObject: EmptyDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy, OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Predeploy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Postdeploy, actual.Postdeploy, dcl.DiffInfo{ObjectFunction: compareDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeployNewStyle, EmptyObject: EmptyDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy, OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Postdeploy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDeliveryPipelineSerialPipelineStagesStrategyStandardPredeployNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy)
	if !ok {
		desiredNotPointer, ok := d.(DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy or *DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy)
	if !ok {
		actualNotPointer, ok := a.(DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Actions, actual.Actions, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Actions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeployNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy)
	if !ok {
		desiredNotPointer, ok := d.(DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy or *DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy)
	if !ok {
		actualNotPointer, ok := a.(DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Actions, actual.Actions, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Actions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDeliveryPipelineSerialPipelineStagesStrategyCanaryNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DeliveryPipelineSerialPipelineStagesStrategyCanary)
	if !ok {
		desiredNotPointer, ok := d.(DeliveryPipelineSerialPipelineStagesStrategyCanary)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyCanary or *DeliveryPipelineSerialPipelineStagesStrategyCanary", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DeliveryPipelineSerialPipelineStagesStrategyCanary)
	if !ok {
		actualNotPointer, ok := a.(DeliveryPipelineSerialPipelineStagesStrategyCanary)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyCanary", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RuntimeConfig, actual.RuntimeConfig, dcl.DiffInfo{ObjectFunction: compareDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigNewStyle, EmptyObject: EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig, OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("RuntimeConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CanaryDeployment, actual.CanaryDeployment, dcl.DiffInfo{ObjectFunction: compareDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentNewStyle, EmptyObject: EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment, OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("CanaryDeployment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CustomCanaryDeployment, actual.CustomCanaryDeployment, dcl.DiffInfo{ObjectFunction: compareDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentNewStyle, EmptyObject: EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment, OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("CustomCanaryDeployment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig)
	if !ok {
		desiredNotPointer, ok := d.(DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig or *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig)
	if !ok {
		actualNotPointer, ok := a.(DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Kubernetes, actual.Kubernetes, dcl.DiffInfo{ObjectFunction: compareDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesNewStyle, EmptyObject: EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes, OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Kubernetes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CloudRun, actual.CloudRun, dcl.DiffInfo{ObjectFunction: compareDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunNewStyle, EmptyObject: EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun, OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("CloudRun")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes)
	if !ok {
		desiredNotPointer, ok := d.(DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes or *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes)
	if !ok {
		actualNotPointer, ok := a.(DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.GatewayServiceMesh, actual.GatewayServiceMesh, dcl.DiffInfo{ObjectFunction: compareDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshNewStyle, EmptyObject: EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh, OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("GatewayServiceMesh")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceNetworking, actual.ServiceNetworking, dcl.DiffInfo{ObjectFunction: compareDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingNewStyle, EmptyObject: EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking, OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("ServiceNetworking")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh)
	if !ok {
		desiredNotPointer, ok := d.(DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh or *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh)
	if !ok {
		actualNotPointer, ok := a.(DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HttpRoute, actual.HttpRoute, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("HttpRoute")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Service, actual.Service, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Service")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Deployment, actual.Deployment, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Deployment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RouteUpdateWaitTime, actual.RouteUpdateWaitTime, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("RouteUpdateWaitTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StableCutbackDuration, actual.StableCutbackDuration, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("StableCutbackDuration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking)
	if !ok {
		desiredNotPointer, ok := d.(DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking or *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking)
	if !ok {
		actualNotPointer, ok := a.(DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Service, actual.Service, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Service")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Deployment, actual.Deployment, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Deployment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisablePodOverprovisioning, actual.DisablePodOverprovisioning, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("DisablePodOverprovisioning")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun)
	if !ok {
		desiredNotPointer, ok := d.(DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun or *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun)
	if !ok {
		actualNotPointer, ok := a.(DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AutomaticTrafficControl, actual.AutomaticTrafficControl, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("AutomaticTrafficControl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CanaryRevisionTags, actual.CanaryRevisionTags, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("CanaryRevisionTags")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PriorRevisionTags, actual.PriorRevisionTags, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("PriorRevisionTags")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StableRevisionTags, actual.StableRevisionTags, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("StableRevisionTags")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment)
	if !ok {
		desiredNotPointer, ok := d.(DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment or *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment)
	if !ok {
		actualNotPointer, ok := a.(DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Percentages, actual.Percentages, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Percentages")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Verify, actual.Verify, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Verify")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Predeploy, actual.Predeploy, dcl.DiffInfo{ObjectFunction: compareDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeployNewStyle, EmptyObject: EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy, OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Predeploy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Postdeploy, actual.Postdeploy, dcl.DiffInfo{ObjectFunction: compareDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeployNewStyle, EmptyObject: EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy, OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Postdeploy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeployNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy)
	if !ok {
		desiredNotPointer, ok := d.(DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy or *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy)
	if !ok {
		actualNotPointer, ok := a.(DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Actions, actual.Actions, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Actions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeployNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy)
	if !ok {
		desiredNotPointer, ok := d.(DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy or *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy)
	if !ok {
		actualNotPointer, ok := a.(DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Actions, actual.Actions, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Actions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment)
	if !ok {
		desiredNotPointer, ok := d.(DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment or *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment)
	if !ok {
		actualNotPointer, ok := a.(DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PhaseConfigs, actual.PhaseConfigs, dcl.DiffInfo{ObjectFunction: compareDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsNewStyle, EmptyObject: EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs, OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("PhaseConfigs")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs)
	if !ok {
		desiredNotPointer, ok := d.(DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs or *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs)
	if !ok {
		actualNotPointer, ok := a.(DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PhaseId, actual.PhaseId, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("PhaseId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percentage, actual.Percentage, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Percentage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Profiles, actual.Profiles, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Profiles")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Verify, actual.Verify, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Verify")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Predeploy, actual.Predeploy, dcl.DiffInfo{ObjectFunction: compareDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeployNewStyle, EmptyObject: EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy, OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Predeploy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Postdeploy, actual.Postdeploy, dcl.DiffInfo{ObjectFunction: compareDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeployNewStyle, EmptyObject: EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy, OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Postdeploy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeployNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy)
	if !ok {
		desiredNotPointer, ok := d.(DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy or *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy)
	if !ok {
		actualNotPointer, ok := a.(DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Actions, actual.Actions, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Actions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeployNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy)
	if !ok {
		desiredNotPointer, ok := d.(DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy or *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy)
	if !ok {
		actualNotPointer, ok := a.(DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Actions, actual.Actions, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Actions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDeliveryPipelineSerialPipelineStagesDeployParametersNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DeliveryPipelineSerialPipelineStagesDeployParameters)
	if !ok {
		desiredNotPointer, ok := d.(DeliveryPipelineSerialPipelineStagesDeployParameters)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesDeployParameters or *DeliveryPipelineSerialPipelineStagesDeployParameters", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DeliveryPipelineSerialPipelineStagesDeployParameters)
	if !ok {
		actualNotPointer, ok := a.(DeliveryPipelineSerialPipelineStagesDeployParameters)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineSerialPipelineStagesDeployParameters", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Values, actual.Values, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Values")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MatchTargetLabels, actual.MatchTargetLabels, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("MatchTargetLabels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDeliveryPipelineConditionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DeliveryPipelineCondition)
	if !ok {
		desiredNotPointer, ok := d.(DeliveryPipelineCondition)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineCondition or *DeliveryPipelineCondition", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DeliveryPipelineCondition)
	if !ok {
		actualNotPointer, ok := a.(DeliveryPipelineCondition)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineCondition", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PipelineReadyCondition, actual.PipelineReadyCondition, dcl.DiffInfo{ObjectFunction: compareDeliveryPipelineConditionPipelineReadyConditionNewStyle, EmptyObject: EmptyDeliveryPipelineConditionPipelineReadyCondition, OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("PipelineReadyCondition")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TargetsPresentCondition, actual.TargetsPresentCondition, dcl.DiffInfo{ObjectFunction: compareDeliveryPipelineConditionTargetsPresentConditionNewStyle, EmptyObject: EmptyDeliveryPipelineConditionTargetsPresentCondition, OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("TargetsPresentCondition")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TargetsTypeCondition, actual.TargetsTypeCondition, dcl.DiffInfo{ObjectFunction: compareDeliveryPipelineConditionTargetsTypeConditionNewStyle, EmptyObject: EmptyDeliveryPipelineConditionTargetsTypeCondition, OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("TargetsTypeCondition")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDeliveryPipelineConditionPipelineReadyConditionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DeliveryPipelineConditionPipelineReadyCondition)
	if !ok {
		desiredNotPointer, ok := d.(DeliveryPipelineConditionPipelineReadyCondition)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineConditionPipelineReadyCondition or *DeliveryPipelineConditionPipelineReadyCondition", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DeliveryPipelineConditionPipelineReadyCondition)
	if !ok {
		actualNotPointer, ok := a.(DeliveryPipelineConditionPipelineReadyCondition)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineConditionPipelineReadyCondition", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Status, actual.Status, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Status")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDeliveryPipelineConditionTargetsPresentConditionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DeliveryPipelineConditionTargetsPresentCondition)
	if !ok {
		desiredNotPointer, ok := d.(DeliveryPipelineConditionTargetsPresentCondition)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineConditionTargetsPresentCondition or *DeliveryPipelineConditionTargetsPresentCondition", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DeliveryPipelineConditionTargetsPresentCondition)
	if !ok {
		actualNotPointer, ok := a.(DeliveryPipelineConditionTargetsPresentCondition)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineConditionTargetsPresentCondition", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Status, actual.Status, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Status")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MissingTargets, actual.MissingTargets, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("MissingTargets")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDeliveryPipelineConditionTargetsTypeConditionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DeliveryPipelineConditionTargetsTypeCondition)
	if !ok {
		desiredNotPointer, ok := d.(DeliveryPipelineConditionTargetsTypeCondition)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineConditionTargetsTypeCondition or *DeliveryPipelineConditionTargetsTypeCondition", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DeliveryPipelineConditionTargetsTypeCondition)
	if !ok {
		actualNotPointer, ok := a.(DeliveryPipelineConditionTargetsTypeCondition)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DeliveryPipelineConditionTargetsTypeCondition", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Status, actual.Status, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("Status")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ErrorDetails, actual.ErrorDetails, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDeliveryPipelineUpdateDeliveryPipelineOperation")}, fn.AddNest("ErrorDetails")); len(ds) != 0 || err != nil {
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
func (r *DeliveryPipeline) urlNormalized() *DeliveryPipeline {
	normalized := dcl.Copy(*r).(DeliveryPipeline)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Uid = dcl.SelfLinkToName(r.Uid)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *DeliveryPipeline) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateDeliveryPipeline" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/deliveryPipelines/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the DeliveryPipeline resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *DeliveryPipeline) marshal(c *Client) ([]byte, error) {
	m, err := expandDeliveryPipeline(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling DeliveryPipeline: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalDeliveryPipeline decodes JSON responses into the DeliveryPipeline resource schema.
func unmarshalDeliveryPipeline(b []byte, c *Client, res *DeliveryPipeline) (*DeliveryPipeline, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapDeliveryPipeline(m, c, res)
}

func unmarshalMapDeliveryPipeline(m map[string]interface{}, c *Client, res *DeliveryPipeline) (*DeliveryPipeline, error) {

	flattened := flattenDeliveryPipeline(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandDeliveryPipeline expands DeliveryPipeline into a JSON request object.
func expandDeliveryPipeline(c *Client, f *DeliveryPipeline) (map[string]interface{}, error) {
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
	if v := f.Annotations; dcl.ValueShouldBeSent(v) {
		m["annotations"] = v
	}
	if v := f.Labels; dcl.ValueShouldBeSent(v) {
		m["labels"] = v
	}
	if v, err := expandDeliveryPipelineSerialPipeline(c, f.SerialPipeline, res); err != nil {
		return nil, fmt.Errorf("error expanding SerialPipeline into serialPipeline: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["serialPipeline"] = v
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
	if v := f.Suspended; dcl.ValueShouldBeSent(v) {
		m["suspended"] = v
	}

	return m, nil
}

// flattenDeliveryPipeline flattens DeliveryPipeline from a JSON request object into the
// DeliveryPipeline type.
func flattenDeliveryPipeline(c *Client, i interface{}, res *DeliveryPipeline) *DeliveryPipeline {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &DeliveryPipeline{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.Uid = dcl.FlattenString(m["uid"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.Annotations = dcl.FlattenKeyValuePairs(m["annotations"])
	resultRes.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.SerialPipeline = flattenDeliveryPipelineSerialPipeline(c, m["serialPipeline"], res)
	resultRes.Condition = flattenDeliveryPipelineCondition(c, m["condition"], res)
	resultRes.Etag = dcl.FlattenString(m["etag"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])
	resultRes.Suspended = dcl.FlattenBool(m["suspended"])

	return resultRes
}

// expandDeliveryPipelineSerialPipelineMap expands the contents of DeliveryPipelineSerialPipeline into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineMap(c *Client, f map[string]DeliveryPipelineSerialPipeline, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDeliveryPipelineSerialPipeline(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDeliveryPipelineSerialPipelineSlice expands the contents of DeliveryPipelineSerialPipeline into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineSlice(c *Client, f []DeliveryPipelineSerialPipeline, res *DeliveryPipeline) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDeliveryPipelineSerialPipeline(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDeliveryPipelineSerialPipelineMap flattens the contents of DeliveryPipelineSerialPipeline from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineMap(c *Client, i interface{}, res *DeliveryPipeline) map[string]DeliveryPipelineSerialPipeline {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DeliveryPipelineSerialPipeline{}
	}

	if len(a) == 0 {
		return map[string]DeliveryPipelineSerialPipeline{}
	}

	items := make(map[string]DeliveryPipelineSerialPipeline)
	for k, item := range a {
		items[k] = *flattenDeliveryPipelineSerialPipeline(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDeliveryPipelineSerialPipelineSlice flattens the contents of DeliveryPipelineSerialPipeline from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineSlice(c *Client, i interface{}, res *DeliveryPipeline) []DeliveryPipelineSerialPipeline {
	a, ok := i.([]interface{})
	if !ok {
		return []DeliveryPipelineSerialPipeline{}
	}

	if len(a) == 0 {
		return []DeliveryPipelineSerialPipeline{}
	}

	items := make([]DeliveryPipelineSerialPipeline, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDeliveryPipelineSerialPipeline(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDeliveryPipelineSerialPipeline expands an instance of DeliveryPipelineSerialPipeline into a JSON
// request object.
func expandDeliveryPipelineSerialPipeline(c *Client, f *DeliveryPipelineSerialPipeline, res *DeliveryPipeline) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandDeliveryPipelineSerialPipelineStagesSlice(c, f.Stages, res); err != nil {
		return nil, fmt.Errorf("error expanding Stages into stages: %w", err)
	} else if v != nil {
		m["stages"] = v
	}

	return m, nil
}

// flattenDeliveryPipelineSerialPipeline flattens an instance of DeliveryPipelineSerialPipeline from a JSON
// response object.
func flattenDeliveryPipelineSerialPipeline(c *Client, i interface{}, res *DeliveryPipeline) *DeliveryPipelineSerialPipeline {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DeliveryPipelineSerialPipeline{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDeliveryPipelineSerialPipeline
	}
	r.Stages = flattenDeliveryPipelineSerialPipelineStagesSlice(c, m["stages"], res)

	return r
}

// expandDeliveryPipelineSerialPipelineStagesMap expands the contents of DeliveryPipelineSerialPipelineStages into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesMap(c *Client, f map[string]DeliveryPipelineSerialPipelineStages, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStages(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDeliveryPipelineSerialPipelineStagesSlice expands the contents of DeliveryPipelineSerialPipelineStages into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesSlice(c *Client, f []DeliveryPipelineSerialPipelineStages, res *DeliveryPipeline) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStages(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDeliveryPipelineSerialPipelineStagesMap flattens the contents of DeliveryPipelineSerialPipelineStages from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesMap(c *Client, i interface{}, res *DeliveryPipeline) map[string]DeliveryPipelineSerialPipelineStages {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DeliveryPipelineSerialPipelineStages{}
	}

	if len(a) == 0 {
		return map[string]DeliveryPipelineSerialPipelineStages{}
	}

	items := make(map[string]DeliveryPipelineSerialPipelineStages)
	for k, item := range a {
		items[k] = *flattenDeliveryPipelineSerialPipelineStages(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDeliveryPipelineSerialPipelineStagesSlice flattens the contents of DeliveryPipelineSerialPipelineStages from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesSlice(c *Client, i interface{}, res *DeliveryPipeline) []DeliveryPipelineSerialPipelineStages {
	a, ok := i.([]interface{})
	if !ok {
		return []DeliveryPipelineSerialPipelineStages{}
	}

	if len(a) == 0 {
		return []DeliveryPipelineSerialPipelineStages{}
	}

	items := make([]DeliveryPipelineSerialPipelineStages, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDeliveryPipelineSerialPipelineStages(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDeliveryPipelineSerialPipelineStages expands an instance of DeliveryPipelineSerialPipelineStages into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStages(c *Client, f *DeliveryPipelineSerialPipelineStages, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.TargetId; !dcl.IsEmptyValueIndirect(v) {
		m["targetId"] = v
	}
	if v := f.Profiles; v != nil {
		m["profiles"] = v
	}
	if v, err := expandDeliveryPipelineSerialPipelineStagesStrategy(c, f.Strategy, res); err != nil {
		return nil, fmt.Errorf("error expanding Strategy into strategy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["strategy"] = v
	}
	if v, err := expandDeliveryPipelineSerialPipelineStagesDeployParametersSlice(c, f.DeployParameters, res); err != nil {
		return nil, fmt.Errorf("error expanding DeployParameters into deployParameters: %w", err)
	} else if v != nil {
		m["deployParameters"] = v
	}

	return m, nil
}

// flattenDeliveryPipelineSerialPipelineStages flattens an instance of DeliveryPipelineSerialPipelineStages from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStages(c *Client, i interface{}, res *DeliveryPipeline) *DeliveryPipelineSerialPipelineStages {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DeliveryPipelineSerialPipelineStages{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDeliveryPipelineSerialPipelineStages
	}
	r.TargetId = dcl.FlattenString(m["targetId"])
	r.Profiles = dcl.FlattenStringSlice(m["profiles"])
	r.Strategy = flattenDeliveryPipelineSerialPipelineStagesStrategy(c, m["strategy"], res)
	r.DeployParameters = flattenDeliveryPipelineSerialPipelineStagesDeployParametersSlice(c, m["deployParameters"], res)

	return r
}

// expandDeliveryPipelineSerialPipelineStagesStrategyMap expands the contents of DeliveryPipelineSerialPipelineStagesStrategy into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyMap(c *Client, f map[string]DeliveryPipelineSerialPipelineStagesStrategy, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategy(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDeliveryPipelineSerialPipelineStagesStrategySlice expands the contents of DeliveryPipelineSerialPipelineStagesStrategy into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategySlice(c *Client, f []DeliveryPipelineSerialPipelineStagesStrategy, res *DeliveryPipeline) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategy(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyMap flattens the contents of DeliveryPipelineSerialPipelineStagesStrategy from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyMap(c *Client, i interface{}, res *DeliveryPipeline) map[string]DeliveryPipelineSerialPipelineStagesStrategy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategy{}
	}

	if len(a) == 0 {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategy{}
	}

	items := make(map[string]DeliveryPipelineSerialPipelineStagesStrategy)
	for k, item := range a {
		items[k] = *flattenDeliveryPipelineSerialPipelineStagesStrategy(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDeliveryPipelineSerialPipelineStagesStrategySlice flattens the contents of DeliveryPipelineSerialPipelineStagesStrategy from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategySlice(c *Client, i interface{}, res *DeliveryPipeline) []DeliveryPipelineSerialPipelineStagesStrategy {
	a, ok := i.([]interface{})
	if !ok {
		return []DeliveryPipelineSerialPipelineStagesStrategy{}
	}

	if len(a) == 0 {
		return []DeliveryPipelineSerialPipelineStagesStrategy{}
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDeliveryPipelineSerialPipelineStagesStrategy(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDeliveryPipelineSerialPipelineStagesStrategy expands an instance of DeliveryPipelineSerialPipelineStagesStrategy into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategy(c *Client, f *DeliveryPipelineSerialPipelineStagesStrategy, res *DeliveryPipeline) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandDeliveryPipelineSerialPipelineStagesStrategyStandard(c, f.Standard, res); err != nil {
		return nil, fmt.Errorf("error expanding Standard into standard: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["standard"] = v
	}
	if v, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanary(c, f.Canary, res); err != nil {
		return nil, fmt.Errorf("error expanding Canary into canary: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["canary"] = v
	}

	return m, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategy flattens an instance of DeliveryPipelineSerialPipelineStagesStrategy from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategy(c *Client, i interface{}, res *DeliveryPipeline) *DeliveryPipelineSerialPipelineStagesStrategy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DeliveryPipelineSerialPipelineStagesStrategy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDeliveryPipelineSerialPipelineStagesStrategy
	}
	r.Standard = flattenDeliveryPipelineSerialPipelineStagesStrategyStandard(c, m["standard"], res)
	r.Canary = flattenDeliveryPipelineSerialPipelineStagesStrategyCanary(c, m["canary"], res)

	return r
}

// expandDeliveryPipelineSerialPipelineStagesStrategyStandardMap expands the contents of DeliveryPipelineSerialPipelineStagesStrategyStandard into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyStandardMap(c *Client, f map[string]DeliveryPipelineSerialPipelineStagesStrategyStandard, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyStandard(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDeliveryPipelineSerialPipelineStagesStrategyStandardSlice expands the contents of DeliveryPipelineSerialPipelineStagesStrategyStandard into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyStandardSlice(c *Client, f []DeliveryPipelineSerialPipelineStagesStrategyStandard, res *DeliveryPipeline) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyStandard(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyStandardMap flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyStandard from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyStandardMap(c *Client, i interface{}, res *DeliveryPipeline) map[string]DeliveryPipelineSerialPipelineStagesStrategyStandard {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyStandard{}
	}

	if len(a) == 0 {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyStandard{}
	}

	items := make(map[string]DeliveryPipelineSerialPipelineStagesStrategyStandard)
	for k, item := range a {
		items[k] = *flattenDeliveryPipelineSerialPipelineStagesStrategyStandard(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyStandardSlice flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyStandard from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyStandardSlice(c *Client, i interface{}, res *DeliveryPipeline) []DeliveryPipelineSerialPipelineStagesStrategyStandard {
	a, ok := i.([]interface{})
	if !ok {
		return []DeliveryPipelineSerialPipelineStagesStrategyStandard{}
	}

	if len(a) == 0 {
		return []DeliveryPipelineSerialPipelineStagesStrategyStandard{}
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyStandard, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDeliveryPipelineSerialPipelineStagesStrategyStandard(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDeliveryPipelineSerialPipelineStagesStrategyStandard expands an instance of DeliveryPipelineSerialPipelineStagesStrategyStandard into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyStandard(c *Client, f *DeliveryPipelineSerialPipelineStagesStrategyStandard, res *DeliveryPipeline) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Verify; !dcl.IsEmptyValueIndirect(v) {
		m["verify"] = v
	}
	if v, err := expandDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy(c, f.Predeploy, res); err != nil {
		return nil, fmt.Errorf("error expanding Predeploy into predeploy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["predeploy"] = v
	}
	if v, err := expandDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy(c, f.Postdeploy, res); err != nil {
		return nil, fmt.Errorf("error expanding Postdeploy into postdeploy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["postdeploy"] = v
	}

	return m, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyStandard flattens an instance of DeliveryPipelineSerialPipelineStagesStrategyStandard from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyStandard(c *Client, i interface{}, res *DeliveryPipeline) *DeliveryPipelineSerialPipelineStagesStrategyStandard {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DeliveryPipelineSerialPipelineStagesStrategyStandard{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDeliveryPipelineSerialPipelineStagesStrategyStandard
	}
	r.Verify = dcl.FlattenBool(m["verify"])
	r.Predeploy = flattenDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy(c, m["predeploy"], res)
	r.Postdeploy = flattenDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy(c, m["postdeploy"], res)

	return r
}

// expandDeliveryPipelineSerialPipelineStagesStrategyStandardPredeployMap expands the contents of DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyStandardPredeployMap(c *Client, f map[string]DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploySlice expands the contents of DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploySlice(c *Client, f []DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy, res *DeliveryPipeline) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyStandardPredeployMap flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyStandardPredeployMap(c *Client, i interface{}, res *DeliveryPipeline) map[string]DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy{}
	}

	if len(a) == 0 {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy{}
	}

	items := make(map[string]DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy)
	for k, item := range a {
		items[k] = *flattenDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploySlice flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploySlice(c *Client, i interface{}, res *DeliveryPipeline) []DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy {
	a, ok := i.([]interface{})
	if !ok {
		return []DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy{}
	}

	if len(a) == 0 {
		return []DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy{}
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy expands an instance of DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy(c *Client, f *DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy, res *DeliveryPipeline) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Actions; v != nil {
		m["actions"] = v
	}

	return m, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy flattens an instance of DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy(c *Client, i interface{}, res *DeliveryPipeline) *DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy
	}
	r.Actions = dcl.FlattenStringSlice(m["actions"])

	return r
}

// expandDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeployMap expands the contents of DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeployMap(c *Client, f map[string]DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploySlice expands the contents of DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploySlice(c *Client, f []DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy, res *DeliveryPipeline) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeployMap flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeployMap(c *Client, i interface{}, res *DeliveryPipeline) map[string]DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy{}
	}

	if len(a) == 0 {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy{}
	}

	items := make(map[string]DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy)
	for k, item := range a {
		items[k] = *flattenDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploySlice flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploySlice(c *Client, i interface{}, res *DeliveryPipeline) []DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy {
	a, ok := i.([]interface{})
	if !ok {
		return []DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy{}
	}

	if len(a) == 0 {
		return []DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy{}
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy expands an instance of DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy(c *Client, f *DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy, res *DeliveryPipeline) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Actions; v != nil {
		m["actions"] = v
	}

	return m, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy flattens an instance of DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy(c *Client, i interface{}, res *DeliveryPipeline) *DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy
	}
	r.Actions = dcl.FlattenStringSlice(m["actions"])

	return r
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryMap expands the contents of DeliveryPipelineSerialPipelineStagesStrategyCanary into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryMap(c *Client, f map[string]DeliveryPipelineSerialPipelineStagesStrategyCanary, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanary(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanarySlice expands the contents of DeliveryPipelineSerialPipelineStagesStrategyCanary into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanarySlice(c *Client, f []DeliveryPipelineSerialPipelineStagesStrategyCanary, res *DeliveryPipeline) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanary(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryMap flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyCanary from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryMap(c *Client, i interface{}, res *DeliveryPipeline) map[string]DeliveryPipelineSerialPipelineStagesStrategyCanary {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyCanary{}
	}

	if len(a) == 0 {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyCanary{}
	}

	items := make(map[string]DeliveryPipelineSerialPipelineStagesStrategyCanary)
	for k, item := range a {
		items[k] = *flattenDeliveryPipelineSerialPipelineStagesStrategyCanary(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanarySlice flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyCanary from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanarySlice(c *Client, i interface{}, res *DeliveryPipeline) []DeliveryPipelineSerialPipelineStagesStrategyCanary {
	a, ok := i.([]interface{})
	if !ok {
		return []DeliveryPipelineSerialPipelineStagesStrategyCanary{}
	}

	if len(a) == 0 {
		return []DeliveryPipelineSerialPipelineStagesStrategyCanary{}
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanary, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDeliveryPipelineSerialPipelineStagesStrategyCanary(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanary expands an instance of DeliveryPipelineSerialPipelineStagesStrategyCanary into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanary(c *Client, f *DeliveryPipelineSerialPipelineStagesStrategyCanary, res *DeliveryPipeline) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig(c, f.RuntimeConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding RuntimeConfig into runtimeConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["runtimeConfig"] = v
	}
	if v, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment(c, f.CanaryDeployment, res); err != nil {
		return nil, fmt.Errorf("error expanding CanaryDeployment into canaryDeployment: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["canaryDeployment"] = v
	}
	if v, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment(c, f.CustomCanaryDeployment, res); err != nil {
		return nil, fmt.Errorf("error expanding CustomCanaryDeployment into customCanaryDeployment: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["customCanaryDeployment"] = v
	}

	return m, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanary flattens an instance of DeliveryPipelineSerialPipelineStagesStrategyCanary from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanary(c *Client, i interface{}, res *DeliveryPipeline) *DeliveryPipelineSerialPipelineStagesStrategyCanary {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DeliveryPipelineSerialPipelineStagesStrategyCanary{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDeliveryPipelineSerialPipelineStagesStrategyCanary
	}
	r.RuntimeConfig = flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig(c, m["runtimeConfig"], res)
	r.CanaryDeployment = flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment(c, m["canaryDeployment"], res)
	r.CustomCanaryDeployment = flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment(c, m["customCanaryDeployment"], res)

	return r
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigMap expands the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigMap(c *Client, f map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigSlice expands the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigSlice(c *Client, f []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig, res *DeliveryPipeline) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigMap flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigMap(c *Client, i interface{}, res *DeliveryPipeline) map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig{}
	}

	if len(a) == 0 {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig{}
	}

	items := make(map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig)
	for k, item := range a {
		items[k] = *flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigSlice flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigSlice(c *Client, i interface{}, res *DeliveryPipeline) []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig{}
	}

	if len(a) == 0 {
		return []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig{}
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig expands an instance of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig(c *Client, f *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig, res *DeliveryPipeline) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes(c, f.Kubernetes, res); err != nil {
		return nil, fmt.Errorf("error expanding Kubernetes into kubernetes: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["kubernetes"] = v
	}
	if v, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun(c, f.CloudRun, res); err != nil {
		return nil, fmt.Errorf("error expanding CloudRun into cloudRun: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cloudRun"] = v
	}

	return m, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig flattens an instance of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig(c *Client, i interface{}, res *DeliveryPipeline) *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig
	}
	r.Kubernetes = flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes(c, m["kubernetes"], res)
	r.CloudRun = flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun(c, m["cloudRun"], res)

	return r
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesMap expands the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesMap(c *Client, f map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesSlice expands the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesSlice(c *Client, f []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes, res *DeliveryPipeline) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesMap flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesMap(c *Client, i interface{}, res *DeliveryPipeline) map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes{}
	}

	if len(a) == 0 {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes{}
	}

	items := make(map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes)
	for k, item := range a {
		items[k] = *flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesSlice flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesSlice(c *Client, i interface{}, res *DeliveryPipeline) []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes {
	a, ok := i.([]interface{})
	if !ok {
		return []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes{}
	}

	if len(a) == 0 {
		return []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes{}
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes expands an instance of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes(c *Client, f *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes, res *DeliveryPipeline) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh(c, f.GatewayServiceMesh, res); err != nil {
		return nil, fmt.Errorf("error expanding GatewayServiceMesh into gatewayServiceMesh: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["gatewayServiceMesh"] = v
	}
	if v, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking(c, f.ServiceNetworking, res); err != nil {
		return nil, fmt.Errorf("error expanding ServiceNetworking into serviceNetworking: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["serviceNetworking"] = v
	}

	return m, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes flattens an instance of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes(c *Client, i interface{}, res *DeliveryPipeline) *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes
	}
	r.GatewayServiceMesh = flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh(c, m["gatewayServiceMesh"], res)
	r.ServiceNetworking = flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking(c, m["serviceNetworking"], res)

	return r
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshMap expands the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshMap(c *Client, f map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshSlice expands the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshSlice(c *Client, f []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh, res *DeliveryPipeline) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshMap flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshMap(c *Client, i interface{}, res *DeliveryPipeline) map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh{}
	}

	if len(a) == 0 {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh{}
	}

	items := make(map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh)
	for k, item := range a {
		items[k] = *flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshSlice flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshSlice(c *Client, i interface{}, res *DeliveryPipeline) []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh {
	a, ok := i.([]interface{})
	if !ok {
		return []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh{}
	}

	if len(a) == 0 {
		return []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh{}
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh expands an instance of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh(c *Client, f *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh, res *DeliveryPipeline) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.HttpRoute; !dcl.IsEmptyValueIndirect(v) {
		m["httpRoute"] = v
	}
	if v := f.Service; !dcl.IsEmptyValueIndirect(v) {
		m["service"] = v
	}
	if v := f.Deployment; !dcl.IsEmptyValueIndirect(v) {
		m["deployment"] = v
	}
	if v := f.RouteUpdateWaitTime; !dcl.IsEmptyValueIndirect(v) {
		m["routeUpdateWaitTime"] = v
	}
	if v := f.StableCutbackDuration; !dcl.IsEmptyValueIndirect(v) {
		m["stableCutbackDuration"] = v
	}

	return m, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh flattens an instance of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh(c *Client, i interface{}, res *DeliveryPipeline) *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh
	}
	r.HttpRoute = dcl.FlattenString(m["httpRoute"])
	r.Service = dcl.FlattenString(m["service"])
	r.Deployment = dcl.FlattenString(m["deployment"])
	r.RouteUpdateWaitTime = dcl.FlattenString(m["routeUpdateWaitTime"])
	r.StableCutbackDuration = dcl.FlattenString(m["stableCutbackDuration"])

	return r
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingMap expands the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingMap(c *Client, f map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingSlice expands the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingSlice(c *Client, f []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking, res *DeliveryPipeline) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingMap flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingMap(c *Client, i interface{}, res *DeliveryPipeline) map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking{}
	}

	if len(a) == 0 {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking{}
	}

	items := make(map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking)
	for k, item := range a {
		items[k] = *flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingSlice flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingSlice(c *Client, i interface{}, res *DeliveryPipeline) []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking {
	a, ok := i.([]interface{})
	if !ok {
		return []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking{}
	}

	if len(a) == 0 {
		return []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking{}
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking expands an instance of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking(c *Client, f *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking, res *DeliveryPipeline) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Service; !dcl.IsEmptyValueIndirect(v) {
		m["service"] = v
	}
	if v := f.Deployment; !dcl.IsEmptyValueIndirect(v) {
		m["deployment"] = v
	}
	if v := f.DisablePodOverprovisioning; !dcl.IsEmptyValueIndirect(v) {
		m["disablePodOverprovisioning"] = v
	}

	return m, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking flattens an instance of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking(c *Client, i interface{}, res *DeliveryPipeline) *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking
	}
	r.Service = dcl.FlattenString(m["service"])
	r.Deployment = dcl.FlattenString(m["deployment"])
	r.DisablePodOverprovisioning = dcl.FlattenBool(m["disablePodOverprovisioning"])

	return r
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunMap expands the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunMap(c *Client, f map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunSlice expands the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunSlice(c *Client, f []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun, res *DeliveryPipeline) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunMap flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunMap(c *Client, i interface{}, res *DeliveryPipeline) map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun{}
	}

	if len(a) == 0 {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun{}
	}

	items := make(map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun)
	for k, item := range a {
		items[k] = *flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunSlice flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunSlice(c *Client, i interface{}, res *DeliveryPipeline) []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun {
	a, ok := i.([]interface{})
	if !ok {
		return []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun{}
	}

	if len(a) == 0 {
		return []DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun{}
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun expands an instance of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun(c *Client, f *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun, res *DeliveryPipeline) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AutomaticTrafficControl; !dcl.IsEmptyValueIndirect(v) {
		m["automaticTrafficControl"] = v
	}
	if v := f.CanaryRevisionTags; v != nil {
		m["canaryRevisionTags"] = v
	}
	if v := f.PriorRevisionTags; v != nil {
		m["priorRevisionTags"] = v
	}
	if v := f.StableRevisionTags; v != nil {
		m["stableRevisionTags"] = v
	}

	return m, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun flattens an instance of DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun(c *Client, i interface{}, res *DeliveryPipeline) *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun
	}
	r.AutomaticTrafficControl = dcl.FlattenBool(m["automaticTrafficControl"])
	r.CanaryRevisionTags = dcl.FlattenStringSlice(m["canaryRevisionTags"])
	r.PriorRevisionTags = dcl.FlattenStringSlice(m["priorRevisionTags"])
	r.StableRevisionTags = dcl.FlattenStringSlice(m["stableRevisionTags"])

	return r
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentMap expands the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentMap(c *Client, f map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentSlice expands the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentSlice(c *Client, f []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment, res *DeliveryPipeline) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentMap flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentMap(c *Client, i interface{}, res *DeliveryPipeline) map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment{}
	}

	if len(a) == 0 {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment{}
	}

	items := make(map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment)
	for k, item := range a {
		items[k] = *flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentSlice flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentSlice(c *Client, i interface{}, res *DeliveryPipeline) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment {
	a, ok := i.([]interface{})
	if !ok {
		return []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment{}
	}

	if len(a) == 0 {
		return []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment{}
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment expands an instance of DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment(c *Client, f *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment, res *DeliveryPipeline) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Percentages; v != nil {
		m["percentages"] = v
	}
	if v := f.Verify; !dcl.IsEmptyValueIndirect(v) {
		m["verify"] = v
	}
	if v, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy(c, f.Predeploy, res); err != nil {
		return nil, fmt.Errorf("error expanding Predeploy into predeploy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["predeploy"] = v
	}
	if v, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy(c, f.Postdeploy, res); err != nil {
		return nil, fmt.Errorf("error expanding Postdeploy into postdeploy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["postdeploy"] = v
	}

	return m, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment flattens an instance of DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment(c *Client, i interface{}, res *DeliveryPipeline) *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment
	}
	r.Percentages = dcl.FlattenIntSlice(m["percentages"])
	r.Verify = dcl.FlattenBool(m["verify"])
	r.Predeploy = flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy(c, m["predeploy"], res)
	r.Postdeploy = flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy(c, m["postdeploy"], res)

	return r
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeployMap expands the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeployMap(c *Client, f map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploySlice expands the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploySlice(c *Client, f []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy, res *DeliveryPipeline) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeployMap flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeployMap(c *Client, i interface{}, res *DeliveryPipeline) map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy{}
	}

	if len(a) == 0 {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy{}
	}

	items := make(map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy)
	for k, item := range a {
		items[k] = *flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploySlice flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploySlice(c *Client, i interface{}, res *DeliveryPipeline) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy {
	a, ok := i.([]interface{})
	if !ok {
		return []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy{}
	}

	if len(a) == 0 {
		return []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy{}
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy expands an instance of DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy(c *Client, f *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy, res *DeliveryPipeline) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Actions; v != nil {
		m["actions"] = v
	}

	return m, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy flattens an instance of DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy(c *Client, i interface{}, res *DeliveryPipeline) *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy
	}
	r.Actions = dcl.FlattenStringSlice(m["actions"])

	return r
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeployMap expands the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeployMap(c *Client, f map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploySlice expands the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploySlice(c *Client, f []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy, res *DeliveryPipeline) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeployMap flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeployMap(c *Client, i interface{}, res *DeliveryPipeline) map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy{}
	}

	if len(a) == 0 {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy{}
	}

	items := make(map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy)
	for k, item := range a {
		items[k] = *flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploySlice flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploySlice(c *Client, i interface{}, res *DeliveryPipeline) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy {
	a, ok := i.([]interface{})
	if !ok {
		return []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy{}
	}

	if len(a) == 0 {
		return []DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy{}
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy expands an instance of DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy(c *Client, f *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy, res *DeliveryPipeline) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Actions; v != nil {
		m["actions"] = v
	}

	return m, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy flattens an instance of DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy(c *Client, i interface{}, res *DeliveryPipeline) *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy
	}
	r.Actions = dcl.FlattenStringSlice(m["actions"])

	return r
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentMap expands the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentMap(c *Client, f map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentSlice expands the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentSlice(c *Client, f []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment, res *DeliveryPipeline) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentMap flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentMap(c *Client, i interface{}, res *DeliveryPipeline) map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment{}
	}

	if len(a) == 0 {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment{}
	}

	items := make(map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment)
	for k, item := range a {
		items[k] = *flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentSlice flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentSlice(c *Client, i interface{}, res *DeliveryPipeline) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment {
	a, ok := i.([]interface{})
	if !ok {
		return []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment{}
	}

	if len(a) == 0 {
		return []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment{}
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment expands an instance of DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment(c *Client, f *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment, res *DeliveryPipeline) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsSlice(c, f.PhaseConfigs, res); err != nil {
		return nil, fmt.Errorf("error expanding PhaseConfigs into phaseConfigs: %w", err)
	} else if v != nil {
		m["phaseConfigs"] = v
	}

	return m, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment flattens an instance of DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment(c *Client, i interface{}, res *DeliveryPipeline) *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment
	}
	r.PhaseConfigs = flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsSlice(c, m["phaseConfigs"], res)

	return r
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsMap expands the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsMap(c *Client, f map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsSlice expands the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsSlice(c *Client, f []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs, res *DeliveryPipeline) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsMap flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsMap(c *Client, i interface{}, res *DeliveryPipeline) map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs{}
	}

	if len(a) == 0 {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs{}
	}

	items := make(map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs)
	for k, item := range a {
		items[k] = *flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsSlice flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsSlice(c *Client, i interface{}, res *DeliveryPipeline) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs {
	a, ok := i.([]interface{})
	if !ok {
		return []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs{}
	}

	if len(a) == 0 {
		return []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs{}
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs expands an instance of DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs(c *Client, f *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.PhaseId; !dcl.IsEmptyValueIndirect(v) {
		m["phaseId"] = v
	}
	if v := f.Percentage; !dcl.IsEmptyValueIndirect(v) {
		m["percentage"] = v
	}
	if v := f.Profiles; v != nil {
		m["profiles"] = v
	}
	if v := f.Verify; !dcl.IsEmptyValueIndirect(v) {
		m["verify"] = v
	}
	if v, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy(c, f.Predeploy, res); err != nil {
		return nil, fmt.Errorf("error expanding Predeploy into predeploy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["predeploy"] = v
	}
	if v, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy(c, f.Postdeploy, res); err != nil {
		return nil, fmt.Errorf("error expanding Postdeploy into postdeploy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["postdeploy"] = v
	}

	return m, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs flattens an instance of DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs(c *Client, i interface{}, res *DeliveryPipeline) *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs
	}
	r.PhaseId = dcl.FlattenString(m["phaseId"])
	r.Percentage = dcl.FlattenInteger(m["percentage"])
	r.Profiles = dcl.FlattenStringSlice(m["profiles"])
	r.Verify = dcl.FlattenBool(m["verify"])
	r.Predeploy = flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy(c, m["predeploy"], res)
	r.Postdeploy = flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy(c, m["postdeploy"], res)

	return r
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeployMap expands the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeployMap(c *Client, f map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploySlice expands the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploySlice(c *Client, f []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy, res *DeliveryPipeline) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeployMap flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeployMap(c *Client, i interface{}, res *DeliveryPipeline) map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy{}
	}

	if len(a) == 0 {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy{}
	}

	items := make(map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy)
	for k, item := range a {
		items[k] = *flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploySlice flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploySlice(c *Client, i interface{}, res *DeliveryPipeline) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy {
	a, ok := i.([]interface{})
	if !ok {
		return []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy{}
	}

	if len(a) == 0 {
		return []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy{}
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy expands an instance of DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy(c *Client, f *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy, res *DeliveryPipeline) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Actions; v != nil {
		m["actions"] = v
	}

	return m, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy flattens an instance of DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy(c *Client, i interface{}, res *DeliveryPipeline) *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy
	}
	r.Actions = dcl.FlattenStringSlice(m["actions"])

	return r
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeployMap expands the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeployMap(c *Client, f map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploySlice expands the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploySlice(c *Client, f []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy, res *DeliveryPipeline) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeployMap flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeployMap(c *Client, i interface{}, res *DeliveryPipeline) map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy{}
	}

	if len(a) == 0 {
		return map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy{}
	}

	items := make(map[string]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy)
	for k, item := range a {
		items[k] = *flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploySlice flattens the contents of DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploySlice(c *Client, i interface{}, res *DeliveryPipeline) []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy {
	a, ok := i.([]interface{})
	if !ok {
		return []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy{}
	}

	if len(a) == 0 {
		return []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy{}
	}

	items := make([]DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy expands an instance of DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy(c *Client, f *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy, res *DeliveryPipeline) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Actions; v != nil {
		m["actions"] = v
	}

	return m, nil
}

// flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy flattens an instance of DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy(c *Client, i interface{}, res *DeliveryPipeline) *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy
	}
	r.Actions = dcl.FlattenStringSlice(m["actions"])

	return r
}

// expandDeliveryPipelineSerialPipelineStagesDeployParametersMap expands the contents of DeliveryPipelineSerialPipelineStagesDeployParameters into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesDeployParametersMap(c *Client, f map[string]DeliveryPipelineSerialPipelineStagesDeployParameters, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesDeployParameters(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDeliveryPipelineSerialPipelineStagesDeployParametersSlice expands the contents of DeliveryPipelineSerialPipelineStagesDeployParameters into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesDeployParametersSlice(c *Client, f []DeliveryPipelineSerialPipelineStagesDeployParameters, res *DeliveryPipeline) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDeliveryPipelineSerialPipelineStagesDeployParameters(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDeliveryPipelineSerialPipelineStagesDeployParametersMap flattens the contents of DeliveryPipelineSerialPipelineStagesDeployParameters from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesDeployParametersMap(c *Client, i interface{}, res *DeliveryPipeline) map[string]DeliveryPipelineSerialPipelineStagesDeployParameters {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DeliveryPipelineSerialPipelineStagesDeployParameters{}
	}

	if len(a) == 0 {
		return map[string]DeliveryPipelineSerialPipelineStagesDeployParameters{}
	}

	items := make(map[string]DeliveryPipelineSerialPipelineStagesDeployParameters)
	for k, item := range a {
		items[k] = *flattenDeliveryPipelineSerialPipelineStagesDeployParameters(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDeliveryPipelineSerialPipelineStagesDeployParametersSlice flattens the contents of DeliveryPipelineSerialPipelineStagesDeployParameters from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesDeployParametersSlice(c *Client, i interface{}, res *DeliveryPipeline) []DeliveryPipelineSerialPipelineStagesDeployParameters {
	a, ok := i.([]interface{})
	if !ok {
		return []DeliveryPipelineSerialPipelineStagesDeployParameters{}
	}

	if len(a) == 0 {
		return []DeliveryPipelineSerialPipelineStagesDeployParameters{}
	}

	items := make([]DeliveryPipelineSerialPipelineStagesDeployParameters, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDeliveryPipelineSerialPipelineStagesDeployParameters(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDeliveryPipelineSerialPipelineStagesDeployParameters expands an instance of DeliveryPipelineSerialPipelineStagesDeployParameters into a JSON
// request object.
func expandDeliveryPipelineSerialPipelineStagesDeployParameters(c *Client, f *DeliveryPipelineSerialPipelineStagesDeployParameters, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Values; !dcl.IsEmptyValueIndirect(v) {
		m["values"] = v
	}
	if v := f.MatchTargetLabels; !dcl.IsEmptyValueIndirect(v) {
		m["matchTargetLabels"] = v
	}

	return m, nil
}

// flattenDeliveryPipelineSerialPipelineStagesDeployParameters flattens an instance of DeliveryPipelineSerialPipelineStagesDeployParameters from a JSON
// response object.
func flattenDeliveryPipelineSerialPipelineStagesDeployParameters(c *Client, i interface{}, res *DeliveryPipeline) *DeliveryPipelineSerialPipelineStagesDeployParameters {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DeliveryPipelineSerialPipelineStagesDeployParameters{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDeliveryPipelineSerialPipelineStagesDeployParameters
	}
	r.Values = dcl.FlattenKeyValuePairs(m["values"])
	r.MatchTargetLabels = dcl.FlattenKeyValuePairs(m["matchTargetLabels"])

	return r
}

// expandDeliveryPipelineConditionMap expands the contents of DeliveryPipelineCondition into a JSON
// request object.
func expandDeliveryPipelineConditionMap(c *Client, f map[string]DeliveryPipelineCondition, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDeliveryPipelineCondition(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDeliveryPipelineConditionSlice expands the contents of DeliveryPipelineCondition into a JSON
// request object.
func expandDeliveryPipelineConditionSlice(c *Client, f []DeliveryPipelineCondition, res *DeliveryPipeline) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDeliveryPipelineCondition(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDeliveryPipelineConditionMap flattens the contents of DeliveryPipelineCondition from a JSON
// response object.
func flattenDeliveryPipelineConditionMap(c *Client, i interface{}, res *DeliveryPipeline) map[string]DeliveryPipelineCondition {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DeliveryPipelineCondition{}
	}

	if len(a) == 0 {
		return map[string]DeliveryPipelineCondition{}
	}

	items := make(map[string]DeliveryPipelineCondition)
	for k, item := range a {
		items[k] = *flattenDeliveryPipelineCondition(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDeliveryPipelineConditionSlice flattens the contents of DeliveryPipelineCondition from a JSON
// response object.
func flattenDeliveryPipelineConditionSlice(c *Client, i interface{}, res *DeliveryPipeline) []DeliveryPipelineCondition {
	a, ok := i.([]interface{})
	if !ok {
		return []DeliveryPipelineCondition{}
	}

	if len(a) == 0 {
		return []DeliveryPipelineCondition{}
	}

	items := make([]DeliveryPipelineCondition, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDeliveryPipelineCondition(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDeliveryPipelineCondition expands an instance of DeliveryPipelineCondition into a JSON
// request object.
func expandDeliveryPipelineCondition(c *Client, f *DeliveryPipelineCondition, res *DeliveryPipeline) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandDeliveryPipelineConditionPipelineReadyCondition(c, f.PipelineReadyCondition, res); err != nil {
		return nil, fmt.Errorf("error expanding PipelineReadyCondition into pipelineReadyCondition: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["pipelineReadyCondition"] = v
	}
	if v, err := expandDeliveryPipelineConditionTargetsPresentCondition(c, f.TargetsPresentCondition, res); err != nil {
		return nil, fmt.Errorf("error expanding TargetsPresentCondition into targetsPresentCondition: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["targetsPresentCondition"] = v
	}
	if v, err := expandDeliveryPipelineConditionTargetsTypeCondition(c, f.TargetsTypeCondition, res); err != nil {
		return nil, fmt.Errorf("error expanding TargetsTypeCondition into targetsTypeCondition: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["targetsTypeCondition"] = v
	}

	return m, nil
}

// flattenDeliveryPipelineCondition flattens an instance of DeliveryPipelineCondition from a JSON
// response object.
func flattenDeliveryPipelineCondition(c *Client, i interface{}, res *DeliveryPipeline) *DeliveryPipelineCondition {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DeliveryPipelineCondition{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDeliveryPipelineCondition
	}
	r.PipelineReadyCondition = flattenDeliveryPipelineConditionPipelineReadyCondition(c, m["pipelineReadyCondition"], res)
	r.TargetsPresentCondition = flattenDeliveryPipelineConditionTargetsPresentCondition(c, m["targetsPresentCondition"], res)
	r.TargetsTypeCondition = flattenDeliveryPipelineConditionTargetsTypeCondition(c, m["targetsTypeCondition"], res)

	return r
}

// expandDeliveryPipelineConditionPipelineReadyConditionMap expands the contents of DeliveryPipelineConditionPipelineReadyCondition into a JSON
// request object.
func expandDeliveryPipelineConditionPipelineReadyConditionMap(c *Client, f map[string]DeliveryPipelineConditionPipelineReadyCondition, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDeliveryPipelineConditionPipelineReadyCondition(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDeliveryPipelineConditionPipelineReadyConditionSlice expands the contents of DeliveryPipelineConditionPipelineReadyCondition into a JSON
// request object.
func expandDeliveryPipelineConditionPipelineReadyConditionSlice(c *Client, f []DeliveryPipelineConditionPipelineReadyCondition, res *DeliveryPipeline) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDeliveryPipelineConditionPipelineReadyCondition(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDeliveryPipelineConditionPipelineReadyConditionMap flattens the contents of DeliveryPipelineConditionPipelineReadyCondition from a JSON
// response object.
func flattenDeliveryPipelineConditionPipelineReadyConditionMap(c *Client, i interface{}, res *DeliveryPipeline) map[string]DeliveryPipelineConditionPipelineReadyCondition {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DeliveryPipelineConditionPipelineReadyCondition{}
	}

	if len(a) == 0 {
		return map[string]DeliveryPipelineConditionPipelineReadyCondition{}
	}

	items := make(map[string]DeliveryPipelineConditionPipelineReadyCondition)
	for k, item := range a {
		items[k] = *flattenDeliveryPipelineConditionPipelineReadyCondition(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDeliveryPipelineConditionPipelineReadyConditionSlice flattens the contents of DeliveryPipelineConditionPipelineReadyCondition from a JSON
// response object.
func flattenDeliveryPipelineConditionPipelineReadyConditionSlice(c *Client, i interface{}, res *DeliveryPipeline) []DeliveryPipelineConditionPipelineReadyCondition {
	a, ok := i.([]interface{})
	if !ok {
		return []DeliveryPipelineConditionPipelineReadyCondition{}
	}

	if len(a) == 0 {
		return []DeliveryPipelineConditionPipelineReadyCondition{}
	}

	items := make([]DeliveryPipelineConditionPipelineReadyCondition, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDeliveryPipelineConditionPipelineReadyCondition(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDeliveryPipelineConditionPipelineReadyCondition expands an instance of DeliveryPipelineConditionPipelineReadyCondition into a JSON
// request object.
func expandDeliveryPipelineConditionPipelineReadyCondition(c *Client, f *DeliveryPipelineConditionPipelineReadyCondition, res *DeliveryPipeline) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Status; !dcl.IsEmptyValueIndirect(v) {
		m["status"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}

	return m, nil
}

// flattenDeliveryPipelineConditionPipelineReadyCondition flattens an instance of DeliveryPipelineConditionPipelineReadyCondition from a JSON
// response object.
func flattenDeliveryPipelineConditionPipelineReadyCondition(c *Client, i interface{}, res *DeliveryPipeline) *DeliveryPipelineConditionPipelineReadyCondition {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DeliveryPipelineConditionPipelineReadyCondition{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDeliveryPipelineConditionPipelineReadyCondition
	}
	r.Status = dcl.FlattenBool(m["status"])
	r.UpdateTime = dcl.FlattenString(m["updateTime"])

	return r
}

// expandDeliveryPipelineConditionTargetsPresentConditionMap expands the contents of DeliveryPipelineConditionTargetsPresentCondition into a JSON
// request object.
func expandDeliveryPipelineConditionTargetsPresentConditionMap(c *Client, f map[string]DeliveryPipelineConditionTargetsPresentCondition, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDeliveryPipelineConditionTargetsPresentCondition(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDeliveryPipelineConditionTargetsPresentConditionSlice expands the contents of DeliveryPipelineConditionTargetsPresentCondition into a JSON
// request object.
func expandDeliveryPipelineConditionTargetsPresentConditionSlice(c *Client, f []DeliveryPipelineConditionTargetsPresentCondition, res *DeliveryPipeline) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDeliveryPipelineConditionTargetsPresentCondition(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDeliveryPipelineConditionTargetsPresentConditionMap flattens the contents of DeliveryPipelineConditionTargetsPresentCondition from a JSON
// response object.
func flattenDeliveryPipelineConditionTargetsPresentConditionMap(c *Client, i interface{}, res *DeliveryPipeline) map[string]DeliveryPipelineConditionTargetsPresentCondition {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DeliveryPipelineConditionTargetsPresentCondition{}
	}

	if len(a) == 0 {
		return map[string]DeliveryPipelineConditionTargetsPresentCondition{}
	}

	items := make(map[string]DeliveryPipelineConditionTargetsPresentCondition)
	for k, item := range a {
		items[k] = *flattenDeliveryPipelineConditionTargetsPresentCondition(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDeliveryPipelineConditionTargetsPresentConditionSlice flattens the contents of DeliveryPipelineConditionTargetsPresentCondition from a JSON
// response object.
func flattenDeliveryPipelineConditionTargetsPresentConditionSlice(c *Client, i interface{}, res *DeliveryPipeline) []DeliveryPipelineConditionTargetsPresentCondition {
	a, ok := i.([]interface{})
	if !ok {
		return []DeliveryPipelineConditionTargetsPresentCondition{}
	}

	if len(a) == 0 {
		return []DeliveryPipelineConditionTargetsPresentCondition{}
	}

	items := make([]DeliveryPipelineConditionTargetsPresentCondition, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDeliveryPipelineConditionTargetsPresentCondition(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDeliveryPipelineConditionTargetsPresentCondition expands an instance of DeliveryPipelineConditionTargetsPresentCondition into a JSON
// request object.
func expandDeliveryPipelineConditionTargetsPresentCondition(c *Client, f *DeliveryPipelineConditionTargetsPresentCondition, res *DeliveryPipeline) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Status; !dcl.IsEmptyValueIndirect(v) {
		m["status"] = v
	}
	if v := f.MissingTargets; v != nil {
		m["missingTargets"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}

	return m, nil
}

// flattenDeliveryPipelineConditionTargetsPresentCondition flattens an instance of DeliveryPipelineConditionTargetsPresentCondition from a JSON
// response object.
func flattenDeliveryPipelineConditionTargetsPresentCondition(c *Client, i interface{}, res *DeliveryPipeline) *DeliveryPipelineConditionTargetsPresentCondition {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DeliveryPipelineConditionTargetsPresentCondition{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDeliveryPipelineConditionTargetsPresentCondition
	}
	r.Status = dcl.FlattenBool(m["status"])
	r.MissingTargets = dcl.FlattenStringSlice(m["missingTargets"])
	r.UpdateTime = dcl.FlattenString(m["updateTime"])

	return r
}

// expandDeliveryPipelineConditionTargetsTypeConditionMap expands the contents of DeliveryPipelineConditionTargetsTypeCondition into a JSON
// request object.
func expandDeliveryPipelineConditionTargetsTypeConditionMap(c *Client, f map[string]DeliveryPipelineConditionTargetsTypeCondition, res *DeliveryPipeline) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDeliveryPipelineConditionTargetsTypeCondition(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDeliveryPipelineConditionTargetsTypeConditionSlice expands the contents of DeliveryPipelineConditionTargetsTypeCondition into a JSON
// request object.
func expandDeliveryPipelineConditionTargetsTypeConditionSlice(c *Client, f []DeliveryPipelineConditionTargetsTypeCondition, res *DeliveryPipeline) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDeliveryPipelineConditionTargetsTypeCondition(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDeliveryPipelineConditionTargetsTypeConditionMap flattens the contents of DeliveryPipelineConditionTargetsTypeCondition from a JSON
// response object.
func flattenDeliveryPipelineConditionTargetsTypeConditionMap(c *Client, i interface{}, res *DeliveryPipeline) map[string]DeliveryPipelineConditionTargetsTypeCondition {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DeliveryPipelineConditionTargetsTypeCondition{}
	}

	if len(a) == 0 {
		return map[string]DeliveryPipelineConditionTargetsTypeCondition{}
	}

	items := make(map[string]DeliveryPipelineConditionTargetsTypeCondition)
	for k, item := range a {
		items[k] = *flattenDeliveryPipelineConditionTargetsTypeCondition(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDeliveryPipelineConditionTargetsTypeConditionSlice flattens the contents of DeliveryPipelineConditionTargetsTypeCondition from a JSON
// response object.
func flattenDeliveryPipelineConditionTargetsTypeConditionSlice(c *Client, i interface{}, res *DeliveryPipeline) []DeliveryPipelineConditionTargetsTypeCondition {
	a, ok := i.([]interface{})
	if !ok {
		return []DeliveryPipelineConditionTargetsTypeCondition{}
	}

	if len(a) == 0 {
		return []DeliveryPipelineConditionTargetsTypeCondition{}
	}

	items := make([]DeliveryPipelineConditionTargetsTypeCondition, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDeliveryPipelineConditionTargetsTypeCondition(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDeliveryPipelineConditionTargetsTypeCondition expands an instance of DeliveryPipelineConditionTargetsTypeCondition into a JSON
// request object.
func expandDeliveryPipelineConditionTargetsTypeCondition(c *Client, f *DeliveryPipelineConditionTargetsTypeCondition, res *DeliveryPipeline) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Status; !dcl.IsEmptyValueIndirect(v) {
		m["status"] = v
	}
	if v := f.ErrorDetails; !dcl.IsEmptyValueIndirect(v) {
		m["errorDetails"] = v
	}

	return m, nil
}

// flattenDeliveryPipelineConditionTargetsTypeCondition flattens an instance of DeliveryPipelineConditionTargetsTypeCondition from a JSON
// response object.
func flattenDeliveryPipelineConditionTargetsTypeCondition(c *Client, i interface{}, res *DeliveryPipeline) *DeliveryPipelineConditionTargetsTypeCondition {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DeliveryPipelineConditionTargetsTypeCondition{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDeliveryPipelineConditionTargetsTypeCondition
	}
	r.Status = dcl.FlattenBool(m["status"])
	r.ErrorDetails = dcl.FlattenString(m["errorDetails"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *DeliveryPipeline) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalDeliveryPipeline(b, c, r)
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

type deliveryPipelineDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         deliveryPipelineApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToDeliveryPipelineDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]deliveryPipelineDiff, error) {
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
	var diffs []deliveryPipelineDiff
	// For each operation name, create a deliveryPipelineDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := deliveryPipelineDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToDeliveryPipelineApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToDeliveryPipelineApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (deliveryPipelineApiOperation, error) {
	switch opName {

	case "updateDeliveryPipelineUpdateDeliveryPipelineOperation":
		return &updateDeliveryPipelineUpdateDeliveryPipelineOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractDeliveryPipelineFields(r *DeliveryPipeline) error {
	vSerialPipeline := r.SerialPipeline
	if vSerialPipeline == nil {
		// note: explicitly not the empty object.
		vSerialPipeline = &DeliveryPipelineSerialPipeline{}
	}
	if err := extractDeliveryPipelineSerialPipelineFields(r, vSerialPipeline); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSerialPipeline) {
		r.SerialPipeline = vSerialPipeline
	}
	vCondition := r.Condition
	if vCondition == nil {
		// note: explicitly not the empty object.
		vCondition = &DeliveryPipelineCondition{}
	}
	if err := extractDeliveryPipelineConditionFields(r, vCondition); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCondition) {
		r.Condition = vCondition
	}
	return nil
}
func extractDeliveryPipelineSerialPipelineFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipeline) error {
	return nil
}
func extractDeliveryPipelineSerialPipelineStagesFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStages) error {
	vStrategy := o.Strategy
	if vStrategy == nil {
		// note: explicitly not the empty object.
		vStrategy = &DeliveryPipelineSerialPipelineStagesStrategy{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyFields(r, vStrategy); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vStrategy) {
		o.Strategy = vStrategy
	}
	return nil
}
func extractDeliveryPipelineSerialPipelineStagesStrategyFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategy) error {
	vStandard := o.Standard
	if vStandard == nil {
		// note: explicitly not the empty object.
		vStandard = &DeliveryPipelineSerialPipelineStagesStrategyStandard{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyStandardFields(r, vStandard); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vStandard) {
		o.Standard = vStandard
	}
	vCanary := o.Canary
	if vCanary == nil {
		// note: explicitly not the empty object.
		vCanary = &DeliveryPipelineSerialPipelineStagesStrategyCanary{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyCanaryFields(r, vCanary); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCanary) {
		o.Canary = vCanary
	}
	return nil
}
func extractDeliveryPipelineSerialPipelineStagesStrategyStandardFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyStandard) error {
	vPredeploy := o.Predeploy
	if vPredeploy == nil {
		// note: explicitly not the empty object.
		vPredeploy = &DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyStandardPredeployFields(r, vPredeploy); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPredeploy) {
		o.Predeploy = vPredeploy
	}
	vPostdeploy := o.Postdeploy
	if vPostdeploy == nil {
		// note: explicitly not the empty object.
		vPostdeploy = &DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeployFields(r, vPostdeploy); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPostdeploy) {
		o.Postdeploy = vPostdeploy
	}
	return nil
}
func extractDeliveryPipelineSerialPipelineStagesStrategyStandardPredeployFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy) error {
	return nil
}
func extractDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeployFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy) error {
	return nil
}
func extractDeliveryPipelineSerialPipelineStagesStrategyCanaryFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyCanary) error {
	vRuntimeConfig := o.RuntimeConfig
	if vRuntimeConfig == nil {
		// note: explicitly not the empty object.
		vRuntimeConfig = &DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigFields(r, vRuntimeConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRuntimeConfig) {
		o.RuntimeConfig = vRuntimeConfig
	}
	vCanaryDeployment := o.CanaryDeployment
	if vCanaryDeployment == nil {
		// note: explicitly not the empty object.
		vCanaryDeployment = &DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentFields(r, vCanaryDeployment); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCanaryDeployment) {
		o.CanaryDeployment = vCanaryDeployment
	}
	vCustomCanaryDeployment := o.CustomCanaryDeployment
	if vCustomCanaryDeployment == nil {
		// note: explicitly not the empty object.
		vCustomCanaryDeployment = &DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentFields(r, vCustomCanaryDeployment); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCustomCanaryDeployment) {
		o.CustomCanaryDeployment = vCustomCanaryDeployment
	}
	return nil
}
func extractDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig) error {
	vKubernetes := o.Kubernetes
	if vKubernetes == nil {
		// note: explicitly not the empty object.
		vKubernetes = &DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesFields(r, vKubernetes); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vKubernetes) {
		o.Kubernetes = vKubernetes
	}
	vCloudRun := o.CloudRun
	if vCloudRun == nil {
		// note: explicitly not the empty object.
		vCloudRun = &DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunFields(r, vCloudRun); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCloudRun) {
		o.CloudRun = vCloudRun
	}
	return nil
}
func extractDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes) error {
	vGatewayServiceMesh := o.GatewayServiceMesh
	if vGatewayServiceMesh == nil {
		// note: explicitly not the empty object.
		vGatewayServiceMesh = &DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshFields(r, vGatewayServiceMesh); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vGatewayServiceMesh) {
		o.GatewayServiceMesh = vGatewayServiceMesh
	}
	vServiceNetworking := o.ServiceNetworking
	if vServiceNetworking == nil {
		// note: explicitly not the empty object.
		vServiceNetworking = &DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingFields(r, vServiceNetworking); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vServiceNetworking) {
		o.ServiceNetworking = vServiceNetworking
	}
	return nil
}
func extractDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh) error {
	return nil
}
func extractDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking) error {
	return nil
}
func extractDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun) error {
	return nil
}
func extractDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment) error {
	vPredeploy := o.Predeploy
	if vPredeploy == nil {
		// note: explicitly not the empty object.
		vPredeploy = &DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeployFields(r, vPredeploy); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPredeploy) {
		o.Predeploy = vPredeploy
	}
	vPostdeploy := o.Postdeploy
	if vPostdeploy == nil {
		// note: explicitly not the empty object.
		vPostdeploy = &DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeployFields(r, vPostdeploy); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPostdeploy) {
		o.Postdeploy = vPostdeploy
	}
	return nil
}
func extractDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeployFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy) error {
	return nil
}
func extractDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeployFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy) error {
	return nil
}
func extractDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment) error {
	return nil
}
func extractDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs) error {
	vPredeploy := o.Predeploy
	if vPredeploy == nil {
		// note: explicitly not the empty object.
		vPredeploy = &DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeployFields(r, vPredeploy); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPredeploy) {
		o.Predeploy = vPredeploy
	}
	vPostdeploy := o.Postdeploy
	if vPostdeploy == nil {
		// note: explicitly not the empty object.
		vPostdeploy = &DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeployFields(r, vPostdeploy); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPostdeploy) {
		o.Postdeploy = vPostdeploy
	}
	return nil
}
func extractDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeployFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy) error {
	return nil
}
func extractDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeployFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy) error {
	return nil
}
func extractDeliveryPipelineSerialPipelineStagesDeployParametersFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesDeployParameters) error {
	return nil
}
func extractDeliveryPipelineConditionFields(r *DeliveryPipeline, o *DeliveryPipelineCondition) error {
	vPipelineReadyCondition := o.PipelineReadyCondition
	if vPipelineReadyCondition == nil {
		// note: explicitly not the empty object.
		vPipelineReadyCondition = &DeliveryPipelineConditionPipelineReadyCondition{}
	}
	if err := extractDeliveryPipelineConditionPipelineReadyConditionFields(r, vPipelineReadyCondition); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPipelineReadyCondition) {
		o.PipelineReadyCondition = vPipelineReadyCondition
	}
	vTargetsPresentCondition := o.TargetsPresentCondition
	if vTargetsPresentCondition == nil {
		// note: explicitly not the empty object.
		vTargetsPresentCondition = &DeliveryPipelineConditionTargetsPresentCondition{}
	}
	if err := extractDeliveryPipelineConditionTargetsPresentConditionFields(r, vTargetsPresentCondition); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTargetsPresentCondition) {
		o.TargetsPresentCondition = vTargetsPresentCondition
	}
	vTargetsTypeCondition := o.TargetsTypeCondition
	if vTargetsTypeCondition == nil {
		// note: explicitly not the empty object.
		vTargetsTypeCondition = &DeliveryPipelineConditionTargetsTypeCondition{}
	}
	if err := extractDeliveryPipelineConditionTargetsTypeConditionFields(r, vTargetsTypeCondition); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTargetsTypeCondition) {
		o.TargetsTypeCondition = vTargetsTypeCondition
	}
	return nil
}
func extractDeliveryPipelineConditionPipelineReadyConditionFields(r *DeliveryPipeline, o *DeliveryPipelineConditionPipelineReadyCondition) error {
	return nil
}
func extractDeliveryPipelineConditionTargetsPresentConditionFields(r *DeliveryPipeline, o *DeliveryPipelineConditionTargetsPresentCondition) error {
	return nil
}
func extractDeliveryPipelineConditionTargetsTypeConditionFields(r *DeliveryPipeline, o *DeliveryPipelineConditionTargetsTypeCondition) error {
	return nil
}

func postReadExtractDeliveryPipelineFields(r *DeliveryPipeline) error {
	vSerialPipeline := r.SerialPipeline
	if vSerialPipeline == nil {
		// note: explicitly not the empty object.
		vSerialPipeline = &DeliveryPipelineSerialPipeline{}
	}
	if err := postReadExtractDeliveryPipelineSerialPipelineFields(r, vSerialPipeline); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSerialPipeline) {
		r.SerialPipeline = vSerialPipeline
	}
	vCondition := r.Condition
	if vCondition == nil {
		// note: explicitly not the empty object.
		vCondition = &DeliveryPipelineCondition{}
	}
	if err := postReadExtractDeliveryPipelineConditionFields(r, vCondition); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCondition) {
		r.Condition = vCondition
	}
	return nil
}
func postReadExtractDeliveryPipelineSerialPipelineFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipeline) error {
	return nil
}
func postReadExtractDeliveryPipelineSerialPipelineStagesFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStages) error {
	vStrategy := o.Strategy
	if vStrategy == nil {
		// note: explicitly not the empty object.
		vStrategy = &DeliveryPipelineSerialPipelineStagesStrategy{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyFields(r, vStrategy); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vStrategy) {
		o.Strategy = vStrategy
	}
	return nil
}
func postReadExtractDeliveryPipelineSerialPipelineStagesStrategyFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategy) error {
	vStandard := o.Standard
	if vStandard == nil {
		// note: explicitly not the empty object.
		vStandard = &DeliveryPipelineSerialPipelineStagesStrategyStandard{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyStandardFields(r, vStandard); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vStandard) {
		o.Standard = vStandard
	}
	vCanary := o.Canary
	if vCanary == nil {
		// note: explicitly not the empty object.
		vCanary = &DeliveryPipelineSerialPipelineStagesStrategyCanary{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyCanaryFields(r, vCanary); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCanary) {
		o.Canary = vCanary
	}
	return nil
}
func postReadExtractDeliveryPipelineSerialPipelineStagesStrategyStandardFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyStandard) error {
	vPredeploy := o.Predeploy
	if vPredeploy == nil {
		// note: explicitly not the empty object.
		vPredeploy = &DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyStandardPredeployFields(r, vPredeploy); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPredeploy) {
		o.Predeploy = vPredeploy
	}
	vPostdeploy := o.Postdeploy
	if vPostdeploy == nil {
		// note: explicitly not the empty object.
		vPostdeploy = &DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeployFields(r, vPostdeploy); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPostdeploy) {
		o.Postdeploy = vPostdeploy
	}
	return nil
}
func postReadExtractDeliveryPipelineSerialPipelineStagesStrategyStandardPredeployFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy) error {
	return nil
}
func postReadExtractDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeployFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy) error {
	return nil
}
func postReadExtractDeliveryPipelineSerialPipelineStagesStrategyCanaryFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyCanary) error {
	vRuntimeConfig := o.RuntimeConfig
	if vRuntimeConfig == nil {
		// note: explicitly not the empty object.
		vRuntimeConfig = &DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigFields(r, vRuntimeConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRuntimeConfig) {
		o.RuntimeConfig = vRuntimeConfig
	}
	vCanaryDeployment := o.CanaryDeployment
	if vCanaryDeployment == nil {
		// note: explicitly not the empty object.
		vCanaryDeployment = &DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentFields(r, vCanaryDeployment); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCanaryDeployment) {
		o.CanaryDeployment = vCanaryDeployment
	}
	vCustomCanaryDeployment := o.CustomCanaryDeployment
	if vCustomCanaryDeployment == nil {
		// note: explicitly not the empty object.
		vCustomCanaryDeployment = &DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentFields(r, vCustomCanaryDeployment); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCustomCanaryDeployment) {
		o.CustomCanaryDeployment = vCustomCanaryDeployment
	}
	return nil
}
func postReadExtractDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig) error {
	vKubernetes := o.Kubernetes
	if vKubernetes == nil {
		// note: explicitly not the empty object.
		vKubernetes = &DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesFields(r, vKubernetes); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vKubernetes) {
		o.Kubernetes = vKubernetes
	}
	vCloudRun := o.CloudRun
	if vCloudRun == nil {
		// note: explicitly not the empty object.
		vCloudRun = &DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunFields(r, vCloudRun); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCloudRun) {
		o.CloudRun = vCloudRun
	}
	return nil
}
func postReadExtractDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes) error {
	vGatewayServiceMesh := o.GatewayServiceMesh
	if vGatewayServiceMesh == nil {
		// note: explicitly not the empty object.
		vGatewayServiceMesh = &DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshFields(r, vGatewayServiceMesh); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vGatewayServiceMesh) {
		o.GatewayServiceMesh = vGatewayServiceMesh
	}
	vServiceNetworking := o.ServiceNetworking
	if vServiceNetworking == nil {
		// note: explicitly not the empty object.
		vServiceNetworking = &DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingFields(r, vServiceNetworking); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vServiceNetworking) {
		o.ServiceNetworking = vServiceNetworking
	}
	return nil
}
func postReadExtractDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh) error {
	return nil
}
func postReadExtractDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking) error {
	return nil
}
func postReadExtractDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun) error {
	return nil
}
func postReadExtractDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment) error {
	vPredeploy := o.Predeploy
	if vPredeploy == nil {
		// note: explicitly not the empty object.
		vPredeploy = &DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeployFields(r, vPredeploy); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPredeploy) {
		o.Predeploy = vPredeploy
	}
	vPostdeploy := o.Postdeploy
	if vPostdeploy == nil {
		// note: explicitly not the empty object.
		vPostdeploy = &DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeployFields(r, vPostdeploy); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPostdeploy) {
		o.Postdeploy = vPostdeploy
	}
	return nil
}
func postReadExtractDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeployFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy) error {
	return nil
}
func postReadExtractDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeployFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy) error {
	return nil
}
func postReadExtractDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment) error {
	return nil
}
func postReadExtractDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs) error {
	vPredeploy := o.Predeploy
	if vPredeploy == nil {
		// note: explicitly not the empty object.
		vPredeploy = &DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeployFields(r, vPredeploy); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPredeploy) {
		o.Predeploy = vPredeploy
	}
	vPostdeploy := o.Postdeploy
	if vPostdeploy == nil {
		// note: explicitly not the empty object.
		vPostdeploy = &DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy{}
	}
	if err := extractDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeployFields(r, vPostdeploy); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPostdeploy) {
		o.Postdeploy = vPostdeploy
	}
	return nil
}
func postReadExtractDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeployFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy) error {
	return nil
}
func postReadExtractDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeployFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy) error {
	return nil
}
func postReadExtractDeliveryPipelineSerialPipelineStagesDeployParametersFields(r *DeliveryPipeline, o *DeliveryPipelineSerialPipelineStagesDeployParameters) error {
	return nil
}
func postReadExtractDeliveryPipelineConditionFields(r *DeliveryPipeline, o *DeliveryPipelineCondition) error {
	vPipelineReadyCondition := o.PipelineReadyCondition
	if vPipelineReadyCondition == nil {
		// note: explicitly not the empty object.
		vPipelineReadyCondition = &DeliveryPipelineConditionPipelineReadyCondition{}
	}
	if err := extractDeliveryPipelineConditionPipelineReadyConditionFields(r, vPipelineReadyCondition); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPipelineReadyCondition) {
		o.PipelineReadyCondition = vPipelineReadyCondition
	}
	vTargetsPresentCondition := o.TargetsPresentCondition
	if vTargetsPresentCondition == nil {
		// note: explicitly not the empty object.
		vTargetsPresentCondition = &DeliveryPipelineConditionTargetsPresentCondition{}
	}
	if err := extractDeliveryPipelineConditionTargetsPresentConditionFields(r, vTargetsPresentCondition); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTargetsPresentCondition) {
		o.TargetsPresentCondition = vTargetsPresentCondition
	}
	vTargetsTypeCondition := o.TargetsTypeCondition
	if vTargetsTypeCondition == nil {
		// note: explicitly not the empty object.
		vTargetsTypeCondition = &DeliveryPipelineConditionTargetsTypeCondition{}
	}
	if err := extractDeliveryPipelineConditionTargetsTypeConditionFields(r, vTargetsTypeCondition); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTargetsTypeCondition) {
		o.TargetsTypeCondition = vTargetsTypeCondition
	}
	return nil
}
func postReadExtractDeliveryPipelineConditionPipelineReadyConditionFields(r *DeliveryPipeline, o *DeliveryPipelineConditionPipelineReadyCondition) error {
	return nil
}
func postReadExtractDeliveryPipelineConditionTargetsPresentConditionFields(r *DeliveryPipeline, o *DeliveryPipelineConditionTargetsPresentCondition) error {
	return nil
}
func postReadExtractDeliveryPipelineConditionTargetsTypeConditionFields(r *DeliveryPipeline, o *DeliveryPipelineConditionTargetsTypeCondition) error {
	return nil
}
