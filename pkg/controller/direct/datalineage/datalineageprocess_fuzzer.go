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

package datalineage

import (
	pb "cloud.google.com/go/datacatalog/lineage/apiv1/lineagepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datalineage/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	"google.golang.org/protobuf/types/known/structpb"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(datalineageProcessFuzzer())
}

func datalineageProcessFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer[*pb.Process, krm.DataLineageProcessSpec, krm.DataLineageProcessStatus](&pb.Process{},
		DataLineageProcessSpec_FromProto, DataLineageProcessSpec_ToProto,
		nil, nil,
	)

	// Identity Field
	f.Unimplemented_Identity(".name")

	// Spec Fields
	f.SpecField(".display_name")
	f.SpecField(".origin")
	f.SpecField(".attributes")

	f.FilterSpec = func(in *pb.Process) {
		for _, v := range in.Attributes {
			normalizeValue(v)
		}
	}

	return f
}

func normalizeValue(v *structpb.Value) {
	if v == nil {
		return
	}
	if v.Kind == nil {
		v.Kind = &structpb.Value_NullValue{NullValue: structpb.NullValue_NULL_VALUE}
		return
	}
	switch k := v.Kind.(type) {
	case *structpb.Value_StructValue:
		if k.StructValue != nil {
			for _, val := range k.StructValue.Fields {
				normalizeValue(val)
			}
		}
	case *structpb.Value_ListValue:
		if k.ListValue != nil {
			for _, val := range k.ListValue.Values {
				normalizeValue(val)
			}
		}
	}
}
