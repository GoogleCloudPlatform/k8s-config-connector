// Copyright 2025 Google LLC
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
// proto.message: google.cloud.alloydb.v1beta.Instance

package alloydb

import (
	pb "cloud.google.com/go/alloydb/apiv1beta/alloydbpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(instanceFuzzer())
}

func instanceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Instance{},
		AlloyDBInstanceSpec_FromProto, AlloyDBInstanceSpec_ToProto,
		AlloyDBInstanceStatus_FromProto, AlloyDBInstanceStatus_ToProto,
	)

	f.SpecField(".display_name")
	f.SpecField(".instance_type")
	f.SpecField(".machine_config")
	f.SpecField(".availability_type")
	f.SpecField(".gce_zone")
	f.SpecField(".database_flags")
	f.SpecField(".read_pool_config")
	f.SpecField(".annotations")
	f.SpecField(".network_config")

	f.StatusField(".uid")
	f.StatusField(".name")
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".state")
	f.StatusField(".ip_address")
	f.StatusField(".public_ip_address")
	f.StatusField(".outbound_public_ip_addresses")

	f.Unimplemented_Internal(".satisfies_pzs")

	f.Unimplemented_LabelsAnnotations(".labels")

	f.Unimplemented_NotYetTriaged(".gemini_config")
	f.Unimplemented_NotYetTriaged(".query_insights_config")
	f.Unimplemented_NotYetTriaged(".observability_config")
	f.Unimplemented_NotYetTriaged(".update_policy")
	f.Unimplemented_NotYetTriaged(".client_connection_config")
	f.Unimplemented_NotYetTriaged(".psc_instance_config")
	f.Unimplemented_NotYetTriaged(".gca_config")
	f.Unimplemented_NotYetTriaged(".delete_time")
	f.Unimplemented_NotYetTriaged(".etag")
	f.Unimplemented_NotYetTriaged(".writable_node")
	f.Unimplemented_NotYetTriaged(".nodes")
	f.Unimplemented_NotYetTriaged(".reconciling")
	f.Unimplemented_NotYetTriaged(".connection_pool_config")

	return f
}
