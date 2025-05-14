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
	"fmt"

	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BigtableAppProfileSpec_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile) *krmv1beta1.BigtableAppProfileSpec {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.BigtableAppProfileSpec{}

	// MISSING: Name
	// MISSING: Etag
	out.Description = direct.LazyPtr(in.GetDescription())
	var isMultiClusterRouting = new(bool)
	if multiClusterRouting := in.GetMultiClusterRoutingUseAny(); multiClusterRouting != nil {
		*isMultiClusterRouting = true
		clusterIds := in.GetMultiClusterRoutingUseAny().ClusterIds
		out.MultiClusterRoutingClusterIds = clusterIds
		rowAffinity := in.GetMultiClusterRoutingUseAny().GetAffinity()
		if rowAffinity != nil {
			*out.MultiClusterRoutingUseAnyRowAffinity = true
		}
	}
	out.MultiClusterRoutingUseAny = isMultiClusterRouting
	out.SingleClusterRouting = AppProfile_SingleClusterRouting_FromProto(mapCtx, in.GetSingleClusterRouting())
	// MISSING: Priority
	out.DataBoostIsolationReadOnly = AppProfile_DataBoostIsolationReadOnly_FromProto(mapCtx, in.GetDataBoostIsolationReadOnly())
	out.StandardIsolation = AppProfile_StandardIsolation_FromProto(mapCtx, in.GetStandardIsolation())
	return out
}
func BigtableAppProfileSpec_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.BigtableAppProfileSpec) *pb.AppProfile {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile{}
	// MISSING: Name
	out.Name = direct.ValueOf(in.ResourceID)

	// MISSING: Etag
	out.Description = direct.ValueOf(in.Description)
	if oneof := in.MultiClusterRoutingUseAny; oneof != nil && *oneof {
		clusterIds := in.MultiClusterRoutingClusterIds
		var affinity *pb.AppProfile_MultiClusterRoutingUseAny_RowAffinity_ = nil
		fmt.Println("CHKPT toProto 1")
		if rowAffinity := in.MultiClusterRoutingUseAnyRowAffinity; rowAffinity != nil && *rowAffinity {
			affinity = &pb.AppProfile_MultiClusterRoutingUseAny_RowAffinity_{}
			fmt.Println("CHKPT toProto 2")
			fmt.Println(affinity)
		}
		out.RoutingPolicy = &pb.AppProfile_MultiClusterRoutingUseAny_{
			MultiClusterRoutingUseAny: &pb.AppProfile_MultiClusterRoutingUseAny{
				ClusterIds: clusterIds,
				Affinity:   affinity,
			},
		}
	}
	if oneof := AppProfile_SingleClusterRouting_ToProto(mapCtx, in.SingleClusterRouting); oneof != nil {
		out.RoutingPolicy = &pb.AppProfile_SingleClusterRouting_{SingleClusterRouting: oneof}
	}
	// MISSING: Priority
	if oneof := AppProfile_StandardIsolation_ToProto(mapCtx, in.StandardIsolation); oneof != nil {
		out.Isolation = &pb.AppProfile_StandardIsolation_{StandardIsolation: oneof}
	}
	if oneof := AppProfile_DataBoostIsolationReadOnly_ToProto(mapCtx, in.DataBoostIsolationReadOnly); oneof != nil {
		out.Isolation = &pb.AppProfile_DataBoostIsolationReadOnly_{DataBoostIsolationReadOnly: oneof}
	}
	fmt.Println("CHKPT toProto 3")
	fmt.Println(out.RoutingPolicy)
	return out
}
