// Copyright 2026 Google LLC
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
// proto.message: google.cloud.datalabeling.v1beta1.Instruction
// api.group: datalabeling.cnrm.cloud.google.com

package datalabelinginstruction

import (
	pb "cloud.google.com/go/datalabeling/apiv1beta1/datalabelingpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(datalabelingInstructionFuzzer())
}

func datalabelingInstructionFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Instruction{},
		DataLabelingInstructionSpec_FromProto, DataLabelingInstructionSpec_ToProto,
		DataLabelingInstructionObservedState_FromProto, DataLabelingInstructionObservedState_ToProto,
	)

	f.SpecField(".display_name")
	f.SpecField(".description")
	f.SpecField(".data_type")
	f.SpecField(".csv_instruction")
	f.SpecField(".csv_instruction.gcs_file_uri")
	f.SpecField(".pdf_instruction")
	f.SpecField(".pdf_instruction.gcs_file_uri")

	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".blocking_resources")

	f.Unimplemented_Identity(".name")

	return f
}
