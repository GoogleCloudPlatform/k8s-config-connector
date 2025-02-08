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

package contentwarehouse

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/contentwarehouse/apiv1/contentwarehousepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contentwarehouse/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ContentwarehouseDocumentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Document) *krm.ContentwarehouseDocumentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ContentwarehouseDocumentObservedState{}
	// MISSING: Name
	// MISSING: ReferenceID
	// MISSING: DisplayName
	// MISSING: Title
	// MISSING: DisplayURI
	// MISSING: DocumentSchemaName
	// MISSING: PlainText
	// MISSING: CloudAiDocument
	// MISSING: StructuredContentURI
	// MISSING: RawDocumentPath
	// MISSING: InlineRawDocument
	// MISSING: Properties
	// MISSING: UpdateTime
	// MISSING: CreateTime
	// MISSING: RawDocumentFileType
	// MISSING: AsyncEnabled
	// MISSING: ContentCategory
	// MISSING: TextExtractionDisabled
	// MISSING: TextExtractionEnabled
	// MISSING: Creator
	// MISSING: Updater
	// MISSING: DispositionTime
	// MISSING: LegalHold
	return out
}
func ContentwarehouseDocumentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ContentwarehouseDocumentObservedState) *pb.Document {
	if in == nil {
		return nil
	}
	out := &pb.Document{}
	// MISSING: Name
	// MISSING: ReferenceID
	// MISSING: DisplayName
	// MISSING: Title
	// MISSING: DisplayURI
	// MISSING: DocumentSchemaName
	// MISSING: PlainText
	// MISSING: CloudAiDocument
	// MISSING: StructuredContentURI
	// MISSING: RawDocumentPath
	// MISSING: InlineRawDocument
	// MISSING: Properties
	// MISSING: UpdateTime
	// MISSING: CreateTime
	// MISSING: RawDocumentFileType
	// MISSING: AsyncEnabled
	// MISSING: ContentCategory
	// MISSING: TextExtractionDisabled
	// MISSING: TextExtractionEnabled
	// MISSING: Creator
	// MISSING: Updater
	// MISSING: DispositionTime
	// MISSING: LegalHold
	return out
}
func ContentwarehouseDocumentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Document) *krm.ContentwarehouseDocumentSpec {
	if in == nil {
		return nil
	}
	out := &krm.ContentwarehouseDocumentSpec{}
	// MISSING: Name
	// MISSING: ReferenceID
	// MISSING: DisplayName
	// MISSING: Title
	// MISSING: DisplayURI
	// MISSING: DocumentSchemaName
	// MISSING: PlainText
	// MISSING: CloudAiDocument
	// MISSING: StructuredContentURI
	// MISSING: RawDocumentPath
	// MISSING: InlineRawDocument
	// MISSING: Properties
	// MISSING: UpdateTime
	// MISSING: CreateTime
	// MISSING: RawDocumentFileType
	// MISSING: AsyncEnabled
	// MISSING: ContentCategory
	// MISSING: TextExtractionDisabled
	// MISSING: TextExtractionEnabled
	// MISSING: Creator
	// MISSING: Updater
	// MISSING: DispositionTime
	// MISSING: LegalHold
	return out
}
func ContentwarehouseDocumentSpec_ToProto(mapCtx *direct.MapContext, in *krm.ContentwarehouseDocumentSpec) *pb.Document {
	if in == nil {
		return nil
	}
	out := &pb.Document{}
	// MISSING: Name
	// MISSING: ReferenceID
	// MISSING: DisplayName
	// MISSING: Title
	// MISSING: DisplayURI
	// MISSING: DocumentSchemaName
	// MISSING: PlainText
	// MISSING: CloudAiDocument
	// MISSING: StructuredContentURI
	// MISSING: RawDocumentPath
	// MISSING: InlineRawDocument
	// MISSING: Properties
	// MISSING: UpdateTime
	// MISSING: CreateTime
	// MISSING: RawDocumentFileType
	// MISSING: AsyncEnabled
	// MISSING: ContentCategory
	// MISSING: TextExtractionDisabled
	// MISSING: TextExtractionEnabled
	// MISSING: Creator
	// MISSING: Updater
	// MISSING: DispositionTime
	// MISSING: LegalHold
	return out
}
func DateTimeArray_FromProto(mapCtx *direct.MapContext, in *pb.DateTimeArray) *krm.DateTimeArray {
	if in == nil {
		return nil
	}
	out := &krm.DateTimeArray{}
	out.Values = direct.Slice_FromProto(mapCtx, in.Values, DateTime_FromProto)
	return out
}
func DateTimeArray_ToProto(mapCtx *direct.MapContext, in *krm.DateTimeArray) *pb.DateTimeArray {
	if in == nil {
		return nil
	}
	out := &pb.DateTimeArray{}
	out.Values = direct.Slice_ToProto(mapCtx, in.Values, DateTime_ToProto)
	return out
}
func Document_FromProto(mapCtx *direct.MapContext, in *pb.Document) *krm.Document {
	if in == nil {
		return nil
	}
	out := &krm.Document{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ReferenceID = direct.LazyPtr(in.GetReferenceId())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.DisplayURI = direct.LazyPtr(in.GetDisplayUri())
	out.DocumentSchemaName = direct.LazyPtr(in.GetDocumentSchemaName())
	out.PlainText = direct.LazyPtr(in.GetPlainText())
	out.CloudAiDocument = Document_FromProto(mapCtx, in.GetCloudAiDocument())
	out.StructuredContentURI = direct.LazyPtr(in.GetStructuredContentUri())
	out.RawDocumentPath = direct.LazyPtr(in.GetRawDocumentPath())
	out.InlineRawDocument = in.GetInlineRawDocument()
	out.Properties = direct.Slice_FromProto(mapCtx, in.Properties, Property_FromProto)
	// MISSING: UpdateTime
	// MISSING: CreateTime
	out.RawDocumentFileType = direct.Enum_FromProto(mapCtx, in.GetRawDocumentFileType())
	out.AsyncEnabled = direct.LazyPtr(in.GetAsyncEnabled())
	out.ContentCategory = direct.Enum_FromProto(mapCtx, in.GetContentCategory())
	out.TextExtractionDisabled = direct.LazyPtr(in.GetTextExtractionDisabled())
	out.TextExtractionEnabled = direct.LazyPtr(in.GetTextExtractionEnabled())
	out.Creator = direct.LazyPtr(in.GetCreator())
	out.Updater = direct.LazyPtr(in.GetUpdater())
	// MISSING: DispositionTime
	// MISSING: LegalHold
	return out
}
func Document_ToProto(mapCtx *direct.MapContext, in *krm.Document) *pb.Document {
	if in == nil {
		return nil
	}
	out := &pb.Document{}
	out.Name = direct.ValueOf(in.Name)
	out.ReferenceId = direct.ValueOf(in.ReferenceID)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Title = direct.ValueOf(in.Title)
	out.DisplayUri = direct.ValueOf(in.DisplayURI)
	out.DocumentSchemaName = direct.ValueOf(in.DocumentSchemaName)
	if oneof := Document_PlainText_ToProto(mapCtx, in.PlainText); oneof != nil {
		out.StructuredContent = oneof
	}
	if oneof := Document_ToProto(mapCtx, in.CloudAiDocument); oneof != nil {
		out.StructuredContent = &pb.Document_CloudAiDocument{CloudAiDocument: oneof}
	}
	out.StructuredContentUri = direct.ValueOf(in.StructuredContentURI)
	if oneof := Document_RawDocumentPath_ToProto(mapCtx, in.RawDocumentPath); oneof != nil {
		out.RawDocument = oneof
	}
	if oneof := Document_InlineRawDocument_ToProto(mapCtx, in.InlineRawDocument); oneof != nil {
		out.RawDocument = oneof
	}
	out.Properties = direct.Slice_ToProto(mapCtx, in.Properties, Property_ToProto)
	// MISSING: UpdateTime
	// MISSING: CreateTime
	out.RawDocumentFileType = direct.Enum_ToProto[pb.RawDocumentFileType](mapCtx, in.RawDocumentFileType)
	out.AsyncEnabled = direct.ValueOf(in.AsyncEnabled)
	out.ContentCategory = direct.Enum_ToProto[pb.ContentCategory](mapCtx, in.ContentCategory)
	out.TextExtractionDisabled = direct.ValueOf(in.TextExtractionDisabled)
	out.TextExtractionEnabled = direct.ValueOf(in.TextExtractionEnabled)
	out.Creator = direct.ValueOf(in.Creator)
	out.Updater = direct.ValueOf(in.Updater)
	// MISSING: DispositionTime
	// MISSING: LegalHold
	return out
}
func DocumentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Document) *krm.DocumentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DocumentObservedState{}
	// MISSING: Name
	// MISSING: ReferenceID
	// MISSING: DisplayName
	// MISSING: Title
	// MISSING: DisplayURI
	// MISSING: DocumentSchemaName
	// MISSING: PlainText
	// MISSING: CloudAiDocument
	// MISSING: StructuredContentURI
	// MISSING: RawDocumentPath
	// MISSING: InlineRawDocument
	// MISSING: Properties
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: RawDocumentFileType
	// MISSING: AsyncEnabled
	// MISSING: ContentCategory
	// MISSING: TextExtractionDisabled
	// MISSING: TextExtractionEnabled
	// MISSING: Creator
	// MISSING: Updater
	out.DispositionTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDispositionTime())
	out.LegalHold = direct.LazyPtr(in.GetLegalHold())
	return out
}
func DocumentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DocumentObservedState) *pb.Document {
	if in == nil {
		return nil
	}
	out := &pb.Document{}
	// MISSING: Name
	// MISSING: ReferenceID
	// MISSING: DisplayName
	// MISSING: Title
	// MISSING: DisplayURI
	// MISSING: DocumentSchemaName
	// MISSING: PlainText
	// MISSING: CloudAiDocument
	// MISSING: StructuredContentURI
	// MISSING: RawDocumentPath
	// MISSING: InlineRawDocument
	// MISSING: Properties
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: RawDocumentFileType
	// MISSING: AsyncEnabled
	// MISSING: ContentCategory
	// MISSING: TextExtractionDisabled
	// MISSING: TextExtractionEnabled
	// MISSING: Creator
	// MISSING: Updater
	out.DispositionTime = direct.StringTimestamp_ToProto(mapCtx, in.DispositionTime)
	out.LegalHold = direct.ValueOf(in.LegalHold)
	return out
}
func EnumArray_FromProto(mapCtx *direct.MapContext, in *pb.EnumArray) *krm.EnumArray {
	if in == nil {
		return nil
	}
	out := &krm.EnumArray{}
	out.Values = in.Values
	return out
}
func EnumArray_ToProto(mapCtx *direct.MapContext, in *krm.EnumArray) *pb.EnumArray {
	if in == nil {
		return nil
	}
	out := &pb.EnumArray{}
	out.Values = in.Values
	return out
}
func EnumValue_FromProto(mapCtx *direct.MapContext, in *pb.EnumValue) *krm.EnumValue {
	if in == nil {
		return nil
	}
	out := &krm.EnumValue{}
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func EnumValue_ToProto(mapCtx *direct.MapContext, in *krm.EnumValue) *pb.EnumValue {
	if in == nil {
		return nil
	}
	out := &pb.EnumValue{}
	out.Value = direct.ValueOf(in.Value)
	return out
}
func FloatArray_FromProto(mapCtx *direct.MapContext, in *pb.FloatArray) *krm.FloatArray {
	if in == nil {
		return nil
	}
	out := &krm.FloatArray{}
	out.Values = in.Values
	return out
}
func FloatArray_ToProto(mapCtx *direct.MapContext, in *krm.FloatArray) *pb.FloatArray {
	if in == nil {
		return nil
	}
	out := &pb.FloatArray{}
	out.Values = in.Values
	return out
}
func IntegerArray_FromProto(mapCtx *direct.MapContext, in *pb.IntegerArray) *krm.IntegerArray {
	if in == nil {
		return nil
	}
	out := &krm.IntegerArray{}
	out.Values = in.Values
	return out
}
func IntegerArray_ToProto(mapCtx *direct.MapContext, in *krm.IntegerArray) *pb.IntegerArray {
	if in == nil {
		return nil
	}
	out := &pb.IntegerArray{}
	out.Values = in.Values
	return out
}
func MapProperty_FromProto(mapCtx *direct.MapContext, in *pb.MapProperty) *krm.MapProperty {
	if in == nil {
		return nil
	}
	out := &krm.MapProperty{}
	// MISSING: Fields
	return out
}
func MapProperty_ToProto(mapCtx *direct.MapContext, in *krm.MapProperty) *pb.MapProperty {
	if in == nil {
		return nil
	}
	out := &pb.MapProperty{}
	// MISSING: Fields
	return out
}
func Property_FromProto(mapCtx *direct.MapContext, in *pb.Property) *krm.Property {
	if in == nil {
		return nil
	}
	out := &krm.Property{}
	out.Name = direct.LazyPtr(in.GetName())
	out.IntegerValues = IntegerArray_FromProto(mapCtx, in.GetIntegerValues())
	out.FloatValues = FloatArray_FromProto(mapCtx, in.GetFloatValues())
	out.TextValues = TextArray_FromProto(mapCtx, in.GetTextValues())
	out.EnumValues = EnumArray_FromProto(mapCtx, in.GetEnumValues())
	out.PropertyValues = PropertyArray_FromProto(mapCtx, in.GetPropertyValues())
	out.DateTimeValues = DateTimeArray_FromProto(mapCtx, in.GetDateTimeValues())
	out.MapProperty = MapProperty_FromProto(mapCtx, in.GetMapProperty())
	out.TimestampValues = TimestampArray_FromProto(mapCtx, in.GetTimestampValues())
	return out
}
func Property_ToProto(mapCtx *direct.MapContext, in *krm.Property) *pb.Property {
	if in == nil {
		return nil
	}
	out := &pb.Property{}
	out.Name = direct.ValueOf(in.Name)
	if oneof := IntegerArray_ToProto(mapCtx, in.IntegerValues); oneof != nil {
		out.Values = &pb.Property_IntegerValues{IntegerValues: oneof}
	}
	if oneof := FloatArray_ToProto(mapCtx, in.FloatValues); oneof != nil {
		out.Values = &pb.Property_FloatValues{FloatValues: oneof}
	}
	if oneof := TextArray_ToProto(mapCtx, in.TextValues); oneof != nil {
		out.Values = &pb.Property_TextValues{TextValues: oneof}
	}
	if oneof := EnumArray_ToProto(mapCtx, in.EnumValues); oneof != nil {
		out.Values = &pb.Property_EnumValues{EnumValues: oneof}
	}
	if oneof := PropertyArray_ToProto(mapCtx, in.PropertyValues); oneof != nil {
		out.Values = &pb.Property_PropertyValues{PropertyValues: oneof}
	}
	if oneof := DateTimeArray_ToProto(mapCtx, in.DateTimeValues); oneof != nil {
		out.Values = &pb.Property_DateTimeValues{DateTimeValues: oneof}
	}
	if oneof := MapProperty_ToProto(mapCtx, in.MapProperty); oneof != nil {
		out.Values = &pb.Property_MapProperty{MapProperty: oneof}
	}
	if oneof := TimestampArray_ToProto(mapCtx, in.TimestampValues); oneof != nil {
		out.Values = &pb.Property_TimestampValues{TimestampValues: oneof}
	}
	return out
}
func PropertyArray_FromProto(mapCtx *direct.MapContext, in *pb.PropertyArray) *krm.PropertyArray {
	if in == nil {
		return nil
	}
	out := &krm.PropertyArray{}
	out.Properties = direct.Slice_FromProto(mapCtx, in.Properties, Property_FromProto)
	return out
}
func PropertyArray_ToProto(mapCtx *direct.MapContext, in *krm.PropertyArray) *pb.PropertyArray {
	if in == nil {
		return nil
	}
	out := &pb.PropertyArray{}
	out.Properties = direct.Slice_ToProto(mapCtx, in.Properties, Property_ToProto)
	return out
}
func TextArray_FromProto(mapCtx *direct.MapContext, in *pb.TextArray) *krm.TextArray {
	if in == nil {
		return nil
	}
	out := &krm.TextArray{}
	out.Values = in.Values
	return out
}
func TextArray_ToProto(mapCtx *direct.MapContext, in *krm.TextArray) *pb.TextArray {
	if in == nil {
		return nil
	}
	out := &pb.TextArray{}
	out.Values = in.Values
	return out
}
func TimestampArray_FromProto(mapCtx *direct.MapContext, in *pb.TimestampArray) *krm.TimestampArray {
	if in == nil {
		return nil
	}
	out := &krm.TimestampArray{}
	out.Values = direct.Slice_FromProto(mapCtx, in.Values, TimestampValue_FromProto)
	return out
}
func TimestampArray_ToProto(mapCtx *direct.MapContext, in *krm.TimestampArray) *pb.TimestampArray {
	if in == nil {
		return nil
	}
	out := &pb.TimestampArray{}
	out.Values = direct.Slice_ToProto(mapCtx, in.Values, TimestampValue_ToProto)
	return out
}
func TimestampValue_FromProto(mapCtx *direct.MapContext, in *pb.TimestampValue) *krm.TimestampValue {
	if in == nil {
		return nil
	}
	out := &krm.TimestampValue{}
	out.TimestampValue = direct.StringTimestamp_FromProto(mapCtx, in.GetTimestampValue())
	out.TextValue = direct.LazyPtr(in.GetTextValue())
	return out
}
func TimestampValue_ToProto(mapCtx *direct.MapContext, in *krm.TimestampValue) *pb.TimestampValue {
	if in == nil {
		return nil
	}
	out := &pb.TimestampValue{}
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.TimestampValue); oneof != nil {
		out.Value = &pb.TimestampValue_TimestampValue{TimestampValue: oneof}
	}
	if oneof := TimestampValue_TextValue_ToProto(mapCtx, in.TextValue); oneof != nil {
		out.Value = oneof
	}
	return out
}
func Value_FromProto(mapCtx *direct.MapContext, in *pb.Value) *krm.Value {
	if in == nil {
		return nil
	}
	out := &krm.Value{}
	out.FloatValue = direct.LazyPtr(in.GetFloatValue())
	out.IntValue = direct.LazyPtr(in.GetIntValue())
	out.StringValue = direct.LazyPtr(in.GetStringValue())
	out.EnumValue = EnumValue_FromProto(mapCtx, in.GetEnumValue())
	out.DatetimeValue = DateTime_FromProto(mapCtx, in.GetDatetimeValue())
	out.TimestampValue = TimestampValue_FromProto(mapCtx, in.GetTimestampValue())
	out.BooleanValue = direct.LazyPtr(in.GetBooleanValue())
	return out
}
func Value_ToProto(mapCtx *direct.MapContext, in *krm.Value) *pb.Value {
	if in == nil {
		return nil
	}
	out := &pb.Value{}
	if oneof := Value_FloatValue_ToProto(mapCtx, in.FloatValue); oneof != nil {
		out.Kind = oneof
	}
	if oneof := Value_IntValue_ToProto(mapCtx, in.IntValue); oneof != nil {
		out.Kind = oneof
	}
	if oneof := Value_StringValue_ToProto(mapCtx, in.StringValue); oneof != nil {
		out.Kind = oneof
	}
	if oneof := EnumValue_ToProto(mapCtx, in.EnumValue); oneof != nil {
		out.Kind = &pb.Value_EnumValue{EnumValue: oneof}
	}
	if oneof := DateTime_ToProto(mapCtx, in.DatetimeValue); oneof != nil {
		out.Kind = &pb.Value_DatetimeValue{DatetimeValue: oneof}
	}
	if oneof := TimestampValue_ToProto(mapCtx, in.TimestampValue); oneof != nil {
		out.Kind = &pb.Value_TimestampValue{TimestampValue: oneof}
	}
	if oneof := Value_BooleanValue_ToProto(mapCtx, in.BooleanValue); oneof != nil {
		out.Kind = oneof
	}
	return out
}
