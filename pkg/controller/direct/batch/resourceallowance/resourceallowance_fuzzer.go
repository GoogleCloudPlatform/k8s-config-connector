// Copyright 2026 Google LLC
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
// proto.message: google.cloud.batch.v1alpha.ResourceAllowance
// api.group: batch.cnrm.cloud.google.com

package resourceallowance

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpclients/generated/google/cloud/batch/v1alpha"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(batchResourceAllowanceFuzzer())
}

func batchResourceAllowanceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ResourceAllowance{},
		CloudBatchResourceAllowanceSpec_FromProto, CloudBatchResourceAllowanceSpec_ToProto,
		CloudBatchResourceAllowanceObservedState_FromProto, CloudBatchResourceAllowanceObservedState_ToProto,
	)

	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".notifications")
	f.SpecFields.Insert(".usage_resource_allowance")

	f.StatusFields.Insert(".name")
	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".usage_resource_allowance")

	return f
}
