// Copyright 2024 Google LLC
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
// proto.message: google.cloud.datastream.v1.ConnectionProfile
// api.group: datastream.cnrm.cloud.google.com

package datastream

import (
	pb "cloud.google.com/go/datastream/apiv1/datastreampb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(datastreamConnectionProfileFuzzer())
}

func datastreamConnectionProfileFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ConnectionProfile{},
		DatastreamConnectionProfileSpec_FromProto, DatastreamConnectionProfileSpec_ToProto,
		DatastreamConnectionProfileObservedState_FromProto, DatastreamConnectionProfileObservedState_ToProto,
	)

	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".oracle_profile")
	f.SpecFields.Insert(".gcs_profile")
	f.SpecFields.Insert(".mysql_profile")
	f.SpecFields.Insert(".bigquery_profile")
	f.SpecFields.Insert(".sql_server_profile")
	f.SpecFields.Insert(".static_service_ip_connectivity")
	f.SpecFields.Insert(".forward_ssh_connectivity")
	f.SpecFields.Insert(".private_connectivity")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".oracle_profile")
	f.StatusFields.Insert(".mysql_profile")

	f.UnimplementedFields.Insert(".name") // special field
	f.UnimplementedFields.Insert(".postgresql_profile")

	// sensitive fields do not roundtrip
	f.UnimplementedFields.Insert(".forward_ssh_connectivity.username")
	f.UnimplementedFields.Insert(".forward_ssh_connectivity.password")
	f.UnimplementedFields.Insert(".mysql_profile.username")
	f.UnimplementedFields.Insert(".mysql_profile.password")
	f.UnimplementedFields.Insert(".oracle_profile.username")
	f.UnimplementedFields.Insert(".oracle_profile.password")
	f.UnimplementedFields.Insert(".oracle_profile.oracle_asm_config.username")
	f.UnimplementedFields.Insert(".oracle_profile.oracle_asm_config.password")
	f.UnimplementedFields.Insert(".sql_server_profile.username")
	f.UnimplementedFields.Insert(".sql_server_profile.password")

	return f
}
