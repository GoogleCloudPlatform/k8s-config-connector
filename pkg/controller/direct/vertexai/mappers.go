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
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
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
		KMSKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
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
		KMSKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
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
		out.KmsKeyNameRef = &kmsv1beta1.KMSCryptoKeyRef{
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

// VertexAIFeaturestoreEntityTypeFeatureSpec_v1alpha1_FromProto maps a protobuf Feature to VertexAIFeaturestoreEntityTypeFeatureSpec.
// Handcoded because ValueType is a required string in KRM to maintain strict baseline CRD schema compatibility,
// whereas in proto it is an enum (pb.Feature_ValueType) which requires converting between string and *string
// for direct.Enum_FromProto and direct.Enum_ToProto.
func VertexAIFeaturestoreEntityTypeFeatureSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Feature) *krmv1alpha1.VertexAIFeaturestoreEntityTypeFeatureSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.VertexAIFeaturestoreEntityTypeFeatureSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	if v := direct.Enum_FromProto(mapCtx, in.GetValueType()); v != nil {
		out.ValueType = *v
	}
	return out
}

// VertexAIFeaturestoreEntityTypeFeatureSpec_v1alpha1_ToProto maps a VertexAIFeaturestoreEntityTypeFeatureSpec to protobuf Feature.
// Handcoded because ValueType is a required string in KRM to maintain strict baseline CRD schema compatibility,
// whereas in proto it is an enum (pb.Feature_ValueType) which requires converting between string and *string
// for direct.Enum_FromProto and direct.Enum_ToProto.
func VertexAIFeaturestoreEntityTypeFeatureSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.VertexAIFeaturestoreEntityTypeFeatureSpec) *pb.Feature {
	if in == nil {
		return nil
	}
	out := &pb.Feature{}
	out.Description = direct.ValueOf(in.Description)
	out.ValueType = direct.Enum_ToProto[pb.Feature_ValueType](mapCtx, &in.ValueType)
	return out
}

// FeaturestoreMonitoringConfig_ThresholdConfig_v1alpha1_FromProto maps a protobuf FeaturestoreMonitoringConfig_ThresholdConfig to FeaturestoreMonitoringConfig_ThresholdConfig.
// Handcoded because Value is a required float64 in KRM to maintain strict baseline CRD schema compatibility,
// whereas in proto Value is an optional float64 within a oneof (Threshold) which requires converting between float64 and *float64.
func FeaturestoreMonitoringConfig_ThresholdConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.FeaturestoreMonitoringConfig_ThresholdConfig) *krmv1alpha1.FeaturestoreMonitoringConfig_ThresholdConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.FeaturestoreMonitoringConfig_ThresholdConfig{}
	out.Value = in.GetValue()
	return out
}

// FeaturestoreMonitoringConfig_ThresholdConfig_v1alpha1_ToProto maps a FeaturestoreMonitoringConfig_ThresholdConfig to protobuf FeaturestoreMonitoringConfig_ThresholdConfig.
// Handcoded because Value is a required float64 in KRM to maintain strict baseline CRD schema compatibility,
// whereas in proto Value is an optional float64 within a oneof (Threshold) which requires converting between float64 and *float64.
func FeaturestoreMonitoringConfig_ThresholdConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.FeaturestoreMonitoringConfig_ThresholdConfig) *pb.FeaturestoreMonitoringConfig_ThresholdConfig {
	if in == nil {
		return nil
	}
	out := &pb.FeaturestoreMonitoringConfig_ThresholdConfig{}
	out.Threshold = &pb.FeaturestoreMonitoringConfig_ThresholdConfig_Value{
		Value: in.Value,
	}
	return out
}

func FeaturestoreMonitoringConfig_ThresholdConfig_Value_ToProto(mapCtx *direct.MapContext, in float64) *pb.FeaturestoreMonitoringConfig_ThresholdConfig_Value {
	return &pb.FeaturestoreMonitoringConfig_ThresholdConfig_Value{Value: in}
}

// FeaturestoreMonitoringConfig_SnapshotAnalysis_v1alpha1_FromProto maps a protobuf FeaturestoreMonitoringConfig_SnapshotAnalysis to FeaturestoreMonitoringConfig_SnapshotAnalysis.
// Handcoded because monitoringIntervalDays and stalenessDays are unformatted integers (*int) in KRM to maintain strict baseline CRD schema compatibility,
// whereas in proto they are int32.
func FeaturestoreMonitoringConfig_SnapshotAnalysis_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.FeaturestoreMonitoringConfig_SnapshotAnalysis) *krmv1alpha1.FeaturestoreMonitoringConfig_SnapshotAnalysis {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.FeaturestoreMonitoringConfig_SnapshotAnalysis{}
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	out.MonitoringInterval = direct.StringDuration_FromProto(mapCtx, in.GetMonitoringInterval())
	out.MonitoringIntervalDays = direct.LazyPtr(int(in.GetMonitoringIntervalDays()))
	out.StalenessDays = direct.LazyPtr(int(in.GetStalenessDays()))
	return out
}

// FeaturestoreMonitoringConfig_SnapshotAnalysis_v1alpha1_ToProto maps a FeaturestoreMonitoringConfig_SnapshotAnalysis to protobuf FeaturestoreMonitoringConfig_SnapshotAnalysis.
// Handcoded because monitoringIntervalDays and stalenessDays are unformatted integers (*int) in KRM to maintain strict baseline CRD schema compatibility,
// whereas in proto they are int32.
func FeaturestoreMonitoringConfig_SnapshotAnalysis_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.FeaturestoreMonitoringConfig_SnapshotAnalysis) *pb.FeaturestoreMonitoringConfig_SnapshotAnalysis {
	if in == nil {
		return nil
	}
	out := &pb.FeaturestoreMonitoringConfig_SnapshotAnalysis{}
	out.Disabled = direct.ValueOf(in.Disabled)
	out.MonitoringInterval = direct.StringDuration_ToProto(mapCtx, in.MonitoringInterval)
	out.MonitoringIntervalDays = int32(direct.ValueOf(in.MonitoringIntervalDays))
	out.StalenessDays = int32(direct.ValueOf(in.StalenessDays))
	return out
}

// VertexAIFeaturestoreEntityTypeSpec_v1alpha1_FromProto maps a protobuf EntityType to VertexAIFeaturestoreEntityTypeSpec.
// Handcoded because offlineStorageTtlDays is an unformatted integer (*int) in KRM to maintain strict baseline CRD schema compatibility,
// whereas in proto it is int32.
func VertexAIFeaturestoreEntityTypeSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.EntityType) *krmv1alpha1.VertexAIFeaturestoreEntityTypeSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.VertexAIFeaturestoreEntityTypeSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.MonitoringConfig = FeaturestoreMonitoringConfig_v1alpha1_FromProto(mapCtx, in.GetMonitoringConfig())
	out.OfflineStorageTTLDays = direct.LazyPtr(int(in.GetOfflineStorageTtlDays()))
	return out
}

// VertexAIFeaturestoreEntityTypeSpec_v1alpha1_ToProto maps a VertexAIFeaturestoreEntityTypeSpec to protobuf EntityType.
// Handcoded because offlineStorageTtlDays is an unformatted integer (*int) in KRM to maintain strict baseline CRD schema compatibility,
// whereas in proto it is int32.
func VertexAIFeaturestoreEntityTypeSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.VertexAIFeaturestoreEntityTypeSpec) *pb.EntityType {
	if in == nil {
		return nil
	}
	out := &pb.EntityType{}
	out.Description = direct.ValueOf(in.Description)
	out.MonitoringConfig = FeaturestoreMonitoringConfig_v1alpha1_ToProto(mapCtx, in.MonitoringConfig)
	out.OfflineStorageTtlDays = int32(direct.ValueOf(in.OfflineStorageTTLDays))
	return out
}
