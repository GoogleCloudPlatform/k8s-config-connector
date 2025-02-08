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

package dataplex

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func DataAttributeBinding_FromProto(mapCtx *direct.MapContext, in *pb.DataAttributeBinding) *krm.DataAttributeBinding {
	if in == nil {
		return nil
	}
	out := &krm.DataAttributeBinding{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Labels = in.Labels
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Resource = direct.LazyPtr(in.GetResource())
	out.Attributes = in.Attributes
	out.Paths = direct.Slice_FromProto(mapCtx, in.Paths, DataAttributeBinding_Path_FromProto)
	return out
}
func DataAttributeBinding_ToProto(mapCtx *direct.MapContext, in *krm.DataAttributeBinding) *pb.DataAttributeBinding {
	if in == nil {
		return nil
	}
	out := &pb.DataAttributeBinding{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.ValueOf(in.Description)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Labels = in.Labels
	out.Etag = direct.ValueOf(in.Etag)
	if oneof := DataAttributeBinding_Resource_ToProto(mapCtx, in.Resource); oneof != nil {
		out.ResourceReference = oneof
	}
	out.Attributes = in.Attributes
	out.Paths = direct.Slice_ToProto(mapCtx, in.Paths, DataAttributeBinding_Path_ToProto)
	return out
}
func DataAttributeBindingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataAttributeBinding) *krm.DataAttributeBindingObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataAttributeBindingObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: Resource
	// MISSING: Attributes
	// MISSING: Paths
	return out
}
func DataAttributeBindingObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataAttributeBindingObservedState) *pb.DataAttributeBinding {
	if in == nil {
		return nil
	}
	out := &pb.DataAttributeBinding{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: Resource
	// MISSING: Attributes
	// MISSING: Paths
	return out
}
func DataAttributeBinding_Path_FromProto(mapCtx *direct.MapContext, in *pb.DataAttributeBinding_Path) *krm.DataAttributeBinding_Path {
	if in == nil {
		return nil
	}
	out := &krm.DataAttributeBinding_Path{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Attributes = in.Attributes
	return out
}
func DataAttributeBinding_Path_ToProto(mapCtx *direct.MapContext, in *krm.DataAttributeBinding_Path) *pb.DataAttributeBinding_Path {
	if in == nil {
		return nil
	}
	out := &pb.DataAttributeBinding_Path{}
	out.Name = direct.ValueOf(in.Name)
	out.Attributes = in.Attributes
	return out
}
func DataplexDataAttributeBindingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataAttributeBinding) *krm.DataplexDataAttributeBindingObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataplexDataAttributeBindingObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: Resource
	// MISSING: Attributes
	// MISSING: Paths
	return out
}
func DataplexDataAttributeBindingObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataplexDataAttributeBindingObservedState) *pb.DataAttributeBinding {
	if in == nil {
		return nil
	}
	out := &pb.DataAttributeBinding{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: Resource
	// MISSING: Attributes
	// MISSING: Paths
	return out
}
func DataplexDataAttributeBindingSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataAttributeBinding) *krm.DataplexDataAttributeBindingSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataplexDataAttributeBindingSpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: Resource
	// MISSING: Attributes
	// MISSING: Paths
	return out
}
func DataplexDataAttributeBindingSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataplexDataAttributeBindingSpec) *pb.DataAttributeBinding {
	if in == nil {
		return nil
	}
	out := &pb.DataAttributeBinding{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: Resource
	// MISSING: Attributes
	// MISSING: Paths
	return out
}
