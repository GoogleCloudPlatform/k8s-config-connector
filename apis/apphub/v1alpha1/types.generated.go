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


// +kcc:proto=google.cloud.apphub.v1.DiscoveredService
type DiscoveredService struct {
	// Identifier. The resource name of the discovered service. Format:
	//  "projects/{host-project-id}/locations/{location}/discoveredServices/{uuid}""
	// +kcc:proto:field=google.cloud.apphub.v1.DiscoveredService.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.apphub.v1.ServiceProperties
type ServiceProperties struct {
}

// +kcc:proto=google.cloud.apphub.v1.ServiceReference
type ServiceReference struct {
}

// +kcc:proto=google.cloud.apphub.v1.DiscoveredService
type DiscoveredServiceObservedState struct {
	// Output only. Reference to an underlying networking resource that can
	//  comprise a Service. These are immutable.
	// +kcc:proto:field=google.cloud.apphub.v1.DiscoveredService.service_reference
	ServiceReference *ServiceReference `json:"serviceReference,omitempty"`

	// Output only. Properties of an underlying compute resource that can comprise
	//  a Service. These are immutable.
	// +kcc:proto:field=google.cloud.apphub.v1.DiscoveredService.service_properties
	ServiceProperties *ServiceProperties `json:"serviceProperties,omitempty"`
}

// +kcc:proto=google.cloud.apphub.v1.ServiceProperties
type ServicePropertiesObservedState struct {
	// Output only. The service project identifier that the underlying cloud
	//  resource resides in.
	// +kcc:proto:field=google.cloud.apphub.v1.ServiceProperties.gcp_project
	GcpProject *string `json:"gcpProject,omitempty"`

	// Output only. The location that the underlying resource resides in, for
	//  example, us-west1.
	// +kcc:proto:field=google.cloud.apphub.v1.ServiceProperties.location
	Location *string `json:"location,omitempty"`

	// Output only. The location that the underlying resource resides in if it is
	//  zonal, for example, us-west1-a).
	// +kcc:proto:field=google.cloud.apphub.v1.ServiceProperties.zone
	Zone *string `json:"zone,omitempty"`
}

// +kcc:proto=google.cloud.apphub.v1.ServiceReference
type ServiceReferenceObservedState struct {
	// Output only. The underlying resource URI (For example, URI of Forwarding
	//  Rule, URL Map, and Backend Service).
	// +kcc:proto:field=google.cloud.apphub.v1.ServiceReference.uri
	URI *string `json:"uri,omitempty"`
}
