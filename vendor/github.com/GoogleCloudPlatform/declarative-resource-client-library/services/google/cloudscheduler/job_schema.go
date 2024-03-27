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
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLJobSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "CloudScheduler/Job",
			Description: "The CloudScheduler Job resource",
			StructName:  "Job",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Job",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "job",
						Required:    true,
						Description: "A full instance of a Job",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Job",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "job",
						Required:    true,
						Description: "A full instance of a Job",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Job",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "job",
						Required:    true,
						Description: "A full instance of a Job",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Job",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "project",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "location",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many Job",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "project",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "location",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
		},
		Components: &dcl.Components{
			Schemas: map[string]*dcl.Component{
				"Job": &dcl.Component{
					Title:           "Job",
					ID:              "projects/{{project}}/locations/{{location}}/jobs/{{name}}",
					ParentContainer: "project",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"project",
							"location",
						},
						Properties: map[string]*dcl.Property{
							"appEngineHttpTarget": &dcl.Property{
								Type:        "object",
								GoName:      "AppEngineHttpTarget",
								GoType:      "JobAppEngineHttpTarget",
								Description: "App Engine HTTP target.",
								Conflicts: []string{
									"pubsubTarget",
									"httpTarget",
								},
								Properties: map[string]*dcl.Property{
									"appEngineRouting": &dcl.Property{
										Type:        "object",
										GoName:      "AppEngineRouting",
										GoType:      "JobAppEngineHttpTargetAppEngineRouting",
										Description: "App Engine Routing setting for the job.",
										Properties: map[string]*dcl.Property{
											"host": &dcl.Property{
												Type:        "string",
												GoName:      "Host",
												ReadOnly:    true,
												Description: "Output only. The host that the job is sent to. For more information about how App Engine requests are routed, see [here](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed). The host is constructed as: * `host = [application_domain_name]` `| [service] + '.' + [application_domain_name]` `| [version] + '.' + [application_domain_name]` `| [version_dot_service]+ '.' + [application_domain_name]` `| [instance] + '.' + [application_domain_name]` `| [instance_dot_service] + '.' + [application_domain_name]` `| [instance_dot_version] + '.' + [application_domain_name]` `| [instance_dot_version_dot_service] + '.' + [application_domain_name]` * `application_domain_name` = The domain name of the app, for example .appspot.com, which is associated with the job's project ID. * `service =` service * `version =` version * `version_dot_service =` version `+ '.' +` service * `instance =` instance * `instance_dot_service =` instance `+ '.' +` service * `instance_dot_version =` instance `+ '.' +` version * `instance_dot_version_dot_service =` instance `+ '.' +` version `+ '.' +` service If service is empty, then the job will be sent to the service which is the default service when the job is attempted. If version is empty, then the job will be sent to the version which is the default version when the job is attempted. If instance is empty, then the job will be sent to an instance which is available when the job is attempted. If service, version, or instance is invalid, then the job will be sent to the default version of the default service when the job is attempted.",
											},
											"instance": &dcl.Property{
												Type:        "string",
												GoName:      "Instance",
												Description: "App instance. By default, the job is sent to an instance which is available when the job is attempted. Requests can only be sent to a specific instance if [manual scaling is used in App Engine Standard](https://cloud.google.com/appengine/docs/python/an-overview-of-app-engine?hl=en_US#scaling_types_and_instance_classes). App Engine Flex does not support instances. For more information, see [App Engine Standard request routing](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed) and [App Engine Flex request routing](https://cloud.google.com/appengine/docs/flexible/python/how-requests-are-routed).",
											},
											"service": &dcl.Property{
												Type:        "string",
												GoName:      "Service",
												Description: "App service. By default, the job is sent to the service which is the default service when the job is attempted.",
											},
											"version": &dcl.Property{
												Type:        "string",
												GoName:      "Version",
												Description: "App version. By default, the job is sent to the version which is the default version when the job is attempted.",
											},
										},
									},
									"body": &dcl.Property{
										Type:        "string",
										GoName:      "Body",
										Description: "Body. HTTP request body. A request body is allowed only if the HTTP method is POST or PUT. It will result in invalid argument error to set a body on a job with an incompatible HttpMethod.",
									},
									"headers": &dcl.Property{
										Type: "object",
										AdditionalProperties: &dcl.Property{
											Type: "string",
										},
										GoName:      "Headers",
										Description: "HTTP request headers. This map contains the header field names and values. Headers can be set when the job is created. Cloud Scheduler sets some headers to default values: * `User-Agent`: By default, this header is `\"App Engine-Google; (+http://code.google.com/appengine)\"`. This header can be modified, but Cloud Scheduler will append `\"App Engine-Google; (+http://code.google.com/appengine)\"` to the modified `User-Agent`. * `X-CloudScheduler`: This header will be set to true. The headers below are output only. They cannot be set or overridden: * `X-Google-*`: For Google internal use only. * `X-App Engine-*`: For Google internal use only. In addition, some App Engine headers, which contain job-specific information, are also be sent to the job handler.",
									},
									"httpMethod": &dcl.Property{
										Type:        "string",
										GoName:      "HttpMethod",
										GoType:      "JobAppEngineHttpTargetHttpMethodEnum",
										Description: "The HTTP method to use for the request. PATCH and OPTIONS are not permitted. Possible values: HTTP_METHOD_UNSPECIFIED, POST, GET, HEAD, PUT, DELETE, PATCH, OPTIONS",
										Enum: []string{
											"HTTP_METHOD_UNSPECIFIED",
											"POST",
											"GET",
											"HEAD",
											"PUT",
											"DELETE",
											"PATCH",
											"OPTIONS",
										},
									},
									"relativeUri": &dcl.Property{
										Type:        "string",
										GoName:      "RelativeUri",
										Description: "The relative URI. The relative URL must begin with \"/\" and must be a valid HTTP relative URL. It can contain a path, query string arguments, and `#` fragments. If the relative URL is empty, then the root path \"/\" will be used. No spaces are allowed, and the maximum length allowed is 2083 characters.",
									},
								},
							},
							"attemptDeadline": &dcl.Property{
								Type:        "string",
								GoName:      "AttemptDeadline",
								Description: "The deadline for job attempts. If the request handler does not respond by this deadline then the request is cancelled and the attempt is marked as a `DEADLINE_EXCEEDED` failure. The failed attempt can be viewed in execution logs. Cloud Scheduler will retry the job according to the RetryConfig. The allowed duration for this deadline is: * For HTTP targets, between 15 seconds and 30 minutes. * For App Engine HTTP targets, between 15 seconds and 24 hours.",
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "Optionally caller-specified in CreateJob or UpdateJob. A human-readable description for the job. This string must not contain more than 500 characters.",
							},
							"httpTarget": &dcl.Property{
								Type:        "object",
								GoName:      "HttpTarget",
								GoType:      "JobHttpTarget",
								Description: "HTTP target.",
								Conflicts: []string{
									"pubsubTarget",
									"appEngineHttpTarget",
								},
								Required: []string{
									"uri",
								},
								Properties: map[string]*dcl.Property{
									"body": &dcl.Property{
										Type:        "string",
										GoName:      "Body",
										Description: "HTTP request body. A request body is allowed only if the HTTP method is POST, PUT, or PATCH. It is an error to set body on a job with an incompatible HttpMethod.",
									},
									"headers": &dcl.Property{
										Type: "object",
										AdditionalProperties: &dcl.Property{
											Type: "string",
										},
										GoName:      "Headers",
										Description: "The user can specify HTTP request headers to send with the job's HTTP request. This map contains the header field names and values. Repeated headers are not supported, but a header value can contain commas. These headers represent a subset of the headers that will accompany the job's HTTP request. Some HTTP request headers will be ignored or replaced. A partial list of headers that will be ignored or replaced is below: - Host: This will be computed by Cloud Scheduler and derived from uri. * `Content-Length`: This will be computed by Cloud Scheduler. * `User-Agent`: This will be set to `\"Google-Cloud-Scheduler\"`. * `X-Google-*`: Google internal use only. * `X-appengine-*`: Google internal use only. The total size of headers must be less than 80KB.",
									},
									"httpMethod": &dcl.Property{
										Type:        "string",
										GoName:      "HttpMethod",
										GoType:      "JobHttpTargetHttpMethodEnum",
										Description: "Which HTTP method to use for the request. Possible values: HTTP_METHOD_UNSPECIFIED, POST, GET, HEAD, PUT, DELETE, PATCH, OPTIONS",
										Enum: []string{
											"HTTP_METHOD_UNSPECIFIED",
											"POST",
											"GET",
											"HEAD",
											"PUT",
											"DELETE",
											"PATCH",
											"OPTIONS",
										},
									},
									"oauthToken": &dcl.Property{
										Type:        "object",
										GoName:      "OAuthToken",
										GoType:      "JobHttpTargetOAuthToken",
										Description: "If specified, an [OAuth token](https://developers.google.com/identity/protocols/OAuth2) will be generated and attached as an `Authorization` header in the HTTP request. This type of authorization should generally only be used when calling Google APIs hosted on *.googleapis.com.",
										Properties: map[string]*dcl.Property{
											"scope": &dcl.Property{
												Type:        "string",
												GoName:      "Scope",
												Description: "OAuth scope to be used for generating OAuth access token. If not specified, \"https://www.googleapis.com/auth/cloud-platform\" will be used.",
											},
											"serviceAccountEmail": &dcl.Property{
												Type:        "string",
												GoName:      "ServiceAccountEmail",
												Description: "[Service account email](https://cloud.google.com/iam/docs/service-accounts) to be used for generating OAuth token. The service account must be within the same project as the job. The caller must have iam.serviceAccounts.actAs permission for the service account.",
												ResourceReferences: []*dcl.PropertyResourceReference{
													&dcl.PropertyResourceReference{
														Resource: "Iam/ServiceAccount",
														Field:    "email",
													},
												},
											},
										},
									},
									"oidcToken": &dcl.Property{
										Type:        "object",
										GoName:      "OidcToken",
										GoType:      "JobHttpTargetOidcToken",
										Description: "If specified, an [OIDC](https://developers.google.com/identity/protocols/OpenIDConnect) token will be generated and attached as an `Authorization` header in the HTTP request. This type of authorization can be used for many scenarios, including calling Cloud Run, or endpoints where you intend to validate the token yourself.",
										Properties: map[string]*dcl.Property{
											"audience": &dcl.Property{
												Type:        "string",
												GoName:      "Audience",
												Description: "Audience to be used when generating OIDC token. If not specified, the URI specified in target will be used.",
											},
											"serviceAccountEmail": &dcl.Property{
												Type:        "string",
												GoName:      "ServiceAccountEmail",
												Description: "[Service account email](https://cloud.google.com/iam/docs/service-accounts) to be used for generating OIDC token. The service account must be within the same project as the job. The caller must have iam.serviceAccounts.actAs permission for the service account.",
												ResourceReferences: []*dcl.PropertyResourceReference{
													&dcl.PropertyResourceReference{
														Resource: "Iam/ServiceAccount",
														Field:    "email",
													},
												},
											},
										},
									},
									"uri": &dcl.Property{
										Type:        "string",
										GoName:      "Uri",
										Description: "Required. The full URI path that the request will be sent to. This string must begin with either \"http://\" or \"https://\". Some examples of valid values for uri are: `http://acme.com` and `https://acme.com/sales:8080`. Cloud Scheduler will encode some characters for safety and compatibility. The maximum allowed URL length is 2083 characters after encoding.",
									},
								},
							},
							"lastAttemptTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "LastAttemptTime",
								ReadOnly:    true,
								Description: "Output only. The time the last job attempt started.",
								Immutable:   true,
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location for the resource",
								Immutable:   true,
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "Optionally caller-specified in CreateJob, after which it becomes output only. The job name. For example: `projects/PROJECT_ID/locations/LOCATION_ID/jobs/JOB_ID`. * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), or periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects) * `LOCATION_ID` is the canonical ID for the job's location. The list of available locations can be obtained by calling ListLocations. For more information, see https://cloud.google.com/about/locations/. * `JOB_ID` can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), or underscores (_). The maximum length is 500 characters.",
							},
							"project": &dcl.Property{
								Type:        "string",
								GoName:      "Project",
								Description: "The project for the resource",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Project",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"pubsubTarget": &dcl.Property{
								Type:        "object",
								GoName:      "PubsubTarget",
								GoType:      "JobPubsubTarget",
								Description: "Pub/Sub target.",
								Conflicts: []string{
									"appEngineHttpTarget",
									"httpTarget",
								},
								Required: []string{
									"topicName",
								},
								Properties: map[string]*dcl.Property{
									"attributes": &dcl.Property{
										Type: "object",
										AdditionalProperties: &dcl.Property{
											Type: "string",
										},
										GoName:      "Attributes",
										Description: "Attributes for PubsubMessage. Pubsub message must contain either non-empty data, or at least one attribute.",
									},
									"data": &dcl.Property{
										Type:        "string",
										GoName:      "Data",
										Description: "The message payload for PubsubMessage. Pubsub message must contain either non-empty data, or at least one attribute.",
									},
									"topicName": &dcl.Property{
										Type:        "string",
										GoName:      "TopicName",
										Description: "Required. The name of the Cloud Pub/Sub topic to which messages will be published when a job is delivered. The topic name must be in the same format as required by Pub/Sub's [PublishRequest.name](https://cloud.google.com/pubsub/docs/reference/rpc/google.pubsub.v1#publishrequest), for example `projects/PROJECT_ID/topics/TOPIC_ID`. The topic must be in the same project as the Cloud Scheduler job.",
										ResourceReferences: []*dcl.PropertyResourceReference{
											&dcl.PropertyResourceReference{
												Resource: "Pubsub/Topic",
												Field:    "name",
											},
										},
									},
								},
							},
							"retryConfig": &dcl.Property{
								Type:        "object",
								GoName:      "RetryConfig",
								GoType:      "JobRetryConfig",
								Description: "Settings that determine the retry behavior.",
								Properties: map[string]*dcl.Property{
									"maxBackoffDuration": &dcl.Property{
										Type:        "string",
										GoName:      "MaxBackoffDuration",
										Description: "The maximum amount of time to wait before retrying a job after it fails. The default value of this field is 1 hour.",
									},
									"maxDoublings": &dcl.Property{
										Type:        "integer",
										Format:      "int64",
										GoName:      "MaxDoublings",
										Description: "The time between retries will double `max_doublings` times. A job's retry interval starts at min_backoff_duration, then doubles `max_doublings` times, then increases linearly, and finally retries at intervals of max_backoff_duration up to retry_count times. For example, if min_backoff_duration is 10s, max_backoff_duration is 300s, and `max_doublings` is 3, then the a job will first be retried in 10s. The retry interval will double three times, and then increase linearly by 2^3 * 10s. Finally, the job will retry at intervals of max_backoff_duration until the job has been attempted retry_count times. Thus, the requests will retry at 10s, 20s, 40s, 80s, 160s, 240s, 300s, 300s, .... The default value of this field is 5.",
									},
									"maxRetryDuration": &dcl.Property{
										Type:        "string",
										GoName:      "MaxRetryDuration",
										Description: "The time limit for retrying a failed job, measured from time when an execution was first attempted. If specified with retry_count, the job will be retried until both limits are reached. The default value for max_retry_duration is zero, which means retry duration is unlimited.",
									},
									"minBackoffDuration": &dcl.Property{
										Type:        "string",
										GoName:      "MinBackoffDuration",
										Description: "The minimum amount of time to wait before retrying a job after it fails. The default value of this field is 5 seconds.",
									},
									"retryCount": &dcl.Property{
										Type:        "integer",
										Format:      "int64",
										GoName:      "RetryCount",
										Description: "The number of attempts that the system will make to run a job using the exponential backoff procedure described by max_doublings. The default value of retry_count is zero. If retry_count is zero, a job attempt will *not* be retried if it fails. Instead the Cloud Scheduler system will wait for the next scheduled execution time. If retry_count is set to a non-zero number then Cloud Scheduler will retry failed attempts, using exponential backoff, retry_count times, or until the next scheduled execution time, whichever comes first. Values greater than 5 and negative values are not allowed.",
									},
								},
							},
							"schedule": &dcl.Property{
								Type:        "string",
								GoName:      "Schedule",
								Description: "Required, except when used with UpdateJob. Describes the schedule on which the job will be executed. The schedule can be either of the following types: * [Crontab](http://en.wikipedia.org/wiki/Cron#Overview) * English-like [schedule](https://cloud.google.com/scheduler/docs/configuring/cron-job-schedules) As a general rule, execution `n + 1` of a job will not begin until execution `n` has finished. Cloud Scheduler will never allow two simultaneously outstanding executions. For example, this implies that if the `n+1`th execution is scheduled to run at 16:00 but the `n`th execution takes until 16:15, the `n+1`th execution will not start until `16:15`. A scheduled start time will be delayed if the previous execution has not ended when its scheduled time occurs. If retry_count > 0 and a job attempt fails, the job will be tried a total of retry_count times, with exponential backoff, until the next scheduled start time.",
							},
							"scheduleTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "ScheduleTime",
								ReadOnly:    true,
								Description: "Output only. The next time the job is scheduled. Note that this may be a retry of a previously failed attempt or the next execution time according to the schedule.",
								Immutable:   true,
							},
							"state": &dcl.Property{
								Type:        "string",
								GoName:      "State",
								GoType:      "JobStateEnum",
								ReadOnly:    true,
								Description: "Output only. State of the job. Possible values: STATE_UNSPECIFIED, ENABLED, PAUSED, DISABLED, UPDATE_FAILED",
								Immutable:   true,
								Enum: []string{
									"STATE_UNSPECIFIED",
									"ENABLED",
									"PAUSED",
									"DISABLED",
									"UPDATE_FAILED",
								},
							},
							"status": &dcl.Property{
								Type:        "object",
								GoName:      "Status",
								GoType:      "JobStatus",
								ReadOnly:    true,
								Description: "Output only. The response from the target for the last attempted execution.",
								Immutable:   true,
								Properties: map[string]*dcl.Property{
									"code": &dcl.Property{
										Type:        "integer",
										Format:      "int64",
										GoName:      "Code",
										Description: "The status code, which should be an enum value of google.rpc.Code.",
										Immutable:   true,
									},
									"details": &dcl.Property{
										Type:        "array",
										GoName:      "Details",
										Description: "A list of messages that carry the error details. There is a common set of message types for APIs to use.",
										Immutable:   true,
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "object",
											GoType: "JobStatusDetails",
											Properties: map[string]*dcl.Property{
												"typeUrl": &dcl.Property{
													Type:        "string",
													GoName:      "TypeUrl",
													Description: "A URL/resource name that uniquely identifies the type of the serialized protocol buffer message. This string must contain at least one \"/\" character. The last segment of the URL's path must represent the fully qualified name of the type (as in `path/google.protobuf.Duration`). The name should be in a canonical form (e.g., leading \".\" is not accepted). In practice, teams usually precompile into the binary all types that they expect it to use in the context of Any. However, for URLs which use the scheme `http`, `https`, or no scheme, one can optionally set up a type server that maps type URLs to message definitions as follows: * If no scheme is provided, `https` is assumed. * An HTTP GET on the URL must yield a google.protobuf.Type value in binary format, or produce an error. * Applications are allowed to cache lookup results based on the URL, or have them precompiled into a binary to avoid any lookup. Therefore, binary compatibility needs to be preserved on changes to types. (Use versioned type names to manage breaking changes.) Note: this functionality is not currently available in the official protobuf release, and it is not used for type URLs beginning with type.googleapis.com. Schemes other than `http`, `https` (or the empty scheme) might be used with implementation specific semantics.",
													Immutable:   true,
												},
												"value": &dcl.Property{
													Type:        "string",
													GoName:      "Value",
													Description: "Must be a valid serialized protocol buffer of the above specified type.",
													Immutable:   true,
												},
											},
										},
									},
									"message": &dcl.Property{
										Type:        "string",
										GoName:      "Message",
										Description: "A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client.",
										Immutable:   true,
									},
								},
							},
							"timeZone": &dcl.Property{
								Type:          "string",
								GoName:        "TimeZone",
								Description:   "Specifies the time zone to be used in interpreting schedule. The value of this field must be a time zone name from the [tz database](http://en.wikipedia.org/wiki/Tz_database). Note that some time zones include a provision for daylight savings time. The rules for daylight saving time are determined by the chosen tz. For UTC use the string \"utc\". If a time zone is not specified, the default will be in UTC (also known as GMT).",
								ServerDefault: true,
							},
							"userUpdateTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "UserUpdateTime",
								ReadOnly:    true,
								Description: "Output only. The creation time of the job.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
