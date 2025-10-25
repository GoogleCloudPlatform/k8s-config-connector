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
// proto.message: google.firestore.v1.Document
// api.group: firestore.cnrm.cloud.google.com

package firestore

import (
	pb "cloud.google.com/go/firestore/apiv1/firestorepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(firestoreDocumentFuzzer())
}

func firestoreDocumentFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Document{},
		FirestoreDocumentSpec_v1alpha1_FromProto, FirestoreDocumentSpec_v1alpha1_ToProto,
		FirestoreDocumentObservedState_v1alpha1_FromProto, FirestoreDocumentObservedState_v1alpha1_ToProto,
	)
	f.FilterSpec = func(in *pb.Document) {
		for _, field := range in.GetFields() {
			removeUnsupportedFieldValues(field)
		}
	}

	f.IdentityField(".name")

	f.SpecField(".fields")

	f.StatusField(".create_time")
	f.StatusField(".update_time")

	return f
}

// removeUnsupportedFieldValues removes field values that do not easily
// round-trip to JSON, replacing them with null values.
// These values are not currently supported in KCC FirestoreDocument resources.
func removeUnsupportedFieldValues(v *pb.Value) {
	if v == nil {
		return
	}
	switch value := v.ValueType.(type) {
	case *pb.Value_MapValue:
		for _, fv := range value.MapValue.Fields {
			removeUnsupportedFieldValues(fv)
		}
	case *pb.Value_ArrayValue:
		for _, av := range value.ArrayValue.Values {
			removeUnsupportedFieldValues(av)
		}

	// These types do not easily round-trip to JSON, so we omit them for now.
	case *pb.Value_BytesValue, *pb.Value_TimestampValue, *pb.Value_ReferenceValue, *pb.Value_GeoPointValue:
		v.ValueType = &pb.Value_NullValue{NullValue: structpb.NullValue_NULL_VALUE}
	}
}
