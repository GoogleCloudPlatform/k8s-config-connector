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
// proto.message: google.cloud.bigquery.biglake.v1.Catalog
// api.group: bigquerybiglake.cnrm.cloud.google.com

package bigquerybiglake

import (
	pb "cloud.google.com/go/bigquery/biglake/apiv1/biglakepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(bigLakeDatabaseFuzzer())
}

func bigLakeDatabaseFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Database{},
		BigLakeDatabaseSpec_FromProto, BigLakeDatabaseSpec_ToProto,
		BigLakeDatabaseObservedState_FromProto, BigLakeDatabaseObservedState_ToProto,
	)

	f.SpecFields.Insert(".hive_options")
	f.SpecFields.Insert(".type")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".delete_time")
	f.StatusFields.Insert(".expire_time")

	f.UnimplementedFields.Insert(".name")

	return f
}
