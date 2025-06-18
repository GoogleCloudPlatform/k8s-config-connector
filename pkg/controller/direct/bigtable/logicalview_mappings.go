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
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BigtableLogicalViewSpec_FromProto(mapCtx *direct.MapContext, in *pb.LogicalView) *krmv1alpha1.BigtableLogicalViewSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BigtableLogicalViewSpec{}
	out.Name = &in.Name
	out.Query = &in.Query
	// TODO: implement this once DeletionProtection is published in adminpb.logicalview.
	// out.DeletionProtection = direct.LazyPtr(in.DeletionProtection)
	return out
}

func BigtableLogicalViewSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BigtableLogicalViewSpec) *pb.LogicalView {
	if in == nil {
		return nil
	}
	return &pb.LogicalView{
		Name:  *in.Name,
		Query: *in.Query,
	}
	// TODO: implement this once DeletionProtection is published in adminpb.logicalview.
	// out.DeletionProtection = direct.LazyPtr(in.DeletionProtection)
}
