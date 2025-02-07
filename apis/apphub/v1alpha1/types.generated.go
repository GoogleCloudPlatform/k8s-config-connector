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


// +kcc:proto=google.cloud.apphub.v1.DiscoveredWorkload
type DiscoveredWorkload struct {
	// Identifier. The resource name of the discovered workload. Format:
	//  "projects/{host-project-id}/locations/{location}/discoveredWorkloads/{uuid}"
	// +kcc:proto:field=google.cloud.apphub.v1.DiscoveredWorkload.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.apphub.v1.WorkloadProperties
type WorkloadProperties struct {
}

// +kcc:proto=google.cloud.apphub.v1.WorkloadReference
type WorkloadReference struct {
}

// +kcc:proto=google.cloud.apphub.v1.DiscoveredWorkload
type DiscoveredWorkloadObservedState struct {
	// Output only. Reference of an underlying compute resource represented by the
	//  Workload. These are immutable.
	// +kcc:proto:field=google.cloud.apphub.v1.DiscoveredWorkload.workload_reference
	WorkloadReference *WorkloadReference `json:"workloadReference,omitempty"`

	// Output only. Properties of an underlying compute resource represented by
	//  the Workload. These are immutable.
	// +kcc:proto:field=google.cloud.apphub.v1.DiscoveredWorkload.workload_properties
	WorkloadProperties *WorkloadProperties `json:"workloadProperties,omitempty"`
}

// +kcc:proto=google.cloud.apphub.v1.WorkloadProperties
type WorkloadPropertiesObservedState struct {
	// Output only. The service project identifier that the underlying cloud
	//  resource resides in. Empty for non cloud resources.
	// +kcc:proto:field=google.cloud.apphub.v1.WorkloadProperties.gcp_project
	GcpProject *string `json:"gcpProject,omitempty"`

	// Output only. The location that the underlying compute resource resides in
	//  (e.g us-west1).
	// +kcc:proto:field=google.cloud.apphub.v1.WorkloadProperties.location
	Location *string `json:"location,omitempty"`

	// Output only. The location that the underlying compute resource resides in
	//  if it is zonal (e.g us-west1-a).
	// +kcc:proto:field=google.cloud.apphub.v1.WorkloadProperties.zone
	Zone *string `json:"zone,omitempty"`
}

// +kcc:proto=google.cloud.apphub.v1.WorkloadReference
type WorkloadReferenceObservedState struct {
	// Output only. The underlying compute resource uri.
	// +kcc:proto:field=google.cloud.apphub.v1.WorkloadReference.uri
	URI *string `json:"uri,omitempty"`
}
