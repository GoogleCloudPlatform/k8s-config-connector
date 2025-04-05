// Copyright 2024 Google LLC
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

package v1beta1

// +kcc:proto=google.cloud.vpcaccess.v1.Connector.Subnet
type Connector_Subnet struct {
	// Subnet name (relative, not fully qualified).
	//  E.g. if the full subnet selfLink is
	//  https://compute.googleapis.com/compute/v1/projects/{project}/regions/{region}/subnetworks/{subnetName}
	//  the correct input for this field would be {subnetName}
	Name *string `json:"name,omitempty"`

	// Project in which the subnet exists.
	//  If not set, this project is assumed to be the project for which
	//  the connector create request was issued.
	ProjectID *string `json:"projectID,omitempty"`
}
