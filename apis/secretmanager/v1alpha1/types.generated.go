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


// +kcc:proto=google.cloud.secretmanager.v1.Topic
type Topic struct {
	// Required. The resource name of the Pub/Sub topic that will be published to,
	//  in the following format: `projects/*/topics/*`. For publication to succeed,
	//  the Secret Manager service agent must have the `pubsub.topic.publish`
	//  permission on the topic. The Pub/Sub Publisher role
	//  (`roles/pubsub.publisher`) includes this permission.
	// +kcc:proto:field=google.cloud.secretmanager.v1.Topic.name
	Name *string `json:"name,omitempty"`
}
