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

// +kcc:proto=google.cloud.apigeeregistry.v1.ApiDeployment
type ApiDeployment struct {
	// Resource name.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.name
	Name *string `json:"name,omitempty"`

	// Human-meaningful name.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A detailed description.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.description
	Description *string `json:"description,omitempty"`

	// The full resource name (including revision ID) of the spec of the API being
	//  served by the deployment. Changes to this value will update the revision.
	//  Format: `apis/{api}/deployments/{deployment}`
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.api_spec_revision
	ApiSpecRevision *string `json:"apiSpecRevision,omitempty"`

	// The address where the deployment is serving. Changes to this value will
	//  update the revision.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.endpoint_uri
	EndpointURI *string `json:"endpointURI,omitempty"`

	// The address of the external channel of the API (e.g., the Developer
	//  Portal). Changes to this value will not affect the revision.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.external_channel_uri
	ExternalChannelURI *string `json:"externalChannelURI,omitempty"`

	// Text briefly identifying the intended audience of the API. Changes to this
	//  value will not affect the revision.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.intended_audience
	IntendedAudience *string `json:"intendedAudience,omitempty"`

	// Text briefly describing how to access the endpoint. Changes to this value
	//  will not affect the revision.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.access_guidance
	AccessGuidance *string `json:"accessGuidance,omitempty"`

	// Labels attach identifying metadata to resources. Identifying metadata can
	//  be used to filter list operations.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//  No more than 64 user labels can be associated with one resource (System
	//  labels are excluded).
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	//  System reserved label keys are prefixed with
	//  `apigeeregistry.googleapis.com/` and cannot be changed.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Annotations attach non-identifying metadata to resources.
	//
	//  Annotation keys and values are less restricted than those of labels, but
	//  should be generally used for small values of broad interest. Larger, topic-
	//  specific metadata should be stored in Artifacts.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// +kcc:proto=google.cloud.apigeeregistry.v1.ApiDeployment
type ApiDeploymentObservedState struct {
	// Output only. Immutable. The revision ID of the deployment.
	//  A new revision is committed whenever the deployment contents are changed.
	//  The format is an 8-character hexadecimal string.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.revision_id
	RevisionID *string `json:"revisionID,omitempty"`

	// Output only. Creation timestamp; when the deployment resource was created.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Revision creation timestamp; when the represented revision was created.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.revision_create_time
	RevisionCreateTime *string `json:"revisionCreateTime,omitempty"`

	// Output only. Last update timestamp: when the represented revision was last modified.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.revision_update_time
	RevisionUpdateTime *string `json:"revisionUpdateTime,omitempty"`
}
