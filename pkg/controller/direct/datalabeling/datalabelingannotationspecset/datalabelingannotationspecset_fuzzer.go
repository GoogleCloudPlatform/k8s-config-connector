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
// proto.message: google.cloud.datalabeling.v1beta1.AnnotationSpecSet
// api.group: datalabeling.cnrm.cloud.google.com

package datalabelingannotationspecset

import (
	pb "cloud.google.com/go/datalabeling/apiv1beta1/datalabelingpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(datalabelingAnnotationSpecSetFuzzer())
}

func datalabelingAnnotationSpecSetFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.AnnotationSpecSet{},
		DataLabelingAnnotationSpecSetSpec_FromProto, DataLabelingAnnotationSpecSetSpec_ToProto,
		DataLabelingAnnotationSpecSetObservedState_FromProto, DataLabelingAnnotationSpecSetObservedState_ToProto,
	)

	f.SpecField(".display_name")
	f.SpecField(".description")
	f.SpecField(".annotation_specs")
	f.SpecField(".annotation_specs[].display_name")
	f.SpecField(".annotation_specs[].description")

	f.StatusField(".blocking_resources")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_NotYetTriaged(".annotation_specs[].index")

	return f
}
