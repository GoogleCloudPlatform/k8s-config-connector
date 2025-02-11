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
// proto.message: google.cloud.alloydb.v1beta.Cluster

package alloydb

import (
	pb "cloud.google.com/go/alloydb/apiv1beta/alloydbpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(alloyDBClusterFuzzer())
}

func alloyDBClusterFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Cluster{},
		AlloyDBClusterSpec_FromProto, AlloyDBClusterSpec_ToProto,
		AlloyDBClusterStatus_FromProto, AlloyDBClusterStatus_ToProto,
	)

	f.UnimplementedFields.Insert(".annotations")
	f.UnimplementedFields.Insert(".etag")
	f.UnimplementedFields.Insert(".labels")
	f.UnimplementedFields.Insert(".gemini_config")
	f.UnimplementedFields.Insert(".tags")
	f.UnimplementedFields.Insert(".ssl_config")
	f.UnimplementedFields.Insert(".cloudsql_backup_run_source")
	f.UnimplementedFields.Insert(".delete_time")
	f.UnimplementedFields.Insert(".psc_config")
	f.UnimplementedFields.Insert(".maintenance_schedule")
	f.UnimplementedFields.Insert(".create_time")
	f.UnimplementedFields.Insert(".primary_config")
	f.UnimplementedFields.Insert(".satisfies_pzs")
	f.UnimplementedFields.Insert(".state")
	f.UnimplementedFields.Insert(".trial_metadata")
	f.UnimplementedFields.Insert(".subscription_type")
	f.UnimplementedFields.Insert(".backup_source.backup_uid")
	f.UnimplementedFields.Insert(".update_time")
	f.UnimplementedFields.Insert(".reconciling")
	// .initial_user.password is unreadable.
	f.UnimplementedFields.Insert(".initial_user.password")

	f.SpecFields.Insert(".deletion_policy")
	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".cluster_type")
	f.SpecFields.Insert(".network")
	f.SpecFields.Insert(".network_config")
	f.SpecFields.Insert(".automated_backup_policy")
	f.SpecFields.Insert(".initial_user")
	f.SpecFields.Insert(".encryption_config")
	f.SpecFields.Insert(".maintenance_update_policy")
	f.SpecFields.Insert(".restore_backup_source")
	f.SpecFields.Insert(".restore_continuous_backup_source")
	f.SpecFields.Insert(".secondary_config")
	f.SpecFields.Insert(".continuous_backup_config")

	f.StatusFields.Insert(".name")
	f.StatusFields.Insert(".database_version")
	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".observed_state")
	f.StatusFields.Insert(".encryption_info")
	f.StatusFields.Insert(".continuous_backup_info")
	f.StatusFields.Insert(".migration_source")
	f.StatusFields.Insert(".backup_source")

	return f
}
