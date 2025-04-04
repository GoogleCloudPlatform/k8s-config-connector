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
// proto.message: google.cloud.gkebackup.v1.Restore
// api.group: gkebackup.cnrm.cloud.google.com

package gkebackup

import (
	pb "cloud.google.com/go/gkebackup/apiv1/gkebackuppb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(gkeBackupRestoreFuzzer())
}

func gkeBackupRestoreFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Restore{},
		GKEBackupRestoreSpec_FromProto, GKEBackupRestoreSpec_ToProto,
		GKEBackupRestoreObservedState_FromProto, GKEBackupRestoreObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // special field

	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".backup")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".filter")
	f.SpecFields.Insert(".volume_data_restore_policy_overrides")

	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".cluster")
	f.StatusFields.Insert(".restore_config")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".state_reason")
	f.StatusFields.Insert(".complete_time")
	f.StatusFields.Insert(".resources_restored_count")
	f.StatusFields.Insert(".resources_excluded_count")
	f.StatusFields.Insert(".resources_failed_count")
	f.StatusFields.Insert(".volumes_restored_count")
	f.StatusFields.Insert(".etag")

	// The default value of `.restore_config.volume_data_restore_policy_bindings.volume_type` is
	// VolumeTypeEnum_VOLUME_TYPE_UNSPECIFIED, which does not roundtrip due to our Enum_FromProto implementation.
	f.UnimplementedFields.Insert(".restore_config.volume_data_restore_policy_bindings")

	return f
}
