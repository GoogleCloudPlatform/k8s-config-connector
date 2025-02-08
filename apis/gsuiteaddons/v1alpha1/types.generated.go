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


// +kcc:proto=google.cloud.gsuiteaddons.v1.InstallStatus
type InstallStatus struct {
	// The canonical full resource name of the deployment install status.
	//
	//  Example:  `projects/123/deployments/my_deployment/installStatus`.
	// +kcc:proto:field=google.cloud.gsuiteaddons.v1.InstallStatus.name
	Name *string `json:"name,omitempty"`

	// True if the deployment is installed for the user
	// +kcc:proto:field=google.cloud.gsuiteaddons.v1.InstallStatus.installed
	Installed *bool `json:"installed,omitempty"`
}
