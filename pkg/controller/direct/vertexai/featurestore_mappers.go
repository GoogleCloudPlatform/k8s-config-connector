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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func EncryptionSpec_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionSpec) *krm.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionSpec{}
	if in.GetKmsKeyName() != "" {
		out.KMSKeyRef = &refs.KMSCryptoKeyRef{External: in.GetKmsKeyName()}
	}
	return out
}
func EncryptionSpec_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionSpec) *pb.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionSpec{}
	if in.KMSKeyRef != nil {
		out.KmsKeyName = in.KMSKeyRef.External
	}
	return out
}
func VertexAIFeaturestoreObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Featurestore) *krm.VertexAIFeaturestoreObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIFeaturestoreObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func VertexAIFeaturestoreObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIFeaturestoreObservedState) *pb.Featurestore {
	if in == nil {
		return nil
	}
	out := &pb.Featurestore{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.Featurestore_State](mapCtx, in.State)
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func VertexAIFeaturestoreSpec_FromProto(mapCtx *direct.MapContext, in *pb.Featurestore) *krm.VertexAIFeaturestoreSpec {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIFeaturestoreSpec{}
	// MISSING: Name
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Labels = in.Labels
	out.OnlineServingConfig = Featurestore_OnlineServingConfig_FromProto(mapCtx, in.GetOnlineServingConfig())
	out.OnlineStorageTTLDays = direct.LazyPtr(in.GetOnlineStorageTtlDays())
	out.EncryptionSpec = EncryptionSpec_FromProto(mapCtx, in.GetEncryptionSpec())
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func VertexAIFeaturestoreSpec_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIFeaturestoreSpec) *pb.Featurestore {
	if in == nil {
		return nil
	}
	out := &pb.Featurestore{}
	// MISSING: Name
	out.Etag = direct.ValueOf(in.Etag)
	out.Labels = in.Labels
	out.OnlineServingConfig = Featurestore_OnlineServingConfig_ToProto(mapCtx, in.OnlineServingConfig)
	out.OnlineStorageTtlDays = direct.ValueOf(in.OnlineStorageTTLDays)
	out.EncryptionSpec = EncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
