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
// proto.message: google.cloud.dataproc.v1.WorkflowTemplate
// api.group: dataproc.cnrm.cloud.google.com

package dataproc

import (
	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dataprocWorkflowTemplateFuzzer())
}

func dataprocWorkflowTemplateFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.WorkflowTemplate{},
		DataprocWorkflowTemplateSpec_FromProto, DataprocWorkflowTemplateSpec_ToProto,
		DataprocWorkflowTemplateObservedState_FromProto, DataprocWorkflowTemplateObservedState_ToProto,
	)

	f.SpecFields.Insert(".id")
	f.SpecFields.Insert(".version")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".placement")
	f.SpecFields.Insert(".jobs")
	f.SpecFields.Insert(".parameters")
	f.SpecFields.Insert(".dag_timeout")
	f.SpecFields.Insert(".encryption_config")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".placement")

	f.UnimplementedFields.Insert(".name") // special field
	f.UnimplementedFields.Insert(".jobs[].hadoop_job.logging_config")
	f.UnimplementedFields.Insert(".jobs[].spark_job.logging_config")
	f.UnimplementedFields.Insert(".jobs[].pyspark_job.logging_config")
	f.UnimplementedFields.Insert(".jobs[].hive_job.logging_config")
	f.UnimplementedFields.Insert(".jobs[].pig_job.logging_config")
	f.UnimplementedFields.Insert(".jobs[].spark_r_job.logging_config")
	f.UnimplementedFields.Insert(".jobs[].spark_sql_job.logging_config")
	f.UnimplementedFields.Insert(".jobs[].presto_job.logging_config")
	f.UnimplementedFields.Insert(".jobs[].trino_job.logging_config")
	f.UnimplementedFields.Insert(".jobs[].flink_job.logging_config")

	return f
}
