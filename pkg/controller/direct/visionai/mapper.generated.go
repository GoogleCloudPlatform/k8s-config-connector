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
func Asset_FromProto(mapCtx *direct.MapContext, in *pb.Asset) *krm.Asset {
	if in == nil {
		return nil
	}
	out := &krm.Asset{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Ttl = direct.StringDuration_FromProto(mapCtx, in.GetTtl())
	// MISSING: AssetGcsSource
	return out
}
func Asset_ToProto(mapCtx *direct.MapContext, in *krm.Asset) *pb.Asset {
	if in == nil {
		return nil
	}
	out := &pb.Asset{}
	out.Name = direct.ValueOf(in.Name)
	out.Ttl = direct.StringDuration_ToProto(mapCtx, in.Ttl)
	// MISSING: AssetGcsSource
	return out
}
func AssetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Asset) *krm.AssetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AssetObservedState{}
	// MISSING: Name
	// MISSING: Ttl
	out.AssetGcsSource = AssetSource_AssetGcsSource_FromProto(mapCtx, in.GetAssetGcsSource())
	return out
}
func AssetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AssetObservedState) *pb.Asset {
	if in == nil {
		return nil
	}
	out := &pb.Asset{}
	// MISSING: Name
	// MISSING: Ttl
	out.AssetGcsSource = AssetSource_AssetGcsSource_ToProto(mapCtx, in.AssetGcsSource)
	return out
}
func VisionaiAssetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Asset) *krm.VisionaiAssetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiAssetObservedState{}
	// MISSING: Name
	// MISSING: Ttl
	// MISSING: AssetGcsSource
	return out
}
func VisionaiAssetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiAssetObservedState) *pb.Asset {
	if in == nil {
		return nil
	}
	out := &pb.Asset{}
	// MISSING: Name
	// MISSING: Ttl
	// MISSING: AssetGcsSource
	return out
}
func VisionaiAssetSpec_FromProto(mapCtx *direct.MapContext, in *pb.Asset) *krm.VisionaiAssetSpec {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiAssetSpec{}
	// MISSING: Name
	// MISSING: Ttl
	// MISSING: AssetGcsSource
	return out
}
func VisionaiAssetSpec_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiAssetSpec) *pb.Asset {
	if in == nil {
		return nil
	}
	out := &pb.Asset{}
	// MISSING: Name
	// MISSING: Ttl
	// MISSING: AssetGcsSource
	return out
}
