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

package dataproc

import (
	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func InstanceGroupAutoscalingPolicyConfig_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupAutoscalingPolicyConfig) *krm.InstanceGroupAutoscalingPolicyConfig {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupAutoscalingPolicyConfig{}
	out.MinInstances = direct.LazyPtr(int64(in.GetMinInstances()))
	out.MaxInstances = direct.LazyPtr(int64(in.GetMaxInstances()))
	out.Weight = direct.LazyPtr(int64(in.GetWeight()))
	return out
}

func InstanceGroupAutoscalingPolicyConfig_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupAutoscalingPolicyConfig) *pb.InstanceGroupAutoscalingPolicyConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupAutoscalingPolicyConfig{}
	out.MinInstances = int32(direct.ValueOf(in.MinInstances))
	out.MaxInstances = int32(direct.ValueOf(in.MaxInstances))
	out.Weight = int32(direct.ValueOf(in.Weight))
	return out
}

func SecondaryInstanceGroupAutoscalingPolicyConfig_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupAutoscalingPolicyConfig) *krm.SecondaryInstanceGroupAutoscalingPolicyConfig {
	if in == nil {
		return nil
	}
	out := &krm.SecondaryInstanceGroupAutoscalingPolicyConfig{}
	out.MinInstances = direct.LazyPtr(int64(in.GetMinInstances()))
	out.MaxInstances = direct.LazyPtr(int64(in.GetMaxInstances()))
	out.Weight = direct.LazyPtr(int64(in.GetWeight()))
	return out
}

func SecondaryInstanceGroupAutoscalingPolicyConfig_ToProto(mapCtx *direct.MapContext, in *krm.SecondaryInstanceGroupAutoscalingPolicyConfig) *pb.InstanceGroupAutoscalingPolicyConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupAutoscalingPolicyConfig{}
	out.MinInstances = int32(direct.ValueOf(in.MinInstances))
	out.MaxInstances = int32(direct.ValueOf(in.MaxInstances))
	out.Weight = int32(direct.ValueOf(in.Weight))
	return out
}
