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


// +kcc:proto=google.cloud.servicedirectory.v1beta1.Endpoint
type Endpoint struct {
	// Immutable. The resource name for the endpoint in the format
	//  `projects/*/locations/*/namespaces/*/services/*/endpoints/*`.
	// +kcc:proto:field=google.cloud.servicedirectory.v1beta1.Endpoint.name
	Name *string `json:"name,omitempty"`

	// Optional. An IPv4 or IPv6 address. Service Directory rejects bad addresses
	//  like:
	//
	//  *   `8.8.8`
	//  *   `8.8.8.8:53`
	//  *   `test:bad:address`
	//  *   `[::1]`
	//  *   `[::1]:8080`
	//
	//  Limited to 45 characters.
	// +kcc:proto:field=google.cloud.servicedirectory.v1beta1.Endpoint.address
	Address *string `json:"address,omitempty"`

	// Optional. Service Directory rejects values outside of `[0, 65535]`.
	// +kcc:proto:field=google.cloud.servicedirectory.v1beta1.Endpoint.port
	Port *int32 `json:"port,omitempty"`

	// Optional. Metadata for the endpoint. This data can be consumed by service
	//  clients.
	//
	//  Restrictions:
	//
	//  *   The entire metadata dictionary may contain up to 512 characters,
	//      spread accoss all key-value pairs. Metadata that goes beyond this
	//      limit are rejected
	//  *   Valid metadata keys have two segments: an optional prefix and name,
	//      separated by a slash (/). The name segment is required and must be 63
	//      characters or less, beginning and ending with an alphanumeric character
	//      ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and
	//      alphanumerics between. The prefix is optional. If specified, the prefix
	//      must be a DNS subdomain: a series of DNS labels separated by dots (.),
	//      not longer than 253 characters in total, followed by a slash (/).
	//      Metadata that fails to meet these requirements are rejected
	//
	//  Note: This field is equivalent to the `annotations` field in the v1 API.
	//  They have the same syntax and read/write to the same location in Service
	//  Directory.
	// +kcc:proto:field=google.cloud.servicedirectory.v1beta1.Endpoint.metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// Immutable. The Google Compute Engine network (VPC) of the endpoint in the
	//  format `projects/<project number>/locations/global/networks/*`.
	//
	//  The project must be specified by project number (project id is rejected).
	//  Incorrectly formatted networks are rejected, but no other validation
	//  is performed on this field (ex. network or project existence, reachability,
	//  or permissions).
	// +kcc:proto:field=google.cloud.servicedirectory.v1beta1.Endpoint.network
	Network *string `json:"network,omitempty"`
}

// +kcc:proto=google.cloud.servicedirectory.v1beta1.Endpoint
type EndpointObservedState struct {
	// Output only. The timestamp when the endpoint was created.
	// +kcc:proto:field=google.cloud.servicedirectory.v1beta1.Endpoint.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the endpoint was last updated.
	// +kcc:proto:field=google.cloud.servicedirectory.v1beta1.Endpoint.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. A globally unique identifier (in UUID4 format) for this
	//  endpoint.
	// +kcc:proto:field=google.cloud.servicedirectory.v1beta1.Endpoint.uid
	Uid *string `json:"uid,omitempty"`
}
