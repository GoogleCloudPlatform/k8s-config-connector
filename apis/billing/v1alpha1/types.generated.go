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


// +kcc:proto=google.cloud.billing.v1.Service
type Service struct {
	// The resource name for the service.
	//  Example: "services/6F81-5844-456A"
	// +kcc:proto:field=google.cloud.billing.v1.Service.name
	Name *string `json:"name,omitempty"`

	// The identifier for the service.
	//  Example: "6F81-5844-456A"
	// +kcc:proto:field=google.cloud.billing.v1.Service.service_id
	ServiceID *string `json:"serviceID,omitempty"`

	// A human readable display name for this service.
	// +kcc:proto:field=google.cloud.billing.v1.Service.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The business under which the service is offered.
	//  Ex. "businessEntities/GCP", "businessEntities/Maps"
	// +kcc:proto:field=google.cloud.billing.v1.Service.business_entity_name
	BusinessEntityName *string `json:"businessEntityName,omitempty"`
}
