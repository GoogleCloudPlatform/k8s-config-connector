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
// proto.message: google.cloud.dataproc.v1.Job
// api.group: dataproc.cnrm.cloud.google.com

package dataproc

import (
	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dataprocJobFuzzer())
}

func dataprocJobFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Job{},
		DataprocJobSpec_FromProto, DataprocJobSpec_ToProto,
		DataprocJobObservedState_FromProto, DataprocJobObservedState_ToProto,
	)

	f.SpecFields.Insert(".reference")
	f.SpecFields.Insert(".placement")
	f.SpecFields.Insert(".hadoop_job")
	f.SpecFields.Insert(".spark_job")
	f.SpecFields.Insert(".pyspark_job")
	f.SpecFields.Insert(".hive_job")
	f.SpecFields.Insert(".pig_job")
	f.SpecFields.Insert(".spark_r_job")
	f.SpecFields.Insert(".spark_sql_job")
	f.SpecFields.Insert(".presto_job")
	f.SpecFields.Insert(".trino_job")
	f.SpecFields.Insert(".flink_job")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".scheduling")
	f.SpecFields.Insert(".driver_scheduling_config")

	f.StatusFields.Insert(".placement")
	f.StatusFields.Insert(".status")
	f.StatusFields.Insert(".yarn_applications")
	f.StatusFields.Insert(".driver_output_resource_uri")
	f.StatusFields.Insert(".driver_control_files_uri")
	f.StatusFields.Insert(".job_uuid")
	f.StatusFields.Insert(".done")

	f.UnimplementedFields.Insert(".status")
	f.UnimplementedFields.Insert(".status_history")

	return f
}
