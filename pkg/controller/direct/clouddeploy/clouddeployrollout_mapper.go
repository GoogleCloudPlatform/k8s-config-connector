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

package clouddeploy

import (
	pb "cloud.google.com/go/deploy/apiv1/deploypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddeploy/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CloudDeployRolloutSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Rollout) *krm.CloudDeployRolloutSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudDeployRolloutSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	if in.Annotations != nil {
		out.Annotations = make(map[string]string)
		for k, v := range in.Annotations {
			out.Annotations[k] = v
		}
	}
	if in.GetTargetId() != "" {
		out.TargetRef = &krm.CloudDeployTargetRef{External: in.GetTargetId()}
	}
	return out
}

func CloudDeployRolloutSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.CloudDeployRolloutSpec) *pb.Rollout {
	if in == nil {
		return nil
	}
	out := &pb.Rollout{}
	out.Description = direct.ValueOf(in.Description)
	if in.Annotations != nil {
		out.Annotations = make(map[string]string)
		for k, v := range in.Annotations {
			out.Annotations[k] = v
		}
	}
	if in.TargetRef != nil {
		out.TargetId = lastSegment(in.TargetRef.External)
	}
	return out
}
