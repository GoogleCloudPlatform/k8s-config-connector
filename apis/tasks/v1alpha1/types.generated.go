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


// +kcc:proto=google.cloud.tasks.v2.AppEngineHttpRequest
type AppEngineHttpRequest struct {
	// The HTTP method to use for the request. The default is POST.
	//
	//  The app's request handler for the task's target URL must be able to handle
	//  HTTP requests with this http_method, otherwise the task attempt fails with
	//  error code 405 (Method Not Allowed). See [Writing a push task request
	//  handler](https://cloud.google.com/appengine/docs/java/taskqueue/push/creating-handlers#writing_a_push_task_request_handler)
	//  and the App Engine documentation for your runtime on [How Requests are
	//  Handled](https://cloud.google.com/appengine/docs/standard/python3/how-requests-are-handled).
	// +kcc:proto:field=google.cloud.tasks.v2.AppEngineHttpRequest.http_method
	HTTPMethod *string `json:"httpMethod,omitempty"`

	// Task-level setting for App Engine routing.
	//
	//  * If [app_engine_routing_override is set on the
	//    queue][google.cloud.tasks.v2.Queue.app_engine_routing_override], this
	//    value is used for all tasks in the queue, no matter what the setting is
	//    for the [task-level
	//    app_engine_routing][google.cloud.tasks.v2.AppEngineHttpRequest.app_engine_routing].
	// +kcc:proto:field=google.cloud.tasks.v2.AppEngineHttpRequest.app_engine_routing
	AppEngineRouting *AppEngineRouting `json:"appEngineRouting,omitempty"`

	// The relative URI.
	//
	//  The relative URI must begin with "/" and must be a valid HTTP relative URI.
	//  It can contain a path and query string arguments.
	//  If the relative URI is empty, then the root path "/" will be used.
	//  No spaces are allowed, and the maximum length allowed is 2083 characters.
	// +kcc:proto:field=google.cloud.tasks.v2.AppEngineHttpRequest.relative_uri
	RelativeURI *string `json:"relativeURI,omitempty"`

	// HTTP request headers.
	//
	//  This map contains the header field names and values.
	//  Headers can be set when the
	//  [task is created][google.cloud.tasks.v2.CloudTasks.CreateTask].
	//  Repeated headers are not supported but a header value can contain commas.
	//
	//  Cloud Tasks sets some headers to default values:
	//
	//  * `User-Agent`: By default, this header is
	//    `"AppEngine-Google; (+http://code.google.com/appengine)"`.
	//    This header can be modified, but Cloud Tasks will append
	//    `"AppEngine-Google; (+http://code.google.com/appengine)"` to the
	//    modified `User-Agent`.
	//
	//  If the task has a [body][google.cloud.tasks.v2.AppEngineHttpRequest.body],
	//  Cloud Tasks sets the following headers:
	//
	//  * `Content-Type`: By default, the `Content-Type` header is set to
	//    `"application/octet-stream"`. The default can be overridden by explicitly
	//    setting `Content-Type` to a particular media type when the
	//    [task is created][google.cloud.tasks.v2.CloudTasks.CreateTask].
	//    For example, `Content-Type` can be set to `"application/json"`.
	//  * `Content-Length`: This is computed by Cloud Tasks. This value is
	//    output only.   It cannot be changed.
	//
	//  The headers below cannot be set or overridden:
	//
	//  * `Host`
	//  * `X-Google-*`
	//  * `X-AppEngine-*`
	//
	//  In addition, Cloud Tasks sets some headers when the task is dispatched,
	//  such as headers containing information about the task; see
	//  [request
	//  headers](https://cloud.google.com/tasks/docs/creating-appengine-handlers#reading_request_headers).
	//  These headers are set only when the task is dispatched, so they are not
	//  visible when the task is returned in a Cloud Tasks response.
	//
	//  Although there is no specific limit for the maximum number of headers or
	//  the size, there is a limit on the maximum size of the
	//  [Task][google.cloud.tasks.v2.Task]. For more information, see the
	//  [CreateTask][google.cloud.tasks.v2.CloudTasks.CreateTask] documentation.
	// +kcc:proto:field=google.cloud.tasks.v2.AppEngineHttpRequest.headers
	Headers map[string]string `json:"headers,omitempty"`

	// HTTP request body.
	//
	//  A request body is allowed only if the HTTP method is POST or PUT. It is
	//  an error to set a body on a task with an incompatible
	//  [HttpMethod][google.cloud.tasks.v2.HttpMethod].
	// +kcc:proto:field=google.cloud.tasks.v2.AppEngineHttpRequest.body
	Body []byte `json:"body,omitempty"`
}

// +kcc:proto=google.cloud.tasks.v2.AppEngineRouting
type AppEngineRouting struct {
	// App service.
	//
	//  By default, the task is sent to the service which is the default
	//  service when the task is attempted.
	//
	//  For some queues or tasks which were created using the App Engine
	//  Task Queue API, [host][google.cloud.tasks.v2.AppEngineRouting.host] is not
	//  parsable into [service][google.cloud.tasks.v2.AppEngineRouting.service],
	//  [version][google.cloud.tasks.v2.AppEngineRouting.version], and
	//  [instance][google.cloud.tasks.v2.AppEngineRouting.instance]. For example,
	//  some tasks which were created using the App Engine SDK use a custom domain
	//  name; custom domains are not parsed by Cloud Tasks. If
	//  [host][google.cloud.tasks.v2.AppEngineRouting.host] is not parsable, then
	//  [service][google.cloud.tasks.v2.AppEngineRouting.service],
	//  [version][google.cloud.tasks.v2.AppEngineRouting.version], and
	//  [instance][google.cloud.tasks.v2.AppEngineRouting.instance] are the empty
	//  string.
	// +kcc:proto:field=google.cloud.tasks.v2.AppEngineRouting.service
	Service *string `json:"service,omitempty"`

	// App version.
	//
	//  By default, the task is sent to the version which is the default
	//  version when the task is attempted.
	//
	//  For some queues or tasks which were created using the App Engine
	//  Task Queue API, [host][google.cloud.tasks.v2.AppEngineRouting.host] is not
	//  parsable into [service][google.cloud.tasks.v2.AppEngineRouting.service],
	//  [version][google.cloud.tasks.v2.AppEngineRouting.version], and
	//  [instance][google.cloud.tasks.v2.AppEngineRouting.instance]. For example,
	//  some tasks which were created using the App Engine SDK use a custom domain
	//  name; custom domains are not parsed by Cloud Tasks. If
	//  [host][google.cloud.tasks.v2.AppEngineRouting.host] is not parsable, then
	//  [service][google.cloud.tasks.v2.AppEngineRouting.service],
	//  [version][google.cloud.tasks.v2.AppEngineRouting.version], and
	//  [instance][google.cloud.tasks.v2.AppEngineRouting.instance] are the empty
	//  string.
	// +kcc:proto:field=google.cloud.tasks.v2.AppEngineRouting.version
	Version *string `json:"version,omitempty"`

	// App instance.
	//
	//  By default, the task is sent to an instance which is available when
	//  the task is attempted.
	//
	//  Requests can only be sent to a specific instance if
	//  [manual scaling is used in App Engine
	//  Standard](https://cloud.google.com/appengine/docs/python/an-overview-of-app-engine?hl=en_US#scaling_types_and_instance_classes).
	//  App Engine Flex does not support instances. For more information, see
	//  [App Engine Standard request
	//  routing](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed)
	//  and [App Engine Flex request
	//  routing](https://cloud.google.com/appengine/docs/flexible/python/how-requests-are-routed).
	// +kcc:proto:field=google.cloud.tasks.v2.AppEngineRouting.instance
	Instance *string `json:"instance,omitempty"`

	// Output only. The host that the task is sent to.
	//
	//  The host is constructed from the domain name of the app associated with
	//  the queue's project ID (for example <app-id>.appspot.com), and the
	//  [service][google.cloud.tasks.v2.AppEngineRouting.service],
	//  [version][google.cloud.tasks.v2.AppEngineRouting.version], and
	//  [instance][google.cloud.tasks.v2.AppEngineRouting.instance]. Tasks which
	//  were created using the App Engine SDK might have a custom domain name.
	//
	//  For more information, see
	//  [How Requests are
	//  Routed](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed).
	// +kcc:proto:field=google.cloud.tasks.v2.AppEngineRouting.host
	Host *string `json:"host,omitempty"`
}

// +kcc:proto=google.cloud.tasks.v2.Attempt
type Attempt struct {
	// Output only. The time that this attempt was scheduled.
	//
	//  `schedule_time` will be truncated to the nearest microsecond.
	// +kcc:proto:field=google.cloud.tasks.v2.Attempt.schedule_time
	ScheduleTime *string `json:"scheduleTime,omitempty"`

	// Output only. The time that this attempt was dispatched.
	//
	//  `dispatch_time` will be truncated to the nearest microsecond.
	// +kcc:proto:field=google.cloud.tasks.v2.Attempt.dispatch_time
	DispatchTime *string `json:"dispatchTime,omitempty"`

	// Output only. The time that this attempt response was received.
	//
	//  `response_time` will be truncated to the nearest microsecond.
	// +kcc:proto:field=google.cloud.tasks.v2.Attempt.response_time
	ResponseTime *string `json:"responseTime,omitempty"`

	// Output only. The response from the worker for this attempt.
	//
	//  If `response_time` is unset, then the task has not been attempted or is
	//  currently running and the `response_status` field is meaningless.
	// +kcc:proto:field=google.cloud.tasks.v2.Attempt.response_status
	ResponseStatus *Status `json:"responseStatus,omitempty"`
}

// +kcc:proto=google.cloud.tasks.v2.HttpRequest
type HttpRequest struct {
	// Required. The full url path that the request will be sent to.
	//
	//  This string must begin with either "http://" or "https://". Some examples
	//  are: `http://acme.com` and `https://acme.com/sales:8080`. Cloud Tasks will
	//  encode some characters for safety and compatibility. The maximum allowed
	//  URL length is 2083 characters after encoding.
	//
	//  The `Location` header response from a redirect response [`300` - `399`]
	//  may be followed. The redirect is not counted as a separate attempt.
	// +kcc:proto:field=google.cloud.tasks.v2.HttpRequest.url
	URL *string `json:"url,omitempty"`

	// The HTTP method to use for the request. The default is POST.
	// +kcc:proto:field=google.cloud.tasks.v2.HttpRequest.http_method
	HTTPMethod *string `json:"httpMethod,omitempty"`

	// HTTP request headers.
	//
	//  This map contains the header field names and values.
	//  Headers can be set when the
	//  [task is created][google.cloud.tasks.v2beta3.CloudTasks.CreateTask].
	//
	//  These headers represent a subset of the headers that will accompany the
	//  task's HTTP request. Some HTTP request headers will be ignored or replaced.
	//
	//  A partial list of headers that will be ignored or replaced is:
	//
	//  * Host: This will be computed by Cloud Tasks and derived from
	//    [HttpRequest.url][google.cloud.tasks.v2.HttpRequest.url].
	//  * Content-Length: This will be computed by Cloud Tasks.
	//  * User-Agent: This will be set to `"Google-Cloud-Tasks"`.
	//  * `X-Google-*`: Google use only.
	//  * `X-AppEngine-*`: Google use only.
	//
	//  `Content-Type` won't be set by Cloud Tasks. You can explicitly set
	//  `Content-Type` to a media type when the
	//   [task is created][google.cloud.tasks.v2beta3.CloudTasks.CreateTask].
	//   For example, `Content-Type` can be set to `"application/octet-stream"` or
	//   `"application/json"`.
	//
	//  Headers which can have multiple values (according to RFC2616) can be
	//  specified using comma-separated values.
	//
	//  The size of the headers must be less than 80KB.
	// +kcc:proto:field=google.cloud.tasks.v2.HttpRequest.headers
	Headers map[string]string `json:"headers,omitempty"`

	// HTTP request body.
	//
	//  A request body is allowed only if the
	//  [HTTP method][google.cloud.tasks.v2.HttpRequest.http_method] is POST, PUT,
	//  or PATCH. It is an error to set body on a task with an incompatible
	//  [HttpMethod][google.cloud.tasks.v2.HttpMethod].
	// +kcc:proto:field=google.cloud.tasks.v2.HttpRequest.body
	Body []byte `json:"body,omitempty"`

	// If specified, an
	//  [OAuth token](https://developers.google.com/identity/protocols/OAuth2)
	//  will be generated and attached as an `Authorization` header in the HTTP
	//  request.
	//
	//  This type of authorization should generally only be used when calling
	//  Google APIs hosted on *.googleapis.com.
	// +kcc:proto:field=google.cloud.tasks.v2.HttpRequest.oauth_token
	OauthToken *OAuthToken `json:"oauthToken,omitempty"`

	// If specified, an
	//  [OIDC](https://developers.google.com/identity/protocols/OpenIDConnect)
	//  token will be generated and attached as an `Authorization` header in the
	//  HTTP request.
	//
	//  This type of authorization can be used for many scenarios, including
	//  calling Cloud Run, or endpoints where you intend to validate the token
	//  yourself.
	// +kcc:proto:field=google.cloud.tasks.v2.HttpRequest.oidc_token
	OidcToken *OidcToken `json:"oidcToken,omitempty"`
}

// +kcc:proto=google.cloud.tasks.v2.OAuthToken
type OAuthToken struct {
	// [Service account email](https://cloud.google.com/iam/docs/service-accounts)
	//  to be used for generating OAuth token.
	//  The service account must be within the same project as the queue. The
	//  caller must have iam.serviceAccounts.actAs permission for the service
	//  account.
	// +kcc:proto:field=google.cloud.tasks.v2.OAuthToken.service_account_email
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty"`

	// OAuth scope to be used for generating OAuth access token.
	//  If not specified, "https://www.googleapis.com/auth/cloud-platform"
	//  will be used.
	// +kcc:proto:field=google.cloud.tasks.v2.OAuthToken.scope
	Scope *string `json:"scope,omitempty"`
}

// +kcc:proto=google.cloud.tasks.v2.OidcToken
type OidcToken struct {
	// [Service account email](https://cloud.google.com/iam/docs/service-accounts)
	//  to be used for generating OIDC token.
	//  The service account must be within the same project as the queue. The
	//  caller must have iam.serviceAccounts.actAs permission for the service
	//  account.
	// +kcc:proto:field=google.cloud.tasks.v2.OidcToken.service_account_email
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty"`

	// Audience to be used when generating OIDC token. If not specified, the URI
	//  specified in target will be used.
	// +kcc:proto:field=google.cloud.tasks.v2.OidcToken.audience
	Audience *string `json:"audience,omitempty"`
}

// +kcc:proto=google.cloud.tasks.v2.Task
type Task struct {
	// Optionally caller-specified in
	//  [CreateTask][google.cloud.tasks.v2.CloudTasks.CreateTask].
	//
	//  The task name.
	//
	//  The task name must have the following format:
	//  `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID/tasks/TASK_ID`
	//
	//  * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]),
	//     hyphens (-), colons (:), or periods (.).
	//     For more information, see
	//     [Identifying
	//     projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects)
	//  * `LOCATION_ID` is the canonical ID for the task's location.
	//     The list of available locations can be obtained by calling
	//     [ListLocations][google.cloud.location.Locations.ListLocations].
	//     For more information, see https://cloud.google.com/about/locations/.
	//  * `QUEUE_ID` can contain letters ([A-Za-z]), numbers ([0-9]), or
	//    hyphens (-). The maximum length is 100 characters.
	//  * `TASK_ID` can contain only letters ([A-Za-z]), numbers ([0-9]),
	//    hyphens (-), or underscores (_). The maximum length is 500 characters.
	// +kcc:proto:field=google.cloud.tasks.v2.Task.name
	Name *string `json:"name,omitempty"`

	// HTTP request that is sent to the App Engine app handler.
	//
	//  An App Engine task is a task that has
	//  [AppEngineHttpRequest][google.cloud.tasks.v2.AppEngineHttpRequest] set.
	// +kcc:proto:field=google.cloud.tasks.v2.Task.app_engine_http_request
	AppEngineHTTPRequest *AppEngineHttpRequest `json:"appEngineHTTPRequest,omitempty"`

	// HTTP request that is sent to the worker.
	//
	//  An HTTP task is a task that has
	//  [HttpRequest][google.cloud.tasks.v2.HttpRequest] set.
	// +kcc:proto:field=google.cloud.tasks.v2.Task.http_request
	HTTPRequest *HttpRequest `json:"httpRequest,omitempty"`

	// The time when the task is scheduled to be attempted or retried.
	//
	//  `schedule_time` will be truncated to the nearest microsecond.
	// +kcc:proto:field=google.cloud.tasks.v2.Task.schedule_time
	ScheduleTime *string `json:"scheduleTime,omitempty"`

	// Output only. The time that the task was created.
	//
	//  `create_time` will be truncated to the nearest second.
	// +kcc:proto:field=google.cloud.tasks.v2.Task.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// The deadline for requests sent to the worker. If the worker does not
	//  respond by this deadline then the request is cancelled and the attempt
	//  is marked as a `DEADLINE_EXCEEDED` failure. Cloud Tasks will retry the
	//  task according to the [RetryConfig][google.cloud.tasks.v2.RetryConfig].
	//
	//  Note that when the request is cancelled, Cloud Tasks will stop listening
	//  for the response, but whether the worker stops processing depends on the
	//  worker. For example, if the worker is stuck, it may not react to cancelled
	//  requests.
	//
	//  The default and maximum values depend on the type of request:
	//
	//  * For [HTTP tasks][google.cloud.tasks.v2.HttpRequest], the default is 10
	//  minutes. The deadline
	//    must be in the interval [15 seconds, 30 minutes].
	//
	//  * For [App Engine tasks][google.cloud.tasks.v2.AppEngineHttpRequest], 0
	//  indicates that the
	//    request has the default deadline. The default deadline depends on the
	//    [scaling
	//    type](https://cloud.google.com/appengine/docs/standard/go/how-instances-are-managed#instance_scaling)
	//    of the service: 10 minutes for standard apps with automatic scaling, 24
	//    hours for standard apps with manual and basic scaling, and 60 minutes for
	//    flex apps. If the request deadline is set, it must be in the interval [15
	//    seconds, 24 hours 15 seconds]. Regardless of the task's
	//    `dispatch_deadline`, the app handler will not run for longer than than
	//    the service's timeout. We recommend setting the `dispatch_deadline` to
	//    at most a few seconds more than the app handler's timeout. For more
	//    information see
	//    [Timeouts](https://cloud.google.com/tasks/docs/creating-appengine-handlers#timeouts).
	//
	//  `dispatch_deadline` will be truncated to the nearest millisecond. The
	//  deadline is an approximate deadline.
	// +kcc:proto:field=google.cloud.tasks.v2.Task.dispatch_deadline
	DispatchDeadline *string `json:"dispatchDeadline,omitempty"`

	// Output only. The number of attempts dispatched.
	//
	//  This count includes attempts which have been dispatched but haven't
	//  received a response.
	// +kcc:proto:field=google.cloud.tasks.v2.Task.dispatch_count
	DispatchCount *int32 `json:"dispatchCount,omitempty"`

	// Output only. The number of attempts which have received a response.
	// +kcc:proto:field=google.cloud.tasks.v2.Task.response_count
	ResponseCount *int32 `json:"responseCount,omitempty"`

	// Output only. The status of the task's first attempt.
	//
	//  Only [dispatch_time][google.cloud.tasks.v2.Attempt.dispatch_time] will be
	//  set. The other [Attempt][google.cloud.tasks.v2.Attempt] information is not
	//  retained by Cloud Tasks.
	// +kcc:proto:field=google.cloud.tasks.v2.Task.first_attempt
	FirstAttempt *Attempt `json:"firstAttempt,omitempty"`

	// Output only. The status of the task's last attempt.
	// +kcc:proto:field=google.cloud.tasks.v2.Task.last_attempt
	LastAttempt *Attempt `json:"lastAttempt,omitempty"`

	// Output only. The view specifies which subset of the
	//  [Task][google.cloud.tasks.v2.Task] has been returned.
	// +kcc:proto:field=google.cloud.tasks.v2.Task.view
	View *string `json:"view,omitempty"`
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
