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

func (r *Service) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "template"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.BinaryAuthorization) {
		if err := r.BinaryAuthorization.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Template) {
		if err := r.Template.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.TerminalCondition) {
		if err := r.TerminalCondition.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceBinaryAuthorization) validate() error {
	return nil
}
func (r *ServiceTemplate) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Scaling) {
		if err := r.Scaling.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.VPCAccess) {
		if err := r.VPCAccess.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceTemplateScaling) validate() error {
	return nil
}
func (r *ServiceTemplateVPCAccess) validate() error {
	return nil
}
func (r *ServiceTemplateContainers) validate() error {
	if err := dcl.Required(r, "image"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Resources) {
		if err := r.Resources.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceTemplateContainersEnv) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"Value", "ValueSource"}, r.Value, r.ValueSource); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.ValueSource) {
		if err := r.ValueSource.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceTemplateContainersEnvValueSource) validate() error {
	if !dcl.IsEmptyValueIndirect(r.SecretKeyRef) {
		if err := r.SecretKeyRef.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceTemplateContainersEnvValueSourceSecretKeyRef) validate() error {
	if err := dcl.Required(r, "secret"); err != nil {
		return err
	}
	return nil
}
func (r *ServiceTemplateContainersResources) validate() error {
	return nil
}
func (r *ServiceTemplateContainersPorts) validate() error {
	return nil
}
func (r *ServiceTemplateContainersVolumeMounts) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "mountPath"); err != nil {
		return err
	}
	return nil
}
func (r *ServiceTemplateVolumes) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"Secret", "CloudSqlInstance"}, r.Secret, r.CloudSqlInstance); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Secret) {
		if err := r.Secret.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CloudSqlInstance) {
		if err := r.CloudSqlInstance.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceTemplateVolumesSecret) validate() error {
	if err := dcl.Required(r, "secret"); err != nil {
		return err
	}
	return nil
}
func (r *ServiceTemplateVolumesSecretItems) validate() error {
	if err := dcl.Required(r, "path"); err != nil {
		return err
	}
	return nil
}
func (r *ServiceTemplateVolumesCloudSqlInstance) validate() error {
	return nil
}
func (r *ServiceTraffic) validate() error {
	return nil
}
func (r *ServiceTerminalCondition) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"Reason", "RevisionReason", "JobReason"}, r.Reason, r.RevisionReason, r.JobReason); err != nil {
		return err
	}
	return nil
}
func (r *ServiceTrafficStatuses) validate() error {
	return nil
}
func (r *Service) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://run.googleapis.com/v2/", params)
}

func (r *Service) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/services/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Service) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/services", nr.basePath(), userBasePath, params), nil

}

func (r *Service) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/services?serviceId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *Service) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/services/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Service) SetPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"project":  *nr.Project,
		"location": *nr.Location,
		"name":     *nr.Name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/services/{{name}}:setIamPolicy", nr.basePath(), userBasePath, fields)
}

func (r *Service) SetPolicyVerb() string {
	return "POST"
}

func (r *Service) getPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"project":  *nr.Project,
		"location": *nr.Location,
		"name":     *nr.Name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/services/{{name}}:getIamPolicy", nr.basePath(), userBasePath, fields)
}

func (r *Service) IAMPolicyVersion() int {
	return 3
}

// serviceApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type serviceApiOperation interface {
	do(context.Context, *Service, *Client) error
}

// newUpdateServiceUpdateServiceRequest creates a request for an
// Service resource's UpdateService update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateServiceUpdateServiceRequest(ctx context.Context, f *Service, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v, err := dcl.DeriveField("projects/%s/locations/%s/services/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Location), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Annotations; !dcl.IsEmptyValueIndirect(v) {
		req["annotations"] = v
	}
	if v := f.Client; !dcl.IsEmptyValueIndirect(v) {
		req["client"] = v
	}
	if v := f.ClientVersion; !dcl.IsEmptyValueIndirect(v) {
		req["clientVersion"] = v
	}
	if v := f.Ingress; !dcl.IsEmptyValueIndirect(v) {
		req["ingress"] = v
	}
	if v := f.LaunchStage; !dcl.IsEmptyValueIndirect(v) {
		req["launchStage"] = v
	}
	if v, err := expandServiceBinaryAuthorization(c, f.BinaryAuthorization, res); err != nil {
		return nil, fmt.Errorf("error expanding BinaryAuthorization into binaryAuthorization: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["binaryAuthorization"] = v
	}
	if v, err := expandServiceTemplate(c, f.Template, res); err != nil {
		return nil, fmt.Errorf("error expanding Template into template: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["template"] = v
	}
	if v, err := expandServiceTrafficSlice(c, f.Traffic, res); err != nil {
		return nil, fmt.Errorf("error expanding Traffic into traffic: %w", err)
	} else if v != nil {
		req["traffic"] = v
	}
	if v, err := expandServiceTerminalCondition(c, f.TerminalCondition, res); err != nil {
		return nil, fmt.Errorf("error expanding TerminalCondition into terminalCondition: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["terminalCondition"] = v
	}
	b, err := c.getServiceRaw(ctx, f)
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

// marshalUpdateServiceUpdateServiceRequest converts the update into
// the final JSON request body.
func marshalUpdateServiceUpdateServiceRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateServiceUpdateServiceOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateServiceUpdateServiceOperation) do(ctx context.Context, r *Service, c *Client) error {
	_, err := c.GetService(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateService")
	if err != nil {
		return err
	}

	req, err := newUpdateServiceUpdateServiceRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateServiceUpdateServiceRequest(c, req)
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

func (c *Client) listServiceRaw(ctx context.Context, r *Service, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ServiceMaxPage {
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

type listServiceOperation struct {
	Services []map[string]interface{} `json:"services"`
	Token    string                   `json:"nextPageToken"`
}

func (c *Client) listService(ctx context.Context, r *Service, pageToken string, pageSize int32) ([]*Service, string, error) {
	b, err := c.listServiceRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listServiceOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Service
	for _, v := range m.Services {
		res, err := unmarshalMapService(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllService(ctx context.Context, f func(*Service) bool, resources []*Service) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteService(ctx, res)
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

type deleteServiceOperation struct{}

func (op *deleteServiceOperation) do(ctx context.Context, r *Service, c *Client) error {
	r, err := c.GetService(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Service not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetService checking for existence. error: %v", err)
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
		return fmt.Errorf("failed to delete Service: %w", err)
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createServiceOperation struct {
	response map[string]interface{}
}

func (op *createServiceOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createServiceOperation) do(ctx context.Context, r *Service, c *Client) error {
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

	if _, err := c.GetService(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getServiceRaw(ctx context.Context, r *Service) ([]byte, error) {

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

func (c *Client) serviceDiffsForRawDesired(ctx context.Context, rawDesired *Service, opts ...dcl.ApplyOption) (initial, desired *Service, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Service
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Service); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Service, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetService(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Service resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Service resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Service resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeServiceDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Service: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Service: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractServiceFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeServiceInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Service: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeServiceDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Service: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffService(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeServiceInitialState(rawInitial, rawDesired *Service) (*Service, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeServiceDesiredState(rawDesired, rawInitial *Service, opts ...dcl.ApplyOption) (*Service, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.BinaryAuthorization = canonicalizeServiceBinaryAuthorization(rawDesired.BinaryAuthorization, nil, opts...)
		rawDesired.Template = canonicalizeServiceTemplate(rawDesired.Template, nil, opts...)
		rawDesired.TerminalCondition = canonicalizeServiceTerminalCondition(rawDesired.TerminalCondition, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Service{}
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
	if dcl.IsZeroValue(rawDesired.Annotations) || (dcl.IsEmptyValueIndirect(rawDesired.Annotations) && dcl.IsEmptyValueIndirect(rawInitial.Annotations)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Annotations = rawInitial.Annotations
	} else {
		canonicalDesired.Annotations = rawDesired.Annotations
	}
	if dcl.StringCanonicalize(rawDesired.Client, rawInitial.Client) {
		canonicalDesired.Client = rawInitial.Client
	} else {
		canonicalDesired.Client = rawDesired.Client
	}
	if dcl.StringCanonicalize(rawDesired.ClientVersion, rawInitial.ClientVersion) {
		canonicalDesired.ClientVersion = rawInitial.ClientVersion
	} else {
		canonicalDesired.ClientVersion = rawDesired.ClientVersion
	}
	if dcl.IsZeroValue(rawDesired.Ingress) || (dcl.IsEmptyValueIndirect(rawDesired.Ingress) && dcl.IsEmptyValueIndirect(rawInitial.Ingress)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Ingress = rawInitial.Ingress
	} else {
		canonicalDesired.Ingress = rawDesired.Ingress
	}
	if dcl.IsZeroValue(rawDesired.LaunchStage) || (dcl.IsEmptyValueIndirect(rawDesired.LaunchStage) && dcl.IsEmptyValueIndirect(rawInitial.LaunchStage)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.LaunchStage = rawInitial.LaunchStage
	} else {
		canonicalDesired.LaunchStage = rawDesired.LaunchStage
	}
	canonicalDesired.BinaryAuthorization = canonicalizeServiceBinaryAuthorization(rawDesired.BinaryAuthorization, rawInitial.BinaryAuthorization, opts...)
	canonicalDesired.Template = canonicalizeServiceTemplate(rawDesired.Template, rawInitial.Template, opts...)
	canonicalDesired.Traffic = canonicalizeServiceTrafficSlice(rawDesired.Traffic, rawInitial.Traffic, opts...)
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

func canonicalizeServiceNewState(c *Client, rawNew, rawDesired *Service) (*Service, error) {

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

	if dcl.IsEmptyValueIndirect(rawNew.Uid) && dcl.IsEmptyValueIndirect(rawDesired.Uid) {
		rawNew.Uid = rawDesired.Uid
	} else {
		if dcl.StringCanonicalize(rawDesired.Uid, rawNew.Uid) {
			rawNew.Uid = rawDesired.Uid
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Generation) && dcl.IsEmptyValueIndirect(rawDesired.Generation) {
		rawNew.Generation = rawDesired.Generation
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Annotations) && dcl.IsEmptyValueIndirect(rawDesired.Annotations) {
		rawNew.Annotations = rawDesired.Annotations
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

	if dcl.IsEmptyValueIndirect(rawNew.DeleteTime) && dcl.IsEmptyValueIndirect(rawDesired.DeleteTime) {
		rawNew.DeleteTime = rawDesired.DeleteTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ExpireTime) && dcl.IsEmptyValueIndirect(rawDesired.ExpireTime) {
		rawNew.ExpireTime = rawDesired.ExpireTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Creator) && dcl.IsEmptyValueIndirect(rawDesired.Creator) {
		rawNew.Creator = rawDesired.Creator
	} else {
		if dcl.StringCanonicalize(rawDesired.Creator, rawNew.Creator) {
			rawNew.Creator = rawDesired.Creator
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.LastModifier) && dcl.IsEmptyValueIndirect(rawDesired.LastModifier) {
		rawNew.LastModifier = rawDesired.LastModifier
	} else {
		if dcl.StringCanonicalize(rawDesired.LastModifier, rawNew.LastModifier) {
			rawNew.LastModifier = rawDesired.LastModifier
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Client) && dcl.IsEmptyValueIndirect(rawDesired.Client) {
		rawNew.Client = rawDesired.Client
	} else {
		if dcl.StringCanonicalize(rawDesired.Client, rawNew.Client) {
			rawNew.Client = rawDesired.Client
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ClientVersion) && dcl.IsEmptyValueIndirect(rawDesired.ClientVersion) {
		rawNew.ClientVersion = rawDesired.ClientVersion
	} else {
		if dcl.StringCanonicalize(rawDesired.ClientVersion, rawNew.ClientVersion) {
			rawNew.ClientVersion = rawDesired.ClientVersion
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Ingress) && dcl.IsEmptyValueIndirect(rawDesired.Ingress) {
		rawNew.Ingress = rawDesired.Ingress
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.LaunchStage) && dcl.IsEmptyValueIndirect(rawDesired.LaunchStage) {
		rawNew.LaunchStage = rawDesired.LaunchStage
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.BinaryAuthorization) && dcl.IsEmptyValueIndirect(rawDesired.BinaryAuthorization) {
		rawNew.BinaryAuthorization = rawDesired.BinaryAuthorization
	} else {
		rawNew.BinaryAuthorization = canonicalizeNewServiceBinaryAuthorization(c, rawDesired.BinaryAuthorization, rawNew.BinaryAuthorization)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Template) && dcl.IsEmptyValueIndirect(rawDesired.Template) {
		rawNew.Template = rawDesired.Template
	} else {
		rawNew.Template = canonicalizeNewServiceTemplate(c, rawDesired.Template, rawNew.Template)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Traffic) && dcl.IsEmptyValueIndirect(rawDesired.Traffic) {
		rawNew.Traffic = rawDesired.Traffic
	} else {
		rawNew.Traffic = canonicalizeNewServiceTrafficSlice(c, rawDesired.Traffic, rawNew.Traffic)
	}

	if dcl.IsEmptyValueIndirect(rawNew.TerminalCondition) && dcl.IsEmptyValueIndirect(rawDesired.TerminalCondition) {
		rawNew.TerminalCondition = rawDesired.TerminalCondition
	} else {
		rawNew.TerminalCondition = canonicalizeNewServiceTerminalCondition(c, rawDesired.TerminalCondition, rawNew.TerminalCondition)
	}

	if dcl.IsEmptyValueIndirect(rawNew.LatestReadyRevision) && dcl.IsEmptyValueIndirect(rawDesired.LatestReadyRevision) {
		rawNew.LatestReadyRevision = rawDesired.LatestReadyRevision
	} else {
		if dcl.StringCanonicalize(rawDesired.LatestReadyRevision, rawNew.LatestReadyRevision) {
			rawNew.LatestReadyRevision = rawDesired.LatestReadyRevision
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.LatestCreatedRevision) && dcl.IsEmptyValueIndirect(rawDesired.LatestCreatedRevision) {
		rawNew.LatestCreatedRevision = rawDesired.LatestCreatedRevision
	} else {
		if dcl.StringCanonicalize(rawDesired.LatestCreatedRevision, rawNew.LatestCreatedRevision) {
			rawNew.LatestCreatedRevision = rawDesired.LatestCreatedRevision
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.TrafficStatuses) && dcl.IsEmptyValueIndirect(rawDesired.TrafficStatuses) {
		rawNew.TrafficStatuses = rawDesired.TrafficStatuses
	} else {
		rawNew.TrafficStatuses = canonicalizeNewServiceTrafficStatusesSlice(c, rawDesired.TrafficStatuses, rawNew.TrafficStatuses)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Uri) && dcl.IsEmptyValueIndirect(rawDesired.Uri) {
		rawNew.Uri = rawDesired.Uri
	} else {
		if dcl.StringCanonicalize(rawDesired.Uri, rawNew.Uri) {
			rawNew.Uri = rawDesired.Uri
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Reconciling) && dcl.IsEmptyValueIndirect(rawDesired.Reconciling) {
		rawNew.Reconciling = rawDesired.Reconciling
	} else {
		if dcl.BoolCanonicalize(rawDesired.Reconciling, rawNew.Reconciling) {
			rawNew.Reconciling = rawDesired.Reconciling
		}
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

func canonicalizeServiceBinaryAuthorization(des, initial *ServiceBinaryAuthorization, opts ...dcl.ApplyOption) *ServiceBinaryAuthorization {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceBinaryAuthorization{}

	if dcl.BoolCanonicalize(des.UseDefault, initial.UseDefault) || dcl.IsZeroValue(des.UseDefault) {
		cDes.UseDefault = initial.UseDefault
	} else {
		cDes.UseDefault = des.UseDefault
	}
	if dcl.StringCanonicalize(des.BreakglassJustification, initial.BreakglassJustification) || dcl.IsZeroValue(des.BreakglassJustification) {
		cDes.BreakglassJustification = initial.BreakglassJustification
	} else {
		cDes.BreakglassJustification = des.BreakglassJustification
	}

	return cDes
}

func canonicalizeServiceBinaryAuthorizationSlice(des, initial []ServiceBinaryAuthorization, opts ...dcl.ApplyOption) []ServiceBinaryAuthorization {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceBinaryAuthorization, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceBinaryAuthorization(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceBinaryAuthorization, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceBinaryAuthorization(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceBinaryAuthorization(c *Client, des, nw *ServiceBinaryAuthorization) *ServiceBinaryAuthorization {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceBinaryAuthorization while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.UseDefault, nw.UseDefault) {
		nw.UseDefault = des.UseDefault
	}
	if dcl.StringCanonicalize(des.BreakglassJustification, nw.BreakglassJustification) {
		nw.BreakglassJustification = des.BreakglassJustification
	}

	return nw
}

func canonicalizeNewServiceBinaryAuthorizationSet(c *Client, des, nw []ServiceBinaryAuthorization) []ServiceBinaryAuthorization {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ServiceBinaryAuthorization
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareServiceBinaryAuthorizationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewServiceBinaryAuthorization(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewServiceBinaryAuthorizationSlice(c *Client, des, nw []ServiceBinaryAuthorization) []ServiceBinaryAuthorization {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceBinaryAuthorization
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceBinaryAuthorization(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplate(des, initial *ServiceTemplate, opts ...dcl.ApplyOption) *ServiceTemplate {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplate{}

	if dcl.StringCanonicalize(des.Revision, initial.Revision) || dcl.IsZeroValue(des.Revision) {
		cDes.Revision = initial.Revision
	} else {
		cDes.Revision = des.Revision
	}
	if dcl.IsZeroValue(des.Labels) || (dcl.IsEmptyValueIndirect(des.Labels) && dcl.IsEmptyValueIndirect(initial.Labels)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Labels = initial.Labels
	} else {
		cDes.Labels = des.Labels
	}
	if dcl.IsZeroValue(des.Annotations) || (dcl.IsEmptyValueIndirect(des.Annotations) && dcl.IsEmptyValueIndirect(initial.Annotations)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Annotations = initial.Annotations
	} else {
		cDes.Annotations = des.Annotations
	}
	cDes.Scaling = canonicalizeServiceTemplateScaling(des.Scaling, initial.Scaling, opts...)
	cDes.VPCAccess = canonicalizeServiceTemplateVPCAccess(des.VPCAccess, initial.VPCAccess, opts...)
	if dcl.IsZeroValue(des.ContainerConcurrency) || (dcl.IsEmptyValueIndirect(des.ContainerConcurrency) && dcl.IsEmptyValueIndirect(initial.ContainerConcurrency)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ContainerConcurrency = initial.ContainerConcurrency
	} else {
		cDes.ContainerConcurrency = des.ContainerConcurrency
	}
	if dcl.StringCanonicalize(des.Timeout, initial.Timeout) || dcl.IsZeroValue(des.Timeout) {
		cDes.Timeout = initial.Timeout
	} else {
		cDes.Timeout = des.Timeout
	}
	if dcl.IsZeroValue(des.ServiceAccount) || (dcl.IsEmptyValueIndirect(des.ServiceAccount) && dcl.IsEmptyValueIndirect(initial.ServiceAccount)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ServiceAccount = initial.ServiceAccount
	} else {
		cDes.ServiceAccount = des.ServiceAccount
	}
	cDes.Containers = canonicalizeServiceTemplateContainersSlice(des.Containers, initial.Containers, opts...)
	cDes.Volumes = canonicalizeServiceTemplateVolumesSlice(des.Volumes, initial.Volumes, opts...)
	if dcl.IsZeroValue(des.ExecutionEnvironment) || (dcl.IsEmptyValueIndirect(des.ExecutionEnvironment) && dcl.IsEmptyValueIndirect(initial.ExecutionEnvironment)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ExecutionEnvironment = initial.ExecutionEnvironment
	} else {
		cDes.ExecutionEnvironment = des.ExecutionEnvironment
	}

	return cDes
}

func canonicalizeServiceTemplateSlice(des, initial []ServiceTemplate, opts ...dcl.ApplyOption) []ServiceTemplate {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplate, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplate(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplate, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplate(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplate(c *Client, des, nw *ServiceTemplate) *ServiceTemplate {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplate while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Revision, nw.Revision) {
		nw.Revision = des.Revision
	}
	nw.Scaling = canonicalizeNewServiceTemplateScaling(c, des.Scaling, nw.Scaling)
	nw.VPCAccess = canonicalizeNewServiceTemplateVPCAccess(c, des.VPCAccess, nw.VPCAccess)
	if dcl.StringCanonicalize(des.Timeout, nw.Timeout) {
		nw.Timeout = des.Timeout
	}
	nw.Containers = canonicalizeNewServiceTemplateContainersSlice(c, des.Containers, nw.Containers)
	nw.Volumes = canonicalizeNewServiceTemplateVolumesSlice(c, des.Volumes, nw.Volumes)

	return nw
}

func canonicalizeNewServiceTemplateSet(c *Client, des, nw []ServiceTemplate) []ServiceTemplate {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ServiceTemplate
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareServiceTemplateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewServiceTemplate(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewServiceTemplateSlice(c *Client, des, nw []ServiceTemplate) []ServiceTemplate {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplate
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplate(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplateScaling(des, initial *ServiceTemplateScaling, opts ...dcl.ApplyOption) *ServiceTemplateScaling {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplateScaling{}

	if dcl.IsZeroValue(des.MinInstanceCount) || (dcl.IsEmptyValueIndirect(des.MinInstanceCount) && dcl.IsEmptyValueIndirect(initial.MinInstanceCount)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MinInstanceCount = initial.MinInstanceCount
	} else {
		cDes.MinInstanceCount = des.MinInstanceCount
	}
	if dcl.IsZeroValue(des.MaxInstanceCount) || (dcl.IsEmptyValueIndirect(des.MaxInstanceCount) && dcl.IsEmptyValueIndirect(initial.MaxInstanceCount)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MaxInstanceCount = initial.MaxInstanceCount
	} else {
		cDes.MaxInstanceCount = des.MaxInstanceCount
	}

	return cDes
}

func canonicalizeServiceTemplateScalingSlice(des, initial []ServiceTemplateScaling, opts ...dcl.ApplyOption) []ServiceTemplateScaling {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplateScaling, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplateScaling(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplateScaling, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplateScaling(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplateScaling(c *Client, des, nw *ServiceTemplateScaling) *ServiceTemplateScaling {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplateScaling while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewServiceTemplateScalingSet(c *Client, des, nw []ServiceTemplateScaling) []ServiceTemplateScaling {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ServiceTemplateScaling
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareServiceTemplateScalingNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewServiceTemplateScaling(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewServiceTemplateScalingSlice(c *Client, des, nw []ServiceTemplateScaling) []ServiceTemplateScaling {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplateScaling
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplateScaling(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplateVPCAccess(des, initial *ServiceTemplateVPCAccess, opts ...dcl.ApplyOption) *ServiceTemplateVPCAccess {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplateVPCAccess{}

	if dcl.IsZeroValue(des.Connector) || (dcl.IsEmptyValueIndirect(des.Connector) && dcl.IsEmptyValueIndirect(initial.Connector)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Connector = initial.Connector
	} else {
		cDes.Connector = des.Connector
	}
	if dcl.IsZeroValue(des.Egress) || (dcl.IsEmptyValueIndirect(des.Egress) && dcl.IsEmptyValueIndirect(initial.Egress)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Egress = initial.Egress
	} else {
		cDes.Egress = des.Egress
	}

	return cDes
}

func canonicalizeServiceTemplateVPCAccessSlice(des, initial []ServiceTemplateVPCAccess, opts ...dcl.ApplyOption) []ServiceTemplateVPCAccess {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplateVPCAccess, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplateVPCAccess(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplateVPCAccess, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplateVPCAccess(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplateVPCAccess(c *Client, des, nw *ServiceTemplateVPCAccess) *ServiceTemplateVPCAccess {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplateVPCAccess while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewServiceTemplateVPCAccessSet(c *Client, des, nw []ServiceTemplateVPCAccess) []ServiceTemplateVPCAccess {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ServiceTemplateVPCAccess
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareServiceTemplateVPCAccessNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewServiceTemplateVPCAccess(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewServiceTemplateVPCAccessSlice(c *Client, des, nw []ServiceTemplateVPCAccess) []ServiceTemplateVPCAccess {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplateVPCAccess
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplateVPCAccess(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplateContainers(des, initial *ServiceTemplateContainers, opts ...dcl.ApplyOption) *ServiceTemplateContainers {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplateContainers{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Image, initial.Image) || dcl.IsZeroValue(des.Image) {
		cDes.Image = initial.Image
	} else {
		cDes.Image = des.Image
	}
	if dcl.StringArrayCanonicalize(des.Command, initial.Command) {
		cDes.Command = initial.Command
	} else {
		cDes.Command = des.Command
	}
	if dcl.StringArrayCanonicalize(des.Args, initial.Args) {
		cDes.Args = initial.Args
	} else {
		cDes.Args = des.Args
	}
	cDes.Env = canonicalizeServiceTemplateContainersEnvSlice(des.Env, initial.Env, opts...)
	cDes.Resources = canonicalizeServiceTemplateContainersResources(des.Resources, initial.Resources, opts...)
	cDes.Ports = canonicalizeServiceTemplateContainersPortsSlice(des.Ports, initial.Ports, opts...)
	cDes.VolumeMounts = canonicalizeServiceTemplateContainersVolumeMountsSlice(des.VolumeMounts, initial.VolumeMounts, opts...)

	return cDes
}

func canonicalizeServiceTemplateContainersSlice(des, initial []ServiceTemplateContainers, opts ...dcl.ApplyOption) []ServiceTemplateContainers {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplateContainers, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplateContainers(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplateContainers, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplateContainers(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplateContainers(c *Client, des, nw *ServiceTemplateContainers) *ServiceTemplateContainers {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplateContainers while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Image, nw.Image) {
		nw.Image = des.Image
	}
	if dcl.StringArrayCanonicalize(des.Command, nw.Command) {
		nw.Command = des.Command
	}
	if dcl.StringArrayCanonicalize(des.Args, nw.Args) {
		nw.Args = des.Args
	}
	nw.Env = canonicalizeNewServiceTemplateContainersEnvSlice(c, des.Env, nw.Env)
	nw.Resources = canonicalizeNewServiceTemplateContainersResources(c, des.Resources, nw.Resources)
	nw.Ports = canonicalizeNewServiceTemplateContainersPortsSlice(c, des.Ports, nw.Ports)
	nw.VolumeMounts = canonicalizeNewServiceTemplateContainersVolumeMountsSlice(c, des.VolumeMounts, nw.VolumeMounts)

	return nw
}

func canonicalizeNewServiceTemplateContainersSet(c *Client, des, nw []ServiceTemplateContainers) []ServiceTemplateContainers {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ServiceTemplateContainers
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareServiceTemplateContainersNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewServiceTemplateContainers(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewServiceTemplateContainersSlice(c *Client, des, nw []ServiceTemplateContainers) []ServiceTemplateContainers {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplateContainers
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplateContainers(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplateContainersEnv(des, initial *ServiceTemplateContainersEnv, opts ...dcl.ApplyOption) *ServiceTemplateContainersEnv {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.Value != nil || (initial != nil && initial.Value != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.ValueSource) {
			des.Value = nil
			if initial != nil {
				initial.Value = nil
			}
		}
	}

	if des.ValueSource != nil || (initial != nil && initial.ValueSource != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Value) {
			des.ValueSource = nil
			if initial != nil {
				initial.ValueSource = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplateContainersEnv{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		cDes.Value = initial.Value
	} else {
		cDes.Value = des.Value
	}
	cDes.ValueSource = canonicalizeServiceTemplateContainersEnvValueSource(des.ValueSource, initial.ValueSource, opts...)

	return cDes
}

func canonicalizeServiceTemplateContainersEnvSlice(des, initial []ServiceTemplateContainersEnv, opts ...dcl.ApplyOption) []ServiceTemplateContainersEnv {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplateContainersEnv, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplateContainersEnv(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplateContainersEnv, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplateContainersEnv(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplateContainersEnv(c *Client, des, nw *ServiceTemplateContainersEnv) *ServiceTemplateContainersEnv {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplateContainersEnv while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}
	nw.ValueSource = canonicalizeNewServiceTemplateContainersEnvValueSource(c, des.ValueSource, nw.ValueSource)

	return nw
}

func canonicalizeNewServiceTemplateContainersEnvSet(c *Client, des, nw []ServiceTemplateContainersEnv) []ServiceTemplateContainersEnv {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ServiceTemplateContainersEnv
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareServiceTemplateContainersEnvNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewServiceTemplateContainersEnv(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewServiceTemplateContainersEnvSlice(c *Client, des, nw []ServiceTemplateContainersEnv) []ServiceTemplateContainersEnv {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplateContainersEnv
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplateContainersEnv(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplateContainersEnvValueSource(des, initial *ServiceTemplateContainersEnvValueSource, opts ...dcl.ApplyOption) *ServiceTemplateContainersEnvValueSource {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplateContainersEnvValueSource{}

	cDes.SecretKeyRef = canonicalizeServiceTemplateContainersEnvValueSourceSecretKeyRef(des.SecretKeyRef, initial.SecretKeyRef, opts...)

	return cDes
}

func canonicalizeServiceTemplateContainersEnvValueSourceSlice(des, initial []ServiceTemplateContainersEnvValueSource, opts ...dcl.ApplyOption) []ServiceTemplateContainersEnvValueSource {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplateContainersEnvValueSource, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplateContainersEnvValueSource(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplateContainersEnvValueSource, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplateContainersEnvValueSource(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplateContainersEnvValueSource(c *Client, des, nw *ServiceTemplateContainersEnvValueSource) *ServiceTemplateContainersEnvValueSource {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplateContainersEnvValueSource while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.SecretKeyRef = canonicalizeNewServiceTemplateContainersEnvValueSourceSecretKeyRef(c, des.SecretKeyRef, nw.SecretKeyRef)

	return nw
}

func canonicalizeNewServiceTemplateContainersEnvValueSourceSet(c *Client, des, nw []ServiceTemplateContainersEnvValueSource) []ServiceTemplateContainersEnvValueSource {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ServiceTemplateContainersEnvValueSource
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareServiceTemplateContainersEnvValueSourceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewServiceTemplateContainersEnvValueSource(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewServiceTemplateContainersEnvValueSourceSlice(c *Client, des, nw []ServiceTemplateContainersEnvValueSource) []ServiceTemplateContainersEnvValueSource {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplateContainersEnvValueSource
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplateContainersEnvValueSource(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplateContainersEnvValueSourceSecretKeyRef(des, initial *ServiceTemplateContainersEnvValueSourceSecretKeyRef, opts ...dcl.ApplyOption) *ServiceTemplateContainersEnvValueSourceSecretKeyRef {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplateContainersEnvValueSourceSecretKeyRef{}

	if dcl.IsZeroValue(des.Secret) || (dcl.IsEmptyValueIndirect(des.Secret) && dcl.IsEmptyValueIndirect(initial.Secret)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Secret = initial.Secret
	} else {
		cDes.Secret = des.Secret
	}
	if dcl.IsZeroValue(des.Version) || (dcl.IsEmptyValueIndirect(des.Version) && dcl.IsEmptyValueIndirect(initial.Version)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Version = initial.Version
	} else {
		cDes.Version = des.Version
	}

	return cDes
}

func canonicalizeServiceTemplateContainersEnvValueSourceSecretKeyRefSlice(des, initial []ServiceTemplateContainersEnvValueSourceSecretKeyRef, opts ...dcl.ApplyOption) []ServiceTemplateContainersEnvValueSourceSecretKeyRef {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplateContainersEnvValueSourceSecretKeyRef, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplateContainersEnvValueSourceSecretKeyRef(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplateContainersEnvValueSourceSecretKeyRef, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplateContainersEnvValueSourceSecretKeyRef(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplateContainersEnvValueSourceSecretKeyRef(c *Client, des, nw *ServiceTemplateContainersEnvValueSourceSecretKeyRef) *ServiceTemplateContainersEnvValueSourceSecretKeyRef {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplateContainersEnvValueSourceSecretKeyRef while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewServiceTemplateContainersEnvValueSourceSecretKeyRefSet(c *Client, des, nw []ServiceTemplateContainersEnvValueSourceSecretKeyRef) []ServiceTemplateContainersEnvValueSourceSecretKeyRef {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ServiceTemplateContainersEnvValueSourceSecretKeyRef
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareServiceTemplateContainersEnvValueSourceSecretKeyRefNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewServiceTemplateContainersEnvValueSourceSecretKeyRef(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewServiceTemplateContainersEnvValueSourceSecretKeyRefSlice(c *Client, des, nw []ServiceTemplateContainersEnvValueSourceSecretKeyRef) []ServiceTemplateContainersEnvValueSourceSecretKeyRef {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplateContainersEnvValueSourceSecretKeyRef
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplateContainersEnvValueSourceSecretKeyRef(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplateContainersResources(des, initial *ServiceTemplateContainersResources, opts ...dcl.ApplyOption) *ServiceTemplateContainersResources {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplateContainersResources{}

	if dcl.IsZeroValue(des.Limits) || (dcl.IsEmptyValueIndirect(des.Limits) && dcl.IsEmptyValueIndirect(initial.Limits)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Limits = initial.Limits
	} else {
		cDes.Limits = des.Limits
	}
	if dcl.BoolCanonicalize(des.CpuIdle, initial.CpuIdle) || dcl.IsZeroValue(des.CpuIdle) {
		cDes.CpuIdle = initial.CpuIdle
	} else {
		cDes.CpuIdle = des.CpuIdle
	}

	return cDes
}

func canonicalizeServiceTemplateContainersResourcesSlice(des, initial []ServiceTemplateContainersResources, opts ...dcl.ApplyOption) []ServiceTemplateContainersResources {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplateContainersResources, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplateContainersResources(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplateContainersResources, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplateContainersResources(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplateContainersResources(c *Client, des, nw *ServiceTemplateContainersResources) *ServiceTemplateContainersResources {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplateContainersResources while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.CpuIdle, nw.CpuIdle) {
		nw.CpuIdle = des.CpuIdle
	}

	return nw
}

func canonicalizeNewServiceTemplateContainersResourcesSet(c *Client, des, nw []ServiceTemplateContainersResources) []ServiceTemplateContainersResources {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ServiceTemplateContainersResources
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareServiceTemplateContainersResourcesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewServiceTemplateContainersResources(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewServiceTemplateContainersResourcesSlice(c *Client, des, nw []ServiceTemplateContainersResources) []ServiceTemplateContainersResources {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplateContainersResources
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplateContainersResources(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplateContainersPorts(des, initial *ServiceTemplateContainersPorts, opts ...dcl.ApplyOption) *ServiceTemplateContainersPorts {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplateContainersPorts{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.IsZeroValue(des.ContainerPort) || (dcl.IsEmptyValueIndirect(des.ContainerPort) && dcl.IsEmptyValueIndirect(initial.ContainerPort)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ContainerPort = initial.ContainerPort
	} else {
		cDes.ContainerPort = des.ContainerPort
	}

	return cDes
}

func canonicalizeServiceTemplateContainersPortsSlice(des, initial []ServiceTemplateContainersPorts, opts ...dcl.ApplyOption) []ServiceTemplateContainersPorts {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplateContainersPorts, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplateContainersPorts(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplateContainersPorts, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplateContainersPorts(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplateContainersPorts(c *Client, des, nw *ServiceTemplateContainersPorts) *ServiceTemplateContainersPorts {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplateContainersPorts while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewServiceTemplateContainersPortsSet(c *Client, des, nw []ServiceTemplateContainersPorts) []ServiceTemplateContainersPorts {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ServiceTemplateContainersPorts
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareServiceTemplateContainersPortsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewServiceTemplateContainersPorts(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewServiceTemplateContainersPortsSlice(c *Client, des, nw []ServiceTemplateContainersPorts) []ServiceTemplateContainersPorts {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplateContainersPorts
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplateContainersPorts(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplateContainersVolumeMounts(des, initial *ServiceTemplateContainersVolumeMounts, opts ...dcl.ApplyOption) *ServiceTemplateContainersVolumeMounts {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplateContainersVolumeMounts{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.StringCanonicalize(des.MountPath, initial.MountPath) || dcl.IsZeroValue(des.MountPath) {
		cDes.MountPath = initial.MountPath
	} else {
		cDes.MountPath = des.MountPath
	}

	return cDes
}

func canonicalizeServiceTemplateContainersVolumeMountsSlice(des, initial []ServiceTemplateContainersVolumeMounts, opts ...dcl.ApplyOption) []ServiceTemplateContainersVolumeMounts {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplateContainersVolumeMounts, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplateContainersVolumeMounts(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplateContainersVolumeMounts, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplateContainersVolumeMounts(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplateContainersVolumeMounts(c *Client, des, nw *ServiceTemplateContainersVolumeMounts) *ServiceTemplateContainersVolumeMounts {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplateContainersVolumeMounts while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.MountPath, nw.MountPath) {
		nw.MountPath = des.MountPath
	}

	return nw
}

func canonicalizeNewServiceTemplateContainersVolumeMountsSet(c *Client, des, nw []ServiceTemplateContainersVolumeMounts) []ServiceTemplateContainersVolumeMounts {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ServiceTemplateContainersVolumeMounts
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareServiceTemplateContainersVolumeMountsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewServiceTemplateContainersVolumeMounts(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewServiceTemplateContainersVolumeMountsSlice(c *Client, des, nw []ServiceTemplateContainersVolumeMounts) []ServiceTemplateContainersVolumeMounts {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplateContainersVolumeMounts
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplateContainersVolumeMounts(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplateVolumes(des, initial *ServiceTemplateVolumes, opts ...dcl.ApplyOption) *ServiceTemplateVolumes {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.Secret != nil || (initial != nil && initial.Secret != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.CloudSqlInstance) {
			des.Secret = nil
			if initial != nil {
				initial.Secret = nil
			}
		}
	}

	if des.CloudSqlInstance != nil || (initial != nil && initial.CloudSqlInstance != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Secret) {
			des.CloudSqlInstance = nil
			if initial != nil {
				initial.CloudSqlInstance = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplateVolumes{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	cDes.Secret = canonicalizeServiceTemplateVolumesSecret(des.Secret, initial.Secret, opts...)
	cDes.CloudSqlInstance = canonicalizeServiceTemplateVolumesCloudSqlInstance(des.CloudSqlInstance, initial.CloudSqlInstance, opts...)

	return cDes
}

func canonicalizeServiceTemplateVolumesSlice(des, initial []ServiceTemplateVolumes, opts ...dcl.ApplyOption) []ServiceTemplateVolumes {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplateVolumes, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplateVolumes(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplateVolumes, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplateVolumes(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplateVolumes(c *Client, des, nw *ServiceTemplateVolumes) *ServiceTemplateVolumes {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplateVolumes while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	nw.Secret = canonicalizeNewServiceTemplateVolumesSecret(c, des.Secret, nw.Secret)
	nw.CloudSqlInstance = canonicalizeNewServiceTemplateVolumesCloudSqlInstance(c, des.CloudSqlInstance, nw.CloudSqlInstance)

	return nw
}

func canonicalizeNewServiceTemplateVolumesSet(c *Client, des, nw []ServiceTemplateVolumes) []ServiceTemplateVolumes {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ServiceTemplateVolumes
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareServiceTemplateVolumesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewServiceTemplateVolumes(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewServiceTemplateVolumesSlice(c *Client, des, nw []ServiceTemplateVolumes) []ServiceTemplateVolumes {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplateVolumes
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplateVolumes(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplateVolumesSecret(des, initial *ServiceTemplateVolumesSecret, opts ...dcl.ApplyOption) *ServiceTemplateVolumesSecret {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplateVolumesSecret{}

	if dcl.IsZeroValue(des.Secret) || (dcl.IsEmptyValueIndirect(des.Secret) && dcl.IsEmptyValueIndirect(initial.Secret)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Secret = initial.Secret
	} else {
		cDes.Secret = des.Secret
	}
	cDes.Items = canonicalizeServiceTemplateVolumesSecretItemsSlice(des.Items, initial.Items, opts...)
	if dcl.IsZeroValue(des.DefaultMode) || (dcl.IsEmptyValueIndirect(des.DefaultMode) && dcl.IsEmptyValueIndirect(initial.DefaultMode)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.DefaultMode = initial.DefaultMode
	} else {
		cDes.DefaultMode = des.DefaultMode
	}

	return cDes
}

func canonicalizeServiceTemplateVolumesSecretSlice(des, initial []ServiceTemplateVolumesSecret, opts ...dcl.ApplyOption) []ServiceTemplateVolumesSecret {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplateVolumesSecret, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplateVolumesSecret(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplateVolumesSecret, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplateVolumesSecret(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplateVolumesSecret(c *Client, des, nw *ServiceTemplateVolumesSecret) *ServiceTemplateVolumesSecret {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplateVolumesSecret while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Items = canonicalizeNewServiceTemplateVolumesSecretItemsSlice(c, des.Items, nw.Items)

	return nw
}

func canonicalizeNewServiceTemplateVolumesSecretSet(c *Client, des, nw []ServiceTemplateVolumesSecret) []ServiceTemplateVolumesSecret {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ServiceTemplateVolumesSecret
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareServiceTemplateVolumesSecretNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewServiceTemplateVolumesSecret(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewServiceTemplateVolumesSecretSlice(c *Client, des, nw []ServiceTemplateVolumesSecret) []ServiceTemplateVolumesSecret {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplateVolumesSecret
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplateVolumesSecret(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplateVolumesSecretItems(des, initial *ServiceTemplateVolumesSecretItems, opts ...dcl.ApplyOption) *ServiceTemplateVolumesSecretItems {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplateVolumesSecretItems{}

	if dcl.StringCanonicalize(des.Path, initial.Path) || dcl.IsZeroValue(des.Path) {
		cDes.Path = initial.Path
	} else {
		cDes.Path = des.Path
	}
	if dcl.IsZeroValue(des.Version) || (dcl.IsEmptyValueIndirect(des.Version) && dcl.IsEmptyValueIndirect(initial.Version)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Version = initial.Version
	} else {
		cDes.Version = des.Version
	}
	if dcl.IsZeroValue(des.Mode) || (dcl.IsEmptyValueIndirect(des.Mode) && dcl.IsEmptyValueIndirect(initial.Mode)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Mode = initial.Mode
	} else {
		cDes.Mode = des.Mode
	}

	return cDes
}

func canonicalizeServiceTemplateVolumesSecretItemsSlice(des, initial []ServiceTemplateVolumesSecretItems, opts ...dcl.ApplyOption) []ServiceTemplateVolumesSecretItems {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplateVolumesSecretItems, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplateVolumesSecretItems(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplateVolumesSecretItems, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplateVolumesSecretItems(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplateVolumesSecretItems(c *Client, des, nw *ServiceTemplateVolumesSecretItems) *ServiceTemplateVolumesSecretItems {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplateVolumesSecretItems while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Path, nw.Path) {
		nw.Path = des.Path
	}

	return nw
}

func canonicalizeNewServiceTemplateVolumesSecretItemsSet(c *Client, des, nw []ServiceTemplateVolumesSecretItems) []ServiceTemplateVolumesSecretItems {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ServiceTemplateVolumesSecretItems
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareServiceTemplateVolumesSecretItemsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewServiceTemplateVolumesSecretItems(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewServiceTemplateVolumesSecretItemsSlice(c *Client, des, nw []ServiceTemplateVolumesSecretItems) []ServiceTemplateVolumesSecretItems {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplateVolumesSecretItems
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplateVolumesSecretItems(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTemplateVolumesCloudSqlInstance(des, initial *ServiceTemplateVolumesCloudSqlInstance, opts ...dcl.ApplyOption) *ServiceTemplateVolumesCloudSqlInstance {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTemplateVolumesCloudSqlInstance{}

	if dcl.StringArrayCanonicalize(des.Instances, initial.Instances) {
		cDes.Instances = initial.Instances
	} else {
		cDes.Instances = des.Instances
	}

	return cDes
}

func canonicalizeServiceTemplateVolumesCloudSqlInstanceSlice(des, initial []ServiceTemplateVolumesCloudSqlInstance, opts ...dcl.ApplyOption) []ServiceTemplateVolumesCloudSqlInstance {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTemplateVolumesCloudSqlInstance, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTemplateVolumesCloudSqlInstance(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTemplateVolumesCloudSqlInstance, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTemplateVolumesCloudSqlInstance(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTemplateVolumesCloudSqlInstance(c *Client, des, nw *ServiceTemplateVolumesCloudSqlInstance) *ServiceTemplateVolumesCloudSqlInstance {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTemplateVolumesCloudSqlInstance while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.Instances, nw.Instances) {
		nw.Instances = des.Instances
	}

	return nw
}

func canonicalizeNewServiceTemplateVolumesCloudSqlInstanceSet(c *Client, des, nw []ServiceTemplateVolumesCloudSqlInstance) []ServiceTemplateVolumesCloudSqlInstance {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ServiceTemplateVolumesCloudSqlInstance
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareServiceTemplateVolumesCloudSqlInstanceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewServiceTemplateVolumesCloudSqlInstance(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewServiceTemplateVolumesCloudSqlInstanceSlice(c *Client, des, nw []ServiceTemplateVolumesCloudSqlInstance) []ServiceTemplateVolumesCloudSqlInstance {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTemplateVolumesCloudSqlInstance
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTemplateVolumesCloudSqlInstance(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTraffic(des, initial *ServiceTraffic, opts ...dcl.ApplyOption) *ServiceTraffic {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTraffic{}

	if dcl.IsZeroValue(des.Type) || (dcl.IsEmptyValueIndirect(des.Type) && dcl.IsEmptyValueIndirect(initial.Type)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Type = initial.Type
	} else {
		cDes.Type = des.Type
	}
	if dcl.StringCanonicalize(des.Revision, initial.Revision) || dcl.IsZeroValue(des.Revision) {
		cDes.Revision = initial.Revision
	} else {
		cDes.Revision = des.Revision
	}
	if dcl.IsZeroValue(des.Percent) || (dcl.IsEmptyValueIndirect(des.Percent) && dcl.IsEmptyValueIndirect(initial.Percent)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Percent = initial.Percent
	} else {
		cDes.Percent = des.Percent
	}
	if dcl.StringCanonicalize(des.Tag, initial.Tag) || dcl.IsZeroValue(des.Tag) {
		cDes.Tag = initial.Tag
	} else {
		cDes.Tag = des.Tag
	}

	return cDes
}

func canonicalizeServiceTrafficSlice(des, initial []ServiceTraffic, opts ...dcl.ApplyOption) []ServiceTraffic {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTraffic, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTraffic(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTraffic, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTraffic(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTraffic(c *Client, des, nw *ServiceTraffic) *ServiceTraffic {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTraffic while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Revision, nw.Revision) {
		nw.Revision = des.Revision
	}
	if dcl.StringCanonicalize(des.Tag, nw.Tag) {
		nw.Tag = des.Tag
	}

	return nw
}

func canonicalizeNewServiceTrafficSet(c *Client, des, nw []ServiceTraffic) []ServiceTraffic {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ServiceTraffic
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareServiceTrafficNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewServiceTraffic(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewServiceTrafficSlice(c *Client, des, nw []ServiceTraffic) []ServiceTraffic {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTraffic
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTraffic(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTerminalCondition(des, initial *ServiceTerminalCondition, opts ...dcl.ApplyOption) *ServiceTerminalCondition {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.Reason != nil || (initial != nil && initial.Reason != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.RevisionReason, des.JobReason) {
			des.Reason = nil
			if initial != nil {
				initial.Reason = nil
			}
		}
	}

	if des.RevisionReason != nil || (initial != nil && initial.RevisionReason != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Reason, des.JobReason) {
			des.RevisionReason = nil
			if initial != nil {
				initial.RevisionReason = nil
			}
		}
	}

	if des.JobReason != nil || (initial != nil && initial.JobReason != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Reason, des.RevisionReason) {
			des.JobReason = nil
			if initial != nil {
				initial.JobReason = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTerminalCondition{}

	if dcl.StringCanonicalize(des.Type, initial.Type) || dcl.IsZeroValue(des.Type) {
		cDes.Type = initial.Type
	} else {
		cDes.Type = des.Type
	}
	if dcl.IsZeroValue(des.State) || (dcl.IsEmptyValueIndirect(des.State) && dcl.IsEmptyValueIndirect(initial.State)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.State = initial.State
	} else {
		cDes.State = des.State
	}
	if dcl.StringCanonicalize(des.Message, initial.Message) || dcl.IsZeroValue(des.Message) {
		cDes.Message = initial.Message
	} else {
		cDes.Message = des.Message
	}
	if dcl.IsZeroValue(des.LastTransitionTime) || (dcl.IsEmptyValueIndirect(des.LastTransitionTime) && dcl.IsEmptyValueIndirect(initial.LastTransitionTime)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.LastTransitionTime = initial.LastTransitionTime
	} else {
		cDes.LastTransitionTime = des.LastTransitionTime
	}
	if dcl.IsZeroValue(des.Severity) || (dcl.IsEmptyValueIndirect(des.Severity) && dcl.IsEmptyValueIndirect(initial.Severity)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Severity = initial.Severity
	} else {
		cDes.Severity = des.Severity
	}
	if dcl.IsZeroValue(des.Reason) || (dcl.IsEmptyValueIndirect(des.Reason) && dcl.IsEmptyValueIndirect(initial.Reason)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Reason = initial.Reason
	} else {
		cDes.Reason = des.Reason
	}
	if dcl.IsZeroValue(des.RevisionReason) || (dcl.IsEmptyValueIndirect(des.RevisionReason) && dcl.IsEmptyValueIndirect(initial.RevisionReason)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.RevisionReason = initial.RevisionReason
	} else {
		cDes.RevisionReason = des.RevisionReason
	}
	if dcl.IsZeroValue(des.JobReason) || (dcl.IsEmptyValueIndirect(des.JobReason) && dcl.IsEmptyValueIndirect(initial.JobReason)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.JobReason = initial.JobReason
	} else {
		cDes.JobReason = des.JobReason
	}

	return cDes
}

func canonicalizeServiceTerminalConditionSlice(des, initial []ServiceTerminalCondition, opts ...dcl.ApplyOption) []ServiceTerminalCondition {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTerminalCondition, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTerminalCondition(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTerminalCondition, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTerminalCondition(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTerminalCondition(c *Client, des, nw *ServiceTerminalCondition) *ServiceTerminalCondition {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTerminalCondition while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Type, nw.Type) {
		nw.Type = des.Type
	}
	if dcl.StringCanonicalize(des.Message, nw.Message) {
		nw.Message = des.Message
	}

	return nw
}

func canonicalizeNewServiceTerminalConditionSet(c *Client, des, nw []ServiceTerminalCondition) []ServiceTerminalCondition {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ServiceTerminalCondition
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareServiceTerminalConditionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewServiceTerminalCondition(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewServiceTerminalConditionSlice(c *Client, des, nw []ServiceTerminalCondition) []ServiceTerminalCondition {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTerminalCondition
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTerminalCondition(c, &d, &n))
	}

	return items
}

func canonicalizeServiceTrafficStatuses(des, initial *ServiceTrafficStatuses, opts ...dcl.ApplyOption) *ServiceTrafficStatuses {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ServiceTrafficStatuses{}

	if dcl.IsZeroValue(des.Type) || (dcl.IsEmptyValueIndirect(des.Type) && dcl.IsEmptyValueIndirect(initial.Type)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Type = initial.Type
	} else {
		cDes.Type = des.Type
	}
	if dcl.StringCanonicalize(des.Revision, initial.Revision) || dcl.IsZeroValue(des.Revision) {
		cDes.Revision = initial.Revision
	} else {
		cDes.Revision = des.Revision
	}
	if dcl.IsZeroValue(des.Percent) || (dcl.IsEmptyValueIndirect(des.Percent) && dcl.IsEmptyValueIndirect(initial.Percent)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Percent = initial.Percent
	} else {
		cDes.Percent = des.Percent
	}
	if dcl.StringCanonicalize(des.Tag, initial.Tag) || dcl.IsZeroValue(des.Tag) {
		cDes.Tag = initial.Tag
	} else {
		cDes.Tag = des.Tag
	}
	if dcl.StringCanonicalize(des.Uri, initial.Uri) || dcl.IsZeroValue(des.Uri) {
		cDes.Uri = initial.Uri
	} else {
		cDes.Uri = des.Uri
	}

	return cDes
}

func canonicalizeServiceTrafficStatusesSlice(des, initial []ServiceTrafficStatuses, opts ...dcl.ApplyOption) []ServiceTrafficStatuses {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ServiceTrafficStatuses, 0, len(des))
		for _, d := range des {
			cd := canonicalizeServiceTrafficStatuses(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ServiceTrafficStatuses, 0, len(des))
	for i, d := range des {
		cd := canonicalizeServiceTrafficStatuses(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewServiceTrafficStatuses(c *Client, des, nw *ServiceTrafficStatuses) *ServiceTrafficStatuses {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ServiceTrafficStatuses while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Revision, nw.Revision) {
		nw.Revision = des.Revision
	}
	if dcl.StringCanonicalize(des.Tag, nw.Tag) {
		nw.Tag = des.Tag
	}
	if dcl.StringCanonicalize(des.Uri, nw.Uri) {
		nw.Uri = des.Uri
	}

	return nw
}

func canonicalizeNewServiceTrafficStatusesSet(c *Client, des, nw []ServiceTrafficStatuses) []ServiceTrafficStatuses {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ServiceTrafficStatuses
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareServiceTrafficStatusesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewServiceTrafficStatuses(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewServiceTrafficStatusesSlice(c *Client, des, nw []ServiceTrafficStatuses) []ServiceTrafficStatuses {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceTrafficStatuses
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceTrafficStatuses(c, &d, &n))
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
func diffService(c *Client, desired, actual *Service, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Generation, actual.Generation, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Generation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Annotations, actual.Annotations, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Annotations")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.ExpireTime, actual.ExpireTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExpireTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Creator, actual.Creator, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Creator")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LastModifier, actual.LastModifier, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LastModifier")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Client, actual.Client, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Client")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClientVersion, actual.ClientVersion, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("ClientVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Ingress, actual.Ingress, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Ingress")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LaunchStage, actual.LaunchStage, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("LaunchStage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BinaryAuthorization, actual.BinaryAuthorization, dcl.DiffInfo{ObjectFunction: compareServiceBinaryAuthorizationNewStyle, EmptyObject: EmptyServiceBinaryAuthorization, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("BinaryAuthorization")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Template, actual.Template, dcl.DiffInfo{ObjectFunction: compareServiceTemplateNewStyle, EmptyObject: EmptyServiceTemplate, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Template")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Traffic, actual.Traffic, dcl.DiffInfo{ObjectFunction: compareServiceTrafficNewStyle, EmptyObject: EmptyServiceTraffic, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Traffic")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TerminalCondition, actual.TerminalCondition, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareServiceTerminalConditionNewStyle, EmptyObject: EmptyServiceTerminalCondition, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TerminalCondition")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LatestReadyRevision, actual.LatestReadyRevision, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LatestReadyRevision")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LatestCreatedRevision, actual.LatestCreatedRevision, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LatestCreatedRevision")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TrafficStatuses, actual.TrafficStatuses, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareServiceTrafficStatusesNewStyle, EmptyObject: EmptyServiceTrafficStatuses, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TrafficStatuses")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Uri, actual.Uri, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Uri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Reconciling, actual.Reconciling, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Reconciling")); len(ds) != 0 || err != nil {
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
func compareServiceBinaryAuthorizationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceBinaryAuthorization)
	if !ok {
		desiredNotPointer, ok := d.(ServiceBinaryAuthorization)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceBinaryAuthorization or *ServiceBinaryAuthorization", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceBinaryAuthorization)
	if !ok {
		actualNotPointer, ok := a.(ServiceBinaryAuthorization)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceBinaryAuthorization", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.UseDefault, actual.UseDefault, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("UseDefault")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BreakglassJustification, actual.BreakglassJustification, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("BreakglassJustification")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplate)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplate or *ServiceTemplate", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplate)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplate", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Revision, actual.Revision, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Revision")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Annotations, actual.Annotations, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Annotations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Scaling, actual.Scaling, dcl.DiffInfo{ObjectFunction: compareServiceTemplateScalingNewStyle, EmptyObject: EmptyServiceTemplateScaling, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Scaling")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VPCAccess, actual.VPCAccess, dcl.DiffInfo{ObjectFunction: compareServiceTemplateVPCAccessNewStyle, EmptyObject: EmptyServiceTemplateVPCAccess, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("VpcAccess")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ContainerConcurrency, actual.ContainerConcurrency, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("MaxInstanceRequestConcurrency")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Timeout, actual.Timeout, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Timeout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceAccount, actual.ServiceAccount, dcl.DiffInfo{ServerDefault: true, Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("ServiceAccount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Containers, actual.Containers, dcl.DiffInfo{ObjectFunction: compareServiceTemplateContainersNewStyle, EmptyObject: EmptyServiceTemplateContainers, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Containers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Volumes, actual.Volumes, dcl.DiffInfo{ObjectFunction: compareServiceTemplateVolumesNewStyle, EmptyObject: EmptyServiceTemplateVolumes, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Volumes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExecutionEnvironment, actual.ExecutionEnvironment, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("ExecutionEnvironment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateScalingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplateScaling)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplateScaling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateScaling or *ServiceTemplateScaling", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplateScaling)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplateScaling)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateScaling", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MinInstanceCount, actual.MinInstanceCount, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("MinInstanceCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxInstanceCount, actual.MaxInstanceCount, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("MaxInstanceCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateVPCAccessNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplateVPCAccess)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplateVPCAccess)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateVPCAccess or *ServiceTemplateVPCAccess", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplateVPCAccess)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplateVPCAccess)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateVPCAccess", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Connector, actual.Connector, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Connector")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Egress, actual.Egress, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Egress")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateContainersNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplateContainers)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplateContainers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainers or *ServiceTemplateContainers", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplateContainers)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplateContainers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainers", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Image, actual.Image, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Image")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Command, actual.Command, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Command")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Args, actual.Args, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Args")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Env, actual.Env, dcl.DiffInfo{ObjectFunction: compareServiceTemplateContainersEnvNewStyle, EmptyObject: EmptyServiceTemplateContainersEnv, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Env")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Resources, actual.Resources, dcl.DiffInfo{ServerDefault: true, ObjectFunction: compareServiceTemplateContainersResourcesNewStyle, EmptyObject: EmptyServiceTemplateContainersResources, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Resources")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Ports, actual.Ports, dcl.DiffInfo{ServerDefault: true, ObjectFunction: compareServiceTemplateContainersPortsNewStyle, EmptyObject: EmptyServiceTemplateContainersPorts, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Ports")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VolumeMounts, actual.VolumeMounts, dcl.DiffInfo{ObjectFunction: compareServiceTemplateContainersVolumeMountsNewStyle, EmptyObject: EmptyServiceTemplateContainersVolumeMounts, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("VolumeMounts")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateContainersEnvNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplateContainersEnv)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplateContainersEnv)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainersEnv or *ServiceTemplateContainersEnv", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplateContainersEnv)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplateContainersEnv)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainersEnv", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ValueSource, actual.ValueSource, dcl.DiffInfo{ObjectFunction: compareServiceTemplateContainersEnvValueSourceNewStyle, EmptyObject: EmptyServiceTemplateContainersEnvValueSource, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("ValueSource")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateContainersEnvValueSourceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplateContainersEnvValueSource)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplateContainersEnvValueSource)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainersEnvValueSource or *ServiceTemplateContainersEnvValueSource", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplateContainersEnvValueSource)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplateContainersEnvValueSource)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainersEnvValueSource", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SecretKeyRef, actual.SecretKeyRef, dcl.DiffInfo{ObjectFunction: compareServiceTemplateContainersEnvValueSourceSecretKeyRefNewStyle, EmptyObject: EmptyServiceTemplateContainersEnvValueSourceSecretKeyRef, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("SecretKeyRef")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateContainersEnvValueSourceSecretKeyRefNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplateContainersEnvValueSourceSecretKeyRef)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplateContainersEnvValueSourceSecretKeyRef)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainersEnvValueSourceSecretKeyRef or *ServiceTemplateContainersEnvValueSourceSecretKeyRef", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplateContainersEnvValueSourceSecretKeyRef)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplateContainersEnvValueSourceSecretKeyRef)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainersEnvValueSourceSecretKeyRef", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Secret, actual.Secret, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Secret")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Version, actual.Version, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Version")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateContainersResourcesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplateContainersResources)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplateContainersResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainersResources or *ServiceTemplateContainersResources", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplateContainersResources)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplateContainersResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainersResources", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Limits, actual.Limits, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Limits")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CpuIdle, actual.CpuIdle, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("CpuIdle")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateContainersPortsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplateContainersPorts)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplateContainersPorts)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainersPorts or *ServiceTemplateContainersPorts", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplateContainersPorts)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplateContainersPorts)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainersPorts", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ContainerPort, actual.ContainerPort, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("ContainerPort")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateContainersVolumeMountsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplateContainersVolumeMounts)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplateContainersVolumeMounts)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainersVolumeMounts or *ServiceTemplateContainersVolumeMounts", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplateContainersVolumeMounts)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplateContainersVolumeMounts)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateContainersVolumeMounts", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MountPath, actual.MountPath, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("MountPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateVolumesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplateVolumes)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplateVolumes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateVolumes or *ServiceTemplateVolumes", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplateVolumes)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplateVolumes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateVolumes", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Secret, actual.Secret, dcl.DiffInfo{ObjectFunction: compareServiceTemplateVolumesSecretNewStyle, EmptyObject: EmptyServiceTemplateVolumesSecret, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Secret")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CloudSqlInstance, actual.CloudSqlInstance, dcl.DiffInfo{ObjectFunction: compareServiceTemplateVolumesCloudSqlInstanceNewStyle, EmptyObject: EmptyServiceTemplateVolumesCloudSqlInstance, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("CloudSqlInstance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateVolumesSecretNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplateVolumesSecret)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplateVolumesSecret)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateVolumesSecret or *ServiceTemplateVolumesSecret", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplateVolumesSecret)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplateVolumesSecret)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateVolumesSecret", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Secret, actual.Secret, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Secret")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Items, actual.Items, dcl.DiffInfo{ObjectFunction: compareServiceTemplateVolumesSecretItemsNewStyle, EmptyObject: EmptyServiceTemplateVolumesSecretItems, OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Items")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultMode, actual.DefaultMode, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("DefaultMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateVolumesSecretItemsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplateVolumesSecretItems)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplateVolumesSecretItems)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateVolumesSecretItems or *ServiceTemplateVolumesSecretItems", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplateVolumesSecretItems)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplateVolumesSecretItems)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateVolumesSecretItems", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Version, actual.Version, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Version")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Mode, actual.Mode, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Mode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTemplateVolumesCloudSqlInstanceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTemplateVolumesCloudSqlInstance)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTemplateVolumesCloudSqlInstance)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateVolumesCloudSqlInstance or *ServiceTemplateVolumesCloudSqlInstance", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTemplateVolumesCloudSqlInstance)
	if !ok {
		actualNotPointer, ok := a.(ServiceTemplateVolumesCloudSqlInstance)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTemplateVolumesCloudSqlInstance", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Instances, actual.Instances, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Instances")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTrafficNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTraffic)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTraffic)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTraffic or *ServiceTraffic", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTraffic)
	if !ok {
		actualNotPointer, ok := a.(ServiceTraffic)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTraffic", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Revision, actual.Revision, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Revision")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percent, actual.Percent, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Percent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Tag, actual.Tag, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Tag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTerminalConditionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTerminalCondition)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTerminalCondition)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTerminalCondition or *ServiceTerminalCondition", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTerminalCondition)
	if !ok {
		actualNotPointer, ok := a.(ServiceTerminalCondition)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTerminalCondition", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Message, actual.Message, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Message")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LastTransitionTime, actual.LastTransitionTime, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("LastTransitionTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Severity, actual.Severity, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Severity")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Reason, actual.Reason, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Reason")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RevisionReason, actual.RevisionReason, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("RevisionReason")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.JobReason, actual.JobReason, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("JobReason")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceTrafficStatusesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceTrafficStatuses)
	if !ok {
		desiredNotPointer, ok := d.(ServiceTrafficStatuses)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTrafficStatuses or *ServiceTrafficStatuses", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceTrafficStatuses)
	if !ok {
		actualNotPointer, ok := a.(ServiceTrafficStatuses)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceTrafficStatuses", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Revision, actual.Revision, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Revision")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Percent, actual.Percent, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Percent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Tag, actual.Tag, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Tag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Uri, actual.Uri, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateServiceUpdateServiceOperation")}, fn.AddNest("Uri")); len(ds) != 0 || err != nil {
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
func (r *Service) urlNormalized() *Service {
	normalized := dcl.Copy(*r).(Service)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Uid = dcl.SelfLinkToName(r.Uid)
	normalized.Creator = dcl.SelfLinkToName(r.Creator)
	normalized.LastModifier = dcl.SelfLinkToName(r.LastModifier)
	normalized.Client = dcl.SelfLinkToName(r.Client)
	normalized.ClientVersion = dcl.SelfLinkToName(r.ClientVersion)
	normalized.LatestReadyRevision = dcl.SelfLinkToName(r.LatestReadyRevision)
	normalized.LatestCreatedRevision = dcl.SelfLinkToName(r.LatestCreatedRevision)
	normalized.Uri = dcl.SelfLinkToName(r.Uri)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Service) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateService" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/services/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Service resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Service) marshal(c *Client) ([]byte, error) {
	m, err := expandService(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Service: %w", err)
	}
	m = EncodeServiceCreateRequest(m)

	return json.Marshal(m)
}

// unmarshalService decodes JSON responses into the Service resource schema.
func unmarshalService(b []byte, c *Client, res *Service) (*Service, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapService(m, c, res)
}

func unmarshalMapService(m map[string]interface{}, c *Client, res *Service) (*Service, error) {

	flattened := flattenService(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandService expands Service into a JSON request object.
func expandService(c *Client, f *Service) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("projects/%s/locations/%s/services/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Location), dcl.SelfLinkToName(f.Name)); err != nil {
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
	if v := f.Client; dcl.ValueShouldBeSent(v) {
		m["client"] = v
	}
	if v := f.ClientVersion; dcl.ValueShouldBeSent(v) {
		m["clientVersion"] = v
	}
	if v := f.Ingress; dcl.ValueShouldBeSent(v) {
		m["ingress"] = v
	}
	if v := f.LaunchStage; dcl.ValueShouldBeSent(v) {
		m["launchStage"] = v
	}
	if v, err := expandServiceBinaryAuthorization(c, f.BinaryAuthorization, res); err != nil {
		return nil, fmt.Errorf("error expanding BinaryAuthorization into binaryAuthorization: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["binaryAuthorization"] = v
	}
	if v, err := expandServiceTemplate(c, f.Template, res); err != nil {
		return nil, fmt.Errorf("error expanding Template into template: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["template"] = v
	}
	if v, err := expandServiceTrafficSlice(c, f.Traffic, res); err != nil {
		return nil, fmt.Errorf("error expanding Traffic into traffic: %w", err)
	} else if v != nil {
		m["traffic"] = v
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

// flattenService flattens Service from a JSON request object into the
// Service type.
func flattenService(c *Client, i interface{}, res *Service) *Service {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Service{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.Uid = dcl.FlattenString(m["uid"])
	resultRes.Generation = dcl.FlattenInteger(m["generation"])
	resultRes.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	resultRes.Annotations = dcl.FlattenKeyValuePairs(m["annotations"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.DeleteTime = dcl.FlattenString(m["deleteTime"])
	resultRes.ExpireTime = dcl.FlattenString(m["expireTime"])
	resultRes.Creator = dcl.FlattenString(m["creator"])
	resultRes.LastModifier = dcl.FlattenString(m["lastModifier"])
	resultRes.Client = dcl.FlattenString(m["client"])
	resultRes.ClientVersion = dcl.FlattenString(m["clientVersion"])
	resultRes.Ingress = flattenServiceIngressEnum(m["ingress"])
	resultRes.LaunchStage = flattenServiceLaunchStageEnum(m["launchStage"])
	resultRes.BinaryAuthorization = flattenServiceBinaryAuthorization(c, m["binaryAuthorization"], res)
	resultRes.Template = flattenServiceTemplate(c, m["template"], res)
	resultRes.Traffic = flattenServiceTrafficSlice(c, m["traffic"], res)
	resultRes.TerminalCondition = flattenServiceTerminalCondition(c, m["terminalCondition"], res)
	resultRes.LatestReadyRevision = dcl.FlattenString(m["latestReadyRevision"])
	resultRes.LatestCreatedRevision = dcl.FlattenString(m["latestCreatedRevision"])
	resultRes.TrafficStatuses = flattenServiceTrafficStatusesSlice(c, m["trafficStatuses"], res)
	resultRes.Uri = dcl.FlattenString(m["uri"])
	resultRes.Reconciling = dcl.FlattenBool(m["reconciling"])
	resultRes.Etag = dcl.FlattenString(m["etag"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// expandServiceBinaryAuthorizationMap expands the contents of ServiceBinaryAuthorization into a JSON
// request object.
func expandServiceBinaryAuthorizationMap(c *Client, f map[string]ServiceBinaryAuthorization, res *Service) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceBinaryAuthorization(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceBinaryAuthorizationSlice expands the contents of ServiceBinaryAuthorization into a JSON
// request object.
func expandServiceBinaryAuthorizationSlice(c *Client, f []ServiceBinaryAuthorization, res *Service) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceBinaryAuthorization(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceBinaryAuthorizationMap flattens the contents of ServiceBinaryAuthorization from a JSON
// response object.
func flattenServiceBinaryAuthorizationMap(c *Client, i interface{}, res *Service) map[string]ServiceBinaryAuthorization {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceBinaryAuthorization{}
	}

	if len(a) == 0 {
		return map[string]ServiceBinaryAuthorization{}
	}

	items := make(map[string]ServiceBinaryAuthorization)
	for k, item := range a {
		items[k] = *flattenServiceBinaryAuthorization(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenServiceBinaryAuthorizationSlice flattens the contents of ServiceBinaryAuthorization from a JSON
// response object.
func flattenServiceBinaryAuthorizationSlice(c *Client, i interface{}, res *Service) []ServiceBinaryAuthorization {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceBinaryAuthorization{}
	}

	if len(a) == 0 {
		return []ServiceBinaryAuthorization{}
	}

	items := make([]ServiceBinaryAuthorization, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceBinaryAuthorization(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandServiceBinaryAuthorization expands an instance of ServiceBinaryAuthorization into a JSON
// request object.
func expandServiceBinaryAuthorization(c *Client, f *ServiceBinaryAuthorization, res *Service) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.UseDefault; !dcl.IsEmptyValueIndirect(v) {
		m["useDefault"] = v
	}
	if v := f.BreakglassJustification; !dcl.IsEmptyValueIndirect(v) {
		m["breakglassJustification"] = v
	}

	return m, nil
}

// flattenServiceBinaryAuthorization flattens an instance of ServiceBinaryAuthorization from a JSON
// response object.
func flattenServiceBinaryAuthorization(c *Client, i interface{}, res *Service) *ServiceBinaryAuthorization {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceBinaryAuthorization{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceBinaryAuthorization
	}
	r.UseDefault = dcl.FlattenBool(m["useDefault"])
	r.BreakglassJustification = dcl.FlattenString(m["breakglassJustification"])

	return r
}

// expandServiceTemplateMap expands the contents of ServiceTemplate into a JSON
// request object.
func expandServiceTemplateMap(c *Client, f map[string]ServiceTemplate, res *Service) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplate(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateSlice expands the contents of ServiceTemplate into a JSON
// request object.
func expandServiceTemplateSlice(c *Client, f []ServiceTemplate, res *Service) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplate(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateMap flattens the contents of ServiceTemplate from a JSON
// response object.
func flattenServiceTemplateMap(c *Client, i interface{}, res *Service) map[string]ServiceTemplate {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplate{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplate{}
	}

	items := make(map[string]ServiceTemplate)
	for k, item := range a {
		items[k] = *flattenServiceTemplate(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenServiceTemplateSlice flattens the contents of ServiceTemplate from a JSON
// response object.
func flattenServiceTemplateSlice(c *Client, i interface{}, res *Service) []ServiceTemplate {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplate{}
	}

	if len(a) == 0 {
		return []ServiceTemplate{}
	}

	items := make([]ServiceTemplate, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplate(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandServiceTemplate expands an instance of ServiceTemplate into a JSON
// request object.
func expandServiceTemplate(c *Client, f *ServiceTemplate, res *Service) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Revision; !dcl.IsEmptyValueIndirect(v) {
		m["revision"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.Annotations; !dcl.IsEmptyValueIndirect(v) {
		m["annotations"] = v
	}
	if v, err := expandServiceTemplateScaling(c, f.Scaling, res); err != nil {
		return nil, fmt.Errorf("error expanding Scaling into scaling: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["scaling"] = v
	}
	if v, err := expandServiceTemplateVPCAccess(c, f.VPCAccess, res); err != nil {
		return nil, fmt.Errorf("error expanding VPCAccess into vpcAccess: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["vpcAccess"] = v
	}
	if v := f.ContainerConcurrency; !dcl.IsEmptyValueIndirect(v) {
		m["maxInstanceRequestConcurrency"] = v
	}
	if v := f.Timeout; !dcl.IsEmptyValueIndirect(v) {
		m["timeout"] = v
	}
	if v := f.ServiceAccount; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccount"] = v
	}
	if v, err := expandServiceTemplateContainersSlice(c, f.Containers, res); err != nil {
		return nil, fmt.Errorf("error expanding Containers into containers: %w", err)
	} else if v != nil {
		m["containers"] = v
	}
	if v, err := expandServiceTemplateVolumesSlice(c, f.Volumes, res); err != nil {
		return nil, fmt.Errorf("error expanding Volumes into volumes: %w", err)
	} else if v != nil {
		m["volumes"] = v
	}
	if v := f.ExecutionEnvironment; !dcl.IsEmptyValueIndirect(v) {
		m["executionEnvironment"] = v
	}

	return m, nil
}

// flattenServiceTemplate flattens an instance of ServiceTemplate from a JSON
// response object.
func flattenServiceTemplate(c *Client, i interface{}, res *Service) *ServiceTemplate {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplate{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplate
	}
	r.Revision = dcl.FlattenString(m["revision"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.Annotations = dcl.FlattenKeyValuePairs(m["annotations"])
	r.Scaling = flattenServiceTemplateScaling(c, m["scaling"], res)
	r.VPCAccess = flattenServiceTemplateVPCAccess(c, m["vpcAccess"], res)
	r.ContainerConcurrency = dcl.FlattenInteger(m["maxInstanceRequestConcurrency"])
	r.Timeout = dcl.FlattenString(m["timeout"])
	r.ServiceAccount = dcl.FlattenString(m["serviceAccount"])
	r.Containers = flattenServiceTemplateContainersSlice(c, m["containers"], res)
	r.Volumes = flattenServiceTemplateVolumesSlice(c, m["volumes"], res)
	r.ExecutionEnvironment = flattenServiceTemplateExecutionEnvironmentEnum(m["executionEnvironment"])

	return r
}

// expandServiceTemplateScalingMap expands the contents of ServiceTemplateScaling into a JSON
// request object.
func expandServiceTemplateScalingMap(c *Client, f map[string]ServiceTemplateScaling, res *Service) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplateScaling(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateScalingSlice expands the contents of ServiceTemplateScaling into a JSON
// request object.
func expandServiceTemplateScalingSlice(c *Client, f []ServiceTemplateScaling, res *Service) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplateScaling(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateScalingMap flattens the contents of ServiceTemplateScaling from a JSON
// response object.
func flattenServiceTemplateScalingMap(c *Client, i interface{}, res *Service) map[string]ServiceTemplateScaling {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateScaling{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateScaling{}
	}

	items := make(map[string]ServiceTemplateScaling)
	for k, item := range a {
		items[k] = *flattenServiceTemplateScaling(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenServiceTemplateScalingSlice flattens the contents of ServiceTemplateScaling from a JSON
// response object.
func flattenServiceTemplateScalingSlice(c *Client, i interface{}, res *Service) []ServiceTemplateScaling {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateScaling{}
	}

	if len(a) == 0 {
		return []ServiceTemplateScaling{}
	}

	items := make([]ServiceTemplateScaling, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateScaling(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandServiceTemplateScaling expands an instance of ServiceTemplateScaling into a JSON
// request object.
func expandServiceTemplateScaling(c *Client, f *ServiceTemplateScaling, res *Service) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MinInstanceCount; !dcl.IsEmptyValueIndirect(v) {
		m["minInstanceCount"] = v
	}
	if v := f.MaxInstanceCount; !dcl.IsEmptyValueIndirect(v) {
		m["maxInstanceCount"] = v
	}

	return m, nil
}

// flattenServiceTemplateScaling flattens an instance of ServiceTemplateScaling from a JSON
// response object.
func flattenServiceTemplateScaling(c *Client, i interface{}, res *Service) *ServiceTemplateScaling {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplateScaling{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplateScaling
	}
	r.MinInstanceCount = dcl.FlattenInteger(m["minInstanceCount"])
	r.MaxInstanceCount = dcl.FlattenInteger(m["maxInstanceCount"])

	return r
}

// expandServiceTemplateVPCAccessMap expands the contents of ServiceTemplateVPCAccess into a JSON
// request object.
func expandServiceTemplateVPCAccessMap(c *Client, f map[string]ServiceTemplateVPCAccess, res *Service) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplateVPCAccess(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateVPCAccessSlice expands the contents of ServiceTemplateVPCAccess into a JSON
// request object.
func expandServiceTemplateVPCAccessSlice(c *Client, f []ServiceTemplateVPCAccess, res *Service) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplateVPCAccess(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateVPCAccessMap flattens the contents of ServiceTemplateVPCAccess from a JSON
// response object.
func flattenServiceTemplateVPCAccessMap(c *Client, i interface{}, res *Service) map[string]ServiceTemplateVPCAccess {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateVPCAccess{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateVPCAccess{}
	}

	items := make(map[string]ServiceTemplateVPCAccess)
	for k, item := range a {
		items[k] = *flattenServiceTemplateVPCAccess(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenServiceTemplateVPCAccessSlice flattens the contents of ServiceTemplateVPCAccess from a JSON
// response object.
func flattenServiceTemplateVPCAccessSlice(c *Client, i interface{}, res *Service) []ServiceTemplateVPCAccess {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateVPCAccess{}
	}

	if len(a) == 0 {
		return []ServiceTemplateVPCAccess{}
	}

	items := make([]ServiceTemplateVPCAccess, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateVPCAccess(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandServiceTemplateVPCAccess expands an instance of ServiceTemplateVPCAccess into a JSON
// request object.
func expandServiceTemplateVPCAccess(c *Client, f *ServiceTemplateVPCAccess, res *Service) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Connector; !dcl.IsEmptyValueIndirect(v) {
		m["connector"] = v
	}
	if v := f.Egress; !dcl.IsEmptyValueIndirect(v) {
		m["egress"] = v
	}

	return m, nil
}

// flattenServiceTemplateVPCAccess flattens an instance of ServiceTemplateVPCAccess from a JSON
// response object.
func flattenServiceTemplateVPCAccess(c *Client, i interface{}, res *Service) *ServiceTemplateVPCAccess {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplateVPCAccess{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplateVPCAccess
	}
	r.Connector = dcl.FlattenString(m["connector"])
	r.Egress = flattenServiceTemplateVPCAccessEgressEnum(m["egress"])

	return r
}

// expandServiceTemplateContainersMap expands the contents of ServiceTemplateContainers into a JSON
// request object.
func expandServiceTemplateContainersMap(c *Client, f map[string]ServiceTemplateContainers, res *Service) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplateContainers(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateContainersSlice expands the contents of ServiceTemplateContainers into a JSON
// request object.
func expandServiceTemplateContainersSlice(c *Client, f []ServiceTemplateContainers, res *Service) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplateContainers(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateContainersMap flattens the contents of ServiceTemplateContainers from a JSON
// response object.
func flattenServiceTemplateContainersMap(c *Client, i interface{}, res *Service) map[string]ServiceTemplateContainers {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateContainers{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateContainers{}
	}

	items := make(map[string]ServiceTemplateContainers)
	for k, item := range a {
		items[k] = *flattenServiceTemplateContainers(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenServiceTemplateContainersSlice flattens the contents of ServiceTemplateContainers from a JSON
// response object.
func flattenServiceTemplateContainersSlice(c *Client, i interface{}, res *Service) []ServiceTemplateContainers {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateContainers{}
	}

	if len(a) == 0 {
		return []ServiceTemplateContainers{}
	}

	items := make([]ServiceTemplateContainers, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateContainers(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandServiceTemplateContainers expands an instance of ServiceTemplateContainers into a JSON
// request object.
func expandServiceTemplateContainers(c *Client, f *ServiceTemplateContainers, res *Service) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Image; !dcl.IsEmptyValueIndirect(v) {
		m["image"] = v
	}
	if v := f.Command; v != nil {
		m["command"] = v
	}
	if v := f.Args; v != nil {
		m["args"] = v
	}
	if v, err := expandServiceTemplateContainersEnvSlice(c, f.Env, res); err != nil {
		return nil, fmt.Errorf("error expanding Env into env: %w", err)
	} else if v != nil {
		m["env"] = v
	}
	if v, err := expandServiceTemplateContainersResources(c, f.Resources, res); err != nil {
		return nil, fmt.Errorf("error expanding Resources into resources: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["resources"] = v
	}
	if v, err := expandServiceTemplateContainersPortsSlice(c, f.Ports, res); err != nil {
		return nil, fmt.Errorf("error expanding Ports into ports: %w", err)
	} else if v != nil {
		m["ports"] = v
	}
	if v, err := expandServiceTemplateContainersVolumeMountsSlice(c, f.VolumeMounts, res); err != nil {
		return nil, fmt.Errorf("error expanding VolumeMounts into volumeMounts: %w", err)
	} else if v != nil {
		m["volumeMounts"] = v
	}

	return m, nil
}

// flattenServiceTemplateContainers flattens an instance of ServiceTemplateContainers from a JSON
// response object.
func flattenServiceTemplateContainers(c *Client, i interface{}, res *Service) *ServiceTemplateContainers {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplateContainers{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplateContainers
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Image = dcl.FlattenString(m["image"])
	r.Command = dcl.FlattenStringSlice(m["command"])
	r.Args = dcl.FlattenStringSlice(m["args"])
	r.Env = flattenServiceTemplateContainersEnvSlice(c, m["env"], res)
	r.Resources = flattenServiceTemplateContainersResources(c, m["resources"], res)
	r.Ports = flattenServiceTemplateContainersPortsSlice(c, m["ports"], res)
	r.VolumeMounts = flattenServiceTemplateContainersVolumeMountsSlice(c, m["volumeMounts"], res)

	return r
}

// expandServiceTemplateContainersEnvMap expands the contents of ServiceTemplateContainersEnv into a JSON
// request object.
func expandServiceTemplateContainersEnvMap(c *Client, f map[string]ServiceTemplateContainersEnv, res *Service) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplateContainersEnv(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateContainersEnvSlice expands the contents of ServiceTemplateContainersEnv into a JSON
// request object.
func expandServiceTemplateContainersEnvSlice(c *Client, f []ServiceTemplateContainersEnv, res *Service) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplateContainersEnv(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateContainersEnvMap flattens the contents of ServiceTemplateContainersEnv from a JSON
// response object.
func flattenServiceTemplateContainersEnvMap(c *Client, i interface{}, res *Service) map[string]ServiceTemplateContainersEnv {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateContainersEnv{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateContainersEnv{}
	}

	items := make(map[string]ServiceTemplateContainersEnv)
	for k, item := range a {
		items[k] = *flattenServiceTemplateContainersEnv(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenServiceTemplateContainersEnvSlice flattens the contents of ServiceTemplateContainersEnv from a JSON
// response object.
func flattenServiceTemplateContainersEnvSlice(c *Client, i interface{}, res *Service) []ServiceTemplateContainersEnv {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateContainersEnv{}
	}

	if len(a) == 0 {
		return []ServiceTemplateContainersEnv{}
	}

	items := make([]ServiceTemplateContainersEnv, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateContainersEnv(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandServiceTemplateContainersEnv expands an instance of ServiceTemplateContainersEnv into a JSON
// request object.
func expandServiceTemplateContainersEnv(c *Client, f *ServiceTemplateContainersEnv, res *Service) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}
	if v, err := expandServiceTemplateContainersEnvValueSource(c, f.ValueSource, res); err != nil {
		return nil, fmt.Errorf("error expanding ValueSource into valueSource: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["valueSource"] = v
	}

	return m, nil
}

// flattenServiceTemplateContainersEnv flattens an instance of ServiceTemplateContainersEnv from a JSON
// response object.
func flattenServiceTemplateContainersEnv(c *Client, i interface{}, res *Service) *ServiceTemplateContainersEnv {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplateContainersEnv{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplateContainersEnv
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Value = dcl.FlattenString(m["value"])
	r.ValueSource = flattenServiceTemplateContainersEnvValueSource(c, m["valueSource"], res)

	return r
}

// expandServiceTemplateContainersEnvValueSourceMap expands the contents of ServiceTemplateContainersEnvValueSource into a JSON
// request object.
func expandServiceTemplateContainersEnvValueSourceMap(c *Client, f map[string]ServiceTemplateContainersEnvValueSource, res *Service) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplateContainersEnvValueSource(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateContainersEnvValueSourceSlice expands the contents of ServiceTemplateContainersEnvValueSource into a JSON
// request object.
func expandServiceTemplateContainersEnvValueSourceSlice(c *Client, f []ServiceTemplateContainersEnvValueSource, res *Service) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplateContainersEnvValueSource(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateContainersEnvValueSourceMap flattens the contents of ServiceTemplateContainersEnvValueSource from a JSON
// response object.
func flattenServiceTemplateContainersEnvValueSourceMap(c *Client, i interface{}, res *Service) map[string]ServiceTemplateContainersEnvValueSource {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateContainersEnvValueSource{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateContainersEnvValueSource{}
	}

	items := make(map[string]ServiceTemplateContainersEnvValueSource)
	for k, item := range a {
		items[k] = *flattenServiceTemplateContainersEnvValueSource(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenServiceTemplateContainersEnvValueSourceSlice flattens the contents of ServiceTemplateContainersEnvValueSource from a JSON
// response object.
func flattenServiceTemplateContainersEnvValueSourceSlice(c *Client, i interface{}, res *Service) []ServiceTemplateContainersEnvValueSource {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateContainersEnvValueSource{}
	}

	if len(a) == 0 {
		return []ServiceTemplateContainersEnvValueSource{}
	}

	items := make([]ServiceTemplateContainersEnvValueSource, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateContainersEnvValueSource(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandServiceTemplateContainersEnvValueSource expands an instance of ServiceTemplateContainersEnvValueSource into a JSON
// request object.
func expandServiceTemplateContainersEnvValueSource(c *Client, f *ServiceTemplateContainersEnvValueSource, res *Service) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandServiceTemplateContainersEnvValueSourceSecretKeyRef(c, f.SecretKeyRef, res); err != nil {
		return nil, fmt.Errorf("error expanding SecretKeyRef into secretKeyRef: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["secretKeyRef"] = v
	}

	return m, nil
}

// flattenServiceTemplateContainersEnvValueSource flattens an instance of ServiceTemplateContainersEnvValueSource from a JSON
// response object.
func flattenServiceTemplateContainersEnvValueSource(c *Client, i interface{}, res *Service) *ServiceTemplateContainersEnvValueSource {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplateContainersEnvValueSource{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplateContainersEnvValueSource
	}
	r.SecretKeyRef = flattenServiceTemplateContainersEnvValueSourceSecretKeyRef(c, m["secretKeyRef"], res)

	return r
}

// expandServiceTemplateContainersEnvValueSourceSecretKeyRefMap expands the contents of ServiceTemplateContainersEnvValueSourceSecretKeyRef into a JSON
// request object.
func expandServiceTemplateContainersEnvValueSourceSecretKeyRefMap(c *Client, f map[string]ServiceTemplateContainersEnvValueSourceSecretKeyRef, res *Service) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplateContainersEnvValueSourceSecretKeyRef(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateContainersEnvValueSourceSecretKeyRefSlice expands the contents of ServiceTemplateContainersEnvValueSourceSecretKeyRef into a JSON
// request object.
func expandServiceTemplateContainersEnvValueSourceSecretKeyRefSlice(c *Client, f []ServiceTemplateContainersEnvValueSourceSecretKeyRef, res *Service) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplateContainersEnvValueSourceSecretKeyRef(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateContainersEnvValueSourceSecretKeyRefMap flattens the contents of ServiceTemplateContainersEnvValueSourceSecretKeyRef from a JSON
// response object.
func flattenServiceTemplateContainersEnvValueSourceSecretKeyRefMap(c *Client, i interface{}, res *Service) map[string]ServiceTemplateContainersEnvValueSourceSecretKeyRef {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateContainersEnvValueSourceSecretKeyRef{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateContainersEnvValueSourceSecretKeyRef{}
	}

	items := make(map[string]ServiceTemplateContainersEnvValueSourceSecretKeyRef)
	for k, item := range a {
		items[k] = *flattenServiceTemplateContainersEnvValueSourceSecretKeyRef(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenServiceTemplateContainersEnvValueSourceSecretKeyRefSlice flattens the contents of ServiceTemplateContainersEnvValueSourceSecretKeyRef from a JSON
// response object.
func flattenServiceTemplateContainersEnvValueSourceSecretKeyRefSlice(c *Client, i interface{}, res *Service) []ServiceTemplateContainersEnvValueSourceSecretKeyRef {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateContainersEnvValueSourceSecretKeyRef{}
	}

	if len(a) == 0 {
		return []ServiceTemplateContainersEnvValueSourceSecretKeyRef{}
	}

	items := make([]ServiceTemplateContainersEnvValueSourceSecretKeyRef, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateContainersEnvValueSourceSecretKeyRef(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandServiceTemplateContainersEnvValueSourceSecretKeyRef expands an instance of ServiceTemplateContainersEnvValueSourceSecretKeyRef into a JSON
// request object.
func expandServiceTemplateContainersEnvValueSourceSecretKeyRef(c *Client, f *ServiceTemplateContainersEnvValueSourceSecretKeyRef, res *Service) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := dcl.SelfLinkToNameExpander(f.Secret); err != nil {
		return nil, fmt.Errorf("error expanding Secret into secret: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["secret"] = v
	}
	if v, err := dcl.SelfLinkToNameExpander(f.Version); err != nil {
		return nil, fmt.Errorf("error expanding Version into version: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}

	return m, nil
}

// flattenServiceTemplateContainersEnvValueSourceSecretKeyRef flattens an instance of ServiceTemplateContainersEnvValueSourceSecretKeyRef from a JSON
// response object.
func flattenServiceTemplateContainersEnvValueSourceSecretKeyRef(c *Client, i interface{}, res *Service) *ServiceTemplateContainersEnvValueSourceSecretKeyRef {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplateContainersEnvValueSourceSecretKeyRef{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplateContainersEnvValueSourceSecretKeyRef
	}
	r.Secret = dcl.FlattenString(m["secret"])
	r.Version = dcl.FlattenString(m["version"])

	return r
}

// expandServiceTemplateContainersResourcesMap expands the contents of ServiceTemplateContainersResources into a JSON
// request object.
func expandServiceTemplateContainersResourcesMap(c *Client, f map[string]ServiceTemplateContainersResources, res *Service) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplateContainersResources(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateContainersResourcesSlice expands the contents of ServiceTemplateContainersResources into a JSON
// request object.
func expandServiceTemplateContainersResourcesSlice(c *Client, f []ServiceTemplateContainersResources, res *Service) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplateContainersResources(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateContainersResourcesMap flattens the contents of ServiceTemplateContainersResources from a JSON
// response object.
func flattenServiceTemplateContainersResourcesMap(c *Client, i interface{}, res *Service) map[string]ServiceTemplateContainersResources {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateContainersResources{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateContainersResources{}
	}

	items := make(map[string]ServiceTemplateContainersResources)
	for k, item := range a {
		items[k] = *flattenServiceTemplateContainersResources(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenServiceTemplateContainersResourcesSlice flattens the contents of ServiceTemplateContainersResources from a JSON
// response object.
func flattenServiceTemplateContainersResourcesSlice(c *Client, i interface{}, res *Service) []ServiceTemplateContainersResources {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateContainersResources{}
	}

	if len(a) == 0 {
		return []ServiceTemplateContainersResources{}
	}

	items := make([]ServiceTemplateContainersResources, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateContainersResources(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandServiceTemplateContainersResources expands an instance of ServiceTemplateContainersResources into a JSON
// request object.
func expandServiceTemplateContainersResources(c *Client, f *ServiceTemplateContainersResources, res *Service) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Limits; !dcl.IsEmptyValueIndirect(v) {
		m["limits"] = v
	}
	if v := f.CpuIdle; !dcl.IsEmptyValueIndirect(v) {
		m["cpuIdle"] = v
	}

	return m, nil
}

// flattenServiceTemplateContainersResources flattens an instance of ServiceTemplateContainersResources from a JSON
// response object.
func flattenServiceTemplateContainersResources(c *Client, i interface{}, res *Service) *ServiceTemplateContainersResources {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplateContainersResources{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplateContainersResources
	}
	r.Limits = dcl.FlattenKeyValuePairs(m["limits"])
	r.CpuIdle = dcl.FlattenBool(m["cpuIdle"])

	return r
}

// expandServiceTemplateContainersPortsMap expands the contents of ServiceTemplateContainersPorts into a JSON
// request object.
func expandServiceTemplateContainersPortsMap(c *Client, f map[string]ServiceTemplateContainersPorts, res *Service) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplateContainersPorts(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateContainersPortsSlice expands the contents of ServiceTemplateContainersPorts into a JSON
// request object.
func expandServiceTemplateContainersPortsSlice(c *Client, f []ServiceTemplateContainersPorts, res *Service) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplateContainersPorts(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateContainersPortsMap flattens the contents of ServiceTemplateContainersPorts from a JSON
// response object.
func flattenServiceTemplateContainersPortsMap(c *Client, i interface{}, res *Service) map[string]ServiceTemplateContainersPorts {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateContainersPorts{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateContainersPorts{}
	}

	items := make(map[string]ServiceTemplateContainersPorts)
	for k, item := range a {
		items[k] = *flattenServiceTemplateContainersPorts(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenServiceTemplateContainersPortsSlice flattens the contents of ServiceTemplateContainersPorts from a JSON
// response object.
func flattenServiceTemplateContainersPortsSlice(c *Client, i interface{}, res *Service) []ServiceTemplateContainersPorts {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateContainersPorts{}
	}

	if len(a) == 0 {
		return []ServiceTemplateContainersPorts{}
	}

	items := make([]ServiceTemplateContainersPorts, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateContainersPorts(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandServiceTemplateContainersPorts expands an instance of ServiceTemplateContainersPorts into a JSON
// request object.
func expandServiceTemplateContainersPorts(c *Client, f *ServiceTemplateContainersPorts, res *Service) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.ContainerPort; !dcl.IsEmptyValueIndirect(v) {
		m["containerPort"] = v
	}

	return m, nil
}

// flattenServiceTemplateContainersPorts flattens an instance of ServiceTemplateContainersPorts from a JSON
// response object.
func flattenServiceTemplateContainersPorts(c *Client, i interface{}, res *Service) *ServiceTemplateContainersPorts {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplateContainersPorts{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplateContainersPorts
	}
	r.Name = dcl.FlattenString(m["name"])
	r.ContainerPort = dcl.FlattenInteger(m["containerPort"])

	return r
}

// expandServiceTemplateContainersVolumeMountsMap expands the contents of ServiceTemplateContainersVolumeMounts into a JSON
// request object.
func expandServiceTemplateContainersVolumeMountsMap(c *Client, f map[string]ServiceTemplateContainersVolumeMounts, res *Service) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplateContainersVolumeMounts(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateContainersVolumeMountsSlice expands the contents of ServiceTemplateContainersVolumeMounts into a JSON
// request object.
func expandServiceTemplateContainersVolumeMountsSlice(c *Client, f []ServiceTemplateContainersVolumeMounts, res *Service) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplateContainersVolumeMounts(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateContainersVolumeMountsMap flattens the contents of ServiceTemplateContainersVolumeMounts from a JSON
// response object.
func flattenServiceTemplateContainersVolumeMountsMap(c *Client, i interface{}, res *Service) map[string]ServiceTemplateContainersVolumeMounts {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateContainersVolumeMounts{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateContainersVolumeMounts{}
	}

	items := make(map[string]ServiceTemplateContainersVolumeMounts)
	for k, item := range a {
		items[k] = *flattenServiceTemplateContainersVolumeMounts(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenServiceTemplateContainersVolumeMountsSlice flattens the contents of ServiceTemplateContainersVolumeMounts from a JSON
// response object.
func flattenServiceTemplateContainersVolumeMountsSlice(c *Client, i interface{}, res *Service) []ServiceTemplateContainersVolumeMounts {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateContainersVolumeMounts{}
	}

	if len(a) == 0 {
		return []ServiceTemplateContainersVolumeMounts{}
	}

	items := make([]ServiceTemplateContainersVolumeMounts, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateContainersVolumeMounts(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandServiceTemplateContainersVolumeMounts expands an instance of ServiceTemplateContainersVolumeMounts into a JSON
// request object.
func expandServiceTemplateContainersVolumeMounts(c *Client, f *ServiceTemplateContainersVolumeMounts, res *Service) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.MountPath; !dcl.IsEmptyValueIndirect(v) {
		m["mountPath"] = v
	}

	return m, nil
}

// flattenServiceTemplateContainersVolumeMounts flattens an instance of ServiceTemplateContainersVolumeMounts from a JSON
// response object.
func flattenServiceTemplateContainersVolumeMounts(c *Client, i interface{}, res *Service) *ServiceTemplateContainersVolumeMounts {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplateContainersVolumeMounts{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplateContainersVolumeMounts
	}
	r.Name = dcl.FlattenString(m["name"])
	r.MountPath = dcl.FlattenString(m["mountPath"])

	return r
}

// expandServiceTemplateVolumesMap expands the contents of ServiceTemplateVolumes into a JSON
// request object.
func expandServiceTemplateVolumesMap(c *Client, f map[string]ServiceTemplateVolumes, res *Service) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplateVolumes(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateVolumesSlice expands the contents of ServiceTemplateVolumes into a JSON
// request object.
func expandServiceTemplateVolumesSlice(c *Client, f []ServiceTemplateVolumes, res *Service) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplateVolumes(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateVolumesMap flattens the contents of ServiceTemplateVolumes from a JSON
// response object.
func flattenServiceTemplateVolumesMap(c *Client, i interface{}, res *Service) map[string]ServiceTemplateVolumes {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateVolumes{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateVolumes{}
	}

	items := make(map[string]ServiceTemplateVolumes)
	for k, item := range a {
		items[k] = *flattenServiceTemplateVolumes(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenServiceTemplateVolumesSlice flattens the contents of ServiceTemplateVolumes from a JSON
// response object.
func flattenServiceTemplateVolumesSlice(c *Client, i interface{}, res *Service) []ServiceTemplateVolumes {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateVolumes{}
	}

	if len(a) == 0 {
		return []ServiceTemplateVolumes{}
	}

	items := make([]ServiceTemplateVolumes, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateVolumes(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandServiceTemplateVolumes expands an instance of ServiceTemplateVolumes into a JSON
// request object.
func expandServiceTemplateVolumes(c *Client, f *ServiceTemplateVolumes, res *Service) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandServiceTemplateVolumesSecret(c, f.Secret, res); err != nil {
		return nil, fmt.Errorf("error expanding Secret into secret: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["secret"] = v
	}
	if v, err := expandServiceTemplateVolumesCloudSqlInstance(c, f.CloudSqlInstance, res); err != nil {
		return nil, fmt.Errorf("error expanding CloudSqlInstance into cloudSqlInstance: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cloudSqlInstance"] = v
	}

	return m, nil
}

// flattenServiceTemplateVolumes flattens an instance of ServiceTemplateVolumes from a JSON
// response object.
func flattenServiceTemplateVolumes(c *Client, i interface{}, res *Service) *ServiceTemplateVolumes {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplateVolumes{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplateVolumes
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Secret = flattenServiceTemplateVolumesSecret(c, m["secret"], res)
	r.CloudSqlInstance = flattenServiceTemplateVolumesCloudSqlInstance(c, m["cloudSqlInstance"], res)

	return r
}

// expandServiceTemplateVolumesSecretMap expands the contents of ServiceTemplateVolumesSecret into a JSON
// request object.
func expandServiceTemplateVolumesSecretMap(c *Client, f map[string]ServiceTemplateVolumesSecret, res *Service) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplateVolumesSecret(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateVolumesSecretSlice expands the contents of ServiceTemplateVolumesSecret into a JSON
// request object.
func expandServiceTemplateVolumesSecretSlice(c *Client, f []ServiceTemplateVolumesSecret, res *Service) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplateVolumesSecret(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateVolumesSecretMap flattens the contents of ServiceTemplateVolumesSecret from a JSON
// response object.
func flattenServiceTemplateVolumesSecretMap(c *Client, i interface{}, res *Service) map[string]ServiceTemplateVolumesSecret {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateVolumesSecret{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateVolumesSecret{}
	}

	items := make(map[string]ServiceTemplateVolumesSecret)
	for k, item := range a {
		items[k] = *flattenServiceTemplateVolumesSecret(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenServiceTemplateVolumesSecretSlice flattens the contents of ServiceTemplateVolumesSecret from a JSON
// response object.
func flattenServiceTemplateVolumesSecretSlice(c *Client, i interface{}, res *Service) []ServiceTemplateVolumesSecret {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateVolumesSecret{}
	}

	if len(a) == 0 {
		return []ServiceTemplateVolumesSecret{}
	}

	items := make([]ServiceTemplateVolumesSecret, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateVolumesSecret(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandServiceTemplateVolumesSecret expands an instance of ServiceTemplateVolumesSecret into a JSON
// request object.
func expandServiceTemplateVolumesSecret(c *Client, f *ServiceTemplateVolumesSecret, res *Service) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Secret; !dcl.IsEmptyValueIndirect(v) {
		m["secret"] = v
	}
	if v, err := expandServiceTemplateVolumesSecretItemsSlice(c, f.Items, res); err != nil {
		return nil, fmt.Errorf("error expanding Items into items: %w", err)
	} else if v != nil {
		m["items"] = v
	}
	if v := f.DefaultMode; !dcl.IsEmptyValueIndirect(v) {
		m["defaultMode"] = v
	}

	return m, nil
}

// flattenServiceTemplateVolumesSecret flattens an instance of ServiceTemplateVolumesSecret from a JSON
// response object.
func flattenServiceTemplateVolumesSecret(c *Client, i interface{}, res *Service) *ServiceTemplateVolumesSecret {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplateVolumesSecret{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplateVolumesSecret
	}
	r.Secret = dcl.FlattenString(m["secret"])
	r.Items = flattenServiceTemplateVolumesSecretItemsSlice(c, m["items"], res)
	r.DefaultMode = dcl.FlattenInteger(m["defaultMode"])

	return r
}

// expandServiceTemplateVolumesSecretItemsMap expands the contents of ServiceTemplateVolumesSecretItems into a JSON
// request object.
func expandServiceTemplateVolumesSecretItemsMap(c *Client, f map[string]ServiceTemplateVolumesSecretItems, res *Service) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplateVolumesSecretItems(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateVolumesSecretItemsSlice expands the contents of ServiceTemplateVolumesSecretItems into a JSON
// request object.
func expandServiceTemplateVolumesSecretItemsSlice(c *Client, f []ServiceTemplateVolumesSecretItems, res *Service) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplateVolumesSecretItems(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateVolumesSecretItemsMap flattens the contents of ServiceTemplateVolumesSecretItems from a JSON
// response object.
func flattenServiceTemplateVolumesSecretItemsMap(c *Client, i interface{}, res *Service) map[string]ServiceTemplateVolumesSecretItems {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateVolumesSecretItems{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateVolumesSecretItems{}
	}

	items := make(map[string]ServiceTemplateVolumesSecretItems)
	for k, item := range a {
		items[k] = *flattenServiceTemplateVolumesSecretItems(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenServiceTemplateVolumesSecretItemsSlice flattens the contents of ServiceTemplateVolumesSecretItems from a JSON
// response object.
func flattenServiceTemplateVolumesSecretItemsSlice(c *Client, i interface{}, res *Service) []ServiceTemplateVolumesSecretItems {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateVolumesSecretItems{}
	}

	if len(a) == 0 {
		return []ServiceTemplateVolumesSecretItems{}
	}

	items := make([]ServiceTemplateVolumesSecretItems, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateVolumesSecretItems(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandServiceTemplateVolumesSecretItems expands an instance of ServiceTemplateVolumesSecretItems into a JSON
// request object.
func expandServiceTemplateVolumesSecretItems(c *Client, f *ServiceTemplateVolumesSecretItems, res *Service) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Path; !dcl.IsEmptyValueIndirect(v) {
		m["path"] = v
	}
	if v, err := dcl.SelfLinkToNameExpander(f.Version); err != nil {
		return nil, fmt.Errorf("error expanding Version into version: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}
	if v := f.Mode; !dcl.IsEmptyValueIndirect(v) {
		m["mode"] = v
	}

	return m, nil
}

// flattenServiceTemplateVolumesSecretItems flattens an instance of ServiceTemplateVolumesSecretItems from a JSON
// response object.
func flattenServiceTemplateVolumesSecretItems(c *Client, i interface{}, res *Service) *ServiceTemplateVolumesSecretItems {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplateVolumesSecretItems{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplateVolumesSecretItems
	}
	r.Path = dcl.FlattenString(m["path"])
	r.Version = dcl.FlattenString(m["version"])
	r.Mode = dcl.FlattenInteger(m["mode"])

	return r
}

// expandServiceTemplateVolumesCloudSqlInstanceMap expands the contents of ServiceTemplateVolumesCloudSqlInstance into a JSON
// request object.
func expandServiceTemplateVolumesCloudSqlInstanceMap(c *Client, f map[string]ServiceTemplateVolumesCloudSqlInstance, res *Service) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTemplateVolumesCloudSqlInstance(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTemplateVolumesCloudSqlInstanceSlice expands the contents of ServiceTemplateVolumesCloudSqlInstance into a JSON
// request object.
func expandServiceTemplateVolumesCloudSqlInstanceSlice(c *Client, f []ServiceTemplateVolumesCloudSqlInstance, res *Service) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTemplateVolumesCloudSqlInstance(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTemplateVolumesCloudSqlInstanceMap flattens the contents of ServiceTemplateVolumesCloudSqlInstance from a JSON
// response object.
func flattenServiceTemplateVolumesCloudSqlInstanceMap(c *Client, i interface{}, res *Service) map[string]ServiceTemplateVolumesCloudSqlInstance {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateVolumesCloudSqlInstance{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateVolumesCloudSqlInstance{}
	}

	items := make(map[string]ServiceTemplateVolumesCloudSqlInstance)
	for k, item := range a {
		items[k] = *flattenServiceTemplateVolumesCloudSqlInstance(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenServiceTemplateVolumesCloudSqlInstanceSlice flattens the contents of ServiceTemplateVolumesCloudSqlInstance from a JSON
// response object.
func flattenServiceTemplateVolumesCloudSqlInstanceSlice(c *Client, i interface{}, res *Service) []ServiceTemplateVolumesCloudSqlInstance {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateVolumesCloudSqlInstance{}
	}

	if len(a) == 0 {
		return []ServiceTemplateVolumesCloudSqlInstance{}
	}

	items := make([]ServiceTemplateVolumesCloudSqlInstance, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateVolumesCloudSqlInstance(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandServiceTemplateVolumesCloudSqlInstance expands an instance of ServiceTemplateVolumesCloudSqlInstance into a JSON
// request object.
func expandServiceTemplateVolumesCloudSqlInstance(c *Client, f *ServiceTemplateVolumesCloudSqlInstance, res *Service) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Instances; v != nil {
		m["instances"] = v
	}

	return m, nil
}

// flattenServiceTemplateVolumesCloudSqlInstance flattens an instance of ServiceTemplateVolumesCloudSqlInstance from a JSON
// response object.
func flattenServiceTemplateVolumesCloudSqlInstance(c *Client, i interface{}, res *Service) *ServiceTemplateVolumesCloudSqlInstance {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTemplateVolumesCloudSqlInstance{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTemplateVolumesCloudSqlInstance
	}
	r.Instances = dcl.FlattenStringSlice(m["instances"])

	return r
}

// expandServiceTrafficMap expands the contents of ServiceTraffic into a JSON
// request object.
func expandServiceTrafficMap(c *Client, f map[string]ServiceTraffic, res *Service) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTraffic(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTrafficSlice expands the contents of ServiceTraffic into a JSON
// request object.
func expandServiceTrafficSlice(c *Client, f []ServiceTraffic, res *Service) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTraffic(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTrafficMap flattens the contents of ServiceTraffic from a JSON
// response object.
func flattenServiceTrafficMap(c *Client, i interface{}, res *Service) map[string]ServiceTraffic {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTraffic{}
	}

	if len(a) == 0 {
		return map[string]ServiceTraffic{}
	}

	items := make(map[string]ServiceTraffic)
	for k, item := range a {
		items[k] = *flattenServiceTraffic(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenServiceTrafficSlice flattens the contents of ServiceTraffic from a JSON
// response object.
func flattenServiceTrafficSlice(c *Client, i interface{}, res *Service) []ServiceTraffic {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTraffic{}
	}

	if len(a) == 0 {
		return []ServiceTraffic{}
	}

	items := make([]ServiceTraffic, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTraffic(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandServiceTraffic expands an instance of ServiceTraffic into a JSON
// request object.
func expandServiceTraffic(c *Client, f *ServiceTraffic, res *Service) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.Revision; !dcl.IsEmptyValueIndirect(v) {
		m["revision"] = v
	}
	if v := f.Percent; !dcl.IsEmptyValueIndirect(v) {
		m["percent"] = v
	}
	if v := f.Tag; !dcl.IsEmptyValueIndirect(v) {
		m["tag"] = v
	}

	return m, nil
}

// flattenServiceTraffic flattens an instance of ServiceTraffic from a JSON
// response object.
func flattenServiceTraffic(c *Client, i interface{}, res *Service) *ServiceTraffic {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTraffic{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTraffic
	}
	r.Type = flattenServiceTrafficTypeEnum(m["type"])
	r.Revision = dcl.FlattenString(m["revision"])
	r.Percent = dcl.FlattenInteger(m["percent"])
	r.Tag = dcl.FlattenString(m["tag"])

	return r
}

// expandServiceTerminalConditionMap expands the contents of ServiceTerminalCondition into a JSON
// request object.
func expandServiceTerminalConditionMap(c *Client, f map[string]ServiceTerminalCondition, res *Service) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTerminalCondition(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTerminalConditionSlice expands the contents of ServiceTerminalCondition into a JSON
// request object.
func expandServiceTerminalConditionSlice(c *Client, f []ServiceTerminalCondition, res *Service) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTerminalCondition(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTerminalConditionMap flattens the contents of ServiceTerminalCondition from a JSON
// response object.
func flattenServiceTerminalConditionMap(c *Client, i interface{}, res *Service) map[string]ServiceTerminalCondition {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTerminalCondition{}
	}

	if len(a) == 0 {
		return map[string]ServiceTerminalCondition{}
	}

	items := make(map[string]ServiceTerminalCondition)
	for k, item := range a {
		items[k] = *flattenServiceTerminalCondition(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenServiceTerminalConditionSlice flattens the contents of ServiceTerminalCondition from a JSON
// response object.
func flattenServiceTerminalConditionSlice(c *Client, i interface{}, res *Service) []ServiceTerminalCondition {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTerminalCondition{}
	}

	if len(a) == 0 {
		return []ServiceTerminalCondition{}
	}

	items := make([]ServiceTerminalCondition, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTerminalCondition(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandServiceTerminalCondition expands an instance of ServiceTerminalCondition into a JSON
// request object.
func expandServiceTerminalCondition(c *Client, f *ServiceTerminalCondition, res *Service) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v := f.Message; !dcl.IsEmptyValueIndirect(v) {
		m["message"] = v
	}
	if v := f.LastTransitionTime; !dcl.IsEmptyValueIndirect(v) {
		m["lastTransitionTime"] = v
	}
	if v := f.Severity; !dcl.IsEmptyValueIndirect(v) {
		m["severity"] = v
	}
	if v := f.Reason; !dcl.IsEmptyValueIndirect(v) {
		m["reason"] = v
	}
	if v := f.RevisionReason; !dcl.IsEmptyValueIndirect(v) {
		m["revisionReason"] = v
	}
	if v := f.JobReason; !dcl.IsEmptyValueIndirect(v) {
		m["jobReason"] = v
	}

	return m, nil
}

// flattenServiceTerminalCondition flattens an instance of ServiceTerminalCondition from a JSON
// response object.
func flattenServiceTerminalCondition(c *Client, i interface{}, res *Service) *ServiceTerminalCondition {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTerminalCondition{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTerminalCondition
	}
	r.Type = dcl.FlattenString(m["type"])
	r.State = flattenServiceTerminalConditionStateEnum(m["state"])
	r.Message = dcl.FlattenString(m["message"])
	r.LastTransitionTime = dcl.FlattenString(m["lastTransitionTime"])
	r.Severity = flattenServiceTerminalConditionSeverityEnum(m["severity"])
	r.Reason = flattenServiceTerminalConditionReasonEnum(m["reason"])
	r.RevisionReason = flattenServiceTerminalConditionRevisionReasonEnum(m["revisionReason"])
	r.JobReason = flattenServiceTerminalConditionJobReasonEnum(m["jobReason"])

	return r
}

// expandServiceTrafficStatusesMap expands the contents of ServiceTrafficStatuses into a JSON
// request object.
func expandServiceTrafficStatusesMap(c *Client, f map[string]ServiceTrafficStatuses, res *Service) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceTrafficStatuses(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceTrafficStatusesSlice expands the contents of ServiceTrafficStatuses into a JSON
// request object.
func expandServiceTrafficStatusesSlice(c *Client, f []ServiceTrafficStatuses, res *Service) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceTrafficStatuses(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceTrafficStatusesMap flattens the contents of ServiceTrafficStatuses from a JSON
// response object.
func flattenServiceTrafficStatusesMap(c *Client, i interface{}, res *Service) map[string]ServiceTrafficStatuses {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTrafficStatuses{}
	}

	if len(a) == 0 {
		return map[string]ServiceTrafficStatuses{}
	}

	items := make(map[string]ServiceTrafficStatuses)
	for k, item := range a {
		items[k] = *flattenServiceTrafficStatuses(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenServiceTrafficStatusesSlice flattens the contents of ServiceTrafficStatuses from a JSON
// response object.
func flattenServiceTrafficStatusesSlice(c *Client, i interface{}, res *Service) []ServiceTrafficStatuses {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTrafficStatuses{}
	}

	if len(a) == 0 {
		return []ServiceTrafficStatuses{}
	}

	items := make([]ServiceTrafficStatuses, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTrafficStatuses(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandServiceTrafficStatuses expands an instance of ServiceTrafficStatuses into a JSON
// request object.
func expandServiceTrafficStatuses(c *Client, f *ServiceTrafficStatuses, res *Service) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.Revision; !dcl.IsEmptyValueIndirect(v) {
		m["revision"] = v
	}
	if v := f.Percent; !dcl.IsEmptyValueIndirect(v) {
		m["percent"] = v
	}
	if v := f.Tag; !dcl.IsEmptyValueIndirect(v) {
		m["tag"] = v
	}
	if v := f.Uri; !dcl.IsEmptyValueIndirect(v) {
		m["uri"] = v
	}

	return m, nil
}

// flattenServiceTrafficStatuses flattens an instance of ServiceTrafficStatuses from a JSON
// response object.
func flattenServiceTrafficStatuses(c *Client, i interface{}, res *Service) *ServiceTrafficStatuses {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceTrafficStatuses{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceTrafficStatuses
	}
	r.Type = flattenServiceTrafficStatusesTypeEnum(m["type"])
	r.Revision = dcl.FlattenString(m["revision"])
	r.Percent = dcl.FlattenInteger(m["percent"])
	r.Tag = dcl.FlattenString(m["tag"])
	r.Uri = dcl.FlattenString(m["uri"])

	return r
}

// flattenServiceIngressEnumMap flattens the contents of ServiceIngressEnum from a JSON
// response object.
func flattenServiceIngressEnumMap(c *Client, i interface{}, res *Service) map[string]ServiceIngressEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceIngressEnum{}
	}

	if len(a) == 0 {
		return map[string]ServiceIngressEnum{}
	}

	items := make(map[string]ServiceIngressEnum)
	for k, item := range a {
		items[k] = *flattenServiceIngressEnum(item.(interface{}))
	}

	return items
}

// flattenServiceIngressEnumSlice flattens the contents of ServiceIngressEnum from a JSON
// response object.
func flattenServiceIngressEnumSlice(c *Client, i interface{}, res *Service) []ServiceIngressEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceIngressEnum{}
	}

	if len(a) == 0 {
		return []ServiceIngressEnum{}
	}

	items := make([]ServiceIngressEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceIngressEnum(item.(interface{})))
	}

	return items
}

// flattenServiceIngressEnum asserts that an interface is a string, and returns a
// pointer to a *ServiceIngressEnum with the same value as that string.
func flattenServiceIngressEnum(i interface{}) *ServiceIngressEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ServiceIngressEnumRef(s)
}

// flattenServiceLaunchStageEnumMap flattens the contents of ServiceLaunchStageEnum from a JSON
// response object.
func flattenServiceLaunchStageEnumMap(c *Client, i interface{}, res *Service) map[string]ServiceLaunchStageEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceLaunchStageEnum{}
	}

	if len(a) == 0 {
		return map[string]ServiceLaunchStageEnum{}
	}

	items := make(map[string]ServiceLaunchStageEnum)
	for k, item := range a {
		items[k] = *flattenServiceLaunchStageEnum(item.(interface{}))
	}

	return items
}

// flattenServiceLaunchStageEnumSlice flattens the contents of ServiceLaunchStageEnum from a JSON
// response object.
func flattenServiceLaunchStageEnumSlice(c *Client, i interface{}, res *Service) []ServiceLaunchStageEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceLaunchStageEnum{}
	}

	if len(a) == 0 {
		return []ServiceLaunchStageEnum{}
	}

	items := make([]ServiceLaunchStageEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceLaunchStageEnum(item.(interface{})))
	}

	return items
}

// flattenServiceLaunchStageEnum asserts that an interface is a string, and returns a
// pointer to a *ServiceLaunchStageEnum with the same value as that string.
func flattenServiceLaunchStageEnum(i interface{}) *ServiceLaunchStageEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ServiceLaunchStageEnumRef(s)
}

// flattenServiceTemplateVPCAccessEgressEnumMap flattens the contents of ServiceTemplateVPCAccessEgressEnum from a JSON
// response object.
func flattenServiceTemplateVPCAccessEgressEnumMap(c *Client, i interface{}, res *Service) map[string]ServiceTemplateVPCAccessEgressEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateVPCAccessEgressEnum{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateVPCAccessEgressEnum{}
	}

	items := make(map[string]ServiceTemplateVPCAccessEgressEnum)
	for k, item := range a {
		items[k] = *flattenServiceTemplateVPCAccessEgressEnum(item.(interface{}))
	}

	return items
}

// flattenServiceTemplateVPCAccessEgressEnumSlice flattens the contents of ServiceTemplateVPCAccessEgressEnum from a JSON
// response object.
func flattenServiceTemplateVPCAccessEgressEnumSlice(c *Client, i interface{}, res *Service) []ServiceTemplateVPCAccessEgressEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateVPCAccessEgressEnum{}
	}

	if len(a) == 0 {
		return []ServiceTemplateVPCAccessEgressEnum{}
	}

	items := make([]ServiceTemplateVPCAccessEgressEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateVPCAccessEgressEnum(item.(interface{})))
	}

	return items
}

// flattenServiceTemplateVPCAccessEgressEnum asserts that an interface is a string, and returns a
// pointer to a *ServiceTemplateVPCAccessEgressEnum with the same value as that string.
func flattenServiceTemplateVPCAccessEgressEnum(i interface{}) *ServiceTemplateVPCAccessEgressEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ServiceTemplateVPCAccessEgressEnumRef(s)
}

// flattenServiceTemplateExecutionEnvironmentEnumMap flattens the contents of ServiceTemplateExecutionEnvironmentEnum from a JSON
// response object.
func flattenServiceTemplateExecutionEnvironmentEnumMap(c *Client, i interface{}, res *Service) map[string]ServiceTemplateExecutionEnvironmentEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTemplateExecutionEnvironmentEnum{}
	}

	if len(a) == 0 {
		return map[string]ServiceTemplateExecutionEnvironmentEnum{}
	}

	items := make(map[string]ServiceTemplateExecutionEnvironmentEnum)
	for k, item := range a {
		items[k] = *flattenServiceTemplateExecutionEnvironmentEnum(item.(interface{}))
	}

	return items
}

// flattenServiceTemplateExecutionEnvironmentEnumSlice flattens the contents of ServiceTemplateExecutionEnvironmentEnum from a JSON
// response object.
func flattenServiceTemplateExecutionEnvironmentEnumSlice(c *Client, i interface{}, res *Service) []ServiceTemplateExecutionEnvironmentEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTemplateExecutionEnvironmentEnum{}
	}

	if len(a) == 0 {
		return []ServiceTemplateExecutionEnvironmentEnum{}
	}

	items := make([]ServiceTemplateExecutionEnvironmentEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTemplateExecutionEnvironmentEnum(item.(interface{})))
	}

	return items
}

// flattenServiceTemplateExecutionEnvironmentEnum asserts that an interface is a string, and returns a
// pointer to a *ServiceTemplateExecutionEnvironmentEnum with the same value as that string.
func flattenServiceTemplateExecutionEnvironmentEnum(i interface{}) *ServiceTemplateExecutionEnvironmentEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ServiceTemplateExecutionEnvironmentEnumRef(s)
}

// flattenServiceTrafficTypeEnumMap flattens the contents of ServiceTrafficTypeEnum from a JSON
// response object.
func flattenServiceTrafficTypeEnumMap(c *Client, i interface{}, res *Service) map[string]ServiceTrafficTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTrafficTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]ServiceTrafficTypeEnum{}
	}

	items := make(map[string]ServiceTrafficTypeEnum)
	for k, item := range a {
		items[k] = *flattenServiceTrafficTypeEnum(item.(interface{}))
	}

	return items
}

// flattenServiceTrafficTypeEnumSlice flattens the contents of ServiceTrafficTypeEnum from a JSON
// response object.
func flattenServiceTrafficTypeEnumSlice(c *Client, i interface{}, res *Service) []ServiceTrafficTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTrafficTypeEnum{}
	}

	if len(a) == 0 {
		return []ServiceTrafficTypeEnum{}
	}

	items := make([]ServiceTrafficTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTrafficTypeEnum(item.(interface{})))
	}

	return items
}

// flattenServiceTrafficTypeEnum asserts that an interface is a string, and returns a
// pointer to a *ServiceTrafficTypeEnum with the same value as that string.
func flattenServiceTrafficTypeEnum(i interface{}) *ServiceTrafficTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ServiceTrafficTypeEnumRef(s)
}

// flattenServiceTerminalConditionStateEnumMap flattens the contents of ServiceTerminalConditionStateEnum from a JSON
// response object.
func flattenServiceTerminalConditionStateEnumMap(c *Client, i interface{}, res *Service) map[string]ServiceTerminalConditionStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTerminalConditionStateEnum{}
	}

	if len(a) == 0 {
		return map[string]ServiceTerminalConditionStateEnum{}
	}

	items := make(map[string]ServiceTerminalConditionStateEnum)
	for k, item := range a {
		items[k] = *flattenServiceTerminalConditionStateEnum(item.(interface{}))
	}

	return items
}

// flattenServiceTerminalConditionStateEnumSlice flattens the contents of ServiceTerminalConditionStateEnum from a JSON
// response object.
func flattenServiceTerminalConditionStateEnumSlice(c *Client, i interface{}, res *Service) []ServiceTerminalConditionStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTerminalConditionStateEnum{}
	}

	if len(a) == 0 {
		return []ServiceTerminalConditionStateEnum{}
	}

	items := make([]ServiceTerminalConditionStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTerminalConditionStateEnum(item.(interface{})))
	}

	return items
}

// flattenServiceTerminalConditionStateEnum asserts that an interface is a string, and returns a
// pointer to a *ServiceTerminalConditionStateEnum with the same value as that string.
func flattenServiceTerminalConditionStateEnum(i interface{}) *ServiceTerminalConditionStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ServiceTerminalConditionStateEnumRef(s)
}

// flattenServiceTerminalConditionSeverityEnumMap flattens the contents of ServiceTerminalConditionSeverityEnum from a JSON
// response object.
func flattenServiceTerminalConditionSeverityEnumMap(c *Client, i interface{}, res *Service) map[string]ServiceTerminalConditionSeverityEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTerminalConditionSeverityEnum{}
	}

	if len(a) == 0 {
		return map[string]ServiceTerminalConditionSeverityEnum{}
	}

	items := make(map[string]ServiceTerminalConditionSeverityEnum)
	for k, item := range a {
		items[k] = *flattenServiceTerminalConditionSeverityEnum(item.(interface{}))
	}

	return items
}

// flattenServiceTerminalConditionSeverityEnumSlice flattens the contents of ServiceTerminalConditionSeverityEnum from a JSON
// response object.
func flattenServiceTerminalConditionSeverityEnumSlice(c *Client, i interface{}, res *Service) []ServiceTerminalConditionSeverityEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTerminalConditionSeverityEnum{}
	}

	if len(a) == 0 {
		return []ServiceTerminalConditionSeverityEnum{}
	}

	items := make([]ServiceTerminalConditionSeverityEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTerminalConditionSeverityEnum(item.(interface{})))
	}

	return items
}

// flattenServiceTerminalConditionSeverityEnum asserts that an interface is a string, and returns a
// pointer to a *ServiceTerminalConditionSeverityEnum with the same value as that string.
func flattenServiceTerminalConditionSeverityEnum(i interface{}) *ServiceTerminalConditionSeverityEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ServiceTerminalConditionSeverityEnumRef(s)
}

// flattenServiceTerminalConditionReasonEnumMap flattens the contents of ServiceTerminalConditionReasonEnum from a JSON
// response object.
func flattenServiceTerminalConditionReasonEnumMap(c *Client, i interface{}, res *Service) map[string]ServiceTerminalConditionReasonEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTerminalConditionReasonEnum{}
	}

	if len(a) == 0 {
		return map[string]ServiceTerminalConditionReasonEnum{}
	}

	items := make(map[string]ServiceTerminalConditionReasonEnum)
	for k, item := range a {
		items[k] = *flattenServiceTerminalConditionReasonEnum(item.(interface{}))
	}

	return items
}

// flattenServiceTerminalConditionReasonEnumSlice flattens the contents of ServiceTerminalConditionReasonEnum from a JSON
// response object.
func flattenServiceTerminalConditionReasonEnumSlice(c *Client, i interface{}, res *Service) []ServiceTerminalConditionReasonEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTerminalConditionReasonEnum{}
	}

	if len(a) == 0 {
		return []ServiceTerminalConditionReasonEnum{}
	}

	items := make([]ServiceTerminalConditionReasonEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTerminalConditionReasonEnum(item.(interface{})))
	}

	return items
}

// flattenServiceTerminalConditionReasonEnum asserts that an interface is a string, and returns a
// pointer to a *ServiceTerminalConditionReasonEnum with the same value as that string.
func flattenServiceTerminalConditionReasonEnum(i interface{}) *ServiceTerminalConditionReasonEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ServiceTerminalConditionReasonEnumRef(s)
}

// flattenServiceTerminalConditionRevisionReasonEnumMap flattens the contents of ServiceTerminalConditionRevisionReasonEnum from a JSON
// response object.
func flattenServiceTerminalConditionRevisionReasonEnumMap(c *Client, i interface{}, res *Service) map[string]ServiceTerminalConditionRevisionReasonEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTerminalConditionRevisionReasonEnum{}
	}

	if len(a) == 0 {
		return map[string]ServiceTerminalConditionRevisionReasonEnum{}
	}

	items := make(map[string]ServiceTerminalConditionRevisionReasonEnum)
	for k, item := range a {
		items[k] = *flattenServiceTerminalConditionRevisionReasonEnum(item.(interface{}))
	}

	return items
}

// flattenServiceTerminalConditionRevisionReasonEnumSlice flattens the contents of ServiceTerminalConditionRevisionReasonEnum from a JSON
// response object.
func flattenServiceTerminalConditionRevisionReasonEnumSlice(c *Client, i interface{}, res *Service) []ServiceTerminalConditionRevisionReasonEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTerminalConditionRevisionReasonEnum{}
	}

	if len(a) == 0 {
		return []ServiceTerminalConditionRevisionReasonEnum{}
	}

	items := make([]ServiceTerminalConditionRevisionReasonEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTerminalConditionRevisionReasonEnum(item.(interface{})))
	}

	return items
}

// flattenServiceTerminalConditionRevisionReasonEnum asserts that an interface is a string, and returns a
// pointer to a *ServiceTerminalConditionRevisionReasonEnum with the same value as that string.
func flattenServiceTerminalConditionRevisionReasonEnum(i interface{}) *ServiceTerminalConditionRevisionReasonEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ServiceTerminalConditionRevisionReasonEnumRef(s)
}

// flattenServiceTerminalConditionJobReasonEnumMap flattens the contents of ServiceTerminalConditionJobReasonEnum from a JSON
// response object.
func flattenServiceTerminalConditionJobReasonEnumMap(c *Client, i interface{}, res *Service) map[string]ServiceTerminalConditionJobReasonEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTerminalConditionJobReasonEnum{}
	}

	if len(a) == 0 {
		return map[string]ServiceTerminalConditionJobReasonEnum{}
	}

	items := make(map[string]ServiceTerminalConditionJobReasonEnum)
	for k, item := range a {
		items[k] = *flattenServiceTerminalConditionJobReasonEnum(item.(interface{}))
	}

	return items
}

// flattenServiceTerminalConditionJobReasonEnumSlice flattens the contents of ServiceTerminalConditionJobReasonEnum from a JSON
// response object.
func flattenServiceTerminalConditionJobReasonEnumSlice(c *Client, i interface{}, res *Service) []ServiceTerminalConditionJobReasonEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTerminalConditionJobReasonEnum{}
	}

	if len(a) == 0 {
		return []ServiceTerminalConditionJobReasonEnum{}
	}

	items := make([]ServiceTerminalConditionJobReasonEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTerminalConditionJobReasonEnum(item.(interface{})))
	}

	return items
}

// flattenServiceTerminalConditionJobReasonEnum asserts that an interface is a string, and returns a
// pointer to a *ServiceTerminalConditionJobReasonEnum with the same value as that string.
func flattenServiceTerminalConditionJobReasonEnum(i interface{}) *ServiceTerminalConditionJobReasonEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ServiceTerminalConditionJobReasonEnumRef(s)
}

// flattenServiceTrafficStatusesTypeEnumMap flattens the contents of ServiceTrafficStatusesTypeEnum from a JSON
// response object.
func flattenServiceTrafficStatusesTypeEnumMap(c *Client, i interface{}, res *Service) map[string]ServiceTrafficStatusesTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceTrafficStatusesTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]ServiceTrafficStatusesTypeEnum{}
	}

	items := make(map[string]ServiceTrafficStatusesTypeEnum)
	for k, item := range a {
		items[k] = *flattenServiceTrafficStatusesTypeEnum(item.(interface{}))
	}

	return items
}

// flattenServiceTrafficStatusesTypeEnumSlice flattens the contents of ServiceTrafficStatusesTypeEnum from a JSON
// response object.
func flattenServiceTrafficStatusesTypeEnumSlice(c *Client, i interface{}, res *Service) []ServiceTrafficStatusesTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceTrafficStatusesTypeEnum{}
	}

	if len(a) == 0 {
		return []ServiceTrafficStatusesTypeEnum{}
	}

	items := make([]ServiceTrafficStatusesTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceTrafficStatusesTypeEnum(item.(interface{})))
	}

	return items
}

// flattenServiceTrafficStatusesTypeEnum asserts that an interface is a string, and returns a
// pointer to a *ServiceTrafficStatusesTypeEnum with the same value as that string.
func flattenServiceTrafficStatusesTypeEnum(i interface{}) *ServiceTrafficStatusesTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ServiceTrafficStatusesTypeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Service) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalService(b, c, r)
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

type serviceDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         serviceApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToServiceDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]serviceDiff, error) {
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
	var diffs []serviceDiff
	// For each operation name, create a serviceDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := serviceDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToServiceApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToServiceApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (serviceApiOperation, error) {
	switch opName {

	case "updateServiceUpdateServiceOperation":
		return &updateServiceUpdateServiceOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractServiceFields(r *Service) error {
	vBinaryAuthorization := r.BinaryAuthorization
	if vBinaryAuthorization == nil {
		// note: explicitly not the empty object.
		vBinaryAuthorization = &ServiceBinaryAuthorization{}
	}
	if err := extractServiceBinaryAuthorizationFields(r, vBinaryAuthorization); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vBinaryAuthorization) {
		r.BinaryAuthorization = vBinaryAuthorization
	}
	vTemplate := r.Template
	if vTemplate == nil {
		// note: explicitly not the empty object.
		vTemplate = &ServiceTemplate{}
	}
	if err := extractServiceTemplateFields(r, vTemplate); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTemplate) {
		r.Template = vTemplate
	}
	vTerminalCondition := r.TerminalCondition
	if vTerminalCondition == nil {
		// note: explicitly not the empty object.
		vTerminalCondition = &ServiceTerminalCondition{}
	}
	if err := extractServiceTerminalConditionFields(r, vTerminalCondition); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTerminalCondition) {
		r.TerminalCondition = vTerminalCondition
	}
	return nil
}
func extractServiceBinaryAuthorizationFields(r *Service, o *ServiceBinaryAuthorization) error {
	return nil
}
func extractServiceTemplateFields(r *Service, o *ServiceTemplate) error {
	vScaling := o.Scaling
	if vScaling == nil {
		// note: explicitly not the empty object.
		vScaling = &ServiceTemplateScaling{}
	}
	if err := extractServiceTemplateScalingFields(r, vScaling); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vScaling) {
		o.Scaling = vScaling
	}
	vVPCAccess := o.VPCAccess
	if vVPCAccess == nil {
		// note: explicitly not the empty object.
		vVPCAccess = &ServiceTemplateVPCAccess{}
	}
	if err := extractServiceTemplateVPCAccessFields(r, vVPCAccess); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vVPCAccess) {
		o.VPCAccess = vVPCAccess
	}
	return nil
}
func extractServiceTemplateScalingFields(r *Service, o *ServiceTemplateScaling) error {
	return nil
}
func extractServiceTemplateVPCAccessFields(r *Service, o *ServiceTemplateVPCAccess) error {
	return nil
}
func extractServiceTemplateContainersFields(r *Service, o *ServiceTemplateContainers) error {
	vResources := o.Resources
	if vResources == nil {
		// note: explicitly not the empty object.
		vResources = &ServiceTemplateContainersResources{}
	}
	if err := extractServiceTemplateContainersResourcesFields(r, vResources); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vResources) {
		o.Resources = vResources
	}
	return nil
}
func extractServiceTemplateContainersEnvFields(r *Service, o *ServiceTemplateContainersEnv) error {
	vValueSource := o.ValueSource
	if vValueSource == nil {
		// note: explicitly not the empty object.
		vValueSource = &ServiceTemplateContainersEnvValueSource{}
	}
	if err := extractServiceTemplateContainersEnvValueSourceFields(r, vValueSource); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vValueSource) {
		o.ValueSource = vValueSource
	}
	return nil
}
func extractServiceTemplateContainersEnvValueSourceFields(r *Service, o *ServiceTemplateContainersEnvValueSource) error {
	vSecretKeyRef := o.SecretKeyRef
	if vSecretKeyRef == nil {
		// note: explicitly not the empty object.
		vSecretKeyRef = &ServiceTemplateContainersEnvValueSourceSecretKeyRef{}
	}
	if err := extractServiceTemplateContainersEnvValueSourceSecretKeyRefFields(r, vSecretKeyRef); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSecretKeyRef) {
		o.SecretKeyRef = vSecretKeyRef
	}
	return nil
}
func extractServiceTemplateContainersEnvValueSourceSecretKeyRefFields(r *Service, o *ServiceTemplateContainersEnvValueSourceSecretKeyRef) error {
	return nil
}
func extractServiceTemplateContainersResourcesFields(r *Service, o *ServiceTemplateContainersResources) error {
	return nil
}
func extractServiceTemplateContainersPortsFields(r *Service, o *ServiceTemplateContainersPorts) error {
	return nil
}
func extractServiceTemplateContainersVolumeMountsFields(r *Service, o *ServiceTemplateContainersVolumeMounts) error {
	return nil
}
func extractServiceTemplateVolumesFields(r *Service, o *ServiceTemplateVolumes) error {
	vSecret := o.Secret
	if vSecret == nil {
		// note: explicitly not the empty object.
		vSecret = &ServiceTemplateVolumesSecret{}
	}
	if err := extractServiceTemplateVolumesSecretFields(r, vSecret); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSecret) {
		o.Secret = vSecret
	}
	vCloudSqlInstance := o.CloudSqlInstance
	if vCloudSqlInstance == nil {
		// note: explicitly not the empty object.
		vCloudSqlInstance = &ServiceTemplateVolumesCloudSqlInstance{}
	}
	if err := extractServiceTemplateVolumesCloudSqlInstanceFields(r, vCloudSqlInstance); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCloudSqlInstance) {
		o.CloudSqlInstance = vCloudSqlInstance
	}
	return nil
}
func extractServiceTemplateVolumesSecretFields(r *Service, o *ServiceTemplateVolumesSecret) error {
	return nil
}
func extractServiceTemplateVolumesSecretItemsFields(r *Service, o *ServiceTemplateVolumesSecretItems) error {
	return nil
}
func extractServiceTemplateVolumesCloudSqlInstanceFields(r *Service, o *ServiceTemplateVolumesCloudSqlInstance) error {
	return nil
}
func extractServiceTrafficFields(r *Service, o *ServiceTraffic) error {
	return nil
}
func extractServiceTerminalConditionFields(r *Service, o *ServiceTerminalCondition) error {
	return nil
}
func extractServiceTrafficStatusesFields(r *Service, o *ServiceTrafficStatuses) error {
	return nil
}

func postReadExtractServiceFields(r *Service) error {
	vBinaryAuthorization := r.BinaryAuthorization
	if vBinaryAuthorization == nil {
		// note: explicitly not the empty object.
		vBinaryAuthorization = &ServiceBinaryAuthorization{}
	}
	if err := postReadExtractServiceBinaryAuthorizationFields(r, vBinaryAuthorization); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vBinaryAuthorization) {
		r.BinaryAuthorization = vBinaryAuthorization
	}
	vTemplate := r.Template
	if vTemplate == nil {
		// note: explicitly not the empty object.
		vTemplate = &ServiceTemplate{}
	}
	if err := postReadExtractServiceTemplateFields(r, vTemplate); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTemplate) {
		r.Template = vTemplate
	}
	vTerminalCondition := r.TerminalCondition
	if vTerminalCondition == nil {
		// note: explicitly not the empty object.
		vTerminalCondition = &ServiceTerminalCondition{}
	}
	if err := postReadExtractServiceTerminalConditionFields(r, vTerminalCondition); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTerminalCondition) {
		r.TerminalCondition = vTerminalCondition
	}
	return nil
}
func postReadExtractServiceBinaryAuthorizationFields(r *Service, o *ServiceBinaryAuthorization) error {
	return nil
}
func postReadExtractServiceTemplateFields(r *Service, o *ServiceTemplate) error {
	vScaling := o.Scaling
	if vScaling == nil {
		// note: explicitly not the empty object.
		vScaling = &ServiceTemplateScaling{}
	}
	if err := extractServiceTemplateScalingFields(r, vScaling); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vScaling) {
		o.Scaling = vScaling
	}
	vVPCAccess := o.VPCAccess
	if vVPCAccess == nil {
		// note: explicitly not the empty object.
		vVPCAccess = &ServiceTemplateVPCAccess{}
	}
	if err := extractServiceTemplateVPCAccessFields(r, vVPCAccess); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vVPCAccess) {
		o.VPCAccess = vVPCAccess
	}
	return nil
}
func postReadExtractServiceTemplateScalingFields(r *Service, o *ServiceTemplateScaling) error {
	return nil
}
func postReadExtractServiceTemplateVPCAccessFields(r *Service, o *ServiceTemplateVPCAccess) error {
	return nil
}
func postReadExtractServiceTemplateContainersFields(r *Service, o *ServiceTemplateContainers) error {
	vResources := o.Resources
	if vResources == nil {
		// note: explicitly not the empty object.
		vResources = &ServiceTemplateContainersResources{}
	}
	if err := extractServiceTemplateContainersResourcesFields(r, vResources); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vResources) {
		o.Resources = vResources
	}
	return nil
}
func postReadExtractServiceTemplateContainersEnvFields(r *Service, o *ServiceTemplateContainersEnv) error {
	vValueSource := o.ValueSource
	if vValueSource == nil {
		// note: explicitly not the empty object.
		vValueSource = &ServiceTemplateContainersEnvValueSource{}
	}
	if err := extractServiceTemplateContainersEnvValueSourceFields(r, vValueSource); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vValueSource) {
		o.ValueSource = vValueSource
	}
	return nil
}
func postReadExtractServiceTemplateContainersEnvValueSourceFields(r *Service, o *ServiceTemplateContainersEnvValueSource) error {
	vSecretKeyRef := o.SecretKeyRef
	if vSecretKeyRef == nil {
		// note: explicitly not the empty object.
		vSecretKeyRef = &ServiceTemplateContainersEnvValueSourceSecretKeyRef{}
	}
	if err := extractServiceTemplateContainersEnvValueSourceSecretKeyRefFields(r, vSecretKeyRef); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSecretKeyRef) {
		o.SecretKeyRef = vSecretKeyRef
	}
	return nil
}
func postReadExtractServiceTemplateContainersEnvValueSourceSecretKeyRefFields(r *Service, o *ServiceTemplateContainersEnvValueSourceSecretKeyRef) error {
	return nil
}
func postReadExtractServiceTemplateContainersResourcesFields(r *Service, o *ServiceTemplateContainersResources) error {
	return nil
}
func postReadExtractServiceTemplateContainersPortsFields(r *Service, o *ServiceTemplateContainersPorts) error {
	return nil
}
func postReadExtractServiceTemplateContainersVolumeMountsFields(r *Service, o *ServiceTemplateContainersVolumeMounts) error {
	return nil
}
func postReadExtractServiceTemplateVolumesFields(r *Service, o *ServiceTemplateVolumes) error {
	vSecret := o.Secret
	if vSecret == nil {
		// note: explicitly not the empty object.
		vSecret = &ServiceTemplateVolumesSecret{}
	}
	if err := extractServiceTemplateVolumesSecretFields(r, vSecret); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSecret) {
		o.Secret = vSecret
	}
	vCloudSqlInstance := o.CloudSqlInstance
	if vCloudSqlInstance == nil {
		// note: explicitly not the empty object.
		vCloudSqlInstance = &ServiceTemplateVolumesCloudSqlInstance{}
	}
	if err := extractServiceTemplateVolumesCloudSqlInstanceFields(r, vCloudSqlInstance); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCloudSqlInstance) {
		o.CloudSqlInstance = vCloudSqlInstance
	}
	return nil
}
func postReadExtractServiceTemplateVolumesSecretFields(r *Service, o *ServiceTemplateVolumesSecret) error {
	return nil
}
func postReadExtractServiceTemplateVolumesSecretItemsFields(r *Service, o *ServiceTemplateVolumesSecretItems) error {
	return nil
}
func postReadExtractServiceTemplateVolumesCloudSqlInstanceFields(r *Service, o *ServiceTemplateVolumesCloudSqlInstance) error {
	return nil
}
func postReadExtractServiceTrafficFields(r *Service, o *ServiceTraffic) error {
	return nil
}
func postReadExtractServiceTerminalConditionFields(r *Service, o *ServiceTerminalCondition) error {
	return nil
}
func postReadExtractServiceTrafficStatusesFields(r *Service, o *ServiceTrafficStatuses) error {
	return nil
}
