// Copyright 2024 Google LLC
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
// proto.message: google.cloud.datacatalog.v1.Tag
// api.group: datacatalog.cnrm.cloud.google.com

package datacatalog

import (
	pb "cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/datacatalog/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataCatalogTagSpec_Template_ToProto(mapCtx *direct.MapContext, in *krm.TagTemplateRef) string {
	if in == nil {
		return ""
	}

	return in.External
}

func DataCatalogTagSpec_Column_ToProto(mapCtx *direct.MapContext, in *string) *pb.Tag_Column {
	if in == nil {
		return nil
	}

	return &pb.Tag_Column{
		Column: direct.ValueOf(in),
	}
}

func TagField_DoubleValue_ToProto(mapCtx *direct.MapContext, in *string) *pb.TagField_DoubleValue {
	if in == nil {
		return nil
	}

	val := direct.StringToFloat64(mapCtx, direct.ValueOf(in))

	return &pb.TagField_DoubleValue{DoubleValue: val}
}

func TagField_StringValue_ToProto(mapCtx *direct.MapContext, in *string) *pb.TagField_StringValue {
	if in == nil {
		return nil
	}

	return &pb.TagField_StringValue{StringValue: direct.ValueOf(in)}
}

func TagField_BoolValue_ToProto(mapCtx *direct.MapContext, in *bool) *pb.TagField_BoolValue {
	if in == nil {
		return nil
	}

	return &pb.TagField_BoolValue{BoolValue: direct.ValueOf(in)}
}

func TagField_RichtextValue_ToProto(mapCtx *direct.MapContext, in *string) *pb.TagField_RichtextValue {
	if in == nil {
		return nil
	}

	return &pb.TagField_RichtextValue{RichtextValue: direct.ValueOf(in)}
}

func init() {
	fuzztesting.RegisterKRMFuzzer(dataCatalogTagFuzzer())
}

func dataCatalogTagFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Tag{},
		DataCatalogTagSpec_FromProto, DataCatalogTagSpec_ToProto,
		DataCatalogTagObservedState_FromProto, DataCatalogTagObservedState_ToProto,
	)

	f.SpecFields.Insert(".template")
	f.SpecFields.Insert(".column")
	f.SpecFields.Insert(".fields")

	f.StatusFields.Insert(".template_display_name")
	f.StatusFields.Insert(".dataplex_transfer_status")
	f.StatusFields.Insert(".fields")

	f.UnimplementedFields.Insert(".name") // identifier, output only

	return f
}

```
</out>


