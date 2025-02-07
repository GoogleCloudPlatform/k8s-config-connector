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


// +kcc:proto=google.cloud.apihub.v1.HostProjectRegistration
type HostProjectRegistration struct {
	// Identifier. The name of the host project registration.
	//  Format:
	//  "projects/{project}/locations/{location}/hostProjectRegistrations/{host_project_registration}".
	// +kcc:proto:field=google.cloud.apihub.v1.HostProjectRegistration.name
	Name *string `json:"name,omitempty"`

	// Required. Immutable. Google cloud project name in the format:
	//  "projects/abc" or "projects/123". As input, project name with either
	//  project id or number are accepted. As output, this field will contain
	//  project number.
	// +kcc:proto:field=google.cloud.apihub.v1.HostProjectRegistration.gcp_project
	GcpProject *string `json:"gcpProject,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.HostProjectRegistration
type HostProjectRegistrationObservedState struct {
	// Output only. The time at which the host project registration was created.
	// +kcc:proto:field=google.cloud.apihub.v1.HostProjectRegistration.create_time
	CreateTime *string `json:"createTime,omitempty"`
}
