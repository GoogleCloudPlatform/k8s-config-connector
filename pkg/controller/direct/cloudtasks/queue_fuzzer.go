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
// proto.message: google.cloud.tasks.v2.Queue
// api.group: tasks.cnrm.cloud.google.com

package cloudtasks

import (
	pb "cloud.google.com/go/cloudtasks/apiv2/cloudtaskspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(tasksQueueFuzzer())
}

func tasksQueueFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Queue{},
		TasksQueueSpec_FromProto, TasksQueueSpec_ToProto,
		TasksQueueObservedState_FromProto, TasksQueueObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // special field

	f.SpecFields.Insert(".app_engine_routing_override")
	f.SpecFields.Insert(".rate_limits")
	f.SpecFields.Insert(".retry_config")
	f.SpecFields.Insert(".stackdriver_logging_config")

	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".purge_time")

	return f
}
