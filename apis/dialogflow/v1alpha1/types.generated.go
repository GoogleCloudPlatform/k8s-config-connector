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


// +kcc:proto=google.cloud.dialogflow.v2.Fulfillment
type Fulfillment struct {
	// Required. The unique identifier of the fulfillment.
	//  Supported formats:
	//
	//  - `projects/<Project ID>/agent/fulfillment`
	//  - `projects/<Project ID>/locations/<Location ID>/agent/fulfillment`
	//
	//  This field is not used for Fulfillment in an Environment.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Fulfillment.name
	Name *string `json:"name,omitempty"`

	// Optional. The human-readable name of the fulfillment, unique within the
	//  agent.
	//
	//  This field is not used for Fulfillment in an Environment.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Fulfillment.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Configuration for a generic web service.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Fulfillment.generic_web_service
	GenericWebService *Fulfillment_GenericWebService `json:"genericWebService,omitempty"`

	// Optional. Whether fulfillment is enabled.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Fulfillment.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Optional. The field defines whether the fulfillment is enabled for certain
	//  features.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Fulfillment.features
	Features []Fulfillment_Feature `json:"features,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.Fulfillment.Feature
type Fulfillment_Feature struct {
	// The type of the feature that enabled for fulfillment.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Fulfillment.Feature.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.Fulfillment.GenericWebService
type Fulfillment_GenericWebService struct {
	// Required. The fulfillment URI for receiving POST requests.
	//  It must use https protocol.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Fulfillment.GenericWebService.uri
	URI *string `json:"uri,omitempty"`

	// Optional. The user name for HTTP Basic authentication.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Fulfillment.GenericWebService.username
	Username *string `json:"username,omitempty"`

	// Optional. The password for HTTP Basic authentication.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Fulfillment.GenericWebService.password
	Password *string `json:"password,omitempty"`

	// Optional. The HTTP request headers to send together with fulfillment
	//  requests.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Fulfillment.GenericWebService.request_headers
	RequestHeaders map[string]string `json:"requestHeaders,omitempty"`

	// Optional. Indicates if generic web service is created through Cloud
	//  Functions integration. Defaults to false.
	//
	//  is_cloud_function is deprecated. Cloud functions can be configured by
	//  its uri as a regular web service now.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Fulfillment.GenericWebService.is_cloud_function
	IsCloudFunction *bool `json:"isCloudFunction,omitempty"`
}
