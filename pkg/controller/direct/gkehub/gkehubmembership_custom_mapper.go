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

package gkehub

import (
	pb "cloud.google.com/go/gkehub/apiv1beta1/gkehubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func KubernetesMetadataObservedState_FromProto(mapCtx *direct.MapContext, in *pb.KubernetesMetadata) *krm.KubernetesMetadataObservedState {
	if in == nil {
		return nil
	}
	out := &krm.KubernetesMetadataObservedState{}
	out.KubernetesAPIServerVersion = direct.LazyPtr(in.GetKubernetesApiServerVersion())
	out.NodeProviderID = direct.LazyPtr(in.GetNodeProviderId())
	out.NodeCount = direct.LazyPtr(in.GetNodeCount())
	out.VcpuCount = direct.LazyPtr(in.GetVcpuCount())
	out.MemoryMb = direct.LazyPtr(in.GetMemoryMb())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func KubernetesMetadataObservedState_ToProto(mapCtx *direct.MapContext, in *krm.KubernetesMetadataObservedState) *pb.KubernetesMetadata {
	if in == nil {
		return nil
	}
	out := &pb.KubernetesMetadata{}
	out.KubernetesApiServerVersion = direct.ValueOf(in.KubernetesAPIServerVersion)
	out.NodeProviderId = direct.ValueOf(in.NodeProviderID)
	out.NodeCount = direct.ValueOf(in.NodeCount)
	out.VcpuCount = direct.ValueOf(in.VcpuCount)
	out.MemoryMb = direct.ValueOf(in.MemoryMb)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
