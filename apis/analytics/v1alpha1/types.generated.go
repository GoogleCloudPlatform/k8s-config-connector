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


// +kcc:proto=google.analytics.admin.v1beta.Property
type Property struct {

	// Immutable. The property type for this Property resource. When creating a
	//  property, if the type is "PROPERTY_TYPE_UNSPECIFIED", then
	//  "ORDINARY_PROPERTY" will be implied.
	// +kcc:proto:field=google.analytics.admin.v1beta.Property.property_type
	PropertyType *string `json:"propertyType,omitempty"`

	// Immutable. Resource name of this property's logical parent.
	//
	//  Note: The Property-Moving UI can be used to change the parent.
	//  Format: accounts/{account}, properties/{property}
	//  Example: "accounts/100", "properties/101"
	// +kcc:proto:field=google.analytics.admin.v1beta.Property.parent
	Parent *string `json:"parent,omitempty"`

	// Required. Human-readable display name for this property.
	//
	//  The max allowed display name length is 100 UTF-16 code units.
	// +kcc:proto:field=google.analytics.admin.v1beta.Property.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Industry associated with this property
	//  Example: AUTOMOTIVE, FOOD_AND_DRINK
	// +kcc:proto:field=google.analytics.admin.v1beta.Property.industry_category
	IndustryCategory *string `json:"industryCategory,omitempty"`

	// Required. Reporting Time Zone, used as the day boundary for reports,
	//  regardless of where the data originates. If the time zone honors DST,
	//  Analytics will automatically adjust for the changes.
	//
	//  NOTE: Changing the time zone only affects data going forward, and is not
	//  applied retroactively.
	//
	//  Format: https://www.iana.org/time-zones
	//  Example: "America/Los_Angeles"
	// +kcc:proto:field=google.analytics.admin.v1beta.Property.time_zone
	TimeZone *string `json:"timeZone,omitempty"`

	// The currency type used in reports involving monetary values.
	//
	//
	//  Format: https://en.wikipedia.org/wiki/ISO_4217
	//  Examples: "USD", "EUR", "JPY"
	// +kcc:proto:field=google.analytics.admin.v1beta.Property.currency_code
	CurrencyCode *string `json:"currencyCode,omitempty"`

	// Immutable. The resource name of the parent account
	//  Format: accounts/{account_id}
	//  Example: "accounts/123"
	// +kcc:proto:field=google.analytics.admin.v1beta.Property.account
	Account *string `json:"account,omitempty"`
}

// +kcc:proto=google.analytics.admin.v1beta.Property
type PropertyObservedState struct {
	// Output only. Resource name of this property.
	//  Format: properties/{property_id}
	//  Example: "properties/1000"
	// +kcc:proto:field=google.analytics.admin.v1beta.Property.name
	Name *string `json:"name,omitempty"`

	// Output only. Time when the entity was originally created.
	// +kcc:proto:field=google.analytics.admin.v1beta.Property.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when entity payload fields were last updated.
	// +kcc:proto:field=google.analytics.admin.v1beta.Property.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The Google Analytics service level that applies to this
	//  property.
	// +kcc:proto:field=google.analytics.admin.v1beta.Property.service_level
	ServiceLevel *string `json:"serviceLevel,omitempty"`

	// Output only. If set, the time at which this property was trashed. If not
	//  set, then this property is not currently in the trash can.
	// +kcc:proto:field=google.analytics.admin.v1beta.Property.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. If set, the time at which this trashed property will be
	//  permanently deleted. If not set, then this property is not currently in the
	//  trash can and is not slated to be deleted.
	// +kcc:proto:field=google.analytics.admin.v1beta.Property.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`
}
