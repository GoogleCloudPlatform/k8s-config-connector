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

func FirestoreDatabaseSpec_FromProto(mapCtx *direct.MapContext, in *pb.Database) *krm.FirestoreDatabaseSpec {
	if in == nil {
		return nil
	}
	out := &krm.FirestoreDatabaseSpec{}
	out.LocationID = direct.LazyPtr(in.GetLocationId())
	out.ConcurrencyMode = direct.Enum_FromProto(mapCtx, in.GetConcurrencyMode())
	out.PointInTimeRecoveryEnablement = direct.Enum_FromProto(mapCtx, in.GetPointInTimeRecoveryEnablement())
	return out
}

func FirestoreDatabaseSpec_ToProto(mapCtx *direct.MapContext, in *krm.FirestoreDatabaseSpec) *pb.Database {
	if in == nil {
		return nil
	}
	out := &pb.Database{}
	out.LocationId = direct.ValueOf(in.LocationID)
	out.ConcurrencyMode = direct.Enum_ToProto[pb.Database_ConcurrencyMode](mapCtx, in.ConcurrencyMode)
	out.PointInTimeRecoveryEnablement = direct.Enum_ToProto[pb.Database_PointInTimeRecoveryEnablement](mapCtx, in.PointInTimeRecoveryEnablement)
	return out
}
