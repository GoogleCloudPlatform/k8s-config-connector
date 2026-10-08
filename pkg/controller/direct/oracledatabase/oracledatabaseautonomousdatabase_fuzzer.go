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

// +tool:fuzz-gen
// proto.message: google.cloud.oracledatabase.v1.AutonomousDatabase
// api.group: oracledatabase.cnrm.cloud.google.com

package oracledatabase

import (
	pb "cloud.google.com/go/oracledatabase/apiv1/oracledatabasepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(oracleDatabaseAutonomousDatabaseFuzzer())
}

func oracleDatabaseAutonomousDatabaseFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.AutonomousDatabase{},
		OracleDatabaseAutonomousDatabaseSpec_FromProto, OracleDatabaseAutonomousDatabaseSpec_ToProto,
		OracleDatabaseAutonomousDatabaseObservedState_FromProto, OracleDatabaseAutonomousDatabaseObservedState_ToProto,
	)

	f.Unimplemented_Identity(".name")

	// Spec fields
	f.SpecField(".database")
	f.SpecField(".display_name")
	f.SpecField(".properties")
	f.SpecField(".labels")
	f.SpecField(".network")
	f.SpecField(".cidr")
	f.SpecField(".odb_network")
	f.SpecField(".odb_subnet")
	f.SpecField(".source_config")
	f.SpecField(".admin_password_secret_version")

	// Status fields
	f.StatusField(".properties")
	f.StatusField(".create_time")
	f.StatusField(".entitlement_id")
	f.StatusField(".gcp_oracle_zone")
	f.StatusField(".peer_autonomous_databases")
	f.StatusField(".disaster_recovery_supported_locations")

	// Unimplemented / write-only fields
	f.Unimplemented_NotYetTriaged(".admin_password")

	// Subfields of .properties - Spec
	f.SpecField(".properties.compute_count")
	f.SpecField(".properties.cpu_core_count")
	f.SpecField(".properties.data_storage_size_tb")
	f.SpecField(".properties.data_storage_size_gb")
	f.SpecField(".properties.db_workload")
	f.SpecField(".properties.db_edition")
	f.SpecField(".properties.character_set")
	f.SpecField(".properties.n_character_set")
	f.SpecField(".properties.private_endpoint_ip")
	f.SpecField(".properties.private_endpoint_label")
	f.SpecField(".properties.db_version")
	f.SpecField(".properties.is_auto_scaling_enabled")
	f.SpecField(".properties.is_storage_auto_scaling_enabled")
	f.SpecField(".properties.license_type")
	f.SpecField(".properties.customer_contacts")
	f.SpecField(".properties.customer_contacts[].email")
	f.SpecField(".properties.secret_id")
	f.SpecField(".properties.vault_id")
	f.SpecField(".properties.maintenance_schedule_type")
	f.SpecField(".properties.mtls_connection_required")
	f.SpecField(".properties.backup_retention_period_days")
	f.SpecField(".properties.encryption_key")
	f.SpecField(".properties.encryption_key.kms_key")
	f.SpecField(".properties.encryption_key.provider")
	f.SpecField(".properties.local_data_guard_enabled")
	f.SpecField(".properties.local_adg_auto_failover_max_data_loss_limit_duration")
	f.SpecField(".properties.refreshable_clone")

	// Subfields of .properties - Status
	f.StatusField(".properties.ocid")
	f.StatusField(".properties.actual_used_data_storage_size_tb")
	f.StatusField(".properties.allocated_storage_size_tb")
	f.StatusField(".properties.apex_details")
	f.StatusField(".properties.apex_details.apex_version")
	f.StatusField(".properties.apex_details.ords_version")
	f.StatusField(".properties.lifecycle_details")
	f.StatusField(".properties.state")
	f.StatusField(".properties.autonomous_container_database_id")
	f.StatusField(".properties.available_upgrade_versions")
	f.StatusField(".properties.connection_strings")
	f.StatusField(".properties.failed_data_recovery_duration")
	f.StatusField(".properties.is_local_data_guard_enabled")
	f.StatusField(".properties.local_adg_auto_failover_max_data_loss_limit")
	f.StatusField(".properties.local_standby_db")
	f.StatusField(".properties.local_disaster_recovery_type")
	f.StatusField(".properties.data_safe_state")
	f.StatusField(".properties.database_management_state")
	f.StatusField(".properties.open_mode")
	f.StatusField(".properties.operations_insights_state")
	f.StatusField(".properties.permission_level")
	f.StatusField(".properties.private_endpoint")
	f.StatusField(".properties.refreshable_mode")
	f.StatusField(".properties.refreshable_state")
	f.StatusField(".properties.role")
	f.StatusField(".properties.scheduled_operation_details")
	f.StatusField(".properties.sql_web_developer_url")
	f.StatusField(".properties.supported_clone_regions")
	f.StatusField(".properties.used_data_storage_size_tbs")
	f.StatusField(".properties.oci_url")
	f.StatusField(".properties.next_long_term_backup_time")
	f.StatusField(".properties.data_guard_role_changed_time")
	f.StatusField(".properties.disaster_recovery_role_changed_time")
	f.StatusField(".properties.maintenance_begin_time")
	f.StatusField(".properties.maintenance_end_time")
	f.StatusField(".properties.encryption_key_history_entries")
	f.StatusField(".properties.service_agent_email")
	f.SpecField(".properties.allowlisted_ips")
	f.StatusField(".properties.are_primary_allowlisted_ips_used")
	f.StatusField(".properties.connection_urls")
	f.StatusField(".properties.memory_table_gbs")
	f.StatusField(".properties.memory_per_oracle_compute_unit_gbs")
	f.StatusField(".properties.peer_db_ids")
	f.StatusField(".properties.total_auto_backup_storage_size_gbs")

	// Subfields of .source_config - Spec
	f.SpecField(".source_config.automatic_backups_replication_enabled")
	f.SpecField(".source_config.autonomous_database")
	f.SpecField(".source_config.autonomous_database_backup")
	f.SpecField(".source_config.backup_time")
	f.SpecField(".source_config.clone_type")
	f.SpecField(".source_config.refreshable_mode")
	f.SpecField(".source_config.type")
	f.SpecField(".source_config.use_latest_available_backup")
	f.SpecField(".source_config.auto_refresh_frequency_seconds")
	f.SpecField(".source_config.auto_refresh_point_lag_seconds")
	f.SpecField(".source_config.auto_refresh_start_time")

	return f
}
