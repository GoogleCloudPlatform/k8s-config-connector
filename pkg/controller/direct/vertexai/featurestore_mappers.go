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
	krmvertexaiv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func VertexAIEncryptionSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionSpec) *krmvertexaiv1alpha1.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &krmvertexaiv1alpha1.EncryptionSpec{}
	// MISSING: KMSKeyName
	return out
}
func VertexAIEncryptionSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmvertexaiv1alpha1.EncryptionSpec) *pb.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionSpec{}
	// MISSING: KMSKeyName
	return out
}

func VertexAIFeaturestoreSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Featurestore) *krmvertexaiv1alpha1.VertexAIFeaturestoreSpec {
	if in == nil {
		return nil
	}
	out := &krmvertexaiv1alpha1.VertexAIFeaturestoreSpec{}
	// MISSING: Name
	// MISSING: Etag
	out.Labels = in.Labels
	out.OnlineServingConfig = Featurestore_OnlineServingConfig_v1alpha1_FromProto(mapCtx, in.GetOnlineServingConfig())
	out.OnlineStorageTTLDays = direct.LazyPtr(in.GetOnlineStorageTtlDays())
	out.EncryptionSpec = VertexAIEncryptionSpec_v1alpha1_FromProto(mapCtx, in.GetEncryptionSpec())
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func VertexAIFeaturestoreSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmvertexaiv1alpha1.VertexAIFeaturestoreSpec) *pb.Featurestore {
	if in == nil {
		return nil
	}
	out := &pb.Featurestore{}
	// MISSING: Name
	// MISSING: Etag
	out.Labels = in.Labels
	out.OnlineServingConfig = Featurestore_OnlineServingConfig_v1alpha1_ToProto(mapCtx, in.OnlineServingConfig)
	out.OnlineStorageTtlDays = direct.ValueOf(in.OnlineStorageTTLDays)
	out.EncryptionSpec = VertexAIEncryptionSpec_v1alpha1_ToProto(mapCtx, in.EncryptionSpec)
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
