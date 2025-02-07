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


// +kcc:proto=google.cloud.apphub.v1.Application
type Application struct {
	// Identifier. The resource name of an Application. Format:
	//  "projects/{host-project-id}/locations/{location}/applications/{application-id}"
	// +kcc:proto:field=google.cloud.apphub.v1.Application.name
	Name *string `json:"name,omitempty"`

	// Optional. User-defined name for the Application.
	//  Can have a maximum length of 63 characters.
	// +kcc:proto:field=google.cloud.apphub.v1.Application.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined description of an Application.
	//  Can have a maximum length of 2048 characters.
	// +kcc:proto:field=google.cloud.apphub.v1.Application.description
	Description *string `json:"description,omitempty"`

	// Optional. Consumer provided attributes.
	// +kcc:proto:field=google.cloud.apphub.v1.Application.attributes
	Attributes *Attributes `json:"attributes,omitempty"`

	// Required. Immutable. Defines what data can be included into this
	//  Application. Limits which Services and Workloads can be registered.
	// +kcc:proto:field=google.cloud.apphub.v1.Application.scope
	Scope *Scope `json:"scope,omitempty"`
}

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

// +kcc:proto=google.cloud.apphub.v1.Scope
type Scope struct {
	// Required. Scope Type.
	// +kcc:proto:field=google.cloud.apphub.v1.Scope.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.apphub.v1.Application
type ApplicationObservedState struct {
	// Output only. Create time.
	// +kcc:proto:field=google.cloud.apphub.v1.Application.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time.
	// +kcc:proto:field=google.cloud.apphub.v1.Application.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. A universally unique identifier (in UUID4 format) for the
	//  `Application`.
	// +kcc:proto:field=google.cloud.apphub.v1.Application.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Application state.
	// +kcc:proto:field=google.cloud.apphub.v1.Application.state
	State *string `json:"state,omitempty"`
}
