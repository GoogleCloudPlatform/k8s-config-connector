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


// +kcc:proto=google.cloud.discoveryengine.v1.Query
type Query struct {
	// Plain text.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Query.text
	Text *string `json:"text,omitempty"`

	// Unique Id for the query.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Query.query_id
	QueryID *string `json:"queryID,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.Session
type Session struct {
	// Immutable. Fully qualified name
	//  `projects/{project}/locations/global/collections/{collection}/engines/{engine}/sessions/*`
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Session.name
	Name *string `json:"name,omitempty"`

	// The state of the session.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Session.state
	State *string `json:"state,omitempty"`

	// A unique identifier for tracking users.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Session.user_pseudo_id
	UserPseudoID *string `json:"userPseudoID,omitempty"`

	// Turns.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Session.turns
	Turns []Session_Turn `json:"turns,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.Session.Turn
type Session_Turn struct {
	// The user query.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Session.Turn.query
	Query *Query `json:"query,omitempty"`

	// The resource name of the answer to the user query.
	//
	//  Only set if the answer generation (/answer API call) happened in this
	//  turn.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Session.Turn.answer
	Answer *string `json:"answer,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.Session
type SessionObservedState struct {
	// Output only. The time the session started.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Session.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. The time the session finished.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Session.end_time
	EndTime *string `json:"endTime,omitempty"`
}
