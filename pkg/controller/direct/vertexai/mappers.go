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

package vertexai

import (
	"encoding/json"

	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
	money "google.golang.org/genproto/googleapis/type/money"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func Featurestore_OnlineServingConfig_FromProto(mapCtx *direct.MapContext, in *pb.Featurestore_OnlineServingConfig) *krmv1alpha1.Featurestore_OnlineServingConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Featurestore_OnlineServingConfig{}
	out.FixedNodeCount = direct.LazyPtr(in.GetFixedNodeCount())
	out.Scaling = Featurestore_OnlineServingConfig_Scaling_FromProto(mapCtx, in.GetScaling())
	return out
}

func Featurestore_OnlineServingConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Featurestore_OnlineServingConfig) *pb.Featurestore_OnlineServingConfig {
	if in == nil {
		return nil
	}
	out := &pb.Featurestore_OnlineServingConfig{}
	out.FixedNodeCount = direct.ValueOf(in.FixedNodeCount)
	out.Scaling = Featurestore_OnlineServingConfig_Scaling_ToProto(mapCtx, in.Scaling)
	return out
}

func Featurestore_OnlineServingConfig_Scaling_FromProto(mapCtx *direct.MapContext, in *pb.Featurestore_OnlineServingConfig_Scaling) *krmv1alpha1.Featurestore_OnlineServingConfig_Scaling {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Featurestore_OnlineServingConfig_Scaling{}
	out.MinNodeCount = direct.LazyPtr(in.GetMinNodeCount())
	out.MaxNodeCount = direct.LazyPtr(in.GetMaxNodeCount())
	out.CPUUtilizationTarget = direct.LazyPtr(in.GetCpuUtilizationTarget())
	return out
}

func Featurestore_OnlineServingConfig_Scaling_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Featurestore_OnlineServingConfig_Scaling) *pb.Featurestore_OnlineServingConfig_Scaling {
	if in == nil {
		return nil
	}
	out := &pb.Featurestore_OnlineServingConfig_Scaling{}
	out.MinNodeCount = direct.ValueOf(in.MinNodeCount)
	out.MaxNodeCount = direct.ValueOf(in.MaxNodeCount)
	out.CpuUtilizationTarget = direct.ValueOf(in.CPUUtilizationTarget)
	return out
}

func EncryptionSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionSpec) *krmv1beta1.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.EncryptionSpec{
		KMSKeyRef: &v1beta1.KMSCryptoKeyRef{
			External: in.KmsKeyName,
		},
	}
	return out
}

func EncryptionSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.EncryptionSpec) *pb.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionSpec{}
	if in.KMSKeyRef != nil {
		out.KmsKeyName = in.KMSKeyRef.External
	}
	return out
}

func EncryptionSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionSpec) *krmv1alpha1.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.EncryptionSpec{
		KMSKeyRef: &v1beta1.KMSCryptoKeyRef{
			External: in.KmsKeyName,
		},
	}
	return out
}

func EncryptionSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.EncryptionSpec) *pb.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionSpec{}
	if in.KMSKeyRef != nil {
		out.KmsKeyName = in.KMSKeyRef.External
	}
	return out
}

func DatasetEncryptionSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionSpec) *krmv1beta1.DatasetEncryptionSpec {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.DatasetEncryptionSpec{}
	if in.KmsKeyName != "" {
		out.KmsKeyNameRef = &v1beta1.KMSCryptoKeyRef{
			External: in.KmsKeyName,
		}
	}
	return out
}

func DatasetEncryptionSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.DatasetEncryptionSpec) *pb.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionSpec{}
	if in.KmsKeyNameRef != nil {
		out.KmsKeyName = in.KmsKeyNameRef.External
	}
	return out
}

// Unversioned wrappers to support old controller code that doesn't expect version suffixes

func VertexAIMetadataStoreSpec_FromProto(mapCtx *direct.MapContext, in *pb.MetadataStore) *krmv1beta1.VertexAIMetadataStoreSpec {
	return VertexAIMetadataStoreSpec_v1beta1_FromProto(mapCtx, in)
}

func VertexAIMetadataStoreSpec_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.VertexAIMetadataStoreSpec) *pb.MetadataStore {
	return VertexAIMetadataStoreSpec_v1beta1_ToProto(mapCtx, in)
}

func VertexAIMetadataStoreObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MetadataStore) *krmv1beta1.VertexAIMetadataStoreObservedState {
	return VertexAIMetadataStoreObservedState_v1beta1_FromProto(mapCtx, in)
}

func VertexAIMetadataStoreObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.VertexAIMetadataStoreObservedState) *pb.MetadataStore {
	return VertexAIMetadataStoreObservedState_v1beta1_ToProto(mapCtx, in)
}

func EncryptionSpecV1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionSpec) *krmv1alpha1.EncryptionSpec {
	return EncryptionSpec_v1alpha1_FromProto(mapCtx, in)
}

func EncryptionSpecV1alpha1_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.EncryptionSpec) *pb.EncryptionSpec {
	return EncryptionSpec_v1alpha1_ToProto(mapCtx, in)
}

func EncryptionSpec_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionSpec) *krmv1beta1.EncryptionSpec {
	return EncryptionSpec_v1beta1_FromProto(mapCtx, in)
}

func EncryptionSpec_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.EncryptionSpec) *pb.EncryptionSpec {
	return EncryptionSpec_v1beta1_ToProto(mapCtx, in)
}

func VertexAIExampleStoreSpec_DisplayName_ToProto(mapCtx *direct.MapContext, in string) string {
	return in
}

func VertexAITensorboardSpec_DisplayName_ToProto(mapCtx *direct.MapContext, in string) string {
	return in
}

func VertexAIDatasetSpec_FromProto(mapCtx *direct.MapContext, in *pb.Dataset) *krmv1beta1.VertexAIDatasetSpec {
	return VertexAIDatasetSpec_v1beta1_FromProto(mapCtx, in)
}

func VertexAIDatasetSpec_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.VertexAIDatasetSpec) *pb.Dataset {
	return VertexAIDatasetSpec_v1beta1_ToProto(mapCtx, in)
}

func VertexAIDatasetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Dataset) *krmv1beta1.VertexAIDatasetObservedState {
	return VertexAIDatasetObservedState_v1beta1_FromProto(mapCtx, in)
}

func VertexAIDatasetObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.VertexAIDatasetObservedState) *pb.Dataset {
	return VertexAIDatasetObservedState_v1beta1_ToProto(mapCtx, in)
}

func Money_v1alpha1_FromProto(mapCtx *direct.MapContext, in *money.Money) *krmv1alpha1.Money {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Money{}
	out.CurrencyCode = direct.LazyPtr(in.GetCurrencyCode())
	out.Units = direct.LazyPtr(in.GetUnits())
	out.Nanos = direct.LazyPtr(in.GetNanos())
	return out
}

func Money_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Money) *money.Money {
	if in == nil {
		return nil
	}
	out := &money.Money{}
	out.CurrencyCode = direct.ValueOf(in.CurrencyCode)
	out.Units = direct.ValueOf(in.Units)
	out.Nanos = direct.ValueOf(in.Nanos)
	return out
}

func JSON_v1alpha1_FromProto(mapCtx *direct.MapContext, in *structpb.Value) *apiextensionsv1.JSON {
	if in == nil {
		return nil
	}
	b, err := protojson.Marshal(in)
	if err != nil {
		mapCtx.Errorf("error marshalling structpb.Value to JSON: %v", err)
		return nil
	}
	out := apiextensionsv1.JSON{Raw: b}
	return &out
}

func JSON_v1alpha1_ToProto(mapCtx *direct.MapContext, in *apiextensionsv1.JSON) *structpb.Value {
	if in == nil {
		return nil
	}
	out := &structpb.Value{}
	if err := protojson.Unmarshal(in.Raw, out); err != nil {
		mapCtx.Errorf("error unmarshalling JSON to structpb.Value: %v", err)
		return nil
	}
	return out
}

func LatLng_v1alpha1_FromProto(mapCtx *direct.MapContext, in *latlng.LatLng) *krmv1alpha1.LatLng {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.LatLng{}
	out.Latitude = direct.LazyPtr(in.GetLatitude())
	out.Longitude = direct.LazyPtr(in.GetLongitude())
	return out
}

func LatLng_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.LatLng) *latlng.LatLng {
	if in == nil {
		return nil
	}
	out := &latlng.LatLng{}
	out.Latitude = direct.ValueOf(in.Latitude)
	out.Longitude = direct.ValueOf(in.Longitude)
	return out
}

func Value_v1alpha1_FromProto(mapCtx *direct.MapContext, in *structpb.Value) *krmv1alpha1.Value {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Value{}
	switch k := in.Kind.(type) {
	case *structpb.Value_NullValue:
		out.NullValue = direct.LazyPtr(k.NullValue.String())
	case *structpb.Value_NumberValue:
		out.NumberValue = direct.LazyPtr(k.NumberValue)
	case *structpb.Value_StringValue:
		out.StringValue = direct.LazyPtr(k.StringValue)
	case *structpb.Value_BoolValue:
		out.BoolValue = direct.LazyPtr(k.BoolValue)
	case *structpb.Value_StructValue:
		b, err := protojson.Marshal(k.StructValue)
		if err != nil {
			mapCtx.Errorf("failed to marshal structpb.Struct: %v", err)
		} else {
			out.StructValue = apiextensionsv1.JSON{Raw: b}
		}
	case *structpb.Value_ListValue:
		// ListValue is disabled due to CRD instability
		// out.ListValue = ListValue_v1alpha1_FromProto(mapCtx, k.ListValue)
	}
	return out
}

func Value_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Value) *structpb.Value {
	if in == nil {
		return nil
	}
	if in.NullValue != nil {
		return structpb.NewNullValue()
	}
	if in.NumberValue != nil {
		return structpb.NewNumberValue(*in.NumberValue)
	}
	if in.StringValue != nil {
		return structpb.NewStringValue(*in.StringValue)
	}
	if in.BoolValue != nil {
		return structpb.NewBoolValue(*in.BoolValue)
	}
	if len(in.StructValue.Raw) > 0 {
		s := &structpb.Struct{}
		if err := protojson.Unmarshal(in.StructValue.Raw, s); err != nil {
			mapCtx.Errorf("failed to unmarshal structpb.Struct: %v", err)
			return nil
		}
		return structpb.NewStructValue(s)
	}
	// ListValue is disabled due to CRD instability
	// if in.ListValue != nil {
	// 	return structpb.NewListValue(ListValue_v1alpha1_ToProto(mapCtx, in.ListValue))
	// }
	return nil
}

func ListValue_v1alpha1_FromProto(mapCtx *direct.MapContext, in *structpb.ListValue) *krmv1alpha1.ListValue {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.ListValue{}
	for _, val := range in.Values {
		if mapped := Value_v1alpha1_FromProto(mapCtx, val); mapped != nil {
			out.Values = append(out.Values, *mapped)
		}
	}
	return out
}

func ListValue_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.ListValue) *structpb.ListValue {
	if in == nil {
		return nil
	}
	out := &structpb.ListValue{}
	for _, val := range in.Values {
		if mapped := Value_v1alpha1_ToProto(mapCtx, &val); mapped != nil {
			out.Values = append(out.Values, mapped)
		}
	}
	return out
}

func CachedContent_UsageMetadata_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.CachedContent_UsageMetadata) *krmv1alpha1.CachedContent_UsageMetadata {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.CachedContent_UsageMetadata{}
	out.TotalTokenCount = direct.LazyPtr(in.GetTotalTokenCount())
	out.TextCount = direct.LazyPtr(in.GetTextCount())
	out.ImageCount = direct.LazyPtr(in.GetImageCount())
	out.VideoDurationSeconds = direct.LazyPtr(in.GetVideoDurationSeconds())
	out.AudioDurationSeconds = direct.LazyPtr(in.GetAudioDurationSeconds())
	return out
}

func CachedContent_UsageMetadata_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.CachedContent_UsageMetadata) *pb.CachedContent_UsageMetadata {
	if in == nil {
		return nil
	}
	out := &pb.CachedContent_UsageMetadata{}
	out.TotalTokenCount = direct.ValueOf(in.TotalTokenCount)
	out.TextCount = direct.ValueOf(in.TextCount)
	out.ImageCount = direct.ValueOf(in.ImageCount)
	out.VideoDurationSeconds = direct.ValueOf(in.VideoDurationSeconds)
	out.AudioDurationSeconds = direct.ValueOf(in.AudioDurationSeconds)
	return out
}

func Schema_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Schema) *krmv1alpha1.Schema {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Schema{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Format = direct.LazyPtr(in.GetFormat())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Nullable = direct.LazyPtr(in.GetNullable())
	out.Default = Value_v1alpha1_FromProto(mapCtx, in.GetDefault())
	if in.GetItems() != nil {
		nested := Schema_v1alpha1_FromProto(mapCtx, in.GetItems())
		if nested != nil {
			b, err := json.Marshal(nested)
			if err != nil {
				mapCtx.Errorf("error marshalling nested schema to JSON: %v", err)
			} else {
				out.Items = &apiextensionsv1.JSON{Raw: b}
			}
		}
	}
	out.MinItems = direct.LazyPtr(in.GetMinItems())
	out.MaxItems = direct.LazyPtr(in.GetMaxItems())
	out.Enum = in.Enum
	out.PropertyOrdering = in.PropertyOrdering
	out.Required = in.Required
	out.MinProperties = direct.LazyPtr(in.GetMinProperties())
	out.MaxProperties = direct.LazyPtr(in.GetMaxProperties())
	out.Minimum = direct.LazyPtr(in.GetMinimum())
	out.Maximum = direct.LazyPtr(in.GetMaximum())
	out.MinLength = direct.LazyPtr(in.GetMinLength())
	out.MaxLength = direct.LazyPtr(in.GetMaxLength())
	out.Pattern = direct.LazyPtr(in.GetPattern())
	out.Example = Value_v1alpha1_FromProto(mapCtx, in.GetExample())
	if in.AnyOf != nil {
		out.AnyOf = make([]apiextensionsv1.JSON, 0, len(in.AnyOf))
		for _, x := range in.AnyOf {
			nested := Schema_v1alpha1_FromProto(mapCtx, x)
			if nested != nil {
				b, err := json.Marshal(nested)
				if err != nil {
					mapCtx.Errorf("error marshalling nested anyOf schema to JSON: %v", err)
				} else {
					out.AnyOf = append(out.AnyOf, apiextensionsv1.JSON{Raw: b})
				}
			}
		}
	}
	out.AdditionalProperties = Value_v1alpha1_FromProto(mapCtx, in.GetAdditionalProperties())
	out.Ref = direct.LazyPtr(in.GetRef())
	return out
}

func Schema_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Schema) *pb.Schema {
	if in == nil {
		return nil
	}
	out := &pb.Schema{}
	out.Type = direct.Enum_ToProto[pb.Type](mapCtx, in.Type)
	out.Format = direct.ValueOf(in.Format)
	out.Title = direct.ValueOf(in.Title)
	out.Description = direct.ValueOf(in.Description)
	out.Nullable = direct.ValueOf(in.Nullable)
	out.Default = Value_v1alpha1_ToProto(mapCtx, in.Default)
	if in.Items != nil {
		var nested krmv1alpha1.Schema
		if err := json.Unmarshal(in.Items.Raw, &nested); err != nil {
			mapCtx.Errorf("error unmarshalling nested schema from JSON: %v", err)
		} else {
			out.Items = Schema_v1alpha1_ToProto(mapCtx, &nested)
		}
	}
	out.MinItems = direct.ValueOf(in.MinItems)
	out.MaxItems = direct.ValueOf(in.MaxItems)
	out.Enum = in.Enum
	out.PropertyOrdering = in.PropertyOrdering
	out.Required = in.Required
	out.MinProperties = direct.ValueOf(in.MinProperties)
	out.MaxProperties = direct.ValueOf(in.MaxProperties)
	out.Minimum = direct.ValueOf(in.Minimum)
	out.Maximum = direct.ValueOf(in.Maximum)
	out.MinLength = direct.ValueOf(in.MinLength)
	out.MaxLength = direct.ValueOf(in.MaxLength)
	out.Pattern = direct.ValueOf(in.Pattern)
	out.Example = Value_v1alpha1_ToProto(mapCtx, in.Example)
	if len(in.AnyOf) > 0 {
		out.AnyOf = make([]*pb.Schema, 0, len(in.AnyOf))
		for _, x := range in.AnyOf {
			var nested krmv1alpha1.Schema
			if err := json.Unmarshal(x.Raw, &nested); err != nil {
				mapCtx.Errorf("error unmarshalling anyOf schema from JSON: %v", err)
			} else {
				out.AnyOf = append(out.AnyOf, Schema_v1alpha1_ToProto(mapCtx, &nested))
			}
		}
	}
	out.AdditionalProperties = Value_v1alpha1_ToProto(mapCtx, in.AdditionalProperties)
	out.Ref = direct.ValueOf(in.Ref)
	return out
}
