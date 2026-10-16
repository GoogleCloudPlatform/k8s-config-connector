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
// proto.message: google.cloud.oracledatabase.v1.ExadbVmCluster
// api.group: oracledatabase.cnrm.cloud.google.com

package oracledatabase

import (
	pb "cloud.google.com/go/oracledatabase/apiv1/oracledatabasepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(oracleDatabaseExadbVMClusterFuzzer())
}

func oracleDatabaseExadbVMClusterFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ExadbVmCluster{},
		OracleDatabaseExadbVMClusterSpec_FromProto, OracleDatabaseExadbVMClusterSpec_ToProto,
		OracleDatabaseExadbVMClusterObservedState_FromProto, OracleDatabaseExadbVMClusterObservedState_ToProto,
	)

	// Spec fields
	f.SpecField(".display_name")
	f.SpecField(".properties")
	f.SpecField(".properties.cluster_name")
	f.SpecField(".properties.grid_image_id")
	f.SpecField(".properties.node_count")
	f.SpecField(".properties.enabled_ecpu_count_per_node")
	f.SpecField(".properties.additional_ecpu_count_per_node")
	f.SpecField(".properties.vm_file_system_storage")
	f.SpecField(".properties.vm_file_system_storage.size_in_gbs_per_node")
	f.SpecField(".properties.license_model")
	f.SpecField(".properties.hostname_prefix")
	f.SpecField(".properties.ssh_public_keys")
	f.SpecField(".properties.data_collection_options")
	f.SpecField(".properties.data_collection_options.is_diagnostics_events_enabled")
	f.SpecField(".properties.data_collection_options.is_health_monitoring_enabled")
	f.SpecField(".properties.data_collection_options.is_incident_logs_enabled")
	f.SpecField(".properties.time_zone")
	f.SpecField(".properties.time_zone.id")
	f.SpecField(".properties.time_zone.version")
	f.SpecField(".properties.shape_attribute")
	f.SpecField(".properties.scan_listener_port_tcp")

	// Status fields
	f.StatusField(".gcp_oracle_zone")
	f.StatusField(".create_time")
	f.StatusField(".entitlement_id")
	f.StatusField(".properties.hostname")
	f.StatusField(".properties.lifecycle_state")
	f.StatusField(".properties.memory_size_gb")
	f.StatusField(".properties.oci_uri")
	f.StatusField(".properties.gi_version")
	f.StatusField(".identity_connector")
	f.StatusField(".identity_connector.service_agent_email")
	f.StatusField(".identity_connector.connection_state")

	// Unimplemented / identity / label fields
	f.Unimplemented_Identity(".name")
	f.Unimplemented_LabelsAnnotations(".labels")

	// References mapped to Ref in Spec
	f.Unimplemented_NotYetTriaged(".odb_network")
	f.Unimplemented_NotYetTriaged(".odb_subnet")
	f.Unimplemented_NotYetTriaged(".backup_odb_subnet")
	f.Unimplemented_NotYetTriaged(".properties.exascale_db_storage_vault")

	return f
}
