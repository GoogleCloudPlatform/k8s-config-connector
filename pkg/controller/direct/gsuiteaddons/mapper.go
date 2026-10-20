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

package gsuiteaddons

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gsuiteaddons/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

func Value_ToProto(mapCtx *direct.MapContext, in *krm.Value) *structpb.Value {
	if in == nil {
		return nil
	}
	out := &structpb.Value{}
	if oneof := Value_NullValue_ToProto(mapCtx, in.NullValue); oneof != nil {
		out.Kind = oneof
	}
	if oneof := Value_NumberValue_ToProto(mapCtx, in.NumberValue); oneof != nil {
		out.Kind = oneof
	}
	if oneof := Value_StringValue_ToProto(mapCtx, in.StringValue); oneof != nil {
		out.Kind = oneof
	}
	if oneof := Value_BoolValue_ToProto(mapCtx, in.BoolValue); oneof != nil {
		out.Kind = oneof
	}
	if oneof := direct.Struct_ToProto(mapCtx, &in.StructValue); oneof != nil {
		out.Kind = &structpb.Value_StructValue{StructValue: oneof}
	}
	if oneof := ListValue_ToProto(mapCtx, in.ListValue); oneof != nil {
		out.Kind = &structpb.Value_ListValue{ListValue: oneof}
	}
	return out
}
