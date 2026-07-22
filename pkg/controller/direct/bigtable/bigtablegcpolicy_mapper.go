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
	"fmt"

	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BigtableGCPolicySpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.GcRule) *krm.BigtableGCPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.BigtableGCPolicySpec{}

	if in.GetMaxNumVersions() != 0 {
		out.MaxVersion = []krm.GcpolicyMaxVersion{
			{
				Number: int64(in.GetMaxNumVersions()),
			},
		}
	}

	if in.GetMaxAge() != nil {
		durationStr := direct.StringDuration_FromProto(mapCtx, in.GetMaxAge())
		out.MaxAge = []krm.GcpolicyMaxAge{
			{
				Duration: durationStr,
			},
		}
	}

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
	}

	if len(in.MaxAge) > 0 {
		var dStr *string
		if in.MaxAge[0].Duration != nil {
			dStr = in.MaxAge[0].Duration
		} else if in.MaxAge[0].Days != nil {
			s := fmt.Sprintf("%dh", *in.MaxAge[0].Days*24)
			dStr = &s
		}
		if dStr != nil {
			out.Rule = &pb.GcRule_MaxAge{
				MaxAge: direct.StringDuration_ToProto(mapCtx, dStr),
			}
		}
	}

	return out
}
