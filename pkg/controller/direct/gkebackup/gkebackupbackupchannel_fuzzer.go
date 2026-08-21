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
// proto.message: google.cloud.gkebackup.v1.BackupChannel
// api.group: gkebackup.cnrm.cloud.google.com

package gkebackup

import (
	pb "cloud.google.com/go/gkebackup/apiv1/gkebackuppb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(GKEBackupBackupChannelFuzzer())
}

func GKEBackupBackupChannelFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.BackupChannel{},
		GKEBackupBackupChannelSpec_FromProto, GKEBackupBackupChannelSpec_ToProto,
		GKEBackupBackupChannelObservedState_FromProto, GKEBackupBackupChannelObservedState_ToProto,
	)

	f.SpecField(".destination_project")
	f.SpecField(".labels")
	f.SpecField(".description")

	f.StatusField(".uid")
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".etag")
	f.StatusField(".destination_project_id")

	f.IdentityField(".name")

	return f
}
