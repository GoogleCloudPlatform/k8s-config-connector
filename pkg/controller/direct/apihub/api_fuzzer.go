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

// +tool:fuzz-gen
// proto.message: google.cloud.apihub.v1.Api

package apihub

import (
	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(apiHubAPIFuzzer())
}

func apiHubAPIFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Api{},
		APIHubAPISpec_FromProto, APIHubAPISpec_ToProto,
		APIHubAPIObservedState_FromProto, APIHubAPIObservedState_ToProto,
	)

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".documentation")
	f.SpecFields.Insert(".owner")
	f.SpecFields.Insert(".target_user")
	f.SpecFields.Insert(".team")
	f.SpecFields.Insert(".business_unit")
	f.SpecFields.Insert(".maturity_level")
	f.SpecFields.Insert(".api_style")
	f.SpecFields.Insert(".selected_version")

	f.StatusFields.Insert(".name")
	f.StatusFields.Insert(".versions")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".target_user.attribute")

	f.UnimplementedFields.Insert(".attributes")

	return f
}
