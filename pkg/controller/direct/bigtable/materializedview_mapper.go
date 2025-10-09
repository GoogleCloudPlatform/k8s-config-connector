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

// +generated:mapper
// krm.group: bigtable.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.bigtable.admin.v2

package bigtable

import (
	gcp "cloud.google.com/go/bigtable"
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BigtableMaterializedViewSpec_FromProto(mapCtx *direct.MapContext, in *pb.MaterializedView) *krmv1alpha1.BigtableMaterializedViewSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BigtableMaterializedViewSpec{}
	_, resourceID, _ := krmv1alpha1.ParseMaterializedViewExternal(in.Name)
	out.ResourceID = &resourceID
	out.Query = &in.Query

	out.DeletionProtection = &in.DeletionProtection
	return out
}

func BigtableMaterializedViewSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BigtableMaterializedViewSpec) *pb.MaterializedView {
	if in == nil {
		return nil
	}
	out := &pb.MaterializedView{
		Query:              *in.Query,
		DeletionProtection: *in.DeletionProtection,
	}
	return out
}

func BigtableMaterializedViewSpec_ToMaterializedViewInfo(mapCtx *direct.MapContext, in *krmv1alpha1.BigtableMaterializedViewSpec, identity *krmv1alpha1.MaterializedViewIdentity) *gcp.MaterializedViewInfo {
	if in == nil || identity == nil {
		return nil
	}

	gcpDeletionProtection := gcp.None
	if in.DeletionProtection != nil {
		if *in.DeletionProtection {
			gcpDeletionProtection = gcp.Protected
		} else {
			gcpDeletionProtection = gcp.Unprotected
		}
	}

	return &gcp.MaterializedViewInfo{
		MaterializedViewID: identity.ID(),
		Query:              *in.Query,
		DeletionProtection: gcpDeletionProtection,
	}
}

func BigtableMaterializedViewInfo_ToBigtableMaterializedView(mapCtx *direct.MapContext, in *gcp.MaterializedViewInfo, identity *krmv1alpha1.MaterializedViewIdentity) *pb.MaterializedView {
	if in == nil || identity == nil {
		return nil
	}

	deletionProtection := false
	if in.DeletionProtection == gcp.Protected {
		deletionProtection = true
	}

	return &pb.MaterializedView{
		Name:               identity.ID(),
		Query:              in.Query,
		DeletionProtection: deletionProtection,
	}
}
