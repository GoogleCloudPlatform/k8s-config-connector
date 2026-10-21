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

package bigtable

import (
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BigtableMemoryLayerObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.MemoryLayer) *krm.BigtableMemoryLayerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigtableMemoryLayerObservedState{}
	// MISSING: Name
	// MISSING: Etag, which is not declarative-friendly and so is not surfaced in the KRM API.
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.MemoryConfig = MemoryLayer_MemoryConfigObservedState_v1alpha1_FromProto(mapCtx, in.GetMemoryConfig())
	return out
}

func BigtableMemoryLayerObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.BigtableMemoryLayerObservedState) *pb.MemoryLayer {
	if in == nil {
		return nil
	}
	out := &pb.MemoryLayer{}
	// MISSING: Name
	// MISSING: Etag, which is not declarative-friendly and so is not surfaced in the KRM API.
	out.State = direct.Enum_ToProto[pb.MemoryLayer_State](mapCtx, in.State)
	out.MemoryConfig = MemoryLayer_MemoryConfigObservedState_v1alpha1_ToProto(mapCtx, in.MemoryConfig)
	return out
}

func MemoryLayer_MemoryConfigObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.MemoryLayer_MemoryConfig) *krm.MemoryLayer_MemoryConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MemoryLayer_MemoryConfigObservedState{}
	out.StorageSizeGib = direct.LazyPtr(in.GetStorageSizeGib())
	return out
}

func MemoryLayer_MemoryConfigObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.MemoryLayer_MemoryConfigObservedState) *pb.MemoryLayer_MemoryConfig {
	if in == nil {
		return nil
	}
	out := &pb.MemoryLayer_MemoryConfig{}
	out.StorageSizeGib = direct.ValueOf(in.StorageSizeGib)
	return out
}

// BigtableMemoryLayerSpec_v1alpha1_FromProto and its ToProto counterpart carry
// no fields: every settable part of a MemoryLayer is either its identity
// (spec.clusterRef, derived from the resource name) or not surfaced in the KRM
// API. They exist so the round-trip fuzzer can cover the proto message, which
// makes a newly added MemoryLayer field fail the fuzz test until it is
// deliberately classified as spec, status, or unimplemented.
func BigtableMemoryLayerSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.MemoryLayer) *krm.BigtableMemoryLayerSpec {
	if in == nil {
		return nil
	}
	// MISSING: Name, which is modelled as spec.clusterRef by the identity.
	return &krm.BigtableMemoryLayerSpec{}
}

func BigtableMemoryLayerSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.BigtableMemoryLayerSpec) *pb.MemoryLayer {
	if in == nil {
		return nil
	}
	// MISSING: Name, which is modelled as spec.clusterRef by the identity.
	return &pb.MemoryLayer{}
}
