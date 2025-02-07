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


// +kcc:proto=google.cloud.bigquery.analyticshub.v1.Subscription
type Subscription struct {
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.Subscription.LinkedResource
type Subscription_LinkedResource struct {
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.Subscription
type SubscriptionObservedState struct {
	// Output only. Resource name of the source Listing.
	//  e.g. projects/123/locations/US/dataExchanges/456/listings/789
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Subscription.listing
	Listing *string `json:"listing,omitempty"`

	// Output only. Resource name of the source Data Exchange.
	//  e.g. projects/123/locations/US/dataExchanges/456
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Subscription.data_exchange
	DataExchange *string `json:"dataExchange,omitempty"`

	// Output only. The resource name of the subscription.
	//  e.g. `projects/myproject/locations/US/subscriptions/123`.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Subscription.name
	Name *string `json:"name,omitempty"`

	// Output only. Timestamp when the subscription was created.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Subscription.creation_time
	CreationTime *string `json:"creationTime,omitempty"`

	// Output only. Timestamp when the subscription was last modified.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Subscription.last_modify_time
	LastModifyTime *string `json:"lastModifyTime,omitempty"`

	// Output only. Organization of the project this subscription belongs to.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Subscription.organization_id
	OrganizationID *string `json:"organizationID,omitempty"`

	// Output only. Display name of the project of this subscription.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Subscription.organization_display_name
	OrganizationDisplayName *string `json:"organizationDisplayName,omitempty"`

	// Output only. Current state of the subscription.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Subscription.state
	State *string `json:"state,omitempty"`

	// TODO: unsupported map type with key string and value message


	// Output only. Email of the subscriber.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.Subscription.subscriber_contact
	SubscriberContact *string `json:"subscriberContact,omitempty"`
}
