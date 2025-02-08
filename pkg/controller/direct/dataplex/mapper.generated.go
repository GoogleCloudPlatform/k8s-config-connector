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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
)
func DataAccessSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataAccessSpec) *krm.DataAccessSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataAccessSpec{}
	out.Readers = in.Readers
	return out
}
func DataAccessSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataAccessSpec) *pb.DataAccessSpec {
	if in == nil {
		return nil
	}
	out := &pb.DataAccessSpec{}
	out.Readers = in.Readers
	return out
}
func DataAttribute_FromProto(mapCtx *direct.MapContext, in *pb.DataAttribute) *krm.DataAttribute {
	if in == nil {
		return nil
	}
	out := &krm.DataAttribute{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Labels = in.Labels
	out.ParentID = direct.LazyPtr(in.GetParentId())
	// MISSING: AttributeCount
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.ResourceAccessSpec = ResourceAccessSpec_FromProto(mapCtx, in.GetResourceAccessSpec())
	out.DataAccessSpec = DataAccessSpec_FromProto(mapCtx, in.GetDataAccessSpec())
	return out
}
func DataAttribute_ToProto(mapCtx *direct.MapContext, in *krm.DataAttribute) *pb.DataAttribute {
	if in == nil {
		return nil
	}
	out := &pb.DataAttribute{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.ValueOf(in.Description)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Labels = in.Labels
	out.ParentId = direct.ValueOf(in.ParentID)
	// MISSING: AttributeCount
	out.Etag = direct.ValueOf(in.Etag)
	out.ResourceAccessSpec = ResourceAccessSpec_ToProto(mapCtx, in.ResourceAccessSpec)
	out.DataAccessSpec = DataAccessSpec_ToProto(mapCtx, in.DataAccessSpec)
	return out
}
func DataAttributeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataAttribute) *krm.DataAttributeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataAttributeObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: ParentID
	out.AttributeCount = direct.LazyPtr(in.GetAttributeCount())
	// MISSING: Etag
	// MISSING: ResourceAccessSpec
	// MISSING: DataAccessSpec
	return out
}
func DataAttributeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataAttributeObservedState) *pb.DataAttribute {
	if in == nil {
		return nil
	}
	out := &pb.DataAttribute{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: ParentID
	out.AttributeCount = direct.ValueOf(in.AttributeCount)
	// MISSING: Etag
	// MISSING: ResourceAccessSpec
	// MISSING: DataAccessSpec
	return out
}
func DataplexDataAttributeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataAttribute) *krm.DataplexDataAttributeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataplexDataAttributeObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: ParentID
	// MISSING: AttributeCount
	// MISSING: Etag
	// MISSING: ResourceAccessSpec
	// MISSING: DataAccessSpec
	return out
}
func DataplexDataAttributeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataplexDataAttributeObservedState) *pb.DataAttribute {
	if in == nil {
		return nil
	}
	out := &pb.DataAttribute{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: ParentID
	// MISSING: AttributeCount
	// MISSING: Etag
	// MISSING: ResourceAccessSpec
	// MISSING: DataAccessSpec
	return out
}
func DataplexDataAttributeSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataAttribute) *krm.DataplexDataAttributeSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataplexDataAttributeSpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: ParentID
	// MISSING: AttributeCount
	// MISSING: Etag
	// MISSING: ResourceAccessSpec
	// MISSING: DataAccessSpec
	return out
}
func DataplexDataAttributeSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataplexDataAttributeSpec) *pb.DataAttribute {
	if in == nil {
		return nil
	}
	out := &pb.DataAttribute{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: ParentID
	// MISSING: AttributeCount
	// MISSING: Etag
	// MISSING: ResourceAccessSpec
	// MISSING: DataAccessSpec
	return out
}
func ResourceAccessSpec_FromProto(mapCtx *direct.MapContext, in *pb.ResourceAccessSpec) *krm.ResourceAccessSpec {
	if in == nil {
		return nil
	}
	out := &krm.ResourceAccessSpec{}
	out.Readers = in.Readers
	out.Writers = in.Writers
	out.Owners = in.Owners
	return out
}
func ResourceAccessSpec_ToProto(mapCtx *direct.MapContext, in *krm.ResourceAccessSpec) *pb.ResourceAccessSpec {
	if in == nil {
		return nil
	}
	out := &pb.ResourceAccessSpec{}
	out.Readers = in.Readers
	out.Writers = in.Writers
	out.Owners = in.Owners
	return out
}
