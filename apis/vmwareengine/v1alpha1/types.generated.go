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


// +kcc:proto=google.cloud.vmwareengine.v1.LoggingServer
type LoggingServer struct {

	// Required. Fully-qualified domain name (FQDN) or IP Address of the logging
	//  server.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.LoggingServer.hostname
	Hostname *string `json:"hostname,omitempty"`

	// Required. Port number at which the logging server receives logs.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.LoggingServer.port
	Port *int32 `json:"port,omitempty"`

	// Required. Protocol used by vCenter to send logs to a logging server.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.LoggingServer.protocol
	Protocol *string `json:"protocol,omitempty"`

	// Required. The type of component that produces logs that will be forwarded
	//  to this logging server.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.LoggingServer.source_type
	SourceType *string `json:"sourceType,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.LoggingServer
type LoggingServerObservedState struct {
	// Output only. The resource name of this logging server.
	//  Resource names are schemeless URIs that follow the conventions in
	//  https://cloud.google.com/apis/design/resource_names.
	//  For example:
	//  `projects/my-project/locations/us-central1-a/privateClouds/my-cloud/loggingServers/my-logging-server`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.LoggingServer.name
	Name *string `json:"name,omitempty"`

	// Output only. Creation time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.LoggingServer.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.LoggingServer.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. System-generated unique identifier for the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.LoggingServer.uid
	Uid *string `json:"uid,omitempty"`
}
