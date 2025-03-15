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
// proto.message: google.cloud.gkebackup.v1.RestorePlan
// api.group: gkebackup.cnrm.cloud.google.com

package gkebackup

import (
	pb "cloud.google.com/go/gkebackup/apiv1/gkebackuppb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(GKEBackupRestorePlanFuzzer())
}

func GKEBackupRestorePlanFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.RestorePlan{},
		GKEBackupRestorePlanSpec_FromProto, GKEBackupRestorePlanSpec_ToProto,
		GKEBackupRestorePlanObservedState_FromProto, GKEBackupRestorePlanObservedState_ToProto,
	)
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".backup_plan")
	f.SpecFields.Insert(".cluster")
	f.SpecFields.Insert(".restore_config")
	f.SpecFields.Insert(".labels")

	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".etag")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".state_reason")

	f.UnimplementedFields.Insert(".name")

	// The default value of `.restore_config.volume_data_restore_policy_bindings.volume_type` is
	// VolumeTypeEnum_VOLUME_TYPE_UNSPECIFIED, which does not roundtrip due to our Enum_FromProto implementation.
	f.UnimplementedFields.Insert(".restore_config.volume_data_restore_policy_bindings")

	return f
}
