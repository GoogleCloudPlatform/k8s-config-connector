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

func BigtableGCPolicySpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.GcRule) *krm.BigtableGCPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.BigtableGCPolicySpec{}

	switch rule := in.Rule.(type) {
	case *pb.GcRule_MaxNumVersions:
		out.MaxVersion = []krm.GcpolicyMaxVersion{
			{
				Number: int(rule.MaxNumVersions),
			},
		}
	case *pb.GcRule_MaxAge:
		durationStr := direct.Duration_FromProto(mapCtx, rule.MaxAge)
		out.MaxAge = []krm.GcpolicyMaxAge{
			{
				Duration: durationStr,
			},
		}
	case *pb.GcRule_Intersection_:
		mode := "INTERSECTION"
		out.Mode = &mode
	case *pb.GcRule_Union_:
		mode := "UNION"
		out.Mode = &mode
	}

	// Note: ColumnFamily, InstanceRef, TableRef, DeletionPolicy, and GcRules
	// are not part of the standard pb.GcRule message, and are managed by the controller.
	return out
}

func BigtableGCPolicySpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BigtableGCPolicySpec) *pb.GcRule {
	if in == nil {
		return nil
	}
	out := &pb.GcRule{}

	if len(in.MaxVersion) > 0 {
		out.Rule = &pb.GcRule_MaxNumVersions{
			MaxNumVersions: int32(in.MaxVersion[0].Number),
		}
	} else if len(in.MaxAge) > 0 {
		out.Rule = &pb.GcRule_MaxAge{
			MaxAge: direct.Duration_ToProto(mapCtx, in.MaxAge[0].Duration),
		}
	}

	// Note: Manual mapping for custom structures or multi-policy (UNION/INTERSECTION)
	// can be handled here or in the direct controller reconciliation logic.
	return out
}
