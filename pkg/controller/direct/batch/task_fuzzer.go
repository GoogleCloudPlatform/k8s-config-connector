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
// proto.message: google.cloud.batch.v1.Task
// api.group: batch.cnrm.cloud.google.com

package batch

import (
	pb "cloud.google.com/go/batch/apiv1/batchpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(batchTaskFuzzer())
}

func batchTaskFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Task{},
		BatchTaskSpec_FromProto, BatchTaskSpec_ToProto,
		BatchTaskObservedState_FromProto, BatchTaskObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")

	f.StatusFields.Insert(".status")

	return f
}
