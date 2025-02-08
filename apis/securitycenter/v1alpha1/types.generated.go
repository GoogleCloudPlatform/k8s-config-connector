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


// +kcc:proto=google.cloud.securitycenter.v2.AttackPath
type AttackPath struct {
	// The attack path name, for example,
	//   `organizations/12/simulations/34/valuedResources/56/attackPaths/78`
	// +kcc:proto:field=google.cloud.securitycenter.v2.AttackPath.name
	Name *string `json:"name,omitempty"`

	// A list of nodes that exist in this attack path.
	// +kcc:proto:field=google.cloud.securitycenter.v2.AttackPath.path_nodes
	PathNodes []AttackPath_AttackPathNode `json:"pathNodes,omitempty"`

	// A list of the edges between nodes in this attack path.
	// +kcc:proto:field=google.cloud.securitycenter.v2.AttackPath.edges
	Edges []AttackPath_AttackPathEdge `json:"edges,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v2.AttackPath.AttackPathEdge
type AttackPath_AttackPathEdge struct {
	// The attack node uuid of the source node.
	// +kcc:proto:field=google.cloud.securitycenter.v2.AttackPath.AttackPathEdge.source
	Source *string `json:"source,omitempty"`

	// The attack node uuid of the destination node.
	// +kcc:proto:field=google.cloud.securitycenter.v2.AttackPath.AttackPathEdge.destination
	Destination *string `json:"destination,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v2.AttackPath.AttackPathNode
type AttackPath_AttackPathNode struct {
	// The name of the resource at this point in the attack path.
	//  The format of the name follows the Cloud Asset Inventory [resource
	//  name
	//  format](https://cloud.google.com/asset-inventory/docs/resource-name-format)
	// +kcc:proto:field=google.cloud.securitycenter.v2.AttackPath.AttackPathNode.resource
	Resource *string `json:"resource,omitempty"`

	// The [supported resource
	//  type](https://cloud.google.com/asset-inventory/docs/supported-asset-types)
	// +kcc:proto:field=google.cloud.securitycenter.v2.AttackPath.AttackPathNode.resource_type
	ResourceType *string `json:"resourceType,omitempty"`

	// Human-readable name of this resource.
	// +kcc:proto:field=google.cloud.securitycenter.v2.AttackPath.AttackPathNode.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The findings associated with this node in the attack path.
	// +kcc:proto:field=google.cloud.securitycenter.v2.AttackPath.AttackPathNode.associated_findings
	AssociatedFindings []AttackPath_AttackPathNode_PathNodeAssociatedFinding `json:"associatedFindings,omitempty"`

	// Unique id of the attack path node.
	// +kcc:proto:field=google.cloud.securitycenter.v2.AttackPath.AttackPathNode.uuid
	Uuid *string `json:"uuid,omitempty"`

	// A list of attack step nodes that exist in this attack path node.
	// +kcc:proto:field=google.cloud.securitycenter.v2.AttackPath.AttackPathNode.attack_steps
	AttackSteps []AttackPath_AttackPathNode_AttackStepNode `json:"attackSteps,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v2.AttackPath.AttackPathNode.AttackStepNode
type AttackPath_AttackPathNode_AttackStepNode struct {
	// Unique ID for one Node
	// +kcc:proto:field=google.cloud.securitycenter.v2.AttackPath.AttackPathNode.AttackStepNode.uuid
	Uuid *string `json:"uuid,omitempty"`

	// Attack step type. Can be either AND, OR or DEFENSE
	// +kcc:proto:field=google.cloud.securitycenter.v2.AttackPath.AttackPathNode.AttackStepNode.type
	Type *string `json:"type,omitempty"`

	// User friendly name of the attack step
	// +kcc:proto:field=google.cloud.securitycenter.v2.AttackPath.AttackPathNode.AttackStepNode.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Attack step labels for metadata
	// +kcc:proto:field=google.cloud.securitycenter.v2.AttackPath.AttackPathNode.AttackStepNode.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Attack step description
	// +kcc:proto:field=google.cloud.securitycenter.v2.AttackPath.AttackPathNode.AttackStepNode.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v2.AttackPath.AttackPathNode.PathNodeAssociatedFinding
type AttackPath_AttackPathNode_PathNodeAssociatedFinding struct {
	// Canonical name of the associated findings. Example:
	//  `organizations/123/sources/456/findings/789`
	// +kcc:proto:field=google.cloud.securitycenter.v2.AttackPath.AttackPathNode.PathNodeAssociatedFinding.canonical_finding
	CanonicalFinding *string `json:"canonicalFinding,omitempty"`

	// The additional taxonomy group within findings from a given source.
	// +kcc:proto:field=google.cloud.securitycenter.v2.AttackPath.AttackPathNode.PathNodeAssociatedFinding.finding_category
	FindingCategory *string `json:"findingCategory,omitempty"`

	// Full resource name of the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v2.AttackPath.AttackPathNode.PathNodeAssociatedFinding.name
	Name *string `json:"name,omitempty"`
}
