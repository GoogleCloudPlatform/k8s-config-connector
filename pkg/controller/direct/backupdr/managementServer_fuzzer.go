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
// proto.message: google.cloud.backupdr.v1.ManagementServer
// api.group: backupdr.cnrm.cloud.google.com

package backupdr

import (
	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(backupDRManagementServerFuzzer())
}

func backupDRManagementServerFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ManagementServer{},
		BackupDRManagementServerSpec_FromProto, BackupDRManagementServerSpec_ToProto,
		BackupDRManagementServerObservedState_FromProto, BackupDRManagementServerObservedState_ToProto,
	)

	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".type")
	f.SpecFields.Insert(".networks")
	f.SpecFields.Insert(".etag")

	f.StatusFields.Insert(".management_uri")
	f.StatusFields.Insert(".workforce_identity_based_management_uri")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".oauth2_client_id")
	f.StatusFields.Insert(".workforce_identity_based_oauth2_client_id")
	f.StatusFields.Insert(".ba_proxy_uri")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".satisfies_pzs")
	f.StatusFields.Insert(".satisfies_pzi")

	f.UnimplementedFields.Insert(".name") // special field

	return f
}
