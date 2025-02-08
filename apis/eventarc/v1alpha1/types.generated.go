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


// +kcc:proto=google.cloud.eventarc.v1.EventType
type EventType struct {
}

// +kcc:proto=google.cloud.eventarc.v1.FilteringAttribute
type FilteringAttribute struct {
}

// +kcc:proto=google.cloud.eventarc.v1.Provider
type Provider struct {
}

// +kcc:proto=google.cloud.eventarc.v1.EventType
type EventTypeObservedState struct {
	// Output only. The full name of the event type (for example,
	//  "google.cloud.storage.object.v1.finalized"). In the form of
	//  {provider-specific-prefix}.{resource}.{version}.{verb}. Types MUST be
	//  versioned and event schemas are guaranteed to remain backward compatible
	//  within one version. Note that event type versions and API versions do not
	//  need to match.
	// +kcc:proto:field=google.cloud.eventarc.v1.EventType.type
	Type *string `json:"type,omitempty"`

	// Output only. Human friendly description of what the event type is about.
	//  For example "Bucket created in Cloud Storage".
	// +kcc:proto:field=google.cloud.eventarc.v1.EventType.description
	Description *string `json:"description,omitempty"`

	// Output only. Filtering attributes for the event type.
	// +kcc:proto:field=google.cloud.eventarc.v1.EventType.filtering_attributes
	FilteringAttributes []FilteringAttribute `json:"filteringAttributes,omitempty"`

	// Output only. URI for the event schema.
	//  For example
	//  "https://github.com/googleapis/google-cloudevents/blob/master/proto/google/events/cloud/storage/v1/events.proto"
	// +kcc:proto:field=google.cloud.eventarc.v1.EventType.event_schema_uri
	EventSchemaURI *string `json:"eventSchemaURI,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.FilteringAttribute
type FilteringAttributeObservedState struct {
	// Output only. Attribute used for filtering the event type.
	// +kcc:proto:field=google.cloud.eventarc.v1.FilteringAttribute.attribute
	Attribute *string `json:"attribute,omitempty"`

	// Output only. Description of the purpose of the attribute.
	// +kcc:proto:field=google.cloud.eventarc.v1.FilteringAttribute.description
	Description *string `json:"description,omitempty"`

	// Output only. If true, the triggers for this provider should always specify
	//  a filter on these attributes. Trigger creation will fail otherwise.
	// +kcc:proto:field=google.cloud.eventarc.v1.FilteringAttribute.required
	Required *bool `json:"required,omitempty"`

	// Output only. If true, the attribute accepts matching expressions in the
	//  Eventarc PathPattern format.
	// +kcc:proto:field=google.cloud.eventarc.v1.FilteringAttribute.path_pattern_supported
	PathPatternSupported *bool `json:"pathPatternSupported,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Provider
type ProviderObservedState struct {
	// Output only. In
	//  `projects/{project}/locations/{location}/providers/{provider_id}` format.
	// +kcc:proto:field=google.cloud.eventarc.v1.Provider.name
	Name *string `json:"name,omitempty"`

	// Output only. Human friendly name for the Provider. For example "Cloud
	//  Storage".
	// +kcc:proto:field=google.cloud.eventarc.v1.Provider.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. Event types for this provider.
	// +kcc:proto:field=google.cloud.eventarc.v1.Provider.event_types
	EventTypes []EventType `json:"eventTypes,omitempty"`
}
