// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

// +tool:krm-type-terraform
// proto.message: google.appengine.v1.DomainMapping
// crd.kind: AppEngineDomainMapping
// crd.version: v1alpha1
// terraform.src: github.com/hashicorp/terraform-provider-google-beta/google-beta/services/appengine/resource_app_engine_domain_mapping.go

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AppEngineDomainMappingGVK = GroupVersion.WithKind("AppEngineDomainMapping")

// AppEngineDomainMappingSpec defines the desired state of AppEngineDomainMapping
// +kcc:spec:proto=google.appengine.v1.DomainMapping
type AppEngineDomainMappingSpec struct {
	// The project that this resource belongs to.
	Project *string `json:"project,omitempty"`

	// TODO: Should be projectRef
	// +required
	// ProjectRef *v1beta1.ProjectRef `json:"projectRef,omitempty"`

	// Location *string `json:"location,omitempty"`

	// The AppEngineDomainMapping name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// NOTYET(teraform)
	// // Relative name of the domain serving the application. Example:
	// //  `example.com`.
	// // +kcc:proto:field=google.appengine.v1.DomainMapping.id
	// ID *string `json:"id,omitempty"`

	// TODO: Remove overrideStrategy?

	// 	Whether the domain creation should override any existing mappings for this domain.
	//  By default, overrides are rejected. Default value: "STRICT" Possible values: ["STRICT", "OVERRIDE"].
	OverrideStrategy *string `json:"overrideStrategy,omitempty"`

	// SSL configuration for this domain. If unconfigured, this domain will not
	//  serve with SSL.
	// +kcc:proto:field=google.appengine.v1.DomainMapping.ssl_settings
	SSLSettings *SSLSettings `json:"sslSettings,omitempty"`
}

// Removed from generated because of Id vs ID (and should be Refs anyway, I think)

// +kcc:proto=google.appengine.v1.SslSettings
type SSLSettings struct {
	// ID of the `AuthorizedCertificate` resource configuring SSL for the
	//  application. Clearing this field will remove SSL support.
	//
	//  By default, a managed certificate is automatically created for every
	//  domain mapping. To omit SSL support or to configure SSL manually, specify
	//  `SslManagementType.MANUAL` on a `CREATE` or `UPDATE` request. You must
	//  be authorized to administer the `AuthorizedCertificate` resource to
	//  manually map it to a `DomainMapping` resource.
	//  Example: `12345`.
	// +kcc:proto:field=google.appengine.v1.SslSettings.certificate_id
	CertificateID *string `json:"certificateId,omitempty"`

	// SSL management type for this domain. If `AUTOMATIC`, a managed certificate
	//  is automatically provisioned. If `MANUAL`, `certificate_id` must be
	//  manually specified in order to configure SSL for this domain.
	// +kcc:proto:field=google.appengine.v1.SslSettings.ssl_management_type
	// +required
	SSLManagementType *string `json:"sslManagementType,omitempty"`

	// ID of the managed `AuthorizedCertificate` resource currently being
	//  provisioned, if applicable. Until the new managed certificate has been
	//  successfully provisioned, the previous SSL state will be preserved. Once
	//  the provisioning process completes, the `certificate_id` field will reflect
	//  the new managed certificate and this field will be left empty. To remove
	//  SSL support while there is still a pending managed certificate, clear the
	//  `certificate_id` field with an `UpdateDomainMappingRequest`.
	//
	//  @OutputOnly
	// +kcc:proto:field=google.appengine.v1.SslSettings.pending_managed_certificate_id
	PendingManagedCertificateID *string `json:"pendingManagedCertificateId,omitempty"`
}

// AppEngineDomainMappingStatus defines the config connector machine state of AppEngineDomainMapping
type AppEngineDomainMappingStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AppEngineDomainMapping resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// Full path to the `DomainMapping` resource in the API. Example:
	//  `apps/myapp/domainMapping/example.com`.
	// +kcc:proto:field=google.appengine.v1.DomainMapping.name
	Name *string `json:"name,omitempty"`

	// The resource records required to configure this domain mapping. These
	//  records must be added to the domain's DNS configuration in order to
	//  serve the application via this domain mapping.
	// +kcc:proto:field=google.appengine.v1.DomainMapping.resource_records
	ResourceRecords []DomainMapping_ResourceRecord `json:"resourceRecords,omitempty"`

	// NOTYET(teraform)
	// // ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *AppEngineDomainMappingObservedState `json:"observedState,omitempty"`
}

// +kcc:proto=google.appengine.v1.ResourceRecord
type DomainMapping_ResourceRecord struct {
	// Relative name of the object affected by this record. Only applicable for CNAME records. Example: 'www'.
	// +kcc:proto:field=google.appengine.v1.ResourceRecord.name
	Name *string `json:"name,omitempty"`

	// Data for this record. Values vary by record type, as defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1).
	// +kcc:proto:field=google.appengine.v1.ResourceRecord.rrdata
	Rrdata *string `json:"rrdata,omitempty"`

	// Resource record type. Example: 'AAAA'. Possible values: ["A", "AAAA", "CNAME"]
	// +kcc:proto:field=google.appengine.v1.ResourceRecord.type
	Type *string `json:"type,omitempty"`
}

// AppEngineDomainMappingObservedState is the state of the AppEngineDomainMapping resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.appengine.v1.DomainMapping
type AppEngineDomainMappingObservedState struct {
	// SSL configuration for this domain. If unconfigured, this domain will not
	//  serve with SSL.
	// +kcc:proto:field=google.appengine.v1.DomainMapping.ssl_settings
	SslSettings *DomainMapping_SslSettingsObservedState `json:"sslSettings,omitempty"`
}

// +kcc:observedstate:proto=google.appengine.v1.DomainMapping.SslSettings
type DomainMapping_SslSettingsObservedState struct {
	// ID of the managed 'AuthorizedCertificate' resource currently being provisioned, if applicable. Until the new
	//  managed certificate has been successfully provisioned, the previous SSL state will be preserved. Once the
	//  provisioning process completes, the 'certificateId' field will reflect the new managed certificate and this
	//  field will be left empty. To remove SSL support while there is still a pending managed certificate, clear the
	//  'certificateId' field with an update request.
	// +kcc:proto:field=google.appengine.v1.DomainMapping.SslSettings.pending_managed_certificate_id
	PendingManagedCertificateID *string `json:"pendingManagedCertificateId,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpappenginedomainmapping;gcpappenginedomainmappings
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AppEngineDomainMapping is the Schema for the AppEngineDomainMapping API
// +k8s:openapi-gen=true
type AppEngineDomainMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppEngineDomainMappingSpec   `json:"spec,omitempty"`
	Status AppEngineDomainMappingStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AppEngineDomainMappingList contains a list of AppEngineDomainMapping
type AppEngineDomainMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppEngineDomainMapping `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AppEngineDomainMapping{}, &AppEngineDomainMappingList{})
}
