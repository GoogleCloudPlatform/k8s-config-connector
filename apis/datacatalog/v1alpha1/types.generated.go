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


// +kcc:proto=google.cloud.datacatalog.lineage.v1.EntityReference
type EntityReference struct {
	// Required. [Fully Qualified Name
	//  (FQN)](https://cloud.google.com/data-catalog/docs/fully-qualified-names)
	//  of the entity.
	// +kcc:proto:field=google.cloud.datacatalog.lineage.v1.EntityReference.fully_qualified_name
	FullyQualifiedName *string `json:"fullyQualifiedName,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.lineage.v1.EventLink
type EventLink struct {
	// Required. Reference to the source entity
	// +kcc:proto:field=google.cloud.datacatalog.lineage.v1.EventLink.source
	Source *EntityReference `json:"source,omitempty"`

	// Required. Reference to the target entity
	// +kcc:proto:field=google.cloud.datacatalog.lineage.v1.EventLink.target
	Target *EntityReference `json:"target,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.lineage.v1.LineageEvent
type LineageEvent struct {
	// Immutable. The resource name of the lineage event.
	//  Format:
	//  `projects/{project}/locations/{location}/processes/{process}/runs/{run}/lineageEvents/{lineage_event}`.
	//  Can be specified or auto-assigned.
	//  {lineage_event} must be not longer than 200 characters and only
	//  contain characters in a set: `a-zA-Z0-9_-:.`
	// +kcc:proto:field=google.cloud.datacatalog.lineage.v1.LineageEvent.name
	Name *string `json:"name,omitempty"`

	// Optional. List of source-target pairs. Can't contain more than 100 tuples.
	// +kcc:proto:field=google.cloud.datacatalog.lineage.v1.LineageEvent.links
	Links []EventLink `json:"links,omitempty"`

	// Required. The beginning of the transformation which resulted in this
	//  lineage event. For streaming scenarios, it should be the beginning of the
	//  period from which the lineage is being reported.
	// +kcc:proto:field=google.cloud.datacatalog.lineage.v1.LineageEvent.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Optional. The end of the transformation which resulted in this lineage
	//  event.  For streaming scenarios, it should be the end of the period from
	//  which the lineage is being reported.
	// +kcc:proto:field=google.cloud.datacatalog.lineage.v1.LineageEvent.end_time
	EndTime *string `json:"endTime,omitempty"`
}
