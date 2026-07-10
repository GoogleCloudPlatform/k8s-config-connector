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
// proto.message: google.iam.admin.v1.ServiceAccount
// krm.group: iam.cnrm.cloud.google.com
// krm.kind: IAMServiceAccount

package iam

import (
	pb "cloud.google.com/go/iam/admin/apiv1/adminpb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(iamServiceAccountFuzzer())
}

func iamServiceAccountFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ServiceAccount{},
		IAMServiceAccountSpec_FromProto, IAMServiceAccountSpec_ToProto,
		IAMServiceAccountStatus_FromProto, IAMServiceAccountStatus_ToProto,
	)

	// KRM Spec field comparison:
	// - Spec.Description maps to .description
	// - Spec.Disabled maps to .disabled
	// - Spec.DisplayName maps to .display_name
	// - Spec.ResourceID is KCC-specific resource ID and not mapped to a direct proto field
	f.SpecField(".description")
	f.SpecField(".disabled")
	f.SpecField(".display_name")

	// KRM Status field comparison:
	// - Status.Email maps to .email
	// - Status.Member is derived as "serviceAccount:{email}" in FromProto, not stored directly as a field in pb.ServiceAccount
	// - Status.Name maps to .name
	// - Status.UniqueId maps to .unique_id
	// - Status.ExternalRef maps to .name
	f.StatusField(".email")
	f.StatusField(".name")
	f.StatusField(".unique_id")

	// Identity/GCP Path fields
	f.Unimplemented_Identity(".project_id")

	// Fields not yet mapped or supported in KRM
	f.Unimplemented_NotYetTriaged(".etag")
	f.Unimplemented_NotYetTriaged(".oauth2_client_id")

	return f
}
