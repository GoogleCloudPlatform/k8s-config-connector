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
package cloudscheduler

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Job struct {
	Name                *string                 `json:"name"`
	Description         *string                 `json:"description"`
	PubsubTarget        *JobPubsubTarget        `json:"pubsubTarget"`
	AppEngineHttpTarget *JobAppEngineHttpTarget `json:"appEngineHttpTarget"`
	HttpTarget          *JobHttpTarget          `json:"httpTarget"`
	Schedule            *string                 `json:"schedule"`
	TimeZone            *string                 `json:"timeZone"`
	UserUpdateTime      *string                 `json:"userUpdateTime"`
	State               *JobStateEnum           `json:"state"`
	Status              *JobStatus              `json:"status"`
	ScheduleTime        *string                 `json:"scheduleTime"`
	LastAttemptTime     *string                 `json:"lastAttemptTime"`
	RetryConfig         *JobRetryConfig         `json:"retryConfig"`
	AttemptDeadline     *string                 `json:"attemptDeadline"`
	Project             *string                 `json:"project"`
	Location            *string                 `json:"location"`
}

func (r *Job) String() string {
	return dcl.SprintResource(r)
}

// The enum JobAppEngineHttpTargetHttpMethodEnum.
type JobAppEngineHttpTargetHttpMethodEnum string

// JobAppEngineHttpTargetHttpMethodEnumRef returns a *JobAppEngineHttpTargetHttpMethodEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobAppEngineHttpTargetHttpMethodEnumRef(s string) *JobAppEngineHttpTargetHttpMethodEnum {
	v := JobAppEngineHttpTargetHttpMethodEnum(s)
	return &v
}

func (v JobAppEngineHttpTargetHttpMethodEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"HTTP_METHOD_UNSPECIFIED", "POST", "GET", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobAppEngineHttpTargetHttpMethodEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobHttpTargetHttpMethodEnum.
type JobHttpTargetHttpMethodEnum string

// JobHttpTargetHttpMethodEnumRef returns a *JobHttpTargetHttpMethodEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobHttpTargetHttpMethodEnumRef(s string) *JobHttpTargetHttpMethodEnum {
	v := JobHttpTargetHttpMethodEnum(s)
	return &v
}

func (v JobHttpTargetHttpMethodEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"HTTP_METHOD_UNSPECIFIED", "POST", "GET", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobHttpTargetHttpMethodEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobStateEnum.
type JobStateEnum string

// JobStateEnumRef returns a *JobStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobStateEnumRef(s string) *JobStateEnum {
	v := JobStateEnum(s)
	return &v
}

func (v JobStateEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"STATE_UNSPECIFIED", "ENABLED", "PAUSED", "DISABLED", "UPDATE_FAILED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type JobPubsubTarget struct {
	empty      bool              `json:"-"`
	TopicName  *string           `json:"topicName"`
	Data       *string           `json:"data"`
	Attributes map[string]string `json:"attributes"`
}

type jsonJobPubsubTarget JobPubsubTarget

func (r *JobPubsubTarget) UnmarshalJSON(data []byte) error {
	var res jsonJobPubsubTarget
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobPubsubTarget
	} else {

		r.TopicName = res.TopicName

		r.Data = res.Data

		r.Attributes = res.Attributes

	}
	return nil
}

// This object is used to assert a desired state where this JobPubsubTarget is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobPubsubTarget *JobPubsubTarget = &JobPubsubTarget{empty: true}

func (r *JobPubsubTarget) Empty() bool {
	return r.empty
}

func (r *JobPubsubTarget) String() string {
	return dcl.SprintResource(r)
}

func (r *JobPubsubTarget) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobAppEngineHttpTarget struct {
	empty            bool                                    `json:"-"`
	HttpMethod       *JobAppEngineHttpTargetHttpMethodEnum   `json:"httpMethod"`
	AppEngineRouting *JobAppEngineHttpTargetAppEngineRouting `json:"appEngineRouting"`
	RelativeUri      *string                                 `json:"relativeUri"`
	Headers          map[string]string                       `json:"headers"`
	Body             *string                                 `json:"body"`
}

type jsonJobAppEngineHttpTarget JobAppEngineHttpTarget

func (r *JobAppEngineHttpTarget) UnmarshalJSON(data []byte) error {
	var res jsonJobAppEngineHttpTarget
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobAppEngineHttpTarget
	} else {

		r.HttpMethod = res.HttpMethod

		r.AppEngineRouting = res.AppEngineRouting

		r.RelativeUri = res.RelativeUri

		r.Headers = res.Headers

		r.Body = res.Body

	}
	return nil
}

// This object is used to assert a desired state where this JobAppEngineHttpTarget is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobAppEngineHttpTarget *JobAppEngineHttpTarget = &JobAppEngineHttpTarget{empty: true}

func (r *JobAppEngineHttpTarget) Empty() bool {
	return r.empty
}

func (r *JobAppEngineHttpTarget) String() string {
	return dcl.SprintResource(r)
}

func (r *JobAppEngineHttpTarget) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobAppEngineHttpTargetAppEngineRouting struct {
	empty    bool    `json:"-"`
	Service  *string `json:"service"`
	Version  *string `json:"version"`
	Instance *string `json:"instance"`
	Host     *string `json:"host"`
}

type jsonJobAppEngineHttpTargetAppEngineRouting JobAppEngineHttpTargetAppEngineRouting

func (r *JobAppEngineHttpTargetAppEngineRouting) UnmarshalJSON(data []byte) error {
	var res jsonJobAppEngineHttpTargetAppEngineRouting
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobAppEngineHttpTargetAppEngineRouting
	} else {

		r.Service = res.Service

		r.Version = res.Version

		r.Instance = res.Instance

		r.Host = res.Host

	}
	return nil
}

// This object is used to assert a desired state where this JobAppEngineHttpTargetAppEngineRouting is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobAppEngineHttpTargetAppEngineRouting *JobAppEngineHttpTargetAppEngineRouting = &JobAppEngineHttpTargetAppEngineRouting{empty: true}

func (r *JobAppEngineHttpTargetAppEngineRouting) Empty() bool {
	return r.empty
}

func (r *JobAppEngineHttpTargetAppEngineRouting) String() string {
	return dcl.SprintResource(r)
}

func (r *JobAppEngineHttpTargetAppEngineRouting) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobHttpTarget struct {
	empty      bool                         `json:"-"`
	Uri        *string                      `json:"uri"`
	HttpMethod *JobHttpTargetHttpMethodEnum `json:"httpMethod"`
	Headers    map[string]string            `json:"headers"`
	Body       *string                      `json:"body"`
	OAuthToken *JobHttpTargetOAuthToken     `json:"oauthToken"`
	OidcToken  *JobHttpTargetOidcToken      `json:"oidcToken"`
}

type jsonJobHttpTarget JobHttpTarget

func (r *JobHttpTarget) UnmarshalJSON(data []byte) error {
	var res jsonJobHttpTarget
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobHttpTarget
	} else {

		r.Uri = res.Uri

		r.HttpMethod = res.HttpMethod

		r.Headers = res.Headers

		r.Body = res.Body

		r.OAuthToken = res.OAuthToken

		r.OidcToken = res.OidcToken

	}
	return nil
}

// This object is used to assert a desired state where this JobHttpTarget is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobHttpTarget *JobHttpTarget = &JobHttpTarget{empty: true}

func (r *JobHttpTarget) Empty() bool {
	return r.empty
}

func (r *JobHttpTarget) String() string {
	return dcl.SprintResource(r)
}

func (r *JobHttpTarget) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobHttpTargetOAuthToken struct {
	empty               bool    `json:"-"`
	ServiceAccountEmail *string `json:"serviceAccountEmail"`
	Scope               *string `json:"scope"`
}

type jsonJobHttpTargetOAuthToken JobHttpTargetOAuthToken

func (r *JobHttpTargetOAuthToken) UnmarshalJSON(data []byte) error {
	var res jsonJobHttpTargetOAuthToken
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobHttpTargetOAuthToken
	} else {

		r.ServiceAccountEmail = res.ServiceAccountEmail

		r.Scope = res.Scope

	}
	return nil
}

// This object is used to assert a desired state where this JobHttpTargetOAuthToken is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobHttpTargetOAuthToken *JobHttpTargetOAuthToken = &JobHttpTargetOAuthToken{empty: true}

func (r *JobHttpTargetOAuthToken) Empty() bool {
	return r.empty
}

func (r *JobHttpTargetOAuthToken) String() string {
	return dcl.SprintResource(r)
}

func (r *JobHttpTargetOAuthToken) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobHttpTargetOidcToken struct {
	empty               bool    `json:"-"`
	ServiceAccountEmail *string `json:"serviceAccountEmail"`
	Audience            *string `json:"audience"`
}

type jsonJobHttpTargetOidcToken JobHttpTargetOidcToken

func (r *JobHttpTargetOidcToken) UnmarshalJSON(data []byte) error {
	var res jsonJobHttpTargetOidcToken
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobHttpTargetOidcToken
	} else {

		r.ServiceAccountEmail = res.ServiceAccountEmail

		r.Audience = res.Audience

	}
	return nil
}

// This object is used to assert a desired state where this JobHttpTargetOidcToken is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobHttpTargetOidcToken *JobHttpTargetOidcToken = &JobHttpTargetOidcToken{empty: true}

func (r *JobHttpTargetOidcToken) Empty() bool {
	return r.empty
}

func (r *JobHttpTargetOidcToken) String() string {
	return dcl.SprintResource(r)
}

func (r *JobHttpTargetOidcToken) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobStatus struct {
	empty   bool               `json:"-"`
	Code    *int64             `json:"code"`
	Message *string            `json:"message"`
	Details []JobStatusDetails `json:"details"`
}

type jsonJobStatus JobStatus

func (r *JobStatus) UnmarshalJSON(data []byte) error {
	var res jsonJobStatus
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobStatus
	} else {

		r.Code = res.Code

		r.Message = res.Message

		r.Details = res.Details

	}
	return nil
}

// This object is used to assert a desired state where this JobStatus is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobStatus *JobStatus = &JobStatus{empty: true}

func (r *JobStatus) Empty() bool {
	return r.empty
}

func (r *JobStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *JobStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

type jsonJobStatusDetails JobStatusDetails

func (r *JobStatusDetails) UnmarshalJSON(data []byte) error {
	var res jsonJobStatusDetails
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobStatusDetails
	} else {

		r.TypeUrl = res.TypeUrl

		r.Value = res.Value

	}
	return nil
}

// This object is used to assert a desired state where this JobStatusDetails is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobStatusDetails *JobStatusDetails = &JobStatusDetails{empty: true}

func (r *JobStatusDetails) Empty() bool {
	return r.empty
}

func (r *JobStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *JobStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobRetryConfig struct {
	empty              bool    `json:"-"`
	RetryCount         *int64  `json:"retryCount"`
	MaxRetryDuration   *string `json:"maxRetryDuration"`
	MinBackoffDuration *string `json:"minBackoffDuration"`
	MaxBackoffDuration *string `json:"maxBackoffDuration"`
	MaxDoublings       *int64  `json:"maxDoublings"`
}

type jsonJobRetryConfig JobRetryConfig

func (r *JobRetryConfig) UnmarshalJSON(data []byte) error {
	var res jsonJobRetryConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobRetryConfig
	} else {

		r.RetryCount = res.RetryCount

		r.MaxRetryDuration = res.MaxRetryDuration

		r.MinBackoffDuration = res.MinBackoffDuration

		r.MaxBackoffDuration = res.MaxBackoffDuration

		r.MaxDoublings = res.MaxDoublings

	}
	return nil
}

// This object is used to assert a desired state where this JobRetryConfig is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobRetryConfig *JobRetryConfig = &JobRetryConfig{empty: true}

func (r *JobRetryConfig) Empty() bool {
	return r.empty
}

func (r *JobRetryConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *JobRetryConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Job) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "cloud_scheduler",
		Type:    "Job",
		Version: "cloudscheduler",
	}
}

func (r *Job) ID() (string, error) {
	if err := extractJobFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                   dcl.ValueOrEmptyString(nr.Name),
		"description":            dcl.ValueOrEmptyString(nr.Description),
		"pubsub_target":          dcl.ValueOrEmptyString(nr.PubsubTarget),
		"app_engine_http_target": dcl.ValueOrEmptyString(nr.AppEngineHttpTarget),
		"http_target":            dcl.ValueOrEmptyString(nr.HttpTarget),
		"schedule":               dcl.ValueOrEmptyString(nr.Schedule),
		"time_zone":              dcl.ValueOrEmptyString(nr.TimeZone),
		"user_update_time":       dcl.ValueOrEmptyString(nr.UserUpdateTime),
		"state":                  dcl.ValueOrEmptyString(nr.State),
		"status":                 dcl.ValueOrEmptyString(nr.Status),
		"schedule_time":          dcl.ValueOrEmptyString(nr.ScheduleTime),
		"last_attempt_time":      dcl.ValueOrEmptyString(nr.LastAttemptTime),
		"retry_config":           dcl.ValueOrEmptyString(nr.RetryConfig),
		"attempt_deadline":       dcl.ValueOrEmptyString(nr.AttemptDeadline),
		"project":                dcl.ValueOrEmptyString(nr.Project),
		"location":               dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/jobs/{{name}}", params), nil
}

const JobMaxPage = -1

type JobList struct {
	Items []*Job

	nextToken string

	pageSize int32

	resource *Job
}

func (l *JobList) HasNext() bool {
	return l.nextToken != ""
}

func (l *JobList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listJob(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListJob(ctx context.Context, project, location string) (*JobList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListJobWithMaxResults(ctx, project, location, JobMaxPage)

}

func (c *Client) ListJobWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*JobList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &Job{
		Project:  &project,
		Location: &location,
	}
	items, token, err := c.listJob(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &JobList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetJob(ctx context.Context, r *Job) (*Job, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractJobFields(r)

	b, err := c.getJobRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalJob(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeJobNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractJobFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteJob(ctx context.Context, r *Job) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Job resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Job...")
	deleteOp := deleteJobOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllJob deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllJob(ctx context.Context, project, location string, filter func(*Job) bool) error {
	listObj, err := c.ListJob(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllJob(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllJob(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyJob(ctx context.Context, rawDesired *Job, opts ...dcl.ApplyOption) (*Job, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Job
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyJobHelper(c, ctx, rawDesired, opts...)
		resultNewState = newState
		if err != nil {
			// If the error is 409, there is conflict in resource update.
			// Here we want to apply changes based on latest state.
			if dcl.IsConflictError(err) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		return nil, nil
	}, c.Config.RetryProvider)
	return resultNewState, err
}

func applyJobHelper(c *Client, ctx context.Context, rawDesired *Job, opts ...dcl.ApplyOption) (*Job, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyJob...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractJobFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.jobDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToJobDiffs(c.Config, fieldDiffs, opts)
	if err != nil {
		return nil, err
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		if dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
			return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Creation blocked by lifecycle params: %#v.", desired)}
		}
		create = true
	} else if dcl.HasLifecycleParam(lp, dcl.BlockAcquire) {
		return nil, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource already exists - apply blocked by lifecycle params: %#v.", initial),
		}
	} else {
		for _, d := range diffs {
			if d.RequiresRecreate {
				return nil, dcl.ApplyInfeasibleError{
					Message: fmt.Sprintf("infeasible update: (%v) would require recreation", d),
				}
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []jobApiOperation
	if create {
		ops = append(ops, &createJobOperation{})
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.InfoWithContextf(ctx, "Performing operation %T %+v", op, op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.InfoWithContextf(ctx, "Failed operation %T %+v: %v", op, op, err)
			return nil, err
		}
		c.Config.Logger.InfoWithContextf(ctx, "Finished operation %T %+v", op, op)
	}
	return applyJobDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyJobDiff(c *Client, ctx context.Context, desired *Job, rawDesired *Job, ops []jobApiOperation, opts ...dcl.ApplyOption) (*Job, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetJob(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createJobOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapJob(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeJobNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeJobNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeJobDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractJobFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractJobFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffJob(c, newDesired, newState)
	if err != nil {
		return newState, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.InfoWithContext(ctx, "No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.InfoWithContextf(ctx, "Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.InfoWithContext(ctx, "Done Apply.")
	return newState, nil
}
