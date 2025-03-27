// Copyright 2024 Google LLC
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
// proto.message: google.cloud.datacatalog.v1.Entry
// api.group: datacatalog.cnrm.cloud.google.com

package datacatalog

import (
	pb "cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dataCatalogEntryFuzzer())
}

func dataCatalogEntryFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Entry{},
		DataCatalogEntrySpec_FromProto, DataCatalogEntrySpec_ToProto,
		DataCatalogEntryObservedState_FromProto, DataCatalogEntryObservedState_ToProto,
	)

	f.SpecFields.Insert(".linked_resource")
	f.SpecFields.Insert(".fully_qualified_name")
	f.SpecFields.Insert(".type")
	f.SpecFields.Insert(".user_specified_type")
	f.SpecFields.Insert(".user_specified_system")
	f.SpecFields.Insert(".sql_database_system_spec")
	f.SpecFields.Insert(".looker_system_spec")
	f.SpecFields.Insert(".cloud_bigtable_system_spec")
	f.SpecFields.Insert(".gcs_fileset_spec")
	f.SpecFields.Insert(".database_table_spec")
	f.SpecFields.Insert(".data_source_connection_spec")
	f.SpecFields.Insert(".routine_spec")
	f.SpecFields.Insert(".dataset_spec")
	f.SpecFields.Insert(".fileset_spec")
	f.SpecFields.Insert(".service_spec")
	f.SpecFields.Insert(".model_spec")
	f.SpecFields.Insert(".feature_online_store_spec")
	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".business_context")
	f.SpecFields.Insert(".schema")
	f.SpecFields.Insert(".source_system_timestamps")
	f.SpecFields.Insert(".usage_signal")
	f.SpecFields.Insert(".labels")

	f.StatusFields.Insert(".gcs_fileset_spec")
	f.StatusFields.Insert(".bigquery_table_spec")
	f.StatusFields.Insert(".bigquery_date_sharded_spec")
	f.StatusFields.Insert(".database_table_spec")
	f.StatusFields.Insert(".feature_online_store_spec")
	f.StatusFields.Insert(".usage_signal")
	f.StatusFields.Insert(".data_source")
	f.StatusFields.Insert(".personal_details")

	f.UnimplementedFields.Insert(".name") // special field
	f.UnimplementedFields.Insert(".integrated_system")

	return f
}
