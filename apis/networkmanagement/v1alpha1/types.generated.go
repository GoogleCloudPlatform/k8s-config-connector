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

// +generated:types
// krm.group: networkmanagement.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.networkmanagement.v1
// resource: NetworkManagement:ConnectivityTest

package v1alpha1

// +kcc:proto=google.cloud.networkmanagement.v1.Endpoint.AppEngineVersionEndpoint
type Endpoint_AppEngineVersionEndpoint struct {
	// An [App Engine](https://cloud.google.com/appengine) [service
	//  version](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions)
	//  name.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.AppEngineVersionEndpoint.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.Endpoint.CloudFunctionEndpoint
type Endpoint_CloudFunctionEndpoint struct {
	// A [Cloud Function](https://cloud.google.com/functions) name.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.CloudFunctionEndpoint.uri
	URI *string `json:"uri,omitempty"`
}
