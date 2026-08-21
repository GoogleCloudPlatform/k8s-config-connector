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
)

func (r *Job) validate() error {

	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"PubsubTarget", "AppEngineHttpTarget", "HttpTarget"}, r.PubsubTarget, r.AppEngineHttpTarget, r.HttpTarget); err != nil {
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
	if !dcl.IsEmptyValueIndirect(r.PubsubTarget) {
		if err := r.PubsubTarget.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.AppEngineHttpTarget) {
		if err := r.AppEngineHttpTarget.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.HttpTarget) {
		if err := r.HttpTarget.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Status) {
		if err := r.Status.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.RetryConfig) {
		if err := r.RetryConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobPubsubTarget) validate() error {
	if err := dcl.Required(r, "topicName"); err != nil {
		return err
	}
	return nil
}
func (r *JobAppEngineHttpTarget) validate() error {
	if !dcl.IsEmptyValueIndirect(r.AppEngineRouting) {
		if err := r.AppEngineRouting.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobAppEngineHttpTargetAppEngineRouting) validate() error {
	return nil
}
func (r *JobHttpTarget) validate() error {
	if err := dcl.Required(r, "uri"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.OAuthToken) {
		if err := r.OAuthToken.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.OidcToken) {
		if err := r.OidcToken.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobHttpTargetOAuthToken) validate() error {
	return nil
}
func (r *JobHttpTargetOidcToken) validate() error {
	return nil
}
func (r *JobStatus) validate() error {
	return nil
}
func (r *JobStatusDetails) validate() error {
	return nil
}
func (r *JobRetryConfig) validate() error {
	return nil
}
func (r *Job) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://cloudscheduler.googleapis.com/v1/", params)
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
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/jobs", nr.basePath(), userBasePath, params), nil

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
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v, err := expandJobPubsubTarget(c, f.PubsubTarget, res); err != nil {
		return nil, fmt.Errorf("error expanding PubsubTarget into pubsubTarget: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["pubsubTarget"] = v
	}
	if v, err := expandJobAppEngineHttpTarget(c, f.AppEngineHttpTarget, res); err != nil {
		return nil, fmt.Errorf("error expanding AppEngineHttpTarget into appEngineHttpTarget: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["appEngineHttpTarget"] = v
	}
	if v, err := expandJobHttpTarget(c, f.HttpTarget, res); err != nil {
		return nil, fmt.Errorf("error expanding HttpTarget into httpTarget: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["httpTarget"] = v
	}
	if v := f.Schedule; !dcl.IsEmptyValueIndirect(v) {
		req["schedule"] = v
	}
	if v := f.TimeZone; !dcl.IsEmptyValueIndirect(v) {
		req["timeZone"] = v
	}
	if v, err := expandJobRetryConfig(c, f.RetryConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding RetryConfig into retryConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["retryConfig"] = v
	}
	if v := f.AttemptDeadline; !dcl.IsEmptyValueIndirect(v) {
		req["attemptDeadline"] = v
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
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
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

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Job: %w", err)
	}

	// We saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// This is the reason we are adding retry to handle that case.
	retriesRemaining := 10
	dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		_, err := c.GetJob(ctx, r)
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

	o, err := dcl.ResponseBodyAsJSON(resp)
	if err != nil {
		return fmt.Errorf("error decoding response body into JSON: %w", err)
	}
	op.response = o

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

	if !dcl.IsZeroValue(rawInitial.PubsubTarget) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.AppEngineHttpTarget, rawInitial.HttpTarget) {
			rawInitial.PubsubTarget = EmptyJobPubsubTarget
		}
	}

	if !dcl.IsZeroValue(rawInitial.AppEngineHttpTarget) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.PubsubTarget, rawInitial.HttpTarget) {
			rawInitial.AppEngineHttpTarget = EmptyJobAppEngineHttpTarget
		}
	}

	if !dcl.IsZeroValue(rawInitial.HttpTarget) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.PubsubTarget, rawInitial.AppEngineHttpTarget) {
			rawInitial.HttpTarget = EmptyJobHttpTarget
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

func canonicalizeJobDesiredState(rawDesired, rawInitial *Job, opts ...dcl.ApplyOption) (*Job, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.PubsubTarget = canonicalizeJobPubsubTarget(rawDesired.PubsubTarget, nil, opts...)
		rawDesired.AppEngineHttpTarget = canonicalizeJobAppEngineHttpTarget(rawDesired.AppEngineHttpTarget, nil, opts...)
		rawDesired.HttpTarget = canonicalizeJobHttpTarget(rawDesired.HttpTarget, nil, opts...)
		rawDesired.Status = canonicalizeJobStatus(rawDesired.Status, nil, opts...)
		rawDesired.RetryConfig = canonicalizeJobRetryConfig(rawDesired.RetryConfig, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Job{}
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
	canonicalDesired.PubsubTarget = canonicalizeJobPubsubTarget(rawDesired.PubsubTarget, rawInitial.PubsubTarget, opts...)
	canonicalDesired.AppEngineHttpTarget = canonicalizeJobAppEngineHttpTarget(rawDesired.AppEngineHttpTarget, rawInitial.AppEngineHttpTarget, opts...)
	canonicalDesired.HttpTarget = canonicalizeJobHttpTarget(rawDesired.HttpTarget, rawInitial.HttpTarget, opts...)
	if dcl.StringCanonicalize(rawDesired.Schedule, rawInitial.Schedule) {
		canonicalDesired.Schedule = rawInitial.Schedule
	} else {
		canonicalDesired.Schedule = rawDesired.Schedule
	}
	if dcl.StringCanonicalize(rawDesired.TimeZone, rawInitial.TimeZone) {
		canonicalDesired.TimeZone = rawInitial.TimeZone
	} else {
		canonicalDesired.TimeZone = rawDesired.TimeZone
	}
	canonicalDesired.RetryConfig = canonicalizeJobRetryConfig(rawDesired.RetryConfig, rawInitial.RetryConfig, opts...)
	if dcl.StringCanonicalize(rawDesired.AttemptDeadline, rawInitial.AttemptDeadline) {
		canonicalDesired.AttemptDeadline = rawInitial.AttemptDeadline
	} else {
		canonicalDesired.AttemptDeadline = rawDesired.AttemptDeadline
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

	if canonicalDesired.PubsubTarget != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.AppEngineHttpTarget, rawDesired.HttpTarget) {
			canonicalDesired.PubsubTarget = EmptyJobPubsubTarget
		}
	}

	if canonicalDesired.AppEngineHttpTarget != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.PubsubTarget, rawDesired.HttpTarget) {
			canonicalDesired.AppEngineHttpTarget = EmptyJobAppEngineHttpTarget
		}
	}

	if canonicalDesired.HttpTarget != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.PubsubTarget, rawDesired.AppEngineHttpTarget) {
			canonicalDesired.HttpTarget = EmptyJobHttpTarget
		}
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

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.PubsubTarget) && dcl.IsEmptyValueIndirect(rawDesired.PubsubTarget) {
		rawNew.PubsubTarget = rawDesired.PubsubTarget
	} else {
		rawNew.PubsubTarget = canonicalizeNewJobPubsubTarget(c, rawDesired.PubsubTarget, rawNew.PubsubTarget)
	}

	if dcl.IsEmptyValueIndirect(rawNew.AppEngineHttpTarget) && dcl.IsEmptyValueIndirect(rawDesired.AppEngineHttpTarget) {
		rawNew.AppEngineHttpTarget = rawDesired.AppEngineHttpTarget
	} else {
		rawNew.AppEngineHttpTarget = canonicalizeNewJobAppEngineHttpTarget(c, rawDesired.AppEngineHttpTarget, rawNew.AppEngineHttpTarget)
	}

	if dcl.IsEmptyValueIndirect(rawNew.HttpTarget) && dcl.IsEmptyValueIndirect(rawDesired.HttpTarget) {
		rawNew.HttpTarget = rawDesired.HttpTarget
	} else {
		rawNew.HttpTarget = canonicalizeNewJobHttpTarget(c, rawDesired.HttpTarget, rawNew.HttpTarget)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Schedule) && dcl.IsEmptyValueIndirect(rawDesired.Schedule) {
		rawNew.Schedule = rawDesired.Schedule
	} else {
		if dcl.StringCanonicalize(rawDesired.Schedule, rawNew.Schedule) {
			rawNew.Schedule = rawDesired.Schedule
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.TimeZone) && dcl.IsEmptyValueIndirect(rawDesired.TimeZone) {
		rawNew.TimeZone = rawDesired.TimeZone
	} else {
		if dcl.StringCanonicalize(rawDesired.TimeZone, rawNew.TimeZone) {
			rawNew.TimeZone = rawDesired.TimeZone
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.UserUpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UserUpdateTime) {
		rawNew.UserUpdateTime = rawDesired.UserUpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Status) && dcl.IsEmptyValueIndirect(rawDesired.Status) {
		rawNew.Status = rawDesired.Status
	} else {
		rawNew.Status = canonicalizeNewJobStatus(c, rawDesired.Status, rawNew.Status)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ScheduleTime) && dcl.IsEmptyValueIndirect(rawDesired.ScheduleTime) {
		rawNew.ScheduleTime = rawDesired.ScheduleTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.LastAttemptTime) && dcl.IsEmptyValueIndirect(rawDesired.LastAttemptTime) {
		rawNew.LastAttemptTime = rawDesired.LastAttemptTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.RetryConfig) && dcl.IsEmptyValueIndirect(rawDesired.RetryConfig) {
		rawNew.RetryConfig = rawDesired.RetryConfig
	} else {
		rawNew.RetryConfig = canonicalizeNewJobRetryConfig(c, rawDesired.RetryConfig, rawNew.RetryConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.AttemptDeadline) && dcl.IsEmptyValueIndirect(rawDesired.AttemptDeadline) {
		rawNew.AttemptDeadline = rawDesired.AttemptDeadline
	} else {
		if dcl.StringCanonicalize(rawDesired.AttemptDeadline, rawNew.AttemptDeadline) {
			rawNew.AttemptDeadline = rawDesired.AttemptDeadline
		}
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeJobPubsubTarget(des, initial *JobPubsubTarget, opts ...dcl.ApplyOption) *JobPubsubTarget {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobPubsubTarget{}

	if dcl.IsZeroValue(des.TopicName) || (dcl.IsEmptyValueIndirect(des.TopicName) && dcl.IsEmptyValueIndirect(initial.TopicName)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.TopicName = initial.TopicName
	} else {
		cDes.TopicName = des.TopicName
	}
	if dcl.StringCanonicalize(des.Data, initial.Data) || dcl.IsZeroValue(des.Data) {
		cDes.Data = initial.Data
	} else {
		cDes.Data = des.Data
	}
	if dcl.IsZeroValue(des.Attributes) || (dcl.IsEmptyValueIndirect(des.Attributes) && dcl.IsEmptyValueIndirect(initial.Attributes)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Attributes = initial.Attributes
	} else {
		cDes.Attributes = des.Attributes
	}

	return cDes
}

func canonicalizeJobPubsubTargetSlice(des, initial []JobPubsubTarget, opts ...dcl.ApplyOption) []JobPubsubTarget {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobPubsubTarget, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobPubsubTarget(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobPubsubTarget, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobPubsubTarget(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobPubsubTarget(c *Client, des, nw *JobPubsubTarget) *JobPubsubTarget {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobPubsubTarget while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Data, nw.Data) {
		nw.Data = des.Data
	}

	return nw
}

func canonicalizeNewJobPubsubTargetSet(c *Client, des, nw []JobPubsubTarget) []JobPubsubTarget {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobPubsubTarget
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobPubsubTargetNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobPubsubTarget(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobPubsubTargetSlice(c *Client, des, nw []JobPubsubTarget) []JobPubsubTarget {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobPubsubTarget
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobPubsubTarget(c, &d, &n))
	}

	return items
}

func canonicalizeJobAppEngineHttpTarget(des, initial *JobAppEngineHttpTarget, opts ...dcl.ApplyOption) *JobAppEngineHttpTarget {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobAppEngineHttpTarget{}

	if dcl.IsZeroValue(des.HttpMethod) || (dcl.IsEmptyValueIndirect(des.HttpMethod) && dcl.IsEmptyValueIndirect(initial.HttpMethod)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.HttpMethod = initial.HttpMethod
	} else {
		cDes.HttpMethod = des.HttpMethod
	}
	cDes.AppEngineRouting = canonicalizeJobAppEngineHttpTargetAppEngineRouting(des.AppEngineRouting, initial.AppEngineRouting, opts...)
	if dcl.StringCanonicalize(des.RelativeUri, initial.RelativeUri) || dcl.IsZeroValue(des.RelativeUri) {
		cDes.RelativeUri = initial.RelativeUri
	} else {
		cDes.RelativeUri = des.RelativeUri
	}
	if dcl.IsZeroValue(des.Headers) || (dcl.IsEmptyValueIndirect(des.Headers) && dcl.IsEmptyValueIndirect(initial.Headers)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Headers = initial.Headers
	} else {
		cDes.Headers = des.Headers
	}
	if dcl.StringCanonicalize(des.Body, initial.Body) || dcl.IsZeroValue(des.Body) {
		cDes.Body = initial.Body
	} else {
		cDes.Body = des.Body
	}

	return cDes
}

func canonicalizeJobAppEngineHttpTargetSlice(des, initial []JobAppEngineHttpTarget, opts ...dcl.ApplyOption) []JobAppEngineHttpTarget {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobAppEngineHttpTarget, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobAppEngineHttpTarget(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobAppEngineHttpTarget, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobAppEngineHttpTarget(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobAppEngineHttpTarget(c *Client, des, nw *JobAppEngineHttpTarget) *JobAppEngineHttpTarget {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobAppEngineHttpTarget while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.AppEngineRouting = canonicalizeNewJobAppEngineHttpTargetAppEngineRouting(c, des.AppEngineRouting, nw.AppEngineRouting)
	if dcl.StringCanonicalize(des.RelativeUri, nw.RelativeUri) {
		nw.RelativeUri = des.RelativeUri
	}
	if dcl.StringCanonicalize(des.Body, nw.Body) {
		nw.Body = des.Body
	}

	return nw
}

func canonicalizeNewJobAppEngineHttpTargetSet(c *Client, des, nw []JobAppEngineHttpTarget) []JobAppEngineHttpTarget {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobAppEngineHttpTarget
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobAppEngineHttpTargetNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobAppEngineHttpTarget(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobAppEngineHttpTargetSlice(c *Client, des, nw []JobAppEngineHttpTarget) []JobAppEngineHttpTarget {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobAppEngineHttpTarget
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobAppEngineHttpTarget(c, &d, &n))
	}

	return items
}

func canonicalizeJobAppEngineHttpTargetAppEngineRouting(des, initial *JobAppEngineHttpTargetAppEngineRouting, opts ...dcl.ApplyOption) *JobAppEngineHttpTargetAppEngineRouting {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobAppEngineHttpTargetAppEngineRouting{}

	if dcl.StringCanonicalize(des.Service, initial.Service) || dcl.IsZeroValue(des.Service) {
		cDes.Service = initial.Service
	} else {
		cDes.Service = des.Service
	}
	if dcl.StringCanonicalize(des.Version, initial.Version) || dcl.IsZeroValue(des.Version) {
		cDes.Version = initial.Version
	} else {
		cDes.Version = des.Version
	}
	if dcl.StringCanonicalize(des.Instance, initial.Instance) || dcl.IsZeroValue(des.Instance) {
		cDes.Instance = initial.Instance
	} else {
		cDes.Instance = des.Instance
	}

	return cDes
}

func canonicalizeJobAppEngineHttpTargetAppEngineRoutingSlice(des, initial []JobAppEngineHttpTargetAppEngineRouting, opts ...dcl.ApplyOption) []JobAppEngineHttpTargetAppEngineRouting {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobAppEngineHttpTargetAppEngineRouting, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobAppEngineHttpTargetAppEngineRouting(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobAppEngineHttpTargetAppEngineRouting, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobAppEngineHttpTargetAppEngineRouting(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobAppEngineHttpTargetAppEngineRouting(c *Client, des, nw *JobAppEngineHttpTargetAppEngineRouting) *JobAppEngineHttpTargetAppEngineRouting {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobAppEngineHttpTargetAppEngineRouting while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Service, nw.Service) {
		nw.Service = des.Service
	}
	if dcl.StringCanonicalize(des.Version, nw.Version) {
		nw.Version = des.Version
	}
	if dcl.StringCanonicalize(des.Instance, nw.Instance) {
		nw.Instance = des.Instance
	}
	if dcl.StringCanonicalize(des.Host, nw.Host) {
		nw.Host = des.Host
	}

	return nw
}

func canonicalizeNewJobAppEngineHttpTargetAppEngineRoutingSet(c *Client, des, nw []JobAppEngineHttpTargetAppEngineRouting) []JobAppEngineHttpTargetAppEngineRouting {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobAppEngineHttpTargetAppEngineRouting
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobAppEngineHttpTargetAppEngineRoutingNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobAppEngineHttpTargetAppEngineRouting(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobAppEngineHttpTargetAppEngineRoutingSlice(c *Client, des, nw []JobAppEngineHttpTargetAppEngineRouting) []JobAppEngineHttpTargetAppEngineRouting {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobAppEngineHttpTargetAppEngineRouting
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobAppEngineHttpTargetAppEngineRouting(c, &d, &n))
	}

	return items
}

func canonicalizeJobHttpTarget(des, initial *JobHttpTarget, opts ...dcl.ApplyOption) *JobHttpTarget {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobHttpTarget{}

	if dcl.StringCanonicalize(des.Uri, initial.Uri) || dcl.IsZeroValue(des.Uri) {
		cDes.Uri = initial.Uri
	} else {
		cDes.Uri = des.Uri
	}
	if dcl.IsZeroValue(des.HttpMethod) || (dcl.IsEmptyValueIndirect(des.HttpMethod) && dcl.IsEmptyValueIndirect(initial.HttpMethod)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.HttpMethod = initial.HttpMethod
	} else {
		cDes.HttpMethod = des.HttpMethod
	}
	if dcl.IsZeroValue(des.Headers) || (dcl.IsEmptyValueIndirect(des.Headers) && dcl.IsEmptyValueIndirect(initial.Headers)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Headers = initial.Headers
	} else {
		cDes.Headers = des.Headers
	}
	if dcl.StringCanonicalize(des.Body, initial.Body) || dcl.IsZeroValue(des.Body) {
		cDes.Body = initial.Body
	} else {
		cDes.Body = des.Body
	}
	cDes.OAuthToken = canonicalizeJobHttpTargetOAuthToken(des.OAuthToken, initial.OAuthToken, opts...)
	cDes.OidcToken = canonicalizeJobHttpTargetOidcToken(des.OidcToken, initial.OidcToken, opts...)

	return cDes
}

func canonicalizeJobHttpTargetSlice(des, initial []JobHttpTarget, opts ...dcl.ApplyOption) []JobHttpTarget {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobHttpTarget, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobHttpTarget(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobHttpTarget, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobHttpTarget(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobHttpTarget(c *Client, des, nw *JobHttpTarget) *JobHttpTarget {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobHttpTarget while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Uri, nw.Uri) {
		nw.Uri = des.Uri
	}
	if dcl.StringCanonicalize(des.Body, nw.Body) {
		nw.Body = des.Body
	}
	nw.OAuthToken = canonicalizeNewJobHttpTargetOAuthToken(c, des.OAuthToken, nw.OAuthToken)
	nw.OidcToken = canonicalizeNewJobHttpTargetOidcToken(c, des.OidcToken, nw.OidcToken)

	return nw
}

func canonicalizeNewJobHttpTargetSet(c *Client, des, nw []JobHttpTarget) []JobHttpTarget {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobHttpTarget
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobHttpTargetNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobHttpTarget(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobHttpTargetSlice(c *Client, des, nw []JobHttpTarget) []JobHttpTarget {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobHttpTarget
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobHttpTarget(c, &d, &n))
	}

	return items
}

func canonicalizeJobHttpTargetOAuthToken(des, initial *JobHttpTargetOAuthToken, opts ...dcl.ApplyOption) *JobHttpTargetOAuthToken {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobHttpTargetOAuthToken{}

	if dcl.IsZeroValue(des.ServiceAccountEmail) || (dcl.IsEmptyValueIndirect(des.ServiceAccountEmail) && dcl.IsEmptyValueIndirect(initial.ServiceAccountEmail)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ServiceAccountEmail = initial.ServiceAccountEmail
	} else {
		cDes.ServiceAccountEmail = des.ServiceAccountEmail
	}
	if dcl.StringCanonicalize(des.Scope, initial.Scope) || dcl.IsZeroValue(des.Scope) {
		cDes.Scope = initial.Scope
	} else {
		cDes.Scope = des.Scope
	}

	return cDes
}

func canonicalizeJobHttpTargetOAuthTokenSlice(des, initial []JobHttpTargetOAuthToken, opts ...dcl.ApplyOption) []JobHttpTargetOAuthToken {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobHttpTargetOAuthToken, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobHttpTargetOAuthToken(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobHttpTargetOAuthToken, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobHttpTargetOAuthToken(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobHttpTargetOAuthToken(c *Client, des, nw *JobHttpTargetOAuthToken) *JobHttpTargetOAuthToken {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobHttpTargetOAuthToken while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Scope, nw.Scope) {
		nw.Scope = des.Scope
	}

	return nw
}

func canonicalizeNewJobHttpTargetOAuthTokenSet(c *Client, des, nw []JobHttpTargetOAuthToken) []JobHttpTargetOAuthToken {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobHttpTargetOAuthToken
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobHttpTargetOAuthTokenNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobHttpTargetOAuthToken(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobHttpTargetOAuthTokenSlice(c *Client, des, nw []JobHttpTargetOAuthToken) []JobHttpTargetOAuthToken {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobHttpTargetOAuthToken
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobHttpTargetOAuthToken(c, &d, &n))
	}

	return items
}

func canonicalizeJobHttpTargetOidcToken(des, initial *JobHttpTargetOidcToken, opts ...dcl.ApplyOption) *JobHttpTargetOidcToken {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobHttpTargetOidcToken{}

	if dcl.IsZeroValue(des.ServiceAccountEmail) || (dcl.IsEmptyValueIndirect(des.ServiceAccountEmail) && dcl.IsEmptyValueIndirect(initial.ServiceAccountEmail)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ServiceAccountEmail = initial.ServiceAccountEmail
	} else {
		cDes.ServiceAccountEmail = des.ServiceAccountEmail
	}
	if dcl.StringCanonicalize(des.Audience, initial.Audience) || dcl.IsZeroValue(des.Audience) {
		cDes.Audience = initial.Audience
	} else {
		cDes.Audience = des.Audience
	}

	return cDes
}

func canonicalizeJobHttpTargetOidcTokenSlice(des, initial []JobHttpTargetOidcToken, opts ...dcl.ApplyOption) []JobHttpTargetOidcToken {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobHttpTargetOidcToken, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobHttpTargetOidcToken(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobHttpTargetOidcToken, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobHttpTargetOidcToken(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobHttpTargetOidcToken(c *Client, des, nw *JobHttpTargetOidcToken) *JobHttpTargetOidcToken {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobHttpTargetOidcToken while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Audience, nw.Audience) {
		nw.Audience = des.Audience
	}

	return nw
}

func canonicalizeNewJobHttpTargetOidcTokenSet(c *Client, des, nw []JobHttpTargetOidcToken) []JobHttpTargetOidcToken {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobHttpTargetOidcToken
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobHttpTargetOidcTokenNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobHttpTargetOidcToken(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobHttpTargetOidcTokenSlice(c *Client, des, nw []JobHttpTargetOidcToken) []JobHttpTargetOidcToken {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobHttpTargetOidcToken
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobHttpTargetOidcToken(c, &d, &n))
	}

	return items
}

func canonicalizeJobStatus(des, initial *JobStatus, opts ...dcl.ApplyOption) *JobStatus {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobStatus{}

	if dcl.IsZeroValue(des.Code) || (dcl.IsEmptyValueIndirect(des.Code) && dcl.IsEmptyValueIndirect(initial.Code)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Code = initial.Code
	} else {
		cDes.Code = des.Code
	}
	if dcl.StringCanonicalize(des.Message, initial.Message) || dcl.IsZeroValue(des.Message) {
		cDes.Message = initial.Message
	} else {
		cDes.Message = des.Message
	}
	cDes.Details = canonicalizeJobStatusDetailsSlice(des.Details, initial.Details, opts...)

	return cDes
}

func canonicalizeJobStatusSlice(des, initial []JobStatus, opts ...dcl.ApplyOption) []JobStatus {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobStatus, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobStatus(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobStatus, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobStatus(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobStatus(c *Client, des, nw *JobStatus) *JobStatus {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobStatus while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Message, nw.Message) {
		nw.Message = des.Message
	}
	nw.Details = canonicalizeNewJobStatusDetailsSlice(c, des.Details, nw.Details)

	return nw
}

func canonicalizeNewJobStatusSet(c *Client, des, nw []JobStatus) []JobStatus {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobStatus
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobStatusNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobStatus(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobStatusSlice(c *Client, des, nw []JobStatus) []JobStatus {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobStatus
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobStatus(c, &d, &n))
	}

	return items
}

func canonicalizeJobStatusDetails(des, initial *JobStatusDetails, opts ...dcl.ApplyOption) *JobStatusDetails {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobStatusDetails{}

	if dcl.StringCanonicalize(des.TypeUrl, initial.TypeUrl) || dcl.IsZeroValue(des.TypeUrl) {
		cDes.TypeUrl = initial.TypeUrl
	} else {
		cDes.TypeUrl = des.TypeUrl
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		cDes.Value = initial.Value
	} else {
		cDes.Value = des.Value
	}

	return cDes
}

func canonicalizeJobStatusDetailsSlice(des, initial []JobStatusDetails, opts ...dcl.ApplyOption) []JobStatusDetails {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobStatusDetails, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobStatusDetails(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobStatusDetails, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobStatusDetails(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobStatusDetails(c *Client, des, nw *JobStatusDetails) *JobStatusDetails {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobStatusDetails while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.TypeUrl, nw.TypeUrl) {
		nw.TypeUrl = des.TypeUrl
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}

	return nw
}

func canonicalizeNewJobStatusDetailsSet(c *Client, des, nw []JobStatusDetails) []JobStatusDetails {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobStatusDetails
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobStatusDetailsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobStatusDetails(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobStatusDetailsSlice(c *Client, des, nw []JobStatusDetails) []JobStatusDetails {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobStatusDetails
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobStatusDetails(c, &d, &n))
	}

	return items
}

func canonicalizeJobRetryConfig(des, initial *JobRetryConfig, opts ...dcl.ApplyOption) *JobRetryConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobRetryConfig{}

	if dcl.IsZeroValue(des.RetryCount) || (dcl.IsEmptyValueIndirect(des.RetryCount) && dcl.IsEmptyValueIndirect(initial.RetryCount)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.RetryCount = initial.RetryCount
	} else {
		cDes.RetryCount = des.RetryCount
	}
	if dcl.StringCanonicalize(des.MaxRetryDuration, initial.MaxRetryDuration) || dcl.IsZeroValue(des.MaxRetryDuration) {
		cDes.MaxRetryDuration = initial.MaxRetryDuration
	} else {
		cDes.MaxRetryDuration = des.MaxRetryDuration
	}
	if dcl.StringCanonicalize(des.MinBackoffDuration, initial.MinBackoffDuration) || dcl.IsZeroValue(des.MinBackoffDuration) {
		cDes.MinBackoffDuration = initial.MinBackoffDuration
	} else {
		cDes.MinBackoffDuration = des.MinBackoffDuration
	}
	if dcl.StringCanonicalize(des.MaxBackoffDuration, initial.MaxBackoffDuration) || dcl.IsZeroValue(des.MaxBackoffDuration) {
		cDes.MaxBackoffDuration = initial.MaxBackoffDuration
	} else {
		cDes.MaxBackoffDuration = des.MaxBackoffDuration
	}
	if dcl.IsZeroValue(des.MaxDoublings) || (dcl.IsEmptyValueIndirect(des.MaxDoublings) && dcl.IsEmptyValueIndirect(initial.MaxDoublings)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MaxDoublings = initial.MaxDoublings
	} else {
		cDes.MaxDoublings = des.MaxDoublings
	}

	return cDes
}

func canonicalizeJobRetryConfigSlice(des, initial []JobRetryConfig, opts ...dcl.ApplyOption) []JobRetryConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobRetryConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobRetryConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobRetryConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobRetryConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobRetryConfig(c *Client, des, nw *JobRetryConfig) *JobRetryConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobRetryConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.MaxRetryDuration, nw.MaxRetryDuration) {
		nw.MaxRetryDuration = des.MaxRetryDuration
	}
	if dcl.StringCanonicalize(des.MinBackoffDuration, nw.MinBackoffDuration) {
		nw.MinBackoffDuration = des.MinBackoffDuration
	}
	if dcl.StringCanonicalize(des.MaxBackoffDuration, nw.MaxBackoffDuration) {
		nw.MaxBackoffDuration = des.MaxBackoffDuration
	}

	return nw
}

func canonicalizeNewJobRetryConfigSet(c *Client, des, nw []JobRetryConfig) []JobRetryConfig {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobRetryConfig
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobRetryConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobRetryConfig(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobRetryConfigSlice(c *Client, des, nw []JobRetryConfig) []JobRetryConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobRetryConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobRetryConfig(c, &d, &n))
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

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PubsubTarget, actual.PubsubTarget, dcl.DiffInfo{ObjectFunction: compareJobPubsubTargetNewStyle, EmptyObject: EmptyJobPubsubTarget, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("PubsubTarget")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AppEngineHttpTarget, actual.AppEngineHttpTarget, dcl.DiffInfo{ObjectFunction: compareJobAppEngineHttpTargetNewStyle, EmptyObject: EmptyJobAppEngineHttpTarget, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("AppEngineHttpTarget")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HttpTarget, actual.HttpTarget, dcl.DiffInfo{ObjectFunction: compareJobHttpTargetNewStyle, EmptyObject: EmptyJobHttpTarget, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("HttpTarget")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Schedule, actual.Schedule, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Schedule")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimeZone, actual.TimeZone, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("TimeZone")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UserUpdateTime, actual.UserUpdateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UserUpdateTime")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Status, actual.Status, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareJobStatusNewStyle, EmptyObject: EmptyJobStatus, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Status")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ScheduleTime, actual.ScheduleTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ScheduleTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LastAttemptTime, actual.LastAttemptTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LastAttemptTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RetryConfig, actual.RetryConfig, dcl.DiffInfo{ObjectFunction: compareJobRetryConfigNewStyle, EmptyObject: EmptyJobRetryConfig, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("RetryConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AttemptDeadline, actual.AttemptDeadline, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("AttemptDeadline")); len(ds) != 0 || err != nil {
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
func compareJobPubsubTargetNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobPubsubTarget)
	if !ok {
		desiredNotPointer, ok := d.(JobPubsubTarget)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobPubsubTarget or *JobPubsubTarget", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobPubsubTarget)
	if !ok {
		actualNotPointer, ok := a.(JobPubsubTarget)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobPubsubTarget", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TopicName, actual.TopicName, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("TopicName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Data, actual.Data, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Data")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Attributes, actual.Attributes, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Attributes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobAppEngineHttpTargetNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobAppEngineHttpTarget)
	if !ok {
		desiredNotPointer, ok := d.(JobAppEngineHttpTarget)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobAppEngineHttpTarget or *JobAppEngineHttpTarget", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobAppEngineHttpTarget)
	if !ok {
		actualNotPointer, ok := a.(JobAppEngineHttpTarget)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobAppEngineHttpTarget", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HttpMethod, actual.HttpMethod, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("HttpMethod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AppEngineRouting, actual.AppEngineRouting, dcl.DiffInfo{ObjectFunction: compareJobAppEngineHttpTargetAppEngineRoutingNewStyle, EmptyObject: EmptyJobAppEngineHttpTargetAppEngineRouting, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("AppEngineRouting")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RelativeUri, actual.RelativeUri, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("RelativeUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Headers, actual.Headers, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Headers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Body, actual.Body, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Body")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobAppEngineHttpTargetAppEngineRoutingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobAppEngineHttpTargetAppEngineRouting)
	if !ok {
		desiredNotPointer, ok := d.(JobAppEngineHttpTargetAppEngineRouting)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobAppEngineHttpTargetAppEngineRouting or *JobAppEngineHttpTargetAppEngineRouting", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobAppEngineHttpTargetAppEngineRouting)
	if !ok {
		actualNotPointer, ok := a.(JobAppEngineHttpTargetAppEngineRouting)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobAppEngineHttpTargetAppEngineRouting", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Service, actual.Service, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Service")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Instance, actual.Instance, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Instance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Host, actual.Host, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Host")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobHttpTargetNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobHttpTarget)
	if !ok {
		desiredNotPointer, ok := d.(JobHttpTarget)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobHttpTarget or *JobHttpTarget", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobHttpTarget)
	if !ok {
		actualNotPointer, ok := a.(JobHttpTarget)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobHttpTarget", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Uri, actual.Uri, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Uri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HttpMethod, actual.HttpMethod, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("HttpMethod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Headers, actual.Headers, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Headers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Body, actual.Body, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Body")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OAuthToken, actual.OAuthToken, dcl.DiffInfo{ObjectFunction: compareJobHttpTargetOAuthTokenNewStyle, EmptyObject: EmptyJobHttpTargetOAuthToken, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("OauthToken")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OidcToken, actual.OidcToken, dcl.DiffInfo{ObjectFunction: compareJobHttpTargetOidcTokenNewStyle, EmptyObject: EmptyJobHttpTargetOidcToken, OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("OidcToken")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobHttpTargetOAuthTokenNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobHttpTargetOAuthToken)
	if !ok {
		desiredNotPointer, ok := d.(JobHttpTargetOAuthToken)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobHttpTargetOAuthToken or *JobHttpTargetOAuthToken", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobHttpTargetOAuthToken)
	if !ok {
		actualNotPointer, ok := a.(JobHttpTargetOAuthToken)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobHttpTargetOAuthToken", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ServiceAccountEmail, actual.ServiceAccountEmail, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("ServiceAccountEmail")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Scope, actual.Scope, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Scope")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobHttpTargetOidcTokenNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobHttpTargetOidcToken)
	if !ok {
		desiredNotPointer, ok := d.(JobHttpTargetOidcToken)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobHttpTargetOidcToken or *JobHttpTargetOidcToken", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobHttpTargetOidcToken)
	if !ok {
		actualNotPointer, ok := a.(JobHttpTargetOidcToken)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobHttpTargetOidcToken", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ServiceAccountEmail, actual.ServiceAccountEmail, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("ServiceAccountEmail")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Audience, actual.Audience, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("Audience")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobStatusNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobStatus)
	if !ok {
		desiredNotPointer, ok := d.(JobStatus)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobStatus or *JobStatus", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobStatus)
	if !ok {
		actualNotPointer, ok := a.(JobStatus)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobStatus", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Code, actual.Code, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Code")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Message, actual.Message, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Message")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Details, actual.Details, dcl.DiffInfo{ObjectFunction: compareJobStatusDetailsNewStyle, EmptyObject: EmptyJobStatusDetails, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Details")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobStatusDetailsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobStatusDetails)
	if !ok {
		desiredNotPointer, ok := d.(JobStatusDetails)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobStatusDetails or *JobStatusDetails", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobStatusDetails)
	if !ok {
		actualNotPointer, ok := a.(JobStatusDetails)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobStatusDetails", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TypeUrl, actual.TypeUrl, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TypeUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobRetryConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobRetryConfig)
	if !ok {
		desiredNotPointer, ok := d.(JobRetryConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobRetryConfig or *JobRetryConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobRetryConfig)
	if !ok {
		actualNotPointer, ok := a.(JobRetryConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobRetryConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RetryCount, actual.RetryCount, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("RetryCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxRetryDuration, actual.MaxRetryDuration, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("MaxRetryDuration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinBackoffDuration, actual.MinBackoffDuration, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("MinBackoffDuration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxBackoffDuration, actual.MaxBackoffDuration, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("MaxBackoffDuration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxDoublings, actual.MaxDoublings, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobUpdateJobOperation")}, fn.AddNest("MaxDoublings")); len(ds) != 0 || err != nil {
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
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Schedule = dcl.SelfLinkToName(r.Schedule)
	normalized.TimeZone = dcl.SelfLinkToName(r.TimeZone)
	normalized.AttemptDeadline = dcl.SelfLinkToName(r.AttemptDeadline)
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
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v, err := expandJobPubsubTarget(c, f.PubsubTarget, res); err != nil {
		return nil, fmt.Errorf("error expanding PubsubTarget into pubsubTarget: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["pubsubTarget"] = v
	}
	if v, err := expandJobAppEngineHttpTarget(c, f.AppEngineHttpTarget, res); err != nil {
		return nil, fmt.Errorf("error expanding AppEngineHttpTarget into appEngineHttpTarget: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["appEngineHttpTarget"] = v
	}
	if v, err := expandJobHttpTarget(c, f.HttpTarget, res); err != nil {
		return nil, fmt.Errorf("error expanding HttpTarget into httpTarget: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["httpTarget"] = v
	}
	if v := f.Schedule; dcl.ValueShouldBeSent(v) {
		m["schedule"] = v
	}
	if v := f.TimeZone; dcl.ValueShouldBeSent(v) {
		m["timeZone"] = v
	}
	if v, err := expandJobRetryConfig(c, f.RetryConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding RetryConfig into retryConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["retryConfig"] = v
	}
	if v := f.AttemptDeadline; dcl.ValueShouldBeSent(v) {
		m["attemptDeadline"] = v
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
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.PubsubTarget = flattenJobPubsubTarget(c, m["pubsubTarget"], res)
	resultRes.AppEngineHttpTarget = flattenJobAppEngineHttpTarget(c, m["appEngineHttpTarget"], res)
	resultRes.HttpTarget = flattenJobHttpTarget(c, m["httpTarget"], res)
	resultRes.Schedule = dcl.FlattenString(m["schedule"])
	resultRes.TimeZone = dcl.FlattenString(m["timeZone"])
	resultRes.UserUpdateTime = dcl.FlattenString(m["userUpdateTime"])
	resultRes.State = flattenJobStateEnum(m["state"])
	resultRes.Status = flattenJobStatus(c, m["status"], res)
	resultRes.ScheduleTime = dcl.FlattenString(m["scheduleTime"])
	resultRes.LastAttemptTime = dcl.FlattenString(m["lastAttemptTime"])
	resultRes.RetryConfig = flattenJobRetryConfig(c, m["retryConfig"], res)
	resultRes.AttemptDeadline = dcl.FlattenString(m["attemptDeadline"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// expandJobPubsubTargetMap expands the contents of JobPubsubTarget into a JSON
// request object.
func expandJobPubsubTargetMap(c *Client, f map[string]JobPubsubTarget, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobPubsubTarget(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobPubsubTargetSlice expands the contents of JobPubsubTarget into a JSON
// request object.
func expandJobPubsubTargetSlice(c *Client, f []JobPubsubTarget, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobPubsubTarget(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobPubsubTargetMap flattens the contents of JobPubsubTarget from a JSON
// response object.
func flattenJobPubsubTargetMap(c *Client, i interface{}, res *Job) map[string]JobPubsubTarget {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobPubsubTarget{}
	}

	if len(a) == 0 {
		return map[string]JobPubsubTarget{}
	}

	items := make(map[string]JobPubsubTarget)
	for k, item := range a {
		items[k] = *flattenJobPubsubTarget(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobPubsubTargetSlice flattens the contents of JobPubsubTarget from a JSON
// response object.
func flattenJobPubsubTargetSlice(c *Client, i interface{}, res *Job) []JobPubsubTarget {
	a, ok := i.([]interface{})
	if !ok {
		return []JobPubsubTarget{}
	}

	if len(a) == 0 {
		return []JobPubsubTarget{}
	}

	items := make([]JobPubsubTarget, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobPubsubTarget(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobPubsubTarget expands an instance of JobPubsubTarget into a JSON
// request object.
func expandJobPubsubTarget(c *Client, f *JobPubsubTarget, res *Job) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.TopicName; !dcl.IsEmptyValueIndirect(v) {
		m["topicName"] = v
	}
	if v := f.Data; !dcl.IsEmptyValueIndirect(v) {
		m["data"] = v
	}
	if v := f.Attributes; !dcl.IsEmptyValueIndirect(v) {
		m["attributes"] = v
	}

	return m, nil
}

// flattenJobPubsubTarget flattens an instance of JobPubsubTarget from a JSON
// response object.
func flattenJobPubsubTarget(c *Client, i interface{}, res *Job) *JobPubsubTarget {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobPubsubTarget{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobPubsubTarget
	}
	r.TopicName = dcl.FlattenString(m["topicName"])
	r.Data = dcl.FlattenString(m["data"])
	r.Attributes = dcl.FlattenKeyValuePairs(m["attributes"])

	return r
}

// expandJobAppEngineHttpTargetMap expands the contents of JobAppEngineHttpTarget into a JSON
// request object.
func expandJobAppEngineHttpTargetMap(c *Client, f map[string]JobAppEngineHttpTarget, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobAppEngineHttpTarget(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobAppEngineHttpTargetSlice expands the contents of JobAppEngineHttpTarget into a JSON
// request object.
func expandJobAppEngineHttpTargetSlice(c *Client, f []JobAppEngineHttpTarget, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobAppEngineHttpTarget(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobAppEngineHttpTargetMap flattens the contents of JobAppEngineHttpTarget from a JSON
// response object.
func flattenJobAppEngineHttpTargetMap(c *Client, i interface{}, res *Job) map[string]JobAppEngineHttpTarget {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobAppEngineHttpTarget{}
	}

	if len(a) == 0 {
		return map[string]JobAppEngineHttpTarget{}
	}

	items := make(map[string]JobAppEngineHttpTarget)
	for k, item := range a {
		items[k] = *flattenJobAppEngineHttpTarget(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobAppEngineHttpTargetSlice flattens the contents of JobAppEngineHttpTarget from a JSON
// response object.
func flattenJobAppEngineHttpTargetSlice(c *Client, i interface{}, res *Job) []JobAppEngineHttpTarget {
	a, ok := i.([]interface{})
	if !ok {
		return []JobAppEngineHttpTarget{}
	}

	if len(a) == 0 {
		return []JobAppEngineHttpTarget{}
	}

	items := make([]JobAppEngineHttpTarget, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobAppEngineHttpTarget(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobAppEngineHttpTarget expands an instance of JobAppEngineHttpTarget into a JSON
// request object.
func expandJobAppEngineHttpTarget(c *Client, f *JobAppEngineHttpTarget, res *Job) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.HttpMethod; !dcl.IsEmptyValueIndirect(v) {
		m["httpMethod"] = v
	}
	if v, err := expandJobAppEngineHttpTargetAppEngineRouting(c, f.AppEngineRouting, res); err != nil {
		return nil, fmt.Errorf("error expanding AppEngineRouting into appEngineRouting: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["appEngineRouting"] = v
	}
	if v := f.RelativeUri; !dcl.IsEmptyValueIndirect(v) {
		m["relativeUri"] = v
	}
	if v := f.Headers; !dcl.IsEmptyValueIndirect(v) {
		m["headers"] = v
	}
	if v := f.Body; !dcl.IsEmptyValueIndirect(v) {
		m["body"] = v
	}

	return m, nil
}

// flattenJobAppEngineHttpTarget flattens an instance of JobAppEngineHttpTarget from a JSON
// response object.
func flattenJobAppEngineHttpTarget(c *Client, i interface{}, res *Job) *JobAppEngineHttpTarget {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobAppEngineHttpTarget{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobAppEngineHttpTarget
	}
	r.HttpMethod = flattenJobAppEngineHttpTargetHttpMethodEnum(m["httpMethod"])
	r.AppEngineRouting = flattenJobAppEngineHttpTargetAppEngineRouting(c, m["appEngineRouting"], res)
	r.RelativeUri = dcl.FlattenString(m["relativeUri"])
	r.Headers = dcl.FlattenKeyValuePairs(m["headers"])
	r.Body = dcl.FlattenString(m["body"])

	return r
}

// expandJobAppEngineHttpTargetAppEngineRoutingMap expands the contents of JobAppEngineHttpTargetAppEngineRouting into a JSON
// request object.
func expandJobAppEngineHttpTargetAppEngineRoutingMap(c *Client, f map[string]JobAppEngineHttpTargetAppEngineRouting, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobAppEngineHttpTargetAppEngineRouting(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobAppEngineHttpTargetAppEngineRoutingSlice expands the contents of JobAppEngineHttpTargetAppEngineRouting into a JSON
// request object.
func expandJobAppEngineHttpTargetAppEngineRoutingSlice(c *Client, f []JobAppEngineHttpTargetAppEngineRouting, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobAppEngineHttpTargetAppEngineRouting(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobAppEngineHttpTargetAppEngineRoutingMap flattens the contents of JobAppEngineHttpTargetAppEngineRouting from a JSON
// response object.
func flattenJobAppEngineHttpTargetAppEngineRoutingMap(c *Client, i interface{}, res *Job) map[string]JobAppEngineHttpTargetAppEngineRouting {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobAppEngineHttpTargetAppEngineRouting{}
	}

	if len(a) == 0 {
		return map[string]JobAppEngineHttpTargetAppEngineRouting{}
	}

	items := make(map[string]JobAppEngineHttpTargetAppEngineRouting)
	for k, item := range a {
		items[k] = *flattenJobAppEngineHttpTargetAppEngineRouting(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobAppEngineHttpTargetAppEngineRoutingSlice flattens the contents of JobAppEngineHttpTargetAppEngineRouting from a JSON
// response object.
func flattenJobAppEngineHttpTargetAppEngineRoutingSlice(c *Client, i interface{}, res *Job) []JobAppEngineHttpTargetAppEngineRouting {
	a, ok := i.([]interface{})
	if !ok {
		return []JobAppEngineHttpTargetAppEngineRouting{}
	}

	if len(a) == 0 {
		return []JobAppEngineHttpTargetAppEngineRouting{}
	}

	items := make([]JobAppEngineHttpTargetAppEngineRouting, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobAppEngineHttpTargetAppEngineRouting(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobAppEngineHttpTargetAppEngineRouting expands an instance of JobAppEngineHttpTargetAppEngineRouting into a JSON
// request object.
func expandJobAppEngineHttpTargetAppEngineRouting(c *Client, f *JobAppEngineHttpTargetAppEngineRouting, res *Job) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Service; !dcl.IsEmptyValueIndirect(v) {
		m["service"] = v
	}
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}
	if v := f.Instance; !dcl.IsEmptyValueIndirect(v) {
		m["instance"] = v
	}

	return m, nil
}

// flattenJobAppEngineHttpTargetAppEngineRouting flattens an instance of JobAppEngineHttpTargetAppEngineRouting from a JSON
// response object.
func flattenJobAppEngineHttpTargetAppEngineRouting(c *Client, i interface{}, res *Job) *JobAppEngineHttpTargetAppEngineRouting {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobAppEngineHttpTargetAppEngineRouting{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobAppEngineHttpTargetAppEngineRouting
	}
	r.Service = dcl.FlattenString(m["service"])
	r.Version = dcl.FlattenString(m["version"])
	r.Instance = dcl.FlattenString(m["instance"])
	r.Host = dcl.FlattenString(m["host"])

	return r
}

// expandJobHttpTargetMap expands the contents of JobHttpTarget into a JSON
// request object.
func expandJobHttpTargetMap(c *Client, f map[string]JobHttpTarget, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobHttpTarget(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobHttpTargetSlice expands the contents of JobHttpTarget into a JSON
// request object.
func expandJobHttpTargetSlice(c *Client, f []JobHttpTarget, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobHttpTarget(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobHttpTargetMap flattens the contents of JobHttpTarget from a JSON
// response object.
func flattenJobHttpTargetMap(c *Client, i interface{}, res *Job) map[string]JobHttpTarget {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobHttpTarget{}
	}

	if len(a) == 0 {
		return map[string]JobHttpTarget{}
	}

	items := make(map[string]JobHttpTarget)
	for k, item := range a {
		items[k] = *flattenJobHttpTarget(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobHttpTargetSlice flattens the contents of JobHttpTarget from a JSON
// response object.
func flattenJobHttpTargetSlice(c *Client, i interface{}, res *Job) []JobHttpTarget {
	a, ok := i.([]interface{})
	if !ok {
		return []JobHttpTarget{}
	}

	if len(a) == 0 {
		return []JobHttpTarget{}
	}

	items := make([]JobHttpTarget, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobHttpTarget(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobHttpTarget expands an instance of JobHttpTarget into a JSON
// request object.
func expandJobHttpTarget(c *Client, f *JobHttpTarget, res *Job) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Uri; !dcl.IsEmptyValueIndirect(v) {
		m["uri"] = v
	}
	if v := f.HttpMethod; !dcl.IsEmptyValueIndirect(v) {
		m["httpMethod"] = v
	}
	if v := f.Headers; !dcl.IsEmptyValueIndirect(v) {
		m["headers"] = v
	}
	if v := f.Body; !dcl.IsEmptyValueIndirect(v) {
		m["body"] = v
	}
	if v, err := expandJobHttpTargetOAuthToken(c, f.OAuthToken, res); err != nil {
		return nil, fmt.Errorf("error expanding OAuthToken into oauthToken: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["oauthToken"] = v
	}
	if v, err := expandJobHttpTargetOidcToken(c, f.OidcToken, res); err != nil {
		return nil, fmt.Errorf("error expanding OidcToken into oidcToken: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["oidcToken"] = v
	}

	return m, nil
}

// flattenJobHttpTarget flattens an instance of JobHttpTarget from a JSON
// response object.
func flattenJobHttpTarget(c *Client, i interface{}, res *Job) *JobHttpTarget {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobHttpTarget{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobHttpTarget
	}
	r.Uri = dcl.FlattenString(m["uri"])
	r.HttpMethod = flattenJobHttpTargetHttpMethodEnum(m["httpMethod"])
	r.Headers = dcl.FlattenKeyValuePairs(m["headers"])
	r.Body = dcl.FlattenString(m["body"])
	r.OAuthToken = flattenJobHttpTargetOAuthToken(c, m["oauthToken"], res)
	r.OidcToken = flattenJobHttpTargetOidcToken(c, m["oidcToken"], res)

	return r
}

// expandJobHttpTargetOAuthTokenMap expands the contents of JobHttpTargetOAuthToken into a JSON
// request object.
func expandJobHttpTargetOAuthTokenMap(c *Client, f map[string]JobHttpTargetOAuthToken, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobHttpTargetOAuthToken(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobHttpTargetOAuthTokenSlice expands the contents of JobHttpTargetOAuthToken into a JSON
// request object.
func expandJobHttpTargetOAuthTokenSlice(c *Client, f []JobHttpTargetOAuthToken, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobHttpTargetOAuthToken(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobHttpTargetOAuthTokenMap flattens the contents of JobHttpTargetOAuthToken from a JSON
// response object.
func flattenJobHttpTargetOAuthTokenMap(c *Client, i interface{}, res *Job) map[string]JobHttpTargetOAuthToken {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobHttpTargetOAuthToken{}
	}

	if len(a) == 0 {
		return map[string]JobHttpTargetOAuthToken{}
	}

	items := make(map[string]JobHttpTargetOAuthToken)
	for k, item := range a {
		items[k] = *flattenJobHttpTargetOAuthToken(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobHttpTargetOAuthTokenSlice flattens the contents of JobHttpTargetOAuthToken from a JSON
// response object.
func flattenJobHttpTargetOAuthTokenSlice(c *Client, i interface{}, res *Job) []JobHttpTargetOAuthToken {
	a, ok := i.([]interface{})
	if !ok {
		return []JobHttpTargetOAuthToken{}
	}

	if len(a) == 0 {
		return []JobHttpTargetOAuthToken{}
	}

	items := make([]JobHttpTargetOAuthToken, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobHttpTargetOAuthToken(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobHttpTargetOAuthToken expands an instance of JobHttpTargetOAuthToken into a JSON
// request object.
func expandJobHttpTargetOAuthToken(c *Client, f *JobHttpTargetOAuthToken, res *Job) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ServiceAccountEmail; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccountEmail"] = v
	}
	if v := f.Scope; !dcl.IsEmptyValueIndirect(v) {
		m["scope"] = v
	}

	return m, nil
}

// flattenJobHttpTargetOAuthToken flattens an instance of JobHttpTargetOAuthToken from a JSON
// response object.
func flattenJobHttpTargetOAuthToken(c *Client, i interface{}, res *Job) *JobHttpTargetOAuthToken {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobHttpTargetOAuthToken{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobHttpTargetOAuthToken
	}
	r.ServiceAccountEmail = dcl.FlattenString(m["serviceAccountEmail"])
	r.Scope = dcl.FlattenString(m["scope"])

	return r
}

// expandJobHttpTargetOidcTokenMap expands the contents of JobHttpTargetOidcToken into a JSON
// request object.
func expandJobHttpTargetOidcTokenMap(c *Client, f map[string]JobHttpTargetOidcToken, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobHttpTargetOidcToken(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobHttpTargetOidcTokenSlice expands the contents of JobHttpTargetOidcToken into a JSON
// request object.
func expandJobHttpTargetOidcTokenSlice(c *Client, f []JobHttpTargetOidcToken, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobHttpTargetOidcToken(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobHttpTargetOidcTokenMap flattens the contents of JobHttpTargetOidcToken from a JSON
// response object.
func flattenJobHttpTargetOidcTokenMap(c *Client, i interface{}, res *Job) map[string]JobHttpTargetOidcToken {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobHttpTargetOidcToken{}
	}

	if len(a) == 0 {
		return map[string]JobHttpTargetOidcToken{}
	}

	items := make(map[string]JobHttpTargetOidcToken)
	for k, item := range a {
		items[k] = *flattenJobHttpTargetOidcToken(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobHttpTargetOidcTokenSlice flattens the contents of JobHttpTargetOidcToken from a JSON
// response object.
func flattenJobHttpTargetOidcTokenSlice(c *Client, i interface{}, res *Job) []JobHttpTargetOidcToken {
	a, ok := i.([]interface{})
	if !ok {
		return []JobHttpTargetOidcToken{}
	}

	if len(a) == 0 {
		return []JobHttpTargetOidcToken{}
	}

	items := make([]JobHttpTargetOidcToken, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobHttpTargetOidcToken(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobHttpTargetOidcToken expands an instance of JobHttpTargetOidcToken into a JSON
// request object.
func expandJobHttpTargetOidcToken(c *Client, f *JobHttpTargetOidcToken, res *Job) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ServiceAccountEmail; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccountEmail"] = v
	}
	if v := f.Audience; !dcl.IsEmptyValueIndirect(v) {
		m["audience"] = v
	}

	return m, nil
}

// flattenJobHttpTargetOidcToken flattens an instance of JobHttpTargetOidcToken from a JSON
// response object.
func flattenJobHttpTargetOidcToken(c *Client, i interface{}, res *Job) *JobHttpTargetOidcToken {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobHttpTargetOidcToken{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobHttpTargetOidcToken
	}
	r.ServiceAccountEmail = dcl.FlattenString(m["serviceAccountEmail"])
	r.Audience = dcl.FlattenString(m["audience"])

	return r
}

// expandJobStatusMap expands the contents of JobStatus into a JSON
// request object.
func expandJobStatusMap(c *Client, f map[string]JobStatus, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobStatus(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobStatusSlice expands the contents of JobStatus into a JSON
// request object.
func expandJobStatusSlice(c *Client, f []JobStatus, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobStatus(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobStatusMap flattens the contents of JobStatus from a JSON
// response object.
func flattenJobStatusMap(c *Client, i interface{}, res *Job) map[string]JobStatus {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobStatus{}
	}

	if len(a) == 0 {
		return map[string]JobStatus{}
	}

	items := make(map[string]JobStatus)
	for k, item := range a {
		items[k] = *flattenJobStatus(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobStatusSlice flattens the contents of JobStatus from a JSON
// response object.
func flattenJobStatusSlice(c *Client, i interface{}, res *Job) []JobStatus {
	a, ok := i.([]interface{})
	if !ok {
		return []JobStatus{}
	}

	if len(a) == 0 {
		return []JobStatus{}
	}

	items := make([]JobStatus, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobStatus(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobStatus expands an instance of JobStatus into a JSON
// request object.
func expandJobStatus(c *Client, f *JobStatus, res *Job) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Code; !dcl.IsEmptyValueIndirect(v) {
		m["code"] = v
	}
	if v := f.Message; !dcl.IsEmptyValueIndirect(v) {
		m["message"] = v
	}
	if v, err := expandJobStatusDetailsSlice(c, f.Details, res); err != nil {
		return nil, fmt.Errorf("error expanding Details into details: %w", err)
	} else if v != nil {
		m["details"] = v
	}

	return m, nil
}

// flattenJobStatus flattens an instance of JobStatus from a JSON
// response object.
func flattenJobStatus(c *Client, i interface{}, res *Job) *JobStatus {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobStatus{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobStatus
	}
	r.Code = dcl.FlattenInteger(m["code"])
	r.Message = dcl.FlattenString(m["message"])
	r.Details = flattenJobStatusDetailsSlice(c, m["details"], res)

	return r
}

// expandJobStatusDetailsMap expands the contents of JobStatusDetails into a JSON
// request object.
func expandJobStatusDetailsMap(c *Client, f map[string]JobStatusDetails, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobStatusDetails(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobStatusDetailsSlice expands the contents of JobStatusDetails into a JSON
// request object.
func expandJobStatusDetailsSlice(c *Client, f []JobStatusDetails, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobStatusDetails(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobStatusDetailsMap flattens the contents of JobStatusDetails from a JSON
// response object.
func flattenJobStatusDetailsMap(c *Client, i interface{}, res *Job) map[string]JobStatusDetails {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobStatusDetails{}
	}

	if len(a) == 0 {
		return map[string]JobStatusDetails{}
	}

	items := make(map[string]JobStatusDetails)
	for k, item := range a {
		items[k] = *flattenJobStatusDetails(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobStatusDetailsSlice flattens the contents of JobStatusDetails from a JSON
// response object.
func flattenJobStatusDetailsSlice(c *Client, i interface{}, res *Job) []JobStatusDetails {
	a, ok := i.([]interface{})
	if !ok {
		return []JobStatusDetails{}
	}

	if len(a) == 0 {
		return []JobStatusDetails{}
	}

	items := make([]JobStatusDetails, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobStatusDetails(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobStatusDetails expands an instance of JobStatusDetails into a JSON
// request object.
func expandJobStatusDetails(c *Client, f *JobStatusDetails, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.TypeUrl; !dcl.IsEmptyValueIndirect(v) {
		m["typeUrl"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenJobStatusDetails flattens an instance of JobStatusDetails from a JSON
// response object.
func flattenJobStatusDetails(c *Client, i interface{}, res *Job) *JobStatusDetails {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobStatusDetails{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobStatusDetails
	}
	r.TypeUrl = dcl.FlattenString(m["typeUrl"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// expandJobRetryConfigMap expands the contents of JobRetryConfig into a JSON
// request object.
func expandJobRetryConfigMap(c *Client, f map[string]JobRetryConfig, res *Job) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobRetryConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobRetryConfigSlice expands the contents of JobRetryConfig into a JSON
// request object.
func expandJobRetryConfigSlice(c *Client, f []JobRetryConfig, res *Job) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobRetryConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobRetryConfigMap flattens the contents of JobRetryConfig from a JSON
// response object.
func flattenJobRetryConfigMap(c *Client, i interface{}, res *Job) map[string]JobRetryConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobRetryConfig{}
	}

	if len(a) == 0 {
		return map[string]JobRetryConfig{}
	}

	items := make(map[string]JobRetryConfig)
	for k, item := range a {
		items[k] = *flattenJobRetryConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobRetryConfigSlice flattens the contents of JobRetryConfig from a JSON
// response object.
func flattenJobRetryConfigSlice(c *Client, i interface{}, res *Job) []JobRetryConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []JobRetryConfig{}
	}

	if len(a) == 0 {
		return []JobRetryConfig{}
	}

	items := make([]JobRetryConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobRetryConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobRetryConfig expands an instance of JobRetryConfig into a JSON
// request object.
func expandJobRetryConfig(c *Client, f *JobRetryConfig, res *Job) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RetryCount; !dcl.IsEmptyValueIndirect(v) {
		m["retryCount"] = v
	}
	if v := f.MaxRetryDuration; !dcl.IsEmptyValueIndirect(v) {
		m["maxRetryDuration"] = v
	}
	if v := f.MinBackoffDuration; !dcl.IsEmptyValueIndirect(v) {
		m["minBackoffDuration"] = v
	}
	if v := f.MaxBackoffDuration; !dcl.IsEmptyValueIndirect(v) {
		m["maxBackoffDuration"] = v
	}
	if v := f.MaxDoublings; !dcl.IsEmptyValueIndirect(v) {
		m["maxDoublings"] = v
	}

	return m, nil
}

// flattenJobRetryConfig flattens an instance of JobRetryConfig from a JSON
// response object.
func flattenJobRetryConfig(c *Client, i interface{}, res *Job) *JobRetryConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobRetryConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobRetryConfig
	}
	r.RetryCount = dcl.FlattenInteger(m["retryCount"])
	r.MaxRetryDuration = dcl.FlattenString(m["maxRetryDuration"])
	r.MinBackoffDuration = dcl.FlattenString(m["minBackoffDuration"])
	r.MaxBackoffDuration = dcl.FlattenString(m["maxBackoffDuration"])
	r.MaxDoublings = dcl.FlattenInteger(m["maxDoublings"])

	return r
}

// flattenJobAppEngineHttpTargetHttpMethodEnumMap flattens the contents of JobAppEngineHttpTargetHttpMethodEnum from a JSON
// response object.
func flattenJobAppEngineHttpTargetHttpMethodEnumMap(c *Client, i interface{}, res *Job) map[string]JobAppEngineHttpTargetHttpMethodEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobAppEngineHttpTargetHttpMethodEnum{}
	}

	if len(a) == 0 {
		return map[string]JobAppEngineHttpTargetHttpMethodEnum{}
	}

	items := make(map[string]JobAppEngineHttpTargetHttpMethodEnum)
	for k, item := range a {
		items[k] = *flattenJobAppEngineHttpTargetHttpMethodEnum(item.(interface{}))
	}

	return items
}

// flattenJobAppEngineHttpTargetHttpMethodEnumSlice flattens the contents of JobAppEngineHttpTargetHttpMethodEnum from a JSON
// response object.
func flattenJobAppEngineHttpTargetHttpMethodEnumSlice(c *Client, i interface{}, res *Job) []JobAppEngineHttpTargetHttpMethodEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobAppEngineHttpTargetHttpMethodEnum{}
	}

	if len(a) == 0 {
		return []JobAppEngineHttpTargetHttpMethodEnum{}
	}

	items := make([]JobAppEngineHttpTargetHttpMethodEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobAppEngineHttpTargetHttpMethodEnum(item.(interface{})))
	}

	return items
}

// flattenJobAppEngineHttpTargetHttpMethodEnum asserts that an interface is a string, and returns a
// pointer to a *JobAppEngineHttpTargetHttpMethodEnum with the same value as that string.
func flattenJobAppEngineHttpTargetHttpMethodEnum(i interface{}) *JobAppEngineHttpTargetHttpMethodEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobAppEngineHttpTargetHttpMethodEnumRef(s)
}

// flattenJobHttpTargetHttpMethodEnumMap flattens the contents of JobHttpTargetHttpMethodEnum from a JSON
// response object.
func flattenJobHttpTargetHttpMethodEnumMap(c *Client, i interface{}, res *Job) map[string]JobHttpTargetHttpMethodEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobHttpTargetHttpMethodEnum{}
	}

	if len(a) == 0 {
		return map[string]JobHttpTargetHttpMethodEnum{}
	}

	items := make(map[string]JobHttpTargetHttpMethodEnum)
	for k, item := range a {
		items[k] = *flattenJobHttpTargetHttpMethodEnum(item.(interface{}))
	}

	return items
}

// flattenJobHttpTargetHttpMethodEnumSlice flattens the contents of JobHttpTargetHttpMethodEnum from a JSON
// response object.
func flattenJobHttpTargetHttpMethodEnumSlice(c *Client, i interface{}, res *Job) []JobHttpTargetHttpMethodEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobHttpTargetHttpMethodEnum{}
	}

	if len(a) == 0 {
		return []JobHttpTargetHttpMethodEnum{}
	}

	items := make([]JobHttpTargetHttpMethodEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobHttpTargetHttpMethodEnum(item.(interface{})))
	}

	return items
}

// flattenJobHttpTargetHttpMethodEnum asserts that an interface is a string, and returns a
// pointer to a *JobHttpTargetHttpMethodEnum with the same value as that string.
func flattenJobHttpTargetHttpMethodEnum(i interface{}) *JobHttpTargetHttpMethodEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobHttpTargetHttpMethodEnumRef(s)
}

// flattenJobStateEnumMap flattens the contents of JobStateEnum from a JSON
// response object.
func flattenJobStateEnumMap(c *Client, i interface{}, res *Job) map[string]JobStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobStateEnum{}
	}

	if len(a) == 0 {
		return map[string]JobStateEnum{}
	}

	items := make(map[string]JobStateEnum)
	for k, item := range a {
		items[k] = *flattenJobStateEnum(item.(interface{}))
	}

	return items
}

// flattenJobStateEnumSlice flattens the contents of JobStateEnum from a JSON
// response object.
func flattenJobStateEnumSlice(c *Client, i interface{}, res *Job) []JobStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobStateEnum{}
	}

	if len(a) == 0 {
		return []JobStateEnum{}
	}

	items := make([]JobStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobStateEnum(item.(interface{})))
	}

	return items
}

// flattenJobStateEnum asserts that an interface is a string, and returns a
// pointer to a *JobStateEnum with the same value as that string.
func flattenJobStateEnum(i interface{}) *JobStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobStateEnumRef(s)
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
	vPubsubTarget := r.PubsubTarget
	if vPubsubTarget == nil {
		// note: explicitly not the empty object.
		vPubsubTarget = &JobPubsubTarget{}
	}
	if err := extractJobPubsubTargetFields(r, vPubsubTarget); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPubsubTarget) {
		r.PubsubTarget = vPubsubTarget
	}
	vAppEngineHttpTarget := r.AppEngineHttpTarget
	if vAppEngineHttpTarget == nil {
		// note: explicitly not the empty object.
		vAppEngineHttpTarget = &JobAppEngineHttpTarget{}
	}
	if err := extractJobAppEngineHttpTargetFields(r, vAppEngineHttpTarget); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAppEngineHttpTarget) {
		r.AppEngineHttpTarget = vAppEngineHttpTarget
	}
	vHttpTarget := r.HttpTarget
	if vHttpTarget == nil {
		// note: explicitly not the empty object.
		vHttpTarget = &JobHttpTarget{}
	}
	if err := extractJobHttpTargetFields(r, vHttpTarget); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vHttpTarget) {
		r.HttpTarget = vHttpTarget
	}
	vStatus := r.Status
	if vStatus == nil {
		// note: explicitly not the empty object.
		vStatus = &JobStatus{}
	}
	if err := extractJobStatusFields(r, vStatus); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vStatus) {
		r.Status = vStatus
	}
	vRetryConfig := r.RetryConfig
	if vRetryConfig == nil {
		// note: explicitly not the empty object.
		vRetryConfig = &JobRetryConfig{}
	}
	if err := extractJobRetryConfigFields(r, vRetryConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRetryConfig) {
		r.RetryConfig = vRetryConfig
	}
	return nil
}
func extractJobPubsubTargetFields(r *Job, o *JobPubsubTarget) error {
	return nil
}
func extractJobAppEngineHttpTargetFields(r *Job, o *JobAppEngineHttpTarget) error {
	vAppEngineRouting := o.AppEngineRouting
	if vAppEngineRouting == nil {
		// note: explicitly not the empty object.
		vAppEngineRouting = &JobAppEngineHttpTargetAppEngineRouting{}
	}
	if err := extractJobAppEngineHttpTargetAppEngineRoutingFields(r, vAppEngineRouting); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAppEngineRouting) {
		o.AppEngineRouting = vAppEngineRouting
	}
	return nil
}
func extractJobAppEngineHttpTargetAppEngineRoutingFields(r *Job, o *JobAppEngineHttpTargetAppEngineRouting) error {
	return nil
}
func extractJobHttpTargetFields(r *Job, o *JobHttpTarget) error {
	vOAuthToken := o.OAuthToken
	if vOAuthToken == nil {
		// note: explicitly not the empty object.
		vOAuthToken = &JobHttpTargetOAuthToken{}
	}
	if err := extractJobHttpTargetOAuthTokenFields(r, vOAuthToken); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vOAuthToken) {
		o.OAuthToken = vOAuthToken
	}
	vOidcToken := o.OidcToken
	if vOidcToken == nil {
		// note: explicitly not the empty object.
		vOidcToken = &JobHttpTargetOidcToken{}
	}
	if err := extractJobHttpTargetOidcTokenFields(r, vOidcToken); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vOidcToken) {
		o.OidcToken = vOidcToken
	}
	return nil
}
func extractJobHttpTargetOAuthTokenFields(r *Job, o *JobHttpTargetOAuthToken) error {
	return nil
}
func extractJobHttpTargetOidcTokenFields(r *Job, o *JobHttpTargetOidcToken) error {
	return nil
}
func extractJobStatusFields(r *Job, o *JobStatus) error {
	return nil
}
func extractJobStatusDetailsFields(r *Job, o *JobStatusDetails) error {
	return nil
}
func extractJobRetryConfigFields(r *Job, o *JobRetryConfig) error {
	return nil
}

func postReadExtractJobFields(r *Job) error {
	vPubsubTarget := r.PubsubTarget
	if vPubsubTarget == nil {
		// note: explicitly not the empty object.
		vPubsubTarget = &JobPubsubTarget{}
	}
	if err := postReadExtractJobPubsubTargetFields(r, vPubsubTarget); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPubsubTarget) {
		r.PubsubTarget = vPubsubTarget
	}
	vAppEngineHttpTarget := r.AppEngineHttpTarget
	if vAppEngineHttpTarget == nil {
		// note: explicitly not the empty object.
		vAppEngineHttpTarget = &JobAppEngineHttpTarget{}
	}
	if err := postReadExtractJobAppEngineHttpTargetFields(r, vAppEngineHttpTarget); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAppEngineHttpTarget) {
		r.AppEngineHttpTarget = vAppEngineHttpTarget
	}
	vHttpTarget := r.HttpTarget
	if vHttpTarget == nil {
		// note: explicitly not the empty object.
		vHttpTarget = &JobHttpTarget{}
	}
	if err := postReadExtractJobHttpTargetFields(r, vHttpTarget); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vHttpTarget) {
		r.HttpTarget = vHttpTarget
	}
	vStatus := r.Status
	if vStatus == nil {
		// note: explicitly not the empty object.
		vStatus = &JobStatus{}
	}
	if err := postReadExtractJobStatusFields(r, vStatus); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vStatus) {
		r.Status = vStatus
	}
	vRetryConfig := r.RetryConfig
	if vRetryConfig == nil {
		// note: explicitly not the empty object.
		vRetryConfig = &JobRetryConfig{}
	}
	if err := postReadExtractJobRetryConfigFields(r, vRetryConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRetryConfig) {
		r.RetryConfig = vRetryConfig
	}
	return nil
}
func postReadExtractJobPubsubTargetFields(r *Job, o *JobPubsubTarget) error {
	return nil
}
func postReadExtractJobAppEngineHttpTargetFields(r *Job, o *JobAppEngineHttpTarget) error {
	vAppEngineRouting := o.AppEngineRouting
	if vAppEngineRouting == nil {
		// note: explicitly not the empty object.
		vAppEngineRouting = &JobAppEngineHttpTargetAppEngineRouting{}
	}
	if err := extractJobAppEngineHttpTargetAppEngineRoutingFields(r, vAppEngineRouting); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAppEngineRouting) {
		o.AppEngineRouting = vAppEngineRouting
	}
	return nil
}
func postReadExtractJobAppEngineHttpTargetAppEngineRoutingFields(r *Job, o *JobAppEngineHttpTargetAppEngineRouting) error {
	return nil
}
func postReadExtractJobHttpTargetFields(r *Job, o *JobHttpTarget) error {
	vOAuthToken := o.OAuthToken
	if vOAuthToken == nil {
		// note: explicitly not the empty object.
		vOAuthToken = &JobHttpTargetOAuthToken{}
	}
	if err := extractJobHttpTargetOAuthTokenFields(r, vOAuthToken); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vOAuthToken) {
		o.OAuthToken = vOAuthToken
	}
	vOidcToken := o.OidcToken
	if vOidcToken == nil {
		// note: explicitly not the empty object.
		vOidcToken = &JobHttpTargetOidcToken{}
	}
	if err := extractJobHttpTargetOidcTokenFields(r, vOidcToken); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vOidcToken) {
		o.OidcToken = vOidcToken
	}
	return nil
}
func postReadExtractJobHttpTargetOAuthTokenFields(r *Job, o *JobHttpTargetOAuthToken) error {
	return nil
}
func postReadExtractJobHttpTargetOidcTokenFields(r *Job, o *JobHttpTargetOidcToken) error {
	return nil
}
func postReadExtractJobStatusFields(r *Job, o *JobStatus) error {
	return nil
}
func postReadExtractJobStatusDetailsFields(r *Job, o *JobStatusDetails) error {
	return nil
}
func postReadExtractJobRetryConfigFields(r *Job, o *JobRetryConfig) error {
	return nil
}
