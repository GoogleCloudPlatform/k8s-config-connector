package bigtable

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	pb "google.golang.org/genproto/googleapis/bigtable/admin/v2"
)

func AuthorizedView_FamilySubsets_FromProto(mapCtx *direct.MapContext, in *pb.AuthorizedView_FamilySubsets) *krm.AuthorizedView_FamilySubsets {
	if in == nil {
		return nil
	}
	out := &krm.AuthorizedView_FamilySubsets{}
	out.Qualifiers = in.Qualifiers
	out.QualifierPrefixes = in.QualifierPrefixes
	return out
}
func AuthorizedView_FamilySubsets_ToProto(mapCtx *direct.MapContext, in *krm.AuthorizedView_FamilySubsets) *pb.AuthorizedView_FamilySubsets {
	if in == nil {
		return nil
	}
	out := &pb.AuthorizedView_FamilySubsets{}
	out.Qualifiers = in.Qualifiers
	out.QualifierPrefixes = in.QualifierPrefixes
	return out
}
func AuthorizedView_SubsetView_FromProto(mapCtx *direct.MapContext, in *pb.AuthorizedView_SubsetView) *krm.AuthorizedView_SubsetView {
	if in == nil {
		return nil
	}
	out := &krm.AuthorizedView_SubsetView{}
	out.RowPrefixes = in.RowPrefixes
	// MISSING: FamilySubsets
	return out
}
func AuthorizedView_SubsetView_ToProto(mapCtx *direct.MapContext, in *krm.AuthorizedView_SubsetView) *pb.AuthorizedView_SubsetView {
	if in == nil {
		return nil
	}
	out := &pb.AuthorizedView_SubsetView{}
	out.RowPrefixes = in.RowPrefixes
	// MISSING: FamilySubsets
	return out
}

func BigtableAuthorizedViewSpec_FromProto(mapCtx *direct.MapContext, in *pb.AuthorizedView) *krm.BigtableAuthorizedViewSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigtableAuthorizedViewSpec{}
	// MISSING: Name
	out.SubsetView = AuthorizedView_SubsetView_FromProto(mapCtx, in.GetSubsetView())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.DeletionProtection = direct.LazyPtr(in.GetDeletionProtection())
	return out
}
func BigtableAuthorizedViewSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigtableAuthorizedViewSpec) *pb.AuthorizedView {
	if in == nil {
		return nil
	}
	out := &pb.AuthorizedView{}
	// MISSING: Name
	if oneof := AuthorizedView_SubsetView_ToProto(mapCtx, in.SubsetView); oneof != nil {
		out.AuthorizedView = &pb.AuthorizedView_SubsetView_{SubsetView: oneof}
	}
	out.Etag = direct.ValueOf(in.Etag)
	out.DeletionProtection = direct.ValueOf(in.DeletionProtection)
	return out
}
