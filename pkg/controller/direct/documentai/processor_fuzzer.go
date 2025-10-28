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
// proto.message: google.cloud.documentai.v1.Processor
// api.group: documentai.cnrm.cloud.google.com

package documentai

import (
	pb "cloud.google.com/go/documentai/apiv1/documentaipb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(DocumentAIProcessorFuzzer())
}

func DocumentAIProcessorFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Processor{},
		DocumentAIProcessorSpec_v1alpha1_FromProto, DocumentAIProcessorSpec_v1alpha1_ToProto,
		DocumentAIProcessorObservedState_v1alpha1_FromProto, DocumentAIProcessorObservedState_v1alpha1_ToProto,
	)
	f.UnimplementedFields.Insert(".name")

	f.SpecFields.Insert(".type")
	f.SpecFields.Insert(".display_name")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".processor_version_aliases")
	f.StatusFields.Insert(".process_endpoint")
	f.StatusFields.Insert(".default_processor_version")

	f.UnimplementedFields.Insert(".kms_key_name")
	f.UnimplementedFields.Insert(".satisfies_pzi")
	f.UnimplementedFields.Insert(".satisfies_pzs")

	return f
}
