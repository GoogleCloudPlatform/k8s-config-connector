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
// proto.message: google.cloud.osconfig.v1beta.GuestPolicy
// api.group: osconfig.cnrm.cloud.google.com

package osconfig

import (
	osconfigpb "cloud.google.com/go/osconfig/apiv1beta/osconfigpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/osconfig/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(osConfigGuestPolicyFuzzer())
}

func osConfigGuestPolicyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer[*osconfigpb.GuestPolicy, krm.OSConfigGuestPolicySpec, krm.OSConfigGuestPolicyStatus](&osconfigpb.GuestPolicy{},
		OSConfigGuestPolicySpec_FromProto, OSConfigGuestPolicySpec_ToProto,
		OSConfigGuestPolicyStatus_FromProto, OSConfigGuestPolicyStatus_ToProto,
	)

	// Field comparison of KRM OSConfigGuestPolicySpec vs Proto:
	// - ResourceID (KRM Spec) maps to metadata.name / GCP resource name (.name is identity)
	// - Description (KRM Spec) maps to description
	// - Assignment (KRM Spec) maps to assignment
	// - Packages (KRM Spec) maps to packages
	// - PackageRepositories (KRM Spec) maps to package_repositories
	// - Recipes (KRM Spec) maps to recipes

	// Spec fields
	f.SpecField(".description")
	f.SpecField(".assignment")
	f.SpecField(".packages")
	f.SpecField(".package_repositories")
	f.SpecField(".recipes")

	// Status fields
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".etag")

	// Identity fields
	f.Unimplemented_Identity(".name")

	return f
}
