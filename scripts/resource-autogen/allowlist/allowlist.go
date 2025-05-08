// Copyright 2022 Google LLC
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

package allowlist

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"
)

var (
	// alphaAllowlist holds the list of the resources to be allowlisted as
	// v1alpha1 CRDs. The format is '[terraform_product_name]/[terraform_type_name]'.
	// 'google_[terraform_product_name]' should be the prefix of
	// '[terraform_type_name]'.
	alphaAllowlist = []string{
		"access_context_manager/google_access_context_manager_access_level_condition",
		"access_context_manager/google_access_context_manager_gcp_user_access_binding",
		"access_context_manager/google_access_context_manager_service_perimeter_resource",
		"alloydb/google_alloydb_backup",
		"alloydb/google_alloydb_cluster",
		"alloydb/google_alloydb_instance",
		"api_gateway/google_api_gateway_api",
		"api_gateway/google_api_gateway_api_config",
		"api_gateway/google_api_gateway_gateway",
		"apigee/google_apigee_addons_config",
		"apigee/google_apigee_nat_address",
		"apigee/google_apigee_sync_authorization",
		"app_engine/google_app_engine_domain_mapping",
		"app_engine/google_app_engine_firewall_rule",
		"app_engine/google_app_engine_flexible_app_version",
		"app_engine/google_app_engine_service_split_traffic",
		"app_engine/google_app_engine_standard_app_version",
		"beyondcorp/google_beyondcorp_app_connection",
		"beyondcorp/google_beyondcorp_app_connector",
		"beyondcorp/google_beyondcorp_app_gateway",
		"bigquery/google_bigquery_dataset_access",
		"bigquery_datapolicy/google_bigquery_datapolicy_data_policy",
		"bigquery_reservation/google_bigquery_capacity_commitment",
		"bigquery_reservation/google_bigquery_reservation",
		"certificate_manager/google_certificate_manager_certificate",
		"certificate_manager/google_certificate_manager_certificate_map",
		"certificate_manager/google_certificate_manager_certificate_map_entry",
		"certificate_manager/google_certificate_manager_dns_authorization",
		"cloud_asset/google_cloud_asset_folder_feed",
		"cloud_asset/google_cloud_asset_organization_feed",
		"cloud_asset/google_cloud_asset_project_feed",
		"cloud_ids/google_cloud_ids_endpoint",
		"cloudfunctions2/google_cloudfunctions2_function",
		"cloudiot/google_cloudiot_device",
		"cloudiot/google_cloudiot_registry",
		"compute/google_compute_autoscaler",
		"compute/google_compute_backend_bucket_signed_url_key",
		"compute/google_compute_backend_service_signed_url_key",
		"compute/google_compute_disk_resource_policy_attachment",
		"compute/google_compute_global_network_endpoint",
		"compute/google_compute_global_network_endpoint_group",
		"compute/google_compute_instance_group_named_port",
		"compute/google_compute_machine_image",
		"compute/google_compute_managed_ssl_certificate",
		"compute/google_compute_network_endpoint",
		"compute/google_compute_network_firewall_policy_rule",
		"compute/google_compute_network_peering_routes_config",
		"compute/google_compute_organization_security_policy",
		"compute/google_compute_organization_security_policy_association",
		"compute/google_compute_organization_security_policy_rule",
		"compute/google_compute_per_instance_config",
		"compute/google_compute_region_autoscaler",
		"compute/google_compute_region_disk_resource_policy_attachment",
		"compute/google_compute_region_per_instance_config",
		"compute/google_compute_region_ssl_policy",
		"container_analysis/google_container_analysis_occurrence",
		"data_catalog/google_data_catalog_entry",
		"data_catalog/google_data_catalog_entry_group",
		"data_catalog/google_data_catalog_tag",
		"data_catalog/google_data_catalog_tag_template",
		"datastore/google_datastore_index",
		"datastream/google_datastream_connection_profile",
		"datastream/google_datastream_private_connection",
		"datastream/google_datastream_stream",
		"deployment_manager/google_deployment_manager_deployment",
		"dialogflow/google_dialogflow_agent",
		"dialogflow/google_dialogflow_entity_type",
		"dialogflow/google_dialogflow_fulfillment",
		"dialogflow/google_dialogflow_intent",
		"dialogflow_cx/google_dialogflow_cx_agent",
		"dialogflow_cx/google_dialogflow_cx_entity_type",
		"dialogflow_cx/google_dialogflow_cx_flow",
		"dialogflow_cx/google_dialogflow_cx_intent",
		"dialogflow_cx/google_dialogflow_cx_page",
		"dialogflow_cx/google_dialogflow_cx_webhook",
		"dns/google_dns_response_policy",
		"dns/google_dns_response_policy_rule",
		"document_ai/google_document_ai_processor",
		"document_ai/google_document_ai_processor_default_version",
		"essential_contacts/google_essential_contacts_contact",
		"filestore/google_filestore_snapshot",
		"firebase/google_firebase_android_app",
		"firebase/google_firebase_project",
		"firebase/google_firebase_web_app",
		"firebase_database/google_firebase_database_instance",
		"firebase_hosting/google_firebase_hosting_channel",
		"firebase_hosting/google_firebase_hosting_site",
		"firebase_storage/google_firebase_storage_bucket",
		"gke_backup/google_gke_backup_backup_plan",
		"healthcare/google_healthcare_consent_store",
		"healthcare/google_healthcare_dataset",
		"healthcare/google_healthcare_dicom_store",
		"healthcare/google_healthcare_fhir_store",
		"healthcare/google_healthcare_hl7_v2_store",
		"identity_platform/google_identity_platform_default_supported_idp_config",
		"identity_platform/google_identity_platform_inbound_saml_config",
		"identity_platform/google_identity_platform_project_default_config",
		"identity_platform/google_identity_platform_tenant_default_supported_idp_config",
		"identity_platform/google_identity_platform_tenant_inbound_saml_config",
		"kms/google_kms_crypto_key_version",
		"kms/google_kms_key_ring_import_job",
		"kms/google_kms_secret_ciphertext",
		"ml_engine/google_ml_engine_model",
		"network_management/google_network_management_connectivity_test",
		"network_services/google_network_services_edge_cache_keyset",
		"network_services/google_network_services_edge_cache_origin",
		"network_services/google_network_services_edge_cache_service",
		"notebooks/google_notebooks_environment",
		"org_policy/google_org_policy_custom_constraint",
		"os_config/google_os_config_patch_deployment",
		"os_login/google_os_login_ssh_public_key",
		"pubsub_lite/google_pubsub_lite_subscription",
		"pubsub_lite/google_pubsub_lite_topic",
		"security_center/google_scc_notification_config",
		"security_center/google_scc_source",
		"service_usage/google_service_usage_consumer_quota_override",
		"storage/google_storage_hmac_key",
		"storage_transfer/google_storage_transfer_agent_pool",
		"tags/google_tags_location_tag_binding",
		"tpu/google_tpu_node",
		"vertex_ai/google_vertex_ai_dataset",
		"vertex_ai/google_vertex_ai_endpoint",
		"vertex_ai/google_vertex_ai_featurestore",
		"vertex_ai/google_vertex_ai_featurestore_entitytype",
		"vertex_ai/google_vertex_ai_featurestore_entitytype_feature",
		"vertex_ai/google_vertex_ai_index",
		"vertex_ai/google_vertex_ai_index_endpoint",
		"vertex_ai/google_vertex_ai_metadata_store",
		"vertex_ai/google_vertex_ai_tensorboard",
	}
	// betaAllowlist holds the list of the resources to be allowlisted as
	// v1beta1 CRDs. The format is '[terraform_product_name]/[terraform_type_name]'.
	// 'google_[terraform_product_name]' should be the prefix of
	// '[terraform_type_name]'.
	betaAllowlist = []string{
		"bigquery/google_bigquery_routine",
		"data_catalog/google_data_catalog_policy_tag",
		"data_catalog/google_data_catalog_taxonomy",
		"tags/google_tags_tag_binding",
		"tags/google_tags_tag_key",
		"tags/google_tags_tag_value",
	}
)

type AutoGenType struct {
	ServiceNameInLC string
	KRMKindName     string
	TFTypeName      string
	Version         string
}

func (a *AutoGenType) loadKRMKindFromSM(smAndRCMap map[string]map[string]string) error {
	service := a.ServiceNameInLC
	tfType := a.TFTypeName
	rcMap, ok := smAndRCMap[a.ServiceNameInLC]
	if !ok {
		return fmt.Errorf("can't find allowlisted service %v "+
			"in generated service mappings", service)
	}
	krmKind, ok := rcMap[tfType]
	if !ok {
		return fmt.Errorf("can't find allowlisted type %v "+
			"under service %v in auto-generated service mappings",
			tfType, service)
	}
	a.KRMKindName = krmKind
	return nil
}

func newAutoGenType(autoGenTypeInString string, version string) (*AutoGenType, error) {
	parts := strings.Split(autoGenTypeInString, "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("type for resource auto-generation should be"+
			" in the format '[terraform_product_name]/[terraform_type_name]', split by one '/',"+
			" but the provided type is %q", autoGenTypeInString)
	}

	if !text.IsSnakeCase(parts[0]) && !text.IsSnakeCase(parts[1]) {
		return nil, fmt.Errorf("type for resource auto-generation should be"+
			" in the format '[terraform_product_name]/[terraform_type_name]', both terraform_product_name"+
			" and terraform_type_name should be in snake case, but the provided"+
			" type is %q", autoGenTypeInString)
	}

	return &AutoGenType{
		ServiceNameInLC: strings.Replace(parts[0], "_", "", -1),
		TFTypeName:      parts[1],
		Version:         version,
	}, nil
}

type AutoGenAllowlist struct {
	ServiceAndTFTypes  map[string]map[string]*AutoGenType
	ServiceAndKRMKinds map[string]map[string]*AutoGenType
	KRMKinds           map[string]*AutoGenType
}

func (l *AutoGenAllowlist) HasService(serviceNameInLC string) bool {
	_, ok := l.ServiceAndTFTypes[serviceNameInLC]
	return ok
}

func (l *AutoGenAllowlist) GetTFTypeInService(serviceNameInLC, tfType string) (*AutoGenType, bool) {
	resourceMap, ok := l.ServiceAndTFTypes[serviceNameInLC]
	if !ok {
		return nil, false
	}
	autoGenType, ok := resourceMap[tfType]
	return autoGenType, ok
}

func (l *AutoGenAllowlist) GetKRMKind(krmKind string) (*AutoGenType, bool) {
	autoGenType, ok := l.KRMKinds[krmKind]
	return autoGenType, ok
}

func (l *AutoGenAllowlist) HasKRMKindInService(serviceNameInLC, krmKind string) bool {
	resourceMap, ok := l.ServiceAndKRMKinds[serviceNameInLC]
	if !ok {
		return false
	}
	_, ok = resourceMap[krmKind]
	return ok
}

func (l *AutoGenAllowlist) addAutoGenType(autoGenType *AutoGenType) error {
	_, ok := l.ServiceAndTFTypes[autoGenType.ServiceNameInLC]
	if !ok {
		l.ServiceAndTFTypes[autoGenType.ServiceNameInLC] = make(map[string]*AutoGenType)
		l.ServiceAndKRMKinds[autoGenType.ServiceNameInLC] = make(map[string]*AutoGenType)
	}
	TFTypeMap, _ := l.ServiceAndTFTypes[autoGenType.ServiceNameInLC]
	KRMKindMap, _ := l.ServiceAndKRMKinds[autoGenType.ServiceNameInLC]
	_, ok = TFTypeMap[autoGenType.TFTypeName]
	if ok {
		return fmt.Errorf("TF type %v has already been allowlisted under "+
			"service %v", autoGenType.TFTypeName, autoGenType.ServiceNameInLC)
	}

	TFTypeMap[autoGenType.TFTypeName] = autoGenType
	KRMKindMap[autoGenType.KRMKindName] = autoGenType
	l.KRMKinds[autoGenType.KRMKindName] = autoGenType
	return nil
}

func NewAutoGenAllowlist() *AutoGenAllowlist {
	return &AutoGenAllowlist{
		ServiceAndTFTypes:  make(map[string]map[string]*AutoGenType),
		ServiceAndKRMKinds: make(map[string]map[string]*AutoGenType),
		KRMKinds:           make(map[string]*AutoGenType),
	}
}

func LoadAutoGenAllowList(generatedSMMap map[string]v1alpha1.ServiceMapping) (*AutoGenAllowlist, error) {
	smAndRCMap := getGeneratedSMAndRCMap(generatedSMMap)
	autoGenAllowlist := NewAutoGenAllowlist()
	for _, typeInString := range alphaAllowlist {
		autoGenType, err := newAutoGenType(typeInString, k8s.KCCAPIVersionV1Alpha1)
		if err != nil {
			return nil, fmt.Errorf("error converting allowlisted type %v from string to AutoGenType: %w", typeInString, err)
		}
		if err := autoGenType.loadKRMKindFromSM(smAndRCMap); err != nil {
			return nil, fmt.Errorf("error loading KRMKind for allowlisted type %v: %w", typeInString, err)
		}
		if err := autoGenAllowlist.addAutoGenType(autoGenType); err != nil {
			return nil, fmt.Errorf("error adding AutoGenType for %v into the AutoGenAllowlist: %w", typeInString, err)
		}
	}
	for _, typeInString := range betaAllowlist {
		autoGenType, err := newAutoGenType(typeInString, k8s.KCCAPIVersionV1Beta1)
		if err != nil {
			return nil, fmt.Errorf("error converting allowlisted type %v from string to AutoGenType: %w", typeInString, err)
		}
		if err := autoGenType.loadKRMKindFromSM(smAndRCMap); err != nil {
			return nil, fmt.Errorf("error loading KRMKind for allowlisted type %v: %w", typeInString, err)
		}
		if err := autoGenAllowlist.addAutoGenType(autoGenType); err != nil {
			return nil, fmt.Errorf("error adding AutoGenType for %v into the AutoGenAllowlist: %w", typeInString, err)
		}
	}
	return autoGenAllowlist, nil
}

func getGeneratedSMAndRCMap(generatedSMMap map[string]v1alpha1.ServiceMapping) map[string]map[string]string {
	smAndRCMap := make(map[string]map[string]string)
	for smName, sm := range generatedSMMap {
		service := strings.TrimSuffix(smName, ".cnrm.cloud.google.com")
		generatedRCMap := make(map[string]string)
		for _, rc := range sm.Spec.Resources {
			generatedRCMap[rc.Name] = rc.Kind
		}
		smAndRCMap[service] = generatedRCMap
	}
	return smAndRCMap
}
