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
// proto.message: google.cloud.netapp.v1.BackupPolicy
// api.group: netapp.cnrm.cloud.google.com

package netapp

import (
	pb "cloud.google.com/go/netapp/apiv1/netapppb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(netAppBackupPolicyFuzzer())
}

func netAppBackupPolicyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.BackupPolicy{},
		NetAppBackupPolicySpec_FromProto, NetAppBackupPolicySpec_ToProto,
		NetAppBackupPolicyObservedState_FromProto, NetAppBackupPolicyObservedState_ToProto,
	)

	f.SpecFields.Insert(".daily_backup_limit")
	f.SpecFields.Insert(".weekly_backup_limit")
	f.SpecFields.Insert(".monthly_backup_limit")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".enabled")
	f.SpecFields.Insert(".labels")

	f.StatusFields.Insert(".assigned_volume_count")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".state")

	f.UnimplementedFields.Insert(".name") // special field

	return f
}


