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

// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.compute.v1

package compute

import (
	"sort"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeProjectMetadataSpec_FromProto(mapCtx *direct.MapContext, in *pb.Metadata) *krm.ComputeProjectMetadataSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeProjectMetadataSpec{}
	if len(in.Items) > 0 {
		out.Metadata = make(map[string]string)
		for _, item := range in.Items {
			if item == nil {
				continue
			}
			out.Metadata[item.GetKey()] = item.GetValue()
		}
	}
	return out
}

func ComputeProjectMetadataSpec_ToProto(mapCtx *direct.MapContext, in *krm.ComputeProjectMetadataSpec) *pb.Metadata {
	if in == nil {
		return nil
	}
	out := &pb.Metadata{}
	for k, v := range in.Metadata {
		out.Items = append(out.Items, &pb.Items{
			Key:   direct.PtrTo(k),
			Value: direct.PtrTo(v),
		})
	}
	sort.Slice(out.Items, func(i, j int) bool {
		return out.Items[i].GetKey() < out.Items[j].GetKey()
	})
	return out
}

func ComputeProjectMetadataStatus_FromProto(mapCtx *direct.MapContext, in *pb.Metadata) *krm.ComputeProjectMetadataStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeProjectMetadataStatus{}
	return out
}

func ComputeProjectMetadataStatus_ToProto(mapCtx *direct.MapContext, in *krm.ComputeProjectMetadataStatus) *pb.Metadata {
	if in == nil {
		return nil
	}
	out := &pb.Metadata{}
	return out
}
