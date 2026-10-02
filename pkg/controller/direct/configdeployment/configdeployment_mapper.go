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

package configdeployment

import (
	pb "cloud.google.com/go/config/apiv1/configpb"
	"google.golang.org/protobuf/types/known/structpb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/configdeployment/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// mapper.generated.go calls these helpers for the ConfigDeployment fields that
// generate-mapper cannot convert yet.

// ProviderConfig_SourceType_ToProto converts ProviderConfig.source_type, a
// proto3 optional enum. generate-mapper reads the field as an enum but writes
// it as if it were a oneof.
func ProviderConfig_SourceType_ToProto(mapCtx *direct.MapContext, in *string) *pb.ProviderConfig_ProviderSource {
	if in == nil {
		return nil
	}
	v := direct.Enum_ToProto[pb.ProviderConfig_ProviderSource](mapCtx, in)
	return &v
}

// Value_FromProto converts a google.protobuf.Value.
func Value_FromProto(mapCtx *direct.MapContext, in *structpb.Value) *krm.Value {
	if in == nil {
		return nil
	}
	out := &krm.Value{}
	switch kind := in.GetKind().(type) {
	case *structpb.Value_NullValue:
		out.NullValue = direct.LazyPtr(kind.NullValue.String())
	case *structpb.Value_NumberValue:
		out.NumberValue = direct.PtrTo(kind.NumberValue)
	case *structpb.Value_StringValue:
		out.StringValue = direct.PtrTo(kind.StringValue)
	case *structpb.Value_BoolValue:
		out.BoolValue = direct.PtrTo(kind.BoolValue)
	case *structpb.Value_StructValue:
		if v := direct.Struct_FromProto(mapCtx, kind.StructValue); v != nil {
			out.StructValue = *v
		}
	case *structpb.Value_ListValue:
		out.ListValue = &krm.ListValue{}
		for _, v := range kind.ListValue.GetValues() {
			if c := Value_FromProto(mapCtx, v); c != nil {
				out.ListValue.Values = append(out.ListValue.Values, *c)
			}
		}
	}
	return out
}

// Value_ToProto converts a google.protobuf.Value.
func Value_ToProto(mapCtx *direct.MapContext, in *krm.Value) *structpb.Value {
	if in == nil {
		return nil
	}
	switch {
	case in.NullValue != nil:
		return structpb.NewNullValue()
	case in.NumberValue != nil:
		return structpb.NewNumberValue(*in.NumberValue)
	case in.StringValue != nil:
		return structpb.NewStringValue(*in.StringValue)
	case in.BoolValue != nil:
		return structpb.NewBoolValue(*in.BoolValue)
	case len(in.StructValue.Raw) > 0:
		return structpb.NewStructValue(direct.Struct_ToProto(mapCtx, &in.StructValue))
	case in.ListValue != nil:
		list := &structpb.ListValue{}
		for i := range in.ListValue.Values {
			list.Values = append(list.Values, Value_ToProto(mapCtx, &in.ListValue.Values[i]))
		}
		return structpb.NewListValue(list)
	}
	return &structpb.Value{}
}
