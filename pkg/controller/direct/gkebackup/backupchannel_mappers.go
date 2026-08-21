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

package gkebackup

import (
	pb "cloud.google.com/go/gkebackup/apiv1/gkebackuppb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkebackup/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func GKEBackupBackupChannelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupChannel) *krm.GKEBackupBackupChannelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GKEBackupBackupChannelObservedState{}
	out.UID = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.DestinationProjectID = direct.LazyPtr(in.GetDestinationProjectId())
	return out
}

func GKEBackupBackupChannelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GKEBackupBackupChannelObservedState) *pb.BackupChannel {
	if in == nil {
		return nil
	}
	out := &pb.BackupChannel{}
	out.Uid = direct.ValueOf(in.UID)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Etag = direct.ValueOf(in.Etag)
	out.DestinationProjectId = direct.ValueOf(in.DestinationProjectID)
	return out
}
