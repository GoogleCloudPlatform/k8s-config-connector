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

package vertexai

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
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

func Presets_Query_ToProto(mapCtx *direct.MapContext, in *string) *pb.Presets_Query {
	if in == nil {
		return nil
	}
	val := direct.Enum_ToProto[pb.Presets_Query](mapCtx, in)
	return &val
}

func Value_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Value) *structpb.Value {
	if in == nil {
		return nil
	}
	out := &structpb.Value{}
	if in.BoolValue != nil {
		out.Kind = &structpb.Value_BoolValue{
			BoolValue: *in.BoolValue,
		}
	} else if in.NumberValue != nil {
		out.Kind = &structpb.Value_NumberValue{
			NumberValue: *in.NumberValue,
		}
	} else if in.StringValue != nil {
		out.Kind = &structpb.Value_StringValue{
			StringValue: *in.StringValue,
		}
	} else if in.NullValue != nil {
		strVal := *in.NullValue
		var value int32
		if val, ok := structpb.NullValue_value[strVal]; ok {
			value = val
		}
		out.Kind = &structpb.Value_NullValue{
			NullValue: structpb.NullValue(value),
		}
	} else if len(in.StructValue.Raw) > 0 {
		s := direct.Struct_ToProto(mapCtx, &in.StructValue)
		out.Kind = &structpb.Value_StructValue{
			StructValue: s,
		}
	} else if in.ListValue != nil {
		out.Kind = &structpb.Value_ListValue{
			ListValue: ListValue_v1alpha1_ToProto(mapCtx, in.ListValue),
		}
	}
	return out
}

func Value_v1alpha1_FromProto(mapCtx *direct.MapContext, in *structpb.Value) *krmv1alpha1.Value {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Value{}
	switch kind := in.GetKind().(type) {
	case *structpb.Value_StringValue:
		value := kind.StringValue
		out.StringValue = &value
	case *structpb.Value_NumberValue:
		value := kind.NumberValue
		out.NumberValue = &value
	case *structpb.Value_NullValue:
		value := kind.NullValue.String()
		out.NullValue = &value
	case *structpb.Value_BoolValue:
		value := kind.BoolValue
		out.BoolValue = &value
	case *structpb.Value_StructValue:
		js := direct.Struct_FromProto(mapCtx, kind.StructValue)
		if js != nil {
			out.StructValue = *js
		}
	case *structpb.Value_ListValue:
		out.ListValue = ListValue_v1alpha1_FromProto(mapCtx, kind.ListValue)
	}
	return out
}

func ListValue_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.ListValue) *structpb.ListValue {
	if in == nil {
		return nil
	}
	out := &structpb.ListValue{}
	for _, v := range in.Values {
		out.Values = append(out.Values, Value_v1alpha1_ToProto(mapCtx, &v))
	}
	return out
}

func ListValue_v1alpha1_FromProto(mapCtx *direct.MapContext, in *structpb.ListValue) *krmv1alpha1.ListValue {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.ListValue{}
	for _, v := range in.GetValues() {
		val := Value_v1alpha1_FromProto(mapCtx, v)
		if val != nil {
			out.Values = append(out.Values, *val)
		}
	}
	return out
}

func ExplanationMetadata_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.ExplanationMetadata) *krmv1alpha1.ExplanationMetadata {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.ExplanationMetadata{}
	out.FeatureAttributionsSchemaURI = direct.LazyPtr(in.GetFeatureAttributionsSchemaUri())
	out.LatentSpaceSource = direct.LazyPtr(in.GetLatentSpaceSource())
	return out
}

func ExplanationMetadata_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.ExplanationMetadata) *pb.ExplanationMetadata {
	if in == nil {
		return nil
	}
	out := &pb.ExplanationMetadata{}
	out.FeatureAttributionsSchemaUri = direct.ValueOf(in.FeatureAttributionsSchemaURI)
	out.LatentSpaceSource = direct.ValueOf(in.LatentSpaceSource)
	return out
}

func VertexAIDataLabelingJobSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.DataLabelingJob) *krmv1alpha1.VertexAIDataLabelingJobSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.VertexAIDataLabelingJobSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	if v := in.GetDatasets(); len(v) != 0 {
		for i := range v {
			out.DatasetRefs = append(out.DatasetRefs, krmv1beta1.VertexAIDatasetRef{External: v[i]})
		}
	}
	out.AnnotationLabels = in.GetAnnotationLabels()
	out.LabelerCount = direct.LazyPtr(in.GetLabelerCount())
	out.InstructionURI = direct.LazyPtr(in.GetInstructionUri())
	out.InputsSchemaURI = direct.LazyPtr(in.GetInputsSchemaUri())
	out.Inputs = JSON_v1alpha1_FromProto(mapCtx, in.GetInputs())
	out.Labels = in.GetLabels()
	out.SpecialistPools = in.GetSpecialistPools()
	out.EncryptionSpec = EncryptionSpec_v1alpha1_FromProto(mapCtx, in.GetEncryptionSpec())
	out.ActiveLearningConfig = ActiveLearningConfig_v1alpha1_FromProto(mapCtx, in.GetActiveLearningConfig())
	return out
}

func VertexAIDataLabelingJobSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.VertexAIDataLabelingJobSpec) *pb.DataLabelingJob {
	if in == nil {
		return nil
	}
	out := &pb.DataLabelingJob{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	if v := in.DatasetRefs; len(v) != 0 {
		for i := range v {
			out.Datasets = append(out.Datasets, v[i].External)
		}
	}
	out.AnnotationLabels = in.AnnotationLabels
	out.LabelerCount = direct.ValueOf(in.LabelerCount)
	out.InstructionUri = direct.ValueOf(in.InstructionURI)
	out.InputsSchemaUri = direct.ValueOf(in.InputsSchemaURI)
	out.Inputs = JSON_v1alpha1_ToProto(mapCtx, in.Inputs)
	out.Labels = in.Labels
	out.SpecialistPools = in.SpecialistPools
	out.EncryptionSpec = EncryptionSpec_v1alpha1_ToProto(mapCtx, in.EncryptionSpec)
	out.ActiveLearningConfig = ActiveLearningConfig_v1alpha1_ToProto(mapCtx, in.ActiveLearningConfig)
	return out
}

func VertexAIModelSpec_DisplayName_ToProto(mapCtx *direct.MapContext, in string) string {
	return in
}
