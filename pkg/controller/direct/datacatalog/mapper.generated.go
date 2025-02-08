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

package datacatalog

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/datacatalog/lineage/apiv1/lineagepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datacatalog/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func DatacatalogProcessObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Process) *krm.DatacatalogProcessObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatacatalogProcessObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Attributes
	// MISSING: Origin
	return out
}
func DatacatalogProcessObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatacatalogProcessObservedState) *pb.Process {
	if in == nil {
		return nil
	}
	out := &pb.Process{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Attributes
	// MISSING: Origin
	return out
}
func DatacatalogProcessSpec_FromProto(mapCtx *direct.MapContext, in *pb.Process) *krm.DatacatalogProcessSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatacatalogProcessSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Attributes
	// MISSING: Origin
	return out
}
func DatacatalogProcessSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatacatalogProcessSpec) *pb.Process {
	if in == nil {
		return nil
	}
	out := &pb.Process{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Attributes
	// MISSING: Origin
	return out
}
func Origin_FromProto(mapCtx *direct.MapContext, in *pb.Origin) *krm.Origin {
	if in == nil {
		return nil
	}
	out := &krm.Origin{}
	out.SourceType = direct.Enum_FromProto(mapCtx, in.GetSourceType())
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func Origin_ToProto(mapCtx *direct.MapContext, in *krm.Origin) *pb.Origin {
	if in == nil {
		return nil
	}
	out := &pb.Origin{}
	out.SourceType = direct.Enum_ToProto[pb.Origin_SourceType](mapCtx, in.SourceType)
	out.Name = direct.ValueOf(in.Name)
	return out
}
func Process_FromProto(mapCtx *direct.MapContext, in *pb.Process) *krm.Process {
	if in == nil {
		return nil
	}
	out := &krm.Process{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Attributes
	out.Origin = Origin_FromProto(mapCtx, in.GetOrigin())
	return out
}
func Process_ToProto(mapCtx *direct.MapContext, in *krm.Process) *pb.Process {
	if in == nil {
		return nil
	}
	out := &pb.Process{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Attributes
	out.Origin = Origin_ToProto(mapCtx, in.Origin)
	return out
}
