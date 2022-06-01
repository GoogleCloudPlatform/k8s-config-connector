## 4.23.0 (Unreleased)

FEATURES:
* **New Data Source:** `google_tags_tag_key` ([#4337](https://github.com/hashicorp/terraform-provider-google-beta/pull/4337))
* **New Data Source:** `google_tags_tag_value` ([#4337](https://github.com/hashicorp/terraform-provider-google-beta/pull/4337))
* **New Resource:** `google_dataplex_lake` ([#4341](https://github.com/hashicorp/terraform-provider-google-beta/pull/4341))

IMPROVEMENTS:
* bigqueryconnection: updated connection types to support v1 ga ([#4323](https://github.com/hashicorp/terraform-provider-google-beta/pull/4323))
* cloudfunctions: added docker registry support for Cloud Functions ([#4324](https://github.com/hashicorp/terraform-provider-google-beta/pull/4324))
* memcache: added `maintenance_policy` and `maintenance_schedule` to `google_memcache_instance` ([#4338](https://github.com/hashicorp/terraform-provider-google-beta/pull/4338))
* service-directory: marked network field immutable in `google_service_directory_endpoint` ([#4334](https://github.com/hashicorp/terraform-provider-google-beta/pull/4334))

BUG FIXES:
* binaryauthorization: fixed permadiff in `google_binary_authorization_attestor` ([#4325](https://github.com/hashicorp/terraform-provider-google-beta/pull/4325))
* service: added re-polling for service account after creation, 404s sometimes due to [eventual consistency](https://cloud.google.com/iam/docs/overview#consistency) ([#4333](https://github.com/hashicorp/terraform-provider-google-beta/pull/4333))

## 4.22.0 (May 24, 2022)

NOTE: Due to technical difficulties encountered in the release process, the `4.22.0` release for `google-beta` occurred several hours after the corresponding `google` provider release.

FEATURES:
* **New Resource:** `google_certificate_manager_certificate` ([#4301](https://github.com/hashicorp/terraform-provider-google-beta/pull/4301))
* **New Resource:** `google_certificate_manager_dns_authorization` ([#4301](https://github.com/hashicorp/terraform-provider-google-beta/pull/4301))
* **New Resource:** `google_clouddeploy_delivery_pipeline` ([#4288](https://github.com/hashicorp/terraform-provider-google-beta/pull/4288))
* **New Resource:** `google_clouddeploy_target` ([#4288](https://github.com/hashicorp/terraform-provider-google-beta/pull/4288))

IMPROVEMENTS:
* bigquery: added connection of type cloud_resource for `google_bigquery_connection` ([#4312](https://github.com/hashicorp/terraform-provider-google-beta/pull/4312))
* cloudfunctions: added `https_trigger_security_level` to `google_cloudfunctions_function` ([#4295](https://github.com/hashicorp/terraform-provider-google-beta/pull/4295))
* cloudrun: added `traffic.tag` and `traffic.url` fields to `google_cloud_run_service` ([#4283](https://github.com/hashicorp/terraform-provider-google-beta/pull/4283))
* compute: added `enable_dynamic_port_allocation` to `google_compute_router_nat` ([#4316](https://github.com/hashicorp/terraform-provider-google-beta/pull/4316))
* compute: added field `update_policy.most_disruptive_allowed_action` to `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` ([#4282](https://github.com/hashicorp/terraform-provider-google-beta/pull/4282))
* compute: added support for NEG type `PRIVATE_SERVICE_CONNECT` in `NetworkEndpointGroup` ([#4303](https://github.com/hashicorp/terraform-provider-google-beta/pull/4303))
* compute: added support for `domain_names` attribute in `google_compute_service_attachment` ([#4313](https://github.com/hashicorp/terraform-provider-google-beta/pull/4313))
* compute: added value `REFRESH` to field update_policy.minimal_action` in `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` ([#4282](https://github.com/hashicorp/terraform-provider-google-beta/pull/4282))
* container: added field `exclusion_options` to `google_container_cluster` ([#4291](https://github.com/hashicorp/terraform-provider-google-beta/pull/4291))
* monitoring: added `checker_type` field to `google_monitoring_uptime_check_config` resource ([#4302](https://github.com/hashicorp/terraform-provider-google-beta/pull/4302))
* privateca: add a new field `desired_state` to manage CertificateAuthority state. ([#4279](https://github.com/hashicorp/terraform-provider-google-beta/pull/4279))
* sql: added `active_directory_config` field in `google_sql_database_instance` ([#4298](https://github.com/hashicorp/terraform-provider-google-beta/pull/4298))
* sql: removed requirement that Cloud SQL Insight is only allowed for Postgres in `google_sql_database_instance` ([#4310](https://github.com/hashicorp/terraform-provider-google-beta/pull/4310))

BUG FIXES:
* cloudfunctions: fixed an issue where `google_cloudfunctions2_function` would not update ([#4278](https://github.com/hashicorp/terraform-provider-google-beta/pull/4278))
* compute: fixed extra diffs generated on `google_security_policy` `rules` when modifying a rule ([#4287](https://github.com/hashicorp/terraform-provider-google-beta/pull/4287))
* container: fixed Autopilot cluster couldn't omit master ipv4 cidr in `google_container_cluster` ([#4280](https://github.com/hashicorp/terraform-provider-google-beta/pull/4280))
* resourcemanager: fixed a bug in wrongly writing to state when creation failed on `google_project_organization_policy` ([#4297](https://github.com/hashicorp/terraform-provider-google-beta/pull/4297))
* storage: not specifying `content` or `source` for `google_storage_bucket_object` now fails at plan-time instead of apply-time. ([#4292](https://github.com/hashicorp/terraform-provider-google-beta/pull/4292))

## 4.21.0 (May 16, 2022)

IMPROVEMENTS:
* cloudfunctions: added CMEK support for Cloud Functions ([#4272](https://github.com/hashicorp/terraform-provider-google-beta/pull/4272))
* compute: added `service_directory_registrations` to `google_compute_forwarding_rule` resource ([#4276](https://github.com/hashicorp/terraform-provider-google-beta/pull/4276))
* compute: removed validation checking against a fixed set of persistent disk types ([#4273](https://github.com/hashicorp/terraform-provider-google-beta/pull/4273))
* container: removed validation checking against a fixed set of persistent disk types ([#4273](https://github.com/hashicorp/terraform-provider-google-beta/pull/4273))
* containeraws: added `image_type` and `instance_placement` to `google_container_aws_node_pool` resource ([#4276](https://github.com/hashicorp/terraform-provider-google-beta/pull/4276))
* containeraws: added `instance_placement` and `logging_config` to `google_container_aws_cluster` resource ([#4276](https://github.com/hashicorp/terraform-provider-google-beta/pull/4276))
* containeraws: added `proxy_config` to `google_container_aws_node_pool` resource ([#4276](https://github.com/hashicorp/terraform-provider-google-beta/pull/4276))
* containerazure: added `image_type` to `google_container_azure_node_pool` resource ([#4276](https://github.com/hashicorp/terraform-provider-google-beta/pull/4276))
* containerazure: added `logging_config` to `google_container_azure_cluster` resource ([#4276](https://github.com/hashicorp/terraform-provider-google-beta/pull/4276))
* containerazure: added `proxy_config` to `google_container_azure_node_pool` resource ([#4276](https://github.com/hashicorp/terraform-provider-google-beta/pull/4276))
* dataproc: removed validation checking against a fixed set of persistent disk types ([#4273](https://github.com/hashicorp/terraform-provider-google-beta/pull/4273))
* dns: added `routing_policy` to `google_dns_record_set` resource ([#4265](https://github.com/hashicorp/terraform-provider-google-beta/pull/4265))

BUG FIXES:
* cloudfunctions: fixed an issue where `google_cloudfunctions2_function` would not update ([#4278](https://github.com/hashicorp/terraform-provider-google-beta/pull/4278))
* compute: fixed a crash in `google_compute_instance` when the instance is deleted outside of Terraform ([#4262](https://github.com/hashicorp/terraform-provider-google-beta/pull/4262))
* provider: removed printing credentials to the console if malformed JSON is given ([#4266](https://github.com/hashicorp/terraform-provider-google-beta/pull/4266))

## 4.20.0 (May 2, 2022)

NOTES:
* `google_privateca_certificate_authority` resources now cannot be destroyed unless `deletion_protection = false` is set in state for the resource. ([#4241](https://github.com/hashicorp/terraform-provider-google-beta/pull/4241))

FEATURES:
* **New Data Source:** `google_compute_disk` ([#4255](https://github.com/hashicorp/terraform-provider-google-beta/pull/4255))

IMPROVEMENTS:
* apigee: `consumer_accept_list` and `service_attachment` to `google_apigee_instance`. ([#4260](https://github.com/hashicorp/terraform-provider-google-beta/pull/4260))
* compute: added `subsetting` field to `google_compute_region_backend_service` ([#4246](https://github.com/hashicorp/terraform-provider-google-beta/pull/4246))
* privateca: added `deletion_protection` for `google_privateca_certificate_authority`. ([#4241](https://github.com/hashicorp/terraform-provider-google-beta/pull/4241))
* privateca: added new output fields on `google_privateca_certificate` including `issuer_certificate_authority`, `pem_certificate_chain` and `certificate_description.x509_description` ([#4242](https://github.com/hashicorp/terraform-provider-google-beta/pull/4242))
* redis: added multi read replica field `read_replicas_mode` and `secondary_ip_range` in `google_redis_instance` ([#4259](https://github.com/hashicorp/terraform-provider-google-beta/pull/4259))

BUG FIXES:
* compute: fixed a crash when `compute.instance` is not found ([#4262](https://github.com/hashicorp/terraform-provider-google-beta/pull/4262))
* provider: removed printing credentials to the console if malformed JSON is given ([#4266](https://github.com/hashicorp/terraform-provider-google-beta/pull/4266))
* sql: fixed bug where `encryption_key_name` was not being propagated to the API. ([#4261](https://github.com/hashicorp/terraform-provider-google-beta/pull/4261))

## 4.19.0 (April 25, 2022)

IMPROVEMENTS:
* cloudbuild: made `CLOUD_LOGGING_ONLY` available as a cloud build logging option. ([#4224](https://github.com/hashicorp/terraform-provider-google-beta/pull/4224))
* compute: added `redirect_options` field for `google_compute_security_policy` rules ([#4217](https://github.com/hashicorp/terraform-provider-google-beta/pull/4217))
* compute: added `FIXED_STANDARD` and `STANDARD` as valid values to the field `network_interface.0.access_configs.0.network_tier` of  `google_compute_instance_template` resource ([#4233](https://github.com/hashicorp/terraform-provider-google-beta/pull/4233))
* compute: added `FIXED_STANDARD` and `STANDARD` as valid values to the field `network_interface.0.access_configs.0.network_tier` of  `google_compute_instance` resource ([#4233](https://github.com/hashicorp/terraform-provider-google-beta/pull/4233))
* compute: added passing `exceed_redirect_options` field for `google_compute_security_policy` rules ([#4238](https://github.com/hashicorp/terraform-provider-google-beta/pull/4238))
* container: added `gke_backup_agent_config` in `addons_config` to `google_container_cluster` (beta) ([#4231](https://github.com/hashicorp/terraform-provider-google-beta/pull/4231))
* filestore: added `kms_key_name` field to `google_filestore_instance` resource to support CMEK ([#11493](https://github.com/hashicorp/terraform-provider-google/pull/11493))
* logging: made `google_logging_*_bucket_config` deletable ([#4234](https://github.com/hashicorp/terraform-provider-google-beta/pull/4234))
* notebooks: updated `container_images` on `google_notebooks_runtime` to default to the value returned by the API if not set ([#4216](https://github.com/hashicorp/terraform-provider-google-beta/pull/4216))
* provider: modified request retry logic to retry all per-minute quota limits returned with a 403 error code. Previously, only read requests were retried. This will generally affect Google Compute Engine resources. ([#4223](https://github.com/hashicorp/terraform-provider-google-beta/pull/4223))

BUG FIXES:
* bigquery: fixed a bug where `encryption_configuration.kms_key_name` stored the version rather than the key name. ([#4221](https://github.com/hashicorp/terraform-provider-google-beta/pull/4221))
* compute: fixed url_mask required mis-annotation in `google_compute_region_network_endpoint_group`, making it optional ([#4227](https://github.com/hashicorp/terraform-provider-google-beta/pull/4227))
* spanner: fixed escaping of database names with Postgres dialect in `google_spanner_database` ([#4228](https://github.com/hashicorp/terraform-provider-google-beta/pull/4228))

## 4.18.0 (April 18, 2022)

FEATURES:
* **New Resource:** `google_privateca_certificate_template_iam_binding` ([#4201](https://github.com/hashicorp/terraform-provider-google-beta/pull/4201))
* **New Resource:** `google_privateca_certificate_template_iam_member` ([#4201](https://github.com/hashicorp/terraform-provider-google-beta/pull/4201))
* **New Resource:** `google_privateca_certificate_template_iam_policy` ([#4201](https://github.com/hashicorp/terraform-provider-google-beta/pull/4201))

IMPROVEMENTS:
* bigtable: added `gc_rules` to `google_bigtable_gc_policy` resource. ([#4212](https://github.com/hashicorp/terraform-provider-google-beta/pull/4212))
* dialogflow: added support for location based dialogflow resources ([#4206](https://github.com/hashicorp/terraform-provider-google-beta/pull/4206))
* metastore: added support for encryption_config during service creation. ([#4204](https://github.com/hashicorp/terraform-provider-google-beta/pull/4204))
* privateca: support update on CertificateAuthority and Certificate ([#4207](https://github.com/hashicorp/terraform-provider-google-beta/pull/4207))

BUG FIXES:
* Update mutex on google_apigee_instance_attachment to lock on org_id. ([#4203](https://github.com/hashicorp/terraform-provider-google-beta/pull/4203))
* vpcaccess: fixed an issue where `google_vpc_access_connector` would be repeatedly recreated when `network` was not specified ([#4205](https://github.com/hashicorp/terraform-provider-google-beta/pull/4205))

## 4.17.0 (April 11, 2022)

FEATURES:
* **New Data Source:** `google_access_approval_folder_service_account` ([#4179](https://github.com/hashicorp/terraform-provider-google-beta/pull/4179))
* **New Data Source:** `google_access_approval_organization_service_account` ([#4179](https://github.com/hashicorp/terraform-provider-google-beta/pull/4179))
* **New Data Source:** `google_access_approval_project_service_account` ([#4179](https://github.com/hashicorp/terraform-provider-google-beta/pull/4179))
* **New Resource:** `google_access_context_manager_access_policy_iam_binding` ([#4180](https://github.com/hashicorp/terraform-provider-google-beta/pull/4180))
* **New Resource:** `google_access_context_manager_access_policy_iam_member` ([#4180](https://github.com/hashicorp/terraform-provider-google-beta/pull/4180))
* **New Resource:** `google_access_context_manager_access_policy_iam_policy` ([#4180](https://github.com/hashicorp/terraform-provider-google-beta/pull/4180))
* **New Resource:** `google_endpoints_service_consumers_iam_binding` ([#4160](https://github.com/hashicorp/terraform-provider-google-beta/pull/4160))
* **New Resource:** `google_endpoints_service_consumers_iam_member` ([#4160](https://github.com/hashicorp/terraform-provider-google-beta/pull/4160))
* **New Resource:** `google_endpoints_service_consumers_iam_policy` ([#4160](https://github.com/hashicorp/terraform-provider-google-beta/pull/4160))
* **New Resource:** `google_iam_deny_policy` ([#4194](https://github.com/hashicorp/terraform-provider-google-beta/pull/4194))

IMPROVEMENTS:
* access approval: added `active_key_version`, `ancestor_has_active_key_version`, and `invalid_key_version` fields to `google_folder_access_approval_settings`, `google_organization_access_approval_settings`, and `google_project_access_approval_settings` resources ([#4179](https://github.com/hashicorp/terraform-provider-google-beta/pull/4179))
* access context manager: added support for scoped policies in `google_access_context_manager_access_policy` ([#4180](https://github.com/hashicorp/terraform-provider-google-beta/pull/4180))
* apigee: added `deployment_type` and `api_proxy_type` to `google_apigee_environment` ([#4177](https://github.com/hashicorp/terraform-provider-google-beta/pull/4177))
* bigtable: updated the examples to show users can create all 3 different flavors of AppProfile ([#4172](https://github.com/hashicorp/terraform-provider-google-beta/pull/4172))
* cloudbuild: added `approval_config` to `google_cloudbuild_trigger` ([#4162](https://github.com/hashicorp/terraform-provider-google-beta/pull/4162))
* composer: added support for `airflow-1` and `airflow-2` aliases in image version argument ([#4185](https://github.com/hashicorp/terraform-provider-google-beta/pull/4185))
* dataflow: added `skip_wait_on_job_termination` attribute to `google_dataflow_job` and `google_dataflow_flex_template_job` resources (issue #10559) ([#4196](https://github.com/hashicorp/terraform-provider-google-beta/pull/4196))
* dataproc: added `presto_config` to `dataproc_job` ([#4171](https://github.com/hashicorp/terraform-provider-google-beta/pull/4171))
* healthcare: added support V3 parser version for Healthcare HL7 stores. ([#4189](https://github.com/hashicorp/terraform-provider-google-beta/pull/4189))
* healthcare: added support for `ANALYTICS_V2 `and `LOSSLESS` BigQueryDestination schema types to `google_healthcare_fhir_store` ([#4186](https://github.com/hashicorp/terraform-provider-google-beta/pull/4186))
* os-config: added field `migInstancesAllowed` to resource `os_config_patch_deployment` ([#4195](https://github.com/hashicorp/terraform-provider-google-beta/pull/4195))
* privateca: added support for IAM conditions to CaPool ([#4170](https://github.com/hashicorp/terraform-provider-google-beta/pull/4170))
* pubsub: added `enable_exactly_once_delivery` to `google_pubsub_subscription` ([#4166](https://github.com/hashicorp/terraform-provider-google-beta/pull/4166))
* spanner: added support for setting database_dialect on `google_spanner_database` ([#4158](https://github.com/hashicorp/terraform-provider-google-beta/pull/4158))

BUG FIXES:
* redis: fixed an issue where older redis instances had a dangerous diff on the field `read_replicas_mode`, adding a default of `READ_REPLICAS_DISABLED`. Now, if the field is not set in config, the value of the field will keep the old value from state. ([#4184](https://github.com/hashicorp/terraform-provider-google-beta/pull/4184))
* tags: fixed issue where tags could not be applied sequentially to the same parent in `google_tags_tag_binding` ([#4191](https://github.com/hashicorp/terraform-provider-google-beta/pull/4191))

## 4.16.0 (April 4, 2022)

FEATURES:
* **New Data Source:** `google_dataproc_metastore_service` ([#4155](https://github.com/hashicorp/terraform-provider-google-beta/pull/4155))
* **New Resource:** `google_firebaserules_release` ([#4132](https://github.com/hashicorp/terraform-provider-google-beta/pull/4132))
* **New Resource:** `google_firebaserules_ruleset` ([#4132](https://github.com/hashicorp/terraform-provider-google-beta/pull/4132))

IMPROVEMENTS:
* bigtable: added support for `autoscaling_config` to `google_bigtable_instance` ([#4150](https://github.com/hashicorp/terraform-provider-google-beta/pull/4150))
* composer: Added support for `composer-1` and `composer-2` aliases in image version argument ([#4131](https://github.com/hashicorp/terraform-provider-google-beta/pull/4131))
* compute: added support for attaching a `edge_security_policy` to `google_compute_backend_bucket` ([#4154](https://github.com/hashicorp/terraform-provider-google-beta/pull/4154))
* compute: added support for field `type` to `google_compute_security_policy` ([#4154](https://github.com/hashicorp/terraform-provider-google-beta/pull/4154))
* eventarc: added gke and workflows destination for eventarc trigger resource. ([#4152](https://github.com/hashicorp/terraform-provider-google-beta/pull/4152))
* networkservices: added `included_cookie_names` to cache key policy configuration ([#4147](https://github.com/hashicorp/terraform-provider-google-beta/pull/4147))
* spanner: added support for setting database_dialect on `google_spanner_database` ([#4158](https://github.com/hashicorp/terraform-provider-google-beta/pull/4158))
* storagetransfer: added `repeat_interval` field to `google_storage_transfer_job` resource ([#4144](https://github.com/hashicorp/terraform-provider-google-beta/pull/4144))

BUG FIXES:
* apikeys: fixed a bug where `google_apikeys_key.key_string` was not being set. ([#4139](https://github.com/hashicorp/terraform-provider-google-beta/pull/4139))
* container: fixed a bug where `google_container_cluster.authenticator_groups_config` could not be set in tandem with `enable_autopilot` ([#4140](https://github.com/hashicorp/terraform-provider-google-beta/pull/4140))
* iam: fixed an issue where special identifiers `allAuthenticatedUsers` and `allUsers` were flattened to lower case in IAM members. ([#4156](https://github.com/hashicorp/terraform-provider-google-beta/pull/4156))
* logging: fixed bug where `google_logging_project_bucket_config` would erroneously write to state after it errored out and wasn't actually created. ([#4141](https://github.com/hashicorp/terraform-provider-google-beta/pull/4141))
* monitoring: fixed a permadiff when `google_monitoring_uptime_check_config.http_check.path` does not begin with "/" ([#4135](https://github.com/hashicorp/terraform-provider-google-beta/pull/4135))
* osconfig: fixed a bug where `recurring_schedule.time_of_day` can not be set to 12am exact time in `google_os_config_patch_deployment` resource ([#4127](https://github.com/hashicorp/terraform-provider-google-beta/pull/4127))
* sql: fixed bug where permadiff of `encryption_key_name` would show on `google_sql_database_instance` for replica instances. ([#4130](https://github.com/hashicorp/terraform-provider-google-beta/pull/4130))
* storage: fixed a bug where `google_storage_bucket` data source would retry for 20 min when bucket was not found. ([#4129](https://github.com/hashicorp/terraform-provider-google-beta/pull/4129))
* storage: fixed bug where `google_storage_transfer_job` that was deleted outside of Terraform would not be recreated on apply. ([#4138](https://github.com/hashicorp/terraform-provider-google-beta/pull/4138))


## 4.15.0 (March 21, 2022)

FEATURES:
* **New Resource:** google_logging_log_view ([#4125](https://github.com/hashicorp/terraform-provider-google-beta/pull/4125))

IMPROVEMENTS:
* apigee: added `billing_type` attribute to `google_apigee_organization` resource. ([#4126](https://github.com/hashicorp/terraform-provider-google-beta/pull/4126))
* networkservices: added `disable_http2` property to `google_network_services_edge_cache_service` resource ([#4119](https://github.com/hashicorp/terraform-provider-google-beta/pull/4119))
* networkservices: updated `google_network_services_edge_cache_origin` resource to read and write the `timeout` property, including a new `read_timeout` field. ([#4122](https://github.com/hashicorp/terraform-provider-google-beta/pull/4122))
* networkservices: updated `google_network_services_edge_cache_origin` to retry_conditions to include `FORBIDDEN` ([#4122](https://github.com/hashicorp/terraform-provider-google-beta/pull/4122))

BUG FIXES:
* dataproc: fixed a crash when `logging_config` only contains `nil` entry  in `google_dataproc_workflow_template` ([#4124](https://github.com/hashicorp/terraform-provider-google-beta/pull/4124))
* sql: fixed crash when one of `settings.database_flags` is nil. ([#4123](https://github.com/hashicorp/terraform-provider-google-beta/pull/4123))

## 4.14.0 (March 14, 2022)

FEATURES:
* **New Resource:** `google_bigqueryreservation_assignment` ([#4098](https://github.com/hashicorp/terraform-provider-google-beta/pull/4098))
* **New Resource:** `google_apikeys_key` ([#4114](https://github.com/hashicorp/terraform-provider-google-beta/pull/4114))

IMPROVEMENTS:
* artifactregistry: added maven config for `google_artifact_registry_repository` ([#4112](https://github.com/hashicorp/terraform-provider-google-beta/pull/4112))
* cloudbuild: added support for manual builds, git source for webhook/pubsub triggered builds and filter field ([#4100](https://github.com/hashicorp/terraform-provider-google-beta/pull/4100))
* container: added support for gvnic to `google_container_node_pool` ([#4111](https://github.com/hashicorp/terraform-provider-google-beta/pull/4111))
* dataproc: added `preemptibility` field to the `preemptible_worker_config` of `google_dataproc_cluster` ([#4107](https://github.com/hashicorp/terraform-provider-google-beta/pull/4107))
* serviceusage: supported `force` behavior for deleting consumer quota override ([#4094](https://github.com/hashicorp/terraform-provider-google-beta/pull/4094))

BUG FIXES:
* dataproc: fixed a crash when `logging_config` only contains `nil` entry  in `google_dataproc_job` ([#4108](https://github.com/hashicorp/terraform-provider-google-beta/pull/4108))

## 4.13.0 (March 7, 2022)

FEATURES:
* **New Resource:** `google_apigee_endpoint_attachment` ([#4074](https://github.com/hashicorp/terraform-provider-google-beta/pull/4074))
* **New Resource:** `google_cloudfunctions2_function` ([#4093](https://github.com/hashicorp/terraform-provider-google-beta/pull/4093))
* **New Resource:** `google_region_backend_service_iam_*` ([#4088](https://github.com/hashicorp/terraform-provider-google-beta/pull/4088))
* **New Datasource:** `google_dns_record_set` ([#4085](https://github.com/hashicorp/terraform-provider-google-beta/pull/4085))
* **New Datasource:** `google_privateca_certificate_authority` ([#4087](https://github.com/hashicorp/terraform-provider-google-beta/pull/4087))

IMPROVEMENTS:
* compute: added support for `keepalive_interval` to `google_compute_router.bgp` ([#4089](https://github.com/hashicorp/terraform-provider-google-beta/pull/4089))
* compute: added update support for `google_compute_reservation.share_settings` ([#4092](https://github.com/hashicorp/terraform-provider-google-beta/pull/4092))
* storagetransfer: added attribute `subject_id` to data source `google_storage_transfer_project_service_account` ([#4073](https://github.com/hashicorp/terraform-provider-google-beta/pull/4073))

BUG FIXES:
* composer: allow region to be undefined in configuration for `google_composer_environment` ([#4083](https://github.com/hashicorp/terraform-provider-google-beta/pull/4083))
* container: fixed a bug where `vertical_pod_autoscaling` would cause autopilot clusters to recreate ([#4076](https://github.com/hashicorp/terraform-provider-google-beta/pull/4076))

## 4.12.0 (February 28, 2022)

NOTE:
* updated to go 1.16.14 ([#4066](https://github.com/hashicorp/terraform-provider-google-beta/pull/4066))

FEATURES:
* **New Resource:** `dns_response_policy*` ([#4046](https://github.com/hashicorp/terraform-provider-google-beta/pull/4046))
* **New Resource:** `dns_response_policy_rule*` ([#4046](https://github.com/hashicorp/terraform-provider-google-beta/pull/4046))

DEPRECATIONS:
* datafusion: deprecated `service_account` in `google_datafusion_instance`. Use `tenant_project_id` instead to extract the tenant project ID (beta) ([#4045](https://github.com/hashicorp/terraform-provider-google-beta/pull/4045))

IMPROVEMENTS:
* bigquery: added support for authorized datasets to `google_bigquery_dataset.access` and `google_bigquery_dataset_access` ([#4047](https://github.com/hashicorp/terraform-provider-google-beta/pull/4047))
* bigtable: added `multi_cluster_routing_cluster_ids` fields to `google_bigtable_app_profile` ([#4051](https://github.com/hashicorp/terraform-provider-google-beta/pull/4051))
* compute: added field `serverless_deployment` to `google_compute_network_endpoint_group` (beta only) for API Gateway resources ([#4041](https://github.com/hashicorp/terraform-provider-google-beta/pull/4041))
* compute: updated `instance` attribute for `google_compute_network_endpoint` to be optional, as Hybrid connectivity NEGs use network endpoints with just IP and Port. ([#4068](https://github.com/hashicorp/terraform-provider-google-beta/pull/4068))
* compute: added `NON_GCP_PRIVATE_IP_PORT` value for `network_endpoint_type` in the `google_compute_network_endpoint_group` resource ([#4068](https://github.com/hashicorp/terraform-provider-google-beta/pull/4068))
* compute: added `provisioning_model` field to `google_compute_instance_template ` resource to support Spot VM(beta) ([#4033](https://github.com/hashicorp/terraform-provider-google-beta/pull/4033))
* compute: added `provisioning_model` field to `google_compute_instance` resource to support Spot VM(beta) ([#4033](https://github.com/hashicorp/terraform-provider-google-beta/pull/4033))
* container: Add support for GKE Compact Placement ([#4043](https://github.com/hashicorp/terraform-provider-google-beta/pull/4043))
* datafusion: added support for `tenant_project_id` and `gcs_bucket` in `google_datafusion_instance` resource. ([#4045](https://github.com/hashicorp/terraform-provider-google-beta/pull/4045))
* provider: added retries for `ReadRequest` errors incorrectly coded as `403` errors, particularly in Google Compute Engine ([#4064](https://github.com/hashicorp/terraform-provider-google-beta/pull/4064))

BUG FIXES:
* apigee: fixed a bug where multiple `google_apigee_instance` could not be used on the same `google_apigee_organization` ([#4059](https://github.com/hashicorp/terraform-provider-google-beta/pull/4059))
* compute: corrected an issue in `google_compute_security_policy` where only alpha values for certain enums were accepted ([#4049](https://github.com/hashicorp/terraform-provider-google-beta/pull/4049))
* compute: fixed permadiff in `google_compute_instance.scheduling.provisioning_model` ([#4044](https://github.com/hashicorp/terraform-provider-google-beta/pull/4044))
* compute: fixed permadiff in `google_compute_instance_template.scheduling.provisioning_model` ([#4052](https://github.com/hashicorp/terraform-provider-google-beta/pull/4052))

## 4.11.0 (February 16, 2022)

IMPROVEMENTS:
* cloudfunctions: Added SecretManager integration support to `google_cloudfunctions_function`. ([#4040](https://github.com/hashicorp/terraform-provider-google-beta/pull/4040))
* compute: Added field `serverless_deployment` to `google_compute_network_endpoint_group` ([#4041](https://github.com/hashicorp/terraform-provider-google-beta/pull/4041))
* dataproc: increased the default timeout for `google_dataproc_cluster` from 20m to 45m ([#4027](https://github.com/hashicorp/terraform-provider-google-beta/pull/4027))
* sql: added field `clone.allocated_ip_range` to support address range picker for clone in resource `google_sql_database_instance` ([#4037](https://github.com/hashicorp/terraform-provider-google-beta/pull/4037))
* storagetransfer: added support for POSIX data source and data sink to `google_storage_transfer_job` via `transfer_spec.posix_data_source` and `transfer_spec.posix_data_sink` fields ([#4029](https://github.com/hashicorp/terraform-provider-google-beta/pull/4029))

BUG FIXES:
* cloudrun: updated `containers.ports.container_port` to be optional instead of required on `google_cloud_run_service` ([#4030](https://github.com/hashicorp/terraform-provider-google-beta/pull/4030))
* compute: marked `project` field optional in `google_compute_instance_template` data source ([#4031](https://github.com/hashicorp/terraform-provider-google-beta/pull/4031))

## 4.10.0 (February 7, 2022)

FEATURES:
* **New Resource:** `google_backend_service_iam_*` ([#4021](https://github.com/hashicorp/terraform-provider-google-beta/pull/4021))

IMPROVEMENTS:
* compute: added `EXTERNAL_MANAGED` as option for `load_balancing_scheme` in `google_compute_global_forwarding_rule` resource ([#4011](https://github.com/hashicorp/terraform-provider-google-beta/pull/4011))
* compute: added field `rate_limit_options` to `google_compute_security_policy` rules ([#4020](https://github.com/hashicorp/terraform-provider-google-beta/pull/4020))
* container: added support for image type configuration on the GKE Node Auto-provisioning ([#4023](https://github.com/hashicorp/terraform-provider-google-beta/pull/4023))
* container: added support for GCPFilestoreCSIDriver addon to `google_container_cluster` resource. ([#4015](https://github.com/hashicorp/terraform-provider-google-beta/pull/4015))
* dataproc: increased the default timeout for `google_dataproc_cluster` from 20m to 45m ([#4027](https://github.com/hashicorp/terraform-provider-google-beta/pull/4027))
* redis: added `maintenance_policy` and `maintenance_schedule` to `google_redis_instance` ([#4010](https://github.com/hashicorp/terraform-provider-google-beta/pull/4010))
* vpcaccess: updated field `network` in `google_vpc_access_connector` to accept `self_link` or `name` ([#4013](https://github.com/hashicorp/terraform-provider-google-beta/pull/4013))

BUG FIXES:
* storage: fixed bug where the provider crashes when `Object.owner` is missing when using `google_storage_object_acl` ([#4019](https://github.com/hashicorp/terraform-provider-google-beta/pull/4019))

## 4.9.0 (January 31, 2022)

BREAKING CHANGES:
* cloudrun: changed the `location` of `google_cloud_run_service` so that modifying the `location` field will recreate the resource rather than causing Terraform to report it would attempt an invalid update ([#3998](https://github.com/hashicorp/terraform-provider-google-beta/pull/3998))

IMPROVEMENTS:
* provider: changed the default timeout for many resources to 20 minutes, the current Terraform default, where it was less than 20 minutes previously ([#4002](https://github.com/hashicorp/terraform-provider-google-beta/pull/4002))
* redis: added `maintenance_policy` and `maintenance_schedule` to `google_redis_instance` ([#4010](https://github.com/hashicorp/terraform-provider-google-beta/pull/4010))
* storage: added field `transfer_spec.aws_s3_data_source.role_arn` to `google_storage_transfer_job` ([#3999](https://github.com/hashicorp/terraform-provider-google-beta/pull/3999))

BUG FIXES:
* cloudrun: fixed a bug where changing the non-updatable `location` of a `google_cloud_run_service` would not force resource recreation ([#3998](https://github.com/hashicorp/terraform-provider-google-beta/pull/3998))
* compute: fixed a bug where `google_compute_firewall` would incorrectly find `source_ranges` to be empty during validation ([#4008](https://github.com/hashicorp/terraform-provider-google-beta/pull/4008))
* notebooks: fixed permadiff in `google_notebooks_runtime.software_config` ([#3997](https://github.com/hashicorp/terraform-provider-google-beta/pull/3997))

## 4.8.0 (January 24, 2022)

BREAKING CHANGES:
* dlp: renamed the `characters_to_ignore.character_to_skip` field to `characters_to_ignore.characters_to_skip` in `google_data_loss_prevention_deidentify_template`. Any affected configurations will have been failing with an error at apply time already. ([#3983](https://github.com/hashicorp/terraform-provider-google-beta/pull/3983))

FEATURES:
* **New Resource:** `google_network_connectivity_spoke` ([#3987](https://github.com/hashicorp/terraform-provider-google-beta/pull/3987))

IMPROVEMENTS:
* apigee: added `ip_range` field to `google_apigee_instance` ([#3989](https://github.com/hashicorp/terraform-provider-google-beta/pull/3989))
* cloudrun: added support for `default_mode` and `mode` settings for created files within `secrets` in `google_cloud_run_service` ([#3984](https://github.com/hashicorp/terraform-provider-google-beta/pull/3984))
* compute: Added `share_settings` in `google_compute_reservation` ([#3980](https://github.com/hashicorp/terraform-provider-google-beta/pull/3980))

BUG FIXES:
* all: Fixed operation polling to support custom endpoints. ([#3986](https://github.com/hashicorp/terraform-provider-google-beta/pull/3986))
* cloudrun: Fixed permadiff in `google_cloud_run_service`'s `template.spec.service_account_name`. ([#3993](https://github.com/hashicorp/terraform-provider-google-beta/pull/3993))
* dlp: Fixed typo in name of `characters_to_ignore.characters_to_skip` field for `google_data_loss_prevention_deidentify_template` ([#3983](https://github.com/hashicorp/terraform-provider-google-beta/pull/3983))
* storagetransfer: fixed bug where `schedule` was required, but really it is optional. ([#3995](https://github.com/hashicorp/terraform-provider-google-beta/pull/3995))

## 4.7.0 (January 19, 2022)

IMPROVEMENTS:
* compute: added `EXTERNAL_MANAGED` as option for `load_balancing_scheme` in `google_compute_backend_service` resource ([#3975](https://github.com/hashicorp/terraform-provider-google-beta/pull/3975))
* container: promoted `dns_config` field of `google_container_cluster` to GA ([#3978](https://github.com/hashicorp/terraform-provider-google-beta/pull/3978))
* monitoring: added `conditionMatchedLog` and `alertStrategy` fields to `google_monitoring_alert_policy` resource ([#3968](https://github.com/hashicorp/terraform-provider-google-beta/pull/3968))

## 4.6.0 (January 10, 2022)

BREAKING CHANGES:
* pubsub: changed `google_pubsub_schema` so that modifiying fields will recreate the resource rather than causing Terraform to report it would attempt an invalid update ([#3933](https://github.com/hashicorp/terraform-provider-google-beta/pull/3933))

FEATURES:
* **New Resource:** `google_apigee_nat_address` ([#3941](https://github.com/hashicorp/terraform-provider-google-beta/pull/3941))
* **New Resource:** `google_network_connectivity_hub` ([#3947](https://github.com/hashicorp/terraform-provider-google-beta/pull/3947))

IMPROVEMENTS:
* bigquery: added ability to create a table with both a schema and view simultaneously to `google_bigquery_table` ([#3950](https://github.com/hashicorp/terraform-provider-google-beta/pull/3950))
* cloud_composer: Added support for Cloud Composer master authorized networks flag ([#3937](https://github.com/hashicorp/terraform-provider-google-beta/pull/3937))
* container: Added field `identity_service_config` to `google_container_cluster` ([#3957](https://github.com/hashicorp/terraform-provider-google-beta/pull/3957))
* osconfig: Added daily os config patch deployments ([#3945](https://github.com/hashicorp/terraform-provider-google-beta/pull/3945))
* storage: added configurable read timeout to `google_storage_bucket` ([#3938](https://github.com/hashicorp/terraform-provider-google-beta/pull/3938))

BUG FIXES:
* billingbudget: fixed a bug where `google_billing_budget.budget_filter.labels` was not updating. ([#3932](https://github.com/hashicorp/terraform-provider-google-beta/pull/3932))
* compute: fixed scenario where `region_instance_group_manager` would not start update if `wait_for_instances` was set and initial status was not `STABLE` ([#3949](https://github.com/hashicorp/terraform-provider-google-beta/pull/3949))
* healthcare: Added back `self_link` functionality which was accidentally removed in `4.0.0` release. ([#3946](https://github.com/hashicorp/terraform-provider-google-beta/pull/3946))
* pubsub: fixed update failure when attempting to change non-updatable resource `google_pubsub_schema` ([#3933](https://github.com/hashicorp/terraform-provider-google-beta/pull/3933))
* storage: fixed a bug where `google_storage_bucket.lifecycle_rule.condition.days_since_custom_time` was not updating. ([#3936](https://github.com/hashicorp/terraform-provider-google-beta/pull/3936))
* vpcaccess: Added back `self_link` functionality which was accidentally removed in `4.0.0` release. ([#3946](https://github.com/hashicorp/terraform-provider-google-beta/pull/3946))

## 4.5.0 (December 20, 2021)

FEATURES:
* **New Data Source:** google_container_aws_versions ([#3928](https://github.com/hashicorp/terraform-provider-google-beta/pull/3928))
* **New Data Source:** google_container_azure_versions ([#3928](https://github.com/hashicorp/terraform-provider-google-beta/pull/3928))
* **New Resource:** google_container_aws_cluster ([#3928](https://github.com/hashicorp/terraform-provider-google-beta/pull/3928))
* **New Resource:** google_container_aws_node_pool ([#3928](https://github.com/hashicorp/terraform-provider-google-beta/pull/3928))
* **New Resource:** google_container_azure_client ([#3928](https://github.com/hashicorp/terraform-provider-google-beta/pull/3928))
* **New Resource:** google_container_azure_cluster ([#3928](https://github.com/hashicorp/terraform-provider-google-beta/pull/3928))
* **New Resource:** google_container_azure_node_pool ([#3928](https://github.com/hashicorp/terraform-provider-google-beta/pull/3928))


IMPROVEMENTS:
* bigquery: added the `return_table_type` field to `google_bigquery_routine` ([#3922](https://github.com/hashicorp/terraform-provider-google-beta/pull/3922))
* cloudbuild: added support for `available_secrets` to `google_cloudbuild_trigger` ([#3907](https://github.com/hashicorp/terraform-provider-google-beta/pull/3907))
* cloudfunctions: added support for `min_instances` to `google_cloudfunctions_function` ([#3904](https://github.com/hashicorp/terraform-provider-google-beta/pull/3904))
* composer: added support for Private Service Connect by adding field `cloud_composer_connection_subnetwork` in `google_composer_environment` ([#3912](https://github.com/hashicorp/terraform-provider-google-beta/pull/3912))
* compute: fixed bug where `google_compute_instance`'s `can_ip_forward` could not be updated without recreating or restarting the instance. ([#3920](https://github.com/hashicorp/terraform-provider-google-beta/pull/3920))
* compute: added field `public_access_prevention` to resource `bucket` ([#3919](https://github.com/hashicorp/terraform-provider-google-beta/pull/3919))
* compute: added support for regional external HTTP(S) load balancer ([#3916](https://github.com/hashicorp/terraform-provider-google-beta/pull/3916))
* privateca: added support for setting default values for basic constraints for `google_privateca_certificate`, `google_privateca_certificate_authority`, and `google_privateca_ca_pool` via the `non_ca` and `zero_max_issuer_path_length` fields ([#3902](https://github.com/hashicorp/terraform-provider-google-beta/pull/3902))
* provider: enabled gRPC requests and response logging ([#3910](https://github.com/hashicorp/terraform-provider-google-beta/pull/3910))

BUG FIXES:
* assuredworkloads: fixed a bug preventing `google_assured_workloads_workload` from being created in any region other than us-central1 ([#3925](https://github.com/hashicorp/terraform-provider-google-beta/pull/3925))

## 4.4.0 (December 13, 2021)

DEPRECATIONS:
* filestore: deprecated `zone` on `google_filestore_instance` in favor of `location` to allow for regional instances ([#3887](https://github.com/hashicorp/terraform-provider-google-beta/pull/3887))

FEATURES:
* **New Resource:** `google_os_config_os_policy_assignment` ([#3892](https://github.com/hashicorp/terraform-provider-google-beta/pull/3892))
* **New Resource:** `google_recaptcha_enterprise_key` ([#3890](https://github.com/hashicorp/terraform-provider-google-beta/pull/3890))

IMPROVEMENTS:
* filestore: added support for `ENTERPRISE` value on `google_filestore_instance` `tier` ([#3887](https://github.com/hashicorp/terraform-provider-google-beta/pull/3887))
* privateca: added support for setting default values for basic constraints for `google_privateca_certificate`, `google_privateca_certificate_authority`, and `google_privateca_ca_pool` via the `non_ca` and `zero_max_issuer_path_length` fields ([#3902](https://github.com/hashicorp/terraform-provider-google-beta/pull/3902))
* sql: added field `allocated_ip_range` to resource `google_sql_database_instance` ([#3897](https://github.com/hashicorp/terraform-provider-google-beta/pull/3897))

BUG FIXES:
* compute: fixed incorrectly failing validation for `INTERNAL_MANAGED` `google_compute_region_backend_service`. ([#3888](https://github.com/hashicorp/terraform-provider-google-beta/pull/3888))
* compute: fixed scenario where `instance_group_manager` would not start update if `wait_for_instances` was set and initial status was not `STABLE` ([#3893](https://github.com/hashicorp/terraform-provider-google-beta/pull/3893))
* container: fixed the `ROUTES` value for the `networking_mode` field in `google_container_cluster`. A recent API change unintentionally changed the default to a `VPC_NATIVE` cluster, and removed the ability to create a `ROUTES`-based one. Provider versions prior to this one will default to `VPC_NATIVE` due to this change, and are unable to create `ROUTES` clusters. ([#3896](https://github.com/hashicorp/terraform-provider-google-beta/pull/3896))

## 4.3.0 (December 7, 2021)

FEATURES:
* **New Data Source:** `google_compute_router_status` ([#3859](https://github.com/hashicorp/terraform-provider-google-beta/pull/3859))
* **New Data Source:** `google_folders` ([#3886](https://github.com/hashicorp/terraform-provider-google-beta/pull/3886))
* **New Resource:** `google_notebooks_runtime` ([#3878](https://github.com/hashicorp/terraform-provider-google-beta/pull/3878))
* **New Resource:** `google_vertex_ai_metadata_store` ([#3885](https://github.com/hashicorp/terraform-provider-google-beta/pull/3885))

IMPROVEMENTS
* apigee: Added IAM support for `google_apigee_environment`. ([#3871](https://github.com/hashicorp/terraform-provider-google-beta/pull/3871)):
* apigee: Added supported values for 'peeringCidrRange' in `google_apigee_instance`. ([#3880](https://github.com/hashicorp/terraform-provider-google-beta/pull/3880))
* cloudbuild: added display_name and annotations to google_cloudbuild_worker_pool for compatibility with new GA. ([#3873](https://github.com/hashicorp/terraform-provider-google-beta/pull/3873))
* container: added `node_group` to `node_config` for container clusters and node pools to support sole tenancy ([#3881](https://github.com/hashicorp/terraform-provider-google-beta/pull/3881))
* container: added `spot` field to `node_config` sub-resource ([#3863](https://github.com/hashicorp/terraform-provider-google-beta/pull/3863))
* redis: Added Multi read replica field `replicaCount `, `nodes`,  `readEndpoint`, `readEndpointPort`, `readReplicasMode` in `google_redis_instance` ([#3870](https://github.com/hashicorp/terraform-provider-google-beta/pull/3870))

BUG FIXES:
* essentialcontacts: marked updating `email` in `google_essential_contacts_contact` as requiring recreation ([#3864](https://github.com/hashicorp/terraform-provider-google-beta/pull/3864))
* privateca: fixed crlAccessUrls in `CertificateAuthority ` ([#3861](https://github.com/hashicorp/terraform-provider-google-beta/pull/3861))

## 4.2.1 (December 3, 2021)

BUG FIXES:
* provider: reverted a requirement in v4.2.0 for Terraform 0.13 and above. This release should be compatible with Terraform 0.12.31

## 4.2.0 (December 2, 2021)

FEATURES:
* **New Data Source:** `google_compute_router_status` ([#3859](https://github.com/hashicorp/terraform-provider-google-beta/pull/3859))

IMPROVEMENTS:
* compute: added support for `queue_count` to `google_compute_instance.network_interface` and `google_compute_instance_template.network_interface` ([#3857](https://github.com/hashicorp/terraform-provider-google-beta/pull/3857))

BUG FIXES:
* all: fixed an issue where some documentation for new resources was not showing up in the GA provider if it was beta-only. ([#3848](https://github.com/hashicorp/terraform-provider-google-beta/pull/3848))
* bigquery: fixed update failure when attempting to change non-updatable fields in `google_bigquery_routine`. ([#3849](https://github.com/hashicorp/terraform-provider-google-beta/pull/3849))
* compute: fixed a bug that would cause `google_instance_from_machine_image` to fail with a resourceInUseByAnotherResource error ([#3855](https://github.com/hashicorp/terraform-provider-google-beta/pull/3855))
* compute: fixed a bug when `cache_mode` is set to FORCE_CACHE_ALL on `google_compute_backend_bucket` ([#3858](https://github.com/hashicorp/terraform-provider-google-beta/pull/3858))
* compute: fixed a perma-diff on `google_compute_region_health_check` when `log_config.enable` is set to false ([#3853](https://github.com/hashicorp/terraform-provider-google-beta/pull/3853))
* servicedirectory: added support for vpc network configuration in `google_service_directory_endpoint`. ([#3856](https://github.com/hashicorp/terraform-provider-google-beta/pull/3856))

## 4.1.0 (November 15, 2021)

IMPROVEMENTS:
* compute: Added `bfd` to `google_compute_router_peer` ([#3822](https://github.com/hashicorp/terraform-provider-google-beta/pull/3822))
* container: added `gcfs_config` to `node_config` of `google_container_node_pool` resource ([#3828](https://github.com/hashicorp/terraform-provider-google-beta/pull/3828))
* provider: added retries for the `resourceNotReady` error returned when attempting to add resources to a recently-modified subnetwork ([#3827](https://github.com/hashicorp/terraform-provider-google-beta/pull/3827))
* pubsub: added `message_retention_duration` field to `google_pubsub_topic` ([#3831](https://github.com/hashicorp/terraform-provider-google-beta/pull/3831))

BUG FIXES:
* apigee: fixed a bug where multiple `google_apigee_instance_attachment` could not be used on the same `google_apigee_instance` ([#3838](https://github.com/hashicorp/terraform-provider-google-beta/pull/3838))
* bigquery: fixed a bug following import where schema is empty on `google_bigquery_table` ([#3839](https://github.com/hashicorp/terraform-provider-google-beta/pull/3839))
* billingbudget: fixed unable to provide `labels` on `google_billing_budget` ([#3823](https://github.com/hashicorp/terraform-provider-google-beta/pull/3823))
* compute: allowed `source_disk` to accept full image path on `google_compute_snapshot` ([#3835](https://github.com/hashicorp/terraform-provider-google-beta/pull/3835))
* compute: fixed a bug in `google_compute_firewall` that would cause changes in `source_ranges` to not correctly be applied ([#3834](https://github.com/hashicorp/terraform-provider-google-beta/pull/3834))
* logging: fixed a bug with updating `description` on `google_logging_project_sink`, `google_logging_folder_sink` and `google_logging_organization_sink` ([#3826](https://github.com/hashicorp/terraform-provider-google-beta/pull/3826))

## 4.0.0 (November 02, 2021)

NOTES:
* compute: Google Compute Engine resources will now call the endpoint appropriate to the provider version rather than the beta endpoint by default ([#3787](https://github.com/hashicorp/terraform-provider-google-beta/pull/3787))
* container: Google Kubernetes Engine resources will now call the endpoint appropriate to the provider version rather than the beta endpoint by default ([#3788](https://github.com/hashicorp/terraform-provider-google-beta/pull/3788))

BREAKING CHANGES:
* appengine: marked `google_app_engine_standard_app_version` `entrypoint` as required ([#3784](https://github.com/hashicorp/terraform-provider-google-beta/pull/3784))
* compute: removed the ability to specify the `trace-append` or `trace-ro` as scopes in `google_compute_instance`, use `trace` instead ([#3759](https://github.com/hashicorp/terraform-provider-google-beta/pull/3759))
* compute: changed `advanced_machine_features` on `google_compute_instance_template` to track changes when the block is undefined in a user's config ([#3786](https://github.com/hashicorp/terraform-provider-google-beta/pull/3786))
* compute: changed `source_ranges` in `google_compute_firewall_rule` to track changes when it is not set in a config file ([#3791](https://github.com/hashicorp/terraform-provider-google-beta/pull/3791))
* compute: changed the import / drift detection behaviours for `metadata_startup_script`, `metadata.startup-script` in `google_compute_instance`. Now, `metadata.startup-script` will be set by default, and `metadata_startup_script` will only be set if present. ([#3765](https://github.com/hashicorp/terraform-provider-google-beta/pull/3765))
* compute: removed `source_disk_link` field from `google_compute_snapshot` ([#3783](https://github.com/hashicorp/terraform-provider-google-beta/pull/3783))
* container: `instance_group_urls` has been removed in favor of `node_pool.instance_group_urls` ([#3796](https://github.com/hashicorp/terraform-provider-google-beta/pull/3796))
* container: changed default for `enable_shielded_nodes` to true for `google_container_cluster` ([#3773](https://github.com/hashicorp/terraform-provider-google-beta/pull/3773))
* container: made `master_auth.client_certificate_config` required ([#3794](https://github.com/hashicorp/terraform-provider-google-beta/pull/3794))
* container: removed `master_auth.username` and `master_auth.password` from `google_container_cluster` ([#3794](https://github.com/hashicorp/terraform-provider-google-beta/pull/3794))
* container: removed `workload_metadata_configuration.node_metadata` in favor of `workload_metadata_configuration.mode` in `google_container_cluster` ([#3772](https://github.com/hashicorp/terraform-provider-google-beta/pull/3772))
* container: removed the `workload_identity_config.0.identity_namespace` field from `google_container_cluster`, use `workload_identity_config.0.workload_pool` instead ([#3776](https://github.com/hashicorp/terraform-provider-google-beta/pull/3776))
* kms: removed `self_link` field from `google_kms_crypto_key` and `google_kms_key_ring` ([#3783](https://github.com/hashicorp/terraform-provider-google-beta/pull/3783))
* project: removed ability to specify `bigquery-json.googleapis.com`, the provider will no longer convert it as the upstream API migration is finished. Use `bigquery.googleapis.com` instead. ([#3751](https://github.com/hashicorp/terraform-provider-google-beta/pull/3751))
* provider: changed `credentials`, `access_token` precedence so that `credentials` values in configuration take precedence over `access_token` values assigned through environment variables ([#3766](https://github.com/hashicorp/terraform-provider-google-beta/pull/3766))
* provider: removed redundant default scopes. The provider's default scopes when authenticating with credentials are now exclusively "https://www.googleapis.com/auth/cloud-platform" and "https://www.googleapis.com/auth/userinfo.email". ([#3756](https://github.com/hashicorp/terraform-provider-google-beta/pull/3756))
* pubsub: removed `path` from `google_pubsub_subscription` ([#3777](https://github.com/hashicorp/terraform-provider-google-beta/pull/3777))
* pubsub: removed `path` field from `google_pubsub_subscription` ([#3783](https://github.com/hashicorp/terraform-provider-google-beta/pull/3783))
* resourcemanager: made `google_project` remove `org_id` and `folder_id` from state when they are removed from config ([#3754](https://github.com/hashicorp/terraform-provider-google-beta/pull/3754))
* resourcemanager: changed the `project` field to `Required` in all `google_project_iam_*` resources ([#3767](https://github.com/hashicorp/terraform-provider-google-beta/pull/3767))
* sql: added drift detection to the following `google_sql_database_instance` fields: `activation_policy` (defaults `ALWAYS`), `availability_type` (defaults `ZONAL`), `disk_type` (defaults `PD_SSD`), `encryption_key_name` ([#3778](https://github.com/hashicorp/terraform-provider-google-beta/pull/3778))
* sql: changed the `database_version` field to `Required` in `google_sql_database_instance` resource ([#3770](https://github.com/hashicorp/terraform-provider-google-beta/pull/3770))
* sql: removed the following `google_sql_database_instance` fields: `authorized_gae_applications`, `crash_safe_replication`, `replication_type` ([#3778](https://github.com/hashicorp/terraform-provider-google-beta/pull/3778))
* storage: removed `bucket_policy_only` from `google_storage_bucket` ([#3769](https://github.com/hashicorp/terraform-provider-google-beta/pull/3769))
* storage: changed the `location` field to required in `google_storage_bucket` ([#3771](https://github.com/hashicorp/terraform-provider-google-beta/pull/3771))

VALIDATION CHANGES:
* bigquery: at least one of `statement_timeout_ms`, `statement_byte_budget`, or `key_result_statement` is required on `google_bigquery_job.query.script_options.` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* bigquery: exactly one of `query`, `load`, `copy` or `extract` is required on `google_bigquery_job` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* bigquery: exactly one of `source_table` or `source_model` is required on `google_bigquery_job.extract` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* cloudbuild: exactly one of `branch_name`, `commit_sha` or `tag_name` is required on `google_cloudbuild_trigger.build.source.repo_source` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `fixed_delay` or `percentage` is required on `google_compute_url_map.default_route_action.fault_injection_policy.delay` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `fixed` or `percent` is required on `google_compute_autoscaler.autoscaling_policy.scale_down_control.max_scaled_down_replicas` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `fixed` or `percent` is required on `google_compute_autoscaler.autoscaling_policy.scale_in_control.max_scaled_in_replicas` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `fixed` or `percent` is required on `google_compute_region_autoscaler.autoscaling_policy.scale_down_control.max_scaled_down_replicas` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `fixed` or `percent` is required on `google_compute_region_autoscaler.autoscaling_policy.scale_in_control.max_scaled_in_replicas` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `max_scaled_down_replicas` or `time_window_sec` is required on `google_compute_autoscaler.autoscaling_policy.scale_down_control` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `max_scaled_down_replicas` or `time_window_sec` is required on `google_compute_region_autoscaler.autoscaling_policy.scale_down_control` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `max_scaled_in_replicas` or `time_window_sec` is required on `google_compute_autoscaler.autoscaling_policy.scale_in_control.0.` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: at least one of `max_scaled_in_replicas` or `time_window_sec` is required on `google_compute_region_autoscaler.autoscaling_policy.scale_in_control.0.` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* compute: required one of `source_tags`, `source_ranges` or `source_service_accounts` on INGRESS `google_compute_firewall` resources ([#3750](https://github.com/hashicorp/terraform-provider-google-beta/pull/3750))
* dlp: at least one of `start_time` or `end_time` is required on `google_data_loss_prevention_trigger.inspect_job.storage_config.timespan_config` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* dlp: exactly one of `url` or `regex_file_set` is required on `google_data_loss_prevention_trigger.inspect_job.storage_config.cloud_storage_options.file_set` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* resourcemanager: added conflict between `org_id`, `folder_id` at plan time in `google_project` ([#3754](https://github.com/hashicorp/terraform-provider-google-beta/pull/3754))
* osconfig: at least one of `linux_exec_step_config` or `windows_exec_step_config` is required on `google_os_config_patch_deployment.patch_config.post_step` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* osconfig: at least one of `linux_exec_step_config` or `windows_exec_step_config` is required on `google_os_config_patch_deployment.patch_config.pre_step` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* osconfig: at least one of `reboot_config`, `apt`, `yum`, `goo` `zypper`, `windows_update`, `pre_step` or `pre_step` is required on `google_os_config_patch_deployment.patch_config` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* osconfig: at least one of `security`, `minimal`, `excludes` or `exclusive_packages` is required on `google_os_config_patch_deployment.patch_config.yum` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* osconfig: at least one of `type`, `excludes` or `exclusive_packages` is required on `google_os_config_patch_deployment.patch_config.apt` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* osconfig: at least one of `with_optional`, `with_update`, `categories`, `severities`, `excludes` or `exclusive_patches` is required on `google_os_config_patch_deployment.patch_config.zypper` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* osconfig: exactly one of `classifications`, `excludes` or `exclusive_patches` is required on `google_os_config_patch_deployment.inspect_job.patch_config.windows_update` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))
* spanner: at least one of `num_nodes` or `processing_units` is required on `google_spanner_instance` ([#3752](https://github.com/hashicorp/terraform-provider-google-beta/pull/3752))

IMPROVEMENTS:
* container: added `managed_instance_group_urls` to `google_container_node_pool` to replace `instance_group_urls` on `google_container_cluster` ([#3815](https://github.com/hashicorp/terraform-provider-google-beta/pull/3815))
* kms: added support for EKM to `google_kms_crypto_key.protection_level` ([#3763](https://github.com/hashicorp/terraform-provider-google-beta/pull/3763))
* project: added support for `billing_project` on `google_project_service` ([#3768](https://github.com/hashicorp/terraform-provider-google-beta/pull/3768))
* spanner: increased the default timeout on `google_spanner_instance` operations from 4 minutes to 20 minutes, significantly reducing the likelihood that resources will time out ([#3789](https://github.com/hashicorp/terraform-provider-google-beta/pull/3789))

BUG FIXES:
* bigquery: fixed a bug of cannot add required fields to an existing schema on `google_bigquery_table` ([#3781](https://github.com/hashicorp/terraform-provider-google-beta/pull/3781))
* compute: fixed a bug in updating multiple `ttl` fields on `google_compute_backend_bucket` ([#3757](https://github.com/hashicorp/terraform-provider-google-beta/pull/3757))
* compute: fixed a perma-diff on `subnetwork` when it is optional on `google_compute_network_endpoint_group` ([#3780](https://github.com/hashicorp/terraform-provider-google-beta/pull/3780))
* compute: fixed perma-diff bug on `log_config.enable` of both `google_compute_backend_service` and `google_compute_region_backend_service` ([#3760](https://github.com/hashicorp/terraform-provider-google-beta/pull/3760))
* compute: fixed the `google_compute_instance_group_manager.update_policy.0.min_ready_sec` field so that updating it to `0` works ([#3810](https://github.com/hashicorp/terraform-provider-google-beta/pull/3810))
* compute: fixed the `google_compute_region_instance_group_manager.update_policy.0.min_ready_sec` field so that updating it to `0` works ([#3810](https://github.com/hashicorp/terraform-provider-google-beta/pull/3810))
* spanner: fixed the schema for `data.google_spanner_instance` so that non-configurable fields are considered outputs ([#3804](https://github.com/hashicorp/terraform-provider-google-beta/pull/3804))

## 3.90.1 (November 02, 2021)

DEPRECATIONS:

* container: fixed an overly-broad deprecation on `master_auth`, constraining it to `master_auth.username` and `master_auth.password`

## 3.90.0 (October 26, 2021)

DEPRECATIONS:
* container: deprecated `workload_identity_config.0.identity_namespace` and it will be removed in a future major release as it has been deprecated in the API. Use `workload_identity_config.0.workload_pool` instead. Switching your configuration from one value to the other will trigger a diff at plan time, and a spurious update. ([#3733](https://github.com/hashicorp/terraform-provider-google-beta/pull/3733))
* container: deprecated the following `google_container_cluster` fields: `instance_group_urls` and `master_auth` ([#3746](https://github.com/hashicorp/terraform-provider-google-beta/pull/3746))

IMPROVEMENTS:
* composer: added field `environment_size` to resource `google_composer_environment` ([#3730](https://github.com/hashicorp/terraform-provider-google-beta/pull/3730))
* container: added `node_config.0.guest_accelerator.0.gpu_partition_size` field to google_container_node_pool ([#3739](https://github.com/hashicorp/terraform-provider-google-beta/pull/3739))
* container: added `workload_identity_config.0.workload_pool` to `google_container_cluster` ([#3733](https://github.com/hashicorp/terraform-provider-google-beta/pull/3733))
* container: made `dns_cache_config` conflict with GKE Autopilot mode ([#3725](https://github.com/hashicorp/terraform-provider-google-beta/pull/3725))
* container_cluster: Updated `monitoring_config` to accept `WORKLOAD` ([#3732](https://github.com/hashicorp/terraform-provider-google-beta/pull/3732))
* provider: Added links to nested types documentation for manually generated pages ([#3736](https://github.com/hashicorp/terraform-provider-google-beta/pull/3736))

BUG FIXES:
* cloudrun: fixed a permadiff on the field `template.spec.containers.ports.name` of the `google_cloud_run_service` resource ([#3740](https://github.com/hashicorp/terraform-provider-google-beta/pull/3740))
* composer: removed `config.node_config.zone` requirement on `google_composer_environment` ([#3745](https://github.com/hashicorp/terraform-provider-google-beta/pull/3745))
* compute: fixed permadiff for `failover_policy` on `google_compute_region_backend_service` ([#3728](https://github.com/hashicorp/terraform-provider-google-beta/pull/3728))
* compute: fixed to make `description` updatable without recreation on `google_compute_instance_group_manager` ([#3735](https://github.com/hashicorp/terraform-provider-google-beta/pull/3735))
* container: fixed a permadiff on `google_container_node_pool.workload_metadata_config.mode` ([#3726](https://github.com/hashicorp/terraform-provider-google-beta/pull/3726))
* iam: fixed request batching bug where failed requests would show unnecessary backslash escaping to the user. ([#3723](https://github.com/hashicorp/terraform-provider-google-beta/pull/3723))
* securitycenter: fixed bug where `google_scc_notification_config.streaming_config.filter` was not updating. ([#3727](https://github.com/hashicorp/terraform-provider-google-beta/pull/3727))

## 3.89.0 (October 18, 2021)

BUG FIXES:
* compute: fixed bug where `google_compute_router_peer` could not set an advertised route priority of 0, causing permadiff. ([#3718](https://github.com/hashicorp/terraform-provider-google-beta/pull/3718))
* container: fixed a crash on `monitoring_config` of `google_container_cluster` ([#3717](https://github.com/hashicorp/terraform-provider-google-beta/pull/3717))
* iam: fixed request batching bug where failed requests would show unnecessary backslash escaping to the user. ([#3723](https://github.com/hashicorp/terraform-provider-google-beta/pull/3723))
* storage: fixed a bug to better handle eventual consistency among `google_storage_bucket` resources. ([#3715](https://github.com/hashicorp/terraform-provider-google-beta/pull/3715))

## 3.88.0 (October 11, 2021)
NOTES:
* reorganized documentation to group all Compute Engine and Monitoring (Stackdriver) resources together. ([#3686](https://github.com/hashicorp/terraform-provider-google-beta/pull/3686))

DEPRECATIONS:
* container: deprecated `workload_metadata_configuration.node_metadata` in favor of `workload_metadata_configuration.mode` in `google_container_cluster` ([#3694](https://github.com/hashicorp/terraform-provider-google-beta/pull/3694))
* dataproc: deprecated the `google_dataproc_workflow_template.version` field, as it wasn't actually useful. The field is used during updates, but updates aren't currently possible with the resource. ([#3675](https://github.com/hashicorp/terraform-provider-google-beta/pull/3675))
BREAKING CHANGES:
* gke_hub: made the `config_membership` field in `google_gke_hub_feature` required, disallowing invalid configurations ([#3681](https://github.com/hashicorp/terraform-provider-google-beta/pull/3681))
* gke_hub: made the `configmanagement`, `feature`, `location`, `membership` fields in `google_gke_hub_feature_membership` required, disallowing invalid configurations ([#3681](https://github.com/hashicorp/terraform-provider-google-beta/pull/3681))

FEATURES:
* **New Data Source:** `google_service_networking_peered_dns_domain` ([#3690](https://github.com/hashicorp/terraform-provider-google-beta/pull/3690))
* **New Data Source:** `google_sourcerepo_repository` ([#3684](https://github.com/hashicorp/terraform-provider-google-beta/pull/3684))
* **New Data Source:** `google_storage_bucket` ([#3678](https://github.com/hashicorp/terraform-provider-google-beta/pull/3678))
* **New Resource:** `google_pubsub_lite_reservation` ([#3708](https://github.com/hashicorp/terraform-provider-google-beta/pull/3708))
* **New Resource:** `google_service_networking_peered_dns_domain` ([#3690](https://github.com/hashicorp/terraform-provider-google-beta/pull/3690))

IMPROVEMENTS:
* composer: added field `enable_privately_used_public_ips` to resource `google_composer_environment` (beta) ([#3697](https://github.com/hashicorp/terraform-provider-google-beta/pull/3697))
* composer: added field `enable_ip_masq_agent` to resource `google_composer_environment` (beta) ([#3705](https://github.com/hashicorp/terraform-provider-google-beta/pull/3705))
* composer: added support for composer v2 fields `workloads_config` and `cloud_composer_network_ipv4_cidr_block` to `composer_environment` ([#3709](https://github.com/hashicorp/terraform-provider-google-beta/pull/3709))
* compute: added NetLB support for Connection Tracking as `connectionTrackingPolicy` in `RegionBackendService`(beta) ([#3698](https://github.com/hashicorp/terraform-provider-google-beta/pull/3698))
* compute: added external IPv6 support on `google_compute_subnetwork` and `google_compute_instance.network_interfaces` ([#3677](https://github.com/hashicorp/terraform-provider-google-beta/pull/3677))
* container: added support for `workload_metadata_configuration.mode` in `google_container_cluster` ([#3694](https://github.com/hashicorp/terraform-provider-google-beta/pull/3694))
* eventarc: added support for `uid` output field, `cloud_function` destination to `google_eventarc_trigger` ([#3681](https://github.com/hashicorp/terraform-provider-google-beta/pull/3681))
* gke_hub: added support for `gcp_service_account_email` when configuring Git sync in `google_gke_hub_feature_membership` ([#3681](https://github.com/hashicorp/terraform-provider-google-beta/pull/3681))
* gke_hub: added support for `resource_state`, `state` outputs to `google_gke_hub_feature` ([#3681](https://github.com/hashicorp/terraform-provider-google-beta/pull/3681))
* pubsub:  added support for references to `google_pubsub_lite_reservation` to `google_pubsub_lite_topic`. ([#3708](https://github.com/hashicorp/terraform-provider-google-beta/pull/3708))

BUG FIXES:
* monitoring: fixed typo in `google_monitoring_uptime_check_config` where `NOT_MATCHES_REGEX` could not be specified. ([#3700](https://github.com/hashicorp/terraform-provider-google-beta/pull/3700))
* servicedirectory: marked `service` on `google_service_directory_endpoint` as ForceNew to trigger recreates on changes ([#3683](https://github.com/hashicorp/terraform-provider-google-beta/pull/3683))

## 3.87.0 (October 04, 2021)

DEPRECATIONS:
* dataproc: deprecated the `google_dataproc_workflow_template.version` field, as it wasn't actually useful. The field is used during updates, but updates aren't currently possible with the resource. ([#3675](https://github.com/hashicorp/terraform-provider-google-beta/pull/3675))

FEATURES:
* **New Resource:** `google_monitoring_monitored_project` ([#3658](https://github.com/hashicorp/terraform-provider-google-beta/pull/3658))
* **New Resource:** `google_org_policy_policy` ([#3637](https://github.com/hashicorp/terraform-provider-google-beta/pull/3637))

IMPROVEMENTS:
* cloudbuild: added field `service_account` to `google_cloudbuild_trigger` ([#3661](https://github.com/hashicorp/terraform-provider-google-beta/pull/3661))
* composer: added field `scheduler_count` to `google_composer_environment` ([#3660](https://github.com/hashicorp/terraform-provider-google-beta/pull/3660))
* compute: Disabled recreation of GCE instances when updating `resource_policies` property ([#3668](https://github.com/hashicorp/terraform-provider-google-beta/pull/3668))
* container: added support for `logging_config` and `monitoring_config` to `google_container_cluster` ([#3641](https://github.com/hashicorp/terraform-provider-google-beta/pull/3641))
* kms: added support for `import_only` to `google_kms_crypto_key` ([#3659](https://github.com/hashicorp/terraform-provider-google-beta/pull/3659))
* networkservices: boosted the default timeout for `google_network_services_edge_cache_origin` from 30m to 60m ([#3674](https://github.com/hashicorp/terraform-provider-google-beta/pull/3674))

BUG FIXES:
* container: fixed an issue where a node pool created with error (eg. GKE_STOCKOUT) would not be captured in state ([#3646](https://github.com/hashicorp/terraform-provider-google-beta/pull/3646))
* filestore: Allowed updating `reserved_ip_range` on `google_filestore_instance` via recreation of the instance ([#3651](https://github.com/hashicorp/terraform-provider-google-beta/pull/3651))
* serviceusage: Made the service api retry failed operation calls in anticipation of transient errors that occur when first enabling the service. ([#3666](https://github.com/hashicorp/terraform-provider-google-beta/pull/3666))

## 3.86.0 (September 27, 2021)

BUG FIXES:
* dns: fixed an issue in `google_dns_record_set` where `rrdatas` could not be updated ([#3625](https://github.com/hashicorp/terraform-provider-google-beta/pull/3625))
* dns: fixed an issue in `google_dns_record_set` where creating the resource would result in an 409 error ([#3625](https://github.com/hashicorp/terraform-provider-google-beta/pull/3625))
* platform: fixed a bug in wrongly writing to state when creation failed on `google_organization_policy` ([#3624](https://github.com/hashicorp/terraform-provider-google-beta/pull/3624))

## 3.85.0 (September 20, 2021)
DEPRECATIONS:
* compute: deprecated `interface` field on `google_compute_disk` and `google_compute_region_disk` ([#3611](https://github.com/hashicorp/terraform-provider-google-beta/pull/3611))

IMPROVEMENTS:
* bigtable: enabled support for `user_project_override` in `google_bigtable_instance` and `google_bigtable_table` ([#3614](https://github.com/hashicorp/terraform-provider-google-beta/pull/3614))
* compute: added `iap` fields to `google_compute_region_backend_service` ([#3605](https://github.com/hashicorp/terraform-provider-google-beta/pull/3605))
* compute: allowed passing an IP address to the `nextHopIlb` field of `google_compute_route` resource ([#3609](https://github.com/hashicorp/terraform-provider-google-beta/pull/3609))
* container: added field `dns_config` to resource `google_container_cluster` ([#3606](https://github.com/hashicorp/terraform-provider-google-beta/pull/3606))
* iam: added `disabled` field to `google_service_account` resource ([#3603](https://github.com/hashicorp/terraform-provider-google-beta/pull/3603))
* provider: added links to nested types documentation within a resource ([#3615](https://github.com/hashicorp/terraform-provider-google-beta/pull/3615))
* storage: added field `path` to `google_storage_transfer_job` ([#3608](https://github.com/hashicorp/terraform-provider-google-beta/pull/3608))

BUG FIXES:
* appengine: fixed bug where `deployment.container.image` would update to an old version even if in `ignore_changes` ([#3613](https://github.com/hashicorp/terraform-provider-google-beta/pull/3613))
* bigquery: fixed a bug where `destination_encryption_config.kms_key_name` stored the version rather than the key name. ([#3616](https://github.com/hashicorp/terraform-provider-google-beta/pull/3616))
* redis: extended the default timeouts on `google_redis_instance` ([#3604](https://github.com/hashicorp/terraform-provider-google-beta/pull/3604))
* serviceusage: fixed an issue in `google_project_service` where users could not reenable services that were disabled outside of Terraform. ([#3607](https://github.com/hashicorp/terraform-provider-google-beta/pull/3607))

## 3.84.0 (September 13, 2021)
DEPRECATIONS:
* compute: deprecated `interface` field on `google_compute_disk` and `google_compute_region_disk` ([#3611](https://github.com/hashicorp/terraform-provider-google-beta/pull/3611))

FEATURES:
* **New Data Source:** `google_secret_manager_secret` ([#3588](https://github.com/hashicorp/terraform-provider-google-beta/pull/3588))

IMPROVEMENTS:
* compute: added update support to `google_compute_service_attachment` ([#3587](https://github.com/hashicorp/terraform-provider-google-beta/pull/3587))
* filestore: added `connect_mode` to `networks` field in `google_filestore_instance` ([#3595](https://github.com/hashicorp/terraform-provider-google-beta/pull/3595))

BUG FIXES:
* container: fixed a bug in failing to remove `maintenance_exclusion` on `google_container_cluster` ([#3600](https://github.com/hashicorp/terraform-provider-google-beta/pull/3600))
* compute: fixed `advanced_machine_features` error messages in `google_compute_instance` ([#3598](https://github.com/hashicorp/terraform-provider-google-beta/pull/3598))
* eventarc: fixed bug where resources deleted outside of Terraform would cause errors ([#3590](https://github.com/hashicorp/terraform-provider-google-beta/pull/3590))
* functions: fixed an error message on `google_cloudfunctions_function` ([#3591](https://github.com/hashicorp/terraform-provider-google-beta/pull/3591))
* logging: fixed the data type for `bucket_options.linear_buckets.width` on `google_logging_metric` ([#3589](https://github.com/hashicorp/terraform-provider-google-beta/pull/3589))
* osconfig: fixed import on google_os_config_guest_policies ([#3594](https://github.com/hashicorp/terraform-provider-google-beta/pull/3594))
* storage: fixed an undetected change on `days_since_noncurrent_time` of `google_storage_bucket` ([#3599](https://github.com/hashicorp/terraform-provider-google-beta/pull/3599))


## 3.83.0 (September 09, 2021)
FEATURES:
* **New Resource:** `google_privateca_certificate_template` ([#3561](https://github.com/hashicorp/terraform-provider-google-beta/pull/3561))

IMPROVEMENTS:
* privateca: added `certificate_template` to `google_privateca_certificate`. ([#3567](https://github.com/hashicorp/terraform-provider-google-beta/pull/3567))
* compute: allowed setting `ip_address` field of `google_compute_router_peer` ([#3565](https://github.com/hashicorp/terraform-provider-google-beta/pull/3565))
* dataproc: added field `metastore_config` to `google_dataproc_cluster` ([#3577](https://github.com/hashicorp/terraform-provider-google-beta/pull/3577))
* kms: added support for `destroy_scheduled_duration` to `google_kms_crypto_key` ([#3563](https://github.com/hashicorp/terraform-provider-google-beta/pull/3563))

BUG FIXES:
* endpoints: fixed a timezone discrepancy in `config_id` on `google_endpoints_service` ([#3564](https://github.com/hashicorp/terraform-provider-google-beta/pull/3564))
* cloudbuild: marked `google_cloudbuild_trigger` as requiring one of branch_name/tag_name/commit_sha  within build.source.repo_source ([#3582](https://github.com/hashicorp/terraform-provider-google-beta/pull/3582))
* compute: fixed a crash on `enable` field of `google_compute_router_peer` ([#3579](https://github.com/hashicorp/terraform-provider-google-beta/pull/3579))
* compute: fixed a permanent diff for `next_hop_instance_zone` on `google_compute_route` when `next_hop_instance` was set to a self link ([#3571](https://github.com/hashicorp/terraform-provider-google-beta/pull/3571))
* compute: fixed an issue in `google_compute_router_nat` where removing `log_config` resulted in a perma-diff ([#3581](https://github.com/hashicorp/terraform-provider-google-beta/pull/3581))
* privateca: fixed a permadiff bug for `publishing_options` on `google_privateca_ca_pool` when both attributes set false ([#3570](https://github.com/hashicorp/terraform-provider-google-beta/pull/3570))
* spanner: fixed instance updates to processing units ([#3575](https://github.com/hashicorp/terraform-provider-google-beta/pull/3575))
* storage: added support for timeouts on `google_storage_bucket_object` ([#3578](https://github.com/hashicorp/terraform-provider-google-beta/pull/3578))

## 3.82.0 (August 30, 2021)
FEATURES:
* **New Resource:** `google_privateca_certificate_template` ([#3561](https://github.com/hashicorp/terraform-provider-google-beta/pull/3561))
* **New Resource:** `google_compute_firewall_policy` ([#3556](https://github.com/hashicorp/terraform-provider-google-beta/pull/3556))
* **New Resource:** `google_compute_firewall_policy_association` ([#3556](https://github.com/hashicorp/terraform-provider-google-beta/pull/3556))
* **New Resource:** `google_compute_firewall_policy_rule` ([#3556](https://github.com/hashicorp/terraform-provider-google-beta/pull/3556))

IMPROVEMENTS:
* notebooks: added support for `nic_type`, `reservation_affinity` to `google_notebooks_instance` ([#3554](https://github.com/hashicorp/terraform-provider-google-beta/pull/3554))
* sql: added field `collation` to `google_sql_database_instance` ([#3557](https://github.com/hashicorp/terraform-provider-google-beta/pull/3557))

BUG FIXES:
* apigateway: fixed import functionality for all `apigateway` resources ([#3549](https://github.com/hashicorp/terraform-provider-google-beta/pull/3549))
* compute: fixed a bug when a `source_machine_image` from a different project is used on `google_compute_instance_from_machine_image` ([#3541](https://github.com/hashicorp/terraform-provider-google-beta/pull/3541))
* dns: fixed not-exists error message on data source `google_dns_managed_zone` ([#3559](https://github.com/hashicorp/terraform-provider-google-beta/pull/3559))
* healthcare: fixed bug where changes to `google_healthcare_hl7_v2_store.parser_config` subfields would error with "...parser_config.version field is immutable..." ([#3560](https://github.com/hashicorp/terraform-provider-google-beta/pull/3560))
* os_config: fixed imports for `google_os_config_guest_policies` ([#3550](https://github.com/hashicorp/terraform-provider-google-beta/pull/3550))
* pubsub: added polling to `google_pubsub_schema` to deal with eventually consistent deletes ([#3544](https://github.com/hashicorp/terraform-provider-google-beta/pull/3544))
* secretmanager: fixed an issue where `replication` fields would not update in `google_secret_manager_secret` ([#3558](https://github.com/hashicorp/terraform-provider-google-beta/pull/3558))
* service_usage: fixed imports on `google_service_usage_consumer_quota_override` ([#3552](https://github.com/hashicorp/terraform-provider-google-beta/pull/3552))
* sql: fixed a permadiff bug for `type` when BUILT_IN on `google_sql_user` ([#3545](https://github.com/hashicorp/terraform-provider-google-beta/pull/3545))
* sql: fixed bug in `google_sql_user` with CLOUD_IAM_USERs on POSTGRES. ([#3542](https://github.com/hashicorp/terraform-provider-google-beta/pull/3542))

## 3.81.0 (August 23, 2021)

IMPROVEMENTS:
* compute: Added `enable` attribute to `google_compute_router_peer` ([#3507](https://github.com/hashicorp/terraform-provider-google-beta/pull/3507))
* compute: added support for `L3_DEFAULT` as `ip_protocol` for `google_compute_forwarding_rule` and `UNSPECIFIED` as `protocol` for `google_compute_region_backend_service` to support network load balancers that forward all protocols and ports. ([#3516](https://github.com/hashicorp/terraform-provider-google-beta/pull/3516))
* compute: added support for `security_settings` to `google_compute_backend_service` ([#3515](https://github.com/hashicorp/terraform-provider-google-beta/pull/3515))
* gkehub: added `google_gke_hub_membership` support for both `//container.googleapis.com/${google_container_cluster.my-cluster.id}` and `google_container_cluster.my-cluster.id` in `endpoint.0.gke_cluster.0.resource_link` ([#3502](https://github.com/hashicorp/terraform-provider-google-beta/pull/3502))
* provider: Added provider support for `request_reason` ([#3513](https://github.com/hashicorp/terraform-provider-google-beta/pull/3513))
* provider: added support for `billing_project` across all resources. If `user_project_override` is set to `true` and a `billing_project` is set, the `X-Goog-User-Project` header will be sent for all resources. ([#3539](https://github.com/hashicorp/terraform-provider-google-beta/pull/3539))

BUG FIXES:
* assuredworkloads: enhanced resource deletion so `google_assured_workloads_workload` can delete what it creates ([#3533](https://github.com/hashicorp/terraform-provider-google-beta/pull/3533))
* bigquery: fixed the permadiff bug on `location` of the `google_bigquery_dataset` ([#3524](https://github.com/hashicorp/terraform-provider-google-beta/pull/3524))
* composer: fixed environment version regexp to explicitly require . (dot) instead of any character after 'preview' (example: composer-2.0.0-preview.0-airflow-2.1.1) ([#3520](https://github.com/hashicorp/terraform-provider-google-beta/pull/3520))
* compute: changed `wait_for_instances` in `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` to no longer block plan / refresh, waiting on managed instance statuses during apply instead ([#3531](https://github.com/hashicorp/terraform-provider-google-beta/pull/3531))
* compute: fixed a bug where `negative_caching_policy` cannot be set always revalidate on `google_compute_backend_service` ([#3529](https://github.com/hashicorp/terraform-provider-google-beta/pull/3529))
* compute: fixed instances where compute resource calls would have their urls appended with a redundant `/projects` after the host ([#3532](https://github.com/hashicorp/terraform-provider-google-beta/pull/3532))
* firestore: removed diff for server generated field `__name__` on `google_firestore_index` ([#3528](https://github.com/hashicorp/terraform-provider-google-beta/pull/3528))
* privateca: Fixed null for `ignore_active_certificates_on_deletion` on the imported `google_privateca_certificate_authority` ([#3511](https://github.com/hashicorp/terraform-provider-google-beta/pull/3511))
* privateca: fixed the creation of subordinate `google_privateca_certificate_authority` with `max_issuer_path_length = 0`. ([#3540](https://github.com/hashicorp/terraform-provider-google-beta/pull/3540))

## 3.80.0 (August 16, 2021)

FEATURES:
* **New Resource:** `google_dialogflow_cx_environment` ([#3488](https://github.com/hashicorp/terraform-provider-google-beta/pull/3488))

IMPROVEMENTS:
* gkehub: added support for both `//container.googleapis.com/${google_container_cluster.my-cluster.id}` and `google_container_cluster.my-cluster.id` references in `google_gke_hub_membership.endpoint.0.gke_cluster.0.resource_link` ([#3502](https://github.com/hashicorp/terraform-provider-google-beta/pull/3502))
* kms: added `name` field to `google_kms_crypto_key_version` datasource ([#3500](https://github.com/hashicorp/terraform-provider-google-beta/pull/3500))

BUG FIXES:
* apigee: fixed update behavior on `google_apigee_envgroup` ([#3489](https://github.com/hashicorp/terraform-provider-google-beta/pull/3489))
* artifact_registry: transitioned the field `format` to be case insensitive in aligning with backend behavior on `google_artifact_registry_repository` ([#3491](https://github.com/hashicorp/terraform-provider-google-beta/pull/3491))
* privateca: fixed a failure to create `google_privateca_certificate_authority` of type `SUBORDINATE` due to an invalid attempt to activate it on creation. ([#3499](https://github.com/hashicorp/terraform-provider-google-beta/pull/3499))

## 3.79.0 (August 09, 2021)

NOTES:
* spanner: The `num_nodes` field on `google_spanner_instance` will have its default removed in a future major release, and either `num_nodes` or `processing_units` will be required. ([#3479](https://github.com/hashicorp/terraform-provider-google-beta/pull/3479))

FEATURES:
* **New Resource:** `google_dialogflow_cx_entity_type` ([#3480](https://github.com/hashicorp/terraform-provider-google-beta/pull/3480))
* **New Resource:** `google_dialogflow_cx_page` ([#3461](https://github.com/hashicorp/terraform-provider-google-beta/pull/3461))

IMPROVEMENTS:
* container: added `network_config` block to `google_container_node_pool` resource ([#3472](https://github.com/hashicorp/terraform-provider-google-beta/pull/3472))
* spanner: added `processing_units` to `google_spanner_instance`. ([#3479](https://github.com/hashicorp/terraform-provider-google-beta/pull/3479))
* storage: added support for `customer_encryption` on `resource_storage_bucket_object` ([#3469](https://github.com/hashicorp/terraform-provider-google-beta/pull/3469))

## 3.78.0 (August 02, 2021)

IMPROVEMENTS:
* composer: added validation for `max_pods_per_node` field. ([#3445](https://github.com/hashicorp/terraform-provider-google-beta/pull/3445))
* servicenetworking: added support for `user_project_override` and `billing_project ` to `google_service_networking_connection` ([#3455](https://github.com/hashicorp/terraform-provider-google-beta/pull/3455))

BUG FIXES:
* storagetransfer: fixed a crash on `azure_blob_storage_data_source` for `google_storage_transfer_job` ([#3447](https://github.com/hashicorp/terraform-provider-google-beta/pull/3447))
* sql: fixed bug that wouldn't insert the `google_sql_user` in state for iam users. ([#3442](https://github.com/hashicorp/terraform-provider-google-beta/pull/3442))
* storage: fixed a crash when `azure_credentials` was defined in `google_storage_transfer_job` ([#3457](https://github.com/hashicorp/terraform-provider-google-beta/pull/3457))

## 3.77.0 (July 26, 2021)

FEATURES:
* **New Resource:** `google_scc_notification_config` ([#3431](https://github.com/hashicorp/terraform-provider-google-beta/pull/3431))

IMPROVEMENTS:
* composer: added field `maintenance_window` to resource `google_composer_environment` ([#3435](https://github.com/hashicorp/terraform-provider-google-beta/pull/3435))
* compute: fixed a permadiff bug in `log_config` field of `google_compute_region_backend_service` ([#3427](https://github.com/hashicorp/terraform-provider-google-beta/pull/3427))
* dlp: added `crypto_replace_ffx_fpe_config` and `crypto_replace_ffx_fpe_config` as primitive transformation types to `google_data_loss_prevention_deidentify_template` ([#3429](https://github.com/hashicorp/terraform-provider-google-beta/pull/3429))

BUG FIXES:
* bigquerydatatransfer: fixed a bug where `destination_dataset_id` was required, it is now optional. ([#3438](https://github.com/hashicorp/terraform-provider-google-beta/pull/3438))
* billing: Fixed ordering of `budget_filter. projects` on `google_billing_budget` ([#3436](https://github.com/hashicorp/terraform-provider-google-beta/pull/3436))
* compute: removed default value of `0.8` from `google_backend_service.backend.max_utilization` and it will now default from API. All `max_connections_xxx` and `max_rate_xxx` will also default from API as these are all conditional on balancing mode. ([#3432](https://github.com/hashicorp/terraform-provider-google-beta/pull/3432))
* sql: fixed bug where the provider would retry on an error if the database instance name couldn't be reused. ([#3434](https://github.com/hashicorp/terraform-provider-google-beta/pull/3434))

## 3.76.0 (July 19, 2021)
FEATURES:
* **New Resource:** `google_assured_workloads_workload` ([#3410](https://github.com/hashicorp/terraform-provider-google-beta/pull/3410))
* **New Resource:** `google_dialogflow_cx_flow` ([#3422](https://github.com/hashicorp/terraform-provider-google-beta/pull/3422))
* **New Resource:** `google_dialogflow_cx_intent` ([#3415](https://github.com/hashicorp/terraform-provider-google-beta/pull/3415))
* **New Resource:** `google_dialogflow_cx_version` ([#3423](https://github.com/hashicorp/terraform-provider-google-beta/pull/3423))
* **New Resource:** `google_network_services_edge_cache_keyset` ([#3417](https://github.com/hashicorp/terraform-provider-google-beta/pull/3417))
* **New Resource:** `google_network_services_edge_cache_origin` ([#3417](https://github.com/hashicorp/terraform-provider-google-beta/pull/3417))
* **New Resource:** `google_network_services_edge_cache_service` ([#3417](https://github.com/hashicorp/terraform-provider-google-beta/pull/3417))
* **New Resource:** `google_vertex_ai_featurestore_entitytype` ([#3416](https://github.com/hashicorp/terraform-provider-google-beta/pull/3416))
* **New Resource:** `google_vertex_ai_featurestore` ([#3416](https://github.com/hashicorp/terraform-provider-google-beta/pull/3416))

IMPROVEMENTS:
* apigee: Added SLASH_22 support for `peering_cidr_range` on `google_apigee_instance` ([#3424](https://github.com/hashicorp/terraform-provider-google-beta/pull/3424))
* cloudbuild: Added `pubsub_config` and `webhook_config` parameter to `google_cloudbuild_trigger`. ([#3418](https://github.com/hashicorp/terraform-provider-google-beta/pull/3418))

BUG FIXES:
* pubsub: fixed pubsublite update issues ([#3421](https://github.com/hashicorp/terraform-provider-google-beta/pull/3421))

## 3.75.0 (July 12, 2021)

BREAKING CHANGES:
* privateca: existing beta resources will no longer function ([#3397](https://github.com/hashicorp/terraform-provider-google-beta/pull/3397))

FEATURES:
* **New Resource:** google_privateca_ca_pool ([#3397](https://github.com/hashicorp/terraform-provider-google-beta/pull/3397))
* **New Resource:** google_privateca_certificate ([#3397](https://github.com/hashicorp/terraform-provider-google-beta/pull/3397))
* **New Resource:** google_privateca_certificate_authority ([#3397](https://github.com/hashicorp/terraform-provider-google-beta/pull/3397))

IMPROVEMENTS:
* bigquery: added `kms_key_version` as an output on `bigquery_table.encryption_configuration` and the `destination_encryption_configuration` blocks of `bigquery_job.query`, `bigquery_job.load`, and `bigquery_copy`. ([#3406](https://github.com/hashicorp/terraform-provider-google-beta/pull/3406))
* compute: added `advanced_machine_features` to `google_compute_instance` ([#3392](https://github.com/hashicorp/terraform-provider-google-beta/pull/3392))
* dlp: Added `replace_with_info_type_config` to `dlp_deidentify_template`. ([#3384](https://github.com/hashicorp/terraform-provider-google-beta/pull/3384))
* storage: added `temporary_hold` and `event_based_hold` attributes to `google_storage_bucket_object` ([#3399](https://github.com/hashicorp/terraform-provider-google-beta/pull/3399))

BUG FIXES:
* bigquery: Fixed permadiff due to lowercase mode/type in `google_bigquery_table.schema` ([#3405](https://github.com/hashicorp/terraform-provider-google-beta/pull/3405))
* billing: made `all_updates_rule.*` fields updatable on `google_billing_budget` ([#3394](https://github.com/hashicorp/terraform-provider-google-beta/pull/3394))
* billing: made `amount.specified_amount.units` updatable on `google_billing_budget` ([#3391](https://github.com/hashicorp/terraform-provider-google-beta/pull/3391))
* compute: fixed perma-diff in `google_compute_instance` ([#3389](https://github.com/hashicorp/terraform-provider-google-beta/pull/3389))
* storage: fixed handling of object paths that contain slashes for `google_storage_object_access_control` ([#3407](https://github.com/hashicorp/terraform-provider-google-beta/pull/3407))

## 3.74.0 (June 28, 2021)

FEATURES:
* **New Resource:** `google_app_engine_service_network_settings` ([#3371](https://github.com/hashicorp/terraform-provider-google-beta/pull/3371))
* **New Resource:** `google_vertex_ai_dataset` ([#3369](https://github.com/hashicorp/terraform-provider-google-beta/pull/3369))
* **New Resource:** `google_cloudbuild_worker_pool` ([#3372](https://github.com/hashicorp/terraform-provider-google-beta/pull/3372))

IMPROVEMENTS:
* bigtable: added `cluster.kms_key_name` field to `google_bigtable_instance` ([#3354](https://github.com/hashicorp/terraform-provider-google-beta/pull/3354))
* composer: added field `max_pods_per_node` to resource `google_composer_environment` (beta) ([#3376](https://github.com/hashicorp/terraform-provider-google-beta/pull/3376))
* secretmanager: added `ttl`, `expire_time`, `topics` and `rotation` fields to `google_secret_manager_secret` ([#3360](https://github.com/hashicorp/terraform-provider-google-beta/pull/3360))

BUG FIXES:
* container: allowed setting `node_config.service_account` at the same time as `enable_autopilot = true` for `google_container_cluster` ([#3361](https://github.com/hashicorp/terraform-provider-google-beta/pull/3361))
* container: fixed issue where creating a node pool with a name that already exists would import that resource. `google_container_node_pool` ([#3378](https://github.com/hashicorp/terraform-provider-google-beta/pull/3378))
* dataproc: fixed crash when creating `google_dataproc_workflow_template` with `secondary_worker_config` empty except for `num_instances = 0` ([#3347](https://github.com/hashicorp/terraform-provider-google-beta/pull/3347))
* filestore: fixed an issue in `google_filestore_instance` where creating two instances simultaneously resulted in an error. ([#3358](https://github.com/hashicorp/terraform-provider-google-beta/pull/3358))
* iam: fixed an issue in `google_iam_workload_identity_pool_provider` where `aws` and `oidc` were not updatable. ([#3350](https://github.com/hashicorp/terraform-provider-google-beta/pull/3350))
* sql: added support for `binary_logging` on replica instances for `googe_sql_database_instance` ([#3379](https://github.com/hashicorp/terraform-provider-google-beta/pull/3379))

## 3.73.0 (June 21, 2021)
FEATURES:
* **New Resource:** `google_compute_service_attachment` ([#3328](https://github.com/hashicorp/terraform-provider-google-beta/pull/3328))
* **New Resource:** `google_dialogflow_cx_agent` ([#3324](https://github.com/hashicorp/terraform-provider-google-beta/pull/3324))
* **New Resource:** `google_gkehub_feature` ([#3330](https://github.com/hashicorp/terraform-provider-google-beta/pull/3330))
* **New Resource:** `google_gkehub_feature_membership` ([#3330](https://github.com/hashicorp/terraform-provider-google-beta/pull/3330))

IMPROVEMENTS:
* provider: added support for [mtls authentication](https://google.aip.dev/auth/4114) ([#3348](https://github.com/hashicorp/terraform-provider-google-beta/pull/3348))
* compute: added field `adaptive_protection_config` to `google_compute_security_policy` ([#3322](https://github.com/hashicorp/terraform-provider-google-beta/pull/3322))
* compute: added `advanced_machine_features` fields to `google_compute_instance_template` ([#3337](https://github.com/hashicorp/terraform-provider-google-beta/pull/3337))
* compute: added a `network_performance_config` block to each of `resource_compute_instance`, `resource_compute_instance_from_template`, and `resource_compute_instance_template` ([#3341](https://github.com/hashicorp/terraform-provider-google-beta/pull/3341))
* redis: allowed `redis_version` to be upgraded on `google_redis_instance` ([#3344](https://github.com/hashicorp/terraform-provider-google-beta/pull/3344))

BUG FIXES:
* apigee: added SLASH_23 support for `peering_cidr_range` on `google_apigee_instance` ([#3327](https://github.com/hashicorp/terraform-provider-google-beta/pull/3327))
* cloudrun: fixed a bug where plan would should a diff on `google_cloud_run_service` if the order of the `template.spec.containers.env` list was re-ordered outside of terraform. ([#3326](https://github.com/hashicorp/terraform-provider-google-beta/pull/3326))
* container: added `user_project_override` support to the ContainerOperationWaiter used by `google_container_cluster` ([#3345](https://github.com/hashicorp/terraform-provider-google-beta/pull/3345))

## 3.72.0 (June 14, 2021)
IMPROVEMENTS:
* container: Allowed specifying a cluster id field for `google_container_node_pool.cluster` to ensure that a node pool is recreated if the associated cluster is recreated. ([#3314](https://github.com/hashicorp/terraform-provider-google-beta/pull/3314))
* storagetransfer: added support for `azure_blob_storage_data_source` to `google_storage_transfer_job` ([#3316](https://github.com/hashicorp/terraform-provider-google-beta/pull/3316))

BUG FIXES:
* bigquery: Fixed `google_bigquery_table.schema` handling of policyTags ([#3307](https://github.com/hashicorp/terraform-provider-google-beta/pull/3307))
* bigtable: fixed bug that would error if creating multiple bigtable gc policies at the same time ([#3311](https://github.com/hashicorp/terraform-provider-google-beta/pull/3311))
* compute: fixed bug where `encryption` showed a perma-diff on resources created prior to the feature being released. ([#3309](https://github.com/hashicorp/terraform-provider-google-beta/pull/3309))
* dataflow: fixed handling of failed `google_dataflow_flex_template_job` updates ([#3318](https://github.com/hashicorp/terraform-provider-google-beta/pull/3318))
* dataflow: made `google_dataflow_flex_template_job` updates fail fast if the job is in the process of cancelling or draining([#3317](https://github.com/hashicorp/terraform-provider-google-beta/pull/3317))

## 3.71.0 (June 07, 2021)
FEATURES:
* **New Resource:** `google_dialogflow_fulfillment` ([#3286](https://github.com/hashicorp/terraform-provider-google-beta/pull/3286))

IMPROVEMENTS:
* compute: added `reservation_affinity` to `google_compute_instance` and `google_compute_instance_template` ([#3288](https://github.com/hashicorp/terraform-provider-google-beta/pull/3288))
* compute: added support for `wait_for_instances_status` on `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` ([#3283](https://github.com/hashicorp/terraform-provider-google-beta/pull/3283))
* compute: added support for output-only `status` field on `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` ([#3283](https://github.com/hashicorp/terraform-provider-google-beta/pull/3283))
* compute: set the default value for log_config.enable on `google_compute_region_health_check` to avoid permanent diff on plan/apply. ([#3291](https://github.com/hashicorp/terraform-provider-google-beta/pull/3291))

BUG FIXES:
* composer: fixed a check that did not allow for preview versions in `google_composer_environment` ([#3287](https://github.com/hashicorp/terraform-provider-google-beta/pull/3287))
* storage: fixed error when `matches_storage_class` is set empty on `google_storage_bucket` ([#3282](https://github.com/hashicorp/terraform-provider-google-beta/pull/3282))
* vpcaccess: fixed permadiff when `max_throughput` is not set on `google_vpc_access_connector` ([#3294](https://github.com/hashicorp/terraform-provider-google-beta/pull/3294))

## 3.70.0 (June 01, 2021)
IMPROVEMENTS:
* compute: added `provisioned_iops` to `google_compute_disk` ([#3269](https://github.com/hashicorp/terraform-provider-google-beta/pull/3269))
* sql: added field `disk_autoresize_limit` to `sql_database_instance` ([#3273](https://github.com/hashicorp/terraform-provider-google-beta/pull/3273))

BUG FIXES:
* cloudrun: fixed a bug where resources would return successfully due to responses based on a previous version of the resource ([#3277](https://github.com/hashicorp/terraform-provider-google-beta/pull/3277))
* compute: fixed issue where `google_compute_region_disk` and `google_compute_disk` would force recreation due to the addition of `interface` property ([#3272](https://github.com/hashicorp/terraform-provider-google-beta/pull/3272))
* compute: fixed missing values for `negative_caching` and `serve_while_stale` on `google_compute_backend_service` ([#3278](https://github.com/hashicorp/terraform-provider-google-beta/pull/3278))
* storage: fixed error when `matches_storage_class` is set empty on `google_storage_bucket` ([#3282](https://github.com/hashicorp/terraform-provider-google-beta/pull/3282))

## 3.69.0 (May 24, 2021)

IMPROVEMENTS:
* apigateway: allowed field `apiconfig` to change on resource `google_apigateway_gateway` ([#3248](https://github.com/hashicorp/terraform-provider-google-beta/pull/3248))
* compute: added "description" field to "google_compute_resource_policy" resource ([#3263](https://github.com/hashicorp/terraform-provider-google-beta/pull/3263))
* compute: added "instance_schedule_policy" field to "google_compute_resource_policy" resource ([#3263](https://github.com/hashicorp/terraform-provider-google-beta/pull/3263))
* compute: added support for IPsec-encrypted Interconnect in the form of new fields on `google_compute_router`, `google_compute_ha_vpn_gateway`, `google_compute_interconnect_attachment` and `google_compute_address` ([#3256](https://github.com/hashicorp/terraform-provider-google-beta/pull/3256))
* dataflow: enabled updates for `google_dataflow_flex_template_job` ([#3246](https://github.com/hashicorp/terraform-provider-google-beta/pull/3246))

BUG FIXES:
* cloudidentity: fixed recreation on the `initial_group_config` of `google_cloud_identity_group` ([#3252](https://github.com/hashicorp/terraform-provider-google-beta/pull/3252))
* compute: added mutex in `google_compute_metadata_item` to reduce retries + quota errors ([#3262](https://github.com/hashicorp/terraform-provider-google-beta/pull/3262))
* container: fixed bug where `enable_shielded_nodes` could not be false on resource `google_container_cluster` ([#3247](https://github.com/hashicorp/terraform-provider-google-beta/pull/3247))

## 3.68.0 (May 18, 2021)
FEATURES:
* **New Resource:** `google_pubsub_schema` ([#3243](https://github.com/hashicorp/terraform-provider-google-beta/pull/3243))

IMPROVEMENTS:
* compute: added `initial_size` in resource `google_compute_node_group` to account for scenarios where size may change under the hood ([#3228](https://github.com/hashicorp/terraform-provider-google-beta/pull/3228))
* compute: added support for setting `kms_key_name` on `google_compute_machine_image` ([#3241](https://github.com/hashicorp/terraform-provider-google-beta/pull/3241))
* dataflow: enabled updates for `google_dataflow_flex_template_job` ([#3246](https://github.com/hashicorp/terraform-provider-google-beta/pull/3246))

BUG FIXES:
* compute: Fixed permadiff for `cdn_policy.serve_while_stale` and `cdn_policy.*_ttl` in `google_compute_region_backend_service` (beta) ([#3230](https://github.com/hashicorp/terraform-provider-google-beta/pull/3230))
* compute: fixed bug where, when an organization security policy association was removed outside of terraform, the next plan/apply would fail. ([#3234](https://github.com/hashicorp/terraform-provider-google-beta/pull/3234))
* container: added validation to check that both `node_version` and `remove_default_node_pool` cannot be set on `google_container_cluster` ([#3237](https://github.com/hashicorp/terraform-provider-google-beta/pull/3237))
* dns: suppressed spurious diffs due to case changes in DS records ([#3236](https://github.com/hashicorp/terraform-provider-google-beta/pull/3236))

## 3.67.0 (May 10, 2021)
NOTES:
* all: changed default HTTP request timeout from 30 seconds to 120 seconds ([#3181](https://github.com/hashicorp/terraform-provider-google-beta/pull/3181))
BREAKING CHANGES:
* bigquery: updating `dataset_id` or `project_id` in `google_bigquery_dataset` will now recreate the resource ([#3185](https://github.com/hashicorp/terraform-provider-google-beta/pull/3185))

IMPROVEMENTS:
* accesscontextmanager: added support for `require_verified_chrome_os` in basic access levels. ([#3223](https://github.com/hashicorp/terraform-provider-google-beta/pull/3223))
* billingbudget: added support for import of `google_billing_budget` ([#3194](https://github.com/hashicorp/terraform-provider-google-beta/pull/3194))
* cloud_identity: added support for `initial_group_config` to the google_cloud_identity_group resource ([#3211](https://github.com/hashicorp/terraform-provider-google-beta/pull/3211))
* cloudrun: added support to bind secrets from Secret Manager to environment variables or files to `google_cloud_run_service` ([#3225](https://github.com/hashicorp/terraform-provider-google-beta/pull/3225))
* compute: added `initial_size` to account for scenarios where size may change under the hood in resource `google_compute_node_group` ([#3228](https://github.com/hashicorp/terraform-provider-google-beta/pull/3228))
* compute: added `interface` field to `google_compute_region_disk` ([#3193](https://github.com/hashicorp/terraform-provider-google-beta/pull/3193))
* healthcare: added support for `stream_configs` in `google_healthcare_dicom_store` ([#3190](https://github.com/hashicorp/terraform-provider-google-beta/pull/3190))
* secretmanager: added support for setting a CMEK on `google_secret_manager_secret` ([#3212](https://github.com/hashicorp/terraform-provider-google-beta/pull/3212))
* spanner: added `force_destroy` to `google_spanner_instance` to delete instances that have backups enabled. ([#3227](https://github.com/hashicorp/terraform-provider-google-beta/pull/3227))
* spanner: added support for setting a CMEK on `google_spanner_database` ([#3181](https://github.com/hashicorp/terraform-provider-google-beta/pull/3181))
* workflows: marked `source_contents` and `service_account` as updatable on `google_workflows_workflow` ([#3205](https://github.com/hashicorp/terraform-provider-google-beta/pull/3205))

BUG FIXES:
* bigquery: fixed `dataset_id` to force new resource if name is changed. ([#3185](https://github.com/hashicorp/terraform-provider-google-beta/pull/3185))
* cloudrun: fixed permadiff on `google_cloud_run_domain_mapping.metadata.labels` ([#3183](https://github.com/hashicorp/terraform-provider-google-beta/pull/3183))
* composer: changed `google_composer_environment.master_ipv4_cidr_block` to draw default from the API ([#3204](https://github.com/hashicorp/terraform-provider-google-beta/pull/3204))
* compute: fixed the failure when `min_required_replicas` is set to 0 on `google_compute_autoscaler` or `google_compute_region_autoscaler` ([#3203](https://github.com/hashicorp/terraform-provider-google-beta/pull/3203))
* container: fixed container node pool not removed from the state when received 404 error on delete call for the resource `google_container_node_pool` ([#3210](https://github.com/hashicorp/terraform-provider-google-beta/pull/3210))
* dns: fixed empty `rrdatas` list on `google_dns_record_set` for AAAA records ([#3207](https://github.com/hashicorp/terraform-provider-google-beta/pull/3207))
* kms: fixed indirectly force replacement via `skip_initial_version_creation` on `google_kms_crypto_key` ([#3192](https://github.com/hashicorp/terraform-provider-google-beta/pull/3192))
* logging: fixed `metric_descriptor.labels` can't be updated on 'google_logging_metric' ([#3217](https://github.com/hashicorp/terraform-provider-google-beta/pull/3217))
* pubsub: fixed diff for `minimum_backoff & maximum_backoff` on `google_pubsub_subscription` ([#3214](https://github.com/hashicorp/terraform-provider-google-beta/pull/3214))
* resourcemanager: fixed broken handling of IAM conditions for `google_organization_iam_member`, `google_organization_iam_binding`, and `google_organization_iam_policy` ([#3213](https://github.com/hashicorp/terraform-provider-google-beta/pull/3213))
* serviceusage: added `google_project_service.service` validation to reject invalid service domains that don't contain a period ([#3191](https://github.com/hashicorp/terraform-provider-google-beta/pull/3191))
* storage: fixed bug where `role_entity` user wouldn't update if the role changed. ([#3199](https://github.com/hashicorp/terraform-provider-google-beta/pull/3199))

## 3.66.1 (April 29, 2021)
BUG FIXES:
* compute: fixed bug where terraform would crash if updating from no `service_account.scopes` to more. ([#3208](https://github.com/hashicorp/terraform-provider-google-beta/pull/3208))

## 3.66.0 (April 28, 2021)

NOTES:
* all: changed default HTTP request timeout from 30 seconds to 120 seconds ([#3181](https://github.com/hashicorp/terraform-provider-google-beta/pull/3181))

BREAKING CHANGES:
* datacatalog: updating `parent` in `google_data_catalog_tag` will now recreate the resource ([#3179](https://github.com/hashicorp/terraform-provider-google-beta/pull/3179))

FEATURES:
* **New Data Source:** `google_compute_ha_vpn_gateway` ([#3173](https://github.com/hashicorp/terraform-provider-google-beta/pull/3173))
* **New Resource:** `google_dataproc_workflow_template` ([#3178](https://github.com/hashicorp/terraform-provider-google-beta/pull/3178))

IMPROVEMENTS:
* bigquery: Added BigTable source format in BigQuery table ([#3165](https://github.com/hashicorp/terraform-provider-google-beta/pull/3165))
* cloudfunctions: removed bounds on the supported memory range in `google_cloudfunctions_function.available_memory_mb` ([#3171](https://github.com/hashicorp/terraform-provider-google-beta/pull/3171))
* compute: marked scheduling.0.node_affinities as updatable in `google_compute_instance` ([#3166](https://github.com/hashicorp/terraform-provider-google-beta/pull/3166))
* dataproc: added `shielded_instance_config` fields to `google_dataproc_cluster` ([#3157](https://github.com/hashicorp/terraform-provider-google-beta/pull/3157))
* spanner: added support for setting a CMEK on `google_spanner_database` ([#3181](https://github.com/hashicorp/terraform-provider-google-beta/pull/3181))

BUG FIXES:
* compute: fixed error when creating empty `scopes` on `google_compute_instance` ([#3174](https://github.com/hashicorp/terraform-provider-google-beta/pull/3174))
* container: fixed a bug that allowed specifying `node_config` on `google_container_cluster` when autopilot is used ([#3155](https://github.com/hashicorp/terraform-provider-google-beta/pull/3155))
* datacatalog: fixed an issue where `parent` in `google_data_catalog_tag` attempted to update the resource when change instead of recreating it ([#3179](https://github.com/hashicorp/terraform-provider-google-beta/pull/3179))
* datacatalog: set default false for `force_delete` on `google_data_catalog_tag_template` ([#3164](https://github.com/hashicorp/terraform-provider-google-beta/pull/3164))
* dns: added missing record types to `google_dns_record_set` resource ([#3160](https://github.com/hashicorp/terraform-provider-google-beta/pull/3160))
* sql: set `clone.point_in_time` optional for `google_sql_database_instance` ([#3180](https://github.com/hashicorp/terraform-provider-google-beta/pull/3180))

## 3.65.0 (April 20, 2021)

FEATURES:
* **New Data Source:** `google_kms_secret_asymmetric` ([#3141](https://github.com/hashicorp/terraform-provider-google-beta/pull/3141))

IMPROVEMENTS:
* compute: added the ability to specify `google_compute_forwarding_rule.ip_address` by a reference in addition to raw IP address ([#3140](https://github.com/hashicorp/terraform-provider-google-beta/pull/3140))
* compute: enabled fields `advertiseMode`, `advertisedGroups`, `peerAsn`, and `peerIpAddress` to be updatable on resource `google_compute_router_peer` ([#3134](https://github.com/hashicorp/terraform-provider-google-beta/pull/3134))
* eventarc: added `transport.pubsub.topic` to `google_eventarc_trigger` ([#3149](https://github.com/hashicorp/terraform-provider-google-beta/pull/3149))

BUG FIXES:
* cloud_identity: fixed google_cloud_identity_group_membership import/update ([#3136](https://github.com/hashicorp/terraform-provider-google-beta/pull/3136))
* compute: removed minimum for `scopes` field on `google_compute_instance` resource ([#3147](https://github.com/hashicorp/terraform-provider-google-beta/pull/3147))
* iam: fixed issue with principle and principleSet members not retaining their casing ([#3133](https://github.com/hashicorp/terraform-provider-google-beta/pull/3133))
* workflows: fixed a bug in `google_workflows_workflow` that could cause inconsistent final plan errors when using the `name` field in other resources ([#3138](https://github.com/hashicorp/terraform-provider-google-beta/pull/3138))

## 3.64.0 (April 12, 2021)

FEATURES:
* **New Resource:** `google_tags_tag_binding` ([#3121](https://github.com/hashicorp/terraform-provider-google-beta/pull/3121))
* **New Resource:** `google_tags_tag_key_iam_binding` ([#3124](https://github.com/hashicorp/terraform-provider-google-beta/pull/3124))
* **New Resource:** `google_tags_tag_key_iam_member` ([#3124](https://github.com/hashicorp/terraform-provider-google-beta/pull/3124))
* **New Resource:** `google_tags_tag_key_iam_policy` ([#3124](https://github.com/hashicorp/terraform-provider-google-beta/pull/3124))
* **New Resource:** `google_tags_tag_value_iam_binding` ([#3124](https://github.com/hashicorp/terraform-provider-google-beta/pull/3124))
* **New Resource:** `google_tags_tag_value_iam_member` ([#3124](https://github.com/hashicorp/terraform-provider-google-beta/pull/3124))
* **New Resource:** `google_tags_tag_value_iam_policy` ([#3124](https://github.com/hashicorp/terraform-provider-google-beta/pull/3124))
* **New Resource:** `google_apigee_envgroup_attachment` ([#3129](https://github.com/hashicorp/terraform-provider-google-beta/pull/3129))

IMPROVEMENTS:
* bigquery: added `require_partition_filter` field to `google_bigquery_table` when provisioning `hive_partitioning_options` ([#3106](https://github.com/hashicorp/terraform-provider-google-beta/pull/3106))
* cloudbuild: added new machine types for `google_cloudbuild_trigger` ([#3115](https://github.com/hashicorp/terraform-provider-google-beta/pull/3115))
* compute: added field `maintenance_window.start_time` to `google_compute_node_group` ([#3125](https://github.com/hashicorp/terraform-provider-google-beta/pull/3125))
* compute: added gVNIC support for `google_compute_instance_template` ([#3123](https://github.com/hashicorp/terraform-provider-google-beta/pull/3123))
* datacatalog: added `description` field to `google_data_catalog_tag_template ` resource ([#3128](https://github.com/hashicorp/terraform-provider-google-beta/pull/3128))
* iam: added support for third party identities via the principle and principleSet IAM members ([#3133](https://github.com/hashicorp/terraform-provider-google-beta/pull/3133))

BUG FIXES:
* compute: reverted datatype change for `mtu` in `google_compute_interconnect_attachment` as it was incompatible with existing state representation ([#3112](https://github.com/hashicorp/terraform-provider-google-beta/pull/3112))
* iam: fixed issue with principle and principleSet members not retaining their casing ([#3133](https://github.com/hashicorp/terraform-provider-google-beta/pull/3133))
* storage: fixed intermittent `Provider produced inconsistent result after apply` error when creating ([#3107](https://github.com/hashicorp/terraform-provider-google-beta/pull/3107))

## 3.63.0 (April 5, 2021)

FEATURES:
* **New Data Source:** `google_monitoring_istio_canonical_service` ([#3092](https://github.com/hashicorp/terraform-provider-google-beta/pull/3092))
* **New Resource:** `google_apigee_instance_attachment` ([#3093](https://github.com/hashicorp/terraform-provider-google-beta/pull/3093))
* **New Resource:** `google_gke_hub_membership` ([#3079](https://github.com/hashicorp/terraform-provider-google-beta/pull/3079))
* **New Resource:** `google_tags_tag_value` ([#3097](https://github.com/hashicorp/terraform-provider-google-beta/pull/3097))

IMPROVEMENTS:
* added support for Apple silicon chip (updated to go 1.16) ([#3057](https://github.com/hashicorp/terraform-provider-google-beta/pull/3057))
* container: 
  * added support for GKE Autopilot in `google_container_cluster`([#3101](https://github.com/hashicorp/terraform-provider-google-beta/pull/3101))
  * added `enable_l4_ilb_subsetting` (beta) and `private_ipv6_google_access` fields to `google_container_cluster` ([#3095](https://github.com/hashicorp/terraform-provider-google-beta/pull/3095))
* sql: changed the default timeout of `google_sql_database_instance` to 30m from 20m ([#3099](https://github.com/hashicorp/terraform-provider-google-beta/pull/3099))

BUG FIXES:
* bigquery: fixed issue where you couldn't extend an existing `schema` with additional columns in `google_bigquery_table` ([#3100](https://github.com/hashicorp/terraform-provider-google-beta/pull/3100))
* cloudidentity: modified `google_cloud_identity_groups` and `google_cloud_identity_group_memberships ` to respect the `user_project_override` and `billing_project` configurations and send the appropriate headers to establish a quota project ([#3081](https://github.com/hashicorp/terraform-provider-google-beta/pull/3081))
* compute: added minimum for `scopes` field to `google_compute_instance` resource ([#3098](https://github.com/hashicorp/terraform-provider-google-beta/pull/3098))
* notebooks: fixed permadiff on labels for `google_notebook_instance` ([#3096](https://github.com/hashicorp/terraform-provider-google-beta/pull/3096))
* secretmanager: set required on `secrest_data` in `google_secret_manager_secret_version` ([#3094](https://github.com/hashicorp/terraform-provider-google-beta/pull/3094))


## 3.62.0 (March 27, 2021)

FEATURES:
* **New Data Source:** `google_compute_health_check` ([#3066](https://github.com/hashicorp/terraform-provider-google-beta/pull/3066))
* **New Data Source:** `google_kms_secret_asymmetric` ([#3076](https://github.com/hashicorp/terraform-provider-google-beta/pull/3076))
* **New Resource:** `google_gke_hub_membership` ([#3079](https://github.com/hashicorp/terraform-provider-google-beta/pull/3079))
* **New Resource:** `google_tags_tag_key` ([#3062](https://github.com/hashicorp/terraform-provider-google-beta/pull/3062))
* **New Resource:** `google_data_catalog_tag_template_iam_*` ([#3071](https://github.com/hashicorp/terraform-provider-google-beta/pull/3071))

IMPROVEMENTS:
* accesscontextmanager: added support for ingress and egress policies to `google_access_context_manager_service_perimeter` ([#3064](https://github.com/hashicorp/terraform-provider-google-beta/pull/3064))
* artifactregistry: relaxed field validations for field `format` on `google_artifact_registry_repository` ([#3068](https://github.com/hashicorp/terraform-provider-google-beta/pull/3068))
* compute: added `proxy_bind` to `google_compute_target_tcp_proxy`, `google_compute_target_http_proxy` and `google_compute_target_https_proxy` ([#3061](https://github.com/hashicorp/terraform-provider-google-beta/pull/3061))

BUG FIXES:
* compute: fixed an issue where exceeding the operation rate limit would fail without retrying ([#3077](https://github.com/hashicorp/terraform-provider-google-beta/pull/3077))
* compute: corrected underlying type to integer for field `mtu` in `google_compute_interconnect_attachment` ([#3075](https://github.com/hashicorp/terraform-provider-google-beta/pull/3075)

## 3.61.0 (March 23, 2021)

IMPROVEMENTS:
* provider: The provider now supports [Workload Identity Federation](https://cloud.google.com/iam/docs/workload-identity-federation). The federated json credentials must be loaded through the `GOOGLE_APPLICATION_CREDENTIALS` environment variable. ([#3054](https://github.com/hashicorp/terraform-provider-google-beta/pull/3054))
* compute: added `proxy_bind` to `google_compute_target_tcp_proxy`, `google_compute_target_http_proxy` and `google_compute_target_https_proxy` ([#3061](https://github.com/hashicorp/terraform-provider-google-beta/pull/3061))
* compute: changed `google_compute_subnetwork` to accept more values in the `purpose` field ([#3043](https://github.com/hashicorp/terraform-provider-google-beta/pull/3043))
* dataflow: added `enable_streaming_engine` argument ([#3049](https://github.com/hashicorp/terraform-provider-google-beta/pull/3049))
* vpcaccess: added `subnet`, `machine_type` beta fields to `google_vpc_access_connector` ([#3042](https://github.com/hashicorp/terraform-provider-google-beta/pull/3042))

BUG FIXES:
* bigtable: fixed bug where gc_policy would attempt to recreate the resource when switching from deprecated attribute but maintaining the same value underlying value ([#3037](https://github.com/hashicorp/terraform-provider-google-beta/pull/3037))
* binaryauthorization: fixed permadiff in `google_binary_authorization_attestor` ([#3035](https://github.com/hashicorp/terraform-provider-google-beta/pull/3035))
* container: Fixed updates on `export_custom_routes` and `import_custom_routes` in `google_compute_network_peering` ([#3045](https://github.com/hashicorp/terraform-provider-google-beta/pull/3045))

## 3.60.0 (March 15, 2021)

FEATURES:
* **New Resource:** `google_workflows_workflow` ([#2989](https://github.com/hashicorp/terraform-provider-google-beta/pull/2989))
* **New Resource:** google_apigee_envgroup ([#3039](https://github.com/hashicorp/terraform-provider-google-beta/pull/3039))
* **New Resource:** google_apigee_environment ([#3020](https://github.com/hashicorp/terraform-provider-google-beta/pull/3020))
* **New Resource:** google_apigee_instance ([#2986](https://github.com/hashicorp/terraform-provider-google-beta/pull/2986))

IMPROVEMENTS:
* cloudrun: suppressed metadata.labels["cloud.googleapis.com/location"] value in `google_cloud_run_service` ([#3005](https://github.com/hashicorp/terraform-provider-google-beta/pull/3005))
* compute: added `mtu` field to `google_compute_interconnect_attachment` ([#3006](https://github.com/hashicorp/terraform-provider-google-beta/pull/3006))
* compute: added autoscaling_policy.cpu_utilization.predictive_method field to `google_compute_autoscaler` and `google_compute_region_autoscaler` ([#2987](https://github.com/hashicorp/terraform-provider-google-beta/pull/2987))
* compute: added support for `nic_type` to `google_compute_instance` (GA only) ([#2998](https://github.com/hashicorp/terraform-provider-google-beta/pull/2998))
* container: added field `ephemeral_storage_config` to resource `google_container_node_pool` and `google_container_cluster` (beta) ([#3023](https://github.com/hashicorp/terraform-provider-google-beta/pull/3023))
* datafusion: added support for the `DEVELOPER` instance type to `google_data_fusion_instance`  ([#3015](https://github.com/hashicorp/terraform-provider-google-beta/pull/3015))
* monitoring: added windows based availability sli to the resource `google_monitoring_slo` ([#3013](https://github.com/hashicorp/terraform-provider-google-beta/pull/3013))
* sql: added `settings.0.backup_configuration.transaction_log_retention_days` and `settings.0.backup_configuration.transaction_log_retention_days` fields to `google_sql_database_instance` ([#3010](https://github.com/hashicorp/terraform-provider-google-beta/pull/3010))
* storage: added `kms_key_name` to `google_storage_bucket_object` resource ([#3026](https://github.com/hashicorp/terraform-provider-google-beta/pull/3026))

BUG FIXES:
* bigquery: fixed materialized view to be recreated when query changes ([#3032](https://github.com/hashicorp/terraform-provider-google-beta/pull/3032))
* bigtable: fixed bug where gc_policy would attempt to recreate the resource when switching from deprecated attribute but maintaining the same underlying value ([#3037](https://github.com/hashicorp/terraform-provider-google-beta/pull/3037))
* bigtable: required resource recreation if any fields change on `resource_bigtable_gc_policy` ([#2991](https://github.com/hashicorp/terraform-provider-google-beta/pull/2991))
* binaryauthorization: fixed permadiff in `google_binary_authorization_attestor` ([#3035](https://github.com/hashicorp/terraform-provider-google-beta/pull/3035))
* cloudfunction: added retry logic for `google_cloudfunctions_function` updates ([#2992](https://github.com/hashicorp/terraform-provider-google-beta/pull/2992))
* cloudidentity: fixed a bug where `google_cloud_identity_group` would periodically fail with a 403 ([#3012](https://github.com/hashicorp/terraform-provider-google-beta/pull/3012))
* compute: fixed a perma-diff for `nat_ips` that were specified as short forms in `google_compute_router_nat` ([#3007](https://github.com/hashicorp/terraform-provider-google-beta/pull/3007))
* compute: fixed perma-diff for cos-family disk images ([#3024](https://github.com/hashicorp/terraform-provider-google-beta/pull/3024))
* compute: Fixed service account scope alias to be updated. ([#3021](https://github.com/hashicorp/terraform-provider-google-beta/pull/3021))
* container: fixed container cluster not removed from the state when received 404 error on delete call for the resource `google_container_cluster` ([#3018](https://github.com/hashicorp/terraform-provider-google-beta/pull/3018))
* container: Fixed failure in deleting `maintenance_exclusion` for `google_container_cluster` ([#3014](https://github.com/hashicorp/terraform-provider-google-beta/pull/3014))
* container: fixed an issue where release channel UNSPECIFIED could not be set ([#3019](https://github.com/hashicorp/terraform-provider-google-beta/pull/3019))
* essentialcontacts: made `language_tag` required for `google_essential_contacts_contact` ([#2994](https://github.com/hashicorp/terraform-provider-google-beta/pull/2994))
* serviceusage: fixed an issue in `google_service_usage_consumer_quota_override` where setting the `override_value` to 0 would result in a permanent diff ([#2985](https://github.com/hashicorp/terraform-provider-google-beta/pull/2985))
* serviceusage: fixed an issue in `google_service_usage_consumer_quota_override` where setting the `override_value` to 0 would result in a permanent diff ([#3025](https://github.com/hashicorp/terraform-provider-google-beta/pull/3025))

## 3.59.0 (March 08, 2021)

FEATURES:
* **New Resource:** `google_dataproc_metastore_service` ([#2977](https://github.com/hashicorp/terraform-provider-google-beta/pull/2977))
* **New Resource:** `google_workflows_workflow` ([#2989](https://github.com/hashicorp/terraform-provider-google-beta/pull/2989))
* **New Resource:** `google_apigee_instance` ([#2986](https://github.com/hashicorp/terraform-provider-google-beta/pull/2986))
* **New Resource:** `google_eventarc_trigger` ([#2972](https://github.com/hashicorp/terraform-provider-google-beta/pull/2972))

IMPROVEMENTS:
* composer: added `encryption_config`  to `google_composer_environment` resource ([#2967](https://github.com/hashicorp/terraform-provider-google-beta/pull/2967))
* compute: Added graceful termination to `google_container_node_pool` create calls so that partially created node pools will resume the original operation if the Terraform process is killed mid create. ([#2969](https://github.com/hashicorp/terraform-provider-google-beta/pull/2969))
* redis : marked `auth_string` on the `resource_redis_instance` resource as sensitive ([#2974](https://github.com/hashicorp/terraform-provider-google-beta/pull/2974))

BUG FIXES:
* apigee: fixed IDs when importing `google_apigee_organization` resource ([#2966](https://github.com/hashicorp/terraform-provider-google-beta/pull/2966))
* artifactregistry: fixed issue where updating `google_artifact_registry_repository` always failed ([#2968](https://github.com/hashicorp/terraform-provider-google-beta/pull/2968))
* compute : fixed a bug where `guest_flush` could not be set to false for the resource `google_compute_resource_policy` ([#2975](https://github.com/hashicorp/terraform-provider-google-beta/pull/2975))
* compute: fixed a panic on empty `target_size` in `google_compute_region_instance_group_manager` ([#2979](https://github.com/hashicorp/terraform-provider-google-beta/pull/2979))
* redis: fixed invalid value error on `auth_string` in `google_redis_instance` ([#2970](https://github.com/hashicorp/terraform-provider-google-beta/pull/2970))

## 3.58.0 (February 23, 2021)

NOTES:
* `google_bigquery_table` resources now cannot be destroyed unless `deletion_protection = false` is set in state for the resource. ([#2954](https://github.com/hashicorp/terraform-provider-google-beta/pull/2954))

FEATURES:
* **New Data Source:** `google_runtimeconfig_variable` ([#2945](https://github.com/hashicorp/terraform-provider-google-beta/pull/2945))
* **New Data Source:** `google_iap_client` ([#2951](https://github.com/hashicorp/terraform-provider-google-beta/pull/2951))

IMPROVEMENTS:
* bigquery: added `deletion_protection` field to `google_bigquery_table` to make deleting them require an explicit intent. ([#2954](https://github.com/hashicorp/terraform-provider-google-beta/pull/2954))
* cloudrun: updated retry logic to attempt to retry 409 errors from the Cloud Run API, which may be returned intermittently on create. ([#2948](https://github.com/hashicorp/terraform-provider-google-beta/pull/2948))
* compute: removed max items limit from `google_compute_target_ssl_proxy`. The API currently allows upto 15 Certificates. ([#2964](https://github.com/hashicorp/terraform-provider-google-beta/pull/2964))
* compute: added support for Private Services Connect for Google APIs in `google_compute_global_address` and `google_compute_global_forwarding_rule` ([#2956](https://github.com/hashicorp/terraform-provider-google-beta/pull/2956))
* iam: added a retry condition that retries editing `iam_binding` and `iam_member` resources on policies that have frequently deleted service accounts ([#2963](https://github.com/hashicorp/terraform-provider-google-beta/pull/2963))
* redis: added transit encryption mode support for `google_redis_instance` ([#2955](https://github.com/hashicorp/terraform-provider-google-beta/pull/2955))
* secretmanager: changed endpoint to use v1 instead of v1beta1 as it is more up-to-date ([#2946](https://github.com/hashicorp/terraform-provider-google-beta/pull/2946))
* sql: added `insights_config` block to `google_sql_database_instance` resource ([#2944](https://github.com/hashicorp/terraform-provider-google-beta/pull/2944))

BUG FIXES:
* compute: fixed an issue where the provider could return an error on a successful delete operation ([#2958](https://github.com/hashicorp/terraform-provider-google-beta/pull/2958))
* datacatalog: fixed import issue for `google_data_catalog_taxonomy` ([#2961](https://github.com/hashicorp/terraform-provider-google-beta/pull/2961))
* dataproc : fixed `max_failure_per_hour` not sent in API request for the resource `google_dataproc_job` ([#2949](https://github.com/hashicorp/terraform-provider-google-beta/pull/2949))
* dlp : modified `google_data_loss_prevention_stored_info_type` `regex.group_indexes` field to trigger resource recreation on update ([#2947](https://github.com/hashicorp/terraform-provider-google-beta/pull/2947))
* sql: fixed diffs based on case for `charset` in `google_sql_database` ([#2957](https://github.com/hashicorp/terraform-provider-google-beta/pull/2957))

## 3.57.0 (February 16, 2021)

DEPRECATIONS:
* compute: deprecated `source_disk_url` field in `google_compute_snapshot`. ([#2939](https://github.com/hashicorp/terraform-provider-google-beta/pull/2939))
* kms: deprecated `self_link` field in `google_kms_keyring` and `google_kms_cryptokey` resource as it is identical value to `id` field. ([#2939](https://github.com/hashicorp/terraform-provider-google-beta/pull/2939))
* pubsub: deprecated `path` field in `google_pubsub_subscription` resource as it is identical value to `id` field. ([#2939](https://github.com/hashicorp/terraform-provider-google-beta/pull/2939))

FEATURES:
* **New Resource:** `google_essential_contacts_contact` ([#2943](https://github.com/hashicorp/terraform-provider-google-beta/pull/2943))
* **New Resource:** `google_privateca_certificate` ([#2924](https://github.com/hashicorp/terraform-provider-google-beta/pull/2924))

IMPROVEMENTS:
* bigquery: added `status` field to `google_bigquery_job` ([#2926](https://github.com/hashicorp/terraform-provider-google-beta/pull/2926))
* compute: added `disk.resource_policies` field to resource `google_compute_instance_template` ([#2929](https://github.com/hashicorp/terraform-provider-google-beta/pull/2929))
* compute: added `nic_type` field to `google_compute_instance_template ` resource to support gVNIC ([#2941](https://github.com/hashicorp/terraform-provider-google-beta/pull/2941))
* compute: added `nic_type` field to `google_compute_instance` resource to support gVNIC ([#2941](https://github.com/hashicorp/terraform-provider-google-beta/pull/2941))
* pubsub: marked `kms_key_name` field in `google_pubsub_topic` as updatable ([#2942](https://github.com/hashicorp/terraform-provider-google-beta/pull/2942))

BUG FIXES:
* appengine: added retry for P4SA propagation delay ([#2938](https://github.com/hashicorp/terraform-provider-google-beta/pull/2938))
* compute: fixed overly-aggressive detection of changes to google_compute_security_policy rules ([#2940](https://github.com/hashicorp/terraform-provider-google-beta/pull/2940))

## 3.56.0 (February 8, 2021)

FEATURES:
* **New Resource:** `google_privateca_certificate` ([#2924](https://github.com/hashicorp/terraform-provider-google-beta/pull/2924))

IMPROVEMENTS:
* all: added plan time validations for fields that expect base64 values. ([#2906](https://github.com/hashicorp/terraform-provider-google-beta/pull/2906))
* compute: added `disk.resource_policies` field to resource `google_compute_instance_template` ([#2929](https://github.com/hashicorp/terraform-provider-google-beta/pull/2929))
* sql: added support for point-in-time-recovery to `google_sql_database_instance` ([#2923](https://github.com/hashicorp/terraform-provider-google-beta/pull/2923))
* monitoring : added `availability` sli metric support for the resource `google_monitoring_slo` ([#2908](https://github.com/hashicorp/terraform-provider-google-beta/pull/2908))

BUG FIXES:
* bigquery: fixed bug where you could not reorder columns on `schema` for resource `google_bigquery_table` ([#2913](https://github.com/hashicorp/terraform-provider-google-beta/pull/2913))
* cloudrun: suppressed `run.googleapis.com/ingress-status` annotation in `google_cloud_run_service` ([#2920](https://github.com/hashicorp/terraform-provider-google-beta/pull/2920))
* serviceaccount: loosened restrictions on `account_id` for datasource `google_service_account` ([#2917](https://github.com/hashicorp/terraform-provider-google-beta/pull/2917))

## 3.55.0 (February 1, 2021)

BREAKING CHANGES:
* Reverted `* bigquery: made incompatible changes to the `google_bigquery_table.schema` field to cause the resource to be recreated ([#8232](https://github.com/hashicorp/terraform-provider-google/pull/8232))` due to unintended interactions with a bug introduced in an earlier version of the resource.

FEATURES:
* **New Data Source:** `google_runtimeconfig_config` ([#8268](https://github.com/hashicorp/terraform-provider-google/pull/8268))

IMPROVEMENTS:
* compute: added `distribution_policy_target_shape` field to `google_compute_region_instance_group_manager` resource ([#8277](https://github.com/hashicorp/terraform-provider-google/pull/8277))
* container: promoted `master_global_access_config`, `tpu_ipv4_cidr_block`, `default_snat_status` and `datapath_provider` fields of `google_container_cluster` to GA. ([#8303](https://github.com/hashicorp/terraform-provider-google/pull/8303))
* dataproc: Added field `temp_bucket` to `google_dataproc_cluster` cluster config. ([#8131](https://github.com/hashicorp/terraform-provider-google/pull/8131))
* notebooks: added `tags`, `service_account_scopes`,`shielded_instance_config` to `google_notebooks_instance` ([#8289](https://github.com/hashicorp/terraform-provider-google/pull/8289))
* provider: added plan time validations for fields that expect base64 values. ([#8304](https://github.com/hashicorp/terraform-provider-google/pull/8304))

BUG FIXES:
* bigquery: fixed permadiff on expiration_ms for `google_bigquery_table` ([#8298](https://github.com/hashicorp/terraform-provider-google/pull/8298))
* billing: fixed perma-diff on currency_code in `google_billing_budget` ([#8266](https://github.com/hashicorp/terraform-provider-google/pull/8266))
 * compute: changed private_ipv6_google_access in `google_compute_subnetwork` to correctly send a fingerprint ([#8290](https://github.com/hashicorp/terraform-provider-google/pull/8290))
* healthcare: add retry logic on healthcare dataset not initialized error ([#8256](https://github.com/hashicorp/terraform-provider-google/pull/8256))

## 3.54.0 (January 25, 2021)

KNOWN ISSUES: New `google_bigquery_table` behaviour introduced in this version had unintended consequences, and may incorrectly flag tables for recreation. We expect to revert this for `3.55.0`.

FEATURES:
* **New Data Source:** `google_cloud_run_locations` ([#2864](https://github.com/hashicorp/terraform-provider-google-beta/pull/2864))
* **New Resource:** `google_privateca_certificate_authority` ([#2877](https://github.com/hashicorp/terraform-provider-google-beta/pull/2877))
* **New Resource:** `google_privateca_certificate_authority_iam_binding` ([#2883](https://github.com/hashicorp/terraform-provider-google-beta/pull/2883))
* **New Resource:** `google_privateca_certificate_authority_iam_member` ([#2883](https://github.com/hashicorp/terraform-provider-google-beta/pull/2883))
* **New Resource:** `google_privateca_certificate_authority_iam_policy` ([#2883](https://github.com/hashicorp/terraform-provider-google-beta/pull/2883))

IMPROVEMENTS:
* bigquery: made incompatible changes to the `google_bigquery_table.schema` field cause the resource to be recreated ([#2876](https://github.com/hashicorp/terraform-provider-google-beta/pull/2876))
* bigtable: fixed an issue where the `google_bigtable_instance` resource was not inferring the zone from the provider. ([#2873](https://github.com/hashicorp/terraform-provider-google-beta/pull/2873))
* cloudscheduler: fixed unnecessary recreate for `google_cloud_scheduler_job` ([#2882](https://github.com/hashicorp/terraform-provider-google-beta/pull/2882))
* compute: added `scaling_schedules` fields to `google_compute_autoscaler` and `google_compute_region_autoscaler` (beta) ([#2879](https://github.com/hashicorp/terraform-provider-google-beta/pull/2879))
* compute: fixed an issue where `google_compute_region_per_instance_config`, `google_compute_per_instance_config`, `google_compute_region_instance_group_manager` resources were not inferring the region/zone from the provider. ([#2874](https://github.com/hashicorp/terraform-provider-google-beta/pull/2874))
* memcache: fixed an issue where `google_memcached_instance` resource was not inferring the region from the provider. ([#2863](https://github.com/hashicorp/terraform-provider-google-beta/pull/2863))
* tpu: fixed an issue where `google_tpu_node` resource was not inferring the zone from the provider. ([#2863](https://github.com/hashicorp/terraform-provider-google-beta/pull/2863))
* vpcaccess: fixed an issue where `google_vpc_access_connector` resource was not inferring the region from the provider. ([#2863](https://github.com/hashicorp/terraform-provider-google-beta/pull/2863))

BUG FIXES:
* bigquery: fixed an issue in `bigquery_dataset_iam_member` where deleted members were not handled correctly ([#2875](https://github.com/hashicorp/terraform-provider-google-beta/pull/2875))
* compute: fixed a perma-diff on `google_compute_health_check` when `log_config.enable` is set to false ([#2866](https://github.com/hashicorp/terraform-provider-google-beta/pull/2866))
* notebooks: fixed permadiff on noRemoveDataDisk for `google_notebooks_instance` ([#2880](https://github.com/hashicorp/terraform-provider-google-beta/pull/2880))
* resourcemanager: fixed an inconsistent result when IAM conditions are specified with `google_folder_iam_*` ([#2878](https://github.com/hashicorp/terraform-provider-google-beta/pull/2878))
* healthcare: added retry logic on healthcare dataset not initialized error ([#2885](https://github.com/hashicorp/terraform-provider-google-beta/pull/2885))

## 3.53.0 (January 19, 2021)

FEATURES:
* **New Data Source:** `google_compute_instance_template` ([#2842](https://github.com/hashicorp/terraform-provider-google-beta/pull/2842))
* **New Resource:** `google_apigee_organization` ([#2856](https://github.com/hashicorp/terraform-provider-google-beta/pull/2856))

IMPROVEMENTS:
* accesscontextmanager: added support for `google_access_context_manager_gcp_user_access_binding` ([#2851](https://github.com/hashicorp/terraform-provider-google-beta/pull/2851))
* memcached: fixed an issue where `google_memcached_instance` resource was not inferring the region from the provider. ([More info](https://github.com/hashicorp/terraform-provider-google/issues/8027))
* serviceaccount: added a `keepers` field to `google_service_account_key` that recreates the field when it is modified ([#2860](https://github.com/hashicorp/terraform-provider-google-beta/pull/2860))
* sql: added restore from backup support to `google_sql_database_instance` ([#2843](https://github.com/hashicorp/terraform-provider-google-beta/pull/2843))
* sql: added support for MYSQL_8_0 on resource `google_sql_source_representation_instance` ([#2841](https://github.com/hashicorp/terraform-provider-google-beta/pull/2841))
* tpu: fixed an issue where `google_tpu_node` resource was not inferring the zone from the provider. ([More info](https://github.com/hashicorp/terraform-provider-google/issues/8027))
* vpcaccess: fixed an issue where `google_vpc_access_connector` resource was not inferring the region from the provider. ([More info](https://github.com/hashicorp/terraform-provider-google/issues/8027))

BUG FIXES:
* bigquery: enhanced diff suppress to ignore certain api divergences on resource `table` ([#2840](https://github.com/hashicorp/terraform-provider-google-beta/pull/2840))
* container: fixed crash due to nil exclusions object when updating an existent cluster with maintenance_policy but without exclusions ([#2839](https://github.com/hashicorp/terraform-provider-google-beta/pull/2839))
* project: fixed a bug in `google_project_access_approval_settings` where the default `project` was used rather than `project_id` ([#2852](https://github.com/hashicorp/terraform-provider-google-beta/pull/2852))

## 3.52.0 (January 11, 2021)

BREAKING CHANGES:
* billing: removed import support for `google_billing_budget` as it never functioned correctly ([#2789](https://github.com/hashicorp/terraform-provider-google-beta/pull/2789))

FEATURES:
* **New Data Source:** `google_sql_backup_run` ([#2824](https://github.com/hashicorp/terraform-provider-google-beta/pull/2824))
* **New Data Source:** `google_storage_bucket_object_content` ([#2785](https://github.com/hashicorp/terraform-provider-google-beta/pull/2785))
* **New Resource:** `google_billing_subaccount` ([#2788](https://github.com/hashicorp/terraform-provider-google-beta/pull/2788))
* **New Resource:** `google_pubsub_lite_subscription` ([#2781](https://github.com/hashicorp/terraform-provider-google-beta/pull/2781))
* **New Resource:** `google_pubsub_lite_topic` ([#2781](https://github.com/hashicorp/terraform-provider-google-beta/pull/2781))

IMPROVEMENTS:
* bigtable: added support for specifying `duration` for `bigtable_gc_policy` to allow durations shorter than a day ([#2815](https://github.com/hashicorp/terraform-provider-google-beta/pull/2815))
* compute: Added support for Google Virtual Network Interface (gVNIC) for `google_compute_image` ([#2779](https://github.com/hashicorp/terraform-provider-google-beta/pull/2779))
* compute: added SHARED_LOADBALANCER_VIP as a valid option for `google_compute_address.purpose` ([#2773](https://github.com/hashicorp/terraform-provider-google-beta/pull/2773))
* compute: added field `multiwriter` to resource `disk` (beta) ([#2822](https://github.com/hashicorp/terraform-provider-google-beta/pull/2822))
* compute: added support for `enable_independent_endpoint_mapping` to `google_compute_router_nat` resource ([#2805](https://github.com/hashicorp/terraform-provider-google-beta/pull/2805))
* compute: added support for `filter.direction` to `google_compute_packet_mirroring` ([#2825](https://github.com/hashicorp/terraform-provider-google-beta/pull/2825))
* compute: promoted `confidential_instance_config` field in `google_compute_instance` and `google_compute_instance_template` to GA ([#2818](https://github.com/hashicorp/terraform-provider-google-beta/pull/2818))
* dataflow: Added optional `kms_key_name` field for `google_dataflow_job` ([#2829](https://github.com/hashicorp/terraform-provider-google-beta/pull/2829))
* dataflow: added documentation about using `parameters` for custom service account and other pipeline options to `google_dataflow_flex_template_job` ([#2776](https://github.com/hashicorp/terraform-provider-google-beta/pull/2776))
* redis: added `auth_string` output to `google_redis_instance` when `auth_enabled` is `true` ([#2819](https://github.com/hashicorp/terraform-provider-google-beta/pull/2819))
* sql: added support for setting the `type` field on `google_sql_user` to support IAM authentication ([#2802](https://github.com/hashicorp/terraform-provider-google-beta/pull/2802))

BUG FIXES:
* bigquery: fixed a bug in `google_bigquery_connection` that caused the resource to function incorrectly when `connection_id` was unset ([#2792](https://github.com/hashicorp/terraform-provider-google-beta/pull/2792))
* compute: removed requirement for `google_compute_region_url_map` default_service, as it should be a choice of default_service or default_url_redirect ([#2810](https://github.com/hashicorp/terraform-provider-google-beta/pull/2810))
* cloud_tasks: fixed permadiff on retry_config.max_retry_duration for `google_cloud_tasks_queue` when the 0s is supplied ([#2812](https://github.com/hashicorp/terraform-provider-google-beta/pull/2812))
* cloudfunctions: fixed a bug where `google_cloudfunctions_function` would sometimes fail to update after being imported from gcloud ([#2780](https://github.com/hashicorp/terraform-provider-google-beta/pull/2780))
* cloudrun: fixed a permanent diff on `google_cloud_run_domain_mapping` `spec.force_override` field ([#2791](https://github.com/hashicorp/terraform-provider-google-beta/pull/2791))
* container: added plan time validation to ensure `enable_private_nodes` is true if `master_ipv4_cidr_block` is set on resource `cluster` ([#2811](https://github.com/hashicorp/terraform-provider-google-beta/pull/2811))
* container: fixed an issue where setting `google_container_cluster.private_cluster_config[0].master_global_access_config.enabled` to `false` caused a permadiff. ([#2816](https://github.com/hashicorp/terraform-provider-google-beta/pull/2816))
* container: fixed setting kubelet_config to disable cpu_cfs_quota does not seem to work ([#2820](https://github.com/hashicorp/terraform-provider-google-beta/pull/2820))
* dataproc: updated jobs to no longer wait for job completion during create ([#2809](https://github.com/hashicorp/terraform-provider-google-beta/pull/2809))
* filestore: updated retry logic to fail fast on quota error which cannot succeed on retry. ([#2814](https://github.com/hashicorp/terraform-provider-google-beta/pull/2814))
* logging: fixed updating on disabled in `google_logging_project_sink` ([#2821](https://github.com/hashicorp/terraform-provider-google-beta/pull/2821))
* scheduler: Fixed syntax error in the Cloud Scheduler HTTP target example. ([#2777](https://github.com/hashicorp/terraform-provider-google-beta/pull/2777))
* sql: fixed a bug in `google_sql_database_instance` that caused a permadiff on `settings.replication_type` ([#2778](https://github.com/hashicorp/terraform-provider-google-beta/pull/2778))
* storage: updated IAM resources to refresh etag sooner on an IAM conflict error, which will make applications of multiple IAM resources much faster. ([#2814](https://github.com/hashicorp/terraform-provider-google-beta/pull/2814))

## 3.51.1 (January 07, 2021)

BUG FIXES:
* all: fixed a bug that would occur in various resources due to comparison of large integers ([#2826](https://github.com/hashicorp/terraform-provider-google-beta/pull/2826))

## 3.51.0 (December 14, 2020)

FEATURES:
* **New Resource:** `google_firestore_document` ([#2759](https://github.com/hashicorp/terraform-provider-google-beta/pull/2759))

IMPROVEMENTS:
* compute: added CDN features to `google_compute_region_backend_service`. ([#2762](https://github.com/hashicorp/terraform-provider-google-beta/pull/2762))
* compute: added Flexible Cache Control features to `google_compute_backend_service`. ([#2762](https://github.com/hashicorp/terraform-provider-google-beta/pull/2762))
* compute: added `replacement_method` field to `update_policy` block of `google_compute_instance_group_manager` ([#2756](https://github.com/hashicorp/terraform-provider-google-beta/pull/2756))
* compute: added `replacement_method` field to `update_policy` block of `google_compute_region_instance_group_manager` ([#2756](https://github.com/hashicorp/terraform-provider-google-beta/pull/2756))
* logging: added plan time validation for `unique_writer_identity` on `google_logging_project_sink` ([#2767](https://github.com/hashicorp/terraform-provider-google-beta/pull/2767))
* storage: added more lifecycle conditions to `google_storage_bucket` resource ([#2761](https://github.com/hashicorp/terraform-provider-google-beta/pull/2761))

BUG FIXES:
* all: bump default request timeout to avoid conflicts if creating a resource takes longer than expected ([#2769](https://github.com/hashicorp/terraform-provider-google-beta/pull/2769))
* project: fixed a bug where `google_project_default_service_accounts` would delete all IAM bindings on a project when run with `action = "DEPRIVILEGE"` ([#2771](https://github.com/hashicorp/terraform-provider-google-beta/pull/2771))
* spanner: fixed an issue in `google_spanner_database` where multi-statement updates were not formatted correctly ([#2766](https://github.com/hashicorp/terraform-provider-google-beta/pull/2766))
* sql: fixed a bug in `google_sql_database_instance` that caused a permadiff on `settings.replication_type` ([#2778](https://github.com/hashicorp/terraform-provider-google-beta/pull/2778))

## 3.50.0 (December 7, 2020)

FEATURES:
* **New Data Source:** `google_composer_environment` ([#2745](https://github.com/hashicorp/terraform-provider-google-beta/pull/2745))
* **New Data Source:** `google_monitoring_cluster_istio_service` ([#2730](https://github.com/hashicorp/terraform-provider-google-beta/pull/2730))
* **New Data Source:** `google_monitoring_mesh_istio_service` ([#2730](https://github.com/hashicorp/terraform-provider-google-beta/pull/2730))

IMPROVEMENTS:
* compute: added `replacement_method` field to `update_policy` block of `google_compute_instance_group_manager` ([#2756](https://github.com/hashicorp/terraform-provider-google-beta/pull/2756))
* compute: added `replacement_method` field to `update_policy` block of `google_compute_region_instance_group_manager` ([#2756](https://github.com/hashicorp/terraform-provider-google-beta/pull/2756))
* compute: added more fields to cdn_policy block of `google_compute_backend_bucket` ([#2741](https://github.com/hashicorp/terraform-provider-google-beta/pull/2741))
* compute: updated `google_compute_url_map`'s fields referring to backend services to be able to refer to backend buckets. ([#2754](https://github.com/hashicorp/terraform-provider-google-beta/pull/2754))
* container: added cluster state check in `resource_container_node_pool` ([#2740](https://github.com/hashicorp/terraform-provider-google-beta/pull/2740))
* google: added support for more import formats to google_project_iam_custom_role ([#2735](https://github.com/hashicorp/terraform-provider-google-beta/pull/2735))
* project: added new restore_policy `REVERT_AND_IGNORE_FAILURE` to `google_project_default_service_accounts` ([#2750](https://github.com/hashicorp/terraform-provider-google-beta/pull/2750))
* serviceusage: Allowed use of field `force` with updates to `google_service_usage_consumer_quota_override` ([#2747](https://github.com/hashicorp/terraform-provider-google-beta/pull/2747))

BUG FIXES:
* bigqueryconnection: fixed failure to import a resource if it has a non-default project or location. ([#2746](https://github.com/hashicorp/terraform-provider-google-beta/pull/2746))
* datacatalog: fixed permadiff on import for tags with a taxonomy set in config. ([#2744](https://github.com/hashicorp/terraform-provider-google-beta/pull/2744))
* iam: fixed iam conflict handling so that optimistic-locking retries will succeed more often. ([#2753](https://github.com/hashicorp/terraform-provider-google-beta/pull/2753))
* storage: fixed an issue in `google_storage_bucket` where `cors` could not be removed ([#2732](https://github.com/hashicorp/terraform-provider-google-beta/pull/2732))

## 3.49.0 (November 24, 2020)

FEATURES:
* **New Resource:** google_healthcare_consent_store ([#2713](https://github.com/hashicorp/terraform-provider-google-beta/pull/2713))
* **New Resource:** google_healthcare_consent_store_iam_binding ([#2713](https://github.com/hashicorp/terraform-provider-google-beta/pull/2713))
* **New Resource:** google_healthcare_consent_store_iam_member ([#2713](https://github.com/hashicorp/terraform-provider-google-beta/pull/2713))
* **New Resource:** google_healthcare_consent_store_iam_policy ([#2713](https://github.com/hashicorp/terraform-provider-google-beta/pull/2713))

IMPROVEMENTS:
* bigquery: added `ORC` as a valid option to `source_format` field of  `google_bigquery_table` resource ([#2714](https://github.com/hashicorp/terraform-provider-google-beta/pull/2714))
* compute: added `custom_response_headers` field to `google_compute_backend_service` resource ([#2722](https://github.com/hashicorp/terraform-provider-google-beta/pull/2722))
* container: added maintenance_exclusions_window to `google_container_cluster` ([#2724](https://github.com/hashicorp/terraform-provider-google-beta/pull/2724))
* logging: added description and disabled to logging sinks ([#2718](https://github.com/hashicorp/terraform-provider-google-beta/pull/2718))
* runtimeconfig: marked value and text fields in `google_runtimeconfig_variable` resource as sensitive ([#2717](https://github.com/hashicorp/terraform-provider-google-beta/pull/2717))
* sql: added `deletion_policy` field to `google_sql_user` to enable abandoning users rather than deleting them ([#2719](https://github.com/hashicorp/terraform-provider-google-beta/pull/2719))

BUG FIXES:
* bigtable: added ignore_warnings flag to create call for `google_bigtable_app_profile` ([#2716](https://github.com/hashicorp/terraform-provider-google-beta/pull/2716))

## 3.48.0 (November 16, 2020)

FEATURES:
* **New Data Source:** `google_iam_workload_identity_pool_provider` ([#2688](https://github.com/hashicorp/terraform-provider-google-beta/pull/2688))

IMPROVEMENTS:
* apigateway: added api_config_id_prefix field to `google_api_gateway_api_config` resoure ([#2692](https://github.com/hashicorp/terraform-provider-google-beta/pull/2692))
* cloudfunctions: fixed a bug with `google_cloudfunction_function` that blocked updates when Organization Policies are enabled. ([#2681](https://github.com/hashicorp/terraform-provider-google-beta/pull/2681))
* compute: added `autoscaling_policy.0.scale_in_control` fields to `google_compute_autoscaler` ([#2703](https://github.com/hashicorp/terraform-provider-google-beta/pull/2703))
* compute: added `autoscaling_policy.0.scale_in_control` fields to `google_compute_region_autoscaler` ([#2703](https://github.com/hashicorp/terraform-provider-google-beta/pull/2703))
* compute: added update support for `google_compute_interconnect_attachment` `bandwidth` field ([#2698](https://github.com/hashicorp/terraform-provider-google-beta/pull/2698))
* dataproc: added "FLINK", "DOCKER", "HBASE" as valid options for field cluster_config.0.software_config.0.optional_components of `google_dataproc_cluster` resource ([#2683](https://github.com/hashicorp/terraform-provider-google-beta/pull/2683))

BUG FIXES:
* cloudrun: added diff suppress function for `google_cloud_run_domain_mapping` `metadata.annotations` to ignore API-set fields ([#2700](https://github.com/hashicorp/terraform-provider-google-beta/pull/2700))
* compute: fixed an issue in `google_compute_packet_mirroring` where updates would fail due to `network` not being updatable ([#2704](https://github.com/hashicorp/terraform-provider-google-beta/pull/2704))
* datacatalog: fixed an issue in `google_data_catalog_taxonomy` and `google_data_catalog_policy_tag` where importing would fail ([#2694](https://github.com/hashicorp/terraform-provider-google-beta/pull/2694))
* spanner: marked `google_spanner_instance.config` as ForceNew as is not updatable ([#2699](https://github.com/hashicorp/terraform-provider-google-beta/pull/2699))

## 3.47.0 (November 09, 2020)

FEATURES:
* **New Data Source:** `google_iam_workload_identity_pool` ([#2663](https://github.com/hashicorp/terraform-provider-google-beta/pull/2663))
* **New Resource:** `google_iam_workload_identity_pool_provider` ([#2670](https://github.com/hashicorp/terraform-provider-google-beta/pull/2670))
* **New Resource:** `google_project_default_service_accounts` ([#2668](https://github.com/hashicorp/terraform-provider-google-beta/pull/2668))

IMPROVEMENTS:
* cloudfunctions: fixed a bug with `google_cloudfunction_function` that blocked updates when Organization Policies are enabled. ([#2681](https://github.com/hashicorp/terraform-provider-google-beta/pull/2681))
* functions: added 4096 as a valid value for available_memory_mb field of `google_cloudfunction_function` ([#2666](https://github.com/hashicorp/terraform-provider-google-beta/pull/2666))
* cloudrun: patched `google_cloud_run_service` to suppress Google generated annotations ([#2679](https://github.com/hashicorp/terraform-provider-google-beta/pull/2679))

BUG FIXES:
* dataflow: removed required validation for zone for `google_data_flow_job` when region is given in the config ([#2662](https://github.com/hashicorp/terraform-provider-google-beta/pull/2662))
* monitoring: Fixed type of `google_monitoring_slo`'s `range` values - some `range` values are doubles, others are integers. ([#2655](https://github.com/hashicorp/terraform-provider-google-beta/pull/2655))
* pubsub: Fixed permadiff on push_config.attributes. ([#2672](https://github.com/hashicorp/terraform-provider-google-beta/pull/2672))
* storage: fixed an issue in `google_storage_bucket` where `lifecycle_rules` were always included in update requests ([#2684](https://github.com/hashicorp/terraform-provider-google-beta/pull/2684))

## 3.46.0 (November 02, 2020)

NOTES:
* compute: updated `google_compute_machine_image` resource to complete once the Image is ready. ([#2637](https://github.com/hashicorp/terraform-provider-google-beta/pull/2637))

FEATURES:
* **New Resource:** `google_api_gateway_api_config_iam_binding` ([#2636](https://github.com/hashicorp/terraform-provider-google-beta/pull/2636))
* **New Resource:** `google_api_gateway_api_config_iam_member` ([#2636](https://github.com/hashicorp/terraform-provider-google-beta/pull/2636))
* **New Resource:** `google_api_gateway_api_config_iam_policy` ([#2636](https://github.com/hashicorp/terraform-provider-google-beta/pull/2636))
* **New Resource:** `google_api_gateway_api_config` ([#2636](https://github.com/hashicorp/terraform-provider-google-beta/pull/2636))
* **New Resource:** `google_api_gateway_api_iam_binding` ([#2636](https://github.com/hashicorp/terraform-provider-google-beta/pull/2636))
* **New Resource:** `google_api_gateway_api_iam_member` ([#2636](https://github.com/hashicorp/terraform-provider-google-beta/pull/2636))
* **New Resource:** `google_api_gateway_api_iam_policy` ([#2636](https://github.com/hashicorp/terraform-provider-google-beta/pull/2636))
* **New Resource:** `google_api_gateway_api` ([#2636](https://github.com/hashicorp/terraform-provider-google-beta/pull/2636))
* **New Resource:** `google_api_gateway_gateway_iam_binding` ([#2636](https://github.com/hashicorp/terraform-provider-google-beta/pull/2636))
* **New Resource:** `google_api_gateway_gateway_iam_member` ([#2636](https://github.com/hashicorp/terraform-provider-google-beta/pull/2636))
* **New Resource:** `google_api_gateway_gateway_iam_policy` ([#2636](https://github.com/hashicorp/terraform-provider-google-beta/pull/2636))
* **New Resource:** `google_api_gateway_gateway` ([#2636](https://github.com/hashicorp/terraform-provider-google-beta/pull/2636))
* **New Resource:** `google_compute_instance_from_machine_image` ([#2637](https://github.com/hashicorp/terraform-provider-google-beta/pull/2637))
* **New Resource:** `google_compute_machine_image_iam_binding` ([#2637](https://github.com/hashicorp/terraform-provider-google-beta/pull/2637))
* **New Resource:** `google_compute_machine_image_iam_member` ([#2637](https://github.com/hashicorp/terraform-provider-google-beta/pull/2637))
* **New Resource:** `google_compute_machine_image_iam_policy` ([#2637](https://github.com/hashicorp/terraform-provider-google-beta/pull/2637))
* **New Resource:** `google_iap_tunnel_iam_binding` ([#2642](https://github.com/hashicorp/terraform-provider-google-beta/pull/2642))
* **New Resource:** `google_iap_tunnel_iam_member` ([#2642](https://github.com/hashicorp/terraform-provider-google-beta/pull/2642))
* **New Resource:** `google_iap_tunnel_iam_policy` ([#2642](https://github.com/hashicorp/terraform-provider-google-beta/pull/2642))

IMPROVEMENTS:
* asset: added conditions to Cloud Asset Feeds ([#2640](https://github.com/hashicorp/terraform-provider-google-beta/pull/2640))
* bigquery: added `email_preferences ` field to `google_bigquery_data_transfer_config` resource ([#2652](https://github.com/hashicorp/terraform-provider-google-beta/pull/2652))
* bigquery: added `schedule_options` field to `google_bigquery_data_transfer_config` resource ([#2641](https://github.com/hashicorp/terraform-provider-google-beta/pull/2641))
* compute: added `private_ipv6_google_access` field to `google_compute_subnetwork` ([#2649](https://github.com/hashicorp/terraform-provider-google-beta/pull/2649))
* compute: added storage_locations & cmek fields to `google_compute_machine_image` resource ([#2637](https://github.com/hashicorp/terraform-provider-google-beta/pull/2637))
* compute: added support for non-destructive updates to `export_custom_routes` and `import_custom_routes` for `google_compute_network_peering` ([#2633](https://github.com/hashicorp/terraform-provider-google-beta/pull/2633))
* compute: relaxed `load_balancing_scheme` validation of `google_compute_region_backend_service` to support external network load-balancers ([#2628](https://github.com/hashicorp/terraform-provider-google-beta/pull/2628))
* container: added `confidential_nodes` field to `google_container_cluster` resource ([#2632](https://github.com/hashicorp/terraform-provider-google-beta/pull/2632))
* datacatalog: added taxonomy and policy_tag `google_data_catalog` ([#2626](https://github.com/hashicorp/terraform-provider-google-beta/pull/2626))
* dlp: added `custom_info_types` to `google_dlp_inspect_template` ([#2648](https://github.com/hashicorp/terraform-provider-google-beta/pull/2648))
* functions: added `build_environment_variables` field to `google_cloudfunction_function` ([#2629](https://github.com/hashicorp/terraform-provider-google-beta/pull/2629))
* kms: added `skip_initial_version_creation` to `google_kms_crypto_key` ([#2645](https://github.com/hashicorp/terraform-provider-google-beta/pull/2645))
* monitoring: added Monitoring Query Language based alerting for `google_monitoring_alert_policy` ([#2651](https://github.com/hashicorp/terraform-provider-google-beta/pull/2651))

BUG FIXES:
* compute: fixed an issue where `google_compute_health_check` `port` values caused a diff when `port_specification` was unset or set to `""` ([#2635](https://github.com/hashicorp/terraform-provider-google-beta/pull/2635))
* monitoring: added more retries for potential failed monitoring operations ([#2639](https://github.com/hashicorp/terraform-provider-google-beta/pull/2639))
* osconfig: fixed an issue where the `rollout.disruption_budget.percentage` field in `google_os_config_patch_deployment` did not correspond to a field in the API ([#2644](https://github.com/hashicorp/terraform-provider-google-beta/pull/2644))
* sql: fixed a case in `google_sql_database_instance` where we inadvertently required the `projects.get` permission for a service networking precheck introduced in `v3.44.0` ([#2634](https://github.com/hashicorp/terraform-provider-google-beta/pull/2634))

## 3.45.0 (October 28, 2020)

BREAKING CHANGES:
* pubsub: changing the value of `google_pubsub_subscription.enable_message_ordering` will now recreate the resource. Previously, an error was returned. ([#2624](https://github.com/hashicorp/terraform-provider-google-beta/pull/2624))
* spanner: `google_spanner_database` resources now cannot be destroyed unless `deletion_protection = false` is set in state for the resource. ([#2612](https://github.com/hashicorp/terraform-provider-google-beta/pull/2612))

NOTES:
* compute: added a warning to `google_compute_vpn_gateway` ([#2607](https://github.com/hashicorp/terraform-provider-google-beta/pull/2607))

FEATURES:
* **New Data Source:** `google_spanner_instance` ([#2602](https://github.com/hashicorp/terraform-provider-google-beta/pull/2602))
* **New Resource:** `google_notebooks_instance_iam_binding` ([#2605](https://github.com/hashicorp/terraform-provider-google-beta/pull/2605))
* **New Resource:** `google_notebooks_instance_iam_member` ([#2605](https://github.com/hashicorp/terraform-provider-google-beta/pull/2605))
* **New Resource:** `google_notebooks_instance_iam_policy` ([#2605](https://github.com/hashicorp/terraform-provider-google-beta/pull/2605))
* **New Resource:** `access_context_manager_access_level_condition` ([#2595](https://github.com/hashicorp/terraform-provider-google-beta/pull/2595))
* **New Resource:** `google_bigquery_routine` ([#2622](https://github.com/hashicorp/terraform-provider-google-beta/pull/2622))
* **New Resource:** `google_iam_workload_identity_pool` ([#2623](https://github.com/hashicorp/terraform-provider-google-beta/pull/2623))
* **New Resource:** `google_data_catalog_taxonomy` ([#2626](https://github.com/hashicorp/terraform-provider-google-beta/pull/2626))
* **New Resource:** `google_data_catalog_policy_tag` ([#2626](https://github.com/hashicorp/terraform-provider-google-beta/pull/2626))
* **New Resource:** `google_data_catalog_taxonomy_iam_binding` ([#2626](https://github.com/hashicorp/terraform-provider-google-beta/pull/2626))
* **New Resource:** `google_data_catalog_taxonomy_iam_member` ([#2626](https://github.com/hashicorp/terraform-provider-google-beta/pull/2626))
* **New Resource:** `google_data_catalog_taxonomy_iam_policy` ([#2626](https://github.com/hashicorp/terraform-provider-google-beta/pull/2626))
* **New Resource:** `google_data_catalog_policy_tag_iam_binding` ([#2626](https://github.com/hashicorp/terraform-provider-google-beta/pull/2626))
* **New Resource:** `google_data_catalog_policy_tag_iam_member` ([#2626](https://github.com/hashicorp/terraform-provider-google-beta/pull/2626))
* **New Resource:** `google_data_catalog_policy_tag_iam_policy` ([#2626](https://github.com/hashicorp/terraform-provider-google-beta/pull/2626))

IMPROVEMENTS:
* billing_budget: added `disable_default_iam_recipients ` field to `google_billing_budget` to allow disable sending email notifications to default recipients. ([#2606](https://github.com/hashicorp/terraform-provider-google-beta/pull/2606))
* compute: added `interface` attribute to `google_compute_disk` ([#2609](https://github.com/hashicorp/terraform-provider-google-beta/pull/2609))
* compute: added `mtu` field to `google_compute_network` resource ([#2617](https://github.com/hashicorp/terraform-provider-google-beta/pull/2617))
* compute: added support for updating `network_interface.[d].network_ip` on `google_compute_instance` when changing network or subnetwork ([#2590](https://github.com/hashicorp/terraform-provider-google-beta/pull/2590))
* compute: promoted HA VPN fields in `google_compute_vpn_tunnel` to GA ([#2607](https://github.com/hashicorp/terraform-provider-google-beta/pull/2607))
* compute: promoted `google_compute_external_vpn_gateway` to GA ([#2607](https://github.com/hashicorp/terraform-provider-google-beta/pull/2607))
* compute: promoted `google_compute_ha_vpn_gateway` to GA ([#2607](https://github.com/hashicorp/terraform-provider-google-beta/pull/2607))
* provider: added support for service account impersonation. ([#2604](https://github.com/hashicorp/terraform-provider-google-beta/pull/2604))
* spanner: added `deletion_protection` field to `google_spanner_database` to make deleting them require an explicit intent. ([#2612](https://github.com/hashicorp/terraform-provider-google-beta/pull/2612))

BUG FIXES:
* all: fixed misleading "empty non-retryable error" message that was appearing in debug logs ([#2618](https://github.com/hashicorp/terraform-provider-google-beta/pull/2618))
* compute: fixed incorrect import format for `google_compute_global_network_endpoint` ([#2594](https://github.com/hashicorp/terraform-provider-google-beta/pull/2594))
* compute: fixed issue where `google_compute_[region_]backend_service.backend.max_utilization` could not be updated ([#2620](https://github.com/hashicorp/terraform-provider-google-beta/pull/2620))
* iap: fixed an eventual consistency bug causing creates for `google_iap_brand` to fail ([#2592](https://github.com/hashicorp/terraform-provider-google-beta/pull/2592))
* provider: fixed an issue where the request headers would grow proportionally to the number of resources in a given `terraform apply` ([#2621](https://github.com/hashicorp/terraform-provider-google-beta/pull/2621))
* serviceusage: fixed bug where concurrent activations/deactivations of project services would fail, now they retry ([#2591](https://github.com/hashicorp/terraform-provider-google-beta/pull/2591))

## 3.44.0 (October 19, 2020)

BREAKING CHANGE:
* Added `deletion_protection` to `google_sql_database_instance`, which defaults to true. SQL instances can no longer be destroyed without setting `deletion_protection = false`. ([#2579](https://github.com/hashicorp/terraform-provider-google-beta/pull/2579))

FEATURES:
* **New Data Source:** `google_app_engine_default_service_account` ([#2568](https://github.com/hashicorp/terraform-provider-google-beta/pull/2568))
* **New Data Source:** `google_pubsub_topic` ([#2556](https://github.com/hashicorp/terraform-provider-google-beta/pull/2556))

IMPROVEMENTS:
* bigquery: added ability for `google_bigquery_dataset_access` to retry quota errors since quota refreshes quickly. ([#2584](https://github.com/hashicorp/terraform-provider-google-beta/pull/2584))
* bigquery: added `MONTH` and `YEAR` as allowed values in `google_bigquery_table.time_partitioning.type` ([#2562](https://github.com/hashicorp/terraform-provider-google-beta/pull/2562))
* cloud_tasks: added `stackdriver_logging_config` field to `cloud_tasks_queue` resource ([#2572](https://github.com/hashicorp/terraform-provider-google-beta/pull/2572))
* compute: added support for updating `network_interface.[d].network_ip` on `google_compute_instance` when changing network or subnetwork ([#2590](https://github.com/hashicorp/terraform-provider-google-beta/pull/2590))
* compute: added `maintenance_policy` field to `google_compute_node_group` ([#2586](https://github.com/hashicorp/terraform-provider-google-beta/pull/2586))
* compute: added filter field to google_compute_image datasource ([#2573](https://github.com/hashicorp/terraform-provider-google-beta/pull/2573))
* dataproc: Added `graceful_decomissioning_timeout` field to `dataproc_cluster` resource ([#2571](https://github.com/hashicorp/terraform-provider-google-beta/pull/2571))
* iam: fixed `google_service_account_id_token` datasource to work with User ADCs and Impersonated Credentials ([#2560](https://github.com/hashicorp/terraform-provider-google-beta/pull/2560))
* logging: Added support for exclusions options for `google_logging_project_sink ` ([#2569](https://github.com/hashicorp/terraform-provider-google-beta/pull/2569))
* logging: added bucket creation based on custom-id given for the resource `google_logging_project_bucket_config` ([#2575](https://github.com/hashicorp/terraform-provider-google-beta/pull/2575))
* oslogin: added ability to set a `project` on `google_os_login_ssh_public_key` ([#2583](https://github.com/hashicorp/terraform-provider-google-beta/pull/2583))
* redis: Added `auth_enabled` field to `google_redis_instance` ([#2570](https://github.com/hashicorp/terraform-provider-google-beta/pull/2570))
* resourcemanager: added a precheck that the serviceusage API is enabled to `google_project` when `auto_create_network` is false, as configuring the GCE API is required in that circumstance ([#2566](https://github.com/hashicorp/terraform-provider-google-beta/pull/2566))
* sql: added a check to `google_sql_database_instance` to catch failures early by seeing if Service Networking Connections already exists for the private network of the instance. ([#2579](https://github.com/hashicorp/terraform-provider-google-beta/pull/2579))

BUG FIXES:
* accessapproval: fixed issue where, due to a recent API change, `google_*_access_approval.enrolled_services.cloud_product` entries specified as a URL would result in a permadiff ([#2565](https://github.com/hashicorp/terraform-provider-google-beta/pull/2565))
* compute: fixed ability to clear `description` field on `google_compute_health_check` and `google_compute_region_health_check` ([#2580](https://github.com/hashicorp/terraform-provider-google-beta/pull/2580))
* monitoring: fixed bug where deleting a `google_monitoring_dashboard` would give an "unsupported protocol scheme" error ([#2558](https://github.com/hashicorp/terraform-provider-google-beta/pull/2558))

## 3.43.0 (October 12, 2020)

FEATURES:
* **New Data Source:** `google_pubsub_topic` ([#2556](https://github.com/hashicorp/terraform-provider-google-beta/pull/2556))
* **New Data Source:** `google_compute_global_forwarding_rule` ([#2548](https://github.com/hashicorp/terraform-provider-google-beta/pull/2548))
* **New Data Source:** `google_cloud_run_service` ([#2539](https://github.com/hashicorp/terraform-provider-google-beta/pull/2539))
* **New Resource:** `google_bigtable_table_iam_member` ([#2536](https://github.com/hashicorp/terraform-provider-google-beta/pull/2536))
* **New Resource:** `google_bigtable_table_iam_binding` ([#2536](https://github.com/hashicorp/terraform-provider-google-beta/pull/2536))
* **New Resource:** `google_bigtable_table_iam_policy` ([#2536](https://github.com/hashicorp/terraform-provider-google-beta/pull/2536))

IMPROVEMENTS:
* appengine: added ability to manage pre-firestore appengine applications. ([#2533](https://github.com/hashicorp/terraform-provider-google-beta/pull/2533))
* bigquery: added support for `google_bigquery_table` `materialized_view` field ([#2532](https://github.com/hashicorp/terraform-provider-google-beta/pull/2532))
* cloudbuild: Added `COMMENTS_ENABLED_FOR_EXTERNAL_CONTRIBUTORS_ONLY` support to `google_cloudbuild_trigger.github.pull_request.comment_control` field ([#2552](https://github.com/hashicorp/terraform-provider-google-beta/pull/2552))
* compute: added additional fields to the `google_compute_forwarding_rule` datasource. ([#2550](https://github.com/hashicorp/terraform-provider-google-beta/pull/2550))
* dns: added `forwarding_path` field to `google_dns_policy` resource ([#2540](https://github.com/hashicorp/terraform-provider-google-beta/pull/2540))
* netblock: changed `google_netblock_ip_ranges` to read from cloud.json file rather than DNS record ([#2543](https://github.com/hashicorp/terraform-provider-google-beta/pull/2543))

BUG FIXES:
* accessapproval: fixed issue where, due to a recent API change, `google_*_access_approval.enrolled_services.cloud_product` entries specified as a URL would result in a permadiff
* artifactregistry: fixed an issue where `google_artifact_registry_repository` would import an empty state ([#2546](https://github.com/hashicorp/terraform-provider-google-beta/pull/2546))
* bigquery: fixed an issue in `google_bigquery_job` where non-US locations could not be read ([#2542](https://github.com/hashicorp/terraform-provider-google-beta/pull/2542))
* cloudrun: fixed an issue in `google_cloud_run_domain_mapping` where labels provided by Google would cause a diff ([#2531](https://github.com/hashicorp/terraform-provider-google-beta/pull/2531))
* compute: Fixed an issue where `google_compute_region_backend_service` required `healthChecks` for a serverless network endpoint group. ([#2547](https://github.com/hashicorp/terraform-provider-google-beta/pull/2547))
* container: fixed `node_config.image_type` perma-diff when specified in lower case. ([#2538](https://github.com/hashicorp/terraform-provider-google-beta/pull/2538))
* datacatalog: fixed an error in `google_data_catalog_tag` when trying to set boolean field to `false` ([#2534](https://github.com/hashicorp/terraform-provider-google-beta/pull/2534))
* monitoring: fixed bug where deleting a `google_monitoring_dashboard` would give an "unsupported protocol scheme" error

## 3.42.0 (October 05, 2020)

FEATURES:
* **New Resource:** google_data_loss_prevention_deidentify_template ([#2524](https://github.com/hashicorp/terraform-provider-google-beta/pull/2524))

IMPROVEMENTS:
* compute: added support for updating `network_interface.[d].network` and `network_interface.[d].subnetwork` properties on `google_compute_instance`. ([#2517](https://github.com/hashicorp/terraform-provider-google-beta/pull/2517))
* container: added `notification_config` to `google_container_cluster` ([#2521](https://github.com/hashicorp/terraform-provider-google-beta/pull/2521))
* dataflow: added `region` field to `google_dataflow_flex_template_job` resource ([#2520](https://github.com/hashicorp/terraform-provider-google-beta/pull/2520))
* healthcare: added field `parser_config.version` to `google_healthcare_hl7_v2_store` ([#2516](https://github.com/hashicorp/terraform-provider-google-beta/pull/2516))

BUG FIXES:
* bigquery: fixed an issue where `google_bigquery_table` would crash while reading an empty schema ([#2518](https://github.com/hashicorp/terraform-provider-google-beta/pull/2518))
* compute: fixed an issue where `google_compute_instance_template` would throw an error for unspecified `disk_size_gb` values while upgrading the provider. ([#2515](https://github.com/hashicorp/terraform-provider-google-beta/pull/2515))
* resourcemanager: fixed an issue in retrieving `google_active_folder` data source when the display name included whitespace ([#2528](https://github.com/hashicorp/terraform-provider-google-beta/pull/2528))

## 3.41.0 (September 28, 2020)

IMPROVEMENTS:
* container: Added support for `datapath_provider` to `google_container_cluster` ([#2492](https://github.com/hashicorp/terraform-provider-google-beta/pull/2492))
* cloudfunctions: added the ALLOW_INTERNAL_AND_GCLB option to `ingress_settings` of `google_cloudfunctions_function` resource. ([#2493](https://github.com/hashicorp/terraform-provider-google-beta/pull/2493))
* composer: allowed in-place updates to webserver and database machine type ([#2491](https://github.com/hashicorp/terraform-provider-google-beta/pull/2491))
* compute: added `SEV_CAPABLE` option to `guestOsFeatures` in `google_compute_image` resource. ([#2503](https://github.com/hashicorp/terraform-provider-google-beta/pull/2503))
* tpu: added `use_service_networking` to `google_tpu_node` which enables Shared VPC Support. ([#2497](https://github.com/hashicorp/terraform-provider-google-beta/pull/2497))

BUG FIXES:
* cloudidentity: Fixed upstream breakage of `google_identity_group`. ([#2507](https://github.com/hashicorp/terraform-provider-google-beta/pull/2507))

## 3.40.0 (September 22, 2020)

DEPRECATIONS:
* bigtable: deprecated `instance_type` for `google_bigtable_instance` - it is now recommended to leave field unspecified. ([#2477](https://github.com/hashicorp/terraform-provider-google-beta/pull/2477))

FEATURES:
* **New Data Source:** `google_compute_region_ssl_certificate` ([#2476](https://github.com/hashicorp/terraform-provider-google-beta/pull/2476))
* **New Resource:** `google_compute_target_grpc_proxy` ([#2488](https://github.com/hashicorp/terraform-provider-google-beta/pull/2488))

IMPROVEMENTS:
* cloudlbuild: added `options` and `artifacts` properties to `google_cloudbuild_trigger` ([#2490](https://github.com/hashicorp/terraform-provider-google-beta/pull/2490))
* compute: added GRPC as a valid value for `google_compute_backend_service.protocol` (and regional equivalent) ([#2478](https://github.com/hashicorp/terraform-provider-google-beta/pull/2478))
* compute: added 'all' option for `google_compute_firewall` ([#2465](https://github.com/hashicorp/terraform-provider-google-beta/pull/2465))
* container: added support for `load_balancer_type` to `google_container_cluster` Cloud Run config addon. ([#2487](https://github.com/hashicorp/terraform-provider-google-beta/pull/2487))
* dataflow: added `transformnameMapping` to `google_dataflow_job` ([#2480](https://github.com/hashicorp/terraform-provider-google-beta/pull/2480))
* serviceusage: added ability to pass google.project.id to `google_project_service.project` ([#2479](https://github.com/hashicorp/terraform-provider-google-beta/pull/2479))
* spanner: added schema update/update ddl support for `google_spanner_database` ([#2489](https://github.com/hashicorp/terraform-provider-google-beta/pull/2489))

BUG FIXES:
* bigtable: fixed the update behaviour of the `single_cluster_routing` sub-fields in `google_bigtable_app_profile` ([#2482](https://github.com/hashicorp/terraform-provider-google-beta/pull/2482))
* dataproc: fixed issues where updating `google_dataproc_cluster.cluster_config.autoscaling_policy` would do nothing, and where there was no way to remove a policy. ([#2483](https://github.com/hashicorp/terraform-provider-google-beta/pull/2483))
* osconfig: fixed a potential crash in `google_os_config_patch_deployment` due to an unchecked nil value in `recurring_schedule` ([#2481](https://github.com/hashicorp/terraform-provider-google-beta/pull/2481))
* serviceusage: fixed intermittent failure when a service is already being modified - added retries ([#2469](https://github.com/hashicorp/terraform-provider-google-beta/pull/2469))
* serviceusage: fixed an issue where `bigquery.googleapis.com` was getting enabled as the `bigquery-json.googleapis.com` alias instead, incorrectly. This had no user impact yet, but the alias may go away in the future. ([#2469](https://github.com/hashicorp/terraform-provider-google-beta/pull/2469))

## 3.39.0 (September 15, 2020)

IMPROVEMENTS:
* compute: added network field to `compute_target_instance` ([#2456](https://github.com/hashicorp/terraform-provider-google-beta/pull/2456))
* compute: added storage_locations field to `google_compute_snapshot` ([#2461](https://github.com/hashicorp/terraform-provider-google-beta/pull/2461))
* compute: added `kms_key_service_account`, `kms_key_self_link ` fields to `snapshot_encryption_key` field in `google_compute_snapshot` ([#2461](https://github.com/hashicorp/terraform-provider-google-beta/pull/2461))
* compute: added `source_disk_encryption_key.kms_key_service_account` field to `google_compute_snapshot` ([#2461](https://github.com/hashicorp/terraform-provider-google-beta/pull/2461))
* container: Added `self_link` to google_container_cluster ([#2457](https://github.com/hashicorp/terraform-provider-google-beta/pull/2457))

BUG FIXES:
* bigquery: fixed a bug when a BigQuery table schema didn't have `name` in the schema. Previously it would panic; now it logs an error. ([#2462](https://github.com/hashicorp/terraform-provider-google-beta/pull/2462))
* bigquery: fixed bug where updating `clustering` would force a new resource rather than update. ([#2459](https://github.com/hashicorp/terraform-provider-google-beta/pull/2459))
* bigquerydatatransfer: fixed `params.secret_access_key` perma-diff for AWS S3 data transfer config types by adding a `sensitive_params` block with the `secret_access_key` attribute. ([#2451](https://github.com/hashicorp/terraform-provider-google-beta/pull/2451))
* compute: fixed bug where `delete_default_routes_on_create=true` was not actually deleting the default routes on create. ([#2460](https://github.com/hashicorp/terraform-provider-google-beta/pull/2460)

## 3.38.0 (September 08, 2020)

DEPRECATIONS:
* storage: deprecated `bucket_policy_only` field in `google_storage_bucket` in favour of `uniform_bucket_level_access` ([#2442](https://github.com/hashicorp/terraform-provider-google-beta/pull/2442))

FEATURES:
* **New Resource:** google_compute_disk_iam_binding ([#2424](https://github.com/hashicorp/terraform-provider-google-beta/pull/2424))
* **New Resource:** google_compute_disk_iam_member ([#2424](https://github.com/hashicorp/terraform-provider-google-beta/pull/2424))
* **New Resource:** google_compute_disk_iam_policy ([#2424](https://github.com/hashicorp/terraform-provider-google-beta/pull/2424))
* **New Resource:** google_compute_region_disk_iam_binding ([#2424](https://github.com/hashicorp/terraform-provider-google-beta/pull/2424))
* **New Resource:** google_compute_region_disk_iam_member ([#2424](https://github.com/hashicorp/terraform-provider-google-beta/pull/2424))
* **New Resource:** google_compute_region_disk_iam_policy ([#2424](https://github.com/hashicorp/terraform-provider-google-beta/pull/2424))
* **New Resource:** google_data_loss_prevention_inspect_template ([#2433](https://github.com/hashicorp/terraform-provider-google-beta/pull/2433))
* **New Resource:** google_data_loss_prevention_job_trigger ([#2433](https://github.com/hashicorp/terraform-provider-google-beta/pull/2433))
* **New Resource:** google_data_loss_prevention_stored_info_type ([#2444](https://github.com/hashicorp/terraform-provider-google-beta/pull/2444))
* **New Resource:** google_project_service_identity ([#2430](https://github.com/hashicorp/terraform-provider-google-beta/pull/2430))

IMPROVEMENTS:
* compute: Added graceful termination to `google_compute_instance_group_manager` create calls so that partially created instance group managers will resume the original operation if the Terraform process is killed mid create. ([#2446](https://github.com/hashicorp/terraform-provider-google-beta/pull/2446))
* container: added project override support to `google_container_cluster` and `google_container_nodepool` ([#2428](https://github.com/hashicorp/terraform-provider-google-beta/pull/2428))
* notebooks: added `PD_BALANCED` as a possible disk type for `google_notebooks_instance` ([#2438](https://github.com/hashicorp/terraform-provider-google-beta/pull/2438))
* osconfig: added rollout field to `google_os_config_patch_deployment` ([#2449](https://github.com/hashicorp/terraform-provider-google-beta/pull/2449))
* provider: added a new field `billing_project` to the provider that's associated as a billing/quota project with most requests when `user_project_override` is true ([#2427](https://github.com/hashicorp/terraform-provider-google-beta/pull/2427))
* resourcemanager: added additional fields to `google_projects` datasource ([#2440](https://github.com/hashicorp/terraform-provider-google-beta/pull/2440))
* serviceusage: added project override support to `google_project_service` ([#2428](https://github.com/hashicorp/terraform-provider-google-beta/pull/2428))

BUG FIXES:
* bigquerydatatransfer: fixed `params.secret_access_key` perma-diff for AWS S3 data transfer config types by adding a `sensitive_params` block with the `secret_access_key` attribute. ([#2451](https://github.com/hashicorp/terraform-provider-google-beta/pull/2451))
* compute: Fixed bug with `google_netblock_ip_ranges` data source failing to read from the correct URL ([#2448](https://github.com/hashicorp/terraform-provider-google-beta/pull/2448))
* compute: fixed updating `google_compute_instance.shielded_instance_config` by adding it to the `allow_stopping_for_update` list ([#2436](https://github.com/hashicorp/terraform-provider-google-beta/pull/2436))
* notebooks: fixed broken `google_notebooks_instance.instance_owners` field by making it a list instead of a string ([#2438](https://github.com/hashicorp/terraform-provider-google-beta/pull/2438))

## 3.37.0 (August 31, 2020)
NOTES:
* Drop recommendation to use -provider= on import in documentation ([#2417](https://github.com/hashicorp/terraform-provider-google-beta/pull/2417))

FEATURES:
* **New Resource:** `google_compute_image_iam_binding` ([#2410](https://github.com/hashicorp/terraform-provider-google-beta/pull/2410))
* **New Resource:** `google_compute_image_iam_member` ([#2410](https://github.com/hashicorp/terraform-provider-google-beta/pull/2410))
* **New Resource:** `google_compute_image_iam_policy` ([#2410](https://github.com/hashicorp/terraform-provider-google-beta/pull/2410))
* **New Resource:** `google_compute_disk_iam_binding` ([#2424](https://github.com/hashicorp/terraform-provider-google-beta/pull/2424))
* **New Resource:** `google_compute_disk_iam_member` ([#2424](https://github.com/hashicorp/terraform-provider-google-beta/pull/2424))
* **New Resource:** `google_compute_disk_iam_policy` ([#2424](https://github.com/hashicorp/terraform-provider-google-beta/pull/2424))
* **New Resource:** `google_compute_region_disk_iam_binding` ([#2424](https://github.com/hashicorp/terraform-provider-google-beta/pull/2424))
* **New Resource:** `google_compute_region_disk_iam_member` ([#2424](https://github.com/hashicorp/terraform-provider-google-beta/pull/2424))
* **New Resource:** `google_compute_region_disk_iam_policy` ([#2424](https://github.com/hashicorp/terraform-provider-google-beta/pull/2424))

IMPROVEMENTS:
* appengine: added `vpc_access_connector` field to `google_app_engine_standard_app_version` resource ([#2405](https://github.com/hashicorp/terraform-provider-google-beta/pull/2405))
* bigquery: added `notification_pubsub_topic` field to `google_bigquery_data_transfer_config` resource ([#2411](https://github.com/hashicorp/terraform-provider-google-beta/pull/2411))
* composer: added `database_config` and `web_server_config` to `google_composer_environment` resource ([#2419](https://github.com/hashicorp/terraform-provider-google-beta/pull/2419))
* compute: Added custom metadata fields and filter expressions to `google_compute_subnetwork` flow log configuration ([#2416](https://github.com/hashicorp/terraform-provider-google-beta/pull/2416))
* compute: Added support to `google_compute_backend_service` for setting a serverless regional network endpoint group as `backend.group` ([#2408](https://github.com/hashicorp/terraform-provider-google-beta/pull/2408))
* compute: added support for pd-balanced disk type for `google_compute_instance` ([#2421](https://github.com/hashicorp/terraform-provider-google-beta/pull/2421))
* container: added support for `kubelet_config` and `linux_node_config` to GKE node pools ([#2279](https://github.com/hashicorp/terraform-provider-google-beta/pull/2279), [#2403](https://github.com/hashicorp/terraform-provider-google-beta/pull/2403))
* container: added support for pd-balanced disk type for `google_container_node_pool` ([#2421](https://github.com/hashicorp/terraform-provider-google-beta/pull/2421))
* memcached: added discovery_endpoint to `resource_memcached_instance` ([#2414](https://github.com/hashicorp/terraform-provider-google-beta/pull/2414))
* pubsub: added `retry_policy` to `google_pubsub_subscription` resource ([#2412](https://github.com/hashicorp/terraform-provider-google-beta/pull/2412))

BUG FIXES:
* compute: fixed an issue where `google_compute_url_map` `path_matcher.default_route_action` would conflict with `default_url_redirect` ([#2406](https://github.com/hashicorp/terraform-provider-google-beta/pull/2406))
* kms: updated `data_source_secret_manager_secret_version` to have consistent id value ([#2415](https://github.com/hashicorp/terraform-provider-google-beta/pull/2415))

## 3.36.0 (August 24, 2020)

FEATURES:
* **New Resource:** `google_active_directory_domain_trust` ([#2401](https://github.com/hashicorp/terraform-provider-google-beta/pull/2401))
* **New Resource:** `google_access_context_manager_service_perimeters` ([#2382](https://github.com/hashicorp/terraform-provider-google-beta/pull/2382))
* **New Resource:** `google_access_context_manager_access_levels` ([#2382](https://github.com/hashicorp/terraform-provider-google-beta/pull/2382))
* **New Resource:** `google_folder_access_approval_settings` ([#2373](https://github.com/hashicorp/terraform-provider-google-beta/pull/2373))
* **New Resource:** `google_organization_access_approval_settings` ([#2373](https://github.com/hashicorp/terraform-provider-google-beta/pull/2373))
* **New Resource:** `google_project_access_approval_settings` ([#2373](https://github.com/hashicorp/terraform-provider-google-beta/pull/2373))
* **New Resource:** `google_bigquery_table_iam_policy` ([#2392](https://github.com/hashicorp/terraform-provider-google-beta/pull/2392))
* **New Resource:** `google_bigquery_table_iam_binding` ([#2392](https://github.com/hashicorp/terraform-provider-google-beta/pull/2392))
* **New Resource:** `google_bigquery_table_iam_member` ([#2392](https://github.com/hashicorp/terraform-provider-google-beta/pull/2392))

IMPROVEMENTS:
* billing: added `last_period_amount` field to `google_billing_budget` to allow setting budget amount automatically to the last billing period's spend. ([#2378](https://github.com/hashicorp/terraform-provider-google-beta/pull/2378))
* compute: added confidential_instance_config block to google_compute_instance ([#2369](https://github.com/hashicorp/terraform-provider-google-beta/pull/2369))
* compute: added confidential_instance_config block to google_compute_instance_template ([#2369](https://github.com/hashicorp/terraform-provider-google-beta/pull/2369))
* compute: added grpc_health_check block to compute_health_check ([#2389](https://github.com/hashicorp/terraform-provider-google-beta/pull/2389))
* compute: added grpc_health_check block to compute_region_health_check ([#2389](https://github.com/hashicorp/terraform-provider-google-beta/pull/2389))
* pubsub: added `enable_message_ordering` support to `google_pubsub_subscription` ([#2390](https://github.com/hashicorp/terraform-provider-google-beta/pull/2390))
* sql: added project field to `google_sql_database_instance` datasource. ([#2370](https://github.com/hashicorp/terraform-provider-google-beta/pull/2370))
* storage: added `ARCHIVE` as an accepted class for `google_storage_bucket` and `google_storage_bucket_object` ([#2385](https://github.com/hashicorp/terraform-provider-google-beta/pull/2385))

BUG FIXES:
* all: updated base urls for compute, dns, storage, and bigquery APIs to their recommended endpoints ([#2396](https://github.com/hashicorp/terraform-provider-google-beta/pull/2396))
* bigquery: fixed a bug where `dataset_access.iam_member` would produce inconsistent results after apply. ([#2397](https://github.com/hashicorp/terraform-provider-google-beta/pull/2397))
* bigquery: fixed an issue with `use_legacy_sql` not being set to `false`. ([#2375](https://github.com/hashicorp/terraform-provider-google-beta/pull/2375))
* cloudidentity: fixed a bug with importing `google_cloud_identity_group` and `google_cloud_identity_group_membership` ([#2379](https://github.com/hashicorp/terraform-provider-google-beta/pull/2379))
* cloudidentity: fixed cloud identity datasources to handle pagination ([#2387](https://github.com/hashicorp/terraform-provider-google-beta/pull/2387))
* compute: set the default value for log_config.enable on `google_compute_health_check` to avoid permanent diff on plan/apply. ([#2399](https://github.com/hashicorp/terraform-provider-google-beta/pull/2399))
* dns: fixed an issue where `google_dns_managed_zone` would not remove `private_visibility_config` on updates ([#2380](https://github.com/hashicorp/terraform-provider-google-beta/pull/2380))
* sql: fixed an issue where `google_sql_database_instance` would throw an error when removing `private_network`. Removing `private_network` now recreates the resource. ([#2400](https://github.com/hashicorp/terraform-provider-google-beta/pull/2400))

## 3.35.0 (August 17, 2020)
NOTES:
* all: Updated lists of enums to display the enum options in the documentation pages. ([#2340](https://github.com/hashicorp/terraform-provider-google-beta/pull/2340))

FEATURES:
* **New Resource:** `google_compute_region_network_endpoint_group` (supports serverless NEGs) ([#2348](https://github.com/hashicorp/terraform-provider-google-beta/pull/2348))

IMPROVEMENTS:
* appengine: converted `google_app_engine_standard_app_version`'s `inbound_services` to an enum array, which enhances docs and provides some client-side validation. ([#2344](https://github.com/hashicorp/terraform-provider-google-beta/pull/2344))
* billing_budget: Added support for `monitoring_notification_channels` to allow sending budget notifications to Cloud Monitoring email notification channels. ([#2366](https://github.com/hashicorp/terraform-provider-google-beta/pull/2366))
* cloudbuild: added tags, source, queue_ttl, logs_bucket, substitutions, and secrets to `google_cloudbuild_trigger` ([#2335](https://github.com/hashicorp/terraform-provider-google-beta/pull/2335))
* cloudfunctions: Updated the `google_cloudfunctions_function` datasource to include new fields available in the API. ([#2334](https://github.com/hashicorp/terraform-provider-google-beta/pull/2334))
* compute: added `source_image` and `source_snapshot` to `google_compute_image` ([#2356](https://github.com/hashicorp/terraform-provider-google-beta/pull/2356))
* compute: added confidential_instance_config block to google_compute_instance ([#2369](https://github.com/hashicorp/terraform-provider-google-beta/pull/2369))
* compute: added confidential_instance_config block to google_compute_instance_template ([#2369](https://github.com/hashicorp/terraform-provider-google-beta/pull/2369))
* iam: Added `public_key_type` field to `google_service_account_key ` ([#2368](https://github.com/hashicorp/terraform-provider-google-beta/pull/2368))
* memcached: added memcacheVersion input and memcacheNodes output field to `google_memcache_instance` ([#2336](https://github.com/hashicorp/terraform-provider-google-beta/pull/2336))
* pubsub: added `filter` field to `google_pubsub_subscription` resource ([#2367](https://github.com/hashicorp/terraform-provider-google-beta/pull/2367))
* resource-manager: updated documentation for `folder_iam_*` and `organization_iam_*` resources. ([#2365](https://github.com/hashicorp/terraform-provider-google-beta/pull/2365))
* sql: added support for point_in_time_recovery for `google_sql_database_instance` ([#2338](https://github.com/hashicorp/terraform-provider-google-beta/pull/2338))

BUG FIXES:
* appengine: Set `iap` to computed in `google_app_engine_application` ([#2342](https://github.com/hashicorp/terraform-provider-google-beta/pull/2342))
* artifactrepository: Fixed import failure of `google_artifact_registry_repository`. ([#2345](https://github.com/hashicorp/terraform-provider-google-beta/pull/2345))
* compute: fixed shielded instance config, which had been failing to apply due to a field rename on the GCP side. ([#2337](https://github.com/hashicorp/terraform-provider-google-beta/pull/2337))
* monitoring: fixed validation rules for `google_monitoring_slo` `windows_based_sli.metric_sum_in_range.max` field ([#2354](https://github.com/hashicorp/terraform-provider-google-beta/pull/2354))
* osconfig: fixed `google_os_config_patch_deployment` `windows_update.classifications` field to work correctly, accepting multiple values. ([#2340](https://github.com/hashicorp/terraform-provider-google-beta/pull/2340))

## 3.34.0 (August 11, 2020)
NOTES:
* redis: explicitly noted in `google_redis_instance` documentation that `"REDIS_5_0"` is supported ([#2323](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2323))
* all: fix markdown formatting while showing enum values in documentation ([#2327](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2327))

FEATURES:
* **New Resource:** `google_compute_compute_organization_security_policy_association` ([#2333](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2333))
* **New Resource:** `google_compute_compute_organization_security_policy_rule` ([#2333](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2333))
* **New Resource:** `google_compute_compute_organization_security_policy` ([#2333](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2333))

IMPROVEMENTS:
* bigtable: added support for labels in `google_bigtable_instance` ([#2325](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2325))
* cloudfunctions: updated the `google_cloudfunctions_function` datasource to include new fields available in the API. ([#2334](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2334))
* compute: masked automatically applied GKE Sandbox node labels and taints on node pools ([#2320](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2320))
* redis: added `persistence_iam_identity` output field to `google_redis_instance` ([#2323](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2323))
* storage: added output-only `media_link` to `google_storage_bucket_object` ([#2331](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2331))

BUG FIXES:
* compute: fixed issue where the `project` field in `data.google_compute_network_endpoint_group` was returning an error when specified ([#2324](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2324))
* notebook: fixed bug where not setting `data_disk_type` or `disk_encryption` would cause a diff on the next plan ([#2332](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2332))
* sourcerepo: fixed perma-diff in `google_sourcerepo_repository` ([#2316](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2316))
* all: fixed crash due to nil context when loading credentials ([#2321](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2321))

## 3.33.0 (August 04, 2020)

DEPRECATIONS:
* compute: deprecated `enable_logging` on `google_compute_firewall`, define `log_config.metadata` to enable logging instead. ([#2310](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2310))

FEATURES:
* **New Resource:** `google_active_directory_domain` ([#2309](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2309))
* **New Resource:** `google_dataflow_flex_template_job` ([#2303](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2303))

IMPROVEMENTS:
* cloudrun: added `ports` field to `google_cloud_run_service` `templates.spec.containers` ([#2311](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2311))
* compute: added `log_config.metadata` to `google_compute_firewall`, defining this will enable logging. ([#2310](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2310))

BUG FIXES:
* container: Fixed a crash in `google_container_cluster` when `""` was specified for `resource_usage_export_config.bigquery_destination.dataset_id`. ([#2296](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2296))
* endpoints: Fixed a crash when `google_endpoints_service` is used on a machine without timezone data ([#2302](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2302))
* resourcemanager: bumped `google_project` timeout defaults to 10 minutes (from 4) ([#2306](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2306)

## 3.32.0 (July 27, 2020)
FEATURES:
* **New Data Source:** `google_sql_database_instance`  #2841 ([#2273](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2273))
* **New Resource:** `google_cloud_asset_folder_feed` ([#2284](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2284))
* **New Resource:** `google_cloud_asset_organization_feed` ([#2284](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2284))
* **New Resource:** `google_cloud_asset_project_feed` ([#2284](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2284))
* **New Resource:** `google_monitoring_metric_descriptor` ([#2290](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2290))
* **New Resource:** `google_os_config_guest_policies` ([#2276](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2276))

IMPROVEMENTS:
* cluster: Added `default_snat_status` field for `google_container_cluster` resource. ([#2283](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2283))
* filestore: Added `nfs_export_options` field on `google_filestore_instance.file_shares`. ([#2289](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2289))
* filestore: Added support for filestore high scale tier. ([#2289](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2289))
* resourcemanager: Added `folder_id` as computed attribute to `google_folder` resource and datasource. ([#2287](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2287))
* compute: Added support to `google_compute_backend_service` for setting a network endpoint group as `backend.group`. ([#2304](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2304))

BUG FIXES:
* container: Fixed `google_container_cluster.pod_security_policy_config` not being set when disabled.
* container: Fixed a crash in `google_container_cluster` when `""` was specified for `resource_usage_export_config.bigquery_destination.dataset_id`. ([#2296](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2296))
* bigquery: Fixed bug where a permadiff would show up when adding a column to the middle of a `bigquery_table.schema` ([#2275](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2275))
* notebook: Fixed bug where many fields were being written as empty to state, causing a diff on the next plan ([#2288](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2288))
* notebook: Fixed bug where setting `network` or `subnet` to a full URL would succeed, but cause a diff on the next plan ([#2288](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2288))
* notebook: Fixed bug where updating certain fields would result in a no-op update call instead of a create/destroy. Now, the only field that is updatable in place is `labels` ([#2288](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2288))

## 3.31.0 (July 20, 2020)
FEATURES:
* **New Data Source:** `google_service_account_id_token` ([#2269](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2269))
* **New Resource:** `google_cloudiot_device` ([#2266](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2266))

IMPROVEMENTS:
* bigquery: added support for BigQuery custom schemas for external data using CSV / NDJSON ([#2264](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2264))
* datafusion: changed `version` field to be settable in `google_data_fusion_instance` resource ([#2268](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2268))

BUG FIXES:
* container: fixed a bug where `useIpAlias` was not defaulting to true inside the `ip_allocation_policy` block ([#2260](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2260))
* memcache: fixed field `memcache_parameters` to work correctly on `google_memcache_instance` ([#2261](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2261))

## 3.30.0 (July 13, 2020)
FEATURES:
* **New Data Source:** `google_game_services_game_server_deployment_rollout` ([#2258](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2258))
* **New Resource:** `google_os_config_patch_deployment` ([#2253](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2253))

IMPROVEMENTS:
* artifactregistry: Added field `kms_key_name` to `google_artifact_registry_repository` ([#2254](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2254))

BUG FIXES:
* container: added the ability to update `database_encryption` without recreating the cluster. ([#2259](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2259))
* container: fixed a bug where useIpAlias was not defaulting to true inside the `ip_allocation_policy` block ([#2260](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2260))
* endpoints: fixed `google_endpoints_service` to allow dependent resources to plan based on the `config_id` value. ([#2248](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2248))
* runtimeconfig: fixed `Requested entity was not found.` error when config was deleted outside of terraform. ([#2257](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2257))

## 3.29.0 (July 06, 2020)
NOTES:
* added the `https://www.googleapis.com/auth/cloud-identity` scope to the provider by default ([#2224](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2224))
* `google_app_engine_*_version`'s `service` field is required; previously it would have passed validation but failed on apply if it were absent. ([#6720](https://github.com/terraform-providers/terraform-provider-google/pull/6720))

FEATURES:
* **New Data Source:** `google_cloud_identity_group_memberships` ([#2240](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2240))
* **New Data Source:** `google_cloud_identity_groups` ([#2240](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2240))
* **New Resource:** `google_cloud_identity_group_membership` ([#2224](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2224))
* **New Resource:** `google_cloud_identity_group` ([#2224](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2224))
* **New Resource:** `google_kms_key_ring_import_job` ([#2225](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2225))
* **New Resource:** `google_folder_iam_audit_config` ([#2237](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2237))

IMPROVEMENTS:
* bigquery: Added `"HOUR"` option for `google_bigquery_table` time partitioning (`type`) ([#2235](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2235))
* compute: Added `mode` to `google_compute_region_autoscaler` `autoscaling_policy` ([#2226](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2226))
* compute: Added `scale_down_control` to `google_compute_region_autoscaler` `autoscaling_policy` ([#2226](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2226))
* container: added `networking_mode` to `google_container_cluster` ([#2243](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2243))
* endpoints: enable `google_endpoints_service`-dependent resources to plan based on the `config_id` value. ([#2248](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2248))
* monitoring: added `request_method`, `content_type`, and `body` fields within the `http_check` object to `google_monitoring_uptime_check_config` resource ([#2233](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2233))

BUG FIXES:
* compute: Fixed an issue in `google_compute_managed_ssl_certificate` where multiple fully qualified domain names would cause a permadiff ([#2241](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2241))
* compute: fixed an issue in `compute_url_map` where `path_matcher` sub-fields would conflict with `default_service` ([#2247](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2247))
* container: fixed an issue in `google_container_cluster` where `workload_metadata_config` would cause a permadiff ([#2242](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2242))

## 3.28.0 (June 29, 2020)

FEATURES:
* **New Data Source:** `google_redis_instance` ([#2209](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2209))
* **New Resource:** `google_notebook_environment` ([#2199](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2199))
* **New Resource:** `google_notebook_instance` ([#2199](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2199))

IMPROVEMENTS:
* appengine: Enabled provisioning Firestore on a new project by adding the option to specify `database_type` in `google_app_engine_application` ([#2193](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2193))
* compute: Added `mode` to `google_compute_autoscaler` `autoscaling_policy` ([#2214](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2214))
* compute: Added `remove_instance_state_on_destroy` to `google_compute_per_instance_config` to control deletion of underlying instance state. ([#2187](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2187))
* compute: Added `remove_instance_state_on_destroy` to `google_compute_region_per_instance_config` to control deletion of underlying instance state. ([#2187](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2187))
* compute: Added `scale_down_control` for `google_compute_autoscaler` `autoscaling_policy` ([#2214](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2214))
* compute: Added SHARED_LOADBALANCER_VIP as an option for `google_compute_address.purpose` ([#2204](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2204))
* dns: enabled `google_dns_policy` to accept network id ([#2189](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2189))

BUG FIXES:
* appengine: Added polling to `google_app_engine_firewall_rule` to prevent issues with eventually consistent creation ([#2197](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2197))
* compute: Allowed updating `google_compute_network_peering_routes_config ` `import_custom_routes` and  `export_custom_routes` to false ([#2190](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2190))
* netblock: fixed the google netblock ranges returned by the `google_netblock_ip_ranges` by targeting json on gstatic domain instead of reading SPF dns records (solution provided by network team) ([#2210](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2210))

## 3.27.0 (June 23, 2020)

IMPROVEMENTS:
* accesscontextmanager: Added `custom` config to `google_access_context_manager_access_level` ([#2180](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2180))
* cloudbuild: Added `invert_regex` flag in Github PullRequestFilter and PushFilter in triggerTemplate ([#2171](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2171))
* cloudrun: Added `template.spec.timeout_seconds` to `google_cloud_run_service` ([#2164](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2164))
* compute: Added `cpu_over_commit_type` to `google_compute_node_template` ([#2176](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2176))
* compute: Added `min_node_cpus` to the `scheduling` blocks on `compute_instance` and `compute_instance_template` ([#2169](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2169))
* compute: Added `export_subnet_routes_with_public_ip` and `import_subnet_routes_with_public_ip` to `google_compute_network_peering` ([#2170](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2170))
* compute: Added `remove_instance_state_on_destroy` to `google_compute_per_instance_config` to control deletion of underlying instance state. ([#2187](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2187))
* container: Added support for `private_cluster_config.master_global_access_config` to `google_container_cluster` ([#2157](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2157))
* compute: Added support for `google_compute_instance_group` `instances` to accept instance id field as well as self_link ([#2161](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2161))
* dns: Added support for `google_dns_policy` network to accept `google_compute_network.id` ([#2189](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2189))
* redis: Added validation for name attribute in `redis_instance` ([#2167](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2167))

BUG FIXES:
* bigquery: Fixed `range_partitioning.range.start` so that the value `0` is sent in `google_bigquery_table` ([#2153](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2153))
* container: Fixed a regression in `google_container_cluster` where the location was not inferred when using a `subnetwork` shortname value like `name` ([#2160](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2160))
* datastore: Added retries to `google_datastore_index` requests when under contention. ([#2154](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2154))
* kms: Fixed the `id` value in the `google_kms_crypto_key_version` datasource to include a `/v1` part following `//cloudkms.googleapis.com/`, making it useful for interpolation into Binary Authorization. ([#2165](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2165))


## 3.26.0 (June 15, 2020)

FEATURES:
* **New Resource:** `google_data_catalog_tag` ([#2144](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2144))
* **New Resource:** `google_bigquery_dataset_iam_binding` ([#2147](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2147))
* **New Resource:** `google_bigquery_dataset_iam_member` ([#2147](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2147))
* **New Resource:** `google_bigquery_dataset_iam_policy` ([#2147](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2147))
* **New Resource:** `google_memcache_instance` ([#2142](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2142))
* **New Resource:** `google_network_management_connectivity_test` ([#2138](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2138))

IMPROVEMENTS:
* compute: added `default_route_action` to `compute_url_map` and `compute_url_map.path_matchers` ([#2143](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2143))
* container : Added cluster_telemetry attribute to `google_container_cluster` ([#2149](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2149))
* dialogflow: Changed `google_dialogflow_agent.time_zone` to be updatable ([#2133](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2133))
* dns: enabled google_dns_managed_zone to accept network id for two attributes ([#2139](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2139))
* healthcare: Added support for `streaming_configs` to `google_healthcare_fhir_store` ([#2145](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2145))
* monitoring: added `matcher` attribute to `content_matchers` block for `google_monitoring_uptime_check_config` ([#2150](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2150))

BUG FIXES:
* compute: fixed issue where trying to update the region of `google_compute_subnetwork` would fail instead of destroying/recreating the subnetwork ([#2134](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2134))
* dataflow: added retries in `google_dataflow_job` for common retryable API errors when waiting for job to update ([#2146](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2146))
* dataflow: changed the update logic for `google_dataflow_job` to wait for the replacement job to start successfully before modifying the resource ID to point to the replacement job ([#2140](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2140))

## 3.25.0 (June 08, 2020)
BREAKING CHANGES:
* bigquery: Add ability to manage credentials to `google_bigquery_connection`.  This field is required as the resource is not useful without them. ([#2111](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2111))

FEATURES:
* **New Resource:** `google_data_catalog_tag_template` ([#2120](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2120))
* **New Resource:** `google_container_analysis_occurence` ([#2114](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2114))

IMPROVEMENTS:
* appengine: added `inbound_services` to `StandardAppVersion` resource ([#2131](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2131))
* bigquery: Added support for `google_bigquery_table` `hive_partitioning_options` ([#2121](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2121))
* container_analysis: Added top-level generic note fields to `google_container_analysis_note` ([#2114](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2114))

BUG FIXES:
* bigquery: Fixed an issue where `google_bigquery_job` would return "was present, but now absent" error after job creation ([#2122](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2122))
* container: Changed retry logic for `google_container_node_pool` deletion to use timeouts and retry errors more specifically when cluster is updating. ([#2115](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2115))
* dataflow: fixed an issue where `google_dataflow_job` would try to update `max_workers` ([#2110](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2110))
* dataflow: fixed an issue where updating `on_delete` in `google_dataflow_job` would cause the job to be replaced ([#2110](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2110))
* compute: fixed issue where removing all target pools from `google_compute_instance_group_manager` or `google_compute_region_instance_group_manager` had no effect ([#2124](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2124))
* functions: Added retry to `google_cloudfunctions_function` creation when API returns error while pulling source from GCS ([#2116](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2116))
* provider: Removed credentials from output error when provider cannot parse given credentials ([#2113](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2113))

## 3.24.0 (June 01, 2020)

BREAKING CHANGES:
* bigquery: Add ability to manage credentials to `google_bigquery_connection`.  This field is required as the resource is not useful without them. ([#2111](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2111))

FEATURES:
* **New Resource:** `google_compute_machine_image` ([#2109](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2109))
* **New Resources:** `google_data_catalog_entry_group_iam_*` ([#2098](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2098))
* **New Resource:** `google_data_catalog_entry_group` ([#2098](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2098))
* **New Resource:** `google_data_catalog_entry` ([#2100](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2100))

IMPROVEMENTS:
* appengine: added `handlers` to `google_flexible_app_version` ([#2105](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2105))
* bigquery: suppressed diffs between fully qualified URLs and relative paths that reference the same table or dataset in `google_bigquery_job` ([#2107](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2107))
* container: Added update support for `node_config.workload_metadata_config` to `google_container_node_pool` ([#2091](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2091))

BUG FIXES:
* appengine: added ability to fully sync `StandardAppVersion` resources ([#2096](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2096))
* bigquery: Fixed an issue with `google_bigquery_dataset_access` failing for primitive role `roles/bigquery.dataViewer` ([#2092](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2092))
* dataflow: fixed an issue where `google_dataflow_job` would try to update `max_workers` ([#2110](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2110))
* dataflow: fixed an issue where updating `on_delete` in `google_dataflow_job` would cause the job to be replaced ([#2110](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2110))
* os_login: Fixed `google_os_login_ssh_public_key` `key` field attempting to update in-place ([#2094](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2094))

## 3.23.0 (May 25, 2020)

BREAKING CHANGES:
* The base url for the `monitoring` endpoint no longer includes the API version (previously "v3/"). If you use a `monitoring_custom_endpoint`, remove the trailing "v3/". ([#2088](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2088))

FEATURES:
* **New Data Source:** `google_iam_testable_permissions` ([#2071](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2071))
* **New Resource:** `google_monitoring_dashboard` ([#2088](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2088))

IMPROVEMENTS:
* bigquery: Added ability for various `table_id` fields (and one `dataset_id` field) in `google_bigquery_job` to specify a relative path instead of just the table id ([#2079](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2079))
* composer: Added support for `google_composer_environment` `config.private_environment_config.cloud_sql_ipv4_cidr_block` ([#2075](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2075))
* composer: Added support for `google_composer_environment` `config.private_environment_config.web_server_ipv4_cidr_block` ([#2075](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2075))
* composer: Added support for `google_composer_environment` `web_server_network_access_control` for private environments ([#2075](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2075))
* container: Added update support for `node_config.workload_metadata_config` to `google_container_node_pool` ([#2091](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2091))
* container: Added `min_cpu_platform` to google_container_cluster.cluster_autoscaling.auto_provisioning_defaults ([#2086](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2086))
* container: Added `release_channel_default_version` to `data.google_container_engine_versions`, allowing you to find the default version for a release channel ([#2068](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2068))
* container: Added the ability to unspecify `google_container_cluster`'s `min_master_version` field ([#2068](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2068))
* container: Added update support to `google_container_cluster`'s `release_channel` field ([#2068](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2068))
* container: Added `config_connector_config` `google_container_cluster` ([#2064](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2064))
* monitoring: Added window-based SLI to `google_monitoring_slo` ([#2070](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2070))

BUG FIXES:
* compute: Fixed an issue where `google_compute_route` creation failed while VPC peering was in progress. ([#2082](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2082))
* Fixed an issue where data source `google_organization` would ignore exact domain matches if multiple domains were found ([#2085](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2085))
* compute: Fixed `google_compute_interconnect_attachment` `edge_availability_domain` diff when the field is unspecified ([#2084](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2084))
* compute: Fixed error where plan would error if `google_compute_region_disk_resource_policy_attachment` had been deleted outside of terraform. ([#2065](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2065))
* compute: Raise limit on number of `src_ip_ranges` values in `google_compute_security_policy` to supported 10 ([#2076](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2076))
* iam: Fixed an issue where `google_service_account` shows an error after creating the resource ([#2074](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2074))

## 3.22.0 (May 18, 2020)
BREAKING CHANGE:
* `google_bigtable_instance` resources now cannot be destroyed unless `deletion_protection = false` is set in state for the resource. ([#2061](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2061))

FEATURES:
* **New Resource:** `google_compute_region_per_instance_config` ([#2046](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2046))
* **New Resource:** `google_dialogflow_entity_type` ([#2052](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2052))

IMPROVEMENTS:
* bigtable: added `deletion_protection` field to `google_bigtable_instance` to make deleting them require an explicit intent. ([#2061](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2061))
* compute: Added `google_compute_region_backend_service` `portName` parameter ([#2048](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2048))
* dataproc: Updated `google_dataproc_cluster.software_config.optional_components` to include new options. ([#2049](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2049))
* monitoring: Added `request_based` SLI support to `google_monitoring_slo` ([#2058](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2058))
* storage: added `google_storage_bucket` bucket name to the error message when the bucket can't be deleted because it's not empty ([#2059](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2059))

BUG FIXES:
* bigquery: Fixed error where `google_bigquery_dataset_access` resources could not be found post-creation if role was set to a predefined IAM role with an equivalent primative role (e.g. `roles/bigquery.dataOwner` and `OWNER`) ([#2039](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2039))
* compute: Fixed permadiff in `google_compute_instance_template`'s `network_tier`. ([#2054](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2054))
* compute: Removed permadiff or errors on update for `google_compute_backend_service` and `google_compute_region_backend_service` when `consistent_hash` values were previously set on  backend service but are not supported by updated value of `locality_lb_policy` ([#2044](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2044))
* sql: Fixed occasional failure to delete `google_sql_database_instance` and `google_sql_user`. ([#2045](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2045))

## 3.21.0 (May 11, 2020)

FEATURES:
* **New Resource:** `google_compute_per_instance_config` ([#2029](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2029))
* **New Resource:** `google_logging_billing_account_bucket_config` ([#2008](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2008))
* **New Resource:** `google_logging_folder_bucket_config` ([#2008](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2008))
* **New Resource:** `google_logging_organization_bucket_config` ([#2008](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2008))
* **New Resource:** `google_logging_project_bucket_config` ([#2008](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2008))

IMPROVEMENTS:
* all: add configurable timeouts to several resources that did not previously have them ([#2007](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2007))
* bigquery: added `service_account_name` field to `google_bigquery_data_transfer_config` resource ([#2004](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2004))
* cloudfunctions: Added validation to label keys for `google_cloudfunctions_function` as API errors aren't useful. ([#2009](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2009))
* compute: Added support for `stateful_disk` to both `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager`. ([#2006](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2006))
* container: added `kalm_config` addon to `google_container_cluster` ([#2027](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2027))
* dataflow: Added drift detection for `google_dataflow_job` `template_gcs_path` and `temp_gcs_location` fields ([#2021](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2021))
* dataflow: Added support for update-by-replacement to `google_dataflow_job` ([#2021](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2021))
* dataflow: added `additional_experiments` field to `google_dataflow_job` ([#2005](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2005))
* dataproc: added component gateway support to `google_dataproc_cluster` ([#2035](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2035))
* storage: Added retries for `google_storage_bucket_iam_*` on 412 (precondition not met) errors for eventually consistent bucket creation. ([#2011](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2011))

BUG FIXES:
* all: fixed bug where timeouts specified in units other than minutes were getting incorrectly rounded. Also fixed several instances of timeout values being used from the wrong method. ([#2002](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2002))
* accesscontextmanager: Fixed setting `require_screen_lock` to true for `google_access_context_manager_access_level` ([#2010](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2010))
* appengine: Changed `google_app_engine_application` to respect updates in `iap` ([#2000](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2000))
* bigquery: Fixed error where `google_bigquery_dataset_access` resources could not be found post-creation if role was set to a predefined IAM role with an equivalent primative role (e.g. `roles/bigquery.dataOwner` and `OWNER`) ([#2039](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2039))
* bigquery: Fixed the `google_sheets_options` at least one of logic. ([#2030](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2030))
* cloudscheduler: Fixed permadiff for `google_cloud_scheduler_job.retry_config.*` block when API provides default values ([#2028](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2028))
* compute: Added lock to prevent `google_compute_route` from changing while peering operations are happening on its network ([#2016](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2016))
* compute: Stopped force-recreation of `google_compute_backend_service` and `google_compute_backend_service` on updating `locality_lb_policy` ([#2012](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2012))
* compute: fixed issue where the default value for the attribute `advertise_mode` on `google_compte_router_peer` was not populated on import ([#2024](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2024))
* container: Fixed occasional error with `container_node_pool` partially-successful creations not being recorded if an error occurs on the GCP side. ([#2038](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2038))
* container: fixed issue where terraform would error if a gke instance group was deleted out-of-band ([#2015](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2015))
* storage: Fixed setting/reading `google_storage_bucket_object`  metadata on API object ([#2025](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2025))
* storage: Marked the credentials field in `google_storage_object_signed_url` as sensitive so it doesn't expose private credentials. ([#2026](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2026))

## 3.20.0 (May 04, 2020)

* **New Resource:** `google_artifact_registry_repository` ([#1981](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1981))
* **New Resource:** `google_artifact_registry_repository_iam_policy` ([#1981](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1981))
* **New Resource:** `google_artifact_registry_repository_iam_binding` ([#1981](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1981))
* **New Resource:** `google_artifact_registry_repository_iam_member` ([#1981](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1981))
* **New Resource:** `google_bigquery_connection` ([#2014](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2014))

IMPROVEMENTS:
* appengine: Added `automatic_scaling`, `basic_scaling`, and `manual_scaling` to `google_app_engine_standard_app_version` ([#1984](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1984))
* bigquery: added `service_account_name` field to `google_bigquery_data_transfer_config` resource ([#2004](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2004))
* bigtable: added ability to add/remove column families in `google_bigtable_table` ([#1988](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1988))
* cloudfunctions: Added validation to label keys for `google_cloudfunctions_function` as API errors aren't useful. ([#2009](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2009))
* compute: Added support for `stateful_disk` to both `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager`. ([#2006](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2006))
* compute: Added support for default URL redirects to `google_compute_url_map` and `google_compute_region_url_map` ([#1998](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1998))
* dataflow: Added `additional_experiments` field to `google_dataflow_job` ([#2005](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2005))
* dns: Added `service_directory_config` field to`google_dns_managed_zone` ([#1976](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1976))
* compute: Added update of `google_compute_backend_service` and `google_compute_backend_service` field `locality_lb_policy ([#2012](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2012))

BUG FIXES:
* accesscontextmanager: Fixed setting `require_screen_lock` to true for `google_access_context_manager_access_level` ([#2010](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2010))
* appengine: Changed `google_app_engine_application` to respect updates in `iap` ([#2000](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2000))
* storage: Added retries for `google_storage_bucket_iam_*` on 412 (precondition not met) errors for eventually consistent bucket creation. ([#2011](https://github.com/terraform-providers/terraform-provider-google-beta/pull/2011))

## 3.19.0 (April 27, 2020)

FEATURES:
* **New Resource:** `google_bigquery_job` ([#1959](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1959))
* **New Resource:** `google_monitoring_slo` ([#1953](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1953))
* **New Resource:** `google_service_directory_endpoint` ([#1964](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1964))
* **New Resource:** `google_service_directory_namespace` ([#1964](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1964))
* **New Resource:** `google_service_directory_service` ([#1964](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1964))

IMPROVEMENTS:
* bigtable: Reduced the minimum number of nodes for the `bigtable_instace` resource from 3 to 1. ([#1968](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1968))
* container: Added support for `google_container_cluster` Compute Engine persistent disk CSI driver ([#1969](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1969))
* compute: Added support for `google_compute_instance` `resource_policies` field ([#1957](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1957))
* compute: Added support for `google_compute_resource_policy` group placement policies ([#1957](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1957))
* healthcare: Added `schema` field to `google_healthcare_hl7_v2_store` ([#1962](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1962))

BUG FIXES:
* dataproc: Fixed diff when `google_dataproc_cluster` `preemptible_worker_config.0.num_instances` is sized to 0 and other `preemptible_worker_config` subfields are set ([#1954](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1954))
* resourcemanager: added a wait to `google_project` so that projects are more likely to be ready before the resource finishes creation ([#1970](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1970))
* sql: Allowed `binary_log_enabled` to be disabled. ([#1973](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1973))
* sql: Fixed behaviour in `google_sql_database` when the parent instance is deleted, removing it from state ([#1972](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1972))

## 3.18.0 (April 20, 2020)

FEATURES:
* **New Data Source:** `google_firebase_web_app_config` ([#1950](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1950))
* **New Data Source:** `google_firebase_web_app` ([#1950](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1950))
* **New Data Source:** `google_monitoring_app_engine_service` ([#1944](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1944))
* **New Resource:** `google_firebase_web_app` ([#1950](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1950))
* **New Resource:** `google_monitoring_custom_service` ([#1944](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1944))
* **New Resource:** `google_compute_global_network_endpoint` ([#1948](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1948))
* **New Resource:** `google_compute_global_network_endpoint_group` ([#1948](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1948))
* **New Resource:** `google_monitoring_slo` ([#1953](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1953))

IMPROVEMENTS:
* appengine: Added `iap.enabled` field to `google_app_engine_application` resource ([#1943](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1943))
* iam: Added `name` field to `google_organization_iam_custom_role` ([#1951](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1951))
* iam: Added `name` field to `google_project_iam_custom_role` ([#1951](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1951))

BUG FIXES:
* container: Fixed importing/reading `google_container_node_pool` resources in non-RUNNING states ([#1952](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1952))
* container: Made `addons_config.cloudrun_config` able to be updated without recreating and destroying. ([#1942](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1942))
* container: Made `addons_config.dns_cache_config` able to be updated without recreating and destroying. ([#1942](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1942))
* monitoring: Made `display_name` optional on `google_monitoring_notification_channel ` ([#1947](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1947))

## 3.17.0 (April 13, 2020)

FEATURES:
* **New Resource:** `google_bigquery_dataset_access` ([#1924](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1924))
* **New Resource:** `google_dialogflow_intent` ([#1936](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1936))
* **New Resource:** `google_os_login_ssh_public_key` ([#1922](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1922))

IMPROVEMENTS:
* accesscontextmanager: added `spec` and `use_explicit_dry_run_spec` to `google_access_context_manager_service_perimeter` to test perimeter configurations in dry-run mode. ([#1940](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1940))
* compute: Added update support for `google_compute_interconnect_attachment` `admin_enabled` ([#1931](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1931))
* compute: Added field `log_config` to `google_compute_health_check` and `google_compute_region_health_check` to enable health check logging. ([#1934](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1934))
* compute: Added more import formats for `google_compute_instance` ([#1933](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1933))
* sourcerepo: allowed `google_sourcerepo_repo` `pubsub_configs.topic` to accept short topic names in addition to full references. ([#1938](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1938))

BUG FIXES:
* compute: Fixed diff on default value for `google_compute_interconnect_attachment` `admin_enabled` ([#1931](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1931))
* compute: Fixed perma-diff on `google_compute_interconnect_attachment` `candidate_subnets` ([#1931](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1931))
* compute: fixed bug where `google_compute_instance_from_template` instance defaults were overriding `scheduling` ([#1939](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1939))
* iap: `project` can now be unset in `iap_web_iam_member` and will read from the default `project` ([#1935](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1935))
* serviceusage: fixed issue where `google_project_services` attempted to read a project before enabling the API that allows that read ([#1937](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1937))
* sql: fixed error that occurred on `google_sql_database_instance` when `settings.ip_configuration` was set but `ipv4_enabled` was not set to true and `private_network` was not configured, by defaulting `ipv4_enabled` to true. ([#1926](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1926))
* storage: fixed bug where deleting a `google_storage_bucket` that contained non-deletable objects would retry indefinitely ([#1929](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1929))

## 3.16.0 (April 06, 2020)
FEATURES:
* **New Data Source:** `google_monitoring_uptime_check_ips` ([#1912](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1912))
* **New Resource:** `firebase_project_location`: finalizes the firebase location. ([#1919](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1919))

IMPROVEMENTS:
* cloudfunctions: Added `ingress_settings` field to `google_cloudfunctions_function` ([#1898](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1898))
* cloudfunctions: added support for `vpc_connector_egress_settings` to `google_cloudfunctions_function` ([#1904](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1904))
* accesscontextmanager: added `status.vpc_accessible_services` to `google_access_context_manager_service_perimeter` to control which services are available from the perimeter's VPC networks to the restricted Google APIs IP address range. ([#1910](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1910))
* cloudrun: added ability to autogenerate revision name ([#1900](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1900))
* compute: added ability to resize `google_compute_reservation` ([#1908](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1908))
* container: added `enable_resource_consumption_metering` to `resource_usage_export_config` in `google_container_cluster` ([#1901](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1901))
* dns: added ability to update `google_dns_managed_zone.dnssec_config` ([#1914](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1914))
* pubsub: Added `dead_letter_policy` support to `google_pubsub_subscription` ([#1913](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1913))

BUG FIXES:
* compute: Fixed an issue where `port` could not be removed from health checks ([#1906](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1906))
* storage: fixed an issue where `google_storage_bucket_iam_member` showed a diff for bucket self links ([#1918](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1918))

## 3.15.0 (March 30, 2020)

FEATURES:
* **New Resource:** `google_compute_instance_group_named_port` ([#1869](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1869))
* **New Resource:** `google_service_usage_consumer_quota_override` ([#1884](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1884))
* **New Resource:** `google_firebase_project`: enables Firebase for a referenced Google project ([#1885](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1885))
* **New Resource:** `google_iap_brand` ([#1848](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1848))
* **New Resource:** `google_iap_client` ([#1848](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1848))
* **New Resource:** `google_appengine_flexible_app_version` ([#1849](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1849))

IMPROVEMENTS:
* accesscontextmanager: Added `regions` field to `google_access_context_manager_access_level` ([#1882](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1882))
* compute: added support for IAM conditions in `google_compute_subnet_iam_*` IAM resources ([#1877](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1877))
* kms: Added new field "Additional Authenticated Data" for Cloud KMS data source `google_kms_secret` ([#1886](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1886))
* kms: Added new field "Additional Authenticated Data" for Cloud KMS resource `google_kms_secret_ciphertext` ([#1886](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1886))

BUG FIXES:
* kms: Fixed an issue in `google_kms_crypto_key_version` where `public_key` would return empty after apply ([#1879](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1879))
* logging: Fixed import issue with `google_logging_metric` in a non-default project. ([#1876](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1876))
* provider: Fixed an error with resources failing to upload large files (e.g. with `google_storage_bucket_object`) during retried requests ([#1894](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1894))

## 3.14.0 (March 23, 2020)

FEATURES:
* **New Data Source:** `google_compute_instance_serial_port` ([#1860](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1860))
* **New Resource:** `google_compute_region_ssl_certificate` ([#1863](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1863))

IMPROVEMENTS:
* compute: Added new attribute reference `current_status` to the `google_compute_instance` resource ([#1857](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1857))
* container: Added `dns_cache_config` field to `google_container_cluster` resource ([#1853](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1853))
* container: Updated `upgrade_settings` to read defaults from API for the `google_container_node_pool` resource ([#1859](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1859))
* provider: Added provider-wide request retries for common temporary GCP error codes and network errors ([#1856](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1856))
* redis: Added `connect_mode` field to `google_redis_instance` resource ([#1854](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1854))

## 3.13.0 (March 16, 2020)

BREAKING CHANGES:
* dialogflow: Changed `google_dialogflow_agent.time_zone` to ForceNew. Updating this field will require recreation. This is due to a change in API behavior. ([#1827](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1827))

FEATURES:
* **New Resource:** `google_bigquery_reservation` ([#1833](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1833))
* **New Resource:** `google_compute_region_disk_resource_policy_attachment` ([#1836](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1836))
* **New Resource:** `google_sql_source_representation_instance` ([#1832](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1832))

IMPROVEMENTS:
* bigtable: Added support for full-name/id `instance_name` value in `google_bigtable_table` and `google_bigtable_gc_policy` ([#1830](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1830))
* compute: Added `autoscaling_policy` to `google_compute_node_group` ([#1841](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1841))
* compute: Added support for full-name/id `network_endpoint_group` value in `google_network_endpoint` ([#1831](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1831))
* dialogflow: Changed `google_dialogflow_agent` to not read `tier` status ([#1829](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1829))
* monitoring: Added `sensitive_labels` to `google_monitoring_notification_channel` so that labels like `password` and `auth_token` can be managed separately from the other labels and marked as sensitive. ([#1844](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1844))

BUG FIXES:
* all: fixed issue where nested objects were getting sent as null values to GCP on create instead of being omitted from requests ([#1822](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1822))
* cloudfunctions: fixed `vpc_connector` to be updated properly in `google_cloudfunctions_function` ([#1825](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1825))
* compute: fixed `google_compute_security_policy` from allowing two rules with the same priority. ([#1828](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1828))
* compute: fixed bug where `google_compute_instance.scheduling.node_affinities.operator` would incorrectly accept `NOT` rather than `NOT_IN`. ([#1835](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1835))
* container: Fixed issue where `google_container_node_pool` resources created in the 2.X series were failing to update after 3.11. ([#1846](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1846))

## 3.12.0 (March 09, 2020)

IMPROVEMENTS:
* serviceusage: `google_project_service` no longer attempts to enable a service that is already enabled. ([#1814](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1814))
* bigtable: Added support for full-name/id `instance` value in `google_bigtable_app_profile` ([#1804](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1804))
* pubsub: Added polling to ensure correct resource state for negative-cached PubSub resources ([#1816](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1816))

BUG FIXES:
* compute: Fixed a scenario where `google_compute_instance_template` would cause a crash. ([#1812](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1812))
* storage: Added check for bucket retention policy list being empty. ([#1807](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1807))
* storage: Added locking for operations involving `google_storage_*_access_control` resources to prevent errors from ACLs being added at the same time. ([#1806](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1806))
* container: Fixed panic when upgrading `google_container_cluster` with autoscaling block. ([#1766](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1766))

## 3.11.0 (March 02, 2020)

FEATURES:
* **New Data Source:** `google_compute_backend_bucket` ([#1778](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1778))
* **New Resource:** `google_app_engine_service_split_traffic` ([#1785](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1785))
* **New Resource:** `google_compute_packet_mirroring` ([#1791](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1791))
* **New Resource:** Added new resource `google_game_services_game_server_cluster` ([#1789](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1789))
* **New Resource:** Added new resource `google_game_services_game_server_config` ([#1789](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1789))
* **New Resource:** Added new resource `google_game_services_game_server_deployment_rollout` ([#1789](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1789))
* **New Resource:** Added new resource `google_game_services_game_server_deployment` ([#1789](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1789))
* **New Resource:** Added new resource `google_game_services_realm` ([#1789](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1789))

IMPROVEMENTS:
* bigquery: Landed support for range-based partitioning in `google_bigquery_table` ([#1782](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1782))
* compute: added check on `google_compute_router` for non-empty advertised_groups or advertised_ip_ranges values when advertise_mode is DEFAULT in the bgp block. ([#1776](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1776))
* compute: added the ability to manage the status of `google_compute_instance` resources with the `desired_status` field ([#1786](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1786))
* iam: `google_project_iam_member` and `google_project_iam_binding`'s `project` field can be specified with an optional `projects/` prefix ([#1780](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1780))
* storage: added `metadata` to `google_storage_bucket_object`. ([#1779](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1779))

BUG FIXES:
* compute: Updated `google_project` to check for valid permissions on the parent billing account before creating and tainting the resource. ([#1777](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1777))
* container: Fixed panic when upgrading `google_container_cluster` with `autoscaling` block ([#1766](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1766))

## 3.10.0 (February 25, 2020)

BREAKING CHANGES:
* container: Fully removed `use_ip_aliases` and `create_subnetwork` fields to fix misleading diff for removed fields ([#1760](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1760))

FEATURES:
* **New Data Source:** `google_dns_keys` ([#1768](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1768))
* **New Resource:** `google_datastore_index` ([#1755](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1755))
* **New Resource:** `google_storage_hmac_key` ([#1765](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1765))
* **New Resource:** `google_endpoints_service_iam_binding` ([#1761](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1761))
* **New Resource:** `google_endpoints_service_iam_member` ([#1761](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1761))
* **New Resource:** `google_endpoints_service_iam_policy` ([#1761](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1761))

IMPROVEMENTS:
* container: Enabled configuring autoscaling profile in GKE clusters (https://cloud.google.com/kubernetes-engine/docs/concepts/cluster-autoscaler#autoscaling_profiles) ([#1756](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1756))
* container: Allowed import/update/deletion of `google_container_cluster` in error states ([#1759](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1759))
* container: Changed `google_container_node_pool` so node pools created in an error state will be marked as tainted on creation. ([#1758](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1758))
* container: Allowed import/update/deletion of `google_container_node_pool` in error states and updated resource to wait for a stable state after any changes. ([#1758](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1758))
* container: added label_fingerprint to `google_container_cluster` ([#1750](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1750))
* dataflow: added `job_id` field to `google_dataflow_job` ([#1754](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1754))
* dataflow: added computed `type` field to `google_dataflow_job`. ([#1771](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1771))
* healthcare: added `version` field to `google_healthcare_fhir_store` ([#1769](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1769))
* provider: Added retries for common network errors we've encountered. ([#1762](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1762))

## 3.9.0 (February 18, 2020)

FEATURES:
* **New Resource:** `google_container_registry` ([#1725](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1725))

IMPROVEMENTS:
* all: improve error handling of 404s. ([#1728](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1728))
* bigtable: added update support for `display_name` and `instance_type` ([#1751](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1751))
* container: `google_container_cluster` will wait for a stable state after updates. ([#1737](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1737))
* container: added support for `autoscaling_profile` to `google_container_cluster` ([#1756](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1756))
* container: added `boot_disk_kms_key` to `node_config` block. ([#1736](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1736))
* dataflow: added `job_id` field to `google_dataflow_job` ([#1754](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1754))
* dialogflow: improve error handling by increasing retry count ([#1730](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1730))
* resourcemanager: fixed retry behavior for updates in `google_project`, added retries for billing metadata requests ([#1735](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1735))
* sql: add `encryption_key_name` to `google_sql_database_instance` ([#1724](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1724))

BUG FIXES:
* cloudrun: fixed permadiff caused by new API default values on `annotations` and `limits` ([#1727](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1727))
* container: Removed restriction on `auto_provisioning_defaults` to allow both `oauth_scopes` and `service_account` to be set ([#1748](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1748))
* firestore: fixed import of `google_firestore_index` when database or collection were non-default. ([#1741](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1741))
* iam: Fixed an erroneous error during import of IAM resources when a provider default project/zone/region is not defined. ([#1734](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1734))
* kms: Fixed issue where `google_kms_crypto_key_version` datasource would throw an Invalid Index error on plan ([#1740](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1740))

## 3.8.0 (February 10, 2020)
NOTES:
* provider: added documentation for the `id` field for many resources, including format ([#1697](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1697))
BREAKING CHANGES:
* compute: Added conditional requirement of `google_compute_**region**_backend_service` `backend.capacity_scaler` to no longer accept the API default if not INTERNAL. Non-INTERNAL backend services must now specify `capacity_scaler` explicitly and have a total capacity greater than 0. In addition, API default of 1.0 must now be explicitly set and will be treated as nil or zero if not set in config. ([#1707](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1707))

FEATURES:
* **New Data Source:** `secret_manager_secret_version` ([#1708](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1708))
* **New Resource:** `google_access_context_manager_service_perimeter_resource` ([#1712](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1712))
* **New Resource:** `secret_manager_secret_version` ([#1708](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1708))
* **New Resource:** `secret_manager_secret` ([#1708](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1708))
* **New Resource:** `google_dialogflow_agent` ([#1706](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1706))

IMPROVEMENTS:
* appengine: added support for `google_app_engine_application.iap` ([#1703](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1703))
* compute: `google_compute_security_policy` `rule.match.expr` field is now GA ([#1692](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1692))
* compute: added additional validation to `google_cloud_router`'s `bgp.asn` field. ([#1699](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1699))

BUG FIXES:
* bigtable: fixed diff for DEVELOPMENT instances that are returned from the API with one node ([#1704](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1704))
* compute: Fixed `backend.capacity_scaler` to actually set zero (0.0) value. ([#1707](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1707))
* compute: Fixed `google_compute_**region**_backend_service` so it no longer has a permadiff if `backend.capacity_scaler` is unset in config by requiring capacity scaler. ([#1707](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1707))
* compute: updated `google_compute_project_metadata_item` to fail on create if its key is already present in the project metadata. ([#1714](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1714))
* logging: updated `bigquery_options` so the default value from the api will be set in state. ([#1694](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1694))
* sql: undeprecated `settings.ip_configuration.authorized_networks.expiration_time` ([#1691](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1691))

## 3.7.0 (February 03, 2020)

IMPROVEMENTS:
* binaryauthorization: moved from beta API to ga API in anticipation of beta API turndown. ([#1689](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1689))
* dns: `google_dns_managed_zone` added support for Non-RFC1918 fields for reverse lookup and fowarding paths. ([#1685](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1685))
* monitoring: Added `labels` and `user_labels` filters to data source `google_monitoring_notification_channel` ([#1666](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1666))

BUG FIXES:
* bigtable: fixed diff for DEVELOPMENT instances that are returned from the API with one node ([#1704](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1704))
* compute: `google_compute_instance_template` added plan time check for any disks marked `boot` outside of the first disk ([#1684](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1684))
* container: Fixed perma-diff in `google_container_cluster`'s `cluster_autoscaling.auto_provisioning_defaults`. ([#1679](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1679))
* logging: updated `bigquery_options` so the default value from the api will be set in state. ([#1694](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1694))
* storage: Stopped `project-owner` showing up in the diff for `google_storage_bucket_acl` ([#1674](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1674))

## 3.6.0 (January 29, 2020)

KNOWN ISSUES:

* bigtable: due to API changes, bigtable DEVELOPMENT instances may show a diff on `num_nodes`. There will be a fix in the 3.7.0 release of the provider. No known workarounds exist at the moment, but will be tracked in https://github.com/terraform-providers/terraform-provider-google/issues/5492.

FEATURES:
* **New Data Source:** google_monitoring_notification_channel ([#1643](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1643))
* **New Resource:** google_compute_network_peering_routes_config ([#1652](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1652))

IMPROVEMENTS:
* compute: added waiting logic to `google_compute_interconnect_attachment` to avoid modifications when the attachment is UNPROVISIONED ([#1664](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1664))
* compute: made the `google_compute_network_peering` routes fields available in GA ([#1650](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1650))
* datafusion: Added `service_account` field to `google_data_fusion_instance` ([#1660](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1660))
* iap: added support for IAM conditions in `google_iap_tunnel_instance_iam_*` IAM resources ([#1654](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1654))
* resourcemanager: restricted the length of the `description` field of `google_service_account`. It is now limited to 256 characters. ([#1646](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1646))
* scheduler: Added `attempt_deadline` to `google_cloud_scheduler_job`. ([#1639](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1639))
* storage: added `default_event_based_hold` to `google_storage_bucket` ([#1626](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1626))

BUG FIXES:
* compute: Fixed `google_compute_instance_from_template` with existing boot disks ([#1655](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1655))
* compute: Fixed a bug in `google_compute_instance` when attempting to update a field that requires stopping and starting an instance with an encrypted disk ([#1658](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1658))

## 3.5.0 (January 22, 2020)

DEPRECATIONS:
* kms: deprecated `data.google_kms_secret_ciphertext` as there was no way to make it idempotent. Instead, use the `google_kms_secret_ciphertext` resource. ([#1586](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1586))
* sql: deprecated first generation-only fields on `google_sql_database_instance` ([#1628](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1628))

FEATURES:
* **New Resource:** `google_kms_secret_ciphertext` ([#1586](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1586))

IMPROVEMENTS:
* bigtable: added the ability to add/remove clusters from `google_bigtable_instance` ([#1589](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1589))
* compute: added support for other resource types (like a Proxy) as a `target` to `google_compute_forwarding_rule` ([#1630](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1630))
* dataproc: added `lifecycle_config` to `google_dataproc_cluster.cluster_config` ([#1593](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1593))
* iam: updated to allow for empty bindings in `data_source_google_iam_policy` data source ([#1173](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1173))
* provider: added retries for batched requests so failed batches will retry each single request separately. ([#1615](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1615))
* resourcemanager: restricted the length of the `description` field of `google_service_account`. It is now limited to 256 characters. ([#1646](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1646))

BUG FIXES:
* bigtable: Fixed error on reading non-existent `google_bigtable_gc_policy`,  `google_bigtable_instance`,  `google_bigtable_table` ([#1597](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1597))
* cloudfunctions: Fixed validation of `google_cloudfunctions_function` name to allow for 63 characters. ([#1640](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1640))
* cloudtasks: Changed `max_dispatches_per_second` to a double instead of an integer. ([#1633](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1633))
* compute: Added validation for `compute_resource_policy` to no longer allow invalid `start_time` values that weren't hourly. ([#1603](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1603))
* compute: Fixed errors from concurrent creation/deletion of overlapping `google_compute_network_peering` resources. ([#1601](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1601))
* compute: Stopped panic when using `usage_export_bucket` and the setting had been disabled manually. ([#1610](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1610))
* compute: fixed `google_compute_router_nat` timeout fields causing a diff when using a long-lived resource ([#1613](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1613))
* compute: fixed `google_compute_target_https_proxy.quic_override` causing a diff when using a long-lived resource ([#1611](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1611))
* identityplatform: fixed `google_identity_platform_default_supported_idp_config` to correctly allow configuration of both `idp_id` and `client_id` separately ([#1638](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1638))
* monitoring: Stopped `labels` from causing a perma diff on `AlertPolicy` ([#1622](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1622))

## 3.4.0 (January 07, 2020)

DEPRECATIONS:
* kms: deprecated `data.google_kms_secret_ciphertext` as there was no way to make it idempotent. Instead, use the `google_kms_secret_ciphertext` resource. ([#1586](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1586))

BREAKING CHANGES:
* `google_iap_web_iam_*`, `google_iap_web_type_compute_iam_*`, `google_iap_web_type_app_engine_*`,  and `google_iap_app_engine_service_iam_*` resources now support IAM Conditions (beta provider only). If any conditions had been created out of band before this release, take extra care to ensure they are present in your Terraform config so the provider doesn't try to create new bindings with no conditions. Terraform will show a diff that it is adding the condition to the resource, which is safe to apply. ([#1527](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1527))
* `google_kms_key_ring_iam_*` and `google_kms_crypto_key_iam_*` resources now support IAM Conditions (beta provider only). If any conditions had been created out of band before this release, take extra care to ensure they are present in your Terraform config so the provider doesn't try to create new bindings with no conditions. Terraform will show a diff that it is adding the condition to the resource, which is safe to apply. ([#1524](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1524))
* cloudrun: Changed `google_cloud_run_domain_mapping` to correctly match Cloud Run API expected format for `spec.route_name`, {serviceName}, instead of invalid projects/{project}/global/services/{serviceName} ([#1563](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1563))
* compute: Added back ConflictsWith restrictions for ExactlyOneOf restrictions that were removed in v3.3.0 for `google_compute_firewall`, `google_compute_health_check`, and `google_compute_region_health_check`. This effectively changes an API-side failure that was only accessible in v3.3.0 to a plan-time one. ([#1534](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1534))
* logging: Changed `google_logging_metric.metric_descriptors.labels` from a list to a set ([#1559](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1559))
* resourcemanager: Added back ConflictsWith restrictions for ExactlyOneOf restrictions that were removed in v3.3.0 for `google_organization_policy`, `google_folder_organization_policy`, and `google_project_organization_policy`. This effectively changes an API-side failure that was only accessible in v3.3.0 to a plan-time one. ([#1534](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1534))

FEATURES:
* **New Data Source:** `google_sql_ca_certs` ([#1580](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1580))
* **New Resource:** `google_identity_platform_default_supported_idp_config` ([#1523](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1523))
* **New Resource:** `google_identity_platform_inbound_saml_config` ([#1523](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1523))
* **New Resource:** `google_identity_platform_oauth_idp_config` ([#1523](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1523))
* **New Resource:** `google_identity_platform_tenant_default_supported_idp_config` ([#1523](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1523))
* **New Resource:** `google_identity_platform_tenant_inbound_saml_config` ([#1523](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1523))
* **New Resource:** `google_identity_platform_tenant_oauth_idp_config` ([#1523](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1523))
* **New Resource:** `google_identity_platform_tenant` ([#1523](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1523))
* **New Resource:** `google_kms_crypto_key_iam_policy` ([#1554](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1554))
* **New Resource:** `google_kms_secret_ciphertext` ([#1586](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1586))

IMPROVEMENTS:
* composer: Increased default timeouts for `google_composer_environment` ([#1539](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1539))
* compute: Added graceful termination to `container_cluster` create calls so that partially created clusters will resume the original operation if the Terraform process is killed mid create. ([#1533](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1533))
* compute: Fixed `google_compute_disk_resource_policy_attachment` parsing of region from zone to allow for provider-level zone and make error message more accurate` ([#1557](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1557))
* datafusion: Increased default timeouts for `google_data_fusion_instance` ([#1545](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1545))
* datafusion: Increased update timeout for updating `google_data_fusion_instance` ([#1538](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1538))
* healthcare: Enabled request batching for (beta-only) Healthcare API IAM resources `google_healthcare_*_iam_*` to reduce likelihood of errors from very low default write quota. ([#1558](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1558))
* iap: added support for IAM Conditions to the `google_iap_web_iam_*`, `google_iap_web_type_compute_iam_*`, `google_iap_web_type_app_engine_*`,  and `google_iap_app_engine_service_iam_*` resources (beta provider only) ([#1527](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1527))
* kms: added support for IAM Conditions to the `google_kms_key_ring_iam_*` and `google_kms_crypto_key_iam_*` resources (beta provider only) ([#1524](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1524))
* provider: Reduced default `send_after` controlling the time interval after which a batched request sends. ([#1565](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1565))

BUG FIXES:
* all: fixed issue where many fields that were removed in 3.0.0 would show a diff when they were removed from config ([#1585](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1585))
* bigquery: fixed `bigquery_table.encryption_configuration` to correctly recreate the table when modified ([#1591](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1591))
* cloudrun:  Changed `google_cloud_run_domain_mapping` to correctly match Cloud Run API expected format for `spec.route_name`, {serviceName}, instead of invalid projects/{project}/global/services/{serviceName} ([#1563](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1563))
* cloudrun: Changed `cloud_run_domain_mapping` to poll for success or failure and throw an appropriate error when ready status returns as false. ([#1564](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1564))
* cloudrun: Fixed `google_cloudrun_service` to allow update instead of force-recreation for changes in `spec` `env` and `command` fields ([#1566](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1566))
* cloudrun: Removed unsupported update for `google_cloud_run_domain_mapping` to allow force-recreation. ([#1556](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1556))
* cloudrun: Stopped returning an error when a `cloud_run_domain_mapping` was waiting on DNS verification. ([#1587](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1587))
* compute: Fixed `google_compute_backend_service` to allow updating `cdn_policy.cache_key_policy.*` fields to false or empty. ([#1569](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1569))
* compute: Fixed behaviour where `google_compute_subnetwork` did not record a value for `name` when `self_link` was specified. ([#1579](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1579))
* container: fixed issue where an empty variable in `tags` would cause a crash ([#1543](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1543))
* endpoints: Added operation wait for `google_endpoints_service` to fix 403 "Service not found" errors during initial creation ([#1560](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1560))
* logging: Made `google_logging_metric.metric_descriptors.labels` a set to prevent diff from ordering ([#1559](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1559))
* resourcemanager: added retries for `data.google_organization` ([#1553](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1553))
* vpcaccess: marked `network` field as required in order to fail invalid configs at plan-time instead of at apply-time ([#1577](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1577))

## 3.3.0 (December 17, 2019)

BREAKING CHANGES:
* `google_storage_bucket_iam_*` resources now support IAM Conditions (beta provider only). If any conditions had been created out of band before this release, take extra care to ensure they are present in your Terraform config so the provider doesn't try to create new bindings with no conditions. Terraform will show a diff that it is adding the condition to the resource, which is safe to apply. ([#1479](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1479))

FEATURES:
* **New Resource:** `google_compute_region_health_check` is now available in GA ([#1507](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1507))
* **New Resource:** `google_deployment_manager_deployment` ([#1498](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1498))

IMPROVEMENTS:
* bigquery: added `PARQUET` as an option in `google_bigquery_table.external_data_configuration.source_format` ([#1514](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1514))
* compute: Added `allow_global_access` for to `google_compute_forwarding_rule` resource. ([#1511](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1511))
* compute: added support for up to 100 domains on `google_compute_managed_ssl_certificate` ([#1519](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1519))
* dataproc: added support for `security_config` to `google_dataproc_cluster` ([#1492](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1492))
* storage: added support for IAM Conditions to the `google_storage_bucket_iam_*` resources (beta provider only) ([#1479](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1479))
* storage: updated `id` and `bucket` fields for `google_storage_bucket_iam_*` resources to use `b/{bucket_name}` ([#1479](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1479))

BUG FIXES:
* compute: Fixed an issue where interpolated values caused plan-time errors in `google_compute_router_interface`. ([#1517](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1517))
* compute: relaxed ExactlyOneOf restrictions on `google_compute_firewall`, `google_compute_health_check`, and `google_compute_region_health_check` to enable the use of dynamic blocks with those resources. ([#1520](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1520))
* iam: Fixed a bug that causes badRequest errors on IAM resources due to deleted serviceAccount principals ([#1501](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1501))
* resourcemanager: relaxed ExactlyOneOf restrictions on `google_organization_policy `, `google_folder_organization_policy `, and `google_project_organization_policy ` to enable the use of dynamic blocks with those resources. ([#1520](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1520))
* sourcerepo: Fixed a bug preventing repository IAM resources from referencing repositories with the `/` character in their name ([#1521](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1521))
* sql: fixed bug where terraform would keep retrying to create new `google_sql_database_instance` with the name of a previously deleted instance ([#1500](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1500))

## 3.2.0 (December 11, 2019)

DEPRECATIONS:
* compute: deprecated `fingerprint` field in `google_compute_subnetwork`. Its value is now always `""`. ([#1482](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1482))

FEATURES:
* **New Data Source:** `data_source_google_bigquery_default_service_account` ([#1471](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1471))
* **New Resource:** cloudrun: Added support for `google_cloud_run_service` IAM resources: `google_cloud_run_service_iam_policy`, `google_cloud_run_service_iam_binding`, `google_cloud_run_service_iam_member` ([#1456](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1456))

IMPROVEMENTS:
* all: Added `synchronous_timeout` to provider block to allow setting higher per-operation-poll timeouts. ([#1449](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1449))
* bigquery: Added KMS support to `google_bigquery_table` ([#1471](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1471))
* cloudresourcemanager: Added `org_id` field to `google_organization` datasource to expose the raw organization id ([#1485](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1485))
* cloudrun: Stopped requiring the root `metadata` block for `google_cloud_run_service`. ([#1478](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1478))
* compute: added support for `expr` to `google_compute_security_policy.rule.match` ([#1465](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1465))
* compute: added support for `path_rules` to `google_compute_region_url_map` ([#1489](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1489))
* compute: added support for `path_rules` to `google_compute_url_map` ([#1483](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1483))
* compute: added support for `route_rules` to `google_compute_region_url_map` ([#1493](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1493))
* compute: added support for header actions and route rules to `google_compute_url_map` ([#1435](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1435))
* dns: Added `visibility` field to `google_dns_managed_zone` data source ([#1462](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1462))
* sourcerepo: added support for `pubsub_configs` to `google_sourcerepo_repository` ([#1455](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1455))

BUG FIXES:
* dns: fixed 503s caused by high numbers of `dns_record_set`s. ([#1477](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1477))
* logging: updated `exponential_buckets.growth_factor` from integer to double. ([#1484](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1484))
* storage: fixed bug where users without storage.objects.list permissions couldn't delete empty buckets ([#1443](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1443))

## 3.1.0 (December 05, 2019)

BREAKING CHANGES:
* compute: field `peer_ip_address` in `google_compute_router_peer` is now required, to match the API behavior. ([#1396](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1396))

FEATURES:
* **New Resource:** `google_billing_budget` ([#1428](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1428))
* **New Resource:** `google_cloud_tasks_queue` ([#1369](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1369))
* **New Resource:** `google_organization_iam_audit_config` ([#1427](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1427))

IMPROVEMENTS:
* accesscontextmanager: added support for `require_admin_approval` and `require_corp_owned` in `google_access_context_manager_access_level`'s `device_policy`. ([#1403](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1403))
* all: added retries for timeouts while fetching operations ([#1356](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1356))
* cloudbuild: Added build timeout to `google_cloudbuild_trigger` ([#1404](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1404))
* cloudresourcemanager: added support for importing `google_folder` in the form of the bare folder id, rather than requiring `folders/{bare_id}` ([#1430](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1430))
* compute: Updated default timeouts on `google_compute_project_metadata_item`. ([#1436](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1436))
* compute: `google_compute_disk` `disk_encryption_key.raw_key` is now sensitive ([#1445](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1445))
* compute: `google_compute_disk` `source_image_encryption_key.raw_key` is now sensitive ([#1452](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1452))
* compute: `google_compute_network_peering` resource can now be imported ([#1439](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1439))
* compute: computed attribute `management_type` in `google_compute_router_peer` is now available. ([#1396](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1396))
* compute: field `network` can now be specified on `google_compute_region_backend_service`, which allows internal load balancers to target the non-primary interface of an instance. ([#1418](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1418))
* container: Added support for `peering_name` in `google_container_cluster.private_cluster_config`. ([#1438](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1438))
* container: added `auto_provisioning_defaults` to `google_container_cluster.cluster_autoscaling` ([#1434](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1434))
* container: added `upgrade_settings` support  to `google_container_node_pool` ([#1400](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1400))
* container: increased timeouts on `google_container_cluster` and `google_container_node_pool` ([#1386](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1386))
* datafusion: Added `private_instance` and `network_config` fields to `google_data_fusion_instance` ([#1411](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1411))
* kms: enabled use of `user_project_override` for the `kms_crypto_key` resource ([#1422](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1422))
* kms: enabled use of `user_project_override` for the `kms_secret_ciphertext` data source ([#1433](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1433))
* sql: added `root_password` field to `google_sql_database_instance` resource ([#1432](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1432))

BUG FIXES:
* bigquery: fixed an issue where bigquery table id formats from the `2.X` series caused an error at plan time ([#1448](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1448))
* cloudbuild: Fixed incorrect dependency between `trigger_template` and `github` in `google_cloud_build_trigger`. ([#1410](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1410))
* cloudfunctions: Fixed inability to set `google_cloud_functions_function` update timeout. ([#1447](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1447))
* cloudrun: Wait for the cloudrun resource to reach a ready state before returning success. ([#1409](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1409))
* compute: `google_compute_disk` `disk_encryption_key.raw_key` is now sensitive ([#1453](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1453))
* compute: `self_link` in several datasources will now error on invalid values instead of crashing ([#1373](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1373))
* compute: field `advertised_ip_ranges` in `google_compute_router_peer` can now be updated without recreating the resource. ([#1396](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1396))
* compute: marked `min_cpu_platform` on `google_compute_instance` as computed so if it is not specified it will not cause diffs ([#1429](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1429))
* dataproc: Changed default for `google_dataproc_autoscaling_policy` `secondary_worker_config.min_instances` from 2 to 0. ([#1408](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1408))
* dns: Fixed issue causing `google_dns_record_set` deletion to fail when the managed zone ceased to exist before the deletion event. ([#1446](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1446))
* iam: disallowed `deleted:` principals in IAM resources ([#1417](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1417))
* sql: added retries to `google_sql_user` create and update to reduce flakiness ([#1399](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1399))

## 3.0.0 (December 04, 2019)

NOTES:

These are the changes between 3.0.0-beta.1 and the 3.0.0 final release. For changes since 2.20.0, see also the 3.0.0-beta.1 changelog entry below.

**Please see [the 3.0.0 upgrade guide](https://www.terraform.io/docs/providers/google/guides/version_3_upgrade.html) for upgrade guidance.**

BREAKING CHANGES:
* cloudrun: updated `cloud_run_service` to v1. Significant updates have been made to the resource including a breaking schema change. ([#1426](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1426))

BUG FIXES:
* compute: fixed a bug in `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` that created an artificial diff when removing a now-removed field from a config ([#1401](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1401))
* dns: Fixed bug causing `google_dns_managed_zone` datasource to always return a 404 ([#1405](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1405))
* service_networking: fixed "An unknown error occurred" bug when creating multiple google_service_networking_connection resources in parallel ([#1246](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1246))

## 3.0.0-beta.1 (November 15, 2019)

BREAKING CHANGES:

* access_context_manager: Made `os_type` required on block `google_access_context_manager_access_level.basic.conditions.device_policy.os_constraints`. [MM#2665](https://github.com/GoogleCloudPlatform/magic-modules/pull/2665)
* all: changed any id values that could not be interpolated as self_links into values that could [MM#2461](https://github.com/GoogleCloudPlatform/magic-modules/pull/2461)
* app_engine: Made `ssl_management_type` required on `google_app_engine_domain_mapping.ssl_settings` [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* app_engine: Made `shell` required on `google_app_engine_standard_app_version.entrypoint`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* app_engine: Made `source_url` required on `google_app_engine_standard_app_version.deployment.files` and `google_app_engine_standard_app_version.deployment.zip`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* app_engine: Made `split_health_checks ` required on `google_app_engine_application.feature_settings` [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* app_engine: Made `script_path` required on `google_app_engine_standard_app_version.handlers.script`. [MM#2665](https://github.com/GoogleCloudPlatform/magic-modules/pull/2665)
* bigtable: Made `cluster_id` required on `google_bigtable_app_profile.single_cluster_routing`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* bigquery: Made at least one of `range` or `skip_leading_rows` required on `google_bigquery_table.external_data_configuration.google_sheets_options`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* bigquery: Made `role` required on `google_bigquery_dataset.access`. [MM#2665](https://github.com/GoogleCloudPlatform/magic-modules/pull/2665)
* bigtable: Made exactly one of `single_cluster_routing` or `multi_cluster_routing_use_any` required on `google_bigtable_app_profile`. [MM#2665](https://github.com/GoogleCloudPlatform/magic-modules/pull/2665)
* binary_authorization: Made `name_pattern` required on `google_binary_authorization_policy.admission_whitelist_patterns`. [MM#2665](https://github.com/GoogleCloudPlatform/magic-modules/pull/2665)
* binary_authorization: Made `evaluation_mode` and `enforcement_mode` required on `google_binary_authorization_policy.cluster_admission_rules`. [MM#2665](https://github.com/GoogleCloudPlatform/magic-modules/pull/2665)
* cloudbuild: made Cloud Build Trigger's trigger template required to match API requirements. [MM#2352](https://github.com/GoogleCloudPlatform/magic-modules/pull/2352)
* cloudbuild: Made `branch` required on `google_cloudbuild_trigger.github`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* cloudbuild: Made `steps` required on `google_cloudbuild_trigger.build`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* cloudbuild: Made `name` required on `google_cloudbuild_trigger.build.steps`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* cloudbuild: Made `name` and `path` required on `google_cloudbuild_trigger.build.steps.volumes`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* cloudbuild: Made exactly one of `filename` or `build` required on `google_cloudbuild_trigger`. [MM#2665](https://github.com/GoogleCloudPlatform/magic-modules/pull/2665)
* cloudfunctions: deprecated `nodejs6` as option for `runtime` in `function` and made it required. [MM#2499](https://github.com/GoogleCloudPlatform/magic-modules/pull/2499)
* cloudscheduler: Made exactly one of `pubsub_target`, `http_target` or `app_engine_http_target` required on `google_cloudscheduler_job`. [MM#2665](https://github.com/GoogleCloudPlatform/magic-modules/pull/2665)
* cloudiot: removed `event_notification_config` (singular) from `google_cloudiot_registry`. Use plural `event_notification_configs` instead. [MM#2390](https://github.com/GoogleCloudPlatform/magic-modules/pull/2390)
* cloudiot: Made `public_key_certificate` required on `google_cloudiot_registry. credentials `. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* cloudscheduler: Made `service_account_email` required on `google_cloudscheduler_job.http_target.oauth_token` and `google_cloudscheduler_job.http_target.oidc_token`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* composer: Made at least one of `airflow_config_overrides`, `pypi_packages`, `env_variables, `image_version`, or `python_version` required on `google_composer_environment.config.software_config`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* composer: Made `use_ip_aliases` required on `google_composer_environment.config.node_config.ip_allocation_policy`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* composer: Made `enable_private_endpoint` required on `google_composer_environment.config.private_environment_config`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* composer: Made at least one of `enable_private_endpoint` or `master_ipv4_cidr_block` required on `google_composer_environment.config.private_environment_config` [MM#2682](https://github.com/GoogleCloudPlatform/magic-modules/pull/2682)
* composer: Made at least one of `node_count`, `node_config`, `software_config` or `private_environment_config` required on `google_composer_environment.config` [MM#2682](https://github.com/GoogleCloudPlatform/magic-modules/pull/2682)
* compute: `google_compute_backend_service`'s `backend` field field now requires the `group` subfield to be set. [MM#2373](https://github.com/GoogleCloudPlatform/magic-modules/pull/2373)
* compute: permanently removed `ip_version` field from `google_compute_forwarding_rule` [MM#2436](https://github.com/GoogleCloudPlatform/magic-modules/pull/2436)
* compute: permanently removed `ipv4_range` field from `google_compute_network`. [MM#2436](https://github.com/GoogleCloudPlatform/magic-modules/pull/2436)
* compute: permanently removed `auto_create_routes` field from `google_compute_network_peering`. [MM#2436](https://github.com/GoogleCloudPlatform/magic-modules/pull/2436)
* compute: added check to only allow `google_compute_instance_template`s with 375gb scratch disks [MM#2495](https://github.com/GoogleCloudPlatform/magic-modules/pull/2495)
* compute: made `google_compute_instance_template` fail at plan time when scratch disks do not have `disk_type` `"local-ssd"`. [MM#2282](https://github.com/GoogleCloudPlatform/magic-modules/pull/2282)
* compute: removed `enable_flow_logs` field from `google_compute_subnetwork`. This is now controlled by the presence of the `log_config` block [MM#2597](https://github.com/GoogleCloudPlatform/magic-modules/pull/2597)
* compute: Made `raw_key` required on `google_compute_snapshot.snapshot_encryption_key`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* compute: Made at least one of `auto_delete`, `device_name`, `disk_encryption_key_raw`, `kms_key_self_link`, `initialize_params`, `mode` or `source` required on `google_compute_instance.boot_disk`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* compute: Made at least one of `size`, `type`, `image`, or `labels` required on `google_compute_instance.boot_disk.initialize_params`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* compute: Made at least one of `enable_secure_boot`, `enable_vtpm`, or `enable_integrity_monitoring` required on `google_compute_instance.shielded_instance_config`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* compute: Made at least one of `on_host_maintenance`, `automatic_restart`, `preemptible`, or `node_affinities` required on `google_compute_instance.scheduling`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* compute: Made `interface` required on `google_compute_instance.scratch_disk`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* compute: Made at least one of `enable_secure_boot`, `enable_vtpm`, or `enable_integrity_monitoring` required on `google_compute_instance_template.shielded_instance_config`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* compute: Made at least one of `on_host_maintenance`, `automatic_restart`, `preemptible`, or `node_affinities` are now required on `google_compute_instance_template.scheduling`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* compute: Made `kms_key_self_link` required on `google_compute_instance_template.disk.disk_encryption_key`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* compute: Made `range` required on `google_compute_router_peer. advertised_ip_ranges`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* compute: Removed `instance_template` for `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager`. Use `version.instance_template` instead. [MM#2595](https://github.com/GoogleCloudPlatform/magic-modules/pull/2595)
* compute: removed `update_strategy` for `google_compute_instance_group_manager`. Use `update_policy` instead. [MM#2595](https://github.com/GoogleCloudPlatform/magic-modules/pull/2595)
* compute: stopped allowing selfLink or path style references as IP addresses for `google_compute_forwarding_rule` or `google_compute_global_forwarding_rule` [MM#2620](https://github.com/GoogleCloudPlatform/magic-modules/pull/2620)
* compute: permanently removed `update_strategy` field from `google_compute_region_instance_group_manager`. [MM#2436](https://github.com/GoogleCloudPlatform/magic-modules/pull/2436)
* compute: Made exactly one of `http_health_check`, `https_health_check`, `http2_health_check`, `tcp_health_check` or `ssl_health_check` required on `google_compute_health_check`. [MM#2665](https://github.com/GoogleCloudPlatform/magic-modules/pull/2665)
* compute: Made exactly one of `http_health_check`, `https_health_check`, `http2_health_check`, `tcp_health_check` or `ssl_health_check` required on `google_compute_region_health_check`. [MM#2665](https://github.com/GoogleCloudPlatform/magic-modules/pull/2665)
* container: permanently removed `zone` and `region` fields from data source `google_container_engine_versions`. [MM#2436](https://github.com/GoogleCloudPlatform/magic-modules/pull/2436)
* container: permanently removed `zone`, `region` and `additional_zones` fields from `google_container_cluster`. [MM#2436](https://github.com/GoogleCloudPlatform/magic-modules/pull/2436)
* container: permanently removed `zone` and `region` fields from `google_container_node_pool`. [MM#2436](https://github.com/GoogleCloudPlatform/magic-modules/pull/2436)
* container: set `google_container_cluster`'s `logging_service` and `monitoring_service` defaults to enable GKE Stackdriver Monitoring. [MM#2471](https://github.com/GoogleCloudPlatform/magic-modules/pull/2471)
* container: removed `kubernetes_dashboard` from `google_container_cluster.addons_config` [MM#2551](https://github.com/GoogleCloudPlatform/magic-modules/pull/2551)
* container: removed automatic suppression of GPU taints in GKE `taint` [MM#2537](https://github.com/GoogleCloudPlatform/magic-modules/pull/2537)
* container: Made `disabled` required on `google_container_cluster.addons_config.http_load_balancing`, `google_container_cluster.addons_config.horizontal_pod_autoscaling`, `google_container_cluster.addons_config.network_policy_config`, `google_container_cluster.addons_config.cloudrun_config`, and `google_container_cluster.addons_config.istio_config`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* container: Made at least one of `http_load_balancing`, `horizontal_pod_autoscaling` , `network_policy_config`, `cloudrun_config`, or `istio_config` required on `google_container_cluster.addons_config`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* container: Made `enabled` required on `google_container_cluster.network_policy`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* container: Made `enable_private_endpoint` required on `google_container_cluster.private_cluster_config`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* container: Made `enabled` required on `google_container_cluster.vertical_pod_autoscaling`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* container: Made `cidr_blocks` required on `google_container_cluster.master_authorized_networks_config`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* container: Made at least one of `username`, `password` or `client_certificate_config` required on `google_container_cluster.master_auth`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* container: Made exactly one of `daily_maintenance_window` or `recurring_window` required on `google_container_cluster.maintenance_policy` [MM#2682](https://github.com/GoogleCloudPlatform/magic-modules/pull/2682)
* container: removed `google_container_cluster` `ip_allocation_policy.use_ip_aliases`. If it's set to true, remove it from your config. If false, remove `ip_allocation_policy` as a whole. [MM#2615](https://github.com/GoogleCloudPlatform/magic-modules/pull/2615)
* container: removed `google_container_cluster` `ip_allocation_policy.create_subnetwork`, `ip_allocation_policy.subnetwork_name`, `ip_allocation_policy.node_ipv4_cidr_block`. Define an explicit `google_compute_subnetwork` and use `subnetwork` instead. [MM#2615](https://github.com/GoogleCloudPlatform/magic-modules/pull/2615)
* container: Made `channel` required on `google_container_cluster.release_channel`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* dataproc: Made at least one of `staging_bucket`, `gce_cluster_config`, `master_config`, `worker_config`, `preemptible_worker_config`, `software_config`, `initialization_action` or `encryption_config` required on `google_dataproc_cluster.cluster_config`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* dataproc: Made at least one of `zone`, `network`, `subnetwork`, `tags`, `service_account`, `service_account_scopes`, `internal_ip_only` or `metadata` required on `google_dataproc_cluster.cluster_config.gce_cluster_config`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* dataproc: Made at least one of `num_instances`, `image_uri`, `machine_type`, `min_cpu_platform`, `disk_config`, or `accelerators` required on `google_dataproc_cluster.cluster_config.master_config` and `google_dataproc_cluster.cluster_config.worker_config`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* dataproc: Made at least one of `num_local_ssds`, `boot_disk_size_gb` or `boot_disk_type` required on `google_dataproc_cluster.cluster_config.preemptible_worker_config.disk_config`, `google_dataproc_cluster.cluster_config.master_config.disk_config` and `google_dataproc_cluster.cluster_config.worker_config.disk_config`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* dataproc: Made at least one of `num_instances` or `disk_config` required on `google_dataproc_cluster.cluster_config.preemptible_worker_config`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* dataproc: Made at least one of `image_version`, `override_properties` or `optional_components` is now required on `google_dataproc_cluster.cluster_config.software_config`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* dataproc: Made `policy_uri` required on `google_dataproc_cluster.cluster_config.autoscaling_config`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* dataproc: Made `max_failures_per_hour` required on `google_dataproc_job.scheduling`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* dataproc: Made `driver_log_levels` required on `google_dataproc_job.pyspark_config.logging_config`, `google_dataproc_job.spark_config.logging_config`, `google_dataproc_job.hadoop_config.logging_config`, `google_dataproc_job.hive_config.logging_config`, `google_dataproc_job.pig_config.logging_config`, `google_dataproc_job.sparksql_config.logging_config`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* dataproc: Made at least one of `main_class` or `main_jar_file_uri` required on `google_dataproc_job.spark_config` and `google_dataproc_job.hadoop_config`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* dataproc: Made at least one of `query_file_uri` or `query_list` required on `google_dataproc_job.hive_config`, `google_dataproc_job.pig_config`, and `google_dataproc_job.sparksql_config`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* dns: Made `networks` required on `google_dns_managed_zone.private_visibility_config`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* dns: Made `network_url` required on `google_dns_managed_zone.private_visibility_config.networks`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* iam: made `iam_audit_config` resources overwrite existing audit config on create. Previous implementations merged config with existing audit configs on create. [MM#2438](https://github.com/GoogleCloudPlatform/magic-modules/pull/2438)
* iam: Made exactly one of `list_policy`, `boolean_policy`, or `restore_policy` required on `google_organization_policy`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* iam: Made exactly one of `all` or `values` required on `google_organization_policy.list_policy.allow` and `google_organization_policy.list_policy.deny`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* iam: `google_project_iam_policy` can handle the `project` field in either of the following forms: `project-id` or `projects/project-id` [MM#2700](https://github.com/GoogleCloudPlatform/magic-modules/pull/2700)
* iam: Made exactly one of `allow` or `deny` required on `google_organization_policy.list_policy` [MM#2682](https://github.com/GoogleCloudPlatform/magic-modules/pull/2682)
* iam: removed the deprecated `pgp_key`, `private_key_encrypted` and `private_key_fingerprint` from `google_service_account_key` [MM#2680](https://github.com/GoogleCloudPlatform/magic-modules/pull/2680)
* monitoring: permanently removed `is_internal` and `internal_checkers` fields from `google_monitoring_uptime_check_config`. [MM#2436](https://github.com/GoogleCloudPlatform/magic-modules/pull/2436)
* monitoring: permanently removed `labels` field from `google_monitoring_alert_policy`. [MM#2436](https://github.com/GoogleCloudPlatform/magic-modules/pull/2436)
* monitoring: Made `content` required on `google_monitoring_uptime_check_config.content_matchers`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* monitoring: Made exactly one of `http_check` or `tcp_check` is now required on `google_monitoring_uptime_check_config`. [MM#2665](https://github.com/GoogleCloudPlatform/magic-modules/pull/2665)
* monitoring: Made at least one of `auth_info`, `port`, `headers`, `path`, `use_ssl`, or `mask_headers` is now required on `google_monitoring_uptime_check_config.http_check` [MM#2665](https://github.com/GoogleCloudPlatform/magic-modules/pull/2665)
* provider: added the `https://www.googleapis.com/auth/userinfo.email` scope to the provider by default [MM#2473](https://github.com/GoogleCloudPlatform/magic-modules/pull/2473)
* pubsub: removed ability to set a full path for `google_pubsub_subscription.name` (e.g. `projects/my-project/subscriptions/my-subscription`). `name` now must be the shortname (e.g. `my-subscription`) [MM#2561](https://github.com/GoogleCloudPlatform/magic-modules/pull/2561)
* resourcemanager: converted `google_folder_organization_policy` and `google_organization_policy` import format to use slashes instead of colons. [MM#2638](https://github.com/GoogleCloudPlatform/magic-modules/pull/2638)
* serviceusage: removed `google_project_services` [MM#2403](https://github.com/GoogleCloudPlatform/magic-modules/pull/2403)
* serviceusage: stopped accepting `bigquery-json.googleapis.com` in `google_project_service`. Specify `biquery.googleapis.com` instead. [MM#2626](https://github.com/GoogleCloudPlatform/magic-modules/pull/2626)
* sql: Made `name` and `value` required on `google_sql_database_instance.settings.database_flags`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* sql: Made at least one of `binary_log_enabled`, `enabled`, `start_time`, and `location` required on `google_sql_database_instance.settings.backup_configuration`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* sql: Made at least one of `authorized_networks`, `ipv4_enabled`, `require_ssl`, and `private_network` required on `google_sql_database_instance.settings.ip_configuration`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* sql: Made at least one of `day`, `hour`, and `update_track` required on `google_sql_database_instance.settings.maintenance_window`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* sql: Made at least one of `cert`, `common_name`, `create_time`, `expiration_time`, or `sha1_fingerprint` required on `google_sql_database_instance.settings.server_ca_cert`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* sql: Made at least one of `ca_certificate`, `client_certificate`, `client_key`, `connect_retry_interval`, `dump_file_path`, `failover_target`, `master_heartbeat_period`, `password`, `ssl_cipher`, `username`, and `verify_server_certificate` required on `google_sql_database_instance.settings.replica_configuration`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* sql: Made `value` required on `google_sql_database_instance.settings.ip_configuration.authorized_networks`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* storage: permanently removed `is_live` flag from `google_storage_bucket`. [MM#2436](https://github.com/GoogleCloudPlatform/magic-modules/pull/2436)
* storage: Made at least one of `main_page_suffix` or `not_found_page` required on `google_storage_bucket.website`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* storage: Made at least one of `min_time_elapsed_since_last_modification`, `max_time_elapsed_since_last_modification`, `include_prefixes`, or `exclude_prefixes` required on `google_storage_transfer_job.transfer_spec.object_conditions`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* storage: Made at least one of `overwrite_objects_already_existing_in_sink`, `delete_objects_unique_in_sink`, and `delete_objects_from_source_after_transfer` required on `google_storage_transfer_job.transfer_spec.transfer_options`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)
* storage: Made at least one of `gcs_data_source`, `aws_s3_data_source`, or `http_data_source` required on `google_storage_transfer_job.transfer_options`. [MM#2608](https://github.com/GoogleCloudPlatform/magic-modules/pull/2608)

## 2.20.3 (March 10, 2020)

NOTES:
* `2.20.3` is a backport release, and some changes will not appear in `3.X` series releases until `3.12.0`.
To upgrade to `3.X` you will need to perform a large jump in versions, and it is _strongly_ advised that you attempt to upgrade to `3.X` instead of using this release.
* `2.20.3` is primarily a preventative fix, in anticipation of a change in API response messages adding a default value.

BUG FIXES:
* compute: fixed error when reading `google_compute_instance_template` resources with `network_interface[*].name` set. ([#1815](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1815))

## 2.20.2 (February 04, 2020)

BUG FIXES:
* bigtable: fixed diff for DEVELOPMENT instances that are returned from the API with one node ([#1704](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1704))

## 2.20.1 (December 13, 2019)

BUG FIXES:
* iam: Fixed a bug that causes badRequest errors on IAM resources due to deleted serviceAccount principals ([#1501](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1501))

## 2.20.2 (February 03, 2020)

BUG FIXES:
* bigtable: fixed diff for DEVELOPMENT instances that are returned from the API with one node ([#1704](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1704))

## 2.20.1 (December 13, 2019)

**Note**: 2.20.1 is a backport release. The changes in it are unavailable in 3.0.0-beta.1 through 3.2.0.

BUG FIXES:
* iam: Fixed a bug that causes badRequest errors on IAM resources due to deleted serviceAccount principals ([#1501](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1501))

## 2.20.0 (November 13, 2019)

BREAKING CHANGES:
* `google_compute_instance_iam_*` resources now support IAM Conditions. If any conditions had been created out of band before this release, take extra care to ensure they are present in your Terraform config so the provider doesn't try to create new bindings with no conditions. Terraform will show a diff that it is adding the condition to the resource, which is safe to apply. ([#1360](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1360))
* `google_iap_app_engine_version_iam_*` resources now support IAM Conditions. If any conditions had been created out of band before this release, take extra care to ensure they are present in your Terraform config so the provider doesn't try to create new bindings with no conditions. Terraform will show a diff that it is adding the condition to the resource, which is safe to apply. ([#1352](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1352))
* `google_iap_web_backend_service_iam_*` resources now support IAM Conditions. If any conditions had been created out of band before this release, take extra care to ensure they are present in your Terraform config so the provider doesn't try to create new bindings with no conditions. Terraform will show a diff that it is adding the condition to the resource, which is safe to apply. ([#1352](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1352))
* `google_project_iam_*` resources now support IAM Conditions. If any conditions had been created out of band before this release, take extra care to ensure they are present in your Terraform config so the provider doesn't try to create new bindings with no conditions. Terraform will show a diff that it is adding the condition to the resource, which is safe to apply. ([#1321](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1321))
* compute: the `backend.group` field is now required for `google_compute_region_backend_service`. Configurations without this would not have worked, so this isn't considered an API break. ([#1311](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1311))

FEATURES:
* **New Resource:** `google_data_fusion_instance` ([#1339](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1339))

IMPROVEMENTS:
* bigtable: added import support to `google_bigtable_table` ([#1350](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1350))
* compute: `load_balancing_scheme` for `google_compute_forwarding_rule` now accepts `INTERNAL_MANAGED` as a value. ([#1311](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1311))
* compute: added support for L7 ILB to google_compute_region_backend_service. ([#1311](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1311))
* compute: extended backend configuration options for `google_compute_region_backend_service` to include `backend.balancing_mode`, `backend.capacity_scaler`, `backend.max_connections`, `backend.max_connections_per_endpoint`, `backend.max_connections_per_instance`, `backend.max_rate`, `backend.max_rate_per_endpoint`, `backend.max_rate_per_instance`, and `backend.max_utilization` ([#1311](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1311))
* iam: changed the `id` for many IAM resources to the reference resource long name. Updated `instance_name` on `google_compute_instance_iam` and `subnetwork` on `google_compute_subnetwork` to their respective long names in state ([#1360](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1360))
* iap: added support for IAM Conditions to the `google_compute_instance_iam_*` resources ([#1360](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1360))
* iap: added support for IAM Conditions to the `google_iap_app_engine_version_iam_*` resources ([#1352](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1352))
* iap: added support for IAM Conditions to the `google_iap_web_backend_service_iam_*` resources ([#1352](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1352))
* logging: added `display_name` field to `google_logging_metric` resource ([#1344](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1344))
* monitoring: Added `validate_ssl` to `google_monitoring_uptime_check_config` ([#1243](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1243))
* project: added batching functionality to `google_project_service` read calls, so fewer API requests are made ([#1354](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1354))
* resourcemanager: added support for IAM Conditions to the `google_project_iam_*` resources ([#1321](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1321))
* storage: added notification_id field to `google_storage_notification` ([#1368](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1368))

BUG FIXES:
* compute: fixed issue where setting a 0 for `min_replicas` in `google_compute_autoscaler` and `google_compute_region_autoscaler` would set that field to its server-side default instead of 0. ([#1351](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1351))
* dns: fixed crash when `network` blocks are defined without `network_url`s ([#1345](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1345))
* google: used the correct update method for google_service_account.description ([#1362](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1362))
* logging: fixed issue where logging exclusion resources silently failed when being mutated in parallel ([#1329](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1329))

## 2.19.0 (November 05, 2019)

DEPRECATIONS:
* `compute`: deprecated `enable_flow_logs` on `google_compute_subnetwork`. The presence of the `log_config` block signals that flow logs are enabled for a subnetwork ([#1320](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1320))
* `compute`: deprecated `instance_template` for `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` . Use `version.instance_template` instead. ([#1309](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1309))
* `compute`: deprecated `update_strategy` for `google_compute_instance_group_manager` . Use `update_policy` instead. ([#1309](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1309))
* `container`: deprecated `google_container_cluster` `ip_allocation_policy.create_subnetwork`, `ip_allocation_policy.subnetwork_name`, `ip_allocation_policy.node_ipv4_cidr_block`. Define an explicit `google_compute_subnetwork` and use `subnetwork` instead. ([#1312](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1312))
* `container`: deprecated `google_container_cluster` `ip_allocation_policy.use_ip_aliases`. If it's set to true, remove it from your config. If false, remove `ip_allocation_policy` as a whole. ([#1312](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1312))
* `iam`: Deprecated `pgp_key` on `google_service_account_key` resource. See https://www.terraform.io/docs/extend/best-practices/sensitive-state.html for more information. ([#1326](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1326))

BREAKING CHANGES:
* `google_service_account_iam_*` resources now support IAM Conditions. If any conditions had been created out of band before this release, take extra care to ensure they are present in your Terraform config so the provider doesn't try to create new bindings with no conditions. Terraform will show a diff that it is adding the condition to the resource, which is safe to apply. ([#1188](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1188))

FEATURES:
* `compute`: added `google_compute_router` datasource ([#1233](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1233))

IMPROVEMENTS:
* `cloudbuild`: added ability to specify `name` for `cloud_build_trigger` to avoid name collisions when creating multiple triggers at once. ([#1277](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1277))
* `compute`: added support for multiple versions of `instance_template` and granular control of the update policies for `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager`. ([#1309](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1309))
* `container`: added `taint` field in GKE resources to the GA `google` provider ([#1296](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1296))
* `container`: fix a diff created in the cloud console when `MaintenanceExclusions` are added. ([#1310](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1310))
* `container`: added `maintenance_policy.recurring_window` support to `google_container_cluster`, significantly increasing expressive range. ([#1292](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1292))
* `compute`: added `google_compute_instance` support for display device (Virtual Displays) ([#1313](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1313))
* `iam`: added support for IAM Conditions to the `google_service_account_iam_*` resources (beta provider only) ([#1188](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1188))
* `iam`: added `description` to `google_service_account`. ([#1291](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1291))

BUG FIXES:
* `appengine`: Resolved permadiff in `google_app_engine_domain_mapping.ssl_settings.certificate_id`. ([#1303](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1303))
* `storage`: Fixed error in `google_storage_bucket` where locked retention policies would cause a bucket to report failure on all updates (even though updates were applied correctly). ([#1307](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1307))
* `container`: Fixed nil reference to ShieldedNodes. ([#1314](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1314))

## 2.18.1 (October 25, 2019)

BUGS:
* `resourcemanager`: fixed deleting the default network in `google_project` ([#1299](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1299))

## 2.18.0 (October 23, 2019)

KNOWN ISSUES:
* `resourcemanager`: `google_project` `auto_create_network` is failing to delete networks when set to `false`. Use an earlier provider version to resolve.

DEPRECATIONS:
* `container`: The `kubernetes_dashboard` addon is deprecated for `google_container_cluster`. ([#1247](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1247))

FEATURES:
* **New Resource:** `google_app_engine_application_url_dispatch_rules` ([#1262](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1262))

IMPROVEMENTS:
* `all`: increased support for custom endpoints across the provider ([#1244](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1244))
* `appengine`: added the ability to delete the parent service of `google_app_engine_standard_app_version` ([#1222](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1222))
* `container`: Added `shielded_instance_config` attribute to `node_config` ([#1198](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1198))
* `container`: Allow the configuration of release channels when creating GKE clusters. ([#1260](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1260))
* `dataflow`: added `ip_configuration` option to `job`. ([#1284](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1284))
* `pubsub`: Added field `oidc_token` to `google_pubsub_subscription` ([#1265](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1265))
* `sql`: added `location` field to `backup_configuration` block in `google_sql_database_instance` ([#1282](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1282))

BUGS:
* `all`: fixed the custom endpoint version used by older legacy REST clients ([#1274](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1274))
* `bigquery`: fix issue with `google_bigquery_data_transfer_config` `params` crashing on boolean values ([#1263](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1263))
* `cloudrun`: fixed the apiVersion sent in `google_cloud_run_domain_mapping` requests ([#1251](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1251))
* `compute`: added support for updating multiple fields at once to `google_compute_subnetwork` ([#1269](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1269))
* `compute`: fixed diffs in `google_compute_instance_group`'s `network` field when equivalent values were specified ([#1286](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1286))
* `compute`: fixed issues updating `google_compute_instance_group`'s `instances` field when config/state values didn't match ([#1286](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1286))
* `iam`: fixed bug where IAM binding wouldn't replace members if they were deleted outside of terraform. ([#1272](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1272))
* `pubsub`: Fixed permadiff due to interaction of organization policies and `google_pubsub_topic`. ([#1281](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1281))

## 2.17.0 (October 08, 2019)

NOTES:
* An [upgrade guide](https://www.terraform.io/docs/providers/google/version_3_upgrade.html) has been started for the upcoming 3.0.0 release. ([#1220](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1220))
* `google_project_services` users of provider versions prior to `2.17.0` should update, as past versions of the provider will not handle an upcoming rename of `bigquery-json.googleapis.com` to `bigquery.googleapis.com` well. See https://github.com/terraform-providers/terraform-provider-google/issues/4590 for details. ([#1234](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1234))

DEPRECATIONS:
* `google_project_services` ([#1218](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1218))

FEATURES:
* **New Resource:** `google_bigtable_gc_policy` ([#1213](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1213))
* **New Resource:** `google_binary_authorization_attestor_iam_policy` ([#1166](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1166))
* **New Resource:** `google_compute_region_ssl_certificate` ([#1183](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1183))
* **New Resource:** `google_compute_region_target_http_proxy` ([#1183](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1183))
* **New Resource:** `google_compute_region_target_https_proxy` ([#1183](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1183))
* **New Resource:** `google_iap_app_engine_service_iam_*` ([#1205](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1205))
* **New Resource:** `google_iap_app_engine_version_iam_*` ([#1205](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1205))
* **New Resource:** `google_storage_bucket_access_control` ([#1177](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1177))

IMPROVEMENTS:
* all: made `monitoring-read` scope available. ([#1208](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1208))
* bigquery: added support for default customer-managed encryption keys (CMEK) for BigQuery datasets. ([#1081](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1081))
* bigtable: import support added to `google_bigtable_instance` ([#1224](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1224))
* cloudbuild: added `github` field in `google_cloudbuild_trigger`. ([#1229](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1229))
* container: moved `default_max_pods_per_node` to ga. ([#1235](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1235))
* containeranalysis: moved `google_containeranalysis_note` to ga ([#1166](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1166))
* projectservice: added mitigations for bigquery-json to bigquery rename in project service resources. ([#1234](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1234))

BUGS:
* cloudscheduler: Fixed permadiff for `app_engine_http_target.app_engine_routing` on `google_cloud_scheduler_job` ([#1131](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1131))
* compute: Added ability to set `quic_override` on `google_compute_https_target_proxy` to empty. ([#1219](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1219))
* compute: Fix bug where changes to `region_backend_service.backends.failover` was not detected. ([#1236](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1236))
* compute: fixed `google_compute_router_peer` to default if empty for `advertise_mode` ([#1163](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1163))
* compute: fixed perma-diff in `google_compute_router_nat` when referencing subnetwork via `name` ([#1194](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1194))
* compute: fixed perma-diff in `google_compute_router_nat` when referencing subnetwork via `name` ([#1194](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1194))
* container: fixed an overly-aggressive validation for `master_ipv4_cidr_block` in `google_container_cluster` ([#1211](https://github.com/terraform-providers/terraform-provider-google-beta/pull/1211))

## 2.16.0 (September 24, 2019)

KNOWN ISSUES:
* Based on an upstream change, users of the `google_project_services` resource may have seen the `bigquery.googleapis.com` service added and the `bigquery-json.googleapis.com` service removed, causing a diff. This was later reverted, causing another diff. This issue is being tracked as https://github.com/terraform-providers/terraform-provider-google/issues/4590.

FEATURES:
* **New Resource**: `google_compute_region_url_map` is now available. To support this, the `protocol` for `google_compute_region_backend_service` can now be set to `HTTP`, `HTTPS`, `HTTP2`, and `SSL`. ([#1161](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1161))
* **New Resource**: Adds `google_runtimeconfig_config_iam_*` resources ([#1138](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1138))
* **New Resource**: Added `google_compute_resource_policy` and `google_compute_disk_resource_policy_attachment` to manage `google_compute_disk` resource policies as fine-grained resources ([#1085](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1085))

ENHANCEMENTS:
* composer: Add `python_version` and ability to set `image_version` in `google_composer_environment` in the GA provider ([#1143](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1143))
* compute: `google_compute_global_forwarding_rule` now supports `metadata_filters`. ([#1160](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1160))
* compute: `google_compute_backend_service` now supports `locality_lb_policy`, `outlier_detection`, `consistent_hash`, and `circuit_breakers`. ([#1118](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1118))
* compute: Add support for `guest_os_features` to resource `google_compute_image` ([#1156](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1156))
* compute: Added `drain_nat_ips` to `google_compute_router_nat` ([#1155](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1155))
* container: google_container_node_pool now supports node_locations to specify specific node zones. ([#1154](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1154))
* googleapis: `google_netblock_ip_ranges` data source now has a `private-googleapis` field, for the IP addresses used for Private Google Access for services that do not support VPC Service Controls API access. ([#1102](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1102))
* project: `google_project_iam_*` Properly set the `project` field in state ([#1158](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1158))

BUG FIXES:
* cloudiot: Fixed error where `subfolder_matches` were not set in `google_cloudiot_registry` `event_notification_configs` ([#1175](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1175))

## 2.15.0 (September 17, 2019)

FEATURES:
* **New Resource**: `google_iap_web_iam_binding/_member/_policy` are now available for managing IAP web IAM permissions ([#1044](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1044))
* **New Resource**: `google_iap_web_backend_service_binding/_member/_policy` are now available for managing IAM permissions on IAP enabled backend services ([#1044](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1044))
* **New Resource**: `google_iap_web_type_compute_iam_binding/_member/_policy` are now available for managing IAM permissions on IAP enabled compute services ([#1044](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1044))
* **New Resource**: `google_iap_web_type_app_engine_iam_binding/_member/_policy` are now available for managing IAM permissions on IAP enabled App Engine applications ([#1044](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1044))
* **New Resource**: Add the new resource `google_app_engine_domain_mapping` ([#1079](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1079))
* **New Resource**: `google_cloudfunctions_function_iam_policy`, `google_cloudfunctions_function_iam_binding`, and `google_cloudfunctions_function_iam_member` ([#1121](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1121))
* **New Resource**: `google_compute_reservation` allows you to reserve instance capacity in GCE. ([#1086](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1086))
* **New Resource**: `google_compute_region_health_check` is now available. This and `google_compute_health_check` now include additional support for HTTP2 health checks. ([#1058](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1058))

ENHANCEMENTS:
* compute: Added full routing options to `google_compute_router_peer` ([#1104](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1104))
* compute: add `tunnel_id` to `google_compute_vpn_tunnel` and `gateway_id` to `google_compute_vpn_gateway` ([#1106](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1106))
* compute: `google_compute_subnetwork` now includes the `purpose` and `role` fields. ([#1051](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1051))
* compute: add `purpose` field to `google_compute_address` ([#1115](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1115))
* compute: add `mode` option to `google_compute_instance.boot_disk` ([#1119](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1119))
* compute: `google_compute_firewall` does not show a diff if allowed or denied rules are specified with uppercase protocol values ([#1144](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1144))
* compute: Add support for the `log_config` block to `compute_backend_service` (Beta only) ([#1137](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1137))
* logging: added `metric_descriptor.unit` to `google_logging_metric` resource ([#1117](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1117))

BUG FIXES:
* all: More classes of generic HTTP errors are retried provider-wide. ([#1120](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1120))
* container: Fix error when `master_authorized_networks_config` is removed from the `google_container_cluster` configuration. ([#1133](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1133))
* iam: Make `google_service_account_` and `google_service_account_iam_*` validation less restrictive to allow for more default service accounts ([#1109](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1109))
* iam: set auditconfigs in state for google_\*\_iam_policy resources ([#1134](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1134))
* logging: `google_logging_metric` `explicit` bucket option can now be set ([#1096](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1096))
* pubsub: Add retry for Pubsub Topic creation when project is still initializing org policies ([#1094](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1094))
* servicenetworking: remove need for provider-level project to delete connection ([#1132](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1132))
* sql: Add more retries for operationInProgress 409 errors for `google_sql_database_instance` ([#1108](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1108))

MISC:
* The User-Agent header that Terraform sends has been updated to correctly report the version of Terraform being run, and has minorly changed the formatting on the Terraform string. ([#1107](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1107))


## 2.14.0 (August 28, 2019)

DEPRECATIONS:
* cloudiot: `resource_cloudiot_registry`'s `event_notification_config` field has been deprecated. ([#1064](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1064))

FEATURES:
* **New Resource**: `google_bigtable_app_profile` is now available ([#988](https://github.com/terraform-providers/terraform-provider-google-beta/issues/988))
* **New Resource**: `google_ml_engine_model` ([#957](https://github.com/terraform-providers/terraform-provider-google-beta/issues/957))
* **New Resource**: `google_dataproc_autoscaling_policy` ([#1078](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1078))
* **New Data Source**: `google_kms_secret_ciphertext` ([#1011](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1011))

ENHANCEMENTS:
* bigquery: Add support for clustering/partitioning to bigquery_table ([#1025](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1025))
* bigtable: `num_nodes` can now be updated in `google_bigtable_instance` ([#1067](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1067))
* cloudiot: `resource_cloudiot_registry` now has fields plural `event_notification_configs` and `log_level`, and `event_notification_config` has been deprecated. ([#1064](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1064))
* cloud_run: New output-only fields have been added to google_cloud_run_service' status. ([#1071](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1071))
* compute: Adding bandwidth attribute to interconnect attachment. ([#1016](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1016))
* compute: `google_compute_region_instance_group_manager.update_policy` now supports `instance_redistribution_type` ([#1073](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1073))
* compute: adds admin_enabled to google_compute_interconnect_attachment ([#1072](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1072))
* compute: The compute routes includes next_hop_ilb attribute support in beta. ([#1076](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1076))
* scheduler: Add support for `oauth_token` and `oidc_token` on resource `google_cloud_scheduler_job` ([#1024](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1024))

BUG FIXES:
* containerregistry: Correctly handle domain-scoped projects ([#1035](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1035))
* iam: Fixed regression in 2.13.0 for permadiff on empty members in IAM policy bindings. ([#1092](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1092))
* project: `google_project_iam_custom_role` now sets the project properly on import. ([#1089](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1089))
* sql: Added back a missing import format for `google_sql_database`. ([#1061](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1061))

## 2.13.0 (August 15, 2019)

KNOWN ISSUES:
* `bigtable`: `google_bigtable_instance` may cause a panic on Terraform `0.11`. This was resolved in `2.17.0`.

FEATURES:
* **New Resource**: added the `google_vpc_access_connector` resource and the `vpc_connector` option on the `google_cloudfunctions_function` resource. ([#1004](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1004))
* **New Resource**: Added `google_scc_source` resource for managing Cloud Security Command Center sources in Terraform ([#1033](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1033))
* **New Data Source**: `google_compute_network_endpoint_group`([#999](https://github.com/terraform-providers/terraform-provider-google-beta/issues/999))

ENHANCEMENTS:
* bigquery: Added support for `google_bigquery_data_transfer_config` (which include scheduled queries). ([#975](https://github.com/terraform-providers/terraform-provider-google-beta/issues/975))
* bigtable: `google_bigtable_instance` max number of `cluster` blocks is now 4 ([#995](https://github.com/terraform-providers/terraform-provider-google-beta/issues/995))
* binary_authorization: Added `globalPolicyEvaluationMode` to `google_binary_authorization_policy`. ([#987](https://github.com/terraform-providers/terraform-provider-google-beta/issues/987))
* cloudfunctions: Allow partial URIs in google_cloudfunctions_function event_trigger.resource ([#1009](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1009))
* compute: Enable update for `google_compute_router_nat` ([#979](https://github.com/terraform-providers/terraform-provider-google-beta/issues/979))
* netblock: extended `google_netblock_ip_ranges` to support multiple useful IP address ranges that have a special meaning on GCP. ([#986](https://github.com/terraform-providers/terraform-provider-google-beta/issues/986))
* project: Wrapped API requests with retries for `google_project`, `google_folder`, and `google_*_organization_policy` ([#971](https://github.com/terraform-providers/terraform-provider-google-beta/issues/971))
* project: IAM and service requests are now batched ([#1014](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1014))
* provider: allow provider's region to be specified as a self_link ([#1022](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1022))
* provider: Adds new provider-level field `user_project_override`, which allows billing, quota checks, and service enablement checks to occur against the project a resource is in instead of the project the credentials are from. ([#1010](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1010))
* pubsub: Pub/Sub topic geo restriction support. ([#989](https://github.com/terraform-providers/terraform-provider-google-beta/issues/989))

BUG FIXES:
* binary_authorization: don't diff when attestation authority note public keys don't have an ID in the config ([#1042](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1042))
* compute: instance descriptions will now be stored in state ([#990](https://github.com/terraform-providers/terraform-provider-google-beta/issues/990))
* container: `key_name` in `google_container_cluster.database_encryption` is no longer a required field. ([#1032](https://github.com/terraform-providers/terraform-provider-google-beta/issues/1032))
* project: ignore errors when deleting a default network that doesn't exist ([#991](https://github.com/terraform-providers/terraform-provider-google-beta/issues/991))

## 2.12.0 (August 01, 2019)

FEATURES:
* **New Data Source**: `google_kms_crypto_key_version` - Provides access to KMS key version data with Google Cloud KMS. ([#964](https://github.com/terraform-providers/terraform-provider-google-beta/issues/964))
* **New Resource**: `google_cloud_run_service` - Set up a cloud run service ([#757](https://github.com/terraform-providers/terraform-provider-google-beta/issues/757))
* **New Resource**: `google_cloud_run_domain_mapping` - Allows custom domains to map to a cloud run service ([#757](https://github.com/terraform-providers/terraform-provider-google-beta/issues/757))

ENHANCEMENTS:
* binary_authorization: Add support for Cloud KMS PKIX keys to `binary_authorization_attestor`. ([#964](https://github.com/terraform-providers/terraform-provider-google-beta/issues/964))
* composer: Add private IP config for `google_composer_environment` ([#908](https://github.com/terraform-providers/terraform-provider-google-beta/issues/908))
* compute: add support for port_specification to resource `google_compute_health_check` ([#933](https://github.com/terraform-providers/terraform-provider-google-beta/issues/933))
* compute: Fixed import formats for `google_compute_network_endpoint` and add location-only import formats ([#947](https://github.com/terraform-providers/terraform-provider-google-beta/issues/947))
* compute: add support for `resource_policies` to resource `google_compute_disk` ([#960](https://github.com/terraform-providers/terraform-provider-google-beta/issues/960))
* compute: Support labelling for compute_instance boot_disks and compute_instance_template disks. ([#982](https://github.com/terraform-providers/terraform-provider-google-beta/issues/982))
* container: `workload_identity_config` in `google_container_cluster` can now be updated without recreating the cluster. ([#896](https://github.com/terraform-providers/terraform-provider-google-beta/issues/896))
* container: validate that master_ipv4_cidr_block is set if enable_private_nodes is true ([#948](https://github.com/terraform-providers/terraform-provider-google-beta/issues/948))
* dataflow: added support for user-defined `labels` on resource `google_dataflow_job` ([#970](https://github.com/terraform-providers/terraform-provider-google-beta/issues/970))
* dataproc: add support for `optional_components` to resource `resource_dataproc_cluster` ([#961](https://github.com/terraform-providers/terraform-provider-google-beta/issues/961))
* project: add checks to import to prevent importing by project number instead of id ([#954](https://github.com/terraform-providers/terraform-provider-google-beta/issues/954))
* storage: add support for `retention_policy` to resource `google_storage_bucket` ([#949](https://github.com/terraform-providers/terraform-provider-google-beta/issues/949))

BUG FIXES:
* access_context_manager: import format checking ([#952](https://github.com/terraform-providers/terraform-provider-google-beta/issues/952))
* dataproc: Suppress diff for `google_dataproc_cluster` `software_config.0.image_version` to prevent permadiff when server uses more specific versions of config value ([#969](https://github.com/terraform-providers/terraform-provider-google-beta/issues/969))
* organization: Add auditConfigs to update masks for setting org and folder IAM policy (`google_organization_iam_policy`, `google_folder_iam_policy`) ([#967](https://github.com/terraform-providers/terraform-provider-google-beta/issues/967))
* storage: `google_storage_bucket` Set website metadata during read ([#925](https://github.com/terraform-providers/terraform-provider-google-beta/issues/925))

## 2.11.0 (July 16, 2019)

NOTES:
* container: We have changed the way container clusters handle cluster state, and they should now wait until the cluster is ready when creating, updating, or refreshing cluster state. This is meant to decrease the frequency of errors where Terraform is operating on a cluster that isn't ready to be operated on. If this change causes a problem, please open an issue with as much information as you can provide, especially [debug logs](https://www.terraform.io/docs/internals/debugging.html). See [terraform-provider-google #3989](https://github.com/terraform-providers/terraform-provider-google/issues/3989) for more info.

FEATURES:
* **New Resources**: `google_bigtable_instance_iam_binding`, `google_bigtable_instance_iam_member`, and `google_bigtable_instance_iam_policy` are now available. ([#923](https://github.com/terraform-providers/terraform-provider-google-beta/issues/923))
* **New Resources**: `google_sourcerepo_repository_iam_*` Add support for source repo repository IAM resources ([#914](https://github.com/terraform-providers/terraform-provider-google-beta/issues/914))

ENHANCEMENTS:
* bigquery: Added support for `external_data_configuration` to `google_bigquery_table`. ([#696](https://github.com/terraform-providers/terraform-provider-google-beta/issues/696))
* compute: Avoid getting project if no diff found for google_compute_instance_template ([#932](https://github.com/terraform-providers/terraform-provider-google-beta/issues/932))
* firestore: `google_firestore_index` `query_scope` can have `COLLECTION_GROUP` specified. ([#919](https://github.com/terraform-providers/terraform-provider-google-beta/issues/919))

BUG FIXES:
* compute: Mark instance KMS self link field kms_key_self_link as computed ([#819](https://github.com/terraform-providers/terraform-provider-google-beta/issues/819))
* compute: Allow security policy to be removed from `google_backend_service` ([#916](https://github.com/terraform-providers/terraform-provider-google-beta/issues/916))
* container: `google_container_cluster` deeper nil checks to prevent crash on empty object ([#934](https://github.com/terraform-providers/terraform-provider-google-beta/issues/934))
* container: `google_container_cluster` keep clusters in state if they are created in an error state and don't get correctly cleaned up. ([#929](https://github.com/terraform-providers/terraform-provider-google-beta/issues/929))
* container: `google_container_node_pool` Correctly set nodepool autoscaling in state when disabled in the API ([#931](https://github.com/terraform-providers/terraform-provider-google-beta/issues/931))
* container: `google_container_cluster` will now wait to act until the cluster can be operated on, respecting timeouts. ([#927](https://github.com/terraform-providers/terraform-provider-google-beta/issues/927))
* monitoring: Fix diff in `google_monitoring_uptime_check_config` on a deprecated field. ([#944](https://github.com/terraform-providers/terraform-provider-google-beta/issues/944))
* service: `google_service_networking_connection` correctly delete the connection when the resource is destroyed. ([#935](https://github.com/terraform-providers/terraform-provider-google-beta/issues/935))
* spanner: Wait for spanner databases to create before returning. Don't wait for databases to delete before returning anymore. ([#922](https://github.com/terraform-providers/terraform-provider-google-beta/issues/922))
* storage: Fixed an issue where `google_storage_transfer_job` `schedule_end_date` caused requests to fail if unset. ([#936](https://github.com/terraform-providers/terraform-provider-google-beta/issues/936))
* storage: `google_storage_object_acl` Prevent panic when using interpolated object names. ([#917](https://github.com/terraform-providers/terraform-provider-google-beta/issues/917))


## 2.10.0 (July 02, 2019)

DEPRECATIONS:
* monitoring: Deprecated non-existent fields `is_internal` and `internal_checkers` from `google_monitoring_uptime_check_config`. ([#888](https://github.com/terraform-providers/terraform-provider-google-beta/issues/888))

FEATURES:
* **New Resource**: `google_compute_project_default_network_tier` ([#882](https://github.com/terraform-providers/terraform-provider-google-beta/issues/882))
* **New Resource** `google_healthcare_dataset_iam_binding` ([#899](https://github.com/terraform-providers/terraform-provider-google-beta/pull/899))
* **New Resource** `google_healthcare_dataset_iam_member` ([8#99](https://github.com/terraform-providers/terraform-provider-google-beta/pull/899))
* **New Resource** `google_healthcare_dataset_iam_policy` ([#899](https://github.com/terraform-providers/terraform-provider-google-beta/pull/899))
* **New Resource** `google_healthcare_dicom_store_iam_binding` ([#899](https://github.com/terraform-providers/terraform-provider-google-beta/pull/899))
* **New Resource** `google_healthcare_dicom_store_iam_member` ([#899](https://github.com/terraform-providers/terraform-provider-google-beta/pull/899))
* **New Resource** `google_healthcare_dicom_store_iam_policy` ([#899](https://github.com/terraform-providers/terraform-provider-google-beta/pull/899))
* **New Resource** `google_healthcare_fhir_store_iam_binding` ([#899](https://github.com/terraform-providers/terraform-provider-google-beta/pull/899))
* **New Resource** `google_healthcare_fhir_store_iam_member` ([#899](https://github.com/terraform-providers/terraform-provider-google-beta/pull/899))
* **New Resource** `google_healthcare_fhir_store_iam_policy` ([#899](https://github.com/terraform-providers/terraform-provider-google-beta/pull/899))
* **New Resource** `google_healthcare_hl7_v2_store_iam_binding` ([#899](https://github.com/terraform-providers/terraform-provider-google-beta/pull/899))
* **New Resource** `google_healthcare_hl7_v2_store_iam_member` ([#899](https://github.com/terraform-providers/terraform-provider-google-beta/pull/899))
* **New Resource** `google_healthcare_hl7_v2_store_iam_policy` ([#899](https://github.com/terraform-providers/terraform-provider-google-beta/pull/899))

ENHANCEMENTS:
* compute: Added fields for managing network endpoint group backends in `google_compute_backend_service`, including `max_connections_per_endpoint` and `max_rate_per_endpoint` ([#854](https://github.com/terraform-providers/terraform-provider-google-beta/issues/854))
* compute: Support custom timeouts in `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` ([#909](https://github.com/terraform-providers/terraform-provider-google-beta/issues/909))
* container: `node_config.sandbox_config` is supported on GKE node pool definitions, allowing you to configure GKE Sandbox. ([#863](https://github.com/terraform-providers/terraform-provider-google-beta/issues/863))
* container: `google_container_cluster` add support for GKE resource usage ([#825](https://github.com/terraform-providers/terraform-provider-google-beta/issues/825))
* folder: `google_folder` improve error message on delete ([#878](https://github.com/terraform-providers/terraform-provider-google-beta/issues/878))
* iam: sort bindings in `google_*_iam_policy` resources to get simpler diffs ([#881](https://github.com/terraform-providers/terraform-provider-google-beta/issues/881))
* kms: `google_kms_crypto_key` now supports labels. ([#885](https://github.com/terraform-providers/terraform-provider-google-beta/issues/885))
* pubsub: `google_pubsub_topic` supports KMS keys with `kms_key_name`. ([#894](https://github.com/terraform-providers/terraform-provider-google-beta/issues/894))

BUG FIXES:
* iam: the member field in iam_* resources is now case-insensitive ([#876](https://github.com/terraform-providers/terraform-provider-google-beta/issues/876))
* servicenetworking: `google_service_networking_connection` fix update ([#871](https://github.com/terraform-providers/terraform-provider-google-beta/issues/871))

## 2.9.1 (June 21, 2019)

BUG FIXES:
* kms: fix regression when reading existing `google_kms_crypto_key` resources ([#873](https://github.com/terraform-providers/terraform-provider-google-beta/issues/873))
* storage: `google_storage_bucket` fix for crash that occurs when running plan on old buckets ([#870](https://github.com/terraform-providers/terraform-provider-google-beta/issues/870))
* storage: `google_storage_bucket` allow updating bucket_policy_only to false ([#870](https://github.com/terraform-providers/terraform-provider-google-beta/issues/870))

## 2.9.0 (June 19, 2019)

FEATURES:
* **Custom Endpoint Support**: The Google provider supports custom endpoints, allowing you to use GCP-like APIs such as emulators. See the [Provider Reference](https://www.terraform.io/docs/providers/google/provider_reference.html) for details. ([#811](https://github.com/terraform-providers/terraform-provider-google-beta/issues/811))
* **New Resource**: `google_compute_resource_policy` is now available which can be used to schedule disk snapshots. ([#1850](https://github.com/GoogleCloudPlatform/magic-modules/pull/1850))
* **New Resource**: `google_compute_external_vpn_gateway` is now available which can be used to connect to external VPN gateways. ([#833](https://github.com/terraform-providers/terraform-provider-google-beta/issues/833))
* **New Resource** Network endpoint groups (`google_compute_network_endpoint_group`) and fine-grained resource endpoints (`google_compute_network_endpoint`) are now available. ([#781](https://github.com/terraform-providers/terraform-provider-google-beta/issues/781))

ENHANCEMENTS:
* increased default timeouts for `google_compute_instance`, `google_container_cluster`, `google_dataproc_cluster`, and `google_sql_database_instance` ([#862](https://github.com/terraform-providers/terraform-provider-google-beta/issues/862))
* container: `google_container_cluster` Stop guest_accelerator from having a permadiff for accelerators with `count=0` ([#851](https://github.com/terraform-providers/terraform-provider-google-beta/issues/851))
* container: `google_container_cluster` supports `authenticator_groups_config` to allow Google Groups-based authentication. ([#669](https://github.com/terraform-providers/terraform-provider-google-beta/issues/669))
* container: `google_container_cluster` supports `enable_intranode_visibility`. ([#801](https://github.com/terraform-providers/terraform-provider-google-beta/issues/801))
* container: `google_container_cluster` supports Workload Identity to access GCP APIs in GKE applications with `workload_identity_config`. ([#824](https://github.com/terraform-providers/terraform-provider-google-beta/issues/824))
* dataproc: `google_dataproc_cluster` supports `min_cpu_platform` ([#424](https://github.com/terraform-providers/terraform-provider-google-beta/issues/424)], [[#848](https://github.com/terraform-providers/terraform-provider-google-beta/issues/848))
* dns: `google_dns_record_set`: allow importing dns record sets in any project ([#853](https://github.com/terraform-providers/terraform-provider-google-beta/issues/853))
* kms: `kms_crypto_key` supports `purpose` ([#845](https://github.com/terraform-providers/terraform-provider-google-beta/issues/845))
* storage: `google_storage_bucket` now supports enabling `bucket_policy_only` access control. ([#1878](https://github.com/GoogleCloudPlatform/magic-modules/pull/1878))
* storage: IAM resources for storage buckets (`google_storage_bucket_iam_*`) now all support import ([#835](https://github.com/terraform-providers/terraform-provider-google-beta/issues/835))
* pubsub: `google_pubsub_topic` Updates for labels are now supported ([#832](https://github.com/terraform-providers/terraform-provider-google-beta/issues/832))

BUG FIXES:
* bigquery: `google_bigquery_dataset` Relax IAM role restrictions on BQ datasets ([#857](https://github.com/terraform-providers/terraform-provider-google-beta/issues/857))
* compute: `google_project_iam` When importing resources `project` no longer needs to be set in the config post import ([#805](https://github.com/terraform-providers/terraform-provider-google-beta/issues/805))
* compute: `google_sql_user` User's can now be updated to change their password ([#810](https://github.com/terraform-providers/terraform-provider-google-beta/issues/810))
* compute: `google_compute_instance_template` Fixed issue so project can now be specified by interpolated varibles. ([#816](https://github.com/terraform-providers/terraform-provider-google-beta/issues/816))
* compute: `google_compute_instance_template` Throw error when using incompatible disk fields instead of continual plan diff ([#812](https://github.com/terraform-providers/terraform-provider-google-beta/issues/812))
* compute: `google_compute_instance_from_template` Make sure disk type is expanded to a URL ([#771](https://github.com/terraform-providers/terraform-provider-google-beta/issues/771))
* comptue: `google_compute_instance_template` Attempt to put disks in state in the same order they were specified ([#771](https://github.com/terraform-providers/terraform-provider-google-beta/issues/771))
* container: `google_container_cluster` and `google_node_pool` now retry correctly when polling for status of an operation. ([#818](https://github.com/terraform-providers/terraform-provider-google-beta/issues/818))
* container: `google_container_cluster` `istio_config.auth` will no longer permadiff on `AUTH_NONE` when an auth method other than TLS is defined. ([#834](https://github.com/terraform-providers/terraform-provider-google-beta/issues/834))
* dns: `google_dns_record_set` overrides all existing record types on create, not just NS ([#850](https://github.com/terraform-providers/terraform-provider-google-beta/issues/850))
* monitoring: `google_monitoring_notification_channel` Allow setting enabled to false ([#864](https://github.com/terraform-providers/terraform-provider-google-beta/issues/864))
* pubsub: `google_pubsub_subscription` and `google_pubsub_topic` resources can be created inside VPC service controls. ([#827](https://github.com/terraform-providers/terraform-provider-google-beta/issues/827))
* redis: `google_redis_instance` Fall back to region from `location_id` when region isn't specified ([#847](https://github.com/terraform-providers/terraform-provider-google-beta/issues/847))

## 2.8.0 (June 04, 2019)

DEPRECATIONS:
* compute: The `auto_create_routes` field on `google_compute_network_peering` has been deprecated because it is not user configurable. ([#3394](https://github.com/terraform-providers/terraform-provider-google/issues/3394))

FEATURES:
* **New Resource**: `google_compute_ha_vpn_gateway` is now available. This is an alternative to `google_compute_vpn_gateway` that can be set up to provide higher availability. ([#704](https://github.com/terraform-providers/terraform-provider-google-beta/pull/704))
* **New Datasource**: `google_compute_ssl_certificate` ([#742](https://github.com/terraform-providers/terraform-provider-google-beta/pull/742))
* **New Datasource**: `google_composer_image_versions` ([#752](https://github.com/terraform-providers/terraform-provider-google-beta/pull/752))

ENHANCEMENTS:
* app_engine: Remove restrictive `app_engine_application` location validation. ([#760](https://github.com/terraform-providers/terraform-provider-google-beta/pull/760))
* compute: `google_compute_vpn_tunnel` supports HA fields `vpn_gateway`, `vpn_gateway_interface`, `peer_gcp_gateway`, `peer_external_gateway`, `vpn_gateway_interface` ([#704](https://github.com/terraform-providers/terraform-provider-google-beta/pull/704))
* compute: `google_container_cluster` add support for vertical pod autoscaling ([#749](https://github.com/terraform-providers/terraform-provider-google-beta/issues/749))
* compute: `google_compute_router_interface` now supports specifying an `interconnect_attachment`. ([#769](https://github.com/terraform-providers/terraform-provider-google-beta/pull/769))
* compute: `google_compute_router_nat` now supports specifying a `log_config` block. ([#743](https://github.com/terraform-providers/terraform-provider-google-beta/pull/743))
* compute: `google_compute_router_nat` now supports more import formats. ([#785](https://github.com/terraform-providers/terraform-provider-google-beta/pull/785))
* compute: `google_compute_network_peering` now supports importing/exporting custom routes ([#754](https://github.com/terraform-providers/terraform-provider-google-beta/pull/754))
* compute: `google_compute_backend_service` now supports self-managed internal load balancing ([#772](https://github.com/terraform-providers/terraform-provider-google-beta/issues/772))
* compute: `google_compute_region_backend_service` now supports failover policies  ([#789](https://github.com/terraform-providers/terraform-provider-google-beta/pull/789))
* compute: Add support for INTERNAL_SELF_MANAGED backend service. Changed Resources: `google_compute_backend_service`, `google_compute_global_forwarding_rule`. ([#772](https://github.com/terraform-providers/terraform-provider-google-beta/pull/772))
* composer: Make cloud composer environment image version updateable ([#741](https://github.com/terraform-providers/terraform-provider-google-beta/pull/741))
* container: `google_container_cluster` now supports `vertical_pod_autoscaling` ([#733](https://github.com/terraform-providers/terraform-provider-google-beta/pull/733))
* container: Expose the `services_ipv4_cidr` for `container_cluster`. ([#804](https://github.com/terraform-providers/terraform-provider-google-beta/pull/804))
* dataflow: `google_dataflow_job` now supports setting machine type ([#1862](https://github.com/GoogleCloudPlatform/magic-modules/pull/1862))
* dns: `google_dns_managed_zone` now supports DNSSec ([#737](https://github.com/terraform-providers/terraform-provider-google-beta/pull/737))
* kms: `google_kms_key_ring` is now autogenerated. ([#748](https://github.com/terraform-providers/terraform-provider-google-beta/pull/748))
* pubsub: `google_pubsub_subscription` supports setting an `expiration_policy` with no `ttl`. ([#783](https://github.com/terraform-providers/terraform-provider-google-beta/pull/783))

BUG FIXES:
* binauth: `google_binary_authorization_policy` can be used with attestors in another project. ([#778](https://github.com/terraform-providers/terraform-provider-google-beta/pull/778))
* compute: allow setting firewall priority to 0 ([#755](https://github.com/terraform-providers/terraform-provider-google-beta/pull/755))
* compute: Resolved an issue where `google_compute_region_backend_service` was unable to perform a state migration. ([#775](https://github.com/terraform-providers/terraform-provider-google-beta/pull/775))
* compute: allow empty metadata.startup-script on instances ([#776](https://github.com/terraform-providers/terraform-provider-google-beta/pull/776))
* compute: Fix flattened custom patchable resources in `google_compute_network`. ([#782](https://github.com/terraform-providers/terraform-provider-google-beta/pull/782))
* compute: `google_compute_vpn_tunnel` now supports sending an empty external gateway interface id. ([#759](https://github.com/terraform-providers/terraform-provider-google-beta/pull/759))
* container: allow AUTH_NONE in istio addon_config ([#664](https://github.com/terraform-providers/terraform-provider-google-beta/pull/664))
* container: allow going from no ip_allocation_policy to a blank-equivalent one ([#774](https://github.com/terraform-providers/terraform-provider-google-beta/pull/774))
* container: `google_container_cluster` will no longer diff unnecessarily on `issue_client_certificate`. ([#788](https://github.com/terraform-providers/terraform-provider-google-beta/pull/788))
* container: `google_container_cluster` can enable client certificates on GKE `1.12+` series releases. ([#788](https://github.com/terraform-providers/terraform-provider-google-beta/pull/788))
* container: `google_container_cluster` now retries the call to remove default node pools during cluster creation ([#799](https://github.com/terraform-providers/terraform-provider-google-beta/pull/799))
* storage: Fix occasional crash when updating storage buckets ([#706](https://github.com/terraform-providers/terraform-provider-google-beta/pull/706))

## 2.7.0 (May 21, 2019)

NOTE:
* Several resources were previously undocumented on the site or changelog; they should be added to both with this release. `google_compute_backend_bucket_signed_url_key` and `google_compute_backend_service_signed_url_key` were introduced in `2.4.0`.

BACKWARDS INCOMPATIBILITIES:
* cloudfunctions: `google_cloudfunctions_function.runtime` now has an explicit default value of `nodejs6`. Users who have a different value set in the API but the value undefined in their config will see a diff. ([#697](https://github.com/terraform-providers/terraform-provider-google-beta/issues/697))

FEATURES:
* **New Resources**: `google_compute_instance_iam_binding`, `google_compute_instance_iam_member`, and `google_compute_instance_iam_policy` are now available. ([#685](https://github.com/terraform-providers/terraform-provider-google-beta/pull/685))
* **New Resources**: IAM resources for Dataproc jobs and clusters (`google_dataproc_job_iam_policy`, `google_dataproc_job_iam_member`, `google_dataproc_job_iam_binding`, `google_dataproc_cluster_iam_policy`, `google_dataproc_cluster_iam_member`, `google_dataproc_cluster_iam_binding`) are now available. [#709](https://github.com/terraform-providers/terraform-provider-google-beta/pull/709)
* **New Resources**: `google_iap_tunnel_instance_iam_binding`, `google_iap_tunnel_instance_iam_member`, and `google_iap_tunnel_instance_iam_policy` are now available. ([#687](https://github.com/terraform-providers/terraform-provider-google-beta/issues/687))

ENHANCEMENTS:
* provider: Add GCP zone to `google_client_config` datasource ([#668](https://github.com/terraform-providers/terraform-provider-google-beta/issues/668))
* compute: Add support for creating instances with CMEK ([#698](https://github.com/terraform-providers/terraform-provider-google-beta/issues/698))
* compute: Can now specify project when importing instance groups.
* compute: `google_compute_instance` now supports `shielded_instance_config` for verifiable integrity of your VM instances. ([#711](https://github.com/terraform-providers/terraform-provider-google-beta/issues/711))
* compute: `google_compute_backend_service` now supports `HTTP2` protocol (beta API feature) [#708](https://github.com/terraform-providers/terraform-provider-google-beta/pull/708)
* compute: `google_compute_instance_template` now supports `shielded_instance_config` for verifiable integrity of your VM instances. ([#711](https://github.com/terraform-providers/terraform-provider-google-beta/issues/711))
* container: use the cluster subnet to look up the node cidr block ([#722](https://github.com/terraform-providers/terraform-provider-google-beta/issues/722))

BUG FIXES:
* cloudfunctions: `google_cloudfunctions_function.runtime` now has an explicit default value of `nodejs6`. ([#697](https://github.com/terraform-providers/terraform-provider-google-beta/issues/697))
* monitoring: updating `google_monitoring_alert_policy` is more likely to succeed ([#684](https://github.com/terraform-providers/terraform-provider-google-beta/issues/684))
* kms: `google_kms_crypto_key` now (in addition to marking all crypto key versions for destruction) correctly disables auto-rotation for destroyed keys ([#705](https://github.com/terraform-providers/terraform-provider-google-beta/issues/705))
* iam: Increase IAM custom role length validation to match API. ([#728](https://github.com/terraform-providers/terraform-provider-google-beta/issues/728))

## 2.6.0 (May 07, 2019)

KNOWN ISSUES:
* cloudfunctions: `google_cloudfunctions_function`s without a `runtime` set will fail to create due to an upstream API change. You can work around this by setting an explicit `runtime` in `2.X` series releases.

DEPRECATIONS:
* monitoring: `google_monitoring_alert_policy` `labels` was deprecated, as the field was never used and it was typed incorrectly. ([#635](https://github.com/terraform-providers/terraform-provider-google-beta/issues/635))

FEATURES:
* **New Datasource**: `google_compute_node_types` for sole-tenant node types is now available. ([#614](https://github.com/terraform-providers/terraform-provider-google-beta/pull/614))
* **New Resource**: `google_compute_node_group` for sole-tenant nodes is now available. ([#643](https://github.com/terraform-providers/terraform-provider-google-beta/pull/643))
* **New Resource**: `google_compute_node_template` for sole-tenant nodes is now available. ([#614](https://github.com/terraform-providers/terraform-provider-google-beta/pull/614))
* **New Resource**: `google_firestore_index` is now available to configure composite indexes on Firestore. ([#632](https://github.com/terraform-providers/terraform-provider-google-beta/issues/632))
* **New Resource**: `google_logging_metric` is now available to configure Stackdriver logs-based metrics. ([#1702](https://github.com/GoogleCloudPlatform/magic-modules/pull/1702))
* **New Resource**: `google_compute_network_endpoint_group` ([#630](https://github.com/terraform-providers/terraform-provider-google-beta/issues/630))
* **New Resource**: `google_security_scanner_scan_config` is now available for configuring scan runs with Cloud Security Scanner. ([#641](https://github.com/terraform-providers/terraform-provider-google-beta/issues/641))

ENHANCEMENTS:
* compute: `google_compute_subnetwork` now supports `log_config` to configure flow logs' logging behaviour. ([#619](https://github.com/terraform-providers/terraform-provider-google-beta/issues/619))
* container: `google_container_cluster` now supports `database_encryption` to configure etcd encryption. ([#649](https://github.com/terraform-providers/terraform-provider-google-beta/issues/649))
* dataflow: `google_dataflow_job`'s `network` and `subnetwork` can be configured. ([#631](https://github.com/terraform-providers/terraform-provider-google-beta/issues/631))
* monitoring: `google_monitoring_alert_policy` `user_labels` support was added. ([#635](https://github.com/terraform-providers/terraform-provider-google-beta/issues/635))
* compute: `google_compute_region_backend_service` is now generated with Magic Modules, adding configurable timeouts, multiple import formats, `creation_timestamp` output. ([#645](https://github.com/terraform-providers/terraform-provider-google-beta/issues/645))
* compute: `iam_compute_subnetwork` is now GA. ([#656](https://github.com/terraform-providers/terraform-provider-google-beta/issues/656))
* pubsub: `google_pubsub_subscription` now supports setting an `expiration_policy`. ([#1703](https://github.com/GoogleCloudPlatform/magic-modules/pull/1703))

BUG FIXES:
* bigquery: `google_bigquery_table` will work with a larger range of projects id formats. ([#658](https://github.com/terraform-providers/terraform-provider-google-beta/issues/658))
* cloudfunctions: `google_cloudfunctions_fucntion` no longer restricts an outdated list of `region`s ([#659](https://github.com/terraform-providers/terraform-provider-google-beta/issues/659))
* compute: `google_compute_instance` now retries updating metadata when fingerprints are mismatched. ([#583](https://github.com/terraform-providers/terraform-provider-google-beta/issues/583))
* compute: `google_compute_instance` and `google_compute_instance_template` now support node affinities for scheduling on sole tenant nodes [[#663](https://github.com/terraform-providers/terraform-provider-google-beta/issues/663)](https://github.com/terraform-providers/terraform-provider-google-beta/pull/663)
* compute: `google_compute_managed_ssl_certificate` will no longer diff when using an absolute FQDN. ([#591](https://github.com/terraform-providers/terraform-provider-google-beta/issues/591))
* compute: `google_compute_disk` resources using `google-beta` will properly detach users at deletion instead of failing. ([#640](https://github.com/terraform-providers/terraform-provider-google-beta/issues/640))
* compute: `google_compute_subnetwork.secondary_ip_ranges` doesn't cause a diff on out of band changes, allows updating to empty list of ranges. ([#3496](https://github.com/terraform-providers/terraform-provider-google-beta/issues/3496))
* container: `google_container_cluster` setting networks / subnetworks by name works with `location`. ([#634](https://github.com/terraform-providers/terraform-provider-google-beta/issues/634))
* container: `google_container_cluster` removed an overly restrictive validation restricting `node_pool` and `remove_default_node_pool` being specified at the same time. ([#637](https://github.com/terraform-providers/terraform-provider-google-beta/issues/637))
* storage: `data_source_google_storage_bucket_object` now correctly URL encodes the slashes in a file name ([#587](https://github.com/terraform-providers/terraform-provider-google-beta/issues/587))

## 2.5.1 (April 22, 2019)

BUG FIXES:
* compute: `google_compute_backend_service` handles empty/nil `iap` block created by previous providers properly. ([#622](https://github.com/terraform-providers/terraform-provider-google-beta/issues/622))
* compute: `google_compute_backend_service` allows multiple instance types in `backends.group` again. ([#625](https://github.com/terraform-providers/terraform-provider-google-beta/issues/625))
* dns: `google_dns_managed_zone` does not permadiff when visiblity is set to default and returned as empty from API ([#624](https://github.com/terraform-providers/terraform-provider-google-beta/issues/624))
* google_projects: Datasource `google_projects` now handles paginated results from listing projects ([#626](https://github.com/terraform-providers/terraform-provider-google-beta/pull/626))
* google_project_iam: `google_project_iam_policy/member/binding` now attempts to retry for read-only operations as well as retrying read-write operations([#620](https://github.com/terraform-providers/terraform-provider-google-beta/pull/620))
* kms: `google_kms_crypto_key.rotation_period` now can be an empty string to allow for unset behavior in modules ([#627](https://github.com/terraform-providers/terraform-provider-google-beta/pull/627))

## 2.5.0 (April 18, 2019)


KNOWN ISSUES:
* compute: `google_compute_subnetwork` will fail to reorder `secondary_ip_range` values at apply time
* compute: `google_compute_subnetwork`s used with a VPC-native GKE cluster will have a diff if that cluster creates secondary ranges automatically.

BACKWARDS INCOMPATIBILITIES:
* all: This is the first release to use the 0.12 SDK required for Terraform 0.12 support. Some provider behaviour may have changed as a result of changes made by the new SDK version.
* compute: `google_compute_instance_group` will not reconcile instances recreated within the same `terraform apply` due to underlying `0.12` SDK changes in the provider. ([#616](https://github.com/terraform-providers/terraform-provider-google-beta/issues/616))
* compute: `google_compute_subnetwork` will have a diff if `secondary_ip_range` values defined in config don't exactly match real state; if so, they will need to be reconciled. ([#3432](https://github.com/terraform-providers/terraform-provider-google-beta/issues/3432))
* container: `google_container_cluster` will have a diff if `master_authorized_networks.cidr_blocks` defined in config doesn't exactly match the real state; if so, it will need to be reconciled. ([#603](https://github.com/terraform-providers/terraform-provider-google-beta/issues/603))


BUG FIXES:
* container: `google_container_cluster` catch out of band changes to `master_authorized_networks.cidr_blocks`. ([#603](https://github.com/terraform-providers/terraform-provider-google-beta/issues/603))


## 2.4.1 (April 30, 2019)

NOTES: This 2.4.1 release is a bugfix release for 2.4.0. It backports the fixes applied in the 2.5.1 release to the 2.4.0 series.

BUG FIXES:
* compute: `google_compute_backend_service` handles empty/nil `iap` block created by previous providers properly. ([#622](https://github.com/terraform-providers/terraform-provider-google-beta/issues/622))
* compute: `google_compute_backend_service` allows multiple instance types in `backends.group` again. ([#625](https://github.com/terraform-providers/terraform-provider-google-beta/issues/625))
* dns: `google_dns_managed_zone` does not permadiff when visiblity is set to default and returned as empty from API ([#624](https://github.com/terraform-providers/terraform-provider-google-beta/issues/624))

## 2.4.0 (April 15, 2019)

KNOWN ISSUES:

* compute: `google_compute_backend_service` resources created with past provider versions won't work with `2.4.0`. You can pin your provider version or manually delete them and recreate them until this is resolved. (https://github.com/terraform-providers/terraform-provider-google/issues/3441)
* dns: `google_dns_managed_zone.visibility` will cause a diff if set to `public`. Setting it to `""` (defaulting to public) will work around this. (https://github.com/terraform-providers/terraform-provider-google/issues/3435)

BACKWARDS INCOMPATIBILITIES:
* accesscontextmanager: `google_access_context_manager_service_perimeter` `unrestricted_services` field was removed based on a removal in the underlying API. ([#576](https://github.com/terraform-providers/terraform-provider-google-beta/issues/576))

FEATURES:
* **New Resource**: `google_compute_backend_bucket_signed_url_key` is now available. ([#530](https://github.com/terraform-providers/terraform-provider-google-beta/issues/530))
* **New Resource**: `google_compute_backend_service_signed_url_key` is now available. ([#577](https://github.com/terraform-providers/terraform-provider-google-beta/issues/577))
* **New Datasource**: `google_service_account_access_token` is now available. ([#575](https://github.com/terraform-providers/terraform-provider-google-beta/issues/575))

ENHANCEMENTS:
* compute: `google_compute_backend_service` is now generated with Magic Modules, adding configurable timeouts, multiple import formats, `creation_timestamp` output. ([#569](https://github.com/terraform-providers/terraform-provider-google-beta/issues/569))
* compute: `google_compute_backend_service` now supports `load_balancing_scheme` and `cdn_policy.signed_url_cache_max_age_sec`. ([#584](https://github.com/terraform-providers/terraform-provider-google-beta/issues/584))
* compute: `google_compute_network` now supports `delete_default_routes_on_create` to delete pre-created routes at network creation time. ([#592](https://github.com/terraform-providers/terraform-provider-google-beta/issues/592))
* compute: `google_compute_autoscaler` now supports `metric.single_instance_assignment` ([#580](https://github.com/terraform-providers/terraform-provider-google-beta/issues/580))
* dns: `google_dns_policy` now supports `enable_logging`. ([#573](https://github.com/terraform-providers/terraform-provider-google-beta/issues/573))
* dns: `google_dns_managed_zone` now supports `peering_config` to enable DNS Peering. ([#572](https://github.com/terraform-providers/terraform-provider-google-beta/issues/572))

BUG FIXES:
* container: `google_container_cluster` will ignore out of band changes on `node_ipv4_cidr_block`. ([#558](https://github.com/terraform-providers/terraform-provider-google-beta/issues/558))
* container: `google_container_cluster` will now reject config with both `node_pool` and `remove_default_node_pool` defined ([#600](https://github.com/terraform-providers/terraform-provider-google-beta/issues/600))
* container: `google_container_cluster` will allow >20 `cidr_blocks` in `master_authorized_networks_config`. ([#594](https://github.com/terraform-providers/terraform-provider-google-beta/issues/594))
* netblock: `data.google_netblock_ip_ranges.cidr_blocks` will better handle ipv6 input. ([#590](https://github.com/terraform-providers/terraform-provider-google-beta/issues/590))
* sql: `google_sql_database_instance` will retry reads during Terraform refreshes if it hits a rate limit. ([#579](https://github.com/terraform-providers/terraform-provider-google-beta/issues/579))

## 2.3.0 (March 26, 2019)

DEPRECATIONS:
* container: `google_container_cluster` `zone` and `region` fields are deprecated in favour of `location`, `additional_zones` in favour of `node_locations`. ([#461](https://github.com/terraform-providers/terraform-provider-google-beta/issues/461))
* container: `google_container_node_pool` `zone` and `region` fields are deprecated in favour of `location`. ([#461](https://github.com/terraform-providers/terraform-provider-google-beta/issues/461))
* container: `data.google_container_cluster` `zone` and `region` fields are deprecated in favour of `location`. ([#461](https://github.com/terraform-providers/terraform-provider-google-beta/issues/461))
* container: `google_container_engine_versions` `zone` and `region` fields are deprecated in favour of `location`. ([#461](https://github.com/terraform-providers/terraform-provider-google-beta/issues/461))

FEATURES:
* **New Datasource**: `google_*_organization_policy` Adding datasources for folder and project org policy ([#468](https://github.com/terraform-providers/terraform-provider-google-beta/issues/468))

ENHANCEMENTS:
* compute: `google_compute_disk`, `google_compute_region_disk` now support `physical_block_size_bytes` ([#526](https://github.com/terraform-providers/terraform-provider-google-beta/issues/526))
* compute: `google_compute_vpn_tunnel will properly apply labels. ([#541](https://github.com/terraform-providers/terraform-provider-google-beta/issues/541))
* container: `google_container_cluster` adds a unified `location` field for regions and zones, `node_locations` to manage extra zones for multi-zonal clusters and specific zones for regional clusters. ([#461](https://github.com/terraform-providers/terraform-provider-google-beta/issues/461))
* container: `google_container_node_pool` adds a unified `location` field for regions and zones. ([#461](https://github.com/terraform-providers/terraform-provider-google-beta/issues/461))
* container: `data.google_container_cluster` adds a unified `location` field for regions and zones. ([#461](https://github.com/terraform-providers/terraform-provider-google-beta/issues/461))
* container: `google_container_engine_versions` adds a unified `location` field for regions and zones. ([#461](https://github.com/terraform-providers/terraform-provider-google-beta/issues/461))
* dataflow: `google_dataflow_job` has support for custom service accounts with `service_account_email`. ([#527](https://github.com/terraform-providers/terraform-provider-google-beta/issues/527))
* monitoring: `google_monitoring_uptime_check` will properly recreate to perform updates. ([#485](https://github.com/terraform-providers/terraform-provider-google-beta/issues/485))
* resourcemanager: `google_*_organization_policy` Add import support for folder and project organization_policies ([#512](https://github.com/terraform-providers/terraform-provider-google-beta/issues/512))
* sql: `google_sql_ssl_cert` Allow project to be specified at resource level ([#524](https://github.com/terraform-providers/terraform-provider-google-beta/issues/524))
* storage: `google_storage_bucket` avoids calls to the compute api during import ([#529](https://github.com/terraform-providers/terraform-provider-google-beta/issues/529))
* storage: `google_storage_bucket.storage_class` supports updating. ([#548](https://github.com/terraform-providers/terraform-provider-google-beta/issues/548))
* various: Some import formats that previously failed will now work as documented. ([#542](https://github.com/terraform-providers/terraform-provider-google-beta/issues/542))

BUG FIXES:
* compute: `google_compute_disk` will properly detach instances again. ([#538](https://github.com/terraform-providers/terraform-provider-google-beta/issues/538))
* container: `google_container_cluster`, `google_container_node_pool` properly suppress new GKE `1.12` `metadata` values. ([#522](https://github.com/terraform-providers/terraform-provider-google-beta/issues/522))
* various: Only 409 concurrent operation errors will be retried, and naming conflicts will not. ([#544](https://github.com/terraform-providers/terraform-provider-google-beta/issues/544))

## 2.2.0 (March 12, 2019)

KNOWN ISSUES:

* compute: `google_compute_disk` is unable to detach instances at deletion time.

---

FEATURES:
* **New Datasource**: `data.google_projects` for retrieving a list of projects based on a filter. ([#493](https://github.com/terraform-providers/terraform-provider-google-beta/issues/493))
* **New Resource**: `google_tpu_node` for Cloud TPU Nodes ([#494](https://github.com/terraform-providers/terraform-provider-google-beta/issues/494))
* **New Resource**: `google_dns_policy` for Cloud DNS policies. ([#488](https://github.com/terraform-providers/terraform-provider-google-beta/pull/488))

ENHANCEMENTS:
* compute: `google_compute_disk` and `google_compute_region_disk` will now detach themselves from a more up to date set of users at delete time. ([#480](https://github.com/terraform-providers/terraform-provider-google-beta/issues/480))
* compute: `google_compute_network` is now generated by Magic Modules, supporting configurable timeouts and more import formats. ([#509](https://github.com/terraform-providers/terraform-provider-google-beta/issues/509))
* compute: `google_compute_firewall` will validate the maximum size of service account lists at plan time. ([#508](https://github.com/terraform-providers/terraform-provider-google-beta/issues/508))
* container: `google_container_cluster` can now disable VPC Native clusters with `ip_allocation_policy.use_ip_aliases` ([#489](https://github.com/terraform-providers/terraform-provider-google-beta/issues/489))
* container: `data.google_container_engine_versions` supports `version_prefix` to allow fuzzy version matching. Using this field, Terraform can match the latest version of a major, minor, or patch release. ([#506](https://github.com/terraform-providers/terraform-provider-google-beta/issues/506))
* pubsub: `google_pubsub_subscription` now supports configuring `message_retention_duration` and `retain_acked_messages`. ([#503](https://github.com/terraform-providers/terraform-provider-google-beta/issues/503))

BUG FIXES:
* app_engine: `google_app_engine_application` correctly outputs `gcr_domain`.  ([#479](https://github.com/terraform-providers/terraform-provider-google-beta/issues/479))
* compute: `data.google_compute_subnetwork` outputs the `self_link` field again. ([#481](https://github.com/terraform-providers/terraform-provider-google-beta/issues/481))
* compute: `google_compute_attached_disk` is now removed from state if the instance was removed. ([#497](https://github.com/terraform-providers/terraform-provider-google-beta/issues/497))
* container: `google_container_cluster` private_cluster_config now has a diff suppress to prevent a permadiff for and allows for empty `master_ipv4_cidr_block`  ([#460](https://github.com/terraform-providers/terraform-provider-google-beta/issues/460))
* container: `google_container_cluster` import behavior fixed/documented for TF-state-only fields (`remove_default_node_pool`, `min_master_version`) ([#476](https://github.com/terraform-providers/terraform-provider-google-beta/issues/476)][[#487](https://github.com/terraform-providers/terraform-provider-google-beta/issues/487)][[#495](https://github.com/terraform-providers/terraform-provider-google-beta/issues/495))
* storagetransfer: `google_storage_transfer_job` will no longer crash when accessing nil dates. ([#499](https://github.com/terraform-providers/terraform-provider-google-beta/issues/499))

## 2.1.0 (February 26, 2019)

FEATURES:
* **New Resource**: Add support for `google_compute_managed_ssl_certificate`.  ([#458](https://github.com/terraform-providers/terraform-provider-google-beta/issues/458))
* **New Datasource**: `google_client_openid_userinfo` for retrieving the `email` used to authenticate with GCP. ([#459](https://github.com/terraform-providers/terraform-provider-google-beta/issues/459))

ENHANCEMENTS:
* compute: `data.google_compute_subnetwork` can now be addressed by `self_link` as an alternative to the existing `name`/`region`/`project` fields. ([#429](https://github.com/terraform-providers/terraform-provider-google-beta/issues/429))
* dns: Support for privately visible zones is added to `google_dns_managed_zone`. ([#268](https://github.com/terraform-providers/terraform-provider-google-beta/issues/268))
* pubsub: `google_pubsub_topic` is now generated using Magic Modules, adding Open in Cloud Shell examples, configurable timeouts, and the `labels` field. ([#432](https://github.com/terraform-providers/terraform-provider-google-beta/issues/432))
* pubsub: `google_pubsub_subscription` is now generated using Magic Modules, adding Open in Cloud Shell examples, configurable timeouts, update support, and the `labels` field. ([#432](https://github.com/terraform-providers/terraform-provider-google-beta/issues/432))
* sql: `google_sql_database_instance` now provides `public_ip_address` and `private_ip_address` outputs of the first public and private IP of the instance respectively. ([#454](https://github.com/terraform-providers/terraform-provider-google-beta/issues/454))


BUG FIXES:
* sql: `google_sql_database_instance` allows the empty string to be set for `private_network`. ([#454](https://github.com/terraform-providers/terraform-provider-google-beta/issues/454))

## 2.0.0 (February 12, 2019)

BACKWARDS INCOMPATIBILITIES:
* bigtable: `google_bigtable_instance` `zone` field is no longer inferred from the provider.
* bigtable: `google_bigtable_table` now reads `family` from the table's column family in Cloud Bigtable instead of creating a new column family ([#70](https://github.com/terraform-providers/terraform-provider-google-beta/issues/70))
* bigtable: `google_bigtable_instance.cluster.num_nodes` will fail at plan time if `DEVELOPMENT` instances have `num_nodes = "0"` set explicitly. If it has been set, unset the field. ([#82](https://github.com/terraform-providers/terraform-provider-google-beta/issues/82))
* cloudbuild: `google_cloudbuild_trigger.build.step.args` is now a list instead of space separated strings. ([#308](https://github.com/terraform-providers/terraform-provider-google-beta/issues/308))
* cloudfunctions: `google_cloudfunctions_function.retry_on_failure` has been removed. Use `event_trigger.failure_policy.retry` instead. ([#75](https://github.com/terraform-providers/terraform-provider-google-beta/issues/75))
* cloudfunctions: `google_cloudfunctions_function.trigger_bucket` and `google_cloudfunctions_function.trigger_topic` have been removed. Use `event trigger` instead. ([#30](https://github.com/terraform-providers/terraform-provider-google-beta/issues/30))
* composer: `google_composer_environment.node_config.zone` is now `Required`. ([#396](https://github.com/terraform-providers/terraform-provider-google-beta/issues/396))
* compute: `google_compute_instance`, `google_compute_instance_from_template` `metadata` field is now authoritative and will remove values not explicitly set in config. [[#2208](https://github.com/terraform-providers/terraform-provider-google-beta/issues/2208)](https://github.com/terraform-providers/terraform-provider-google/pull/2208)
* compute: `google_compute_region_instance_group_manager` field `update_strategy` is now deprecated in the beta provider only. It will only function in the `google` provider, ([#76](https://github.com/terraform-providers/terraform-provider-google-beta/issues/76))
* compute: `google_compute_global_forwarding_rule` field `labels` is now removed ([#81](https://github.com/terraform-providers/terraform-provider-google-beta/issues/81))
* compute: `google_compute_project_metadata` resource is now authoritative and will remove values not explicitly set in config. [[#2205](https://github.com/terraform-providers/terraform-provider-google-beta/issues/2205)](https://github.com/terraform-providers/terraform-provider-google/pull/2205)
* compute: `google_compute_url_map` resource is now authoritative and will remove values not explicitly set in config. [[#2245](https://github.com/terraform-providers/terraform-provider-google-beta/issues/2245)](https://github.com/terraform-providers/terraform-provider-google/pull/2245)
* compute: `google_compute_snapshot.snapshot_encryption_key_raw`, `google_compute_snapshot.snapshot_encryption_key_sha256`, `google_compute_snapshot.source_disk_encryption_key_raw`, `google_compute_snapshot.source_disk_encryption_key_sha256` fields are now removed. Use `google_compute_snapshot.snapshot_encryption_key.0.raw_key`, `google_compute_snapshot.snapshot_encryption_key.0.sha256`, `google_compute_snapshot.source_disk_encryption_key.0.raw_key`, `google_compute_snapshot.source_disk_encryption_key.0.sha256` instead. ([#202](https://github.com/terraform-providers/terraform-provider-google-beta/issues/202))
* compute: `google_compute_instance_group_manager` is no longer imported by the provider-level region. Set the appropriate provider-level zone instead. ([#248](https://github.com/terraform-providers/terraform-provider-google-beta/issues/248))
* compute: `google_compute_region_instance_group_manager.update_strategy` in the `google-beta` provider has been removed. ([#189](https://github.com/terraform-providers/terraform-provider-google-beta/issues/189))
* compute: `google_compute_instance`, `google_compute_instance_template`, `google_compute_instance_from_template` have had the `network_interface.address` field removed. ([#190](https://github.com/terraform-providers/terraform-provider-google-beta/issues/190))
* compute: `google_compute_instance` has had the `network_interface.access_config.assigned_nat_ip` field removed ([#48](https://github.com/terraform-providers/terraform-provider-google-beta/issues/48))
* compute: `google_compute_disk` is no longer imported by the provider-level region. Set the appropriate provider-level zone instead. ([#249](https://github.com/terraform-providers/terraform-provider-google-beta/issues/249))
* compute: `google_compute_router_nat.subnetwork.source_ip_ranges_to_nat` is now Required inside `subnetwork` blocks. ([#281](https://github.com/terraform-providers/terraform-provider-google-beta/issues/281))
* compute: `google_compute_ssl_certificate`'s `private_key` field is no longer stored in state in cleartext; it is now SHA256 encoded. ([#400](https://github.com/terraform-providers/terraform-provider-google-beta/issues/400))
* container: `google_container_cluster` fields (`private_cluster`, `master_ipv4_cidr_block`) are removed. Use `private_cluster_config` and `private_cluster_config.master_ipv4_cidr_block` instead. ([#78](https://github.com/terraform-providers/terraform-provider-google-beta/issues/78))
* container: `google_container_node_pool`'s `name_prefix` field has been restored and is no longer deprecated. ([#2975](https://github.com/terraform-providers/terraform-provider-google-beta/issues/2975))
* sql: `google_sql_database_instance` resource is now authoritative and will remove values not explicitly set in config. [[#2203](https://github.com/terraform-providers/terraform-provider-google-beta/issues/2203)](https://github.com/terraform-providers/terraform-provider-google/pull/2203)
* endpoints: `google_endpoints_service.protoc_output` was removed. Use `google_endpoints_service.protoc_output_base64` instead. ([#79](https://github.com/terraform-providers/terraform-provider-google-beta/issues/79))
* resourcemanager: `google_project_iam_policy` is now authoritative and will remove values not explicitly set in config. Several fields were removed that made it authoritative: `authoritative`, `restore_policy`, and `disable_project`. This resource is very dangerous! Ensure you are not using the removed fields (`authoritative`, `restore_policy`, `disable_project`). ([#25](https://github.com/terraform-providers/terraform-provider-google-beta/issues/25))
* resourcemanager: Datasource `google_service_account_key.service_account_id` has been removed. Use the `name` field instead. ([#80](https://github.com/terraform-providers/terraform-provider-google-beta/issues/80))
* resourcemanager: `google_project.app_engine` has been removed. Use the `google_app_engine_application` resource instead. ([#74](https://github.com/terraform-providers/terraform-provider-google-beta/issues/74))
* resourcemanager: `google_organization_custom_role.deleted` is now an output-only attribute. Use `terraform destroy`, or remove the resource from your config instead. ([#191](https://github.com/terraform-providers/terraform-provider-google-beta/issues/191))
* resourcemanager: `google_project_custom_role.deleted` is now an output-only attribute. Use `terraform destroy`, or remove the resource from your config instead. ([#199](https://github.com/terraform-providers/terraform-provider-google-beta/issues/199))
* serviceusage: `google_project_service` will now error instead of silently disabling dependent services if `disable_dependent_services` is unset. ([#384](https://github.com/terraform-providers/terraform-provider-google-beta/issues/384))
* storage: `google_storage_object_acl.role_entity` is now authoritative and will remove values not explicitly set in config. Use `google_storage_object_access_control` for fine-grained management. ([#26](https://github.com/terraform-providers/terraform-provider-google-beta/issues/26))
* storage: `google_storage_default_object_acl.role_entity` is now authoritative and will remove values not explicitly set in config. ([#47](https://github.com/terraform-providers/terraform-provider-google-beta/issues/47))
* iam: `google_*_iam_binding` Change all IAM bindings to be authoritative ([#291](https://github.com/terraform-providers/terraform-provider-google-beta/issues/291))

FEATURES:
* **New Resource**: `google_access_context_manager_access_policy` for managing the container for an organization's access levels. ([#96](https://github.com/terraform-providers/terraform-provider-google-beta/issues/96))
* **New Resource**: `google_access_context_manager_access_level` for managing an organization's access levels. ([#149](https://github.com/terraform-providers/terraform-provider-google-beta/issues/149))
* **New Resource**: `google_access_context_manager_service_perimeter` for managing service perimeters in an access policy. ([#246](https://github.com/terraform-providers/terraform-provider-google-beta/issues/246))
* **New Resource**: `google_app_engine_firewall_rule` ([#271](https://github.com/terraform-providers/terraform-provider-google-beta/issues/271)][[#336](https://github.com/terraform-providers/terraform-provider-google-beta/issues/336))
* **New Resource**: `google_monitoring_group` ([#120](https://github.com/terraform-providers/terraform-provider-google-beta/issues/120))
* **New Resource**: `google_project_iam_audit_config` ([#265](https://github.com/terraform-providers/terraform-provider-google-beta/issues/265))
* **New Resource**: `google_storage_transfer_job` for managing recurring storage transfers with Google Cloud Storage. ([#256](https://github.com/terraform-providers/terraform-provider-google-beta/issues/256))
* **New Resource**: `google_cloud_scheduler_job` for managing the cron job scheduling service with Google Cloud Scheduler. ([#378](https://github.com/terraform-providers/terraform-provider-google-beta/issues/378))
* **New Datasource**: `google_storage_bucket_object` ([#223](https://github.com/terraform-providers/terraform-provider-google-beta/issues/223))
* **New Datasource**: `google_storage_transfer_project_service_account` data source for retrieving the Storage Transfer service account for a project ([#247](https://github.com/terraform-providers/terraform-provider-google-beta/issues/247))
* **New Datasource**: `google_kms_crypto_key` ([#359](https://github.com/terraform-providers/terraform-provider-google-beta/issues/359))
* **New Datasource**: `google_kms_key_ring` ([#359](https://github.com/terraform-providers/terraform-provider-google-beta/issues/359))

ENHANCEMENTS:
* provider: Add `access_token` config option to allow Terraform to authenticate using short-lived Google OAuth 2.0 access token ([#330](https://github.com/terraform-providers/terraform-provider-google-beta/issues/330))
* bigquery: Add new locations `europe-west2` and `australia-southeast1` to valid location set for `google_bigquery_dataset` ([#41](https://github.com/terraform-providers/terraform-provider-google-beta/issues/41))
* bigquery: Add `default_partition_expiration_ms` field to `google_bigquery_dataset` resource. ([#127](https://github.com/terraform-providers/terraform-provider-google-beta/issues/127))
* bigquery: Add `delete_contents_on_destroy` field to `google_bigquery_dataset` resource. ([#413](https://github.com/terraform-providers/terraform-provider-google-beta/issues/413))
* bigquery: Add `time_partitioning.require_partition_filter` to `google_bigquery_table` resource. ([#324](https://github.com/terraform-providers/terraform-provider-google-beta/issues/324))
* bigquery: Allow more BigQuery regions ([#269](https://github.com/terraform-providers/terraform-provider-google-beta/issues/269))
* bigtable: Add `column_family` at create time to `google_bigtable_table`. [[#2228](https://github.com/terraform-providers/terraform-provider-google-beta/issues/2228)](https://github.com/terraform
