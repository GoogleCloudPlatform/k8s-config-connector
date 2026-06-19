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

package apihubattribute

import (
	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(apiHubAttributeFuzzer())
}

func apiHubAttributeFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Attribute{},
		APIHubAttributeSpec_FromProto, APIHubAttributeSpec_ToProto,
		APIHubAttributeObservedState_FromProto, APIHubAttributeObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // special field

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".scope")
	f.SpecFields.Insert(".data_type")
	f.SpecFields.Insert(".allowed_values")
	f.SpecFields.Insert(".cardinality")

	f.StatusFields.Insert(".definition_type")
	f.StatusFields.Insert(".mandatory")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")

	return f
}
