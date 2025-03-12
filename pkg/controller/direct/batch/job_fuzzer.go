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
// proto.message: google.cloud.bigquery.datatransfer.v1.TransferConfig
// api.group: bigquerydatatransfer.cnrm.cloud.google.com

package batch

import (
	pb "cloud.google.com/go/batch/apiv1/batchpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(batchJobFuzzer())
}

func batchJobFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Job{},
		BatchJobSpec_FromProto, BatchJobSpec_ToProto,
		BatchJobObservedState_FromProto, BatchJobObservedState_ToProto,
	)

	f.SpecFields.Insert(".priority")
	f.SpecFields.Insert(".task_groups")
	f.SpecFields.Insert(".allocation_policy")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".logs_policy")
	f.SpecFields.Insert(".notifications")
	f.SpecFields.Insert(".parent")

	f.StatusFields.Insert(".name")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".task_groups")
	f.StatusFields.Insert(".status")

	return f
}
