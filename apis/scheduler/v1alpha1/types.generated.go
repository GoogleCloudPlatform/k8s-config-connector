// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1


// +kcc:proto=google.cloud.scheduler.v1beta1.AppEngineHttpTarget
type AppEngineHttpTarget struct {
	// The HTTP method to use for the request. PATCH and OPTIONS are not
	//  permitted.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.AppEngineHttpTarget.http_method
	HTTPMethod *string `json:"httpMethod,omitempty"`

	// App Engine Routing setting for the job.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.AppEngineHttpTarget.app_engine_routing
	AppEngineRouting *AppEngineRouting `json:"appEngineRouting,omitempty"`

	// The relative URI.
	//
	//  The relative URL must begin with "/" and must be a valid HTTP relative URL.
	//  It can contain a path, query string arguments, and `#` fragments.
	//  If the relative URL is empty, then the root path "/" will be used.
	//  No spaces are allowed, and the maximum length allowed is 2083 characters.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.AppEngineHttpTarget.relative_uri
	RelativeURI *string `json:"relativeURI,omitempty"`

	// HTTP request headers.
	//
	//  This map contains the header field names and values. Headers can be set
	//  when the job is created.
	//
	//  Cloud Scheduler sets some headers to default values:
	//
	//  * `User-Agent`: By default, this header is
	//    `"AppEngine-Google; (+http://code.google.com/appengine)"`.
	//    This header can be modified, but Cloud Scheduler will append
	//    `"AppEngine-Google; (+http://code.google.com/appengine)"` to the
	//    modified `User-Agent`.
	//  * `X-CloudScheduler`: This header will be set to true.
	//  * `X-CloudScheduler-JobName`: This header will contain the job name.
	//  * `X-CloudScheduler-ScheduleTime`: For Cloud Scheduler jobs specified in
	//  the unix-cron format, this header will contain the job schedule as an
	//  offset of UTC parsed according to RFC3339.
	//
	//  If the job has an
	//  [body][google.cloud.scheduler.v1beta1.AppEngineHttpTarget.body], Cloud
	//  Scheduler sets the following headers:
	//
	//  * `Content-Type`: By default, the `Content-Type` header is set to
	//    `"application/octet-stream"`. The default can be overridden by explictly
	//    setting `Content-Type` to a particular media type when the job is
	//    created.
	//    For example, `Content-Type` can be set to `"application/json"`.
	//  * `Content-Length`: This is computed by Cloud Scheduler. This value is
	//    output only. It cannot be changed.
	//
	//  The headers below are output only. They cannot be set or overridden:
	//
	//  * `X-Google-*`: For Google internal use only.
	//  * `X-AppEngine-*`: For Google internal use only.
	//
	//  In addition, some App Engine headers, which contain
	//  job-specific information, are also be sent to the job handler.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.AppEngineHttpTarget.headers
	Headers map[string]string `json:"headers,omitempty"`

	// Body.
	//
	//  HTTP request body. A request body is allowed only if the HTTP method is
	//  POST or PUT. It will result in invalid argument error to set a body on a
	//  job with an incompatible
	//  [HttpMethod][google.cloud.scheduler.v1beta1.HttpMethod].
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.AppEngineHttpTarget.body
	Body []byte `json:"body,omitempty"`
}

// +kcc:proto=google.cloud.scheduler.v1beta1.AppEngineRouting
type AppEngineRouting struct {
	// App service.
	//
	//  By default, the job is sent to the service which is the default
	//  service when the job is attempted.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.AppEngineRouting.service
	Service *string `json:"service,omitempty"`

	// App version.
	//
	//  By default, the job is sent to the version which is the default
	//  version when the job is attempted.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.AppEngineRouting.version
	Version *string `json:"version,omitempty"`

	// App instance.
	//
	//  By default, the job is sent to an instance which is available when
	//  the job is attempted.
	//
	//  Requests can only be sent to a specific instance if
	//  [manual scaling is used in App Engine
	//  Standard](https://cloud.google.com/appengine/docs/python/an-overview-of-app-engine?#scaling_types_and_instance_classes).
	//  App Engine Flex does not support instances. For more information, see
	//  [App Engine Standard request
	//  routing](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed)
	//  and [App Engine Flex request
	//  routing](https://cloud.google.com/appengine/docs/flexible/python/how-requests-are-routed).
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.AppEngineRouting.instance
	Instance *string `json:"instance,omitempty"`

	// Output only. The host that the job is sent to.
	//
	//  For more information about how App Engine requests are routed, see
	//  [here](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed).
	//
	//  The host is constructed as:
	//
	//
	//  * `host = [application_domain_name]`</br>
	//    `| [service] + '.' + [application_domain_name]`</br>
	//    `| [version] + '.' + [application_domain_name]`</br>
	//    `| [version_dot_service]+ '.' + [application_domain_name]`</br>
	//    `| [instance] + '.' + [application_domain_name]`</br>
	//    `| [instance_dot_service] + '.' + [application_domain_name]`</br>
	//    `| [instance_dot_version] + '.' + [application_domain_name]`</br>
	//    `| [instance_dot_version_dot_service] + '.' + [application_domain_name]`
	//
	//  * `application_domain_name` = The domain name of the app, for
	//    example <app-id>.appspot.com, which is associated with the
	//    job's project ID.
	//
	//  * `service =`
	//  [service][google.cloud.scheduler.v1beta1.AppEngineRouting.service]
	//
	//  * `version =`
	//  [version][google.cloud.scheduler.v1beta1.AppEngineRouting.version]
	//
	//  * `version_dot_service =`
	//    [version][google.cloud.scheduler.v1beta1.AppEngineRouting.version] `+ '.'
	//    +` [service][google.cloud.scheduler.v1beta1.AppEngineRouting.service]
	//
	//  * `instance =`
	//  [instance][google.cloud.scheduler.v1beta1.AppEngineRouting.instance]
	//
	//  * `instance_dot_service =`
	//    [instance][google.cloud.scheduler.v1beta1.AppEngineRouting.instance] `+
	//    '.' +` [service][google.cloud.scheduler.v1beta1.AppEngineRouting.service]
	//
	//  * `instance_dot_version =`
	//    [instance][google.cloud.scheduler.v1beta1.AppEngineRouting.instance] `+
	//    '.' +` [version][google.cloud.scheduler.v1beta1.AppEngineRouting.version]
	//
	//  * `instance_dot_version_dot_service =`
	//    [instance][google.cloud.scheduler.v1beta1.AppEngineRouting.instance] `+
	//    '.' +` [version][google.cloud.scheduler.v1beta1.AppEngineRouting.version]
	//    `+ '.' +`
	//    [service][google.cloud.scheduler.v1beta1.AppEngineRouting.service]
	//
	//
	//  If [service][google.cloud.scheduler.v1beta1.AppEngineRouting.service] is
	//  empty, then the job will be sent to the service which is the default
	//  service when the job is attempted.
	//
	//  If [version][google.cloud.scheduler.v1beta1.AppEngineRouting.version] is
	//  empty, then the job will be sent to the version which is the default
	//  version when the job is attempted.
	//
	//  If [instance][google.cloud.scheduler.v1beta1.AppEngineRouting.instance] is
	//  empty, then the job will be sent to an instance which is available when the
	//  job is attempted.
	//
	//  If [service][google.cloud.scheduler.v1beta1.AppEngineRouting.service],
	//  [version][google.cloud.scheduler.v1beta1.AppEngineRouting.version], or
	//  [instance][google.cloud.scheduler.v1beta1.AppEngineRouting.instance] is
	//  invalid, then the job will be sent to the default version of the default
	//  service when the job is attempted.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.AppEngineRouting.host
	Host *string `json:"host,omitempty"`
}

// +kcc:proto=google.cloud.scheduler.v1beta1.HttpTarget
type HttpTarget struct {
	// Required. The full URI path that the request will be sent to. This string
	//  must begin with either "http://" or "https://". Some examples of
	//  valid values for [uri][google.cloud.scheduler.v1beta1.HttpTarget.uri] are:
	//  `http://acme.com` and `https://acme.com/sales:8080`. Cloud Scheduler will
	//  encode some characters for safety and compatibility. The maximum allowed
	//  URL length is 2083 characters after encoding.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.HttpTarget.uri
	URI *string `json:"uri,omitempty"`

	// Which HTTP method to use for the request.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.HttpTarget.http_method
	HTTPMethod *string `json:"httpMethod,omitempty"`

	// The user can specify HTTP request headers to send with the job's
	//  HTTP request. This map contains the header field names and
	//  values. Repeated headers are not supported, but a header value can
	//  contain commas. These headers represent a subset of the headers
	//  that will accompany the job's HTTP request. Some HTTP request
	//  headers will be ignored or replaced. A partial list of headers that
	//  will be ignored or replaced is below:
	//  - Host: This will be computed by Cloud Scheduler and derived from
	//  [uri][google.cloud.scheduler.v1beta1.HttpTarget.uri].
	//  * `Content-Length`: This will be computed by Cloud Scheduler.
	//  * `User-Agent`: This will be set to `"Google-Cloud-Scheduler"`.
	//  * `X-Google-*`: Google internal use only.
	//  * `X-AppEngine-*`: Google internal use only.
	//  * `X-CloudScheduler`: This header will be set to true.
	//  * `X-CloudScheduler-JobName`: This header will contain the job name.
	//  * `X-CloudScheduler-ScheduleTime`: For Cloud Scheduler jobs specified in
	//  the unix-cron format, this header will contain the job schedule as an
	//  offset of UTC parsed according to RFC3339.
	//
	//  The total size of headers must be less than 80KB.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.HttpTarget.headers
	Headers map[string]string `json:"headers,omitempty"`

	// HTTP request body. A request body is allowed only if the HTTP
	//  method is POST, PUT, or PATCH. It is an error to set body on a job with an
	//  incompatible [HttpMethod][google.cloud.scheduler.v1beta1.HttpMethod].
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.HttpTarget.body
	Body []byte `json:"body,omitempty"`

	// If specified, an
	//  [OAuth token](https://developers.google.com/identity/protocols/OAuth2)
	//  will be generated and attached as an `Authorization` header in the HTTP
	//  request.
	//
	//  This type of authorization should generally only be used when calling
	//  Google APIs hosted on *.googleapis.com.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.HttpTarget.oauth_token
	OauthToken *OAuthToken `json:"oauthToken,omitempty"`

	// If specified, an
	//  [OIDC](https://developers.google.com/identity/protocols/OpenIDConnect)
	//  token will be generated and attached as an `Authorization` header in the
	//  HTTP request.
	//
	//  This type of authorization can be used for many scenarios, including
	//  calling Cloud Run, or endpoints where you intend to validate the token
	//  yourself.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.HttpTarget.oidc_token
	OidcToken *OidcToken `json:"oidcToken,omitempty"`
}

// +kcc:proto=google.cloud.scheduler.v1beta1.Job
type Job struct {
	// Optionally caller-specified in
	//  [CreateJob][google.cloud.scheduler.v1beta1.CloudScheduler.CreateJob], after
	//  which it becomes output only.
	//
	//  The job name. For example:
	//  `projects/PROJECT_ID/locations/LOCATION_ID/jobs/JOB_ID`.
	//
	//  * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]),
	//     hyphens (-), colons (:), or periods (.).
	//     For more information, see
	//     [Identifying
	//     projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects)
	//  * `LOCATION_ID` is the canonical ID for the job's location.
	//     The list of available locations can be obtained by calling
	//     [ListLocations][google.cloud.location.Locations.ListLocations].
	//     For more information, see https://cloud.google.com/about/locations/.
	//  * `JOB_ID` can contain only letters ([A-Za-z]), numbers ([0-9]),
	//     hyphens (-), or underscores (_). The maximum length is 500 characters.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.Job.name
	Name *string `json:"name,omitempty"`

	// Optionally caller-specified in
	//  [CreateJob][google.cloud.scheduler.v1beta1.CloudScheduler.CreateJob] or
	//  [UpdateJob][google.cloud.scheduler.v1beta1.CloudScheduler.UpdateJob].
	//
	//  A human-readable description for the job. This string must not contain
	//  more than 500 characters.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.Job.description
	Description *string `json:"description,omitempty"`

	// Pub/Sub target.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.Job.pubsub_target
	PubsubTarget *PubsubTarget `json:"pubsubTarget,omitempty"`

	// App Engine HTTP target.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.Job.app_engine_http_target
	AppEngineHTTPTarget *AppEngineHttpTarget `json:"appEngineHTTPTarget,omitempty"`

	// HTTP target.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.Job.http_target
	HTTPTarget *HttpTarget `json:"httpTarget,omitempty"`

	// Required, except when used with
	//  [UpdateJob][google.cloud.scheduler.v1beta1.CloudScheduler.UpdateJob].
	//
	//  Describes the schedule on which the job will be executed.
	//
	//  The schedule can be either of the following types:
	//
	//  * [Crontab](https://en.wikipedia.org/wiki/Cron#Overview)
	//  * English-like
	//  [schedule](https://cloud.google.com/scheduler/docs/configuring/cron-job-schedules)
	//
	//  As a general rule, execution `n + 1` of a job will not begin
	//  until execution `n` has finished. Cloud Scheduler will never
	//  allow two simultaneously outstanding executions. For example,
	//  this implies that if the `n+1`th execution is scheduled to run at
	//  16:00 but the `n`th execution takes until 16:15, the `n+1`th
	//  execution will not start until `16:15`.
	//  A scheduled start time will be delayed if the previous
	//  execution has not ended when its scheduled time occurs.
	//
	//  If [retry_count][google.cloud.scheduler.v1beta1.RetryConfig.retry_count] >
	//  0 and a job attempt fails, the job will be tried a total of
	//  [retry_count][google.cloud.scheduler.v1beta1.RetryConfig.retry_count]
	//  times, with exponential backoff, until the next scheduled start
	//  time.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.Job.schedule
	Schedule *string `json:"schedule,omitempty"`

	// Specifies the time zone to be used in interpreting
	//  [schedule][google.cloud.scheduler.v1beta1.Job.schedule]. The value of this
	//  field must be a time zone name from the [tz
	//  database](http://en.wikipedia.org/wiki/Tz_database).
	//
	//  Note that some time zones include a provision for
	//  daylight savings time. The rules for daylight saving time are
	//  determined by the chosen tz. For UTC use the string "utc". If a
	//  time zone is not specified, the default will be in UTC (also known
	//  as GMT).
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.Job.time_zone
	TimeZone *string `json:"timeZone,omitempty"`

	// Output only. The creation time of the job.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.Job.user_update_time
	UserUpdateTime *string `json:"userUpdateTime,omitempty"`

	// Output only. State of the job.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.Job.state
	State *string `json:"state,omitempty"`

	// Output only. The response from the target for the last attempted execution.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.Job.status
	Status *Status `json:"status,omitempty"`

	// Output only. The next time the job is scheduled. Note that this may be a
	//  retry of a previously failed attempt or the next execution time
	//  according to the schedule.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.Job.schedule_time
	ScheduleTime *string `json:"scheduleTime,omitempty"`

	// Output only. The time the last job attempt started.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.Job.last_attempt_time
	LastAttemptTime *string `json:"lastAttemptTime,omitempty"`

	// Settings that determine the retry behavior.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.Job.retry_config
	RetryConfig *RetryConfig `json:"retryConfig,omitempty"`

	// The deadline for job attempts. If the request handler does not respond by
	//  this deadline then the request is cancelled and the attempt is marked as a
	//  `DEADLINE_EXCEEDED` failure. The failed attempt can be viewed in
	//  execution logs. Cloud Scheduler will retry the job according
	//  to the [RetryConfig][google.cloud.scheduler.v1beta1.RetryConfig].
	//
	//  The default and the allowed values depend on the type of target:
	//
	//  * For [HTTP targets][google.cloud.scheduler.v1beta1.Job.http_target], the
	//  default is 3 minutes. The deadline must be in the interval [15 seconds, 30
	//  minutes].
	//
	//  * For [App Engine HTTP
	//  targets][google.cloud.scheduler.v1beta1.Job.app_engine_http_target], 0
	//  indicates that the request has the default deadline. The default deadline
	//  depends on the scaling type of the service: 10 minutes for standard apps
	//  with automatic scaling, 24 hours for standard apps with manual and basic
	//  scaling, and 60 minutes for flex apps. If the request deadline is set, it
	//  must be in the interval [15 seconds, 24 hours 15 seconds].
	//
	//  * For [Pub/Sub targets][google.cloud.scheduler.v1beta1.Job.pubsub_target],
	//  this field is ignored.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.Job.attempt_deadline
	AttemptDeadline *string `json:"attemptDeadline,omitempty"`

	// Immutable. This field is used to manage the legacy App Engine Cron jobs
	//  using the Cloud Scheduler API. If the field is set to true, the job will be
	//  considered a legacy job. Note that App Engine Cron jobs have fewer
	//  features than Cloud Scheduler jobs, e.g., are only limited to App Engine
	//  targets.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.Job.legacy_app_engine_cron
	LegacyAppEngineCron *bool `json:"legacyAppEngineCron,omitempty"`
}

// +kcc:proto=google.cloud.scheduler.v1beta1.OAuthToken
type OAuthToken struct {
	// [Service account email](https://cloud.google.com/iam/docs/service-accounts)
	//  to be used for generating OAuth token.
	//  The service account must be within the same project as the job. The caller
	//  must have iam.serviceAccounts.actAs permission for the service account.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.OAuthToken.service_account_email
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty"`

	// OAuth scope to be used for generating OAuth access token.
	//  If not specified, "https://www.googleapis.com/auth/cloud-platform"
	//  will be used.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.OAuthToken.scope
	Scope *string `json:"scope,omitempty"`
}

// +kcc:proto=google.cloud.scheduler.v1beta1.OidcToken
type OidcToken struct {
	// [Service account email](https://cloud.google.com/iam/docs/service-accounts)
	//  to be used for generating OIDC token.
	//  The service account must be within the same project as the job. The caller
	//  must have iam.serviceAccounts.actAs permission for the service account.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.OidcToken.service_account_email
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty"`

	// Audience to be used when generating OIDC token. If not specified, the URI
	//  specified in target will be used.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.OidcToken.audience
	Audience *string `json:"audience,omitempty"`
}

// +kcc:proto=google.cloud.scheduler.v1beta1.PubsubTarget
type PubsubTarget struct {
	// Required. The name of the Cloud Pub/Sub topic to which messages will
	//  be published when a job is delivered. The topic name must be in the
	//  same format as required by Pub/Sub's
	//  [PublishRequest.name](https://cloud.google.com/pubsub/docs/reference/rpc/google.pubsub.v1#publishrequest),
	//  for example `projects/PROJECT_ID/topics/TOPIC_ID`.
	//
	//  The topic must be in the same project as the Cloud Scheduler job.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.PubsubTarget.topic_name
	TopicName *string `json:"topicName,omitempty"`

	// The message payload for PubsubMessage.
	//
	//  Pubsub message must contain either non-empty data, or at least one
	//  attribute.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.PubsubTarget.data
	Data []byte `json:"data,omitempty"`

	// Attributes for PubsubMessage.
	//
	//  Pubsub message must contain either non-empty data, or at least one
	//  attribute.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.PubsubTarget.attributes
	Attributes map[string]string `json:"attributes,omitempty"`
}

// +kcc:proto=google.cloud.scheduler.v1beta1.RetryConfig
type RetryConfig struct {
	// The number of attempts that the system will make to run a job using the
	//  exponential backoff procedure described by
	//  [max_doublings][google.cloud.scheduler.v1beta1.RetryConfig.max_doublings].
	//
	//  The default value of retry_count is zero.
	//
	//  If retry_count is zero, a job attempt will *not* be retried if
	//  it fails. Instead the Cloud Scheduler system will wait for the
	//  next scheduled execution time.
	//
	//  If retry_count is set to a non-zero number then Cloud Scheduler
	//  will retry failed attempts, using exponential backoff,
	//  retry_count times, or until the next scheduled execution time,
	//  whichever comes first.
	//
	//  Values greater than 5 and negative values are not allowed.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.RetryConfig.retry_count
	RetryCount *int32 `json:"retryCount,omitempty"`

	// The time limit for retrying a failed job, measured from time when an
	//  execution was first attempted. If specified with
	//  [retry_count][google.cloud.scheduler.v1beta1.RetryConfig.retry_count], the
	//  job will be retried until both limits are reached.
	//
	//  The default value for max_retry_duration is zero, which means retry
	//  duration is unlimited.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.RetryConfig.max_retry_duration
	MaxRetryDuration *string `json:"maxRetryDuration,omitempty"`

	// The minimum amount of time to wait before retrying a job after
	//  it fails.
	//
	//  The default value of this field is 5 seconds.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.RetryConfig.min_backoff_duration
	MinBackoffDuration *string `json:"minBackoffDuration,omitempty"`

	// The maximum amount of time to wait before retrying a job after
	//  it fails.
	//
	//  The default value of this field is 1 hour.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.RetryConfig.max_backoff_duration
	MaxBackoffDuration *string `json:"maxBackoffDuration,omitempty"`

	// The time between retries will double `max_doublings` times.
	//
	//  A job's retry interval starts at
	//  [min_backoff_duration][google.cloud.scheduler.v1beta1.RetryConfig.min_backoff_duration],
	//  then doubles `max_doublings` times, then increases linearly, and finally
	//  retries at intervals of
	//  [max_backoff_duration][google.cloud.scheduler.v1beta1.RetryConfig.max_backoff_duration]
	//  up to [retry_count][google.cloud.scheduler.v1beta1.RetryConfig.retry_count]
	//  times.
	//
	//  For example, if
	//  [min_backoff_duration][google.cloud.scheduler.v1beta1.RetryConfig.min_backoff_duration]
	//  is 10s,
	//  [max_backoff_duration][google.cloud.scheduler.v1beta1.RetryConfig.max_backoff_duration]
	//  is 300s, and `max_doublings` is 3, then the a job will first be retried in
	//  10s. The retry interval will double three times, and then increase linearly
	//  by 2^3 * 10s.  Finally, the job will retry at intervals of
	//  [max_backoff_duration][google.cloud.scheduler.v1beta1.RetryConfig.max_backoff_duration]
	//  until the job has been attempted
	//  [retry_count][google.cloud.scheduler.v1beta1.RetryConfig.retry_count]
	//  times. Thus, the requests will retry at 10s, 20s, 40s, 80s, 160s, 240s,
	//  300s, 300s, ....
	//
	//  The default value of this field is 5.
	// +kcc:proto:field=google.cloud.scheduler.v1beta1.RetryConfig.max_doublings
	MaxDoublings *int32 `json:"maxDoublings,omitempty"`
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}
