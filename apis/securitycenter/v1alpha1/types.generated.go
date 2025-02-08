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


// +kcc:proto=google.cloud.securitycenter.v2.ExternalSystem
type ExternalSystem struct {
	// Full resource name of the external system. The following list
	//  shows some examples:
	//
	//  + `organizations/1234/sources/5678/findings/123456/externalSystems/jira`
	//  +
	//  `organizations/1234/sources/5678/locations/us/findings/123456/externalSystems/jira`
	//  + `folders/1234/sources/5678/findings/123456/externalSystems/jira`
	//  +
	//  `folders/1234/sources/5678/locations/us/findings/123456/externalSystems/jira`
	//  + `projects/1234/sources/5678/findings/123456/externalSystems/jira`
	//  +
	//  `projects/1234/sources/5678/locations/us/findings/123456/externalSystems/jira`
	// +kcc:proto:field=google.cloud.securitycenter.v2.ExternalSystem.name
	Name *string `json:"name,omitempty"`

	// References primary/secondary etc assignees in the external system.
	// +kcc:proto:field=google.cloud.securitycenter.v2.ExternalSystem.assignees
	Assignees []string `json:"assignees,omitempty"`

	// The identifier that's used to track the finding's corresponding case in the
	//  external system.
	// +kcc:proto:field=google.cloud.securitycenter.v2.ExternalSystem.external_uid
	ExternalUid *string `json:"externalUid,omitempty"`

	// The most recent status of the finding's corresponding case, as reported by
	//  the external system.
	// +kcc:proto:field=google.cloud.securitycenter.v2.ExternalSystem.status
	Status *string `json:"status,omitempty"`

	// The time when the case was last updated, as reported by the external
	//  system.
	// +kcc:proto:field=google.cloud.securitycenter.v2.ExternalSystem.external_system_update_time
	ExternalSystemUpdateTime *string `json:"externalSystemUpdateTime,omitempty"`

	// The link to the finding's corresponding case in the external system.
	// +kcc:proto:field=google.cloud.securitycenter.v2.ExternalSystem.case_uri
	CaseURI *string `json:"caseURI,omitempty"`

	// The priority of the finding's corresponding case in the external system.
	// +kcc:proto:field=google.cloud.securitycenter.v2.ExternalSystem.case_priority
	CasePriority *string `json:"casePriority,omitempty"`

	// The SLA of the finding's corresponding case in the external system.
	// +kcc:proto:field=google.cloud.securitycenter.v2.ExternalSystem.case_sla
	CaseSla *string `json:"caseSla,omitempty"`

	// The time when the case was created, as reported by the external system.
	// +kcc:proto:field=google.cloud.securitycenter.v2.ExternalSystem.case_create_time
	CaseCreateTime *string `json:"caseCreateTime,omitempty"`

	// The time when the case was closed, as reported by the external system.
	// +kcc:proto:field=google.cloud.securitycenter.v2.ExternalSystem.case_close_time
	CaseCloseTime *string `json:"caseCloseTime,omitempty"`

	// Information about the ticket, if any, that is being used to track the
	//  resolution of the issue that is identified by this finding.
	// +kcc:proto:field=google.cloud.securitycenter.v2.ExternalSystem.ticket_info
	TicketInfo *ExternalSystem_TicketInfo `json:"ticketInfo,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v2.ExternalSystem.TicketInfo
type ExternalSystem_TicketInfo struct {
	// The identifier of the ticket in the ticket system.
	// +kcc:proto:field=google.cloud.securitycenter.v2.ExternalSystem.TicketInfo.id
	ID *string `json:"id,omitempty"`

	// The assignee of the ticket in the ticket system.
	// +kcc:proto:field=google.cloud.securitycenter.v2.ExternalSystem.TicketInfo.assignee
	Assignee *string `json:"assignee,omitempty"`

	// The description of the ticket in the ticket system.
	// +kcc:proto:field=google.cloud.securitycenter.v2.ExternalSystem.TicketInfo.description
	Description *string `json:"description,omitempty"`

	// The link to the ticket in the ticket system.
	// +kcc:proto:field=google.cloud.securitycenter.v2.ExternalSystem.TicketInfo.uri
	URI *string `json:"uri,omitempty"`

	// The latest status of the ticket, as reported by the ticket system.
	// +kcc:proto:field=google.cloud.securitycenter.v2.ExternalSystem.TicketInfo.status
	Status *string `json:"status,omitempty"`

	// The time when the ticket was last updated, as reported by the ticket
	//  system.
	// +kcc:proto:field=google.cloud.securitycenter.v2.ExternalSystem.TicketInfo.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
