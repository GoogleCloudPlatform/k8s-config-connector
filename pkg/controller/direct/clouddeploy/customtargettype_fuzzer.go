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
// proto.message: google.cloud.deploy.v1.CustomTargetType
// api.group: clouddeploy.googleapis.com

package clouddeploy

import (
	pb "cloud.google.com/go/deploy/apiv1/deploypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(customTargetTypeFuzzer())
}

func customTargetTypeFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.CustomTargetType{},
		CustomTargetType_FromProto, CustomTargetType_ToProto,
		CustomTargetTypeObservedState_FromProto, CustomTargetTypeObservedState_ToProto,
	)

	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".custom_actions")

	f.StatusFields.Insert(".custom_target_type_id")
	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")

	f.UnimplementedFields.Insert(".name")        // special field
	f.UnimplementedFields.Insert(".annotations") // NOT YET
	f.UnimplementedFields.Insert(".labels")      // NOT YET
	f.UnimplementedFields.Insert(".etag")        // NOT YET
	return f
}
