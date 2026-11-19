// Copyright 2026 Google LLC
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

// API sources for BigLakeIcebergCatalog, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/biglake/v1/iceberg_rest_catalog.proto
// +kcc:source:service-docs=https://cloud.google.com/products/lakehouse
// +kcc:source:resource-docs=https://cloud.google.com/bigquery/docs/iceberg-rest-catalog-about

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	secretmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	servicedirectoryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/servicedirectory/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigLakeIcebergCatalogGVK = GroupVersion.WithKind("BigLakeIcebergCatalog")

// BigLakeIcebergCatalogSpec defines the desired state of BigLakeIcebergCatalog
// +kcc:spec:proto=google.cloud.biglake.v1.IcebergCatalog
// +kcc:required-from-proto
type BigLakeIcebergCatalogSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The BigLakeIcebergCatalog name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. The credential mode for the catalog.
	// +kcc:proto:field=google.cloud.biglake.v1.IcebergCatalog.credential_mode
	CredentialMode *string `json:"credentialMode,omitempty"`

	// Required. The catalog type. Required for CreateIcebergCatalog.
	// +kcc:proto:field=google.cloud.biglake.v1.IcebergCatalog.catalog_type
	// +required
	CatalogType *string `json:"catalogType,omitempty"`

	// Optional. The default storage location for the catalog, e.g.,
	//  `gs://my-bucket`. For Google Cloud Storage bucket catalogs, this is output
	//  only.
	//
	//  For BigLake catalogs, this field must be provided and point to a
	//  Google Cloud Storage bucket or a path within that bucket. This path serves
	//  as the base directory for constructing the full path to a table's data and
	//  metadata directories when a location is not specified at the namespace or
	//  table level. The full path is formed by appending the namespace and table
	//  identifiers to the default location.
	// +kcc:proto:field=google.cloud.biglake.v1.IcebergCatalog.default_location
	DefaultLocation *string `json:"defaultLocation,omitempty"`

	// Optional. A user-provided description of the catalog. The description must
	//  be a UTF-8 string with a maximum length of 1024 characters.
	// +kcc:proto:field=google.cloud.biglake.v1.IcebergCatalog.description
	Description *string `json:"description,omitempty"`

	// Optional. Restricted locations configuration. This field is currently only
	//  used for BigLake catalogs.
	//
	//  If this field is unset, or if
	//  `restricted_locations_config.restricted_locations` is empty, all
	//  accessible locations are allowed. If
	//  `restricted_locations_config.restricted_locations` is not empty, only
	//  locations in `default_location` and
	//  `restricted_locations_config.restricted_locations` are allowed.
	// +kcc:proto:field=google.cloud.biglake.v1.IcebergCatalog.restricted_locations_config
	RestrictedLocationsConfig *IcebergCatalog_RestrictedLocationsConfig `json:"restrictedLocationsConfig,omitempty"`

	// Optional. Configuration options for federated catalogs.
	// +kcc:proto:field=google.cloud.biglake.v1.IcebergCatalog.federated_catalog_options
	FederatedCatalogOptions *IcebergCatalog_FederatedCatalogOptions `json:"federatedCatalogOptions,omitempty"`
}

// +kcc:proto=google.cloud.biglake.v1.IcebergCatalog.FederatedCatalogOptions
type IcebergCatalog_FederatedCatalogOptions struct {
	// Optional. Info specific to Unity Catalog by Databricks.
	// +kcc:proto:field=google.cloud.biglake.v1.IcebergCatalog.FederatedCatalogOptions.unity_catalog_info
	UnityCatalogInfo *IcebergCatalog_FederatedCatalogOptions_UnityCatalogInfo `json:"unityCatalogInfo,omitempty"`

	// Optional. Info specific to an AWS Glue Catalog.
	// +kcc:proto:field=google.cloud.biglake.v1.IcebergCatalog.FederatedCatalogOptions.glue_catalog_info
	GlueCatalogInfo *IcebergCatalog_FederatedCatalogOptions_GlueCatalogInfo `json:"glueCatalogInfo,omitempty"`

	// Optional. Info specific to a Snowflake Catalog.
	// +kcc:proto:field=google.cloud.biglake.v1.IcebergCatalog.FederatedCatalogOptions.snowflake_catalog_info
	SnowflakeCatalogInfo *IcebergCatalog_FederatedCatalogOptions_SnowflakeCatalogInfo `json:"snowflakeCatalogInfo,omitempty"`

	// Optional. The secret resource name in Secret Manager, in the format
	//  `projects/{project_id}/locations/{location}/secrets/{secret_id}` or
	//  `projects/{project_id}/locations/{location}/secrets/{secret_id}/versions/{version_id}`.
	//
	//  The project ID must match the catalog's project and location must match
	//  the catalog's location.
	//  If the version is not specified, the latest version will be used.
	//
	//  This field is not used when
	//  [google.cloud.biglake.v1main.IcebergCatalog.FederatedCatalogOptions.UnityCatalogInfo.service_principal_application_id][google.cloud.biglake.v1main.IcebergCatalog.FederatedCatalogOptions.UnityCatalogInfo.service_principal_application_id]
	//  or
	//  [google.cloud.biglake.v1main.IcebergCatalog.FederatedCatalogOptions.SnowflakeCatalogInfo.snowflake_role][google.cloud.biglake.v1main.IcebergCatalog.FederatedCatalogOptions.SnowflakeCatalogInfo.snowflake_role]
	//  is set.
	// +kcc:proto:field=google.cloud.biglake.v1.IcebergCatalog.FederatedCatalogOptions.secret_name
	SecretRef *secretmanagerv1beta1.SecretManagerSecretRef `json:"secretRef,omitempty"`

	// Optional. The service directory resource name for routing traffic over a
	//  private network connection through Cross-Cloud Interconnect, in the
	//  format
	//  `projects/{project_id}/locations/{location_id}/namespaces/{namespace_id}/services/{service_id}`.
	// +kcc:proto:field=google.cloud.biglake.v1.IcebergCatalog.FederatedCatalogOptions.service_directory_name
	ServiceDirectoryRef *servicedirectoryv1beta1.ServiceDirectoryServiceRef `json:"serviceDirectoryRef,omitempty"`

	// Optional. Refresh configuration.
	// +kcc:proto:field=google.cloud.biglake.v1.IcebergCatalog.FederatedCatalogOptions.refresh_options
	RefreshOptions *IcebergCatalog_FederatedCatalogOptions_RefreshOptions `json:"refreshOptions,omitempty"`
}

// BigLakeIcebergCatalogStatus defines the config connector machine state of BigLakeIcebergCatalog
type BigLakeIcebergCatalogStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigLakeIcebergCatalog resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigLakeIcebergCatalogObservedState `json:"observedState,omitempty"`
}

// BigLakeIcebergCatalogObservedState is the state of the BigLakeIcebergCatalog resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.biglake.v1.IcebergCatalog
type BigLakeIcebergCatalogObservedState struct {
	// Output only. The service account used for credential vending, output only.
	//  Might be empty if Credential vending was never enabled for the catalog. For
	//  federated catalogs, the service account will be always provisioned and will
	//  be used to access the remote Iceberg REST Catalog using access to Secret
	//  Manager secret or identity federation.
	// +kcc:proto:field=google.cloud.biglake.v1.IcebergCatalog.biglake_service_account
	BiglakeServiceAccount *string `json:"biglakeServiceAccount,omitempty"`

	// Output only. The unique ID of the service account. This is used for
	//  federation scenarios.
	// +kcc:proto:field=google.cloud.biglake.v1.IcebergCatalog.biglake_service_account_unique_id
	BiglakeServiceAccountUniqueID *string `json:"biglakeServiceAccountUniqueID,omitempty"`

	// Output only. The GCP region(s) of the default location's bucket, e.g.
	//  `us-central1`, `nam4` or `us`. This will contain one value for all
	//  locations, except for the catalogs that are configured to use custom dual
	//  region buckets, in which case it will contain the two regions of the
	//  bucket. The region(s) of this field should be in the jurisdiction of or
	//  nearby the primary location of the catalog.
	// +kcc:proto:field=google.cloud.biglake.v1.IcebergCatalog.storage_regions
	StorageRegions []string `json:"storageRegions,omitempty"`

	// Output only. When the catalog was created.
	// +kcc:proto:field=google.cloud.biglake.v1.IcebergCatalog.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. When the catalog was last updated.
	// +kcc:proto:field=google.cloud.biglake.v1.IcebergCatalog.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The replicas for the catalog metadata.
	// +kcc:proto:field=google.cloud.biglake.v1.IcebergCatalog.replicas
	Replicas []IcebergCatalog_ReplicaObservedState `json:"replicas,omitempty"`

	// Optional. Configuration options for federated catalogs.
	// +kcc:proto:field=google.cloud.biglake.v1.IcebergCatalog.federated_catalog_options
	FederatedCatalogOptions *IcebergCatalog_FederatedCatalogOptionsObservedState `json:"federatedCatalogOptions,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbiglakeicebergcatalog;gcpbiglakeicebergcatalogs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigLakeIcebergCatalog is the Schema for the BigLakeIcebergCatalog API
// +k8s:openapi-gen=true
type BigLakeIcebergCatalog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigLakeIcebergCatalogSpec   `json:"spec,omitempty"`
	Status BigLakeIcebergCatalogStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigLakeIcebergCatalogList contains a list of BigLakeIcebergCatalog
type BigLakeIcebergCatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigLakeIcebergCatalog `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigLakeIcebergCatalog{}, &BigLakeIcebergCatalogList{})
}
