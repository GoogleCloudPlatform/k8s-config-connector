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


// +kcc:proto=google.cloud.visionai.v1.Asset
type Asset struct {
	// Resource name of the asset.
	//  Format:
	//  `projects/{project_number}/locations/{location_id}/corpora/{corpus_id}/assets/{asset_id}`
	// +kcc:proto:field=google.cloud.visionai.v1.Asset.name
	Name *string `json:"name,omitempty"`

	// The duration for which all media assets, associated metadata, and search
	//  documents can exist. If not set, then it will using the default ttl in the
	//  parent corpus resource.
	// +kcc:proto:field=google.cloud.visionai.v1.Asset.ttl
	Ttl *string `json:"ttl,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.AssetSource.AssetGcsSource
type AssetSource_AssetGcsSource struct {
	// Cloud storage uri.
	// +kcc:proto:field=google.cloud.visionai.v1.AssetSource.AssetGcsSource.gcs_uri
	GcsURI *string `json:"gcsURI,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Asset
type AssetObservedState struct {
	// Output only. The original cloud storage source uri that is associated with
	//  this asset.
	// +kcc:proto:field=google.cloud.visionai.v1.Asset.asset_gcs_source
	AssetGcsSource *AssetSource_AssetGcsSource `json:"assetGcsSource,omitempty"`
}
