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


// +kcc:proto=google.cloud.migrationcenter.v1.ImportDataFile
type ImportDataFile struct {

	// User-friendly display name. Maximum length is 63 characters.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportDataFile.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The payload format.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportDataFile.format
	Format *string `json:"format,omitempty"`

	// Information about a file that is uploaded to a storage service.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportDataFile.upload_file_info
	UploadFileInfo *UploadFileInfo `json:"uploadFileInfo,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.UploadFileInfo
type UploadFileInfo struct {
}

// +kcc:proto=google.cloud.migrationcenter.v1.ImportDataFile
type ImportDataFileObservedState struct {
	// Output only. The name of the file.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportDataFile.name
	Name *string `json:"name,omitempty"`

	// Output only. The timestamp when the file was created.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportDataFile.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The state of the import data file.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportDataFile.state
	State *string `json:"state,omitempty"`

	// Information about a file that is uploaded to a storage service.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportDataFile.upload_file_info
	UploadFileInfo *UploadFileInfoObservedState `json:"uploadFileInfo,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.UploadFileInfo
type UploadFileInfoObservedState struct {
	// Output only. Upload URI for the file.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.UploadFileInfo.signed_uri
	SignedURI *string `json:"signedURI,omitempty"`

	// Output only. The headers that were used to sign the URI.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.UploadFileInfo.headers
	Headers map[string]string `json:"headers,omitempty"`

	// Output only. Expiration time of the upload URI.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.UploadFileInfo.uri_expiration_time
	URIExpirationTime *string `json:"uriExpirationTime,omitempty"`
}
