// Copyright 2025 Google LLC
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
// proto.message: google.cloud.clouddms.v1.MigrationJob
// api.group: clouddms.cnrm.cloud.google.com

package clouddms

import (
	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dmsMigrationJobFuzzer())
}

func dmsMigrationJobFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.MigrationJob{},
		CloudDMSMigrationJobSpec_FromProto, CloudDMSMigrationJobSpec_ToProto,
		CloudDMSMigrationJobObservedState_FromProto, CloudDMSMigrationJobObservedState_ToProto,
	)

	f.SpecField(".conversion_workspace")
	f.SpecField(".destination")
	f.SpecField(".display_name")
	f.SpecField(".source")
	f.SpecField(".staticIPConnectivity")
	f.SpecField(".type")

	f.StatusField(".create_time")
	f.StatusField(".duration")
	f.StatusField(".error")
	f.StatusField(".phase")
	f.StatusField(".update_time")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_Etag()
	f.Unimplemented_LabelsAnnotations(".labels")
	// The 'state' is not considered as observedState because it does not have the OUTPUT-ONLY field mask.
	f.Unimplemented_Internal(".state")

	return f
}
