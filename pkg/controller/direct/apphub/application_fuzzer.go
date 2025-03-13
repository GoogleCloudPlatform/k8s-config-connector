// Copyright 2024 Google LLC
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
// proto.message: google.cloud.apphub.v1.Application
// api.group: apphub.cnrm.cloud.google.com

package apphub

import (
	pb "cloud.google.com/go/apphub/apiv1/apphubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(apphubApplicationFuzzer())
}

func apphubApplicationFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Application{},
		AppHubApplicationSpec_FromProto, AppHubApplicationSpec_ToProto,
		AppHubApplicationStatus_FromProto, AppHubApplicationStatus_ToProto,
	)
	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".attributes")
	f.SpecFields.Insert(".scope")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".state")

	f.UnimplementedFields.Insert(".name")
	return f
}
