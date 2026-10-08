// Copyright 2026 Google LLC
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
// proto.message: google.cloud.aiplatform.v1.SpecialistPool
// api.group: aiplatform.cnrm.cloud.google.com

package aiplatform

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(aiPlatformSpecialistPoolFuzzer())
}

func aiPlatformSpecialistPoolFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.SpecialistPool{},
		AIPlatformSpecialistPoolSpec_FromProto, AIPlatformSpecialistPoolSpec_ToProto,
		AIPlatformSpecialistPoolObservedState_FromProto, AIPlatformSpecialistPoolObservedState_ToProto,
	)

	f.Unimplemented_Identity(".name")

	f.SpecField(".display_name")
	f.SpecField(".specialist_manager_emails")
	f.SpecField(".specialist_worker_emails")

	f.StatusField(".specialist_managers_count")
	f.StatusField(".pending_data_labeling_jobs")

	return f
}
