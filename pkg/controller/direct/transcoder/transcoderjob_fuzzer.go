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
// proto.message: google.cloud.video.transcoder.v1.Job
// api.group: transcoder.cnrm.cloud.google.com

package transcoder

import (
	pb "cloud.google.com/go/video/transcoder/apiv1/transcoderpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(transcoderJobFuzzer())
}

func transcoderJobFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Job{},
		TranscoderJobSpec_FromProto, TranscoderJobSpec_ToProto,
		TranscoderJobObservedState_FromProto, TranscoderJobObservedState_ToProto,
	)

	// Identity and Labels fields that are not KRM fields
	f.Unimplemented_Identity(".name")
	f.Unimplemented_LabelsAnnotations(".labels")

	// Spec fields
	f.SpecField(".input_uri")
	f.SpecField(".output_uri")
	f.SpecField(".template_id")
	f.SpecField(".config")
	f.SpecField(".ttl_after_completion_days")
	f.SpecField(".mode")
	f.SpecField(".batch_mode_priority")
	f.SpecField(".optimization")
	f.SpecField(".fill_content_gaps")

	// Status/ObservedState fields
	f.StatusField(".state")
	f.StatusField(".create_time")
	f.StatusField(".start_time")
	f.StatusField(".end_time")
	f.StatusField(".error")

	// Unimplemented fields
	f.Unimplemented_NotYetTriaged(".error.details")
	f.Unimplemented_NotYetTriaged(".config.inputs[].attributes.track_definitions[].detected_languages")

	return f
}
