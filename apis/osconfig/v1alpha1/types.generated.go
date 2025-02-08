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


// +kcc:proto=google.cloud.osconfig.v1.CVSSv3
type CVSSv3 struct {
	// The base score is a function of the base metric scores.
	//  https://www.first.org/cvss/specification-document#Base-Metrics
	// +kcc:proto:field=google.cloud.osconfig.v1.CVSSv3.base_score
	BaseScore *float32 `json:"baseScore,omitempty"`

	// The Exploitability sub-score equation is derived from the Base
	//  Exploitability metrics.
	//  https://www.first.org/cvss/specification-document#2-1-Exploitability-Metrics
	// +kcc:proto:field=google.cloud.osconfig.v1.CVSSv3.exploitability_score
	ExploitabilityScore *float32 `json:"exploitabilityScore,omitempty"`

	// The Impact sub-score equation is derived from the Base Impact metrics.
	// +kcc:proto:field=google.cloud.osconfig.v1.CVSSv3.impact_score
	ImpactScore *float32 `json:"impactScore,omitempty"`

	// This metric reflects the context by which vulnerability exploitation is
	//  possible.
	// +kcc:proto:field=google.cloud.osconfig.v1.CVSSv3.attack_vector
	AttackVector *string `json:"attackVector,omitempty"`

	// This metric describes the conditions beyond the attacker's control that
	//  must exist in order to exploit the vulnerability.
	// +kcc:proto:field=google.cloud.osconfig.v1.CVSSv3.attack_complexity
	AttackComplexity *string `json:"attackComplexity,omitempty"`

	// This metric describes the level of privileges an attacker must possess
	//  before successfully exploiting the vulnerability.
	// +kcc:proto:field=google.cloud.osconfig.v1.CVSSv3.privileges_required
	PrivilegesRequired *string `json:"privilegesRequired,omitempty"`

	// This metric captures the requirement for a human user, other than the
	//  attacker, to participate in the successful compromise of the vulnerable
	//  component.
	// +kcc:proto:field=google.cloud.osconfig.v1.CVSSv3.user_interaction
	UserInteraction *string `json:"userInteraction,omitempty"`

	// The Scope metric captures whether a vulnerability in one vulnerable
	//  component impacts resources in components beyond its security scope.
	// +kcc:proto:field=google.cloud.osconfig.v1.CVSSv3.scope
	Scope *string `json:"scope,omitempty"`

	// This metric measures the impact to the confidentiality of the information
	//  resources managed by a software component due to a successfully exploited
	//  vulnerability.
	// +kcc:proto:field=google.cloud.osconfig.v1.CVSSv3.confidentiality_impact
	ConfidentialityImpact *string `json:"confidentialityImpact,omitempty"`

	// This metric measures the impact to integrity of a successfully exploited
	//  vulnerability.
	// +kcc:proto:field=google.cloud.osconfig.v1.CVSSv3.integrity_impact
	IntegrityImpact *string `json:"integrityImpact,omitempty"`

	// This metric measures the impact to the availability of the impacted
	//  component resulting from a successfully exploited vulnerability.
	// +kcc:proto:field=google.cloud.osconfig.v1.CVSSv3.availability_impact
	AvailabilityImpact *string `json:"availabilityImpact,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.VulnerabilityReport
type VulnerabilityReport struct {
}

// +kcc:proto=google.cloud.osconfig.v1.VulnerabilityReport.Vulnerability
type VulnerabilityReport_Vulnerability struct {
	// Contains metadata as per the upstream feed of the operating system and
	//  NVD.
	// +kcc:proto:field=google.cloud.osconfig.v1.VulnerabilityReport.Vulnerability.details
	Details *VulnerabilityReport_Vulnerability_Details `json:"details,omitempty"`

	// Corresponds to the `INSTALLED_PACKAGE` inventory item on the VM.
	//  This field displays the inventory items affected by this vulnerability.
	//  If the vulnerability report was not updated after the VM inventory
	//  update, these values might not display in VM inventory. For some distros,
	//  this field may be empty.
	// +kcc:proto:field=google.cloud.osconfig.v1.VulnerabilityReport.Vulnerability.installed_inventory_item_ids
	InstalledInventoryItemIds []string `json:"installedInventoryItemIds,omitempty"`

	// Corresponds to the `AVAILABLE_PACKAGE` inventory item on the VM.
	//  If the vulnerability report was not updated after the VM inventory
	//  update, these values might not display in VM inventory. If there is no
	//  available fix, the field is empty. The `inventory_item` value specifies
	//  the latest `SoftwarePackage` available to the VM that fixes the
	//  vulnerability.
	// +kcc:proto:field=google.cloud.osconfig.v1.VulnerabilityReport.Vulnerability.available_inventory_item_ids
	AvailableInventoryItemIds []string `json:"availableInventoryItemIds,omitempty"`

	// The timestamp for when the vulnerability was first detected.
	// +kcc:proto:field=google.cloud.osconfig.v1.VulnerabilityReport.Vulnerability.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// The timestamp for when the vulnerability was last modified.
	// +kcc:proto:field=google.cloud.osconfig.v1.VulnerabilityReport.Vulnerability.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// List of items affected by the vulnerability.
	// +kcc:proto:field=google.cloud.osconfig.v1.VulnerabilityReport.Vulnerability.items
	Items []VulnerabilityReport_Vulnerability_Item `json:"items,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.VulnerabilityReport.Vulnerability.Details
type VulnerabilityReport_Vulnerability_Details struct {
	// The CVE of the vulnerability. CVE cannot be
	//  empty and the combination of <cve, classification> should be unique
	//  across vulnerabilities for a VM.
	// +kcc:proto:field=google.cloud.osconfig.v1.VulnerabilityReport.Vulnerability.Details.cve
	Cve *string `json:"cve,omitempty"`

	// The CVSS V2 score of this vulnerability. CVSS V2 score is on a scale of
	//  0 - 10 where 0 indicates low severity and 10 indicates high severity.
	// +kcc:proto:field=google.cloud.osconfig.v1.VulnerabilityReport.Vulnerability.Details.cvss_v2_score
	CvssV2Score *float32 `json:"cvssV2Score,omitempty"`

	// The full description of the CVSSv3 for this vulnerability from NVD.
	// +kcc:proto:field=google.cloud.osconfig.v1.VulnerabilityReport.Vulnerability.Details.cvss_v3
	CvssV3 *CVSSv3 `json:"cvssV3,omitempty"`

	// Assigned severity/impact ranking from the distro.
	// +kcc:proto:field=google.cloud.osconfig.v1.VulnerabilityReport.Vulnerability.Details.severity
	Severity *string `json:"severity,omitempty"`

	// The note or description describing the vulnerability from the distro.
	// +kcc:proto:field=google.cloud.osconfig.v1.VulnerabilityReport.Vulnerability.Details.description
	Description *string `json:"description,omitempty"`

	// Corresponds to the references attached to the `VulnerabilityDetails`.
	// +kcc:proto:field=google.cloud.osconfig.v1.VulnerabilityReport.Vulnerability.Details.references
	References []VulnerabilityReport_Vulnerability_Details_Reference `json:"references,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.VulnerabilityReport.Vulnerability.Details.Reference
type VulnerabilityReport_Vulnerability_Details_Reference struct {
	// The url of the reference.
	// +kcc:proto:field=google.cloud.osconfig.v1.VulnerabilityReport.Vulnerability.Details.Reference.url
	URL *string `json:"url,omitempty"`

	// The source of the reference e.g. NVD.
	// +kcc:proto:field=google.cloud.osconfig.v1.VulnerabilityReport.Vulnerability.Details.Reference.source
	Source *string `json:"source,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.VulnerabilityReport.Vulnerability.Item
type VulnerabilityReport_Vulnerability_Item struct {
	// Corresponds to the `INSTALLED_PACKAGE` inventory item on the VM.
	//  This field displays the inventory items affected by this vulnerability.
	//  If the vulnerability report was not updated after the VM inventory
	//  update, these values might not display in VM inventory. For some
	//  operating systems, this field might be empty.
	// +kcc:proto:field=google.cloud.osconfig.v1.VulnerabilityReport.Vulnerability.Item.installed_inventory_item_id
	InstalledInventoryItemID *string `json:"installedInventoryItemID,omitempty"`

	// Corresponds to the `AVAILABLE_PACKAGE` inventory item on the VM.
	//  If the vulnerability report was not updated after the VM inventory
	//  update, these values might not display in VM inventory. If there is no
	//  available fix, the field is empty. The `inventory_item` value specifies
	//  the latest `SoftwarePackage` available to the VM that fixes the
	//  vulnerability.
	// +kcc:proto:field=google.cloud.osconfig.v1.VulnerabilityReport.Vulnerability.Item.available_inventory_item_id
	AvailableInventoryItemID *string `json:"availableInventoryItemID,omitempty"`

	// The recommended [CPE URI](https://cpe.mitre.org/specification/) update
	//  that contains a fix for this vulnerability.
	// +kcc:proto:field=google.cloud.osconfig.v1.VulnerabilityReport.Vulnerability.Item.fixed_cpe_uri
	FixedCpeURI *string `json:"fixedCpeURI,omitempty"`

	// The upstream OS patch, packages or KB that fixes the vulnerability.
	// +kcc:proto:field=google.cloud.osconfig.v1.VulnerabilityReport.Vulnerability.Item.upstream_fix
	UpstreamFix *string `json:"upstreamFix,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.VulnerabilityReport
type VulnerabilityReportObservedState struct {
	// Output only. The `vulnerabilityReport` API resource name.
	//
	//  Format:
	//  `projects/{project_number}/locations/{location}/instances/{instance_id}/vulnerabilityReport`
	// +kcc:proto:field=google.cloud.osconfig.v1.VulnerabilityReport.name
	Name *string `json:"name,omitempty"`

	// Output only. List of vulnerabilities affecting the VM.
	// +kcc:proto:field=google.cloud.osconfig.v1.VulnerabilityReport.vulnerabilities
	Vulnerabilities []VulnerabilityReport_Vulnerability `json:"vulnerabilities,omitempty"`

	// Output only. The timestamp for when the last vulnerability report was generated for the
	//  VM.
	// +kcc:proto:field=google.cloud.osconfig.v1.VulnerabilityReport.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
