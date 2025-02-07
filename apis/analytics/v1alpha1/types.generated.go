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


// +kcc:proto=google.analytics.admin.v1beta.ConversionEvent
type ConversionEvent struct {

	// Immutable. The event name for this conversion event.
	//  Examples: 'click', 'purchase'
	// +kcc:proto:field=google.analytics.admin.v1beta.ConversionEvent.event_name
	EventName *string `json:"eventName,omitempty"`

	// Optional. The method by which conversions will be counted across multiple
	//  events within a session. If this value is not provided, it will be set to
	//  `ONCE_PER_EVENT`.
	// +kcc:proto:field=google.analytics.admin.v1beta.ConversionEvent.counting_method
	CountingMethod *string `json:"countingMethod,omitempty"`

	// Optional. Defines a default value/currency for a conversion event.
	// +kcc:proto:field=google.analytics.admin.v1beta.ConversionEvent.default_conversion_value
	DefaultConversionValue *ConversionEvent_DefaultConversionValue `json:"defaultConversionValue,omitempty"`
}

// +kcc:proto=google.analytics.admin.v1beta.ConversionEvent.DefaultConversionValue
type ConversionEvent_DefaultConversionValue struct {
	// This value will be used to populate the value for all conversions
	//  of the specified event_name where the event "value" parameter is unset.
	// +kcc:proto:field=google.analytics.admin.v1beta.ConversionEvent.DefaultConversionValue.value
	Value *float64 `json:"value,omitempty"`

	// When a conversion event for this event_name has no set currency,
	//  this currency will be applied as the default. Must be in ISO 4217
	//  currency code format. See https://en.wikipedia.org/wiki/ISO_4217 for
	//  more information.
	// +kcc:proto:field=google.analytics.admin.v1beta.ConversionEvent.DefaultConversionValue.currency_code
	CurrencyCode *string `json:"currencyCode,omitempty"`
}

// +kcc:proto=google.analytics.admin.v1beta.ConversionEvent
type ConversionEventObservedState struct {
	// Output only. Resource name of this conversion event.
	//  Format: properties/{property}/conversionEvents/{conversion_event}
	// +kcc:proto:field=google.analytics.admin.v1beta.ConversionEvent.name
	Name *string `json:"name,omitempty"`

	// Output only. Time when this conversion event was created in the property.
	// +kcc:proto:field=google.analytics.admin.v1beta.ConversionEvent.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. If set, this event can currently be deleted with
	//  DeleteConversionEvent.
	// +kcc:proto:field=google.analytics.admin.v1beta.ConversionEvent.deletable
	Deletable *bool `json:"deletable,omitempty"`

	// Output only. If set to true, this conversion event refers to a custom
	//  event.  If set to false, this conversion event refers to a default event in
	//  GA. Default events typically have special meaning in GA. Default events are
	//  usually created for you by the GA system, but in some cases can be created
	//  by property admins. Custom events count towards the maximum number of
	//  custom conversion events that may be created per property.
	// +kcc:proto:field=google.analytics.admin.v1beta.ConversionEvent.custom
	Custom *bool `json:"custom,omitempty"`
}
