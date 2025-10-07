// Copyright 2024 Google LLC
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

package firestore

import (
	pb "cloud.google.com/go/firestore/apiv1/admin/adminpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func FirestoreDatabaseObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Database) *krm.FirestoreDatabaseObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FirestoreDatabaseObservedState{}
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.VersionRetentionPeriod = direct.Duration_FromProto(mapCtx, in.GetVersionRetentionPeriod())
	out.EarliestVersionTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEarliestVersionTime())
	out.KeyPrefix = direct.LazyPtr(in.GetKeyPrefix())
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}

func FirestoreDatabaseObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FirestoreDatabaseObservedState) *pb.Database {
	if in == nil {
		return nil
	}
	out := &pb.Database{}
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.VersionRetentionPeriod = direct.Duration_ToProto(mapCtx, in.VersionRetentionPeriod)
	out.EarliestVersionTime = direct.StringTimestamp_ToProto(mapCtx, in.EarliestVersionTime)
	out.KeyPrefix = direct.ValueOf(in.KeyPrefix)
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
