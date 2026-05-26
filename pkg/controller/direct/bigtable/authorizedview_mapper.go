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

package bigtable

import (
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AuthorizedView_SubsetView_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AuthorizedView_SubsetView) *krm.AuthorizedView_SubsetView {
	if in == nil {
		return nil
	}
	out := &krm.AuthorizedView_SubsetView{}
	out.RowPrefixes = in.RowPrefixes
	if in.GetFamilySubsets() != nil {
		out.FamilySubsets = []krm.AuthorizedView_FamilySubsets{}
		for k, v := range in.GetFamilySubsets() {
			if v == nil {
				continue
			}
			if fs := AuthorizedView_FamilySubsets_v1beta1_FromProto(mapCtx, v); fs != nil {
				fs.Name = k
				out.FamilySubsets = append(out.FamilySubsets, *fs)
			}
		}
	}
	return out
}

func AuthorizedView_SubsetView_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.AuthorizedView_SubsetView) *pb.AuthorizedView_SubsetView {
	if in == nil {
		return nil
	}
	out := &pb.AuthorizedView_SubsetView{}
	out.RowPrefixes = in.RowPrefixes
	if in.FamilySubsets != nil {
		out.FamilySubsets = map[string]*pb.AuthorizedView_FamilySubsets{}
		for _, v := range in.FamilySubsets {
			out.FamilySubsets[v.Name] = AuthorizedView_FamilySubsets_v1beta1_ToProto(mapCtx, &v)
		}
	}
	return out
}
