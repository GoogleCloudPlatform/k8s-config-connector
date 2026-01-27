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
// proto.message: google.cloud.sql.v1beta4.DatabaseInstance
// api.group: sql.cnrm.cloud.google.com

package sql

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpclients/generated/google/cloud/sql/v1beta4"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(sqlInstanceFuzzer())
}

func sqlInstanceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.DatabaseInstance{},
		DatabaseInstance_FromProto, DatabaseInstance_ToProto,
		DatabaseInstanceObservedState_FromProto, DatabaseInstanceObservedState_ToProto,
	)

	f.SpecField(".kind")
	f.SpecField(".state")
	f.SpecField(".database_version")
	f.SpecField(".settings")
	f.SpecField(".etag")
	f.SpecField(".failover_replica")
	f.SpecField(".master_instance_name")
	f.SpecField(".replica_names")
	f.SpecField(".max_disk_size")
	f.SpecField(".current_disk_size")
	f.SpecField(".ip_addresses")
	f.SpecField(".server_ca_cert")
	f.SpecField(".instance_type")
	f.SpecField(".project")
	f.SpecField(".ipv6_address")
	f.SpecField(".service_account_email_address")
	f.SpecField(".on_premises_configuration")
	f.SpecField(".replica_configuration")
	f.SpecField(".backend_type")
	f.SpecField(".self_link")
	f.SpecField(".suspension_reason")
	f.SpecField(".connection_name")
	f.SpecField(".name")
	f.SpecField(".region")
	f.SpecField(".gce_zone")
	f.SpecField(".secondary_gce_zone")
	f.SpecField(".disk_encryption_configuration")
	f.SpecField(".disk_encryption_status")
	f.SpecField(".root_password")
	f.SpecField(".scheduled_maintenance")
	f.SpecField(".satisfies_pzs")
	f.SpecField(".out_of_disk_report")
	f.SpecField(".maintenance_version")
	f.Unimplemented_NotYetTriaged(".sql_network_architecture")
	f.SpecField(".replication_cluster")
	f.SpecField(".gemini_config")

	f.StatusField(".settings")
	f.StatusField(".database_installed_version")
	f.StatusField(".create_time")
	f.StatusField(".available_maintenance_versions")
	f.StatusField(".upgradable_database_versions")
	f.StatusField(".psc_service_attachment_link")
	f.StatusField(".dns_name")
	f.StatusField(".primary_dns_name")
	f.StatusField(".write_endpoint")
	f.StatusField(".replication_cluster")
	f.StatusField(".gemini_config")

	return f
}
