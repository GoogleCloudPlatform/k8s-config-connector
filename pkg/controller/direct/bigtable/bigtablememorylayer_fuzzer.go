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
// proto.message: google.bigtable.admin.v2.MemoryLayer
// api.group: bigtable.cnrm.cloud.google.com

package bigtable

import (
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(BigtableMemoryLayerFuzzer())
}

func BigtableMemoryLayerFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.MemoryLayer{},
		BigtableMemoryLayerSpec_v1alpha1_FromProto, BigtableMemoryLayerSpec_v1alpha1_ToProto,
		BigtableMemoryLayerObservedState_v1alpha1_FromProto, BigtableMemoryLayerObservedState_v1alpha1_ToProto,
	)

	// A memory layer has no settable fields today: memory_config is output only
	// apart from its presence, which the KRM API models as the existence of the
	// BigtableMemoryLayer object.
	f.StatusFields.Insert(".memory_config")
	f.StatusFields.Insert(".state")

	// name is the resource identity, modelled as spec.clusterRef.
	f.UnimplementedFields.Insert(".name")
	// etag is not declarative-friendly, so it is not surfaced in the KRM API.
	f.UnimplementedFields.Insert(".etag")
	return f
}
