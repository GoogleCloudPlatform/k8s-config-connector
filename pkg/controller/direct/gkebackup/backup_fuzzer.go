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
// proto.message: google.cloud.gkebackup.v1.Backup
// api.group: gkebackup.cnrm.cloud.google.com

package gkebackup

import (
	pb "cloud.google.com/go/gkebackup/apiv1/gkebackuppb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(GKEBackupBackupFuzzer())
}

func GKEBackupBackupFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Backup{},
		GKEBackupBackupSpec_FromProto, GKEBackupBackupSpec_ToProto,
		GKEBackupBackupObservedState_FromProto, GKEBackupBackupObservedState_ToProto,
	)

	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".delete_lock_days")
	f.SpecFields.Insert(".retain_days")
	f.SpecFields.Insert(".description")

	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".manual")
	f.StatusFields.Insert(".delete_lock_expire_time")
	f.StatusFields.Insert(".retain_expire_time")
	f.StatusFields.Insert(".encryption_key")
	f.StatusFields.Insert(".all_namespaces")
	f.StatusFields.Insert(".selected_namespaces")
	f.StatusFields.Insert(".selected_applications")
	f.StatusFields.Insert(".contains_volume_data")
	f.StatusFields.Insert(".contains_secrets")
	f.StatusFields.Insert(".cluster_metadata")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".state_reason")
	f.StatusFields.Insert(".complete_time")
	f.StatusFields.Insert(".resource_count")
	f.StatusFields.Insert(".volume_count")
	f.StatusFields.Insert(".size_bytes")
	f.StatusFields.Insert(".etag")
	f.StatusFields.Insert(".pod_count")
	f.StatusFields.Insert(".config_backup_size_bytes")
	f.StatusFields.Insert(".permissive_mode")

	f.UnimplementedFields.Insert(".name")

	// New fields that could potentially be added
	f.UnimplementedFields.Insert(".satisfies_pzs")
	f.UnimplementedFields.Insert(".satisfies_pzi")

	return f
}
