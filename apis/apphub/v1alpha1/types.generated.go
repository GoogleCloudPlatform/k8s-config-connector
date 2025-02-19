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

// +kcc:proto=google.cloud.apphub.v1.Attributes
type Attributes struct {
	// Optional. User-defined criticality information.
	// +kcc:proto:field=google.cloud.apphub.v1.Attributes.criticality
	Criticality *Criticality `json:"criticality,omitempty"`

	// Optional. User-defined environment information.
	// +kcc:proto:field=google.cloud.apphub.v1.Attributes.environment
	Environment *Environment `json:"environment,omitempty"`

	// Optional. Developer team that owns development and coding.
	// +kcc:proto:field=google.cloud.apphub.v1.Attributes.developer_owners
	DeveloperOwners []ContactInfo `json:"developerOwners,omitempty"`

	// Optional. Operator team that ensures runtime and operations.
	// +kcc:proto:field=google.cloud.apphub.v1.Attributes.operator_owners
	OperatorOwners []ContactInfo `json:"operatorOwners,omitempty"`

	// Optional. Business team that ensures user needs are met and value is
	//  delivered
	// +kcc:proto:field=google.cloud.apphub.v1.Attributes.business_owners
	BusinessOwners []ContactInfo `json:"businessOwners,omitempty"`
}

// +kcc:proto=google.cloud.apphub.v1.ContactInfo
type ContactInfo struct {
	// Optional. Contact's name.
	//  Can have a maximum length of 63 characters.
	// +kcc:proto:field=google.cloud.apphub.v1.ContactInfo.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Email address of the contacts.
	// +kcc:proto:field=google.cloud.apphub.v1.ContactInfo.email
	Email *string `json:"email,omitempty"`
}

// +kcc:proto=google.cloud.apphub.v1.Criticality
type Criticality struct {
	// Required. Criticality Type.
	// +kcc:proto:field=google.cloud.apphub.v1.Criticality.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.apphub.v1.Environment
type Environment struct {
	// Required. Environment Type.
	// +kcc:proto:field=google.cloud.apphub.v1.Environment.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.apphub.v1.Workload
type Workload struct {
	// Identifier. The resource name of the Workload. Format:
	//  "projects/{host-project-id}/locations/{location}/applications/{application-id}/workloads/{workload-id}"
	// +kcc:proto:field=google.cloud.apphub.v1.Workload.name
	Name *string `json:"name,omitempty"`

	// Optional. User-defined name for the Workload.
	//  Can have a maximum length of 63 characters.
	// +kcc:proto:field=google.cloud.apphub.v1.Workload.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined description of a Workload.
	//  Can have a maximum length of 2048 characters.
	// +kcc:proto:field=google.cloud.apphub.v1.Workload.description
	Description *string `json:"description,omitempty"`

	// Required. Immutable. The resource name of the original discovered workload.
	// +kcc:proto:field=google.cloud.apphub.v1.Workload.discovered_workload
	DiscoveredWorkload *string `json:"discoveredWorkload,omitempty"`

	// Optional. Consumer provided attributes.
	// +kcc:proto:field=google.cloud.apphub.v1.Workload.attributes
	Attributes *Attributes `json:"attributes,omitempty"`
}

// +kcc:proto=google.cloud.apphub.v1.WorkloadProperties
type WorkloadProperties struct {
}

// +kcc:proto=google.cloud.apphub.v1.WorkloadReference
type WorkloadReference struct {
}

// +kcc:proto=google.cloud.apphub.v1.Workload
type WorkloadObservedState struct {
	// Output only. Reference of an underlying compute resource represented by the
	//  Workload. These are immutable.
	// +kcc:proto:field=google.cloud.apphub.v1.Workload.workload_reference
	WorkloadReference *WorkloadReference `json:"workloadReference,omitempty"`

	// Output only. Properties of an underlying compute resource represented by
	//  the Workload. These are immutable.
	// +kcc:proto:field=google.cloud.apphub.v1.Workload.workload_properties
	WorkloadProperties *WorkloadProperties `json:"workloadProperties,omitempty"`

	// Output only. Create time.
	// +kcc:proto:field=google.cloud.apphub.v1.Workload.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time.
	// +kcc:proto:field=google.cloud.apphub.v1.Workload.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. A universally unique identifier (UUID) for the `Workload` in
	//  the UUID4 format.
	// +kcc:proto:field=google.cloud.apphub.v1.Workload.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Workload state.
	// +kcc:proto:field=google.cloud.apphub.v1.Workload.state
	State *string `json:"state,omitempty"`
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
