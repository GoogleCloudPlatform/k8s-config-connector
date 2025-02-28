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
// proto.message: google.cloud.bigquery.biglake.v1.Table
// api.group: bigquerybiglake.cnrm.cloud.google.com

package bigquerybiglake

import (
	pb "cloud.google.com/go/bigquery/biglake/apiv1/biglakepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(bigLakeTableFuzzer())
}

func bigLakeTableFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Table{},
		BigLakeTableSpec_FromProto, BigLakeTableSpec_ToProto,
		BigLakeTableObservedState_FromProto, BigLakeTableObservedState_ToProto,
	)

	f.SpecFields.Insert(".type")
	f.SpecFields.Insert(".hive_options")
	f.SpecFields.Insert(".etag")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".delete_time")
	f.StatusFields.Insert(".expire_time")

	f.UnimplementedFields.Insert(".name")

	return f
}
