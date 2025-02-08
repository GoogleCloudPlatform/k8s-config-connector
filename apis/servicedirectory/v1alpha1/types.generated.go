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


// +kcc:proto=google.cloud.servicedirectory.v1.Endpoint
type Endpoint struct {
	// Immutable. The resource name for the endpoint in the format
	//  `projects/*/locations/*/namespaces/*/services/*/endpoints/*`.
	// +kcc:proto:field=google.cloud.servicedirectory.v1.Endpoint.name
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
	// +kcc:proto:field=google.cloud.servicedirectory.v1.Endpoint.address
	Address *string `json:"address,omitempty"`

	// Optional. Service Directory rejects values outside of `[0, 65535]`.
	// +kcc:proto:field=google.cloud.servicedirectory.v1.Endpoint.port
	Port *int32 `json:"port,omitempty"`

	// Optional. Annotations for the endpoint. This data can be consumed by
	//  service clients.
	//
	//  Restrictions:
	//
	//  *   The entire annotations dictionary may contain up to 512 characters,
	//      spread accoss all key-value pairs. Annotations that go beyond this
	//      limit are rejected
	//  *   Valid annotation keys have two segments: an optional prefix and name,
	//      separated by a slash (/). The name segment is required and must be 63
	//      characters or less, beginning and ending with an alphanumeric character
	//      ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and
	//      alphanumerics between. The prefix is optional. If specified, the prefix
	//      must be a DNS subdomain: a series of DNS labels separated by dots (.),
	//      not longer than 253 characters in total, followed by a slash (/)
	//      Annotations that fails to meet these requirements are rejected.
	//
	//  Note: This field is equivalent to the `metadata` field in the v1beta1 API.
	//  They have the same syntax and read/write to the same location in Service
	//  Directory.
	// +kcc:proto:field=google.cloud.servicedirectory.v1.Endpoint.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Immutable. The Google Compute Engine network (VPC) of the endpoint in the
	//  format `projects/<project number>/locations/global/networks/*`.
	//
	//  The project must be specified by project number (project id is rejected).
	//  Incorrectly formatted networks are rejected, we also check to make sure
	//  that you have the servicedirectory.networks.attach permission on the
	//  project specified.
	// +kcc:proto:field=google.cloud.servicedirectory.v1.Endpoint.network
	Network *string `json:"network,omitempty"`
}

// +kcc:proto=google.cloud.servicedirectory.v1.Endpoint
type EndpointObservedState struct {
	// Output only. The globally unique identifier of the endpoint in the UUID4
	//  format.
	// +kcc:proto:field=google.cloud.servicedirectory.v1.Endpoint.uid
	Uid *string `json:"uid,omitempty"`
}
