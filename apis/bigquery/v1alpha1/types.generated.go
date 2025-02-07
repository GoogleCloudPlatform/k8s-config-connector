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


// +kcc:proto=google.cloud.bigquery.dataexchange.v1beta1.DataExchange
type DataExchange struct {

	// Required. Human-readable display name of the data exchange. The display name must
	//  contain only Unicode letters, numbers (0-9), underscores (_), dashes (-),
	//  spaces ( ), ampersands (&) and must not start or end with spaces.
	//  Default value is an empty string.
	//  Max length: 63 bytes.
	// +kcc:proto:field=google.cloud.bigquery.dataexchange.v1beta1.DataExchange.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description of the data exchange. The description must not contain Unicode
	//  non-characters as well as C0 and C1 control codes except tabs (HT),
	//  new lines (LF), carriage returns (CR), and page breaks (FF).
	//  Default value is an empty string.
	//  Max length: 2000 bytes.
	// +kcc:proto:field=google.cloud.bigquery.dataexchange.v1beta1.DataExchange.description
	Description *string `json:"description,omitempty"`

	// Optional. Email or URL of the primary point of contact of the data exchange.
	//  Max Length: 1000 bytes.
	// +kcc:proto:field=google.cloud.bigquery.dataexchange.v1beta1.DataExchange.primary_contact
	PrimaryContact *string `json:"primaryContact,omitempty"`

	// Optional. Documentation describing the data exchange.
	// +kcc:proto:field=google.cloud.bigquery.dataexchange.v1beta1.DataExchange.documentation
	Documentation *string `json:"documentation,omitempty"`

	// Optional. Base64 encoded image representing the data exchange. Max Size: 3.0MiB
	//  Expected image dimensions are 512x512 pixels, however the API only
	//  performs validation on size of the encoded data.
	//  Note: For byte fields, the content of the fields are base64-encoded (which
	//  increases the size of the data by 33-36%) when using JSON on the wire.
	// +kcc:proto:field=google.cloud.bigquery.dataexchange.v1beta1.DataExchange.icon
	Icon []byte `json:"icon,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.dataexchange.v1beta1.DataExchange
type DataExchangeObservedState struct {
	// Output only. The resource name of the data exchange.
	//  e.g. `projects/myproject/locations/US/dataExchanges/123`.
	// +kcc:proto:field=google.cloud.bigquery.dataexchange.v1beta1.DataExchange.name
	Name *string `json:"name,omitempty"`

	// Output only. Number of listings contained in the data exchange.
	// +kcc:proto:field=google.cloud.bigquery.dataexchange.v1beta1.DataExchange.listing_count
	ListingCount *int32 `json:"listingCount,omitempty"`
}
