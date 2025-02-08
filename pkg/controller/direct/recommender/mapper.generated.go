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
func InsightType_FromProto(mapCtx *direct.MapContext, in *pb.InsightType) *krm.InsightType {
	if in == nil {
		return nil
	}
	out := &krm.InsightType{}
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func InsightType_ToProto(mapCtx *direct.MapContext, in *krm.InsightType) *pb.InsightType {
	if in == nil {
		return nil
	}
	out := &pb.InsightType{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
func RecommenderInsightTypeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InsightType) *krm.RecommenderInsightTypeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RecommenderInsightTypeObservedState{}
	// MISSING: Name
	return out
}
func RecommenderInsightTypeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RecommenderInsightTypeObservedState) *pb.InsightType {
	if in == nil {
		return nil
	}
	out := &pb.InsightType{}
	// MISSING: Name
	return out
}
func RecommenderInsightTypeSpec_FromProto(mapCtx *direct.MapContext, in *pb.InsightType) *krm.RecommenderInsightTypeSpec {
	if in == nil {
		return nil
	}
	out := &krm.RecommenderInsightTypeSpec{}
	// MISSING: Name
	return out
}
func RecommenderInsightTypeSpec_ToProto(mapCtx *direct.MapContext, in *krm.RecommenderInsightTypeSpec) *pb.InsightType {
	if in == nil {
		return nil
	}
	out := &pb.InsightType{}
	// MISSING: Name
	return out
}
