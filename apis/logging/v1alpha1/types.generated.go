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

// +kcc:proto=google.api.MonitoredResource
type MonitoredResource struct {
	// Required. The monitored resource type. This field must match
	//  the `type` field of a
	//  [MonitoredResourceDescriptor][google.api.MonitoredResourceDescriptor]
	//  object. For example, the type of a Compute Engine VM instance is
	//  `gce_instance`. Some descriptors include the service name in the type; for
	//  example, the type of a Datastream stream is
	//  `datastream.googleapis.com/Stream`.
	// +kcc:proto:field=google.api.MonitoredResource.type
	Type *string `json:"type,omitempty"`

	// Required. Values for all of the labels listed in the associated monitored
	//  resource descriptor. For example, Compute Engine VM instances use the
	//  labels `"project_id"`, `"instance_id"`, and `"zone"`.
	// +kcc:proto:field=google.api.MonitoredResource.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.logging.type.HttpRequest
type HTTPRequest struct {
	// The request method. Examples: `"GET"`, `"HEAD"`, `"PUT"`, `"POST"`.
	// +kcc:proto:field=google.logging.type.HttpRequest.request_method
	RequestMethod *string `json:"requestMethod,omitempty"`

	// The scheme (http, https), the host name, the path and the query
	//  portion of the URL that was requested.
	//  Example: `"http://example.com/some/info?color=red"`.
	// +kcc:proto:field=google.logging.type.HttpRequest.request_url
	RequestURL *string `json:"requestURL,omitempty"`

	// The size of the HTTP request message in bytes, including the request
	//  headers and the request body.
	// +kcc:proto:field=google.logging.type.HttpRequest.request_size
	RequestSize *int64 `json:"requestSize,omitempty"`

	// The response code indicating the status of response.
	//  Examples: 200, 404.
	// +kcc:proto:field=google.logging.type.HttpRequest.status
	Status *int32 `json:"status,omitempty"`

	// The size of the HTTP response message sent back to the client, in bytes,
	//  including the response headers and the response body.
	// +kcc:proto:field=google.logging.type.HttpRequest.response_size
	ResponseSize *int64 `json:"responseSize,omitempty"`

	// The user agent sent by the client. Example:
	//  `"Mozilla/4.0 (compatible; MSIE 6.0; Windows 98; Q312461; .NET
	//  CLR 1.0.3705)"`.
	// +kcc:proto:field=google.logging.type.HttpRequest.user_agent
	UserAgent *string `json:"userAgent,omitempty"`

	// The IP address (IPv4 or IPv6) of the client that issued the HTTP
	//  request. This field can include port information. Examples:
	//  `"192.168.1.1"`, `"10.0.0.1:80"`, `"FE80::0202:B3FF:FE1E:8329"`.
	// +kcc:proto:field=google.logging.type.HttpRequest.remote_ip
	RemoteIP *string `json:"remoteIP,omitempty"`

	// The IP address (IPv4 or IPv6) of the origin server that the request was
	//  sent to. This field can include port information. Examples:
	//  `"192.168.1.1"`, `"10.0.0.1:80"`, `"FE80::0202:B3FF:FE1E:8329"`.
	// +kcc:proto:field=google.logging.type.HttpRequest.server_ip
	ServerIP *string `json:"serverIP,omitempty"`

	// The referer URL of the request, as defined in
	//  [HTTP/1.1 Header Field
	//  Definitions](https://datatracker.ietf.org/doc/html/rfc2616#section-14.36).
	// +kcc:proto:field=google.logging.type.HttpRequest.referer
	Referer *string `json:"referer,omitempty"`

	// The request processing latency on the server, from the time the request was
	//  received until the response was sent.
	// +kcc:proto:field=google.logging.type.HttpRequest.latency
	Latency *string `json:"latency,omitempty"`

	// Whether or not a cache lookup was attempted.
	// +kcc:proto:field=google.logging.type.HttpRequest.cache_lookup
	CacheLookup *bool `json:"cacheLookup,omitempty"`

	// Whether or not an entity was served from cache
	//  (with or without validation).
	// +kcc:proto:field=google.logging.type.HttpRequest.cache_hit
	CacheHit *bool `json:"cacheHit,omitempty"`

	// Whether or not the response was validated with the origin server before
	//  being served from cache. This field is only meaningful if `cache_hit` is
	//  True.
	// +kcc:proto:field=google.logging.type.HttpRequest.cache_validated_with_origin_server
	CacheValidatedWithOriginServer *bool `json:"cacheValidatedWithOriginServer,omitempty"`

	// The number of HTTP response bytes inserted into cache. Set only when a
	//  cache fill was attempted.
	// +kcc:proto:field=google.logging.type.HttpRequest.cache_fill_bytes
	CacheFillBytes *int64 `json:"cacheFillBytes,omitempty"`

	// Protocol used for the request. Examples: "HTTP/1.1", "HTTP/2", "websocket"
	// +kcc:proto:field=google.logging.type.HttpRequest.protocol
	Protocol *string `json:"protocol,omitempty"`
}

// +kcc:proto=google.logging.v2.LogEntry
type LogEntry struct {
	// Required. The resource name of the log to which this log entry belongs:
	//
	//      "projects/[PROJECT_ID]/logs/[LOG_ID]"
	//      "organizations/[ORGANIZATION_ID]/logs/[LOG_ID]"
	//      "billingAccounts/[BILLING_ACCOUNT_ID]/logs/[LOG_ID]"
	//      "folders/[FOLDER_ID]/logs/[LOG_ID]"
	//
	//  A project number may be used in place of PROJECT_ID. The project number is
	//  translated to its corresponding PROJECT_ID internally and the `log_name`
	//  field will contain PROJECT_ID in queries and exports.
	//
	//  `[LOG_ID]` must be URL-encoded within `log_name`. Example:
	//  `"organizations/1234567890/logs/cloudresourcemanager.googleapis.com%2Factivity"`.
	//
	//  `[LOG_ID]` must be less than 512 characters long and can only include the
	//  following characters: upper and lower case alphanumeric characters,
	//  forward-slash, underscore, hyphen, and period.
	//
	//  For backward compatibility, if `log_name` begins with a forward-slash, such
	//  as `/projects/...`, then the log entry is ingested as usual, but the
	//  forward-slash is removed. Listing the log entry will not show the leading
	//  slash and filtering for a log name with a leading slash will never return
	//  any results.
	// +kcc:proto:field=google.logging.v2.LogEntry.log_name
	LogName *string `json:"logName,omitempty"`

	// Required. The monitored resource that produced this log entry.
	//
	//  Example: a log entry that reports a database error would be associated with
	//  the monitored resource designating the particular database that reported
	//  the error.
	// +kcc:proto:field=google.logging.v2.LogEntry.resource
	Resource *MonitoredResource `json:"resource,omitempty"`

	// The log entry payload, represented as a protocol buffer. Some Google
	//  Cloud Platform services use this field for their log entry payloads.
	//
	//  The following protocol buffer types are supported; user-defined types
	//  are not supported:
	//
	//    "type.googleapis.com/google.cloud.audit.AuditLog"
	//    "type.googleapis.com/google.appengine.logging.v1.RequestLog"
	// +kcc:proto:field=google.logging.v2.LogEntry.proto_payload
	ProtoPayload *Any `json:"protoPayload,omitempty"`

	// The log entry payload, represented as a Unicode string (UTF-8).
	// +kcc:proto:field=google.logging.v2.LogEntry.text_payload
	TextPayload *string `json:"textPayload,omitempty"`

	// The log entry payload, represented as a structure that is
	//  expressed as a JSON object.
	// +kcc:proto:field=google.logging.v2.LogEntry.json_payload
	JsonPayload map[string]string `json:"jsonPayload,omitempty"`

	// Optional. The time the event described by the log entry occurred. This time
	//  is used to compute the log entry's age and to enforce the logs retention
	//  period. If this field is omitted in a new log entry, then Logging assigns
	//  it the current time. Timestamps have nanosecond accuracy, but trailing
	//  zeros in the fractional seconds might be omitted when the timestamp is
	//  displayed.
	//
	//  Incoming log entries must have timestamps that don't exceed the
	//  [logs retention
	//  period](https://cloud.google.com/logging/quotas#logs_retention_periods) in
	//  the past, and that don't exceed 24 hours in the future. Log entries outside
	//  those time boundaries aren't ingested by Logging.
	// +kcc:proto:field=google.logging.v2.LogEntry.timestamp
	Timestamp *string `json:"timestamp,omitempty"`

	// Optional. The severity of the log entry. The default value is
	//  `LogSeverity.DEFAULT`.
	// +kcc:proto:field=google.logging.v2.LogEntry.severity
	Severity *string `json:"severity,omitempty"`

	// Optional. A unique identifier for the log entry. If you provide a value,
	//  then Logging considers other log entries in the same project, with the same
	//  `timestamp`, and with the same `insert_id` to be duplicates which are
	//  removed in a single query result. However, there are no guarantees of
	//  de-duplication in the export of logs.
	//
	//  If the `insert_id` is omitted when writing a log entry, the Logging API
	//  assigns its own unique identifier in this field.
	//
	//  In queries, the `insert_id` is also used to order log entries that have
	//  the same `log_name` and `timestamp` values.
	// +kcc:proto:field=google.logging.v2.LogEntry.insert_id
	InsertID *string `json:"insertID,omitempty"`

	// Optional. Information about the HTTP request associated with this log
	//  entry, if applicable.
	// +kcc:proto:field=google.logging.v2.LogEntry.http_request
	HTTPRequest *HTTPRequest `json:"httpRequest,omitempty"`

	// Optional. A map of key, value pairs that provides additional information
	//  about the log entry. The labels can be user-defined or system-defined.
	//
	//  User-defined labels are arbitrary key, value pairs that you can use to
	//  classify logs.
	//
	//  System-defined labels are defined by GCP services for platform logs.
	//  They have two components - a service namespace component and the
	//  attribute name. For example: `compute.googleapis.com/resource_name`.
	//
	//  Cloud Logging truncates label keys that exceed 512 B and label
	//  values that exceed 64 KB upon their associated log entry being
	//  written. The truncation is indicated by an ellipsis at the
	//  end of the character string.
	// +kcc:proto:field=google.logging.v2.LogEntry.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Information about an operation associated with the log entry, if
	//  applicable.
	// +kcc:proto:field=google.logging.v2.LogEntry.operation
	Operation *LogEntryOperation `json:"operation,omitempty"`

	// Optional. The REST resource name of the trace being written to
	//  [Cloud Trace](https://cloud.google.com/trace) in
	//  association with this log entry. For example, if your trace data is stored
	//  in the Cloud project "my-trace-project" and if the service that is creating
	//  the log entry receives a trace header that includes the trace ID "12345",
	//  then the service should use "projects/my-tracing-project/traces/12345".
	//
	//  The `trace` field provides the link between logs and traces. By using
	//  this field, you can navigate from a log entry to a trace.
	// +kcc:proto:field=google.logging.v2.LogEntry.trace
	Trace *string `json:"trace,omitempty"`

	// Optional. The ID of the [Cloud Trace](https://cloud.google.com/trace) span
	//  associated with the current operation in which the log is being written.
	//  For example, if a span has the REST resource name of
	//  "projects/some-project/traces/some-trace/spans/some-span-id", then the
	//  `span_id` field is "some-span-id".
	//
	//  A
	//  [Span](https://cloud.google.com/trace/docs/reference/v2/rest/v2/projects.traces/batchWrite#Span)
	//  represents a single operation within a trace. Whereas a trace may involve
	//  multiple different microservices running on multiple different machines,
	//  a span generally corresponds to a single logical operation being performed
	//  in a single instance of a microservice on one specific machine. Spans
	//  are the nodes within the tree that is a trace.
	//
	//  Applications that are [instrumented for
	//  tracing](https://cloud.google.com/trace/docs/setup) will generally assign a
	//  new, unique span ID on each incoming request. It is also common to create
	//  and record additional spans corresponding to internal processing elements
	//  as well as issuing requests to dependencies.
	//
	//  The span ID is expected to be a 16-character, hexadecimal encoding of an
	//  8-byte array and should not be zero. It should be unique within the trace
	//  and should, ideally, be generated in a manner that is uniformly random.
	//
	//  Example values:
	//
	//    - `000000000000004a`
	//    - `7a2190356c3fc94b`
	//    - `0000f00300090021`
	//    - `d39223e101960076`
	// +kcc:proto:field=google.logging.v2.LogEntry.span_id
	SpanID *string `json:"spanID,omitempty"`

	// Optional. The sampling decision of the trace associated with the log entry.
	//
	//  True means that the trace resource name in the `trace` field was sampled
	//  for storage in a trace backend. False means that the trace was not sampled
	//  for storage when this log entry was written, or the sampling decision was
	//  unknown at the time. A non-sampled `trace` value is still useful as a
	//  request correlation identifier. The default is False.
	// +kcc:proto:field=google.logging.v2.LogEntry.trace_sampled
	TraceSampled *bool `json:"traceSampled,omitempty"`

	// Optional. Source code location information associated with the log entry,
	//  if any.
	// +kcc:proto:field=google.logging.v2.LogEntry.source_location
	SourceLocation *LogEntrySourceLocation `json:"sourceLocation,omitempty"`

	// Optional. Information indicating this LogEntry is part of a sequence of
	//  multiple log entries split from a single LogEntry.
	// +kcc:proto:field=google.logging.v2.LogEntry.split
	Split *LogSplit `json:"split,omitempty"`
}

// +kcc:proto=google.logging.v2.LogEntryOperation
type LogEntryOperation struct {
	// Optional. An arbitrary operation identifier. Log entries with the same
	//  identifier are assumed to be part of the same operation.
	// +kcc:proto:field=google.logging.v2.LogEntryOperation.id
	ID *string `json:"id,omitempty"`

	// Optional. An arbitrary producer identifier. The combination of `id` and
	//  `producer` must be globally unique. Examples for `producer`:
	//  `"MyDivision.MyBigCompany.com"`, `"github.com/MyProject/MyApplication"`.
	// +kcc:proto:field=google.logging.v2.LogEntryOperation.producer
	Producer *string `json:"producer,omitempty"`

	// Optional. Set this to True if this is the first log entry in the operation.
	// +kcc:proto:field=google.logging.v2.LogEntryOperation.first
	First *bool `json:"first,omitempty"`

	// Optional. Set this to True if this is the last log entry in the operation.
	// +kcc:proto:field=google.logging.v2.LogEntryOperation.last
	Last *bool `json:"last,omitempty"`
}

// +kcc:proto=google.logging.v2.LogEntrySourceLocation
type LogEntrySourceLocation struct {
	// Optional. Source file name. Depending on the runtime environment, this
	//  might be a simple name or a fully-qualified name.
	// +kcc:proto:field=google.logging.v2.LogEntrySourceLocation.file
	File *string `json:"file,omitempty"`

	// Optional. Line within the source file. 1-based; 0 indicates no line number
	//  available.
	// +kcc:proto:field=google.logging.v2.LogEntrySourceLocation.line
	Line *int64 `json:"line,omitempty"`

	// Optional. Human-readable name of the function or method being invoked, with
	//  optional context such as the class or package name. This information may be
	//  used in contexts such as the logs viewer, where a file and line number are
	//  less meaningful. The format can vary by language. For example:
	//  `qual.if.ied.Class.method` (Java), `dir/package.func` (Go), `function`
	//  (Python).
	// +kcc:proto:field=google.logging.v2.LogEntrySourceLocation.function
	Function *string `json:"function,omitempty"`
}

// +kcc:proto=google.logging.v2.LogSplit
type LogSplit struct {
	// A globally unique identifier for all log entries in a sequence of split log
	//  entries. All log entries with the same |LogSplit.uid| are assumed to be
	//  part of the same sequence of split log entries.
	// +kcc:proto:field=google.logging.v2.LogSplit.uid
	Uid *string `json:"uid,omitempty"`

	// The index of this LogEntry in the sequence of split log entries. Log
	//  entries are given |index| values 0, 1, ..., n-1 for a sequence of n log
	//  entries.
	// +kcc:proto:field=google.logging.v2.LogSplit.index
	Index *int32 `json:"index,omitempty"`

	// The total number of log entries that the original LogEntry was split into.
	// +kcc:proto:field=google.logging.v2.LogSplit.total_splits
	TotalSplits *int32 `json:"totalSplits,omitempty"`
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

// +kcc:proto=google.logging.v2.LogEntry
type LogEntryObservedState struct {
	// Output only. The time the log entry was received by Logging.
	// +kcc:proto:field=google.logging.v2.LogEntry.receive_timestamp
	ReceiveTimestamp *string `json:"receiveTimestamp,omitempty"`
}
