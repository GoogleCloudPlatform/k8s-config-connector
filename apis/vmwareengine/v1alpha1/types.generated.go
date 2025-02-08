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


// +kcc:proto=google.cloud.vmwareengine.v1.DnsBindPermission
type DnsBindPermission struct {
}

// +kcc:proto=google.cloud.vmwareengine.v1.Principal
type Principal struct {
	// The user who needs to be granted permission.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Principal.user
	User *string `json:"user,omitempty"`

	// The service account which needs to be granted the permission.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Principal.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.DnsBindPermission
type DnsBindPermissionObservedState struct {
	// Required. Output only. The name of the resource which stores the
	//  users/service accounts having the permission to bind to the corresponding
	//  intranet VPC of the consumer project. DnsBindPermission is a global
	//  resource and location can only be global. Resource names are schemeless
	//  URIs that follow the conventions in
	//  https://cloud.google.com/apis/design/resource_names. For example:
	//  `projects/my-project/locations/global/dnsBindPermission`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.DnsBindPermission.name
	Name *string `json:"name,omitempty"`

	// Output only. Users/Service accounts which have access for binding on the
	//  intranet VPC project corresponding to the consumer project.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.DnsBindPermission.principals
	Principals []Principal `json:"principals,omitempty"`
}
