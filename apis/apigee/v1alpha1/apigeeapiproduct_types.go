// Copyright 2026 Google LLC
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

import (
	apigeev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ApigeeAPIProductGVK = GroupVersion.WithKind("ApigeeAPIProduct")

// ApigeeAPIProductSpec defines the desired state of ApigeeAPIProduct
// +kcc:spec:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ApiProduct
type ApigeeAPIProductSpec struct {
	// The Apigee Organization that this resource belongs to.
	// +required
	OrganizationRef *apigeev1beta1.ApigeeOrganizationRef `json:"organizationRef"`

	// The ApigeeAPIProduct name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Comma-separated list of API resources to be bundled in the API product.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ApiProduct.api_resources
	APIResources []string `json:"apiResources,omitempty"`

	// Flag that specifies how API keys are approved to access the APIs defined by the API product. If set to `manual`, the consumer key is generated and returned in "pending" state. In this case, the API keys won't work until they have been explicitly approved. If set to `auto`, the consumer key is generated and returned in "approved" state and can be used immediately.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ApiProduct.approval_type
	ApprovalType *string `json:"approvalType,omitempty"`

	// Array of attributes that may be used to extend the default API product profile with customer-specific metadata.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ApiProduct.attributes
	Attributes []GoogleCloudApigeeV1Attribute `json:"attributes,omitempty"`

	// Description of the API product.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ApiProduct.description
	Description *string `json:"description,omitempty"`

	// Name displayed in the UI or developer portal to developers registering for API access.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ApiProduct.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Comma-separated list of environment names to which the API product is bound.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ApiProduct.environments
	Environments []string `json:"environments,omitempty"`

	// Configuration used to group Apigee proxies or remote services with graphQL operation name, graphQL operation type and quotas.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ApiProduct.graphql_operation_group
	GraphQLOperationGroup *GoogleCloudApigeeV1GraphQlOperationGroup `json:"graphqlOperationGroup,omitempty"`

	// Configuration used to group Apigee proxies with gRPC services and method names.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ApiProduct.grpc_operation_group
	GrpcOperationGroup *GoogleCloudApigeeV1GrpcOperationGroup `json:"grpcOperationGroup,omitempty"`

	// Configuration used to group Apigee proxies or remote services with LLM operation name and quotas.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ApiProduct.llm_operation_group
	LlmOperationGroup *GoogleCloudApigeeV1LlmOperationGroup `json:"llmOperationGroup,omitempty"`

	// Configuration used to group Apigee proxies or remote services with operations and quotas.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ApiProduct.operation_group
	OperationGroup *GoogleCloudApigeeV1OperationGroup `json:"operationGroup,omitempty"`

	// Comma-separated list of API proxy names to which this API product is bound.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ApiProduct.proxies
	Proxies []string `json:"proxies,omitempty"`

	// Quota associated with the API product.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ApiProduct.quota
	Quota *string `json:"quota,omitempty"`

	// Scope of the quota decides how the quota counter gets applied and evaluate for quota violation.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ApiProduct.quota_counter_scope
	QuotaCounterScope *string `json:"quotaCounterScope,omitempty"`

	// Time interval over which the number of request messages is calculated.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ApiProduct.quota_interval
	QuotaInterval *string `json:"quotaInterval,omitempty"`

	// Time unit defined for the `quota_interval`. Valid values include `minute`, `hour`, `day`, or `month`.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ApiProduct.quota_time_unit
	QuotaTimeUnit *string `json:"quotaTimeUnit,omitempty"`

	// Comma-separated list of OAuth scopes validated at runtime.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ApiProduct.scopes
	Scopes []string `json:"scopes,omitempty"`
}

// ApigeeAPIProductStatus defines the config connector machine state of ApigeeAPIProduct
type ApigeeAPIProductStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ApigeeAPIProduct resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ApigeeAPIProductObservedState `json:"observedState,omitempty"`
}

// ApigeeAPIProductObservedState is the state of the ApigeeAPIProduct resource as most recently observed in GCP.
// +kcc:observedstate:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ApiProduct
type ApigeeAPIProductObservedState struct {
	// Response only. Creation time of this environment as milliseconds since epoch.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ApiProduct.created_at
	CreatedAt *int64 `json:"createdAt,omitempty"`

	// Response only. Modified time of this environment as milliseconds since epoch.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ApiProduct.last_modified_at
	LastModifiedAt *int64 `json:"lastModifiedAt,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapigeeapiproduct;gcpapigeeapiproducts
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ApigeeAPIProduct is the Schema for the ApigeeAPIProduct API
// +k8s:openapi-gen=true
type ApigeeAPIProduct struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ApigeeAPIProductSpec   `json:"spec,omitempty"`
	Status ApigeeAPIProductStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ApigeeAPIProductList contains a list of ApigeeAPIProduct
type ApigeeAPIProductList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApigeeAPIProduct `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApigeeAPIProduct{}, &ApigeeAPIProductList{})
}
