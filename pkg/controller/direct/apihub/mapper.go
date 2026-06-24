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

package apihub

import (
	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apihub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AttributeValues_FromProto(mapCtx *direct.MapContext, in *pb.AttributeValues) *krm.AttributeValues {
	if in == nil {
		return nil
	}
	out := &krm.AttributeValues{}
	out.EnumValues = AttributeValues_EnumAttributeValues_FromProto(mapCtx, in.GetEnumValues())
	out.StringValues = AttributeValues_StringAttributeValues_FromProto(mapCtx, in.GetStringValues())
	out.JSONValues = AttributeValues_StringAttributeValues_FromProto(mapCtx, in.GetJsonValues())
	return out
}

func AttributeValues_ToProto(mapCtx *direct.MapContext, in *krm.AttributeValues) *pb.AttributeValues {
	if in == nil {
		return nil
	}
	out := &pb.AttributeValues{}
	if oneof := AttributeValues_EnumAttributeValues_ToProto(mapCtx, in.EnumValues); oneof != nil {
		out.Value = &pb.AttributeValues_EnumValues{EnumValues: oneof}
	}
	if oneof := AttributeValues_StringAttributeValues_ToProto(mapCtx, in.StringValues); oneof != nil {
		out.Value = &pb.AttributeValues_StringValues{StringValues: oneof}
	}
	if oneof := AttributeValues_StringAttributeValues_ToProto(mapCtx, in.JSONValues); oneof != nil {
		out.Value = &pb.AttributeValues_JsonValues{JsonValues: oneof}
	}
	return out
}
