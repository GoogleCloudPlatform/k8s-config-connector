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
// proto.message: google.iam.v3.PolicyBinding
// api.group: iam.cnrm.cloud.google.com

package iam

import (
	pb "cloud.google.com/go/iam/apiv3/iampb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(IAMPolicyBindingFuzzer())
}

func IAMPolicyBindingFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.PolicyBinding{},
		IAMPolicyBindingSpec_FromProto, IAMPolicyBindingSpec_ToProto,
		IAMPolicyBindingObservedState_FromProto, IAMPolicyBindingObservedState_ToProto,
	)

	f.SpecFields.Insert(".etag")
	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".annotations")
	f.SpecFields.Insert(".target")
	f.SpecFields.Insert(".policy_kind")
	f.SpecFields.Insert(".policy")
	f.SpecFields.Insert(".condition")

	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".policy_uid")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")

	f.UnimplementedFields.Insert(".name") // special field

	return f
}
