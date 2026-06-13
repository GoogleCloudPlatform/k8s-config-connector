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

package aiplatform

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/structpb"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func JSON_FromProto(mapCtx *direct.MapContext, in *structpb.Value) *apiextensionsv1.JSON {
	if in == nil || in.Kind == nil {
		return nil
	}
	b, err := in.MarshalJSON()
	if err != nil {
		mapCtx.Errorf("error marshaling structpb.Value to JSON: %v", err)
		return nil
	}
	return &apiextensionsv1.JSON{Raw: b}
}

func JSON_ToProto(mapCtx *direct.MapContext, in *apiextensionsv1.JSON) *structpb.Value {
	if in == nil {
		return nil
	}
	out := &structpb.Value{}
	if err := out.UnmarshalJSON(in.Raw); err != nil {
		mapCtx.Errorf("error unmarshaling JSON to structpb.Value: %v", err)
		return nil
	}
	return out
}

func VertexAIEncryptionSpec_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionSpec) *krm.VertexAIEncryptionSpec {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIEncryptionSpec{}
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	return out
}

func VertexAIEncryptionSpec_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIEncryptionSpec) *pb.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionSpec{}
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	return out
}

func VertexAIDeployedIndexRefObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DeployedIndexRef) *krm.VertexAIDeployedIndexRefObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIDeployedIndexRefObservedState{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}

func VertexAIDeployedIndexRefObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIDeployedIndexRefObservedState) *pb.DeployedIndexRef {
	if in == nil {
		return nil
	}
	out := &pb.DeployedIndexRef{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}

func VertexAIIndexStatsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.IndexStats) *krm.VertexAIIndexStatsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIIndexStatsObservedState{}
	out.VectorsCount = direct.LazyPtr(in.GetVectorsCount())
	out.SparseVectorsCount = direct.LazyPtr(in.GetSparseVectorsCount())
	out.ShardsCount = direct.LazyPtr(in.GetShardsCount())
	return out
}

func VertexAIIndexStatsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIIndexStatsObservedState) *pb.IndexStats {
	if in == nil {
		return nil
	}
	out := &pb.IndexStats{}
	out.VectorsCount = direct.ValueOf(in.VectorsCount)
	out.SparseVectorsCount = direct.ValueOf(in.SparseVectorsCount)
	out.ShardsCount = direct.ValueOf(in.ShardsCount)
	return out
}

func VertexAIIndexObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Index) *krm.VertexAIIndexObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIIndexObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DeployedIndexes = direct.Slice_FromProto(mapCtx, in.DeployedIndexes, VertexAIDeployedIndexRefObservedState_FromProto)
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.IndexStats = VertexAIIndexStatsObservedState_FromProto(mapCtx, in.GetIndexStats())
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	return out
}

func VertexAIIndexObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIIndexObservedState) *pb.Index {
	if in == nil {
		return nil
	}
	out := &pb.Index{}
	out.Name = direct.ValueOf(in.Name)
	out.DeployedIndexes = direct.Slice_ToProto(mapCtx, in.DeployedIndexes, VertexAIDeployedIndexRefObservedState_ToProto)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.IndexStats = VertexAIIndexStatsObservedState_ToProto(mapCtx, in.IndexStats)
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.SatisfiesPzi = direct.ValueOf(in.SatisfiesPzi)
	return out
}

func VertexAIIndexSpec_FromProto(mapCtx *direct.MapContext, in *pb.Index) *krm.VertexAIIndexSpec {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIIndexSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.MetadataSchemaURI = direct.LazyPtr(in.GetMetadataSchemaUri())
	out.Metadata = JSON_FromProto(mapCtx, in.GetMetadata())
	out.Labels = in.Labels
	out.IndexUpdateMethod = direct.Enum_FromProto(mapCtx, in.GetIndexUpdateMethod())
	out.EncryptionSpec = VertexAIEncryptionSpec_FromProto(mapCtx, in.GetEncryptionSpec())
	return out
}

func VertexAIIndexSpec_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIIndexSpec) *pb.Index {
	if in == nil {
		return nil
	}
	out := &pb.Index{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.MetadataSchemaUri = direct.ValueOf(in.MetadataSchemaURI)
	out.Metadata = JSON_ToProto(mapCtx, in.Metadata)
	out.Labels = in.Labels
	out.IndexUpdateMethod = direct.Enum_ToProto[pb.Index_IndexUpdateMethod](mapCtx, in.IndexUpdateMethod)
	out.EncryptionSpec = VertexAIEncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)
	return out
}
