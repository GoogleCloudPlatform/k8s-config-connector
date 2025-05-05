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

package cloudidentity

import (
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/apps/cloudidentity/groups/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(cloudIdentityMembershipFuzzer())
}

func cloudIdentityMembershipFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Membership{},
		CloudIdentityMembershipSpec_FromProto, CloudIdentityMembershipSpec_ToProto,
		CloudIdentityMembershipStatus_FromProto, CloudIdentityMembershipStatus_ToProto,
	)

	f.SpecFields.Insert(".member_key")
	f.SpecFields.Insert(".preferred_member_key")
	f.SpecFields.Insert(".roles")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".delivery_setting")
	f.StatusFields.Insert(".type")

	f.UnimplementedFields.Insert(".name")

	return f
}
