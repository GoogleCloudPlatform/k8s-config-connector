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
// proto.message: google.cloud.automl.v1.Dataset
// api.group: automl.cnrm.cloud.google.com

package automldataset

import (
	pb "cloud.google.com/go/automl/apiv1/automlpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/automl"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(automlDatasetFuzzer())
}

func automlDatasetFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Dataset{},
		automl.AutoMLDatasetSpec_FromProto, automl.AutoMLDatasetSpec_ToProto,
		automl.AutoMLDatasetObservedState_FromProto, automl.AutoMLDatasetObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // special field

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".translation_dataset_metadata")
	f.SpecFields.Insert(".image_classification_dataset_metadata")
	f.SpecFields.Insert(".text_classification_dataset_metadata")
	f.SpecFields.Insert(".image_object_detection_dataset_metadata")
	f.SpecFields.Insert(".text_extraction_dataset_metadata")
	f.SpecFields.Insert(".text_sentiment_dataset_metadata")

	f.StatusFields.Insert(".example_count")
	f.StatusFields.Insert(".create_time")

	f.StatusFields.Insert(".etag") // we mapped it to status observed state

	return f
}
