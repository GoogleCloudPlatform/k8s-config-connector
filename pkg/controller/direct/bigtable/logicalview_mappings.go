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

func BigtableLogicalViewSpec_FromProto(mapCtx *direct.MapContext, in *pb.LogicalView) *krmv1alpha1.BigtableLogicalViewSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BigtableLogicalViewSpec{}
	_, resourceID, _ := krmv1alpha1.ParseLogicalViewExternal(in.Name)
	out.ResourceID = &resourceID
	out.Query = &in.Query
	out.DeletionProtection = &in.DeletionProtection
	return out
}

func BigtableLogicalViewSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BigtableLogicalViewSpec) *pb.LogicalView {
	if in == nil {
		return nil
	}

	out := &pb.LogicalView{}

	if in.Query != nil {
		out.Query = *in.Query
	}

	if in.DeletionProtection != nil {
		out.DeletionProtection = *in.DeletionProtection
	}
	return out
}

func BigtableLogicalViewSpec_ToLogicalViewInfo(mapCtx *direct.MapContext, in *krmv1alpha1.BigtableLogicalViewSpec, identity *krmv1alpha1.LogicalViewIdentity) *gcp.LogicalViewInfo {
	if in == nil {
		return nil
	}
	if identity == nil {
		return nil
	}
	gcpDeletionProtection := gcp.Unprotected
	if in.DeletionProtection != nil {
		if *in.DeletionProtection {
			gcpDeletionProtection = gcp.Protected
		} else {
			gcpDeletionProtection = gcp.Unprotected
		}
	}
	out := &gcp.LogicalViewInfo{
		LogicalViewID:      identity.ID(),
		DeletionProtection: gcpDeletionProtection,
	}

	if in.Query != nil {
		out.Query = *in.Query
	}

	return out
}
