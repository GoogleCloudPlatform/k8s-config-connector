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


// +kcc:proto=google.cloud.servicehealth.v1.Event
type Event struct {

	// Google Cloud products and locations impacted by the event.
	// +kcc:proto:field=google.cloud.servicehealth.v1.Event.event_impacts
	EventImpacts []EventImpact `json:"eventImpacts,omitempty"`
}

// +kcc:proto=google.cloud.servicehealth.v1.EventImpact
type EventImpact struct {
	// Google Cloud product impacted by the event.
	// +kcc:proto:field=google.cloud.servicehealth.v1.EventImpact.product
	Product *Product `json:"product,omitempty"`

	// Location impacted by the event.
	// +kcc:proto:field=google.cloud.servicehealth.v1.EventImpact.location
	Location *Location `json:"location,omitempty"`
}

// +kcc:proto=google.cloud.servicehealth.v1.EventUpdate
type EventUpdate struct {
}

// +kcc:proto=google.cloud.servicehealth.v1.Location
type Location struct {
	// Location impacted by the event. Example: `"us-central1"`
	// +kcc:proto:field=google.cloud.servicehealth.v1.Location.location_name
	LocationName *string `json:"locationName,omitempty"`
}

// +kcc:proto=google.cloud.servicehealth.v1.Product
type Product struct {
	// Google Cloud product impacted by the event. Example: `"Google Cloud SQL"`
	// +kcc:proto:field=google.cloud.servicehealth.v1.Product.product_name
	ProductName *string `json:"productName,omitempty"`

	// Unique identifier for the product.
	// +kcc:proto:field=google.cloud.servicehealth.v1.Product.id
	ID *string `json:"id,omitempty"`
}

// +kcc:proto=google.cloud.servicehealth.v1.Event
type EventObservedState struct {
	// Output only. Identifier. Name of the event. Unique name of the event in
	//  this scope including project and location using the form
	//  `projects/{project_id}/locations/{location}/events/{event_id}`.
	// +kcc:proto:field=google.cloud.servicehealth.v1.Event.name
	Name *string `json:"name,omitempty"`

	// Output only. Brief description for the event.
	// +kcc:proto:field=google.cloud.servicehealth.v1.Event.title
	Title *string `json:"title,omitempty"`

	// Output only. Free-form, human-readable description.
	// +kcc:proto:field=google.cloud.servicehealth.v1.Event.description
	Description *string `json:"description,omitempty"`

	// Output only. The category of the event.
	// +kcc:proto:field=google.cloud.servicehealth.v1.Event.category
	Category *string `json:"category,omitempty"`

	// Output only. The detailed category of the event.
	// +kcc:proto:field=google.cloud.servicehealth.v1.Event.detailed_category
	DetailedCategory *string `json:"detailedCategory,omitempty"`

	// Output only. The current state of the event.
	// +kcc:proto:field=google.cloud.servicehealth.v1.Event.state
	State *string `json:"state,omitempty"`

	// Output only. The current detailed state of the incident.
	// +kcc:proto:field=google.cloud.servicehealth.v1.Event.detailed_state
	DetailedState *string `json:"detailedState,omitempty"`

	// Output only. Communicates why a given event is deemed relevant in the
	//  context of a given project.
	// +kcc:proto:field=google.cloud.servicehealth.v1.Event.relevance
	Relevance *string `json:"relevance,omitempty"`

	// Output only. Event updates are correspondence from Google.
	// +kcc:proto:field=google.cloud.servicehealth.v1.Event.updates
	Updates []EventUpdate `json:"updates,omitempty"`

	// Output only. When `detailed_state`=`MERGED`, `parent_event` contains the
	//  name of the parent event. All further updates will be published to the
	//  parent event.
	// +kcc:proto:field=google.cloud.servicehealth.v1.Event.parent_event
	ParentEvent *string `json:"parentEvent,omitempty"`

	// Output only. The time when the event was last modified.
	// +kcc:proto:field=google.cloud.servicehealth.v1.Event.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The start time of the event, if applicable.
	// +kcc:proto:field=google.cloud.servicehealth.v1.Event.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. The end time of the event, if applicable.
	// +kcc:proto:field=google.cloud.servicehealth.v1.Event.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. The time when the next update can be expected.
	// +kcc:proto:field=google.cloud.servicehealth.v1.Event.next_update_time
	NextUpdateTime *string `json:"nextUpdateTime,omitempty"`
}

// +kcc:proto=google.cloud.servicehealth.v1.EventUpdate
type EventUpdateObservedState struct {
	// Output only. The time the update was posted.
	// +kcc:proto:field=google.cloud.servicehealth.v1.EventUpdate.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Brief title for the event.
	// +kcc:proto:field=google.cloud.servicehealth.v1.EventUpdate.title
	Title *string `json:"title,omitempty"`

	// Output only. Free-form, human-readable description.
	// +kcc:proto:field=google.cloud.servicehealth.v1.EventUpdate.description
	Description *string `json:"description,omitempty"`

	// Output only. Symptoms of the event, if available.
	// +kcc:proto:field=google.cloud.servicehealth.v1.EventUpdate.symptom
	Symptom *string `json:"symptom,omitempty"`

	// Output only. Workaround steps to remediate the event impact, if available.
	// +kcc:proto:field=google.cloud.servicehealth.v1.EventUpdate.workaround
	Workaround *string `json:"workaround,omitempty"`
}
