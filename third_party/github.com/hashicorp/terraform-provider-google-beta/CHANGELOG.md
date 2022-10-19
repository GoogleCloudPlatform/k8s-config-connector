## 4.41.0 (Unreleased)
BREAKING CHANGES:
* sql: updated `google_sql_user.sql_server_user_details` to be read only. Any configuration attempting to set this field is invalid and will cause the provider to fail during plan time. ([#4764](https://github.com/hashicorp/terraform-provider-google-beta/pull/4764))

FEATURES:
* **New Resource:**  `google_cloud_ids_endpoint` ([#4765](https://github.com/hashicorp/terraform-provider-google-beta/pull/4765))
* **New Resource:** `google_bigquery_analytics_hub_listing_iam_binding` ([#4771](https://github.com/hashicorp/terraform-provider-google-beta/pull/4771))
* **New Resource:** `google_bigquery_analytics_hub_listing_iam_member` ([#4771](https://github.com/hashicorp/terraform-provider-google-beta/pull/4771))
* **New Resource:** `google_bigquery_analytics_hub_listing_iam_policy` ([#4771](https://github.com/hashicorp/terraform-provider-google-beta/pull/4771))
* **New Resource:** `google_bigquery_analytics_hub_listing` ([#4771](https://github.com/hashicorp/terraform-provider-google-beta/pull/4771))

IMPROVEMENTS:
* appengine: added support for `service_account` field to `google_app_engine_standard_app_version` resource ([#4757](https://github.com/hashicorp/terraform-provider-google-beta/pull/4757))
* bigquery: added `avro_options` field to `google_bigquery_table` resource ([#4768](https://github.com/hashicorp/terraform-provider-google-beta/pull/4768))
* cloudrun: added field `startup_probe` to resource `google_cloud_run_service` ([#4773](https://github.com/hashicorp/terraform-provider-google-beta/pull/4773))
* compute: added `node_config.0.guest_accelerator.0.gpu_sharing_config` field to `google_container_node_pool` resource ([#4758](https://github.com/hashicorp/terraform-provider-google-beta/pull/4758))
* datafusion: added `crypto_key_config` field to `google_data_fusion_instance` resource ([#4761](https://github.com/hashicorp/terraform-provider-google-beta/pull/4761))
* filestore: removed constraint that forced multiple `google_filestore_instance` creations to occur serially ([#4770](https://github.com/hashicorp/terraform-provider-google-beta/pull/4770))

BUG FIXES:
* kms: fixed apply failure when `google_kms_crypto_key` is removed after its versions were destroyed earlier ([#4769](https://github.com/hashicorp/terraform-provider-google-beta/pull/4769))
* monitoring: fixed a bug causing a perma-diff in `google_monitoring_alert_policy` when `cross_series_reducer` was set to "REDUCE_NONE" ([#4763](https://github.com/hashicorp/terraform-provider-google-beta/pull/4763))

## 4.40.0 (October 10, 2022)

FEATURES:
* **New Data Source:** `google_cloudfunctions2_function` ([#4732](https://github.com/hashicorp/terraform-provider-google-beta/pull/4732))
* **New Data Source:** `google_compute_snapshot` ([#4731](https://github.com/hashicorp/terraform-provider-google-beta/pull/4731))
* **New Resource:** `google_compute_region_target_tcp_proxy` ([#4749](https://github.com/hashicorp/terraform-provider-google-beta/pull/4749))
* **New Resource:** `google_identity_platform_config` ([#4729](https://github.com/hashicorp/terraform-provider-google-beta/pull/4729))
* **New Resource:** `google_bigquery_datapolicy_data_policy` ([#4754](https://github.com/hashicorp/terraform-provider-google-beta/pull/4754))
* **New Resource:** `google_bigquery_datapolicy_data_policy_iam_binding` ([#4754](https://github.com/hashicorp/terraform-provider-google-beta/pull/4754))
* **New Resource:** `google_bigquery_datapolicy_data_policy_iam_member` ([#4754](https://github.com/hashicorp/terraform-provider-google-beta/pull/4754))
* **New Resource:** `google_bigquery_datapolicy_data_policy_iam_policy` ([#4754](https://github.com/hashicorp/terraform-provider-google-beta/pull/4754))
* **New Resource:** `google_org_policy_custom_constraint` ([#4741](https://github.com/hashicorp/terraform-provider-google-beta/pull/4741))
* **New Resource:** `google_vertex_ai_featurestore_entitytype_feature` ([#4736](https://github.com/hashicorp/terraform-provider-google-beta/pull/4736))

IMPROVEMENTS:
* bigqueryreservation: added `concurrency` and `multiRegionAuxiliary` to `google_bigquery_reservation` ([#4739](https://github.com/hashicorp/terraform-provider-google-beta/pull/4739))
* bigtable: added additional retry GC policy operations with a longer poll interval to avoid quota issues ([#4750](https://github.com/hashicorp/terraform-provider-google-beta/pull/4750))
* bigtable: improved error messaging ([#4746](https://github.com/hashicorp/terraform-provider-google-beta/pull/4746))
* compute: added support for `compression_mode` field in `google_compute_backend_bucket` and `google_compute_backend_service` resource ([#4733](https://github.com/hashicorp/terraform-provider-google-beta/pull/4733))
* dataflow : added support of `labels` to resource `google_dataflow_flextemplate_job` ([#4748](https://github.com/hashicorp/terraform-provider-google-beta/pull/4748))
* datastream: added field `bigquery_profile` to `google_datastream_connection_profile` ([#4742](https://github.com/hashicorp/terraform-provider-google-beta/pull/4742))
* dns: added general field `cloud_logging_config` to `google_dns_managed_zone` ([#4734](https://github.com/hashicorp/terraform-provider-google-beta/pull/4734))
* metastore: added bigquery support for `google_dataproc_metastore_service` ([#4753](https://github.com/hashicorp/terraform-provider-google-beta/pull/4753))
* storage: added `custom_placement_config` field to `google_storage_bucket` resource to support custom dual-region GCS buckets ([#4752](https://github.com/hashicorp/terraform-provider-google-beta/pull/4752))
* sql: added  `password_policy` field to `google_sql_user` resource ([#4730](https://github.com/hashicorp/terraform-provider-google-beta/pull/4730))

BUG FIXES:
* storage: fixed a bug where user specified labels get overwritten by Dataplex auto generated labels ([#4743](https://github.com/hashicorp/terraform-provider-google-beta/pull/4743))
* storagetransfer: fixed a crash in `google_storagetransfer_job` refreshes when `transfer_schedule` was empty ([#4745](https://github.com/hashicorp/terraform-provider-google-beta/pull/4745))


## 4.39.0 (October 3, 2022)

FEATURES:
* **New Data Source:** `google_artifact_registry_repository` ([#4714](https://github.com/hashicorp/terraform-provider-google-beta/pull/4714))
* **New Resource:** `google_identity_platform_config` ([#4729](https://github.com/hashicorp/terraform-provider-google-beta/pull/4729))

IMPROVEMENTS:
* certificatemanager: added public/private PEM fields `pem_certificate` / `pem_private_key` and deprecated `certificate_pem` / `private_key_pem` ([#4728](https://github.com/hashicorp/terraform-provider-google-beta/pull/4728))
* clouddeploy: added `serial_pipeline.stages.strategy` field to `google_clouddeploy_delivery_pipeline` ([#4707](https://github.com/hashicorp/terraform-provider-google-beta/pull/4707))
* container: added `notification_config.pubsub.filter` field to `google_container_cluster` ([#4718](https://github.com/hashicorp/terraform-provider-google-beta/pull/4718))
* eventarc: added `channels` and `conditions` fields to `google_eventarc_trigger` ([#4707](https://github.com/hashicorp/terraform-provider-google-beta/pull/4707))
* healthcare: added `notification_configs ` field to `google_healthcare_fhir_store` resource ([#4720](https://github.com/hashicorp/terraform-provider-google-beta/pull/4720))
* iap: added ability to import `google_iap_brand` using ID using {{project}}/{{brand_id}} format ([#4712](https://github.com/hashicorp/terraform-provider-google-beta/pull/4712))
* secretmanager: added output field 'version' to resource 'secret_manager_secret_version' ([#4724](https://github.com/hashicorp/terraform-provider-google-beta/pull/4724))
* sql: added `maintenance_version` and `available_maintenance_versions` fields to `google_sql_database_instance` resource ([#4725](https://github.com/hashicorp/terraform-provider-google-beta/pull/4725))
* storagetransfer: added `notification_config` field to `google_storage_transfer_job` resource ([#4709](https://github.com/hashicorp/terraform-provider-google-beta/pull/4709))
* tags: added `purpose` and `purpose_data` properties to `google_tags_tag_key` ([#4721](https://github.com/hashicorp/terraform-provider-google-beta/pull/4721))

BUG FIXES:
* bigquery: fixed a bug where `allow_quoted_newlines` and `allow_jagged_rows` could not be set to false on `google_bigquery_table` ([#4711](https://github.com/hashicorp/terraform-provider-google-beta/pull/4711))
* cloudfunction: fixed inability to update `docker_repository` and `kms_key_name` on `google_cloudfunctions_function` ([#4727](https://github.com/hashicorp/terraform-provider-google-beta/pull/4727))
* compute: fixed inability to manage Cloud Armor `adaptive_protection_config` on `google_compute_security_policy` ([#4726](https://github.com/hashicorp/terraform-provider-google-beta/pull/4726))
* container: fixed a bug where upgrading provider version breaks on `node_pool_auto_config` or `node_pool_defaults` ([#4706](https://github.com/hashicorp/terraform-provider-google-beta/pull/4706))
* iam: fixed diffs between `policy_data` from `google_iam_policy` data source and policy data in API responses ([#4722](https://github.com/hashicorp/terraform-provider-google-beta/pull/4722))
* iam: fixed permadiff resulting from empty fields being sent in requests to set conditional IAM policies ([#4723](https://github.com/hashicorp/terraform-provider-google-beta/pull/4723))
* secretmanager: fixed a bug where `google_secret_manager_secret_version` that was destroyed outside of Terraform would not be recreated on apply ([#4719](https://github.com/hashicorp/terraform-provider-google-beta/pull/4719))
* storagetransfer: fixed a crash in `google_storagetransfer_job` when `transfer_schedule` is empty ([#4745](https://github.com/hashicorp/terraform-provider-google-beta/pull/4745))
 
## 4.38.0 (September 26, 2022)

FEATURES:
* **New Data Source:** `google_vpc_access_connector` ([#4693](https://github.com/hashicorp/terraform-provider-google-beta/pull/4693))
* **New Resource:** `google_datastream_private_connection` ([#4691](https://github.com/hashicorp/terraform-provider-google-beta/pull/4691))

IMPROVEMENTS:
* appengine: Added `egress_setting` for field `vpc_access_connector` to `google_app_engine_standard_app_version` ([#4701](https://github.com/hashicorp/terraform-provider-google-beta/pull/4701))
* bigquery: added `json_extension` field to the `load` block of `google_bigquery_job` resource ([#4699](https://github.com/hashicorp/terraform-provider-google-beta/pull/4699))
* cloudfunctions: Added `build_worker_pool` to `google_cloudfunctions_function` ([#4696](https://github.com/hashicorp/terraform-provider-google-beta/pull/4696))
* compute: added `json_custom_config` field to `google_compute_security_policy` resource ([#4703](https://github.com/hashicorp/terraform-provider-google-beta/pull/4703))
* redis: Added `persistence_config` field to the `google_redis_instance` resource. ([#4688](https://github.com/hashicorp/terraform-provider-google-beta/pull/4688))
* storage: added support for `overwriteWhen` field to `transfer_options` in `google_storage_transfer_job` resource ([#4690](https://github.com/hashicorp/terraform-provider-google-beta/pull/4690))

BUG FIXES:
* bigtable: added drift detection on `gc_rules` for `google_bigtable_gc_policy` ([#4687](https://github.com/hashicorp/terraform-provider-google-beta/pull/4687))
* compute: fixed the inability to update `most_disruptive_allowed_action` for both `google_compute_per_instance_config` and `google_compute_region_per_instance_config` ([#4685](https://github.com/hashicorp/terraform-provider-google-beta/pull/4685))
* container: fixed allow passing empty list to monitoring_config and logging_config in `google_container_cluster` ([#4700](https://github.com/hashicorp/terraform-provider-google-beta/pull/4700))
* sql: fixed a bug causing a perma-diff on `disk_type` due to API values being downcased ([#4686](https://github.com/hashicorp/terraform-provider-google-beta/pull/4686))
* storage: fixed the inability to set 0 for `lifecycle_rule.condition.age` on `google_storage_bucket` ([#4698](https://github.com/hashicorp/terraform-provider-google-beta/pull/4698))

## 4.37.0 (September 19, 2022)

FEATURES:
* **New Resource:** `google_apigee_nat_address` ([#4676](https://github.com/hashicorp/terraform-provider-google-beta/pull/4676))
* **New Resource:** `google_dialogflow_cx_webhook` ([#4667](https://github.com/hashicorp/terraform-provider-google-beta/pull/4667))
* **New Resource:** `google_filestore_snapshot` ([#4661](https://github.com/hashicorp/terraform-provider-google-beta/pull/4661))

IMPROVEMENTS:
* apigee: added read-only field `connection_state` to `google_apigee_endpoint_attachment` ([#4668](https://github.com/hashicorp/terraform-provider-google-beta/pull/4668))
* bigtable: added support for `autoscaling_config.storage_target` to `google_bigtable_instance` ([#4671](https://github.com/hashicorp/terraform-provider-google-beta/pull/4671))
* cloudbuild: added support for `BITBUCKET` option to `git_source.repo_type` in `google_cloudbuild_trigger` ([#4679](https://github.com/hashicorp/terraform-provider-google-beta/pull/4679))
* dns: added in validation for trailing dot at end of DNS record name ([#4674](https://github.com/hashicorp/terraform-provider-google-beta/pull/4674))
* project: added validation for field `project_id` in `google_project` datasource. ([#4684](https://github.com/hashicorp/terraform-provider-google-beta/pull/4684))
* serviceaccount: added `expires_in` attribute for generating `exp` claim to `google_service_account_jwt` datasource. ([#4677](https://github.com/hashicorp/terraform-provider-google-beta/pull/4677))

BUG FIXES:
* notebooks: fixed perma-diff in `google_notebooks_instance` ([#4664](https://github.com/hashicorp/terraform-provider-google-beta/pull/4664))
* privateca: fixed an issue that blocked subordinate CA data sources when `state` was not `AWAITING_USER_ACTIVATION` ([#4672](https://github.com/hashicorp/terraform-provider-google-beta/pull/4672))
* storage: fixed permdiff on the field `versioning` of `google_storage_bucket` ([#4665](https://github.com/hashicorp/terraform-provider-google-beta/pull/4665))

## 4.36.0 (September 12, 2022)

FEATURES:
* **New Resource:** `google_bigquery_analytics_hub_data_exchange_iam_binding` ([#4656](https://github.com/hashicorp/terraform-provider-google-beta/pull/4656))
* **New Resource:** `google_bigquery_analytics_hub_data_exchange_iam_member` ([#4656](https://github.com/hashicorp/terraform-provider-google-beta/pull/4656))
* **New Resource:** `google_bigquery_analytics_hub_data_exchange_iam_policy` ([#4656](https://github.com/hashicorp/terraform-provider-google-beta/pull/4656))
* **New Resource:** `google_bigquery_analytics_hub_data_exchange` ([#4656](https://github.com/hashicorp/terraform-provider-google-beta/pull/4656))
* **New Resource:** `google_datastream_connection_profile` ([#4657](https://github.com/hashicorp/terraform-provider-google-beta/pull/4657))

IMPROVEMENTS:
* appengine: added field `service_account` to `google_app_engine_flexible_app_version` ([#4653](https://github.com/hashicorp/terraform-provider-google-beta/pull/4653))
* bigtable: increased timeout in `google_bigtable_table` creation. ([#4655](https://github.com/hashicorp/terraform-provider-google-beta/pull/4655))
* cloudbuild: added `location` field to `google_cloudbuild_trigger` resource ([#4646](https://github.com/hashicorp/terraform-provider-google-beta/pull/4646))
* compute: added `certificate_map` to `compute_target_ssl_proxy` resource ([#4654](https://github.com/hashicorp/terraform-provider-google-beta/pull/4654))
* compute: added field `chain_name` to `google_compute_resource_policy.snapshot_properties` ([#4660](https://github.com/hashicorp/terraform-provider-google-beta/pull/4660))
* compute: added field `chain_name` to resource `google_compute_snapshot` ([#4660](https://github.com/hashicorp/terraform-provider-google-beta/pull/4660))
* container: added `autoscaling.total_min_node_count`, `autoscaling.total_max_node_count`, and `autoscaling.location_policy` to `google_container_cluster.node_pool` ([#4649](https://github.com/hashicorp/terraform-provider-google-beta/pull/4649))
* container: added `autoscaling.total_min_node_count`, `autoscaling.total_max_node_count`, and `autoscaling.location_policy` to `google_container_node_pool` resource ([#4649](https://github.com/hashicorp/terraform-provider-google-beta/pull/4649))
* container: added field `node_pool_defaults` to `resource_container_cluster`. ([#4648](https://github.com/hashicorp/terraform-provider-google-beta/pull/4648))
* dataproc: added option `shielded_instance_config` to resource `google_dataproc_workflow_template`. ([#4647](https://github.com/hashicorp/terraform-provider-google-beta/pull/4647))
* metastore: extended default timeouts for `google_dataproc_metastore_service` from 40m to 60m ([#4652](https://github.com/hashicorp/terraform-provider-google-beta/pull/4652))
* pubsub: made `google_pubsub_subscription.enable_exactly_once_delivery` mutable so that it updates subscription without recreation. ([#4645](https://github.com/hashicorp/terraform-provider-google-beta/pull/4645))

## 4.35.0 (September 6, 2022)

IMPROVEMENTS:
* apigee: added support for `nodeConfig` in `google_apigee_environment` ([#4632](https://github.com/hashicorp/terraform-provider-google-beta/pull/4632))
* apigee: added a `properties` field to `google_apigee_organization` ([#4644](https://github.com/hashicorp/terraform-provider-google-beta/pull/4644))
* cloudfunctions2: added `secret_environment_variables` and `secret_volumes` to `google_cloudfunctions2_function` ([#4641](https://github.com/hashicorp/terraform-provider-google-beta/pull/4641))
* compute: added support for param `visible_core_count` in `google_compute_instance` and `google_compute_instance_template` under `advanced_machine_features` ([#4635](https://github.com/hashicorp/terraform-provider-google-beta/pull/4635))
* compute: added support documentation links to error messages for certain Compute Operation errors. ([#4642](https://github.com/hashicorp/terraform-provider-google-beta/pull/4642))
* container: added `service_external_ips_config` support to `cluster_container` resource. ([#4639](https://github.com/hashicorp/terraform-provider-google-beta/pull/4639))
* container: added `enable_cost_allocation` to `google_container_cluster` ([#4640](https://github.com/hashicorp/terraform-provider-google-beta/pull/4640))
* dns: added `behavior` field to `google_dns_response_policy_rule` resource ([#4637](https://github.com/hashicorp/terraform-provider-google-beta/pull/4637))
* monitoring: added `force_delete` field to `google_monitoring_notification_channel` resource ([#4638](https://github.com/hashicorp/terraform-provider-google-beta/pull/4638))
* pubsub: made `enable_exactly_once_delivery` mutable so that it updates subscription in-place and avoids recreation of the subscription. ([#4645](https://github.com/hashicorp/terraform-provider-google-beta/pull/4645))
* vertexai: added `encryption_spec` field to `google_vertex_ai_featurestore` resource (beta) ([#4643](https://github.com/hashicorp/terraform-provider-google-beta/pull/4643))

BUG FIXES:
* compute: fixed the `id` format of the data source `google_compute_instance` ([#4636](https://github.com/hashicorp/terraform-provider-google-beta/pull/4636))

## 4.34.0 (August 29, 2022)
NOTES:
* updated Bigtable go client version from 1.13 to 1.16. ([#4613](https://github.com/hashicorp/terraform-provider-google-beta/pull/4613))

IMPROVEMENTS:
* apigee: added support for specifying retention when deleting `google_apigee_organization` ([#4604](https://github.com/hashicorp/terraform-provider-google-beta/pull/4604))
* appengine: added `app_engine_apis` field to `google_app_engine_standard_app_version` resource ([#4607](https://github.com/hashicorp/terraform-provider-google-beta/pull/4607))
* compute: improved error messaging for compute errors ([#4602](https://github.com/hashicorp/terraform-provider-google-beta/pull/4602))
* container: added general field `reservation_affinity` to `google_container_node_pool` ([#4622](https://github.com/hashicorp/terraform-provider-google-beta/pull/4622))
* container: added field `auto_provisioning_network_tags` to `google_container_cluster` (beta) ([#4611](https://github.com/hashicorp/terraform-provider-google-beta/pull/4611))
* sql: added support for major version upgrade to `google_sql_database_instance ` resource  ([#4606](https://github.com/hashicorp/terraform-provider-google-beta/pull/4606))

BUG FIXES:
* bigtable: fixed comparing column family name when reading a GC policy. ([#4624](https://github.com/hashicorp/terraform-provider-google-beta/pull/4624))
* bigtable: passed `isTopeLevel` in getGCPolicyFromJSON() instead of hardcoding it to true. ([#4615](https://github.com/hashicorp/terraform-provider-google-beta/pull/4615))
* cloud iam: made `denial_condition` optional on `google_iam_deny_policy` ([#4617](https://github.com/hashicorp/terraform-provider-google-beta/pull/4617))

## 4.33.0 (August 22, 2022)
IMPROVEMENTS:
* container: added update support for `authenticator_groups_config` in `google_container_cluster` ([#4591](https://github.com/hashicorp/terraform-provider-google-beta/pull/4591))
* dataflow: added ability to import `google_dataflow_job` ([#4595](https://github.com/hashicorp/terraform-provider-google-beta/pull/4595))
* dns: added `managed_zone_id` attribute to `google_dns_managed_zone` data source ([#4593](https://github.com/hashicorp/terraform-provider-google-beta/pull/4593))
* metastore: added `metadata_integration` and `hive_metastore_config.auxiliary_versions` fields to `google_dataproc_metastore_service` resource ([#4598](https://github.com/hashicorp/terraform-provider-google-beta/pull/4598))
* monitoring: added `accepted_response_status_codes` to `google_monitoring_uptime_check_config` ([#4594](https://github.com/hashicorp/terraform-provider-google-beta/pull/4594))
* sql: added `password_validation_policy` field to `google_cloud_sql` resource ([#4597](https://github.com/hashicorp/terraform-provider-google-beta/pull/4597))

BUG FIXES:
* bigquery: removed force replacement for `display_name` on `google_bigquery_data_transfer_config` ([#4592](https://github.com/hashicorp/terraform-provider-google-beta/pull/4592))
* compute: fixed permadiff for `instance_termination_action` in `google_compute_instance_template` ([#4590](https://github.com/hashicorp/terraform-provider-google-beta/pull/4590))

## 4.32.0 (August 15, 2022)

NOTES:
* Updated to Golang 1.18 ([#4564](https://github.com/hashicorp/terraform-provider-google-beta/pull/4564))

FEATURES:
* **New Resource:** `google_dataplex_asset` ([#4543](https://github.com/hashicorp/terraform-provider-google-beta/pull/4543))
* **New Resource:** `google_gke_hub_membership_iam_binding` ([#4583](https://github.com/hashicorp/terraform-provider-google-beta/pull/4583))
* **New Resource:** `google_gke_hub_membership_iam_member` ([#4583](https://github.com/hashicorp/terraform-provider-google-beta/pull/4583))
* **New Resource:** `google_gke_hub_membership_iam_policy` ([#4583](https://github.com/hashicorp/terraform-provider-google-beta/pull/4583))

IMPROVEMENTS:
* certificatemanager: added `state`, `authorization_attempt_info` and `provisioning_issue` output fields to `google_certificate_manager_certificate` ([#4548](https://github.com/hashicorp/terraform-provider-google-beta/pull/4548))
* cloudfunctions2: added field `event_filters` to resource `google_cloudfunctions2_function` ([#4547](https://github.com/hashicorp/terraform-provider-google-beta/pull/4547))
* compute: added `certificate_map` to `compute_target_https_proxy` resource ([#4550](https://github.com/hashicorp/terraform-provider-google-beta/pull/4550))
* compute: added validation for name field on `google_compute_network` ([#4579](https://github.com/hashicorp/terraform-provider-google-beta/pull/4579))
* compute: made `port` optional in `google_compute_network_endpoint` to allow network endpoints to be associated with `GCE_VM_IP` network endpoint groups ([#4575](https://github.com/hashicorp/terraform-provider-google-beta/pull/4575))
* container: added support for additional values `APISERVER`, `CONTROLLER_MANAGER`, and `SCHEDULER` in `google_container_cluster.monitoring_config` ([#4565](https://github.com/hashicorp/terraform-provider-google-beta/pull/4565))
* gkehub: added `monitoring` and `mutation_enabled` fields to resource `feature_membership` ([#4572](https://github.com/hashicorp/terraform-provider-google-beta/pull/4572))
* gkehub: added better support for import for `google_gke_hub_membership` ([#4542](https://github.com/hashicorp/terraform-provider-google-beta/pull/4542))
* pubsub: added `bigquery_config` to `google_pubsub_subscription` ([#4545](https://github.com/hashicorp/terraform-provider-google-beta/pull/4545))
* scheduler: added `paused` field to `google_cloud_scheduler_job` ([#4535](https://github.com/hashicorp/terraform-provider-google-beta/pull/4535))
* scheduler: added `state` output field to `google_cloud_scheduler_job` ([#4535](https://github.com/hashicorp/terraform-provider-google-beta/pull/4535))

BUG FIXES:
* apigee: fixed an issue where `google_apigee_instance` creation would fail due to multiple concurrent instances ([#4584](https://github.com/hashicorp/terraform-provider-google-beta/pull/4584))
* billingbudget: fixed a bug where `google_billing_budget.budget_filter.services` was not updating. ([#4577](https://github.com/hashicorp/terraform-provider-google-beta/pull/4577))
* compute: fixed perma-diff on `google_compute_disk` for new arm64 images ([#4533](https://github.com/hashicorp/terraform-provider-google-beta/pull/4533))
* dataflow: fixed bug where permadiff would show on `google_dataflow_job.additional_experiments` ([#4576](https://github.com/hashicorp/terraform-provider-google-beta/pull/4576))
* storage: fixed a bug in `google_storage_bucket` where `name` was incorrectly validated. ([#4566](https://github.com/hashicorp/terraform-provider-google-beta/pull/4566))

## 4.31.0 (Aug 1, 2022)

FEATURES:
* **New Resource:** `google_dataplex_zone` ([#4511](https://github.com/hashicorp/terraform-provider-google-beta/pull/4511))

IMPROVEMENTS:
* bucket: added support for `matches_prefix` and `matches_suffix` in `condition` of a `lifecycle_rule` in  `google_storage_bucket` ([#4527](https://github.com/hashicorp/terraform-provider-google-beta/pull/4527))
* compute: added `network` and `subnetwork` fields to `google_compute_region_network_endpoint_group` for PSC. ([#4528](https://github.com/hashicorp/terraform-provider-google-beta/pull/4528))
* container: added field `boot_disk_kms_key` to `auto_provisioning_defaults` in `google_container_cluster` ([#4524](https://github.com/hashicorp/terraform-provider-google-beta/pull/4524))
* notebooks: added `bootDiskType` support for `PD_EXTREME` in `google_notebooks_instance` ([#4530](https://github.com/hashicorp/terraform-provider-google-beta/pull/4530))
* notebooks: added `softwareConfig.upgradeable`, `softwareConfig.postStartupScriptBehavior`, `softwareConfig.kernels` in `google_notebooks_runtime` ([#4530](https://github.com/hashicorp/terraform-provider-google-beta/pull/4530))
* storage: added name validation for `google_storage_bucket` ([#4532](https://github.com/hashicorp/terraform-provider-google-beta/pull/4532))

BUG FIXES:
* compute: fixed perma-diff on `google_compute_disk` for new arm64 images ([#4533](https://github.com/hashicorp/terraform-provider-google-beta/pull/4533))
* dns: fixed a bug where `google_dns_record_set` would create an inconsistent plan when using interpolated values in `rrdatas` ([#4515](https://github.com/hashicorp/terraform-provider-google-beta/pull/4515))
* kms: fixed setting of resource id post-import for `google_kms_crypto_key` ([#4520](https://github.com/hashicorp/terraform-provider-google-beta/pull/4520))
* provider: fixed a bug where user-agent was showing "dev" rather than the provider version ([#4509](https://github.com/hashicorp/terraform-provider-google-beta/pull/4509))

## 4.30.0 (July 25, 2022)

FEATURES:
* **New Data Source:** `google_service_account_jwt` ([#4489](https://github.com/hashicorp/terraform-provider-google-beta/pull/4489))
* **New Resource:** `google_certificate_map_entry` ([#4501](https://github.com/hashicorp/terraform-provider-google-beta/pull/4501))
* **New Resource:** `google_certificate_map` ([#4501](https://github.com/hashicorp/terraform-provider-google-beta/pull/4501))
* **New Resource:** `google_compute_backend_bucket_iam_binding` ([#4484](https://github.com/hashicorp/terraform-provider-google-beta/pull/4484))
* **New Resource:** `google_compute_backend_bucket_iam_member` ([#4484](https://github.com/hashicorp/terraform-provider-google-beta/pull/4484))
* **New Resource:** `google_compute_backend_bucket_iam_policy` ([#4484](https://github.com/hashicorp/terraform-provider-google-beta/pull/4484))
* **New Resource:** `google_dataproc_metastore_federation` ([#4482](https://github.com/hashicorp/terraform-provider-google-beta/pull/4482))
* **New Resource:** `google_dataproc_metastore_federation_iam_binding` ([#4482](https://github.com/hashicorp/terraform-provider-google-beta/pull/4482))
* **New Resource:** `google_dataproc_metastore_federation_iam_member` ([#4482](https://github.com/hashicorp/terraform-provider-google-beta/pull/4482))
* **New Resource:** `google_dataproc_metastore_federation_iam_policy` ([#4482](https://github.com/hashicorp/terraform-provider-google-beta/pull/4482))

IMPROVEMENTS:
* billingbudget: made `thresholdRules` optional in `google_billing_budget` ([#4480](https://github.com/hashicorp/terraform-provider-google-beta/pull/4480))
* compute: added `instance_termination_action` field to `google_compute_instance_template` resource to support Spot VM termination action ([#4488](https://github.com/hashicorp/terraform-provider-google-beta/pull/4488))
* compute: added `instance_termination_action` field to `google_compute_instance` resource to support Spot VM termination action ([#4488](https://github.com/hashicorp/terraform-provider-google-beta/pull/4488))
* compute: added `request_coalescing` and `bypass_cache_on_request_headers` fields to `compute_backend_bucket` ([#4484](https://github.com/hashicorp/terraform-provider-google-beta/pull/4484))
* compute: added field `all_instances_config` to `google_compute_instance_group_manager` and `google_compute_region_instance_group_manager` ([#4506](https://github.com/hashicorp/terraform-provider-google-beta/pull/4506))
* compute: added support for `esp` protocol in `google_compute_packet_mirroring.filters.ip_protocols` ([#4496](https://github.com/hashicorp/terraform-provider-google-beta/pull/4496))
* monitoring: added `evaluation_missing_data` field to `google_monitoring_alert_policy` ([#4502](https://github.com/hashicorp/terraform-provider-google-beta/pull/4502))
* notebooks: added field `reserved_ip_range` to `google_notebooks_runtime` ([#4492](https://github.com/hashicorp/terraform-provider-google-beta/pull/4492))

BUG FIXES:
* bigtable: fixed an incorrect diff when adding two or more clusters ([#4490](https://github.com/hashicorp/terraform-provider-google-beta/pull/4490))
* compute: allowed properly updating `adaptive_protection_config` in `compute_security_policy` ([#4478](https://github.com/hashicorp/terraform-provider-google-beta/pull/4478))
* notebooks: fixed a bug where`google_notebooks_runtime` can't be updated ([#4492](https://github.com/hashicorp/terraform-provider-google-beta/pull/4492))
* sql: fixed an issue in `google_sql_database_instance` where updates would fail because of the `collation` field ([#4505](https://github.com/hashicorp/terraform-provider-google-beta/pull/4505))

## 4.29.0 (July 18, 2022)

FEATURES:
* **New Resource:** `google_cloudiot_registry_iam_binding` ([#4452](https://github.com/hashicorp/terraform-provider-google-beta/pull/4452))
* **New Resource:** `google_cloudiot_registry_iam_member` ([#4452](https://github.com/hashicorp/terraform-provider-google-beta/pull/4452))
* **New Resource:** `google_cloudiot_registry_iam_policy` ([#4452](https://github.com/hashicorp/terraform-provider-google-beta/pull/4452))
* **New Resource:** `google_compute_snapshot_iam_binding` ([#4445](https://github.com/hashicorp/terraform-provider-google-beta/pull/4445))
* **New Resource:** `google_compute_snapshot_iam_member` ([#4445](https://github.com/hashicorp/terraform-provider-google-beta/pull/4445))
* **New Resource:** `google_compute_snapshot_iam_policy` ([#4445](https://github.com/hashicorp/terraform-provider-google-beta/pull/4445))

IMPROVEMENTS:
* container: added `binauthz_evaluation_mode` field to `resource_container_cluster`. ([#4451](https://github.com/hashicorp/terraform-provider-google-beta/pull/4451))
* kms: added support for MAC value in `google_kms_crypto_key.purpose` ([#4458](https://github.com/hashicorp/terraform-provider-google-beta/pull/4458))
* metastore: added `databaseType`, `releaseChannel`, and `hiveMetastoreConfig.endpointProtocol` arguments ([#4443](https://github.com/hashicorp/terraform-provider-google-beta/pull/4443))

BUG FIXES:
* bigquery: fixed case-sensitivity for `user_by_email` and `group_by_email` on `google_bigquery_dataset_access` ([#4446](https://github.com/hashicorp/terraform-provider-google-beta/pull/4446))
* clouddeploy: fixed permadiff on `execution_configs` in `google_clouddeploy_target` resource ([#4450](https://github.com/hashicorp/terraform-provider-google/pull/4450))
* cloudscheduler: fixed a diff on the last slash of uri on `google_cloud_scheduler_job` ([#4444](https://github.com/hashicorp/terraform-provider-google-beta/pull/4444))
* compute: fixed force recreation on `provisioned_iops` of `google_compute_disk` ([#4464](https://github.com/hashicorp/terraform-provider-google-beta/pull/4464))
* compute: fixed missing `network_interface.0.ipv6_access_config.0.external_ipv6` output on `google_compute_instance` ([#4470](https://github.com/hashicorp/terraform-provider-google-beta/pull/4470))
* documentai: fixed a bug where eu region could not be utilized for documentai resources ([#4472](https://github.com/hashicorp/terraform-provider-google-beta/pull/4472))
* gkehub: fixed a bug where `issuer` can't be updated on `google_gke_hub_membership` ([#4471](https://github.com/hashicorp/terraform-provider-google-beta/pull/4471))

## 4.28.0 (July 11, 2022)

FEATURES:
* **New Resource:** google_bigquery_connection_iam_binding ([#4437](https://github.com/hashicorp/terraform-provider-google-beta/pull/4437))
* **New Resource:** google_bigquery_connection_iam_member ([#4437](https://github.com/hashicorp/terraform-provider-google-beta/pull/4437))
* **New Resource:** google_bigquery_connection_iam_policy ([#4437](https://github.com/hashicorp/terraform-provider-google-beta/pull/4437))
* **New Resource:** google_cloud_tasks_queue_iam_binding ([#4427](https://github.com/hashicorp/terraform-provider-google-beta/pull/4427))
* **New Resource:** google_cloud_tasks_queue_iam_member ([#4427](https://github.com/hashicorp/terraform-provider-google-beta/pull/4427))
* **New Resource:** google_cloud_tasks_queue_iam_policy ([#4427](https://github.com/hashicorp/terraform-provider-google-beta/pull/4427))
* **New Resource:** google_dataproc_autoscaling_policy_iam_binding ([#4441](https://github.com/hashicorp/terraform-provider-google-beta/pull/4441))
* **New Resource:** google_dataproc_autoscaling_policy_iam_member ([#4441](https://github.com/hashicorp/terraform-provider-google-beta/pull/4441))
* **New Resource:** google_dataproc_autoscaling_policy_iam_policy ([#4441](https://github.com/hashicorp/terraform-provider-google-beta/pull/4441))
* **New Resource:** google_dataproc_metastore_service_iam_binding ([#4416](https://github.com/hashicorp/terraform-provider-google-beta/pull/4416))
* **New Resource:** google_dataproc_metastore_service_iam_member ([#4416](https://github.com/hashicorp/terraform-provider-google-beta/pull/4416))
* **New Resource:** google_dataproc_metastore_service_iam_policy ([#4416](https://github.com/hashicorp/terraform-provider-google-beta/pull/4416))

IMPROVEMENTS:
* bigquery: fixed a permadiff in `google_bigquery_job.query. destination_table` ([#4401](https://github.com/hashicorp/terraform-provider-google-beta/pull/4401))
* billing: added `calendar_period` and `custom_period` fields to `google_billing_budget` ([#4429](https://github.com/hashicorp/terraform-provider-google-beta/pull/4429))
* cloudsql: added attribute `project` to data source `google_sql_backup_run` ([#4402](https://github.com/hashicorp/terraform-provider-google-beta/pull/4402))
* composer: added CMEK, PUPI and IP_masq_agent support for Composer 2 in `google_composer_environment` resource ([#4430](https://github.com/hashicorp/terraform-provider-google-beta/pull/4430))
* compute: added `max_ports_per_vm` field to `google_compute_router_nat` resource ([#4400](https://github.com/hashicorp/terraform-provider-google-beta/pull/4400))
* compute: added `GCE_VM_IP` support to `google_compute_network_endpoint_group` resource. ([#4434](https://github.com/hashicorp/terraform-provider-google-beta/pull/4434))
* privateca: added support to subordinate CA activation ([#4422](https://github.com/hashicorp/terraform-provider-google-beta/pull/4422))
* redis: added CMEK key field `customer_managed_key` in `google_redis_instance ` ([#4435](https://github.com/hashicorp/terraform-provider-google-beta/pull/4435))
* spanner: added field `version_retention_period` to `google_spanner_database` resource ([#4424](https://github.com/hashicorp/terraform-provider-google-beta/pull/4424))
* sql: added `settings.location_preference.secondary_zone` field in `google_sql_database_instance` ([#4433](https://github.com/hashicorp/terraform-provider-google-beta/pull/4433))
* sql: added `sql_server_audit_config` field in `google_sql_database_instance` ([#4403](https://github.com/hashicorp/terraform-provider-google-beta/pull/4403))

BUG FIXES:
* composer: fixed a problem with updating Cloud Composer's `scheduler_count` field (https://github.com/hashicorp/terraform-provider-google/issues/11940) ([#4408](https://github.com/hashicorp/terraform-provider-google-beta/pull/4408))
* composer: fixed permadiff on `private_environment_config.cloud_composer_connection_subnetwork` ([#4411](https://github.com/hashicorp/terraform-provider-google-beta/pull/4411))
* container: fixed an issue where `node_config.min_cpu_platform` could cause a perma-diff in `google_container_cluster` ([#4426](https://github.com/hashicorp/terraform-provider-google-beta/pull/4426))
* filestore: fixed a case where `google_filestore_instance.networks.network` would incorrectly see a diff between state and config when the network `id` format was used ([#4431](https://github.com/hashicorp/terraform-provider-google-beta/pull/4431))
* serviceusage: fixed an issue where `google_project_service_identity` didn't handle service identities without emails correctly ([#4432](https://github.com/hashicorp/terraform-provider-google-beta/pull/4432))


## 4.27.0 (June 27, 2022)

IMPROVEMENTS:
* clouddeploy: added `suspend` field to `google_clouddeploy_delivery_pipeline` resource ([#4394](https://github.com/hashicorp/terraform-provider-google-beta/pull/4394))
* compute: added maxPortsPerVm field to `google_compute_router_nat` resource ([#4400](https://github.com/hashicorp/terraform-provider-google-beta/pull/4400))
* compute: added `psc_connection_id` and `psc_connection_status` output fields to `google_compute_forwarding_rule` and `google_compute_global_forwarding_rule` resources ([#4392](https://github.com/hashicorp/terraform-provider-google-beta/pull/4392))
* container: added `tpu_config` to `google_container_cluster` (beta only) ([#4390](https://github.com/hashicorp/terraform-provider-google-beta/pull/4390))
* containeraws: made `config.instance_type` field updatable in `google_container_aws_node_pool` ([#4392](https://github.com/hashicorp/terraform-provider-google-beta/pull/4392))

BUG FIXES:
* compute: fixed default handling for `enable_dynamic_port_allocation ` to be managed by the api ([#4391](https://github.com/hashicorp/terraform-provider-google-beta/pull/4391))
* vertexai: Fixed a bug where terraform crashes when `force_destroy` is set in `google_vertex_ai_featurestore` resource ([#4398](https://github.com/hashicorp/terraform-provider-google-beta/pull/4398))

## 4.26.0 (June 21, 2022)

FEATURES:
* **New Resource:** `google_cloudfunctions2_function_iam_binding` ([#4377](https://github.com/hashicorp/terraform-provider-google-beta/pull/4377))
* **New Resource:** `google_cloudfunctions2_function_iam_member` ([#4377](https://github.com/hashicorp/terraform-provider-google-beta/pull/4377))
* **New Resource:** `google_cloudfunctions2_function_iam_policy` ([#4377](https://github.com/hashicorp/terraform-provider-google-beta/pull/4377))
* **New Resource:** `google_compute_region_ssl_policy` ([#4376](https://github.com/hashicorp/terraform-provider-google-beta/pull/4376))
* **New Resource:** `google_documentai_processor` ([#4389](https://github.com/hashicorp/terraform-provider-google-beta/pull/4389))
* **New Resource:** `google_documentai_processor_default_version` ([#4389](https://github.com/hashicorp/terraform-provider-google-beta/pull/4389))

IMPROVEMENTS:
* accesscontextmanager: Added `external_resources` to `egress_to` in `google_access_context_manager_service_perimeter` and `google_access_context_manager_service_perimeters` resource ([#4378](https://github.com/hashicorp/terraform-provider-google-beta/pull/4378))
* apigateway: Added `grpc_services` and `managed_service_configs` to `google_api_gateway_api_config` ([#4388](https://github.com/hashicorp/terraform-provider-google-beta/pull/4388))
* cloudbuild: Added `include_build_logs` to `google_cloudbuild_trigger` ([#4380](https://github.com/hashicorp/terraform-provider-google-beta/pull/4380))
* compute: Added `ssl_policy` field to `google_compute_region_target_https_proxy` ([#4376](https://github.com/hashicorp/terraform-provider-google-beta/pull/4376))
* container: Added `managed_prometheus` to `monitoring_config` in `google_container_cluster` ([#4373](https://github.com/hashicorp/terraform-provider-google-beta/pull/4373))
* container: Added `tpu_config` to `google_container_cluster` ([#4390](https://github.com/hashicorp/terraform-provider-google-beta/pull/4390))

BUG FIXES:
* dns: Fixed a bug where `google_dns_record_set` resource can not be changed from default routing to Geo routing policy. ([#4383](https://github.com/hashicorp/terraform-provider-google-beta/pull/4383))
* sql: Fixed a bug where `google_sql_database_instance` would fail if a replica was created, with an encryption key, in a different region than the master instance. ([#4379](https://github.com/hashicorp/terraform-provider-google-beta/pull/4379))

## 4.25.0 (June 15, 2022)

IMPROVEMENTS:
* bigquery: added `connection_id` to `external_data_configuration` for `google_bigquery_table` ([#4365](https://github.com/hashicorp/terraform-provider-google-beta/pull/4365))
* cloudfunctions2: added support for configuring `service_account_email` to `google_cloudfunctions2_function` resource ([#4367](https://github.com/hashicorp/terraform-provider-google-beta/pull/4367))
* compute: added `advanced_options_config` to `google_compute_security_policy` ([#4354](https://github.com/hashicorp/terraform-provider-google-beta/pull/4354))
* compute: added `cache_key_policy` field to `google_compute_backend_bucket` resource ([#4349](https://github.com/hashicorp/terraform-provider-google-beta/pull/4349))
* compute: added `include_named_cookies` to `cdn_policy` on `compute_backend_service` resource ([#4358](https://github.com/hashicorp/terraform-provider-google-beta/pull/4358))
* compute: added internal IPv6 support on `google_compute_network` and `google_compute_subnetwork` ([#4368](https://github.com/hashicorp/terraform-provider-google-beta/pull/4368))
* container: added `managed_prometheus` to `monitoring_config` in `google_container_cluster` ([#4373](https://github.com/hashicorp/terraform-provider-google-beta/pull/4373))
* container: added `spot` field to `node_config` sub-resource ([#4350](https://github.com/hashicorp/terraform-provider-google-beta/pull/4350))
* gkehub: added `prevent_drift` field to `google_gke_hub_feature_membership` resource ([#4370](https://github.com/hashicorp/terraform-provider-google-beta/pull/4370))
* monitoring: added support for JSONPath content matchers to `google_monitoring_uptime_check_config` resource ([#4361](https://github.com/hashicorp/terraform-provider-google-beta/pull/4361))
* monitoring: added support for `user_labels` to `google_monitoring_slo` resource ([#4363](https://github.com/hashicorp/terraform-provider-google-beta/pull/4363))
* sql: added `sql_server_user_details` field to `google_sql_user` resource ([#4364](https://github.com/hashicorp/terraform-provider-google-beta/pull/4364))

BUG FIXES:
* certificatemanager: fixed bug where `DEFAULT` scope would permadiff and force replace the certificate. ([#4356](https://github.com/hashicorp/terraform-provider-google-beta/pull/4356))
* dns: fixed perma-diff for updated labels in `google_dns_managed_zone` ([#4372](https://github.com/hashicorp/terraform-provider-google-beta/pull/4372))
* storagetransfer: fixed perm diff on transfer_options for `google_storage_transfer_job` ([#4357](https://github.com/hashicorp/terraform-provider-google-beta/pull/4357))

## 4.24.0 (June 6, 2022)

IMPROVEMENTS:
* compute: added `cache_key_policy` field to `google_compute_backend_bucket` resource ([#4349](https://github.com/hashicorp/terraform-provider-google-beta/pull/4349))

## 4.23.0 (June 1, 2022)

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
