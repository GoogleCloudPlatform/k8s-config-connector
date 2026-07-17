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
// proto.message: google.cloud.filestore.v1.Instance
// api.group: filestore.cnrm.cloud.google.com

package filestore

import (
	pb "cloud.google.com/go/filestore/apiv1/filestorepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(filestoreInstanceFuzzer())
}

func filestoreInstanceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Instance{},
		FilestoreInstanceSpec_FromProto, FilestoreInstanceSpec_ToProto,
		FilestoreInstanceStatus_FromProto, FilestoreInstanceStatus_ToProto,
	)

	// Field comparisons:
	// - description maps to .description
	// - fileShares maps to .file_shares
	// - location maps to parent URL/resource name (not direct spec field)
	// - networks maps to .networks
	// - projectRef maps to parent URL/resource name (not direct spec field)
	// - resourceID maps to GCP resource Name (handled by Unimplemented_Identity)
	// - tier maps to .tier

	f.Unimplemented_Identity(".name")

	f.SpecField(".description")
	f.SpecField(".file_shares")
	f.SpecField(".networks")
	f.SpecField(".tier")

	f.StatusField(".create_time")
	f.StatusField(".etag")
	f.StatusField(".state")
	f.StatusField(".status_message")

	f.Unimplemented_LabelsAnnotations(".labels")

	f.Unimplemented_NotYetTriaged(".satisfies_pzs")
	f.Unimplemented_NotYetTriaged(".satisfies_pzi")
	f.Unimplemented_NotYetTriaged(".kms_key_name")
	f.Unimplemented_NotYetTriaged(".suspension_reasons")
	f.Unimplemented_NotYetTriaged(".replication")
	f.Unimplemented_NotYetTriaged(".tags")
	f.Unimplemented_NotYetTriaged(".protocol")
	f.Unimplemented_NotYetTriaged(".custom_performance_supported")
	f.Unimplemented_NotYetTriaged(".performance_config")
	f.Unimplemented_NotYetTriaged(".performance_limits")
	f.Unimplemented_NotYetTriaged(".deletion_protection_enabled")
	f.Unimplemented_NotYetTriaged(".deletion_protection_reason")
	f.Unimplemented_NotYetTriaged(".networks[].connect_mode")

	return f
}
