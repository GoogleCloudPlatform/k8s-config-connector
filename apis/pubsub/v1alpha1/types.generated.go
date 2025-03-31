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

// +kcc:proto=google.pubsub.v1.Snapshot
type Snapshot struct {
	// Optional. The name of the snapshot.
	// +kcc:proto:field=google.pubsub.v1.Snapshot.name
	Name *string `json:"name,omitempty"`

	// Optional. The name of the topic from which this snapshot is retaining
	//  messages.
	// +kcc:proto:field=google.pubsub.v1.Snapshot.topic
	Topic *string `json:"topic,omitempty"`

	// Optional. The snapshot is guaranteed to exist up until this time.
	//  A newly-created snapshot expires no later than 7 days from the time of its
	//  creation. Its exact lifetime is determined at creation by the existing
	//  backlog in the source subscription. Specifically, the lifetime of the
	//  snapshot is `7 days - (age of oldest unacked message in the subscription)`.
	//  For example, consider a subscription whose oldest unacked message is 3 days
	//  old. If a snapshot is created from this subscription, the snapshot -- which
	//  will always capture this 3-day-old backlog as long as the snapshot
	//  exists -- will expire in 4 days. The service will refuse to create a
	//  snapshot that would expire in less than 1 hour after creation.
	// +kcc:proto:field=google.pubsub.v1.Snapshot.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Optional. See [Creating and managing labels]
	//  (https://cloud.google.com/pubsub/docs/labels).
	// +kcc:proto:field=google.pubsub.v1.Snapshot.labels
	Labels map[string]string `json:"labels,omitempty"`
}
