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

package dataplex

import (
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func MetadataFeed_Filters_FromProto(mapCtx *direct.MapContext, in *pb.MetadataFeed_Filters) *krm.MetadataFeed_Filters {
	if in == nil {
		return nil
	}
	out := &krm.MetadataFeed_Filters{}
	out.EntryTypeRefs = MetadataFeed_Filters_EntryTypeRefs_FromProto(mapCtx, in.EntryTypes)
	out.AspectTypeRefs = MetadataFeed_Filters_AspectTypeRefs_FromProto(mapCtx, in.AspectTypes)
	out.ChangeTypes = direct.EnumSlice_FromProto(mapCtx, in.ChangeTypes)
	return out
}

func MetadataFeed_Filters_ToProto(mapCtx *direct.MapContext, in *krm.MetadataFeed_Filters) *pb.MetadataFeed_Filters {
	if in == nil {
		return nil
	}
	out := &pb.MetadataFeed_Filters{}
	out.EntryTypes = MetadataFeed_Filters_EntryTypeRefs_ToProto(mapCtx, in.EntryTypeRefs)
	out.AspectTypes = MetadataFeed_Filters_AspectTypeRefs_ToProto(mapCtx, in.AspectTypeRefs)
	out.ChangeTypes = direct.EnumSlice_ToProto[pb.MetadataFeed_Filters_ChangeType](mapCtx, in.ChangeTypes)
	return out
}

func MetadataFeed_Scope_FromProto(mapCtx *direct.MapContext, in *pb.MetadataFeed_Scope) *krm.MetadataFeed_Scope {
	if in == nil {
		return nil
	}
	out := &krm.MetadataFeed_Scope{}
	out.OrganizationLevel = direct.LazyPtr(in.GetOrganizationLevel())
	out.ProjectRefs = MetadataFeed_Scope_ProjectRefs_FromProto(mapCtx, in.Projects)
	out.EntryGroupRefs = MetadataFeed_Scope_EntryGroupRefs_FromProto(mapCtx, in.EntryGroups)
	return out
}

func MetadataFeed_Scope_ToProto(mapCtx *direct.MapContext, in *krm.MetadataFeed_Scope) *pb.MetadataFeed_Scope {
	if in == nil {
		return nil
	}
	out := &pb.MetadataFeed_Scope{}
	out.OrganizationLevel = direct.ValueOf(in.OrganizationLevel)
	out.Projects = MetadataFeed_Scope_ProjectRefs_ToProto(mapCtx, in.ProjectRefs)
	out.EntryGroups = MetadataFeed_Scope_EntryGroupRefs_ToProto(mapCtx, in.EntryGroupRefs)
	return out
}

func MetadataFeed_Filters_EntryTypeRefs_FromProto(mapCtx *direct.MapContext, in []string) []krm.EntryTypeRef {
	if in == nil {
		return nil
	}
	out := make([]krm.EntryTypeRef, len(in))
	for i, v := range in {
		out[i] = krm.EntryTypeRef{External: v}
	}
	return out
}

func MetadataFeed_Filters_EntryTypeRefs_ToProto(mapCtx *direct.MapContext, in []krm.EntryTypeRef) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	for i, v := range in {
		out[i] = v.External
	}
	return out
}

func MetadataFeed_Filters_AspectTypeRefs_FromProto(mapCtx *direct.MapContext, in []string) []krm.AspectTypeRef {
	if in == nil {
		return nil
	}
	out := make([]krm.AspectTypeRef, len(in))
	for i, v := range in {
		out[i] = krm.AspectTypeRef{External: v}
	}
	return out
}

func MetadataFeed_Filters_AspectTypeRefs_ToProto(mapCtx *direct.MapContext, in []krm.AspectTypeRef) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	for i, v := range in {
		out[i] = v.External
	}
	return out
}

func MetadataFeed_Scope_ProjectRefs_FromProto(mapCtx *direct.MapContext, in []string) []refsv1beta1.ProjectRef {
	if in == nil {
		return nil
	}
	out := make([]refsv1beta1.ProjectRef, len(in))
	for i, v := range in {
		out[i] = refsv1beta1.ProjectRef{External: v}
	}
	return out
}

func MetadataFeed_Scope_ProjectRefs_ToProto(mapCtx *direct.MapContext, in []refsv1beta1.ProjectRef) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	for i, v := range in {
		out[i] = v.External
	}
	return out
}

func MetadataFeed_Scope_EntryGroupRefs_FromProto(mapCtx *direct.MapContext, in []string) []krm.EntryGroupRef {
	if in == nil {
		return nil
	}
	out := make([]krm.EntryGroupRef, len(in))
	for i, v := range in {
		out[i] = krm.EntryGroupRef{External: v}
	}
	return out
}

func MetadataFeed_Scope_EntryGroupRefs_ToProto(mapCtx *direct.MapContext, in []krm.EntryGroupRef) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	for i, v := range in {
		out[i] = v.External
	}
	return out
}
