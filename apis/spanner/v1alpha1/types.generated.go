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


// +kcc:proto=google.spanner.v1.Session
type Session struct {

	// The labels for the session.
	//
	//   * Label keys must be between 1 and 63 characters long and must conform to
	//     the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?`.
	//   * Label values must be between 0 and 63 characters long and must conform
	//     to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`.
	//   * No more than 64 labels can be associated with a given session.
	//
	//  See https://goo.gl/xmQnxf for more information on and examples of labels.
	// +kcc:proto:field=google.spanner.v1.Session.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The database role which created this session.
	// +kcc:proto:field=google.spanner.v1.Session.creator_role
	CreatorRole *string `json:"creatorRole,omitempty"`

	// Optional. If true, specifies a multiplexed session. A multiplexed session
	//  may be used for multiple, concurrent read-only operations but can not be
	//  used for read-write transactions, partitioned reads, or partitioned
	//  queries. Multiplexed sessions can be created via
	//  [CreateSession][google.spanner.v1.Spanner.CreateSession] but not via
	//  [BatchCreateSessions][google.spanner.v1.Spanner.BatchCreateSessions].
	//  Multiplexed sessions may not be deleted nor listed.
	// +kcc:proto:field=google.spanner.v1.Session.multiplexed
	Multiplexed *bool `json:"multiplexed,omitempty"`
}

// +kcc:proto=google.spanner.v1.Session
type SessionObservedState struct {
	// Output only. The name of the session. This is always system-assigned.
	// +kcc:proto:field=google.spanner.v1.Session.name
	Name *string `json:"name,omitempty"`

	// Output only. The timestamp when the session is created.
	// +kcc:proto:field=google.spanner.v1.Session.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The approximate timestamp when the session is last used. It is
	//  typically earlier than the actual last use time.
	// +kcc:proto:field=google.spanner.v1.Session.approximate_last_use_time
	ApproximateLastUseTime *string `json:"approximateLastUseTime,omitempty"`
}
