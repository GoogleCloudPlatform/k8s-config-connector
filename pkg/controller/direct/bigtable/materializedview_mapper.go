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

func BigtableMaterializedViewSpec_ToMaterializedViewInfo(mapCtx *direct.MapContext, in *krmv1alpha1.BigtableMaterializedViewSpec) *gcp.MaterializedViewInfo {
	if in == nil {
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
		Query:              *in.Query,
		DeletionProtection: gcpDeletionProtection,
	}
}

func BigtableMaterializedViewInfo_ToBigtableMaterializedViewSpec(mapCtx *direct.MapContext, in *gcp.MaterializedViewInfo) *krmv1alpha1.BigtableMaterializedViewSpec {
	if in == nil {
		return nil
	}

	var deletionProtection *bool
	switch in.DeletionProtection {
	case gcp.Protected:
		dp := true
		deletionProtection = &dp
	case gcp.Unprotected:
		dp := false
		deletionProtection = &dp
	}

	return &krmv1alpha1.BigtableMaterializedViewSpec{
		Query:              &in.Query,
		DeletionProtection: deletionProtection,
	}
}

func BigtableMaterializedViewInfo_ToBigtableMaterializedView(mapCtx *direct.MapContext, in *gcp.MaterializedViewInfo) *pb.MaterializedView {
	if in == nil {
		return nil
	}

	deletionProtection := false
	if in.DeletionProtection == gcp.Protected {
		deletionProtection = true
	}

	return &pb.MaterializedView{
		Query:              in.Query,
		DeletionProtection: deletionProtection,
	}
}

func BigtableMaterializedView_ToBigtableMaterializedViewInfo(mapCtx *direct.MapContext, in *pb.MaterializedView) *gcp.MaterializedViewInfo {
	if in == nil {
		return nil
	}

	deletionProtection := gcp.None
	if in.DeletionProtection {
		deletionProtection = gcp.Protected
	} else {
		deletionProtection = gcp.Unprotected
	}

	return &gcp.MaterializedViewInfo{
		Query:              in.Query,
		DeletionProtection: deletionProtection,
	}
}
