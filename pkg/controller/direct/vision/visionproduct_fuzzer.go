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
// proto.message: google.cloud.vision.v1.Product
// api.group: vision.cnrm.cloud.google.com

package vision

import (
	pb "cloud.google.com/go/vision/v2/apiv1/visionpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(visionProductFuzzer())
}

func visionProductFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Product{},
		VisionProductSpec_FromProto, VisionProductSpec_ToProto,
		VisionProductObservedState_FromProto, VisionProductObservedState_ToProto,
	)

	// Identity fields that are not in KRM fields
	f.Unimplemented_Identity(".name")

	// Spec fields to fuzz
	f.SpecField(".display_name")
	f.SpecField(".description")
	f.SpecField(".product_category")
	f.SpecField(".product_labels")

	return f
}
