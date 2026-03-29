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
// proto.message: google.cloud.modelarmor.v1.Template
// api.group: modelarmor.cnrm.cloud.google.com

package modelarmor

import (
	pb "cloud.google.com/go/modelarmor/apiv1/modelarmorpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(templateFuzzer())
}

func templateFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Template{},
		ModelArmorTemplateSpec_FromProto, ModelArmorTemplateSpec_ToProto,
		ModelArmorTemplateObservedState_FromProto, ModelArmorTemplateObservedState_ToProto,
	)

	f.SpecFields.Insert(".projectRef")
	f.SpecFields.Insert(".location")
	f.SpecFields.Insert(".resourceID")

	f.StatusFields.Insert(".conditions")
	f.StatusFields.Insert(".observedGeneration")
	f.StatusFields.Insert(".externalRef")

	f.Unimplemented_NotYetTriaged(".name")
	f.Unimplemented_NotYetTriaged(".create_time")
	f.Unimplemented_NotYetTriaged(".update_time")
	f.Unimplemented_NotYetTriaged(".filter_config")
	f.Unimplemented_NotYetTriaged(".template_metadata")
	f.Unimplemented_NotYetTriaged(".labels")

	return f
}
