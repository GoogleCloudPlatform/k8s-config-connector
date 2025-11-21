// Copyright 2025 Google LLC
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
// proto.message: google.cloud.assuredworkloads.v1.Workload
// api.group: assuredworkloads.cnrm.cloud.google.com

package assuredworkloads

import (
	pb "cloud.google.com/go/assuredworkloads/apiv1/assuredworkloadspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(assuredWorkloadsWorkloadFuzzer())
}

func assuredWorkloadsWorkloadFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Workload{},
		AssuredWorkloadsWorkloadSpec_FromProto, AssuredWorkloadsWorkloadSpec_ToProto,
		AssuredWorkloadsWorkloadObservedState_FromProto, AssuredWorkloadsWorkloadObservedState_ToProto,
	)

	f.SpecField(".billing_account")
	f.SpecField(".resource_settings")

	f.StatusField(".create_time")
	f.StatusField(".resources")
	f.StatusField(".kaj_enrollment_state")
	f.StatusField(".saa_enrollment_response")
	f.StatusField(".compliant_but_disallowed_services")
	f.Unimplemented_NotYetTriaged(".provisioned_resources_parent")
	f.Unimplemented_NotYetTriaged(".kms_settings")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_Etag()
	f.Unimplemented_LabelsAnnotations(".labels")

	return f
}
