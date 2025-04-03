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

// +generated:mapper
// krm.group: datacatalog.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.datacatalog.v1

package datacatalog

import (
	pb "cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datacatalog/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataCatalogTagSpec_FromProto(mapCtx *direct.MapContext, in *pb.Tag) *krmv1alpha1.DataCatalogTagSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DataCatalogTagSpec{}
	// MISSING: Name
	out.Template = in.GetTemplate()
	out.Column = direct.LazyPtr(in.GetColumn())
	out.Fields = make(map[string]krmv1alpha1.TagField)
	for k, v := range in.GetFields() {
		out.Fields[k] = *TagField_FromProto(mapCtx, v)
	}
	return out
}
func DataCatalogTagSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DataCatalogTagSpec) *pb.Tag {
	if in == nil {
		return nil
	}
	out := &pb.Tag{}
	// MISSING: Name
	out.Template = in.Template
	if val := direct.ValueOf(in.Column); val != "" {
		out.Scope = &pb.Tag_Column{Column: val}
	}
	out.Fields = make(map[string]*pb.TagField)
	for k, v := range in.Fields {
		out.Fields[k] = TagField_ToProto(mapCtx, &v)
	}
	return out
}

func TagField_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.TagField) *pb.TagField {
	if in == nil {
		return nil
	}
	out := &pb.TagField{}
	// MISSING: DisplayName
	if in.DoubleValue != nil {
		out.Kind = &pb.TagField_DoubleValue{DoubleValue: *in.DoubleValue}
	}
	if in.StringValue != nil {
		out.Kind = &pb.TagField_StringValue{StringValue: *in.StringValue}
	}
	if in.BoolValue != nil {
		out.Kind = &pb.TagField_BoolValue{BoolValue: *in.BoolValue}
	}
	if in.TimestampValue != nil {
		out.Kind = &pb.TagField_TimestampValue{TimestampValue: direct.StringTimestamp_ToProto(mapCtx, in.TimestampValue)}
	}
	if in.EnumValue != nil {
		out.Kind = &pb.TagField_EnumValue_{EnumValue: TagField_EnumValue_ToProto(mapCtx, in.EnumValue)}
	}
	if in.RichtextValue != nil {
		out.Kind = &pb.TagField_RichtextValue{RichtextValue: *in.RichtextValue}
	}
	// MISSING: Order
	return out
}
