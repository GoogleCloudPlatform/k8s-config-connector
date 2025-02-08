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

package recommender

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/recommender/apiv1beta1/recommenderpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/recommender/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func RecommenderRecommenderTypeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RecommenderType) *krm.RecommenderRecommenderTypeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RecommenderRecommenderTypeObservedState{}
	// MISSING: Name
	return out
}
func RecommenderRecommenderTypeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RecommenderRecommenderTypeObservedState) *pb.RecommenderType {
	if in == nil {
		return nil
	}
	out := &pb.RecommenderType{}
	// MISSING: Name
	return out
}
func RecommenderRecommenderTypeSpec_FromProto(mapCtx *direct.MapContext, in *pb.RecommenderType) *krm.RecommenderRecommenderTypeSpec {
	if in == nil {
		return nil
	}
	out := &krm.RecommenderRecommenderTypeSpec{}
	// MISSING: Name
	return out
}
func RecommenderRecommenderTypeSpec_ToProto(mapCtx *direct.MapContext, in *krm.RecommenderRecommenderTypeSpec) *pb.RecommenderType {
	if in == nil {
		return nil
	}
	out := &pb.RecommenderType{}
	// MISSING: Name
	return out
}
func RecommenderType_FromProto(mapCtx *direct.MapContext, in *pb.RecommenderType) *krm.RecommenderType {
	if in == nil {
		return nil
	}
	out := &krm.RecommenderType{}
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func RecommenderType_ToProto(mapCtx *direct.MapContext, in *krm.RecommenderType) *pb.RecommenderType {
	if in == nil {
		return nil
	}
	out := &pb.RecommenderType{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
