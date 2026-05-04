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

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NodeGroupProjectMap_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ShareSettingsProjectConfig) *krm.NodeGroupProjectMap {
	if in == nil {
		return nil
	}
	out := &krm.NodeGroupProjectMap{}
	if in.ProjectId != nil {
		out.ProjectIdRef = krm.ComputeNodeGroupProjectRef{External: *in.ProjectId}
	}
	return out
}

func NodeGroupProjectMap_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.NodeGroupProjectMap) *pb.ShareSettingsProjectConfig {
	if in == nil {
		return nil
	}
	out := &pb.ShareSettingsProjectConfig{}
	if in.ProjectIdRef.External != "" {
		out.ProjectId = &in.ProjectIdRef.External
	}
	return out
}

func NodeGroupShareSettings_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ShareSettings) *krm.NodeGroupShareSettings {
	if in == nil {
		return nil
	}
	out := &krm.NodeGroupShareSettings{}
	if in.ProjectMap != nil {
		for k, v := range in.ProjectMap {
			item := NodeGroupProjectMap_v1beta1_FromProto(mapCtx, v)
			item.IdRef = krm.ComputeNodeGroupProjectRef{External: k}
			out.ProjectMap = append(out.ProjectMap, *item)
		}
	}
	out.ShareType = in.ShareType
	return out
}

func NodeGroupShareSettings_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.NodeGroupShareSettings) *pb.ShareSettings {
	if in == nil {
		return nil
	}
	out := &pb.ShareSettings{}
	if in.ProjectMap != nil {
		out.ProjectMap = make(map[string]*pb.ShareSettingsProjectConfig)
		for _, item := range in.ProjectMap {
			val := NodeGroupProjectMap_v1beta1_ToProto(mapCtx, &item)
			out.ProjectMap[item.IdRef.External] = val
		}
	}
	out.ShareType = in.ShareType
	return out
}
