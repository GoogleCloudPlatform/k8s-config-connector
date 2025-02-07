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


// +kcc:proto=google.analytics.admin.v1beta.KeyEvent
type KeyEvent struct {

	// Immutable. The event name for this key event.
	//  Examples: 'click', 'purchase'
	// +kcc:proto:field=google.analytics.admin.v1beta.KeyEvent.event_name
	EventName *string `json:"eventName,omitempty"`

	// Required. The method by which Key Events will be counted across multiple
	//  events within a session.
	// +kcc:proto:field=google.analytics.admin.v1beta.KeyEvent.counting_method
	CountingMethod *string `json:"countingMethod,omitempty"`

	// Optional. Defines a default value/currency for a key event.
	// +kcc:proto:field=google.analytics.admin.v1beta.KeyEvent.default_value
	DefaultValue *KeyEvent_DefaultValue `json:"defaultValue,omitempty"`
}

// +kcc:proto=google.analytics.admin.v1beta.KeyEvent.DefaultValue
type KeyEvent_DefaultValue struct {
	// Required. This will be used to populate the "value" parameter for all
	//  occurrences of this Key Event (specified by event_name) where that
	//  parameter is unset.
	// +kcc:proto:field=google.analytics.admin.v1beta.KeyEvent.DefaultValue.numeric_value
	NumericValue *float64 `json:"numericValue,omitempty"`

	// Required. When an occurrence of this Key Event (specified by event_name)
	//  has no set currency this currency will be applied as the default. Must be
	//  in ISO 4217 currency code format.
	//
	//  See https://en.wikipedia.org/wiki/ISO_4217 for more information.
	// +kcc:proto:field=google.analytics.admin.v1beta.KeyEvent.DefaultValue.currency_code
	CurrencyCode *string `json:"currencyCode,omitempty"`
}

// +kcc:proto=google.analytics.admin.v1beta.KeyEvent
type KeyEventObservedState struct {
	// Output only. Resource name of this key event.
	//  Format: properties/{property}/keyEvents/{key_event}
	// +kcc:proto:field=google.analytics.admin.v1beta.KeyEvent.name
	Name *string `json:"name,omitempty"`

	// Output only. Time when this key event was created in the property.
	// +kcc:proto:field=google.analytics.admin.v1beta.KeyEvent.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. If set to true, this event can be deleted.
	// +kcc:proto:field=google.analytics.admin.v1beta.KeyEvent.deletable
	Deletable *bool `json:"deletable,omitempty"`

	// Output only. If set to true, this key event refers to a custom event.  If
	//  set to false, this key event refers to a default event in GA. Default
	//  events typically have special meaning in GA. Default events are usually
	//  created for you by the GA system, but in some cases can be created by
	//  property admins. Custom events count towards the maximum number of
	//  custom key events that may be created per property.
	// +kcc:proto:field=google.analytics.admin.v1beta.KeyEvent.custom
	Custom *bool `json:"custom,omitempty"`
}
