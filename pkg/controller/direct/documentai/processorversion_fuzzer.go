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
// proto.message: google.cloud.documentai.v1beta3.ProcessorVersion
// api.group: documentai.cnrm.cloud.google.com

package documentai

import (
	pb "cloud.google.com/go/documentai/apiv1beta3/documentaipb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(DocumentAIProcessorVersionFuzzer())
}

func DocumentAIProcessorVersionFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ProcessorVersion{},
		DocumentAIProcessorVersionSpec_FromProto, DocumentAIProcessorVersionSpec_ToProto,
		DocumentAIProcessorVersionObservedState_FromProto, DocumentAIProcessorVersionObservedState_ToProto,
	)

	f.SpecFields.Insert(".name")
	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".kms_key_name")
	f.SpecFields.Insert(".kms_key_version_name")
	f.SpecFields.Insert(".deprecation_info")

	f.StatusFields.Insert(".document_schema")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".latest_evaluation")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".google_managed")
	f.StatusFields.Insert(".model_type")
	f.StatusFields.Insert(".satisfies_pzs")
	f.StatusFields.Insert(".satisfies_pzi")
	f.StatusFields.Insert(".gen_ai_model_info")

	return f
}
