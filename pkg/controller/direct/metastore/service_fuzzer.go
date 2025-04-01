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
// proto.message: google.cloud.metastore.v1.Service
// api.group: metastore.cnrm.cloud.google.com

package metastore

import (
	pb "cloud.google.com/go/metastore/apiv1/metastorepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(metastoreServiceFuzzer())
}

func metastoreServiceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Service{},
		MetastoreServiceSpec_FromProto, MetastoreServiceSpec_ToProto,
		MetastoreServiceObservedState_FromProto, MetastoreServiceObservedState_ToProto,
	)

	f.SpecFields.Insert(".hive_metastore_config")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".network")
	f.SpecFields.Insert(".port")
	f.SpecFields.Insert(".tier")
	f.SpecFields.Insert(".maintenance_window")
	f.SpecFields.Insert(".release_channel")
	f.SpecFields.Insert(".encryption_config")
	f.SpecFields.Insert(".network_config")
	f.SpecFields.Insert(".database_type")
	f.SpecFields.Insert(".telemetry_config")
	f.SpecFields.Insert(".scaling_config")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".endpoint_uri")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".state_message")
	f.StatusFields.Insert(".artifact_gcs_uri")
	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".metadata_management_activity")
	f.StatusFields.Insert(".network_config")

	f.UnimplementedFields.Insert(".name") // special field
	f.UnimplementedFields.Insert(".hive_metastore_config.auxiliary_versions")
	f.UnimplementedFields.Insert(".metadata_management_activity.metadata_exports")
	f.UnimplementedFields.Insert(".metadata_management_activity.restores")
	f.UnimplementedFields.Insert(".network_config.consumers.endpoint_uri")
	f.UnimplementedFields.Insert(".network_config.consumers.endpoint_location")

	return f
}
