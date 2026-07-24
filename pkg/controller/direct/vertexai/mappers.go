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
