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
// proto.message: google.cloud.bigquery.connection.v1.Connection
// api.group: bigqueryconnection.cnrm.cloud.google.com

package bigqueryconnection

import (
	pb "cloud.google.com/go/bigquery/connection/apiv1/connectionpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(bigQueryConnectionConnectionFuzzer())
}

func bigQueryConnectionConnectionFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Connection{},
		BigQueryConnectionConnectionSpec_FromProto, BigQueryConnectionConnectionSpec_ToProto,
		BigQueryConnectionConnectionObservedState_FromProto, BigQueryConnectionConnectionObservedState_ToProto,
	)

	f.SpecFields.Insert(".friendly_name")
	f.SpecFields.Insert(".description")

	// aws spec fields
	f.SpecFields.Insert(".aws.access_role.iam_role_id")

	// azure spec fields
	f.SpecFields.Insert(".azure.customer_tenant_id")
	f.SpecFields.Insert(".azure.federated_application_client_id")

	// cloud_resource spec fields
	f.SpecFields.Insert(".cloud_resource")

	// cloud_sql spec fields
	f.SpecFields.Insert(".cloud_sql.instance_id")
	f.SpecFields.Insert(".cloud_sql.database")
	f.SpecFields.Insert(".cloud_sql.type")

	// cloud_spanner spec fields
	f.SpecFields.Insert(".cloud_spanner.database")
	f.SpecFields.Insert(".cloud_spanner.use_parallelism")
	f.SpecFields.Insert(".cloud_spanner.use_data_boost")
	f.SpecFields.Insert(".cloud_spanner.use_serverless_analytics")
	f.SpecFields.Insert(".cloud_spanner.max_parallelism")
	f.SpecFields.Insert(".cloud_spanner.database_role")

	// spark spec fields
	f.SpecFields.Insert(".spark.metastore_service_config.metastore_service")
	f.SpecFields.Insert(".spark.spark_history_server_config.dataproc_cluster")

	f.StatusFields.Insert(".has_credential")

	// Status/ObservedState mappings:
	f.StatusFields.Insert(".aws.access_role.identity")
	f.StatusFields.Insert(".azure.application")
	f.StatusFields.Insert(".azure.client_id")
	f.StatusFields.Insert(".azure.object_id")
	f.StatusFields.Insert(".azure.redirect_uri")
	f.StatusFields.Insert(".azure.identity")
	f.StatusFields.Insert(".cloud_resource.service_account_id")
	f.StatusFields.Insert(".cloud_sql.service_account_id")
	f.StatusFields.Insert(".spark.service_account_id")

	f.UnimplementedFields.Insert(".name")
	f.UnimplementedFields.Insert(".creation_time")
	f.UnimplementedFields.Insert(".last_modified_time")

	f.Unimplemented_NotYetTriaged(".salesforce_data_cloud")
	f.Unimplemented_NotYetTriaged(".aws.cross_account_role")
	f.Unimplemented_NotYetTriaged(".cloud_sql.credential.username")
	f.Unimplemented_NotYetTriaged(".cloud_sql.credential.password")

	f.FilterSpec = func(in *pb.Connection) {
		if in.GetCloudSql() != nil {
			if in.GetCloudSql().InstanceId == "" {
				in.GetCloudSql().InstanceId = "projects/p/instances/i"
			}
			if in.GetCloudSql().Database == "" {
				in.GetCloudSql().Database = "db"
			}
		}
		if in.GetCloudSpanner() != nil {
			if in.GetCloudSpanner().Database == "" {
				in.GetCloudSpanner().Database = "projects/p/instances/i/databases/d"
			}
		}
		if in.GetSpark() != nil {
			if in.GetSpark().GetSparkHistoryServerConfig() != nil {
				if in.GetSpark().GetSparkHistoryServerConfig().DataprocCluster == "" {
					in.GetSpark().GetSparkHistoryServerConfig().DataprocCluster = "projects/p/regions/r/clusters/c"
				}
			}
			if in.GetSpark().GetMetastoreServiceConfig() != nil {
				if in.GetSpark().GetMetastoreServiceConfig().MetastoreService == "" {
					in.GetSpark().GetMetastoreServiceConfig().MetastoreService = "projects/p/locations/l/services/s"
				}
			}
		}
		cleanEmptyMessages(in.ProtoReflect())
	}

	f.FilterStatus = func(in *pb.Connection) {
		cleanEmptyMessages(in.ProtoReflect())
	}

	return f
}

func cleanEmptyMessages(m protoreflect.Message) {
	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if fd.Kind() == protoreflect.MessageKind {
			if fd.IsList() || fd.IsMap() {
				return true
			}
			sub := v.Message()
			cleanEmptyMessages(sub)
			// check if sub has any populated fields now
			hasFields := false
			sub.Range(func(fd2 protoreflect.FieldDescriptor, v2 protoreflect.Value) bool {
				hasFields = true
				return false
			})
			if !hasFields {
				m.Clear(fd)
			}
		}
		return true
	})
}
