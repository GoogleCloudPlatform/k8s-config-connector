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

package contentwarehouse

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contentwarehouse/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	pb "google.golang.org/genproto/googleapis/cloud/contentwarehouse/v1"
	iam "google.golang.org/genproto/googleapis/iam/v1"
)

// Manual mapper for google.iam.v1.Policy since it's not in the service package
func Policy_FromProto(mapCtx *direct.MapContext, in *iam.Policy) *krm.Policy {
	return nil
}

func Policy_ToProto(mapCtx *direct.MapContext, in *krm.Policy) *iam.Policy {
	return nil
}

// Manual mapper for SchemaSource because it's missing from the Go SDK
func PropertyDefinition_SchemaSource_FromProto(mapCtx *direct.MapContext, in any) *krm.PropertyDefinition_SchemaSource {
	return nil
}

func PropertyDefinition_SchemaSource_ToProto(mapCtx *direct.MapContext, in *krm.PropertyDefinition_SchemaSource) any {
	return nil
}

// Manual mapper for PropertyDefinition because the Go SDK is missing RetrievalImportance and SchemaSources
func PropertyDefinition_FromProto(mapCtx *direct.MapContext, in *pb.PropertyDefinition) *krm.PropertyDefinition {
	if in == nil {
		return nil
	}
	out := &krm.PropertyDefinition{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.IsRepeatable = direct.LazyPtr(in.GetIsRepeatable())
	out.IsFilterable = direct.LazyPtr(in.GetIsFilterable())
	out.IsSearchable = direct.LazyPtr(in.GetIsSearchable())
	out.IsMetadata = direct.LazyPtr(in.GetIsMetadata())
	out.IsRequired = direct.LazyPtr(in.GetIsRequired())
	out.IntegerTypeOptions = IntegerTypeOptions_FromProto(mapCtx, in.GetIntegerTypeOptions())
	out.FloatTypeOptions = FloatTypeOptions_FromProto(mapCtx, in.GetFloatTypeOptions())
	out.TextTypeOptions = TextTypeOptions_FromProto(mapCtx, in.GetTextTypeOptions())
	out.PropertyTypeOptions = PropertyTypeOptions_FromProto(mapCtx, in.GetPropertyTypeOptions())
	out.EnumTypeOptions = EnumTypeOptions_FromProto(mapCtx, in.GetEnumTypeOptions())
	out.DateTimeTypeOptions = DateTimeTypeOptions_FromProto(mapCtx, in.GetDateTimeTypeOptions())
	out.MapTypeOptions = MapTypeOptions_FromProto(mapCtx, in.GetMapTypeOptions())
	out.TimestampTypeOptions = TimestampTypeOptions_FromProto(mapCtx, in.GetTimestampTypeOptions())
	return out
}

func PropertyDefinition_ToProto(mapCtx *direct.MapContext, in *krm.PropertyDefinition) *pb.PropertyDefinition {
	if in == nil {
		return nil
	}
	out := &pb.PropertyDefinition{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.IsRepeatable = direct.ValueOf(in.IsRepeatable)
	out.IsFilterable = direct.ValueOf(in.IsFilterable)
	out.IsSearchable = direct.ValueOf(in.IsSearchable)
	out.IsMetadata = direct.ValueOf(in.IsMetadata)
	out.IsRequired = direct.ValueOf(in.IsRequired)
	if oneof := IntegerTypeOptions_ToProto(mapCtx, in.IntegerTypeOptions); oneof != nil {
		out.ValueTypeOptions = &pb.PropertyDefinition_IntegerTypeOptions{IntegerTypeOptions: oneof}
	}
	if oneof := FloatTypeOptions_ToProto(mapCtx, in.FloatTypeOptions); oneof != nil {
		out.ValueTypeOptions = &pb.PropertyDefinition_FloatTypeOptions{FloatTypeOptions: oneof}
	}
	if oneof := TextTypeOptions_ToProto(mapCtx, in.TextTypeOptions); oneof != nil {
		out.ValueTypeOptions = &pb.PropertyDefinition_TextTypeOptions{TextTypeOptions: oneof}
	}
	if oneof := PropertyTypeOptions_ToProto(mapCtx, in.PropertyTypeOptions); oneof != nil {
		out.ValueTypeOptions = &pb.PropertyDefinition_PropertyTypeOptions{PropertyTypeOptions: oneof}
	}
	if oneof := EnumTypeOptions_ToProto(mapCtx, in.EnumTypeOptions); oneof != nil {
		out.ValueTypeOptions = &pb.PropertyDefinition_EnumTypeOptions{EnumTypeOptions: oneof}
	}
	if oneof := DateTimeTypeOptions_ToProto(mapCtx, in.DateTimeTypeOptions); oneof != nil {
		out.ValueTypeOptions = &pb.PropertyDefinition_DateTimeTypeOptions{DateTimeTypeOptions: oneof}
	}
	if oneof := MapTypeOptions_ToProto(mapCtx, in.MapTypeOptions); oneof != nil {
		out.ValueTypeOptions = &pb.PropertyDefinition_MapTypeOptions{MapTypeOptions: oneof}
	}
	if oneof := TimestampTypeOptions_ToProto(mapCtx, in.TimestampTypeOptions); oneof != nil {
		out.ValueTypeOptions = &pb.PropertyDefinition_TimestampTypeOptions{TimestampTypeOptions: oneof}
	}
	return out
}
