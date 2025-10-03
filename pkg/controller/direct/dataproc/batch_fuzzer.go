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
// proto.message: google.cloud.dataproc.v1.Batch
// api.group: dataproc.cnrm.cloud.google.com

package dataproc

import (
	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dataprocBatchFuzzer())
}

func dataprocBatchFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Batch{},
		DataprocBatchSpec_FromProto, DataprocBatchSpec_ToProto,
		DataprocBatchObservedState_FromProto, DataprocBatchObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")      // special field
	f.UnimplementedFields.Insert(".operation") // only returned in LRO
	f.UnimplementedFields.Insert(".environment_config.execution_config.authentication_config")

	f.SpecFields.Insert(".pyspark_batch")
	f.SpecFields.Insert(".spark_batch")
	f.SpecFields.Insert(".spark_r_batch")
	f.SpecFields.Insert(".spark_sql_batch")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".runtime_config")
	f.SpecFields.Insert(".environment_config")

	f.StatusFields.Insert(".uuid")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".runtime_info")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".state_message")
	f.StatusFields.Insert(".state_time")
	f.StatusFields.Insert(".creator")
	f.StatusFields.Insert(".state_history")

	return f
}
