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

package bigtable

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	pb "google.golang.org/genproto/googleapis/bigtable/admin/v2"
)

func SchemaBundle_FamilySubsets_FromProto(mapCtx *direct.MapContext, in *pb.SchemaBundle_FamilySubsets) *krm.SchemaBundle_FamilySubsets {
	if in == nil {
		return nil
	}
	out := &krm.SchemaBundle_FamilySubsets{}
	out.Qualifiers = in.Qualifiers
	out.QualifierPrefixes = in.QualifierPrefixes
	return out
}
func SchemaBundle_FamilySubsets_ToProto(mapCtx *direct.MapContext, in *krm.SchemaBundle_FamilySubsets) *pb.SchemaBundle_FamilySubsets {
	if in == nil {
		return nil
	}
	out := &pb.SchemaBundle_FamilySubsets{}
	out.Qualifiers = in.Qualifiers
	out.QualifierPrefixes = in.QualifierPrefixes
	return out
}
func SchemaBundle_SubsetView_FromProto(mapCtx *direct.MapContext, in *pb.SchemaBundle_SubsetView) *krm.SchemaBundle_SubsetView {
	if in == nil {
		return nil
	}
	out := &krm.SchemaBundle_SubsetView{}
	out.RowPrefixes = in.RowPrefixes
	// MISSING: FamilySubsets
	return out
}
func SchemaBundle_SubsetView_ToProto(mapCtx *direct.MapContext, in *krm.SchemaBundle_SubsetView) *pb.SchemaBundle_SubsetView {
	if in == nil {
		return nil
	}
	out := &pb.SchemaBundle_SubsetView{}
	out.RowPrefixes = in.RowPrefixes
	// MISSING: FamilySubsets
	return out
}

func BigtableSchemaBundleSpec_FromProto(mapCtx *direct.MapContext, in *pb.SchemaBundle) *krm.BigtableSchemaBundleSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigtableSchemaBundleSpec{}
	// MISSING: Name
	out.ProtoSchema = SchemaBundle_ProtoSchema_FromProto(mapCtx, in.GetSubsetView())
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func BigtableSchemaBundleSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigtableSchemaBundleSpec) *pb.SchemaBundle {
	if in == nil {
		return nil
	}
	out := &pb.SchemaBundle{}
	// MISSING: Name
	if oneof := SchemaBundle_ProtoSchema_ToProto(mapCtx, in.ProtoSchema); oneof != nil {
		out.SchemaBundle = &pb.SchemaBundle_ProtoSchema_{ProtoSchema: oneof}
	}
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
