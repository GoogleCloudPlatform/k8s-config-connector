// Copyright 2023 Google LLC. All Rights Reserved.
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

func (r *Job) validate() error {

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
	if !dcl.IsEmptyValueIndirect(r.LatestSucceededExecution) {
		if err := r.LatestSucceededExecution.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.LatestCreatedExecution) {
		if err := r.LatestCreatedExecution.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobBinaryAuthorization) validate() error {
	return nil
}
func (r *JobTemplate) validate() error {
	if err := dcl.Required(r, "template"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Template) {
		if err := r.Template.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTemplateTemplate) validate() error {
	if !dcl.IsEmptyValueIndirect(r.VPCAccess) {
		if err := r.VPCAccess.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTemplateTemplateContainers) validate() error {
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
func (r *JobTemplateTemplateContainersEnv) validate() error {
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
func (r *JobTemplateTemplateContainersEnvValueSource) validate() error {
	if !dcl.IsEmptyValueIndirect(r.SecretKeyRef) {
		if err := r.SecretKeyRef.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTemplateTemplateContainersEnvValueSourceSecretKeyRef) validate() error {
	if err := dcl.Required(r, "secret"); err != nil {
		return err
	}
	return nil
}
func (r *JobTemplateTemplateContainersResources) validate() error {
	return nil
}
func (r *JobTemplateTemplateContainersPorts) validate() error {
	return nil
}
func (r *JobTemplateTemplateContainersVolumeMounts) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "mountPath"); err != nil {
		return err
	}
	return nil
}
func (r *JobTemplateTemplateVolumes) validate() error {
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
func (r *JobTemplateTemplateVolumesSecret) validate() error {
	if err := dcl.Required(r, "secret"); err != nil {
		return err
	}
	return nil
}
func (r *JobTemplateTemplateVolumesSecretItems) validate() error {
	if err := dcl.Required(r, "path"); err != nil {
		return err
	}
	return nil
}
func (r *JobTemplateTemplateVolumesCloudSqlInstance) validate() error {
	return nil
}
func (r *JobTemplateTemplateVPCAccess) validate() error {
	return nil
}
func (r *JobTerminalCondition) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"Reason", "InternalReason", "DomainMappingReason", "RevisionReason", "ExecutionReason"}, r.Reason, r.InternalReason, r.DomainMappingReason, r.RevisionReason, r.ExecutionReason); err != nil {
		return err
	}
	return nil
}
func (r *JobConditions) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"Reason", "RevisionReason", "ExecutionReason"}, r.Reason, r.RevisionReason, r.ExecutionReason); err != nil {
		return err
	}
	return nil
}
func (r *JobLatestSucceededExecution) validate() error {
	return nil
}
func (r *JobLatestCreatedExecution) validate() error {
	return nil
}
func (r *Job) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://run.googleapis.com/v2/", params)
}

func (r *Job) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/jobs/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Job) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/jobs", nr.basePath(), userBasePath, params), nil

}

func (r *Job) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/jobs?jobId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *Job) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/jobs/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Job) SetPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"project":  *nr.Project,
		"location": *nr.Location,
		"name":     *nr.Name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/jobs/{{name}}", nr.basePath(), userBasePath, fields)
}

func (r *Job) SetPolicyVerb() string {
	return "POST"
}

func (r *Job) getPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"project":  *nr.Project,
		"location": *nr.Location,
		"name":     *nr.Name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/jobs/{{name}}", nr.basePath(), userBasePath, fields)
}

func (r *Job) IAMPolicyVersion() int {
	return 3
}

// jobApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type jobApiOperation interface {
	do(context.Context, *Job, *Client) error
}

// newUpdateJobUpdateJobRequest creates a request for an
// Job resource's UpdateJob update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateJobUpdateJobRequest(ctx context.Context, f *Job, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v, err := dcl.DeriveField("projects/%s/locations/%s/jobs/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Location), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
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
	if v := f.LaunchStage; !dcl.IsEmptyValueIndirect(v) {
		req["launchStage"] = v
	}
	if v, err := expandJobBinaryAuthorization(c, f.BinaryAuthorization, res); err != nil {
		return nil, fmt.Errorf("error expanding BinaryAuthorization into binaryAuthorization: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["binaryAuthorization"] = v
	}
	if v, err := expandJobTemplate(c, f.Template, res); err != nil {
		return nil, fmt.Errorf("error expanding Template into template: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["template"] = v
	}
	if v, err := expandJobTerminalCondition(c, f.TerminalCondition, res); err != nil {
		return nil, fmt.Errorf("error expanding TerminalCondition into terminalCondition: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["terminalCondition"] = v
	}
	if v, err := expandJobLatestSucceededExecution(c, f.LatestSucceededExecution, res); err != nil {
		return nil, fmt.Errorf("error expanding LatestSucceededExecution into latestSucceededExecution: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["latestSucceededExecution"] = v
	}
	if v, err := expandJobLatestCreatedExecution(c, f.LatestCreatedExecution, res); err != nil {
		return nil, fmt.Errorf("error expanding LatestCreatedExecution into latestCreatedExecution: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["latestCreatedExecution"] = v
	}
	b, err := c.getJobRaw(ctx, f)
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

// marshalUpdateJobUpdateJobRequest converts the update into
// the final JSON request body.
func marshalUpdateJobUpdateJobRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateJobUpdateJobOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateJobUpdateJobOperation) do(ctx context.Context, r *Job, c *Client) error {
	_, err := c.GetJob(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateJob")
	if err != nil {
		return err
	}

	req, err := newUpdateJobUpdateJobRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateJobUpdateJobRequest(c, req)
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

func (c *Client) listJobRaw(ctx context.Context, r *Job, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != JobMaxPage {
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

type listJobOperation struct {
	Jobs  []map[string]interface{} `json:"jobs"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listJob(ctx context.Context, r *Job, pageToken string, pageSize int32) ([]*Job, string, error) {
	b, err := c.listJobRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listJobOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Job
	for _, v := range m.Jobs {
		res, err := unmarshalMapJob(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllJob(ctx context.Context, f func(*Job) bool, resources []*Job) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteJob(ctx, res)
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

type deleteJobOperation struct{}

func (op *deleteJobOperation) do(ctx context.Context, r *Job, c *Client) error {
	r, err := c.GetJob(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Job not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetJob checking for existence. error: %v", err)
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
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Job: %w", err)
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createJobOperation struct {
	response map[string]interface{}
}

func (op *createJobOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createJobOperation) do(ctx context.Context, r *Job, c *Client) error {
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

	if _, err := c.GetJob(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getJobRaw(ctx context.Context, r *Job) ([]byte, error) {

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

func (c *Client) jobDiffsForRawDesired(ctx context.Context, rawDesired *Job, opts ...dcl.ApplyOption) (initial, desired *Job, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Job
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Job); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Job, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetJob(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Job resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Job resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Job resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeJobDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Job: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Job: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractJobFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeJobInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Job: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeJobDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Job: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffJob(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeJobInitialState(rawInitial, rawDesired *Job) (*Job, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeJobDesiredState(rawDesired, rawInitial *Job, opts ...dcl.ApplyOption) (*Job, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.BinaryAuthorization = canonicalizeJobBinaryAuthorization(rawDesired.BinaryAuthorization, nil, opts...)
		rawDesired.Template = canonicalizeJobTemplate(rawDesired.Template, nil, opts...)
		rawDesired.TerminalCondition = canonicalizeJobTerminalCondition(rawDesired.TerminalCondition, nil, opts...)
		rawDesired.LatestSucceededExecution = canonicalizeJobLatestSucceededExecution(rawDesired.LatestSucceededExecution, nil, opts...)
		rawDesired.LatestCreatedExecution = canonicalizeJobLatestCreatedExecution(rawDesired.LatestCreatedExecution, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Job{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
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
	if dcl.IsZeroValue(rawDesired.LaunchStage) || (dcl.IsEmptyValueIndirect(rawDesired.LaunchStage) && dcl.IsEmptyValueIndirect(rawInitial.LaunchStage)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.LaunchStage = rawInitial.LaunchStage
	} else {
		canonicalDesired.LaunchStage = rawDesired.LaunchStage
	}
	canonicalDesired.BinaryAuthorization = canonicalizeJobBinaryAuthorization(rawDesired.BinaryAuthorization, rawInitial.BinaryAuthorization, opts...)
	canonicalDesired.Template = canonicalizeJobTemplate(rawDesired.Template, rawInitial.Template, opts...)
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

func canonicalizeJobNewState(c *Client, rawNew, rawDesired *Job) (*Job, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
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

	if dcl.IsEmptyValueIndirect(rawNew.LaunchStage) && dcl.IsEmptyValueIndirect(rawDesired.LaunchStage) {
		rawNew.LaunchStage = rawDesired.LaunchStage
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.BinaryAuthorization) && dcl.IsEmptyValueIndirect(rawDesired.BinaryAuthorization) {
		rawNew.BinaryAuthorization = rawDesired.BinaryAuthorization
	} else {
		rawNew.BinaryAuthorization = canonicalizeNewJobBinaryAuthorization(c, rawDesired.BinaryAuthorization, rawNew.BinaryAuthorization)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Template) && dcl.IsEmptyValueIndirect(rawDesired.Template) {
		rawNew.Template = rawDesired.Template
	} else {
		rawNew.Template = canonicalizeNewJobTemplate(c, rawDesired.Template, rawNew.Template)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ObservedGeneration) && dcl.IsEmptyValueIndirect(rawDesired.ObservedGeneration) {
		rawNew.ObservedGeneration = rawDesired.ObservedGeneration
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.TerminalCondition) && dcl.IsEmptyValueIndirect(rawDesired.TerminalCondition) {
		rawNew.TerminalCondition = rawDesired.TerminalCondition
	} else {
		rawNew.TerminalCondition = canonicalizeNewJobTerminalCondition(c, rawDesired.TerminalCondition, rawNew.TerminalCondition)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Conditions) && dcl.IsEmptyValueIndirect(rawDesired.Conditions) {
		rawNew.Conditions = rawDesired.Conditions
	} else {
		rawNew.Conditions = canonicalizeNewJobConditionsSlice(c, rawDesired.Conditions, rawNew.Conditions)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ExecutionCount) && dcl.IsEmptyValueIndirect(rawDesired.ExecutionCount) {
		rawNew.ExecutionCount = rawDesired.ExecutionCount
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.LatestSucceededExecution) && dcl.IsEmptyValueIndirect(rawDesired.LatestSucceededExecution) {
		rawNew.LatestSucceededExecution = rawDesired.LatestSucceededExecution
	} else {
		rawNew.LatestSucceededExecution = canonicalizeNewJobLatestSucceededExecution(c, rawDesired.LatestSucceededExecution, rawNew.LatestSucceededExecution)
	}

	if dcl.IsEmptyValueIndirect(rawNew.LatestCreatedExecution) && dcl.IsEmptyValueIndirect(rawDesired.LatestCreatedExecution) {
		rawNew.LatestCreatedExecution = rawDesired.LatestCreatedExecution
	} else {
		rawNew.LatestCreatedExecution = canonicalizeNewJobLatestCreatedExecution(c, rawDesired.LatestCreatedExecution, rawNew.LatestCreatedExecution)
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

func canonicalizeJobBinaryAuthorization(des, initial *JobBinaryAuthorization, opts ...dcl.ApplyOption) *JobBinaryAuthorization {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobBinaryAuthorization{}

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

func canonicalizeJobBinaryAuthorizationSlice(des, initial []JobBinaryAuthorization, opts ...dcl.ApplyOption) []JobBinaryAuthorization {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobBinaryAuthorization, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobBinaryAuthorization(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobBinaryAuthorization, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobBinaryAuthorization(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobBinaryAuthorization(c *Client, des, nw *JobBinaryAuthorization) *JobBinaryAuthorization {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobBinaryAuthorization while comparing non-nil desired to nil actual.  Returning desired object.")
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

func canonicalizeNewJobBinaryAuthorizationSet(c *Client, des, nw []JobBinaryAuthorization) []JobBinaryAuthorization {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobBinaryAuthorization
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobBinaryAuthorizationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobBinaryAuthorization(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobBinaryAuthorizationSlice(c *Client, des, nw []JobBinaryAuthorization) []JobBinaryAuthorization {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobBinaryAuthorization
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobBinaryAuthorization(c, &d, &n))
	}

	return items
}

func canonicalizeJobTemplate(des, initial *JobTemplate, opts ...dcl.ApplyOption) *JobTemplate {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTemplate{}

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
	if dcl.IsZeroValue(des.Parallelism) || (dcl.IsEmptyValueIndirect(des.Parallelism) && dcl.IsEmptyValueIndirect(initial.Parallelism)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Parallelism = initial.Parallelism
	} else {
		cDes.Parallelism = des.Parallelism
	}
	if dcl.IsZeroValue(des.TaskCount) || (dcl.IsEmptyValueIndirect(des.TaskCount) && dcl.IsEmptyValueIndirect(initial.TaskCount)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.TaskCount = initial.TaskCount
	} else {
		cDes.TaskCount = des.TaskCount
	}
	cDes.Template = canonicalizeJobTemplateTemplate(des.Template, initial.Template, opts...)

	return cDes
}

func canonicalizeJobTemplateSlice(des, initial []JobTemplate, opts ...dcl.ApplyOption) []JobTemplate {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTemplate, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTemplate(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTemplate, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTemplate(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTemplate(c *Client, des, nw *JobTemplate) *JobTemplate {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTemplate while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Template = canonicalizeNewJobTemplateTemplate(c, des.Template, nw.Template)

	return nw
}

func canonicalizeNewJobTemplateSet(c *Client, des, nw []JobTemplate) []JobTemplate {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTemplate
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTemplateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTemplate(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTemplateSlice(c *Client, des, nw []JobTemplate) []JobTemplate {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTemplate
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTemplate(c, &d, &n))
	}

	return items
}

func canonicalizeJobTemplateTemplate(des, initial *JobTemplateTemplate, opts ...dcl.ApplyOption) *JobTemplateTemplate {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTemplateTemplate{}

	cDes.Containers = canonicalizeJobTemplateTemplateContainersSlice(des.Containers, initial.Containers, opts...)
	cDes.Volumes = canonicalizeJobTemplateTemplateVolumesSlice(des.Volumes, initial.Volumes, opts...)
	if dcl.IsZeroValue(des.MaxRetries) || (dcl.IsEmptyValueIndirect(des.MaxRetries) && dcl.IsEmptyValueIndirect(initial.MaxRetries)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MaxRetries = initial.MaxRetries
	} else {
		cDes.MaxRetries = des.MaxRetries
	}
	if dcl.StringCanonicalize(des.Timeout, initial.Timeout) || dcl.IsZeroValue(des.Timeout) {
		cDes.Timeout = initial.Timeout
	} else {
		cDes.Timeout = des.Timeout
	}
	if dcl.StringCanonicalize(des.ServiceAccount, initial.ServiceAccount) || dcl.IsZeroValue(des.ServiceAccount) {
		cDes.ServiceAccount = initial.ServiceAccount
	} else {
		cDes.ServiceAccount = des.ServiceAccount
	}
	if dcl.IsZeroValue(des.ExecutionEnvironment) || (dcl.IsEmptyValueIndirect(des.ExecutionEnvironment) && dcl.IsEmptyValueIndirect(initial.ExecutionEnvironment)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ExecutionEnvironment = initial.ExecutionEnvironment
	} else {
		cDes.ExecutionEnvironment = des.ExecutionEnvironment
	}
	if dcl.IsZeroValue(des.EncryptionKey) || (dcl.IsEmptyValueIndirect(des.EncryptionKey) && dcl.IsEmptyValueIndirect(initial.EncryptionKey)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.EncryptionKey = initial.EncryptionKey
	} else {
		cDes.EncryptionKey = des.EncryptionKey
	}
	cDes.VPCAccess = canonicalizeJobTemplateTemplateVPCAccess(des.VPCAccess, initial.VPCAccess, opts...)

	return cDes
}

func canonicalizeJobTemplateTemplateSlice(des, initial []JobTemplateTemplate, opts ...dcl.ApplyOption) []JobTemplateTemplate {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTemplateTemplate, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTemplateTemplate(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTemplateTemplate, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTemplateTemplate(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTemplateTemplate(c *Client, des, nw *JobTemplateTemplate) *JobTemplateTemplate {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTemplateTemplate while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Containers = canonicalizeNewJobTemplateTemplateContainersSlice(c, des.Containers, nw.Containers)
	nw.Volumes = canonicalizeNewJobTemplateTemplateVolumesSlice(c, des.Volumes, nw.Volumes)
	if dcl.StringCanonicalize(des.Timeout, nw.Timeout) {
		nw.Timeout = des.Timeout
	}
	if dcl.StringCanonicalize(des.ServiceAccount, nw.ServiceAccount) {
		nw.ServiceAccount = des.ServiceAccount
	}
	nw.VPCAccess = canonicalizeNewJobTemplateTemplateVPCAccess(c, des.VPCAccess, nw.VPCAccess)

	return nw
}

func canonicalizeNewJobTemplateTemplateSet(c *Client, des, nw []JobTemplateTemplate) []JobTemplateTemplate {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTemplateTemplate
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTemplateTemplateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTemplateTemplate(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTemplateTemplateSlice(c *Client, des, nw []JobTemplateTemplate) []JobTemplateTemplate {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTemplateTemplate
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTemplateTemplate(c, &d, &n))
	}

	return items
}

func canonicalizeJobTemplateTemplateContainers(des, initial *JobTemplateTemplateContainers, opts ...dcl.ApplyOption) *JobTemplateTemplateContainers {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTemplateTemplateContainers{}

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
	cDes.Env = canonicalizeJobTemplateTemplateContainersEnvSlice(des.Env, initial.Env, opts...)
	cDes.Resources = canonicalizeJobTemplateTemplateContainersResources(des.Resources, initial.Resources, opts...)
	cDes.Ports = canonicalizeJobTemplateTemplateContainersPortsSlice(des.Ports, initial.Ports, opts...)
	cDes.VolumeMounts = canonicalizeJobTemplateTemplateContainersVolumeMountsSlice(des.VolumeMounts, initial.VolumeMounts, opts...)

	return cDes
}

func canonicalizeJobTemplateTemplateContainersSlice(des, initial []JobTemplateTemplateContainers, opts ...dcl.ApplyOption) []JobTemplateTemplateContainers {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTemplateTemplateContainers, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTemplateTemplateContainers(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTemplateTemplateContainers, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTemplateTemplateContainers(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTemplateTemplateContainers(c *Client, des, nw *JobTemplateTemplateContainers) *JobTemplateTemplateContainers {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTemplateTemplateContainers while comparing non-nil desired to nil actual.  Returning desired object.")
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
	nw.Env = canonicalizeNewJobTemplateTemplateContainersEnvSlice(c, des.Env, nw.Env)
	nw.Resources = canonicalizeNewJobTemplateTemplateContainersResources(c, des.Resources, nw.Resources)
	nw.Ports = canonicalizeNewJobTemplateTemplateContainersPortsSlice(c, des.Ports, nw.Ports)
	nw.VolumeMounts = canonicalizeNewJobTemplateTemplateContainersVolumeMountsSlice(c, des.VolumeMounts, nw.VolumeMounts)

	return nw
}

func canonicalizeNewJobTemplateTemplateContainersSet(c *Client, des, nw []JobTemplateTemplateContainers) []JobTemplateTemplateContainers {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTemplateTemplateContainers
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTemplateTemplateContainersNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTemplateTemplateContainers(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTemplateTemplateContainersSlice(c *Client, des, nw []JobTemplateTemplateContainers) []JobTemplateTemplateContainers {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTemplateTemplateContainers
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTemplateTemplateContainers(c, &d, &n))
	}

	return items
}

func canonicalizeJobTemplateTemplateContainersEnv(des, initial *JobTemplateTemplateContainersEnv, opts ...dcl.ApplyOption) *JobTemplateTemplateContainersEnv {
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

	cDes := &JobTemplateTemplateContainersEnv{}

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
	cDes.ValueSource = canonicalizeJobTemplateTemplateContainersEnvValueSource(des.ValueSource, initial.ValueSource, opts...)

	return cDes
}

func canonicalizeJobTemplateTemplateContainersEnvSlice(des, initial []JobTemplateTemplateContainersEnv, opts ...dcl.ApplyOption) []JobTemplateTemplateContainersEnv {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTemplateTemplateContainersEnv, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTemplateTemplateContainersEnv(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTemplateTemplateContainersEnv, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTemplateTemplateContainersEnv(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTemplateTemplateContainersEnv(c *Client, des, nw *JobTemplateTemplateContainersEnv) *JobTemplateTemplateContainersEnv {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTemplateTemplateContainersEnv while comparing non-nil desired to nil actual.  Returning desired object.")
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
	nw.ValueSource = canonicalizeNewJobTemplateTemplateContainersEnvValueSource(c, des.ValueSource, nw.ValueSource)

	return nw
}

func canonicalizeNewJobTemplateTemplateContainersEnvSet(c *Client, des, nw []JobTemplateTemplateContainersEnv) []JobTemplateTemplateContainersEnv {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTemplateTemplateContainersEnv
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTemplateTemplateContainersEnvNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTemplateTemplateContainersEnv(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTemplateTemplateContainersEnvSlice(c *Client, des, nw []JobTemplateTemplateContainersEnv) []JobTemplateTemplateContainersEnv {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTemplateTemplateContainersEnv
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTemplateTemplateContainersEnv(c, &d, &n))
	}

	return items
}

func canonicalizeJobTemplateTemplateContainersEnvValueSource(des, initial *JobTemplateTemplateContainersEnvValueSource, opts ...dcl.ApplyOption) *JobTemplateTemplateContainersEnvValueSource {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTemplateTemplateContainersEnvValueSource{}

	cDes.SecretKeyRef = canonicalizeJobTemplateTemplateContainersEnvValueSourceSecretKeyRef(des.SecretKeyRef, initial.SecretKeyRef, opts...)

	return cDes
}

func canonicalizeJobTemplateTemplateContainersEnvValueSourceSlice(des, initial []JobTemplateTemplateContainersEnvValueSource, opts ...dcl.ApplyOption) []JobTemplateTemplateContainersEnvValueSource {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTemplateTemplateContainersEnvValueSource, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTemplateTemplateContainersEnvValueSource(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTemplateTemplateContainersEnvValueSource, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTemplateTemplateContainersEnvValueSource(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTemplateTemplateContainersEnvValueSource(c *Client, des, nw *JobTemplateTemplateContainersEnvValueSource) *JobTemplateTemplateContainersEnvValueSource {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTemplateTemplateContainersEnvValueSource while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.SecretKeyRef = canonicalizeNewJobTemplateTemplateContainersEnvValueSourceSecretKeyRef(c, des.SecretKeyRef, nw.SecretKeyRef)

	return nw
}

func canonicalizeNewJobTemplateTemplateContainersEnvValueSourceSet(c *Client, des, nw []JobTemplateTemplateContainersEnvValueSource) []JobTemplateTemplateContainersEnvValueSource {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTemplateTemplateContainersEnvValueSource
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTemplateTemplateContainersEnvValueSourceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTemplateTemplateContainersEnvValueSource(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTemplateTemplateContainersEnvValueSourceSlice(c *Client, des, nw []JobTemplateTemplateContainersEnvValueSource) []JobTemplateTemplateContainersEnvValueSource {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTemplateTemplateContainersEnvValueSource
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTemplateTemplateContainersEnvValueSource(c, &d, &n))
	}

	return items
}

func canonicalizeJobTemplateTemplateContainersEnvValueSourceSecretKeyRef(des, initial *JobTemplateTemplateContainersEnvValueSourceSecretKeyRef, opts ...dcl.ApplyOption) *JobTemplateTemplateContainersEnvValueSourceSecretKeyRef {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTemplateTemplateContainersEnvValueSourceSecretKeyRef{}

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

func canonicalizeJobTemplateTemplateContainersEnvValueSourceSecretKeyRefSlice(des, initial []JobTemplateTemplateContainersEnvValueSourceSecretKeyRef, opts ...dcl.ApplyOption) []JobTemplateTemplateContainersEnvValueSourceSecretKeyRef {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTemplateTemplateContainersEnvValueSourceSecretKeyRef, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTemplateTemplateContainersEnvValueSourceSecretKeyRef(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTemplateTemplateContainersEnvValueSourceSecretKeyRef, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTemplateTemplateContainersEnvValueSourceSecretKeyRef(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTemplateTemplateContainersEnvValueSourceSecretKeyRef(c *Client, des, nw *JobTemplateTemplateContainersEnvValueSourceSecretKeyRef) *JobTemplateTemplateContainersEnvValueSourceSecretKeyRef {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTemplateTemplateContainersEnvValueSourceSecretKeyRef while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewJobTemplateTemplateContainersEnvValueSourceSecretKeyRefSet(c *Client, des, nw []JobTemplateTemplateContainersEnvValueSourceSecretKeyRef) []JobTemplateTemplateContainersEnvValueSourceSecretKeyRef {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTemplateTemplateContainersEnvValueSourceSecretKeyRef
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTemplateTemplateContainersEnvValueSourceSecretKeyRefNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTemplateTemplateContainersEnvValueSourceSecretKeyRef(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTemplateTemplateContainersEnvValueSourceSecretKeyRefSlice(c *Client, des, nw []JobTemplateTemplateContainersEnvValueSourceSecretKeyRef) []JobTemplateTemplateContainersEnvValueSourceSecretKeyRef {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTemplateTemplateContainersEnvValueSourceSecretKeyRef
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTemplateTemplateContainersEnvValueSourceSecretKeyRef(c, &d, &n))
	}

	return items
}

func canonicalizeJobTemplateTemplateContainersResources(des, initial *JobTemplateTemplateContainersResources, opts ...dcl.ApplyOption) *JobTemplateTemplateContainersResources {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTemplateTemplateContainersResources{}

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

func canonicalizeJobTemplateTemplateContainersResourcesSlice(des, initial []JobTemplateTemplateContainersResources, opts ...dcl.ApplyOption) []JobTemplateTemplateContainersResources {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTemplateTemplateContainersResources, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTemplateTemplateContainersResources(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTemplateTemplateContainersResources, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTemplateTemplateContainersResources(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTemplateTemplateContainersResources(c *Client, des, nw *JobTemplateTemplateContainersResources) *JobTemplateTemplateContainersResources {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTemplateTemplateContainersResources while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.CpuIdle, nw.CpuIdle) {
		nw.CpuIdle = des.CpuIdle
	}

	return nw
}

func canonicalizeNewJobTemplateTemplateContainersResourcesSet(c *Client, des, nw []JobTemplateTemplateContainersResources) []JobTemplateTemplateContainersResources {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTemplateTemplateContainersResources
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTemplateTemplateContainersResourcesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTemplateTemplateContainersResources(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTemplateTemplateContainersResourcesSlice(c *Client, des, nw []JobTemplateTemplateContainersResources) []JobTemplateTemplateContainersResources {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTemplateTemplateContainersResources
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTemplateTemplateContainersResources(c, &d, &n))
	}

	return items
}

func canonicalizeJobTemplateTemplateContainersPorts(des, initial *JobTemplateTemplateContainersPorts, opts ...dcl.ApplyOption) *JobTemplateTemplateContainersPorts {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTemplateTemplateContainersPorts{}

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

func canonicalizeJobTemplateTemplateContainersPortsSlice(des, initial []JobTemplateTemplateContainersPorts, opts ...dcl.ApplyOption) []JobTemplateTemplateContainersPorts {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTemplateTemplateContainersPorts, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTemplateTemplateContainersPorts(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTemplateTemplateContainersPorts, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTemplateTemplateContainersPorts(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTemplateTemplateContainersPorts(c *Client, des, nw *JobTemplateTemplateContainersPorts) *JobTemplateTemplateContainersPorts {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTemplateTemplateContainersPorts while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewJobTemplateTemplateContainersPortsSet(c *Client, des, nw []JobTemplateTemplateContainersPorts) []JobTemplateTemplateContainersPorts {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTemplateTemplateContainersPorts
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTemplateTemplateContainersPortsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTemplateTemplateContainersPorts(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTemplateTemplateContainersPortsSlice(c *Client, des, nw []JobTemplateTemplateContainersPorts) []JobTemplateTemplateContainersPorts {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTemplateTemplateContainersPorts
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTemplateTemplateContainersPorts(c, &d, &n))
	}

	return items
}

func canonicalizeJobTemplateTemplateContainersVolumeMounts(des, initial *JobTemplateTemplateContainersVolumeMounts, opts ...dcl.ApplyOption) *JobTemplateTemplateContainersVolumeMounts {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTemplateTemplateContainersVolumeMounts{}

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

func canonicalizeJobTemplateTemplateContainersVolumeMountsSlice(des, initial []JobTemplateTemplateContainersVolumeMounts, opts ...dcl.ApplyOption) []JobTemplateTemplateContainersVolumeMounts {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTemplateTemplateContainersVolumeMounts, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTemplateTemplateContainersVolumeMounts(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTemplateTemplateContainersVolumeMounts, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTemplateTemplateContainersVolumeMounts(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTemplateTemplateContainersVolumeMounts(c *Client, des, nw *JobTemplateTemplateContainersVolumeMounts) *JobTemplateTemplateContainersVolumeMounts {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTemplateTemplateContainersVolumeMounts while comparing non-nil desired to nil actual.  Returning desired object.")
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

func canonicalizeNewJobTemplateTemplateContainersVolumeMountsSet(c *Client, des, nw []JobTemplateTemplateContainersVolumeMounts) []JobTemplateTemplateContainersVolumeMounts {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTemplateTemplateContainersVolumeMounts
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTemplateTemplateContainersVolumeMountsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTemplateTemplateContainersVolumeMounts(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTemplateTemplateContainersVolumeMountsSlice(c *Client, des, nw []JobTemplateTemplateContainersVolumeMounts) []JobTemplateTemplateContainersVolumeMounts {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTemplateTemplateContainersVolumeMounts
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTemplateTemplateContainersVolumeMounts(c, &d, &n))
	}

	return items
}

func canonicalizeJobTemplateTemplateVolumes(des, initial *JobTemplateTemplateVolumes, opts ...dcl.ApplyOption) *JobTemplateTemplateVolumes {
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

	cDes := &JobTemplateTemplateVolumes{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	cDes.Secret = canonicalizeJobTemplateTemplateVolumesSecret(des.Secret, initial.Secret, opts...)
	cDes.CloudSqlInstance = canonicalizeJobTemplateTemplateVolumesCloudSqlInstance(des.CloudSqlInstance, initial.CloudSqlInstance, opts...)

	return cDes
}

func canonicalizeJobTemplateTemplateVolumesSlice(des, initial []JobTemplateTemplateVolumes, opts ...dcl.ApplyOption) []JobTemplateTemplateVolumes {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTemplateTemplateVolumes, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTemplateTemplateVolumes(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTemplateTemplateVolumes, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTemplateTemplateVolumes(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTemplateTemplateVolumes(c *Client, des, nw *JobTemplateTemplateVolumes) *JobTemplateTemplateVolumes {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTemplateTemplateVolumes while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	nw.Secret = canonicalizeNewJobTemplateTemplateVolumesSecret(c, des.Secret, nw.Secret)
	nw.CloudSqlInstance = canonicalizeNewJobTemplateTemplateVolumesCloudSqlInstance(c, des.CloudSqlInstance, nw.CloudSqlInstance)

	return nw
}

func canonicalizeNewJobTemplateTemplateVolumesSet(c *Client, des, nw []JobTemplateTemplateVolumes) []JobTemplateTemplateVolumes {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTemplateTemplateVolumes
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTemplateTemplateVolumesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTemplateTemplateVolumes(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTemplateTemplateVolumesSlice(c *Client, des, nw []JobTemplateTemplateVolumes) []JobTemplateTemplateVolumes {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTemplateTemplateVolumes
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTemplateTemplateVolumes(c, &d, &n))
	}

	return items
}

func canonicalizeJobTemplateTemplateVolumesSecret(des, initial *JobTemplateTemplateVolumesSecret, opts ...dcl.ApplyOption) *JobTemplateTemplateVolumesSecret {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTemplateTemplateVolumesSecret{}

	if dcl.StringCanonicalize(des.Secret, initial.Secret) || dcl.IsZeroValue(des.Secret) {
		cDes.Secret = initial.Secret
	} else {
		cDes.Secret = des.Secret
	}
	cDes.Items = canonicalizeJobTemplateTemplateVolumesSecretItemsSlice(des.Items, initial.Items, opts...)
	if dcl.IsZeroValue(des.DefaultMode) || (dcl.IsEmptyValueIndirect(des.DefaultMode) && dcl.IsEmptyValueIndirect(initial.DefaultMode)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.DefaultMode = initial.DefaultMode
	} else {
		cDes.DefaultMode = des.DefaultMode
	}

	return cDes
}

func canonicalizeJobTemplateTemplateVolumesSecretSlice(des, initial []JobTemplateTemplateVolumesSecret, opts ...dcl.ApplyOption) []JobTemplateTemplateVolumesSecret {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTemplateTemplateVolumesSecret, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTemplateTemplateVolumesSecret(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTemplateTemplateVolumesSecret, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTemplateTemplateVolumesSecret(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTemplateTemplateVolumesSecret(c *Client, des, nw *JobTemplateTemplateVolumesSecret) *JobTemplateTemplateVolumesSecret {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTemplateTemplateVolumesSecret while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Secret, nw.Secret) {
		nw.Secret = des.Secret
	}
	nw.Items = canonicalizeNewJobTemplateTemplateVolumesSecretItemsSlice(c, des.Items, nw.Items)

	return nw
}

func canonicalizeNewJobTemplateTemplateVolumesSecretSet(c *Client, des, nw []JobTemplateTemplateVolumesSecret) []JobTemplateTemplateVolumesSecret {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTemplateTemplateVolumesSecret
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTemplateTemplateVolumesSecretNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTemplateTemplateVolumesSecret(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTemplateTemplateVolumesSecretSlice(c *Client, des, nw []JobTemplateTemplateVolumesSecret) []JobTemplateTemplateVolumesSecret {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTemplateTemplateVolumesSecret
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTemplateTemplateVolumesSecret(c, &d, &n))
	}

	return items
}

func canonicalizeJobTemplateTemplateVolumesSecretItems(des, initial *JobTemplateTemplateVolumesSecretItems, opts ...dcl.ApplyOption) *JobTemplateTemplateVolumesSecretItems {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTemplateTemplateVolumesSecretItems{}

	if dcl.StringCanonicalize(des.Path, initial.Path) || dcl.IsZeroValue(des.Path) {
		cDes.Path = initial.Path
	} else {
		cDes.Path = des.Path
	}
	if dcl.StringCanonicalize(des.Version, initial.Version) || dcl.IsZeroValue(des.Version) {
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

func canonicalizeJobTemplateTemplateVolumesSecretItemsSlice(des, initial []JobTemplateTemplateVolumesSecretItems, opts ...dcl.ApplyOption) []JobTemplateTemplateVolumesSecretItems {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTemplateTemplateVolumesSecretItems, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTemplateTemplateVolumesSecretItems(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTemplateTemplateVolumesSecretItems, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTemplateTemplateVolumesSecretItems(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTemplateTemplateVolumesSecretItems(c *Client, des, nw *JobTemplateTemplateVolumesSecretItems) *JobTemplateTemplateVolumesSecretItems {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTemplateTemplateVolumesSecretItems while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Path, nw.Path) {
		nw.Path = des.Path
	}
	if dcl.StringCanonicalize(des.Version, nw.Version) {
		nw.Version = des.Version
	}

	return nw
}

func canonicalizeNewJobTemplateTemplateVolumesSecretItemsSet(c *Client, des, nw []JobTemplateTemplateVolumesSecretItems) []JobTemplateTemplateVolumesSecretItems {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTemplateTemplateVolumesSecretItems
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTemplateTemplateVolumesSecretItemsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTemplateTemplateVolumesSecretItems(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTemplateTemplateVolumesSecretItemsSlice(c *Client, des, nw []JobTemplateTemplateVolumesSecretItems) []JobTemplateTemplateVolumesSecretItems {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTemplateTemplateVolumesSecretItems
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTemplateTemplateVolumesSecretItems(c, &d, &n))
	}

	return items
}

func canonicalizeJobTemplateTemplateVolumesCloudSqlInstance(des, initial *JobTemplateTemplateVolumesCloudSqlInstance, opts ...dcl.ApplyOption) *JobTemplateTemplateVolumesCloudSqlInstance {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTemplateTemplateVolumesCloudSqlInstance{}

	if dcl.StringArrayCanonicalize(des.Instances, initial.Instances) {
		cDes.Instances = initial.Instances
	} else {
		cDes.Instances = des.Instances
	}

	return cDes
}

func canonicalizeJobTemplateTemplateVolumesCloudSqlInstanceSlice(des, initial []JobTemplateTemplateVolumesCloudSqlInstance, opts ...dcl.ApplyOption) []JobTemplateTemplateVolumesCloudSqlInstance {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTemplateTemplateVolumesCloudSqlInstance, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTemplateTemplateVolumesCloudSqlInstance(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTemplateTemplateVolumesCloudSqlInstance, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTemplateTemplateVolumesCloudSqlInstance(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTemplateTemplateVolumesCloudSqlInstance(c *Client, des, nw *JobTemplateTemplateVolumesCloudSqlInstance) *JobTemplateTemplateVolumesCloudSqlInstance {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTemplateTemplateVolumesCloudSqlInstance while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.Instances, nw.Instances) {
		nw.Instances = des.Instances
	}

	return nw
}

func canonicalizeNewJobTemplateTemplateVolumesCloudSqlInstanceSet(c *Client, des, nw []JobTemplateTemplateVolumesCloudSqlInstance) []JobTemplateTemplateVolumesCloudSqlInstance {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTemplateTemplateVolumesCloudSqlInstance
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTemplateTemplateVolumesCloudSqlInstanceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTemplateTemplateVolumesCloudSqlInstance(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTemplateTemplateVolumesCloudSqlInstanceSlice(c *Client, des, nw []JobTemplateTemplateVolumesCloudSqlInstance) []JobTemplateTemplateVolumesCloudSqlInstance {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTemplateTemplateVolumesCloudSqlInstance
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTemplateTemplateVolumesCloudSqlInstance(c, &d, &n))
	}

	return items
}

func canonicalizeJobTemplateTemplateVPCAccess(des, initial *JobTemplateTemplateVPCAccess, opts ...dcl.ApplyOption) *JobTemplateTemplateVPCAccess {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTemplateTemplateVPCAccess{}

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

func canonicalizeJobTemplateTemplateVPCAccessSlice(des, initial []JobTemplateTemplateVPCAccess, opts ...dcl.ApplyOption) []JobTemplateTemplateVPCAccess {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTemplateTemplateVPCAccess, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTemplateTemplateVPCAccess(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTemplateTemplateVPCAccess, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTemplateTemplateVPCAccess(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTemplateTemplateVPCAccess(c *Client, des, nw *JobTemplateTemplateVPCAccess) *JobTemplateTemplateVPCAccess {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTemplateTemplateVPCAccess while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewJobTemplateTemplateVPCAccessSet(c *Client, des, nw []JobTemplateTemplateVPCAccess) []JobTemplateTemplateVPCAccess {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTemplateTemplateVPCAccess
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTemplateTemplateVPCAccessNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTemplateTemplateVPCAccess(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTemplateTemplateVPCAccessSlice(c *Client, des, nw []JobTemplateTemplateVPCAccess) []JobTemplateTemplateVPCAccess {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTemplateTemplateVPCAccess
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTemplateTemplateVPCAccess(c, &d, &n))
	}

	return items
}

func canonicalizeJobTerminalCondition(des, initial *JobTerminalCondition, opts ...dcl.ApplyOption) *JobTerminalCondition {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.Reason != nil || (initial != nil && initial.Reason != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.InternalReason, des.DomainMappingReason, des.RevisionReason, des.ExecutionReason) {
			des.Reason = nil
			if initial != nil {
				initial.Reason = nil
			}
		}
	}

	if des.InternalReason != nil || (initial != nil && initial.InternalReason != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Reason, des.DomainMappingReason, des.RevisionReason, des.ExecutionReason) {
			des.InternalReason = nil
			if initial != nil {
				initial.InternalReason = nil
			}
		}
	}

	if des.DomainMappingReason != nil || (initial != nil && initial.DomainMappingReason != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Reason, des.InternalReason, des.RevisionReason, des.ExecutionReason) {
			des.DomainMappingReason = nil
			if initial != nil {
				initial.DomainMappingReason = nil
			}
		}
	}

	if des.RevisionReason != nil || (initial != nil && initial.RevisionReason != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Reason, des.InternalReason, des.DomainMappingReason, des.ExecutionReason) {
			des.RevisionReason = nil
			if initial != nil {
				initial.RevisionReason = nil
			}
		}
	}

	if des.ExecutionReason != nil || (initial != nil && initial.ExecutionReason != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Reason, des.InternalReason, des.DomainMappingReason, des.RevisionReason) {
			des.ExecutionReason = nil
			if initial != nil {
				initial.ExecutionReason = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &JobTerminalCondition{}

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
	if dcl.IsZeroValue(des.InternalReason) || (dcl.IsEmptyValueIndirect(des.InternalReason) && dcl.IsEmptyValueIndirect(initial.InternalReason)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.InternalReason = initial.InternalReason
	} else {
		cDes.InternalReason = des.InternalReason
	}
	if dcl.IsZeroValue(des.DomainMappingReason) || (dcl.IsEmptyValueIndirect(des.DomainMappingReason) && dcl.IsEmptyValueIndirect(initial.DomainMappingReason)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.DomainMappingReason = initial.DomainMappingReason
	} else {
		cDes.DomainMappingReason = des.DomainMappingReason
	}
	if dcl.IsZeroValue(des.RevisionReason) || (dcl.IsEmptyValueIndirect(des.RevisionReason) && dcl.IsEmptyValueIndirect(initial.RevisionReason)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.RevisionReason = initial.RevisionReason
	} else {
		cDes.RevisionReason = des.RevisionReason
	}
	if dcl.IsZeroValue(des.ExecutionReason) || (dcl.IsEmptyValueIndirect(des.ExecutionReason) && dcl.IsEmptyValueIndirect(initial.ExecutionReason)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ExecutionReason = initial.ExecutionReason
	} else {
		cDes.ExecutionReason = des.ExecutionReason
	}

	return cDes
}

func canonicalizeJobTerminalConditionSlice(des, initial []JobTerminalCondition, opts ...dcl.ApplyOption) []JobTerminalCondition {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTerminalCondition, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTerminalCondition(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTerminalCondition, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTerminalCondition(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTerminalCondition(c *Client, des, nw *JobTerminalCondition) *JobTerminalCondition {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTerminalCondition while comparing non-nil desired to nil actual.  Returning desired object.")
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

func canonicalizeNewJobTerminalConditionSet(c *Client, des, nw []JobTerminalCondition) []JobTerminalCondition {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTerminalCondition
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTerminalConditionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTerminalCondition(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTerminalConditionSlice(c *Client, des, nw []JobTerminalCondition) []JobTerminalCondition {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTerminalCondition
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTerminalCondition(c, &d, &n))
	}

	return items
}

func canonicalizeJobConditions(des, initial *JobConditions, opts ...dcl.ApplyOption) *JobConditions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.Reason != nil || (initial != nil && initial.Reason != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.RevisionReason, des.ExecutionReason) {
			des.Reason = nil
			if initial != nil {
				initial.Reason = nil
			}
		}
	}

	if des.RevisionReason != nil || (initial != nil && initial.RevisionReason != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Reason, des.ExecutionReason) {
			des.RevisionReason = nil
			if initial != nil {
				initial.RevisionReason = nil
			}
		}
	}

	if des.ExecutionReason != nil || (initial != nil && initial.ExecutionReason != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Reason, des.RevisionReason) {
			des.ExecutionReason = nil
			if initial != nil {
				initial.ExecutionReason = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &JobConditions{}

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
	if dcl.IsZeroValue(des.ExecutionReason) || (dcl.IsEmptyValueIndirect(des.ExecutionReason) && dcl.IsEmptyValueIndirect(initial.ExecutionReason)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ExecutionReason = initial.ExecutionReason
	} else {
		cDes.ExecutionReason = des.ExecutionReason
	}

	return cDes
}

func canonicalizeJobConditionsSlice(des, initial []JobConditions, opts ...dcl.ApplyOption) []JobConditions {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobConditions, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobConditions(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobConditions, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobConditions(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobConditions(c *Client, des, nw *JobConditions) *JobConditions {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobConditions while comparing non-nil desired to nil actual.  Returning desired object.")
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

func canonicalizeNewJobConditionsSet(c *Client, des, nw []JobConditions) []JobConditions {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobConditions
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobConditionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobConditions(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobConditionsSlice(c *Client, des, nw []JobConditions) []JobConditions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobConditions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobConditions(c, &d, &n))
	}

	return items
}

func canonicalizeJobLatestSucceededExecution(des, initial *JobLatestSucceededExecution, opts ...dcl.ApplyOption) *JobLatestSucceededExecution {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobLatestSucceededExecution{}

	if dcl.IsZeroValue(des.Name) || (dcl.IsEmptyValueIndirect(des.Name) && dcl.IsEmptyValueIndirect(initial.Name)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.IsZeroValue(des.CreateTime) || (dcl.IsEmptyValueIndirect(des.CreateTime) && dcl.IsEmptyValueIndirect(initial.CreateTime)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.CreateTime = initial.CreateTime
	} else {
		cDes.CreateTime = des.CreateTime
	}

	return cDes
}

func canonicalizeJobLatestSucceededExecutionSlice(des, initial []JobLatestSucceededExecution, opts ...dcl.ApplyOption) []JobLatestSucceededExecution {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobLatestSucceededExecution, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobLatestSucceededExecution(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobLatestSucceededExecution, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobLatestSucceededExecution(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobLatestSucceededExecution(c *Client, des, nw *JobLatestSucceededExecution) *JobLatestSucceededExecution {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobLatestSucceededExecution while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewJobLatestSucceededExecutionSet(c *Client, des, nw []JobLatestSucceededExecution) []JobLatestSucceededExecution {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobLatestSucceededExecution
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobLatestSucceededExecutionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobLatestSucceededExecution(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobLatestSucceededExecutionSlice(c *Client, des, nw []JobLatestSucceededExecution) []JobLatestSucceededExecution {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobLatestSucceededExecution
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobLatestSucceededExecution(c, &d, &n))
	}

	return items
}

func canonicalizeJobLatestCreatedExecution(des, initial *JobLatestCreatedExecution, opts ...dcl.ApplyOption) *JobLatestCreatedExecution {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobLatestCreatedExecution{}

	if dcl.IsZeroValue(des.Name) || (dcl.IsEmptyValueIndirect(des.Name) && dcl.IsEmptyValueIndirect(initial.Name)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.IsZeroValue(des.CreateTime) || (dcl.IsEmptyValueIndirect(des.CreateTime) && dcl.IsEmptyValueIndirect(initial.CreateTime)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.CreateTime = initial.CreateTime
	} else {
		cDes.CreateTime = des.CreateTime
	}

	return cDes
}

func canonicalizeJobLatestCreatedExecutionSlice(des, initial []JobLatestCreatedExecution, opts ...dcl.ApplyOption) []JobLatestCreatedExecution {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobLatestCreatedExecution, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobLatestCreatedExecution(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobLatestCreatedExecution, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobLatestCreatedExecution(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobLatestCreatedExecution(c *Client, des, nw *JobLatestCreatedExecution) *JobLatestCreatedExecution {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobLatestCreatedExecution while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewJobLatestCreatedExecutionSet(c *Client, des, nw []JobLatestCreatedExecution) []JobLatestCreatedExecution {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobLatestCreatedExecution
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobLatestCreatedExecutionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobLatestCreatedExecution(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobLatestCreatedExecutionSlice(c *Client, des, nw []JobLatestCreatedExecution) []JobLatestCreatedExecution {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobLatestCreatedExecution
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobLatestCreatedExecution(c, &d, &n))
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
func diffJob(c *Client, desired, actual *Job, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Annotations, actual.Annotations, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Annotations")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Client, actual.Client, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Client")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClientVersion, actual.ClientVersion, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("ClientVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LaunchStage, actual.LaunchStage, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("LaunchStage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BinaryAuthorization, actual.BinaryAuthorization, dcl.DiffInfo{ObjectFunction: compareJobBinaryAuthorizationNewStyle, EmptyObject: EmptyJobBinaryAuthorization, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("BinaryAuthorization")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Template, actual.Template, dcl.DiffInfo{ObjectFunction: compareJobTemplateNewStyle, EmptyObject: EmptyJobTemplate, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Template")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ObservedGeneration, actual.ObservedGeneration, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ObservedGeneration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TerminalCondition, actual.TerminalCondition, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareJobTerminalConditionNewStyle, EmptyObject: EmptyJobTerminalCondition, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TerminalCondition")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Conditions, actual.Conditions, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareJobConditionsNewStyle, EmptyObject: EmptyJobConditions, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Conditions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExecutionCount, actual.ExecutionCount, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExecutionCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LatestSucceededExecution, actual.LatestSucceededExecution, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareJobLatestSucceededExecutionNewStyle, EmptyObject: EmptyJobLatestSucceededExecution, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LatestSucceededExecution")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LatestCreatedExecution, actual.LatestCreatedExecution, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareJobLatestCreatedExecutionNewStyle, EmptyObject: EmptyJobLatestCreatedExecution, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LatestCreatedExecution")); len(ds) != 0 || err != nil {
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
func compareJobBinaryAuthorizationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobBinaryAuthorization)
	if !ok {
		desiredNotPointer, ok := d.(JobBinaryAuthorization)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobBinaryAuthorization or *JobBinaryAuthorization", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobBinaryAuthorization)
	if !ok {
		actualNotPointer, ok := a.(JobBinaryAuthorization)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobBinaryAuthorization", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.UseDefault, actual.UseDefault, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("UseDefault")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BreakglassJustification, actual.BreakglassJustification, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("BreakglassJustification")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTemplateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTemplate)
	if !ok {
		desiredNotPointer, ok := d.(JobTemplate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplate or *JobTemplate", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTemplate)
	if !ok {
		actualNotPointer, ok := a.(JobTemplate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplate", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Annotations, actual.Annotations, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Annotations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Parallelism, actual.Parallelism, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Parallelism")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TaskCount, actual.TaskCount, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("TaskCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Template, actual.Template, dcl.DiffInfo{ObjectFunction: compareJobTemplateTemplateNewStyle, EmptyObject: EmptyJobTemplateTemplate, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Template")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTemplateTemplateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTemplateTemplate)
	if !ok {
		desiredNotPointer, ok := d.(JobTemplateTemplate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplateTemplate or *JobTemplateTemplate", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTemplateTemplate)
	if !ok {
		actualNotPointer, ok := a.(JobTemplateTemplate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplateTemplate", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Containers, actual.Containers, dcl.DiffInfo{ObjectFunction: compareJobTemplateTemplateContainersNewStyle, EmptyObject: EmptyJobTemplateTemplateContainers, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Containers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Volumes, actual.Volumes, dcl.DiffInfo{ObjectFunction: compareJobTemplateTemplateVolumesNewStyle, EmptyObject: EmptyJobTemplateTemplateVolumes, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Volumes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxRetries, actual.MaxRetries, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("MaxRetries")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Timeout, actual.Timeout, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Timeout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceAccount, actual.ServiceAccount, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("ServiceAccount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExecutionEnvironment, actual.ExecutionEnvironment, dcl.DiffInfo{ServerDefault: true, Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("ExecutionEnvironment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EncryptionKey, actual.EncryptionKey, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("EncryptionKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VPCAccess, actual.VPCAccess, dcl.DiffInfo{ObjectFunction: compareJobTemplateTemplateVPCAccessNewStyle, EmptyObject: EmptyJobTemplateTemplateVPCAccess, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("VpcAccess")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTemplateTemplateContainersNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTemplateTemplateContainers)
	if !ok {
		desiredNotPointer, ok := d.(JobTemplateTemplateContainers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplateTemplateContainers or *JobTemplateTemplateContainers", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTemplateTemplateContainers)
	if !ok {
		actualNotPointer, ok := a.(JobTemplateTemplateContainers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplateTemplateContainers", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Image, actual.Image, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Image")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Command, actual.Command, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Command")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Args, actual.Args, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Args")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Env, actual.Env, dcl.DiffInfo{ObjectFunction: compareJobTemplateTemplateContainersEnvNewStyle, EmptyObject: EmptyJobTemplateTemplateContainersEnv, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Env")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Resources, actual.Resources, dcl.DiffInfo{ServerDefault: true, ObjectFunction: compareJobTemplateTemplateContainersResourcesNewStyle, EmptyObject: EmptyJobTemplateTemplateContainersResources, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Resources")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Ports, actual.Ports, dcl.DiffInfo{ObjectFunction: compareJobTemplateTemplateContainersPortsNewStyle, EmptyObject: EmptyJobTemplateTemplateContainersPorts, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Ports")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VolumeMounts, actual.VolumeMounts, dcl.DiffInfo{ObjectFunction: compareJobTemplateTemplateContainersVolumeMountsNewStyle, EmptyObject: EmptyJobTemplateTemplateContainersVolumeMounts, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("VolumeMounts")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTemplateTemplateContainersEnvNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTemplateTemplateContainersEnv)
	if !ok {
		desiredNotPointer, ok := d.(JobTemplateTemplateContainersEnv)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplateTemplateContainersEnv or *JobTemplateTemplateContainersEnv", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTemplateTemplateContainersEnv)
	if !ok {
		actualNotPointer, ok := a.(JobTemplateTemplateContainersEnv)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplateTemplateContainersEnv", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ValueSource, actual.ValueSource, dcl.DiffInfo{ObjectFunction: compareJobTemplateTemplateContainersEnvValueSourceNewStyle, EmptyObject: EmptyJobTemplateTemplateContainersEnvValueSource, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("ValueSource")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTemplateTemplateContainersEnvValueSourceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTemplateTemplateContainersEnvValueSource)
	if !ok {
		desiredNotPointer, ok := d.(JobTemplateTemplateContainersEnvValueSource)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplateTemplateContainersEnvValueSource or *JobTemplateTemplateContainersEnvValueSource", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTemplateTemplateContainersEnvValueSource)
	if !ok {
		actualNotPointer, ok := a.(JobTemplateTemplateContainersEnvValueSource)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplateTemplateContainersEnvValueSource", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SecretKeyRef, actual.SecretKeyRef, dcl.DiffInfo{ObjectFunction: compareJobTemplateTemplateContainersEnvValueSourceSecretKeyRefNewStyle, EmptyObject: EmptyJobTemplateTemplateContainersEnvValueSourceSecretKeyRef, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("SecretKeyRef")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTemplateTemplateContainersEnvValueSourceSecretKeyRefNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTemplateTemplateContainersEnvValueSourceSecretKeyRef)
	if !ok {
		desiredNotPointer, ok := d.(JobTemplateTemplateContainersEnvValueSourceSecretKeyRef)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplateTemplateContainersEnvValueSourceSecretKeyRef or *JobTemplateTemplateContainersEnvValueSourceSecretKeyRef", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTemplateTemplateContainersEnvValueSourceSecretKeyRef)
	if !ok {
		actualNotPointer, ok := a.(JobTemplateTemplateContainersEnvValueSourceSecretKeyRef)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplateTemplateContainersEnvValueSourceSecretKeyRef", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Secret, actual.Secret, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Secret")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Version, actual.Version, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Version")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTemplateTemplateContainersResourcesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTemplateTemplateContainersResources)
	if !ok {
		desiredNotPointer, ok := d.(JobTemplateTemplateContainersResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplateTemplateContainersResources or *JobTemplateTemplateContainersResources", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTemplateTemplateContainersResources)
	if !ok {
		actualNotPointer, ok := a.(JobTemplateTemplateContainersResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplateTemplateContainersResources", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Limits, actual.Limits, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Limits")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CpuIdle, actual.CpuIdle, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("CpuIdle")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTemplateTemplateContainersPortsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTemplateTemplateContainersPorts)
	if !ok {
		desiredNotPointer, ok := d.(JobTemplateTemplateContainersPorts)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplateTemplateContainersPorts or *JobTemplateTemplateContainersPorts", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTemplateTemplateContainersPorts)
	if !ok {
		actualNotPointer, ok := a.(JobTemplateTemplateContainersPorts)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplateTemplateContainersPorts", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ContainerPort, actual.ContainerPort, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("ContainerPort")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTemplateTemplateContainersVolumeMountsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTemplateTemplateContainersVolumeMounts)
	if !ok {
		desiredNotPointer, ok := d.(JobTemplateTemplateContainersVolumeMounts)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplateTemplateContainersVolumeMounts or *JobTemplateTemplateContainersVolumeMounts", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTemplateTemplateContainersVolumeMounts)
	if !ok {
		actualNotPointer, ok := a.(JobTemplateTemplateContainersVolumeMounts)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplateTemplateContainersVolumeMounts", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MountPath, actual.MountPath, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("MountPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTemplateTemplateVolumesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTemplateTemplateVolumes)
	if !ok {
		desiredNotPointer, ok := d.(JobTemplateTemplateVolumes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplateTemplateVolumes or *JobTemplateTemplateVolumes", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTemplateTemplateVolumes)
	if !ok {
		actualNotPointer, ok := a.(JobTemplateTemplateVolumes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplateTemplateVolumes", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Secret, actual.Secret, dcl.DiffInfo{ObjectFunction: compareJobTemplateTemplateVolumesSecretNewStyle, EmptyObject: EmptyJobTemplateTemplateVolumesSecret, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Secret")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CloudSqlInstance, actual.CloudSqlInstance, dcl.DiffInfo{ObjectFunction: compareJobTemplateTemplateVolumesCloudSqlInstanceNewStyle, EmptyObject: EmptyJobTemplateTemplateVolumesCloudSqlInstance, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("CloudSqlInstance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTemplateTemplateVolumesSecretNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTemplateTemplateVolumesSecret)
	if !ok {
		desiredNotPointer, ok := d.(JobTemplateTemplateVolumesSecret)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplateTemplateVolumesSecret or *JobTemplateTemplateVolumesSecret", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTemplateTemplateVolumesSecret)
	if !ok {
		actualNotPointer, ok := a.(JobTemplateTemplateVolumesSecret)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplateTemplateVolumesSecret", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Secret, actual.Secret, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Secret")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Items, actual.Items, dcl.DiffInfo{ObjectFunction: compareJobTemplateTemplateVolumesSecretItemsNewStyle, EmptyObject: EmptyJobTemplateTemplateVolumesSecretItems, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Items")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultMode, actual.DefaultMode, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("DefaultMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTemplateTemplateVolumesSecretItemsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTemplateTemplateVolumesSecretItems)
	if !ok {
		desiredNotPointer, ok := d.(JobTemplateTemplateVolumesSecretItems)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplateTemplateVolumesSecretItems or *JobTemplateTemplateVolumesSecretItems", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTemplateTemplateVolumesSecretItems)
	if !ok {
		actualNotPointer, ok := a.(JobTemplateTemplateVolumesSecretItems)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplateTemplateVolumesSecretItems", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Version, actual.Version, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Version")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Mode, actual.Mode, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Mode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTemplateTemplateVolumesCloudSqlInstanceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTemplateTemplateVolumesCloudSqlInstance)
	if !ok {
		desiredNotPointer, ok := d.(JobTemplateTemplateVolumesCloudSqlInstance)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplateTemplateVolumesCloudSqlInstance or *JobTemplateTemplateVolumesCloudSqlInstance", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTemplateTemplateVolumesCloudSqlInstance)
	if !ok {
		actualNotPointer, ok := a.(JobTemplateTemplateVolumesCloudSqlInstance)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplateTemplateVolumesCloudSqlInstance", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Instances, actual.Instances, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Instances")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTemplateTemplateVPCAccessNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTemplateTemplateVPCAccess)
	if !ok {
		desiredNotPointer, ok := d.(JobTemplateTemplateVPCAccess)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplateTemplateVPCAccess or *JobTemplateTemplateVPCAccess", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTemplateTemplateVPCAccess)
	if !ok {
		actualNotPointer, ok := a.(JobTemplateTemplateVPCAccess)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTemplateTemplateVPCAccess", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Connector, actual.Connector, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Connector")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Egress, actual.Egress, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Egress")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTerminalConditionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTerminalCondition)
	if !ok {
		desiredNotPointer, ok := d.(JobTerminalCondition)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTerminalCondition or *JobTerminalCondition", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTerminalCondition)
	if !ok {
		actualNotPointer, ok := a.(JobTerminalCondition)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTerminalCondition", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Message, actual.Message, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Message")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LastTransitionTime, actual.LastTransitionTime, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("LastTransitionTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Severity, actual.Severity, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Severity")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Reason, actual.Reason, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Reason")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InternalReason, actual.InternalReason, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("InternalReason")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DomainMappingReason, actual.DomainMappingReason, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("DomainMappingReason")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RevisionReason, actual.RevisionReason, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("RevisionReason")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExecutionReason, actual.ExecutionReason, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("ExecutionReason")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobConditionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobConditions)
	if !ok {
		desiredNotPointer, ok := d.(JobConditions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobConditions or *JobConditions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobConditions)
	if !ok {
		actualNotPointer, ok := a.(JobConditions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobConditions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Message, actual.Message, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Message")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LastTransitionTime, actual.LastTransitionTime, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("LastTransitionTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Severity, actual.Severity, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Severity")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Reason, actual.Reason, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Reason")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RevisionReason, actual.RevisionReason, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("RevisionReason")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExecutionReason, actual.ExecutionReason, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("ExecutionReason")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobLatestSucceededExecutionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobLatestSucceededExecution)
	if !ok {
		desiredNotPointer, ok := d.(JobLatestSucceededExecution)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobLatestSucceededExecution or *JobLatestSucceededExecution", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobLatestSucceededExecution)
	if !ok {
		actualNotPointer, ok := a.(JobLatestSucceededExecution)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobLatestSucceededExecution", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobLatestCreatedExecutionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobLatestCreatedExecution)
	if !ok {
		desiredNotPointer, ok := d.(JobLatestCreatedExecution)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobLatestCreatedExecution or *JobLatestCreatedExecution", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobLatestCreatedExecution)
	if !ok {
		actualNotPointer, ok := a.(JobLatestCreatedExecution)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobLatestCreatedExecution", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
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
func (r *Job) urlNormalized() *Job {
	normalized := dcl.Copy(*r).(Job)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Uid = dcl.SelfLinkToName(r.Uid)
	normalized.Creator = dcl.SelfLinkToName(r.Creator)
	normalized.LastModifier = dcl.SelfLinkToName(r.LastModifier)
	normalized.Client = dcl.SelfLinkToName(r.Client)
	normalized.ClientVersion = dcl.SelfLinkToName(r.ClientVersion)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Job) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateJob" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/jobs/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Job resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Job) marshal(c *Client) ([]byte, error) {
	m, err := expandJob(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Job: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalJob decodes JSON responses into the Job resource schema.
func unmarshalJob(b []byte, c *Client, res *Job) (*Job, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapJob(m, c, res)
}

func unmarshalMapJob(m map[string]interface{}, c *Client, res *Job) (*Job, error) {

	flattened := flattenJob(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandJob expands Job into a JSON request object.
func expandJob(c *Client, f *Job) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("projects/%s/locations/%s/jobs/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Location), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
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
	if v := f.LaunchStage; dcl.ValueShouldBeSent(v) {
		m["launchStage"] = v
	}
	if v, err := expandJobBinaryAuthorization(c, f.BinaryAuthorization, res); err != nil {
		return nil, fmt.Errorf("error expanding BinaryAuthorization into binaryAuthorization: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["binaryAuthorization"] = v
	}
	if v, err := expandJobTemplate(c, f.Template, res); err != nil {
		return nil, fmt.Errorf("error expanding Template into template: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["template"] = v
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

// flattenJob flattens Job from a JSON request object into the
// Job type.
func flattenJob(c *Client, i interface{}, res *Job) *Job {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Job{}
	resultRes.Name = dcl.FlattenString(m["name"])
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
	resultRes.LaunchStage = flattenJobLaunchStageEnum(m["launchStage"])
	resultRes.BinaryAuthorization = flattenJobBinaryAuthorization(c, m["binaryAuthorization"], res)
	resultRes.Template = flattenJobTemplate(c, m["template"], res)
	resultRes.ObservedGeneration = dcl.FlattenInteger(m["observedGeneration"])
	resultRes.TerminalCondition = flattenJobTerminalCondition(c, m["terminalCondition"], res)
	resultRes.Conditions = flattenJobConditionsSlice(c, m["conditions"], res)
	resultRes.ExecutionCount = dcl.FlattenInteger(m["executionCount"])
	resultRes.LatestSucceededExecution = flattenJobLatestSucceededExecution(c, m["latestSucceededExecution"], res)
	resultRes.LatestCreatedExecution = flattenJobLatestCreatedExecution(c, m["latestCreatedExecution"], res)
	resultRes.Reconciling = dcl.FlattenBool(m["reconciling"])
	resultRes.Etag = dcl.FlattenString(m["etag"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// expandJobBinaryAuthorizationMap expands the contents of JobBinaryAuthorization into a JSON
// request object.
func expandJobBinaryAuthorizationMap(c *Client, f map[string]JobBinaryAuthorization, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobBinaryAuthorization(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobBinaryAuthorizationSlice expands the contents of JobBinaryAuthorization into a JSON
// request object.
func expandJobBinaryAuthorizationSlice(c *Client, f []JobBinaryAuthorization, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobBinaryAuthorization(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobBinaryAuthorizationMap flattens the contents of JobBinaryAuthorization from a JSON
// response object.
func flattenJobBinaryAuthorizationMap(c *Client, i interface{}, res *Job) map[string]JobBinaryAuthorization {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobBinaryAuthorization{}
	}

	if len(a) == 0 {
		return map[string]JobBinaryAuthorization{}
	}

	items := make(map[string]JobBinaryAuthorization)
	for k, item := range a {
		items[k] = *flattenJobBinaryAuthorization(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobBinaryAuthorizationSlice flattens the contents of JobBinaryAuthorization from a JSON
// response object.
func flattenJobBinaryAuthorizationSlice(c *Client, i interface{}, res *Job) []JobBinaryAuthorization {
	a, ok := i.([]interface{})
	if !ok {
		return []JobBinaryAuthorization{}
	}

	if len(a) == 0 {
		return []JobBinaryAuthorization{}
	}

	items := make([]JobBinaryAuthorization, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobBinaryAuthorization(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobBinaryAuthorization expands an instance of JobBinaryAuthorization into a JSON
// request object.
func expandJobBinaryAuthorization(c *Client, f *JobBinaryAuthorization, res *Job) (map[string]interface{}, error) {
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

// flattenJobBinaryAuthorization flattens an instance of JobBinaryAuthorization from a JSON
// response object.
func flattenJobBinaryAuthorization(c *Client, i interface{}, res *Job) *JobBinaryAuthorization {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobBinaryAuthorization{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobBinaryAuthorization
	}
	r.UseDefault = dcl.FlattenBool(m["useDefault"])
	r.BreakglassJustification = dcl.FlattenString(m["breakglassJustification"])

	return r
}

// expandJobTemplateMap expands the contents of JobTemplate into a JSON
// request object.
func expandJobTemplateMap(c *Client, f map[string]JobTemplate, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTemplate(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTemplateSlice expands the contents of JobTemplate into a JSON
// request object.
func expandJobTemplateSlice(c *Client, f []JobTemplate, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTemplate(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTemplateMap flattens the contents of JobTemplate from a JSON
// response object.
func flattenJobTemplateMap(c *Client, i interface{}, res *Job) map[string]JobTemplate {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTemplate{}
	}

	if len(a) == 0 {
		return map[string]JobTemplate{}
	}

	items := make(map[string]JobTemplate)
	for k, item := range a {
		items[k] = *flattenJobTemplate(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTemplateSlice flattens the contents of JobTemplate from a JSON
// response object.
func flattenJobTemplateSlice(c *Client, i interface{}, res *Job) []JobTemplate {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTemplate{}
	}

	if len(a) == 0 {
		return []JobTemplate{}
	}

	items := make([]JobTemplate, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTemplate(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTemplate expands an instance of JobTemplate into a JSON
// request object.
func expandJobTemplate(c *Client, f *JobTemplate, res *Job) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.Annotations; !dcl.IsEmptyValueIndirect(v) {
		m["annotations"] = v
	}
	if v := f.Parallelism; !dcl.IsEmptyValueIndirect(v) {
		m["parallelism"] = v
	}
	if v := f.TaskCount; !dcl.IsEmptyValueIndirect(v) {
		m["taskCount"] = v
	}
	if v, err := expandJobTemplateTemplate(c, f.Template, res); err != nil {
		return nil, fmt.Errorf("error expanding Template into template: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["template"] = v
	}

	return m, nil
}

// flattenJobTemplate flattens an instance of JobTemplate from a JSON
// response object.
func flattenJobTemplate(c *Client, i interface{}, res *Job) *JobTemplate {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTemplate{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTemplate
	}
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.Annotations = dcl.FlattenKeyValuePairs(m["annotations"])
	r.Parallelism = dcl.FlattenInteger(m["parallelism"])
	r.TaskCount = dcl.FlattenInteger(m["taskCount"])
	r.Template = flattenJobTemplateTemplate(c, m["template"], res)

	return r
}

// expandJobTemplateTemplateMap expands the contents of JobTemplateTemplate into a JSON
// request object.
func expandJobTemplateTemplateMap(c *Client, f map[string]JobTemplateTemplate, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTemplateTemplate(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTemplateTemplateSlice expands the contents of JobTemplateTemplate into a JSON
// request object.
func expandJobTemplateTemplateSlice(c *Client, f []JobTemplateTemplate, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTemplateTemplate(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTemplateTemplateMap flattens the contents of JobTemplateTemplate from a JSON
// response object.
func flattenJobTemplateTemplateMap(c *Client, i interface{}, res *Job) map[string]JobTemplateTemplate {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTemplateTemplate{}
	}

	if len(a) == 0 {
		return map[string]JobTemplateTemplate{}
	}

	items := make(map[string]JobTemplateTemplate)
	for k, item := range a {
		items[k] = *flattenJobTemplateTemplate(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTemplateTemplateSlice flattens the contents of JobTemplateTemplate from a JSON
// response object.
func flattenJobTemplateTemplateSlice(c *Client, i interface{}, res *Job) []JobTemplateTemplate {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTemplateTemplate{}
	}

	if len(a) == 0 {
		return []JobTemplateTemplate{}
	}

	items := make([]JobTemplateTemplate, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTemplateTemplate(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTemplateTemplate expands an instance of JobTemplateTemplate into a JSON
// request object.
func expandJobTemplateTemplate(c *Client, f *JobTemplateTemplate, res *Job) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandJobTemplateTemplateContainersSlice(c, f.Containers, res); err != nil {
		return nil, fmt.Errorf("error expanding Containers into containers: %w", err)
	} else if v != nil {
		m["containers"] = v
	}
	if v, err := expandJobTemplateTemplateVolumesSlice(c, f.Volumes, res); err != nil {
		return nil, fmt.Errorf("error expanding Volumes into volumes: %w", err)
	} else if v != nil {
		m["volumes"] = v
	}
	if v := f.MaxRetries; !dcl.IsEmptyValueIndirect(v) {
		m["maxRetries"] = v
	}
	if v := f.Timeout; !dcl.IsEmptyValueIndirect(v) {
		m["timeout"] = v
	}
	if v := f.ServiceAccount; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccount"] = v
	}
	if v := f.ExecutionEnvironment; !dcl.IsEmptyValueIndirect(v) {
		m["executionEnvironment"] = v
	}
	if v := f.EncryptionKey; !dcl.IsEmptyValueIndirect(v) {
		m["encryptionKey"] = v
	}
	if v, err := expandJobTemplateTemplateVPCAccess(c, f.VPCAccess, res); err != nil {
		return nil, fmt.Errorf("error expanding VPCAccess into vpcAccess: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["vpcAccess"] = v
	}

	return m, nil
}

// flattenJobTemplateTemplate flattens an instance of JobTemplateTemplate from a JSON
// response object.
func flattenJobTemplateTemplate(c *Client, i interface{}, res *Job) *JobTemplateTemplate {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTemplateTemplate{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTemplateTemplate
	}
	r.Containers = flattenJobTemplateTemplateContainersSlice(c, m["containers"], res)
	r.Volumes = flattenJobTemplateTemplateVolumesSlice(c, m["volumes"], res)
	r.MaxRetries = dcl.FlattenInteger(m["maxRetries"])
	r.Timeout = dcl.FlattenString(m["timeout"])
	r.ServiceAccount = dcl.FlattenString(m["serviceAccount"])
	r.ExecutionEnvironment = flattenJobTemplateTemplateExecutionEnvironmentEnum(m["executionEnvironment"])
	r.EncryptionKey = dcl.FlattenString(m["encryptionKey"])
	r.VPCAccess = flattenJobTemplateTemplateVPCAccess(c, m["vpcAccess"], res)

	return r
}

// expandJobTemplateTemplateContainersMap expands the contents of JobTemplateTemplateContainers into a JSON
// request object.
func expandJobTemplateTemplateContainersMap(c *Client, f map[string]JobTemplateTemplateContainers, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTemplateTemplateContainers(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTemplateTemplateContainersSlice expands the contents of JobTemplateTemplateContainers into a JSON
// request object.
func expandJobTemplateTemplateContainersSlice(c *Client, f []JobTemplateTemplateContainers, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTemplateTemplateContainers(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTemplateTemplateContainersMap flattens the contents of JobTemplateTemplateContainers from a JSON
// response object.
func flattenJobTemplateTemplateContainersMap(c *Client, i interface{}, res *Job) map[string]JobTemplateTemplateContainers {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTemplateTemplateContainers{}
	}

	if len(a) == 0 {
		return map[string]JobTemplateTemplateContainers{}
	}

	items := make(map[string]JobTemplateTemplateContainers)
	for k, item := range a {
		items[k] = *flattenJobTemplateTemplateContainers(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTemplateTemplateContainersSlice flattens the contents of JobTemplateTemplateContainers from a JSON
// response object.
func flattenJobTemplateTemplateContainersSlice(c *Client, i interface{}, res *Job) []JobTemplateTemplateContainers {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTemplateTemplateContainers{}
	}

	if len(a) == 0 {
		return []JobTemplateTemplateContainers{}
	}

	items := make([]JobTemplateTemplateContainers, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTemplateTemplateContainers(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTemplateTemplateContainers expands an instance of JobTemplateTemplateContainers into a JSON
// request object.
func expandJobTemplateTemplateContainers(c *Client, f *JobTemplateTemplateContainers, res *Job) (map[string]interface{}, error) {
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
	if v, err := expandJobTemplateTemplateContainersEnvSlice(c, f.Env, res); err != nil {
		return nil, fmt.Errorf("error expanding Env into env: %w", err)
	} else if v != nil {
		m["env"] = v
	}
	if v, err := expandJobTemplateTemplateContainersResources(c, f.Resources, res); err != nil {
		return nil, fmt.Errorf("error expanding Resources into resources: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["resources"] = v
	}
	if v, err := expandJobTemplateTemplateContainersPortsSlice(c, f.Ports, res); err != nil {
		return nil, fmt.Errorf("error expanding Ports into ports: %w", err)
	} else if v != nil {
		m["ports"] = v
	}
	if v, err := expandJobTemplateTemplateContainersVolumeMountsSlice(c, f.VolumeMounts, res); err != nil {
		return nil, fmt.Errorf("error expanding VolumeMounts into volumeMounts: %w", err)
	} else if v != nil {
		m["volumeMounts"] = v
	}

	return m, nil
}

// flattenJobTemplateTemplateContainers flattens an instance of JobTemplateTemplateContainers from a JSON
// response object.
func flattenJobTemplateTemplateContainers(c *Client, i interface{}, res *Job) *JobTemplateTemplateContainers {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTemplateTemplateContainers{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTemplateTemplateContainers
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Image = dcl.FlattenString(m["image"])
	r.Command = dcl.FlattenStringSlice(m["command"])
	r.Args = dcl.FlattenStringSlice(m["args"])
	r.Env = flattenJobTemplateTemplateContainersEnvSlice(c, m["env"], res)
	r.Resources = flattenJobTemplateTemplateContainersResources(c, m["resources"], res)
	r.Ports = flattenJobTemplateTemplateContainersPortsSlice(c, m["ports"], res)
	r.VolumeMounts = flattenJobTemplateTemplateContainersVolumeMountsSlice(c, m["volumeMounts"], res)

	return r
}

// expandJobTemplateTemplateContainersEnvMap expands the contents of JobTemplateTemplateContainersEnv into a JSON
// request object.
func expandJobTemplateTemplateContainersEnvMap(c *Client, f map[string]JobTemplateTemplateContainersEnv, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTemplateTemplateContainersEnv(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTemplateTemplateContainersEnvSlice expands the contents of JobTemplateTemplateContainersEnv into a JSON
// request object.
func expandJobTemplateTemplateContainersEnvSlice(c *Client, f []JobTemplateTemplateContainersEnv, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTemplateTemplateContainersEnv(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTemplateTemplateContainersEnvMap flattens the contents of JobTemplateTemplateContainersEnv from a JSON
// response object.
func flattenJobTemplateTemplateContainersEnvMap(c *Client, i interface{}, res *Job) map[string]JobTemplateTemplateContainersEnv {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTemplateTemplateContainersEnv{}
	}

	if len(a) == 0 {
		return map[string]JobTemplateTemplateContainersEnv{}
	}

	items := make(map[string]JobTemplateTemplateContainersEnv)
	for k, item := range a {
		items[k] = *flattenJobTemplateTemplateContainersEnv(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTemplateTemplateContainersEnvSlice flattens the contents of JobTemplateTemplateContainersEnv from a JSON
// response object.
func flattenJobTemplateTemplateContainersEnvSlice(c *Client, i interface{}, res *Job) []JobTemplateTemplateContainersEnv {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTemplateTemplateContainersEnv{}
	}

	if len(a) == 0 {
		return []JobTemplateTemplateContainersEnv{}
	}

	items := make([]JobTemplateTemplateContainersEnv, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTemplateTemplateContainersEnv(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTemplateTemplateContainersEnv expands an instance of JobTemplateTemplateContainersEnv into a JSON
// request object.
func expandJobTemplateTemplateContainersEnv(c *Client, f *JobTemplateTemplateContainersEnv, res *Job) (map[string]interface{}, error) {
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
	if v, err := expandJobTemplateTemplateContainersEnvValueSource(c, f.ValueSource, res); err != nil {
		return nil, fmt.Errorf("error expanding ValueSource into valueSource: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["valueSource"] = v
	}

	return m, nil
}

// flattenJobTemplateTemplateContainersEnv flattens an instance of JobTemplateTemplateContainersEnv from a JSON
// response object.
func flattenJobTemplateTemplateContainersEnv(c *Client, i interface{}, res *Job) *JobTemplateTemplateContainersEnv {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTemplateTemplateContainersEnv{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTemplateTemplateContainersEnv
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Value = dcl.FlattenString(m["value"])
	r.ValueSource = flattenJobTemplateTemplateContainersEnvValueSource(c, m["valueSource"], res)

	return r
}

// expandJobTemplateTemplateContainersEnvValueSourceMap expands the contents of JobTemplateTemplateContainersEnvValueSource into a JSON
// request object.
func expandJobTemplateTemplateContainersEnvValueSourceMap(c *Client, f map[string]JobTemplateTemplateContainersEnvValueSource, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTemplateTemplateContainersEnvValueSource(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTemplateTemplateContainersEnvValueSourceSlice expands the contents of JobTemplateTemplateContainersEnvValueSource into a JSON
// request object.
func expandJobTemplateTemplateContainersEnvValueSourceSlice(c *Client, f []JobTemplateTemplateContainersEnvValueSource, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTemplateTemplateContainersEnvValueSource(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTemplateTemplateContainersEnvValueSourceMap flattens the contents of JobTemplateTemplateContainersEnvValueSource from a JSON
// response object.
func flattenJobTemplateTemplateContainersEnvValueSourceMap(c *Client, i interface{}, res *Job) map[string]JobTemplateTemplateContainersEnvValueSource {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTemplateTemplateContainersEnvValueSource{}
	}

	if len(a) == 0 {
		return map[string]JobTemplateTemplateContainersEnvValueSource{}
	}

	items := make(map[string]JobTemplateTemplateContainersEnvValueSource)
	for k, item := range a {
		items[k] = *flattenJobTemplateTemplateContainersEnvValueSource(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTemplateTemplateContainersEnvValueSourceSlice flattens the contents of JobTemplateTemplateContainersEnvValueSource from a JSON
// response object.
func flattenJobTemplateTemplateContainersEnvValueSourceSlice(c *Client, i interface{}, res *Job) []JobTemplateTemplateContainersEnvValueSource {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTemplateTemplateContainersEnvValueSource{}
	}

	if len(a) == 0 {
		return []JobTemplateTemplateContainersEnvValueSource{}
	}

	items := make([]JobTemplateTemplateContainersEnvValueSource, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTemplateTemplateContainersEnvValueSource(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTemplateTemplateContainersEnvValueSource expands an instance of JobTemplateTemplateContainersEnvValueSource into a JSON
// request object.
func expandJobTemplateTemplateContainersEnvValueSource(c *Client, f *JobTemplateTemplateContainersEnvValueSource, res *Job) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandJobTemplateTemplateContainersEnvValueSourceSecretKeyRef(c, f.SecretKeyRef, res); err != nil {
		return nil, fmt.Errorf("error expanding SecretKeyRef into secretKeyRef: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["secretKeyRef"] = v
	}

	return m, nil
}

// flattenJobTemplateTemplateContainersEnvValueSource flattens an instance of JobTemplateTemplateContainersEnvValueSource from a JSON
// response object.
func flattenJobTemplateTemplateContainersEnvValueSource(c *Client, i interface{}, res *Job) *JobTemplateTemplateContainersEnvValueSource {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTemplateTemplateContainersEnvValueSource{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTemplateTemplateContainersEnvValueSource
	}
	r.SecretKeyRef = flattenJobTemplateTemplateContainersEnvValueSourceSecretKeyRef(c, m["secretKeyRef"], res)

	return r
}

// expandJobTemplateTemplateContainersEnvValueSourceSecretKeyRefMap expands the contents of JobTemplateTemplateContainersEnvValueSourceSecretKeyRef into a JSON
// request object.
func expandJobTemplateTemplateContainersEnvValueSourceSecretKeyRefMap(c *Client, f map[string]JobTemplateTemplateContainersEnvValueSourceSecretKeyRef, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTemplateTemplateContainersEnvValueSourceSecretKeyRef(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTemplateTemplateContainersEnvValueSourceSecretKeyRefSlice expands the contents of JobTemplateTemplateContainersEnvValueSourceSecretKeyRef into a JSON
// request object.
func expandJobTemplateTemplateContainersEnvValueSourceSecretKeyRefSlice(c *Client, f []JobTemplateTemplateContainersEnvValueSourceSecretKeyRef, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTemplateTemplateContainersEnvValueSourceSecretKeyRef(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTemplateTemplateContainersEnvValueSourceSecretKeyRefMap flattens the contents of JobTemplateTemplateContainersEnvValueSourceSecretKeyRef from a JSON
// response object.
func flattenJobTemplateTemplateContainersEnvValueSourceSecretKeyRefMap(c *Client, i interface{}, res *Job) map[string]JobTemplateTemplateContainersEnvValueSourceSecretKeyRef {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTemplateTemplateContainersEnvValueSourceSecretKeyRef{}
	}

	if len(a) == 0 {
		return map[string]JobTemplateTemplateContainersEnvValueSourceSecretKeyRef{}
	}

	items := make(map[string]JobTemplateTemplateContainersEnvValueSourceSecretKeyRef)
	for k, item := range a {
		items[k] = *flattenJobTemplateTemplateContainersEnvValueSourceSecretKeyRef(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTemplateTemplateContainersEnvValueSourceSecretKeyRefSlice flattens the contents of JobTemplateTemplateContainersEnvValueSourceSecretKeyRef from a JSON
// response object.
func flattenJobTemplateTemplateContainersEnvValueSourceSecretKeyRefSlice(c *Client, i interface{}, res *Job) []JobTemplateTemplateContainersEnvValueSourceSecretKeyRef {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTemplateTemplateContainersEnvValueSourceSecretKeyRef{}
	}

	if len(a) == 0 {
		return []JobTemplateTemplateContainersEnvValueSourceSecretKeyRef{}
	}

	items := make([]JobTemplateTemplateContainersEnvValueSourceSecretKeyRef, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTemplateTemplateContainersEnvValueSourceSecretKeyRef(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTemplateTemplateContainersEnvValueSourceSecretKeyRef expands an instance of JobTemplateTemplateContainersEnvValueSourceSecretKeyRef into a JSON
// request object.
func expandJobTemplateTemplateContainersEnvValueSourceSecretKeyRef(c *Client, f *JobTemplateTemplateContainersEnvValueSourceSecretKeyRef, res *Job) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Secret; !dcl.IsEmptyValueIndirect(v) {
		m["secret"] = v
	}
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}

	return m, nil
}

// flattenJobTemplateTemplateContainersEnvValueSourceSecretKeyRef flattens an instance of JobTemplateTemplateContainersEnvValueSourceSecretKeyRef from a JSON
// response object.
func flattenJobTemplateTemplateContainersEnvValueSourceSecretKeyRef(c *Client, i interface{}, res *Job) *JobTemplateTemplateContainersEnvValueSourceSecretKeyRef {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTemplateTemplateContainersEnvValueSourceSecretKeyRef{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTemplateTemplateContainersEnvValueSourceSecretKeyRef
	}
	r.Secret = dcl.FlattenString(m["secret"])
	r.Version = dcl.FlattenString(m["version"])

	return r
}

// expandJobTemplateTemplateContainersResourcesMap expands the contents of JobTemplateTemplateContainersResources into a JSON
// request object.
func expandJobTemplateTemplateContainersResourcesMap(c *Client, f map[string]JobTemplateTemplateContainersResources, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTemplateTemplateContainersResources(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTemplateTemplateContainersResourcesSlice expands the contents of JobTemplateTemplateContainersResources into a JSON
// request object.
func expandJobTemplateTemplateContainersResourcesSlice(c *Client, f []JobTemplateTemplateContainersResources, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTemplateTemplateContainersResources(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTemplateTemplateContainersResourcesMap flattens the contents of JobTemplateTemplateContainersResources from a JSON
// response object.
func flattenJobTemplateTemplateContainersResourcesMap(c *Client, i interface{}, res *Job) map[string]JobTemplateTemplateContainersResources {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTemplateTemplateContainersResources{}
	}

	if len(a) == 0 {
		return map[string]JobTemplateTemplateContainersResources{}
	}

	items := make(map[string]JobTemplateTemplateContainersResources)
	for k, item := range a {
		items[k] = *flattenJobTemplateTemplateContainersResources(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTemplateTemplateContainersResourcesSlice flattens the contents of JobTemplateTemplateContainersResources from a JSON
// response object.
func flattenJobTemplateTemplateContainersResourcesSlice(c *Client, i interface{}, res *Job) []JobTemplateTemplateContainersResources {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTemplateTemplateContainersResources{}
	}

	if len(a) == 0 {
		return []JobTemplateTemplateContainersResources{}
	}

	items := make([]JobTemplateTemplateContainersResources, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTemplateTemplateContainersResources(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTemplateTemplateContainersResources expands an instance of JobTemplateTemplateContainersResources into a JSON
// request object.
func expandJobTemplateTemplateContainersResources(c *Client, f *JobTemplateTemplateContainersResources, res *Job) (map[string]interface{}, error) {
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

// flattenJobTemplateTemplateContainersResources flattens an instance of JobTemplateTemplateContainersResources from a JSON
// response object.
func flattenJobTemplateTemplateContainersResources(c *Client, i interface{}, res *Job) *JobTemplateTemplateContainersResources {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTemplateTemplateContainersResources{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTemplateTemplateContainersResources
	}
	r.Limits = dcl.FlattenKeyValuePairs(m["limits"])
	r.CpuIdle = dcl.FlattenBool(m["cpuIdle"])

	return r
}

// expandJobTemplateTemplateContainersPortsMap expands the contents of JobTemplateTemplateContainersPorts into a JSON
// request object.
func expandJobTemplateTemplateContainersPortsMap(c *Client, f map[string]JobTemplateTemplateContainersPorts, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTemplateTemplateContainersPorts(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTemplateTemplateContainersPortsSlice expands the contents of JobTemplateTemplateContainersPorts into a JSON
// request object.
func expandJobTemplateTemplateContainersPortsSlice(c *Client, f []JobTemplateTemplateContainersPorts, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTemplateTemplateContainersPorts(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTemplateTemplateContainersPortsMap flattens the contents of JobTemplateTemplateContainersPorts from a JSON
// response object.
func flattenJobTemplateTemplateContainersPortsMap(c *Client, i interface{}, res *Job) map[string]JobTemplateTemplateContainersPorts {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTemplateTemplateContainersPorts{}
	}

	if len(a) == 0 {
		return map[string]JobTemplateTemplateContainersPorts{}
	}

	items := make(map[string]JobTemplateTemplateContainersPorts)
	for k, item := range a {
		items[k] = *flattenJobTemplateTemplateContainersPorts(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTemplateTemplateContainersPortsSlice flattens the contents of JobTemplateTemplateContainersPorts from a JSON
// response object.
func flattenJobTemplateTemplateContainersPortsSlice(c *Client, i interface{}, res *Job) []JobTemplateTemplateContainersPorts {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTemplateTemplateContainersPorts{}
	}

	if len(a) == 0 {
		return []JobTemplateTemplateContainersPorts{}
	}

	items := make([]JobTemplateTemplateContainersPorts, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTemplateTemplateContainersPorts(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTemplateTemplateContainersPorts expands an instance of JobTemplateTemplateContainersPorts into a JSON
// request object.
func expandJobTemplateTemplateContainersPorts(c *Client, f *JobTemplateTemplateContainersPorts, res *Job) (map[string]interface{}, error) {
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

// flattenJobTemplateTemplateContainersPorts flattens an instance of JobTemplateTemplateContainersPorts from a JSON
// response object.
func flattenJobTemplateTemplateContainersPorts(c *Client, i interface{}, res *Job) *JobTemplateTemplateContainersPorts {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTemplateTemplateContainersPorts{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTemplateTemplateContainersPorts
	}
	r.Name = dcl.FlattenString(m["name"])
	r.ContainerPort = dcl.FlattenInteger(m["containerPort"])

	return r
}

// expandJobTemplateTemplateContainersVolumeMountsMap expands the contents of JobTemplateTemplateContainersVolumeMounts into a JSON
// request object.
func expandJobTemplateTemplateContainersVolumeMountsMap(c *Client, f map[string]JobTemplateTemplateContainersVolumeMounts, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTemplateTemplateContainersVolumeMounts(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTemplateTemplateContainersVolumeMountsSlice expands the contents of JobTemplateTemplateContainersVolumeMounts into a JSON
// request object.
func expandJobTemplateTemplateContainersVolumeMountsSlice(c *Client, f []JobTemplateTemplateContainersVolumeMounts, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTemplateTemplateContainersVolumeMounts(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTemplateTemplateContainersVolumeMountsMap flattens the contents of JobTemplateTemplateContainersVolumeMounts from a JSON
// response object.
func flattenJobTemplateTemplateContainersVolumeMountsMap(c *Client, i interface{}, res *Job) map[string]JobTemplateTemplateContainersVolumeMounts {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTemplateTemplateContainersVolumeMounts{}
	}

	if len(a) == 0 {
		return map[string]JobTemplateTemplateContainersVolumeMounts{}
	}

	items := make(map[string]JobTemplateTemplateContainersVolumeMounts)
	for k, item := range a {
		items[k] = *flattenJobTemplateTemplateContainersVolumeMounts(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTemplateTemplateContainersVolumeMountsSlice flattens the contents of JobTemplateTemplateContainersVolumeMounts from a JSON
// response object.
func flattenJobTemplateTemplateContainersVolumeMountsSlice(c *Client, i interface{}, res *Job) []JobTemplateTemplateContainersVolumeMounts {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTemplateTemplateContainersVolumeMounts{}
	}

	if len(a) == 0 {
		return []JobTemplateTemplateContainersVolumeMounts{}
	}

	items := make([]JobTemplateTemplateContainersVolumeMounts, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTemplateTemplateContainersVolumeMounts(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTemplateTemplateContainersVolumeMounts expands an instance of JobTemplateTemplateContainersVolumeMounts into a JSON
// request object.
func expandJobTemplateTemplateContainersVolumeMounts(c *Client, f *JobTemplateTemplateContainersVolumeMounts, res *Job) (map[string]interface{}, error) {
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

// flattenJobTemplateTemplateContainersVolumeMounts flattens an instance of JobTemplateTemplateContainersVolumeMounts from a JSON
// response object.
func flattenJobTemplateTemplateContainersVolumeMounts(c *Client, i interface{}, res *Job) *JobTemplateTemplateContainersVolumeMounts {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTemplateTemplateContainersVolumeMounts{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTemplateTemplateContainersVolumeMounts
	}
	r.Name = dcl.FlattenString(m["name"])
	r.MountPath = dcl.FlattenString(m["mountPath"])

	return r
}

// expandJobTemplateTemplateVolumesMap expands the contents of JobTemplateTemplateVolumes into a JSON
// request object.
func expandJobTemplateTemplateVolumesMap(c *Client, f map[string]JobTemplateTemplateVolumes, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTemplateTemplateVolumes(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTemplateTemplateVolumesSlice expands the contents of JobTemplateTemplateVolumes into a JSON
// request object.
func expandJobTemplateTemplateVolumesSlice(c *Client, f []JobTemplateTemplateVolumes, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTemplateTemplateVolumes(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTemplateTemplateVolumesMap flattens the contents of JobTemplateTemplateVolumes from a JSON
// response object.
func flattenJobTemplateTemplateVolumesMap(c *Client, i interface{}, res *Job) map[string]JobTemplateTemplateVolumes {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTemplateTemplateVolumes{}
	}

	if len(a) == 0 {
		return map[string]JobTemplateTemplateVolumes{}
	}

	items := make(map[string]JobTemplateTemplateVolumes)
	for k, item := range a {
		items[k] = *flattenJobTemplateTemplateVolumes(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTemplateTemplateVolumesSlice flattens the contents of JobTemplateTemplateVolumes from a JSON
// response object.
func flattenJobTemplateTemplateVolumesSlice(c *Client, i interface{}, res *Job) []JobTemplateTemplateVolumes {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTemplateTemplateVolumes{}
	}

	if len(a) == 0 {
		return []JobTemplateTemplateVolumes{}
	}

	items := make([]JobTemplateTemplateVolumes, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTemplateTemplateVolumes(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTemplateTemplateVolumes expands an instance of JobTemplateTemplateVolumes into a JSON
// request object.
func expandJobTemplateTemplateVolumes(c *Client, f *JobTemplateTemplateVolumes, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandJobTemplateTemplateVolumesSecret(c, f.Secret, res); err != nil {
		return nil, fmt.Errorf("error expanding Secret into secret: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["secret"] = v
	}
	if v, err := expandJobTemplateTemplateVolumesCloudSqlInstance(c, f.CloudSqlInstance, res); err != nil {
		return nil, fmt.Errorf("error expanding CloudSqlInstance into cloudSqlInstance: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cloudSqlInstance"] = v
	}

	return m, nil
}

// flattenJobTemplateTemplateVolumes flattens an instance of JobTemplateTemplateVolumes from a JSON
// response object.
func flattenJobTemplateTemplateVolumes(c *Client, i interface{}, res *Job) *JobTemplateTemplateVolumes {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTemplateTemplateVolumes{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTemplateTemplateVolumes
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Secret = flattenJobTemplateTemplateVolumesSecret(c, m["secret"], res)
	r.CloudSqlInstance = flattenJobTemplateTemplateVolumesCloudSqlInstance(c, m["cloudSqlInstance"], res)

	return r
}

// expandJobTemplateTemplateVolumesSecretMap expands the contents of JobTemplateTemplateVolumesSecret into a JSON
// request object.
func expandJobTemplateTemplateVolumesSecretMap(c *Client, f map[string]JobTemplateTemplateVolumesSecret, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTemplateTemplateVolumesSecret(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTemplateTemplateVolumesSecretSlice expands the contents of JobTemplateTemplateVolumesSecret into a JSON
// request object.
func expandJobTemplateTemplateVolumesSecretSlice(c *Client, f []JobTemplateTemplateVolumesSecret, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTemplateTemplateVolumesSecret(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTemplateTemplateVolumesSecretMap flattens the contents of JobTemplateTemplateVolumesSecret from a JSON
// response object.
func flattenJobTemplateTemplateVolumesSecretMap(c *Client, i interface{}, res *Job) map[string]JobTemplateTemplateVolumesSecret {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTemplateTemplateVolumesSecret{}
	}

	if len(a) == 0 {
		return map[string]JobTemplateTemplateVolumesSecret{}
	}

	items := make(map[string]JobTemplateTemplateVolumesSecret)
	for k, item := range a {
		items[k] = *flattenJobTemplateTemplateVolumesSecret(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTemplateTemplateVolumesSecretSlice flattens the contents of JobTemplateTemplateVolumesSecret from a JSON
// response object.
func flattenJobTemplateTemplateVolumesSecretSlice(c *Client, i interface{}, res *Job) []JobTemplateTemplateVolumesSecret {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTemplateTemplateVolumesSecret{}
	}

	if len(a) == 0 {
		return []JobTemplateTemplateVolumesSecret{}
	}

	items := make([]JobTemplateTemplateVolumesSecret, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTemplateTemplateVolumesSecret(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTemplateTemplateVolumesSecret expands an instance of JobTemplateTemplateVolumesSecret into a JSON
// request object.
func expandJobTemplateTemplateVolumesSecret(c *Client, f *JobTemplateTemplateVolumesSecret, res *Job) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Secret; !dcl.IsEmptyValueIndirect(v) {
		m["secret"] = v
	}
	if v, err := expandJobTemplateTemplateVolumesSecretItemsSlice(c, f.Items, res); err != nil {
		return nil, fmt.Errorf("error expanding Items into items: %w", err)
	} else if v != nil {
		m["items"] = v
	}
	if v := f.DefaultMode; !dcl.IsEmptyValueIndirect(v) {
		m["defaultMode"] = v
	}

	return m, nil
}

// flattenJobTemplateTemplateVolumesSecret flattens an instance of JobTemplateTemplateVolumesSecret from a JSON
// response object.
func flattenJobTemplateTemplateVolumesSecret(c *Client, i interface{}, res *Job) *JobTemplateTemplateVolumesSecret {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTemplateTemplateVolumesSecret{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTemplateTemplateVolumesSecret
	}
	r.Secret = dcl.FlattenString(m["secret"])
	r.Items = flattenJobTemplateTemplateVolumesSecretItemsSlice(c, m["items"], res)
	r.DefaultMode = dcl.FlattenInteger(m["defaultMode"])

	return r
}

// expandJobTemplateTemplateVolumesSecretItemsMap expands the contents of JobTemplateTemplateVolumesSecretItems into a JSON
// request object.
func expandJobTemplateTemplateVolumesSecretItemsMap(c *Client, f map[string]JobTemplateTemplateVolumesSecretItems, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTemplateTemplateVolumesSecretItems(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTemplateTemplateVolumesSecretItemsSlice expands the contents of JobTemplateTemplateVolumesSecretItems into a JSON
// request object.
func expandJobTemplateTemplateVolumesSecretItemsSlice(c *Client, f []JobTemplateTemplateVolumesSecretItems, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTemplateTemplateVolumesSecretItems(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTemplateTemplateVolumesSecretItemsMap flattens the contents of JobTemplateTemplateVolumesSecretItems from a JSON
// response object.
func flattenJobTemplateTemplateVolumesSecretItemsMap(c *Client, i interface{}, res *Job) map[string]JobTemplateTemplateVolumesSecretItems {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTemplateTemplateVolumesSecretItems{}
	}

	if len(a) == 0 {
		return map[string]JobTemplateTemplateVolumesSecretItems{}
	}

	items := make(map[string]JobTemplateTemplateVolumesSecretItems)
	for k, item := range a {
		items[k] = *flattenJobTemplateTemplateVolumesSecretItems(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTemplateTemplateVolumesSecretItemsSlice flattens the contents of JobTemplateTemplateVolumesSecretItems from a JSON
// response object.
func flattenJobTemplateTemplateVolumesSecretItemsSlice(c *Client, i interface{}, res *Job) []JobTemplateTemplateVolumesSecretItems {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTemplateTemplateVolumesSecretItems{}
	}

	if len(a) == 0 {
		return []JobTemplateTemplateVolumesSecretItems{}
	}

	items := make([]JobTemplateTemplateVolumesSecretItems, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTemplateTemplateVolumesSecretItems(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTemplateTemplateVolumesSecretItems expands an instance of JobTemplateTemplateVolumesSecretItems into a JSON
// request object.
func expandJobTemplateTemplateVolumesSecretItems(c *Client, f *JobTemplateTemplateVolumesSecretItems, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Path; !dcl.IsEmptyValueIndirect(v) {
		m["path"] = v
	}
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}
	if v := f.Mode; !dcl.IsEmptyValueIndirect(v) {
		m["mode"] = v
	}

	return m, nil
}

// flattenJobTemplateTemplateVolumesSecretItems flattens an instance of JobTemplateTemplateVolumesSecretItems from a JSON
// response object.
func flattenJobTemplateTemplateVolumesSecretItems(c *Client, i interface{}, res *Job) *JobTemplateTemplateVolumesSecretItems {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTemplateTemplateVolumesSecretItems{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTemplateTemplateVolumesSecretItems
	}
	r.Path = dcl.FlattenString(m["path"])
	r.Version = dcl.FlattenString(m["version"])
	r.Mode = dcl.FlattenInteger(m["mode"])

	return r
}

// expandJobTemplateTemplateVolumesCloudSqlInstanceMap expands the contents of JobTemplateTemplateVolumesCloudSqlInstance into a JSON
// request object.
func expandJobTemplateTemplateVolumesCloudSqlInstanceMap(c *Client, f map[string]JobTemplateTemplateVolumesCloudSqlInstance, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTemplateTemplateVolumesCloudSqlInstance(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTemplateTemplateVolumesCloudSqlInstanceSlice expands the contents of JobTemplateTemplateVolumesCloudSqlInstance into a JSON
// request object.
func expandJobTemplateTemplateVolumesCloudSqlInstanceSlice(c *Client, f []JobTemplateTemplateVolumesCloudSqlInstance, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTemplateTemplateVolumesCloudSqlInstance(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTemplateTemplateVolumesCloudSqlInstanceMap flattens the contents of JobTemplateTemplateVolumesCloudSqlInstance from a JSON
// response object.
func flattenJobTemplateTemplateVolumesCloudSqlInstanceMap(c *Client, i interface{}, res *Job) map[string]JobTemplateTemplateVolumesCloudSqlInstance {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTemplateTemplateVolumesCloudSqlInstance{}
	}

	if len(a) == 0 {
		return map[string]JobTemplateTemplateVolumesCloudSqlInstance{}
	}

	items := make(map[string]JobTemplateTemplateVolumesCloudSqlInstance)
	for k, item := range a {
		items[k] = *flattenJobTemplateTemplateVolumesCloudSqlInstance(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTemplateTemplateVolumesCloudSqlInstanceSlice flattens the contents of JobTemplateTemplateVolumesCloudSqlInstance from a JSON
// response object.
func flattenJobTemplateTemplateVolumesCloudSqlInstanceSlice(c *Client, i interface{}, res *Job) []JobTemplateTemplateVolumesCloudSqlInstance {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTemplateTemplateVolumesCloudSqlInstance{}
	}

	if len(a) == 0 {
		return []JobTemplateTemplateVolumesCloudSqlInstance{}
	}

	items := make([]JobTemplateTemplateVolumesCloudSqlInstance, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTemplateTemplateVolumesCloudSqlInstance(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTemplateTemplateVolumesCloudSqlInstance expands an instance of JobTemplateTemplateVolumesCloudSqlInstance into a JSON
// request object.
func expandJobTemplateTemplateVolumesCloudSqlInstance(c *Client, f *JobTemplateTemplateVolumesCloudSqlInstance, res *Job) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Instances; v != nil {
		m["instances"] = v
	}

	return m, nil
}

// flattenJobTemplateTemplateVolumesCloudSqlInstance flattens an instance of JobTemplateTemplateVolumesCloudSqlInstance from a JSON
// response object.
func flattenJobTemplateTemplateVolumesCloudSqlInstance(c *Client, i interface{}, res *Job) *JobTemplateTemplateVolumesCloudSqlInstance {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTemplateTemplateVolumesCloudSqlInstance{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTemplateTemplateVolumesCloudSqlInstance
	}
	r.Instances = dcl.FlattenStringSlice(m["instances"])

	return r
}

// expandJobTemplateTemplateVPCAccessMap expands the contents of JobTemplateTemplateVPCAccess into a JSON
// request object.
func expandJobTemplateTemplateVPCAccessMap(c *Client, f map[string]JobTemplateTemplateVPCAccess, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTemplateTemplateVPCAccess(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTemplateTemplateVPCAccessSlice expands the contents of JobTemplateTemplateVPCAccess into a JSON
// request object.
func expandJobTemplateTemplateVPCAccessSlice(c *Client, f []JobTemplateTemplateVPCAccess, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTemplateTemplateVPCAccess(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTemplateTemplateVPCAccessMap flattens the contents of JobTemplateTemplateVPCAccess from a JSON
// response object.
func flattenJobTemplateTemplateVPCAccessMap(c *Client, i interface{}, res *Job) map[string]JobTemplateTemplateVPCAccess {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTemplateTemplateVPCAccess{}
	}

	if len(a) == 0 {
		return map[string]JobTemplateTemplateVPCAccess{}
	}

	items := make(map[string]JobTemplateTemplateVPCAccess)
	for k, item := range a {
		items[k] = *flattenJobTemplateTemplateVPCAccess(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTemplateTemplateVPCAccessSlice flattens the contents of JobTemplateTemplateVPCAccess from a JSON
// response object.
func flattenJobTemplateTemplateVPCAccessSlice(c *Client, i interface{}, res *Job) []JobTemplateTemplateVPCAccess {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTemplateTemplateVPCAccess{}
	}

	if len(a) == 0 {
		return []JobTemplateTemplateVPCAccess{}
	}

	items := make([]JobTemplateTemplateVPCAccess, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTemplateTemplateVPCAccess(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTemplateTemplateVPCAccess expands an instance of JobTemplateTemplateVPCAccess into a JSON
// request object.
func expandJobTemplateTemplateVPCAccess(c *Client, f *JobTemplateTemplateVPCAccess, res *Job) (map[string]interface{}, error) {
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

// flattenJobTemplateTemplateVPCAccess flattens an instance of JobTemplateTemplateVPCAccess from a JSON
// response object.
func flattenJobTemplateTemplateVPCAccess(c *Client, i interface{}, res *Job) *JobTemplateTemplateVPCAccess {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTemplateTemplateVPCAccess{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTemplateTemplateVPCAccess
	}
	r.Connector = dcl.FlattenString(m["connector"])
	r.Egress = flattenJobTemplateTemplateVPCAccessEgressEnum(m["egress"])

	return r
}

// expandJobTerminalConditionMap expands the contents of JobTerminalCondition into a JSON
// request object.
func expandJobTerminalConditionMap(c *Client, f map[string]JobTerminalCondition, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTerminalCondition(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTerminalConditionSlice expands the contents of JobTerminalCondition into a JSON
// request object.
func expandJobTerminalConditionSlice(c *Client, f []JobTerminalCondition, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTerminalCondition(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTerminalConditionMap flattens the contents of JobTerminalCondition from a JSON
// response object.
func flattenJobTerminalConditionMap(c *Client, i interface{}, res *Job) map[string]JobTerminalCondition {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTerminalCondition{}
	}

	if len(a) == 0 {
		return map[string]JobTerminalCondition{}
	}

	items := make(map[string]JobTerminalCondition)
	for k, item := range a {
		items[k] = *flattenJobTerminalCondition(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTerminalConditionSlice flattens the contents of JobTerminalCondition from a JSON
// response object.
func flattenJobTerminalConditionSlice(c *Client, i interface{}, res *Job) []JobTerminalCondition {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTerminalCondition{}
	}

	if len(a) == 0 {
		return []JobTerminalCondition{}
	}

	items := make([]JobTerminalCondition, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTerminalCondition(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTerminalCondition expands an instance of JobTerminalCondition into a JSON
// request object.
func expandJobTerminalCondition(c *Client, f *JobTerminalCondition, res *Job) (map[string]interface{}, error) {
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
	if v := f.InternalReason; !dcl.IsEmptyValueIndirect(v) {
		m["internalReason"] = v
	}
	if v := f.DomainMappingReason; !dcl.IsEmptyValueIndirect(v) {
		m["domainMappingReason"] = v
	}
	if v := f.RevisionReason; !dcl.IsEmptyValueIndirect(v) {
		m["revisionReason"] = v
	}
	if v := f.ExecutionReason; !dcl.IsEmptyValueIndirect(v) {
		m["executionReason"] = v
	}

	return m, nil
}

// flattenJobTerminalCondition flattens an instance of JobTerminalCondition from a JSON
// response object.
func flattenJobTerminalCondition(c *Client, i interface{}, res *Job) *JobTerminalCondition {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTerminalCondition{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTerminalCondition
	}
	r.Type = dcl.FlattenString(m["type"])
	r.State = flattenJobTerminalConditionStateEnum(m["state"])
	r.Message = dcl.FlattenString(m["message"])
	r.LastTransitionTime = dcl.FlattenString(m["lastTransitionTime"])
	r.Severity = flattenJobTerminalConditionSeverityEnum(m["severity"])
	r.Reason = flattenJobTerminalConditionReasonEnum(m["reason"])
	r.InternalReason = flattenJobTerminalConditionInternalReasonEnum(m["internalReason"])
	r.DomainMappingReason = flattenJobTerminalConditionDomainMappingReasonEnum(m["domainMappingReason"])
	r.RevisionReason = flattenJobTerminalConditionRevisionReasonEnum(m["revisionReason"])
	r.ExecutionReason = flattenJobTerminalConditionExecutionReasonEnum(m["executionReason"])

	return r
}

// expandJobConditionsMap expands the contents of JobConditions into a JSON
// request object.
func expandJobConditionsMap(c *Client, f map[string]JobConditions, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobConditions(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobConditionsSlice expands the contents of JobConditions into a JSON
// request object.
func expandJobConditionsSlice(c *Client, f []JobConditions, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobConditions(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobConditionsMap flattens the contents of JobConditions from a JSON
// response object.
func flattenJobConditionsMap(c *Client, i interface{}, res *Job) map[string]JobConditions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobConditions{}
	}

	if len(a) == 0 {
		return map[string]JobConditions{}
	}

	items := make(map[string]JobConditions)
	for k, item := range a {
		items[k] = *flattenJobConditions(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobConditionsSlice flattens the contents of JobConditions from a JSON
// response object.
func flattenJobConditionsSlice(c *Client, i interface{}, res *Job) []JobConditions {
	a, ok := i.([]interface{})
	if !ok {
		return []JobConditions{}
	}

	if len(a) == 0 {
		return []JobConditions{}
	}

	items := make([]JobConditions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobConditions(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobConditions expands an instance of JobConditions into a JSON
// request object.
func expandJobConditions(c *Client, f *JobConditions, res *Job) (map[string]interface{}, error) {
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
	if v := f.ExecutionReason; !dcl.IsEmptyValueIndirect(v) {
		m["executionReason"] = v
	}

	return m, nil
}

// flattenJobConditions flattens an instance of JobConditions from a JSON
// response object.
func flattenJobConditions(c *Client, i interface{}, res *Job) *JobConditions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobConditions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobConditions
	}
	r.Type = dcl.FlattenString(m["type"])
	r.State = flattenJobConditionsStateEnum(m["state"])
	r.Message = dcl.FlattenString(m["message"])
	r.LastTransitionTime = dcl.FlattenString(m["lastTransitionTime"])
	r.Severity = flattenJobConditionsSeverityEnum(m["severity"])
	r.Reason = flattenJobConditionsReasonEnum(m["reason"])
	r.RevisionReason = flattenJobConditionsRevisionReasonEnum(m["revisionReason"])
	r.ExecutionReason = flattenJobConditionsExecutionReasonEnum(m["executionReason"])

	return r
}

// expandJobLatestSucceededExecutionMap expands the contents of JobLatestSucceededExecution into a JSON
// request object.
func expandJobLatestSucceededExecutionMap(c *Client, f map[string]JobLatestSucceededExecution, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobLatestSucceededExecution(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobLatestSucceededExecutionSlice expands the contents of JobLatestSucceededExecution into a JSON
// request object.
func expandJobLatestSucceededExecutionSlice(c *Client, f []JobLatestSucceededExecution, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobLatestSucceededExecution(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobLatestSucceededExecutionMap flattens the contents of JobLatestSucceededExecution from a JSON
// response object.
func flattenJobLatestSucceededExecutionMap(c *Client, i interface{}, res *Job) map[string]JobLatestSucceededExecution {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobLatestSucceededExecution{}
	}

	if len(a) == 0 {
		return map[string]JobLatestSucceededExecution{}
	}

	items := make(map[string]JobLatestSucceededExecution)
	for k, item := range a {
		items[k] = *flattenJobLatestSucceededExecution(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobLatestSucceededExecutionSlice flattens the contents of JobLatestSucceededExecution from a JSON
// response object.
func flattenJobLatestSucceededExecutionSlice(c *Client, i interface{}, res *Job) []JobLatestSucceededExecution {
	a, ok := i.([]interface{})
	if !ok {
		return []JobLatestSucceededExecution{}
	}

	if len(a) == 0 {
		return []JobLatestSucceededExecution{}
	}

	items := make([]JobLatestSucceededExecution, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobLatestSucceededExecution(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobLatestSucceededExecution expands an instance of JobLatestSucceededExecution into a JSON
// request object.
func expandJobLatestSucceededExecution(c *Client, f *JobLatestSucceededExecution, res *Job) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}

	return m, nil
}

// flattenJobLatestSucceededExecution flattens an instance of JobLatestSucceededExecution from a JSON
// response object.
func flattenJobLatestSucceededExecution(c *Client, i interface{}, res *Job) *JobLatestSucceededExecution {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobLatestSucceededExecution{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobLatestSucceededExecution
	}
	r.Name = dcl.FlattenString(m["name"])
	r.CreateTime = dcl.FlattenString(m["createTime"])

	return r
}

// expandJobLatestCreatedExecutionMap expands the contents of JobLatestCreatedExecution into a JSON
// request object.
func expandJobLatestCreatedExecutionMap(c *Client, f map[string]JobLatestCreatedExecution, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobLatestCreatedExecution(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobLatestCreatedExecutionSlice expands the contents of JobLatestCreatedExecution into a JSON
// request object.
func expandJobLatestCreatedExecutionSlice(c *Client, f []JobLatestCreatedExecution, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobLatestCreatedExecution(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobLatestCreatedExecutionMap flattens the contents of JobLatestCreatedExecution from a JSON
// response object.
func flattenJobLatestCreatedExecutionMap(c *Client, i interface{}, res *Job) map[string]JobLatestCreatedExecution {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobLatestCreatedExecution{}
	}

	if len(a) == 0 {
		return map[string]JobLatestCreatedExecution{}
	}

	items := make(map[string]JobLatestCreatedExecution)
	for k, item := range a {
		items[k] = *flattenJobLatestCreatedExecution(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobLatestCreatedExecutionSlice flattens the contents of JobLatestCreatedExecution from a JSON
// response object.
func flattenJobLatestCreatedExecutionSlice(c *Client, i interface{}, res *Job) []JobLatestCreatedExecution {
	a, ok := i.([]interface{})
	if !ok {
		return []JobLatestCreatedExecution{}
	}

	if len(a) == 0 {
		return []JobLatestCreatedExecution{}
	}

	items := make([]JobLatestCreatedExecution, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobLatestCreatedExecution(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobLatestCreatedExecution expands an instance of JobLatestCreatedExecution into a JSON
// request object.
func expandJobLatestCreatedExecution(c *Client, f *JobLatestCreatedExecution, res *Job) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}

	return m, nil
}

// flattenJobLatestCreatedExecution flattens an instance of JobLatestCreatedExecution from a JSON
// response object.
func flattenJobLatestCreatedExecution(c *Client, i interface{}, res *Job) *JobLatestCreatedExecution {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobLatestCreatedExecution{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobLatestCreatedExecution
	}
	r.Name = dcl.FlattenString(m["name"])
	r.CreateTime = dcl.FlattenString(m["createTime"])

	return r
}

// flattenJobLaunchStageEnumMap flattens the contents of JobLaunchStageEnum from a JSON
// response object.
func flattenJobLaunchStageEnumMap(c *Client, i interface{}, res *Job) map[string]JobLaunchStageEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobLaunchStageEnum{}
	}

	if len(a) == 0 {
		return map[string]JobLaunchStageEnum{}
	}

	items := make(map[string]JobLaunchStageEnum)
	for k, item := range a {
		items[k] = *flattenJobLaunchStageEnum(item.(interface{}))
	}

	return items
}

// flattenJobLaunchStageEnumSlice flattens the contents of JobLaunchStageEnum from a JSON
// response object.
func flattenJobLaunchStageEnumSlice(c *Client, i interface{}, res *Job) []JobLaunchStageEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobLaunchStageEnum{}
	}

	if len(a) == 0 {
		return []JobLaunchStageEnum{}
	}

	items := make([]JobLaunchStageEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobLaunchStageEnum(item.(interface{})))
	}

	return items
}

// flattenJobLaunchStageEnum asserts that an interface is a string, and returns a
// pointer to a *JobLaunchStageEnum with the same value as that string.
func flattenJobLaunchStageEnum(i interface{}) *JobLaunchStageEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobLaunchStageEnumRef(s)
}

// flattenJobTemplateTemplateExecutionEnvironmentEnumMap flattens the contents of JobTemplateTemplateExecutionEnvironmentEnum from a JSON
// response object.
func flattenJobTemplateTemplateExecutionEnvironmentEnumMap(c *Client, i interface{}, res *Job) map[string]JobTemplateTemplateExecutionEnvironmentEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTemplateTemplateExecutionEnvironmentEnum{}
	}

	if len(a) == 0 {
		return map[string]JobTemplateTemplateExecutionEnvironmentEnum{}
	}

	items := make(map[string]JobTemplateTemplateExecutionEnvironmentEnum)
	for k, item := range a {
		items[k] = *flattenJobTemplateTemplateExecutionEnvironmentEnum(item.(interface{}))
	}

	return items
}

// flattenJobTemplateTemplateExecutionEnvironmentEnumSlice flattens the contents of JobTemplateTemplateExecutionEnvironmentEnum from a JSON
// response object.
func flattenJobTemplateTemplateExecutionEnvironmentEnumSlice(c *Client, i interface{}, res *Job) []JobTemplateTemplateExecutionEnvironmentEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTemplateTemplateExecutionEnvironmentEnum{}
	}

	if len(a) == 0 {
		return []JobTemplateTemplateExecutionEnvironmentEnum{}
	}

	items := make([]JobTemplateTemplateExecutionEnvironmentEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTemplateTemplateExecutionEnvironmentEnum(item.(interface{})))
	}

	return items
}

// flattenJobTemplateTemplateExecutionEnvironmentEnum asserts that an interface is a string, and returns a
// pointer to a *JobTemplateTemplateExecutionEnvironmentEnum with the same value as that string.
func flattenJobTemplateTemplateExecutionEnvironmentEnum(i interface{}) *JobTemplateTemplateExecutionEnvironmentEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobTemplateTemplateExecutionEnvironmentEnumRef(s)
}

// flattenJobTemplateTemplateVPCAccessEgressEnumMap flattens the contents of JobTemplateTemplateVPCAccessEgressEnum from a JSON
// response object.
func flattenJobTemplateTemplateVPCAccessEgressEnumMap(c *Client, i interface{}, res *Job) map[string]JobTemplateTemplateVPCAccessEgressEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTemplateTemplateVPCAccessEgressEnum{}
	}

	if len(a) == 0 {
		return map[string]JobTemplateTemplateVPCAccessEgressEnum{}
	}

	items := make(map[string]JobTemplateTemplateVPCAccessEgressEnum)
	for k, item := range a {
		items[k] = *flattenJobTemplateTemplateVPCAccessEgressEnum(item.(interface{}))
	}

	return items
}

// flattenJobTemplateTemplateVPCAccessEgressEnumSlice flattens the contents of JobTemplateTemplateVPCAccessEgressEnum from a JSON
// response object.
func flattenJobTemplateTemplateVPCAccessEgressEnumSlice(c *Client, i interface{}, res *Job) []JobTemplateTemplateVPCAccessEgressEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTemplateTemplateVPCAccessEgressEnum{}
	}

	if len(a) == 0 {
		return []JobTemplateTemplateVPCAccessEgressEnum{}
	}

	items := make([]JobTemplateTemplateVPCAccessEgressEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTemplateTemplateVPCAccessEgressEnum(item.(interface{})))
	}

	return items
}

// flattenJobTemplateTemplateVPCAccessEgressEnum asserts that an interface is a string, and returns a
// pointer to a *JobTemplateTemplateVPCAccessEgressEnum with the same value as that string.
func flattenJobTemplateTemplateVPCAccessEgressEnum(i interface{}) *JobTemplateTemplateVPCAccessEgressEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobTemplateTemplateVPCAccessEgressEnumRef(s)
}

// flattenJobTerminalConditionStateEnumMap flattens the contents of JobTerminalConditionStateEnum from a JSON
// response object.
func flattenJobTerminalConditionStateEnumMap(c *Client, i interface{}, res *Job) map[string]JobTerminalConditionStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTerminalConditionStateEnum{}
	}

	if len(a) == 0 {
		return map[string]JobTerminalConditionStateEnum{}
	}

	items := make(map[string]JobTerminalConditionStateEnum)
	for k, item := range a {
		items[k] = *flattenJobTerminalConditionStateEnum(item.(interface{}))
	}

	return items
}

// flattenJobTerminalConditionStateEnumSlice flattens the contents of JobTerminalConditionStateEnum from a JSON
// response object.
func flattenJobTerminalConditionStateEnumSlice(c *Client, i interface{}, res *Job) []JobTerminalConditionStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTerminalConditionStateEnum{}
	}

	if len(a) == 0 {
		return []JobTerminalConditionStateEnum{}
	}

	items := make([]JobTerminalConditionStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTerminalConditionStateEnum(item.(interface{})))
	}

	return items
}

// flattenJobTerminalConditionStateEnum asserts that an interface is a string, and returns a
// pointer to a *JobTerminalConditionStateEnum with the same value as that string.
func flattenJobTerminalConditionStateEnum(i interface{}) *JobTerminalConditionStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobTerminalConditionStateEnumRef(s)
}

// flattenJobTerminalConditionSeverityEnumMap flattens the contents of JobTerminalConditionSeverityEnum from a JSON
// response object.
func flattenJobTerminalConditionSeverityEnumMap(c *Client, i interface{}, res *Job) map[string]JobTerminalConditionSeverityEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTerminalConditionSeverityEnum{}
	}

	if len(a) == 0 {
		return map[string]JobTerminalConditionSeverityEnum{}
	}

	items := make(map[string]JobTerminalConditionSeverityEnum)
	for k, item := range a {
		items[k] = *flattenJobTerminalConditionSeverityEnum(item.(interface{}))
	}

	return items
}

// flattenJobTerminalConditionSeverityEnumSlice flattens the contents of JobTerminalConditionSeverityEnum from a JSON
// response object.
func flattenJobTerminalConditionSeverityEnumSlice(c *Client, i interface{}, res *Job) []JobTerminalConditionSeverityEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTerminalConditionSeverityEnum{}
	}

	if len(a) == 0 {
		return []JobTerminalConditionSeverityEnum{}
	}

	items := make([]JobTerminalConditionSeverityEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTerminalConditionSeverityEnum(item.(interface{})))
	}

	return items
}

// flattenJobTerminalConditionSeverityEnum asserts that an interface is a string, and returns a
// pointer to a *JobTerminalConditionSeverityEnum with the same value as that string.
func flattenJobTerminalConditionSeverityEnum(i interface{}) *JobTerminalConditionSeverityEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobTerminalConditionSeverityEnumRef(s)
}

// flattenJobTerminalConditionReasonEnumMap flattens the contents of JobTerminalConditionReasonEnum from a JSON
// response object.
func flattenJobTerminalConditionReasonEnumMap(c *Client, i interface{}, res *Job) map[string]JobTerminalConditionReasonEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTerminalConditionReasonEnum{}
	}

	if len(a) == 0 {
		return map[string]JobTerminalConditionReasonEnum{}
	}

	items := make(map[string]JobTerminalConditionReasonEnum)
	for k, item := range a {
		items[k] = *flattenJobTerminalConditionReasonEnum(item.(interface{}))
	}

	return items
}

// flattenJobTerminalConditionReasonEnumSlice flattens the contents of JobTerminalConditionReasonEnum from a JSON
// response object.
func flattenJobTerminalConditionReasonEnumSlice(c *Client, i interface{}, res *Job) []JobTerminalConditionReasonEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTerminalConditionReasonEnum{}
	}

	if len(a) == 0 {
		return []JobTerminalConditionReasonEnum{}
	}

	items := make([]JobTerminalConditionReasonEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTerminalConditionReasonEnum(item.(interface{})))
	}

	return items
}

// flattenJobTerminalConditionReasonEnum asserts that an interface is a string, and returns a
// pointer to a *JobTerminalConditionReasonEnum with the same value as that string.
func flattenJobTerminalConditionReasonEnum(i interface{}) *JobTerminalConditionReasonEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobTerminalConditionReasonEnumRef(s)
}

// flattenJobTerminalConditionInternalReasonEnumMap flattens the contents of JobTerminalConditionInternalReasonEnum from a JSON
// response object.
func flattenJobTerminalConditionInternalReasonEnumMap(c *Client, i interface{}, res *Job) map[string]JobTerminalConditionInternalReasonEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTerminalConditionInternalReasonEnum{}
	}

	if len(a) == 0 {
		return map[string]JobTerminalConditionInternalReasonEnum{}
	}

	items := make(map[string]JobTerminalConditionInternalReasonEnum)
	for k, item := range a {
		items[k] = *flattenJobTerminalConditionInternalReasonEnum(item.(interface{}))
	}

	return items
}

// flattenJobTerminalConditionInternalReasonEnumSlice flattens the contents of JobTerminalConditionInternalReasonEnum from a JSON
// response object.
func flattenJobTerminalConditionInternalReasonEnumSlice(c *Client, i interface{}, res *Job) []JobTerminalConditionInternalReasonEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTerminalConditionInternalReasonEnum{}
	}

	if len(a) == 0 {
		return []JobTerminalConditionInternalReasonEnum{}
	}

	items := make([]JobTerminalConditionInternalReasonEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTerminalConditionInternalReasonEnum(item.(interface{})))
	}

	return items
}

// flattenJobTerminalConditionInternalReasonEnum asserts that an interface is a string, and returns a
// pointer to a *JobTerminalConditionInternalReasonEnum with the same value as that string.
func flattenJobTerminalConditionInternalReasonEnum(i interface{}) *JobTerminalConditionInternalReasonEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobTerminalConditionInternalReasonEnumRef(s)
}

// flattenJobTerminalConditionDomainMappingReasonEnumMap flattens the contents of JobTerminalConditionDomainMappingReasonEnum from a JSON
// response object.
func flattenJobTerminalConditionDomainMappingReasonEnumMap(c *Client, i interface{}, res *Job) map[string]JobTerminalConditionDomainMappingReasonEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTerminalConditionDomainMappingReasonEnum{}
	}

	if len(a) == 0 {
		return map[string]JobTerminalConditionDomainMappingReasonEnum{}
	}

	items := make(map[string]JobTerminalConditionDomainMappingReasonEnum)
	for k, item := range a {
		items[k] = *flattenJobTerminalConditionDomainMappingReasonEnum(item.(interface{}))
	}

	return items
}

// flattenJobTerminalConditionDomainMappingReasonEnumSlice flattens the contents of JobTerminalConditionDomainMappingReasonEnum from a JSON
// response object.
func flattenJobTerminalConditionDomainMappingReasonEnumSlice(c *Client, i interface{}, res *Job) []JobTerminalConditionDomainMappingReasonEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTerminalConditionDomainMappingReasonEnum{}
	}

	if len(a) == 0 {
		return []JobTerminalConditionDomainMappingReasonEnum{}
	}

	items := make([]JobTerminalConditionDomainMappingReasonEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTerminalConditionDomainMappingReasonEnum(item.(interface{})))
	}

	return items
}

// flattenJobTerminalConditionDomainMappingReasonEnum asserts that an interface is a string, and returns a
// pointer to a *JobTerminalConditionDomainMappingReasonEnum with the same value as that string.
func flattenJobTerminalConditionDomainMappingReasonEnum(i interface{}) *JobTerminalConditionDomainMappingReasonEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobTerminalConditionDomainMappingReasonEnumRef(s)
}

// flattenJobTerminalConditionRevisionReasonEnumMap flattens the contents of JobTerminalConditionRevisionReasonEnum from a JSON
// response object.
func flattenJobTerminalConditionRevisionReasonEnumMap(c *Client, i interface{}, res *Job) map[string]JobTerminalConditionRevisionReasonEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTerminalConditionRevisionReasonEnum{}
	}

	if len(a) == 0 {
		return map[string]JobTerminalConditionRevisionReasonEnum{}
	}

	items := make(map[string]JobTerminalConditionRevisionReasonEnum)
	for k, item := range a {
		items[k] = *flattenJobTerminalConditionRevisionReasonEnum(item.(interface{}))
	}

	return items
}

// flattenJobTerminalConditionRevisionReasonEnumSlice flattens the contents of JobTerminalConditionRevisionReasonEnum from a JSON
// response object.
func flattenJobTerminalConditionRevisionReasonEnumSlice(c *Client, i interface{}, res *Job) []JobTerminalConditionRevisionReasonEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTerminalConditionRevisionReasonEnum{}
	}

	if len(a) == 0 {
		return []JobTerminalConditionRevisionReasonEnum{}
	}

	items := make([]JobTerminalConditionRevisionReasonEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTerminalConditionRevisionReasonEnum(item.(interface{})))
	}

	return items
}

// flattenJobTerminalConditionRevisionReasonEnum asserts that an interface is a string, and returns a
// pointer to a *JobTerminalConditionRevisionReasonEnum with the same value as that string.
func flattenJobTerminalConditionRevisionReasonEnum(i interface{}) *JobTerminalConditionRevisionReasonEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobTerminalConditionRevisionReasonEnumRef(s)
}

// flattenJobTerminalConditionExecutionReasonEnumMap flattens the contents of JobTerminalConditionExecutionReasonEnum from a JSON
// response object.
func flattenJobTerminalConditionExecutionReasonEnumMap(c *Client, i interface{}, res *Job) map[string]JobTerminalConditionExecutionReasonEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTerminalConditionExecutionReasonEnum{}
	}

	if len(a) == 0 {
		return map[string]JobTerminalConditionExecutionReasonEnum{}
	}

	items := make(map[string]JobTerminalConditionExecutionReasonEnum)
	for k, item := range a {
		items[k] = *flattenJobTerminalConditionExecutionReasonEnum(item.(interface{}))
	}

	return items
}

// flattenJobTerminalConditionExecutionReasonEnumSlice flattens the contents of JobTerminalConditionExecutionReasonEnum from a JSON
// response object.
func flattenJobTerminalConditionExecutionReasonEnumSlice(c *Client, i interface{}, res *Job) []JobTerminalConditionExecutionReasonEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTerminalConditionExecutionReasonEnum{}
	}

	if len(a) == 0 {
		return []JobTerminalConditionExecutionReasonEnum{}
	}

	items := make([]JobTerminalConditionExecutionReasonEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTerminalConditionExecutionReasonEnum(item.(interface{})))
	}

	return items
}

// flattenJobTerminalConditionExecutionReasonEnum asserts that an interface is a string, and returns a
// pointer to a *JobTerminalConditionExecutionReasonEnum with the same value as that string.
func flattenJobTerminalConditionExecutionReasonEnum(i interface{}) *JobTerminalConditionExecutionReasonEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobTerminalConditionExecutionReasonEnumRef(s)
}

// flattenJobConditionsStateEnumMap flattens the contents of JobConditionsStateEnum from a JSON
// response object.
func flattenJobConditionsStateEnumMap(c *Client, i interface{}, res *Job) map[string]JobConditionsStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobConditionsStateEnum{}
	}

	if len(a) == 0 {
		return map[string]JobConditionsStateEnum{}
	}

	items := make(map[string]JobConditionsStateEnum)
	for k, item := range a {
		items[k] = *flattenJobConditionsStateEnum(item.(interface{}))
	}

	return items
}

// flattenJobConditionsStateEnumSlice flattens the contents of JobConditionsStateEnum from a JSON
// response object.
func flattenJobConditionsStateEnumSlice(c *Client, i interface{}, res *Job) []JobConditionsStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobConditionsStateEnum{}
	}

	if len(a) == 0 {
		return []JobConditionsStateEnum{}
	}

	items := make([]JobConditionsStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobConditionsStateEnum(item.(interface{})))
	}

	return items
}

// flattenJobConditionsStateEnum asserts that an interface is a string, and returns a
// pointer to a *JobConditionsStateEnum with the same value as that string.
func flattenJobConditionsStateEnum(i interface{}) *JobConditionsStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobConditionsStateEnumRef(s)
}

// flattenJobConditionsSeverityEnumMap flattens the contents of JobConditionsSeverityEnum from a JSON
// response object.
func flattenJobConditionsSeverityEnumMap(c *Client, i interface{}, res *Job) map[string]JobConditionsSeverityEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobConditionsSeverityEnum{}
	}

	if len(a) == 0 {
		return map[string]JobConditionsSeverityEnum{}
	}

	items := make(map[string]JobConditionsSeverityEnum)
	for k, item := range a {
		items[k] = *flattenJobConditionsSeverityEnum(item.(interface{}))
	}

	return items
}

// flattenJobConditionsSeverityEnumSlice flattens the contents of JobConditionsSeverityEnum from a JSON
// response object.
func flattenJobConditionsSeverityEnumSlice(c *Client, i interface{}, res *Job) []JobConditionsSeverityEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobConditionsSeverityEnum{}
	}

	if len(a) == 0 {
		return []JobConditionsSeverityEnum{}
	}

	items := make([]JobConditionsSeverityEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobConditionsSeverityEnum(item.(interface{})))
	}

	return items
}

// flattenJobConditionsSeverityEnum asserts that an interface is a string, and returns a
// pointer to a *JobConditionsSeverityEnum with the same value as that string.
func flattenJobConditionsSeverityEnum(i interface{}) *JobConditionsSeverityEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobConditionsSeverityEnumRef(s)
}

// flattenJobConditionsReasonEnumMap flattens the contents of JobConditionsReasonEnum from a JSON
// response object.
func flattenJobConditionsReasonEnumMap(c *Client, i interface{}, res *Job) map[string]JobConditionsReasonEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobConditionsReasonEnum{}
	}

	if len(a) == 0 {
		return map[string]JobConditionsReasonEnum{}
	}

	items := make(map[string]JobConditionsReasonEnum)
	for k, item := range a {
		items[k] = *flattenJobConditionsReasonEnum(item.(interface{}))
	}

	return items
}

// flattenJobConditionsReasonEnumSlice flattens the contents of JobConditionsReasonEnum from a JSON
// response object.
func flattenJobConditionsReasonEnumSlice(c *Client, i interface{}, res *Job) []JobConditionsReasonEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobConditionsReasonEnum{}
	}

	if len(a) == 0 {
		return []JobConditionsReasonEnum{}
	}

	items := make([]JobConditionsReasonEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobConditionsReasonEnum(item.(interface{})))
	}

	return items
}

// flattenJobConditionsReasonEnum asserts that an interface is a string, and returns a
// pointer to a *JobConditionsReasonEnum with the same value as that string.
func flattenJobConditionsReasonEnum(i interface{}) *JobConditionsReasonEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobConditionsReasonEnumRef(s)
}

// flattenJobConditionsRevisionReasonEnumMap flattens the contents of JobConditionsRevisionReasonEnum from a JSON
// response object.
func flattenJobConditionsRevisionReasonEnumMap(c *Client, i interface{}, res *Job) map[string]JobConditionsRevisionReasonEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobConditionsRevisionReasonEnum{}
	}

	if len(a) == 0 {
		return map[string]JobConditionsRevisionReasonEnum{}
	}

	items := make(map[string]JobConditionsRevisionReasonEnum)
	for k, item := range a {
		items[k] = *flattenJobConditionsRevisionReasonEnum(item.(interface{}))
	}

	return items
}

// flattenJobConditionsRevisionReasonEnumSlice flattens the contents of JobConditionsRevisionReasonEnum from a JSON
// response object.
func flattenJobConditionsRevisionReasonEnumSlice(c *Client, i interface{}, res *Job) []JobConditionsRevisionReasonEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobConditionsRevisionReasonEnum{}
	}

	if len(a) == 0 {
		return []JobConditionsRevisionReasonEnum{}
	}

	items := make([]JobConditionsRevisionReasonEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobConditionsRevisionReasonEnum(item.(interface{})))
	}

	return items
}

// flattenJobConditionsRevisionReasonEnum asserts that an interface is a string, and returns a
// pointer to a *JobConditionsRevisionReasonEnum with the same value as that string.
func flattenJobConditionsRevisionReasonEnum(i interface{}) *JobConditionsRevisionReasonEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobConditionsRevisionReasonEnumRef(s)
}

// flattenJobConditionsExecutionReasonEnumMap flattens the contents of JobConditionsExecutionReasonEnum from a JSON
// response object.
func flattenJobConditionsExecutionReasonEnumMap(c *Client, i interface{}, res *Job) map[string]JobConditionsExecutionReasonEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobConditionsExecutionReasonEnum{}
	}

	if len(a) == 0 {
		return map[string]JobConditionsExecutionReasonEnum{}
	}

	items := make(map[string]JobConditionsExecutionReasonEnum)
	for k, item := range a {
		items[k] = *flattenJobConditionsExecutionReasonEnum(item.(interface{}))
	}

	return items
}

// flattenJobConditionsExecutionReasonEnumSlice flattens the contents of JobConditionsExecutionReasonEnum from a JSON
// response object.
func flattenJobConditionsExecutionReasonEnumSlice(c *Client, i interface{}, res *Job) []JobConditionsExecutionReasonEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobConditionsExecutionReasonEnum{}
	}

	if len(a) == 0 {
		return []JobConditionsExecutionReasonEnum{}
	}

	items := make([]JobConditionsExecutionReasonEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobConditionsExecutionReasonEnum(item.(interface{})))
	}

	return items
}

// flattenJobConditionsExecutionReasonEnum asserts that an interface is a string, and returns a
// pointer to a *JobConditionsExecutionReasonEnum with the same value as that string.
func flattenJobConditionsExecutionReasonEnum(i interface{}) *JobConditionsExecutionReasonEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobConditionsExecutionReasonEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Job) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalJob(b, c, r)
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

type jobDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         jobApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToJobDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]jobDiff, error) {
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
	var diffs []jobDiff
	// For each operation name, create a jobDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := jobDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToJobApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToJobApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (jobApiOperation, error) {
	switch opName {

	case "updateJobUpdateJobOperation":
		return &updateJobUpdateJobOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractJobFields(r *Job) error {
	vBinaryAuthorization := r.BinaryAuthorization
	if vBinaryAuthorization == nil {
		// note: explicitly not the empty object.
		vBinaryAuthorization = &JobBinaryAuthorization{}
	}
	if err := extractJobBinaryAuthorizationFields(r, vBinaryAuthorization); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vBinaryAuthorization) {
		r.BinaryAuthorization = vBinaryAuthorization
	}
	vTemplate := r.Template
	if vTemplate == nil {
		// note: explicitly not the empty object.
		vTemplate = &JobTemplate{}
	}
	if err := extractJobTemplateFields(r, vTemplate); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTemplate) {
		r.Template = vTemplate
	}
	vTerminalCondition := r.TerminalCondition
	if vTerminalCondition == nil {
		// note: explicitly not the empty object.
		vTerminalCondition = &JobTerminalCondition{}
	}
	if err := extractJobTerminalConditionFields(r, vTerminalCondition); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTerminalCondition) {
		r.TerminalCondition = vTerminalCondition
	}
	vLatestSucceededExecution := r.LatestSucceededExecution
	if vLatestSucceededExecution == nil {
		// note: explicitly not the empty object.
		vLatestSucceededExecution = &JobLatestSucceededExecution{}
	}
	if err := extractJobLatestSucceededExecutionFields(r, vLatestSucceededExecution); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLatestSucceededExecution) {
		r.LatestSucceededExecution = vLatestSucceededExecution
	}
	vLatestCreatedExecution := r.LatestCreatedExecution
	if vLatestCreatedExecution == nil {
		// note: explicitly not the empty object.
		vLatestCreatedExecution = &JobLatestCreatedExecution{}
	}
	if err := extractJobLatestCreatedExecutionFields(r, vLatestCreatedExecution); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLatestCreatedExecution) {
		r.LatestCreatedExecution = vLatestCreatedExecution
	}
	return nil
}
func extractJobBinaryAuthorizationFields(r *Job, o *JobBinaryAuthorization) error {
	return nil
}
func extractJobTemplateFields(r *Job, o *JobTemplate) error {
	vTemplate := o.Template
	if vTemplate == nil {
		// note: explicitly not the empty object.
		vTemplate = &JobTemplateTemplate{}
	}
	if err := extractJobTemplateTemplateFields(r, vTemplate); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTemplate) {
		o.Template = vTemplate
	}
	return nil
}
func extractJobTemplateTemplateFields(r *Job, o *JobTemplateTemplate) error {
	vVPCAccess := o.VPCAccess
	if vVPCAccess == nil {
		// note: explicitly not the empty object.
		vVPCAccess = &JobTemplateTemplateVPCAccess{}
	}
	if err := extractJobTemplateTemplateVPCAccessFields(r, vVPCAccess); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vVPCAccess) {
		o.VPCAccess = vVPCAccess
	}
	return nil
}
func extractJobTemplateTemplateContainersFields(r *Job, o *JobTemplateTemplateContainers) error {
	vResources := o.Resources
	if vResources == nil {
		// note: explicitly not the empty object.
		vResources = &JobTemplateTemplateContainersResources{}
	}
	if err := extractJobTemplateTemplateContainersResourcesFields(r, vResources); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vResources) {
		o.Resources = vResources
	}
	return nil
}
func extractJobTemplateTemplateContainersEnvFields(r *Job, o *JobTemplateTemplateContainersEnv) error {
	vValueSource := o.ValueSource
	if vValueSource == nil {
		// note: explicitly not the empty object.
		vValueSource = &JobTemplateTemplateContainersEnvValueSource{}
	}
	if err := extractJobTemplateTemplateContainersEnvValueSourceFields(r, vValueSource); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vValueSource) {
		o.ValueSource = vValueSource
	}
	return nil
}
func extractJobTemplateTemplateContainersEnvValueSourceFields(r *Job, o *JobTemplateTemplateContainersEnvValueSource) error {
	vSecretKeyRef := o.SecretKeyRef
	if vSecretKeyRef == nil {
		// note: explicitly not the empty object.
		vSecretKeyRef = &JobTemplateTemplateContainersEnvValueSourceSecretKeyRef{}
	}
	if err := extractJobTemplateTemplateContainersEnvValueSourceSecretKeyRefFields(r, vSecretKeyRef); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSecretKeyRef) {
		o.SecretKeyRef = vSecretKeyRef
	}
	return nil
}
func extractJobTemplateTemplateContainersEnvValueSourceSecretKeyRefFields(r *Job, o *JobTemplateTemplateContainersEnvValueSourceSecretKeyRef) error {
	return nil
}
func extractJobTemplateTemplateContainersResourcesFields(r *Job, o *JobTemplateTemplateContainersResources) error {
	return nil
}
func extractJobTemplateTemplateContainersPortsFields(r *Job, o *JobTemplateTemplateContainersPorts) error {
	return nil
}
func extractJobTemplateTemplateContainersVolumeMountsFields(r *Job, o *JobTemplateTemplateContainersVolumeMounts) error {
	return nil
}
func extractJobTemplateTemplateVolumesFields(r *Job, o *JobTemplateTemplateVolumes) error {
	vSecret := o.Secret
	if vSecret == nil {
		// note: explicitly not the empty object.
		vSecret = &JobTemplateTemplateVolumesSecret{}
	}
	if err := extractJobTemplateTemplateVolumesSecretFields(r, vSecret); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSecret) {
		o.Secret = vSecret
	}
	vCloudSqlInstance := o.CloudSqlInstance
	if vCloudSqlInstance == nil {
		// note: explicitly not the empty object.
		vCloudSqlInstance = &JobTemplateTemplateVolumesCloudSqlInstance{}
	}
	if err := extractJobTemplateTemplateVolumesCloudSqlInstanceFields(r, vCloudSqlInstance); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCloudSqlInstance) {
		o.CloudSqlInstance = vCloudSqlInstance
	}
	return nil
}
func extractJobTemplateTemplateVolumesSecretFields(r *Job, o *JobTemplateTemplateVolumesSecret) error {
	return nil
}
func extractJobTemplateTemplateVolumesSecretItemsFields(r *Job, o *JobTemplateTemplateVolumesSecretItems) error {
	return nil
}
func extractJobTemplateTemplateVolumesCloudSqlInstanceFields(r *Job, o *JobTemplateTemplateVolumesCloudSqlInstance) error {
	return nil
}
func extractJobTemplateTemplateVPCAccessFields(r *Job, o *JobTemplateTemplateVPCAccess) error {
	return nil
}
func extractJobTerminalConditionFields(r *Job, o *JobTerminalCondition) error {
	return nil
}
func extractJobConditionsFields(r *Job, o *JobConditions) error {
	return nil
}
func extractJobLatestSucceededExecutionFields(r *Job, o *JobLatestSucceededExecution) error {
	return nil
}
func extractJobLatestCreatedExecutionFields(r *Job, o *JobLatestCreatedExecution) error {
	return nil
}

func postReadExtractJobFields(r *Job) error {
	vBinaryAuthorization := r.BinaryAuthorization
	if vBinaryAuthorization == nil {
		// note: explicitly not the empty object.
		vBinaryAuthorization = &JobBinaryAuthorization{}
	}
	if err := postReadExtractJobBinaryAuthorizationFields(r, vBinaryAuthorization); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vBinaryAuthorization) {
		r.BinaryAuthorization = vBinaryAuthorization
	}
	vTemplate := r.Template
	if vTemplate == nil {
		// note: explicitly not the empty object.
		vTemplate = &JobTemplate{}
	}
	if err := postReadExtractJobTemplateFields(r, vTemplate); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTemplate) {
		r.Template = vTemplate
	}
	vTerminalCondition := r.TerminalCondition
	if vTerminalCondition == nil {
		// note: explicitly not the empty object.
		vTerminalCondition = &JobTerminalCondition{}
	}
	if err := postReadExtractJobTerminalConditionFields(r, vTerminalCondition); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTerminalCondition) {
		r.TerminalCondition = vTerminalCondition
	}
	vLatestSucceededExecution := r.LatestSucceededExecution
	if vLatestSucceededExecution == nil {
		// note: explicitly not the empty object.
		vLatestSucceededExecution = &JobLatestSucceededExecution{}
	}
	if err := postReadExtractJobLatestSucceededExecutionFields(r, vLatestSucceededExecution); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLatestSucceededExecution) {
		r.LatestSucceededExecution = vLatestSucceededExecution
	}
	vLatestCreatedExecution := r.LatestCreatedExecution
	if vLatestCreatedExecution == nil {
		// note: explicitly not the empty object.
		vLatestCreatedExecution = &JobLatestCreatedExecution{}
	}
	if err := postReadExtractJobLatestCreatedExecutionFields(r, vLatestCreatedExecution); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLatestCreatedExecution) {
		r.LatestCreatedExecution = vLatestCreatedExecution
	}
	return nil
}
func postReadExtractJobBinaryAuthorizationFields(r *Job, o *JobBinaryAuthorization) error {
	return nil
}
func postReadExtractJobTemplateFields(r *Job, o *JobTemplate) error {
	vTemplate := o.Template
	if vTemplate == nil {
		// note: explicitly not the empty object.
		vTemplate = &JobTemplateTemplate{}
	}
	if err := extractJobTemplateTemplateFields(r, vTemplate); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTemplate) {
		o.Template = vTemplate
	}
	return nil
}
func postReadExtractJobTemplateTemplateFields(r *Job, o *JobTemplateTemplate) error {
	vVPCAccess := o.VPCAccess
	if vVPCAccess == nil {
		// note: explicitly not the empty object.
		vVPCAccess = &JobTemplateTemplateVPCAccess{}
	}
	if err := extractJobTemplateTemplateVPCAccessFields(r, vVPCAccess); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vVPCAccess) {
		o.VPCAccess = vVPCAccess
	}
	return nil
}
func postReadExtractJobTemplateTemplateContainersFields(r *Job, o *JobTemplateTemplateContainers) error {
	vResources := o.Resources
	if vResources == nil {
		// note: explicitly not the empty object.
		vResources = &JobTemplateTemplateContainersResources{}
	}
	if err := extractJobTemplateTemplateContainersResourcesFields(r, vResources); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vResources) {
		o.Resources = vResources
	}
	return nil
}
func postReadExtractJobTemplateTemplateContainersEnvFields(r *Job, o *JobTemplateTemplateContainersEnv) error {
	vValueSource := o.ValueSource
	if vValueSource == nil {
		// note: explicitly not the empty object.
		vValueSource = &JobTemplateTemplateContainersEnvValueSource{}
	}
	if err := extractJobTemplateTemplateContainersEnvValueSourceFields(r, vValueSource); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vValueSource) {
		o.ValueSource = vValueSource
	}
	return nil
}
func postReadExtractJobTemplateTemplateContainersEnvValueSourceFields(r *Job, o *JobTemplateTemplateContainersEnvValueSource) error {
	vSecretKeyRef := o.SecretKeyRef
	if vSecretKeyRef == nil {
		// note: explicitly not the empty object.
		vSecretKeyRef = &JobTemplateTemplateContainersEnvValueSourceSecretKeyRef{}
	}
	if err := extractJobTemplateTemplateContainersEnvValueSourceSecretKeyRefFields(r, vSecretKeyRef); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSecretKeyRef) {
		o.SecretKeyRef = vSecretKeyRef
	}
	return nil
}
func postReadExtractJobTemplateTemplateContainersEnvValueSourceSecretKeyRefFields(r *Job, o *JobTemplateTemplateContainersEnvValueSourceSecretKeyRef) error {
	return nil
}
func postReadExtractJobTemplateTemplateContainersResourcesFields(r *Job, o *JobTemplateTemplateContainersResources) error {
	return nil
}
func postReadExtractJobTemplateTemplateContainersPortsFields(r *Job, o *JobTemplateTemplateContainersPorts) error {
	return nil
}
func postReadExtractJobTemplateTemplateContainersVolumeMountsFields(r *Job, o *JobTemplateTemplateContainersVolumeMounts) error {
	return nil
}
func postReadExtractJobTemplateTemplateVolumesFields(r *Job, o *JobTemplateTemplateVolumes) error {
	vSecret := o.Secret
	if vSecret == nil {
		// note: explicitly not the empty object.
		vSecret = &JobTemplateTemplateVolumesSecret{}
	}
	if err := extractJobTemplateTemplateVolumesSecretFields(r, vSecret); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSecret) {
		o.Secret = vSecret
	}
	vCloudSqlInstance := o.CloudSqlInstance
	if vCloudSqlInstance == nil {
		// note: explicitly not the empty object.
		vCloudSqlInstance = &JobTemplateTemplateVolumesCloudSqlInstance{}
	}
	if err := extractJobTemplateTemplateVolumesCloudSqlInstanceFields(r, vCloudSqlInstance); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCloudSqlInstance) {
		o.CloudSqlInstance = vCloudSqlInstance
	}
	return nil
}
func postReadExtractJobTemplateTemplateVolumesSecretFields(r *Job, o *JobTemplateTemplateVolumesSecret) error {
	return nil
}
func postReadExtractJobTemplateTemplateVolumesSecretItemsFields(r *Job, o *JobTemplateTemplateVolumesSecretItems) error {
	return nil
}
func postReadExtractJobTemplateTemplateVolumesCloudSqlInstanceFields(r *Job, o *JobTemplateTemplateVolumesCloudSqlInstance) error {
	return nil
}
func postReadExtractJobTemplateTemplateVPCAccessFields(r *Job, o *JobTemplateTemplateVPCAccess) error {
	return nil
}
func postReadExtractJobTerminalConditionFields(r *Job, o *JobTerminalCondition) error {
	return nil
}
func postReadExtractJobConditionsFields(r *Job, o *JobConditions) error {
	return nil
}
func postReadExtractJobLatestSucceededExecutionFields(r *Job, o *JobLatestSucceededExecution) error {
	return nil
}
func postReadExtractJobLatestCreatedExecutionFields(r *Job, o *JobLatestCreatedExecution) error {
	return nil
}
