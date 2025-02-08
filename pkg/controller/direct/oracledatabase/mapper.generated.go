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

package oracledatabase

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/oracledatabase/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/oracledatabase/apiv1/oracledatabasepb"
)
func DbSystemShape_FromProto(mapCtx *direct.MapContext, in *pb.DbSystemShape) *krm.DbSystemShape {
	if in == nil {
		return nil
	}
	out := &krm.DbSystemShape{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Shape = direct.LazyPtr(in.GetShape())
	out.MinNodeCount = direct.LazyPtr(in.GetMinNodeCount())
	out.MaxNodeCount = direct.LazyPtr(in.GetMaxNodeCount())
	out.MinStorageCount = direct.LazyPtr(in.GetMinStorageCount())
	out.MaxStorageCount = direct.LazyPtr(in.GetMaxStorageCount())
	out.AvailableCoreCountPerNode = direct.LazyPtr(in.GetAvailableCoreCountPerNode())
	out.AvailableMemoryPerNodeGB = direct.LazyPtr(in.GetAvailableMemoryPerNodeGb())
	out.AvailableDataStorageTb = direct.LazyPtr(in.GetAvailableDataStorageTb())
	out.MinCoreCountPerNode = direct.LazyPtr(in.GetMinCoreCountPerNode())
	out.MinMemoryPerNodeGB = direct.LazyPtr(in.GetMinMemoryPerNodeGb())
	out.MinDbNodeStoragePerNodeGB = direct.LazyPtr(in.GetMinDbNodeStoragePerNodeGb())
	return out
}
func DbSystemShape_ToProto(mapCtx *direct.MapContext, in *krm.DbSystemShape) *pb.DbSystemShape {
	if in == nil {
		return nil
	}
	out := &pb.DbSystemShape{}
	out.Name = direct.ValueOf(in.Name)
	out.Shape = direct.ValueOf(in.Shape)
	out.MinNodeCount = direct.ValueOf(in.MinNodeCount)
	out.MaxNodeCount = direct.ValueOf(in.MaxNodeCount)
	out.MinStorageCount = direct.ValueOf(in.MinStorageCount)
	out.MaxStorageCount = direct.ValueOf(in.MaxStorageCount)
	out.AvailableCoreCountPerNode = direct.ValueOf(in.AvailableCoreCountPerNode)
	out.AvailableMemoryPerNodeGb = direct.ValueOf(in.AvailableMemoryPerNodeGB)
	out.AvailableDataStorageTb = direct.ValueOf(in.AvailableDataStorageTb)
	out.MinCoreCountPerNode = direct.ValueOf(in.MinCoreCountPerNode)
	out.MinMemoryPerNodeGb = direct.ValueOf(in.MinMemoryPerNodeGB)
	out.MinDbNodeStoragePerNodeGb = direct.ValueOf(in.MinDbNodeStoragePerNodeGB)
	return out
}
func OracledatabaseDbSystemShapeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DbSystemShape) *krm.OracledatabaseDbSystemShapeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OracledatabaseDbSystemShapeObservedState{}
	// MISSING: Name
	// MISSING: Shape
	// MISSING: MinNodeCount
	// MISSING: MaxNodeCount
	// MISSING: MinStorageCount
	// MISSING: MaxStorageCount
	// MISSING: AvailableCoreCountPerNode
	// MISSING: AvailableMemoryPerNodeGB
	// MISSING: AvailableDataStorageTb
	// MISSING: MinCoreCountPerNode
	// MISSING: MinMemoryPerNodeGB
	// MISSING: MinDbNodeStoragePerNodeGB
	return out
}
func OracledatabaseDbSystemShapeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OracledatabaseDbSystemShapeObservedState) *pb.DbSystemShape {
	if in == nil {
		return nil
	}
	out := &pb.DbSystemShape{}
	// MISSING: Name
	// MISSING: Shape
	// MISSING: MinNodeCount
	// MISSING: MaxNodeCount
	// MISSING: MinStorageCount
	// MISSING: MaxStorageCount
	// MISSING: AvailableCoreCountPerNode
	// MISSING: AvailableMemoryPerNodeGB
	// MISSING: AvailableDataStorageTb
	// MISSING: MinCoreCountPerNode
	// MISSING: MinMemoryPerNodeGB
	// MISSING: MinDbNodeStoragePerNodeGB
	return out
}
func OracledatabaseDbSystemShapeSpec_FromProto(mapCtx *direct.MapContext, in *pb.DbSystemShape) *krm.OracledatabaseDbSystemShapeSpec {
	if in == nil {
		return nil
	}
	out := &krm.OracledatabaseDbSystemShapeSpec{}
	// MISSING: Name
	// MISSING: Shape
	// MISSING: MinNodeCount
	// MISSING: MaxNodeCount
	// MISSING: MinStorageCount
	// MISSING: MaxStorageCount
	// MISSING: AvailableCoreCountPerNode
	// MISSING: AvailableMemoryPerNodeGB
	// MISSING: AvailableDataStorageTb
	// MISSING: MinCoreCountPerNode
	// MISSING: MinMemoryPerNodeGB
	// MISSING: MinDbNodeStoragePerNodeGB
	return out
}
func OracledatabaseDbSystemShapeSpec_ToProto(mapCtx *direct.MapContext, in *krm.OracledatabaseDbSystemShapeSpec) *pb.DbSystemShape {
	if in == nil {
		return nil
	}
	out := &pb.DbSystemShape{}
	// MISSING: Name
	// MISSING: Shape
	// MISSING: MinNodeCount
	// MISSING: MaxNodeCount
	// MISSING: MinStorageCount
	// MISSING: MaxStorageCount
	// MISSING: AvailableCoreCountPerNode
	// MISSING: AvailableMemoryPerNodeGB
	// MISSING: AvailableDataStorageTb
	// MISSING: MinCoreCountPerNode
	// MISSING: MinMemoryPerNodeGB
	// MISSING: MinDbNodeStoragePerNodeGB
	return out
}
