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

package v1alpha1

// +kcc:proto=google.cloud.workstations.v1.WorkstationCluster.PrivateClusterConfig
type WorkstationCluster_PrivateClusterConfig struct {
	// Immutable. Whether Workstations endpoint is private.
	EnablePrivateEndpoint *bool `json:"enablePrivateEndpoint,omitempty"`

	// NOTYET: This is an output field in a message that is mostly spec
	// // Output only. Hostname for the workstation cluster. This field will be
	// //  populated only when private endpoint is enabled. To access workstations
	// //  in the workstation cluster, create a new DNS zone mapping this domain
	// //  name to an internal IP address and a forwarding rule mapping that address
	// //  to the service attachment.
	// ClusterHostname *string `json:"clusterHostname,omitempty"`

	// NOTYET: This is an output field in a message that is mostly spec
	// // Output only. Service attachment URI for the workstation cluster. The
	// //  service attachment is created when private endpoint is enabled. To access
	// //  workstations in the workstation cluster, configure access to the managed
	// //  service using [Private Service
	// //  Connect](https://cloud.google.com/vpc/docs/configure-private-service-connect-services).
	// ServiceAttachmentUri *string `json:"serviceAttachmentUri,omitempty"`

	// Optional. Additional projects that are allowed to attach to the
	//  workstation cluster's service attachment. By default, the workstation
	//  cluster's project and the VPC host project (if different) are allowed.
	AllowedProjects []string `json:"allowedProjects,omitempty"`
}
