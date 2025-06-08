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
// krm.group: storage.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.storage.control.v2

package storagecontrol

import (
	pb "cloud.google.com/go/storage/control/apiv2/controlpb"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func StorageAnywhereCacheObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AnywhereCache) *krm.StorageAnywhereCacheObservedState {
	if in == nil {
		return nil
	}
	out := &krm.StorageAnywhereCacheObservedState{}
	out.State = direct.LazyPtr(in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())

	pendingUpdate := in.GetPendingUpdate()
	out.PendingUpdate = &pendingUpdate
	return out
}
func StorageAnywhereCacheObservedState_ToProto(mapCtx *direct.MapContext, in *krm.StorageAnywhereCacheObservedState) *pb.AnywhereCache {
	if in == nil {
		return nil
	}
	out := &pb.AnywhereCache{}
	out.State = direct.ValueOf(in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.PendingUpdate = direct.ValueOf(in.PendingUpdate)
	return out
}
func StorageAnywhereCacheSpec_FromProto(mapCtx *direct.MapContext, in *pb.AnywhereCache) *krm.StorageAnywhereCacheSpec {
	if in == nil {
		return nil
	}
	out := &krm.StorageAnywhereCacheSpec{}

	if in.GetName() != "" {
		parent, resourceID, err := krm.ParseAnywhereCacheExternal(in.GetName())
		if err != nil {
			mapCtx.Errorf("Error while parsing name %s %w", in.GetName(), err)
			return nil
		}

		out.BucketRef = &refs.StorageBucketRef{External: parent.String()}
		out.ResourceID = &resourceID
	}
	out.Zone = direct.LazyPtr(in.GetZone())
	out.Ttl = direct.StringDuration_FromProto(mapCtx, in.GetTtl())
	out.AdmissionPolicy = direct.LazyPtr(in.GetAdmissionPolicy())
	return out
}
func StorageAnywhereCacheSpec_ToProto(mapCtx *direct.MapContext, in *krm.StorageAnywhereCacheSpec) *pb.AnywhereCache {
	if in == nil {
		return nil
	}
	out := &pb.AnywhereCache{}
	out.Zone = direct.ValueOf(in.Zone)
	out.Ttl = direct.StringDuration_ToProto(mapCtx, in.Ttl)
	out.AdmissionPolicy = direct.ValueOf(in.AdmissionPolicy)
	return out
}
