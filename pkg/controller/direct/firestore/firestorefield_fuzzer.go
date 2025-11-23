// Copyright 2025 Google LLC
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
// proto.message: google.firestore.admin.v1.Field
// api.group: firestore.cnrm.cloud.google.com

package firestore

import (
	pb "cloud.google.com/go/firestore/apiv1/admin/adminpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(firestoreFieldFuzzer())
}

func firestoreFieldFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Field{},
		FirestoreFieldSpec_v1alpha1_FromProto, FirestoreFieldSpec_v1alpha1_ToProto,
		FirestoreFieldObservedState_v1alpha1_FromProto, FirestoreFieldObservedState_v1alpha1_ToProto,
	)
	f.FilterSpec = func(in *pb.Field) {
		for _, index := range in.GetIndexConfig().GetIndexes() {
			for _, field := range index.GetFields() {
				if x, ok := field.GetValueMode().(*pb.Index_IndexField_ArrayConfig_); ok {
					// If we specify an ArrayConfig, it should not be the unspecified value.
					if x.ArrayConfig == pb.Index_IndexField_ARRAY_CONFIG_UNSPECIFIED {
						x.ArrayConfig = pb.Index_IndexField_CONTAINS
					}
				}
			}
		}
	}

	f.IdentityField(".name")

	f.SpecField(".index_config")
	f.StatusField(".index_config.ancestor_field")
	f.StatusField(".index_config.reverting")
	f.StatusField(".index_config.uses_ancestor_config")
	f.StatusField(".index_config.indexes[].name")
	f.StatusField(".index_config.indexes[].state")
	f.StatusField(".ttl_config.state")

	f.Unimplemented_NotYetTriaged(".ttl_config")

	return f
}
