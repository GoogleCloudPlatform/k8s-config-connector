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
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/cloudidentity/groups/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(cloudIdentityGroupFuzzer())
}

func cloudIdentityGroupFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Group{},
		CloudIdentityGroupSpec_FromProto, CloudIdentityGroupSpec_ToProto,
		CloudIdentityGroupStatus_FromProto, CloudIdentityGroupStatus_ToProto,
	)

	f.UnimplementedFields.Insert(".posix_groups")
	f.UnimplementedFields.Insert(".dynamic_group_metadata")

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".group_key")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".parent")

	f.StatusFields.Insert(".name")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".additional_group_keys")

	return f
}
