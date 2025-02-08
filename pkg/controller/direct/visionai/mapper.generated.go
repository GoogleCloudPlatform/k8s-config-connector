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

package visionai

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/visionai/apiv1/visionaipb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/visionai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func SearchHypernym_FromProto(mapCtx *direct.MapContext, in *pb.SearchHypernym) *krm.SearchHypernym {
	if in == nil {
		return nil
	}
	out := &krm.SearchHypernym{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Hypernym = direct.LazyPtr(in.GetHypernym())
	out.Hyponyms = in.Hyponyms
	return out
}
func SearchHypernym_ToProto(mapCtx *direct.MapContext, in *krm.SearchHypernym) *pb.SearchHypernym {
	if in == nil {
		return nil
	}
	out := &pb.SearchHypernym{}
	out.Name = direct.ValueOf(in.Name)
	out.Hypernym = direct.ValueOf(in.Hypernym)
	out.Hyponyms = in.Hyponyms
	return out
}
func VisionaiSearchHypernymObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SearchHypernym) *krm.VisionaiSearchHypernymObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiSearchHypernymObservedState{}
	// MISSING: Name
	// MISSING: Hypernym
	// MISSING: Hyponyms
	return out
}
func VisionaiSearchHypernymObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiSearchHypernymObservedState) *pb.SearchHypernym {
	if in == nil {
		return nil
	}
	out := &pb.SearchHypernym{}
	// MISSING: Name
	// MISSING: Hypernym
	// MISSING: Hyponyms
	return out
}
func VisionaiSearchHypernymSpec_FromProto(mapCtx *direct.MapContext, in *pb.SearchHypernym) *krm.VisionaiSearchHypernymSpec {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiSearchHypernymSpec{}
	// MISSING: Name
	// MISSING: Hypernym
	// MISSING: Hyponyms
	return out
}
func VisionaiSearchHypernymSpec_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiSearchHypernymSpec) *pb.SearchHypernym {
	if in == nil {
		return nil
	}
	out := &pb.SearchHypernym{}
	// MISSING: Name
	// MISSING: Hypernym
	// MISSING: Hyponyms
	return out
}
