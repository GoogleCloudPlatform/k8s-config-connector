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
// proto.message: google.iam.v2.Policy
// krm.group: iam.cnrm.cloud.google.com
// krm.kind: IAMDenyPolicy

package iam

import (
	pb "cloud.google.com/go/iam/apiv2/iampb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(denyPolicyFuzzer())
}

func denyPolicyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Policy{},
		IAMDenyPolicySpec_FromProto, IAMDenyPolicySpec_ToProto,
		IAMDenyPolicyObservedState_FromProto, IAMDenyPolicyObservedState_ToProto,
	)

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".rules")

	// Annotations/labels - need a roadmap
	f.UnimplementedFields.Insert(".annotations")

	// Fields that we don't yet know if we should support
	f.UnimplementedFields.Insert(".managing_authority")

	// Output fields that we don't currently see a need for
	f.UnimplementedFields.Insert(".uid")
	f.UnimplementedFields.Insert(".etag")
	f.UnimplementedFields.Insert(".create_time")
	f.UnimplementedFields.Insert(".update_time")
	f.UnimplementedFields.Insert(".delete_time")

	// Identity fields that are handled specially
	f.UnimplementedFields.Insert(".name")

	// System fields that aren't relevant
	f.UnimplementedFields.Insert(".kind")

	return f
}
