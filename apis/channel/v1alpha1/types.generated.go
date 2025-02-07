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


// +kcc:proto=google.cloud.channel.v1.MarketingInfo
type MarketingInfo struct {
	// Human readable name.
	// +kcc:proto:field=google.cloud.channel.v1.MarketingInfo.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Human readable description. Description can contain HTML.
	// +kcc:proto:field=google.cloud.channel.v1.MarketingInfo.description
	Description *string `json:"description,omitempty"`

	// Default logo.
	// +kcc:proto:field=google.cloud.channel.v1.MarketingInfo.default_logo
	DefaultLogo *Media `json:"defaultLogo,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.Media
type Media struct {
	// Title of the media.
	// +kcc:proto:field=google.cloud.channel.v1.Media.title
	Title *string `json:"title,omitempty"`

	// URL of the media.
	// +kcc:proto:field=google.cloud.channel.v1.Media.content
	Content *string `json:"content,omitempty"`

	// Type of the media.
	// +kcc:proto:field=google.cloud.channel.v1.Media.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.Sku
type Sku struct {
	// Resource Name of the SKU.
	//  Format: products/{product_id}/skus/{sku_id}
	// +kcc:proto:field=google.cloud.channel.v1.Sku.name
	Name *string `json:"name,omitempty"`

	// Marketing information for the SKU.
	// +kcc:proto:field=google.cloud.channel.v1.Sku.marketing_info
	MarketingInfo *MarketingInfo `json:"marketingInfo,omitempty"`

	// Product the SKU is associated with.
	// +kcc:proto:field=google.cloud.channel.v1.Sku.product
	Product *Product `json:"product,omitempty"`
}
