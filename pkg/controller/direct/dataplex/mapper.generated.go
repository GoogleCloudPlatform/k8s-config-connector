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
func DataTaxonomy_FromProto(mapCtx *direct.MapContext, in *pb.DataTaxonomy) *krm.DataTaxonomy {
	if in == nil {
		return nil
	}
	out := &krm.DataTaxonomy{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Labels = in.Labels
	// MISSING: AttributeCount
	out.Etag = direct.LazyPtr(in.GetEtag())
	// MISSING: ClassCount
	return out
}
func DataTaxonomy_ToProto(mapCtx *direct.MapContext, in *krm.DataTaxonomy) *pb.DataTaxonomy {
	if in == nil {
		return nil
	}
	out := &pb.DataTaxonomy{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.ValueOf(in.Description)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Labels = in.Labels
	// MISSING: AttributeCount
	out.Etag = direct.ValueOf(in.Etag)
	// MISSING: ClassCount
	return out
}
func DataTaxonomyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataTaxonomy) *krm.DataTaxonomyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataTaxonomyObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	out.AttributeCount = direct.LazyPtr(in.GetAttributeCount())
	// MISSING: Etag
	out.ClassCount = direct.LazyPtr(in.GetClassCount())
	return out
}
func DataTaxonomyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataTaxonomyObservedState) *pb.DataTaxonomy {
	if in == nil {
		return nil
	}
	out := &pb.DataTaxonomy{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	out.AttributeCount = direct.ValueOf(in.AttributeCount)
	// MISSING: Etag
	out.ClassCount = direct.ValueOf(in.ClassCount)
	return out
}
func DataplexDataTaxonomyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataTaxonomy) *krm.DataplexDataTaxonomyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataplexDataTaxonomyObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: AttributeCount
	// MISSING: Etag
	// MISSING: ClassCount
	return out
}
func DataplexDataTaxonomyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataplexDataTaxonomyObservedState) *pb.DataTaxonomy {
	if in == nil {
		return nil
	}
	out := &pb.DataTaxonomy{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: AttributeCount
	// MISSING: Etag
	// MISSING: ClassCount
	return out
}
func DataplexDataTaxonomySpec_FromProto(mapCtx *direct.MapContext, in *pb.DataTaxonomy) *krm.DataplexDataTaxonomySpec {
	if in == nil {
		return nil
	}
	out := &krm.DataplexDataTaxonomySpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: AttributeCount
	// MISSING: Etag
	// MISSING: ClassCount
	return out
}
func DataplexDataTaxonomySpec_ToProto(mapCtx *direct.MapContext, in *krm.DataplexDataTaxonomySpec) *pb.DataTaxonomy {
	if in == nil {
		return nil
	}
	out := &pb.DataTaxonomy{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: AttributeCount
	// MISSING: Etag
	// MISSING: ClassCount
	return out
}
