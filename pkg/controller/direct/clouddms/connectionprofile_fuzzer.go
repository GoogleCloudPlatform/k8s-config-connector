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
// proto.message: google.cloud.clouddms.v1.ConnectionProfile
// api.group: clouddms.cnrm.cloud.google.com

package clouddms

import (
	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(clouddmsConnectionProfileFuzzer())
}

func clouddmsConnectionProfileFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ConnectionProfile{},
		CloudDMSConnectionProfileSpec_FromProto, CloudDMSConnectionProfileSpec_ToProto,
		CloudDMSConnectionProfileObservedState_FromProto, CloudDMSConnectionProfileObservedState_ToProto,
	)

	f.SpecField(".displayName")
	f.SpecField(".mysql")
	f.SpecField(".postgresql")
	f.SpecField(".oracle")
	f.SpecField(".cloudsql")
	f.SpecField(".alloydb")
	f.SpecField(".provider")

	f.StatusField(".state")
	f.StatusField(".createTime")
	f.StatusField(".updateTime")
	f.StatusField(".error")

	f.Unimplemented_LabelsAnnotations(".labels") // labels can be either Spec or Status. The fuzzer will complain if it is not handled.
	f.IdentityField(".name")                     // The resource name is derived from metadata.name
	return f
}
