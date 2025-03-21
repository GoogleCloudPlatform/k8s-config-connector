// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//\thttp://www.apache.org/licenses/LICENSE-2.0
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
\tpb "cloud.google.com/go/metastore/apiv1/metastorepb"
\t"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
\tfuzztesting.RegisterKRMFuzzer(MetastoreServiceFuzzer())
}

func MetastoreServiceFuzzer() fuzztesting.KRMFuzzer {
\tf := fuzztesting.NewKRMTypedFuzzer(&pb.Service{},\n\t\tMetastoreServiceSpec_FromProto, MetastoreServiceSpec_ToProto,\n\t\tMetastoreServiceObservedState_FromProto, MetastoreServiceObservedState_ToProto,\n\t)

\tf.SpecFields.Insert(".hive_metastore_config")
\tf.SpecFields.Insert(".labels")
\tf.SpecFields.Insert(".network")
\tf.SpecFields.Insert(".port")
\tf.SpecFields.Insert(".tier")
\tf.SpecFields.Insert(".maintenance_window")
\tf.SpecFields.Insert(".release_channel")
\tf.SpecFields.Insert(".encryption_config")
\tf.SpecFields.Insert(".network_config")
\tf.SpecFields.Insert(".database_type")
\tf.SpecFields.Insert(".telemetry_config")
\tf.SpecFields.Insert(".scaling_config")

\tf.StatusFields.Insert(".create_time")
\tf.StatusFields.Insert(".update_time")
\tf.StatusFields.Insert(".endpoint_uri")
\tf.StatusFields.Insert(".state")
\tf.StatusFields.Insert(".state_message")
\tf.StatusFields.Insert(".artifact_gcs_uri")
\tf.StatusFields.Insert(".uid")
\tf.StatusFields.Insert(".metadata_management_activity")
\tf.StatusFields.Insert(".network_config")
\tf.UnimplementedFields.Insert("name")
\treturn f
}
