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
// proto.message: google.cloud.vectorsearch.v1.Collection
// api.group: vectorsearch.cnrm.cloud.google.com

package vectorsearch

import (
	pb "cloud.google.com/go/vectorsearch/apiv1/vectorsearchpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(vectorSearchCollectionFuzzer())
}

func vectorSearchCollectionFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Collection{},
		VectorSearchCollectionSpec_FromProto, VectorSearchCollectionSpec_ToProto,
		VectorSearchCollectionObservedState_FromProto, VectorSearchCollectionObservedState_ToProto,
	)

	// Identity fields that are not in KRM fields
	f.Unimplemented_Identity(".name")
	f.Unimplemented_NotYetTriaged(".encryption_spec")

	// Spec fields to fuzz
	f.SpecField(".display_name")
	f.SpecField(".description")
	f.SpecField(".labels")
	f.SpecField(".vector_schema")
	f.SpecField(".data_schema")

	// Status fields (ObservedState) to fuzz
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	return f
}
