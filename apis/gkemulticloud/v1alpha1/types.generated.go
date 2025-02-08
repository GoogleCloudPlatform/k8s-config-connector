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


// +kcc:proto=google.cloud.gkemulticloud.v1.AzureClient
type AzureClient struct {
	// The name of this resource.
	//
	//  `AzureClient` resource names are formatted as
	//  `projects/<project-number>/locations/<region>/azureClients/<client-id>`.
	//
	//  See [Resource Names](https://cloud.google.com/apis/design/resource_names)
	//  for more details on Google Cloud resource names.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureClient.name
	Name *string `json:"name,omitempty"`

	// Required. The Azure Active Directory Tenant ID.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureClient.tenant_id
	TenantID *string `json:"tenantID,omitempty"`

	// Required. The Azure Active Directory Application ID.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureClient.application_id
	ApplicationID *string `json:"applicationID,omitempty"`

	// Optional. Annotations on the resource.
	//
	//  This field has the same restrictions as Kubernetes annotations.
	//  The total size of all keys and values combined is limited to 256k.
	//  Keys can have 2 segments: prefix (optional) and name (required),
	//  separated by a slash (/).
	//  Prefix must be a DNS subdomain.
	//  Name must be 63 characters or less, begin and end with alphanumerics,
	//  with dashes (-), underscores (_), dots (.), and alphanumerics between.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureClient.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AzureClient
type AzureClientObservedState struct {
	// Output only. If set, there are currently pending changes to the client.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureClient.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. The PEM encoded x509 certificate.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureClient.pem_certificate
	PemCertificate *string `json:"pemCertificate,omitempty"`

	// Output only. A globally unique identifier for the client.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureClient.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time at which this resource was created.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureClient.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which this client was last updated.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureClient.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
