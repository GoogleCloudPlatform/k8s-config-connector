// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.clouddms.v1.MigrationJob
// api.group: clouddms.cnrm.cloud.google.com

package clouddms

import (
	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dmsMigrationJobFuzzer())
}

func dmsMigrationJobFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.MigrationJob{},
		CloudDMSMigrationJobSpec_FromProto, CloudDMSMigrationJobSpec_ToProto,
		CloudDMSMigrationJobObservedState_FromProto, CloudDMSMigrationJobObservedState_ToProto,
	)

	f.SpecFields.Insert(".type")
	f.SpecFields.Insert(".source")
	f.SpecFields.Insert(".destination")
	f.SpecFields.Insert(".conversion_workspace")

	f.Unimplemented_NotYetTriaged(".phase")
	f.Unimplemented_NotYetTriaged(".create_time")
	f.Unimplemented_NotYetTriaged(".update_time")
	f.Unimplemented_NotYetTriaged(".duration")
	f.Unimplemented_NotYetTriaged(".error")

	f.Unimplemented_NotYetTriaged(".labels")
	f.Unimplemented_NotYetTriaged(".display_name")
	f.Unimplemented_NotYetTriaged(".state")
	f.Unimplemented_NotYetTriaged(".cmek_key_name")
	f.Unimplemented_NotYetTriaged(".filter")
	f.Unimplemented_NotYetTriaged(".dump_path")
	f.Unimplemented_NotYetTriaged(".dump_flags")
	f.Unimplemented_NotYetTriaged(".dump_flags.dump_flags")
	f.Unimplemented_NotYetTriaged(".source_database")
	f.Unimplemented_NotYetTriaged(".source_database.provider")
	f.Unimplemented_NotYetTriaged(".source_database.engine")
	f.Unimplemented_NotYetTriaged(".destination_database")
	f.Unimplemented_NotYetTriaged(".destination_database.provider")
	f.Unimplemented_NotYetTriaged(".destination_database.engine")
	f.Unimplemented_NotYetTriaged(".performance_config")
	f.Unimplemented_NotYetTriaged(".performance_config.dump_parallel_level")
	f.Unimplemented_NotYetTriaged(".reverse_ssh_connectivity")
	f.Unimplemented_NotYetTriaged(".reverse_ssh_connectivity.vm_ip")
	f.Unimplemented_NotYetTriaged(".reverse_ssh_connectivity.vm_port")
	f.Unimplemented_NotYetTriaged(".reverse_ssh_connectivity.vm")
	f.Unimplemented_NotYetTriaged(".reverse_ssh_connectivity.vpc")
	f.Unimplemented_NotYetTriaged(".vpc_peering_connectivity")
	f.Unimplemented_NotYetTriaged(".vpc_peering_connectivity.vpc")
	f.Unimplemented_NotYetTriaged(".static_ip_connectivity")
	f.Unimplemented_NotYetTriaged(".end_time")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_Etag()

	return f
}
