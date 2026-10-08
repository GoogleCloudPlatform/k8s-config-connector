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
// proto.message: google.cloud.dataplex.v1.DataTaxonomy
// api.group: dataplex.cnrm.cloud.google.com

package dataplex

import (
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dataplexDataTaxonomyFuzzer())
}

func dataplexDataTaxonomyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.DataTaxonomy{},
		DataplexDataTaxonomySpec_FromProto, DataplexDataTaxonomySpec_ToProto,
		DataplexDataTaxonomyObservedState_FromProto, DataplexDataTaxonomyObservedState_ToProto,
	)

	f.SpecField(".display_name")
	f.SpecField(".description")

	f.StatusField(".uid")
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".attribute_count")
	f.StatusField(".class_count")

	f.UnimplementedFields.Insert(".name")
	f.Unimplemented_LabelsAnnotations(".labels")
	f.UnimplementedFields.Insert(".etag")

	return f
}
