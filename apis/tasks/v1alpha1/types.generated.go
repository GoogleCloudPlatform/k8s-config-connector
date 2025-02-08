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


// +kcc:proto=google.cloud.tasks.v2beta2.AppEngineHttpTarget
type AppEngineHttpTarget struct {
	// Overrides for the
	//  [task-level
	//  app_engine_routing][google.cloud.tasks.v2beta2.AppEngineHttpRequest.app_engine_routing].
	//
	//  If set, `app_engine_routing_override` is used for all tasks in
	//  the queue, no matter what the setting is for the
	//  [task-level
	//  app_engine_routing][google.cloud.tasks.v2beta2.AppEngineHttpRequest.app_engine_routing].
	// +kcc:proto:field=google.cloud.tasks.v2beta2.AppEngineHttpTarget.app_engine_routing_override
	AppEngineRoutingOverride *AppEngineRouting `json:"appEngineRoutingOverride,omitempty"`
}

// +kcc:proto=google.cloud.tasks.v2beta2.AppEngineRouting
type AppEngineRouting struct {
	// App service.
	//
	//  By default, the task is sent to the service which is the default
	//  service when the task is attempted.
	//
	//  For some queues or tasks which were created using the App Engine
	//  Task Queue API, [host][google.cloud.tasks.v2beta2.AppEngineRouting.host] is
	//  not parsable into
	//  [service][google.cloud.tasks.v2beta2.AppEngineRouting.service],
	//  [version][google.cloud.tasks.v2beta2.AppEngineRouting.version], and
	//  [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance]. For
	//  example, some tasks which were created using the App Engine SDK use a
	//  custom domain name; custom domains are not parsed by Cloud Tasks. If
	//  [host][google.cloud.tasks.v2beta2.AppEngineRouting.host] is not parsable,
	//  then [service][google.cloud.tasks.v2beta2.AppEngineRouting.service],
	//  [version][google.cloud.tasks.v2beta2.AppEngineRouting.version], and
	//  [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance] are the
	//  empty string.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.AppEngineRouting.service
	Service *string `json:"service,omitempty"`

	// App version.
	//
	//  By default, the task is sent to the version which is the default
	//  version when the task is attempted.
	//
	//  For some queues or tasks which were created using the App Engine
	//  Task Queue API, [host][google.cloud.tasks.v2beta2.AppEngineRouting.host] is
	//  not parsable into
	//  [service][google.cloud.tasks.v2beta2.AppEngineRouting.service],
	//  [version][google.cloud.tasks.v2beta2.AppEngineRouting.version], and
	//  [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance]. For
	//  example, some tasks which were created using the App Engine SDK use a
	//  custom domain name; custom domains are not parsed by Cloud Tasks. If
	//  [host][google.cloud.tasks.v2beta2.AppEngineRouting.host] is not parsable,
	//  then [service][google.cloud.tasks.v2beta2.AppEngineRouting.service],
	//  [version][google.cloud.tasks.v2beta2.AppEngineRouting.version], and
	//  [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance] are the
	//  empty string.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.AppEngineRouting.version
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
	// +kcc:proto:field=google.cloud.tasks.v2beta2.AppEngineRouting.instance
	Instance *string `json:"instance,omitempty"`

	// Output only. The host that the task is sent to.
	//
	//  For more information, see
	//  [How Requests are
	//  Routed](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed).
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
	//    queue's project ID. Some tasks which were created using the App Engine
	//    SDK use a custom domain name.
	//
	//  * `service =`
	//  [service][google.cloud.tasks.v2beta2.AppEngineRouting.service]
	//
	//  * `version =`
	//  [version][google.cloud.tasks.v2beta2.AppEngineRouting.version]
	//
	//  * `version_dot_service =`
	//    [version][google.cloud.tasks.v2beta2.AppEngineRouting.version] `+ '.' +`
	//    [service][google.cloud.tasks.v2beta2.AppEngineRouting.service]
	//
	//  * `instance =`
	//  [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance]
	//
	//  * `instance_dot_service =`
	//    [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance] `+ '.'
	//    +` [service][google.cloud.tasks.v2beta2.AppEngineRouting.service]
	//
	//  * `instance_dot_version =`
	//    [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance] `+ '.'
	//    +` [version][google.cloud.tasks.v2beta2.AppEngineRouting.version]
	//
	//  * `instance_dot_version_dot_service =`
	//    [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance] `+ '.'
	//    +` [version][google.cloud.tasks.v2beta2.AppEngineRouting.version] `+ '.'
	//    +` [service][google.cloud.tasks.v2beta2.AppEngineRouting.service]
	//
	//  If [service][google.cloud.tasks.v2beta2.AppEngineRouting.service] is empty,
	//  then the task will be sent to the service which is the default service when
	//  the task is attempted.
	//
	//  If [version][google.cloud.tasks.v2beta2.AppEngineRouting.version] is empty,
	//  then the task will be sent to the version which is the default version when
	//  the task is attempted.
	//
	//  If [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance] is
	//  empty, then the task will be sent to an instance which is available when
	//  the task is attempted.
	//
	//  If [service][google.cloud.tasks.v2beta2.AppEngineRouting.service],
	//  [version][google.cloud.tasks.v2beta2.AppEngineRouting.version], or
	//  [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance] is
	//  invalid, then the task will be sent to the default version of the default
	//  service when the task is attempted.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.AppEngineRouting.host
	Host *string `json:"host,omitempty"`
}

// +kcc:proto=google.cloud.tasks.v2beta2.HttpTarget
type HttpTarget struct {
	// Uri override.
	//
	//  When specified, overrides the execution Uri for all the tasks in the queue.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.HttpTarget.uri_override
	URIOverride *UriOverride `json:"uriOverride,omitempty"`

	// The HTTP method to use for the request.
	//
	//  When specified, it overrides
	//  [HttpRequest][google.cloud.tasks.v2beta2.HttpTarget.http_method] for the
	//  task. Note that if the value is set to [HttpMethod][GET] the
	//  [HttpRequest][body] of the task will be ignored at execution time.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.HttpTarget.http_method
	HTTPMethod *string `json:"httpMethod,omitempty"`

	// HTTP target headers.
	//
	//  This map contains the header field names and values.
	//  Headers will be set when running the
	//  [task is created][google.cloud.tasks.v2beta2.CloudTasks.CreateTask] and/or
	//  [task is created][google.cloud.tasks.v2beta2.CloudTasks.BufferTask].
	//
	//  These headers represent a subset of the headers that will accompany the
	//  task's HTTP request. Some HTTP request headers will be ignored or replaced.
	//
	//  A partial list of headers that will be ignored or replaced is:
	//  * Any header that is prefixed with "X-CloudTasks-" will be treated
	//  as service header. Service headers define properties of the task and are
	//  predefined in CloudTask.
	//  * Host: This will be computed by Cloud Tasks and derived from
	//    [HttpRequest.url][google.cloud.tasks.v2beta2.HttpRequest.url].
	//  * Content-Length: This will be computed by Cloud Tasks.
	//  * User-Agent: This will be set to `"Google-CloudTasks"`.
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
	//  Queue-level headers to override headers of all the tasks in the queue.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.HttpTarget.header_overrides
	HeaderOverrides []HttpTarget_HeaderOverride `json:"headerOverrides,omitempty"`

	// If specified, an
	//  [OAuth token](https://developers.google.com/identity/protocols/OAuth2)
	//  will be generated and attached as an `Authorization` header in the HTTP
	//  request.
	//
	//  This type of authorization should generally only be used when calling
	//  Google APIs hosted on *.googleapis.com.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.HttpTarget.oauth_token
	OauthToken *OAuthToken `json:"oauthToken,omitempty"`

	// If specified, an
	//  [OIDC](https://developers.google.com/identity/protocols/OpenIDConnect)
	//  token will be generated and attached as an `Authorization` header in the
	//  HTTP request.
	//
	//  This type of authorization can be used for many scenarios, including
	//  calling Cloud Run, or endpoints where you intend to validate the token
	//  yourself.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.HttpTarget.oidc_token
	OidcToken *OidcToken `json:"oidcToken,omitempty"`
}

// +kcc:proto=google.cloud.tasks.v2beta2.HttpTarget.Header
type HttpTarget_Header struct {
	// The key of the header.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.HttpTarget.Header.key
	Key *string `json:"key,omitempty"`

	// The value of the header.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.HttpTarget.Header.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.tasks.v2beta2.HttpTarget.HeaderOverride
type HttpTarget_HeaderOverride struct {
	// header embodying a key and a value.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.HttpTarget.HeaderOverride.header
	Header *HttpTarget_Header `json:"header,omitempty"`
}

// +kcc:proto=google.cloud.tasks.v2beta2.OAuthToken
type OAuthToken struct {
	// [Service account email](https://cloud.google.com/iam/docs/service-accounts)
	//  to be used for generating OAuth token.
	//  The service account must be within the same project as the queue. The
	//  caller must have iam.serviceAccounts.actAs permission for the service
	//  account.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.OAuthToken.service_account_email
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty"`

	// OAuth scope to be used for generating OAuth access token.
	//  If not specified, "https://www.googleapis.com/auth/cloud-platform"
	//  will be used.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.OAuthToken.scope
	Scope *string `json:"scope,omitempty"`
}

// +kcc:proto=google.cloud.tasks.v2beta2.OidcToken
type OidcToken struct {
	// [Service account email](https://cloud.google.com/iam/docs/service-accounts)
	//  to be used for generating OIDC token.
	//  The service account must be within the same project as the queue. The
	//  caller must have iam.serviceAccounts.actAs permission for the service
	//  account.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.OidcToken.service_account_email
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty"`

	// Audience to be used when generating OIDC token. If not specified, the URI
	//  specified in target will be used.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.OidcToken.audience
	Audience *string `json:"audience,omitempty"`
}

// +kcc:proto=google.cloud.tasks.v2beta2.PathOverride
type PathOverride struct {
	// The URI path (e.g., /users/1234). Default is an empty string.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.PathOverride.path
	Path *string `json:"path,omitempty"`
}

// +kcc:proto=google.cloud.tasks.v2beta2.PullTarget
type PullTarget struct {
}

// +kcc:proto=google.cloud.tasks.v2beta2.QueryOverride
type QueryOverride struct {
	// The query parameters (e.g., qparam1=123&qparam2=456). Default is an empty
	//  string.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.QueryOverride.query_params
	QueryParams *string `json:"queryParams,omitempty"`
}

// +kcc:proto=google.cloud.tasks.v2beta2.Queue
type Queue struct {
	// Caller-specified and required in
	//  [CreateQueue][google.cloud.tasks.v2beta2.CloudTasks.CreateQueue], after
	//  which it becomes output only.
	//
	//  The queue name.
	//
	//  The queue name must have the following format:
	//  `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID`
	//
	//  * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]),
	//     hyphens (-), colons (:), or periods (.).
	//     For more information, see
	//     [Identifying
	//     projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects)
	//  * `LOCATION_ID` is the canonical ID for the queue's location.
	//     The list of available locations can be obtained by calling
	//     [ListLocations][google.cloud.location.Locations.ListLocations].
	//     For more information, see https://cloud.google.com/about/locations/.
	//  * `QUEUE_ID` can contain letters ([A-Za-z]), numbers ([0-9]), or
	//    hyphens (-). The maximum length is 100 characters.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.Queue.name
	Name *string `json:"name,omitempty"`

	// App Engine HTTP target.
	//
	//  An App Engine queue is a queue that has an
	//  [AppEngineHttpTarget][google.cloud.tasks.v2beta2.AppEngineHttpTarget].
	// +kcc:proto:field=google.cloud.tasks.v2beta2.Queue.app_engine_http_target
	AppEngineHTTPTarget *AppEngineHttpTarget `json:"appEngineHTTPTarget,omitempty"`

	// Pull target.
	//
	//  A pull queue is a queue that has a
	//  [PullTarget][google.cloud.tasks.v2beta2.PullTarget].
	// +kcc:proto:field=google.cloud.tasks.v2beta2.Queue.pull_target
	PullTarget *PullTarget `json:"pullTarget,omitempty"`

	// An http_target is used to override the target values for HTTP tasks.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.Queue.http_target
	HTTPTarget *HttpTarget `json:"httpTarget,omitempty"`

	// Rate limits for task dispatches.
	//
	//  [rate_limits][google.cloud.tasks.v2beta2.Queue.rate_limits] and
	//  [retry_config][google.cloud.tasks.v2beta2.Queue.retry_config] are related
	//  because they both control task attempts however they control how tasks are
	//  attempted in different ways:
	//
	//  * [rate_limits][google.cloud.tasks.v2beta2.Queue.rate_limits] controls the
	//  total rate of
	//    dispatches from a queue (i.e. all traffic dispatched from the
	//    queue, regardless of whether the dispatch is from a first
	//    attempt or a retry).
	//  * [retry_config][google.cloud.tasks.v2beta2.Queue.retry_config] controls
	//  what happens to
	//    particular a task after its first attempt fails. That is,
	//    [retry_config][google.cloud.tasks.v2beta2.Queue.retry_config] controls
	//    task retries (the second attempt, third attempt, etc).
	// +kcc:proto:field=google.cloud.tasks.v2beta2.Queue.rate_limits
	RateLimits *RateLimits `json:"rateLimits,omitempty"`

	// Settings that determine the retry behavior.
	//
	//  * For tasks created using Cloud Tasks: the queue-level retry settings
	//    apply to all tasks in the queue that were created using Cloud Tasks.
	//    Retry settings cannot be set on individual tasks.
	//  * For tasks created using the App Engine SDK: the queue-level retry
	//    settings apply to all tasks in the queue which do not have retry settings
	//    explicitly set on the task and were created by the App Engine SDK. See
	//    [App Engine
	//    documentation](https://cloud.google.com/appengine/docs/standard/python/taskqueue/push/retrying-tasks).
	// +kcc:proto:field=google.cloud.tasks.v2beta2.Queue.retry_config
	RetryConfig *RetryConfig `json:"retryConfig,omitempty"`

	// Output only. The state of the queue.
	//
	//  `state` can only be changed by called
	//  [PauseQueue][google.cloud.tasks.v2beta2.CloudTasks.PauseQueue],
	//  [ResumeQueue][google.cloud.tasks.v2beta2.CloudTasks.ResumeQueue], or
	//  uploading
	//  [queue.yaml/xml](https://cloud.google.com/appengine/docs/python/config/queueref).
	//  [UpdateQueue][google.cloud.tasks.v2beta2.CloudTasks.UpdateQueue] cannot be
	//  used to change `state`.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.Queue.state
	State *string `json:"state,omitempty"`

	// Output only. The last time this queue was purged.
	//
	//  All tasks that were [created][google.cloud.tasks.v2beta2.Task.create_time]
	//  before this time were purged.
	//
	//  A queue can be purged using
	//  [PurgeQueue][google.cloud.tasks.v2beta2.CloudTasks.PurgeQueue], the [App
	//  Engine Task Queue SDK, or the Cloud
	//  Console](https://cloud.google.com/appengine/docs/standard/python/taskqueue/push/deleting-tasks-and-queues#purging_all_tasks_from_a_queue).
	//
	//  Purge time will be truncated to the nearest microsecond. Purge
	//  time will be unset if the queue has never been purged.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.Queue.purge_time
	PurgeTime *string `json:"purgeTime,omitempty"`

	// The maximum amount of time that a task will be retained in
	//  this queue.
	//
	//  Queues created by Cloud Tasks have a default `task_ttl` of 31 days.
	//  After a task has lived for `task_ttl`, the task will be deleted
	//  regardless of whether it was dispatched or not.
	//
	//  The `task_ttl` for queues created via queue.yaml/xml is equal to the
	//  maximum duration because there is a
	//  [storage quota](https://cloud.google.com/appengine/quotas#Task_Queue) for
	//  these queues. To view the maximum valid duration, see the documentation for
	//  [Duration][google.protobuf.Duration].
	// +kcc:proto:field=google.cloud.tasks.v2beta2.Queue.task_ttl
	TaskTtl *string `json:"taskTtl,omitempty"`

	// The task tombstone time to live (TTL).
	//
	//  After a task is deleted or completed, the task's tombstone is
	//  retained for the length of time specified by `tombstone_ttl`.
	//  The tombstone is used by task de-duplication; another task with the same
	//  name can't be created until the tombstone has expired. For more information
	//  about task de-duplication, see the documentation for
	//  [CreateTaskRequest][google.cloud.tasks.v2beta2.CreateTaskRequest.task].
	//
	//  Queues created by Cloud Tasks have a default `tombstone_ttl` of 1 hour.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.Queue.tombstone_ttl
	TombstoneTtl *string `json:"tombstoneTtl,omitempty"`
}

// +kcc:proto=google.cloud.tasks.v2beta2.QueueStats
type QueueStats struct {
}

// +kcc:proto=google.cloud.tasks.v2beta2.RateLimits
type RateLimits struct {
	// The maximum rate at which tasks are dispatched from this queue.
	//
	//  If unspecified when the queue is created, Cloud Tasks will pick the
	//  default.
	//
	//  * For [App Engine queues][google.cloud.tasks.v2beta2.AppEngineHttpTarget],
	//  the maximum allowed value
	//    is 500.
	//  * This field is output only   for [pull
	//  queues][google.cloud.tasks.v2beta2.PullTarget]. In addition to the
	//    `max_tasks_dispatched_per_second` limit, a maximum of 10 QPS of
	//    [LeaseTasks][google.cloud.tasks.v2beta2.CloudTasks.LeaseTasks] requests
	//    are allowed per pull queue.
	//
	//
	//  This field has the same meaning as
	//  [rate in
	//  queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#rate).
	// +kcc:proto:field=google.cloud.tasks.v2beta2.RateLimits.max_tasks_dispatched_per_second
	MaxTasksDispatchedPerSecond *float64 `json:"maxTasksDispatchedPerSecond,omitempty"`

	// The max burst size.
	//
	//  Max burst size limits how fast tasks in queue are processed when
	//  many tasks are in the queue and the rate is high. This field
	//  allows the queue to have a high rate so processing starts shortly
	//  after a task is enqueued, but still limits resource usage when
	//  many tasks are enqueued in a short period of time.
	//
	//  The [token bucket](https://wikipedia.org/wiki/Token_Bucket)
	//  algorithm is used to control the rate of task dispatches. Each
	//  queue has a token bucket that holds tokens, up to the maximum
	//  specified by `max_burst_size`. Each time a task is dispatched, a
	//  token is removed from the bucket. Tasks will be dispatched until
	//  the queue's bucket runs out of tokens. The bucket will be
	//  continuously refilled with new tokens based on
	//  [max_dispatches_per_second][RateLimits.max_dispatches_per_second].
	//
	//  The default value of `max_burst_size` is picked by Cloud Tasks
	//  based on the value of
	//  [max_dispatches_per_second][RateLimits.max_dispatches_per_second].
	//
	//  The maximum value of `max_burst_size` is 500.
	//
	//  For App Engine queues that were created or updated using
	//  `queue.yaml/xml`, `max_burst_size` is equal to
	//  [bucket_size](https://cloud.google.com/appengine/docs/standard/python/config/queueref#bucket_size).
	//  If
	//  [UpdateQueue][google.cloud.tasks.v2beta2.CloudTasks.UpdateQueue] is called
	//  on a queue without explicitly setting a value for `max_burst_size`,
	//  `max_burst_size` value will get updated if
	//  [UpdateQueue][google.cloud.tasks.v2beta2.CloudTasks.UpdateQueue] is
	//  updating [max_dispatches_per_second][RateLimits.max_dispatches_per_second].
	// +kcc:proto:field=google.cloud.tasks.v2beta2.RateLimits.max_burst_size
	MaxBurstSize *int32 `json:"maxBurstSize,omitempty"`

	// The maximum number of concurrent tasks that Cloud Tasks allows
	//  to be dispatched for this queue. After this threshold has been
	//  reached, Cloud Tasks stops dispatching tasks until the number of
	//  concurrent requests decreases.
	//
	//  If unspecified when the queue is created, Cloud Tasks will pick the
	//  default.
	//
	//
	//  The maximum allowed value is 5,000.
	//
	//  This field is output only for
	//  [pull queues][google.cloud.tasks.v2beta2.PullTarget] and always -1, which
	//  indicates no limit. No other queue types can have `max_concurrent_tasks`
	//  set to -1.
	//
	//
	//  This field has the same meaning as
	//  [max_concurrent_requests in
	//  queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#max_concurrent_requests).
	// +kcc:proto:field=google.cloud.tasks.v2beta2.RateLimits.max_concurrent_tasks
	MaxConcurrentTasks *int32 `json:"maxConcurrentTasks,omitempty"`
}

// +kcc:proto=google.cloud.tasks.v2beta2.RetryConfig
type RetryConfig struct {
	// The maximum number of attempts for a task.
	//
	//  Cloud Tasks will attempt the task `max_attempts` times (that
	//  is, if the first attempt fails, then there will be
	//  `max_attempts - 1` retries).  Must be > 0.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.RetryConfig.max_attempts
	MaxAttempts *int32 `json:"maxAttempts,omitempty"`

	// If true, then the number of attempts is unlimited.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.RetryConfig.unlimited_attempts
	UnlimitedAttempts *bool `json:"unlimitedAttempts,omitempty"`

	// If positive, `max_retry_duration` specifies the time limit for
	//  retrying a failed task, measured from when the task was first
	//  attempted. Once `max_retry_duration` time has passed *and* the
	//  task has been attempted
	//  [max_attempts][google.cloud.tasks.v2beta2.RetryConfig.max_attempts] times,
	//  no further attempts will be made and the task will be deleted.
	//
	//  If zero, then the task age is unlimited.
	//
	//  If unspecified when the queue is created, Cloud Tasks will pick the
	//  default.
	//
	//  This field is output only for [pull
	//  queues][google.cloud.tasks.v2beta2.PullTarget].
	//
	//
	//  `max_retry_duration` will be truncated to the nearest second.
	//
	//  This field has the same meaning as
	//  [task_age_limit in
	//  queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	// +kcc:proto:field=google.cloud.tasks.v2beta2.RetryConfig.max_retry_duration
	MaxRetryDuration *string `json:"maxRetryDuration,omitempty"`

	// A task will be [scheduled][google.cloud.tasks.v2beta2.Task.schedule_time]
	//  for retry between
	//  [min_backoff][google.cloud.tasks.v2beta2.RetryConfig.min_backoff] and
	//  [max_backoff][google.cloud.tasks.v2beta2.RetryConfig.max_backoff] duration
	//  after it fails, if the queue's
	//  [RetryConfig][google.cloud.tasks.v2beta2.RetryConfig] specifies that the
	//  task should be retried.
	//
	//  If unspecified when the queue is created, Cloud Tasks will pick the
	//  default.
	//
	//  This field is output only for [pull
	//  queues][google.cloud.tasks.v2beta2.PullTarget].
	//
	//
	//  `min_backoff` will be truncated to the nearest second.
	//
	//  This field has the same meaning as
	//  [min_backoff_seconds in
	//  queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	// +kcc:proto:field=google.cloud.tasks.v2beta2.RetryConfig.min_backoff
	MinBackoff *string `json:"minBackoff,omitempty"`

	// A task will be [scheduled][google.cloud.tasks.v2beta2.Task.schedule_time]
	//  for retry between
	//  [min_backoff][google.cloud.tasks.v2beta2.RetryConfig.min_backoff] and
	//  [max_backoff][google.cloud.tasks.v2beta2.RetryConfig.max_backoff] duration
	//  after it fails, if the queue's
	//  [RetryConfig][google.cloud.tasks.v2beta2.RetryConfig] specifies that the
	//  task should be retried.
	//
	//  If unspecified when the queue is created, Cloud Tasks will pick the
	//  default.
	//
	//  This field is output only for [pull
	//  queues][google.cloud.tasks.v2beta2.PullTarget].
	//
	//
	//  `max_backoff` will be truncated to the nearest second.
	//
	//  This field has the same meaning as
	//  [max_backoff_seconds in
	//  queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	// +kcc:proto:field=google.cloud.tasks.v2beta2.RetryConfig.max_backoff
	MaxBackoff *string `json:"maxBackoff,omitempty"`

	// The time between retries will double `max_doublings` times.
	//
	//  A task's retry interval starts at
	//  [min_backoff][google.cloud.tasks.v2beta2.RetryConfig.min_backoff], then
	//  doubles `max_doublings` times, then increases linearly, and finally retries
	//  at intervals of
	//  [max_backoff][google.cloud.tasks.v2beta2.RetryConfig.max_backoff] up to
	//  [max_attempts][google.cloud.tasks.v2beta2.RetryConfig.max_attempts] times.
	//
	//  For example, if
	//  [min_backoff][google.cloud.tasks.v2beta2.RetryConfig.min_backoff] is 10s,
	//  [max_backoff][google.cloud.tasks.v2beta2.RetryConfig.max_backoff] is 300s,
	//  and `max_doublings` is 3, then the a task will first be retried in 10s. The
	//  retry interval will double three times, and then increase linearly by 2^3 *
	//  10s.  Finally, the task will retry at intervals of
	//  [max_backoff][google.cloud.tasks.v2beta2.RetryConfig.max_backoff] until the
	//  task has been attempted
	//  [max_attempts][google.cloud.tasks.v2beta2.RetryConfig.max_attempts] times.
	//  Thus, the requests will retry at 10s, 20s, 40s, 80s, 160s, 240s, 300s,
	//  300s, ....
	//
	//  If unspecified when the queue is created, Cloud Tasks will pick the
	//  default.
	//
	//  This field is output only for [pull
	//  queues][google.cloud.tasks.v2beta2.PullTarget].
	//
	//
	//  This field has the same meaning as
	//  [max_doublings in
	//  queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	// +kcc:proto:field=google.cloud.tasks.v2beta2.RetryConfig.max_doublings
	MaxDoublings *int32 `json:"maxDoublings,omitempty"`
}

// +kcc:proto=google.cloud.tasks.v2beta2.UriOverride
type UriOverride struct {
	// Scheme override.
	//
	//  When specified, the task URI scheme is replaced by the provided value (HTTP
	//  or HTTPS).
	// +kcc:proto:field=google.cloud.tasks.v2beta2.UriOverride.scheme
	Scheme *string `json:"scheme,omitempty"`

	// Host override.
	//
	//  When specified, replaces the host part of the task URL. For example,
	//  if the task URL is "https://www.google.com," and host value is set to
	//  "example.net", the overridden URI will be changed to "https://example.net."
	//  Host value cannot be an empty string (INVALID_ARGUMENT).
	// +kcc:proto:field=google.cloud.tasks.v2beta2.UriOverride.host
	Host *string `json:"host,omitempty"`

	// Port override.
	//
	//  When specified, replaces the port part of the task URI. For instance,
	//  for a URI http://www.google.com/foo and port=123, the overridden URI
	//  becomes http://www.google.com:123/foo. Note that the port value must be a
	//  positive integer. Setting the port to 0 (Zero) clears the URI port.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.UriOverride.port
	Port *int64 `json:"port,omitempty"`

	// URI path.
	//
	//  When specified, replaces the existing path of the task URL. Setting the
	//  path value to an empty string clears the URI path segment.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.UriOverride.path_override
	PathOverride *PathOverride `json:"pathOverride,omitempty"`

	// URI Query.
	//
	//  When specified, replaces the query part of the task URI. Setting the
	//  query value to an empty string clears the URI query segment.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.UriOverride.query_override
	QueryOverride *QueryOverride `json:"queryOverride,omitempty"`

	// URI Override Enforce Mode
	//
	//  When specified, determines the Target UriOverride mode. If not specified,
	//  it defaults to ALWAYS.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.UriOverride.uri_override_enforce_mode
	URIOverrideEnforceMode *string `json:"uriOverrideEnforceMode,omitempty"`
}

// +kcc:proto=google.cloud.tasks.v2beta2.Queue
type QueueObservedState struct {
	// Output only. The realtime, informational statistics for a queue. In order
	//  to receive the statistics the caller should include this field in the
	//  FieldMask.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.Queue.stats
	Stats *QueueStats `json:"stats,omitempty"`
}

// +kcc:proto=google.cloud.tasks.v2beta2.QueueStats
type QueueStatsObservedState struct {
	// Output only. An estimation of the number of tasks in the queue, that is,
	//  the tasks in the queue that haven't been executed, the tasks in the queue
	//  which the queue has dispatched but has not yet received a reply for, and
	//  the failed tasks that the queue is retrying.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.QueueStats.tasks_count
	TasksCount *int64 `json:"tasksCount,omitempty"`

	// Output only. An estimation of the nearest time in the future where a task
	//  in the queue is scheduled to be executed.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.QueueStats.oldest_estimated_arrival_time
	OldestEstimatedArrivalTime *string `json:"oldestEstimatedArrivalTime,omitempty"`

	// Output only. The number of tasks that the queue has dispatched and received
	//  a reply for during the last minute. This variable counts both successful
	//  and non-successful executions.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.QueueStats.executed_last_minute_count
	ExecutedLastMinuteCount *int64 `json:"executedLastMinuteCount,omitempty"`

	// Output only. The number of requests that the queue has dispatched but has
	//  not received a reply for yet.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.QueueStats.concurrent_dispatches_count
	ConcurrentDispatchesCount *int64 `json:"concurrentDispatchesCount,omitempty"`

	// Output only. The current maximum number of tasks per second executed by the
	//  queue. The maximum value of this variable is controlled by the RateLimits
	//  of the Queue. However, this value could be less to avoid overloading the
	//  endpoints tasks in the queue are targeting.
	// +kcc:proto:field=google.cloud.tasks.v2beta2.QueueStats.effective_execution_rate
	EffectiveExecutionRate *float64 `json:"effectiveExecutionRate,omitempty"`
}
