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
